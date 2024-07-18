// Copyright 2020 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package eth68

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/forkid"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/rlp"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
)

const (
	// softResponseLimit is the target maximum size of replies to data retrievals.
	softResponseLimit = 2 * 1024 * 1024

	// maxHeadersServe is the maximum number of block headers to serve. This number
	// is there to limit the number of disk lookups.
	maxHeadersServe = 1024

	// maxBodiesServe is the maximum number of block bodies to serve. This number
	// is mostly there to limit the number of disk lookups. With 24KB block sizes
	// nowadays, the practical limit will always be softResponseLimit.
	maxBodiesServe = 1024

	// maxReceiptsServe is the maximum number of block receipts to serve. This
	// number is mostly there to limit the number of disk lookups. With block
	// containing 200+ transactions nowadays, the practical limit will always
	// be softResponseLimit.
	maxReceiptsServe = 1024
)

// Handler is a callback to invoke from an outside runner after the boilerplate
// exchanges have passed.
type Handler func(peer *eth.Peer) error

// TxPool defines the methods needed by the protocol handler to serve transactions.
type TxPool interface {
	// Get retrieves the transaction from the local txpool with the given hash.
	Get(hash common.Hash) *types.Transaction
}

// MakeProtocols constructs the P2P protocol definitions for `eth`.
func MakeProtocols(backend eth.Backend, network uint64, dnsdisc enode.Iterator) []p2p.Protocol {
	protocols := make([]p2p.Protocol, 0, len(ProtocolVersions))
	for _, version := range ProtocolVersions {
		version := version // Closure

		protocols = append(protocols, p2p.Protocol{
			Name:    ProtocolName,
			Version: version,
			Length:  protocolLengths[version],
			Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
				peer := eth.NewPeer(version, p, rw, backend.TxPool())
				defer peer.Close()

				return backend.RunPeer(peer, func(peer *eth.Peer) error {
					return Handle(backend, peer)
				})
			},
			NodeInfo: func() interface{} {
				return nodeInfo(backend.Chain(), network)
			},
			PeerInfo: func(id enode.ID) interface{} {
				return backend.PeerInfo(id)
			},
			Attributes:     []enr.Entry{currentENREntry(backend.Chain())},
			DialCandidates: dnsdisc,
		})
	}
	return protocols
}

// nodeInfo retrieves some `eth` protocol metadata about the running host node.
func nodeInfo(chain *core.BlockChain, network uint64) *eth.NodeInfo {
	head := chain.CurrentBlock()
	hash := head.Hash()

	return &eth.NodeInfo{
		Network:    network,
		Difficulty: chain.GetTd(hash, head.NumberU64()),
		Genesis:    chain.Genesis().Hash(),
		Config:     chain.Config(),
		Head:       hash,
	}
}

// Handle is invoked whenever an `eth` connection is made that successfully passes
// the protocol handshake. This method will keep processing messages until the
// connection is torn down.
func Handle(backend eth.Backend, peer *eth.Peer) error {
	for {
		if err := handleMessage(backend, peer); err != nil {
			peer.Log().Debug("Message handling failed in `eth`", "err", err)
			return err
		}
	}
}

type msgHandler func(backend eth.Backend, msg Decoder, peer *eth.Peer) error
type Decoder interface {
	Decode(val interface{}) error
	Time() time.Time
}

var eth68 = map[uint64]msgHandler{
	NewBlockHashesMsg:             handleNewBlockhashes,
	NewBlockMsg:                   handleNewBlock,
	TransactionsMsg:               handleTransactions,
	NewPooledTransactionHashesMsg: handleNewPooledTransactionHashes,
	GetBlockHeadersMsg:            handleGetBlockHeaders,
	BlockHeadersMsg:               handleBlockHeaders,
	GetBlockBodiesMsg:             handleGetBlockBodies,
	BlockBodiesMsg:                handleBlockBodies,
	GetReceiptsMsg:                handleGetReceipts,
	ReceiptsMsg:                   handleReceipts,
	GetPooledTransactionsMsg:      handleGetPooledTransactions,
	PooledTransactionsMsg:         handlePooledTransactions,
}

// handleMessage is invoked whenever an inbound message is received from a remote
// peer. The remote connection is torn down upon returning any error.
func handleMessage(backend eth.Backend, peer *eth.Peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := peer.RW().ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > maxMessageSize {
		return fmt.Errorf("%w: %v > %v", errMsgTooLarge, msg.Size, maxMessageSize)
	}
	defer msg.Discard()

	var handlers = eth68

	// Track the amount of time it takes to serve the request and run the handler
	if metrics.Enabled {
		h := fmt.Sprintf("%s/%s/%d/%#02x", p2p.HandleHistName, ProtocolName, peer.Version(), msg.Code)
		defer func(start time.Time) {
			sampler := func() metrics.Sample {
				return metrics.ResettingSample(
					metrics.NewExpDecaySample(1028, 0.015),
				)
			}
			metrics.GetOrRegisterHistogramLazy(h, nil, sampler).Update(time.Since(start).Microseconds())
		}(time.Now())
	}
	if handler := handlers[msg.Code]; handler != nil {
		return handler(backend, msg, peer)
	}
	return fmt.Errorf("%w: %v", errInvalidMsgCode, msg.Code)
}

type enrEntry struct {
	ForkID forkid.ID // Fork identifier per EIP-2124

	// Ignore additional fields (for forward compatibility).
	Rest []rlp.RawValue `rlp:"tail"`
}

// currentENREntry constructs an `eth` ENR entry based on the current state of the chain.
func currentENREntry(chain *core.BlockChain) *enrEntry {
	head := chain.CurrentHeader()
	return &enrEntry{
		ForkID: forkid.NewID(chain.Config(), chain.Genesis().Hash(), head.Number.Uint64()),
	}
}

// ENRKey implements enr.Entry.
func (e enrEntry) ENRKey() string {
	return "eth"
}
