// SPDX-License-Identifier: MIT
// LightLink Hummingbird v0.2.0

pragma solidity ^0.8.0;

import "blobstream-contracts/src/IDAOracle.sol";

contract MockDAOracle is IDAOracle {
    bool public _result = true;

    function verifyAttestation(
        uint256 _tupleRootNonce,
        DataRootTuple memory _tuple,
        BinaryMerkleProof memory _proof
    ) external view returns (bool) {
        _tupleRootNonce; // To silence the compiler warning
        _tuple; // To silence the compiler warning
        _proof; // To silence the compiler warning
        return _result;
    }

    function setResult(bool result) external {
        _result = result;
    }
}
