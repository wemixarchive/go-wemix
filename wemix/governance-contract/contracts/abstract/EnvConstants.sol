// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EnvConstants {
    bytes32 public constant BLOCKS_PER_NAME = keccak256("blocksPer");

    bytes32 public constant BALLOT_DURATION_MIN_NAME =
        keccak256("ballotDurationMin");

    bytes32 public constant BALLOT_DURATION_MAX_NAME =
        keccak256("ballotDurationMax");

    bytes32 public constant STAKING_MIN_NAME = keccak256("stakingMin");

    bytes32 public constant STAKING_MAX_NAME = keccak256("stakingMax");

    bytes32 public constant MAX_IDLE_BLOCK_INTERVAL_NAME =
        keccak256("MaxIdleBlockInterval");

    //=======NXTMeta========/

    bytes32 public constant BALLOT_DURATION_MIN_MAX_NAME =
        keccak256("ballotDurationMinMax");
    bytes32 public constant STAKING_MIN_MAX_NAME = keccak256("stakingMinMax");

    bytes32 public constant BLOCK_CREATION_TIME_NAME =
        keccak256("blockCreationTime");
    bytes32 public constant BLOCK_REWARD_AMOUNT_NAME =
        keccak256("blockRewardAmount");
    // unit = gwei
    bytes32 public constant MAX_PRIORITY_FEE_PER_GAS_NAME =
        keccak256("maxPriorityFeePerGas");

    bytes32 public constant BLOCK_REWARD_DISTRIBUTION_METHOD_NAME =
        keccak256("blockRewardDistribution");
    bytes32 public constant BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER_NAME =
        keccak256("blockRewardDistributionBlockProducer");
    bytes32 public constant BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD_NAME =
        keccak256("blockRewardDistributionStakingReward");
    bytes32 public constant BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM_NAME =
        keccak256("blockRewardDistributionEcosystem");
    bytes32 public constant BLOCK_REWARD_DISTRIBUTION_MAINTENANCE_NAME =
        keccak256("blockRewardDistributionMaintenance");

    bytes32 public constant GASLIMIT_AND_BASE_FEE_NAME =
        keccak256("gasLimitAndBaseFee");
    bytes32 public constant BLOCK_GASLIMIT_NAME = keccak256("blockGasLimit");
    bytes32 public constant BASE_FEE_MAX_CHANGE_RATE_NAME =
        keccak256("baseFeeMaxChangeRate");
    bytes32 public constant GAS_TARGET_PERCENTAGE_NAME =
        keccak256("gasTargetPercentage");

    bytes32 public constant MAX_BASE_FEE_NAME = keccak256("maxBaseFee");

    uint256 public constant DENOMINATOR = 10000;
}
