// Copyright 2025 The go-wemix Authors
// This file is part of the go-wemix library.

// Reproduction for the etcdReady data race.
//
// etcdReady was a plain bool written by the etcdEventHandler goroutine and read
// through etcdIsReady on other goroutines with no synchronization. This test
// runs the production handler through both of its writes, ready and stopped,
// while another goroutine keeps calling etcdIsReady. Run it with -race: an
// unsynchronized write in either place is reported here.
package wemix

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestEtcdReadyConcurrentAccessNoRace(t *testing.T) {
	survivor, victim := newTestEtcdCluster(t)

	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
				victim.etcdIsReady()
			}
		}
	}()
	defer func() {
		close(done)
		wg.Wait()
	}()

	atomic.StoreUint32(&etcdReady, 0)
	go etcdEventHandler()
	waitForEtcdCondition(t, "the handler to set etcdReady", victim.etcdIsReady)

	removeTestEtcdMember(t, survivor, uint64(victim.etcd.Server.ID()))
	waitForEtcdCondition(t, "the handler to clear etcdReady", func() bool { return !victim.etcdIsReady() })
}
