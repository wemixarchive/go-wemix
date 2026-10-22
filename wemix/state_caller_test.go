package wemix

import (
	"bytes"
	"context"
	"encoding/hex"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	gov "github.com/ethereum/go-ethereum/wemix/bind"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

// autoCommitBackend seals a block per transaction so that gas is always
// estimated against the block the transaction lands in.
type autoCommitBackend struct {
	*backends.SimulatedBackend
}

func (b *autoCommitBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if err := b.SimulatedBackend.SendTransaction(ctx, tx); err != nil {
		return err
	}
	b.Commit()
	return nil
}

// deployGovForTest deploys and initializes the governance contracts on an
// in-memory chain, returning the backend and the registry owner.
func deployGovForTest(t *testing.T) (*backends.SimulatedBackend, common.Address) {
	t.Helper()

	// deploy on a plain chain; PoA would require the wemix admin to seal blocks
	orig := params.ConsensusMethod
	params.ConsensusMethod = params.ConsensusPoW
	t.Cleanup(func() { params.ConsensusMethod = orig })

	key, _ := crypto.GenerateKey()
	owner := crypto.PubkeyToAddress(key.PublicKey)
	balance := new(big.Int).Mul(big.NewInt(1e9), big.NewInt(params.Ether))
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{owner: {Balance: balance}}, 1_000_000_000)
	t.Cleanup(func() { sim.Close() })

	opts, err := bind.NewKeyedTransactorWithChainID(key, params.AllEthashProtocolChanges.ChainID)
	if err != nil {
		t.Fatal(err)
	}

	backend := &autoCommitBackend{sim}
	contracts, err := gov.DeployGovContracts(opts, backend, map[string]common.Address{
		gov.DOMAIN_StakingReward: common.HexToAddress("0x1001"),
		gov.DOMAIN_Ecosystem:     common.HexToAddress("0x1002"),
		gov.DOMAIN_Maintenance:   common.HexToAddress("0x1003"),
		gov.DOMAIN_FeeCollector:  common.HexToAddress("0x1004"),
	})
	if err != nil {
		t.Fatal(err)
	}
	env := gov.DefaultInitEnvStorage
	members := gov.InitMembers{{
		Staker:  owner,
		Voter:   owner,
		Reward:  owner,
		Name:    "node1",
		Enode:   hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)[1:]),
		Ip:      "127.0.0.1",
		Port:    8589,
		Deposit: env.STAKING_MIN,
	}}
	if err := gov.ExecuteInitialize(contracts, opts, backend, env.STAKING_MIN, env, members); err != nil {
		t.Fatal(err)
	}
	return sim, owner
}

func TestBlockStateCallerMatchesEthCall(t *testing.T) {
	sim, owner := deployGovForTest(t)
	head := sim.Blockchain().CurrentBlock().Header()
	headState, err := sim.Blockchain().StateAt(head.Root)
	if err != nil {
		t.Fatal(err)
	}

	// the simulated backend serves eth_call on the head block only
	want, err := (&wemixAdmin{bootAccount: owner}).getRewardParams(context.Background(), sim, head.Number)
	if err != nil {
		t.Fatal(err)
	}
	got, err := (&wemixAdmin{bootAccount: owner}).getRewardParams(context.Background(), newBlockStateCaller(sim.Blockchain(), head, headState), head.Number)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("reward params mismatch\n got: %+v\nwant: %+v", got, want)
	}
	if len(got.members) != 1 || got.members[0].Stake.Cmp(gov.DefaultInitEnvStorage.STAKING_MIN) != 0 {
		t.Fatalf("unexpected members: %+v", got.members)
	}
	if got.feeCollector == nil || *got.feeCollector != common.HexToAddress("0x1004") {
		t.Fatalf("unexpected fee collector: %v", got.feeCollector)
	}
}

func TestBlockStateCallerRejectsOtherBlock(t *testing.T) {
	sim, owner := deployGovForTest(t)
	head := sim.Blockchain().CurrentBlock().Header()
	headState, err := sim.Blockchain().StateAt(head.Root)
	if err != nil {
		t.Fatal(err)
	}
	caller := newBlockStateCaller(sim.Blockchain(), head, headState)

	for _, height := range []*big.Int{nil, new(big.Int).Add(head.Number, common.Big1)} {
		if _, err := (&wemixAdmin{bootAccount: owner}).getRewardParams(context.Background(), caller, height); err == nil {
			t.Fatalf("height %v: expected an error", height)
		}
	}
}

type credit struct {
	addr common.Address
	amt  *big.Int
}

func TestCalculateRewardsMatchesEthCall(t *testing.T) {
	sim, owner := deployGovForTest(t)
	bc := sim.Blockchain()
	parent := bc.CurrentBlock().Header()

	prevAdmin := admin
	admin = &wemixAdmin{bootAccount: owner}
	t.Cleanup(func() { admin = prevAdmin })
	header := &types.Header{ParentHash: parent.Hash(), Number: new(big.Int).Add(parent.Number, common.Big1), Fees: big.NewInt(12345)}

	// expected: reward parameters through eth_call at the parent (head) block
	rp, err := (&wemixAdmin{bootAccount: owner}).getRewardParams(context.Background(), sim, parent.Number)
	if err != nil {
		t.Fatal(err)
	}
	var wantCredits []credit
	want, err := calculateRewardsWithParams(bc.Config(), rp, header.Number, header.Fees, func(a common.Address, v *big.Int) {
		wantCredits = append(wantCredits, credit{a, v})
	})
	if err != nil {
		t.Fatal(err)
	}

	var gotCredits []credit
	got, err := calculateRewards(bc, header, bc.StateCache(), func(a common.Address, v *big.Int) {
		gotCredits = append(gotCredits, credit{a, v})
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("rewards mismatch\n got: %s\nwant: %s", got, want)
	}
	if !reflect.DeepEqual(gotCredits, wantCredits) {
		t.Fatalf("credits mismatch\n got: %v\nwant: %v", gotCredits, wantCredits)
	}
}

func TestCalculateRewardsWithoutParent(t *testing.T) {
	sim, _ := deployGovForTest(t)
	bc := sim.Blockchain()
	header := &types.Header{ParentHash: common.HexToHash("0xdead"), Number: big.NewInt(2), Fees: common.Big0}
	if _, err := calculateRewards(bc, header, bc.StateCache(), nil); err != wemixminer.ErrNotInitialized {
		t.Fatalf("err = %v, want ErrNotInitialized", err)
	}
}
