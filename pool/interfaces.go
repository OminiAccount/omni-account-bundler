package pool

import (
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/state"
	"github.com/OAB/utils/merkletree"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type (
	EthereumInterface interface {
		EstimateGas(uint64, *big.Int, []SyncRouter.BaseStructPackedUserOperation) (*big.Int, error)
		UpdateEntryPointRoot(hexutil.Bytes, []EntryPoint.BaseStructBatchData, EntryPoint.BaseStructChainsExecuteInfo) (common.Hash, error)
	}

	StateInterface interface {
		GetHisIns() *state.HistoryManager
		GetTree() *merkletree.SMT
	}
)
