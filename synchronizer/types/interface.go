package types

import (
	"context"
	"github.com/OAB/state/types"
)

type (
	EthereumInterface interface {
		//WatchEntryPointEvent(ctx context.Context, chainId chains.ChainId, fromBlock uint64, ticketChannel chan<- pool.TicketFull) error
		WatchAAFactoryEvent(ctx context.Context, fromBlock uint64, accountMappingChannel chan<- types.AccountMapping) error
	}

	StateInterface interface {
		AddNewMapping(mapping types.AccountMapping) error
	}
)
