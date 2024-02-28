import { ethers, network } from "hardhat";
import { verify } from "../../utils/verify";

// Set current addresses
const challengeContractProxyAddr = "CONTRACT_ADDRESS";

const main = async () => {
  // Log network name and chain id selected for deployment
  const chainIdHex = await network.provider.send("eth_chainId");
  const chainId = parseInt(chainIdHex, 16);
  console.log("Network name:", network.name);
  console.log("Network chain id:", chainId + "\n");

  // Get deployer/signer account
  const [owner, publisher] = await ethers.getSigners();
  const ownerAddr = await owner.getAddress();
  const publisherAddr = await publisher.getAddress();

  console.log("Owner address is set to:", ownerAddr);
  console.log("Publisher address is set to:", publisherAddr);

  ///
  /// Deploy contracts
  ///

  // Deploy new Challenge contract implementation
  console.log("Deploying Challenge...");
  const challengeFactory: any = await ethers.getContractFactory("Challenge");
  const challengeImplementation = await challengeFactory.deploy();
  await challengeImplementation.waitForDeployment();
  const challengeImplementationAddr =
    await challengeImplementation.getAddress();

  console.log(
    `→ Challenge implementation deployed to ${challengeImplementationAddr}`
  );

  // Update Challenge contract proxy to new implementation
  console.log("Updating Challenge proxy to new implementation...");
  const challengeProxy = await ethers.getContractAt(
    "Challenge",
    challengeContractProxyAddr
  );
  await challengeProxy.upgradeToAndCall(challengeImplementationAddr, "0x");
  console.log(
    `→ Challenge proxy (${challengeContractProxyAddr}) updated to ${challengeImplementationAddr}` + "\n"
  );

  console.log("All Contracts deployed successfully! \n");

  ///
  /// Verify contracts
  ///

  // Verify contract (after 1 min)
  console.log("Waiting for 1 min before verifying contracts..");
  await new Promise((r) => setTimeout(r, 60000));

  // Verify Challenge Implementation
  await verify(
    challengeImplementationAddr,
    [],
    "contracts/challenge/Challenge.sol:Challenge"
  );
  console.log(
    `Verified Challenge impl contract at ${challengeImplementationAddr}`
  );
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
