// Copyright 2025 The go-wemix Authors
// This file is part of the go-wemix library.

// Flag semantics for etcd liveness tracking. This test pins the contract the run
// loop relies on to restart etcd: clearing etcdRunning makes both etcdIsRunning()
// and etcdIsReady() return false. etcd_stop_recovery_test.go covers the handler
// clearing it on a real stop.
package wemix

import (
	"sync/atomic"
	"testing"
)

func TestEtcdRunningFlagGatesReadiness(t *testing.T) {
	ma := newTestEtcdAdmin(t) // starts a ready single-node etcd

	if !ma.etcdIsRunning() || !ma.etcdIsReady() {
		t.Fatalf("fresh etcd should be running and ready: running=%v ready=%v",
			ma.etcdIsRunning(), ma.etcdIsReady())
	}

	// Mirror what etcdEventHandler does when the server stops.
	atomic.StoreUint32(&etcdRunning, 0)

	if ma.etcdIsRunning() {
		t.Error("etcdIsRunning() must be false once etcdRunning is cleared")
	}
	if ma.etcdIsReady() {
		t.Error("etcdIsReady() must be false once the server is no longer running")
	}

	// Restore so the harness cleanup sees a consistent live server.
	atomic.StoreUint32(&etcdRunning, 1)
}
