package config

import (
	"github.com/BurntSushi/toml"
	_ "github.com/spf13/viper"
	"os"
)

// Config represents the `spv_develop.toml` file used to configure the processor
type Config struct {
	DB        DBConfig
	API       APIConfig
	RPCs      RPCsConfig      `toml:"rpcs"`
	Contracts ContractsConfig `toml:"contracts"`
	Instant   InstantConfig
}

// DBConfig configures the mysql database
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// APIConfig configures the API server
type APIConfig struct {
	Host      string
	Port      int
	ZkPools   [][]interface{} `toml:"zk-pools"`
	Makers    [][]interface{}
	Moralis   []string
	HeartBeat string `toml:"heart-beat"`
}

// RPCsConfig configures the RPC urls
type RPCsConfig struct {
	Ethereum     string `toml:"ethereum"`
	PolygonZkEVM string `toml:"polygon-zkevm"`
}

// ContractsConfig configures the contracts address
type ContractsConfig struct {
	EntryPoint string `toml:"entry-point"`
}

// InstantConfig configures the networks submit or commit wait time
type InstantConfig struct {
	EthereumSubmitWait     int64 `toml:"ethereum-submit-wait" mapstructure:"ethereum-submit-wait"`
	PolygonZkEVMSubmitWait int64 `toml:"polygon-zkevm-submit-wait" mapstructure:"polygon-zkevm-submit-wait"`
	EthereumCommitWait     int64 `mapstructure:"ethereum-commit-wait"`
	PolygonZkEVMCommitWait int64 `mapstructure:"polygon-zkevm-commit-wait"`
}

// LoadConfig loads the `spv_develop.toml` config file from a given path
func LoadConfig(path string) (Config, error) {
	var conf Config

	// Read the config file.
	data, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}

	// Replace environment variables.
	data = []byte(os.ExpandEnv(string(data)))

	// Decode the TOML data.
	if _, err := toml.Decode(string(data), &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
