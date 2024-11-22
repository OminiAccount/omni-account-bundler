package chains

// ChainId represents a chain ID as uint64
type ChainId uint64

// ChainInfo holds the ID and name of a chain
type ChainInfo struct {
	ID   ChainId
	Name string
}

// Define all supported ChainId and names
var (
	MainnetChain         = ChainInfo{ID: 1, Name: "Mainnet"}
	SepoliaChain         = ChainInfo{ID: 11155111, Name: "Sepolia"}
	ArbitrumMainnetChain = ChainInfo{ID: 42161, Name: "Arbitrum Mainnet"}
	ArbitrumSepoliaChain = ChainInfo{ID: 421614, Name: "Arbitrum Sepolia"}
	LocalNode            = ChainInfo{ID: 31337, Name: "LocalNode"}
	// Add more chains as needed...
)

// supportedChains is a map with all the ChainInfo
var supportedChains map[ChainId]ChainInfo

const MaxChainInfoLength = 200

// The init function is used to initialize the supportedChains
func init() {
	supportedChains = make(map[ChainId]ChainInfo, MaxChainInfoLength)
	chains := []ChainInfo{
		MainnetChain,
		SepoliaChain,
		ArbitrumMainnetChain,
		ArbitrumSepoliaChain,
		LocalNode,
		// Add more chains as needed...
	}

	for _, chain := range chains {
		supportedChains[chain.ID] = chain
	}
}

// IsSupport checks if the ChainId is supported
func (id ChainId) IsSupport() bool {
	_, ok := supportedChains[id]
	return ok
}

// Name returns the name of the ChainId if supported
func (id ChainId) Name() (string, bool) {
	if info, ok := supportedChains[id]; ok {
		return info.Name, true
	}
	return "", false
}

// Equals checks if the ChainId equals the given ChainInfo
func (id ChainId) Equals(chain ChainInfo) bool {
	return id == chain.ID
}
