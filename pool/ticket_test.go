package pool

import (
	"fmt"
	"github.com/OAB/utils/merkletree"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestEip712(t *testing.T) {
	sus := CreateUserOps()

	// encode context,used for circuit
	encodeBytes := EncodeEip712Context(sus)
	fmt.Println("BatchData Input", common.Bytes2Hex(encodeBytes))

	// BatchData Hash
	fmt.Println("BatchData Hash", HashSignedUserOperationV1s(sus))

	tree := merkletree.NewSMT(nil, false)

	fmt.Println("=====================")

	for _, userOperation := range sus {
		fmt.Println("=====================")
		fmt.Printf("Address: %s ,Nonce: %d ,ChainId %s ,OperationType: %d ,operationValue: %d \n", userOperation.Owner, userOperation.Nonce, userOperation.ChainId.String(), userOperation.OperationType, userOperation.OperationValue)

		// balance
		var balance big.Int
		balanceU256, _ := tree.GetAccountBalance(userOperation.Sender)
		balance.SetBytes(balanceU256.Bytes())
		fmt.Println("oldBalance", &balance)
		operationValue := new(big.Int).SetUint64(userOperation.OperationValue.Uint64())

		if userOperation.OperationType == 1 {
			fmt.Printf("balance %d + operationValue %d\n", &balance, operationValue)
			balance.Add(&balance, operationValue)
		} else if userOperation.OperationType == 2 {
			if balance.Cmp(operationValue) < 0 {
				fmt.Printf("current account balance %s is insufficient to withdraw %s", &balance, operationValue)
			}
			fmt.Printf("balance %d - operationValue %d\n", &balance, operationValue)
			balance.Sub(&balance, operationValue)
		}

		fmt.Printf("balance %d - CalculateGasUsed %d\n", &balance, userOperation.CalculateGasUsed())
		balance.Sub(&balance, userOperation.CalculateGasUsed())

		_, err := tree.SetAccountBalance(userOperation.Sender.String(), &balance)
		if err != nil {
			fmt.Println("accountBalance error", err)
		}
		fmt.Println("newBalance", &balance)

		// nonce
		var nonce big.Int
		nonceU256, err := tree.GetAccountNonce(userOperation.Sender, userOperation.ChainId.String())
		if err != nil {
			fmt.Println("accountNonce error", err)
		}
		nonce.SetBytes(nonceU256.Bytes())
		fmt.Println("oldNonce", nonce.Uint64())
		nonce.Add(&nonce, big.NewInt(1))
		fmt.Println("newNonce", nonce.Uint64())
		tree.SetAccountNonce(userOperation.Sender.String(), &nonce, userOperation.ChainId.String())

		//expectGasBalance, expectNonce := s.storage.Account.GetAccountInfoByAccountAndChainId(userOperation.Sender, userOperation.ChainId.ToInt().Uint64())
		//fmt.Println("expectGasBalance", expectGasBalance)
		//fmt.Println("smtGasBalance", balance)
		//fmt.Println("expectNonce", expectNonce)
		//fmt.Println("smtNonce", nonce)
		if userOperation.Nonce.Uint64() == nonce.Uint64() {
			continue
		} else {
			fmt.Println("neq")
		}
	}

	fmt.Println("newRoot", common.BigToHash(tree.LastRoot()))
}
