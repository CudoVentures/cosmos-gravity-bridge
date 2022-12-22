#!/bin/bash
npx ts-node \
contract-deployer.ts \
--cosmos-node="http://localhost:26657" \
--eth-node="http://localhost:8545" \
--eth-privkey="0xb1bab011e03a9862664706fc3bbaa1b16651528e5f0e7fbfcbfdd8be302a13e7" \
--contract="artifacts/contracts/Gravity.sol/Gravity.json" \
--test-mode=true \
--cudos-access-control="0x0412C7c846bb6b7DC462CF6B453f76D8440b2609" \
--cudos-token-address="0x9676519d99E390A180Ab1445d5d857E3f6869065"