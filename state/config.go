package state

type Config struct {
	VizingChainID uint64 `mapstructure:"vizing-chain-id"`
	MaxBatches    int    `mapstructure:"max-batches"`
}
