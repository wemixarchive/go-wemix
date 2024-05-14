package wemix

import (
	"math/big"
	"testing"
)

func TestHalveRewards(t *testing.T) {
	testcases := []struct {
		reward   *big.Int
		period   *big.Int
		past     *big.Int
		expected *big.Int
	}{
		// sample test
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(0), big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(99), big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(100), big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(100), big.NewInt(101), big.NewInt(25e16)},

		// brioche halving test
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(0), big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 - 1), big.NewInt(5e17)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200), big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200*2 - 1), big.NewInt(25e16)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 2), big.NewInt(125e15)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 3), big.NewInt(625e14)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 4), big.NewInt(3125e13)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 5), big.NewInt(15625e12)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 6), big.NewInt(78125e11)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 7), big.NewInt(390625e10)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 8), big.NewInt(1953125e9)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 9), big.NewInt(9765625e8)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 10), big.NewInt(48828125e7)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 11), big.NewInt(244140625e6)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 12), big.NewInt(1220703125e5)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 13), big.NewInt(6103515625e4)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 14), big.NewInt(30517578125e3)},
		{big.NewInt(1e18), big.NewInt(63115200), big.NewInt(63115200 * 15), big.NewInt(152587890625e2)},
	}

	for _, testcase := range testcases {
		halved := halveRewards(testcase.reward, testcase.period, testcase.past)
		if testcase.expected.Cmp(halved) != 0 {
			t.Errorf("halveRewards mismatched (expected=%v, actual=%v)", testcase.expected, halved)
		}
	}
}
