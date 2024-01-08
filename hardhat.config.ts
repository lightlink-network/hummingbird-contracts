import { config as cfg } from 'dotenv'
import { HardhatUserConfig, task } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-verify";

cfg()

const config: HardhatUserConfig = {
  solidity: {
    compilers: [{ version: "0.8.22" }, { version: "0.7.6" }],
  },
  networks: {
    hardhat: {
      forking: {
        url: process.env.SEPOLIA_PROVIDER_URL || '',
      },
    },
    sepolia: {
      url: process.env.SEPOLIA_PROVIDER_URL || '',
      accounts: [process.env.SEPOLIA_PRIVATE_KEY ?? ''],
    },
    ethereum: {
      url: process.env.ETHEREUM_PROVIDER_URL || '',
      accounts: [process.env.ETHEREUM_PRIVATE_KEY ?? ''],
    },
  },
  mocha: {
    timeout: 200000,
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: 'USD',
  },
  etherscan: {
    apiKey: {
      hardhat: '123456',
      mainnet: process.env.ETHERSCAN_API_KEY || '',
      sepolia: process.env.ETHERSCAN_API_KEY || '',
    },
  },
};

export default config;
