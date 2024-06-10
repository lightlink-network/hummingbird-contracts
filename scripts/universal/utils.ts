import { ethers } from "hardhat";

/**
 * Deploy a Proxy and contract and initialize it.
 * @param signer The signer to deploy the contract.
 * @param implementationFactory The factory of the implementation contract.
 * @param args The arguments to initialize on the contract.
 * @returns The proxy and implementation contract.
 */
export const proxyDeployAndInitialize = async (
  signer: any,
  implementationFactory: any,
  args: any[],
) => {
  // step 1: deploy implementation contract
  const implementation = await implementationFactory.connect(signer).deploy();
  await implementation.waitForDeployment();
  const implementationAddress = await implementation.getAddress();

  // step 2: deploy proxy contract
  const proxyFactory: any = await ethers.getContractFactory("CoreProxy");
  const contract = await proxyFactory
    .connect(signer)
    .deploy(
      implementationAddress,
      implementation.interface.encodeFunctionData("initialize", args),
    );
  await contract.waitForDeployment();
  const proxyAddress = await contract.getAddress();

  return {
    contract: await implementationFactory.attach(proxyAddress),
    address: proxyAddress,
    implementation,
    implementationAddress,
  };
};
