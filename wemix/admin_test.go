// Copyright 2026 The go-wemix Authors
// This file is part of the go-wemix library.

package wemix

import (
	"fmt"
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
	wemixapi "github.com/ethereum/go-ethereum/wemix/api"
)

// fakeMinerAdminAPI stands in for the "admin" RPC namespace used by getMiners.
// Every peer is reported up, and the queued responses are published only on the
// last status request so their arrival order is deterministic.
type fakeMinerAdminAPI struct {
	mu        sync.Mutex
	calls     int
	total     int
	responses []*wemixapi.WemixMinerStatus
}

func (f *fakeMinerAdminAPI) PeerInfo(id string) (*p2p.NodeInfo, error) {
	return &p2p.NodeInfo{ID: id}, nil
}

func (f *fakeMinerAdminAPI) RequestMinerStatus(id string) error {
	f.mu.Lock()
	f.calls++
	last := f.calls == f.total
	f.mu.Unlock()
	if last {
		for _, s := range f.responses {
			wemixapi.GotStatusEx(s)
		}
	}
	return nil
}

func newTestMinersAdmin(t *testing.T, names []string, api *fakeMinerAdminAPI) {
	t.Helper()

	server := rpc.NewServer()
	if err := server.RegisterName("admin", api); err != nil {
		t.Fatalf("register fake admin api: %v", err)
	}
	cli := rpc.DialInProc(server)

	nodes := map[string]*wemixNode{}
	for i, name := range names {
		id := fmt.Sprintf("id-%s", name)
		nodes[id] = &wemixNode{Name: name, Id: id, Ip: "127.0.0.1", Port: 30000 + i}
	}

	prev := admin
	admin = &wemixAdmin{rpcCli: cli, lock: &sync.Mutex{}, nodes: nodes}
	t.Cleanup(func() {
		admin = prev
		cli.Close()
		server.Stop()
	})
}

func minerStatus(name string, height int64, hash common.Hash) *wemixapi.WemixMinerStatus {
	return &wemixapi.WemixMinerStatus{
		NodeName:          name,
		Status:            "up",
		LatestBlockHeight: big.NewInt(height),
		LatestBlockHash:   hash,
	}
}

// A single registered validator repeating its StatusEx response must be
// counted once, so it cannot satisfy the recovery-vote majority on its own.
func TestGetMiners_DuplicateResponses(t *testing.T) {
	names := []string{"attacker", "honest1", "honest2", "honest3", "honest4"}
	attackerHeight, attackerHash := int64(99999), common.HexToHash("0xdeadbeef")
	honestHeight, honestHash := int64(100), common.HexToHash("0x01")

	api := &fakeMinerAdminAPI{total: len(names)}
	for i := 0; i < 6; i++ {
		api.responses = append(api.responses, minerStatus("attacker", attackerHeight, attackerHash))
	}
	for _, name := range names[1:] {
		api.responses = append(api.responses, minerStatus(name, honestHeight, honestHash))
	}
	newTestMinersAdmin(t, names, api)

	states := getMiners("", 5)

	if len(states) != len(names) {
		t.Fatalf("got %d states, want %d (one per registered node)", len(states), len(names))
	}
	seen := map[string]int{}
	for _, s := range states {
		seen[s.NodeName]++
	}
	for _, name := range names {
		if seen[name] != 1 {
			t.Errorf("node %q counted %d times, want 1", name, seen[name])
		}
	}

	height, hash := findConsensusBlock(states)
	if height == nil || height.Int64() != honestHeight || hash != honestHash {
		t.Fatalf("consensus = (%v, %x), want honest (%d, %x)", height, hash, honestHeight, honestHash)
	}
}
