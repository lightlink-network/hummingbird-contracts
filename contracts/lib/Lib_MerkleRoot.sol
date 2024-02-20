// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

library Lib_MerkleRoot {
    // Calculate the Merkle root from an array of leaf nodes
    // @dev The number of leaf nodes must be a power of 2
    function calculateMerkleRoot(
        bytes32[] memory leafNodes
    ) public pure returns (bytes32) {
        require(leafNodes.length > 0, "Must have at least one leaf node");
        require(
            // bitwise AND to check if the number of leaf nodes is a power of 2
            // explanation: if the num is a power of 2, it will have only one bit set to 1
            // so, if we subtract 1 from it, all the bits will be set to 1
            // e.g. 8 (1000) - 1 = 7 (0111)
            // so, if we do a bitwise AND, it will return 0 as no bits match
            (leafNodes.length & (leafNodes.length - 1)) == 0,
            "Number of leaf nodes must be a power of 2"
        );

        while (leafNodes.length > 1) {
            bytes32[] memory parentNodes = new bytes32[](leafNodes.length / 2);

            for (uint i = 0; i < leafNodes.length; i += 2) {
                parentNodes[i / 2] = keccak256(
                    abi.encodePacked(leafNodes[i], leafNodes[i + 1])
                );
            }

            leafNodes = parentNodes;
        }

        // The remaining element is the Merkle root
        return leafNodes[0];
    }
}
