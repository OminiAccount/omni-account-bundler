package types

import (
	"github.com/OAB/etherman"
	"github.com/OAB/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
)

type (
	EthereumInterface interface {
		GetChains() map[uint64]*etherman.EthereumClient
	}

	StateInterface interface {
		UpdateUserOpStatus(uint64, int, pgx.Tx) error
		UpdateUserOpPhase(uint64, int, pgx.Tx) error
		CreateAccountInfo(uint64, state.AccountMapping)
		AddAccountGas(*state.SignedUserOperation, pgx.Tx) error
		GetSignedUserOp(common.Address, common.Address, string, int) (*state.SignedUserOperation, error)
		GetHisIns() *state.HistoryManager
		GetAccountInfo(user, account common.Address) (*state.AccountInfo, error)
		GetUser(user, account common.Address) *state.AccountInfo
	}

	PoolInterface interface {
		CheckFlush(bool)
	}
)
