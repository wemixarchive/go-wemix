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
	statusExLock = &sync.Mutex{}
	statusExCh   chan *MetadiumMinerStatus

	Info           func() interface{}
	GetMinerStatus func() *MetadiumMinerStatus
	GetMiners      func(node string) []*MetadiumMinerStatus
)

func SetStatusExChannel(ch chan *MetadiumMinerStatus) {
	statusExLock.Lock()
	defer statusExLock.Unlock()
	statusExCh = ch
}

func GotStatusEx(status *MetadiumMinerStatus) {
	statusExLock.Lock()
	defer statusExLock.Unlock()
	if statusExCh != nil {
		statusExCh <- status
	}
}

// EOF
