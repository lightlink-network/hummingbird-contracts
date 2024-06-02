import {
  BinaryMerkleProofStruct,
  ChainOracle,
  SharesProofStruct,
} from "../../typechain-types/contracts/ChainOracle";

export const provideHeader = async (
  oracle: ChainOracle,
  rblockHash: string,
  pointerIndex: number,
  proof: SharesProofStruct,
  ranges: ChainOracle.ShareRangeStruct[],
  pointerProof: BinaryMerkleProofStruct[],
) => {
  const shareKey = await oracle.ShareKey(rblockHash, proof.data);
  await oracle.provideShares(rblockHash, pointerIndex, proof);
  await oracle.provideHeader(shareKey, ranges);
};

export const provideLegacyTx = async (
  oracle: ChainOracle,
  rblockHash: string,
  pointerIndex: number,
  proof: SharesProofStruct,
  ranges: ChainOracle.ShareRangeStruct[],
  pointerProof: BinaryMerkleProofStruct[],
) => {
  const shareKey = await oracle.ShareKey(rblockHash, proof.data);
  await oracle.provideShares(rblockHash, pointerIndex, proof);
  await oracle.provideLegacyTx(shareKey, ranges);
};
