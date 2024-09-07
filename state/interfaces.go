package state

import (
	"github.com/OAAC/pool"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	PoolInterface interface {
		Context() chan pool.BatchContext
		Cache()
	}

	EthereumInterface interface {
		UpdateEntryPointRoot(proof hexutil.Bytes, pubicValues hexutil.Bytes) (*types.Transaction, error)
	}
)
