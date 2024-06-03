import { config as cfg } from "dotenv";
import { HardhatUserConfig, task } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-verify";
import "@solarity/hardhat-gobind";
import "./tasks/rollupHead";
import "./tasks/challengeL2Header";
import "./tasks/pushRBlock";
import "solidity-coverage";

cfg();

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.8.22",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
          viaIR: true,
        },
      },
      {
        version: "0.7.6",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ],
  },
  defaultNetwork: "hardhat",
  networks: {
    hardhat: {},
    localhost: {
      chainId: 1337,
    },
    sepolia: {
      url: process.env.SEPOLIA_PROVIDER_URL || "",
      accounts: [
        process.env.SEPOLIA_OWNER_PRIVATE_KEY ??
          "0000000000000000000000000000000000000000000000000000000000000000",
        process.env.SEPOLIA_PUBLISHER_PRIVATE_KEY ??
          "0000000000000000000000000000000000000000000000000000000000000000",
      ],
    },
    ethereum: {
      url: process.env.ETHEREUM_PROVIDER_URL || "",
      accounts: [
        process.env.ETHEREUM_OWNER_PRIVATE_KEY ??
          "0000000000000000000000000000000000000000000000000000000000000000",
        process.env.ETHEREUM_PUBLISHER_PRIVATE_KEY ??
          "0000000000000000000000000000000000000000000000000000000000000000",
      ],
    },
    pegasus: {
      url: process.env.PEGASUS_PROVIDER_URL || "",
    },
    arbSepolia: {
      url: process.env.ARBITRUM_SEPOLIA_PROVIDER_URL || "",
      accounts: [
        process.env.ARBITRUM_SEPOLIA_OWNER_PRIVATE_KEY ??
          "0000000000000000000000000000000000000000000000000000000000000000",
        process.env.ARBITRUM_SEPOLIA_PUBLISHER_PRIVATE_KEY ??
          "0000000000000000000000000000000000000000000000000000000000000000",
      ],
    },
  },
  mocha: {
    timeout: 200000,
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: "USD",
    coinmarketcap: process.env.COINMARKETCAP_API_KEY,
    L1: "ethereum",
    L1Etherscan: process.env.ETHERSCAN_API_KEY,
    darkMode: true,
  },
  etherscan: {
    apiKey: {
      hardhat: "123456",
      ethereum: process.env.ETHERSCAN_API_KEY || "",
      sepolia: process.env.ETHERSCAN_API_KEY || "",
      arbSepolia: process.env.ETHERSCAN_API_KEY || "",
    },
    customChains: [
      {
        network: "arbSepolia",
        chainId: 421614,
        urls: {
          apiURL: "https://api-sepolia.arbiscan.io/api",
          browserURL: "https://sepolia.arbiscan.io",
        },
      },
    ],
  },
  gobind: {
    outdir: "./generated-types/bindings",
    deployable: false,
    runOnCompile: false,
    verbose: false,
    onlyFiles: [
      "contracts/CanonicalStateChain.sol",
      "contracts/challenge/Challenge.sol",
      "contracts/ChainOracle.sol",
      "contracts/interfaces/IBlobstreamX.sol",
    ],
    skipFiles: ["@openzeppelin", "@solarity"],
  },
};

export default config;
