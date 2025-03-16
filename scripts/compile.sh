#!/usr/bin/env bash

# Compile Solidity contracts (solc 0.8.25)
./bin/macos/solc --evm-version paris --base-path / --include-path node_modules --optimize --optimize-runs=1000000 \
--abi --bin -o bin/contracts --overwrite contracts/TrustManager.sol

# Generate Go bindings (abigen 1.14.0)
./bin/macos/abigen --abi bin/contracts/TrustManager.abi --bin bin/contracts/TrustManager.bin --type TrustManager --pkg bindings --out core/bindings/trust-manager.go
./bin/macos/abigen --abi bin/contracts/TrustNode.abi --type TrustNode --pkg bindings --out core/bindings/trust-node.go
