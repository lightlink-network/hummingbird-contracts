// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ChallengeBase.sol";

/// @title  ChallengeHeader
/// @author LightLink Hummingbird
/// @custom:version v1.0.0-alpha
/// @notice ChallengeHeader lets anyone challenge a block header against some basic validity checks.
///         If the header is invalid, the chain is rolled back to the previous block.
///         Note: This challenge is free and has no payout.
///
///         The challenge is made in a single step by calling invalidateHeader. This function directly checks
///         the validity of the header without requiring the defender to respond.
///
///         The following checks are made:
///         1. The epoch is greater than the previous epoch.
///         2. The l2Height is greater than the previous l2Height.
///         3. The prevHash is the previous block hash.
///         4. The bundle size is less than the max bundle size.
///
///         If any of these checks fail, the chain is rolled back to the previous block.
///         Just like with all challenges, the challenge window must be open.
abstract contract ChallengeHeader is ChallengeBase {
    /// @notice The reasons a header can be invalid.
    /// @param InvalidEpoch - The epoch is less than or equal to the previous epoch.
    /// @param InvalidL2Height - The l2Height is less than or equal to the previous l2Height.
    /// @param InvalidPrevHash - The prevHash is not the previous block hash.
    /// @param InvalidBundleSize - The bundle size is greater than the max bundle size.
    enum InvalidHeaderReason {
        InvalidEpoch,
        InvalidL2Height,
        InvalidPrevHash,
        InvalidBundleSize
    }

    /// @notice Emitted when a header is invalid.
    /// @param _blockIndex - The block index of the invalid header.
    /// @param _hash - The hash of the invalid header.
    /// @param _reason - The reason the header is invalid.
    event InvalidHeader(
        uint256 indexed _blockIndex,
        bytes32 indexed _hash,
        InvalidHeaderReason indexed _reason
    );

    /// @notice Whether the header challenge is enabled.
    /// @dev Is disabled by default.
    bool public isHeaderChallengeEnabled;

    /// @notice The maximum bundle size.
    uint256 public maxBundleSize;

    /// @notice Initializes the contract.
    function __ChallengeHeader_init() internal {
        maxBundleSize = 14000;
    }

    /// @notice Invalidate challenges a block header by checking that the header is valid.
    /// @param _blockIndex - The block index of the header to challenge.
    function invalidateHeader(
        uint256 _blockIndex
    )
        external
        mustBeCanonical(_blockIndex)
        mustBeWithinChallengeWindow(_blockIndex)
    {
        require(isHeaderChallengeEnabled, "header challenge is disabled");

        bytes32 _hash = chain.chain(_blockIndex);
        ICanonicalStateChain.Header memory header = chain.getHeaderByNum(
            _blockIndex
        );

        // check header validity.
        require(!_isHeaderValid(header, _hash, _blockIndex), "header is valid");

        // rollback the chain.
        chain.rollback(_blockIndex - 1, _hash);
    }

    /// @notice Checks if a header is valid.
    /// @param _header - The header to check.
    /// @param _hash - The hash of the header.
    /// @param _blockIndex - The block index of the header.
    /// @return True if the header is valid.
    function _isHeaderValid(
        ICanonicalStateChain.Header memory _header,
        bytes32 _hash,
        uint256 _blockIndex
    ) internal returns (bool) {
        // check that the blocks epoch is greater than the previous epoch.
        if (_header.epoch <= chain.getHeaderByNum(_blockIndex - 1).epoch) {
            emit InvalidHeader(
                _header.epoch,
                _hash,
                InvalidHeaderReason.InvalidEpoch
            );
            return false;
        }

        // check that the l2 height is greater than the previous l2 height.
        if (
            _header.l2Height <= chain.getHeaderByNum(_blockIndex - 1).l2Height
        ) {
            emit InvalidHeader(
                _header.epoch,
                _hash,
                InvalidHeaderReason.InvalidL2Height
            );
            return false;
        }

        // check that the prevHash is the previous block hash.
        if (_header.prevHash != chain.chain(_blockIndex - 1)) {
            emit InvalidHeader(
                _header.epoch,
                _hash,
                InvalidHeaderReason.InvalidPrevHash
            );
            return false;
        }

        // check that the bundle size is less than the max bundle size.
        if (
            _header.l2Height - chain.getHeaderByNum(_blockIndex - 1).l2Height >
            maxBundleSize
        ) {
            emit InvalidHeader(
                _header.epoch,
                _hash,
                InvalidHeaderReason.InvalidBundleSize
            );
            return false;
        }

        return true;
    }

    /// @notice Enables or disables the header challenge.
    /// @param _status - The status to set.
    function toggleHeaderChallenge(bool _status) external onlyOwner {
        isHeaderChallengeEnabled = _status;
    }

    /// @notice Sets the maximum bundle size.
    /// @param _maxBundleSize - The new maximum bundle size.
    function setMaxBundleSize(uint256 _maxBundleSize) external onlyOwner {
        maxBundleSize = _maxBundleSize;
    }

    uint256[50] private __gap;
}
