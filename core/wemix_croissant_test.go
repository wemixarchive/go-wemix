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

// TestCheckInitCodeSize checks the EIP-3860 initcode limit shared by the state
// transition and both transaction pools.
func TestCheckInitCodeSize(t *testing.T) {
	atLimit := make([]byte, params.MaxInitCodeSize)
	oversized := make([]byte, params.MaxInitCodeSize+1)

	for _, tt := range []struct {
		name     string
		data     []byte
		creation bool
		eip3860  bool
		wantErr  bool
	}{
		{"oversized creation after the fork", oversized, true, true, true},
		{"creation at the limit after the fork", atLimit, true, true, false},
		{"oversized creation before the fork", oversized, true, false, false},
		{"oversized call after the fork", oversized, false, true, false},
	} {
		err := CheckInitCodeSize(tt.data, tt.creation, tt.eip3860)
		if got := errors.Is(err, ErrMaxInitCodeSizeExceeded); got != tt.wantErr {
			t.Errorf("%s: got %v, want ErrMaxInitCodeSizeExceeded=%v", tt.name, err, tt.wantErr)
		}
	}
}

var (
	testSender   = common.HexToAddress("0x1000000000000000000000000000000000000001")
	testCoinbase = common.HexToAddress("0x2000000000000000000000000000000000000002")
)

// applyAt runs a message from testSender at block number on a london-era chain
// that activates Croissant at croissantForkBlock. A nil to means contract
// creation. code, when given, is installed at to before the message runs.
func applyAt(t *testing.T, block int64, to *common.Address, code, data []byte) (*ExecutionResult, error) {
	t.Helper()
	cfg := *params.AllEthashProtocolChanges
	cfg.CroissantBlock = big.NewInt(croissantForkBlock)

	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatal(err)
	}
	if to != nil && code != nil {
		statedb.SetCode(*to, code)
	}
	const gasLimit = 30_000_000
	blockCtx := vm.BlockContext{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		Coinbase:    testCoinbase,
		BlockNumber: big.NewInt(block),
		Time:        big.NewInt(0),
		Difficulty:  big.NewInt(1), // go-wemix PoA headers carry difficulty 1
		GasLimit:    gasLimit,
		BaseFee:     big.NewInt(0),
	}
	evm := vm.NewEVM(blockCtx, vm.TxContext{GasPrice: big.NewInt(0)}, statedb, &cfg, vm.Config{NoBaseFee: true})

	zero := big.NewInt(0)
	msg := types.NewMessage(testSender, to, 0, zero, gasLimit, zero, zero, zero, data, nil, true)
	return ApplyMessage(evm, msg, new(GasPool).AddGas(math.MaxUint64))
}

// applyCreation runs a contract creation message with the given initcode.
func applyCreation(t *testing.T, block int64, initcode []byte) (*ExecutionResult, error) {
	t.Helper()
	return applyAt(t, block, nil, nil, initcode)
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

// TestEIP3651WarmCoinbase checks the gas effect of EIP-3651 through the state
// transition: reading the coinbase balance costs a cold account access before
// Croissant and a warm one from the fork on.
func TestEIP3651WarmCoinbase(t *testing.T) {
	// COINBASE, BALANCE, POP, STOP
	balanceOfCoinbase := common.FromHex("41315000")
	contract := common.HexToAddress("0x3000000000000000000000000000000000000003")

	usedGas := func(block int64) uint64 {
		res, err := applyAt(t, block, &contract, balanceOfCoinbase, nil)
		if err != nil {
			t.Fatalf("block %d: %v", block, err)
		}
		if res.Failed() {
			t.Fatalf("block %d: execution failed: %v", block, res.Err)
		}
		return res.UsedGas
	}

	pre, post := usedGas(croissantForkBlock-1), usedGas(croissantForkBlock)
	want := params.ColdAccountAccessCostEIP2929 - params.WarmStorageReadCostEIP2929
	if pre-post != want {
		t.Fatalf("Croissant saves %d gas on a coinbase BALANCE, want %d (pre=%d post=%d)", pre-post, want, pre, post)
	}
}
