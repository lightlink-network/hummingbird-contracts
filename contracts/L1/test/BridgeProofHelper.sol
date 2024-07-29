pragma solidity 0.8.22;

import "../LightLinkPortal.sol";

// This contract is used to check the validity of the output and withdrawal proofs.
// This contract is not meant to be deployed in production, it is only used for testing the
// validity of client generated proofs.
contract BridgeProofHelper {
    function checkProof(
        bytes32 outputRoot,
        Types.OutputRootProof calldata outputProof,
        bytes32 withdrawalHash,
        bytes[] calldata withdrawalProof
    ) public pure returns (bool) {
        require(
            Hashing.hashOutputRootProof(outputProof) == outputRoot,
            "INVALID_OUTPUT_ROOT"
        );

        bytes32 storageKey = keccak256(
            abi.encode(
                withdrawalHash,
                uint256(0) // The withdrawals mapping is at the first slot in the layout.
            )
        );

        require(
            SecureMerkleTrie.verifyInclusionProof({
                _key: abi.encode(storageKey),
                _value: hex"01",
                _proof: withdrawalProof,
                _root: outputProof.messagePasserStorageRoot
            }),
            "LightLinkPortal: invalid withdrawal inclusion proof"
        );

        return true;
    }

    function hashWithdrawalTx(
        Types.WithdrawalTransaction memory _tx
    ) public pure returns (bytes32) {
        return Hashing.hashWithdrawal(_tx);
    }
}
