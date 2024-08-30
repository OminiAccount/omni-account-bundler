package pool

type pool interface {
	AddUserOp(op SignedUserOperation)
	AddTicket(ticket Ticket)
}
