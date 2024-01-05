pragma solidity ^0.8.0;

import "./ChallengeBase.sol";
import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";

// TODO: ADD Non-reentrancy
// TODO: settle on expiry if no one responds

abstract contract ChallengeDataAvailability is ChallengeBase {
    enum ChallengeDAStatus {
        None,
        ChallengerInitiated,
        DefenderResponded,
        ChallengerWon,
        DefenderWon
    }

    struct ChallengeDA {
        uint256 blockIndex;
        address challenger;
        uint256 expiry;
        ChallengeDAStatus status;
    }

    struct ChallengeDAProof {
        uint256 blockIndex;
        uint256 rootNonce;
        BinaryMerkleProof proof;
    }

    event ChallengeDAUpdate(
        uint256 indexed _blockIndex,
        uint256 _expiry,
        ChallengeDAStatus indexed _status
    );

    // one challenge per block.
    mapping(uint256 => ChallengeDA) public daChallenges;

    function challengeDataRootInclusion(
        uint256 _blockIndex
    )
        external
        payable
        mustBeCanonical(_blockIndex)
        mustBeWithinChallengeWindow(_blockIndex)
        requireChallengeFee
        returns (uint256)
    {
        // check if there is already a challenge for this block.
        ChallengeDA storage challenge = daChallenges[_blockIndex];
        require(
            challenge.status == ChallengeDAStatus.None,
            "challenge already exists"
        );

        // create a new challenge.
        daChallenges[_blockIndex] = ChallengeDA(
            _blockIndex,
            msg.sender,
            block.timestamp + challengePeriod,
            ChallengeDAStatus.ChallengerInitiated
        );

        emit ChallengeDAUpdate(
            _blockIndex,
            block.timestamp + challengePeriod,
            ChallengeDAStatus.ChallengerInitiated
        );

        return _blockIndex;
    }

    function proveDataRootInclusion(
        uint256 _challengeId,
        ChallengeDAProof memory _proof
    ) public {
        ChallengeDA storage challenge = daChallenges[_challengeId];
        require(
            challenge.status == ChallengeDAStatus.ChallengerInitiated,
            "challenge is not in the correct state"
        );

        ICanonicalStateChain.Header memory header = chain.headers(
            chain.chain(challenge.blockIndex)
        );

        // verify the proof.
        require(
            daOracle.verifyAttestation(
                _proof.rootNonce,
                DataRootTuple(
                    uint256(header.celestiaHeight),
                    header.celestiaDataRoot
                ),
                _proof.proof
            ),
            "invalid proof"
        );

        // update the challenge.
        challenge.status = ChallengeDAStatus.DefenderWon;
        emit ChallengeDAUpdate(
            challenge.blockIndex,
            challenge.expiry,
            ChallengeDAStatus.DefenderWon
        );

        // pay out the reward.
        payable(defender).transfer(challengeReward);
    }

    // settle the challenge in favor of the challenger if the defender does not respond.
    function settle(uint256 _challengeId) public {
        ChallengeDA storage challenge = daChallenges[_challengeId];
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
            challenge.blockIndex,
            challenge.expiry,
            ChallengeDAStatus.ChallengerWon
        );

        // pay out the reward.
        payable(challenge.challenger).transfer(challengeReward);
    }
}
