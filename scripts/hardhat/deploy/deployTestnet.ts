import { ethers } from "hardhat";
import { setupCanonicalStateChain } from "../../../test/lib/chain";
import {
  Challenge,
  LightLinkPortal,
  L2CrossDomainMessenger,
  L1CrossDomainMessenger,
} from "../../../typechain-types";
import {
  proxyDeployAndInitialize,
  uupsProxyDeployAndInitialize,
} from "../lib/deploy";
import { startNetworks } from "../lib/startNetworks";

const main = async () => {
  // Start Anvil network instances
  const networks = await startNetworks();
  const l1Network = networks.l1Network;
  const l2Network = networks.l2Network;

  console.log("Anvil networks started");

  // Set up L1 provider and signer
  const l1Provider = new ethers.JsonRpcProvider("http://0.0.0.0:8545");
  const l1Deployer = (await l1Provider.getSigner(0)) as any;

  // Set up L2 provider and signer
  const l2Provider = new ethers.JsonRpcProvider("http://0.0.0.0:8546");
  const l2Deployer = (await l2Provider.getSigner(0)) as any;

  // Deploy L1 contracts

  // - CanonicalStateChain
  const _chain = await setupCanonicalStateChain(l1Deployer, l1Deployer.address);
  const canonicalStateChain = _chain.canonicalStateChain;
  console.log("CanonicalStateChain deployed");

  // - Challenge
  const challengeDeployment = await uupsProxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("Challenge"),
    [
      await canonicalStateChain.getAddress(),
      ethers.ZeroAddress,
      ethers.ZeroAddress,
    ],
  );
  const challenge = challengeDeployment.contract as Challenge;

  // - LightLinkPortal
  const lightLinkPortalDeployment = await proxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("LightLinkPortal"),
    [
      await canonicalStateChain.getAddress(),
      await challengeDeployment.address,
      ethers.ZeroAddress, // L1Block address
    ],
  );
  const lightLinkPortal = lightLinkPortalDeployment.contract as LightLinkPortal;

  // - BridgeProofHelper
  const bridgeProofHelperFactory = await ethers.getContractFactory(
    "contracts/L1/test/BridgeProofHelper.sol:BridgeProofHelper",
    l1Deployer,
  );
  const BridgeProofHelper = (await bridgeProofHelperFactory.deploy()) as any;
  await BridgeProofHelper.waitForDeployment();

  // Deploy L2 contracts

  // - L2ToL1MessagePasser
  const L2ToL1MessagePasserFactory = await ethers.getContractFactory(
    "contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser",
    l2Deployer,
  );
  const l2ToL1MessagePasser =
    (await L2ToL1MessagePasserFactory.deploy()) as any;
  await l2ToL1MessagePasser.waitForDeployment();

  // - L1Block
  const L1BlockFactory = await ethers.getContractFactory(
    "contracts/L2/L1Block.sol:L1Block",
    l2Deployer,
  );
  const l1Block = (await L1BlockFactory.deploy()) as any;

  // Cross domain messengers
  // - Infer deployment addresses before deploying
  const l2CrossDomainMessengerAddr = ethers.getCreateAddress({
    from: l2Deployer.address,
    nonce: (await l2Provider.getTransactionCount(l2Deployer.address)) + 1,
    // +1 because implementation will be deployed first
  });
  const l1CrossDomainMessengerAddr = ethers.getCreateAddress({
    from: l1Deployer.address,
    nonce: (await l1Provider.getTransactionCount(l1Deployer.address)) + 1,
    // +1 because implementation will be deployed first
  });

  // - Deploy cross domain messengers
  console.log("Deploying cross domain messengers");
  const L2CrossDomainMessengerDeployment = await proxyDeployAndInitialize(
    l2Deployer,
    await ethers.getContractFactory("L2CrossDomainMessenger"),
    [
      l1CrossDomainMessengerAddr,
      await l2ToL1MessagePasser.getAddress(),
      await l1Block.getAddress(),
    ],
  );
  const l2CrossDomainMessenger =
    L2CrossDomainMessengerDeployment.contract as L2CrossDomainMessenger;

  const L1CrossDomainMessengerDeployment = await proxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("L1CrossDomainMessenger"),
    [await lightLinkPortal.getAddress(), l2CrossDomainMessengerAddr],
  );
  const l1CrossDomainMessenger =
    L1CrossDomainMessengerDeployment.contract as L1CrossDomainMessenger;

  // Impersonate l2 Depositor account
  console.log("Impersonating L2 depositor account");
  await l2Provider.send("hardhat_impersonateAccount", [
    "0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001",
  ]);

  const l2Depositor = (await l2Provider.getSigner(
    "0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001",
  )) as any;

  console.log("L2 depositor account impersonated - funding...");
  await l2Deployer.sendTransaction({
    to: l2Depositor.address,
    value: ethers.parseEther("1"),
  });

  // Setup GasPayingToken in L1Block
  console.log("Setting up GasPayingToken in L1Block");
  await l1Block.connect(l2Depositor).setGasPayingToken(
    "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", // Constants.ETHER
    18,
    ethers.encodeBytes32String("Ether"),
    ethers.encodeBytes32String("ETH"),
  );

  // setup mining interval
  await l1Provider.send("evm_setIntervalMining", [5_000]);
  await l2Provider.send("evm_setIntervalMining", [2_500]);

  // done

  console.log("\n\n");
  console.log(
    " canonicalStateChain:",
    `"${await canonicalStateChain.getAddress()}"`,
  );
  console.log(" challenge:", `"${await challenge.getAddress()}"`);
  console.log(" chainOracle:", `"${ethers.ZeroAddress}"`);
  console.log(" blobstreamX:", `"${ethers.ZeroAddress}"`);
  console.log(" rlpReader:", `"${ethers.ZeroAddress}" \n`);

  console.log("\n\n");
  console.log(
    " l2CrossDomainMessenger:",
    `"${await l2CrossDomainMessenger.getAddress()}"`,
  );
  console.log(
    " l1CrossDomainMessenger:",
    `"${await l1CrossDomainMessenger.getAddress()}"`,
  );
  console.log(
    " l2ToL1MessagePasser:",
    `"${await l2ToL1MessagePasser.getAddress()}"`,
  );
  console.log(" lightLinkPortal:", `"${await lightLinkPortal.getAddress()}"`);
  console.log(" l1Block:", `"${await l1Block.getAddress()}"`);

  console.log("\n\n");
  console.log(" ... running, press Ctrl+C to stop");
  while (true) {
    await new Promise((r) => setTimeout(r, 1000));
  }
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
