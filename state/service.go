package state

import (
	"context"
	"fmt"
	"github.com/OAAC/config"
	"github.com/OAAC/ethereum"
	"github.com/OAAC/jsonrpc"
	"github.com/OAAC/synchronizer"
	"github.com/OAAC/utils/chains"
	"github.com/ethereum/go-ethereum/ethdb"
)

type ServiceConfig struct {
	Context   context.Context
	LevelDB   ethdb.Database
	Networks  map[chains.ChainId]ethereum.Network
	zkPools   []jsonrpc.Rpc
	heartBeat jsonrpc.Rpc
}

func NewServiceConfig(ctx context.Context, db ethdb.Database, cfg config.Config) (ServiceConfig, error) {
	serviceConfig := ServiceConfig{
		Context:   ctx,
		LevelDB:   db,
		zkPools:   cfg.JsonRpc.ZkPools,
		heartBeat: cfg.JsonRpc.HeartBeat,
		Networks:  make(map[chains.ChainId]ethereum.Network, chains.MaxChainInfoLength),
	}

	for _, network := range cfg.Ethereum.Networks {
		chainId := chains.ChainId(network.ChainId)
		if !chainId.IsSupport() {
			return serviceConfig, fmt.Errorf("ChainId %d is not supported", network.ChainId)
		}
		serviceConfig.Networks[chainId] = network
	}

	return serviceConfig, nil
}

type Service struct {
	cfg    ServiceConfig
	ctx    context.Context
	cancel func()

	synchronizer *synchronizer.Synchronizer

	// collect Tickets
	//collectTicketsChannel *CollectTicketsChannel
	// Submit a challenge to the zk-pool
	//submitChallengesChannel *SubmitChallengesChannel
	// Commit a challenge to the contract
	//commitChallengesChannel *CommitChallengesChannel

	//heartBeatChannel *HeartBeatChannel
}

var service *Service

func GetService() *Service {
	return service
}

func NewService(sc ServiceConfig) (*Service, error) {

	ctx, cancel := context.WithCancel(sc.Context)
	defer cancel()

	service = &Service{
		cfg:    sc,
		ctx:    ctx,
		cancel: cancel,
	}

	return service, nil
}
