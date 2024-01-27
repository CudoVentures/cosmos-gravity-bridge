set -e
trap 'kill $(jobs -p) 2>/dev/null' EXIT

cd ../solidity


echo Starting EVM chain
npm exec hardhat node >../baum/hardhat.log 2>&1 &

echo Prepping Cosmos chain
cudos node init >../baum/cosmos.log 2>&1

echo Starting Cosmos chain
cudos node run >>../baum/cosmos.log 2>&1 &

echo Waiting for no reason...
sleep 5

echo Deploying contracts
eval $(
  npx ts-node contract-deployer.ts \
  --cudos-access-control=0x25d16867e01197d048c1433e4335edeff43cd75b \
  --cudos-token-address=0x0 \
  --eth-node=http://127.0.0.1:8545/ \
  --eth-privkey=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
  --cosmos-node=http://127.0.0.1:26657/ \
  --contract=Gravity.json \
  --test-mode=true
)

echo Token is at $TOKEN
echo Bridge is at $BRIDGE

echo Starting orchestrator
../baum/gbt -a cudos orchestrator -g $BRIDGE >../baum/orchestrator.log 2>&1 &

echo "Sending test Cosmos -> ETH transaction"
cudos node exec -- tx gravity send-to-eth 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 100000000acudos 100000000acudos --from account0 -y

echo Everything\'s started up...

while true; do
  sleep 60
done
