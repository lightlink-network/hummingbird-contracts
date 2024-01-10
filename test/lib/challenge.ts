// struct ChallengeDAProof {
//   uint256 rootNonce;
//   BinaryMerkleProof proof;
// }
//
// struct BinaryMerkleProof {
//   // List of side nodes to verify and calculate tree.
//   bytes32[] sideNodes;
//   // The key of the leaf to verify.
//   uint256 key;
//   // The number of leaves in the tree
//   uint256 numLeaves;
// }

interface ChallengeDAProof {
  rootNonce: bigint;
  proof: BinaryMerkleProof;
}

interface BinaryMerkleProof {
  sideNodes: string[];
  key: bigint;
  numLeaves: bigint;
}
