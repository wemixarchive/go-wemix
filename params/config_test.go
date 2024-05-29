// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCheckCompatible(t *testing.T) {
	type test struct {
		stored, new *ChainConfig
		head        uint64
		wantErr     *ConfigCompatError
	}
	tests := []test{
		{stored: AllEthashProtocolChanges, new: AllEthashProtocolChanges, head: 0, wantErr: nil},
		{stored: AllEthashProtocolChanges, new: AllEthashProtocolChanges, head: 100, wantErr: nil},
		{
			stored:  &ChainConfig{EIP150Block: big.NewInt(10)},
			new:     &ChainConfig{EIP150Block: big.NewInt(20)},
			head:    9,
			wantErr: nil,
		},
		{
			stored: AllEthashProtocolChanges,
			new:    &ChainConfig{HomesteadBlock: nil},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "Homestead fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    nil,
				RewindTo:     0,
			},
		},
		{
			stored: AllEthashProtocolChanges,
			new:    &ChainConfig{HomesteadBlock: big.NewInt(1)},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "Homestead fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    big.NewInt(1),
				RewindTo:     0,
			},
		},
		{
			stored: &ChainConfig{HomesteadBlock: big.NewInt(30), EIP150Block: big.NewInt(10)},
			new:    &ChainConfig{HomesteadBlock: big.NewInt(25), EIP150Block: big.NewInt(20)},
			head:   25,
			wantErr: &ConfigCompatError{
				What:         "EIP150 fork block",
				StoredConfig: big.NewInt(10),
				NewConfig:    big.NewInt(20),
				RewindTo:     9,
			},
		},
		{
			stored:  &ChainConfig{ConstantinopleBlock: big.NewInt(30)},
			new:     &ChainConfig{ConstantinopleBlock: big.NewInt(30), PetersburgBlock: big.NewInt(30)},
			head:    40,
			wantErr: nil,
		},
		{
			stored: &ChainConfig{ConstantinopleBlock: big.NewInt(30)},
			new:    &ChainConfig{ConstantinopleBlock: big.NewInt(30), PetersburgBlock: big.NewInt(31)},
			head:   40,
			wantErr: &ConfigCompatError{
				What:         "Petersburg fork block",
				StoredConfig: nil,
				NewConfig:    big.NewInt(31),
				RewindTo:     30,
			},
		},
	}

	for _, test := range tests {
		err := test.stored.CheckCompatible(test.new, test.head)
		if !reflect.DeepEqual(err, test.wantErr) {
			t.Errorf("error mismatch:\nstored: %v\nnew: %v\nhead: %v\nerr: %v\nwant: %v", test.stored, test.new, test.head, err, test.wantErr)
		}
	}
}

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
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), 2, 200, big.NewInt(4e18)},

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
		brioche := &BriocheConfig{
			BlockReward:       tc.reward,
			FirstHalvingBlock: big.NewInt(0),
			HalvingPeriod:     tc.period,
			HalvingTimes:      tc.times,
			HalvingRate:       tc.rate,
		}
		halved := brioche.calcHalvedReward(tc.reward, tc.past)
		if tc.expected.Cmp(halved) != 0 {
			t.Errorf("halved reward mismatched (expected=%v, actual=%v)", tc.expected, halved)
		}
	}
}

func TestGetBriocheBlockReward(t *testing.T) {
	defaultBlockReward := big.NewInt(1234e14)
	testcases := []struct {
		id            int32
		briocheConfig *BriocheConfig
		blockNum      *big.Int
		expected      *big.Int
	}{
		// nil case
		{
			id:            1,
			briocheConfig: nil,
			blockNum:      big.NewInt(100),
			expected:      defaultBlockReward,
		},

		// normal case
		{
			id: 2,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(5e17),
		},
		{
			id: 3,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(201),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(200),
			expected: big.NewInt(25e16),
		},

		// base block reward variations
		{
			id: 4,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(7e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(35e17),
		},
		{
			id: 5,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(3),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(1),
		},
		{
			id: 6,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(1),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(0),
		},
		{
			id: 7,
			briocheConfig: &BriocheConfig{
				BlockReward:       nil, // it will use the default block reward
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: new(big.Int).Div(defaultBlockReward, big.NewInt(2)),
		},

		// no halving
		{
			id: 8,
			briocheConfig: &BriocheConfig{
				BlockReward:       nil, // it will use the default block reward
				FirstHalvingBlock: nil,
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: defaultBlockReward,
		},
		{
			id: 9,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     nil,
				FinishRewardBlock: big.NewInt(101),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(10),
		},
		{
			id: 10,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(101),
				HalvingTimes:      0,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(10),
		},
		{
			id: 11,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(200),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(99), // not yet halving time
			expected: big.NewInt(10),
		},
		{
			id: 12,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(200),
				HalvingTimes:      10,
				HalvingRate:       100, // no halving rate
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(10),
		},

		// no reward case
		{
			id: 13,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(200),
				HalvingTimes:      10,
				HalvingRate:       0, // no reward
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(0),
		},
		{
			id: 14,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(200),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(199),
			expected: big.NewInt(5),
		},
		{
			id: 15,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(10),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(200),
				HalvingTimes:      10,
				HalvingRate:       50,
			},
			blockNum: big.NewInt(200),
			expected: big.NewInt(0),
		},

		// halving rate variations
		{
			id: 16,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(1000),
				HalvingTimes:      10,
				HalvingRate:       10,
			},
			blockNum: big.NewInt(100),
			expected: big.NewInt(1e17),
		},
		{
			id: 17,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(1000),
				HalvingTimes:      10,
				HalvingRate:       10,
			},
			blockNum: big.NewInt(200),
			expected: big.NewInt(1e16),
		},
		{
			id: 18,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(1000),
				HalvingTimes:      10,
				HalvingRate:       1,
			},
			blockNum: big.NewInt(300),
			expected: big.NewInt(1e12),
		},
		{
			id: 19,
			briocheConfig: &BriocheConfig{
				BlockReward:       big.NewInt(1e18),
				FirstHalvingBlock: big.NewInt(100),
				HalvingPeriod:     big.NewInt(100),
				FinishRewardBlock: big.NewInt(1000),
				HalvingTimes:      10,
				HalvingRate:       99,
			},
			blockNum: big.NewInt(300),
			expected: big.NewInt(970299e12),
		},
	}

	for _, tc := range testcases {
		actual := tc.briocheConfig.GetBriocheBlockReward(defaultBlockReward, tc.blockNum)
		if tc.expected.Cmp(actual) != 0 {
			t.Errorf("getBriocheReward mismatched (id=%d, expected=%v, actual=%v)", tc.id, tc.expected, actual)
		}
	}
}
