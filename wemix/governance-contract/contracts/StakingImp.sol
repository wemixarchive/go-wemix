// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./GovChecker.sol";
import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import "./interface/IEnvStorage.sol";
import "./interface/IStaking.sol";
import "./interface/INCPStaking.sol";
import "./interface/INCPExit.sol";

contract StakingImp is GovChecker, UUPSUpgradeable, ReentrancyGuardUpgradeable, IStaking {
	//balance of each NCP
	mapping(address => uint256) private _balance;
	mapping(address => uint256) private _lockedBalance;
	uint256 private _totalLockedBalance;
	bool private revoked;

	//====NXTMeta====//
	event Staked(address indexed payee, uint256 amount, uint256 total, uint256 available);
	event Unstaked(address indexed payee, uint256 amount, uint256 total, uint256 available);
	event Locked(address indexed payee, uint256 amount, uint256 total, uint256 available);
	event Unlocked(address indexed payee, uint256 amount, uint256 total, uint256 available);
	event TransferLocked(address indexed payee, uint256 amount, uint256 total, uint256 available);
	event Revoked(address indexed owner, uint256 amount);

	//====Phase2-Staking====//
	event DelegateStaked(address indexed payee, uint256 amount, address indexed ncp, uint256 ncpTotalLocked, uint256 userTotalLocked);
	event DelegateUnstaked(address indexed payee, uint256 amount, address indexed ncp, uint256 ncpTotalLocked, uint256 userTotalLocked);
	event NCPAddrChanged(address indexed ncp);

	constructor() {
		_disableInitializers();
	}

	function init(address registry, bytes memory data) external initializer {
		_totalLockedBalance = 0;
		revoked = false;
		// _transferOwnership(_msgSender());
		__ReentrancyGuard_init();
		__Ownable_init();
		setRegistry(registry);

		// data is only for test purpose
		if (data.length == 0) return;

		// []{address, amount}
		address addr;
		uint amount;
		uint ix;
		uint eix;
		assembly {
			ix := add(data, 0x20)
		}
		eix = ix + data.length;
		while (ix < eix) {
			assembly {
				addr := mload(ix)
			}
			ix += 0x20;
			require(ix < eix);
			assembly {
				amount := mload(ix)
			}
			ix += 0x20;

			_balance[addr] = amount;
		}
	}

	receive() external payable {
		revert();
	}

	/**
	 * @dev Deposit from a sender.
	 */
	function deposit() external payable override nonReentrant notRevoked {
		require(msg.value > 0, "Deposit amount should be greater than zero");

		_balance[msg.sender] = _balance[msg.sender] + msg.value;

		if (IGov(getGovAddress()).isMember(msg.sender)) {
			uint256 minimum_staking = IEnvStorage(getEnvStorageAddress()).getStakingMin();
			//if minimum lock is going higher than current locked value, lock more
			if (minimum_staking > _lockedBalance[msg.sender] && availableBalanceOf(msg.sender) >= (minimum_staking - _lockedBalance[msg.sender])) {
				if (ncpStaking != address(0)) {
					uint256 pid = INCPStaking(ncpStaking).ncpToIdx(msg.sender);
					if (INCPStaking(ncpStaking).getUserInfo(pid, msg.sender).lastRewardClaimed != 0) {
						INCPStaking(ncpStaking).ncpDeposit(minimum_staking - _lockedBalance[msg.sender], payable(msg.sender));
					}
				}
				_lock(msg.sender, minimum_staking - _lockedBalance[msg.sender]);
			}
		}

		emit Staked(msg.sender, msg.value, _balance[msg.sender], availableBalanceOf(msg.sender));
	}

	/**
	 * @dev Withdraw for a sender.
	 * @param amount The amount of funds will be withdrawn and transferred to.
	 */
	function withdraw(uint256 amount) external override nonReentrant notRevoked {
		require(amount > 0, "Amount should be bigger than zero");

		//if minimum is changed unlock staked value
		uint256 minimum_staking = IEnvStorage(getEnvStorageAddress()).getStakingMin();

		bool unlockBalance = false;
		if (lockedBalanceOf(msg.sender) - _lockedUserBalanceToNCPTotal[msg.sender] >= minimum_staking + amount) {
			_unlock(msg.sender, amount);
			unlockBalance = true;
		}

		require(amount <= availableBalanceOf(msg.sender), "Withdraw amount should be equal or less than balance");

		_balance[msg.sender] = _balance[msg.sender] - amount;

		if (ncpStaking != address(0) && unlockBalance) {
			(bool succ, ) = payable(ncpStaking).call{ value: amount }("");
			require(succ, "Transfer to NCP staking failed");
			INCPStaking(ncpStaking).ncpWithdraw(amount, payable(msg.sender));
		} else {
			//payable(msg.sender).transfer(amount);
			(bool succ, ) = payable(msg.sender).call{ value: amount }("");
			require(succ, "Transfer to sender failed");
		}

		emit Unstaked(msg.sender, amount, _balance[msg.sender], availableBalanceOf(msg.sender));
	}

	/**
	 * @dev Lock fund
	 * @param payee The address whose funds will be locked.
	 * @param lockAmount The amount of funds will be locked.
	 */
	function lock(address payee, uint256 lockAmount) external override onlyGov {
		_lock(payee, lockAmount);
	}

	function lockMore(uint256 lockAmount) external onlyGovStaker {
		_lock(msg.sender, lockAmount);

		if (ncpStaking != address(0)) {
			uint256 pid = INCPStaking(ncpStaking).ncpToIdx(msg.sender);
			if (INCPStaking(ncpStaking).getUserInfo(pid, msg.sender).lastRewardClaimed != 0) {
				INCPStaking(ncpStaking).ncpDeposit(lockAmount, payable(msg.sender));
			}
		}
	}

	function _lock(address payee, uint256 lockAmount) internal {
		if (lockAmount == 0) return;
		require(_balance[payee] >= lockAmount, "Lock amount should be equal or less than balance");
		require(availableBalanceOf(payee) >= lockAmount, "Insufficient balance that can be locked");
		uint256 maximum = IEnvStorage(getEnvStorageAddress()).getStakingMax();

		_lockedBalance[payee] = _lockedBalance[payee] + lockAmount;
		require(_lockedBalance[payee] <= maximum, "Locked balance is larger than max");

		_totalLockedBalance = _totalLockedBalance + lockAmount;

		emit Locked(payee, lockAmount, _balance[payee], availableBalanceOf(payee));
	}

	/**
	 * @dev Transfer locked value(slashing, ncpLockMore, ncpUserTotal) to Ecosystem, NCPExit
	 * @param from The address whose funds will be transfered.
	 * @param slashing The amount of funds will be transfered.
	 * @param ext ncpLockMore + ncpUserTotal
	 */
	function transferLocked(address from, uint256 slashing, uint256 ext) external override onlyGov {
		INCPExit ncpExit = INCPExit(getContractAddress(bytes32("NCPExit")));

		// Ecosystem
		unlock(from, slashing);
		_balance[from] = _balance[from] - slashing;
		address ecosystem = getEcosystemAddress();
		_balance[ecosystem] = _balance[ecosystem] + slashing;

		// NCPLockMore
		uint256 ncpLockMore = ext - _lockedUserBalanceToNCPTotal[from];
		unlock(from, ncpLockMore);

		// To NCPExit
		uint256 transferedBalance = lockedBalanceOf(from);
		require(
			transferedBalance >= _lockedUserBalanceToNCPTotal[from],
			"transferedBalance must be greater than or equal to _lockedUserBalanceToNCPTotal."
		);

		unlock(from, transferedBalance);
		_balance[from] = _balance[from] - transferedBalance;
		ncpExit.depositExitAmount{ value: transferedBalance }(from, transferedBalance, _lockedUserBalanceToNCPTotal[from]);
		if (ncpStaking != address(0)) {
			_lockedUserBalanceToNCPTotal[from] = 0;
			_lockedUserBalanceToNCP[from][ncpStaking] = 0;
		}
		emit TransferLocked(from, slashing + transferedBalance, _balance[from], availableBalanceOf(from));
	}

	/**
	 * @dev Unlock fund
	 * @param payee The address whose funds will be unlocked.
	 * @param unlockAmount The amount of funds will be unlocked.
	 */
	function unlock(address payee, uint256 unlockAmount) public override onlyGov {
		_unlock(payee, unlockAmount);
	}

	function _unlock(address payee, uint256 unlockAmount) internal {
		if (unlockAmount == 0) return;
		require(_lockedBalance[payee] >= unlockAmount, "Unlock amount should be equal or less than balance locked");
		_lockedBalance[payee] = _lockedBalance[payee] - unlockAmount;
		_totalLockedBalance = _totalLockedBalance - unlockAmount;

		emit Unlocked(payee, unlockAmount, _balance[payee], availableBalanceOf(payee));
	}

	function balanceOf(address payee) public view override returns (uint256) {
		return _balance[payee];
	}

	function lockedBalanceOf(address payee) public view override returns (uint256) {
		return _lockedBalance[payee];
	}

	function availableBalanceOf(address payee) public view override returns (uint256) {
		return _balance[payee] - _lockedBalance[payee];
	}

	/**
	 * @dev Calculate voting weight which range between 0 and 100.
	 * @param payee The address whose funds were locked.
	 */
	function calcVotingWeight(address payee) public view override returns (uint256) {
		return calcVotingWeightWithScaleFactor(payee, 1e2);
	}

	/**
	 * @dev Calculate voting weight with a scale factor.
	 * @param payee The address whose funds were locked.
	 * @param factor The scale factor for weight. For instance:
	 *               if 1e1, result range is between 0 ~ 10
	 *               if 1e2, result range is between 0 ~ 100
	 *               if 1e3, result range is between 0 ~ 1000
	 */
	function calcVotingWeightWithScaleFactor(address payee, uint32 factor) public view override returns (uint256) {
		if (_lockedBalance[payee] == 0 || factor == 0) return 0;
		return (_lockedBalance[payee] * factor) / _totalLockedBalance;
	}

	function isRevoked() public view returns (bool) {
		return revoked;
	}

	modifier notRevoked() {
		require(!revoked, "Is revoked");
		_;
	}

	/**
	 * @dev Allows the owner to revoke the staking. Funds already staked are returned to the owner
	 */
	function revoke() public onlyOwner notRevoked {
		address contractOwner = owner();
		uint256 balance = address(this).balance;

		require(balance > 0, "balance = 0");

		payable(contractOwner).transfer(balance);
		revoked = true;

		emit Revoked(contractOwner, balance);
	}

	function upgradeStaking(address newImp) external onlyOwner {
		if (newImp != address(0)) {
			_authorizeUpgrade(newImp);
			_upgradeToAndCallUUPS(newImp, new bytes(0), false);
		}
	}

	function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

	/**
	 * @dev This empty reserved space is put in place to allow future versions to add new
	 * variables without shifting down storage in the inheritance chain.
	 * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
	 */
	uint256[49] private __gap;

	//Phase2 - Staking
	mapping(address /* NCP */ => mapping(address /* user */ => uint256 /* amount */)) private _lockedUserBalanceToNCP;
	mapping(address /* NCP */ => uint256 /* amount */) private _lockedUserBalanceToNCPTotal;

	address public ncpStaking;

	modifier onlyNCPStaking() {
		require(msg.sender == ncpStaking, "Only NCPStaking contract can call this function");
		_;
	}

	function userBalanceOf(address ncp, address user) external view override returns (uint256) {
		return _lockedUserBalanceToNCP[ncp][user];
	}

	function userTotalBalanceOf(address ncp) external view override returns (uint256) {
		return _lockedUserBalanceToNCPTotal[ncp];
	}

	function getRatioOfUserBalance(address ncp) external view override returns (uint256) {
		uint256 user = _lockedUserBalanceToNCPTotal[ncp];
		uint256 locked = _lockedBalance[ncp];
		if (user == 0 || locked == 0) return 0;
		return (user * 1e2) / locked;
	}

	function setNCPStaking(address _ncpStaking) external onlyOwner {
		require(_ncpStaking != address(0), "NCPStaking is the zero address");
		ncpStaking = _ncpStaking;
		emit NCPAddrChanged(_ncpStaking);
	}

	/**
	 * @dev Deposit from a user for delegate staking.
	 */
	function delegateDepositAndLockMore(address ncp) external payable override nonReentrant notRevoked onlyNCPStaking {
		require(msg.value > 0, "Deposit amount should be greater than zero");
		require(IGov(getGovAddress()).isMember(ncp), "NCP should be a member");

		// console.log("delegateDepositAndLockMore: msg.value: %s",ncp);

		uint256 userDepositValue = msg.value;
		//added value to ncp balance
		_balance[ncp] = _balance[ncp] + userDepositValue;

		uint256 minimum_staking = IEnvStorage(getEnvStorageAddress()).getStakingMin();
		uint256 maximum_staking = IEnvStorage(getEnvStorageAddress()).getStakingMax();
		require(
			minimum_staking <= _lockedBalance[ncp] && (_lockedBalance[ncp] + userDepositValue) <= maximum_staking,
			"user should be in staking range"
		);
		_lock(ncp, userDepositValue);

		_lockedUserBalanceToNCP[ncp][msg.sender] = _lockedUserBalanceToNCP[ncp][msg.sender] + userDepositValue;
		_lockedUserBalanceToNCPTotal[ncp] = _lockedUserBalanceToNCPTotal[ncp] + userDepositValue;

		emit DelegateStaked(msg.sender, userDepositValue, ncp, _lockedUserBalanceToNCPTotal[ncp], _lockedUserBalanceToNCP[ncp][msg.sender]);
	}

	/**
	 * @dev Withdraw for a user.
	 * @param amount The amount of funds will be unlocked, withdrawn and transferred to.
	 */
	function delegateUnlockAndWithdraw(address ncp, uint256 amount) external override nonReentrant notRevoked onlyNCPStaking {
		require(amount > 0, "Amount should be bigger than zero");
		require(IGov(getGovAddress()).isMember(ncp), "NCP should be a member");

		uint256 userWithdrawValue = amount;
		uint256 minimum_staking = IEnvStorage(getEnvStorageAddress()).getStakingMin();
		require(
			userWithdrawValue <= _lockedUserBalanceToNCP[ncp][msg.sender] && lockedBalanceOf(ncp) - userWithdrawValue >= minimum_staking,
			"Withdraw amount should be equal or less than balance"
		);
		_unlock(ncp, userWithdrawValue);

		_balance[ncp] = _balance[ncp] - userWithdrawValue;
		_lockedUserBalanceToNCP[ncp][msg.sender] = _lockedUserBalanceToNCP[ncp][msg.sender] - userWithdrawValue;
		_lockedUserBalanceToNCPTotal[ncp] = _lockedUserBalanceToNCPTotal[ncp] - userWithdrawValue;

		// payable(ncpStaking).transfer(userWithdrawValue);
		(bool succ, ) = payable(ncpStaking).call{ value: userWithdrawValue }("");
		require(succ, "Transfer to NCP staking failed");

		emit DelegateUnstaked(msg.sender, userWithdrawValue, ncp, _lockedUserBalanceToNCPTotal[ncp], _lockedUserBalanceToNCP[ncp][msg.sender]);
	}

	function getTotalLockedBalance() external view override returns (uint256) {
		return _totalLockedBalance;
	}
}
