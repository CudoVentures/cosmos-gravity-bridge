#!/bin/bash
npx ts-node \
contract-deployer.ts \
--cosmos-node="${COSMOS_NODE}" \
--eth-node="${ETH_NODE}" \
--eth-privkey="${ETH_PRIV_KEY_HEX}" \
--contract="artifacts/contracts/Gravity.sol/Gravity.json" \
--test-mode=true \
--cudos-access-control="${CUDOS_ACCESS_CONTROL_ADDRESS}" \
--cudos-token-address="${CUDOS_TOKEN_ADDRESS}"