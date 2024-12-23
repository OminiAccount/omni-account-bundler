package types

import (
	"github.com/OAB/etherman"
	"github.com/OAB/state"
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
		CreateAccountInfo(chains.ChainId, state.AccountMapping)
		AddAccountGas(*state.SignedUserOperation, pgx.Tx) error
		GetSignedUserOp(common.Address, common.Address, string, int) (*state.SignedUserOperation, error)
		GetHisIns() *state.HistoryManager
		GetAccountInfo(user, account common.Address) (*state.AccountInfo, error)
	}

	PoolInterface interface {
		CheckFlush(bool)
	}
)
