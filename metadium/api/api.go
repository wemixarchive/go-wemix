// metadium/api/api.go

package api

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type MetadiumMinerStatus struct {
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
	msgChannelLock = &sync.Mutex{}
	msgChannel     chan interface{}

	Info func() interface{}

	GetMinerStatus func() *MetadiumMinerStatus
	GetMiners      func(node string, timeout int) []*MetadiumMinerStatus

	EtcdInit         func() error
	EtcdAddMember    func(name string) (string, error)
	EtcdRemoveMember func(name string) (string, error)
	EtcdJoin         func(cluster string) error
	EtcdMoveLeader   func(name string) error
	EtcdGetWork      func() (string, error)
	EtcdDeleteWork   func() error
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
