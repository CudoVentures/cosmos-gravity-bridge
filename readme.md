![Gravity Bridge](./gravity-bridge.svg)

Gravity bridge is Cosmos <-> Ethereum bridge designed to run on the [Cosmos SDK blockchains](https://github.com/cosmos/cosmos-sdk) like the [Cosmos Hub](https://github.com/cosmos/gaia) focused on maximum design simplicity and efficiency.

Gravity is currently can transfer ERC20 assets originating on Cosmos or Ethereum to and from Ethereum, as well as move Cosmos assets to Ethereum as ERC20 representations.

## Documentation

### High level documentation

To understand Gravity at a high level, read [this blog post](https://blog.althea.net/how-gravity-works/). It is accessible and more concise than the rest of these docs, but does not cover every detail.

### Code documentation

This documentation lives with the code it references and helps to understand the functions and data structures involved. This is useful if you are reviewing or working on the code.

[Solidity Ethereum contract documentation](https://github.com/althea-net/cosmos-gravity-bridge/blob/main/solidity/contracts/contract-explanation.md)

[Go Cosmos module documentation](https://github.com/althea-net/cosmos-gravity-bridge/tree/main/module/x/gravity/spec)

### Specs

These specs cover specific areas of the bridge that a lot of thought went into. They explore the tradeoffs involved and decisions made.

[slashing-spec](/spec/slashing-spec.md)

[batch-creation-spec](/spec/batch-creation-spec.md)

[valset-creation-spec](/spec/valset-creation-spec.md)

### Design docs

These are mid-level docs which go into the most detail on various topics relating to the bridge.

[design overview](/docs/design/overview.md)

[Bootstrapping the bridge](/docs/design/bootstrapping.md)

[Minting and locking tokens in Gravity](/docs/design/mint-lock.md)

[Oracle design](/docs/design/oracle.md)

[Ethereum signing](/docs/design/ethereum-signing.md)

[Messages](/docs/design/messages.md)

[Parameters](/docs/design/parameters.md)

[Incentives](/docs/design/incentives.md)

[arbitrary logic](/docs/design/arbitrary-logic.md)

[relaying semantics](/docs/design/relaying-semantics.md)

### Developer Guide

To contribute to Gravity, refer to these guides.

[Development environment setup](/docs/developer/environment-setup.md)

[Code structure](/docs/developer/code-structure.md)

[Adding integration tests](/docs/developer/modifying-integration-tests.md)

[Security hotspots](/docs/developer/hotspots.md)

## Status

Gravity bridge is under development and will be undergoing audits soon. Instructions for deployment and use are provided in the hope that they will be useful.

It is your responsibility to understand the financial, legal, and other risks of using this software. There is no guarantee of functionality or safety. You use Gravity bridge entirely at your own risk.

You can keep up with the latest development by watching our [public standups](https://www.youtube.com/playlist?list=PL1MwlVJloJeyeE23-UmXeIx2NSxs_CV4b) feel free to join yourself and ask questions.

- Solidity Contract
  - [x] Multiple ERC20 support
  - [x] Tested with 100+ validators
  - [x] Unit tests for every throw condition
  - [x] Audit for assets originating on Ethereum
  - [x] Support for issuing Cosmos assets on Ethereum
- Cosmos Module
  - [x] Basic validator set syncing
  - [x] Basic transaction batch generation
  - [x] Ethereum -> Cosmos Token issuing
  - [x] Cosmos -> Ethereum Token issuing
  - [x] Bootstrapping
  - [x] Genesis file save/load
  - [x] Validator set syncing edge cases
  - [x] Slashing
  - [x] Relaying edge cases
  - [x] Transaction batch edge cases
  - [x] Support for issuing Cosmos assets on Ethereum
  - [ ] Audit
- Orchestrator / Relayer
  - [x] Validator set update relaying
  - [x] Ethereum -> Cosmos Oracle
  - [x] Transaction batch relaying
  - [ ] Tendermint KMS support
  - [ ] Audit

## The design of Gravity Bridge

- Trust in the integrity of the Gravity bridge is anchored on the Cosmos side. The signing of fraudulent validator set updates and transaction batches meant for the Ethereum contract is punished by slashing on the Cosmos chain. If you trust the Cosmos chain, you can trust the Gravity bridge operated by it, as long as it is operated within certain parameters.
- It is mandatory for peg zone validators to maintain a trusted Ethereum node. This removes all trust and game theory implications that usually arise from independent relayers, once again dramatically simplifying the design.

## Key design Components

- A highly efficient way of mirroring Cosmos validator voting onto Ethereum. The Gravity solidity contract has validator set updates costing ~500,000 gas ($2 @ 20gwei), tested on a snapshot of the Cosmos Hub validator set with 125 validators. Verifying the votes of the validator set is the most expensive on chain operation Gravity has to perform. Our highly optimized Solidity code provides enormous cost savings. Existing bridges incur more than double the gas costs for signature sets as small as 8 signers.
- Transactions from Cosmos to ethereum are batched, batches have a base cost of ~500,000 gas ($2 @ 20gwei). Batches may contain arbitrary numbers of transactions within the limits of ERC20 sends per block, allowing for costs to be heavily amortized on high volume bridges.

## Operational parameters ensuring security

- There must be a validator set update made on the Ethereum contract by calling the `updateValset` method at least once every Cosmos unbonding period (usually 2 weeks). This is because if there has not been an update for longer than the unbonding period, the validator set stored by the Ethereum contract could contain validators who cannot be slashed for misbehavior.
- Cosmos full nodes do not verify events coming from Ethereum. These events are accepted into the Cosmos state based purely on the signatures of the current validator set. It is possible for the validators with >2/3 of the stake to put events into the Cosmos state which never happened on Ethereum. In this case observers of both chains will need to "raise the alarm". We have built this functionality into the relayer.

## Cudos changes to original Althea GravityBridge
Since we forked the project, there are several changes that we've made to the repo. They are described below

### Cosmos sdk, tendermint, osmosis-lab ibc version updates
We've replaced the cosmos-sdk repo to our own fork of it.

We've updated the tendermint version because of the above.

We've replaced the osmosis-lab ibc module to cosmos-sdk's one. This was also needed because of the cosmos-sdk update.

### Removed hardcoded bridge fee while using gbt
We've removed the ahrdcoded bridge fee that was set in the msgs, sent from gbt. We needed it removed for test purposes and since it is not used in production we left it at that.

### Replaced all ETH address checks to lowercase
We replaced all ETH addresses checks and tests to lowr cases for consistency sake, since the addresses we receive from events are not consistent.

### Export all params to genesis
Some params were missing in genesis, so we added them in.

### Fix param in genesis order
We've added a key sorting priod to getting the values, so we can ensure that we always get them in the same order. This fixes a non-determinism issue that might occur.

### Removed market feature

### Removed slashing
Removed the slashing feature and and the tests for it

### Handle non-running eth/cosmos nodes
Prior to our change, when a probe to RPC connection to a node failed, the orchestrator qould panic. We've changed it so it rather throws an error and retries,

### Remove logic calls
We've removed all LogicCall functions from the contract, the module and the orchestrator and the tests for them.

### Added check for 0th address in MsgSendToEth

### Added env variables to orchestrator

### Initialize chain with 0th Gravity contract address vy default

### Added access control
We've added an access control to the Gravity contract, which limits who can call some functions.

### Fixed some typos and outputs

### Added minimum amount to send to ETH
We've implemented a minimum amount of acudos to send to ETH as to prevent spammint with 1 acudo transactions. This includes new parameted in Gravity module, new checks to the messages, new message for setting the minimum amount by admin and some unit tests.

### Added automatic fee calculation in orchestrator

### Added whitelist functionality on some functions

### Added minimum bridge fee for MsgSendToEth
Since the orchestrators sign ethereum transactions in order to validate the transfers from Cudos to Ethereum, they need to receive some minimum amount of CUDOS in order to not be at a loss at the end. We've implemented a minimum amount of bridge fee that needs to be set in each transfer.

To do this we've implemented a new parameter, messages for setting it, that can be ran only by adminToken holders and checks in MsgSendToEth. Also we've redone some other tests to accomodate that change.

### Fixed and improved MsgCancelSendToEth
The CancelSendToEth functionality we had from the fork had a bug. We've fixed it and improved it by adding queries for transfers that are not yet included in a batch and can be canceled. This is done for ease of UI use.

Also more tests added.

### Added Gravity contract verifivation on etherscan

### Updated ECRECOVER function
There are a few potential problems with the standard ecrecover function. That is why we implemented a zero address check after the ecrecover and also decided to use OpenZeppelins' tryEcrecover function. The latter required us to update our solidity version to ^0.8.0. From the update a few changes to the imports and a little change to the CosmosToken were required, but nothing major.

### Added pause functionality on Gravity contract

### Only admin to some functions

### Added check for empty bytecode address on sendToCosmos
This makes sure that the ERC20 contract address leads to a deployed contract, because there is a case where SendToCosmos could be called with not yet deployed contract.

### Added list of supported tokens
We've added a list with supported tokens in which the functions check if the given ERC20 address is valid.

### Added gas optimizations

### Only the highest power orchestrator sends submitBatch transactions
This is a change in the orchestrator only. We check if the current orchestartor is the one with the highest power, and only if it is, we let it send the submit batch transaction. 

This is done so not all orchestrators waste gas on resubmiting batches.

### Tests reformatting for orchestrator
Prerequisites: validator with highest power sends batch, retry getting gravity id, automatic fee calculations, 

Ater all the changes, some test reforms had to be made in order to incorporate everything.

### Removed test uniswap luquidity

### Added static valset functionality
We've added a static list of validators that participate in the orchestrating process. This is set during the init of the chain and can only be changed with a fork.

This is done on the module level, where if a validator tries to set itself as orchestrator, and is not in the static valset list, an error is thrown.