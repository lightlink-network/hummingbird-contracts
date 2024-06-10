// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title Enterprise Contract
 * @notice This contract manages the creation of enterprises, their gas balances, and whitelisted contracts.
 */
contract EnterprisePortal is UUPSUpgradeable, OwnableUpgradeable {
    /**
     * @dev Struct to store enterprise details.
     * @param owner The owner of the enterprise.
     * @param gasBalance The gas balance of the enterprise.
     * @param whitelistedContracts The whitelisted contracts for the enterprise.
     */
    struct EnterpriseStruct {
        address owner;
        uint256 gasBalance;
        mapping(address => bool) whitelistedContracts;
    }

    /// @notice Mapping of enterprises to their addresses.
    mapping(address => EnterpriseStruct) private enterprises;

    /**
     * @notice Emitted when a new enterprise is created.
     * @param enterpriseAddress The public address of the enterprise.
     */
    event EnterpriseCreated(address indexed enterpriseAddress);

    /**
     * @dev Emitted when a contract is added to the whitelist.
     * @param enterpriseAddress The public address of the enterprise.
     * @param whitelistedContract The address of the whitelisted contract.
     */
    event WhitelistedContractAdded(
        address indexed enterpriseAddress,
        address indexed whitelistedContract
    );

    /**
     * @dev Emitted when a contract is removed from the whitelist.
     * @param enterpriseAddress The public address of the enterprise.
     * @param whitelistedContract The address of the removed contract.
     */
    event WhitelistedContractRemoved(
        address indexed enterpriseAddress,
        address indexed whitelistedContract
    );

    /**
     * @dev Emitted when an enterprise is removed.
     * @param enterpriseAddress The address of the enterprise.
     */
    event EnterpriseRemoved(address indexed enterpriseAddress);

    /**
     * @dev Emitted when gas is added to an enterprise.
     * @param enterpriseAddress The public address of the enterprise.
     * @param amount The amount of gas added.
     */
    event GasAdded(address indexed enterpriseAddress, uint256 amount);

    /**
     * @dev Emitted when gas is deducted from an enterprise.
     * @param enterpriseAddress The public address of the enterprise.
     * @param amount The amount of gas deducted.
     */
    event GasDeducted(address indexed enterpriseAddress, uint256 amount);

    /**
     * @dev This function is a special internal function that's part of
     * the UUPS upgradeable contract's lifecycle. When you want to
     * upgrade the contract to a new version, _authorizeUpgrade is
     * called to check whether the upgrade is authorized, thus
     * preventing anyone from just upgrading the contract.
     * @dev Only the owner can call this function.
     */
    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}

    /**
     * @notice Constructor for the EnterprisePortal contract.
     */
    function initialize() public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
    }

    modifier onlyEnterpriseOwner(address enterpriseAddress) {
        require(
            enterprises[enterpriseAddress].owner == msg.sender,
            "Not authorized"
        );
        _;
    }

    /**
     * @notice Creates a new enterprise.
     * @dev Only the owner can create a new enterprise.
     * @param _enterpriseAddress The public address of the enterprise.
     */
    function createEnterprise(address _enterpriseAddress) external onlyOwner {
        require(_enterpriseAddress != address(0), "Invalid address");
        require(
            enterprises[_enterpriseAddress].owner == address(0),
            "Enterprise already exists"
        );

        enterprises[_enterpriseAddress].owner = msg.sender;
        emit EnterpriseCreated(_enterpriseAddress);
    }

    /**
     * @notice Adds a whitelisted contract to an enterprise.
     * @param _enterpriseAddress The public address of the enterprise.
     * @param _whitelistedContract The address of the contract to be whitelisted.
     */
    function addWhitelistedContract(
        address _enterpriseAddress,
        address _whitelistedContract
    ) external onlyEnterpriseOwner(_enterpriseAddress) {
        require(
            !enterprises[_enterpriseAddress].whitelistedContracts[
                _whitelistedContract
            ],
            "Contract already whitelisted"
        );

        enterprises[_enterpriseAddress].whitelistedContracts[
            _whitelistedContract
        ] = true;
        emit WhitelistedContractAdded(_enterpriseAddress, _whitelistedContract);
    }

    /**
     * @notice Removes a whitelisted contract from an enterprise.
     * @param _enterpriseAddress The public address of the enterprise.
     * @param _whitelistedContract The address of the contract to be removed from the whitelist.
     */
    function removeWhitelistedContract(
        address _enterpriseAddress,
        address _whitelistedContract
    ) external onlyEnterpriseOwner(_enterpriseAddress) {
        require(
            enterprises[_enterpriseAddress].whitelistedContracts[
                _whitelistedContract
            ],
            "Contract not whitelisted"
        );

        delete enterprises[_enterpriseAddress].whitelistedContracts[
            _whitelistedContract
        ];
        emit WhitelistedContractRemoved(
            _enterpriseAddress,
            _whitelistedContract
        );
    }

    /**
     * @notice Removes an enterprise.
     * @param _enterpriseAddress The public address of the enterprise.
     */
    function removeEnterprise(
        address _enterpriseAddress
    ) external onlyEnterpriseOwner(_enterpriseAddress) {
        delete enterprises[_enterpriseAddress];
        emit EnterpriseRemoved(_enterpriseAddress);
    }

    /**
     * @notice Adds gas to an enterprise.
     * @dev Only the owner can add gas to an enterprise.
     * @param _enterpriseAddress The public address of the enterprise.
     * @param _amount The amount of gas to be added.
     */
    function addGas(
        address _enterpriseAddress,
        uint256 _amount
    ) external onlyOwner {
        enterprises[_enterpriseAddress].gasBalance += _amount;
        emit GasAdded(_enterpriseAddress, _amount);
    }

    /**
     * @notice Deducts gas from an enterprise.
     * @param _enterpriseAddress The public address of the enterprise.
     * @param _amount The amount of gas to be deducted.
     */
    function deductGas(
        address _enterpriseAddress,
        uint256 _amount
    ) external onlyEnterpriseOwner(_enterpriseAddress) {
        require(
            enterprises[_enterpriseAddress].gasBalance >= _amount,
            "Not enough gas"
        );

        enterprises[_enterpriseAddress].gasBalance -= _amount;
        emit GasDeducted(_enterpriseAddress, _amount);
    }

    /**
     * @notice Checks if a contract is whitelisted for an enterprise.
     * @param _enterpriseAddress The public address of the enterprise.
     * @param _whitelistedContract The address of the contract to check if it is whitelisted.
     * @return True if the contract is whitelisted, otherwise false.
     */
    function isContractWhitelisted(
        address _enterpriseAddress,
        address _whitelistedContract
    ) external view returns (bool) {
        require(
            enterprises[_enterpriseAddress].owner != address(0),
            "Enterprise does not exist"
        );

        return
            enterprises[_enterpriseAddress].whitelistedContracts[
                _whitelistedContract
            ];
    }

    /**
     * @notice Gets the gas balance of an enterprise.
     * @param _enterpriseAddress The public address of the enterprise.
     * @return gasBalance The gas balance of the enterprise.
     */
    function getEnterpriseGasBalance(
        address _enterpriseAddress
    ) external view returns (uint256) {
        require(
            enterprises[_enterpriseAddress].owner != address(0),
            "Enterprise does not exist"
        );

        return enterprises[_enterpriseAddress].gasBalance;
    }
}
