package pool

import (
	"encoding/json"
	"fmt"
	"github.com/OAB/database/leveldb"
	"github.com/OAB/lib/common/hexutil"
	"github.com/OAB/state"
	"github.com/OAB/utils"
	"github.com/OAB/utils/db"
	"github.com/OAB/utils/merkletree"
	"github.com/OAB/utils/packutils"
	"github.com/OAB/utils/poseidon"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"path/filepath"
	"testing"
)

func TestOpUnmarshalJSON(t *testing.T) {
	raw := `{"operationType":1,"operationValue":"0x0","sender":"0x569db0654a0c9844257a7f496e02f9e7bc805c0b","exec":{"nonce":"0x1","chainId":"0x6f64","callData":"0x","mainChainGasLimit":"0x0","destChainGasLimit":"0x3d090","zkVerificationGasLimit":"0x898","mainChainGasPrice":"0x5f5e100","destChainGasPrice":"0x3b9aca00"},"innerExec":{"nonce":"0x0","chainId":"0x0","callData":"0x","mainChainGasLimit":"0x0","destChainGasLimit":"0x0","zkVerificationGasLimit":"0x0","mainChainGasPrice":"0x0","destChainGasPrice":"0x0"},"signature":"0xd0c0f69178773015bae93e1e61f1c80e3065d25074230ab48e43e96cd5fb0331014397c3afa3963a4d1837853276b4be213b92d13ecda23677e3dd7bda3c427a1c"}`
	var uo state.SignedUserOperation
	err := json.Unmarshal([]byte(raw), &uo)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", uo.UserOperation)
}

func TestPackUint(t *testing.T) {
	a := big.NewInt(12)
	b := big.NewInt(13)
	value, _ := packutils.PackUints(a, b)
	t.Log(value)
	a, b, _ = packutils.UnpackUints(value)
	t.Log(a.Uint64(), b.Uint64())

	result := poseidon.H4ToScalar([]uint64{13, 14, 0, 0})

	value, _ = packutils.PackUints(a, result)
	t.Log(value)

	a, b, _ = packutils.UnpackUints(value)
	padded := make([]byte, 16)
	copy(padded[16-len(b.Bytes()):], b.Bytes())
	t.Log(padded)
	b = big.NewInt(0).SetBytes(padded[8:])
	c := big.NewInt(0).SetBytes(padded[:8])
	t.Log(a.Uint64(), b.Uint64(), c.Uint64())
}

func TestEncodeCircuitInput(t *testing.T) {
	userOpDeposit := state.UserOperation{
		OperationType:  1,
		OperationValue: (*hexutil.Big)(big.NewInt(100)),
		Sender:         common.HexToAddress("0xd09d22e15b8c387a023811e5c1021b441b8f0e5a"),
		Exec: state.ExecData{
			Nonce:                  1,
			ChainId:                hexutil.Uint64(28516),
			CallData:               common.FromHex("0x"),
			MainChainGasLimit:      0x30d40,
			DestChainGasLimit:      0,
			ZkVerificationGasLimit: 0x898,
			MainChainGasPrice:      (*hexutil.Big)(big.NewInt(100)),
			DestChainGasPrice:      (*hexutil.Big)(big.NewInt(100)),
		},
	}
	t.Logf("%+v", userOpDeposit)
	dataHash := userOpDeposit.CalculateEIP712TypeDataHash()
	signedUserOperation := state.SignedUserOperation{
		UserOperation: &userOpDeposit,
		Signature:     signMessage(dataHash),
	}
	t.Logf("%+v", signedUserOperation.Signature)
	var sus []*state.SignedUserOperation
	sus = append(sus, &signedUserOperation)
	by := encodeCircuitInput(sus)
	t.Log(by)
}

func TestHashSignedUserOperationV1s(t *testing.T) {
	dataHash := common.FromHex("0xcd812b0013bde63ae0b83e7500f07ab18edde7b86ba1bf0c1e86de4f8377a165")
	sig := common.FromHex("0xbaad87420f9b3f3c7f024e0f6ec45b429ce0ef719971cdb2f929832fb1c1daa707b36a7b1c48f2f7af36b9caea53f9a0e9c546aa59dab067643b21b2f5fefa4e1b")

	// 1. Add the prefix
	//prefixedMessage := []byte(fmt.Sprintf("\u0019Ethereum Signed Message:\n%d%s", len(dataHash), dataHash))
	//msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(dataHash), string(dataHash))
	//hasher := sha3.NewLegacyKeccak256()
	//hasher.Write([]byte(msg))
	// 2. Hash the message using keccak256 (SHA3)
	messageHash := accounts.TextHash(dataHash)
	//hash := sha3.NewLegacyKeccak256()
	//hash.Write(prefixedMessage)
	//messageHash := hasher.Sum(nil)
	fmt.Println("message", common.BytesToHash(messageHash))

	signature := common.CopyBytes(sig)
	if len(signature) != crypto.SignatureLength {
		fmt.Printf("signature must be %d bytes long", crypto.SignatureLength)
	}
	if signature[crypto.RecoveryIDOffset] != 27 && signature[crypto.RecoveryIDOffset] != 28 {
		fmt.Printf("invalid Ethereum signature (V is not 27 or 28)")
	}
	signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	sigPublicKey, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		fmt.Printf("Failed to recover public key from signature: %v", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*sigPublicKey)

	fmt.Println(recoveredAddr.Hex())
}

func newDB() *db.MemDb {
	levelDBDir, err := filepath.Abs("./db_test")
	if _, err = utils.PathExists(levelDBDir); err != nil {
		return nil
	}
	levelDB, err := leveldb.NewLevelDB(levelDBDir)
	if err != nil {
		return nil
	}
	return db.NewMemDb(levelDB)
}

func TestEip712(t *testing.T) {
	sus := CreateUserOps()

	// encode context,used for circuit
	encodeBytes := state.EncodeEip712Context(sus)
	fmt.Println("BatchData Input", common.Bytes2Hex(encodeBytes))

	// BatchData Hash
	fmt.Println("AccInput Hash", state.HashSignedUserOperationV1s(sus))

	tree := merkletree.NewSMT(newDB(), false)

	fmt.Println("=====================")

	for _, userOperation := range sus {
		fmt.Println("=====================")
		fmt.Printf("Address: %s ,Nonce: %d ,ChainId %+v ,OperationType: %d ,operationValue: %d \n",
			userOperation.Owner, userOperation.Exec.Nonce, userOperation.Exec.ChainId, userOperation.OperationType, userOperation.OperationValue.Uint64())

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
		nonceU256, err := tree.GetAccountNonce(userOperation.Sender, userOperation.Exec.ChainId.String())
		if err != nil {
			fmt.Println("accountNonce error", err)
		}
		nonce.SetBytes(nonceU256.Bytes())
		fmt.Println("oldNonce", nonce.Uint64())
		nonce.Add(&nonce, big.NewInt(1))
		fmt.Println("newNonce", nonce.Uint64())
		tree.SetAccountNonce(userOperation.Sender.String(), &nonce, userOperation.Exec.ChainId.String())

		//expectGasBalance, expectNonce := s.storage.Account.GetAccountInfoByAccountAndChainId(userOperation.Sender, userOperation.ChainId.ToInt().Uint64())
		//fmt.Println("expectGasBalance", expectGasBalance)
		//fmt.Println("smtGasBalance", balance)
		//fmt.Println("expectNonce", expectNonce)
		//fmt.Println("smtNonce", nonce)
		if userOperation.Exec.Nonce.Uint64() == nonce.Uint64() {
			continue
		} else {
			fmt.Println("neq")
		}
	}

	fmt.Println("newRoot", common.BigToHash(tree.LastRoot()))
}
