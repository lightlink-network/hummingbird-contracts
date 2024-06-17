import { readFileSync } from "fs";
import { task } from "hardhat/config";

task("pushRBlock", "Push a block to the chain")
  .addPositionalParam("csc", "Address of the CanonicalStateChain contract")
  .addPositionalParam("filepath", "Path to the block file")
  .setAction(async (args, hre) => {
    const [owner, publisher] = await hre.ethers.getSigners();

    // get csc contract
    console.log("Connecting to CanonicalStateChain contract at", args.csc);
    const csc = await hre.ethers.getContractAt("CanonicalStateChain", args.csc);

    // read block from file
    console.log("Reading block from", args.filepath);
    const block = JSON.parse(readFileSync(args.filepath, "utf8"));

    // validate block
    validateBlock(block);
    console.log("Block is valid");

    // override prevHash
    const prevHash = await csc.chain(await csc.chainHead());
    block.prevHash = prevHash;

    // push block to the chain
    const tx = await csc.connect(publisher).pushBlock(block);
    await tx.wait();
    console.log("Block pushed to the chain", tx.hash);
    const head = await csc.getHead();
    const headHash = await csc.chain(await csc.chainHead());

    console.log("The new chain head is", headHash, head);
  });

// // check block has the right format
// e.g.
// epoch, l2Height, prevHash, txRoot, blockRoot, outputRoot, celestiaHeight, celestiaShareStart, celestiaShareLen
const validateBlock = (block: any) => {
  if (!block.epoch) throw new Error("Block is missing epoch");
  if (!block.l2Height) throw new Error("Block is missing l2Height");
  if (!block.prevHash) throw new Error("Block is missing prevHash");
  if (!block.txRoot) throw new Error("Block is missing txRoot");
  if (!block.blockRoot) throw new Error("Block is missing blockRoot");
  if (!block.outputRoot) throw new Error("Block is missing stateRoot");
  if (!block.celestiaHeight) throw new Error("Block is missing celestiaHeight");
  if (!block.celestiaShareStart)
    throw new Error("Block is missing celestiaShareStart");
  if (!block.celestiaShareLen)
    throw new Error("Block is missing celestiaShareLen");
};
