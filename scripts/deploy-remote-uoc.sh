#!/usr/bin/env bash

# Global
ACCOUNT_CMD="cd .ethereum; ./bin/geth --datadir . --verbosity 0 account list | grep -o '[0-9a-fA-F]\{40\}' | head -1"
BOOTNODE="enode://8f9079250bf8a523f46a438560433a962ce2d3f9ba20e99145ac1b52fba27216eef5b6774d6cb9dd83fe564d968ed1725c80e031f988d316be05881419bce6e2@192.168.1.2:30301"

# Prepare deployment environment (remove previous data and initialize genesis block)
CMD="pkill bootnode; pkill geth; pkill nxgent; cd .ethereum; rm -rf geth; rm -f geth.ipc geth.out history nxgent.out; \
./bin/geth --datadir . init genesis.json"
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.1 "$CMD" > /dev/null 2>&1 # PC1
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.2 "$CMD" > /dev/null 2>&1 # PC2
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.3 "$CMD" > /dev/null 2>&1 # PC3
sshpass -p "vyos" ssh vyos@172.16.2.4 "$CMD"  > /dev/null 2>&1 # R1
sshpass -p "vyos" ssh vyos@172.16.2.5 "$CMD"  > /dev/null 2>&1 # R2

# Create Ethereum accounts
#CMD="cd .ethereum; ./bin/geth --datadir . account new --password secret.txt"
#ssh -i ~/.ssh/uoc_rsa kison@172.16.2.1 "$CMD" > /dev/null 2>&1 # PC1
#ssh -i ~/.ssh/uoc_rsa kison@172.16.2.2 "$CMD" > /dev/null 2>&1 # PC2
#ssh -i ~/.ssh/uoc_rsa kison@172.16.2.3 "$CMD" > /dev/null 2>&1 # PC3
#sshpass -p "vyos" ssh vyos@172.16.2.4 "$CMD"  > /dev/null 2>&1 # R1
#sshpass -p "vyos" ssh vyos@172.16.2.5 "$CMD"  > /dev/null 2>&1 # R2

# Execute blockchain bootnode
CMD="cd .ethereum; nohup ./bin/bootnode -nodekey boot.key > /dev/null 2>&1 &"
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.1 "$CMD" > /dev/null 2>&1 # PC1

# Execute sealer node (1)
ACCOUNT=`ssh -i ~/.ssh/uoc_rsa kison@172.16.2.1 "$ACCOUNT_CMD"` # PC1
CMD="cd .ethereum; nohup ./bin/geth --datadir . --networkid 12345 --bootnodes "$BOOTNODE" \
--authrpc.port 8551 --http --allow-insecure-unlock --unlock "0x$ACCOUNT" --password secret.txt \
--miner.etherbase "0x$ACCOUNT" --miner.gaslimit 4294967295 --mine > geth.out 2>&1 &"
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.1 "$CMD" # PC1

# Execute regular node (2)
ACCOUNT=`ssh -i ~/.ssh/uoc_rsa kison@172.16.2.2 "$ACCOUNT_CMD"` # PC2
CMD="cd .ethereum; nohup ./bin/geth --datadir . --networkid 12345 --bootnodes "$BOOTNODE" \
--authrpc.port 8551 --unlock "0x$ACCOUNT" --password secret.txt > geth.out 2>&1 &"
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.2 "$CMD" # PC2

# Execute regular node (3)
ACCOUNT=`ssh -i ~/.ssh/uoc_rsa kison@172.16.2.3 "$ACCOUNT_CMD"` # PC3
CMD="cd .ethereum; nohup ./bin/geth --datadir . --networkid 12345 --bootnodes "$BOOTNODE" \
--authrpc.port 8551 --unlock "0x$ACCOUNT" --password secret.txt > geth.out 2>&1 &"
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.3 "$CMD" # PC3

# Execute regular node (4)
ACCOUNT=`sshpass -p "vyos" ssh vyos@172.16.2.4 "$ACCOUNT_CMD"` # R1
CMD="cd .ethereum; nohup ./bin/geth --datadir . --networkid 12345 --bootnodes "$BOOTNODE" \
--authrpc.port 8551 --unlock "0x$ACCOUNT" --password secret.txt --rpc.evmtimeout 0 > geth.out 2>&1 &"
sshpass -p "vyos" ssh vyos@172.16.2.4 "$CMD" # R1

# Execute regular node (5)
ACCOUNT=`sshpass -p "vyos" ssh vyos@172.16.2.5 "$ACCOUNT_CMD"` # R2
CMD="cd .ethereum; nohup ./bin/geth --datadir . --networkid 12345 --bootnodes "$BOOTNODE" \
--authrpc.port 8551 --unlock "0x$ACCOUNT" --password secret.txt --rpc.evmtimeout 0 > geth.out 2>&1 &"
sshpass -p "vyos" ssh vyos@172.16.2.5 "$CMD" # R2

# Compile NxGenT
env GOOS=linux GOARCH=amd64 go build -o bin/linux/nxgent main.go >/dev/null 2>&1

# Upload NxGenT
scp -i ~/.ssh/uoc_rsa bin/linux/nxgent testbed_uoc/.env kison@172.16.2.1:/home/kison/.ethereum > /dev/null 2>&1 # PC1
scp -i ~/.ssh/uoc_rsa bin/linux/nxgent testbed_uoc/.env kison@172.16.2.2:/home/kison/.ethereum > /dev/null 2>&1 # PC2
scp -i ~/.ssh/uoc_rsa bin/linux/nxgent testbed_uoc/.env kison@172.16.2.3:/home/kison/.ethereum > /dev/null 2>&1 # PC3
sshpass -p "vyos" scp bin/linux/nxgent testbed_uoc/.env vyos@172.16.2.4:/home/vyos/.ethereum > /dev/null 2>&1 # R1
sshpass -p "vyos" scp bin/linux/nxgent testbed_uoc/.env vyos@172.16.2.5:/home/vyos/.ethereum > /dev/null 2>&1 # R2

# Deploy NxGenT smart contract
CMD="cd .ethereum; ./nxgent deploy ."
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.1 "$CMD" # PC1

sleep 5

# Register trust nodes
CMD="cd .ethereum; ./nxgent register ."
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.1 "$CMD" # PC1
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.2 "$CMD" # PC2
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.3 "$CMD" # PC3
sshpass -p "vyos" ssh vyos@172.16.2.4 "$CMD"  # R1
sshpass -p "vyos" ssh vyos@172.16.2.5 "$CMD"  # R2

sleep 5

# Run NxGenT
CMD="cd .ethereum; nohup ./nxgent run . 8888 > nxgent.out 2>&1 &"
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.1 "$CMD" # PC1
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.2 "$CMD" # PC2
ssh -i ~/.ssh/uoc_rsa kison@172.16.2.3 "$CMD" # PC3
sshpass -p "vyos" ssh vyos@172.16.2.4 "$CMD"  # R1
sshpass -p "vyos" ssh vyos@172.16.2.5 "$CMD"  # R2

echo "--> Done!"
