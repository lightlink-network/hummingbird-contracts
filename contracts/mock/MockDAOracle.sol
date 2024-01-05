pragma solidity ^0.8.0;

import "blobstream-contracts/src/IDAOracle.sol";

contract MockDAOracle is IDAOracle {
    bool private result = true;

    function setMockResult(bool _result) external {
        result = _result;
    }

    function verifyAttestation(
        uint256 _tupleRootNonce,
        DataRootTuple memory _tuple,
        BinaryMerkleProof memory _proof
    ) external view override returns (bool) {
        return result;
    }
}
