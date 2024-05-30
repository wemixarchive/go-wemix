package sim

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

// specified by EIP1967
// bytes32(uint256(keccak256('eip1967.proxy.implementation')) - 1)
const IMPLEMENTATION_SLOT = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"

type Client struct {
	Backend            *backends.SimulatedBackend
	ChainID            *big.Int
	Owner              common.Address
	ContractsByAlias   map[string]*Contract
	ContractsByAddress map[common.Address]*Contract
	Signer             bind.SignerFn
}

func NewClient(t *testing.T) *Client {
	owner := GetOrNewEOA("owner")

	MaxUint128 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 128), common.Big1)

	backend := backends.NewSimulatedBackend(
		core.GenesisAlloc{
			owner: {Balance: MaxUint128},
		},
		params.MaxGasLimit,
	)

	signer := types.LatestSignerForChainID(params.AllEthashProtocolChanges.ChainID)

	client := &Client{
		ContractsByAlias:   make(map[string]*Contract),
		ContractsByAddress: make(map[common.Address]*Contract),
		Owner:              owner,
		Backend:            backend,
		ChainID:            params.AllEthashProtocolChanges.ChainID,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), EOAKey[address])
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}

	return client
}

func (c *Client) SetOwner(owner common.Address) {
	c.Owner = owner
}

func (c *Client) EnrollContract(t *testing.T, contract *Contract) (logic *Contract) {
	c.ContractsByAlias[contract.Alias] = contract
	c.ContractsByAddress[contract.Address] = contract

	res, err := c.Backend.StorageAt(ctx, contract.Address, common.HexToHash(IMPLEMENTATION_SLOT), nil)
	require.NoError(t, err)
	logic = c.ContractsByAddress[common.BytesToAddress(res)]
	return
}

func (c *Client) FindContractByAddress(ca common.Address) *Contract {
	return c.ContractsByAddress[ca]
}

func (p *Client) SendTransaction(t *testing.T, sender common.Address, to common.Address, value *big.Int, payload []byte) (*types.Transaction, *types.Receipt) {
	tx := p.SendTransactionWithoutCommit(t, sender, to, value, payload)
	var (
		receipt *types.Receipt
		err     error
	)
	p.Backend.Commit()
	receipt, err = p.Backend.TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	return tx, receipt
}

func (p *Client) SendTransactionWithoutCommit(t *testing.T, sender common.Address, to common.Address, value *big.Int, payload []byte) (tx *types.Transaction) {
	if sender == (common.Address{}) {
		sender = p.Owner
	}

	nonce, err := p.Backend.PendingNonceAt(ctx, sender)
	require.NoError(t, err)

	head, err := p.Backend.HeaderByNumber(ctx, nil)
	require.NoError(t, err)

	tip, err := p.Backend.SuggestGasTipCap(ctx)
	require.NoError(t, err)

	cap := new(big.Int).Add(tip, new(big.Int).Mul(head.BaseFee, big.NewInt(2))) // tip + (head.baseFee*2)

	tx = types.NewTx(&types.DynamicFeeTx{
		To: func() *common.Address {
			if (to != common.Address{}) {
				return &to
			}
			return nil
		}(),
		Nonce:     nonce,
		GasFeeCap: cap,
		GasTipCap: tip,
		Gas:       defaultGasLimit,
		Value:     value,
		Data:      payload,
	})

	tx, err = p.Signer(sender, tx)
	require.NoError(t, err)

	err = p.Backend.SendTransaction(ctx, tx)
	require.NoError(t, err)
	return
}

func (p *Client) Commit() {
	p.Backend.Commit()
}

func (p *Client) CurrentTime(t *testing.T) uint64 {
	ctx := context.Background()
	header, err := p.Backend.HeaderByNumber(ctx, nil)
	require.NoError(t, err)
	return header.Time
}

func (p *Client) TimeAt(t *testing.T, blockNumber *big.Int) uint64 {
	block, err := p.Backend.BlockByNumber(context.Background(), blockNumber)
	require.NoError(t, err)
	return block.Time()
}

func (p *Client) AdjustTime(t *testing.T, adjustment time.Duration) {
	err := p.Backend.AdjustTime(adjustment)
	require.NoError(t, err)
	p.Backend.Commit()
}
