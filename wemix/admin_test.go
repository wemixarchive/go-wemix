package wemix

import (
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"testing"
)

func TestHalveRewards(t *testing.T) {
	testcases := []struct {
		reward   *big.Int
		period   *big.Int
		past     *big.Int
		times    uint64
		expected *big.Int
	}{
		// sample test
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(0), 4, big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(99), 4, big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 4, big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(101), 4, big.NewInt(25e16)},

		// times test
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(0), 0, big.NewInt(1e18)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 0, big.NewInt(1e18)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 1, big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 2, big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 3, big.NewInt(25e16)},

		// brioche halving test
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(0), 16, big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 - 1), 16, big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200), 16, big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200*2 - 1), 16, big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 2), 16, big.NewInt(125e15)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 3), 16, big.NewInt(625e14)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 4), 16, big.NewInt(3125e13)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 5), 16, big.NewInt(15625e12)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 6), 16, big.NewInt(78125e11)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 7), 16, big.NewInt(390625e10)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 8), 16, big.NewInt(1953125e9)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 9), 16, big.NewInt(9765625e8)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 10), 16, big.NewInt(48828125e7)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 11), 16, big.NewInt(244140625e6)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 12), 16, big.NewInt(1220703125e5)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 13), 16, big.NewInt(6103515625e4)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 14), 16, big.NewInt(30517578125e3)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 15), 16, big.NewInt(152587890625e2)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 15), 17, big.NewInt(152587890625e2)},
	}

	for _, tc := range testcases {
		halved := halveRewards(tc.reward, tc.period, tc.past, tc.times)
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
				HalvingTimes:      10,
				NoRewardHereAfter: big.NewInt(101),
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(5e17),
		},

		// base block reward variations
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(7e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				HalvingTimes:      10,
				NoRewardHereAfter: big.NewInt(101),
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(35e17),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(3),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				HalvingTimes:      10,
				NoRewardHereAfter: big.NewInt(101),
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(1),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(1),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				HalvingTimes:      10,
				NoRewardHereAfter: big.NewInt(101),
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(0),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       nil, // it will use the default block reward
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				HalvingTimes:      10,
				NoRewardHereAfter: big.NewInt(101),
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
				HalvingTimes:      10,
				NoRewardHereAfter: big.NewInt(101),
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(1e18),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     nil,
				HalvingTimes:      10,
				NoRewardHereAfter: big.NewInt(101),
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(10),
		},
		{
			briocheConfig: &params.BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				HalvingTimes:      0,
				NoRewardHereAfter: big.NewInt(101),
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(10),
		},
	}

	for _, tc := range testcases {
		actual := getBriocheBlockReward(tc.briocheConfig, tc.blockNum)
		if tc.expected.Cmp(actual) != 0 {
			t.Errorf("getBriocheReward mismatched (expected=%v, actual=%v)", tc.expected, actual)
		}
	}
}
