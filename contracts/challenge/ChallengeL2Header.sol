// SPDX-License-Identifier: MIT
// LightLink Hummingbird UNRELEASED v0.1.1

// NOT FOR RELEASE
// NOT IMPLEMENTED YET

pragma solidity ^0.8.0;

import "./ChallengeBase.sol";

contract ChallengeL2Header is ChallengeBase {
    enum L2HeaderChallengeStatus {
        None,
        Initiated,
        ChallengerWon,
        DefenderWon
    }

    struct L2HeaderPointer {
        bytes32 rblock;
        uint256 number;
    }

    struct L2HeaderChallenge {
        L2HeaderPointer header; // The header being challenged.
        L2HeaderPointer prevHeader; // The previous header.
        uint256 challengeEnd; // The end of the challenge period.
        address challenger; // The address of the challenger.
        L2HeaderChallengeStatus status; // The status of the challenge.
    }

    event L2HeaderChallengeUpdate(
        bytes32 indexed challengeHash,
        bytes32 indexed l2Number,
        bytes32 rblock,
        uint256 expiry,
        L2HeaderChallengeStatus indexed status
    );

    mapping(bytes32 => L2HeaderChallenge) public l2HeaderChallenges;

    function challengeL2Header(
        uint256 _rblockNum,
        uint256 _l2Num
    )
        external
        payable
        mustBeCanonical(_rblockNum)
        mustBeWithinChallengeWindow(_rblockNum)
        requireChallengeFee
        returns (bytes32)
    {
        // 1. Load the rblock and the previous rblock
        bytes32 rblockHash = chain.chain(_rblockNum);
        ICanonicalStateChain.Header memory rblock = chain.headers(rblockHash);
        ICanonicalStateChain.Header memory prevRBlock = chain.headers(
            rblock.prevHash
        );

        // 2. Check that this exact L2 header is not already challenged
        bytes32 challengeHash = keccak256(abi.encodePacked(rblockHash, _l2Num));
        require(
            l2HeaderChallenges[challengeHash].status ==
                L2HeaderChallengeStatus.None,
            "challenge already exists"
        );

        // 3. Check that the L2 header is within the rblock bundle range
        require(
            _l2Num > prevRBlock.l2Height && _l2Num < rblock.l2Height,
            "L2 header must be within the rblock bundle range"
        );

        // 4. Create pointer to the L2 header
        L2HeaderPointer memory header = L2HeaderPointer(rblockHash, _l2Num);

        // 5. Create a pointer the previous L2 header
        L2HeaderPointer memory prevHeader = L2HeaderPointer(
            rblockHash,
            _l2Num - 1
        );

        if (_l2Num == prevRBlock.l2Height) {
            // If the L2 header is the first in the rblock, then the previous header is in the previous rblock
            prevHeader = L2HeaderPointer(
                prevRBlock.prevHash,
                prevRBlock.l2Height
            );
        }

        // 6. Create the challenge
        l2HeaderChallenges[challengeHash] = L2HeaderChallenge(
            header,
            prevHeader,
            block.timestamp + challengePeriod,
            msg.sender,
            L2HeaderChallengeStatus.Initiated
        );

        // 7. Emit the challenge event
        emit L2HeaderChallengeUpdate(
            challengeHash,
            bytes32(_l2Num),
            rblockHash,
            block.timestamp + challengePeriod,
            L2HeaderChallengeStatus.Initiated
        );

        return challengeHash;
    }

    function defendL2Header(
        bytes32 _challengeHash,
        bytes32 _headerHash,
        bytes32 _headerPrevHash
    ) external nonReentrant {
        L2HeaderChallenge storage challenge = l2HeaderChallenges[
            _challengeHash
        ];
        require(
            challenge.status == L2HeaderChallengeStatus.Initiated,
            "challenge is not in the correct state"
        );

        // 1. Load the header and previous header from the ChainOracle
        IChainOracle.L2Header memory header = chainOracle.headers(_headerHash);
        IChainOracle.L2Header memory prevHeader = chainOracle.headers(
            _headerPrevHash
        );

        // 2. Check the headers has the correct number
        require(
            header.number == challenge.header.number,
            "header number does not match"
        );
        require(
            prevHeader.number == challenge.prevHeader.number,
            "previous header number does not match"
        );

        // 3. Check the blocks are sequential
        require(
            header.parentHash == _headerPrevHash,
            "header does not point to the previous header"
        );

        // 4. Check the timestamp is correct
        require(
            header.timestamp > prevHeader.timestamp,
            "header timestamp is too late"
        );
        require(
            header.timestamp < block.timestamp,
            "header timestamp is in the future"
        );
        require(
            prevHeader.timestamp < block.timestamp,
            "previous header timestamp is the future"
        );

        // finalise the challenge
        challenge.status = L2HeaderChallengeStatus.DefenderWon;

        // payout the caller
        (bool success, ) = payable(msg.sender).call{value: challengeFee}("");
        require(success, "failed to pay defender");

        // emit the event
        emit L2HeaderChallengeUpdate(
            _challengeHash,
            bytes32(challenge.header.number),
            challenge.header.rblock,
            challenge.challengeEnd,
            L2HeaderChallengeStatus.DefenderWon
        );
    }

    function settleL2HeaderChallenge(bytes32 _challengeHash) external {
        L2HeaderChallenge storage challenge = l2HeaderChallenges[
            _challengeHash
        ];
        require(
            challenge.status == L2HeaderChallengeStatus.Initiated,
            "challenge is not in the correct state"
        );
        require(
            block.timestamp > challenge.challengeEnd,
            "challenge period has not ended"
        );

        challenge.status = L2HeaderChallengeStatus.ChallengerWon;

        emit L2HeaderChallengeUpdate(
            _challengeHash,
            bytes32(challenge.header.number),
            challenge.header.rblock,
            challenge.challengeEnd,
            L2HeaderChallengeStatus.ChallengerWon
        );

        // pay out the challenger
        (bool success, ) = challenge.challenger.call{value: challengeFee}("");
        require(success, "failed to pay challenger");
    }
}
