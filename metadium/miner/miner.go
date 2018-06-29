// Copyright 2018 The go-metadium Authors

package miner

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	// notification channel when a new transaction arrives
	TxNotifier = make(chan bool, 1)

	// allow mining without PoA
	//DoMineWithout = false
	DoMineWithout = true
)

func TxNotify() {
	if params.ConsensusMethod != params.ConsensusPoW {
		TxNotifier <- true
	}
}

func IsMiner(height int) bool {
	return IsPoW() || DoMineWithout
}

func IsPoW() bool {
	return params.ConsensusMethod == params.ConsensusPoW
}

// EOF
