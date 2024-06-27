// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/// @author @seunghwalee
interface INCPStaking {
	struct UserInfo {
		uint256 amount;
		uint256 rewardDebt;
		uint256 pendingReward;
		uint256 pendingAmountReward;
		uint256 lastRewardClaimed;
	}
	function ncpDeposit(uint256 amount, address payable to) external payable;
	function ncpWithdraw(uint256 amount, address payable to) external payable;
	function getUserInfo(uint256 pid, address account) external view returns (UserInfo memory info);
	function ncpToIdx(address ncp) external view returns (uint256);
}
