package merkletree

import (
	"fmt"
	"github.com/OAB/lib/common"
	"github.com/OAB/utils/merkletreeutils"
	common2 "github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

//ky: [1719563600532945246 14890561418312037703 11938254450375952587 5329719403986311437]
//Node value: [12346118620362682079 13209918650071895561 1050462047988786108 10297651453442781900 15020833855946683413 2541896837400596712 5158482081674306993 15736419290823331982 1 0 0 0]
//sender nonce chainId01 == 1 1
//ky: [1719563600532945246 14890561418312037703 11938254450375952587 5329719403986311437]
//Node value: [12346118620362682079 13209918650071895561 1050462047988786108 10297651453442781900 15020833855946683413 2541896837400596712 5158482081674306993 15736419290823331982 1 0 0 0]

func TestSMT1(t *testing.T) {
	tree := NewSMT(nil, false)

	sender := common.HexToAddress("0x123")
	{
		keyNonce := merkletreeutils.KeyEthAddrNonceChainId(sender.String(), "0x01")
		source, _ := tree.Db.GetKeySource(keyNonce)
		fmt.Println("source", len(source) == 0)
	}
	tree.SetAccountNonce(sender.String(), new(big.Int).SetUint64(1), "0x01")
	{
		keyNonce := merkletreeutils.KeyEthAddrNonceChainId(sender.String(), "0x01")
		source, _ := tree.Db.GetKeySource(keyNonce)
		fmt.Println("source", source)
	}
	tree.SetAccountNonce(sender.String(), new(big.Int).SetUint64(2), "0x02")

	nonce, _ := tree.GetAccountNonce(common2.Address(sender), "0x01")
	fmt.Println("sender nonce chainId01 == 1", nonce.Uint64())

	nonce, _ = tree.GetAccountNonce(common2.Address(sender), "0x02")
	fmt.Println("sender nonce chainId02 == 1", nonce.Uint64())

	fmt.Println("root", tree.LastRoot())
	//
	//tree.SetAccountNonce(sender.String(), new(big.Int).SetUint64(1), "0x02")
	//
	//fmt.Println("root", tree.LastRoot())

}

func TestSMT(t *testing.T) {
	tree := NewSMT(nil, false)
	tree.SetAccountNonce(common.HexToAddress("0x11").String(), new(big.Int).SetUint64(1), "0x11")

	tree.updateDepth(0)

	sender := common.HexToAddress("0x123")
	sender1 := common.HexToAddress("0x1234")

	// EMPTY
	nonce, _ := tree.GetAccountNonce(common2.Address(sender), "0x01")
	fmt.Println("sender nonce chainId01 == 0", nonce.Uint64())

	nonce, _ = tree.GetAccountNonce(common2.Address(sender), "0x02")
	fmt.Println("sender nonce chainId02 == 0", nonce.Uint64())

	nonce, _ = tree.GetAccountNonce(common2.Address(sender1), "0x01")
	fmt.Println("sender1 nonce chainId01 == 0", nonce.Uint64())

	nonce, _ = tree.GetAccountNonce(common2.Address(sender1), "0x02")
	fmt.Println("sender1 nonce chainId02 == 0", nonce.Uint64())

	balance, _ := tree.GetAccountBalance(common2.Address(sender))
	fmt.Println("sender balance == 0", balance.Uint64())

	balance, _ = tree.GetAccountBalance(common2.Address(sender1))
	fmt.Println("sender1 balance == 0", balance.Uint64())

	tree.SetAccountNonce(sender.String(), new(big.Int).SetUint64(1), "0x01")
	tree.SetAccountNonce(sender.String(), new(big.Int).SetUint64(2), "0x02")

	tree.SetAccountNonce(sender1.String(), new(big.Int).SetUint64(2), "0x01")
	tree.SetAccountNonce(sender1.String(), new(big.Int).SetUint64(3), "0x02")

	tree.SetAccountBalance(sender.String(), new(big.Int).SetUint64(10))
	tree.SetAccountBalance(sender1.String(), new(big.Int).SetUint64(100))

	nonce, _ = tree.GetAccountNonce(common2.Address(sender), "0x01")
	fmt.Println("sender nonce chainId01 == 1", nonce.Uint64())

	nonce, _ = tree.GetAccountNonce(common2.Address(sender), "0x02")
	fmt.Println("sender nonce chainId02 == 2", nonce.Uint64())

	nonce, _ = tree.GetAccountNonce(common2.Address(sender1), "0x01")
	fmt.Println("sender1 nonce chainId01 == 2", nonce.Uint64())

	nonce, _ = tree.GetAccountNonce(common2.Address(sender1), "0x02")
	fmt.Println("sender1 nonce chainId02 == 3", nonce.Uint64())

	balance, _ = tree.GetAccountBalance(common2.Address(sender))
	fmt.Println("sender balance == 10", balance.Uint64())

	balance, _ = tree.GetAccountBalance(common2.Address(sender1))
	fmt.Println("sender1 balance == 100", balance.Uint64())

	// update
	tree.SetAccountNonce(sender.String(), new(big.Int).SetUint64(2), "0x01")
	tree.SetAccountNonce(sender.String(), new(big.Int).SetUint64(3), "0x02")

	tree.SetAccountNonce(sender1.String(), new(big.Int).SetUint64(3), "0x01")
	tree.SetAccountNonce(sender1.String(), new(big.Int).SetUint64(4), "0x02")

	tree.SetAccountBalance(sender.String(), new(big.Int).SetUint64(20))
	tree.SetAccountBalance(sender1.String(), new(big.Int).SetUint64(200))

	nonce, _ = tree.GetAccountNonce(common2.Address(sender), "0x01")
	fmt.Println("sender nonce chainId01 == 2", nonce.Uint64())

	nonce, _ = tree.GetAccountNonce(common2.Address(sender), "0x02")
	fmt.Println("sender nonce chainId02 == 3", nonce.Uint64())

	nonce, _ = tree.GetAccountNonce(common2.Address(sender1), "0x01")
	fmt.Println("sender1 nonce chainId01 == 3", nonce.Uint64())

	nonce, _ = tree.GetAccountNonce(common2.Address(sender1), "0x02")
	fmt.Println("sender1 nonce chainId02 == 4", nonce.Uint64())

	balance, _ = tree.GetAccountBalance(common2.Address(sender))
	fmt.Println("sender balance == 20", balance.Uint64())

	balance, _ = tree.GetAccountBalance(common2.Address(sender1))
	fmt.Println("sender1 balance == 200", balance.Uint64())
}

// Test vectors from JS implementation: https://github.com/0xPolygonHermez/zkevm-testvectors/blob/main/merkle-tree/smt-raw.json
// Test 'want' values come from running the JS tests with debugger attached

func TestSMT_SingleInsert(t *testing.T) {
	scenarios := []struct {
		name     string
		oldRoot  *big.Int
		k        *big.Int
		v        *big.Int
		expected string
	}{
		{
			name:     "TestSMT_Insert_0Key_0Value",
			oldRoot:  big.NewInt(0),
			k:        big.NewInt(0),
			v:        big.NewInt(0),
			expected: "0x0",
		},
		{
			name:     "TestSMT_Insert_0Key_1Value",
			oldRoot:  big.NewInt(0),
			k:        big.NewInt(0),
			v:        big.NewInt(1),
			expected: "0x42bb2f66296df03552203ae337815976ca9c1bf52cc1bdd59399ede8fea8a822",
		},
		{
			name:     "TestSMT_Insert1Key_XValue",
			oldRoot:  big.NewInt(0),
			k:        big.NewInt(1),
			v:        new(big.Int).SetUint64(1),
			expected: "0xb26e0de762d186d2efc35d9ff4388def6c96ec15f942d83d779141386fe1d2e1",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			s := NewSMT(nil, false)
			// set scenario old root if fail
			newRoot, err := s.InsertBI(scenario.k, scenario.v)
			if err != nil {
				t.Errorf("Insert failed: %v", err)
			}
			fmt.Println("root ", common.BigToHash(newRoot.NewRootScalar.ToBigInt()))
			hex := fmt.Sprintf("0x%0x", newRoot.NewRootScalar.ToBigInt())

			value12, err := s.getValueInBytes(merkletreeutils.ScalarToNodeKey(scenario.k))
			if err != nil {
				fmt.Printf("error %s", err)
			}

			fmt.Println("value ", value12)
			printNode(newRoot)
			if hex != scenario.expected {
				t.Errorf("root hash is not as expected, got %v wanted %v", hex, scenario.expected)
			}
		})
	}
}

func TestSMT_MultipleInsert(t *testing.T) {
	s := NewSMT(nil, false)
	testCases := []struct {
		root  *big.Int
		key   *big.Int
		value *big.Int
		want  string
		mode  string
	}{
		{
			big.NewInt(0),
			big.NewInt(1),
			big.NewInt(1),
			"0xb26e0de762d186d2efc35d9ff4388def6c96ec15f942d83d779141386fe1d2e1",
			"insertNotFound",
		},
		{
			nil,
			big.NewInt(2),
			big.NewInt(2),
			"0xa399847134a9987c648deabc85a7310fbe854315cbeb6dc3a7efa1a4fa2a2c86",
			"insertFound",
		},
		{
			nil,
			big.NewInt(3),
			big.NewInt(3),
			"0xb5a4b1b7a8c3a7c11becc339bbd7f639229cd14f14f76ee3a0e9170074399da4",
			"insertFound",
		},
		{
			nil,
			big.NewInt(3),
			big.NewInt(0),
			"0xa399847134a9987c648deabc85a7310fbe854315cbeb6dc3a7efa1a4fa2a2c86",
			"deleteFound",
		},
	}

	var root *big.Int
	for i, testCase := range testCases {
		if i > 0 {
			testCase.root = root
		}
		r, err := s.InsertBI(testCase.key, testCase.value)
		if err != nil {
			t.Errorf("Test case %d: Insert failed: %v", i, err)
			continue
		}

		got := toHex(r.NewRootScalar.ToBigInt())
		if got != testCase.want {
			t.Errorf("Test case %d: Root hash is not as expected, got %v, want %v", i, got, testCase.want)
		}
		if testCase.mode != r.Mode {
			t.Errorf("Test case %d: Mode is not as expected, got %v, want %v", i, r.Mode, testCase.mode)
		}

		root = r.NewRootScalar.ToBigInt()
	}
}

func TestSMT_MultipleInsert3(t *testing.T) {
	s := NewSMT(nil, false)
	testCases := []struct {
		root  *big.Int
		key   *big.Int
		value *big.Int
		want  string
		mode  string
	}{
		{
			big.NewInt(0),
			big.NewInt(18),
			big.NewInt(18),
			"0xc5eb28fcdad974a3b988a7e5441cb9dfef6f7159bbab0f9387aaf4ea192b5b88",
			"insertNotFound",
		},
		{
			nil,
			big.NewInt(19),
			big.NewInt(19),
			"0xfa2d3062e11e44668ab79c595c0c916a82036a017408377419d74523569858ea",
			"insertFound",
		},
	}

	var root *big.Int
	for i, testCase := range testCases {
		if i > 0 {
			testCase.root = root
		}
		r, err := s.InsertBI(testCase.key, testCase.value)
		if err != nil {
			t.Errorf("Test case %d: Insert failed: %v", i, err)
			continue
		}

		got := toHex(r.NewRootScalar.ToBigInt())
		if got != testCase.want {
			t.Errorf("Test case %d: Root hash is not as expected, got %v, want %v", i, got, testCase.want)
		}
		if testCase.mode != r.Mode {
			t.Errorf("Test case %d: Mode is not as expected, got %v, want %v", i, r.Mode, testCase.mode)
		}

		root = r.NewRootScalar.ToBigInt()
	}
}

func TestSMT_UpdateElement1(t *testing.T) {
	s := NewSMT(nil, false)
	testCases := []struct {
		root  *big.Int
		key   *big.Int
		value *big.Int
		want  string
		mode  string
	}{
		{
			big.NewInt(0),
			big.NewInt(1),
			big.NewInt(2),
			"0x7212762089bfe2505ebbd8f1696acb835ecaf394d0f8d191e4c026dab9ddcfa5",
			"insertNotFound",
		},
		{big.NewInt(0),
			big.NewInt(1),
			big.NewInt(3),
			"0x0f740b94e3935291daf0998666160414f14a93bb7be05ad56df4df21ff817c1d",
			"update",
		},
		{big.NewInt(0),
			big.NewInt(1),
			big.NewInt(2),
			"0x7212762089bfe2505ebbd8f1696acb835ecaf394d0f8d191e4c026dab9ddcfa5",
			"update",
		},
	}

	// set up the first root as that from the first testCase
	rs := merkletreeutils.ScalarToRoot(testCases[0].root)
	r := &SMTResponse{
		NewRootScalar: &rs,
	}
	var err error

	for i, testCase := range testCases {
		r, err = s.InsertBI(testCase.key, testCase.value)
		if err != nil {
			t.Errorf("Test case %d: Insert failed: %v", i, err)
			continue
		}

		got := toHex(r.NewRootScalar.ToBigInt())
		if got != testCase.want {
			t.Errorf("Test case %d: Root hash is not as expected, got %v, want %v", i, got, testCase.want)
		}
		if testCase.mode != r.Mode {
			t.Errorf("Test case %d: Mode is not as expected, got %v, want %v", i, r.Mode, testCase.mode)
		}
	}
}

func TestSMT_AddSharedElement2(t *testing.T) {
	s := NewSMT(nil, false)

	r1, err := s.InsertBI(big.NewInt(8), big.NewInt(2))
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	printNode(r1)
	r2, err := s.InsertBI(big.NewInt(9), big.NewInt(3))
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	printNode(r2)
	r3, err := s.InsertBI(big.NewInt(8), big.NewInt(0))
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	printNode(r3)
	r4, err := s.InsertBI(big.NewInt(9), big.NewInt(0))
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	printNode(r4)
}

func TestSMT_AddRemove128Elements(t *testing.T) {
	s := NewSMT(nil, false)
	N := 128
	var r *SMTResponse

	for i := 0; i < N; i++ {
		r, _ = s.InsertBI(big.NewInt(int64(i)), big.NewInt(int64(i+1000)))
	}

	for i := 0; i < N; i++ {
		r, _ = s.InsertBI(big.NewInt(int64(i)), big.NewInt(0))
		if r.Mode != "deleteFound" && i != N-1 {
			t.Errorf("Mode is not deleteFound, got %v", r.Mode)
		} else if r.Mode != "deleteLast" && i == N-1 {
			t.Errorf("Mode is not deleteLast, got %v", r.Mode)
		}
	}

	if r.NewRootScalar.ToBigInt().Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Root hash is not zero, got %v", toHex(r.NewRootScalar.ToBigInt()))
	}
}

func TestSMT_MultipleInsert2(t *testing.T) {
	s := NewSMT(nil, false)
	testCases := []struct {
		root  *big.Int
		key   merkletreeutils.NodeKey
		value merkletreeutils.NodeValue8
	}{
		{
			big.NewInt(0),
			merkletreeutils.ScalarToNodeKey(big.NewInt(2)),
			merkletreeutils.ScalarToNodeValue8(big.NewInt(1)),
		},
		{
			nil,
			merkletreeutils.ScalarToNodeKey(big.NewInt(10)),
			merkletreeutils.ScalarToNodeValue8(big.NewInt(1)),
		},
	}

	var root *big.Int
	for i, testCase := range testCases {
		if i > 0 {
			testCase.root = root
		}
		fmt.Println(testCase.key.GetPath())
		r, err := s.insertSingle(testCase.key, testCase.value, [4]uint64{})
		if err != nil {
			t.Errorf("Test case %d: Insert failed: %v", i, err)
			continue
		}

		root = r.NewRootScalar.ToBigInt()
		s.PrintTree()
	}
}

func printNode(n *SMTResponse) {
	fmt.Printf("Root: %s Mode: %s\n", toHex(n.NewRootScalar.ToBigInt()), n.Mode)
}

func toHex(i *big.Int) string {
	return fmt.Sprintf("0x%064x", i)
}
