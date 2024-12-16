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

// BaseStructMemoryUserOp is an auto generated low-level Go binding around an user-defined struct.
type BaseStructMemoryUserOp struct {
	Phase                  uint8
	Sender                 common.Address
	Owner                  common.Address
	InnerExec              bool
	ChainId                *big.Int
	ZkVerificationGasLimit uint64
	MainChainGasLimit      uint64
	DestChainGasLimit      uint64
	MainChainGasPrice      *big.Int
	DestChainGasPrice      *big.Int
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

// BaseStructUserOpInfo is an auto generated low-level Go binding around an user-defined struct.
type BaseStructUserOpInfo struct {
	MUserOp       BaseStructMemoryUserOp
	UserOpHash    [32]byte
	Prefund       *big.Int
	ContextOffset *big.Int
	PreOpGas      *big.Int
}

// IConfigManagerConfig is an auto generated low-level Go binding around an user-defined struct.
type IConfigManagerConfig struct {
	EntryPoint common.Address
	Router     common.Address
}

// BaseStructMetaData contains all meta data concerning the BaseStruct contract.
var BaseStructMetaData = &bind.MetaData{
	ABI: "[]",
}

// BaseStructABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseStructMetaData.ABI instead.
var BaseStructABI = BaseStructMetaData.ABI

// BaseStruct is an auto generated Go binding around an Ethereum contract.
type BaseStruct struct {
	BaseStructCaller     // Read-only binding to the contract
	BaseStructTransactor // Write-only binding to the contract
	BaseStructFilterer   // Log filterer for contract events
}

// BaseStructCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseStructCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseStructTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseStructTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseStructFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseStructFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseStructSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseStructSession struct {
	Contract     *BaseStruct       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseStructCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseStructCallerSession struct {
	Contract *BaseStructCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BaseStructTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseStructTransactorSession struct {
	Contract     *BaseStructTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BaseStructRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseStructRaw struct {
	Contract *BaseStruct // Generic contract binding to access the raw methods on
}

// BaseStructCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseStructCallerRaw struct {
	Contract *BaseStructCaller // Generic read-only contract binding to access the raw methods on
}

// BaseStructTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseStructTransactorRaw struct {
	Contract *BaseStructTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseStruct creates a new instance of BaseStruct, bound to a specific deployed contract.
func NewBaseStruct(address common.Address, backend bind.ContractBackend) (*BaseStruct, error) {
	contract, err := bindBaseStruct(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseStruct{BaseStructCaller: BaseStructCaller{contract: contract}, BaseStructTransactor: BaseStructTransactor{contract: contract}, BaseStructFilterer: BaseStructFilterer{contract: contract}}, nil
}

// NewBaseStructCaller creates a new read-only instance of BaseStruct, bound to a specific deployed contract.
func NewBaseStructCaller(address common.Address, caller bind.ContractCaller) (*BaseStructCaller, error) {
	contract, err := bindBaseStruct(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseStructCaller{contract: contract}, nil
}

// NewBaseStructTransactor creates a new write-only instance of BaseStruct, bound to a specific deployed contract.
func NewBaseStructTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseStructTransactor, error) {
	contract, err := bindBaseStruct(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseStructTransactor{contract: contract}, nil
}

// NewBaseStructFilterer creates a new log filterer instance of BaseStruct, bound to a specific deployed contract.
func NewBaseStructFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseStructFilterer, error) {
	contract, err := bindBaseStruct(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseStructFilterer{contract: contract}, nil
}

// bindBaseStruct binds a generic wrapper to an already deployed contract.
func bindBaseStruct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseStructMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseStruct *BaseStructRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseStruct.Contract.BaseStructCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseStruct *BaseStructRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseStruct.Contract.BaseStructTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseStruct *BaseStructRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseStruct.Contract.BaseStructTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseStruct *BaseStructCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseStruct.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseStruct *BaseStructTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseStruct.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseStruct *BaseStructTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseStruct.Contract.contract.Transact(opts, method, params...)
}

// ERC165MetaData contains all meta data concerning the ERC165 contract.
var ERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC165MetaData.ABI instead.
var ERC165ABI = ERC165MetaData.ABI

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// EntryPointMetaData contains all meta data concerning the EntryPoint contract.
var EntryPointMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"name\":\"DelegateAndRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"EstimateRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"inner\",\"type\":\"bytes\"}],\"name\":\"FailedOpWithRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"PostOpReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotEqual\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"did\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"innerExec\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FORK_ID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchNumToState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accInputRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"estimateCrossMessageParamsCrossGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"estimateSubmitDepositOperationByRemoteGas\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getChainConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"internalType\":\"structIConfigManager.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPreGasBalanceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSync\",\"type\":\"bool\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"innerExec\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"mainChainGasPrice\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"destChainGasPrice\",\"type\":\"uint128\"}],\"internalType\":\"structBaseStruct.MemoryUserOp\",\"name\":\"mUserOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contextOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"}],\"internalType\":\"structBaseStruct.UserOpInfo\",\"name\":\"opInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"innerHandleOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"sendUserOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitDepositOperation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitDepositOperationByRemote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submitWithdrawOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"}],\"name\":\"syncBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"internalType\":\"structIConfigManager.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"updateChainConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"}],\"internalType\":\"structBaseStruct.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainUserOperationsNumber\",\"type\":\"uint64\"}],\"internalType\":\"structBaseStruct.ChainExecuteExtra[]\",\"name\":\"chainExtra\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structBaseStruct.ChainsExecuteInfo\",\"name\":\"chainsExecuteInfos\",\"type\":\"tuple\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052346200002d5762000014620000d4565b6200001e62000033565b615bd0620001b88239615bd090f35b62000039565b60405190565b5f80fd5b5f1b90565b906200005560018060a01b03916200003d565b9181191691161790565b60018060a01b031690565b90565b62000086620000806200008c926200005f565b6200006a565b6200005f565b90565b6200009a906200006d565b90565b620000a8906200008f565b90565b90565b90620000c8620000c2620000d0926200009d565b620000ab565b825462000042565b9055565b620000de620000ed565b620000eb336004620000ae565b565b620000f7620000f9565b565b620001036200019f565b565b90565b90565b620001246200011e6200012a9262000105565b6200006a565b62000108565b90565b6200013960016200010b565b90565b906200014a5f19916200003d565b9181191691161790565b6200016d62000167620001739262000108565b6200006a565b62000108565b90565b90565b90620001936200018d6200019b9262000154565b62000176565b82546200013c565b9055565b620001b5620001ad6200012d565b600362000179565b56fe60806040526004361015610013575b611313565b61001d5f3561019c565b806301ffc9a7146101975780631b3348f0146101925780632885a84f1461018d5780632b7ac3f31461018857806335999cb31461018357806339875d1a1461017e57806342f1ec6914610179578063653e75ec146101745780636b0862371461016f578063750c67ba1461016a57806377627f0c1461016557806377c2719c146101605780637fcb36531461015b578063850aaf62146101565780638da5cb5b1461015157806397fc007c1461014c5780639abe790314610147578063b224160014610142578063bb8ce5261461013d578063d4115a0114610138578063df08ec3114610133578063e85160921461012e578063e9ab53d1146101295763f2fde38b0361000e576112e0565b6112b3565b6111a7565b61115f565b611126565b610ee6565b610e77565b610e4d565b610ded565b610db8565b610d9d565b610ce3565b610c6e565b6108ac565b610784565b61072c565b610651565b6105b9565b610522565b6104aa565b610405565b610344565b6102ae565b610228565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b63ffffffff60e01b1690565b6101c9816101b4565b036101d057565b5f80fd5b905035906101e1826101c0565b565b906020828203126101fc576101f9915f016101d4565b90565b6101ac565b151590565b61020f90610201565b9052565b9190610226905f60208501940190610206565b565b346102585761025461024361023e3660046101e3565b61131b565b61024b6101a2565b91829182610213565b0390f35b6101a8565b5f80fd5b9081604091031261026f5790565b61025d565b906020828203126102a4575f82013567ffffffffffffffff811161029f5761029c9201610261565b90565b6101b0565b6101ac565b5f0190565b6102c16102bc366004610274565b6119fa565b6102c96101a2565b806102d3816102a9565b0390f35b908160e09103126102e55790565b61025d565b9060208282031261031a575f82013567ffffffffffffffff81116103155761031292016102d7565b90565b6101b0565b6101ac565b90565b61032b9061031f565b9052565b9190610342905f60208501940190610322565b565b346103745761037061035f61035a3660046102ea565b611ab7565b6103676101a2565b9182918261032f565b0390f35b6101a8565b5f91031261038357565b6101ac565b1c90565b60018060a01b031690565b6103a79060086103ac9302610388565b61038c565b90565b906103ba9154610397565b90565b6103c960055f906103af565b90565b60018060a01b031690565b6103e0906103cc565b90565b6103ec906103d7565b9052565b9190610403905f602085019401906103e3565b565b3461043557610415366004610379565b6104316104206103bd565b6104286101a2565b918291826103f0565b0390f35b6101a8565b67ffffffffffffffff1690565b6104508161043a565b0361045757565b5f80fd5b9050359061046882610447565b565b908160409103126104785790565b61025d565b91906060838203126104a557806104996104a2925f860161045b565b9360200161046a565b90565b6101ac565b346104d9576104c36104bd36600461047d565b90611bf3565b6104cb6101a2565b806104d5816102a9565b0390f35b6101a8565b90565b6104ea816104de565b036104f157565b5f80fd5b90503590610502826104e1565b565b9060208282031261051d5761051a915f016104f5565b90565b6101ac565b346105505761053a610535366004610504565b611c22565b6105426101a2565b8061054c816102a9565b0390f35b6101a8565b9060208282031261056e5761056b915f0161045b565b90565b6101ac565b61057c906103d7565b9052565b906020806105a2936105985f8201515f860190610573565b0151910190610573565b565b91906105b7905f60408501940190610580565b565b346105e9576105e56105d46105cf366004610555565b611cfd565b6105dc6101a2565b918291826105a4565b0390f35b6101a8565b6105f7816103d7565b036105fe57565b5f80fd5b9050359061060f826105ee565b565b9060208282031261062a57610627915f01610602565b90565b6101ac565b610638906104de565b9052565b919061064f905f6020850194019061062f565b565b346106815761067d61066c610667366004610611565b611d5a565b6106746101a2565b9182918261063c565b0390f35b6101a8565b90565b61069d6106986106a29261043a565b610686565b61043a565b90565b906106af90610689565b5f5260205260405f2090565b5f1c90565b90565b6106cf6106d4916106bb565b6106c0565b90565b6106e190546106c3565b90565b6106ee905f6106a5565b9061070660016106ff5f85016106d7565b93016106d7565b90565b91602061072a92949361072360408201965f830190610322565b0190610322565b565b3461075d5761074461073f366004610555565b6106e4565b906107596107506101a2565b92839283610709565b0390f35b6101a8565b61076b9061043a565b9052565b9190610782905f60208501940190610762565b565b346107b457610794366004610379565b6107b061079f611daa565b6107a76101a2565b9182918261076f565b0390f35b6101a8565b5f80fd5b5f80fd5b5f80fd5b909182601f830112156107ff5781359167ffffffffffffffff83116107fa5760200192602083028401116107f557565b6107c1565b6107bd565b6107b9565b61080d906103cc565b90565b61081981610804565b0361082057565b5f80fd5b9050359061083182610810565b565b61083c81610201565b0361084357565b5f80fd5b9050359061085482610833565b565b9190916060818403126108a7575f8101359167ffffffffffffffff83116108a2576108868461089f9484016107c5565b9390946108968160208601610824565b93604001610847565b90565b6101b0565b6101ac565b346108de576108c86108bf366004610856565b929190916121bf565b6108d06101a2565b806108da816102a9565b0390f35b6101a8565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061090b906108e3565b810190811067ffffffffffffffff82111761092557604052565b6108ed565b9061093d6109366101a2565b9283610901565b565b67ffffffffffffffff81116109575760208091020190565b6108ed565b5f80fd5b5f80fd5b60ff1690565b61097381610964565b0361097a57565b5f80fd5b9050359061098b8261096a565b565b5f80fd5b67ffffffffffffffff81116109af576109ab6020916108e3565b0190565b6108ed565b90825f939282370152565b909291926109d46109cf82610991565b61092a565b938185526020850190828401116109f0576109ee926109b4565b565b61098d565b9080601f83011215610a1357816020610a10933591016109bf565b90565b6107b9565b91909161010081840312610ad857610a3161010061092a565b92610a3e815f840161045b565b5f850152610a4f816020840161045b565b6020850152610a61816040840161045b565b6040850152610a73816060840161045b565b6060850152610a85816080840161045b565b6080850152610a978160a0840161045b565b60a0850152610aa98160c0840161045b565b60c085015260e082013567ffffffffffffffff8111610ad357610acc92016109f5565b60e0830152565b610960565b61095c565b91909160e081840312610ba157610af460e061092a565b92610b01815f840161097e565b5f850152610b12816020840161097e565b6020850152610b2481604084016104f5565b6040850152610b368160608401610602565b6060850152610b488160808401610602565b608085015260a082013567ffffffffffffffff8111610b9c5781610b6d918401610a18565b60a085015260c082013567ffffffffffffffff8111610b9757610b909201610a18565b60c0830152565b610960565b610960565b61095c565b929190610bba610bb58261093f565b61092a565b9381855260208086019202810191838311610c115781905b838210610be0575050505050565b813567ffffffffffffffff8111610c0c57602091610c018784938701610add565b815201910190610bd2565b6107b9565b6107c1565b9080601f83011215610c3457816020610c3193359101610ba6565b90565b6107b9565b90602082820312610c69575f82013567ffffffffffffffff8111610c6457610c619201610c16565b90565b6101b0565b6101ac565b34610c9c57610c86610c81366004610c39565b612272565b610c8e6101a2565b80610c98816102a9565b0390f35b6101a8565b67ffffffffffffffff1690565b610cbe906008610cc39302610388565b610ca1565b90565b90610cd19154610cae565b90565b610ce060015f90610cc6565b90565b34610d1357610cf3366004610379565b610d0f610cfe610cd4565b610d066101a2565b9182918261076f565b0390f35b6101a8565b909182601f83011215610d525781359167ffffffffffffffff8311610d4d576020019260018302840111610d4857565b6107c1565b6107bd565b6107b9565b919091604081840312610d9857610d70835f8301610602565b92602082013567ffffffffffffffff8111610d9357610d8f9201610d18565b9091565b6101b0565b6101ac565b34610db357610dad366004610d57565b91612373565b6101a8565b34610de857610dc8366004610379565b610de4610dd36123c3565b610ddb6101a2565b918291826103f0565b0390f35b6101a8565b34610e1b57610e05610e00366004610611565b612416565b610e0d6101a2565b80610e17816102a9565b0390f35b6101a8565b9190604083820312610e485780610e3c610e45925f86016104f5565b936020016104f5565b90565b6101ac565b610e61610e5b366004610e20565b906124af565b610e696101a2565b80610e73816102a9565b0390f35b34610ea757610ea3610e92610e8d366004610274565b61252f565b610e9a6101a2565b9182918261063c565b0390f35b6101a8565b9091606082840312610ee157610ede610ec7845f8501610602565b93610ed581602086016104f5565b936040016104f5565b90565b6101ac565b610efa610ef4366004610eac565b916125d9565b610f026101a2565b80610f0c816102a9565b0390f35b6fffffffffffffffffffffffffffffffff1690565b610f2e81610f10565b03610f3557565b5f80fd5b90503590610f4682610f25565b565b9190610140838203126110185761101090610f6461014061092a565b93610f71825f830161097e565b5f860152610f828260208301610602565b6020860152610f948260408301610602565b6040860152610fa68260608301610847565b6060860152610fb882608083016104f5565b6080860152610fca8260a0830161045b565b60a0860152610fdc8260c0830161045b565b60c0860152610fee8260e0830161045b565b60e0860152611001826101008301610f39565b61010086015261012001610f39565b610120830152565b61095c565b6110268161031f565b0361102d57565b5f80fd5b9050359061103e8261101d565b565b91906101c0838203126110b5576110ae9061105b60a061092a565b93611068825f8301610f48565b5f86015261107a826101408301611031565b602086015261108d8261016083016104f5565b60408601526110a08261018083016104f5565b60608601526101a0016104f5565b6080830152565b61095c565b9161020083830312611121575f83013567ffffffffffffffff811161111c57826110e59185016109f5565b926110f38360208301611040565b926101e082013567ffffffffffffffff8111611117576111139201610d18565b9091565b6101b0565b6101b0565b6101ac565b3461115a5761115661114561113c3660046110ba565b929190916126ee565b61114d6101a2565b9182918261063c565b0390f35b6101a8565b61116a366004610eac565b916128b9565b90565b61118761118261118c92611170565b610686565b61043a565b90565b6111996001611173565b90565b6111a461118f565b90565b346111d7576111b7366004610379565b6111d36111c261119c565b6111ca6101a2565b9182918261076f565b0390f35b6101a8565b909182601f830112156112165781359167ffffffffffffffff831161121157602001926020830284011161120c57565b6107c1565b6107bd565b6107b9565b908160609103126112295790565b61025d565b906060828203126112ae575f82013567ffffffffffffffff81116112a95781611258918401610d18565b929093602082013567ffffffffffffffff81116112a4578361127b9184016111dc565b929093604082013567ffffffffffffffff811161129f5761129c920161121b565b90565b6101b0565b6101b0565b6101b0565b6101ac565b6112ca6112c136600461122e565b93929092613217565b6112d26101a2565b806112dc816102a9565b0390f35b3461130e576112f86112f3366004610611565b6137a2565b6113006101a2565b8061130a816102a9565b0390f35b6101a8565b5f80fd5b5f90565b611323611317565b508061134761134163bfb4b26b60e01b63c60751f560e01b186101b4565b916101b4565b148015611393575b8015611372575b908115611362575b5090565b61136c91506137ad565b5f61135e565b508061138d61138763c60751f560e01b6101b4565b916101b4565b14611356565b50806113ae6113a863bfb4b26b60e01b6101b4565b916101b4565b1461134f565b6113c86113c36113cd926104de565b610686565b61043a565b90565b6113da90516103d7565b90565b6113f16113ec6113f6926103cc565b610686565b6103cc565b90565b611402906113dd565b90565b61140e906113f9565b90565b61141a906113dd565b90565b61142690611411565b90565b5f80fd5b60e01b90565b5f91031261143d57565b6101ac565b5f80fd5b5f80fd5b5f80fd5b903560016020038236030381121561148f57016020813591019167ffffffffffffffff821161148a57602082023603831361148557565b611446565b611442565b61144a565b60209181520190565b90565b506114af90602081019061097e565b90565b6114bb90610964565b9052565b506114ce9060208101906104f5565b90565b6114da906104de565b9052565b506114ed906020810190610602565b90565b903560016101000382360303811215611507570190565b61144a565b5061151b90602081019061045b565b90565b6115279061043a565b9052565b903560016020038236030381121561156c57016020813591019167ffffffffffffffff821161156757600182023603831361156257565b611446565b611442565b61144a565b60209181520190565b91906115948161158d8161159995611571565b80956109b4565b6108e3565b0190565b611676916116686101008201926115c26115b95f83018361150c565b5f85019061151e565b6115dc6115d2602083018361150c565b602085019061151e565b6115f66115ec604083018361150c565b604085019061151e565b611610611606606083018361150c565b606085019061151e565b61162a611620608083018361150c565b608085019061151e565b61164461163a60a083018361150c565b60a085019061151e565b61165e61165460c083018361150c565b60c085019061151e565b60e081019061152b565b9160e081850391015261157a565b90565b61173a9161172c61172160e0830161169f6116965f8701876114a0565b5f8601906114b2565b6116b96116af60208701876114a0565b60208601906114b2565b6116d36116c960408701876114bf565b60408601906114d1565b6116ed6116e360608701876114de565b6060860190610573565b6117076116fd60808701876114de565b6080860190610573565b61171460a08601866114f0565b84820360a086015261159d565b9260c08101906114f0565b9060c081840391015261159d565b90565b9061174791611679565b90565b9035600160e00382360303811215611760570190565b61144a565b60200190565b918161177691611494565b90816117876020830284019461149d565b92835f925b84841061179c5750505050505090565b90919293949560206117c76117c183856001950388526117bc8b8861174a565b61173d565b98611765565b94019401929493919061178c565b9035600161018003823603038112156117ec570190565b61144a565b62ffffff1690565b611802816117f1565b0361180957565b5f80fd5b9050359061181a826117f9565b565b5061182b90602081019061180d565b90565b611837906117f1565b9052565b61198c9161197d611971611954610180840161186561185c5f8801886114a0565b5f8701906114b2565b61187f611875602088018861181c565b602087019061182e565b61189961188f604088018861150c565b604087019061151e565b6118b36118a9606088018861150c565b606087019061151e565b6118cd6118c3608088018861150c565b608087019061151e565b6118e76118dd60a088018861150c565b60a087019061151e565b6119016118f760c08801886114de565b60c0870190610573565b61191b61191160e08801886114de565b60e0870190610573565b61193761192c6101008801886114bf565b6101008701906114d1565b61194561012087018761152b565b9086830361012088015261157a565b61196261014086018661152b565b9085830361014087015261157a565b9261016081019061152b565b9161016081850391015261157a565b90565b6119cf916119c16119b6604083016119a95f86018661144e565b908583035f87015261176b565b9260208101906117d5565b90602081840391015261183b565b90565b6119e79160208201915f81840391015261198f565b90565b6119f26101a2565b3d5f823e3d90fd5b611a26611a21611a1c6020611a16611a11466113b4565b611cfd565b016113d0565b611405565b61141d565b90631b3348f0349290929190803b15611aa457611a565f93611a6195611a4a6101a2565b9687958694859361142d565b8352600483016119d2565b03925af18015611a9f57611a73575b50565b611a92905f3d8111611a98575b611a8a8183610901565b810190611433565b5f611a70565b503d611a80565b6119ea565b611429565b5f90565b60200190565b5190565b611ac990611ac3611aa9565b50613843565b611adb611ad582611ab3565b91611aad565b2090565b611aeb611af0916106bb565b61038c565b90565b611afd9054611adf565b90565b15611b0757565b5f80fd5b90611b3a91611b3533611b2f611b29611b246004611af3565b6103d7565b916103d7565b14611b00565b611bae565b565b35611b46816105ee565b90565b90611b5390610689565b5f5260205260405f2090565b5f1b90565b90611b7560018060a01b0391611b5f565b9181191691161790565b611b8890611411565b90565b90565b90611ba3611b9e611baa92611b7f565b611b8b565b8254611b64565b9055565b6001611beb611be36020611bf195611bdd611bca5f8301611b3c565b5f611bd760068a90611b49565b01611b8e565b01611b3c565b926006611b49565b01611b8e565b565b90611bfd91611b0b565b565b916020611c20929493611c1960408201965f83019061062f565b019061062f565b565b3342611c4e7f2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a92611b7f565b92611c63611c5a6101a2565b92839283611bff565b0390a2565b611c72604061092a565b90565b5f90565b611c81611c68565b9060208083611c8e611c75565b815201611c99611c75565b81525050565b611ca7611c79565b90565b90611cb4906103d7565b9052565b90611cef611ce66001611cc9611c68565b94611ce0611cd85f8301611af3565b5f8801611caa565b01611af3565b60208401611caa565b565b611cfa90611cb8565b90565b611d14611d1991611d0c611c9f565b506006611b49565b611cf1565b90565b5f90565b90611d2a90611b7f565b5f5260205260405f2090565b90565b611d45611d4a916106bb565b611d36565b90565b611d579054611d39565b90565b611d71611d7691611d69611d1c565b506002611d20565b611d4d565b90565b5f90565b90565b611d94611d8f611d9992611d7d565b610686565b61043a565b90565b611da7616f64611d80565b90565b611db2611d79565b50611dbb611d9c565b90565b90611dd2939291611dcd6138f8565b612032565b611dda613964565b565b5090565b67ffffffffffffffff8111611df85760208091020190565b6108ed565b90611e0f611e0a83611de0565b61092a565b918252565b611e1e60a061092a565b90565b611e2c61014061092a565b90565b5f90565b5f90565b5f90565b5f90565b5f90565b611e4b611e21565b9060208080808080808080808b611e60611e2f565b815201611e6b611c75565b815201611e76611c75565b815201611e81611e33565b815201611e8c611e37565b815201611e97611e3b565b815201611ea2611e3b565b815201611ead611e3b565b815201611eb8611e3f565b815201611ec3611e3f565b81525050565b611ed1611e43565b90565b5f90565b611ee0611e14565b9060208080808086611ef0611ec9565b815201611efb611ed4565b815201611f06611e37565b815201611f11611e37565b815201611f1c611e37565b81525050565b611f2a611ed8565b90565b5f5b828110611f3b57505050565b602090611f46611f22565b8184015201611f2f565b90611f75611f5d83611dfd565b92602080611f6b8693611de0565b9201910390611f2d565b565b90565b611f8e611f89611f9392611f77565b610686565b6104de565b90565b6001611fa291016104de565b90565b634e487b7160e01b5f52603260045260245ffd5b5190565b90611fc782611fb9565b811015611fd8576020809102010190565b611fa5565b5f80fd5b5f80fd5b5f80fd5b903590600160e00381360303821215612000570190565b611fdd565b9082101561201f57602061201c9202810190611fe9565b90565b611fa5565b9061202f91016104de565b90565b9392909192612042858490611ddc565b61204b81611f50565b926120555f611f7a565b5b80612069612063856104de565b916104de565b10156120a15761209c90612097612081878390611fbd565b5182906120908c8b8691612005565b9091613a8f565b611f96565b612056565b509293919590946120b15f611f7a565b937fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f9726120db6101a2565b806120e5816102a9565b0390a16120f15f611f7a565b945b856121066121008a6104de565b916104de565b10156121945761215c9061212461211f86888a91612005565b613cc7565b5f14612162576121559061214f61213d87898b91612005565b6121488b8b90611fbd565b5190614472565b90612024565b955b611f96565b946120f3565b61218e9061218888612176888a8c91612005565b6121818c8c90611fbd565b5191614112565b90612024565b95612157565b955092509450506121a6915015610201565b6121af575b5050565b6121b891614528565b5f806121ab565b906121cb939291611dbe565b565b6122109061220b6121dd466113b4565b6122056121ff6121fa60016121f433956006611b49565b01611af3565b6103d7565b916103d7565b14611b00565b612252565b565b61222661222161222b92611170565b610686565b6103cc565b90565b61223790612212565b90565b612243906113dd565b90565b61224f9061223a565b90565b61227090612268612263600161222e565b612246565b6001916147a6565b565b61227b906121cd565b565b905090565b909182612292816122999361227d565b80936109b4565b0190565b90916122a892612282565b90565b906122bd6122b883610991565b61092a565b918252565b606090565b3d5f146122e2576122d73d6122ab565b903d5f602084013e5b565b6122ea6122c2565b906122e0565b60209181520190565b5f5b83811061230b575050905f910152565b8060209183015181850152016122fb565b61233b6123446020936123499361233281611ab3565b938480936122f0565b958691016122f9565b6108e3565b0190565b916123709261236360408201935f830190610206565b602081840391015261231c565b90565b905f9283929161238d6123846101a2565b9283928361229d565b03915af46123996122c7565b906123bb6123a56101a2565b928392632650415560e21b84526004840161234d565b0390fd5b5f90565b6123cb6123bf565b506123d66004611af3565b90565b61240790612402336123fc6123f66123f16004611af3565b6103d7565b916103d7565b14611b00565b612409565b565b612414906005611b8e565b565b61241f906123d9565b565b634e487b7160e01b5f52601160045260245ffd5b61244461244a919392936104de565b926104de565b820180921161245557565b612421565b906124665f1991611b5f565b9181191691161790565b61248461247f612489926104de565b610686565b6104de565b90565b90565b906124a461249f6124ab92612470565b61248c565b825461245a565b9055565b612500916124cf346124c96124c3856104de565b916104de565b14611b00565b6124f7826124f16124e260023390611d20565b916124ec83611d4d565b612435565b9061248f565b339190916148b5565b565b9050519061250f826104e1565b565b9060208282031261252a57612527915f01612502565b90565b6101ac565b60206125909161253d611d1c565b5061256961256461255f84612559612554466113b4565b611cfd565b016113d0565b611405565b61141d565b612585633169eb136125796101a2565b9586948593849361142d565b8352600483016119d2565b03915afa9081156125d4575f916125a6575b5090565b6125c7915060203d81116125cd575b6125bf8183610901565b810190612511565b5f6125a2565b503d6125b5565b6119ea565b916125e692919091614978565b565b6125f190611411565b90565b6125fe905161043a565b90565b90565b61261861261361261d92612601565b610686565b61043a565b90565b61262c6126329161043a565b9161043a565b029061263d8261043a565b91820361264657565b612421565b61265f61265a6126649261043a565b610686565b6104de565b90565b90565b61267e61267961268392612667565b610686565b6104de565b90565b61269161080061266a565b90565b61269e905161031f565b90565b6126aa9061031f565b90565b6126c29160208201915f81840391015261231c565b90565b906126d091036104de565b90565b6126dd90516104de565b90565b6126eb9136916109bf565b90565b916127a26127a894936126ff611d1c565b505a6127253361271f612719612714306125e8565b6103d7565b916103d7565b14611b00565b5f85015161275061274b61273b60c084016125f4565b6127456064612604565b90612620565b61264b565b925f9361275c82611ab3565b61276e6127685f611f7a565b916104de565b116127ab575b505050612785612797915a906126c5565b612791608087016126d3565b90612024565b9093929490946126e0565b91614ade565b90565b6127d7916127d1916127bf602086016113d0565b6127cb5f939293611f7a565b906149cb565b15610201565b6127e3575b8080612774565b9091506127f66127f1612686565b6149e6565b6127ff81611ab3565b61281161280b5f611f7a565b916104de565b11612829575b505061279761278560019291506127dc565b6128406020612839818901612694565b93016113d0565b909161288a6128786128727f0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb936126a1565b93611b7f565b936128816101a2565b918291826126ad565b0390a35f80612817565b6128a36128a9919392936104de565b926104de565b82039182116128b457565b612421565b6128ce6128d5926128f3945a93919091614978565b5a90612894565b6128dd6101a2565b918291637f8f5bcd60e01b83526004830161063c565b0390fd5b67ffffffffffffffff811161290c5760200290565b6108ed565b90929192612926612921826128f7565b61092a565b93602085920283019281841161295e57915b8383106129455750505050565b6020809161295384866104f5565b815201920191612938565b6107c1565b9080601f8301121561297e5761297b91600290612911565b90565b6107b9565b67ffffffffffffffff81116129985760200290565b6108ed565b909291926129b26129ad82612983565b61092a565b9360408592028301928184116129eb57915b8383106129d15750505050565b60206040916129e08486612963565b8152019201916129c4565b6107c1565b9080601f83011215612a0b57612a089160029061299d565b90565b6107b9565b909161010082840312612a4657612a43612a2c845f8501612963565b93612a3a81604086016129f0565b9360c001612963565b90565b6101ac565b612a57612a5c916106bb565b610ca1565b90565b612a699054612a4b565b90565b5090565b612a7c612a829161043a565b9161043a565b019067ffffffffffffffff8211612a9557565b612421565b612aa6612aac9161043a565b9161043a565b90039067ffffffffffffffff8211612ac057565b612421565b903590600160400381360303821215612adc570190565b611fdd565b90821015612afb576020612af89202810190612ac5565b90565b611fa5565b35612b0a8161101d565b90565b612b16906113dd565b90565b612b2290612b0d565b90565b612b2e90611411565b90565b67ffffffffffffffff8111612b465760200290565b6108ed565b612b57612b5c91612b31565b61092a565b90565b612b84612b7b92602092612b7281611ab3565b9485809361227d565b938491016122f9565b0190565b612b9191612b5f565b90565b612ba0612ba5916106bb565b612470565b90565b90565b612bbf612bba612bc492612ba8565b610686565b6104de565b90565b612bf07f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001612bab565b90565b634e487b7160e01b5f52601260045260245ffd5b612c13612c19916104de565b916104de565b908115612c24570690565b612bf3565b90612c33906104de565b9052565b90505190612c4482610833565b565b90602082820312612c5f57612c5c915f01612c37565b90565b6101ac565b50600290565b905090565b90565b90612c7f816020936114d1565b0190565b60200190565b612ca5612c9f612c9883612c64565b8094612c6a565b91612c6f565b5f915b838310612cb55750505050565b612ccb612cc56001928451612c72565b92612c83565b92019190612ca8565b50600290565b905090565b90565b905090565b612d03612cfd612cf683612c64565b8094612ce2565b91612c6f565b5f915b838310612d135750505050565b612d29612d236001928451612c72565b92612c83565b92019190612d06565b90612d3f81604093612ce7565b0190565b60200190565b612d65612d5f612d5883612cd4565b8094612cda565b91612cdf565b5f915b838310612d755750505050565b612d8b612d856001928451612d32565b92612d43565b92019190612d68565b50600190565b905090565b90565b60200190565b612dc4612dbe612db783612d94565b8094612d9a565b91612d9f565b5f915b838310612dd45750505050565b612dea612de46001928451612c72565b92612da2565b92019190612dc7565b612e2a612e3194612e2061010094989795612e1661012086019a5f870190612c89565b6040850190612d49565b60c0830190612c89565b0190612da8565b565b903590600160200381360303821215612e75570180359067ffffffffffffffff8211612e7057602001916060820236038313612e6b57565b611fe5565b611fe1565b611fdd565b5090565b67ffffffffffffffff8111612e965760208091020190565b6108ed565b90612ead612ea883612e7e565b61092a565b918252565b612ebc604061092a565b90565b612ec9606061092a565b90565b612ed4612ebf565b906020808084612ee2611e3b565b815201612eed611e3b565b815201612ef8611e3b565b81525050565b612f06612ecc565b90565b606090565b612f16612eb2565b9060208083612f23612efe565b815201612f2e612f09565b81525050565b612f3c612f0e565b90565b5f5b828110612f4d57505050565b602090612f58612f34565b8184015201612f41565b90612f87612f6f83612e9b565b92602080612f7d8693612e7e565b9201910390612f3f565b565b9190811015612f99576060020190565b611fa5565b9190606083820312612fea57612fe390612fb8606061092a565b93612fc5825f830161045b565b5f860152612fd6826020830161045b565b602086015260400161045b565b6040830152565b61095c565b612ffa903690612f9e565b90565b61301161300c61301692611f77565b610686565b6103cc565b90565b61302290612ffd565b90565b906130376130328361093f565b61092a565b918252565b61304660e061092a565b90565b61305461010061092a565b90565b606090565b613064613049565b9060208080808080808089613077611e3b565b815201613082611e3b565b81520161308d611e3b565b815201613098611e3b565b8152016130a3611e3b565b8152016130ae611e3b565b8152016130b9611e3b565b8152016130c4613057565b81525050565b6130d261305c565b90565b6130dd61303c565b906020808080808080886130ef611e2f565b8152016130fa611e2f565b815201613105611e37565b815201613110611c75565b81520161311b611c75565b8152016131266130ca565b8152016131316130ca565b81525050565b61313f6130d5565b90565b5f5b82811061315057505050565b60209061315b613137565b8184015201613144565b9061318a61317283613025565b92602080613180869361093f565b9201910390613142565b565b9035906001602003813603038212156131ce570180359067ffffffffffffffff82116131c9576020019160208202360383136131c457565b611fe5565b611fe1565b611fdd565b5190565b5190565b906131e5826131d7565b8110156131f6576020809102010190565b611fa5565b61320f61320a61321492611170565b610686565b6104de565b90565b90613229919593929590810190612a10565b91906132f46132386001612a5f565b61325e6132456001612a5f565b6132586132538c8a90612a6c565b6113b4565b90612a70565b896132c260206132bc8a61328660016132805f61327a83612a5f565b906106a5565b016106d7565b946132b66132b16132a161329c84958095612a6c565b6113b4565b6132ab6001611173565b90612a9a565b61264b565b91612ae1565b01612b00565b906132e15f6132db816132d56001612a5f565b906106a5565b016106d7565b926132ee60208d01612b00565b94614d97565b61330e6133096133046005611af3565b612b19565b612b25565b906343753b4d9392949360205f61333a6133286001612b4b565b946133316101a2565b91829182612b88565b039060025afa156137605761339c6133919260209661337d6133756133676133625f51611b5f565b612b94565b61336f612bc7565b90612c07565b5f8401612c29565b6133856101a2565b9889978896879661142d565b865260048601612df3565b03915afa801561375b576133b8915f9161372d575b5015610201565b61370a579190926133dd6133d86133d2845f810190612e33565b90612e7a565b612f62565b926133e75f611f7a565b915b8261341161340b613406613400885f810190612e33565b90612e7a565b6104de565b916104de565b10156135c45761343661343161342a865f810190612e33565b8691612f89565b612fef565b926134425f85016125f4565b61345461344e466104de565b9161264b565b141580613583575b6135605761347c613477613472604087016125f4565b61264b565b613165565b96613485611d1c565b9061348f5f611f7a565b915b826134b66134b06134ab6134a6898b90612a6c565b6113b4565b61264b565b916104de565b10156135265761351a613520916135148c61350f6135065f8d6135006134fb8f8f6134f1918f6134e892919091612ae1565b8581019061318c565b94909493016125f4565b61264b565b91614ea6565b91828591615071565b6131d3565b90612024565b92611f96565b91613491565b50959394613555939891505f61353d8985906131db565b510152602061354d8884906131db565b510152611f96565b9294929190926133e9565b6135686101a2565b636f670bbb60e11b81528061357f600482016102a9565b0390fd5b506135a45f61359e60066135988389016125f4565b90611b49565b01611af3565b6135be6135b86135b35f613019565b6103d7565b916103d7565b1461345c565b6136a493925061365b61366592613660929796976136566136016135e86001612a5f565b6135fb6135f6858790612a6c565b6113b4565b90612a70565b61360d60208901612b00565b613650602061364a868861364461363f61362f61362a8d8d612a6c565b6113b4565b6136396001611173565b90612a9a565b61264b565b91612ae1565b01612b00565b91615191565b612a6c565b6113b4565b615209565b61369d5a91613696613691604060206136878a6136815f611f7a565b906131db565b5101519301611b3c565b612246565b5f916147a6565b5a90612894565b506136af60016131fb565b5b806136cb6136c56136c0866131d7565b6104de565b916104de565b101561370557613700906136fa6136e38583906131db565b516136f45f806006930151016125f4565b90611b49565b50611f96565b6136b0565b509050565b6137126101a2565b6309bde33960e01b815280613729600482016102a9565b0390fd5b61374e915060203d8111613754575b6137468183610901565b810190612c46565b5f6133b1565b503d61373c565b6119ea565b6119ea565b6137939061378e3361378861378261377d6004611af3565b6103d7565b916103d7565b14611b00565b613795565b565b6137a0906004611b8e565b565b6137ab90613765565b565b6137b5611317565b506137cf6137c96301ffc9a760e01b6101b4565b916101b4565b1490565b903590600161010003813603038212156137eb570190565b611fdd565b93613825613840969461381b6138329561381160a08a01965f8b0190610322565b60208901906103e3565b60408701906103e3565b848203606086015261231c565b91608081840391015261231c565b90565b61384b6122c2565b506138c96138588261525c565b916138ba61386860608301611b3c565b9161387560808201611b3c565b949294906138a56138a06138956138908460a08101906137d3565b6152e7565b9260c08101906137d3565b6152e7565b916138ae6101a2565b968795602087016137f0565b60208201810382520382610901565b90565b90565b6138e36138de6138e8926138cc565b610686565b6104de565b90565b6138f560026138cf565b90565b6139026003611d4d565b61391b6139156139106138eb565b6104de565b916104de565b146139345761393261392b6138eb565b600361248f565b565b61393c6101a2565b633ee5aeb560e01b815280613953600482016102a9565b0390fd5b61396160016131fb565b90565b61397661396f613957565b600361248f565b565b906139829061031f565b9052565b6139909051610f10565b90565b6139a76139a26139ac9261043a565b610686565b610f10565b90565b6139c36139be6139c892610f10565b610686565b6104de565b90565b6effffffffffffffffffffffffffffff1690565b6139f36139ee6139f8926139cb565b610686565b6104de565b90565b613a04906113dd565b90565b613a10906139fb565b90565b613a1c90611411565b90565b60209181520190565b5f7f4141206f776e657220766572696669636174696f6e2066616c69656400000000910152565b613a5c601c602092613a1f565b613a6581613a28565b0190565b90613a8c91613a7f60408201925f83019061062f565b6020818303910152613a4f565b90565b905a905f84015192613aa28285906153cb565b613ab7613aae83611ab7565b60208701613978565b613b35613b0d613ac960c087016125f4565b613ad560a088016125f4565b17613ae260e088016125f4565b17613af9613af36101008901613986565b91613993565b17613b076101208801613986565b176139af565b613b2e613b286effffffffffffffffffffffffffffff6139df565b916104de565b1115611b00565b613ba26020613b43866154df565b93613b68613b63613b5e84613b5785615594565b9a016113d0565b613a07565b613a13565b90613b975f635c51ffd1613b8260808c9692969501611b3c565b94613b8b6101a2565b9788968795869361142d565b8352600483016103f0565b0393f18015613c6357613bbd915f91613c35575b5015610201565b613c0f5750613c0d9392613c01608093613bfa613bf1613c0695613bec613be26122c2565b9160408b01612c29565b6155a9565b60608801612c29565b5a906126c5565b612024565b9101612c29565b565b613c3190613c1b6101a2565b918291631101335b60e11b835260048301613a69565b0390fd5b613c56915060203d8111613c5c575b613c4e8183610901565b810190612c46565b5f613bb6565b503d613c44565b6119ea565b35613c728161096a565b90565b613c89613c84613c8e92611170565b610686565b610964565b90565b613c9b6001613c75565b90565b613cb2613cad613cb7926138cc565b610686565b610964565b90565b613cc46002613c9e565b90565b613ccf611317565b50613cdc60208201613c68565b613cf5613cef613cea613c91565b610964565b91610964565b14908115613d02575b5090565b613d0f9150602001613c68565b613d28613d22613d1d613cba565b610964565b91610964565b145f613cfe565b36905f90565b613d49613d44613d4e92611f77565b610686565b610964565b90565b903590600160200381360303821215613d93570180359067ffffffffffffffff8211613d8e57602001916001820236038313613d8957565b611fe5565b611fe1565b611fdd565b5f90565b63ffffffff1690565b613db9613db4613dbe92613d9c565b61142d565b6101b4565b90565b9190613ddb81613dd481613de0956122f0565b80956109b4565b6108e3565b0190565b613ded90610201565b9052565b613dfa90610f10565b9052565b9061012080613eb393613e175f8201515f8601906114b2565b613e2960208201516020860190610573565b613e3b60408201516040860190610573565b613e4d60608201516060860190613de4565b613e5f608082015160808601906114d1565b613e7160a082015160a086019061151e565b613e8360c082015160c086019061151e565b613e9560e082015160e086019061151e565b613ea9610100820151610100860190613df1565b0151910190613df1565b565b613ebe9061031f565b9052565b906101a06080613f1f93613edc5f8201515f860190613dfe565b613eef6020820151610140860190613eb5565b613f0260408201516101608601906114d1565b613f1560608201516101808601906114d1565b01519101906114d1565b565b90613f40613f4b91613f5996946102008501918583035f870152613dc1565b936020830190613ec2565b6101e081840391015261231c565b90565b61401d9161400f61400460e08301613f82613f795f8701876114a0565b5f8601906114b2565b613f9c613f9260208701876114a0565b60208601906114b2565b613fb6613fac60408701876114bf565b60408601906114d1565b613fd0613fc660608701876114de565b6060860190610573565b613fea613fe060808701876114de565b6080860190610573565b613ff760a08601866114f0565b84820360a086015261159d565b9260c08101906114f0565b9060c081840391015261159d565b90565b9291602061403c6140449360408701908782035f890152613f5c565b940190610322565b565b61406e61406361407c95936102008401908482035f86015261231c565b936020830190613ec2565b6101e081840391015261231c565b90565b63deaddead60e01b90565b61409261407f565b90565b63deadaa5160e01b90565b6140a8614095565b90565b5f7f41413935206f7574206f66206761730000000000000000000000000000000000910152565b6140df600f602092613a1f565b6140e8816140ab565b0190565b9061410f9161410260408201925f83019061062f565b60208183039101526140d2565b90565b90929161411d611d1c565b505a61413361412e606084016126d3565b6155b5565b61413b611317565b50614144611d1c565b5061426b60205f600460405199614159613d2f565b905050614167838201613c68565b61417961417385613d35565b91610964565b1483146144515761419b6141918260a08101906137d3565b60e0810190613d51565b905b6141a56122c2565b506141ae613d98565b60038311614449575b6141d06141ca63a943c00f60e01b6101b4565b916101b4565b1485146144025750506142559061421c6141eb868b01612694565b9161420e6141f76101a2565b938492878a850163a943c00f60e01b815201614020565b878201810382520382610901565b61424763d4115a018a61422f8a92613da5565b936142386101a2565b9687958a870190815201614046565b858201810382520382610901565b5b828151910182305af15f519760405215610201565b614276575b50505050565b909192939450614284611aa9565b3d6020146143f5575b806142a76142a161429c61408a565b61031f565b9161031f565b145f146142d4576142d0856142ba6101a2565b918291631101335b60e11b8352600483016140ec565b0390fd5b90919293506142f26142ec6142e76140a0565b61031f565b9161031f565b145f14614348575061431d61430b61433d925a90612894565b614317608085016126d3565b90612435565b614329604084016126d3565b92614333816155c1565b905f8491926156a8565b905b5f808080614270565b6143e36143d1849361435f60206143ef9701612694565b61436e60205f880151016113d0565b61437e614379612686565b6149e6565b916143c76143b56143af7f676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e936126a1565b93611b7f565b936143be6101a2565b918291826126ad565b0390a35a90612894565b6143dd608085016126d3565b90612435565b90600292909192614ade565b9061433f565b5060205f803e5f5161428d565b61444492506144369063d4115a0190928b9061441e8b93613da5565b946144276101a2565b9788968b880190815201613f21565b858201810382520382610901565b614256565b5080356141b7565b61446c6144628260c08101906137d3565b60e0810190613d51565b9061419d565b6144be91604091614481611d1c565b5061448e60208201613c68565b6144a76144a161449c613c91565b610964565b91610964565b145f146144c1576144b790615811565b5b016126d3565b90565b6144cd60208201613c68565b6144e66144e06144db613cba565b610964565b91610964565b146144f2575b506144b8565b6144fb90615752565b5f6144ec565b61450a90611411565b90565b6145185f809261227d565b0190565b6145259061450d565b90565b5f6145829261455c82936145578161455061454a61454588613019565b6103d7565b91614501565b1415611b00565b614501565b906145656101a2565b90816145708161451c565b03925af161457c6122c7565b50611b00565b565b60209181520190565b60200190565b6145b26145bb6020936145c0936145a981611ab3565b93848093611571565b958691016122f9565b6108e3565b0190565b61465c9160e06101008201926145e05f8201515f85019061151e565b6145f26020820151602085019061151e565b6146046040820151604085019061151e565b6146166060820151606085019061151e565b6146286080820151608085019061151e565b61463a60a082015160a085019061151e565b61464c60c082015160c085019061151e565b01519060e0818403910152614593565b90565b6146e79160c06146d660e0830161467c5f8601515f8601906114b2565b61468e602086015160208601906114b2565b6146a0604086015160408601906114d1565b6146b260608601516060860190610573565b6146c460808601516080860190610573565b60a085015184820360a08601526145c4565b9201519060c08184039101526145c4565b90565b906146f49161465f565b90565b60200190565b9061471161470a836131d3565b8092614584565b90816147226020830284019461458d565b925f915b83831061473557505050505090565b90919293946020614757614751838560019503875289516146ea565b976146f7565b9301930191939290614726565b61476d90610804565b9052565b60409061479d6147926147a49597969460608401908482035f8601526146fd565b966020830190614764565b0190610206565b565b906147b0306125e8565b6377627f0c92919392813b1561482c575f6147de916147e982966147d26101a2565b9889978896879561142d565b855260048501614771565b03925af18015614827576147fb575b50565b61481a905f3d8111614820575b6148128183610901565b810190611433565b5f6147f8565b503d614808565b6119ea565b611429565b60601b90565b61484090614831565b90565b61484c90614837565b90565b61485b614860916103d7565b614843565b9052565b90565b614873614878916104de565b614864565b9052565b926148a960206148b1946148a1601488614899859b9a869961484f565b018092614867565b018092614867565b018092614867565b0190565b916148e66148cd6148c860028690611d20565b611d4d565b6148df6148d9856104de565b916104de565b1015611b00565b61491483614905469385906148f96101a2565b9586946020860161487c565b60208201810382520382610901565b61492661492082611ab3565b91611aad565b209190914261495e6149587f081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc936126a1565b93611b7f565b9361497361496a6101a2565b92839283611bff565b0390a3565b916149c9926149993461499361498d866104de565b916104de565b14611b00565b6149c1836149bb6149ac60028590611d20565b916149b683611d4d565b612435565b9061248f565b9190916148b5565b565b925f93918493926149da611317565b5060208451940192f190565b6149ee6122c2565b503d90808211614a13575b50604051906020810182016040528082525f602083013e90565b90505f6149f9565b90614a26910161043a565b90565b90565b614a40614a3b614a4592614a29565b610686565b6104de565b90565b614a52600a614a2c565b90565b90614a6091026104de565b90565b614a77614a72614a7c92612601565b610686565b6104de565b90565b614a8b614a91916104de565b916104de565b908115614a9c570490565b612bf3565b634e487b7160e01b5f52602160045260245ffd5b60031115614abf57565b614aa1565b90614ace82614ab5565b565b90614ada90610964565b9052565b91509291614aea611d1c565b505a614af46123bf565b50614b63614b5e614b2a5f88015194614b24614b0f876158c6565b95614b1c602089016113d0565b505a906126c5565b90612024565b93614b5860e0614b51614b3f60a085016125f4565b614b4b60c086016125f4565b90614a1b565b92016125f4565b90614a1b565b61264b565b614b7983614b73608089016126d3565b906126c5565b9080614b8d614b87846104de565b916104de565b11614c4e575b5050614b9f9082614a55565b91614bac604086016126d3565b9283614bc0614bba836104de565b916104de565b105f14614c0c5750614bdb614bd56002614ac4565b91614ac4565b145f14614bfd57614bfa9193614bf0816155c1565b905f8591926156a8565b5b565b63deadaa5160e01b5f5260205ffd5b949190614c499350614c26614c205f614ac4565b91614ac4565b1491614c40614c356001613c75565b5f8084015101614ad0565b918591926156a8565b614bfb565b92614c84614c74614c66614c8a94614b9f96976126c5565b614c6e614a48565b90614a55565b614c7e6064614a63565b90614a7f565b90612024565b91905f614b93565b614ca6614ca1614cab92611f77565b610686565b61043a565b90565b614cc2614cbd614cc792611f77565b611b5f565b61031f565b90565b90565b614cd9614cde9161031f565b614cca565b9052565b60c01b90565b614cf190614ce2565b90565b614d00614d059161043a565b614ce8565b9052565b96614d7b6008809c9b98614d7382614d8b9a614d6b614d939f9a60209f9c60209d8e8693614d5b614d839f839f90614d53601487614d4b614d6399899661484f565b018092614ccd565b018092614ccd565b018092614cf4565b018092614cf4565b018092614cf4565b018092614ccd565b018092614ccd565b018092614ccd565b018092614cf4565b0190565b90614e419295614e509594614daa6122c2565b50614df4614db6611aa9565b85614dc9614dc35f614c92565b9161043a565b148015614e53575b614dda90611b00565b614ded614de7849261031f565b9161031f565b1415611b00565b73f39fd6e51aad88f6f4ce6ab8827279cfffb9226695979391614e176001611173565b614e1f61118f565b919293614e2b5f614cae565b9596614e356101a2565b9b8c9a60208c01614d09565b60208201810382520382610901565b90565b50614dda8a614e6a614e648461031f565b9161031f565b14159050614dd1565b606090565b90614e82826131d3565b811015614e93576020809102010190565b611fa5565b614ea3903690610add565b90565b90929192614eb2614e73565b50614ebc5f611f7a565b90614ed0614ecb848390611ddc565b613165565b614ed95f611f7a565b5b80614ef7614ef1614eec888790611ddc565b6104de565b916104de565b1015614f8357614f1c6020614f16614f1188878691612005565b615939565b016125f4565b614f2e614f28896104de565b9161264b565b14614f42575b614f3d90611f96565b614eda565b92614f7b614f3d91614f74614f5988878991612005565b85614f648492614e98565b614f6e8383614e78565b52614e78565b5150611f96565b939050614f34565b509250509250614f9283613165565b91614f9c5f611f7a565b5b80614fb0614faa876104de565b916104de565b1015614fe857614fe390614fdc614fc8858390614e78565b51868391614fd68383614e78565b52614e78565b5150611f96565b614f9d565b5092505090565b5f7f54617267657420617272617920746f6f20736d616c6c00000000000000000000910152565b6150236016602092613a1f565b61502c81614fef565b0190565b6150459060208101905f818303910152615016565b90565b1561504f57565b6150576101a2565b62461bcd60e51b81528061506d60048201615030565b0390fd5b9291906150ab615080856131d3565b6150a461509e61509986615093876131d3565b90612435565b6104de565b916104de565b1015615048565b6150b45f611f7a565b5b806150d06150ca6150c5856131d3565b6104de565b916104de565b10156151125761510d906151066150e8848390614e78565b51876150f5878590612435565b916151008383614e78565b52614e78565b5150611f96565b6150b5565b5050509050565b615123604061092a565b90565b61512f906106bb565b90565b9061514761514261514e926126a1565b615126565b825461245a565b9055565b9061517d60206001615183946151755f820161516f5f8801612694565b90615132565b019201612694565b90615132565b565b9061518f91615152565b565b906151c7926151bb6151c292916151b26151a9615119565b935f8501613978565b60208301613978565b915f6106a5565b615185565b565b906151dc67ffffffffffffffff91611b5f565b9181191691161790565b90565b906151fe6151f961520592610689565b6151e6565b82546151c9565b9055565b61521f6152269161521a6001612a5f565b612a70565b60016151e9565b565b35615232816104e1565b90565b60018060f81b031690565b61525461524f615259926104de565b610686565b615235565b90565b615264611aa9565b50615285615280604061527960208501613c68565b9301615228565b615240565b9060f81b1790565b3561529781610447565b90565b909594926152e5946152d46152de926152ca6080966152c060a088019c5f890190610322565b6020870190610322565b6040850190610322565b606083019061062f565b0190610322565b565b6152ef6122c2565b506152f86122c2565b906153056020820161528d565b6153176153115f614c92565b9161043a565b03615321575b5090565b905061539b61532f82615994565b9161538c61534a6153448360e0810190613d51565b906159d3565b91615354816159e9565b9461537261536c6153676080850161528d565b61264b565b92615a29565b9093959190916153806101a2565b9687956020870161529a565b60208201810382520382610901565b5f61531d565b906153ab90610201565b9052565b906153b99061043a565b9052565b906153c790610f10565b9052565b9060e06154ba60606154226154c1956153f16153e8848301611b3c565b60208801611caa565b61540961540060808301611b3c565b60408801611caa565b61541d61541582615a69565b8488016153a1565b615939565b615442615439615434602084016125f4565b61264b565b60808701612c29565b61545a615451608083016125f4565b60a087016153af565b61547b61547161546c60a084016125f4565b613993565b61010087016153bd565b61549c61549261548d60c084016125f4565b613993565b61012087016153bd565b6154b46154ab604083016125f4565b60c087016153af565b016125f4565b91016153af565b565b906154ce9102610f10565b90565b906154dc9101610f10565b90565b61555f615564916154ee611d1c565b5061555961553161551661550460c085016125f4565b61551060a086016125f4565b90614a1b565b61552c6155266101008601613986565b91613993565b6154c3565b9161555461554e61012061554760e085016125f4565b9301613986565b91613993565b6154c3565b906154d1565b6139af565b90565b90565b61557e61557961558392615567565b610686565b6104de565b90565b61559161271061556a565b90565b5061559d611d1c565b506155a6615586565b90565b6155b1611d1c565b5090565b6155bd6122c2565b5090565b6155db60205f6155d2828501612694565b930151016113d0565b61560e6156087f1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff28936126a1565b91611b7f565b916156176101a2565b80615621816102a9565b0390a3565b6156309051610964565b90565b61563d9051610201565b90565b61564990610964565b9052565b919461569561569f9298979561568b60a0966156816156a69a61567760c08a019e5f8b01906103e3565b6020890190615640565b6040870190610206565b6060850190610206565b608083019061062f565b019061062f565b565b9290926156b760208201612694565b916157416156ca60205f850151016113d0565b946156da60405f860151016113d0565b926156f860605f6156ef81808a015101615626565b97015101615633565b9791909161572f6157297f6072b7fd64d9718ff132984edd6caa7702fcca448873489af52aba3349477f33986126a1565b98611b7f565b986157386101a2565b9687968761564d565b0390a3565b61574f90611411565b90565b61576a604061576360608401611b3c565b9201615228565b9061579161577730615746565b3161578a615784856104de565b916104de565b1015611b00565b6157cd5f806157a76157a285612246565b614501565b856157b06101a2565b90816157bb8161451c565b03925af16157c76122c7565b50611b00565b61580c6157fa7fe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f92611b7f565b926158036101a2565b9182918261063c565b0390a2565b615829604061582260608401611b3c565b9201615228565b9061585a61584161583c60028490611d20565b611d4d565b61585361584d856104de565b916104de565b1015611b00565b6158828261587c61586d60028590611d20565b9161587783611d4d565b612894565b9061248f565b6158c16158af7f9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d92611b7f565b926158b86101a2565b9182918261063c565b0390a2565b61590b906158d2611d1c565b506158df608082016126d3565b6158f16158eb466104de565b916104de565b145f1461590e576101006159059101613986565b5b6139af565b90565b61012061591b9101613986565b615906565b61592861305c565b90565b615936903690610a18565b90565b61597e90615945615920565b506159515f8201613c68565b61596361595d5f613d35565b91610964565b145f14615981576159789060a08101906137d3565b5b61592b565b90565b61598f9060c08101906137d3565b615979565b6159d0906159a0611aa9565b506159ca6159c46159be60206159b75f860161528d565b940161528d565b9261264b565b9161264b565b90615b16565b90565b906159dc611aa9565b5060405190809282372090565b615a26906159f5611aa9565b50615a20615a1a615a146060615a0d6040860161528d565b940161528d565b9261264b565b9161264b565b90615b16565b90565b615a6690615a35611aa9565b50615a60615a5a615a5460c0615a4d60a0860161528d565b940161528d565b9261264b565b9161264b565b90615b16565b90565b6020615a85615a8b92615a7a611317565b5060c08101906137d3565b0161528d565b615a9d615a975f614c92565b9161043a565b145f14615aa9575f5b90565b6001615aa6565b90565b615ac7615ac2615acc92615ab0565b610686565b610964565b90565b1b90565b615af290615aec615ae6615af794610964565b916104de565b90615acf565b6104de565b90565b615b0e615b09615b13926104de565b611b5f565b61031f565b90565b90615b91615b9792615b26611aa9565b50615b5481615b4d615b476fffffffffffffffffffffffffffffffff6139af565b916104de565b1115611b00565b615b8183615b7a615b746fffffffffffffffffffffffffffffffff6139af565b916104de565b1115611b00565b615b8b6080615ab3565b90615ad3565b17615afa565b9056fea2646970667358221220d7684b8e9b7b62979f039cf2e67a03a93a8e4e8ea001d96deb869805e233593a64736f6c63430008180033",
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

// EstimateCrossMessageParamsCrossGas is a free data retrieval call binding the contract method 0xb2241600.
//
// Solidity: function estimateCrossMessageParamsCrossGas(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256)
func (_EntryPoint *EntryPointCaller) EstimateCrossMessageParamsCrossGas(opts *bind.CallOpts, params BaseStructCrossMessageParams) (*big.Int, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "estimateCrossMessageParamsCrossGas", params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCrossMessageParamsCrossGas is a free data retrieval call binding the contract method 0xb2241600.
//
// Solidity: function estimateCrossMessageParamsCrossGas(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256)
func (_EntryPoint *EntryPointSession) EstimateCrossMessageParamsCrossGas(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _EntryPoint.Contract.EstimateCrossMessageParamsCrossGas(&_EntryPoint.CallOpts, params)
}

// EstimateCrossMessageParamsCrossGas is a free data retrieval call binding the contract method 0xb2241600.
//
// Solidity: function estimateCrossMessageParamsCrossGas(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256)
func (_EntryPoint *EntryPointCallerSession) EstimateCrossMessageParamsCrossGas(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _EntryPoint.Contract.EstimateCrossMessageParamsCrossGas(&_EntryPoint.CallOpts, params)
}

// GetChainConfigs is a free data retrieval call binding the contract method 0x42f1ec69.
//
// Solidity: function getChainConfigs(uint64 chainId) view returns((address,address))
func (_EntryPoint *EntryPointCaller) GetChainConfigs(opts *bind.CallOpts, chainId uint64) (IConfigManagerConfig, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getChainConfigs", chainId)

	if err != nil {
		return *new(IConfigManagerConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IConfigManagerConfig)).(*IConfigManagerConfig)

	return out0, err

}

// GetChainConfigs is a free data retrieval call binding the contract method 0x42f1ec69.
//
// Solidity: function getChainConfigs(uint64 chainId) view returns((address,address))
func (_EntryPoint *EntryPointSession) GetChainConfigs(chainId uint64) (IConfigManagerConfig, error) {
	return _EntryPoint.Contract.GetChainConfigs(&_EntryPoint.CallOpts, chainId)
}

// GetChainConfigs is a free data retrieval call binding the contract method 0x42f1ec69.
//
// Solidity: function getChainConfigs(uint64 chainId) view returns((address,address))
func (_EntryPoint *EntryPointCallerSession) GetChainConfigs(chainId uint64) (IConfigManagerConfig, error) {
	return _EntryPoint.Contract.GetChainConfigs(&_EntryPoint.CallOpts, chainId)
}

// GetMainChainId is a free data retrieval call binding the contract method 0x750c67ba.
//
// Solidity: function getMainChainId() pure returns(uint64)
func (_EntryPoint *EntryPointCaller) GetMainChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getMainChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetMainChainId is a free data retrieval call binding the contract method 0x750c67ba.
//
// Solidity: function getMainChainId() pure returns(uint64)
func (_EntryPoint *EntryPointSession) GetMainChainId() (uint64, error) {
	return _EntryPoint.Contract.GetMainChainId(&_EntryPoint.CallOpts)
}

// GetMainChainId is a free data retrieval call binding the contract method 0x750c67ba.
//
// Solidity: function getMainChainId() pure returns(uint64)
func (_EntryPoint *EntryPointCallerSession) GetMainChainId() (uint64, error) {
	return _EntryPoint.Contract.GetMainChainId(&_EntryPoint.CallOpts)
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

// GetUserOpHash is a free data retrieval call binding the contract method 0x2885a84f.
//
// Solidity: function getUserOpHash((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) pure returns(bytes32)
func (_EntryPoint *EntryPointCaller) GetUserOpHash(opts *bind.CallOpts, userOp BaseStructPackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0x2885a84f.
//
// Solidity: function getUserOpHash((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) pure returns(bytes32)
func (_EntryPoint *EntryPointSession) GetUserOpHash(userOp BaseStructPackedUserOperation) ([32]byte, error) {
	return _EntryPoint.Contract.GetUserOpHash(&_EntryPoint.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x2885a84f.
//
// Solidity: function getUserOpHash((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) pure returns(bytes32)
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

// EstimateSubmitDepositOperationByRemoteGas is a paid mutator transaction binding the contract method 0xdf08ec31.
//
// Solidity: function estimateSubmitDepositOperationByRemoteGas(address sender, uint256 amount, uint256 nonce) payable returns()
func (_EntryPoint *EntryPointTransactor) EstimateSubmitDepositOperationByRemoteGas(opts *bind.TransactOpts, sender common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "estimateSubmitDepositOperationByRemoteGas", sender, amount, nonce)
}

// EstimateSubmitDepositOperationByRemoteGas is a paid mutator transaction binding the contract method 0xdf08ec31.
//
// Solidity: function estimateSubmitDepositOperationByRemoteGas(address sender, uint256 amount, uint256 nonce) payable returns()
func (_EntryPoint *EntryPointSession) EstimateSubmitDepositOperationByRemoteGas(sender common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.EstimateSubmitDepositOperationByRemoteGas(&_EntryPoint.TransactOpts, sender, amount, nonce)
}

// EstimateSubmitDepositOperationByRemoteGas is a paid mutator transaction binding the contract method 0xdf08ec31.
//
// Solidity: function estimateSubmitDepositOperationByRemoteGas(address sender, uint256 amount, uint256 nonce) payable returns()
func (_EntryPoint *EntryPointTransactorSession) EstimateSubmitDepositOperationByRemoteGas(sender common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.EstimateSubmitDepositOperationByRemoteGas(&_EntryPoint.TransactOpts, sender, amount, nonce)
}

// HandleOps is a paid mutator transaction binding the contract method 0x77627f0c.
//
// Solidity: function handleOps((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] ops, address beneficiary, bool isSync) returns()
func (_EntryPoint *EntryPointTransactor) HandleOps(opts *bind.TransactOpts, ops []BaseStructPackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "handleOps", ops, beneficiary, isSync)
}

// HandleOps is a paid mutator transaction binding the contract method 0x77627f0c.
//
// Solidity: function handleOps((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] ops, address beneficiary, bool isSync) returns()
func (_EntryPoint *EntryPointSession) HandleOps(ops []BaseStructPackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, beneficiary, isSync)
}

// HandleOps is a paid mutator transaction binding the contract method 0x77627f0c.
//
// Solidity: function handleOps((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] ops, address beneficiary, bool isSync) returns()
func (_EntryPoint *EntryPointTransactorSession) HandleOps(ops []BaseStructPackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, beneficiary, isSync)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0xd4115a01.
//
// Solidity: function innerHandleOp(bytes callData, ((uint8,address,address,bool,uint256,uint64,uint64,uint64,uint128,uint128),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactor) InnerHandleOp(opts *bind.TransactOpts, callData []byte, opInfo BaseStructUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "innerHandleOp", callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0xd4115a01.
//
// Solidity: function innerHandleOp(bytes callData, ((uint8,address,address,bool,uint256,uint64,uint64,uint64,uint128,uint128),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointSession) InnerHandleOp(callData []byte, opInfo BaseStructUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0xd4115a01.
//
// Solidity: function innerHandleOp(bytes callData, ((uint8,address,address,bool,uint256,uint64,uint64,uint64,uint128,uint128),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactorSession) InnerHandleOp(callData []byte, opInfo BaseStructUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_EntryPoint *EntryPointTransactor) SendUserOmniMessage(opts *bind.TransactOpts, params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "sendUserOmniMessage", params)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_EntryPoint *EntryPointSession) SendUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _EntryPoint.Contract.SendUserOmniMessage(&_EntryPoint.TransactOpts, params)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_EntryPoint *EntryPointTransactorSession) SendUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _EntryPoint.Contract.SendUserOmniMessage(&_EntryPoint.TransactOpts, params)
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

// SubmitDepositOperationByRemote is a paid mutator transaction binding the contract method 0xbb8ce526.
//
// Solidity: function submitDepositOperationByRemote(address sender, uint256 amount, uint256 nonce) payable returns()
func (_EntryPoint *EntryPointTransactor) SubmitDepositOperationByRemote(opts *bind.TransactOpts, sender common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "submitDepositOperationByRemote", sender, amount, nonce)
}

// SubmitDepositOperationByRemote is a paid mutator transaction binding the contract method 0xbb8ce526.
//
// Solidity: function submitDepositOperationByRemote(address sender, uint256 amount, uint256 nonce) payable returns()
func (_EntryPoint *EntryPointSession) SubmitDepositOperationByRemote(sender common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.SubmitDepositOperationByRemote(&_EntryPoint.TransactOpts, sender, amount, nonce)
}

// SubmitDepositOperationByRemote is a paid mutator transaction binding the contract method 0xbb8ce526.
//
// Solidity: function submitDepositOperationByRemote(address sender, uint256 amount, uint256 nonce) payable returns()
func (_EntryPoint *EntryPointTransactorSession) SubmitDepositOperationByRemote(sender common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.SubmitDepositOperationByRemote(&_EntryPoint.TransactOpts, sender, amount, nonce)
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

// SyncBatches is a paid mutator transaction binding the contract method 0x77c2719c.
//
// Solidity: function syncBatches((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOps) returns()
func (_EntryPoint *EntryPointTransactor) SyncBatches(opts *bind.TransactOpts, userOps []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "syncBatches", userOps)
}

// SyncBatches is a paid mutator transaction binding the contract method 0x77c2719c.
//
// Solidity: function syncBatches((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOps) returns()
func (_EntryPoint *EntryPointSession) SyncBatches(userOps []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _EntryPoint.Contract.SyncBatches(&_EntryPoint.TransactOpts, userOps)
}

// SyncBatches is a paid mutator transaction binding the contract method 0x77c2719c.
//
// Solidity: function syncBatches((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOps) returns()
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

// UpdateChainConfig is a paid mutator transaction binding the contract method 0x35999cb3.
//
// Solidity: function updateChainConfig(uint64 _chainId, (address,address) _config) returns()
func (_EntryPoint *EntryPointTransactor) UpdateChainConfig(opts *bind.TransactOpts, _chainId uint64, _config IConfigManagerConfig) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "updateChainConfig", _chainId, _config)
}

// UpdateChainConfig is a paid mutator transaction binding the contract method 0x35999cb3.
//
// Solidity: function updateChainConfig(uint64 _chainId, (address,address) _config) returns()
func (_EntryPoint *EntryPointSession) UpdateChainConfig(_chainId uint64, _config IConfigManagerConfig) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateChainConfig(&_EntryPoint.TransactOpts, _chainId, _config)
}

// UpdateChainConfig is a paid mutator transaction binding the contract method 0x35999cb3.
//
// Solidity: function updateChainConfig(uint64 _chainId, (address,address) _config) returns()
func (_EntryPoint *EntryPointTransactorSession) UpdateChainConfig(_chainId uint64, _config IConfigManagerConfig) (*types.Transaction, error) {
	return _EntryPoint.Contract.UpdateChainConfig(&_EntryPoint.TransactOpts, _chainId, _config)
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

// VerifyBatches is a paid mutator transaction binding the contract method 0xe9ab53d1.
//
// Solidity: function verifyBatches(bytes proof, ((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],bytes32)[] batches, ((uint64,uint64,uint64)[],bytes32,address) chainsExecuteInfos) payable returns()
func (_EntryPoint *EntryPointTransactor) VerifyBatches(opts *bind.TransactOpts, proof []byte, batches []BaseStructBatchData, chainsExecuteInfos BaseStructChainsExecuteInfo) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "verifyBatches", proof, batches, chainsExecuteInfos)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0xe9ab53d1.
//
// Solidity: function verifyBatches(bytes proof, ((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],bytes32)[] batches, ((uint64,uint64,uint64)[],bytes32,address) chainsExecuteInfos) payable returns()
func (_EntryPoint *EntryPointSession) VerifyBatches(proof []byte, batches []BaseStructBatchData, chainsExecuteInfos BaseStructChainsExecuteInfo) (*types.Transaction, error) {
	return _EntryPoint.Contract.VerifyBatches(&_EntryPoint.TransactOpts, proof, batches, chainsExecuteInfos)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0xe9ab53d1.
//
// Solidity: function verifyBatches(bytes proof, ((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],bytes32)[] batches, ((uint64,uint64,uint64)[],bytes32,address) chainsExecuteInfos) payable returns()
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
	Owner         common.Address
	Phase         uint8
	InnerExec     bool
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserOperationEvent is a free log retrieval operation binding the contract event 0x6072b7fd64d9718ff132984edd6caa7702fcca448873489af52aba3349477f33.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address owner, uint8 phase, bool innerExec, bool success, uint256 actualGasCost, uint256 actualGasUsed)
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

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0x6072b7fd64d9718ff132984edd6caa7702fcca448873489af52aba3349477f33.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address owner, uint8 phase, bool innerExec, bool success, uint256 actualGasCost, uint256 actualGasUsed)
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

// ParseUserOperationEvent is a log parse operation binding the contract event 0x6072b7fd64d9718ff132984edd6caa7702fcca448873489af52aba3349477f33.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address owner, uint8 phase, bool innerExec, bool success, uint256 actualGasCost, uint256 actualGasUsed)
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

// ExecMetaData contains all meta data concerning the Exec contract.
var ExecMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x608060405234601d57600e6021565b603e602c823930815050603e90f35b6027565b60405190565b5f80fdfe60806040525f80fdfea26469706673582212201a0d7e416c3e9ad15c9bf613e5eb28050f07f732616ec6eebf48eb32b066362264736f6c63430008180033",
}

// ExecABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecMetaData.ABI instead.
var ExecABI = ExecMetaData.ABI

// ExecBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecMetaData.Bin instead.
var ExecBin = ExecMetaData.Bin

// DeployExec deploys a new Ethereum contract, binding an instance of Exec to it.
func DeployExec(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Exec, error) {
	parsed, err := ExecMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Exec{ExecCaller: ExecCaller{contract: contract}, ExecTransactor: ExecTransactor{contract: contract}, ExecFilterer: ExecFilterer{contract: contract}}, nil
}

// Exec is an auto generated Go binding around an Ethereum contract.
type Exec struct {
	ExecCaller     // Read-only binding to the contract
	ExecTransactor // Write-only binding to the contract
	ExecFilterer   // Log filterer for contract events
}

// ExecCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecSession struct {
	Contract     *Exec             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecCallerSession struct {
	Contract *ExecCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ExecTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecTransactorSession struct {
	Contract     *ExecTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecRaw struct {
	Contract *Exec // Generic contract binding to access the raw methods on
}

// ExecCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecCallerRaw struct {
	Contract *ExecCaller // Generic read-only contract binding to access the raw methods on
}

// ExecTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecTransactorRaw struct {
	Contract *ExecTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExec creates a new instance of Exec, bound to a specific deployed contract.
func NewExec(address common.Address, backend bind.ContractBackend) (*Exec, error) {
	contract, err := bindExec(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exec{ExecCaller: ExecCaller{contract: contract}, ExecTransactor: ExecTransactor{contract: contract}, ExecFilterer: ExecFilterer{contract: contract}}, nil
}

// NewExecCaller creates a new read-only instance of Exec, bound to a specific deployed contract.
func NewExecCaller(address common.Address, caller bind.ContractCaller) (*ExecCaller, error) {
	contract, err := bindExec(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecCaller{contract: contract}, nil
}

// NewExecTransactor creates a new write-only instance of Exec, bound to a specific deployed contract.
func NewExecTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecTransactor, error) {
	contract, err := bindExec(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecTransactor{contract: contract}, nil
}

// NewExecFilterer creates a new log filterer instance of Exec, bound to a specific deployed contract.
func NewExecFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecFilterer, error) {
	contract, err := bindExec(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecFilterer{contract: contract}, nil
}

// bindExec binds a generic wrapper to an already deployed contract.
func bindExec(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exec *ExecRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exec.Contract.ExecCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exec *ExecRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exec.Contract.ExecTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exec *ExecRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exec.Contract.ExecTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exec *ExecCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exec.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exec *ExecTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exec.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exec *ExecTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exec.Contract.contract.Transact(opts, method, params...)
}

// GoldilocksPoseidonMetaData contains all meta data concerning the GoldilocksPoseidon contract.
var GoldilocksPoseidonMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"encoded\",\"type\":\"bytes32\"}],\"name\":\"decodeHashOut\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"output\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"value\",\"type\":\"uint256[4]\"}],\"name\":\"encodeHashOut\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"encoded\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"input\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numOutputs\",\"type\":\"uint256\"}],\"name\":\"hashNToMNoPad\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"output\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"input\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[4]\",\"name\":\"capacity\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"numOutputs\",\"type\":\"uint256\"}],\"name\":\"hashNToMWithCap\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"output\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[12]\",\"name\":\"state\",\"type\":\"uint256[12]\"}],\"name\":\"permute\",\"outputs\":[{\"internalType\":\"uint256[12]\",\"name\":\"newState\",\"type\":\"uint256[12]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"right\",\"type\":\"bytes32\"}],\"name\":\"twoToOne\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"output\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052346200002957620000146200002f565b6201e1c16200003a8239308150506201e1c190f35b62000035565b60405190565b5f80fdfe6080604052600436101562000015575b62000795565b620000215f356200008c565b806310e2e35814620000865780631a10dabc1462000080578063496f3b22146200007a57806349abd3541462000074578063bbb62147146200006e5763f8978579036200000f5762000762565b6200070d565b62000635565b6200057e565b62000408565b62000387565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90620000ce90620000a4565b810190811067ffffffffffffffff821117620000e957604052565b620000ae565b9062000106620000fe62000092565b9283620000c2565b565b67ffffffffffffffff8111620001215760208091020190565b620000ae565b5f80fd5b90565b62000139816200012b565b036200014157565b5f80fd5b9050359062000154826200012e565b565b909291926200016f620001698262000108565b620000ef565b9381855260208086019202830192818411620001b157915b838310620001955750505050565b60208091620001a5848662000145565b81520192019162000187565b62000127565b9080601f83011215620001d857816020620001d59335910162000156565b90565b620000a0565b67ffffffffffffffff8111620001f45760200290565b620000ae565b90929192620002136200020d82620001de565b620000ef565b9360208592028301928184116200025057915b838310620002345750505050565b6020809162000244848662000145565b81520192019162000226565b62000127565b9080601f8301121562000274576200027191600490620001fa565b90565b620000a0565b909160c082840312620002d0575f82013567ffffffffffffffff8111620002ca57620002ad84620002c7928501620001b7565b93620002bd816020860162000256565b9360a00162000145565b90565b6200009c565b62000098565b5190565b60209181520190565b60200190565b620002f4906200012b565b9052565b906200030781602093620002e9565b0190565b60200190565b90620003346200032d6200032584620002d6565b8093620002da565b92620002e3565b905f5b818110620003455750505090565b909192620003626200035b6001928651620002f8565b946200030b565b910191909162000337565b620003849160208201915f81840391015262000311565b90565b620003b7620003a36200039c3660046200027a565b916200082e565b620003ad62000092565b918291826200036d565b0390f35b919060408382031262000402575f8301359067ffffffffffffffff8211620003fc57620003ef81620003f9938601620001b7565b9360200162000145565b90565b6200009c565b62000098565b62000438620004246200041d366004620003bb565b90620008cf565b6200042e62000092565b918291826200036d565b0390f35b67ffffffffffffffff8111620004525760200290565b620000ae565b90929192620004716200046b826200043c565b620000ef565b936020859202830192818411620004ae57915b838310620004925750505050565b60208091620004a2848662000145565b81520192019162000484565b62000127565b9080601f83011215620004d257620004cf91600c9062000458565b90565b620000a0565b9061018082820312620004f557620004f2915f01620004b4565b90565b62000098565b50600c90565b905090565b90565b60200190565b620005316200052a6200052283620004fb565b809462000501565b9162000506565b5f915b838310620005425750505050565b6200055c620005556001928451620002f8565b9262000509565b9201919062000534565b91906200057c905f61018085019401906200050f565b565b620005ad6200059962000593366004620004d8565b62000a12565b620005a362000092565b9182918262000566565b0390f35b90565b620005bf81620005b1565b03620005c757565b5f80fd5b90503590620005da82620005b4565b565b9190604083820312620006095780620005fc62000606925f8601620005cb565b93602001620005cb565b90565b62000098565b6200061a90620005b1565b9052565b919062000633905f602085019401906200060f565b565b62000665620006516200064a366004620005dc565b9062000c1f565b6200065b62000092565b918291826200061e565b0390f35b90602082820312620006855762000682915f01620005cb565b90565b62000098565b50600490565b905090565b90565b60200190565b620006c1620006ba620006b2836200068b565b809462000691565b9162000696565b5f915b838310620006d25750505050565b620006ec620006e56001928451620002f8565b9262000699565b92019190620006c4565b91906200070b905f608085019401906200069f565b565b6200073c620007286200072236600462000669565b620010da565b6200073262000092565b91829182620006f6565b0390f35b906080828203126200075c5762000759915f0162000256565b90565b62000098565b620007916200077d6200077736600462000740565b6200132f565b6200078762000092565b918291826200061e565b0390f35b5f80fd5b606090565b90565b90565b620007bd620007b7620007c3926200079e565b620007a1565b6200012b565b90565b6001620007d491016200012b565b90565b634e487b7160e01b5f52603260045260245ffd5b90620007f782620002d6565b81101562000809576020809102010190565b620007d7565b6200081b90516200012b565b90565b906200082a906200012b565b9052565b92906200083a62000799565b50620008465f620007a4565b5b8062000868620008616200085b88620002d6565b6200012b565b916200012b565b1015620008ba57620008b490620008ae62000899620008936200088d898590620007eb565b6200080f565b6200143f565b620008a88891849092620007eb565b6200081e565b620007c6565b62000847565b50909291620008cc9291909162001659565b90565b91620008da62000799565b50620008e65f620007a4565b5b806200090862000901620008fb87620002d6565b6200012b565b916200012b565b10156200095a5762000954906200094e62000939620009336200092d888590620007eb565b6200080f565b6200143f565b620009488791849092620007eb565b6200081e565b620007c6565b620008e7565b5091906200096891620019f0565b90565b6200097a62000980916200043c565b620000ef565b90565b369037565b90620009ac62000998836200096b565b92620009a584916200043c565b9062000983565b565b620009ba600c62000988565b90565b90565b620009d9620009d3620009df92620009bd565b620007a1565b6200012b565b90565b620009ee600c620009c0565b90565b90620009fd82620004fb565b81101562000a0c576020020190565b620007d7565b9062000a2862000a21620009ae565b9262002452565b62000a335f620007a4565b5b8062000a5462000a4d62000a47620009e2565b6200012b565b916200012b565b101562000aa65762000aa09062000a9a62000a8562000a7f62000a79868590620009f1565b6200080f565b6200143f565b62000a948791849092620009f1565b6200081e565b620007c6565b62000a34565b5050565b5f90565b9062000aba826200068b565b81101562000ac9576020020190565b620007d7565b90565b62000aeb62000ae562000af19262000acf565b620007a1565b6200012b565b90565b90565b62000b1062000b0a62000b169262000af4565b620007a1565b6200012b565b90565b90565b62000b3562000b2f62000b3b9262000b19565b620007a1565b6200012b565b90565b90565b62000b5a62000b5462000b609262000b3e565b620007a1565b6200012b565b90565b90565b62000b7f62000b7962000b859262000b63565b620007a1565b6200012b565b90565b90565b62000ba462000b9e62000baa9262000b88565b620007a1565b6200012b565b90565b90565b62000bc962000bc362000bcf9262000bad565b620007a1565b6200012b565b90565b62000be162000be791620001de565b620000ef565b90565b9062000c0e62000bfa8362000bd2565b9262000c078491620001de565b9062000983565b565b62000c1c600462000bea565b90565b62000fbb9162000e6f62000e5162000e4b62000c5562000c4e62000e759662000c4762000aaa565b50620010da565b94620010da565b62000d6d62000d4f62000d4962000c6b620009ae565b9762000cae62000c9162000c8b8362000c845f620007a4565b9062000aae565b6200080f565b62000ca88b62000ca15f620007a4565b90620009f1565b6200081e565b62000cf262000cd462000cce8362000cc7600162000ad2565b9062000aae565b6200080f565b62000cec8b62000ce5600162000ad2565b90620009f1565b6200081e565b62000d3662000d1862000d128362000d0b600262000af7565b9062000aae565b6200080f565b62000d308b62000d29600262000af7565b90620009f1565b6200081e565b62000d42600362000b1c565b9062000aae565b6200080f565b62000d678762000d60600362000b1c565b90620009f1565b6200081e565b62000db062000d9262000d8c8362000d855f620007a4565b9062000aae565b6200080f565b62000daa8762000da3600462000b41565b90620009f1565b6200081e565b62000df462000dd662000dd08362000dc9600162000ad2565b9062000aae565b6200080f565b62000dee8762000de7600562000b66565b90620009f1565b6200081e565b62000e3862000e1a62000e148362000e0d600262000af7565b9062000aae565b6200080f565b62000e328762000e2b600662000b8b565b90620009f1565b6200081e565b62000e44600362000b1c565b9062000aae565b6200080f565b62000e698362000e62600762000bb0565b90620009f1565b6200081e565b62002452565b62000fb562000f9762000f9162000f8b62000e8f62000c10565b9462000edc62000ebf62000eb962000eb38462000eac5f620007a4565b90620009f1565b6200080f565b6200143f565b62000ed68862000ecf5f620007a4565b9062000aae565b6200081e565b62000f2a62000f0c62000f0662000f008462000ef9600162000ad2565b90620009f1565b6200080f565b6200143f565b62000f248862000f1d600162000ad2565b9062000aae565b6200081e565b62000f7862000f5a62000f5462000f4e8462000f47600262000af7565b90620009f1565b6200080f565b6200143f565b62000f728862000f6b600262000af7565b9062000aae565b6200081e565b62000f84600362000b1c565b90620009f1565b6200080f565b6200143f565b62000faf8362000fa8600362000b1c565b9062000aae565b6200081e565b6200132f565b90565b62000fc990620005b1565b9052565b919062000fe2905f6020850194019062000fbe565b565b9050519062000ff3826200012e565b565b9060208282031262001011576200100e915f0162000fe4565b90565b62000098565b5190565b90565b62001037620010316200103d926200101b565b620007a1565b6200012b565b90565b634e487b7160e01b5f52601260045260245ffd5b620010636200106a916200012b565b916200012b565b90811562001076570690565b62001040565b90565b60ff1690565b6200109e62001098620010a4926200107c565b620007a1565b6200107f565b90565b1c90565b620010d190620010ca620010c3620010d7946200107f565b916200012b565b90620010a7565b6200012b565b90565b906200125e62001240620011eb620011966200114262001119620011296200110162000c10565b986200110c62000092565b9283916020830162000fcd565b60208201810382520382620000c2565b6020620011368262001017565b81830101910162000ff5565b6200118362001166826200115f680100000000000000006200101e565b9062001054565b6200117d89620011765f620007a4565b9062000aae565b6200081e565b6200118f604062001085565b90620010ab565b620011d8620011ba82620011b3680100000000000000006200101e565b9062001054565b620011d288620011cb600162000ad2565b9062000aae565b6200081e565b620011e4604062001085565b90620010ab565b6200122d6200120f8262001208680100000000000000006200101e565b9062001054565b620012278762001220600262000af7565b9062000aae565b6200081e565b62001239604062001085565b90620010ab565b620012588462001251600362000b1c565b9062000aae565b6200081e565b565b67ffffffffffffffff1690565b62001286620012806200128c926200012b565b620007a1565b62001260565b90565b60c01b90565b620012a0906200128f565b90565b620012b2620012b89162001260565b62001295565b9052565b92620012ef6008620012f894620012e68288620012dd829b9a8399620012a3565b018092620012a3565b018092620012a3565b018092620012a3565b0190565b905051906200130b82620005b4565b565b90602082820312620013295762001326915f01620012fc565b90565b62000098565b62001438906200133e62000aaa565b506200141f6200136f6200136962001363846200135c600362000b1c565b9062000aae565b6200080f565b6200126d565b6200140f6200139f6200139962001393866200138c600262000af7565b9062000aae565b6200080f565b6200126d565b93620013f7620013f1620013eb620013d8620013d2620013cc86620013c5600162000ad2565b9062000aae565b6200080f565b6200126d565b93620013e45f620007a4565b9062000aae565b6200080f565b6200126d565b906200140262000092565b95869460208601620012bc565b60208201810382520382620000c2565b60206200142c8262001017565b8183010191016200130d565b90565b5f90565b67ffffffff0000000190620014536200143b565b500690565b90565b620014746200146e6200147a9262001458565b620007a1565b6200012b565b90565b60209181520190565b5f7f746f6f206d616e79206f75747075747300000000000000000000000000000000910152565b620014bc60106020926200147d565b620014c78162001486565b0190565b620014e29060208101905f818303910152620014ad565b90565b15620014ed57565b620014f762000092565b62461bcd60e51b8152806200150f60048201620014cb565b0390fd5b9062001529620015238362000108565b620000ef565b918252565b90620015596200153e8362001513565b926020806200154e869362000108565b920191039062000983565b565b634e487b7160e01b5f52601160045260245ffd5b6200158162001588919392936200012b565b926200012b565b82039182116200159457565b6200155b565b620015a660086200145b565b90565b620015b8620015bf916200012b565b916200012b565b908115620015cb570490565b62001040565b620015e3620015ea919392936200012b565b926200012b565b8201809211620015f657565b6200155b565b62001607906200012b565b5f198114620016165760010190565b6200155b565b6200162e62001635919392936200012b565b926200012b565b91620016438382026200012b565b9281840414901517156200165357565b6200155b565b9291926200166662000799565b506200168a84620016836200167c60086200145b565b916200012b565b10620014e5565b6200169581620002d6565b620016ab620016a45f620007a4565b916200012b565b14620019df57620016e8620016d7620016c483620002d6565b620016d0600162000ad2565b906200156f565b620016e16200159a565b90620015a9565b936200173762001724620017136200170085620002d6565b6200170c600162000ad2565b906200156f565b6200171d6200159a565b9062001054565b62001730600162000ad2565b90620015d1565b9462001742620009ae565b956200174e5f620007a4565b5b806200176762001760600462000b41565b916200012b565b1015620017c657620017c090620017ba6200178e6200178889849062000aae565b6200080f565b620017b48b91620017ac85620017a560086200145b565b90620015d1565b9092620009f1565b6200081e565b620015fc565b6200174f565b509350939193620017d75f620007a4565b935b84620017f0620017e9846200012b565b916200012b565b1015620018b457620018025f620007a4565b5b80620018236200181c620018166200159a565b6200012b565b916200012b565b10156200188c576200188690620018806200186b620018658a6200185e620018568c6200184f6200159a565b906200161c565b8690620015d1565b90620007eb565b6200080f565b6200187a8b91849092620009f1565b6200081e565b620015fc565b62001803565b50919395620018a2620018a99194929462002452565b96620015fc565b9391929092620017d9565b94929190959350620018c65f620007a4565b5b80620018de620018d7886200012b565b916200012b565b1015620019475762001941906200193b62001926620019208762001919620019118d6200190a6200159a565b906200161c565b8690620015d1565b90620007eb565b6200080f565b620019358891849092620009f1565b6200081e565b620007c6565b620018c7565b509350935050620019589062002452565b62001963826200152e565b926200196f5f620007a4565b5b806200198762001980866200012b565b916200012b565b1015620019d957620019d390620019cd620019b8620019b2620019ac878590620009f1565b6200080f565b6200143f565b620019c78891849092620007eb565b6200081e565b620007c6565b62001970565b50915050565b505090620019ed906200152e565b90565b91620019fb62000799565b5062001a1f8262001a1862001a1160086200145b565b916200012b565b10620014e5565b62001a2a83620002d6565b62001a4062001a395f620007a4565b916200012b565b1462001f6d5762001a7d62001a6c62001a5985620002d6565b62001a65600162000ad2565b906200156f565b62001a766200159a565b90620015a9565b9262001acc62001ab962001aa862001a9584620002d6565b62001aa1600162000ad2565b906200156f565b62001ab26200159a565b9062001054565b62001ac5600162000ad2565b90620015d1565b62001ad6620009ae565b9262001ae25f620007a4565b935b8462001afb62001af4896200012b565b916200012b565b101562001e455762001e3762001e3e9162001b6e62001b5162001b4b8862001b4462001b328c62001b2b6200159a565b906200161c565b62001b3d5f620007a4565b90620015d1565b90620007eb565b6200080f565b62001b688362001b615f620007a4565b90620009f1565b6200081e565b62001bd362001bb562001baf8862001ba862001b958c62001b8e6200159a565b906200161c565b62001ba1600162000ad2565b90620015d1565b90620007eb565b6200080f565b62001bcd8362001bc6600162000ad2565b90620009f1565b6200081e565b62001c3862001c1a62001c148862001c0d62001bfa8c62001bf36200159a565b906200161c565b62001c06600262000af7565b90620015d1565b90620007eb565b6200080f565b62001c328362001c2b600262000af7565b90620009f1565b6200081e565b62001c9d62001c7f62001c798862001c7262001c5f8c62001c586200159a565b906200161c565b62001c6b600362000b1c565b90620015d1565b90620007eb565b6200080f565b62001c978362001c90600362000b1c565b90620009f1565b6200081e565b62001d0262001ce462001cde8862001cd762001cc48c62001cbd6200159a565b906200161c565b62001cd0600462000b41565b90620015d1565b90620007eb565b6200080f565b62001cfc8362001cf5600462000b41565b90620009f1565b6200081e565b62001d6762001d4962001d438862001d3c62001d298c62001d226200159a565b906200161c565b62001d35600562000b66565b90620015d1565b90620007eb565b6200080f565b62001d618362001d5a600562000b66565b90620009f1565b6200081e565b62001dcc62001dae62001da88862001da162001d8e8c62001d876200159a565b906200161c565b62001d9a600662000b8b565b90620015d1565b90620007eb565b6200080f565b62001dc68362001dbf600662000b8b565b90620009f1565b6200081e565b62001e3162001e1362001e0d8862001e0662001df38c62001dec6200159a565b906200161c565b62001dff600762000bb0565b90620015d1565b90620007eb565b6200080f565b62001e2b8362001e24600762000bb0565b90620009f1565b6200081e565b62002452565b94620007c6565b9362001ae4565b9350909362001e545f620007a4565b5b8062001e6c62001e65886200012b565b916200012b565b101562001ed55762001ecf9062001ec962001eb462001eae8762001ea762001e9f8d62001e986200159a565b906200161c565b8690620015d1565b90620007eb565b6200080f565b62001ec38891849092620009f1565b6200081e565b620007c6565b62001e55565b50935093505062001ee69062002452565b62001ef1826200152e565b9262001efd5f620007a4565b5b8062001f1562001f0e866200012b565b916200012b565b101562001f675762001f619062001f5b62001f4662001f4062001f3a878590620009f1565b6200080f565b6200143f565b62001f558891849092620007eb565b6200081e565b620007c6565b62001efe565b50915050565b915062001f7a906200152e565b90565b90565b62001f9962001f9362001f9f9262001f7d565b620007a1565b6200012b565b90565b9062001faf91016200012b565b90565b90565b62001fce62001fc862001fd49262001fb2565b620007a1565b6200012b565b90565b90565b62001ff362001fed62001ff99262001fd7565b620007a1565b6200012b565b90565b90565b62002018620020126200201e9262001ffc565b620007a1565b6200012b565b90565b90565b6200203d62002037620020439262002021565b620007a1565b6200012b565b90565b90565b620020626200205c620020689262002046565b620007a1565b6200012b565b90565b90565b62002087620020816200208d926200206b565b620007a1565b6200012b565b90565b90565b620020ac620020a6620020b29262002090565b620007a1565b6200012b565b90565b90565b620020d1620020cb620020d792620020b5565b620007a1565b6200012b565b90565b90565b620020f6620020f0620020fc92620020da565b620007a1565b6200012b565b90565b90565b6200211b620021156200212192620020ff565b620007a1565b6200012b565b90565b90565b620021406200213a620021469262002124565b620007a1565b6200012b565b90565b90565b620021656200215f6200216b9262002149565b620007a1565b6200012b565b90565b90565b6200218a6200218462002190926200216e565b620007a1565b6200012b565b90565b90565b620021af620021a9620021b59262002193565b620007a1565b6200012b565b90565b90565b620021d4620021ce620021da92620021b8565b620007a1565b6200012b565b90565b90565b620021f9620021f3620021ff92620021dd565b620007a1565b6200012b565b90565b90565b6200221e62002218620022249262002202565b620007a1565b6200012b565b90565b90565b620022436200223d620022499262002227565b620007a1565b6200012b565b90565b90565b62002268620022626200226e926200224c565b620007a1565b6200012b565b90565b90565b6200228d62002287620022939262002271565b620007a1565b6200012b565b90565b90565b620022b2620022ac620022b89262002296565b620007a1565b6200012b565b90565b90565b620022d7620022d1620022dd92620022bb565b620007a1565b6200012b565b90565b90565b620022fc620022f66200230292620022e0565b620007a1565b6200012b565b90565b90565b620023216200231b620023279262002305565b620007a1565b6200012b565b90565b90565b62002346620023406200234c926200232a565b620007a1565b6200012b565b90565b90565b6200236b6200236562002371926200234f565b620007a1565b6200012b565b90565b90565b620023906200238a620023969262002374565b620007a1565b6200012b565b90565b90565b620023b5620023af620023bb9262002399565b620007a1565b6200012b565b90565b90565b620023da620023d4620023e092620023be565b620007a1565b6200012b565b90565b90565b620023ff620023f96200240592620023e3565b620007a1565b6200012b565b90565b90565b620024246200241e6200242a9262002408565b620007a1565b6200012b565b90565b90565b62002449620024436200244f926200242d565b620007a1565b6200012b565b90565b62003078620030656200305f6200304c62003046620030336200302d6200301a62002fa562002f2862002eab62002e2e62002db162002d3462002cb76200307e9f62002bbd62002b4062002ac362002a46620029c96200294c620028cf62002852620027d562002758620026db6200265e620025e1620025656200255f62002c3a9f62002546620025406200252d62002527620025146200250e6200255996620024fb620009ae565b50620025075f620007a4565b9062003091565b62003661565b62002520600162000ad2565b9062003091565b62003661565b62002539600262000af7565b9062003091565b62003661565b62002552600362000b1c565b9062003091565b62003661565b62003ef3565b620025cf620025b262002598620025926200258c85620025855f620007a4565b90620009f1565b6200080f565b62004519565b620025ab6774cb2e819ae421ab62001f80565b9062001fa2565b620025c983620025c25f620007a4565b90620009f1565b6200081e565b620025da5f620007a4565b9062008a7a565b6200264b6200262e620026146200260e6200260885620026015f620007a4565b90620009f1565b6200080f565b62004519565b6200262767d2559d2370e7f66362001fb5565b9062001fa2565b62002645836200263e5f620007a4565b90620009f1565b6200081e565b62002657600162000ad2565b9062008a7a565b620026c8620026ab620026916200268b62002685856200267e5f620007a4565b90620009f1565b6200080f565b62004519565b620026a46762bf78acf843d17c62001fda565b9062001fa2565b620026c283620026bb5f620007a4565b90620009f1565b6200081e565b620026d4600262000af7565b9062008a7a565b62002745620027286200270e620027086200270285620026fb5f620007a4565b90620009f1565b6200080f565b62004519565b6200272167d5ab7b67e14d1fb462001fff565b9062001fa2565b6200273f83620027385f620007a4565b90620009f1565b6200081e565b62002751600362000b1c565b9062008a7a565b620027c2620027a56200278b620027856200277f85620027785f620007a4565b90620009f1565b6200080f565b62004519565b6200279e67b9fe2ae6e0969bdc62002024565b9062001fa2565b620027bc83620027b55f620007a4565b90620009f1565b6200081e565b620027ce600462000b41565b9062008a7a565b6200283f620028226200280862002802620027fc85620027f55f620007a4565b90620009f1565b6200080f565b62004519565b6200281b67e33fdf79f92a10e862002049565b9062001fa2565b6200283983620028325f620007a4565b90620009f1565b6200081e565b6200284b600562000b66565b9062008a7a565b620028bc6200289f620028856200287f6200287985620028725f620007a4565b90620009f1565b6200080f565b62004519565b62002898670ea2bb4c2b25989b6200206e565b9062001fa2565b620028b683620028af5f620007a4565b90620009f1565b6200081e565b620028c8600662000b8b565b9062008a7a565b620029396200291c62002902620028fc620028f685620028ef5f620007a4565b90620009f1565b6200080f565b62004519565b6200291567ca9121fbf9d38f0662002093565b9062001fa2565b62002933836200292c5f620007a4565b90620009f1565b6200081e565b62002945600762000bb0565b9062008a7a565b620029b6620029996200297f6200297962002973856200296c5f620007a4565b90620009f1565b6200080f565b62004519565b6200299267bdd9b0aa81f58fa4620020b8565b9062001fa2565b620029b083620029a95f620007a4565b90620009f1565b6200081e565b620029c260086200145b565b9062008a7a565b62002a3362002a16620029fc620029f6620029f085620029e95f620007a4565b90620009f1565b6200080f565b62004519565b62002a0f6783079fa4ecf20d7e620020dd565b9062001fa2565b62002a2d8362002a265f620007a4565b90620009f1565b6200081e565b62002a3f600962002102565b9062008a7a565b62002ab062002a9362002a7962002a7362002a6d8562002a665f620007a4565b90620009f1565b6200080f565b62004519565b62002a8c67650b838edfcc4ad362002127565b9062001fa2565b62002aaa8362002aa35f620007a4565b90620009f1565b6200081e565b62002abc600a6200214c565b9062008a7a565b62002b2d62002b1062002af662002af062002aea8562002ae35f620007a4565b90620009f1565b6200080f565b62004519565b62002b096777180c88583c76ac62002171565b9062001fa2565b62002b278362002b205f620007a4565b90620009f1565b6200081e565b62002b39600b62002196565b9062008a7a565b62002baa62002b8d62002b7362002b6d62002b678562002b605f620007a4565b90620009f1565b6200080f565b62004519565b62002b8667af8c20753143a180620021bb565b9062001fa2565b62002ba48362002b9d5f620007a4565b90620009f1565b6200081e565b62002bb6600c620009c0565b9062008a7a565b62002c2762002c0a62002bf062002bea62002be48562002bdd5f620007a4565b90620009f1565b6200080f565b62004519565b62002c0367b8ccfe9989a39175620021e0565b9062001fa2565b62002c218362002c1a5f620007a4565b90620009f1565b6200081e565b62002c33600d62002205565b9062008a7a565b62002ca462002c8762002c6d62002c6762002c618562002c5a5f620007a4565b90620009f1565b6200080f565b62004519565b62002c8067954a1729f60cc9c56200222a565b9062001fa2565b62002c9e8362002c975f620007a4565b90620009f1565b6200081e565b62002cb0600e6200224f565b9062008a7a565b62002d2162002d0462002cea62002ce462002cde8562002cd75f620007a4565b90620009f1565b6200080f565b62004519565b62002cfd67deb5b550c4dca53b62002274565b9062001fa2565b62002d1b8362002d145f620007a4565b90620009f1565b6200081e565b62002d2d600f62002299565b9062008a7a565b62002d9e62002d8162002d6762002d6162002d5b8562002d545f620007a4565b90620009f1565b6200080f565b62004519565b62002d7a67f01bb0b00f77011e620022be565b9062001fa2565b62002d988362002d915f620007a4565b90620009f1565b6200081e565b62002daa6010620022e3565b9062008a7a565b62002e1b62002dfe62002de462002dde62002dd88562002dd15f620007a4565b90620009f1565b6200080f565b62004519565b62002df767a1ebb404b676afd962002308565b9062001fa2565b62002e158362002e0e5f620007a4565b90620009f1565b6200081e565b62002e2760116200232d565b9062008a7a565b62002e9862002e7b62002e6162002e5b62002e558562002e4e5f620007a4565b90620009f1565b6200080f565b62004519565b62002e7467860b6e1597a0173e62002352565b9062001fa2565b62002e928362002e8b5f620007a4565b90620009f1565b6200081e565b62002ea4601262002377565b9062008a7a565b62002f1562002ef862002ede62002ed862002ed28562002ecb5f620007a4565b90620009f1565b6200080f565b62004519565b62002ef167308bb65a036acbce6200239c565b9062001fa2565b62002f0f8362002f085f620007a4565b90620009f1565b6200081e565b62002f216013620023c1565b9062008a7a565b62002f9262002f7562002f5b62002f5562002f4f8562002f485f620007a4565b90620009f1565b6200080f565b62004519565b62002f6e671aca78f31c97c876620023e6565b9062001fa2565b62002f8c8362002f855f620007a4565b90620009f1565b6200081e565b62002f9e60146200240b565b9062008a7a565b6200300762002fea62002fd862002fd262002fcc8562002fc55f620007a4565b90620009f1565b6200080f565b62004519565b62002fe35f620007a4565b9062001fa2565b620030018362002ffa5f620007a4565b90620009f1565b6200081e565b62003013601562002430565b9062008a7a565b62003026600462000b41565b9062003091565b62003661565b6200303f600562000b66565b9062003091565b62003661565b62003058600662000b8b565b9062003091565b62003661565b62003071600762000bb0565b9062003091565b62003661565b90565b906200308e91026200012b565b90565b620036416200363b6200365f92949394620036346200362e6200361a62003614620030d0620030bf620009ae565b9a620030ca620009e2565b62003081565b94620031336200311662003110620030fe620030f885620030f15f620007a4565b90620009f1565b6200080f565b620031098a62017a9b565b9062001fa2565b6200143f565b6200312d8d620031265f620007a4565b90620009f1565b6200081e565b620031ae620031906200318a620031616200315b8562003154600162000ad2565b90620009f1565b6200080f565b620031836200317d8b62003176600162000ad2565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b620031a88d620031a1600162000ad2565b90620009f1565b6200081e565b620032296200320b62003205620031dc620031d685620031cf600262000af7565b90620009f1565b6200080f565b620031fe620031f88b620031f1600262000af7565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b620032238d6200321c600262000af7565b90620009f1565b6200081e565b620032a462003286620032806200325762003251856200324a600362000b1c565b90620009f1565b6200080f565b62003279620032738b6200326c600362000b1c565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b6200329e8d62003297600362000b1c565b90620009f1565b6200081e565b6200331f62003301620032fb620032d2620032cc85620032c5600462000b41565b90620009f1565b6200080f565b620032f4620032ee8b620032e7600462000b41565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b620033198d62003312600462000b41565b90620009f1565b6200081e565b6200339a6200337c620033766200334d620033478562003340600562000b66565b90620009f1565b6200080f565b6200336f620033698b62003362600562000b66565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b620033948d6200338d600562000b66565b90620009f1565b6200081e565b62003415620033f7620033f1620033c8620033c285620033bb600662000b8b565b90620009f1565b6200080f565b620033ea620033e48b620033dd600662000b8b565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b6200340f8d62003408600662000b8b565b90620009f1565b6200081e565b62003490620034726200346c620034436200343d8562003436600762000bb0565b90620009f1565b6200080f565b620034656200345f8b62003458600762000bb0565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b6200348a8d62003483600762000bb0565b90620009f1565b6200081e565b6200350b620034ed620034e7620034be620034b885620034b160086200145b565b90620009f1565b6200080f565b620034e0620034da8b620034d360086200145b565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b620035058d620034fe60086200145b565b90620009f1565b6200081e565b6200358662003568620035626200353962003533856200352c600962002102565b90620009f1565b6200080f565b6200355b620035558b6200354e600962002102565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b620035808d62003579600962002102565b90620009f1565b6200081e565b62003601620035e3620035dd620035b4620035ae85620035a7600a6200214c565b90620009f1565b6200080f565b620035d6620035d08b620035c9600a6200214c565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b620035fb8d620035f4600a6200214c565b90620009f1565b6200081e565b6200360d600b62002196565b90620009f1565b6200080f565b9262003627600b62002196565b9062001fa2565b62017a9b565b9062001fa2565b6200143f565b620036598462003652600b62002196565b90620009f1565b6200081e565b565b9062003c2d62003c0f62003674620009ae565b93620036c1620036a46200369e6200369884620036915f620007a4565b90620009f1565b6200080f565b62004519565b620036bb83620036b45f620007a4565b90620009f1565b6200081e565b6200370f620036f1620036eb620036e584620036de600162000ad2565b90620009f1565b6200080f565b62004519565b620037098362003702600162000ad2565b90620009f1565b6200081e565b6200375d6200373f6200373962003733846200372c600262000af7565b90620009f1565b6200080f565b62004519565b620037578362003750600262000af7565b90620009f1565b6200081e565b620037ab6200378d6200378762003781846200377a600362000b1c565b90620009f1565b6200080f565b62004519565b620037a5836200379e600362000b1c565b90620009f1565b6200081e565b620037f9620037db620037d5620037cf84620037c8600462000b41565b90620009f1565b6200080f565b62004519565b620037f383620037ec600462000b41565b90620009f1565b6200081e565b6200384762003829620038236200381d8462003816600562000b66565b90620009f1565b6200080f565b62004519565b62003841836200383a600562000b66565b90620009f1565b6200081e565b6200389562003877620038716200386b8462003864600662000b8b565b90620009f1565b6200080f565b62004519565b6200388f8362003888600662000b8b565b90620009f1565b6200081e565b620038e3620038c5620038bf620038b984620038b2600762000bb0565b90620009f1565b6200080f565b62004519565b620038dd83620038d6600762000bb0565b90620009f1565b6200081e565b62003931620039136200390d62003907846200390060086200145b565b90620009f1565b6200080f565b62004519565b6200392b836200392460086200145b565b90620009f1565b6200081e565b6200397f620039616200395b62003955846200394e600962002102565b90620009f1565b6200080f565b62004519565b620039798362003972600962002102565b90620009f1565b6200081e565b620039cd620039af620039a9620039a3846200399c600a6200214c565b90620009f1565b6200080f565b62004519565b620039c783620039c0600a6200214c565b90620009f1565b6200081e565b62003a1b620039fd620039f7620039f184620039ea600b62002196565b90620009f1565b6200080f565b62004519565b62003a158362003a0e600b62002196565b90620009f1565b6200081e565b62003a4762003a2a8262018ec2565b62003a418762003a3a5f620007a4565b90620009f1565b6200081e565b62003a7462003a5682620191f5565b62003a6e8762003a67600162000ad2565b90620009f1565b6200081e565b62003aa162003a838262019513565b62003a9b8762003a94600262000af7565b90620009f1565b6200081e565b62003ace62003ab08262019831565b62003ac88762003ac1600362000b1c565b90620009f1565b6200081e565b62003afb62003add8262019b4f565b62003af58762003aee600462000b41565b90620009f1565b6200081e565b62003b2862003b0a8262019e6d565b62003b228762003b1b600562000b66565b90620009f1565b6200081e565b62003b5562003b37826201a18b565b62003b4f8762003b48600662000b8b565b90620009f1565b6200081e565b62003b8262003b64826201a4a9565b62003b7c8762003b75600762000bb0565b90620009f1565b6200081e565b62003baf62003b91826201a7c7565b62003ba98762003ba260086200145b565b90620009f1565b6200081e565b62003bdc62003bbe826201aae5565b62003bd68762003bcf600962002102565b90620009f1565b6200081e565b62003c0962003beb826201ae03565b62003c038762003bfc600a6200214c565b90620009f1565b6200081e565b6201b121565b62003c278462003c20600b62002196565b90620009f1565b6200081e565b565b90565b62003c4b62003c4562003c519262003c2f565b620007a1565b6200012b565b90565b62003c6767e993fd841e7e97f162003c32565b90565b90565b62003c8662003c8062003c8c9262003c6a565b620007a1565b6200012b565b90565b62003ca267f2831d3575f0f3af62003c6d565b90565b90565b62003cc162003cbb62003cc79262003ca5565b620007a1565b6200012b565b90565b62003cdd67d2500e0a350994ca62003ca8565b90565b90565b62003cfc62003cf662003d029262003ce0565b620007a1565b6200012b565b90565b62003d1867c5571f35d728863362003ce3565b90565b90565b62003d3762003d3162003d3d9262003d1b565b620007a1565b6200012b565b90565b62003d536791d89c5184109a0262003d1e565b90565b90565b62003d7262003d6c62003d789262003d56565b620007a1565b6200012b565b90565b62003d8e67f37f925d04e5667b62003d59565b90565b90565b62003dad62003da762003db39262003d91565b620007a1565b6200012b565b90565b62003dc9672d6e448371955a6962003d94565b90565b90565b62003de862003de262003dee9262003dcc565b620007a1565b6200012b565b90565b62003e0467740ef19ce01398a162003dcf565b90565b90565b62003e2362003e1d62003e299262003e07565b620007a1565b6200012b565b90565b62003e3f67694d24c0752fdf4562003e0a565b90565b90565b62003e5e62003e5862003e649262003e42565b620007a1565b6200012b565b90565b62003e7a6760936af96ee2f14862003e45565b90565b90565b62003e9962003e9362003e9f9262003e7d565b620007a1565b6200012b565b90565b62003eb567c33448feadc78f0c62003e80565b90565b90565b62003ed462003ece62003eda9262003eb8565b620007a1565b6200012b565b90565b62003ef0673cc3f892184df40862003ebb565b90565b9062004517620044fa620044e9620044e362003f0e620009ae565b9562003f6862003f4a62003f3962003f338462003f2c600162000ad2565b90620009f1565b6200080f565b62003f4362003c54565b906201b43f565b62003f628362003f5b600162000ad2565b90620009f1565b6200081e565b62003fc162003fa362003f9262003f8c8462003f85600262000af7565b90620009f1565b6200080f565b62003f9c62003c8f565b906201b43f565b62003fbb8362003fb4600262000af7565b90620009f1565b6200081e565b6200401a62003ffc62003feb62003fe58462003fde600362000b1c565b90620009f1565b6200080f565b62003ff562003cca565b906201b43f565b62004014836200400d600362000b1c565b90620009f1565b6200081e565b6200407362004055620040446200403e8462004037600462000b41565b90620009f1565b6200080f565b6200404e62003d05565b906201b43f565b6200406d8362004066600462000b41565b90620009f1565b6200081e565b620040cc620040ae6200409d620040978462004090600562000b66565b90620009f1565b6200080f565b620040a762003d40565b906201b43f565b620040c683620040bf600562000b66565b90620009f1565b6200081e565b6200412562004107620040f6620040f084620040e9600662000b8b565b90620009f1565b6200080f565b6200410062003d7b565b906201b43f565b6200411f8362004118600662000b8b565b90620009f1565b6200081e565b6200417e620041606200414f620041498462004142600762000bb0565b90620009f1565b6200080f565b6200415962003db6565b906201b43f565b620041788362004171600762000bb0565b90620009f1565b6200081e565b620041d7620041b9620041a8620041a2846200419b60086200145b565b90620009f1565b6200080f565b620041b262003df1565b906201b43f565b620041d183620041ca60086200145b565b90620009f1565b6200081e565b620042306200421262004201620041fb84620041f4600962002102565b90620009f1565b6200080f565b6200420b62003e2c565b906201b43f565b6200422a8362004223600962002102565b90620009f1565b6200081e565b620042896200426b6200425a62004254846200424d600a6200214c565b90620009f1565b6200080f565b6200426462003e67565b906201b43f565b62004283836200427c600a6200214c565b90620009f1565b6200081e565b620042e2620042c4620042b3620042ad84620042a6600b62002196565b90620009f1565b6200080f565b620042bd62003ea2565b906201b43f565b620042dc83620042d5600b62002196565b90620009f1565b6200081e565b6200430f620042f1826201b5f0565b620043098962004302600162000ad2565b90620009f1565b6200081e565b6200433c6200431e826201baa0565b62004336896200432f600362000b1c565b90620009f1565b6200081e565b620043696200434b826201bf2b565b62004363896200435c600462000b41565b90620009f1565b6200081e565b6200439662004378826201c391565b620043908962004389600562000b66565b90620009f1565b6200081e565b620043c3620043a5826201c7d2565b620043bd89620043b6600662000b8b565b90620009f1565b6200081e565b620043f0620043d2826201cbed565b620043ea89620043e3600762000bb0565b90620009f1565b6200081e565b6200441d620043ff826201cfe4565b62004417896200441060086200145b565b90620009f1565b6200081e565b6200444a6200442c826201d3b6565b62004444896200443d600962002102565b90620009f1565b6200081e565b6200447762004459826201d763565b62004471896200446a600a6200214c565b90620009f1565b6200081e565b620044a462004486826201daeb565b6200449e8962004497600b62002196565b90620009f1565b6200081e565b620044d1620044b3826201de4d565b620044cb89620044c4600262000af7565b90620009f1565b6200081e565b620044dc5f620007a4565b90620009f1565b6200080f565b620044f362003edd565b906201b43f565b62004511846200450a5f620007a4565b90620009f1565b6200081e565b565b6200456890620045286200143b565b50620045336200143b565b50620045626200455b620045556200454d84859062003081565b849062003081565b6200143f565b8062003081565b62003081565b90565b6200457760116200232d565b90565b6200458660086200145b565b90565b90565b620045a56200459f620045ab9262004589565b620007a1565b6200012b565b90565b90565b620045ca620045c4620045d092620045ae565b620007a1565b6200012b565b90565b90565b620045ef620045e9620045f592620045d3565b620007a1565b6200012b565b90565b90565b620046146200460e6200461a92620045f8565b620007a1565b6200012b565b90565b90565b62004639620046336200463f926200461d565b620007a1565b6200012b565b90565b90565b6200465e62004658620046649262004642565b620007a1565b6200012b565b90565b90565b620046836200467d620046899262004667565b620007a1565b6200012b565b90565b90565b620046a8620046a2620046ae926200468c565b620007a1565b6200012b565b90565b90565b620046cd620046c7620046d392620046b1565b620007a1565b6200012b565b90565b90565b620046f2620046ec620046f892620046d6565b620007a1565b6200012b565b90565b90565b62004717620047116200471d92620046fb565b620007a1565b6200012b565b90565b90565b6200473c62004736620047429262004720565b620007a1565b6200012b565b90565b90565b620047616200475b620047679262004745565b620007a1565b6200012b565b90565b90565b62004786620047806200478c926200476a565b620007a1565b6200012b565b90565b90565b620047ab620047a5620047b1926200478f565b620007a1565b6200012b565b90565b90565b620047d0620047ca620047d692620047b4565b620007a1565b6200012b565b90565b90565b620047f5620047ef620047fb92620047d9565b620007a1565b6200012b565b90565b90565b6200481a620048146200482092620047fe565b620007a1565b6200012b565b90565b90565b6200483f62004839620048459262004823565b620007a1565b6200012b565b90565b90565b620048646200485e6200486a9262004848565b620007a1565b6200012b565b90565b90565b62004889620048836200488f926200486d565b620007a1565b6200012b565b90565b90565b620048ae620048a8620048b49262004892565b620007a1565b6200012b565b90565b90565b620048d3620048cd620048d992620048b7565b620007a1565b6200012b565b90565b90565b620048f8620048f2620048fe92620048dc565b620007a1565b6200012b565b90565b90565b6200491d62004917620049239262004901565b620007a1565b6200012b565b90565b90565b620049426200493c620049489262004926565b620007a1565b6200012b565b90565b90565b62004967620049616200496d926200494b565b620007a1565b6200012b565b90565b90565b6200498c62004986620049929262004970565b620007a1565b6200012b565b90565b90565b620049b1620049ab620049b79262004995565b620007a1565b6200012b565b90565b90565b620049d6620049d0620049dc92620049ba565b620007a1565b6200012b565b90565b90565b620049fb620049f562004a0192620049df565b620007a1565b6200012b565b90565b90565b62004a2062004a1a62004a269262004a04565b620007a1565b6200012b565b90565b90565b62004a4562004a3f62004a4b9262004a29565b620007a1565b6200012b565b90565b90565b62004a6a62004a6462004a709262004a4e565b620007a1565b6200012b565b90565b90565b62004a8f62004a8962004a959262004a73565b620007a1565b6200012b565b90565b90565b62004ab462004aae62004aba9262004a98565b620007a1565b6200012b565b90565b90565b62004ad962004ad362004adf9262004abd565b620007a1565b6200012b565b90565b90565b62004afe62004af862004b049262004ae2565b620007a1565b6200012b565b90565b90565b62004b2362004b1d62004b299262004b07565b620007a1565b6200012b565b90565b90565b62004b4862004b4262004b4e9262004b2c565b620007a1565b6200012b565b90565b90565b62004b6d62004b6762004b739262004b51565b620007a1565b6200012b565b90565b90565b62004b9262004b8c62004b989262004b76565b620007a1565b6200012b565b90565b90565b62004bb762004bb162004bbd9262004b9b565b620007a1565b6200012b565b90565b90565b62004bdc62004bd662004be29262004bc0565b620007a1565b6200012b565b90565b90565b62004c0162004bfb62004c079262004be5565b620007a1565b6200012b565b90565b90565b62004c2662004c2062004c2c9262004c0a565b620007a1565b6200012b565b90565b90565b62004c4b62004c4562004c519262004c2f565b620007a1565b6200012b565b90565b90565b62004c7062004c6a62004c769262004c54565b620007a1565b6200012b565b90565b90565b62004c9562004c8f62004c9b9262004c79565b620007a1565b6200012b565b90565b90565b62004cba62004cb462004cc09262004c9e565b620007a1565b6200012b565b90565b90565b62004cdf62004cd962004ce59262004cc3565b620007a1565b6200012b565b90565b90565b62004d0462004cfe62004d0a9262004ce8565b620007a1565b6200012b565b90565b90565b62004d2962004d2362004d2f9262004d0d565b620007a1565b6200012b565b90565b90565b62004d4e62004d4862004d549262004d32565b620007a1565b6200012b565b90565b90565b62004d7362004d6d62004d799262004d57565b620007a1565b6200012b565b90565b90565b62004d9862004d9262004d9e9262004d7c565b620007a1565b6200012b565b90565b90565b62004dbd62004db762004dc39262004da1565b620007a1565b6200012b565b90565b90565b62004de262004ddc62004de89262004dc6565b620007a1565b6200012b565b90565b90565b62004e0762004e0162004e0d9262004deb565b620007a1565b6200012b565b90565b90565b62004e2c62004e2662004e329262004e10565b620007a1565b6200012b565b90565b90565b62004e5162004e4b62004e579262004e35565b620007a1565b6200012b565b90565b90565b62004e7662004e7062004e7c9262004e5a565b620007a1565b6200012b565b90565b90565b62004e9b62004e9562004ea19262004e7f565b620007a1565b6200012b565b90565b90565b62004ec062004eba62004ec69262004ea4565b620007a1565b6200012b565b90565b90565b62004ee562004edf62004eeb9262004ec9565b620007a1565b6200012b565b90565b90565b62004f0a62004f0462004f109262004eee565b620007a1565b6200012b565b90565b90565b62004f2f62004f2962004f359262004f13565b620007a1565b6200012b565b90565b90565b62004f5462004f4e62004f5a9262004f38565b620007a1565b6200012b565b90565b90565b62004f7962004f7362004f7f9262004f5d565b620007a1565b6200012b565b90565b90565b62004f9e62004f9862004fa49262004f82565b620007a1565b6200012b565b90565b90565b62004fc362004fbd62004fc99262004fa7565b620007a1565b6200012b565b90565b90565b62004fe862004fe262004fee9262004fcc565b620007a1565b6200012b565b90565b90565b6200500d62005007620050139262004ff1565b620007a1565b6200012b565b90565b90565b620050326200502c620050389262005016565b620007a1565b6200012b565b90565b90565b62005057620050516200505d926200503b565b620007a1565b6200012b565b90565b90565b6200507c62005076620050829262005060565b620007a1565b6200012b565b90565b90565b620050a16200509b620050a79262005085565b620007a1565b6200012b565b90565b90565b620050c6620050c0620050cc92620050aa565b620007a1565b6200012b565b90565b90565b620050eb620050e5620050f192620050cf565b620007a1565b6200012b565b90565b90565b620051106200510a6200511692620050f4565b620007a1565b6200012b565b90565b90565b620051356200512f6200513b9262005119565b620007a1565b6200012b565b90565b90565b6200515a6200515462005160926200513e565b620007a1565b6200012b565b90565b90565b6200517f62005179620051859262005163565b620007a1565b6200012b565b90565b90565b620051a46200519e620051aa9262005188565b620007a1565b6200012b565b90565b90565b620051c9620051c3620051cf92620051ad565b620007a1565b6200012b565b90565b90565b620051ee620051e8620051f492620051d2565b620007a1565b6200012b565b90565b90565b620052136200520d6200521992620051f7565b620007a1565b6200012b565b90565b90565b62005238620052326200523e926200521c565b620007a1565b6200012b565b90565b90565b6200525d62005257620052639262005241565b620007a1565b6200012b565b90565b90565b620052826200527c620052889262005266565b620007a1565b6200012b565b90565b90565b620052a7620052a1620052ad926200528b565b620007a1565b6200012b565b90565b90565b620052cc620052c6620052d292620052b0565b620007a1565b6200012b565b90565b90565b620052f1620052eb620052f792620052d5565b620007a1565b6200012b565b90565b90565b62005316620053106200531c92620052fa565b620007a1565b6200012b565b90565b90565b6200533b6200533562005341926200531f565b620007a1565b6200012b565b90565b90565b620053606200535a620053669262005344565b620007a1565b6200012b565b90565b90565b620053856200537f6200538b9262005369565b620007a1565b6200012b565b90565b90565b620053aa620053a4620053b0926200538e565b620007a1565b6200012b565b90565b90565b620053cf620053c9620053d592620053b3565b620007a1565b6200012b565b90565b90565b620053f4620053ee620053fa92620053d8565b620007a1565b6200012b565b90565b90565b62005419620054136200541f92620053fd565b620007a1565b6200012b565b90565b90565b6200543e62005438620054449262005422565b620007a1565b6200012b565b90565b90565b620054636200545d620054699262005447565b620007a1565b6200012b565b90565b90565b62005488620054826200548e926200546c565b620007a1565b6200012b565b90565b90565b620054ad620054a7620054b39262005491565b620007a1565b6200012b565b90565b90565b620054d2620054cc620054d892620054b6565b620007a1565b6200012b565b90565b90565b620054f7620054f1620054fd92620054db565b620007a1565b6200012b565b90565b90565b6200551c62005516620055229262005500565b620007a1565b6200012b565b90565b90565b620055416200553b620055479262005525565b620007a1565b6200012b565b90565b90565b62005566620055606200556c926200554a565b620007a1565b6200012b565b90565b90565b6200558b6200558562005591926200556f565b620007a1565b6200012b565b90565b90565b620055b0620055aa620055b69262005594565b620007a1565b6200012b565b90565b90565b620055d5620055cf620055db92620055b9565b620007a1565b6200012b565b90565b90565b620055fa620055f46200560092620055de565b620007a1565b6200012b565b90565b90565b6200561f62005619620056259262005603565b620007a1565b6200012b565b90565b90565b620056446200563e6200564a9262005628565b620007a1565b6200012b565b90565b90565b62005669620056636200566f926200564d565b620007a1565b6200012b565b90565b90565b6200568e62005688620056949262005672565b620007a1565b6200012b565b90565b90565b620056b3620056ad620056b99262005697565b620007a1565b6200012b565b90565b90565b620056d8620056d2620056de92620056bc565b620007a1565b6200012b565b90565b90565b620056fd620056f76200570392620056e1565b620007a1565b6200012b565b90565b90565b620057226200571c620057289262005706565b620007a1565b6200012b565b90565b90565b62005747620057416200574d926200572b565b620007a1565b6200012b565b90565b90565b6200576c62005766620057729262005750565b620007a1565b6200012b565b90565b90565b620057916200578b620057979262005775565b620007a1565b6200012b565b90565b90565b620057b6620057b0620057bc926200579a565b620007a1565b6200012b565b90565b90565b620057db620057d5620057e192620057bf565b620007a1565b6200012b565b90565b90565b62005800620057fa6200580692620057e4565b620007a1565b6200012b565b90565b90565b620058256200581f6200582b9262005809565b620007a1565b6200012b565b90565b90565b6200584a6200584462005850926200582e565b620007a1565b6200012b565b90565b90565b6200586f62005869620058759262005853565b620007a1565b6200012b565b90565b90565b620058946200588e6200589a9262005878565b620007a1565b6200012b565b90565b90565b620058b9620058b3620058bf926200589d565b620007a1565b6200012b565b90565b90565b620058de620058d8620058e492620058c2565b620007a1565b6200012b565b90565b90565b62005903620058fd6200590992620058e7565b620007a1565b6200012b565b90565b90565b62005928620059226200592e926200590c565b620007a1565b6200012b565b90565b90565b6200594d62005947620059539262005931565b620007a1565b6200012b565b90565b90565b620059726200596c620059789262005956565b620007a1565b6200012b565b90565b90565b62005997620059916200599d926200597b565b620007a1565b6200012b565b90565b90565b620059bc620059b6620059c292620059a0565b620007a1565b6200012b565b90565b90565b620059e1620059db620059e792620059c5565b620007a1565b6200012b565b90565b90565b62005a0662005a0062005a0c92620059ea565b620007a1565b6200012b565b90565b90565b62005a2b62005a2562005a319262005a0f565b620007a1565b6200012b565b90565b90565b62005a5062005a4a62005a569262005a34565b620007a1565b6200012b565b90565b90565b62005a7562005a6f62005a7b9262005a59565b620007a1565b6200012b565b90565b90565b62005a9a62005a9462005aa09262005a7e565b620007a1565b6200012b565b90565b90565b62005abf62005ab962005ac59262005aa3565b620007a1565b6200012b565b90565b90565b62005ae462005ade62005aea9262005ac8565b620007a1565b6200012b565b90565b90565b62005b0962005b0362005b0f9262005aed565b620007a1565b6200012b565b90565b90565b62005b2e62005b2862005b349262005b12565b620007a1565b6200012b565b90565b90565b62005b5362005b4d62005b599262005b37565b620007a1565b6200012b565b90565b90565b62005b7862005b7262005b7e9262005b5c565b620007a1565b6200012b565b90565b90565b62005b9d62005b9762005ba39262005b81565b620007a1565b6200012b565b90565b90565b62005bc262005bbc62005bc89262005ba6565b620007a1565b6200012b565b90565b90565b62005be762005be162005bed9262005bcb565b620007a1565b6200012b565b90565b90565b62005c0c62005c0662005c129262005bf0565b620007a1565b6200012b565b90565b90565b62005c3162005c2b62005c379262005c15565b620007a1565b6200012b565b90565b90565b62005c5662005c5062005c5c9262005c3a565b620007a1565b6200012b565b90565b90565b62005c7b62005c7562005c819262005c5f565b620007a1565b6200012b565b90565b90565b62005ca062005c9a62005ca69262005c84565b620007a1565b6200012b565b90565b90565b62005cc562005cbf62005ccb9262005ca9565b620007a1565b6200012b565b90565b90565b62005cea62005ce462005cf09262005cce565b620007a1565b6200012b565b90565b90565b62005d0f62005d0962005d159262005cf3565b620007a1565b6200012b565b90565b90565b62005d3462005d2e62005d3a9262005d18565b620007a1565b6200012b565b90565b90565b62005d5962005d5362005d5f9262005d3d565b620007a1565b6200012b565b90565b90565b62005d7e62005d7862005d849262005d62565b620007a1565b6200012b565b90565b90565b62005da362005d9d62005da99262005d87565b620007a1565b6200012b565b90565b90565b62005dc862005dc262005dce9262005dac565b620007a1565b6200012b565b90565b90565b62005ded62005de762005df39262005dd1565b620007a1565b6200012b565b90565b90565b62005e1262005e0c62005e189262005df6565b620007a1565b6200012b565b90565b90565b62005e3762005e3162005e3d9262005e1b565b620007a1565b6200012b565b90565b90565b62005e5c62005e5662005e629262005e40565b620007a1565b6200012b565b90565b90565b62005e8162005e7b62005e879262005e65565b620007a1565b6200012b565b90565b90565b62005ea662005ea062005eac9262005e8a565b620007a1565b6200012b565b90565b90565b62005ecb62005ec562005ed19262005eaf565b620007a1565b6200012b565b90565b90565b62005ef062005eea62005ef69262005ed4565b620007a1565b6200012b565b90565b90565b62005f1562005f0f62005f1b9262005ef9565b620007a1565b6200012b565b90565b90565b62005f3a62005f3462005f409262005f1e565b620007a1565b6200012b565b90565b90565b62005f5f62005f5962005f659262005f43565b620007a1565b6200012b565b90565b90565b62005f8462005f7e62005f8a9262005f68565b620007a1565b6200012b565b90565b90565b62005fa962005fa362005faf9262005f8d565b620007a1565b6200012b565b90565b90565b62005fce62005fc862005fd49262005fb2565b620007a1565b6200012b565b90565b90565b62005ff362005fed62005ff99262005fd7565b620007a1565b6200012b565b90565b90565b62006018620060126200601e9262005ffc565b620007a1565b6200012b565b90565b90565b6200603d62006037620060439262006021565b620007a1565b6200012b565b90565b90565b620060626200605c620060689262006046565b620007a1565b6200012b565b90565b90565b62006087620060816200608d926200606b565b620007a1565b6200012b565b90565b90565b620060ac620060a6620060b29262006090565b620007a1565b6200012b565b90565b90565b620060d1620060cb620060d792620060b5565b620007a1565b6200012b565b90565b90565b620060f6620060f0620060fc92620060da565b620007a1565b6200012b565b90565b90565b6200611b620061156200612192620060ff565b620007a1565b6200012b565b90565b90565b620061406200613a620061469262006124565b620007a1565b6200012b565b90565b90565b620061656200615f6200616b9262006149565b620007a1565b6200012b565b90565b90565b6200618a6200618462006190926200616e565b620007a1565b6200012b565b90565b90565b620061af620061a9620061b59262006193565b620007a1565b6200012b565b90565b90565b620061d4620061ce620061da92620061b8565b620007a1565b6200012b565b90565b90565b620061f9620061f3620061ff92620061dd565b620007a1565b6200012b565b90565b90565b6200621e62006218620062249262006202565b620007a1565b6200012b565b90565b90565b620062436200623d620062499262006227565b620007a1565b6200012b565b90565b90565b62006268620062626200626e926200624c565b620007a1565b6200012b565b90565b90565b6200628d62006287620062939262006271565b620007a1565b6200012b565b90565b90565b620062b2620062ac620062b89262006296565b620007a1565b6200012b565b90565b90565b620062d7620062d1620062dd92620062bb565b620007a1565b6200012b565b90565b90565b620062fc620062f66200630292620062e0565b620007a1565b6200012b565b90565b90565b620063216200631b620063279262006305565b620007a1565b6200012b565b90565b90565b62006346620063406200634c926200632a565b620007a1565b6200012b565b90565b90565b6200636b6200636562006371926200634f565b620007a1565b6200012b565b90565b90565b620063906200638a620063969262006374565b620007a1565b6200012b565b90565b90565b620063b5620063af620063bb9262006399565b620007a1565b6200012b565b90565b90565b620063da620063d4620063e092620063be565b620007a1565b6200012b565b90565b90565b620063ff620063f96200640592620063e3565b620007a1565b6200012b565b90565b90565b620064246200641e6200642a9262006408565b620007a1565b6200012b565b90565b90565b62006449620064436200644f926200642d565b620007a1565b6200012b565b90565b90565b6200646e62006468620064749262006452565b620007a1565b6200012b565b90565b90565b620064936200648d620064999262006477565b620007a1565b6200012b565b90565b90565b620064b8620064b2620064be926200649c565b620007a1565b6200012b565b90565b90565b620064dd620064d7620064e392620064c1565b620007a1565b6200012b565b90565b90565b62006502620064fc6200650892620064e6565b620007a1565b6200012b565b90565b90565b62006527620065216200652d926200650b565b620007a1565b6200012b565b90565b90565b6200654c62006546620065529262006530565b620007a1565b6200012b565b90565b90565b620065716200656b620065779262006555565b620007a1565b6200012b565b90565b90565b62006596620065906200659c926200657a565b620007a1565b6200012b565b90565b90565b620065bb620065b5620065c1926200659f565b620007a1565b6200012b565b90565b90565b620065e0620065da620065e692620065c4565b620007a1565b6200012b565b90565b90565b62006605620065ff6200660b92620065e9565b620007a1565b6200012b565b90565b90565b6200662a6200662462006630926200660e565b620007a1565b6200012b565b90565b90565b6200664f62006649620066559262006633565b620007a1565b6200012b565b90565b90565b620066746200666e6200667a9262006658565b620007a1565b6200012b565b90565b90565b62006699620066936200669f926200667d565b620007a1565b6200012b565b90565b90565b620066be620066b8620066c492620066a2565b620007a1565b6200012b565b90565b90565b620066e3620066dd620066e992620066c7565b620007a1565b6200012b565b90565b90565b62006708620067026200670e92620066ec565b620007a1565b6200012b565b90565b90565b6200672d62006727620067339262006711565b620007a1565b6200012b565b90565b90565b620067526200674c620067589262006736565b620007a1565b6200012b565b90565b90565b62006777620067716200677d926200675b565b620007a1565b6200012b565b90565b90565b6200679c62006796620067a29262006780565b620007a1565b6200012b565b90565b90565b620067c1620067bb620067c792620067a5565b620007a1565b6200012b565b90565b90565b620067e6620067e0620067ec92620067ca565b620007a1565b6200012b565b90565b90565b6200680b620068056200681192620067ef565b620007a1565b6200012b565b90565b90565b620068306200682a620068369262006814565b620007a1565b6200012b565b90565b90565b620068556200684f6200685b9262006839565b620007a1565b6200012b565b90565b90565b6200687a6200687462006880926200685e565b620007a1565b6200012b565b90565b90565b6200689f62006899620068a59262006883565b620007a1565b6200012b565b90565b90565b620068c4620068be620068ca92620068a8565b620007a1565b6200012b565b90565b90565b620068e9620068e3620068ef92620068cd565b620007a1565b6200012b565b90565b90565b6200690e620069086200691492620068f2565b620007a1565b6200012b565b90565b90565b620069336200692d620069399262006917565b620007a1565b6200012b565b90565b90565b62006958620069526200695e926200693c565b620007a1565b6200012b565b90565b90565b6200697d62006977620069839262006961565b620007a1565b6200012b565b90565b90565b620069a26200699c620069a89262006986565b620007a1565b6200012b565b90565b90565b620069c7620069c1620069cd92620069ab565b620007a1565b6200012b565b90565b90565b620069ec620069e6620069f292620069d0565b620007a1565b6200012b565b90565b90565b62006a1162006a0b62006a1792620069f5565b620007a1565b6200012b565b90565b90565b62006a3662006a3062006a3c9262006a1a565b620007a1565b6200012b565b90565b90565b62006a5b62006a5562006a619262006a3f565b620007a1565b6200012b565b90565b90565b62006a8062006a7a62006a869262006a64565b620007a1565b6200012b565b90565b90565b62006aa562006a9f62006aab9262006a89565b620007a1565b6200012b565b90565b90565b62006aca62006ac462006ad09262006aae565b620007a1565b6200012b565b90565b90565b62006aef62006ae962006af59262006ad3565b620007a1565b6200012b565b90565b90565b62006b1462006b0e62006b1a9262006af8565b620007a1565b6200012b565b90565b90565b62006b3962006b3362006b3f9262006b1d565b620007a1565b6200012b565b90565b90565b62006b5e62006b5862006b649262006b42565b620007a1565b6200012b565b90565b90565b62006b8362006b7d62006b899262006b67565b620007a1565b6200012b565b90565b90565b62006ba862006ba262006bae9262006b8c565b620007a1565b6200012b565b90565b90565b62006bcd62006bc762006bd39262006bb1565b620007a1565b6200012b565b90565b90565b62006bf262006bec62006bf89262006bd6565b620007a1565b6200012b565b90565b90565b62006c1762006c1162006c1d9262006bfb565b620007a1565b6200012b565b90565b90565b62006c3c62006c3662006c429262006c20565b620007a1565b6200012b565b90565b90565b62006c6162006c5b62006c679262006c45565b620007a1565b6200012b565b90565b90565b62006c8662006c8062006c8c9262006c6a565b620007a1565b6200012b565b90565b90565b62006cab62006ca562006cb19262006c8f565b620007a1565b6200012b565b90565b90565b62006cd062006cca62006cd69262006cb4565b620007a1565b6200012b565b90565b90565b62006cf562006cef62006cfb9262006cd9565b620007a1565b6200012b565b90565b90565b62006d1a62006d1462006d209262006cfe565b620007a1565b6200012b565b90565b90565b62006d3f62006d3962006d459262006d23565b620007a1565b6200012b565b90565b90565b62006d6462006d5e62006d6a9262006d48565b620007a1565b6200012b565b90565b90565b62006d8962006d8362006d8f9262006d6d565b620007a1565b6200012b565b90565b90565b62006dae62006da862006db49262006d92565b620007a1565b6200012b565b90565b90565b62006dd362006dcd62006dd99262006db7565b620007a1565b6200012b565b90565b90565b62006df862006df262006dfe9262006ddc565b620007a1565b6200012b565b90565b90565b62006e1d62006e1762006e239262006e01565b620007a1565b6200012b565b90565b90565b62006e4262006e3c62006e489262006e26565b620007a1565b6200012b565b90565b90565b62006e6762006e6162006e6d9262006e4b565b620007a1565b6200012b565b90565b90565b62006e8c62006e8662006e929262006e70565b620007a1565b6200012b565b90565b90565b62006eb162006eab62006eb79262006e95565b620007a1565b6200012b565b90565b90565b62006ed662006ed062006edc9262006eba565b620007a1565b6200012b565b90565b90565b62006efb62006ef562006f019262006edf565b620007a1565b6200012b565b90565b90565b62006f2062006f1a62006f269262006f04565b620007a1565b6200012b565b90565b90565b62006f4562006f3f62006f4b9262006f29565b620007a1565b6200012b565b90565b90565b62006f6a62006f6462006f709262006f4e565b620007a1565b6200012b565b90565b90565b62006f8f62006f8962006f959262006f73565b620007a1565b6200012b565b90565b90565b62006fb462006fae62006fba9262006f98565b620007a1565b6200012b565b90565b90565b62006fd962006fd362006fdf9262006fbd565b620007a1565b6200012b565b90565b90565b62006ffe62006ff8620070049262006fe2565b620007a1565b6200012b565b90565b90565b620070236200701d620070299262007007565b620007a1565b6200012b565b90565b90565b62007048620070426200704e926200702c565b620007a1565b6200012b565b90565b90565b6200706d62007067620070739262007051565b620007a1565b6200012b565b90565b90565b620070926200708c620070989262007076565b620007a1565b6200012b565b90565b90565b620070b7620070b1620070bd926200709b565b620007a1565b6200012b565b90565b90565b620070dc620070d6620070e292620070c0565b620007a1565b6200012b565b90565b90565b62007101620070fb6200710792620070e5565b620007a1565b6200012b565b90565b90565b62007126620071206200712c926200710a565b620007a1565b6200012b565b90565b90565b6200714b6200714562007151926200712f565b620007a1565b6200012b565b90565b90565b620071706200716a620071769262007154565b620007a1565b6200012b565b90565b90565b620071956200718f6200719b9262007179565b620007a1565b6200012b565b90565b90565b620071ba620071b4620071c0926200719e565b620007a1565b6200012b565b90565b90565b620071df620071d9620071e592620071c3565b620007a1565b6200012b565b90565b90565b62007204620071fe6200720a92620071e8565b620007a1565b6200012b565b90565b90565b62007229620072236200722f926200720d565b620007a1565b6200012b565b90565b90565b6200724e62007248620072549262007232565b620007a1565b6200012b565b90565b90565b620072736200726d620072799262007257565b620007a1565b6200012b565b90565b90565b62007298620072926200729e926200727c565b620007a1565b6200012b565b90565b90565b620072bd620072b7620072c392620072a1565b620007a1565b6200012b565b90565b90565b620072e2620072dc620072e892620072c6565b620007a1565b6200012b565b90565b90565b62007307620073016200730d92620072eb565b620007a1565b6200012b565b90565b90565b6200732c62007326620073329262007310565b620007a1565b6200012b565b90565b90565b620073516200734b620073579262007335565b620007a1565b6200012b565b90565b90565b62007376620073706200737c926200735a565b620007a1565b6200012b565b90565b90565b6200739b62007395620073a1926200737f565b620007a1565b6200012b565b90565b90565b620073c0620073ba620073c692620073a4565b620007a1565b6200012b565b90565b90565b620073e5620073df620073eb92620073c9565b620007a1565b6200012b565b90565b90565b6200740a620074046200741092620073ee565b620007a1565b6200012b565b90565b90565b6200742f62007429620074359262007413565b620007a1565b6200012b565b90565b90565b620074546200744e6200745a9262007438565b620007a1565b6200012b565b90565b90565b62007479620074736200747f926200745d565b620007a1565b6200012b565b90565b90565b6200749e62007498620074a49262007482565b620007a1565b6200012b565b90565b90565b620074c3620074bd620074c992620074a7565b620007a1565b6200012b565b90565b90565b620074e8620074e2620074ee92620074cc565b620007a1565b6200012b565b90565b90565b6200750d620075076200751392620074f1565b620007a1565b6200012b565b90565b90565b620075326200752c620075389262007516565b620007a1565b6200012b565b90565b90565b62007557620075516200755d926200753b565b620007a1565b6200012b565b90565b90565b6200757c62007576620075829262007560565b620007a1565b6200012b565b90565b90565b620075a16200759b620075a79262007585565b620007a1565b6200012b565b90565b90565b620075c6620075c0620075cc92620075aa565b620007a1565b6200012b565b90565b90565b620075eb620075e5620075f192620075cf565b620007a1565b6200012b565b90565b90565b620076106200760a6200761692620075f4565b620007a1565b6200012b565b90565b90565b620076356200762f6200763b9262007619565b620007a1565b6200012b565b90565b90565b6200765a6200765462007660926200763e565b620007a1565b6200012b565b90565b90565b6200767f62007679620076859262007663565b620007a1565b6200012b565b90565b90565b620076a46200769e620076aa9262007688565b620007a1565b6200012b565b90565b90565b620076c9620076c3620076cf92620076ad565b620007a1565b6200012b565b90565b90565b620076ee620076e8620076f492620076d2565b620007a1565b6200012b565b90565b90565b620077136200770d6200771992620076f7565b620007a1565b6200012b565b90565b90565b62007738620077326200773e926200771c565b620007a1565b6200012b565b90565b90565b6200775d62007757620077639262007741565b620007a1565b6200012b565b90565b90565b620077826200777c620077889262007766565b620007a1565b6200012b565b90565b90565b620077a7620077a1620077ad926200778b565b620007a1565b6200012b565b90565b90565b620077cc620077c6620077d292620077b0565b620007a1565b6200012b565b90565b90565b620077f1620077eb620077f792620077d5565b620007a1565b6200012b565b90565b90565b62007816620078106200781c92620077fa565b620007a1565b6200012b565b90565b90565b6200783b6200783562007841926200781f565b620007a1565b6200012b565b90565b90565b620078606200785a620078669262007844565b620007a1565b6200012b565b90565b90565b620078856200787f6200788b9262007869565b620007a1565b6200012b565b90565b90565b620078aa620078a4620078b0926200788e565b620007a1565b6200012b565b90565b90565b620078cf620078c9620078d592620078b3565b620007a1565b6200012b565b90565b90565b620078f4620078ee620078fa92620078d8565b620007a1565b6200012b565b90565b90565b62007919620079136200791f92620078fd565b620007a1565b6200012b565b90565b90565b6200793e62007938620079449262007922565b620007a1565b6200012b565b90565b90565b620079636200795d620079699262007947565b620007a1565b6200012b565b90565b90565b62007988620079826200798e926200796c565b620007a1565b6200012b565b90565b90565b620079ad620079a7620079b39262007991565b620007a1565b6200012b565b90565b90565b620079d2620079cc620079d892620079b6565b620007a1565b6200012b565b90565b90565b620079f7620079f1620079fd92620079db565b620007a1565b6200012b565b90565b90565b62007a1c62007a1662007a229262007a00565b620007a1565b6200012b565b90565b90565b62007a4162007a3b62007a479262007a25565b620007a1565b6200012b565b90565b90565b62007a6662007a6062007a6c9262007a4a565b620007a1565b6200012b565b90565b90565b62007a8b62007a8562007a919262007a6f565b620007a1565b6200012b565b90565b90565b62007ab062007aaa62007ab69262007a94565b620007a1565b6200012b565b90565b90565b62007ad562007acf62007adb9262007ab9565b620007a1565b6200012b565b90565b90565b62007afa62007af462007b009262007ade565b620007a1565b6200012b565b90565b90565b62007b1f62007b1962007b259262007b03565b620007a1565b6200012b565b90565b90565b62007b4462007b3e62007b4a9262007b28565b620007a1565b6200012b565b90565b90565b62007b6962007b6362007b6f9262007b4d565b620007a1565b6200012b565b90565b90565b62007b8e62007b8862007b949262007b72565b620007a1565b6200012b565b90565b90565b62007bb362007bad62007bb99262007b97565b620007a1565b6200012b565b90565b90565b62007bd862007bd262007bde9262007bbc565b620007a1565b6200012b565b90565b90565b62007bfd62007bf762007c039262007be1565b620007a1565b6200012b565b90565b90565b62007c2262007c1c62007c289262007c06565b620007a1565b6200012b565b90565b90565b62007c4762007c4162007c4d9262007c2b565b620007a1565b6200012b565b90565b90565b62007c6c62007c6662007c729262007c50565b620007a1565b6200012b565b90565b90565b62007c9162007c8b62007c979262007c75565b620007a1565b6200012b565b90565b90565b62007cb662007cb062007cbc9262007c9a565b620007a1565b6200012b565b90565b90565b62007cdb62007cd562007ce19262007cbf565b620007a1565b6200012b565b90565b90565b62007d0062007cfa62007d069262007ce4565b620007a1565b6200012b565b90565b90565b62007d2562007d1f62007d2b9262007d09565b620007a1565b6200012b565b90565b90565b62007d4a62007d4462007d509262007d2e565b620007a1565b6200012b565b90565b90565b62007d6f62007d6962007d759262007d53565b620007a1565b6200012b565b90565b90565b62007d9462007d8e62007d9a9262007d78565b620007a1565b6200012b565b90565b90565b62007db962007db362007dbf9262007d9d565b620007a1565b6200012b565b90565b90565b62007dde62007dd862007de49262007dc2565b620007a1565b6200012b565b90565b90565b62007e0362007dfd62007e099262007de7565b620007a1565b6200012b565b90565b90565b62007e2862007e2262007e2e9262007e0c565b620007a1565b6200012b565b90565b90565b62007e4d62007e4762007e539262007e31565b620007a1565b6200012b565b90565b90565b62007e7262007e6c62007e789262007e56565b620007a1565b6200012b565b90565b90565b62007e9762007e9162007e9d9262007e7b565b620007a1565b6200012b565b90565b90565b62007ebc62007eb662007ec29262007ea0565b620007a1565b6200012b565b90565b90565b62007ee162007edb62007ee79262007ec5565b620007a1565b6200012b565b90565b90565b62007f0662007f0062007f0c9262007eea565b620007a1565b6200012b565b90565b90565b62007f2b62007f2562007f319262007f0f565b620007a1565b6200012b565b90565b90565b62007f5062007f4a62007f569262007f34565b620007a1565b6200012b565b90565b90565b62007f7562007f6f62007f7b9262007f59565b620007a1565b6200012b565b90565b90565b62007f9a62007f9462007fa09262007f7e565b620007a1565b6200012b565b90565b90565b62007fbf62007fb962007fc59262007fa3565b620007a1565b6200012b565b90565b90565b62007fe462007fde62007fea9262007fc8565b620007a1565b6200012b565b90565b90565b62008009620080036200800f9262007fed565b620007a1565b6200012b565b90565b90565b6200802e62008028620080349262008012565b620007a1565b6200012b565b90565b90565b620080536200804d620080599262008037565b620007a1565b6200012b565b90565b90565b62008078620080726200807e926200805c565b620007a1565b6200012b565b90565b90565b6200809d62008097620080a39262008081565b620007a1565b6200012b565b90565b90565b620080c2620080bc620080c892620080a6565b620007a1565b6200012b565b90565b90565b620080e7620080e1620080ed92620080cb565b620007a1565b6200012b565b90565b90565b6200810c620081066200811292620080f0565b620007a1565b6200012b565b90565b90565b620081316200812b620081379262008115565b620007a1565b6200012b565b90565b90565b62008156620081506200815c926200813a565b620007a1565b6200012b565b90565b90565b6200817b6200817562008181926200815f565b620007a1565b6200012b565b90565b90565b620081a06200819a620081a69262008184565b620007a1565b6200012b565b90565b90565b620081c5620081bf620081cb92620081a9565b620007a1565b6200012b565b90565b90565b620081ea620081e4620081f092620081ce565b620007a1565b6200012b565b90565b90565b6200820f620082096200821592620081f3565b620007a1565b6200012b565b90565b90565b620082346200822e6200823a9262008218565b620007a1565b6200012b565b90565b90565b62008259620082536200825f926200823d565b620007a1565b6200012b565b90565b90565b6200827e62008278620082849262008262565b620007a1565b6200012b565b90565b90565b620082a36200829d620082a99262008287565b620007a1565b6200012b565b90565b90565b620082c8620082c2620082ce92620082ac565b620007a1565b6200012b565b90565b90565b620082ed620082e7620082f392620082d1565b620007a1565b6200012b565b90565b90565b620083126200830c6200831892620082f6565b620007a1565b6200012b565b90565b90565b62008337620083316200833d926200831b565b620007a1565b6200012b565b90565b90565b6200835c62008356620083629262008340565b620007a1565b6200012b565b90565b90565b620083816200837b620083879262008365565b620007a1565b6200012b565b90565b90565b620083a6620083a0620083ac926200838a565b620007a1565b6200012b565b90565b90565b620083cb620083c5620083d192620083af565b620007a1565b6200012b565b90565b90565b620083f0620083ea620083f692620083d4565b620007a1565b6200012b565b90565b90565b620084156200840f6200841b92620083f9565b620007a1565b6200012b565b90565b90565b6200843a6200843462008440926200841e565b620007a1565b6200012b565b90565b90565b6200845f62008459620084659262008443565b620007a1565b6200012b565b90565b90565b620084846200847e6200848a9262008468565b620007a1565b6200012b565b90565b90565b620084a9620084a3620084af926200848d565b620007a1565b6200012b565b90565b90565b620084ce620084c8620084d492620084b2565b620007a1565b6200012b565b90565b90565b620084f3620084ed620084f992620084d7565b620007a1565b6200012b565b90565b90565b62008518620085126200851e92620084fc565b620007a1565b6200012b565b90565b90565b6200853d62008537620085439262008521565b620007a1565b6200012b565b90565b90565b620085626200855c620085689262008546565b620007a1565b6200012b565b90565b90565b62008587620085816200858d926200856b565b620007a1565b6200012b565b90565b90565b620085ac620085a6620085b29262008590565b620007a1565b6200012b565b90565b90565b620085d1620085cb620085d792620085b5565b620007a1565b6200012b565b90565b90565b620085f6620085f0620085fc92620085da565b620007a1565b6200012b565b90565b90565b6200861b620086156200862192620085ff565b620007a1565b6200012b565b90565b90565b620086406200863a620086469262008624565b620007a1565b6200012b565b90565b90565b620086656200865f6200866b9262008649565b620007a1565b6200012b565b90565b90565b6200868a6200868462008690926200866e565b620007a1565b6200012b565b90565b90565b620086af620086a9620086b59262008693565b620007a1565b6200012b565b90565b90565b620086d4620086ce620086da92620086b8565b620007a1565b6200012b565b90565b90565b620086f9620086f3620086ff92620086dd565b620007a1565b6200012b565b90565b90565b6200871e62008718620087249262008702565b620007a1565b6200012b565b90565b90565b620087436200873d620087499262008727565b620007a1565b6200012b565b90565b90565b62008768620087626200876e926200874c565b620007a1565b6200012b565b90565b90565b6200878d62008787620087939262008771565b620007a1565b6200012b565b90565b90565b620087b2620087ac620087b89262008796565b620007a1565b6200012b565b90565b90565b620087d7620087d1620087dd92620087bb565b620007a1565b6200012b565b90565b90565b620087fc620087f66200880292620087e0565b620007a1565b6200012b565b90565b90565b620088216200881b620088279262008805565b620007a1565b6200012b565b90565b90565b62008846620088406200884c926200882a565b620007a1565b6200012b565b90565b90565b6200886b6200886562008871926200884f565b620007a1565b6200012b565b90565b90565b620088906200888a620088969262008874565b620007a1565b6200012b565b90565b90565b620088b5620088af620088bb9262008899565b620007a1565b6200012b565b90565b90565b620088da620088d4620088e092620088be565b620007a1565b6200012b565b90565b90565b620088ff620088f96200890592620088e3565b620007a1565b6200012b565b90565b90565b620089246200891e6200892a9262008908565b620007a1565b6200012b565b90565b90565b62008949620089436200894f926200892d565b620007a1565b6200012b565b90565b90565b6200896e62008968620089749262008952565b620007a1565b6200012b565b90565b90565b620089936200898d620089999262008977565b620007a1565b6200012b565b90565b90565b620089b8620089b2620089be926200899c565b620007a1565b6200012b565b90565b90565b620089dd620089d7620089e392620089c1565b620007a1565b6200012b565b90565b90565b62008a02620089fc62008a0892620089e6565b620007a1565b6200012b565b90565b90565b62008a2762008a2162008a2d9262008a0b565b620007a1565b6200012b565b90565b90565b62008a4c62008a4662008a529262008a30565b620007a1565b6200012b565b90565b90565b62008a7162008a6b62008a779262008a55565b620007a1565b6200012b565b90565b919062008a86620009ae565b9262008ad262008aac62008aa68362008a9f5f620007a4565b90620009f1565b6200080f565b62008acb62008aba6200456b565b62008ac46200457a565b9062001fa2565b9062003081565b918062008aea62008ae35f620007a4565b916200012b565b145f14620094e65750620094bf8162008e6762008e4a62008e44620094e39662008e3d62008df662008daf62008d6862008d2162008cda8a62008cd362008cb9620094c59f62008c9f62008c5862008c1162008bca62008b8362008b6962008b6362008cb39762008b5c600162000ad2565b90620009f1565b6200080f565b62008b7c673d999c961b7c63b06200874f565b9062003081565b62008bc362008ba962008ba38b62008b9c600262000af7565b90620009f1565b6200080f565b62008bbc67814e82efcd17252962008774565b9062003081565b9062001fa2565b62008c0a62008bf062008bea8a62008be3600362000b1c565b90620009f1565b6200080f565b62008c03672421e5d23670458862008799565b9062003081565b9062001fa2565b62008c5162008c3762008c318962008c2a600462000b41565b90620009f1565b6200080f565b62008c4a67887af7d4dd482328620087be565b9062003081565b9062001fa2565b62008c9862008c7e62008c788862008c71600562000b66565b90620009f1565b6200080f565b62008c9167a5e9c291f6119b27620087e3565b9062003081565b9062001fa2565b9362008cac600662000b8b565b90620009f1565b6200080f565b62008ccc67bdc52b2676a4b4aa62008808565b9062003081565b9062001fa2565b62008d1a62008d0062008cfa8d62008cf3600762000bb0565b90620009f1565b6200080f565b62008d136764832009d29bcf576200882d565b9062003081565b9062001fa2565b62008d6162008d4762008d418c62008d3a60086200145b565b90620009f1565b6200080f565b62008d5a6709c4155174a552cc62008852565b9062003081565b9062001fa2565b62008da862008d8e62008d888b62008d81600962002102565b90620009f1565b6200080f565b62008da167463f9ee03d29081062008877565b9062003081565b9062001fa2565b62008def62008dd562008dcf8a62008dc8600a6200214c565b90620009f1565b6200080f565b62008de867c810936e649825426200889c565b9062003081565b9062001fa2565b62008e3662008e1c62008e168962008e0f600b62002196565b90620009f1565b6200080f565b62008e2f67043b1c289f7bc3ac620088c1565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b62008e618862008e5a5f620007a4565b90620009f1565b6200081e565b62008eff62008ee162008edb62008e9562008e8f8562008e88600162000ad2565b90620009f1565b6200080f565b62008ed462008eba62008eb48762008ead5f620007a4565b90620009f1565b6200080f565b62008ecd6794877900674181c3620088e6565b9062003081565b9062001fa2565b6200143f565b62008ef98862008ef2600162000ad2565b90620009f1565b6200081e565b62008f9762008f7962008f7362008f2d62008f278562008f20600262000af7565b90620009f1565b6200080f565b62008f6c62008f5262008f4c8762008f455f620007a4565b90620009f1565b6200080f565b62008f6567c6c67cc37a2a2bbd6200890b565b9062003081565b9062001fa2565b6200143f565b62008f918862008f8a600262000af7565b90620009f1565b6200081e565b6200902f620090116200900b62008fc562008fbf8562008fb8600362000b1c565b90620009f1565b6200080f565b6200900462008fea62008fe48762008fdd5f620007a4565b90620009f1565b6200080f565b62008ffd67d667c2055387940f62008930565b9062003081565b9062001fa2565b6200143f565b620090298862009022600362000b1c565b90620009f1565b6200081e565b620090c7620090a9620090a36200905d620090578562009050600462000b41565b90620009f1565b6200080f565b6200909c620090826200907c87620090755f620007a4565b90620009f1565b6200080f565b62009095670ba63a63e94b5ff062008955565b9062003081565b9062001fa2565b6200143f565b620090c188620090ba600462000b41565b90620009f1565b6200081e565b6200915f620091416200913b620090f5620090ef85620090e8600562000b66565b90620009f1565b6200080f565b620091346200911a62009114876200910d5f620007a4565b90620009f1565b6200080f565b6200912d6799460cc41b8f079f6200897a565b9062003081565b9062001fa2565b6200143f565b620091598862009152600562000b66565b90620009f1565b6200081e565b620091f7620091d9620091d36200918d620091878562009180600662000b8b565b90620009f1565b6200080f565b620091cc620091b2620091ac87620091a55f620007a4565b90620009f1565b6200080f565b620091c5677ff02375ed524bb36200899f565b9062003081565b9062001fa2565b6200143f565b620091f188620091ea600662000b8b565b90620009f1565b6200081e565b6200928f620092716200926b620092256200921f8562009218600762000bb0565b90620009f1565b6200080f565b620092646200924a62009244876200923d5f620007a4565b90620009f1565b6200080f565b6200925d67ea0870b47a8caf0e620089c4565b9062003081565b9062001fa2565b6200143f565b620092898862009282600762000bb0565b90620009f1565b6200081e565b620093276200930962009303620092bd620092b785620092b060086200145b565b90620009f1565b6200080f565b620092fc620092e2620092dc87620092d55f620007a4565b90620009f1565b6200080f565b620092f567abcad82633b7bc9d620089e9565b9062003081565b9062001fa2565b6200143f565b62009321886200931a60086200145b565b90620009f1565b6200081e565b620093bf620093a16200939b620093556200934f8562009348600962002102565b90620009f1565b6200080f565b620093946200937a62009374876200936d5f620007a4565b90620009f1565b6200080f565b6200938d673b8d13526105224162008a0e565b9062003081565b9062001fa2565b6200143f565b620093b988620093b2600962002102565b90620009f1565b6200081e565b620094576200943962009433620093ed620093e785620093e0600a6200214c565b90620009f1565b6200080f565b6200942c620094126200940c87620094055f620007a4565b90620009f1565b6200080f565b6200942567fb4515f5e5b0d53962008a33565b9062003081565b9062001fa2565b6200143f565b62009451886200944a600a6200214c565b90620009f1565b6200081e565b620094b86200949e62009498620094856200947f8562009478600b62002196565b90620009f1565b6200080f565b93620094915f620007a4565b90620009f1565b6200080f565b620094b1673ee8011c2b37f77c62008a58565b9062003081565b9062001fa2565b6200143f565b620094dd84620094d6600b62002196565b90620009f1565b6200081e565b5b565b80620094fe620094f7600162000ad2565b916200012b565b145f1462009efd575062009ed2816200987a6200985d6200985762009ef6966200985062009809620097c36200977c62009735620096ee8a620096e7620096cd62009ed89f620096b36200966c62009625620095de620095976200957d62009577620096c79762009570600162000ad2565b90620009f1565b6200080f565b6200959067673655aae8be5a8b62008421565b9062003081565b620095d7620095bd620095b78b620095b0600262000af7565b90620009f1565b6200080f565b620095d067d510fe714f39fa1062008446565b9062003081565b9062001fa2565b6200961e62009604620095fe8a620095f7600362000b1c565b90620009f1565b6200080f565b62009617672c68a099b51c9e736200846b565b9062003081565b9062001fa2565b620096656200964b62009645896200963e600462000b41565b90620009f1565b6200080f565b6200965e67a667bfa9aa96999d62008490565b9062003081565b9062001fa2565b620096ac620096926200968c8862009685600562000b66565b90620009f1565b6200080f565b620096a5674d67e72f063e2108620084b5565b9062003081565b9062001fa2565b93620096c0600662000b8b565b90620009f1565b6200080f565b620096e067f84dde3e6acda179620084da565b9062003081565b9062001fa2565b6200972e620097146200970e8d62009707600762000bb0565b90620009f1565b6200080f565b620097276740f9cc8c08f80981620084ff565b9062003081565b9062001fa2565b620097756200975b620097558c6200974e60086200145b565b90620009f1565b6200080f565b6200976e675ead03205009714262008524565b9062003081565b9062001fa2565b620097bc620097a26200979c8b62009795600962002102565b90620009f1565b6200080f565b620097b5676591b02092d671bb62008549565b9062003081565b9062001fa2565b62009802620097e9620097e38a620097dc600a6200214c565b90620009f1565b6200080f565b620097fb66e18c71963dd1b76200856e565b9062003081565b9062001fa2565b620098496200982f620098298962009822600b62002196565b90620009f1565b6200080f565b62009842678a21bcd24a14218a62008593565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b62009874886200986d5f620007a4565b90620009f1565b6200081e565b62009912620098f4620098ee620098a8620098a2856200989b600162000ad2565b90620009f1565b6200080f565b620098e7620098cd620098c787620098c05f620007a4565b90620009f1565b6200080f565b620098e0670adef3740e71c726620085b8565b9062003081565b9062001fa2565b6200143f565b6200990c8862009905600162000ad2565b90620009f1565b6200081e565b620099aa6200998c62009986620099406200993a8562009933600262000af7565b90620009f1565b6200080f565b6200997f620099656200995f87620099585f620007a4565b90620009f1565b6200080f565b6200997867a37bf67c6f986559620085dd565b9062003081565b9062001fa2565b6200143f565b620099a4886200999d600262000af7565b90620009f1565b6200081e565b62009a4262009a2462009a1e620099d8620099d285620099cb600362000b1c565b90620009f1565b6200080f565b62009a17620099fd620099f787620099f05f620007a4565b90620009f1565b6200080f565b62009a1067c6b16f7ed4fa1b0062008602565b9062003081565b9062001fa2565b6200143f565b62009a3c8862009a35600362000b1c565b90620009f1565b6200081e565b62009ada62009abc62009ab662009a7062009a6a8562009a63600462000b41565b90620009f1565b6200080f565b62009aaf62009a9562009a8f8762009a885f620007a4565b90620009f1565b6200080f565b62009aa8676a065da88d8bfc3c62008627565b9062003081565b9062001fa2565b6200143f565b62009ad48862009acd600462000b41565b90620009f1565b6200081e565b62009b7262009b5462009b4e62009b0862009b028562009afb600562000b66565b90620009f1565b6200080f565b62009b4762009b2d62009b278762009b205f620007a4565b90620009f1565b6200080f565b62009b40674cabc0916844b46f6200864c565b9062003081565b9062001fa2565b6200143f565b62009b6c8862009b65600562000b66565b90620009f1565b6200081e565b62009c0a62009bec62009be662009ba062009b9a8562009b93600662000b8b565b90620009f1565b6200080f565b62009bdf62009bc562009bbf8762009bb85f620007a4565b90620009f1565b6200080f565b62009bd867407faac0f02e78d162008671565b9062003081565b9062001fa2565b6200143f565b62009c048862009bfd600662000b8b565b90620009f1565b6200081e565b62009ca262009c8462009c7e62009c3862009c328562009c2b600762000bb0565b90620009f1565b6200080f565b62009c7762009c5d62009c578762009c505f620007a4565b90620009f1565b6200080f565b62009c706707a786d9cf0852cf62008696565b9062003081565b9062001fa2565b6200143f565b62009c9c8862009c95600762000bb0565b90620009f1565b6200081e565b62009d3a62009d1c62009d1662009cd062009cca8562009cc360086200145b565b90620009f1565b6200080f565b62009d0f62009cf562009cef8762009ce85f620007a4565b90620009f1565b6200080f565b62009d086742433fb6949a629a620086bb565b9062003081565b9062001fa2565b6200143f565b62009d348862009d2d60086200145b565b90620009f1565b6200081e565b62009dd262009db462009dae62009d6862009d628562009d5b600962002102565b90620009f1565b6200080f565b62009da762009d8d62009d878762009d805f620007a4565b90620009f1565b6200080f565b62009da067891682a147ce43b0620086e0565b9062003081565b9062001fa2565b6200143f565b62009dcc8862009dc5600962002102565b90620009f1565b6200081e565b62009e6a62009e4c62009e4662009e0062009dfa8562009df3600a6200214c565b90620009f1565b6200080f565b62009e3f62009e2562009e1f8762009e185f620007a4565b90620009f1565b6200080f565b62009e386726cfd58e7b003b5562008705565b9062003081565b9062001fa2565b6200143f565b62009e648862009e5d600a6200214c565b90620009f1565b6200081e565b62009ecb62009eb162009eab62009e9862009e928562009e8b600b62002196565b90620009f1565b6200080f565b9362009ea45f620007a4565b90620009f1565b6200080f565b62009ec4672bbf0ed7b657acb36200872a565b9062003081565b9062001fa2565b6200143f565b62009ef08462009ee9600b62002196565b90620009f1565b6200081e565b5b620094e4565b8062009f1562009f0e600262000af7565b916200012b565b145f146200a91557506200a8ea816200a2926200a2756200a26f6200a90e966200a2686200a2216200a1da6200a1936200a14c6200a1058a6200a0fe6200a0e46200a8f09f6200a0ca6200a0836200a03c62009ff562009fae62009f9462009f8e6200a0de9762009f87600162000ad2565b90620009f1565b6200080f565b62009fa767202800f4addbdc87620080f3565b9062003081565b62009fee62009fd462009fce8b62009fc7600262000af7565b90620009f1565b6200080f565b62009fe767e4b5bdb1cc3504ff62008118565b9062003081565b9062001fa2565b6200a0356200a01b6200a0158a6200a00e600362000b1c565b90620009f1565b6200080f565b6200a02e67be32b32a825596e76200813d565b9062003081565b9062001fa2565b6200a07c6200a0626200a05c896200a055600462000b41565b90620009f1565b6200080f565b6200a075678e0f68c5dc223b9a62008162565b9062003081565b9062001fa2565b6200a0c36200a0a96200a0a3886200a09c600562000b66565b90620009f1565b6200080f565b6200a0bc6758022d9e1c256ce362008187565b9062003081565b9062001fa2565b936200a0d7600662000b8b565b90620009f1565b6200080f565b6200a0f767584d29227aa073ac620081ac565b9062003081565b9062001fa2565b6200a1456200a12b6200a1258d6200a11e600762000bb0565b90620009f1565b6200080f565b6200a13e678b9352ad04bef9e7620081d1565b9062003081565b9062001fa2565b6200a18c6200a1726200a16c8c6200a16560086200145b565b90620009f1565b6200080f565b6200a18567aead42a3f445ecbf620081f6565b9062003081565b9062001fa2565b6200a1d36200a1b96200a1b38b6200a1ac600962002102565b90620009f1565b6200080f565b6200a1cc673c667a1d833a3cca6200821b565b9062003081565b9062001fa2565b6200a21a6200a2006200a1fa8a6200a1f3600a6200214c565b90620009f1565b6200080f565b6200a21367da6f61838efa1ffe62008240565b9062003081565b9062001fa2565b6200a2616200a2476200a241896200a23a600b62002196565b90620009f1565b6200080f565b6200a25a67e8f749470bd7c44662008265565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6200a28c886200a2855f620007a4565b90620009f1565b6200081e565b6200a32a6200a30c6200a3066200a2c06200a2ba856200a2b3600162000ad2565b90620009f1565b6200080f565b6200a2ff6200a2e56200a2df876200a2d85f620007a4565b90620009f1565b6200080f565b6200a2f867481ac7746b159c676200828a565b9062003081565b9062001fa2565b6200143f565b6200a324886200a31d600162000ad2565b90620009f1565b6200081e565b6200a3c26200a3a46200a39e6200a3586200a352856200a34b600262000af7565b90620009f1565b6200080f565b6200a3976200a37d6200a377876200a3705f620007a4565b90620009f1565b6200080f565b6200a39067e367de32f108e278620082af565b9062003081565b9062001fa2565b6200143f565b6200a3bc886200a3b5600262000af7565b90620009f1565b6200081e565b6200a45a6200a43c6200a4366200a3f06200a3ea856200a3e3600362000b1c565b90620009f1565b6200080f565b6200a42f6200a4156200a40f876200a4085f620007a4565b90620009f1565b6200080f565b6200a4286773f260087ad28bec620082d4565b9062003081565b9062001fa2565b6200143f565b6200a454886200a44d600362000b1c565b90620009f1565b6200081e565b6200a4f26200a4d46200a4ce6200a4886200a482856200a47b600462000b41565b90620009f1565b6200080f565b6200a4c76200a4ad6200a4a7876200a4a05f620007a4565b90620009f1565b6200080f565b6200a4c0675cfc82216bc1bdca620082f9565b9062003081565b9062001fa2565b6200143f565b6200a4ec886200a4e5600462000b41565b90620009f1565b6200081e565b6200a58a6200a56c6200a5666200a5206200a51a856200a513600562000b66565b90620009f1565b6200080f565b6200a55f6200a5456200a53f876200a5385f620007a4565b90620009f1565b6200080f565b6200a55867caccc870a2663a0e6200831e565b9062003081565b9062001fa2565b6200143f565b6200a584886200a57d600562000b66565b90620009f1565b6200081e565b6200a6226200a6046200a5fe6200a5b86200a5b2856200a5ab600662000b8b565b90620009f1565b6200080f565b6200a5f76200a5dd6200a5d7876200a5d05f620007a4565b90620009f1565b6200080f565b6200a5f067db69cd7b4298c45d62008343565b9062003081565b9062001fa2565b6200143f565b6200a61c886200a615600662000b8b565b90620009f1565b6200081e565b6200a6ba6200a69c6200a6966200a6506200a64a856200a643600762000bb0565b90620009f1565b6200080f565b6200a68f6200a6756200a66f876200a6685f620007a4565b90620009f1565b6200080f565b6200a688677bc9e0c57243e62d62008368565b9062003081565b9062001fa2565b6200143f565b6200a6b4886200a6ad600762000bb0565b90620009f1565b6200081e565b6200a7526200a7346200a72e6200a6e86200a6e2856200a6db60086200145b565b90620009f1565b6200080f565b6200a7276200a70d6200a707876200a7005f620007a4565b90620009f1565b6200080f565b6200a720673cc51c5d368693ae6200838d565b9062003081565b9062001fa2565b6200143f565b6200a74c886200a74560086200145b565b90620009f1565b6200081e565b6200a7ea6200a7cc6200a7c66200a7806200a77a856200a773600962002102565b90620009f1565b6200080f565b6200a7bf6200a7a56200a79f876200a7985f620007a4565b90620009f1565b6200080f565b6200a7b867366b4e8cc068895b620083b2565b9062003081565b9062001fa2565b6200143f565b6200a7e4886200a7dd600962002102565b90620009f1565b6200081e565b6200a8826200a8646200a85e6200a8186200a812856200a80b600a6200214c565b90620009f1565b6200080f565b6200a8576200a83d6200a837876200a8305f620007a4565b90620009f1565b6200080f565b6200a850672bd18715cdabbca4620083d7565b9062003081565b9062001fa2565b6200143f565b6200a87c886200a875600a6200214c565b90620009f1565b6200081e565b6200a8e36200a8c96200a8c36200a8b06200a8aa856200a8a3600b62002196565b90620009f1565b6200080f565b936200a8bc5f620007a4565b90620009f1565b6200080f565b6200a8dc67a752061c4f33b8cf620083fc565b9062003081565b9062001fa2565b6200143f565b6200a908846200a901600b62002196565b90620009f1565b6200081e565b5b62009ef7565b806200a92d6200a926600362000b1c565b916200012b565b145f146200b32d57506200b302816200acaa6200ac8d6200ac876200b326966200ac806200ac396200abf26200abab6200ab646200ab1d8a6200ab166200aafc6200b3089f6200aae26200aa9b6200aa546200aa0d6200a9c66200a9ac6200a9a66200aaf6976200a99f600162000ad2565b90620009f1565b6200080f565b6200a9bf67c5b85bab9e5b386962007dc5565b9062003081565b6200aa066200a9ec6200a9e68b6200a9df600262000af7565b90620009f1565b6200080f565b6200a9ff6745245258aec51cf762007dea565b9062003081565b9062001fa2565b6200aa4d6200aa336200aa2d8a6200aa26600362000b1c565b90620009f1565b6200080f565b6200aa466716e6b8e68b93183062007e0f565b9062003081565b9062001fa2565b6200aa946200aa7a6200aa74896200aa6d600462000b41565b90620009f1565b6200080f565b6200aa8d67e2ae0f051418112c62007e34565b9062003081565b9062001fa2565b6200aadb6200aac16200aabb886200aab4600562000b66565b90620009f1565b6200080f565b6200aad4670470e26a0093a65b62007e59565b9062003081565b9062001fa2565b936200aaef600662000b8b565b90620009f1565b6200080f565b6200ab0f676bef71973a8146ed62007e7e565b9062003081565b9062001fa2565b6200ab5d6200ab436200ab3d8d6200ab36600762000bb0565b90620009f1565b6200080f565b6200ab5667119265be51812daf62007ea3565b9062003081565b9062001fa2565b6200aba46200ab8a6200ab848c6200ab7d60086200145b565b90620009f1565b6200080f565b6200ab9d67b0be7356254bea2e62007ec8565b9062003081565b9062001fa2565b6200abeb6200abd16200abcb8b6200abc4600962002102565b90620009f1565b6200080f565b6200abe4678584defff7589bd762007eed565b9062003081565b9062001fa2565b6200ac326200ac186200ac128a6200ac0b600a6200214c565b90620009f1565b6200080f565b6200ac2b673c5fe4aeb1fb52ba62007f12565b9062003081565b9062001fa2565b6200ac796200ac5f6200ac59896200ac52600b62002196565b90620009f1565b6200080f565b6200ac72679e7cd88acf543a5e62007f37565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6200aca4886200ac9d5f620007a4565b90620009f1565b6200081e565b6200ad426200ad246200ad1e6200acd86200acd2856200accb600162000ad2565b90620009f1565b6200080f565b6200ad176200acfd6200acf7876200acf05f620007a4565b90620009f1565b6200080f565b6200ad1067b22d2432b72d509862007f5c565b9062003081565b9062001fa2565b6200143f565b6200ad3c886200ad35600162000ad2565b90620009f1565b6200081e565b6200adda6200adbc6200adb66200ad706200ad6a856200ad63600262000af7565b90620009f1565b6200080f565b6200adaf6200ad956200ad8f876200ad885f620007a4565b90620009f1565b6200080f565b6200ada8679e18a487f44d2fe462007f81565b9062003081565b9062001fa2565b6200143f565b6200add4886200adcd600262000af7565b90620009f1565b6200081e565b6200ae726200ae546200ae4e6200ae086200ae02856200adfb600362000b1c565b90620009f1565b6200080f565b6200ae476200ae2d6200ae27876200ae205f620007a4565b90620009f1565b6200080f565b6200ae40674b39e14ce22abd3c62007fa6565b9062003081565b9062001fa2565b6200143f565b6200ae6c886200ae65600362000b1c565b90620009f1565b6200081e565b6200af0a6200aeec6200aee66200aea06200ae9a856200ae93600462000b41565b90620009f1565b6200080f565b6200aedf6200aec56200aebf876200aeb85f620007a4565b90620009f1565b6200080f565b6200aed8679e77fde2eb315e0d62007fcb565b9062003081565b9062001fa2565b6200143f565b6200af04886200aefd600462000b41565b90620009f1565b6200081e565b6200afa26200af846200af7e6200af386200af32856200af2b600562000b66565b90620009f1565b6200080f565b6200af776200af5d6200af57876200af505f620007a4565b90620009f1565b6200080f565b6200af7067ca5e0385fe67014d62007ff0565b9062003081565b9062001fa2565b6200143f565b6200af9c886200af95600562000b66565b90620009f1565b6200081e565b6200b03a6200b01c6200b0166200afd06200afca856200afc3600662000b8b565b90620009f1565b6200080f565b6200b00f6200aff56200afef876200afe85f620007a4565b90620009f1565b6200080f565b6200b008670c2cb99bf1b6bddb62008015565b9062003081565b9062001fa2565b6200143f565b6200b034886200b02d600662000b8b565b90620009f1565b6200081e565b6200b0d26200b0b46200b0ae6200b0686200b062856200b05b600762000bb0565b90620009f1565b6200080f565b6200b0a76200b08d6200b087876200b0805f620007a4565b90620009f1565b6200080f565b6200b0a06799ec1cd2a4460bfe6200803a565b9062003081565b9062001fa2565b6200143f565b6200b0cc886200b0c5600762000bb0565b90620009f1565b6200081e565b6200b16a6200b14c6200b1466200b1006200b0fa856200b0f360086200145b565b90620009f1565b6200080f565b6200b13f6200b1256200b11f876200b1185f620007a4565b90620009f1565b6200080f565b6200b138678577a815a2ff843f6200805f565b9062003081565b9062001fa2565b6200143f565b6200b164886200b15d60086200145b565b90620009f1565b6200081e565b6200b2026200b1e46200b1de6200b1986200b192856200b18b600962002102565b90620009f1565b6200080f565b6200b1d76200b1bd6200b1b7876200b1b05f620007a4565b90620009f1565b6200080f565b6200b1d0677d80a6b4fd6518a562008084565b9062003081565b9062001fa2565b6200143f565b6200b1fc886200b1f5600962002102565b90620009f1565b6200081e565b6200b29a6200b27c6200b2766200b2306200b22a856200b223600a6200214c565b90620009f1565b6200080f565b6200b26f6200b2556200b24f876200b2485f620007a4565b90620009f1565b6200080f565b6200b26867eb6c67123eab62cb620080a9565b9062003081565b9062001fa2565b6200143f565b6200b294886200b28d600a6200214c565b90620009f1565b6200081e565b6200b2fb6200b2e16200b2db6200b2c86200b2c2856200b2bb600b62002196565b90620009f1565b6200080f565b936200b2d45f620007a4565b90620009f1565b6200080f565b6200b2f4678f7851650eca21a5620080ce565b9062003081565b9062001fa2565b6200143f565b6200b320846200b319600b62002196565b90620009f1565b6200081e565b5b6200a90f565b806200b3456200b33e600462000b41565b916200012b565b145f146200bd4557506200bd1a816200b6c26200b6a56200b69f6200bd3e966200b6986200b6516200b60a6200b5c36200b57c6200b5358a6200b52e6200b5146200bd209f6200b4fa6200b4b36200b46c6200b4256200b3de6200b3c46200b3be6200b50e976200b3b7600162000ad2565b90620009f1565b6200080f565b6200b3d767179be4bba87f0a8c62007a97565b9062003081565b6200b41e6200b4046200b3fe8b6200b3f7600262000af7565b90620009f1565b6200080f565b6200b41767acf63d95d888735562007abc565b9062003081565b9062001fa2565b6200b4656200b44b6200b4458a6200b43e600362000b1c565b90620009f1565b6200080f565b6200b45e676696670196b0074f62007ae1565b9062003081565b9062001fa2565b6200b4ac6200b4926200b48c896200b485600462000b41565b90620009f1565b6200080f565b6200b4a567d99ddf1fe75085f962007b06565b9062003081565b9062001fa2565b6200b4f36200b4d96200b4d3886200b4cc600562000b66565b90620009f1565b6200080f565b6200b4ec67c2597881fef0283b62007b2b565b9062003081565b9062001fa2565b936200b507600662000b8b565b90620009f1565b6200080f565b6200b52767cf48395ee6c54f1462007b50565b9062003081565b9062001fa2565b6200b5756200b55b6200b5558d6200b54e600762000bb0565b90620009f1565b6200080f565b6200b56e6715226a8e4cd8d3b662007b75565b9062003081565b9062001fa2565b6200b5bc6200b5a26200b59c8c6200b59560086200145b565b90620009f1565b6200080f565b6200b5b567c053297389af5d3b62007b9a565b9062003081565b9062001fa2565b6200b6036200b5e96200b5e38b6200b5dc600962002102565b90620009f1565b6200080f565b6200b5fc672c08893f0d1580e262007bbf565b9062003081565b9062001fa2565b6200b64a6200b6306200b62a8a6200b623600a6200214c565b90620009f1565b6200080f565b6200b643670ed3cbcff6fcc5ba62007be4565b9062003081565b9062001fa2565b6200b6916200b6776200b671896200b66a600b62002196565b90620009f1565b6200080f565b6200b68a67c82f510ecf81f6d062007c09565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6200b6bc886200b6b55f620007a4565b90620009f1565b6200081e565b6200b75a6200b73c6200b7366200b6f06200b6ea856200b6e3600162000ad2565b90620009f1565b6200080f565b6200b72f6200b7156200b70f876200b7085f620007a4565b90620009f1565b6200080f565b6200b7286711ba9a1b81718c2a62007c2e565b9062003081565b9062001fa2565b6200143f565b6200b754886200b74d600162000ad2565b90620009f1565b6200081e565b6200b7f26200b7d46200b7ce6200b7886200b782856200b77b600262000af7565b90620009f1565b6200080f565b6200b7c76200b7ad6200b7a7876200b7a05f620007a4565b90620009f1565b6200080f565b6200b7c0679f7d798a3323410c62007c53565b9062003081565b9062001fa2565b6200143f565b6200b7ec886200b7e5600262000af7565b90620009f1565b6200081e565b6200b88a6200b86c6200b8666200b8206200b81a856200b813600362000b1c565b90620009f1565b6200080f565b6200b85f6200b8456200b83f876200b8385f620007a4565b90620009f1565b6200080f565b6200b85867a821855c8c1cf5e562007c78565b9062003081565b9062001fa2565b6200143f565b6200b884886200b87d600362000b1c565b90620009f1565b6200081e565b6200b9226200b9046200b8fe6200b8b86200b8b2856200b8ab600462000b41565b90620009f1565b6200080f565b6200b8f76200b8dd6200b8d7876200b8d05f620007a4565b90620009f1565b6200080f565b6200b8f067535e8d6fac0031b262007c9d565b9062003081565b9062001fa2565b6200143f565b6200b91c886200b915600462000b41565b90620009f1565b6200081e565b6200b9ba6200b99c6200b9966200b9506200b94a856200b943600562000b66565b90620009f1565b6200080f565b6200b98f6200b9756200b96f876200b9685f620007a4565b90620009f1565b6200080f565b6200b98867404e7c751b63432062007cc2565b9062003081565b9062001fa2565b6200143f565b6200b9b4886200b9ad600562000b66565b90620009f1565b6200081e565b6200ba526200ba346200ba2e6200b9e86200b9e2856200b9db600662000b8b565b90620009f1565b6200080f565b6200ba276200ba0d6200ba07876200ba005f620007a4565b90620009f1565b6200080f565b6200ba2067a729353f6e55d35462007ce7565b9062003081565b9062001fa2565b6200143f565b6200ba4c886200ba45600662000b8b565b90620009f1565b6200081e565b6200baea6200bacc6200bac66200ba806200ba7a856200ba73600762000bb0565b90620009f1565b6200080f565b6200babf6200baa56200ba9f876200ba985f620007a4565b90620009f1565b6200080f565b6200bab8674db97d92e58bb83162007d0c565b9062003081565b9062001fa2565b6200143f565b6200bae4886200badd600762000bb0565b90620009f1565b6200081e565b6200bb826200bb646200bb5e6200bb186200bb12856200bb0b60086200145b565b90620009f1565b6200080f565b6200bb576200bb3d6200bb37876200bb305f620007a4565b90620009f1565b6200080f565b6200bb5067b53926c27897bf7d62007d31565b9062003081565b9062001fa2565b6200143f565b6200bb7c886200bb7560086200145b565b90620009f1565b6200081e565b6200bc1a6200bbfc6200bbf66200bbb06200bbaa856200bba3600962002102565b90620009f1565b6200080f565b6200bbef6200bbd56200bbcf876200bbc85f620007a4565b90620009f1565b6200080f565b6200bbe867965040d52fe115c562007d56565b9062003081565b9062001fa2565b6200143f565b6200bc14886200bc0d600962002102565b90620009f1565b6200081e565b6200bcb26200bc946200bc8e6200bc486200bc42856200bc3b600a6200214c565b90620009f1565b6200080f565b6200bc876200bc6d6200bc67876200bc605f620007a4565b90620009f1565b6200080f565b6200bc80679565fa41ebd31fd762007d7b565b9062003081565b9062001fa2565b6200143f565b6200bcac886200bca5600a6200214c565b90620009f1565b6200081e565b6200bd136200bcf96200bcf36200bce06200bcda856200bcd3600b62002196565b90620009f1565b6200080f565b936200bcec5f620007a4565b90620009f1565b6200080f565b6200bd0c67aae4438c877ea8f462007da0565b9062003081565b9062001fa2565b6200143f565b6200bd38846200bd31600b62002196565b90620009f1565b6200081e565b5b6200b327565b806200bd5d6200bd56600562000b66565b916200012b565b145f146200c75d57506200c732816200c0da6200c0bd6200c0b76200c756966200c0b06200c0696200c0226200bfdb6200bf946200bf4d8a6200bf466200bf2c6200c7389f6200bf126200becb6200be846200be3d6200bdf66200bddc6200bdd66200bf26976200bdcf600162000ad2565b90620009f1565b6200080f565b6200bdef6794b06183acb715cc62007769565b9062003081565b6200be366200be1c6200be168b6200be0f600262000af7565b90620009f1565b6200080f565b6200be2f67500392ed0d4311376200778e565b9062003081565b9062001fa2565b6200be7d6200be636200be5d8a6200be56600362000b1c565b90620009f1565b6200080f565b6200be7667861cc95ad5c86323620077b3565b9062003081565b9062001fa2565b6200bec46200beaa6200bea4896200be9d600462000b41565b90620009f1565b6200080f565b6200bebd6705830a443f86c4ac620077d8565b9062003081565b9062001fa2565b6200bf0b6200bef16200beeb886200bee4600562000b66565b90620009f1565b6200080f565b6200bf04673b68225874a20a7c620077fd565b9062003081565b9062001fa2565b936200bf1f600662000b8b565b90620009f1565b6200080f565b6200bf3f6710b3309838e236fb62007822565b9062003081565b9062001fa2565b6200bf8d6200bf736200bf6d8d6200bf66600762000bb0565b90620009f1565b6200080f565b6200bf86679b77fc8bcd559e2c62007847565b9062003081565b9062001fa2565b6200bfd46200bfba6200bfb48c6200bfad60086200145b565b90620009f1565b6200080f565b6200bfcd67bdecf5e0cb9cb2136200786c565b9062003081565b9062001fa2565b6200c01b6200c0016200bffb8b6200bff4600962002102565b90620009f1565b6200080f565b6200c0146730276f1221ace5fa62007891565b9062003081565b9062001fa2565b6200c0626200c0486200c0428a6200c03b600a6200214c565b90620009f1565b6200080f565b6200c05b677935dd342764a144620078b6565b9062003081565b9062001fa2565b6200c0a96200c08f6200c089896200c082600b62002196565b90620009f1565b6200080f565b6200c0a267eac6db520bb03708620078db565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6200c0d4886200c0cd5f620007a4565b90620009f1565b6200081e565b6200c1726200c1546200c14e6200c1086200c102856200c0fb600162000ad2565b90620009f1565b6200080f565b6200c1476200c12d6200c127876200c1205f620007a4565b90620009f1565b6200080f565b6200c1406737f4e36af6073c6e62007900565b9062003081565b9062001fa2565b6200143f565b6200c16c886200c165600162000ad2565b90620009f1565b6200081e565b6200c20a6200c1ec6200c1e66200c1a06200c19a856200c193600262000af7565b90620009f1565b6200080f565b6200c1df6200c1c56200c1bf876200c1b85f620007a4565b90620009f1565b6200080f565b6200c1d8674edc0918210800e962007925565b9062003081565b9062001fa2565b6200143f565b6200c204886200c1fd600262000af7565b90620009f1565b6200081e565b6200c2a26200c2846200c27e6200c2386200c232856200c22b600362000b1c565b90620009f1565b6200080f565b6200c2776200c25d6200c257876200c2505f620007a4565b90620009f1565b6200080f565b6200c27067c44998e99eae41886200794a565b9062003081565b9062001fa2565b6200143f565b6200c29c886200c295600362000b1c565b90620009f1565b6200081e565b6200c33a6200c31c6200c3166200c2d06200c2ca856200c2c3600462000b41565b90620009f1565b6200080f565b6200c30f6200c2f56200c2ef876200c2e85f620007a4565b90620009f1565b6200080f565b6200c308679f4310d05d0683386200796f565b9062003081565b9062001fa2565b6200143f565b6200c334886200c32d600462000b41565b90620009f1565b6200081e565b6200c3d26200c3b46200c3ae6200c3686200c362856200c35b600562000b66565b90620009f1565b6200080f565b6200c3a76200c38d6200c387876200c3805f620007a4565b90620009f1565b6200080f565b6200c3a0679ec7fe4350680f2962007994565b9062003081565b9062001fa2565b6200143f565b6200c3cc886200c3c5600562000b66565b90620009f1565b6200081e565b6200c46a6200c44c6200c4466200c4006200c3fa856200c3f3600662000b8b565b90620009f1565b6200080f565b6200c43f6200c4256200c41f876200c4185f620007a4565b90620009f1565b6200080f565b6200c43867c5b2c1fdc0b50874620079b9565b9062003081565b9062001fa2565b6200143f565b6200c464886200c45d600662000b8b565b90620009f1565b6200081e565b6200c5026200c4e46200c4de6200c4986200c492856200c48b600762000bb0565b90620009f1565b6200080f565b6200c4d76200c4bd6200c4b7876200c4b05f620007a4565b90620009f1565b6200080f565b6200c4d067a01920c5ef8b2ebe620079de565b9062003081565b9062001fa2565b6200143f565b6200c4fc886200c4f5600762000bb0565b90620009f1565b6200081e565b6200c59a6200c57c6200c5766200c5306200c52a856200c52360086200145b565b90620009f1565b6200080f565b6200c56f6200c5556200c54f876200c5485f620007a4565b90620009f1565b6200080f565b6200c5686759fa6f8bd91d58ba62007a03565b9062003081565b9062001fa2565b6200143f565b6200c594886200c58d60086200145b565b90620009f1565b6200081e565b6200c6326200c6146200c60e6200c5c86200c5c2856200c5bb600962002102565b90620009f1565b6200080f565b6200c6076200c5ed6200c5e7876200c5e05f620007a4565b90620009f1565b6200080f565b6200c600678bfc9eb89b515a8262007a28565b9062003081565b9062001fa2565b6200143f565b6200c62c886200c625600962002102565b90620009f1565b6200081e565b6200c6ca6200c6ac6200c6a66200c6606200c65a856200c653600a6200214c565b90620009f1565b6200080f565b6200c69f6200c6856200c67f876200c6785f620007a4565b90620009f1565b6200080f565b6200c69867be86a7a2555ae77562007a4d565b9062003081565b9062001fa2565b6200143f565b6200c6c4886200c6bd600a6200214c565b90620009f1565b6200081e565b6200c72b6200c7116200c70b6200c6f86200c6f2856200c6eb600b62002196565b90620009f1565b6200080f565b936200c7045f620007a4565b90620009f1565b6200080f565b6200c72467cbb8bbaa3810babf62007a72565b9062003081565b9062001fa2565b6200143f565b6200c750846200c749600b62002196565b90620009f1565b6200081e565b5b6200bd3f565b806200c7756200c76e600662000b8b565b916200012b565b145f146200d17557506200d14a816200caf26200cad56200cacf6200d16e966200cac86200ca816200ca3a6200c9f36200c9ac6200c9658a6200c95e6200c9446200d1509f6200c92a6200c8e36200c89c6200c8556200c80e6200c7f46200c7ee6200c93e976200c7e7600162000ad2565b90620009f1565b6200080f565b6200c807677186a80551025f8f6200743b565b9062003081565b6200c84e6200c8346200c82e8b6200c827600262000af7565b90620009f1565b6200080f565b6200c84767622247557e9b537162007460565b9062003081565b9062001fa2565b6200c8956200c87b6200c8758a6200c86e600362000b1c565b90620009f1565b6200080f565b6200c88e67c4cbe326d1ad974262007485565b9062003081565b9062001fa2565b6200c8dc6200c8c26200c8bc896200c8b5600462000b41565b90620009f1565b6200080f565b6200c8d56755f1523ac6a23ea2620074aa565b9062003081565b9062001fa2565b6200c9236200c9096200c903886200c8fc600562000b66565b90620009f1565b6200080f565b6200c91c67a13dfe77a3d52f53620074cf565b9062003081565b9062001fa2565b936200c937600662000b8b565b90620009f1565b6200080f565b6200c95767e30750b6301c0452620074f4565b9062003081565b9062001fa2565b6200c9a56200c98b6200c9858d6200c97e600762000bb0565b90620009f1565b6200080f565b6200c99e6708bd488070a3a32b62007519565b9062003081565b9062001fa2565b6200c9ec6200c9d26200c9cc8c6200c9c560086200145b565b90620009f1565b6200080f565b6200c9e567cd800caef5b72ae36200753e565b9062003081565b9062001fa2565b6200ca336200ca196200ca138b6200ca0c600962002102565b90620009f1565b6200080f565b6200ca2c6783329c90f04233ce62007563565b9062003081565b9062001fa2565b6200ca7a6200ca606200ca5a8a6200ca53600a6200214c565b90620009f1565b6200080f565b6200ca7367b5b99e6664a0a3ee62007588565b9062003081565b9062001fa2565b6200cac16200caa76200caa1896200ca9a600b62002196565b90620009f1565b6200080f565b6200caba676b0731849e200a7f620075ad565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6200caec886200cae55f620007a4565b90620009f1565b6200081e565b6200cb8a6200cb6c6200cb666200cb206200cb1a856200cb13600162000ad2565b90620009f1565b6200080f565b6200cb5f6200cb456200cb3f876200cb385f620007a4565b90620009f1565b6200080f565b6200cb5867577f9a9e7ee3f9c2620075d2565b9062003081565b9062001fa2565b6200143f565b6200cb84886200cb7d600162000ad2565b90620009f1565b6200081e565b6200cc226200cc046200cbfe6200cbb86200cbb2856200cbab600262000af7565b90620009f1565b6200080f565b6200cbf76200cbdd6200cbd7876200cbd05f620007a4565b90620009f1565b6200080f565b6200cbf06788c522b949ace7b1620075f7565b9062003081565b9062001fa2565b6200143f565b6200cc1c886200cc15600262000af7565b90620009f1565b6200081e565b6200ccba6200cc9c6200cc966200cc506200cc4a856200cc43600362000b1c565b90620009f1565b6200080f565b6200cc8f6200cc756200cc6f876200cc685f620007a4565b90620009f1565b6200080f565b6200cc886782f07007c8b721066200761c565b9062003081565b9062001fa2565b6200143f565b6200ccb4886200ccad600362000b1c565b90620009f1565b6200081e565b6200cd526200cd346200cd2e6200cce86200cce2856200ccdb600462000b41565b90620009f1565b6200080f565b6200cd276200cd0d6200cd07876200cd005f620007a4565b90620009f1565b6200080f565b6200cd20678283d37c6675b50e62007641565b9062003081565b9062001fa2565b6200143f565b6200cd4c886200cd45600462000b41565b90620009f1565b6200081e565b6200cdea6200cdcc6200cdc66200cd806200cd7a856200cd73600562000b66565b90620009f1565b6200080f565b6200cdbf6200cda56200cd9f876200cd985f620007a4565b90620009f1565b6200080f565b6200cdb86798b074d9bbac112362007666565b9062003081565b9062001fa2565b6200143f565b6200cde4886200cddd600562000b66565b90620009f1565b6200081e565b6200ce826200ce646200ce5e6200ce186200ce12856200ce0b600662000b8b565b90620009f1565b6200080f565b6200ce576200ce3d6200ce37876200ce305f620007a4565b90620009f1565b6200080f565b6200ce506775c56fb7758317c16200768b565b9062003081565b9062001fa2565b6200143f565b6200ce7c886200ce75600662000b8b565b90620009f1565b6200081e565b6200cf1a6200cefc6200cef66200ceb06200ceaa856200cea3600762000bb0565b90620009f1565b6200080f565b6200ceef6200ced56200cecf876200cec85f620007a4565b90620009f1565b6200080f565b6200cee867fed24e206052bc72620076b0565b9062003081565b9062001fa2565b6200143f565b6200cf14886200cf0d600762000bb0565b90620009f1565b6200081e565b6200cfb26200cf946200cf8e6200cf486200cf42856200cf3b60086200145b565b90620009f1565b6200080f565b6200cf876200cf6d6200cf67876200cf605f620007a4565b90620009f1565b6200080f565b6200cf806726d7c3d1bc07dae5620076d5565b9062003081565b9062001fa2565b6200143f565b6200cfac886200cfa560086200145b565b90620009f1565b6200081e565b6200d04a6200d02c6200d0266200cfe06200cfda856200cfd3600962002102565b90620009f1565b6200080f565b6200d01f6200d0056200cfff876200cff85f620007a4565b90620009f1565b6200080f565b6200d01867f88c5e441e28dbb4620076fa565b9062003081565b9062001fa2565b6200143f565b6200d044886200d03d600962002102565b90620009f1565b6200081e565b6200d0e26200d0c46200d0be6200d0786200d072856200d06b600a6200214c565b90620009f1565b6200080f565b6200d0b76200d09d6200d097876200d0905f620007a4565b90620009f1565b6200080f565b6200d0b0674fe27f9f966152706200771f565b9062003081565b9062001fa2565b6200143f565b6200d0dc886200d0d5600a6200214c565b90620009f1565b6200081e565b6200d1436200d1296200d1236200d1106200d10a856200d103600b62002196565b90620009f1565b6200080f565b936200d11c5f620007a4565b90620009f1565b6200080f565b6200d13c67514d4ba49c2b14fe62007744565b9062003081565b9062001fa2565b6200143f565b6200d168846200d161600b62002196565b90620009f1565b6200081e565b5b6200c757565b806200d18d6200d186600762000bb0565b916200012b565b145f146200db8d57506200db62816200d50a6200d4ed6200d4e76200db86966200d4e06200d4996200d4526200d40b6200d3c46200d37d8a6200d3766200d35c6200db689f6200d3426200d2fb6200d2b46200d26d6200d2266200d20c6200d2066200d356976200d1ff600162000ad2565b90620009f1565b6200080f565b6200d21f67ec3fabc192b017996200710d565b9062003081565b6200d2666200d24c6200d2468b6200d23f600262000af7565b90620009f1565b6200080f565b6200d25f67382b38cee8ee537562007132565b9062003081565b9062001fa2565b6200d2ad6200d2936200d28d8a6200d286600362000b1c565b90620009f1565b6200080f565b6200d2a6673bfb6c3f0e61657262007157565b9062003081565b9062001fa2565b6200d2f46200d2da6200d2d4896200d2cd600462000b41565b90620009f1565b6200080f565b6200d2ed67514abd0cf6c7bc866200717c565b9062003081565b9062001fa2565b6200d33b6200d3216200d31b886200d314600562000b66565b90620009f1565b6200080f565b6200d3346747521b1361dcc546620071a1565b9062003081565b9062001fa2565b936200d34f600662000b8b565b90620009f1565b6200080f565b6200d36f67178093843f863d14620071c6565b9062003081565b9062001fa2565b6200d3bd6200d3a36200d39d8d6200d396600762000bb0565b90620009f1565b6200080f565b6200d3b667ad1003c5d28918e7620071eb565b9062003081565b9062001fa2565b6200d4046200d3ea6200d3e48c6200d3dd60086200145b565b90620009f1565b6200080f565b6200d3fd67738450e42495bc8162007210565b9062003081565b9062001fa2565b6200d44b6200d4316200d42b8b6200d424600962002102565b90620009f1565b6200080f565b6200d44467af947c59af5e404762007235565b9062003081565b9062001fa2565b6200d4926200d4786200d4728a6200d46b600a6200214c565b90620009f1565b6200080f565b6200d48b674653fb0685084ef26200725a565b9062003081565b9062001fa2565b6200d4d96200d4bf6200d4b9896200d4b2600b62002196565b90620009f1565b6200080f565b6200d4d267057fde2062ae35bf6200727f565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6200d504886200d4fd5f620007a4565b90620009f1565b6200081e565b6200d5a26200d5846200d57e6200d5386200d532856200d52b600162000ad2565b90620009f1565b6200080f565b6200d5776200d55d6200d557876200d5505f620007a4565b90620009f1565b6200080f565b6200d57067f02a3ac068ee110b620072a4565b9062003081565b9062001fa2565b6200143f565b6200d59c886200d595600162000ad2565b90620009f1565b6200081e565b6200d63a6200d61c6200d6166200d5d06200d5ca856200d5c3600262000af7565b90620009f1565b6200080f565b6200d60f6200d5f56200d5ef876200d5e85f620007a4565b90620009f1565b6200080f565b6200d608670a3630dafb8ae2d7620072c9565b9062003081565b9062001fa2565b6200143f565b6200d634886200d62d600262000af7565b90620009f1565b6200081e565b6200d6d26200d6b46200d6ae6200d6686200d662856200d65b600362000b1c565b90620009f1565b6200080f565b6200d6a76200d68d6200d687876200d6805f620007a4565b90620009f1565b6200080f565b6200d6a067ce0dc874eaf9b55c620072ee565b9062003081565b9062001fa2565b6200143f565b6200d6cc886200d6c5600362000b1c565b90620009f1565b6200081e565b6200d76a6200d74c6200d7466200d7006200d6fa856200d6f3600462000b41565b90620009f1565b6200080f565b6200d73f6200d7256200d71f876200d7185f620007a4565b90620009f1565b6200080f565b6200d738679a95f6cff5b55c7e62007313565b9062003081565b9062001fa2565b6200143f565b6200d764886200d75d600462000b41565b90620009f1565b6200081e565b6200d8026200d7e46200d7de6200d7986200d792856200d78b600562000b66565b90620009f1565b6200080f565b6200d7d76200d7bd6200d7b7876200d7b05f620007a4565b90620009f1565b6200080f565b6200d7d067626d76abfed00c7b62007338565b9062003081565b9062001fa2565b6200143f565b6200d7fc886200d7f5600562000b66565b90620009f1565b6200081e565b6200d89a6200d87c6200d8766200d8306200d82a856200d823600662000b8b565b90620009f1565b6200080f565b6200d86f6200d8556200d84f876200d8485f620007a4565b90620009f1565b6200080f565b6200d86867a0c1cf1251c204ad6200735d565b9062003081565b9062001fa2565b6200143f565b6200d894886200d88d600662000b8b565b90620009f1565b6200081e565b6200d9326200d9146200d90e6200d8c86200d8c2856200d8bb600762000bb0565b90620009f1565b6200080f565b6200d9076200d8ed6200d8e7876200d8e05f620007a4565b90620009f1565b6200080f565b6200d90067daebd3006321052c62007382565b9062003081565b9062001fa2565b6200143f565b6200d92c886200d925600762000bb0565b90620009f1565b6200081e565b6200d9ca6200d9ac6200d9a66200d9606200d95a856200d95360086200145b565b90620009f1565b6200080f565b6200d99f6200d9856200d97f876200d9785f620007a4565b90620009f1565b6200080f565b6200d998673d4bd48b625a8065620073a7565b9062003081565b9062001fa2565b6200143f565b6200d9c4886200d9bd60086200145b565b90620009f1565b6200081e565b6200da626200da446200da3e6200d9f86200d9f2856200d9eb600962002102565b90620009f1565b6200080f565b6200da376200da1d6200da17876200da105f620007a4565b90620009f1565b6200080f565b6200da30677f1e584e071f6ed2620073cc565b9062003081565b9062001fa2565b6200143f565b6200da5c886200da55600962002102565b90620009f1565b6200081e565b6200dafa6200dadc6200dad66200da906200da8a856200da83600a6200214c565b90620009f1565b6200080f565b6200dacf6200dab56200daaf876200daa85f620007a4565b90620009f1565b6200080f565b6200dac867720574f0501caed3620073f1565b9062003081565b9062001fa2565b6200143f565b6200daf4886200daed600a6200214c565b90620009f1565b6200081e565b6200db5b6200db416200db3b6200db286200db22856200db1b600b62002196565b90620009f1565b6200080f565b936200db345f620007a4565b90620009f1565b6200080f565b6200db5467e3260ba93d23540a62007416565b9062003081565b9062001fa2565b6200143f565b6200db80846200db79600b62002196565b90620009f1565b6200081e565b5b6200d16f565b806200dba56200db9e60086200145b565b916200012b565b145f146200e5a557506200e57a816200df226200df056200deff6200e59e966200def86200deb16200de6a6200de236200dddc6200dd958a6200dd8e6200dd746200e5809f6200dd5a6200dd136200dccc6200dc856200dc3e6200dc246200dc1e6200dd6e976200dc17600162000ad2565b90620009f1565b6200080f565b6200dc3767e376678d843ce55e62006ddf565b9062003081565b6200dc7e6200dc646200dc5e8b6200dc57600262000af7565b90620009f1565b6200080f565b6200dc776766f3860d7514e7fc62006e04565b9062003081565b9062001fa2565b6200dcc56200dcab6200dca58a6200dc9e600362000b1c565b90620009f1565b6200080f565b6200dcbe677817f3dfff8b4ffa62006e29565b9062003081565b9062001fa2565b6200dd0c6200dcf26200dcec896200dce5600462000b41565b90620009f1565b6200080f565b6200dd05673929624a9def725b62006e4e565b9062003081565b9062001fa2565b6200dd536200dd396200dd33886200dd2c600562000b66565b90620009f1565b6200080f565b6200dd4c670126ca37f215a80a62006e73565b9062003081565b9062001fa2565b936200dd67600662000b8b565b90620009f1565b6200080f565b6200dd8767fce2f5d02762a30362006e98565b9062003081565b9062001fa2565b6200ddd56200ddbb6200ddb58d6200ddae600762000bb0565b90620009f1565b6200080f565b6200ddce671bc927375febbad762006ebd565b9062003081565b9062001fa2565b6200de1c6200de026200ddfc8c6200ddf560086200145b565b90620009f1565b6200080f565b6200de156785b481e5243f60bf62006ee2565b9062003081565b9062001fa2565b6200de636200de496200de438b6200de3c600962002102565b90620009f1565b6200080f565b6200de5c672d3c5f42a39c91a062006f07565b9062003081565b9062001fa2565b6200deaa6200de906200de8a8a6200de83600a6200214c565b90620009f1565b6200080f565b6200dea3670811719919351ae862006f2c565b9062003081565b9062001fa2565b6200def16200ded76200ded1896200deca600b62002196565b90620009f1565b6200080f565b6200deea67f669de0add99313162006f51565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6200df1c886200df155f620007a4565b90620009f1565b6200081e565b6200dfba6200df9c6200df966200df506200df4a856200df43600162000ad2565b90620009f1565b6200080f565b6200df8f6200df756200df6f876200df685f620007a4565b90620009f1565b6200080f565b6200df8867ab1cbd41d8c1e33562006f76565b9062003081565b9062001fa2565b6200143f565b6200dfb4886200dfad600162000ad2565b90620009f1565b6200081e565b6200e0526200e0346200e02e6200dfe86200dfe2856200dfdb600262000af7565b90620009f1565b6200080f565b6200e0276200e00d6200e007876200e0005f620007a4565b90620009f1565b6200080f565b6200e020679322ed4c0bc2df0162006f9b565b9062003081565b9062001fa2565b6200143f565b6200e04c886200e045600262000af7565b90620009f1565b6200081e565b6200e0ea6200e0cc6200e0c66200e0806200e07a856200e073600362000b1c565b90620009f1565b6200080f565b6200e0bf6200e0a56200e09f876200e0985f620007a4565b90620009f1565b6200080f565b6200e0b86751c3c0983d4284e562006fc0565b9062003081565b9062001fa2565b6200143f565b6200e0e4886200e0dd600362000b1c565b90620009f1565b6200081e565b6200e1826200e1646200e15e6200e1186200e112856200e10b600462000b41565b90620009f1565b6200080f565b6200e1576200e13d6200e137876200e1305f620007a4565b90620009f1565b6200080f565b6200e1506794178e291145c23162006fe5565b9062003081565b9062001fa2565b6200143f565b6200e17c886200e175600462000b41565b90620009f1565b6200081e565b6200e21a6200e1fc6200e1f66200e1b06200e1aa856200e1a3600562000b66565b90620009f1565b6200080f565b6200e1ef6200e1d56200e1cf876200e1c85f620007a4565b90620009f1565b6200080f565b6200e1e867fd0f1a973d6b20856200700a565b9062003081565b9062001fa2565b6200143f565b6200e214886200e20d600562000b66565b90620009f1565b6200081e565b6200e2b26200e2946200e28e6200e2486200e242856200e23b600662000b8b565b90620009f1565b6200080f565b6200e2876200e26d6200e267876200e2605f620007a4565b90620009f1565b6200080f565b6200e28067d427ad96e2b397196200702f565b9062003081565b9062001fa2565b6200143f565b6200e2ac886200e2a5600662000b8b565b90620009f1565b6200081e565b6200e34a6200e32c6200e3266200e2e06200e2da856200e2d3600762000bb0565b90620009f1565b6200080f565b6200e31f6200e3056200e2ff876200e2f85f620007a4565b90620009f1565b6200080f565b6200e318678a52437fecaac06b62007054565b9062003081565b9062001fa2565b6200143f565b6200e344886200e33d600762000bb0565b90620009f1565b6200081e565b6200e3e26200e3c46200e3be6200e3786200e372856200e36b60086200145b565b90620009f1565b6200080f565b6200e3b76200e39d6200e397876200e3905f620007a4565b90620009f1565b6200080f565b6200e3b067dc20ee4b8c4c9a8062007079565b9062003081565b9062001fa2565b6200143f565b6200e3dc886200e3d560086200145b565b90620009f1565b6200081e565b6200e47a6200e45c6200e4566200e4106200e40a856200e403600962002102565b90620009f1565b6200080f565b6200e44f6200e4356200e42f876200e4285f620007a4565b90620009f1565b6200080f565b6200e44867a2c98e9549da21006200709e565b9062003081565b9062001fa2565b6200143f565b6200e474886200e46d600962002102565b90620009f1565b6200081e565b6200e5126200e4f46200e4ee6200e4a86200e4a2856200e49b600a6200214c565b90620009f1565b6200080f565b6200e4e76200e4cd6200e4c7876200e4c05f620007a4565b90620009f1565b6200080f565b6200e4e0671603fe12613db5b6620070c3565b9062003081565b9062001fa2565b6200143f565b6200e50c886200e505600a6200214c565b90620009f1565b6200081e565b6200e5736200e5596200e5536200e5406200e53a856200e533600b62002196565b90620009f1565b6200080f565b936200e54c5f620007a4565b90620009f1565b6200080f565b6200e56c670e174929433c5505620070e8565b9062003081565b9062001fa2565b6200143f565b6200e598846200e591600b62002196565b90620009f1565b6200081e565b5b6200db87565b806200e5bd6200e5b6600962002102565b916200012b565b145f146200efbd57506200ef92816200e93a6200e91d6200e9176200efb6966200e9106200e8c96200e8826200e83b6200e7f46200e7ad8a6200e7a66200e78c6200ef989f6200e7726200e72b6200e6e46200e69d6200e6566200e63c6200e6366200e786976200e62f600162000ad2565b90620009f1565b6200080f565b6200e64f677de38bae084da92d62006ab1565b9062003081565b6200e6966200e67c6200e6768b6200e66f600262000af7565b90620009f1565b6200080f565b6200e68f675b848442237e8a9b62006ad6565b9062003081565b9062001fa2565b6200e6dd6200e6c36200e6bd8a6200e6b6600362000b1c565b90620009f1565b6200080f565b6200e6d667f6c705da84d5731062006afb565b9062003081565b9062001fa2565b6200e7246200e70a6200e704896200e6fd600462000b41565b90620009f1565b6200080f565b6200e71d6731e6a4bdb6a4901762006b20565b9062003081565b9062001fa2565b6200e76b6200e7516200e74b886200e744600562000b66565b90620009f1565b6200080f565b6200e76467889489706e5c5c0f62006b45565b9062003081565b9062001fa2565b936200e77f600662000b8b565b90620009f1565b6200080f565b6200e79f670e4a205459692a1b62006b6a565b9062003081565b9062001fa2565b6200e7ed6200e7d36200e7cd8d6200e7c6600762000bb0565b90620009f1565b6200080f565b6200e7e667bac3fa75ee26f29962006b8f565b9062003081565b9062001fa2565b6200e8346200e81a6200e8148c6200e80d60086200145b565b90620009f1565b6200080f565b6200e82d675f5894f4057d755e62006bb4565b9062003081565b9062001fa2565b6200e87b6200e8616200e85b8b6200e854600962002102565b90620009f1565b6200080f565b6200e87467b0dc3ecd724bb07662006bd9565b9062003081565b9062001fa2565b6200e8c26200e8a86200e8a28a6200e89b600a6200214c565b90620009f1565b6200080f565b6200e8bb675e34d8554a6452ba62006bfe565b9062003081565b9062001fa2565b6200e9096200e8ef6200e8e9896200e8e2600b62002196565b90620009f1565b6200080f565b6200e9026704f78fd8c1fdcc5f62006c23565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6200e934886200e92d5f620007a4565b90620009f1565b6200081e565b6200e9d26200e9b46200e9ae6200e9686200e962856200e95b600162000ad2565b90620009f1565b6200080f565b6200e9a76200e98d6200e987876200e9805f620007a4565b90620009f1565b6200080f565b6200e9a0673d4eab2b8ef5f79662006c48565b9062003081565b9062001fa2565b6200143f565b6200e9cc886200e9c5600162000ad2565b90620009f1565b6200081e565b6200ea6a6200ea4c6200ea466200ea006200e9fa856200e9f3600262000af7565b90620009f1565b6200080f565b6200ea3f6200ea256200ea1f876200ea185f620007a4565b90620009f1565b6200080f565b6200ea3867cfff421583896e2262006c6d565b9062003081565b9062001fa2565b6200143f565b6200ea64886200ea5d600262000af7565b90620009f1565b6200081e565b6200eb026200eae46200eade6200ea986200ea92856200ea8b600362000b1c565b90620009f1565b6200080f565b6200ead76200eabd6200eab7876200eab05f620007a4565b90620009f1565b6200080f565b6200ead0674143cb32d39ac3d962006c92565b9062003081565b9062001fa2565b6200143f565b6200eafc886200eaf5600362000b1c565b90620009f1565b6200081e565b6200eb9a6200eb7c6200eb766200eb306200eb2a856200eb23600462000b41565b90620009f1565b6200080f565b6200eb6f6200eb556200eb4f876200eb485f620007a4565b90620009f1565b6200080f565b6200eb686722365051b78a5b6562006cb7565b9062003081565b9062001fa2565b6200143f565b6200eb94886200eb8d600462000b41565b90620009f1565b6200081e565b6200ec326200ec146200ec0e6200ebc86200ebc2856200ebbb600562000b66565b90620009f1565b6200080f565b6200ec076200ebed6200ebe7876200ebe05f620007a4565b90620009f1565b6200080f565b6200ec00676f7fd010d027c9b662006cdc565b9062003081565b9062001fa2565b6200143f565b6200ec2c886200ec25600562000b66565b90620009f1565b6200081e565b6200ecca6200ecac6200eca66200ec606200ec5a856200ec53600662000b8b565b90620009f1565b6200080f565b6200ec9f6200ec856200ec7f876200ec785f620007a4565b90620009f1565b6200080f565b6200ec9867d9dd36fba77522ab62006d01565b9062003081565b9062001fa2565b6200143f565b6200ecc4886200ecbd600662000b8b565b90620009f1565b6200081e565b6200ed626200ed446200ed3e6200ecf86200ecf2856200eceb600762000bb0565b90620009f1565b6200080f565b6200ed376200ed1d6200ed17876200ed105f620007a4565b90620009f1565b6200080f565b6200ed3067a44cf1cb33e3716562006d26565b9062003081565b9062001fa2565b6200143f565b6200ed5c886200ed55600762000bb0565b90620009f1565b6200081e565b6200edfa6200eddc6200edd66200ed906200ed8a856200ed8360086200145b565b90620009f1565b6200080f565b6200edcf6200edb56200edaf876200eda85f620007a4565b90620009f1565b6200080f565b6200edc8673fc83d3038c8641762006d4b565b9062003081565b9062001fa2565b6200143f565b6200edf4886200eded60086200145b565b90620009f1565b6200081e565b6200ee926200ee746200ee6e6200ee286200ee22856200ee1b600962002102565b90620009f1565b6200080f565b6200ee676200ee4d6200ee47876200ee405f620007a4565b90620009f1565b6200080f565b6200ee6067c4588d418e88d27062006d70565b9062003081565b9062001fa2565b6200143f565b6200ee8c886200ee85600962002102565b90620009f1565b6200081e565b6200ef2a6200ef0c6200ef066200eec06200eeba856200eeb3600a6200214c565b90620009f1565b6200080f565b6200eeff6200eee56200eedf876200eed85f620007a4565b90620009f1565b6200080f565b6200eef867ce1320f10ab80fe262006d95565b9062003081565b9062001fa2565b6200143f565b6200ef24886200ef1d600a6200214c565b90620009f1565b6200081e565b6200ef8b6200ef716200ef6b6200ef586200ef52856200ef4b600b62002196565b90620009f1565b6200080f565b936200ef645f620007a4565b90620009f1565b6200080f565b6200ef8467db5eadbbec18de5d62006dba565b9062003081565b9062001fa2565b6200143f565b6200efb0846200efa9600b62002196565b90620009f1565b6200081e565b5b6200e59f565b806200efd56200efce600a6200214c565b916200012b565b145f146200f9d557506200f9aa816200f3526200f3356200f32f6200f9ce966200f3286200f2e16200f29a6200f2536200f20c6200f1c58a6200f1be6200f1a46200f9b09f6200f18a6200f1436200f0fc6200f0b56200f06e6200f0546200f04e6200f19e976200f047600162000ad2565b90620009f1565b6200080f565b6200f067674dd19c38779512ea62006783565b9062003081565b6200f0ae6200f0946200f08e8b6200f087600262000af7565b90620009f1565b6200080f565b6200f0a767db79ba02704620e9620067a8565b9062003081565b9062001fa2565b6200f0f56200f0db6200f0d58a6200f0ce600362000b1c565b90620009f1565b6200080f565b6200f0ee6792a29a3675a5d2be620067cd565b9062003081565b9062001fa2565b6200f13c6200f1226200f11c896200f115600462000b41565b90620009f1565b6200080f565b6200f13567d5177029fe495166620067f2565b9062003081565b9062001fa2565b6200f1836200f1696200f163886200f15c600562000b66565b90620009f1565b6200080f565b6200f17c67d32b3298a13330c162006817565b9062003081565b9062001fa2565b936200f197600662000b8b565b90620009f1565b6200080f565b6200f1b767251c4a3eb2c5f8fd6200683c565b9062003081565b9062001fa2565b6200f2056200f1eb6200f1e58d6200f1de600762000bb0565b90620009f1565b6200080f565b6200f1fe67e1c48b26e0d9882562006861565b9062003081565b9062001fa2565b6200f24c6200f2326200f22c8c6200f22560086200145b565b90620009f1565b6200080f565b6200f245673301d3362a4ffccb62006886565b9062003081565b9062001fa2565b6200f2936200f2796200f2738b6200f26c600962002102565b90620009f1565b6200080f565b6200f28c6709bb6c88de8cd178620068ab565b9062003081565b9062001fa2565b6200f2da6200f2c06200f2ba8a6200f2b3600a6200214c565b90620009f1565b6200080f565b6200f2d367dc05b676564f538a620068d0565b9062003081565b9062001fa2565b6200f3216200f3076200f301896200f2fa600b62002196565b90620009f1565b6200080f565b6200f31a6760192d883e473fee620068f5565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6200f34c886200f3455f620007a4565b90620009f1565b6200081e565b6200f3ea6200f3cc6200f3c66200f3806200f37a856200f373600162000ad2565b90620009f1565b6200080f565b6200f3bf6200f3a56200f39f876200f3985f620007a4565b90620009f1565b6200080f565b6200f3b8671183dfce7c454afd6200691a565b9062003081565b9062001fa2565b6200143f565b6200f3e4886200f3dd600162000ad2565b90620009f1565b6200081e565b6200f4826200f4646200f45e6200f4186200f412856200f40b600262000af7565b90620009f1565b6200080f565b6200f4576200f43d6200f437876200f4305f620007a4565b90620009f1565b6200080f565b6200f4506721cea4aa3d3ed9496200693f565b9062003081565b9062001fa2565b6200143f565b6200f47c886200f475600262000af7565b90620009f1565b6200081e565b6200f51a6200f4fc6200f4f66200f4b06200f4aa856200f4a3600362000b1c565b90620009f1565b6200080f565b6200f4ef6200f4d56200f4cf876200f4c85f620007a4565b90620009f1565b6200080f565b6200f4e8670fce6f70303f230462006964565b9062003081565b9062001fa2565b6200143f565b6200f514886200f50d600362000b1c565b90620009f1565b6200081e565b6200f5b26200f5946200f58e6200f5486200f542856200f53b600462000b41565b90620009f1565b6200080f565b6200f5876200f56d6200f567876200f5605f620007a4565b90620009f1565b6200080f565b6200f5806719557d34b55551be62006989565b9062003081565b9062001fa2565b6200143f565b6200f5ac886200f5a5600462000b41565b90620009f1565b6200081e565b6200f64a6200f62c6200f6266200f5e06200f5da856200f5d3600562000b66565b90620009f1565b6200080f565b6200f61f6200f6056200f5ff876200f5f85f620007a4565b90620009f1565b6200080f565b6200f618674c56f689afc5bbc9620069ae565b9062003081565b9062001fa2565b6200143f565b6200f644886200f63d600562000b66565b90620009f1565b6200081e565b6200f6e26200f6c46200f6be6200f6786200f672856200f66b600662000b8b565b90620009f1565b6200080f565b6200f6b76200f69d6200f697876200f6905f620007a4565b90620009f1565b6200080f565b6200f6b067a1e920844334f944620069d3565b9062003081565b9062001fa2565b6200143f565b6200f6dc886200f6d5600662000b8b565b90620009f1565b6200081e565b6200f77a6200f75c6200f7566200f7106200f70a856200f703600762000bb0565b90620009f1565b6200080f565b6200f74f6200f7356200f72f876200f7285f620007a4565b90620009f1565b6200080f565b6200f74867bad66d423d2ec861620069f8565b9062003081565b9062001fa2565b6200143f565b6200f774886200f76d600762000bb0565b90620009f1565b6200081e565b6200f8126200f7f46200f7ee6200f7a86200f7a2856200f79b60086200145b565b90620009f1565b6200080f565b6200f7e76200f7cd6200f7c7876200f7c05f620007a4565b90620009f1565b6200080f565b6200f7e067f318c785dc9e047962006a1d565b9062003081565b9062001fa2565b6200143f565b6200f80c886200f80560086200145b565b90620009f1565b6200081e565b6200f8aa6200f88c6200f8866200f8406200f83a856200f833600962002102565b90620009f1565b6200080f565b6200f87f6200f8656200f85f876200f8585f620007a4565b90620009f1565b6200080f565b6200f8786799e2032e765ddd8162006a42565b9062003081565b9062001fa2565b6200143f565b6200f8a4886200f89d600962002102565b90620009f1565b6200081e565b6200f9426200f9246200f91e6200f8d86200f8d2856200f8cb600a6200214c565b90620009f1565b6200080f565b6200f9176200f8fd6200f8f7876200f8f05f620007a4565b90620009f1565b6200080f565b6200f91067400ccc9906d66f4562006a67565b9062003081565b9062001fa2565b6200143f565b6200f93c886200f935600a6200214c565b90620009f1565b6200081e565b6200f9a36200f9896200f9836200f9706200f96a856200f963600b62002196565b90620009f1565b6200080f565b936200f97c5f620007a4565b90620009f1565b6200080f565b6200f99c67e1197454db2e0dd962006a8c565b9062003081565b9062001fa2565b6200143f565b6200f9c8846200f9c1600b62002196565b90620009f1565b6200081e565b5b6200efb7565b806200f9ed6200f9e6600b62002196565b916200012b565b145f14620103ed5750620103c2816200fd6a6200fd4d6200fd47620103e6966200fd406200fcf96200fcb26200fc6b6200fc246200fbdd8a6200fbd66200fbbc620103c89f6200fba26200fb5b6200fb146200facd6200fa866200fa6c6200fa666200fbb6976200fa5f600162000ad2565b90620009f1565b6200080f565b6200fa7f6716b9774801ac44a062006455565b9062003081565b6200fac66200faac6200faa68b6200fa9f600262000af7565b90620009f1565b6200080f565b6200fabf673cb8411e786d3c8e6200647a565b9062003081565b9062001fa2565b6200fb0d6200faf36200faed8a6200fae6600362000b1c565b90620009f1565b6200080f565b6200fb0667a86e9cf5050724916200649f565b9062003081565b9062001fa2565b6200fb546200fb3a6200fb34896200fb2d600462000b41565b90620009f1565b6200080f565b6200fb4d670178928152e109ae620064c4565b9062003081565b9062001fa2565b6200fb9b6200fb816200fb7b886200fb74600562000b66565b90620009f1565b6200080f565b6200fb94675317b905a6e1ab7b620064e9565b9062003081565b9062001fa2565b936200fbaf600662000b8b565b90620009f1565b6200080f565b6200fbcf67da20b3be7f53d59f6200650e565b9062003081565b9062001fa2565b6200fc1d6200fc036200fbfd8d6200fbf6600762000bb0565b90620009f1565b6200080f565b6200fc1667cb97dedecebee9ad62006533565b9062003081565b9062001fa2565b6200fc646200fc4a6200fc448c6200fc3d60086200145b565b90620009f1565b6200080f565b6200fc5d674bd545218c59f58d62006558565b9062003081565b9062001fa2565b6200fcab6200fc916200fc8b8b6200fc84600962002102565b90620009f1565b6200080f565b6200fca46777dc8d856c05a44a6200657d565b9062003081565b9062001fa2565b6200fcf26200fcd86200fcd28a6200fccb600a6200214c565b90620009f1565b6200080f565b6200fceb6787948589e4f243fd620065a2565b9062003081565b9062001fa2565b6200fd396200fd1f6200fd19896200fd12600b62002196565b90620009f1565b6200080f565b6200fd32677e5217af969952c2620065c7565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6200fd64886200fd5d5f620007a4565b90620009f1565b6200081e565b6200fe026200fde46200fdde6200fd986200fd92856200fd8b600162000ad2565b90620009f1565b6200080f565b6200fdd76200fdbd6200fdb7876200fdb05f620007a4565b90620009f1565b6200080f565b6200fdd06784d1ecc4d53d2ff1620065ec565b9062003081565b9062001fa2565b6200143f565b6200fdfc886200fdf5600162000ad2565b90620009f1565b6200081e565b6200fe9a6200fe7c6200fe766200fe306200fe2a856200fe23600262000af7565b90620009f1565b6200080f565b6200fe6f6200fe556200fe4f876200fe485f620007a4565b90620009f1565b6200080f565b6200fe6867d8af8b9ceb4e11b662006611565b9062003081565b9062001fa2565b6200143f565b6200fe94886200fe8d600262000af7565b90620009f1565b6200081e565b6200ff326200ff146200ff0e6200fec86200fec2856200febb600362000b1c565b90620009f1565b6200080f565b6200ff076200feed6200fee7876200fee05f620007a4565b90620009f1565b6200080f565b6200ff0067335856bb527b52f462006636565b9062003081565b9062001fa2565b6200143f565b6200ff2c886200ff25600362000b1c565b90620009f1565b6200081e565b6200ffca6200ffac6200ffa66200ff606200ff5a856200ff53600462000b41565b90620009f1565b6200080f565b6200ff9f6200ff856200ff7f876200ff785f620007a4565b90620009f1565b6200080f565b6200ff9867c756f17fb59be5956200665b565b9062003081565b9062001fa2565b6200143f565b6200ffc4886200ffbd600462000b41565b90620009f1565b6200081e565b62010062620100446201003e6200fff86200fff2856200ffeb600562000b66565b90620009f1565b6200080f565b620100376201001d6201001787620100105f620007a4565b90620009f1565b6200080f565b6201003067c0654e4ea5553a7862006680565b9062003081565b9062001fa2565b6200143f565b6201005c8862010055600562000b66565b90620009f1565b6200081e565b620100fa620100dc620100d6620100906201008a8562010083600662000b8b565b90620009f1565b6200080f565b620100cf620100b5620100af87620100a85f620007a4565b90620009f1565b6200080f565b620100c8679e9a46b61f2ea942620066a5565b9062003081565b9062001fa2565b6200143f565b620100f488620100ed600662000b8b565b90620009f1565b6200081e565b62010192620101746201016e6201012862010122856201011b600762000bb0565b90620009f1565b6200080f565b620101676201014d6201014787620101405f620007a4565b90620009f1565b6200080f565b620101606714fc8b5b3b809127620066ca565b9062003081565b9062001fa2565b6200143f565b6201018c8862010185600762000bb0565b90620009f1565b6200081e565b6201022a6201020c62010206620101c0620101ba85620101b360086200145b565b90620009f1565b6200080f565b620101ff620101e5620101df87620101d85f620007a4565b90620009f1565b6200080f565b620101f867d7009f0f103be413620066ef565b9062003081565b9062001fa2565b6200143f565b62010224886201021d60086200145b565b90620009f1565b6200081e565b620102c2620102a46201029e6201025862010252856201024b600962002102565b90620009f1565b6200080f565b620102976201027d6201027787620102705f620007a4565b90620009f1565b6200080f565b62010290673e0ee7b7a9fb460162006714565b9062003081565b9062001fa2565b6200143f565b620102bc88620102b5600962002102565b90620009f1565b6200081e565b6201035a6201033c62010336620102f0620102ea85620102e3600a6200214c565b90620009f1565b6200080f565b6201032f620103156201030f87620103085f620007a4565b90620009f1565b6200080f565b6201032867a74e888922085ed762006739565b9062003081565b9062001fa2565b6200143f565b62010354886201034d600a6200214c565b90620009f1565b6200081e565b620103bb620103a16201039b6201038862010382856201037b600b62002196565b90620009f1565b6200080f565b93620103945f620007a4565b90620009f1565b6200080f565b620103b467e80a7cde3d4ac5266200675e565b9062003081565b9062001fa2565b6200143f565b620103e084620103d9600b62002196565b90620009f1565b6200081e565b5b6200f9cf565b8062010405620103fe600c620009c0565b916200012b565b145f1462010e05575062010dda8162010782620107656201075f62010dfe966201075862010711620106ca620106836201063c620105f58a620105ee620105d462010de09f620105ba620105736201052c620104e56201049e620104846201047e620105ce9762010477600162000ad2565b90620009f1565b6200080f565b6201049767bc58987d06a84e4d62006127565b9062003081565b620104de620104c4620104be8b620104b7600262000af7565b90620009f1565b6200080f565b620104d7670b5d420244c9cae36200614c565b9062003081565b9062001fa2565b620105256201050b620105058a620104fe600362000b1c565b90620009f1565b6200080f565b6201051e67a3c4711b938c02c062006171565b9062003081565b9062001fa2565b6201056c620105526201054c8962010545600462000b41565b90620009f1565b6200080f565b62010565673aace640a3e0399062006196565b9062003081565b9062001fa2565b620105b36201059962010593886201058c600562000b66565b90620009f1565b6200080f565b620105ac67865a0f3249aacd8a620061bb565b9062003081565b9062001fa2565b93620105c7600662000b8b565b90620009f1565b6200080f565b620105e7678d00b2a7dbed06c7620061e0565b9062003081565b9062001fa2565b620106356201061b620106158d6201060e600762000bb0565b90620009f1565b6200080f565b6201062e676eacb905beb7e2f862006205565b9062003081565b9062001fa2565b6201067c620106626201065c8c6201065560086200145b565b90620009f1565b6200080f565b6201067567045322b216ec3ec76200622a565b9062003081565b9062001fa2565b620106c3620106a9620106a38b6201069c600962002102565b90620009f1565b6200080f565b620106bc67eb9de00d594828e66200624f565b9062003081565b9062001fa2565b6201070a620106f0620106ea8a620106e3600a6200214c565b90620009f1565b6200080f565b6201070367088c5f20df9e5c2662006274565b9062003081565b9062001fa2565b620107516201073762010731896201072a600b62002196565b90620009f1565b6200080f565b6201074a67f555f4112b19781f62006299565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6201077c88620107755f620007a4565b90620009f1565b6200081e565b6201081a620107fc620107f6620107b0620107aa85620107a3600162000ad2565b90620009f1565b6200080f565b620107ef620107d5620107cf87620107c85f620007a4565b90620009f1565b6200080f565b620107e867238aa6daa612186d620062be565b9062003081565b9062001fa2565b6200143f565b62010814886201080d600162000ad2565b90620009f1565b6200081e565b620108b2620108946201088e6201084862010842856201083b600262000af7565b90620009f1565b6200080f565b620108876201086d6201086787620108605f620007a4565b90620009f1565b6200080f565b62010880679137a5c630bad4b4620062e3565b9062003081565b9062001fa2565b6200143f565b620108ac88620108a5600262000af7565b90620009f1565b6200081e565b6201094a6201092c62010926620108e0620108da85620108d3600362000b1c565b90620009f1565b6200080f565b6201091f62010905620108ff87620108f85f620007a4565b90620009f1565b6200080f565b6201091867c7db3817870c5eda62006308565b9062003081565b9062001fa2565b6200143f565b62010944886201093d600362000b1c565b90620009f1565b6200081e565b620109e2620109c4620109be6201097862010972856201096b600462000b41565b90620009f1565b6200080f565b620109b76201099d6201099787620109905f620007a4565b90620009f1565b6200080f565b620109b067217e4f04e5718dc96200632d565b9062003081565b9062001fa2565b6200143f565b620109dc88620109d5600462000b41565b90620009f1565b6200081e565b62010a7a62010a5c62010a5662010a1062010a0a8562010a03600562000b66565b90620009f1565b6200080f565b62010a4f62010a3562010a2f8762010a285f620007a4565b90620009f1565b6200080f565b62010a4867cae814e2817bd99d62006352565b9062003081565b9062001fa2565b6200143f565b62010a748862010a6d600562000b66565b90620009f1565b6200081e565b62010b1262010af462010aee62010aa862010aa28562010a9b600662000b8b565b90620009f1565b6200080f565b62010ae762010acd62010ac78762010ac05f620007a4565b90620009f1565b6200080f565b62010ae067e3292e7ab770a8ba62006377565b9062003081565b9062001fa2565b6200143f565b62010b0c8862010b05600662000b8b565b90620009f1565b6200081e565b62010baa62010b8c62010b8662010b4062010b3a8562010b33600762000bb0565b90620009f1565b6200080f565b62010b7f62010b6562010b5f8762010b585f620007a4565b90620009f1565b6200080f565b62010b78677bb36ef70b6b94826200639c565b9062003081565b9062001fa2565b6200143f565b62010ba48862010b9d600762000bb0565b90620009f1565b6200081e565b62010c4262010c2462010c1e62010bd862010bd28562010bcb60086200145b565b90620009f1565b6200080f565b62010c1762010bfd62010bf78762010bf05f620007a4565b90620009f1565b6200080f565b62010c10673c7835fb85bca2d3620063c1565b9062003081565b9062001fa2565b6200143f565b62010c3c8862010c3560086200145b565b90620009f1565b6200081e565b62010cda62010cbc62010cb662010c7062010c6a8562010c63600962002102565b90620009f1565b6200080f565b62010caf62010c9562010c8f8762010c885f620007a4565b90620009f1565b6200080f565b62010ca867fe2cdf8ee3c25e86620063e6565b9062003081565b9062001fa2565b6200143f565b62010cd48862010ccd600962002102565b90620009f1565b6200081e565b62010d7262010d5462010d4e62010d0862010d028562010cfb600a6200214c565b90620009f1565b6200080f565b62010d4762010d2d62010d278762010d205f620007a4565b90620009f1565b6200080f565b62010d406761b3915ad7274b206200640b565b9062003081565b9062001fa2565b6200143f565b62010d6c8862010d65600a6200214c565b90620009f1565b6200081e565b62010dd362010db962010db362010da062010d9a8562010d93600b62002196565b90620009f1565b6200080f565b9362010dac5f620007a4565b90620009f1565b6200080f565b62010dcc67eab75ca7c918e4ef62006430565b9062003081565b9062001fa2565b6200143f565b62010df88462010df1600b62002196565b90620009f1565b6200081e565b5b620103e7565b8062010e1d62010e16600d62002205565b916200012b565b145f146201181d5750620117f2816201119a6201117d6201117762011816966201117062011129620110e26201109b620110546201100d8a6201100662010fec620117f89f62010fd262010f8b62010f4462010efd62010eb662010e9c62010e9662010fe69762010e8f600162000ad2565b90620009f1565b6200080f565b62010eaf67a8cedbff1813d3a762005df9565b9062003081565b62010ef662010edc62010ed68b62010ecf600262000af7565b90620009f1565b6200080f565b62010eef6750dcaee0fd27d16462005e1e565b9062003081565b9062001fa2565b62010f3d62010f2362010f1d8a62010f16600362000b1c565b90620009f1565b6200080f565b62010f3667f1cb02417e23bd8262005e43565b9062003081565b9062001fa2565b62010f8462010f6a62010f648962010f5d600462000b41565b90620009f1565b6200080f565b62010f7d67faf322786e2abe8b62005e68565b9062003081565b9062001fa2565b62010fcb62010fb162010fab8862010fa4600562000b66565b90620009f1565b6200080f565b62010fc467937a4315beb5d9b662005e8d565b9062003081565b9062001fa2565b9362010fdf600662000b8b565b90620009f1565b6200080f565b62010fff671b18992921a11d8562005eb2565b9062003081565b9062001fa2565b6201104d620110336201102d8d62011026600762000bb0565b90620009f1565b6200080f565b62011046677d66c4368b3c497b62005ed7565b9062003081565b9062001fa2565b620110946201107a620110748c6201106d60086200145b565b90620009f1565b6200080f565b6201108d670e7946317a6b4e9962005efc565b9062003081565b9062001fa2565b620110db620110c1620110bb8b620110b4600962002102565b90620009f1565b6200080f565b620110d467be4430134182978b62005f21565b9062003081565b9062001fa2565b6201112262011108620111028a620110fb600a6200214c565b90620009f1565b6200080f565b6201111b673771e82493ab262d62005f46565b9062003081565b9062001fa2565b620111696201114f620111498962011142600b62002196565b90620009f1565b6200080f565b6201116267a671690d8095ce8262005f6b565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b62011194886201118d5f620007a4565b90620009f1565b6200081e565b62011232620112146201120e620111c8620111c285620111bb600162000ad2565b90620009f1565b6200080f565b62011207620111ed620111e787620111e05f620007a4565b90620009f1565b6200080f565b6201120067d6e15ffc055e154e62005f90565b9062003081565b9062001fa2565b6200143f565b6201122c8862011225600162000ad2565b90620009f1565b6200081e565b620112ca620112ac620112a6620112606201125a8562011253600262000af7565b90620009f1565b6200080f565b6201129f620112856201127f87620112785f620007a4565b90620009f1565b6200080f565b6201129867ec67881f381a32bf62005fb5565b9062003081565b9062001fa2565b6200143f565b620112c488620112bd600262000af7565b90620009f1565b6200081e565b62011362620113446201133e620112f8620112f285620112eb600362000b1c565b90620009f1565b6200080f565b620113376201131d6201131787620113105f620007a4565b90620009f1565b6200080f565b6201133067fbb1196092bf409c62005fda565b9062003081565b9062001fa2565b6200143f565b6201135c8862011355600362000b1c565b90620009f1565b6200081e565b620113fa620113dc620113d6620113906201138a8562011383600462000b41565b90620009f1565b6200080f565b620113cf620113b5620113af87620113a85f620007a4565b90620009f1565b6200080f565b620113c867dc9d2e07830ba22662005fff565b9062003081565b9062001fa2565b6200143f565b620113f488620113ed600462000b41565b90620009f1565b6200081e565b62011492620114746201146e6201142862011422856201141b600562000b66565b90620009f1565b6200080f565b620114676201144d6201144787620114405f620007a4565b90620009f1565b6200080f565b62011460670698ef3245ff798862006024565b9062003081565b9062001fa2565b6200143f565b6201148c8862011485600562000b66565b90620009f1565b6200081e565b6201152a6201150c62011506620114c0620114ba85620114b3600662000b8b565b90620009f1565b6200080f565b620114ff620114e5620114df87620114d85f620007a4565b90620009f1565b6200080f565b620114f867194fae2974f8b57662006049565b9062003081565b9062001fa2565b6200143f565b62011524886201151d600662000b8b565b90620009f1565b6200081e565b620115c2620115a46201159e6201155862011552856201154b600762000bb0565b90620009f1565b6200080f565b620115976201157d6201157787620115705f620007a4565b90620009f1565b6200080f565b62011590677a5d9bea6ca4910e6200606e565b9062003081565b9062001fa2565b6200143f565b620115bc88620115b5600762000bb0565b90620009f1565b6200081e565b6201165a6201163c62011636620115f0620115ea85620115e360086200145b565b90620009f1565b6200080f565b6201162f620116156201160f87620116085f620007a4565b90620009f1565b6200080f565b62011628677aebfea95ccdd1c962006093565b9062003081565b9062001fa2565b6200143f565b62011654886201164d60086200145b565b90620009f1565b6200081e565b620116f2620116d4620116ce6201168862011682856201167b600962002102565b90620009f1565b6200080f565b620116c7620116ad620116a787620116a05f620007a4565b90620009f1565b6200080f565b620116c067f9bd38a67d5f0e86620060b8565b9062003081565b9062001fa2565b6200143f565b620116ec88620116e5600962002102565b90620009f1565b6200081e565b6201178a6201176c62011766620117206201171a8562011713600a6200214c565b90620009f1565b6200080f565b6201175f620117456201173f87620117385f620007a4565b90620009f1565b6200080f565b6201175867fa65539de65492d8620060dd565b9062003081565b9062001fa2565b6200143f565b62011784886201177d600a6200214c565b90620009f1565b6200081e565b620117eb620117d1620117cb620117b8620117b285620117ab600b62002196565b90620009f1565b6200080f565b93620117c45f620007a4565b90620009f1565b6200080f565b620117e467f0dfcbe7653ff78762006102565b9062003081565b9062001fa2565b6200143f565b620118108462011809600b62002196565b90620009f1565b6200081e565b5b62010dff565b80620118356201182e600e6200224f565b916200012b565b145f146201223557506201220a8162011bb262011b9562011b8f6201222e9662011b8862011b4162011afa62011ab362011a6c62011a258a62011a1e62011a04620122109f620119ea620119a36201195c62011915620118ce620118b4620118ae620119fe97620118a7600162000ad2565b90620009f1565b6200080f565b620118c767b035585f6e929d9d62005acb565b9062003081565b6201190e620118f4620118ee8b620118e7600262000af7565b90620009f1565b6200080f565b6201190767ba1579c7e219b95462005af0565b9062003081565b9062001fa2565b620119556201193b620119358a6201192e600362000b1c565b90620009f1565b6200080f565b6201194e67cb201cf846db4ba362005b15565b9062003081565b9062001fa2565b6201199c620119826201197c8962011975600462000b41565b90620009f1565b6200080f565b6201199567287bf9177372cf4562005b3a565b9062003081565b9062001fa2565b620119e3620119c9620119c388620119bc600562000b66565b90620009f1565b6200080f565b620119dc67a350e4f61147d0a662005b5f565b9062003081565b9062001fa2565b93620119f7600662000b8b565b90620009f1565b6200080f565b62011a1767d5d0ecfb50bcff9962005b84565b9062003081565b9062001fa2565b62011a6562011a4b62011a458d62011a3e600762000bb0565b90620009f1565b6200080f565b62011a5e672e166aa6c776ed2162005ba9565b9062003081565b9062001fa2565b62011aac62011a9262011a8c8c62011a8560086200145b565b90620009f1565b6200080f565b62011aa567e1e66c991990e28262005bce565b9062003081565b9062001fa2565b62011af362011ad962011ad38b62011acc600962002102565b90620009f1565b6200080f565b62011aec67662b329b01e7bb3862005bf3565b9062003081565b9062001fa2565b62011b3a62011b2062011b1a8a62011b13600a6200214c565b90620009f1565b6200080f565b62011b33678aa674b36144d9a962005c18565b9062003081565b9062001fa2565b62011b8162011b6762011b618962011b5a600b62002196565b90620009f1565b6200080f565b62011b7a67cbabf78f97f95e6562005c3d565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b62011bac8862011ba55f620007a4565b90620009f1565b6200081e565b62011c4a62011c2c62011c2662011be062011bda8562011bd3600162000ad2565b90620009f1565b6200080f565b62011c1f62011c0562011bff8762011bf85f620007a4565b90620009f1565b6200080f565b62011c18670bd87ad39042025862005c62565b9062003081565b9062001fa2565b6200143f565b62011c448862011c3d600162000ad2565b90620009f1565b6200081e565b62011ce262011cc462011cbe62011c7862011c728562011c6b600262000af7565b90620009f1565b6200080f565b62011cb762011c9d62011c978762011c905f620007a4565b90620009f1565b6200080f565b62011cb0670ad8617bca9e33c862005c87565b9062003081565b9062001fa2565b6200143f565b62011cdc8862011cd5600262000af7565b90620009f1565b6200081e565b62011d7a62011d5c62011d5662011d1062011d0a8562011d03600362000b1c565b90620009f1565b6200080f565b62011d4f62011d3562011d2f8762011d285f620007a4565b90620009f1565b6200080f565b62011d48670c00ad377a1e266662005cac565b9062003081565b9062001fa2565b6200143f565b62011d748862011d6d600362000b1c565b90620009f1565b6200081e565b62011e1262011df462011dee62011da862011da28562011d9b600462000b41565b90620009f1565b6200080f565b62011de762011dcd62011dc78762011dc05f620007a4565b90620009f1565b6200080f565b62011de0670ac6fc58b3f0518f62005cd1565b9062003081565b9062001fa2565b6200143f565b62011e0c8862011e05600462000b41565b90620009f1565b6200081e565b62011eaa62011e8c62011e8662011e4062011e3a8562011e33600562000b66565b90620009f1565b6200080f565b62011e7f62011e6562011e5f8762011e585f620007a4565b90620009f1565b6200080f565b62011e78670c0cc8a892cc417362005cf6565b9062003081565b9062001fa2565b6200143f565b62011ea48862011e9d600562000b66565b90620009f1565b6200081e565b62011f4262011f2462011f1e62011ed862011ed28562011ecb600662000b8b565b90620009f1565b6200080f565b62011f1762011efd62011ef78762011ef05f620007a4565b90620009f1565b6200080f565b62011f10670c210accb117bc2162005d1b565b9062003081565b9062001fa2565b6200143f565b62011f3c8862011f35600662000b8b565b90620009f1565b6200081e565b62011fda62011fbc62011fb662011f7062011f6a8562011f63600762000bb0565b90620009f1565b6200080f565b62011faf62011f9562011f8f8762011f885f620007a4565b90620009f1565b6200080f565b62011fa8670b73630dbb46ca1862005d40565b9062003081565b9062001fa2565b6200143f565b62011fd48862011fcd600762000bb0565b90620009f1565b6200081e565b62012072620120546201204e62012008620120028562011ffb60086200145b565b90620009f1565b6200080f565b620120476201202d6201202787620120205f620007a4565b90620009f1565b6200080f565b62012040670c8be4920cbd4a5462005d65565b9062003081565b9062001fa2565b6200143f565b6201206c886201206560086200145b565b90620009f1565b6200081e565b6201210a620120ec620120e6620120a06201209a8562012093600962002102565b90620009f1565b6200080f565b620120df620120c5620120bf87620120b85f620007a4565b90620009f1565b6200080f565b620120d8670bfe877a21be169062005d8a565b9062003081565b9062001fa2565b6200143f565b6201210488620120fd600962002102565b90620009f1565b6200081e565b620121a2620121846201217e6201213862012132856201212b600a6200214c565b90620009f1565b6200080f565b620121776201215d6201215787620121505f620007a4565b90620009f1565b6200080f565b62012170670ae790559b0ded8162005daf565b9062003081565b9062001fa2565b6200143f565b6201219c8862012195600a6200214c565b90620009f1565b6200081e565b62012203620121e9620121e3620121d0620121ca85620121c3600b62002196565b90620009f1565b6200080f565b93620121dc5f620007a4565b90620009f1565b6200080f565b620121fc670bf50db2f8d6ce3162005dd4565b9062003081565b9062001fa2565b6200143f565b620122288462012221600b62002196565b90620009f1565b6200081e565b5b62011817565b806201224d62012246600f62002299565b916200012b565b145f1462012c42575062012c1781620125ca620125ad620125a762012c3b96620125a06201255962012512620124cb620124846201243d8a620124366201241c62012c1d9f62012402620123bb620123746201232d620122e6620122cc620122c66201241697620122bf600162000ad2565b90620009f1565b6200080f565b620122df67eec24b15a06b53fe6200579d565b9062003081565b620123266201230c620123068b620122ff600262000af7565b90620009f1565b6200080f565b6201231f67c8a7aa07c5633533620057c2565b9062003081565b9062001fa2565b6201236d620123536201234d8a62012346600362000b1c565b90620009f1565b6200080f565b6201236667efe9c6fa4311ad51620057e7565b9062003081565b9062001fa2565b620123b46201239a62012394896201238d600462000b41565b90620009f1565b6200080f565b620123ad67b9173f13977109a16200580c565b9062003081565b9062001fa2565b620123fb620123e1620123db88620123d4600562000b66565b90620009f1565b6200080f565b620123f46769ce43c9cc94aedc62005831565b9062003081565b9062001fa2565b936201240f600662000b8b565b90620009f1565b6200080f565b6201242f67ecf623c9cd11881562005856565b9062003081565b9062001fa2565b6201247d620124636201245d8d62012456600762000bb0565b90620009f1565b6200080f565b620124766728625def198c33c76200587b565b9062003081565b9062001fa2565b620124c4620124aa620124a48c6201249d60086200145b565b90620009f1565b6200080f565b620124bd67ccfc5f7de5c3636a620058a0565b9062003081565b9062001fa2565b6201250b620124f1620124eb8b620124e4600962002102565b90620009f1565b6200080f565b6201250467f5e6c40f1621c299620058c5565b9062003081565b9062001fa2565b6201255262012538620125328a6201252b600a6200214c565b90620009f1565b6200080f565b6201254b67cec0e58c34cb64b1620058ea565b9062003081565b9062001fa2565b620125996201257f620125798962012572600b62002196565b90620009f1565b6200080f565b6201259267a868ea113387939f6200590f565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b620125c488620125bd5f620007a4565b90620009f1565b6200081e565b62012661620126436201263d620125f8620125f285620125eb600162000ad2565b90620009f1565b6200080f565b620126366201261d6201261787620126105f620007a4565b90620009f1565b6200080f565b6201262f660cf29427ff7c5862005934565b9062003081565b9062001fa2565b6200143f565b6201265b8862012654600162000ad2565b90620009f1565b6200081e565b620126f8620126da620126d46201268f620126898562012682600262000af7565b90620009f1565b6200080f565b620126cd620126b4620126ae87620126a75f620007a4565b90620009f1565b6200080f565b620126c6660bd9b3cf49eec862005959565b9062003081565b9062001fa2565b6200143f565b620126f288620126eb600262000af7565b90620009f1565b6200081e565b6201278f620127716201276b62012726620127208562012719600362000b1c565b90620009f1565b6200080f565b620127646201274b62012745876201273e5f620007a4565b90620009f1565b6200080f565b6201275d660d1dc8aa81fb266200597e565b9062003081565b9062001fa2565b6200143f565b620127898862012782600362000b1c565b90620009f1565b6200081e565b620128266201280862012802620127bd620127b785620127b0600462000b41565b90620009f1565b6200080f565b620127fb620127e2620127dc87620127d55f620007a4565b90620009f1565b6200080f565b620127f4660bc792d5c394ef620059a3565b9062003081565b9062001fa2565b6200143f565b620128208862012819600462000b41565b90620009f1565b6200081e565b620128bd6201289f62012899620128546201284e8562012847600562000b66565b90620009f1565b6200080f565b620128926201287962012873876201286c5f620007a4565b90620009f1565b6200080f565b6201288b660d2ae0b2266453620059c8565b9062003081565b9062001fa2565b6200143f565b620128b788620128b0600562000b66565b90620009f1565b6200081e565b620129546201293662012930620128eb620128e585620128de600662000b8b565b90620009f1565b6200080f565b62012929620129106201290a87620129035f620007a4565b90620009f1565b6200080f565b62012922660d413f12c496c1620059ed565b9062003081565b9062001fa2565b6200143f565b6201294e8862012947600662000b8b565b90620009f1565b6200081e565b620129eb620129cd620129c7620129826201297c8562012975600762000bb0565b90620009f1565b6200080f565b620129c0620129a7620129a1876201299a5f620007a4565b90620009f1565b6200080f565b620129b9660c84128cfed61862005a12565b9062003081565b9062001fa2565b6200143f565b620129e588620129de600762000bb0565b90620009f1565b6200081e565b62012a8262012a6462012a5e62012a1962012a138562012a0c60086200145b565b90620009f1565b6200080f565b62012a5762012a3e62012a388762012a315f620007a4565b90620009f1565b6200080f565b62012a50660db5ebd48fc0d462005a37565b9062003081565b9062001fa2565b6200143f565b62012a7c8862012a7560086200145b565b90620009f1565b6200081e565b62012b1962012afb62012af562012ab062012aaa8562012aa3600962002102565b90620009f1565b6200080f565b62012aee62012ad562012acf8762012ac85f620007a4565b90620009f1565b6200080f565b62012ae7660d1b77326dcb9062005a5c565b9062003081565b9062001fa2565b6200143f565b62012b138862012b0c600962002102565b90620009f1565b6200081e565b62012bb062012b9262012b8c62012b4762012b418562012b3a600a6200214c565b90620009f1565b6200080f565b62012b8562012b6c62012b668762012b5f5f620007a4565b90620009f1565b6200080f565b62012b7e660beb0ccc14542162005a81565b9062003081565b9062001fa2565b6200143f565b62012baa8862012ba3600a6200214c565b90620009f1565b6200081e565b62012c1062012bf762012bf162012bde62012bd88562012bd1600b62002196565b90620009f1565b6200080f565b9362012bea5f620007a4565b90620009f1565b6200080f565b62012c09660d10e5b22b11d162005aa6565b9062003081565b9062001fa2565b6200143f565b62012c358462012c2e600b62002196565b90620009f1565b6200081e565b5b6201222f565b8062012c5a62012c536010620022e3565b916200012b565b145f14620136445750620136198162012fd762012fba62012fb46201363d9662012fad62012f6662012f1f62012ed862012e9162012e4a8a62012e4362012e296201361f9f62012e0f62012dc862012d8162012d3a62012cf362012cd962012cd362012e239762012ccc600162000ad2565b90620009f1565b6200080f565b62012cec67d8dddbdc5ce4ef456200546f565b9062003081565b62012d3362012d1962012d138b62012d0c600262000af7565b90620009f1565b6200080f565b62012d2c67acfc51de8131458c62005494565b9062003081565b9062001fa2565b62012d7a62012d6062012d5a8a62012d53600362000b1c565b90620009f1565b6200080f565b62012d7367146bb3c0fe499ac0620054b9565b9062003081565b9062001fa2565b62012dc162012da762012da18962012d9a600462000b41565b90620009f1565b6200080f565b62012dba679e65309f15943903620054de565b9062003081565b9062001fa2565b62012e0862012dee62012de88862012de1600562000b66565b90620009f1565b6200080f565b62012e016780d0ad980773aa7062005503565b9062003081565b9062001fa2565b9362012e1c600662000b8b565b90620009f1565b6200080f565b62012e3c67f97817d4ddbf060762005528565b9062003081565b9062001fa2565b62012e8a62012e7062012e6a8d62012e63600762000bb0565b90620009f1565b6200080f565b62012e8367e4626620a75ba2766200554d565b9062003081565b9062001fa2565b62012ed162012eb762012eb18c62012eaa60086200145b565b90620009f1565b6200080f565b62012eca670dfdc7fd6fc74f6662005572565b9062003081565b9062001fa2565b62012f1862012efe62012ef88b62012ef1600962002102565b90620009f1565b6200080f565b62012f1167f464864ad6f2bb9362005597565b9062003081565b9062001fa2565b62012f5f62012f4562012f3f8a62012f38600a6200214c565b90620009f1565b6200080f565b62012f586702d55e52a5d44414620055bc565b9062003081565b9062001fa2565b62012fa662012f8c62012f868962012f7f600b62002196565b90620009f1565b6200080f565b62012f9f67dd8de62487c40925620055e1565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b62012fd18862012fca5f620007a4565b90620009f1565b6200081e565b6201306d6201304f620130496201300562012fff8562012ff8600162000ad2565b90620009f1565b6200080f565b620130426201302a62013024876201301d5f620007a4565b90620009f1565b6200080f565b6201303b650e24c99adad862005606565b9062003081565b9062001fa2565b6200143f565b620130678862013060600162000ad2565b90620009f1565b6200081e565b62013103620130e5620130df6201309b62013095856201308e600262000af7565b90620009f1565b6200080f565b620130d8620130c0620130ba87620130b35f620007a4565b90620009f1565b6200080f565b620130d1650cf389ed4bc86200562b565b9062003081565b9062001fa2565b6200143f565b620130fd88620130f6600262000af7565b90620009f1565b6200081e565b620131996201317b62013175620131316201312b8562013124600362000b1c565b90620009f1565b6200080f565b6201316e620131566201315087620131495f620007a4565b90620009f1565b6200080f565b62013167650e580cbf696662005650565b9062003081565b9062001fa2565b6200143f565b62013193886201318c600362000b1c565b90620009f1565b6200081e565b6201322f620132116201320b620131c7620131c185620131ba600462000b41565b90620009f1565b6200080f565b62013204620131ec620131e687620131df5f620007a4565b90620009f1565b6200080f565b620131fd650cde5fd7e04f62005675565b9062003081565b9062001fa2565b6200143f565b620132298862013222600462000b41565b90620009f1565b6200081e565b620132c5620132a7620132a16201325d620132578562013250600562000b66565b90620009f1565b6200080f565b6201329a620132826201327c87620132755f620007a4565b90620009f1565b6200080f565b62013293650e63628041b36200569a565b9062003081565b9062001fa2565b6200143f565b620132bf88620132b8600562000b66565b90620009f1565b6200081e565b6201335b6201333d62013337620132f3620132ed85620132e6600662000b8b565b90620009f1565b6200080f565b620133306201331862013312876201330b5f620007a4565b90620009f1565b6200080f565b62013329650e7e81a87361620056bf565b9062003081565b9062001fa2565b6200143f565b62013355886201334e600662000b8b565b90620009f1565b6200081e565b620133f1620133d3620133cd6201338962013383856201337c600762000bb0565b90620009f1565b6200080f565b620133c6620133ae620133a887620133a15f620007a4565b90620009f1565b6200080f565b620133bf650dabe78f6d98620056e4565b9062003081565b9062001fa2565b6200143f565b620133eb88620133e4600762000bb0565b90620009f1565b6200081e565b6201348762013469620134636201341f62013419856201341260086200145b565b90620009f1565b6200080f565b6201345c620134446201343e87620134375f620007a4565b90620009f1565b6200080f565b62013455650efb14cac55462005709565b9062003081565b9062001fa2565b6200143f565b62013481886201347a60086200145b565b90620009f1565b6200081e565b6201351d620134ff620134f9620134b5620134af85620134a8600962002102565b90620009f1565b6200080f565b620134f2620134da620134d487620134cd5f620007a4565b90620009f1565b6200080f565b620134eb650e5574743b106200572e565b9062003081565b9062001fa2565b6200143f565b620135178862013510600962002102565b90620009f1565b6200081e565b620135b3620135956201358f6201354b62013545856201353e600a6200214c565b90620009f1565b6200080f565b62013588620135706201356a87620135635f620007a4565b90620009f1565b6200080f565b62013581650d05709f42c162005753565b9062003081565b9062001fa2565b6200143f565b620135ad88620135a6600a6200214c565b90620009f1565b6200081e565b62013612620135fa620135f4620135e1620135db85620135d4600b62002196565b90620009f1565b6200080f565b93620135ed5f620007a4565b90620009f1565b6200080f565b6201360b650e4690c96af162005778565b9062003081565b9062001fa2565b6200143f565b620136378462013630600b62002196565b90620009f1565b6200081e565b5b62012c3c565b806201365c6201365560116200232d565b916200012b565b145f146201403b57506201401081620139d9620139bc620139b66201403496620139af6201396862013921620138da620138936201384c8a620138456201382b620140169f62013811620137ca620137836201373c620136f5620136db620136d56201382597620136ce600162000ad2565b90620009f1565b6200080f565b620136ee67c15acf44759545a362005141565b9062003081565b620137356201371b620137158b6201370e600262000af7565b90620009f1565b6200080f565b6201372e67cbfdcf39869719d462005166565b9062003081565b9062001fa2565b6201377c620137626201375c8a62013755600362000b1c565b90620009f1565b6200080f565b620137756733f62042e2f802256200518b565b9062003081565b9062001fa2565b620137c3620137a9620137a3896201379c600462000b41565b90620009f1565b6200080f565b620137bc672599c5ead81d8fa3620051b0565b9062003081565b9062001fa2565b6201380a620137f0620137ea88620137e3600562000b66565b90620009f1565b6200080f565b62013803670b306cb6c1d7c8d0620051d5565b9062003081565b9062001fa2565b936201381e600662000b8b565b90620009f1565b6200080f565b6201383e67658c80d3df3729b1620051fa565b9062003081565b9062001fa2565b6201388c620138726201386c8d62013865600762000bb0565b90620009f1565b6200080f565b6201388567e8d1b2b21b41429c6200521f565b9062003081565b9062001fa2565b620138d3620138b9620138b38c620138ac60086200145b565b90620009f1565b6200080f565b620138cc67a1b67f09d4b3ccb862005244565b9062003081565b9062001fa2565b6201391a62013900620138fa8b620138f3600962002102565b90620009f1565b6200080f565b62013913670e1adf8b8443718062005269565b9062003081565b9062001fa2565b6201396162013947620139418a6201393a600a6200214c565b90620009f1565b6200080f565b6201395a670d593a5e584af47b6200528e565b9062003081565b9062001fa2565b620139a86201398e620139888962013981600b62002196565b90620009f1565b6200080f565b620139a167a023d94c56e151c7620052b3565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b620139d388620139cc5f620007a4565b90620009f1565b6200081e565b62013a6e62013a5062013a4a62013a0762013a0185620139fa600162000ad2565b90620009f1565b6200080f565b62013a4362013a2c62013a268762013a1f5f620007a4565b90620009f1565b6200080f565b62013a3c640f7157bc98620052d8565b9062003081565b9062001fa2565b6200143f565b62013a688862013a61600162000ad2565b90620009f1565b6200081e565b62013b0362013ae562013adf62013a9c62013a968562013a8f600262000af7565b90620009f1565b6200080f565b62013ad862013ac162013abb8762013ab45f620007a4565b90620009f1565b6200080f565b62013ad1640e3006d948620052fd565b9062003081565b9062001fa2565b6200143f565b62013afd8862013af6600262000af7565b90620009f1565b6200081e565b62013b9862013b7a62013b7462013b3162013b2b8562013b24600362000b1c565b90620009f1565b6200080f565b62013b6d62013b5662013b508762013b495f620007a4565b90620009f1565b6200080f565b62013b66640fa65811e662005322565b9062003081565b9062001fa2565b6200143f565b62013b928862013b8b600362000b1c565b90620009f1565b6200081e565b62013c2d62013c0f62013c0962013bc662013bc08562013bb9600462000b41565b90620009f1565b6200080f565b62013c0262013beb62013be58762013bde5f620007a4565b90620009f1565b6200080f565b62013bfb640e0d127e2f62005347565b9062003081565b9062001fa2565b6200143f565b62013c278862013c20600462000b41565b90620009f1565b6200081e565b62013cc262013ca462013c9e62013c5b62013c558562013c4e600562000b66565b90620009f1565b6200080f565b62013c9762013c8062013c7a8762013c735f620007a4565b90620009f1565b6200080f565b62013c90640fc18bfe536200536c565b9062003081565b9062001fa2565b6200143f565b62013cbc8862013cb5600562000b66565b90620009f1565b6200081e565b62013d5762013d3962013d3362013cf062013cea8562013ce3600662000b8b565b90620009f1565b6200080f565b62013d2c62013d1562013d0f8762013d085f620007a4565b90620009f1565b6200080f565b62013d25640fd002d90162005391565b9062003081565b9062001fa2565b6200143f565b62013d518862013d4a600662000b8b565b90620009f1565b6200081e565b62013dec62013dce62013dc862013d8562013d7f8562013d78600762000bb0565b90620009f1565b6200080f565b62013dc162013daa62013da48762013d9d5f620007a4565b90620009f1565b6200080f565b62013dba640eed6461d8620053b6565b9062003081565b9062001fa2565b6200143f565b62013de68862013ddf600762000bb0565b90620009f1565b6200081e565b62013e8162013e6362013e5d62013e1a62013e148562013e0d60086200145b565b90620009f1565b6200080f565b62013e5662013e3f62013e398762013e325f620007a4565b90620009f1565b6200080f565b62013e4f641068562754620053db565b9062003081565b9062001fa2565b6200143f565b62013e7b8862013e7460086200145b565b90620009f1565b6200081e565b62013f1662013ef862013ef262013eaf62013ea98562013ea2600962002102565b90620009f1565b6200080f565b62013eeb62013ed462013ece8762013ec75f620007a4565b90620009f1565b6200080f565b62013ee4640fa0236f5062005400565b9062003081565b9062001fa2565b6200143f565b62013f108862013f09600962002102565b90620009f1565b6200081e565b62013fab62013f8d62013f8762013f4462013f3e8562013f37600a6200214c565b90620009f1565b6200080f565b62013f8062013f6962013f638762013f5c5f620007a4565b90620009f1565b6200080f565b62013f79640e3af13ee162005425565b9062003081565b9062001fa2565b6200143f565b62013fa58862013f9e600a6200214c565b90620009f1565b6200081e565b6201400962013ff262013fec62013fd962013fd38562013fcc600b62002196565b90620009f1565b6200080f565b9362013fe55f620007a4565b90620009f1565b6200080f565b62014002640fa460f6d16200544a565b9062003081565b9062001fa2565b6200143f565b6201402e8462014027600b62002196565b90620009f1565b6200081e565b5b6201363e565b80620140536201404c601262002377565b916200012b565b145f1462014a275750620149fc81620143d0620143b3620143ad62014a2096620143a66201435f62014318620142d16201428a620142438a6201423c6201422262014a029f62014208620141c16201417a62014133620140ec620140d2620140cc6201421c97620140c5600162000ad2565b90620009f1565b6200080f565b620140e56749026cc3a4afc5a662004e13565b9062003081565b6201412c620141126201410c8b62014105600262000af7565b90620009f1565b6200080f565b6201412567e06dff00ab25b91b62004e38565b9062003081565b9062001fa2565b6201417362014159620141538a6201414c600362000b1c565b90620009f1565b6200080f565b6201416c670ab38c561e8850ff62004e5d565b9062003081565b9062001fa2565b620141ba620141a06201419a8962014193600462000b41565b90620009f1565b6200080f565b620141b36792c3c8275e105eeb62004e82565b9062003081565b9062001fa2565b62014201620141e7620141e188620141da600562000b66565b90620009f1565b6200080f565b620141fa67b65256e546889bd062004ea7565b9062003081565b9062001fa2565b9362014215600662000b8b565b90620009f1565b6200080f565b62014235673c0468236ea142f662004ecc565b9062003081565b9062001fa2565b6201428362014269620142638d6201425c600762000bb0565b90620009f1565b6200080f565b6201427c67ee61766b889e18f262004ef1565b9062003081565b9062001fa2565b620142ca620142b0620142aa8c620142a360086200145b565b90620009f1565b6200080f565b620142c367a206f41b12c3041562004f16565b9062003081565b9062001fa2565b62014311620142f7620142f18b620142ea600962002102565b90620009f1565b6200080f565b6201430a6702fe9d756c9f12d162004f3b565b9062003081565b9062001fa2565b620143586201433e620143388a62014331600a6200214c565b90620009f1565b6200080f565b6201435167e9633210630cbf1262004f60565b9062003081565b9062001fa2565b6201439f620143856201437f8962014378600b62002196565b90620009f1565b6200080f565b62014398671ffea9fe85a0b0b162004f85565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b620143ca88620143c35f620007a4565b90620009f1565b6200081e565b620144646201444662014440620143fe620143f885620143f1600162000ad2565b90620009f1565b6200080f565b62014439620144236201441d87620144165f620007a4565b90620009f1565b6200080f565b62014432631113173862004faa565b9062003081565b9062001fa2565b6200143f565b6201445e8862014457600162000ad2565b90620009f1565b6200081e565b620144f8620144da620144d4620144926201448c8562014485600262000af7565b90620009f1565b6200080f565b620144cd620144b7620144b187620144aa5f620007a4565b90620009f1565b6200080f565b620144c6630f56d58862004fcf565b9062003081565b9062001fa2565b6200143f565b620144f288620144eb600262000af7565b90620009f1565b6200081e565b6201458c6201456e6201456862014526620145208562014519600362000b1c565b90620009f1565b6200080f565b620145616201454b62014545876201453e5f620007a4565b90620009f1565b6200080f565b6201455a6311050f8662004ff4565b9062003081565b9062001fa2565b6200143f565b62014586886201457f600362000b1c565b90620009f1565b6200081e565b6201462062014602620145fc620145ba620145b485620145ad600462000b41565b90620009f1565b6200080f565b620145f5620145df620145d987620145d25f620007a4565b90620009f1565b6200080f565b620145ee630f848f4f62005019565b9062003081565b9062001fa2565b6200143f565b6201461a8862014613600462000b41565b90620009f1565b6200081e565b620146b462014696620146906201464e620146488562014641600562000b66565b90620009f1565b6200080f565b62014689620146736201466d87620146665f620007a4565b90620009f1565b6200080f565b6201468263111527d36200503e565b9062003081565b9062001fa2565b6200143f565b620146ae88620146a7600562000b66565b90620009f1565b6200081e565b620147486201472a62014724620146e2620146dc85620146d5600662000b8b565b90620009f1565b6200080f565b6201471d620147076201470187620146fa5f620007a4565b90620009f1565b6200080f565b6201471663114369a162005063565b9062003081565b9062001fa2565b6200143f565b62014742886201473b600662000b8b565b90620009f1565b6200081e565b620147dc620147be620147b862014776620147708562014769600762000bb0565b90620009f1565b6200080f565b620147b16201479b62014795876201478e5f620007a4565b90620009f1565b6200080f565b620147aa63106f2f3862005088565b9062003081565b9062001fa2565b6200143f565b620147d688620147cf600762000bb0565b90620009f1565b6200081e565b62014870620148526201484c6201480a6201480485620147fd60086200145b565b90620009f1565b6200080f565b620148456201482f6201482987620148225f620007a4565b90620009f1565b6200080f565b6201483e6311e2ca94620050ad565b9062003081565b9062001fa2565b6200143f565b6201486a886201486360086200145b565b90620009f1565b6200081e565b62014904620148e6620148e06201489e620148988562014891600962002102565b90620009f1565b6200080f565b620148d9620148c3620148bd87620148b65f620007a4565b90620009f1565b6200080f565b620148d263110a29f0620050d2565b9062003081565b9062001fa2565b6200143f565b620148fe88620148f7600962002102565b90620009f1565b6200081e565b620149986201497a62014974620149326201492c8562014925600a6200214c565b90620009f1565b6200080f565b6201496d6201495762014951876201494a5f620007a4565b90620009f1565b6200080f565b62014966630fa9f5c1620050f7565b9062003081565b9062001fa2565b6200143f565b62014992886201498b600a6200214c565b90620009f1565b6200081e565b620149f5620149df620149d9620149c6620149c085620149b9600b62002196565b90620009f1565b6200080f565b93620149d25f620007a4565b90620009f1565b6200080f565b620149ee6310f625d16200511c565b9062003081565b9062001fa2565b6200143f565b62014a1a8462014a13600b62002196565b90620009f1565b6200081e565b5b62014035565b8062014a3f62014a386013620023c1565b916200012b565b145f14620154085750620153dd8162014dbc62014d9f62014d99620154019662014d9262014d4b62014d0462014cbd62014c7662014c2f8a62014c2862014c0e620153e39f62014bf462014bad62014b6662014b1f62014ad862014abe62014ab862014c089762014ab1600162000ad2565b90620009f1565b6200080f565b62014ad16781d1ae8cc50240f362004ae5565b9062003081565b62014b1862014afe62014af88b62014af1600262000af7565b90620009f1565b6200080f565b62014b1167f4c77a079a4607d762004b0a565b9062003081565b9062001fa2565b62014b5f62014b4562014b3f8a62014b38600362000b1c565b90620009f1565b6200080f565b62014b5867ed446b2315e3efc162004b2f565b9062003081565b9062001fa2565b62014ba662014b8c62014b868962014b7f600462000b41565b90620009f1565b6200080f565b62014b9f670b0a6b70915178c362004b54565b9062003081565b9062001fa2565b62014bed62014bd362014bcd8862014bc6600562000b66565b90620009f1565b6200080f565b62014be667b11ff3e089f15d9a62004b79565b9062003081565b9062001fa2565b9362014c01600662000b8b565b90620009f1565b6200080f565b62014c21671d4dba0b7ae9cc1862004b9e565b9062003081565b9062001fa2565b62014c6f62014c5562014c4f8d62014c48600762000bb0565b90620009f1565b6200080f565b62014c686765d74e2f43b48d0562004bc3565b9062003081565b9062001fa2565b62014cb662014c9c62014c968c62014c8f60086200145b565b90620009f1565b6200080f565b62014caf67a2df8c6b8ae0804a62004be8565b9062003081565b9062001fa2565b62014cfd62014ce362014cdd8b62014cd6600962002102565b90620009f1565b6200080f565b62014cf667a4e6f0a8c33348a662004c0d565b9062003081565b9062001fa2565b62014d4462014d2a62014d248a62014d1d600a6200214c565b90620009f1565b6200080f565b62014d3d67c0a26efc7be5669b62004c32565b9062003081565b9062001fa2565b62014d8b62014d7162014d6b8962014d64600b62002196565b90620009f1565b6200080f565b62014d8467a6b6582c547d0d6062004c57565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b62014db68862014daf5f620007a4565b90620009f1565b6200081e565b62014e4f62014e3162014e2b62014dea62014de48562014ddd600162000ad2565b90620009f1565b6200080f565b62014e2462014e0f62014e098762014e025f620007a4565b90620009f1565b6200080f565b62014e1d6211f71862004c7c565b9062003081565b9062001fa2565b6200143f565b62014e498862014e42600162000ad2565b90620009f1565b6200081e565b62014ee262014ec462014ebe62014e7d62014e778562014e70600262000af7565b90620009f1565b6200080f565b62014eb762014ea262014e9c8762014e955f620007a4565b90620009f1565b6200080f565b62014eb06210b6c862004ca1565b9062003081565b9062001fa2565b6200143f565b62014edc8862014ed5600262000af7565b90620009f1565b6200081e565b62014f7562014f5762014f5162014f1062014f0a8562014f03600362000b1c565b90620009f1565b6200080f565b62014f4a62014f3562014f2f8762014f285f620007a4565b90620009f1565b6200080f565b62014f4362134a9662004cc6565b9062003081565b9062001fa2565b6200143f565b62014f6f8862014f68600362000b1c565b90620009f1565b6200081e565b6201500862014fea62014fe462014fa362014f9d8562014f96600462000b41565b90620009f1565b6200080f565b62014fdd62014fc862014fc28762014fbb5f620007a4565b90620009f1565b6200080f565b62014fd66210cf7f62004ceb565b9062003081565b9062001fa2565b6200143f565b620150028862014ffb600462000b41565b90620009f1565b6200081e565b6201509b6201507d6201507762015036620150308562015029600562000b66565b90620009f1565b6200080f565b620150706201505b62015055876201504e5f620007a4565b90620009f1565b6200080f565b6201506962124d0362004d10565b9062003081565b9062001fa2565b6200143f565b62015095886201508e600562000b66565b90620009f1565b6200081e565b6201512e620151106201510a620150c9620150c385620150bc600662000b8b565b90620009f1565b6200080f565b62015103620150ee620150e887620150e15f620007a4565b90620009f1565b6200080f565b620150fc6213f8a162004d35565b9062003081565b9062001fa2565b6200143f565b620151288862015121600662000b8b565b90620009f1565b6200081e565b620151c1620151a36201519d6201515c62015156856201514f600762000bb0565b90620009f1565b6200080f565b62015196620151816201517b87620151745f620007a4565b90620009f1565b6200080f565b6201518f62117c5862004d5a565b9062003081565b9062001fa2565b6200143f565b620151bb88620151b4600762000bb0565b90620009f1565b6200081e565b620152546201523662015230620151ef620151e985620151e260086200145b565b90620009f1565b6200080f565b62015229620152146201520e87620152075f620007a4565b90620009f1565b6200080f565b6201522262132c9462004d7f565b9062003081565b9062001fa2565b6200143f565b6201524e886201524760086200145b565b90620009f1565b6200081e565b620152e7620152c9620152c3620152826201527c8562015275600962002102565b90620009f1565b6200080f565b620152bc620152a7620152a1876201529a5f620007a4565b90620009f1565b6200080f565b620152b562134fc062004da4565b9062003081565b9062001fa2565b6200143f565b620152e188620152da600962002102565b90620009f1565b6200081e565b6201537a6201535c62015356620153156201530f8562015308600a6200214c565b90620009f1565b6200080f565b6201534f6201533a62015334876201532d5f620007a4565b90620009f1565b6200080f565b620153486210a09162004dc9565b9062003081565b9062001fa2565b6200143f565b62015374886201536d600a6200214c565b90620009f1565b6200081e565b620153d6620153c1620153bb620153a8620153a2856201539b600b62002196565b90620009f1565b6200080f565b93620153b45f620007a4565b90620009f1565b6200080f565b620153cf6212896162004dee565b9062003081565b9062001fa2565b6200143f565b620153fb84620153f4600b62002196565b90620009f1565b6200081e565b5b62014a21565b80620154206201541960146200240b565b916200012b565b145f1462015dde575062015db3816201579d620157806201577a62015dd796620157736201572c620156e56201569e62015657620156108a62015609620155ef62015db99f620155d56201558e6201554762015500620154b96201549f62015499620155e99762015492600162000ad2565b90620009f1565b6200080f565b620154b26784afc741f1c13213620047b7565b9062003081565b620154f9620154df620154d98b620154d2600262000af7565b90620009f1565b6200080f565b620154f2672f8f43734fc906f3620047dc565b9062003081565b9062001fa2565b6201554062015526620155208a62015519600362000b1c565b90620009f1565b6200080f565b6201553967de682d72da0a02d962004801565b9062003081565b9062001fa2565b620155876201556d620155678962015560600462000b41565b90620009f1565b6200080f565b62015580670bb005236adb9ef262004826565b9062003081565b9062001fa2565b620155ce620155b4620155ae88620155a7600562000b66565b90620009f1565b6200080f565b620155c7675bdf35c10a8b56246200484b565b9062003081565b9062001fa2565b93620155e2600662000b8b565b90620009f1565b6200080f565b62015602670739a8a34395001062004870565b9062003081565b9062001fa2565b6201565062015636620156308d62015629600762000bb0565b90620009f1565b6200080f565b620156496752f515f44785cfbc62004895565b9062003081565b9062001fa2565b620156976201567d620156778c6201567060086200145b565b90620009f1565b6200080f565b6201569067cbaf4e5d82856c60620048ba565b9062003081565b9062001fa2565b620156de620156c4620156be8b620156b7600962002102565b90620009f1565b6200080f565b620156d767ac9ea09074e3e150620048df565b9062003081565b9062001fa2565b620157256201570b620157058a620156fe600a6200214c565b90620009f1565b6200080f565b6201571e678f0fa011a2035fb062004904565b9062003081565b9062001fa2565b6201576c620157526201574c8962015745600b62002196565b90620009f1565b6200080f565b62015765671a37905d8450904a62004929565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b6201579788620157905f620007a4565b90620009f1565b6200081e565b6201582f620158116201580b620157cb620157c585620157be600162000ad2565b90620009f1565b6200080f565b62015804620157f0620157ea87620157e35f620007a4565b90620009f1565b6200080f565b620157fd6113006200494e565b9062003081565b9062001fa2565b6200143f565b620158298862015822600162000ad2565b90620009f1565b6200081e565b620158c1620158a36201589d6201585d620158578562015850600262000af7565b90620009f1565b6200080f565b62015896620158826201587c87620158755f620007a4565b90620009f1565b6200080f565b6201588f61175062004973565b9062003081565b9062001fa2565b6200143f565b620158bb88620158b4600262000af7565b90620009f1565b6200081e565b62015953620159356201592f620158ef620158e985620158e2600362000b1c565b90620009f1565b6200080f565b62015928620159146201590e87620159075f620007a4565b90620009f1565b6200080f565b6201592161114e62004998565b9062003081565b9062001fa2565b6200143f565b6201594d8862015946600362000b1c565b90620009f1565b6200081e565b620159e5620159c7620159c1620159816201597b8562015974600462000b41565b90620009f1565b6200080f565b620159ba620159a6620159a087620159995f620007a4565b90620009f1565b6200080f565b620159b361131f620049bd565b9062003081565b9062001fa2565b6200143f565b620159df88620159d8600462000b41565b90620009f1565b6200081e565b62015a7762015a5962015a5362015a1362015a0d8562015a06600562000b66565b90620009f1565b6200080f565b62015a4c62015a3862015a328762015a2b5f620007a4565b90620009f1565b6200080f565b62015a4561167b620049e2565b9062003081565b9062001fa2565b6200143f565b62015a718862015a6a600562000b66565b90620009f1565b6200081e565b62015b0962015aeb62015ae562015aa562015a9f8562015a98600662000b8b565b90620009f1565b6200080f565b62015ade62015aca62015ac48762015abd5f620007a4565b90620009f1565b6200080f565b62015ad761137162004a07565b9062003081565b9062001fa2565b6200143f565b62015b038862015afc600662000b8b565b90620009f1565b6200081e565b62015b9b62015b7d62015b7762015b3762015b318562015b2a600762000bb0565b90620009f1565b6200080f565b62015b7062015b5c62015b568762015b4f5f620007a4565b90620009f1565b6200080f565b62015b6961123062004a2c565b9062003081565b9062001fa2565b6200143f565b62015b958862015b8e600762000bb0565b90620009f1565b6200081e565b62015c2d62015c0f62015c0962015bc962015bc38562015bbc60086200145b565b90620009f1565b6200080f565b62015c0262015bee62015be88762015be15f620007a4565b90620009f1565b6200080f565b62015bfb61182c62004a51565b9062003081565b9062001fa2565b6200143f565b62015c278862015c2060086200145b565b90620009f1565b6200081e565b62015cbf62015ca162015c9b62015c5b62015c558562015c4e600962002102565b90620009f1565b6200080f565b62015c9462015c8062015c7a8762015c735f620007a4565b90620009f1565b6200080f565b62015c8d61136862004a76565b9062003081565b9062001fa2565b6200143f565b62015cb98862015cb2600962002102565b90620009f1565b6200081e565b62015d5162015d3362015d2d62015ced62015ce78562015ce0600a6200214c565b90620009f1565b6200080f565b62015d2662015d1262015d0c8762015d055f620007a4565b90620009f1565b6200080f565b62015d1f610f3162004a9b565b9062003081565b9062001fa2565b6200143f565b62015d4b8862015d44600a6200214c565b90620009f1565b6200081e565b62015dac62015d9862015d9262015d7f62015d798562015d72600b62002196565b90620009f1565b6200080f565b9362015d8b5f620007a4565b90620009f1565b6200080f565b62015da56115c962004ac0565b9062003081565b9062001fa2565b6200143f565b62015dd18462015dca600b62002196565b90620009f1565b6200081e565b5b62015402565b62015df562015dee601562002430565b916200012b565b1462015e04575b505062015dd8565b6201678381620161786201615b62016155620167a7966201614e62016107620160c0620160796201603262015feb8a62015fe462015fca620167899f62015fb062015f6962015f2262015edb62015e9462015e7a62015e7462015fc49762015e6d600162000ad2565b90620009f1565b6200080f565b62015e8d673abeb80def61cc856200458c565b9062003081565b62015ed462015eba62015eb48b62015ead600262000af7565b90620009f1565b6200080f565b62015ecd679d19c9dd4eac4133620045b1565b9062003081565b9062001fa2565b62015f1b62015f0162015efb8a62015ef4600362000b1c565b90620009f1565b6200080f565b62015f1467075a652d9641a985620045d6565b9062003081565b9062001fa2565b62015f6262015f4862015f428962015f3b600462000b41565b90620009f1565b6200080f565b62015f5b679daf69ae1b67e667620045fb565b9062003081565b9062001fa2565b62015fa962015f8f62015f898862015f82600562000b66565b90620009f1565b6200080f565b62015fa267364f71da77920a1862004620565b9062003081565b9062001fa2565b9362015fbd600662000b8b565b90620009f1565b6200080f565b62015fdd6750bd769f745c95b162004645565b9062003081565b9062001fa2565b6201602b620160116201600b8d62016004600762000bb0565b90620009f1565b6200080f565b6201602467f223d1180dbbf3fc6200466a565b9062003081565b9062001fa2565b6201607262016058620160528c6201604b60086200145b565b90620009f1565b6200080f565b6201606b672f885e584e04aa996200468f565b9062003081565b9062001fa2565b620160b96201609f620160998b62016092600962002102565b90620009f1565b6200080f565b620160b267b69a0fa70aea684a620046b4565b9062003081565b9062001fa2565b62016100620160e6620160e08a620160d9600a6200214c565b90620009f1565b6200080f565b620160f96709584acaa6e062a0620046d9565b9062003081565b9062001fa2565b620161476201612d620161278962016120600b62002196565b90620009f1565b6200080f565b62016140670bc051640145b19b620046fe565b9062003081565b9062001fa2565b9062001fa2565b6200143f565b62016172886201616b5f620007a4565b90620009f1565b6200081e565b62016209620161eb620161e5620161a6620161a08562016199600162000ad2565b90620009f1565b6200080f565b620161de620161cb620161c587620161be5f620007a4565b90620009f1565b6200080f565b620161d760146200240b565b9062003081565b9062001fa2565b6200143f565b6201620388620161fc600162000ad2565b90620009f1565b6200081e565b6201629a6201627c620162766201623762016231856201622a600262000af7565b90620009f1565b6200080f565b6201626f6201625c62016256876201624f5f620007a4565b90620009f1565b6200080f565b62016268602262004723565b9062003081565b9062001fa2565b6200143f565b62016294886201628d600262000af7565b90620009f1565b6200081e565b6201632b6201630d62016307620162c8620162c285620162bb600362000b1c565b90620009f1565b6200080f565b62016300620162ed620162e787620162e05f620007a4565b90620009f1565b6200080f565b620162f9601262002377565b9062003081565b9062001fa2565b6200143f565b62016325886201631e600362000b1c565b90620009f1565b6200081e565b620163bc6201639e620163986201635962016353856201634c600462000b41565b90620009f1565b6200080f565b620163916201637e6201637887620163715f620007a4565b90620009f1565b6200080f565b6201638a602762004748565b9062003081565b9062001fa2565b6200143f565b620163b688620163af600462000b41565b90620009f1565b6200081e565b6201644d6201642f62016429620163ea620163e485620163dd600562000b66565b90620009f1565b6200080f565b620164226201640f6201640987620164025f620007a4565b90620009f1565b6200080f565b6201641b600d62002205565b9062003081565b9062001fa2565b6200143f565b620164478862016440600562000b66565b90620009f1565b6200081e565b620164de620164c0620164ba6201647b62016475856201646e600662000b8b565b90620009f1565b6200080f565b620164b3620164a06201649a87620164935f620007a4565b90620009f1565b6200080f565b620164ac600d62002205565b9062003081565b9062001fa2565b6200143f565b620164d888620164d1600662000b8b565b90620009f1565b6200081e565b6201656f620165516201654b6201650c6201650685620164ff600762000bb0565b90620009f1565b6200080f565b62016544620165316201652b87620165245f620007a4565b90620009f1565b6200080f565b6201653d601c6200476d565b9062003081565b9062001fa2565b6200143f565b620165698862016562600762000bb0565b90620009f1565b6200081e565b62016600620165e2620165dc6201659d62016597856201659060086200145b565b90620009f1565b6200080f565b620165d5620165c2620165bc87620165b55f620007a4565b90620009f1565b6200080f565b620165ce600262000af7565b9062003081565b9062001fa2565b6200143f565b620165fa88620165f360086200145b565b90620009f1565b6200081e565b62016691620166736201666d6201662e620166288562016621600962002102565b90620009f1565b6200080f565b62016666620166536201664d87620166465f620007a4565b90620009f1565b6200080f565b6201665f6010620022e3565b9062003081565b9062001fa2565b6200143f565b6201668b8862016684600962002102565b90620009f1565b6200081e565b6201672262016704620166fe620166bf620166b985620166b2600a6200214c565b90620009f1565b6200080f565b620166f7620166e4620166de87620166d75f620007a4565b90620009f1565b6200080f565b620166f0602962004792565b9062003081565b9062001fa2565b6200143f565b6201671c8862016715600a6200214c565b90620009f1565b6200081e565b6201677c6201676962016763620167506201674a8562016743600b62002196565b90620009f1565b6200080f565b936201675c5f620007a4565b90620009f1565b6200080f565b62016775600f62002299565b9062003081565b9062001fa2565b6200143f565b620167a1846201679a600b62002196565b90620009f1565b6200081e565b5f8062015dfc565b90565b620167cb620167c5620167d192620167af565b620007a1565b6200012b565b90565b90565b620167f0620167ea620167f692620167d4565b620007a1565b6200012b565b90565b90565b620168156201680f6201681b92620167f9565b620007a1565b6200012b565b90565b90565b6201683a6201683462016840926201681e565b620007a1565b6200012b565b90565b90565b6201685f62016859620168659262016843565b620007a1565b6200012b565b90565b90565b620168846201687e6201688a9262016868565b620007a1565b6200012b565b90565b90565b620168a9620168a3620168af926201688d565b620007a1565b6200012b565b90565b90565b620168ce620168c8620168d492620168b2565b620007a1565b6200012b565b90565b90565b620168f3620168ed620168f992620168d7565b620007a1565b6200012b565b90565b90565b62016918620169126201691e92620168fc565b620007a1565b6200012b565b90565b90565b6201693d62016937620169439262016921565b620007a1565b6200012b565b90565b90565b620169626201695c620169689262016946565b620007a1565b6200012b565b90565b90565b62016987620169816201698d926201696b565b620007a1565b6200012b565b90565b90565b620169ac620169a6620169b29262016990565b620007a1565b6200012b565b90565b90565b620169d1620169cb620169d792620169b5565b620007a1565b6200012b565b90565b90565b620169f6620169f0620169fc92620169da565b620007a1565b6200012b565b90565b90565b62016a1b62016a1562016a2192620169ff565b620007a1565b6200012b565b90565b90565b62016a4062016a3a62016a469262016a24565b620007a1565b6200012b565b90565b90565b62016a6562016a5f62016a6b9262016a49565b620007a1565b6200012b565b90565b90565b62016a8a62016a8462016a909262016a6e565b620007a1565b6200012b565b90565b90565b62016aaf62016aa962016ab59262016a93565b620007a1565b6200012b565b90565b90565b62016ad462016ace62016ada9262016ab8565b620007a1565b6200012b565b90565b90565b62016af962016af362016aff9262016add565b620007a1565b6200012b565b90565b90565b62016b1e62016b1862016b249262016b02565b620007a1565b6200012b565b90565b90565b62016b4362016b3d62016b499262016b27565b620007a1565b6200012b565b90565b90565b62016b6862016b6262016b6e9262016b4c565b620007a1565b6200012b565b90565b90565b62016b8d62016b8762016b939262016b71565b620007a1565b6200012b565b90565b90565b62016bb262016bac62016bb89262016b96565b620007a1565b6200012b565b90565b90565b62016bd762016bd162016bdd9262016bbb565b620007a1565b6200012b565b90565b90565b62016bfc62016bf662016c029262016be0565b620007a1565b6200012b565b90565b90565b62016c2162016c1b62016c279262016c05565b620007a1565b6200012b565b90565b90565b62016c4662016c4062016c4c9262016c2a565b620007a1565b6200012b565b90565b90565b62016c6b62016c6562016c719262016c4f565b620007a1565b6200012b565b90565b90565b62016c9062016c8a62016c969262016c74565b620007a1565b6200012b565b90565b90565b62016cb562016caf62016cbb9262016c99565b620007a1565b6200012b565b90565b90565b62016cda62016cd462016ce09262016cbe565b620007a1565b6200012b565b90565b90565b62016cff62016cf962016d059262016ce3565b620007a1565b6200012b565b90565b90565b62016d2462016d1e62016d2a9262016d08565b620007a1565b6200012b565b90565b90565b62016d4962016d4362016d4f9262016d2d565b620007a1565b6200012b565b90565b90565b62016d6e62016d6862016d749262016d52565b620007a1565b6200012b565b90565b90565b62016d9362016d8d62016d999262016d77565b620007a1565b6200012b565b90565b90565b62016db862016db262016dbe9262016d9c565b620007a1565b6200012b565b90565b90565b62016ddd62016dd762016de39262016dc1565b620007a1565b6200012b565b90565b90565b62016e0262016dfc62016e089262016de6565b620007a1565b6200012b565b90565b90565b62016e2762016e2162016e2d9262016e0b565b620007a1565b6200012b565b90565b90565b62016e4c62016e4662016e529262016e30565b620007a1565b6200012b565b90565b90565b62016e7162016e6b62016e779262016e55565b620007a1565b6200012b565b90565b90565b62016e9662016e9062016e9c9262016e7a565b620007a1565b6200012b565b90565b90565b62016ebb62016eb562016ec19262016e9f565b620007a1565b6200012b565b90565b90565b62016ee062016eda62016ee69262016ec4565b620007a1565b6200012b565b90565b90565b62016f0562016eff62016f0b9262016ee9565b620007a1565b6200012b565b90565b90565b62016f2a62016f2462016f309262016f0e565b620007a1565b6200012b565b90565b90565b62016f4f62016f4962016f559262016f33565b620007a1565b6200012b565b90565b90565b62016f7462016f6e62016f7a9262016f58565b620007a1565b6200012b565b90565b90565b62016f9962016f9362016f9f9262016f7d565b620007a1565b6200012b565b90565b90565b62016fbe62016fb862016fc49262016fa2565b620007a1565b6200012b565b90565b90565b62016fe362016fdd62016fe99262016fc7565b620007a1565b6200012b565b90565b90565b62017008620170026201700e9262016fec565b620007a1565b6200012b565b90565b90565b6201702d62017027620170339262017011565b620007a1565b6200012b565b90565b90565b620170526201704c620170589262017036565b620007a1565b6200012b565b90565b90565b62017077620170716201707d926201705b565b620007a1565b6200012b565b90565b90565b6201709c62017096620170a29262017080565b620007a1565b6200012b565b90565b90565b620170c1620170bb620170c792620170a5565b620007a1565b6200012b565b90565b90565b620170e6620170e0620170ec92620170ca565b620007a1565b6200012b565b90565b90565b6201710b620171056201711192620170ef565b620007a1565b6200012b565b90565b90565b620171306201712a620171369262017114565b620007a1565b6200012b565b90565b90565b620171556201714f6201715b9262017139565b620007a1565b6200012b565b90565b90565b6201717a6201717462017180926201715e565b620007a1565b6200012b565b90565b90565b6201719f62017199620171a59262017183565b620007a1565b6200012b565b90565b90565b620171c4620171be620171ca92620171a8565b620007a1565b6200012b565b90565b90565b620171e9620171e3620171ef92620171cd565b620007a1565b6200012b565b90565b90565b6201720e620172086201721492620171f2565b620007a1565b6200012b565b90565b90565b620172336201722d620172399262017217565b620007a1565b6200012b565b90565b90565b62017258620172526201725e926201723c565b620007a1565b6200012b565b90565b90565b6201727d62017277620172839262017261565b620007a1565b6200012b565b90565b90565b620172a26201729c620172a89262017286565b620007a1565b6200012b565b90565b90565b620172c7620172c1620172cd92620172ab565b620007a1565b6200012b565b90565b90565b620172ec620172e6620172f292620172d0565b620007a1565b6200012b565b90565b90565b620173116201730b6201731792620172f5565b620007a1565b6200012b565b90565b620173336201732d62017339926200107c565b620007a1565b6200012b565b90565b90565b62017358620173526201735e926201733c565b620007a1565b6200012b565b90565b90565b6201737d62017377620173839262017361565b620007a1565b6200012b565b90565b90565b620173a26201739c620173a89262017386565b620007a1565b6200012b565b90565b90565b620173c7620173c1620173cd92620173ab565b620007a1565b6200012b565b90565b90565b620173ec620173e6620173f292620173d0565b620007a1565b6200012b565b90565b90565b620174116201740b6201741792620173f5565b620007a1565b6200012b565b90565b90565b62017436620174306201743c926201741a565b620007a1565b6200012b565b90565b90565b6201745b6201745562017461926201743f565b620007a1565b6200012b565b90565b90565b620174806201747a620174869262017464565b620007a1565b6200012b565b90565b90565b620174a56201749f620174ab9262017489565b620007a1565b6200012b565b90565b90565b620174ca620174c4620174d092620174ae565b620007a1565b6200012b565b90565b90565b620174ef620174e9620174f592620174d3565b620007a1565b6200012b565b90565b90565b620175146201750e6201751a92620174f8565b620007a1565b6200012b565b90565b90565b62017539620175336201753f926201751d565b620007a1565b6200012b565b90565b90565b6201755e62017558620175649262017542565b620007a1565b6200012b565b90565b90565b620175836201757d620175899262017567565b620007a1565b6200012b565b90565b90565b620175a8620175a2620175ae926201758c565b620007a1565b6200012b565b90565b90565b620175cd620175c7620175d392620175b1565b620007a1565b6200012b565b90565b90565b620175f2620175ec620175f892620175d6565b620007a1565b6200012b565b90565b90565b62017617620176116201761d92620175fb565b620007a1565b6200012b565b90565b90565b6201763c62017636620176429262017620565b620007a1565b6200012b565b90565b90565b620176616201765b620176679262017645565b620007a1565b6200012b565b90565b90565b62017686620176806201768c926201766a565b620007a1565b6200012b565b90565b90565b620176ab620176a5620176b1926201768f565b620007a1565b6200012b565b90565b90565b620176d0620176ca620176d692620176b4565b620007a1565b6200012b565b90565b90565b620176f5620176ef620176fb92620176d9565b620007a1565b6200012b565b90565b90565b6201771a620177146201772092620176fe565b620007a1565b6200012b565b90565b90565b6201773f62017739620177459262017723565b620007a1565b6200012b565b90565b90565b620177646201775e6201776a9262017748565b620007a1565b6200012b565b90565b90565b62017789620177836201778f926201776d565b620007a1565b6200012b565b90565b90565b620177ae620177a8620177b49262017792565b620007a1565b6200012b565b90565b90565b620177d3620177cd620177d992620177b7565b620007a1565b6200012b565b90565b90565b620177f8620177f2620177fe92620177dc565b620007a1565b6200012b565b90565b90565b6201781d62017817620178239262017801565b620007a1565b6200012b565b90565b90565b620178426201783c620178489262017826565b620007a1565b6200012b565b90565b90565b62017867620178616201786d926201784b565b620007a1565b6200012b565b90565b90565b6201788c62017886620178929262017870565b620007a1565b6200012b565b90565b90565b620178b1620178ab620178b79262017895565b620007a1565b6200012b565b90565b90565b620178d6620178d0620178dc92620178ba565b620007a1565b6200012b565b90565b90565b620178fb620178f56201790192620178df565b620007a1565b6200012b565b90565b90565b620179206201791a620179269262017904565b620007a1565b6200012b565b90565b90565b620179456201793f6201794b9262017929565b620007a1565b6200012b565b90565b90565b6201796a6201796462017970926201794e565b620007a1565b6200012b565b90565b90565b6201798f62017989620179959262017973565b620007a1565b6200012b565b90565b90565b620179b4620179ae620179ba9262017998565b620007a1565b6200012b565b90565b90565b620179d9620179d3620179df92620179bd565b620007a1565b6200012b565b90565b90565b620179fe620179f862017a0492620179e2565b620007a1565b6200012b565b90565b90565b62017a2362017a1d62017a299262017a07565b620007a1565b6200012b565b90565b90565b62017a4862017a4262017a4e9262017a2c565b620007a1565b6200012b565b90565b90565b62017a6d62017a6762017a739262017a51565b620007a1565b6200012b565b90565b90565b62017a9262017a8c62017a989262017a76565b620007a1565b6200012b565b90565b62017aa56200143b565b508062017abe62017ab76030620167b2565b916200012b565b106201844e578062017adc62017ad560486201705e565b916200012b565b1062017f98578062017afa62017af360546201758f565b916200012b565b1062017d4c578062017b1862017b11605862017829565b916200012b565b1062017c9c578062017b3662017b2f605c62017907565b916200012b565b1062017bec578062017b5462017b4d605e620179e5565b916200012b565b1062017ba35762017b7162017b6a605e620179e5565b916200012b565b1462017b8d5762017b8a67bc8dfb627fe558fc62017a79565b90565b62017ba067dfd1c4febcc8123862017a54565b90565b62017bba62017bb3605c62017907565b916200012b565b1462017bd65762017bd367f172d73e004fc90d62017a2f565b90565b62017be9674543d9df5476d3cb62017a0a565b90565b8062017c0462017bfd605a6201792c565b916200012b565b1062017c535762017c2162017c1a605a6201792c565b916200012b565b1462017c3d5762017c3a677f795e2a5f915440620179c0565b90565b62017c506770fc91eb5937085c6201799b565b90565b62017c6a62017c63605862017829565b916200012b565b1462017c865762017c83678ffd96bbf9c9c81d62017976565b90565b62017c9967aaed34074b16434662017951565b90565b8062017cb462017cad60566201784e565b916200012b565b1062017d035762017cd162017cca60566201784e565b916200012b565b1462017ced5762017cea6725032aa7c4cb1811620178e2565b90565b62017d0067551f5b0fbe7b1840620178bd565b90565b62017d1a62017d1360546201758f565b916200012b565b1462017d365762017d336740b9e922ed9771e262017898565b90565b62017d4967f3c12fe54d5c653b62017873565b90565b8062017d6462017d5d604c620175b4565b916200012b565b1062017ee8578062017d8262017d7b605062017692565b916200012b565b1062017e38578062017da062017d99605262017770565b916200012b565b1062017def5762017dbd62017db6605262017770565b916200012b565b1462017dd95762017dd667164bb2de1bbeddc862017804565b90565b62017dec676e7c8f1a84265034620177df565b90565b62017e0662017dff605062017692565b916200012b565b1462017e225762017e1f67bb9879c6e61fd62a620177ba565b90565b62017e356780cefd2b7d99ff8362017795565b90565b8062017e5062017e49604e620176b7565b916200012b565b1062017e9f5762017e6d62017e66604e620176b7565b916200012b565b1462017e895762017e86672a440b51a4945ee56201774b565b90565b62017e9c67e8928a4c8911475062017726565b90565b62017eb662017eaf604c620175b4565b916200012b565b1462017ed25762017ecf673762b59a8ae83efa62017701565b90565b62017ee567166fc0cc438a3c72620176dc565b90565b8062017f0062017ef9604a620175d9565b916200012b565b1062017f4f5762017f1d62017f16604a620175d9565b916200012b565b1462017f395762017f36674135c2bebfe148c66201766d565b90565b62017f4c670230bca8f4df054462017648565b90565b62017f6662017f5f60486201705e565b916200012b565b1462017f825762017f7f67a3c95eaec244aa3062017623565b90565b62017f95677be5b9ffda905e1c620175fe565b90565b8062017fb062017fa9603c62017083565b916200012b565b1062018202578062017fce62017fc760406201731a565b916200012b565b1062018152578062017fec62017fe56044620173f8565b916200012b565b10620180a257806201800a620180036046620174d6565b916200012b565b10620180595762018027620180206046620174d6565b916200012b565b1462018043576201804067a49229d24d0143436201756a565b90565b6201805667debe0853b1a1d37862017545565b90565b62018070620180696044620173f8565b916200012b565b146201808c576201808967dae8645996edd6a562017520565b90565b6201809f679e9cbd781628fc5b620174fb565b90565b80620180ba620180b360426201741d565b916200012b565b106201810957620180d7620180d060426201741d565b916200012b565b14620180f357620180f067ba0dec7103255dd4620174b1565b90565b62018106678e7fba02983224bd6201748c565b90565b620181206201811960406201731a565b916200012b565b146201813c5762018139673996ce73dd832c1c62017467565b90565b6201814f671c86d265f15750cd62017442565b90565b806201816a62018163603e6201733f565b916200012b565b10620181b9576201818762018180603e6201733f565b916200012b565b14620181a357620181a067d1f507e43ab749b2620173d3565b90565b620181b66742747a7245e7fa84620173ae565b90565b620181d0620181c9603c62017083565b916200012b565b14620181ec57620181e9676f90ff3b6a65f10862017389565b90565b620181ff67e70201e960cb78b862017364565b90565b806201821a620182136034620170a8565b916200012b565b106201839e57806201823862018231603862017186565b916200012b565b10620182ee5780620182566201824f603a62017264565b916200012b565b10620182a557620182736201826c603a62017264565b916200012b565b146201828f576201828c67c02ac5e47312d3ca620172f8565b90565b620182a267f6f31db1899b12d5620172d3565b90565b620182bc620182b5603862017186565b916200012b565b14620182d857620182d567594e29c42473d06b620172ae565b90565b620182eb671ff7327017aa2a6e62017289565b90565b8062018306620182ff6036620171ab565b916200012b565b106201835557620183236201831c6036620171ab565b916200012b565b146201833f576201833c67e587cc0d69a733466201723f565b90565b620183526711b4c74eda62beef6201721a565b90565b6201836c620183656034620170a8565b916200012b565b146201838857620183856701a8d1a588085737620171f5565b90565b6201839b67e4a3dff1d7d6fef9620171d0565b90565b80620183b6620183af6032620170cd565b916200012b565b106201840557620183d3620183cc6032620170cd565b916200012b565b14620183ef57620183ec6715147108121967d762017161565b90565b6201840267023e7414af6630686201713c565b90565b6201841c620184156030620167b2565b916200012b565b146201843857620184356718a42105c31b7e8862017117565b90565b6201844b67475cd3205a3bdcde620170f2565b90565b80620184666201845f6018620167d7565b916200012b565b10620189225780620184846201847d602462016b99565b916200012b565b10620186d65780620184a26201849b602862016de9565b916200012b565b10620186265780620184c0620184b9602c62016ec7565b916200012b565b10620185765780620184de620184d7602e62016fa5565b916200012b565b106201852d57620184fb620184f4602e62016fa5565b916200012b565b1462018517576201851467a637093ecb2ad63162017039565b90565b6201852a6740b12cbf09ef74bf62017014565b90565b620185446201853d602c62016ec7565b916200012b565b1462018560576201855d67f5c0bc1de6da869962016fef565b90565b6201857367e8be3e5da8043e5762016fca565b90565b806201858e62018587602a62016eec565b916200012b565b10620185dd57620185ab620185a4602a62016eec565b916200012b565b14620185c757620185c4674b959b48b9c1073362016f80565b90565b620185da6785dc499b11d77b7562016f5b565b90565b620185f4620185ed602862016de9565b916200012b565b1462018610576201860d6761df517b86a4643962016f36565b90565b62018623677cf416cfe7e14ca162016f11565b90565b806201863e62018637602662016e0e565b916200012b565b106201868d576201865b62018654602662016e0e565b916200012b565b14620186775762018674676f6f2ed77246635262016ea2565b90565b6201868a67019d5ee2af82ec1c62016e7d565b90565b620186a46201869d602462016b99565b916200012b565b14620186c057620186bd6770e741ebfee9658662016e58565b90565b620186d36792a756e67e2b941362016e33565b90565b80620186ee620186e7601c6200476d565b916200012b565b106201887257806201870c62018705602062016c77565b916200012b565b10620187c257806201872a62018723602262004723565b916200012b565b1062018779576201874762018740602262004723565b916200012b565b146201876357620187606728b8ec3ae9c2963362016dc4565b90565b6201877667f061274fdb150d6162016d9f565b90565b6201879062018789602062016c77565b916200012b565b14620187ac57620187a9670da2abef589d644e62016d7a565b90565b620187bf6766bbd30e6ce0e58362016d55565b90565b80620187da620187d3601e62016c9c565b916200012b565b106201882957620187f7620187f0601e62016c9c565b916200012b565b146201881357620188106757a1550c110b304162016d30565b90565b6201882667502ad7dc18c2ad8762016d0b565b90565b6201884062018839601c6200476d565b916200012b565b146201885c576201885967ac51d9f5fcf8535e62016ce6565b90565b6201886f67ad6805e0e30b2c8a62016cc1565b90565b806201888a62018883601a62016bbe565b916200012b565b10620188d957620188a7620188a0601a62016bbe565b916200012b565b14620188c357620188c0678297132b32825daf62016c52565b90565b620188d667f7d713e80639116562016c2d565b90565b620188f0620188e96018620167d7565b916200012b565b146201890c576201890967f56f6d48959a600d62016c08565b90565b6201891f67e9fa634a21de008262016be3565b90565b806201893a62018933600c620009c0565b916200012b565b1062018b8c578062018958620189516010620022e3565b916200012b565b1062018adc5780620189766201896f60146200240b565b916200012b565b1062018a2c5780620189946201898d601662016ae0565b916200012b565b10620189e357620189b1620189aa601662016ae0565b916200012b565b14620189cd57620189ca67c0572c8c08cbbbad62016b74565b90565b620189e067849c2656969c4be762016b4f565b90565b620189fa620189f360146200240b565b916200012b565b1462018a165762018a1367c35eae071903ff0b62016b2a565b90565b62018a2967d5c7718fbfc647fb62016b05565b90565b8062018a4462018a3d601262002377565b916200012b565b1062018a935762018a6162018a5a601262002377565b916200012b565b1462018a7d5762018a7a67e360c6e0ae486f3862016abb565b90565b62018a90679a2eedb70d8f8cfa62016a96565b90565b62018aaa62018aa36010620022e3565b916200012b565b1462018ac65762018ac3676e4a40c18f30c09c62016a71565b90565b62018ad967308bbd23dc5416cc62016a4c565b90565b8062018af462018aed600e6200224f565b916200012b565b1062018b435762018b1162018b0a600e6200224f565b916200012b565b1462018b2d5762018b2a67a484c4c5ef6a078162016a27565b90565b62018b4067c3b919ad495dc57462016a02565b90565b62018b5a62018b53600c620009c0565b916200012b565b1462018b765762018b736759cd1a8a41c18e55620169dd565b90565b62018b896786287821f722c881620169b8565b90565b8062018ba462018b9d600462000b41565b916200012b565b1062018d28578062018bc262018bbb60086200145b565b916200012b565b1062018c78578062018be062018bd9600a6200214c565b916200012b565b1062018c2f5762018bfd62018bf6600a6200214c565b916200012b565b1462018c195762018c1667c54302f225db2c7662016993565b90565b62018c2c67e17653a29da578a16201696e565b90565b62018c4662018c3f60086200145b565b916200012b565b1462018c625762018c5f67277fa208bf337bff62016949565b90565b62018c7567cdc4a239b0c4442662016924565b90565b8062018c9062018c89600662000b8b565b916200012b565b1062018cdf5762018cad62018ca6600662000b8b565b916200012b565b1462018cc95762018cc667ef781c7ce35b4c3d620168ff565b90565b62018cdc67d438539c95f63e9f620168da565b90565b62018cf662018cef600462000b41565b916200012b565b1462018d125762018d0f678cae14cb07d09bf1620168b5565b90565b62018d2567e10d666650f4e01262016890565b90565b8062018d4062018d39600262000af7565b916200012b565b1062018d8f5762018d5d62018d56600262000af7565b916200012b565b1462018d795762018d76670f6760a4803427d76201686b565b90565b62018d8c67b2fb0d31cee799b462016846565b90565b62018da562018d9e5f620007a4565b916200012b565b1462018dc15762018dbe677746a55f43921ad762016821565b90565b62018dd467b585f766f2144405620167fc565b90565b62018de3600f62002299565b90565b62018df2602962004792565b90565b62018e0e62018e0862018e149262000b3e565b620007a1565b6200107f565b90565b1b90565b62018e419062018e3a62018e3362018e47946200107f565b916200012b565b9062018e17565b6200012b565b90565b62018e56600262000af7565b90565b62018e65601c6200476d565b90565b62018e74600d62002205565b90565b62018e83600d62002205565b90565b62018e92602762004748565b90565b62018ea1601262002377565b90565b62018eb0602262004723565b90565b62018ebf60146200240b565b90565b620191f29062018ed16200143b565b50620191eb620191c5620191bf620191ac6201916e62019130620190f2620190b462019076620190388a62019031620190206201901a6201900662018fc862018f8862018f4a62018f3962018f338962018f2c600162000ad2565b90620009f1565b6200080f565b62018f4362018dd7565b9062003081565b62018f8162018f7062018f6a8a62018f63600262000af7565b90620009f1565b6200080f565b62018f7a62018de6565b9062003081565b9062001fa2565b62018fc162018fae62018fa88962018fa1600362000b1c565b90620009f1565b6200080f565b62018fba600462018df5565b9062018e1b565b9062001fa2565b62018fff62018fee62018fe88862018fe1600462000b41565b90620009f1565b6200080f565b62018ff862018e4a565b9062003081565b9062001fa2565b9362019013600562000b66565b90620009f1565b6200080f565b6201902a62018e59565b9062003081565b9062001fa2565b6201906f6201905e620190588d62019051600662000b8b565b90620009f1565b6200080f565b6201906862018e68565b9062003081565b9062001fa2565b620190ad6201909c620190968c6201908f600762000bb0565b90620009f1565b6200080f565b620190a662018e77565b9062003081565b9062001fa2565b620190eb620190da620190d48b620190cd60086200145b565b90620009f1565b6200080f565b620190e462018e86565b9062003081565b9062001fa2565b6201912962019118620191128a6201910b600962002102565b90620009f1565b6200080f565b6201912262018e95565b9062003081565b9062001fa2565b6201916762019156620191508962019149600a6200214c565b90620009f1565b6200080f565b6201916062018ea4565b9062003081565b9062001fa2565b620191a5620191946201918e8862019187600b62002196565b90620009f1565b6200080f565b6201919e62018eb3565b9062003081565b9062001fa2565b93620191b85f620007a4565b90620009f1565b6200080f565b620191e4620191d36200456b565b620191dd6200457a565b9062001fa2565b9062003081565b9062001fa2565b90565b6201951090620192046200143b565b5062019509620194f8620194f2620194df620194a16201946362019425620193e7620193a96201936b8a62019364620193536201934d62019339620192f9620192bb6201927d6201926c62019266896201925f600162000ad2565b90620009f1565b6200080f565b620192766200456b565b9062003081565b620192b4620192a36201929d8a62019296600262000af7565b90620009f1565b6200080f565b620192ad62018dd7565b9062003081565b9062001fa2565b620192f2620192e1620192db89620192d4600362000b1c565b90620009f1565b6200080f565b620192eb62018de6565b9062003081565b9062001fa2565b620193326201931f620193198862019312600462000b41565b90620009f1565b6200080f565b6201932b600462018df5565b9062018e1b565b9062001fa2565b9362019346600562000b66565b90620009f1565b6200080f565b6201935d62018e4a565b9062003081565b9062001fa2565b620193a2620193916201938b8d62019384600662000b8b565b90620009f1565b6200080f565b6201939b62018e59565b9062003081565b9062001fa2565b620193e0620193cf620193c98c620193c2600762000bb0565b90620009f1565b6200080f565b620193d962018e68565b9062003081565b9062001fa2565b6201941e6201940d620194078b6201940060086200145b565b90620009f1565b6200080f565b6201941762018e77565b9062003081565b9062001fa2565b6201945c6201944b620194458a6201943e600962002102565b90620009f1565b6200080f565b6201945562018e86565b9062003081565b9062001fa2565b6201949a6201948962019483896201947c600a6200214c565b90620009f1565b6200080f565b6201949362018e95565b9062003081565b9062001fa2565b620194d8620194c7620194c188620194ba600b62002196565b90620009f1565b6200080f565b620194d162018ea4565b9062003081565b9062001fa2565b93620194eb5f620007a4565b90620009f1565b6200080f565b6201950262018eb3565b9062003081565b9062001fa2565b90565b6201982e90620195226200143b565b50620198276201981662019810620197fc620197bf620197816201974362019705620196c7620196898a62019682620196716201966b6201965762019617620195d96201959b6201958a62019584896201957d600262000af7565b90620009f1565b6200080f565b620195946200456b565b9062003081565b620195d2620195c1620195bb8a620195b4600362000b1c565b90620009f1565b6200080f565b620195cb62018dd7565b9062003081565b9062001fa2565b62019610620195ff620195f989620195f2600462000b41565b90620009f1565b6200080f565b6201960962018de6565b9062003081565b9062001fa2565b620196506201963d620196378862019630600562000b66565b90620009f1565b6200080f565b62019649600462018df5565b9062018e1b565b9062001fa2565b9362019664600662000b8b565b90620009f1565b6200080f565b6201967b62018e4a565b9062003081565b9062001fa2565b620196c0620196af620196a98d620196a2600762000bb0565b90620009f1565b6200080f565b620196b962018e59565b9062003081565b9062001fa2565b620196fe620196ed620196e78c620196e060086200145b565b90620009f1565b6200080f565b620196f762018e68565b9062003081565b9062001fa2565b6201973c6201972b620197258b6201971e600962002102565b90620009f1565b6200080f565b6201973562018e77565b9062003081565b9062001fa2565b6201977a62019769620197638a6201975c600a6200214c565b90620009f1565b6200080f565b6201977362018e86565b9062003081565b9062001fa2565b620197b8620197a7620197a1896201979a600b62002196565b90620009f1565b6200080f565b620197b162018e95565b9062003081565b9062001fa2565b620197f5620197e4620197de88620197d75f620007a4565b90620009f1565b6200080f565b620197ee62018ea4565b9062003081565b9062001fa2565b9362019809600162000ad2565b90620009f1565b6200080f565b6201982062018eb3565b9062003081565b9062001fa2565b90565b62019b4c90620198406200143b565b5062019b4562019b3462019b2e62019b1a62019adc62019a9f62019a6162019a23620199e5620199a78a620199a06201998f620199896201997562019935620198f7620198b9620198a8620198a2896201989b600362000b1c565b90620009f1565b6200080f565b620198b26200456b565b9062003081565b620198f0620198df620198d98a620198d2600462000b41565b90620009f1565b6200080f565b620198e962018dd7565b9062003081565b9062001fa2565b6201992e6201991d620199178962019910600562000b66565b90620009f1565b6200080f565b6201992762018de6565b9062003081565b9062001fa2565b6201996e6201995b62019955886201994e600662000b8b565b90620009f1565b6200080f565b62019967600462018df5565b9062018e1b565b9062001fa2565b9362019982600762000bb0565b90620009f1565b6200080f565b6201999962018e4a565b9062003081565b9062001fa2565b620199de620199cd620199c78d620199c060086200145b565b90620009f1565b6200080f565b620199d762018e59565b9062003081565b9062001fa2565b62019a1c62019a0b62019a058c620199fe600962002102565b90620009f1565b6200080f565b62019a1562018e68565b9062003081565b9062001fa2565b62019a5a62019a4962019a438b62019a3c600a6200214c565b90620009f1565b6200080f565b62019a5362018e77565b9062003081565b9062001fa2565b62019a9862019a8762019a818a62019a7a600b62002196565b90620009f1565b6200080f565b62019a9162018e86565b9062003081565b9062001fa2565b62019ad562019ac462019abe8962019ab75f620007a4565b90620009f1565b6200080f565b62019ace62018e95565b9062003081565b9062001fa2565b62019b1362019b0262019afc8862019af5600162000ad2565b90620009f1565b6200080f565b62019b0c62018ea4565b9062003081565b9062001fa2565b9362019b27600262000af7565b90620009f1565b6200080f565b62019b3e62018eb3565b9062003081565b9062001fa2565b90565b62019e6a9062019b5e6200143b565b5062019e6362019e5262019e4c62019e3862019dfa62019dbc62019d7f62019d4162019d0362019cc58a62019cbe62019cad62019ca762019c9362019c5362019c1562019bd762019bc662019bc08962019bb9600462000b41565b90620009f1565b6200080f565b62019bd06200456b565b9062003081565b62019c0e62019bfd62019bf78a62019bf0600562000b66565b90620009f1565b6200080f565b62019c0762018dd7565b9062003081565b9062001fa2565b62019c4c62019c3b62019c358962019c2e600662000b8b565b90620009f1565b6200080f565b62019c4562018de6565b9062003081565b9062001fa2565b62019c8c62019c7962019c738862019c6c600762000bb0565b90620009f1565b6200080f565b62019c85600462018df5565b9062018e1b565b9062001fa2565b9362019ca060086200145b565b90620009f1565b6200080f565b62019cb762018e4a565b9062003081565b9062001fa2565b62019cfc62019ceb62019ce58d62019cde600962002102565b90620009f1565b6200080f565b62019cf562018e59565b9062003081565b9062001fa2565b62019d3a62019d2962019d238c62019d1c600a6200214c565b90620009f1565b6200080f565b62019d3362018e68565b9062003081565b9062001fa2565b62019d7862019d6762019d618b62019d5a600b62002196565b90620009f1565b6200080f565b62019d7162018e77565b9062003081565b9062001fa2565b62019db562019da462019d9e8a62019d975f620007a4565b90620009f1565b6200080f565b62019dae62018e86565b9062003081565b9062001fa2565b62019df362019de262019ddc8962019dd5600162000ad2565b90620009f1565b6200080f565b62019dec62018e95565b9062003081565b9062001fa2565b62019e3162019e2062019e1a8862019e13600262000af7565b90620009f1565b6200080f565b62019e2a62018ea4565b9062003081565b9062001fa2565b9362019e45600362000b1c565b90620009f1565b6200080f565b62019e5c62018eb3565b9062003081565b9062001fa2565b90565b6201a1889062019e7c6200143b565b506201a1816201a1706201a16a6201a1566201a1186201a0da6201a09c6201a05f6201a02162019fe38a62019fdc62019fcb62019fc562019fb162019f7162019f3362019ef562019ee462019ede8962019ed7600562000b66565b90620009f1565b6200080f565b62019eee6200456b565b9062003081565b62019f2c62019f1b62019f158a62019f0e600662000b8b565b90620009f1565b6200080f565b62019f2562018dd7565b9062003081565b9062001fa2565b62019f6a62019f5962019f538962019f4c600762000bb0565b90620009f1565b6200080f565b62019f6362018de6565b9062003081565b9062001fa2565b62019faa62019f9762019f918862019f8a60086200145b565b90620009f1565b6200080f565b62019fa3600462018df5565b9062018e1b565b9062001fa2565b9362019fbe600962002102565b90620009f1565b6200080f565b62019fd562018e4a565b9062003081565b9062001fa2565b6201a01a6201a0096201a0038d62019ffc600a6200214c565b90620009f1565b6200080f565b6201a01362018e59565b9062003081565b9062001fa2565b6201a0586201a0476201a0418c6201a03a600b62002196565b90620009f1565b6200080f565b6201a05162018e68565b9062003081565b9062001fa2565b6201a0956201a0846201a07e8b6201a0775f620007a4565b90620009f1565b6200080f565b6201a08e62018e77565b9062003081565b9062001fa2565b6201a0d36201a0c26201a0bc8a6201a0b5600162000ad2565b90620009f1565b6200080f565b6201a0cc62018e86565b9062003081565b9062001fa2565b6201a1116201a1006201a0fa896201a0f3600262000af7565b90620009f1565b6200080f565b6201a10a62018e95565b9062003081565b9062001fa2565b6201a14f6201a13e6201a138886201a131600362000b1c565b90620009f1565b6200080f565b6201a14862018ea4565b9062003081565b9062001fa2565b936201a163600462000b41565b90620009f1565b6200080f565b6201a17a62018eb3565b9062003081565b9062001fa2565b90565b6201a4a6906201a19a6200143b565b506201a49f6201a48e6201a4886201a4746201a4366201a3f86201a3ba6201a37c6201a33f6201a3018a6201a2fa6201a2e96201a2e36201a2cf6201a28f6201a2516201a2136201a2026201a1fc896201a1f5600662000b8b565b90620009f1565b6200080f565b6201a20c6200456b565b9062003081565b6201a24a6201a2396201a2338a6201a22c600762000bb0565b90620009f1565b6200080f565b6201a24362018dd7565b9062003081565b9062001fa2565b6201a2886201a2776201a271896201a26a60086200145b565b90620009f1565b6200080f565b6201a28162018de6565b9062003081565b9062001fa2565b6201a2c86201a2b56201a2af886201a2a8600962002102565b90620009f1565b6200080f565b6201a2c1600462018df5565b9062018e1b565b9062001fa2565b936201a2dc600a6200214c565b90620009f1565b6200080f565b6201a2f362018e4a565b9062003081565b9062001fa2565b6201a3386201a3276201a3218d6201a31a600b62002196565b90620009f1565b6200080f565b6201a33162018e59565b9062003081565b9062001fa2565b6201a3756201a3646201a35e8c6201a3575f620007a4565b90620009f1565b6200080f565b6201a36e62018e68565b9062003081565b9062001fa2565b6201a3b36201a3a26201a39c8b6201a395600162000ad2565b90620009f1565b6200080f565b6201a3ac62018e77565b9062003081565b9062001fa2565b6201a3f16201a3e06201a3da8a6201a3d3600262000af7565b90620009f1565b6200080f565b6201a3ea62018e86565b9062003081565b9062001fa2565b6201a42f6201a41e6201a418896201a411600362000b1c565b90620009f1565b6200080f565b6201a42862018e95565b9062003081565b9062001fa2565b6201a46d6201a45c6201a456886201a44f600462000b41565b90620009f1565b6200080f565b6201a46662018ea4565b9062003081565b9062001fa2565b936201a481600562000b66565b90620009f1565b6200080f565b6201a49862018eb3565b9062003081565b9062001fa2565b90565b6201a7c4906201a4b86200143b565b506201a7bd6201a7ac6201a7a66201a7926201a7546201a7166201a6d86201a69a6201a65c6201a61f8a6201a6186201a6076201a6016201a5ed6201a5ad6201a56f6201a5316201a5206201a51a896201a513600762000bb0565b90620009f1565b6200080f565b6201a52a6200456b565b9062003081565b6201a5686201a5576201a5518a6201a54a60086200145b565b90620009f1565b6200080f565b6201a56162018dd7565b9062003081565b9062001fa2565b6201a5a66201a5956201a58f896201a588600962002102565b90620009f1565b6200080f565b6201a59f62018de6565b9062003081565b9062001fa2565b6201a5e66201a5d36201a5cd886201a5c6600a6200214c565b90620009f1565b6200080f565b6201a5df600462018df5565b9062018e1b565b9062001fa2565b936201a5fa600b62002196565b90620009f1565b6200080f565b6201a61162018e4a565b9062003081565b9062001fa2565b6201a6556201a6446201a63e8d6201a6375f620007a4565b90620009f1565b6200080f565b6201a64e62018e59565b9062003081565b9062001fa2565b6201a6936201a6826201a67c8c6201a675600162000ad2565b90620009f1565b6200080f565b6201a68c62018e68565b9062003081565b9062001fa2565b6201a6d16201a6c06201a6ba8b6201a6b3600262000af7565b90620009f1565b6200080f565b6201a6ca62018e77565b9062003081565b9062001fa2565b6201a70f6201a6fe6201a6f88a6201a6f1600362000b1c565b90620009f1565b6200080f565b6201a70862018e86565b9062003081565b9062001fa2565b6201a74d6201a73c6201a736896201a72f600462000b41565b90620009f1565b6200080f565b6201a74662018e95565b9062003081565b9062001fa2565b6201a78b6201a77a6201a774886201a76d600562000b66565b90620009f1565b6200080f565b6201a78462018ea4565b9062003081565b9062001fa2565b936201a79f600662000b8b565b90620009f1565b6200080f565b6201a7b662018eb3565b9062003081565b9062001fa2565b90565b6201aae2906201a7d66200143b565b506201aadb6201aaca6201aac46201aab06201aa726201aa346201a9f66201a9b86201a97a6201a93c8a6201a9356201a9246201a91e6201a90b6201a8cb6201a88d6201a84f6201a83e6201a838896201a83160086200145b565b90620009f1565b6200080f565b6201a8486200456b565b9062003081565b6201a8866201a8756201a86f8a6201a868600962002102565b90620009f1565b6200080f565b6201a87f62018dd7565b9062003081565b9062001fa2565b6201a8c46201a8b36201a8ad896201a8a6600a6200214c565b90620009f1565b6200080f565b6201a8bd62018de6565b9062003081565b9062001fa2565b6201a9046201a8f16201a8eb886201a8e4600b62002196565b90620009f1565b6200080f565b6201a8fd600462018df5565b9062018e1b565b9062001fa2565b936201a9175f620007a4565b90620009f1565b6200080f565b6201a92e62018e4a565b9062003081565b9062001fa2565b6201a9736201a9626201a95c8d6201a955600162000ad2565b90620009f1565b6200080f565b6201a96c62018e59565b9062003081565b9062001fa2565b6201a9b16201a9a06201a99a8c6201a993600262000af7565b90620009f1565b6200080f565b6201a9aa62018e68565b9062003081565b9062001fa2565b6201a9ef6201a9de6201a9d88b6201a9d1600362000b1c565b90620009f1565b6200080f565b6201a9e862018e77565b9062003081565b9062001fa2565b6201aa2d6201aa1c6201aa168a6201aa0f600462000b41565b90620009f1565b6200080f565b6201aa2662018e86565b9062003081565b9062001fa2565b6201aa6b6201aa5a6201aa54896201aa4d600562000b66565b90620009f1565b6200080f565b6201aa6462018e95565b9062003081565b9062001fa2565b6201aaa96201aa986201aa92886201aa8b600662000b8b565b90620009f1565b6200080f565b6201aaa262018ea4565b9062003081565b9062001fa2565b936201aabd600762000bb0565b90620009f1565b6200080f565b6201aad462018eb3565b9062003081565b9062001fa2565b90565b6201ae00906201aaf46200143b565b506201adf96201ade86201ade26201adce6201ad906201ad526201ad146201acd66201ac986201ac5a8a6201ac536201ac426201ac3c6201ac286201abe96201abab6201ab6d6201ab5c6201ab56896201ab4f600962002102565b90620009f1565b6200080f565b6201ab666200456b565b9062003081565b6201aba46201ab936201ab8d8a6201ab86600a6200214c565b90620009f1565b6200080f565b6201ab9d62018dd7565b9062003081565b9062001fa2565b6201abe26201abd16201abcb896201abc4600b62002196565b90620009f1565b6200080f565b6201abdb62018de6565b9062003081565b9062001fa2565b6201ac216201ac0e6201ac08886201ac015f620007a4565b90620009f1565b6200080f565b6201ac1a600462018df5565b9062018e1b565b9062001fa2565b936201ac35600162000ad2565b90620009f1565b6200080f565b6201ac4c62018e4a565b9062003081565b9062001fa2565b6201ac916201ac806201ac7a8d6201ac73600262000af7565b90620009f1565b6200080f565b6201ac8a62018e59565b9062003081565b9062001fa2565b6201accf6201acbe6201acb88c6201acb1600362000b1c565b90620009f1565b6200080f565b6201acc862018e68565b9062003081565b9062001fa2565b6201ad0d6201acfc6201acf68b6201acef600462000b41565b90620009f1565b6200080f565b6201ad0662018e77565b9062003081565b9062001fa2565b6201ad4b6201ad3a6201ad348a6201ad2d600562000b66565b90620009f1565b6200080f565b6201ad4462018e86565b9062003081565b9062001fa2565b6201ad896201ad786201ad72896201ad6b600662000b8b565b90620009f1565b6200080f565b6201ad8262018e95565b9062003081565b9062001fa2565b6201adc76201adb66201adb0886201ada9600762000bb0565b90620009f1565b6200080f565b6201adc062018ea4565b9062003081565b9062001fa2565b936201addb60086200145b565b90620009f1565b6200080f565b6201adf262018eb3565b9062003081565b9062001fa2565b90565b6201b11e906201ae126200143b565b506201b1176201b1066201b1006201b0ec6201b0ae6201b0706201b0326201aff46201afb66201af788a6201af716201af606201af5a6201af466201af066201aec96201ae8b6201ae7a6201ae74896201ae6d600a6200214c565b90620009f1565b6200080f565b6201ae846200456b565b9062003081565b6201aec26201aeb16201aeab8a6201aea4600b62002196565b90620009f1565b6200080f565b6201aebb62018dd7565b9062003081565b9062001fa2565b6201aeff6201aeee6201aee8896201aee15f620007a4565b90620009f1565b6200080f565b6201aef862018de6565b9062003081565b9062001fa2565b6201af3f6201af2c6201af26886201af1f600162000ad2565b90620009f1565b6200080f565b6201af38600462018df5565b9062018e1b565b9062001fa2565b936201af53600262000af7565b90620009f1565b6200080f565b6201af6a62018e4a565b9062003081565b9062001fa2565b6201afaf6201af9e6201af988d6201af91600362000b1c565b90620009f1565b6200080f565b6201afa862018e59565b9062003081565b9062001fa2565b6201afed6201afdc6201afd68c6201afcf600462000b41565b90620009f1565b6200080f565b6201afe662018e68565b9062003081565b9062001fa2565b6201b02b6201b01a6201b0148b6201b00d600562000b66565b90620009f1565b6200080f565b6201b02462018e77565b9062003081565b9062001fa2565b6201b0696201b0586201b0528a6201b04b600662000b8b565b90620009f1565b6200080f565b6201b06262018e86565b9062003081565b9062001fa2565b6201b0a76201b0966201b090896201b089600762000bb0565b90620009f1565b6200080f565b6201b0a062018e95565b9062003081565b9062001fa2565b6201b0e56201b0d46201b0ce886201b0c760086200145b565b90620009f1565b6200080f565b6201b0de62018ea4565b9062003081565b9062001fa2565b936201b0f9600962002102565b90620009f1565b6200080f565b6201b11062018eb3565b9062003081565b9062001fa2565b90565b6201b43c906201b1306200143b565b506201b4356201b4246201b41e6201b40a6201b3cc6201b38e6201b3506201b3126201b2d46201b2968a6201b28f6201b27e6201b2786201b2646201b2246201b1e66201b1a96201b1986201b192896201b18b600b62002196565b90620009f1565b6200080f565b6201b1a26200456b565b9062003081565b6201b1df6201b1ce6201b1c88a6201b1c15f620007a4565b90620009f1565b6200080f565b6201b1d862018dd7565b9062003081565b9062001fa2565b6201b21d6201b20c6201b206896201b1ff600162000ad2565b90620009f1565b6200080f565b6201b21662018de6565b9062003081565b9062001fa2565b6201b25d6201b24a6201b244886201b23d600262000af7565b90620009f1565b6200080f565b6201b256600462018df5565b9062018e1b565b9062001fa2565b936201b271600362000b1c565b90620009f1565b6200080f565b6201b28862018e4a565b9062003081565b9062001fa2565b6201b2cd6201b2bc6201b2b68d6201b2af600462000b41565b90620009f1565b6200080f565b6201b2c662018e59565b9062003081565b9062001fa2565b6201b30b6201b2fa6201b2f48c6201b2ed600562000b66565b90620009f1565b6200080f565b6201b30462018e68565b9062003081565b9062001fa2565b6201b3496201b3386201b3328b6201b32b600662000b8b565b90620009f1565b6200080f565b6201b34262018e77565b9062003081565b9062001fa2565b6201b3876201b3766201b3708a6201b369600762000bb0565b90620009f1565b6200080f565b6201b38062018e86565b9062003081565b9062001fa2565b6201b3c56201b3b46201b3ae896201b3a760086200145b565b90620009f1565b6200080f565b6201b3be62018e95565b9062003081565b9062001fa2565b6201b4036201b3f26201b3ec886201b3e5600962002102565b90620009f1565b6200080f565b6201b3fc62018ea4565b9062003081565b9062001fa2565b936201b417600a6200214c565b90620009f1565b6200080f565b6201b42e62018eb3565b9062003081565b9062001fa2565b90565b9067ffffffff00000001916201b4546200143b565b500890565b90565b6201b4756201b46f6201b47b926201b459565b620007a1565b6200012b565b90565b90565b6201b49a6201b4946201b4a0926201b47e565b620007a1565b6200012b565b90565b90565b6201b4bf6201b4b96201b4c5926201b4a3565b620007a1565b6200012b565b90565b90565b6201b4e46201b4de6201b4ea926201b4c8565b620007a1565b6200012b565b90565b90565b6201b5096201b5036201b50f926201b4ed565b620007a1565b6200012b565b90565b90565b6201b52e6201b5286201b534926201b512565b620007a1565b6200012b565b90565b90565b6201b5536201b54d6201b559926201b537565b620007a1565b6200012b565b90565b90565b6201b5786201b5726201b57e926201b55c565b620007a1565b6200012b565b90565b90565b6201b59d6201b5976201b5a3926201b581565b620007a1565b6200012b565b90565b90565b6201b5c26201b5bc6201b5c8926201b5a6565b620007a1565b6200012b565b90565b90565b6201b5e76201b5e16201b5ed926201b5cb565b620007a1565b6200012b565b90565b6201b92b906201b5ff6200143b565b506201b9246201b90a6201b9046201b8f06201b8a96201b8626201b81b6201b7d46201b78d6201b7468a6201b73f6201b7256201b71f6201b70b6201b6c46201b67d6201b6636201b65d886201b656600162000ad2565b90620009f1565b6200080f565b6201b6766780772dc2645b280b6201b45c565b9062003081565b6201b6bd6201b6a36201b69d896201b696600262000af7565b90620009f1565b6200080f565b6201b6b667e796d293a47a64cb6201b481565b9062003081565b9062001fa2565b6201b7046201b6ea6201b6e4886201b6dd600362000b1c565b90620009f1565b6200080f565b6201b6fd67dcedab70f40718ba6201b4a6565b9062003081565b9062001fa2565b936201b718600462000b41565b90620009f1565b6200080f565b6201b73867f4a437f2888ae9096201b4cb565b9062003081565b9062001fa2565b6201b7866201b76c6201b7668d6201b75f600562000b66565b90620009f1565b6200080f565b6201b77f67f97abba0dffb6c506201b4f0565b9062003081565b9062001fa2565b6201b7cd6201b7b36201b7ad8c6201b7a6600662000b8b565b90620009f1565b6200080f565b6201b7c6677f8e41e0b0a6cdff6201b515565b9062003081565b9062001fa2565b6201b8146201b7fa6201b7f48b6201b7ed600762000bb0565b90620009f1565b6200080f565b6201b80d67726af914971c13746201b53a565b9062003081565b9062001fa2565b6201b85b6201b8416201b83b8a6201b83460086200145b565b90620009f1565b6200080f565b6201b8546764dd936da878404d6201b55f565b9062003081565b9062001fa2565b6201b8a26201b8886201b882896201b87b600962002102565b90620009f1565b6200080f565b6201b89b6785418a9fef8a98906201b584565b9062003081565b9062001fa2565b6201b8e96201b8cf6201b8c9886201b8c2600a6200214c565b90620009f1565b6200080f565b6201b8e267156048ee7a7381546201b5a9565b9062003081565b9062001fa2565b936201b8fd600b62002196565b90620009f1565b6200080f565b6201b91d67d841e8ef9dde8ba06201b5ce565b9062003081565b9062001fa2565b90565b90565b6201b94a6201b9446201b950926201b92e565b620007a1565b6200012b565b90565b90565b6201b96f6201b9696201b975926201b953565b620007a1565b6200012b565b90565b90565b6201b9946201b98e6201b99a926201b978565b620007a1565b6200012b565b90565b90565b6201b9b96201b9b36201b9bf926201b99d565b620007a1565b6200012b565b90565b90565b6201b9de6201b9d86201b9e4926201b9c2565b620007a1565b6200012b565b90565b90565b6201ba036201b9fd6201ba09926201b9e7565b620007a1565b6200012b565b90565b90565b6201ba286201ba226201ba2e926201ba0c565b620007a1565b6200012b565b90565b90565b6201ba4d6201ba476201ba53926201ba31565b620007a1565b6200012b565b90565b90565b6201ba726201ba6c6201ba78926201ba56565b620007a1565b6200012b565b90565b90565b6201ba976201ba916201ba9d926201ba7b565b620007a1565b6200012b565b90565b6201bddb906201baaf6200143b565b506201bdd46201bdba6201bdb46201bda06201bd596201bd126201bccb6201bc846201bc3d6201bbf68a6201bbef6201bbd56201bbcf6201bbbb6201bb746201bb2d6201bb136201bb0d886201bb06600162000ad2565b90620009f1565b6200080f565b6201bb2667c1978156516879ad6201b931565b9062003081565b6201bb6d6201bb536201bb4d896201bb46600262000af7565b90620009f1565b6200080f565b6201bb66670ee5dc0ce131268a6201b956565b9062003081565b9062001fa2565b6201bbb46201bb9a6201bb94886201bb8d600362000b1c565b90620009f1565b6200080f565b6201bbad674715b8e5ab34653b6201b97b565b9062003081565b9062001fa2565b936201bbc8600462000b41565b90620009f1565b6200080f565b6201bbe8677f68007619fd8ba96201b9a0565b9062003081565b9062001fa2565b6201bc366201bc1c6201bc168d6201bc0f600562000b66565b90620009f1565b6200080f565b6201bc2f675996a80497e24a6b6201b9c5565b9062003081565b9062001fa2565b6201bc7d6201bc636201bc5d8c6201bc56600662000b8b565b90620009f1565b6200080f565b6201bc7667623708f28fca70e86201b9ea565b9062003081565b9062001fa2565b6201bcc46201bcaa6201bca48b6201bc9d600762000bb0565b90620009f1565b6200080f565b6201bcbd6718737784700c75cd6201ba0f565b9062003081565b9062001fa2565b6201bd0b6201bcf16201bceb8a6201bce460086200145b565b90620009f1565b6200080f565b6201bd0467be2e19f6d07f1a836201ba34565b9062003081565b9062001fa2565b6201bd526201bd386201bd32896201bd2b600962002102565b90620009f1565b6200080f565b6201bd4b67bfe85ababed2d8826201ba59565b9062003081565b9062001fa2565b6201bd996201bd7f6201bd79886201bd72600a6200214c565b90620009f1565b6200080f565b6201bd9267d8a2eb7ef5e707ad6201ba7e565b9062003081565b9062001fa2565b936201bdad600b62002196565b90620009f1565b6200080f565b6201bdcd6785418a9fef8a98906201b584565b9062003081565b9062001fa2565b90565b90565b6201bdfa6201bdf46201be00926201bdde565b620007a1565b6200012b565b90565b90565b6201be1f6201be196201be25926201be03565b620007a1565b6200012b565b90565b90565b6201be446201be3e6201be4a926201be28565b620007a1565b6200012b565b90565b90565b6201be696201be636201be6f926201be4d565b620007a1565b6200012b565b90565b90565b6201be8e6201be886201be94926201be72565b620007a1565b6200012b565b90565b90565b6201beb36201bead6201beb9926201be97565b620007a1565b6200012b565b90565b90565b6201bed86201bed26201bede926201bebc565b620007a1565b6200012b565b90565b90565b6201befd6201bef76201bf03926201bee1565b620007a1565b6200012b565b90565b90565b6201bf226201bf1c6201bf28926201bf06565b620007a1565b6200012b565b90565b6201c266906201bf3a6200143b565b506201c25f6201c2456201c23f6201c22b6201c1e46201c19d6201c1566201c10f6201c0c86201c0818a6201c07a6201c0606201c05a6201c0466201bfff6201bfb86201bf9e6201bf98886201bf91600162000ad2565b90620009f1565b6200080f565b6201bfb16790e80c591f48b6036201bde1565b9062003081565b6201bff86201bfde6201bfd8896201bfd1600262000af7565b90620009f1565b6200080f565b6201bff167a9032a52f930fae66201be06565b9062003081565b9062001fa2565b6201c03f6201c0256201c01f886201c018600362000b1c565b90620009f1565b6200080f565b6201c038671e8916a99c93a88e6201be2b565b9062003081565b9062001fa2565b936201c053600462000b41565b90620009f1565b6200080f565b6201c07367a4911db6a32612da6201be50565b9062003081565b9062001fa2565b6201c0c16201c0a76201c0a18d6201c09a600562000b66565b90620009f1565b6200080f565b6201c0ba6707084430a7307c9a6201be75565b9062003081565b9062001fa2565b6201c1086201c0ee6201c0e88c6201c0e1600662000b8b565b90620009f1565b6200080f565b6201c10167bf150dc4914d380f6201be9a565b9062003081565b9062001fa2565b6201c14f6201c1356201c12f8b6201c128600762000bb0565b90620009f1565b6200080f565b6201c148677fb45d605dd828386201bebf565b9062003081565b9062001fa2565b6201c1966201c17c6201c1768a6201c16f60086200145b565b90620009f1565b6200080f565b6201c18f6702290fe23c20351a6201bee4565b9062003081565b9062001fa2565b6201c1dd6201c1c36201c1bd896201c1b6600962002102565b90620009f1565b6200080f565b6201c1d667be2e19f6d07f1a836201ba34565b9062003081565b9062001fa2565b6201c2246201c20a6201c204886201c1fd600a6200214c565b90620009f1565b6200080f565b6201c21d674db9a2ead2bd72626201bf09565b9062003081565b9062001fa2565b936201c238600b62002196565b90620009f1565b6200080f565b6201c2586764dd936da878404d6201b55f565b9062003081565b9062001fa2565b90565b90565b6201c2856201c27f6201c28b926201c269565b620007a1565b6200012b565b90565b90565b6201c2aa6201c2a46201c2b0926201c28e565b620007a1565b6200012b565b90565b90565b6201c2cf6201c2c96201c2d5926201c2b3565b620007a1565b6200012b565b90565b90565b6201c2f46201c2ee6201c2fa926201c2d8565b620007a1565b6200012b565b90565b90565b6201c3196201c3136201c31f926201c2fd565b620007a1565b6200012b565b90565b90565b6201c33e6201c3386201c344926201c322565b620007a1565b6200012b565b90565b90565b6201c3636201c35d6201c369926201c347565b620007a1565b6200012b565b90565b90565b6201c3886201c3826201c38e926201c36c565b620007a1565b6200012b565b90565b6201c6cc906201c3a06200143b565b506201c6c56201c6ab6201c6a56201c6916201c64a6201c6036201c5bc6201c5756201c52e6201c4e78a6201c4e06201c4c66201c4c06201c4ac6201c4656201c41e6201c4046201c3fe886201c3f7600162000ad2565b90620009f1565b6200080f565b6201c417673a2432625475e3ae6201c26c565b9062003081565b6201c45e6201c4446201c43e896201c437600262000af7565b90620009f1565b6200080f565b6201c457677e33ca8c814280de6201c291565b9062003081565b9062001fa2565b6201c4a56201c48b6201c485886201c47e600362000b1c565b90620009f1565b6200080f565b6201c49e67bba4b5d86b9a3b2c6201c2b6565b9062003081565b9062001fa2565b936201c4b9600462000b41565b90620009f1565b6200080f565b6201c4d9672f7e9aade3fdaec16201c2db565b9062003081565b9062001fa2565b6201c5276201c50d6201c5078d6201c500600562000b66565b90620009f1565b6200080f565b6201c52067ad2f570a5b8545aa6201c300565b9062003081565b9062001fa2565b6201c56e6201c5546201c54e8c6201c547600662000b8b565b90620009f1565b6200080f565b6201c56767c26a0835547671066201c325565b9062003081565b9062001fa2565b6201c5b56201c59b6201c5958b6201c58e600762000bb0565b90620009f1565b6200080f565b6201c5ae67862361aeab0f9b6e6201c34a565b9062003081565b9062001fa2565b6201c5fc6201c5e26201c5dc8a6201c5d560086200145b565b90620009f1565b6200080f565b6201c5f5677fb45d605dd828386201bebf565b9062003081565b9062001fa2565b6201c6436201c6296201c623896201c61c600962002102565b90620009f1565b6200080f565b6201c63c6718737784700c75cd6201ba0f565b9062003081565b9062001fa2565b6201c68a6201c6706201c66a886201c663600a6200214c565b90620009f1565b6200080f565b6201c683671d7f8a2cce1a9d006201c36f565b9062003081565b9062001fa2565b936201c69e600b62002196565b90620009f1565b6200080f565b6201c6be67726af914971c13746201b53a565b9062003081565b9062001fa2565b90565b90565b6201c6eb6201c6e56201c6f1926201c6cf565b620007a1565b6200012b565b90565b90565b6201c7106201c70a6201c716926201c6f4565b620007a1565b6200012b565b90565b90565b6201c7356201c72f6201c73b926201c719565b620007a1565b6200012b565b90565b90565b6201c75a6201c7546201c760926201c73e565b620007a1565b6200012b565b90565b90565b6201c77f6201c7796201c785926201c763565b620007a1565b6200012b565b90565b90565b6201c7a46201c79e6201c7aa926201c788565b620007a1565b6200012b565b90565b90565b6201c7c96201c7c36201c7cf926201c7ad565b620007a1565b6200012b565b90565b6201cb0c906201c7e16200143b565b506201cb056201caeb6201cae56201cad16201ca8a6201ca436201c9fc6201c9b56201c96e6201c9278a6201c9206201c9066201c9006201c8ec6201c8a56201c85e6201c8456201c83f886201c838600162000ad2565b90620009f1565b6200080f565b6201c85766a2d4321cca94fe6201c6d2565b9062003081565b6201c89e6201c8846201c87e896201c877600262000af7565b90620009f1565b6200080f565b6201c89767ad11180f69a8c29e6201c6f7565b9062003081565b9062001fa2565b6201c8e56201c8cb6201c8c5886201c8be600362000b1c565b90620009f1565b6200080f565b6201c8de67e76649f9bd5d5c2e6201c71c565b9062003081565b9062001fa2565b936201c8f9600462000b41565b90620009f1565b6200080f565b6201c91967e7ffd578da4ea43d6201c741565b9062003081565b9062001fa2565b6201c9676201c94d6201c9478d6201c940600562000b66565b90620009f1565b6200080f565b6201c96067ab7f81fef42747706201c766565b9062003081565b9062001fa2565b6201c9ae6201c9946201c98e8c6201c987600662000b8b565b90620009f1565b6200080f565b6201c9a767753b8b1126665c226201c78b565b9062003081565b9062001fa2565b6201c9f56201c9db6201c9d58b6201c9ce600762000bb0565b90620009f1565b6200080f565b6201c9ee67c26a0835547671066201c325565b9062003081565b9062001fa2565b6201ca3c6201ca226201ca1c8a6201ca1560086200145b565b90620009f1565b6200080f565b6201ca3567bf150dc4914d380f6201be9a565b9062003081565b9062001fa2565b6201ca836201ca696201ca63896201ca5c600962002102565b90620009f1565b6200080f565b6201ca7c67623708f28fca70e86201b9ea565b9062003081565b9062001fa2565b6201caca6201cab06201caaa886201caa3600a6200214c565b90620009f1565b6200080f565b6201cac3674b1ba8d40afca97d6201c7b0565b9062003081565b9062001fa2565b936201cade600b62002196565b90620009f1565b6200080f565b6201cafe677f8e41e0b0a6cdff6201b515565b9062003081565b9062001fa2565b90565b90565b6201cb2b6201cb256201cb31926201cb0f565b620007a1565b6200012b565b90565b90565b6201cb506201cb4a6201cb56926201cb34565b620007a1565b6200012b565b90565b90565b6201cb756201cb6f6201cb7b926201cb59565b620007a1565b6200012b565b90565b90565b6201cb9a6201cb946201cba0926201cb7e565b620007a1565b6200012b565b90565b90565b6201cbbf6201cbb96201cbc5926201cba3565b620007a1565b6200012b565b90565b90565b6201cbe46201cbde6201cbea926201cbc8565b620007a1565b6200012b565b90565b6201cf28906201cbfc6200143b565b506201cf216201cf076201cf016201ceed6201cea66201ce5f6201ce186201cdd16201cd8a6201cd438a6201cd3c6201cd226201cd1c6201cd086201ccc16201cc7a6201cc606201cc5a886201cc53600162000ad2565b90620009f1565b6200080f565b6201cc736777736f524010c9326201cb12565b9062003081565b6201ccba6201cca06201cc9a896201cc93600262000af7565b90620009f1565b6200080f565b6201ccb367c75ac6d5b5a10ff36201cb37565b9062003081565b9062001fa2565b6201cd016201cce76201cce1886201ccda600362000b1c565b90620009f1565b6200080f565b6201ccfa67af8e2518a1ece54d6201cb5c565b9062003081565b9062001fa2565b936201cd15600462000b41565b90620009f1565b6200080f565b6201cd356743a608e7afa6b5c26201cb81565b9062003081565b9062001fa2565b6201cd836201cd696201cd638d6201cd5c600562000b66565b90620009f1565b6200080f565b6201cd7c67cb81f535cf98c9e96201cba6565b9062003081565b9062001fa2565b6201cdca6201cdb06201cdaa8c6201cda3600662000b8b565b90620009f1565b6200080f565b6201cdc367ab7f81fef42747706201c766565b9062003081565b9062001fa2565b6201ce116201cdf76201cdf18b6201cdea600762000bb0565b90620009f1565b6200080f565b6201ce0a67ad2f570a5b8545aa6201c300565b9062003081565b9062001fa2565b6201ce586201ce3e6201ce388a6201ce3160086200145b565b90620009f1565b6200080f565b6201ce516707084430a7307c9a6201be75565b9062003081565b9062001fa2565b6201ce9f6201ce856201ce7f896201ce78600962002102565b90620009f1565b6200080f565b6201ce98675996a80497e24a6b6201b9c5565b9062003081565b9062001fa2565b6201cee66201cecc6201cec6886201cebf600a6200214c565b90620009f1565b6200080f565b6201cedf675e40f0c9bb82aab56201cbcb565b9062003081565b9062001fa2565b936201cefa600b62002196565b90620009f1565b6200080f565b6201cf1a67f97abba0dffb6c506201b4f0565b9062003081565b9062001fa2565b90565b90565b6201cf476201cf416201cf4d926201cf2b565b620007a1565b6200012b565b90565b90565b6201cf6c6201cf666201cf72926201cf50565b620007a1565b6200012b565b90565b90565b6201cf916201cf8b6201cf97926201cf75565b620007a1565b6200012b565b90565b90565b6201cfb66201cfb06201cfbc926201cf9a565b620007a1565b6200012b565b90565b90565b6201cfdb6201cfd56201cfe1926201cfbf565b620007a1565b6200012b565b90565b6201d31f906201cff36200143b565b506201d3186201d2fe6201d2f86201d2e46201d29d6201d2566201d20f6201d1c86201d1816201d13a8a6201d1336201d1196201d1136201d0ff6201d0b86201d0716201d0576201d051886201d04a600162000ad2565b90620009f1565b6200080f565b6201d06a67904d3f2804a36c546201cf2e565b9062003081565b6201d0b16201d0976201d091896201d08a600262000af7565b90620009f1565b6200080f565b6201d0aa67f0674a8dc5a387ec6201cf53565b9062003081565b9062001fa2565b6201d0f86201d0de6201d0d8886201d0d1600362000b1c565b90620009f1565b6200080f565b6201d0f167dcda1344cdca873f6201cf78565b9062003081565b9062001fa2565b936201d10c600462000b41565b90620009f1565b6200080f565b6201d12c67ca46546aa99e15756201cf9d565b9062003081565b9062001fa2565b6201d17a6201d1606201d15a8d6201d153600562000b66565b90620009f1565b6200080f565b6201d1736743a608e7afa6b5c26201cb81565b9062003081565b9062001fa2565b6201d1c16201d1a76201d1a18c6201d19a600662000b8b565b90620009f1565b6200080f565b6201d1ba67e7ffd578da4ea43d6201c741565b9062003081565b9062001fa2565b6201d2086201d1ee6201d1e88b6201d1e1600762000bb0565b90620009f1565b6200080f565b6201d201672f7e9aade3fdaec16201c2db565b9062003081565b9062001fa2565b6201d24f6201d2356201d22f8a6201d22860086200145b565b90620009f1565b6200080f565b6201d24867a4911db6a32612da6201be50565b9062003081565b9062001fa2565b6201d2966201d27c6201d276896201d26f600962002102565b90620009f1565b6200080f565b6201d28f677f68007619fd8ba96201b9a0565b9062003081565b9062001fa2565b6201d2dd6201d2c36201d2bd886201d2b6600a6200214c565b90620009f1565b6200080f565b6201d2d667c537d44dc28754036201cfc2565b9062003081565b9062001fa2565b936201d2f1600b62002196565b90620009f1565b6200080f565b6201d31167f4a437f2888ae9096201b4cb565b9062003081565b9062001fa2565b90565b90565b6201d33e6201d3386201d344926201d322565b620007a1565b6200012b565b90565b90565b6201d3636201d35d6201d369926201d347565b620007a1565b6200012b565b90565b90565b6201d3886201d3826201d38e926201d36c565b620007a1565b6200012b565b90565b90565b6201d3ad6201d3a76201d3b3926201d391565b620007a1565b6200012b565b90565b6201d6f1906201d3c56200143b565b506201d6ea6201d6d06201d6ca6201d6b66201d66f6201d6286201d5e16201d59a6201d5536201d50c8a6201d5056201d4eb6201d4e56201d4d16201d48a6201d4436201d4296201d423886201d41c600162000ad2565b90620009f1565b6200080f565b6201d43c67bf9b39e28a16f3546201d325565b9062003081565b6201d4836201d4696201d463896201d45c600262000af7565b90620009f1565b6200080f565b6201d47c67b36d43120eaa5e2b6201d34a565b9062003081565b9062001fa2565b6201d4ca6201d4b06201d4aa886201d4a3600362000b1c565b90620009f1565b6200080f565b6201d4c367cd080204256088e56201d36f565b9062003081565b9062001fa2565b936201d4de600462000b41565b90620009f1565b6200080f565b6201d4fe67dcda1344cdca873f6201cf78565b9062003081565b9062001fa2565b6201d54c6201d5326201d52c8d6201d525600562000b66565b90620009f1565b6200080f565b6201d54567af8e2518a1ece54d6201cb5c565b9062003081565b9062001fa2565b6201d5936201d5796201d5738c6201d56c600662000b8b565b90620009f1565b6200080f565b6201d58c67e76649f9bd5d5c2e6201c71c565b9062003081565b9062001fa2565b6201d5da6201d5c06201d5ba8b6201d5b3600762000bb0565b90620009f1565b6200080f565b6201d5d367bba4b5d86b9a3b2c6201c2b6565b9062003081565b9062001fa2565b6201d6216201d6076201d6018a6201d5fa60086200145b565b90620009f1565b6200080f565b6201d61a671e8916a99c93a88e6201be2b565b9062003081565b9062001fa2565b6201d6686201d64e6201d648896201d641600962002102565b90620009f1565b6200080f565b6201d661674715b8e5ab34653b6201b97b565b9062003081565b9062001fa2565b6201d6af6201d6956201d68f886201d688600a6200214c565b90620009f1565b6200080f565b6201d6a86714a4a64da0b2668f6201d394565b9062003081565b9062001fa2565b936201d6c3600b62002196565b90620009f1565b6200080f565b6201d6e367dcedab70f40718ba6201b4a6565b9062003081565b9062001fa2565b90565b90565b6201d7106201d70a6201d716926201d6f4565b620007a1565b6200012b565b90565b90565b6201d7356201d72f6201d73b926201d719565b620007a1565b6200012b565b90565b90565b6201d75a6201d7546201d760926201d73e565b620007a1565b6200012b565b90565b6201da9e906201d7726200143b565b506201da976201da7d6201da776201da636201da1c6201d9d56201d98e6201d9476201d9006201d8b98a6201d8b26201d8986201d8926201d87e6201d8376201d7f06201d7d66201d7d0886201d7c9600162000ad2565b90620009f1565b6200080f565b6201d7e9673a1ded54a6cd058b6201d6f7565b9062003081565b6201d8306201d8166201d810896201d809600262000af7565b90620009f1565b6200080f565b6201d829676f232aab4b533a256201d71c565b9062003081565b9062001fa2565b6201d8776201d85d6201d857886201d850600362000b1c565b90620009f1565b6200080f565b6201d87067b36d43120eaa5e2b6201d34a565b9062003081565b9062001fa2565b936201d88b600462000b41565b90620009f1565b6200080f565b6201d8ab67f0674a8dc5a387ec6201cf53565b9062003081565b9062001fa2565b6201d8f96201d8df6201d8d98d6201d8d2600562000b66565b90620009f1565b6200080f565b6201d8f267c75ac6d5b5a10ff36201cb37565b9062003081565b9062001fa2565b6201d9406201d9266201d9208c6201d919600662000b8b565b90620009f1565b6200080f565b6201d93967ad11180f69a8c29e6201c6f7565b9062003081565b9062001fa2565b6201d9876201d96d6201d9678b6201d960600762000bb0565b90620009f1565b6200080f565b6201d980677e33ca8c814280de6201c291565b9062003081565b9062001fa2565b6201d9ce6201d9b46201d9ae8a6201d9a760086200145b565b90620009f1565b6200080f565b6201d9c767a9032a52f930fae66201be06565b9062003081565b9062001fa2565b6201da156201d9fb6201d9f5896201d9ee600962002102565b90620009f1565b6200080f565b6201da0e670ee5dc0ce131268a6201b956565b9062003081565b9062001fa2565b6201da5c6201da426201da3c886201da35600a6200214c565b90620009f1565b6200080f565b6201da5567b124c33152a2421a6201d741565b9062003081565b9062001fa2565b936201da70600b62002196565b90620009f1565b6200080f565b6201da9067e796d293a47a64cb6201b481565b9062003081565b9062001fa2565b90565b90565b6201dabd6201dab76201dac3926201daa1565b620007a1565b6200012b565b90565b90565b6201dae26201dadc6201dae8926201dac6565b620007a1565b6200012b565b90565b6201de25906201dafa6200143b565b506201de1e6201de046201ddfe6201ddea6201dda36201dd5c6201dd156201dcce6201dc886201dc418a6201dc3a6201dc206201dc1a6201dc066201dbbf6201db786201db5e6201db58886201db51600162000ad2565b90620009f1565b6200080f565b6201db716742392870da5737cf6201daa4565b9062003081565b6201dbb86201db9e6201db98896201db91600262000af7565b90620009f1565b6200080f565b6201dbb1673a1ded54a6cd058b6201d6f7565b9062003081565b9062001fa2565b6201dbff6201dbe56201dbdf886201dbd8600362000b1c565b90620009f1565b6200080f565b6201dbf867bf9b39e28a16f3546201d325565b9062003081565b9062001fa2565b936201dc13600462000b41565b90620009f1565b6200080f565b6201dc3367904d3f2804a36c546201cf2e565b9062003081565b9062001fa2565b6201dc816201dc676201dc618d6201dc5a600562000b66565b90620009f1565b6200080f565b6201dc7a6777736f524010c9326201cb12565b9062003081565b9062001fa2565b6201dcc76201dcae6201dca88c6201dca1600662000b8b565b90620009f1565b6200080f565b6201dcc066a2d4321cca94fe6201c6d2565b9062003081565b9062001fa2565b6201dd0e6201dcf46201dcee8b6201dce7600762000bb0565b90620009f1565b6200080f565b6201dd07673a2432625475e3ae6201c26c565b9062003081565b9062001fa2565b6201dd556201dd3b6201dd358a6201dd2e60086200145b565b90620009f1565b6200080f565b6201dd4e6790e80c591f48b6036201bde1565b9062003081565b9062001fa2565b6201dd9c6201dd826201dd7c896201dd75600962002102565b90620009f1565b6200080f565b6201dd9567c1978156516879ad6201b931565b9062003081565b9062001fa2565b6201dde36201ddc96201ddc3886201ddbc600a6200214c565b90620009f1565b6200080f565b6201dddc67dc927721da922cf86201dac9565b9062003081565b9062001fa2565b936201ddf7600b62002196565b90620009f1565b6200080f565b6201de176780772dc2645b280b6201b45c565b9062003081565b9062001fa2565b90565b90565b6201de446201de3e6201de4a926201de28565b620007a1565b6200012b565b90565b6201e188906201de5c6200143b565b506201e1816201e1676201e1616201e14d6201e1066201e0bf6201e0786201e0316201dfea6201dfa38a6201df9c6201df826201df7c6201df686201df216201deda6201dec06201deba886201deb3600162000ad2565b90620009f1565b6200080f565b6201ded367dc927721da922cf86201dac9565b9062003081565b6201df1a6201df006201defa896201def3600262000af7565b90620009f1565b6200080f565b6201df1367b124c33152a2421a6201d741565b9062003081565b9062001fa2565b6201df616201df476201df41886201df3a600362000b1c565b90620009f1565b6200080f565b6201df5a6714a4a64da0b2668f6201d394565b9062003081565b9062001fa2565b936201df75600462000b41565b90620009f1565b6200080f565b6201df9567c537d44dc28754036201cfc2565b9062003081565b9062001fa2565b6201dfe36201dfc96201dfc38d6201dfbc600562000b66565b90620009f1565b6200080f565b6201dfdc675e40f0c9bb82aab56201cbcb565b9062003081565b9062001fa2565b6201e02a6201e0106201e00a8c6201e003600662000b8b565b90620009f1565b6200080f565b6201e023674b1ba8d40afca97d6201c7b0565b9062003081565b9062001fa2565b6201e0716201e0576201e0518b6201e04a600762000bb0565b90620009f1565b6200080f565b6201e06a671d7f8a2cce1a9d006201c36f565b9062003081565b9062001fa2565b6201e0b86201e09e6201e0988a6201e09160086200145b565b90620009f1565b6200080f565b6201e0b1674db9a2ead2bd72626201bf09565b9062003081565b9062001fa2565b6201e0ff6201e0e56201e0df896201e0d8600962002102565b90620009f1565b6200080f565b6201e0f867d8a2eb7ef5e707ad6201ba7e565b9062003081565b9062001fa2565b6201e1466201e12c6201e126886201e11f600a6200214c565b90620009f1565b6200080f565b6201e13f6791f7562377e81df56201de2b565b9062003081565b9062001fa2565b936201e15a600b62002196565b90620009f1565b6200080f565b6201e17a67156048ee7a7381546201b5a9565b9062003081565b9062001fa2565b9056fea26469706673582212207e2c4364a07d3a3bc1118958ebfc0621bca7eaa19e33c2acabb072f2c31abed564736f6c63430008180033",
}

// GoldilocksPoseidonABI is the input ABI used to generate the binding from.
// Deprecated: Use GoldilocksPoseidonMetaData.ABI instead.
var GoldilocksPoseidonABI = GoldilocksPoseidonMetaData.ABI

// GoldilocksPoseidonBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GoldilocksPoseidonMetaData.Bin instead.
var GoldilocksPoseidonBin = GoldilocksPoseidonMetaData.Bin

// DeployGoldilocksPoseidon deploys a new Ethereum contract, binding an instance of GoldilocksPoseidon to it.
func DeployGoldilocksPoseidon(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GoldilocksPoseidon, error) {
	parsed, err := GoldilocksPoseidonMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GoldilocksPoseidonBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GoldilocksPoseidon{GoldilocksPoseidonCaller: GoldilocksPoseidonCaller{contract: contract}, GoldilocksPoseidonTransactor: GoldilocksPoseidonTransactor{contract: contract}, GoldilocksPoseidonFilterer: GoldilocksPoseidonFilterer{contract: contract}}, nil
}

// GoldilocksPoseidon is an auto generated Go binding around an Ethereum contract.
type GoldilocksPoseidon struct {
	GoldilocksPoseidonCaller     // Read-only binding to the contract
	GoldilocksPoseidonTransactor // Write-only binding to the contract
	GoldilocksPoseidonFilterer   // Log filterer for contract events
}

// GoldilocksPoseidonCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoldilocksPoseidonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldilocksPoseidonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoldilocksPoseidonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldilocksPoseidonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoldilocksPoseidonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldilocksPoseidonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoldilocksPoseidonSession struct {
	Contract     *GoldilocksPoseidon // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GoldilocksPoseidonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoldilocksPoseidonCallerSession struct {
	Contract *GoldilocksPoseidonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GoldilocksPoseidonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoldilocksPoseidonTransactorSession struct {
	Contract     *GoldilocksPoseidonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GoldilocksPoseidonRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoldilocksPoseidonRaw struct {
	Contract *GoldilocksPoseidon // Generic contract binding to access the raw methods on
}

// GoldilocksPoseidonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoldilocksPoseidonCallerRaw struct {
	Contract *GoldilocksPoseidonCaller // Generic read-only contract binding to access the raw methods on
}

// GoldilocksPoseidonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoldilocksPoseidonTransactorRaw struct {
	Contract *GoldilocksPoseidonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoldilocksPoseidon creates a new instance of GoldilocksPoseidon, bound to a specific deployed contract.
func NewGoldilocksPoseidon(address common.Address, backend bind.ContractBackend) (*GoldilocksPoseidon, error) {
	contract, err := bindGoldilocksPoseidon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GoldilocksPoseidon{GoldilocksPoseidonCaller: GoldilocksPoseidonCaller{contract: contract}, GoldilocksPoseidonTransactor: GoldilocksPoseidonTransactor{contract: contract}, GoldilocksPoseidonFilterer: GoldilocksPoseidonFilterer{contract: contract}}, nil
}

// NewGoldilocksPoseidonCaller creates a new read-only instance of GoldilocksPoseidon, bound to a specific deployed contract.
func NewGoldilocksPoseidonCaller(address common.Address, caller bind.ContractCaller) (*GoldilocksPoseidonCaller, error) {
	contract, err := bindGoldilocksPoseidon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoldilocksPoseidonCaller{contract: contract}, nil
}

// NewGoldilocksPoseidonTransactor creates a new write-only instance of GoldilocksPoseidon, bound to a specific deployed contract.
func NewGoldilocksPoseidonTransactor(address common.Address, transactor bind.ContractTransactor) (*GoldilocksPoseidonTransactor, error) {
	contract, err := bindGoldilocksPoseidon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoldilocksPoseidonTransactor{contract: contract}, nil
}

// NewGoldilocksPoseidonFilterer creates a new log filterer instance of GoldilocksPoseidon, bound to a specific deployed contract.
func NewGoldilocksPoseidonFilterer(address common.Address, filterer bind.ContractFilterer) (*GoldilocksPoseidonFilterer, error) {
	contract, err := bindGoldilocksPoseidon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoldilocksPoseidonFilterer{contract: contract}, nil
}

// bindGoldilocksPoseidon binds a generic wrapper to an already deployed contract.
func bindGoldilocksPoseidon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GoldilocksPoseidonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoldilocksPoseidon *GoldilocksPoseidonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoldilocksPoseidon.Contract.GoldilocksPoseidonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoldilocksPoseidon *GoldilocksPoseidonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoldilocksPoseidon.Contract.GoldilocksPoseidonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoldilocksPoseidon *GoldilocksPoseidonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoldilocksPoseidon.Contract.GoldilocksPoseidonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoldilocksPoseidon *GoldilocksPoseidonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoldilocksPoseidon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoldilocksPoseidon *GoldilocksPoseidonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoldilocksPoseidon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoldilocksPoseidon *GoldilocksPoseidonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoldilocksPoseidon.Contract.contract.Transact(opts, method, params...)
}

// DecodeHashOut is a free data retrieval call binding the contract method 0xbbb62147.
//
// Solidity: function decodeHashOut(bytes32 encoded) pure returns(uint256[4] output)
func (_GoldilocksPoseidon *GoldilocksPoseidonCaller) DecodeHashOut(opts *bind.CallOpts, encoded [32]byte) ([4]*big.Int, error) {
	var out []interface{}
	err := _GoldilocksPoseidon.contract.Call(opts, &out, "decodeHashOut", encoded)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// DecodeHashOut is a free data retrieval call binding the contract method 0xbbb62147.
//
// Solidity: function decodeHashOut(bytes32 encoded) pure returns(uint256[4] output)
func (_GoldilocksPoseidon *GoldilocksPoseidonSession) DecodeHashOut(encoded [32]byte) ([4]*big.Int, error) {
	return _GoldilocksPoseidon.Contract.DecodeHashOut(&_GoldilocksPoseidon.CallOpts, encoded)
}

// DecodeHashOut is a free data retrieval call binding the contract method 0xbbb62147.
//
// Solidity: function decodeHashOut(bytes32 encoded) pure returns(uint256[4] output)
func (_GoldilocksPoseidon *GoldilocksPoseidonCallerSession) DecodeHashOut(encoded [32]byte) ([4]*big.Int, error) {
	return _GoldilocksPoseidon.Contract.DecodeHashOut(&_GoldilocksPoseidon.CallOpts, encoded)
}

// EncodeHashOut is a free data retrieval call binding the contract method 0xf8978579.
//
// Solidity: function encodeHashOut(uint256[4] value) pure returns(bytes32 encoded)
func (_GoldilocksPoseidon *GoldilocksPoseidonCaller) EncodeHashOut(opts *bind.CallOpts, value [4]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _GoldilocksPoseidon.contract.Call(opts, &out, "encodeHashOut", value)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EncodeHashOut is a free data retrieval call binding the contract method 0xf8978579.
//
// Solidity: function encodeHashOut(uint256[4] value) pure returns(bytes32 encoded)
func (_GoldilocksPoseidon *GoldilocksPoseidonSession) EncodeHashOut(value [4]*big.Int) ([32]byte, error) {
	return _GoldilocksPoseidon.Contract.EncodeHashOut(&_GoldilocksPoseidon.CallOpts, value)
}

// EncodeHashOut is a free data retrieval call binding the contract method 0xf8978579.
//
// Solidity: function encodeHashOut(uint256[4] value) pure returns(bytes32 encoded)
func (_GoldilocksPoseidon *GoldilocksPoseidonCallerSession) EncodeHashOut(value [4]*big.Int) ([32]byte, error) {
	return _GoldilocksPoseidon.Contract.EncodeHashOut(&_GoldilocksPoseidon.CallOpts, value)
}

// HashNToMNoPad is a free data retrieval call binding the contract method 0x1a10dabc.
//
// Solidity: function hashNToMNoPad(uint256[] input, uint256 numOutputs) pure returns(uint256[] output)
func (_GoldilocksPoseidon *GoldilocksPoseidonCaller) HashNToMNoPad(opts *bind.CallOpts, input []*big.Int, numOutputs *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _GoldilocksPoseidon.contract.Call(opts, &out, "hashNToMNoPad", input, numOutputs)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// HashNToMNoPad is a free data retrieval call binding the contract method 0x1a10dabc.
//
// Solidity: function hashNToMNoPad(uint256[] input, uint256 numOutputs) pure returns(uint256[] output)
func (_GoldilocksPoseidon *GoldilocksPoseidonSession) HashNToMNoPad(input []*big.Int, numOutputs *big.Int) ([]*big.Int, error) {
	return _GoldilocksPoseidon.Contract.HashNToMNoPad(&_GoldilocksPoseidon.CallOpts, input, numOutputs)
}

// HashNToMNoPad is a free data retrieval call binding the contract method 0x1a10dabc.
//
// Solidity: function hashNToMNoPad(uint256[] input, uint256 numOutputs) pure returns(uint256[] output)
func (_GoldilocksPoseidon *GoldilocksPoseidonCallerSession) HashNToMNoPad(input []*big.Int, numOutputs *big.Int) ([]*big.Int, error) {
	return _GoldilocksPoseidon.Contract.HashNToMNoPad(&_GoldilocksPoseidon.CallOpts, input, numOutputs)
}

// HashNToMWithCap is a free data retrieval call binding the contract method 0x10e2e358.
//
// Solidity: function hashNToMWithCap(uint256[] input, uint256[4] capacity, uint256 numOutputs) pure returns(uint256[] output)
func (_GoldilocksPoseidon *GoldilocksPoseidonCaller) HashNToMWithCap(opts *bind.CallOpts, input []*big.Int, capacity [4]*big.Int, numOutputs *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _GoldilocksPoseidon.contract.Call(opts, &out, "hashNToMWithCap", input, capacity, numOutputs)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// HashNToMWithCap is a free data retrieval call binding the contract method 0x10e2e358.
//
// Solidity: function hashNToMWithCap(uint256[] input, uint256[4] capacity, uint256 numOutputs) pure returns(uint256[] output)
func (_GoldilocksPoseidon *GoldilocksPoseidonSession) HashNToMWithCap(input []*big.Int, capacity [4]*big.Int, numOutputs *big.Int) ([]*big.Int, error) {
	return _GoldilocksPoseidon.Contract.HashNToMWithCap(&_GoldilocksPoseidon.CallOpts, input, capacity, numOutputs)
}

// HashNToMWithCap is a free data retrieval call binding the contract method 0x10e2e358.
//
// Solidity: function hashNToMWithCap(uint256[] input, uint256[4] capacity, uint256 numOutputs) pure returns(uint256[] output)
func (_GoldilocksPoseidon *GoldilocksPoseidonCallerSession) HashNToMWithCap(input []*big.Int, capacity [4]*big.Int, numOutputs *big.Int) ([]*big.Int, error) {
	return _GoldilocksPoseidon.Contract.HashNToMWithCap(&_GoldilocksPoseidon.CallOpts, input, capacity, numOutputs)
}

// Permute is a free data retrieval call binding the contract method 0x496f3b22.
//
// Solidity: function permute(uint256[12] state) pure returns(uint256[12] newState)
func (_GoldilocksPoseidon *GoldilocksPoseidonCaller) Permute(opts *bind.CallOpts, state [12]*big.Int) ([12]*big.Int, error) {
	var out []interface{}
	err := _GoldilocksPoseidon.contract.Call(opts, &out, "permute", state)

	if err != nil {
		return *new([12]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([12]*big.Int)).(*[12]*big.Int)

	return out0, err

}

// Permute is a free data retrieval call binding the contract method 0x496f3b22.
//
// Solidity: function permute(uint256[12] state) pure returns(uint256[12] newState)
func (_GoldilocksPoseidon *GoldilocksPoseidonSession) Permute(state [12]*big.Int) ([12]*big.Int, error) {
	return _GoldilocksPoseidon.Contract.Permute(&_GoldilocksPoseidon.CallOpts, state)
}

// Permute is a free data retrieval call binding the contract method 0x496f3b22.
//
// Solidity: function permute(uint256[12] state) pure returns(uint256[12] newState)
func (_GoldilocksPoseidon *GoldilocksPoseidonCallerSession) Permute(state [12]*big.Int) ([12]*big.Int, error) {
	return _GoldilocksPoseidon.Contract.Permute(&_GoldilocksPoseidon.CallOpts, state)
}

// TwoToOne is a free data retrieval call binding the contract method 0x49abd354.
//
// Solidity: function twoToOne(bytes32 left, bytes32 right) pure returns(bytes32 output)
func (_GoldilocksPoseidon *GoldilocksPoseidonCaller) TwoToOne(opts *bind.CallOpts, left [32]byte, right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GoldilocksPoseidon.contract.Call(opts, &out, "twoToOne", left, right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TwoToOne is a free data retrieval call binding the contract method 0x49abd354.
//
// Solidity: function twoToOne(bytes32 left, bytes32 right) pure returns(bytes32 output)
func (_GoldilocksPoseidon *GoldilocksPoseidonSession) TwoToOne(left [32]byte, right [32]byte) ([32]byte, error) {
	return _GoldilocksPoseidon.Contract.TwoToOne(&_GoldilocksPoseidon.CallOpts, left, right)
}

// TwoToOne is a free data retrieval call binding the contract method 0x49abd354.
//
// Solidity: function twoToOne(bytes32 left, bytes32 right) pure returns(bytes32 output)
func (_GoldilocksPoseidon *GoldilocksPoseidonCallerSession) TwoToOne(left [32]byte, right [32]byte) ([32]byte, error) {
	return _GoldilocksPoseidon.Contract.TwoToOne(&_GoldilocksPoseidon.CallOpts, left, right)
}

// IAccountMetaData contains all meta data concerning the IAccount contract.
var IAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"validationResult\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountMetaData.ABI instead.
var IAccountABI = IAccountMetaData.ABI

// IAccount is an auto generated Go binding around an Ethereum contract.
type IAccount struct {
	IAccountCaller     // Read-only binding to the contract
	IAccountTransactor // Write-only binding to the contract
	IAccountFilterer   // Log filterer for contract events
}

// IAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountSession struct {
	Contract     *IAccount         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountCallerSession struct {
	Contract *IAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountTransactorSession struct {
	Contract     *IAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountRaw struct {
	Contract *IAccount // Generic contract binding to access the raw methods on
}

// IAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountCallerRaw struct {
	Contract *IAccountCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountTransactorRaw struct {
	Contract *IAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccount creates a new instance of IAccount, bound to a specific deployed contract.
func NewIAccount(address common.Address, backend bind.ContractBackend) (*IAccount, error) {
	contract, err := bindIAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccount{IAccountCaller: IAccountCaller{contract: contract}, IAccountTransactor: IAccountTransactor{contract: contract}, IAccountFilterer: IAccountFilterer{contract: contract}}, nil
}

// NewIAccountCaller creates a new read-only instance of IAccount, bound to a specific deployed contract.
func NewIAccountCaller(address common.Address, caller bind.ContractCaller) (*IAccountCaller, error) {
	contract, err := bindIAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountCaller{contract: contract}, nil
}

// NewIAccountTransactor creates a new write-only instance of IAccount, bound to a specific deployed contract.
func NewIAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountTransactor, error) {
	contract, err := bindIAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountTransactor{contract: contract}, nil
}

// NewIAccountFilterer creates a new log filterer instance of IAccount, bound to a specific deployed contract.
func NewIAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountFilterer, error) {
	contract, err := bindIAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountFilterer{contract: contract}, nil
}

// bindIAccount binds a generic wrapper to an already deployed contract.
func bindIAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccount *IAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccount.Contract.IAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccount *IAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccount.Contract.IAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccount *IAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccount.Contract.IAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccount *IAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccount *IAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccount *IAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccount.Contract.contract.Transact(opts, method, params...)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x5c51ffd1.
//
// Solidity: function validateUserOp(address _owner) returns(bool validationResult)
func (_IAccount *IAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _IAccount.contract.Transact(opts, "validateUserOp", _owner)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x5c51ffd1.
//
// Solidity: function validateUserOp(address _owner) returns(bool validationResult)
func (_IAccount *IAccountSession) ValidateUserOp(_owner common.Address) (*types.Transaction, error) {
	return _IAccount.Contract.ValidateUserOp(&_IAccount.TransactOpts, _owner)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x5c51ffd1.
//
// Solidity: function validateUserOp(address _owner) returns(bool validationResult)
func (_IAccount *IAccountTransactorSession) ValidateUserOp(_owner common.Address) (*types.Transaction, error) {
	return _IAccount.Contract.ValidateUserOp(&_IAccount.TransactOpts, _owner)
}

// IAccountExecuteMetaData contains all meta data concerning the IAccountExecute contract.
var IAccountExecuteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"executeUserOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAccountExecuteABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountExecuteMetaData.ABI instead.
var IAccountExecuteABI = IAccountExecuteMetaData.ABI

// IAccountExecute is an auto generated Go binding around an Ethereum contract.
type IAccountExecute struct {
	IAccountExecuteCaller     // Read-only binding to the contract
	IAccountExecuteTransactor // Write-only binding to the contract
	IAccountExecuteFilterer   // Log filterer for contract events
}

// IAccountExecuteCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountExecuteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountExecuteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountExecuteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountExecuteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountExecuteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountExecuteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountExecuteSession struct {
	Contract     *IAccountExecute  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountExecuteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountExecuteCallerSession struct {
	Contract *IAccountExecuteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IAccountExecuteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountExecuteTransactorSession struct {
	Contract     *IAccountExecuteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAccountExecuteRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountExecuteRaw struct {
	Contract *IAccountExecute // Generic contract binding to access the raw methods on
}

// IAccountExecuteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountExecuteCallerRaw struct {
	Contract *IAccountExecuteCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountExecuteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountExecuteTransactorRaw struct {
	Contract *IAccountExecuteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccountExecute creates a new instance of IAccountExecute, bound to a specific deployed contract.
func NewIAccountExecute(address common.Address, backend bind.ContractBackend) (*IAccountExecute, error) {
	contract, err := bindIAccountExecute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccountExecute{IAccountExecuteCaller: IAccountExecuteCaller{contract: contract}, IAccountExecuteTransactor: IAccountExecuteTransactor{contract: contract}, IAccountExecuteFilterer: IAccountExecuteFilterer{contract: contract}}, nil
}

// NewIAccountExecuteCaller creates a new read-only instance of IAccountExecute, bound to a specific deployed contract.
func NewIAccountExecuteCaller(address common.Address, caller bind.ContractCaller) (*IAccountExecuteCaller, error) {
	contract, err := bindIAccountExecute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountExecuteCaller{contract: contract}, nil
}

// NewIAccountExecuteTransactor creates a new write-only instance of IAccountExecute, bound to a specific deployed contract.
func NewIAccountExecuteTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountExecuteTransactor, error) {
	contract, err := bindIAccountExecute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountExecuteTransactor{contract: contract}, nil
}

// NewIAccountExecuteFilterer creates a new log filterer instance of IAccountExecute, bound to a specific deployed contract.
func NewIAccountExecuteFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountExecuteFilterer, error) {
	contract, err := bindIAccountExecute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountExecuteFilterer{contract: contract}, nil
}

// bindIAccountExecute binds a generic wrapper to an already deployed contract.
func bindIAccountExecute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccountExecuteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountExecute *IAccountExecuteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountExecute.Contract.IAccountExecuteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountExecute *IAccountExecuteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountExecute.Contract.IAccountExecuteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountExecute *IAccountExecuteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountExecute.Contract.IAccountExecuteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountExecute *IAccountExecuteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountExecute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountExecute *IAccountExecuteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountExecute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountExecute *IAccountExecuteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountExecute.Contract.contract.Transact(opts, method, params...)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0xa943c00f.
//
// Solidity: function executeUserOp((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp, bytes32 userOpHash) returns()
func (_IAccountExecute *IAccountExecuteTransactor) ExecuteUserOp(opts *bind.TransactOpts, userOp BaseStructPackedUserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _IAccountExecute.contract.Transact(opts, "executeUserOp", userOp, userOpHash)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0xa943c00f.
//
// Solidity: function executeUserOp((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp, bytes32 userOpHash) returns()
func (_IAccountExecute *IAccountExecuteSession) ExecuteUserOp(userOp BaseStructPackedUserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _IAccountExecute.Contract.ExecuteUserOp(&_IAccountExecute.TransactOpts, userOp, userOpHash)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0xa943c00f.
//
// Solidity: function executeUserOp((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp, bytes32 userOpHash) returns()
func (_IAccountExecute *IAccountExecuteTransactorSession) ExecuteUserOp(userOp BaseStructPackedUserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _IAccountExecute.Contract.ExecuteUserOp(&_IAccountExecute.TransactOpts, userOp, userOpHash)
}

// IConfigManagerMetaData contains all meta data concerning the IConfigManager contract.
var IConfigManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getChainConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"internalType\":\"structIConfigManager.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IConfigManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IConfigManagerMetaData.ABI instead.
var IConfigManagerABI = IConfigManagerMetaData.ABI

// IConfigManager is an auto generated Go binding around an Ethereum contract.
type IConfigManager struct {
	IConfigManagerCaller     // Read-only binding to the contract
	IConfigManagerTransactor // Write-only binding to the contract
	IConfigManagerFilterer   // Log filterer for contract events
}

// IConfigManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IConfigManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfigManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IConfigManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfigManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IConfigManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfigManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IConfigManagerSession struct {
	Contract     *IConfigManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IConfigManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IConfigManagerCallerSession struct {
	Contract *IConfigManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IConfigManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IConfigManagerTransactorSession struct {
	Contract     *IConfigManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IConfigManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IConfigManagerRaw struct {
	Contract *IConfigManager // Generic contract binding to access the raw methods on
}

// IConfigManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IConfigManagerCallerRaw struct {
	Contract *IConfigManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IConfigManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IConfigManagerTransactorRaw struct {
	Contract *IConfigManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIConfigManager creates a new instance of IConfigManager, bound to a specific deployed contract.
func NewIConfigManager(address common.Address, backend bind.ContractBackend) (*IConfigManager, error) {
	contract, err := bindIConfigManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IConfigManager{IConfigManagerCaller: IConfigManagerCaller{contract: contract}, IConfigManagerTransactor: IConfigManagerTransactor{contract: contract}, IConfigManagerFilterer: IConfigManagerFilterer{contract: contract}}, nil
}

// NewIConfigManagerCaller creates a new read-only instance of IConfigManager, bound to a specific deployed contract.
func NewIConfigManagerCaller(address common.Address, caller bind.ContractCaller) (*IConfigManagerCaller, error) {
	contract, err := bindIConfigManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IConfigManagerCaller{contract: contract}, nil
}

// NewIConfigManagerTransactor creates a new write-only instance of IConfigManager, bound to a specific deployed contract.
func NewIConfigManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IConfigManagerTransactor, error) {
	contract, err := bindIConfigManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IConfigManagerTransactor{contract: contract}, nil
}

// NewIConfigManagerFilterer creates a new log filterer instance of IConfigManager, bound to a specific deployed contract.
func NewIConfigManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IConfigManagerFilterer, error) {
	contract, err := bindIConfigManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IConfigManagerFilterer{contract: contract}, nil
}

// bindIConfigManager binds a generic wrapper to an already deployed contract.
func bindIConfigManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IConfigManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConfigManager *IConfigManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IConfigManager.Contract.IConfigManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConfigManager *IConfigManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConfigManager.Contract.IConfigManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConfigManager *IConfigManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConfigManager.Contract.IConfigManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConfigManager *IConfigManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IConfigManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConfigManager *IConfigManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConfigManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConfigManager *IConfigManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConfigManager.Contract.contract.Transact(opts, method, params...)
}

// GetChainConfigs is a free data retrieval call binding the contract method 0x42f1ec69.
//
// Solidity: function getChainConfigs(uint64 chainId) view returns((address,address))
func (_IConfigManager *IConfigManagerCaller) GetChainConfigs(opts *bind.CallOpts, chainId uint64) (IConfigManagerConfig, error) {
	var out []interface{}
	err := _IConfigManager.contract.Call(opts, &out, "getChainConfigs", chainId)

	if err != nil {
		return *new(IConfigManagerConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IConfigManagerConfig)).(*IConfigManagerConfig)

	return out0, err

}

// GetChainConfigs is a free data retrieval call binding the contract method 0x42f1ec69.
//
// Solidity: function getChainConfigs(uint64 chainId) view returns((address,address))
func (_IConfigManager *IConfigManagerSession) GetChainConfigs(chainId uint64) (IConfigManagerConfig, error) {
	return _IConfigManager.Contract.GetChainConfigs(&_IConfigManager.CallOpts, chainId)
}

// GetChainConfigs is a free data retrieval call binding the contract method 0x42f1ec69.
//
// Solidity: function getChainConfigs(uint64 chainId) view returns((address,address))
func (_IConfigManager *IConfigManagerCallerSession) GetChainConfigs(chainId uint64) (IConfigManagerConfig, error) {
	return _IConfigManager.Contract.GetChainConfigs(&_IConfigManager.CallOpts, chainId)
}

// GetMainChainId is a free data retrieval call binding the contract method 0x750c67ba.
//
// Solidity: function getMainChainId() view returns(uint64)
func (_IConfigManager *IConfigManagerCaller) GetMainChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IConfigManager.contract.Call(opts, &out, "getMainChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetMainChainId is a free data retrieval call binding the contract method 0x750c67ba.
//
// Solidity: function getMainChainId() view returns(uint64)
func (_IConfigManager *IConfigManagerSession) GetMainChainId() (uint64, error) {
	return _IConfigManager.Contract.GetMainChainId(&_IConfigManager.CallOpts)
}

// GetMainChainId is a free data retrieval call binding the contract method 0x750c67ba.
//
// Solidity: function getMainChainId() view returns(uint64)
func (_IConfigManager *IConfigManagerCallerSession) GetMainChainId() (uint64, error) {
	return _IConfigManager.Contract.GetMainChainId(&_IConfigManager.CallOpts)
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IEntryPointMetaData contains all meta data concerning the IEntryPoint contract.
var IEntryPointMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"name\":\"DelegateAndRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"inner\",\"type\":\"bytes\"}],\"name\":\"FailedOpWithRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"PostOpReverted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotEqual\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"did\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"innerExec\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"estimateCrossMessageParamsCrossGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getChainConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"internalType\":\"structIConfigManager.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPreGasBalanceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSync\",\"type\":\"bool\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"sendUserOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitDepositOperation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitDepositOperationByRemote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submitWithdrawOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"}],\"name\":\"syncBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEntryPointABI is the input ABI used to generate the binding from.
// Deprecated: Use IEntryPointMetaData.ABI instead.
var IEntryPointABI = IEntryPointMetaData.ABI

// IEntryPoint is an auto generated Go binding around an Ethereum contract.
type IEntryPoint struct {
	IEntryPointCaller     // Read-only binding to the contract
	IEntryPointTransactor // Write-only binding to the contract
	IEntryPointFilterer   // Log filterer for contract events
}

// IEntryPointCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEntryPointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntryPointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEntryPointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntryPointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEntryPointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntryPointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEntryPointSession struct {
	Contract     *IEntryPoint      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEntryPointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEntryPointCallerSession struct {
	Contract *IEntryPointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IEntryPointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEntryPointTransactorSession struct {
	Contract     *IEntryPointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IEntryPointRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEntryPointRaw struct {
	Contract *IEntryPoint // Generic contract binding to access the raw methods on
}

// IEntryPointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEntryPointCallerRaw struct {
	Contract *IEntryPointCaller // Generic read-only contract binding to access the raw methods on
}

// IEntryPointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEntryPointTransactorRaw struct {
	Contract *IEntryPointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEntryPoint creates a new instance of IEntryPoint, bound to a specific deployed contract.
func NewIEntryPoint(address common.Address, backend bind.ContractBackend) (*IEntryPoint, error) {
	contract, err := bindIEntryPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEntryPoint{IEntryPointCaller: IEntryPointCaller{contract: contract}, IEntryPointTransactor: IEntryPointTransactor{contract: contract}, IEntryPointFilterer: IEntryPointFilterer{contract: contract}}, nil
}

// NewIEntryPointCaller creates a new read-only instance of IEntryPoint, bound to a specific deployed contract.
func NewIEntryPointCaller(address common.Address, caller bind.ContractCaller) (*IEntryPointCaller, error) {
	contract, err := bindIEntryPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEntryPointCaller{contract: contract}, nil
}

// NewIEntryPointTransactor creates a new write-only instance of IEntryPoint, bound to a specific deployed contract.
func NewIEntryPointTransactor(address common.Address, transactor bind.ContractTransactor) (*IEntryPointTransactor, error) {
	contract, err := bindIEntryPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEntryPointTransactor{contract: contract}, nil
}

// NewIEntryPointFilterer creates a new log filterer instance of IEntryPoint, bound to a specific deployed contract.
func NewIEntryPointFilterer(address common.Address, filterer bind.ContractFilterer) (*IEntryPointFilterer, error) {
	contract, err := bindIEntryPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEntryPointFilterer{contract: contract}, nil
}

// bindIEntryPoint binds a generic wrapper to an already deployed contract.
func bindIEntryPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEntryPointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEntryPoint *IEntryPointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEntryPoint.Contract.IEntryPointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEntryPoint *IEntryPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEntryPoint.Contract.IEntryPointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEntryPoint *IEntryPointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEntryPoint.Contract.IEntryPointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEntryPoint *IEntryPointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEntryPoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEntryPoint *IEntryPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEntryPoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEntryPoint *IEntryPointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEntryPoint.Contract.contract.Transact(opts, method, params...)
}

// EstimateCrossMessageParamsCrossGas is a free data retrieval call binding the contract method 0xb2241600.
//
// Solidity: function estimateCrossMessageParamsCrossGas(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256)
func (_IEntryPoint *IEntryPointCaller) EstimateCrossMessageParamsCrossGas(opts *bind.CallOpts, params BaseStructCrossMessageParams) (*big.Int, error) {
	var out []interface{}
	err := _IEntryPoint.contract.Call(opts, &out, "estimateCrossMessageParamsCrossGas", params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCrossMessageParamsCrossGas is a free data retrieval call binding the contract method 0xb2241600.
//
// Solidity: function estimateCrossMessageParamsCrossGas(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256)
func (_IEntryPoint *IEntryPointSession) EstimateCrossMessageParamsCrossGas(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _IEntryPoint.Contract.EstimateCrossMessageParamsCrossGas(&_IEntryPoint.CallOpts, params)
}

// EstimateCrossMessageParamsCrossGas is a free data retrieval call binding the contract method 0xb2241600.
//
// Solidity: function estimateCrossMessageParamsCrossGas(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256)
func (_IEntryPoint *IEntryPointCallerSession) EstimateCrossMessageParamsCrossGas(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _IEntryPoint.Contract.EstimateCrossMessageParamsCrossGas(&_IEntryPoint.CallOpts, params)
}

// GetChainConfigs is a free data retrieval call binding the contract method 0x42f1ec69.
//
// Solidity: function getChainConfigs(uint64 chainId) view returns((address,address))
func (_IEntryPoint *IEntryPointCaller) GetChainConfigs(opts *bind.CallOpts, chainId uint64) (IConfigManagerConfig, error) {
	var out []interface{}
	err := _IEntryPoint.contract.Call(opts, &out, "getChainConfigs", chainId)

	if err != nil {
		return *new(IConfigManagerConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IConfigManagerConfig)).(*IConfigManagerConfig)

	return out0, err

}

// GetChainConfigs is a free data retrieval call binding the contract method 0x42f1ec69.
//
// Solidity: function getChainConfigs(uint64 chainId) view returns((address,address))
func (_IEntryPoint *IEntryPointSession) GetChainConfigs(chainId uint64) (IConfigManagerConfig, error) {
	return _IEntryPoint.Contract.GetChainConfigs(&_IEntryPoint.CallOpts, chainId)
}

// GetChainConfigs is a free data retrieval call binding the contract method 0x42f1ec69.
//
// Solidity: function getChainConfigs(uint64 chainId) view returns((address,address))
func (_IEntryPoint *IEntryPointCallerSession) GetChainConfigs(chainId uint64) (IConfigManagerConfig, error) {
	return _IEntryPoint.Contract.GetChainConfigs(&_IEntryPoint.CallOpts, chainId)
}

// GetMainChainId is a free data retrieval call binding the contract method 0x750c67ba.
//
// Solidity: function getMainChainId() view returns(uint64)
func (_IEntryPoint *IEntryPointCaller) GetMainChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEntryPoint.contract.Call(opts, &out, "getMainChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetMainChainId is a free data retrieval call binding the contract method 0x750c67ba.
//
// Solidity: function getMainChainId() view returns(uint64)
func (_IEntryPoint *IEntryPointSession) GetMainChainId() (uint64, error) {
	return _IEntryPoint.Contract.GetMainChainId(&_IEntryPoint.CallOpts)
}

// GetMainChainId is a free data retrieval call binding the contract method 0x750c67ba.
//
// Solidity: function getMainChainId() view returns(uint64)
func (_IEntryPoint *IEntryPointCallerSession) GetMainChainId() (uint64, error) {
	return _IEntryPoint.Contract.GetMainChainId(&_IEntryPoint.CallOpts)
}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_IEntryPoint *IEntryPointCaller) GetPreGasBalanceInfo(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IEntryPoint.contract.Call(opts, &out, "getPreGasBalanceInfo", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_IEntryPoint *IEntryPointSession) GetPreGasBalanceInfo(account common.Address) (*big.Int, error) {
	return _IEntryPoint.Contract.GetPreGasBalanceInfo(&_IEntryPoint.CallOpts, account)
}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_IEntryPoint *IEntryPointCallerSession) GetPreGasBalanceInfo(account common.Address) (*big.Int, error) {
	return _IEntryPoint.Contract.GetPreGasBalanceInfo(&_IEntryPoint.CallOpts, account)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x2885a84f.
//
// Solidity: function getUserOpHash((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) view returns(bytes32)
func (_IEntryPoint *IEntryPointCaller) GetUserOpHash(opts *bind.CallOpts, userOp BaseStructPackedUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _IEntryPoint.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0x2885a84f.
//
// Solidity: function getUserOpHash((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) view returns(bytes32)
func (_IEntryPoint *IEntryPointSession) GetUserOpHash(userOp BaseStructPackedUserOperation) ([32]byte, error) {
	return _IEntryPoint.Contract.GetUserOpHash(&_IEntryPoint.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0x2885a84f.
//
// Solidity: function getUserOpHash((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp) view returns(bytes32)
func (_IEntryPoint *IEntryPointCallerSession) GetUserOpHash(userOp BaseStructPackedUserOperation) ([32]byte, error) {
	return _IEntryPoint.Contract.GetUserOpHash(&_IEntryPoint.CallOpts, userOp)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_IEntryPoint *IEntryPointTransactor) DelegateAndRevert(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _IEntryPoint.contract.Transact(opts, "delegateAndRevert", target, data)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_IEntryPoint *IEntryPointSession) DelegateAndRevert(target common.Address, data []byte) (*types.Transaction, error) {
	return _IEntryPoint.Contract.DelegateAndRevert(&_IEntryPoint.TransactOpts, target, data)
}

// DelegateAndRevert is a paid mutator transaction binding the contract method 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (_IEntryPoint *IEntryPointTransactorSession) DelegateAndRevert(target common.Address, data []byte) (*types.Transaction, error) {
	return _IEntryPoint.Contract.DelegateAndRevert(&_IEntryPoint.TransactOpts, target, data)
}

// HandleOps is a paid mutator transaction binding the contract method 0x77627f0c.
//
// Solidity: function handleOps((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] ops, address beneficiary, bool isSync) returns()
func (_IEntryPoint *IEntryPointTransactor) HandleOps(opts *bind.TransactOpts, ops []BaseStructPackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _IEntryPoint.contract.Transact(opts, "handleOps", ops, beneficiary, isSync)
}

// HandleOps is a paid mutator transaction binding the contract method 0x77627f0c.
//
// Solidity: function handleOps((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] ops, address beneficiary, bool isSync) returns()
func (_IEntryPoint *IEntryPointSession) HandleOps(ops []BaseStructPackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _IEntryPoint.Contract.HandleOps(&_IEntryPoint.TransactOpts, ops, beneficiary, isSync)
}

// HandleOps is a paid mutator transaction binding the contract method 0x77627f0c.
//
// Solidity: function handleOps((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] ops, address beneficiary, bool isSync) returns()
func (_IEntryPoint *IEntryPointTransactorSession) HandleOps(ops []BaseStructPackedUserOperation, beneficiary common.Address, isSync bool) (*types.Transaction, error) {
	return _IEntryPoint.Contract.HandleOps(&_IEntryPoint.TransactOpts, ops, beneficiary, isSync)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_IEntryPoint *IEntryPointTransactor) SendUserOmniMessage(opts *bind.TransactOpts, params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _IEntryPoint.contract.Transact(opts, "sendUserOmniMessage", params)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_IEntryPoint *IEntryPointSession) SendUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _IEntryPoint.Contract.SendUserOmniMessage(&_IEntryPoint.TransactOpts, params)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_IEntryPoint *IEntryPointTransactorSession) SendUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _IEntryPoint.Contract.SendUserOmniMessage(&_IEntryPoint.TransactOpts, params)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_IEntryPoint *IEntryPointTransactor) SubmitDepositOperation(opts *bind.TransactOpts, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IEntryPoint.contract.Transact(opts, "submitDepositOperation", amount, nonce)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_IEntryPoint *IEntryPointSession) SubmitDepositOperation(amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IEntryPoint.Contract.SubmitDepositOperation(&_IEntryPoint.TransactOpts, amount, nonce)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_IEntryPoint *IEntryPointTransactorSession) SubmitDepositOperation(amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IEntryPoint.Contract.SubmitDepositOperation(&_IEntryPoint.TransactOpts, amount, nonce)
}

// SubmitDepositOperationByRemote is a paid mutator transaction binding the contract method 0xbb8ce526.
//
// Solidity: function submitDepositOperationByRemote(address sender, uint256 amount, uint256 nonce) payable returns()
func (_IEntryPoint *IEntryPointTransactor) SubmitDepositOperationByRemote(opts *bind.TransactOpts, sender common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IEntryPoint.contract.Transact(opts, "submitDepositOperationByRemote", sender, amount, nonce)
}

// SubmitDepositOperationByRemote is a paid mutator transaction binding the contract method 0xbb8ce526.
//
// Solidity: function submitDepositOperationByRemote(address sender, uint256 amount, uint256 nonce) payable returns()
func (_IEntryPoint *IEntryPointSession) SubmitDepositOperationByRemote(sender common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IEntryPoint.Contract.SubmitDepositOperationByRemote(&_IEntryPoint.TransactOpts, sender, amount, nonce)
}

// SubmitDepositOperationByRemote is a paid mutator transaction binding the contract method 0xbb8ce526.
//
// Solidity: function submitDepositOperationByRemote(address sender, uint256 amount, uint256 nonce) payable returns()
func (_IEntryPoint *IEntryPointTransactorSession) SubmitDepositOperationByRemote(sender common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IEntryPoint.Contract.SubmitDepositOperationByRemote(&_IEntryPoint.TransactOpts, sender, amount, nonce)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_IEntryPoint *IEntryPointTransactor) SubmitWithdrawOperation(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IEntryPoint.contract.Transact(opts, "submitWithdrawOperation", amount)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_IEntryPoint *IEntryPointSession) SubmitWithdrawOperation(amount *big.Int) (*types.Transaction, error) {
	return _IEntryPoint.Contract.SubmitWithdrawOperation(&_IEntryPoint.TransactOpts, amount)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_IEntryPoint *IEntryPointTransactorSession) SubmitWithdrawOperation(amount *big.Int) (*types.Transaction, error) {
	return _IEntryPoint.Contract.SubmitWithdrawOperation(&_IEntryPoint.TransactOpts, amount)
}

// SyncBatches is a paid mutator transaction binding the contract method 0x77c2719c.
//
// Solidity: function syncBatches((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOps) returns()
func (_IEntryPoint *IEntryPointTransactor) SyncBatches(opts *bind.TransactOpts, userOps []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _IEntryPoint.contract.Transact(opts, "syncBatches", userOps)
}

// SyncBatches is a paid mutator transaction binding the contract method 0x77c2719c.
//
// Solidity: function syncBatches((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOps) returns()
func (_IEntryPoint *IEntryPointSession) SyncBatches(userOps []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _IEntryPoint.Contract.SyncBatches(&_IEntryPoint.TransactOpts, userOps)
}

// SyncBatches is a paid mutator transaction binding the contract method 0x77c2719c.
//
// Solidity: function syncBatches((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOps) returns()
func (_IEntryPoint *IEntryPointTransactorSession) SyncBatches(userOps []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _IEntryPoint.Contract.SyncBatches(&_IEntryPoint.TransactOpts, userOps)
}

// IEntryPointAccountDeployedIterator is returned from FilterAccountDeployed and is used to iterate over the raw logs and unpacked data for AccountDeployed events raised by the IEntryPoint contract.
type IEntryPointAccountDeployedIterator struct {
	Event *IEntryPointAccountDeployed // Event containing the contract specifics and raw log

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
func (it *IEntryPointAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointAccountDeployed)
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
		it.Event = new(IEntryPointAccountDeployed)
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
func (it *IEntryPointAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointAccountDeployed represents a AccountDeployed event raised by the IEntryPoint contract.
type IEntryPointAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDeployed is a free log retrieval operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_IEntryPoint *IEntryPointFilterer) FilterAccountDeployed(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*IEntryPointAccountDeployedIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPoint.contract.FilterLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointAccountDeployedIterator{contract: _IEntryPoint.contract, event: "AccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAccountDeployed is a free log subscription operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_IEntryPoint *IEntryPointFilterer) WatchAccountDeployed(opts *bind.WatchOpts, sink chan<- *IEntryPointAccountDeployed, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPoint.contract.WatchLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointAccountDeployed)
				if err := _IEntryPoint.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
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
func (_IEntryPoint *IEntryPointFilterer) ParseAccountDeployed(log types.Log) (*IEntryPointAccountDeployed, error) {
	event := new(IEntryPointAccountDeployed)
	if err := _IEntryPoint.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointBeforeExecutionIterator is returned from FilterBeforeExecution and is used to iterate over the raw logs and unpacked data for BeforeExecution events raised by the IEntryPoint contract.
type IEntryPointBeforeExecutionIterator struct {
	Event *IEntryPointBeforeExecution // Event containing the contract specifics and raw log

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
func (it *IEntryPointBeforeExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointBeforeExecution)
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
		it.Event = new(IEntryPointBeforeExecution)
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
func (it *IEntryPointBeforeExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointBeforeExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointBeforeExecution represents a BeforeExecution event raised by the IEntryPoint contract.
type IEntryPointBeforeExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBeforeExecution is a free log retrieval operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_IEntryPoint *IEntryPointFilterer) FilterBeforeExecution(opts *bind.FilterOpts) (*IEntryPointBeforeExecutionIterator, error) {

	logs, sub, err := _IEntryPoint.contract.FilterLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return &IEntryPointBeforeExecutionIterator{contract: _IEntryPoint.contract, event: "BeforeExecution", logs: logs, sub: sub}, nil
}

// WatchBeforeExecution is a free log subscription operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_IEntryPoint *IEntryPointFilterer) WatchBeforeExecution(opts *bind.WatchOpts, sink chan<- *IEntryPointBeforeExecution) (event.Subscription, error) {

	logs, sub, err := _IEntryPoint.contract.WatchLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointBeforeExecution)
				if err := _IEntryPoint.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
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
func (_IEntryPoint *IEntryPointFilterer) ParseBeforeExecution(log types.Log) (*IEntryPointBeforeExecution, error) {
	event := new(IEntryPointBeforeExecution)
	if err := _IEntryPoint.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointDepositTicketAddedIterator is returned from FilterDepositTicketAdded and is used to iterate over the raw logs and unpacked data for DepositTicketAdded events raised by the IEntryPoint contract.
type IEntryPointDepositTicketAddedIterator struct {
	Event *IEntryPointDepositTicketAdded // Event containing the contract specifics and raw log

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
func (it *IEntryPointDepositTicketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointDepositTicketAdded)
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
		it.Event = new(IEntryPointDepositTicketAdded)
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
func (it *IEntryPointDepositTicketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointDepositTicketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointDepositTicketAdded represents a DepositTicketAdded event raised by the IEntryPoint contract.
type IEntryPointDepositTicketAdded struct {
	Did       [32]byte
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketAdded is a free log retrieval operation binding the contract event 0x081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc.
//
// Solidity: event DepositTicketAdded(bytes32 indexed did, address indexed account, uint256 amount, uint256 timestamp)
func (_IEntryPoint *IEntryPointFilterer) FilterDepositTicketAdded(opts *bind.FilterOpts, did [][32]byte, account []common.Address) (*IEntryPointDepositTicketAddedIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPoint.contract.FilterLogs(opts, "DepositTicketAdded", didRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointDepositTicketAddedIterator{contract: _IEntryPoint.contract, event: "DepositTicketAdded", logs: logs, sub: sub}, nil
}

// WatchDepositTicketAdded is a free log subscription operation binding the contract event 0x081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc.
//
// Solidity: event DepositTicketAdded(bytes32 indexed did, address indexed account, uint256 amount, uint256 timestamp)
func (_IEntryPoint *IEntryPointFilterer) WatchDepositTicketAdded(opts *bind.WatchOpts, sink chan<- *IEntryPointDepositTicketAdded, did [][32]byte, account []common.Address) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPoint.contract.WatchLogs(opts, "DepositTicketAdded", didRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointDepositTicketAdded)
				if err := _IEntryPoint.contract.UnpackLog(event, "DepositTicketAdded", log); err != nil {
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
func (_IEntryPoint *IEntryPointFilterer) ParseDepositTicketAdded(log types.Log) (*IEntryPointDepositTicketAdded, error) {
	event := new(IEntryPointDepositTicketAdded)
	if err := _IEntryPoint.contract.UnpackLog(event, "DepositTicketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointDepositTicketDeletedIterator is returned from FilterDepositTicketDeleted and is used to iterate over the raw logs and unpacked data for DepositTicketDeleted events raised by the IEntryPoint contract.
type IEntryPointDepositTicketDeletedIterator struct {
	Event *IEntryPointDepositTicketDeleted // Event containing the contract specifics and raw log

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
func (it *IEntryPointDepositTicketDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointDepositTicketDeleted)
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
		it.Event = new(IEntryPointDepositTicketDeleted)
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
func (it *IEntryPointDepositTicketDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointDepositTicketDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointDepositTicketDeleted represents a DepositTicketDeleted event raised by the IEntryPoint contract.
type IEntryPointDepositTicketDeleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketDeleted is a free log retrieval operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_IEntryPoint *IEntryPointFilterer) FilterDepositTicketDeleted(opts *bind.FilterOpts, account []common.Address) (*IEntryPointDepositTicketDeletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPoint.contract.FilterLogs(opts, "DepositTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointDepositTicketDeletedIterator{contract: _IEntryPoint.contract, event: "DepositTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchDepositTicketDeleted is a free log subscription operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_IEntryPoint *IEntryPointFilterer) WatchDepositTicketDeleted(opts *bind.WatchOpts, sink chan<- *IEntryPointDepositTicketDeleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPoint.contract.WatchLogs(opts, "DepositTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointDepositTicketDeleted)
				if err := _IEntryPoint.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
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
func (_IEntryPoint *IEntryPointFilterer) ParseDepositTicketDeleted(log types.Log) (*IEntryPointDepositTicketDeleted, error) {
	event := new(IEntryPointDepositTicketDeleted)
	if err := _IEntryPoint.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointPostOpRevertReasonIterator is returned from FilterPostOpRevertReason and is used to iterate over the raw logs and unpacked data for PostOpRevertReason events raised by the IEntryPoint contract.
type IEntryPointPostOpRevertReasonIterator struct {
	Event *IEntryPointPostOpRevertReason // Event containing the contract specifics and raw log

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
func (it *IEntryPointPostOpRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointPostOpRevertReason)
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
		it.Event = new(IEntryPointPostOpRevertReason)
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
func (it *IEntryPointPostOpRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointPostOpRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointPostOpRevertReason represents a PostOpRevertReason event raised by the IEntryPoint contract.
type IEntryPointPostOpRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostOpRevertReason is a free log retrieval operation binding the contract event 0x676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_IEntryPoint *IEntryPointFilterer) FilterPostOpRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*IEntryPointPostOpRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPoint.contract.FilterLogs(opts, "PostOpRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointPostOpRevertReasonIterator{contract: _IEntryPoint.contract, event: "PostOpRevertReason", logs: logs, sub: sub}, nil
}

// WatchPostOpRevertReason is a free log subscription operation binding the contract event 0x676a9eebb6acdde3e7a8a121f493aa3c68d165eb138523bfc9f733b566bacd9e.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_IEntryPoint *IEntryPointFilterer) WatchPostOpRevertReason(opts *bind.WatchOpts, sink chan<- *IEntryPointPostOpRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPoint.contract.WatchLogs(opts, "PostOpRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointPostOpRevertReason)
				if err := _IEntryPoint.contract.UnpackLog(event, "PostOpRevertReason", log); err != nil {
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
func (_IEntryPoint *IEntryPointFilterer) ParsePostOpRevertReason(log types.Log) (*IEntryPointPostOpRevertReason, error) {
	event := new(IEntryPointPostOpRevertReason)
	if err := _IEntryPoint.contract.UnpackLog(event, "PostOpRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointUserOperationEventIterator is returned from FilterUserOperationEvent and is used to iterate over the raw logs and unpacked data for UserOperationEvent events raised by the IEntryPoint contract.
type IEntryPointUserOperationEventIterator struct {
	Event *IEntryPointUserOperationEvent // Event containing the contract specifics and raw log

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
func (it *IEntryPointUserOperationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointUserOperationEvent)
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
		it.Event = new(IEntryPointUserOperationEvent)
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
func (it *IEntryPointUserOperationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointUserOperationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointUserOperationEvent represents a UserOperationEvent event raised by the IEntryPoint contract.
type IEntryPointUserOperationEvent struct {
	UserOpHash    [32]byte
	Sender        common.Address
	Owner         common.Address
	Phase         uint8
	InnerExec     bool
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserOperationEvent is a free log retrieval operation binding the contract event 0x6072b7fd64d9718ff132984edd6caa7702fcca448873489af52aba3349477f33.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address owner, uint8 phase, bool innerExec, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_IEntryPoint *IEntryPointFilterer) FilterUserOperationEvent(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*IEntryPointUserOperationEventIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPoint.contract.FilterLogs(opts, "UserOperationEvent", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointUserOperationEventIterator{contract: _IEntryPoint.contract, event: "UserOperationEvent", logs: logs, sub: sub}, nil
}

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0x6072b7fd64d9718ff132984edd6caa7702fcca448873489af52aba3349477f33.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address owner, uint8 phase, bool innerExec, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_IEntryPoint *IEntryPointFilterer) WatchUserOperationEvent(opts *bind.WatchOpts, sink chan<- *IEntryPointUserOperationEvent, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPoint.contract.WatchLogs(opts, "UserOperationEvent", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointUserOperationEvent)
				if err := _IEntryPoint.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
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

// ParseUserOperationEvent is a log parse operation binding the contract event 0x6072b7fd64d9718ff132984edd6caa7702fcca448873489af52aba3349477f33.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address owner, uint8 phase, bool innerExec, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_IEntryPoint *IEntryPointFilterer) ParseUserOperationEvent(log types.Log) (*IEntryPointUserOperationEvent, error) {
	event := new(IEntryPointUserOperationEvent)
	if err := _IEntryPoint.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointUserOperationPrefundTooLowIterator is returned from FilterUserOperationPrefundTooLow and is used to iterate over the raw logs and unpacked data for UserOperationPrefundTooLow events raised by the IEntryPoint contract.
type IEntryPointUserOperationPrefundTooLowIterator struct {
	Event *IEntryPointUserOperationPrefundTooLow // Event containing the contract specifics and raw log

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
func (it *IEntryPointUserOperationPrefundTooLowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointUserOperationPrefundTooLow)
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
		it.Event = new(IEntryPointUserOperationPrefundTooLow)
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
func (it *IEntryPointUserOperationPrefundTooLowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointUserOperationPrefundTooLowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointUserOperationPrefundTooLow represents a UserOperationPrefundTooLow event raised by the IEntryPoint contract.
type IEntryPointUserOperationPrefundTooLow struct {
	UserOpHash [32]byte
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserOperationPrefundTooLow is a free log retrieval operation binding the contract event 0x1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff28.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender)
func (_IEntryPoint *IEntryPointFilterer) FilterUserOperationPrefundTooLow(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*IEntryPointUserOperationPrefundTooLowIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPoint.contract.FilterLogs(opts, "UserOperationPrefundTooLow", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointUserOperationPrefundTooLowIterator{contract: _IEntryPoint.contract, event: "UserOperationPrefundTooLow", logs: logs, sub: sub}, nil
}

// WatchUserOperationPrefundTooLow is a free log subscription operation binding the contract event 0x1eefc21737217315131acde0288642f4328ead1376833b69db86aa68812bff28.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender)
func (_IEntryPoint *IEntryPointFilterer) WatchUserOperationPrefundTooLow(opts *bind.WatchOpts, sink chan<- *IEntryPointUserOperationPrefundTooLow, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPoint.contract.WatchLogs(opts, "UserOperationPrefundTooLow", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointUserOperationPrefundTooLow)
				if err := _IEntryPoint.contract.UnpackLog(event, "UserOperationPrefundTooLow", log); err != nil {
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
func (_IEntryPoint *IEntryPointFilterer) ParseUserOperationPrefundTooLow(log types.Log) (*IEntryPointUserOperationPrefundTooLow, error) {
	event := new(IEntryPointUserOperationPrefundTooLow)
	if err := _IEntryPoint.contract.UnpackLog(event, "UserOperationPrefundTooLow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointUserOperationRevertReasonIterator is returned from FilterUserOperationRevertReason and is used to iterate over the raw logs and unpacked data for UserOperationRevertReason events raised by the IEntryPoint contract.
type IEntryPointUserOperationRevertReasonIterator struct {
	Event *IEntryPointUserOperationRevertReason // Event containing the contract specifics and raw log

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
func (it *IEntryPointUserOperationRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointUserOperationRevertReason)
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
		it.Event = new(IEntryPointUserOperationRevertReason)
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
func (it *IEntryPointUserOperationRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointUserOperationRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointUserOperationRevertReason represents a UserOperationRevertReason event raised by the IEntryPoint contract.
type IEntryPointUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserOperationRevertReason is a free log retrieval operation binding the contract event 0x0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_IEntryPoint *IEntryPointFilterer) FilterUserOperationRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*IEntryPointUserOperationRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPoint.contract.FilterLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointUserOperationRevertReasonIterator{contract: _IEntryPoint.contract, event: "UserOperationRevertReason", logs: logs, sub: sub}, nil
}

// WatchUserOperationRevertReason is a free log subscription operation binding the contract event 0x0c99075125a36c86fcce1c24656ca4511d30e30babceee95632cf3e66d3d1beb.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, bytes revertReason)
func (_IEntryPoint *IEntryPointFilterer) WatchUserOperationRevertReason(opts *bind.WatchOpts, sink chan<- *IEntryPointUserOperationRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IEntryPoint.contract.WatchLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointUserOperationRevertReason)
				if err := _IEntryPoint.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
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
func (_IEntryPoint *IEntryPointFilterer) ParseUserOperationRevertReason(log types.Log) (*IEntryPointUserOperationRevertReason, error) {
	event := new(IEntryPointUserOperationRevertReason)
	if err := _IEntryPoint.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointWithdrawTicketAddedIterator is returned from FilterWithdrawTicketAdded and is used to iterate over the raw logs and unpacked data for WithdrawTicketAdded events raised by the IEntryPoint contract.
type IEntryPointWithdrawTicketAddedIterator struct {
	Event *IEntryPointWithdrawTicketAdded // Event containing the contract specifics and raw log

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
func (it *IEntryPointWithdrawTicketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointWithdrawTicketAdded)
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
		it.Event = new(IEntryPointWithdrawTicketAdded)
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
func (it *IEntryPointWithdrawTicketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointWithdrawTicketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointWithdrawTicketAdded represents a WithdrawTicketAdded event raised by the IEntryPoint contract.
type IEntryPointWithdrawTicketAdded struct {
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketAdded is a free log retrieval operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_IEntryPoint *IEntryPointFilterer) FilterWithdrawTicketAdded(opts *bind.FilterOpts, account []common.Address) (*IEntryPointWithdrawTicketAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPoint.contract.FilterLogs(opts, "WithdrawTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointWithdrawTicketAddedIterator{contract: _IEntryPoint.contract, event: "WithdrawTicketAdded", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketAdded is a free log subscription operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_IEntryPoint *IEntryPointFilterer) WatchWithdrawTicketAdded(opts *bind.WatchOpts, sink chan<- *IEntryPointWithdrawTicketAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPoint.contract.WatchLogs(opts, "WithdrawTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointWithdrawTicketAdded)
				if err := _IEntryPoint.contract.UnpackLog(event, "WithdrawTicketAdded", log); err != nil {
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
func (_IEntryPoint *IEntryPointFilterer) ParseWithdrawTicketAdded(log types.Log) (*IEntryPointWithdrawTicketAdded, error) {
	event := new(IEntryPointWithdrawTicketAdded)
	if err := _IEntryPoint.contract.UnpackLog(event, "WithdrawTicketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEntryPointWithdrawTicketDeletedIterator is returned from FilterWithdrawTicketDeleted and is used to iterate over the raw logs and unpacked data for WithdrawTicketDeleted events raised by the IEntryPoint contract.
type IEntryPointWithdrawTicketDeletedIterator struct {
	Event *IEntryPointWithdrawTicketDeleted // Event containing the contract specifics and raw log

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
func (it *IEntryPointWithdrawTicketDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEntryPointWithdrawTicketDeleted)
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
		it.Event = new(IEntryPointWithdrawTicketDeleted)
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
func (it *IEntryPointWithdrawTicketDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEntryPointWithdrawTicketDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEntryPointWithdrawTicketDeleted represents a WithdrawTicketDeleted event raised by the IEntryPoint contract.
type IEntryPointWithdrawTicketDeleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketDeleted is a free log retrieval operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_IEntryPoint *IEntryPointFilterer) FilterWithdrawTicketDeleted(opts *bind.FilterOpts, account []common.Address) (*IEntryPointWithdrawTicketDeletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPoint.contract.FilterLogs(opts, "WithdrawTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEntryPointWithdrawTicketDeletedIterator{contract: _IEntryPoint.contract, event: "WithdrawTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketDeleted is a free log subscription operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_IEntryPoint *IEntryPointFilterer) WatchWithdrawTicketDeleted(opts *bind.WatchOpts, sink chan<- *IEntryPointWithdrawTicketDeleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEntryPoint.contract.WatchLogs(opts, "WithdrawTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEntryPointWithdrawTicketDeleted)
				if err := _IEntryPoint.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
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
func (_IEntryPoint *IEntryPointFilterer) ParseWithdrawTicketDeleted(log types.Log) (*IEntryPointWithdrawTicketDeleted, error) {
	event := new(IEntryPointWithdrawTicketDeleted)
	if err := _IEntryPoint.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymasterMetaData contains all meta data concerning the IPaymaster contract.
var IPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use IPaymasterMetaData.ABI instead.
var IPaymasterABI = IPaymasterMetaData.ABI

// IPaymaster is an auto generated Go binding around an Ethereum contract.
type IPaymaster struct {
	IPaymasterCaller     // Read-only binding to the contract
	IPaymasterTransactor // Write-only binding to the contract
	IPaymasterFilterer   // Log filterer for contract events
}

// IPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPaymasterSession struct {
	Contract     *IPaymaster       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPaymasterCallerSession struct {
	Contract *IPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPaymasterTransactorSession struct {
	Contract     *IPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPaymasterRaw struct {
	Contract *IPaymaster // Generic contract binding to access the raw methods on
}

// IPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPaymasterCallerRaw struct {
	Contract *IPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// IPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPaymasterTransactorRaw struct {
	Contract *IPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPaymaster creates a new instance of IPaymaster, bound to a specific deployed contract.
func NewIPaymaster(address common.Address, backend bind.ContractBackend) (*IPaymaster, error) {
	contract, err := bindIPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPaymaster{IPaymasterCaller: IPaymasterCaller{contract: contract}, IPaymasterTransactor: IPaymasterTransactor{contract: contract}, IPaymasterFilterer: IPaymasterFilterer{contract: contract}}, nil
}

// NewIPaymasterCaller creates a new read-only instance of IPaymaster, bound to a specific deployed contract.
func NewIPaymasterCaller(address common.Address, caller bind.ContractCaller) (*IPaymasterCaller, error) {
	contract, err := bindIPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPaymasterCaller{contract: contract}, nil
}

// NewIPaymasterTransactor creates a new write-only instance of IPaymaster, bound to a specific deployed contract.
func NewIPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*IPaymasterTransactor, error) {
	contract, err := bindIPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPaymasterTransactor{contract: contract}, nil
}

// NewIPaymasterFilterer creates a new log filterer instance of IPaymaster, bound to a specific deployed contract.
func NewIPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*IPaymasterFilterer, error) {
	contract, err := bindIPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPaymasterFilterer{contract: contract}, nil
}

// bindIPaymaster binds a generic wrapper to an already deployed contract.
func bindIPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPaymaster *IPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPaymaster.Contract.IPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPaymaster *IPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPaymaster.Contract.IPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPaymaster *IPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPaymaster.Contract.IPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPaymaster *IPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPaymaster *IPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPaymaster *IPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPaymaster.Contract.contract.Transact(opts, method, params...)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_IPaymaster *IPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _IPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_IPaymaster *IPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _IPaymaster.Contract.PostOp(&_IPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_IPaymaster *IPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _IPaymaster.Contract.PostOp(&_IPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0xa2365dde.
//
// Solidity: function validatePaymasterUserOp((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_IPaymaster *IPaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp BaseStructPackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _IPaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0xa2365dde.
//
// Solidity: function validatePaymasterUserOp((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_IPaymaster *IPaymasterSession) ValidatePaymasterUserOp(userOp BaseStructPackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _IPaymaster.Contract.ValidatePaymasterUserOp(&_IPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0xa2365dde.
//
// Solidity: function validatePaymasterUserOp((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes)) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_IPaymaster *IPaymasterTransactorSession) ValidatePaymasterUserOp(userOp BaseStructPackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _IPaymaster.Contract.ValidatePaymasterUserOp(&_IPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// IPreGasManagerMetaData contains all meta data concerning the IPreGasManager contract.
var IPreGasManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotEqual\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"did\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPreGasBalanceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitDepositOperation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submitWithdrawOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPreGasManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IPreGasManagerMetaData.ABI instead.
var IPreGasManagerABI = IPreGasManagerMetaData.ABI

// IPreGasManager is an auto generated Go binding around an Ethereum contract.
type IPreGasManager struct {
	IPreGasManagerCaller     // Read-only binding to the contract
	IPreGasManagerTransactor // Write-only binding to the contract
	IPreGasManagerFilterer   // Log filterer for contract events
}

// IPreGasManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPreGasManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPreGasManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPreGasManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPreGasManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPreGasManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPreGasManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPreGasManagerSession struct {
	Contract     *IPreGasManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPreGasManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPreGasManagerCallerSession struct {
	Contract *IPreGasManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IPreGasManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPreGasManagerTransactorSession struct {
	Contract     *IPreGasManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IPreGasManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPreGasManagerRaw struct {
	Contract *IPreGasManager // Generic contract binding to access the raw methods on
}

// IPreGasManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPreGasManagerCallerRaw struct {
	Contract *IPreGasManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IPreGasManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPreGasManagerTransactorRaw struct {
	Contract *IPreGasManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPreGasManager creates a new instance of IPreGasManager, bound to a specific deployed contract.
func NewIPreGasManager(address common.Address, backend bind.ContractBackend) (*IPreGasManager, error) {
	contract, err := bindIPreGasManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPreGasManager{IPreGasManagerCaller: IPreGasManagerCaller{contract: contract}, IPreGasManagerTransactor: IPreGasManagerTransactor{contract: contract}, IPreGasManagerFilterer: IPreGasManagerFilterer{contract: contract}}, nil
}

// NewIPreGasManagerCaller creates a new read-only instance of IPreGasManager, bound to a specific deployed contract.
func NewIPreGasManagerCaller(address common.Address, caller bind.ContractCaller) (*IPreGasManagerCaller, error) {
	contract, err := bindIPreGasManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPreGasManagerCaller{contract: contract}, nil
}

// NewIPreGasManagerTransactor creates a new write-only instance of IPreGasManager, bound to a specific deployed contract.
func NewIPreGasManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IPreGasManagerTransactor, error) {
	contract, err := bindIPreGasManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPreGasManagerTransactor{contract: contract}, nil
}

// NewIPreGasManagerFilterer creates a new log filterer instance of IPreGasManager, bound to a specific deployed contract.
func NewIPreGasManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IPreGasManagerFilterer, error) {
	contract, err := bindIPreGasManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPreGasManagerFilterer{contract: contract}, nil
}

// bindIPreGasManager binds a generic wrapper to an already deployed contract.
func bindIPreGasManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPreGasManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPreGasManager *IPreGasManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPreGasManager.Contract.IPreGasManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPreGasManager *IPreGasManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPreGasManager.Contract.IPreGasManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPreGasManager *IPreGasManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPreGasManager.Contract.IPreGasManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPreGasManager *IPreGasManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPreGasManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPreGasManager *IPreGasManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPreGasManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPreGasManager *IPreGasManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPreGasManager.Contract.contract.Transact(opts, method, params...)
}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_IPreGasManager *IPreGasManagerCaller) GetPreGasBalanceInfo(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPreGasManager.contract.Call(opts, &out, "getPreGasBalanceInfo", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_IPreGasManager *IPreGasManagerSession) GetPreGasBalanceInfo(account common.Address) (*big.Int, error) {
	return _IPreGasManager.Contract.GetPreGasBalanceInfo(&_IPreGasManager.CallOpts, account)
}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_IPreGasManager *IPreGasManagerCallerSession) GetPreGasBalanceInfo(account common.Address) (*big.Int, error) {
	return _IPreGasManager.Contract.GetPreGasBalanceInfo(&_IPreGasManager.CallOpts, account)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_IPreGasManager *IPreGasManagerTransactor) SubmitDepositOperation(opts *bind.TransactOpts, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IPreGasManager.contract.Transact(opts, "submitDepositOperation", amount, nonce)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_IPreGasManager *IPreGasManagerSession) SubmitDepositOperation(amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IPreGasManager.Contract.SubmitDepositOperation(&_IPreGasManager.TransactOpts, amount, nonce)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_IPreGasManager *IPreGasManagerTransactorSession) SubmitDepositOperation(amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IPreGasManager.Contract.SubmitDepositOperation(&_IPreGasManager.TransactOpts, amount, nonce)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_IPreGasManager *IPreGasManagerTransactor) SubmitWithdrawOperation(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IPreGasManager.contract.Transact(opts, "submitWithdrawOperation", amount)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_IPreGasManager *IPreGasManagerSession) SubmitWithdrawOperation(amount *big.Int) (*types.Transaction, error) {
	return _IPreGasManager.Contract.SubmitWithdrawOperation(&_IPreGasManager.TransactOpts, amount)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_IPreGasManager *IPreGasManagerTransactorSession) SubmitWithdrawOperation(amount *big.Int) (*types.Transaction, error) {
	return _IPreGasManager.Contract.SubmitWithdrawOperation(&_IPreGasManager.TransactOpts, amount)
}

// IPreGasManagerDepositTicketAddedIterator is returned from FilterDepositTicketAdded and is used to iterate over the raw logs and unpacked data for DepositTicketAdded events raised by the IPreGasManager contract.
type IPreGasManagerDepositTicketAddedIterator struct {
	Event *IPreGasManagerDepositTicketAdded // Event containing the contract specifics and raw log

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
func (it *IPreGasManagerDepositTicketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPreGasManagerDepositTicketAdded)
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
		it.Event = new(IPreGasManagerDepositTicketAdded)
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
func (it *IPreGasManagerDepositTicketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPreGasManagerDepositTicketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPreGasManagerDepositTicketAdded represents a DepositTicketAdded event raised by the IPreGasManager contract.
type IPreGasManagerDepositTicketAdded struct {
	Did       [32]byte
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketAdded is a free log retrieval operation binding the contract event 0x081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc.
//
// Solidity: event DepositTicketAdded(bytes32 indexed did, address indexed account, uint256 amount, uint256 timestamp)
func (_IPreGasManager *IPreGasManagerFilterer) FilterDepositTicketAdded(opts *bind.FilterOpts, did [][32]byte, account []common.Address) (*IPreGasManagerDepositTicketAddedIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IPreGasManager.contract.FilterLogs(opts, "DepositTicketAdded", didRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &IPreGasManagerDepositTicketAddedIterator{contract: _IPreGasManager.contract, event: "DepositTicketAdded", logs: logs, sub: sub}, nil
}

// WatchDepositTicketAdded is a free log subscription operation binding the contract event 0x081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc.
//
// Solidity: event DepositTicketAdded(bytes32 indexed did, address indexed account, uint256 amount, uint256 timestamp)
func (_IPreGasManager *IPreGasManagerFilterer) WatchDepositTicketAdded(opts *bind.WatchOpts, sink chan<- *IPreGasManagerDepositTicketAdded, did [][32]byte, account []common.Address) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IPreGasManager.contract.WatchLogs(opts, "DepositTicketAdded", didRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPreGasManagerDepositTicketAdded)
				if err := _IPreGasManager.contract.UnpackLog(event, "DepositTicketAdded", log); err != nil {
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
func (_IPreGasManager *IPreGasManagerFilterer) ParseDepositTicketAdded(log types.Log) (*IPreGasManagerDepositTicketAdded, error) {
	event := new(IPreGasManagerDepositTicketAdded)
	if err := _IPreGasManager.contract.UnpackLog(event, "DepositTicketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPreGasManagerDepositTicketDeletedIterator is returned from FilterDepositTicketDeleted and is used to iterate over the raw logs and unpacked data for DepositTicketDeleted events raised by the IPreGasManager contract.
type IPreGasManagerDepositTicketDeletedIterator struct {
	Event *IPreGasManagerDepositTicketDeleted // Event containing the contract specifics and raw log

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
func (it *IPreGasManagerDepositTicketDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPreGasManagerDepositTicketDeleted)
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
		it.Event = new(IPreGasManagerDepositTicketDeleted)
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
func (it *IPreGasManagerDepositTicketDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPreGasManagerDepositTicketDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPreGasManagerDepositTicketDeleted represents a DepositTicketDeleted event raised by the IPreGasManager contract.
type IPreGasManagerDepositTicketDeleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketDeleted is a free log retrieval operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_IPreGasManager *IPreGasManagerFilterer) FilterDepositTicketDeleted(opts *bind.FilterOpts, account []common.Address) (*IPreGasManagerDepositTicketDeletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IPreGasManager.contract.FilterLogs(opts, "DepositTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &IPreGasManagerDepositTicketDeletedIterator{contract: _IPreGasManager.contract, event: "DepositTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchDepositTicketDeleted is a free log subscription operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_IPreGasManager *IPreGasManagerFilterer) WatchDepositTicketDeleted(opts *bind.WatchOpts, sink chan<- *IPreGasManagerDepositTicketDeleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IPreGasManager.contract.WatchLogs(opts, "DepositTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPreGasManagerDepositTicketDeleted)
				if err := _IPreGasManager.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
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
func (_IPreGasManager *IPreGasManagerFilterer) ParseDepositTicketDeleted(log types.Log) (*IPreGasManagerDepositTicketDeleted, error) {
	event := new(IPreGasManagerDepositTicketDeleted)
	if err := _IPreGasManager.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPreGasManagerWithdrawTicketAddedIterator is returned from FilterWithdrawTicketAdded and is used to iterate over the raw logs and unpacked data for WithdrawTicketAdded events raised by the IPreGasManager contract.
type IPreGasManagerWithdrawTicketAddedIterator struct {
	Event *IPreGasManagerWithdrawTicketAdded // Event containing the contract specifics and raw log

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
func (it *IPreGasManagerWithdrawTicketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPreGasManagerWithdrawTicketAdded)
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
		it.Event = new(IPreGasManagerWithdrawTicketAdded)
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
func (it *IPreGasManagerWithdrawTicketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPreGasManagerWithdrawTicketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPreGasManagerWithdrawTicketAdded represents a WithdrawTicketAdded event raised by the IPreGasManager contract.
type IPreGasManagerWithdrawTicketAdded struct {
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketAdded is a free log retrieval operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_IPreGasManager *IPreGasManagerFilterer) FilterWithdrawTicketAdded(opts *bind.FilterOpts, account []common.Address) (*IPreGasManagerWithdrawTicketAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IPreGasManager.contract.FilterLogs(opts, "WithdrawTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &IPreGasManagerWithdrawTicketAddedIterator{contract: _IPreGasManager.contract, event: "WithdrawTicketAdded", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketAdded is a free log subscription operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_IPreGasManager *IPreGasManagerFilterer) WatchWithdrawTicketAdded(opts *bind.WatchOpts, sink chan<- *IPreGasManagerWithdrawTicketAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IPreGasManager.contract.WatchLogs(opts, "WithdrawTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPreGasManagerWithdrawTicketAdded)
				if err := _IPreGasManager.contract.UnpackLog(event, "WithdrawTicketAdded", log); err != nil {
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
func (_IPreGasManager *IPreGasManagerFilterer) ParseWithdrawTicketAdded(log types.Log) (*IPreGasManagerWithdrawTicketAdded, error) {
	event := new(IPreGasManagerWithdrawTicketAdded)
	if err := _IPreGasManager.contract.UnpackLog(event, "WithdrawTicketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPreGasManagerWithdrawTicketDeletedIterator is returned from FilterWithdrawTicketDeleted and is used to iterate over the raw logs and unpacked data for WithdrawTicketDeleted events raised by the IPreGasManager contract.
type IPreGasManagerWithdrawTicketDeletedIterator struct {
	Event *IPreGasManagerWithdrawTicketDeleted // Event containing the contract specifics and raw log

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
func (it *IPreGasManagerWithdrawTicketDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPreGasManagerWithdrawTicketDeleted)
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
		it.Event = new(IPreGasManagerWithdrawTicketDeleted)
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
func (it *IPreGasManagerWithdrawTicketDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPreGasManagerWithdrawTicketDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPreGasManagerWithdrawTicketDeleted represents a WithdrawTicketDeleted event raised by the IPreGasManager contract.
type IPreGasManagerWithdrawTicketDeleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketDeleted is a free log retrieval operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_IPreGasManager *IPreGasManagerFilterer) FilterWithdrawTicketDeleted(opts *bind.FilterOpts, account []common.Address) (*IPreGasManagerWithdrawTicketDeletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IPreGasManager.contract.FilterLogs(opts, "WithdrawTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &IPreGasManagerWithdrawTicketDeletedIterator{contract: _IPreGasManager.contract, event: "WithdrawTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketDeleted is a free log subscription operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_IPreGasManager *IPreGasManagerFilterer) WatchWithdrawTicketDeleted(opts *bind.WatchOpts, sink chan<- *IPreGasManagerWithdrawTicketDeleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IPreGasManager.contract.WatchLogs(opts, "WithdrawTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPreGasManagerWithdrawTicketDeleted)
				if err := _IPreGasManager.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
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
func (_IPreGasManager *IPreGasManagerFilterer) ParseWithdrawTicketDeleted(log types.Log) (*IPreGasManagerWithdrawTicketDeleted, error) {
	event := new(IPreGasManagerWithdrawTicketDeleted)
	if err := _IPreGasManager.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISyncRouterMetaData contains all meta data concerning the ISyncRouter contract.
var ISyncRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"fetchOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"fetchUserOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainUsedFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"userOperations\",\"type\":\"tuple[]\"}],\"name\":\"sendOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operationValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"exec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkVerificationGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mainChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.ExecData\",\"name\":\"innerExec\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.PackedUserOperation[]\",\"name\":\"_packedUserOperation\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"way\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"gasLimit\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxArrivalTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"selectedRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainExecuteUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"packCrossParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseStruct.CrossHookMessageParams\",\"name\":\"_hookMessageParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseStruct.CrossMessageParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"sendUserOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ISyncRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use ISyncRouterMetaData.ABI instead.
var ISyncRouterABI = ISyncRouterMetaData.ABI

// ISyncRouter is an auto generated Go binding around an Ethereum contract.
type ISyncRouter struct {
	ISyncRouterCaller     // Read-only binding to the contract
	ISyncRouterTransactor // Write-only binding to the contract
	ISyncRouterFilterer   // Log filterer for contract events
}

// ISyncRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISyncRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISyncRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISyncRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISyncRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISyncRouterSession struct {
	Contract     *ISyncRouter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISyncRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISyncRouterCallerSession struct {
	Contract *ISyncRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ISyncRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISyncRouterTransactorSession struct {
	Contract     *ISyncRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ISyncRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISyncRouterRaw struct {
	Contract *ISyncRouter // Generic contract binding to access the raw methods on
}

// ISyncRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISyncRouterCallerRaw struct {
	Contract *ISyncRouterCaller // Generic read-only contract binding to access the raw methods on
}

// ISyncRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISyncRouterTransactorRaw struct {
	Contract *ISyncRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISyncRouter creates a new instance of ISyncRouter, bound to a specific deployed contract.
func NewISyncRouter(address common.Address, backend bind.ContractBackend) (*ISyncRouter, error) {
	contract, err := bindISyncRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISyncRouter{ISyncRouterCaller: ISyncRouterCaller{contract: contract}, ISyncRouterTransactor: ISyncRouterTransactor{contract: contract}, ISyncRouterFilterer: ISyncRouterFilterer{contract: contract}}, nil
}

// NewISyncRouterCaller creates a new read-only instance of ISyncRouter, bound to a specific deployed contract.
func NewISyncRouterCaller(address common.Address, caller bind.ContractCaller) (*ISyncRouterCaller, error) {
	contract, err := bindISyncRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISyncRouterCaller{contract: contract}, nil
}

// NewISyncRouterTransactor creates a new write-only instance of ISyncRouter, bound to a specific deployed contract.
func NewISyncRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ISyncRouterTransactor, error) {
	contract, err := bindISyncRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISyncRouterTransactor{contract: contract}, nil
}

// NewISyncRouterFilterer creates a new log filterer instance of ISyncRouter, bound to a specific deployed contract.
func NewISyncRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ISyncRouterFilterer, error) {
	contract, err := bindISyncRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISyncRouterFilterer{contract: contract}, nil
}

// bindISyncRouter binds a generic wrapper to an already deployed contract.
func bindISyncRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISyncRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISyncRouter *ISyncRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISyncRouter.Contract.ISyncRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISyncRouter *ISyncRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISyncRouter.Contract.ISyncRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISyncRouter *ISyncRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISyncRouter.Contract.ISyncRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISyncRouter *ISyncRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISyncRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISyncRouter *ISyncRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISyncRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISyncRouter *ISyncRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISyncRouter.Contract.contract.Transact(opts, method, params...)
}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0x3eba5864.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) view returns(uint256)
func (_ISyncRouter *ISyncRouterCaller) FetchOmniMessageFee(opts *bind.CallOpts, destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*big.Int, error) {
	var out []interface{}
	err := _ISyncRouter.contract.Call(opts, &out, "fetchOmniMessageFee", destChainId, destContract, destChainUsedFee, userOperations)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0x3eba5864.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) view returns(uint256)
func (_ISyncRouter *ISyncRouterSession) FetchOmniMessageFee(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*big.Int, error) {
	return _ISyncRouter.Contract.FetchOmniMessageFee(&_ISyncRouter.CallOpts, destChainId, destContract, destChainUsedFee, userOperations)
}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0x3eba5864.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) view returns(uint256)
func (_ISyncRouter *ISyncRouterCallerSession) FetchOmniMessageFee(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*big.Int, error) {
	return _ISyncRouter.Contract.FetchOmniMessageFee(&_ISyncRouter.CallOpts, destChainId, destContract, destChainUsedFee, userOperations)
}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0x3169eb13.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256)
func (_ISyncRouter *ISyncRouterCaller) FetchUserOmniMessageFee(opts *bind.CallOpts, params BaseStructCrossMessageParams) (*big.Int, error) {
	var out []interface{}
	err := _ISyncRouter.contract.Call(opts, &out, "fetchUserOmniMessageFee", params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0x3169eb13.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256)
func (_ISyncRouter *ISyncRouterSession) FetchUserOmniMessageFee(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _ISyncRouter.Contract.FetchUserOmniMessageFee(&_ISyncRouter.CallOpts, params)
}

// FetchUserOmniMessageFee is a free data retrieval call binding the contract method 0x3169eb13.
//
// Solidity: function fetchUserOmniMessageFee(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) view returns(uint256)
func (_ISyncRouter *ISyncRouterCallerSession) FetchUserOmniMessageFee(params BaseStructCrossMessageParams) (*big.Int, error) {
	return _ISyncRouter.Contract.FetchUserOmniMessageFee(&_ISyncRouter.CallOpts, params)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0x93afae93.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) payable returns()
func (_ISyncRouter *ISyncRouterTransactor) SendOmniMessage(opts *bind.TransactOpts, destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _ISyncRouter.contract.Transact(opts, "sendOmniMessage", destChainId, destContract, destChainUsedFee, userOperations)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0x93afae93.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) payable returns()
func (_ISyncRouter *ISyncRouterSession) SendOmniMessage(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _ISyncRouter.Contract.SendOmniMessage(&_ISyncRouter.TransactOpts, destChainId, destContract, destChainUsedFee, userOperations)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0x93afae93.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainUsedFee, (uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[] userOperations) payable returns()
func (_ISyncRouter *ISyncRouterTransactorSession) SendOmniMessage(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, userOperations []BaseStructPackedUserOperation) (*types.Transaction, error) {
	return _ISyncRouter.Contract.SendOmniMessage(&_ISyncRouter.TransactOpts, destChainId, destContract, destChainUsedFee, userOperations)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_ISyncRouter *ISyncRouterTransactor) SendUserOmniMessage(opts *bind.TransactOpts, params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _ISyncRouter.contract.Transact(opts, "sendUserOmniMessage", params)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_ISyncRouter *ISyncRouterSession) SendUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _ISyncRouter.Contract.SendUserOmniMessage(&_ISyncRouter.TransactOpts, params)
}

// SendUserOmniMessage is a paid mutator transaction binding the contract method 0x1b3348f0.
//
// Solidity: function sendUserOmniMessage(((uint8,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,bytes))[],(uint8,uint24,uint64,uint64,uint64,uint64,address,address,uint256,bytes,bytes,bytes)) params) payable returns()
func (_ISyncRouter *ISyncRouterTransactorSession) SendUserOmniMessage(params BaseStructCrossMessageParams) (*types.Transaction, error) {
	return _ISyncRouter.Contract.SendUserOmniMessage(&_ISyncRouter.TransactOpts, params)
}

// IVerifierMetaData contains all meta data concerning the IVerifier contract.
var IVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[1]\",\"name\":\"_pubSignals\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IVerifierMetaData.ABI instead.
var IVerifierABI = IVerifierMetaData.ABI

// IVerifier is an auto generated Go binding around an Ethereum contract.
type IVerifier struct {
	IVerifierCaller     // Read-only binding to the contract
	IVerifierTransactor // Write-only binding to the contract
	IVerifierFilterer   // Log filterer for contract events
}

// IVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVerifierSession struct {
	Contract     *IVerifier        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVerifierCallerSession struct {
	Contract *IVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVerifierTransactorSession struct {
	Contract     *IVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVerifierRaw struct {
	Contract *IVerifier // Generic contract binding to access the raw methods on
}

// IVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVerifierCallerRaw struct {
	Contract *IVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVerifierTransactorRaw struct {
	Contract *IVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVerifier creates a new instance of IVerifier, bound to a specific deployed contract.
func NewIVerifier(address common.Address, backend bind.ContractBackend) (*IVerifier, error) {
	contract, err := bindIVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVerifier{IVerifierCaller: IVerifierCaller{contract: contract}, IVerifierTransactor: IVerifierTransactor{contract: contract}, IVerifierFilterer: IVerifierFilterer{contract: contract}}, nil
}

// NewIVerifierCaller creates a new read-only instance of IVerifier, bound to a specific deployed contract.
func NewIVerifierCaller(address common.Address, caller bind.ContractCaller) (*IVerifierCaller, error) {
	contract, err := bindIVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVerifierCaller{contract: contract}, nil
}

// NewIVerifierTransactor creates a new write-only instance of IVerifier, bound to a specific deployed contract.
func NewIVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IVerifierTransactor, error) {
	contract, err := bindIVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVerifierTransactor{contract: contract}, nil
}

// NewIVerifierFilterer creates a new log filterer instance of IVerifier, bound to a specific deployed contract.
func NewIVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IVerifierFilterer, error) {
	contract, err := bindIVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVerifierFilterer{contract: contract}, nil
}

// bindIVerifier binds a generic wrapper to an already deployed contract.
func bindIVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVerifier *IVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVerifier.Contract.IVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVerifier *IVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVerifier.Contract.IVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVerifier *IVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVerifier.Contract.IVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVerifier *IVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVerifier *IVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVerifier *IVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[1] _pubSignals) view returns(bool)
func (_IVerifier *IVerifierCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _IVerifier.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[1] _pubSignals) view returns(bool)
func (_IVerifier *IVerifierSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [1]*big.Int) (bool, error) {
	return _IVerifier.Contract.VerifyProof(&_IVerifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[1] _pubSignals) view returns(bool)
func (_IVerifier *IVerifierCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [1]*big.Int) (bool, error) {
	return _IVerifier.Contract.VerifyProof(&_IVerifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// PoseidonMetaData contains all meta data concerning the Poseidon contract.
var PoseidonMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"}],\"name\":\"hashMessage\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"parts\",\"type\":\"uint256[4]\"}],\"name\":\"mergeUint64ToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002357610011610028565b61146a61003382393081505061146a90f35b61002e565b60405190565b5f80fdfe60806040526004361015610013575b610351565b61001d5f3561003c565b80635219823a146100375763eb407d9b0361000e57610326565b610205565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061008090610058565b810190811067ffffffffffffffff82111761009a57604052565b610062565b906100b26100ab610042565b9283610076565b565b67ffffffffffffffff81116100d2576100ce602091610058565b0190565b610062565b90825f939282370152565b909291926100f76100f2826100b4565b61009f565b9381855260208501908284011161011357610111926100d7565b565b610054565b9080601f8301121561013657816020610133933591016100e2565b90565b610050565b9060208282031261016b575f82013567ffffffffffffffff8111610166576101639201610118565b90565b61004c565b610048565b50600490565b905090565b90565b90565b61018a9061017e565b9052565b9061019b81602093610181565b0190565b60200190565b6101c16101bb6101b483610170565b8094610176565b9161017b565b5f915b8383106101d15750505050565b6101e76101e1600192845161018e565b9261019f565b920191906101c4565b9190610203905f608085019401906101a5565b565b61022c61021b61021636600461013b565b610a86565b610223610042565b918291826101f0565b0390f35b67ffffffffffffffff81116102455760200290565b610062565b5f80fd5b6102578161017e565b0361025e57565b5f80fd5b9050359061026f8261024e565b565b9092919261028661028182610230565b61009f565b9360208592028301928184116102be57915b8383106102a55750505050565b602080916102b38486610262565b815201920191610298565b61024a565b9080601f830112156102de576102db91600490610271565b90565b610050565b906080828203126102fc576102f9915f016102c3565b90565b610048565b90565b61030d90610301565b9052565b9190610324905f60208501940190610304565b565b61034d61033c6103373660046102e3565b61106e565b610344610042565b91829182610311565b0390f35b5f80fd5b61036161036691610230565b61009f565b90565b369037565b9061038c61037b83610355565b926103868491610230565b90610369565b565b610398600461036e565b90565b90565b60ff60f81b1690565b60f81b90565b6103c16103bc6103c69261039b565b6103a7565b61039e565b90565b5190565b905090565b5f5b8381106103e4575050905f910152565b8060209183015181850152016103d4565b61041a61041192602092610408816103c9565b948580936103cd565b938491016103d2565b0190565b90565b61042d6104329161039e565b61041e565b9052565b61044661044d91600194936103f5565b8092610421565b0190565b90565b90565b61046b61046661047092610451565b610454565b61017e565b90565b634e487b7160e01b5f52601260045260245ffd5b6104936104999161017e565b9161017e565b9081156104a4570690565b610473565b90565b6104c06104bb6104c5926104a9565b610454565b61017e565b90565b6104dc6104d76104e1926104a9565b6103a7565b61039e565b90565b90565b6104fb6104f6610500926104e4565b6103a7565b61039e565b90565b61051761051261051c9261039b565b610454565b61017e565b90565b634e487b7160e01b5f52601160045260245ffd5b6105426105489193929361017e565b9261017e565b820391821161055357565b61051f565b634e487b7160e01b5f52603260045260245ffd5b90610576826103c9565b81101561058857600160209102010190565b610558565b610597905161039e565b90565b90565b6105b16105ac6105b69261059a565b610454565b61017e565b90565b6105c3600861059d565b90565b90565b6105dd6105d86105e2926105c6565b610454565b61017e565b90565b6105ef60076105c9565b90565b6106016106079193929361017e565b9261017e565b9161061383820261017e565b92818404149015171561062257565b61051f565b6106406106326105b9565b61063a6105e5565b906105f2565b90565b61064f6106559161017e565b9161017e565b908115610660570490565b610473565b9061066f9061017e565b9052565b67ffffffffffffffff81116106885760200290565b610062565b61069961069e91610673565b61009f565b90565b906106bf6106ae8361068d565b926106b98491610673565b90610369565b565b6106cb60086106a1565b90565b906106d9910161017e565b90565b5f90565b60ff1690565b6106fa6106f56106ff9261059a565b610454565b6106e0565b90565b1c90565b66ffffffffffffff1690565b6107319061072b610725610736946106e0565b91610706565b90610702565b610706565b90565b60f81c90565b61075361074e610758926106e0565b610454565b6106e0565b90565b61076761076c91610739565b61073f565b90565b61078361077e610788926106e0565b610454565b610706565b90565b90565b6107a261079d6107a79261078b565b610454565b6106e0565b90565b1b90565b6107cd906107c76107c16107d2946106e0565b91610706565b906107aa565b610706565b90565b60016107e1910161017e565b90565b6107f86107f36107fd92610706565b610454565b61017e565b90565b50600890565b9061081082610800565b81101561081e576020020190565b610558565b61083761083261083c926104a9565b610454565b610706565b90565b90565b61084c905161017e565b90565b90565b61086661086161086b9261084f565b610454565b61017e565b90565b90565b61088561088061088a9261086e565b610454565b61017e565b90565b90565b6108a461089f6108a99261088d565b610454565b61017e565b90565b90565b6108c36108be6108c8926108ac565b610454565b61017e565b90565b90565b6108e26108dd6108e7926108cb565b610454565b61017e565b90565b60e01b90565b67ffffffffffffffff81116109085760208091020190565b610062565b9050519061091a8261024e565b565b9092919261093161092c826108f0565b61009f565b938185526020808601920283019281841161096e57915b8383106109555750505050565b60208091610963848661090d565b815201920191610948565b61024a565b9080601f830112156109915781602061098e9351910161091c565b90565b610050565b906020828203126109c6575f82015167ffffffffffffffff81116109c1576109be9201610973565b90565b61004c565b610048565b5190565b60209181520190565b60200190565b60200190565b90610a016109fb6109f4846109cb565b80936109cf565b926109d8565b905f5b818110610a115750505090565b909192610a2a610a24600192865161018e565b946109de565b9101919091610a04565b610a3d90610890565b9052565b60a090610a6d610a62610a749597969460c08401908482035f8601526109e4565b9660208301906101a5565b0190610a34565b565b610a7e610042565b3d5f823e3d90fd5b610a8e61038e565b50610ac3610a9c60016103ad565b91610ab4610aa8610042565b93849260208401610436565b60208201810382520382610076565b5b610ae0610ad0826103c9565b610ada6038610457565b90610487565b610af2610aec5f6104ac565b9161017e565b14610b2f57610b2a610b035f6104c8565b91610b1b610b0f610042565b93849260208401610436565b60208201810382520382610076565b610ac4565b90610b3a60806104e7565b610b778391610b5b610b4b866103c9565b610b556001610503565b90610533565b90610b6e610b69838661056c565b61058d565b175f1a9261056c565b53610b92610b84836103c9565b610b8c610627565b90610643565b90610b9d6004610355565b610bb1610ba95f6104ac565b5f8301610665565b610bc6610bbd5f6104ac565b60208301610665565b610bdb610bd25f6104ac565b60408301610665565b610bf0610be75f6104ac565b60608301610665565b91610bfa5f6104ac565b91610c045f6104ac565b925b83610c19610c138561017e565b9161017e565b1015610fa357610c276106c1565b93610c4d610c3e8884610c38610627565b91611226565b92610c47610627565b906106ce565b96610c566106dc565b97610c605f6104ac565b91610c6a5f6104ac565b96610c745f6104ac565b9a5b8b610c90610c8a610c85610627565b61017e565b9161017e565b1015610da757610d10908c8881610cb7610cb1610cac846103c9565b61017e565b9161017e565b105f14610d8b57610cf4610cee610cdc610cd7610d0995610cf99561056c565b61058d565b5b94610ce860086106e6565b90610712565b9361075b565b61076f565b610d03603061078e565b906107ae565b17946107d5565b9a8b610d2b610d25610d206105e5565b61017e565b9161017e565b14610d41575b610d3a906107d5565b9a93610c76565b9a509296610d67610d55610d6c93996107e4565b610d628b91849092610806565b610665565b6107d5565b95610d765f610823565b92610d3a610d835f6104ac565b9b9050610d31565b5050610d09610cf9610cf4610cee610da25f6104c8565b610cdd565b509650969350975050945f90610dd073__$40e4749874fae718bed7a15ca600323270$__61083f565b610f44610f296310e2e35895610f24610f1b610f16610def600861068d565b93610e16610e0e610e0983610e038e6104ac565b90610806565b610842565b8b8701610665565b610e3e610e35610e3083610e2a6001610503565b90610806565b610842565b60208701610665565b610e66610e5d610e5883610e526002610852565b90610806565b610842565b60408701610665565b610e8e610e85610e8083610e7a6003610871565b90610806565b610842565b60608701610665565b610eb6610ead610ea883610ea26004610890565b90610806565b610842565b60808701610665565b610ede610ed5610ed083610eca60056108af565b90610806565b610842565b60a08701610665565b610f06610efd610ef883610ef260066108ce565b90610806565b610842565b60c08701610665565b610f1060076105c9565b90610806565b610842565b60e08301610665565b61130b565b9294610f4f6004610f38610042565b978896879586956108ea565b855260048501610a41565b03915af4908115610f9e57610f7591610f6f915f91610f7c575b50611386565b946107d5565b9294610c06565b610f9891503d805f833e610f908183610076565b810190610996565b5f610f69565b610a76565b509350505090565b5f90565b90610fb982610170565b811015610fc7576020020190565b610558565b90565b610fe3610fde610fe892610fcc565b610454565b6106e0565b90565b61100a90611004610ffe61100f946106e0565b9161017e565b906107aa565b61017e565b90565b61102661102161102b926104e4565b610454565b6106e0565b90565b90565b61104561104061104a9261102e565b610454565b6106e0565b90565b5f1b90565b61106661106161106b9261017e565b61104d565b610301565b90565b61112b9061107a610fab565b506111256111206110b06110a061109b856110956003610871565b90610faf565b610842565b6110aa60c0610fcf565b90610feb565b6110df6110cf6110ca866110c46002610852565b90610faf565b610842565b6110d96080611012565b90610feb565b1761110f6110ff6110fa866110f46001610503565b90610faf565b610842565b6111096040611031565b90610feb565b179261111a5f6104ac565b90610faf565b610842565b17611052565b90565b606090565b6111426111489193929361017e565b9261017e565b820180921161115357565b61051f565b60209181520190565b5f7f736c696365206f7574206f6620626f756e647300000000000000000000000000910152565b6111956013602092611158565b61119e81611161565b0190565b6111b79060208101905f818303910152611188565b90565b156111c157565b6111c9610042565b62461bcd60e51b8152806111df600482016111a2565b0390fd5b906111f56111f0836100b4565b61009f565b918252565b369037565b9061122461120c836111e3565b9260208061121a86936100b4565b92019103906111fa565b565b919061123061112e565b5061126061123f828490611133565b61125961125361124e876103c9565b61017e565b9161017e565b11156111ba565b602061126b836111ff565b930190602084019101905f915b838310611286575050505090565b60208080928187116112a3575b8051855201920192019190611278565b611293565b606090565b906112bf6112ba836108f0565b61009f565b918252565b906112e96112d1836112ad565b926020806112df86936108f0565b9201910390610369565b565b906112f5826109cb565b811015611306576020809102010190565b610558565b6113136112a8565b50611326611321600861059d565b6112c4565b906113305f6104ac565b5b8061134561133f600861059d565b9161017e565b10156113815761137c90611377611365611360858490610806565b610842565b61137286918490926112eb565b610665565b6107d5565b611331565b505090565b9061138f61038e565b916113995f6104ac565b5b806113ae6113a86004610890565b9161017e565b10156114305780611408916113d36113cd6113c8866109cb565b61017e565b9161017e565b105f1461140d576114026113f06113eb8584906112eb565b610842565b6113fd8791849092610faf565b610665565b5b6107d5565b61139a565b61142b6114195f6104ac565b6114268791849092610faf565b610665565b611403565b505056fea2646970667358221220793fabbc7243f1ce3dc70695ce62019d15a6c4c7ad69b7648478a5cb075875b864736f6c63430008180033",
}

// PoseidonABI is the input ABI used to generate the binding from.
// Deprecated: Use PoseidonMetaData.ABI instead.
var PoseidonABI = PoseidonMetaData.ABI

// PoseidonBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoseidonMetaData.Bin instead.
var PoseidonBin = PoseidonMetaData.Bin

// DeployPoseidon deploys a new Ethereum contract, binding an instance of Poseidon to it.
func DeployPoseidon(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Poseidon, error) {
	parsed, err := PoseidonMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	goldilocksPoseidonAddr, _, _, _ := DeployGoldilocksPoseidon(auth, backend)
	PoseidonBin = strings.ReplaceAll(PoseidonBin, "__$40e4749874fae718bed7a15ca600323270$__", goldilocksPoseidonAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoseidonBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Poseidon{PoseidonCaller: PoseidonCaller{contract: contract}, PoseidonTransactor: PoseidonTransactor{contract: contract}, PoseidonFilterer: PoseidonFilterer{contract: contract}}, nil
}

// Poseidon is an auto generated Go binding around an Ethereum contract.
type Poseidon struct {
	PoseidonCaller     // Read-only binding to the contract
	PoseidonTransactor // Write-only binding to the contract
	PoseidonFilterer   // Log filterer for contract events
}

// PoseidonCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoseidonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoseidonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoseidonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoseidonSession struct {
	Contract     *Poseidon         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoseidonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoseidonCallerSession struct {
	Contract *PoseidonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PoseidonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoseidonTransactorSession struct {
	Contract     *PoseidonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PoseidonRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoseidonRaw struct {
	Contract *Poseidon // Generic contract binding to access the raw methods on
}

// PoseidonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoseidonCallerRaw struct {
	Contract *PoseidonCaller // Generic read-only contract binding to access the raw methods on
}

// PoseidonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoseidonTransactorRaw struct {
	Contract *PoseidonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoseidon creates a new instance of Poseidon, bound to a specific deployed contract.
func NewPoseidon(address common.Address, backend bind.ContractBackend) (*Poseidon, error) {
	contract, err := bindPoseidon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poseidon{PoseidonCaller: PoseidonCaller{contract: contract}, PoseidonTransactor: PoseidonTransactor{contract: contract}, PoseidonFilterer: PoseidonFilterer{contract: contract}}, nil
}

// NewPoseidonCaller creates a new read-only instance of Poseidon, bound to a specific deployed contract.
func NewPoseidonCaller(address common.Address, caller bind.ContractCaller) (*PoseidonCaller, error) {
	contract, err := bindPoseidon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoseidonCaller{contract: contract}, nil
}

// NewPoseidonTransactor creates a new write-only instance of Poseidon, bound to a specific deployed contract.
func NewPoseidonTransactor(address common.Address, transactor bind.ContractTransactor) (*PoseidonTransactor, error) {
	contract, err := bindPoseidon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoseidonTransactor{contract: contract}, nil
}

// NewPoseidonFilterer creates a new log filterer instance of Poseidon, bound to a specific deployed contract.
func NewPoseidonFilterer(address common.Address, filterer bind.ContractFilterer) (*PoseidonFilterer, error) {
	contract, err := bindPoseidon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoseidonFilterer{contract: contract}, nil
}

// bindPoseidon binds a generic wrapper to an already deployed contract.
func bindPoseidon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoseidonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poseidon *PoseidonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poseidon.Contract.PoseidonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poseidon *PoseidonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poseidon.Contract.PoseidonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poseidon *PoseidonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poseidon.Contract.PoseidonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poseidon *PoseidonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poseidon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poseidon *PoseidonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poseidon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poseidon *PoseidonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poseidon.Contract.contract.Transact(opts, method, params...)
}

// HashMessage is a free data retrieval call binding the contract method 0x5219823a.
//
// Solidity: function hashMessage(bytes bytecode) pure returns(uint256[4])
func (_Poseidon *PoseidonCaller) HashMessage(opts *bind.CallOpts, bytecode []byte) ([4]*big.Int, error) {
	var out []interface{}
	err := _Poseidon.contract.Call(opts, &out, "hashMessage", bytecode)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// HashMessage is a free data retrieval call binding the contract method 0x5219823a.
//
// Solidity: function hashMessage(bytes bytecode) pure returns(uint256[4])
func (_Poseidon *PoseidonSession) HashMessage(bytecode []byte) ([4]*big.Int, error) {
	return _Poseidon.Contract.HashMessage(&_Poseidon.CallOpts, bytecode)
}

// HashMessage is a free data retrieval call binding the contract method 0x5219823a.
//
// Solidity: function hashMessage(bytes bytecode) pure returns(uint256[4])
func (_Poseidon *PoseidonCallerSession) HashMessage(bytecode []byte) ([4]*big.Int, error) {
	return _Poseidon.Contract.HashMessage(&_Poseidon.CallOpts, bytecode)
}

// MergeUint64ToBytes32 is a free data retrieval call binding the contract method 0xeb407d9b.
//
// Solidity: function mergeUint64ToBytes32(uint256[4] parts) pure returns(bytes32)
func (_Poseidon *PoseidonCaller) MergeUint64ToBytes32(opts *bind.CallOpts, parts [4]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Poseidon.contract.Call(opts, &out, "mergeUint64ToBytes32", parts)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MergeUint64ToBytes32 is a free data retrieval call binding the contract method 0xeb407d9b.
//
// Solidity: function mergeUint64ToBytes32(uint256[4] parts) pure returns(bytes32)
func (_Poseidon *PoseidonSession) MergeUint64ToBytes32(parts [4]*big.Int) ([32]byte, error) {
	return _Poseidon.Contract.MergeUint64ToBytes32(&_Poseidon.CallOpts, parts)
}

// MergeUint64ToBytes32 is a free data retrieval call binding the contract method 0xeb407d9b.
//
// Solidity: function mergeUint64ToBytes32(uint256[4] parts) pure returns(bytes32)
func (_Poseidon *PoseidonCallerSession) MergeUint64ToBytes32(parts [4]*big.Int) ([32]byte, error) {
	return _Poseidon.Contract.MergeUint64ToBytes32(&_Poseidon.CallOpts, parts)
}

// PreGasManagerMetaData contains all meta data concerning the PreGasManager contract.
var PreGasManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotEqual\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"did\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositTicketDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTicketDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPreGasBalanceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitDepositOperation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submitWithdrawOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523461001f57610011610024565b6105b361002f82396105b390f35b61002a565b60405190565b5f80fdfe60806040526004361015610013575b6101e2565b61001d5f3561004c565b806339875d1a14610047578063653e75ec1461004257639abe79030361000e576101b8565b610156565b6100a9565b60e01c90565b60405190565b5f80fd5b5f80fd5b90565b61006c81610060565b0361007357565b5f80fd5b9050359061008482610063565b565b9060208282031261009f5761009c915f01610077565b90565b61005c565b5f0190565b346100d7576100c16100bc366004610086565b610240565b6100c9610052565b806100d3816100a4565b0390f35b610058565b60018060a01b031690565b6100f0906100dc565b90565b6100fc816100e7565b0361010357565b5f80fd5b90503590610114826100f3565b565b9060208282031261012f5761012c915f01610107565b90565b61005c565b61013d90610060565b9052565b9190610154905f60208501940190610134565b565b346101865761018261017161016c366004610116565b6102c9565b610179610052565b91829182610141565b0390f35b610058565b91906040838203126101b357806101a76101b0925f8601610077565b93602001610077565b90565b61005c565b6101cc6101c636600461018b565b90610385565b6101d4610052565b806101de816100a4565b0390f35b5f80fd5b90565b6101fd6101f8610202926100dc565b6101e6565b6100dc565b90565b61020e906101e9565b90565b61021a90610205565b90565b91602061023e92949361023760408201965f830190610134565b0190610134565b565b334261026c7f2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a92610211565b92610281610278610052565b9283928361021d565b0390a2565b5f90565b9061029490610211565b5f5260205260405f2090565b5f1c90565b90565b6102b46102b9916102a0565b6102a5565b90565b6102c690546102a8565b90565b6102df6102e4916102d8610286565b505f61028a565b6102bc565b90565b156102ee57565b5f80fd5b634e487b7160e01b5f52601160045260245ffd5b61031561031b91939293610060565b92610060565b820180921161032657565b6102f2565b5f1b90565b9061033c5f199161032b565b9181191691161790565b61035a61035561035f92610060565b6101e6565b610060565b90565b90565b9061037a61037561038192610346565b610362565b8254610330565b9055565b6103d5916103a53461039f61039985610060565b91610060565b146102e7565b6103cc826103c66103b75f339061028a565b916103c1836102bc565b610306565b90610365565b339190916104bb565b565b60601b90565b6103e6906103d7565b90565b6103f2906103dd565b90565b610401610406916100e7565b6103e9565b9052565b90565b61041961041e91610060565b61040a565b9052565b9261044f60206104579461044760148861043f859b9a86996103f5565b01809261040d565b01809261040d565b01809261040d565b0190565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906104839061045b565b810190811067ffffffffffffffff82111761049d57604052565b610465565b60200190565b5190565b90565b6104b8906104ac565b90565b916104eb6104d26104cd5f869061028a565b6102bc565b6104e46104de85610060565b91610060565b10156102e7565b6105198361050a469385906104fe610052565b95869460208601610422565b60208201810382520382610479565b61052b610525826104a8565b916104a2565b209190914261056361055d7f081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc936104af565b93610211565b9361057861056f610052565b9283928361021d565b0390a356fea26469706673582212207c4494f39853cde456406a7255148ad3b8626ed7355d090ab5ecbd661b7892fb64736f6c63430008180033",
}

// PreGasManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PreGasManagerMetaData.ABI instead.
var PreGasManagerABI = PreGasManagerMetaData.ABI

// PreGasManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PreGasManagerMetaData.Bin instead.
var PreGasManagerBin = PreGasManagerMetaData.Bin

// DeployPreGasManager deploys a new Ethereum contract, binding an instance of PreGasManager to it.
func DeployPreGasManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PreGasManager, error) {
	parsed, err := PreGasManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PreGasManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PreGasManager{PreGasManagerCaller: PreGasManagerCaller{contract: contract}, PreGasManagerTransactor: PreGasManagerTransactor{contract: contract}, PreGasManagerFilterer: PreGasManagerFilterer{contract: contract}}, nil
}

// PreGasManager is an auto generated Go binding around an Ethereum contract.
type PreGasManager struct {
	PreGasManagerCaller     // Read-only binding to the contract
	PreGasManagerTransactor // Write-only binding to the contract
	PreGasManagerFilterer   // Log filterer for contract events
}

// PreGasManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreGasManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreGasManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreGasManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreGasManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreGasManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreGasManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreGasManagerSession struct {
	Contract     *PreGasManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PreGasManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreGasManagerCallerSession struct {
	Contract *PreGasManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PreGasManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreGasManagerTransactorSession struct {
	Contract     *PreGasManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PreGasManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreGasManagerRaw struct {
	Contract *PreGasManager // Generic contract binding to access the raw methods on
}

// PreGasManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreGasManagerCallerRaw struct {
	Contract *PreGasManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PreGasManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreGasManagerTransactorRaw struct {
	Contract *PreGasManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreGasManager creates a new instance of PreGasManager, bound to a specific deployed contract.
func NewPreGasManager(address common.Address, backend bind.ContractBackend) (*PreGasManager, error) {
	contract, err := bindPreGasManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PreGasManager{PreGasManagerCaller: PreGasManagerCaller{contract: contract}, PreGasManagerTransactor: PreGasManagerTransactor{contract: contract}, PreGasManagerFilterer: PreGasManagerFilterer{contract: contract}}, nil
}

// NewPreGasManagerCaller creates a new read-only instance of PreGasManager, bound to a specific deployed contract.
func NewPreGasManagerCaller(address common.Address, caller bind.ContractCaller) (*PreGasManagerCaller, error) {
	contract, err := bindPreGasManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreGasManagerCaller{contract: contract}, nil
}

// NewPreGasManagerTransactor creates a new write-only instance of PreGasManager, bound to a specific deployed contract.
func NewPreGasManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PreGasManagerTransactor, error) {
	contract, err := bindPreGasManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreGasManagerTransactor{contract: contract}, nil
}

// NewPreGasManagerFilterer creates a new log filterer instance of PreGasManager, bound to a specific deployed contract.
func NewPreGasManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PreGasManagerFilterer, error) {
	contract, err := bindPreGasManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreGasManagerFilterer{contract: contract}, nil
}

// bindPreGasManager binds a generic wrapper to an already deployed contract.
func bindPreGasManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PreGasManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreGasManager *PreGasManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreGasManager.Contract.PreGasManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreGasManager *PreGasManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreGasManager.Contract.PreGasManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreGasManager *PreGasManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreGasManager.Contract.PreGasManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreGasManager *PreGasManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreGasManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreGasManager *PreGasManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreGasManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreGasManager *PreGasManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreGasManager.Contract.contract.Transact(opts, method, params...)
}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_PreGasManager *PreGasManagerCaller) GetPreGasBalanceInfo(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PreGasManager.contract.Call(opts, &out, "getPreGasBalanceInfo", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_PreGasManager *PreGasManagerSession) GetPreGasBalanceInfo(account common.Address) (*big.Int, error) {
	return _PreGasManager.Contract.GetPreGasBalanceInfo(&_PreGasManager.CallOpts, account)
}

// GetPreGasBalanceInfo is a free data retrieval call binding the contract method 0x653e75ec.
//
// Solidity: function getPreGasBalanceInfo(address account) view returns(uint256)
func (_PreGasManager *PreGasManagerCallerSession) GetPreGasBalanceInfo(account common.Address) (*big.Int, error) {
	return _PreGasManager.Contract.GetPreGasBalanceInfo(&_PreGasManager.CallOpts, account)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_PreGasManager *PreGasManagerTransactor) SubmitDepositOperation(opts *bind.TransactOpts, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _PreGasManager.contract.Transact(opts, "submitDepositOperation", amount, nonce)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_PreGasManager *PreGasManagerSession) SubmitDepositOperation(amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _PreGasManager.Contract.SubmitDepositOperation(&_PreGasManager.TransactOpts, amount, nonce)
}

// SubmitDepositOperation is a paid mutator transaction binding the contract method 0x9abe7903.
//
// Solidity: function submitDepositOperation(uint256 amount, uint256 nonce) payable returns()
func (_PreGasManager *PreGasManagerTransactorSession) SubmitDepositOperation(amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _PreGasManager.Contract.SubmitDepositOperation(&_PreGasManager.TransactOpts, amount, nonce)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_PreGasManager *PreGasManagerTransactor) SubmitWithdrawOperation(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PreGasManager.contract.Transact(opts, "submitWithdrawOperation", amount)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_PreGasManager *PreGasManagerSession) SubmitWithdrawOperation(amount *big.Int) (*types.Transaction, error) {
	return _PreGasManager.Contract.SubmitWithdrawOperation(&_PreGasManager.TransactOpts, amount)
}

// SubmitWithdrawOperation is a paid mutator transaction binding the contract method 0x39875d1a.
//
// Solidity: function submitWithdrawOperation(uint256 amount) returns()
func (_PreGasManager *PreGasManagerTransactorSession) SubmitWithdrawOperation(amount *big.Int) (*types.Transaction, error) {
	return _PreGasManager.Contract.SubmitWithdrawOperation(&_PreGasManager.TransactOpts, amount)
}

// PreGasManagerDepositTicketAddedIterator is returned from FilterDepositTicketAdded and is used to iterate over the raw logs and unpacked data for DepositTicketAdded events raised by the PreGasManager contract.
type PreGasManagerDepositTicketAddedIterator struct {
	Event *PreGasManagerDepositTicketAdded // Event containing the contract specifics and raw log

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
func (it *PreGasManagerDepositTicketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreGasManagerDepositTicketAdded)
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
		it.Event = new(PreGasManagerDepositTicketAdded)
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
func (it *PreGasManagerDepositTicketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreGasManagerDepositTicketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreGasManagerDepositTicketAdded represents a DepositTicketAdded event raised by the PreGasManager contract.
type PreGasManagerDepositTicketAdded struct {
	Did       [32]byte
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketAdded is a free log retrieval operation binding the contract event 0x081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc.
//
// Solidity: event DepositTicketAdded(bytes32 indexed did, address indexed account, uint256 amount, uint256 timestamp)
func (_PreGasManager *PreGasManagerFilterer) FilterDepositTicketAdded(opts *bind.FilterOpts, did [][32]byte, account []common.Address) (*PreGasManagerDepositTicketAddedIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PreGasManager.contract.FilterLogs(opts, "DepositTicketAdded", didRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &PreGasManagerDepositTicketAddedIterator{contract: _PreGasManager.contract, event: "DepositTicketAdded", logs: logs, sub: sub}, nil
}

// WatchDepositTicketAdded is a free log subscription operation binding the contract event 0x081cdaa84d746e50d6d31f53d3392249b7b54f556a6b5e52a4c3d0396ca622dc.
//
// Solidity: event DepositTicketAdded(bytes32 indexed did, address indexed account, uint256 amount, uint256 timestamp)
func (_PreGasManager *PreGasManagerFilterer) WatchDepositTicketAdded(opts *bind.WatchOpts, sink chan<- *PreGasManagerDepositTicketAdded, did [][32]byte, account []common.Address) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PreGasManager.contract.WatchLogs(opts, "DepositTicketAdded", didRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreGasManagerDepositTicketAdded)
				if err := _PreGasManager.contract.UnpackLog(event, "DepositTicketAdded", log); err != nil {
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
func (_PreGasManager *PreGasManagerFilterer) ParseDepositTicketAdded(log types.Log) (*PreGasManagerDepositTicketAdded, error) {
	event := new(PreGasManagerDepositTicketAdded)
	if err := _PreGasManager.contract.UnpackLog(event, "DepositTicketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreGasManagerDepositTicketDeletedIterator is returned from FilterDepositTicketDeleted and is used to iterate over the raw logs and unpacked data for DepositTicketDeleted events raised by the PreGasManager contract.
type PreGasManagerDepositTicketDeletedIterator struct {
	Event *PreGasManagerDepositTicketDeleted // Event containing the contract specifics and raw log

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
func (it *PreGasManagerDepositTicketDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreGasManagerDepositTicketDeleted)
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
		it.Event = new(PreGasManagerDepositTicketDeleted)
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
func (it *PreGasManagerDepositTicketDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreGasManagerDepositTicketDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreGasManagerDepositTicketDeleted represents a DepositTicketDeleted event raised by the PreGasManager contract.
type PreGasManagerDepositTicketDeleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositTicketDeleted is a free log retrieval operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_PreGasManager *PreGasManagerFilterer) FilterDepositTicketDeleted(opts *bind.FilterOpts, account []common.Address) (*PreGasManagerDepositTicketDeletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PreGasManager.contract.FilterLogs(opts, "DepositTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &PreGasManagerDepositTicketDeletedIterator{contract: _PreGasManager.contract, event: "DepositTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchDepositTicketDeleted is a free log subscription operation binding the contract event 0x9694d3a9f64b53ccb1de2bdfe41a6c45ddd7ac92f4e89f100edb56fcbf39af1d.
//
// Solidity: event DepositTicketDeleted(address indexed account, uint256 amount)
func (_PreGasManager *PreGasManagerFilterer) WatchDepositTicketDeleted(opts *bind.WatchOpts, sink chan<- *PreGasManagerDepositTicketDeleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PreGasManager.contract.WatchLogs(opts, "DepositTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreGasManagerDepositTicketDeleted)
				if err := _PreGasManager.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
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
func (_PreGasManager *PreGasManagerFilterer) ParseDepositTicketDeleted(log types.Log) (*PreGasManagerDepositTicketDeleted, error) {
	event := new(PreGasManagerDepositTicketDeleted)
	if err := _PreGasManager.contract.UnpackLog(event, "DepositTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreGasManagerWithdrawTicketAddedIterator is returned from FilterWithdrawTicketAdded and is used to iterate over the raw logs and unpacked data for WithdrawTicketAdded events raised by the PreGasManager contract.
type PreGasManagerWithdrawTicketAddedIterator struct {
	Event *PreGasManagerWithdrawTicketAdded // Event containing the contract specifics and raw log

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
func (it *PreGasManagerWithdrawTicketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreGasManagerWithdrawTicketAdded)
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
		it.Event = new(PreGasManagerWithdrawTicketAdded)
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
func (it *PreGasManagerWithdrawTicketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreGasManagerWithdrawTicketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreGasManagerWithdrawTicketAdded represents a WithdrawTicketAdded event raised by the PreGasManager contract.
type PreGasManagerWithdrawTicketAdded struct {
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketAdded is a free log retrieval operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_PreGasManager *PreGasManagerFilterer) FilterWithdrawTicketAdded(opts *bind.FilterOpts, account []common.Address) (*PreGasManagerWithdrawTicketAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PreGasManager.contract.FilterLogs(opts, "WithdrawTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PreGasManagerWithdrawTicketAddedIterator{contract: _PreGasManager.contract, event: "WithdrawTicketAdded", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketAdded is a free log subscription operation binding the contract event 0x2e20ec63884d9292df4974b2a60c1037e4f5b6eefb22756aa5be4e4cf78b611a.
//
// Solidity: event WithdrawTicketAdded(address indexed account, uint256 amount, uint256 timestamp)
func (_PreGasManager *PreGasManagerFilterer) WatchWithdrawTicketAdded(opts *bind.WatchOpts, sink chan<- *PreGasManagerWithdrawTicketAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PreGasManager.contract.WatchLogs(opts, "WithdrawTicketAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreGasManagerWithdrawTicketAdded)
				if err := _PreGasManager.contract.UnpackLog(event, "WithdrawTicketAdded", log); err != nil {
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
func (_PreGasManager *PreGasManagerFilterer) ParseWithdrawTicketAdded(log types.Log) (*PreGasManagerWithdrawTicketAdded, error) {
	event := new(PreGasManagerWithdrawTicketAdded)
	if err := _PreGasManager.contract.UnpackLog(event, "WithdrawTicketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreGasManagerWithdrawTicketDeletedIterator is returned from FilterWithdrawTicketDeleted and is used to iterate over the raw logs and unpacked data for WithdrawTicketDeleted events raised by the PreGasManager contract.
type PreGasManagerWithdrawTicketDeletedIterator struct {
	Event *PreGasManagerWithdrawTicketDeleted // Event containing the contract specifics and raw log

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
func (it *PreGasManagerWithdrawTicketDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreGasManagerWithdrawTicketDeleted)
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
		it.Event = new(PreGasManagerWithdrawTicketDeleted)
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
func (it *PreGasManagerWithdrawTicketDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreGasManagerWithdrawTicketDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreGasManagerWithdrawTicketDeleted represents a WithdrawTicketDeleted event raised by the PreGasManager contract.
type PreGasManagerWithdrawTicketDeleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTicketDeleted is a free log retrieval operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_PreGasManager *PreGasManagerFilterer) FilterWithdrawTicketDeleted(opts *bind.FilterOpts, account []common.Address) (*PreGasManagerWithdrawTicketDeletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PreGasManager.contract.FilterLogs(opts, "WithdrawTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &PreGasManagerWithdrawTicketDeletedIterator{contract: _PreGasManager.contract, event: "WithdrawTicketDeleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawTicketDeleted is a free log subscription operation binding the contract event 0xe971a029a1941c0419902b24a3744c31cd3431b3233d016c7980d9bed21b168f.
//
// Solidity: event WithdrawTicketDeleted(address indexed account, uint256 amount)
func (_PreGasManager *PreGasManagerFilterer) WatchWithdrawTicketDeleted(opts *bind.WatchOpts, sink chan<- *PreGasManagerWithdrawTicketDeleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PreGasManager.contract.WatchLogs(opts, "WithdrawTicketDeleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreGasManagerWithdrawTicketDeleted)
				if err := _PreGasManager.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
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
func (_PreGasManager *PreGasManagerFilterer) ParseWithdrawTicketDeleted(log types.Log) (*PreGasManagerWithdrawTicketDeleted, error) {
	event := new(PreGasManagerWithdrawTicketDeleted)
	if err := _PreGasManager.contract.UnpackLog(event, "WithdrawTicketDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardMetaData contains all meta data concerning the ReentrancyGuard contract.
var ReentrancyGuardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"}]",
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardMetaData.ABI instead.
var ReentrancyGuardABI = ReentrancyGuardMetaData.ABI

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReentrancyGuardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// StateManagerMetaData contains all meta data concerning the StateManager contract.
var StateManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchNumToState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accInputRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523461001f57610011610024565b61027061002f823961027090f35b61002a565b60405190565b5f80fdfe60806040526004361015610013575b610236565b61001d5f3561003c565b80636b0862371461003757637fcb36530361000e57610201565b610154565b60e01c90565b60405190565b5f80fd5b5f80fd5b67ffffffffffffffff1690565b61006681610050565b0361006d57565b5f80fd5b9050359061007e8261005d565b565b9060208282031261009957610096915f01610071565b90565b61004c565b90565b6100b56100b06100ba92610050565b61009e565b610050565b90565b906100c7906100a1565b5f5260205260405f2090565b5f1c90565b90565b6100e76100ec916100d3565b6100d8565b90565b6100f990546100db565b90565b610106905f6100bd565b9061011e60016101175f85016100ef565b93016100ef565b90565b90565b61012d90610121565b9052565b91602061015292949361014b60408201965f830190610124565b0190610124565b565b346101855761016c610167366004610080565b6100fc565b90610181610178610042565b92839283610131565b0390f35b610048565b5f91031261019457565b61004c565b1c90565b67ffffffffffffffff1690565b6101ba9060086101bf9302610199565b61019d565b90565b906101cd91546101aa565b90565b6101dc60015f906101c2565b90565b6101e890610050565b9052565b91906101ff905f602085019401906101df565b565b346102315761021136600461018a565b61022d61021c6101d0565b610224610042565b918291826101ec565b0390f35b610048565b5f80fdfea2646970667358221220ecb48a2891744d2f5fa44e29a86d8d220be4ebe0c75cd318dac32d04c8c3ba8164736f6c63430008180033",
}

// StateManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StateManagerMetaData.ABI instead.
var StateManagerABI = StateManagerMetaData.ABI

// StateManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateManagerMetaData.Bin instead.
var StateManagerBin = StateManagerMetaData.Bin

// DeployStateManager deploys a new Ethereum contract, binding an instance of StateManager to it.
func DeployStateManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StateManager, error) {
	parsed, err := StateManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateManager{StateManagerCaller: StateManagerCaller{contract: contract}, StateManagerTransactor: StateManagerTransactor{contract: contract}, StateManagerFilterer: StateManagerFilterer{contract: contract}}, nil
}

// StateManager is an auto generated Go binding around an Ethereum contract.
type StateManager struct {
	StateManagerCaller     // Read-only binding to the contract
	StateManagerTransactor // Write-only binding to the contract
	StateManagerFilterer   // Log filterer for contract events
}

// StateManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateManagerSession struct {
	Contract     *StateManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateManagerCallerSession struct {
	Contract *StateManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StateManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateManagerTransactorSession struct {
	Contract     *StateManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StateManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateManagerRaw struct {
	Contract *StateManager // Generic contract binding to access the raw methods on
}

// StateManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateManagerCallerRaw struct {
	Contract *StateManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StateManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateManagerTransactorRaw struct {
	Contract *StateManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateManager creates a new instance of StateManager, bound to a specific deployed contract.
func NewStateManager(address common.Address, backend bind.ContractBackend) (*StateManager, error) {
	contract, err := bindStateManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateManager{StateManagerCaller: StateManagerCaller{contract: contract}, StateManagerTransactor: StateManagerTransactor{contract: contract}, StateManagerFilterer: StateManagerFilterer{contract: contract}}, nil
}

// NewStateManagerCaller creates a new read-only instance of StateManager, bound to a specific deployed contract.
func NewStateManagerCaller(address common.Address, caller bind.ContractCaller) (*StateManagerCaller, error) {
	contract, err := bindStateManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateManagerCaller{contract: contract}, nil
}

// NewStateManagerTransactor creates a new write-only instance of StateManager, bound to a specific deployed contract.
func NewStateManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StateManagerTransactor, error) {
	contract, err := bindStateManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateManagerTransactor{contract: contract}, nil
}

// NewStateManagerFilterer creates a new log filterer instance of StateManager, bound to a specific deployed contract.
func NewStateManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StateManagerFilterer, error) {
	contract, err := bindStateManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateManagerFilterer{contract: contract}, nil
}

// bindStateManager binds a generic wrapper to an already deployed contract.
func bindStateManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StateManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateManager *StateManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateManager.Contract.StateManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateManager *StateManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateManager.Contract.StateManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateManager *StateManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateManager.Contract.StateManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateManager *StateManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateManager *StateManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateManager *StateManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateManager.Contract.contract.Transact(opts, method, params...)
}

// BatchNumToState is a free data retrieval call binding the contract method 0x6b086237.
//
// Solidity: function batchNumToState(uint64 ) view returns(bytes32 stateRoot, bytes32 accInputRoot)
func (_StateManager *StateManagerCaller) BatchNumToState(opts *bind.CallOpts, arg0 uint64) (struct {
	StateRoot    [32]byte
	AccInputRoot [32]byte
}, error) {
	var out []interface{}
	err := _StateManager.contract.Call(opts, &out, "batchNumToState", arg0)

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
func (_StateManager *StateManagerSession) BatchNumToState(arg0 uint64) (struct {
	StateRoot    [32]byte
	AccInputRoot [32]byte
}, error) {
	return _StateManager.Contract.BatchNumToState(&_StateManager.CallOpts, arg0)
}

// BatchNumToState is a free data retrieval call binding the contract method 0x6b086237.
//
// Solidity: function batchNumToState(uint64 ) view returns(bytes32 stateRoot, bytes32 accInputRoot)
func (_StateManager *StateManagerCallerSession) BatchNumToState(arg0 uint64) (struct {
	StateRoot    [32]byte
	AccInputRoot [32]byte
}, error) {
	return _StateManager.Contract.BatchNumToState(&_StateManager.CallOpts, arg0)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_StateManager *StateManagerCaller) LastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _StateManager.contract.Call(opts, &out, "lastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_StateManager *StateManagerSession) LastVerifiedBatch() (uint64, error) {
	return _StateManager.Contract.LastVerifiedBatch(&_StateManager.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_StateManager *StateManagerCallerSession) LastVerifiedBatch() (uint64, error) {
	return _StateManager.Contract.LastVerifiedBatch(&_StateManager.CallOpts)
}

// UserOperationLibMetaData contains all meta data concerning the UserOperationLib contract.
var UserOperationLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x608060405234601d57600e6021565b603e602c823930815050603e90f35b6027565b60405190565b5f80fdfe60806040525f80fdfea26469706673582212206b1f9c961649695bf54ef8ec872f2be3edff8fdec9f696e7b16ada9400e8b5c964736f6c63430008180033",
}

// UserOperationLibABI is the input ABI used to generate the binding from.
// Deprecated: Use UserOperationLibMetaData.ABI instead.
var UserOperationLibABI = UserOperationLibMetaData.ABI

// UserOperationLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserOperationLibMetaData.Bin instead.
var UserOperationLibBin = UserOperationLibMetaData.Bin

// DeployUserOperationLib deploys a new Ethereum contract, binding an instance of UserOperationLib to it.
func DeployUserOperationLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserOperationLib, error) {
	parsed, err := UserOperationLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserOperationLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserOperationLib{UserOperationLibCaller: UserOperationLibCaller{contract: contract}, UserOperationLibTransactor: UserOperationLibTransactor{contract: contract}, UserOperationLibFilterer: UserOperationLibFilterer{contract: contract}}, nil
}

// UserOperationLib is an auto generated Go binding around an Ethereum contract.
type UserOperationLib struct {
	UserOperationLibCaller     // Read-only binding to the contract
	UserOperationLibTransactor // Write-only binding to the contract
	UserOperationLibFilterer   // Log filterer for contract events
}

// UserOperationLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserOperationLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserOperationLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserOperationLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserOperationLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserOperationLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserOperationLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserOperationLibSession struct {
	Contract     *UserOperationLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserOperationLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserOperationLibCallerSession struct {
	Contract *UserOperationLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UserOperationLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserOperationLibTransactorSession struct {
	Contract     *UserOperationLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UserOperationLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserOperationLibRaw struct {
	Contract *UserOperationLib // Generic contract binding to access the raw methods on
}

// UserOperationLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserOperationLibCallerRaw struct {
	Contract *UserOperationLibCaller // Generic read-only contract binding to access the raw methods on
}

// UserOperationLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserOperationLibTransactorRaw struct {
	Contract *UserOperationLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserOperationLib creates a new instance of UserOperationLib, bound to a specific deployed contract.
func NewUserOperationLib(address common.Address, backend bind.ContractBackend) (*UserOperationLib, error) {
	contract, err := bindUserOperationLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserOperationLib{UserOperationLibCaller: UserOperationLibCaller{contract: contract}, UserOperationLibTransactor: UserOperationLibTransactor{contract: contract}, UserOperationLibFilterer: UserOperationLibFilterer{contract: contract}}, nil
}

// NewUserOperationLibCaller creates a new read-only instance of UserOperationLib, bound to a specific deployed contract.
func NewUserOperationLibCaller(address common.Address, caller bind.ContractCaller) (*UserOperationLibCaller, error) {
	contract, err := bindUserOperationLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserOperationLibCaller{contract: contract}, nil
}

// NewUserOperationLibTransactor creates a new write-only instance of UserOperationLib, bound to a specific deployed contract.
func NewUserOperationLibTransactor(address common.Address, transactor bind.ContractTransactor) (*UserOperationLibTransactor, error) {
	contract, err := bindUserOperationLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserOperationLibTransactor{contract: contract}, nil
}

// NewUserOperationLibFilterer creates a new log filterer instance of UserOperationLib, bound to a specific deployed contract.
func NewUserOperationLibFilterer(address common.Address, filterer bind.ContractFilterer) (*UserOperationLibFilterer, error) {
	contract, err := bindUserOperationLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserOperationLibFilterer{contract: contract}, nil
}

// bindUserOperationLib binds a generic wrapper to an already deployed contract.
func bindUserOperationLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserOperationLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserOperationLib *UserOperationLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserOperationLib.Contract.UserOperationLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserOperationLib *UserOperationLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserOperationLib.Contract.UserOperationLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserOperationLib *UserOperationLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserOperationLib.Contract.UserOperationLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserOperationLib *UserOperationLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserOperationLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserOperationLib *UserOperationLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserOperationLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserOperationLib *UserOperationLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserOperationLib.Contract.contract.Transact(opts, method, params...)
}

// UserOperationsLibMetaData contains all meta data concerning the UserOperationsLib contract.
var UserOperationsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x608060405234601d57600e6021565b603e602c823930815050603e90f35b6027565b60405190565b5f80fdfe60806040525f80fdfea2646970667358221220fb2914d495dbf7ee87050a0f67e2718d70b3b258d02259e170bdb80dfec8119764736f6c63430008180033",
}

// UserOperationsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use UserOperationsLibMetaData.ABI instead.
var UserOperationsLibABI = UserOperationsLibMetaData.ABI

// UserOperationsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserOperationsLibMetaData.Bin instead.
var UserOperationsLibBin = UserOperationsLibMetaData.Bin

// DeployUserOperationsLib deploys a new Ethereum contract, binding an instance of UserOperationsLib to it.
func DeployUserOperationsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserOperationsLib, error) {
	parsed, err := UserOperationsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserOperationsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserOperationsLib{UserOperationsLibCaller: UserOperationsLibCaller{contract: contract}, UserOperationsLibTransactor: UserOperationsLibTransactor{contract: contract}, UserOperationsLibFilterer: UserOperationsLibFilterer{contract: contract}}, nil
}

// UserOperationsLib is an auto generated Go binding around an Ethereum contract.
type UserOperationsLib struct {
	UserOperationsLibCaller     // Read-only binding to the contract
	UserOperationsLibTransactor // Write-only binding to the contract
	UserOperationsLibFilterer   // Log filterer for contract events
}

// UserOperationsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserOperationsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserOperationsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserOperationsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserOperationsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserOperationsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserOperationsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserOperationsLibSession struct {
	Contract     *UserOperationsLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UserOperationsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserOperationsLibCallerSession struct {
	Contract *UserOperationsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UserOperationsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserOperationsLibTransactorSession struct {
	Contract     *UserOperationsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UserOperationsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserOperationsLibRaw struct {
	Contract *UserOperationsLib // Generic contract binding to access the raw methods on
}

// UserOperationsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserOperationsLibCallerRaw struct {
	Contract *UserOperationsLibCaller // Generic read-only contract binding to access the raw methods on
}

// UserOperationsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserOperationsLibTransactorRaw struct {
	Contract *UserOperationsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserOperationsLib creates a new instance of UserOperationsLib, bound to a specific deployed contract.
func NewUserOperationsLib(address common.Address, backend bind.ContractBackend) (*UserOperationsLib, error) {
	contract, err := bindUserOperationsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserOperationsLib{UserOperationsLibCaller: UserOperationsLibCaller{contract: contract}, UserOperationsLibTransactor: UserOperationsLibTransactor{contract: contract}, UserOperationsLibFilterer: UserOperationsLibFilterer{contract: contract}}, nil
}

// NewUserOperationsLibCaller creates a new read-only instance of UserOperationsLib, bound to a specific deployed contract.
func NewUserOperationsLibCaller(address common.Address, caller bind.ContractCaller) (*UserOperationsLibCaller, error) {
	contract, err := bindUserOperationsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserOperationsLibCaller{contract: contract}, nil
}

// NewUserOperationsLibTransactor creates a new write-only instance of UserOperationsLib, bound to a specific deployed contract.
func NewUserOperationsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*UserOperationsLibTransactor, error) {
	contract, err := bindUserOperationsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserOperationsLibTransactor{contract: contract}, nil
}

// NewUserOperationsLibFilterer creates a new log filterer instance of UserOperationsLib, bound to a specific deployed contract.
func NewUserOperationsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*UserOperationsLibFilterer, error) {
	contract, err := bindUserOperationsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserOperationsLibFilterer{contract: contract}, nil
}

// bindUserOperationsLib binds a generic wrapper to an already deployed contract.
func bindUserOperationsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserOperationsLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserOperationsLib *UserOperationsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserOperationsLib.Contract.UserOperationsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserOperationsLib *UserOperationsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserOperationsLib.Contract.UserOperationsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserOperationsLib *UserOperationsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserOperationsLib.Contract.UserOperationsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserOperationsLib *UserOperationsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserOperationsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserOperationsLib *UserOperationsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserOperationsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserOperationsLib *UserOperationsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserOperationsLib.Contract.contract.Transact(opts, method, params...)
}

// VerifierHelperMetaData contains all meta data concerning the VerifierHelper contract.
var VerifierHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"getSnarkProof\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002357610011610028565b61055561003382393081505061055590f35b61002e565b60405190565b5f80fdfe60806040526004361015610013575b610263565b61001d5f3561002c565b63f3b9a8b60361000e57610234565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b909182601f830112156100865781359167ffffffffffffffff831161008157602001926001830284011161007c57565b610048565b610044565b610040565b906020828203126100bc575f82013567ffffffffffffffff81116100b7576100b3920161004c565b9091565b61003c565b610038565b50600290565b905090565b90565b90565b6100db906100cf565b9052565b906100ec816020936100d2565b0190565b60200190565b61011261010c610105836100c1565b80946100c7565b916100cc565b5f915b8383106101225750505050565b61013861013260019284516100df565b926100f0565b92019190610115565b50600290565b905090565b90565b905090565b61017061016a610163836100c1565b809461014f565b916100cc565b5f915b8383106101805750505050565b61019661019060019284516100df565b926100f0565b92019190610173565b906101ac81604093610154565b0190565b60200190565b6101d26101cc6101c583610141565b8094610147565b9161014c565b5f915b8383106101e25750505050565b6101f86101f2600192845161019f565b926101b0565b920191906101d5565b60c09061022b61023294969593966102216101008401985f8501906100f6565b60408301906101b6565b01906100f6565b565b61025f61024b61024536600461008b565b906104f1565b610256939193610032565b93849384610201565b0390f35b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061028f90610267565b810190811067ffffffffffffffff8211176102a957604052565b610271565b906102c16102ba610032565b9283610285565b565b67ffffffffffffffff81116102d85760200290565b610271565b6102e96102ee916102c3565b6102ae565b90565b369037565b90610314610303836102dd565b9261030e84916102c3565b906102f1565b565b61032060026102f6565b90565b67ffffffffffffffff81116103385760200290565b610271565b61034961034e91610323565b6102ae565b90565b61035b60026102f6565b90565b5f5b82811061036c57505050565b602090610377610351565b8184015201610360565b9061039f61038e8361033d565b926103998491610323565b9061035e565b565b6103ab6002610381565b90565b6103b7816100cf565b036103be57565b5f80fd5b905035906103cf826103ae565b565b909291926103e66103e1826102c3565b6102ae565b93602085920283019281841161041e57915b8383106104055750505050565b6020809161041384866103c2565b8152019201916103f8565b610048565b9080601f8301121561043e5761043b916002906103d1565b90565b610040565b9092919261045861045382610323565b6102ae565b93604085920283019281841161049157915b8383106104775750505050565b60206040916104868486610423565b81520192019161046a565b610048565b9080601f830112156104b1576104ae91600290610443565b90565b610040565b9091610100828403126104ec576104e96104d2845f8501610423565b936104e08160408601610496565b9360c001610423565b90565b610038565b9061051a916104fe610316565b506105076103a1565b50610510610316565b50908101906104b6565b90919256fea264697066735822122028caa143ef92f17d346c376f9eedf27160500ba2d300f937a4baaf0d297547ef64736f6c63430008180033",
}

// VerifierHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierHelperMetaData.ABI instead.
var VerifierHelperABI = VerifierHelperMetaData.ABI

// VerifierHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifierHelperMetaData.Bin instead.
var VerifierHelperBin = VerifierHelperMetaData.Bin

// DeployVerifierHelper deploys a new Ethereum contract, binding an instance of VerifierHelper to it.
func DeployVerifierHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VerifierHelper, error) {
	parsed, err := VerifierHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierHelper{VerifierHelperCaller: VerifierHelperCaller{contract: contract}, VerifierHelperTransactor: VerifierHelperTransactor{contract: contract}, VerifierHelperFilterer: VerifierHelperFilterer{contract: contract}}, nil
}

// VerifierHelper is an auto generated Go binding around an Ethereum contract.
type VerifierHelper struct {
	VerifierHelperCaller     // Read-only binding to the contract
	VerifierHelperTransactor // Write-only binding to the contract
	VerifierHelperFilterer   // Log filterer for contract events
}

// VerifierHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierHelperSession struct {
	Contract     *VerifierHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierHelperCallerSession struct {
	Contract *VerifierHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VerifierHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierHelperTransactorSession struct {
	Contract     *VerifierHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VerifierHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierHelperRaw struct {
	Contract *VerifierHelper // Generic contract binding to access the raw methods on
}

// VerifierHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierHelperCallerRaw struct {
	Contract *VerifierHelperCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierHelperTransactorRaw struct {
	Contract *VerifierHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierHelper creates a new instance of VerifierHelper, bound to a specific deployed contract.
func NewVerifierHelper(address common.Address, backend bind.ContractBackend) (*VerifierHelper, error) {
	contract, err := bindVerifierHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierHelper{VerifierHelperCaller: VerifierHelperCaller{contract: contract}, VerifierHelperTransactor: VerifierHelperTransactor{contract: contract}, VerifierHelperFilterer: VerifierHelperFilterer{contract: contract}}, nil
}

// NewVerifierHelperCaller creates a new read-only instance of VerifierHelper, bound to a specific deployed contract.
func NewVerifierHelperCaller(address common.Address, caller bind.ContractCaller) (*VerifierHelperCaller, error) {
	contract, err := bindVerifierHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierHelperCaller{contract: contract}, nil
}

// NewVerifierHelperTransactor creates a new write-only instance of VerifierHelper, bound to a specific deployed contract.
func NewVerifierHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierHelperTransactor, error) {
	contract, err := bindVerifierHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierHelperTransactor{contract: contract}, nil
}

// NewVerifierHelperFilterer creates a new log filterer instance of VerifierHelper, bound to a specific deployed contract.
func NewVerifierHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierHelperFilterer, error) {
	contract, err := bindVerifierHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierHelperFilterer{contract: contract}, nil
}

// bindVerifierHelper binds a generic wrapper to an already deployed contract.
func bindVerifierHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierHelper *VerifierHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierHelper.Contract.VerifierHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierHelper *VerifierHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierHelper.Contract.VerifierHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierHelper *VerifierHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierHelper.Contract.VerifierHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierHelper *VerifierHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierHelper *VerifierHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierHelper *VerifierHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierHelper.Contract.contract.Transact(opts, method, params...)
}

// GetSnarkProof is a free data retrieval call binding the contract method 0xf3b9a8b6.
//
// Solidity: function getSnarkProof(bytes proof) pure returns(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC)
func (_VerifierHelper *VerifierHelperCaller) GetSnarkProof(opts *bind.CallOpts, proof []byte) (struct {
	PA [2]*big.Int
	PB [2][2]*big.Int
	PC [2]*big.Int
}, error) {
	var out []interface{}
	err := _VerifierHelper.contract.Call(opts, &out, "getSnarkProof", proof)

	outstruct := new(struct {
		PA [2]*big.Int
		PB [2][2]*big.Int
		PC [2]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PA = *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)
	outstruct.PB = *abi.ConvertType(out[1], new([2][2]*big.Int)).(*[2][2]*big.Int)
	outstruct.PC = *abi.ConvertType(out[2], new([2]*big.Int)).(*[2]*big.Int)

	return *outstruct, err

}

// GetSnarkProof is a free data retrieval call binding the contract method 0xf3b9a8b6.
//
// Solidity: function getSnarkProof(bytes proof) pure returns(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC)
func (_VerifierHelper *VerifierHelperSession) GetSnarkProof(proof []byte) (struct {
	PA [2]*big.Int
	PB [2][2]*big.Int
	PC [2]*big.Int
}, error) {
	return _VerifierHelper.Contract.GetSnarkProof(&_VerifierHelper.CallOpts, proof)
}

// GetSnarkProof is a free data retrieval call binding the contract method 0xf3b9a8b6.
//
// Solidity: function getSnarkProof(bytes proof) pure returns(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC)
func (_VerifierHelper *VerifierHelperCallerSession) GetSnarkProof(proof []byte) (struct {
	PA [2]*big.Int
	PB [2][2]*big.Int
	PC [2]*big.Int
}, error) {
	return _VerifierHelper.Contract.GetSnarkProof(&_VerifierHelper.CallOpts, proof)
}
