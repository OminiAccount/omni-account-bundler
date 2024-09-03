package types

import (
	"context"
	"github.com/OAAC/pool"
	"github.com/OAAC/utils/chains"
)

type PoolInterface interface {
	AddTicket(ticket pool.TicketFull)
}

type EthereumInterface interface {
	WatchEntryPointEvent(ctx context.Context, chainId chains.ChainId, fromBlock uint64, ticketChannel chan<- pool.TicketFull) error
}
