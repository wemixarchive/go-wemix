package types

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

// TestTxExs2TxsDropsAdvertisedSender checks that a TransactionEx decoded from
// the wire never seeds the sender cache with its advertised From.
func TestTxExs2TxsDropsAdvertisedSender(t *testing.T) {
	key, err := crypto.ToECDSA(common.LeftPadBytes([]byte{0x01, 0x83}, 32))
	if err != nil {
		t.Fatalf("key: %v", err)
	}
	signerAddr := crypto.PubkeyToAddress(key.PublicKey)
	signer := LatestSignerForChainID(big.NewInt(1))
	tx, err := SignTx(NewTransaction(0, common.Address{0xaa}, big.NewInt(1), 21000, big.NewInt(1), nil), signer, key)
	if err != nil {
		t.Fatalf("sign: %v", err)
	}

	tests := []struct {
		name       string
		advertised common.Address
	}{
		{name: "advertised differs from signer", advertised: common.Address{0xde, 0xad}},
		{name: "zero sender", advertised: common.Address{}},
		{name: "advertised matches signer", advertised: signerAddr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Wire bytes built by hand so From can be set independently of the signer.
			raw, err := rlp.EncodeToBytes([]interface{}{tx, tt.advertised})
			if err != nil {
				t.Fatalf("encode: %v", err)
			}
			var txexs []*TransactionEx
			if err := rlp.DecodeBytes(raw, &txexs); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if txexs[0].From != tt.advertised {
				t.Fatalf("decoded From = %x, want %x", txexs[0].From, tt.advertised)
			}
			txs := TxExs2Txs(txexs)
			if GetSender(signer, txs[0]) != nil {
				t.Fatal("sender cache was seeded from peer input")
			}
			from, err := Sender(signer, txs[0])
			if err != nil || from != signerAddr {
				t.Fatalf("Sender = %x (err %v), want %x", from, err, signerAddr)
			}
		})
	}
}
