package eth

import (
	"fmt"

	metaapi "github.com/ethereum/go-ethereum/metadium/api"
	metaminer "github.com/ethereum/go-ethereum/metadium/miner"
)

func handleGetPendingTxs(backend Backend, msg Decoder, peer *Peer) error {
	// not supported, just ignore it.
	return nil
}

func handleGetStatusEx(backend Backend, msg Decoder, peer *Peer) error {
	if !metaminer.AmPartner() || !metaminer.IsPartner(peer.ID()) {
		return nil
	}

	go func() {
		statusEx := metaapi.GetMinerStatus()
		statusEx.LatestBlockTd = backend.Chain().GetTd(statusEx.LatestBlockHash,
			statusEx.LatestBlockHeight.Uint64())
		if err := peer.SendStatusEx(statusEx); err != nil {
			// ignore the error
		}
	}()

	return nil
}

func handleStatusEx(backend Backend, msg Decoder, peer *Peer) error {
	if !metaminer.AmPartner() || !metaminer.IsPartner(peer.ID()) {
		return nil
	}
	var status metaapi.MetadiumMinerStatus
	if err := msg.Decode(&status); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}

	go func() {
		if _, td := peer.Head(); status.LatestBlockTd.Cmp(td) > 0 {
			peer.SetHead(status.LatestBlockHash, status.LatestBlockTd)
		}
		metaapi.GotStatusEx(&status)
	}()

	return nil
}

func handleEtcdAddMember(backend Backend, msg Decoder, peer *Peer) error {
	if !metaminer.AmPartner() || !metaminer.IsPartner(peer.ID()) {
		return nil
	}

	go func() {
		cluster, _ := metaapi.EtcdAddMember(peer.ID())
		if err := peer.SendEtcdCluster(cluster); err != nil {
			// ignore the error
		}
	}()

	return nil
}

func handleEtcdCluster(backend Backend, msg Decoder, peer *Peer) error {
	if !metaminer.AmPartner() || !metaminer.IsPartner(peer.ID()) {
		return nil
	}
	var cluster string
	if err := msg.Decode(&cluster); err != nil {
		return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
	}

	go metaapi.GotEtcdCluster(cluster)

	return nil
}
