# Movement Bridge Script

This script is used for bridge token from other chain to movement.

## Pre-requirement:

The requirement to bridge token is that token is configured with layerzero oft adapter on source chain and peer with oft or oft adapter on movement.

## How to use:

- clone repo:
```bash
git clone https://github.com/kitelabs-io/movement-bridge-script
cd movement-bridge-script
```
- setup env and run script:
```bash
export RPC_URL=evm_source_chain_rpc
export ADAPTER_ADDRESS=oft_adapter_on_source_chain
export BRIDGE_AMOUNT=amount_in_wei_on_source_chain
export MOVE_ADDRESS=recipient_address 
export PRIVATE_KEY=evm_source_chain_private_key

go run main.go
```

Some configured `ADAPTER_ADDRESS` on ethereum mainnet:

- `0xf1df43a3053cd18e477233b59a25fc483c2cbe0f`: MOVE adapter for the MOVE token (`0x3073f7aAA4DB83f95e9FFf17424F71D4751a3073`) with decimals = 8
- `0xc209a627a7b0a19f16a963d9f7281667a2d9eff2`: USDC adapter for the USC token(`0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48`) with decimals =  6
- `0x5e87d7e75b272fb7150b4d1a05afb6bd71474950`: USDT adapter for the USDT token(`0xdAC17F958D2ee523a2206206994597C13D831ec7`) with decimals = 6
- `0x06E01cB086fea9C644a2C105A9F20cfC21A526e8`: WETH adapter for the WETH token(`0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2`) with decimals = 18
