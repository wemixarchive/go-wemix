// wemix/api/api.go

package api

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

type WemixMinerStatus struct {
	NodeName    string `json:"name"`
	Enode       string `json:"enode"`
	Id          string `json:"id"`
	Addr        string `json:"addr"`
	Status      string `json:"status"`
	Miner       bool   `json:"miner"`
	MiningPeers string `json:"miningPeers"`

	LatestBlockHeight *big.Int    `json:"latestBlockHeight"`
	LatestBlockHash   common.Hash `json:"latestBlockHash"`
	LatestBlockTd     *big.Int    `json:"latestBlockTd"`

	RttMs *big.Int `json:"rttMs"`
}

var (
	// miner status & etcd cluster events
	minerStatusFeed event.Feed
	etcdClusterFeed event.Feed

	msgChannelLock = &sync.Mutex{}
	msgChannel     chan interface{}

	Info func() interface{}

	GetMinerStatus func() *WemixMinerStatus
	GetMiners      func(node string, timeout int) []*WemixMinerStatus

	EtcdInit         func() error
	EtcdAddMember    func(name string) (string, error)
	EtcdRemoveMember func(name string) (string, error)
	EtcdJoin         func(cluster string) error
	EtcdMoveLeader   func(name string) error
	EtcdGetWork      func() (string, error)
	EtcdDeleteWork   func() error

	// for debugging
	EtcdPut    func(string, string) error
	EtcdGet    func(string) (string, error)
	EtcdDelete func(string) error
)

func (s *WemixMinerStatus) Clone() *WemixMinerStatus {
	// RLP payloads that omit *big.Int fields decode to nil, and
	// new(big.Int).Set(nil) panics. Copy each field nil-safely so a single
	// crafted StatusEx cannot crash the handler goroutine.
	safeBig := func(b *big.Int) *big.Int {
		if b == nil {
			return nil
		}
		return new(big.Int).Set(b)
	}
	return &WemixMinerStatus{
		NodeName:          s.NodeName,
		Enode:             s.Enode,
		Id:                s.Id,
		Addr:              s.Addr,
		Status:            s.Status,
		Miner:             s.Miner,
		MiningPeers:       s.MiningPeers,
		LatestBlockHeight: safeBig(s.LatestBlockHeight),
		LatestBlockHash:   s.LatestBlockHash,
		LatestBlockTd:     safeBig(s.LatestBlockTd),
		RttMs:             safeBig(s.RttMs),
	}
}

func SubscribeToMinerStatus(ch chan *WemixMinerStatus) event.Subscription {
	return minerStatusFeed.Subscribe(ch)
}

func SubscribeToEtcdCluster(ch chan string) event.Subscription {
	return etcdClusterFeed.Subscribe(ch)
}

func GotStatusEx(status *WemixMinerStatus) {
	minerStatusFeed.Send(status)
}

func GotEtcdCluster(cluster string) {
	etcdClusterFeed.Send(cluster)
}

// EOF
