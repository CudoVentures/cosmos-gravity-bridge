npx ts-node \
contract-deployer.ts \
--cosmos-node-api="${COSMOS_NODE_API}" \
--cosmos-node-cometbft-rpc="${COSMOS_NODE_COMETBFT_RPC}" \
--eth-node="${ETH_NODE}" \
--eth-privkey="${ETH_PRIV_KEY_HEX}" \
--contract="artifacts/contracts/Gravity.sol/Gravity.json" \
--test-mode=false \
--cudos-access-control="${CUDOS_ACCESS_CONTROL_ADDRESS}" \
--cudos-token-address="${CUDOS_TOKEN_ADDRESS}"
