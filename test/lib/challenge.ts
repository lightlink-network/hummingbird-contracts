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
  dataRootTuple: DataRootTuple;
  proof: BinaryMerkleProof;
}

// struct DataRootTuple {
//   // Celestia block height the data root was included in.
//   // Genesis block is height = 0.
//   // First queryable block is height = 1.
//   uint256 height;
//   // Data root.
//   bytes32 dataRoot;
// }

interface DataRootTuple {
  height: bigint;
  dataRoot: string;
}

interface BinaryMerkleProof {
  sideNodes: string[];
  key: bigint;
  numLeaves: bigint;
}
