// sync.go

package metadium

import (
	"bytes"
	"context"
	"encoding/json"
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

	msgch := make(chan interface{}, 16)
	metaapi.SetMsgChannel(msgch)
	defer func() {
		metaapi.SetMsgChannel(nil)
		close(msgch)
	}()

	timer := time.NewTimer(60 * time.Second)
	err = ma.rpcCli.CallContext(ctx, nil, "admin_requestMinerStatus", &node.Id)
	if err != nil {
		log.Info("Metadium RequestMinerStatus Failed", "id", node.Id, "error", err)
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
			if s.NodeName != node.Name {
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
func (ma *metaAdmin) syncWith(node *metaNode) error {
	tsync := time.Now()
	height, hash, td, err := ma.getLatestBlockInfo(node)
	if err != nil {
		log.Error("Metadium", "failed to synchronize with", node.Name,
			"error", err, "took", time.Since(tsync))
		return err
	} else {
		log.Info("Metadium", "synchronizing with", node.Name,
			"height", height, "hash", hash, "td", td, "took", time.Since(tsync))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = ma.rpcCli.CallContext(ctx, nil, "admin_synchroniseWith", &node.Id)
	if err != nil {
		log.Error("Metadium", "failed to synchronize with", node.Name,
			"error", err, "took", time.Since(tsync))
	} else {
		log.Info("Metadium", "synchronized with", node.Name, "took", time.Since(tsync))
	}
	return err
}

// return true if this node still is the miner after update
func (ma *metaAdmin) updateMiner(locked bool) bool {
	if ma.etcd == nil {
		return false
	}

	syncLock.Lock()
	defer syncLock.Unlock()

	lid, lnode := ma.etcdLeader(locked)
	if lid == leaderId || lid == 0 {
		return lnode == ma.self
	}

	_, oldLeader := leaderId, leader
	leaderId, leader = lid, lnode
	if leader == ma.self && oldLeader != nil {
		// We are the new leader. Make sure we have the latest block.
		// Otherwise, punt the leadership to the next in line.
		// If all fails, accept the potential fork and move on.

		log.Info("Metadium: we are the new leader")
		tstart := time.Now()

		// get the latest work info from etcd
		getLatestWork := func() (*metaWork, error) {
			var (
				workInfo string
				work     *metaWork
				retries  = 60
				err      error
			)

			for ; retries > 0; retries-- {
				workInfo, err = ma.etcdGet("metadium-work")
				if err != nil {
					// TODO: ignore if error is not found
					log.Info("Metadium - cannot get the latest work info",
						"error", err, "took", time.Since(tstart))
					continue
				}

				if workInfo == "" {
					log.Info("Metadium - the latest work info not logged yet")
					return nil, nil
				} else {
					if err = json.Unmarshal([]byte(workInfo), &work); err != nil {
						log.Error("Metadium - cannot get the latest work info",
							"error", err, "took", time.Since(tstart))
						return nil, err
					}
					log.Info("Metadium - got the latest work info",
						"height", work.Height, "hash", work.Hash,
						"took", time.Since(tstart))
					return work, nil
				}
			}
			return nil, ethereum.NotFound
		}

		// check if we are in sync with the latest work info recorded
		inSync := func(work *metaWork) (synced bool, latestNum uint64, curNum uint64) {
			synced, latestNum, curNum = false, 0, 0

			if work == nil {
				synced = true
				return
			}
			latestNum = uint64(work.Height)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			cur, err := ma.cli.HeaderByNumber(ctx, big.NewInt(work.Height))
			if err != nil {
				return
			}
			curNum = uint64(cur.Number.Int64())
			synced = bytes.Equal(cur.Hash().Bytes(), work.Hash.Bytes())
			return
		}

		// if we are not in sync, punt the leadership to the next in line
		// if all fails, just move on
		puntLeadership := func() error {
			nodes := ma.getNodes()
			if len(nodes) == 0 {
				return ethereum.NotFound
			}

			ix := 0
			for i, node := range nodes {
				if node.Id == ma.self.Id {
					ix = i
					break
				}
			}
			if ix >= len(nodes) {
				return ethereum.NotFound
			}

			var err error
			for i, j := 0, (ix+1)%len(nodes); i < len(nodes)-1; i++ {
				err = ma.etcdMoveLeader(nodes[j].Name)
				if err == nil {
					return nil
				}
				j = (j + 1) % len(nodes)
			}

			return err
		}

		work, err := getLatestWork()
		if err != nil {
			log.Error("Metadium - cannot get the latest work information. Yielding leadeship")
			err = puntLeadership()
			if err != nil {
				log.Error("Metadium - leadership yielding failed", "error", err)
			} else {
				log.Info("Metadium - yielded leadership")
			}
		} else if work == nil {
			// this must be the first block, juts move on
			log.Info("Metadium - not initialized yet. Starting mining")
		} else if synced, _, height := inSync(work); synced {
			log.Info("Metadium - in sync. Starting mining", "height", height)
		} else {
			// sync with the previous leader
			ma.syncWith(oldLeader)

			// check sync again
			work, err = getLatestWork()
			if err != nil {
				log.Error("Metadium - cannot get the latest work", "error", err)
			} else if work == nil {
				// this must be the first block, juts move on
			} else if synced, _, height := inSync(work); !synced {
				// if still not in sync, give up leadership
				err = puntLeadership()
				if err != nil {
					log.Error("Metadium - not in sync. Leadership yielding failed",
						"latest", work.Height, "current", height, "error", err)
				} else {
					log.Error("Metadium - not in sync. Yielded leadership",
						"latest", work.Height, "current", height)
				}
			}
		}

		// update leader info again
		lid, lnode = ma.etcdLeader(locked)
		if lid != leaderId && lid != 0 {
			leaderId, leader = lid, lnode
		}
	}

	return leader == ma.self
}

func IsMiner() bool {
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

		if admin.etcdIsLeader() {
			return admin.updateMiner(false)
		} else {
			admin.blocksMined = 0
			return false
		}
	} else {
		return false
	}
}

// EOF
