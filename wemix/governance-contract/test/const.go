package test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/wemix/governance-contract/common/bn"
)

var (
	LOCK_AMOUNT *big.Int = bn.ToWei18(1500000)
)

var EnvConstants = struct {
	BLOCKS_PER                               env
	BALLOT_DURATION_MIN                      env
	BALLOT_DURATION_MAX                      env
	STAKING_MIN                              env
	STAKING_MAX                              env
	MAX_IDLE_BLOCK_INTERVAL                  env
	BALLOT_DURATION_MIN_MAX                  env
	STAKING_MIN_MAX                          env
	BLOCK_CREATION_TIME                      env
	BLOCK_REWARD_AMOUNT                      env
	MAX_PRIORITY_FEE_PER_GAS                 env
	BLOCK_REWARD_DISTRIBUTION_METHOD         env
	BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER env
	BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD env
	BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM      env
	BLOCK_REWARD_DISTRIBUTION_MAINTANANCE    env
	GASLIMIT_AND_BASE_FEE                    env
	BLOCK_GASLIMIT                           env
	BASE_FEE_MAX_CHANGE_RATE                 env
	GAS_TARGET_PERCENTAGE                    env
	MAX_BASE_FEE                             env
}{
	BLOCKS_PER:                               newEnv("blocksPer", 1),
	BALLOT_DURATION_MIN:                      newEnv("ballotDurationMin", 86400),
	BALLOT_DURATION_MAX:                      newEnv("ballotDurationMax", 604800),
	STAKING_MIN:                              newEnv("stakingMin", bn.ToWei18(1500000)),
	STAKING_MAX:                              newEnv("stakingMax", bn.ToWei18(1500000)),
	MAX_IDLE_BLOCK_INTERVAL:                  newEnv("MaxIdleBlockInterval", 5),
	BALLOT_DURATION_MIN_MAX:                  newEnv("ballotDurationMinMax", nil),
	STAKING_MIN_MAX:                          newEnv("stakingMinMax", nil),
	BLOCK_CREATION_TIME:                      newEnv("blockCreationTime", 1000),
	BLOCK_REWARD_AMOUNT:                      newEnv("blockRewardAmount", bn.ToWei18(1)),
	MAX_PRIORITY_FEE_PER_GAS:                 newEnv("maxPriorityFeePerGas", bn.Mul(100, params.GWei)),
	BLOCK_REWARD_DISTRIBUTION_METHOD:         newEnv("blockRewardDistribution", nil),
	BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER: newEnv("blockRewardDistributionBlockProducer", 4000),
	BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD: newEnv("blockRewardDistributionStakingReward", 1000),
	BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM:      newEnv("blockRewardDistributionEcosystem", 2500),
	BLOCK_REWARD_DISTRIBUTION_MAINTANANCE:    newEnv("blockRewardDistributionMaintenance", 2500),
	GASLIMIT_AND_BASE_FEE:                    newEnv("gasLimitAndBaseFee", nil),
	BLOCK_GASLIMIT:                           newEnv("blockGasLimit", 1050000000),
	BASE_FEE_MAX_CHANGE_RATE:                 newEnv("baseFeeMaxChangeRate", 46),
	GAS_TARGET_PERCENTAGE:                    newEnv("gasTargetPercentage", 30),
	MAX_BASE_FEE:                             newEnv("maxBaseFee", bn.Mul(5000, params.GWei)),
}

type env struct {
	Name  [32]byte
	Value *big.Int
}

func newEnv(name string, value interface{}) env {
	if value == nil {
		value = bn.New(0)
	}
	return env{
		Name:  crypto.Keccak256Hash([]byte(name)),
		Value: bn.New(value),
	}
}

func makeEnvParams(envs ...env) (names [][32]byte, values []*big.Int) {
	length := len(envs)
	names = make([][32]byte, length)
	values = make([]*big.Int, length)
	for i, e := range envs {
		names[i] = e.Name
		values[i] = e.Value
	}
	return
}

type MemberInfo struct {
	Staker     common.Address `json:"staker"`
	Voter      common.Address `json:"voter"`
	Reward     common.Address `json:"reward"`
	Name       []byte         `json:"name"`
	Enode      []byte         `json:"enode"`
	Ip         []byte         `json:"ip"`
	Port       *big.Int       `json:"port"`
	LockAmount *big.Int       `json:"lockAmount"`
	Memo       []byte         `json:"memo"`
	Duration   *big.Int       `json:"duration"`
}
