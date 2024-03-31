// SPDX-License-Identifier: MIT
// LightLink Hummingbird v0.2.0

pragma solidity ^0.8.0;

import "./ChallengeBase.sol";
import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";

// ChallengeDataAvailability is a challenge for verifying a rollup blocks data
// root has been included. (via Celestia Blobstream).
//
// This is a challenge game between two parties: the challenger and the defender.
// There can only be one challenge per rblock hash.
//
// The Challenge goes through the following steps:
// 1. A challenger initiates a challenge by calling challengeDataRootInclusion.
// 2. The defending block publisher must provide a proof of inclusion for the
//    data root. If the proof is valid, the defender wins the challenge and
//    receives the challenge fee.
// 3. Otherwise the challenge expires and the challenger wins the challenge and
//    the block is rolled back.
//
// You can trigger a challenge easily via the hummingbird client:
//      `hb challenge-da <block-index>`.

abstract contract ChallengeDataAvailability is ChallengeBase {
    enum ChallengeDAStatus {
        None,
        ChallengerInitiated,
        ChallengerWon,
        DefenderWon
    }

    struct ChallengeDA {
        bytes32 blockHash;
        uint256 blockIndex;
        uint8 pointerIndex;
        address challenger;
        uint256 expiry;
        ChallengeDAStatus status;
    }

    struct ChallengeDAProof {
        uint256 rootNonce;
        DataRootTuple dataRootTuple;
        BinaryMerkleProof proof;
    }

    event ChallengeDAUpdate(
        bytes32 indexed _blockHash,
        uint256 indexed _pointerIndex,
        uint256 _blockIndex,
        uint256 _expiry,
        ChallengeDAStatus indexed _status
    );

    // a mapping of challengeKey to challenges.
    // Note: There should only be one challenge per blockhash-celestiapointer pair.
    mapping(bytes32 => ChallengeDA) public daChallenges;

    function dataRootInclusionChallengeKey(
        bytes32 _blockHash,
        uint8 _pointerIndex
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_blockHash, _pointerIndex));
    }

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
            ChallengeDAStatus.ChallengerInitiated
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

        // pay out the reward.
        // use call to prevent failing receiver is a contract.
        (bool success, ) = defender.call{value: challengeFee}("");
        require(success, "failed to pay defender");
    }

    // settle the challenge in favor of the challenger if the defender does not respond
    // within the challenge period.
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

        // pay out the reward.
        // use call to prevent failing receiver is a contract.
        (bool success, ) = challenge.challenger.call{value: challengeFee}("");
        require(success, "failed to pay challenger");

        // rollback the chain.
        chain.rollback(challenge.blockIndex - 1);
    }
}
