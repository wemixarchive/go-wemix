package sim

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
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

type RevertError struct {
	ABI    abi.Error
	Output interface{}
}

func (r *RevertError) Error() string {
	return fmt.Sprintf("%s: %s %v", vm.ErrExecutionReverted, r.ABI.Sig, r.Output)
}
