package eth

import (
	"fmt"
	"sync"
	"time"

	wemixapi "github.com/ethereum/go-ethereum/wemix/api"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

// Per-peer minimum interval between accepted StatusEx messages. Bounds the
// goroutine/CPU cost from a single compromised partner flooding StatusEx and
// the (height, hash) churn in miningPeers. SyncIdleThreshold is 30s and the
// block interval is ~15s, so 5s is well above any legitimate cadence.
const statusExMinInterval = 5 * time.Second

// statusExLastSeen maps peer.ID() -> time.Time of the last accepted StatusEx.
// Bounded by the registered partner set (governance-controlled), so no
// eviction is needed; stale entries are harmless.
var statusExLastSeen sync.Map

func handleGetPendingTxs(backend Backend, msg Decoder, peer *Peer) error {
	// not supported, just ignore it.
	return nil
}

func handleGetStatusEx(backend Backend, msg Decoder, peer *Peer) error {
	if !wemixminer.AmPartner() || !wemixminer.IsPartner(peer.ID()) {
		return nil
	}

	go func() {
		statusEx := wemixapi.GetMinerStatus()
		if statusEx == nil {
			// ignore the error, most likely server is shutting down
			return
		}
		statusEx.LatestBlockTd = backend.Chain().GetTd(statusEx.LatestBlockHash,
			statusEx.LatestBlockHeight.Uint64())
		if err := peer.SendStatusEx(statusEx); err != nil {
			// ignore the error
		}
	}()

	return nil
}

func handleStatusEx(backend Backend, msg Decoder, peer *Peer) error {
	if !wemixminer.AmPartner() || !wemixminer.IsPartner(peer.ID()) {
		return nil
	}

	// Drop StatusEx arriving faster than statusExMinInterval from the same peer.
	// Performed before Decode so the spammer pays the I/O cost but we don't pay
	// the decode + goroutine-spawn cost.
	now := time.Now()
	if prev, ok := statusExLastSeen.Load(peer.ID()); ok {
		if t, ok2 := prev.(time.Time); ok2 && now.Sub(t) < statusExMinInterval {
			return nil
		}
	}
	statusExLastSeen.Store(peer.ID(), now)

	var status wemixapi.WemixMinerStatus
	if err := msg.Decode(&status); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}

	// Rebind NodeName from the attacker-controlled string in the RLP payload to
	// the verified peer.ID(). miningPeers (wemix/sync.go) thus ends up keyed by
	// peer.ID(), preventing a single peer from injecting multiple entries under
	// spoofed NodeNames.
	status.NodeName = peer.ID()

	go func() {
		// Guard against a crafted StatusEx that omits LatestBlockTd: without this
		// check Cmp() would panic on a nil *big.Int, turning a single message into
		// a handler-goroutine DoS.
		if _, td := peer.Head(); status.LatestBlockTd != nil && status.LatestBlockTd.Cmp(td) > 0 {
			peer.SetHead(status.LatestBlockHash, status.LatestBlockTd)
		}
		wemixapi.GotStatusEx(&status)
	}()

	return nil
}

func handleEtcdAddMember(backend Backend, msg Decoder, peer *Peer) error {
	if !wemixminer.AmPartner() || !wemixminer.IsPartner(peer.ID()) {
		return nil
	}

	go func() {
		cluster, _ := wemixapi.EtcdAddMember(peer.ID())
		if err := peer.SendEtcdCluster(cluster); err != nil {
			// ignore the error
		}
	}()

	return nil
}

func handleEtcdCluster(backend Backend, msg Decoder, peer *Peer) error {
	if !wemixminer.AmPartner() || !wemixminer.IsPartner(peer.ID()) {
		return nil
	}
	var cluster string
	if err := msg.Decode(&cluster); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}

	go wemixapi.GotEtcdCluster(cluster)

	return nil
}
