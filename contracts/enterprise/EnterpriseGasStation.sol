// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import "./EnterprisePortal.sol";
import "./LLERC20.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

/**
 * @title Enterprise Gas Station.
 * @notice This contract manages gas plans, allows gas purchases, and handles gas claims for enterprises.
 */
contract EnterpriseGasStation is
    UUPSUpgradeable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    /// @notice The address of the Enterprise Portal contract.
    EnterprisePortal public enterprisePortalContract;

    /// @notice The address of the LightLink ERC20 token contract.
    LLERC20 public llToken;

    /// @notice The address of the reserve wallet.
    /// @dev The reserve wallet receives a % of the gas purchase cost.
    address public reserveWallet;

    /// @notice Struct to store gas plan details.
    /// @param id The ID of the gas plan.
    /// @param costInTokens The cost of the gas plan in tokens.
    /// @param gasAmount The amount of gas provided by the gas plan.
    /// @param isActive Whether the gas plan is active.
    struct GasPlan {
        uint id;
        uint costInTokens;
        uint256 gasAmount;
        bool isActive;
    }

    /// @notice Mapping of gas plans.
    mapping(uint => GasPlan) public gasPlans;

    /// @notice The ID of the next gas plan.
    /// @dev This is used to generate unique gas plan IDs.
    uint public nextPlanId;

    /**
     * @dev Emitted when a new gas plan is added.
     * @param planId The ID of the new gas plan.
     * @param costInTokens The cost of the gas plan in tokens.
     * @param gasAmount The amount of gas provided by the gas plan.
     * @param isActive Whether the gas plan is active.
     */
    event GasPlanAdded(
        uint indexed planId,
        uint costInTokens,
        uint256 gasAmount,
        bool isActive
    );

    /**
     * @dev Emitted when a gas plan's status is updated.
     * @param planId The ID of the gas plan.
     * @param isActive The new status of the gas plan.
     */
    event GasPlanStatusUpdated(uint indexed planId, bool isActive);

    /**
     * @dev Emitted when gas is purchased.
     * @param enterpriseId The ID of the enterprise.
     * @param planId The ID of the gas plan.
     * @param amount The amount of gas purchased.
     */
    event GasPurchased(
        address indexed enterpriseId,
        uint indexed planId,
        uint amount
    );

    /**
     * @dev Emitted when gas is claimed.
     * @param enterpriseAddress The public address of the enterprise.
     * @param amount The amount of gas claimed.
     */
    event GasClaimed(address indexed enterpriseAddress, uint256 amount);

    /**
     * @dev This function is a special internal function that's part of
     * the UUPS upgradeable contract's lifecycle. When you want to
     * upgrade the contract to a new version, _authorizeUpgrade is
     * called to check whether the upgrade is authorized, thus
     * preventing anyone from just upgrading the contract.
     * @dev Only the owner can call this function.
     */
    function _authorizeUpgrade(address) internal override onlyOwner {}

    /**
     * @notice Constructor for the Gas Management contract.
     * @param _enterprisePortalContract The address of the Enterprise Portal contract.
     * @param _llToken The address of the LLERC20 token contract.
     * @param _reserveWallet The address of the reserve wallet.
     */
    function initialize(
        address _enterprisePortalContract,
        address _llToken,
        address _reserveWallet
    ) public initializer {
        __Ownable_init(msg.sender);
        enterprisePortalContract = EnterprisePortal(_enterprisePortalContract);
        llToken = LLERC20(_llToken);
        //require(_reserveWallet != address(0), "Invalid reserve wallet address");
        reserveWallet = _reserveWallet;
    }

    /**
     * @notice Adds a new gas plan.
     * @param costInTokens The cost of the gas plan in tokens.
     * @param gasAmount The amount of gas provided by the gas plan.
     * @param isActive Whether the gas plan is active.
     */
    function addGasPlan(
        uint costInTokens,
        uint256 gasAmount,
        bool isActive
    ) external onlyOwner {
        gasPlans[nextPlanId] = GasPlan(
            nextPlanId,
            costInTokens,
            gasAmount,
            isActive
        );
        emit GasPlanAdded(nextPlanId, costInTokens, gasAmount, isActive);
        nextPlanId++;
    }

    /**
     * @notice Updates the status of an existing gas plan.
     * @param planId The ID of the gas plan to update.
     * @param isActive The new status of the gas plan.
     */
    function setGasPlanStatus(uint planId, bool isActive) external onlyOwner {
        require(gasPlans[planId].id == planId, "Gas plan does not exist");
        gasPlans[planId].isActive = isActive;
        emit GasPlanStatusUpdated(planId, isActive);
    }

    /**
     * @notice Claims gas from an enterprise's balance.
     * @param enterpriseAddress The address of the enterprise.
     * @param amount The amount of gas to claim.
     * @dev Only whitelisted contracts can claim gas.
     * @dev The caller must be a whitelisted contract.
     */
    function claimGas(
        address enterpriseAddress,
        uint256 amount
    ) external nonReentrant {
        require(
            enterprisePortalContract.isContractWhitelisted(
                enterpriseAddress,
                msg.sender
            ),
            "Caller is not a whitelisted contract"
        );
        enterprisePortalContract.deductGas(enterpriseAddress, amount);
        emit GasClaimed(enterpriseAddress, amount);
    }

    /**
     * @notice Purchases gas for an enterprise using a specified gas plan.
     * @param enterpriseAddress The address of the enterprise.
     * @param planId The ID of the gas plan to use.
     */
    function purchaseGas(
        address enterpriseAddress,
        uint planId
    ) external nonReentrant {
        require(gasPlans[planId].isActive, "Gas plan is not active");

        uint256 cost = gasPlans[planId].costInTokens;
        uint256 gasAmount = gasPlans[planId].gasAmount;

        uint256 burnAmount = cost / 10;
        uint256 reserveAmount = cost - burnAmount;

        // Transfer tokens from buyer to this contract
        llToken.transferFrom(msg.sender, address(this), cost);

        // Burn 10% of tokens
        llToken.burn(burnAmount);

        // Transfer 90% to reserve wallet
        llToken.transfer(reserveWallet, reserveAmount);

        // Add gas to enterprise
        enterprisePortalContract.addGas(enterpriseAddress, gasAmount);
        emit GasPurchased(enterpriseAddress, planId, gasAmount);
    }
}
