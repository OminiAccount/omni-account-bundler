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
	PackedUserOperation BaseStructPackedUserOperation
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vizingPad\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_Hook\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LandingPadAccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplement\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"packHookMessage\",\"type\":\"bytes\"}],\"name\":\"ReceiveTouchHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"VizingSwapEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Hook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LandingPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LaunchPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LockWay\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"MirrorEntryPoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBridgeMode\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGasPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGaslimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destChainid\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"additionParams\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"estimateVizingGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vizingGasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"fetchOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"_packedUserOperation\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"fetchUserOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"_packedUserOperation\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"getUserOmniMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveExtraMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveStandardMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selectedRelayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"sendOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"_packedUserOperation\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"sendUserOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_way\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"_lockState\",\"type\":\"bytes1\"}],\"name\":\"setLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"}],\"name\":\"setMirrorEntryPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"testReceiveMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e06040523462000068576200001f6200001862000161565b9162000706565b620000296200006e565b615f8c62000a8b82396080518181816106c901526129d0015260a051818181610ca801526129f2015260c051818181610aa30152612a140152615f8c90f35b62000074565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90620000a29062000078565b810190811060018060401b03821117620000bb57604052565b62000082565b90620000d8620000d06200006e565b928362000096565b565b5f80fd5b60018060a01b031690565b620000f490620000de565b90565b6200010281620000e9565b036200010a57565b5f80fd5b905051906200011d82620000f7565b565b90916060828403126200015b57620001586200013e845f85016200010e565b936200014e81602086016200010e565b936040016200010e565b90565b620000da565b6200018462006a17803803806200017881620000c1565b9283398101906200011f565b909192565b5f1b90565b906200019c60ff9162000189565b9181191691161790565b90565b60ff60f81b1690565b60f81b90565b620001d1620001cb620001d792620001a6565b620001b2565b620001a9565b90565b60f81c90565b620001eb90620001da565b90565b9062000208620002026200021092620001b8565b620001e0565b82546200018e565b9055565b90565b90565b90565b62000236620002306200023c9262000214565b6200021a565b62000217565b90565b60018060401b0381116200025e576200025a60209162000078565b0190565b62000082565b906200027a62000274836200023f565b620000c1565b918252565b369037565b90620002af620002948362000264565b92602080620002a486936200023f565b92019103906200027f565b565b5190565b634e487b7160e01b5f52602260045260245ffd5b9060016002830492168015620002ec575b6020831014620002e657565b620002b5565b91607f1691620002da565b5f5260205f2090565b601f602091010490565b1b90565b919060086200032d910291620003265f19846200030a565b926200030a565b9181191691161790565b620003506200034a620003569262000217565b6200021a565b62000217565b90565b90565b91906200037762000371620003809362000337565b62000359565b9083546200030e565b9055565b5f90565b6200039e916200039762000384565b916200035c565b565b5b818110620003ad575050565b80620003bc5f60019362000388565b01620003a1565b9190601f8111620003d4575b505050565b620003e36200040e93620002f7565b906020620003f18462000300565b8301931062000417575b620004069062000300565b0190620003a0565b5f8080620003cf565b91506200040681929050620003fb565b1c90565b906200043d905f199060080262000427565b191690565b816200044e916200042b565b906002021790565b906200046281620002b1565b9060018060401b03821162000533576200048982620004828554620002c9565b85620003c3565b602090601f8311600114620004c257918091620004b0935f92620004b5575b505062000442565b90555b565b90915001515f80620004a8565b601f19831691620004d385620002f7565b925f5b8181106200051a57509160029391856001969410620004fd575b50505002019055620004b3565b6200050f910151601f8416906200042b565b90555f8080620004f0565b91936020600181928787015181550195019201620004d6565b62000082565b90620005459162000456565b565b906200055762ffffff9162000189565b9181191691161790565b90565b62ffffff1690565b620005856200057f6200058b9262000561565b6200021a565b62000564565b90565b90565b90620005ab620005a5620005b3926200056c565b6200058e565b825462000547565b9055565b60181b90565b90620005d46301000000600160581b0391620005b7565b9181191691161790565b90565b60018060401b031690565b62000605620005ff6200060b92620005de565b6200021a565b620005e1565b90565b90565b906200062b620006256200063392620005ec565b6200060e565b8254620005bd565b9055565b60081b90565b9062000652610100600160a81b039162000637565b9181191691161790565b620006756200066f6200067b92620000de565b6200021a565b620000de565b90565b62000689906200065c565b90565b62000697906200067e565b90565b90565b90620006b7620006b1620006bf926200068c565b6200069a565b82546200063d565b9055565b90620006d660018060a01b039162000189565b9181191691161790565b90620006fa620006f462000702926200068c565b6200069a565b8254620006c3565b9055565b6200077b92916200071c6200077392336200077d565b6200072a60016004620001ee565b6200074b620007436200073d5f6200021d565b62000284565b600662000539565b6200075a61c350600762000591565b6200076b633b9aca00600762000611565b60046200069d565b6005620006e0565b565b9062000789916200078b565b565b90620007979162000799565b565b90620007a59162000816565b565b620007c0620007ba620007c692620001a6565b6200021a565b62000217565b90565b620007d56001620007a7565b90565b90620007e65f199162000189565b9181191691161790565b906200080a62000804620008129262000337565b62000359565b8254620007d8565b9055565b90620008229162000890565b6200083862000830620007c9565b6003620007f0565b565b620008536200084d620008599262000214565b6200021a565b620000de565b90565b62000867906200083a565b90565b6200087590620000e9565b9052565b91906200088e905f602085019401906200086a565b565b906200089c9062000904565b80620008bd620008b6620008b05f6200085c565b620000e9565b91620000e9565b14620008d057620008ce9062000999565b565b62000900620008df5f6200085c565b620008e96200006e565b918291631e4fbdf760e01b83526004830162000879565b0390fd5b6200090f9062000911565b565b6200091d90806200091f565b565b6200092e620009349262000936565b62000a59565b565b620009419062000943565b565b6200094e9062000950565b565b6200095b9062000a72565b565b5f1c90565b60018060a01b031690565b6200097c62000982916200095d565b62000962565b90565b6200099190546200096d565b90565b5f0190565b620009a5600262000985565b620009b2826002620006e0565b90620009ea620009e37f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936200068c565b916200068c565b91620009f56200006e565b8062000a018162000994565b0390a3565b62000a11906200065c565b90565b62000a1f9062000a06565b90565b62000a2d9062000a06565b90565b90565b9062000a4d62000a4762000a559262000a22565b62000a30565b8254620006c3565b9055565b62000a6862000a709162000a14565b600162000a33565b565b62000a8162000a889162000a14565b5f62000a33565b56fe60806040526004361015610015575b36610faa57005b61001f5f356101dd565b806273b555146101d85780630e82845d146101d35780631490a358146101ce578063397a567d146101c95780633eba5864146101c457806345636279146101bf5780635ad3ad06146101ba5780635aeb4d77146101b55780635e45da23146101b05780636c7cfb52146101ab578063715018a6146101a657806376c81312146101a15780638da5cb5b1461019c5780638ea23ddf1461019757806393afae931461019257806399e581fa1461018d578063ad5c464814610188578063b0cfd4d214610183578063d298a7e91461017e578063d5f3e00814610179578063de4b1d6d14610174578063de8aeda01461016f578063e0b838e91461016a578063e7b4294c14610165578063ebe98c6f14610160578063f2fde38b1461015b578063f3148925146101565763ff16742b0361000e57610f75565b610ec5565b610e0c565b610db9565b610d84565b610d0c565b610cca565b610c71565b610b9c565b610b46565b610ac5565b610a6c565b610a27565b6109bd565b610988565b610944565b610908565b61085a565b610826565b6107a1565b61076c565b61070d565b610692565b61062e565b610520565b6104b5565b6103b8565b6102e9565b60e01c90565b60405190565b5f80fd5b5f80fd5b67ffffffffffffffff1690565b610207816101f1565b0361020e57565b5f80fd5b9050359061021f826101fe565b565b90565b61022d81610221565b0361023457565b5f80fd5b9050359061024582610224565b565b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561028d5781359167ffffffffffffffff831161028857602001926001830284011161028357565b61024f565b61024b565b610247565b916060838303126102df576102a9825f8501610212565b926102b78360208301610238565b92604082013567ffffffffffffffff81116102da576102d69201610253565b9091565b6101ed565b6101e9565b5f0190565b6102f4366004610292565b9291909161102d565b5f80fd5b5f91031261030b57565b6101e9565b1c90565b60018060a01b031690565b61032f9060086103349302610310565b610314565b90565b90610342915461031f565b90565b61035160015f90610337565b90565b60018060a01b031690565b90565b61037661037161037b92610354565b61035f565b610354565b90565b61038790610362565b90565b6103939061037e565b90565b61039f9061038a565b9052565b91906103b6905f60208501940190610396565b565b346103e8576103c8366004610301565b6103e46103d3610345565b6103db6101e3565b918291826103a3565b0390f35b6102fd565b9060208282031261040657610403915f01610212565b90565b6101e9565b61041f61041a610424926101f1565b61035f565b6101f1565b90565b906104319061040b565b5f5260205260405f2090565b60018060a01b031690565b61045890600861045d9302610310565b61043d565b90565b9061046b9154610448565b90565b6104849061047f6008915f92610427565b610460565b90565b61049090610354565b90565b61049c90610487565b9052565b91906104b3905f60208501940190610493565b565b346104e5576104e16104d06104cb3660046103ed565b61046e565b6104d86101e3565b918291826104a0565b0390f35b6102fd565b9060208282031261051b575f82013567ffffffffffffffff8111610516576105129201610253565b9091565b6101ed565b6101e9565b61053461052e3660046104ea565b90611c99565b61053c6101e3565b80610546816102e4565b0390f35b61055381610487565b0361055a57565b5f80fd5b9050359061056b8261054a565b565b909182601f830112156105a75781359167ffffffffffffffff83116105a257602001926020830284011161059d57565b61024f565b61024b565b610247565b90608082820312610607576105c3815f8401610212565b926105d1826020850161055e565b926105df8360408301610238565b92606082013567ffffffffffffffff8111610602576105fe920161056d565b9091565b6101ed565b6101e9565b61061590610221565b9052565b919061062c905f6020850194019061060c565b565b346106625761065e61064d6106443660046105ac565b9392909261253a565b6106556101e3565b91829182610619565b0390f35b6102fd565b60ff60f81b1690565b61067990610667565b9052565b9190610690905f60208501940190610670565b565b346106c2576106a2366004610301565b6106be6106ad61262a565b6106b56101e3565b9182918261067d565b0390f35b6102fd565b7f000000000000000000000000000000000000000000000000000000000000000090565b6106f4906101f1565b9052565b919061070b905f602085019401906106eb565b565b3461073d5761071d366004610301565b6107396107286106c7565b6107306101e3565b918291826106f8565b0390f35b6102fd565b62ffffff1690565b61075390610742565b9052565b919061076a905f6020850194019061074a565b565b3461079c5761077c366004610301565b610798610787612639565b61078f6101e3565b91829182610757565b0390f35b6102fd565b346107d1576107b1366004610301565b6107cd6107bc612644565b6107c46101e3565b91829182610757565b0390f35b6102fd565b6107df81610667565b036107e657565b5f80fd5b905035906107f7826107d6565b565b9190604083820312610821578061081561081e925f8601610238565b936020016107ea565b90565b6101e9565b346108555761083f6108393660046107f9565b906126d2565b6108476101e3565b80610851816102e4565b0390f35b6102fd565b346108885761086a366004610301565b61087261272b565b61087a6101e3565b80610884816102e4565b0390f35b6102fd565b9190608083820312610903576108a5815f8501610238565b926108b38260208301610212565b92604082013567ffffffffffffffff81116108fe57836108d4918401610253565b929093606082013567ffffffffffffffff81116108f9576108f59201610253565b9091565b6101ed565b6101ed565b6101e9565b3461093f5761093b61092a61091e36600461088d565b94939093929192612743565b6109326101e3565b91829182610619565b0390f35b6102fd565b3461097457610954366004610301565b61097061095f612778565b6109676101e3565b918291826104a0565b0390f35b6102fd565b61098560055f90610460565b90565b346109b857610998366004610301565b6109b46109a3610979565b6109ab6101e3565b918291826104a0565b0390f35b6102fd565b6109d46109cb3660046105ac565b93929092612ab4565b6109dc6101e3565b806109e6816102e4565b0390f35b62ffffff1690565b610a02906008610a079302610310565b6109ea565b90565b90610a1591546109f2565b90565b610a2460075f90610a0a565b90565b34610a5757610a37366004610301565b610a53610a42610a18565b610a4a6101e3565b91829182610757565b0390f35b6102fd565b610a696004600190610460565b90565b34610a9c57610a7c366004610301565b610a98610a87610a5c565b610a8f6101e3565b918291826104a0565b0390f35b6102fd565b7f000000000000000000000000000000000000000000000000000000000000000090565b34610af557610ad5366004610301565b610af1610ae0610aa1565b610ae86101e3565b918291826104a0565b0390f35b6102fd565b5f80fd5b90816040910312610b0c5790565b610afa565b90602082820312610b41575f82013567ffffffffffffffff8111610b3c57610b399201610afe565b90565b6101ed565b6101e9565b610b59610b54366004610b11565b614395565b610b616101e3565b80610b6b816102e4565b0390f35b9190604083820312610b975780610b8b610b94925f8601610212565b9360200161055e565b90565b6101e9565b34610bcb57610bb5610baf366004610b6f565b90614414565b610bbd6101e3565b80610bc7816102e4565b0390f35b6102fd565b90602082820312610be957610be6915f01610238565b90565b6101e9565b610c02610bfd610c0792610221565b61035f565b610221565b90565b90610c1490610bee565b5f5260205260405f2090565b60f81b90565b610c2f90610c20565b90565b610c42906008610c479302610310565b610c26565b90565b90610c559154610c32565b90565b610c6e90610c696009915f92610c0a565b610c4a565b90565b34610ca157610c9d610c8c610c87366004610bd0565b610c58565b610c946101e3565b9182918261067d565b0390f35b6102fd565b7f000000000000000000000000000000000000000000000000000000000000000090565b34610cfa57610cda366004610301565b610cf6610ce5610ca6565b610ced6101e3565b918291826106f8565b0390f35b6102fd565b610d095f80610337565b90565b34610d3c57610d1c366004610301565b610d38610d27610cff565b610d2f6101e3565b918291826103a3565b0390f35b6102fd565b67ffffffffffffffff1690565b610d5e906008610d639302610310565b610d41565b90565b90610d719154610d4e565b90565b610d816007600390610d66565b90565b34610db457610d94366004610301565b610db0610d9f610d74565b610da76101e3565b918291826106f8565b0390f35b6102fd565b34610de957610de5610dd4610dcf366004610b11565b6145fc565b610ddc6101e3565b91829182610619565b0390f35b6102fd565b90602082820312610e0757610e04915f0161055e565b90565b6101e9565b34610e3a57610e24610e1f366004610dee565b6149a9565b610e2c6101e3565b80610e36816102e4565b0390f35b6102fd565b90565b610e4b81610e3f565b03610e5257565b5f80fd5b90503590610e6382610e42565b565b90608082820312610ec057610e7c815f8401610e56565b92610e8a8260208501610212565b92610e988360408301610238565b92606082013567ffffffffffffffff8111610ebb57610eb79201610253565b9091565b6101ed565b6101e9565b610edc610ed3366004610e65565b93929092614a22565b610ee46101e3565b80610eee816102e4565b0390f35b5190565b60209181520190565b5f5b838110610f11575050905f910152565b806020918301518185015201610f01565b601f801991011690565b610f4b610f54602093610f5993610f4281610ef2565b93848093610ef6565b95869101610eff565b610f22565b0190565b610f729160208201915f818403910152610f2c565b90565b34610fa557610fa1610f90610f8b366004610b11565b614a31565b610f986101e3565b91829182610f5d565b0390f35b6102fd565b5f80fd5b5f1c90565b610fbf610fc491610fae565b610314565b90565b610fd19054610fb3565b90565b33610ff8610ff2610fed610fe86001610fc7565b61038a565b610487565b91610487565b0315611021576110066101e3565b637d92a0f560e11b81528061101d600482016102e4565b0390fd5b828491929091926156d8565b610fd4565b61104661104161104b92610221565b61035f565b6101f1565b90565b61105a61105f91610fae565b61043d565b90565b61106c905461104e565b90565b61107890610362565b90565b6110849061106f565b90565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b906110a990610f22565b810190811067ffffffffffffffff8211176110c357604052565b61108b565b906110db6110d46101e3565b928361109f565b565b5f80fd5b60ff1690565b6110f0816110e1565b036110f757565b5f80fd5b90503590611108826110e7565b565b5f80fd5b67ffffffffffffffff811161112c57611128602091610f22565b0190565b61108b565b90825f939282370152565b9092919261115161114c8261110e565b6110c8565b9381855260208501908284011161116d5761116b92611131565b565b61110a565b9080601f830112156111905781602061118d9335910161113c565b90565b610247565b91909161010081840312611255576111ae6101006110c8565b926111bb815f8401610212565b5f8501526111cc8160208401610212565b60208501526111de8160408401610212565b60408501526111f08160608401610212565b60608501526112028160808401610212565b60808501526112148160a08401610212565b60a08501526112268160c08401610212565b60c085015260e082013567ffffffffffffffff8111611250576112499201611172565b60e0830152565b6110dd565b611087565b91909160e08184031261131e5761127160e06110c8565b9261127e815f84016110fb565b5f85015261128f81602084016110fb565b60208501526112a18160408401610238565b60408501526112b3816060840161055e565b60608501526112c5816080840161055e565b608085015260a082013567ffffffffffffffff811161131957816112ea918401611195565b60a085015260c082013567ffffffffffffffff81116113145761130d9201611195565b60c0830152565b6110dd565b6110dd565b611087565b61132c81610742565b0361133357565b5f80fd5b9050359061134482611323565b565b919091610180818403126114865761135f6101806110c8565b9261136c815f84016110fb565b5f85015261137d8160208401611337565b602085015261138f8160408401610212565b60408501526113a18160608401610212565b60608501526113b38160808401610212565b60808501526113c58160a08401610212565b60a08501526113d78160c0840161055e565b60c08501526113e98160e0840161055e565b60e08501526113fc816101008401610238565b61010085015261012082013567ffffffffffffffff81116114815781611423918401611172565b61012085015261014082013567ffffffffffffffff811161147c578161144a918401611172565b61014085015261016082013567ffffffffffffffff81116114775761146f9201611172565b610160830152565b6110dd565b6110dd565b6110dd565b611087565b9190916040818403126114f5576114a260406110c8565b925f82013567ffffffffffffffff81116114f057816114c291840161125a565b5f850152602082013567ffffffffffffffff81116114eb576114e49201611346565b6020830152565b6110dd565b6110dd565b611087565b9060208282031261152a575f82013567ffffffffffffffff811161152557611522920161148b565b90565b6101ed565b6101e9565b6115389061037e565b90565b5f80fd5b60e01b90565b90505190611552826101fe565b565b9060208282031261156d5761156a915f01611545565b90565b6101e9565b61157a6101e3565b3d5f823e3d90fd5b909291926115976115928261110e565b6110c8565b938185526020850190828401116115b3576115b192610eff565b565b61110a565b9080601f830112156115d6578160206115d393519101611582565b90565b610247565b9060208282031261160b575f82015167ffffffffffffffff81116116065761160392016115b8565b90565b6101ed565b6101e9565b67ffffffffffffffff81116116285760208091020190565b61108b565b9050519061163a826110e7565b565b9050519061164982610224565b565b905051906116588261054a565b565b9190916101008184031261171a576116736101006110c8565b92611680815f8401611545565b5f8501526116918160208401611545565b60208501526116a38160408401611545565b60408501526116b58160608401611545565b60608501526116c78160808401611545565b60808501526116d98160a08401611545565b60a08501526116eb8160c08401611545565b60c085015260e082015167ffffffffffffffff81116117155761170e92016115b8565b60e0830152565b6110dd565b611087565b91909160e0818403126117e35761173660e06110c8565b92611743815f840161162d565b5f850152611754816020840161162d565b6020850152611766816040840161163c565b6040850152611778816060840161164b565b606085015261178a816080840161164b565b608085015260a082015167ffffffffffffffff81116117de57816117af91840161165a565b60a085015260c082015167ffffffffffffffff81116117d9576117d2920161165a565b60c0830152565b6110dd565b6110dd565b611087565b9291906117fc6117f782611610565b6110c8565b93818552602080860192028101918383116118535781905b838210611822575050505050565b815167ffffffffffffffff811161184e57602091611843878493870161171f565b815201910190611814565b610247565b61024f565b9080601f8301121561187657816020611873935191016117e8565b90565b610247565b906020828203126118ab575f82015167ffffffffffffffff81116118a6576118a39201611858565b90565b6101ed565b6101e9565b5190565b90565b6118cb6118c66118d0926118b4565b61035f565b610221565b90565b5f9103126118dd57565b6101e9565b60209181520190565b60200190565b6118fa906110e1565b9052565b61190790610221565b9052565b61191490610487565b9052565b611921906101f1565b9052565b60209181520190565b61194d61195660209361195b9361194481610ef2565b93848093611925565b95869101610eff565b610f22565b0190565b6119f79160e061010082019261197b5f8201515f850190611918565b61198d60208201516020850190611918565b61199f60408201516040850190611918565b6119b160608201516060850190611918565b6119c360808201516080850190611918565b6119d560a082015160a0850190611918565b6119e760c082015160c0850190611918565b01519060e081840391015261192e565b90565b611a829160c0611a7160e08301611a175f8601515f8601906118f1565b611a29602086015160208601906118f1565b611a3b604086015160408601906118fe565b611a4d6060860151606086019061190b565b611a5f6080860151608086019061190b565b60a085015184820360a086015261195f565b9201519060c081840391015261195f565b90565b90611a8f916119fa565b90565b60200190565b90611aac611aa5836118b0565b80926118e2565b9081611abd602083028401946118eb565b925f915b838310611ad057505050505090565b90919293946020611af2611aec83856001950387528951611a85565b97611a92565b9301930191939290611ac1565b611b149160208201915f818403910152611a98565b90565b5f90565b606090565b9190604083820312611b5a57611b5390611b3a60406110c8565b93611b47825f830161163c565b5f86015260200161164b565b6020830152565b611087565b90604082820312611b7857611b75915f01611b20565b90565b6101e9565b611b8790516110e1565b90565b90565b611ba1611b9c611ba692611b8a565b61035f565b6110e1565b90565b90565b611bc0611bbb611bc592611ba9565b61035f565b6110e1565b90565b611bd19061037e565b90565b611bde9051610221565b90565b90611bf3611bee8361110e565b6110c8565b918252565b3d5f14611c1357611c083d611be1565b903d5f602084013e5b565b611c1b611b1b565b90611c11565b611c2b9051610487565b90565b905090565b611c3e5f8092611c2e565b0190565b611c4b90611c33565b90565b151590565b611c5c90611c4e565b9052565b91611c969391611c8891611c7b60608601925f870190611c53565b8482036020860152610f2c565b916040818403910152610f2c565b90565b90611ccc90611cc2611cbd611cb86008611cb246611032565b90610427565b611062565b61107b565b92908101906114fa565b90611cf16020611cdb8361152f565b63750c67ba90611ce96101e3565b93849261153f565b82528180611d01600482016102e4565b03915afa908115612062575f91612034575b50611d2e611d28611d2346611032565b6101f1565b916101f1565b03611f45575b50611d3d611b17565b50611d46611b1b565b505f80611d6c610160602085015101516020611d6182610ef2565b818301019101611b5f565b611d7b82602086015101611b7d565b611d8e611d8860fe611b8d565b916110e1565b148214611e2257611dbc82611db5611db06008611daa46611032565b90610427565b611062565b9201611bd4565b6101406020860151015190602082019151925af1906101406020611dde611bf8565b935b93920151015191611e1d7fc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb39807944693611e146101e3565b93849384611c60565b0390a1565b611e3182602086015101611b7d565b611e44611e3e60ff611bac565b916110e1565b148214611e9b57611e7282611e6b611e666008611e6046611032565b90610427565b611062565b9201611bd4565b6101406020860151015190602082019151925af1906101406020611e94611bf8565b935b611de0565b611ead61014060208601510151610ef2565b611ebf611eb9846118b7565b91610221565b148214611f0957611edd82611ed660208401611c21565b9201611bd4565b611ee56101e3565b9081611ef081611c42565b03925af1906101406020611f02611bf8565b935b611e96565b611f1d82611f1630611bc8565b9201611bd4565b6101406020860151015190602082019151925af1906101406020611f3f611bf8565b93611f04565b611f81611f6b610120602085015101516020611f6082610ef2565b8183010191016115db565b6020611f7682610ef2565b81830101910161187b565b90611f8b826118b0565b611f9d611f975f6118b7565b91610221565b03611fa9575b50611d34565b611fb29061152f565b906377c2719c90823b1561202f57611fe992611fde5f8094611fd26101e3565b9687958694859361153f565b835260048301611aff565b03925af1801561202a57611ffe575b80611fa3565b61201d905f3d8111612023575b612015818361109f565b8101906118d3565b5f611ff8565b503d61200b565b611572565b61153b565b612055915060203d811161205b575b61204d818361109f565b810190611554565b5f611d13565b503d612043565b611572565b5f90565b90565b5061207d9060208101906110fb565b90565b5061208f906020810190610238565b90565b506120a190602081019061055e565b90565b5f80fd5b9035600161010003823603038112156120bf570190565b6120a4565b506120d3906020810190610212565b90565b5f80fd5b5f80fd5b903560016020038236030381121561211f57016020813591019167ffffffffffffffff821161211a57600182023603831361211557565b6120da565b6120d6565b6120a4565b919061213e816121378161214395611925565b8095611131565b610f22565b0190565b6122209161221261010082019261216c6121635f8301836120c4565b5f850190611918565b61218661217c60208301836120c4565b6020850190611918565b6121a061219660408301836120c4565b6040850190611918565b6121ba6121b060608301836120c4565b6060850190611918565b6121d46121ca60808301836120c4565b6080850190611918565b6121ee6121e460a08301836120c4565b60a0850190611918565b6122086121fe60c08301836120c4565b60c0850190611918565b60e08101906120de565b9160e0818503910152612124565b90565b6122e4916122d66122cb60e083016122496122405f87018761206e565b5f8601906118f1565b612263612259602087018761206e565b60208601906118f1565b61227d6122736040870187612080565b60408601906118fe565b61229761228d6060870187612092565b606086019061190b565b6122b16122a76080870187612092565b608086019061190b565b6122be60a08601866120a8565b84820360a0860152612147565b9260c08101906120a8565b9060c0818403910152612147565b90565b906122f191612223565b90565b9035600160e0038236030381121561230a570190565b6120a4565b60200190565b9181612320916118e2565b90816123316020830284019461206b565b92835f925b8484106123465750505050505090565b909192939495602061237161236b83856001950388526123668b886122f4565b6122e7565b9861230f565b940194019294939190612336565b90916123969260208301925f818503910152612315565b90565b6123a56123aa91610fae565b610c26565b90565b6123b79054612399565b90565b6123c66123cb91610fae565b6109ea565b90565b6123d890546123ba565b90565b60181c90565b6123ed6123f2916123db565b610d41565b90565b6123ff90546123e1565b90565b9060208282031261241b57612418915f0161163c565b90565b6101e9565b634e487b7160e01b5f52602260045260245ffd5b9060016002830492168015612454575b602083101461244f57565b612420565b91607f1691612444565b5f5260205f2090565b905f929180549061248161247a83612434565b8094610ef6565b916001811690815f146124d8575060011461249c575b505050565b6124a9919293945061245e565b915f925b8184106124c057505001905f8080612497565b600181602092959395548486015201910192906124ad565b92949550505060ff19168252151560200201905f8080612497565b92906125299161251c612537969461251260808801945f89019061060c565b60208701906106eb565b8482036040860152612467565b916060818403910152610f2c565b90565b91939061259b9161257460209561254f612067565b5061256661255b6101e3565b93849289840161237f565b87820181038252038261109f565b61257e60046123ad565b9161258960076123ce565b9061259460076123f5565b929361580f565b926125d26125b06125ab5f610fc7565b61038a565b916125dd6385fdd54291949660066125c66101e3565b9889978896879661153f565b8652600486016124f3565b03915afa908115612621575f916125f3575b5090565b612614915060203d811161261a575b61260c818361109f565b810190612402565b5f6125ef565b503d612602565b611572565b5f90565b612632612626565b90565b5f90565b612641612635565b90565b61264c612635565b90565b906126619161265c61585e565b6126bc565b565b5f1b90565b9061267460ff91612663565b9181191691161790565b61268790610667565b90565b60f81c90565b6126999061268a565b90565b906126b16126ac6126b89261267e565b612690565b8254612668565b9055565b6126cb6126d092916009610c0a565b61269c565b565b906126dc9161264f565b565b6126e661585e565b6126ee612718565b565b6127046126ff612709926118b4565b61035f565b610354565b90565b612715906126f0565b90565b6127296127245f61270c565b6158b3565b565b6127336126de565b565b61274091369161113c565b90565b9261276561276b92956127719795612759612067565b50969492909592612735565b93612735565b9261595b565b90565b5f90565b612780612774565b5061278b6002611062565b90565b60209181520190565b5f7f4d45500000000000000000000000000000000000000000000000000000000000910152565b6127cb600360209261278e565b6127d481612797565b0190565b6127ed9060208101905f8183039101526127be565b90565b156127f757565b6127ff6101e3565b62461bcd60e51b815280612815600482016127d8565b0390fd5b9061285d9493929161285861282d46611032565b61285261284c61284761284233946008610427565b611062565b610487565b91610487565b146127f0565b612923565b565b634e487b7160e01b5f52601160045260245ffd5b61288261288891939293610221565b92610221565b820180921161289357565b61285f565b1561289f57565b5f80fd5b9692612912956128f16129209a9893966128e76128fb946128dd6129059860208f6128d661010082019e5f8301906106eb565b01906106eb565b60408d0190610493565b60608b0190610493565b608089019061060c565b60a08701906106eb565b84820360c0860152612467565b9160e0818403910152610f2c565b90565b92936129b29161298e9161298261293a60046123ad565b8361294560076123ce565b61294f60076123f5565b918b9361297d879561296e6129626101e3565b9788926020840161237f565b6020820181038252038661109f565b61580f565b9686928691929361253a565b6129ab6129a56129a034938690612873565b610221565b91610221565b1015612898565b6129c36129be5f610fc7565b61038a565b9063209afe5690349291927f0000000000000000000000000000000000000000000000000000000000000000937f0000000000000000000000000000000000000000000000000000000000000000967f0000000000000000000000000000000000000000000000000000000000000000903394979060069091873b15612aaf575f99612a6197612a6c95612a556101e3565b9d8e9c8d9b8c9a61153f565b8a5260048a016128a3565b03925af18015612aaa57612a7e575b50565b612a9d905f3d8111612aa3575b612a95818361109f565b8101906118d3565b5f612a7b565b503d612a8b565b611572565b61153b565b90612ac194939291612819565b565b612ad490612acf615a77565b612afa565b612adc615ae3565b565b612af2612aed612af7926118b4565b610c20565b610667565b90565b612b3990612b34612b1c612b17612b105f6118b7565b6009610c0a565b6123ad565b612b2e612b285f612ade565b91610667565b14612898565b6133ed565b565b5f80fd5b5f80fd5b5f80fd5b90359060016101800381360303821215612b5f570190565b612b3b565b35612b6e816110e7565b90565b612b85612b80612b8a926118b4565b61035f565b6110e1565b90565b90565b612ba4612b9f612ba992612b8d565b61035f565b6110e1565b90565b90565b612bc3612bbe612bc892612bac565b61035f565b6110e1565b90565b903590600160200381360303821215612c0d570180359067ffffffffffffffff8211612c0857602001916001820236038313612c0357565b612b43565b612b3f565b612b3b565b9190604083820312612c4c57612c4590612c2c60406110c8565b93612c39825f8301610238565b5f86015260200161055e565b6020830152565b611087565b90604082820312612c6a57612c67915f01612c12565b90565b6101e9565b612c7881610354565b03612c7f57565b5f80fd5b90503590612c9082612c6f565b565b919061016083820312612d7657612d6e90612cae6101606110c8565b93612cbb825f83016110fb565b5f860152612ccc82602083016110fb565b6020860152612cde8260408301611337565b6040860152612cf08260608301611337565b6060860152612d028260808301612c83565b6080860152612d148260a08301612c83565b60a0860152612d268260c0830161055e565b60c0860152612d388260e0830161055e565b60e0860152612d4b82610100830161055e565b610100860152612d5f826101208301610238565b61012086015261014001610238565b610140830152565b611087565b9061016082820312612d9557612d92915f01612c92565b90565b6101e9565b612da49051610742565b90565b612db19051610354565b90565b612dbf6101006110c8565b90565b90612dcc906110e1565b9052565b90612dda90610742565b9052565b90612de890610354565b9052565b90612df690610487565b9052565b90612e0490610221565b9052565b612e1190610362565b90565b612e1d90612e08565b90565b612e299061037e565b90565b612e3581611c4e565b03612e3c57565b5f80fd5b90505190612e4d82612e2c565b565b90602082820312612e6857612e65915f01612e40565b90565b6101e9565b916020612e8e929493612e8760408201965f830190610493565b019061060c565b565b612e9990610362565b90565b612ea590612e90565b90565b612eb19061037e565b90565b612ebd90610742565b9052565b612eca90610354565b9052565b9060e080612f5c93612ee65f8201515f8601906118f1565b612ef860208201516020860190612eb4565b612f0a60408201516040860190612ec1565b612f1c6060820151606086019061190b565b612f2e6080820151608086019061190b565b612f4060a082015160a086019061190b565b612f5260c082015160c08601906118fe565b01519101906118fe565b565b9190612f72905f6101008501940190612ece565b565b63ffffffff1690565b63ffffffff60e01b1690565b612f9d612f98612fa292612f74565b61153f565b612f7d565b90565b612faf60406110c8565b90565b90602080612fd493612fca5f8201515f8601906118fe565b015191019061190b565b565b9190612fe9905f60408501940190612fb2565b565b9190610100838203126130935761308c906130076101006110c8565b93613014825f83016110fb565b5f86015261302582602083016110fb565b60208601526130378260408301610238565b60408601526130498260608301610238565b606086015261305b826080830161055e565b608086015261306d8260a0830161055e565b60a086015261307f8260c0830161055e565b60c086015260e001610238565b60e0830152565b611087565b90610100828203126130b2576130af915f01612feb565b90565b6101e9565b606090565b634e487b7160e01b5f52603260045260245ffd5b5190565b906130de826130d0565b8110156130ef576020809102010190565b6130bc565b61310861310361310d92612b8d565b61035f565b610221565b90565b61311a60c06110c8565b90565b52565b60209181520190565b60200190565b9061313c8160209361190b565b0190565b60200190565b9061316361315d613156846130d0565b8093613120565b92613129565b905f5b8181106131735750505090565b90919261318c613186600192865161312f565b94613140565b9101919091613166565b906132089060a0806131eb60c084016131b55f8801515f8701906118f1565b6131c7602088015160208701906118fe565b6131d9604088015160408701906118fe565b60608701518582036060870152613146565b946131fe6080820151608086019061190b565b01519101906118fe565b90565b6132209160208201915f818403910152613196565b90565b3561322d81611323565b90565b3561323a816101fe565b90565b356132478161054a565b90565b3561325481610224565b90565b6132626101806110c8565b90565b9061326f906101f1565b9052565b52565b903590600160e0038136030382121561328d570190565b612b3b565b61329c60406110c8565b90565b6132aa90369061125a565b90565b52565b52565b6133a19161016061338f61337b61018084016132d55f8701515f8701906118f1565b6132e760208701516020870190612eb4565b6132f960408701516040870190611918565b61330b60608701516060870190611918565b61331d60808701516080870190611918565b61332f60a087015160a0870190611918565b61334160c087015160c087019061190b565b61335360e087015160e087019061190b565b6133676101008701516101008701906118fe565b61012086015185820361012087015261192e565b61014085015184820361014086015261192e565b9201519061016081840391015261192e565b90565b6133d29160206133c1604083015f8501518482035f8601526119fa565b9201519060208184039101526132b3565b90565b6133ea9160208201915f8184039101526133a4565b90565b6133f5612067565b506133fe611b1b565b90613407611b1b565b506134205f61341a836020810190612b47565b01612b64565b61343261342c5f612b71565b916110e1565b145f146139785761346a5f61346461345c613451856020810190612b47565b610160810190612bcb565b810190612c51565b01611bd4565b9161349061348a61347f846020810190612b47565b610160810190612bcb565b90612735565b905b82602081016134a091612b47565b5f016134ab90612b64565b9183602081016134ba91612b47565b6020016134c690613223565b9184602081016134d591612b47565b6040016134e190613230565b85602081016134ef91612b47565b6060016134fb90613230565b866020810161350991612b47565b60800161351590613230565b876020810161352391612b47565b60a00161352f90613230565b886020810161353d91612b47565b60c0016135499061323d565b90896020810161355891612b47565b60e0016135649061323d565b928a6020810161357391612b47565b610100016135809061324a565b948b6020810161358f91612b47565b610120810161359d91612bcb565b979098999a6135aa613257565b9c5f8e01906135b891612dc2565b60208d01906135c691612dd0565b60408c01906135d491613265565b60608b01906135e291613265565b60808a01906135f091613265565b60a08901906135fe91613265565b60c088019061360c91612dec565b60e087019061361a91612dec565b61010086019061362991612dfa565b61363291612735565b61012084019061364191613273565b61014083019061365091613273565b61016082019061365f91613273565b61366960046123ad565b90826020810161367891612b47565b60c0016136849061323d565b836020810161369291612b47565b60200161369e90613223565b84602081016136ac91612b47565b6040016136b890613230565b91855f81016136c691613276565b936136cf613292565b946136d99061329f565b5f8601906136e6916132ad565b60208501906136f4916132b0565b6136fc6101e3565b8094602082019061370c916133d5565b60208201810382520361371f908561109f565b6137289461580f565b916137325f610fc7565b61373b9061038a565b6385fdd54290836020810161374f91612b47565b6101000161375c9061324a565b8361376691612873565b91846020810161377591612b47565b60600161378190613230565b906006938761378e6101e3565b95869461379b869561153f565b855260048501936137ab946124f3565b03815a93602094fa8015613973576137f76137fc91613802935f91613945575b50936137f234956137ec6101006137e68a6020810190612b47565b0161324a565b90612873565b612873565b610221565b91610221565b10613922576138186138135f610fc7565b61038a565b9163209afe56349390939161383c6080613836866020810190612b47565b01613230565b9261385660a0613850876020810190612b47565b01613230565b9561387060e061386a886020810190612b47565b0161323d565b33936138a5606061389f61389461010061388e8d6020810190612b47565b0161324a565b9a6020810190612b47565b01613230565b9060069091873b1561391d575f996138cf976138da956138c36101e3565b9d8e9c8d9b8c9a61153f565b8a5260048a016128a3565b03925af18015613918576138ec575b50565b61390b905f3d8111613911575b613903818361109f565b8101906118d3565b5f6138e9565b503d6138f9565b611572565b61153b565b61392a6101e3565b631e9acf1760e31b815280613941600482016102e4565b0390fd5b613966915060203d811161396c575b61395e818361109f565b810190612402565b5f6137cb565b503d613954565b611572565b6139905f61398a836020810190612b47565b01612b64565b6139a361399d6001612b90565b916110e1565b145f14613ddb576139d16139c96139be836020810190612b47565b610160810190612bcb565b810190613098565b6139d96130b7565b926139e660808301611c21565b613a006139fa6139f55f61270c565b610487565b91610487565b145f14613b8c5750613b76613b856004613b335f613afb613a2360408801611bd4565b98613a48613a308461270c565b613a4383613a3d876118b7565b906130d4565b612dec565b613a70613a5760a08a01611c21565b613a6b83613a6560016130f4565b906130d4565b612dec565b613af288613ae9613a8360208301611b7d565b93613ae0613ad7613a9660408601611bd4565b92613ad28a9194613ac9613ab860e0613ab160c08c01611c21565b9a01611bd4565b9a613ac1613110565b9d8e01612dc2565b60208c01612dfa565b6118b7565b60408901612dfa565b6060870161311d565b60808501612dec565b60a08301612dfa565b613b24613b0b630c07a79e612f89565b91613b146101e3565b948593602085019081520161320b565b6020820181038252038261109f565b925b613b62613b4560c0889301611c21565b613b59613b50612fa5565b935f8501612dfa565b60208301612dec565b613b6a6101e3565b92839160208301612fd6565b6020820181038252038261109f565b905b613492565b9092613bb5613b9d60808601611c21565b613bb083613baa5f6118b7565b906130d4565b612dec565b613bda613bc15f61270c565b613bd583613bcf60016130f4565b906130d4565b612dec565b613c66613be85f8601611b7d565b91613c5d86613c54613bfc60408301611bd4565b93613c4b613c0c60608501611bd4565b613c42613c2460e0613c1d30611bc8565b9701611bd4565b97613c39613c30613110565b9b5f8d01612dc2565b60208b01612dfa565b60408901612dfa565b6060870161311d565b60808501612dec565b60a08301612dfa565b613c9a613c7d613c7860808701611c21565b612e14565b33613c8730611bc8565b90613c9460408901611bd4565b92615b29565b613cb6613cb1613cac60808701611c21565b612e14565b612e20565b90602063095ea7b392613cc96005611062565b90613cf25f613cda60408b01611bd4565b96613cfd613ce66101e3565b9889968795869461153f565b845260048401612e6d565b03925af1908115613dd657613d5792602092613dab575b50613d2f613d2a613d256005611062565b612e9c565b612ea8565b613d4c5f630c07a79e613d406101e3565b9687958694859361153f565b83526004830161320b565b03925af1908115613da657613b7691613b85915f91613d78575b5094613b35565b613d99915060203d8111613d9f575b613d91818361109f565b810190612402565b5f613d71565b503d613d87565b611572565b613dca90833d8111613dcf575b613dc2818361109f565b810190612e4f565b613d14565b503d613db8565b611572565b613df35f613ded836020810190612b47565b01612b64565b613e06613e006002612baf565b916110e1565b145f1461425357613e34613e2c613e21836020810190612b47565b610160810190612bcb565b810190612d7b565b91613e4160c08401611c21565b613e5b613e55613e505f61270c565b610487565b91610487565b145f14613ff657508161012001613e7190611bd4565b9180602001613e7f90611b7d565b81606001613e8c90612d9a565b8260a001613e9990612da7565b5f613ea39061270c565b8460e001613eb090611c21565b8561010001613ebe90611c21565b918661012001613ecd90611bd4565b938761014001613edc90611bd4565b95613ee5612db4565b975f890190613ef391612dc2565b6020880190613f0191612dd0565b6040870190613f0f91612dde565b6060860190613f1d91612dec565b6080850190613f2b91612dec565b60a0840190613f3991612dec565b60c0830190613f4791612dfa565b60e0820190613f5591612dfa565b63f87f5859613f6390612f89565b613f6b6101e3565b9182916020830190815260040190613f8291612f5e565b602082018103825203613f95908261109f565b90613fef613fe0915b613fcc613faf610100889301611c21565b613fc3613fba612fa5565b935f8501612dfa565b60208301612dec565b613fd46101e3565b92839160208301612fd6565b6020820181038252038261109f565b905b613b87565b825f0161400290611b7d565b8360400161400f90612d9a565b8460800161401c90612da7565b8560c00161402990611c21565b5f6140339061270c565b3061403d90611bc8565b91886101200161404c90611bd4565b93896101400161405b90611bd4565b95614064612db4565b975f89019061407291612dc2565b602088019061408091612dd0565b604087019061408e91612dde565b606086019061409c91612dec565b60808501906140aa91612dec565b60a08401906140b891612dec565b60c08301906140c691612dfa565b60e08201906140d491612dfa565b806060016140e190611c21565b6140ea90612e14565b33306140f590611bc8565b8360c00161410290611bd4565b9161410c93615b29565b8060600161411990611c21565b61412290612e14565b61412b90612e20565b9063095ea7b39161413c6005611062565b928260c00161414a90611bd4565b906141536101e3565b948592614160849361153f565b8352600483019161417092612e6d565b03815a6020945f91f190811561424e576141cf92602092614223575b506141a76141a261419d6005611062565b612e9c565b612ea8565b6141c45f63f87f58596141b86101e3565b9687958694859361153f565b835260048301612f5e565b03925af190811561421e57613fe091613fef915f916141f0575b5094613f9e565b614211915060203d8111614217575b614209818361109f565b810190612402565b5f6141e9565b503d6141ff565b611572565b61424290833d8111614247575b61423a818361109f565b810190612e4f565b61418c565b503d614230565b611572565b61426b5f614265836020810190612b47565b01612b64565b61427e61427860fe611b8d565b916110e1565b145f146142e3576142b65f6142b06142a861429d856020810190612b47565b610160810190612bcb565b810190612c51565b01611bd4565b916142dc6142d66142cb846020810190612b47565b610160810190612bcb565b90612735565b905b613ff1565b6142fb5f6142f5836020810190612b47565b01612b64565b61430e61430860ff611bac565b916110e1565b145f14614372576143465f61434061433861432d856020810190612b47565b610160810190612bcb565b810190612c51565b01611bd4565b9161436c61436661435b846020810190612b47565b610160810190612bcb565b90612735565b906142de565b61437a6101e3565b634157854360e11b815280614391600482016102e4565b0390fd5b61439e90612ac3565b565b906143b2916143ad61585e565b6143fe565b565b906143c560018060a01b0391612663565b9181191691161790565b6143d89061037e565b90565b90565b906143f36143ee6143fa926143cf565b6143db565b82546143b4565b9055565b61440d61441292916008610427565b6143de565b565b9061441e916143a0565b565b903560016101800382360303811215614437570190565b6120a4565b5061444b906020810190611337565b90565b61459f91614590614584614567610180840161447861446f5f88018861206e565b5f8701906118f1565b614492614488602088018861443c565b6020870190612eb4565b6144ac6144a260408801886120c4565b6040870190611918565b6144c66144bc60608801886120c4565b6060870190611918565b6144e06144d660808801886120c4565b6080870190611918565b6144fa6144f060a08801886120c4565b60a0870190611918565b61451461450a60c0880188612092565b60c087019061190b565b61452e61452460e0880188612092565b60e087019061190b565b61454a61453f610100880188612080565b6101008701906118fe565b6145586101208701876120de565b90868303610120880152612124565b6145756101408601866120de565b90858303610140870152612124565b926101608101906120de565b91610160818503910152612124565b90565b6145e1916145d36145c8604083016145bc5f8601866122f4565b8482035f860152612223565b926020810190614420565b90602081840391015261444e565b90565b6145f99160208201915f8184039101526145a2565b90565b614604612067565b5061462061462f826146146101e3565b928391602083016145e4565b6020820181038252038261109f565b90614638612067565b906146515f61464b836020810190612b47565b01612b64565b61466361465d5f612b71565b916110e1565b145f146147c057602091506146fd6146a15f61469b6146936146888688810190612b47565b610160810190612bcb565b810190612c51565b01611bd4565b935b6146ad60046123ad565b906146c660c06146c08688810190612b47565b0161323d565b6146dd866146d78782810190612b47565b01613223565b906146f660406146f0888a810190612b47565b01613230565b929361580f565b9261476c61471261470d5f610fc7565b61038a565b91614777614755606061474f6147456385fdd5429561474061010061473a8c8e810190612b47565b0161324a565b612873565b9789810190612b47565b01613230565b9660066147606101e3565b9889978896879661153f565b8652600486016124f3565b03915afa9081156147bb575f9161478d575b5090565b6147ae915060203d81116147b4575b6147a6818361109f565b810190612402565b5f614789565b503d61479c565b611572565b6147d85f6147d2836020810190612b47565b01612b64565b6147eb6147e56001612b90565b916110e1565b145f1461486d57614819614811614806836020810190612b47565b610160810190612bcb565b810190613098565b61482560808201611c21565b61483f6148396148345f61270c565b610487565b91610487565b14614853575b506146fd602092935b6146a3565b6020925061486660406146fd9201611bd4565b9250614845565b6148855f61487f836020810190612b47565b01612b64565b6148986148926002612baf565b916110e1565b145f1461491a576148c66148be6148b3836020810190612b47565b610160810190612bcb565b810190612d7b565b6148d260c08201611c21565b6148ec6148e66148e15f61270c565b610487565b91610487565b146148ff575b506146fd6020929361484e565b602092506149136101206146fd9201611bd4565b92506148f2565b6149226101e3565b634157854360e11b815280614939600482016102e4565b0390fd5b61494e9061494961585e565b614950565b565b8061496b6149656149605f61270c565b610487565b91610487565b1461497b57614979906158b3565b565b6149a56149875f61270c565b61498f6101e3565b918291631e4fbdf760e01b8352600483016104a0565b0390fd5b6149b29061493d565b565b93929190336149dc6149d66149d16149cc6001610fc7565b61038a565b610487565b91610487565b036149ec576149ea94614a0f565b565b6149f46101e3565b637d92a0f560e11b815280614a0b600482016102e4565b0390fd5b91614a209492939190919293615c17565b565b90614a2f949392916149b4565b565b614a39611b1b565b50614a42612067565b50614a4b611b1b565b614a53611b1b565b50614a6c5f614a66846020810190612b47565b01612b64565b614a7e614a785f612b71565b916110e1565b145f14614d1457614ab65f614ab0614aa8614a9d866020810190612b47565b610160810190612bcb565b810190612c51565b01611bd4565b50614adc614ad6614acb846020810190612b47565b610160810190612bcb565b90612735565b905b8260208101614aec91612b47565b5f01614af790612b64565b918360208101614b0691612b47565b602001614b1290613223565b918460208101614b2191612b47565b604001614b2d90613230565b8560208101614b3b91612b47565b606001614b4790613230565b8660208101614b5591612b47565b608001614b6190613230565b8760208101614b6f91612b47565b60a001614b7b90613230565b8860208101614b8991612b47565b60c001614b959061323d565b908960208101614ba491612b47565b60e001614bb09061323d565b928a60208101614bbf91612b47565b61010001614bcc9061324a565b948b60208101614bdb91612b47565b6101208101614be991612bcb565b979098999a614bf6613257565b9c5f8e0190614c0491612dc2565b60208d0190614c1291612dd0565b60408c0190614c2091613265565b60608b0190614c2e91613265565b60808a0190614c3c91613265565b60a0890190614c4a91613265565b60c0880190614c5891612dec565b60e0870190614c6691612dec565b610100860190614c7591612dfa565b614c7e91612735565b610120840190614c8d91613273565b610140830190614c9c91613273565b610160820190614cab91613273565b905f8101614cb891613276565b90614cc1613292565b91614ccb9061329f565b5f830190614cd8916132ad565b6020820190614ce6916132b0565b614cee6101e3565b80916020820190614cfe916133d5565b602082018103825203614d11908261109f565b90565b614d2c5f614d26846020810190612b47565b01612b64565b614d3f614d396001612b90565b916110e1565b145f1461517f57614d6d614d65614d5a846020810190612b47565b610160810190612bcb565b810190613098565b614d756130b7565b91614d8260808301611c21565b614d9c614d96614d915f61270c565b610487565b91610487565b145f14614f285750614f2181614efe614ee160c06004614ed85f614ea0614dc86040614f129b01611bd4565b9a614ded614dd58461270c565b614de883614de2876118b7565b906130d4565b612dec565b614e15614dfc60a08b01611c21565b614e1083614e0a60016130f4565b906130d4565b612dec565b614e9789614e8e614e2860208301611b7d565b93614e85614e7c614e3b60408601611bd4565b92614e778a9194614e6e8f60e0614e56614e5d928c01611c21565b9a01611bd4565b9a614e66613110565b9d8e01612dc2565b60208c01612dfa565b6118b7565b60408901612dfa565b6060870161311d565b60808501612dec565b60a08301612dfa565b614ec9614eb0630c07a79e612f89565b91614eb96101e3565b948593602085019081520161320b565b6020820181038252038261109f565b965b9301611c21565b614ef5614eec612fa5565b935f8501612dfa565b60208301612dec565b614f066101e3565b92839160208301612fd6565b6020820181038252038261109f565b905b614ade565b91614f50614f3860808401611c21565b614f4b83614f455f6118b7565b906130d4565b612dec565b614f75614f5c5f61270c565b614f7083614f6a60016130f4565b906130d4565b612dec565b615002614f835f8401611b7d565b91614ff9614f9360408601611bd4565b91614ff0614fa360608801611bd4565b91614fe7614fb030611bc8565b93614fde614fc060e08c01611bd4565b97614fd5614fcc613110565b9b5f8d01612dc2565b60208b01612dfa565b60408901612dfa565b6060870161311d565b60808501612dec565b60a08301612dfa565b9061503761501a61501560808401611c21565b612e14565b3361502430611bc8565b9061503160408601611bd4565b92615b29565b61505361504e61504960808401611c21565b612e14565b612e20565b91602063095ea7b3936150666005611062565b9061508f5f61507760408801611bd4565b9761509a6150836101e3565b998a968795869461153f565b845260048401612e6d565b03925af190811561517a576150f49360209261514f575b506150cc6150c76150c26005611062565b612e9c565b612ea8565b6150e95f630c07a79e6150dd6101e3565b9788958694859361153f565b83526004830161320b565b03925af190811561514a57614efe614ee160c0614f2194614f12965f9161511c575b50614eda565b61513d915060203d8111615143575b615135818361109f565b810190612402565b5f615116565b503d61512b565b611572565b61516e90833d8111615173575b615166818361109f565b810190612e4f565b6150b1565b503d61515c565b611572565b6151975f615191846020810190612b47565b01612b64565b6151aa6151a46002612baf565b916110e1565b145f14615601576151d86151d06151c5846020810190612b47565b610160810190612bcb565b810190612d7b565b906151e560c08301611c21565b6151ff6151f96151f45f61270c565b610487565b91610487565b145f1461539a5750806101200161521590611bd4565b908060200161522390611b7d565b8160600161523090612d9a565b8260a00161523d90612da7565b5f6152479061270c565b8460e00161525490611c21565b856101000161526290611c21565b91866101200161527190611bd4565b93876101400161528090611bd4565b95615289612db4565b975f89019061529791612dc2565b60208801906152a591612dd0565b60408701906152b391612dde565b60608601906152c191612dec565b60808501906152cf91612dec565b60a08401906152dd91612dec565b60c08301906152eb91612dfa565b60e08201906152f991612dfa565b63f87f585961530790612f89565b61530f6101e3565b918291602083019081526004019061532691612f5e565b602082018103825203615339908261109f565b9161538491615370615353610100615393945b9301611c21565b61536761535e612fa5565b935f8501612dfa565b60208301612dec565b6153786101e3565b92839160208301612fd6565b6020820181038252038261109f565b905b614f23565b90805f016153a790611b7d565b816040016153b490612d9a565b826080016153c190612da7565b8360c0016153ce90611c21565b5f6153d89061270c565b306153e290611bc8565b9186610120016153f190611bd4565b93876101400161540090611bd4565b95615409612db4565b975f89019061541791612dc2565b602088019061542591612dd0565b604087019061543391612dde565b606086019061544191612dec565b608085019061544f91612dec565b60a084019061545d91612dec565b60c083019061546b91612dfa565b60e082019061547991612dfa565b908160600161548790611c21565b61549090612e14565b333061549b90611bc8565b8460c0016154a890611bd4565b916154b293615b29565b816060016154bf90611c21565b6154c890612e14565b6154d190612e20565b9163095ea7b3926154e26005611062565b938260c0016154f090611bd4565b906154f96101e3565b958692615506849361153f565b8352600483019161551692612e6d565b03815a6020945f91f19081156155fc57615575936020926155d1575b5061554d6155486155436005611062565b612e9c565b612ea8565b61556a5f63f87f585961555e6101e3565b9788958694859361153f565b835260048301612f5e565b03925af19081156155cc5761537061535361010061539394615384965f9161559e575b5061534c565b6155bf915060203d81116155c5575b6155b7818361109f565b810190612402565b5f615598565b503d6155ad565b611572565b6155f090833d81116155f5575b6155e8818361109f565b810190612e4f565b615532565b503d6155de565b611572565b5061561a5f615614836020810190612b47565b01612b64565b61562d61562760ff611bac565b916110e1565b145f146156b5576156655f61565f61565761564c856020810190612b47565b610160810190612bcb565b810190612c51565b01611bd4565b5061568b61568561567a836020810190612b47565b610160810190612bcb565b90612735565b6156b06156aa61569f846020810190612b47565b610140810190612bcb565b90612735565b615395565b6156bd6101e3565b634157854360e11b8152806156d4600482016102e4565b0390fd5b6156e06101e3565b637fca089960e11b8152806156f7600482016102e4565b0390fd5b61570490610362565b90565b61571b61571661572092610354565b61035f565b610221565b90565b90565b61573261573791610667565b615723565b9052565b90565b61574a61574f91610221565b61573b565b9052565b60e81b90565b61576290615753565b90565b61577161577691610742565b615759565b9052565b60c01b90565b6157899061577a565b90565b61579861579d916101f1565b615780565b9052565b6157c66157bd926020926157b481610ef2565b94858093611c2e565b93849101610eff565b0190565b6008936157fd602061580c9897956157f56001866157ed6003986158059a615726565b01809261573e565b018092615765565b01809261578c565b01906157a1565b90565b9261584c91929461583461582f61585b96615828611b1b565b50956156fb565b615707565b959190916158406101e3565b968795602087016157ca565b6020820181038252038261109f565b90565b615866612778565b61587f615879615874615e8e565b610487565b91610487565b0361588657565b6158af615891615e8e565b6158996101e3565b91829163118cdaa760e01b8352600483016104a0565b0390fd5b6158bd6002611062565b6158c88260026143de565b906158fc6158f67f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936143cf565b916143cf565b916159056101e3565b8061590f816102e4565b0390a3565b929061594a9161593d615958969461593360808801945f89019061060c565b60208701906106eb565b8482036040860152610f2c565b916060818403910152610f2c565b90565b9261599d6020939461596b612067565b506159a861598061597b5f610fc7565b61038a565b936385fdd5429295976159916101e3565b9889978896879661153f565b865260048601615914565b03915afa9081156159ec575f916159be575b5090565b6159df915060203d81116159e5575b6159d7818361109f565b810190612402565b5f6159ba565b503d6159cd565b611572565b90565b615a00615a0591610fae565b6159f1565b90565b615a1290546159f4565b90565b615a29615a24615a2e92612bac565b61035f565b610221565b90565b615a3b6002615a15565b90565b90615a4a5f1991612663565b9181191691161790565b90565b90615a6c615a67615a7392610bee565b615a54565b8254615a3e565b9055565b615a816003615a08565b615a9a615a94615a8f615a31565b610221565b91610221565b14615ab357615ab1615aaa615a31565b6003615a57565b565b615abb6101e3565b633ee5aeb560e01b815280615ad2600482016102e4565b0390fd5b615ae060016130f4565b90565b615af5615aee615ad6565b6003615a57565b565b604090615b20615b279496959396615b1660608401985f850190610493565b6020830190610493565b019061060c565b565b600492615b63615b779593615b729394615b4a6323b872dd92949192612f89565b93615b536101e3565b9788956020870190815201615af7565b6020820181038252038361109f565b615e9b565b565b615b8d615b88615b9292610221565b61035f565b610354565b90565b5f7f496e76616c696420636f6e747261637400000000000000000000000000000000910152565b615bc9601060209261278e565b615bd281615b95565b0190565b615beb9060208101905f818303910152615bbc565b90565b15615bf557565b615bfd6101e3565b62461bcd60e51b815280615c1360048201615bd6565b0390fd5b50615c65939291615c56615c50615c4b615c46615c40615c3b615c5c976008610427565b611062565b94615b79565b61037e565b610487565b91610487565b14615bee565b908101906114fa565b615ca1615c8b610120602084015101516020615c8082610ef2565b8183010191016115db565b6020615c9682610ef2565b81830101910161187b565b615caa816118b0565b615cbc615cb65f6118b7565b91610221565b03615de1575b50615ccb611b17565b50615cd4611b1b565b505f80615cfa610160602085015101516020615cef82610ef2565b818301019101611b5f565b615d0c61014060208601510151610ef2565b615d1e615d18846118b7565b91610221565b148214615da557615d3c82615d3560208401611c21565b9201611bd4565b615d446101e3565b9081615d4f81611c42565b03925af1906101406020615d61611bf8565b935b93920151015191615da07fc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb39807944693615d976101e3565b93849384611c60565b0390a1565b615db982615db230611bc8565b9201611bd4565b6101406020860151015190602082019151925af1906101406020615ddb611bf8565b93615d63565b615e0d615e08615e03615dfe6008615df846611032565b90610427565b611062565b61107b565b61152f565b906377c2719c90823b15615e8957615e4492615e395f8094615e2d6101e3565b9687958694859361153f565b835260048301611aff565b03925af18015615e8457615e58575b615cc2565b615e77905f3d8111615e7d575b615e6f818361109f565b8101906118d3565b5f615e53565b503d615e65565b611572565b61153b565b615e96612774565b503390565b905f602091615ea8612067565b50615eb1612067565b50828151910182855af115615f4b573d5f5190615ed6615ed05f6118b7565b91610221565b145f14615f315750615ee781612e20565b3b615efa615ef45f6118b7565b91610221565b145b615f035750565b615f0f615f2d91612e20565b615f176101e3565b918291635274afe760e01b8352600483016104a0565b0390fd5b615f44615f3e60016130f4565b91610221565b1415615efc565b6040513d5f823e3d90fdfea2646970667358221220c192ab26307ef1077a72b1edf73dbb7d6982c356afd74f9fc74d635956ba2f5864736f6c63430008180033",
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

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0xebe98c6f.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterCaller) FetchUserOmniMessageFee(opts *bind.CallOpts, params BaseStructCrossMessageParams) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "fetchUserOmniMessageFee", params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0xebe98c6f.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterSession) FetchUserOmniMessageFee(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _SyncRouter.Contract.FetchUserOmniMessageFee(&_SyncRouter.CallOpts, params)
}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0xebe98c6f.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterCallerSession) FetchUserOmniMessageFee(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _SyncRouter.Contract.FetchUserOmniMessageFee(&_SyncRouter.CallOpts, params)
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

// GetUserOmniMessage is a paid mutator transaction binding the contract method 0xff16742b.
//
// Solidity: function getUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) returns(bytes)
func (_SyncRouter *SyncRouterTransactor) GetUserOmniMessage(opts *bind.TransactOpts, params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "getUserOmniMessage", params)
}

// GetUserOmniMessage is a paid mutator transaction binding the contract method 0xff16742b.
//
// Solidity: function getUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) returns(bytes)
func (_SyncRouter *SyncRouterSession) GetUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.Contract.GetUserOmniMessage(&_SyncRouter.TransactOpts, params)
}

// GetUserOmniMessage is a paid mutator transaction binding the contract method 0xff16742b.
//
// Solidity: function getUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) returns(bytes)
func (_SyncRouter *SyncRouterTransactorSession) GetUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.Contract.GetUserOmniMessage(&_SyncRouter.TransactOpts, params)
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

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0xd298a7e9.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_SyncRouter *SyncRouterTransactor) SendUserOmniMessage(opts *bind.TransactOpts, params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "sendUserOmniMessage", params)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0xd298a7e9.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_SyncRouter *SyncRouterSession) SendUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendUserOmniMessage(&_SyncRouter.TransactOpts, params)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0xd298a7e9.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_SyncRouter *SyncRouterTransactorSession) SendUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendUserOmniMessage(&_SyncRouter.TransactOpts, params)
}

// SetLock is a paid mutator transaction binding the contract method 0x6c7cfb52.
//
// Solidity: function setLock(uint256 _way, bytes1 _lockState) returns()
func (_SyncRouter *SyncRouterTransactor) SetLock(opts *bind.TransactOpts, _way *big.Int, _lockState [1]byte) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "setLock", _way, _lockState)
}

// SetLock is a paid mutator transaction binding the contract method 0x6c7cfb52.
//
// Solidity: function setLock(uint256 _way, bytes1 _lockState) returns()
func (_SyncRouter *SyncRouterSession) SetLock(_way *big.Int, _lockState [1]byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.SetLock(&_SyncRouter.TransactOpts, _way, _lockState)
}

// SetLock is a paid mutator transaction binding the contract method 0x6c7cfb52.
//
// Solidity: function setLock(uint256 _way, bytes1 _lockState) returns()
func (_SyncRouter *SyncRouterTransactorSession) SetLock(_way *big.Int, _lockState [1]byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.SetLock(&_SyncRouter.TransactOpts, _way, _lockState)
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
