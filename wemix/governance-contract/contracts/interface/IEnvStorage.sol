// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IEnvStorage {
	function setBlocksPerByBytes(bytes memory) external;
	function setBallotDurationMinByBytes(bytes memory) external;
	function setBallotDurationMaxByBytes(bytes memory) external;
	function setStakingMinByBytes(bytes memory) external;
	function setStakingMaxByBytes(bytes memory) external;
	function setMaxIdleBlockIntervalByBytes(bytes memory) external;
	function setBlockCreationTimeByBytes(bytes memory _value) external;
	function setBlockRewardAmountByBytes(bytes memory _value) external;
	function setMaxPriorityFeePerGasByBytes(bytes memory _value) external;
	function setBallotDurationMinMax(uint256 _min, uint256 _max) external;
	function setBlockRewardDistributionMethodByBytes(bytes memory _value) external;
	function setGasLimitAndBaseFeeByBytes(bytes memory _value) external;
	function setMaxBaseFeeByBytes(bytes memory _value) external;
	function setBallotDurationMinMaxByBytes(bytes memory _value) external;
	function setStakingMinMaxByBytes(bytes memory _value) external;
	function getBlockCreationTime() external view returns (uint256);
	function getBlockRewardAmount() external view returns (uint256);
	function getMaxPriorityFeePerGas() external view returns (uint256);
	function getStakingMinMax() external view returns (uint256, uint256);
	function getBlockRewardDistributionMethod() external view returns (uint256, uint256, uint256, uint256);
	function getGasLimitAndBaseFee() external view returns (uint256, uint256, uint256);
	function getMaxBaseFee() external view returns (uint256);
	function getBlocksPer() external view returns (uint256);
	function getStakingMin() external view returns (uint256);
	function getStakingMax() external view returns (uint256);
	function getBallotDurationMin() external view returns (uint256);
	function getBallotDurationMax() external view returns (uint256);
	function getBallotDurationMinMax() external view returns (uint256, uint256);
	function getMaxIdleBlockInterval() external view returns (uint256);
	function checkVariableCondition(bytes32 envKey, bytes memory envVal) external pure returns (bool);
	function setVariable(bytes32 envKey, bytes memory envVal) external;
}
