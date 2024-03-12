import { task } from "hardhat/config";

task("rollupHead", "Get the head of the chain")
  .addPositionalParam("csc", "Address of the CanonicalStateChain contract")
  .setAction(async (args, hre) => {
    const csc = await hre.ethers.getContractAt("CanonicalStateChain", args.csc);
    const head = await csc.getHead();
    const headHash = await csc.chain(await csc.chainHead());
    console.log("The chain head is", headHash, head);
  });
