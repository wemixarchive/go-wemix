// Copyright 2018 The go-metadium Authors

package miner

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

var (
	IsMinerFunc          func(int) bool
	IsPartnerFunc        func(string) bool
	CalculateRewardsFunc func(*big.Int, *big.Int, *big.Int, func(common.Address, *big.Int)) ([]byte, error)
	VerifyRewardsFunc    func(*big.Int, string) error
	RequirePendingTxsFunc  func() bool
)

func IsMiner(height int) bool {
	if IsMinerFunc == nil {
		return false
	} else {
		return IsMinerFunc(height)
	}
}

func IsPartner(id string) bool {
	if IsPartnerFunc == nil {
		return false
	} else {
		return IsPartnerFunc(id)
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

func RequirePendingTxs() bool {
	if RequirePendingTxsFunc == nil {
		return false
	} else {
		return RequirePendingTxsFunc()
	}
}

// EOF
