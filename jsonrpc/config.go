package jsonrpc

type Rpc struct {
	name string
	url  string
}

// RpcsConfig configures the JsonRpc server
type RpcsConfig struct {
	Host      string
	Port      int
}
