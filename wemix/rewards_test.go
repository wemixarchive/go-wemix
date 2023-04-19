// test_rewards.go

package wemix

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// TestDistributeRewards tests the DistributeRewards function
func TestDistributeRewards(t *testing.T) {

	hexToAddressPtr := func(addr string) *common.Address {
		address := common.HexToAddress(addr)
		return &address
	}

	hexToBigInt := func(hexNum string) *big.Int {
		if num, ok := new(big.Int).SetString(hexNum[2:], 16); ok {
			return num
		} else {
			return nil
		}
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
			rewards, err := distributeRewards(tt.height, tt.rp, tt.fees)
			rewardsString, _ := json.Marshal(rewards)
			if string(rewardsString) != tt.want {
				t.Errorf("distributeRewards() failed: %v, %v <-> %v", err, tt.want, string(rewardsString))
			}
		})
	}
}

// EOF
