package types

import (
	"context"
	"github.com/OAAC/pool"
	"github.com/OAAC/state/types"
	"github.com/OAAC/utils/chains"
)

type (
	PoolInterface interface {
		AddTicket(ticket pool.TicketFull)
	}

	EthereumInterface interface {
		WatchEntryPointEvent(ctx context.Context, chainId chains.ChainId, fromBlock uint64, ticketChannel chan<- pool.TicketFull) error
		WatchAAFactoryEvent(ctx context.Context, fromBlock uint64, accountMappingChannel chan<- types.AccountMapping) error
	}

	StateInterface interface {
		AddNewMapping(mapping types.AccountMapping) error
	}
)
