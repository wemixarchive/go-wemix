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
	_lock     = &sync.Mutex{}
	_chainId  *big.Int
	_gasPrice *big.Int
	_nonces   = map[common.Address]*big.Int{}
)

func GetOpportunisticTxParams(ctx context.Context, cli *ethclient.Client, addr common.Address, refresh bool, incNonce bool) (chainId, gasPrice, nonce *big.Int, err error) {
	_lock.Lock()
	defer _lock.Unlock()

	if _nonce, ok := _nonces[addr]; refresh || _chainId == nil || _gasPrice == nil || !ok {
		// pass
	} else {
		// cache's good
		chainId = big.NewInt(_chainId.Int64())
		gasPrice = big.NewInt(_gasPrice.Int64())
		nonce = big.NewInt(_nonce.Int64())
		if incNonce {
			_nonce.Add(_nonce, common.Big1)
		}
		return
	}

	if refresh || _chainId == nil {
		var cid *big.Int
		cid, err = cli.NetworkID(ctx)
		if err != nil {
			return
		}
		_chainId = big.NewInt(cid.Int64())
		chainId = big.NewInt(cid.Int64())
	} else {
		chainId = big.NewInt(_chainId.Int64())
	}
	if refresh || _gasPrice == nil {
		var gp *big.Int
		gp, err = cli.SuggestGasPrice(ctx)
		if err != nil {
			return
		}
		_gasPrice = big.NewInt(gp.Int64())
		gasPrice = big.NewInt(gp.Int64())
	} else {
		gasPrice = big.NewInt(_gasPrice.Int64())
	}

	var _n1, _n2 uint64
	_n1, err = cli.PendingNonceAt(ctx, addr)
	if err != nil {
		return
	}
	_n2, err = cli.NonceAt(ctx, addr, nil)
	if err != nil {
		return
	}
	if _n1 < _n2 {
		_n1 = _n2
	}

	_nonces[addr] = big.NewInt(int64(_n1))
	nonce = big.NewInt(_nonces[addr].Int64())
	if incNonce {
		_nonces[addr].Add(_nonces[addr], common.Big1)
	}

	return
}

// EOF
