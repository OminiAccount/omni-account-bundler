package processor

import (
	"github.com/OAAC/jsonrpc"
	"github.com/OAAC/jsonrpc/rpcs"
	"github.com/OAAC/pool"
	"github.com/OAAC/state"
)

func createJSONRPCServer(cfg rpcs.RpcsConfig, pool *pool.Pool, state *state.State) *jsonrpc.Server {
	s := jsonrpc.Service{
		Name:    jsonrpc.APIEth,
		Service: jsonrpc.NewEthEndpoints(pool, state),
	}
	return jsonrpc.NewServer(cfg, s)
}
