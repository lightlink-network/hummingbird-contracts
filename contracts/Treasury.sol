// SPDX-License-Identifier: UNLICENSED
// LightLink Hummingbird v0.0.1

// NOT FOR RELEASE

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

// Treasury is a contract that holds extra funds for the rollup chain. It is used to
// to incentivize/reimburse honest challenges. It is owned by the hummingbird challenge contract.
//
// The treasury contract operates optimistically. If there are not enough funds to cover
// a payment, a bond is created and the id is returned. The bond can be claimed by the
// owner of the bond at a later time.
//
// Typical income sources are:
// - % of Fees from deposits via the native bridge.
// - % of Fees from unsuccessful challenges.
// - Donations from the DAO and other sources.
contract Treasury is Ownable {
    struct Bond {
        address owner;
        uint256 amount;
    }

    event BondCreated(
        uint256 indexed id,
        address indexed owner,
        uint256 amount
    );
    event BondClaimed(
        uint256 indexed id,
        address indexed owner,
        uint256 amount
    );

    Bond[] public bonds;
    uint256 public commitments = 0;

    constructor() Ownable(msg.sender) {}

    // claimBond allows the owner of a bond to claim the bond. The bond must have
    // enough funds to cover the amount.
    function claimBond(
        uint256 _id,
        uint256 _amount,
        address payable _beneficary
    ) external {
        Bond storage bond = bonds[_id];
        require(bond.owner == msg.sender, "only the bond owner can claim");
        require(
            bond.amount >= _amount,
            "bond amount must be greater than amount"
        );

        bond.amount -= _amount;
        commitments -= _amount;

        _beneficary.transfer(_amount);
        emit BondClaimed(_id, msg.sender, _amount);
    }

    // makePayment makes a payment to the beneficiary. If there are not enough funds
    // in the treasury, a bond is created and the id is returned.
    // If -1 is returned, the payment was successful, otherwise its the id of the bond.
    function makePayment(
        address payable _beneficary,
        uint256 _amount
    ) public onlyOwner returns (int256) {
        if (_amount > disposableFunds()) {
            // create bond
            uint256 id = bonds.length;
            bonds.push(Bond(_beneficary, _amount));

            // update commitments
            commitments += _amount;

            emit BondCreated(id, _beneficary, _amount);
            return int256(id);
        }

        _beneficary.transfer(_amount);
        return -1;
    }

    function transfer(
        address payable _beneficary,
        uint256 _amount
    ) external payable onlyOwner {
        _beneficary.transfer(_amount);
    }

    function transferToken(
        address payable _beneficary,
        uint256 _amount,
        address _token
    ) external payable onlyOwner {
        (bool success, ) = _token.call(
            abi.encodeWithSignature(
                "transfer(address,uint256)",
                _beneficary,
                _amount
            )
        );
        require(success, "transfer failed");
    }

    function disposableFunds() public view returns (uint256) {
        return address(this).balance - commitments;
    }
}
