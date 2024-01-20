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
      accounts: [process.env.SEPOLIA_OWNER_PRIVATE_KEY ?? '', process.env.SEPOLIA_PUBLISHER_PRIVATE_KEY ?? '0x000'],
    },
    ethereum: {
      url: process.env.ETHEREUM_PROVIDER_URL || '',
      accounts: [process.env.ETHEREUM_PRIVATE_KEY ?? ''],
    },
    pegasus: {
      url: process.env.PEGASUS_PROVIDER_URL || '',
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
