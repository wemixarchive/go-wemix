// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
// pragma abicod

import "../abstract/EnvConstants.sol";
import "../abstract/AEnvStorage.sol";
import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import "../interface/IEnvStorage.sol";

contract EnvStorageImp is
    AEnvStorage,
    EnvConstants,
    UUPSUpgradeable,
    IEnvStorage
{
    event UpgradeImplementation(address indexed implementation);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address _registry,
        bytes32[] memory names,
        uint256[] memory infos
    ) public initializer {
        require(
            _registry != _getImplementation(),
            "registry should not be same as implementation"
        );
        __Ownable_init();
        setRegistry(_registry);

        for (uint i = 0; i < infos.length; i++) {
            uint256 temp = getUint(names[i]);
            require(infos[i] != 0 || temp != 0, "invalid variable");
            if (temp == 0) {
                setUint(names[i], infos[i]);
            }
        }
    }
    function upgradeTo(address newImplementation) public override onlyOwner {
        _upgradeToAndCallUUPS(newImplementation, "", false);
        emit UpgradeImplementation(newImplementation);
    }

    function getBlocksPer() public view override returns (uint256) {
        return getUint(BLOCKS_PER_NAME);
    }

    function getBallotDurationMin() public view override returns (uint256) {
        return getUint(BALLOT_DURATION_MIN_NAME);
    }

    function getBallotDurationMax() public view override returns (uint256) {
        return getUint(BALLOT_DURATION_MAX_NAME);
    }

    function getStakingMin() public view override returns (uint256) {
        return getUint(STAKING_MIN_NAME);
    }

    function getStakingMax() public view override returns (uint256) {
        return getUint(STAKING_MAX_NAME);
    }

    // function getGasPrice() public override view returns (uint256) {
    //     return getUint(GAS_PRICE_NAME);
    // }

    function getMaxIdleBlockInterval() public view override returns (uint256) {
        return getUint(MAX_IDLE_BLOCK_INTERVAL_NAME);
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyGov {}

    function setBlocksPer(uint256 _value) public onlyGov {
        setUint(BLOCKS_PER_NAME, _value);
    }

    function setBallotDurationMin(uint256 _value) public onlyGov {
        setUint(BALLOT_DURATION_MIN_NAME, _value);
    }

    function setBallotDurationMax(uint256 _value) public onlyGov {
        setUint(BALLOT_DURATION_MAX_NAME, _value);
    }

    function setStakingMin(uint256 _value) public onlyGov {
        setUint(STAKING_MIN_NAME, _value);
    }

    function setStakingMax(uint256 _value) public onlyGov {
        setUint(STAKING_MAX_NAME, _value);
    }

    // function setGasPrice(uint256 _value)  public override onlyGov {
    //     setUint(GAS_PRICE_NAME, _value);
    // }

    function setMaxIdleBlockInterval(uint256 _value) public onlyGov {
        setUint(MAX_IDLE_BLOCK_INTERVAL_NAME, _value);
    }

    function setBlocksPerByBytes(bytes memory _value) public override onlyGov {
        setBlocksPer(toUint(_value));
    }

    function setBallotDurationMinByBytes(
        bytes memory _value
    ) public override onlyGov {
        setBallotDurationMin(toUint(_value));
    }

    function setBallotDurationMaxByBytes(
        bytes memory _value
    ) public override onlyGov {
        setBallotDurationMax(toUint(_value));
    }

    function setStakingMinByBytes(bytes memory _value) public override onlyGov {
        setStakingMin(toUint(_value));
    }

    function setStakingMaxByBytes(bytes memory _value) public override onlyGov {
        setStakingMax(toUint(_value));
    }

    // function setGasPriceByBytes(bytes memory _value )  public override onlyGov {
    //     setGasPrice(toUint(_value));
    // }

    function setMaxIdleBlockIntervalByBytes(
        bytes memory _value
    ) public override onlyGov {
        setMaxIdleBlockInterval(toUint(_value));
    }

    //=======NXTMeta=======/

    function getBallotDurationMinMax()
        public
        view
        override
        returns (uint256, uint256)
    {
        return (
            getUint(BALLOT_DURATION_MIN_NAME),
            getUint(BALLOT_DURATION_MAX_NAME)
        );
    }

    function getStakingMinMax()
        public
        view
        override
        returns (uint256, uint256)
    {
        return (getUint(STAKING_MIN_NAME), getUint(STAKING_MAX_NAME));
    }

    function getBlockCreationTime() public view override returns (uint256) {
        return getUint(BLOCK_CREATION_TIME_NAME);
    }

    function getBlockRewardAmount() public view override returns (uint256) {
        return getUint(BLOCK_REWARD_AMOUNT_NAME);
    }

    function getMaxPriorityFeePerGas() public view override returns (uint256) {
        return getUint(MAX_PRIORITY_FEE_PER_GAS_NAME);
    }

    function getMaxBaseFee() public view override returns (uint256) {
        return getUint(MAX_BASE_FEE_NAME);
    }

    function getBlockRewardDistributionMethod()
        public
        view
        override
        returns (uint256, uint256, uint256, uint256)
    {
        return (
            getUint(BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME),
            getUint(BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME),
            getUint(BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME),
            getUint(BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME)
        );
    }

    function getGasLimitAndBaseFee()
        public
        view
        override
        returns (uint256, uint256, uint256)
    {
        return (
            getUint(BLOCK_GASLIMIT_NAME),
            getUint(BASE_FEE_MAX_CHANGE_RATE_NAME),
            getUint(GAS_TARGET_PERCENTAGE_NAME)
        );
    }

    // function getStakingRewardAddress() public override view returns(address){
    //     return getAddress(STAKING_REWARD_ADDRESS_NAME);
    // }

    // function getEcofundAddress() public override view returns(address){
    //     return getAddress(ECOFUND_ADDRESS_NAME);
    // }

    // function getMaintenanceAddress() public override view returns(address){
    //     return getAddress(MAINTENANCE_ADDRESS_NAME);
    // }

    function setBallotDurationMinMax(
        uint256 _min,
        uint256 _max
    ) public override onlyGov {
        require(
            _min <= _max,
            "Minimum duration must be smaller and equal than maximum duration"
        );
        setUint(BALLOT_DURATION_MIN_NAME, _min);
        setUint(BALLOT_DURATION_MAX_NAME, _max);
    }

    function setStakingMinMax(uint256 _min, uint256 _max) public onlyGov {
        require(
            _min <= _max,
            "Minimum staking must be smaller and equal than maximum staking"
        );
        setUint(STAKING_MIN_NAME, _min);
        setUint(STAKING_MAX_NAME, _max);
    }

    function setBlockCreationTime(uint256 _value) public onlyGov {
        setUint(BLOCK_CREATION_TIME_NAME, _value);
    }

    function setBlockRewardAmount(uint256 _value) public onlyGov {
        setUint(BLOCK_REWARD_AMOUNT_NAME, _value);
    }

    function setMaxPriorityFeePerGas(uint256 _value) public onlyGov {
        setUint(MAX_PRIORITY_FEE_PER_GAS_NAME, _value);
    }

    function setMaxBaseFee(uint256 _value) public onlyGov {
        setUint(MAX_BASE_FEE_NAME, _value);
    }

    function setBlockRewardDistributionMethod(
        uint256 _block_producer,
        uint256 _staking_reward,
        uint256 _ecofund,
        uint256 _maintenance
    ) public onlyGov {
        require(
            (_block_producer + _staking_reward + _ecofund + _maintenance) ==
                DENOMINATOR,
            "Wrong reward distrubtion ratio"
        );
        setUint(BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME, _block_producer);
        setUint(BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME, _staking_reward);
        setUint(BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME, _ecofund);
        setUint(BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME, _maintenance);
    }

    function setGasLimitAndBaseFee(
        uint256 _block_GasLimit,
        uint256 _baseFeeMaxChangeRate,
        uint256 _gasTargetPercentage,
        uint256 _maxBaseFee
    ) public onlyGov {
        setUint(BLOCK_GASLIMIT_NAME, _block_GasLimit);
        setUint(BASE_FEE_MAX_CHANGE_RATE_NAME, _baseFeeMaxChangeRate);
        setUint(GAS_TARGET_PERCENTAGE_NAME, _gasTargetPercentage);
        setUint(MAX_BASE_FEE_NAME, _maxBaseFee);
    }

    function setBallotDurationMinMaxByBytes(
        bytes memory _value
    ) public override onlyGov {
        (uint256 _min, uint256 _max) = to2Uint(_value);
        setBallotDurationMinMax(_min, _max);
    }

    function setStakingMinMaxByBytes(
        bytes memory _value
    ) public override onlyGov {
        (uint256 _min, uint256 _max) = to2Uint(_value);
        setStakingMinMax(_min, _max);
    }

    function setBlockCreationTimeByBytes(
        bytes memory _value
    ) public override onlyGov {
        setBlockCreationTime(toUint(_value));
    }

    function setBlockRewardAmountByBytes(
        bytes memory _value
    ) public override onlyGov {
        setBlockRewardAmount(toUint(_value));
    }

    function setMaxPriorityFeePerGasByBytes(
        bytes memory _value
    ) public override onlyGov {
        setMaxPriorityFeePerGas(toUint(_value));
    }

    function setMaxBaseFeeByBytes(bytes memory _value) public override onlyGov {
        setMaxBaseFee(toUint(_value));
    }

    function setBlockRewardDistributionMethodByBytes(
        bytes memory _value
    ) public override onlyGov {
        (
            uint256 _block_producer,
            uint256 _staking_reward,
            uint256 _ecosystem,
            uint256 _maintenance
        ) = to4Uint(_value);
        setBlockRewardDistributionMethod(
            _block_producer,
            _staking_reward,
            _ecosystem,
            _maintenance
        );
    }

    function setGasLimitAndBaseFeeByBytes(
        bytes memory _value
    ) public override onlyGov {
        (
            uint256 _block_GasLimit,
            uint256 _baseFeeMaxChangeRate,
            uint256 _gasTargetPercentage,
            uint256 _maxBaseFee
        ) = to4Uint(_value);
        setGasLimitAndBaseFee(
            _block_GasLimit,
            _baseFeeMaxChangeRate,
            _gasTargetPercentage,
            _maxBaseFee
        );
    }

    function checkVariableCondition(
        bytes32 envKey,
        bytes memory envVal
    ) external pure override returns (bool) {
        if (envKey == BLOCK_REWARD_DISTRIBUTION_METHOD_NAME) {
            (
                uint256 _block_producer,
                uint256 _staking_reward,
                uint256 _ecofund,
                uint256 _maintenance
            ) = abi.decode(envVal, (uint256, uint256, uint256, uint256));
            if (
                (_block_producer + _staking_reward + _ecofund + _maintenance) !=
                DENOMINATOR
            ) return false;
        } else if (
            envKey == STAKING_MIN_MAX_NAME ||
            envKey == BALLOT_DURATION_MIN_MAX_NAME
        ) {
            (uint256 min, uint256 max) = abi.decode(envVal, (uint256, uint256));
            if (min > max) return false;
        } else if (envKey == BLOCK_CREATION_TIME_NAME) {
            uint256 time = abi.decode(envVal, (uint256));
            if (time < 1000) return false;
        }
        return true;
    }

    function setVariable(
        bytes32 envKey,
        bytes memory envVal
    ) external override {
        if (envKey == BLOCKS_PER_NAME) {
            setBlocksPerByBytes(envVal);
        } else if (envKey == BALLOT_DURATION_MIN_MAX_NAME) {
            setBallotDurationMinMaxByBytes(envVal);
        } else if (envKey == STAKING_MIN_MAX_NAME) {
            setStakingMinMaxByBytes(envVal);
        }
        // else if (envKey == GAS_PRICE_NAME) {
        //     setGasPriceByBytes(envVal);
        // }
        else if (envKey == MAX_IDLE_BLOCK_INTERVAL_NAME) {
            setMaxIdleBlockIntervalByBytes(envVal);
        } else if (envKey == BLOCK_CREATION_TIME_NAME) {
            setBlockCreationTimeByBytes(envVal);
        } else if (envKey == BLOCK_REWARD_AMOUNT_NAME) {
            setBlockRewardAmountByBytes(envVal);
        } else if (envKey == MAX_PRIORITY_FEE_PER_GAS_NAME) {
            setMaxPriorityFeePerGasByBytes(envVal);
        } else if (envKey == BLOCK_REWARD_DISTRIBUTION_METHOD_NAME) {
            setBlockRewardDistributionMethodByBytes(envVal);
        } else if (envKey == GASLIMIT_AND_BASE_FEE_NAME) {
            setGasLimitAndBaseFeeByBytes(envVal);
        } else if (envKey == MAX_BASE_FEE_NAME) {
            setMaxBaseFeeByBytes(envVal);
        }
    }

    function toBytes32(
        bytes memory _input
    ) internal pure returns (bytes32 _output) {
        assembly {
            _output := mload(add(_input, 32))
        }
    }

    function toInt(bytes memory _input) internal pure returns (int256 _output) {
        assembly {
            _output := mload(add(_input, 32))
        }
    }

    function toUint(
        bytes memory _input
    ) internal pure returns (uint256 _output) {
        assembly {
            _output := mload(add(_input, 32))
        }
    }

    function to2Uint(
        bytes memory _input
    ) internal pure returns (uint256 _output0, uint256 _output1) {
        assembly {
            _output0 := mload(add(_input, 32))
            _output1 := mload(add(_input, 64))
        }
    }

    function to3Uint(
        bytes memory _input
    )
        internal
        pure
        returns (uint256 _output0, uint256 _output1, uint256 _output2)
    {
        assembly {
            _output0 := mload(add(_input, 32))
            _output1 := mload(add(_input, 64))
            _output2 := mload(add(_input, 96))
        }
    }

    function to4Uint(
        bytes memory _input
    )
        internal
        pure
        returns (
            uint256 _output0,
            uint256 _output1,
            uint256 _output2,
            uint256 _output3
        )
    {
        assembly {
            _output0 := mload(add(_input, 32))
            _output1 := mload(add(_input, 64))
            _output2 := mload(add(_input, 96))
            _output3 := mload(add(_input, 128))
        }
    }

    function toAddress(
        bytes memory _input
    ) internal pure returns (address _output) {
        _output = abi.decode(_input, (address));
        // assembly {
        //     _output := mload(add(_input, 20))
        // }
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
