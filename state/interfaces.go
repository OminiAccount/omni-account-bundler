package state

import (
	"github.com/OAAC/pool"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type (
	PoolInterface interface {
		Context() chan pool.BatchContext
	}

	EthereumInterface interface {
		UpdateEntryPointRoot(proof hexutil.Bytes, pubicValues hexutil.Bytes) error
	}
)
