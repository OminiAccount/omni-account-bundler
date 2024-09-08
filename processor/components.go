package processor

import (
	"github.com/OAB/jsonrpc"
	"github.com/OAB/jsonrpc/rpcs"
	"github.com/OAB/pool"
	"github.com/OAB/state"
)

func createJSONRPCServer(cfg rpcs.RpcsConfig, pool *pool.Pool, state *state.State) *jsonrpc.Server {
	s := jsonrpc.Service{
		Name:    jsonrpc.APIEth,
		Service: jsonrpc.NewEthEndpoints(pool, state),
	}
	return jsonrpc.NewServer(cfg, s)
}
