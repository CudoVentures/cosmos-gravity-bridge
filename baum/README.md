Some scripts I have known.

## Warning

This uses hard-coded private keys derived from the well-known "test test test test test test test test test test test junk" mnemonic.

These keys must be presumed compromised. Do not use them for anything real!

## You will need...

* Compile the Solidity contracts...
  in ../solidity/
  
  $ npm exec hardhat compile

* A copy of `cudos`

  $ npm i -g cudos

* No other Cosmos node running on the typical ports.

## So do run...

* start.sh, in the baum/ directory.

  This will start the chains, deploy the contracts, start the orchestrator, and send a transaction.
