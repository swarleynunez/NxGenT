#!/usr/bin/env bash

# Settings
TRUST_NODES=15
CHAIN_ID=12345
BOOTNODE="enode://8f9079250bf8a523f46a438560433a962ce2d3f9ba20e99145ac1b52fba27216eef5b6774d6cb9dd83fe564d968ed1725c80e031f988d316be05881419bce6e2@172.18.0.2:30301"
ETH_BALANCE=1000000000000000000000 # 1000 ETH for each trust node

# Prepare deployment environment
docker rm -f $(docker ps -aq) > /dev/null 2>&1
docker volume rm -f $(docker volume ls -q) > /dev/null 2>&1
docker build -t trust-bootnode --target trust-bootnode . > /dev/null 2>&1
docker build -t trust-node --target trust-node . > /dev/null 2>&1
docker network create 6g-network > /dev/null 2>&1

# For each trust node
for i in $(seq 1 $TRUST_NODES); do
  NODE_DIR="$(pwd)/testbed_local/N$i"
  mkdir -p "$NODE_DIR/keystore"

  # Create Ethereum account
  if ! [ "$(ls -A "$NODE_DIR"/keystore)" ]; then
    ./bin/macos/geth --datadir "$NODE_DIR" account new --password testbed_local/secret.txt > /dev/null 2>&1
  fi

  # Remove previous data
  rm -rf "$NODE_DIR"/geth "$NODE_DIR"/exp > /dev/null 2>&1
  rm -f "$NODE_DIR"/history "$NODE_DIR"/nxgent.out > /dev/null 2>&1

  # Initialize environment
  ./bin/macos/geth --datadir "$NODE_DIR" init testbed_local/genesis.json > /dev/null 2>&1
  mkdir -p "$NODE_DIR/exp"
done

# Execute blockchain bootnode
docker run -d --name trust-bootnode --network 6g-network trust-bootnode > /dev/null 2>&1

# Execute blockchain sealer node
SEALER_ACCOUNT=$(./bin/macos/geth --datadir testbed_local/N1 --verbosity 0 account list | grep -o "[0-9a-fA-F]\{40\}" | head -1)
docker run -d --name trust-node-1 --network 6g-network \
  -v "$(pwd)"/testbed_local/N1/geth:/trust/.ethereum/geth \
  -v "$(pwd)"/testbed_local/N1/keystore:/trust/.ethereum/keystore \
  -v "$(pwd)"/testbed_local/N1/exp:/trust/exp \
  -p 8545:8545 \
  -p 8888:8888 \
  trust-node ./geth --datadir .ethereum --networkid "$CHAIN_ID" --bootnodes "$BOOTNODE" --authrpc.port 8551 \
  --http --http.addr 0.0.0.0 --allow-insecure-unlock \
  --unlock "0x$SEALER_ACCOUNT" --password secret.txt \
  --miner.etherbase "0x$SEALER_ACCOUNT" --miner.gaslimit 4294967295 --mine > /dev/null 2>&1

sleep 3

# For each blockchain regular node
for i in $(seq 2 $TRUST_NODES); do
  # Send ETH from the sealer
  ACCOUNT=$(./bin/macos/geth --datadir "testbed_local/N$i" --verbosity 0 account list | grep -o "[0-9a-fA-F]\{40\}" | head -1)
  TXN="eth.sendTransaction({from: \"0x$SEALER_ACCOUNT\", to: \"0x$ACCOUNT\", value: $ETH_BALANCE})"
  docker exec trust-node-1 ./geth --datadir .ethereum attach --exec "$TXN" > /dev/null 2>&1

  # Execute blockchain regular node
  docker run -d --name "trust-node-$i" --network 6g-network \
    -v "$(pwd)/testbed_local/N$i/geth":/trust/.ethereum/geth \
    -v "$(pwd)/testbed_local/N$i/keystore":/trust/.ethereum/keystore \
    -v "$(pwd)/testbed_local/N$i/exp":/trust/exp \
    trust-node ./geth --datadir .ethereum --networkid "$CHAIN_ID" --bootnodes "$BOOTNODE" --authrpc.port 8551 \
    --unlock "0x$ACCOUNT" --password secret.txt > /dev/null 2>&1
done

# Compile trust client
env GOOS=linux GOARCH=arm64 go build -o bin/linux/nxgent main.go

# Compile trust client (cross-compile MacOS-Linux)
#docker run --rm --name trust-compile -v $(pwd):/trust -w /trust trust-compile env GOOS=linux GOARCH=arm64 CGO_ENABLED=1 go build -o bin/linux/nxgent main.go

# Initialize trust system
for i in $(seq 1 $TRUST_NODES); do
  # Copy files into the container
  docker cp bin/linux/nxgent "trust-node-$i":/trust > /dev/null 2>&1
  docker cp testbed_local/.env "trust-node-$i":/trust > /dev/null 2>&1
  docker cp scripts/res-monitor.sh "trust-node-$i":/trust > /dev/null 2>&1

  # Deploy smart contract
  if [ "$i" == 1 ]; then
    docker exec trust-node-1 ./nxgent deploy

    # Get address of the deployed contract
    docker cp trust-node-1:/trust/.env testbed_local > /dev/null 2>&1

    sleep 10
  fi

  # Register trust node
  docker exec "trust-node-$i" ./nxgent register
done

sleep 3

for i in $(seq 1 $TRUST_NODES); do
  # Run trust client
  docker exec "trust-node-$i" ./nxgent run > "testbed_local/N$i/nxgent.out" 2>&1 &
done

sleep 300

for i in $(seq 1 $TRUST_NODES); do
  # Run resource monitor
  docker exec "trust-node-$i" ./res-monitor-local.sh > /dev/null 2>&1 &
done

sleep 180
echo "--> Done: "$(date +%s)
