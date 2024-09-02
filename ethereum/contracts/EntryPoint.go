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
	Nonce                         *big.Int
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
}

// IEntryPointUserOpsPerAggregator is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointUserOpsPerAggregator struct {
	UserOps    []PackedUserOperation
	Aggregator common.Address
	Signature  []byte
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
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// EntryPointMetaData contains all meta data concerning the EntryPoint contract.
var EntryPointMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addDepositTicket\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addStake\",\"inputs\":[{\"name\":\"unstakeDelaySec\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addWithdrawTicket\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegateAndRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositTickets\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staked\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"stake\",\"type\":\"uint112\",\"internalType\":\"uint112\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"withdrawTime\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositInfo\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structIStakeManager.DepositInfo\",\"components\":[{\"name\":\"deposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staked\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"stake\",\"type\":\"uint112\",\"internalType\":\"uint112\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"withdrawTime\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSenderAddress\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getUserOpHash\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleAggregatedOps\",\"inputs\":[{\"name\":\"opsPerAggregator\",\"type\":\"tuple[]\",\"internalType\":\"structIEntryPoint.UserOpsPerAggregator[]\",\"components\":[{\"name\":\"userOps\",\"type\":\"tuple[]\",\"internalType\":\"structPackedUserOperation[]\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"contractIAggregator\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleOps\",\"inputs\":[{\"name\":\"ops\",\"type\":\"tuple[]\",\"internalType\":\"structPackedUserOperation[]\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpsAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"incrementNonce\",\"inputs\":[{\"name\":\"key\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"innerHandleOp\",\"inputs\":[{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"opInfo\",\"type\":\"tuple\",\"internalType\":\"structEntryPoint.UserOpInfo\",\"components\":[{\"name\":\"mUserOp\",\"type\":\"tuple\",\"internalType\":\"structEntryPoint.MemoryUserOp\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterVerificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterPostOpGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymaster\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prefund\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preOpGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"actualGasCost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonceSequenceNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"smtRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawStake\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTickets\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zkAAEntryPoint\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubInput\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newSmtRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"userOps\",\"type\":\"tuple[]\",\"internalType\":\"structPackedUserOperation[]\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpsAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"depositTickets\",\"type\":\"tuple[]\",\"internalType\":\"structITicketManager.Ticket[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"withdrawTickets\",\"type\":\"tuple[]\",\"internalType\":\"structITicketManager.Ticket[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AccountDeployed\",\"inputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"factory\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"paymaster\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeforeExecution\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositTicketAdded\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ticketHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositTicketDeleted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ticketHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"totalDeposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PostOpRevertReason\",\"inputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertReason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignatureAggregatorChanged\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeLocked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"totalStaked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"unstakeDelaySec\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeUnlocked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeWithdrawn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserOperationEvent\",\"inputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymaster\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"actualGasCost\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"actualGasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserOperationPrefundTooLow\",\"inputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserOperationRevertReason\",\"inputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"revertReason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawTicketAdded\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ticketHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawTicketDeleted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ticketHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DelegateAndRevert\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"ret\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"FailedOp\",\"inputs\":[{\"name\":\"opIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"FailedOpWithRevert\",\"inputs\":[{\"name\":\"opIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"inner\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"PostOpReverted\",\"inputs\":[{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderAddressResult\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SignatureValidationFailed\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a06040523462000036576200001462000059565b6040516157d0620001118239608051818181611b24015261539d01526157d090f35b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b6040513d5f823e3d90fd5b62000063620000ad565b60405161039b810167ffffffffffffffff811182821017620000a75762000091829161039b620058e1843990565b03905ff08015620000a157608052565b6200004e565b6200003a565b620000b7620000b9565b565b620000b7620000fa565b620000d3620000d3620000d39290565b90565b620000d36001620000c3565b90620000d3620000d3620000f692620000c3565b9055565b620000b762000108620000d6565b6005620000e256fe60806040526004361015610011575f80fd5b5f3560e01c806301ffc9a7146101705780630396cb601461016b5780630bd28e3b146101665780631b2e01b8146101615780631f14e0011461015c57806322cdde4c1461015757806335567e1a146101525780635287ce121461014d57806368e9b1ea1461014857806370a0823114610143578063850aaf621461013e5780639b249f6914610139578063a6e79d9b14610134578063afc8d5321461012f578063b2dd2ba31461012a578063bb9fe6bf14610125578063c23a5cea14610120578063cd9ae1711461011b578063d384c38214610116578063dbed18e014610111578063e997c9f91461010c5763fc7e286d036101a057610d54565b610c07565b610adf565b610a8e565b610a69565b610a51565b610a39565b610a1a565b610740565b610702565b61068f565b610653565b6105ce565b6105b2565b610503565b61047b565b610460565b610410565b6103cf565b610285565b61022c565b6101ce565b7fffffffff0000000000000000000000000000000000000000000000000000000081165b036101a057565b5f80fd5b905035906101b182610175565b565b906020828203126101a0576101c7916101a4565b90565b9052565b346101a0576101fb6101e96101e43660046101b3565b610d85565b60405191829182901515815260200190565b0390f35b63ffffffff8116610199565b905035906101b1826101ff565b906020828203126101a0576101c79161020b565b61023f61023a366004610218565b6112ec565b604051005b77ffffffffffffffffffffffffffffffffffffffffffffffff8116610199565b905035906101b182610244565b906020828203126101a0576101c791610264565b346101a05761023f610298366004610271565b611495565b73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff8116610199565b905035906101b1826102b6565b91906040838203126101a0576101c7906102f981856102d2565b93602001610264565b61029d6101c76101c79273ffffffffffffffffffffffffffffffffffffffff1690565b6101c790610302565b6101c790610325565b906103419061032e565b5f5260205260405f2090565b6103746101c76101c79277ffffffffffffffffffffffffffffffffffffffffffffffff1690565b77ffffffffffffffffffffffffffffffffffffffffffffffff1690565b906103419061034d565b6101c7916008021c81565b906101c7915461039b565b6103ca6101c7926103c56002935f94610337565b610391565b6103a6565b346101a0576101fb6103eb6103e53660046102df565b906103b1565b6040515b9182918290815260200190565b5f9103126101a057565b6101c75f806103a6565b346101a0576104203660046103fc565b6101fb6103eb610406565b90816101209103126101a05790565b906020828203126101a057813567ffffffffffffffff81116101a0576101c7920161042b565b346101a0576101fb6103eb61047636600461043a565b6114ed565b346101a0576101fb6103eb6104913660046102df565b9061158e565b906020828203126101a0576101c7916102d2565b60a0810192916101b19190805182526020808201511515908301526040808201516dffffffffffffffffffffffffffff169083015260608082015163ffffffff169083015260809081015165ffffffffffff16910152565b346101a0576101fb61051e610519366004610497565b611683565b604051918291826104ab565b909182601f830112156101a05781359167ffffffffffffffff83116101a05760200192602083028401116101a057565b6060818303126101a057803567ffffffffffffffff81116101a0578261058191830161052a565b929093602083013567ffffffffffffffff81116101a0576105a7836101c792860161052a565b9390946040016102d2565b346101a05761023f6105c536600461055a565b9392909261196d565b346101a0576101fb6103eb6105e4366004610497565b61197a565b909182601f830112156101a05781359167ffffffffffffffff83116101a05760200192600183028401116101a057565b9190916040818403126101a05761063083826102d2565b92602082013567ffffffffffffffff81116101a05761064f92016105e9565b9091565b346101a057610663366004610619565b91611a5d565b906020828203126101a057813567ffffffffffffffff81116101a05761064f92016105e9565b346101a05761069f366004610669565b90611b1b565b80610199565b905035906101b1826106a5565b906020828203126101a0576101c7916106ab565b90610341565b6101c7916008021c5b60ff1690565b906101c791546106d2565b6101c7906106fd6004915f926106cc565b6106e1565b346101a0576101fb6101e96107183660046106b8565b6106ec565b91906040838203126101a0576101c79061073781856102d2565b936020016106ab565b346101a05761023f61075336600461071d565b90611cac565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b90601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810190811067ffffffffffffffff8211176107c657604052565b610759565b906101b16107d860405190565b9283610786565b67ffffffffffffffff81116107c657602090601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160190565b0190565b90825f939282370152565b9092919261083e610839826107df565b6107cb565b938185526020850190828401116101a0576101b19261081e565b9080601f830112156101a0578160206101c793359101610829565b9190610140838203126101a0576109379061088f6101406107cb565b9361089a82826102d2565b85526108a982602083016106ab565b60208601526108bb82604083016106ab565b60408601526108cd82606083016106ab565b60608601526108df82608083016106ab565b60808601526108f18260a083016106ab565b60a08601526109038260c083016106ab565b60c08601526109158260e083016102d2565b60e08601526109288261010083016106ab565b610100860152610120016106ab565b610120830152565b91906101e0838203126101a0576109bc9061095a60c06107cb565b936109658282610873565b85526109758261014083016106ab565b60208601526109888261016083016106ab565b604086015261099b8261018083016106ab565b60608601526109ae826101a083016106ab565b60808601526101c0016102d2565b60a0830152565b91610220838303126101a057823567ffffffffffffffff81116101a057826109ec918501610858565b926109fa836020830161093f565b9261020082013567ffffffffffffffff81116101a05761064f92016105e9565b346101a0576101fb6103eb610a303660046109c3565b92919091611e62565b346101a057610a493660046103fc565b61023f612187565b346101a05761023f610a64366004610497565b612407565b61023f610a7736600461071d565b9061259a565b6101c7906106fd6003915f926106cc565b346101a0576101fb6101e9610aa43660046106b8565b610a7d565b916040838303126101a057823567ffffffffffffffff81116101a057610ad4836101c792860161052a565b9390946020016102d2565b346101a05761023f610af2366004610aa9565b91612de6565b909182601f830112156101a05781359167ffffffffffffffff83116101a05760200192606083028401116101a057565b90610100828203126101a057813567ffffffffffffffff81116101a05781610b519184016105e9565b929093610b6183602084016106ab565b92610b6f81604085016106ab565b92606081013567ffffffffffffffff81116101a05782610b9091830161052a565b929093608083013567ffffffffffffffff81116101a05782610bb391850161052a565b92909360a081013567ffffffffffffffff81116101a05782610bd6918301610af8565b92909360c083013567ffffffffffffffff81116101a057610bfc836101c7928601610af8565b93909460e0016102d2565b346101a05761023f610c1a366004610b28565b9b9a909a999199989298979397969496612e5d565b6101c79081565b6101c79054610c2f565b6101c7906106db565b6101c79054610c40565b6101c79060081c5b6dffffffffffffffffffffffffffff1690565b6101c79054610c53565b6101c79060781c5b63ffffffff1690565b6101c79054610c78565b6101c79060981c5b65ffffffffffff1690565b6101c79054610c93565b610cbb906001610337565b610cc481610c36565b91610cd160018301610c49565b91610cde60018201610c6e565b916101c76001610cef818501610c89565b9301610ca6565b909594926101b194610d39610d4692610d22608096610d1960a088019c5f890152565b15156020870152565b6dffffffffffffffffffffffffffff166040850152565b63ffffffff166060830152565b019065ffffffffffff169052565b346101a0576101fb610d6f610d6a366004610497565b610cb0565b91610d7c95939560405190565b95869586610cf6565b80610dd17fe9778a7a000000000000000000000000000000000000000000000000000000005b917fffffffff000000000000000000000000000000000000000000000000000000001690565b148015610e5e575b8015610e2d575b8015610dfc575b908115610df2575090565b6101c79150612ea5565b5080610e277f3e84f02100000000000000000000000000000000000000000000000000000000610dab565b14610de7565b5080610e587f58143d1600000000000000000000000000000000000000000000000000000000610dab565b14610de0565b5080610e897f8fe7474d00000000000000000000000000000000000000000000000000000000610dab565b14610dd9565b610c806101c76101c79290565b15610ea357565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6d757374207370656369667920756e7374616b652064656c61790000000000006044820152606490fd5b0390fd5b15610f0d57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f63616e6e6f7420646563726561736520756e7374616b652074696d65000000006044820152606490fd5b6101c76101c76101c7926dffffffffffffffffffffffffffff1690565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b9190610fc1565b9290565b8201809211610fcc57565b610f89565b6101c76101c76101c79290565b15610fe557565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6e6f207374616b652073706563696669656400000000000000000000000000006044820152606490fd5b1561104b57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f7374616b65206f766572666c6f770000000000000000000000000000000000006044820152606490fd5b610c5b6101c76101c79290565b6101c760a06107cb565b610c9b6101c76101c79290565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff905b9181191691161790565b9061110c6101c761111392610fd1565b82546110ce565b9055565b9060ff906110f2565b906111306101c761111392151590565b8254611117565b906effffffffffffffffffffffffffff009060081b6110f2565b610c5b6101c76101c7926dffffffffffffffffffffffffffff1690565b9061117e6101c761111392611151565b8254611137565b9072ffffffff0000000000000000000000000000009060781b6110f2565b610c806101c76101c79263ffffffff1690565b906111c66101c7611113926111a3565b8254611185565b9078ffffffffffff000000000000000000000000000000000000009060981b6110f2565b610c9b6101c76101c79265ffffffffffff1690565b906112166101c7611113926111f1565b82546111cd565b6101b191906112aa9060809060019061123e81611238875190565b906110fc565b6112578282016112516020880151151590565b90611120565b61127e82820161127860408801516dffffffffffffffffffffffffffff1690565b9061116e565b61129b828201611295606088015163ffffffff1690565b906111b6565b0192015165ffffffffffff1690565b90611206565b906101b19161121d565b6101c76101c76101c79263ffffffff1690565b6101ca906112ba565b9081526040810192916101b191602001906112cd565b6112fa6101c7336001610337565b61131a8261131461130a5f610e8f565b9163ffffffff1690565b11610e9c565b6113378261133061130a610c8060018601610c89565b1015610f06565b6114216113a25f61135c61134d60018601610c6e565b6113573491610f6c565b610fb6565b936113778561137161136d85610fd1565b9190565b11610fde565b61139c8561139561136d6dffffffffffffffffffffffffffff610f6c565b1115611044565b01610c36565b61141161140260016113fd6113b6876110aa565b6113f089916113d95f956113d06113cb6110b7565b9a8b52565b151560208a0152565b6dffffffffffffffffffffffffffff166040880152565b63ffffffff166060860152565b6110c1565b65ffffffffffff166080830152565b61141c336001610337565b6112b0565b907fa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c019061144d3361032e565b9261146361145a60405190565b928392836112d6565b0390a2565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610fcc5760010190565b6114a86101b1916103c560023390610337565b6112386114b482610c36565b611468565b90815273ffffffffffffffffffffffffffffffffffffffff90911660208201526060810192916101b19160400152565b0152565b6114f690612ed2565b61152b6115023061032e565b9161151c4661151060405190565b948593602085016114b9565b60208201810382520382610786565b61153d611536825190565b9160200190565b2090565b6101c76101c76101c79277ffffffffffffffffffffffffffffffffffffffffffffffff1690565b6106db6101c76101c79290565b6101c79061158961136d6101c79460ff1690565b901b90565b6115bf6115b96115b46115ad6115cf946115a55f90565b506002610337565b8590610391565b610c36565b92611541565b6115c96040611568565b90611575565b1790565b6115db6110b7565b90602080808080865f5b8152015f8152015f5b8152015f8152015f905250565b6101c76115d3565b906101b161166b60016116146110b7565b9461162561162182610c36565b8752565b61163c611633838301610c49565b15156020880152565b61164a6113d9838301610c6e565b611665611658838301610c89565b63ffffffff166060880152565b01610ca6565b65ffffffffffff166080840152565b6101c790611603565b61169a6101c7916116926115fb565b506001610337565b61167a565b906116b4949392916116af612ee5565b61184d565b6101b1612f49565b67ffffffffffffffff81116107c65760208091020190565b906116e1610839836116bc565b918252565b6101c760c06107cb565b6101c76101406107cb565b6117036116f0565b9060208080808080808080808b5f8152015f8152015f8152015f8152015f5b8152015f6115e5565b6101c76116fb565b61173b6116e6565b90602080808080808761172261172b565b6101c7611733565b5f5b82811061176257505050565b60209061176d61174c565b8184015201611756565b906101b1611784836116d4565b9260208061179286936116bc565b9201910390611754565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b906117d2825190565b8110156117e3576020809102010190565b61179c565b91908110156117e3576020020190565b356101c7816102b6565b9035907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee136829003018212156101a0570190565b908210156117e35760206101c79202810190611802565b9493919290918261185d81611777565b916118675f610fd1565b828110156118d5576118d0906118c861188082876117c9565b516118b4611897611892858b8e6117e8565b6117f8565b73ffffffffffffffffffffffffffffffffffffffff1660a0830152565b82906118c1828b8f611836565b9091612fe2565b505060010190565b611867565b50945094909291506118e65f610fd1565b937fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f97261191160405190565b5f90a161191d5f610fd1565b945b8686101561195c5761194f6119569161081a8861193e8c898c91611836565b6119488b8b6117c9565b519161346d565b9560010190565b9461191f565b945092509450506101b19250613862565b906101b19493929161169f565b5f61139c6101c7926116925f90565b909161081a908390809361081e565b90916101c792611989565b906116e1610839836107df565b3d156119c9576119bf3d6119a3565b903d5f602084013e565b606090565b5f5b8381106119df5750505f910152565b81810151838201526020016119d0565b611a10611a1960209361081a93611a04815190565b80835293849260200190565b958691016119ce565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690565b90151581526101c79160408201915b60208184039101526119ef565b905f92839291611a78611a6f60405190565b92839283611998565b03915af4611a846119b0565b90610f02611a9160405190565b9283927f9941055400000000000000000000000000000000000000000000000000000000845260048401611a41565b905051906101b1826102b6565b906020828203126101a0576101c791611ac0565b9190611a1981611af88161081a9560209181520190565b809561081e565b60208082526101c793910191611ae1565b6040513d5f823e3d90fd5b90602090611b4d7f000000000000000000000000000000000000000000000000000000000000000061032e565b61032e565b611b725f63570e1a36959395611b7d611b6560405190565b9788968795869460e01b90565b845260048401611aff565b03925af18015611c0e57610f02915f91611be0575b506040519182917f6ca7b8060000000000000000000000000000000000000000000000000000000083526004830173ffffffffffffffffffffffffffffffffffffffff909116815260200190565b611c01915060203d8111611c07575b611bf98183610786565b810190611acd565b5f611b92565b503d611bef565b611b10565b61029d6101c76101c79290565b6101c790611c13565b15611c3057565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600560248201527f55534552300000000000000000000000000000000000000000000000000000006044820152606490fd5b6101c760606107cb565b9081526040810192916101b19160200152565b90611ce182611cda611cc061029d5f611c20565b9173ffffffffffffffffffffffffffffffffffffffff1690565b1415611c29565b4291611d2981611d2485611d1d86611d16611cfa611c8f565b73ffffffffffffffffffffffffffffffffffffffff9096168652565b6020850152565b6040830152565b6138e9565b90611d3f6001611d3a8460046106cc565b611120565b909192611d73611d6f7ff2d23449d3c6fc1266c59d62d6cb28ebcd25d872fc4c70d906c571e7247a8b009361032e565b9390565b93611d89611d8060405190565b92839283611c99565b0390a3565b15611d9557565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4141393220696e7465726e616c2063616c6c206f6e6c790000000000000000006044820152606490fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b8115611e2b570490565b611df4565b6101c7612710610fd1565b6101c7610800610fd1565b9081526101c7916040820191611a50565b6101c7913691610829565b92915a93611e7f33611e79611cc061029d3061032e565b14611d8e565b825194611e8d606087015190565b90611eb3611ea35a611e9f603f610fd1565b0290565b611ead6040610fd1565b90611e21565b611ed361136d6101c7611ecb8661081a60a08e015190565b61081a611e30565b10612028576101c796611f1e935f93611eea825190565b611ef661136d5f610fd1565b11611f24575b505050611f0c611f15915a900390565b60808601510190565b94909392611e57565b91613a19565b611f6091611f5c91611f4a855173ffffffffffffffffffffffffffffffffffffffff1690565b611f565f939293610fd1565b90613932565b1590565b611f6c575b8080611efc565b909150611f7f611f7a611e3b565b613944565b90611f88825190565b611f9461136d5f610fd1565b11611fa8575b506001919050611f0c611f65565b6020860151611fd66020611fd0845173ffffffffffffffffffffffffffffffffffffffff1690565b93015190565b926120086120027f1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a2019390565b9361032e565b9361201e61201560405190565b92839283611e46565b0390a35f80611f9a565b7fdeaddead000000000000000000000000000000000000000000000000000000005f5260205ffd5b1561205757565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f6e6f74207374616b6564000000000000000000000000000000000000000000006044820152606490fd5b156120bd57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f616c726561647920756e7374616b696e670000000000000000000000000000006044820152606490fd5b610c9b6101c76101c79263ffffffff1690565b6121489065ffffffffffff165b9165ffffffffffff1690565b019065ffffffffffff8211610fcc57565b6101c76101c76101c79265ffffffffffff1690565b6101ca90612159565b6020810192916101b1919061216e565b6121956101c7336001610337565b6121b76121a460018301610c89565b6121b061130a5f610e8f565b1415612050565b6121cb6121c660018301610c49565b6120b6565b61220c6121f46121da426110c1565b6121ee6121e960018601610c89565b61211c565b9061212f565b916122028360018301611206565b60015f9101611120565b7ffa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a6114636122393361032e565b9261224360405190565b91829182612177565b1561225357565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f4e6f207374616b6520746f2077697468647261770000000000000000000000006044820152606490fd5b156122b957565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6d7573742063616c6c20756e6c6f636b5374616b6528292066697273740000006044820152606490fd5b1561231f57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f5374616b65207769746864726177616c206973206e6f742064756500000000006044820152606490fd5b6101ca9061032e565b9160206101b19294936114e960408201965f83019061237e565b156123a857565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661696c656420746f207769746864726177207374616b6500000000000000006044820152606490fd5b5f806101b19261251561241f6101c760013390610337565b916124cb61243761243260018601610c6e565b610f6c565b9361244e8561244861136d89610fd1565b1161224c565b61246f61245d60018301610ca6565b61246961213c896110c1565b116122b2565b61248f61247e60018301610ca6565b6124884291612159565b1115612318565b6124a461249b87610e8f565b600183016111b6565b6124b96124b0876110c1565b60018301611206565b60016124c4876110aa565b910161116e565b807fb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3846124f73361032e565b9261250d61250460405190565b92839283612387565b0390a261032e565b9061251f60405190565b90818003925af161252e6119b0565b506123a1565b1561253b57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600360248201527f564e4500000000000000000000000000000000000000000000000000000000006044820152606490fd5b906125a6813414612534565b6125b982611cda611cc061029d5f611c20565b42916125d281611d2485611d1d86611d16611cfa611c8f565b906125e36001611d3a8460036106cc565b909192611d73611d6f7f2537ccb2b2fe11ea53e1afd15883c085006e0571b5b4f33dbf8becaf183f4e0b9361032e565b906116b49291612621612ee5565b6129bf565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa136829003018212156101a0570190565b908210156117e35760206101c79202810190612626565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe136829003018212156101a0570180359067ffffffffffffffff82116101a0576020019160208202360383136101a057565b156126cd57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4141393620696e76616c69642061676772656761746f720000000000000000006044820152606490fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe136829003018212156101a0570180359067ffffffffffffffff82116101a057602001913682900383136101a057565b506101c79060208101906102d2565b506101c79060208101906106ab565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe136839003018112156101a057016020813591019167ffffffffffffffff82116101a0573682900383136101a057565b6101c7916128e46128d861287861285d610120850161282b612811888061277f565b73ffffffffffffffffffffffffffffffffffffffff168752565b61284261283b602089018961278e565b6020880152565b61284f604088018861279d565b908783036040890152611ae1565b61286a606087018761279d565b908683036060880152611ae1565b61288f612888608087018761278e565b6080860152565b6128a661289f60a087018761278e565b60a0860152565b6128bd6128b660c087018761278e565b60c0860152565b6128ca60e086018661279d565b9085830360e0870152611ae1565b9261010081019061279d565b91610100818503910152611ae1565b906101c7916127ef565b90357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee136839003018112156101a0570190565b8183529160200190816129466020830284019490565b92835f925b84841061295b5750505050505090565b9091929394956020612987612980838560019503885261297b8b886128fd565b6128f3565b9860200190565b94019401929493919061294b565b92906129b1906101c7959360408601918683035f880152612930565b926020818503910152611ae1565b92919080936129cd5f610fd1565b936129d75f610fd1565b945b86861015612b5c576129ec86858561265a565b6129f68180612671565b91612a03602082016117f8565b90612a27612a108361032e565b612a20611cc061029d6001611c20565b14156126c6565b612a308261032e565b612a3f611cc061029d5f611c20565b03612a63575b505091612a5761194f92612a5d945090565b90610fb6565b946129d9565b612a6c8261032e565b632dd81133908490612a838794604081019061272c565b9093823b156101a057612ab6612aab925f96612a9e60405190565b9889978896879660e01b90565b865260048601612995565b03915afa9081612b3f575b50612b2d57612ad4565b92945092612a45565b612ae0610f029161032e565b6040519182917f86a9f7500000000000000000000000000000000000000000000000000000000083526004830173ffffffffffffffffffffffffffffffffffffffff909116815260200190565b5091612a5761194f92612a5d94612acb565b612b56905f612b4e8183610786565b8101906103fc565b5f612ac1565b612b6b91939294969550611777565b91612b755f610fd1565b90612b7f5f610fd1565b925b86841015612c4e57919695602094929391612b9d83868961265a565b94612bb5612bab8780612671565b98909897016117f8565b978699612bc15f610fd1565b945b8b865b1015612c2a578b612bc6612c2061136d9f612c19908f8f8f918f612c078f91936114b4956118c1612bfb612c13978a906117c9565b51918993908691611836565b9390939193919261032e565b92613cc0565b9860010190565b979e505050612bc3565b95939a5095985095509550612c40915060010190565b929493969196959095612b81565b96959493925090507fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972612c8060405190565b5f90a1612c8c5f610fd1565b93612c965f610fd1565b94612ca05f610fd1565b915b80831015612d9557612d0b612cbf87959998939a9788869161265a565b612cce611b48602083016117f8565b612cf87f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d9161032e565b90612d0260405190565b5f90a280612671565b909690948599612d1a5f610fd1565b995b8b8b5b1015612d71578b612d1f612d6661136d9f612d598f918f8f908f916119488f612d53612a5795612d5f998396919091611836565b926117c9565b99611468565b9d60010190565b9c979e505050612d1c565b949b995094919697509450612d8891985060010190565b9192949397959096612ca2565b50945050945050506101b191612daa5f611c20565b612dd47f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d9161032e565b90612dde60405190565b5f90a2613862565b906101b19291612613565b15612df857565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815280610f026004820160208082526004908201527f4c4e455100000000000000000000000000000000000000000000000000000000604082015260600190565b5050509795612e96919296612ea0986101b19b97612e8d612e7e878c905090565b612e878a61136d565b14612df1565b92909192613f71565b939091929361196d565b6140ca565b612ece7f01ffc9a700000000000000000000000000000000000000000000000000000000610dab565b1490565b61152b90614143565b6101c76002610fd1565b612eef6005610c36565b612efd61136d6101c7612edb565b14612f14576101b1612f0d612edb565b60056110fc565b6040517f3ee5aeb5000000000000000000000000000000000000000000000000000000008152600490fd5b6101c76001610fd1565b6101b1612f0d612f3f565b6101c76101c76101c7926effffffffffffffffffffffffffffff1690565b15612f7957565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f41413934206761732076616c756573206f766572666c6f7700000000000000006044820152606490fd5b356101c7816106a5565b9290915f925a8251612ff4818461424c565b613007613000846114ed565b6020860152565b60408101519661307d61301b60c084015190565b8917613028606085015190565b17613034608085015190565b1761304060a085015190565b1761304d61010085015190565b1761305a61012085015190565b1761307661136d6effffffffffffffffffffffffffffff612f54565b1115612f72565b613086826143a2565b90613094898388888561449c565b986130c4611f5c6130b9865173ffffffffffffffffffffffffffffffffffffffff1690565b60208701519061476e565b61320c576130d861136d610fbd875a900390565b116131975761310060e0606094015173ffffffffffffffffffffffffffffffffffffffff1690565b61310f611cc061029d5f611c20565b03613154575b509261081a60a061314860809561314261313b61314f976101c76101b19c9b60408d0152565b60608a0152565b5a900390565b9201612fd8565b910152565b6101c7975061314f925060a061314860809561314261313b6101b19a999661318361081a97898d9084926148ad565b9e9099505096999a50505095505050613115565b610f02906131a460405190565b9182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052601e908201527f41413236206f76657220766572696669636174696f6e4761734c696d69740000606082015260800190565b610f028261321960405190565b9182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052601a908201527f4141323520696e76616c6964206163636f756e74206e6f6e6365000000000000606082015260800190565b61329a6132946101c79263ffffffff1690565b60e01b90565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b805173ffffffffffffffffffffffffffffffffffffffff1682526101b191906101209081906132f360208201516020860152565b61330260408201516040860152565b61331160608201516060860152565b61332060808201516080860152565b61332f60a082015160a0860152565b61333e60c082015160c0860152565b60e08181015173ffffffffffffffffffffffffffffffffffffffff169085015261336f610100820151610100860152565b0151910152565b906101c060a06101b1936133905f8201515f8601906132bf565b6133a06020820151610140860152565b6133b06040820151610160860152565b6133c06060820151610180860152565b6133d060808201516101a0860152565b015173ffffffffffffffffffffffffffffffffffffffff16910152565b6102208082526101c7959391926134159261340a92850191611ae1565b936020830190613376565b6102008184039101526119ef565b9291602061343f6101b19360408701908782035f8901526127ef565b940152565b6102208082526101c794926134159161340a91908401906119ef565b91908203918211610fcc57565b9092915a61347f6101c7606084015190565b61358060205f60046040519961349981606081019061272c565b905f6003831161378e575b6134cd7f8dd7712f00000000000000000000000000000000000000000000000000000000610dab565b14851461375557505061356b906135326134e7868b015190565b916135246134f460405190565b938492878a85017f8dd7712f00000000000000000000000000000000000000000000000000000000815201613423565b878201810382520382610786565b61355d878a61354463b2dd2ba3613281565b9361354e60405190565b9687958a870190815201613444565b858201810382520382610786565b828151910182305af15f519760405215151590565b61358b575b50505050565b9091929394506135985f90565b3d602014613748575b806135cb7fdeaddead0000000000000000000000000000000000000000000000000000000061136d565b0361364557610f02856135dd60405190565b9182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052600f908201527f41413935206f7574206f66206761730000000000000000000000000000000000606082015260800190565b92935090916136737fdeadaa510000000000000000000000000000000000000000000000000000000061136d565b036136b7575061369461368a6136ac925a90613460565b6080840151612a57565b6040830151926136a381614b21565b83905f90614ba7565b905b5f808080613585565b61373a61368a84936136cd602061374297015190565b85515173ffffffffffffffffffffffffffffffffffffffff16865160200151916136f8611f7a611e3b565b6137236120027ff62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f47929390565b9361373061201560405190565b0390a35a90613460565b916002613a19565b906136ae565b5060205f803e5f516135a1565b61378992509061355d9088908b9061377063b2dd2ba3613281565b9461377a60405190565b9788968b8801908152016133ed565b61356b565b5080356134a4565b1561379d57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4141393020696e76616c69642062656e656669636961727900000000000000006044820152606490fd5b1561380357565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f41413931206661696c65642073656e6420746f2062656e6566696369617279006044820152606490fd5b5f6101b19261388c8293611b488161388561387f61029d88611c20565b9161032e565b1415613796565b9061389660405190565b90818003925af16138a56119b0565b506137fc565b6101ca9073ffffffffffffffffffffffffffffffffffffffff1660601b90565b602093926138e36014836138e3889561081a976138ab565b01918252565b61152b61390a825173ffffffffffffffffffffffffffffffffffffffff1690565b61151c613923604061391d602087015190565b95015190565b604051948593602085016138cb565b5f93849392909160208451940192f190565b3d90808211613968575b50604051906020810182016040528082525f602083013e90565b90505f61394e565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600311156139a757565b613970565b906101b18261399d565b6101c7906139ac565b6101ca906139b6565b6114e96139f66060936101b1969897956139e960808601925f8701906139bf565b84820360208601526119ef565b966040830152565b60208082526101c7929101906119ef565b6101c7600a610fd1565b909392905a92855192613a2b84614c54565b91613a4d60e086015173ffffffffffffffffffffffffffffffffffffffff1690565b80613a5d611cc061029d5f611c20565b03613b975750613aab91613a94915061081a613a8d875173ffffffffffffffffffffffffffffffffffffffff1690565b975a900390565b9361081a60a0613aa5606084015190565b92015190565b613abf84613aba60808a015190565b900390565b90818111613b66575b505082025b60408601519381851015613b34575050613af0613aea60026139ac565b916139ac565b03613b0c576101b19193613b0381614b21565b84905f90614ba7565b7fdeadaa51000000000000000000000000000000000000000000000000000000005f5260205ffd5b613b5090613b4a6101b196989594939889900390565b90614c7a565b50613b5d613aea5f6139ac565b14859192614ba7565b9361081a613b85613b7d613b8f94613acd96980390565b611e9f613a0f565b611ead6064610fd1565b92905f613ac8565b809691613ba2815190565b613bae61136d5f610fd1565b11613bc6575b5050613aab9161081a613a9492613142565b8385029186613bd8613aea60026139ac565b03613be4575b50613bb4565b611b48613bf09161032e565b90637c627b21613c0160a08a015190565b92909290889294908890803b156101a0575f95613c3e8793613c3395613c2660405190565b9a8b998a98899660e01b90565b8652600486016139c8565b0393f19081613cab575b50613c9b57613c5b565b92829450613bde565b610f02613c69611f7a611e3b565b6040519182917fad7954bc000000000000000000000000000000000000000000000000000000008352600483016139fe565b613aab9161081a613a9492613c52565b613cba905f612b4e8183610786565b5f613c48565b9290613ced613cd1611cc092614ca0565b93929073ffffffffffffffffffffffffffffffffffffffff1690565b03613ea057613e2b57613cff90614ca0565b90613d0f611cc061029d5f611c20565b03613db657613d1b5750565b610f0290613d2860405190565b9182917f220266b6000000000000000000000000000000000000000000000000000000008352600483019081526040602082018190526021908201527f41413332207061796d61737465722065787069726564206f72206e6f7420647560608201527f6500000000000000000000000000000000000000000000000000000000000000608082015260a00190565b610f0282613dc360405190565b9182917f220266b6000000000000000000000000000000000000000000000000000000008352600483019081526040602082018190526014908201527f41413334207369676e6174757265206572726f72000000000000000000000000606082015260800190565b610f0282613e3860405190565b9182917f220266b6000000000000000000000000000000000000000000000000000000008352600483019081526040602082018190526017908201527f414132322065787069726564206f72206e6f7420647565000000000000000000606082015260800190565b610f0283613ead60405190565b9182917f220266b6000000000000000000000000000000000000000000000000000000008352600483019081526040602082018190526014908201527f41413234207369676e6174757265206572726f72000000000000000000000000606082015260800190565b91908110156117e3576060020190565b91906060838203126101a057611d1d90613f3f60606107cb565b93613f4a82826102d2565b8552613f5982602083016106ab565b60208601526040016106ab565b6101c7903690613f25565b9093849391928291613f825f610fd1565b86811015613fe157613fdc90613fd6613fa4613f9f838c8b613f15565b613f66565b613fad81614d9d565b613fd06020613aa5835173ffffffffffffffffffffffffffffffffffffffff1690565b90614e2c565b60010190565b613f82565b5092955092509250613ff25f610fd1565b838110156140495761404490613fd661400f613f9f838988613f15565b61401881614ed5565b61403e6020613aa5611b48845173ffffffffffffffffffffffffffffffffffffffff1690565b9061501e565b613ff2565b5092505050565b1561405757565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815280610f026004820160208082526004908201527f4e52495a00000000000000000000000000000000000000000000000000000000604082015260600190565b9061110c6101c76111139290565b6101b1906140e8816140e161136d6101c75f610fd1565b1415614050565b5f6140bc565b959391989796949290986101008701995f880161411f9173ffffffffffffffffffffffffffffffffffffffff169052565b602087015260408601526060850152608084015260a083015260c082015260e00152565b6101c761414f826150b7565b9161151c61415f60208301612fd8565b91614176614170604083018361272c565b906150d1565b94614187614170606084018461272c565b9161419460808201612fd8565b6141a060a08301612fd8565b906141be6141706141b360c08601612fd8565b9460e081019061272c565b9396989490919293946141d060405190565b998a9860208a016140ee565b6101c76034610fd1565b156141ed57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4141393320696e76616c6964207061796d6173746572416e64446174610000006044820152606490fd5b6142e89061427661425c826117f8565b73ffffffffffffffffffffffffffffffffffffffff168452565b614285611d1660208301612fd8565b6142a861429c61429760808401612fd8565b615152565b60608601526040850152565b6142be6142b760a08301612fd8565b60c0850152565b6142de6142d061429760c08401612fd8565b610100860152610120850152565b60e081019061272c565b90816142f661136d5f610fd1565b1115614355576101b19261433b92909161432991906143248261431d61136d6101c76141dc565b10156141e6565b6152a9565b60a084015260e0830193919260800152565b73ffffffffffffffffffffffffffffffffffffffff169052565b50506101b1906143846143675f611c20565b73ffffffffffffffffffffffffffffffffffffffff1660e0830152565b6143976143905f610fd1565b6080830152565b60a061314f5f610fd1565b6101c790611e9f610100613aa56143de6143d56143cc6143c3604088015190565b60608801510190565b60808701510190565b60a08601510190565b60c08501510190565b905051906101b1826106a5565b906020828203126101a0576101c7916143e7565b6040906144316144296101b19597969460608401908482035f8601526127ef565b966020830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b908152606060208201819052600d908201527f414132332072657665727465640000000000000000000000000000000000000060808201526101c79160a082015b9160408184039101526119ef565b92919493906144a85f90565b86519096909260209061450a60e06144d4875173ffffffffffffffffffffffffffffffffffffffff1690565b966144ef868b6144e760408b018b61272c565b92909161534f565b015173ffffffffffffffffffffffffffffffffffffffff1690565b936145145f610fd1565b9185614525611cc061029d5f611c20565b146146db575b5f61457c9161453c611b488a61032e565b61458761456b60a06393fa30d29493949397989a015173ffffffffffffffffffffffffffffffffffffffff1690565b6040515b998a988997889560e01b90565b855260048501614408565b0393f15f91816146ab575b506146a4575061465c565b6145ac611cc061029d5f611c20565b146145b657505050565b6101c76145c4916001610337565b916145ce83610c36565b908183116145e757506101b192915f91035b91016110fc565b610f02906145f460405190565b9182917f220266b6000000000000000000000000000000000000000000000000000000008352600483019081526040602082018190526017908201527f41413231206469646e2774207061792070726566756e64000000000000000000606082015260800190565b83614668611f7a611e3b565b90610f0261467560405190565b9283927f65c8fd4d0000000000000000000000000000000000000000000000000000000084526004840161444d565b955061459d565b6146cd91925060203d81116146d4575b6146c58183610786565b8101906143f4565b905f614592565b503d6146bb565b91506146e68661197a565b87811115614705575061457c5f6146fc81610fd1565b9391505061452b565b5f61471261457c928a0390565b6146fc565b6101c79061472b61136d6101c79460ff1690565b901c90565b6103746101c76101c79290565b61474a6101c76101c79290565b67ffffffffffffffff1690565b6101c76101c76101c79267ffffffffffffffff1690565b9061136d6147b1612ece926147805f90565b506103c56147a96147a361479e846147986040611568565b90614717565b614730565b9261473d565b956002610337565b6147c76147bd82610c36565b9161123883611468565b92614757565b909291926147dd610839826107df565b938185526020850190828401116101a0576101b1926119ce565b9080601f830112156101a05781516101c7926020016147cd565b91906040838203126101a05782519067ffffffffffffffff82116101a05761483e816101c79386016147f7565b936020016143e7565b6040906114e96144296101b19597969460608401908482035f8601526127ef565b908152606060208201819052600d908201527f414133332072657665727465640000000000000000000000000000000000000060808201526101c79160a0820161448e565b919391906060945f945a915f810151936148de60e086015173ffffffffffffffffffffffffffffffffffffffff1690565b946148ed6101c7876001610337565b6148f681610c36565b858110614aa857614929611b4861492360805f98979661491e61494c978b6145e08e829a0390565b015190565b9961032e565b6352b7512c9061495761494260208c9394939798015190565b9861456f60405190565b855260048501614847565b0393f1805f80939092614a82575b50614a6d5750506001614a255761136d610fbd614982925a900390565b1161498a5750565b610f029061499760405190565b9182917f220266b6000000000000000000000000000000000000000000000000000000008352600483019081526040602082018190526027908201527f41413336206f766572207061796d6173746572566572696669636174696f6e4760608201527f61734c696d697400000000000000000000000000000000000000000000000000608082015260a00190565b82614a31611f7a611e3b565b90610f02614a3e60405190565b9283927f65c8fd4d00000000000000000000000000000000000000000000000000000000845260048401614868565b965094506149829061136d90610fbd90613142565b909250614aa191503d805f833e614a998183610786565b810190614811565b915f614965565b610f0289614ab560405190565b9182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052601e908201527f41413331207061796d6173746572206465706f73697420746f6f206c6f770000606082015260800190565b5090565b6020818101519151805191015173ffffffffffffffffffffffffffffffffffffffff9091169091611d89614b766120027f67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e9390565b936103ef60405190565b6114e96101b194611d1d606094989795614b9e608086019a5f870152565b15156020850152565b91614bb3602084015190565b83515190919073ffffffffffffffffffffffffffffffffffffffff1692614c4f614c0560205f614bfd60e0828b0151015173ffffffffffffffffffffffffffffffffffffffff1690565b980151015190565b919296614c3c614c36614c367f49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f9790565b9761032e565b97614c4660405190565b94859485614b80565b0390a4565b614c66610120613aa561010084015190565b808214614b1d576101c791904801906156e3565b614c98614c8e6101c76101c7936116925f90565b9261135784610c36565b9182906110fc565b80614cad61136d5f610fd1565b14614d2857614cbb90615711565b42614cda61136d614cd5604085015165ffffffffffff1690565b612159565b118015614d00575b905173ffffffffffffffffffffffffffffffffffffffff1691565b91565b50614cfd5f42614d1f61136d614cd5602087015165ffffffffffff1690565b10915050614ce2565b50614d325f611c20565b905f90565b15614d3e57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600560248201527f54484e45310000000000000000000000000000000000000000000000000000006044820152606490fd5b614da6816138e9565b90614dc2614dbd614db88460036106cc565b610c49565b614d37565b614dd15f611d3a8460036106cc565b614df46020613aa5835173ffffffffffffffffffffffffffffffffffffffff1690565b91614e1f7f03b8ee44084307fb90dbd84aef4a2bbaef367bb8a49bfbfc6e930a12f2f749ea9261032e565b92611463611d8060405190565b90614e379082614c7a565b90611463614e657f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c49261032e565b926103ef60405190565b15614e7657565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600560248201527f54484e45320000000000000000000000000000000000000000000000000000006044820152606490fd5b614ede816138e9565b90614ef5614ef0614db88460046106cc565b614e6f565b614f045f611d3a8460046106cc565b614f276020613aa5835173ffffffffffffffffffffffffffffffffffffffff1690565b91614e1f7ff89de5c5fbaa24098f4988752b4e8c7b8f11fc756a1688b549173fcb9396ea199261032e565b15614f5957565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f576974686472617720616d6f756e7420746f6f206c61726765000000000000006044820152606490fd5b15614fbf57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6661696c656420746f20776974686472617700000000000000000000000000006044820152606490fd5b5f6101b192615098829361506c61503a6101c760013390610337565b6150568561504f61136d6101c78a8601610c36565b1115614f52565b856145e0615065828401610c36565b8790613460565b807fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb846124f73361032e565b906150a260405190565b90818003925af16150b16119b0565b50614fb8565b6150cc6101c7916150c55f90565b5035610325565b610325565b60405190829082372090565b6150fc6101c76101c7926fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b6101c79060801c6150dd565b6101c790610fd1565b6150fc6101c76101c79290565b6101c76101c76101c7926fffffffffffffffffffffffffffffffff1690565b906101c76151a061519a61519561518f7fffffffffffffffffffffffffffffffff000000000000000000000000000000008716615111565b615111565b9561511d565b615126565b93615133565b92615133565b6101c76014610fd1565b909392938483116101a05784116101a0578101920390565b357fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001690601481106151f8575090565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000615227916014036008021b90565b1690565b6101c79060601c610302565b6101c79061522b565b6101c76024610fd1565b357fffffffffffffffffffffffffffffffff0000000000000000000000000000000016906010811061527a575090565b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000615227916010036008021b90565b916101c76151a061531c61518a6153036152e06152db6152d58a8a6152cc6151a6565b905f91926151b0565b906151c8565b615237565b9661530961518a6153038b846152f46151a6565b906152fd615240565b926151b0565b9061524a565b9890615313615240565b906152fd6141dc565b9394615133565b73ffffffffffffffffffffffffffffffffffffffff90911681526040810192916101b191602090614431565b909290918161536061136d5f610fd1565b0361536b5750505050565b83515173ffffffffffffffffffffffffffffffffffffffff1692833b61539361136d5f610fd1565b0361566e576153c17f000000000000000000000000000000000000000000000000000000000000000061032e565b85516040015160209063570e1a3690929092611b725f879389966153f46153e760405190565b9889978896879460e01b90565b0393f1908115611c0e575f91615650575b5080615416611cc061029d5f611c20565b146155db578061543b73ffffffffffffffffffffffffffffffffffffffff8716611cc0565b03615567573b61544d61136d5f610fd1565b146154f35750615474916152db916152d5916014906152fd61546e5f610fd1565b92610fd1565b906154a660e05f615486602087015190565b9394950151015173ffffffffffffffffffffffffffffffffffffffff1690565b6154d16120027fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d9390565b936154e76154de60405190565b92839283615323565b0390a35f808080613585565b610f029061550060405190565b9182917f220266b60000000000000000000000000000000000000000000000000000000083526004830190815260406020808301829052908201527f4141313520696e6974436f6465206d757374206372656174652073656e646572606082015260800190565b610f028261557460405190565b9182917f220266b60000000000000000000000000000000000000000000000000000000083526004830190815260406020808301829052908201527f4141313420696e6974436f6465206d7573742072657475726e2073656e646572606082015260800190565b610f02826155e860405190565b9182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052601b908201527f4141313320696e6974436f6465206661696c6564206f72204f4f470000000000606082015260800190565b615668915060203d8111611c0757611bf98183610786565b5f615405565b610f029061567b60405190565b9182917f220266b600000000000000000000000000000000000000000000000000000000835260048301908152604060208201819052601f908201527f414131302073656e64657220616c726561647920636f6e737472756374656400606082015260800190565b90808210156156f0575090565b905090565b6156fd611c8f565b9060208080845f6115ee565b6101c76156f5565b615719615709565b506101c76157296150cc83611c13565b9161573b6113fd8261479860a0611568565b908161574961213c5f6110c1565b1461578c575b6157636113fd61577d9261479860d0611568565b61576e611cfa611c8f565b65ffffffffffff166020850152565b65ffffffffffff166040830152565b65ffffffffffff915061574f56fea2646970667358221220db2d6a36f277e9f30812c5c247d16617172c8d202f7d0d8243990c6d31ad3fd264736f6c6343000817003360806040523461001a5760405161037c61001f823961037c90f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c63570e1a3603610055576100c1565b909182601f830112156100555781359167ffffffffffffffff831161005557602001926001830284011161005557565b5f80fd5b9060208282031261005557813567ffffffffffffffff81116100555761007f9201610025565b9091565b73ffffffffffffffffffffffffffffffffffffffff1690565b90565b73ffffffffffffffffffffffffffffffffffffffff909116815260200190565b565b34610055576100e96100dd6100d7366004610059565b906102c9565b6040519182918261009f565b0390f35b61009c61009c61009c9290565b90939293848311610055578411610055578101920390565b357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000169060148110610142575090565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000610171916014036008021b90565b1690565b61008361009c61009c9273ffffffffffffffffffffffffffffffffffffffff1690565b61009c9060601c610175565b61009c90610198565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b90601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810190811067ffffffffffffffff82111761021a57604052565b6101ad565b906100bf61022c60405190565b92836101da565b67ffffffffffffffff811161021a57602090601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160190565b90825f939282370152565b9092919261028e61028982610233565b61021f565b93818552602085019082840111610055576100bf9261026e565b61009c913691610279565b61008361009c61009c9290565b61009c906102b3565b5f906103266103206020946102db5f90565b5061030c610307610301866014856102fb6102f58c6100ed565b926100ed565b926100fa565b90610112565b6101a4565b939061031860146100ed565b9080926100fa565b906102a8565b90828483519301915af15f5191901561033b57565b905061009c5f6102c056fea2646970667358221220856f05374126b02ed11d10fcbab965a8054c6e488cfa70a58c943edb14c0a39a64736f6c63430008170033",
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

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_EntryPoint *EntryPointCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_EntryPoint *EntryPointSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _EntryPoint.Contract.GetNonce(&_EntryPoint.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_EntryPoint *EntryPointCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _EntryPoint.Contract.GetNonce(&_EntryPoint.CallOpts, sender, key)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointCaller) GetUserOpHash(opts *bind.CallOpts, userOp PackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _EntryPoint.Contract.GetUserOpHash(&_EntryPoint.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointCallerSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _EntryPoint.Contract.GetUserOpHash(&_EntryPoint.CallOpts, userOp)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_EntryPoint *EntryPointCaller) NonceSequenceNumber(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "nonceSequenceNumber", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_EntryPoint *EntryPointSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _EntryPoint.Contract.NonceSequenceNumber(&_EntryPoint.CallOpts, arg0, arg1)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_EntryPoint *EntryPointCallerSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _EntryPoint.Contract.NonceSequenceNumber(&_EntryPoint.CallOpts, arg0, arg1)
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

// AddDepositTicket is a paid mutator transaction binding the contract method 0xcd9ae171.
//
// Solidity: function addDepositTicket(address user, uint256 amount) payable returns()
func (_EntryPoint *EntryPointTransactor) AddDepositTicket(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "addDepositTicket", user, amount)
}

// AddDepositTicket is a paid mutator transaction binding the contract method 0xcd9ae171.
//
// Solidity: function addDepositTicket(address user, uint256 amount) payable returns()
func (_EntryPoint *EntryPointSession) AddDepositTicket(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.AddDepositTicket(&_EntryPoint.TransactOpts, user, amount)
}

// AddDepositTicket is a paid mutator transaction binding the contract method 0xcd9ae171.
//
// Solidity: function addDepositTicket(address user, uint256 amount) payable returns()
func (_EntryPoint *EntryPointTransactorSession) AddDepositTicket(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.AddDepositTicket(&_EntryPoint.TransactOpts, user, amount)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_EntryPoint *EntryPointTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_EntryPoint *EntryPointSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _EntryPoint.Contract.AddStake(&_EntryPoint.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_EntryPoint *EntryPointTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _EntryPoint.Contract.AddStake(&_EntryPoint.TransactOpts, unstakeDelaySec)
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

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactor) HandleAggregatedOps(opts *bind.TransactOpts, opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "handleAggregatedOps", opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_EntryPoint *EntryPointSession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleAggregatedOps(&_EntryPoint.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactorSession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleAggregatedOps(&_EntryPoint.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x68e9b1ea.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address[] userOpsAddrs, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactor) HandleOps(opts *bind.TransactOpts, ops []PackedUserOperation, userOpsAddrs []common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "handleOps", ops, userOpsAddrs, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x68e9b1ea.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address[] userOpsAddrs, address beneficiary) returns()
func (_EntryPoint *EntryPointSession) HandleOps(ops []PackedUserOperation, userOpsAddrs []common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, userOpsAddrs, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x68e9b1ea.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address[] userOpsAddrs, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactorSession) HandleOps(ops []PackedUserOperation, userOpsAddrs []common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, userOpsAddrs, beneficiary)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_EntryPoint *EntryPointTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_EntryPoint *EntryPointSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.IncrementNonce(&_EntryPoint.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_EntryPoint *EntryPointTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.IncrementNonce(&_EntryPoint.TransactOpts, key)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0xb2dd2ba3.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256,address) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactor) InnerHandleOp(opts *bind.TransactOpts, callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "innerHandleOp", callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0xb2dd2ba3.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256,address) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0xb2dd2ba3.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256,address) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactorSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_EntryPoint *EntryPointTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_EntryPoint *EntryPointSession) UnlockStake() (*types.Transaction, error) {
	return _EntryPoint.Contract.UnlockStake(&_EntryPoint.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_EntryPoint *EntryPointTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _EntryPoint.Contract.UnlockStake(&_EntryPoint.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_EntryPoint *EntryPointTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_EntryPoint *EntryPointSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.WithdrawStake(&_EntryPoint.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_EntryPoint *EntryPointTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.WithdrawStake(&_EntryPoint.TransactOpts, withdrawAddress)
}

// ZkAAEntryPoint is a paid mutator transaction binding the contract method 0xe997c9f9.
//
// Solidity: function zkAAEntryPoint(bytes proof, bytes32 pubInput, bytes32 newSmtRoot, (address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, address[] userOpsAddrs, (address,uint256,uint256)[] depositTickets, (address,uint256,uint256)[] withdrawTickets, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactor) ZkAAEntryPoint(opts *bind.TransactOpts, proof []byte, pubInput [32]byte, newSmtRoot [32]byte, userOps []PackedUserOperation, userOpsAddrs []common.Address, depositTickets []ITicketManagerTicket, withdrawTickets []ITicketManagerTicket, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "zkAAEntryPoint", proof, pubInput, newSmtRoot, userOps, userOpsAddrs, depositTickets, withdrawTickets, beneficiary)
}

// ZkAAEntryPoint is a paid mutator transaction binding the contract method 0xe997c9f9.
//
// Solidity: function zkAAEntryPoint(bytes proof, bytes32 pubInput, bytes32 newSmtRoot, (address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, address[] userOpsAddrs, (address,uint256,uint256)[] depositTickets, (address,uint256,uint256)[] withdrawTickets, address beneficiary) returns()
func (_EntryPoint *EntryPointSession) ZkAAEntryPoint(proof []byte, pubInput [32]byte, newSmtRoot [32]byte, userOps []PackedUserOperation, userOpsAddrs []common.Address, depositTickets []ITicketManagerTicket, withdrawTickets []ITicketManagerTicket, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.ZkAAEntryPoint(&_EntryPoint.TransactOpts, proof, pubInput, newSmtRoot, userOps, userOpsAddrs, depositTickets, withdrawTickets, beneficiary)
}

// ZkAAEntryPoint is a paid mutator transaction binding the contract method 0xe997c9f9.
//
// Solidity: function zkAAEntryPoint(bytes proof, bytes32 pubInput, bytes32 newSmtRoot, (address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, address[] userOpsAddrs, (address,uint256,uint256)[] depositTickets, (address,uint256,uint256)[] withdrawTickets, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactorSession) ZkAAEntryPoint(proof []byte, pubInput [32]byte, newSmtRoot [32]byte, userOps []PackedUserOperation, userOpsAddrs []common.Address, depositTickets []ITicketManagerTicket, withdrawTickets []ITicketManagerTicket, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.ZkAAEntryPoint(&_EntryPoint.TransactOpts, proof, pubInput, newSmtRoot, userOps, userOpsAddrs, depositTickets, withdrawTickets, beneficiary)
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
	Nonce        *big.Int
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostOpRevertReason is a free log retrieval operation binding the contract event 0xf62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
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

// WatchPostOpRevertReason is a free log subscription operation binding the contract event 0xf62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
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

// ParsePostOpRevertReason is a log parse operation binding the contract event 0xf62676f440ff169a3a9afdbf812e89e7f95975ee8e5c31214ffdef631c5f4792.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
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

// EntryPointStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the EntryPoint contract.
type EntryPointStakeLockedIterator struct {
	Event *EntryPointStakeLocked // Event containing the contract specifics and raw log

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
func (it *EntryPointStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointStakeLocked)
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
		it.Event = new(EntryPointStakeLocked)
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
func (it *EntryPointStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointStakeLocked represents a StakeLocked event raised by the EntryPoint contract.
type EntryPointStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_EntryPoint *EntryPointFilterer) FilterStakeLocked(opts *bind.FilterOpts, account []common.Address) (*EntryPointStakeLockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointStakeLockedIterator{contract: _EntryPoint.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_EntryPoint *EntryPointFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *EntryPointStakeLocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointStakeLocked)
				if err := _EntryPoint.contract.UnpackLog(event, "StakeLocked", log); err != nil {
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

// ParseStakeLocked is a log parse operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_EntryPoint *EntryPointFilterer) ParseStakeLocked(log types.Log) (*EntryPointStakeLocked, error) {
	event := new(EntryPointStakeLocked)
	if err := _EntryPoint.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointStakeUnlockedIterator is returned from FilterStakeUnlocked and is used to iterate over the raw logs and unpacked data for StakeUnlocked events raised by the EntryPoint contract.
type EntryPointStakeUnlockedIterator struct {
	Event *EntryPointStakeUnlocked // Event containing the contract specifics and raw log

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
func (it *EntryPointStakeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointStakeUnlocked)
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
		it.Event = new(EntryPointStakeUnlocked)
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
func (it *EntryPointStakeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointStakeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointStakeUnlocked represents a StakeUnlocked event raised by the EntryPoint contract.
type EntryPointStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeUnlocked is a free log retrieval operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_EntryPoint *EntryPointFilterer) FilterStakeUnlocked(opts *bind.FilterOpts, account []common.Address) (*EntryPointStakeUnlockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointStakeUnlockedIterator{contract: _EntryPoint.contract, event: "StakeUnlocked", logs: logs, sub: sub}, nil
}

// WatchStakeUnlocked is a free log subscription operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_EntryPoint *EntryPointFilterer) WatchStakeUnlocked(opts *bind.WatchOpts, sink chan<- *EntryPointStakeUnlocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointStakeUnlocked)
				if err := _EntryPoint.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
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

// ParseStakeUnlocked is a log parse operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_EntryPoint *EntryPointFilterer) ParseStakeUnlocked(log types.Log) (*EntryPointStakeUnlocked, error) {
	event := new(EntryPointStakeUnlocked)
	if err := _EntryPoint.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the EntryPoint contract.
type EntryPointStakeWithdrawnIterator struct {
	Event *EntryPointStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *EntryPointStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointStakeWithdrawn)
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
		it.Event = new(EntryPointStakeWithdrawn)
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
func (it *EntryPointStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointStakeWithdrawn represents a StakeWithdrawn event raised by the EntryPoint contract.
type EntryPointStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, account []common.Address) (*EntryPointStakeWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointStakeWithdrawnIterator{contract: _EntryPoint.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *EntryPointStakeWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointStakeWithdrawn)
				if err := _EntryPoint.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) ParseStakeWithdrawn(log types.Log) (*EntryPointStakeWithdrawn, error) {
	event := new(EntryPointStakeWithdrawn)
	if err := _EntryPoint.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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
	Nonce         *big.Int
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserOperationEvent is a free log retrieval operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
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

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
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

// ParseUserOperationEvent is a log parse operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
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
	Nonce      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserOperationPrefundTooLow is a free log retrieval operation binding the contract event 0x67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
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

// WatchUserOperationPrefundTooLow is a free log subscription operation binding the contract event 0x67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
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

// ParseUserOperationPrefundTooLow is a log parse operation binding the contract event 0x67b4fa9642f42120bf031f3051d1824b0fe25627945b27b8a6a65d5761d5482e.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
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
	Nonce        *big.Int
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserOperationRevertReason is a free log retrieval operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
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

// WatchUserOperationRevertReason is a free log subscription operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
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

// ParseUserOperationRevertReason is a log parse operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
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
