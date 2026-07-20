// sync.go

package wemix

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
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
	wemixWorkKey      = "work"
	wemixTokenKey     = "token"
	MiningTokenTTL    = 10 // seconds; must be >= 2 to ensure the syncCheck peer collection timeout (MiningTokenTTL/2) is > 0.
	SyncIdleThreshold = 30 // seconds
)

var (
	// the latest update time of block or work
	latestUpdateTime atomic.Value

	// the latest block info: *core.types.Header
	latestBlock atomic.Value

	// the latest etcd leader ID: uint64
	latestEtcdLeader atomic.Value

	// the latest mining token info: []byte, json'ed WemixToken
	latestMiningToken atomic.Value

	// the latest work info: []byte, json'ed wemixWork
	latestWemixWork atomic.Value

	// governance data at modifiedBlock height
	// cached coinbase -> enode & enode existence check at given modifiedBlock
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

// updates 'latestBlock' upon building or receiving one
func (ma *wemixAdmin) handleNewBlocks() {
	ch := make(chan *types.Header, 128)
	sub, err := ma.cli.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		panic(fmt.Sprintf("cannot subscribe to new block head: %v", err))
	}
	defer sub.Unsubscribe()

	for {
		<-ch

		latestUpdateTime.Store(time.Now())
		if header, err := admin.cli.HeaderByNumber(context.Background(), nil); err == nil {
			latestBlock.Store(header)
			refreshCoinbaseEnodeCache(header.Number)
		}
		if admin != nil {
			admin.update()
		}
	}
}

// updates 'miningPeers' upon receiving WemixMinerStatus from peers
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

// checks if this node is boot node that can / should generate blocks before
// a governance gets set up.
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

// loads saved mining token from 'miningToken'
func loadMiningToken() *WemixToken {
	var (
		lck *WemixToken
		ok  bool
	)
	if lckData := miningToken.Load(); lckData != nil {
		if lck, ok = lckData.(*WemixToken); !ok {
			return nil
		}
	}
	return lck
}

// acquires mining token via etcd, iff (the token doesn't exist or got expired) and (the work matches or doesn't exist).
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
	lck, err := admin.acquireTokenSync(ctx, height, parentHash, MiningTokenTTL)
	if err != nil {
		return false, err
	}
	miningToken.Store(lck)
	return true, nil
}

// logs the latest block & release the mining token
// iff we're still holding the token & the work matches
func releaseMiningToken(height *big.Int, hash, parentHash common.Hash) error {
	if isBootNodeBeforeGenesis() {
		return nil
	}
	lck := loadMiningToken()
	if lck == nil || lck.ttl() < 0 {
		return wemixminer.ErrNotInitialized
	}
	var err error
	for range []int{1, 2} {
		// retry in case it fails to release due to leader changes, etc.
		ctx, cancel := context.WithTimeout(context.Background(),
			admin.etcd.Server.Cfg.ReqTimeout())
		err = lck.releaseTokenSync(ctx, height, hash, parentHash)
		cancel()
		if err == nil {
			break
		}
	}

	// invalidate the saved token
	miningToken.Store(&WemixToken{})
	return err
}

// checks the cache to see if we're holding mining token
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

// finds a block that the majority of the miners have
// considers the latest blocks only
func findConsensusBlock(states []*wemixapi.WemixMinerStatus) (height *big.Int, hash common.Hash) {
	type anon struct {
		height *big.Int
		hash   *common.Hash
		count  int
	}

	n := len(states)/2 + 1
	m := map[string]*anon{}
	for _, state := range states {
		if state.LatestBlockHeight == nil {
			continue
		}
		key := string(append(state.LatestBlockHeight.Bytes(), state.LatestBlockHash.Bytes()...))
		if f, ok := m[key]; !ok {
			m[key] = &anon{
				height: state.LatestBlockHeight,
				hash:   &state.LatestBlockHash,
				count:  1,
			}
		} else {
			f.count += 1
		}
	}
	for _, f := range m {
		if f.count >= n {
			return f.height, *f.hash
		}
	}
	return nil, common.Hash{}
}

// handles inconsistencies if any
//  0. updates of 'latestBlock' of 'latestWemixWork' hasn't occurred over 30 seconds
//     -> i.e. the system is likely stuck
//  1. the token is invalid, i.e. present but json.Unmarshal fails
//     -> remove
//  2. our latest block is ahead of the recorded 'work'
//     -> if recorded 'work' exists and valid,
//     then revert back to the 'work' block using 'debug.setHead()'
//  3. the recorded 'work' doesn't exist, i.e. no mining peers has it as their
//     latest block, or
//     the 'work' is invalid, i.e. present but json.Unmarshal fails
//     -> find the consensus block, a block that the majority of the miners have
//     as the latest block, sets it to the 'work'
func syncCheck() error {
	if admin == nil || !admin.amPartner() || admin.self == nil || !admin.etcdIsRunning() {
		return nil
	}

	secondsSinceUpdate := func() int {
		if v := latestUpdateTime.Load(); v == nil {
			latestUpdateTime.Store(time.Now())
			return -1
		} else if t, ok := v.(time.Time); ok {
			return int(time.Since(t).Seconds())
		} else {
			panic("invalid update time")
		}
	}

	if t := secondsSinceUpdate(); t <= SyncIdleThreshold {
		return nil
	} else {
		log.Debug("sync check", "last update in seconds", t)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// check token string
	if tokenData, err := admin.etcdGet(wemixTokenKey); err == nil {
		var token = &WemixToken{}
		if err = json.Unmarshal([]byte(tokenData), token); err != nil {
			// invalid token string
			err = admin.etcdDelete(wemixTokenKey)
			log.Error("sync check: reset the invalid token", "token", tokenData, "error", err)
		}
	}

	header, err := admin.cli.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Error("sync check: failed to get the latest block", "error", err)
		return err
	}
	num := new(big.Int).Add(header.Number, common.Big1)
	token, err := admin.acquireToken(ctx, num, MiningTokenTTL)
	if err != nil {
		return err
	}
	defer func() {
		if token != nil {
			token.release(ctx)
		}
	}()

	// check update time again
	if t := secondsSinceUpdate(); t <= SyncIdleThreshold {
		log.Debug("sync check", "new update occurred", t)
		return nil
	}

	// need to refresh the latest block & current work
	header, err = admin.cli.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Error("sync check: failed to get the latest block", "error", err)
		return err
	}

	var (
		workData string
		work     = &wemixWork{}
	)
	if workData, err = admin.etcdGet(wemixWorkKey); err != nil {
		return err
	} else {
		if err = json.Unmarshal([]byte(workData), work); err != nil {
			// invalid work data
			log.Error("sync check: ignoring invalid work", "work", workData)
			work = nil
		}
	}

	// if we're in sync, nothing we can do
	if work != nil && work.Height == header.Number.Int64() && work.Hash == header.Hash() {
		log.Debug("sync check: in sync", "height", work.Height, "hash", work.Hash)
		return err
	}

	// we're ahead of 'work'
	if work != nil && work.Height < header.Number.Int64() {
		// It's too dangerous to go back to the recorded work.
		// If the network, except this node, is in sync, this node will eventually
		// follow the longest chain.
		log.Error("sync check: ahead of work, aborting", "our-height", header.Number, "work-height", work.Height)
		return nil
	}

	// collects mining peers' latest blocks
	// getMiners timeout must be shorter than MiningTokenTTL to ensure etcdPut and subsequent operations
	// complete while the token is still held. if the token expires first, another node may advance
	// work before etcdPut, causing a stale overwrite. MiningTokenTTL/2 leaves headroom for those operations.
	nodes := admin.getNodes()
	states := getMiners("", MiningTokenTTL/2)
	// The original guard `len(nodes)==0 && len(nodes)!=len(states)` was a dead
	// branch in any healthy environment (nodes>=5). Recovered intent: if the
	// governance cache is empty we cannot reason about quorum, so abort. The
	// strict `nodes!=states` comparison was dropped because it false-positives
	// on partial peer outages.
	if len(nodes) == 0 {
		log.Error("sync check: no governance nodes loaded, aborting", "states", len(states))
		return nil
	}

	// 'work' is ahead of us
	if work != nil && work.Height > header.Number.Int64() {
		// Detect catch-up via a single peer-state echo to silence the
		// rebuild-path log noise that a reachability-only gate previously
		// emitted every cycle (HeaderByHash(work.Hash) is guaranteed to
		// return NotFound while we are still importing).
		//
		// Counting matches to require quorum was considered but does not
		// fundamentally resist poison: a single connected malicious peer
		// can echo any value, so any quorum size short of all-honest is
		// still influenced by one attacker message. Given that limit, the
		// policy here is side-effect minimization — quiet the legitimate
		// catch-up case and leave poison recovery to the downstream
		// consensus rebuild plus its reachability and regression guards.
		// Operator intervention remains the backstop for the byzantine
		// case.
		exist := false
		for _, state := range states {
			if state.LatestBlockHeight == nil {
				continue
			}
			if state.LatestBlockHeight.Int64() == work.Height &&
				work.Hash == state.LatestBlockHash {
				exist = true
				break
			}
		}
		if exist {
			log.Debug("sync check: catching up to cluster head",
				"local", header.Number, "work-height", work.Height,
				"work-hash", work.Hash, "states", len(states))
			return nil
		}
		// No peer state echoes the stored value: either the cluster moved
		// past it (legitimate stale) or it was poisoned. Both route through
		// the findConsensusBlock rebuild below; the downstream reachability
		// check and regression guard reject unsafe writes.
		log.Warn("sync check: work-key not echoed by any peer, rebuilding via consensus",
			"work-height", work.Height, "work-hash", work.Hash,
			"states", len(states))
	}

	// work record doesn't exist or recorded block doesn't exist
	consensusHeight, consensusHash := findConsensusBlock(states)
	if consensusHeight == nil {
		// no consensus
		if work == nil {
			log.Error("sync check: no consensus block found, aborting", "work", "invalid")
		} else {
			log.Error("sync check: no consensus block found, aborting", "work-height", work.Height, "work-hash", work.Hash)
		}
		return nil
	}
	// Before persisting to etcd, require that consensusHash is reachable from
	// the local chain. Writing an unreachable hash into wemixWorkKey would
	// permanently fail every validator's acquireTokenSync (ErrInvalidWork), so
	// reject unknown hashes outright.
	peerHeader, herr := admin.cli.HeaderByHash(ctx, consensusHash)
	if herr != nil || peerHeader == nil {
		log.Error("sync check: consensus block not reachable from local chain, aborting",
			"height", consensusHeight, "hash", consensusHash, "error", herr)
		return nil
	}
	// Verify that consensusHash and consensusHeight are a consistent pair.
	// A real hash echoed under a fake height (quorum-tampered
	// findConsensusBlock output) would otherwise pass the reachability check
	// and poison wemixWorkKey with an internally inconsistent value.
	if peerHeader.Number.Cmp(consensusHeight) != 0 {
		log.Error("sync check: consensus hash/height pair mismatch, suspected tampering, aborting",
			"claimed-height", consensusHeight, "actual-height", peerHeader.Number,
			"hash", consensusHash)
		return nil
	}
	// consensusHeight must not regress the local head. A consensus tuple
	// pointing below header.Number indicates poisoned states echoing an old
	// (still-reachable) hash; persisting it would push every validator's
	// mining target backward and stall block production permanently.
	if consensusHeight.Cmp(header.Number) < 0 {
		log.Error("sync check: consensus height regresses local head, aborting",
			"local", header.Number, "consensus", consensusHeight, "consensus-hash", consensusHash)
		return nil
	}
	newWork := &wemixWork{
		Height: consensusHeight.Int64(),
		Hash:   consensusHash,
	}
	if newWorkData, err := json.Marshal(newWork); err != nil {
		panic("failed to marshal work data")
	} else {
		admin.etcdPut(wemixWorkKey, string(newWorkData))
	}
	log.Error("sync check: found consensus block, setting work", "height", consensusHeight, "hash", consensusHash, "error", err)
	return err
}

// EOF
