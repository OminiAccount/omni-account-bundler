package types

import "github.com/OAAC/pool"

type PoolInterface interface {
	AddUserOp(op pool.SignedUserOperation)
}
