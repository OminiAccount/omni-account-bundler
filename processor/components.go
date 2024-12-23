package processor

import (
	"github.com/OAB/jsonrpc"
	"github.com/OAB/pool"
	"github.com/OAB/state"
)

func createJSONRPCServer(cfg jsonrpc.RpcsConfig, p *pool.Pool, state *state.State) *jsonrpc.Server {
	s := jsonrpc.Service{
		Name:    jsonrpc.APIEth,
		Service: jsonrpc.NewEthEndpoints(p, state, state.GetHisIns()),
	}
	return jsonrpc.NewServer(cfg, s)
}
