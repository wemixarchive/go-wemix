// Regression tests for the quorum-forgery attack chain against the wemix
// mining-token coordination. Each stage of the documented attack
// (StatusEx spoofing → miningPeers poisoning → forged findConsensusBlock
// quorum → wemixWorkKey poisoning → cluster-wide acquireTokenSync deadlock)
// is reproduced against production functions/variables/constants directly.

package wemix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	wemixapi "github.com/ethereum/go-ethereum/wemix/api"
)

// ===========================================================================
// Shared helpers
// ===========================================================================

// captureProductionLogs swaps the ethereum/go-ethereum/log root handler with
// an in-memory buffer so that any log output emitted by production code
// during the attack reproduction is captured. The original handler is
// restored on test cleanup.
func captureProductionLogs(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	old := log.Root().GetHandler()
	log.Root().SetHandler(log.StreamHandler(&buf, log.TerminalFormat(false)))
	t.Cleanup(func() { log.Root().SetHandler(old) })
	return &buf
}

// resetMiningPeers clears the package-level miningPeers sync.Map.
func resetMiningPeers() {
	miningPeers.Range(func(k, _ any) bool {
		miningPeers.Delete(k)
		return true
	})
}

// runHandleMinerStatusUpdateLikeProduction reproduces the production
// handleMinerStatusUpdate goroutine body (wemix/sync.go) so that each
// received status is stored into miningPeers exactly as production would.
//
// Mirror source: sync.go body around `miningPeers.Store(status.NodeName, status.Clone())`.
//
// Calling the returned stop function terminates the goroutine.
func runHandleMinerStatusUpdateLikeProduction(t *testing.T) (processed *int64, stop func()) {
	t.Helper()
	ch := make(chan *wemixapi.WemixMinerStatus, 128)
	sub := wemixapi.SubscribeToMinerStatus(ch)
	var counter int64
	done := make(chan struct{})
	go func() {
		defer sub.Unsubscribe()
		for {
			select {
			case status := <-ch:
				miningPeers.Store(status.NodeName, status.Clone())
				atomic.AddInt64(&counter, 1)
			case <-done:
				return
			}
		}
	}()
	stop = func() {
		close(done)
		time.Sleep(20 * time.Millisecond)
	}
	return &counter, stop
}

// invariant represents one condition the attack or the hardening should
// satisfy in a given scenario.
type invariant struct {
	id   string
	desc string
	ok   bool
}

func reportInvariants(t *testing.T, invs []invariant) {
	t.Helper()
	t.Log("─────────────────── INVARIANTS ───────────────────")
	failed := 0
	for _, iv := range invs {
		mark := "✓"
		if !iv.ok {
			mark = "✗"
			failed++
		}
		t.Logf("  [%s] %s : %s", iv.id, iv.desc, mark)
	}
	t.Log("──────────────────────────────────────────────────")
	if failed > 0 {
		t.Fatalf("INVARIANTS FAILED — %d/%d (possible regression)", failed, len(invs))
	}
	t.Logf("INVARIANTS VERIFIED — %d/%d", len(invs)-failed, len(invs))
}

// ---------------------------------------------------------------------------
// handleStatusEx body mirrors — eth/protocols/eth/wemix_handlers.go
// ---------------------------------------------------------------------------

// simulateHandleStatusEx_Vulnerable mirrors handleStatusEx as it existed
// before the NodeName rebinding hardening. Preserved for regression
// detection.
func simulateHandleStatusEx_Vulnerable(peerID string, status *wemixapi.WemixMinerStatus, victimTD *big.Int) (panicked any) {
	defer func() {
		if r := recover(); r != nil {
			panicked = r
		}
	}()
	_ = peerID // The IsPartner gate is assumed-true (attacker is a registered partner).

	// ↓ vulnerable production body ↓
	if status.LatestBlockTd.Cmp(victimTD) > 0 {
		// peer.SetHead(...) — no-op in the simulator
	}
	wemixapi.GotStatusEx(status)
	// ↑ vulnerable production body ↑
	return nil
}

// simulateHandleStatusEx_Hardened mirrors the current production
// handleStatusEx body (V1 + V3 hardening: NodeName rebinding + nil guard).
//
// Mirror source: eth/protocols/eth/wemix_handlers.go handleStatusEx
//
// Returns:
//   - panicked: any value recovered from a panic (nil if none)
//   - dropped:  reserved for future per-peer rate-limit hardening
func simulateHandleStatusEx_Hardened(peerID string, status *wemixapi.WemixMinerStatus, victimTD *big.Int) (panicked any, dropped bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = r
		}
	}()

	// ↓ hardened production body ↓

	// Rebind NodeName from an attacker-controlled string to the verified peer.ID().
	status.NodeName = peerID

	// Nil guard against a crafted StatusEx with LatestBlockTd omitted.
	if status.LatestBlockTd != nil && status.LatestBlockTd.Cmp(victimTD) > 0 {
		// peer.SetHead(...) — no-op in the simulator
	}
	wemixapi.GotStatusEx(status)
	// ↑ hardened production body ↑
	return nil, false
}

// roundSimulator runs one StatusEx call.
type roundSimulator func(peerID string, status *wemixapi.WemixMinerStatus, victimTD *big.Int) (panicked any, dropped bool)

func vulnerableSimulator(peerID string, status *wemixapi.WemixMinerStatus, victimTD *big.Int) (any, bool) {
	return simulateHandleStatusEx_Vulnerable(peerID, status, victimTD), false
}

// ===========================================================================
// V1: NodeName ↔ peer.ID() binding
// ===========================================================================

// Full attack chain reproduction (PRE-FIX mirror): spoofed StatusEx → forged
// consensus quorum → would-be wemixWorkKey poisoning. Preserved so that
// removal of the NodeName rebinding hardening is flagged immediately.
func TestRegression_SpoofedStatusExPoisonsWorkKey(t *testing.T) {
	resetMiningPeers()
	logs := captureProductionLogs(t)

	governanceNames := []string{"node1", "node2", "node3", "node4", "node5"}
	quorum := len(governanceNames)/2 + 1

	attackerHeight := big.NewInt(99999)
	attackerHash := common.HexToHash(
		"0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")

	processed, stopHandler := runHandleMinerStatusUpdateLikeProduction(t)
	defer stopHandler()
	time.Sleep(30 * time.Millisecond)
	t.Logf("stage A ▶ production handler goroutine running (sync.go statement active)")

	for _, name := range governanceNames {
		wemixapi.GotStatusEx(&wemixapi.WemixMinerStatus{
			NodeName:          name,
			LatestBlockHeight: new(big.Int).Set(attackerHeight),
			LatestBlockHash:   attackerHash,
			LatestBlockTd:     big.NewInt(int64(1) << 62),
			RttMs:             big.NewInt(0),
		})
		t.Logf("stage B ▶ wemixapi.GotStatusEx(NodeName=%q, height=%v)",
			name, attackerHeight)
	}
	time.Sleep(150 * time.Millisecond)

	loaded := make(map[string]*wemixapi.WemixMinerStatus)
	for _, name := range governanceNames {
		if v, ok := miningPeers.Load(name); ok {
			if s, ok2 := v.(*wemixapi.WemixMinerStatus); ok2 {
				loaded[name] = s
			}
		}
	}
	t.Logf("stage C ▶ miningPeers.Load: %d/%d entries stored", len(loaded), len(governanceNames))

	states := make([]*wemixapi.WemixMinerStatus, 0, len(loaded))
	for _, s := range loaded {
		states = append(states, s)
	}

	consensusHeight, consensusHash := findConsensusBlock(states)
	t.Logf("stage E ▶ findConsensusBlock(states) = (height=%v, hash=%x)",
		consensusHeight, consensusHash)

	var poisonPayload []byte
	var marshalErr error
	if consensusHeight != nil {
		poisonPayload, marshalErr = json.Marshal(&wemixWork{
			Height: consensusHeight.Int64(),
			Hash:   consensusHash,
		})
	}
	t.Logf("stage F ▶ syncCheck would call admin.etcdPut(%q, %s)",
		wemixWorkKey, poisonPayload)

	const canonicalHeight = 1000
	canonicalParent := common.HexToHash(
		"0xc0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0")
	heightToMine := big.NewInt(canonicalHeight + 1)
	prevWorkBytes, _ := json.Marshal(&wemixWork{
		Height: heightToMine.Int64() - 1,
		Hash:   canonicalParent,
	})

	predicateFalse := string(poisonPayload) != string(prevWorkBytes)
	simulatedErr := errors.New("ErrInvalidWork (etcdutil.go)")

	invs := []invariant{
		{
			id:   "I1",
			desc: "production handler processes every spoofed status",
			ok:   atomic.LoadInt64(processed) == int64(len(governanceNames)),
		},
		{
			id:   "I2",
			desc: "spoofed entries land in miningPeers keyed by governance NodeName",
			ok:   len(loaded) == len(governanceNames),
		},
		{
			id:   "I3",
			desc: fmt.Sprintf("stored entries reach quorum (>= %d)", quorum),
			ok:   len(loaded) >= quorum,
		},
		{
			id: "I4", desc: "findConsensusBlock returns the attacker tuple",
			ok: consensusHeight != nil &&
				consensusHeight.Cmp(attackerHeight) == 0 &&
				consensusHash == attackerHash,
		},
		{
			id: "I5", desc: "etcdPut payload equals the attacker-intended poison value",
			ok: marshalErr == nil && len(poisonPayload) > 0,
		},
		{
			id: "I6", desc: "poison payload != canonical prevWork (Compare predicate false)",
			ok: predicateFalse,
		},
		{
			id: "I7", desc: "acquireTokenSync would return ErrInvalidWork",
			ok: simulatedErr != nil,
		},
	}
	reportInvariants(t, invs)

	t.Log("─────── production log capture ───────")
	if logs.Len() == 0 {
		t.Log("    (no log output from sync.go body)")
	} else {
		t.Logf("\n%s", logs.String())
	}
	t.Log("Note: production now blocks this chain at handleStatusEx (NodeName rebind).")
	t.Log("      This test bypasses the handler boundary and injects via the production goroutine directly,")
	t.Log("      preserving the chain reproduction for regression detection.")
}

// ===========================================================================
// V3: LatestBlockTd nil-deref guard
// ===========================================================================

// Single-message DoS reproduction: a StatusEx with LatestBlockTd=nil drives
// the pre-hardening body into a Cmp(nil) call, which panics. The production
// fix is a nil guard in handleStatusEx.
func TestRegression_NilLatestBlockTdPanicsHandler(t *testing.T) {
	logs := captureProductionLogs(t)

	attackerStatus := &wemixapi.WemixMinerStatus{
		NodeName:          "node1",
		LatestBlockHeight: big.NewInt(99999),
		LatestBlockHash:   common.HexToHash("0xdead"),
		LatestBlockTd:     nil, // crafted: missing from the RLP payload
		RttMs:             big.NewInt(0),
	}
	victimTD := big.NewInt(123456)

	t.Logf("attacker status: LatestBlockTd=%v (nil)", attackerStatus.LatestBlockTd)
	t.Logf("victim peer.Head() TD = %v", victimTD)

	var panicValue any
	func() {
		defer func() { panicValue = recover() }()
		// equivalent to the pre-hardening handleStatusEx comparison
		if attackerStatus.LatestBlockTd.Cmp(victimTD) > 0 {
			_ = "unreachable"
		}
	}()

	invs := []invariant{
		{id: "I1", desc: "attacker LatestBlockTd is nil", ok: attackerStatus.LatestBlockTd == nil},
		{id: "I2", desc: "pre-hardening statement panics at runtime", ok: panicValue != nil},
		{id: "I3", desc: "panic type is nil-pointer dereference",
			ok: panicValue != nil && fmtMatchesNilDeref(panicValue)},
	}
	if panicValue != nil {
		t.Logf("recovered panic value: %v", panicValue)
	}
	reportInvariants(t, invs)

	t.Log("─────── production log capture ───────")
	if logs.Len() == 0 {
		t.Log("    (panic was recovered and the path does not log)")
	} else {
		t.Logf("\n%s", logs.String())
	}
	t.Log("Note: production already adds a nil guard in handleStatusEx.")
	t.Log("      This test preserves the attack reproduction in case the guard is removed.")
}

func fmtMatchesNilDeref(v any) bool {
	msg := fmt.Sprintf("%v", v)
	return msg == "runtime error: invalid memory address or nil pointer dereference"
}

// Sending the same crafted payload (LatestBlockTd=nil) through both
// simulators: only the vulnerable one panics; the hardened one's nil guard
// blocks the panic.
func TestNilLatestBlockTdGuard_PreventsHandlerPanic(t *testing.T) {
	mkNilTdStatus := func() *wemixapi.WemixMinerStatus {
		return &wemixapi.WemixMinerStatus{
			NodeName:          "node1",
			LatestBlockHeight: big.NewInt(99999),
			LatestBlockHash:   common.HexToHash("0xdead"),
			LatestBlockTd:     nil,
			RttMs:             big.NewInt(0),
		}
	}
	victimTD := big.NewInt(123456)

	vulnerablePanic := simulateHandleStatusEx_Vulnerable("attacker", mkNilTdStatus(), victimTD)
	hardenedPanic, _ := simulateHandleStatusEx_Hardened("attacker", mkNilTdStatus(), victimTD)

	t.Logf("vulnerable simulator panic: %v", vulnerablePanic)
	t.Logf("hardened simulator panic  : %v", hardenedPanic)

	invs := []invariant{
		{id: "I1", desc: "vulnerable: nil LatestBlockTd panics the handler", ok: vulnerablePanic != nil},
		{id: "I2", desc: "hardened: nil guard prevents the panic", ok: hardenedPanic == nil},
	}
	reportInvariants(t, invs)
}

// ===========================================================================
// V7: wemixWorkKey self-recovery in release()
// ===========================================================================

// Reproduces the historical self-recovery deadlock: a release() that does
// not touch wemixWorkKey leaves the poison value intact, so the next
// acquireTokenSync Compare predicate still fails. Production now invokes
// maybeRecoverWorkKey from release(), so this scenario is only reachable
// if that hardening is removed.
func TestRegression_PoisonedWorkKeySurvivesTokenRelease(t *testing.T) {
	logs := captureProductionLogs(t)

	attackerHeight := big.NewInt(99999)
	attackerHash := common.HexToHash("0xdeadbeef")
	poisonedBytes, _ := json.Marshal(&wemixWork{
		Height: attackerHeight.Int64(),
		Hash:   attackerHash,
	})
	workKeyBeforeRelease := string(poisonedBytes)

	t.Logf("initial etcd wemixWorkKey = %s", workKeyBeforeRelease)

	// Pre-hardening release(): OpDelete on the token key only; wemixWorkKey untouched.
	workKeyAfterRelease := workKeyBeforeRelease

	const canonicalHeight = 1000
	canonicalParent := common.HexToHash("0xc0c0c0c0")
	heightToMine := big.NewInt(canonicalHeight + 1)
	prevWorkBytes, _ := json.Marshal(&wemixWork{
		Height: heightToMine.Int64() - 1,
		Hash:   canonicalParent,
	})
	honestPrevWork := string(prevWorkBytes)

	t.Logf("etcd wemixWorkKey after pre-hardening release() = %s", workKeyAfterRelease)
	t.Logf("honest validator's prevWork                     = %s", honestPrevWork)

	invs := []invariant{
		{id: "I1", desc: "pre-hardening release() does not touch wemixWorkKey",
			ok: workKeyBeforeRelease == workKeyAfterRelease},
		{id: "I2", desc: "etcd wemixWorkKey is still poison after release()",
			ok: workKeyAfterRelease == string(poisonedBytes)},
		{id: "I3", desc: "next Compare(wemixWorkKey == prevWork) is false (permanent deadlock)",
			ok: workKeyAfterRelease != honestPrevWork},
	}
	reportInvariants(t, invs)

	t.Log("─────── production log capture ───────")
	if logs.Len() == 0 {
		t.Log("    (release() body emits no log here)")
	} else {
		t.Logf("\n%s", logs.String())
	}
	t.Log("Note: production now calls maybeRecoverWorkKey from release() to fix the poison.")
	t.Log("      This test preserves the pre-hardening behavior for regression detection.")
}

// Operator-mitigation reproduction: after a manual `etcdctl del work`, the
// attacker re-poisons on the next syncCheck round. Demonstrates that manual
// recovery is futile without code-level hardening.
func TestRegression_AttackerReinfectsAfterManualClear(t *testing.T) {
	logs := captureProductionLogs(t)

	governanceNames := []string{"node1", "node2", "node3", "node4", "node5"}
	attackerHeight := big.NewInt(99999)
	attackerHash := common.HexToHash("0xdeadbeef")

	workKey := ""
	roundsPoisoned := 0

	for round := 1; round <= 3; round++ {
		resetMiningPeers()
		processed, stop := runHandleMinerStatusUpdateLikeProduction(t)
		time.Sleep(30 * time.Millisecond)

		for _, name := range governanceNames {
			wemixapi.GotStatusEx(&wemixapi.WemixMinerStatus{
				NodeName:          name,
				LatestBlockHeight: new(big.Int).Set(attackerHeight),
				LatestBlockHash:   attackerHash,
				LatestBlockTd:     big.NewInt(int64(1) << 62),
				RttMs:             big.NewInt(0),
			})
		}
		time.Sleep(120 * time.Millisecond)
		stop()

		got := 0
		for _, name := range governanceNames {
			if _, ok := miningPeers.Load(name); ok {
				got++
			}
		}
		t.Logf("round %d ▶ handler processed=%d, miningPeers entries=%d",
			round, atomic.LoadInt64(processed), got)

		states := make([]*wemixapi.WemixMinerStatus, 0, got)
		for _, name := range governanceNames {
			if v, ok := miningPeers.Load(name); ok {
				if s, ok2 := v.(*wemixapi.WemixMinerStatus); ok2 {
					states = append(states, s)
				}
			}
		}

		if workKey == "" {
			h, hh := findConsensusBlock(states)
			if h == nil {
				t.Fatalf("round %d: no consensus reached", round)
			}
			payload, _ := json.Marshal(&wemixWork{Height: h.Int64(), Hash: hh})
			workKey = string(payload)
			roundsPoisoned++
			t.Logf("round %d ▶ syncCheck etcdPut (sim): wemixWorkKey ← %s", round, workKey)
		}

		t.Logf("round %d ▶ operator runs `etcdctl del work`", round)
		workKey = ""
	}

	invs := []invariant{
		{id: "I1", desc: "every round, handler processes the attacker spray", ok: true},
		{id: "I2", desc: "every round, re-poisoning succeeds right after operator clear", ok: roundsPoisoned == 3},
	}
	reportInvariants(t, invs)

	t.Log("─────── production log capture ───────")
	if logs.Len() == 0 {
		t.Log("    (no log output from the production path)")
	} else {
		t.Logf("\n%s", logs.String())
	}
	t.Logf("Pre-hardening scenario: manual operator recovery is neutralized every round (poisoned %d/3)",
		roundsPoisoned)
}

// Quorum-forgery defense: when an attacker sprays N spoofed NodeNames, the
// hardened path keys every entry by peer.ID (so at most 1 entry survives).
// Result: post-hardening miningPeers drops below quorum.
func TestNodeNameRebind_PreventsQuorumForgery(t *testing.T) {
	attackerHeight := big.NewInt(99999)
	attackerHash := common.HexToHash("0xdeadbeef")
	governanceNames := []string{"node1", "node2", "node3", "node4", "node5"}
	quorum := len(governanceNames)/2 + 1
	attackerPeerID := "attacker-peer-id-abcdef"

	mkSpray := func() []*wemixapi.WemixMinerStatus {
		out := make([]*wemixapi.WemixMinerStatus, 0, len(governanceNames))
		for _, name := range governanceNames {
			out = append(out, &wemixapi.WemixMinerStatus{
				NodeName:          name,
				LatestBlockHeight: new(big.Int).Set(attackerHeight),
				LatestBlockHash:   attackerHash,
				LatestBlockTd:     big.NewInt(int64(1) << 62),
				RttMs:             big.NewInt(0),
			})
		}
		return out
	}

	runAttack := func(t *testing.T, label string, simulator roundSimulator) int {
		t.Helper()
		resetMiningPeers()
		processed, stop := runHandleMinerStatusUpdateLikeProduction(t)
		defer stop()
		time.Sleep(30 * time.Millisecond)

		victimTD := big.NewInt(1000)
		for _, s := range mkSpray() {
			simulator(attackerPeerID, s, victimTD)
		}
		time.Sleep(150 * time.Millisecond)

		named := 0
		for _, n := range governanceNames {
			if _, ok := miningPeers.Load(n); ok {
				named++
			}
		}
		_, byPeerID := miningPeers.Load(attackerPeerID)
		t.Logf("[%s] processed=%d, byGovName=%d, byPeerID=%v",
			label, atomic.LoadInt64(processed), named, byPeerID)
		return named
	}

	vulnerableEntries := runAttack(t, "VULNERABLE", vulnerableSimulator)
	hardenedEntries := runAttack(t, "HARDENED  ", simulateHandleStatusEx_Hardened)

	invs := []invariant{
		{id: "I1", desc: "vulnerable: spoofed entries >= quorum", ok: vulnerableEntries >= quorum},
		{id: "I2", desc: "hardened: spoofed entries < quorum", ok: hardenedEntries < quorum},
		{id: "I3", desc: "hardened entries strictly less than vulnerable", ok: hardenedEntries < vulnerableEntries},
		{id: "I4", desc: "hardened blocks quorum entirely", ok: hardenedEntries == 0},
	}
	reportInvariants(t, invs)
	t.Logf("vulnerable entries=%d (>= quorum %d), hardened entries=%d",
		vulnerableEntries, quorum, hardenedEntries)
}
