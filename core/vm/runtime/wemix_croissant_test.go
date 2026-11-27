package runtime

import (
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

// croissantForkBlock is where the test chain activates Croissant. Running at
// croissantForkBlock-1 exercises the pre-fork rules on the same config.
const croissantForkBlock = 10

var (
	// push0ReturnsOne runs PUSH1 1, PUSH0, MSTORE, PUSH1 32, PUSH0, RETURN and
	// returns a 32-byte word holding 1. Only PUSH0 is new.
	push0ReturnsOne = common.FromHex("60015f5260205ff3")

	// difficultyReturned runs DIFFICULTY, PUSH1 0, MSTORE, PUSH1 32, PUSH1 0,
	// RETURN and returns the block difficulty as a 32-byte word.
	difficultyReturned = common.FromHex("4460005260206000f3")
)

// wemixCroissantConfig is a london-era chain config, like go-wemix mainnet,
// that activates Croissant at croissantForkBlock.
func wemixCroissantConfig() *params.ChainConfig {
	zero := big.NewInt(0)
	return &params.ChainConfig{
		ChainID:             big.NewInt(1),
		HomesteadBlock:      zero,
		EIP150Block:         zero,
		EIP155Block:         zero,
		EIP158Block:         zero,
		ByzantiumBlock:      zero,
		ConstantinopleBlock: zero,
		PetersburgBlock:     zero,
		IstanbulBlock:       zero,
		MuirGlacierBlock:    zero,
		BerlinBlock:         zero,
		LondonBlock:         zero,
		CroissantBlock:      big.NewInt(croissantForkBlock),
	}
}

func configAt(block int64) *Config {
	return &Config{
		ChainConfig: wemixCroissantConfig(),
		BlockNumber: big.NewInt(block),
		Difficulty:  big.NewInt(1), // go-wemix PoA headers carry difficulty 1
	}
}

// TestEIP3855Push0 checks that PUSH0 is an invalid opcode before Croissant and
// pushes zero from the fork block on.
func TestEIP3855Push0(t *testing.T) {
	pre, _, err := Execute(push0ReturnsOne, nil, configAt(croissantForkBlock-1))
	var invalid *vm.ErrInvalidOpCode
	if !errors.As(err, &invalid) {
		t.Fatalf("pre-fork PUSH0: got (%x, %v), want ErrInvalidOpCode", pre, err)
	}

	post, _, err := Execute(push0ReturnsOne, nil, configAt(croissantForkBlock))
	if err != nil {
		t.Fatalf("post-fork PUSH0: %v", err)
	}
	if got := new(big.Int).SetBytes(post); got.Cmp(big.NewInt(1)) != 0 {
		t.Fatalf("post-fork PUSH0 returned %v, want 1", got)
	}
}

// TestCroissantDifficultyOpcode checks that DIFFICULTY still returns the header
// difficulty after Croissant. go-wemix has no PREVRANDAO value, so a jump table
// built on the merge set dereferences a nil Random here and panics.
func TestCroissantDifficultyOpcode(t *testing.T) {
	for _, block := range []int64{croissantForkBlock - 1, croissantForkBlock} {
		ret, _, err := Execute(difficultyReturned, nil, configAt(block))
		if err != nil {
			t.Fatalf("block %d: %v", block, err)
		}
		if got := new(big.Int).SetBytes(ret); got.Cmp(big.NewInt(1)) != 0 {
			t.Fatalf("block %d: DIFFICULTY returned %v, want 1", block, got)
		}
	}
}

// createOpcodeReturningAddress returns code that runs op (CREATE or CREATE2)
// on size zero bytes of memory and returns the created address as a word.
// Zero initcode is all STOPs, so the inner creation succeeds whenever the
// size is allowed.
func createOpcodeReturningAddress(op vm.OpCode, size uint32) []byte {
	code := []byte{}
	if op == vm.CREATE2 {
		code = append(code, byte(vm.PUSH1), 0) // salt
	}
	code = append(code, byte(vm.PUSH3), byte(size>>16), byte(size>>8), byte(size)) // size
	code = append(code,
		byte(vm.PUSH1), 0, // offset
		byte(vm.PUSH1), 0, // value
		byte(op),
		byte(vm.PUSH1), 0, byte(vm.MSTORE),
		byte(vm.PUSH1), 32, byte(vm.PUSH1), 0, byte(vm.RETURN),
	)
	return code
}

// TestEIP3860CreateOpcodeLimit checks that CREATE and CREATE2 refuse initcode
// over params.MaxInitCodeSize from the fork block on, accept it at the limit,
// and keep accepting oversized initcode before the fork. The interpreter
// reports a failing dynamic gas function as out of gas and keeps its cause.
func TestEIP3860CreateOpcodeLimit(t *testing.T) {
	const (
		atLimit = uint32(params.MaxInitCodeSize)
		over    = atLimit + 1
	)
	for _, op := range []vm.OpCode{vm.CREATE, vm.CREATE2} {
		for _, tt := range []struct {
			block   int64
			size    uint32
			allowed bool
		}{
			{croissantForkBlock - 1, over, true},
			{croissantForkBlock, atLimit, true},
			{croissantForkBlock, over, false},
		} {
			ret, _, err := Execute(createOpcodeReturningAddress(op, tt.size), nil, configAt(tt.block))
			if !tt.allowed {
				if !errors.Is(err, vm.ErrOutOfGas) || !strings.Contains(err.Error(), vm.ErrMaxInitCodeSizeExceeded.Error()) {
					t.Errorf("%v size %d at block %d: got (%x, %v), want ErrOutOfGas caused by %v", op, tt.size, tt.block, ret, err, vm.ErrMaxInitCodeSizeExceeded)
				}
				continue
			}
			if err != nil {
				t.Errorf("%v size %d at block %d: %v", op, tt.size, tt.block, err)
				continue
			}
			if common.BytesToAddress(ret) == (common.Address{}) {
				t.Errorf("%v size %d at block %d: creation failed", op, tt.size, tt.block)
			}
		}
	}
}
