package types

import (
	"github.com/OAB/etherman"
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/chains"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
)

type (
	EthereumInterface interface {
		GetChains() map[chains.ChainId]*etherman.EthereumClient
	}

	StateInterface interface {
		UpdateUserOpStatus(uint64, int, pgx.Tx) error
		UpdateUserOpPhase(uint64, int, pgx.Tx) error
		CreateAccountInfo(chains.ChainId, types.AccountMapping)
		AddAccountGas(*pool.SignedUserOperation, pgx.Tx) error
		GetSignedUserOp(common.Address, common.Address, string, int) (*pool.SignedUserOperation, error)
		GetHisIns() *state.HistoryManager
	}
)
