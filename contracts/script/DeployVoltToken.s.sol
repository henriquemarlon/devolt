// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {Volt} from "@contracts/Volt.sol";

contract DeployVoltToken is Script {
    function run() external {
        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));
        Volt volt = new Volt(vm.envAddress("INITIAL_ONWER"));
        vm.stopBroadcast();
        console.log(
            "Volt address:",
            address(volt),
            "deployed at:",
            block.timestamp
        );
    }
}