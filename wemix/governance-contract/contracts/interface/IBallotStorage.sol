// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IBallotStorage {
	function createBallotForMember(
		uint256,
		uint256,
		uint256,
		address,
		address,
		address,
		address,
		address,
		bytes memory,
		bytes memory,
		bytes memory,
		uint
	) external;

	function createBallotForAddress(uint256, uint256, uint256, address, address) external returns (uint256);
	function createBallotForVariable(uint256, uint256, uint256, address, bytes32, uint256, bytes memory) external returns (uint256);
	function createBallotForExit(uint256, uint256, uint256) external;
	function createVote(uint256, uint256, address, uint256, uint256) external;
	function finalizeBallot(uint256, uint256) external;
	function startBallot(uint256, uint256, uint256) external;
	function updateBallotMemo(uint256, bytes memory) external;
	function updateBallotDuration(uint256, uint256) external;
	function updateBallotMemberLockAmount(uint256, uint256) external;

	function getBallotPeriod(uint256) external view returns (uint256, uint256, uint256);
	function getBallotVotingInfo(uint256) external view returns (uint256, uint256, uint256);
	function getBallotState(uint256) external view returns (uint256, uint256, bool);

	function getBallotBasic(
		uint256
	) external view returns (uint256, uint256, uint256, address, bytes memory, uint256, uint256, uint256, uint256, bool, uint256);

	function getBallotMember(
		uint256
	) external view returns (address, address, address, address, bytes memory, bytes memory, bytes memory, uint256, uint256);
	function getBallotAddress(uint256) external view returns (address);
	function getBallotVariable(uint256) external view returns (bytes32, uint256, bytes memory);
	function getBallotForExit(uint256) external view returns (uint256, uint256);
}
