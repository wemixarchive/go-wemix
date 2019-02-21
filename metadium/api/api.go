// metadium/api/api.go

package api

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type MetadiumMinerStatus struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Addr        string `json:"addr"`
	Status      string `json:"status"`
	Miner       bool   `json:"miner"`
	MiningPeers string `json:"mining-peers"`

	Mining            bool        `json:"mining"`
	LatestBlockHeight *big.Int    `json:"latest-block-height"`
	LatestBlockHash   common.Hash `json:"latest-block-hash"`

	Syncing      bool                  `json:"syncing"`
	SyncProgress ethereum.SyncProgress `json:"sync-progress"`
}

var (
	msgChannelLock = &sync.Mutex{}
	msgChannel     chan interface{}

	Info func() interface{}

	GetMinerStatus func() *MetadiumMinerStatus
	GetMiners      func(node string) []*MetadiumMinerStatus

	EtcdInit         func() error
	EtcdAddMember    func(name string) (string, error)
	EtcdRemoveMember func(name string) (string, error)
	EtcdJoin         func(cluster string) error
	EtcdMoveLeader   func(name string) error
)

func SetMsgChannel(ch chan interface{}) {
	msgChannelLock.Lock()
	defer msgChannelLock.Unlock()
	msgChannel = ch
}

func GotStatusEx(status *MetadiumMinerStatus) {
	msgChannelLock.Lock()
	defer msgChannelLock.Unlock()
	if msgChannel != nil {
		msgChannel <- status
	}
}

func GotEtcdCluster(cluster string) {
	msgChannelLock.Lock()
	defer msgChannelLock.Unlock()
	if msgChannel != nil {
		msgChannel <- cluster
	}
}

// EOF
