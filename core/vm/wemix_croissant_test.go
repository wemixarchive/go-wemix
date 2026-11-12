package vm

import (
	"reflect"
	"testing"
)

// executorOf returns an identity for the execute function of op, so two jump
// tables can be compared. Every slot is non-nil after validate(), because
// undefined opcodes are filled with opUndefined.
func executorOf(jt JumpTable, op OpCode) uintptr {
	return reflect.ValueOf(jt[op].execute).Pointer()
}

// TestCroissantPush0 checks that PUSH0 (EIP-3855) is wired only in the
// Croissant instruction set.
func TestCroissantPush0(t *testing.T) {
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
