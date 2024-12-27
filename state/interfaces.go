package state

import (
	"github.com/OAB/etherman"
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SyncRouter"
	"github.com/OAB/etherman/contracts/ZKVizingAADataHelp"
	"github.com/OAB/lib/common/hexutil"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type (
	EthereumInterface interface {
		GetChainCli(c uint64) *etherman.EthereumClient
		EstimateGas(uint64, *big.Int, []SyncRouter.BaseStructPackedUserOperation) (*big.Int, error)
		UpdateEntryPointRoot(hexutil.Bytes, []EntryPoint.BaseStructBatchData, EntryPoint.BaseStructChainsExecuteInfo) (common.Hash, error)
		CreateVizingAccount(common.Address) (*common.Address, uint64, error)
		CreateOtherAccount(common.Address, uint64, uint64) (*common.Address, error)
		DecodeSwapCalldata(uint64, hexutil.Bytes) (ZKVizingAADataHelp.BaseStructV3SwapParams, error)
	}
)
