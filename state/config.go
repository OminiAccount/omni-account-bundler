package state

import (
	"context"
	"github.com/OAB/config"
	"github.com/OAB/jsonrpc/rpcs"
	"github.com/ethereum/go-ethereum/ethdb"
)

type Config struct {
	Context   context.Context
	LevelDB   ethdb.Database
	zkPools   []rpcs.Rpc
	heartBeat rpcs.Rpc
}

func NewConfig(ctx context.Context, db ethdb.Database, cfg config.Config) (Config, error) {
	c := Config{
		Context:   ctx,
		LevelDB:   db,
		zkPools:   cfg.JsonRpc.ZkPools,
		heartBeat: cfg.JsonRpc.HeartBeat,
	}

	return c, nil
}
