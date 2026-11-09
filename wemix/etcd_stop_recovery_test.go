// Copyright 2025 The go-wemix Authors
// This file is part of the go-wemix library.

// Stop handling of etcdEventHandler against a real three-member embedded
// cluster. Each test drives the production handler through an event that can
// end a member, and pins what the run loop relies on to recover: etcdIsRunning
// turns false only after the server is released, and the rejoin marker is left
// only when the cluster has rejected this member.
package wemix

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	wemixapi "github.com/ethereum/go-ethereum/wemix/api"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
	"go.etcd.io/etcd/server/v3/etcdserver/api/membership"
	"go.etcd.io/etcd/server/v3/etcdserver/api/v3client"
)

// etcdStopTimeout bounds how long a stop may take to reach the handler. A
// removed member stops about a second after the removal, then Close waits on
// the request timeout.
const etcdStopTimeout = 60 * time.Second

// etcdListenerFailureWindow is how long a member must stay ready after losing a
// listener. The handler sees the failure within milliseconds, so a stop shows
// up long before this.
const etcdListenerFailureWindow = 2 * time.Second

// startTestEtcdMember adds a member called name through leader and starts it
// as a joiner, the way etcdJoin does.
func startTestEtcdMember(t *testing.T, leader *wemixAdmin, name string) *wemixAdmin {
	t.Helper()

	ma := &wemixAdmin{
		self:        &wemixNode{Name: name, Ip: "127.0.0.1", Port: freeBasePort(t)},
		etcdDir:     t.TempDir(),
		etcdTimeout: 5 * time.Second,
	}
	peerURL := fmt.Sprintf("https://%s:%d", ma.self.Ip, ma.self.Port+1)

	ctx, cancel := context.WithTimeout(context.Background(), etcdStopTimeout)
	defer cancel()
	// etcd refuses to add a member until the last one has been active for its
	// health interval, so a fresh cluster answers "unhealthy cluster" for a while.
	var resp *clientv3.MemberAddResponse
	waitForEtcdCondition(t, "MemberAdd("+name+") to be accepted", func() bool {
		var err error
		resp, err = leader.etcdCli.MemberAdd(ctx, []string{peerURL})
		return err == nil
	})

	var cluster []string
	for _, m := range resp.Members {
		memberName := m.Name
		if memberName == "" {
			memberName = name
		}
		cluster = append(cluster, fmt.Sprintf("%s=%s", memberName, m.PeerURLs[0]))
	}

	cfg := ma.etcdNewConfig(false)
	cfg.InitialCluster = strings.Join(cluster, ",")
	etcd, err := embed.StartEtcd(cfg)
	if err != nil {
		t.Fatalf("StartEtcd(%s) failed: %v", name, err)
	}
	select {
	case <-etcd.Server.ReadyNotify():
	case <-time.After(etcdStopTimeout):
		t.Fatalf("member %s did not become ready", name)
	}
	ma.etcd = etcd
	ma.etcdCli = v3client.New(etcd.Server)
	return ma
}

// newTestEtcdCluster returns a survivor and a victim from a three-member
// cluster, leaving a quorum on the survivors when the victim goes away. The
// package-level admin points at the victim, whose handler the test drives.
func newTestEtcdCluster(t *testing.T) (survivor, victim *wemixAdmin) {
	t.Helper()

	survivor = newTestEtcdAdmin(t)
	bystander := startTestEtcdMember(t, survivor, "bystander")
	t.Cleanup(bystander.etcd.Close)
	victim = startTestEtcdMember(t, survivor, "victim")

	prevAdmin := admin
	admin = victim
	t.Cleanup(func() {
		// The handler closes a server it saw stop, and closing one twice panics,
		// so only stop the server and give a live handler time to finish before
		// taking admin away from it.
		if atomic.LoadUint32(&etcdRunning) == 1 {
			admin.etcd.Server.Stop()
			deadline := time.Now().Add(etcdStopTimeout)
			for atomic.LoadUint32(&etcdRunning) == 1 && time.Now().Before(deadline) {
				time.Sleep(10 * time.Millisecond)
			}
		}
		admin = prevAdmin
	})
	return survivor, victim
}

func removeTestEtcdMember(t *testing.T, survivor *wemixAdmin, id uint64) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), etcdStopTimeout)
	defer cancel()
	// The same health check as MemberAdd applies to a removal.
	waitForEtcdCondition(t, "MemberRemove to be accepted", func() bool {
		_, err := survivor.etcdCli.MemberRemove(ctx, id)
		return err == nil
	})
}

func waitForEtcdCondition(t *testing.T, what string, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(etcdStopTimeout)
	for !cond() {
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for %s", what)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// assertEtcdReleased fails unless the victim's ports can be bound again, which
// a restart needs as soon as etcdIsRunning turns false.
func assertEtcdReleased(t *testing.T, ma *wemixAdmin) {
	t.Helper()
	for _, port := range []int{ma.self.Port + 1, ma.self.Port + 2} {
		l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
		if err != nil {
			t.Fatalf("port %d still held once etcdIsRunning is false: %v", port, err)
		}
		l.Close()
	}
}

func TestEtcdRemovedWhileRunningRequestsRejoin(t *testing.T) {
	survivor, victim := newTestEtcdCluster(t)

	atomic.StoreUint32(&etcdReady, 0)
	go etcdEventHandler()
	waitForEtcdCondition(t, "the victim to become ready", victim.etcdIsReady)

	removeTestEtcdMember(t, survivor, uint64(victim.etcd.Server.ID()))
	waitForEtcdCondition(t, "the victim to stop", func() bool { return !victim.etcdIsRunning() })

	if victim.etcdIsReady() {
		t.Error("etcdIsReady() must be false once the member is removed")
	}
	if !victim.rejoinNeeded() {
		t.Error("the rejoin marker must be on disk by the time etcdIsRunning is false")
	}
	assertEtcdReleased(t, victim)
}

// A node that is down while it is removed has no handler to record the removal.
// It restarts on its old directory, so the handler must catch the rejection
// before the server ever becomes ready.
func TestEtcdRemovedWhileOfflineRequestsRejoin(t *testing.T) {
	survivor, victim := newTestEtcdCluster(t)

	id := uint64(victim.etcd.Server.ID())
	victim.etcd.Close()
	atomic.StoreUint32(&etcdRunning, 0)
	atomic.StoreUint32(&etcdReady, 0)
	removeTestEtcdMember(t, survivor, id)

	if victim.rejoinNeeded() {
		t.Fatal("an offline node cannot have recorded its own removal")
	}
	if err := victim.etcdStart(); err != nil {
		t.Fatalf("etcdStart failed: %v", err)
	}
	waitForEtcdCondition(t, "the rejected restart to stop", func() bool { return !victim.etcdIsRunning() })

	if !victim.rejoinNeeded() {
		t.Error("a member rejected on restart must request a rejoin")
	}
	assertEtcdReleased(t, victim)
}

// A failed listener leaves this node a healthy member. The block producer talks
// to the server in-process and outbound peer traffic does not use the listener,
// so the member keeps serving instead of giving up its place in the cluster.
func TestEtcdListenerFailureKeepsMemberServing(t *testing.T) {
	listeners := map[string]func(*wemixAdmin) error{
		"client": func(ma *wemixAdmin) error { return ma.etcd.Clients[0].Close() },
		"peer":   func(ma *wemixAdmin) error { return ma.etcd.Peers[0].Listener.Close() },
	}
	for name, closeListener := range listeners {
		closeListener := closeListener
		t.Run(name, func(t *testing.T) {
			_, victim := newTestEtcdCluster(t)

			atomic.StoreUint32(&etcdReady, 0)
			go etcdEventHandler()
			waitForEtcdCondition(t, "the victim to become ready", victim.etcdIsReady)

			if err := closeListener(victim); err != nil {
				t.Fatalf("closing the listener failed: %v", err)
			}
			deadline := time.Now().Add(etcdListenerFailureWindow)
			for time.Now().Before(deadline) {
				if !victim.etcdIsReady() {
					t.Fatal("a listener failure must not stop a healthy member")
				}
				time.Sleep(10 * time.Millisecond)
			}
			if _, err := victim.etcdPut("listener-failure", name); err != nil {
				t.Errorf("the member must keep serving: %v", err)
			}
			if victim.rejoinNeeded() {
				t.Error("a member that was not removed must not request a rejoin")
			}
		})
	}
}

// A crash while clearing the directory must leave the marker behind. Without
// it the restart path opens whatever remains, and a partial WAL ends the process.
func TestResetEtcdDirKeepsMarkerUntilDataIsGone(t *testing.T) {
	ma := &wemixAdmin{etcdDir: t.TempDir()}
	member := filepath.Join(ma.etcdDir, "member")
	if err := os.Mkdir(member, 0700); err != nil {
		t.Fatal(err)
	}
	ma.markRejoinNeeded(1)

	interrupted := errors.New("interrupted")
	err := clearDirMarkerLast(ma.etcdDir, etcdRejoinMarker, func(path string) error {
		if path == member {
			return interrupted
		}
		return os.RemoveAll(path)
	})
	if !errors.Is(err, interrupted) {
		t.Fatalf("clearDirMarkerLast() = %v, want %v", err, interrupted)
	}
	if !ma.rejoinNeeded() {
		t.Fatal("the marker must survive until the data is gone")
	}

	if err := ma.resetEtcdDir(); err != nil {
		t.Fatalf("resetEtcdDir() = %v", err)
	}
	entries, err := os.ReadDir(ma.etcdDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Errorf("the directory must end up empty, found %d entries", len(entries))
	}
}

// etcdJoinPeer stands in for the peer that answers a join request: it hands back
// the cluster string the way the eth handler does.
type etcdJoinPeer struct{ answer func() string }

func (p *etcdJoinPeer) RequestEtcdAddMember(string) error {
	go func() { wemixapi.GotEtcdCluster(p.answer()) }()
	return nil
}

// serveEtcdJoin wires ma so that joining the node named "peer" gets answer's
// cluster string, and makes ma the package-level admin.
func serveEtcdJoin(t *testing.T, ma *wemixAdmin, answer func() string) {
	t.Helper()

	ma.lock = &sync.Mutex{}
	ma.nodes = map[string]*wemixNode{"peer": {Name: "peer", Id: "peer"}}

	server := rpc.NewServer()
	t.Cleanup(server.Stop)
	if err := server.RegisterName("admin", &etcdJoinPeer{answer: answer}); err != nil {
		t.Fatal(err)
	}
	ma.rpcCli = rpc.DialInProc(server)
	t.Cleanup(ma.rpcCli.Close)

	prevAdmin := admin
	admin = ma
	t.Cleanup(func() { admin = prevAdmin })
}

// A join that lost the race against a start must leave the running server alone.
// The rejoin path clears the data directory, which that server is using.
func TestEtcdJoinRefusesRunningServer(t *testing.T) {
	ma := newTestEtcdAdmin(t)
	serveEtcdJoin(t, ma, func() string {
		return fmt.Sprintf("%s=https://%s:%d", ma.self.Name, ma.self.Ip, ma.self.Port+1)
	})

	ma.markRejoinNeeded(1)
	if err := ma.etcdJoin("peer"); !errors.Is(err, ErrAlreadyRunning) {
		t.Errorf("etcdJoin() = %v, want %v", err, ErrAlreadyRunning)
	}
	if matches, _ := filepath.Glob(filepath.Join(ma.etcdDir, "member", "wal", "*.wal")); len(matches) == 0 {
		t.Error("the running server's WAL must survive")
	}
}

// addMemberForJoin does what a peer does on a join request: it adds joiner and
// returns the new cluster string. It retries while etcd refuses the change.
func addMemberForJoin(survivor, joiner *wemixAdmin) string {
	peerURL := url.URL{Scheme: "https", Host: fmt.Sprintf("%s:%d", joiner.self.Ip, joiner.self.Port+1)}
	now := time.Now()
	member := membership.NewMember(joiner.self.Name, []url.URL{peerURL}, etcdClusterName, &now)

	ctx, cancel := context.WithTimeout(context.Background(), etcdStopTimeout)
	defer cancel()
	for ctx.Err() == nil {
		members, err := survivor.etcd.Server.AddMember(ctx, *member)
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		var cluster []string
		for _, m := range members {
			cluster = append(cluster, fmt.Sprintf("%s=%s", m.Attributes.Name, m.RaftAttributes.PeerURLs[0]))
		}
		return strings.Join(cluster, ",")
	}
	return ""
}

// A removed member rejoins through the production join path as a new member on
// a clean directory, however it learned that it was removed.
func TestEtcdRemovedMemberRejoins(t *testing.T) {
	removals := map[string]func(t *testing.T, survivor, victim *wemixAdmin){
		"while running": func(t *testing.T, survivor, victim *wemixAdmin) {
			atomic.StoreUint32(&etcdReady, 0)
			go etcdEventHandler()
			waitForEtcdCondition(t, "the victim to become ready", victim.etcdIsReady)
			removeTestEtcdMember(t, survivor, uint64(victim.etcd.Server.ID()))
		},
		"while offline": func(t *testing.T, survivor, victim *wemixAdmin) {
			id := uint64(victim.etcd.Server.ID())
			victim.etcd.Close()
			atomic.StoreUint32(&etcdRunning, 0)
			atomic.StoreUint32(&etcdReady, 0)
			removeTestEtcdMember(t, survivor, id)
			if err := victim.etcdStart(); err != nil {
				t.Fatalf("etcdStart failed: %v", err)
			}
		},
	}
	for name, remove := range removals {
		remove := remove
		t.Run(name, func(t *testing.T) {
			survivor, victim := newTestEtcdCluster(t)
			serveEtcdJoin(t, victim, func() string { return addMemberForJoin(survivor, victim) })
			oldID := victim.etcd.Server.ID()

			remove(t, survivor, victim)
			waitForEtcdCondition(t, "the victim to stop", func() bool { return !victim.etcdIsRunning() })

			if err := victim.etcdJoin("peer"); err != nil {
				t.Fatalf("etcdJoin failed: %v", err)
			}
			waitForEtcdCondition(t, "the victim to rejoin", victim.etcdIsReady)

			if victim.etcd.Server.ID() == oldID {
				t.Error("a rejoined node must come back as a new member")
			}
			if victim.rejoinNeeded() {
				t.Error("the rejoin marker must be gone once the member is ready")
			}
			if _, err := victim.etcdPut("rejoined", "yes"); err != nil {
				t.Fatalf("the rejoined member must serve writes: %v", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), etcdStopTimeout)
			defer cancel()
			resp, err := survivor.etcdCli.Get(ctx, "rejoined")
			if err != nil || resp.Count != 1 {
				t.Errorf("the survivor must see the write: count=%d err=%v", resp.Count, err)
			}
		})
	}
}

// The run loop must not restart on a directory the cluster has rejected. That
// server would only be refused again, and a partly cleared one can end the
// process. The join it hands over to is held back here by the auto-join lock.
func TestEtcdStartSkipsRestartWhenRejoinNeeded(t *testing.T) {
	// Held past the cluster teardown: EtcdStart's join goroutine has nothing to
	// wait on, and it must find the lock taken whenever it gets scheduled.
	etcdAutoJoinLock <- struct{}{}
	t.Cleanup(func() { <-etcdAutoJoinLock })

	_, victim := newTestEtcdCluster(t)
	victim.lock = &sync.Mutex{}
	victim.self.Id = "victim"
	victim.nodes = map[string]*wemixNode{victim.self.Id: victim.self}
	victim.etcd.Close()
	atomic.StoreUint32(&etcdRunning, 0)
	atomic.StoreUint32(&etcdReady, 0)
	victim.markRejoinNeeded(1)

	EtcdStart()

	if victim.etcdIsRunning() {
		t.Error("EtcdStart must not restart a server that needs a rejoin")
	}
	if !victim.rejoinNeeded() {
		t.Error("EtcdStart must leave the marker for the join path")
	}
}

// A node dropped from governance has no place in the cluster, so nothing should
// try to bring its etcd back. ma.self outlives the removal, which is why the
// membership has to be read from the node list.
func TestEtcdStartSkipsNodeOutsideGovernance(t *testing.T) {
	etcdAutoJoinLock <- struct{}{}
	t.Cleanup(func() { <-etcdAutoJoinLock })

	_, victim := newTestEtcdCluster(t)
	victim.lock = &sync.Mutex{}
	victim.self.Id = "victim"
	victim.nodes = map[string]*wemixNode{"other": {Name: "other", Id: "other"}}
	victim.etcd.Close()
	atomic.StoreUint32(&etcdRunning, 0)
	atomic.StoreUint32(&etcdReady, 0)

	EtcdStart()

	if victim.etcdIsRunning() {
		t.Fatal("EtcdStart must leave a node outside governance alone")
	}
}

// The node's own entry in MiningPeers must agree with the Status it reports.
func TestMiningPeersFollowsEtcdReadiness(t *testing.T) {
	ma := newTestEtcdAdmin(t)
	ma.lock = &sync.Mutex{}
	ma.self.Id = "self"
	ma.nodes = map[string]*wemixNode{ma.self.Id: ma.self}
	ma.blocksPer = 1

	prevAdmin := admin
	admin = ma
	t.Cleanup(func() { admin = prevAdmin })

	entry := func() string {
		_, _, nodes := ma.getMinerNodes(1, false)
		return ma.toMiningPeers(nodes)
	}
	if got, want := entry(), ma.self.Name+"/up"; !strings.HasPrefix(got, want) {
		t.Errorf("ready node: MiningPeers = %q, want prefix %q", got, want)
	}

	atomic.StoreUint32(&etcdReady, 0)
	if got, want := entry(), ma.self.Name+"/down"; !strings.HasPrefix(got, want) {
		t.Errorf("not ready node: MiningPeers = %q, want prefix %q", got, want)
	}
	atomic.StoreUint32(&etcdReady, 1)
}

// A rejection does not prove a removal. A member that clears its directory while
// the cluster still counts it panics on the leader's commit index, so it keeps
// its data and drops the marker. Before this check the test binary itself died.
func TestEtcdJoinKeepsDataOfLiveMember(t *testing.T) {
	survivor, victim := newTestEtcdCluster(t)

	ctx, cancel := context.WithTimeout(context.Background(), etcdStopTimeout)
	defer cancel()
	list, err := survivor.etcdCli.MemberList(ctx)
	if err != nil {
		t.Fatalf("MemberList failed: %v", err)
	}
	var cluster []string
	for _, m := range list.Members {
		cluster = append(cluster, fmt.Sprintf("%s=%s", m.Name, m.PeerURLs[0]))
	}
	serveEtcdJoin(t, victim, func() string { return strings.Join(cluster, ",") })

	id := uint64(victim.etcd.Server.ID())
	victim.etcd.Close()
	atomic.StoreUint32(&etcdRunning, 0)
	atomic.StoreUint32(&etcdReady, 0)
	victim.markRejoinNeeded(id)

	if err := victim.etcdJoin("peer"); !errors.Is(err, ErrNotRemoved) {
		t.Errorf("etcdJoin() = %v, want %v", err, ErrNotRemoved)
	}
	if victim.rejoinNeeded() {
		t.Error("a member that was not removed must lose the marker")
	}
	if matches, _ := filepath.Glob(filepath.Join(victim.etcdDir, "member", "wal", "*.wal")); len(matches) == 0 {
		t.Error("the member's WAL must survive")
	}
}

// With no peer to ask, nothing is known, so nothing may be cleared.
func TestEtcdRemovalConfirmedNeedsAnAnswer(t *testing.T) {
	ma := &wemixAdmin{
		self:    &wemixNode{Name: "self", Ip: "127.0.0.1", Port: 1},
		etcdDir: t.TempDir(),
	}
	if removed, err := ma.etcdRemovalConfirmed("self=https://127.0.0.1:2"); removed || err != nil {
		t.Errorf("no marker: etcdRemovalConfirmed() = %v, %v, want false, nil", removed, err)
	}
	ma.markRejoinNeeded(1)
	if removed, err := ma.etcdRemovalConfirmed("self=https://127.0.0.1:2"); removed || err == nil {
		t.Errorf("no peers: etcdRemovalConfirmed() = %v, %v, want false and an error", removed, err)
	}
}

// A removal never comes undone, so one peer that no longer lists the member is
// proof, however many others are behind. Trusting the first answer would turn a
// lagging peer into a failed rejoin.
func TestEtcdRemovalConfirmedByAnyPeer(t *testing.T) {
	peerURL, err := url.Parse("https://127.0.0.1:2")
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	removed := membership.NewMember("removed", []url.URL{*peerURL}, etcdClusterName, &now)
	other := membership.NewMember("other", []url.URL{*peerURL}, etcdClusterName, nil)

	serve := func(members ...*membership.Member) string {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("X-Etcd-Cluster-ID", "1")
			if err := json.NewEncoder(w).Encode(members); err != nil {
				t.Error(err)
			}
		}))
		t.Cleanup(server.Close)
		return server.URL
	}
	behind, current := serve(removed, other), serve(other)

	tests := []struct {
		name  string
		peers []string
		want  bool
	}{
		{"first peer is behind", []string{behind, current}, true},
		{"every peer still lists it", []string{behind, behind}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := &wemixAdmin{
				self:    &wemixNode{Name: "self", Ip: "127.0.0.1", Port: 1},
				etcdDir: t.TempDir(),
			}
			ma.markRejoinNeeded(uint64(removed.ID))
			var cluster []string
			for i, peer := range tt.peers {
				cluster = append(cluster, fmt.Sprintf("peer%d=%s", i, peer))
			}
			got, err := ma.etcdRemovalConfirmed(strings.Join(cluster, ","))
			if err != nil {
				t.Fatalf("etcdRemovalConfirmed failed: %v", err)
			}
			if got != tt.want {
				t.Errorf("etcdRemovalConfirmed() = %v, want %v", got, tt.want)
			}
		})
	}
}
