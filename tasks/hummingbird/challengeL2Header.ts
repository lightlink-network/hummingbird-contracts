import { task } from "hardhat/config";

task("challengeL2Header", "Challenge L2 header")
  .addPositionalParam("challenge", "Address of the Challenge contract")
  .addPositionalParam("rblock", "RBlock to challenge")
  .addPositionalParam("header", "Header to challenge")
  .setAction(async (args, hre) => {
    const [owner] = await hre.ethers.getSigners();

    const challenge = await hre.ethers.getContractAt(
      "Challenge",
      args.challenge,
    );
    const csc = await hre.ethers.getContractAt(
      "CanonicalStateChain",
      await challenge.chain(),
    );
    const rblockHash = await csc.chain(args.rblock);
    const challengeFee = await challenge.challengeFee();

    const challengeKey = await challenge.l2HeaderChallengeHash(
      rblockHash,
      args.header,
    );
    console.log("Challenge key", challengeKey);

    const tx = await challenge
      .connect(owner)
      .challengeL2Header(args.rblock, args.header, { value: challengeFee });
    console.log(
      "Challenged header",
      args.header,
      "in rblock",
      args.rblock,
      rblockHash,
      "with key",
      challengeKey,
      "in tx",
      tx.hash,
    );

    await tx.wait();
  });
