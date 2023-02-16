// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package forwarder

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
)

// ERC20ForwardRequestTypesERC20ForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type ERC20ForwardRequestTypesERC20ForwardRequest struct {
	From          common.Address
	To            common.Address
	Token         common.Address
	TxGas         *big.Int
	TokenGasPrice *big.Int
	BatchId       *big.Int
	BatchNonce    *big.Int
	Deadline      *big.Int
	Data          []byte
}

// ForwarderMetaData contains all meta data concerning the Forwarder contract.
var ForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"domainValue\",\"type\":\"bytes\"}],\"name\":\"DomainRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_TYPE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REQUEST_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"domains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structERC20ForwardRequestTypes.ERC20ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"executeEIP712\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structERC20ForwardRequestTypes.ERC20ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"executePersonalSign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"registerDomainSeparator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structERC20ForwardRequestTypes.ERC20ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verifyEIP712\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structERC20ForwardRequestTypes.ERC20ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verifyPersonalSign\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use ForwarderMetaData.ABI instead.
var ForwarderABI = ForwarderMetaData.ABI

// Forwarder is an auto generated Go binding around an Ethereum contract.
type Forwarder struct {
	ForwarderCaller     // Read-only binding to the contract
	ForwarderTransactor // Write-only binding to the contract
	ForwarderFilterer   // Log filterer for contract events
}

// ForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForwarderSession struct {
	Contract     *Forwarder        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForwarderCallerSession struct {
	Contract *ForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForwarderTransactorSession struct {
	Contract     *ForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForwarderRaw struct {
	Contract *Forwarder // Generic contract binding to access the raw methods on
}

// ForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForwarderCallerRaw struct {
	Contract *ForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// ForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForwarderTransactorRaw struct {
	Contract *ForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForwarder creates a new instance of Forwarder, bound to a specific deployed contract.
func NewForwarder(address common.Address, backend bind.ContractBackend) (*Forwarder, error) {
	contract, err := bindForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Forwarder{ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

// NewForwarderCaller creates a new read-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderCaller(address common.Address, caller bind.ContractCaller) (*ForwarderCaller, error) {
	contract, err := bindForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderCaller{contract: contract}, nil
}

// NewForwarderTransactor creates a new write-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ForwarderTransactor, error) {
	contract, err := bindForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderTransactor{contract: contract}, nil
}

// NewForwarderFilterer creates a new log filterer instance of Forwarder, bound to a specific deployed contract.
func NewForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ForwarderFilterer, error) {
	contract, err := bindForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForwarderFilterer{contract: contract}, nil
}

// bindForwarder binds a generic wrapper to an already deployed contract.
func bindForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForwarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.ForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAINTYPE is a free data retrieval call binding the contract method 0xc3f28abd.
//
// Solidity: function EIP712_DOMAIN_TYPE() view returns(string)
func (_Forwarder *ForwarderCaller) EIP712DOMAINTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "EIP712_DOMAIN_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712DOMAINTYPE is a free data retrieval call binding the contract method 0xc3f28abd.
//
// Solidity: function EIP712_DOMAIN_TYPE() view returns(string)
func (_Forwarder *ForwarderSession) EIP712DOMAINTYPE() (string, error) {
	return _Forwarder.Contract.EIP712DOMAINTYPE(&_Forwarder.CallOpts)
}

// EIP712DOMAINTYPE is a free data retrieval call binding the contract method 0xc3f28abd.
//
// Solidity: function EIP712_DOMAIN_TYPE() view returns(string)
func (_Forwarder *ForwarderCallerSession) EIP712DOMAINTYPE() (string, error) {
	return _Forwarder.Contract.EIP712DOMAINTYPE(&_Forwarder.CallOpts)
}

// REQUESTTYPEHASH is a free data retrieval call binding the contract method 0x9e39b73e.
//
// Solidity: function REQUEST_TYPEHASH() view returns(bytes32)
func (_Forwarder *ForwarderCaller) REQUESTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "REQUEST_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REQUESTTYPEHASH is a free data retrieval call binding the contract method 0x9e39b73e.
//
// Solidity: function REQUEST_TYPEHASH() view returns(bytes32)
func (_Forwarder *ForwarderSession) REQUESTTYPEHASH() ([32]byte, error) {
	return _Forwarder.Contract.REQUESTTYPEHASH(&_Forwarder.CallOpts)
}

// REQUESTTYPEHASH is a free data retrieval call binding the contract method 0x9e39b73e.
//
// Solidity: function REQUEST_TYPEHASH() view returns(bytes32)
func (_Forwarder *ForwarderCallerSession) REQUESTTYPEHASH() ([32]byte, error) {
	return _Forwarder.Contract.REQUESTTYPEHASH(&_Forwarder.CallOpts)
}

// Domains is a free data retrieval call binding the contract method 0xc722f177.
//
// Solidity: function domains(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderCaller) Domains(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "domains", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Domains is a free data retrieval call binding the contract method 0xc722f177.
//
// Solidity: function domains(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderSession) Domains(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.Domains(&_Forwarder.CallOpts, arg0)
}

// Domains is a free data retrieval call binding the contract method 0xc722f177.
//
// Solidity: function domains(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderCallerSession) Domains(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.Domains(&_Forwarder.CallOpts, arg0)
}

// GetNonce is a free data retrieval call binding the contract method 0x89535803.
//
// Solidity: function getNonce(address from, uint256 batchId) view returns(uint256)
func (_Forwarder *ForwarderCaller) GetNonce(opts *bind.CallOpts, from common.Address, batchId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getNonce", from, batchId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x89535803.
//
// Solidity: function getNonce(address from, uint256 batchId) view returns(uint256)
func (_Forwarder *ForwarderSession) GetNonce(from common.Address, batchId *big.Int) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, from, batchId)
}

// GetNonce is a free data retrieval call binding the contract method 0x89535803.
//
// Solidity: function getNonce(address from, uint256 batchId) view returns(uint256)
func (_Forwarder *ForwarderCallerSession) GetNonce(from common.Address, batchId *big.Int) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, from, batchId)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Forwarder *ForwarderCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Forwarder *ForwarderSession) IsOwner() (bool, error) {
	return _Forwarder.Contract.IsOwner(&_Forwarder.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Forwarder *ForwarderCallerSession) IsOwner() (bool, error) {
	return _Forwarder.Contract.IsOwner(&_Forwarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Forwarder *ForwarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Forwarder *ForwarderSession) Owner() (common.Address, error) {
	return _Forwarder.Contract.Owner(&_Forwarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Forwarder *ForwarderCallerSession) Owner() (common.Address, error) {
	return _Forwarder.Contract.Owner(&_Forwarder.CallOpts)
}

// VerifyEIP712 is a free data retrieval call binding the contract method 0xa41a03f2.
//
// Solidity: function verifyEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) view returns()
func (_Forwarder *ForwarderCaller) VerifyEIP712(opts *bind.CallOpts, req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) error {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "verifyEIP712", req, domainSeparator, sig)

	if err != nil {
		return err
	}

	return err

}

// VerifyEIP712 is a free data retrieval call binding the contract method 0xa41a03f2.
//
// Solidity: function verifyEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) view returns()
func (_Forwarder *ForwarderSession) VerifyEIP712(req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) error {
	return _Forwarder.Contract.VerifyEIP712(&_Forwarder.CallOpts, req, domainSeparator, sig)
}

// VerifyEIP712 is a free data retrieval call binding the contract method 0xa41a03f2.
//
// Solidity: function verifyEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) view returns()
func (_Forwarder *ForwarderCallerSession) VerifyEIP712(req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) error {
	return _Forwarder.Contract.VerifyEIP712(&_Forwarder.CallOpts, req, domainSeparator, sig)
}

// VerifyPersonalSign is a free data retrieval call binding the contract method 0x6e4cb075.
//
// Solidity: function verifyPersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) view returns()
func (_Forwarder *ForwarderCaller) VerifyPersonalSign(opts *bind.CallOpts, req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) error {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "verifyPersonalSign", req, sig)

	if err != nil {
		return err
	}

	return err

}

// VerifyPersonalSign is a free data retrieval call binding the contract method 0x6e4cb075.
//
// Solidity: function verifyPersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) view returns()
func (_Forwarder *ForwarderSession) VerifyPersonalSign(req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) error {
	return _Forwarder.Contract.VerifyPersonalSign(&_Forwarder.CallOpts, req, sig)
}

// VerifyPersonalSign is a free data retrieval call binding the contract method 0x6e4cb075.
//
// Solidity: function verifyPersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) view returns()
func (_Forwarder *ForwarderCallerSession) VerifyPersonalSign(req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) error {
	return _Forwarder.Contract.VerifyPersonalSign(&_Forwarder.CallOpts, req, sig)
}

// ExecuteEIP712 is a paid mutator transaction binding the contract method 0x41706c4e.
//
// Solidity: function executeEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactor) ExecuteEIP712(opts *bind.TransactOpts, req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "executeEIP712", req, domainSeparator, sig)
}

// ExecuteEIP712 is a paid mutator transaction binding the contract method 0x41706c4e.
//
// Solidity: function executeEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) returns(bool success, bytes ret)
func (_Forwarder *ForwarderSession) ExecuteEIP712(req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.ExecuteEIP712(&_Forwarder.TransactOpts, req, domainSeparator, sig)
}

// ExecuteEIP712 is a paid mutator transaction binding the contract method 0x41706c4e.
//
// Solidity: function executeEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactorSession) ExecuteEIP712(req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.ExecuteEIP712(&_Forwarder.TransactOpts, req, domainSeparator, sig)
}

// ExecutePersonalSign is a paid mutator transaction binding the contract method 0x8171e632.
//
// Solidity: function executePersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactor) ExecutePersonalSign(opts *bind.TransactOpts, req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "executePersonalSign", req, sig)
}

// ExecutePersonalSign is a paid mutator transaction binding the contract method 0x8171e632.
//
// Solidity: function executePersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) returns(bool success, bytes ret)
func (_Forwarder *ForwarderSession) ExecutePersonalSign(req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.ExecutePersonalSign(&_Forwarder.TransactOpts, req, sig)
}

// ExecutePersonalSign is a paid mutator transaction binding the contract method 0x8171e632.
//
// Solidity: function executePersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactorSession) ExecutePersonalSign(req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.ExecutePersonalSign(&_Forwarder.TransactOpts, req, sig)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Forwarder *ForwarderTransactor) RegisterDomainSeparator(opts *bind.TransactOpts, name string, version string) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "registerDomainSeparator", name, version)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Forwarder *ForwarderSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterDomainSeparator(&_Forwarder.TransactOpts, name, version)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Forwarder *ForwarderTransactorSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterDomainSeparator(&_Forwarder.TransactOpts, name, version)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Forwarder *ForwarderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Forwarder *ForwarderSession) RenounceOwnership() (*types.Transaction, error) {
	return _Forwarder.Contract.RenounceOwnership(&_Forwarder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Forwarder *ForwarderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Forwarder.Contract.RenounceOwnership(&_Forwarder.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Forwarder *ForwarderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Forwarder *ForwarderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Forwarder.Contract.TransferOwnership(&_Forwarder.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Forwarder *ForwarderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Forwarder.Contract.TransferOwnership(&_Forwarder.TransactOpts, newOwner)
}

// ForwarderDomainRegisteredIterator is returned from FilterDomainRegistered and is used to iterate over the raw logs and unpacked data for DomainRegistered events raised by the Forwarder contract.
type ForwarderDomainRegisteredIterator struct {
	Event *ForwarderDomainRegistered // Event containing the contract specifics and raw log

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
func (it *ForwarderDomainRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderDomainRegistered)
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
		it.Event = new(ForwarderDomainRegistered)
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
func (it *ForwarderDomainRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderDomainRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderDomainRegistered represents a DomainRegistered event raised by the Forwarder contract.
type ForwarderDomainRegistered struct {
	DomainSeparator [32]byte
	DomainValue     []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDomainRegistered is a free log retrieval operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Forwarder *ForwarderFilterer) FilterDomainRegistered(opts *bind.FilterOpts, domainSeparator [][32]byte) (*ForwarderDomainRegisteredIterator, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderDomainRegisteredIterator{contract: _Forwarder.contract, event: "DomainRegistered", logs: logs, sub: sub}, nil
}

// WatchDomainRegistered is a free log subscription operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Forwarder *ForwarderFilterer) WatchDomainRegistered(opts *bind.WatchOpts, sink chan<- *ForwarderDomainRegistered, domainSeparator [][32]byte) (event.Subscription, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderDomainRegistered)
				if err := _Forwarder.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
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

// ParseDomainRegistered is a log parse operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Forwarder *ForwarderFilterer) ParseDomainRegistered(log types.Log) (*ForwarderDomainRegistered, error) {
	event := new(ForwarderDomainRegistered)
	if err := _Forwarder.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForwarderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Forwarder contract.
type ForwarderOwnershipTransferredIterator struct {
	Event *ForwarderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ForwarderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderOwnershipTransferred)
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
		it.Event = new(ForwarderOwnershipTransferred)
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
func (it *ForwarderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderOwnershipTransferred represents a OwnershipTransferred event raised by the Forwarder contract.
type ForwarderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Forwarder *ForwarderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ForwarderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderOwnershipTransferredIterator{contract: _Forwarder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Forwarder *ForwarderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ForwarderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderOwnershipTransferred)
				if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Forwarder *ForwarderFilterer) ParseOwnershipTransferred(log types.Log) (*ForwarderOwnershipTransferred, error) {
	event := new(ForwarderOwnershipTransferred)
	if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
