import { ethers, network } from "hardhat";
import { verify } from "../../utils/verify";
import { log } from "../lib/log";

// These contracts should eventually be deployed as predeployed contracts
// and not as part of the deployment script.
// 
// But for testnet deployment, we can deploy them this way.
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

    // Step 3. Deploy L1Block contract
    const L1Block = await ethers.getContractFactory("L1Block");
    const l1Block = await L1Block.deploy();

    console.log("L1ToL2MessagePasser:", await l1ToL2MessagePasser.getAddress());
    console.log("L1Block:", await l1Block.getAddress());
};

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });