// sync.go

package metadium

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	metaapi "github.com/ethereum/go-ethereum/metadium/api"
	"github.com/ethereum/go-ethereum/params"
)

var (
	syncLock = &sync.Mutex{}
	leaderId uint64
	leader   *metaNode
)

func (ma *metaAdmin) getLatestBlockInfo(node *metaNode) (height *big.Int, hash common.Hash, td *big.Int, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	msgch := make(chan interface{}, 1)
	metaapi.SetMsgChannel(msgch)
	defer func() {
		metaapi.SetMsgChannel(nil)
		close(msgch)
	}()

	timer := time.NewTimer(60 * time.Second)
	err = ma.rpcCli.CallContext(ctx, nil, "admin_requestMinerStatus", &node.Id)
	if err != nil {
		log.Error("Metadium RequestMinerStatus Failed", "id", node.Id, "error", err)
		return
	}

	done := false
	for {
		if done {
			break
		}
		select {
		case msg := <-msgch:
			s, ok := msg.(*metaapi.MetadiumMinerStatus)
			if !ok {
				continue
			}
			if s.Name != node.Name {
				continue
			}
			height, hash, td, err = s.LatestBlockHeight, s.LatestBlockHash, s.LatestBlockTd, nil
			return

		case <-timer.C:
			err = ErrNotRunning
			return
		}
	}
	err = ethereum.NotFound
	return
}

// syncLock should be held by the caller
func (ma *metaAdmin) syncWith(node *metaNode) {
	height, hash, td, err := ma.getLatestBlockInfo(node)
	if err != nil {
		log.Error("Metadium", "failed to synchronize with", node.Name,
			"error", err)
		return
	} else {
		log.Error("Metadium", "synchronizing with", node.Name,
			"height", height, "hash", hash, "td", td)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = ma.rpcCli.CallContext(ctx, nil, "admin_synchroniseWith", &node.Id)
	if err != nil {
		log.Error("Metadium", "failed to synchronize with", node.Name,
			"error", err)
	} else {
		log.Error("Metadium", "synchronized with", node.Name)
	}
}

func (ma *metaAdmin) updateMiner(locked bool) {
	if ma.etcd == nil {
		return
	}

	lid, lnode := ma.etcdLeader(locked)
	if lid != leaderId && lid != 0 {
		syncLock.Lock()
		_, oldLeader := leaderId, leader
		leaderId, leader = lid, lnode
		if leader == ma.self && oldLeader != nil {
			log.Error("Metadium: we are the new leader",
				"syncing with", oldLeader.Name)
			ma.syncWith(oldLeader)
			log.Error("Metadium", "sync done with", oldLeader.Name)
		}
		syncLock.Unlock()
	}
}

func IsMiner(height int) bool {
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

		if height != admin.lastBlock {
			admin.update()
		}

		admin.updateMiner(false)

		if admin.etcdIsLeader() {
			return true
		} else {
			admin.blocksMined = 0
			return false
		}
	} else {
		return false
	}
}

// EOF
