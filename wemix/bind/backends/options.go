package backends

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	gov "github.com/ethereum/go-ethereum/wemix/bind"
)

type OptionFn func(nodeConf *node.Config, ethConf *ethconfig.Config, envConfig *gov.InitEnvStorage)

func SetLogLevel(level int) OptionFn {
	return func(nodeConf *node.Config, ethConf *ethconfig.Config, envConfig *gov.InitEnvStorage) {
		log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(level), log.StreamHandler(os.Stdout, log.TerminalFormat(true))))
	}
}

// WithBlockGasLimit configures the simulated backend to target a specific gas limit
// when producing blocks.
func WithBlockGasLimit(gaslimit uint64) OptionFn {
	return func(nodeConf *node.Config, ethConf *ethconfig.Config, envConfig *gov.InitEnvStorage) {
		ethConf.Genesis.GasLimit = gaslimit
		ethConf.Miner.GasCeil = gaslimit
	}
}

// WithCallGasLimit configures the simulated backend to cap eth_calls to a specific
// gas limit when running client operations.
func WithCallGasLimit(gaslimit uint64) OptionFn {
	return func(nodeConf *node.Config, ethConf *ethconfig.Config, envConfig *gov.InitEnvStorage) {
		ethConf.RPCGasCap = gaslimit
	}
}

// WithMinerMinTip configures the simulated backend to require a specific minimum
// gas tip for a transaction to be included.
//
// 0 is not possible as a live Geth node would reject that due to DoS protection,
// so the simulated backend will replicate that behavior for consistency.
func WithMinerMinTip(tip *big.Int) OptionFn {
	if tip == nil || tip.Sign() <= 0 {
		panic("invalid miner minimum tip")
	}
	return func(nodeConf *node.Config, ethConf *ethconfig.Config, envConfig *gov.InitEnvStorage) {
		ethConf.Miner.GasPrice = tip
	}
}
