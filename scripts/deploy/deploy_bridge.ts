import { ethers, network } from "hardhat";
import { verify } from "../../utils/verify";

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

  const proxyFactory: any = await ethers.getContractFactory("CoreProxy");

  // Deploy OptimismPortal contract as proxy
  console.log("Deploying OptimismPortal...");
  const optimismPortalFactory: any =
    await ethers.getContractFactory("OptimismPortal");
  const optimismPortalImplementation = await optimismPortalFactory.deploy();
  await optimismPortalImplementation.waitForDeployment();
  const optimismPortalImplementationAddr =
    await optimismPortalImplementation.getAddress();

  const optimismPortalProxy = await proxyFactory.deploy(
    optimismPortalImplementationAddr,
    optimismPortalImplementation.interface.encodeFunctionData("initialize", []),
  );
  await optimismPortalProxy.waitForDeployment();
  const optimismPortal = optimismPortalFactory.attach(
    await optimismPortalProxy.getAddress(),
  );
  const optimismPortalContractAddr = await optimismPortal.getAddress();
  console.log(
    `→ OptimismPortal proxy deployed to ${optimismPortalContractAddr}`,
  );
  console.log(
    `→ OptimismPortal implementation deployed to ${optimismPortalImplementationAddr}`,
  );

  // Deploy L1CrossDomainMessenger contract as proxy
  console.log("Deploying L1CrossDomainMessenger...");
  const l1CrossDomainMessengerFactory: any = await ethers.getContractFactory(
    "L1CrossDomainMessenger",
  );
  const l1CrossDomainMessengerImplementation =
    await l1CrossDomainMessengerFactory.deploy();
  await l1CrossDomainMessengerImplementation.waitForDeployment();
  const l1CrossDomainMessengerImplementationAddr =
    await l1CrossDomainMessengerImplementation.getAddress();

  const l1CrossDomainMessengerProxy = await proxyFactory.deploy(
    l1CrossDomainMessengerImplementationAddr,
    l1CrossDomainMessengerImplementation.interface.encodeFunctionData(
      "initialize",
      [optimismPortalContractAddr],
    ),
  );
  await l1CrossDomainMessengerProxy.waitForDeployment();
  const l1CrossDomainMessenger = l1CrossDomainMessengerFactory.attach(
    await l1CrossDomainMessengerProxy.getAddress(),
  );
  const l1CrossDomainMessengerContractAddr =
    await l1CrossDomainMessenger.getAddress();
  console.log(
    `→ L1CrossDomainMessenger proxy deployed to ${l1CrossDomainMessengerContractAddr}`,
  );
  console.log(
    `→ L1CrossDomainMessenger implementation deployed to ${l1CrossDomainMessengerImplementationAddr}`,
  );

  ///
  /// Verify contracts
  ///

  // Verify contract (after 1 min)
  console.log("Waiting for 2 min before verifying contracts..");
  await new Promise((r) => setTimeout(r, 120000));

  // Verify OptimismPortal Implementation
  await verify(
    optimismPortalImplementationAddr,
    [],
    "contracts/L1/OptimismPortal.sol:OptimismPortal",
  );
  console.log(
    `Verified OptimismPortal impl contract at ${optimismPortalImplementationAddr}`,
  );

  // Verify OptimismPortal Proxy
  await verify(
    optimismPortalContractAddr,
    [
      optimismPortalImplementationAddr,
      optimismPortalImplementation.interface.encodeFunctionData(
        "initialize",
        [],
      ),
    ],
    "contracts/universal/CoreProxy.sol:CoreProxy",
  );
  console.log(
    `Verified OptimismPortal proxy contract at ${optimismPortalContractAddr}`,
  );

  // Verify L1CrossDomainMessenger Implementation
  await verify(
    l1CrossDomainMessengerImplementationAddr,
    [],
    "contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger",
  );
  console.log(
    `Verified L1CrossDomainMessenger impl contract at ${l1CrossDomainMessengerImplementationAddr}`,
  );

  // Verify L1CrossDomainMessenger Proxy
  await verify(
    l1CrossDomainMessengerContractAddr,
    [
      l1CrossDomainMessengerImplementationAddr,
      l1CrossDomainMessengerImplementation.interface.encodeFunctionData(
        "initialize",
        [optimismPortalContractAddr],
      ),
    ],
    "contracts/universal/CoreProxy.sol:CoreProxy",
  );
  console.log(
    `Verified L1CrossDomainMessenger proxy contract at ${l1CrossDomainMessengerContractAddr}`,
  );

  console.log("Deployment completed!");
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
