package etherman

import (
	"github.com/OAB/utils/chains"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	VizingChainID chains.ChainId `mapstructure:"vizing-chain-id"`
	Networks      []Network      `mapstructure:"networks"`
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
	PrivateKeys    KeystoreFileConfig `mapstructure:"private_keys"`
	GenBlockNumber uint64             `mapstructure:"gen-block"`
	IsSync         uint8              `mapstructure:"is-sync"`
}
