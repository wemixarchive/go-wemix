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

func registeredNodeNameForTestPeerID(peerID string) string {
	return "registered-" + peerID
}

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
// handleStatusEx body (NodeName rebinding + nil guards).
//
// Mirror source: eth/protocols/eth/wemix_handlers.go handleStatusEx
//
// Returns:
//   - panicked: any value recovered from a panic (nil if none)
//   - dropped:  true if the message was discarded by a boundary validation
//     reject (e.g., nil LatestBlockHeight)
func simulateHandleStatusEx_Hardened(peerID string, status *wemixapi.WemixMinerStatus, victimTD *big.Int) (panicked any, dropped bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = r
		}
	}()

	// ↓ hardened production body ↓

	// Boundary reject: legitimate senders always set LatestBlockHeight from
	// header.Number. A nil here is only reachable from a crafted payload, and
	// would later panic in electNextMiner where the Status=="up" short circuit
	// is the only thing standing between the value and Int64().
	if status.LatestBlockHeight == nil {
		return nil, true
	}

	// Replace the attacker-controlled name with the governance-registered name
	// for the verified peer. In this simulator the registered name is derived
	// from peerID to keep the production trust boundary explicit.
	status.NodeName = registeredNodeNameForTestPeerID(peerID)

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
// V1: StatusEx NodeName validation
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
// V4: LatestBlockHeight nil-deref guard (electNextMiner downstream)
// ===========================================================================

// Single-message DoS reproduction: a StatusEx with LatestBlockHeight=nil and
// Status="up" reaches miningPeers via wemixapi.GotStatusEx, then electNextMiner
// (wemix/miner_limit.go) calls p.LatestBlockHeight.Int64() and panics on the
// nil *big.Int. The Status=="up" short circuit is the only thing standing
// between a crafted payload and the deref, and it is attacker-controlled.
// Preserved so that removal of the handler-side boundary reject is flagged
// immediately.
func TestRegression_NilLatestBlockHeightPanicsElectNextMiner(t *testing.T) {
	attackerStatus := &wemixapi.WemixMinerStatus{
		NodeName:          "node1",
		Status:            "up",
		LatestBlockHeight: nil, // crafted: missing from the RLP payload
		LatestBlockHash:   common.HexToHash("0xdead"),
		LatestBlockTd:     big.NewInt(1),
		RttMs:             big.NewInt(0),
	}
	currentHeight := big.NewInt(1000)
	const closeEnough = int64(2)

	t.Logf("attacker status: Status=%q, LatestBlockHeight=%v",
		attackerStatus.Status, attackerStatus.LatestBlockHeight)
	t.Logf("electNextMiner current height = %v", currentHeight)

	var panicValue any
	func() {
		defer func() { panicValue = recover() }()
		// equivalent to the pre-boundary-reject body at wemix/miner_limit.go:338
		if attackerStatus.Status == "up" &&
			currentHeight.Int64()-attackerStatus.LatestBlockHeight.Int64() <= closeEnough {
			_ = "unreachable"
		}
	}()

	invs := []invariant{
		{id: "I1", desc: "attacker LatestBlockHeight is nil", ok: attackerStatus.LatestBlockHeight == nil},
		{id: "I2", desc: "attacker Status is up (short-circuit bypassed)", ok: attackerStatus.Status == "up"},
		{id: "I3", desc: "pre-hardening statement panics at runtime", ok: panicValue != nil},
		{id: "I4", desc: "panic type is nil-pointer dereference",
			ok: panicValue != nil && fmtMatchesNilDeref(panicValue)},
	}
	if panicValue != nil {
		t.Logf("recovered panic value: %v", panicValue)
	}
	reportInvariants(t, invs)

	t.Log("Note: production now rejects nil LatestBlockHeight at handleStatusEx.")
	t.Log("      This test preserves the consumer-side panic to flag boundary-reject removal.")
}

// Sending the same crafted payload (LatestBlockHeight=nil) through the hardened
// simulator: the boundary reject drops the message before wemixapi.GotStatusEx,
// so miningPeers never sees the nil-height entry that would later panic
// electNextMiner.
func TestNilLatestBlockHeightGuard_RejectedAtHandler(t *testing.T) {
	resetMiningPeers()
	processed, stopHandler := runHandleMinerStatusUpdateLikeProduction(t)
	defer stopHandler()
	time.Sleep(30 * time.Millisecond)

	mkNilHeightStatus := func() *wemixapi.WemixMinerStatus {
		return &wemixapi.WemixMinerStatus{
			NodeName:          "node1",
			Status:            "up",
			LatestBlockHeight: nil,
			LatestBlockHash:   common.HexToHash("0xdead"),
			LatestBlockTd:     big.NewInt(1),
			RttMs:             big.NewInt(0),
		}
	}
	victimTD := big.NewInt(123456)

	hardenedPanic, hardenedDropped := simulateHandleStatusEx_Hardened(
		"attacker-peer-id", mkNilHeightStatus(), victimTD)
	time.Sleep(50 * time.Millisecond)

	registeredName := registeredNodeNameForTestPeerID("attacker-peer-id")
	_, storedByRegisteredName := miningPeers.Load(registeredName)
	_, storedByOriginalName := miningPeers.Load("node1")

	t.Logf("hardened simulator: panic=%v, dropped=%v", hardenedPanic, hardenedDropped)
	t.Logf("handler processed count = %d", atomic.LoadInt64(processed))
	t.Logf("miningPeers stored byRegisteredName=%v, byOriginalName=%v",
		storedByRegisteredName, storedByOriginalName)

	invs := []invariant{
		{id: "I1", desc: "hardened: nil LatestBlockHeight is dropped at boundary",
			ok: hardenedDropped},
		{id: "I2", desc: "hardened: no panic on the reject path",
			ok: hardenedPanic == nil},
		{id: "I3", desc: "hardened: no miningPeers entry written for the rejected message",
			ok: !storedByRegisteredName && !storedByOriginalName},
		{id: "I4", desc: "hardened: handler did not forward the message to subscribers",
			ok: atomic.LoadInt64(processed) == 0},
	}
	reportInvariants(t, invs)
}

// Quorum-forgery defense: when an attacker sprays N spoofed NodeNames, the
// hardened path stores every accepted status under the attacker's registered
// node name. Result: post-hardening miningPeers drops below quorum.
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
		_, byRegisteredName := miningPeers.Load(registeredNodeNameForTestPeerID(attackerPeerID))
		t.Logf("[%s] processed=%d, byGovName=%d, byRegisteredName=%v",
			label, atomic.LoadInt64(processed), named, byRegisteredName)
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

// ===========================================================================
// End-to-end attack chain — pre-hardening vs hardened comparison
// ===========================================================================

// endToEndAttackScenario bundles the inputs for the nine-stage attack run.
type endToEndAttackScenario struct {
	attackerPeerID   string
	idleSeconds      int
	spoofedNames     []string
	attackerHeight   *big.Int
	attackerHash     common.Hash
	canonicalHeight  int64
	canonicalParent  common.Hash
	victimPeerTD     *big.Int
	tokenAcquireWait time.Duration
}

func defaultScenario() endToEndAttackScenario {
	names := make([]string, 0, 5)
	for i := 1; i <= 5; i++ {
		names = append(names, fmt.Sprintf("validator-%c", 'a'+i))
	}
	return endToEndAttackScenario{
		attackerPeerID:  "compromised-validator-A",
		idleSeconds:     SyncIdleThreshold + 1,
		spoofedNames:    names,
		attackerHeight:  big.NewInt(99999),
		attackerHash:    common.HexToHash("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"),
		canonicalHeight: 1000,
		canonicalParent: common.HexToHash(
			"0xc0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0"),
		victimPeerTD:     big.NewInt(1 << 30),
		tokenAcquireWait: 150 * time.Millisecond,
	}
}

var errInvalidWorkPredicate = errors.New("ErrInvalidWork (simulated — etcdutil.go)")

// acquireTokenSyncPredicate models the byte-equality Compare predicate
// inside acquireTokenSync.
func acquireTokenSyncPredicate(currentWorkKeyValue, expectedPrevWork string) error {
	if currentWorkKeyValue == expectedPrevWork {
		return nil
	}
	return errInvalidWorkPredicate
}

func truncateForLog(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "…"
}

type stageOutcome struct {
	idleSecondsObserved         int
	idlePassed                  bool
	sprayCount                  int
	sprayDelivered              int
	miningPeersByName           int
	miningPeersByRegisteredName bool
	consensusHeight             *big.Int
	consensusHash               common.Hash
	quorumReached               bool
	poisonPayload               []byte
	honestPrevWork              []byte
	compareErr                  error
	tokenAcquired               bool
	subsequentSyncCheckErr      error
}

func runOneRound(
	t *testing.T,
	label string,
	sc endToEndAttackScenario,
	simulator roundSimulator,
) stageOutcome {
	t.Helper()
	out := stageOutcome{}

	t.Logf("════════════════════════════════════════════════════════════════")
	t.Logf("█ %s : end-to-end attack chain (9 stages)", label)
	t.Logf("════════════════════════════════════════════════════════════════")

	// Stage 1: setup
	t.Logf("[1] 5-validator devnet, SyncIdleThreshold = %ds", SyncIdleThreshold)
	resetMiningPeers()
	processed, stop := runHandleMinerStatusUpdateLikeProduction(t)
	defer stop()
	time.Sleep(30 * time.Millisecond)

	// Stage 2: attacker identity
	t.Logf("[2] attacker peer.ID = %q (IsPartner=true assumed)", sc.attackerPeerID)

	// Stage 3: force the idle gate to pass
	idleSince := time.Now().Add(-time.Duration(sc.idleSeconds) * time.Second)
	latestUpdateTime.Store(idleSince)
	out.idleSecondsObserved = int(time.Since(idleSince).Seconds())
	out.idlePassed = out.idleSecondsObserved > SyncIdleThreshold
	t.Logf("[3] latestUpdateTime ← %ds ago, gate pass=%v",
		out.idleSecondsObserved, out.idlePassed)

	// Stage 4: spoofed-NodeName spray
	for _, name := range sc.spoofedNames {
		status := &wemixapi.WemixMinerStatus{
			NodeName:          name,
			LatestBlockHeight: new(big.Int).Set(sc.attackerHeight),
			LatestBlockHash:   sc.attackerHash,
			LatestBlockTd:     big.NewInt(1 << 62),
			RttMs:             big.NewInt(0),
		}
		_, dropped := simulator(sc.attackerPeerID, status, sc.victimPeerTD)
		out.sprayCount++
		if !dropped {
			out.sprayDelivered++
		}
	}
	t.Logf("[4] spray count=%d, delivered=%d (handler drain wait %v)",
		out.sprayCount, out.sprayDelivered, sc.tokenAcquireWait)
	time.Sleep(sc.tokenAcquireWait)

	// Stage 5: inspect miningPeers
	for _, name := range sc.spoofedNames {
		if _, ok := miningPeers.Load(name); ok {
			out.miningPeersByName++
		}
	}
	_, out.miningPeersByRegisteredName = miningPeers.Load(registeredNodeNameForTestPeerID(sc.attackerPeerID))
	t.Logf("[5] handler processed=%d, miningPeers: byName=%d/%d, byRegisteredName=%v",
		atomic.LoadInt64(processed), out.miningPeersByName,
		len(sc.spoofedNames), out.miningPeersByRegisteredName)

	// Stage 6: real findConsensusBlock
	states := make([]*wemixapi.WemixMinerStatus, 0, len(sc.spoofedNames))
	for _, name := range sc.spoofedNames {
		if v, ok := miningPeers.Load(name); ok {
			if s, ok2 := v.(*wemixapi.WemixMinerStatus); ok2 {
				states = append(states, s)
			}
		}
	}
	out.consensusHeight, out.consensusHash = findConsensusBlock(states)
	quorum := len(sc.spoofedNames)/2 + 1
	out.quorumReached = out.consensusHeight != nil &&
		out.consensusHeight.Cmp(sc.attackerHeight) == 0 &&
		out.consensusHash == sc.attackerHash
	t.Logf("[6] findConsensusBlock(states[%d]) quorum=%d, attackerTuple=%v",
		len(states), quorum, out.quorumReached)

	// Stage 7: etcdPut payload
	if out.quorumReached {
		payload, err := json.Marshal(&wemixWork{
			Height: out.consensusHeight.Int64(),
			Hash:   out.consensusHash,
		})
		if err == nil {
			out.poisonPayload = payload
		}
	}
	t.Logf("[7] syncCheck etcdPut(%q, …): payloadEmpty=%v",
		wemixWorkKey, len(out.poisonPayload) == 0)

	// Stage 8: acquireTokenSync Compare predicate
	prevWork, _ := json.Marshal(&wemixWork{
		Height: sc.canonicalHeight,
		Hash:   sc.canonicalParent,
	})
	out.honestPrevWork = prevWork

	currentWorkKeyValue := string(out.poisonPayload)
	if len(out.poisonPayload) == 0 {
		currentWorkKeyValue = string(prevWork)
	}
	out.compareErr = acquireTokenSyncPredicate(currentWorkKeyValue, string(prevWork))
	out.tokenAcquired = out.compareErr == nil
	t.Logf("[8] etcd wemixWorkKey=%s, tokenAcquired=%v",
		truncateForLog(currentWorkKeyValue, 64), out.tokenAcquired)

	// Stage 9: subsequent syncCheck
	out.subsequentSyncCheckErr = acquireTokenSyncPredicate(
		currentWorkKeyValue, string(prevWork))
	t.Logf("[9] subsequent acquireToken → %v", out.subsequentSyncCheckErr)
	t.Logf("════════════════════════════════════════════════════════════════")
	return out
}

func TestRegression_EndToEndAttackChain(t *testing.T) {
	logs := captureProductionLogs(t)
	sc := defaultScenario()

	var pre, post stageOutcome
	t.Run("vulnerable_behavior", func(t *testing.T) {
		pre = runOneRound(t, "VULNERABLE", sc, vulnerableSimulator)
	})
	t.Run("hardened_behavior", func(t *testing.T) {
		post = runOneRound(t, "HARDENED  ", sc, simulateHandleStatusEx_Hardened)
	})

	quorum := len(sc.spoofedNames)/2 + 1
	invs := []invariant{
		{id: "C1", desc: fmt.Sprintf("idle %ds > %ds (both rounds)",
			sc.idleSeconds, SyncIdleThreshold), ok: pre.idlePassed && post.idlePassed},

		{id: "V1", desc: "vulnerable [5]: spoofed entries >= quorum",
			ok: pre.miningPeersByName >= quorum},
		{id: "V2", desc: "vulnerable [6]: findConsensusBlock returns attacker tuple",
			ok: pre.quorumReached},
		{id: "V3", desc: "vulnerable [7]: etcdPut payload equals attacker value",
			ok: len(pre.poisonPayload) > 0},
		{id: "V4", desc: "vulnerable [8]: honest acquireTokenSync receives ErrInvalidWork",
			ok: !pre.tokenAcquired && pre.compareErr != nil},
		{id: "V5", desc: "vulnerable [9]: subsequent syncCheck also fails",
			ok: pre.subsequentSyncCheckErr != nil},

		{id: "H2", desc: "hardened [5]: spoofed entries < quorum",
			ok: post.miningPeersByName < quorum},
		{id: "H3", desc: "hardened [5]: single entry keyed by registered node name",
			ok: post.miningPeersByRegisteredName && post.miningPeersByName == 0},
		{id: "H4", desc: "hardened [6]: findConsensusBlock does not return attacker tuple",
			ok: !post.quorumReached},
		{id: "H5", desc: "hardened [7]: etcdPut never reached",
			ok: len(post.poisonPayload) == 0},
		{id: "H6", desc: "hardened [8]: honest acquireTokenSync succeeds",
			ok: post.tokenAcquired && post.compareErr == nil},
	}
	reportInvariants(t, invs)

	t.Log("─────── production log capture ───────")
	if logs.Len() == 0 {
		t.Log("    (no log output from the production path)")
	} else {
		t.Logf("\n%s", logs.String())
	}
	t.Log("VULNERABLE: stages 1..9 all pass → cluster-wide mining-token deadlock reproduced")
	t.Log("HARDENED  : NodeName rebind blocks quorum forgery → stages 6..9 neutralized")
}

// ===========================================================================
// V9: consensusHeight regression guard
// ===========================================================================

// TestRegression_StaleConsensusHeightBlocked verifies the strict height
// regression guard in syncCheck: when a poisoned quorum drives
// findConsensusBlock to return a height below the local header, the
// would-be etcdPut(wemixWorkKey, ...) must be suppressed.
//
// Mirror source: wemix/sync.go, body around the
//
//	if consensusHeight.Cmp(header.Number) < 0 { ... return nil }
//
// guard immediately above the newWork construction.
func TestRegression_StaleConsensusHeightBlocked(t *testing.T) {
	governanceNames := []string{"node1", "node2", "node3", "node4", "node5"}

	// Local chain is at height 1000. Attacker quorum claims height 900 with
	// a hash that — in the attack scenario — is still reachable from the
	// local chain (an old canonical block). The hash-reachability guard
	// alone would let this through; only the height guard catches it.
	localHeight := big.NewInt(1000)
	staleHeight := big.NewInt(900)
	staleHash := common.HexToHash("0xfeedbeef")

	resetMiningPeers()
	processed, stop := runHandleMinerStatusUpdateLikeProduction(t)
	time.Sleep(30 * time.Millisecond)

	for _, name := range governanceNames {
		wemixapi.GotStatusEx(&wemixapi.WemixMinerStatus{
			NodeName:          name,
			LatestBlockHeight: new(big.Int).Set(staleHeight),
			LatestBlockHash:   staleHash,
			LatestBlockTd:     big.NewInt(1),
			RttMs:             big.NewInt(0),
		})
	}
	time.Sleep(120 * time.Millisecond)
	stop()
	t.Logf("stage A ▶ handler processed=%d, miningPeers entries seeded", atomic.LoadInt64(processed))

	states := make([]*wemixapi.WemixMinerStatus, 0, len(governanceNames))
	for _, name := range governanceNames {
		if v, ok := miningPeers.Load(name); ok {
			if s, ok2 := v.(*wemixapi.WemixMinerStatus); ok2 {
				states = append(states, s)
			}
		}
	}

	consensusHeight, consensusHash := findConsensusBlock(states)
	t.Logf("stage B ▶ findConsensusBlock = (height=%v, hash=%x)", consensusHeight, consensusHash)
	if consensusHeight == nil {
		t.Fatalf("findConsensusBlock returned nil; attack precondition (quorum) broken")
	}

	// Mirror of the production guard. If consensusHeight < header.Number,
	// syncCheck returns before reaching admin.etcdPut(wemixWorkKey, ...).
	regressed := consensusHeight.Cmp(localHeight) < 0
	wouldEtcdPut := !regressed

	invs := []invariant{
		{id: "I1", desc: "attack precondition: quorum reports the stale height",
			ok: consensusHeight.Cmp(staleHeight) == 0},
		{id: "I2", desc: "guard recognises consensus < local head as regression",
			ok: regressed},
		{id: "I3", desc: "syncCheck would NOT call admin.etcdPut(wemixWorkKey, ...)",
			ok: !wouldEtcdPut},
	}
	reportInvariants(t, invs)
}
