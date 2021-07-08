Official Gravity Bridge readme: https://github.com/althea-net/cosmos-gravity-bridge/blob/main/readme.md

Current fees for each transaction:

Acudos transfer from Cudos network -> Ethereum network
 - acudos fee taken from ALL cudos orchestrator wallets
 - ethereum tx fee taken from ALL ethereum orchestrator wallets
 - 1 acudos given to contract deployer ethereum wallet
 - fee taken from cudos sender wallet

Acudos transfer from Ethereum network -> Cudos network
 - acudos fee taken from ALL cudos orchestrator wallets
 - ethereum tx fee taken from senders' ethereum wallet

If more than 1/3 of the orchestrator cudos wallets don't have sufficient funds for the fees, or have stopped their orchestrator relayers, the transaction can't pass.

Link to a better explanation of the fees and incentives: https://github.com/althea-net/cosmos-gravity-bridge/blob/main/docs/design/incentives.md