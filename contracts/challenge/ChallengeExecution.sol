pragma solidity ^0.8.0;

import "./ChallengeBase.sol";

// TODO: NOT IMPLEMENTED YET

abstract contract ChallengeExecution is ChallengeBase {
    enum ChallengeStatus {
        None,
        ChallengerInitiated,
        DefenderResponded,
        ChallengerWon,
        DefenderWon
    }

    struct Challenge {
        ChallengeStatus status;
        uint256 headerIndex;
        uint256 blockIndex;
        uint64 mipSteps;
        bytes32 assertionRoot;
        bytes32 finalSystemState;
        uint256 mipsChallengeId;
    }

    mapping(uint256 => Challenge) public execChallenges;

    function initiateChallenge(
        uint256 _headerIndex,
        uint256 _blockIndex,
        uint64 _mipSteps,
        bytes32 _assertionRoot,
        bytes32 _finalSystemState
    )
        external
        payable
        mustBeCanonical(_blockIndex)
        mustBeWithinChallengeWindow(_blockIndex)
        requireChallengeFee
        returns (uint256)
    {
        // check if there is already a challenge for this block.
        Challenge storage challenge = execChallenges[_blockIndex];
        require(
            challenge.status == ChallengeStatus.None,
            "challenge already exists"
        );

        // create a new challenge.
        execChallenges[_blockIndex] = Challenge(
            ChallengeStatus.ChallengerInitiated,
            _headerIndex,
            _blockIndex,
            _mipSteps,
            _assertionRoot,
            _finalSystemState,
            0
        );

        // TODO: implement data loading

        // return the challenge id.
        return _blockIndex;
    }
}
