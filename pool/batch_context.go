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

func (b *BatchContext) Tickets() []TicketFull {
	return b.tickets
}
