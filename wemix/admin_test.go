package wemix

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/params"
)

func TestHalveRewards(t *testing.T) {
	testcases := []struct {
		reward   *big.Int
		period   *big.Int
		past     *big.Int
		times    uint64
		rate     uint32
		expected *big.Int
	}{
		// sample test
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(0), 4, 50, big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(99), 4, 50, big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 4, 50, big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(101), 4, 50, big.NewInt(25e16)},

		// times test
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(0), 0, 50, big.NewInt(1e18)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 0, 50, big.NewInt(1e18)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 1, 50, big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 2, 50, big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 3, 50, big.NewInt(25e16)},

		// rate test
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 2, 10, big.NewInt(1e16)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(200), 3, 10, big.NewInt(1e15)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 2, 20, big.NewInt(4e16)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 2, 30, big.NewInt(9e16)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 2, 90, big.NewInt(81e16)},

		// brioche halving test
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(0), 16, 50, big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 - 1), 16, 50, big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200), 16, 50, big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200*2 - 1), 16, 50, big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 2), 16, 50, big.NewInt(125e15)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 3), 16, 50, big.NewInt(625e14)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 4), 16, 50, big.NewInt(3125e13)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 5), 16, 50, big.NewInt(15625e12)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 6), 16, 50, big.NewInt(78125e11)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 7), 16, 50, big.NewInt(390625e10)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 8), 16, 50, big.NewInt(1953125e9)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 9), 16, 50, big.NewInt(9765625e8)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 10), 16, 50, big.NewInt(48828125e7)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 11), 16, 50, big.NewInt(244140625e6)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 12), 16, 50, big.NewInt(1220703125e5)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 13), 16, 50, big.NewInt(6103515625e4)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 14), 16, 50, big.NewInt(30517578125e3)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 15), 16, 50, big.NewInt(152587890625e2)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 15), 17, 50, big.NewInt(152587890625e2)},
	}

	for _, tc := range testcases {
		halved := halveRewards(tc.reward, tc.period, tc.past, tc.times, tc.rate)
		if tc.expected.Cmp(halved) != 0 {
			t.Errorf("halveRewards mismatched (expected=%v, actual=%v)", tc.expected, halved)
		}
	}
}

func TestGetBriocheBlockReward(t *testing.T) {
	testcases := []struct {
		briocheConfig *params.BriocheConfig
		blockNum      *big.Int
		expected      *big.Int
	}{
		// normal case
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(5e17),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(201),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(200),
			expected: big.NewInt(25e16),
		},

		// base block reward variations
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(7e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(35e17),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(3),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(1),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(1),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(0),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       nil, // it will use the default block reward
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(5e17),
		},

		// no halving
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       nil, // it will use the default block reward
				FirstHalvingBlock: nil,
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(1e18),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     nil,
				NoRewardHereafter: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(10),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(101),
				HalvingTimes:      0,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(10),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(200),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(99), // not yet halving time
			expected: big.NewInt(10),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(200),
				HalvingTimes:      10,
				HalvingRate:       100, // no halving rate
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(10),
		},

		// no reward case
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(200),
				HalvingTimes:      10,
				HalvingRate:       0, // no reward
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(0),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(200),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(199),
			expected: big.NewInt(5),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(200),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(200),
			expected: big.NewInt(0),
		},

		// halving rate variations
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(1000),
				HalvingTimes:      10,
				HalvingRate:       10,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(1e17),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(1000),
				HalvingTimes:      10,
				HalvingRate:       10,
			},
			blockNum: big.NewInt(200),
			expected: big.NewInt(1e16),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(1000),
				HalvingTimes:      10,
				HalvingRate:       1,
			},
			blockNum: big.NewInt(300),
			expected: big.NewInt(1e12),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				NoRewardHereafter: big.NewInt(1000),
				HalvingTimes:      10,
				HalvingRate:       99,
			},
			blockNum: big.NewInt(300),
			expected: big.NewInt(970299e12),
		},
	}

	for _, tc := range testcases {
		actual := getBriocheBlockReward(tc.briocheConfig, tc.blockNum)
		if tc.expected.Cmp(actual) != 0 {
			t.Errorf("getBriocheReward mismatched (expected=%v, actual=%v)", tc.expected, actual)
		}
	}
}
