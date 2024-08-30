package pool

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
	"time"
)

func TestPoolNew(t *testing.T) {
	pool := NewMemoryPool(10, 5*time.Second)
	pool.StartAutoFlush()

	// Simulate adding operations
	for i := 0; i < 25; i++ {
		fmt.Println("i", i)
		op := SignedUserOperation{
			UserOperation: UserOperation{
				Sender: common.HexToAddress(fmt.Sprintf("0x%040x", i)),
			},
			Signature: []byte(fmt.Sprintf("signature-%d", i)),
		}
		pool.AddUserOp(op)
		time.Sleep(1 * time.Second)
	}
}
