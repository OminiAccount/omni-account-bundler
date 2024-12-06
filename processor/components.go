package processor

import (
	"github.com/OAB/jsonrpc"
	"github.com/OAB/state"
)

func createJSONRPCServer(cfg jsonrpc.RpcsConfig, state *state.State) *jsonrpc.Server {
	s := jsonrpc.Service{
		Name:    jsonrpc.APIEth,
		Service: jsonrpc.NewEthEndpoints(state),
	}
	return jsonrpc.NewServer(cfg, s)
}
