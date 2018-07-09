// Copyright 2018 The go-metadium Authors

package miner

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	// notification channel when a new transaction arrives
	TxNotifier = make(chan bool, 1)
	IsMinerFunc func(int) bool
)

func TxNotify() {
	if params.ConsensusMethod != params.ConsensusPoW {
		TxNotifier <- true
	}
}

func IsMiner(height int) bool {
	if IsMinerFunc == nil {
		return false
	} else {
		return IsMinerFunc(height)
	}
}

func IsPoW() bool {
	return params.ConsensusMethod == params.ConsensusPoW
}

// EOF
