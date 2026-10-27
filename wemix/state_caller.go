// state_caller.go

package wemix

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
)

// same gas as eth_call when no gas cap is configured
const stateCallGas = uint64(math.MaxUint64 / 2)

// blockStateCaller runs contract calls against the state of a fixed block,
// configured like eth_call at that block.
type blockStateCaller struct {
	header *types.Header
	state  *state.StateDB
	evm    *vm.EVM
}

// newBlockStateCallerAtParent opens the parent state of header from db.
func newBlockStateCallerAtParent(chain consensus.ChainHeaderReader, header *types.Header, db state.Database) (*blockStateCaller, error) {
	if header.Number == nil || header.Number.Sign() <= 0 {
		return nil, consensus.ErrUnknownAncestor
	}
	number := header.Number.Uint64() - 1
	parent := chain.GetHeader(header.ParentHash, number)
	if parent == nil {
		// core.GenerateChain hands over a chain reader without headers
		if reader, ok := db.TrieDB().DiskDB().(ethdb.Reader); ok {
			parent = rawdb.ReadHeader(reader, header.ParentHash, number)
		}
		if parent == nil {
			return nil, consensus.ErrUnknownAncestor
		}
	}
	statedb, err := state.New(parent.Root, db, nil)
	if err != nil {
		return nil, err
	}
	return newBlockStateCaller(chain, parent, statedb), nil
}

func newBlockStateCaller(chain consensus.ChainHeaderReader, header *types.Header, statedb *state.StateDB) *blockStateCaller {
	config := chain.Config()
	blockCtx := core.NewEVMBlockContext(header, headerChainContext{chain}, &header.Coinbase)
	return &blockStateCaller{
		header: header,
		state:  statedb,
		evm:    vm.NewEVM(blockCtx, vm.TxContext{GasPrice: new(big.Int)}, statedb, config, vm.Config{NoBaseFee: true}),
	}
}

// headerChainContext satisfies core.ChainContext for BLOCKHASH lookups only;
// Engine is never called because the author is given explicitly.
type headerChainContext struct {
	consensus.ChainHeaderReader
}

func (headerChainContext) Engine() consensus.Engine { return nil }

func (c *blockStateCaller) checkBlock(blockNumber *big.Int) error {
	if blockNumber == nil {
		return errors.New("block number is required")
	}
	if blockNumber.Cmp(c.header.Number) != 0 {
		return fmt.Errorf("block number %v does not match state block %v", blockNumber, c.header.Number)
	}
	return nil
}

func (c *blockStateCaller) CodeAt(_ context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	if err := c.checkBlock(blockNumber); err != nil {
		return nil, err
	}
	return c.state.GetCode(contract), nil
}

func (c *blockStateCaller) CallContract(_ context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	if err := c.checkBlock(blockNumber); err != nil {
		return nil, err
	}
	if call.To == nil {
		return nil, errors.New("missing contract address")
	}
	ret, _, err := c.evm.StaticCall(vm.AccountRef(call.From), *call.To, call.Data, stateCallGas)
	return ret, err
}
