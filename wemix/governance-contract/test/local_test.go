package test

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	gov "github.com/ethereum/go-ethereum/wemix/bind"
	"github.com/stretchr/testify/require"
)

var (
	ctx      = context.Background()
	callOpts = new(bind.CallOpts)
)

func TestTemp(t *testing.T) {
	client, err := ethclient.DialContext(ctx, "/Users/wm-bd000036/workspace/.gwemix/node1/gwemix.ipc")
	require.NoError(t, err)
	contracts := getContracts(t, client)
	length, err := contracts.GovImp.GetNodeLength(callOpts)
	require.NoError(t, err)
	for i := int64(1); i <= length.Int64(); i++ {
		member, err := contracts.GovImp.GetMember(callOpts, big.NewInt(i))
		require.NoError(t, err)
		t.Log("member", i, member.Hex())
		node, err := contracts.GovImp.GetNode(callOpts, big.NewInt(i))
		require.NoError(t, err)
		t.Log("node", i, "\n\t", "Name", string(node.Name), "\n\t", "Enode", hex.EncodeToString(node.Enode), "\n\t", "Ip", string(node.Ip), "\n\t", "Port", node.Port)
	}
}

func getContracts(t *testing.T, client *ethclient.Client) *gov.GovContracts {
	block, err := client.BlockByNumber(ctx, common.Big0)
	require.NoError(t, err)
	contracts, err := gov.GetGovContractsByOwner(callOpts, client, block.Coinbase())
	require.NoError(t, err)
	return contracts
}
