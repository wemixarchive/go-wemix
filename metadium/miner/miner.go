// Copyright 2018 The go-metadium Authors

package miner

import (

)

var (
	// notification channel when a new transaction arrives
	TxNotifier = make(chan bool, 1)

	// allow mining without PoA
	//DoMineWithout = false
	DoMineWithout = true
)

func TxNotify() {
	TxNotifier <- true
}

func IsMiner(height int) bool {
	return DoMineWithout
}

// EOF
