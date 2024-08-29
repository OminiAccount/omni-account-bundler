package services

import (
	"context"
	"fmt"
	"github.com/OAAC/config"
	"github.com/ethereum/go-ethereum/ethdb"
)

type ServiceConfig struct {
	Context   context.Context
	LevelDB   ethdb.Database
	Networks  map[ChainId]config.NetworkConfig
	zkPools   []config.Rpc
	heartBeat config.Rpc
}

func NewServiceConfig(ctx context.Context, db ethdb.Database, cfg config.Config) (ServiceConfig, error) {
	serviceConfig := ServiceConfig{
		Context:   ctx,
		LevelDB:   db,
		zkPools:   cfg.API.ZkPools,
		heartBeat: cfg.API.HeartBeat,
		Networks:  make(map[ChainId]config.NetworkConfig, MaxChainInfoLength),
	}

	for _, network := range cfg.Networks {
		chainId := ChainId(network.ChainId)
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

	// collect Tickets
	collectTicketsChannel *CollectTicketsChannel
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
