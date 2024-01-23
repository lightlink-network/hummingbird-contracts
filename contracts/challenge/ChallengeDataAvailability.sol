// SPDX-License-Identifier: UNLICENSED
// LightLink Hummingbird v0.0.1

pragma solidity ^0.8.0;

import "./ChallengeBase.sol";
import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";


// TODO: settle on expiry if no one responds

// no constructor
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
        uint256 rootNonce;
        BinaryMerkleProof proof;
    }

    event ChallengeDAUpdate(
        bytes32 indexed _blockHash,
        uint256 indexed _blockIndex,
        uint256 _expiry,
        ChallengeDAStatus indexed _status
    );

    // a mapping of block hashes to challenges.
    // Note: There should only be one challenge per block.
    mapping(bytes32 => ChallengeDA) public daChallenges;

    function challengeDataRootInclusion(
        uint256 _blockIndex
    )
        external
        payable
        mustBeCanonical(_blockIndex)
        mustBeWithinChallengeWindow(_blockIndex)
        returns (uint256)
    {
        bytes32 h = chain.chain(_blockIndex);

        // check if there is already a challenge for this block.
        ChallengeDA storage challenge = daChallenges[h];
        require(
            challenge.status == ChallengeDAStatus.None,
            "challenge already exists"
        );

        // create a new challenge.
        daChallenges[h] = ChallengeDA(
            _blockIndex,
            msg.sender,
            block.timestamp + challengePeriod,
            ChallengeDAStatus.ChallengerInitiated
        );

        emit ChallengeDAUpdate(
            h,
            _blockIndex,
            block.timestamp + challengePeriod,
            ChallengeDAStatus.ChallengerInitiated
        );

        return _blockIndex;
    }

    function defendDataRootInclusion(
        bytes32 _blockHash,
        ChallengeDAProof memory _proof
    ) public nonReentrant {
        ChallengeDA storage challenge = daChallenges[_blockHash];
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
            _blockHash,
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
    function settleDataRootInclusion(bytes32 _blockhash) public nonReentrant {
        ChallengeDA storage challenge = daChallenges[_blockhash];
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
            _blockhash,
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
