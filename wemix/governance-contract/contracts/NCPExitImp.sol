// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./interface/INCPStaking.sol";
import "./interface/INCPExit.sol";
import "./interface/IGovStaking.sol";
import "./GovChecker.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";

contract NCPExitImp is GovChecker, INCPExit, UUPSUpgradeable, ReentrancyGuardUpgradeable {
	/* =========== STATE VARIABLES ===========*/
	using AddressUpgradeable for address payable;

	address private _administrator;
	address private _administratorSetter;

	mapping(address /* ncp address */ => /* user + a */ uint256) private _receivedTotalAmount;
	mapping(address /* ncp address */ => /* user */ uint256) private _lockedUserBalanceToNCPTotal;

	/* =========== MODIFIERES ===========*/
	modifier onlyGovStaking() {
		require(msg.sender == getStakingAddress(), "Only governance staking contract can call this function.");
		_;
	}

	modifier onlyNcpStaking() {
		require(msg.sender == IGovStaking(getStakingAddress()).ncpStaking(), "Only NcpStaking can call this function.");
		_;
	}

	modifier onlyAdministrator() {
		require(msg.sender == _administrator, "Only Administrator can call this function.");
		_;
	}

	modifier onlyAdministratorSetter() {
		require(msg.sender == _administratorSetter, "Caller is not AdministratorSetter.");
		_;
	}

	receive() external payable {}

	/* =========== FUNCTIONS ===========*/
	/// @custom:oz-upgrades-unsafe-allow constructor
	constructor() {
		_disableInitializers();
	}
	function initialize(address _registry) external initializer {
		__Ownable_init();
		__ReentrancyGuard_init();
		setRegistry(_registry);
		_administrator = owner();
		_administratorSetter = owner();
	}

	function setAdministrator(address _newAdministrator) external override onlyAdministratorSetter {
		require(_newAdministrator != address(0), "Address should be non-zero");
		_administrator = _newAdministrator;
	}

	function setAdministratorSetter(address _newAdministratorSetter) external override onlyAdministratorSetter {
		require(_newAdministratorSetter != address(0), "Address should be non-zero");
		_administratorSetter = _newAdministratorSetter;
	}

	function depositExitAmount(
		address exitNcp,
		uint256 totalAmount,
		uint256 lockedUserBalanceToNCPTotal
	) external payable override nonReentrant onlyGovStaking {
		require(totalAmount == msg.value);
		_receivedTotalAmount[exitNcp] = totalAmount;

		// administrator
		_lockedUserBalanceToNCPTotal[exitNcp] = lockedUserBalanceToNCPTotal;
	}

	function withdrawForUser(address exitNcp, address exitUser, uint256 amount) external override nonReentrant onlyNcpStaking {
		require(_lockedUserBalanceToNCPTotal[exitNcp] >= amount, "_lockedUserBalanceToNCPTotal[exitNcp] >= amount");

		_receivedTotalAmount[exitNcp] = _receivedTotalAmount[exitNcp] - amount;
		_lockedUserBalanceToNCPTotal[exitNcp] = _lockedUserBalanceToNCPTotal[exitNcp] - amount;

		payable(exitUser).sendValue(amount);
	}

	function withdrawForAdministrator(address exitNcp, uint256 amount, address to) external override nonReentrant onlyAdministrator {
		// For admin
		require(_receivedTotalAmount[exitNcp] - _lockedUserBalanceToNCPTotal[exitNcp] >= amount);
		_receivedTotalAmount[exitNcp] = _receivedTotalAmount[exitNcp] - amount;

		payable(to).sendValue(amount);
	}

	function getAvailableAmountForAdministrator(address exitNcp) external view override onlyAdministrator returns (uint256) {
		// For admin
		return _receivedTotalAmount[exitNcp] - _lockedUserBalanceToNCPTotal[exitNcp];
	}

	function upgradeNCPExit(address newImp) external onlyOwner {
		if (newImp != address(0)) {
			_authorizeUpgrade(newImp);
			_upgradeToAndCallUUPS(newImp, new bytes(0), false);
		}
	}

	function getLockedUserBalanceToNCPTotal(address exitNcp) external view override returns (uint256) {
		return _lockedUserBalanceToNCPTotal[exitNcp];
	}

	function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
}
