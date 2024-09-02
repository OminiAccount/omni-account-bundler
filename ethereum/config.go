package ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Networks []Network `toml:"networks"`
}

// Network configures the network's chainId, rpc and contract's address
type Network struct {
	ChainId    uint64         `toml:"chain-id"`
	Rpc        string         `toml:"rpc"`
	EntryPoint common.Address `toml:"entry-point"`
}

func (n *Network) UnmarshalTOML(data interface{}) error {
	// Decode the data into a map
	m, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected type: %T", data)
	}

	// Convert ChainID to uint64
	if chainID, ok := m["chain-id"].(int64); ok {
		n.ChainId = uint64(chainID)
	} else {
		return fmt.Errorf("invalid chain-id")
	}

	if rpc, ok := m["rpc"].(string); ok {
		n.Rpc = rpc
	} else {
		return fmt.Errorf("invalid rpc")
	}

	// Convert Contract to common.Address
	if contract, ok := m["entry-point"].(string); ok {
		n.EntryPoint = common.HexToAddress(contract)
	} else {
		return fmt.Errorf("invalid contract address")
	}

	return nil
}
