# HB Rollup/Validium

> **Note:** This document is a work in progress and may be subject to change.

Our optimisic rollup/validium implementation is designed to allow a good ratio of low operational fees (meaning low L2 fees) and high security. Each rollup block is a bundle of 1000s of L2 blocks that are posted to our data avaliability layer, and commited to on L1. Celestia blobsteamX is acts as the glue allowing us to securely load
data posted off chain into the L1 contract when needed. 

# Lifecycle of the Rollup Chain

1. The sequencer creates and executes a new L2 Block by aggregating transactions from the mempool and self-queued transactions from L1. 
2. The publisher periodically bundles these L2 Blocks to create the Rollup Block body and posts the bundle to Celestia.
3. The publisher posts the Rollup Header to the CanonicalStateChain contract on L1.
4. Honest actors have a window of time to challenge the block. 
5. If no challenges are successful, the block is considered finalized and the chain is updated.

## Parameters

| Constant           | Value |
| ------------------ | ----- |
| `CHALLENGE_WINDOW` | `3 days` |
| `CHALLENGE_PERIOD` | `2 days` |
| `CHAIN_PUBLISHER`  | `TBA` |
| `MAX_BUNDLE_SIZE`  | `3000` |
| `CHALLENGE_FEE`    | `0.5 eth` |
| `DEFEND_FEE`       | `0.3 eth` |

## Rollup Header 

Rollup headers are posted to the `CanonicalStateChain` contract on L1. Each header serves as a summary of a the blocks body – a bundle of L2 Blocks. It includes claims about the height of the L2 chain, the expected state after correct execution, and pointers to the block body that was published on Celestia. Clients can use the headers to infer the state of the L2 chain, and to challenge incorrect claims.

- [Lifecycle of a Header](#lifecycle-of-a-header)
- [Header Publisher](#header-publisher)
- [Header Structure](#header-structure)
- [Hashing headers](#hashing-headers)
- [Header Validation](#header-validation)
- [Header challenges](#header-challenges)

### Lifecycle of a Header
Once a header is posted to Layer 1 it is pending until `CHALLENGE_WINDOW` has passed. During this period,
anybody can challenge the correctness of the header. If no challenges succeed during the `CHALLENGE_WINDOW`, the header is considered finalized – This header can never revert and the chain up to it is immutable.  Otherwise, if a challenge succeeds, and the header does make incorrect claims, the chain is rolled back to the last correct header.

### Header Publisher
Only `CHAIN_PUBLISHER` can post headers to Layer 1. If the `CHAIN_PUBLISHER` acts badly, as demonstrated by multiple successful challenges, it will be slashed and eventually replaced. The publisher will run the [hummingbird-client](https://github.com/lightlink-network/hummingbird-client), which contains runtimes for automattically posting and defending.

### Header Structure 

Rollup Headers contain the following fields:

| Field                | Type    | Description |
|----------------------|---------|-------------|
| **`epoch`**          | `uint64`| The epoch is the ethereum block number that the header is synced to. It's used to ensure that all self queued transactions, posted prior to this epoch, have been included in rollup block. |
| **`l2Height`**       | `uint64`| The height of the Layer 2 chain. The layer 2 block range included in rollup block N, start at `header(N-1).l2Height` and end at `header(N).l2Height`. |
| **`prevHash`**       | `bytes32`| The hash of the previous rollup block header. |
| ~~**`txRoot`**~~         | `bytes32`| The Merkle root of a tree containing all transactions included in this block. (**_deprecated_**). |
| ~~**`blockRoot`**~~      | `bytes32`| The Merkle root of all a tree containing all blocks in the bundle. (**_deprecated_**). |
| **`stateRoot`**      | `bytes32`| The root of the state trie after executing all layer 2 blocks in the bundle. |
| **`celestiaHeight`** | `bytes32`| The height of the block on Celestia where the bundle was posted. |
| **`celestiaShareStart`** | `bytes32`| The start index of the share of the bundle in the Celestia block. |
| **`celestiaShareLen`** | `bytes32`| The length of the share of the bundle in the Celestia block. |
### Hashing headers

The hash of a rollup block header is obtained by applying the keccak256 hash function to the RLP-encoded header fields, sequenced as follows: [`epoch`, `l2Height`, `prevHash`, `txRoot`, `blockRoot`, `stateRoot`, `celestiaHeight`, `celestiaShareStart`, `celestiaShareLen`].

```solidity
struct Header {
    uint256 epoch;
    uint256 l2Height;
    bytes32 prevHash;
    bytes32 txRoot;
    bytes32 blockRoot;
    bytes32 stateRoot;
    uint256 celestiaHeight;
    uint256 celestiaShareStart;
    uint256 celestiaShareLen;
}

function hashHeader(Header memory _header) public view returns (bytes32) {
    return keccak256(abi.encode(_header));
}
```

### Header Validation

Basic validation of a header is performed when the header is submitted to the `CanonicalStateChain` contract. The following checks are performed:

```
header(N).epoch > header(N-1).epoch && header(N).l2Height > header(N-1).l2Height && header(N).prevHash == hashHeader(header(N-1))
```

- The `epoch` is greater than the previous header's `epoch`.
- The `l2Height` is greater than the previous header's `l2Height`.
- The `prevHash` matches the hash of the previous header in the chain.

The remaining fields are "optimistically" assumed to be correct, but they can be challenged by anyone during the `CHALLENGE_WINDOW`. Most challenges follow a similar flow, where once initiated the defender has until the `CHALLENGE_PERIOD` to respond to the challenge. If the defender fails to respond, the challenger wins and the header is reverted. If the defender responds with a valid proof, the challenger is slashed and the header is considered finalized.


#### Header challenges

| Name | Purpose | Summary | Fee | Reward | Steps |
| ---- | ------- | ------- | --- | ------ | ----- |
| **ChallengeHeader**    | Ensures most of the Rollup block headers fields are valid.                | A single shot challenge that verifies the `epoch` is greater than the previous header's epoch; The `l2Height` is greater than or equal to the previous header's `l2Height`; The `prevHash` matches the hash of the previous header; Finally that the bundle size is less than or equal to `MAX_BUNDLE_SIZE`. (Bundle size is calculated as `bundleSize = header(N).l2Height - header(N-1).l2Height`).                                                                                                                                                                                                                                                                             | 0.0             | 0.0                          | 1     |
| **ChallengeL2Header**  | Ensures that L2 headers are available and most fields are valid.          | Launches a challenge game which requires a defender to load a given L2 block at height `X` on to L1 where `header(N-1).l2Height < X <= header(N)`. The L2 blocks are loaded via the ChainOracle which requires celestia shares within the `celestiaShareStart` and `celestiaShareLen` range to be decoded into a block, as well as a valid data attestations. If the defend fails to load a header with a correct timestamp, and parent hash, the challenger wins.                                                                                                                                                                                                                | `CHALLENGE_FEE` | `CHALLENGE_FEE` | 2-3   |
| **ChallengeL2Tx**      | Ensure that L2 transactions are available and most fields are valid.      | Launches a challenge game targeting a specific L2 header at height `X` on to L1 where `header(N-1).l2Height < X <= header(N)`.The challenger must preload the L2 Header into the ChainOracle. The defender must first provide the transactions hashes `TXS` as (leafs) of MerkleTree with a root matching the `TxRoot` field of the target L2 block. On completion, the challenger can target a specific transaction hash at Index `Y` where `Y < len(TXS)`. Finally the defender must load the given transaction into the ChainOracle, where some of transactions fields are also verified. If the defender fails to complete within the `CHALLENGE_PERIOD` the challenger wins. | `CHALLENGE_FEE` | `CHALLENGE_FEE + DEFEND_FEE`              | 2-5   |
| **ChallengeState**     | Ensure that the state root is correct by checking execution using MIPSEVM | **TODO** | _ | _ | _ |
| **ChallengeTxInclusion** | Ensure that self-queued transactions are included in the rollup block.    | **TODO** | _ | _ | _ |


