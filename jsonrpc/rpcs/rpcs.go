package rpcs

type Rpc struct {
	name string
	url  string
}

// RpcsConfig configures the JsonRpc server
type RpcsConfig struct {
	Host      string
	Port      int
	ZkPools   []Rpc `toml:"zk-pools"`
	Moralis   []string
	HeartBeat Rpc `toml:"heart-beat ,omitempty"`
}
