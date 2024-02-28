import { ethers, network } from "hardhat";
import { verify } from "../../utils/verify";

// Set current addresses
const canonicalStateChainContractProxyAddr = "CONTRACT_ADDRESS";

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

  // Deploy new CanonicalStateChain contract implementation
  console.log("Deploying CanonicalStateChain...");
  const canonicalStateChainFactory: any = await ethers.getContractFactory("CanonicalStateChain");
  const canonicalStateChainImplementation = await canonicalStateChainFactory.deploy();
  await canonicalStateChainImplementation.waitForDeployment();
  const canonicalStateChainImplementationAddr =
    await canonicalStateChainImplementation.getAddress();

  console.log(
    `→ CanonicalStateChain implementation deployed to ${canonicalStateChainImplementationAddr}`
  );

  // Update CanonicalStateChain contract proxy to new implementation
  console.log("Updating CanonicalStateChain proxy to new implementation...");
  const canonicalStateChainProxy = await ethers.getContractAt(
    "CanonicalStateChain",
    canonicalStateChainContractProxyAddr
  );
  await canonicalStateChainProxy.upgradeToAndCall(canonicalStateChainImplementationAddr, "0x");
  console.log(
    `→ CanonicalStateChain proxy (${canonicalStateChainContractProxyAddr}) updated to ${canonicalStateChainImplementationAddr}` + "\n"
  );

  console.log("All Contracts deployed successfully! \n");

  ///
  /// Verify contracts
  ///

  // Verify contract (after 1 min)
  console.log("Waiting for 1 min before verifying contracts..");
  await new Promise((r) => setTimeout(r, 60000));

  // Verify CanonicalStateChain Implementation
  await verify(
    canonicalStateChainImplementationAddr,
    [],
    "contracts/CanonicalStateChain.sol:CanonicalStateChain"
  );
  console.log(
    `Verified CanonicalStateChain impl contract at ${canonicalStateChainImplementationAddr}`
  );
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
