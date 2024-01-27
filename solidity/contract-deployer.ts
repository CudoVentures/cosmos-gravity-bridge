import { Gravity } from "./typechain/Gravity";
import { TestERC20A } from "./typechain/TestERC20A";
import { TestERC20B } from "./typechain/TestERC20B";
import { TestERC20C } from "./typechain/TestERC20C";
// import { TestUniswapLiquidity } from "./typechain/TestUniswapLiquidity";
import { ethers } from "ethers";
import fs from "fs";
import commandLineArgs from "command-line-args";
import axios from "axios";
import { exit } from "process";
import hre from "hardhat";

const args = commandLineArgs([
  // the ethernum node used to deploy the contract
  { name: "eth-node", type: String },
  // the cosmos node that will be used to grab the validator set via RPC (TODO),
  { name: "cosmos-node", type: String },
  // the Ethereum private key that will contain the gas required to pay for the contact deployment
  { name: "eth-privkey", type: String },
  // the gravity contract .json file
  { name: "contract", type: String },
  // test mode, if enabled this script deploys three ERC20 contracts for testing
  { name: "test-mode", type: String },
  // the address of the cudos access control smart contract
  { name: "cudos-access-control", type: String },
  // the address of the cudos token
  { name: "cudos-token-address", type: String },
]);

// 4. Now, the deployer script hits a full node api, gets the Eth signatures of the valset from the latest block, and deploys the Ethereum contract.
//     - We will consider the scenario that many deployers deploy many valid gravity eth contracts.
// 5. The deployer submits the address of the gravity contract that it deployed to Ethereum.
//     - The gravity module checks the Ethereum chain for each submitted address, and makes sure that the gravity contract at that address is using the correct source code, and has the correct validator set.
type Validator = {
  power: number;
  ethereum_address: string;
};
type ValsetTypeWrapper = {
  type: string;
  value: Valset;
};
type Valset = {
  members: Validator[];
  nonce: number;
};
type ABCIWrapper = {
  jsonrpc: string;
  id: string;
  result: ABCIResponse;
};
type ABCIResponse = {
  response: ABCIResult;
};
type ABCIResult = {
  code: number;
  log: string;
  info: string;
  index: string;
  value: string;
  height: string;
  codespace: string;
};
type StatusWrapper = {
  jsonrpc: string;
  id: string;
  result: NodeStatus;
};
type NodeInfo = {
  protocol_version: JSON;
  id: string;
  listen_addr: string;
  network: string;
  version: string;
  channels: string;
  moniker: string;
  other: JSON;
};
type SyncInfo = {
  latest_block_hash: string;
  latest_app_hash: string;
  latest_block_height: number;
  latest_block_time: string;
  earliest_block_hash: string;
  earliest_app_hash: string;
  earliest_block_height: number;
  earliest_block_time: string;
  catching_up: boolean;
};
type NodeStatus = {
  node_info: NodeInfo;
  sync_info: SyncInfo;
  validator_info: JSON;
};

// sets the gas price for all contract deployments
const overrides = {
  //gasPrice: 100000000000
};

async function deploy() {
  const provider = new ethers.providers.JsonRpcProvider(args["eth-node"]);
  const wallet = new ethers.Wallet(args["eth-privkey"], provider);

  console.error("deploying")
  {
    console.error("deploying erc20")
    const erc20_a_path = "TestERC20A.json";
    const { abi, bytecode } = getContractArtifacts(erc20_a_path);
    const erc20Factory = new ethers.ContractFactory(abi, bytecode, wallet);
    const testERC20 = (await erc20Factory.deploy(overrides)) as TestERC20A;
    await testERC20.deployed();
    args["cudos-token-address"] = testERC20.address;
  }

  console.error("deploying bridge")
  const gravityIdString = await getGravityId();
  const gravityId = ethers.utils.formatBytes32String(gravityIdString);

  const cudosAccessControl = args["cudos-access-control"];

  const { abi, bytecode } = getContractArtifacts(args.contract);
  const factory = new ethers.ContractFactory(abi, bytecode, wallet);

  const latestValset = await getLatestValset();

  const eth_addresses = [];
  const powers = [];
  let powers_sum = 0;
  // this MUST be sorted uniformly across all components of Gravity in this
  // case we perform the sorting in module/x/gravity/keeper/types.go to the
  // output of the endpoint should always be sorted correctly. If you're
  // having strange problems with updating the validator set you should go
  // look there.
  for (let i = 0; i < latestValset.members.length; i++) {
    if (latestValset.members[i].ethereum_address == null) {
      continue;
    }
    eth_addresses.push(latestValset.members[i].ethereum_address);
    powers.push(latestValset.members[i].power);
    powers_sum += latestValset.members[i].power;
  }

  // 66% of uint32_max
  const vote_power = 2834678415;
  if (powers_sum < vote_power) {
    console.error(
      "Refusing to deploy! Incorrect power! Please inspect the validator set below"
    );
    console.error(
      "If less than 66% of the current voting power has unset Ethereum Addresses we refuse to deploy"
    );
    console.error(latestValset);
    exit(1);
  }

  const cudosToken = args["cudos-token-address"];
  const gravity = (await factory.deploy(
    gravityId,
    vote_power,
    eth_addresses,
    powers,
    cudosAccessControl,
    cudosToken
  )) as Gravity;

  await gravity.deployed();

  console.log(`TOKEN=${cudosToken.toLowerCase()}`)
  console.log(`BRIDGE=${gravity.address.toLowerCase()}`)

  await gravity.deployTransaction.wait(1);

}

function getContractArtifacts(path: string): { bytecode: string; abi: string } {
  const { bytecode, abi } = JSON.parse(fs.readFileSync(path, "utf8").toString());
  return { bytecode, abi };
}
const decode = (str: string): string =>
  Buffer.from(str, "base64").toString("binary");

async function getLatestValset(): Promise<Valset> {
  const block_height_request_string = `${args["cosmos-node"]}/status`;
  const block_height_response = await axios.get(block_height_request_string);
  const info: StatusWrapper = await block_height_response.data;
  const block_height = info.result.sync_info.latest_block_height;
  if (info.result.sync_info.catching_up) {
    console.error(
      "This node is still syncing! You can not deploy using this validator set!"
    );
    exit(1);
  }
  const request_string = `${args["cosmos-node"]}/abci_query`;
  const response = await axios.get(request_string, {
    params: {
      path: '"/custom/gravity/currentValset/"',
      height: block_height,
      prove: "false",
    },
  });
  const valsets: ABCIWrapper = await response.data;
  const valset: ValsetTypeWrapper = JSON.parse(
    decode(valsets.result.response.value)
  );
  return valset.value;
}

async function getGravityId(): Promise<string> {
  const block_height_request_string = `${args["cosmos-node"]}/status`;
  const block_height_response = await axios.get(block_height_request_string);
  const info: StatusWrapper = await block_height_response.data;
  const block_height = info.result.sync_info.latest_block_height;
  if (info.result.sync_info.catching_up) {
    console.error(
      "This node is still syncing! You can not deploy using this gravityID!"
    );
    exit(1);
  }
  const request_string = `${args["cosmos-node"]}/abci_query`;
  const response = await axios.get(request_string, {
    params: {
      path: '"/custom/gravity/gravityID/"',
      height: block_height,
      prove: "false",
    },
  });
  const gravityIDABCIResponse: ABCIWrapper = await response.data;
  const gravityID: string = JSON.parse(
    decode(gravityIDABCIResponse.result.response.value)
  );
  return gravityID;
}

async function main() {
  await deploy();
}

main();
