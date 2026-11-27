package params

import (
	"math/big"
	"testing"
)

// TestIsCroissantGating checks that Croissant is gated by block number, unlike
// the timestamp-gated Shanghai fork on Ethereum mainnet.
func TestIsCroissantGating(t *testing.T) {
	const forkBlock = 100
	cfg := &ChainConfig{CroissantBlock: big.NewInt(forkBlock)}
	for _, tt := range []struct {
		num  int64
		want bool
	}{
		{0, false},
		{forkBlock - 1, false},
		{forkBlock, true},
		{forkBlock + 1, true},
	} {
		if got := cfg.IsCroissant(big.NewInt(tt.num)); got != tt.want {
			t.Errorf("IsCroissant(%d) = %v, want %v", tt.num, got, tt.want)
		}
	}

	unset := &ChainConfig{}
	if unset.IsCroissant(big.NewInt(1 << 40)) {
		t.Error("IsCroissant must be false when CroissantBlock is nil")
	}
}

// TestCroissantRulesPropagation checks that Rules carries the Croissant flag,
// which is what the EVM and the state transition actually read.
func TestCroissantRulesPropagation(t *testing.T) {
	const forkBlock = 10
	cfg := &ChainConfig{ChainID: big.NewInt(1), CroissantBlock: big.NewInt(forkBlock)}
	if cfg.Rules(big.NewInt(forkBlock-1), false).IsCroissant {
		t.Error("Rules.IsCroissant must be false before the fork block")
	}
	if !cfg.Rules(big.NewInt(forkBlock), false).IsCroissant {
		t.Error("Rules.IsCroissant must be true at the fork block")
	}
}

// TestCroissantCompatibility checks that rescheduling an already passed
// Croissant block is reported as an incompatible config change.
func TestCroissantCompatibility(t *testing.T) {
	stored := &ChainConfig{CroissantBlock: big.NewInt(10)}
	moved := &ChainConfig{CroissantBlock: big.NewInt(20)}

	if err := stored.CheckCompatible(moved, 5); err != nil {
		t.Errorf("moving a fork that has not been reached must be allowed: %v", err)
	}
	if err := stored.CheckCompatible(moved, 15); err == nil {
		t.Error("moving a fork the head has already passed must be rejected")
	}
}

// TestWemixNetworksScheduleCroissant pins the Croissant block of mainnet and
// testnet, keeps Dalgona unset, and checks the fork order. go-wbft must use the
// same Croissant blocks.
func TestWemixNetworksScheduleCroissant(t *testing.T) {
	for _, tt := range []struct {
		name      string
		cfg       *ChainConfig
		croissant int64
	}{
		{"mainnet", WemixMainnetChainConfig, 134_004_000},
		{"testnet", WemixTestnetChainConfig, 138_903_600},
	} {
		if tt.cfg.CroissantBlock == nil || tt.cfg.CroissantBlock.Int64() != tt.croissant {
			t.Errorf("%s: CroissantBlock = %v, want %d", tt.name, tt.cfg.CroissantBlock, tt.croissant)
		}
		if tt.cfg.DalgonaBlock != nil {
			t.Errorf("%s: DalgonaBlock = %v, want nil until scheduled", tt.name, tt.cfg.DalgonaBlock)
		}
		if err := tt.cfg.CheckConfigForkOrder(); err != nil {
			t.Errorf("%s: %v", tt.name, err)
		}
	}
}

// TestMaxInitCodeSizeDerivation pins EIP-3860's MAX_INITCODE_SIZE to twice
// go-wemix's MaxCodeSize, which is larger than Ethereum's.
func TestMaxInitCodeSizeDerivation(t *testing.T) {
	const wemixMaxInitCodeSize = 507904
	if MaxInitCodeSize != 2*MaxCodeSize {
		t.Fatalf("MaxInitCodeSize = %d, want 2*MaxCodeSize = %d", MaxInitCodeSize, 2*MaxCodeSize)
	}
	if MaxInitCodeSize != wemixMaxInitCodeSize {
		t.Fatalf("MaxInitCodeSize = %d, want %d", MaxInitCodeSize, wemixMaxInitCodeSize)
	}
}
