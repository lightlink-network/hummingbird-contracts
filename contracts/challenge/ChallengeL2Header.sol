// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ChallengeBase.sol";

/// @title  ChallengeL2Header
/// @author LightLink Hummingbird
/// @custom:version v1.0.0-alpha
/// @notice ChallengeL2Header is a two party challenge game where the defender must provide
///         a valid L2 header to defend against a challenge.
///
///         The Challenge goes through the following steps:
///         1. A challenger initiates a challenge by calling challengeL2Header with the rblock
///            number and the number of the L2 block it should contain.
///         2. The defending block publisher must provide valid L2 headers to the chainOracle
///            for both the challenged block and the previous block.
///         3. If the headers are valid, the defender wins the challenge and receives the
///            challenge fee.
///         4. Otherwise the challenge expires and the challenger wins the challenge and the
///            block is rolled back.
contract ChallengeL2Header is ChallengeBase {
    /// @notice The different states a L2 header challenge can be in.
    /// @param None - The L2 header challenge has not been initiated.
    /// @param Initiated - The L2 header challenge has been initiated by the challenger.
    /// @param ChallengerWon - The L2 header challenge has been won by the challenger.
    /// @param DefenderWon - The L2 header challenge has been won by the defender.
    enum L2HeaderChallengeStatus {
        None,
        Initiated,
        ChallengerWon,
        DefenderWon
    }

    /// @notice The pointer to an L2 header.
    /// @param rblock - The rblock hash of the L2 header.
    /// @param number - The number of the L2 header.
    struct L2HeaderPointer {
        bytes32 rblock;
        uint256 number;
    }

    /// @notice The data structure for an L2 header challenge.
    /// @param header - The header being challenged.
    /// @param prevHeader - The previous header.
    /// @param challengeEnd - The end of the challenge period.
    /// @param challenger - The address of the challenger.
    /// @param status - The status of the challenge.
    struct L2HeaderChallenge {
        L2HeaderPointer header;
        L2HeaderPointer prevHeader;
        uint256 challengeEnd;
        address challenger;
        L2HeaderChallengeStatus status;
    }

    /// @notice Emitted when an L2 header challenge is updated.
    /// @param challengeHash - The hash of the challenge.
    /// @param l2Number - The number of the L2 header being challenged.
    /// @param rblock - The rblock hash of the L2 header.
    /// @param expiry - The expiry time of the challenge.
    /// @param status - The status of the challenge.
    event L2HeaderChallengeUpdate(
        bytes32 indexed challengeHash,
        uint256 indexed l2Number,
        bytes32 rblock,
        uint256 expiry,
        L2HeaderChallengeStatus indexed status
    );

    /// @notice Stores the L2 header challenges.
    mapping(bytes32 => L2HeaderChallenge) public l2HeaderChallenges;

    /// @notice Whether the L2 header challenge is enabled.
    /// @dev Is disabled by default.
    bool public isL2HeaderChallengeEnabled;

    /// @notice Challenges an L2 header by providing the rblock number and the L2 number.
    /// @param _rblockNum - The rblock number of the L2 header.
    /// @param _l2Num - The number of the L2 header.
    /// @return The hash of the challenge.
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
        require(isL2HeaderChallengeEnabled, "L2 header challenge is disabled");

        // 1. Load the rblock and the previous rblock
        bytes32 rblockHash = chain.chain(_rblockNum);
        ICanonicalStateChain.Header memory rblock = chain.getHeaderByNum(
            _rblockNum
        );
        ICanonicalStateChain.Header memory prevRBlock = chain.getHeaderByNum(
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
            _l2Num > prevRBlock.l2Height && _l2Num <= rblock.l2Height,
            "L2 header must be within the rblock bundle range"
        );

        // 4. Check that the L2 header is not the first in the first rblock
        require(
            !(_rblockNum == 1 && _l2Num == prevRBlock.l2Height + 1),
            "Cannot challenge the first L2 header in the first rblock"
        );

        // 5. Create pointer to the L2 header
        L2HeaderPointer memory header = L2HeaderPointer(rblockHash, _l2Num);

        // 6. Create a pointer the previous L2 header
        L2HeaderPointer memory prevHeader = L2HeaderPointer(
            rblockHash,
            _l2Num - 1
        );

        if (_l2Num == prevRBlock.l2Height + 1) {
            // If the L2 header is the first in the rblock, then the previous header is in the previous rblock
            prevHeader = L2HeaderPointer(rblock.prevHash, prevRBlock.l2Height);
        }

        // 7. Create the challenge
        l2HeaderChallenges[challengeHash] = L2HeaderChallenge(
            header,
            prevHeader,
            block.timestamp + challengePeriod,
            msg.sender,
            L2HeaderChallengeStatus.Initiated
        );

        // 8. Emit the challenge event
        emit L2HeaderChallengeUpdate(
            challengeHash,
            _l2Num,
            rblockHash,
            block.timestamp + challengePeriod,
            L2HeaderChallengeStatus.Initiated
        );

        return challengeHash;
    }

    /// @notice Defends an L2 header challenge by providing the L2 header and the previous L2 header.
    /// @param _challengeHash - The hash of the challenge.
    /// @param _headerHash - The hash of the L2 header.
    /// @param _headerPrevHash - The hash of the previous L2 header.
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

    /// @notice Settles an L2 header challenge by paying out the challenger.
    /// @param _challengeHash - The hash of the challenge.
    /// @dev Can only be called after the challenge period has ended and a
    ///      defender has not responded.
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

    /// @notice Returns the hash of an L2 header challenge.
    /// @param _rblockHash - The rblock hash of the L2 header.
    /// @param _l2Num - The number of the L2 header.
    function l2HeaderChallengeHash(
        bytes32 _rblockHash,
        uint256 _l2Num
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_rblockHash, _l2Num));
    }

    /// @notice Toggles the L2 header challenges on or off.
    /// @param _status - The status of the L2 header challenges.
    /// @dev Only the owner can call this function.
    function toggleL2HeaderChallenge(bool _status) external onlyOwner {
        isL2HeaderChallengeEnabled = _status;
    }
}
