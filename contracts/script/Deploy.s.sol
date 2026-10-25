// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {MockNoAARecovery} from "../src/MockNoAARecovery.sol";

// Avoids a forge-std dependency.
interface Vm {
    function envUint(string calldata name) external view returns (uint256);
    function envString(string calldata name) external view returns (string memory);
    function startBroadcast(uint256 privateKey) external;
    function stopBroadcast() external;
    function writeFile(string calldata path, string calldata data) external;
    function toString(address value) external pure returns (string memory);
    function toString(uint256 value) external pure returns (string memory);
}

contract Deploy {
    Vm private constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    function run() external {
        uint256 startBlock = block.number;

        vm.startBroadcast(vm.envUint("DEPLOYER_PRIVATE_KEY"));
        MockNoAARecovery c =
            new MockNoAARecovery(uint64(vm.envUint("RECOVERY_DELAY")), uint64(vm.envUint("LOCK_DURATION")));
        vm.stopBroadcast();

        vm.writeFile(
            vm.envString("DEPLOY_OUT"),
            string.concat("NOAA_ADDRESS=", vm.toString(address(c)), "\nNOAA_BLOCK=", vm.toString(startBlock), "\n")
        );
    }
}
