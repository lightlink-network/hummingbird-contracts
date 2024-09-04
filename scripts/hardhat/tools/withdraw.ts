import { ethers } from "hardhat";
import {
  L2ToL1MessagePasser,
  LightLinkPortal,
  CanonicalStateChain,
} from "../../../typechain-types";
import {
  getWithdrawalProofs,
  parseMessagePassedEvent,
  hashWithdrawalTx,
  hashMessageHash,
} from "../../../test/lib/bridge";

const withdrawalTxHash =
  "0x6b2a64008dcc75eec7f9a1e0cb764f9556f2fd5f855f838fb57c5cd0c53c9ae4";

const lightLinkPortalAddress = "0x9aBbc181B8b6F6591e70BF785bBb6d999C314925";
const l2ToL1MessagePasserAddress = "0xd8D3a8C83a598550D1C2Cf40932EAE9C9C217E21";
const canonicalStateChainAddress = "0x91C0b1164aB51c3310A7B0ceAEb208016671B1b9";

const l1Provider = new ethers.JsonRpcProvider(process.env.L1_RPC_URL);
const l2Provider = new ethers.JsonRpcProvider(process.env.L2_RPC_URL);

const l1wallet = new ethers.Wallet(process.env.L1_DEPLOYER_KEY!, l1Provider);
const l2wallet = new ethers.Wallet(process.env.L2_DEPLOYER_KEY!, l2Provider);

const main = async () => {
  // attach LightLinkPortal contract
  const LightLinkPortalFactory =
    await ethers.getContractFactory("LightLinkPortal");
  const lightLinkPortal = LightLinkPortalFactory.connect(l1wallet).attach(
    lightLinkPortalAddress,
  ) as LightLinkPortal;

  // attach L2toL1MessagePasser contract
  const L2ToL1MessagePasserFactory = await ethers.getContractFactory(
    "L2ToL1MessagePasser",
  );
  const l2ToL1MessagePasser = L2ToL1MessagePasserFactory.connect(
    l2wallet,
  ).attach(l2ToL1MessagePasserAddress) as L2ToL1MessagePasser;

  // get tx receipt
  const withdrawalReceipt =
    await l2Provider.getTransactionReceipt(withdrawalTxHash);

  // get csc header that includes withdrawal
  let header = await findHeaderForWithdrawal(withdrawalReceipt?.blockNumber!);

  // parse withdrawal tx from receipt event logs
  const { withdrawalTx } = parseMessagePassedEvent(
    l2ToL1MessagePasser.interface,
    withdrawalReceipt!.logs[2],
  );

  // calc withdrawal hash and message slot
  const withdrawalHash = hashWithdrawalTx(withdrawalTx);
  const messageSlot = hashMessageHash(withdrawalHash);

  // generate withdrawal proofs
  const { withdrawalProof, outputProof, outputRoot } =
    await getWithdrawalProofs(
      l2Provider,
      "0x" + header.header.l2Height.toString(16), // last block in header
      l2ToL1MessagePasser,
      messageSlot,
    );

  // send withdrawal proof to L1
  const proveTx = await lightLinkPortal
    .connect(l1wallet)
    .proveWithdrawalTransaction(
      withdrawalTx,
      header.headerNumber,
      outputProof,
      withdrawalProof.storageProof,
    );

  console.log("Withdrawal Proven 🎉 Tx Hash:", proveTx.hash);
};

async function findHeaderForWithdrawal(blockNumber: number): Promise<{
  header: CanonicalStateChain.HeaderStructOutput;
  headerNumber: bigint;
}> {
  try {
    if (blockNumber <= 0) {
      throw new Error(
        "Invalid block number. Block number should be greater than 0.",
      );
    }

    // attach CanonicalStateChain contract
    const canonicalStateChainFactory = await ethers.getContractFactory(
      "CanonicalStateChain",
    );
    const canonicalStateChain = canonicalStateChainFactory
      .connect(l1wallet)
      .attach(canonicalStateChainAddress) as CanonicalStateChain;

    const chainHead = await canonicalStateChain.chainHead();
    let currentHeaderNum = chainHead;
    let header = await canonicalStateChain.getHeaderByNum(currentHeaderNum);

    console.log("Finding header for withdrawal in L1 CSC...");
    console.log("Target block number:", blockNumber);

    if (!header || header.l2Height < blockNumber) {
      console.error("Withdrawal not included in L1 CSC yet");
      process.exit(1);
    }

    while (header && header.l2Height >= blockNumber) {
      console.log(
        `Checking header number: ${currentHeaderNum} with L2 height: ${header.l2Height}`,
      );

      const previousHeaderNum = currentHeaderNum - BigInt(1);
      const previousHeader =
        await canonicalStateChain.getHeaderByNum(previousHeaderNum);

      if (!previousHeader || blockNumber > previousHeader.l2Height) {
        console.log(
          `Withdrawal included in L1 CSC within the range of header number: ${currentHeaderNum}`,
        );
        return { header, headerNumber: currentHeaderNum };
      }

      currentHeaderNum = previousHeaderNum;
      header = previousHeader;
    }

    console.error("Withdrawal not included in L1 CSC yet");
    process.exit(1);
  } catch (error) {
    console.error("An error occurred:", error);
    process.exit(1);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
