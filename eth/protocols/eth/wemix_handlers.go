package eth

import (
	"fmt"

	wemixapi "github.com/ethereum/go-ethereum/wemix/api"
	wemixminer "github.com/ethereum/go-ethereum/wemix/miner"
)

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
