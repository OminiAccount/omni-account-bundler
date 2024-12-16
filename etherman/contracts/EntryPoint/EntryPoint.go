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

// BaseStructBatchData is an auto generated low-level Go binding around an user-defined struct.
type BaseStructBatchData struct {
	UserOperations []BaseStructPackedUserOperation
	AccInputHash   [32]byte
}

// BaseStructChainExecuteExtra is an auto generated low-level Go binding around an user-defined struct.
type BaseStructChainExecuteExtra struct {
	ChainId                   uint64
	ChainFee                  uint64
	ChainUserOperationsNumber uint64
}

// BaseStructChainsExecuteInfo is an auto generated low-level Go binding around an user-defined struct.
type BaseStructChainsExecuteInfo struct {
	ChainExtra   []BaseStructChainExecuteExtra
	NewStateRoot [32]byte
	Beneficiary  common.Address
}

// BaseStructMemoryUserOp is an auto generated low-level Go binding around an user-defined struct.
type BaseStructMemoryUserOp struct {
	Sender                 common.Address
	ChainId                *big.Int
	OperationValue         *big.Int
	ZkVerificationGasLimit uint64
	MainChainGasLimit      uint64
	DestChainGasLimit      uint64
	MainChainGasPrice      *big.Int
	DestChainGasPrice      *big.Int
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

// BaseStructUserOpInfo is an auto generated low-level Go binding around an user-defined struct.
type BaseStructUserOpInfo struct {
	MUserOp       BaseStructMemoryUserOp
	UserOpHash    [32]byte
	Prefund       *big.Int
	ContextOffset *big.Int
	PreOpGas      *big.Int
}

// ConfigManagerConfig is an auto generated low-level Go binding around an user-defined struct.
type ConfigManagerConfig struct {
	EntryPoint common.Address
}

// EntryPointMetaData contains all meta data concerning the EntryPoint contract.
var EntryPointMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"name\":\"DelegateAndRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"inner\",\"type\":\"bytes\"}],\"name\":\"FailedOpWithRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NumberIsNotEqual\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"PostOpReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotEqual\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"did\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FORK_ID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchNumToState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accInputRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"chainConfigs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"oldAccInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPreGasBalanceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSync\",\"type\":\"bool\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"mainChainGasPrice\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"destChainGasPrice\",\"type\":\"uint128\"}],\"internalType\":\"structBaseStruct.MemoryUserOp\",\"name\":\"mUserOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contextOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"}],\"internalType\":\"structBaseStruct.UserOpInfo\",\"name\":\"opInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"innerHandleOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"preGasBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"redeemGasOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitDepositOperation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submitWithdrawOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"}],\"name\":\"syncBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"_chainIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"}],\"internalType\":\"structConfigManager.Config[]\",\"name\":\"_config\",\"type\":\"tuple[]\"}],\"name\":\"updateChainConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_syncRouter\",\"type\":\"address\"}],\"name\":\"updateSyncRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"}],\"internalType\":\"structBaseStruct.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainUserOperationsNumber\",\"type\":\"uint64\"}],\"internalType\":\"structBaseStruct.ChainExecuteExtra[]\",\"name\":\"chainExtra\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.ChainsExecuteInfo\",\"name\":\"chainsExecuteInfos\",\"type\":\"tuple\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b503360016005819055505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200008d575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000849190620001ab565b60405180910390fd5b6200009e81620000a560201b60201c565b50620001c6565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620001938262000168565b9050919050565b620001a58162000187565b82525050565b5f602082019050620001c05f8301846200019a565b92915050565b615aae80620001d45f395ff3fe608060405260043610610165575f3560e01c80637fcb3653116100d0578063a4c4499911610089578063c017d13911610063578063c017d13914610508578063e2aeaa3514610544578063e851609214610580578063f2fde38b146105aa57610165565b8063a4c449991461047c578063a6afc1de146104a4578063ba64b242146104e057610165565b80637fcb3653146103945780637fcb5227146103be578063850aaf62146103e65780638da5cb5b1461040e57806397fc007c146104385780639abe79031461046057610165565b8063654e8b2311610122578063654e8b23146102975780636b086237146102bf5780636cbe19af146102fc5780636e89195914610326578063715018a6146103425780637b031fbc1461035857610165565b806301ffc9a7146101695780632b4280f5146101a55780632b7ac3f3146101cd57806339875d1a146101f75780634e425f761461021f578063653e75ec1461025b575b5f80fd5b348015610174575f80fd5b5061018f600480360381019061018a91906132e6565b6105d2565b60405161019c919061332b565b60405180910390f35b3480156101b0575f80fd5b506101cb60048036038101906101c69190613429565b61073d565b005b3480156101d8575f80fd5b506101e1610940565b6040516101ee91906134ba565b60405180910390f35b348015610202575f80fd5b5061021d60048036038101906102189190613506565b610965565b005b34801561022a575f80fd5b506102456004803603810190610240919061356e565b6109b8565b60405161025291906134ba565b60405180910390f35b348015610266575f80fd5b50610281600480360381019061027c91906135c3565b6109f1565b60405161028e91906135fd565b60405180910390f35b3480156102a2575f80fd5b506102bd60048036038101906102b89190613616565b610a36565b005b3480156102ca575f80fd5b506102e560048036038101906102e0919061356e565b610b2d565b6040516102f392919061366c565b60405180910390f35b348015610307575f80fd5b50610310610b4d565b60405161031d91906134ba565b60405180910390f35b610340600480360381019061033b919061375f565b610b73565b005b34801561034d575f80fd5b50610356611216565b005b348015610363575f80fd5b5061037e6004803603810190610379919061382b565b611229565b60405161038b9190613872565b60405180910390f35b34801561039f575f80fd5b506103a8611241565b6040516103b5919061389a565b60405180910390f35b3480156103c9575f80fd5b506103e460048036038101906103df91906135c3565b61125a565b005b3480156103f1575f80fd5b5061040c600480360381019061040791906138b3565b6112a6565b005b348015610419575f80fd5b50610422611350565b60405161042f91906134ba565b60405180910390f35b348015610443575f80fd5b5061045e600480360381019061045991906135c3565b611378565b005b61047a60048036038101906104759190613616565b6113c3565b005b348015610487575f80fd5b506104a2600480360381019061049d91906139ba565b61145c565b005b3480156104af575f80fd5b506104ca60048036038101906104c591906135c3565b611536565b6040516104d791906135fd565b60405180910390f35b3480156104eb575f80fd5b5061050660048036038101906105019190613dc7565b61154a565b005b348015610513575f80fd5b5061052e60048036038101906105299190613fd2565b6115e9565b60405161053b91906135fd565b60405180910390f35b34801561054f575f80fd5b5061056a60048036038101906105659190614061565b611782565b6040516105779190614164565b60405180910390f35b34801561058b575f80fd5b50610594611866565b6040516105a1919061389a565b60405180910390f35b3480156105b5575f80fd5b506105d060048036038101906105cb91906135c3565b61186b565b005b5f7fa349dad6000000000000000000000000000000000000000000000000000000007f6f2f826900000000000000000000000000000000000000000000000000000000187bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806106be57507f6f2f8269000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061072657507fa349dad6000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806107365750610735826118ef565b5b9050919050565b610745611958565b5f8484905090505f8167ffffffffffffffff81111561076757610766613a48565b5b6040519080825280602002602001820160405280156107a057816020015b61078d613090565b8152602001906001900390816107855790505b5090505f5b82811015610808575f8282815181106107c1576107c0614184565b5b602002602001015190506107fa828989858181106107e2576107e1614184565b5b90506020028101906107f491906141bd565b8361199e565b5080806001019150506107a5565b505f7fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f97260405160405180910390a15f5b8381101561091e5761086d88888381811061085657610855614184565b5b905060200281019061086891906141bd565b611c2b565b156108c3576108ba88888381811061088857610887614184565b5b905060200281019061089a91906141bd565b8483815181106108ad576108ac614184565b5b6020026020010151611c70565b82019150610911565b61090c818989848181106108da576108d9614184565b5b90506020028101906108ec91906141bd565b8584815181106108ff576108fe614184565b5b6020026020010151611cd5565b820191505b8080600101915050610838565b508361092f5761092e858261205c565b5b50505061093a612177565b50505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff167f2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a82426040516109ad9291906141e5565b60405180910390a250565b6004602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b815f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610aac576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1633468385604051602001610ada9493929190614271565b604051602081830303815290604052805190602001207f081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc8442604051610b219291906141e5565b60405180910390a35050565b6001602052805f5260405f205f91509050805f0154908060010154905082565b600260089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8787810190610b85919061441a565b9250925092505f610c9060025f9054906101000a900467ffffffffffffffff168888905060025f9054906101000a900467ffffffffffffffff16610bc99190614498565b60015f60025f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20600101548a8a60018d8d9050610c1991906144d3565b67ffffffffffffffff16818110610c3357610c32614184565b5b9050602002810190610c45919061450e565b6020013560015f60025f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f01548a60200135611782565b905060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166343753b4d85858560405180602001604052807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001600289604051610d0e919061456f565b602060405180830381855afa158015610d29573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610d4c9190614599565b5f1c610d5891906145f1565b8152506040518563ffffffff1660e01b8152600401610d7a949392919061483b565b602060405180830381865afa158015610d95573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610db99190614894565b610def576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f85805f0190610dff91906148bf565b905067ffffffffffffffff811115610e1a57610e19613a48565b5b604051908082528060200260200182016040528015610e5357816020015b610e406130c4565b815260200190600190039081610e385790505b5090505f5b86805f0190610e6791906148bf565b90508110156110c2575f87805f0190610e8091906148bf565b83818110610e9157610e90614184565b5b905060600201803603810190610ea79190614982565b905046815f015167ffffffffffffffff1614158015610f3a57505f73ffffffffffffffffffffffffffffffffffffffff1660045f835f015167ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b15610f71576040517fdece177600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f816040015167ffffffffffffffff1667ffffffffffffffff811115610f9a57610f99613a48565b5b604051908082528060200260200182016040528015610fd357816020015b610fc06130e4565b815260200190600190039081610fb85790505b5090505f805f90505b8c8c905067ffffffffffffffff1681101561106e575f611045855f015167ffffffffffffffff168f8f8581811061101657611015614184565b5b9050602002810190611028919061450e565b805f019061103691906149ad565b6121819290919263ffffffff16565b905061105c8184866123449092919063ffffffff16565b80518301925081600101915050610fdc565b508285858151811061108357611082614184565b5b60200260200101515f0181905250818585815181106110a5576110a4614184565b5b602002602001015160200181905250836001019350505050610e58565b506111368888905060025f9054906101000a900467ffffffffffffffff166110ea9190614498565b87602001358a8a60018d8d905061110191906144d3565b67ffffffffffffffff1681811061111b5761111a614184565b5b905060200281019061112d919061450e565b602001356123f6565b6111428888905061244f565b61117d815f8151811061115857611157614184565b5b60200260200101516020015187604001602081019061117791906135c3565b5f61249c565b5f600190505b8151811015611209575f8282815181106111a05761119f614184565b5b602002602001015190505f60045f835f01515f015167ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905050508080600101915050611183565b5050505050505050505050565b61121e61250b565b6112275f612592565b565b5f61123382612655565b805190602001209050919050565b60025f9054906101000a900467ffffffffffffffff1681565b611262612782565b80600260086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f808473ffffffffffffffffffffffffffffffffffffffff1684846040516112cf929190614a33565b5f60405180830381855af49150503d805f8114611307576040519150601f19603f3d011682016040523d82523d5f602084013e61130c565b606091505b509150915081816040517f99410554000000000000000000000000000000000000000000000000000000008152600401611347929190614a4b565b60405180910390fd5b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611380612782565b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8134146113fc576040517f2a255cf200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546114479190614a79565b925050819055506114588282610a36565b5050565b611464612782565b8181905084849050146114a3576040517f77fdf30d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8484905081101561152f578282828181106114c3576114c2614184565b5b90506020020160045f8787858181106114df576114de614184565b5b90506020020160208101906114f4919061356e565b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2081816115219190614b9c565b9050508060010190506114a5565b5050505050565b5f602052805f5260405f205f915090505481565b600260089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146115da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115d190614c04565b60405180910390fd5b6115e68160018061249c565b50565b5f805a90503073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461165c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161165390614c6c565b60405180910390fd5b5f855f015190505f606482608001516116759190614c8a565b67ffffffffffffffff1690505f8089511115611718575f61169b845f01515f8c8661278c565b905080611716575f6116ae6108006127a3565b90505f8151111561171057845f015173ffffffffffffffffffffffffffffffffffffffff168a602001517f0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb836040516117079190614164565b60405180910390a35b60019250505b505b5f88608001515a8603019050611773828a8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050846127d2565b95505050505050949350505050565b60605f8767ffffffffffffffff161415801561179f57505f801b85145b156117d6576040517f6818c29e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f801b8403611811576040517f66385b5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73f39fd6e51aad88f6f4ce6ab8827279cfffb92266838689600180878a5f801b8e60405160200161184b9a99989796959493929190614d1a565b60405160208183030381529060405290509695505050505050565b600181565b61187361250b565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036118e3575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016118da91906134ba565b60405180910390fd5b6118ec81612592565b50565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b600260055403611994576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600581905550565b5f5a90505f825f015190506119b3848261291f565b5f816060015167ffffffffffffffff1690505f8260e001518360c001518460a0015185606001518660800151171767ffffffffffffffff1617176fffffffffffffffffffffffffffffffff1690506effffffffffffffffffffffffffffff8016811115611a55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a4c90614e17565b60405180910390fd5b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff801683604001511115611abd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ab490614e7f565b60405180910390fd5b5f611ac784612ad1565b90505f611ad388612b1c565b90505f855f015173ffffffffffffffffffffffffffffffffffffffff16635c51ffd1838b610160016020810190611b0a91906135c3565b6040518363ffffffff1660e01b8152600401611b2691906134ba565b6020604051808303815f8887f1158015611b42573d5f803e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190611b679190614894565b905080611bab57896040517f220266b6000000000000000000000000000000000000000000000000000000008152600401611ba29190614ee7565b60405180910390fd5b815a88031115611bf257896040517f220266b6000000000000000000000000000000000000000000000000000000008152600401611be99190614f83565b60405180910390fd5b606083896040018181525050611c0781612b27565b896060018181525050825a8903018960800181815250505050505050505050505050565b5f600160ff16825f016020810190611c439190614faf565b60ff161480611c695750600260ff16825f016020810190611c649190614faf565b60ff16145b9050919050565b5f600160ff16835f016020810190611c889190614faf565b60ff1603611c9e57611c9983612b30565b611cc8565b600260ff16835f016020810190611cb59190614faf565b60ff1603611cc757611cc683612c69565b5b5b8160400151905092915050565b5f805a90505f611ce88460600151612db4565b90505f806040519050365f888060a00190611d039190614fda565b9150915060605f826003811115611d1957843591505b506331aa7ade60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603611e40575f8b8b60200151604051602401611d7e9291906152ce565b6040516020818303038152906040526331aa7ade60e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090503073ffffffffffffffffffffffffffffffffffffffff1663c017d139828d8b604051602401611df693929190615425565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050925050611eb8565b3073ffffffffffffffffffffffffffffffffffffffff1663c017d13985858d8b604051602401611e739493929190615496565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505091505b60205f8351602085015f305af195505f51985084604052505050505080612052575f3d80602003611eed5760205f803e5f5191505b507fdeaddead000000000000000000000000000000000000000000000000000000008103611f5257876040517f220266b6000000000000000000000000000000000000000000000000000000008152600401611f499190615527565b60405180910390fd5b7fdeadaa51000000000000000000000000000000000000000000000000000000008103611fbe575f86608001515a86611f8b9190615553565b611f959190614a79565b90505f87604001519050611fa888612dbe565b611fb4885f8385612e0f565b8096505050612050565b855f01515f015173ffffffffffffffffffffffffffffffffffffffff1686602001517f676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e61200c6108006127a3565b6040516120199190614164565b60405180910390a35f86608001515a866120339190615553565b61203d9190614a79565b905061204c60028886846127d2565b9550505b505b5050509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036120ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120c1906155d0565b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff16826040516120ef90615611565b5f6040518083038185875af1925050503d805f8114612129576040519150601f19603f3d011682016040523d82523d5f602084013e61212e565b606091505b5050905080612172576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121699061566f565b60405180910390fd5b505050565b6001600581905550565b60605f808585905067ffffffffffffffff8111156121a2576121a1613a48565b5b6040519080825280602002602001820160405280156121db57816020015b6121c86130e4565b8152602001906001900390816121c05790505b5090505f5b8686905081101561229057848787838181106121ff576121fe614184565b5b905060200281019061221191906141bd565b6080016020810190612223919061356e565b67ffffffffffffffff16036122855786868281811061224557612244614184565b5b905060200281019061225791906141bd565b6122609061568d565b82848151811061227357612272614184565b5b60200260200101819052508260010192505b8060010190506121e0565b505f8267ffffffffffffffff8111156122ac576122ab613a48565b5b6040519080825280602002602001820160405280156122e557816020015b6122d26130e4565b8152602001906001900390816122ca5790505b5090505f5b838110156123365782818151811061230557612304614184565b5b60200260200101518282815181106123205761231f614184565b5b60200260200101819052508060010190506122ea565b508093505050509392505050565b8151816123519190614a79565b83511015612394576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161238b906156e9565b60405180910390fd5b5f5b82518110156123f0578281815181106123b2576123b1614184565b5b60200260200101518482846123c79190614a79565b815181106123d8576123d7614184565b5b60200260200101819052508080600101915050612396565b50505050565b5f60405180604001604052808481526020018381525090508060015f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f820151815f01556020820151816001015590505050505050565b8060025f8282829054906101000a900467ffffffffffffffff166124739190614498565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b3073ffffffffffffffffffffffffffffffffffffffff16632b4280f58484846040518463ffffffff1660e01b81526004016124d99392919061590a565b5f604051808303815f87803b1580156124f0575f80fd5b505af1158015612502573d5f803e3d5ffd5b50505050505050565b612513612e72565b73ffffffffffffffffffffffffffffffffffffffff16612531611350565b73ffffffffffffffffffffffffffffffffffffffff161461259057612554612e72565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161258791906134ba565b60405180910390fd5b565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60605f60f467ffffffffffffffff81111561267357612672613a48565b5b6040519080825280601f01601f1916602001820160405280156126a55781602001600182028036833780820191505090505b5090505f6126b284612e79565b90505f8460400160208101906126c891906135c3565b90505f6126d486612ea9565b90505f6126ef878060a001906126ea9190614fda565b612ef3565b90505f6126fb88612f09565b90505f88610100016020810190612712919061356e565b67ffffffffffffffff1690505f6127288a612f53565b90505f8a61016001602081019061273f91906135c3565b90508760208a01528660408a01528560608a01528460808a01528360a08a01528260c08a01528160e08a01528060f48a0152889950505050505050505050919050565b61278a61250b565b565b5f805f845160208601878987f19050949350505050565b60603d828111156127b2578290505b604051602082018101604052818152815f602083013e8092505050919050565b5f805a90505f80865f015190505f6127e982612f9f565b9050815f015192505a8403860195505f8260a0015183608001518460600151010167ffffffffffffffff1690505f8960800151880390508082111561284e575f81830390505f6064600a830281612843576128426145c4565b5b049050808a01995050505b505080860294505f88604001519050858110156128da5760028081111561287857612877615946565b5b8a600281111561288b5761288a615946565b5b036128ad5780955061289c89612dbe565b6128a8895f888a612e0f565b6128d5565b7fdeadaa51000000000000000000000000000000000000000000000000000000005f5260205ffd5b612912565b5f8060028111156128ee576128ed615946565b5b8b600281111561290157612900615946565b5b1490506129108a82898b612e0f565b505b5050505050949350505050565b81604001602081019061293291906135c3565b815f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505081608001602081019061297b919061356e565b67ffffffffffffffff168160200181815250508160200135816040018181525050816101000160208101906129b0919061356e565b816060019067ffffffffffffffff16908167ffffffffffffffff1681525050816101200160208101906129e3919061356e565b67ffffffffffffffff168160c001906fffffffffffffffffffffffffffffffff1690816fffffffffffffffffffffffffffffffff168152505081610140016020810190612a30919061356e565b67ffffffffffffffff168160e001906fffffffffffffffffffffffffffffffff1690816fffffffffffffffffffffffffffffffff16815250508160c0016020810190612a7c919061356e565b816080019067ffffffffffffffff16908167ffffffffffffffff16815250508160e0016020810190612aae919061356e565b8160a0019067ffffffffffffffff16908167ffffffffffffffff16815250505050565b5f8160e001518260a0015167ffffffffffffffff16028260c00151836060015184608001510167ffffffffffffffff1602016fffffffffffffffffffffffffffffffff169050919050565b5f6127109050919050565b5f819050919050565b5f816040016020810190612b4491906135c3565b90505f82602001359050805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015612bc4576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612c0f9190615553565b925050819055508173ffffffffffffffffffffffffffffffffffffffff167f9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d82604051612c5c91906135fd565b60405180910390a2505050565b5f816040016020810190612c7d91906135c3565b90505f8260200135905080471015612cc1576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff1682604051612ce690615611565b5f6040518083038185875af1925050503d805f8114612d20576040519150601f19603f3d011682016040523d82523d5f602084013e612d25565b606091505b5050905080612d60576040517f3204506f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff167fe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f83604051612da691906135fd565b60405180910390a250505050565b6060819050919050565b805f01515f015173ffffffffffffffffffffffffffffffffffffffff1681602001517f1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff2860405160405180910390a350565b835f01515f015173ffffffffffffffffffffffffffffffffffffffff1684602001517ffb68fae769a73ec810729fb1cf34aa9faf18a18c0d69fb5388726c793abbf5aa858585604051612e6493929190615973565b60405180910390a350505050565b5f33905090565b5f80825f016020810190612e8d9190614faf565b90505f836020013590508160f81b925080831792505050919050565b5f612eec826060016020810190612ec0919061356e565b67ffffffffffffffff16836080016020810190612edd919061356e565b67ffffffffffffffff16612fd4565b9050919050565b5f60405182808583378082209250505092915050565b5f612f4c8260c0016020810190612f20919061356e565b67ffffffffffffffff168360e0016020810190612f3d919061356e565b67ffffffffffffffff16612fd4565b9050919050565b5f612f9882610120016020810190612f6b919061356e565b67ffffffffffffffff1683610140016020810190612f89919061356e565b67ffffffffffffffff16612fd4565b9050919050565b5f46826020015114612fb5578160e00151612fbb565b8160c001515b6fffffffffffffffffffffffffffffffff169050919050565b5f6fffffffffffffffffffffffffffffffff801683111561302a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613021906159f2565b60405180910390fd5b6fffffffffffffffffffffffffffffffff801682111561307f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161307690615a5a565b60405180910390fd5b81608084901b175f1b905092915050565b6040518060a001604052806130a36131af565b81526020015f80191681526020015f81526020015f81526020015f81525090565b60405180604001604052806130d7613244565b8152602001606081525090565b6040518061018001604052805f60ff1681526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b6040518061010001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f6fffffffffffffffffffffffffffffffff1681526020015f6fffffffffffffffffffffffffffffffff1681525090565b60405180606001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681525090565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6132c581613291565b81146132cf575f80fd5b50565b5f813590506132e0816132bc565b92915050565b5f602082840312156132fb576132fa613289565b5b5f613308848285016132d2565b91505092915050565b5f8115159050919050565b61332581613311565b82525050565b5f60208201905061333e5f83018461331c565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261336557613364613344565b5b8235905067ffffffffffffffff81111561338257613381613348565b5b60208301915083602082028301111561339e5761339d61334c565b5b9250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6133ce826133a5565b9050919050565b6133de816133c4565b81146133e8575f80fd5b50565b5f813590506133f9816133d5565b92915050565b61340881613311565b8114613412575f80fd5b50565b5f81359050613423816133ff565b92915050565b5f805f806060858703121561344157613440613289565b5b5f85013567ffffffffffffffff81111561345e5761345d61328d565b5b61346a87828801613350565b9450945050602061347d878288016133eb565b925050604061348e87828801613415565b91505092959194509250565b5f6134a4826133a5565b9050919050565b6134b48161349a565b82525050565b5f6020820190506134cd5f8301846134ab565b92915050565b5f819050919050565b6134e5816134d3565b81146134ef575f80fd5b50565b5f81359050613500816134dc565b92915050565b5f6020828403121561351b5761351a613289565b5b5f613528848285016134f2565b91505092915050565b5f67ffffffffffffffff82169050919050565b61354d81613531565b8114613557575f80fd5b50565b5f8135905061356881613544565b92915050565b5f6020828403121561358357613582613289565b5b5f6135908482850161355a565b91505092915050565b6135a28161349a565b81146135ac575f80fd5b50565b5f813590506135bd81613599565b92915050565b5f602082840312156135d8576135d7613289565b5b5f6135e5848285016135af565b91505092915050565b6135f7816134d3565b82525050565b5f6020820190506136105f8301846135ee565b92915050565b5f806040838503121561362c5761362b613289565b5b5f613639858286016134f2565b925050602061364a858286016134f2565b9150509250929050565b5f819050919050565b61366681613654565b82525050565b5f60408201905061367f5f83018561365d565b61368c602083018461365d565b9392505050565b5f8083601f8401126136a8576136a7613344565b5b8235905067ffffffffffffffff8111156136c5576136c4613348565b5b6020830191508360018202830111156136e1576136e061334c565b5b9250929050565b5f8083601f8401126136fd576136fc613344565b5b8235905067ffffffffffffffff81111561371a57613719613348565b5b6020830191508360208202830111156137365761373561334c565b5b9250929050565b5f80fd5b5f606082840312156137565761375561373d565b5b81905092915050565b5f805f805f6060868803121561377857613777613289565b5b5f86013567ffffffffffffffff8111156137955761379461328d565b5b6137a188828901613693565b9550955050602086013567ffffffffffffffff8111156137c4576137c361328d565b5b6137d0888289016136e8565b9350935050604086013567ffffffffffffffff8111156137f3576137f261328d565b5b6137ff88828901613741565b9150509295509295909350565b5f61018082840312156138225761382161373d565b5b81905092915050565b5f602082840312156138405761383f613289565b5b5f82013567ffffffffffffffff81111561385d5761385c61328d565b5b6138698482850161380c565b91505092915050565b5f6020820190506138855f83018461365d565b92915050565b61389481613531565b82525050565b5f6020820190506138ad5f83018461388b565b92915050565b5f805f604084860312156138ca576138c9613289565b5b5f6138d7868287016135af565b935050602084013567ffffffffffffffff8111156138f8576138f761328d565b5b61390486828701613693565b92509250509250925092565b5f8083601f84011261392557613924613344565b5b8235905067ffffffffffffffff81111561394257613941613348565b5b60208301915083602082028301111561395e5761395d61334c565b5b9250929050565b5f8083601f84011261397a57613979613344565b5b8235905067ffffffffffffffff81111561399757613996613348565b5b6020830191508360208202830111156139b3576139b261334c565b5b9250929050565b5f805f80604085870312156139d2576139d1613289565b5b5f85013567ffffffffffffffff8111156139ef576139ee61328d565b5b6139fb87828801613910565b9450945050602085013567ffffffffffffffff811115613a1e57613a1d61328d565b5b613a2a87828801613965565b925092505092959194509250565b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b613a7e82613a38565b810181811067ffffffffffffffff82111715613a9d57613a9c613a48565b5b80604052505050565b5f613aaf613280565b9050613abb8282613a75565b919050565b5f67ffffffffffffffff821115613ada57613ad9613a48565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f60ff82169050919050565b613b0881613af3565b8114613b12575f80fd5b50565b5f81359050613b2381613aff565b92915050565b5f80fd5b5f67ffffffffffffffff821115613b4757613b46613a48565b5b613b5082613a38565b9050602081019050919050565b828183375f83830152505050565b5f613b7d613b7884613b2d565b613aa6565b905082815260208101848484011115613b9957613b98613b29565b5b613ba4848285613b5d565b509392505050565b5f82601f830112613bc057613bbf613344565b5b8135613bd0848260208601613b6b565b91505092915050565b5f6101808284031215613bef57613bee613aeb565b5b613bfa610180613aa6565b90505f613c0984828501613b15565b5f830152506020613c1c848285016134f2565b6020830152506040613c30848285016135af565b6040830152506060613c448482850161355a565b6060830152506080613c588482850161355a565b60808301525060a082013567ffffffffffffffff811115613c7c57613c7b613aef565b5b613c8884828501613bac565b60a08301525060c0613c9c8482850161355a565b60c08301525060e0613cb08482850161355a565b60e083015250610100613cc58482850161355a565b61010083015250610120613cdb8482850161355a565b61012083015250610140613cf18482850161355a565b61014083015250610160613d07848285016135af565b6101608301525092915050565b5f613d26613d2184613ac0565b613aa6565b90508083825260208201905060208402830185811115613d4957613d4861334c565b5b835b81811015613d9057803567ffffffffffffffff811115613d6e57613d6d613344565b5b808601613d7b8982613bd9565b85526020850194505050602081019050613d4b565b5050509392505050565b5f82601f830112613dae57613dad613344565b5b8135613dbe848260208601613d14565b91505092915050565b5f60208284031215613ddc57613ddb613289565b5b5f82013567ffffffffffffffff811115613df957613df861328d565b5b613e0584828501613d9a565b91505092915050565b5f6fffffffffffffffffffffffffffffffff82169050919050565b613e3281613e0e565b8114613e3c575f80fd5b50565b5f81359050613e4d81613e29565b92915050565b5f6101008284031215613e6957613e68613aeb565b5b613e74610100613aa6565b90505f613e83848285016135af565b5f830152506020613e96848285016134f2565b6020830152506040613eaa848285016134f2565b6040830152506060613ebe8482850161355a565b6060830152506080613ed28482850161355a565b60808301525060a0613ee68482850161355a565b60a08301525060c0613efa84828501613e3f565b60c08301525060e0613f0e84828501613e3f565b60e08301525092915050565b613f2381613654565b8114613f2d575f80fd5b50565b5f81359050613f3e81613f1a565b92915050565b5f6101808284031215613f5a57613f59613aeb565b5b613f6460a0613aa6565b90505f613f7384828501613e53565b5f83015250610100613f8784828501613f30565b602083015250610120613f9c848285016134f2565b604083015250610140613fb1848285016134f2565b606083015250610160613fc6848285016134f2565b60808301525092915050565b5f805f806101c08587031215613feb57613fea613289565b5b5f85013567ffffffffffffffff8111156140085761400761328d565b5b61401487828801613bac565b945050602061402587828801613f44565b9350506101a085013567ffffffffffffffff8111156140475761404661328d565b5b61405387828801613693565b925092505092959194509250565b5f805f805f8060c0878903121561407b5761407a613289565b5b5f61408889828a0161355a565b965050602061409989828a0161355a565b95505060406140aa89828a01613f30565b94505060606140bb89828a01613f30565b93505060806140cc89828a01613f30565b92505060a06140dd89828a01613f30565b9150509295509295509295565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015614121578082015181840152602081019050614106565b5f8484015250505050565b5f614136826140ea565b61414081856140f4565b9350614150818560208601614104565b61415981613a38565b840191505092915050565b5f6020820190508181035f83015261417c818461412c565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f80fd5b5f80fd5b5f80fd5b5f82356001610180038336030381126141d9576141d86141b1565b5b80830191505092915050565b5f6040820190506141f85f8301856135ee565b61420560208301846135ee565b9392505050565b5f8160601b9050919050565b5f6142228261420c565b9050919050565b5f61423382614218565b9050919050565b61424b6142468261349a565b614229565b82525050565b5f819050919050565b61426b614266826134d3565b614251565b82525050565b5f61427c828761423a565b60148201915061428c828661425a565b60208201915061429c828561425a565b6020820191506142ac828461425a565b60208201915081905095945050505050565b5f67ffffffffffffffff8211156142d8576142d7613a48565b5b602082029050919050565b5f6142f56142f0846142be565b613aa6565b9050806020840283018581111561430f5761430e61334c565b5b835b81811015614338578061432488826134f2565b845260208401935050602081019050614311565b5050509392505050565b5f82601f83011261435657614355613344565b5b60026143638482856142e3565b91505092915050565b5f67ffffffffffffffff82111561438657614385613a48565b5b602082029050919050565b5f6143a361439e8461436c565b613aa6565b905080604084028301858111156143bd576143bc61334c565b5b835b818110156143e657806143d28882614342565b8452602084019350506040810190506143bf565b5050509392505050565b5f82601f83011261440457614403613344565b5b6002614411848285614391565b91505092915050565b5f805f610100848603121561443257614431613289565b5b5f61443f86828701614342565b9350506040614450868287016143f0565b92505060c061446186828701614342565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6144a282613531565b91506144ad83613531565b9250828201905067ffffffffffffffff8111156144cd576144cc61446b565b5b92915050565b5f6144dd82613531565b91506144e883613531565b9250828203905067ffffffffffffffff8111156145085761450761446b565b5b92915050565b5f82356001604003833603038112614529576145286141b1565b5b80830191505092915050565b5f81905092915050565b5f614549826140ea565b6145538185614535565b9350614563818560208601614104565b80840191505092915050565b5f61457a828461453f565b915081905092915050565b5f8151905061459381613f1a565b92915050565b5f602082840312156145ae576145ad613289565b5b5f6145bb84828501614585565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6145fb826134d3565b9150614606836134d3565b925082614616576146156145c4565b5b828206905092915050565b5f60029050919050565b5f81905092915050565b5f819050919050565b614647816134d3565b82525050565b5f614658838361463e565b60208301905092915050565b5f602082019050919050565b61467981614621565b614683818461462b565b925061468e82614635565b805f5b838110156146be5781516146a5878261464d565b96506146b083614664565b925050600181019050614691565b505050505050565b5f60029050919050565b5f81905092915050565b5f819050919050565b5f81905092915050565b6146f681614621565b61470081846146e3565b925061470b82614635565b805f5b8381101561473b578151614722878261464d565b965061472d83614664565b92505060018101905061470e565b505050505050565b5f61474e83836146ed565b60408301905092915050565b5f602082019050919050565b61476f816146c6565b61477981846146d0565b9250614784826146da565b805f5b838110156147b457815161479b8782614743565b96506147a68361475a565b925050600181019050614787565b505050505050565b5f60019050919050565b5f81905092915050565b5f819050919050565b5f602082019050919050565b6147ee816147bc565b6147f881846147c6565b9250614803826147d0565b805f5b8381101561483357815161481a878261464d565b9650614825836147d9565b925050600181019050614806565b505050505050565b5f6101208201905061484f5f830187614670565b61485c6040830186614766565b61486960c0830185614670565b6148776101008301846147e5565b95945050505050565b5f8151905061488e816133ff565b92915050565b5f602082840312156148a9576148a8613289565b5b5f6148b684828501614880565b91505092915050565b5f80833560016020038436030381126148db576148da6141b1565b5b80840192508235915067ffffffffffffffff8211156148fd576148fc6141b5565b5b602083019250606082023603831315614919576149186141b9565b5b509250929050565b5f6060828403121561493657614935613aeb565b5b6149406060613aa6565b90505f61494f8482850161355a565b5f8301525060206149628482850161355a565b60208301525060406149768482850161355a565b60408301525092915050565b5f6060828403121561499757614996613289565b5b5f6149a484828501614921565b91505092915050565b5f80833560016020038436030381126149c9576149c86141b1565b5b80840192508235915067ffffffffffffffff8211156149eb576149ea6141b5565b5b602083019250602082023603831315614a0757614a066141b9565b5b509250929050565b5f614a1a8385614535565b9350614a27838584613b5d565b82840190509392505050565b5f614a3f828486614a0f565b91508190509392505050565b5f604082019050614a5e5f83018561331c565b8181036020830152614a70818461412c565b90509392505050565b5f614a83826134d3565b9150614a8e836134d3565b9250828201905080821115614aa657614aa561446b565b5b92915050565b5f8135614ab881613599565b80915050919050565b5f815f1b9050919050565b5f73ffffffffffffffffffffffffffffffffffffffff614aeb84614ac1565b9350801983169250808416831791505092915050565b5f819050919050565b5f614b24614b1f614b1a846133a5565b614b01565b6133a5565b9050919050565b5f614b3582614b0a565b9050919050565b5f614b4682614b2b565b9050919050565b5f819050919050565b614b5f82614b3c565b614b72614b6b82614b4d565b8354614acc565b8255505050565b5f81015f830180614b8981614aac565b9050614b958184614b56565b5050505050565b614ba68282614b79565b5050565b5f82825260208201905092915050565b7f4e455153525200000000000000000000000000000000000000000000000000005f82015250565b5f614bee600683614baa565b9150614bf982614bba565b602082019050919050565b5f6020820190508181035f830152614c1b81614be2565b9050919050565b7f4141393220696e7465726e616c2063616c6c206f6e6c790000000000000000005f82015250565b5f614c56601783614baa565b9150614c6182614c22565b602082019050919050565b5f6020820190508181035f830152614c8381614c4a565b9050919050565b5f614c9482613531565b9150614c9f83613531565b9250828202614cad81613531565b9150808214614cbf57614cbe61446b565b5b5092915050565b5f819050919050565b614ce0614cdb82613654565b614cc6565b82525050565b5f8160c01b9050919050565b5f614cfc82614ce6565b9050919050565b614d14614d0f82613531565b614cf2565b82525050565b5f614d25828d61423a565b601482019150614d35828c614ccf565b602082019150614d45828b614ccf565b602082019150614d55828a614d03565b600882019150614d658289614d03565b600882019150614d758288614d03565b600882019150614d858287614ccf565b602082019150614d958286614ccf565b602082019150614da58285614ccf565b602082019150614db58284614d03565b6008820191508190509b9a5050505050505050505050565b7f41413934206761732076616c756573206f766572666c6f7700000000000000005f82015250565b5f614e01601883614baa565b9150614e0c82614dcd565b602082019050919050565b5f6020820190508181035f830152614e2e81614df5565b9050919050565b7f41413934206f7065726174696f6e2076616c7565206f766572666c6f770000005f82015250565b5f614e69601d83614baa565b9150614e7482614e35565b602082019050919050565b5f6020820190508181035f830152614e9681614e5d565b9050919050565b7f4141206f776e657220766572696669636174696f6e2066616c696564000000005f82015250565b5f614ed1601c83614baa565b9150614edc82614e9d565b602082019050919050565b5f604082019050614efa5f8301846135ee565b8181036020830152614f0b81614ec5565b905092915050565b7f41413236206f76657220766572696669636174696f6e4f776e65724761734c695f8201527f6d69740000000000000000000000000000000000000000000000000000000000602082015250565b5f614f6d602383614baa565b9150614f7882614f13565b604082019050919050565b5f604082019050614f965f8301846135ee565b8181036020830152614fa781614f61565b905092915050565b5f60208284031215614fc457614fc3613289565b5b5f614fd184828501613b15565b91505092915050565b5f8083356001602003843603038112614ff657614ff56141b1565b5b80840192508235915067ffffffffffffffff821115615018576150176141b5565b5b602083019250600182023603831315615034576150336141b9565b5b509250929050565b5f61504a6020840184613b15565b905092915050565b61505b81613af3565b82525050565b5f61506f60208401846134f2565b905092915050565b5f61508560208401846135af565b905092915050565b6150968161349a565b82525050565b5f6150aa602084018461355a565b905092915050565b6150bb81613531565b82525050565b5f80fd5b5f80fd5b5f80fd5b5f80833560016020038436030381126150e9576150e86150c9565b5b83810192508235915060208301925067ffffffffffffffff821115615111576151106150c1565b5b600182023603831315615127576151266150c5565b5b509250929050565b5f82825260208201905092915050565b5f61514a838561512f565b9350615157838584613b5d565b61516083613a38565b840190509392505050565b5f610180830161517d5f84018461503c565b6151895f860182615052565b506151976020840184615061565b6151a4602086018261463e565b506151b26040840184615077565b6151bf604086018261508d565b506151cd606084018461509c565b6151da60608601826150b2565b506151e8608084018461509c565b6151f560808601826150b2565b5061520360a08401846150cd565b85830360a087015261521683828461513f565b9250505061522760c084018461509c565b61523460c08601826150b2565b5061524260e084018461509c565b61524f60e08601826150b2565b5061525e61010084018461509c565b61526c6101008601826150b2565b5061527b61012084018461509c565b6152896101208601826150b2565b5061529861014084018461509c565b6152a66101408601826150b2565b506152b5610160840184615077565b6152c361016086018261508d565b508091505092915050565b5f6040820190508181035f8301526152e6818561516b565b90506152f5602083018461365d565b9392505050565b61530581613e0e565b82525050565b61010082015f8201516153205f85018261508d565b506020820151615333602085018261463e565b506040820151615346604085018261463e565b50606082015161535960608501826150b2565b50608082015161536c60808501826150b2565b5060a082015161537f60a08501826150b2565b5060c082015161539260c08501826152fc565b5060e08201516153a560e08501826152fc565b50505050565b6153b481613654565b82525050565b61018082015f8201516153cf5f85018261530b565b5060208201516153e36101008501826153ab565b5060408201516153f761012085018261463e565b50606082015161540b61014085018261463e565b50608082015161541f61016085018261463e565b50505050565b5f6101c0820190508181035f83015261543e818661412c565b905061544d60208301856153ba565b8181036101a0830152615460818461412c565b9050949350505050565b5f61547583856140f4565b9350615482838584613b5d565b61548b83613a38565b840190509392505050565b5f6101c0820190508181035f8301526154b081868861546a565b90506154bf60208301856153ba565b8181036101a08301526154d2818461412c565b905095945050505050565b7f41413935206f7574206f662067617300000000000000000000000000000000005f82015250565b5f615511600f83614baa565b915061551c826154dd565b602082019050919050565b5f60408201905061553a5f8301846135ee565b818103602083015261554b81615505565b905092915050565b5f61555d826134d3565b9150615568836134d3565b92508282039050818111156155805761557f61446b565b5b92915050565b7f4141393020696e76616c69642062656e656669636961727900000000000000005f82015250565b5f6155ba601883614baa565b91506155c582615586565b602082019050919050565b5f6020820190508181035f8301526155e7816155ae565b9050919050565b50565b5f6155fc5f83614535565b9150615607826155ee565b5f82019050919050565b5f61561b826155f1565b9150819050919050565b7f41413931206661696c65642073656e6420746f2062656e6566696369617279005f82015250565b5f615659601f83614baa565b915061566482615625565b602082019050919050565b5f6020820190508181035f8301526156868161564d565b9050919050565b5f6156983683613bd9565b9050919050565b7f54617267657420617272617920746f6f20736d616c6c000000000000000000005f82015250565b5f6156d3601683614baa565b91506156de8261569f565b602082019050919050565b5f6020820190508181035f830152615700816156c7565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f61573a826140ea565b615744818561512f565b9350615754818560208601614104565b61575d81613a38565b840191505092915050565b5f61018083015f83015161577e5f860182615052565b506020830151615791602086018261463e565b5060408301516157a4604086018261508d565b5060608301516157b760608601826150b2565b5060808301516157ca60808601826150b2565b5060a083015184820360a08601526157e28282615730565b91505060c08301516157f760c08601826150b2565b5060e083015161580a60e08601826150b2565b5061010083015161581f6101008601826150b2565b506101208301516158346101208601826150b2565b506101408301516158496101408601826150b2565b5061016083015161585e61016086018261508d565b508091505092915050565b5f6158748383615768565b905092915050565b5f602082019050919050565b5f61589282615707565b61589c8185615711565b9350836020820285016158ae85615721565b805f5b858110156158e957848403895281516158ca8582615869565b94506158d58361587c565b925060208a019950506001810190506158b1565b50829750879550505050505092915050565b615904816133c4565b82525050565b5f6060820190508181035f8301526159228186615888565b905061593160208301856158fb565b61593e604083018461331c565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f6060820190506159865f83018661331c565b61599360208301856135ee565b6159a060408301846135ee565b949350505050565b7f6869676831323820657863656564732075696e743132382072616e67650000005f82015250565b5f6159dc601d83614baa565b91506159e7826159a8565b602082019050919050565b5f6020820190508181035f830152615a09816159d0565b9050919050565b7f6c6f7731323820657863656564732075696e743132382072616e6765000000005f82015250565b5f615a44601c83614baa565b9150615a4f82615a10565b602082019050919050565b5f6020820190508181035f830152615a7181615a38565b905091905056fea264697066735822122050aa55a8f9946da6fe1e57eb73e2b23bbe9af948e9e17af109f1515b39a360e164736f6c63430008180033",
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

// FORKID is a free data retrieval call binding the contract method 0xe8516092.
//
// Solidity: function FORK_ID() view returns(uint64)
func (_EntryPoint *EntryPointCaller) FORKID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "FORK_ID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FORKID is a free data retrieval call binding the contract method 0xe8516092.
//
// Solidity: function FORK_ID() view returns(uint64)
func (_EntryPoint *EntryPointSession) FORKID() (uint64, error) {
	return _EntryPoint.Contract.FORKID(&_EntryPoint.CallOpts)
}

// FORKID is a free data retrieval call binding the contract method 0xe8516092.
//
// Solidity: function FORK_ID() view returns(uint64)
func (_EntryPoint *EntryPointCallerSession) FORKID() (uint64, error) {
	return _EntryPoint.Contract.FORKID(&_EntryPoint.CallOpts)
}

// BatchNumToState is a free data retrieval call binding the contract method 0x6b086237.
//
// Solidity: function batchNumToState(uint64 ) view returns(bytes32 stateRoot, bytes32 accInputRoot)
func (_EntryPoint *EntryPointCaller) BatchNumToState(opts *bind.CallOpts, arg0 uint64) (struct {
	StateRoot    [32]byte
	AccInputRoot [32]byte
}, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "batchNumToState", arg0)

	outstruct := new(struct {
		StateRoot    [32]byte
		AccInputRoot [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StateRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.AccInputRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// BatchNumToState is a free data retrieval call binding the contract method 0x6b086237.
//
// Solidity: function batchNumToState(uint64 ) view returns(bytes32 stateRoot, bytes32 accInputRoot)
func (_EntryPoint *EntryPointSession) BatchNumToState(arg0 uint64) (struct {
	StateRoot    [32]byte
	AccInputRoot [32]byte
}, error) {
	return _EntryPoint.Contract.BatchNumToState(&_EntryPoint.CallOpts, arg0)
}

// BatchNumToState is a free data retrieval call binding the contract method 0x6b086237.
//
// Solidity: function batchNumToState(uint64 ) view returns(bytes32 stateRoot, bytes32 accInputRoot)
func (_EntryPoint *EntryPointCallerSession) BatchNumToState(arg0 uint64) (struct {
	StateRoot    [32]byte
	AccInputRoot [32]byte
}, error) {
	return _EntryPoint.Contract.BatchNumToState(&_EntryPoint.CallOpts, arg0)
}

// ChainConfigs is a free data retrieval call binding the contract method 0x4e425f76.
//
// Solidity: function chainConfigs(uint64 ) view returns(address entryPoint)
func (_EntryPoint *EntryPointCaller) ChainConfigs(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "chainConfigs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChainConfigs is a free data retrieval call binding the contract method 0x4e425f76.
//
// Solidity: function chainConfigs(uint64 ) view returns(address entryPoint)
func (_EntryPoint *EntryPointSession) ChainConfigs(arg0 uint64) (common.Address, error) {
	return _EntryPoint.Contract.ChainConfigs(&_EntryPoint.CallOpts, arg0)
}

// ChainConfigs is a free data retrieval call binding the contract method 0x4e425f76.
//
// Solidity: function chainConfigs(uint64 ) view returns(address entryPoint)
func (_EntryPoint *EntryPointCallerSession) ChainConfigs(arg0 uint64) (common.Address, error) {
	return _EntryPoint.Contract.ChainConfigs(&_EntryPoint.CallOpts, arg0)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0xe2aeaa35.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 oldAccInputHash, bytes32 newAccInputHash, bytes32 oldStateRoot, bytes32 newStateRoot) pure returns(bytes)
func (_EntryPoint *EntryPointCaller) GetInputSnarkBytes(opts *bind.CallOpts, initNumBatch uint64, finalNewBatch uint64, oldAccInputHash [32]byte, newAccInputHash [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getInputSnarkBytes", initNumBatch, finalNewBatch, oldAccInputHash, newAccInputHash, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0xe2aeaa35.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 oldAccInputHash, bytes32 newAccInputHash, bytes32 oldStateRoot, bytes32 newStateRoot) pure returns(bytes)
func (_EntryPoint *EntryPointSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, oldAccInputHash [32]byte, newAccInputHash [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _EntryPoint.Contract.GetInputSnarkBytes(&_EntryPoint.CallOpts, initNumBatch, finalNewBatch, oldAccInputHash, newAccInputHash, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0xe2aeaa35.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 oldAccInputHash, bytes32 newAccInputHash, bytes32 oldStateRoot, bytes32 newStateRoot) pure returns(bytes)
func (_EntryPoint *EntryPointCallerSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, oldAccInputHash [32]byte, newAccInputHash [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _EntryPoint.Contract.GetInputSnarkBytes(&_EntryPoint.CallOpts, initNumBatch, finalNewBatch, oldAccInputHash, newAccInputHash, oldStateRoot, newStateRoot)
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

// GetUserOpHash is a free data retrieval call binding the contract method 0x7b031fbc.
//
// Solidity: function getUserOpHash((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address) userOp) pure returns(bytes32)
func (_EntryPoint *EntryPointCaller) GetUserOpHash(opts *bind.CallOpts, userOp BaseStructPackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0x7b031fbc.
//
// Solidity: function getUserOpHash((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address) userOp) pure returns(bytes32)
func (_EntryPoint *EntryPointSession) GetUserOpHash(userOp BaseStructPackedUserOperation) ([32]byte, error) {
	return _EntryPoint.Contract.GetUserOpHash(&_EntryPoint.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x7b031fbc.
//
// Solidity: function getUserOpHash((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address) userOp) pure returns(bytes32)
func (_EntryPoint *EntryPointCallerSession) GetUserOpHash(userOp BaseStructPackedUserOperation) ([32]byte, error) {
	return _EntryPoint.Contract.GetUserOpHash(&_EntryPoint.CallOpts, userOp)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_EntryPoint *EntryPointCaller) LastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "lastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_EntryPoint *EntryPointSession) LastVerifiedBatch() (uint64, error) {
	return _EntryPoint.Contract.LastVerifiedBatch(&_EntryPoint.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_EntryPoint *EntryPointCallerSession) LastVerifiedBatch() (uint64, error) {
	return _EntryPoint.Contract.LastVerifiedBatch(&_EntryPoint.CallOpts)
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

// HandleOps is a paid mutator transaction binding the contract method 0x2b4280f5.
//
// Solidity: function handleOps((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] ops, address beneficiary, bool isSync) returns()
func (_EntryPoint *EntryPointTransactor) HandleOps(opts *bind.TransactOpts, ops []BaseStructPackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "handleOps", ops, beneficiary, isSync)
}

// HandleOps is a paid mutator transaction binding the contract method 0x2b4280f5.
//
// Solidity: function handleOps((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] ops, address beneficiary, bool isSync) returns()
func (_EntryPoint *EntryPointSession) HandleOps(ops []BaseStructPackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, beneficiary, isSync)
}

// HandleOps is a paid mutator transaction binding the contract method 0x2b4280f5.
//
// Solidity: function handleOps((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] ops, address beneficiary, bool isSync) returns()
func (_EntryPoint *EntryPointTransactorSession) HandleOps(ops []BaseStructPackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, beneficiary, isSync)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0xc017d139.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint64,uint64,uint64,uint128,uint128),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactor) InnerHandleOp(opts *bind.TransactOpts, callData []byte, opInfo BaseStructUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "innerHandleOp", callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0xc017d139.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint64,uint64,uint64,uint128,uint128),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointSession) InnerHandleOp(callData []byte, opInfo BaseStructUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0xc017d139.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint64,uint64,uint64,uint128,uint128),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactorSession) InnerHandleOp(callData []byte, opInfo BaseStructUserOpInfo, context []byte) (*types.Transaction, error) {
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

// SyncBatches is a paid mutator transaction binding the contract method 0xba64b242.
//
// Solidity: function syncBatches((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] userOps) returns()
func (_EntryPoint *EntryPointTransactor) SyncBatches(opts *bind.TransactOpts, userOps []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "syncBatches", userOps)
}

// SyncBatches is a paid mutator transaction binding the contract method 0xba64b242.
//
// Solidity: function syncBatches((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] userOps) returns()
func (_EntryPoint *EntryPointSession) SyncBatches(userOps []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _EntryPoint.Contract.SyncBatches(&_EntryPoint.TransactOpts, userOps)
}

// SyncBatches is a paid mutator transaction binding the contract method 0xba64b242.
//
// Solidity: function syncBatches((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[] userOps) returns()
func (_EntryPoint *EntryPointTransactorSession) SyncBatches(userOps []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _EntryPoint.Contract.SyncBatches(&_EntryPoint.TransactOpts, userOps)
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

// UpdateChainConfigs is a paid mutator transaction binding the contract method 0xa4c44999.
//
// Solidity: function updateChainConfigs(uint64[] _chainIds, (address)[] _config) returns()
func (_EntryPoint *EntryPointTransactor) UpdateChainConfigs(opts *bind.TransactOpts, _chainIds []uint64, _config []ConfigManagerConfig) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "updateChainConfigs", _chainIds, _config)
}

// UpdateChainConfigs is a paid mutator transaction binding the contract method 0xa4c44999.
//
// Solidity: function updateChainConfigs(uint64[] _chainIds, (address)[] _config) returns()
func (_EntryPoint *EntryPointSession) UpdateChainConfigs(_chainIds []uint64, _config []ConfigManagerConfig) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateChainConfigs(&_EntryPoint.TransactOpts, _chainIds, _config)
}

// UpdateChainConfigs is a paid mutator transaction binding the contract method 0xa4c44999.
//
// Solidity: function updateChainConfigs(uint64[] _chainIds, (address)[] _config) returns()
func (_EntryPoint *EntryPointTransactorSession) UpdateChainConfigs(_chainIds []uint64, _config []ConfigManagerConfig) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateChainConfigs(&_EntryPoint.TransactOpts, _chainIds, _config)
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

// VerifyBatches is a paid mutator transaction binding the contract method 0x6e891959.
//
// Solidity: function verifyBatches(bytes proof, ((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[],bytes32)[] batches, ((uint64,uint64,uint64)[],bytes32,address) chainsExecuteInfos) payable returns()
func (_EntryPoint *EntryPointTransactor) VerifyBatches(opts *bind.TransactOpts, proof []byte, batches []BaseStructBatchData, chainsExecuteInfos BaseStructChainsExecuteInfo) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "verifyBatches", proof, batches, chainsExecuteInfos)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x6e891959.
//
// Solidity: function verifyBatches(bytes proof, ((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[],bytes32)[] batches, ((uint64,uint64,uint64)[],bytes32,address) chainsExecuteInfos) payable returns()
func (_EntryPoint *EntryPointSession) VerifyBatches(proof []byte, batches []BaseStructBatchData, chainsExecuteInfos BaseStructChainsExecuteInfo) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatches(&_EntryPoint.TransactOpts, proof, batches, chainsExecuteInfos)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x6e891959.
//
// Solidity: function verifyBatches(bytes proof, ((uint8,uint256,address,uint64,uint64,bytes,uint64,uint64,uint64,uint64,uint64,address)[],bytes32)[] batches, ((uint64,uint64,uint64)[],bytes32,address) chainsExecuteInfos) payable returns()
func (_EntryPoint *EntryPointTransactorSession) VerifyBatches(proof []byte, batches []BaseStructBatchData, chainsExecuteInfos BaseStructChainsExecuteInfo) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatches(&_EntryPoint.TransactOpts, proof, batches, chainsExecuteInfos)
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
