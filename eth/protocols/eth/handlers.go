// Copyright 2021 The go-ethereum Authors
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

package eth

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

// handleGetBlockHeaders handles Block header query, collect the requested headers and reply
func handleGetBlockHeaders(backend Backend, msg Decoder, peer *Peer) error {
	// Decode the complex header query
	var query GetBlockHeadersPacket
	if err := msg.Decode(&query); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		response := answerGetBlockHeadersQuery(backend, &query, peer)
		return peer.SendBlockHeaders(response)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func answerGetBlockHeadersQuery(backend Backend, query *GetBlockHeadersPacket, peer *Peer) []*types.Header {
	hashMode := query.Origin.Hash != (common.Hash{})
	first := true
	maxNonCanonical := uint64(100)

	// Gather headers until the fetch or network limits is reached
	var (
		bytes   common.StorageSize
		headers []*types.Header
		unknown bool
		lookups int
	)
	for !unknown && len(headers) < int(query.Amount) && bytes < softResponseLimit &&
		len(headers) < maxHeadersServe && lookups < 2*maxHeadersServe {
		lookups++
		// Retrieve the next header satisfying the query
		var origin *types.Header
		if hashMode {
			if first {
				first = false
				origin = backend.Chain().GetHeaderByHash(query.Origin.Hash)
				if origin != nil {
					query.Origin.Number = origin.Number.Uint64()
				}
			} else {
				origin = backend.Chain().GetHeader(query.Origin.Hash, query.Origin.Number)
			}
		} else {
			origin = backend.Chain().GetHeaderByNumber(query.Origin.Number)
		}
		if origin == nil {
			break
		}
		headers = append(headers, origin)
		bytes += estHeaderSize

		// Advance to the next header of the query
		switch {
		case hashMode && query.Reverse:
			// Hash based traversal towards the genesis block
			ancestor := query.Skip + 1
			if ancestor == 0 {
				unknown = true
			} else {
				query.Origin.Hash, query.Origin.Number = backend.Chain().GetAncestor(query.Origin.Hash, query.Origin.Number, ancestor, &maxNonCanonical)
				unknown = (query.Origin.Hash == common.Hash{})
			}
		case hashMode && !query.Reverse:
			// Hash based traversal towards the leaf block
			var (
				current = origin.Number.Uint64()
				next    = current + query.Skip + 1
			)
			if next <= current {
				infos, _ := json.MarshalIndent(peer.Peer.Info(), "", "  ")
				peer.Log().Warn("GetBlockHeaders skip overflow attack", "current", current, "skip", query.Skip, "next", next, "attacker", infos)
				unknown = true
			} else {
				if header := backend.Chain().GetHeaderByNumber(next); header != nil {
					nextHash := header.Hash()
					expOldHash, _ := backend.Chain().GetAncestor(nextHash, next, query.Skip+1, &maxNonCanonical)
					if expOldHash == query.Origin.Hash {
						query.Origin.Hash, query.Origin.Number = nextHash, next
					} else {
						unknown = true
					}
				} else {
					unknown = true
				}
			}
		case query.Reverse:
			// Number based traversal towards the genesis block
			if query.Origin.Number >= query.Skip+1 {
				query.Origin.Number -= query.Skip + 1
			} else {
				unknown = true
			}

		case !query.Reverse:
			// Number based traversal towards the leaf block
			query.Origin.Number += query.Skip + 1
		}
	}
	return headers
}

// handleGetBlockHeaders66 is the eth/66 version of handleGetBlockHeaders
func handleGetBlockHeaders66(backend Backend, msg Decoder, peer *Peer) error {
	// Decode the complex header query
	var query GetBlockHeadersPacket66
	if err := msg.Decode(&query); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		response := ServiceGetBlockHeadersQuery(backend.Chain(), query.GetBlockHeadersPacket, peer)
		if len(response) == int(query.GetBlockHeadersPacket.Amount) {
			return peer.ReplyBlockHeadersRLP(query.RequestId, response)
		} else {
			// Wemix: fall back to old behavior
			response2 := answerGetBlockHeadersQuery(backend, query.GetBlockHeadersPacket, peer)
			if len(response2) > len(response) {
				return peer.ReplyBlockHeaders(query.RequestId, response2)
			} else {
				return peer.ReplyBlockHeadersRLP(query.RequestId, response)
			}
		}
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

// ServiceGetBlockHeadersQuery assembles the response to a header query. It is
// exposed to allow external packages to test protocol behavior.
func ServiceGetBlockHeadersQuery(chain *core.BlockChain, query *GetBlockHeadersPacket, peer *Peer) []rlp.RawValue {
	if query.Skip == 0 {
		// The fast path: when the request is for a contiguous segment of headers.
		return serviceContiguousBlockHeaderQuery(chain, query)
	} else {
		return serviceNonContiguousBlockHeaderQuery(chain, query, peer)
	}
}

func serviceNonContiguousBlockHeaderQuery(chain *core.BlockChain, query *GetBlockHeadersPacket, peer *Peer) []rlp.RawValue {
	hashMode := query.Origin.Hash != (common.Hash{})
	first := true
	maxNonCanonical := uint64(100)

	// Gather headers until the fetch or network limits is reached
	var (
		bytes   common.StorageSize
		headers []rlp.RawValue
		unknown bool
		lookups int
	)
	for !unknown && len(headers) < int(query.Amount) && bytes < softResponseLimit &&
		len(headers) < maxHeadersServe && lookups < 2*maxHeadersServe {
		lookups++
		// Retrieve the next header satisfying the query
		var origin *types.Header
		if hashMode {
			if first {
				first = false
				origin = chain.GetHeaderByHash(query.Origin.Hash)
				if origin != nil {
					query.Origin.Number = origin.Number.Uint64()
				}
			} else {
				origin = chain.GetHeader(query.Origin.Hash, query.Origin.Number)
			}
		} else {
			origin = chain.GetHeaderByNumber(query.Origin.Number)
		}
		if origin == nil {
			break
		}
		if rlpData, err := rlp.EncodeToBytes(origin); err != nil {
			log.Crit("Unable to decode our own headers", "err", err)
		} else {
			headers = append(headers, rlp.RawValue(rlpData))
			bytes += common.StorageSize(len(rlpData))
		}
		// Advance to the next header of the query
		switch {
		case hashMode && query.Reverse:
			// Hash based traversal towards the genesis block
			ancestor := query.Skip + 1
			if ancestor == 0 {
				unknown = true
			} else {
				query.Origin.Hash, query.Origin.Number = chain.GetAncestor(query.Origin.Hash, query.Origin.Number, ancestor, &maxNonCanonical)
				unknown = (query.Origin.Hash == common.Hash{})
			}
		case hashMode && !query.Reverse:
			// Hash based traversal towards the leaf block
			var (
				current = origin.Number.Uint64()
				next    = current + query.Skip + 1
			)
			if next <= current {
				infos, _ := json.MarshalIndent(peer.Peer.Info(), "", "  ")
				peer.Log().Warn("GetBlockHeaders skip overflow attack", "current", current, "skip", query.Skip, "next", next, "attacker", infos)
				unknown = true
			} else {
				if header := chain.GetHeaderByNumber(next); header != nil {
					nextHash := header.Hash()
					expOldHash, _ := chain.GetAncestor(nextHash, next, query.Skip+1, &maxNonCanonical)
					if expOldHash == query.Origin.Hash {
						query.Origin.Hash, query.Origin.Number = nextHash, next
					} else {
						unknown = true
					}
				} else {
					unknown = true
				}
			}
		case query.Reverse:
			// Number based traversal towards the genesis block
			if query.Origin.Number >= query.Skip+1 {
				query.Origin.Number -= query.Skip + 1
			} else {
				unknown = true
			}

		case !query.Reverse:
			// Number based traversal towards the leaf block
			query.Origin.Number += query.Skip + 1
		}
	}
	return headers
}

func serviceContiguousBlockHeaderQuery(chain *core.BlockChain, query *GetBlockHeadersPacket) []rlp.RawValue {
	count := query.Amount
	if count > maxHeadersServe {
		count = maxHeadersServe
	}
	if query.Origin.Hash == (common.Hash{}) {
		// Number mode, just return the canon chain segment. The backend
		// delivers in [N, N-1, N-2..] descending order, so we need to
		// accommodate for that.
		from := query.Origin.Number
		if !query.Reverse {
			from = from + count - 1
		}
		headers := chain.GetHeadersFrom(from, count)
		if !query.Reverse {
			for i, j := 0, len(headers)-1; i < j; i, j = i+1, j-1 {
				headers[i], headers[j] = headers[j], headers[i]
			}
		}
		return headers
	}
	// Hash mode.
	var (
		headers []rlp.RawValue
		hash    = query.Origin.Hash
		header  = chain.GetHeaderByHash(hash)
	)
	if header != nil {
		rlpData, _ := rlp.EncodeToBytes(header)
		headers = append(headers, rlpData)
	} else {
		// We don't even have the origin header
		return headers
	}
	num := header.Number.Uint64()
	if !query.Reverse {
		// Theoretically, we are tasked to deliver header by hash H, and onwards.
		// However, if H is not canon, we will be unable to deliver any descendants of
		// H.
		if canonHash := chain.GetCanonicalHash(num); canonHash != hash {
			// Not canon, we can't deliver descendants
			return headers
		}
		descendants := chain.GetHeadersFrom(num+count-1, count-1)
		for i, j := 0, len(descendants)-1; i < j; i, j = i+1, j-1 {
			descendants[i], descendants[j] = descendants[j], descendants[i]
		}
		headers = append(headers, descendants...)
		return headers
	}
	{ // Last mode: deliver ancestors of H
		for i := uint64(1); header != nil && i < count; i++ {
			header = chain.GetHeaderByHash(header.ParentHash)
			if header == nil {
				break
			}
			rlpData, _ := rlp.EncodeToBytes(header)
			headers = append(headers, rlpData)
		}
		return headers
	}
}

func handleGetBlockBodies(backend Backend, msg Decoder, peer *Peer) error {
	// Decode the block body retrieval message
	var query GetBlockBodiesPacket
	if err := msg.Decode(&query); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		response := ServiceGetBlockBodiesQuery(backend.Chain(), query)
		return peer.SendBlockBodiesRLP(response)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleGetBlockBodies66(backend Backend, msg Decoder, peer *Peer) error {
	// Decode the block body retrieval message
	var query GetBlockBodiesPacket66
	if err := msg.Decode(&query); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		response := ServiceGetBlockBodiesQuery(backend.Chain(), query.GetBlockBodiesPacket)
		return peer.ReplyBlockBodiesRLP(query.RequestId, response)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

// ServiceGetBlockBodiesQuery assembles the response to a body query. It is
// exposed to allow external packages to test protocol behavior.
func ServiceGetBlockBodiesQuery(chain *core.BlockChain, query GetBlockBodiesPacket) []rlp.RawValue {
	// Gather blocks until the fetch or network limits is reached
	var (
		bytes  int
		bodies []rlp.RawValue
	)
	for lookups, hash := range query {
		if bytes >= softResponseLimit || len(bodies) >= maxBodiesServe ||
			lookups >= 2*maxBodiesServe {
			break
		}
		if data := chain.GetBodyRLP(hash); len(data) != 0 {
			bodies = append(bodies, data)
			bytes += len(data)
		}
	}
	return bodies
}

func handleGetNodeData(backend Backend, msg Decoder, peer *Peer) error {
	// Decode the trie node data retrieval message
	var query GetNodeDataPacket
	if err := msg.Decode(&query); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		response := ServiceGetNodeDataQuery(backend.Chain(), query)
		return peer.SendNodeData(response)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleGetNodeData66(backend Backend, msg Decoder, peer *Peer) error {
	// Decode the trie node data retrieval message
	var query GetNodeDataPacket66
	if err := msg.Decode(&query); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		response := ServiceGetNodeDataQuery(backend.Chain(), query.GetNodeDataPacket)
		return peer.ReplyNodeData(query.RequestId, response)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

// ServiceGetNodeDataQuery assembles the response to a node data query. It is
// exposed to allow external packages to test protocol behavior.
func ServiceGetNodeDataQuery(chain *core.BlockChain, query GetNodeDataPacket) [][]byte {
	// Gather state data until the fetch or network limits is reached
	var (
		bytes int
		nodes [][]byte
	)
	for lookups, hash := range query {
		if bytes >= softResponseLimit || len(nodes) >= maxNodeDataServe ||
			lookups >= 2*maxNodeDataServe {
			break
		}
		// Retrieve the requested state entry
		entry, err := chain.TrieNode(hash)
		if len(entry) == 0 || err != nil {
			// Read the contract code with prefix only to save unnecessary lookups.
			entry, err = chain.ContractCodeWithPrefix(hash)
		}
		if err == nil && len(entry) > 0 {
			nodes = append(nodes, entry)
			bytes += len(entry)
		}
	}
	return nodes
}

func handleGetReceipts(backend Backend, msg Decoder, peer *Peer) error {
	// Decode the block receipts retrieval message
	var query GetReceiptsPacket
	if err := msg.Decode(&query); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		response := ServiceGetReceiptsQuery(backend.Chain(), query)
		return peer.SendReceiptsRLP(response)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleGetReceipts66(backend Backend, msg Decoder, peer *Peer) error {
	// Decode the block receipts retrieval message
	var query GetReceiptsPacket66
	if err := msg.Decode(&query); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		response := ServiceGetReceiptsQuery(backend.Chain(), query.GetReceiptsPacket)
		return peer.ReplyReceiptsRLP(query.RequestId, response)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

// ServiceGetReceiptsQuery assembles the response to a receipt query. It is
// exposed to allow external packages to test protocol behavior.
func ServiceGetReceiptsQuery(chain *core.BlockChain, query GetReceiptsPacket) []rlp.RawValue {
	// Gather state data until the fetch or network limits is reached
	var (
		bytes    int
		receipts []rlp.RawValue
	)
	for lookups, hash := range query {
		if bytes >= softResponseLimit || len(receipts) >= maxReceiptsServe ||
			lookups >= 2*maxReceiptsServe {
			break
		}
		// Retrieve the requested block's receipts
		results := chain.GetReceiptsByHash(hash)
		if results == nil {
			if header := chain.GetHeaderByHash(hash); header == nil || header.ReceiptHash != types.EmptyRootHash {
				continue
			}
		}
		// If known, encode and queue for response packet
		if encoded, err := rlp.EncodeToBytes(results); err != nil {
			log.Error("Failed to encode receipt", "err", err)
		} else {
			receipts = append(receipts, encoded)
			bytes += len(encoded)
		}
	}
	return receipts
}

func handleNewBlockhashes(backend Backend, msg Decoder, peer *Peer) error {
	// A batch of new block announcements just arrived
	ann := new(NewBlockHashesPacket)
	if err := msg.Decode(ann); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		// Mark the hashes as present at the remote node
		for _, block := range *ann {
			peer.markBlock(block.Hash)
		}
		// Deliver them all to the backend for queuing
		return backend.Handle(peer, ann)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleNewBlock(backend Backend, msg Decoder, peer *Peer) error {
	// Retrieve and decode the propagated block
	ann := new(NewBlockPacket)
	if err := msg.Decode(ann); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	if err := ann.sanityCheck(); err != nil {
		return err
	}
	f := func() error {
		if hash := types.CalcUncleHash(ann.Block.Uncles()); hash != ann.Block.UncleHash() {
			log.Warn("Propagated block has invalid uncles", "have", hash, "exp", ann.Block.UncleHash())
			return nil // TODO(karalabe): return error eventually, but wait a few releases
		}
		if hash := types.DeriveSha(ann.Block.Transactions(), trie.NewStackTrie(nil)); hash != ann.Block.TxHash() {
			log.Warn("Propagated block has invalid body", "have", hash, "exp", ann.Block.TxHash())
			return nil // TODO(karalabe): return error eventually, but wait a few releases
		}
		ann.Block.ReceivedAt = msg.Time()
		ann.Block.ReceivedFrom = peer

		// Mark the peer as owning the block
		peer.markBlock(ann.Block.Hash())

		return backend.Handle(peer, ann)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleBlockHeaders(backend Backend, msg Decoder, peer *Peer) error {
	// A batch of headers arrived to one of our previous requests
	res := new(BlockHeadersPacket)
	if err := msg.Decode(res); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		metadata := func() interface{} {
			hashes := make([]common.Hash, len(*res))
			for i, header := range *res {
				hashes[i] = header.Hash()
			}
			return hashes
		}
		return peer.dispatchResponse(&Response{
			id:   peer.genRequestId(BlockHeadersMsg),
			code: BlockHeadersMsg,
			Res:  res,
		}, metadata)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleBlockHeaders66(backend Backend, msg Decoder, peer *Peer) error {
	// A batch of headers arrived to one of our previous requests
	res := new(BlockHeadersPacket66)
	if err := msg.Decode(res); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		metadata := func() interface{} {
			hashes := make([]common.Hash, len(res.BlockHeadersPacket))
			for i, header := range res.BlockHeadersPacket {
				hashes[i] = header.Hash()
			}
			return hashes
		}
		return peer.dispatchResponse(&Response{
			id:   res.RequestId,
			code: BlockHeadersMsg,
			Res:  &res.BlockHeadersPacket,
		}, metadata)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleBlockBodies(backend Backend, msg Decoder, peer *Peer) error {
	// A batch of block bodies arrived to one of our previous requests
	res := new(BlockBodiesPacket)
	if err := msg.Decode(res); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		metadata := func() interface{} {
			var (
				txsHashes   = make([]common.Hash, len(*res))
				uncleHashes = make([]common.Hash, len(*res))
			)
			hasher := trie.NewStackTrie(nil)
			for i, body := range *res {
				txsHashes[i] = types.DeriveSha(types.Transactions(body.Transactions), hasher)
				uncleHashes[i] = types.CalcUncleHash(body.Uncles)
			}
			return [][]common.Hash{txsHashes, uncleHashes}
		}
		return peer.dispatchResponse(&Response{
			id:   peer.genRequestId(BlockBodiesMsg),
			code: BlockBodiesMsg,
			Res:  res,
		}, metadata)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleBlockBodies66(backend Backend, msg Decoder, peer *Peer) error {
	// A batch of block bodies arrived to one of our previous requests
	res := new(BlockBodiesPacket66)
	if err := msg.Decode(res); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		metadata := func() interface{} {
			var (
				txsHashes   = make([]common.Hash, len(res.BlockBodiesPacket))
				uncleHashes = make([]common.Hash, len(res.BlockBodiesPacket))
			)
			hasher := trie.NewStackTrie(nil)
			for i, body := range res.BlockBodiesPacket {
				txsHashes[i] = types.DeriveSha(types.Transactions(body.Transactions), hasher)
				uncleHashes[i] = types.CalcUncleHash(body.Uncles)
			}
			return [][]common.Hash{txsHashes, uncleHashes}
		}
		return peer.dispatchResponse(&Response{
			id:   res.RequestId,
			code: BlockBodiesMsg,
			Res:  &res.BlockBodiesPacket,
		}, metadata)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleNodeData(backend Backend, msg Decoder, peer *Peer) error {
	// A batch of node state data arrived to one of our previous requests
	res := new(NodeDataPacket)
	if err := msg.Decode(res); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		return peer.dispatchResponse(&Response{
			id:   peer.genRequestId(NodeDataMsg),
			code: NodeDataMsg,
			Res:  res,
		}, nil) // No post-processing, we're not using this packet anymore
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleNodeData66(backend Backend, msg Decoder, peer *Peer) error {
	// A batch of node state data arrived to one of our previous requests
	res := new(NodeDataPacket66)
	if err := msg.Decode(res); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		return peer.dispatchResponse(&Response{
			id:   res.RequestId,
			code: NodeDataMsg,
			Res:  &res.NodeDataPacket,
		}, nil) // No post-processing, we're not using this packet anymore
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleReceipts(backend Backend, msg Decoder, peer *Peer) error {
	// A batch of receipts arrived to one of our previous requests
	res := new(ReceiptsPacket)
	if err := msg.Decode(res); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		metadata := func() interface{} {
			hasher := trie.NewStackTrie(nil)
			hashes := make([]common.Hash, len(*res))
			for i, receipt := range *res {
				hashes[i] = types.DeriveSha(types.Receipts(receipt), hasher)
			}
			return hashes
		}
		return peer.dispatchResponse(&Response{
			id:   peer.genRequestId(ReceiptsMsg),
			code: ReceiptsMsg,
			Res:  res,
		}, metadata)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleReceipts66(backend Backend, msg Decoder, peer *Peer) error {
	// A batch of receipts arrived to one of our previous requests
	res := new(ReceiptsPacket66)
	if err := msg.Decode(res); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		metadata := func() interface{} {
			hasher := trie.NewStackTrie(nil)
			hashes := make([]common.Hash, len(res.ReceiptsPacket))
			for i, receipt := range res.ReceiptsPacket {
				hashes[i] = types.DeriveSha(types.Receipts(receipt), hasher)
			}
			return hashes
		}
		return peer.dispatchResponse(&Response{
			id:   res.RequestId,
			code: ReceiptsMsg,
			Res:  &res.ReceiptsPacket,
		}, metadata)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleNewPooledTransactionHashes(backend Backend, msg Decoder, peer *Peer) error {
	// New transaction announcement arrived, make sure we have
	// a valid and fresh chain to handle them
	if !backend.AcceptTxs() {
		return nil
	}
	ann := new(NewPooledTransactionHashesPacket)
	if err := msg.Decode(ann); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		// Schedule all the unknown hashes for retrieval
		for _, hash := range *ann {
			peer.markTransaction(hash)
		}
		return backend.Handle(peer, ann)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleGetPooledTransactions(backend Backend, msg Decoder, peer *Peer) error {
	// Decode the pooled transactions retrieval message
	var query GetPooledTransactionsPacket
	if err := msg.Decode(&query); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		hashes, txs := answerGetPooledTransactions(backend, query, peer)
		return peer.SendPooledTransactionsRLP(hashes, txs)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleGetPooledTransactions66(backend Backend, msg Decoder, peer *Peer) error {
	// Decode the pooled transactions retrieval message
	var query GetPooledTransactionsPacket66
	if err := msg.Decode(&query); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		hashes, txs := answerGetPooledTransactions(backend, query.GetPooledTransactionsPacket, peer)
		return peer.ReplyPooledTransactionsRLP(query.RequestId, hashes, txs)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func answerGetPooledTransactions(backend Backend, query GetPooledTransactionsPacket, peer *Peer) ([]common.Hash, []rlp.RawValue) {
	// Gather transactions until the fetch or network limits is reached
	var (
		bytes  int
		hashes []common.Hash
		txs    []rlp.RawValue
	)
	for _, hash := range query {
		if bytes >= softResponseLimit {
			break
		}
		// Retrieve the requested transaction, skipping if unknown to us
		tx := backend.TxPool().Get(hash)
		if tx == nil {
			continue
		}
		// If known, encode and queue for response packet
		if encoded, err := rlp.EncodeToBytes(tx); err != nil {
			log.Error("Failed to encode transaction", "err", err)
		} else {
			hashes = append(hashes, hash)
			txs = append(txs, encoded)
			bytes += len(encoded)
		}
	}
	return hashes, txs
}

func handleTransactions(backend Backend, msg Decoder, peer *Peer) error {
	// Transactions arrived, make sure we have a valid and fresh chain to handle them
	if !backend.AcceptTxs() {
		return nil
	}
	// Transactions can be processed, parse all of them and deliver to the pool
	var txs TransactionsPacket
	if err := msg.Decode(&txs); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		for i, tx := range txs {
			// Validate and mark the remote transaction
			if tx == nil {
				return fmt.Errorf("%w: transaction %d is nil", errDecode, i)
			}
			peer.markTransaction(tx.Hash())
		}
		return backend.Handle(peer, &txs)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handleTransactionsEx(backend Backend, msg Decoder, peer *Peer) error {
	// Transactions arrived, make sure we have a valid and fresh chain to handle them
	if !backend.AcceptTxs() {
		return nil
	}
	// Transactions can be processed, parse all of them and deliver to the pool
	var txexs TransactionsExPacket
	if err := msg.Decode(&txexs); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		signer := types.MakeSigner(backend.Chain().Config(), backend.Chain().CurrentBlock().Number())
		txs := types.TxExs2Txs(signer, txexs, wemixminer.IsPartner(peer.ID()))
		for i, tx := range txs {
			// Validate and mark the remote transaction
			if tx == nil {
				return fmt.Errorf("%w: transaction %d is nil", errDecode, i)
			}
			peer.markTransaction(tx.Hash())
		}
		txsp := TransactionsPacket(txs)
		return backend.Handle(peer, &txsp)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handlePooledTransactions(backend Backend, msg Decoder, peer *Peer) error {
	// Transactions arrived, make sure we have a valid and fresh chain to handle them
	if !backend.AcceptTxs() {
		return nil
	}
	// Transactions can be processed, parse all of them and deliver to the pool
	var txs PooledTransactionsPacket
	if err := msg.Decode(&txs); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		for i, tx := range txs {
			// Validate and mark the remote transaction
			if tx == nil {
				return fmt.Errorf("%w: transaction %d is nil", errDecode, i)
			}
			peer.markTransaction(tx.Hash())
		}
		return backend.Handle(peer, &txs)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}

func handlePooledTransactions66(backend Backend, msg Decoder, peer *Peer) error {
	// Transactions arrived, make sure we have a valid and fresh chain to handle them
	if !backend.AcceptTxs() {
		return nil
	}
	// Transactions can be processed, parse all of them and deliver to the pool
	var txs PooledTransactionsPacket66
	if err := msg.Decode(&txs); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}
	f := func() error {
		for i, tx := range txs.PooledTransactionsPacket {
			// Validate and mark the remote transaction
			if tx == nil {
				return fmt.Errorf("%w: transaction %d is nil", errDecode, i)
			}
			peer.markTransaction(tx.Hash())
		}
		requestTracker.Fulfil(peer.id, peer.version, PooledTransactionsMsg, txs.RequestId)

		return backend.Handle(peer, &txs.PooledTransactionsPacket)
	}
	if params.ConsensusMethod == params.ConsensusPoW {
		return f()
	}
	go f()
	return nil
}
