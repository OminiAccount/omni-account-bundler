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

// BaseStructPackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type BaseStructPackedUserOperation struct {
	OperationType          uint8
	OperationValue         *big.Int
	Sender                 common.Address
	Nonce                  uint64
	ChainId                uint64
	CallData               []byte
	MainChainGasLimit      uint64
	DestChainGasLimit      uint64
	ZkVerificationGasLimit uint64
	MainChainGasPrice      uint64
	DestChainGasPrice      uint64
	Owner                  common.Address
}

// SyncRouterMetaData contains all meta data concerning the SyncRouter contract.
var SyncRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vizingPad\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_Hook\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LandingPadAccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplement\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"packHookMessage\",\"type\":\"bytes\"}],\"name\":\"ReceiveTouchHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"VizingSwapEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Hook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LandingPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LaunchPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LockWay\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"MirrorEntryPoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBridgeMode\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGasPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGaslimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destChainid\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"additionParams\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"estimateVizingGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vizingGasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"fetchOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"_packedUserOperation\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"fetchUserOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveExtraMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveStandardMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selectedRelayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"sendOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"_packedUserOperation\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"sendUserOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_way\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"_lockState\",\"type\":\"bytes1\"}],\"name\":\"setLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"}],\"name\":\"setMirrorEntryPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e0604052600160f81b60045f6101000a81548160ff021916908360f81c02179055505f67ffffffffffffffff8111156200003f576200003e6200038b565b5b6040519080825280601f01601f191660200182016040528015620000725781602001600182028036833780820191505090505b5060069081620000839190620005ef565b5061c35060075f6101000a81548162ffffff021916908362ffffff160217905550633b9aca00600760036101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550348015620000dd575f80fd5b50604051620056f7380380620056f7833981810160405281019062000103919062000738565b3383808162000118816200024360201b60201c565b506200012a816200028560201b60201c565b50505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200019f575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620001969190620007a2565b60405180910390fd5b620001b081620002c860201b60201c565b50600160038190555081600460016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620007bd565b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200040757607f821691505b6020821081036200041d576200041c620003c2565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620004817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000444565b6200048d868362000444565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620004d7620004d1620004cb84620004a5565b620004ae565b620004a5565b9050919050565b5f819050919050565b620004f283620004b7565b6200050a6200050182620004de565b84845462000450565b825550505050565b5f90565b6200052062000512565b6200052d818484620004e7565b505050565b5b818110156200055457620005485f8262000516565b60018101905062000533565b5050565b601f821115620005a3576200056d8162000423565b620005788462000435565b8101602085101562000588578190505b620005a0620005978562000435565b83018262000532565b50505b505050565b5f82821c905092915050565b5f620005c55f1984600802620005a8565b1980831691505092915050565b5f620005df8383620005b4565b9150826002028217905092915050565b620005fa82620003b8565b67ffffffffffffffff8111156200061657620006156200038b565b5b620006228254620003ef565b6200062f82828562000558565b5f60209050601f83116001811462000665575f841562000650578287015190505b6200065c8582620005d2565b865550620006cb565b601f198416620006758662000423565b5f5b828110156200069e5784890151825560018201915060208501945060208101905062000677565b86831015620006be5784890151620006ba601f891682620005b4565b8355505b6001600288020188555050505b505050505050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6200070282620006d7565b9050919050565b6200071481620006f6565b81146200071f575f80fd5b50565b5f81519050620007328162000709565b92915050565b5f805f60608486031215620007525762000751620006d3565b5b5f620007618682870162000722565b9350506020620007748682870162000722565b9250506040620007878682870162000722565b9150509250925092565b6200079c81620006f6565b82525050565b5f602082019050620007b75f83018462000791565b92915050565b60805160a05160c051614efa620007fd5f395f8181611d2d0152611f0e01525f8181611eed015261201c01525f8181610a360152611ecc0152614efa5ff3fe608060405260043610610183575f3560e01c80638ea23ddf116100d0578063d5f3e00811610089578063e0b838e911610063578063e0b838e914610542578063e7b4294c1461056c578063f2fde38b14610596578063f3148925146105be5761018a565b8063d5f3e008146104b4578063de4b1d6d146104dc578063de8aeda0146105185761018a565b80638ea23ddf146103b457806399e581fa146103de578063a400b25014610408578063ad5c464814610444578063b0cfd4d21461046e578063cd6a15f2146104985761018a565b80635aeb4d771161013d578063715018a611610117578063715018a61461031c57806376c81312146103325780637e08eb861461036e5780638da5cb5b1461038a5761018a565b80635aeb4d77146102a05780635e45da23146102ca5780636c7cfb52146102f45761018a565b806273b5551461018e5780630e82845d146101aa5780630ec9db66146101d45780631490a35814610210578063456362791461024c5780635ad3ad06146102765761018a565b3661018a57005b5f80fd5b6101a860048036038101906101a3919061297c565b6105da565b005b3480156101b5575f80fd5b506101be610672565b6040516101cb9190612a67565b60405180910390f35b3480156101df575f80fd5b506101fa60048036038101906101f59190612aa2565b610697565b6040516102079190612af8565b60405180910390f35b34801561021b575f80fd5b5061023660048036038101906102319190612b11565b610a00565b6040516102439190612b5c565b60405180910390f35b348015610257575f80fd5b50610260610a30565b60405161026d9190612baf565b60405180910390f35b348015610281575f80fd5b5061028a610a34565b6040516102979190612bd7565b60405180910390f35b3480156102ab575f80fd5b506102b4610a58565b6040516102c19190612c0d565b60405180910390f35b3480156102d5575f80fd5b506102de610a5c565b6040516102eb9190612c0d565b60405180910390f35b3480156102ff575f80fd5b5061031a60048036038101906103159190612c50565b610a60565b005b348015610327575f80fd5b50610330610a95565b005b34801561033d575f80fd5b5061035860048036038101906103539190612c8e565b610aa8565b6040516103659190612af8565b60405180910390f35b61038860048036038101906103839190612aa2565b610b47565b005b348015610395575f80fd5b5061039e611b8f565b6040516103ab9190612b5c565b60405180910390f35b3480156103bf575f80fd5b506103c8611bb7565b6040516103d59190612b5c565b60405180910390f35b3480156103e9575f80fd5b506103f2611bdc565b6040516103ff9190612c0d565b60405180910390f35b348015610413575f80fd5b5061042e60048036038101906104299190612db0565b611bf0565b60405161043b9190612af8565b60405180910390f35b34801561044f575f80fd5b50610458611d05565b6040516104659190612b5c565b60405180910390f35b348015610479575f80fd5b50610482611d2b565b60405161048f9190612b5c565b60405180910390f35b6104b260048036038101906104ad9190612db0565b611d4f565b005b3480156104bf575f80fd5b506104da60048036038101906104d59190612e34565b611f8e565b005b3480156104e7575f80fd5b5061050260048036038101906104fd9190612e72565b611ffd565b60405161050f9190612baf565b60405180910390f35b348015610523575f80fd5b5061052c61201a565b6040516105399190612bd7565b60405180910390f35b34801561054d575f80fd5b5061055661203e565b6040516105639190612a67565b60405180910390f35b348015610577575f80fd5b50610580612061565b60405161058d9190612bd7565b60405180910390f35b3480156105a1575f80fd5b506105bc60048036038101906105b79190612e9d565b61207b565b005b6105d860048036038101906105d39190612efb565b6120ff565b005b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610660576040517ffb2541ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61066c84848484612199565b50505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80826040516020016106aa91906134d9565b60405160208183030381529060405290505f808480602001906106cd9190613505565b5f0160208101906106de919061352d565b60ff1603610722575f8480602001906106f79190613505565b8061016001906107079190613558565b8101906107149190613687565b9050805f0151915050610897565b60018480602001906107349190613505565b5f016020810190610745919061352d565b60ff16036107c2575f84806020019061075e9190613505565b80610160019061076e9190613558565b81019061077b9190613779565b90505f73ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff16036107bc57806040015191505b50610896565b60028480602001906107d49190613505565b5f0160208101906107e5919061352d565b60ff1603610863575f8480602001906107fe9190613505565b80610160019061080e9190613558565b81019061081b91906138d8565b90505f73ffffffffffffffffffffffffffffffffffffffff168160c0015173ffffffffffffffffffffffffffffffffffffffff160361085d5780610120015191505b50610895565b6040517f82af0a8600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b5b5f61091660045f9054906101000a900460f81b8680602001906108ba9190613505565b60c00160208101906108cc9190612e9d565b8780602001906108dc9190613505565b60200160208101906108ee9190613904565b8880602001906108fe9190613505565b60400160208101906109109190612b11565b876121cb565b90505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166385fdd542838780602001906109649190613505565b6101000135610973919061395c565b8780602001906109839190613505565b60600160208101906109959190612b11565b6006856040518563ffffffff1660e01b81526004016109b79493929190613af9565b602060405180830381865afa1580156109d2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109f69190613b5e565b9350505050919050565b6008602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f90565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f90565b5f90565b610a68612216565b8060095f8481526020019081526020015f205f6101000a81548160ff021916908360f81c02179055505050565b610a9d612216565b610aa65f61229d565b565b5f610b3b878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505086868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050612360565b90509695505050505050565b610b4f612409565b5f8060f81b60095f8381526020019081526020015f205f9054906101000a900460f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610b9d575f80fd5b5f6060805f858060200190610bb29190613505565b5f016020810190610bc3919061352d565b60ff1603610c6b575f858060200190610bdc9190613505565b806101600190610bec9190613558565b810190610bf99190613687565b9050805f01519350858060200190610c119190613505565b806101600190610c219190613558565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050915050611611565b6001858060200190610c7d9190613505565b5f016020810190610c8e919061352d565b60ff1603611165575f858060200190610ca79190613505565b806101600190610cb79190613558565b810190610cc49190613779565b905060605f73ffffffffffffffffffffffffffffffffffffffff16826080015173ffffffffffffffffffffffffffffffffffffffff1603610e9857816040015194505f815f81518110610d1a57610d19613b89565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508160a0015181600181518110610d6d57610d6c613b89565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505f6040518060c00160405280846020015160ff168152602001846040015181526020015f81526020018381526020018460c0015173ffffffffffffffffffffffffffffffffffffffff1681526020018460e00151815250905060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630c07a79e82604051602401610e4e9190613ce4565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050945050611109565b8160800151815f81518110610eb057610eaf613b89565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505f81600181518110610eff57610efe613b89565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505f6040518060c00160405280845f015160ff16815260200184604001518152602001846060015181526020018381526020013073ffffffffffffffffffffffffffffffffffffffff1681526020018460e001518152509050610fc633308560400151866080015173ffffffffffffffffffffffffffffffffffffffff1661244f909392919063ffffffff16565b826080015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b360055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685604001516040518363ffffffff1660e01b815260040161102a929190613d04565b6020604051808303815f875af1158015611046573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061106a9190613d60565b5060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630c07a79e826040518263ffffffff1660e01b81526004016110c59190613ce4565b6020604051808303815f875af11580156110e1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111059190613b5e565b9550505b5f60405180604001604052808781526020018460c0015173ffffffffffffffffffffffffffffffffffffffff1681525090508060405160200161114c9190613db8565b6040516020818303038152906040529350505050611610565b60028580602001906111779190613505565b5f016020810190611188919061352d565b60ff16036115dd575f8580602001906111a19190613505565b8061016001906111b19190613558565b8101906111be91906138d8565b90505f73ffffffffffffffffffffffffffffffffffffffff168160c0015173ffffffffffffffffffffffffffffffffffffffff16036113535780610120015193505f604051806101000160405280836020015160ff168152602001836060015162ffffff1681526020018360a0015173ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020018360e0015173ffffffffffffffffffffffffffffffffffffffff16815260200183610100015173ffffffffffffffffffffffffffffffffffffffff1681526020018361012001518152602001836101400151815250905060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f87f5859826040516024016113099190613e80565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050935050611581565b5f604051806101000160405280835f015160ff168152602001836040015162ffffff168152602001836080015173ffffffffffffffffffffffffffffffffffffffff1681526020018360c0015173ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff1681526020018361012001518152602001836101400151815250905061143e33308360c00151846060015173ffffffffffffffffffffffffffffffffffffffff1661244f909392919063ffffffff16565b806060015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b360055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168360c001516040518363ffffffff1660e01b81526004016114a2929190613d04565b6020604051808303815f875af11580156114be573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114e29190613d60565b5060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f87f5859826040518263ffffffff1660e01b815260040161153d9190613e80565b6020604051808303815f875af1158015611559573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061157d9190613b5e565b9450505b5f604051806040016040528086815260200183610100015173ffffffffffffffffffffffffffffffffffffffff168152509050806040516020016115c59190613db8565b6040516020818303038152906040529250505061160f565b6040517f82af0a8600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b5b5f60405180610180016040528087806020019061162e9190613505565b5f01602081019061163f919061352d565b60ff1681526020018780602001906116579190613505565b60200160208101906116699190613904565b62ffffff1681526020018780602001906116839190613505565b60400160208101906116959190612b11565b67ffffffffffffffff1681526020018780602001906116b49190613505565b60600160208101906116c69190612b11565b67ffffffffffffffff1681526020018780602001906116e59190613505565b60800160208101906116f79190612b11565b67ffffffffffffffff1681526020018780602001906117169190613505565b60a00160208101906117289190612b11565b67ffffffffffffffff1681526020018780602001906117479190613505565b60c00160208101906117599190612e9d565b73ffffffffffffffffffffffffffffffffffffffff1681526020018780602001906117849190613505565b60e00160208101906117969190612e9d565b73ffffffffffffffffffffffffffffffffffffffff1681526020018780602001906117c19190613505565b610100013581526020018780602001906117db9190613505565b8061012001906117eb9190613558565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018481526020018381525090505f61190760045f9054906101000a900460f81b8880602001906118619190613505565b60c00160208101906118739190612e9d565b8980602001906118839190613505565b60200160208101906118959190613904565b8a80602001906118a59190613505565b60400160208101906118b79190612b11565b60405180604001604052808d805f01906118d19190613e9a565b6118da9061409f565b8152602001888152506040516020016118f3919061433a565b6040516020818303038152906040526121cb565b90505f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166385fdd542878a80602001906119579190613505565b6101000135611966919061395c565b8a80602001906119769190613505565b60600160208101906119889190612b11565b6006866040518563ffffffff1660e01b81526004016119aa9493929190613af9565b602060405180830381865afa1580156119c5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119e99190613b5e565b9050858880602001906119fc9190613505565b610100013582611a0c919061395c565b611a16919061395c565b341015611a4f576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663209afe56348a8060200190611a9b9190613505565b6080016020810190611aad9190612b11565b8b8060200190611abd9190613505565b60a0016020810190611acf9190612b11565b8c8060200190611adf9190613505565b60e0016020810190611af19190612e9d565b338e8060200190611b029190613505565b61010001358f8060200190611b179190613505565b6060016020810190611b299190612b11565b60068b6040518a63ffffffff1660e01b8152600401611b4f98979695949392919061435a565b5f604051808303818588803b158015611b66575f80fd5b505af1158015611b78573d5f803e3d5ffd5b505050505050505050505050611b8c6124d1565b50565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60075f9054906101000a900462ffffff1681565b5f808383604051602001611c0592919061448e565b60405160208183030381529060405290505f611c5860045f9054906101000a900460f81b8860075f9054906101000a900462ffffff16600760039054906101000a900467ffffffffffffffff16866121cb565b90505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166385fdd542878a6006856040518563ffffffff1660e01b8152600401611cb99493929190613af9565b602060405180830381865afa158015611cd4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611cf89190613b5e565b9250505095945050505050565b600460019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b4660085f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611e02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611df99061450a565b60405180910390fd5b5f611e6560045f9054906101000a900460f81b8760075f9054906101000a900462ffffff16600760039054906101000a900467ffffffffffffffff168888604051602001611e5192919061448e565b6040516020818303038152906040526121cb565b90505f611e758888888888611bf0565b90508581611e83919061395c565b341015611e8e575f80fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663209afe56347f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000338c8f60068b6040518a63ffffffff1660e01b8152600401611f5698979695949392919061435a565b5f604051808303818588803b158015611f6d575f80fd5b505af1158015611f7f573d5f803e3d5ffd5b50505050505050505050505050565b611f96612216565b8060085f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6009602052805f5260405f205f915054906101000a900460f81b81565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760039054906101000a900467ffffffffffffffff1681565b612083612216565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036120f3575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016120ea9190612b5c565b60405180910390fd5b6120fc8161229d565b50565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612185576040517ffb2541ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61219285858585856124db565b5050505050565b6040517fff94113200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060858573ffffffffffffffffffffffffffffffffffffffff168585856040516020016121fc95949392919061460a565b604051602081830303815290604052905095945050505050565b61221e6127f8565b73ffffffffffffffffffffffffffffffffffffffff1661223c611b8f565b73ffffffffffffffffffffffffffffffffffffffff161461229b5761225f6127f8565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016122929190612b5c565b60405180910390fd5b565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166385fdd542868686866040518563ffffffff1660e01b81526004016123c09493929190614664565b602060405180830381865afa1580156123db573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123ff9190613b5e565b9050949350505050565b600260035403612445576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600381905550565b6124cb848573ffffffffffffffffffffffffffffffffffffffff166323b872dd868686604051602401612484939291906146b5565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506127ff565b50505050565b6001600381905550565b8273ffffffffffffffffffffffffffffffffffffffff1660085f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461258d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161258490614734565b60405180910390fd5b5f828281019061259d919061494a565b90505f816020015161012001518060200190518101906125bd91906149ff565b90505f818060200190518101906125d49190614c9b565b905060085f4667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ba64b242826040518263ffffffff1660e01b81526004016126539190614d8d565b5f604051808303815f87803b15801561266a575f80fd5b505af115801561267c573d5f803e3d5ffd5b505050505f60605f856020015161016001518060200190518101906126a19190614dfa565b90505f86602001516101400151510361272c57806020015173ffffffffffffffffffffffffffffffffffffffff16815f01516040516126df90614e48565b5f6040518083038185875af1925050503d805f8114612719576040519150601f19603f3d011682016040523d82523d5f602084013e61271e565b606091505b5080935081945050506127a7565b3073ffffffffffffffffffffffffffffffffffffffff16815f01518760200151610140015160405161275e9190614e5c565b5f6040518083038185875af1925050503d805f8114612798576040519150601f19603f3d011682016040523d82523d5f602084013e61279d565b606091505b5080935081945050505b7fc9a6a0c9b51b72a63e2fa16e1e1293c367cbab5054a81dad8d8c1cb3980794468383886020015161014001516040516127e393929190614e81565b60405180910390a15050505050505050505050565b5f33905090565b5f8060205f8451602086015f885af18061281e576040513d5f823e3d81fd5b3d92505f519150505f8214612837576001811415612852565b5f8473ffffffffffffffffffffffffffffffffffffffff163b145b1561289457836040517f5274afe700000000000000000000000000000000000000000000000000000000815260040161288b9190612b5c565b60405180910390fd5b50505050565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b6128c7816128ab565b81146128d1575f80fd5b50565b5f813590506128e2816128be565b92915050565b5f819050919050565b6128fa816128e8565b8114612904575f80fd5b50565b5f81359050612915816128f1565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261293c5761293b61291b565b5b8235905067ffffffffffffffff8111156129595761295861291f565b5b60208301915083600182028301111561297557612974612923565b5b9250929050565b5f805f8060608587031215612994576129936128a3565b5b5f6129a1878288016128d4565b94505060206129b287828801612907565b935050604085013567ffffffffffffffff8111156129d3576129d26128a7565b5b6129df87828801612927565b925092505092959194509250565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f612a2f612a2a612a25846129ed565b612a0c565b6129ed565b9050919050565b5f612a4082612a15565b9050919050565b5f612a5182612a36565b9050919050565b612a6181612a47565b82525050565b5f602082019050612a7a5f830184612a58565b92915050565b5f80fd5b5f60408284031215612a9957612a98612a80565b5b81905092915050565b5f60208284031215612ab757612ab66128a3565b5b5f82013567ffffffffffffffff811115612ad457612ad36128a7565b5b612ae084828501612a84565b91505092915050565b612af2816128e8565b82525050565b5f602082019050612b0b5f830184612ae9565b92915050565b5f60208284031215612b2657612b256128a3565b5b5f612b33848285016128d4565b91505092915050565b5f612b46826129ed565b9050919050565b612b5681612b3c565b82525050565b5f602082019050612b6f5f830184612b4d565b92915050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b612ba981612b75565b82525050565b5f602082019050612bc25f830184612ba0565b92915050565b612bd1816128ab565b82525050565b5f602082019050612bea5f830184612bc8565b92915050565b5f62ffffff82169050919050565b612c0781612bf0565b82525050565b5f602082019050612c205f830184612bfe565b92915050565b612c2f81612b75565b8114612c39575f80fd5b50565b5f81359050612c4a81612c26565b92915050565b5f8060408385031215612c6657612c656128a3565b5b5f612c7385828601612907565b9250506020612c8485828601612c3c565b9150509250929050565b5f805f805f8060808789031215612ca857612ca76128a3565b5b5f612cb589828a01612907565b9650506020612cc689828a016128d4565b955050604087013567ffffffffffffffff811115612ce757612ce66128a7565b5b612cf389828a01612927565b9450945050606087013567ffffffffffffffff811115612d1657612d156128a7565b5b612d2289828a01612927565b92509250509295509295509295565b612d3a81612b3c565b8114612d44575f80fd5b50565b5f81359050612d5581612d31565b92915050565b5f8083601f840112612d7057612d6f61291b565b5b8235905067ffffffffffffffff811115612d8d57612d8c61291f565b5b602083019150836020820283011115612da957612da8612923565b5b9250929050565b5f805f805f60808688031215612dc957612dc86128a3565b5b5f612dd6888289016128d4565b9550506020612de788828901612d47565b9450506040612df888828901612907565b935050606086013567ffffffffffffffff811115612e1957612e186128a7565b5b612e2588828901612d5b565b92509250509295509295909350565b5f8060408385031215612e4a57612e496128a3565b5b5f612e57858286016128d4565b9250506020612e6885828601612d47565b9150509250929050565b5f60208284031215612e8757612e866128a3565b5b5f612e9484828501612907565b91505092915050565b5f60208284031215612eb257612eb16128a3565b5b5f612ebf84828501612d47565b91505092915050565b5f819050919050565b612eda81612ec8565b8114612ee4575f80fd5b50565b5f81359050612ef581612ed1565b92915050565b5f805f805f60808688031215612f1457612f136128a3565b5b5f612f2188828901612ee7565b9550506020612f32888289016128d4565b9450506040612f4388828901612907565b935050606086013567ffffffffffffffff811115612f6457612f636128a7565b5b612f7088828901612927565b92509250509295509295909350565b5f80fd5b5f8235600161018003833603038112612f9f57612f9e612f7f565b5b82810191505092915050565b5f60ff82169050919050565b612fc081612fab565b8114612fca575f80fd5b50565b5f81359050612fdb81612fb7565b92915050565b5f612fef6020840184612fcd565b905092915050565b61300081612fab565b82525050565b5f6130146020840184612907565b905092915050565b613025816128e8565b82525050565b5f6130396020840184612d47565b905092915050565b61304a81612b3c565b82525050565b5f61305e60208401846128d4565b905092915050565b61306f816128ab565b82525050565b5f80fd5b5f80fd5b5f808335600160200384360303811261309957613098612f7f565b5b83810192508235915060208301925067ffffffffffffffff8211156130c1576130c0613075565b5b6001820236038313156130d7576130d6613079565b5b509250929050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f61311883856130df565b93506131258385846130ef565b61312e836130fd565b840190509392505050565b5f610180830161314b5f840184612fe1565b6131575f860182612ff7565b506131656020840184613006565b613172602086018261301c565b50613180604084018461302b565b61318d6040860182613041565b5061319b6060840184613050565b6131a86060860182613066565b506131b66080840184613050565b6131c36080860182613066565b506131d160a084018461307d565b85830360a08701526131e483828461310d565b925050506131f560c0840184613050565b61320260c0860182613066565b5061321060e0840184613050565b61321d60e0860182613066565b5061322c610100840184613050565b61323a610100860182613066565b50613249610120840184613050565b613257610120860182613066565b50613266610140840184613050565b613274610140860182613066565b5061328361016084018461302b565b613291610160860182613041565b508091505092915050565b5f82356001610180038336030381126132b8576132b7612f7f565b5b82810191505092915050565b6132cd81612bf0565b81146132d7575f80fd5b50565b5f813590506132e8816132c4565b92915050565b5f6132fc60208401846132da565b905092915050565b61330d81612bf0565b82525050565b5f61018083016133255f840184612fe1565b6133315f860182612ff7565b5061333f60208401846132ee565b61334c6020860182613304565b5061335a6040840184613050565b6133676040860182613066565b506133756060840184613050565b6133826060860182613066565b506133906080840184613050565b61339d6080860182613066565b506133ab60a0840184613050565b6133b860a0860182613066565b506133c660c084018461302b565b6133d360c0860182613041565b506133e160e084018461302b565b6133ee60e0860182613041565b506133fd610100840184613006565b61340b61010086018261301c565b5061341a61012084018461307d565b85830361012087015261342e83828461310d565b9250505061344061014084018461307d565b85830361014087015261345483828461310d565b9250505061346661016084018461307d565b85830361016087015261347a83828461310d565b925050508091505092915050565b5f604083016134995f840184612f83565b8482035f8601526134aa8282613139565b9150506134ba602084018461329c565b84820360208601526134cc8282613313565b9150508091505092915050565b5f6020820190508181035f8301526134f18184613488565b905092915050565b5f80fd5b5f80fd5b5f80fd5b5f8235600161018003833603038112613521576135206134f9565b5b80830191505092915050565b5f60208284031215613542576135416128a3565b5b5f61354f84828501612fcd565b91505092915050565b5f8083356001602003843603038112613574576135736134f9565b5b80840192508235915067ffffffffffffffff821115613596576135956134fd565b5b6020830192506001820236038313156135b2576135b1613501565b5b509250929050565b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6135f4826130fd565b810181811067ffffffffffffffff82111715613613576136126135be565b5b80604052505050565b5f61362561289a565b905061363182826135eb565b919050565b5f80fd5b5f6040828403121561364f5761364e6135ba565b5b613659604061361c565b90505f61366884828501612907565b5f83015250602061367b84828501612d47565b60208301525092915050565b5f6040828403121561369c5761369b6128a3565b5b5f6136a98482850161363a565b91505092915050565b5f61010082840312156136c8576136c76135ba565b5b6136d361010061361c565b90505f6136e284828501612fcd565b5f8301525060206136f584828501612fcd565b602083015250604061370984828501612907565b604083015250606061371d84828501612907565b606083015250608061373184828501612d47565b60808301525060a061374584828501612d47565b60a08301525060c061375984828501612d47565b60c08301525060e061376d84828501612907565b60e08301525092915050565b5f610100828403121561378f5761378e6128a3565b5b5f61379c848285016136b2565b91505092915050565b6137ae816129ed565b81146137b8575f80fd5b50565b5f813590506137c9816137a5565b92915050565b5f61016082840312156137e5576137e46135ba565b5b6137f061016061361c565b90505f6137ff84828501612fcd565b5f83015250602061381284828501612fcd565b6020830152506040613826848285016132da565b604083015250606061383a848285016132da565b606083015250608061384e848285016137bb565b60808301525060a0613862848285016137bb565b60a08301525060c061387684828501612d47565b60c08301525060e061388a84828501612d47565b60e08301525061010061389f84828501612d47565b610100830152506101206138b584828501612907565b610120830152506101406138cb84828501612907565b6101408301525092915050565b5f61016082840312156138ee576138ed6128a3565b5b5f6138fb848285016137cf565b91505092915050565b5f60208284031215613919576139186128a3565b5b5f613926848285016132da565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f613966826128e8565b9150613971836128e8565b92508282019050808211156139895761398861392f565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806139d357607f821691505b6020821081036139e6576139e561398f565b5b50919050565b5f82825260208201905092915050565b5f819050815f5260205f209050919050565b5f8154613a1a816139bc565b613a2481866139ec565b9450600182165f8114613a3e5760018114613a5457613a86565b60ff198316865281151560200286019350613a86565b613a5d856139fc565b5f5b83811015613a7e57815481890152600182019150602081019050613a5f565b808801955050505b50505092915050565b5f81519050919050565b5f5b83811015613ab6578082015181840152602081019050613a9b565b5f8484015250505050565b5f613acb82613a8f565b613ad581856139ec565b9350613ae5818560208601613a99565b613aee816130fd565b840191505092915050565b5f608082019050613b0c5f830187612ae9565b613b196020830186612bc8565b8181036040830152613b2b8185613a0e565b90508181036060830152613b3f8184613ac1565b905095945050505050565b5f81519050613b58816128f1565b92915050565b5f60208284031215613b7357613b726128a3565b5b5f613b8084828501613b4a565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f613bea8383613041565b60208301905092915050565b5f602082019050919050565b5f613c0c82613bb6565b613c168185613bc0565b9350613c2183613bd0565b805f5b83811015613c51578151613c388882613bdf565b9750613c4383613bf6565b925050600181019050613c24565b5085935050505092915050565b5f60c083015f830151613c735f860182612ff7565b506020830151613c86602086018261301c565b506040830151613c99604086018261301c565b5060608301518482036060860152613cb18282613c02565b9150506080830151613cc66080860182613041565b5060a0830151613cd960a086018261301c565b508091505092915050565b5f6020820190508181035f830152613cfc8184613c5e565b905092915050565b5f604082019050613d175f830185612b4d565b613d246020830184612ae9565b9392505050565b5f8115159050919050565b613d3f81613d2b565b8114613d49575f80fd5b50565b5f81519050613d5a81613d36565b92915050565b5f60208284031215613d7557613d746128a3565b5b5f613d8284828501613d4c565b91505092915050565b604082015f820151613d9f5f85018261301c565b506020820151613db26020850182613041565b50505050565b5f604082019050613dcb5f830184613d8b565b92915050565b613dda816129ed565b82525050565b61010082015f820151613df55f850182612ff7565b506020820151613e086020850182613304565b506040820151613e1b6040850182613dd1565b506060820151613e2e6060850182613041565b506080820151613e416080850182613041565b5060a0820151613e5460a0850182613041565b5060c0820151613e6760c085018261301c565b5060e0820151613e7a60e085018261301c565b50505050565b5f61010082019050613e945f830184613de0565b92915050565b5f8235600161018003833603038112613eb657613eb56134f9565b5b80830191505092915050565b5f80fd5b5f67ffffffffffffffff821115613ee057613edf6135be565b5b613ee9826130fd565b9050602081019050919050565b5f613f08613f0384613ec6565b61361c565b905082815260208101848484011115613f2457613f23613ec2565b5b613f2f8482856130ef565b509392505050565b5f82601f830112613f4b57613f4a61291b565b5b8135613f5b848260208601613ef6565b91505092915050565b5f6101808284031215613f7a57613f796135ba565b5b613f8561018061361c565b90505f613f9484828501612fcd565b5f830152506020613fa784828501612907565b6020830152506040613fbb84828501612d47565b6040830152506060613fcf848285016128d4565b6060830152506080613fe3848285016128d4565b60808301525060a082013567ffffffffffffffff81111561400757614006613636565b5b61401384828501613f37565b60a08301525060c0614027848285016128d4565b60c08301525060e061403b848285016128d4565b60e083015250610100614050848285016128d4565b61010083015250610120614066848285016128d4565b6101208301525061014061407c848285016128d4565b6101408301525061016061409284828501612d47565b6101608301525092915050565b5f6140aa3683613f64565b9050919050565b5f6140bb82613a8f565b6140c581856130df565b93506140d5818560208601613a99565b6140de816130fd565b840191505092915050565b5f61018083015f8301516140ff5f860182612ff7565b506020830151614112602086018261301c565b5060408301516141256040860182613041565b5060608301516141386060860182613066565b50608083015161414b6080860182613066565b5060a083015184820360a086015261416382826140b1565b91505060c083015161417860c0860182613066565b5060e083015161418b60e0860182613066565b506101008301516141a0610100860182613066565b506101208301516141b5610120860182613066565b506101408301516141ca610140860182613066565b506101608301516141df610160860182613041565b508091505092915050565b5f61018083015f8301516142005f860182612ff7565b5060208301516142136020860182613304565b5060408301516142266040860182613066565b5060608301516142396060860182613066565b50608083015161424c6080860182613066565b5060a083015161425f60a0860182613066565b5060c083015161427260c0860182613041565b5060e083015161428560e0860182613041565b5061010083015161429a61010086018261301c565b506101208301518482036101208601526142b482826140b1565b9150506101408301518482036101408601526142d082826140b1565b9150506101608301518482036101608601526142ec82826140b1565b9150508091505092915050565b5f604083015f8301518482035f86015261431382826140e9565b9150506020830151848203602086015261432d82826141ea565b9150508091505092915050565b5f6020820190508181035f83015261435281846142f9565b905092915050565b5f6101008201905061436e5f83018b612bc8565b61437b602083018a612bc8565b6143886040830189612b4d565b6143956060830188612b4d565b6143a26080830187612ae9565b6143af60a0830186612bc8565b81810360c08301526143c18185613a0e565b905081810360e08301526143d58184613ac1565b90509998505050505050505050565b5f82825260208201905092915050565b5f819050919050565b5f6144088383613139565b905092915050565b5f602082019050919050565b5f61442783856143e4565b935083602084028501614439846143f4565b805f5b8781101561447c5784840389526144538284612f83565b61445d85826143fd565b945061446883614410565b925060208a0199505060018101905061443c565b50829750879450505050509392505050565b5f6020820190508181035f8301526144a781848661441c565b90509392505050565b5f82825260208201905092915050565b7f4d455000000000000000000000000000000000000000000000000000000000005f82015250565b5f6144f46003836144b0565b91506144ff826144c0565b602082019050919050565b5f6020820190508181035f830152614521816144e8565b9050919050565b5f819050919050565b61454261453d82612b75565b614528565b82525050565b5f819050919050565b61456261455d826128e8565b614548565b82525050565b5f8160e81b9050919050565b5f61457e82614568565b9050919050565b61459661459182612bf0565b614574565b82525050565b5f8160c01b9050919050565b5f6145b28261459c565b9050919050565b6145ca6145c5826128ab565b6145a8565b82525050565b5f81905092915050565b5f6145e482613a8f565b6145ee81856145d0565b93506145fe818560208601613a99565b80840191505092915050565b5f6146158288614531565b6001820191506146258287614551565b6020820191506146358286614585565b60038201915061464582856145b9565b60088201915061465582846145da565b91508190509695505050505050565b5f6080820190506146775f830187612ae9565b6146846020830186612bc8565b81810360408301526146968185613ac1565b905081810360608301526146aa8184613ac1565b905095945050505050565b5f6060820190506146c85f830186612b4d565b6146d56020830185612b4d565b6146e26040830184612ae9565b949350505050565b7f496e76616c696420636f6e7472616374000000000000000000000000000000005f82015250565b5f61471e6010836144b0565b9150614729826146ea565b602082019050919050565b5f6020820190508181035f83015261474b81614712565b9050919050565b5f6101808284031215614768576147676135ba565b5b61477361018061361c565b90505f61478284828501612fcd565b5f830152506020614795848285016132da565b60208301525060406147a9848285016128d4565b60408301525060606147bd848285016128d4565b60608301525060806147d1848285016128d4565b60808301525060a06147e5848285016128d4565b60a08301525060c06147f984828501612d47565b60c08301525060e061480d84828501612d47565b60e08301525061010061482284828501612907565b6101008301525061012082013567ffffffffffffffff81111561484857614847613636565b5b61485484828501613f37565b6101208301525061014082013567ffffffffffffffff81111561487a57614879613636565b5b61488684828501613f37565b6101408301525061016082013567ffffffffffffffff8111156148ac576148ab613636565b5b6148b884828501613f37565b6101608301525092915050565b5f604082840312156148da576148d96135ba565b5b6148e4604061361c565b90505f82013567ffffffffffffffff81111561490357614902613636565b5b61490f84828501613f64565b5f83015250602082013567ffffffffffffffff81111561493257614931613636565b5b61493e84828501614752565b60208301525092915050565b5f6020828403121561495f5761495e6128a3565b5b5f82013567ffffffffffffffff81111561497c5761497b6128a7565b5b614988848285016148c5565b91505092915050565b5f6149a361499e84613ec6565b61361c565b9050828152602081018484840111156149bf576149be613ec2565b5b6149ca848285613a99565b509392505050565b5f82601f8301126149e6576149e561291b565b5b81516149f6848260208601614991565b91505092915050565b5f60208284031215614a1457614a136128a3565b5b5f82015167ffffffffffffffff811115614a3157614a306128a7565b5b614a3d848285016149d2565b91505092915050565b5f67ffffffffffffffff821115614a6057614a5f6135be565b5b602082029050602081019050919050565b5f81519050614a7f81612fb7565b92915050565b5f81519050614a9381612d31565b92915050565b5f81519050614aa7816128be565b92915050565b5f6101808284031215614ac357614ac26135ba565b5b614ace61018061361c565b90505f614add84828501614a71565b5f830152506020614af084828501613b4a565b6020830152506040614b0484828501614a85565b6040830152506060614b1884828501614a99565b6060830152506080614b2c84828501614a99565b60808301525060a082015167ffffffffffffffff811115614b5057614b4f613636565b5b614b5c848285016149d2565b60a08301525060c0614b7084828501614a99565b60c08301525060e0614b8484828501614a99565b60e083015250610100614b9984828501614a99565b61010083015250610120614baf84828501614a99565b61012083015250610140614bc584828501614a99565b61014083015250610160614bdb84828501614a85565b6101608301525092915050565b5f614bfa614bf584614a46565b61361c565b90508083825260208201905060208402830185811115614c1d57614c1c612923565b5b835b81811015614c6457805167ffffffffffffffff811115614c4257614c4161291b565b5b808601614c4f8982614aad565b85526020850194505050602081019050614c1f565b5050509392505050565b5f82601f830112614c8257614c8161291b565b5b8151614c92848260208601614be8565b91505092915050565b5f60208284031215614cb057614caf6128a3565b5b5f82015167ffffffffffffffff811115614ccd57614ccc6128a7565b5b614cd984828501614c6e565b91505092915050565b5f81519050919050565b5f819050602082019050919050565b5f614d0683836140e9565b905092915050565b5f602082019050919050565b5f614d2482614ce2565b614d2e81856143e4565b935083602082028501614d4085614cec565b805f5b85811015614d7b5784840389528151614d5c8582614cfb565b9450614d6783614d0e565b925060208a01995050600181019050614d43565b50829750879550505050505092915050565b5f6020820190508181035f830152614da58184614d1a565b905092915050565b5f60408284031215614dc257614dc16135ba565b5b614dcc604061361c565b90505f614ddb84828501613b4a565b5f830152506020614dee84828501614a85565b60208301525092915050565b5f60408284031215614e0f57614e0e6128a3565b5b5f614e1c84828501614dad565b91505092915050565b50565b5f614e335f836145d0565b9150614e3e82614e25565b5f82019050919050565b5f614e5282614e28565b9150819050919050565b5f614e6782846145da565b915081905092915050565b614e7b81613d2b565b82525050565b5f606082019050614e945f830186614e72565b8181036020830152614ea68185613ac1565b90508181036040830152614eba8184613ac1565b905094935050505056fea264697066735822122045d603ce3186cf71d1c7d3c5e69c90a50cd07949c5a5f9e1bb3cebbe5200d1e564736f6c63430008180033",
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

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0xa400b250.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] userOperations) view returns(uint256)
func (_SyncRouter *SyncRouterCaller) FetchOmniMessageFee(opts *bind.CallOpts, destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "fetchOmniMessageFee", destChainId, destContract, destChainUsedFee, userOperations)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0xa400b250.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] userOperations) view returns(uint256)
func (_SyncRouter *SyncRouterSession) FetchOmniMessageFee(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*big.Int, error) {
	return _SyncRouter.Contract.FetchOmniMessageFee(&_SyncRouter.CallOpts, destChainId, destContract, destChainUsedFee, userOperations)
}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0xa400b250.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] userOperations) view returns(uint256)
func (_SyncRouter *SyncRouterCallerSession) FetchOmniMessageFee(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*big.Int, error) {
	return _SyncRouter.Contract.FetchOmniMessageFee(&_SyncRouter.CallOpts, destChainId, destContract, destChainUsedFee, userOperations)
}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0x0ec9db66.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterCaller) FetchUserOmniMessageFee(opts *bind.CallOpts, params BaseStructCrossMessageParams) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "fetchUserOmniMessageFee", params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0x0ec9db66.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256 _gasFee)
func (_SyncRouter *SyncRouterSession) FetchUserOmniMessageFee(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _SyncRouter.Contract.FetchUserOmniMessageFee(&_SyncRouter.CallOpts, params)
}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0x0ec9db66.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256 _gasFee)
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

// SendOmniMessage is a paid mutator transaction binding the contract method 0xcd6a15f2.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainExecuteUsedFee, (uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] userOperations) payable returns()
func (_SyncRouter *SyncRouterTransactor) SendOmniMessage(opts *bind.TransactOpts, destChainId uint64, destContract common.Address, destChainExecuteUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "sendOmniMessage", destChainId, destContract, destChainExecuteUsedFee, userOperations)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0xcd6a15f2.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainExecuteUsedFee, (uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] userOperations) payable returns()
func (_SyncRouter *SyncRouterSession) SendOmniMessage(destChainId uint64, destContract common.Address, destChainExecuteUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendOmniMessage(&_SyncRouter.TransactOpts, destChainId, destContract, destChainExecuteUsedFee, userOperations)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0xcd6a15f2.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainExecuteUsedFee, (uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] userOperations) payable returns()
func (_SyncRouter *SyncRouterTransactorSession) SendOmniMessage(destChainId uint64, destContract common.Address, destChainExecuteUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendOmniMessage(&_SyncRouter.TransactOpts, destChainId, destContract, destChainExecuteUsedFee, userOperations)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x7e08eb86.
//
// Solidity: function sendUserOmniMessage(((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_SyncRouter *SyncRouterTransactor) SendUserOmniMessage(opts *bind.TransactOpts, params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "sendUserOmniMessage", params)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x7e08eb86.
//
// Solidity: function sendUserOmniMessage(((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_SyncRouter *SyncRouterSession) SendUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendUserOmniMessage(&_SyncRouter.TransactOpts, params)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x7e08eb86.
//
// Solidity: function sendUserOmniMessage(((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address),(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
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
