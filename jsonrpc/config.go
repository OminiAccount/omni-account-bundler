package jsonrpc

type Rpc struct {
	name string
	url  string
}

// Config configures the JsonRpc server
type Config struct {
	Host      string
	Port      int
	ZkPools   []Rpc `toml:"zk-pools"`
	Moralis   []string
	HeartBeat Rpc `toml:"heart-beat ,omitempty"`
}
