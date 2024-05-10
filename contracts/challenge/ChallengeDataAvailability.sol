// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ChallengeBase.sol";
import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";

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
        uint8 _pointerIndex
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_blockHash, _pointerIndex));
    }

    /// @notice Challenges the data root inclusion of a block.
    /// @param _blockIndex - The index of the block to challenge.
    /// @param _pointerIndex - The index of the celestia pointer to challenge.
    /// @return The index of the block being challenged.
    function challengeDataRootInclusion(
        uint256 _blockIndex,
        uint8 _pointerIndex
    )
        external
        payable
        mustBeCanonical(_blockIndex)
        mustBeWithinChallengeWindow(_blockIndex) // TODO: use custom challenge period.
        requireChallengeFee
        returns (uint256)
    {
        require(isDAChallengeEnabled, "DA challenges are disabled");

        bytes32 h = chain.chain(_blockIndex);
        bytes32 challengeKey = dataRootInclusionChallengeKey(h, _pointerIndex);

        // check if there is already a challenge for this block.
        ChallengeDA storage challenge = daChallenges[challengeKey];
        require(
            challenge.status == ChallengeDAStatus.None,
            "challenge already exists"
        );
        require(
            _pointerIndex <
                chain.getHeaderByNum(_blockIndex).celestiaPointers.length,
            "invalid pointer index"
        );

        // create a new challenge.
        daChallenges[challengeKey] = ChallengeDA(
            h,
            _blockIndex,
            _pointerIndex,
            msg.sender,
            block.timestamp + challengePeriod,
            ChallengeDAStatus.ChallengerInitiated,
            false
        );

        emit ChallengeDAUpdate(
            h,
            _pointerIndex,
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
        ChallengeDAProof memory _proof
    ) public nonReentrant {
        ChallengeDA storage challenge = daChallenges[_challengeKey];
        require(
            challenge.status == ChallengeDAStatus.ChallengerInitiated,
            "challenge is not in the correct state"
        );

        ICanonicalStateChain.Header memory header = chain.getHeaderByNum(
            challenge.blockIndex
        );

        require(
            header.celestiaPointers[challenge.pointerIndex].height ==
                _proof.dataRootTuple.height,
            "invalid celestia height"
        );

        // verify the proof.
        require(
            daOracle.verifyAttestation(
                _proof.rootNonce,
                _proof.dataRootTuple,
                _proof.proof
            ),
            "invalid proof"
        );

        // update the challenge.
        challenge.status = ChallengeDAStatus.DefenderWon;
        emit ChallengeDAUpdate(
            challenge.blockHash,
            challenge.pointerIndex,
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
            challenge.blockIndex,
            challenge.expiry,
            ChallengeDAStatus.ChallengerWon
        );

        // rollback the chain.
        chain.rollback(challenge.blockIndex - 1);

        // The challenger can now call claimDAChallengeReward to claim the reward.
    }

    /// @notice Toggles the data availability challenges on or off.
    /// @param _status - The status of the data availability challenges.
    function toggleDAChallenge(bool _status) external onlyOwner {
        isDAChallengeEnabled = _status;
    }

    function claimDAChallengeReward(bytes32 _challengeKey) external {
        ChallengeDA storage challenge = daChallenges[_challengeKey];
        require(challenge.claimed == false, "challenge has already been claimed");
        require(challenge.status == ChallengeDAStatus.ChallengerWon || challenge.status == ChallengeDAStatus.DefenderWon, "challenge is not in the correct state");

        if (challenge.status == ChallengeDAStatus.ChallengerWon) {
            (bool success, ) = challenge.challenger.call{value: challengeFee}("");
            require(success, "failed to pay challenger");
        } else {
            (bool success, ) = defender.call{value: challengeFee}("");
            require(success, "failed to pay defender");
        }

       challenge.claimed = true;
    }
}
