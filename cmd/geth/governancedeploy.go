// governancedeploy.js

package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	gov "github.com/ethereum/go-ethereum/wemix/bind"
	"github.com/ethereum/go-ethereum/wemix/metclient"
	"gopkg.in/urfave/cli.v1"
)

func getInitialGovernanceMembersAndNodes(configJsFile string) (members gov.InitMembers, rewardPoolAccount, maintenanceAccount *common.Address, err error) {
	var fin *os.File
	if fin, err = os.Open(configJsFile); err != nil {
		return
	}
	defer fin.Close()

	var cfg *genesisConfig
	cfg, err = loadGenesisConfig(fin)
	if err != nil {
		return
	}

	zeroAddress := common.Address{}

	if cfg.RewardPool != zeroAddress {
		rewardPoolAccount = &cfg.RewardPool
	}
	if cfg.Maintenance != zeroAddress {
		maintenanceAccount = &cfg.Maintenance
	}
	members = make(gov.InitMembers, len(cfg.Members))
	for i, member := range cfg.Members {
		var (
			staker, voter, reward common.Address = member.Staker, member.Voter, member.Reward
		)
		if addr := member.Addr; addr != zeroAddress {
			if staker == zeroAddress {
				staker = addr
			}
			if voter == zeroAddress {
				voter = addr
			}
			if reward == zeroAddress {
				reward = addr
			}
		}
		members[i] = gov.InitMember{
			Staker:  staker,
			Voter:   voter,
			Reward:  reward,
			Name:    member.Name,
			Enode:   member.Id,
			Ip:      member.Ip,
			Port:    member.Port,
			Deposit: member.Stake,
		}
	}

	return
}

// governance-contract.js config.js
func deployGovernanceContracts(ctx *cli.Context) error {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stdout, log.TerminalFormat(true))))
	// get command line arguments
	args := ctx.Args()

	var (
		errInvalidArguments     error = errors.New("invalid Arguments")
		configFile, accountFile string
		lockAmount              *big.Int
	)
	switch len(args) {
	case 3:
		configFile, accountFile = args[0], args[1]
		lockAmount, ok := new(big.Int).SetString(args[2], 10)
		if !ok || lockAmount.Sign() <= 0 {
			return errInvalidArguments
		}
	case 2:
		configFile, accountFile = args[0], args[1]
		lockAmount = gov.DefaultInitEnvStorage.STAKING_MIN
	default:
		return errInvalidArguments
	}

	client, err := ethclient.Dial(ctx.String(urlFlag.Name))
	if err != nil {
		return err
	}
	chainID, err := client.ChainID(context.TODO())
	if err != nil {
		return err
	}

	var opts *bind.TransactOpts
	if from, err := metclient.LoadAccount(utils.GetPassPhraseWithList("", false, 0, utils.MakePasswordList(ctx)), accountFile); err != nil {
		return err
	} else if opts, err = bind.NewKeyedTransactorWithChainID(from.PrivateKey, chainID); err != nil {
		return err
	} else {
		if ctx.IsSet(gasFlag.Name) {
			opts.GasLimit = ctx.Uint64(gasFlag.Name)
		}
		if ctx.IsSet(gasPriceFlag.Name) {
			opts.GasPrice = big.NewInt(ctx.Int64(gasPriceFlag.Name))
		}
	}

	return deployGovernance(client, opts, lockAmount, configFile)
}

type Client interface {
	bind.ContractBackend
	bind.DeployBackend
}

func deployGovernance(cli Client, ownerOpts *bind.TransactOpts, lockAmount *big.Int, configJsFile string) error {
	envConfig := gov.DefaultInitEnvStorage
	if lockAmount.Cmp(envConfig.STAKING_MIN) < 0 || lockAmount.Cmp(envConfig.STAKING_MAX) > 0 {
		return fmt.Errorf("invalid lock amount, input:%v, min:%v, max:%v", lockAmount, envConfig.STAKING_MIN, envConfig.STAKING_MAX)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ownerOpts.Context = ctx

	// initial members and nodes data
	members, rewardPoolAccount, maintenanceAccount, err := getInitialGovernanceMembersAndNodes(configJsFile)
	if err != nil {
		return err
	}

	// deploy & set domains
	contracts, impAddress, err := gov.DeployGovContracts(ownerOpts, cli, func() map[string]common.Address {
		if rewardPoolAccount == nil && maintenanceAccount == nil {
			return nil
		}
		optionDomains := make(map[string]common.Address)
		if rewardPoolAccount != nil {
			optionDomains["RewardPool"] = *rewardPoolAccount
		}
		if maintenanceAccount != nil {
			optionDomains["Maintenance"] = *maintenanceAccount
		}
		return optionDomains
	}())
	if err != nil {
		return err
	}
	// init contracts
	err = gov.ExecuteInitialize(contracts, ownerOpts, cli, lockAmount, envConfig, members)
	if err != nil {
		return err
	}

	// print the addresses
	address := contracts.Address()
	fmt.Printf(`{
  "REGISTRY_ADDRESS": "%s",
  "STAKING_ADDRESS": "%s",
  "ENV_STORAGE_ADDRESS": "%s",
  "BALLOT_STORAGE_ADDRESS": "%s",
  "GOV_ADDRESS": "%s",
  "GOV_IMP_ADDRESS": "%s",
}
`,
		address.Registry, address.Staking, address.EnvStorage, address.BallotStorage, address.Gov, impAddress["GovImp"])

	return nil
}

// EOF
