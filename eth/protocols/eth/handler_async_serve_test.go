// Coverage that every retrieval request type still serves a response on the
// asynchronous (PoA) path, i.e. each one routes through peer.dispatchHandler
// rather than bypassing the per-peer/global bound. The concurrency bound itself
// is handler-agnostic and covered by handler_concurrency_test.go; this test
// pins that headers, bodies, node data, receipts and pooled transactions all
// still reply when dispatched asynchronously.
package eth

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
)

func TestAsyncHandlersServeAllRetrievalTypes(t *testing.T) {
	// No t.Parallel(): this flips the package global params.ConsensusMethod.
	backend := newTestBackend(maxBodiesServe + 15)
	defer backend.close()

	peer, _ := newTestPeer("peer", ETH66, backend)
	defer peer.close()

	// Switch to the async (PoA) serving path only after the backend and peer are
	// built; generating the test chain under PoA would need governance hooks this
	// package does not wire up. The request send below happens-after this write
	// and the handler reads the mode after receiving it, so there is no race.
	prev := params.ConsensusMethod
	params.ConsensusMethod = params.ConsensusPoA
	defer func() { params.ConsensusMethod = prev }()

	genesis := backend.chain.GetBlockByNumber(0).Hash()
	root := backend.chain.CurrentBlock().Root()

	cases := []struct {
		name     string
		reqCode  uint64
		req      interface{}
		respCode uint64
	}{
		{
			"headers", GetBlockHeadersMsg,
			&GetBlockHeadersPacket66{RequestId: 1, GetBlockHeadersPacket: &GetBlockHeadersPacket{Origin: HashOrNumber{Number: 0}, Amount: 1}},
			BlockHeadersMsg,
		},
		{
			"bodies", GetBlockBodiesMsg,
			&GetBlockBodiesPacket66{RequestId: 2, GetBlockBodiesPacket: GetBlockBodiesPacket{genesis}},
			BlockBodiesMsg,
		},
		{
			"nodedata", GetNodeDataMsg,
			&GetNodeDataPacket66{RequestId: 3, GetNodeDataPacket: GetNodeDataPacket{root}},
			NodeDataMsg,
		},
		{
			"receipts", GetReceiptsMsg,
			&GetReceiptsPacket66{RequestId: 4, GetReceiptsPacket: GetReceiptsPacket{genesis}},
			ReceiptsMsg,
		},
		{
			"pooledtx", GetPooledTransactionsMsg,
			&GetPooledTransactionsPacket66{RequestId: 5, GetPooledTransactionsPacket: GetPooledTransactionsPacket{common.Hash{}}},
			PooledTransactionsMsg,
		},
	}
	for _, c := range cases {
		if err := p2p.Send(peer.app, c.reqCode, c.req); err != nil {
			t.Fatalf("%s: send request: %v", c.name, err)
		}
		msg, err := peer.app.ReadMsg()
		if err != nil {
			t.Fatalf("%s: read response: %v", c.name, err)
		}
		code := msg.Code
		msg.Discard()
		if code != c.respCode {
			t.Fatalf("%s: response code = %d, want %d", c.name, code, c.respCode)
		}
	}
}
