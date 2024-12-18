package etherman

import (
	"github.com/OAB/etherman/contracts/EntryPoint"
	"github.com/OAB/etherman/contracts/SimpleAccountFactory"
	"github.com/OAB/utils/chains"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"time"
)

type ethClientInterface interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.LogFilterer
	ethereum.TransactionReader
	ethereum.TransactionSender

	bind.DeployBackend
}

type EventOrder string

const (
	AccountCreatedOrder      EventOrder = "AccountCreated"
	VizingDepositTicketOrder EventOrder = "VizingDepositTicket"
	DepositTicketAddedOrder  EventOrder = "DepositTicketAdded"
	WithdrawTicketAddedOrder EventOrder = "WithdrawTicketAdded"
	ExecOpOrder              EventOrder = "ExecOpEvent"
)

type Order struct {
	Name EventOrder
	Pos  int
}

type Block struct {
	ID          uint64
	BlockNumber uint64
	BlockHash   common.Hash
	ParentHash  common.Hash
	NetworkID   uint
	ReceivedAt  time.Time

	or []Order
	ac []AccountCreateData
	dp []DepositData
	wt []WithdrawData
	eo []ExecOpData
}

type AccountCreateData struct {
	ChainID chains.ChainId
	Account common.Address
	Owner   common.Address
}

func ToAccountCreateData(evt *SimpleAccountFactory.SimpleAccountFactoryAccountCreated) AccountCreateData {
	return AccountCreateData{
		Owner:   evt.Owner,
		Account: evt.Account,
	}
}

type DepositData struct {
	Did     string
	User    common.Address
	Account common.Address
	Amount  *big.Int
	TxHash  string
	ChainID chains.ChainId
}

func ToVizingDepositData(evt *EntryPoint.EntryPointDepositTicketAdded) DepositData {
	return DepositData{
		Did:     "0x" + common.Bytes2Hex(evt.Did[:]),
		User:    evt.Account,
		Account: evt.Account,
		Amount:  evt.Amount,
		TxHash:  evt.Raw.TxHash.Hex(),
	}
}

func ToDepositData(evt *EntryPoint.EntryPointValueDepositAmount) DepositData {
	var encodeBytes []byte
	encodeBytes = append(encodeBytes, evt.Owner.Bytes()...)
	encodeBytes = append(encodeBytes, evt.Sender.Bytes()...)
	encodeBytes = append(encodeBytes, common.LeftPadBytes(evt.ValueDepositAmount.Bytes(), 32)...)
	return DepositData{
		Did:     crypto.Keccak256Hash(encodeBytes).Hex(),
		User:    evt.Owner,
		Account: evt.Sender,
		Amount:  evt.ValueDepositAmount,
		TxHash:  evt.Raw.TxHash.Hex(),
	}
}

type WithdrawData struct {
	Account common.Address
	Amount  *big.Int
}

func ToWithdrawData(evt *EntryPoint.EntryPointWithdrawTicketAdded) WithdrawData {
	return WithdrawData{
		Account: evt.Account,
		Amount:  evt.Amount,
	}
}

type ExecOpData struct {
	ID        string
	Phase     uint8
	InnerExec bool
	Success   bool
	Account   common.Address
	Owner     common.Address
}

func ToExecOpData(evt *EntryPoint.EntryPointUserOperationEvent) ExecOpData {
	return ExecOpData{
		ID:        common.Bytes2Hex(evt.UserOpHash[:]),
		Phase:     evt.Phase,
		InnerExec: evt.InnerExec,
		Success:   evt.Success,
		Owner:     evt.Owner,
		Account:   evt.Sender,
	}
}
