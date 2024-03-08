import { ethers, network } from "hardhat";
import { verify } from "../../utils/verify";

// Set DAOracle address
const DAOracleAddr = "0x3a5cbB6EF4756DA0b3f6DAE7aB6430fD8c46d247";

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
  console.log("DAOracle address set to:", DAOracleAddr + "\n");

  // Add new provider for pegasus rpc
  const pegasus = new ethers.JsonRpcProvider(process.env.PEGASUS_PROVIDER_URL);

  // Call pegasus rpc to get the latest blocks state root
  const latestBlock = await pegasus.provider.send("eth_getBlockByNumber", [
    "latest",
    true,
  ]);
  console.log(
    "Latest L2 block number for L1 genesis:",
    parseInt(latestBlock?.number, 16),
  );
  console.log("Latest L2 block hash for L1 genesis:", latestBlock?.hash);
  console.log(
    "Latest L2 block state root for L1 genesis:",
    latestBlock?.stateRoot + "\n",
  );

  // Build genesis header from latest L2 block
  const genesisHeader = {
    epoch: 0,
    l2Height: parseInt(latestBlock?.number, 16),
    prevHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    blockRoot: latestBlock?.hash,
    stateRoot: latestBlock?.stateRoot, // fix state root
    celestiaHeight: 0,
    celestiaShareStart: 0,
    celestiaShareLen: 0,
  };

  ///
  /// Deploy contracts
  ///

  const proxyFactory: any = await ethers.getContractFactory("CoreProxy");

  // Deploy CanonicalStateChain contract as proxy
  console.log("Deploying CanonicalStateChain...");
  const canonicalStateChainFactory: any = await ethers.getContractFactory(
    "CanonicalStateChain",
  );
  const canonicalStateChainImplementation =
    await canonicalStateChainFactory.deploy();
  await canonicalStateChainImplementation.waitForDeployment();
  const canonicalStateChainImplementationAddr =
    await canonicalStateChainImplementation.getAddress();

  const canonicalStateChainProxy = await proxyFactory.deploy(
    canonicalStateChainImplementationAddr,
    canonicalStateChainImplementation.interface.encodeFunctionData(
      "initialize",
      [publisherAddr, genesisHeader],
    ),
  );
  await canonicalStateChainProxy.waitForDeployment();
  const canonicalStateChain = canonicalStateChainFactory.attach(
    await canonicalStateChainProxy.getAddress(),
  );
  const canonicalStateChainContractAddr =
    await canonicalStateChain.getAddress();
  console.log(
    `→ CanonicalStateChain proxy deployed to ${canonicalStateChainContractAddr}`,
  );
  console.log(
    `→ CanonicalStateChain implementation deployed to ${canonicalStateChainImplementationAddr}`,
  );

  // Deploy Treasury contract
  console.log("Deploying Treasury...");
  const Treasury = await ethers.getContractFactory("Treasury");
  const treasury = await Treasury.deploy();
  await treasury.waitForDeployment();
  const treasuryAddr = await treasury.getAddress();
  console.log(`→ Treasury deployed to ${treasuryAddr}`);

  // Deploy required RLPReader lib
  console.log("Deploying RLPReader...");
  const RLPReader = await ethers.getContractFactory("RLPReader");
  const rlpReader = await RLPReader.deploy();
  await rlpReader.waitForDeployment();
  const rlpReaderAddr = await rlpReader.getAddress();
  console.log(`→ RLPReader deployed to ${rlpReaderAddr}`);

  // Deploying ChainOracle contract as a proxy
  console.log("Deploying ChainOracle...");
  const chainOracleFactory = await ethers.getContractFactory("ChainOracle");
  const chainOracleImplementation = await chainOracleFactory.deploy();

  await chainOracleImplementation.waitForDeployment();
  const chainOracleImplementationAddr =
    await chainOracleImplementation.getAddress();
  const chainOracleProxy = await proxyFactory.deploy(
    chainOracleImplementationAddr,
    chainOracleImplementation.interface.encodeFunctionData("initialize", [
      canonicalStateChainContractAddr,
      DAOracleAddr,
      rlpReaderAddr,
    ]),
  );
  await chainOracleProxy.waitForDeployment();
  const chainOracle = chainOracleFactory.attach(
    await chainOracleProxy.getAddress(),
  );
  const chainOracleContractAddr = await chainOracle.getAddress();
  console.log(`→ ChainOracle proxy deployed to ${chainOracleContractAddr}`);
  console.log(
    `→ ChainOracle implementation deployed to ${chainOracleImplementationAddr}`,
  );

  // Deploy Challenge contract as a proxy
  console.log("Deploying Challenge...");
  const challengeFactory: any = await ethers.getContractFactory("Challenge");
  const challengeImplementation = await challengeFactory.deploy();
  await challengeImplementation.waitForDeployment();
  const challengeImplementationAddr =
    await challengeImplementation.getAddress();

  const challengeProxy = await proxyFactory.deploy(
    challengeImplementationAddr,
    challengeImplementation.interface.encodeFunctionData("initialize", [
      treasuryAddr,
      canonicalStateChainContractAddr,
      DAOracleAddr,
      ethers.ZeroAddress,
      chainOracleContractAddr,
    ]),
  );
  await challengeProxy.waitForDeployment();
  const challenge = challengeFactory.attach(await challengeProxy.getAddress());
  const challengeContractAddr = await challenge.getAddress();
  console.log(`→ Challenge proxy deployed to ${challengeContractAddr}`);
  console.log(
    `→ Challenge implementation deployed to ${challengeImplementationAddr}`,
  );

  ///
  /// Set contract setters
  ///

  // set CanonicalStateChain.challengeContract() to challengeContractAddr
  await canonicalStateChain.setChallengeContract(challengeContractAddr);
  console.log(
    `→ → CanonicalStateChain.challengeContract() set to ${challengeContractAddr}` +
      "\n",
  );

  // set Challenge.setDefender() to publisherAddr
  await challenge.setDefender(publisherAddr);
  console.log(`→ → Challenge.setDefender() set to ${publisherAddr}`);

  console.log("All Contracts deployed successfully! \n");

  ///
  /// Verify contracts
  ///

  if (chainId !== 31337 && chainId !== 1337) {
    // Verify contract (after 1 min)
    console.log("Waiting for 1 min before verifying contracts..");
    await new Promise((r) => setTimeout(r, 120000));

    // Verify CanonicalStateChain Implementation
    await verify(
      canonicalStateChainImplementationAddr,
      [],
      "contracts/CanonicalStateChain.sol:CanonicalStateChain",
    );
    console.log(
      `Verified CanonicalStateChain impl contract at ${canonicalStateChainImplementationAddr}`,
    );

    // Verify CanonicalStateChain Proxy
    await verify(
      canonicalStateChainContractAddr,
      [
        canonicalStateChainImplementationAddr,
        canonicalStateChainImplementation.interface.encodeFunctionData(
          "initialize",
          [publisherAddr, genesisHeader],
        ),
      ],
      "contracts/proxy/CoreProxy.sol:CoreProxy",
    );

    // Verify Treasury
    await verify(treasuryAddr, [], "contracts/Treasury.sol:Treasury");

    // Verify RLPReader
    await verify(rlpReaderAddr, [], "contracts/RLPReader.sol:RLPReader");

    // Verify ChainOracle Implementation
    await verify(
      chainOracleImplementationAddr,
      [],
      "contracts/ChainOracle.sol:ChainOracle",
    );
    console.log(
      `Verified ChainOracle impl contract at ${chainOracleImplementationAddr}`,
    );

    // Verify ChainOracle Proxy
    await verify(
      chainOracleContractAddr,
      [
        chainOracleImplementationAddr,
        chainOracleImplementation.interface.encodeFunctionData("initialize", [
          canonicalStateChainContractAddr,
          DAOracleAddr,
          rlpReaderAddr,
        ]),
      ],
      "contracts/proxy/CoreProxy.sol:CoreProxy",
    );

    // Verify Challenge Implementation
    await verify(
      challengeImplementationAddr,
      [],
      "contracts/challenge/Challenge.sol:Challenge",
    );
    console.log(
      `Verified Challenge impl contract at ${challengeImplementationAddr}`,
    );

    // Verify Challenge Proxy
    await verify(
      challengeContractAddr,
      [
        challengeImplementationAddr,
        challengeImplementation.interface.encodeFunctionData("initialize", [
          ethers.ZeroAddress,
          canonicalStateChainContractAddr,
          DAOracleAddr,
          ethers.ZeroAddress,
          chainOracleContractAddr,
        ]),
      ],
      "contracts/proxy/CoreProxy.sol:CoreProxy",
    );
  }

  // print contract addresses
  console.log("\n");
  console.log("Contract addresses:");
  console.log(" canonicalStateChain:", canonicalStateChainContractAddr);
  console.log(" chainOracle:", chainOracleContractAddr);
  console.log(" challenge:", challengeContractAddr);
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
