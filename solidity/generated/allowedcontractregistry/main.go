// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package allowedcontractregistry

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

// AllowedcontractregistryMetaData contains all meta data concerning the Allowedcontractregistry contract.
var AllowedcontractregistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ADD_CONTRACT_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ALLOWED_CONTRACT_REGISTRY_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SWITCH_FLAG_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"name\":\"addAllowedContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"name\":\"isAllowedToDeploy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"name\":\"isDeployed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterAccess\",\"outputs\":[{\"internalType\":\"contractMasterAccessManagement\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"name\":\"toggleDeployedFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AllowedcontractregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AllowedcontractregistryMetaData.ABI instead.
var AllowedcontractregistryABI = AllowedcontractregistryMetaData.ABI

// Allowedcontractregistry is an auto generated Go binding around an Ethereum contract.
type Allowedcontractregistry struct {
	AllowedcontractregistryCaller     // Read-only binding to the contract
	AllowedcontractregistryTransactor // Write-only binding to the contract
	AllowedcontractregistryFilterer   // Log filterer for contract events
}

// AllowedcontractregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllowedcontractregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowedcontractregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllowedcontractregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowedcontractregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllowedcontractregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowedcontractregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllowedcontractregistrySession struct {
	Contract     *Allowedcontractregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AllowedcontractregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllowedcontractregistryCallerSession struct {
	Contract *AllowedcontractregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// AllowedcontractregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllowedcontractregistryTransactorSession struct {
	Contract     *AllowedcontractregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// AllowedcontractregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllowedcontractregistryRaw struct {
	Contract *Allowedcontractregistry // Generic contract binding to access the raw methods on
}

// AllowedcontractregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllowedcontractregistryCallerRaw struct {
	Contract *AllowedcontractregistryCaller // Generic read-only contract binding to access the raw methods on
}

// AllowedcontractregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllowedcontractregistryTransactorRaw struct {
	Contract *AllowedcontractregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllowedcontractregistry creates a new instance of Allowedcontractregistry, bound to a specific deployed contract.
func NewAllowedcontractregistry(address common.Address, backend bind.ContractBackend) (*Allowedcontractregistry, error) {
	contract, err := bindAllowedcontractregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Allowedcontractregistry{AllowedcontractregistryCaller: AllowedcontractregistryCaller{contract: contract}, AllowedcontractregistryTransactor: AllowedcontractregistryTransactor{contract: contract}, AllowedcontractregistryFilterer: AllowedcontractregistryFilterer{contract: contract}}, nil
}

// NewAllowedcontractregistryCaller creates a new read-only instance of Allowedcontractregistry, bound to a specific deployed contract.
func NewAllowedcontractregistryCaller(address common.Address, caller bind.ContractCaller) (*AllowedcontractregistryCaller, error) {
	contract, err := bindAllowedcontractregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllowedcontractregistryCaller{contract: contract}, nil
}

// NewAllowedcontractregistryTransactor creates a new write-only instance of Allowedcontractregistry, bound to a specific deployed contract.
func NewAllowedcontractregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AllowedcontractregistryTransactor, error) {
	contract, err := bindAllowedcontractregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllowedcontractregistryTransactor{contract: contract}, nil
}

// NewAllowedcontractregistryFilterer creates a new log filterer instance of Allowedcontractregistry, bound to a specific deployed contract.
func NewAllowedcontractregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AllowedcontractregistryFilterer, error) {
	contract, err := bindAllowedcontractregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllowedcontractregistryFilterer{contract: contract}, nil
}

// bindAllowedcontractregistry binds a generic wrapper to an already deployed contract.
func bindAllowedcontractregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AllowedcontractregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Allowedcontractregistry *AllowedcontractregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Allowedcontractregistry.Contract.AllowedcontractregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Allowedcontractregistry *AllowedcontractregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.AllowedcontractregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Allowedcontractregistry *AllowedcontractregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.AllowedcontractregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Allowedcontractregistry *AllowedcontractregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Allowedcontractregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Allowedcontractregistry *AllowedcontractregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Allowedcontractregistry *AllowedcontractregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.contract.Transact(opts, method, params...)
}

// ADDCONTRACTPERMISSION is a free data retrieval call binding the contract method 0x3a85451f.
//
// Solidity: function ADD_CONTRACT_PERMISSION() view returns(string)
func (_Allowedcontractregistry *AllowedcontractregistryCaller) ADDCONTRACTPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Allowedcontractregistry.contract.Call(opts, &out, "ADD_CONTRACT_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ADDCONTRACTPERMISSION is a free data retrieval call binding the contract method 0x3a85451f.
//
// Solidity: function ADD_CONTRACT_PERMISSION() view returns(string)
func (_Allowedcontractregistry *AllowedcontractregistrySession) ADDCONTRACTPERMISSION() (string, error) {
	return _Allowedcontractregistry.Contract.ADDCONTRACTPERMISSION(&_Allowedcontractregistry.CallOpts)
}

// ADDCONTRACTPERMISSION is a free data retrieval call binding the contract method 0x3a85451f.
//
// Solidity: function ADD_CONTRACT_PERMISSION() view returns(string)
func (_Allowedcontractregistry *AllowedcontractregistryCallerSession) ADDCONTRACTPERMISSION() (string, error) {
	return _Allowedcontractregistry.Contract.ADDCONTRACTPERMISSION(&_Allowedcontractregistry.CallOpts)
}

// ALLOWEDCONTRACTREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x23c6b840.
//
// Solidity: function ALLOWED_CONTRACT_REGISTRY_RESOURCE() view returns(string)
func (_Allowedcontractregistry *AllowedcontractregistryCaller) ALLOWEDCONTRACTREGISTRYRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Allowedcontractregistry.contract.Call(opts, &out, "ALLOWED_CONTRACT_REGISTRY_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ALLOWEDCONTRACTREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x23c6b840.
//
// Solidity: function ALLOWED_CONTRACT_REGISTRY_RESOURCE() view returns(string)
func (_Allowedcontractregistry *AllowedcontractregistrySession) ALLOWEDCONTRACTREGISTRYRESOURCE() (string, error) {
	return _Allowedcontractregistry.Contract.ALLOWEDCONTRACTREGISTRYRESOURCE(&_Allowedcontractregistry.CallOpts)
}

// ALLOWEDCONTRACTREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x23c6b840.
//
// Solidity: function ALLOWED_CONTRACT_REGISTRY_RESOURCE() view returns(string)
func (_Allowedcontractregistry *AllowedcontractregistryCallerSession) ALLOWEDCONTRACTREGISTRYRESOURCE() (string, error) {
	return _Allowedcontractregistry.Contract.ALLOWEDCONTRACTREGISTRYRESOURCE(&_Allowedcontractregistry.CallOpts)
}

// SWITCHFLAGPERMISSION is a free data retrieval call binding the contract method 0xbfe44f57.
//
// Solidity: function SWITCH_FLAG_PERMISSION() view returns(string)
func (_Allowedcontractregistry *AllowedcontractregistryCaller) SWITCHFLAGPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Allowedcontractregistry.contract.Call(opts, &out, "SWITCH_FLAG_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SWITCHFLAGPERMISSION is a free data retrieval call binding the contract method 0xbfe44f57.
//
// Solidity: function SWITCH_FLAG_PERMISSION() view returns(string)
func (_Allowedcontractregistry *AllowedcontractregistrySession) SWITCHFLAGPERMISSION() (string, error) {
	return _Allowedcontractregistry.Contract.SWITCHFLAGPERMISSION(&_Allowedcontractregistry.CallOpts)
}

// SWITCHFLAGPERMISSION is a free data retrieval call binding the contract method 0xbfe44f57.
//
// Solidity: function SWITCH_FLAG_PERMISSION() view returns(string)
func (_Allowedcontractregistry *AllowedcontractregistryCallerSession) SWITCHFLAGPERMISSION() (string, error) {
	return _Allowedcontractregistry.Contract.SWITCHFLAGPERMISSION(&_Allowedcontractregistry.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Allowedcontractregistry *AllowedcontractregistryCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Allowedcontractregistry.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Allowedcontractregistry *AllowedcontractregistrySession) GetInjector() (common.Address, error) {
	return _Allowedcontractregistry.Contract.GetInjector(&_Allowedcontractregistry.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Allowedcontractregistry *AllowedcontractregistryCallerSession) GetInjector() (common.Address, error) {
	return _Allowedcontractregistry.Contract.GetInjector(&_Allowedcontractregistry.CallOpts)
}

// IsAllowedToDeploy is a free data retrieval call binding the contract method 0x331d711e.
//
// Solidity: function isAllowedToDeploy(bytes32 hash_) view returns(bool)
func (_Allowedcontractregistry *AllowedcontractregistryCaller) IsAllowedToDeploy(opts *bind.CallOpts, hash_ [32]byte) (bool, error) {
	var out []interface{}
	err := _Allowedcontractregistry.contract.Call(opts, &out, "isAllowedToDeploy", hash_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToDeploy is a free data retrieval call binding the contract method 0x331d711e.
//
// Solidity: function isAllowedToDeploy(bytes32 hash_) view returns(bool)
func (_Allowedcontractregistry *AllowedcontractregistrySession) IsAllowedToDeploy(hash_ [32]byte) (bool, error) {
	return _Allowedcontractregistry.Contract.IsAllowedToDeploy(&_Allowedcontractregistry.CallOpts, hash_)
}

// IsAllowedToDeploy is a free data retrieval call binding the contract method 0x331d711e.
//
// Solidity: function isAllowedToDeploy(bytes32 hash_) view returns(bool)
func (_Allowedcontractregistry *AllowedcontractregistryCallerSession) IsAllowedToDeploy(hash_ [32]byte) (bool, error) {
	return _Allowedcontractregistry.Contract.IsAllowedToDeploy(&_Allowedcontractregistry.CallOpts, hash_)
}

// IsDeployed is a free data retrieval call binding the contract method 0x74c0ff4f.
//
// Solidity: function isDeployed(bytes32 hash_) view returns(bool)
func (_Allowedcontractregistry *AllowedcontractregistryCaller) IsDeployed(opts *bind.CallOpts, hash_ [32]byte) (bool, error) {
	var out []interface{}
	err := _Allowedcontractregistry.contract.Call(opts, &out, "isDeployed", hash_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeployed is a free data retrieval call binding the contract method 0x74c0ff4f.
//
// Solidity: function isDeployed(bytes32 hash_) view returns(bool)
func (_Allowedcontractregistry *AllowedcontractregistrySession) IsDeployed(hash_ [32]byte) (bool, error) {
	return _Allowedcontractregistry.Contract.IsDeployed(&_Allowedcontractregistry.CallOpts, hash_)
}

// IsDeployed is a free data retrieval call binding the contract method 0x74c0ff4f.
//
// Solidity: function isDeployed(bytes32 hash_) view returns(bool)
func (_Allowedcontractregistry *AllowedcontractregistryCallerSession) IsDeployed(hash_ [32]byte) (bool, error) {
	return _Allowedcontractregistry.Contract.IsDeployed(&_Allowedcontractregistry.CallOpts, hash_)
}

// MasterAccess is a free data retrieval call binding the contract method 0x4cbd1467.
//
// Solidity: function masterAccess() view returns(address)
func (_Allowedcontractregistry *AllowedcontractregistryCaller) MasterAccess(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Allowedcontractregistry.contract.Call(opts, &out, "masterAccess")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterAccess is a free data retrieval call binding the contract method 0x4cbd1467.
//
// Solidity: function masterAccess() view returns(address)
func (_Allowedcontractregistry *AllowedcontractregistrySession) MasterAccess() (common.Address, error) {
	return _Allowedcontractregistry.Contract.MasterAccess(&_Allowedcontractregistry.CallOpts)
}

// MasterAccess is a free data retrieval call binding the contract method 0x4cbd1467.
//
// Solidity: function masterAccess() view returns(address)
func (_Allowedcontractregistry *AllowedcontractregistryCallerSession) MasterAccess() (common.Address, error) {
	return _Allowedcontractregistry.Contract.MasterAccess(&_Allowedcontractregistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Allowedcontractregistry *AllowedcontractregistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Allowedcontractregistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Allowedcontractregistry *AllowedcontractregistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Allowedcontractregistry.Contract.SupportsInterface(&_Allowedcontractregistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Allowedcontractregistry *AllowedcontractregistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Allowedcontractregistry.Contract.SupportsInterface(&_Allowedcontractregistry.CallOpts, interfaceId)
}

// AddAllowedContract is a paid mutator transaction binding the contract method 0x985d0776.
//
// Solidity: function addAllowedContract(bytes32 hash_) returns()
func (_Allowedcontractregistry *AllowedcontractregistryTransactor) AddAllowedContract(opts *bind.TransactOpts, hash_ [32]byte) (*types.Transaction, error) {
	return _Allowedcontractregistry.contract.Transact(opts, "addAllowedContract", hash_)
}

// AddAllowedContract is a paid mutator transaction binding the contract method 0x985d0776.
//
// Solidity: function addAllowedContract(bytes32 hash_) returns()
func (_Allowedcontractregistry *AllowedcontractregistrySession) AddAllowedContract(hash_ [32]byte) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.AddAllowedContract(&_Allowedcontractregistry.TransactOpts, hash_)
}

// AddAllowedContract is a paid mutator transaction binding the contract method 0x985d0776.
//
// Solidity: function addAllowedContract(bytes32 hash_) returns()
func (_Allowedcontractregistry *AllowedcontractregistryTransactorSession) AddAllowedContract(hash_ [32]byte) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.AddAllowedContract(&_Allowedcontractregistry.TransactOpts, hash_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address registryAddress_, bytes ) returns()
func (_Allowedcontractregistry *AllowedcontractregistryTransactor) SetDependencies(opts *bind.TransactOpts, registryAddress_ common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Allowedcontractregistry.contract.Transact(opts, "setDependencies", registryAddress_, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address registryAddress_, bytes ) returns()
func (_Allowedcontractregistry *AllowedcontractregistrySession) SetDependencies(registryAddress_ common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.SetDependencies(&_Allowedcontractregistry.TransactOpts, registryAddress_, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address registryAddress_, bytes ) returns()
func (_Allowedcontractregistry *AllowedcontractregistryTransactorSession) SetDependencies(registryAddress_ common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.SetDependencies(&_Allowedcontractregistry.TransactOpts, registryAddress_, arg1)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Allowedcontractregistry *AllowedcontractregistryTransactor) SetInjector(opts *bind.TransactOpts, injector_ common.Address) (*types.Transaction, error) {
	return _Allowedcontractregistry.contract.Transact(opts, "setInjector", injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Allowedcontractregistry *AllowedcontractregistrySession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.SetInjector(&_Allowedcontractregistry.TransactOpts, injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Allowedcontractregistry *AllowedcontractregistryTransactorSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.SetInjector(&_Allowedcontractregistry.TransactOpts, injector_)
}

// ToggleDeployedFlag is a paid mutator transaction binding the contract method 0xffa535c1.
//
// Solidity: function toggleDeployedFlag(bytes32 hash_) returns()
func (_Allowedcontractregistry *AllowedcontractregistryTransactor) ToggleDeployedFlag(opts *bind.TransactOpts, hash_ [32]byte) (*types.Transaction, error) {
	return _Allowedcontractregistry.contract.Transact(opts, "toggleDeployedFlag", hash_)
}

// ToggleDeployedFlag is a paid mutator transaction binding the contract method 0xffa535c1.
//
// Solidity: function toggleDeployedFlag(bytes32 hash_) returns()
func (_Allowedcontractregistry *AllowedcontractregistrySession) ToggleDeployedFlag(hash_ [32]byte) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.ToggleDeployedFlag(&_Allowedcontractregistry.TransactOpts, hash_)
}

// ToggleDeployedFlag is a paid mutator transaction binding the contract method 0xffa535c1.
//
// Solidity: function toggleDeployedFlag(bytes32 hash_) returns()
func (_Allowedcontractregistry *AllowedcontractregistryTransactorSession) ToggleDeployedFlag(hash_ [32]byte) (*types.Transaction, error) {
	return _Allowedcontractregistry.Contract.ToggleDeployedFlag(&_Allowedcontractregistry.TransactOpts, hash_)
}
