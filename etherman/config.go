package etherman

import (
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	VizingNetwork Network `mapstructure:"vizing"`
	ArbitNetwork  Network `mapstructure:"arbitrum"`
	BaseNetwork   Network `mapstructure:"base"`
	OpNetwork     Network `mapstructure:"optimism"`
}

// KeystoreFileConfig has all the information needed to load a private key from a key store file
type KeystoreFileConfig struct {
	// Path is the file path for the key store file
	Path string `mapstructure:"path"`

	// Password is the password to decrypt the key store file
	Password string `mapstructure:"password"`
}

// Network configures the network's chainId, rpc and contract's address
type Network struct {
	ChainId        uint64             `mapstructure:"chain-id"`
	Rpc            string             `mapstructure:"rpc"`
	EntryPoint     common.Address     `mapstructure:"entry-point"`
	AccountFactory common.Address     `mapstructure:"account-factory"`
	SyncRouter     common.Address     `mapstructure:"sync-router"`
	DataHelp       common.Address     `mapstructure:"data-help"`
	PrivateKeys    KeystoreFileConfig `mapstructure:"private_keys"`
	GenBlockNumber uint64             `mapstructure:"gen-block"`
	IsSync         uint8              `mapstructure:"is-sync"`
	BlockCheckNum  uint64             `mapstructure:"block-check-num"`
}
