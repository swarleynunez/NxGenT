#!/usr/bin/env bash

# Settings
WORKDIR="deployment"
TRUST_NODES=10
CHAIN_ID=12345
BOOTNODE="enode://8f9079250bf8a523f46a438560433a962ce2d3f9ba20e99145ac1b52fba27216eef5b6774d6cb9dd83fe564d968ed1725c80e031f988d316be05881419bce6e2@127.0.0.1:30301"
GETH_PORT=30303
GETH_RPC_PORT=8551
RPC_PORT=20202
ETH_BALANCE=1000000000000000000000 # 1000 ETH for each trust node

# Prepare deployment environment
mkdir -p $WORKDIR
killall bootnode geth nxgent >/dev/null 2>&1

# For each trust node
for i in $(seq 1 $TRUST_NODES); do
  NODE_DIR="$(pwd)/$WORKDIR/N$i"
  mkdir -p "$NODE_DIR/keystore"

  # Create Ethereum account
  if ! [ $(ls -A "$NODE_DIR"/keystore) ]; then
    ./bin/geth --datadir "$NODE_DIR" account new --password secret.txt >/dev/null 2>&1
  fi

  # Remove previous data
  rm -rf "$NODE_DIR"/geth > /dev/null 2>&1
  rm -f "$NODE_DIR"/geth.ipc "$NODE_DIR"/history "$NODE_DIR"/geth.out "$NODE_DIR"/nxgent.out "$NODE_DIR"/exp.out > /dev/null 2>&1

  # Initialize environment
  ./bin/geth --datadir "$NODE_DIR" init genesis.json >/dev/null 2>&1
  touch "$NODE_DIR"/nxgent.out
done

# Execute blockchain bootnode
nohup ./bin/bootnode -nodekey boot.key >/dev/null 2>&1 &

# Execute blockchain sealer node
SEALER_NODE_DIR="$(pwd)/$WORKDIR/N1"
SEALER_ACCOUNT=$(./bin/geth --datadir "$SEALER_NODE_DIR" --verbosity 0 account list | grep -o "[0-9a-fA-F]\{40\}" | head -1)
nohup ./bin/geth --datadir "$SEALER_NODE_DIR" --networkid "$CHAIN_ID" \
  --bootnodes "$BOOTNODE" --port "$GETH_PORT" --authrpc.port "$GETH_RPC_PORT" \
  --http --allow-insecure-unlock --unlock "0x$SEALER_ACCOUNT" --password secret.txt \
  --miner.etherbase "0x$SEALER_ACCOUNT" --mine > "$SEALER_NODE_DIR/geth.out" 2>&1 &
#./bin/nxgent monitor $! "$SEALER_NODE_DIR/nxgent.out" "Geth_N1" >> "$SEALER_NODE_DIR/exp.out" 2>&1 &

# Next ports
GETH_PORT=$((GETH_PORT + 1))
GETH_RPC_PORT=$((GETH_RPC_PORT + 1))

sleep 3

# For each blockchain regular node
for i in $(seq 2 $TRUST_NODES); do
  NODE_DIR="$(pwd)/$WORKDIR/N$i"

  # Send ETH from the sealer
  ACCOUNT=$(./bin/geth --datadir "$NODE_DIR" --verbosity 0 account list | grep -o "[0-9a-fA-F]\{40\}" | head -1)
  TXN="eth.sendTransaction({from: \"0x$SEALER_ACCOUNT\", to: \"0x$ACCOUNT\", value: $ETH_BALANCE})"
  ./bin/geth --datadir "$(pwd)/$WORKDIR/N1" attach --exec "$TXN" >/dev/null 2>&1

  # Execute blockchain regular node
  nohup ./bin/geth --datadir "$NODE_DIR" --networkid "$CHAIN_ID" \
    --bootnodes "$BOOTNODE" --port "$GETH_PORT" --authrpc.port "$GETH_RPC_PORT" \
    --unlock "0x$ACCOUNT" --password secret.txt > "$NODE_DIR/geth.out" 2>&1 &
  #./bin/nxgent monitor $! "$SEALER_NODE_DIR/nxgent.out" "Geth_N$i" >> "$SEALER_NODE_DIR/exp.out" 2>&1 &

  # Next ports
  GETH_PORT=$((GETH_PORT + 1))
  GETH_RPC_PORT=$((GETH_RPC_PORT + 1))
done

# Initialize NxGenT
for i in $(seq 1 $TRUST_NODES); do
  NODE_DIR="$(pwd)/$WORKDIR/N$i"

  # Deploy smart contract
  if [ "$i" == 1 ]; then
    ./bin/nxgent deploy "$NODE_DIR"
    sleep 15
  fi

  # Register trust node
  ./bin/nxgent register "$NODE_DIR" "192.168.0.$i"
done

sleep 3

# Run NxGenT
for i in $(seq 1 $TRUST_NODES); do
  NODE_DIR="$(pwd)/$WORKDIR/N$i"

  ./bin/nxgent run "$NODE_DIR" $RPC_PORT >> "$NODE_DIR/nxgent.out" 2>&1 &
  #./bin/nxgent monitor $! "$SEALER_NODE_DIR/nxgent.out" "NxGenT_N$i" >> "$SEALER_NODE_DIR/exp.out" 2>&1 &

  # Next ports
  RPC_PORT=$((RPC_PORT + 1))
done

echo "--> Done: "$(date +%s)
