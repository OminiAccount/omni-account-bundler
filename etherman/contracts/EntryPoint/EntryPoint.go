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
	Sender                 common.Address
	ChainId                *big.Int
	OperationValue         *big.Int
	ZkVerificationGasLimit *big.Int
	MainChainGasLimit      *big.Int
	DestChainGasLimit      *big.Int
	MainChainGasPrice      *big.Int
	DestChainGasPrice      *big.Int
}

// EntryPointUserOpInfo is an auto generated low-level Go binding around an user-defined struct.
type EntryPointUserOpInfo struct {
	MUserOp       EntryPointMemoryUserOp
	UserOpHash    [32]byte
	Prefund       *big.Int
	ContextOffset *big.Int
	PreOpGas      *big.Int
}

// IEntryPointBatchData is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointBatchData struct {
	UserOperations []PackedUserOperation
	OldStateRoot   [32]byte
	NewStateRoot   [32]byte
	AccInputHash   [32]byte
}

// IEntryPointChainExecuteExtra is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointChainExecuteExtra struct {
	ChainId                   *big.Int
	ChainFee                  *big.Int
	ChainUserOperationsNumber *big.Int
}

// IEntryPointChainsExecuteInfo is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointChainsExecuteInfo struct {
	ChainExtra  []IEntryPointChainExecuteExtra
	Beneficiary common.Address
}

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	OperationType          uint8
	OperationValue         *big.Int
	Sender                 common.Address
	Nonce                  uint64
	ChainId                uint64
	CallData               []byte
	MainChainGasLimit      *big.Int
	DestChainGasLimit      *big.Int
	ZkVerificationGasLimit *big.Int
	MainChainGasPrice      *big.Int
	DestChainGasPrice      *big.Int
	Owner                  common.Address
}

// EntryPointMetaData contains all meta data concerning the EntryPoint contract.
var EntryPointMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"name\":\"DelegateAndRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"inner\",\"type\":\"bytes\"}],\"name\":\"FailedOpWithRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"PostOpReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotEqual\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"did\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dstEids\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPreGasBalanceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"mainChainGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mainChainGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"mainChainGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mainChainGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSync\",\"type\":\"bool\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mainChainGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mainChainGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainGasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structEntryPoint.MemoryUserOp\",\"name\":\"mUserOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contextOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"}],\"internalType\":\"structEntryPoint.UserOpInfo\",\"name\":\"opInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"innerHandleOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"preGasBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"redeemGasOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"smtRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitDepositOperation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submitWithdrawOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"mainChainGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mainChainGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"}],\"name\":\"syncBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstCoeffGas\",\"type\":\"uint256\"}],\"name\":\"updateDstCoeffGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstConGas\",\"type\":\"uint256\"}],\"name\":\"updateDstConGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_dstEids\",\"type\":\"uint32[]\"}],\"name\":\"updateDstEids\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oldSmtRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newSmtRoot\",\"type\":\"bytes32\"}],\"name\":\"updateSmtRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_syncRouter\",\"type\":\"address\"}],\"name\":\"updateSyncRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"mainChainGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mainChainGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIEntryPoint.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainUserOperationsNumber\",\"type\":\"uint256\"}],\"internalType\":\"structIEntryPoint.ChainExecuteExtra[]\",\"name\":\"chainExtra\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structIEntryPoint.ChainsExecuteInfo\",\"name\":\"chainsExecuteInfos\",\"type\":\"tuple\"}],\"name\":\"verifyBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040527fb178c245c947ea7e21ecede07728941a6ab1b706143c06873baff8ebd6de63085f1b5f5534801562000035575f80fd5b503360016007819055505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000b2575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000a99190620001d0565b60405180910390fd5b620000c381620000ca60201b60201c565b50620001eb565b5f60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620001b8826200018d565b9050919050565b620001ca81620001ac565b82525050565b5f602082019050620001e55f830184620001bf565b92915050565b61508780620001f95f395ff3fe608060405260043610610165575f3560e01c806384b2b1b1116100d057806397fc007c116100895780639e998d4b116100635780639e998d4b146104df578063a6afc1de14610507578063c989e90c14610543578063f2fde38b1461057f57610165565b806397fc007c14610473578063988249081461049b5780639abe7903146104c357610165565b806384b2b1b11461036d578063850aaf621461039557806389b287f0146103bd5780638cb53415146103e55780638da5cb5b1461040d57806396899d3a1461043757610165565b8063653e75ec11610122578063653e75ec14610285578063654e8b23146102c1578063671d6c35146102e95780636cbe19af14610305578063715018a61461032f5780637fcb52271461034557610165565b806301ffc9a71461016957806305dc0922146101a55780631f14e001146101e15780632b7ac3f31461020b57806339875d1a146102355780633d939c4e1461025d575b5f80fd5b348015610174575f80fd5b5061018f600480360381019061018a9190612b77565b6105a7565b60405161019c9190612bbc565b60405180910390f35b3480156101b0575f80fd5b506101cb60048036038101906101c69190612c08565b610712565b6040516101d89190612c51565b60405180910390f35b3480156101ec575f80fd5b506101f5610749565b6040516102029190612c82565b60405180910390f35b348015610216575f80fd5b5061021f61074e565b60405161022c9190612cda565b60405180910390f35b348015610240575f80fd5b5061025b60048036038101906102569190612c08565b610773565b005b348015610268575f80fd5b50610283600480360381019061027e9190612db9565b6107c6565b005b348015610290575f80fd5b506102ab60048036038101906102a69190612e54565b6109c9565b6040516102b89190612e8e565b60405180910390f35b3480156102cc575f80fd5b506102e760048036038101906102e29190612ea7565b610a0f565b005b61030360048036038101906102fe9190612fb1565b610b07565b005b348015610310575f80fd5b50610319610dc5565b6040516103269190612cda565b60405180910390f35b34801561033a575f80fd5b50610343610dea565b005b348015610350575f80fd5b5061036b60048036038101906103669190612e54565b610dfd565b005b348015610378575f80fd5b50610393600480360381019061038e919061342a565b610e48565b005b3480156103a0575f80fd5b506103bb60048036038101906103b69190613471565b610ee6565b005b3480156103c8575f80fd5b506103e360048036038101906103de9190613523565b610f90565b005b3480156103f0575f80fd5b5061040b60048036038101906104069190612c08565b610fae565b005b348015610418575f80fd5b50610421610fc0565b60405161042e9190612cda565b60405180910390f35b348015610442575f80fd5b5061045d600480360381019061045891906136ed565b610fe8565b60405161046a9190612e8e565b60405180910390f35b34801561047e575f80fd5b5061049960048036038101906104949190612e54565b611177565b005b3480156104a6575f80fd5b506104c160048036038101906104bc919061377c565b6111c2565b005b6104dd60048036038101906104d89190612ea7565b611253565b005b3480156104ea575f80fd5b5061050560048036038101906105009190612c08565b6112ed565b005b348015610512575f80fd5b5061052d60048036038101906105289190612e54565b6112ff565b60405161053a9190612e8e565b60405180910390f35b34801561054e575f80fd5b50610569600480360381019061056491906137d9565b611314565b6040516105769190612c82565b60405180910390f35b34801561058a575f80fd5b506105a560048036038101906105a09190612e54565b61132c565b005b5f7fa349dad6000000000000000000000000000000000000000000000000000000007ff5a26b9100000000000000000000000000000000000000000000000000000000187bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061069357507ff5a26b91000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806106fb57507fa349dad6000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061070b575061070a826113b0565b5b9050919050565b60048181548110610721575f80fd5b905f5260205f209060089182820401919006600402915054906101000a900463ffffffff1681565b5f5481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff167f2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a82426040516107bb929190613820565b60405180910390a250565b6107ce611419565b5f8484905090505f8167ffffffffffffffff8111156107f0576107ef61306e565b5b60405190808252806020026020018201604052801561082957816020015b6108166128e5565b81526020019060019003908161080e5790505b5090505f5b82811015610891575f82828151811061084a57610849613847565b5b602002602001015190506108838289898581811061086b5761086a613847565b5b905060200281019061087d9190613880565b8361145f565b50808060010191505061082e565b505f7fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f97260405160405180910390a15f5b838110156109a7576108f68888838181106108df576108de613847565b5b90506020028101906108f19190613880565b6116c6565b1561094c5761094388888381811061091157610910613847565b5b90506020028101906109239190613880565b84838151811061093657610935613847565b5b602002602001015161170b565b8201915061099a565b6109958189898481811061096357610962613847565b5b90506020028101906109759190613880565b85848151811061098857610987613847565b5b6020026020010151611770565b820191505b80806001019150506108c1565b50836109b8576109b78582611af7565b5b5050506109c3611c12565b50505050565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b8160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610a86576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1633468385604051602001610ab4949392919061390d565b604051602081830303815290604052805190602001207f081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc8442604051610afb929190613820565b60405180910390a35050565b5f81805f0190610b17919061395a565b905067ffffffffffffffff811115610b3257610b3161306e565b5b604051908082528060200260200182016040528015610b6b57816020015b610b58612919565b815260200190600190039081610b505790505b5090505f8484905067ffffffffffffffff811115610b8c57610b8b61306e565b5b604051908082528060200260200182016040528015610bba5781602001602082028036833780820191505090505b5090505f5b83805f0190610bce919061395a565b9050811015610d43575f84805f0190610be7919061395a565b83818110610bf857610bf7613847565b5b905060600201803603810190610c0e9190613a1d565b90505f816040015167ffffffffffffffff811115610c2f57610c2e61306e565b5b604051908082528060200260200182016040528015610c6857816020015b610c55612939565b815260200190600190039081610c4d5790505b5090505f805f90505b89899050811015610cef575f610cc6855f01518c8c85818110610c9757610c96613847565b5b9050602002810190610ca99190613a48565b805f0190610cb79190613a6f565b611c1c9290919263ffffffff16565b9050610cdd818486611ddf9092919063ffffffff16565b80518301925081600101915050610c71565b5082868581518110610d0457610d03613847565b5b60200260200101515f018190525081868581518110610d2657610d25613847565b5b602002602001015160200181905250836001019350505050610bbf565b505f5a9050610d83835f81518110610d5e57610d5d613847565b5b602002602001015160200151856020016020810190610d7d9190612e54565b5f611e91565b5f5a82610d909190613afe565b90505f878785604051602001610da893929190614126565b604051602081830303815290604052905050505050505050505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610df2611f00565b610dfb5f611f87565b565b610e0561204a565b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ed7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ece906141b7565b60405180910390fd5b610ee381600180611e91565b50565b5f808473ffffffffffffffffffffffffffffffffffffffff168484604051610f0f929190614203565b5f60405180830381855af49150503d805f8114610f47576040519150601f19603f3d011682016040523d82523d5f602084013e610f4c565b606091505b509150915081816040517f99410554000000000000000000000000000000000000000000000000000000008152600401610f87929190614295565b60405180910390fd5b610f9861204a565b818160049190610fa99291906129d2565b505050565b610fb661204a565b8060068190555050565b5f60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f805a90503073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461105b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110529061430d565b60405180910390fd5b5f855f015190505f60648260800151611074919061432b565b90505f808951111561110d575f611090845f01515f8c86612054565b90508061110b575f6110a361080061206b565b90505f8151111561110557845f015173ffffffffffffffffffffffffffffffffffffffff168a602001517f0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb836040516110fc919061436c565b60405180910390a35b60019250505b505b5f88608001515a8603019050611168828a8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508461209a565b95505050505050949350505050565b61117f61204a565b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f548214611205576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111fc906143d6565b60405180910390fd5b5f801b8103611249576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112409061443e565b60405180910390fd5b805f819055505050565b81341461128c576040517f2a255cf200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546112d8919061445c565b925050819055506112e98282610a0f565b5050565b6112f561204a565b8060058190555050565b6001602052805f5260405f205f915090505481565b5f61131e826121dd565b805190602001209050919050565b611334611f00565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036113a4575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161139b9190612cda565b60405180910390fd5b6113ad81611f87565b50565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b600260075403611455576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600781905550565b5f5a90505f825f0151905061147484826122f2565b5f816060015190505f8260e001518360c001518460a00151856060015186608001511717171790506effffffffffffffffffffffffffffff80168111156114f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114e7906144d9565b60405180910390fd5b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff801683604001511115611558576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161154f90614541565b60405180910390fd5b5f611562846123bc565b90505f61156e886123e1565b90505f855f015173ffffffffffffffffffffffffffffffffffffffff16635c51ffd1838b6101600160208101906115a59190612e54565b6040518363ffffffff1660e01b81526004016115c19190612cda565b6020604051808303815f8887f11580156115dd573d5f803e3d5ffd5b50505050506040513d601f19601f820116820180604052508101906116029190614573565b90508061164657896040517f220266b600000000000000000000000000000000000000000000000000000000815260040161163d91906145e8565b60405180910390fd5b815a8803111561168d57896040517f220266b60000000000000000000000000000000000000000000000000000000081526004016116849190614684565b60405180910390fd5b6060838960400181815250506116a2816123ec565b896060018181525050825a8903018960800181815250505050505050505050505050565b5f600160ff16825f0160208101906116de91906146b0565b60ff1614806117045750600260ff16825f0160208101906116ff91906146b0565b60ff16145b9050919050565b5f600160ff16835f01602081019061172391906146b0565b60ff160361173957611734836123f5565b611763565b600260ff16835f01602081019061175091906146b0565b60ff16036117625761176183612530565b5b5b8160400151905092915050565b5f805a90505f611783846060015161267b565b90505f806040519050365f888060a0019061179e91906146db565b9150915060605f8260038111156117b457843591505b5063c1d1e79660e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916036118db575f8b8b602001516040516024016118199291906148a0565b60405160208183030381529060405263c1d1e79660e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090503073ffffffffffffffffffffffffffffffffffffffff166396899d3a828d8b604051602401611891939291906149d9565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050925050611953565b3073ffffffffffffffffffffffffffffffffffffffff166396899d3a85858d8b60405160240161190e9493929190614a4a565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505091505b60205f8351602085015f305af195505f51985084604052505050505080611aed575f3d806020036119885760205f803e5f5191505b507fdeaddead0000000000000000000000000000000000000000000000000000000081036119ed57876040517f220266b60000000000000000000000000000000000000000000000000000000081526004016119e49190614adb565b60405180910390fd5b7fdeadaa51000000000000000000000000000000000000000000000000000000008103611a59575f86608001515a86611a269190613afe565b611a30919061445c565b90505f87604001519050611a4388612685565b611a4f885f83856126d6565b8096505050611aeb565b855f01515f015173ffffffffffffffffffffffffffffffffffffffff1686602001517f676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e611aa761080061206b565b604051611ab4919061436c565b60405180910390a35f86608001515a86611ace9190613afe565b611ad8919061445c565b9050611ae7600288868461209a565b9550505b505b5050509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611b65576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b5c90614b51565b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff1682604051611b8a90614b92565b5f6040518083038185875af1925050503d805f8114611bc4576040519150601f19603f3d011682016040523d82523d5f602084013e611bc9565b606091505b5050905080611c0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c0490614bf0565b60405180910390fd5b505050565b6001600781905550565b60605f808585905067ffffffffffffffff811115611c3d57611c3c61306e565b5b604051908082528060200260200182016040528015611c7657816020015b611c63612939565b815260200190600190039081611c5b5790505b5090505f5b86869050811015611d2b5784878783818110611c9a57611c99613847565b5b9050602002810190611cac9190613880565b6080016020810190611cbe9190614c0e565b67ffffffffffffffff1603611d2057868682818110611ce057611cdf613847565b5b9050602002810190611cf29190613880565b611cfb90614c39565b828481518110611d0e57611d0d613847565b5b60200260200101819052508260010192505b806001019050611c7b565b505f8267ffffffffffffffff811115611d4757611d4661306e565b5b604051908082528060200260200182016040528015611d8057816020015b611d6d612939565b815260200190600190039081611d655790505b5090505f5b83811015611dd157828181518110611da057611d9f613847565b5b6020026020010151828281518110611dbb57611dba613847565b5b6020026020010181905250806001019050611d85565b508093505050509392505050565b815181611dec919061445c565b83511015611e2f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e2690614c95565b60405180910390fd5b5f5b8251811015611e8b57828181518110611e4d57611e4c613847565b5b6020026020010151848284611e62919061445c565b81518110611e7357611e72613847565b5b60200260200101819052508080600101915050611e31565b50505050565b3073ffffffffffffffffffffffffffffffffffffffff16633d939c4e8484846040518463ffffffff1660e01b8152600401611ece93929190614eb6565b5f604051808303815f87803b158015611ee5575f80fd5b505af1158015611ef7573d5f803e3d5ffd5b50505050505050565b611f08612739565b73ffffffffffffffffffffffffffffffffffffffff16611f26610fc0565b73ffffffffffffffffffffffffffffffffffffffff1614611f8557611f49612739565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401611f7c9190612cda565b60405180910390fd5b565b5f60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b612052611f00565b565b5f805f845160208601878987f19050949350505050565b60603d8281111561207a578290505b604051602082018101604052818152815f602083013e8092505050919050565b5f805a90505f80865f015190505f6120b182612740565b9050815f015192505a8403860195505f8260a0015183608001518460600151010190505f8960800151880390508082111561210c575f81830390505f6064600a83028161210157612100614ef2565b5b049050808a01995050505b505080860294505f88604001519050858110156121985760028081111561213657612135614f1f565b5b8a600281111561214957612148614f1f565b5b0361216b5780955061215a89612685565b612166895f888a6126d6565b612193565b7fdeadaa51000000000000000000000000000000000000000000000000000000005f5260205ffd5b6121d0565b5f8060028111156121ac576121ab614f1f565b5b8b60028111156121bf576121be614f1f565b5b1490506121ce8a82898b6126d6565b505b5050505050949350505050565b60605f60f467ffffffffffffffff8111156121fb576121fa61306e565b5b6040519080825280601f01601f19166020018201604052801561222d5781602001600182028036833780820191505090505b5090505f61223a84612763565b90505f8460400160208101906122509190612e54565b90505f61225c86612793565b90505f612277878060a0019061227291906146db565b6127dd565b90505f612283886127f3565b90505f88610100013590505f6122988a61280d565b90505f8a6101600160208101906122af9190612e54565b90508760208a01528660408a01528560608a01528460808a01528360a08a01528260c08a01528160e08a01528060f48a0152889950505050505050505050919050565b8160400160208101906123059190612e54565b815f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505081608001602081019061234e9190614c0e565b67ffffffffffffffff1681602001818152505081602001358160400181815250508161010001358160600181815250508161012001358160c00181815250508161014001358160e00181815250508160c001358160800181815250508160e001358160a00181815250505050565b5f8160e001518260a00151028260c00151836060015184608001510102019050919050565b5f6127109050919050565b5f819050919050565b5f8160400160208101906124099190612e54565b90505f826020013590508060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101561248a576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546124d69190613afe565b925050819055508173ffffffffffffffffffffffffffffffffffffffff167f9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d826040516125239190612e8e565b60405180910390a2505050565b5f8160400160208101906125449190612e54565b90505f8260200135905080471015612588576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff16826040516125ad90614b92565b5f6040518083038185875af1925050503d805f81146125e7576040519150601f19603f3d011682016040523d82523d5f602084013e6125ec565b606091505b5050905080612627576040517f3204506f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff167fe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f8360405161266d9190612e8e565b60405180910390a250505050565b6060819050919050565b805f01515f015173ffffffffffffffffffffffffffffffffffffffff1681602001517f1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff2860405160405180910390a350565b835f01515f015173ffffffffffffffffffffffffffffffffffffffff1684602001517ffb68fae769a73ec810729fb1cf34aa9faf18a18c0d69fb5388726c793abbf5aa85858560405161272b93929190614f4c565b60405180910390a350505050565b5f33905090565b5f46826020015114612756578160e0015161275c565b8160c001515b9050919050565b5f80825f01602081019061277791906146b0565b90505f836020013590508160f81b925080831792505050919050565b5f6127d68260600160208101906127aa9190614c0e565b67ffffffffffffffff168360800160208101906127c79190614c0e565b67ffffffffffffffff16612829565b9050919050565b5f60405182808583378082209250505092915050565b5f6128068260c001358360e00135612829565b9050919050565b5f612822826101200135836101400135612829565b9050919050565b5f6fffffffffffffffffffffffffffffffff801683111561287f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161287690614fcb565b60405180910390fd5b6fffffffffffffffffffffffffffffffff80168211156128d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128cb90615033565b60405180910390fd5b81608084901b175f1b905092915050565b6040518060a001604052806128f8612a85565b81526020015f80191681526020015f81526020015f81526020015f81525090565b604051806040016040528061292c612ad8565b8152602001606081525090565b6040518061018001604052805f60ff1681526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f81526020015f81526020015f81526020015f81526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b828054828255905f5260205f2090600701600890048101928215612a74579160200282015f5b83821115612a4257833563ffffffff1683826101000a81548163ffffffff021916908363ffffffff16021790555092602001926004016020816003010492830192600103026129f8565b8015612a725782816101000a81549063ffffffff0219169055600401602081600301049283019260010302612a42565b505b509050612a819190612af6565b5090565b6040518061010001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b60405180606001604052805f81526020015f81526020015f81525090565b5b80821115612b0d575f815f905550600101612af7565b5090565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b612b5681612b22565b8114612b60575f80fd5b50565b5f81359050612b7181612b4d565b92915050565b5f60208284031215612b8c57612b8b612b1a565b5b5f612b9984828501612b63565b91505092915050565b5f8115159050919050565b612bb681612ba2565b82525050565b5f602082019050612bcf5f830184612bad565b92915050565b5f819050919050565b612be781612bd5565b8114612bf1575f80fd5b50565b5f81359050612c0281612bde565b92915050565b5f60208284031215612c1d57612c1c612b1a565b5b5f612c2a84828501612bf4565b91505092915050565b5f63ffffffff82169050919050565b612c4b81612c33565b82525050565b5f602082019050612c645f830184612c42565b92915050565b5f819050919050565b612c7c81612c6a565b82525050565b5f602082019050612c955f830184612c73565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f612cc482612c9b565b9050919050565b612cd481612cba565b82525050565b5f602082019050612ced5f830184612ccb565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112612d1457612d13612cf3565b5b8235905067ffffffffffffffff811115612d3157612d30612cf7565b5b602083019150836020820283011115612d4d57612d4c612cfb565b5b9250929050565b5f612d5e82612c9b565b9050919050565b612d6e81612d54565b8114612d78575f80fd5b50565b5f81359050612d8981612d65565b92915050565b612d9881612ba2565b8114612da2575f80fd5b50565b5f81359050612db381612d8f565b92915050565b5f805f8060608587031215612dd157612dd0612b1a565b5b5f85013567ffffffffffffffff811115612dee57612ded612b1e565b5b612dfa87828801612cff565b94509450506020612e0d87828801612d7b565b9250506040612e1e87828801612da5565b91505092959194509250565b612e3381612cba565b8114612e3d575f80fd5b50565b5f81359050612e4e81612e2a565b92915050565b5f60208284031215612e6957612e68612b1a565b5b5f612e7684828501612e40565b91505092915050565b612e8881612bd5565b82525050565b5f602082019050612ea15f830184612e7f565b92915050565b5f8060408385031215612ebd57612ebc612b1a565b5b5f612eca85828601612bf4565b9250506020612edb85828601612bf4565b9150509250929050565b5f8083601f840112612efa57612ef9612cf3565b5b8235905067ffffffffffffffff811115612f1757612f16612cf7565b5b602083019150836001820283011115612f3357612f32612cfb565b5b9250929050565b5f8083601f840112612f4f57612f4e612cf3565b5b8235905067ffffffffffffffff811115612f6c57612f6b612cf7565b5b602083019150836020820283011115612f8857612f87612cfb565b5b9250929050565b5f80fd5b5f60408284031215612fa857612fa7612f8f565b5b81905092915050565b5f805f805f60608688031215612fca57612fc9612b1a565b5b5f86013567ffffffffffffffff811115612fe757612fe6612b1e565b5b612ff388828901612ee5565b9550955050602086013567ffffffffffffffff81111561301657613015612b1e565b5b61302288828901612f3a565b9350935050604086013567ffffffffffffffff81111561304557613044612b1e565b5b61305188828901612f93565b9150509295509295909350565b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6130a48261305e565b810181811067ffffffffffffffff821117156130c3576130c261306e565b5b80604052505050565b5f6130d5612b11565b90506130e1828261309b565b919050565b5f67ffffffffffffffff821115613100576130ff61306e565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f60ff82169050919050565b61312e81613119565b8114613138575f80fd5b50565b5f8135905061314981613125565b92915050565b5f67ffffffffffffffff82169050919050565b61316b8161314f565b8114613175575f80fd5b50565b5f8135905061318681613162565b92915050565b5f80fd5b5f67ffffffffffffffff8211156131aa576131a961306e565b5b6131b38261305e565b9050602081019050919050565b828183375f83830152505050565b5f6131e06131db84613190565b6130cc565b9050828152602081018484840111156131fc576131fb61318c565b5b6132078482856131c0565b509392505050565b5f82601f83011261322357613222612cf3565b5b81356132338482602086016131ce565b91505092915050565b5f610180828403121561325257613251613111565b5b61325d6101806130cc565b90505f61326c8482850161313b565b5f83015250602061327f84828501612bf4565b602083015250604061329384828501612e40565b60408301525060606132a784828501613178565b60608301525060806132bb84828501613178565b60808301525060a082013567ffffffffffffffff8111156132df576132de613115565b5b6132eb8482850161320f565b60a08301525060c06132ff84828501612bf4565b60c08301525060e061331384828501612bf4565b60e08301525061010061332884828501612bf4565b6101008301525061012061333e84828501612bf4565b6101208301525061014061335484828501612bf4565b6101408301525061016061336a84828501612e40565b6101608301525092915050565b5f613389613384846130e6565b6130cc565b905080838252602082019050602084028301858111156133ac576133ab612cfb565b5b835b818110156133f357803567ffffffffffffffff8111156133d1576133d0612cf3565b5b8086016133de898261323c565b855260208501945050506020810190506133ae565b5050509392505050565b5f82601f83011261341157613410612cf3565b5b8135613421848260208601613377565b91505092915050565b5f6020828403121561343f5761343e612b1a565b5b5f82013567ffffffffffffffff81111561345c5761345b612b1e565b5b613468848285016133fd565b91505092915050565b5f805f6040848603121561348857613487612b1a565b5b5f61349586828701612e40565b935050602084013567ffffffffffffffff8111156134b6576134b5612b1e565b5b6134c286828701612ee5565b92509250509250925092565b5f8083601f8401126134e3576134e2612cf3565b5b8235905067ffffffffffffffff811115613500576134ff612cf7565b5b60208301915083602082028301111561351c5761351b612cfb565b5b9250929050565b5f806020838503121561353957613538612b1a565b5b5f83013567ffffffffffffffff81111561355657613555612b1e565b5b613562858286016134ce565b92509250509250929050565b5f610100828403121561358457613583613111565b5b61358f6101006130cc565b90505f61359e84828501612e40565b5f8301525060206135b184828501612bf4565b60208301525060406135c584828501612bf4565b60408301525060606135d984828501612bf4565b60608301525060806135ed84828501612bf4565b60808301525060a061360184828501612bf4565b60a08301525060c061361584828501612bf4565b60c08301525060e061362984828501612bf4565b60e08301525092915050565b61363e81612c6a565b8114613648575f80fd5b50565b5f8135905061365981613635565b92915050565b5f610180828403121561367557613674613111565b5b61367f60a06130cc565b90505f61368e8482850161356e565b5f830152506101006136a28482850161364b565b6020830152506101206136b784828501612bf4565b6040830152506101406136cc84828501612bf4565b6060830152506101606136e184828501612bf4565b60808301525092915050565b5f805f806101c0858703121561370657613705612b1a565b5b5f85013567ffffffffffffffff81111561372357613722612b1e565b5b61372f8782880161320f565b94505060206137408782880161365f565b9350506101a085013567ffffffffffffffff81111561376257613761612b1e565b5b61376e87828801612ee5565b925092505092959194509250565b5f806040838503121561379257613791612b1a565b5b5f61379f8582860161364b565b92505060206137b08582860161364b565b9150509250929050565b5f61018082840312156137d0576137cf612f8f565b5b81905092915050565b5f602082840312156137ee576137ed612b1a565b5b5f82013567ffffffffffffffff81111561380b5761380a612b1e565b5b613817848285016137ba565b91505092915050565b5f6040820190506138335f830185612e7f565b6138406020830184612e7f565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f80fd5b5f80fd5b5f80fd5b5f823560016101800383360303811261389c5761389b613874565b5b80830191505092915050565b5f8160601b9050919050565b5f6138be826138a8565b9050919050565b5f6138cf826138b4565b9050919050565b6138e76138e282612cba565b6138c5565b82525050565b5f819050919050565b61390761390282612bd5565b6138ed565b82525050565b5f61391882876138d6565b60148201915061392882866138f6565b60208201915061393882856138f6565b60208201915061394882846138f6565b60208201915081905095945050505050565b5f808335600160200384360303811261397657613975613874565b5b80840192508235915067ffffffffffffffff82111561399857613997613878565b5b6020830192506060820236038313156139b4576139b361387c565b5b509250929050565b5f606082840312156139d1576139d0613111565b5b6139db60606130cc565b90505f6139ea84828501612bf4565b5f8301525060206139fd84828501612bf4565b6020830152506040613a1184828501612bf4565b60408301525092915050565b5f60608284031215613a3257613a31612b1a565b5b5f613a3f848285016139bc565b91505092915050565b5f82356001608003833603038112613a6357613a62613874565b5b80830191505092915050565b5f8083356001602003843603038112613a8b57613a8a613874565b5b80840192508235915067ffffffffffffffff821115613aad57613aac613878565b5b602083019250602082023603831315613ac957613ac861387c565b5b509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f613b0882612bd5565b9150613b1383612bd5565b9250828203905081811115613b2b57613b2a613ad1565b5b92915050565b5f82825260208201905092915050565b5f819050919050565b5f80fd5b5f80fd5b5f80fd5b5f8083356001602003843603038112613b7257613b71613b52565b5b83810192508235915060208301925067ffffffffffffffff821115613b9a57613b99613b4a565b5b602082023603831315613bb057613baf613b4e565b5b509250929050565b5f82825260208201905092915050565b5f819050919050565b5f613bdf602084018461313b565b905092915050565b613bf081613119565b82525050565b5f613c046020840184612bf4565b905092915050565b613c1581612bd5565b82525050565b5f613c296020840184612e40565b905092915050565b613c3a81612cba565b82525050565b5f613c4e6020840184613178565b905092915050565b613c5f8161314f565b82525050565b5f8083356001602003843603038112613c8157613c80613b52565b5b83810192508235915060208301925067ffffffffffffffff821115613ca957613ca8613b4a565b5b600182023603831315613cbf57613cbe613b4e565b5b509250929050565b5f82825260208201905092915050565b5f613ce28385613cc7565b9350613cef8385846131c0565b613cf88361305e565b840190509392505050565b5f6101808301613d155f840184613bd1565b613d215f860182613be7565b50613d2f6020840184613bf6565b613d3c6020860182613c0c565b50613d4a6040840184613c1b565b613d576040860182613c31565b50613d656060840184613c40565b613d726060860182613c56565b50613d806080840184613c40565b613d8d6080860182613c56565b50613d9b60a0840184613c65565b85830360a0870152613dae838284613cd7565b92505050613dbf60c0840184613bf6565b613dcc60c0860182613c0c565b50613dda60e0840184613bf6565b613de760e0860182613c0c565b50613df6610100840184613bf6565b613e04610100860182613c0c565b50613e13610120840184613bf6565b613e21610120860182613c0c565b50613e30610140840184613bf6565b613e3e610140860182613c0c565b50613e4d610160840184613c1b565b613e5b610160860182613c31565b508091505092915050565b5f613e718383613d03565b905092915050565b5f8235600161018003833603038112613e9557613e94613b52565b5b82810191505092915050565b5f602082019050919050565b5f613eb88385613bb8565b935083602084028501613eca84613bc8565b805f5b87811015613f0d578484038952613ee48284613e79565b613eee8582613e66565b9450613ef983613ea1565b925060208a01995050600181019050613ecd565b50829750879450505050509392505050565b5f613f2d602084018461364b565b905092915050565b613f3e81612c6a565b82525050565b5f60808301613f555f840184613b56565b8583035f870152613f67838284613ead565b92505050613f786020840184613f1f565b613f856020860182613f35565b50613f936040840184613f1f565b613fa06040860182613f35565b50613fae6060840184613f1f565b613fbb6060860182613f35565b508091505092915050565b5f613fd18383613f44565b905092915050565b5f82356001608003833603038112613ff457613ff3613b52565b5b82810191505092915050565b5f602082019050919050565b5f6140178385613b31565b93508360208402850161402984613b41565b805f5b8781101561406c5784840389526140438284613fd9565b61404d8582613fc6565b945061405883614000565b925060208a0199505060018101905061402c565b50829750879450505050509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6140b28383613f35565b60208301905092915050565b5f602082019050919050565b5f6140d48261407e565b6140de8185614088565b93506140e983614098565b805f5b8381101561411957815161410088826140a7565b975061410b836140be565b9250506001810190506140ec565b5085935050505092915050565b5f6040820190508181035f83015261413f81858761400c565b9050818103602083015261415381846140ca565b9050949350505050565b5f82825260208201905092915050565b7f4e455153525200000000000000000000000000000000000000000000000000005f82015250565b5f6141a160068361415d565b91506141ac8261416d565b602082019050919050565b5f6020820190508181035f8301526141ce81614195565b9050919050565b5f81905092915050565b5f6141ea83856141d5565b93506141f78385846131c0565b82840190509392505050565b5f61420f8284866141df565b91508190509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015614252578082015181840152602081019050614237565b5f8484015250505050565b5f6142678261421b565b6142718185614225565b9350614281818560208601614235565b61428a8161305e565b840191505092915050565b5f6040820190506142a85f830185612bad565b81810360208301526142ba818461425d565b90509392505050565b7f4141393220696e7465726e616c2063616c6c206f6e6c790000000000000000005f82015250565b5f6142f760178361415d565b9150614302826142c3565b602082019050919050565b5f6020820190508181035f830152614324816142eb565b9050919050565b5f61433582612bd5565b915061434083612bd5565b925082820261434e81612bd5565b9150828204841483151761436557614364613ad1565b5b5092915050565b5f6020820190508181035f830152614384818461425d565b905092915050565b7f4e455153520000000000000000000000000000000000000000000000000000005f82015250565b5f6143c060058361415d565b91506143cb8261438c565b602082019050919050565b5f6020820190508181035f8301526143ed816143b4565b9050919050565b7f4e52495a000000000000000000000000000000000000000000000000000000005f82015250565b5f61442860048361415d565b9150614433826143f4565b602082019050919050565b5f6020820190508181035f8301526144558161441c565b9050919050565b5f61446682612bd5565b915061447183612bd5565b925082820190508082111561448957614488613ad1565b5b92915050565b7f41413934206761732076616c756573206f766572666c6f7700000000000000005f82015250565b5f6144c360188361415d565b91506144ce8261448f565b602082019050919050565b5f6020820190508181035f8301526144f0816144b7565b9050919050565b7f41413934206f7065726174696f6e2076616c7565206f766572666c6f770000005f82015250565b5f61452b601d8361415d565b9150614536826144f7565b602082019050919050565b5f6020820190508181035f8301526145588161451f565b9050919050565b5f8151905061456d81612d8f565b92915050565b5f6020828403121561458857614587612b1a565b5b5f6145958482850161455f565b91505092915050565b7f4141206f776e657220766572696669636174696f6e2066616c696564000000005f82015250565b5f6145d2601c8361415d565b91506145dd8261459e565b602082019050919050565b5f6040820190506145fb5f830184612e7f565b818103602083015261460c816145c6565b905092915050565b7f41413236206f76657220766572696669636174696f6e4f776e65724761734c695f8201527f6d69740000000000000000000000000000000000000000000000000000000000602082015250565b5f61466e60238361415d565b915061467982614614565b604082019050919050565b5f6040820190506146975f830184612e7f565b81810360208301526146a881614662565b905092915050565b5f602082840312156146c5576146c4612b1a565b5b5f6146d28482850161313b565b91505092915050565b5f80833560016020038436030381126146f7576146f6613874565b5b80840192508235915067ffffffffffffffff82111561471957614718613878565b5b6020830192506001820236038313156147355761473461387c565b5b509250929050565b5f610180830161474f5f840184613bd1565b61475b5f860182613be7565b506147696020840184613bf6565b6147766020860182613c0c565b506147846040840184613c1b565b6147916040860182613c31565b5061479f6060840184613c40565b6147ac6060860182613c56565b506147ba6080840184613c40565b6147c76080860182613c56565b506147d560a0840184613c65565b85830360a08701526147e8838284613cd7565b925050506147f960c0840184613bf6565b61480660c0860182613c0c565b5061481460e0840184613bf6565b61482160e0860182613c0c565b50614830610100840184613bf6565b61483e610100860182613c0c565b5061484d610120840184613bf6565b61485b610120860182613c0c565b5061486a610140840184613bf6565b614878610140860182613c0c565b50614887610160840184613c1b565b614895610160860182613c31565b508091505092915050565b5f6040820190508181035f8301526148b8818561473d565b90506148c76020830184612c73565b9392505050565b61010082015f8201516148e35f850182613c31565b5060208201516148f66020850182613c0c565b5060408201516149096040850182613c0c565b50606082015161491c6060850182613c0c565b50608082015161492f6080850182613c0c565b5060a082015161494260a0850182613c0c565b5060c082015161495560c0850182613c0c565b5060e082015161496860e0850182613c0c565b50505050565b61018082015f8201516149835f8501826148ce565b506020820151614997610100850182613f35565b5060408201516149ab610120850182613c0c565b5060608201516149bf610140850182613c0c565b5060808201516149d3610160850182613c0c565b50505050565b5f6101c0820190508181035f8301526149f2818661425d565b9050614a01602083018561496e565b8181036101a0830152614a14818461425d565b9050949350505050565b5f614a298385614225565b9350614a368385846131c0565b614a3f8361305e565b840190509392505050565b5f6101c0820190508181035f830152614a64818688614a1e565b9050614a73602083018561496e565b8181036101a0830152614a86818461425d565b905095945050505050565b7f41413935206f7574206f662067617300000000000000000000000000000000005f82015250565b5f614ac5600f8361415d565b9150614ad082614a91565b602082019050919050565b5f604082019050614aee5f830184612e7f565b8181036020830152614aff81614ab9565b905092915050565b7f4141393020696e76616c69642062656e656669636961727900000000000000005f82015250565b5f614b3b60188361415d565b9150614b4682614b07565b602082019050919050565b5f6020820190508181035f830152614b6881614b2f565b9050919050565b50565b5f614b7d5f836141d5565b9150614b8882614b6f565b5f82019050919050565b5f614b9c82614b72565b9150819050919050565b7f41413931206661696c65642073656e6420746f2062656e6566696369617279005f82015250565b5f614bda601f8361415d565b9150614be582614ba6565b602082019050919050565b5f6020820190508181035f830152614c0781614bce565b9050919050565b5f60208284031215614c2357614c22612b1a565b5b5f614c3084828501613178565b91505092915050565b5f614c44368361323c565b9050919050565b7f54617267657420617272617920746f6f20736d616c6c000000000000000000005f82015250565b5f614c7f60168361415d565b9150614c8a82614c4b565b602082019050919050565b5f6020820190508181035f830152614cac81614c73565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f614ce68261421b565b614cf08185613cc7565b9350614d00818560208601614235565b614d098161305e565b840191505092915050565b5f61018083015f830151614d2a5f860182613be7565b506020830151614d3d6020860182613c0c565b506040830151614d506040860182613c31565b506060830151614d636060860182613c56565b506080830151614d766080860182613c56565b5060a083015184820360a0860152614d8e8282614cdc565b91505060c0830151614da360c0860182613c0c565b5060e0830151614db660e0860182613c0c565b50610100830151614dcb610100860182613c0c565b50610120830151614de0610120860182613c0c565b50610140830151614df5610140860182613c0c565b50610160830151614e0a610160860182613c31565b508091505092915050565b5f614e208383614d14565b905092915050565b5f602082019050919050565b5f614e3e82614cb3565b614e488185614cbd565b935083602082028501614e5a85614ccd565b805f5b85811015614e955784840389528151614e768582614e15565b9450614e8183614e28565b925060208a01995050600181019050614e5d565b50829750879550505050505092915050565b614eb081612d54565b82525050565b5f6060820190508181035f830152614ece8186614e34565b9050614edd6020830185614ea7565b614eea6040830184612bad565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f606082019050614f5f5f830186612bad565b614f6c6020830185612e7f565b614f796040830184612e7f565b949350505050565b7f6869676831323820657863656564732075696e743132382072616e67650000005f82015250565b5f614fb5601d8361415d565b9150614fc082614f81565b602082019050919050565b5f6020820190508181035f830152614fe281614fa9565b9050919050565b7f6c6f7731323820657863656564732075696e743132382072616e6765000000005f82015250565b5f61501d601c8361415d565b915061502882614fe9565b602082019050919050565b5f6020820190508181035f83015261504a81615011565b905091905056fea26469706673582212204859403cfa9e98ce7491486ae238bd93139aa2172e2988dd43e284f951d1508364736f6c63430008180033",
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

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_EntryPoint *EntryPointCaller) GetPreGasBalanceInfo(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getPreGasBalanceInfo", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_EntryPoint *EntryPointSession) GetPreGasBalanceInfo(account common.Address) (*big.Int, error) {
	return _EntryPoint.Contract.GetPreGasBalanceInfo(&_EntryPoint.CallOpts, account)
}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_EntryPoint *EntryPointCallerSession) GetPreGasBalanceInfo(account common.Address) (*big.Int, error) {
	return _EntryPoint.Contract.GetPreGasBalanceInfo(&_EntryPoint.CallOpts, account)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0xc989e90c.
//
// Solidity: function getUserOpHash((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointCaller) GetUserOpHash(opts *bind.CallOpts, userOp PackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0xc989e90c.
//
// Solidity: function getUserOpHash((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointSession) GetUserOpHash(userOp PackedUserOperation) ([32]byte, error) {
	return _EntryPoint.Contract.GetUserOpHash(&_EntryPoint.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0xc989e90c.
//
// Solidity: function getUserOpHash((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address) userOp) view returns(bytes32)
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

// PreGasBalance is a free data retrieval call binding the contract method 0xa6afc1de.
//
// Solidity: function preGasBalance(address account) view returns(uint256 amount)
func (_EntryPoint *EntryPointCaller) PreGasBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "preGasBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreGasBalance is a free data retrieval call binding the contract method 0xa6afc1de.
//
// Solidity: function preGasBalance(address account) view returns(uint256 amount)
func (_EntryPoint *EntryPointSession) PreGasBalance(account common.Address) (*big.Int, error) {
	return _EntryPoint.Contract.PreGasBalance(&_EntryPoint.CallOpts, account)
}

// PreGasBalance is a free data retrieval call binding the contract method 0xa6afc1de.
//
// Solidity: function preGasBalance(address account) view returns(uint256 amount)
func (_EntryPoint *EntryPointCallerSession) PreGasBalance(account common.Address) (*big.Int, error) {
	return _EntryPoint.Contract.PreGasBalance(&_EntryPoint.CallOpts, account)
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

// HandleOps is a paid mutator transaction binding the contract method 0x3d939c4e.
//
// Solidity: function handleOps((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address)[] ops, address beneficiary, bool isSync) returns()
func (_EntryPoint *EntryPointTransactor) HandleOps(opts *bind.TransactOpts, ops []PackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "handleOps", ops, beneficiary, isSync)
}

// HandleOps is a paid mutator transaction binding the contract method 0x3d939c4e.
//
// Solidity: function handleOps((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address)[] ops, address beneficiary, bool isSync) returns()
func (_EntryPoint *EntryPointSession) HandleOps(ops []PackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, beneficiary, isSync)
}

// HandleOps is a paid mutator transaction binding the contract method 0x3d939c4e.
//
// Solidity: function handleOps((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address)[] ops, address beneficiary, bool isSync) returns()
func (_EntryPoint *EntryPointTransactorSession) HandleOps(ops []PackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, beneficiary, isSync)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x96899d3a.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactor) InnerHandleOp(opts *bind.TransactOpts, callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "innerHandleOp", callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x96899d3a.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x96899d3a.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactorSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// RedeemGasOperation is a paid mutator transaction binding the contract method 0x654e8b23.
//
// Solidity: function redeemGasOperation(uint256 amount, uint256 nonce) returns()
func (_EntryPoint *EntryPointTransactor) RedeemGasOperation(opts *bind.TransactOpts, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "redeemGasOperation", amount, nonce)
}

// RedeemGasOperation is a paid mutator transaction binding the contract method 0x654e8b23.
//
// Solidity: function redeemGasOperation(uint256 amount, uint256 nonce) returns()
func (_EntryPoint *EntryPointSession) RedeemGasOperation(amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.RedeemGasOperation(&_EntryPoint.TransactOpts, amount, nonce)
}

// RedeemGasOperation is a paid mutator transaction binding the contract method 0x654e8b23.
//
// Solidity: function redeemGasOperation(uint256 amount, uint256 nonce) returns()
func (_EntryPoint *EntryPointTransactorSession) RedeemGasOperation(amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.RedeemGasOperation(&_EntryPoint.TransactOpts, amount, nonce)
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

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_EntryPoint *EntryPointTransactor) SubmitDepositOperation(opts *bind.TransactOpts, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "submitDepositOperation", amount, nonce)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_EntryPoint *EntryPointSession) SubmitDepositOperation(amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.SubmitDepositOperation(&_EntryPoint.TransactOpts, amount, nonce)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_EntryPoint *EntryPointTransactorSession) SubmitDepositOperation(amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.SubmitDepositOperation(&_EntryPoint.TransactOpts, amount, nonce)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_EntryPoint *EntryPointTransactor) SubmitWithdrawOperation(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "submitWithdrawOperation", amount)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_EntryPoint *EntryPointSession) SubmitWithdrawOperation(amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.SubmitWithdrawOperation(&_EntryPoint.TransactOpts, amount)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_EntryPoint *EntryPointTransactorSession) SubmitWithdrawOperation(amount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.SubmitWithdrawOperation(&_EntryPoint.TransactOpts, amount)
}

// SyncBatch is a paid mutator transaction binding the contract method 0x84b2b1b1.
//
// Solidity: function syncBatch((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address)[] userOps) returns()
func (_EntryPoint *EntryPointTransactor) SyncBatch(opts *bind.TransactOpts, userOps []PackedUserOperation) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "syncBatch", userOps)
}

// SyncBatch is a paid mutator transaction binding the contract method 0x84b2b1b1.
//
// Solidity: function syncBatch((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address)[] userOps) returns()
func (_EntryPoint *EntryPointSession) SyncBatch(userOps []PackedUserOperation) (*types.Transaction, error) {
	return _EntryPoint.Contract.SyncBatch(&_EntryPoint.TransactOpts, userOps)
}

// SyncBatch is a paid mutator transaction binding the contract method 0x84b2b1b1.
//
// Solidity: function syncBatch((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address)[] userOps) returns()
func (_EntryPoint *EntryPointTransactorSession) SyncBatch(userOps []PackedUserOperation) (*types.Transaction, error) {
	return _EntryPoint.Contract.SyncBatch(&_EntryPoint.TransactOpts, userOps)
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

// UpdateDstCoeffGas is a paid mutator transaction binding the contract method 0x9e998d4b.
//
// Solidity: function updateDstCoeffGas(uint256 _dstCoeffGas) returns()
func (_EntryPoint *EntryPointTransactor) UpdateDstCoeffGas(opts *bind.TransactOpts, _dstCoeffGas *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "updateDstCoeffGas", _dstCoeffGas)
}

// UpdateDstCoeffGas is a paid mutator transaction binding the contract method 0x9e998d4b.
//
// Solidity: function updateDstCoeffGas(uint256 _dstCoeffGas) returns()
func (_EntryPoint *EntryPointSession) UpdateDstCoeffGas(_dstCoeffGas *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateDstCoeffGas(&_EntryPoint.TransactOpts, _dstCoeffGas)
}

// UpdateDstCoeffGas is a paid mutator transaction binding the contract method 0x9e998d4b.
//
// Solidity: function updateDstCoeffGas(uint256 _dstCoeffGas) returns()
func (_EntryPoint *EntryPointTransactorSession) UpdateDstCoeffGas(_dstCoeffGas *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateDstCoeffGas(&_EntryPoint.TransactOpts, _dstCoeffGas)
}

// UpdateDstConGas is a paid mutator transaction binding the contract method 0x8cb53415.
//
// Solidity: function updateDstConGas(uint256 _dstConGas) returns()
func (_EntryPoint *EntryPointTransactor) UpdateDstConGas(opts *bind.TransactOpts, _dstConGas *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "updateDstConGas", _dstConGas)
}

// UpdateDstConGas is a paid mutator transaction binding the contract method 0x8cb53415.
//
// Solidity: function updateDstConGas(uint256 _dstConGas) returns()
func (_EntryPoint *EntryPointSession) UpdateDstConGas(_dstConGas *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateDstConGas(&_EntryPoint.TransactOpts, _dstConGas)
}

// UpdateDstConGas is a paid mutator transaction binding the contract method 0x8cb53415.
//
// Solidity: function updateDstConGas(uint256 _dstConGas) returns()
func (_EntryPoint *EntryPointTransactorSession) UpdateDstConGas(_dstConGas *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateDstConGas(&_EntryPoint.TransactOpts, _dstConGas)
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

// UpdateSmtRoot is a paid mutator transaction binding the contract method 0x98824908.
//
// Solidity: function updateSmtRoot(bytes32 oldSmtRoot, bytes32 newSmtRoot) returns()
func (_EntryPoint *EntryPointTransactor) UpdateSmtRoot(opts *bind.TransactOpts, oldSmtRoot [32]byte, newSmtRoot [32]byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "updateSmtRoot", oldSmtRoot, newSmtRoot)
}

// UpdateSmtRoot is a paid mutator transaction binding the contract method 0x98824908.
//
// Solidity: function updateSmtRoot(bytes32 oldSmtRoot, bytes32 newSmtRoot) returns()
func (_EntryPoint *EntryPointSession) UpdateSmtRoot(oldSmtRoot [32]byte, newSmtRoot [32]byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateSmtRoot(&_EntryPoint.TransactOpts, oldSmtRoot, newSmtRoot)
}

// UpdateSmtRoot is a paid mutator transaction binding the contract method 0x98824908.
//
// Solidity: function updateSmtRoot(bytes32 oldSmtRoot, bytes32 newSmtRoot) returns()
func (_EntryPoint *EntryPointTransactorSession) UpdateSmtRoot(oldSmtRoot [32]byte, newSmtRoot [32]byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateSmtRoot(&_EntryPoint.TransactOpts, oldSmtRoot, newSmtRoot)
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

// VerifyBatch is a paid mutator transaction binding the contract method 0x671d6c35.
//
// Solidity: function verifyBatch(bytes proof, ((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address)[],bytes32,bytes32,bytes32)[] batches, ((uint256,uint256,uint256)[],address) chainsExecuteInfos) payable returns()
func (_EntryPoint *EntryPointTransactor) VerifyBatch(opts *bind.TransactOpts, proof []byte, batches []IEntryPointBatchData, chainsExecuteInfos IEntryPointChainsExecuteInfo) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "verifyBatch", proof, batches, chainsExecuteInfos)
}

// VerifyBatch is a paid mutator transaction binding the contract method 0x671d6c35.
//
// Solidity: function verifyBatch(bytes proof, ((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address)[],bytes32,bytes32,bytes32)[] batches, ((uint256,uint256,uint256)[],address) chainsExecuteInfos) payable returns()
func (_EntryPoint *EntryPointSession) VerifyBatch(proof []byte, batches []IEntryPointBatchData, chainsExecuteInfos IEntryPointChainsExecuteInfo) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatch(&_EntryPoint.TransactOpts, proof, batches, chainsExecuteInfos)
}

// VerifyBatch is a paid mutator transaction binding the contract method 0x671d6c35.
//
// Solidity: function verifyBatch(bytes proof, ((uint8,uint256,address,uint64,uint64,bytes,uint256,uint256,uint256,uint256,uint256,address)[],bytes32,bytes32,bytes32)[] batches, ((uint256,uint256,uint256)[],address) chainsExecuteInfos) payable returns()
func (_EntryPoint *EntryPointTransactorSession) VerifyBatch(proof []byte, batches []IEntryPointBatchData, chainsExecuteInfos IEntryPointChainsExecuteInfo) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatch(&_EntryPoint.TransactOpts, proof, batches, chainsExecuteInfos)
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
	Did       [32]byte
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketAdded is a free log retrieval operation binding the contract event 0x081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc.
//
// Solidity: event DepositTicketAdded(bytes32 indexed did, address indexed account, uint256 amount, uint256 timestamp)
func (_EntryPoint *EntryPointFilterer) FilterDepositTicketAdded(opts *bind.FilterOpts, did [][32]byte, account []common.Address) (*EntryPointDepositTicketAddedIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "DepositTicketAdded", didRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointDepositTicketAddedIterator{contract: _EntryPoint.contract, event: "DepositTicketAdded", logs: logs, sub: sub}, nil
}

// WatchDepositTicketAdded is a free log subscription operation binding the contract event 0x081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc.
//
// Solidity: event DepositTicketAdded(bytes32 indexed did, address indexed account, uint256 amount, uint256 timestamp)
func (_EntryPoint *EntryPointFilterer) WatchDepositTicketAdded(opts *bind.WatchOpts, sink chan<- *EntryPointDepositTicketAdded, did [][32]byte, account []common.Address) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "DepositTicketAdded", didRule, accountRule)
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

// ParseDepositTicketAdded is a log parse operation binding the contract event 0x081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc.
//
// Solidity: event DepositTicketAdded(bytes32 indexed did, address indexed account, uint256 amount, uint256 timestamp)
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
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketDeleted is a free log retrieval operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_EntryPoint *EntryPointFilterer) FilterDepositTicketDeleted(opts *bind.FilterOpts, account []common.Address) (*EntryPointDepositTicketDeletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "DepositTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointDepositTicketDeletedIterator{contract: _EntryPoint.contract, event: "DepositTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchDepositTicketDeleted is a free log subscription operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_EntryPoint *EntryPointFilterer) WatchDepositTicketDeleted(opts *bind.WatchOpts, sink chan<- *EntryPointDepositTicketDeleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "DepositTicketDeleted", accountRule)
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

// ParseDepositTicketDeleted is a log parse operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_EntryPoint *EntryPointFilterer) ParseDepositTicketDeleted(log types.Log) (*EntryPointDepositTicketDeleted, error) {
	event := new(EntryPointDepositTicketDeleted)
	if err := _EntryPoint.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
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
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserOperationEvent is a free log retrieval operation binding the contract event 0xfb68fae769a73ec810729fb1cf34aa9faf18a18c0d69fb5388726c793abbf5aa.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_EntryPoint *EntryPointFilterer) FilterUserOperationEvent(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointUserOperationEventIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "UserOperationEvent", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointUserOperationEventIterator{contract: _EntryPoint.contract, event: "UserOperationEvent", logs: logs, sub: sub}, nil
}

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0xfb68fae769a73ec810729fb1cf34aa9faf18a18c0d69fb5388726c793abbf5aa.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_EntryPoint *EntryPointFilterer) WatchUserOperationEvent(opts *bind.WatchOpts, sink chan<- *EntryPointUserOperationEvent, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "UserOperationEvent", userOpHashRule, senderRule)
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

// ParseUserOperationEvent is a log parse operation binding the contract event 0xfb68fae769a73ec810729fb1cf34aa9faf18a18c0d69fb5388726c793abbf5aa.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, bool success, uint256 actualGasCost, uint256 actualGasUsed)
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
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketAdded is a free log retrieval operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_EntryPoint *EntryPointFilterer) FilterWithdrawTicketAdded(opts *bind.FilterOpts, account []common.Address) (*EntryPointWithdrawTicketAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "WithdrawTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointWithdrawTicketAddedIterator{contract: _EntryPoint.contract, event: "WithdrawTicketAdded", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketAdded is a free log subscription operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_EntryPoint *EntryPointFilterer) WatchWithdrawTicketAdded(opts *bind.WatchOpts, sink chan<- *EntryPointWithdrawTicketAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "WithdrawTicketAdded", accountRule)
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

// ParseWithdrawTicketAdded is a log parse operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
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
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketDeleted is a free log retrieval operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_EntryPoint *EntryPointFilterer) FilterWithdrawTicketDeleted(opts *bind.FilterOpts, account []common.Address) (*EntryPointWithdrawTicketDeletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "WithdrawTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointWithdrawTicketDeletedIterator{contract: _EntryPoint.contract, event: "WithdrawTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketDeleted is a free log subscription operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_EntryPoint *EntryPointFilterer) WatchWithdrawTicketDeleted(opts *bind.WatchOpts, sink chan<- *EntryPointWithdrawTicketDeleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "WithdrawTicketDeleted", accountRule)
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

// ParseWithdrawTicketDeleted is a log parse operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_EntryPoint *EntryPointFilterer) ParseWithdrawTicketDeleted(log types.Log) (*EntryPointWithdrawTicketDeleted, error) {
	event := new(EntryPointWithdrawTicketDeleted)
	if err := _EntryPoint.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
