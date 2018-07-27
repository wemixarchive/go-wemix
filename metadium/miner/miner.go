// Copyright 2018 The go-metadium Authors

package miner

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

var (
	// notification channel when a new transaction arrives
	TxNotifier           = make(chan bool, 1)
	IsMinerFunc          func(int) bool
	CalculateRewardsFunc func(*big.Int, *big.Int, *big.Int, func(common.Address, *big.Int)) ([]byte, error)
	VerifyRewardsFunc    func(*big.Int, string) error
)

func TxNotify() {
	if params.ConsensusMethod != params.ConsensusPoW {
		TxNotifier <- true
	}
}

func IsMiner(height int) bool {
	if IsMinerFunc == nil {
		return false
	} else {
		return IsMinerFunc(height)
	}
}

func IsPoW() bool {
	return params.ConsensusMethod == params.ConsensusPoW
}

func CalculateRewards(num, blockReward, fees *big.Int, addBalance func(common.Address, *big.Int)) ([]byte, error) {
	if CalculateRewardsFunc == nil {
		return nil, fmt.Errorf("Not initialized")
	} else {
		return CalculateRewardsFunc(num, blockReward, fees, addBalance)
	}
}

func VerifyRewards(num *big.Int, rewards string) error {
	if VerifyRewardsFunc == nil {
		return fmt.Errorf("Not initialized")
	} else {
		return VerifyRewardsFunc(num, rewards)
	}
}

// EOF
