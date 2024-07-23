pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import "contracts/L2/L1Block.sol";

contract Predeploy is Script {
    function setUp() public {}

    function run() public {
        vm.startBroadcast(vm.envUint("PREDEPLOY_DEPLOY_KEY"));
        L1Block l1block = new L1Block();

        vm.stopBroadcast();
    }
}
