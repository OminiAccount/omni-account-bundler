package etherman

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	DepositTicketAddedOrder  EventOrder = "DepositTicketAdded"
	WithdrawTicketAddedOrder EventOrder = "WithdrawTicketAdded"
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
}

type AccountCreateData struct {
	Account common.Address
	Owner   common.Address
}

type DepositData struct {
	Account common.Address
	Amount  *big.Int
}

type WithdrawData struct {
	Account common.Address
	Amount  *big.Int
}
