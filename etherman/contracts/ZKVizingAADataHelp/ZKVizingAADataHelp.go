// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ZKVizingAADataHelp

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

// BaseStructCrossETHParams is an auto generated low-level Go binding around an user-defined struct.
type BaseStructCrossETHParams struct {
	Amount   *big.Int
	Reciever common.Address
}

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

// BaseStructCrossV2SwapParams is an auto generated low-level Go binding around an user-defined struct.
type BaseStructCrossV2SwapParams struct {
	SourceIndex  uint8
	TargetIndex  uint8
	AmountIn     *big.Int
	AmountOutMin *big.Int
	SourceToken  common.Address
	TargetToken  common.Address
	Sender       common.Address
	Receiver     common.Address
	Deadline     *big.Int
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

// BaseStructV2SwapParams is an auto generated low-level Go binding around an user-defined struct.
type BaseStructV2SwapParams struct {
	Index        uint8
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	Sender       common.Address
	Receiver     common.Address
	Deadline     *big.Int
}

// BaseStructV3SwapParams is an auto generated low-level Go binding around an user-defined struct.
type BaseStructV3SwapParams struct {
	Index             uint8
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
	TokenIn           common.Address
	TokenOut          common.Address
	Sender            common.Address
	Receiver          common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
}

// ZKVizingAADataHelpMetaData contains all meta data concerning the ZKVizingAADataHelp contract.
var ZKVizingAADataHelpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"FORK_ID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"callDatas\",\"type\":\"bytes[]\"}],\"name\":\"batchDecodeCrossHookMessageParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"callDatas\",\"type\":\"bytes[]\"}],\"name\":\"batchDecodeCrossMessageParams\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams[]\",\"name\":\"paramsGroup\",\"type\":\"tuple[]\"}],\"name\":\"batchEncodeCrossHookMessageParams\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams[]\",\"name\":\"paramsGroup\",\"type\":\"tuple[]\"}],\"name\":\"batchEncodeCrossMessageParams\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"decodeCrossETHData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reciever\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.CrossETHParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"decodeCrossHookMessageParamsData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"decodeCrossMessageParamsData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"decodePackedUserOperationGroup\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"decodeUniswapV2Data\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structBaseStruct.V2SwapParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"decodeUniswapV3Data\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structBaseStruct.V3SwapParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reciever\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.CrossETHParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"encodeCrossETHParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"bytesCrossETHAmount\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"encodeCrossHookMessageParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"bytesCrossHookMessageParams\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"encodeCrossMessageParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"crossMessage\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"sourceIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"targetIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structBaseStruct.CrossV2SwapParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"encodeCrossV2SwapParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"crossV2MessageBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"sourceIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"targetIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structBaseStruct.CrossV2SwapParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"encodeCrossV3SwapParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"crossV3MessageBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"paramsGroup\",\"type\":\"tuple[]\"}],\"name\":\"encodePackedUserOperationGroup\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"bytesUserOperation\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structBaseStruct.V2SwapParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"encodeV2SwapParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"v2SwapMessageBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structBaseStruct.V3SwapParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"encodeV3SwapParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"v3SwapMessageBytes\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getHasInnerExec\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"}],\"name\":\"getPackChainGasLimit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"}],\"name\":\"getPackChainGasPrice\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"}],\"name\":\"getPackOpInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getPackOperation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"encoded\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"high64\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"low64\",\"type\":\"uint64\"}],\"name\":\"getPackUint64s\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"packed\",\"type\":\"bytes16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"high128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"low128\",\"type\":\"uint256\"}],\"name\":\"getPackUints\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"packed\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523462000023576200001462000029565b6138a46200003482396138a490f35b6200002f565b60405190565b5f80fdfe60806040526004361015610013575b6117ec565b61001d5f356101bc565b806303e4218e146101b75780631398fcdf146101b25780631ddd00d3146101ad5780632198e30c146101a85780632205804a146101a3578063278225a71461019e5780633a7e52f5146101995780633dfab0e41461019457806359295fef1461018f5780635c6432df1461018a57806370e496251461018557806375c9f2df146101805780638892b4fb1461017b57806391dc161a1461017657806392e8fd521461017157806397237cc81461016c5780639f17b72214610167578063a6750a0114610162578063aea79baf1461015d578063af0ecfff14610158578063c27670f814610153578063c65179111461014e578063e851609214610149578063f4bbddad14610144578063fb9f784d1461013f5763fce9e8290361000e576117b6565b6116f9565b6116c3565b611619565b611579565b61142b565b611350565b6112a6565b61123e565b611209565b6111a1565b61113b565b611040565b610e50565b610dd3565b610d16565b610a0b565b6109d6565b61086d565b6106bc565b610687565b610609565b61053a565b610496565b610413565b610365565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561021a5781359167ffffffffffffffff831161021557602001926020830284011161021057565b6101dc565b6101d8565b6101d4565b90602082820312610250575f82013567ffffffffffffffff811161024b5761024792016101e0565b9091565b6101d0565b6101cc565b5190565b60209181520190565b60200190565b5190565b60209181520190565b5f5b838110610287575050905f910152565b806020918301518185015201610277565b601f801991011690565b6102c16102ca6020936102cf936102b881610268565b9384809361026c565b95869101610275565b610298565b0190565b906102dd916102a2565b90565b60200190565b906102fa6102f383610255565b8092610259565b908161030b60208302840194610262565b925f915b83831061031e57505050505090565b9091929394602061034061033a838560019503875289516102d3565b976102e0565b930193019193929061030f565b6103629160208201915f8184039101526102e6565b90565b346103965761039261038161037b36600461021f565b90611b6a565b6103896101c2565b9182918261034d565b0390f35b6101c8565b90565b6103a78161039b565b036103ae57565b5f80fd5b905035906103bf8261039e565b565b91906040838203126103e957806103dd6103e6925f86016103b2565b936020016103b2565b90565b6101cc565b90565b6103fa906103ee565b9052565b9190610411905f602085019401906103f1565b565b346104445761044061042f6104293660046103c1565b90611c19565b6104376101c2565b918291826103fe565b0390f35b6101c8565b5f80fd5b908161010091031261045c5790565b610449565b90602082820312610491575f82013567ffffffffffffffff811161048c57610489920161044d565b90565b6101d0565b6101cc565b346104c6576104c26104b16104ac366004610461565b611c2f565b6104b96101c2565b918291826103fe565b0390f35b6101c8565b908160e09103126104d95790565b610449565b9060208282031261050e575f82013567ffffffffffffffff81116105095761050692016104cb565b90565b6101d0565b6101cc565b151590565b61052190610513565b9052565b9190610538905f60208501940190610518565b565b3461056a576105666105556105503660046104de565b611c48565b61055d6101c2565b91829182610525565b0390f35b6101c8565b9081604091031261057d5790565b610449565b906020828203126105b2575f82013567ffffffffffffffff81116105ad576105aa920161056f565b90565b6101d0565b6101cc565b60209181520190565b6105df6105e86020936105ed936105d681610268565b938480936105b7565b95869101610275565b610298565b0190565b6106069160208201915f8184039101526105c0565b90565b346106395761063561062461061f366004610582565b6120aa565b61062c6101c2565b918291826105f1565b0390f35b6101c8565b908161016091031261064d5790565b610449565b90602082820312610682575f82013567ffffffffffffffff811161067d5761067a920161063e565b90565b6101d0565b6101cc565b346106b7576106b36106a261069d366004610652565b6120e0565b6106aa6101c2565b918291826105f1565b0390f35b6101c8565b346106ec576106e86106d76106d23660046104de565b612116565b6106df6101c2565b918291826103fe565b0390f35b6101c8565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b9061071390610298565b810190811067ffffffffffffffff82111761072d57604052565b6106f5565b9061074561073e6101c2565b9283610709565b565b67ffffffffffffffff811161076557610761602091610298565b0190565b6106f5565b90825f939282370152565b9092919261078a61078582610747565b610732565b938185526020850190828401116107a6576107a49261076a565b565b6106f1565b9080601f830112156107c9578160206107c693359101610775565b90565b6101d4565b906020828203126107fe575f82013567ffffffffffffffff81116107f9576107f692016107ab565b90565b6101d0565b6101cc565b61080c9061039b565b9052565b60018060a01b031690565b61082490610810565b90565b6108309061081b565b9052565b906020806108569361084c5f8201515f860190610803565b0151910190610827565b565b919061086b905f60408501940190610834565b565b3461089d576108996108886108833660046107ce565b6121f4565b6108906101c2565b91829182610858565b0390f35b6101c8565b60ff1690565b6108b1906108a2565b9052565b62ffffff1690565b6108c6906108b5565b9052565b67ffffffffffffffff1690565b6108e0906108ca565b9052565b6109bb916101406109a961016083016109035f8601515f8601906108a8565b610915602086015160208601906108bd565b610927604086015160408601906108d7565b610939606086015160608601906108d7565b61094b608086015160808601906108d7565b61095d60a086015160a08601906108d7565b61096f60c086015160c0860190610827565b61098160e086015160e0860190610827565b610995610100860151610100860190610803565b6101208501518482036101208601526102a2565b920151906101408184039101526102a2565b90565b6109d39160208201915f8184039101526108e4565b90565b34610a0657610a026109f16109ec3660046107ce565b6124a5565b6109f96101c2565b918291826109be565b0390f35b6101c8565b34610a3b57610a37610a26610a21366004610461565b6124cb565b610a2e6101c2565b918291826103fe565b0390f35b6101c8565b5190565b60209181520190565b60200190565b610aeb9160e0610100820192610a6f5f8201515f8501906108d7565b610a81602082015160208501906108d7565b610a93604082015160408501906108d7565b610aa5606082015160608501906108d7565b610ab7608082015160808501906108d7565b610ac960a082015160a08501906108d7565b610adb60c082015160c08501906108d7565b01519060e08184039101526102a2565b90565b610b769160c0610b6560e08301610b0b5f8601515f8601906108a8565b610b1d602086015160208601906108a8565b610b2f60408601516040860190610803565b610b4160608601516060860190610827565b610b5360808601516080860190610827565b60a085015184820360a0860152610a53565b9201519060c0818403910152610a53565b90565b90610b8391610aee565b90565b60200190565b90610ba0610b9983610a40565b8092610a44565b9081610bb160208302840194610a4d565b925f915b838310610bc457505050505090565b90919293946020610be6610be083856001950387528951610b79565b97610b86565b9301930191939290610bb5565b610cca91610140610cb86101608301610c125f8601515f8601906108a8565b610c24602086015160208601906108bd565b610c36604086015160408601906108d7565b610c48606086015160608601906108d7565b610c5a608086015160808601906108d7565b610c6c60a086015160a08601906108d7565b610c7e60c086015160c0860190610827565b610c9060e086015160e0860190610827565b610ca4610100860151610100860190610803565b6101208501518482036101208601526102a2565b920151906101408184039101526102a2565b90565b610cfb916020610cea604083015f8501518482035f860152610b8c565b920151906020818403910152610bf3565b90565b610d139160208201915f818403910152610ccd565b90565b34610d4657610d42610d31610d2c3660046107ce565b612810565b610d396101c2565b91829182610cfe565b0390f35b6101c8565b60209181520190565b90610d68610d6183610a40565b8092610d4b565b9081610d7960208302840194610a4d565b925f915b838310610d8c57505050505090565b90919293946020610dae610da883856001950387528951610b79565b97610b86565b9301930191939290610d7d565b610dd09160208201915f818403910152610d54565b90565b34610e0357610dff610dee610de93660046107ce565b612870565b610df66101c2565b91829182610dbb565b0390f35b6101c8565b908160e0910312610e165790565b610449565b90602082820312610e4b575f82013567ffffffffffffffff8111610e4657610e439201610e08565b90565b6101d0565b6101cc565b34610e8057610e7c610e6b610e66366004610e1b565b612a12565b610e736101c2565b918291826105f1565b0390f35b6101c8565b67ffffffffffffffff8111610e9d5760208091020190565b6106f5565b929190610eb6610eb182610e85565b610732565b9381855260208086019202810191838311610f0d5781905b838210610edc575050505050565b813567ffffffffffffffff8111610f0857602091610efd87849387016107ab565b815201910190610ece565b6101d4565b6101dc565b9080601f83011215610f3057816020610f2d93359101610ea2565b90565b6101d4565b90602082820312610f65575f82013567ffffffffffffffff8111610f6057610f5d9201610f12565b90565b6101d0565b6101cc565b5190565b60209181520190565b60200190565b610fab916020610f9a604083015f8501518482035f860152610b8c565b920151906020818403910152610bf3565b90565b90610fb891610f7d565b90565b60200190565b90610fd5610fce83610f6a565b8092610f6e565b9081610fe660208302840194610f77565b925f915b838310610ff957505050505090565b9091929394602061101b61101583856001950387528951610fae565b97610fbb565b9301930191939290610fea565b61103d9160208201915f818403910152610fc1565b90565b346110705761106c61105b611056366004610f35565b612a6d565b6110636101c2565b91829182611028565b0390f35b6101c8565b61107e90610810565b9052565b90610100806111239361109b5f8201515f8601906108a8565b6110ad602082015160208601906108bd565b6110bf60408201516040860190611075565b6110d160608201516060860190610827565b6110e360808201516080860190610827565b6110f560a082015160a0860190610827565b61110760c082015160c0860190610827565b61111960e082015160e0860190610803565b0151910190610803565b565b9190611139905f6101208501940190611082565b565b3461116b576111676111566111513660046107ce565b612c93565b61115e6101c2565b91829182611125565b0390f35b6101c8565b9081604091031261117e5790565b610449565b9060408282031261119c57611199915f01611170565b90565b6101cc565b346111d1576111cd6111bc6111b7366004611183565b612d03565b6111c46101c2565b918291826105f1565b0390f35b6101c8565b90816101209103126111e55790565b610449565b906101208282031261120457611201915f016111d6565b90565b6101cc565b346112395761123561122461121f3660046111ea565b612e5c565b61122c6101c2565b918291826105f1565b0390f35b6101c8565b3461126e5761126a611259611254366004610461565b612e92565b6112616101c2565b918291826103fe565b0390f35b6101c8565b90816101209103126112825790565b610449565b90610120828203126112a15761129e915f01611273565b90565b6101cc565b346112d6576112d26112c16112bc366004611287565b612fa9565b6112c96101c2565b918291826105f1565b0390f35b6101c8565b909182601f830112156113155781359167ffffffffffffffff831161131057602001926020830284011161130b57565b6101dc565b6101d8565b6101d4565b9060208282031261134b575f82013567ffffffffffffffff81116113465761134292016112db565b9091565b6101d0565b6101cc565b346113815761137d61136c61136636600461131a565b90613063565b6113746101c2565b918291826105f1565b0390f35b6101c8565b5190565b60209181520190565b60200190565b906113a391610bf3565b90565b60200190565b906113c06113b983611386565b809261138a565b90816113d160208302840194611393565b925f915b8383106113e457505050505090565b9091929394602061140661140083856001950387528951611399565b976113a6565b93019301919392906113d5565b6114289160208201915f8184039101526113ac565b90565b3461145b57611457611446611441366004610f35565b6130be565b61144e6101c2565b91829182611413565b0390f35b6101c8565b5190565b60209181520190565b60200190565b9061148081602093610827565b0190565b60200190565b906114a76114a161149a84611460565b8093611464565b9261146d565b905f5b8181106114b75750505090565b9091926114d06114ca6001928651611473565b94611484565b91019190916114aa565b9061155e9060c08061152f60e084016114f95f8801515f8701906108a8565b61150b60208801516020870190610803565b61151d60408801516040870190610803565b6060870151858203606087015261148a565b9461154260808201516080860190610827565b61155460a082015160a0860190610827565b0151910190610803565b90565b6115769160208201915f8184039101526114da565b90565b346115a9576115a561159461158f3660046107ce565b613347565b61159c6101c2565b91829182611561565b0390f35b6101c8565b5f9103126115b857565b6101cc565b90565b90565b6115d76115d26115dc926115bd565b6115c0565b6108ca565b90565b6115e960016115c3565b90565b6115f46115df565b90565b611600906108ca565b9052565b9190611617905f602085019401906115f7565b565b34611649576116293660046115ae565b6116456116346115ec565b61163c6101c2565b91829182611604565b0390f35b6101c8565b909182601f830112156116885781359167ffffffffffffffff831161168357602001926020830284011161167e57565b6101dc565b6101d8565b6101d4565b906020828203126116be575f82013567ffffffffffffffff81116116b9576116b5920161164e565b9091565b6101d0565b6101cc565b346116f4576116f06116df6116d936600461168d565b906133ac565b6116e76101c2565b9182918261034d565b0390f35b6101c8565b346117295761172561171461170f366004611287565b613457565b61171c6101c2565b918291826105f1565b0390f35b6101c8565b611737816108ca565b0361173e57565b5f80fd5b9050359061174f8261172e565b565b9190604083820312611779578061176d611776925f8601611742565b93602001611742565b90565b6101cc565b6fffffffffffffffffffffffffffffffff191690565b61179d9061177e565b9052565b91906117b4905f60208501940190611794565b565b346117e7576117e36117d26117cc366004611751565b90613491565b6117da6101c2565b918291826117a1565b0390f35b6101c8565b5f80fd5b606090565b5090565b9061180b61180683610e85565b610732565b918252565b606090565b5f5b82811061182357505050565b60209061182e611810565b8184015201611817565b9061185d611845836117f9565b926020806118538693610e85565b9201910390611815565b565b5f90565b600161186f910161039b565b90565b634e487b7160e01b5f52603260045260245ffd5b5f80fd5b903590600161016003813603038212156118a2570190565b611886565b908210156118c15760206118be920281019061188a565b90565b611872565b6118cf816108a2565b036118d657565b5f80fd5b905035906118e7826118c6565b565b506118f89060208101906118da565b90565b611904816108b5565b0361190b57565b5f80fd5b9050359061191c826118fb565b565b5061192d90602081019061190f565b90565b5061193f906020810190611742565b90565b61194b8161081b565b0361195257565b5f80fd5b9050359061196382611942565b565b50611974906020810190611956565b90565b506119869060208101906103b2565b90565b5f80fd5b5f80fd5b5f80fd5b90356001602003823603038112156119d657016020813591019167ffffffffffffffff82116119d15760018202360383136119cc57565b61198d565b611989565b611991565b91906119f5816119ee816119fa9561026c565b809561076a565b610298565b0190565b611b2f91611b20611b146101608301611a25611a1c5f8701876118e9565b5f8601906108a8565b611a3f611a35602087018761191e565b60208601906108bd565b611a59611a4f6040870187611930565b60408601906108d7565b611a73611a696060870187611930565b60608601906108d7565b611a8d611a836080870187611930565b60808601906108d7565b611aa7611a9d60a0870187611930565b60a08601906108d7565b611ac1611ab760c0870187611965565b60c0860190610827565b611adb611ad160e0870187611965565b60e0860190610827565b611af7611aec610100870187611977565b610100860190610803565b611b05610120860186611995565b908583036101208701526119db565b92610140810190611995565b916101408185039101526119db565b90565b611b479160208201915f8184039101526119fe565b90565b90611b5482610255565b811015611b65576020809102010190565b611872565b90611b736117f0565b50611b87611b828383906117f5565b611838565b91611b9061185f565b5b80611bae611ba8611ba38587906117f5565b61039b565b9161039b565b1015611c0f57611c0a90611c03611be1611bf0611bcd868886916118a7565b611bd56101c2565b92839160208301611b32565b60208201810382520382610709565b868391611bfd8383611b4a565b52611b4a565b5150611863565b611b91565b50505090565b5f90565b90611c2c91611c26611c15565b5061354e565b90565b611c4190611c3b611c15565b506135fb565b90565b5f90565b611c5a90611c54611c44565b50613677565b90565b606090565b9035600160200382360303811215611ca357016020813591019167ffffffffffffffff8211611c9e576020820236038313611c9957565b61198d565b611989565b611991565b90565b903560016101000382360303811215611cc2570190565b611991565b611da091611d92610100820192611cec611ce35f830183611930565b5f8501906108d7565b611d06611cfc6020830183611930565b60208501906108d7565b611d20611d166040830183611930565b60408501906108d7565b611d3a611d306060830183611930565b60608501906108d7565b611d54611d4a6080830183611930565b60808501906108d7565b611d6e611d6460a0830183611930565b60a08501906108d7565b611d88611d7e60c0830183611930565b60c08501906108d7565b60e0810190611995565b9160e08185039101526119db565b90565b611e6491611e56611e4b60e08301611dc9611dc05f8701876118e9565b5f8601906108a8565b611de3611dd960208701876118e9565b60208601906108a8565b611dfd611df36040870187611977565b6040860190610803565b611e17611e0d6060870187611965565b6060860190610827565b611e31611e276080870187611965565b6080860190610827565b611e3e60a0860186611cab565b84820360a0860152611cc7565b9260c0810190611cab565b9060c0818403910152611cc7565b90565b90611e7191611da3565b90565b9035600160e00382360303811215611e8a570190565b611991565b60200190565b9181611ea091610a44565b9081611eb160208302840194611ca8565b92835f925b848410611ec65750505050505090565b9091929394956020611ef1611eeb8385600195038852611ee68b88611e74565b611e67565b98611e8f565b940194019294939190611eb6565b903560016101600382360303811215611f16570190565b611991565b61204c9161203d6120316101608301611f42611f395f8701876118e9565b5f8601906108a8565b611f5c611f52602087018761191e565b60208601906108bd565b611f76611f6c6040870187611930565b60408601906108d7565b611f90611f866060870187611930565b60608601906108d7565b611faa611fa06080870187611930565b60808601906108d7565b611fc4611fba60a0870187611930565b60a08601906108d7565b611fde611fd460c0870187611965565b60c0860190610827565b611ff8611fee60e0870187611965565b60e0860190610827565b612014612009610100870187611977565b610100860190610803565b612022610120860186611995565b908583036101208701526119db565b92610140810190611995565b916101408185039101526119db565b90565b61208f91612081612076604083016120695f860186611c62565b908583035f870152611e95565b926020810190611eff565b906020818403910152611f1b565b90565b6120a79160208201915f81840391015261204f565b90565b6120dd6120ce916120b9611c5d565b506120c26101c2565b92839160208301612092565b60208201810382520382610709565b90565b612113612104916120ef611c5d565b506120f86101c2565b92839160208301611b32565b60208201810382520382610709565b90565b61212890612122611c15565b506136ff565b90565b6121356040610732565b90565b5f90565b5f90565b61214861212b565b9060208083612155612138565b81520161216061213c565b81525050565b61216e612140565b90565b5f80fd5b5f80fd5b905051906121868261039e565b565b9050519061219582611942565b565b91906040838203126121d1576121ca906121b16040610732565b936121be825f8301612179565b5f860152602001612188565b6020830152565b612171565b906040828203126121ef576121ec915f01612197565b90565b6101cc565b61221790612200612166565b50602061220c82610268565b8183010191016121d6565b90565b612225610160610732565b90565b5f90565b5f90565b5f90565b61223c61221a565b906020808080808080808080808c612252612228565b81520161225d61222c565b815201612268612230565b815201612273612230565b81520161227e612230565b815201612289612230565b81520161229461213c565b81520161229f61213c565b8152016122aa612138565b8152016122b5611810565b8152016122c0611810565b81525050565b6122ce612234565b90565b905051906122de826118c6565b565b905051906122ed826118fb565b565b905051906122fc8261172e565b565b9092919261231361230e82610747565b610732565b9381855260208501908284011161232f5761232d92610275565b565b6106f1565b9080601f830112156123525781602061234f935191016122fe565b90565b6101d4565b9190916101608184031261246b57612370610160610732565b9261237d815f84016122d1565b5f85015261238e81602084016122e0565b60208501526123a081604084016122ef565b60408501526123b281606084016122ef565b60608501526123c481608084016122ef565b60808501526123d68160a084016122ef565b60a08501526123e88160c08401612188565b60c08501526123fa8160e08401612188565b60e085015261240d816101008401612179565b61010085015261012082015167ffffffffffffffff81116124665781612434918401612334565b61012085015261014082015167ffffffffffffffff8111612461576124599201612334565b610140830152565b612175565b612175565b612171565b906020828203126124a0575f82015167ffffffffffffffff811161249b576124989201612357565b90565b6101d0565b6101cc565b6124c8906124b16122c6565b5060206124bd82610268565b818301019101612470565b90565b6124dd906124d7611c15565b50613730565b90565b6124ea6040610732565b90565b606090565b6124fa612234565b90565b6125056124e0565b90602080836125126124ed565b81520161251d6124f2565b81525050565b61252b6124fd565b90565b67ffffffffffffffff81116125465760208091020190565b6106f5565b9190916101008184031261260b57612564610100610732565b92612571815f84016122ef565b5f85015261258281602084016122ef565b602085015261259481604084016122ef565b60408501526125a681606084016122ef565b60608501526125b881608084016122ef565b60808501526125ca8160a084016122ef565b60a08501526125dc8160c084016122ef565b60c085015260e082015167ffffffffffffffff8111612606576125ff9201612334565b60e0830152565b612175565b612171565b91909160e0818403126126d45761262760e0610732565b92612634815f84016122d1565b5f85015261264581602084016122d1565b60208501526126578160408401612179565b60408501526126698160608401612188565b606085015261267b8160808401612188565b608085015260a082015167ffffffffffffffff81116126cf57816126a091840161254b565b60a085015260c082015167ffffffffffffffff81116126ca576126c3920161254b565b60c0830152565b612175565b612175565b612171565b9291906126ed6126e88261252e565b610732565b93818552602080860192028101918383116127445781905b838210612713575050505050565b815167ffffffffffffffff811161273f576020916127348784938701612610565b815201910190612705565b6101d4565b6101dc565b9080601f8301121561276757816020612764935191016126d9565b90565b6101d4565b9190916040818403126127d6576127836040610732565b925f82015167ffffffffffffffff81116127d157816127a3918401612749565b5f850152602082015167ffffffffffffffff81116127cc576127c59201612357565b6020830152565b612175565b612175565b612171565b9060208282031261280b575f82015167ffffffffffffffff811161280657612803920161276c565b90565b6101d0565b6101cc565b6128339061281c612523565b50602061282882610268565b8183010191016127db565b90565b606090565b9060208282031261286b575f82015167ffffffffffffffff8111612866576128639201612749565b90565b6101d0565b6101cc565b6128939061287c612836565b50602061288882610268565b81830101910161283b565b90565b90356001602003823603038112156128d757016020813591019167ffffffffffffffff82116128d25760208202360383136128cd57565b61198d565b611989565b611991565b90565b60200190565b916128f3826128f992611464565b926128dc565b90815f905b82821061290c575050505090565b9091929361292e6129286001926129238886611965565b611473565b956128df565b9201909291926128fe565b906129f79060c06129ef6129b160e084016129626129595f8901896118e9565b5f8701906108a8565b61297c6129726020890189611977565b6020870190610803565b61299661298c6040890189611977565b6040870190610803565b6129a36060880188612896565b9086830360608801526128e5565b946129cc6129c26080830183611965565b6080860190610827565b6129e66129dc60a0830183611965565b60a0860190610827565b82810190611977565b910190610803565b90565b612a0f9160208201915f818403910152612939565b90565b612a45612a3691612a21611c5d565b50612a2a6101c2565b928391602083016129fa565b60208201810382520382610709565b90565b606090565b90612a5782610f6a565b811015612a68576020809102010190565b611872565b612a75612a48565b50612a7e612a48565b90612a8761185f565b5b80612aa3612a9d612a9885610255565b61039b565b9161039b565b1015612af457612aef90612ae8612ad5612abe858490611b4a565b516020612aca82610268565b8183010191016127db565b858391612ae28383612a4d565b52612a4d565b5150611863565b612a88565b505090565b612b04610120610732565b90565b5f90565b612b13612af9565b90602080808080808080808a612b27612228565b815201612b3261222c565b815201612b3d612b07565b815201612b4861213c565b815201612b5361213c565b815201612b5e61213c565b815201612b6961213c565b815201612b74612138565b815201612b7f612138565b81525050565b612b8d612b0b565b90565b612b9981610810565b03612ba057565b5f80fd5b90505190612bb182612b90565b565b919061012083820312612c6f57612c6790612bcf610120610732565b93612bdc825f83016122d1565b5f860152612bed82602083016122e0565b6020860152612bff8260408301612ba4565b6040860152612c118260608301612188565b6060860152612c238260808301612188565b6080860152612c358260a08301612188565b60a0860152612c478260c08301612188565b60c0860152612c598260e08301612179565b60e086015261010001612179565b610100830152565b612171565b9061012082820312612c8e57612c8b915f01612bb3565b90565b6101cc565b612cb690612c9f612b85565b506020612cab82610268565b818301019101612c74565b90565b906020612ce4612cec93612cdb612cd25f830183611977565b5f860190610803565b82810190611965565b910190610827565b565b9190612d01905f60408501940190612cb9565b565b612d36612d2791612d12611c5d565b50612d1b6101c2565b92839160208301612cee565b60208201810382520382610709565b90565b90503590612d4682612b90565b565b50612d57906020810190612d39565b90565b90610100612e3c612e4493612d7d612d745f8301836118e9565b5f8601906108a8565b612d97612d8d602083018361191e565b60208601906108bd565b612db1612da76040830183612d48565b6040860190611075565b612dcb612dc16060830183611965565b6060860190610827565b612de5612ddb6080830183611965565b6080860190610827565b612dff612df560a0830183611965565b60a0860190610827565b612e19612e0f60c0830183611965565b60c0860190610827565b612e33612e2960e0830183611977565b60e0860190610803565b82810190611977565b910190610803565b565b9190612e5a905f6101208501940190612d5a565b565b612e8f612e8091612e6b611c5d565b50612e746101c2565b92839160208301612e46565b60208201810382520382610709565b90565b612ea490612e9e611c15565b5061376f565b90565b90610100612f89612f9193612eca612ec15f8301836118e9565b5f8601906108a8565b612ee4612eda60208301836118e9565b60208601906108a8565b612efe612ef46040830183611977565b6040860190610803565b612f18612f0e6060830183611977565b6060860190610803565b612f32612f286080830183611965565b6080860190610827565b612f4c612f4260a0830183611965565b60a0860190610827565b612f66612f5c60c0830183611965565b60c0860190610827565b612f80612f7660e0830183611965565b60e0860190610827565b82810190611977565b910190610803565b565b9190612fa7905f6101208501940190612ea7565b565b612fdc612fcd91612fb8611c5d565b50612fc16101c2565b92839160208301612f93565b60208201810382520382610709565b90565b9181612fea91610d4b565b9081612ffb60208302840194611ca8565b92835f925b8484106130105750505050505090565b909192939495602061303b61303583856001950388526130308b88611e74565b611e67565b98611e8f565b940194019294939190613000565b90916130609260208301925f818503910152612fdf565b90565b6130969061306f611c5d565b5061308761307b6101c2565b93849260208401613049565b60208201810382520382610709565b90565b606090565b906130a882611386565b8110156130b9576020809102010190565b611872565b6130c6613099565b506130cf613099565b906130d861185f565b5b806130f46130ee6130e985610255565b61039b565b9161039b565b1015613145576131409061313961312661310f858490611b4a565b51602061311b82610268565b818301019101612470565b858391613133838361309e565b5261309e565b5150611863565b6130d9565b505090565b61315460e0610732565b90565b606090565b61316461314a565b90602080808080808088613176612228565b815201613181612138565b81520161318c612138565b815201613197613157565b8152016131a261213c565b8152016131ad61213c565b8152016131b8612138565b81525050565b6131c661315c565b90565b67ffffffffffffffff81116131e15760208091020190565b6106f5565b909291926131fb6131f6826131c9565b610732565b938185526020808601920283019281841161323857915b83831061321f5750505050565b6020809161322d8486612188565b815201920191613212565b6101dc565b9080601f8301121561325b57816020613258935191016131e6565b90565b6101d4565b91909160e08184031261330d5761327760e0610732565b92613284815f84016122d1565b5f8501526132958160208401612179565b60208501526132a78160408401612179565b604085015260608201519167ffffffffffffffff8311613308576132d08261330194830161323d565b60608601526132e28260808301612188565b60808601526132f48260a08301612188565b60a086015260c001612179565b60c0830152565b612175565b612171565b90602082820312613342575f82015167ffffffffffffffff811161333d5761333a9201613260565b90565b6101d0565b6101cc565b61336a906133536131be565b50602061335f82610268565b818301019101613312565b90565b5090565b903590600160400381360303821215613388570190565b611886565b908210156133a75760206133a49202810190613371565b90565b611872565b906133b56117f0565b506133c96133c483839061336d565b611838565b916133d261185f565b5b806133f06133ea6133e585879061336d565b61039b565b9161039b565b10156134515761344c9061344561342361343261340f8688869161338d565b6134176101c2565b92839160208301612092565b60208201810382520382610709565b86839161343f8383611b4a565b52611b4a565b5150611863565b6133d3565b50505090565b61348a61347b91613466611c5d565b5061346f6101c2565b92839160208301612f93565b60208201810382520382610709565b90565b5f90565b906134a49161349e61348d565b50613833565b90565b6fffffffffffffffffffffffffffffffff1690565b6134d06134cb6134d5926134a7565b6115c0565b61039b565b90565b156134df57565b5f80fd5b90565b6134fa6134f56134ff926134e3565b6115c0565b6108a2565b90565b1b90565b6135259061351f61351961352a946108a2565b9161039b565b90613502565b61039b565b90565b5f1b90565b61354661354161354b9261039b565b61352d565b6103ee565b90565b906135c96135cf9261355e611c15565b5061358c8161358561357f6fffffffffffffffffffffffffffffffff6134bc565b9161039b565b11156134d8565b6135b9836135b26135ac6fffffffffffffffffffffffffffffffff6134bc565b9161039b565b11156134d8565b6135c360806134e6565b90613506565b17613532565b90565b356135dc8161172e565b90565b6135f36135ee6135f8926108ca565b6115c0565b61039b565b90565b61363890613607611c15565b5061363261362c613626606061361f604086016135d2565b94016135d2565b926135df565b916135df565b9061354e565b90565b90359060016101000381360303821215613653570190565b611886565b90565b61366f61366a61367492613658565b6115c0565b6108ca565b90565b602061369361369992613688611c44565b5060c081019061363b565b016135d2565b6136ab6136a55f61365b565b916108ca565b145f146136b7575f5b90565b60016136b4565b356136c8816118c6565b90565b356136d58161039e565b90565b60018060f81b031690565b6136f76136f26136fc9261039b565b6115c0565b6136d8565b90565b613707611c15565b50613728613723604061371c602085016136be565b93016136cb565b6136e3565b9060f81b1790565b61376c9061373c611c15565b5061376661376061375a60206137535f86016135d2565b94016135d2565b926135df565b916135df565b9061354e565b90565b6137ac9061377b611c15565b506137a66137a061379a60c061379360a086016135d2565b94016135d2565b926135df565b916135df565b9061354e565b90565b6137c36137be6137c8926108ca565b6115c0565b6134a7565b90565b90565b6137e26137dd6137e7926137cb565b6115c0565b6108a2565b90565b613809906138036137fd61380e946108a2565b916134a7565b90613502565b6134a7565b90565b60801b90565b61382b613826613830926134a7565b613811565b61177e565b90565b9061386561385f61384f61386b9461384961348d565b506137af565b61385960406137ce565b906137ea565b916137af565b17613817565b9056fea2646970667358221220c9fc1098ce4518c07359f75a2c50b73635acc1df32c1b2ac4bdef31488af16dc64736f6c63430008180033",
}

// ZKVizingAADataHelpABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKVizingAADataHelpMetaData.ABI instead.
var ZKVizingAADataHelpABI = ZKVizingAADataHelpMetaData.ABI

// ZKVizingAADataHelpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKVizingAADataHelpMetaData.Bin instead.
var ZKVizingAADataHelpBin = ZKVizingAADataHelpMetaData.Bin

// DeployZKVizingAADataHelp deploys a new Ethereum contract, binding an instance of ZKVizingAADataHelp to it.
func DeployZKVizingAADataHelp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZKVizingAADataHelp, error) {
	parsed, err := ZKVizingAADataHelpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKVizingAADataHelpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZKVizingAADataHelp{ZKVizingAADataHelpCaller: ZKVizingAADataHelpCaller{contract: contract}, ZKVizingAADataHelpTransactor: ZKVizingAADataHelpTransactor{contract: contract}, ZKVizingAADataHelpFilterer: ZKVizingAADataHelpFilterer{contract: contract}}, nil
}

// ZKVizingAADataHelp is an auto generated Go binding around an Ethereum contract.
type ZKVizingAADataHelp struct {
	ZKVizingAADataHelpCaller     // Read-only binding to the contract
	ZKVizingAADataHelpTransactor // Write-only binding to the contract
	ZKVizingAADataHelpFilterer   // Log filterer for contract events
}

// ZKVizingAADataHelpCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKVizingAADataHelpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKVizingAADataHelpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKVizingAADataHelpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKVizingAADataHelpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKVizingAADataHelpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKVizingAADataHelpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKVizingAADataHelpSession struct {
	Contract     *ZKVizingAADataHelp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZKVizingAADataHelpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKVizingAADataHelpCallerSession struct {
	Contract *ZKVizingAADataHelpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ZKVizingAADataHelpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKVizingAADataHelpTransactorSession struct {
	Contract     *ZKVizingAADataHelpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ZKVizingAADataHelpRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKVizingAADataHelpRaw struct {
	Contract *ZKVizingAADataHelp // Generic contract binding to access the raw methods on
}

// ZKVizingAADataHelpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKVizingAADataHelpCallerRaw struct {
	Contract *ZKVizingAADataHelpCaller // Generic read-only contract binding to access the raw methods on
}

// ZKVizingAADataHelpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKVizingAADataHelpTransactorRaw struct {
	Contract *ZKVizingAADataHelpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKVizingAADataHelp creates a new instance of ZKVizingAADataHelp, bound to a specific deployed contract.
func NewZKVizingAADataHelp(address common.Address, backend bind.ContractBackend) (*ZKVizingAADataHelp, error) {
	contract, err := bindZKVizingAADataHelp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKVizingAADataHelp{ZKVizingAADataHelpCaller: ZKVizingAADataHelpCaller{contract: contract}, ZKVizingAADataHelpTransactor: ZKVizingAADataHelpTransactor{contract: contract}, ZKVizingAADataHelpFilterer: ZKVizingAADataHelpFilterer{contract: contract}}, nil
}

// NewZKVizingAADataHelpCaller creates a new read-only instance of ZKVizingAADataHelp, bound to a specific deployed contract.
func NewZKVizingAADataHelpCaller(address common.Address, caller bind.ContractCaller) (*ZKVizingAADataHelpCaller, error) {
	contract, err := bindZKVizingAADataHelp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKVizingAADataHelpCaller{contract: contract}, nil
}

// NewZKVizingAADataHelpTransactor creates a new write-only instance of ZKVizingAADataHelp, bound to a specific deployed contract.
func NewZKVizingAADataHelpTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKVizingAADataHelpTransactor, error) {
	contract, err := bindZKVizingAADataHelp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKVizingAADataHelpTransactor{contract: contract}, nil
}

// NewZKVizingAADataHelpFilterer creates a new log filterer instance of ZKVizingAADataHelp, bound to a specific deployed contract.
func NewZKVizingAADataHelpFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKVizingAADataHelpFilterer, error) {
	contract, err := bindZKVizingAADataHelp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKVizingAADataHelpFilterer{contract: contract}, nil
}

// bindZKVizingAADataHelp binds a generic wrapper to an already deployed contract.
func bindZKVizingAADataHelp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKVizingAADataHelpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKVizingAADataHelp *ZKVizingAADataHelpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKVizingAADataHelp.Contract.ZKVizingAADataHelpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKVizingAADataHelp *ZKVizingAADataHelpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKVizingAADataHelp.Contract.ZKVizingAADataHelpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKVizingAADataHelp *ZKVizingAADataHelpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKVizingAADataHelp.Contract.ZKVizingAADataHelpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKVizingAADataHelp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKVizingAADataHelp *ZKVizingAADataHelpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKVizingAADataHelp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKVizingAADataHelp *ZKVizingAADataHelpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKVizingAADataHelp.Contract.contract.Transact(opts, method, params...)
}

// FORKID is a free data retrieval call binding the contract method 0xe8516092.
//
// Solidity: function FORK_ID() view returns(uint64)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) FORKID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "FORK_ID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FORKID is a free data retrieval call binding the contract method 0xe8516092.
//
// Solidity: function FORK_ID() view returns(uint64)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) FORKID() (uint64, error) {
	return _ZKVizingAADataHelp.Contract.FORKID(&_ZKVizingAADataHelp.CallOpts)
}

// FORKID is a free data retrieval call binding the contract method 0xe8516092.
//
// Solidity: function FORK_ID() view returns(uint64)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) FORKID() (uint64, error) {
	return _ZKVizingAADataHelp.Contract.FORKID(&_ZKVizingAADataHelp.CallOpts)
}

// BatchDecodeCrossHookMessageParams is a free data retrieval call binding the contract method 0xc27670f8.
//
// Solidity: function batchDecodeCrossHookMessageParams(bytes[] callDatas) pure returns((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) BatchDecodeCrossHookMessageParams(opts *bind.CallOpts, callDatas [][]byte) ([]BaseStructCrossHookMessageParams, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "batchDecodeCrossHookMessageParams", callDatas)

	if err != nil {
		return *new([]BaseStructCrossHookMessageParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]BaseStructCrossHookMessageParams)).(*[]BaseStructCrossHookMessageParams)

	return out0, err

}

// BatchDecodeCrossHookMessageParams is a free data retrieval call binding the contract method 0xc27670f8.
//
// Solidity: function batchDecodeCrossHookMessageParams(bytes[] callDatas) pure returns((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) BatchDecodeCrossHookMessageParams(callDatas [][]byte) ([]BaseStructCrossHookMessageParams, error) {
	return _ZKVizingAADataHelp.Contract.BatchDecodeCrossHookMessageParams(&_ZKVizingAADataHelp.CallOpts, callDatas)
}

// BatchDecodeCrossHookMessageParams is a free data retrieval call binding the contract method 0xc27670f8.
//
// Solidity: function batchDecodeCrossHookMessageParams(bytes[] callDatas) pure returns((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) BatchDecodeCrossHookMessageParams(callDatas [][]byte) ([]BaseStructCrossHookMessageParams, error) {
	return _ZKVizingAADataHelp.Contract.BatchDecodeCrossHookMessageParams(&_ZKVizingAADataHelp.CallOpts, callDatas)
}

// BatchDecodeCrossMessageParams is a free data retrieval call binding the contract method 0x91dc161a.
//
// Solidity: function batchDecodeCrossMessageParams(bytes[] callDatas) pure returns(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes))[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) BatchDecodeCrossMessageParams(opts *bind.CallOpts, callDatas [][]byte) ([]BaseStructCrossMessageParams, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "batchDecodeCrossMessageParams", callDatas)

	if err != nil {
		return *new([]BaseStructCrossMessageParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]BaseStructCrossMessageParams)).(*[]BaseStructCrossMessageParams)

	return out0, err

}

// BatchDecodeCrossMessageParams is a free data retrieval call binding the contract method 0x91dc161a.
//
// Solidity: function batchDecodeCrossMessageParams(bytes[] callDatas) pure returns(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes))[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) BatchDecodeCrossMessageParams(callDatas [][]byte) ([]BaseStructCrossMessageParams, error) {
	return _ZKVizingAADataHelp.Contract.BatchDecodeCrossMessageParams(&_ZKVizingAADataHelp.CallOpts, callDatas)
}

// BatchDecodeCrossMessageParams is a free data retrieval call binding the contract method 0x91dc161a.
//
// Solidity: function batchDecodeCrossMessageParams(bytes[] callDatas) pure returns(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes))[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) BatchDecodeCrossMessageParams(callDatas [][]byte) ([]BaseStructCrossMessageParams, error) {
	return _ZKVizingAADataHelp.Contract.BatchDecodeCrossMessageParams(&_ZKVizingAADataHelp.CallOpts, callDatas)
}

// BatchEncodeCrossHookMessageParams is a free data retrieval call binding the contract method 0x03e4218e.
//
// Solidity: function batchEncodeCrossHookMessageParams((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)[] paramsGroup) pure returns(bytes[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) BatchEncodeCrossHookMessageParams(opts *bind.CallOpts, paramsGroup []BaseStructCrossHookMessageParams) ([][]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "batchEncodeCrossHookMessageParams", paramsGroup)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// BatchEncodeCrossHookMessageParams is a free data retrieval call binding the contract method 0x03e4218e.
//
// Solidity: function batchEncodeCrossHookMessageParams((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)[] paramsGroup) pure returns(bytes[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) BatchEncodeCrossHookMessageParams(paramsGroup []BaseStructCrossHookMessageParams) ([][]byte, error) {
	return _ZKVizingAADataHelp.Contract.BatchEncodeCrossHookMessageParams(&_ZKVizingAADataHelp.CallOpts, paramsGroup)
}

// BatchEncodeCrossHookMessageParams is a free data retrieval call binding the contract method 0x03e4218e.
//
// Solidity: function batchEncodeCrossHookMessageParams((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)[] paramsGroup) pure returns(bytes[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) BatchEncodeCrossHookMessageParams(paramsGroup []BaseStructCrossHookMessageParams) ([][]byte, error) {
	return _ZKVizingAADataHelp.Contract.BatchEncodeCrossHookMessageParams(&_ZKVizingAADataHelp.CallOpts, paramsGroup)
}

// BatchEncodeCrossMessageParams is a free data retrieval call binding the contract method 0xf4bbddad.
//
// Solidity: function batchEncodeCrossMessageParams(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes))[] paramsGroup) pure returns(bytes[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) BatchEncodeCrossMessageParams(opts *bind.CallOpts, paramsGroup []BaseStructCrossMessageParams) ([][]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "batchEncodeCrossMessageParams", paramsGroup)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// BatchEncodeCrossMessageParams is a free data retrieval call binding the contract method 0xf4bbddad.
//
// Solidity: function batchEncodeCrossMessageParams(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes))[] paramsGroup) pure returns(bytes[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) BatchEncodeCrossMessageParams(paramsGroup []BaseStructCrossMessageParams) ([][]byte, error) {
	return _ZKVizingAADataHelp.Contract.BatchEncodeCrossMessageParams(&_ZKVizingAADataHelp.CallOpts, paramsGroup)
}

// BatchEncodeCrossMessageParams is a free data retrieval call binding the contract method 0xf4bbddad.
//
// Solidity: function batchEncodeCrossMessageParams(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes))[] paramsGroup) pure returns(bytes[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) BatchEncodeCrossMessageParams(paramsGroup []BaseStructCrossMessageParams) ([][]byte, error) {
	return _ZKVizingAADataHelp.Contract.BatchEncodeCrossMessageParams(&_ZKVizingAADataHelp.CallOpts, paramsGroup)
}

// DecodeCrossETHData is a free data retrieval call binding the contract method 0x3dfab0e4.
//
// Solidity: function decodeCrossETHData(bytes callData) pure returns((uint256,address))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) DecodeCrossETHData(opts *bind.CallOpts, callData []byte) (BaseStructCrossETHParams, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "decodeCrossETHData", callData)

	if err != nil {
		return *new(BaseStructCrossETHParams), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseStructCrossETHParams)).(*BaseStructCrossETHParams)

	return out0, err

}

// DecodeCrossETHData is a free data retrieval call binding the contract method 0x3dfab0e4.
//
// Solidity: function decodeCrossETHData(bytes callData) pure returns((uint256,address))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) DecodeCrossETHData(callData []byte) (BaseStructCrossETHParams, error) {
	return _ZKVizingAADataHelp.Contract.DecodeCrossETHData(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodeCrossETHData is a free data retrieval call binding the contract method 0x3dfab0e4.
//
// Solidity: function decodeCrossETHData(bytes callData) pure returns((uint256,address))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) DecodeCrossETHData(callData []byte) (BaseStructCrossETHParams, error) {
	return _ZKVizingAADataHelp.Contract.DecodeCrossETHData(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodeCrossHookMessageParamsData is a free data retrieval call binding the contract method 0x59295fef.
//
// Solidity: function decodeCrossHookMessageParamsData(bytes callData) pure returns((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) DecodeCrossHookMessageParamsData(opts *bind.CallOpts, callData []byte) (BaseStructCrossHookMessageParams, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "decodeCrossHookMessageParamsData", callData)

	if err != nil {
		return *new(BaseStructCrossHookMessageParams), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseStructCrossHookMessageParams)).(*BaseStructCrossHookMessageParams)

	return out0, err

}

// DecodeCrossHookMessageParamsData is a free data retrieval call binding the contract method 0x59295fef.
//
// Solidity: function decodeCrossHookMessageParamsData(bytes callData) pure returns((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) DecodeCrossHookMessageParamsData(callData []byte) (BaseStructCrossHookMessageParams, error) {
	return _ZKVizingAADataHelp.Contract.DecodeCrossHookMessageParamsData(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodeCrossHookMessageParamsData is a free data retrieval call binding the contract method 0x59295fef.
//
// Solidity: function decodeCrossHookMessageParamsData(bytes callData) pure returns((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) DecodeCrossHookMessageParamsData(callData []byte) (BaseStructCrossHookMessageParams, error) {
	return _ZKVizingAADataHelp.Contract.DecodeCrossHookMessageParamsData(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodeCrossMessageParamsData is a free data retrieval call binding the contract method 0x70e49625.
//
// Solidity: function decodeCrossMessageParamsData(bytes callData) pure returns(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) DecodeCrossMessageParamsData(opts *bind.CallOpts, callData []byte) (BaseStructCrossMessageParams, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "decodeCrossMessageParamsData", callData)

	if err != nil {
		return *new(BaseStructCrossMessageParams), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseStructCrossMessageParams)).(*BaseStructCrossMessageParams)

	return out0, err

}

// DecodeCrossMessageParamsData is a free data retrieval call binding the contract method 0x70e49625.
//
// Solidity: function decodeCrossMessageParamsData(bytes callData) pure returns(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) DecodeCrossMessageParamsData(callData []byte) (BaseStructCrossMessageParams, error) {
	return _ZKVizingAADataHelp.Contract.DecodeCrossMessageParamsData(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodeCrossMessageParamsData is a free data retrieval call binding the contract method 0x70e49625.
//
// Solidity: function decodeCrossMessageParamsData(bytes callData) pure returns(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) DecodeCrossMessageParamsData(callData []byte) (BaseStructCrossMessageParams, error) {
	return _ZKVizingAADataHelp.Contract.DecodeCrossMessageParamsData(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodePackedUserOperationGroup is a free data retrieval call binding the contract method 0x75c9f2df.
//
// Solidity: function decodePackedUserOperationGroup(bytes callData) pure returns((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) DecodePackedUserOperationGroup(opts *bind.CallOpts, callData []byte) ([]BaseStructPackedUserOperation, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "decodePackedUserOperationGroup", callData)

	if err != nil {
		return *new([]BaseStructPackedUserOperation), err
	}

	out0 := *abi.ConvertType(out[0], new([]BaseStructPackedUserOperation)).(*[]BaseStructPackedUserOperation)

	return out0, err

}

// DecodePackedUserOperationGroup is a free data retrieval call binding the contract method 0x75c9f2df.
//
// Solidity: function decodePackedUserOperationGroup(bytes callData) pure returns((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) DecodePackedUserOperationGroup(callData []byte) ([]BaseStructPackedUserOperation, error) {
	return _ZKVizingAADataHelp.Contract.DecodePackedUserOperationGroup(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodePackedUserOperationGroup is a free data retrieval call binding the contract method 0x75c9f2df.
//
// Solidity: function decodePackedUserOperationGroup(bytes callData) pure returns((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[])
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) DecodePackedUserOperationGroup(callData []byte) ([]BaseStructPackedUserOperation, error) {
	return _ZKVizingAADataHelp.Contract.DecodePackedUserOperationGroup(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodeUniswapV2Data is a free data retrieval call binding the contract method 0xc6517911.
//
// Solidity: function decodeUniswapV2Data(bytes callData) pure returns((uint8,uint256,uint256,address[],address,address,uint256))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) DecodeUniswapV2Data(opts *bind.CallOpts, callData []byte) (BaseStructV2SwapParams, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "decodeUniswapV2Data", callData)

	if err != nil {
		return *new(BaseStructV2SwapParams), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseStructV2SwapParams)).(*BaseStructV2SwapParams)

	return out0, err

}

// DecodeUniswapV2Data is a free data retrieval call binding the contract method 0xc6517911.
//
// Solidity: function decodeUniswapV2Data(bytes callData) pure returns((uint8,uint256,uint256,address[],address,address,uint256))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) DecodeUniswapV2Data(callData []byte) (BaseStructV2SwapParams, error) {
	return _ZKVizingAADataHelp.Contract.DecodeUniswapV2Data(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodeUniswapV2Data is a free data retrieval call binding the contract method 0xc6517911.
//
// Solidity: function decodeUniswapV2Data(bytes callData) pure returns((uint8,uint256,uint256,address[],address,address,uint256))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) DecodeUniswapV2Data(callData []byte) (BaseStructV2SwapParams, error) {
	return _ZKVizingAADataHelp.Contract.DecodeUniswapV2Data(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodeUniswapV3Data is a free data retrieval call binding the contract method 0x92e8fd52.
//
// Solidity: function decodeUniswapV3Data(bytes callData) pure returns((uint8,uint24,uint160,address,address,address,address,uint256,uint256))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) DecodeUniswapV3Data(opts *bind.CallOpts, callData []byte) (BaseStructV3SwapParams, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "decodeUniswapV3Data", callData)

	if err != nil {
		return *new(BaseStructV3SwapParams), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseStructV3SwapParams)).(*BaseStructV3SwapParams)

	return out0, err

}

// DecodeUniswapV3Data is a free data retrieval call binding the contract method 0x92e8fd52.
//
// Solidity: function decodeUniswapV3Data(bytes callData) pure returns((uint8,uint24,uint160,address,address,address,address,uint256,uint256))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) DecodeUniswapV3Data(callData []byte) (BaseStructV3SwapParams, error) {
	return _ZKVizingAADataHelp.Contract.DecodeUniswapV3Data(&_ZKVizingAADataHelp.CallOpts, callData)
}

// DecodeUniswapV3Data is a free data retrieval call binding the contract method 0x92e8fd52.
//
// Solidity: function decodeUniswapV3Data(bytes callData) pure returns((uint8,uint24,uint160,address,address,address,address,uint256,uint256))
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) DecodeUniswapV3Data(callData []byte) (BaseStructV3SwapParams, error) {
	return _ZKVizingAADataHelp.Contract.DecodeUniswapV3Data(&_ZKVizingAADataHelp.CallOpts, callData)
}

// EncodeCrossETHParams is a free data retrieval call binding the contract method 0x97237cc8.
//
// Solidity: function encodeCrossETHParams((uint256,address) params) pure returns(bytes bytesCrossETHAmount)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) EncodeCrossETHParams(opts *bind.CallOpts, params BaseStructCrossETHParams) ([]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "encodeCrossETHParams", params)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeCrossETHParams is a free data retrieval call binding the contract method 0x97237cc8.
//
// Solidity: function encodeCrossETHParams((uint256,address) params) pure returns(bytes bytesCrossETHAmount)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) EncodeCrossETHParams(params BaseStructCrossETHParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeCrossETHParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeCrossETHParams is a free data retrieval call binding the contract method 0x97237cc8.
//
// Solidity: function encodeCrossETHParams((uint256,address) params) pure returns(bytes bytesCrossETHAmount)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) EncodeCrossETHParams(params BaseStructCrossETHParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeCrossETHParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeCrossHookMessageParams is a free data retrieval call binding the contract method 0x278225a7.
//
// Solidity: function encodeCrossHookMessageParams((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes) params) pure returns(bytes bytesCrossHookMessageParams)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) EncodeCrossHookMessageParams(opts *bind.CallOpts, params BaseStructCrossHookMessageParams) ([]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "encodeCrossHookMessageParams", params)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeCrossHookMessageParams is a free data retrieval call binding the contract method 0x278225a7.
//
// Solidity: function encodeCrossHookMessageParams((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes) params) pure returns(bytes bytesCrossHookMessageParams)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) EncodeCrossHookMessageParams(params BaseStructCrossHookMessageParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeCrossHookMessageParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeCrossHookMessageParams is a free data retrieval call binding the contract method 0x278225a7.
//
// Solidity: function encodeCrossHookMessageParams((uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes) params) pure returns(bytes bytesCrossHookMessageParams)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) EncodeCrossHookMessageParams(params BaseStructCrossHookMessageParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeCrossHookMessageParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeCrossMessageParams is a free data retrieval call binding the contract method 0x2205804a.
//
// Solidity: function encodeCrossMessageParams(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) params) pure returns(bytes crossMessage)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) EncodeCrossMessageParams(opts *bind.CallOpts, params BaseStructCrossMessageParams) ([]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "encodeCrossMessageParams", params)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeCrossMessageParams is a free data retrieval call binding the contract method 0x2205804a.
//
// Solidity: function encodeCrossMessageParams(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) params) pure returns(bytes crossMessage)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) EncodeCrossMessageParams(params BaseStructCrossMessageParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeCrossMessageParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeCrossMessageParams is a free data retrieval call binding the contract method 0x2205804a.
//
// Solidity: function encodeCrossMessageParams(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes)) params) pure returns(bytes crossMessage)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) EncodeCrossMessageParams(params BaseStructCrossMessageParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeCrossMessageParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeCrossV2SwapParams is a free data retrieval call binding the contract method 0xfb9f784d.
//
// Solidity: function encodeCrossV2SwapParams((uint8,uint8,uint256,uint256,address,address,address,address,uint256) params) pure returns(bytes crossV2MessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) EncodeCrossV2SwapParams(opts *bind.CallOpts, params BaseStructCrossV2SwapParams) ([]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "encodeCrossV2SwapParams", params)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeCrossV2SwapParams is a free data retrieval call binding the contract method 0xfb9f784d.
//
// Solidity: function encodeCrossV2SwapParams((uint8,uint8,uint256,uint256,address,address,address,address,uint256) params) pure returns(bytes crossV2MessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) EncodeCrossV2SwapParams(params BaseStructCrossV2SwapParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeCrossV2SwapParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeCrossV2SwapParams is a free data retrieval call binding the contract method 0xfb9f784d.
//
// Solidity: function encodeCrossV2SwapParams((uint8,uint8,uint256,uint256,address,address,address,address,uint256) params) pure returns(bytes crossV2MessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) EncodeCrossV2SwapParams(params BaseStructCrossV2SwapParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeCrossV2SwapParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeCrossV3SwapParams is a free data retrieval call binding the contract method 0xaea79baf.
//
// Solidity: function encodeCrossV3SwapParams((uint8,uint8,uint256,uint256,address,address,address,address,uint256) params) pure returns(bytes crossV3MessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) EncodeCrossV3SwapParams(opts *bind.CallOpts, params BaseStructCrossV2SwapParams) ([]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "encodeCrossV3SwapParams", params)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeCrossV3SwapParams is a free data retrieval call binding the contract method 0xaea79baf.
//
// Solidity: function encodeCrossV3SwapParams((uint8,uint8,uint256,uint256,address,address,address,address,uint256) params) pure returns(bytes crossV3MessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) EncodeCrossV3SwapParams(params BaseStructCrossV2SwapParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeCrossV3SwapParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeCrossV3SwapParams is a free data retrieval call binding the contract method 0xaea79baf.
//
// Solidity: function encodeCrossV3SwapParams((uint8,uint8,uint256,uint256,address,address,address,address,uint256) params) pure returns(bytes crossV3MessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) EncodeCrossV3SwapParams(params BaseStructCrossV2SwapParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeCrossV3SwapParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodePackedUserOperationGroup is a free data retrieval call binding the contract method 0xaf0ecfff.
//
// Solidity: function encodePackedUserOperationGroup((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] paramsGroup) pure returns(bytes bytesUserOperation)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) EncodePackedUserOperationGroup(opts *bind.CallOpts, paramsGroup []BaseStructPackedUserOperation) ([]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "encodePackedUserOperationGroup", paramsGroup)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePackedUserOperationGroup is a free data retrieval call binding the contract method 0xaf0ecfff.
//
// Solidity: function encodePackedUserOperationGroup((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] paramsGroup) pure returns(bytes bytesUserOperation)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) EncodePackedUserOperationGroup(paramsGroup []BaseStructPackedUserOperation) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodePackedUserOperationGroup(&_ZKVizingAADataHelp.CallOpts, paramsGroup)
}

// EncodePackedUserOperationGroup is a free data retrieval call binding the contract method 0xaf0ecfff.
//
// Solidity: function encodePackedUserOperationGroup((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] paramsGroup) pure returns(bytes bytesUserOperation)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) EncodePackedUserOperationGroup(paramsGroup []BaseStructPackedUserOperation) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodePackedUserOperationGroup(&_ZKVizingAADataHelp.CallOpts, paramsGroup)
}

// EncodeV2SwapParams is a free data retrieval call binding the contract method 0x8892b4fb.
//
// Solidity: function encodeV2SwapParams((uint8,uint256,uint256,address[],address,address,uint256) params) pure returns(bytes v2SwapMessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) EncodeV2SwapParams(opts *bind.CallOpts, params BaseStructV2SwapParams) ([]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "encodeV2SwapParams", params)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeV2SwapParams is a free data retrieval call binding the contract method 0x8892b4fb.
//
// Solidity: function encodeV2SwapParams((uint8,uint256,uint256,address[],address,address,uint256) params) pure returns(bytes v2SwapMessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) EncodeV2SwapParams(params BaseStructV2SwapParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeV2SwapParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeV2SwapParams is a free data retrieval call binding the contract method 0x8892b4fb.
//
// Solidity: function encodeV2SwapParams((uint8,uint256,uint256,address[],address,address,uint256) params) pure returns(bytes v2SwapMessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) EncodeV2SwapParams(params BaseStructV2SwapParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeV2SwapParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeV3SwapParams is a free data retrieval call binding the contract method 0x9f17b722.
//
// Solidity: function encodeV3SwapParams((uint8,uint24,uint160,address,address,address,address,uint256,uint256) params) pure returns(bytes v3SwapMessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) EncodeV3SwapParams(opts *bind.CallOpts, params BaseStructV3SwapParams) ([]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "encodeV3SwapParams", params)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeV3SwapParams is a free data retrieval call binding the contract method 0x9f17b722.
//
// Solidity: function encodeV3SwapParams((uint8,uint24,uint160,address,address,address,address,uint256,uint256) params) pure returns(bytes v3SwapMessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) EncodeV3SwapParams(params BaseStructV3SwapParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeV3SwapParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// EncodeV3SwapParams is a free data retrieval call binding the contract method 0x9f17b722.
//
// Solidity: function encodeV3SwapParams((uint8,uint24,uint160,address,address,address,address,uint256,uint256) params) pure returns(bytes v3SwapMessageBytes)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) EncodeV3SwapParams(params BaseStructV3SwapParams) ([]byte, error) {
	return _ZKVizingAADataHelp.Contract.EncodeV3SwapParams(&_ZKVizingAADataHelp.CallOpts, params)
}

// GetHasInnerExec is a free data retrieval call binding the contract method 0x2198e30c.
//
// Solidity: function getHasInnerExec((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) pure returns(bool)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) GetHasInnerExec(opts *bind.CallOpts, userOp BaseStructPackedUserOperation) (bool, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "getHasInnerExec", userOp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetHasInnerExec is a free data retrieval call binding the contract method 0x2198e30c.
//
// Solidity: function getHasInnerExec((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) pure returns(bool)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) GetHasInnerExec(userOp BaseStructPackedUserOperation) (bool, error) {
	return _ZKVizingAADataHelp.Contract.GetHasInnerExec(&_ZKVizingAADataHelp.CallOpts, userOp)
}

// GetHasInnerExec is a free data retrieval call binding the contract method 0x2198e30c.
//
// Solidity: function getHasInnerExec((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) pure returns(bool)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) GetHasInnerExec(userOp BaseStructPackedUserOperation) (bool, error) {
	return _ZKVizingAADataHelp.Contract.GetHasInnerExec(&_ZKVizingAADataHelp.CallOpts, userOp)
}

// GetPackChainGasLimit is a free data retrieval call binding the contract method 0x1ddd00d3.
//
// Solidity: function getPackChainGasLimit((uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes) exec) pure returns(bytes32)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) GetPackChainGasLimit(opts *bind.CallOpts, exec BaseStructExecData) ([32]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "getPackChainGasLimit", exec)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPackChainGasLimit is a free data retrieval call binding the contract method 0x1ddd00d3.
//
// Solidity: function getPackChainGasLimit((uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes) exec) pure returns(bytes32)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) GetPackChainGasLimit(exec BaseStructExecData) ([32]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackChainGasLimit(&_ZKVizingAADataHelp.CallOpts, exec)
}

// GetPackChainGasLimit is a free data retrieval call binding the contract method 0x1ddd00d3.
//
// Solidity: function getPackChainGasLimit((uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes) exec) pure returns(bytes32)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) GetPackChainGasLimit(exec BaseStructExecData) ([32]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackChainGasLimit(&_ZKVizingAADataHelp.CallOpts, exec)
}

// GetPackChainGasPrice is a free data retrieval call binding the contract method 0xa6750a01.
//
// Solidity: function getPackChainGasPrice((uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes) exec) pure returns(bytes32)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) GetPackChainGasPrice(opts *bind.CallOpts, exec BaseStructExecData) ([32]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "getPackChainGasPrice", exec)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPackChainGasPrice is a free data retrieval call binding the contract method 0xa6750a01.
//
// Solidity: function getPackChainGasPrice((uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes) exec) pure returns(bytes32)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) GetPackChainGasPrice(exec BaseStructExecData) ([32]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackChainGasPrice(&_ZKVizingAADataHelp.CallOpts, exec)
}

// GetPackChainGasPrice is a free data retrieval call binding the contract method 0xa6750a01.
//
// Solidity: function getPackChainGasPrice((uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes) exec) pure returns(bytes32)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) GetPackChainGasPrice(exec BaseStructExecData) ([32]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackChainGasPrice(&_ZKVizingAADataHelp.CallOpts, exec)
}

// GetPackOpInfo is a free data retrieval call binding the contract method 0x5c6432df.
//
// Solidity: function getPackOpInfo((uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes) exec) pure returns(bytes32)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) GetPackOpInfo(opts *bind.CallOpts, exec BaseStructExecData) ([32]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "getPackOpInfo", exec)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPackOpInfo is a free data retrieval call binding the contract method 0x5c6432df.
//
// Solidity: function getPackOpInfo((uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes) exec) pure returns(bytes32)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) GetPackOpInfo(exec BaseStructExecData) ([32]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackOpInfo(&_ZKVizingAADataHelp.CallOpts, exec)
}

// GetPackOpInfo is a free data retrieval call binding the contract method 0x5c6432df.
//
// Solidity: function getPackOpInfo((uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes) exec) pure returns(bytes32)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) GetPackOpInfo(exec BaseStructExecData) ([32]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackOpInfo(&_ZKVizingAADataHelp.CallOpts, exec)
}

// GetPackOperation is a free data retrieval call binding the contract method 0x3a7e52f5.
//
// Solidity: function getPackOperation((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) pure returns(bytes32 encoded)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) GetPackOperation(opts *bind.CallOpts, userOp BaseStructPackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "getPackOperation", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPackOperation is a free data retrieval call binding the contract method 0x3a7e52f5.
//
// Solidity: function getPackOperation((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) pure returns(bytes32 encoded)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) GetPackOperation(userOp BaseStructPackedUserOperation) ([32]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackOperation(&_ZKVizingAADataHelp.CallOpts, userOp)
}

// GetPackOperation is a free data retrieval call binding the contract method 0x3a7e52f5.
//
// Solidity: function getPackOperation((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) pure returns(bytes32 encoded)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) GetPackOperation(userOp BaseStructPackedUserOperation) ([32]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackOperation(&_ZKVizingAADataHelp.CallOpts, userOp)
}

// GetPackUint64s is a free data retrieval call binding the contract method 0xfce9e829.
//
// Solidity: function getPackUint64s(uint64 high64, uint64 low64) pure returns(bytes16 packed)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) GetPackUint64s(opts *bind.CallOpts, high64 uint64, low64 uint64) ([16]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "getPackUint64s", high64, low64)

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// GetPackUint64s is a free data retrieval call binding the contract method 0xfce9e829.
//
// Solidity: function getPackUint64s(uint64 high64, uint64 low64) pure returns(bytes16 packed)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) GetPackUint64s(high64 uint64, low64 uint64) ([16]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackUint64s(&_ZKVizingAADataHelp.CallOpts, high64, low64)
}

// GetPackUint64s is a free data retrieval call binding the contract method 0xfce9e829.
//
// Solidity: function getPackUint64s(uint64 high64, uint64 low64) pure returns(bytes16 packed)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) GetPackUint64s(high64 uint64, low64 uint64) ([16]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackUint64s(&_ZKVizingAADataHelp.CallOpts, high64, low64)
}

// GetPackUints is a free data retrieval call binding the contract method 0x1398fcdf.
//
// Solidity: function getPackUints(uint256 high128, uint256 low128) pure returns(bytes32 packed)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCaller) GetPackUints(opts *bind.CallOpts, high128 *big.Int, low128 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZKVizingAADataHelp.contract.Call(opts, &out, "getPackUints", high128, low128)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPackUints is a free data retrieval call binding the contract method 0x1398fcdf.
//
// Solidity: function getPackUints(uint256 high128, uint256 low128) pure returns(bytes32 packed)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpSession) GetPackUints(high128 *big.Int, low128 *big.Int) ([32]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackUints(&_ZKVizingAADataHelp.CallOpts, high128, low128)
}

// GetPackUints is a free data retrieval call binding the contract method 0x1398fcdf.
//
// Solidity: function getPackUints(uint256 high128, uint256 low128) pure returns(bytes32 packed)
func (_ZKVizingAADataHelp *ZKVizingAADataHelpCallerSession) GetPackUints(high128 *big.Int, low128 *big.Int) ([32]byte, error) {
	return _ZKVizingAADataHelp.Contract.GetPackUints(&_ZKVizingAADataHelp.CallOpts, high128, low128)
}
