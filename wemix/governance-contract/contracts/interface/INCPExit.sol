// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface INCPExit {
	/**
	 * @dev Sets a new administrator.
	 * @param _newAdministrator Address of the new administrator.
	 */
	function setAdministrator(address _newAdministrator) external;

	/**
	 * @dev Sets a new administrator setter.
	 * @param _newAdministratorSetter Address of the new administrator setter.
	 */
	function setAdministratorSetter(address _newAdministratorSetter) external;

	/**
	 * @dev Exits from the contract.
	 * @param exitNcp Address of the NCP to exit.
	 * @param totalAmount Total amount of ether to exit with.
	 * @param lockedUserBalanceToNCPTotal Total locked user balance to NCP.
	 */
	function depositExitAmount(address exitNcp, uint256 totalAmount, uint256 lockedUserBalanceToNCPTotal) external payable;

	/**
	 * @dev Withdraws amount for a user.
	 * @param exitNcp Address of the ncp
	 * @param exitUser Address of the user to withdraw for.
	 * @param amount Amount to withdraw.
	 */
	function withdrawForUser(address exitNcp, address exitUser, uint256 amount) external;

	/**
	 * @dev Withdraws amount for the administrator.
	 * @param exitNcp Address of the NCP to withdraw from.
	 */
	function withdrawForAdministrator(address exitNcp, uint256 amount, address to) external;

	function getAvailableAmountForAdministrator(address exitNcp) external view returns (uint256);

	function getLockedUserBalanceToNCPTotal(address exitNcp) external view returns (uint256);
}
