// Copyright 2026 The go-wemix Authors

package rewards

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

// db is taken instead of the block's statedb so rewards never see the block's own transactions.
var CalculateRewardsFunc func(chain consensus.ChainHeaderReader, header *types.Header, db state.Database, addBalance func(common.Address, *big.Int)) ([]byte, error)

func CalculateRewards(chain consensus.ChainHeaderReader, header *types.Header, db state.Database, addBalance func(common.Address, *big.Int)) ([]byte, error) {
	if CalculateRewardsFunc == nil {
		return nil, wemixminer.ErrNotInitialized
	}
	return CalculateRewardsFunc(chain, header, db, addBalance)
}
