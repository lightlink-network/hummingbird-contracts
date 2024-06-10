// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract LLERC20 is ERC20 {
    constructor(uint256 initialSupply) ERC20("LL Token", "LL") {
        _mint(msg.sender, initialSupply);
    }

    function burn(uint256 amount) external {
        _burn(msg.sender, amount);
    }
}
