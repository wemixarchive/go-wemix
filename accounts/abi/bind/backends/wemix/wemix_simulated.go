package wemix_backends

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/wemix"
	gov "github.com/ethereum/go-ethereum/wemix/governance-contract/bind"
)

type SimClient interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.LogFilterer
	ethereum.PendingStateReader
	ethereum.PendingContractCaller
	ethereum.TransactionReader
	ethereum.TransactionSender
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	GovContracts() *gov.GovContracts
	Commit()
}

var ChainID = params.AllEthashProtocolChanges.ChainID

type wemixSimulatedBackend struct {
	stack      *node.Node
	eth        *eth.Ethereum
	backend    *backends.SimulatedBackend
	governance *gov.GovContracts
}

func NewWemixSimulatedBackend(pk *ecdsa.PrivateKey, datadir string, alloc core.GenesisAlloc, options ...func(nodeConf *node.Config, ethConf *ethconfig.Config, envConfig *gov.EnvInitializeConfig)) (SimClient, error) {
	params.ConsensusMethod = params.ConsensusPoA

	nodeConfig := node.DefaultConfig
	nodeConfig.DataDir = datadir
	nodeConfig.KeyStoreDir = filepath.Join(datadir, "keystore")
	nodeConfig.P2P = p2p.Config{
		PrivateKey:  pk,
		NoDiscovery: true,
	}

	ethConfig := ethconfig.Defaults
	ethConfig.Genesis = &core.Genesis{
		Config:   params.AllEthashProtocolChanges,
		GasLimit: ethconfig.Defaults.Miner.GasCeil,
		Alloc:    alloc,
		Coinbase: crypto.PubkeyToAddress(pk.PublicKey),
	}
	ethConfig.SyncMode = downloader.FullSync
	ethConfig.TxPool.NoLocals = true
	ethConfig.Ethash = ethash.Config{PowMode: ethash.ModeFake, Log: log.Root()}
	ethConfig.NoPruning = true
	envConfig := gov.DefaultEnvInitializeConfig

	for _, option := range options {
		option(&nodeConfig, &ethConfig, &envConfig)
	}

	stack, err := node.New(&nodeConfig)
	if err != nil {
		return nil, err
	}

	extraData := strings.Split(strings.TrimLeft(stack.Server().NodeInfo().Enode, "enode://"), "@")[0]
	ethConfig.Genesis.ExtraData = append(ethConfig.Genesis.ExtraData, []byte("0x"+extraData)...)
	ethConfig.NetworkId = ethConfig.Genesis.Config.ChainID.Uint64()

	backend, err := eth.New(stack, &ethConfig)
	if err != nil {
		return nil, err
	}
	ethConfig.Genesis.MustCommit(backend.ChainDb())

	if err := stack.Start(); err != nil {
		return nil, err
	}

	wemix.StartAdmin(stack, nodeConfig.DataDir)

	ks := keystore.NewPlaintextKeyStore(nodeConfig.KeyStoreDir)
	account, err := ks.ImportECDSA(pk, "")
	if err != nil {
		return nil, err
	}
	if err := ks.Unlock(account, ""); err != nil {
		return nil, err
	}

	backend.AccountManager().AddBackend(ks)
	backend.TxPool().SetGasPrice(ethConfig.Miner.GasPrice)
	if err := backend.StartMining(1); err != nil {
		return nil, err
	}
	sleepCount := 0
	for backend.APIBackend.CurrentBlock().NumberU64() == 0 {
		sleepCount++
		if sleepCount > 10 {
			return nil, errors.New("mining error")
		}
		time.Sleep(0.2e9)
	}

	rpcClient, err := stack.Attach()
	if err != nil {
		return nil, err
	}
	ethClient := ethclient.NewClient(rpcClient)

	opts, err := bind.NewKeyedTransactorWithChainID(pk, ethConfig.Genesis.Config.ChainID)
	if err != nil {
		return nil, err
	}

	contracts, err := gov.DeployGovContracts(opts, ethClient)
	if err != nil {
		return nil, err
	}

	contracts, err = gov.ExecuteInitialize(contracts, opts, ethClient, &envConfig, func(Gov gov.IGovInitFuncs) (*types.Transaction, error) {
		nodeInfo := stack.Server().NodeInfo()
		if enode, err := hex.DecodeString(extraData); err != nil {
			return nil, err
		} else {
			return Gov.Init(opts, contracts.Registry.Address(), envConfig.STAKING_MIN, []byte(nodeInfo.Name), enode, []byte(nodeInfo.IP), big.NewInt(8589))
		}
	})
	if err != nil {
		return nil, err
	}

	return &wemixSimulatedBackend{
		stack:      stack,
		eth:        backend,
		backend:    backends.NewSimulatedBackendWithEthereum(backend),
		governance: contracts,
	}, nil
}

func (w *wemixSimulatedBackend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return w.backend.BlockByHash(ctx, hash)
}
func (w *wemixSimulatedBackend) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return w.backend.BlockByNumber(ctx, number)
}
func (w *wemixSimulatedBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return w.backend.HeaderByHash(ctx, hash)
}
func (w *wemixSimulatedBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return w.backend.HeaderByNumber(ctx, number)
}
func (w *wemixSimulatedBackend) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	return w.backend.TransactionCount(ctx, blockHash)
}
func (w *wemixSimulatedBackend) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	return w.backend.TransactionInBlock(ctx, blockHash, index)
}
func (w *wemixSimulatedBackend) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return w.backend.SubscribeNewHead(ctx, ch)
}
func (w *wemixSimulatedBackend) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return w.backend.BalanceAt(ctx, account, blockNumber)
}
func (w *wemixSimulatedBackend) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	return w.backend.StorageAt(ctx, account, key, blockNumber)
}
func (w *wemixSimulatedBackend) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return w.backend.CodeAt(ctx, account, blockNumber)
}
func (w *wemixSimulatedBackend) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return w.backend.NonceAt(ctx, account, blockNumber)
}
func (w *wemixSimulatedBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return w.backend.CallContract(ctx, call, blockNumber)
}
func (w *wemixSimulatedBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return w.backend.EstimateGas(ctx, call)
}
func (w *wemixSimulatedBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return w.backend.SuggestGasPrice(ctx)
}
func (w *wemixSimulatedBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return w.backend.SuggestGasTipCap(ctx)
}
func (w *wemixSimulatedBackend) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return w.backend.FilterLogs(ctx, q)
}
func (w *wemixSimulatedBackend) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return w.backend.SubscribeFilterLogs(ctx, q, ch)
}
func (w *wemixSimulatedBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return w.backend.PendingCodeAt(ctx, account)
}
func (w *wemixSimulatedBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return w.backend.PendingNonceAt(ctx, account)
}
func (w *wemixSimulatedBackend) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	return w.backend.PendingCallContract(ctx, call)
}
func (w *wemixSimulatedBackend) TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return w.backend.TransactionByHash(ctx, txHash)
}
func (w *wemixSimulatedBackend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	receipt, err := w.backend.TransactionReceipt(ctx, txHash)
	if err == ethereum.NotFound {
		w.backend.Commit()
		return w.TransactionReceipt(ctx, txHash)
	}
	return receipt, err
}
func (w *wemixSimulatedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return w.backend.SendTransaction(ctx, tx)
}
func (w *wemixSimulatedBackend) Commit() {
	w.backend.Commit()
}
func (w *wemixSimulatedBackend) GovContracts() *gov.GovContracts {
	return w.governance
}

func (w *wemixSimulatedBackend) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	// return w.backend.PendingBalanceAt(ctx, account)
	return nil, errors.New("PendingBalanceAt")
}
func (w *wemixSimulatedBackend) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	// return w.backend.PendingStorageAt(ctx, account, key)
	return nil, errors.New("PendingStorageAt")
}
func (w *wemixSimulatedBackend) PendingTransactionCount(ctx context.Context) (uint, error) {
	// return w.backend.PendingTransactionCount(ctx)
	return 0, errors.New("PendingTransactionCount")
}
