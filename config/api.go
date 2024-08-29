package config

type Rpc struct {
	name string
	url  string
}

// APIConfig configures the API server
type APIConfig struct {
	Host      string
	Port      int
	ZkPools   []Rpc `toml:"zk-pools"`
	Moralis   []string
	HeartBeat Rpc `toml:"heart-beat ,omitempty"`
}
