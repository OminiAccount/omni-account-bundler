package ticket

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Ticket struct {
	User      common.Address
	Amount    hexutil.Big
	TimeStamp uint64
}
