// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IRegistry {
	function getContractAddress(bytes32) external view returns (address);
}
