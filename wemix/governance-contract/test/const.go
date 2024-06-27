package test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

func towei(x int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(x), big.NewInt(params.Ether))
}

func toGwei(x int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(x), big.NewInt(params.GWei))
}

var (
	LOCK_AMOUNT *big.Int = towei(1500000)
)

var EnvConstants = struct {
	BLOCKS_PER                               env
	BALLOT_DURATION_MIN                      env
	BALLOT_DURATION_MAX                      env
	STAKING_MIN                              env
	STAKING_MAX                              env
	MAX_IDLE_BLOCK_INTERVAL                  env
	BLOCK_CREATION_TIME                      env
	BLOCK_REWARD_AMOUNT                      env
	MAX_PRIORITY_FEE_PER_GAS                 env
	BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER env
	BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD env
	BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM      env
	BLOCK_REWARD_DISTRIBUTION_MAINTENANCE    env
	MAX_BASE_FEE                             env
	BLOCK_GASLIMIT                           env
	BASE_FEE_MAX_CHANGE_RATE                 env
	GAS_TARGET_PERCENTAGE                    env
}{
	BLOCKS_PER:                               newEnvInt64("blocksPer", 1),
	BALLOT_DURATION_MIN:                      newEnvInt64("ballotDurationMin", 86400),
	BALLOT_DURATION_MAX:                      newEnvInt64("ballotDurationMax", 604800),
	STAKING_MIN:                              newEnvBig("stakingMin", towei(1500000)),
	STAKING_MAX:                              newEnvBig("stakingMax", towei(1500000)),
	MAX_IDLE_BLOCK_INTERVAL:                  newEnvInt64("MaxIdleBlockInterval", 5),
	BLOCK_CREATION_TIME:                      newEnvInt64("blockCreationTime", 1000),
	BLOCK_REWARD_AMOUNT:                      newEnvBig("blockRewardAmount", towei(1)),
	MAX_PRIORITY_FEE_PER_GAS:                 newEnvBig("maxPriorityFeePerGas", toGwei(100)),
	BLOCK_REWARD_DISTRIBUTION_BLOCK_PRODUCER: newEnvInt64("blockRewardDistributionBlockProducer", 4000),
	BLOCK_REWARD_DISTRIBUTION_STAKING_REWARD: newEnvInt64("blockRewardDistributionStakingReward", 1000),
	BLOCK_REWARD_DISTRIBUTION_ECOSYSTEM:      newEnvInt64("blockRewardDistributionEcosystem", 2500),
	BLOCK_REWARD_DISTRIBUTION_MAINTENANCE:    newEnvInt64("blockRewardDistributionMaintenance", 2500),
	MAX_BASE_FEE:                             newEnvBig("maxBaseFee", toGwei(5000)),
	BLOCK_GASLIMIT:                           newEnvInt64("blockGasLimit", 1050000000),
	BASE_FEE_MAX_CHANGE_RATE:                 newEnvInt64("baseFeeMaxChangeRate", 46),
	GAS_TARGET_PERCENTAGE:                    newEnvInt64("gasTargetPercentage", 30),
}

type env struct {
	Name  [32]byte
	Value *big.Int
}

func newEnvInt64(name string, value int64) env {
	return env{
		Name:  crypto.Keccak256Hash([]byte(name)),
		Value: big.NewInt(value),
	}
}

func newEnvBig(name string, value *big.Int) env {
	if value == nil {
		value = common.Big0
	}
	return env{
		Name:  crypto.Keccak256Hash([]byte(name)),
		Value: value,
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

var EnvTypes = struct {
	Invalid *big.Int
	Int     *big.Int
	Uint    *big.Int
	Address *big.Int
	Bytes32 *big.Int
	Bytes   *big.Int
	String  *big.Int
}{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5), big.NewInt(6)}

var BallotStates = struct {
	Invalid    *big.Int
	Ready      *big.Int
	InProgress *big.Int
	Accepted   *big.Int
	Rejected   *big.Int
}{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4)}
