package bn

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

// MaxUint256 = ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
var MaxUint256 = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)
var MaxInt256 = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 255), common.Big1)
var MaxUint128 = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 128), common.Big1)

func New(x interface{}) *big.Int {
	return sanitize(x)
}

func News(args ...interface{}) []*big.Int {
	result := make([]*big.Int, len(args))
	for i, x := range args {
		result[i] = New(x)
	}
	return result
}

func sanitize(x interface{}) (y *big.Int) {
	switch v := x.(type) {
	case *big.Int:
		if y = v; y == nil || y.Sign() == 0 { // 0인경우 초기값으로 설정. assert에서 데이터타입이 달라 예외처리될수 있음.
			y = new(big.Int)
		}
	case big.Int:
		y = new(big.Int).Set(&v)
	case *decimal.Decimal:
		if v == nil {
			y = new(big.Int)
		} else {
			y = v.BigInt()
		}
	case decimal.Decimal:
		y = v.BigInt()
	case string:
		y, _ = new(big.Int).SetString(v, 10)
	case bool:
		if v {
			y = big.NewInt(1)
		} else {
			y = big.NewInt(0)
		}
	case uint64, uint32, uint16, uint8, uint:
		y = new(big.Int).SetUint64(reflect.ValueOf(v).Uint())
	case int64, int32, int16, int8, int:
		y = new(big.Int).SetInt64(reflect.ValueOf(v).Int())
	case float32, float64:
		y = new(big.Int).SetInt64(int64(reflect.ValueOf(v).Float()))
	case []byte:
		y = new(big.Int).SetBytes(v)
	case [20]byte, [32]byte, common.Hash, common.Address:
		value := reflect.ValueOf(v)
		slice := reflect.MakeSlice(reflect.TypeOf([]byte{}), value.Len(), value.Len())
		reflect.Copy(slice, value)
		y = new(big.Int).SetBytes(slice.Bytes())
	default:
		panic(fmt.Sprintf("sanitize %T", x))
	}
	return
}

// z = a+b
func Add(x, y interface{}) (z *big.Int) {
	return new(big.Int).Add(sanitize(x), sanitize(y))
}

// z = x-y
func Sub(x, y interface{}) (z *big.Int) {
	return new(big.Int).Sub(sanitize(x), sanitize(y))
}

// z = x*y
func Mul(x, y interface{}) (z *big.Int) {
	return new(big.Int).Mul(sanitize(x), sanitize(y))
}

// z = x/y
func Div(x, y interface{}) (z *big.Int) {
	return new(big.Int).Div(sanitize(x), sanitize(y))
}

// z = x%y
func Mod(x, y interface{}) (z *big.Int) {
	return new(big.Int).Mod(sanitize(x), sanitize(y))
}

// z = x**y
func Exp(x, y interface{}) (z *big.Int) {
	z = new(big.Int).Exp(sanitize(x), sanitize(y), nil)
	return
}

// z = |x|
func Abs(x interface{}) (z *big.Int) {
	z = new(big.Int).Abs(sanitize(x))
	return
}

// z = ⌊√x⌋
func Sqrt(x interface{}) (z *big.Int) {
	z = new(big.Int).Sqrt(sanitize(x))
	return
}

// z = ⌊√(x*y)⌊
func MulSqrt(x, y interface{}) (z *big.Int) {
	return Sqrt(Mul(x, y))
}

// -1, 0, 1
func Sign(x interface{}) int {
	return sanitize(x).Sign()
}

// z = (m1*m2)/d
func MulDiv(m1, m2, d interface{}) (z *big.Int) {
	z = Div(Mul(m1, m2), d)
	return
}

// (e1**e2)*m
func ExpMul(e1, e2, m interface{}) (z *big.Int) {
	z = Mul(Exp(e1, e2), m)
	return
}

// ether to wei by fixed 18 decimals
func ToWei18(ether interface{}) *big.Int {
	return ToWeiN(ether, 18)
}

// wei to ether by fixed 18 decimals
func ToEther18(wei interface{}) decimal.Decimal {
	return ToEtherN(wei, 18)
}

// ToWeiN returns x * 10 ^ N decimals
func ToWeiN(x interface{}, decimals uint8) *big.Int {
	var y decimal.Decimal
	switch v := x.(type) {
	case *big.Int:
		if v == nil {
			y = decimal.New(0, 0)
		} else {
			y = decimal.NewFromBigInt(v, 0)
		}
	case big.Int:
		y = decimal.NewFromBigInt(&v, 0)
	case string:
		y, _ = decimal.NewFromString(v)
	case uint64, uint32, uint16, uint8, uint:
		y = decimal.NewFromBigInt(New(reflect.ValueOf(v).Uint()), 0)
	case int64, int32, int16, int8, int:
		y = decimal.NewFromInt(reflect.ValueOf(v).Int())
	case float32, float64:
		y = decimal.NewFromFloat(reflect.ValueOf(v).Float())
	case decimal.Decimal:
		y = v
	case *decimal.Decimal:
		if v == nil {
			y = decimal.New(0, 0)
		} else {
			y = *v
		}
	case []byte:
		y = decimal.NewFromBigInt(new(big.Int).SetBytes(v), 0)
	case [20]byte, [32]byte, common.Hash, common.Address:
		value := reflect.ValueOf(v)
		slice := reflect.MakeSlice(reflect.TypeOf([]byte{}), value.Len(), value.Len())
		reflect.Copy(slice, value)
		y = decimal.NewFromBigInt(New(slice.Bytes()), 0)
	default:
		panic(fmt.Sprintf("sanitize %T", x))
	}

	mul := decimal.NewFromFloat(10).Pow(decimal.NewFromFloat(float64(decimals)))
	return New(y.Mul(mul))
}

// ToEtherN returns x / 10 ^ N decimals
func ToEtherN(x interface{}, decimals uint8) decimal.Decimal {
	denominator := decimal.NewFromFloat(10).Pow(decimal.NewFromFloat(float64(decimals)))
	return decimal.NewFromBigInt(sanitize(x), 0).Div(denominator)
}

// z = params sum / len(params)
func Avg(params ...interface{}) (z *big.Int) {
	sum := sanitize(0)
	for _, e := range params {
		sum.Add(sum, sanitize(e))
	}
	z = Div(sum, len(params))
	return
}

// -1 if x <  y
// 0 if x == y
// +1 if x >  y
func Cmp(x, y interface{}) int {
	return sanitize(x).Cmp(sanitize(y))
}

func Min(x, y interface{}) *big.Int {
	if Cmp(x, y) < 0 {
		return sanitize(x)
	}
	return sanitize(y)
}

func Max(x, y interface{}) *big.Int {
	if Cmp(x, y) > 0 {
		return sanitize(x)
	}
	return sanitize(y)
}

// compare x and y with margin
func CmpMargin(x, y interface{}, margin interface{}) int {
	if r := Abs(Sub(x, y)).Cmp(New(margin)); r < 0 {
		return 0
	} else {
		return Cmp(x, y)
	}
}

// z = params sum
func Sum(params ...interface{}) (z *big.Int) {
	z = sanitize(0)
	for _, e := range params {
		z.Add(z, sanitize(e))
	}
	return
}

// *big.Int => *hexutil.Big
func ToHexBig(x interface{}) *hexutil.Big {
	if x == nil {
		return nil
	}
	hbn := hexutil.Big(*sanitize(x))
	return &hbn
}

// []*big.Int => []*hexutil.Big
func ToHexBigs(x ...interface{}) (bignums []*hexutil.Big) {
	for _, v := range x {
		bignums = append(bignums, ToHexBig(v))
	}
	return
}
