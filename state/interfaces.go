package state

import (
	"github.com/OAB/etherman"
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/pool"
	"github.com/OAB/utils/chains"
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
		GetChainCli(c chains.ChainId) *etherman.EthereumClient
		EstimateGas(uint64, *big.Int, []SyncRouter.BaseStructPackedUserOperation) (*big.Int, error)
		UpdateEntryPointRoot(hexutil.Bytes, []EntryPoint.BaseStructBatchData, EntryPoint.BaseStructChainsExecuteInfo) (common.Hash, error)
		CreateAccount(common.Address) *common.Address
	}
)
