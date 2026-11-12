package core

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

// croissantForkBlock is where the test chain activates Croissant.
const croissantForkBlock = 10

// TestIntrinsicGasEIP3860 checks that the per-word initcode charge applies to
// contract creation from the fork on, and never to ordinary calls.
func TestIntrinsicGasEIP3860(t *testing.T) {
	const words = 2
	data := make([]byte, words*32)
	wantDelta := words * params.InitCodeWordGas

	for _, tt := range []struct {
		name     string
		creation bool
		delta    uint64
	}{
		{"creation", true, wantDelta},
		{"call", false, 0},
	} {
		pre, err := IntrinsicGas(data, nil, tt.creation, true, true, false)
		if err != nil {
			t.Fatalf("%s before the fork: %v", tt.name, err)
		}
		post, err := IntrinsicGas(data, nil, tt.creation, true, true, true)
		if err != nil {
			t.Fatalf("%s after the fork: %v", tt.name, err)
		}
		if post-pre != tt.delta {
			t.Errorf("%s: Croissant adds %d gas, want %d", tt.name, post-pre, tt.delta)
		}
	}
}

// applyCreation runs a contract creation message with the given initcode at
// block number and returns what ApplyMessage reports.
func applyCreation(t *testing.T, block int64, initcode []byte) (*ExecutionResult, error) {
	t.Helper()
	cfg := *params.AllEthashProtocolChanges
	cfg.CroissantBlock = big.NewInt(croissantForkBlock)

	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatal(err)
	}
	const gasLimit = 30_000_000
	blockCtx := vm.BlockContext{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		BlockNumber: big.NewInt(block),
		Time:        big.NewInt(0),
		Difficulty:  big.NewInt(1),
		GasLimit:    gasLimit,
		BaseFee:     big.NewInt(0),
	}
	evm := vm.NewEVM(blockCtx, vm.TxContext{GasPrice: big.NewInt(0)}, statedb, &cfg, vm.Config{NoBaseFee: true})

	from := common.HexToAddress("0x1000000000000000000000000000000000000001")
	zero := big.NewInt(0)
	msg := types.NewMessage(from, nil, 0, zero, gasLimit, zero, zero, zero, initcode, nil, true)
	return ApplyMessage(evm, msg, new(GasPool).AddGas(math.MaxUint64))
}

// TestEIP3860CreationTxRejected checks that a creation transaction with
// initcode over the limit is invalid from the fork on: ApplyMessage returns a
// consensus error instead of including it as a failed execution.
func TestEIP3860CreationTxRejected(t *testing.T) {
	oversized := make([]byte, params.MaxInitCodeSize+1)

	if _, err := applyCreation(t, croissantForkBlock, oversized); !errors.Is(err, ErrMaxInitCodeSizeExceeded) {
		t.Fatalf("oversized initcode after the fork: got %v, want ErrMaxInitCodeSizeExceeded", err)
	}
	for _, tt := range []struct {
		name     string
		block    int64
		initcode []byte
	}{
		{"oversized before the fork", croissantForkBlock - 1, oversized},
		{"at the limit after the fork", croissantForkBlock, make([]byte, params.MaxInitCodeSize)},
	} {
		res, err := applyCreation(t, tt.block, tt.initcode)
		if err != nil {
			t.Fatalf("%s: %v", tt.name, err)
		}
		if res.Failed() {
			t.Fatalf("%s: execution failed: %v", tt.name, res.Err)
		}
	}
}
