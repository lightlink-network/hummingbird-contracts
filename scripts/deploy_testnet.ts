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
    celestiaDataRoot:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
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

  // Deploying challenge/Challenge.sol contract
  console.log("Deploying Challenge...");
  const challenge = await ethers.getContractFactory("Challenge");
  const challengeContract = await challenge.deploy(
    ethers.ZeroAddress, // treasury
    await canonicalStateChain.getAddress(), // chain
    await mockDAOracle.getAddress(), // daOracle
    ethers.ZeroAddress // mips
  );
  await challengeContract.waitForDeployment();
  console.log(
    `→ Challenge deployed to ${await challengeContract.getAddress()}`
  );

  // setup challenge contract
  await challengeContract.setDefender(publisherAddr);
  console.log(`– Set defender to ${publisherAddr}`);
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
