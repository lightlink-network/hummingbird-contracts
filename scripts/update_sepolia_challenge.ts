import { ethers, network } from "hardhat";
import { verify } from "../utils/verify";
import { Contract } from "ethers";

// Set current addresses
const canonicalStateChainAddr = "0x2A66396cd5D8a29B1b7442596060FB34E4d14989";
const challengeContractProxyAddr = "0x6C757cE003377975DE0B06247c257e1Fdef2158C";
const challengeImplAddr = "0x79ce810967eef9c43ac21b7431fdbd10c134cee5";
const treasuryAddr = "0x3a5cbB6EF4756DA0b3f6DAE7aB6430fD8c46d247";
const DAOracleAddr = "0xec30E779dcB18A4d2C6039a84b060F66d689fb23";

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

  const data = challengeImplementation.interface.encodeFunctionData(
    "initialize",
    [treasuryAddr, canonicalStateChainAddr, DAOracleAddr, ethers.ZeroAddress],
    ownerAddr
  );

  console.log("Updating Challenge proxy to new implementation...");
  console.log("→ → Calling Challenge.upgradeToAndCall() with data:", data);

  const challengeProxy = await ethers.getContractAt(
    "Challenge",
    challengeContractProxyAddr
  );
  await challengeProxy.upgradeToAndCall(challengeImplementationAddr, data);
  console.log(
    `→ Challenge proxy updated to ${challengeImplementationAddr}` + "\n"
  );

  ///
  /// Set contract setters
  ///

  // set Challenge.setDefender() to publisherAddr
  await challengeProxy.setDefender(publisherAddr);
  console.log(`→ → Challenge.setDefender() set to ${publisherAddr}`);

  // get currently deployed CanonicalStateChain contract from address
  const canonicalStateChain = await ethers.getContractAt(
    "CanonicalStateChain",
    canonicalStateChainAddr
  );

  // set CanonicalStateChain.challengeContract() to challengeContractAddr
  await canonicalStateChain.setChallengeContract(challengeImplementationAddr);
  console.log(
    `→ → CanonicalStateChain.challengeContract() set to ${challengeImplementationAddr}` +
      "\n"
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
