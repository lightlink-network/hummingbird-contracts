// SPDX-License-Identifier: MIT
// LightLink Hummingbird v0.2.0

pragma solidity ^0.8.0;

import "./ChallengeBase.sol";

// ChallengeL2Header is a two party challenge game where the defender must provide
// a valid L2 header to defend against a challenge.
//
// The Challenge goes through the following steps:
// 1. A challenger initiates a challenge by calling challengeL2Header with the rblock number and the number of the L2 block it should contain.
// 2. The defending block publisher must provide valid L2 headers to the chainOracle for both the challenged block and the previous block.
// 3. If the headers are valid, the defender wins the challenge and receives the challenge fee.
// 4. Otherwise the challenge expires and the challenger wins the challenge and the block is rolled back.
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
        uint256 indexed l2Number,
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
        ICanonicalStateChain.Header memory rblock = chain.getBlock(_rblockNum);
        ICanonicalStateChain.Header memory prevRBlock = chain.getBlock(
            _rblockNum - 1
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
            _l2Num,
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

        // 0. Check that the header and previous headers are part of the correct rblocks
        // - This prevents rolled back l2 headers from being used to defend
        require(
            chainOracle.headerToRblock(_headerHash) == challenge.header.rblock,
            "l2 header not loaded for the given rblock"
        );
        require(
            chainOracle.headerToRblock(_headerPrevHash) ==
                challenge.prevHeader.rblock,
            "previous l2 header not loaded for the given rblock"
        );

        // 1. Load the header and previous header from the ChainOracle
        IChainOracle.L2Header memory header = chainOracle.getHeader(
            _headerHash
        );
        IChainOracle.L2Header memory prevHeader = chainOracle.getHeader(
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
            header.timestamp >= prevHeader.timestamp,
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
            challenge.header.number,
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
            challenge.header.number,
            challenge.header.rblock,
            challenge.challengeEnd,
            L2HeaderChallengeStatus.ChallengerWon
        );

        // pay out the challenger
        (bool success, ) = challenge.challenger.call{value: challengeFee}("");
        require(success, "failed to pay challenger");
    }

    function l2HeaderChallengeHash(
        bytes32 _rblockHash,
        uint256 _l2Num
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_rblockHash, _l2Num));
    }
}
