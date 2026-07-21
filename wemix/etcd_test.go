// Copyright 2025 The go-wemix Authors
// This file is part of the go-wemix library.

// Tests that require a real embedded etcd cluster. The test spins up a
// single-node embedded etcd so that etcd transaction semantics are exercised
// directly, rather than mirrored.

package wemix

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.etcd.io/etcd/server/v3/embed"
	"go.etcd.io/etcd/server/v3/etcdserver/api/v3client"
)

// freeBasePort returns a base port p such that p+1 (etcd peer) and p+2 (etcd
// client) are usable. It probes three consecutive ports and only returns once
// all of them can be bound, to avoid flaky startup collisions.
func freeBasePort(t *testing.T) int {
	t.Helper()
	for attempt := 0; attempt < 20; attempt++ {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			continue
		}
		base := l.Addr().(*net.TCPAddr).Port
		l.Close()

		ok := true
		for _, off := range []int{1, 2} {
			probe, perr := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", base+off))
			if perr != nil {
				ok = false
				break
			}
			probe.Close()
		}
		if ok {
			return base
		}
	}
	t.Fatal("could not find three consecutive free ports for embedded etcd")
	return 0
}

// newTestEtcdAdmin starts a single-node embedded etcd and returns a
// *wemixAdmin bound to it.
//
// It deliberately does NOT go through (*wemixAdmin).etcdInit: that path spawns
// the production etcdEventHandler goroutine, which is a process-wide singleton
// keyed off the package-level `admin` and is not designed to be started and
// torn down repeatedly inside tests. Since etcdResetWork only depends on a
// ready etcd client/server (etcdIsReady) and ReqTimeout(), we start the
// embedded server directly and set etcdReady ourselves.
func newTestEtcdAdmin(t *testing.T) *wemixAdmin {
	t.Helper()

	base := freeBasePort(t)
	ma := &wemixAdmin{
		self: &wemixNode{
			Name: "testnode",
			Ip:   "127.0.0.1",
			Port: base,
		},
		etcdDir:     t.TempDir(),
		etcdTimeout: 5 * time.Second,
	}

	cfg := ma.etcdNewConfig(true)
	etcd, err := embed.StartEtcd(cfg)
	if err != nil {
		t.Fatalf("StartEtcd failed: %v", err)
	}

	select {
	case <-etcd.Server.ReadyNotify():
	case <-time.After(30 * time.Second):
		etcd.Server.Stop()
		<-etcd.Server.StopNotify()
		t.Fatal("embedded etcd did not become ready in time")
	}

	ma.etcd = etcd
	ma.etcdCli = v3client.New(etcd.Server)

	// etcdIsReady() also checks the package-level etcdReady flag.
	prevReady := etcdReady
	etcdReady = true

	t.Cleanup(func() {
		etcdReady = prevReady
		if ma.etcdCli != nil {
			ma.etcdCli.Close()
		}
		etcd.Close()
		// The data dir is a t.TempDir(), cleaned up automatically.
	})

	return ma
}

// mustGetWork reads the current wemixWorkKey value, treating a not-found as "".
func mustGetWork(t *testing.T, ma *wemixAdmin) string {
	t.Helper()
	v, err := ma.etcdGet(wemixWorkKey)
	if err == ErrNotFound {
		return ""
	}
	if err != nil {
		t.Fatalf("etcdGet(work) failed: %v", err)
	}
	return v
}

func marshalWork(t *testing.T, height int64, hashHex string) string {
	t.Helper()
	b, err := json.Marshal(&wemixWork{Height: height, Hash: common.HexToHash(hashHex)})
	if err != nil {
		t.Fatalf("marshal work: %v", err)
	}
	return string(b)
}

// TestEtcdResetWork_TokenHeld verifies that a caller holding the exact token
// stored under wemixTokenKey can overwrite wemixWorkKey.
func TestEtcdResetWork_TokenHeld(t *testing.T) {
	ma := newTestEtcdAdmin(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Acquire a real token exactly like production (syncCheck) does.
	token, err := ma.acquireToken(ctx, big.NewInt(1001), MiningTokenTTL)
	if err != nil {
		t.Fatalf("acquireToken failed: %v", err)
	}

	// Seed an initial work value.
	oldWork := marshalWork(t, 1000, "0x1111111111111111111111111111111111111111111111111111111111111111")
	if _, err := ma.etcdPut(wemixWorkKey, oldWork); err != nil {
		t.Fatalf("seed work failed: %v", err)
	}

	newWork := marshalWork(t, 1001, "0x2222222222222222222222222222222222222222222222222222222222222222")
	if err := ma.etcdResetWork(token, newWork); err != nil {
		t.Fatalf("etcdResetWork returned error while holding token: %v", err)
	}

	if got := mustGetWork(t, ma); got != newWork {
		t.Fatalf("work not updated: want %q, got %q", newWork, got)
	}
}

// TestEtcdResetWork_TokenMismatch verifies that when the stored token differs
// from the caller's token, the work value is left untouched and ErrInvalidToken
// is returned.
func TestEtcdResetWork_TokenMismatch(t *testing.T) {
	ma := newTestEtcdAdmin(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// The token actually stored in etcd.
	stored, err := ma.acquireToken(ctx, big.NewInt(1001), MiningTokenTTL)
	if err != nil {
		t.Fatalf("acquireToken failed: %v", err)
	}

	oldWork := marshalWork(t, 1000, "0x1111111111111111111111111111111111111111111111111111111111111111")
	if _, err := ma.etcdPut(wemixWorkKey, oldWork); err != nil {
		t.Fatalf("seed work failed: %v", err)
	}

	// A different token (different Height) — same shape, not what is stored.
	other := *stored
	other.Height = big.NewInt(2002)

	newWork := marshalWork(t, 1001, "0x2222222222222222222222222222222222222222222222222222222222222222")
	if err := ma.etcdResetWork(&other, newWork); err != ErrInvalidToken {
		t.Fatalf("expected ErrInvalidToken, got %v", err)
	}

	if got := mustGetWork(t, ma); got != oldWork {
		t.Fatalf("work must be unchanged on mismatch: want %q, got %q", oldWork, got)
	}
}

// TestEtcdResetWork_TokenAbsent verifies that when no token is present (e.g. it
// expired and was deleted by another node), the work value is left untouched.
func TestEtcdResetWork_TokenAbsent(t *testing.T) {
	ma := newTestEtcdAdmin(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token, err := ma.acquireToken(ctx, big.NewInt(1001), MiningTokenTTL)
	if err != nil {
		t.Fatalf("acquireToken failed: %v", err)
	}

	oldWork := marshalWork(t, 1000, "0x1111111111111111111111111111111111111111111111111111111111111111")
	if _, err := ma.etcdPut(wemixWorkKey, oldWork); err != nil {
		t.Fatalf("seed work failed: %v", err)
	}

	// Remove the token, simulating expiry/cleanup by another node.
	if err := ma.etcdDelete(wemixTokenKey); err != nil {
		t.Fatalf("delete token failed: %v", err)
	}

	newWork := marshalWork(t, 1001, "0x2222222222222222222222222222222222222222222222222222222222222222")
	if err := ma.etcdResetWork(token, newWork); err != ErrInvalidToken {
		t.Fatalf("expected ErrInvalidToken when token absent, got %v", err)
	}

	if got := mustGetWork(t, ma); got != oldWork {
		t.Fatalf("work must be unchanged when token absent: want %q, got %q", oldWork, got)
	}
}

// TestEtcdResetWork_ExpiredThenSuperseded reproduces the real stale-overwrite
// scenario: this node acquires a token, the token expires, and another miner
// then acquires the token (which, per acquireToken, deletes the expired token
// and writes a fresh one). When the original — now stale — caller finally runs
// etcdResetWork, the stored token no longer matches, so the reset is rejected
// with ErrInvalidToken and the work value the new holder set is preserved.
func TestEtcdResetWork_ExpiredThenSuperseded(t *testing.T) {
	ma := newTestEtcdAdmin(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. This node acquires a short-lived token.
	staleToken, err := ma.acquireToken(ctx, big.NewInt(1001), 1)
	if err != nil {
		t.Fatalf("acquireToken failed: %v", err)
	}

	oldWork := marshalWork(t, 1000, "0x1111111111111111111111111111111111111111111111111111111111111111")
	if _, err := ma.etcdPut(wemixWorkKey, oldWork); err != nil {
		t.Fatalf("seed work failed: %v", err)
	}

	// 2. The token expires.
	time.Sleep(2200 * time.Millisecond) // > TTL(1s) plus a full-second margin for Unix() truncation
	if staleToken.ttl() >= 0 {
		t.Fatalf("token should be logically expired, ttl=%d", staleToken.ttl())
	}

	// 3. Another miner acquires the token. acquireToken deletes the expired
	//    token and writes a fresh one; the stored value now differs from
	//    staleToken. Simulate that miner also advancing the work value.
	newHolder, err := ma.acquireToken(ctx, big.NewInt(1002), MiningTokenTTL)
	if err != nil {
		t.Fatalf("second acquireToken failed: %v", err)
	}
	newHolderWork := marshalWork(t, 1001, "0x3333333333333333333333333333333333333333333333333333333333333333")
	if err := ma.etcdResetWork(newHolder, newHolderWork); err != nil {
		t.Fatalf("new holder should be able to reset work: %v", err)
	}

	// 4. The original, now-stale caller tries to reset work with its expired
	//    token. It must be rejected, leaving the new holder's work intact.
	staleWork := marshalWork(t, 1001, "0x2222222222222222222222222222222222222222222222222222222222222222")
	if err := ma.etcdResetWork(staleToken, staleWork); err != ErrInvalidToken {
		t.Fatalf("stale caller must be rejected with ErrInvalidToken, got %v", err)
	}
	if got := mustGetWork(t, ma); got != newHolderWork {
		t.Fatalf("stale caller overwrote work: want %q, got %q", newHolderWork, got)
	}
}

// TestEtcdResetWork_NotReady verifies the guard when etcd is not running.
func TestEtcdResetWork_NotReady(t *testing.T) {
	ma := &wemixAdmin{} // no etcd
	if err := ma.etcdResetWork(&WemixToken{}, "x"); err != ErrNotRunning {
		t.Fatalf("expected ErrNotRunning, got %v", err)
	}
}
