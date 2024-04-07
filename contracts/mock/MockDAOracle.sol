// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "blobstream-contracts/src/IDAOracle.sol";

/// @title MockDAOracle
/// @author LightLink Hummingbird
/// @custom:version v1.0.0-alpha
/// @notice A mock implementation of the DAOracle contract.
contract MockDAOracle is IDAOracle {
    /// @notice The result of the verification.
    bool public _result = true;

    /// @notice Verifies the attestation.
    /// @param _tupleRootNonce - The nonce of the tuple root.
    /// @param _tuple - The data root tuple.
    /// @param _proof - The binary merkle proof.
    function verifyAttestation(
        uint256 _tupleRootNonce,
        DataRootTuple memory _tuple,
        BinaryMerkleProof memory _proof
    ) external view returns (bool) {
        _tupleRootNonce;
        _tuple;
        _proof;
        return _result;
    }

    /// @notice Sets the result of the verification.
    /// @param result - The result of the verification.
    function setResult(bool result) external {
        _result = result;
    }
}
