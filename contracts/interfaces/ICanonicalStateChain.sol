// SPDX-License-Identifier: UNLICENSED
// LightLink Hummingbird v0.0.1

pragma solidity ^0.8.0;

interface ICanonicalStateChain {
    struct Header {
        bytes32 prevHash; // PrevHash is the hash of the previous block bundle.
        uint32 epoch; // Epoch refers to a block number on the Ethereum blockchain.
        uint32 l2Height; // L2Height is the index of the Last L2 Block in this bundle.
        // Pointers to the blocks contents on celestia.
        // See `Span` from https://docs.celestia.org/developers/blobstream-offchain#defining-a-chain
        CelestiaPointer[] pointers;
    }

    struct CelestiaPointer {
        uint32 celestiaHeight;
        uint32 celestiaShareStart;
        uint32 celestiaShareLen;
    }

    struct HeaderMetadata {
        address publisher;
        uint32 timestamp;
    }

    event BlockAdded(uint256 indexed blockNumber);
    event Rolledback(uint256 indexed blockNumber);
    event PublisherChanged(address indexed publisher);

    // publisher is the verified address that can add new blocks to the chain.
    // This address can be replaced by the owner of the contract, (expected to be
    // the rollup contract).
    function publisher() external view returns (address);

    // Challenge is the address of the challenge contract. This contract can
    // rollback the chain after a successful challenge is made.
    function challenge() external view returns (address);

    // Rollup Blockchain.
    function chainHead() external view returns (uint256);

    function headers(bytes32) external view returns (Header memory);

    function getBlockByHash(
        bytes32 _hash
    ) external view returns (Header memory);

    function headerMetadata(
        bytes32
    ) external view returns (HeaderMetadata memory);

    function chain(uint256) external view returns (bytes32);

    // pushBlock optimistically pushes block headers to the canonical chain.
    // The only fields that are checked are the epoch and prevHash.
    function pushBlock(Header calldata _header) external;

    function hash(Header memory _header) external pure returns (bytes32);

    function getBlock(uint256 _index) external view returns (Header memory);

    function getHead() external view returns (Header memory);

    // Rollback reverts the chain to a previous state, It can only be called by the challenge
    // contract.
    function rollback(uint256 _blockNumber) external;

    function setPublisher(address _publisher) external;
}
