// miner_limit.go

package wemix

import (
	"bytes"
	"context"
	"encoding/hex"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	wemixapi "github.com/ethereum/go-ethereum/wemix/api"
	"github.com/ethereum/go-ethereum/wemix/metclient"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

// collect mining peers states
func (ma *wemixAdmin) collectMinerStates(height *big.Int) []*wemixapi.WemixMinerStatus {
	var (
		ctx context.Context
		gov *metclient.RemoteContract
		err error
	)
	ctx = context.Background()
	if _, gov, _, _, err = ma.getRegGovEnvContracts(ctx, height); err != nil {
		return nil
	}
	e, err := getCoinbaseEnodeCache(ctx, height, gov)
	if err != nil {
		return nil
	}

	var states []*wemixapi.WemixMinerStatus
	for _, n := range e.nodes {
		if s, ok := miningPeers.Load(n.Name); ok {
			if status, ok := s.(*wemixapi.WemixMinerStatus); ok {
				states = append(states, status)
			}
		}
	}
	return states
}

// get governance nodes at modifiedBlock at height
func getCoinbaseEnodeCache(ctx context.Context, height *big.Int, gov *metclient.RemoteContract) (*coinbaseEnodeEntry, error) {
	var modifiedBlock *big.Int
	if err := metclient.CallContract(ctx, gov, "modifiedBlock", nil, &modifiedBlock, height); err != nil {
		return nil, err
	} else if modifiedBlock.Int64() == 0 {
		return nil, wemixminer.ErrNotInitialized
	}
	// if found in cache, use it
	if e, ok := coinbaseEnodeCache.Load(modifiedBlock.Int64()); ok {
		return e.(*coinbaseEnodeEntry), nil
	}
	// otherwise, load it from the governance
	var (
		count, port     *big.Int
		addr            common.Address
		name, enode, ip []byte
		output          = []interface{}{&name, &enode, &ip, &port}
		e               = &coinbaseEnodeEntry{
			modifiedBlock:  modifiedBlock,
			coinbase2enode: map[string][]byte{},
			enode2index:    map[string]int{},
		}
	)
	if err := metclient.CallContract(ctx, gov, "getNodeLength", nil, &count, height); err != nil {
		return nil, err
	}
	for i := int64(1); i <= count.Int64(); i++ {
		ix := big.NewInt(i)
		if err := metclient.CallContract(ctx, gov, "getReward", &ix, &addr, height); err != nil {
			return nil, err
		}
		if err := metclient.CallContract(ctx, gov, "getNode", &ix, &output, height); err != nil {
			return nil, err
		}
		idv4, _ := toIdv4(hex.EncodeToString(enode))
		e.nodes = append(e.nodes, &wemixNode{
			Name:  string(name),
			Enode: string(enode), // note that this is not in hex unlike wemixAdmin
			Id:    idv4,
			Addr:  addr,
		})
		e.coinbase2enode[string(addr[:])] = enode
		e.enode2index[string(enode)] = int(i) // 1-based, not 0-based
	}
	coinbaseEnodeCache.Store(modifiedBlock.Int64(), e)
	return e, nil
}

// returns coinbase's enode if exists in governance at given height - 1
func coinbaseExists(ctx context.Context, height *big.Int, gov *metclient.RemoteContract, coinbase *common.Address) ([]byte, error) {
	e, err := getCoinbaseEnodeCache(ctx, new(big.Int).Sub(height, common.Big1), gov)
	if err != nil {
		return nil, err
	}
	enode, ok := e.coinbase2enode[string(coinbase[:])]
	if !ok {
		return nil, nil
	}
	return enode, nil
}

// returns true if enode exists in governance at given height-1
func enodeExists(ctx context.Context, height *big.Int, gov *metclient.RemoteContract, enode []byte) (bool, error) {
	e, err := getCoinbaseEnodeCache(ctx, new(big.Int).Sub(height, common.Big1), gov)
	if err != nil {
		return false, err
	}
	ix, ok := e.enode2index[string(enode)]
	if !ok {
		return false, nil
	}
	return ix >= 1, nil
}

// returns wemix nodes at given height
func getNodesAt(height *big.Int) ([]*wemixNode, error) {
	ctx := context.Background()
	if _, gov, _, _, err := admin.getRegGovEnvContracts(ctx, height); err != nil {
		return nil, wemixminer.ErrNotInitialized
	} else if e, err := getCoinbaseEnodeCache(ctx, height, gov); err != nil {
		return nil, err
	} else {
		return e.nodes, nil
	}
}

// return block's miner node id
func getBlockMiner(ctx context.Context, cli *ethclient.Client, entry *coinbaseEnodeEntry, height *big.Int) ([]byte, error) {
	// if already cached, use it
	if enode := height2enode.Get(height.Uint64()); enode != nil {
		return enode.([]byte), nil
	}
	block, err := cli.HeaderByNumber(ctx, height)
	if err != nil {
		return nil, err
	}
	if len(block.MinerNodeId) == 0 {
		enode, ok := entry.coinbase2enode[string(block.Coinbase[:])]
		if !ok {
			return nil, nil
		}
		height2enode.Put(height.Uint64(), enode)
		return enode, nil
	} else {
		if _, ok := entry.enode2index[string(block.MinerNodeId)]; !ok {
			return nil, nil
		}
		height2enode.Put(height.Uint64(), block.MinerNodeId)
		return block.MinerNodeId, nil
	}
}

// A miner can mine one or less block within (member count / 2 + 1) blocks.
// It's reset when governance gets updated, i.e. search doesn't go back
// beyond modifiedBlock + 1.
// Not enforced if member count <= 2.
func (ma *wemixAdmin) verifyMinerLimit(ctx context.Context, height *big.Int, gov *metclient.RemoteContract, coinbase *common.Address, enode []byte) (bool, error) {
	// parent block number
	prev := new(big.Int).Sub(height, common.Big1)
	e, err := getCoinbaseEnodeCache(ctx, prev, gov)
	if err != nil {
		return false, err
	}
	// if count <= 2, not enforced
	if len(e.nodes) <= 2 {
		return true, nil
	}
	// if enode is not given, derive it from the coinbase
	if len(enode) == 0 {
		enode2, ok := e.coinbase2enode[string(coinbase[:])]
		if !ok {
			return false, nil
		}
		enode = enode2
	}
	var miners [][]byte
	// the enode should not appear within the last (member count / 2) blocks
	limit := len(e.nodes) / 2
	if limit > int(height.Int64()-e.modifiedBlock.Int64()-1) {
		limit = int(height.Int64() - e.modifiedBlock.Int64() - 1)
	}
	for h := new(big.Int).Set(prev); limit > 0; h, limit = h.Sub(h, common.Big1), limit-1 {
		blockMinerEnode, err := getBlockMiner(ctx, ma.cli, e, h)
		if err != nil {
			return false, err
		}
		miners = append(miners, blockMinerEnode[:])
		if bytes.Equal(enode[:], blockMinerEnode[:]) {
			return false, nil
		}
	}
	return true, nil
}

// check if self is eligible to mine height block at height-1
func (ma *wemixAdmin) isEligibleMiner(height *big.Int) (bool, error) {
	var (
		ctx   context.Context
		enode []byte
		gov   *metclient.RemoteContract
		err   error
	)
	ctx = context.Background()
	prev := new(big.Int).Sub(height, common.Big1)
	enode, err = hex.DecodeString(ma.self.Enode)
	if err != nil {
		return false, wemixminer.ErrNotInitialized
	}
	if _, gov, _, _, err = ma.getRegGovEnvContracts(ctx, prev); err != nil {
		return false, wemixminer.ErrNotInitialized
	}
	e, err := getCoinbaseEnodeCache(ctx, prev, gov)
	if err != nil {
		return false, err
	}
	// if count <= 2, not enforced
	if len(e.nodes) <= 2 {
		return true, nil
	}
	// the enode should not appear within the last (member count / 2) blocks
	limit := len(e.nodes) / 2
	if limit > int(height.Int64()-e.modifiedBlock.Int64()-1) {
		limit = int(height.Int64() - e.modifiedBlock.Int64() - 1)
	}
	for h := new(big.Int).Set(prev); limit > 0; h, limit = h.Sub(h, common.Big1), limit-1 {
		blockMinerEnode, err := getBlockMiner(ctx, ma.cli, e, h)
		if err != nil {
			return false, err
		}
		if bytes.Equal(enode[:], blockMinerEnode[:]) {
			return false, nil
		}
	}
	return true, nil
}

// It's reset when governance gets updated, i.e. search doesn't go back
// beyond modifiedBlock + 1.
// Not enforced if member count <= 2.
func (ma *wemixAdmin) nextMinerCandidates(height *big.Int) ([]*wemixNode, error) {
	var (
		ctx context.Context
		gov *metclient.RemoteContract
		err error
	)
	ctx = context.Background()
	if _, gov, _, _, err = ma.getRegGovEnvContracts(ctx, height); err != nil {
		return nil, wemixminer.ErrNotInitialized
	}
	e, err := getCoinbaseEnodeCache(ctx, height, gov)
	if err != nil {
		return nil, err
	}
	m := map[string]float64{}
	dix := (int(height.Int64()) + 1) % len(e.nodes) // default miner = height % member count
	for i, n := range e.nodes {
		if i < dix {
			i += len(e.nodes)
		}
		// score, the closer to the default miner, the lower the score
		m[n.Enode] = float64(len(e.nodes)-(i-dix)) / float64(2*len(e.nodes))
	}
	limit := len(e.nodes) / 2
	if len(e.nodes) <= 2 {
		limit = 0
	} else if limit > int(height.Int64()-e.modifiedBlock.Int64()-1) {
		limit = int(height.Int64() - e.modifiedBlock.Int64() - 1)
	}
	tooBig := float64(1000000000.0)
	for ix, h := 0, new(big.Int).Set(height); ix < len(e.nodes); h, ix = h.Sub(h, common.Big1), ix+1 {
		blockMinerEnode, err := getBlockMiner(ctx, admin.cli, e, h)
		if err != nil {
			continue
		}
		senode := string(blockMinerEnode)
		score := 1.0
		if ix < limit {
			score = tooBig
		}
		m[senode] += score
	}

	var keys []string
	for k, v := range m {
		if v <= tooBig {
			keys = append(keys, k)
		}
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})
	var miners []*wemixNode
	for _, k := range keys {
		if nix, ok := e.enode2index[k]; !ok || nix <= 0 {
			continue
		} else {
			miners = append(miners, e.nodes[nix-1])
		}
	}
	return miners, nil
}

// elect the most eligible candidate as the next leader
func (ma *wemixAdmin) electNextMiner(height *big.Int) error {
	tstart := time.Now()
	enode, err := hex.DecodeString(ma.self.Enode)
	if err != nil {
		return wemixminer.ErrNotInitialized
	}
	candidates, err := ma.nextMinerCandidates(height)
	if err != nil {
		return err
	}
	// miningPeers map
	peers := ma.collectMinerStates(height)
	closeEnough := int64(2)
	peersMap := map[string]*wemixapi.WemixMinerStatus{}
	if len(peers) > 0 {
		for _, p := range peers {
			if p.Status == "up" && height.Int64()-p.LatestBlockHeight.Int64() <= closeEnough {
				peersMap[p.NodeName] = p
			}
		}
	}

	for _, next := range candidates {
		if bytes.Equal(enode, []byte(next.Enode)) {
			return nil
		}
		if len(peers) > 0 {
			if _, ok := peersMap[next.Name]; !ok {
				continue
			}
		} else if !admin.isPeerUp(next.Id) {
			continue
		}
		if err = admin.etcdMoveLeader(next.Name); err == nil {
			log.Debug("new miner elected", "leader", next.Name, "height", height.Uint64()+1, "took", time.Since(tstart))
			return nil
		}
	}
	log.Error("failed to elect a new miner", "height", height.Uint64()+1, "took", time.Since(tstart))
	return nil
}

func refreshCoinbaseEnodeCache(height *big.Int) {
	_, _ = admin.nextMinerCandidates(height)
}

// EOF
