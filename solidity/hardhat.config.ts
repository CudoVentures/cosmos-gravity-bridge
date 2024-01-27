import "@nomiclabs/hardhat-waffle";
import "hardhat-gas-reporter";
import "hardhat-typechain";
import { task } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan"


task("accounts", "Prints the list of accounts", async (args, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
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
  },
  {
    version: "0.8.1",
    settings: {
      optimizer: {
        enabled: true
      }
    }  
  }
]
  },
  defaultNetwork: process.env.DEFAULT_NETWORK,
  networks: {
    rinkeby: {
      url: `${process.env.ETH_NODE}`
    },
    mainnet: {
      url: `${process.env.ETH_NODE}`
    },
    hardhat: {
      timeout: 2000000,
      accounts: {
        mnemonic: "test test test test test test test test test test test junk",
        path: "m/44'/60'/0'/0",
        initialIndex: 0,
        count: 20,
        passphrase: "",
      }
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
  etherscan:{
    apiKey: `${process.env.ETHERSCAN_API_KEY}`
  }
};
