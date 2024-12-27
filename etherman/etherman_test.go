package etherman

import (
	"context"
	"crypto/ecdsa"
	"github.com/OAB/etherman/contracts/SimpleAccount"
	"github.com/OAB/etherman/contracts/SimpleAccountFactory"
	"github.com/OAB/etherman/contracts/ZKVizingAADataHelp"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"testing"
)

func TestDepositGas(t *testing.T) {
	cli, err := ethclient.Dial("https://rpc-sepolia.vizing.com")
	if err != nil {
		t.Fatal(err)
	}

	smcAdr := common.HexToAddress("0x8e7cb5af07117d972a91394e988d4aed95973694")
	smcMgr, err := SimpleAccount.NewSimpleAccount(smcAdr, cli)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("aa contract is loaded")
	balance, err := cli.BalanceAt(context.Background(), smcAdr, nil)
	if err != nil {
		t.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	t.Log("合约账户余额：", ethValue)

	auth := getAuth(t)
	auth.Value = big.NewInt(3000000000000000000)
	tx, err := smcMgr.DepositGas(auth, big.NewInt(3))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx.Hash().Hex())
}

func TestDecodeCalldata(t *testing.T) {
	cli, err := ethclient.Dial("https://rpc-sepolia.vizing.com")
	if err != nil {
		t.Fatal(err)
	}

	smcAdr := common.HexToAddress("0xA517139395309572D6d6D4436b6E754743e032C3")
	smcMgr, err := ZKVizingAADataHelp.NewZKVizingAADataHelp(smcAdr, cli)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dataHelp contract is loaded")

	opts := new(bind.CallOpts)
	opts.From = common.HexToAddress("0x7c38C1646213255f62dB509688B8fA062e0Ed8e4")
	decodeParams, err := smcMgr.DecodeUniswapV3Data(opts, []byte{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(decodeParams)
}

func TestCreateAccount(t *testing.T) {
	cli, err := ethclient.Dial("https://rpc-sepolia.vizing.com")
	if err != nil {
		t.Fatal(err)
	}

	smcAdr := common.HexToAddress("0xd35D83D4d325558a42756c9E304421222dC3708D")
	smcMgr, err := SimpleAccountFactory.NewSimpleAccountFactory(smcAdr, cli)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dataHelp contract is loaded")

	owner := common.HexToAddress("0x19B6250c9FF3a4F4fC21A1fa3bec3849cA80dB69")
	nonce := big.NewInt(0)
	opts := new(bind.CallOpts)
	opts.From = common.HexToAddress("0x7c38C1646213255f62dB509688B8fA062e0Ed8e4")
	adr, err := smcMgr.GetAccountAddress(opts, owner, nonce)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(adr)

	auth := getAuth(t)
	tx, err := smcMgr.CreateAccount(auth, owner, nonce)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx.Hash())
}

func TestAccountInfo(t *testing.T) {
	cli, err := ethclient.Dial("https://rpc-sepolia.vizing.com")
	if err != nil {
		t.Fatal(err)
	}

	smcAdr := common.HexToAddress("0xd35D83D4d325558a42756c9E304421222dC3708D")
	smcMgr, err := SimpleAccountFactory.NewSimpleAccountFactory(smcAdr, cli)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dataHelp contract is loaded")

	opts := new(bind.CallOpts)
	//opts.From = common.HexToAddress("0x7c38C1646213255f62dB509688B8fA062e0Ed8e4")
	owner := common.HexToAddress("0x69299A9dFcC793E9780A0115BF3B45b4dECa2463")
	ua, err := smcMgr.Accounts(opts, owner)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ua)
}

func getAuth(t *testing.T) *bind.TransactOpts {
	prk := "28c17d6cb8a410397cc6622e35b23fecd4fa04efb6144ea04ada720c455082b6"
	privateKey, err := crypto.HexToECDSA(prk)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	t.Logf("eoa: %s", addr)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.From = addr
	auth.Value = big.NewInt(0)
	auth.Nonce = nil
	auth.GasPrice = nil
	auth.GasLimit = 0
	return auth
}
