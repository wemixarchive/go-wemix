// tx_params.go

package metclient

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// account nonce cache map. Uses a global lock for now. If necessary,
// need to use sync.Map instead.
// Also we might have to maintain window ala tcp window to accommodate nonces
// used in failed transactions.
var (
	txParamsCache = &sync.Map{}
)

func GetOpportunisticTxParams(ctx context.Context, cli *ethclient.Client, addr common.Address, refresh bool, incNonce bool) (chainId, gasPrice, nonce *big.Int, err error) {
	n, n_ok := txParamsCache.Load(addr)
	cid, cid_ok := txParamsCache.Load("chain-id")
	gp, gp_ok := txParamsCache.Load("gas-price")

	if !refresh && n_ok && cid_ok && gp_ok {
		// cache's good
		chainId = new(big.Int).Set(cid.(*big.Int))
		gasPrice = new(big.Int).Set(gp.(*big.Int))
		nonce = new(big.Int).Set(n.(*big.Int))
		if incNonce {
			txParamsCache.Store(addr, new(big.Int).Add(nonce, common.Big1))
		}
		return
	}

	if !refresh && cid_ok {
		chainId = new(big.Int).Set(cid.(*big.Int))
	} else {
		cid, err = cli.NetworkID(ctx)
		if err != nil {
			return
		}
		txParamsCache.Store("chain-id", cid)
		chainId = new(big.Int).Set(cid.(*big.Int))
	}
	if !refresh && gp_ok {
		gasPrice = new(big.Int).Set(gp.(*big.Int))
	} else {
		gp, err = cli.SuggestGasPrice(ctx)
		if err != nil {
			return
		}
		txParamsCache.Store("gas-price", gp)
		gasPrice = new(big.Int).Set(gp.(*big.Int))
	}

	var n1, n2 uint64
	n1, err = cli.PendingNonceAt(ctx, addr)
	if err != nil {
		return
	}
	n2, err = cli.NonceAt(ctx, addr, nil)
	if err != nil {
		return
	}
	if n1 < n2 {
		n1 = n2
	}

	nonce = big.NewInt(int64(n1))
	if !incNonce {
		txParamsCache.Store(addr, nonce)
	} else {
		txParamsCache.Store(addr, new(big.Int).Add(nonce, common.Big1))
	}

	return
}

// EOF
