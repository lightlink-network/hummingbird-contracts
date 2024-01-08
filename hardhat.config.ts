import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: {
    compilers: [{ version: "0.8.22" }, { version: "0.7.6" }],
  },
  networks: {
    ganache: {
      url: "http://138.68.156.191:8545",
      accounts: [
        "0xd5bc419beb60ea28eaaa67fb2f0809a43c22ee985522b651141d285fcb289d60",
        "0xec9a1d0c5cdc025927957c97112bfdc08de59fcd994d558495653056bc874de4",
        "0x0b6941926cea7aaed5576651c1c17a0ae21d50895a6b8e4926908e98cae3415b",
        "0x1960459a507dbe386171a0d18f9d81ea6505453b843c72cc74677a8b5292f8c7",
        "0x5c94a55a3cff94de626280003484df6440fdd56e2adbe583b43dc17af10a1354",
        "0xcae6a10074eba3ac5d07b5b7067afdcfb68335f98b5d55ebaccc17a76df3844c",
        "0xb1ff96b3ce8232e4f4da2f6c762cfeab1d5421f9232204cbc4f2e109f6c6f0b3",
        "0x8a685fbcee5e24390991ca357be749e0c1f45e83c08a0cfbd2cec3e0f5c46e25",
        "0x022b8376bf30f84786f30ac94fccea5d85e7acdc4fa2259fa70382ff92061f4e",
        "0x9f60347bef0b7c8f2f1f431bd3c2d379745769637ac88a881e44b18dc24f88e5",
      ],
    },
  },
};

export default config;
