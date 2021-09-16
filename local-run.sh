ADDRESS_PREFIX="cudos"
FEES="100acudos"
GRPC="http://localhost:9090"
ETHRPC="http://104.198.157.197:8545"
CONTRACT_ADDR="0x9fdE6D55dDa637806DbF016a03B6970613630333"
COSMOS_ORCH_MNEMONIC="pencil spirit also middle brave celery obtain merge hurt rocket slice damp account actor fire first science organ charge ring vessel square extra general"
ETH_PRIV_KEY_HEX="ed1eb2a53e4f15d2d60009b6b4be1d999cb0f64712678d09de972ce681e82882"

cd ./orchestrator && cargo build && cp ./target/debug/gbt /usr/local/bin/gbt

gbt --address-prefix="$ADDRESS_PREFIX" orchestrator \
    --fees="$FEES" \
    --cosmos-grpc="$GRPC" \
    --ethereum-rpc="$ETHRPC" \
    --gravity-contract-address="$CONTRACT_ADDR" \
    --ethereum-key="${ETH_PRIV_KEY_HEX}" \
    --cosmos-phrase="$COSMOS_ORCH_MNEMONIC"