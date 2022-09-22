// sync.go

package wemix

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	wemixapi "github.com/ethereum/go-ethereum/wemix/api"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

// cached governance data to derive miner's enode
type coinbaseEnodeEntry struct {
	modifiedBlock  *big.Int
	nodes          []*wemixNode
	coinbase2enode map[string][]byte // string(common.Address[:]) => []byte
	enode2index    map[string]int    // string([]byte) => int
}

const (
	wemixWorkKey  = "work"
	wemixLockKey  = "lock"
	MiningLockTTL = 10 // seconds
)

var (
	syncLock = &sync.Mutex{}

	// the latest block info: *core.types.Header
	latestBlock atomic.Value

	// the latest etcd leader ID: uint64
	latestEtcdLeader atomic.Value

	// the latest mining token info: []byte, json'ed WemixEtcdLock
	latestMiningToken atomic.Value

	// the latest work info: []byte, json'ed wemixWork
	latestWemixWork atomic.Value

	// governance data at modifiedBlock height
	// cached coinbase -> enode & enode existance check at given modifiedBlock
	// sync.Map[int]*cointbaseEnodeEntry
	coinbaseEnodeCache = &sync.Map{}

	// lru cache: block height => enode
	height2enode = lru.NewLruCache(10000, true)

	// cached mining peer status
	// sync.Map[string]*wemixapi.WemixMinerStatus
	miningPeers = &sync.Map{}

	// mining token
	miningToken atomic.Value
)

func (ma *wemixAdmin) handleNewBlocks() {
	ch := make(chan *types.Header, 128)
	sub, err := ma.cli.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		panic(fmt.Sprintf("cannot subscribe to new block head: %v", err))
	}
	defer sub.Unsubscribe()

	for {
		_ = <-ch
		if header, err := admin.cli.HeaderByNumber(context.Background(), nil); err == nil {
			latestBlock.Store(header)
			refreshCoinbaseEnodeCache(header.Number)
		}
		admin.update()
	}
}

func handleMinerStatusUpdate() {
	ch := make(chan *wemixapi.WemixMinerStatus, 128)
	sub := wemixapi.SubscribeToMinerStatus(ch)
	defer func() {
		sub.Unsubscribe()
		close(ch)
	}()
	for {
		status := <-ch
		miningPeers.Store(status.NodeName, status.Clone())
	}
}

func isBootNodeBeforeGenesis() bool {
	if params.ConsensusMethod == params.ConsensusPoW {
		return true
	} else if params.ConsensusMethod == params.ConsensusETCD {
		return false
	} else if params.ConsensusMethod == params.ConsensusPoA {
		if admin == nil {
			return false
		} else if admin.self == nil || len(admin.nodes) <= 0 {
			if admin.nodeInfo != nil && admin.nodeInfo.ID == admin.bootNodeId {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func loadMiningToken() *WemixEtcdLock {
	var (
		lck *WemixEtcdLock
		ok  bool
	)
	if lckData := miningToken.Load(); lckData != nil {
		if lck, ok = lckData.(*WemixEtcdLock); !ok {
			return nil
		}
	}
	return lck
}

func acquireMiningToken(height *big.Int, parentHash common.Hash) (bool, error) {
	if isBootNodeBeforeGenesis() {
		return true, nil
	}
	if admin == nil || !admin.etcdIsRunning() {
		return false, ErrNotRunning
	}
	ctx, cancel := context.WithTimeout(context.Background(),
		admin.etcd.Server.Cfg.ReqTimeout())
	defer cancel()
	lck, err := admin.acquireLockSync(ctx, height, parentHash, MiningLockTTL)
	if err != nil {
		return false, err
	}
	miningToken.Store(lck)
	return true, nil
}

// log the latest block & release the mining token
func releaseMiningToken(height *big.Int, hash, parentHash common.Hash) error {
	if isBootNodeBeforeGenesis() {
		return nil
	}
	lck := loadMiningToken()
	if lck == nil || lck.ttl() < 0 {
		return wemixminer.ErrNotInitialized
	}
	ctx, cancel := context.WithTimeout(context.Background(),
		admin.etcd.Server.Cfg.ReqTimeout())
	defer cancel()
	err := lck.releaseLockSync(ctx, height, hash, parentHash)

	// invalidate the saved token
	miningToken.Store(&WemixEtcdLock{})
	return err
}

func hasMiningToken() bool {
	if isBootNodeBeforeGenesis() {
		return true
	}
	lck := loadMiningToken()
	if lck == nil || lck.ttl() < 0 {
		return false
	}
	return true
}

// EOF
