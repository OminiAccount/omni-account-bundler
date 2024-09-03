package state

import (
	"context"
	"fmt"
	"github.com/OAAC/config"
	"github.com/OAAC/ethereum"
	"github.com/OAAC/jsonrpc/rpcs"
	"github.com/OAAC/utils/chains"
	"github.com/ethereum/go-ethereum/ethdb"
)

type Config struct {
	Context   context.Context
	LevelDB   ethdb.Database
	Networks  map[chains.ChainId]ethereum.Network
	zkPools   []rpcs.Rpc
	heartBeat rpcs.Rpc
}

func NewConfig(ctx context.Context, db ethdb.Database, cfg config.Config) (Config, error) {
	c := Config{
		Context:   ctx,
		LevelDB:   db,
		zkPools:   cfg.JsonRpc.ZkPools,
		heartBeat: cfg.JsonRpc.HeartBeat,
		Networks:  make(map[chains.ChainId]ethereum.Network, chains.MaxChainInfoLength),
	}

	for _, network := range cfg.Ethereum.Networks {
		chainId := chains.ChainId(network.ChainId)
		if !chainId.IsSupport() {
			return c, fmt.Errorf("ChainId %d is not supported", network.ChainId)
		}
		c.Networks[chainId] = network
	}

	return c, nil
}
