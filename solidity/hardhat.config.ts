import "@nomiclabs/hardhat-waffle";
import "hardhat-gas-reporter";
import "hardhat-typechain";
import { task } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import * as dotenv from "dotenv";

const lazyImport = async (module: any) => {
  return await import(module);
};

task("accounts", "Prints the list of accounts", async (args, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

task("verify-contracts", "Verifies contracts").setAction(async () => {
  const { verifyContracts } = await lazyImport("./scripts/verify");
  await verifyContracts();
});

// // This is a sample Buidler task. To learn how to create your own go to
// // https://buidler.dev/guides/create-task.html
// task("accounts", "Prints the list of accounts", async (taskArgs, bre) => {
//   const accounts = await bre.ethers.getSigners();

//   for (const account of accounts) {
//     console.log(await account.getAddress());
//   }
// });

// You have to export an object to set up your config
// This object can have the following optional entries:
// defaultNetwork, networks, solc, and paths.
// Go to https://buidler.dev/config/ to learn more
module.exports = {
  // This is a sample solc configuration that specifies which version of solc to use
  solidity: {
    compilers: [
    {
    version: "0.6.6",
    settings: {
      optimizer: {
        enabled: true
      }
    }  
  },
  {
    version: "0.6.12",
    settings: {
      optimizer: {
        enabled: true
      }
    }  
  }
]
  },
  defaultNetwork: "rinkeby",
  networks: {
    rinkeby: {
      url: "https://rinkeby.infura.io/v3/62f1fc27624f401283692320734e387e"
    },
    hardhat: {
      timeout: 2000000
    }
  },
  typechain: {
    outDir: "typechain",
    target: "ethers-v5",
    runOnCompile: true
  },
  gasReporter: {
    enabled: true
  },
  mocha: {
    timeout: 2000000
  },
  etherscan: {
    apiKey: "2MD12YN15Z3CKRNJE3RCMAJ1QUN3W1CDXC"
  }
};
