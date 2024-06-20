package test

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

type allErrorsType map[[4]byte]abi.Error

var (
	allErrorsLock = sync.RWMutex{}
	allErrors     = allErrorsType{} // eventID => abi.Event
)

func collectErrors(errors map[string]abi.Error) error {
	for _, e := range errors {
		func() {
			allErrorsLock.Lock()
			defer allErrorsLock.Unlock()
			sig := [4]byte{}
			copy(sig[:], e.ID[:4])
			if _, ok := allErrors[sig]; !ok {
				allErrors[sig] = e
			}
		}()
	}
	return nil
}

type RevertError struct {
	ABI    abi.Error
	Output interface{}
}

func (r *RevertError) Error() string {
	return fmt.Sprintf("%s: %s %v", vm.ErrExecutionReverted, r.ABI.Sig, r.Output)
}

func UnpackError(err error) error {
	if err == nil {
		return nil
	}

	var result []byte
	if revert, ok := err.(interface {
		ErrorData() interface{}
	}); !ok {
		return err
	} else if data, ok := revert.ErrorData().(string); !ok {
		return err
	} else if bytes, decodeErr := hexutil.Decode(data); decodeErr != nil {
		return err
	} else {
		result = bytes
	}

	sig := [4]byte{}
	copy(sig[:], result[:4])
	if errABI, ok := allErrors[sig]; !ok {
		return err
	} else if output, unpackErr := errABI.Unpack(result); unpackErr != nil {
		return err
	} else {
		return &RevertError{errABI, output}
	}
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

func ToBytes32(str string) [32]byte {
	bytes := []byte(str)
	if len(bytes) > 32 {
		bytes = bytes[:32]
	}
	var copied = [32]byte{}
	copy(copied[:], bytes)
	return copied
}

var eoas = make(map[string]*bind.TransactOpts)

func getTxOpt(t *testing.T, alias string) *bind.TransactOpts {
	if eoa, ok := eoas[alias]; ok {
		return eoa
	} else {
		pk, err := crypto.GenerateKey()
		require.NoError(t, err)
		opts, err := bind.NewKeyedTransactorWithChainID(pk, params.AllEthashProtocolChanges.ChainID)
		require.NoError(t, err)
		eoas[alias] = opts
		return opts
	}
}

type IBackend interface {
	bind.ContractBackend
	bind.DeployBackend
	SuggestGasTipCap(context.Context) (*big.Int, error)
}

func CreateDynamicTx(backend IBackend, opts *bind.TransactOpts, to *common.Address, input []byte) (*types.Transaction, error) {
	// Normalize value
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	// Estimate TipCap
	gasTipCap := opts.GasTipCap
	if gasTipCap == nil {
		tip, err := backend.SuggestGasTipCap(ensureContext(opts.Context))
		if err != nil {
			return nil, err
		}
		gasTipCap = tip
	}
	// Estimate FeeCap
	gasFeeCap := opts.GasFeeCap
	if gasFeeCap == nil {
		gasFeeCap = new(big.Int).Add(gasTipCap, big.NewInt(1e9)) // 101gwei is recommended for maxFeeCap
	}
	if gasFeeCap.Cmp(gasTipCap) < 0 {
		return nil, fmt.Errorf("maxFeePerGas (%v) < maxPriorityFeePerGas (%v)", gasFeeCap, gasTipCap)
	}
	// Estimate GasLimit
	gasLimit := opts.GasLimit
	if opts.GasLimit == 0 {
		var err error
		gasLimit, err = estimateGasLimit(backend, opts, to, input, nil, gasTipCap, gasFeeCap, value)
		if err != nil {
			return nil, err
		}
	}
	// create the transaction
	nonce, err := getNonce(backend, opts)
	if err != nil {
		return nil, err
	}

	baseTx := &types.DynamicFeeTx{
		To:        to,
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		Value:     value,
		Data:      input,
	}

	return opts.Signer(opts.From, types.NewTx(baseTx))
}

func TransferCoin(backend IBackend, opts *bind.TransactOpts, value *big.Int, to *common.Address) (*types.Transaction, error) {
	opts.Value = value
	defer func() { opts.Value = nil }()
	if tx, err := CreateDynamicTx(backend, opts, to, []byte{}); err != nil {
		return nil, err
	} else {
		return tx, backend.SendTransaction(ensureContext(opts.Context), tx)
	}
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}

func estimateGasLimit(backend interface {
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error)
}, opts *bind.TransactOpts, to *common.Address, input []byte, gasPrice, gasTipCap, gasFeeCap, value *big.Int) (uint64, error) {
	msg := ethereum.CallMsg{
		From:      opts.From,
		To:        to,
		GasPrice:  gasPrice,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     value,
		Data:      input,
	}
	return backend.EstimateGas(ensureContext(opts.Context), msg)
}

func getNonce(backend interface {
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
}, opts *bind.TransactOpts) (uint64, error) {
	if opts.Nonce == nil {
		return backend.PendingNonceAt(ensureContext(opts.Context), opts.From)
	} else {
		return opts.Nonce.Uint64(), nil
	}
}
