package types

import (
	"github.com/OAB/etherman"
	"github.com/OAB/pool"
	"github.com/OAB/state/types"
	"github.com/OAB/utils/chains"
	"github.com/ethereum/go-ethereum/common"
)

type (
	PoolInterface interface {
		AddSignedUserOperation(op *pool.SignedUserOperation)
		AddStepUserOperation(op *pool.SignedUserOperation)
		AddTicket(ticket *pool.TicketFull)
		GetTicket(id string) *pool.TicketFull
	}

	EthereumInterface interface {
		GetChains() map[chains.ChainId]*etherman.EthereumClient
	}

	StateInterface interface {
		AddNewMapping(mapping types.AccountMapping) error
		AddAccountGas(*pool.SignedUserOperation) error
		GetSignedUserOp(common.Address, common.Address, string) (*pool.SignedUserOperation, error)
	}
)
