require("@nomicfoundation/hardhat-toolbox");
require("@nomicfoundation/hardhat-ethers");
require("@openzeppelin/hardhat-upgrades");
require("hardhat-deploy");
require("hardhat-deploy-ethers");
require("dotenv").config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      priceFeeds: {
        ethUsd: "0x694AA1769357215DE4FAC081bf1f309aDC325306",
        usdcUsd: "0x0A6513e40db6EB1b165753AD52E80663aeA50545",
        daiUsd: "0x0d79df66BE487753B02D015Fb622DED7f0E9798d",
        usdtUsd: "0xc16679aad21eb7d1d45ab1fa80c05e81ea1cd02f"
      }
    },
    sepolia: {
      url: "https://sepolia.infura.io/v3/9472bfbb916f4d07af74d5bb77a0bd0a",
      accounts: [process.env.PRIVATE_KEY], // 这里填你的私钥
      namedAccounts: {
        firstAccount: {
          default: "0xF829C797962fE6DA0001550cc897d2Dd3c9883Df"
        }
      },
      
      priceFeeds: {
        ethUsd: "0x694AA1769357215DE4FAC081bf1f309aDC325306",
        usdcUsd: "0x0A6513e40db6EB1b165753AD52E80663aeA50545",
        daiUsd: "0x0d79df66BE487753B02D015Fb622DED7f0E9798d"
      }
    },
    hardhat: {
      priceFeeds: {
        ethUsd: "0x694AA1769357215DE4FAC081bf1f309aDC325306",
        usdcUsd: "0x0A6513e40db6EB1b165753AD52E80663aeA50545",
        daiUsd: "0x0d79df66BE487753B02D015Fb622DED7f0E9798d",
        usdtUsd: "0xc16679aad21eb7d1d45ab1fa80c05e81ea1cd02f"
      }
    }
  }
};


