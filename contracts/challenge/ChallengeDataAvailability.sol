// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ChallengeBase.sol";
import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";
import "hardhat/console.sol";

/// @title  ChallengeDataAvailability
/// @author LightLink Hummingbird
/// @custom:version v1.0.0-alpha
/// @notice ChallengeDataAvailability is a challenge for verifying a rollup blocks
///         data root has been included. (via Celestia Blobstream).
///
///         This is a challenge game between two parties: the challenger and the defender.
///         There can only be one challenge per rblock hash.
///
///         The Challenge goes through the following steps:
///         1. A challenger initiates a challenge by calling challengeDataRootInclusion.
///         2. The defending block publisher must provide a proof of inclusion for the
///            data root. If the proof is valid, the defender wins the challenge and
///            receives the challenge fee.
///         3. Otherwise the challenge expires and the challenger wins the challenge and
///            the block is rolled back.
///
///         You can trigger a challenge easily via the hummingbird client:
///         `hb challenge-da <block-index>`.
abstract contract ChallengeDataAvailability is ChallengeBase {
    /// @notice The different states a DA challenge can be in.
    /// @param None - The DA challenge has not been initiated.
    /// @param ChallengerInitiated - The DA challenge has been initiated by the challenger.
    /// @param ChallengerWon - The DA challenge has been won by the challenger.
    /// @param DefenderWon - The DA challenge has been won by the defender.
    enum ChallengeDAStatus {
        None,
        ChallengerInitiated,
        ChallengerWon,
        DefenderWon
    }

    /// @notice The data structure for a DA challenge.
    /// @param blockHash - The block hash of the block being challenged.
    /// @param blockIndex - The index of the block being challenged.
    /// @param pointerIndex - The index of the celestia pointer being challenged.
    /// @param challenger - The address of the challenger.
    /// @param expiry - The expiry time of the challenge.
    /// @param status - The status of the challenge.
    struct ChallengeDA {
        bytes32 blockHash;
        uint256 blockIndex;
        uint8 pointerIndex;
        uint32 shareIndex;
        address challenger;
        uint256 expiry;
        ChallengeDAStatus status;
        bool claimed;
    }

    /// @notice The data structure for a DA challenge proof.
    /// @param rootNonce - The nonce of the data root.
    /// @param dataRootTuple - The data root tuple.
    /// @param proof - The binary merkle proof.
    struct ChallengeDAProof {
        uint256 rootNonce;
        DataRootTuple dataRootTuple;
        BinaryMerkleProof proof;
    }

    /// @notice Emitted when a DA challenge is updated.
    /// @param _blockHash - The block hash of the block being challenged.
    /// @param _pointerIndex - The index of the celestia pointer being challenged.
    /// @param _blockIndex - The index of the block being challenged.
    /// @param _expiry - The expiry time of the challenge.
    /// @param _status - The status of the challenge.
    event ChallengeDAUpdate(
        bytes32 indexed _blockHash,
        uint256 indexed _pointerIndex,
        uint32 _shareIndex,
        uint256 _blockIndex,
        uint256 _expiry,
        ChallengeDAStatus indexed _status
    );

    /// @notice The mapping of challengeKey to challenges.
    /// @dev There should only be one challenge per blockhash-celestiapointer pair.
    mapping(bytes32 => ChallengeDA) public daChallenges;

    /// @notice The fee required to make a challenge.
    /// @dev Is disabled by default.
    bool public isDAChallengeEnabled;

    /// @notice Returns the reference key for a DA challenge.
    /// @param _blockHash - The block hash of the block being challenged.
    /// @param _pointerIndex - The index of the celestia pointer being challenged.
    /// @return The reference key for the DA challenge.
    function dataRootInclusionChallengeKey(
        bytes32 _blockHash,
        uint8 _pointerIndex,
        uint32 _shareIndex
    ) public pure returns (bytes32) {
        return
            keccak256(abi.encodePacked(_blockHash, _pointerIndex, _shareIndex));
    }

    /// @notice Challenges the data root inclusion of a block.
    /// @param _blockIndex - The index of the block to challenge.
    /// @param _pointerIndex - The index of the celestia pointer to challenge.
    /// @return The index of the block being challenged.
    function challengeDataRootInclusion(
        uint256 _blockIndex,
        uint8 _pointerIndex,
        uint32 _shareIndex
    )
        external
        payable
        mustBeCanonical(_blockIndex)
        mustBeWithinChallengeWindow(_blockIndex) // TODO: use custom challenge period.
        requireChallengeFee
        returns (uint256)
    {
        require(isDAChallengeEnabled, "DA challenges are disabled");

        bytes32 challengeBlockHash = chain.chain(_blockIndex);
        bytes32 challengeKey = dataRootInclusionChallengeKey(
            challengeBlockHash,
            _pointerIndex,
            _shareIndex
        );

        // check if there is already a challenge for this block.
        ChallengeDA storage challenge = daChallenges[challengeKey];
        require(
            challenge.status == ChallengeDAStatus.None,
            "challenge already exists"
        );

        ICanonicalStateChain.Header memory header = chain.getHeaderByNum(
            _blockIndex
        );

        require(
            _pointerIndex < header.celestiaPointers.length,
            "invalid pointer index"
        );
        require(
            _shareIndex >= header.celestiaPointers[_pointerIndex].shareStart &&
                _shareIndex <
                header.celestiaPointers[_pointerIndex].shareStart +
                    header.celestiaPointers[_pointerIndex].shareLen,
            "invalid share index: not in pointers range"
        );

        // create a new challenge.
        daChallenges[challengeKey] = ChallengeDA(
            challengeBlockHash,
            _blockIndex,
            _pointerIndex,
            _shareIndex,
            msg.sender,
            block.timestamp + challengePeriod,
            ChallengeDAStatus.ChallengerInitiated,
            false
        );

        emit ChallengeDAUpdate(
            challengeBlockHash,
            _pointerIndex,
            _shareIndex,
            _blockIndex,
            block.timestamp + challengePeriod,
            ChallengeDAStatus.ChallengerInitiated
        );

        return _blockIndex;
    }

    /// @notice Defends the data root inclusion of a block.
    /// @param _challengeKey - The reference key of the challenge.
    /// @param _proof - The proof of inclusion.
    function defendDataRootInclusion(
        bytes32 _challengeKey,
        SharesProof memory _proof
    ) public nonReentrant {
        ChallengeDA storage challenge = daChallenges[_challengeKey];
        ICanonicalStateChain.Header memory header = chain.getHeaderByNum(
            challenge.blockIndex
        );

        require(
            challenge.status == ChallengeDAStatus.ChallengerInitiated,
            "challenge is not in the correct state"
        );
        require(
            header.celestiaPointers[challenge.pointerIndex].height ==
                _proof.attestationProof.tuple.height,
            "invalid celestia height"
        );

        // check the namespace
        require(_proof.namespace.equalTo(daNamespace), "invalid namespace");

        // verify the provided proof is valid – this also calls verifyAttestations.
        (bool success, ) = DAVerifier.verifySharesToDataRootTupleRoot(
            daOracle,
            _proof,
            _proof.attestationProof.tuple.dataRoot
        );
        require(success, "failed to verify shares to data root tuple root");

        // calculate squaresize
        (uint256 squaresize, ) = DAVerifier.computeSquareSizeFromRowProof(
            _proof.rowProofs[0]
        );

        // check that the share index is within the celestia pointer range.
        uint256 shareIndexInRow = _proof.shareProofs[0].beginKey;
        uint256 shareIndexInRowMajorOrder = shareIndexInRow +
            squaresize *
            _proof.rowProofs[0].key;
        require(
            shareIndexInRowMajorOrder == challenge.shareIndex,
            "proof must be provided for the challenged share index"
        );

        // update the challenge.
        challenge.status = ChallengeDAStatus.DefenderWon;
        emit ChallengeDAUpdate(
            challenge.blockHash,
            challenge.pointerIndex,
            challenge.shareIndex,
            challenge.blockIndex,
            challenge.expiry,
            ChallengeDAStatus.DefenderWon
        );

        // The defender can now call claimDAChallengeReward to claim the reward.
    }

    /// @notice Settles the data root inclusion challenge in favor of the challenger
    ///         if the defender does not respond within the challenge period.
    /// @param _challengeKey - The reference key of the challenge.
    function settleDataRootInclusion(
        bytes32 _challengeKey
    ) public nonReentrant {
        ChallengeDA storage challenge = daChallenges[_challengeKey];
        require(
            challenge.status == ChallengeDAStatus.ChallengerInitiated,
            "challenge is not in the correct state"
        );
        require(
            block.timestamp > challenge.expiry,
            "challenge has not expired"
        );

        // update the challenge.
        challenge.status = ChallengeDAStatus.ChallengerWon;
        emit ChallengeDAUpdate(
            challenge.blockHash,
            challenge.pointerIndex,
            challenge.shareIndex,
            challenge.blockIndex,
            challenge.expiry,
            ChallengeDAStatus.ChallengerWon
        );

        // rollback the chain.
        chain.rollback(challenge.blockIndex - 1, challenge.blockHash);

        // The challenger can now call claimDAChallengeReward to claim the reward.
    }

    /// @notice Toggles the data availability challenges on or off.
    /// @param _status - The status of the data availability challenges.
    function toggleDAChallenge(bool _status) external onlyOwner {
        isDAChallengeEnabled = _status;
    }

    function claimDAChallengeReward(
        bytes32 _challengeKey
    ) external nonReentrant {
        ChallengeDA storage challenge = daChallenges[_challengeKey];
        require(
            challenge.claimed == false,
            "challenge has already been claimed"
        );
        require(
            challenge.status == ChallengeDAStatus.ChallengerWon ||
                challenge.status == ChallengeDAStatus.DefenderWon,
            "challenge is not in the correct state"
        );

        challenge.claimed = true;
        if (challenge.status == ChallengeDAStatus.ChallengerWon) {
            (bool success, ) = challenge.challenger.call{value: challengeFee}(
                ""
            );
            require(success, "failed to pay challenger");
        } else {
            (bool success, ) = defender.call{value: challengeFee}("");
            require(success, "failed to pay defender");
        }
    }
}
