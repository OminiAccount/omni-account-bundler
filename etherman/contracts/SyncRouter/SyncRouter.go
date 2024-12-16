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
	BatchsMessage           []byte
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vizingPad\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_Hook\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExecutionCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LandingPadAccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplement\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"packHookMessage\",\"type\":\"bytes\"}],\"name\":\"ReceiveTouchHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"VizingSwapEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"DataExcecuteNumber\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Hook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LandingPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LaunchPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LockWay\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"MirrorEntryPoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"newGaslimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"newGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newRelayer\",\"type\":\"address\"}],\"name\":\"changeDefaultSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBridgeMode\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGasPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGaslimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destChainid\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"additionParams\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"estimateVizingGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vizingGasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"fetchOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"fetchUserOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"cmp\",\"type\":\"tuple\"}],\"name\":\"getUserOmniEncodeMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveExtraMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveStandardMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selectedRelayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"sendOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"cmp\",\"type\":\"tuple\"}],\"name\":\"sendUserOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"}],\"name\":\"setMirrorEntryPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"testReceiveMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thisRelayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523462000038576200001f6200001862000131565b9162000694565b620000296200003e565b614d0d62000a1a8239614d0d90f35b62000044565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90620000729062000048565b810190811060018060401b038211176200008b57604052565b62000052565b90620000a8620000a06200003e565b928362000066565b565b5f80fd5b60018060a01b031690565b620000c490620000ae565b90565b620000d281620000b9565b03620000da57565b5f80fd5b90505190620000ed82620000c7565b565b90916060828403126200012b57620001286200010e845f8501620000de565b936200011e8160208601620000de565b93604001620000de565b90565b620000aa565b620001546200572780380380620001488162000091565b928339810190620000ef565b909192565b60a01b90565b906200017060ff60a01b9162000159565b9181191691161790565b90565b60ff60f81b1690565b60f81b90565b620001a56200019f620001ab926200017a565b62000186565b6200017d565b90565b60f81c90565b620001bf90620001ae565b90565b90620001dc620001d6620001e4926200018c565b620001b4565b82546200015f565b9055565b90565b90565b90565b6200020a620002046200021092620001e8565b620001ee565b620001eb565b90565b60018060401b03811162000232576200022e60209162000048565b0190565b62000052565b906200024e620002488362000213565b62000091565b918252565b369037565b9062000283620002688362000238565b9260208062000278869362000213565b920191039062000253565b565b5190565b634e487b7160e01b5f52602260045260245ffd5b9060016002830492168015620002c0575b6020831014620002ba57565b62000289565b91607f1691620002ae565b5f5260205f2090565b601f602091010490565b1b90565b9190600862000301910291620002fa5f1984620002de565b92620002de565b9181191691161790565b620003246200031e6200032a92620001eb565b620001ee565b620001eb565b90565b90565b91906200034b6200034562000354936200030b565b6200032d565b908354620002e2565b9055565b5f90565b62000372916200036b62000358565b9162000330565b565b5b81811062000381575050565b80620003905f6001936200035c565b0162000375565b9190601f8111620003a8575b505050565b620003b7620003e293620002cb565b906020620003c584620002d4565b83019310620003eb575b620003da90620002d4565b019062000374565b5f8080620003a3565b9150620003da81929050620003cf565b1c90565b9062000411905f1990600802620003fb565b191690565b816200042291620003ff565b906002021790565b90620004368162000285565b9060018060401b03821162000507576200045d826200045685546200029d565b8562000397565b602090601f8311600114620004965791809162000484935f9262000489575b505062000416565b90555b565b90915001515f806200047c565b601f19831691620004a785620002cb565b925f5b818110620004ee57509160029391856001969410620004d1575b5050500201905562000487565b620004e3910151601f841690620003ff565b90555f8080620004c4565b91936020600181928787015181550195019201620004aa565b62000052565b9062000519916200042a565b565b5f1b90565b906200053062ffffff916200051b565b9181191691161790565b90565b62ffffff1690565b6200055e6200055862000564926200053a565b620001ee565b6200053d565b90565b90565b90620005846200057e6200058c9262000545565b62000567565b825462000520565b9055565b60181b90565b90620005ad6301000000600160581b039162000590565b9181191691161790565b90565b60018060401b031690565b620005de620005d8620005e492620005b7565b620001ee565b620005ba565b90565b90565b9062000604620005fe6200060c92620005c5565b620005e7565b825462000596565b9055565b906200062360018060a01b03916200051b565b9181191691161790565b62000646620006406200064c92620000ae565b620001ee565b620000ae565b90565b6200065a906200062d565b90565b62000668906200064f565b90565b90565b90620006886200068262000690926200065d565b6200066b565b825462000610565b9055565b6200070a9291620006aa6200070292336200070c565b620006b860016006620001c2565b620006d9620006d1620006cb5f620001f1565b62000258565b60076200050d565b620006e96207a12060086200056a565b620006fa633b9aca006008620005ea565b60056200066e565b60066200066e565b565b9062000718916200071a565b565b90620007269162000728565b565b906200073491620007a5565b565b6200074f6200074962000755926200017a565b620001ee565b620001eb565b90565b62000764600162000736565b90565b90620007755f19916200051b565b9181191691161790565b906200079962000793620007a1926200030b565b6200032d565b825462000767565b9055565b90620007b1916200081f565b620007c7620007bf62000758565b60036200077f565b565b620007e2620007dc620007e892620001e8565b620001ee565b620000ae565b90565b620007f690620007c9565b90565b6200080490620000b9565b9052565b91906200081d905f60208501940190620007f9565b565b906200082b9062000893565b806200084c620008456200083f5f620007eb565b620000b9565b91620000b9565b146200085f576200085d9062000928565b565b6200088f6200086e5f620007eb565b620008786200003e565b918291631e4fbdf760e01b83526004830162000808565b0390fd5b6200089e90620008a0565b565b620008ac9080620008ae565b565b620008bd620008c392620008c5565b620009e8565b565b620008d090620008d2565b565b620008dd90620008df565b565b620008ea9062000a01565b565b5f1c90565b60018060a01b031690565b6200090b6200091191620008ec565b620008f1565b90565b620009209054620008fc565b90565b5f0190565b62000934600262000914565b620009418260026200066e565b9062000979620009727f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936200065d565b916200065d565b91620009846200003e565b80620009908162000923565b0390a3565b620009a0906200062d565b90565b620009ae9062000995565b90565b620009bc9062000995565b90565b90565b90620009dc620009d6620009e492620009b1565b620009bf565b825462000610565b9055565b620009f7620009ff91620009a3565b6001620009c2565b565b62000a1062000a1791620009a3565b5f620009c2565b56fe60806040526004361015610015575b366116e657005b61001f5f3561020c565b80622232e31461020757806273b555146102025780630e82845d146101fd5780631490a358146101f85780631b3348f0146101f35780633169eb13146101ee578063397a567d146101e95780633eba5864146101e457806345636279146101df5780635ad3ad06146101da5780635aeb4d77146101d55780635e45da23146101d057806360b43029146101cb578063715018a6146101c657806376c81312146101c15780638da5cb5b146101bc5780638ea23ddf146101b757806393afae93146101b257806399e581fa146101ad578063a834a900146101a8578063ad5c4648146101a3578063b0cfd4d21461019e578063c5f378d814610199578063c967f88614610194578063d5f3e0081461018f578063de4b1d6d1461018a578063de8aeda014610185578063e0b838e914610180578063e7b4294c1461017b578063f2fde38b146101765763f31489250361000e576116b9565b611600565b6115ad565b611568565b611526565b6114e1565b61140c565b6113a9565b610ec2565b610e7e565b610e49565b610e06565b610d74565b610d0a565b610cd5565b610c91565b610c55565b610ba7565b610b72565b610b05565b610ad0565b610a71565b6109d7565b610973565b610887565b61081c565b6107d1565b610750565b610653565b61058c565b610466565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906102549061022c565b810190811067ffffffffffffffff82111761026e57604052565b610236565b9061028661027f610212565b928361024a565b565b67ffffffffffffffff81116102a6576102a260209161022c565b0190565b610236565b90825f939282370152565b909291926102cb6102c682610288565b610273565b938185526020850190828401116102e7576102e5926102ab565b565b610228565b9080601f8301121561030a57816020610307933591016102b6565b90565b610224565b9060208282031261033f575f82013567ffffffffffffffff811161033a5761033792016102ec565b90565b610220565b61021c565b5190565b905090565b5f5b83811061035f575050905f910152565b80602091830151818501520161034f565b61039561038c9260209261038381610344565b94858093610348565b9384910161034d565b0190565b90565b90565b6103ab6103b091610399565b61039c565b9052565b6103c46103cb9160209493610370565b809261039f565b0190565b6103e36103da610212565b928392836103b4565b03902090565b6103f2916103cf565b90565b1c90565b60ff1690565b61040f90600861041493026103f5565b6103f9565b90565b9061042291546103ff565b90565b61043b90610436600c915f926103e9565b610417565b90565b60ff1690565b61044d9061043e565b9052565b9190610464905f60208501940190610444565b565b346104965761049261048161047c36600461030f565b610425565b610489610212565b91829182610451565b0390f35b610218565b67ffffffffffffffff1690565b6104b18161049b565b036104b857565b5f80fd5b905035906104c9826104a8565b565b6104d481610399565b036104db57565b5f80fd5b905035906104ec826104cb565b565b5f80fd5b5f80fd5b909182601f830112156105305781359167ffffffffffffffff831161052b57602001926001830284011161052657565b6104f2565b6104ee565b610224565b916060838303126105825761054c825f85016104bc565b9261055a83602083016104df565b92604082013567ffffffffffffffff811161057d5761057992016104f6565b9091565b610220565b61021c565b5f0190565b610597366004610535565b92919091611769565b5f9103126105aa57565b61021c565b60018060a01b031690565b6105ca9060086105cf93026103f5565b6105af565b90565b906105dd91546105ba565b90565b6105ec60015f906105d2565b90565b60018060a01b031690565b90565b61061161060c610616926105ef565b6105fa565b6105ef565b90565b610622906105fd565b90565b61062e90610619565b90565b61063a90610625565b9052565b9190610651905f60208501940190610631565b565b34610683576106633660046105a0565b61067f61066e6105e0565b610676610212565b9182918261063e565b0390f35b610218565b906020828203126106a15761069e915f016104bc565b90565b61021c565b6106ba6106b56106bf9261049b565b6105fa565b61049b565b90565b906106cc906106a6565b5f5260205260405f2090565b60018060a01b031690565b6106f39060086106f893026103f5565b6106d8565b90565b9061070691546106e3565b90565b61071f9061071a600a915f926106c2565b6106fb565b90565b61072b906105ef565b90565b61073790610722565b9052565b919061074e905f6020850194019061072e565b565b346107805761077c61076b610766366004610688565b610709565b610773610212565b9182918261073b565b0390f35b610218565b5f80fd5b908160409103126107975790565b610785565b906020828203126107cc575f82013567ffffffffffffffff81116107c7576107c49201610789565b90565b610220565b61021c565b6107e46107df36600461079c565b611d76565b6107ec610212565b806107f681610587565b0390f35b61080390610399565b9052565b919061081a905f602085019401906107fa565b565b3461084c5761084861083761083236600461079c565b6123c0565b61083f610212565b91829182610807565b0390f35b610218565b90602082820312610882575f82013567ffffffffffffffff811161087d5761087992016104f6565b9091565b610220565b61021c565b61089b610895366004610851565b9061295c565b6108a3610212565b806108ad81610587565b0390f35b6108ba81610722565b036108c157565b5f80fd5b905035906108d2826108b1565b565b909182601f8301121561090e5781359167ffffffffffffffff831161090957602001926020830284011161090457565b6104f2565b6104ee565b610224565b9060808282031261096e5761092a815f84016104bc565b9261093882602085016108c5565b9261094683604083016104df565b92606082013567ffffffffffffffff81116109695761096592016108d4565b9091565b610220565b61021c565b346109a7576109a3610992610989366004610913565b93929092612d0a565b61099a610212565b91829182610807565b0390f35b610218565b60ff60f81b1690565b6109be906109ac565b9052565b91906109d5905f602085019401906109b5565b565b34610a07576109e73660046105a0565b610a036109f2612dfa565b6109fa610212565b918291826109c2565b0390f35b610218565b67ffffffffffffffff1690565b610a29906008610a2e93026103f5565b610a0c565b90565b90610a3c9154610a19565b90565b610a4c6008600b90610a31565b90565b610a589061049b565b9052565b9190610a6f905f60208501940190610a4f565b565b34610aa157610a813660046105a0565b610a9d610a8c610a3f565b610a94610212565b91829182610a5c565b0390f35b610218565b62ffffff1690565b610ab790610aa6565b9052565b9190610ace905f60208501940190610aae565b565b34610b0057610ae03660046105a0565b610afc610aeb612e09565b610af3610212565b91829182610abb565b0390f35b610218565b34610b3557610b153660046105a0565b610b31610b20612e14565b610b28610212565b91829182610abb565b0390f35b610218565b90565b610b4d906008610b5293026103f5565b610b3a565b90565b90610b609154610b3d565b90565b610b6f60045f90610b55565b90565b34610ba257610b823660046105a0565b610b9e610b8d610b63565b610b95610212565b91829182610807565b0390f35b610218565b34610bd557610bb73660046105a0565b610bbf612e6c565b610bc7610212565b80610bd181610587565b0390f35b610218565b9190608083820312610c5057610bf2815f85016104df565b92610c0082602083016104bc565b92604082013567ffffffffffffffff8111610c4b5783610c219184016104f6565b929093606082013567ffffffffffffffff8111610c4657610c4292016104f6565b9091565b610220565b610220565b61021c565b34610c8c57610c88610c77610c6b366004610bda565b94939093929192612e84565b610c7f610212565b91829182610807565b0390f35b610218565b34610cc157610ca13660046105a0565b610cbd610cac612eb9565b610cb4610212565b9182918261073b565b0390f35b610218565b610cd260065f906106fb565b90565b34610d0557610ce53660046105a0565b610d01610cf0610cc6565b610cf8610212565b9182918261073b565b0390f35b610218565b610d21610d18366004610913565b9392909261313c565b610d29610212565b80610d3381610587565b0390f35b62ffffff1690565b610d4f906008610d5493026103f5565b610d37565b90565b90610d629154610d3f565b90565b610d7160085f90610d57565b90565b34610da457610d843660046105a0565b610da0610d8f610d65565b610d97610212565b91829182610abb565b0390f35b610218565b610db281610aa6565b03610db957565b5f80fd5b90503590610dca82610da9565b565b9091606082840312610e0157610dfe610de7845f8501610dbd565b93610df581602086016104bc565b936040016108c5565b90565b61021c565b34610e3557610e1f610e19366004610dcc565b91613272565b610e27610212565b80610e3181610587565b0390f35b610218565b610e4660055f906106fb565b90565b34610e7957610e593660046105a0565b610e75610e64610e3a565b610e6c610212565b9182918261073b565b0390f35b610218565b34610eae57610e8e3660046105a0565b610eaa610e9961327f565b610ea1610212565b9182918261073b565b0390f35b610218565b610ebf60095f906106fb565b90565b34610ef257610ed23660046105a0565b610eee610edd610eb3565b610ee5610212565b9182918261073b565b0390f35b610218565b5f80fd5b5f80fd5b67ffffffffffffffff8111610f175760208091020190565b610236565b610f258161043e565b03610f2c57565b5f80fd5b90503590610f3d82610f1c565b565b91909161010081840312610fff57610f58610100610273565b92610f65815f84016104bc565b5f850152610f7681602084016104bc565b6020850152610f8881604084016104bc565b6040850152610f9a81606084016104bc565b6060850152610fac81608084016104bc565b6080850152610fbe8160a084016104bc565b60a0850152610fd08160c084016104bc565b60c085015260e082013567ffffffffffffffff8111610ffa57610ff392016102ec565b60e0830152565b610efb565b610ef7565b91909160e0818403126110c85761101b60e0610273565b92611028815f8401610f30565b5f8501526110398160208401610f30565b602085015261104b81604084016104df565b604085015261105d81606084016108c5565b606085015261106f81608084016108c5565b608085015260a082013567ffffffffffffffff81116110c35781611094918401610f3f565b60a085015260c082013567ffffffffffffffff81116110be576110b79201610f3f565b60c0830152565b610efb565b610efb565b610ef7565b9291906110e16110dc82610eff565b610273565b93818552602080860192028101918383116111385781905b838210611107575050505050565b813567ffffffffffffffff8111611133576020916111288784938701611004565b8152019101906110f9565b610224565b6104f2565b9080601f8301121561115b57816020611158933591016110cd565b90565b610224565b919091610180818403126112a057611179610180610273565b92611186815f8401610f30565b5f8501526111978160208401610dbd565b60208501526111a981604084016104bc565b60408501526111bb81606084016104bc565b60608501526111cd81608084016104bc565b60808501526111df8160a084016104bc565b60a08501526111f18160c084016108c5565b60c08501526112038160e084016108c5565b60e08501526112168161010084016104df565b61010085015261012082013567ffffffffffffffff811161129b578161123d9184016102ec565b61012085015261014082013567ffffffffffffffff811161129657816112649184016102ec565b61014085015261016082013567ffffffffffffffff81116112915761128992016102ec565b610160830152565b610efb565b610efb565b610efb565b610ef7565b91909160408184031261130f576112bc6040610273565b925f82013567ffffffffffffffff811161130a57816112dc91840161113d565b5f850152602082013567ffffffffffffffff8111611305576112fe9201611160565b6020830152565b610efb565b610efb565b610ef7565b90602082820312611344575f82013567ffffffffffffffff811161133f5761133c92016112a5565b90565b610220565b61021c565b60209181520190565b61137161137a60209361137f9361136881610344565b93848093611349565b9586910161034d565b61022c565b0190565b916113a69261139960408201935f8301906107fa565b6020818403910152611352565b90565b346113da576113c16113bc366004611314565b613b55565b906113d66113cd610212565b92839283611383565b0390f35b610218565b919060408382031261140757806113fb611404925f86016104bc565b936020016108c5565b90565b61021c565b3461143b5761142561141f3660046113df565b9061444f565b61142d610212565b8061143781610587565b0390f35b610218565b9060208282031261145957611456915f016104df565b90565b61021c565b61147261146d61147792610399565b6105fa565b610399565b90565b906114849061145e565b5f5260205260405f2090565b60f81b90565b61149f90611490565b90565b6114b29060086114b793026103f5565b611496565b90565b906114c591546114a2565b90565b6114de906114d9600b915f9261147a565b6114ba565b90565b346115115761150d6114fc6114f7366004611440565b6114c8565b611504610212565b918291826109c2565b0390f35b610218565b6115236008601390610a31565b90565b34611556576115363660046105a0565b611552611541611516565b611549610212565b91829182610a5c565b0390f35b610218565b6115655f806105d2565b90565b34611598576115783660046105a0565b61159461158361155b565b61158b610212565b9182918261063e565b0390f35b610218565b6115aa6008600390610a31565b90565b346115dd576115bd3660046105a0565b6115d96115c861159d565b6115d0610212565b91829182610a5c565b0390f35b610218565b906020828203126115fb576115f8915f016108c5565b90565b61021c565b3461162e576116186116133660046115e2565b6144c7565b611620610212565b8061162a81610587565b0390f35b610218565b90565b61163f81611633565b0361164657565b5f80fd5b9050359061165782611636565b565b906080828203126116b457611670815f840161164a565b9261167e82602085016104bc565b9261168c83604083016104df565b92606082013567ffffffffffffffff81116116af576116ab92016104f6565b9091565b610220565b61021c565b6116d06116c7366004611659565b93929092614540565b6116d8610212565b806116e281610587565b0390f35b5f80fd5b5f1c90565b6116fb611700916116ea565b6105af565b90565b61170d90546116ef565b90565b3361173461172e6117296117246001611703565b610625565b610722565b91610722565b031561175d57611742610212565b637d92a0f560e11b81528061175960048201610587565b0390fd5b8284919290919261454f565b611710565b61177f9061177a6145f5565b611acf565b61178761467d565b565b6117949036906112a5565b90565b60a01c90565b6117a96117ae91611797565b611496565b90565b6117bb905461179d565b90565b5f80fd5b5f80fd5b5f80fd5b903590600161018003813603038212156117e2570190565b6117be565b356117f1816108b1565b90565b356117fe81610da9565b90565b3561180b816104a8565b90565b35611818816104cb565b90565b634e487b7160e01b5f52601160045260245ffd5b61183e61184491939293610399565b92610399565b820180921161184f57565b61181b565b5f80fd5b60e01b90565b9050519061186b826104cb565b565b9060208282031261188657611883915f0161185e565b90565b61021c565b634e487b7160e01b5f52602260045260245ffd5b90600160028304921680156118bf575b60208310146118ba57565b61188b565b91607f16916118af565b5f5260205f2090565b905f92918054906118ec6118e58361189f565b8094611349565b916001811690815f146119435750600114611907575b505050565b61191491929394506118c9565b915f925b81841061192b57505001905f8080611902565b60018160209295939554848601520191019290611918565b92949550505060ff19168252151560200201905f8080611902565b9290611994916119876119a2969461197d60808801945f8901906107fa565b6020870190610a4f565b84820360408601526118d2565b916060818403910152611352565b90565b6119ad610212565b3d5f823e3d90fd5b60209181520190565b5f7f53656e642065746820496e73756666696369656e740000000000000000000000910152565b6119f260156020926119b5565b6119fb816119be565b0190565b611a149060208101905f8183039101526119e5565b90565b15611a1e57565b611a26610212565b62461bcd60e51b815280611a3c600482016119ff565b0390fd5b5f910312611a4a57565b61021c565b9692611abe95611a9d611acc9a989396611a93611aa794611a89611ab19860208f611a8261010082019e5f830190610a4f565b0190610a4f565b60408d019061072e565b60608b019061072e565b60808901906107fa565b60a0870190610a4f565b84820360c08601526118d2565b9160e0818403910152611352565b90565b611b4490611ae4611adf82611789565b613b55565b929092611af160066117b1565b90611b0b60c0611b058660208101906117ca565b016117e7565b611b236020611b1d87828101906117ca565b016117f4565b90611b3d6040611b378860208101906117ca565b01611801565b9293614764565b91611b56611b515f611703565b610625565b60206385fdd54291611b81611b7a610100611b7488868101906117ca565b0161180e565b859061182f565b90611bb9611b9d6060611b9789878101906117ca565b01611801565b94611bc460078a90611bad610212565b98899788968796611858565b86526004860161195e565b03915afa8015611d7157611c21915f91611d43575b50611c1a611c14611c0f611c083494611c02610100611bfc8b60208101906117ca565b0161180e565b9061182f565b869061182f565b610399565b91610399565b1015611a17565b611c32611c2d5f611703565b610625565b8263209afe5691349092611c556080611c4f8560208101906117ca565b01611801565b94611c6f60a0611c698660208101906117ca565b01611801565b97611cc66060611cc0611cb5611c9460e0611c8e8b60208101906117ca565b016117e7565b95611cb0610100611caa339c60208101906117ca565b0161180e565b61182f565b9a60208101906117ca565b01611801565b9060079091873b15611d3e575f99611cf097611cfb95611ce4610212565b9d8e9c8d9b8c9a611858565b8a5260048a01611a4f565b03925af18015611d3957611d0d575b50565b611d2c905f3d8111611d32575b611d24818361024a565b810190611a40565b5f611d0a565b503d611d1a565b6119a5565b611854565b611d64915060203d8111611d6a575b611d5c818361024a565b81019061186d565b5f611bd9565b503d611d52565b6119a5565b611d7f9061176e565b565b5f90565b5f80fd5b5f80fd5b5f80fd5b9035600160200382360303811215611dd257016020813591019167ffffffffffffffff8211611dcd576020820236038313611dc857565b611d89565b611d85565b611d8d565b60209181520190565b90565b50611df2906020810190610f30565b90565b611dfe9061043e565b9052565b50611e119060208101906104df565b90565b611e1d90610399565b9052565b50611e309060208101906108c5565b90565b611e3c90610722565b9052565b903560016101000382360303811215611e57570190565b611d8d565b50611e6b9060208101906104bc565b90565b611e779061049b565b9052565b9035600160200382360303811215611ebc57016020813591019167ffffffffffffffff8211611eb7576001820236038313611eb257565b611d89565b611d85565b611d8d565b60209181520190565b9190611ee481611edd81611ee995611ec1565b80956102ab565b61022c565b0190565b611fc691611fb8610100820192611f12611f095f830183611e5c565b5f850190611e6e565b611f2c611f226020830183611e5c565b6020850190611e6e565b611f46611f3c6040830183611e5c565b6040850190611e6e565b611f60611f566060830183611e5c565b6060850190611e6e565b611f7a611f706080830183611e5c565b6080850190611e6e565b611f94611f8a60a0830183611e5c565b60a0850190611e6e565b611fae611fa460c0830183611e5c565b60c0850190611e6e565b60e0810190611e7b565b9160e0818503910152611eca565b90565b61208a9161207c61207160e08301611fef611fe65f870187611de3565b5f860190611df5565b612009611fff6020870187611de3565b6020860190611df5565b6120236120196040870187611e02565b6040860190611e14565b61203d6120336060870187611e21565b6060860190611e33565b61205761204d6080870187611e21565b6080860190611e33565b61206460a0860186611e40565b84820360a0860152611eed565b9260c0810190611e40565b9060c0818403910152611eed565b90565b9061209791611fc9565b90565b9035600160e003823603038112156120b0570190565b611d8d565b60200190565b91816120c691611dd7565b90816120d760208302840194611de0565b92835f925b8484106120ec5750505050505090565b9091929394956020612117612111838560019503885261210c8b8861209a565b61208d565b986120b5565b9401940192949391906120dc565b90356001610180038236030381121561213c570190565b611d8d565b50612150906020810190610dbd565b90565b61215c90610aa6565b9052565b6122b1916122a2612296612279610180840161218a6121815f880188611de3565b5f870190611df5565b6121a461219a6020880188612141565b6020870190612153565b6121be6121b46040880188611e5c565b6040870190611e6e565b6121d86121ce6060880188611e5c565b6060870190611e6e565b6121f26121e86080880188611e5c565b6080870190611e6e565b61220c61220260a0880188611e5c565b60a0870190611e6e565b61222661221c60c0880188611e21565b60c0870190611e33565b61224061223660e0880188611e21565b60e0870190611e33565b61225c612251610100880188611e02565b610100870190611e14565b61226a610120870187611e7b565b90868303610120880152611eca565b612287610140860186611e7b565b90858303610140870152611eca565b92610160810190611e7b565b91610160818503910152611eca565b90565b6122f4916122e66122db604083016122ce5f860186611d91565b908583035f8701526120bb565b926020810190612125565b906020818403910152612160565b90565b61230c9160208201915f8184039101526122b4565b90565b903590600160200381360303821215612351570180359067ffffffffffffffff821161234c5760200191600182023603831361234757565b6117c6565b6117c2565b6117be565b919060408382031261239057612389906123706040610273565b9361237d825f83016104df565b5f8601526020016108c5565b6020830152565b610ef7565b906040828203126123ae576123ab915f01612356565b90565b61021c565b6123bd9051610399565b90565b6123c8611d81565b5060206123e56123f3836123da610212565b9283918583016122f7565b84820181038252038261024a565b916124826124275f61242161241961240e86888101906117ca565b61016081019061230f565b810190612395565b016123b3565b9361243260066117b1565b9061244b60c061244586888101906117ca565b016117e7565b6124628661245c87828101906117ca565b016117f4565b9061247b6040612475888a8101906117ca565b01611801565b9293614764565b926124f16124976124925f611703565b610625565b916124fc6124da60606124d46124ca6385fdd542956124c56101006124bf8c8e8101906117ca565b0161180e565b61182f565b97898101906117ca565b01611801565b9660076124e5610212565b98899788968796611858565b86526004860161195e565b03915afa908115612540575f91612512575b5090565b612533915060203d8111612539575b61252b818361024a565b81019061186d565b5f61250e565b503d612521565b6119a5565b5190565b90565b61256061255b61256592612549565b6105fa565b610399565b90565b61257c61257761258192610399565b6105fa565b61049b565b90565b612590612595916116ea565b6106d8565b90565b6125a29054612584565b90565b6125ae906105fd565b90565b6125ba906125a5565b90565b6125c690610619565b90565b60209181520190565b60200190565b6125f7612600602093612605936125ee81610344565b93848093611ec1565b9586910161034d565b61022c565b0190565b6126a19160e06101008201926126255f8201515f850190611e6e565b61263760208201516020850190611e6e565b61264960408201516040850190611e6e565b61265b60608201516060850190611e6e565b61266d60808201516080850190611e6e565b61267f60a082015160a0850190611e6e565b61269160c082015160c0850190611e6e565b01519060e08184039101526125d8565b90565b61272c9160c061271b60e083016126c15f8601515f860190611df5565b6126d360208601516020860190611df5565b6126e560408601516040860190611e14565b6126f760608601516060860190611e33565b61270960808601516080860190611e33565b60a085015184820360a0860152612609565b9201519060c0818403910152612609565b90565b90612739916126a4565b90565b60200190565b9061275661274f83612545565b80926125c9565b9081612767602083028401946125d2565b925f915b83831061277a57505050505090565b9091929394602061279c6127968385600195038752895161272f565b9761273c565b930193019193929061276b565b6127be9160208201915f818403910152612742565b90565b5f90565b606090565b905051906127d7826108b1565b565b91906040838203126128135761280c906127f36040610273565b93612800825f830161185e565b5f8601526020016127ca565b6020830152565b610ef7565b906040828203126128315761282e915f016127d9565b90565b61021c565b612840905161043e565b90565b90565b61285a61285561285f92612843565b6105fa565b61043e565b90565b90565b61287961287461287e92612862565b6105fa565b61043e565b90565b61289561289061289a92612549565b6105fa565b61043e565b90565b6128a690610619565b90565b906128bb6128b683610288565b610273565b918252565b3d5f146128db576128d03d6128a9565b903d5f602084013e5b565b6128e36127c5565b906128d9565b6128f39051610722565b90565b6129015f8092610348565b0190565b61290e906128f6565b90565b151590565b61291f90612911565b9052565b91612959939161294b9161293e60608601925f870190612916565b8482036020860152611352565b916040818403910152611352565b90565b9061296a9190810190611314565b6129765f820151612545565b6129886129825f61254c565b91610399565b03612b8e575b6129966127c1565b5061299f6127c5565b505f806129c56101606020850151015160206129ba82610344565b818301019101612818565b6129d482602086015101612836565b6129e76129e160fe612846565b9161043e565b148214612a6e57612a05826129fe602084016128e9565b92016123b3565b612a0d610212565b9081612a1881612905565b03925af1906101406020612a2a6128c0565b935b93920151015191612a697fc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb39807944693612a60610212565b93849384612923565b0390a1565b612a7d82602086015101612836565b612a90612a8a60ff612865565b9161043e565b148214612ae757612abe82612ab7612ab2600a612aac46612568565b906106c2565b612598565b92016123b3565b6101406020860151015190602082019151925af1906101406020612ae06128c0565b935b612a2c565b612af682602086015101612836565b612b08612b0284612881565b9161043e565b148214612b5257612b2682612b1f602084016128e9565b92016123b3565b612b2e610212565b9081612b3981612905565b03925af1906101406020612b4b6128c0565b935b612ae2565b612b6682612b5f3061289d565b92016123b3565b6101406020860151015190602082019151925af1906101406020612b886128c0565b93612b4d565b612bba612bb5612bb0612bab600a612ba546612568565b906106c2565b612598565b6125b1565b6125bd565b6377c2719c5f830151823b15612c3957612bf392612be85f8094612bdc610212565b96879586948593611858565b8352600483016127a9565b03925af18015612c3457612c08575b5061298e565b612c27905f3d8111612c2d575b612c1f818361024a565b810190611a40565b5f612c02565b503d612c15565b6119a5565b611854565b9181612c49916125c9565b9081612c5a60208302840194611de0565b92835f925b848410612c6f5750505050505090565b9091929394956020612c9a612c948385600195038852612c8f8b8861209a565b61208d565b986120b5565b940194019294939190612c5f565b9091612cbf9260208301925f818503910152612c3e565b90565b612cce612cd3916116ea565b610d37565b90565b612ce09054612cc2565b90565b60181c90565b612cf5612cfa91612ce3565b610a0c565b90565b612d079054612ce9565b90565b919390612d6b91612d44602095612d1f611d81565b50612d36612d2b610212565b938492898401612ca8565b87820181038252038261024a565b612d4e60066117b1565b91612d596008612cd6565b90612d646008612cfd565b9293614764565b92612da2612d80612d7b5f611703565b610625565b91612dad6385fdd5429194966007612d96610212565b98899788968796611858565b86526004860161195e565b03915afa908115612df1575f91612dc3575b5090565b612de4915060203d8111612dea575b612ddc818361024a565b81019061186d565b5f612dbf565b503d612dd2565b6119a5565b5f90565b612e02612df6565b90565b5f90565b612e11612e05565b90565b612e1c612e05565b90565b612e276147b3565b612e2f612e59565b565b612e45612e40612e4a92612549565b6105fa565b6105ef565b90565b612e5690612e31565b90565b612e6a612e655f612e4d565b614808565b565b612e74612e1f565b565b612e819136916102b6565b90565b92612ea6612eac9295612eb29795612e9a611d81565b50969492909592612e76565b93612e76565b926148b0565b90565b5f90565b612ec1612eb5565b50612ecc6002612598565b90565b5f7f4d45500000000000000000000000000000000000000000000000000000000000910152565b612f0360036020926119b5565b612f0c81612ecf565b0190565b612f259060208101905f818303910152612ef6565b90565b15612f2f57565b612f37610212565b62461bcd60e51b815280612f4d60048201612f10565b0390fd5b90612f9594939291612f90612f6546612568565b612f8a612f84612f7f612f7a3394600a6106c2565b612598565b610722565b91610722565b14612f28565b612ff0565b565b15612f9e57565b5f80fd5b60581c90565b612fb4612fb991612fa2565b610a0c565b90565b612fc69054612fa8565b90565b60981c90565b612fdb612fe091612fc9565b610a0c565b90565b612fed9054612fcf565b90565b929361307f9161305b9161304f61300760066117b1565b836130126008612cd6565b61301c6008612cfd565b918b9361304a879561303b61302f610212565b97889260208401612ca8565b6020820181038252038661024a565b614764565b96869286919293612d0a565b61307861307261306d3493869061182f565b610399565b91610399565b1015612f97565b61309061308b5f611703565b610625565b9063209afe5690349291926130a56008612fbc565b936130b06008612fe3565b966130bb6009612598565b903394979060079091873b15613137575f996130e9976130f4956130dd610212565b9d8e9c8d9b8c9a611858565b8a5260048a01611a4f565b03925af1801561313257613106575b50565b613125905f3d811161312b575b61311d818361024a565b810190611a40565b5f613103565b503d613113565b6119a5565b611854565b9061314994939291612f51565b565b9061315e92916131596147b3565b61324f565b565b5f1b90565b9061317362ffffff91613160565b9181191691161790565b61319161318c61319692610aa6565b6105fa565b610aa6565b90565b90565b906131b16131ac6131b89261317d565b613199565b8254613165565b9055565b60181b90565b906131d86affffffffffffffff000000916131bc565b9181191691161790565b90565b906131fa6131f5613201926106a6565b6131e2565b82546131c2565b9055565b9061321660018060a01b0391613160565b9181191691161790565b61322990610619565b90565b90565b9061324461323f61324b92613220565b61322c565b8254613205565b9055565b613270929161326261326992600861319c565b60086131e5565b600961322f565b565b9061327d929161314b565b565b613287612eb5565b90565b90565b6132a161329c6132a69261328a565b6105fa565b61043e565b90565b90565b6132c06132bb6132c5926132a9565b6105fa565b61043e565b90565b6132d3610120610273565b90565b5f90565b5f90565b5f90565b5f90565b5f90565b6132f26132c8565b90602080808080808080808a6133066132d6565b8152016133116132da565b81520161331c6132de565b8152016133276132e2565b8152016133326132e2565b81520161333d6132e2565b8152016133486132e2565b8152016133536132e6565b81520161335e6132e6565b81525050565b61336c6132ea565b90565b634e487b7160e01b5f52603260045260245ffd5b9061338d82612545565b81101561339e576020809102010190565b61336f565b6133af6133b4916116ea565b6103f9565b90565b6133c190546133a3565b90565b905051906133d182610f1c565b565b905051906133e082610da9565b565b6133eb816105ef565b036133f257565b5f80fd5b90505190613403826133e2565b565b9190610120838203126134c1576134b990613421610120610273565b9361342e825f83016133c4565b5f86015261343f82602083016133d3565b602086015261345182604083016133f6565b604086015261346382606083016127ca565b606086015261347582608083016127ca565b60808601526134878260a083016127ca565b60a08601526134998260c083016127ca565b60c08601526134ab8260e0830161185e565b60e08601526101000161185e565b610100830152565b610ef7565b90610120828203126134e0576134dd915f01613405565b90565b61021c565b63ffffffff1690565b63ffffffff60e01b1690565b61350e613509613513926134e5565b611858565b6134ee565b90565b61351f906105ef565b9052565b90610100806135c49361353c5f8201515f860190611df5565b61354e60208201516020860190612153565b61356060408201516040860190613516565b61357260608201516060860190611e33565b61358460808201516080860190611e33565b61359660a082015160a0860190611e33565b6135a860c082015160c0860190611e33565b6135ba60e082015160e0860190611e14565b0151910190611e14565b565b91906135da905f6101208501940190613523565b565b6135e660e0610273565b90565b606090565b6135f66135dc565b906020808080808080886136086132d6565b8152016136136132e6565b81520161361e6132e6565b8152016136296135e9565b8152016136346132e2565b81520161363f6132e2565b81520161364a6132e6565b81525050565b6136586135ee565b90565b67ffffffffffffffff81116136735760208091020190565b610236565b9092919261368d6136888261365b565b610273565b93818552602080860192028301928184116136ca57915b8383106136b15750505050565b602080916136bf84866127ca565b8152019201916136a4565b6104f2565b9080601f830112156136ed578160206136ea93519101613678565b90565b610224565b91909160e08184031261379f5761370960e0610273565b92613716815f84016133c4565b5f850152613727816020840161185e565b6020850152613739816040840161185e565b604085015260608201519167ffffffffffffffff831161379a57613762826137939483016136cf565b606086015261377482608083016127ca565b60808601526137868260a083016127ca565b60a086015260c00161185e565b60c0830152565b610efb565b610ef7565b906020828203126137d4575f82015167ffffffffffffffff81116137cf576137cc92016136f2565b90565b610220565b61021c565b5190565b60209181520190565b60200190565b906137f981602093611e33565b0190565b60200190565b9061382061381a613813846137d9565b80936137dd565b926137e6565b905f5b8181106138305750505090565b90919261384961384360019286516137ec565b946137fd565b9101919091613823565b906138d79060c0806138a860e084016138725f8801515f870190611df5565b61388460208801516020870190611e14565b61389660408801516040870190611e14565b60608701518582036060870152613803565b946138bb60808201516080860190611e33565b6138cd60a082015160a0860190611e33565b0151910190611e14565b90565b6138ef9160208201915f818403910152613853565b90565b6138fc6040610273565b90565b6139076138f2565b90602080836139146132e6565b81520161391f6132e2565b81525050565b61392d6138ff565b90565b61393a9051610aa6565b90565b613947905161049b565b90565b613955610180610273565b90565b906139629061043e565b9052565b9061397090610aa6565b9052565b9061397e9061049b565b9052565b9061398c90610722565b9052565b9061399a90610399565b9052565b52565b6139ab6040610273565b90565b52565b52565b906139c86139c183612545565b8092611dd7565b90816139d9602083028401946125d2565b925f915b8383106139ec57505050505090565b90919293946020613a0e613a088385600195038752895161272f565b9761273c565b93019301919392906139dd565b613b0991610160613af7613ae36101808401613a3d5f8701515f870190611df5565b613a4f60208701516020870190612153565b613a6160408701516040870190611e6e565b613a7360608701516060870190611e6e565b613a8560808701516080870190611e6e565b613a9760a087015160a0870190611e6e565b613aa960c087015160c0870190611e33565b613abb60e087015160e0870190611e33565b613acf610100870151610100870190611e14565b6101208601518582036101208701526125d8565b6101408501518482036101408601526125d8565b920151906101608184039101526125d8565b90565b613b3a916020613b29604083015f8501518482035f8601526139b4565b920151906020818403910152613a1b565b90565b613b529160208201915f818403910152613b0c565b90565b90613b5e611d81565b50613b676127c5565b50613b70611d81565b613b786127c5565b50613b816127c5565b613b905f602086015101612836565b613ba2613b9c5f612881565b9161043e565b145f14613f3c575050613bb3613925565b50613be6613be1600c60e060a0613bd65f880151613bd05f61254c565b90613383565b5101510151906103e9565b6133b7565b613bf8613bf25f612881565b9161043e565b145f14613e2557613c3760e060a0613c1c5f860151613c165f61254c565b90613383565b51015101516020613c2c82610344565b818301019101612818565b91613c635f60e060a0613c5683860151613c508561254c565b90613383565b5101510151945b016123b3565b92610140602083015101515b82602001515f01613c7f90612836565b918360200151602001613c9190613930565b918460200151604001613ca39061393d565b8560200151606001613cb49061393d565b8660200151608001613cc59061393d565b876020015160a001613cd69061393d565b886020015160c001613ce7906128e9565b91896020015160e001613cf9906128e9565b938a6020015161010001613d0c906123b3565b958b602001516101200151979899613d2261394a565b9b5f8d0190613d3091613958565b60208c0190613d3e91613966565b60408b0190613d4c91613974565b60608a0190613d5a91613974565b6080890190613d6891613974565b60a0880190613d7691613974565b60c0870190613d8491613982565b60e0860190613d9291613982565b610100850190613da191613990565b610120840190613db09161399e565b610140830190613dbf9161399e565b610160820190613dce9161399e565b905f015190613ddb6139a1565b915f830190613de9916139ae565b6020820190613df7916139b1565b613dff610212565b80916020820190613e0f91613b3d565b602082018103825203613e22908261024a565b90565b613e57613e52600c60e060a0613e475f880151613e415f61254c565b90613383565b5101510151906103e9565b6133b7565b613e6a613e64600161328d565b9161043e565b1480613efc575b5f14613ed957613eaf60e060c0613e945f860151613e8e5f61254c565b90613383565b51015101516020613ea482610344565b818301019101612818565b91613c635f60e060c0613ece83860151613ec88561254c565b90613383565b510151015194613c5d565b613ee1610212565b6382077b6960e01b815280613ef860048201610587565b0390fd5b50613f2460e060c0613f1a5f860151613f145f61254c565b90613383565b5101510151610344565b613f36613f305f61254c565b91610399565b11613e71565b613f4d5f6020869496015101612836565b613f60613f5a600161328d565b9161043e565b145f1461412e57613f6f613650565b50613fa2613f9d600c60e060a0613f925f880151613f8c5f61254c565b90613383565b5101510151906103e9565b6133b7565b613fb4613fae5f612881565b9161043e565b145f14614037576004614031613ff860e060a0613fdd5f880151613fd75f61254c565b90613383565b51015101516020613fed82610344565b8183010191016137a4565b5b6140226140096374a736b96134fa565b91614012610212565b94859360208501908152016138da565b6020820181038252038261024a565b5b613c6f565b614069614064600c60e060a06140595f8801516140535f61254c565b90613383565b5101510151906103e9565b6133b7565b61407c614076600161328d565b9161043e565b14806140ee575b5f146140cb5760046140316140c660e060c06140ab5f8801516140a55f61254c565b90613383565b510151015160206140bb82610344565b8183010191016137a4565b613ff9565b6140d3610212565b6382077b6960e01b8152806140ea60048201610587565b0390fd5b5061411660e060c061410c5f8601516141065f61254c565b90613383565b5101510151610344565b6141286141225f61254c565b91610399565b11614083565b61413d5f602084015101612836565b61415061414a60026132ac565b9161043e565b145f1461431e5761415f613364565b5061419261418d600c60e060a06141825f88015161417c5f61254c565b90613383565b5101510151906103e9565b6133b7565b6141a461419e5f612881565b9161043e565b145f146142275760046142216141e860e060a06141cd5f8801516141c75f61254c565b90613383565b510151015160206141dd82610344565b8183010191016134c6565b5b6142126141f963bc33e6206134fa565b91614202610212565b94859360208501908152016135c6565b6020820181038252038261024a565b5b614032565b614259614254600c60e060a06142495f8801516142435f61254c565b90613383565b5101510151906103e9565b6133b7565b61426c614266600161328d565b9161043e565b14806142de575b5f146142bb5760046142216142b660e060c061429b5f8801516142955f61254c565b90613383565b510151015160206142ab82610344565b8183010191016134c6565b6141e9565b6142c3610212565b6382077b6960e01b8152806142da60048201610587565b0390fd5b5061430660e060c06142fc5f8601516142f65f61254c565b90613383565b5101510151610344565b6143186143125f61254c565b91610399565b11614273565b5091506143305f602084015101612836565b61434361433d60fe612846565b9161043e565b145f14614392576143775f61437161016060208601510151602061436682610344565b818301019101612818565b016123b3565b9161016060208201510151610140602083015101515b614222565b6143a15f602084015101612836565b6143b46143ae60ff612865565b9161043e565b145f14614402576143e85f6143e26101606020860151015160206143d782610344565b818301019101612818565b016123b3565b91610160602082015101516101406020830151015161438d565b61440a610212565b634157854360e11b81528061442160048201610587565b0390fd5b90614437916144326147b3565b614439565b565b61444861444d9291600a6106c2565b61322f565b565b9061445991614425565b565b61446c906144676147b3565b61446e565b565b8061448961448361447e5f612e4d565b610722565b91610722565b146144995761449790614808565b565b6144c36144a55f612e4d565b6144ad610212565b918291631e4fbdf760e01b83526004830161073b565b0390fd5b6144d09061445b565b565b93929190336144fa6144f46144ef6144ea6001611703565b610625565b610722565b91610722565b0361450a576145089461452d565b565b614512610212565b637d92a0f560e11b81528061452960048201610587565b0390fd5b9161453e94929391909192936149e4565b565b9061454d949392916144d2565b565b614557610212565b637fca089960e11b81528061456e60048201610587565b0390fd5b61457e614583916116ea565b610b3a565b90565b6145909054614572565b90565b6145a76145a26145ac926132a9565b6105fa565b610399565b90565b6145b96002614593565b90565b906145c85f1991613160565b9181191691161790565b90565b906145ea6145e56145f19261145e565b6145d2565b82546145bc565b9055565b6145ff6003614586565b61461861461261460d6145af565b610399565b91610399565b146146315761462f6146286145af565b60036145d5565b565b614639610212565b633ee5aeb560e01b81528061465060048201610587565b0390fd5b61466861466361466d9261328a565b6105fa565b610399565b90565b61467a6001614654565b90565b61468f614688614670565b60036145d5565b565b61469a906105fd565b90565b6146b16146ac6146b6926105ef565b6105fa565b610399565b90565b90565b6146c86146cd916109ac565b6146b9565b9052565b60e81b90565b6146e0906146d1565b90565b6146ef6146f491610aa6565b6146d7565b9052565b60c01b90565b614707906146f8565b90565b61471661471b9161049b565b6146fe565b9052565b600893614752602061476198979561474a60018661474260039861475a9a6146bc565b01809261039f565b0180926146e3565b01809261470a565b0190610370565b90565b926147a19192946147896147846147b09661477d6127c5565b5095614691565b61469d565b95919091614795610212565b9687956020870161471f565b6020820181038252038261024a565b90565b6147bb612eb9565b6147d46147ce6147c9614cca565b610722565b91610722565b036147db57565b6148046147e6614cca565b6147ee610212565b91829163118cdaa760e01b83526004830161073b565b0390fd5b6148126002612598565b61481d82600261322f565b9061485161484b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093613220565b91613220565b9161485a610212565b8061486481610587565b0390a3565b929061489f916148926148ad969461488860808801945f8901906107fa565b6020870190610a4f565b8482036040860152611352565b916060818403910152611352565b90565b926148f2602093946148c0611d81565b506148fd6148d56148d05f611703565b610625565b936385fdd5429295976148e6610212565b98899788968796611858565b865260048601614869565b03915afa908115614941575f91614913575b5090565b614934915060203d811161493a575b61492c818361024a565b81019061186d565b5f61490f565b503d614922565b6119a5565b61495a61495561495f92610399565b6105fa565b6105ef565b90565b5f7f496e76616c696420636f6e747261637400000000000000000000000000000000910152565b61499660106020926119b5565b61499f81614962565b0190565b6149b89060208101905f818303910152614989565b90565b156149c257565b6149ca610212565b62461bcd60e51b8152806149e0600482016149a3565b0390fd5b50614a32939291614a23614a1d614a18614a13614a0d614a08614a2997600a6106c2565b612598565b94614946565b610619565b610722565b91610722565b146149bb565b90810190611314565b614a3e5f820151612545565b614a50614a4a5f61254c565b91610399565b03614c1a575b614a5e6127c1565b50614a676127c5565b505f80614a8d610160602085015101516020614a8282610344565b818301019101612818565b614a9c82602086015101612836565b614aaf614aa960ff612865565b9161043e565b148214614b4357614add82614ad6614ad1600a614acb46612568565b906106c2565b612598565b92016123b3565b6101406020860151015190602082019151925af1906101406020614aff6128c0565b935b93920151015191614b3e7fc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb39807944693614b35610212565b93849384612923565b0390a1565b614b5282602086015101612836565b614b64614b5e84612881565b9161043e565b148015614bf1575b8214614bb557614b8982614b82602084016128e9565b92016123b3565b614b91610212565b9081614b9c81612905565b03925af1906101406020614bae6128c0565b935b614b01565b614bc982614bc23061289d565b92016123b3565b6101406020860151015190602082019151925af1906101406020614beb6128c0565b93614bb0565b50614c0182602086015101612836565b614c14614c0e60fe612846565b9161043e565b14614b6c565b614c46614c41614c3c614c37600a614c3146612568565b906106c2565b612598565b6125b1565b6125bd565b6377c2719c5f830151823b15614cc557614c7f92614c745f8094614c68610212565b96879586948593611858565b8352600483016127a9565b03925af18015614cc057614c94575b50614a56565b614cb3905f3d8111614cb9575b614cab818361024a565b810190611a40565b5f614c8e565b503d614ca1565b6119a5565b611854565b614cd2612eb5565b50339056fea264697066735822122002cde2778813f012f1ea5fb6155706165c3bba885c6b22a199ceb90113bf705a64736f6c63430008180033",
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

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0x3169eb13.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterCaller) FetchUserOmniMessageFee(opts *bind.CallOpts, params BaseStructCrossMessageParams) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "fetchUserOmniMessageFee", params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0x3169eb13.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterSession) FetchUserOmniMessageFee(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _SyncRouter.Contract.FetchUserOmniMessageFee(&_SyncRouter.CallOpts, params)
}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0x3169eb13.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterCallerSession) FetchUserOmniMessageFee(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _SyncRouter.Contract.FetchUserOmniMessageFee(&_SyncRouter.CallOpts, params)
}

// GetUserOmniEncodeMessage is a free data retrieval call binding the contract method 0xc967f886.
//
// Solidity: function getUserOmniEncodeMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) cmp) view returns(uint256, bytes)
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

// GetUserOmniEncodeMessage is a free data retrieval call binding the contract method 0xc967f886.
//
// Solidity: function getUserOmniEncodeMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) cmp) view returns(uint256, bytes)
func (_SyncRouter *SyncRouterSession) GetUserOmniEncodeMessage(cmp BaseStructCrossMessageParams) (*big.Int, []byte, error) {
	return _SyncRouter.Contract.GetUserOmniEncodeMessage(&_SyncRouter.CallOpts, cmp)
}

// GetUserOmniEncodeMessage is a free data retrieval call binding the contract method 0xc967f886.
//
// Solidity: function getUserOmniEncodeMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) cmp) view returns(uint256, bytes)
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

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) cmp) payable returns()
func (_SyncRouter *SyncRouterTransactor) SendUserOmniMessage(opts *bind.TransactOpts, cmp BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "sendUserOmniMessage", cmp)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) cmp) payable returns()
func (_SyncRouter *SyncRouterSession) SendUserOmniMessage(cmp BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendUserOmniMessage(&_SyncRouter.TransactOpts, cmp)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) cmp) payable returns()
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
