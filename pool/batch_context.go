package pool

type BatchContext struct {
	userOps []SignedUserOperation
	tickets []TicketFull
}

func NewBatchContext(userOps []SignedUserOperation, tickets []TicketFull) *BatchContext {
	return &BatchContext{userOps: userOps, tickets: tickets}
}

func (b *BatchContext) SignedUserOperations() []SignedUserOperation {
	return b.userOps
}

// SortedTickets Returns the sorted tickets, sorted in the order of depositTicket before withdrawTicket
func (b *BatchContext) SortedTickets() []TicketFull {
	var depositTickets, withdrawTickets []TicketFull

	for _, ticket := range b.tickets {
		if ticket.Type == Deposit {
			depositTickets = append(depositTickets, ticket)
		} else if ticket.Type == Withdraw {
			withdrawTickets = append(withdrawTickets, ticket)
		}
	}

	return append(depositTickets, withdrawTickets...)
}
