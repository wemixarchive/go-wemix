// Copyright 2018-2022 The go-metadium / go-wemix Authors

package miner

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

var (
	ErrNotInitialized = errors.New("not initialized")

	IsMinerFunc                 func() bool
	AmPartnerFunc               func() bool
	IsPartnerFunc               func(string) bool
	AmHubFunc                   func(string) int
	LogBlockFunc                func(int64, common.Hash)
	CalculateRewardsFunc        func(*big.Int, *big.Int, *big.Int, func(common.Address, *big.Int)) (*common.Address, []byte, error)
	VerifyRewardsFunc           func(*big.Int, string) error
	SignBlockFunc               func(height *big.Int, hash common.Hash) (coinbase common.Address, sig []byte, err error)
	VerifyBlockSigFunc          func(height *big.Int, coinbase common.Address, nodeId []byte, hash common.Hash, sig []byte) bool
	RequirePendingTxsFunc       func() bool
	VerifyBlockRewardsFunc      func(height *big.Int) interface{}
	SuggestGasPriceFunc         func() *big.Int
	GetBlockBuildParametersFunc func(height *big.Int) (blockInterval int64, maxBaseFee, gasLimit *big.Int, baseFeeMaxChangeRate, gasTargetPercentage int64, err error)
)

func IsMiner() bool {
	if IsMinerFunc == nil {
		return false
	} else {
		return IsMinerFunc()
	}
}

func IsPartner(id string) bool {
	if IsPartnerFunc == nil {
		return false
	} else {
		return IsPartnerFunc(id)
	}
}

func AmPartner() bool {
	if AmPartnerFunc == nil {
		return false
	} else {
		return AmPartnerFunc()
	}
}

func AmHub(id string) int {
	if AmHubFunc == nil {
		return -1
	} else {
		return AmHubFunc(id)
	}
}

func LogBlock(height int64, hash common.Hash) {
	if LogBlockFunc != nil {
		LogBlockFunc(height, hash)
	}
}

func IsPoW() bool {
	return params.ConsensusMethod == params.ConsensusPoW
}

func CalculateRewards(num, blockReward, fees *big.Int, addBalance func(common.Address, *big.Int)) (*common.Address, []byte, error) {
	if CalculateRewardsFunc == nil {
		return nil, nil, ErrNotInitialized
	} else {
		return CalculateRewardsFunc(num, blockReward, fees, addBalance)
	}
}

func VerifyRewards(num *big.Int, rewards string) error {
	if VerifyRewardsFunc == nil {
		return ErrNotInitialized
	} else {
		return VerifyRewardsFunc(num, rewards)
	}
}

func SignBlock(height *big.Int, hash common.Hash) (coinbase common.Address, sig []byte, err error) {
	if SignBlockFunc == nil {
		err = ErrNotInitialized
	} else {
		coinbase, sig, err = SignBlockFunc(height, hash)
	}
	return
}

func VerifyBlockSig(height *big.Int, coinbase common.Address, nodeId []byte, hash common.Hash, sig []byte) bool {
	if VerifyBlockSigFunc == nil {
		return false
	} else {
		return VerifyBlockSigFunc(height, coinbase, nodeId, hash, sig)
	}
}

func RequirePendingTxs() bool {
	if RequirePendingTxsFunc == nil {
		return false
	} else {
		return RequirePendingTxsFunc()
	}
}

func VerifyBlockRewards(height *big.Int) interface{} {
	if VerifyBlockRewardsFunc == nil {
		return false
	} else {
		return VerifyBlockRewardsFunc(height)
	}
}

func SuggestGasPrice() *big.Int {
	if SuggestGasPriceFunc == nil {
		return big.NewInt(100 * params.GWei)
	} else {
		return SuggestGasPriceFunc()
	}
}

func GetBlockBuildParameters(height *big.Int) (blockInterval int64, maxBaseFee, gasLimit *big.Int, baseFeeMaxChangeRate, gasTargetPercentage int64, err error) {
	if GetBlockBuildParametersFunc == nil {
		// default values
		return 15, big.NewInt(0), big.NewInt(0), 0, 100, ErrNotInitialized
	} else {
		return GetBlockBuildParametersFunc(height)
	}
}

var leadershipSink *chan struct{}

func SubscribeToLeadership(ch *chan struct{}) {
	leadershipSink = ch
}

func UnsubscribeToLeadership() {
	leadershipSink = nil
}

func FeedLeadership() {
	if leadershipSink != nil {
		select {
		case *leadershipSink <- struct{}{}:
		default:
		}
	}
}

type WemixBlockHead struct {
	Height int64
	Hash   common.Hash
}

var blockImportedSink *chan WemixBlockHead

func SubscribeToBlockImported(ch *chan WemixBlockHead) {
	blockImportedSink = ch
}

func UnsubscribeToBlockImported() {
	blockImportedSink = nil
}

func FeedBlockImported(height int64, hash common.Hash) {
	if blockImportedSink != nil {
		select {
		case *blockImportedSink <- WemixBlockHead{Height: height, Hash: hash}:
		default:
			// if full, replace it
			<-*blockImportedSink
			*blockImportedSink <- WemixBlockHead{Height: height, Hash: hash}
		}
	}
}

// EOF
