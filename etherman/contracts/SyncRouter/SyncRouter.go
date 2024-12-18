// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SyncRouter

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

// BaseStructCrossHookMessageParams is an auto generated low-level Go binding around an user-defined struct.
type BaseStructCrossHookMessageParams struct {
	Way                     uint8
	GasLimit                *big.Int
	GasPrice                uint64
	DestChainId             uint64
	MinArrivalTime          uint64
	MaxArrivalTime          uint64
	DestContract            common.Address
	SelectedRelayer         common.Address
	DestChainExecuteUsedFee *big.Int
	PackCrossMessage        []byte
	PackCrossParams         []byte
}

// BaseStructCrossMessageParams is an auto generated low-level Go binding around an user-defined struct.
type BaseStructCrossMessageParams struct {
	PackedUserOperation []BaseStructPackedUserOperation
	HookMessageParams   BaseStructCrossHookMessageParams
}

// BaseStructExecData is an auto generated low-level Go binding around an user-defined struct.
type BaseStructExecData struct {
	Nonce                  uint64
	ChainId                uint64
	MainChainGasLimit      uint64
	DestChainGasLimit      uint64
	ZkVerificationGasLimit uint64
	MainChainGasPrice      uint64
	DestChainGasPrice      uint64
	CallData               []byte
}

// BaseStructPackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type BaseStructPackedUserOperation struct {
	Phase          uint8
	OperationType  uint8
	OperationValue *big.Int
	Sender         common.Address
	Owner          common.Address
	Exec           BaseStructExecData
	InnerExec      BaseStructExecData
}

// SyncRouterMetaData contains all meta data concerning the SyncRouter contract.
var SyncRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vizingPad\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_Hook\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExecutionCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LandingPadAccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplement\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"packHookMessage\",\"type\":\"bytes\"}],\"name\":\"ReceiveTouchHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"valueDepositAmount\",\"type\":\"uint256\"}],\"name\":\"ValueDepositAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"VizingSwapEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"DataExcecuteNumber\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Hook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LandingPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LaunchPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LockWay\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"MirrorEntryPoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"newGaslimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"newGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newRelayer\",\"type\":\"address\"}],\"name\":\"changeDefaultSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBridgeMode\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGasPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGaslimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destChainid\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"additionParams\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"estimateVizingGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vizingGasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"fetchOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"fetchUserOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"cmp\",\"type\":\"tuple\"}],\"name\":\"getUserOmniEncodeMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveExtraMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveStandardMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selectedRelayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"sendOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"cmp\",\"type\":\"tuple\"}],\"name\":\"sendUserOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"}],\"name\":\"setMirrorEntryPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"testReceiveMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thisRelayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523462000038576200001f6200001862000131565b9162000694565b620000296200003e565b614d3262000a1a8239614d3290f35b62000044565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90620000729062000048565b810190811060018060401b038211176200008b57604052565b62000052565b90620000a8620000a06200003e565b928362000066565b565b5f80fd5b60018060a01b031690565b620000c490620000ae565b90565b620000d281620000b9565b03620000da57565b5f80fd5b90505190620000ed82620000c7565b565b90916060828403126200012b57620001286200010e845f8501620000de565b936200011e8160208601620000de565b93604001620000de565b90565b620000aa565b620001546200574c80380380620001488162000091565b928339810190620000ef565b909192565b60a01b90565b906200017060ff60a01b9162000159565b9181191691161790565b90565b60ff60f81b1690565b60f81b90565b620001a56200019f620001ab926200017a565b62000186565b6200017d565b90565b60f81c90565b620001bf90620001ae565b90565b90620001dc620001d6620001e4926200018c565b620001b4565b82546200015f565b9055565b90565b90565b90565b6200020a620002046200021092620001e8565b620001ee565b620001eb565b90565b60018060401b03811162000232576200022e60209162000048565b0190565b62000052565b906200024e620002488362000213565b62000091565b918252565b369037565b9062000283620002688362000238565b9260208062000278869362000213565b920191039062000253565b565b5190565b634e487b7160e01b5f52602260045260245ffd5b9060016002830492168015620002c0575b6020831014620002ba57565b62000289565b91607f1691620002ae565b5f5260205f2090565b601f602091010490565b1b90565b9190600862000301910291620002fa5f1984620002de565b92620002de565b9181191691161790565b620003246200031e6200032a92620001eb565b620001ee565b620001eb565b90565b90565b91906200034b6200034562000354936200030b565b6200032d565b908354620002e2565b9055565b5f90565b62000372916200036b62000358565b9162000330565b565b5b81811062000381575050565b80620003905f6001936200035c565b0162000375565b9190601f8111620003a8575b505050565b620003b7620003e293620002cb565b906020620003c584620002d4565b83019310620003eb575b620003da90620002d4565b019062000374565b5f8080620003a3565b9150620003da81929050620003cf565b1c90565b9062000411905f1990600802620003fb565b191690565b816200042291620003ff565b906002021790565b90620004368162000285565b9060018060401b03821162000507576200045d826200045685546200029d565b8562000397565b602090601f8311600114620004965791809162000484935f9262000489575b505062000416565b90555b565b90915001515f806200047c565b601f19831691620004a785620002cb565b925f5b818110620004ee57509160029391856001969410620004d1575b5050500201905562000487565b620004e3910151601f841690620003ff565b90555f8080620004c4565b91936020600181928787015181550195019201620004aa565b62000052565b9062000519916200042a565b565b5f1b90565b906200053062ffffff916200051b565b9181191691161790565b90565b62ffffff1690565b6200055e6200055862000564926200053a565b620001ee565b6200053d565b90565b90565b90620005846200057e6200058c9262000545565b62000567565b825462000520565b9055565b60181b90565b90620005ad6301000000600160581b039162000590565b9181191691161790565b90565b60018060401b031690565b620005de620005d8620005e492620005b7565b620001ee565b620005ba565b90565b90565b9062000604620005fe6200060c92620005c5565b620005e7565b825462000596565b9055565b906200062360018060a01b03916200051b565b9181191691161790565b62000646620006406200064c92620000ae565b620001ee565b620000ae565b90565b6200065a906200062d565b90565b62000668906200064f565b90565b90565b90620006886200068262000690926200065d565b6200066b565b825462000610565b9055565b6200070a9291620006aa6200070292336200070c565b620006b860016006620001c2565b620006d9620006d1620006cb5f620001f1565b62000258565b60076200050d565b620006e96207a12060086200056a565b620006fa633b9aca006008620005ea565b60056200066e565b60066200066e565b565b9062000718916200071a565b565b90620007269162000728565b565b906200073491620007a5565b565b6200074f6200074962000755926200017a565b620001ee565b620001eb565b90565b62000764600162000736565b90565b90620007755f19916200051b565b9181191691161790565b906200079962000793620007a1926200030b565b6200032d565b825462000767565b9055565b90620007b1916200081f565b620007c7620007bf62000758565b60036200077f565b565b620007e2620007dc620007e892620001e8565b620001ee565b620000ae565b90565b620007f690620007c9565b90565b6200080490620000b9565b9052565b91906200081d905f60208501940190620007f9565b565b906200082b9062000893565b806200084c620008456200083f5f620007eb565b620000b9565b91620000b9565b146200085f576200085d9062000928565b565b6200088f6200086e5f620007eb565b620008786200003e565b918291631e4fbdf760e01b83526004830162000808565b0390fd5b6200089e90620008a0565b565b620008ac9080620008ae565b565b620008bd620008c392620008c5565b620009e8565b565b620008d090620008d2565b565b620008dd90620008df565b565b620008ea9062000a01565b565b5f1c90565b60018060a01b031690565b6200090b6200091191620008ec565b620008f1565b90565b620009209054620008fc565b90565b5f0190565b62000934600262000914565b620009418260026200066e565b9062000979620009727f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936200065d565b916200065d565b91620009846200003e565b80620009908162000923565b0390a3565b620009a0906200062d565b90565b620009ae9062000995565b90565b620009bc9062000995565b90565b90565b90620009dc620009d6620009e492620009b1565b620009bf565b825462000610565b9055565b620009f7620009ff91620009a3565b6001620009c2565b565b62000a1062000a1791620009a3565b5f620009c2565b56fe60806040526004361015610015575b366116ba57005b61001f5f3561020c565b80622232e31461020757806273b555146102025780630e82845d146101fd5780631490a358146101f8578063397a567d146101f35780633eba5864146101ee57806345636279146101e95780635ad3ad06146101e45780635aeb4d77146101df5780635e45da23146101da57806360b43029146101d5578063715018a6146101d057806372afdea1146101cb57806376c81312146101c65780638da5cb5b146101c15780638ea23ddf146101bc57806393afae93146101b757806399e581fa146101b2578063a834a900146101ad578063ad5c4648146101a8578063afb3e17d146101a3578063b0cfd4d21461019e578063c5f378d814610199578063d5f3e00814610194578063de4b1d6d1461018f578063de8aeda01461018a578063e0b838e914610185578063e7b4294c14610180578063f2fde38b1461017b578063f3148925146101765763f329f52f0361000e57611684565b6111d1565b611118565b6110c5565b611080565b61103e565b610ff9565b610f24565b610ec2565b610e7e565b610e49565b610e14565b610dd1565b610d3f565b610cd5565b610ca0565b610c5c565b610c20565b610b7c565b610afd565b610ac8565b610a5b565b610a26565b6109c7565b61092d565b6108c9565b6107bb565b610750565b610653565b61058c565b610466565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906102549061022c565b810190811067ffffffffffffffff82111761026e57604052565b610236565b9061028661027f610212565b928361024a565b565b67ffffffffffffffff81116102a6576102a260209161022c565b0190565b610236565b90825f939282370152565b909291926102cb6102c682610288565b610273565b938185526020850190828401116102e7576102e5926102ab565b565b610228565b9080601f8301121561030a57816020610307933591016102b6565b90565b610224565b9060208282031261033f575f82013567ffffffffffffffff811161033a5761033792016102ec565b90565b610220565b61021c565b5190565b905090565b5f5b83811061035f575050905f910152565b80602091830151818501520161034f565b61039561038c9260209261038381610344565b94858093610348565b9384910161034d565b0190565b90565b90565b6103ab6103b091610399565b61039c565b9052565b6103c46103cb9160209493610370565b809261039f565b0190565b6103e36103da610212565b928392836103b4565b03902090565b6103f2916103cf565b90565b1c90565b60ff1690565b61040f90600861041493026103f5565b6103f9565b90565b9061042291546103ff565b90565b61043b90610436600c915f926103e9565b610417565b90565b60ff1690565b61044d9061043e565b9052565b9190610464905f60208501940190610444565b565b346104965761049261048161047c36600461030f565b610425565b610489610212565b91829182610451565b0390f35b610218565b67ffffffffffffffff1690565b6104b18161049b565b036104b857565b5f80fd5b905035906104c9826104a8565b565b6104d481610399565b036104db57565b5f80fd5b905035906104ec826104cb565b565b5f80fd5b5f80fd5b909182601f830112156105305781359167ffffffffffffffff831161052b57602001926001830284011161052657565b6104f2565b6104ee565b610224565b916060838303126105825761054c825f85016104bc565b9261055a83602083016104df565b92604082013567ffffffffffffffff811161057d5761057992016104f6565b9091565b610220565b61021c565b5f0190565b610597366004610535565b9291909161173d565b5f9103126105aa57565b61021c565b60018060a01b031690565b6105ca9060086105cf93026103f5565b6105af565b90565b906105dd91546105ba565b90565b6105ec60015f906105d2565b90565b60018060a01b031690565b90565b61061161060c610616926105ef565b6105fa565b6105ef565b90565b610622906105fd565b90565b61062e90610619565b90565b61063a90610625565b9052565b9190610651905f60208501940190610631565b565b34610683576106633660046105a0565b61067f61066e6105e0565b610676610212565b9182918261063e565b0390f35b610218565b906020828203126106a15761069e915f016104bc565b90565b61021c565b6106ba6106b56106bf9261049b565b6105fa565b61049b565b90565b906106cc906106a6565b5f5260205260405f2090565b60018060a01b031690565b6106f39060086106f893026103f5565b6106d8565b90565b9061070691546106e3565b90565b61071f9061071a600a915f926106c2565b6106fb565b90565b61072b906105ef565b90565b61073790610722565b9052565b919061074e905f6020850194019061072e565b565b346107805761077c61076b610766366004610688565b610709565b610773610212565b9182918261073b565b0390f35b610218565b906020828203126107b6575f82013567ffffffffffffffff81116107b1576107ad92016104f6565b9091565b610220565b61021c565b6107cf6107c9366004610785565b90611bdb565b6107d7610212565b806107e181610587565b0390f35b6107ee81610722565b036107f557565b5f80fd5b90503590610806826107e5565b565b909182601f830112156108425781359167ffffffffffffffff831161083d57602001926020830284011161083857565b6104f2565b6104ee565b610224565b906080828203126108a25761085e815f84016104bc565b9261086c82602085016107f9565b9261087a83604083016104df565b92606082013567ffffffffffffffff811161089d576108999201610808565b9091565b610220565b61021c565b6108b090610399565b9052565b91906108c7905f602085019401906108a7565b565b346108fd576108f96108e86108df366004610847565b9392909261235a565b6108f0610212565b918291826108b4565b0390f35b610218565b60ff60f81b1690565b61091490610902565b9052565b919061092b905f6020850194019061090b565b565b3461095d5761093d3660046105a0565b61095961094861244a565b610950610212565b91829182610918565b0390f35b610218565b67ffffffffffffffff1690565b61097f90600861098493026103f5565b610962565b90565b90610992915461096f565b90565b6109a26008600b90610987565b90565b6109ae9061049b565b9052565b91906109c5905f602085019401906109a5565b565b346109f7576109d73660046105a0565b6109f36109e2610995565b6109ea610212565b918291826109b2565b0390f35b610218565b62ffffff1690565b610a0d906109fc565b9052565b9190610a24905f60208501940190610a04565b565b34610a5657610a363660046105a0565b610a52610a41612459565b610a49610212565b91829182610a11565b0390f35b610218565b34610a8b57610a6b3660046105a0565b610a87610a76612464565b610a7e610212565b91829182610a11565b0390f35b610218565b90565b610aa3906008610aa893026103f5565b610a90565b90565b90610ab69154610a93565b90565b610ac560045f90610aab565b90565b34610af857610ad83660046105a0565b610af4610ae3610ab9565b610aeb610212565b918291826108b4565b0390f35b610218565b34610b2b57610b0d3660046105a0565b610b156124bc565b610b1d610212565b80610b2781610587565b0390f35b610218565b5f80fd5b90816040910312610b425790565b610b30565b90602082820312610b77575f82013567ffffffffffffffff8111610b7257610b6f9201610b34565b90565b610220565b61021c565b610b8f610b8a366004610b47565b612937565b610b97610212565b80610ba181610587565b0390f35b9190608083820312610c1b57610bbd815f85016104df565b92610bcb82602083016104bc565b92604082013567ffffffffffffffff8111610c165783610bec9184016104f6565b929093606082013567ffffffffffffffff8111610c1157610c0d92016104f6565b9091565b610220565b610220565b61021c565b34610c5757610c53610c42610c36366004610ba5565b94939093929192612950565b610c4a610212565b918291826108b4565b0390f35b610218565b34610c8c57610c6c3660046105a0565b610c88610c77612985565b610c7f610212565b9182918261073b565b0390f35b610218565b610c9d60065f906106fb565b90565b34610cd057610cb03660046105a0565b610ccc610cbb610c91565b610cc3610212565b9182918261073b565b0390f35b610218565b610cec610ce3366004610847565b93929092612c08565b610cf4610212565b80610cfe81610587565b0390f35b62ffffff1690565b610d1a906008610d1f93026103f5565b610d02565b90565b90610d2d9154610d0a565b90565b610d3c60085f90610d22565b90565b34610d6f57610d4f3660046105a0565b610d6b610d5a610d30565b610d62610212565b91829182610a11565b0390f35b610218565b610d7d816109fc565b03610d8457565b5f80fd5b90503590610d9582610d74565b565b9091606082840312610dcc57610dc9610db2845f8501610d88565b93610dc081602086016104bc565b936040016107f9565b90565b61021c565b34610e0057610dea610de4366004610d97565b91612d3e565b610df2610212565b80610dfc81610587565b0390f35b610218565b610e1160055f906106fb565b90565b34610e4457610e243660046105a0565b610e40610e2f610e05565b610e37610212565b9182918261073b565b0390f35b610218565b34610e7957610e75610e64610e5f366004610b47565b613072565b610e6c610212565b918291826108b4565b0390f35b610218565b34610eae57610e8e3660046105a0565b610eaa610e996131f7565b610ea1610212565b9182918261073b565b0390f35b610218565b610ebf60095f906106fb565b90565b34610ef257610ed23660046105a0565b610eee610edd610eb3565b610ee5610212565b9182918261073b565b0390f35b610218565b9190604083820312610f1f5780610f13610f1c925f86016104bc565b936020016107f9565b90565b61021c565b34610f5357610f3d610f37366004610ef7565b9061322c565b610f45610212565b80610f4f81610587565b0390f35b610218565b90602082820312610f7157610f6e915f016104df565b90565b61021c565b610f8a610f85610f8f92610399565b6105fa565b610399565b90565b90610f9c90610f76565b5f5260205260405f2090565b60f81b90565b610fb790610fa8565b90565b610fca906008610fcf93026103f5565b610fae565b90565b90610fdd9154610fba565b90565b610ff690610ff1600b915f92610f92565b610fd2565b90565b346110295761102561101461100f366004610f58565b610fe0565b61101c610212565b91829182610918565b0390f35b610218565b61103b6008601390610987565b90565b3461106e5761104e3660046105a0565b61106a61105961102e565b611061610212565b918291826109b2565b0390f35b610218565b61107d5f806105d2565b90565b346110b0576110903660046105a0565b6110ac61109b611073565b6110a3610212565b9182918261063e565b0390f35b610218565b6110c26008600390610987565b90565b346110f5576110d53660046105a0565b6110f16110e06110b5565b6110e8610212565b918291826109b2565b0390f35b610218565b9060208282031261111357611110915f016107f9565b90565b61021c565b346111465761113061112b3660046110fa565b6132a4565b611138610212565b8061114281610587565b0390f35b610218565b90565b6111578161114b565b0361115e57565b5f80fd5b9050359061116f8261114e565b565b906080828203126111cc57611188815f8401611162565b9261119682602085016104bc565b926111a483604083016104df565b92606082013567ffffffffffffffff81116111c7576111c392016104f6565b9091565b610220565b61021c565b6111e86111df366004611171565b9392909261331d565b6111f0610212565b806111fa81610587565b0390f35b5f80fd5b5f80fd5b67ffffffffffffffff811161121e5760208091020190565b610236565b61122c8161043e565b0361123357565b5f80fd5b9050359061124482611223565b565b919091610100818403126113065761125f610100610273565b9261126c815f84016104bc565b5f85015261127d81602084016104bc565b602085015261128f81604084016104bc565b60408501526112a181606084016104bc565b60608501526112b381608084016104bc565b60808501526112c58160a084016104bc565b60a08501526112d78160c084016104bc565b60c085015260e082013567ffffffffffffffff8111611301576112fa92016102ec565b60e0830152565b611202565b6111fe565b91909160e0818403126113cf5761132260e0610273565b9261132f815f8401611237565b5f8501526113408160208401611237565b602085015261135281604084016104df565b604085015261136481606084016107f9565b606085015261137681608084016107f9565b608085015260a082013567ffffffffffffffff81116113ca578161139b918401611246565b60a085015260c082013567ffffffffffffffff81116113c5576113be9201611246565b60c0830152565b611202565b611202565b6111fe565b9291906113e86113e382611206565b610273565b938185526020808601920281019183831161143f5781905b83821061140e575050505050565b813567ffffffffffffffff811161143a5760209161142f878493870161130b565b815201910190611400565b610224565b6104f2565b9080601f830112156114625781602061145f933591016113d4565b90565b610224565b9190916101608184031261157b57611480610160610273565b9261148d815f8401611237565b5f85015261149e8160208401610d88565b60208501526114b081604084016104bc565b60408501526114c281606084016104bc565b60608501526114d481608084016104bc565b60808501526114e68160a084016104bc565b60a08501526114f88160c084016107f9565b60c085015261150a8160e084016107f9565b60e085015261151d8161010084016104df565b61010085015261012082013567ffffffffffffffff811161157657816115449184016102ec565b61012085015261014082013567ffffffffffffffff81116115715761156992016102ec565b610140830152565b611202565b611202565b6111fe565b9190916040818403126115ea576115976040610273565b925f82013567ffffffffffffffff81116115e557816115b7918401611444565b5f850152602082013567ffffffffffffffff81116115e0576115d99201611467565b6020830152565b611202565b611202565b6111fe565b9060208282031261161f575f82013567ffffffffffffffff811161161a576116179201611580565b90565b610220565b61021c565b60209181520190565b61164c61165560209361165a9361164381610344565b93848093611624565b9586910161034d565b61022c565b0190565b916116819261167460408201935f8301906108a7565b602081840391015261162d565b90565b346116b55761169c6116973660046115ef565b613be0565b906116b16116a8610212565b9283928361165e565b0390f35b610218565b5f80fd5b5f1c90565b6116cf6116d4916116be565b6105af565b90565b6116e190546116c3565b90565b336117086117026116fd6116f860016116d7565b610625565b610722565b91610722565b031561173157611716610212565b637d92a0f560e11b81528061172d60048201610587565b0390fd5b82849192909192614485565b6116e4565b5190565b90565b61175d61175861176292611746565b6105fa565b610399565b90565b61177961177461177e92610399565b6105fa565b61049b565b90565b61178d611792916116be565b6106d8565b90565b61179f9054611781565b90565b6117ab906105fd565b90565b6117b7906117a2565b90565b6117c390610619565b90565b5f80fd5b60e01b90565b5f9103126117da57565b61021c565b60209181520190565b60200190565b6117f79061043e565b9052565b61180490610399565b9052565b61181190610722565b9052565b61181e9061049b565b9052565b60209181520190565b61184a6118536020936118589361184181610344565b93848093611822565b9586910161034d565b61022c565b0190565b6118f49160e06101008201926118785f8201515f850190611815565b61188a60208201516020850190611815565b61189c60408201516040850190611815565b6118ae60608201516060850190611815565b6118c060808201516080850190611815565b6118d260a082015160a0850190611815565b6118e460c082015160c0850190611815565b01519060e081840391015261182b565b90565b61197f9160c061196e60e083016119145f8601515f8601906117ee565b611926602086015160208601906117ee565b611938604086015160408601906117fb565b61194a60608601516060860190611808565b61195c60808601516080860190611808565b60a085015184820360a086015261185c565b9201519060c081840391015261185c565b90565b9061198c916118f7565b90565b60200190565b906119a96119a283611742565b80926117df565b90816119ba602083028401946117e8565b925f915b8383106119cd57505050505090565b909192939460206119ef6119e983856001950387528951611982565b9761198f565b93019301919392906119be565b611a119160208201915f818403910152611995565b90565b611a1c610212565b3d5f823e3d90fd5b5f90565b606090565b90505190611a3a826104cb565b565b90505190611a49826107e5565b565b9190604083820312611a8557611a7e90611a656040610273565b93611a72825f8301611a2d565b5f860152602001611a3c565b6020830152565b6111fe565b90604082820312611aa357611aa0915f01611a4b565b90565b61021c565b611ab2905161043e565b90565b90565b611acc611ac7611ad192611ab5565b6105fa565b61043e565b90565b611ae8611ae3611aed92611746565b6105fa565b61043e565b90565b90565b611b07611b02611b0c92611af0565b6105fa565b61043e565b90565b611b1890610619565b90565b611b259051610399565b90565b90611b3a611b3583610288565b610273565b918252565b3d5f14611b5a57611b4f3d611b28565b903d5f602084013e5b565b611b62611a28565b90611b58565b611b729051610722565b90565b611b805f8092610348565b0190565b611b8d90611b75565b90565b151590565b611b9e90611b90565b9052565b91611bd89391611bca91611bbd60608601925f870190611b95565b848203602086015261162d565b91604081840391015261162d565b90565b90611be991908101906115ef565b611bf55f820151611742565b611c07611c015f611749565b91610399565b03611dd1575b611c15611a24565b50611c1e611a28565b505f80611c44610140602085015101516020611c3982610344565b818301019101611a8a565b611c5382602086015101611aa8565b611c66611c6060ff611ab8565b9161043e565b148214611cfa57611c9482611c8d611c88600a611c8246611765565b906106c2565b611795565b9201611b1b565b6101206020860151015190602082019151925af1906101206020611cb6611b3f565b935b93920151015191611cf57fc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb39807944693611cec610212565b93849384611ba2565b0390a1565b611d0982602086015101611aa8565b611d1b611d1584611ad4565b9161043e565b148015611da8575b8214611d6c57611d4082611d3960208401611b68565b9201611b1b565b611d48610212565b9081611d5381611b84565b03925af1906101206020611d65611b3f565b935b611cb8565b611d8082611d7930611b0f565b9201611b1b565b6101206020860151015190602082019151925af1906101206020611da2611b3f565b93611d67565b50611db882602086015101611aa8565b611dcb611dc560fe611af3565b9161043e565b14611d23565b611dfd611df8611df3611dee600a611de846611765565b906106c2565b611795565b6117ae565b6117ba565b6377c2719c5f830151823b15611e7c57611e3692611e2b5f8094611e1f610212565b968795869485936117ca565b8352600483016119fc565b03925af18015611e7757611e4b575b50611c0d565b611e6a905f3d8111611e70575b611e62818361024a565b8101906117d0565b5f611e45565b503d611e58565b611a14565b6117c6565b5f90565b90565b50611e97906020810190611237565b90565b50611ea99060208101906104df565b90565b50611ebb9060208101906107f9565b90565b5f80fd5b903560016101000382360303811215611ed9570190565b611ebe565b50611eed9060208101906104bc565b90565b5f80fd5b5f80fd5b9035600160200382360303811215611f3957016020813591019167ffffffffffffffff8211611f34576001820236038313611f2f57565b611ef4565b611ef0565b611ebe565b9190611f5881611f5181611f5d95611822565b80956102ab565b61022c565b0190565b61203a9161202c610100820192611f86611f7d5f830183611ede565b5f850190611815565b611fa0611f966020830183611ede565b6020850190611815565b611fba611fb06040830183611ede565b6040850190611815565b611fd4611fca6060830183611ede565b6060850190611815565b611fee611fe46080830183611ede565b6080850190611815565b612008611ffe60a0830183611ede565b60a0850190611815565b61202261201860c0830183611ede565b60c0850190611815565b60e0810190611ef8565b9160e0818503910152611f3e565b90565b6120fe916120f06120e560e0830161206361205a5f870187611e88565b5f8601906117ee565b61207d6120736020870187611e88565b60208601906117ee565b61209761208d6040870187611e9a565b60408601906117fb565b6120b16120a76060870187611eac565b6060860190611808565b6120cb6120c16080870187611eac565b6080860190611808565b6120d860a0860186611ec2565b84820360a0860152611f61565b9260c0810190611ec2565b9060c0818403910152611f61565b90565b9061210b9161203d565b90565b9035600160e00382360303811215612124570190565b611ebe565b60200190565b918161213a916117df565b908161214b60208302840194611e85565b92835f925b8484106121605750505050505090565b909192939495602061218b61218583856001950388526121808b8861210e565b612101565b98612129565b940194019294939190612150565b90916121b09260208301925f81850391015261212f565b90565b60a01c90565b6121c56121ca916121b3565b610fae565b90565b6121d790546121b9565b90565b6121e66121eb916116be565b610d02565b90565b6121f890546121da565b90565b60181c90565b61220d612212916121fb565b610962565b90565b61221f9054612201565b90565b9060208282031261223b57612238915f01611a2d565b90565b61021c565b634e487b7160e01b5f52602260045260245ffd5b9060016002830492168015612274575b602083101461226f57565b612240565b91607f1691612264565b5f5260205f2090565b905f92918054906122a161229a83612254565b8094611624565b916001811690815f146122f857506001146122bc575b505050565b6122c9919293945061227e565b915f925b8184106122e057505001905f80806122b7565b600181602092959395548486015201910192906122cd565b92949550505060ff19168252151560200201905f80806122b7565b92906123499161233c612357969461233260808801945f8901906108a7565b60208701906109a5565b8482036040860152612287565b91606081840391015261162d565b90565b9193906123bb9161239460209561236f611e81565b5061238661237b610212565b938492898401612199565b87820181038252038261024a565b61239e60066121cd565b916123a960086121ee565b906123b46008612215565b929361457b565b926123f26123d06123cb5f6116d7565b610625565b916123fd6385fdd54291949660076123e6610212565b988997889687966117ca565b865260048601612313565b03915afa908115612441575f91612413575b5090565b612434915060203d811161243a575b61242c818361024a565b810190612222565b5f61240f565b503d612422565b611a14565b5f90565b612452612446565b90565b5f90565b612461612455565b90565b61246c612455565b90565b6124776145ca565b61247f6124a9565b565b61249561249061249a92611746565b6105fa565b6105ef565b90565b6124a690612481565b90565b6124ba6124b55f61249d565b61461f565b565b6124c461246f565b565b6124d7906124d2614703565b612690565b6124df61478b565b565b6124ec903690611580565b90565b5f80fd5b5f80fd5b5f80fd5b90359060016101600381360303821215612513570190565b6124ef565b35612522816107e5565b90565b3561252f81610d74565b90565b3561253c816104a8565b90565b35612549816104cb565b90565b634e487b7160e01b5f52601160045260245ffd5b61256f61257591939293610399565b92610399565b820180921161258057565b61254c565b60209181520190565b5f7f53656e642065746820496e73756666696369656e740000000000000000000000910152565b6125c26015602092612585565b6125cb8161258e565b0190565b6125e49060208101905f8183039101526125b5565b90565b156125ee57565b6125f6610212565b62461bcd60e51b81528061260c600482016125cf565b0390fd5b969261267f9561265e61268d9a9893966126546126689461264a6126729860208f61264361010082019e5f8301906109a5565b01906109a5565b60408d019061072e565b60608b019061072e565b60808901906108a7565b60a08701906109a5565b84820360c0860152612287565b9160e081840391015261162d565b90565b612705906126a56126a0826124e1565b613be0565b9290926126b260066121cd565b906126cc60c06126c68660208101906124fb565b01612518565b6126e460206126de87828101906124fb565b01612525565b906126fe60406126f88860208101906124fb565b01612532565b929361457b565b916127176127125f6116d7565b610625565b60206385fdd5429161274261273b61010061273588868101906124fb565b0161253f565b8590612560565b9061277a61275e606061275889878101906124fb565b01612532565b9461278560078a9061276e610212565b988997889687966117ca565b865260048601612313565b03915afa8015612932576127e2915f91612904575b506127db6127d56127d06127c934946127c36101006127bd8b60208101906124fb565b0161253f565b90612560565b8690612560565b610399565b91610399565b10156125e7565b6127f36127ee5f6116d7565b610625565b8263209afe569134909261281660806128108560208101906124fb565b01612532565b9461283060a061282a8660208101906124fb565b01612532565b97612887606061288161287661285560e061284f8b60208101906124fb565b01612518565b9561287161010061286b339c60208101906124fb565b0161253f565b612560565b9a60208101906124fb565b01612532565b9060079091873b156128ff575f996128b1976128bc956128a5610212565b9d8e9c8d9b8c9a6117ca565b8a5260048a01612610565b03925af180156128fa576128ce575b50565b6128ed905f3d81116128f3575b6128e5818361024a565b8101906117d0565b5f6128cb565b503d6128db565b611a14565b6117c6565b612925915060203d811161292b575b61291d818361024a565b810190612222565b5f61279a565b503d612913565b611a14565b612940906124c6565b565b61294d9136916102b6565b90565b92612972612978929561297e9795612966611e81565b50969492909592612942565b93612942565b926147e6565b90565b5f90565b61298d612981565b506129986002611795565b90565b5f7f4d45500000000000000000000000000000000000000000000000000000000000910152565b6129cf6003602092612585565b6129d88161299b565b0190565b6129f19060208101905f8183039101526129c2565b90565b156129fb57565b612a03610212565b62461bcd60e51b815280612a19600482016129dc565b0390fd5b90612a6194939291612a5c612a3146611765565b612a56612a50612a4b612a463394600a6106c2565b611795565b610722565b91610722565b146129f4565b612abc565b565b15612a6a57565b5f80fd5b60581c90565b612a80612a8591612a6e565b610962565b90565b612a929054612a74565b90565b60981c90565b612aa7612aac91612a95565b610962565b90565b612ab99054612a9b565b90565b9293612b4b91612b2791612b1b612ad360066121cd565b83612ade60086121ee565b612ae86008612215565b918b93612b168795612b07612afb610212565b97889260208401612199565b6020820181038252038661024a565b61457b565b9686928691929361235a565b612b44612b3e612b3934938690612560565b610399565b91610399565b1015612a63565b612b5c612b575f6116d7565b610625565b9063209afe569034929192612b716008612a88565b93612b7c6008612aaf565b96612b876009611795565b903394979060079091873b15612c03575f99612bb597612bc095612ba9610212565b9d8e9c8d9b8c9a6117ca565b8a5260048a01612610565b03925af18015612bfe57612bd2575b50565b612bf1905f3d8111612bf7575b612be9818361024a565b8101906117d0565b5f612bcf565b503d612bdf565b611a14565b6117c6565b90612c1594939291612a1d565b565b90612c2a9291612c256145ca565b612d1b565b565b5f1b90565b90612c3f62ffffff91612c2c565b9181191691161790565b612c5d612c58612c62926109fc565b6105fa565b6109fc565b90565b90565b90612c7d612c78612c8492612c49565b612c65565b8254612c31565b9055565b60181b90565b90612ca46affffffffffffffff00000091612c88565b9181191691161790565b90565b90612cc6612cc1612ccd926106a6565b612cae565b8254612c8e565b9055565b90612ce260018060a01b0391612c2c565b9181191691161790565b612cf590610619565b90565b90565b90612d10612d0b612d1792612cec565b612cf8565b8254612cd1565b9055565b612d3c9291612d2e612d35926008612c68565b6008612cb1565b6009612cfb565b565b90612d499291612c17565b565b9035600160200382360303811215612d8c57016020813591019167ffffffffffffffff8211612d87576020820236038313612d8257565b611ef4565b611ef0565b611ebe565b60209181520190565b9181612da591612d91565b9081612db660208302840194611e85565b92835f925b848410612dcb5750505050505090565b9091929394956020612df6612df08385600195038852612deb8b8861210e565b612101565b98612129565b940194019294939190612dbb565b903560016101600382360303811215612e1b570190565b611ebe565b50612e2f906020810190610d88565b90565b612e3b906109fc565b9052565b612f7091612f61612f556101608301612e66612e5d5f870187611e88565b5f8601906117ee565b612e80612e766020870187612e20565b6020860190612e32565b612e9a612e906040870187611ede565b6040860190611815565b612eb4612eaa6060870187611ede565b6060860190611815565b612ece612ec46080870187611ede565b6080860190611815565b612ee8612ede60a0870187611ede565b60a0860190611815565b612f02612ef860c0870187611eac565b60c0860190611808565b612f1c612f1260e0870187611eac565b60e0860190611808565b612f38612f2d610100870187611e9a565b6101008601906117fb565b612f46610120860186611ef8565b90858303610120870152611f3e565b92610140810190611ef8565b91610140818503910152611f3e565b90565b612fb391612fa5612f9a60408301612f8d5f860186612d4b565b908583035f870152612d9a565b926020810190612e04565b906020818403910152612e3f565b90565b612fcb9160208201915f818403910152612f73565b90565b903590600160200381360303821215613010570180359067ffffffffffffffff821161300b5760200191600182023603831361300657565b6124f7565b6124f3565b6124ef565b919060408382031261304f576130489061302f6040610273565b9361303c825f83016104df565b5f8601526020016107f9565b6020830152565b6111fe565b9060408282031261306d5761306a915f01613015565b90565b61021c565b61307a611e81565b5060206130976130a58361308c610212565b928391858301612fb6565b84820181038252038261024a565b916131346130d95f6130d36130cb6130c086888101906124fb565b610140810190612fce565b810190613054565b01611b1b565b936130e460066121cd565b906130fd60c06130f786888101906124fb565b01612518565b6131148661310e87828101906124fb565b01612525565b9061312d6040613127888a8101906124fb565b01612532565b929361457b565b926131a36131496131445f6116d7565b610625565b916131ae61318c606061318661317c6385fdd542956131776101006131718c8e8101906124fb565b0161253f565b612560565b97898101906124fb565b01612532565b966007613197610212565b988997889687966117ca565b865260048601612313565b03915afa9081156131f2575f916131c4575b5090565b6131e5915060203d81116131eb575b6131dd818361024a565b810190612222565b5f6131c0565b503d6131d3565b611a14565b6131ff612981565b90565b906132149161320f6145ca565b613216565b565b61322561322a9291600a6106c2565b612cfb565b565b9061323691613202565b565b613249906132446145ca565b61324b565b565b8061326661326061325b5f61249d565b610722565b91610722565b14613276576132749061461f565b565b6132a06132825f61249d565b61328a610212565b918291631e4fbdf760e01b83526004830161073b565b0390fd5b6132ad90613238565b565b93929190336132d76132d16132cc6132c760016116d7565b610625565b610722565b91610722565b036132e7576132e59461330a565b565b6132ef610212565b637d92a0f560e11b81528061330660048201610587565b0390fd5b9161331b9492939190919293614977565b565b9061332a949392916132af565b565b90565b61334361333e6133489261332c565b6105fa565b61043e565b90565b90565b61336261335d6133679261334b565b6105fa565b61043e565b90565b613375610120610273565b90565b5f90565b5f90565b5f90565b5f90565b5f90565b61339461336a565b90602080808080808080808a6133a8613378565b8152016133b361337c565b8152016133be613380565b8152016133c9613384565b8152016133d4613384565b8152016133df613384565b8152016133ea613384565b8152016133f5613388565b815201613400613388565b81525050565b61340e61338c565b90565b634e487b7160e01b5f52603260045260245ffd5b9061342f82611742565b811015613440576020809102010190565b613411565b613451613456916116be565b6103f9565b90565b6134639054613445565b90565b9050519061347382611223565b565b9050519061348282610d74565b565b61348d816105ef565b0361349457565b5f80fd5b905051906134a582613484565b565b9190610120838203126135635761355b906134c3610120610273565b936134d0825f8301613466565b5f8601526134e18260208301613475565b60208601526134f38260408301613498565b60408601526135058260608301611a3c565b60608601526135178260808301611a3c565b60808601526135298260a08301611a3c565b60a086015261353b8260c08301611a3c565b60c086015261354d8260e08301611a2d565b60e086015261010001611a2d565b610100830152565b6111fe565b90610120828203126135825761357f915f016134a7565b90565b61021c565b63ffffffff1690565b63ffffffff60e01b1690565b6135b06135ab6135b592613587565b6117ca565b613590565b90565b6135c1906105ef565b9052565b9061010080613666936135de5f8201515f8601906117ee565b6135f060208201516020860190612e32565b613602604082015160408601906135b8565b61361460608201516060860190611808565b61362660808201516080860190611808565b61363860a082015160a0860190611808565b61364a60c082015160c0860190611808565b61365c60e082015160e08601906117fb565b01519101906117fb565b565b919061367c905f61012085019401906135c5565b565b61368860e0610273565b90565b606090565b61369861367e565b906020808080808080886136aa613378565b8152016136b5613388565b8152016136c0613388565b8152016136cb61368b565b8152016136d6613384565b8152016136e1613384565b8152016136ec613388565b81525050565b6136fa613690565b90565b67ffffffffffffffff81116137155760208091020190565b610236565b9092919261372f61372a826136fd565b610273565b938185526020808601920283019281841161376c57915b8383106137535750505050565b602080916137618486611a3c565b815201920191613746565b6104f2565b9080601f8301121561378f5781602061378c9351910161371a565b90565b610224565b91909160e081840312613841576137ab60e0610273565b926137b8815f8401613466565b5f8501526137c98160208401611a2d565b60208501526137db8160408401611a2d565b604085015260608201519167ffffffffffffffff831161383c5761380482613835948301613771565b60608601526138168260808301611a3c565b60808601526138288260a08301611a3c565b60a086015260c001611a2d565b60c0830152565b611202565b6111fe565b90602082820312613876575f82015167ffffffffffffffff81116138715761386e9201613794565b90565b610220565b61021c565b5190565b60209181520190565b60200190565b9061389b81602093611808565b0190565b60200190565b906138c26138bc6138b58461387b565b809361387f565b92613888565b905f5b8181106138d25750505090565b9091926138eb6138e5600192865161388e565b9461389f565b91019190916138c5565b906139799060c08061394a60e084016139145f8801515f8701906117ee565b613926602088015160208701906117fb565b613938604088015160408701906117fb565b606087015185820360608701526138a5565b9461395d60808201516080860190611808565b61396f60a082015160a0860190611808565b01519101906117fb565b90565b6139919160208201915f8184039101526138f5565b90565b61399e6040610273565b90565b6139a9613994565b90602080836139b6613388565b8152016139c1613384565b81525050565b6139cf6139a1565b90565b6139dc90516109fc565b90565b6139e9905161049b565b90565b6139f7610160610273565b90565b90613a049061043e565b9052565b90613a12906109fc565b9052565b90613a209061049b565b9052565b90613a2e90610722565b9052565b90613a3c90610399565b9052565b52565b613a4d6040610273565b90565b52565b52565b90613a6a613a6383611742565b8092612d91565b9081613a7b602083028401946117e8565b925f915b838310613a8e57505050505090565b90919293946020613ab0613aaa83856001950387528951611982565b9761198f565b9301930191939290613a7f565b613b9491610140613b826101608301613adc5f8601515f8601906117ee565b613aee60208601516020860190612e32565b613b0060408601516040860190611815565b613b1260608601516060860190611815565b613b2460808601516080860190611815565b613b3660a086015160a0860190611815565b613b4860c086015160c0860190611808565b613b5a60e086015160e0860190611808565b613b6e6101008601516101008601906117fb565b61012085015184820361012086015261182b565b9201519061014081840391015261182b565b90565b613bc5916020613bb4604083015f8501518482035f860152613a56565b920151906020818403910152613abd565b90565b613bdd9160208201915f818403910152613b97565b90565b90613be9611e81565b50613bf2611a28565b50613bfb611e81565b613c03611a28565b613c0b611a28565b90613c1b5f602087015101611aa8565b613c2d613c275f611ad4565b9161043e565b145f14613faf57505050613c3f6139c7565b50613c72613c6d600c60e060a0613c625f880151613c5c5f611749565b90613425565b5101510151906103e9565b613459565b613c84613c7e5f611ad4565b9161043e565b145f14613e9857613cc360e060a0613ca85f860151613ca25f611749565b90613425565b51015101516020613cb882610344565b818301019101611a8a565b91613cef5f60e060a0613ce283860151613cdc85611749565b90613425565b5101510151945b01611b1b565b92610120602083015101515b82602001515f01613d0b90611aa8565b918360200151602001613d1d906139d2565b918460200151604001613d2f906139df565b8560200151606001613d40906139df565b8660200151608001613d51906139df565b876020015160a001613d62906139df565b90886020015160c001613d7490611b68565b92896020015160e001613d8690611b68565b948a6020015161010001613d9990611b1b565b969798613da46139ec565b9a5f8c0190613db2916139fa565b60208b0190613dc091613a08565b60408a0190613dce91613a16565b6060890190613ddc91613a16565b6080880190613dea91613a16565b60a0870190613df891613a16565b60c0860190613e0691613a24565b60e0850190613e1491613a24565b610100840190613e2391613a32565b610120830190613e3291613a40565b610140820190613e4191613a40565b905f015190613e4e613a43565b915f830190613e5c91613a50565b6020820190613e6a91613a53565b613e72610212565b80916020820190613e8291613bc8565b602082018103825203613e95908261024a565b90565b613eca613ec5600c60e060a0613eba5f880151613eb45f611749565b90613425565b5101510151906103e9565b613459565b613edd613ed7600161332f565b9161043e565b1480613f6f575b5f14613f4c57613f2260e060c0613f075f860151613f015f611749565b90613425565b51015101516020613f1782610344565b818301019101611a8a565b91613cef5f60e060c0613f4183860151613f3b85611749565b90613425565b510151015194613ce9565b613f54610212565b6382077b6960e01b815280613f6b60048201610587565b0390fd5b50613f9760e060c0613f8d5f860151613f875f611749565b90613425565b5101510151610344565b613fa9613fa35f611749565b91610399565b11613ee4565b613fc05f6020879597015101611aa8565b613fd3613fcd600161332f565b9161043e565b145f146141a25750613fe36136f2565b50614016614011600c60e060a06140065f8801516140005f611749565b90613425565b5101510151906103e9565b613459565b6140286140225f611ad4565b9161043e565b145f146140ab5760046140a561406c60e060a06140515f88015161404b5f611749565b90613425565b5101510151602061406182610344565b818301019101613846565b5b61409661407d6374a736b961359c565b91614086610212565b948593602085019081520161397c565b6020820181038252038261024a565b5b613cfb565b6140dd6140d8600c60e060a06140cd5f8801516140c75f611749565b90613425565b5101510151906103e9565b613459565b6140f06140ea600161332f565b9161043e565b1480614162575b5f1461413f5760046140a561413a60e060c061411f5f8801516141195f611749565b90613425565b5101510151602061412f82610344565b818301019101613846565b61406d565b614147610212565b6382077b6960e01b81528061415e60048201610587565b0390fd5b5061418a60e060c06141805f86015161417a5f611749565b90613425565b5101510151610344565b61419c6141965f611749565b91610399565b116140f7565b6141b15f602085015101611aa8565b6141c46141be600261334e565b9161043e565b145f1461439357506141d4613406565b50614207614202600c60e060a06141f75f8801516141f15f611749565b90613425565b5101510151906103e9565b613459565b6142196142135f611ad4565b9161043e565b145f1461429c57600461429661425d60e060a06142425f88015161423c5f611749565b90613425565b5101510151602061425282610344565b818301019101613568565b5b61428761426e63bc33e62061359c565b91614277610212565b9485936020850190815201613668565b6020820181038252038261024a565b5b6140a6565b6142ce6142c9600c60e060a06142be5f8801516142b85f611749565b90613425565b5101510151906103e9565b613459565b6142e16142db600161332f565b9161043e565b1480614353575b5f1461433057600461429661432b60e060c06143105f88015161430a5f611749565b90613425565b5101510151602061432082610344565b818301019101613568565b61425e565b614338610212565b6382077b6960e01b81528061434f60048201610587565b0390fd5b5061437b60e060c06143715f86015161436b5f611749565b90613425565b5101510151610344565b61438d6143875f611749565b91610399565b116142e8565b6143a45f6020859795015101611aa8565b6143b76143b160fe611af3565b9161043e565b145f14614409575050506143ee5f6143e86101406020860151015160206143dd82610344565b818301019101611a8a565b01611b1b565b9161014060208201510151610120602083015101515b614297565b6144185f602087015101611aa8565b61442b61442560ff611ab8565b9161043e565b145f1461447d575050506144625f61445c61014060208601510151602061445182610344565b818301019101611a8a565b01611b1b565b9161014060208201510151610120602083015101515b614404565b939193614478565b61448d610212565b637fca089960e11b8152806144a460048201610587565b0390fd5b6144b1906105fd565b90565b6144c86144c36144cd926105ef565b6105fa565b610399565b90565b90565b6144df6144e491610902565b6144d0565b9052565b60e81b90565b6144f7906144e8565b90565b61450661450b916109fc565b6144ee565b9052565b60c01b90565b61451e9061450f565b90565b61452d6145329161049b565b614515565b9052565b60089361456960206145789897956145616001866145596003986145719a6144d3565b01809261039f565b0180926144fa565b018092614521565b0190610370565b90565b926145b89192946145a061459b6145c796614594611a28565b50956144a8565b6144b4565b959190916145ac610212565b96879560208701614536565b6020820181038252038261024a565b90565b6145d2612985565b6145eb6145e56145e0614cef565b610722565b91610722565b036145f257565b61461b6145fd614cef565b614605610212565b91829163118cdaa760e01b83526004830161073b565b0390fd5b6146296002611795565b614634826002612cfb565b906146686146627f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093612cec565b91612cec565b91614671610212565b8061467b81610587565b0390a3565b61468c614691916116be565b610a90565b90565b61469e9054614680565b90565b6146b56146b06146ba9261334b565b6105fa565b610399565b90565b6146c760026146a1565b90565b906146d65f1991612c2c565b9181191691161790565b90565b906146f86146f36146ff92610f76565b6146e0565b82546146ca565b9055565b61470d6003614694565b61472661472061471b6146bd565b610399565b91610399565b1461473f5761473d6147366146bd565b60036146e3565b565b614747610212565b633ee5aeb560e01b81528061475e60048201610587565b0390fd5b61477661477161477b9261332c565b6105fa565b610399565b90565b6147886001614762565b90565b61479d61479661477e565b60036146e3565b565b92906147d5916147c86147e396946147be60808801945f8901906108a7565b60208701906109a5565b848203604086015261162d565b91606081840391015261162d565b90565b92614828602093946147f6611e81565b5061483361480b6148065f6116d7565b610625565b936385fdd54292959761481c610212565b988997889687966117ca565b86526004860161479f565b03915afa908115614877575f91614849575b5090565b61486a915060203d8111614870575b614862818361024a565b810190612222565b5f614845565b503d614858565b611a14565b91906040838203126148b6576148af906148966040610273565b936148a3825f8301611a3c565b5f860152602001611a3c565b6020830152565b6111fe565b906040828203126148d4576148d1915f0161487c565b90565b61021c565b6148ed6148e86148f292610399565b6105fa565b6105ef565b90565b5f7f496e76616c696420636f6e747261637400000000000000000000000000000000910152565b6149296010602092612585565b614932816148f5565b0190565b61494b9060208101905f81830391015261491c565b90565b1561495557565b61495d610212565b62461bcd60e51b81528061497360048201614936565b0390fd5b5060406149cd91949392946149a66149a161499c614997600a85906106c2565b611795565b6117ae565b6117ba565b6149c26342f1ec696149b6610212565b958694859384936117ca565b8352600483016109b2565b03915afa938415614cea57614a15614a0f614a0a614a056149ff6020614a249a614a1b985f91614cbc575b5001611b68565b946148d9565b610619565b610722565b91610722565b1461494e565b908101906115ef565b614a305f820151611742565b614a42614a3c5f611749565b91610399565b03614c0c575b614a50611a24565b50614a59611a28565b505f80614a7f610140602085015101516020614a7482610344565b818301019101611a8a565b614a8e82602086015101611aa8565b614aa1614a9b60ff611ab8565b9161043e565b148214614b3557614acf82614ac8614ac3600a614abd46611765565b906106c2565b611795565b9201611b1b565b6101206020860151015190602082019151925af1906101206020614af1611b3f565b935b93920151015191614b307fc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb39807944693614b27610212565b93849384611ba2565b0390a1565b614b4482602086015101611aa8565b614b56614b5084611ad4565b9161043e565b148015614be3575b8214614ba757614b7b82614b7460208401611b68565b9201611b1b565b614b83610212565b9081614b8e81611b84565b03925af1906101206020614ba0611b3f565b935b614af3565b614bbb82614bb430611b0f565b9201611b1b565b6101206020860151015190602082019151925af1906101206020614bdd611b3f565b93614ba2565b50614bf382602086015101611aa8565b614c06614c0060fe611af3565b9161043e565b14614b5e565b614c38614c33614c2e614c29600a614c2346611765565b906106c2565b611795565b6117ae565b6117ba565b6377c2719c5f830151823b15614cb757614c7192614c665f8094614c5a610212565b968795869485936117ca565b8352600483016119fc565b03925af18015614cb257614c86575b50614a48565b614ca5905f3d8111614cab575b614c9d818361024a565b8101906117d0565b5f614c80565b503d614c93565b611a14565b6117c6565b614cdd915060403d8111614ce3575b614cd5818361024a565b8101906148bb565b5f6149f8565b503d614ccb565b611a14565b614cf7612981565b50339056fea264697066735822122082dd4bdcdb5d8532830715b5cf2cc7f6a20f0de585850c284880478b5c7260d864736f6c63430008180033",
}

// SyncRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SyncRouterMetaData.ABI instead.
var SyncRouterABI = SyncRouterMetaData.ABI

// SyncRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SyncRouterMetaData.Bin instead.
var SyncRouterBin = SyncRouterMetaData.Bin

// DeploySyncRouter deploys a new Ethereum contract, binding an instance of SyncRouter to it.
func DeploySyncRouter(auth *bind.TransactOpts, backend bind.ContractBackend, _vizingPad common.Address, _WETH common.Address, _Hook common.Address) (common.Address, *types.Transaction, *SyncRouter, error) {
	parsed, err := SyncRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SyncRouterBin), backend, _vizingPad, _WETH, _Hook)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SyncRouter{SyncRouterCaller: SyncRouterCaller{contract: contract}, SyncRouterTransactor: SyncRouterTransactor{contract: contract}, SyncRouterFilterer: SyncRouterFilterer{contract: contract}}, nil
}

// SyncRouter is an auto generated Go binding around an Ethereum contract.
type SyncRouter struct {
	SyncRouterCaller     // Read-only binding to the contract
	SyncRouterTransactor // Write-only binding to the contract
	SyncRouterFilterer   // Log filterer for contract events
}

// SyncRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SyncRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyncRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SyncRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyncRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SyncRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyncRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SyncRouterSession struct {
	Contract     *SyncRouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SyncRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SyncRouterCallerSession struct {
	Contract *SyncRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SyncRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SyncRouterTransactorSession struct {
	Contract     *SyncRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SyncRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SyncRouterRaw struct {
	Contract *SyncRouter // Generic contract binding to access the raw methods on
}

// SyncRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SyncRouterCallerRaw struct {
	Contract *SyncRouterCaller // Generic read-only contract binding to access the raw methods on
}

// SyncRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SyncRouterTransactorRaw struct {
	Contract *SyncRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSyncRouter creates a new instance of SyncRouter, bound to a specific deployed contract.
func NewSyncRouter(address common.Address, backend bind.ContractBackend) (*SyncRouter, error) {
	contract, err := bindSyncRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SyncRouter{SyncRouterCaller: SyncRouterCaller{contract: contract}, SyncRouterTransactor: SyncRouterTransactor{contract: contract}, SyncRouterFilterer: SyncRouterFilterer{contract: contract}}, nil
}

// NewSyncRouterCaller creates a new read-only instance of SyncRouter, bound to a specific deployed contract.
func NewSyncRouterCaller(address common.Address, caller bind.ContractCaller) (*SyncRouterCaller, error) {
	contract, err := bindSyncRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SyncRouterCaller{contract: contract}, nil
}

// NewSyncRouterTransactor creates a new write-only instance of SyncRouter, bound to a specific deployed contract.
func NewSyncRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*SyncRouterTransactor, error) {
	contract, err := bindSyncRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SyncRouterTransactor{contract: contract}, nil
}

// NewSyncRouterFilterer creates a new log filterer instance of SyncRouter, bound to a specific deployed contract.
func NewSyncRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*SyncRouterFilterer, error) {
	contract, err := bindSyncRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SyncRouterFilterer{contract: contract}, nil
}

// bindSyncRouter binds a generic wrapper to an already deployed contract.
func bindSyncRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SyncRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyncRouter *SyncRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyncRouter.Contract.SyncRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyncRouter *SyncRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyncRouter.Contract.SyncRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyncRouter *SyncRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyncRouter.Contract.SyncRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyncRouter *SyncRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyncRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyncRouter *SyncRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyncRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyncRouter *SyncRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyncRouter.Contract.contract.Transact(opts, method, params...)
}

// DataExcecuteNumber is a free data retrieval call binding the contract method 0x002232e3.
//
// Solidity: function DataExcecuteNumber(bytes ) view returns(uint8)
func (_SyncRouter *SyncRouterCaller) DataExcecuteNumber(opts *bind.CallOpts, arg0 []byte) (uint8, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "DataExcecuteNumber", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DataExcecuteNumber is a free data retrieval call binding the contract method 0x002232e3.
//
// Solidity: function DataExcecuteNumber(bytes ) view returns(uint8)
func (_SyncRouter *SyncRouterSession) DataExcecuteNumber(arg0 []byte) (uint8, error) {
	return _SyncRouter.Contract.DataExcecuteNumber(&_SyncRouter.CallOpts, arg0)
}

// DataExcecuteNumber is a free data retrieval call binding the contract method 0x002232e3.
//
// Solidity: function DataExcecuteNumber(bytes ) view returns(uint8)
func (_SyncRouter *SyncRouterCallerSession) DataExcecuteNumber(arg0 []byte) (uint8, error) {
	return _SyncRouter.Contract.DataExcecuteNumber(&_SyncRouter.CallOpts, arg0)
}

// Hook is a free data retrieval call binding the contract method 0x8ea23ddf.
//
// Solidity: function Hook() view returns(address)
func (_SyncRouter *SyncRouterCaller) Hook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "Hook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hook is a free data retrieval call binding the contract method 0x8ea23ddf.
//
// Solidity: function Hook() view returns(address)
func (_SyncRouter *SyncRouterSession) Hook() (common.Address, error) {
	return _SyncRouter.Contract.Hook(&_SyncRouter.CallOpts)
}

// Hook is a free data retrieval call binding the contract method 0x8ea23ddf.
//
// Solidity: function Hook() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) Hook() (common.Address, error) {
	return _SyncRouter.Contract.Hook(&_SyncRouter.CallOpts)
}

// LandingPad is a free data retrieval call binding the contract method 0x0e82845d.
//
// Solidity: function LandingPad() view returns(address)
func (_SyncRouter *SyncRouterCaller) LandingPad(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "LandingPad")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LandingPad is a free data retrieval call binding the contract method 0x0e82845d.
//
// Solidity: function LandingPad() view returns(address)
func (_SyncRouter *SyncRouterSession) LandingPad() (common.Address, error) {
	return _SyncRouter.Contract.LandingPad(&_SyncRouter.CallOpts)
}

// LandingPad is a free data retrieval call binding the contract method 0x0e82845d.
//
// Solidity: function LandingPad() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) LandingPad() (common.Address, error) {
	return _SyncRouter.Contract.LandingPad(&_SyncRouter.CallOpts)
}

// LaunchPad is a free data retrieval call binding the contract method 0xe0b838e9.
//
// Solidity: function LaunchPad() view returns(address)
func (_SyncRouter *SyncRouterCaller) LaunchPad(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "LaunchPad")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LaunchPad is a free data retrieval call binding the contract method 0xe0b838e9.
//
// Solidity: function LaunchPad() view returns(address)
func (_SyncRouter *SyncRouterSession) LaunchPad() (common.Address, error) {
	return _SyncRouter.Contract.LaunchPad(&_SyncRouter.CallOpts)
}

// LaunchPad is a free data retrieval call binding the contract method 0xe0b838e9.
//
// Solidity: function LaunchPad() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) LaunchPad() (common.Address, error) {
	return _SyncRouter.Contract.LaunchPad(&_SyncRouter.CallOpts)
}

// LockWay is a free data retrieval call binding the contract method 0xde4b1d6d.
//
// Solidity: function LockWay(uint256 ) view returns(bytes1)
func (_SyncRouter *SyncRouterCaller) LockWay(opts *bind.CallOpts, arg0 *big.Int) ([1]byte, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "LockWay", arg0)

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// LockWay is a free data retrieval call binding the contract method 0xde4b1d6d.
//
// Solidity: function LockWay(uint256 ) view returns(bytes1)
func (_SyncRouter *SyncRouterSession) LockWay(arg0 *big.Int) ([1]byte, error) {
	return _SyncRouter.Contract.LockWay(&_SyncRouter.CallOpts, arg0)
}

// LockWay is a free data retrieval call binding the contract method 0xde4b1d6d.
//
// Solidity: function LockWay(uint256 ) view returns(bytes1)
func (_SyncRouter *SyncRouterCallerSession) LockWay(arg0 *big.Int) ([1]byte, error) {
	return _SyncRouter.Contract.LockWay(&_SyncRouter.CallOpts, arg0)
}

// MirrorEntryPoint is a free data retrieval call binding the contract method 0x1490a358.
//
// Solidity: function MirrorEntryPoint(uint64 ) view returns(address)
func (_SyncRouter *SyncRouterCaller) MirrorEntryPoint(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "MirrorEntryPoint", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MirrorEntryPoint is a free data retrieval call binding the contract method 0x1490a358.
//
// Solidity: function MirrorEntryPoint(uint64 ) view returns(address)
func (_SyncRouter *SyncRouterSession) MirrorEntryPoint(arg0 uint64) (common.Address, error) {
	return _SyncRouter.Contract.MirrorEntryPoint(&_SyncRouter.CallOpts, arg0)
}

// MirrorEntryPoint is a free data retrieval call binding the contract method 0x1490a358.
//
// Solidity: function MirrorEntryPoint(uint64 ) view returns(address)
func (_SyncRouter *SyncRouterCallerSession) MirrorEntryPoint(arg0 uint64) (common.Address, error) {
	return _SyncRouter.Contract.MirrorEntryPoint(&_SyncRouter.CallOpts, arg0)
}

// OrderId is a free data retrieval call binding the contract method 0x60b43029.
//
// Solidity: function OrderId() view returns(uint256)
func (_SyncRouter *SyncRouterCaller) OrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "OrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderId is a free data retrieval call binding the contract method 0x60b43029.
//
// Solidity: function OrderId() view returns(uint256)
func (_SyncRouter *SyncRouterSession) OrderId() (*big.Int, error) {
	return _SyncRouter.Contract.OrderId(&_SyncRouter.CallOpts)
}

// OrderId is a free data retrieval call binding the contract method 0x60b43029.
//
// Solidity: function OrderId() view returns(uint256)
func (_SyncRouter *SyncRouterCallerSession) OrderId() (*big.Int, error) {
	return _SyncRouter.Contract.OrderId(&_SyncRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_SyncRouter *SyncRouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_SyncRouter *SyncRouterSession) WETH() (common.Address, error) {
	return _SyncRouter.Contract.WETH(&_SyncRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) WETH() (common.Address, error) {
	return _SyncRouter.Contract.WETH(&_SyncRouter.CallOpts)
}

// DefaultBridgeMode is a free data retrieval call binding the contract method 0x45636279.
//
// Solidity: function defaultBridgeMode() view returns(bytes1)
func (_SyncRouter *SyncRouterCaller) DefaultBridgeMode(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "defaultBridgeMode")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DefaultBridgeMode is a free data retrieval call binding the contract method 0x45636279.
//
// Solidity: function defaultBridgeMode() view returns(bytes1)
func (_SyncRouter *SyncRouterSession) DefaultBridgeMode() ([1]byte, error) {
	return _SyncRouter.Contract.DefaultBridgeMode(&_SyncRouter.CallOpts)
}

// DefaultBridgeMode is a free data retrieval call binding the contract method 0x45636279.
//
// Solidity: function defaultBridgeMode() view returns(bytes1)
func (_SyncRouter *SyncRouterCallerSession) DefaultBridgeMode() ([1]byte, error) {
	return _SyncRouter.Contract.DefaultBridgeMode(&_SyncRouter.CallOpts)
}

// DefaultGasPrice is a free data retrieval call binding the contract method 0xe7b4294c.
//
// Solidity: function defaultGasPrice() view returns(uint64)
func (_SyncRouter *SyncRouterCaller) DefaultGasPrice(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "defaultGasPrice")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DefaultGasPrice is a free data retrieval call binding the contract method 0xe7b4294c.
//
// Solidity: function defaultGasPrice() view returns(uint64)
func (_SyncRouter *SyncRouterSession) DefaultGasPrice() (uint64, error) {
	return _SyncRouter.Contract.DefaultGasPrice(&_SyncRouter.CallOpts)
}

// DefaultGasPrice is a free data retrieval call binding the contract method 0xe7b4294c.
//
// Solidity: function defaultGasPrice() view returns(uint64)
func (_SyncRouter *SyncRouterCallerSession) DefaultGasPrice() (uint64, error) {
	return _SyncRouter.Contract.DefaultGasPrice(&_SyncRouter.CallOpts)
}

// DefaultGaslimit is a free data retrieval call binding the contract method 0x99e581fa.
//
// Solidity: function defaultGaslimit() view returns(uint24)
func (_SyncRouter *SyncRouterCaller) DefaultGaslimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "defaultGaslimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultGaslimit is a free data retrieval call binding the contract method 0x99e581fa.
//
// Solidity: function defaultGaslimit() view returns(uint24)
func (_SyncRouter *SyncRouterSession) DefaultGaslimit() (*big.Int, error) {
	return _SyncRouter.Contract.DefaultGaslimit(&_SyncRouter.CallOpts)
}

// DefaultGaslimit is a free data retrieval call binding the contract method 0x99e581fa.
//
// Solidity: function defaultGaslimit() view returns(uint24)
func (_SyncRouter *SyncRouterCallerSession) DefaultGaslimit() (*big.Int, error) {
	return _SyncRouter.Contract.DefaultGaslimit(&_SyncRouter.CallOpts)
}

// EstimateVizingGasFee is a free data retrieval call binding the contract method 0x76c81312.
//
// Solidity: function estimateVizingGasFee(uint256 value, uint64 destChainid, bytes additionParams, bytes message) view returns(uint256 vizingGasFee)
func (_SyncRouter *SyncRouterCaller) EstimateVizingGasFee(opts *bind.CallOpts, value *big.Int, destChainid uint64, additionParams []byte, message []byte) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "estimateVizingGasFee", value, destChainid, additionParams, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateVizingGasFee is a free data retrieval call binding the contract method 0x76c81312.
//
// Solidity: function estimateVizingGasFee(uint256 value, uint64 destChainid, bytes additionParams, bytes message) view returns(uint256 vizingGasFee)
func (_SyncRouter *SyncRouterSession) EstimateVizingGasFee(value *big.Int, destChainid uint64, additionParams []byte, message []byte) (*big.Int, error) {
	return _SyncRouter.Contract.EstimateVizingGasFee(&_SyncRouter.CallOpts, value, destChainid, additionParams, message)
}

// EstimateVizingGasFee is a free data retrieval call binding the contract method 0x76c81312.
//
// Solidity: function estimateVizingGasFee(uint256 value, uint64 destChainid, bytes additionParams, bytes message) view returns(uint256 vizingGasFee)
func (_SyncRouter *SyncRouterCallerSession) EstimateVizingGasFee(value *big.Int, destChainid uint64, additionParams []byte, message []byte) (*big.Int, error) {
	return _SyncRouter.Contract.EstimateVizingGasFee(&_SyncRouter.CallOpts, value, destChainid, additionParams, message)
}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0x3eba5864.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) view returns(uint256)
func (_SyncRouter *SyncRouterCaller) FetchOmniMessageFee(opts *bind.CallOpts, destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "fetchOmniMessageFee", destChainId, destContract, destChainUsedFee, userOperations)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0x3eba5864.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) view returns(uint256)
func (_SyncRouter *SyncRouterSession) FetchOmniMessageFee(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*big.Int, error) {
	return _SyncRouter.Contract.FetchOmniMessageFee(&_SyncRouter.CallOpts, destChainId, destContract, destChainUsedFee, userOperations)
}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0x3eba5864.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) view returns(uint256)
func (_SyncRouter *SyncRouterCallerSession) FetchOmniMessageFee(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*big.Int, error) {
	return _SyncRouter.Contract.FetchOmniMessageFee(&_SyncRouter.CallOpts, destChainId, destContract, destChainUsedFee, userOperations)
}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0xafb3e17d.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterCaller) FetchUserOmniMessageFee(opts *bind.CallOpts, params BaseStructCrossMessageParams) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "fetchUserOmniMessageFee", params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0xafb3e17d.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterSession) FetchUserOmniMessageFee(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _SyncRouter.Contract.FetchUserOmniMessageFee(&_SyncRouter.CallOpts, params)
}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0xafb3e17d.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterCallerSession) FetchUserOmniMessageFee(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _SyncRouter.Contract.FetchUserOmniMessageFee(&_SyncRouter.CallOpts, params)
}

// GetUserOmniEncodeMessage is a free data retrieval call binding the contract method 0xf329f52f.
//
// Solidity: function getUserOmniEncodeMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) cmp) view returns(uint256, bytes)
func (_SyncRouter *SyncRouterCaller) GetUserOmniEncodeMessage(opts *bind.CallOpts, cmp BaseStructCrossMessageParams) (*big.Int, []byte, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "getUserOmniEncodeMessage", cmp)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetUserOmniEncodeMessage is a free data retrieval call binding the contract method 0xf329f52f.
//
// Solidity: function getUserOmniEncodeMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) cmp) view returns(uint256, bytes)
func (_SyncRouter *SyncRouterSession) GetUserOmniEncodeMessage(cmp BaseStructCrossMessageParams) (*big.Int, []byte, error) {
	return _SyncRouter.Contract.GetUserOmniEncodeMessage(&_SyncRouter.CallOpts, cmp)
}

// GetUserOmniEncodeMessage is a free data retrieval call binding the contract method 0xf329f52f.
//
// Solidity: function getUserOmniEncodeMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) cmp) view returns(uint256, bytes)
func (_SyncRouter *SyncRouterCallerSession) GetUserOmniEncodeMessage(cmp BaseStructCrossMessageParams) (*big.Int, []byte, error) {
	return _SyncRouter.Contract.GetUserOmniEncodeMessage(&_SyncRouter.CallOpts, cmp)
}

// MaxArrivalTime is a free data retrieval call binding the contract method 0xde8aeda0.
//
// Solidity: function maxArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterCaller) MaxArrivalTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "maxArrivalTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxArrivalTime is a free data retrieval call binding the contract method 0xde8aeda0.
//
// Solidity: function maxArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterSession) MaxArrivalTime() (uint64, error) {
	return _SyncRouter.Contract.MaxArrivalTime(&_SyncRouter.CallOpts)
}

// MaxArrivalTime is a free data retrieval call binding the contract method 0xde8aeda0.
//
// Solidity: function maxArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterCallerSession) MaxArrivalTime() (uint64, error) {
	return _SyncRouter.Contract.MaxArrivalTime(&_SyncRouter.CallOpts)
}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterCaller) MaxGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "maxGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterSession) MaxGasLimit() (*big.Int, error) {
	return _SyncRouter.Contract.MaxGasLimit(&_SyncRouter.CallOpts)
}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterCallerSession) MaxGasLimit() (*big.Int, error) {
	return _SyncRouter.Contract.MaxGasLimit(&_SyncRouter.CallOpts)
}

// MinArrivalTime is a free data retrieval call binding the contract method 0x5ad3ad06.
//
// Solidity: function minArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterCaller) MinArrivalTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "minArrivalTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinArrivalTime is a free data retrieval call binding the contract method 0x5ad3ad06.
//
// Solidity: function minArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterSession) MinArrivalTime() (uint64, error) {
	return _SyncRouter.Contract.MinArrivalTime(&_SyncRouter.CallOpts)
}

// MinArrivalTime is a free data retrieval call binding the contract method 0x5ad3ad06.
//
// Solidity: function minArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterCallerSession) MinArrivalTime() (uint64, error) {
	return _SyncRouter.Contract.MinArrivalTime(&_SyncRouter.CallOpts)
}

// MinGasLimit is a free data retrieval call binding the contract method 0x5aeb4d77.
//
// Solidity: function minGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterCaller) MinGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "minGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinGasLimit is a free data retrieval call binding the contract method 0x5aeb4d77.
//
// Solidity: function minGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterSession) MinGasLimit() (*big.Int, error) {
	return _SyncRouter.Contract.MinGasLimit(&_SyncRouter.CallOpts)
}

// MinGasLimit is a free data retrieval call binding the contract method 0x5aeb4d77.
//
// Solidity: function minGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterCallerSession) MinGasLimit() (*big.Int, error) {
	return _SyncRouter.Contract.MinGasLimit(&_SyncRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyncRouter *SyncRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyncRouter *SyncRouterSession) Owner() (common.Address, error) {
	return _SyncRouter.Contract.Owner(&_SyncRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) Owner() (common.Address, error) {
	return _SyncRouter.Contract.Owner(&_SyncRouter.CallOpts)
}

// SelectedRelayer is a free data retrieval call binding the contract method 0xb0cfd4d2.
//
// Solidity: function selectedRelayer() view returns(address)
func (_SyncRouter *SyncRouterCaller) SelectedRelayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "selectedRelayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SelectedRelayer is a free data retrieval call binding the contract method 0xb0cfd4d2.
//
// Solidity: function selectedRelayer() view returns(address)
func (_SyncRouter *SyncRouterSession) SelectedRelayer() (common.Address, error) {
	return _SyncRouter.Contract.SelectedRelayer(&_SyncRouter.CallOpts)
}

// SelectedRelayer is a free data retrieval call binding the contract method 0xb0cfd4d2.
//
// Solidity: function selectedRelayer() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) SelectedRelayer() (common.Address, error) {
	return _SyncRouter.Contract.SelectedRelayer(&_SyncRouter.CallOpts)
}

// ThisRelayer is a free data retrieval call binding the contract method 0xc5f378d8.
//
// Solidity: function thisRelayer() view returns(address)
func (_SyncRouter *SyncRouterCaller) ThisRelayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "thisRelayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ThisRelayer is a free data retrieval call binding the contract method 0xc5f378d8.
//
// Solidity: function thisRelayer() view returns(address)
func (_SyncRouter *SyncRouterSession) ThisRelayer() (common.Address, error) {
	return _SyncRouter.Contract.ThisRelayer(&_SyncRouter.CallOpts)
}

// ThisRelayer is a free data retrieval call binding the contract method 0xc5f378d8.
//
// Solidity: function thisRelayer() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) ThisRelayer() (common.Address, error) {
	return _SyncRouter.Contract.ThisRelayer(&_SyncRouter.CallOpts)
}

// ChangeDefaultSet is a paid mutator transaction binding the contract method 0xa834a900.
//
// Solidity: function changeDefaultSet(uint24 newGaslimit, uint64 newGasPrice, address newRelayer) returns()
func (_SyncRouter *SyncRouterTransactor) ChangeDefaultSet(opts *bind.TransactOpts, newGaslimit *big.Int, newGasPrice uint64, newRelayer common.Address) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "changeDefaultSet", newGaslimit, newGasPrice, newRelayer)
}

// ChangeDefaultSet is a paid mutator transaction binding the contract method 0xa834a900.
//
// Solidity: function changeDefaultSet(uint24 newGaslimit, uint64 newGasPrice, address newRelayer) returns()
func (_SyncRouter *SyncRouterSession) ChangeDefaultSet(newGaslimit *big.Int, newGasPrice uint64, newRelayer common.Address) (*types.Transaction, error) {
	return _SyncRouter.Contract.ChangeDefaultSet(&_SyncRouter.TransactOpts, newGaslimit, newGasPrice, newRelayer)
}

// ChangeDefaultSet is a paid mutator transaction binding the contract method 0xa834a900.
//
// Solidity: function changeDefaultSet(uint24 newGaslimit, uint64 newGasPrice, address newRelayer) returns()
func (_SyncRouter *SyncRouterTransactorSession) ChangeDefaultSet(newGaslimit *big.Int, newGasPrice uint64, newRelayer common.Address) (*types.Transaction, error) {
	return _SyncRouter.Contract.ChangeDefaultSet(&_SyncRouter.TransactOpts, newGaslimit, newGasPrice, newRelayer)
}

// ReceiveExtraMessage is a paid mutator transaction binding the contract method 0xf3148925.
//
// Solidity: function receiveExtraMessage(bytes32 messageId, uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterTransactor) ReceiveExtraMessage(opts *bind.TransactOpts, messageId [32]byte, srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "receiveExtraMessage", messageId, srcChainId, srcContract, message)
}

// ReceiveExtraMessage is a paid mutator transaction binding the contract method 0xf3148925.
//
// Solidity: function receiveExtraMessage(bytes32 messageId, uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterSession) ReceiveExtraMessage(messageId [32]byte, srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.ReceiveExtraMessage(&_SyncRouter.TransactOpts, messageId, srcChainId, srcContract, message)
}

// ReceiveExtraMessage is a paid mutator transaction binding the contract method 0xf3148925.
//
// Solidity: function receiveExtraMessage(bytes32 messageId, uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterTransactorSession) ReceiveExtraMessage(messageId [32]byte, srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.ReceiveExtraMessage(&_SyncRouter.TransactOpts, messageId, srcChainId, srcContract, message)
}

// ReceiveStandardMessage is a paid mutator transaction binding the contract method 0x0073b555.
//
// Solidity: function receiveStandardMessage(uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterTransactor) ReceiveStandardMessage(opts *bind.TransactOpts, srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "receiveStandardMessage", srcChainId, srcContract, message)
}

// ReceiveStandardMessage is a paid mutator transaction binding the contract method 0x0073b555.
//
// Solidity: function receiveStandardMessage(uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterSession) ReceiveStandardMessage(srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.ReceiveStandardMessage(&_SyncRouter.TransactOpts, srcChainId, srcContract, message)
}

// ReceiveStandardMessage is a paid mutator transaction binding the contract method 0x0073b555.
//
// Solidity: function receiveStandardMessage(uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterTransactorSession) ReceiveStandardMessage(srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.ReceiveStandardMessage(&_SyncRouter.TransactOpts, srcChainId, srcContract, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyncRouter *SyncRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyncRouter *SyncRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _SyncRouter.Contract.RenounceOwnership(&_SyncRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyncRouter *SyncRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SyncRouter.Contract.RenounceOwnership(&_SyncRouter.TransactOpts)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0x93afae93.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainExecuteUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) payable returns()
func (_SyncRouter *SyncRouterTransactor) SendOmniMessage(opts *bind.TransactOpts, destChainId uint64, destContract common.Address, destChainExecuteUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "sendOmniMessage", destChainId, destContract, destChainExecuteUsedFee, userOperations)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0x93afae93.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainExecuteUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) payable returns()
func (_SyncRouter *SyncRouterSession) SendOmniMessage(destChainId uint64, destContract common.Address, destChainExecuteUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendOmniMessage(&_SyncRouter.TransactOpts, destChainId, destContract, destChainExecuteUsedFee, userOperations)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0x93afae93.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainExecuteUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) payable returns()
func (_SyncRouter *SyncRouterTransactorSession) SendOmniMessage(destChainId uint64, destContract common.Address, destChainExecuteUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendOmniMessage(&_SyncRouter.TransactOpts, destChainId, destContract, destChainExecuteUsedFee, userOperations)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x72afdea1.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) cmp) payable returns()
func (_SyncRouter *SyncRouterTransactor) SendUserOmniMessage(opts *bind.TransactOpts, cmp BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "sendUserOmniMessage", cmp)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x72afdea1.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) cmp) payable returns()
func (_SyncRouter *SyncRouterSession) SendUserOmniMessage(cmp BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendUserOmniMessage(&_SyncRouter.TransactOpts, cmp)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x72afdea1.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) cmp) payable returns()
func (_SyncRouter *SyncRouterTransactorSession) SendUserOmniMessage(cmp BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendUserOmniMessage(&_SyncRouter.TransactOpts, cmp)
}

// SetMirrorEntryPoint is a paid mutator transaction binding the contract method 0xd5f3e008.
//
// Solidity: function setMirrorEntryPoint(uint64 chainId, address entryPoint) returns()
func (_SyncRouter *SyncRouterTransactor) SetMirrorEntryPoint(opts *bind.TransactOpts, chainId uint64, entryPoint common.Address) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "setMirrorEntryPoint", chainId, entryPoint)
}

// SetMirrorEntryPoint is a paid mutator transaction binding the contract method 0xd5f3e008.
//
// Solidity: function setMirrorEntryPoint(uint64 chainId, address entryPoint) returns()
func (_SyncRouter *SyncRouterSession) SetMirrorEntryPoint(chainId uint64, entryPoint common.Address) (*types.Transaction, error) {
	return _SyncRouter.Contract.SetMirrorEntryPoint(&_SyncRouter.TransactOpts, chainId, entryPoint)
}

// SetMirrorEntryPoint is a paid mutator transaction binding the contract method 0xd5f3e008.
//
// Solidity: function setMirrorEntryPoint(uint64 chainId, address entryPoint) returns()
func (_SyncRouter *SyncRouterTransactorSession) SetMirrorEntryPoint(chainId uint64, entryPoint common.Address) (*types.Transaction, error) {
	return _SyncRouter.Contract.SetMirrorEntryPoint(&_SyncRouter.TransactOpts, chainId, entryPoint)
}

// TestReceiveMessage is a paid mutator transaction binding the contract method 0x397a567d.
//
// Solidity: function testReceiveMessage(bytes message) payable returns()
func (_SyncRouter *SyncRouterTransactor) TestReceiveMessage(opts *bind.TransactOpts, message []byte) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "testReceiveMessage", message)
}

// TestReceiveMessage is a paid mutator transaction binding the contract method 0x397a567d.
//
// Solidity: function testReceiveMessage(bytes message) payable returns()
func (_SyncRouter *SyncRouterSession) TestReceiveMessage(message []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.TestReceiveMessage(&_SyncRouter.TransactOpts, message)
}

// TestReceiveMessage is a paid mutator transaction binding the contract method 0x397a567d.
//
// Solidity: function testReceiveMessage(bytes message) payable returns()
func (_SyncRouter *SyncRouterTransactorSession) TestReceiveMessage(message []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.TestReceiveMessage(&_SyncRouter.TransactOpts, message)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyncRouter *SyncRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyncRouter *SyncRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SyncRouter.Contract.TransferOwnership(&_SyncRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyncRouter *SyncRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SyncRouter.Contract.TransferOwnership(&_SyncRouter.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SyncRouter *SyncRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyncRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SyncRouter *SyncRouterSession) Receive() (*types.Transaction, error) {
	return _SyncRouter.Contract.Receive(&_SyncRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SyncRouter *SyncRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _SyncRouter.Contract.Receive(&_SyncRouter.TransactOpts)
}

// SyncRouterAccountDeployedIterator is returned from FilterAccountDeployed and is used to iterate over the raw logs and unpacked data for AccountDeployed events raised by the SyncRouter contract.
type SyncRouterAccountDeployedIterator struct {
	Event *SyncRouterAccountDeployed // Event containing the contract specifics and raw log

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
func (it *SyncRouterAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterAccountDeployed)
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
		it.Event = new(SyncRouterAccountDeployed)
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
func (it *SyncRouterAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterAccountDeployed represents a AccountDeployed event raised by the SyncRouter contract.
type SyncRouterAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDeployed is a free log retrieval operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_SyncRouter *SyncRouterFilterer) FilterAccountDeployed(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*SyncRouterAccountDeployedIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterAccountDeployedIterator{contract: _SyncRouter.contract, event: "AccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAccountDeployed is a free log subscription operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_SyncRouter *SyncRouterFilterer) WatchAccountDeployed(opts *bind.WatchOpts, sink chan<- *SyncRouterAccountDeployed, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterAccountDeployed)
				if err := _SyncRouter.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
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
func (_SyncRouter *SyncRouterFilterer) ParseAccountDeployed(log types.Log) (*SyncRouterAccountDeployed, error) {
	event := new(SyncRouterAccountDeployed)
	if err := _SyncRouter.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterBeforeExecutionIterator is returned from FilterBeforeExecution and is used to iterate over the raw logs and unpacked data for BeforeExecution events raised by the SyncRouter contract.
type SyncRouterBeforeExecutionIterator struct {
	Event *SyncRouterBeforeExecution // Event containing the contract specifics and raw log

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
func (it *SyncRouterBeforeExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterBeforeExecution)
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
		it.Event = new(SyncRouterBeforeExecution)
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
func (it *SyncRouterBeforeExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterBeforeExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterBeforeExecution represents a BeforeExecution event raised by the SyncRouter contract.
type SyncRouterBeforeExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBeforeExecution is a free log retrieval operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_SyncRouter *SyncRouterFilterer) FilterBeforeExecution(opts *bind.FilterOpts) (*SyncRouterBeforeExecutionIterator, error) {

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return &SyncRouterBeforeExecutionIterator{contract: _SyncRouter.contract, event: "BeforeExecution", logs: logs, sub: sub}, nil
}

// WatchBeforeExecution is a free log subscription operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_SyncRouter *SyncRouterFilterer) WatchBeforeExecution(opts *bind.WatchOpts, sink chan<- *SyncRouterBeforeExecution) (event.Subscription, error) {

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterBeforeExecution)
				if err := _SyncRouter.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
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
func (_SyncRouter *SyncRouterFilterer) ParseBeforeExecution(log types.Log) (*SyncRouterBeforeExecution, error) {
	event := new(SyncRouterBeforeExecution)
	if err := _SyncRouter.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterDepositTicketAddedIterator is returned from FilterDepositTicketAdded and is used to iterate over the raw logs and unpacked data for DepositTicketAdded events raised by the SyncRouter contract.
type SyncRouterDepositTicketAddedIterator struct {
	Event *SyncRouterDepositTicketAdded // Event containing the contract specifics and raw log

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
func (it *SyncRouterDepositTicketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterDepositTicketAdded)
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
		it.Event = new(SyncRouterDepositTicketAdded)
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
func (it *SyncRouterDepositTicketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterDepositTicketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterDepositTicketAdded represents a DepositTicketAdded event raised by the SyncRouter contract.
type SyncRouterDepositTicketAdded struct {
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketAdded is a free log retrieval operation binding the contract event 0xc3907160f19cdbb26e95bf12cfad1fb236fecb3b180f61834e0215bf52fd8768.
//
// Solidity: event DepositTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_SyncRouter *SyncRouterFilterer) FilterDepositTicketAdded(opts *bind.FilterOpts, account []common.Address) (*SyncRouterDepositTicketAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "DepositTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterDepositTicketAddedIterator{contract: _SyncRouter.contract, event: "DepositTicketAdded", logs: logs, sub: sub}, nil
}

// WatchDepositTicketAdded is a free log subscription operation binding the contract event 0xc3907160f19cdbb26e95bf12cfad1fb236fecb3b180f61834e0215bf52fd8768.
//
// Solidity: event DepositTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_SyncRouter *SyncRouterFilterer) WatchDepositTicketAdded(opts *bind.WatchOpts, sink chan<- *SyncRouterDepositTicketAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "DepositTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterDepositTicketAdded)
				if err := _SyncRouter.contract.UnpackLog(event, "DepositTicketAdded", log); err != nil {
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

// ParseDepositTicketAdded is a log parse operation binding the contract event 0xc3907160f19cdbb26e95bf12cfad1fb236fecb3b180f61834e0215bf52fd8768.
//
// Solidity: event DepositTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_SyncRouter *SyncRouterFilterer) ParseDepositTicketAdded(log types.Log) (*SyncRouterDepositTicketAdded, error) {
	event := new(SyncRouterDepositTicketAdded)
	if err := _SyncRouter.contract.UnpackLog(event, "DepositTicketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterDepositTicketDeletedIterator is returned from FilterDepositTicketDeleted and is used to iterate over the raw logs and unpacked data for DepositTicketDeleted events raised by the SyncRouter contract.
type SyncRouterDepositTicketDeletedIterator struct {
	Event *SyncRouterDepositTicketDeleted // Event containing the contract specifics and raw log

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
func (it *SyncRouterDepositTicketDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterDepositTicketDeleted)
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
		it.Event = new(SyncRouterDepositTicketDeleted)
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
func (it *SyncRouterDepositTicketDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterDepositTicketDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterDepositTicketDeleted represents a DepositTicketDeleted event raised by the SyncRouter contract.
type SyncRouterDepositTicketDeleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketDeleted is a free log retrieval operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_SyncRouter *SyncRouterFilterer) FilterDepositTicketDeleted(opts *bind.FilterOpts, account []common.Address) (*SyncRouterDepositTicketDeletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "DepositTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterDepositTicketDeletedIterator{contract: _SyncRouter.contract, event: "DepositTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchDepositTicketDeleted is a free log subscription operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_SyncRouter *SyncRouterFilterer) WatchDepositTicketDeleted(opts *bind.WatchOpts, sink chan<- *SyncRouterDepositTicketDeleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "DepositTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterDepositTicketDeleted)
				if err := _SyncRouter.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
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

// ParseDepositTicketDeleted is a log parse operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_SyncRouter *SyncRouterFilterer) ParseDepositTicketDeleted(log types.Log) (*SyncRouterDepositTicketDeleted, error) {
	event := new(SyncRouterDepositTicketDeleted)
	if err := _SyncRouter.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SyncRouter contract.
type SyncRouterOwnershipTransferredIterator struct {
	Event *SyncRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SyncRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterOwnershipTransferred)
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
		it.Event = new(SyncRouterOwnershipTransferred)
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
func (it *SyncRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterOwnershipTransferred represents a OwnershipTransferred event raised by the SyncRouter contract.
type SyncRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SyncRouter *SyncRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SyncRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterOwnershipTransferredIterator{contract: _SyncRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SyncRouter *SyncRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SyncRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterOwnershipTransferred)
				if err := _SyncRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SyncRouter *SyncRouterFilterer) ParseOwnershipTransferred(log types.Log) (*SyncRouterOwnershipTransferred, error) {
	event := new(SyncRouterOwnershipTransferred)
	if err := _SyncRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterPostOpRevertReasonIterator is returned from FilterPostOpRevertReason and is used to iterate over the raw logs and unpacked data for PostOpRevertReason events raised by the SyncRouter contract.
type SyncRouterPostOpRevertReasonIterator struct {
	Event *SyncRouterPostOpRevertReason // Event containing the contract specifics and raw log

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
func (it *SyncRouterPostOpRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterPostOpRevertReason)
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
		it.Event = new(SyncRouterPostOpRevertReason)
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
func (it *SyncRouterPostOpRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterPostOpRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterPostOpRevertReason represents a PostOpRevertReason event raised by the SyncRouter contract.
type SyncRouterPostOpRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostOpRevertReason is a free log retrieval operation binding the contract event 0x676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_SyncRouter *SyncRouterFilterer) FilterPostOpRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*SyncRouterPostOpRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "PostOpRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterPostOpRevertReasonIterator{contract: _SyncRouter.contract, event: "PostOpRevertReason", logs: logs, sub: sub}, nil
}

// WatchPostOpRevertReason is a free log subscription operation binding the contract event 0x676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_SyncRouter *SyncRouterFilterer) WatchPostOpRevertReason(opts *bind.WatchOpts, sink chan<- *SyncRouterPostOpRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "PostOpRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterPostOpRevertReason)
				if err := _SyncRouter.contract.UnpackLog(event, "PostOpRevertReason", log); err != nil {
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
func (_SyncRouter *SyncRouterFilterer) ParsePostOpRevertReason(log types.Log) (*SyncRouterPostOpRevertReason, error) {
	event := new(SyncRouterPostOpRevertReason)
	if err := _SyncRouter.contract.UnpackLog(event, "PostOpRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterReceiveTouchHookIterator is returned from FilterReceiveTouchHook and is used to iterate over the raw logs and unpacked data for ReceiveTouchHook events raised by the SyncRouter contract.
type SyncRouterReceiveTouchHookIterator struct {
	Event *SyncRouterReceiveTouchHook // Event containing the contract specifics and raw log

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
func (it *SyncRouterReceiveTouchHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterReceiveTouchHook)
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
		it.Event = new(SyncRouterReceiveTouchHook)
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
func (it *SyncRouterReceiveTouchHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterReceiveTouchHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterReceiveTouchHook represents a ReceiveTouchHook event raised by the SyncRouter contract.
type SyncRouterReceiveTouchHook struct {
	Success         bool
	Data            []byte
	PackHookMessage []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReceiveTouchHook is a free log retrieval operation binding the contract event 0xc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb398079446.
//
// Solidity: event ReceiveTouchHook(bool success, bytes data, bytes packHookMessage)
func (_SyncRouter *SyncRouterFilterer) FilterReceiveTouchHook(opts *bind.FilterOpts) (*SyncRouterReceiveTouchHookIterator, error) {

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "ReceiveTouchHook")
	if err != nil {
		return nil, err
	}
	return &SyncRouterReceiveTouchHookIterator{contract: _SyncRouter.contract, event: "ReceiveTouchHook", logs: logs, sub: sub}, nil
}

// WatchReceiveTouchHook is a free log subscription operation binding the contract event 0xc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb398079446.
//
// Solidity: event ReceiveTouchHook(bool success, bytes data, bytes packHookMessage)
func (_SyncRouter *SyncRouterFilterer) WatchReceiveTouchHook(opts *bind.WatchOpts, sink chan<- *SyncRouterReceiveTouchHook) (event.Subscription, error) {

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "ReceiveTouchHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterReceiveTouchHook)
				if err := _SyncRouter.contract.UnpackLog(event, "ReceiveTouchHook", log); err != nil {
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

// ParseReceiveTouchHook is a log parse operation binding the contract event 0xc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb398079446.
//
// Solidity: event ReceiveTouchHook(bool success, bytes data, bytes packHookMessage)
func (_SyncRouter *SyncRouterFilterer) ParseReceiveTouchHook(log types.Log) (*SyncRouterReceiveTouchHook, error) {
	event := new(SyncRouterReceiveTouchHook)
	if err := _SyncRouter.contract.UnpackLog(event, "ReceiveTouchHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterRefundEventIterator is returned from FilterRefundEvent and is used to iterate over the raw logs and unpacked data for RefundEvent events raised by the SyncRouter contract.
type SyncRouterRefundEventIterator struct {
	Event *SyncRouterRefundEvent // Event containing the contract specifics and raw log

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
func (it *SyncRouterRefundEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterRefundEvent)
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
		it.Event = new(SyncRouterRefundEvent)
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
func (it *SyncRouterRefundEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterRefundEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterRefundEvent represents a RefundEvent event raised by the SyncRouter contract.
type SyncRouterRefundEvent struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefundEvent is a free log retrieval operation binding the contract event 0xfe72f5b9769cf4f44c25a2cae288955570e053c444408fa295182f65deceb086.
//
// Solidity: event RefundEvent(address indexed token, address indexed receiver, uint256 amount)
func (_SyncRouter *SyncRouterFilterer) FilterRefundEvent(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*SyncRouterRefundEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "RefundEvent", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterRefundEventIterator{contract: _SyncRouter.contract, event: "RefundEvent", logs: logs, sub: sub}, nil
}

// WatchRefundEvent is a free log subscription operation binding the contract event 0xfe72f5b9769cf4f44c25a2cae288955570e053c444408fa295182f65deceb086.
//
// Solidity: event RefundEvent(address indexed token, address indexed receiver, uint256 amount)
func (_SyncRouter *SyncRouterFilterer) WatchRefundEvent(opts *bind.WatchOpts, sink chan<- *SyncRouterRefundEvent, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "RefundEvent", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterRefundEvent)
				if err := _SyncRouter.contract.UnpackLog(event, "RefundEvent", log); err != nil {
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

// ParseRefundEvent is a log parse operation binding the contract event 0xfe72f5b9769cf4f44c25a2cae288955570e053c444408fa295182f65deceb086.
//
// Solidity: event RefundEvent(address indexed token, address indexed receiver, uint256 amount)
func (_SyncRouter *SyncRouterFilterer) ParseRefundEvent(log types.Log) (*SyncRouterRefundEvent, error) {
	event := new(SyncRouterRefundEvent)
	if err := _SyncRouter.contract.UnpackLog(event, "RefundEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterUserOperationEventIterator is returned from FilterUserOperationEvent and is used to iterate over the raw logs and unpacked data for UserOperationEvent events raised by the SyncRouter contract.
type SyncRouterUserOperationEventIterator struct {
	Event *SyncRouterUserOperationEvent // Event containing the contract specifics and raw log

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
func (it *SyncRouterUserOperationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterUserOperationEvent)
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
		it.Event = new(SyncRouterUserOperationEvent)
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
func (it *SyncRouterUserOperationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterUserOperationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterUserOperationEvent represents a UserOperationEvent event raised by the SyncRouter contract.
type SyncRouterUserOperationEvent struct {
	UserOpHash    [32]byte
	Sender        common.Address
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserOperationEvent is a free log retrieval operation binding the contract event 0xfb68fae769a73ec810729fb1cf34aa9faf18a18c0d69fb5388726c793abbf5aa.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_SyncRouter *SyncRouterFilterer) FilterUserOperationEvent(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*SyncRouterUserOperationEventIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "UserOperationEvent", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterUserOperationEventIterator{contract: _SyncRouter.contract, event: "UserOperationEvent", logs: logs, sub: sub}, nil
}

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0xfb68fae769a73ec810729fb1cf34aa9faf18a18c0d69fb5388726c793abbf5aa.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_SyncRouter *SyncRouterFilterer) WatchUserOperationEvent(opts *bind.WatchOpts, sink chan<- *SyncRouterUserOperationEvent, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "UserOperationEvent", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterUserOperationEvent)
				if err := _SyncRouter.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
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

// ParseUserOperationEvent is a log parse operation binding the contract event 0xfb68fae769a73ec810729fb1cf34aa9faf18a18c0d69fb5388726c793abbf5aa.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_SyncRouter *SyncRouterFilterer) ParseUserOperationEvent(log types.Log) (*SyncRouterUserOperationEvent, error) {
	event := new(SyncRouterUserOperationEvent)
	if err := _SyncRouter.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterUserOperationPrefundTooLowIterator is returned from FilterUserOperationPrefundTooLow and is used to iterate over the raw logs and unpacked data for UserOperationPrefundTooLow events raised by the SyncRouter contract.
type SyncRouterUserOperationPrefundTooLowIterator struct {
	Event *SyncRouterUserOperationPrefundTooLow // Event containing the contract specifics and raw log

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
func (it *SyncRouterUserOperationPrefundTooLowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterUserOperationPrefundTooLow)
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
		it.Event = new(SyncRouterUserOperationPrefundTooLow)
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
func (it *SyncRouterUserOperationPrefundTooLowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterUserOperationPrefundTooLowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterUserOperationPrefundTooLow represents a UserOperationPrefundTooLow event raised by the SyncRouter contract.
type SyncRouterUserOperationPrefundTooLow struct {
	UserOpHash [32]byte
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserOperationPrefundTooLow is a free log retrieval operation binding the contract event 0x1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff28.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender)
func (_SyncRouter *SyncRouterFilterer) FilterUserOperationPrefundTooLow(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*SyncRouterUserOperationPrefundTooLowIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "UserOperationPrefundTooLow", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterUserOperationPrefundTooLowIterator{contract: _SyncRouter.contract, event: "UserOperationPrefundTooLow", logs: logs, sub: sub}, nil
}

// WatchUserOperationPrefundTooLow is a free log subscription operation binding the contract event 0x1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff28.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender)
func (_SyncRouter *SyncRouterFilterer) WatchUserOperationPrefundTooLow(opts *bind.WatchOpts, sink chan<- *SyncRouterUserOperationPrefundTooLow, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "UserOperationPrefundTooLow", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterUserOperationPrefundTooLow)
				if err := _SyncRouter.contract.UnpackLog(event, "UserOperationPrefundTooLow", log); err != nil {
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
func (_SyncRouter *SyncRouterFilterer) ParseUserOperationPrefundTooLow(log types.Log) (*SyncRouterUserOperationPrefundTooLow, error) {
	event := new(SyncRouterUserOperationPrefundTooLow)
	if err := _SyncRouter.contract.UnpackLog(event, "UserOperationPrefundTooLow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterUserOperationRevertReasonIterator is returned from FilterUserOperationRevertReason and is used to iterate over the raw logs and unpacked data for UserOperationRevertReason events raised by the SyncRouter contract.
type SyncRouterUserOperationRevertReasonIterator struct {
	Event *SyncRouterUserOperationRevertReason // Event containing the contract specifics and raw log

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
func (it *SyncRouterUserOperationRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterUserOperationRevertReason)
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
		it.Event = new(SyncRouterUserOperationRevertReason)
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
func (it *SyncRouterUserOperationRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterUserOperationRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterUserOperationRevertReason represents a UserOperationRevertReason event raised by the SyncRouter contract.
type SyncRouterUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserOperationRevertReason is a free log retrieval operation binding the contract event 0x0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_SyncRouter *SyncRouterFilterer) FilterUserOperationRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*SyncRouterUserOperationRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterUserOperationRevertReasonIterator{contract: _SyncRouter.contract, event: "UserOperationRevertReason", logs: logs, sub: sub}, nil
}

// WatchUserOperationRevertReason is a free log subscription operation binding the contract event 0x0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_SyncRouter *SyncRouterFilterer) WatchUserOperationRevertReason(opts *bind.WatchOpts, sink chan<- *SyncRouterUserOperationRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterUserOperationRevertReason)
				if err := _SyncRouter.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
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
func (_SyncRouter *SyncRouterFilterer) ParseUserOperationRevertReason(log types.Log) (*SyncRouterUserOperationRevertReason, error) {
	event := new(SyncRouterUserOperationRevertReason)
	if err := _SyncRouter.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterValueDepositAmountIterator is returned from FilterValueDepositAmount and is used to iterate over the raw logs and unpacked data for ValueDepositAmount events raised by the SyncRouter contract.
type SyncRouterValueDepositAmountIterator struct {
	Event *SyncRouterValueDepositAmount // Event containing the contract specifics and raw log

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
func (it *SyncRouterValueDepositAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterValueDepositAmount)
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
		it.Event = new(SyncRouterValueDepositAmount)
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
func (it *SyncRouterValueDepositAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterValueDepositAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterValueDepositAmount represents a ValueDepositAmount event raised by the SyncRouter contract.
type SyncRouterValueDepositAmount struct {
	Sender             common.Address
	Owner              common.Address
	ValueDepositAmount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValueDepositAmount is a free log retrieval operation binding the contract event 0x8b5f503e9ee0a7309afb2323db56ba3b483f5d7545b8500147ebce30b5edac5e.
//
// Solidity: event ValueDepositAmount(address indexed sender, address indexed owner, uint256 indexed valueDepositAmount)
func (_SyncRouter *SyncRouterFilterer) FilterValueDepositAmount(opts *bind.FilterOpts, sender []common.Address, owner []common.Address, valueDepositAmount []*big.Int) (*SyncRouterValueDepositAmountIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var valueDepositAmountRule []interface{}
	for _, valueDepositAmountItem := range valueDepositAmount {
		valueDepositAmountRule = append(valueDepositAmountRule, valueDepositAmountItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "ValueDepositAmount", senderRule, ownerRule, valueDepositAmountRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterValueDepositAmountIterator{contract: _SyncRouter.contract, event: "ValueDepositAmount", logs: logs, sub: sub}, nil
}

// WatchValueDepositAmount is a free log subscription operation binding the contract event 0x8b5f503e9ee0a7309afb2323db56ba3b483f5d7545b8500147ebce30b5edac5e.
//
// Solidity: event ValueDepositAmount(address indexed sender, address indexed owner, uint256 indexed valueDepositAmount)
func (_SyncRouter *SyncRouterFilterer) WatchValueDepositAmount(opts *bind.WatchOpts, sink chan<- *SyncRouterValueDepositAmount, sender []common.Address, owner []common.Address, valueDepositAmount []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var valueDepositAmountRule []interface{}
	for _, valueDepositAmountItem := range valueDepositAmount {
		valueDepositAmountRule = append(valueDepositAmountRule, valueDepositAmountItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "ValueDepositAmount", senderRule, ownerRule, valueDepositAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterValueDepositAmount)
				if err := _SyncRouter.contract.UnpackLog(event, "ValueDepositAmount", log); err != nil {
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

// ParseValueDepositAmount is a log parse operation binding the contract event 0x8b5f503e9ee0a7309afb2323db56ba3b483f5d7545b8500147ebce30b5edac5e.
//
// Solidity: event ValueDepositAmount(address indexed sender, address indexed owner, uint256 indexed valueDepositAmount)
func (_SyncRouter *SyncRouterFilterer) ParseValueDepositAmount(log types.Log) (*SyncRouterValueDepositAmount, error) {
	event := new(SyncRouterValueDepositAmount)
	if err := _SyncRouter.contract.UnpackLog(event, "ValueDepositAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterVizingSwapEventIterator is returned from FilterVizingSwapEvent and is used to iterate over the raw logs and unpacked data for VizingSwapEvent events raised by the SyncRouter contract.
type SyncRouterVizingSwapEventIterator struct {
	Event *SyncRouterVizingSwapEvent // Event containing the contract specifics and raw log

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
func (it *SyncRouterVizingSwapEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterVizingSwapEvent)
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
		it.Event = new(SyncRouterVizingSwapEvent)
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
func (it *SyncRouterVizingSwapEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterVizingSwapEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterVizingSwapEvent represents a VizingSwapEvent event raised by the SyncRouter contract.
type SyncRouterVizingSwapEvent struct {
	Sender    common.Address
	TokenIn   common.Address
	TokenOut  common.Address
	Receiver  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVizingSwapEvent is a free log retrieval operation binding the contract event 0x9f93dc48f356aa24b3e518eface42b72c97cf0ea07f286aa492e0a91e301f6fd.
//
// Solidity: event VizingSwapEvent(address indexed sender, address indexed tokenIn, address indexed tokenOut, address receiver, uint256 amountIn, uint256 amountOut)
func (_SyncRouter *SyncRouterFilterer) FilterVizingSwapEvent(opts *bind.FilterOpts, sender []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*SyncRouterVizingSwapEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "VizingSwapEvent", senderRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterVizingSwapEventIterator{contract: _SyncRouter.contract, event: "VizingSwapEvent", logs: logs, sub: sub}, nil
}

// WatchVizingSwapEvent is a free log subscription operation binding the contract event 0x9f93dc48f356aa24b3e518eface42b72c97cf0ea07f286aa492e0a91e301f6fd.
//
// Solidity: event VizingSwapEvent(address indexed sender, address indexed tokenIn, address indexed tokenOut, address receiver, uint256 amountIn, uint256 amountOut)
func (_SyncRouter *SyncRouterFilterer) WatchVizingSwapEvent(opts *bind.WatchOpts, sink chan<- *SyncRouterVizingSwapEvent, sender []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "VizingSwapEvent", senderRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterVizingSwapEvent)
				if err := _SyncRouter.contract.UnpackLog(event, "VizingSwapEvent", log); err != nil {
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

// ParseVizingSwapEvent is a log parse operation binding the contract event 0x9f93dc48f356aa24b3e518eface42b72c97cf0ea07f286aa492e0a91e301f6fd.
//
// Solidity: event VizingSwapEvent(address indexed sender, address indexed tokenIn, address indexed tokenOut, address receiver, uint256 amountIn, uint256 amountOut)
func (_SyncRouter *SyncRouterFilterer) ParseVizingSwapEvent(log types.Log) (*SyncRouterVizingSwapEvent, error) {
	event := new(SyncRouterVizingSwapEvent)
	if err := _SyncRouter.contract.UnpackLog(event, "VizingSwapEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterWithdrawTicketAddedIterator is returned from FilterWithdrawTicketAdded and is used to iterate over the raw logs and unpacked data for WithdrawTicketAdded events raised by the SyncRouter contract.
type SyncRouterWithdrawTicketAddedIterator struct {
	Event *SyncRouterWithdrawTicketAdded // Event containing the contract specifics and raw log

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
func (it *SyncRouterWithdrawTicketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterWithdrawTicketAdded)
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
		it.Event = new(SyncRouterWithdrawTicketAdded)
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
func (it *SyncRouterWithdrawTicketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterWithdrawTicketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterWithdrawTicketAdded represents a WithdrawTicketAdded event raised by the SyncRouter contract.
type SyncRouterWithdrawTicketAdded struct {
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketAdded is a free log retrieval operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_SyncRouter *SyncRouterFilterer) FilterWithdrawTicketAdded(opts *bind.FilterOpts, account []common.Address) (*SyncRouterWithdrawTicketAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "WithdrawTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterWithdrawTicketAddedIterator{contract: _SyncRouter.contract, event: "WithdrawTicketAdded", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketAdded is a free log subscription operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_SyncRouter *SyncRouterFilterer) WatchWithdrawTicketAdded(opts *bind.WatchOpts, sink chan<- *SyncRouterWithdrawTicketAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "WithdrawTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterWithdrawTicketAdded)
				if err := _SyncRouter.contract.UnpackLog(event, "WithdrawTicketAdded", log); err != nil {
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

// ParseWithdrawTicketAdded is a log parse operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_SyncRouter *SyncRouterFilterer) ParseWithdrawTicketAdded(log types.Log) (*SyncRouterWithdrawTicketAdded, error) {
	event := new(SyncRouterWithdrawTicketAdded)
	if err := _SyncRouter.contract.UnpackLog(event, "WithdrawTicketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncRouterWithdrawTicketDeletedIterator is returned from FilterWithdrawTicketDeleted and is used to iterate over the raw logs and unpacked data for WithdrawTicketDeleted events raised by the SyncRouter contract.
type SyncRouterWithdrawTicketDeletedIterator struct {
	Event *SyncRouterWithdrawTicketDeleted // Event containing the contract specifics and raw log

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
func (it *SyncRouterWithdrawTicketDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterWithdrawTicketDeleted)
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
		it.Event = new(SyncRouterWithdrawTicketDeleted)
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
func (it *SyncRouterWithdrawTicketDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterWithdrawTicketDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterWithdrawTicketDeleted represents a WithdrawTicketDeleted event raised by the SyncRouter contract.
type SyncRouterWithdrawTicketDeleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketDeleted is a free log retrieval operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_SyncRouter *SyncRouterFilterer) FilterWithdrawTicketDeleted(opts *bind.FilterOpts, account []common.Address) (*SyncRouterWithdrawTicketDeletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "WithdrawTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterWithdrawTicketDeletedIterator{contract: _SyncRouter.contract, event: "WithdrawTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketDeleted is a free log subscription operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_SyncRouter *SyncRouterFilterer) WatchWithdrawTicketDeleted(opts *bind.WatchOpts, sink chan<- *SyncRouterWithdrawTicketDeleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "WithdrawTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterWithdrawTicketDeleted)
				if err := _SyncRouter.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
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

// ParseWithdrawTicketDeleted is a log parse operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_SyncRouter *SyncRouterFilterer) ParseWithdrawTicketDeleted(log types.Log) (*SyncRouterWithdrawTicketDeleted, error) {
	event := new(SyncRouterWithdrawTicketDeleted)
	if err := _SyncRouter.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
