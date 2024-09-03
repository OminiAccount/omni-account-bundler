package pool

import "fmt"

func (p *Pool) mockPool() {
	fmt.Println("Mock user op")
	sender := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	private_key_hex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	operation, sig, _, _, _ := MockSignedUserOperation(sender, private_key_hex, uint64(42161), 0)
	mockOp := SignedUserOperation{
		UserOperation: operation,
		Signature:     sig,
	}

	p.AddUserOp(mockOp)
}
