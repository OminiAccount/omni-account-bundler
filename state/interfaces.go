package state

import (
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/pool"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type (
	PoolInterface interface {
		AddTicket(ticket *pool.TicketFull)
		GetTicket(string) *pool.TicketFull
		AddSignedUserOperation(op *pool.SignedUserOperation)
		BatchContext() chan *pool.BatchContext
		Cache() error
		LoadCache()
	}

	EthereumInterface interface {
		EstimateGas(uint64, *big.Int, []byte) (*big.Int, error)
		UpdateEntryPointRoot(hexutil.Bytes, []EntryPoint.IEntryPointBatchData, EntryPoint.IEntryPointChainsExecuteInfo) (common.Hash, error)
	}
)
