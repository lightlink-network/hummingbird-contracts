import {
  ChainOracle,
  SharesProofStruct,
} from "../../typechain-types/contracts/ChainOracle";

export const provideHeader = async (
  oracle: ChainOracle,
  rblockHash: string,
  proof: SharesProofStruct,
  ranges: ChainOracle.ShareRangeStruct[],
) => {
  const shareKey = await oracle.ShareKey(rblockHash, proof.data);
  await oracle.provideShares(rblockHash, proof);
  await oracle.provideHeader(shareKey, ranges);
};

export const provideLegacyTx = async (
  oracle: ChainOracle,
  rblockHash: string,
  proof: SharesProofStruct,
  ranges: ChainOracle.ShareRangeStruct[],
) => {
  const shareKey = await oracle.ShareKey(rblockHash, proof.data);
  await oracle.provideShares(rblockHash, proof);
  await oracle.provideLegacyTx(shareKey, ranges);
};
