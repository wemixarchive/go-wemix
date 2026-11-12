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

// TestWemixNetworksScheduleCroissant checks that mainnet and testnet both
// declare a Croissant block, so neither network silently stays without one.
func TestWemixNetworksScheduleCroissant(t *testing.T) {
	for name, cfg := range map[string]*ChainConfig{
		"mainnet": WemixMainnetChainConfig,
		"testnet": WemixTestnetChainConfig,
	} {
		if cfg.CroissantBlock == nil {
			t.Errorf("%s: CroissantBlock is nil", name)
		}
		if err := cfg.CheckConfigForkOrder(); err != nil {
			t.Errorf("%s: %v", name, err)
		}
	}
}

// TestCroissantBeforeDalgona checks that a config scheduling Dalgona must also
// schedule Croissant at or before it. The rest of the go-wemix fork order is
// not enforced, so only this pair is checked.
func TestCroissantBeforeDalgona(t *testing.T) {
	for _, tt := range []struct {
		name      string
		croissant *big.Int
		dalgona   *big.Int
		wantErr   bool
	}{
		{"neither", nil, nil, false},
		{"croissant only", big.NewInt(10), nil, false},
		{"croissant before dalgona", big.NewInt(10), big.NewInt(20), false},
		{"same block", big.NewInt(10), big.NewInt(10), false},
		{"dalgona without croissant", nil, big.NewInt(20), true},
		{"croissant after dalgona", big.NewInt(30), big.NewInt(20), true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &ChainConfig{CroissantBlock: tt.croissant, DalgonaBlock: tt.dalgona}
			err := cfg.CheckConfigForkOrder()
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckConfigForkOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
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
