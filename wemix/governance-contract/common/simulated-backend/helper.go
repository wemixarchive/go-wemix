package sim

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/stretchr/testify/require"
)

func parseABI(abiDefinition interface{}) (*abi.ABI, error) {
	s, ok := abiDefinition.(string)
	if !ok {
		if bytes, err := json.Marshal(abiDefinition); err != nil {
			return nil, err
		} else {
			s = string(bytes)
		}
	}
	if abi, err := abi.JSON(strings.NewReader(s)); err != nil {
		return nil, err
	} else {
		return &abi, nil
	}
}

func StringToBytes32(str string) [32]byte {
	dest := [32]byte{}
	copy(dest[:], []byte(str))
	return dest
}

func ABIEncode(types []string, data ...interface{}) ([]byte, error) {
	var args abi.Arguments
	for i := 0; i < len(types); i++ {
		if tpy, err := abi.NewType(types[i], "", nil); err != nil {
			return nil, err
		} else {
			args = append(args, abi.Argument{Type: tpy})
		}
	}
	return args.Pack(data...)
}

func ExpectedRevert(t *testing.T, err error, args ...interface{}) {
	require.Error(t, err)
	length := len(args)
	if revert, ok := err.(*RevertError); ok {
		if length > 0 {
			name, ok := args[0].(string)
			require.True(t, ok)
			require.Equal(t, revert.ABI.Name, name)
		}
		if length > 1 {
			output, ok := revert.Output.([]interface{})
			require.True(t, ok)
			for i := 1; i < length; i++ {
				arg := args[i]
				if arg != nil {
					require.Equal(t, args[i], output[i-1])
				}
			}
		}
	} else {
		if length > 0 {
			errStr := strings.TrimLeft(err.Error(), vm.ErrExecutionReverted.Error()+":")
			message, ok := args[0].(string)
			require.True(t, ok)
			require.Equal(t, strings.TrimSpace(message), strings.TrimSpace(errStr))
		}
	}
}
