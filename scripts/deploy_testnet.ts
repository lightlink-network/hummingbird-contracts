import { ethers } from "hardhat";

const main = async () => {
  // Get deployer/signer account
  const [deployer, publisher] = await ethers.getSigners();
  const deployerAddr = await deployer.getAddress();
  const publisherAddr = await deployer.getAddress();
  console.log("Deploying contracts with the deployer address:", deployerAddr);

  // 1. CanonicalStateChain
  // Build genesis header block 62207259
  const genesisHeader = {
    epoch: 0,
    l2Height: 62207259,
    prevHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    blockRoot:
      "0x3810c405e3789ebee1ffb7ab8d36debdcefb629f129bb9db00d3271148b55f83",
    stateRoot:
      "0x491cc7d79299f9569e4bfddef640ade68091ab486d68a08ed1477c678db34103",
    celestiaHeight: 0,
    celestiaShareStart: 0,
    celestiaShareLen: 0,
  };

  // Deploy CanonicalStateChain contract
  console.log("Deploying CanonicalStateChain...");
  const CanonicalStateChain = await ethers.getContractFactory(
    "CanonicalStateChain"
  );
  const canonicalStateChain = await CanonicalStateChain.deploy(
    publisherAddr,
    genesisHeader
  );
  await canonicalStateChain.waitForDeployment();

  console.log(
    `→ CanonicalStateChain deployed to ${await canonicalStateChain.getAddress()}`
  );

  // 2. MockDAOracle
  // Deploying MockDAOracle contract
  console.log("Deploying MockDAOracle...");
  const MockDAOracle = await ethers.getContractFactory("MockDAOracle");
  const mockDAOracle = await MockDAOracle.deploy();
  await mockDAOracle.waitForDeployment();
  console.log(`→ MockDAOracle deployed to ${await mockDAOracle.getAddress()}`);

  // Deploy Challenge contract as a proxy
  console.log("Deploying Challenge...");
  const proxyFactory: any = await ethers.getContractFactory("CoreProxy");
  const challengeFactory: any = await ethers.getContractFactory("Challenge");
  const challengeImplementation = await challengeFactory.deploy();
  await challengeImplementation.waitForDeployment();
  const challengeImplementationAddr =
    await challengeImplementation.getAddress();

  const proxy = await proxyFactory.deploy(
    challengeImplementationAddr,
    challengeImplementation.interface.encodeFunctionData("initialize", [
      ethers.ZeroAddress,
      await canonicalStateChain.getAddress(),
      await mockDAOracle.getAddress(),
      ethers.ZeroAddress,
    ])
  );
  await proxy.waitForDeployment();
  const challenge = challengeFactory.attach(await proxy.getAddress());
  const challengeContractAddr = await challenge.getAddress();
  console.log(`→ Challenge proxy deployed to ${challengeContractAddr}`);
  console.log(
    `→ Challenge implementation deployed to ${challengeImplementationAddr}`
  );

  ///
  /// Set contract setters
  ///

  // set Challenge.setDefender() to publisherAddr
  await challenge.setDefender(publisherAddr);
  console.log(`→ → Challenge.setDefender() set to ${publisherAddr}`);

  // set CanonicalStateChain.challengeContract() to challengeContractAddr
  await canonicalStateChain.setChallengeContract(challengeContractAddr);
  console.log(
    `→ → CanonicalStateChain.challengeContract() set to ${challengeContractAddr}` +
      "\n"
  );

  console.log("All Contracts deployed successfully! \n");

  // setup challenge contract
  await challenge.setDefender(publisherAddr);
  console.log(`– Set defender to ${publisherAddr}`);
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
