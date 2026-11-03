// Copyright 2026 The go-wemix Authors

package wemix

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
)

var errHeaderNotFound = errors.New("header not found")

// offlineChainAPI serves, straight from a BlockChain, the eth namespace
// methods wemixAdmin reads chain and governance state through. Results match
// what the same methods return on a running node, so the admin code cannot
// tell the two apart.
type offlineChainAPI struct {
	chain *core.BlockChain
}

func (api *offlineChainAPI) headerByNumber(number rpc.BlockNumber) *types.Header {
	if number < 0 {
		// latest, pending and the like all mean the current head offline.
		return api.chain.CurrentHeader()
	}
	return api.chain.GetHeaderByNumber(uint64(number))
}

func (api *offlineChainAPI) headerByNumberOrHash(blockNrOrHash rpc.BlockNumberOrHash) *types.Header {
	if number, ok := blockNrOrHash.Number(); ok {
		return api.headerByNumber(number)
	}
	hash, _ := blockNrOrHash.Hash()
	return api.chain.GetHeaderByHash(hash)
}

func (api *offlineChainAPI) stateAndHeader(blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	header := api.headerByNumberOrHash(blockNrOrHash)
	if header == nil {
		return nil, nil, errHeaderNotFound
	}
	statedb, err := api.chain.StateAt(header.Root)
	if err != nil {
		return nil, nil, err
	}
	return statedb, header, nil
}

// GetBlockByNumber returns the header fields of the block, or nil if the
// chain does not have it. Transactions are never included.
func (api *offlineChainAPI) GetBlockByNumber(ctx context.Context, number rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	if header := api.headerByNumber(number); header != nil {
		return ethapi.RPCMarshalHeader(header), nil
	}
	return nil, nil
}

// GetBlockByHash returns the header fields of the block, or nil if the chain
// does not have it. Transactions are never included.
func (api *offlineChainAPI) GetBlockByHash(ctx context.Context, hash common.Hash, fullTx bool) (map[string]interface{}, error) {
	if header := api.chain.GetHeaderByHash(hash); header != nil {
		return ethapi.RPCMarshalHeader(header), nil
	}
	return nil, nil
}

// GetCode returns the contract code at the given block.
func (api *offlineChainAPI) GetCode(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	statedb, _, err := api.stateAndHeader(blockNrOrHash)
	if err != nil {
		return nil, err
	}
	return statedb.GetCode(address), statedb.Error()
}

// Call executes a read-only message call on the state of the given block
// with the same gas cap and timeout a node uses by default.
func (api *offlineChainAPI) Call(ctx context.Context, args ethapi.TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	statedb, header, err := api.stateAndHeader(blockNrOrHash)
	if err != nil {
		return nil, err
	}
	msg, err := args.ToMessage(ethconfig.Defaults.RPCGasCap, header.BaseFee)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, ethconfig.Defaults.RPCEVMTimeout)
	defer cancel()

	blockCtx := core.NewEVMBlockContext(header, api.chain, nil)
	evm := vm.NewEVM(blockCtx, core.NewEVMTxContext(msg), statedb, api.chain.Config(), vm.Config{NoBaseFee: true})
	go func() {
		<-ctx.Done()
		evm.Cancel()
	}()
	result, err := core.ApplyMessage(evm, msg, new(core.GasPool).AddGas(math.MaxUint64))
	if evm.Cancelled() {
		return nil, fmt.Errorf("execution aborted (timeout = %v)", ethconfig.Defaults.RPCEVMTimeout)
	}
	if err != nil {
		return nil, err
	}
	return result.Return(), result.Err
}

// StartOfflineAdmin sets up the admin for commands that process blocks
// without a running node, such as import. The admin talks to an in-process
// RPC server that answers from chain, in place of the node's own RPC, so
// block signatures, rewards and base fees are checked the same way a running
// node checks them. Only the methods governance reads need are served; any
// other call fails with a method-not-found error.
//
// The returned function clears the admin and shuts the server down. It
// returns ErrAlreadyRunning if an admin is already set up.
func StartOfflineAdmin(chain *core.BlockChain) (func(), error) {
	if admin != nil {
		return nil, ErrAlreadyRunning
	}
	server := rpc.NewServer()
	if err := server.RegisterName("eth", &offlineChainAPI{chain: chain}); err != nil {
		return nil, err
	}
	rpcCli := rpc.DialInProc(server)
	stop := func() {
		rpcCli.Close()
		server.Stop()
	}

	ma := &wemixAdmin{
		lock:   &sync.Mutex{},
		rpcCli: rpcCli,
		cli:    ethclient.NewClient(rpcCli),
	}
	var err error
	if ma.bootNodeId, ma.bootAccount, err = ma.getGenesisInfo(); err != nil {
		stop()
		return nil, err
	}
	admin = ma
	return func() {
		admin = nil
		stop()
	}, nil
}
