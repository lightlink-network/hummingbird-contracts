pragma solidity 0.8.22;

import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract SystemConfig is UUPSUpgradeable, OwnableUpgradeable {
    uint256 private _startBlock;
    address private _optimismPortal;

    /// @notice This function is a special internal function that's part of
    ///         the UUPS upgradeable contract's lifecycle. When you want to
    ///         upgrade the contract to a new version, _authorizeUpgrade is
    ///         called to check whether the upgrade is authorized, thus
    ///         preventing anyone from just upgrading the contract.
    /// @dev Only the owner can call this function.
    function _authorizeUpgrade(address) internal override onlyOwner {}

    function initialize(
        uint256 __startBlock,
        address __optimismPortal
    ) public initializer {
        __Ownable_init(msg.sender);
        _startBlock = __startBlock;
        _optimismPortal = __optimismPortal;
    }

    /// @notice Getter for the StartBlock number.
    function startBlock() external view returns (uint256 startBlock_) {
        return _startBlock;
    }

    // @notice Getter for the OptimismPortal address.
    function optimismPortal() external view returns (address) {
        return _optimismPortal;
    }

    /// @notice Setter for the StartBlock number.
    function setStartBlock(uint256 __startBlock) external onlyOwner {
        _startBlock = __startBlock;
    }

    /// @notice Setter for the OptimismPortal address.
    function setOptimismPortal(address __optimismPortal) external onlyOwner {
        _optimismPortal = __optimismPortal;
    }

    uint256[50] private __gap;
}
