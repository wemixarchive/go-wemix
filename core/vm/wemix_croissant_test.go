package vm

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

// executorOf returns an identity for the execute function of op, so two jump
// tables can be compared. Every slot is non-nil after validate(), because
// undefined opcodes are filled with opUndefined.
func executorOf(jt JumpTable, op OpCode) uintptr {
	return reflect.ValueOf(jt[op].execute).Pointer()
}

// TestEIP3855InCroissantSet checks that PUSH0 (EIP-3855) is wired only in the
// Croissant instruction set.
func TestEIP3855InCroissantSet(t *testing.T) {
	push0 := reflect.ValueOf(opPush0).Pointer()

	if executorOf(newLondonInstructionSet(), PUSH0) == push0 {
		t.Error("PUSH0 must not be enabled before Croissant")
	}
	if executorOf(newCroissantInstructionSet(), PUSH0) != push0 {
		t.Error("PUSH0 must execute opPush0 in the Croissant set")
	}
}

// TestCroissantKeepsLondonOpcodes checks that Croissant only adds to the london
// set that go-wemix runs today. In particular 0x44 must stay DIFFICULTY:
// go-wemix never went through the merge, so PREVRANDAO has no value to read.
func TestCroissantKeepsLondonOpcodes(t *testing.T) {
	london := newLondonInstructionSet()
	croissant := newCroissantInstructionSet()

	if executorOf(croissant, DIFFICULTY) != reflect.ValueOf(opDifficulty).Pointer() {
		t.Fatal("0x44 must execute opDifficulty in the Croissant set")
	}
	for i := 0; i < len(london); i++ {
		op := OpCode(i)
		if op == PUSH0 {
			continue
		}
		if executorOf(london, op) != executorOf(croissant, op) {
			t.Errorf("%v changed its executor in the Croissant set", op)
		}
	}
}

// TestEIP3860CreateGasInCroissantSet checks that CREATE and CREATE2 switch to
// their EIP-3860 dynamic gas functions in the Croissant set only.
func TestEIP3860CreateGasInCroissantSet(t *testing.T) {
	london := newLondonInstructionSet()
	croissant := newCroissantInstructionSet()
	for op, want := range map[OpCode]gasFunc{CREATE: gasCreateEip3860, CREATE2: gasCreate2Eip3860} {
		wantPtr := reflect.ValueOf(want).Pointer()
		if reflect.ValueOf(london[op].dynamicGas).Pointer() == wantPtr {
			t.Errorf("%v must not meter initcode before Croissant", op)
		}
		if reflect.ValueOf(croissant[op].dynamicGas).Pointer() != wantPtr {
			t.Errorf("%v must meter initcode in the Croissant set", op)
		}
	}
}

// createStack builds the stack CREATE and CREATE2 read, where the initcode
// size is always Back(2). CREATE2 additionally reads a salt at Back(3).
func createStack(size uint64) *Stack {
	st := newstack()
	st.push(new(uint256.Int))                 // salt, only read by CREATE2
	st.push(new(uint256.Int).SetUint64(size)) // size
	st.push(new(uint256.Int))                 // offset
	st.push(new(uint256.Int))                 // value
	return st
}

// TestEIP3860CreateGas checks the per-word initcode charge and the size cap of
// the EIP-3860 gas functions against the pre-Croissant ones.
func TestEIP3860CreateGas(t *testing.T) {
	for _, tt := range []struct {
		name     string
		pre, eip gasFunc
	}{
		{"CREATE", gasCreate, gasCreateEip3860},
		{"CREATE2", gasCreate2, gasCreate2Eip3860},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for _, size := range []uint64{0, 1, 32, 33, params.MaxInitCodeSize} {
				pre, err := tt.pre(nil, nil, createStack(size), NewMemory(), size)
				if err != nil {
					t.Fatalf("size %d: pre-Croissant gas: %v", size, err)
				}
				got, err := tt.eip(nil, nil, createStack(size), NewMemory(), size)
				if err != nil {
					t.Fatalf("size %d: EIP-3860 gas: %v", size, err)
				}
				if want := pre + params.InitCodeWordGas*toWordSize(size); got != want {
					t.Errorf("size %d: gas = %d, want %d", size, got, want)
				}
			}

			over := uint64(params.MaxInitCodeSize + 1)
			if _, err := tt.eip(nil, nil, createStack(over), NewMemory(), over); !errors.Is(err, ErrMaxInitCodeSizeExceeded) {
				t.Errorf("initcode over the limit: got %v, want ErrMaxInitCodeSizeExceeded", err)
			}
		})
	}
}
