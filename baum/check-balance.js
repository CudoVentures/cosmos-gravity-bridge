const { Web3 } = require("web3");
const endpointUrl = "http://localhost:8545"
const httpProvider = new Web3.providers.HttpProvider(endpointUrl);
const web3Client = new Web3(httpProvider)

const minABI = [
  // balanceOf
  {
    constant: true,
    inputs: [{ name: '_owner', type: 'address' }],
    name: 'balanceOf',
    outputs: [{ name: 'balance', type: 'uint256'}],
    type: 'function',
  },
]

const tokenAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
const walletAddress = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8";

const contract = new web3Client.eth.Contract(minABI, tokenAddress);

async function getBalance() {
  const result = await contract.methods.balanceOf(walletAddress).call();
  console.log(result);
}

getBalance();
