// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ITestnetGov {
	function isReward(address addr) external view returns (bool);
	function isVoter(address addr) external view returns (bool);
	function isStaker(address addr) external view returns (bool);
	function isMember(address) external view returns (bool);
	function getMember(uint256) external view returns (address);
	function getMemberLength() external view returns (uint256);
	function getReward(uint256) external view returns (address);
	function getNodeIdxFromMember(address) external view returns (uint256);
	function getMemberFromNodeIdx(uint256) external view returns (address);
	function getNodeLength() external view returns (uint256);
	function getNode(uint256) external view returns (bytes memory, bytes memory, bytes memory, uint);
	function getBallotInVoting() external view returns (uint256);
	function getVoter(uint256 idx) external view returns (address);
}
