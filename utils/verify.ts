import { run } from "hardhat";

export async function verify(
  contractAddress: string,
  args?: any[],
  contractName?: string,
) {
  console.log("Verifying contract: " + contractAddress);
  try {
    await run("verify:verify", {
      address: contractAddress,
      constructorArguments: args,
      contract: contractName,
    });
  } catch (e: any) {
    if (e.message.toLowerCase().includes("already verified")) {
      console.log("Already verified!");
    } else {
      console.log(e);
    }
  }
  console.log("Verified successfully! \n");
}

module.exports = {
  verify,
};
