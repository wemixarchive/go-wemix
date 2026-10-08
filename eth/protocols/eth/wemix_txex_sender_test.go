package eth

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

// A TransactionsExMsg carries a From field next to each transaction. The
// receiver must derive the sender from the signature only, so a From that
// does not match the signer is ignored, whether or not the peer is a partner.

const (
	txexAccountBalance = 1_000_000_000_000_000_000
	txexTransferGas    = 21000
	txexOtherPrice     = 1_000_000_000
	txexSignerPrice    = 2_000_000_000 // above the 10% price bump, so it can replace
)

// Test-only keys built from small scalars.
var (
	txexOtherKey   = common.LeftPadBytes([]byte{0x01, 0x82}, 32)
	txexOtherAddr  = keyAddr(txexOtherKey)
	txexSignerKey  = common.LeftPadBytes([]byte{0x01, 0x83}, 32)
	txexSignerAddr = keyAddr(txexSignerKey)
)

func keyAddr(key []byte) common.Address {
	k, err := crypto.ToECDSA(key)
	if err != nil {
		panic(err)
	}
	return crypto.PubkeyToAddress(k.PublicKey)
}

// poolBackend forwards received transactions to a real txpool, the way
// eth.ethHandler does through the tx fetcher (pool.AddRemotes).
type poolBackend struct {
	*testBackend
}

func (b *poolBackend) AcceptTxs() bool { return true }

func (b *poolBackend) Handle(_ *Peer, packet Packet) error {
	txs, ok := packet.(*TransactionsPacket)
	if !ok {
		return nil
	}
	b.txpool.AddRemotesSync(*txs)
	return nil
}

func newPoolBackend(t *testing.T) *poolBackend {
	t.Helper()
	db := rawdb.NewMemoryDatabase()
	(&core.Genesis{
		Config: params.TestChainConfig,
		Alloc: core.GenesisAlloc{
			txexOtherAddr:  {Balance: big.NewInt(txexAccountBalance)},
			txexSignerAddr: {Balance: big.NewInt(txexAccountBalance)},
		},
	}).MustCommit(db)
	chain, err := core.NewBlockChain(db, nil, params.TestChainConfig, ethash.NewFaker(), vm.Config{}, nil, nil)
	if err != nil {
		t.Fatalf("new chain: %v", err)
	}
	txconfig := core.DefaultTxPoolConfig
	txconfig.Journal = ""
	b := &poolBackend{testBackend: &testBackend{
		db:     db,
		chain:  chain,
		txpool: core.NewTxPool(txconfig, params.TestChainConfig, chain),
	}}
	t.Cleanup(b.close)
	return b
}

func signTransfer(t *testing.T, key []byte, nonce uint64, price int64) *types.Transaction {
	t.Helper()
	k, err := crypto.ToECDSA(key)
	if err != nil {
		t.Fatalf("key: %v", err)
	}
	tx := types.NewTransaction(nonce, common.Address{0xaa}, big.NewInt(1), txexTransferGas, big.NewInt(price), nil)
	signed, err := types.SignTx(tx, types.LatestSigner(params.TestChainConfig), k)
	if err != nil {
		t.Fatalf("sign: %v", err)
	}
	return signed
}

// txexPacketMsg builds the wire bytes by hand: TransactionEx.EncodeRLP writes
// the tx's sender cache, not its From field, so making From differ from the
// signer requires writing the list [tx, from] directly, which
// TransactionEx.DecodeRLP reads back as {Tx: tx, From: from}.
func txexPacketMsg(t *testing.T, tx *types.Transaction, advertised common.Address) p2p.Msg {
	t.Helper()
	payload, err := rlp.EncodeToBytes([]interface{}{tx, advertised})
	if err != nil {
		t.Fatalf("encode: %v", err)
	}
	return p2p.Msg{Code: TransactionsExMsg, Size: uint32(len(payload)), Payload: bytes.NewReader(payload)}
}

// pinPartnerAndConsensus makes every peer look like a partner (or not) and
// runs handleTransactionsEx synchronously (PoW branch) so the pool state is
// settled when the handler returns. Both globals are restored on cleanup.
func pinPartnerAndConsensus(t *testing.T, partner bool) {
	t.Helper()
	prevPartner, prevMethod := wemixminer.IsPartnerFunc, params.ConsensusMethod
	t.Cleanup(func() {
		wemixminer.IsPartnerFunc = prevPartner
		params.ConsensusMethod = prevMethod
	})
	wemixminer.IsPartnerFunc = func(string) bool { return partner }
	params.ConsensusMethod = params.ConsensusPoW
}

func newTxexSenderPeer(t *testing.T, b Backend) *Peer {
	t.Helper()
	app, net := p2p.MsgPipe()
	var id enode.ID
	id[0] = 0x16
	peer := NewPeer(ETH66, p2p.NewPeer(id, "txex-sender", nil), net, b.TxPool())
	t.Cleanup(func() {
		peer.Close()
		_ = app.Close()
	})
	return peer
}

// TestTransactionsExIgnoresAdvertisedSender sends a transaction whose advertised
// From differs from its signer. The pool must file it under the signer, and must
// not disturb an unrelated pending tx filed under the advertised address.
func TestTransactionsExIgnoresAdvertisedSender(t *testing.T) {
	tests := []struct {
		name         string
		partner      bool
		otherPending bool // other already has nonce 0 pending
	}{
		{name: "partner signed sender", partner: true},
		{name: "partner signed sender replaces other tx", partner: true, otherPending: true},
		{name: "non-partner control", partner: false},
		{name: "non-partner control with other pending tx", partner: false, otherPending: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pinPartnerAndConsensus(t, tt.partner)
			b := newPoolBackend(t)
			peer := newTxexSenderPeer(t, b)

			var otherTx *types.Transaction
			if tt.otherPending {
				otherTx = signTransfer(t, txexOtherKey, 0, txexOtherPrice)
				if errs := b.txpool.AddRemotesSync([]*types.Transaction{otherTx}); errs[0] != nil {
					t.Fatalf("other tx rejected: %v", errs[0])
				}
			}

			signed := signTransfer(t, txexSignerKey, 0, txexSignerPrice)
			if err := handleTransactionsEx(b, txexPacketMsg(t, signed, txexOtherAddr), peer); err != nil {
				t.Fatalf("handleTransactionsEx: %v", err)
			}

			pending, _ := b.txpool.Content()
			if got := pending[txexSignerAddr]; len(got) != 1 || got[0].Hash() != signed.Hash() {
				t.Errorf("tx not pending under its signer %x: got %d txs", txexSignerAddr, len(got))
			}
			other := pending[txexOtherAddr]
			for _, tx := range other {
				if tx.Hash() == signed.Hash() {
					t.Errorf("tx %x is pending under the advertised From %x, not its signer", signed.Hash(), txexOtherAddr)
				}
			}
			if tt.otherPending && (len(other) != 1 || other[0].Hash() != otherTx.Hash()) {
				t.Errorf("unrelated pending tx under the advertised address was replaced or dropped (pending=%d)", len(other))
			}
			// The pooled object is the one decoded from the packet, so its
			// sender cache is what the miner will read when building a block.
			if pooled := b.txpool.Get(signed.Hash()); pooled != nil {
				if from, err := types.Sender(types.LatestSigner(params.TestChainConfig), pooled); err != nil || from != txexSignerAddr {
					t.Errorf("sender cache of pooled tx = %x (err %v), want signer %x", from, err, txexSignerAddr)
				}
			}
		})
	}
}
