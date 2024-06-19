import { ethers, network } from "hardhat";
import { verify } from "../../utils/verify";
import { log } from "../lib/log";

const main = async () => {
    // Log network name and chain id selected for deployment
    const chainIdHex = await network.provider.send("eth_chainId");
    const chainId = parseInt(chainIdHex, 16);
    log("Network name:", network.name);
    log("Network chain id:", chainId + "\n");

    // Step 1. Get deployer/signer account
    const [owner, publisher] = await ethers.getSigners();
    const [ownerAddr] = [
        await owner.getAddress(),
    ];
    log("Owner address:", ownerAddr);

    // Step 2. Deploy L2ToL1MessagePasser contract
    const L1ToL2MessagePasser = await ethers.getContractFactory("L2ToL1MessagePasser");
    const l1ToL2MessagePasser = await L1ToL2MessagePasser.deploy();

    console.log("L1ToL2MessagePasser:", await l1ToL2MessagePasser.getAddress());
};

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });