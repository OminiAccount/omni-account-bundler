package processor

import (
	"github.com/OAAC/jsonrpc"
	"github.com/OAAC/pool"
)

func createPool(cfg pool.Config) *pool.Pool {
	return pool.NewMemoryPool(cfg)
}

func createJSONRPCServer(cfg jsonrpc.Config, pool *pool.Pool) *jsonrpc.Server {
	s := jsonrpc.Service{
		Name:    jsonrpc.APIEth,
		Service: jsonrpc.NewEthEndpoints(pool),
	}
	return jsonrpc.NewServer(cfg, s)
}
