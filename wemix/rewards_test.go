// test_rewards.go

package wemix

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

func hexToBigInt(hexNum string) *big.Int {
	if num, ok := new(big.Int).SetString(hexNum[2:], 16); ok {
		return num
	} else {
		return nil
	}
}

// TestDistributeRewards tests the DistributeRewards function
func TestDistributeRewards(t *testing.T) {

	hexToAddressPtr := func(addr string) *common.Address {
		address := common.HexToAddress(addr)
		return &address
	}

	// Test cases
	tests := []struct {
		name   string
		height *big.Int
		rp     *rewardParameters
		fees   *big.Int
		want   string
	}{
		{
			name:   "basic test",
			height: big.NewInt(100),
			rp: &rewardParameters{
				rewardAmount: big.NewInt(10000000000),
				staker:       hexToAddressPtr("0x1111111111111111111111111111111111111111"),
				ecoSystem:    hexToAddressPtr("0x2222222222222222222222222222222222222222"),
				maintenance:  hexToAddressPtr("0x3333333333333333333333333333333333333333"),
				feeCollector: nil,
				members: []*wemixMember{
					{
						Staker: common.HexToAddress("0x4444444444444444444444444444444444444444"),
						Reward: common.HexToAddress("0x4444444444444444444444444444444444444444"),
						Stake:  big.NewInt(500),
					},
					{
						Staker: common.HexToAddress("0x5555555555555555555555555555555555555555"),
						Reward: common.HexToAddress("0x5555555555555555555555555555555555555555"),
						Stake:  big.NewInt(500),
					},
				},
				distributionMethod: []*big.Int{big.NewInt(5000), big.NewInt(3000), big.NewInt(2000)},
				blocksPer:          10,
			},
			fees: big.NewInt(100),
			want: `[{"addr":"0x4444444444444444444444444444444444444444","reward":2500000000},{"addr":"0x5555555555555555555555555555555555555555","reward":2500000000},{"addr":"0x1111111111111111111111111111111111111111","reward":3000000000},{"addr":"0x2222222222222222222222222222222222222222","reward":2000000000},{"addr":"0x3333333333333333333333333333333333333333","reward":100}]`,
		},
		{
			name:   "case 1",
			height: hexToBigInt("0x10a187"),
			rp: &rewardParameters{
				rewardAmount: hexToBigInt("0xde0b6b3a7640000"),
				staker:       hexToAddressPtr("0x6f488615e6b462ce8909e9cd34c3f103994ab2fb"),
				ecoSystem:    hexToAddressPtr("0x6bd26c4a45e7d7cac2a389142f99f12e5713d719"),
				maintenance:  hexToAddressPtr("0x816e30b6c314ba5d1a67b1b54be944ce4554ed87"),
				feeCollector: nil,
				members: []*wemixMember{
					{
						Staker: common.HexToAddress("0x02b4b2d83786c8ee315db2ddac704794850d2149"),
						Reward: common.HexToAddress("0x02b4b2d83786c8ee315db2ddac704794850d2149"),
						Stake:  hexToBigInt("0x1a784379d99db42000000"),
					},
					{
						Staker: common.HexToAddress("0xb16d2494fddfa4c000deaf642d47673e5ca74e07"),
						Reward: common.HexToAddress("0xb16d2494fddfa4c000deaf642d47673e5ca74e07"),
						Stake:  hexToBigInt("0xe8ef1e96ae3897800000"),
					},
					{
						Staker: common.HexToAddress("0x452893ed818c0e3ea6f415aeab8ef08778087fc6"),
						Reward: common.HexToAddress("0x452893ed818c0e3ea6f415aeab8ef08778087fc6"),
						Stake:  hexToBigInt("0xc92b9a6adc4825c00000"),
					},
					{
						Staker: common.HexToAddress("0xf4404494ab647d29a62c57118c3a739c521fa004"),
						Reward: common.HexToAddress("0xf4404494ab647d29a62c57118c3a739c521fa004"),
						Stake:  hexToBigInt("0x1a0bd7e67e11f6e00000"),
					},
				},
				blocksPer:          1,
				distributionMethod: []*big.Int{big.NewInt(4000), big.NewInt(1000), big.NewInt(2500), big.NewInt(2500)},
			},
			fees: hexToBigInt("0x0"),
			want: `[{"addr":"0x02b4b2d83786c8ee315db2ddac704794850d2149","reward":191708602923556195},{"addr":"0xb16d2494fddfa4c000deaf642d47673e5ca74e07","reward":105439731607955907},{"addr":"0x452893ed818c0e3ea6f415aeab8ef08778087fc6","reward":91061586388689192},{"addr":"0xf4404494ab647d29a62c57118c3a739c521fa004","reward":11790079079798706},{"addr":"0x6f488615e6b462ce8909e9cd34c3f103994ab2fb","reward":100000000000000000},{"addr":"0x6bd26c4a45e7d7cac2a389142f99f12e5713d719","reward":250000000000000000},{"addr":"0x816e30b6c314ba5d1a67b1b54be944ce4554ed87","reward":250000000000000000}]`,
		},
		{
			name:   "with feeCollector",
			height: hexToBigInt("0x2a"),
			rp: &rewardParameters{
				rewardAmount: hexToBigInt("0xde0b6b3a7640000"),
				staker:       hexToAddressPtr("0xf5ed7476157980e831516cdd5493f5334a35b23e"),
				ecoSystem:    hexToAddressPtr("0x6bd26c4a45e7d7cac2a389142f99f12e5713d719"),
				maintenance:  hexToAddressPtr("0x816e30b6c314ba5d1a67b1b54be944ce4554ed87"),
				feeCollector: hexToAddressPtr("0x6f488615e6b462ce8909e9cd34c3f103994ab2fb"),
				members: []*wemixMember{
					{
						Staker: common.HexToAddress("0x02b4b2d83786c8ee315db2ddac704794850d2149"),
						Reward: common.HexToAddress("0x02b4b2d83786c8ee315db2ddac704794850d2149"),
						Stake:  hexToBigInt("0x1a784379d99db42000000"),
					},
					{
						Staker: common.HexToAddress("0xb16d2494fddfa4c000deaf642d47673e5ca74e07"),
						Reward: common.HexToAddress("0xb16d2494fddfa4c000deaf642d47673e5ca74e07"),
						Stake:  hexToBigInt("0xe8ef1e96ae3897800000"),
					},
					{
						Staker: common.HexToAddress("0x452893ed818c0e3ea6f415aeab8ef08778087fc6"),
						Reward: common.HexToAddress("0x452893ed818c0e3ea6f415aeab8ef08778087fc6"),
						Stake:  hexToBigInt("0xc92b9a6adc4825c00000"),
					},
				},
				blocksPer:          1,
				distributionMethod: []*big.Int{big.NewInt(4000), big.NewInt(1000), big.NewInt(2500), big.NewInt(2500)},
			},
			fees: hexToBigInt("0xadc5e885b956557f"),
			want: `[{"addr":"0x02b4b2d83786c8ee315db2ddac704794850d2149","reward":197530864197530865},{"addr":"0xb16d2494fddfa4c000deaf642d47673e5ca74e07","reward":108641975308641975},{"addr":"0x452893ed818c0e3ea6f415aeab8ef08778087fc6","reward":93827160493827160},{"addr":"0xf5ed7476157980e831516cdd5493f5334a35b23e","reward":100000000000000000},{"addr":"0x6bd26c4a45e7d7cac2a389142f99f12e5713d719","reward":250000000000000000},{"addr":"0x816e30b6c314ba5d1a67b1b54be944ce4554ed87","reward":250000000000000000},{"addr":"0x6f488615e6b462ce8909e9cd34c3f103994ab2fb","reward":12521670000011269503}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the distributeRewards function
			rewards, err := distributeRewards(tt.height, tt.rp, tt.rp.rewardAmount, tt.fees)
			rewardsString, _ := json.Marshal(rewards)
			if string(rewardsString) != tt.want {
				t.Errorf("distributeRewards() failed: %v, %v <-> %v", err, tt.want, string(rewardsString))
			}

			distTotal := new(big.Int)
			for _, dist := range tt.rp.distributionMethod {
				distTotal.Add(distTotal, dist)
			}
			totalRewards := new(big.Int)
			memberRewards := new(big.Int)
			for i, reward := range rewards {
				totalRewards.Add(totalRewards, reward.Reward)
				if i < len(tt.rp.members) {
					memberRewards.Add(memberRewards, reward.Reward)
				}
			}
			totalAmount := new(big.Int).Set(tt.rp.rewardAmount)
			totalAmount.Add(totalAmount, tt.fees)
			memberAmount := new(big.Int).Set(tt.rp.rewardAmount)
			memberAmount = memberAmount.Mul(memberAmount, tt.rp.distributionMethod[0])
			memberAmount = memberAmount.Div(memberAmount, distTotal)
			if memberRewards.Cmp(memberAmount) != 0 {
				t.Errorf("members reward amount mismatched! sum=%d, rewards=%d", memberRewards, memberAmount)
			}
			if totalRewards.Cmp(totalAmount) != 0 {
				t.Errorf("total reward amount mismatched! sum=%d, rewards=%d", totalRewards, totalAmount)
			}
		})
	}
}

func makeCalculateRewardFunc(rp *rewardParameters) func(config *params.ChainConfig, num, fees *big.Int, addBalance func(common.Address, *big.Int)) ([]byte, error) {
	return func(config *params.ChainConfig, num, fees *big.Int, addBalance func(common.Address, *big.Int)) ([]byte, error) {
		return calculateRewardsWithParams(config, rp, num, fees, addBalance)
	}
}

func makeSignBlockFunc(privateKey *ecdsa.PrivateKey) func(height *big.Int, hash common.Hash) (common.Address, []byte, error) {
	return func(height *big.Int, hash common.Hash) (coinbase common.Address, sig []byte, err error) {
		data := append(height.Bytes(), hash.Bytes()...)
		data = crypto.Keccak256(data)
		sig, _ = crypto.Sign(data, privateKey)
		return crypto.PubkeyToAddress(privateKey.PublicKey), sig, nil
	}
}

func verifyBlockSigForTest(height *big.Int, coinbase common.Address, nodeId []byte, hash common.Hash, sig []byte, checkMinerLimit bool) bool {
	var data []byte
	data = append(height.Bytes(), hash.Bytes()...)
	data = crypto.Keccak256(data)
	pubKey, err := crypto.SigToPub(data, sig)
	if err != nil {
		return false
	}
	signer := crypto.PubkeyToAddress(*pubKey)
	return coinbase == signer
}

func TestRewardValidation(t *testing.T) {
	// use wemix consensus
	params.ConsensusMethod = params.ConsensusPoA

	var (
		db         = rawdb.NewMemoryDatabase()
		key, _     = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
		address    = crypto.PubkeyToAddress(key.PublicKey)
		funds      = big.NewInt(100000000000000)
		deleteAddr = common.Address{1}
		gspec      = &core.Genesis{
			Config: &params.ChainConfig{
				ChainID:      big.NewInt(1),
				LondonBlock:  common.Big0,
				BriocheBlock: common.Big0,
				Brioche: &params.BriocheConfig{
					BlockReward:       big.NewInt(100),
					FirstHalvingBlock: big.NewInt(0),
					HalvingPeriod:     big.NewInt(10),
					FinishRewardBlock: big.NewInt(30),
					HalvingTimes:      3,
					HalvingRate:       50,
				}},
			Alloc: core.GenesisAlloc{address: {Balance: funds}, deleteAddr: {Balance: big.NewInt(0)}},
		}
		genesis = gspec.MustCommit(db)
		signer  = types.LatestSigner(gspec.Config)
	)

	rp := &rewardParameters{
		rewardAmount: big.NewInt(1e18),
		staker:       &common.Address{0x11},
		ecoSystem:    &common.Address{0x22},
		maintenance:  &common.Address{0x33},
		feeCollector: &common.Address{0x44},
		members: []*wemixMember{
			{
				Staker: common.HexToAddress("0x02b4b2d83786c8ee315db2ddac704794850d2149"),
				Reward: common.HexToAddress("0x02b4b2d83786c8ee315db2ddac704794850d2149"),
				Stake:  hexToBigInt("0x1a784379d99db42000000"),
			},
			{
				Staker: common.HexToAddress("0xb16d2494fddfa4c000deaf642d47673e5ca74e07"),
				Reward: common.HexToAddress("0xb16d2494fddfa4c000deaf642d47673e5ca74e07"),
				Stake:  hexToBigInt("0xe8ef1e96ae3897800000"),
			},
			{
				Staker: common.HexToAddress("0x452893ed818c0e3ea6f415aeab8ef08778087fc6"),
				Reward: common.HexToAddress("0x452893ed818c0e3ea6f415aeab8ef08778087fc6"),
				Stake:  hexToBigInt("0xc92b9a6adc4825c00000"),
			},
		},
		blocksPer:          1,
		distributionMethod: []*big.Int{big.NewInt(4000), big.NewInt(1000), big.NewInt(2500), big.NewInt(2500)},
	}

	wemixminer.CalculateRewardsFunc = makeCalculateRewardFunc(rp)
	wemixminer.SignBlockFunc = makeSignBlockFunc(key)
	wemixminer.VerifyBlockSigFunc = verifyBlockSigForTest

	blockchain, _ := core.NewBlockChain(db, nil, gspec.Config, ethash.NewFaker(), vm.Config{}, nil, nil)
	defer blockchain.Stop()

	byzantineConfig := &params.ChainConfig{
		ChainID:      gspec.Config.ChainID,
		LondonBlock:  gspec.Config.LondonBlock,
		BriocheBlock: gspec.Config.BriocheBlock,
		Brioche: &params.BriocheConfig{
			BlockReward:       big.NewInt(200), // different reward!!
			FirstHalvingBlock: gspec.Config.Brioche.FirstHalvingBlock,
			HalvingPeriod:     gspec.Config.Brioche.HalvingPeriod,
			FinishRewardBlock: gspec.Config.Brioche.FinishRewardBlock,
			HalvingTimes:      gspec.Config.Brioche.HalvingTimes,
			HalvingRate:       gspec.Config.Brioche.HalvingRate,
		}}
	blocks, _ := core.GenerateChain(byzantineConfig, genesis, ethash.NewFaker(), db, 1, func(i int, gen *core.BlockGen) {
		tx, err := types.SignTx(types.NewTransaction(gen.TxNonce(address), common.Address{0x00}, big.NewInt(1), params.TxGas, gen.BaseFee(), nil), signer, key)
		if err != nil {
			panic(err)
		}
		gen.AddTx(tx)
	})

	if _, err := blockchain.InsertChain(blocks); err != nil {
		if !strings.HasPrefix(err.Error(), "remote block hash is different") {
			t.Fatal(err)
		}
	} else {
		t.Fatal("reward validation failed")
	}
}

func TestBriocheHardFork(t *testing.T) {
	// use wemix consensus
	params.ConsensusMethod = params.ConsensusPoA

	var (
		db         = rawdb.NewMemoryDatabase()
		key, _     = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
		address    = crypto.PubkeyToAddress(key.PublicKey)
		funds      = big.NewInt(100000000000000000)
		deleteAddr = common.Address{1}
		gspec      = &core.Genesis{
			Config: &params.ChainConfig{
				ChainID:      big.NewInt(1),
				LondonBlock:  common.Big0,
				BriocheBlock: big.NewInt(2),
				Brioche: &params.BriocheConfig{
					// 1 block reward: 1e18
					// 2 block reward: 4e17 (brioche start)
					// 3 block reward: 4e17
					// 4 block reward: 2e17 (first halving)
					// 5 block reward: 2e17
					// 6 block reward: 1e17 (second halving)
					// 7~ block reward: 0
					BlockReward:       big.NewInt(4e17),
					FirstHalvingBlock: big.NewInt(4),
					HalvingPeriod:     big.NewInt(2),
					FinishRewardBlock: big.NewInt(7),
					HalvingTimes:      2,
					HalvingRate:       50,
				}},
			Alloc: core.GenesisAlloc{address: {Balance: funds}, deleteAddr: {Balance: big.NewInt(0)}},
		}
		genesis = gspec.MustCommit(db)
		signer  = types.LatestSigner(gspec.Config)
	)

	expectedBlockReward := []*big.Int{
		big.NewInt(0), // zero block reward; not used
		big.NewInt(1e18),
		big.NewInt(4e17),
		big.NewInt(4e17),
		big.NewInt(2e17),
		big.NewInt(2e17),
		big.NewInt(1e17),
		big.NewInt(0),
	}

	miners := []common.Address{
		common.HexToAddress("0x02b4b2d83786c8ee315db2ddac704794850d2149"),
		common.HexToAddress("0xb16d2494fddfa4c000deaf642d47673e5ca74e07"),
	}
	rp := &rewardParameters{
		rewardAmount: big.NewInt(1e18),
		staker:       &common.Address{0x11},
		ecoSystem:    &common.Address{0x22},
		maintenance:  &common.Address{0x33},
		feeCollector: &common.Address{0x44},
		members: []*wemixMember{
			{
				Staker: miners[0],
				Reward: miners[0],
				Stake:  hexToBigInt("0xFE1C215E8F838E00000"), // 75e21 (75%)
			},
			{
				Staker: miners[1],
				Reward: miners[1],
				Stake:  hexToBigInt("0x54B40B1F852BDA00000"), // 25e21 (25%)
			},
		},
		blocksPer:          1,
		distributionMethod: []*big.Int{big.NewInt(5000), big.NewInt(0), big.NewInt(2500), big.NewInt(2500)}, // miner, staker, eco, maintenance
	}

	wemixminer.CalculateRewardsFunc = makeCalculateRewardFunc(rp)
	wemixminer.SignBlockFunc = makeSignBlockFunc(key)
	wemixminer.VerifyBlockSigFunc = verifyBlockSigForTest

	blockchain, _ := core.NewBlockChain(db, nil, gspec.Config, ethash.NewFaker(), vm.Config{}, nil, nil)
	defer blockchain.Stop()

	parent := genesis
	for i := 1; i <= 7; i++ {
		statedb, _ := blockchain.State()
		miner0Bal := statedb.GetBalance(miners[0])
		miner1Bal := statedb.GetBalance(miners[1])

		blocks, _ := core.GenerateChain(gspec.Config, parent, ethash.NewFaker(), db, 1, func(i int, gen *core.BlockGen) {
			tx, err := types.SignTx(types.NewTransaction(gen.TxNonce(address), common.Address{0x00}, big.NewInt(1), params.TxGas, gen.BaseFee(), nil), signer, key)
			if err != nil {
				panic(err)
			}
			gen.AddTx(tx)
		})

		if _, err := blockchain.InsertChain(blocks); err != nil {
			t.Fatal(err)
		}
		statedb, _ = blockchain.State()

		miner0Reward := new(big.Int).Div(expectedBlockReward[i], big.NewInt(2))
		miner0Reward = miner0Reward.Mul(miner0Reward, big.NewInt(3))
		miner0Reward = miner0Reward.Div(miner0Reward, big.NewInt(4))

		miner1Reward := new(big.Int).Div(expectedBlockReward[i], big.NewInt(2))
		miner1Reward = miner1Reward.Div(miner1Reward, big.NewInt(4))

		miner0Bal = new(big.Int).Add(miner0Bal, miner0Reward)
		miner1Bal = new(big.Int).Add(miner1Bal, miner1Reward)
		if statedb.GetBalance(miners[0]).Cmp(miner0Bal) != 0 {
			t.Logf("miner bal = %v, expected = %v", statedb.GetBalance(miners[0]), miner0Bal)
			t.Fatal("block reward mismatched for miner0")
		}
		if statedb.GetBalance(miners[1]).Cmp(miner1Bal) != 0 {
			t.Fatal("block reward mismatched for miner1")
		}

		parent = blocks[0]
	}
}

// EOF
