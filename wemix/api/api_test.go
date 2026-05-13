// Unit tests for wemix/api/api.go's WemixMinerStatus.Clone(), focused on
// nil-field safety.
//
// Background: a crafted StatusEx that omits a *big.Int field decodes to nil.
// The earlier Clone() implementation called new(big.Int).Set(nil) and
// panicked. A panic inside handleMinerStatusUpdate's goroutine would halt
// the miningPeers update channel.

package api

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// Clone() on a WemixMinerStatus whose *big.Int fields are nil must not panic
// and must preserve those fields as nil in the copy.
func TestWemixMinerStatus_Clone_NilSafe(t *testing.T) {
	src := &WemixMinerStatus{
		NodeName:          "node1",
		LatestBlockHeight: nil,
		LatestBlockHash:   common.HexToHash("0xdead"),
		LatestBlockTd:     nil,
		RttMs:             nil,
	}

	var panicked any
	var cloned *WemixMinerStatus
	func() {
		defer func() { panicked = recover() }()
		cloned = src.Clone()
	}()

	if panicked != nil {
		t.Fatalf("Clone() panicked on nil big.Int fields: %v", panicked)
	}
	if cloned == nil {
		t.Fatal("Clone() returned nil")
	}
	if cloned.LatestBlockHeight != nil {
		t.Errorf("LatestBlockHeight should remain nil, got %v", cloned.LatestBlockHeight)
	}
	if cloned.LatestBlockTd != nil {
		t.Errorf("LatestBlockTd should remain nil, got %v", cloned.LatestBlockTd)
	}
	if cloned.RttMs != nil {
		t.Errorf("RttMs should remain nil, got %v", cloned.RttMs)
	}
	if cloned.NodeName != "node1" || cloned.LatestBlockHash != common.HexToHash("0xdead") {
		t.Errorf("non-nil fields not preserved: %+v", cloned)
	}
}

// When all *big.Int fields are non-nil, Clone() must deep-copy them so that
// mutating the source after cloning does not affect the clone.
func TestWemixMinerStatus_Clone_DeepCopiesBigInts(t *testing.T) {
	srcHeight := big.NewInt(100)
	srcTd := big.NewInt(200)
	srcRtt := big.NewInt(50)
	src := &WemixMinerStatus{
		NodeName:          "node1",
		LatestBlockHeight: srcHeight,
		LatestBlockHash:   common.HexToHash("0xabcd"),
		LatestBlockTd:     srcTd,
		RttMs:             srcRtt,
	}

	cloned := src.Clone()
	if cloned == nil {
		t.Fatal("Clone() returned nil")
	}

	srcHeight.SetInt64(999)
	srcTd.SetInt64(999)
	srcRtt.SetInt64(999)

	if cloned.LatestBlockHeight.Int64() != 100 {
		t.Errorf("LatestBlockHeight aliased: got %v, want 100", cloned.LatestBlockHeight)
	}
	if cloned.LatestBlockTd.Int64() != 200 {
		t.Errorf("LatestBlockTd aliased: got %v, want 200", cloned.LatestBlockTd)
	}
	if cloned.RttMs.Int64() != 50 {
		t.Errorf("RttMs aliased: got %v, want 50", cloned.RttMs)
	}
}
