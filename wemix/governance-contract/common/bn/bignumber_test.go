package bn

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestToBigInt(t *testing.T) {
	t.Run("case big.Int", func(t *testing.T) {
		require.Equal(t, big.NewInt(10), New(big.NewInt(10)))
		require.Equal(t, big.NewInt(10), New(*big.NewInt(10)))
	})
	t.Run("case decimal.Decimal", func(t *testing.T) {
		actual := decimal.NewFromInt(10)
		require.Equal(t, big.NewInt(10), New(&actual))
		// case decimal.Decimal:
		require.Equal(t, big.NewInt(10), New(actual))
	})
	t.Run("case string", func(t *testing.T) {
		require.Equal(t, big.NewInt(10), New("10"))
	})
	t.Run("case bool", func(t *testing.T) {
		require.Equal(t, []*big.Int{big.NewInt(0), big.NewInt(1)}, News(false, true))
	})
	t.Run("case integer", func(t *testing.T) {
		// case uint64, uint32, uint16, uint8, uint:
		expected := []*big.Int{big.NewInt(10), big.NewInt(10), big.NewInt(10), big.NewInt(10), big.NewInt(10)}
		require.Equal(t, expected, News(uint64(10), uint32(10), uint16(10), uint8(10), uint(10)))
		// case int64, int32, int16, int8, int:
		require.Equal(t, expected, News(int64(10), int32(10), int16(10), int8(10), int(10)))
	})
	t.Run("case float", func(t *testing.T) {
		// case float32, float64:
		require.Equal(t, []*big.Int{big.NewInt(10), big.NewInt(10)}, News(float32(10), float64(10)))
	})
	t.Run("case bytes", func(t *testing.T) {
		// []byte
		bytesSlice := []byte{0xff, 0xff}
		require.Equal(t, new(big.Int).SetBytes(bytesSlice), New(bytesSlice))
		// [20]byte,
		bytes20 := bytes.Repeat([]byte{0xff}, 20)
		require.Equal(t, new(big.Int).SetBytes(bytes20), New(bytes20))
		// [32]byte
		bytes32 := bytes.Repeat([]byte{0xff}, 32)
		require.Equal(t, new(big.Int).SetBytes(bytes32), New(bytes32))
		// common.Hash
		hash := common.BytesToHash(bytes32)
		require.Equal(t, new(big.Int).SetBytes(hash[:]), New(hash))
		// common.Address
		address := common.BytesToAddress(bytes20)
		require.Equal(t, new(big.Int).SetBytes(address[:]), New(address))
	})
	t.Run("case unsupported type", func(t *testing.T) {
		defer func() {
			require.NotNil(t, recover(), "didn't panic for unsupported type")
		}()
		New([21]byte{})
	})
}

func TestToWei(t *testing.T) {
	t.Run("case big.Int", func(t *testing.T) {
		require.Equal(t, big.NewInt(10_000), ToWeiN(big.NewInt(10), 3))
		require.Equal(t, big.NewInt(10_000), ToWeiN(*big.NewInt(10), 3))
	})
	t.Run("case decimal.Decimal", func(t *testing.T) {
		actual := decimal.NewFromInt(10)
		require.Equal(t, big.NewInt(10_000), ToWeiN(&actual, 3))
		// case decimal.Decimal:
		require.Equal(t, big.NewInt(10_000), ToWeiN(actual, 3))
	})
	t.Run("case string", func(t *testing.T) {
		require.Equal(t, big.NewInt(10_000), ToWeiN("10", 3))
	})
	t.Run("case integer", func(t *testing.T) {
		// case uint64, uint32, uint16, uint8, uint:
		expected := []*big.Int{big.NewInt(10_000), big.NewInt(10_000), big.NewInt(10_000), big.NewInt(10_000), big.NewInt(10_000)}
		require.Equal(t, expected, []*big.Int{ToWeiN(uint64(10), 3), ToWeiN(uint32(10), 3), ToWeiN(uint16(10), 3), ToWeiN(uint8(10), 3), ToWeiN(uint(10), 3)})
		// case int64, int32, int16, int8, int:
		require.Equal(t, expected, []*big.Int{ToWeiN(int64(10), 3), ToWeiN(int32(10), 3), ToWeiN(int16(10), 3), ToWeiN(int8(10), 3), ToWeiN(int(10), 3)})
	})
	t.Run("case float", func(t *testing.T) {
		// case float32, float64:
		require.Equal(t, []*big.Int{big.NewInt(100), big.NewInt(100)}, []*big.Int{ToWeiN(float32(0.1), 3), ToWeiN(float64(0.1), 3)})
	})
	t.Run("case bytes", func(t *testing.T) {
		// []byte
		bytesSlice := []byte{0xff, 0xff}
		require.Equal(t, new(big.Int).SetBytes(bytesSlice), ToWeiN(bytesSlice, 0))
		// [20]byte,
		bytes20 := bytes.Repeat([]byte{0xff}, 20)
		require.Equal(t, new(big.Int).SetBytes(bytes20), ToWeiN(bytes20, 0))
		// [32]byte
		bytes32 := bytes.Repeat([]byte{0xff}, 32)
		require.Equal(t, new(big.Int).SetBytes(bytes32), ToWeiN(bytes32, 0))
		// common.Hash
		hash := common.BytesToHash(bytes32)
		require.Equal(t, new(big.Int).SetBytes(hash[:]), ToWeiN(hash, 0))
		// common.Address
		address := common.BytesToAddress(bytes20)
		require.Equal(t, new(big.Int).SetBytes(address[:]), ToWeiN(address, 0))
	})
	t.Run("case unsupported type", func(t *testing.T) {
		defer func() {
			require.NotNil(t, recover(), "didn't panic for unsupported type")
		}()
		ToWeiN([21]byte{}, 2)
	})
}
