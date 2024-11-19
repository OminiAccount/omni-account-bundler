package pool

type BatchContext struct {
	userOps []*SignedUserOperation
}

func NewBatchContext(userOps []*SignedUserOperation) *BatchContext {
	return &BatchContext{userOps: userOps}
}

func (b *BatchContext) SignedUserOperations() []*SignedUserOperation {
	return b.userOps
}
