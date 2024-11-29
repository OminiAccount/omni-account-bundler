// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EntryPoint

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EntryPointMemoryUserOp is an auto generated low-level Go binding around an user-defined struct.
type EntryPointMemoryUserOp struct {
	Sender                        common.Address
	VerificationGasLimit          *big.Int
	CallGasLimit                  *big.Int
	PaymasterVerificationGasLimit *big.Int
	PaymasterPostOpGasLimit       *big.Int
	PreVerificationGas            *big.Int
	Paymaster                     common.Address
	MaxFeePerGas                  *big.Int
	MaxPriorityFeePerGas          *big.Int
}

// EntryPointUserOpInfo is an auto generated low-level Go binding around an user-defined struct.
type EntryPointUserOpInfo struct {
	MUserOp       EntryPointMemoryUserOp
	UserOpHash    [32]byte
	Prefund       *big.Int
	ContextOffset *big.Int
	PreOpGas      *big.Int
	Sender        common.Address
	UserAddr      common.Address
}

// IStakeManagerDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerDepositInfo struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}

// ITicketManagerTicket is an auto generated low-level Go binding around an user-defined struct.
type ITicketManagerTicket struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
}

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	ChainId            *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	UserAddr           common.Address
}

// EntryPointMetaData contains all meta data concerning the EntryPoint contract.
var EntryPointMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accountOwners\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"accountAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addDepositTicket\",\"inputs\":[{\"name\":\"accountAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addWithdrawTicket\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegateAndRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositTickets\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staked\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"stake\",\"type\":\"uint112\",\"internalType\":\"uint112\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"withdrawTime\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dstEids\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositInfo\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structIStakeManager.DepositInfo\",\"components\":[{\"name\":\"deposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staked\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"stake\",\"type\":\"uint112\",\"internalType\":\"uint112\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"withdrawTime\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSenderAddress\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getUserOpHash\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"userAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleOps\",\"inputs\":[{\"name\":\"ops\",\"type\":\"tuple[]\",\"internalType\":\"structPackedUserOperation[]\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"userAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"innerHandleOp\",\"inputs\":[{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"opInfo\",\"type\":\"tuple\",\"internalType\":\"structEntryPoint.UserOpInfo\",\"components\":[{\"name\":\"mUserOp\",\"type\":\"tuple\",\"internalType\":\"structEntryPoint.MemoryUserOp\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterVerificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterPostOpGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymaster\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prefund\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preOpGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"actualGasCost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"smtRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"syncBatch\",\"inputs\":[{\"name\":\"syncInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"syncBatchMock\",\"inputs\":[{\"name\":\"syncInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"syncRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDstEids\",\"inputs\":[{\"name\":\"_dstEids\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSyncRouter\",\"inputs\":[{\"name\":\"_syncRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateVerifier\",\"inputs\":[{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyBatch\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicValues\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyBatchMock\",\"inputs\":[{\"name\":\"depositTickets\",\"type\":\"tuple[]\",\"internalType\":\"structITicketManager.Ticket[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"withdrawTickets\",\"type\":\"tuple[]\",\"internalType\":\"structITicketManager.Ticket[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyBatchMockUserOp\",\"inputs\":[{\"name\":\"allUserOps\",\"type\":\"tuple[]\",\"internalType\":\"structPackedUserOperation[]\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"userAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawTickets\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AccountDeployed\",\"inputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"factory\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"paymaster\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeforeExecution\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositTicketAdded\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ticketHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositTicketDeleted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ticketHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"totalDeposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PostOpRevertReason\",\"inputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"revertReason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignatureAggregatorChanged\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserOperationEvent\",\"inputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymaster\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"actualGasCost\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"actualGasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserOperationPrefundTooLow\",\"inputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserOperationRevertReason\",\"inputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"revertReason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawTicketAdded\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ticketHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawTicketDeleted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ticketHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DelegateAndRevert\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"ret\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"FailedOp\",\"inputs\":[{\"name\":\"opIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"FailedOpWithRevert\",\"inputs\":[{\"name\":\"opIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"inner\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionType\",\"inputs\":[{\"name\":\"optionType\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PostOpReverted\",\"inputs\":[{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderAddressResult\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SignatureValidationFailed\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a06040523462000036576200001462000059565b604051615b1962000315823960805181818161285e015261574c0152615b1990f35b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b6040513d5f823e3d90fd5b6200006433620000ae565b60405161039b810167ffffffffffffffff811182821017620000a85762000092829161039b62005e2e843990565b03905ff08015620000a257608052565b6200004e565b6200003a565b620000b990620000bb565b565b620000b9906200011e565b620000d9620000d6620000d69290565b90565b73ffffffffffffffffffffffffffffffffffffffff1690565b620000d690620000c6565b6200010890620000d9565b9052565b602081019291620000b99190620000fd565b62000128620001a0565b8062000149620001426200013c5f620000f2565b620000d9565b91620000d9565b146200015a57620000b990620002b3565b6200019c620001695f620000f2565b6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152918291600483016200010c565b0390fd5b620000b962000214565b620000d6620000d6620000d69290565b620000d66001620001aa565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff905b9181191691161790565b9062000208620000d66200021092620001aa565b8254620001c6565b9055565b620000b962000222620001ba565b6008620001f4565b620000d690620000d9565b620000d690546200022a565b9073ffffffffffffffffffffffffffffffffffffffff90620001ea565b620000d690620000d99073ffffffffffffffffffffffffffffffffffffffff1682565b620000d6906200025e565b620000d69062000281565b90620002ab620000d662000210926200028c565b825462000241565b620002bf600962000235565b620002cc82600962000297565b9062000304620002fd7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936200028c565b916200028c565b916200030f60405190565b5f90a356fe60806040526004361015610011575f80fd5b5f3560e01c806301ffc9a7146101e057806305dc0922146101db5780631f14e001146101d657806325350f08146101d15780632b7ac3f3146101cc5780632f6feea7146101c75780633a106262146101c2578063400c1a7c146101bd57806340eef1d0146101b85780635287ce12146101b3578063611f11e9146101ae5780636cbe19af146101a957806370a08231146101a4578063715018a61461019f57806374e811941461019a5780637fcb522714610195578063850aaf621461019057806389b287f01461018b5780638da5cb5b1461018657806397fc007c146101815780639b249f691461017c578063a6e79d9b14610177578063afc8d53214610172578063bc6be42e1461016d578063c9c8a7bd14610168578063cd71176514610163578063d384c3821461015e578063f2fde38b146101595763fc7e286d0361021057610e3b565b610d05565b610cea565b610cc2565b610c5c565b610c43565b610c2a565b610bec565b610ba0565b610b88565b610b6d565b610b54565b610b18565b610aca565b610ab6565b610a6e565b610a53565b610a38565b610a19565b61098c565b610915565b61063c565b6105c7565b610571565b6104ce565b610481565b610386565b61032a565b61023e565b7fffffffff0000000000000000000000000000000000000000000000000000000081165b0361021057565b5f80fd5b90503590610221826101e5565b565b906020828203126102105761023791610214565b90565b9052565b346102105761026b610259610254366004610223565b610e6c565b60405191829182901515815260200190565b0390f35b80610209565b905035906102218261026f565b906020828203126102105761023791610275565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b80549192918310156102ea5760086102e06004925f5260205f2090565b8185040193060290565b610296565b610237916008021c5b63ffffffff1690565b9061023791546102ef565b600780548210156102105761023791610324916102c3565b90610301565b346102105761026b610345610340366004610282565b61030c565b6040519182918263ffffffff909116815260200190565b5f91031261021057565b610237916008021c81565b906102379154610366565b6102375f80610371565b346102105761039636600461035c565b61026b6103a161037c565b6040515b9182918290815260200190565b73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff8116610209565b90503590610221826103cb565b9060208282031261021057610237916103e7565b6103b26102376102379273ffffffffffffffffffffffffffffffffffffffff1690565b61023790610408565b6102379061042b565b9061044790610434565b5f5260205260405f2090565b610237916008021c6103b2565b906102379154610453565b6102379061047c6004915f9261043d565b610460565b346102105761026b61049c6104973660046103f4565b61046b565b6040519182918273ffffffffffffffffffffffffffffffffffffffff909116815260200190565b6102375f6006610460565b34610210576104de36600461035c565b61026b61049c6104c3565b909182601f830112156102105781359167ffffffffffffffff831161021057602001926060830284011161021057565b60608183031261021057803567ffffffffffffffff811161021057826105409183016104e9565b929093602083013567ffffffffffffffff811161021057610566836102379286016104e9565b9390946040016103e7565b346102105761058d610584366004610519565b93929092610ff8565b604051005b90816101209103126102105790565b9060208282031261021057813567ffffffffffffffff8111610210576102379201610592565b346102105761026b6103a16105dd3660046105a1565b61104e565b909182601f830112156102105781359167ffffffffffffffff831161021057602001926001830284011161021057565b9060208282031261021057813567ffffffffffffffff81116102105761063892016105e2565b9091565b346102105761058d61064f366004610612565b90611609565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b90601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810190811067ffffffffffffffff8211176106c257604052565b610655565b906102216106d460405190565b9283610682565b67ffffffffffffffff81116106c257602090601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160190565b0190565b90825f939282370152565b9092919261073a610735826106db565b6106c7565b93818552602085019082840111610210576102219261071a565b9080601f830112156102105781602061023793359101610725565b9190610120838203126102105761081f9061078b6101206106c7565b9361079682826103e7565b85526107a58260208301610275565b60208601526107b78260408301610275565b60408601526107c98260608301610275565b60608601526107db8260808301610275565b60808601526107ed8260a08301610275565b60a08601526107ff8260c083016103e7565b60c08601526108118260e08301610275565b60e086015261010001610275565b610100830152565b91906101e083820312610210576108b79061084260e06106c7565b9361084d828261076f565b855261085d826101208301610275565b6020860152610870826101408301610275565b6040860152610883826101608301610275565b6060860152610896826101808301610275565b60808601526108a9826101a083016103e7565b60a08601526101c0016103e7565b60c0830152565b916102208383031261021057823567ffffffffffffffff811161021057826108e7918501610754565b926108f58360208301610827565b9261020082013567ffffffffffffffff81116102105761063892016105e2565b346102105761026b6103a161092b3660046108be565b929190916116ff565b60a0810192916102219190805182526020808201511515908301526040808201516dffffffffffffffffffffffffffff169083015260608082015163ffffffff169083015260809081015165ffffffffffff16910152565b346102105761026b6109a76109a23660046103f4565b6119b4565b60405191829182610934565b909182601f830112156102105781359167ffffffffffffffff831161021057602001926020830284011161021057565b9160408383031261021057823567ffffffffffffffff811161021057610a0e836102379286016109b3565b9390946020016103e7565b61058d610a273660046109e3565b916120f6565b6102375f6005610460565b3461021057610a4836600461035c565b61026b61049c610a2d565b346102105761026b6103a1610a693660046103f4565b61221c565b3461021057610a7e36600461035c565b61058d612268565b909160608284031261021057610237610a9f84846103e7565b93610aad81602086016103e7565b93604001610275565b61058d610ac4366004610a86565b916123b6565b346102105761058d610add3660046103f4565b6124de565b91909160408184031261021057610af983826103e7565b92602082013567ffffffffffffffff81116102105761063892016105e2565b3461021057610b28366004610ae2565b91612539565b9060208282031261021057813567ffffffffffffffff81116102105761063892016109b3565b346102105761058d610b67366004610b2e565b906127fc565b3461021057610b7d36600461035c565b61026b61049c612806565b346102105761058d610b9b3660046103f4565b612827565b3461021057610bb0366004610612565b90612855565b90610447565b610237916008021c5b60ff1690565b906102379154610bbc565b61023790610be76003915f92610bb6565b610bcb565b346102105761026b610259610c02366004610282565b610bd6565b91906040838203126102105761023790610c2181856103e7565b93602001610275565b346102105761058d610c3d366004610c07565b90612936565b346102105761058d610c56366004610612565b90612b87565b346102105761058d610c6f3660046109e3565b91612db5565b60608183031261021057803567ffffffffffffffff81116102105782610c9c9183016105e2565b929093602083013567ffffffffffffffff811161021057610566836102379286016105e2565b61058d610cd0366004610c75565b93929092612eb4565b61023790610be76002915f92610bb6565b346102105761026b610259610d00366004610282565b610cd9565b346102105761058d610d183660046103f4565b613099565b6102379081565b6102379054610d1d565b61023790610bc5565b6102379054610d2e565b6102379060081c6dffffffffffffffffffffffffffff1690565b6102379054610d41565b6102379060781c6102f8565b6102379054610d65565b6102379060981c65ffffffffffff1690565b6102379054610d7b565b610da290600161043d565b610dab81610d24565b91610db860018301610d37565b91610dc560018201610d5b565b916102376001610dd6818501610d71565b9301610d8d565b9095949261022194610e20610e2d92610e09608096610e0060a088019c5f890152565b15156020870152565b6dffffffffffffffffffffffffffff166040850152565b63ffffffff166060830152565b019065ffffffffffff169052565b346102105761026b610e56610e513660046103f4565b610d97565b91610e6395939560405190565b95869586610ddd565b80610eb87f73ba5dd9000000000000000000000000000000000000000000000000000000005b917fffffffff000000000000000000000000000000000000000000000000000000001690565b148015610f0d575b8015610edc575b908115610ed2575090565b61023791506130a2565b5080610f077f22274c2300000000000000000000000000000000000000000000000000000000610e92565b14610ec7565b5080610f387f519d11fa00000000000000000000000000000000000000000000000000000000610e92565b14610ec0565b67ffffffffffffffff81116106c25760208091020190565b919060608382031261021057610f9790610f7060606106c7565b93610f7b82826103e7565b8552610f8a8260208301610275565b6020860152604001610275565b6040830152565b90929192610fae61073582610f3e565b93818552606060208601920283019281841161021057915b838310610fd35750505050565b6020606091610fe28486610f56565b815201920191610fc6565b610237913691610f9e565b61022194506110149261100e9192909492610fed565b92610fed565b906130cf565b90815273ffffffffffffffffffffffffffffffffffffffff90911660208201526060810192916102219160400152565b0152565b611057906131df565b61108c61106330610434565b9161107d4661107160405190565b9485936020850161101a565b60208201810382520382610682565b61109e611097825190565b9160200190565b2090565b610237906103b2565b61023790546110a2565b156110bc57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600560248201527f4e455153520000000000000000000000000000000000000000000000000000006044820152606490fd5b0390fd5b9061022191611158336111526111386103b260056110ab565b9173ffffffffffffffffffffffffffffffffffffffff1690565b146110b5565b611569565b91906040838203126102105782359067ffffffffffffffff82116102105761118a81610237938601610754565b936020016103e7565b90505190610221826103cb565b905051906102218261026f565b5f5b8381106111be5750505f910152565b81810151838201526020016111af565b909291926111de610735826106db565b9381855260208501908284011161021057610221926111ad565b9080601f83011215610210578151610237926020016111ce565b919091610120818403126102105761122b6101206106c7565b926112368183611193565b845261124581602084016111a0565b6020850152604082015167ffffffffffffffff8111610210578161126a9184016111f8565b6040850152606082015167ffffffffffffffff8111610210578161128f9184016111f8565b60608501526112a181608084016111a0565b60808501526112b38160a084016111a0565b60a08501526112c58160c084016111a0565b60c085015260e08201519167ffffffffffffffff8311610210576112ee8261081f9483016111f8565b60e086015261010001611193565b92919061130b61073582610f3e565b93818552602080860192028101918383116102105781905b838210611331575050505050565b815167ffffffffffffffff8111610210576020916113528784938701611212565b815201910190611323565b9080601f83011215610210578151610237926020016112fc565b9060208282031261021057815167ffffffffffffffff811161021057610237920161135d565b6113be6113c7602093610716936113b2815190565b80835293849260200190565b958691016111ad565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690565b805173ffffffffffffffffffffffffffffffffffffffff168252906102379061010080611492611453611441610120860161142f60208a01516020890152565b6040890151878203604089015261139d565b6060880151868203606088015261139d565b61146260808801516080870152565b61147160a088015160a0870152565b61148060c088015160c0870152565b60e087015185820360e087015261139d565b94015173ffffffffffffffffffffffffffffffffffffffff16910152565b90610237916113ef565b906114d06114c6835190565b8083529160200190565b90816114e26020830284019460200190565b925f915b8383106114f557505050505090565b90919293946020611518611511838560019503875289516114b0565b9760200190565b93019301919392906114e6565b929160206115416102219360408701908782035f8901526114ba565b94019073ffffffffffffffffffffffffffffffffffffffff169052565b6040513d5f823e3d90fd5b61157d6115949261159b929081019061115d565b92906020611589825190565b818301019101611377565b46906131e8565b6115a430610434565b63c9c8a7bd9190803b15610210576115d15f80946115dc6115c460405190565b9788968795869460e01b90565b845260048401611525565b03925af18015611604576115ed5750565b610221905f6115fc8183610682565b81019061035c565b61155e565b906102219161111f565b1561161a57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4141393220696e7465726e616c2063616c6c206f6e6c790000000000000000006044820152606490fd5b6102376102376102379290565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b906116bd565b9190565b9081156116c8570490565b611686565b610237612710611679565b610237610800611679565b60208082526102379291019061139d565b610237913691610725565b92915a9361171c336117166111386103b230610434565b14611613565b82519461172a604087015190565b906117506117405a61173c603f611679565b0290565b61174a6040611679565b906116b3565b6117706116b96102376117688661071660808e015190565b6107166116cd565b106118bc57610237966117bb935f93611787825190565b6117936116b95f611679565b116117c1575b5050506117a96117b2915a900390565b60808601510190565b949093926116f4565b91613364565b6117fd916117f9916117e7855173ffffffffffffffffffffffffffffffffffffffff1690565b6117f35f939293611679565b9061328e565b1590565b611809575b8080611799565b90915061181c6118176116d8565b6132a0565b805161182a6116b95f611679565b1161183e575b5060019190506117a9611802565b6118695f61184d602089015190565b93015173ffffffffffffffffffffffffffffffffffffffff1690565b90916118b261189f6118997f0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb9390565b93610434565b936118a960405190565b918291826116e3565b0390a35f80611830565b7fdeaddead000000000000000000000000000000000000000000000000000000005f5260205ffd5b61023760a06106c7565b6118f66118e4565b90602080808080865f5b8152015f8152015f8152015f8152015f905250565b6102376118ee565b9061022161199c600161192e6118e4565b9461193f61193b82610d24565b8752565b61195661194d838301610d37565b15156020880152565b61197b611964838301610d5b565b6dffffffffffffffffffffffffffff166040880152565b611996611989838301610d71565b63ffffffff166060880152565b01610d8d565b65ffffffffffff166080840152565b6102379061191d565b6119cb610237916119c3611915565b50600161043d565b6119ab565b91909161012081840312610210576119e96101206106c7565b926119f481836103e7565b8452611a038160208401610275565b6020850152604082013567ffffffffffffffff81116102105781611a28918401610754565b6040850152606082013567ffffffffffffffff81116102105781611a4d918401610754565b6060850152611a5f8160808401610275565b6080850152611a718160a08401610275565b60a0850152611a838160c08401610275565b60c085015260e08201359167ffffffffffffffff831161021057611aac8261081f948301610754565b60e0860152610100016103e7565b929190611ac961073582610f3e565b93818552602080860192028101918383116102105781905b838210611aef575050505050565b813567ffffffffffffffff811161021057602091611b1087849387016119d0565b815201910190611ae1565b610237913691611aba565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b9190611b5e565b9290565b8203918211611b6957565b611b26565b90611b7b610735836106db565b918252565b611b8a6007611b6e565b7f6761735573656400000000000000000000000000000000000000000000000000602082015290565b610237611b80565b506102379060208101906103e7565b50610237906020810190610275565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1368390030181121561021057016020813591019167ffffffffffffffff82116102105736829003831361021057565b91906113c781611c42816107169560209181520190565b809561071a565b9061023790610100611d40611d36611cd6611cbb6101208601611c89611c6f8a80611bbb565b73ffffffffffffffffffffffffffffffffffffffff168852565b611ca0611c9960208b018b611bca565b6020890152565b611cad60408a018a611bd9565b9088830360408a0152611c2b565b611cc86060890189611bd9565b908783036060890152611c2b565b611ced611ce66080890189611bca565b6080870152565b611d04611cfd60a0890189611bca565b60a0870152565b611d1b611d1460c0890189611bca565b60c0870152565b611d2860e0880188611bd9565b9086830360e0880152611c2b565b9482810190611bbb565b73ffffffffffffffffffffffffffffffffffffffff16910152565b9061023791611c49565b90357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee13683900301811215610210570190565b818352916020019081611dae6020830284019490565b92835f925b848410611dc35750505050505090565b9091929394956020611def611de88385600195038852611de38b88611d65565b611d5b565b9860200190565b940194019294939190611db3565b9392906115416020916102219460408801918883035f8a0152611d98565b81810292918115918404141715611b6957565b611e3b6102376102379290565b6fffffffffffffffffffffffffffffffff1690565b610237906102f8565b6102379060201c6102f8565b6102379060401c6102f8565b6102379060601c6102f8565b6102379060801c6102f8565b6102379060a01c6102f8565b6102379060c01c6102f8565b6102379060e01c6102f8565b90600190611ed3611ec9611ebf855490565b8084529260200190565b935f5260205f2090565b5f9261201d575b611ee5565b50505090565b5490808310612000575b808310611fe3575b808310611fc6575b808310611fa9575b808310611f8c575b808310611f6f575b808310611f52575b8210611f2c575b80611edf565b82611f4960019394611f3f602094611ea1565b63ffffffff169052565b0191015f611f26565b9192602081611f66600193611f3f86611e95565b01930191611f1f565b9192602081611f83600193611f3f86611e89565b01930191611f17565b9192602081611fa0600193611f3f86611e7d565b01930191611f0f565b9192602081611fbd600193611f3f86611e71565b01930191611f07565b9192602081611fda600193611f3f86611e65565b01930191611eff565b9192602081611ff7600193611f3f86611e59565b01930191611ef7565b9192602081612014600193611f3f86611e50565b01930191611eef565b60078301821115611eda579260016020611f3f6120aa600894838080808080808f549761204d81611f3f8b611e50565b0161205b81611f3f8a611e59565b0161206981611f3f89611e65565b0161207781611f3f88611e71565b0161208581611f3f87611e7d565b0161209381611f3f86611e89565b016120a181611f3f85611e95565b01928391611ea1565b01940192019161201d565b949391610221936120e8611541926120da60609560808b01908b82035f8d0152611ead565b9089820360208b015261139d565b90878203604089015261139d565b91905a906121076115948286611b1b565b61211030610434565b849163c9c8a7bd91803b15610210576115d15f80946121316115c460405190565b03925af18015611604576121b39361217f9261215492612207575b505a90611b53565b9461216686612161611bb3565b613627565b61107d8561217360405190565b94859360208501611dfd565b926121a461219f61218e6136aa565b926121996002611679565b90611e1b565b611e2e565b6121ad5f611e2e565b91613787565b916121c96121c46121c460056110ab565b610434565b909291600791903463513ff44d833b15610210576115dc6121fc935f976121ef60405190565b998a988997889660e01b90565b8652600486016120b5565b612216905f6115fc8183610682565b5f61214c565b5f61222b610237926119c35f90565b01610d24565b612239613794565b610221612257565b6103b26102376102379290565b61023790612241565b6102216122635f61224e565b6137e0565b610221612231565b1561227757565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600360248201527f564e4500000000000000000000000000000000000000000000000000000000006044820152606490fd5b156122dd57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600560248201527f55534552300000000000000000000000000000000000000000000000000000006044820152606490fd5b61023760606106c7565b9060ff905b9181191691161790565b9061236561023761236c92151590565b8254612346565b9055565b9073ffffffffffffffffffffffffffffffffffffffff9061234b565b9061239c61023761236c92610434565b8254612370565b9081526040810192916102219160200152565b91906123c3823414612270565b6123dd836123d66111386103b25f61224e565b14156122d6565b429261241e8161241986610f97876124126123f661233c565b73ffffffffffffffffffffffffffffffffffffffff9096168652565b6020850152565b613878565b91612434600161242f856002610bb6565b612355565b61244761244282600461043d565b6110ab565b6124566111386103b25f61224e565b146124ac575b5090919261249161248d7f2537ccb2b2fe11ea53e1afd15883c085006e0571b5b4f33dbf8becaf183f4e0b93610434565b9390565b936124a761249e60405190565b928392836123a3565b0390a3565b6124c1906124bc8391600461043d565b61238c565b5f61245c565b610221906124d36138c9565b61022190600561238c565b610221906124c7565b9091610716908390809361071a565b9091610237926124e7565b3d1561251a576125103d611b6e565b903d5f602084013e565b606090565b90151581526040602082018190526102379291019061139d565b905f9283929161255461254b60405190565b928392836124f6565b03915af4612560612501565b9061111b61256d60405190565b9283927f994105540000000000000000000000000000000000000000000000000000000084526004840161251f565b90610221916125a96138c9565b6127f0565b9190600861234b9102916125e17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff841b90565b921b90565b91906125f761023761236c93611679565b9083546125ae565b610221915f916125e6565b818110612615575050565b806126225f6001936125ff565b0161260a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff612658916020036008021c90565b8154169055565b91909182821061266e57505050565b6102219260049060089060070181900461269260086007870104945f5260205f2090565b93840193019306025f81116126a8575b5061260a565b6126b59060018303612628565b5f6126a2565b906801000000000000000081116106c257816126d8610221935490565b9082815561265f565b63ffffffff8116610209565b35610237816126e1565b9190600861234b9102916125e163ffffffff841b90565b919067ffffffffffffffff82116106c25761273061273a9161023784866126bb565b925f5260205f2090565b60088204915f5b8381106127a357506008830290035f811161275d575b50505050565b925f935f5b8181106127775750505001555f808080612757565b909194602061279960019261278e6102378a6126ed565b9085600402906126f7565b9601929101612762565b5f805b600881106127bb575083820155600101612741565b959060206127dc6001926127d1610237866126ed565b908a600402906126f7565b920196016127a6565b90610221929161270e565b906102219160076127e5565b906102219161259c565b61023760096110ab565b6102219061281c6138c9565b61022190600661238c565b61022190612810565b906020828203126102105761023791611193565b602080825261023793910191611c2b565b906020906128827f0000000000000000000000000000000000000000000000000000000000000000610434565b61289a5f63570e1a369593956128a56115c460405190565b845260048401612844565b03925af180156116045761111b915f91612908575b506040519182917f6ca7b8060000000000000000000000000000000000000000000000000000000083526004830173ffffffffffffffffffffffffffffffffffffffff909116815260200190565b612929915060203d811161292f575b6129218183610682565b810190612830565b5f6128ba565b503d612917565b9061294a826123d66111386103b25f61224e565b42916129638161241985610f97866124126123f661233c565b90612974600161242f846003610bb6565b90919261249161248d7ff2d23449d3c6fc1266c59d62d6cb28ebcd25d872fc4c70d906c571e7247a8b0093610434565b90610221916129bd336111526111386103b260056110ab565b612ae2565b919060608382031261021057610f97906129dc60606106c7565b936129e78282611193565b85526129f682602083016111a0565b60208601526040016111a0565b90929192612a1361073582610f3e565b93818552606060208601920283019281841161021057915b838310612a385750505050565b6020606091612a4784866129c2565b815201920191612a2b565b9080601f8301121561021057815161023792602001612a03565b9060808282031261021057815167ffffffffffffffff81116102105781612a9491840161135d565b92612aa282602085016111a0565b92604081015167ffffffffffffffff81116102105783612ac3918301612a52565b92606082015167ffffffffffffffff8111610210576102379201612a52565b612b22612af9612b1093612b1c939081019061115d565b93906020612b05825190565b818301019101612a6c565b919490939246906131e8565b926130cf565b612b2b30610434565b63c9c8a7bd9190803b15610210576115d15f8094612b58612b4b60405190565b9889968795869460e01b90565b03925af19182156116045761022192612b72575b5061397a565b612b81905f6115fc8183610682565b5f612b6c565b90610221916129a4565b90612ba49291612b9f6139b8565b612cd0565b610221613a1c565b90611b7b61073583610f3e565b61023760e06106c7565b6102376101206106c7565b612bd6612bc3565b90602080808080808080808a5f8152015f8152015f5b8152015f8152015f611900565b610237612bce565b612c09612bb9565b90602080808080808088612bec612bf9565b610237612c01565b5f5b828110612c3157505050565b602090612c3c612c1b565b8184015201612c25565b90610221612c5383612bac565b92602080612c618693610f3e565b9201910390612c23565b90612c74825190565b8110156102ea576020809102010190565b9035907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee13682900301821215610210570190565b908210156102ea5760206102379202810190612c85565b929181612cdc81612c46565b91612ce65f611679565b82811015612d2157612d1c90612d14612cff8287612c6b565b518290612d0d828a8d612cb9565b9091613ab5565b505060010190565b612ce6565b50929093612d2e5f611679565b937fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972612d5960405190565b5f90a1612d655f611679565b945b86861015612da457612d97612d9e9161071688612d868c898c91612cb9565b612d908b8b612c6b565b5191613efb565b9560010190565b94612d67565b9450925094505061022192506142e1565b906102219291612b91565b9290612ddc90610237959360408601918683035f880152611c2b565b926020818503910152611c2b565b9080601f830112156102105781602061023793359101611aba565b9080601f830112156102105781602061023793359101610f9e565b9060808282031261021057813567ffffffffffffffff81116102105781612e48918401612dea565b92612e568260208501610275565b92604081013567ffffffffffffffff81116102105783612e77918301612e05565b92606082013567ffffffffffffffff8111610210576102379201612e05565b9392906115416020916102219460408801918883035f8a0152611c2b565b9190612ec66121c46121c460066110ab565b84939183919063b8e72af6823b1561021057612f02612ef7925f96612eea60405190565b998a978896879660e01b90565b865260048601612dc0565b03915afa91821561160457612f3c92613001575b50612f36612f2684830183612e20565b95929390939193959146906131e8565b946130cf565b612f4530610434565b859363c9c8a7bd91803b15610210576115d15f8094612f73612f6660405190565b998a968795869460e01b90565b03925af190811561160457612fa993612f9092612b72575061397a565b61107d84612f9d60405190565b94859360208501612e96565b90612fca612fb56136aa565b5f6121ad612fc461c350611e2e565b91611e2e565b91612fdb6121c46121c460056110ab565b63513ff44d926007929490823b15610210575f946115dc86926121fc946121ef60405190565b613010905f6115fc8183610682565b5f612f16565b61022190613022613794565b806130326111386103b25f61224e565b1461304057610221906137e0565b61111b61304c5f61224e565b6040519182917f1e4fbdf70000000000000000000000000000000000000000000000000000000083526004830173ffffffffffffffffffffffffffffffffffffffff909116815260200190565b61022190613016565b6130cb7f01ffc9a700000000000000000000000000000000000000000000000000000000610e92565b1490565b91906130d9835190565b916130e2825190565b906130ec5f611679565b84811015613149576131449061313e6131058289612c6b565b5161310f8161438a565b6131386020613132835173ffffffffffffffffffffffffffffffffffffffff1690565b92015190565b9061441e565b60010190565b6130ec565b5091935091506131585f611679565b828110156131d9576131d49061313e6131718287612c6b565b5161317b816144c7565b6131a761244260046131a1845173ffffffffffffffffffffffffffffffffffffffff1690565b9061043d565b906131ce60206131326121c4845173ffffffffffffffffffffffffffffffffffffffff1690565b91614633565b613158565b50915050565b61108c90614748565b6131f15f611679565b6131fa5f611679565b806132096116b9610237865190565b101561328657613224602061321e8386612c6b565b51015190565b8414613238575b6001016131fa565b6131fa565b90613256613233918361324b6116b98390565b0361325e5760010190565b91905061322b565b61327f61326b8587612c6b565b518683916132798383612c6b565b52612c6b565b5060010190565b508152919050565b5f93849392909160208451940192f190565b3d908082116132c4575b50604051906020810182016040528082525f602083013e90565b90505f6132aa565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6003111561330357565b6132cc565b90610221826132f9565b61023790613308565b61023a90613312565b61104a6133526060936102219698979561334560808601925f87019061331b565b848203602086015261139d565b966040830152565b610237600a611679565b909392905a92855192613376846147e1565b9161339860c086015173ffffffffffffffffffffffffffffffffffffffff1690565b806133a86111386103b25f61224e565b036134dd57506133f1916133e091506107166133d8875173ffffffffffffffffffffffffffffffffffffffff1690565b975b5a900390565b936107166080613132604084015190565b6134058461340060808a015190565b900390565b908181116134ac575b505082025b6040860151938185101561347a5750506134366134306002613308565b91613308565b03613452576102219193613449816148fd565b84905f90614854565b7fdeadaa51000000000000000000000000000000000000000000000000000000005f5260205ffd5b6134969061349061022196989594939889900390565b9061480a565b506134a36134305f613308565b14859192614854565b936107166134cb6134c36134d59461341396980390565b61173c61335a565b61174a6064611679565b92905f61340e565b8096916134e8815190565b6134f46116b95f611679565b1161350c575b50506133f1916107166133e0926133da565b838502918661351e6134306002613308565b0361352a575b506134fa565b6121c461353691610434565b90637c627b2161354760808a015190565b92909290889294908890803b15610210575f9561358487936135799561356c60405190565b9a8b998a98899660e01b90565b865260048601613324565b0393f190816135f1575b506135e1576135a1565b92829450613524565b61111b6135af6118176116d8565b6040519182917fad7954bc000000000000000000000000000000000000000000000000000000008352600483016116e3565b6133f1916107166133e092613598565b613600905f6115fc8183610682565b5f61358e565b929160206136226102219360408701908782035f89015261139d565b940152565b9061366d6102219261107d61363b60405190565b9384926004602085017f9710a9d000000000000000000000000000000000000000000000000000000000815201613606565b614956565b61367f6102376102379290565b61ffff1690565b6102376003613672565b61023a9061ffff1660f01b90565b61071681600293613690565b61107d6102376136b8613686565b6040519283916020830161369e565b9392919080946136e0866136da5f611679565b906149c7565b6136f66136ee61367f613686565b9161ffff1690565b0361370657610237949550613768565b61111b613716876136da5f611679565b6040519182917f3a51740d0000000000000000000000000000000000000000000000000000000083526004830161ffff909116815260200190565b610bc56102376102379290565b6102376001613751565b9261377891926102379450614a37565b61378061375e565b9091614b97565b61023792919060606136c7565b61379c612806565b6137a8611138336103b2565b036137af57565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602490fd5b6137ea60096110ab565b6137f582600961238c565b906138296138237f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610434565b91610434565b9161383360405190565b80806124a7565b61023a9073ffffffffffffffffffffffffffffffffffffffff1660601b90565b6020939261387260148361387288956107169761383a565b01918252565b61108c613899825173ffffffffffffffffffffffffffffffffffffffff1690565b61107d6138b260406138ac602087015190565b95015190565b6040519485936020850161385a565b610221613794565b6102216138c1565b156138d857565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081528061111b6004820160208082526004908201527f4e52495a00000000000000000000000000000000000000000000000000000000604082015260600190565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9061234b565b9061397361023761236c9290565b825461393d565b61022190613998816139916116b96102375f611679565b14156138d1565b5f613965565b6102376002611679565b9061397361023761236c92611679565b6139c26008610d24565b6139d06116b961023761399e565b146139e7576102216139e061399e565b60086139a8565b6040517f3ee5aeb5000000000000000000000000000000000000000000000000000000008152600490fd5b6102376001611679565b6102216139e0613a12565b610237610237610237926effffffffffffffffffffffffffffff1690565b15613a4c57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f41413934206761732076616c756573206f766572666c6f7700000000000000006044820152606490fd5b356102378161026f565b9290915f925a8251613ac78184614c1e565b613ada613ad38461104e565b6020860152565b602081015196613b4f613aee60a084015190565b8917613afb604085015190565b17613b07606085015190565b17613b13608085015190565b17613b1f60e085015190565b17613b2c61010085015190565b17613b486116b96effffffffffffffffffffffffffffff613a27565b1115613a45565b613b5882614d64565b90613b668983888885614e2c565b98613b71855a900390565b8110613c2c5750613b9b60c0606094015173ffffffffffffffffffffffffffffffffffffffff1690565b613baa6111386103b25f61224e565b03613be9575b509261071660a0613bdd6080956133da613bd6613be4976102376102219c9b60408d0152565b60608a0152565b9201613aab565b910152565b6102379750613be4925060a0613bdd6080956133da613bd66102219a9996613c1861071697898d90849261517c565b9e9099505096999a50505095505050613bb0565b90613c4661111b92613c41613c41885a900390565b61508f565b6040519182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052601e908201527f41413236206f76657220766572696669636174696f6e4761734c696d69740000606082015260800190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe13682900301821215610210570180359067ffffffffffffffff8211610210576020019136829003831361021057565b613d1d613d176102379263ffffffff1690565b60e01b90565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b805173ffffffffffffffffffffffffffffffffffffffff1682526102219190610100908190613d7660208201516020860152565b613d8560408201516040860152565b613d9460608201516060860152565b613da360808201516080860152565b613db260a082015160a0860152565b60c08181015173ffffffffffffffffffffffffffffffffffffffff1690850152613de160e082015160e0860152565b0151910152565b906101c060c061022193613e025f8201515f860190613d42565b613e126020820151610120860152565b613e226040820151610140860152565b613e326060820151610160860152565b613e426080820151610180860152565b60a081015173ffffffffffffffffffffffffffffffffffffffff166101a0850152015173ffffffffffffffffffffffffffffffffffffffff16910152565b61022080825261023795939192613ea892613e9d92850191611c2b565b936020830190613de8565b61020081840391015261139d565b929160206136226102219360408701908782035f890152611c49565b6102208082526102379492613ea891613e9d919084019061139d565b91908201809211611b6957565b9092915a613f0d610237606084015190565b61400e60205f600460405199613f27816060810190613cb1565b905f6003831161420d575b613f5b7fa647e9ed00000000000000000000000000000000000000000000000000000000610e92565b1485146141d4575050613ff990613fc0613f75868b015190565b91613fb2613f8260405190565b938492878a85017fa647e9ed00000000000000000000000000000000000000000000000000000000815201613eb6565b878201810382520382610682565b613feb878a613fd26340eef1d0613d04565b93613fdc60405190565b9687958a870190815201613ed2565b858201810382520382610682565b828151910182305af15f519760405215151590565b6140185750505050565b9091929394506140255f90565b3d6020146141c7575b806140587fdeaddead000000000000000000000000000000000000000000000000000000006116b9565b036140d25761111b8561406a60405190565b9182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052600f908201527f41413935206f7574206f66206761730000000000000000000000000000000000606082015260800190565b92935090916141007fdeadaa51000000000000000000000000000000000000000000000000000000006116b9565b03614146575061412361411761413b925a90611b53565b60808401515b90613eee565b604083015192614132816148fd565b83905f90614854565b905b5f808080612757565b6141b9614117849361415c60206141c197015190565b85515173ffffffffffffffffffffffffffffffffffffffff166141806118176116d8565b916141af61189f6118997f676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e9390565b0390a35a90611b53565b916002613364565b9061413d565b5060205f803e5f5161402e565b614208925090613feb9088908b906141ef6340eef1d0613d04565b946141f960405190565b9788968b880190815201613e80565b613ff9565b508035613f32565b1561421c57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4141393020696e76616c69642062656e656669636961727900000000000000006044820152606490fd5b1561428257565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f41413931206661696c65642073656e6420746f2062656e6566696369617279006044820152606490fd5b5f6102219261430582936121c4816142fe6138236103b28861224e565b1415614215565b9061430f60405190565b90818003925af161431e612501565b5061427b565b1561432b57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600560248201527f54484e45310000000000000000000000000000000000000000000000000000006044820152606490fd5b61439381613878565b906143af6143aa6143a5846002610bb6565b610d37565b614324565b6143be5f61242f846002610bb6565b6143e16020613132835173ffffffffffffffffffffffffffffffffffffffff1690565b9161440c7f03b8ee44084307fb90dbd84aef4a2bbaef367bb8a49bfbfc6e930a12f2f749ea92610434565b9261441961249e60405190565b0390a2565b90614429908261480a565b906144196144577f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c492610434565b926103a560405190565b1561446857565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600560248201527f54484e45320000000000000000000000000000000000000000000000000000006044820152606490fd5b6144d081613878565b906144e76144e26143a5846003610bb6565b614461565b6144f65f61242f846003610bb6565b6145196020613132835173ffffffffffffffffffffffffffffffffffffffff1690565b9161440c7ff89de5c5fbaa24098f4988752b4e8c7b8f11fc756a1688b549173fcb9396ea1992610434565b1561454b57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f576974686472617720616d6f756e7420746f6f206c61726765000000000000006044820152606490fd5b61023a90610434565b91602061022192949361104a60408201965f8301906145aa565b156145d457565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6661696c656420746f20776974686472617700000000000000000000000000006044820152606490fd5b610221926146d45f9384936146896146506102376001849061043d565b61466c866146656116b96102378b8601610d24565b1115614544565b8661468261467b828401610d24565b8890611b53565b91016139a8565b8190846146b67fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb92610434565b926146cc6146c360405190565b928392836145b3565b0390a2610434565b906146de60405190565b90818003925af16146ed612501565b506145cd565b959391989796949290986101008701995f88016147249173ffffffffffffffffffffffffffffffffffffffff169052565b602087015260408601526060850152608084015260a083015260c082015260e00152565b610237614754826153f9565b9161107d61476460208301613aab565b9161477b6147756040830183613cb1565b90615413565b9461478c6147756060840184613cb1565b9161479960808201613aab565b6147a560a08301613aab565b906147c36147756147b860c08601613aab565b9460e0810190613cb1565b9396989490919293946147d560405190565b998a9860208a016146f3565b6147f261010061313260e084015190565b80821461480657610237919048019061541f565b5090565b61482d61481e610237610237936119c35f90565b9261482884610d24565b613eee565b9182906139a8565b901515815260608101939261022192909160409161104a906020830152565b909192614862602083015190565b906148aa60c05f61488c8180880151015173ffffffffffffffffffffffffffffffffffffffff1690565b950151015173ffffffffffffffffffffffffffffffffffffffff1690565b939490946148f86148e56148df6148df7f58eb13224f9e7bb63720c3b18c69c98cfe7fb1459db62324c36836a7d30146f39690565b96610434565b966148ef60405190565b93849384614835565b0390a4565b61492b5f8061490d602085015190565b930151015173ffffffffffffffffffffffffffffffffffffffff1690565b6138296138237f1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff289390565b610221906001615ad3565b1561496857565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f746f55696e7431365f6f75744f66426f756e64730000000000000000000000006044820152606490fd5b6002906149ef6149d5825190565b6149e86116b96102378761411d88611679565b1015614961565b01015190565b61023a906fffffffffffffffffffffffffffffffff1660801b90565b601081614a23610716938396956149f5565b0180926149f5565b610716816010936149f5565b9080614a5b614a455f611e2e565b916fffffffffffffffffffffffffffffffff1690565b03614a7d575061023761107d91614a7160405190565b92839160208301614a2b565b906102379061107d614a8e60405190565b93849260208401614a11565b939291908094614aad866136da5f611679565b614abb6136ee61367f613686565b0361370657610237949550614b50565b614ad89061ffff166136ee565b019061ffff8211611b6957565b610716614afd92602092614af7815190565b94859290565b938491016111ad565b61023a9060ff1660f81b90565b600193614b4185610237989795614b39614b32600296614b4998614ae5565b8092614b06565b018092613690565b018092614b06565b0190614ae5565b506102379161107d90614b6161375e565b94614b85614b75614b70835190565b615485565b614b7f6001613672565b90614acb565b60405191929196879560208701614b13565b6102379291906060614a9a565b35610237816103cb565b6102376034611679565b15614bbf57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4141393320696e76616c6964207061796d6173746572416e64446174610000006044820152606490fd5b614caa90614c48614c2e82614ba4565b73ffffffffffffffffffffffffffffffffffffffff168452565b614c6b614c5f614c5a60808401613aab565b615532565b60408601526020850152565b614c81614c7a60a08301613aab565b60a0850152565b614ca0614c93614c5a60c08401613aab565b60e0860152610100850152565b60e0810190613cb1565b9081614cb86116b95f611679565b1115614d175761022192614cfd929091614ceb9190614ce682614cdf6116b9610237614bae565b1015614bb8565b615684565b608084015260c0830193919260600152565b73ffffffffffffffffffffffffffffffffffffffff169052565b505061022190614d46614d295f61224e565b73ffffffffffffffffffffffffffffffffffffffff1660c0830152565b614d59614d525f611679565b6060830152565b6080613be45f611679565b6102379061173c60e0613132614d966117a9614d8d614d84602088015190565b60408801510190565b60608701510190565b60a08501510190565b9060208282031261021057610237916111a0565b73ffffffffffffffffffffffffffffffffffffffff90911681526040810192916102219160200152565b908152606060208201819052600d908201527f414132332072657665727465640000000000000000000000000000000000000060808201526102379160a082015b91604081840391015261139d565b93949392915f956020614e9560c05f850151614e7a614e615f83015173ffffffffffffffffffffffffffffffffffffffff1690565b978a614e7289926040810190613cb1565b9290916156fe565b015173ffffffffffffffffffffffffffffffffffffffff1690565b92614e9f5f611679565b9084614eb06111386103b25f61224e565b14615065575b614f095f614ec66121c489610434565b63f563d33c9693614eee9060c0015173ffffffffffffffffffffffffffffffffffffffff1690565b9496614f14614efc60405190565b9889978896879460e01b90565b845260048401614db3565b0393f15f9181615035575b5061502e5750614fe6565b614f396111386103b25f61224e565b14614f4357505050565b610237614f5191600161043d565b91614f5b83610d24565b90818311614f71575061022192915f9103614682565b61111b90614f7e60405190565b9182917f220266b6000000000000000000000000000000000000000000000000000000008352600483019081526040602082018190526017908201527f41413231206469646e2774207061792070726566756e64000000000000000000606082015260800190565b83614ff26118176116d8565b9061111b614fff60405190565b9283927f65c8fd4d00000000000000000000000000000000000000000000000000000000845260048401614ddd565b9550614f2a565b61505791925060203d811161505e575b61504f8183610682565b810190614d9f565b905f614f1f565b503d615045565b90506150708561221c565b8681111561508857506150825f611679565b90614eb6565b8603615082565b61107d61366d610221926150a260405190565b9283916004602084017ff5b1bba90000000000000000000000000000000000000000000000000000000081520190815260200190565b91906040838203126102105782519067ffffffffffffffff821161021057615105816102379386016111f8565b936020016111a0565b60409061104a61512f6102219597969460608401908482035f860152611c49565b966020830152565b908152606060208201819052600d908201527f414133332072657665727465640000000000000000000000000000000000000060808201526102379160a08201614e1e565b919391906060945f945a915f810151936151ad60c086015173ffffffffffffffffffffffffffffffffffffffff1690565b946151bc61023787600161043d565b6151c581610d24565b858110615384576151f86121c46151f260605f9897966151ed615228978b6146828e829a0390565b015190565b99610434565b63854678fb9061523361521160208c9394939798015190565b9861521b60405190565b998a988997889560e01b90565b85526004850161510e565b0393f1805f8093909261535e575b506153495750506001615301576116b9611b5a61525e925a900390565b116152665750565b61111b9061527360405190565b9182917f220266b6000000000000000000000000000000000000000000000000000000008352600483019081526040602082018190526027908201527f41413336206f766572207061796d6173746572566572696669636174696f6e4760608201527f61734c696d697400000000000000000000000000000000000000000000000000608082015260a00190565b8261530d6118176116d8565b9061111b61531a60405190565b9283927f65c8fd4d00000000000000000000000000000000000000000000000000000000845260048401615137565b9650945061525e906116b990611b5a906133da565b90925061537d91503d805f833e6153758183610682565b8101906150d8565b915f615241565b61111b8961539160405190565b9182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052601e908201527f41413331207061796d6173746572206465706f73697420746f6f206c6f770000606082015260800190565b61540e610237916154075f90565b503561042b565b61042b565b60405190829082372090565b908082101561542c575090565b905090565b5f809161543c815190565b906020016a636f6e736f6c652e6c6f675afa50565b6102376102376102379261ffff1690565b61023a90613751565b91602061022192949361104a60408201965f830190615462565b806154946116b961ffff615451565b116154a25761023790613672565b601061111b6154b060405190565b9283927f6dfcc6500000000000000000000000000000000000000000000000000000000084526004840161546b565b611e3b610237610237926fffffffffffffffffffffffffffffffff1690565b6102379060801c6154df565b61023790611679565b610237610237610237926fffffffffffffffffffffffffffffffff1690565b9061023761557b61557561219f61556f7fffffffffffffffffffffffffffffffff0000000000000000000000000000000087166154fe565b6154fe565b9561550a565b93615513565b92615513565b6102376014611679565b90939293848311610210578411610210578101920390565b357fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001690601481106155d3575090565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000615602916014036008021b90565b1690565b6102379060601c610408565b61023790615606565b6102376024611679565b357fffffffffffffffffffffffffffffffff00000000000000000000000000000000169060108110615655575090565b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000615602916010036008021b90565b9161023761557b6156f761556a6156de6156bb6156b66156b08a8a6156a7615581565b905f919261558b565b906155a3565b615612565b966156e461556a6156de8b846156cf615581565b906156d861561b565b9261558b565b90615625565b98906156ee61561b565b906156d8614bae565b9394615513565b909290918161570f6116b95f611679565b0361571a5750505050565b83515173ffffffffffffffffffffffffffffffffffffffff1692833b6157426116b95f611679565b03615a31576157707f0000000000000000000000000000000000000000000000000000000000000000610434565b855160209081015163570e1a369092909261289a5f87938996615795614efc60405190565b0393f1908115611604575f91615a13575b50806157b76111386103b25f61224e565b1461599e57806157dc73ffffffffffffffffffffffffffffffffffffffff8716611138565b0361592a573b6157ee6116b95f611679565b146158b65750615815916156b6916156b0916014906156d861580f5f611679565b92611679565b9061584760c05f615827602087015190565b9394950151015173ffffffffffffffffffffffffffffffffffffffff1690565b6158726118997fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d9390565b936158aa61587f60405190565b9283928373ffffffffffffffffffffffffffffffffffffffff91821681529116602082015260400190565b0390a35f808080612757565b61111b906158c360405190565b9182917f220266b60000000000000000000000000000000000000000000000000000000083526004830190815260406020808301829052908201527f4141313520696e6974436f6465206d757374206372656174652073656e646572606082015260800190565b61111b8261593760405190565b9182917f220266b60000000000000000000000000000000000000000000000000000000083526004830190815260406020808301829052908201527f4141313420696e6974436f6465206d7573742072657475726e2073656e646572606082015260800190565b61111b826159ab60405190565b9182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052601b908201527f4141313320696e6974436f6465206661696c6564206f72204f4f470000000000606082015260800190565b615a2b915060203d811161292f576129218183610682565b5f6157a6565b61111b90615a3e60405190565b9182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052601f908201527f414131302073656e64657220616c726561647920636f6e737472756374656400606082015260800190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffd5b600103615aa6576102219061543156fea26469706673582212200b78d549e315b33ac349511785f7232e71a9f635a80cd73a0960058d96b1ca9564736f6c6343000817003360806040523461001a5760405161037c61001f823961037c90f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c63570e1a3603610055576100c1565b909182601f830112156100555781359167ffffffffffffffff831161005557602001926001830284011161005557565b5f80fd5b9060208282031261005557813567ffffffffffffffff81116100555761007f9201610025565b9091565b73ffffffffffffffffffffffffffffffffffffffff1690565b90565b73ffffffffffffffffffffffffffffffffffffffff909116815260200190565b565b34610055576100e96100dd6100d7366004610059565b906102c9565b6040519182918261009f565b0390f35b61009c61009c61009c9290565b90939293848311610055578411610055578101920390565b357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000169060148110610142575090565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000610171916014036008021b90565b1690565b61008361009c61009c9273ffffffffffffffffffffffffffffffffffffffff1690565b61009c9060601c610175565b61009c90610198565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b90601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810190811067ffffffffffffffff82111761021a57604052565b6101ad565b906100bf61022c60405190565b92836101da565b67ffffffffffffffff811161021a57602090601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160190565b90825f939282370152565b9092919261028e61028982610233565b61021f565b93818552602085019082840111610055576100bf9261026e565b61009c913691610279565b61008361009c61009c9290565b61009c906102b3565b5f906103266103206020946102db5f90565b5061030c610307610301866014856102fb6102f58c6100ed565b926100ed565b926100fa565b90610112565b6101a4565b939061031860146100ed565b9080926100fa565b906102a8565b90828483519301915af15f5191901561033b57565b905061009c5f6102c056fea264697066735822122052d3a20a86c952f15263581c8402e720e450e11a46c04704b2537be6115e14b464736f6c63430008170033",
}

// EntryPointABI is the input ABI used to generate the binding from.
// Deprecated: Use EntryPointMetaData.ABI instead.
var EntryPointABI = EntryPointMetaData.ABI

// EntryPointBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EntryPointMetaData.Bin instead.
var EntryPointBin = EntryPointMetaData.Bin

// DeployEntryPoint deploys a new Ethereum contract, binding an instance of EntryPoint to it.
func DeployEntryPoint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EntryPoint, error) {
	parsed, err := EntryPointMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EntryPointBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EntryPoint{EntryPointCaller: EntryPointCaller{contract: contract}, EntryPointTransactor: EntryPointTransactor{contract: contract}, EntryPointFilterer: EntryPointFilterer{contract: contract}}, nil
}

// EntryPoint is an auto generated Go binding around an Ethereum contract.
type EntryPoint struct {
	EntryPointCaller     // Read-only binding to the contract
	EntryPointTransactor // Write-only binding to the contract
	EntryPointFilterer   // Log filterer for contract events
}

// EntryPointCaller is an auto generated read-only Go binding around an Ethereum contract.
type EntryPointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryPointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EntryPointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryPointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EntryPointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryPointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EntryPointSession struct {
	Contract     *EntryPoint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntryPointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EntryPointCallerSession struct {
	Contract *EntryPointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EntryPointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EntryPointTransactorSession struct {
	Contract     *EntryPointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EntryPointRaw is an auto generated low-level Go binding around an Ethereum contract.
type EntryPointRaw struct {
	Contract *EntryPoint // Generic contract binding to access the raw methods on
}

// EntryPointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EntryPointCallerRaw struct {
	Contract *EntryPointCaller // Generic read-only contract binding to access the raw methods on
}

// EntryPointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EntryPointTransactorRaw struct {
	Contract *EntryPointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntryPoint creates a new instance of EntryPoint, bound to a specific deployed contract.
func NewEntryPoint(address common.Address, backend bind.ContractBackend) (*EntryPoint, error) {
	contract, err := bindEntryPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EntryPoint{EntryPointCaller: EntryPointCaller{contract: contract}, EntryPointTransactor: EntryPointTransactor{contract: contract}, EntryPointFilterer: EntryPointFilterer{contract: contract}}, nil
}

// NewEntryPointCaller creates a new read-only instance of EntryPoint, bound to a specific deployed contract.
func NewEntryPointCaller(address common.Address, caller bind.ContractCaller) (*EntryPointCaller, error) {
	contract, err := bindEntryPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntryPointCaller{contract: contract}, nil
}

// NewEntryPointTransactor creates a new write-only instance of EntryPoint, bound to a specific deployed contract.
func NewEntryPointTransactor(address common.Address, transactor bind.ContractTransactor) (*EntryPointTransactor, error) {
	contract, err := bindEntryPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntryPointTransactor{contract: contract}, nil
}

// NewEntryPointFilterer creates a new log filterer instance of EntryPoint, bound to a specific deployed contract.
func NewEntryPointFilterer(address common.Address, filterer bind.ContractFilterer) (*EntryPointFilterer, error) {
	contract, err := bindEntryPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntryPointFilterer{contract: contract}, nil
}

// bindEntryPoint binds a generic wrapper to an already deployed contract.
func bindEntryPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EntryPointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntryPoint *EntryPointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntryPoint.Contract.EntryPointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntryPoint *EntryPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPoint.Contract.EntryPointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntryPoint *EntryPointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntryPoint.Contract.EntryPointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntryPoint *EntryPointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntryPoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntryPoint *EntryPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntryPoint *EntryPointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntryPoint.Contract.contract.Transact(opts, method, params...)
}

// AccountOwners is a free data retrieval call binding the contract method 0x25350f08.
//
// Solidity: function accountOwners(address owner) view returns(address accountAddress)
func (_EntryPoint *EntryPointCaller) AccountOwners(opts *bind.CallOpts, owner common.Address) (common.Address, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "accountOwners", owner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountOwners is a free data retrieval call binding the contract method 0x25350f08.
//
// Solidity: function accountOwners(address owner) view returns(address accountAddress)
func (_EntryPoint *EntryPointSession) AccountOwners(owner common.Address) (common.Address, error) {
	return _EntryPoint.Contract.AccountOwners(&_EntryPoint.CallOpts, owner)
}

// AccountOwners is a free data retrieval call binding the contract method 0x25350f08.
//
// Solidity: function accountOwners(address owner) view returns(address accountAddress)
func (_EntryPoint *EntryPointCallerSession) AccountOwners(owner common.Address) (common.Address, error) {
	return _EntryPoint.Contract.AccountOwners(&_EntryPoint.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EntryPoint *EntryPointCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EntryPoint *EntryPointSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EntryPoint.Contract.BalanceOf(&_EntryPoint.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EntryPoint *EntryPointCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EntryPoint.Contract.BalanceOf(&_EntryPoint.CallOpts, account)
}

// DepositTickets is a free data retrieval call binding the contract method 0xd384c382.
//
// Solidity: function depositTickets(bytes32 ) view returns(bool)
func (_EntryPoint *EntryPointCaller) DepositTickets(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "depositTickets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositTickets is a free data retrieval call binding the contract method 0xd384c382.
//
// Solidity: function depositTickets(bytes32 ) view returns(bool)
func (_EntryPoint *EntryPointSession) DepositTickets(arg0 [32]byte) (bool, error) {
	return _EntryPoint.Contract.DepositTickets(&_EntryPoint.CallOpts, arg0)
}

// DepositTickets is a free data retrieval call binding the contract method 0xd384c382.
//
// Solidity: function depositTickets(bytes32 ) view returns(bool)
func (_EntryPoint *EntryPointCallerSession) DepositTickets(arg0 [32]byte) (bool, error) {
	return _EntryPoint.Contract.DepositTickets(&_EntryPoint.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_EntryPoint *EntryPointCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Deposit         *big.Int
		Staked          bool
		Stake           *big.Int
		UnstakeDelaySec uint32
		WithdrawTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Staked = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Stake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeDelaySec = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.WithdrawTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_EntryPoint *EntryPointSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _EntryPoint.Contract.Deposits(&_EntryPoint.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_EntryPoint *EntryPointCallerSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _EntryPoint.Contract.Deposits(&_EntryPoint.CallOpts, arg0)
}

// DstEids is a free data retrieval call binding the contract method 0x05dc0922.
//
// Solidity: function dstEids(uint256 ) view returns(uint32)
func (_EntryPoint *EntryPointCaller) DstEids(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "dstEids", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DstEids is a free data retrieval call binding the contract method 0x05dc0922.
//
// Solidity: function dstEids(uint256 ) view returns(uint32)
func (_EntryPoint *EntryPointSession) DstEids(arg0 *big.Int) (uint32, error) {
	return _EntryPoint.Contract.DstEids(&_EntryPoint.CallOpts, arg0)
}

// DstEids is a free data retrieval call binding the contract method 0x05dc0922.
//
// Solidity: function dstEids(uint256 ) view returns(uint32)
func (_EntryPoint *EntryPointCallerSession) DstEids(arg0 *big.Int) (uint32, error) {
	return _EntryPoint.Contract.DstEids(&_EntryPoint.CallOpts, arg0)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_EntryPoint *EntryPointCaller) GetDepositInfo(opts *bind.CallOpts, account common.Address) (IStakeManagerDepositInfo, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getDepositInfo", account)

	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)

	return out0, err

}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_EntryPoint *EntryPointSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _EntryPoint.Contract.GetDepositInfo(&_EntryPoint.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (_EntryPoint *EntryPointCallerSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _EntryPoint.Contract.GetDepositInfo(&_EntryPoint.CallOpts, account)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x3a106262.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,address) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointCaller) GetUserOpHash(opts *bind.CallOpts, userOp PackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0x3a106262.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,address) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _EntryPoint.Contract.GetUserOpHash(&_EntryPoint.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x3a106262.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,address) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointCallerSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _EntryPoint.Contract.GetUserOpHash(&_EntryPoint.CallOpts, userOp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EntryPoint *EntryPointCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EntryPoint *EntryPointSession) Owner() (common.Address, error) {
	return _EntryPoint.Contract.Owner(&_EntryPoint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EntryPoint *EntryPointCallerSession) Owner() (common.Address, error) {
	return _EntryPoint.Contract.Owner(&_EntryPoint.CallOpts)
}

// SmtRoot is a free data retrieval call binding the contract method 0x1f14e001.
//
// Solidity: function smtRoot() view returns(bytes32)
func (_EntryPoint *EntryPointCaller) SmtRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "smtRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SmtRoot is a free data retrieval call binding the contract method 0x1f14e001.
//
// Solidity: function smtRoot() view returns(bytes32)
func (_EntryPoint *EntryPointSession) SmtRoot() ([32]byte, error) {
	return _EntryPoint.Contract.SmtRoot(&_EntryPoint.CallOpts)
}

// SmtRoot is a free data retrieval call binding the contract method 0x1f14e001.
//
// Solidity: function smtRoot() view returns(bytes32)
func (_EntryPoint *EntryPointCallerSession) SmtRoot() ([32]byte, error) {
	return _EntryPoint.Contract.SmtRoot(&_EntryPoint.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EntryPoint *EntryPointCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EntryPoint *EntryPointSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EntryPoint.Contract.SupportsInterface(&_EntryPoint.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EntryPoint *EntryPointCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EntryPoint.Contract.SupportsInterface(&_EntryPoint.CallOpts, interfaceId)
}

// SyncRouter is a free data retrieval call binding the contract method 0x6cbe19af.
//
// Solidity: function syncRouter() view returns(address)
func (_EntryPoint *EntryPointCaller) SyncRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "syncRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SyncRouter is a free data retrieval call binding the contract method 0x6cbe19af.
//
// Solidity: function syncRouter() view returns(address)
func (_EntryPoint *EntryPointSession) SyncRouter() (common.Address, error) {
	return _EntryPoint.Contract.SyncRouter(&_EntryPoint.CallOpts)
}

// SyncRouter is a free data retrieval call binding the contract method 0x6cbe19af.
//
// Solidity: function syncRouter() view returns(address)
func (_EntryPoint *EntryPointCallerSession) SyncRouter() (common.Address, error) {
	return _EntryPoint.Contract.SyncRouter(&_EntryPoint.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_EntryPoint *EntryPointCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_EntryPoint *EntryPointSession) Verifier() (common.Address, error) {
	return _EntryPoint.Contract.Verifier(&_EntryPoint.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_EntryPoint *EntryPointCallerSession) Verifier() (common.Address, error) {
	return _EntryPoint.Contract.Verifier(&_EntryPoint.CallOpts)
}

// WithdrawTickets is a free data retrieval call binding the contract method 0xa6e79d9b.
//
// Solidity: function withdrawTickets(bytes32 ) view returns(bool)
func (_EntryPoint *EntryPointCaller) WithdrawTickets(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "withdrawTickets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawTickets is a free data retrieval call binding the contract method 0xa6e79d9b.
//
// Solidity: function withdrawTickets(bytes32 ) view returns(bool)
func (_EntryPoint *EntryPointSession) WithdrawTickets(arg0 [32]byte) (bool, error) {
	return _EntryPoint.Contract.WithdrawTickets(&_EntryPoint.CallOpts, arg0)
}

// WithdrawTickets is a free data retrieval call binding the contract method 0xa6e79d9b.
//
// Solidity: function withdrawTickets(bytes32 ) view returns(bool)
func (_EntryPoint *EntryPointCallerSession) WithdrawTickets(arg0 [32]byte) (bool, error) {
	return _EntryPoint.Contract.WithdrawTickets(&_EntryPoint.CallOpts, arg0)
}

// AddDepositTicket is a paid mutator transaction binding the contract method 0x74e81194.
//
// Solidity: function addDepositTicket(address accountAddress, address owner, uint256 amount) payable returns()
func (_EntryPoint *EntryPointTransactor) AddDepositTicket(opts *bind.TransactOpts, accountAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "addDepositTicket", accountAddress, owner, amount)
}

// AddDepositTicket is a paid mutator transaction binding the contract method 0x74e81194.
//
// Solidity: function addDepositTicket(address accountAddress, address owner, uint256 amount) payable returns()
func (_EntryPoint *EntryPointSession) AddDepositTicket(accountAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.AddDepositTicket(&_EntryPoint.TransactOpts, accountAddress, owner, amount)
}

// AddDepositTicket is a paid mutator transaction binding the contract method 0x74e81194.
//
// Solidity: function addDepositTicket(address accountAddress, address owner, uint256 amount) payable returns()
func (_EntryPoint *EntryPointTransactorSession) AddDepositTicket(accountAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.AddDepositTicket(&_EntryPoint.TransactOpts, accountAddress, owner, amount)
}

// AddWithdrawTicket is a paid mutator transaction binding the contract method 0xafc8d532.
//
// Solidity: function addWithdrawTicket(address user, uint256 amount) returns()
func (_EntryPoint *EntryPointTransactor) AddWithdrawTicket(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "addWithdrawTicket", user, amount)
}

// AddWithdrawTicket is a paid mutator transaction binding the contract method 0xafc8d532.
//
// Solidity: function addWithdrawTicket(address user, uint256 amount) returns()
func (_EntryPoint *EntryPointSession) AddWithdrawTicket(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.AddWithdrawTicket(&_EntryPoint.TransactOpts, user, amount)
}

// AddWithdrawTicket is a paid mutator transaction binding the contract method 0xafc8d532.
//
// Solidity: function addWithdrawTicket(address user, uint256 amount) returns()
func (_EntryPoint *EntryPointTransactorSession) AddWithdrawTicket(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.AddWithdrawTicket(&_EntryPoint.TransactOpts, user, amount)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_EntryPoint *EntryPointTransactor) DelegateAndRevert(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "delegateAndRevert", target, data)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_EntryPoint *EntryPointSession) DelegateAndRevert(target common.Address, data []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.DelegateAndRevert(&_EntryPoint.TransactOpts, target, data)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_EntryPoint *EntryPointTransactorSession) DelegateAndRevert(target common.Address, data []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.DelegateAndRevert(&_EntryPoint.TransactOpts, target, data)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_EntryPoint *EntryPointTransactor) GetSenderAddress(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "getSenderAddress", initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_EntryPoint *EntryPointSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.GetSenderAddress(&_EntryPoint.TransactOpts, initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_EntryPoint *EntryPointTransactorSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.GetSenderAddress(&_EntryPoint.TransactOpts, initCode)
}

// HandleOps is a paid mutator transaction binding the contract method 0xc9c8a7bd.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,address)[] ops, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactor) HandleOps(opts *bind.TransactOpts, ops []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0xc9c8a7bd.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,address)[] ops, address beneficiary) returns()
func (_EntryPoint *EntryPointSession) HandleOps(ops []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0xc9c8a7bd.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,address)[] ops, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactorSession) HandleOps(ops []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, beneficiary)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x40eef1d0.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256,address,address) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactor) InnerHandleOp(opts *bind.TransactOpts, callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "innerHandleOp", callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x40eef1d0.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256,address,address) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x40eef1d0.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256,address,address) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactorSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EntryPoint *EntryPointTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EntryPoint *EntryPointSession) RenounceOwnership() (*types.Transaction, error) {
	return _EntryPoint.Contract.RenounceOwnership(&_EntryPoint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EntryPoint *EntryPointTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EntryPoint.Contract.RenounceOwnership(&_EntryPoint.TransactOpts)
}

// SyncBatch is a paid mutator transaction binding the contract method 0xbc6be42e.
//
// Solidity: function syncBatch(bytes syncInfo) returns()
func (_EntryPoint *EntryPointTransactor) SyncBatch(opts *bind.TransactOpts, syncInfo []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "syncBatch", syncInfo)
}

// SyncBatch is a paid mutator transaction binding the contract method 0xbc6be42e.
//
// Solidity: function syncBatch(bytes syncInfo) returns()
func (_EntryPoint *EntryPointSession) SyncBatch(syncInfo []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.SyncBatch(&_EntryPoint.TransactOpts, syncInfo)
}

// SyncBatch is a paid mutator transaction binding the contract method 0xbc6be42e.
//
// Solidity: function syncBatch(bytes syncInfo) returns()
func (_EntryPoint *EntryPointTransactorSession) SyncBatch(syncInfo []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.SyncBatch(&_EntryPoint.TransactOpts, syncInfo)
}

// SyncBatchMock is a paid mutator transaction binding the contract method 0x400c1a7c.
//
// Solidity: function syncBatchMock(bytes syncInfo) returns()
func (_EntryPoint *EntryPointTransactor) SyncBatchMock(opts *bind.TransactOpts, syncInfo []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "syncBatchMock", syncInfo)
}

// SyncBatchMock is a paid mutator transaction binding the contract method 0x400c1a7c.
//
// Solidity: function syncBatchMock(bytes syncInfo) returns()
func (_EntryPoint *EntryPointSession) SyncBatchMock(syncInfo []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.SyncBatchMock(&_EntryPoint.TransactOpts, syncInfo)
}

// SyncBatchMock is a paid mutator transaction binding the contract method 0x400c1a7c.
//
// Solidity: function syncBatchMock(bytes syncInfo) returns()
func (_EntryPoint *EntryPointTransactorSession) SyncBatchMock(syncInfo []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.SyncBatchMock(&_EntryPoint.TransactOpts, syncInfo)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EntryPoint *EntryPointTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EntryPoint *EntryPointSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.TransferOwnership(&_EntryPoint.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EntryPoint *EntryPointTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.TransferOwnership(&_EntryPoint.TransactOpts, newOwner)
}

// UpdateDstEids is a paid mutator transaction binding the contract method 0x89b287f0.
//
// Solidity: function updateDstEids(uint32[] _dstEids) returns()
func (_EntryPoint *EntryPointTransactor) UpdateDstEids(opts *bind.TransactOpts, _dstEids []uint32) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "updateDstEids", _dstEids)
}

// UpdateDstEids is a paid mutator transaction binding the contract method 0x89b287f0.
//
// Solidity: function updateDstEids(uint32[] _dstEids) returns()
func (_EntryPoint *EntryPointSession) UpdateDstEids(_dstEids []uint32) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateDstEids(&_EntryPoint.TransactOpts, _dstEids)
}

// UpdateDstEids is a paid mutator transaction binding the contract method 0x89b287f0.
//
// Solidity: function updateDstEids(uint32[] _dstEids) returns()
func (_EntryPoint *EntryPointTransactorSession) UpdateDstEids(_dstEids []uint32) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateDstEids(&_EntryPoint.TransactOpts, _dstEids)
}

// UpdateSyncRouter is a paid mutator transaction binding the contract method 0x7fcb5227.
//
// Solidity: function updateSyncRouter(address _syncRouter) returns()
func (_EntryPoint *EntryPointTransactor) UpdateSyncRouter(opts *bind.TransactOpts, _syncRouter common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "updateSyncRouter", _syncRouter)
}

// UpdateSyncRouter is a paid mutator transaction binding the contract method 0x7fcb5227.
//
// Solidity: function updateSyncRouter(address _syncRouter) returns()
func (_EntryPoint *EntryPointSession) UpdateSyncRouter(_syncRouter common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateSyncRouter(&_EntryPoint.TransactOpts, _syncRouter)
}

// UpdateSyncRouter is a paid mutator transaction binding the contract method 0x7fcb5227.
//
// Solidity: function updateSyncRouter(address _syncRouter) returns()
func (_EntryPoint *EntryPointTransactorSession) UpdateSyncRouter(_syncRouter common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateSyncRouter(&_EntryPoint.TransactOpts, _syncRouter)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _verifier) returns()
func (_EntryPoint *EntryPointTransactor) UpdateVerifier(opts *bind.TransactOpts, _verifier common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "updateVerifier", _verifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _verifier) returns()
func (_EntryPoint *EntryPointSession) UpdateVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateVerifier(&_EntryPoint.TransactOpts, _verifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _verifier) returns()
func (_EntryPoint *EntryPointTransactorSession) UpdateVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateVerifier(&_EntryPoint.TransactOpts, _verifier)
}

// VerifyBatch is a paid mutator transaction binding the contract method 0xcd711765.
//
// Solidity: function verifyBatch(bytes proof, bytes publicValues, address beneficiary) payable returns()
func (_EntryPoint *EntryPointTransactor) VerifyBatch(opts *bind.TransactOpts, proof []byte, publicValues []byte, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "verifyBatch", proof, publicValues, beneficiary)
}

// VerifyBatch is a paid mutator transaction binding the contract method 0xcd711765.
//
// Solidity: function verifyBatch(bytes proof, bytes publicValues, address beneficiary) payable returns()
func (_EntryPoint *EntryPointSession) VerifyBatch(proof []byte, publicValues []byte, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatch(&_EntryPoint.TransactOpts, proof, publicValues, beneficiary)
}

// VerifyBatch is a paid mutator transaction binding the contract method 0xcd711765.
//
// Solidity: function verifyBatch(bytes proof, bytes publicValues, address beneficiary) payable returns()
func (_EntryPoint *EntryPointTransactorSession) VerifyBatch(proof []byte, publicValues []byte, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatch(&_EntryPoint.TransactOpts, proof, publicValues, beneficiary)
}

// VerifyBatchMock is a paid mutator transaction binding the contract method 0x2f6feea7.
//
// Solidity: function verifyBatchMock((address,uint256,uint256)[] depositTickets, (address,uint256,uint256)[] withdrawTickets, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactor) VerifyBatchMock(opts *bind.TransactOpts, depositTickets []ITicketManagerTicket, withdrawTickets []ITicketManagerTicket, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "verifyBatchMock", depositTickets, withdrawTickets, beneficiary)
}

// VerifyBatchMock is a paid mutator transaction binding the contract method 0x2f6feea7.
//
// Solidity: function verifyBatchMock((address,uint256,uint256)[] depositTickets, (address,uint256,uint256)[] withdrawTickets, address beneficiary) returns()
func (_EntryPoint *EntryPointSession) VerifyBatchMock(depositTickets []ITicketManagerTicket, withdrawTickets []ITicketManagerTicket, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatchMock(&_EntryPoint.TransactOpts, depositTickets, withdrawTickets, beneficiary)
}

// VerifyBatchMock is a paid mutator transaction binding the contract method 0x2f6feea7.
//
// Solidity: function verifyBatchMock((address,uint256,uint256)[] depositTickets, (address,uint256,uint256)[] withdrawTickets, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactorSession) VerifyBatchMock(depositTickets []ITicketManagerTicket, withdrawTickets []ITicketManagerTicket, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatchMock(&_EntryPoint.TransactOpts, depositTickets, withdrawTickets, beneficiary)
}

// VerifyBatchMockUserOp is a paid mutator transaction binding the contract method 0x611f11e9.
//
// Solidity: function verifyBatchMockUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,address)[] allUserOps, address beneficiary) payable returns()
func (_EntryPoint *EntryPointTransactor) VerifyBatchMockUserOp(opts *bind.TransactOpts, allUserOps []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "verifyBatchMockUserOp", allUserOps, beneficiary)
}

// VerifyBatchMockUserOp is a paid mutator transaction binding the contract method 0x611f11e9.
//
// Solidity: function verifyBatchMockUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,address)[] allUserOps, address beneficiary) payable returns()
func (_EntryPoint *EntryPointSession) VerifyBatchMockUserOp(allUserOps []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatchMockUserOp(&_EntryPoint.TransactOpts, allUserOps, beneficiary)
}

// VerifyBatchMockUserOp is a paid mutator transaction binding the contract method 0x611f11e9.
//
// Solidity: function verifyBatchMockUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,address)[] allUserOps, address beneficiary) payable returns()
func (_EntryPoint *EntryPointTransactorSession) VerifyBatchMockUserOp(allUserOps []PackedUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatchMockUserOp(&_EntryPoint.TransactOpts, allUserOps, beneficiary)
}

// EntryPointAccountDeployedIterator is returned from FilterAccountDeployed and is used to iterate over the raw logs and unpacked data for AccountDeployed events raised by the EntryPoint contract.
type EntryPointAccountDeployedIterator struct {
	Event *EntryPointAccountDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointAccountDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointAccountDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointAccountDeployed represents a AccountDeployed event raised by the EntryPoint contract.
type EntryPointAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDeployed is a free log retrieval operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_EntryPoint *EntryPointFilterer) FilterAccountDeployed(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointAccountDeployedIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointAccountDeployedIterator{contract: _EntryPoint.contract, event: "AccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAccountDeployed is a free log subscription operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_EntryPoint *EntryPointFilterer) WatchAccountDeployed(opts *bind.WatchOpts, sink chan<- *EntryPointAccountDeployed, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointAccountDeployed)
				if err := _EntryPoint.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountDeployed is a log parse operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_EntryPoint *EntryPointFilterer) ParseAccountDeployed(log types.Log) (*EntryPointAccountDeployed, error) {
	event := new(EntryPointAccountDeployed)
	if err := _EntryPoint.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointBeforeExecutionIterator is returned from FilterBeforeExecution and is used to iterate over the raw logs and unpacked data for BeforeExecution events raised by the EntryPoint contract.
type EntryPointBeforeExecutionIterator struct {
	Event *EntryPointBeforeExecution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointBeforeExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointBeforeExecution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointBeforeExecution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointBeforeExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointBeforeExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointBeforeExecution represents a BeforeExecution event raised by the EntryPoint contract.
type EntryPointBeforeExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBeforeExecution is a free log retrieval operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_EntryPoint *EntryPointFilterer) FilterBeforeExecution(opts *bind.FilterOpts) (*EntryPointBeforeExecutionIterator, error) {

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return &EntryPointBeforeExecutionIterator{contract: _EntryPoint.contract, event: "BeforeExecution", logs: logs, sub: sub}, nil
}

// WatchBeforeExecution is a free log subscription operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_EntryPoint *EntryPointFilterer) WatchBeforeExecution(opts *bind.WatchOpts, sink chan<- *EntryPointBeforeExecution) (event.Subscription, error) {

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointBeforeExecution)
				if err := _EntryPoint.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeforeExecution is a log parse operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_EntryPoint *EntryPointFilterer) ParseBeforeExecution(log types.Log) (*EntryPointBeforeExecution, error) {
	event := new(EntryPointBeforeExecution)
	if err := _EntryPoint.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointDepositTicketAddedIterator is returned from FilterDepositTicketAdded and is used to iterate over the raw logs and unpacked data for DepositTicketAdded events raised by the EntryPoint contract.
type EntryPointDepositTicketAddedIterator struct {
	Event *EntryPointDepositTicketAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointDepositTicketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointDepositTicketAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointDepositTicketAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointDepositTicketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointDepositTicketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointDepositTicketAdded represents a DepositTicketAdded event raised by the EntryPoint contract.
type EntryPointDepositTicketAdded struct {
	User       common.Address
	TicketHash [32]byte
	Amount     *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketAdded is a free log retrieval operation binding the contract event 0x2537ccb2b2fe11ea53e1afd15883c085006e0571b5b4f33dbf8becaf183f4e0b.
//
// Solidity: event DepositTicketAdded(address indexed user, bytes32 indexed ticketHash, uint256 amount, uint256 timestamp)
func (_EntryPoint *EntryPointFilterer) FilterDepositTicketAdded(opts *bind.FilterOpts, user []common.Address, ticketHash [][32]byte) (*EntryPointDepositTicketAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var ticketHashRule []interface{}
	for _, ticketHashItem := range ticketHash {
		ticketHashRule = append(ticketHashRule, ticketHashItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "DepositTicketAdded", userRule, ticketHashRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointDepositTicketAddedIterator{contract: _EntryPoint.contract, event: "DepositTicketAdded", logs: logs, sub: sub}, nil
}

// WatchDepositTicketAdded is a free log subscription operation binding the contract event 0x2537ccb2b2fe11ea53e1afd15883c085006e0571b5b4f33dbf8becaf183f4e0b.
//
// Solidity: event DepositTicketAdded(address indexed user, bytes32 indexed ticketHash, uint256 amount, uint256 timestamp)
func (_EntryPoint *EntryPointFilterer) WatchDepositTicketAdded(opts *bind.WatchOpts, sink chan<- *EntryPointDepositTicketAdded, user []common.Address, ticketHash [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var ticketHashRule []interface{}
	for _, ticketHashItem := range ticketHash {
		ticketHashRule = append(ticketHashRule, ticketHashItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "DepositTicketAdded", userRule, ticketHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointDepositTicketAdded)
				if err := _EntryPoint.contract.UnpackLog(event, "DepositTicketAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositTicketAdded is a log parse operation binding the contract event 0x2537ccb2b2fe11ea53e1afd15883c085006e0571b5b4f33dbf8becaf183f4e0b.
//
// Solidity: event DepositTicketAdded(address indexed user, bytes32 indexed ticketHash, uint256 amount, uint256 timestamp)
func (_EntryPoint *EntryPointFilterer) ParseDepositTicketAdded(log types.Log) (*EntryPointDepositTicketAdded, error) {
	event := new(EntryPointDepositTicketAdded)
	if err := _EntryPoint.contract.UnpackLog(event, "DepositTicketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointDepositTicketDeletedIterator is returned from FilterDepositTicketDeleted and is used to iterate over the raw logs and unpacked data for DepositTicketDeleted events raised by the EntryPoint contract.
type EntryPointDepositTicketDeletedIterator struct {
	Event *EntryPointDepositTicketDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointDepositTicketDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointDepositTicketDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointDepositTicketDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointDepositTicketDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointDepositTicketDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointDepositTicketDeleted represents a DepositTicketDeleted event raised by the EntryPoint contract.
type EntryPointDepositTicketDeleted struct {
	User       common.Address
	Amount     *big.Int
	TicketHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketDeleted is a free log retrieval operation binding the contract event 0x03b8ee44084307fb90dbd84aef4a2bbaef367bb8a49bfbfc6e930a12f2f749ea.
//
// Solidity: event DepositTicketDeleted(address indexed user, uint256 amount, bytes32 ticketHash)
func (_EntryPoint *EntryPointFilterer) FilterDepositTicketDeleted(opts *bind.FilterOpts, user []common.Address) (*EntryPointDepositTicketDeletedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "DepositTicketDeleted", userRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointDepositTicketDeletedIterator{contract: _EntryPoint.contract, event: "DepositTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchDepositTicketDeleted is a free log subscription operation binding the contract event 0x03b8ee44084307fb90dbd84aef4a2bbaef367bb8a49bfbfc6e930a12f2f749ea.
//
// Solidity: event DepositTicketDeleted(address indexed user, uint256 amount, bytes32 ticketHash)
func (_EntryPoint *EntryPointFilterer) WatchDepositTicketDeleted(opts *bind.WatchOpts, sink chan<- *EntryPointDepositTicketDeleted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "DepositTicketDeleted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointDepositTicketDeleted)
				if err := _EntryPoint.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositTicketDeleted is a log parse operation binding the contract event 0x03b8ee44084307fb90dbd84aef4a2bbaef367bb8a49bfbfc6e930a12f2f749ea.
//
// Solidity: event DepositTicketDeleted(address indexed user, uint256 amount, bytes32 ticketHash)
func (_EntryPoint *EntryPointFilterer) ParseDepositTicketDeleted(log types.Log) (*EntryPointDepositTicketDeleted, error) {
	event := new(EntryPointDepositTicketDeleted)
	if err := _EntryPoint.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the EntryPoint contract.
type EntryPointDepositedIterator struct {
	Event *EntryPointDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointDeposited represents a Deposited event raised by the EntryPoint contract.
type EntryPointDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_EntryPoint *EntryPointFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*EntryPointDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointDepositedIterator{contract: _EntryPoint.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_EntryPoint *EntryPointFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *EntryPointDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointDeposited)
				if err := _EntryPoint.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_EntryPoint *EntryPointFilterer) ParseDeposited(log types.Log) (*EntryPointDeposited, error) {
	event := new(EntryPointDeposited)
	if err := _EntryPoint.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EntryPoint contract.
type EntryPointOwnershipTransferredIterator struct {
	Event *EntryPointOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointOwnershipTransferred represents a OwnershipTransferred event raised by the EntryPoint contract.
type EntryPointOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EntryPoint *EntryPointFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EntryPointOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointOwnershipTransferredIterator{contract: _EntryPoint.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EntryPoint *EntryPointFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EntryPointOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointOwnershipTransferred)
				if err := _EntryPoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EntryPoint *EntryPointFilterer) ParseOwnershipTransferred(log types.Log) (*EntryPointOwnershipTransferred, error) {
	event := new(EntryPointOwnershipTransferred)
	if err := _EntryPoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointPostOpRevertReasonIterator is returned from FilterPostOpRevertReason and is used to iterate over the raw logs and unpacked data for PostOpRevertReason events raised by the EntryPoint contract.
type EntryPointPostOpRevertReasonIterator struct {
	Event *EntryPointPostOpRevertReason // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointPostOpRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointPostOpRevertReason)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointPostOpRevertReason)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointPostOpRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointPostOpRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointPostOpRevertReason represents a PostOpRevertReason event raised by the EntryPoint contract.
type EntryPointPostOpRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostOpRevertReason is a free log retrieval operation binding the contract event 0x676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_EntryPoint *EntryPointFilterer) FilterPostOpRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointPostOpRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "PostOpRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointPostOpRevertReasonIterator{contract: _EntryPoint.contract, event: "PostOpRevertReason", logs: logs, sub: sub}, nil
}

// WatchPostOpRevertReason is a free log subscription operation binding the contract event 0x676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_EntryPoint *EntryPointFilterer) WatchPostOpRevertReason(opts *bind.WatchOpts, sink chan<- *EntryPointPostOpRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "PostOpRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointPostOpRevertReason)
				if err := _EntryPoint.contract.UnpackLog(event, "PostOpRevertReason", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePostOpRevertReason is a log parse operation binding the contract event 0x676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_EntryPoint *EntryPointFilterer) ParsePostOpRevertReason(log types.Log) (*EntryPointPostOpRevertReason, error) {
	event := new(EntryPointPostOpRevertReason)
	if err := _EntryPoint.contract.UnpackLog(event, "PostOpRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSignatureAggregatorChangedIterator is returned from FilterSignatureAggregatorChanged and is used to iterate over the raw logs and unpacked data for SignatureAggregatorChanged events raised by the EntryPoint contract.
type EntryPointSignatureAggregatorChangedIterator struct {
	Event *EntryPointSignatureAggregatorChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointSignatureAggregatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSignatureAggregatorChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointSignatureAggregatorChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointSignatureAggregatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSignatureAggregatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSignatureAggregatorChanged represents a SignatureAggregatorChanged event raised by the EntryPoint contract.
type EntryPointSignatureAggregatorChanged struct {
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSignatureAggregatorChanged is a free log retrieval operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_EntryPoint *EntryPointFilterer) FilterSignatureAggregatorChanged(opts *bind.FilterOpts, aggregator []common.Address) (*EntryPointSignatureAggregatorChangedIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSignatureAggregatorChangedIterator{contract: _EntryPoint.contract, event: "SignatureAggregatorChanged", logs: logs, sub: sub}, nil
}

// WatchSignatureAggregatorChanged is a free log subscription operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_EntryPoint *EntryPointFilterer) WatchSignatureAggregatorChanged(opts *bind.WatchOpts, sink chan<- *EntryPointSignatureAggregatorChanged, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSignatureAggregatorChanged)
				if err := _EntryPoint.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignatureAggregatorChanged is a log parse operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_EntryPoint *EntryPointFilterer) ParseSignatureAggregatorChanged(log types.Log) (*EntryPointSignatureAggregatorChanged, error) {
	event := new(EntryPointSignatureAggregatorChanged)
	if err := _EntryPoint.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointUserOperationEventIterator is returned from FilterUserOperationEvent and is used to iterate over the raw logs and unpacked data for UserOperationEvent events raised by the EntryPoint contract.
type EntryPointUserOperationEventIterator struct {
	Event *EntryPointUserOperationEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointUserOperationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointUserOperationEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointUserOperationEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointUserOperationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointUserOperationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointUserOperationEvent represents a UserOperationEvent event raised by the EntryPoint contract.
type EntryPointUserOperationEvent struct {
	UserOpHash    [32]byte
	Sender        common.Address
	Paymaster     common.Address
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserOperationEvent is a free log retrieval operation binding the contract event 0x58eb13224f9e7bb63720c3b18c69c98cfe7fb1459db62324c36836a7d30146f3.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_EntryPoint *EntryPointFilterer) FilterUserOperationEvent(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (*EntryPointUserOperationEventIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointUserOperationEventIterator{contract: _EntryPoint.contract, event: "UserOperationEvent", logs: logs, sub: sub}, nil
}

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0x58eb13224f9e7bb63720c3b18c69c98cfe7fb1459db62324c36836a7d30146f3.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_EntryPoint *EntryPointFilterer) WatchUserOperationEvent(opts *bind.WatchOpts, sink chan<- *EntryPointUserOperationEvent, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointUserOperationEvent)
				if err := _EntryPoint.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserOperationEvent is a log parse operation binding the contract event 0x58eb13224f9e7bb63720c3b18c69c98cfe7fb1459db62324c36836a7d30146f3.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_EntryPoint *EntryPointFilterer) ParseUserOperationEvent(log types.Log) (*EntryPointUserOperationEvent, error) {
	event := new(EntryPointUserOperationEvent)
	if err := _EntryPoint.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointUserOperationPrefundTooLowIterator is returned from FilterUserOperationPrefundTooLow and is used to iterate over the raw logs and unpacked data for UserOperationPrefundTooLow events raised by the EntryPoint contract.
type EntryPointUserOperationPrefundTooLowIterator struct {
	Event *EntryPointUserOperationPrefundTooLow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointUserOperationPrefundTooLowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointUserOperationPrefundTooLow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointUserOperationPrefundTooLow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointUserOperationPrefundTooLowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointUserOperationPrefundTooLowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointUserOperationPrefundTooLow represents a UserOperationPrefundTooLow event raised by the EntryPoint contract.
type EntryPointUserOperationPrefundTooLow struct {
	UserOpHash [32]byte
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserOperationPrefundTooLow is a free log retrieval operation binding the contract event 0x1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff28.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender)
func (_EntryPoint *EntryPointFilterer) FilterUserOperationPrefundTooLow(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointUserOperationPrefundTooLowIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "UserOperationPrefundTooLow", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointUserOperationPrefundTooLowIterator{contract: _EntryPoint.contract, event: "UserOperationPrefundTooLow", logs: logs, sub: sub}, nil
}

// WatchUserOperationPrefundTooLow is a free log subscription operation binding the contract event 0x1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff28.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender)
func (_EntryPoint *EntryPointFilterer) WatchUserOperationPrefundTooLow(opts *bind.WatchOpts, sink chan<- *EntryPointUserOperationPrefundTooLow, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "UserOperationPrefundTooLow", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointUserOperationPrefundTooLow)
				if err := _EntryPoint.contract.UnpackLog(event, "UserOperationPrefundTooLow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserOperationPrefundTooLow is a log parse operation binding the contract event 0x1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff28.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender)
func (_EntryPoint *EntryPointFilterer) ParseUserOperationPrefundTooLow(log types.Log) (*EntryPointUserOperationPrefundTooLow, error) {
	event := new(EntryPointUserOperationPrefundTooLow)
	if err := _EntryPoint.contract.UnpackLog(event, "UserOperationPrefundTooLow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointUserOperationRevertReasonIterator is returned from FilterUserOperationRevertReason and is used to iterate over the raw logs and unpacked data for UserOperationRevertReason events raised by the EntryPoint contract.
type EntryPointUserOperationRevertReasonIterator struct {
	Event *EntryPointUserOperationRevertReason // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointUserOperationRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointUserOperationRevertReason)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointUserOperationRevertReason)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointUserOperationRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointUserOperationRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointUserOperationRevertReason represents a UserOperationRevertReason event raised by the EntryPoint contract.
type EntryPointUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserOperationRevertReason is a free log retrieval operation binding the contract event 0x0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_EntryPoint *EntryPointFilterer) FilterUserOperationRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointUserOperationRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointUserOperationRevertReasonIterator{contract: _EntryPoint.contract, event: "UserOperationRevertReason", logs: logs, sub: sub}, nil
}

// WatchUserOperationRevertReason is a free log subscription operation binding the contract event 0x0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_EntryPoint *EntryPointFilterer) WatchUserOperationRevertReason(opts *bind.WatchOpts, sink chan<- *EntryPointUserOperationRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointUserOperationRevertReason)
				if err := _EntryPoint.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserOperationRevertReason is a log parse operation binding the contract event 0x0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_EntryPoint *EntryPointFilterer) ParseUserOperationRevertReason(log types.Log) (*EntryPointUserOperationRevertReason, error) {
	event := new(EntryPointUserOperationRevertReason)
	if err := _EntryPoint.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointWithdrawTicketAddedIterator is returned from FilterWithdrawTicketAdded and is used to iterate over the raw logs and unpacked data for WithdrawTicketAdded events raised by the EntryPoint contract.
type EntryPointWithdrawTicketAddedIterator struct {
	Event *EntryPointWithdrawTicketAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointWithdrawTicketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointWithdrawTicketAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointWithdrawTicketAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointWithdrawTicketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointWithdrawTicketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointWithdrawTicketAdded represents a WithdrawTicketAdded event raised by the EntryPoint contract.
type EntryPointWithdrawTicketAdded struct {
	User       common.Address
	TicketHash [32]byte
	Amount     *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketAdded is a free log retrieval operation binding the contract event 0xf2d23449d3c6fc1266c59d62d6cb28ebcd25d872fc4c70d906c571e7247a8b00.
//
// Solidity: event WithdrawTicketAdded(address indexed user, bytes32 indexed ticketHash, uint256 amount, uint256 timestamp)
func (_EntryPoint *EntryPointFilterer) FilterWithdrawTicketAdded(opts *bind.FilterOpts, user []common.Address, ticketHash [][32]byte) (*EntryPointWithdrawTicketAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var ticketHashRule []interface{}
	for _, ticketHashItem := range ticketHash {
		ticketHashRule = append(ticketHashRule, ticketHashItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "WithdrawTicketAdded", userRule, ticketHashRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointWithdrawTicketAddedIterator{contract: _EntryPoint.contract, event: "WithdrawTicketAdded", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketAdded is a free log subscription operation binding the contract event 0xf2d23449d3c6fc1266c59d62d6cb28ebcd25d872fc4c70d906c571e7247a8b00.
//
// Solidity: event WithdrawTicketAdded(address indexed user, bytes32 indexed ticketHash, uint256 amount, uint256 timestamp)
func (_EntryPoint *EntryPointFilterer) WatchWithdrawTicketAdded(opts *bind.WatchOpts, sink chan<- *EntryPointWithdrawTicketAdded, user []common.Address, ticketHash [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var ticketHashRule []interface{}
	for _, ticketHashItem := range ticketHash {
		ticketHashRule = append(ticketHashRule, ticketHashItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "WithdrawTicketAdded", userRule, ticketHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointWithdrawTicketAdded)
				if err := _EntryPoint.contract.UnpackLog(event, "WithdrawTicketAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawTicketAdded is a log parse operation binding the contract event 0xf2d23449d3c6fc1266c59d62d6cb28ebcd25d872fc4c70d906c571e7247a8b00.
//
// Solidity: event WithdrawTicketAdded(address indexed user, bytes32 indexed ticketHash, uint256 amount, uint256 timestamp)
func (_EntryPoint *EntryPointFilterer) ParseWithdrawTicketAdded(log types.Log) (*EntryPointWithdrawTicketAdded, error) {
	event := new(EntryPointWithdrawTicketAdded)
	if err := _EntryPoint.contract.UnpackLog(event, "WithdrawTicketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointWithdrawTicketDeletedIterator is returned from FilterWithdrawTicketDeleted and is used to iterate over the raw logs and unpacked data for WithdrawTicketDeleted events raised by the EntryPoint contract.
type EntryPointWithdrawTicketDeletedIterator struct {
	Event *EntryPointWithdrawTicketDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointWithdrawTicketDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointWithdrawTicketDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointWithdrawTicketDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointWithdrawTicketDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointWithdrawTicketDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointWithdrawTicketDeleted represents a WithdrawTicketDeleted event raised by the EntryPoint contract.
type EntryPointWithdrawTicketDeleted struct {
	User       common.Address
	Amount     *big.Int
	TicketHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketDeleted is a free log retrieval operation binding the contract event 0xf89de5c5fbaa24098f4988752b4e8c7b8f11fc756a1688b549173fcb9396ea19.
//
// Solidity: event WithdrawTicketDeleted(address indexed user, uint256 amount, bytes32 ticketHash)
func (_EntryPoint *EntryPointFilterer) FilterWithdrawTicketDeleted(opts *bind.FilterOpts, user []common.Address) (*EntryPointWithdrawTicketDeletedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "WithdrawTicketDeleted", userRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointWithdrawTicketDeletedIterator{contract: _EntryPoint.contract, event: "WithdrawTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketDeleted is a free log subscription operation binding the contract event 0xf89de5c5fbaa24098f4988752b4e8c7b8f11fc756a1688b549173fcb9396ea19.
//
// Solidity: event WithdrawTicketDeleted(address indexed user, uint256 amount, bytes32 ticketHash)
func (_EntryPoint *EntryPointFilterer) WatchWithdrawTicketDeleted(opts *bind.WatchOpts, sink chan<- *EntryPointWithdrawTicketDeleted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "WithdrawTicketDeleted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointWithdrawTicketDeleted)
				if err := _EntryPoint.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawTicketDeleted is a log parse operation binding the contract event 0xf89de5c5fbaa24098f4988752b4e8c7b8f11fc756a1688b549173fcb9396ea19.
//
// Solidity: event WithdrawTicketDeleted(address indexed user, uint256 amount, bytes32 ticketHash)
func (_EntryPoint *EntryPointFilterer) ParseWithdrawTicketDeleted(log types.Log) (*EntryPointWithdrawTicketDeleted, error) {
	event := new(EntryPointWithdrawTicketDeleted)
	if err := _EntryPoint.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the EntryPoint contract.
type EntryPointWithdrawnIterator struct {
	Event *EntryPointWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EntryPointWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EntryPointWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EntryPointWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointWithdrawn represents a Withdrawn event raised by the EntryPoint contract.
type EntryPointWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*EntryPointWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointWithdrawnIterator{contract: _EntryPoint.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *EntryPointWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointWithdrawn)
				if err := _EntryPoint.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) ParseWithdrawn(log types.Log) (*EntryPointWithdrawn, error) {
	event := new(EntryPointWithdrawn)
	if err := _EntryPoint.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
