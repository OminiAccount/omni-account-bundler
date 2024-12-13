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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vizingPad\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_Hook\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExecutionCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LandingPadAccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplement\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"packHookMessage\",\"type\":\"bytes\"}],\"name\":\"ReceiveTouchHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"VizingSwapEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"DataExcecuteNumber\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Hook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LandingPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LaunchPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LockWay\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"MirrorEntryPoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"newGaslimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"newGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newRelayer\",\"type\":\"address\"}],\"name\":\"changeDefaultSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBridgeMode\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGasPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGaslimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destChainid\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"additionParams\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"estimateVizingGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vizingGasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"fetchOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"fetchUserOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"cmp\",\"type\":\"tuple\"}],\"name\":\"getUserOmniEncodeMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveExtraMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveStandardMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selectedRelayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"sendOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"cmp\",\"type\":\"tuple\"}],\"name\":\"sendUserOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"}],\"name\":\"setMirrorEntryPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thisRelayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523462000038576200001f6200001862000131565b9162000694565b620000296200003e565b614cba62000a1a8239614cba90f35b62000044565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90620000729062000048565b810190811060018060401b038211176200008b57604052565b62000052565b90620000a8620000a06200003e565b928362000066565b565b5f80fd5b60018060a01b031690565b620000c490620000ae565b90565b620000d281620000b9565b03620000da57565b5f80fd5b90505190620000ed82620000c7565b565b90916060828403126200012b57620001286200010e845f8501620000de565b936200011e8160208601620000de565b93604001620000de565b90565b620000aa565b62000154620056d480380380620001488162000091565b928339810190620000ef565b909192565b60a01b90565b906200017060ff60a01b9162000159565b9181191691161790565b90565b60ff60f81b1690565b60f81b90565b620001a56200019f620001ab926200017a565b62000186565b6200017d565b90565b60f81c90565b620001bf90620001ae565b90565b90620001dc620001d6620001e4926200018c565b620001b4565b82546200015f565b9055565b90565b90565b90565b6200020a620002046200021092620001e8565b620001ee565b620001eb565b90565b60018060401b03811162000232576200022e60209162000048565b0190565b62000052565b906200024e620002488362000213565b62000091565b918252565b369037565b9062000283620002688362000238565b9260208062000278869362000213565b920191039062000253565b565b5190565b634e487b7160e01b5f52602260045260245ffd5b9060016002830492168015620002c0575b6020831014620002ba57565b62000289565b91607f1691620002ae565b5f5260205f2090565b601f602091010490565b1b90565b9190600862000301910291620002fa5f1984620002de565b92620002de565b9181191691161790565b620003246200031e6200032a92620001eb565b620001ee565b620001eb565b90565b90565b91906200034b6200034562000354936200030b565b6200032d565b908354620002e2565b9055565b5f90565b62000372916200036b62000358565b9162000330565b565b5b81811062000381575050565b80620003905f6001936200035c565b0162000375565b9190601f8111620003a8575b505050565b620003b7620003e293620002cb565b906020620003c584620002d4565b83019310620003eb575b620003da90620002d4565b019062000374565b5f8080620003a3565b9150620003da81929050620003cf565b1c90565b9062000411905f1990600802620003fb565b191690565b816200042291620003ff565b906002021790565b90620004368162000285565b9060018060401b03821162000507576200045d826200045685546200029d565b8562000397565b602090601f8311600114620004965791809162000484935f9262000489575b505062000416565b90555b565b90915001515f806200047c565b601f19831691620004a785620002cb565b925f5b818110620004ee57509160029391856001969410620004d1575b5050500201905562000487565b620004e3910151601f841690620003ff565b90555f8080620004c4565b91936020600181928787015181550195019201620004aa565b62000052565b9062000519916200042a565b565b5f1b90565b906200053062ffffff916200051b565b9181191691161790565b90565b62ffffff1690565b6200055e6200055862000564926200053a565b620001ee565b6200053d565b90565b90565b90620005846200057e6200058c9262000545565b62000567565b825462000520565b9055565b60181b90565b90620005ad6301000000600160581b039162000590565b9181191691161790565b90565b60018060401b031690565b620005de620005d8620005e492620005b7565b620001ee565b620005ba565b90565b90565b9062000604620005fe6200060c92620005c5565b620005e7565b825462000596565b9055565b906200062360018060a01b03916200051b565b9181191691161790565b62000646620006406200064c92620000ae565b620001ee565b620000ae565b90565b6200065a906200062d565b90565b62000668906200064f565b90565b90565b90620006886200068262000690926200065d565b6200066b565b825462000610565b9055565b6200070a9291620006aa6200070292336200070c565b620006b860016006620001c2565b620006d9620006d1620006cb5f620001f1565b62000258565b60076200050d565b620006e96207a12060086200056a565b620006fa633b9aca006008620005ea565b60056200066e565b60066200066e565b565b9062000718916200071a565b565b90620007269162000728565b565b906200073491620007a5565b565b6200074f6200074962000755926200017a565b620001ee565b620001eb565b90565b62000764600162000736565b90565b90620007755f19916200051b565b9181191691161790565b906200079962000793620007a1926200030b565b6200032d565b825462000767565b9055565b90620007b1916200081f565b620007c7620007bf62000758565b60036200077f565b565b620007e2620007dc620007e892620001e8565b620001ee565b620000ae565b90565b620007f690620007c9565b90565b6200080490620000b9565b9052565b91906200081d905f60208501940190620007f9565b565b906200082b9062000893565b806200084c620008456200083f5f620007eb565b620000b9565b91620000b9565b146200085f576200085d9062000928565b565b6200088f6200086e5f620007eb565b620008786200003e565b918291631e4fbdf760e01b83526004830162000808565b0390fd5b6200089e90620008a0565b565b620008ac9080620008ae565b565b620008bd620008c392620008c5565b620009e8565b565b620008d090620008d2565b565b620008dd90620008df565b565b620008ea9062000a01565b565b5f1c90565b60018060a01b031690565b6200090b6200091191620008ec565b620008f1565b90565b620009209054620008fc565b90565b5f0190565b62000934600262000914565b620009418260026200066e565b9062000979620009727f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936200065d565b916200065d565b91620009846200003e565b80620009908162000923565b0390a3565b620009a0906200062d565b90565b620009ae9062000995565b90565b620009bc9062000995565b90565b90565b90620009dc620009d6620009e492620009b1565b620009bf565b825462000610565b9055565b620009f7620009ff91620009a3565b6001620009c2565b565b62000a1062000a1791620009a3565b5f620009c2565b56fe60806040526004361015610015575b3661167657005b61001f5f356101fc565b80622232e3146101f757806273b555146101f25780630e82845d146101ed5780631490a358146101e85780631b3348f0146101e35780633169eb13146101de5780633eba5864146101d957806345636279146101d45780635ad3ad06146101cf5780635aeb4d77146101ca5780635e45da23146101c557806360b43029146101c0578063715018a6146101bb57806376c81312146101b65780638da5cb5b146101b15780638ea23ddf146101ac57806393afae93146101a757806399e581fa146101a2578063a834a9001461019d578063ad5c464814610198578063b0cfd4d214610193578063c5f378d81461018e578063c967f88614610189578063d5f3e00814610184578063de4b1d6d1461017f578063de8aeda01461017a578063e0b838e914610175578063e7b4294c14610170578063f2fde38b1461016b5763f31489250361000e57611649565b611590565b61153d565b6114f8565b6114b6565b611471565b61139c565b611339565b610e52565b610e0e565b610dd9565b610d96565b610d04565b610c9a565b610c65565b610c21565b610be5565b610b37565b610b02565b610a95565b610a60565b610a01565b610967565b610903565b61080c565b6107c1565b610740565b610643565b61057c565b610456565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906102449061021c565b810190811067ffffffffffffffff82111761025e57604052565b610226565b9061027661026f610202565b928361023a565b565b67ffffffffffffffff81116102965761029260209161021c565b0190565b610226565b90825f939282370152565b909291926102bb6102b682610278565b610263565b938185526020850190828401116102d7576102d59261029b565b565b610218565b9080601f830112156102fa578160206102f7933591016102a6565b90565b610214565b9060208282031261032f575f82013567ffffffffffffffff811161032a5761032792016102dc565b90565b610210565b61020c565b5190565b905090565b5f5b83811061034f575050905f910152565b80602091830151818501520161033f565b61038561037c9260209261037381610334565b94858093610338565b9384910161033d565b0190565b90565b90565b61039b6103a091610389565b61038c565b9052565b6103b46103bb9160209493610360565b809261038f565b0190565b6103d36103ca610202565b928392836103a4565b03902090565b6103e2916103bf565b90565b1c90565b60ff1690565b6103ff90600861040493026103e5565b6103e9565b90565b9061041291546103ef565b90565b61042b90610426600c915f926103d9565b610407565b90565b60ff1690565b61043d9061042e565b9052565b9190610454905f60208501940190610434565b565b346104865761048261047161046c3660046102ff565b610415565b610479610202565b91829182610441565b0390f35b610208565b67ffffffffffffffff1690565b6104a18161048b565b036104a857565b5f80fd5b905035906104b982610498565b565b6104c481610389565b036104cb57565b5f80fd5b905035906104dc826104bb565b565b5f80fd5b5f80fd5b909182601f830112156105205781359167ffffffffffffffff831161051b57602001926001830284011161051657565b6104e2565b6104de565b610214565b916060838303126105725761053c825f85016104ac565b9261054a83602083016104cf565b92604082013567ffffffffffffffff811161056d5761056992016104e6565b9091565b610210565b61020c565b5f0190565b610587366004610525565b929190916116f9565b5f91031261059a57565b61020c565b60018060a01b031690565b6105ba9060086105bf93026103e5565b61059f565b90565b906105cd91546105aa565b90565b6105dc60015f906105c2565b90565b60018060a01b031690565b90565b6106016105fc610606926105df565b6105ea565b6105df565b90565b610612906105ed565b90565b61061e90610609565b90565b61062a90610615565b9052565b9190610641905f60208501940190610621565b565b3461067357610653366004610590565b61066f61065e6105d0565b610666610202565b9182918261062e565b0390f35b610208565b906020828203126106915761068e915f016104ac565b90565b61020c565b6106aa6106a56106af9261048b565b6105ea565b61048b565b90565b906106bc90610696565b5f5260205260405f2090565b60018060a01b031690565b6106e39060086106e893026103e5565b6106c8565b90565b906106f691546106d3565b90565b61070f9061070a600a915f926106b2565b6106eb565b90565b61071b906105df565b90565b61072790610712565b9052565b919061073e905f6020850194019061071e565b565b346107705761076c61075b610756366004610678565b6106f9565b610763610202565b9182918261072b565b0390f35b610208565b5f80fd5b908160409103126107875790565b610775565b906020828203126107bc575f82013567ffffffffffffffff81116107b7576107b49201610779565b90565b610210565b61020c565b6107d46107cf36600461078c565b611d06565b6107dc610202565b806107e681610577565b0390f35b6107f390610389565b9052565b919061080a905f602085019401906107ea565b565b3461083c5761083861082761082236600461078c565b612350565b61082f610202565b918291826107f7565b0390f35b610208565b61084a81610712565b0361085157565b5f80fd5b9050359061086282610841565b565b909182601f8301121561089e5781359167ffffffffffffffff831161089957602001926020830284011161089457565b6104e2565b6104de565b610214565b906080828203126108fe576108ba815f84016104ac565b926108c88260208501610855565b926108d683604083016104cf565b92606082013567ffffffffffffffff81116108f9576108f59201610864565b9091565b610210565b61020c565b34610937576109336109226109193660046108a3565b939290926125aa565b61092a610202565b918291826107f7565b0390f35b610208565b60ff60f81b1690565b61094e9061093c565b9052565b9190610965905f60208501940190610945565b565b3461099757610977366004610590565b61099361098261269a565b61098a610202565b91829182610952565b0390f35b610208565b67ffffffffffffffff1690565b6109b99060086109be93026103e5565b61099c565b90565b906109cc91546109a9565b90565b6109dc6008600b906109c1565b90565b6109e89061048b565b9052565b91906109ff905f602085019401906109df565b565b34610a3157610a11366004610590565b610a2d610a1c6109cf565b610a24610202565b918291826109ec565b0390f35b610208565b62ffffff1690565b610a4790610a36565b9052565b9190610a5e905f60208501940190610a3e565b565b34610a9057610a70366004610590565b610a8c610a7b6126a9565b610a83610202565b91829182610a4b565b0390f35b610208565b34610ac557610aa5366004610590565b610ac1610ab06126b4565b610ab8610202565b91829182610a4b565b0390f35b610208565b90565b610add906008610ae293026103e5565b610aca565b90565b90610af09154610acd565b90565b610aff60045f90610ae5565b90565b34610b3257610b12366004610590565b610b2e610b1d610af3565b610b25610202565b918291826107f7565b0390f35b610208565b34610b6557610b47366004610590565b610b4f61270f565b610b57610202565b80610b6181610577565b0390f35b610208565b9190608083820312610be057610b82815f85016104cf565b92610b9082602083016104ac565b92604082013567ffffffffffffffff8111610bdb5783610bb19184016104e6565b929093606082013567ffffffffffffffff8111610bd657610bd292016104e6565b9091565b610210565b610210565b61020c565b34610c1c57610c18610c07610bfb366004610b6a565b94939093929192612727565b610c0f610202565b918291826107f7565b0390f35b610208565b34610c5157610c31366004610590565b610c4d610c3c61277d565b610c44610202565b9182918261072b565b0390f35b610208565b610c6260065f906106eb565b90565b34610c9557610c75366004610590565b610c91610c80610c56565b610c88610202565b9182918261072b565b0390f35b610208565b610cb1610ca83660046108a3565b93929092612a1c565b610cb9610202565b80610cc381610577565b0390f35b62ffffff1690565b610cdf906008610ce493026103e5565b610cc7565b90565b90610cf29154610ccf565b90565b610d0160085f90610ce7565b90565b34610d3457610d14366004610590565b610d30610d1f610cf5565b610d27610202565b91829182610a4b565b0390f35b610208565b610d4281610a36565b03610d4957565b5f80fd5b90503590610d5a82610d39565b565b9091606082840312610d9157610d8e610d77845f8501610d4d565b93610d8581602086016104ac565b93604001610855565b90565b61020c565b34610dc557610daf610da9366004610d5c565b91612b52565b610db7610202565b80610dc181610577565b0390f35b610208565b610dd660055f906106eb565b90565b34610e0957610de9366004610590565b610e05610df4610dca565b610dfc610202565b9182918261072b565b0390f35b610208565b34610e3e57610e1e366004610590565b610e3a610e29612b5f565b610e31610202565b9182918261072b565b0390f35b610208565b610e4f60095f906106eb565b90565b34610e8257610e62366004610590565b610e7e610e6d610e43565b610e75610202565b9182918261072b565b0390f35b610208565b5f80fd5b5f80fd5b67ffffffffffffffff8111610ea75760208091020190565b610226565b610eb58161042e565b03610ebc57565b5f80fd5b90503590610ecd82610eac565b565b91909161010081840312610f8f57610ee8610100610263565b92610ef5815f84016104ac565b5f850152610f0681602084016104ac565b6020850152610f1881604084016104ac565b6040850152610f2a81606084016104ac565b6060850152610f3c81608084016104ac565b6080850152610f4e8160a084016104ac565b60a0850152610f608160c084016104ac565b60c085015260e082013567ffffffffffffffff8111610f8a57610f8392016102dc565b60e0830152565b610e8b565b610e87565b91909160e08184031261105857610fab60e0610263565b92610fb8815f8401610ec0565b5f850152610fc98160208401610ec0565b6020850152610fdb81604084016104cf565b6040850152610fed8160608401610855565b6060850152610fff8160808401610855565b608085015260a082013567ffffffffffffffff81116110535781611024918401610ecf565b60a085015260c082013567ffffffffffffffff811161104e576110479201610ecf565b60c0830152565b610e8b565b610e8b565b610e87565b92919061107161106c82610e8f565b610263565b93818552602080860192028101918383116110c85781905b838210611097575050505050565b813567ffffffffffffffff81116110c3576020916110b88784938701610f94565b815201910190611089565b610214565b6104e2565b9080601f830112156110eb578160206110e89335910161105d565b90565b610214565b9190916101808184031261123057611109610180610263565b92611116815f8401610ec0565b5f8501526111278160208401610d4d565b602085015261113981604084016104ac565b604085015261114b81606084016104ac565b606085015261115d81608084016104ac565b608085015261116f8160a084016104ac565b60a08501526111818160c08401610855565b60c08501526111938160e08401610855565b60e08501526111a68161010084016104cf565b61010085015261012082013567ffffffffffffffff811161122b57816111cd9184016102dc565b61012085015261014082013567ffffffffffffffff811161122657816111f49184016102dc565b61014085015261016082013567ffffffffffffffff81116112215761121992016102dc565b610160830152565b610e8b565b610e8b565b610e8b565b610e87565b91909160408184031261129f5761124c6040610263565b925f82013567ffffffffffffffff811161129a578161126c9184016110cd565b5f850152602082013567ffffffffffffffff81116112955761128e92016110f0565b6020830152565b610e8b565b610e8b565b610e87565b906020828203126112d4575f82013567ffffffffffffffff81116112cf576112cc9201611235565b90565b610210565b61020c565b60209181520190565b61130161130a60209361130f936112f881610334565b938480936112d9565b9586910161033d565b61021c565b0190565b916113369261132960408201935f8301906107ea565b60208184039101526112e2565b90565b3461136a5761135161134c3660046112a4565b6136aa565b9061136661135d610202565b92839283611313565b0390f35b610208565b9190604083820312611397578061138b611394925f86016104ac565b93602001610855565b90565b61020c565b346113cb576113b56113af36600461136f565b90613fa4565b6113bd610202565b806113c781610577565b0390f35b610208565b906020828203126113e9576113e6915f016104cf565b90565b61020c565b6114026113fd61140792610389565b6105ea565b610389565b90565b90611414906113ee565b5f5260205260405f2090565b60f81b90565b61142f90611420565b90565b61144290600861144793026103e5565b611426565b90565b906114559154611432565b90565b61146e90611469600b915f9261140a565b61144a565b90565b346114a15761149d61148c6114873660046113d0565b611458565b611494610202565b91829182610952565b0390f35b610208565b6114b360086013906109c1565b90565b346114e6576114c6366004610590565b6114e26114d16114a6565b6114d9610202565b918291826109ec565b0390f35b610208565b6114f55f806105c2565b90565b3461152857611508366004610590565b6115246115136114eb565b61151b610202565b9182918261062e565b0390f35b610208565b61153a60086003906109c1565b90565b3461156d5761154d366004610590565b61156961155861152d565b611560610202565b918291826109ec565b0390f35b610208565b9060208282031261158b57611588915f01610855565b90565b61020c565b346115be576115a86115a3366004611572565b61401c565b6115b0610202565b806115ba81610577565b0390f35b610208565b90565b6115cf816115c3565b036115d657565b5f80fd5b905035906115e7826115c6565b565b9060808282031261164457611600815f84016115da565b9261160e82602085016104ac565b9261161c83604083016104cf565b92606082013567ffffffffffffffff811161163f5761163b92016104e6565b9091565b610210565b61020c565b6116606116573660046115e9565b93929092614095565b611668610202565b8061167281610577565b0390f35b5f80fd5b5f1c90565b61168b6116909161167a565b61059f565b90565b61169d905461167f565b90565b336116c46116be6116b96116b46001611693565b610615565b610712565b91610712565b03156116ed576116d2610202565b637d92a0f560e11b8152806116e960048201610577565b0390fd5b828491929091926140a4565b6116a0565b61170f9061170a61414a565b611a5f565b6117176141d2565b565b611724903690611235565b90565b60a01c90565b61173961173e91611727565b611426565b90565b61174b905461172d565b90565b5f80fd5b5f80fd5b5f80fd5b90359060016101800381360303821215611772570190565b61174e565b3561178181610841565b90565b3561178e81610d39565b90565b3561179b81610498565b90565b356117a8816104bb565b90565b634e487b7160e01b5f52601160045260245ffd5b6117ce6117d491939293610389565b92610389565b82018092116117df57565b6117ab565b5f80fd5b60e01b90565b905051906117fb826104bb565b565b9060208282031261181657611813915f016117ee565b90565b61020c565b634e487b7160e01b5f52602260045260245ffd5b906001600283049216801561184f575b602083101461184a57565b61181b565b91607f169161183f565b5f5260205f2090565b905f929180549061187c6118758361182f565b80946112d9565b916001811690815f146118d35750600114611897575b505050565b6118a49192939450611859565b915f925b8184106118bb57505001905f8080611892565b600181602092959395548486015201910192906118a8565b92949550505060ff19168252151560200201905f8080611892565b929061192491611917611932969461190d60808801945f8901906107ea565b60208701906109df565b8482036040860152611862565b9160608184039101526112e2565b90565b61193d610202565b3d5f823e3d90fd5b60209181520190565b5f7f53656e642065746820496e73756666696369656e740000000000000000000000910152565b6119826015602092611945565b61198b8161194e565b0190565b6119a49060208101905f818303910152611975565b90565b156119ae57565b6119b6610202565b62461bcd60e51b8152806119cc6004820161198f565b0390fd5b5f9103126119da57565b61020c565b9692611a4e95611a2d611a5c9a989396611a23611a3794611a19611a419860208f611a1261010082019e5f8301906109df565b01906109df565b60408d019061071e565b60608b019061071e565b60808901906107ea565b60a08701906109df565b84820360c0860152611862565b9160e08184039101526112e2565b90565b611ad490611a74611a6f82611719565b6136aa565b929092611a816006611741565b90611a9b60c0611a9586602081019061175a565b01611777565b611ab36020611aad878281019061175a565b01611784565b90611acd6040611ac788602081019061175a565b01611791565b92936142b9565b91611ae6611ae15f611693565b610615565b60206385fdd54291611b11611b0a610100611b04888681019061175a565b0161179e565b85906117bf565b90611b49611b2d6060611b27898781019061175a565b01611791565b94611b5460078a90611b3d610202565b988997889687966117e8565b8652600486016118ee565b03915afa8015611d0157611bb1915f91611cd3575b50611baa611ba4611b9f611b983494611b92610100611b8c8b602081019061175a565b0161179e565b906117bf565b86906117bf565b610389565b91610389565b10156119a7565b611bc2611bbd5f611693565b610615565b8263209afe5691349092611be56080611bdf85602081019061175a565b01611791565b94611bff60a0611bf986602081019061175a565b01611791565b97611c566060611c50611c45611c2460e0611c1e8b602081019061175a565b01611777565b95611c40610100611c3a339c602081019061175a565b0161179e565b6117bf565b9a602081019061175a565b01611791565b9060079091873b15611cce575f99611c8097611c8b95611c74610202565b9d8e9c8d9b8c9a6117e8565b8a5260048a016119df565b03925af18015611cc957611c9d575b50565b611cbc905f3d8111611cc2575b611cb4818361023a565b8101906119d0565b5f611c9a565b503d611caa565b611935565b6117e4565b611cf4915060203d8111611cfa575b611cec818361023a565b8101906117fd565b5f611b69565b503d611ce2565b611935565b611d0f906116fe565b565b5f90565b5f80fd5b5f80fd5b5f80fd5b9035600160200382360303811215611d6257016020813591019167ffffffffffffffff8211611d5d576020820236038313611d5857565b611d19565b611d15565b611d1d565b60209181520190565b90565b50611d82906020810190610ec0565b90565b611d8e9061042e565b9052565b50611da19060208101906104cf565b90565b611dad90610389565b9052565b50611dc0906020810190610855565b90565b611dcc90610712565b9052565b903560016101000382360303811215611de7570190565b611d1d565b50611dfb9060208101906104ac565b90565b611e079061048b565b9052565b9035600160200382360303811215611e4c57016020813591019167ffffffffffffffff8211611e47576001820236038313611e4257565b611d19565b611d15565b611d1d565b60209181520190565b9190611e7481611e6d81611e7995611e51565b809561029b565b61021c565b0190565b611f5691611f48610100820192611ea2611e995f830183611dec565b5f850190611dfe565b611ebc611eb26020830183611dec565b6020850190611dfe565b611ed6611ecc6040830183611dec565b6040850190611dfe565b611ef0611ee66060830183611dec565b6060850190611dfe565b611f0a611f006080830183611dec565b6080850190611dfe565b611f24611f1a60a0830183611dec565b60a0850190611dfe565b611f3e611f3460c0830183611dec565b60c0850190611dfe565b60e0810190611e0b565b9160e0818503910152611e5a565b90565b61201a9161200c61200160e08301611f7f611f765f870187611d73565b5f860190611d85565b611f99611f8f6020870187611d73565b6020860190611d85565b611fb3611fa96040870187611d92565b6040860190611da4565b611fcd611fc36060870187611db1565b6060860190611dc3565b611fe7611fdd6080870187611db1565b6080860190611dc3565b611ff460a0860186611dd0565b84820360a0860152611e7d565b9260c0810190611dd0565b9060c0818403910152611e7d565b90565b9061202791611f59565b90565b9035600160e00382360303811215612040570190565b611d1d565b60200190565b918161205691611d67565b908161206760208302840194611d70565b92835f925b84841061207c5750505050505090565b90919293949560206120a76120a1838560019503885261209c8b8861202a565b61201d565b98612045565b94019401929493919061206c565b9035600161018003823603038112156120cc570190565b611d1d565b506120e0906020810190610d4d565b90565b6120ec90610a36565b9052565b61224191612232612226612209610180840161211a6121115f880188611d73565b5f870190611d85565b61213461212a60208801886120d1565b60208701906120e3565b61214e6121446040880188611dec565b6040870190611dfe565b61216861215e6060880188611dec565b6060870190611dfe565b6121826121786080880188611dec565b6080870190611dfe565b61219c61219260a0880188611dec565b60a0870190611dfe565b6121b66121ac60c0880188611db1565b60c0870190611dc3565b6121d06121c660e0880188611db1565b60e0870190611dc3565b6121ec6121e1610100880188611d92565b610100870190611da4565b6121fa610120870187611e0b565b90868303610120880152611e5a565b612217610140860186611e0b565b90858303610140870152611e5a565b92610160810190611e0b565b91610160818503910152611e5a565b90565b6122849161227661226b6040830161225e5f860186611d21565b908583035f87015261204b565b9260208101906120b5565b9060208184039101526120f0565b90565b61229c9160208201915f818403910152612244565b90565b9035906001602003813603038212156122e1570180359067ffffffffffffffff82116122dc576020019160018202360383136122d757565b611756565b611752565b61174e565b919060408382031261232057612319906123006040610263565b9361230d825f83016104cf565b5f860152602001610855565b6020830152565b610e87565b9060408282031261233e5761233b915f016122e6565b90565b61020c565b61234d9051610389565b90565b612358611d11565b5060206123756123838361236a610202565b928391858301612287565b84820181038252038261023a565b916124126123b75f6123b16123a961239e868881019061175a565b61016081019061229f565b810190612325565b01612343565b936123c26006611741565b906123db60c06123d5868881019061175a565b01611777565b6123f2866123ec878281019061175a565b01611784565b9061240b6040612405888a81019061175a565b01611791565b92936142b9565b926124816124276124225f611693565b610615565b9161248c61246a606061246461245a6385fdd5429561245561010061244f8c8e81019061175a565b0161179e565b6117bf565b978981019061175a565b01611791565b966007612475610202565b988997889687966117e8565b8652600486016118ee565b03915afa9081156124d0575f916124a2575b5090565b6124c3915060203d81116124c9575b6124bb818361023a565b8101906117fd565b5f61249e565b503d6124b1565b611935565b60209181520190565b91816124e9916124d5565b90816124fa60208302840194611d70565b92835f925b84841061250f5750505050505090565b909192939495602061253a612534838560019503885261252f8b8861202a565b61201d565b98612045565b9401940192949391906124ff565b909161255f9260208301925f8185039101526124de565b90565b61256e6125739161167a565b610cc7565b90565b6125809054612562565b90565b60181c90565b61259561259a91612583565b61099c565b90565b6125a79054612589565b90565b91939061260b916125e46020956125bf611d11565b506125d66125cb610202565b938492898401612548565b87820181038252038261023a565b6125ee6006611741565b916125f96008612576565b90612604600861259d565b92936142b9565b9261264261262061261b5f611693565b610615565b9161264d6385fdd5429194966007612636610202565b988997889687966117e8565b8652600486016118ee565b03915afa908115612691575f91612663575b5090565b612684915060203d811161268a575b61267c818361023a565b8101906117fd565b5f61265f565b503d612672565b611935565b5f90565b6126a2612696565b90565b5f90565b6126b16126a5565b90565b6126bc6126a5565b90565b6126c7614308565b6126cf6126fc565b565b90565b6126e86126e36126ed926126d1565b6105ea565b6105df565b90565b6126f9906126d4565b90565b61270d6127085f6126f0565b61435d565b565b6127176126bf565b565b6127249136916102a6565b90565b9261274961274f9295612755979561273d611d11565b50969492909592612719565b93612719565b92614405565b90565b5f90565b61276861276d9161167a565b6106c8565b90565b61277a905461275c565b90565b612785612758565b506127906002612770565b90565b6127a76127a26127ac92610389565b6105ea565b61048b565b90565b5f7f4d45500000000000000000000000000000000000000000000000000000000000910152565b6127e36003602092611945565b6127ec816127af565b0190565b6128059060208101905f8183039101526127d6565b90565b1561280f57565b612817610202565b62461bcd60e51b81528061282d600482016127f0565b0390fd5b906128759493929161287061284546612793565b61286a61286461285f61285a3394600a6106b2565b612770565b610712565b91610712565b14612808565b6128d0565b565b1561287e57565b5f80fd5b60581c90565b61289461289991612882565b61099c565b90565b6128a69054612888565b90565b60981c90565b6128bb6128c0916128a9565b61099c565b90565b6128cd90546128af565b90565b929361295f9161293b9161292f6128e76006611741565b836128f26008612576565b6128fc600861259d565b918b9361292a879561291b61290f610202565b97889260208401612548565b6020820181038252038661023a565b6142b9565b968692869192936125aa565b61295861295261294d349386906117bf565b610389565b91610389565b1015612877565b61297061296b5f611693565b610615565b9063209afe569034929192612985600861289c565b9361299060086128c3565b9661299b6009612770565b903394979060079091873b15612a17575f996129c9976129d4956129bd610202565b9d8e9c8d9b8c9a6117e8565b8a5260048a016119df565b03925af18015612a12576129e6575b50565b612a05905f3d8111612a0b575b6129fd818361023a565b8101906119d0565b5f6129e3565b503d6129f3565b611935565b6117e4565b90612a2994939291612831565b565b90612a3e9291612a39614308565b612b2f565b565b5f1b90565b90612a5362ffffff91612a40565b9181191691161790565b612a71612a6c612a7692610a36565b6105ea565b610a36565b90565b90565b90612a91612a8c612a9892612a5d565b612a79565b8254612a45565b9055565b60181b90565b90612ab86affffffffffffffff00000091612a9c565b9181191691161790565b90565b90612ada612ad5612ae192610696565b612ac2565b8254612aa2565b9055565b90612af660018060a01b0391612a40565b9181191691161790565b612b0990610609565b90565b90565b90612b24612b1f612b2b92612b00565b612b0c565b8254612ae5565b9055565b612b509291612b42612b49926008612a7c565b6008612ac5565b6009612b0f565b565b90612b5d9291612a2b565b565b612b67612758565b90565b606090565b612b79905161042e565b90565b612b90612b8b612b95926126d1565b6105ea565b61042e565b90565b90565b612baf612baa612bb492612b98565b6105ea565b61042e565b90565b90565b612bce612bc9612bd392612bb7565b6105ea565b61042e565b90565b90565b612bed612be8612bf292612bd6565b6105ea565b61042e565b90565b90565b612c0c612c07612c1192612bf5565b6105ea565b61042e565b90565b90505190612c2182610841565b565b9190604083820312612c5d57612c5690612c3d6040610263565b93612c4a825f83016117ee565b5f860152602001612c14565b6020830152565b610e87565b90604082820312612c7b57612c78915f01612c23565b90565b61020c565b612c8b610120610263565b90565b5f90565b5f90565b5f90565b5f90565b5f90565b612caa612c80565b90602080808080808080808a612cbe612c8e565b815201612cc9612c92565b815201612cd4612c96565b815201612cdf612c9a565b815201612cea612c9a565b815201612cf5612c9a565b815201612d00612c9a565b815201612d0b612c9e565b815201612d16612c9e565b81525050565b612d24612ca2565b90565b634e487b7160e01b5f52603260045260245ffd5b5190565b90612d4982612d3b565b811015612d5a576020809102010190565b612d27565b612d73612d6e612d78926126d1565b6105ea565b610389565b90565b612d87612d8c9161167a565b6103e9565b90565b612d999054612d7b565b90565b90505190612da982610eac565b565b90505190612db882610d39565b565b612dc3816105df565b03612dca57565b5f80fd5b90505190612ddb82612dba565b565b919061012083820312612e9957612e9190612df9610120610263565b93612e06825f8301612d9c565b5f860152612e178260208301612dab565b6020860152612e298260408301612dce565b6040860152612e3b8260608301612c14565b6060860152612e4d8260808301612c14565b6080860152612e5f8260a08301612c14565b60a0860152612e718260c08301612c14565b60c0860152612e838260e083016117ee565b60e0860152610100016117ee565b610100830152565b610e87565b9061012082820312612eb857612eb5915f01612ddd565b90565b61020c565b63ffffffff1690565b63ffffffff60e01b1690565b612ee6612ee1612eeb92612ebd565b6117e8565b612ec6565b90565b612ef7906105df565b9052565b9061010080612f9c93612f145f8201515f860190611d85565b612f26602082015160208601906120e3565b612f3860408201516040860190612eee565b612f4a60608201516060860190611dc3565b612f5c60808201516080860190611dc3565b612f6e60a082015160a0860190611dc3565b612f8060c082015160c0860190611dc3565b612f9260e082015160e0860190611da4565b0151910190611da4565b565b9190612fb2905f6101208501940190612efb565b565b612fbe60e0610263565b90565b606090565b612fce612fb4565b90602080808080808088612fe0612c8e565b815201612feb612c9e565b815201612ff6612c9e565b815201613001612fc1565b81520161300c612c9a565b815201613017612c9a565b815201613022612c9e565b81525050565b613030612fc6565b90565b67ffffffffffffffff811161304b5760208091020190565b610226565b9092919261306561306082613033565b610263565b93818552602080860192028301928184116130a257915b8383106130895750505050565b602080916130978486612c14565b81520192019161307c565b6104e2565b9080601f830112156130c5578160206130c293519101613050565b90565b610214565b91909160e081840312613177576130e160e0610263565b926130ee815f8401612d9c565b5f8501526130ff81602084016117ee565b602085015261311181604084016117ee565b604085015260608201519167ffffffffffffffff83116131725761313a8261316b9483016130a7565b606086015261314c8260808301612c14565b608086015261315e8260a08301612c14565b60a086015260c0016117ee565b60c0830152565b610e8b565b610e87565b906020828203126131ac575f82015167ffffffffffffffff81116131a7576131a492016130ca565b90565b610210565b61020c565b5190565b60209181520190565b60200190565b906131d181602093611dc3565b0190565b60200190565b906131f86131f26131eb846131b1565b80936131b5565b926131be565b905f5b8181106132085750505090565b90919261322161321b60019286516131c4565b946131d5565b91019190916131fb565b906132af9060c08061328060e0840161324a5f8801515f870190611d85565b61325c60208801516020870190611da4565b61326e60408801516040870190611da4565b606087015185820360608701526131db565b9461329360808201516080860190611dc3565b6132a560a082015160a0860190611dc3565b0151910190611da4565b90565b6132c79160208201915f81840391015261322b565b90565b6132d46040610263565b90565b6132df6132ca565b90602080836132ec612c9e565b8152016132f7612c9a565b81525050565b6133056132d7565b90565b6133129051610a36565b90565b61331f905161048b565b90565b61332c9051610712565b90565b61333a610180610263565b90565b906133479061042e565b9052565b9061335590610a36565b9052565b906133639061048b565b9052565b9061337190610712565b9052565b9061337f90610389565b9052565b52565b6133906040610263565b90565b52565b52565b60200190565b6133be6133c76020936133cc936133b581610334565b93848093611e51565b9586910161033d565b61021c565b0190565b6134689160e06101008201926133ec5f8201515f850190611dfe565b6133fe60208201516020850190611dfe565b61341060408201516040850190611dfe565b61342260608201516060850190611dfe565b61343460808201516080850190611dfe565b61344660a082015160a0850190611dfe565b61345860c082015160c0850190611dfe565b01519060e081840391015261339f565b90565b6134f39160c06134e260e083016134885f8601515f860190611d85565b61349a60208601516020860190611d85565b6134ac60408601516040860190611da4565b6134be60608601516060860190611dc3565b6134d060808601516080860190611dc3565b60a085015184820360a08601526133d0565b9201519060c08184039101526133d0565b90565b906135009161346b565b90565b60200190565b9061351d61351683612d3b565b8092611d67565b908161352e60208302840194613399565b925f915b83831061354157505050505090565b9091929394602061356361355d838560019503875289516134f6565b97613503565b9301930191939290613532565b61365e9161016061364c61363861018084016135925f8701515f870190611d85565b6135a4602087015160208701906120e3565b6135b660408701516040870190611dfe565b6135c860608701516060870190611dfe565b6135da60808701516080870190611dfe565b6135ec60a087015160a0870190611dfe565b6135fe60c087015160c0870190611dc3565b61361060e087015160e0870190611dc3565b613624610100870151610100870190611da4565b61012086015185820361012087015261339f565b61014085015184820361014086015261339f565b9201519061016081840391015261339f565b90565b61368f91602061367e604083015f8501518482035f860152613509565b920151906020818403910152613570565b90565b6136a79160208201915f818403910152613661565b90565b906136b3611d11565b506136bc612b6a565b506136c5611d11565b6136cd612b6a565b506136d6612b6a565b6136e55f602086015101612b6f565b6136f76136f15f612b7c565b9161042e565b145f14613a915750506137086132fd565b5061373b613736600c60e060a061372b5f8801516137255f612d5f565b90612d3f565b5101510151906103d9565b612d8f565b61374d6137475f612b7c565b9161042e565b145f1461397a5761378c60e060a06137715f86015161376b5f612d5f565b90612d3f565b5101510151602061378182610334565b818301019101612c62565b916137b85f60e060a06137ab838601516137a585612d5f565b90612d3f565b5101510151945b01612343565b92610140602083015101515b82602001515f016137d490612b6f565b9183602001516020016137e690613308565b9184602001516040016137f890613315565b856020015160600161380990613315565b866020015160800161381a90613315565b876020015160a00161382b90613315565b886020015160c00161383c90613322565b91896020015160e00161384e90613322565b938a602001516101000161386190612343565b958b60200151610120015197989961387761332f565b9b5f8d01906138859161333d565b60208c01906138939161334b565b60408b01906138a191613359565b60608a01906138af91613359565b60808901906138bd91613359565b60a08801906138cb91613359565b60c08701906138d991613367565b60e08601906138e791613367565b6101008501906138f691613375565b61012084019061390591613383565b61014083019061391491613383565b61016082019061392391613383565b905f015190613930613386565b915f83019061393e91613393565b602082019061394c91613396565b613954610202565b8091602082019061396491613692565b602082018103825203613977908261023a565b90565b6139ac6139a7600c60e060a061399c5f8801516139965f612d5f565b90612d3f565b5101510151906103d9565b612d8f565b6139bf6139b96001612b9b565b9161042e565b1480613a51575b5f14613a2e57613a0460e060c06139e95f8601516139e35f612d5f565b90612d3f565b510151015160206139f982610334565b818301019101612c62565b916137b85f60e060c0613a2383860151613a1d85612d5f565b90612d3f565b5101510151946137b2565b613a36610202565b6382077b6960e01b815280613a4d60048201610577565b0390fd5b50613a7960e060c0613a6f5f860151613a695f612d5f565b90612d3f565b5101510151610334565b613a8b613a855f612d5f565b91610389565b116139c6565b613aa25f6020869496015101612b6f565b613ab5613aaf6001612b9b565b9161042e565b145f14613c8357613ac4613028565b50613af7613af2600c60e060a0613ae75f880151613ae15f612d5f565b90612d3f565b5101510151906103d9565b612d8f565b613b09613b035f612b7c565b9161042e565b145f14613b8c576004613b86613b4d60e060a0613b325f880151613b2c5f612d5f565b90612d3f565b51015101516020613b4282610334565b81830101910161317c565b5b613b77613b5e6374a736b9612ed2565b91613b67610202565b94859360208501908152016132b2565b6020820181038252038261023a565b5b6137c4565b613bbe613bb9600c60e060a0613bae5f880151613ba85f612d5f565b90612d3f565b5101510151906103d9565b612d8f565b613bd1613bcb6001612b9b565b9161042e565b1480613c43575b5f14613c20576004613b86613c1b60e060c0613c005f880151613bfa5f612d5f565b90612d3f565b51015101516020613c1082610334565b81830101910161317c565b613b4e565b613c28610202565b6382077b6960e01b815280613c3f60048201610577565b0390fd5b50613c6b60e060c0613c615f860151613c5b5f612d5f565b90612d3f565b5101510151610334565b613c7d613c775f612d5f565b91610389565b11613bd8565b613c925f602084015101612b6f565b613ca5613c9f6002612bba565b9161042e565b145f14613e7357613cb4612d1c565b50613ce7613ce2600c60e060a0613cd75f880151613cd15f612d5f565b90612d3f565b5101510151906103d9565b612d8f565b613cf9613cf35f612b7c565b9161042e565b145f14613d7c576004613d76613d3d60e060a0613d225f880151613d1c5f612d5f565b90612d3f565b51015101516020613d3282610334565b818301019101612e9e565b5b613d67613d4e63bc33e620612ed2565b91613d57610202565b9485936020850190815201612f9e565b6020820181038252038261023a565b5b613b87565b613dae613da9600c60e060a0613d9e5f880151613d985f612d5f565b90612d3f565b5101510151906103d9565b612d8f565b613dc1613dbb6001612b9b565b9161042e565b1480613e33575b5f14613e10576004613d76613e0b60e060c0613df05f880151613dea5f612d5f565b90612d3f565b51015101516020613e0082610334565b818301019101612e9e565b613d3e565b613e18610202565b6382077b6960e01b815280613e2f60048201610577565b0390fd5b50613e5b60e060c0613e515f860151613e4b5f612d5f565b90612d3f565b5101510151610334565b613e6d613e675f612d5f565b91610389565b11613dc8565b509150613e855f602084015101612b6f565b613e98613e9260fe612bd9565b9161042e565b145f14613ee757613ecc5f613ec6610160602086015101516020613ebb82610334565b818301019101612c62565b01612343565b9161016060208201510151610140602083015101515b613d77565b613ef65f602084015101612b6f565b613f09613f0360ff612bf8565b9161042e565b145f14613f5757613f3d5f613f37610160602086015101516020613f2c82610334565b818301019101612c62565b01612343565b916101606020820151015161014060208301510151613ee2565b613f5f610202565b634157854360e11b815280613f7660048201610577565b0390fd5b90613f8c91613f87614308565b613f8e565b565b613f9d613fa29291600a6106b2565b612b0f565b565b90613fae91613f7a565b565b613fc190613fbc614308565b613fc3565b565b80613fde613fd8613fd35f6126f0565b610712565b91610712565b14613fee57613fec9061435d565b565b614018613ffa5f6126f0565b614002610202565b918291631e4fbdf760e01b83526004830161072b565b0390fd5b61402590613fb0565b565b939291903361404f61404961404461403f6001611693565b610615565b610712565b91610712565b0361405f5761405d94614082565b565b614067610202565b637d92a0f560e11b81528061407e60048201610577565b0390fd5b916140939492939190919293614950565b565b906140a294939291614027565b565b6140ac610202565b637fca089960e11b8152806140c360048201610577565b0390fd5b6140d36140d89161167a565b610aca565b90565b6140e590546140c7565b90565b6140fc6140f761410192612bb7565b6105ea565b610389565b90565b61410e60026140e8565b90565b9061411d5f1991612a40565b9181191691161790565b90565b9061413f61413a614146926113ee565b614127565b8254614111565b9055565b61415460036140db565b61416d614167614162614104565b610389565b91610389565b146141865761418461417d614104565b600361412a565b565b61418e610202565b633ee5aeb560e01b8152806141a560048201610577565b0390fd5b6141bd6141b86141c292612b98565b6105ea565b610389565b90565b6141cf60016141a9565b90565b6141e46141dd6141c5565b600361412a565b565b6141ef906105ed565b90565b61420661420161420b926105df565b6105ea565b610389565b90565b90565b61421d6142229161093c565b61420e565b9052565b60e81b90565b61423590614226565b90565b61424461424991610a36565b61422c565b9052565b60c01b90565b61425c9061424d565b90565b61426b6142709161048b565b614253565b9052565b6008936142a760206142b698979561429f6001866142976003986142af9a614211565b01809261038f565b018092614238565b01809261425f565b0190610360565b90565b926142f69192946142de6142d9614305966142d2612b6a565b50956141e6565b6141f2565b959190916142ea610202565b96879560208701614274565b6020820181038252038261023a565b90565b61431061277d565b61432961432361431e614c77565b610712565b91610712565b0361433057565b61435961433b614c77565b614343610202565b91829163118cdaa760e01b83526004830161072b565b0390fd5b6143676002612770565b614372826002612b0f565b906143a66143a07f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093612b00565b91612b00565b916143af610202565b806143b981610577565b0390a3565b92906143f4916143e761440296946143dd60808801945f8901906107ea565b60208701906109df565b84820360408601526112e2565b9160608184039101526112e2565b90565b9261444760209394614415611d11565b5061445261442a6144255f611693565b610615565b936385fdd54292959761443b610202565b988997889687966117e8565b8652600486016143be565b03915afa908115614496575f91614468575b5090565b614489915060203d811161448f575b614481818361023a565b8101906117fd565b5f614464565b503d614477565b611935565b6144af6144aa6144b492610389565b6105ea565b6105df565b90565b5f7f496e76616c696420636f6e747261637400000000000000000000000000000000910152565b6144eb6010602092611945565b6144f4816144b7565b0190565b61450d9060208101905f8183039101526144de565b90565b1561451757565b61451f610202565b62461bcd60e51b815280614535600482016144f8565b0390fd5b9050519061454682610498565b565b9092919261455d61455882610278565b610263565b93818552602085019082840111614579576145779261033d565b565b610218565b9080601f8301121561459c5781602061459993519101614548565b90565b610214565b91909161010081840312614661576145ba610100610263565b926145c7815f8401614539565b5f8501526145d88160208401614539565b60208501526145ea8160408401614539565b60408501526145fc8160608401614539565b606085015261460e8160808401614539565b60808501526146208160a08401614539565b60a08501526146328160c08401614539565b60c085015260e082015167ffffffffffffffff811161465c57614655920161457e565b60e0830152565b610e8b565b610e87565b91909160e08184031261472a5761467d60e0610263565b9261468a815f8401612d9c565b5f85015261469b8160208401612d9c565b60208501526146ad81604084016117ee565b60408501526146bf8160608401612c14565b60608501526146d18160808401612c14565b608085015260a082015167ffffffffffffffff811161472557816146f69184016145a1565b60a085015260c082015167ffffffffffffffff81116147205761471992016145a1565b60c0830152565b610e8b565b610e8b565b610e87565b92919061474361473e82610e8f565b610263565b938185526020808601920281019183831161479a5781905b838210614769575050505050565b815167ffffffffffffffff81116147955760209161478a8784938701614666565b81520191019061475b565b610214565b6104e2565b9080601f830112156147bd578160206147ba9351910161472f565b90565b610214565b906020828203126147f2575f82015167ffffffffffffffff81116147ed576147ea920161479f565b90565b610210565b61020c565b614800906105ed565b90565b61480c906147f7565b90565b61481890610609565b90565b9061482f61482883612d3b565b80926124d5565b908161484060208302840194613399565b925f915b83831061485357505050505090565b9091929394602061487561486f838560019503875289516134f6565b97613503565b9301930191939290614844565b6148979160208201915f81840391015261481b565b90565b5f90565b6148a790610609565b90565b906148bc6148b783610278565b610263565b918252565b3d5f146148dc576148d13d6148aa565b903d5f602084013e5b565b6148e4612b6a565b906148da565b6148f55f8092610338565b0190565b614902906148ea565b90565b151590565b61491390614905565b9052565b9161494d939161493f9161493260608601925f87019061490a565b84820360208601526112e2565b9160408184039101526112e2565b90565b5061499e93929161498f61498961498461497f61497961497461499597600a6106b2565b612770565b9461449b565b610609565b610712565b91610712565b14614510565b908101906112a4565b6149c16101206020830151015160206149b682610334565b8183010191016147c2565b6149ed6149e86149e36149de600a6149d846612793565b906106b2565b612770565b614803565b61480f565b906377c2719c90823b15614c7257614a2492614a195f8094614a0d610202565b968795869485936117e8565b835260048301614882565b03925af18015614c6d57614c41575b50614a3c61489a565b50614a45612b6a565b505f80614a6b610160602085015101516020614a6082610334565b818301019101612c62565b614a7a82602086015101612b6f565b614a8d614a8760fe612bd9565b9161042e565b148214614b2157614abb82614ab4614aaf600a614aa946612793565b906106b2565b612770565b9201612343565b6101406020860151015190602082019151925af1906101406020614add6148c1565b935b93920151015191614b1c7fc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb39807944693614b13610202565b93849384614917565b0390a1565b614b3082602086015101612b6f565b614b43614b3d60ff612bf8565b9161042e565b148214614b9a57614b7182614b6a614b65600a614b5f46612793565b906106b2565b612770565b9201612343565b6101406020860151015190602082019151925af1906101406020614b936148c1565b935b614adf565b614ba982602086015101612b6f565b614bbb614bb584612b7c565b9161042e565b148214614c0557614bd982614bd260208401613322565b9201612343565b614be1610202565b9081614bec816148f9565b03925af1906101406020614bfe6148c1565b935b614b95565b614c1982614c123061489e565b9201612343565b6101406020860151015190602082019151925af1906101406020614c3b6148c1565b93614c00565b614c60905f3d8111614c66575b614c58818361023a565b8101906119d0565b5f614a33565b503d614c4e565b611935565b6117e4565b614c7f612758565b50339056fea264697066735822122001d6a9766451fa27cf692849cc81e5d8ed068e8377f94b42b59380c8fc80b64864736f6c63430008180033",
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
