package state

import "github.com/OAB/utils/apitypes"

type BaseConfig struct {
	VizingChainID uint64
}

type Config struct {
	BaseConfig
	HisInterval   apitypes.Duration `mapstructure:"his-interval"`
	CrossChainAPI string            `mapstructure:"cross-chain-api"`
}
