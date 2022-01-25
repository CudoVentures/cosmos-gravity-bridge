//@ts-nocheck
import chai from "chai";
import { ethers } from "hardhat";
import { solidity } from "ethereum-waffle";

import { BridgeAccessControl } from "../typechain/BridgeAccessControl";
import { deployContracts } from "../test-utils";
import {
  getSignerAddresses,
  makeCheckpoint,
  signHash,
  makeTxBatchHash,
  examplePowers
} from "../test-utils/pure";

chai.use(solidity);
const { expect } = chai;

let bridgeAccessControl:any

async function runTest(opts: {}) {


  const BridgeAccessControl = await ethers.getContractFactory("BridgeAccessControl");
  bridgeAccessControl = (await BridgeAccessControl.deploy());
  // Prep and deploy contract
  // ========================
  const signers = await ethers.getSigners();
  const gravityId = ethers.utils.formatBytes32String("foo");
  // This is the power distribution on the Cosmos hub as of 7/14/2020
  let powers = examplePowers();
  let validators = signers.slice(0, powers.length);
  const powerThreshold = 6666;
  const {
    gravity,
    testERC20,
    checkpoint: deployCheckpoint
  } = await deployContracts(gravityId, powerThreshold, validators, powers, bridgeAccessControl.address);


  // Transfer out to Cosmos, locking coins
  // =====================================
  await testERC20.functions.approve(gravity.address, 1000);
  await expect(gravity.functions.sendToCosmos(
    testERC20.address,
    ethers.utils.formatBytes32String("myCosmosAddress"),
    1000
  )).to.emit(gravity, 'SendToCosmosEvent').withArgs(
      testERC20.address,
      await signers[0].getAddress(),
      ethers.utils.formatBytes32String("myCosmosAddress"),
      1000, 
      2
    );

  expect((await testERC20.functions.balanceOf(gravity.address))[0]).to.equal(1000);
  expect((await gravity.functions.state_lastEventNonce())[0]).to.equal(2);


    
  // Do it again
  // =====================================
  await testERC20.functions.approve(gravity.address, 1000);
  await expect(gravity.functions.sendToCosmos(
    testERC20.address,
    ethers.utils.formatBytes32String("myCosmosAddress"),
    1000
  )).to.emit(gravity, 'SendToCosmosEvent').withArgs(
      testERC20.address,
      await signers[0].getAddress(),
      ethers.utils.formatBytes32String("myCosmosAddress"),
      1000, 
      3
    );

  expect((await testERC20.functions.balanceOf(gravity.address))[0]).to.equal(2000);
  expect((await gravity.functions.state_lastEventNonce())[0]).to.equal(3);
}

describe("sendToCosmos tests", function () {
  it("works right", async function () {
    await runTest({})
  });
});