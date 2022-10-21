// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// NeoFMetaData contains all meta data concerning the NeoF contract.
var NeoFMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INIT_CODE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NET_PROXYADMIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROX_CODE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pch\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftsoData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ftso1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ftso2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"f1prc\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recip\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_refPa\",\"type\":\"address\"}],\"name\":\"setDefProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ftso1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ftso2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ftso1percent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recip\",\"type\":\"address\"}],\"name\":\"setFtsoDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_refImp\",\"type\":\"address\"}],\"name\":\"setLiqPoolRefImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NeoFABI is the input ABI used to generate the binding from.
// Deprecated: Use NeoFMetaData.ABI instead.
var NeoFABI = NeoFMetaData.ABI

// NeoF is an auto generated Go binding around an Ethereum contract.
type NeoF struct {
	NeoFCaller     // Read-only binding to the contract
	NeoFTransactor // Write-only binding to the contract
	NeoFFilterer   // Log filterer for contract events
}

// NeoFCaller is an auto generated read-only Go binding around an Ethereum contract.
type NeoFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeoFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NeoFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeoFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NeoFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeoFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NeoFSession struct {
	Contract     *NeoF             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NeoFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NeoFCallerSession struct {
	Contract *NeoFCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NeoFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NeoFTransactorSession struct {
	Contract     *NeoFTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NeoFRaw is an auto generated low-level Go binding around an Ethereum contract.
type NeoFRaw struct {
	Contract *NeoF // Generic contract binding to access the raw methods on
}

// NeoFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NeoFCallerRaw struct {
	Contract *NeoFCaller // Generic read-only contract binding to access the raw methods on
}

// NeoFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NeoFTransactorRaw struct {
	Contract *NeoFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNeoF creates a new instance of NeoF, bound to a specific deployed contract.
func NewNeoF(address common.Address, backend bind.ContractBackend) (*NeoF, error) {
	contract, err := bindNeoF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NeoF{NeoFCaller: NeoFCaller{contract: contract}, NeoFTransactor: NeoFTransactor{contract: contract}, NeoFFilterer: NeoFFilterer{contract: contract}}, nil
}

// NewNeoFCaller creates a new read-only instance of NeoF, bound to a specific deployed contract.
func NewNeoFCaller(address common.Address, caller bind.ContractCaller) (*NeoFCaller, error) {
	contract, err := bindNeoF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NeoFCaller{contract: contract}, nil
}

// NewNeoFTransactor creates a new write-only instance of NeoF, bound to a specific deployed contract.
func NewNeoFTransactor(address common.Address, transactor bind.ContractTransactor) (*NeoFTransactor, error) {
	contract, err := bindNeoF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NeoFTransactor{contract: contract}, nil
}

// NewNeoFFilterer creates a new log filterer instance of NeoF, bound to a specific deployed contract.
func NewNeoFFilterer(address common.Address, filterer bind.ContractFilterer) (*NeoFFilterer, error) {
	contract, err := bindNeoF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NeoFFilterer{contract: contract}, nil
}

// bindNeoF binds a generic wrapper to an already deployed contract.
func bindNeoF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NeoFABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NeoF *NeoFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NeoF.Contract.NeoFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NeoF *NeoFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoF.Contract.NeoFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NeoF *NeoFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NeoF.Contract.NeoFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NeoF *NeoFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NeoF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NeoF *NeoFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NeoF *NeoFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NeoF.Contract.contract.Transact(opts, method, params...)
}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_NeoF *NeoFCaller) INITCODEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "INIT_CODE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_NeoF *NeoFSession) INITCODEHASH() ([32]byte, error) {
	return _NeoF.Contract.INITCODEHASH(&_NeoF.CallOpts)
}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_NeoF *NeoFCallerSession) INITCODEHASH() ([32]byte, error) {
	return _NeoF.Contract.INITCODEHASH(&_NeoF.CallOpts)
}

// LPIMPLEMENTATION is a free data retrieval call binding the contract method 0x07a765ae.
//
// Solidity: function LP_IMPLEMENTATION() view returns(address)
func (_NeoF *NeoFCaller) LPIMPLEMENTATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "LP_IMPLEMENTATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LPIMPLEMENTATION is a free data retrieval call binding the contract method 0x07a765ae.
//
// Solidity: function LP_IMPLEMENTATION() view returns(address)
func (_NeoF *NeoFSession) LPIMPLEMENTATION() (common.Address, error) {
	return _NeoF.Contract.LPIMPLEMENTATION(&_NeoF.CallOpts)
}

// LPIMPLEMENTATION is a free data retrieval call binding the contract method 0x07a765ae.
//
// Solidity: function LP_IMPLEMENTATION() view returns(address)
func (_NeoF *NeoFCallerSession) LPIMPLEMENTATION() (common.Address, error) {
	return _NeoF.Contract.LPIMPLEMENTATION(&_NeoF.CallOpts)
}

// NETPROXYADMIN is a free data retrieval call binding the contract method 0x07ad79af.
//
// Solidity: function NET_PROXYADMIN() view returns(address)
func (_NeoF *NeoFCaller) NETPROXYADMIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "NET_PROXYADMIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NETPROXYADMIN is a free data retrieval call binding the contract method 0x07ad79af.
//
// Solidity: function NET_PROXYADMIN() view returns(address)
func (_NeoF *NeoFSession) NETPROXYADMIN() (common.Address, error) {
	return _NeoF.Contract.NETPROXYADMIN(&_NeoF.CallOpts)
}

// NETPROXYADMIN is a free data retrieval call binding the contract method 0x07ad79af.
//
// Solidity: function NET_PROXYADMIN() view returns(address)
func (_NeoF *NeoFCallerSession) NETPROXYADMIN() (common.Address, error) {
	return _NeoF.Contract.NETPROXYADMIN(&_NeoF.CallOpts)
}

// PROXCODEHASH is a free data retrieval call binding the contract method 0xb0797877.
//
// Solidity: function PROX_CODE_HASH() view returns(bytes32 _pch)
func (_NeoF *NeoFCaller) PROXCODEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "PROX_CODE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROXCODEHASH is a free data retrieval call binding the contract method 0xb0797877.
//
// Solidity: function PROX_CODE_HASH() view returns(bytes32 _pch)
func (_NeoF *NeoFSession) PROXCODEHASH() ([32]byte, error) {
	return _NeoF.Contract.PROXCODEHASH(&_NeoF.CallOpts)
}

// PROXCODEHASH is a free data retrieval call binding the contract method 0xb0797877.
//
// Solidity: function PROX_CODE_HASH() view returns(bytes32 _pch)
func (_NeoF *NeoFCallerSession) PROXCODEHASH() ([32]byte, error) {
	return _NeoF.Contract.PROXCODEHASH(&_NeoF.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_NeoF *NeoFCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_NeoF *NeoFSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _NeoF.Contract.AllPairs(&_NeoF.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_NeoF *NeoFCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _NeoF.Contract.AllPairs(&_NeoF.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_NeoF *NeoFCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_NeoF *NeoFSession) AllPairsLength() (*big.Int, error) {
	return _NeoF.Contract.AllPairsLength(&_NeoF.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_NeoF *NeoFCallerSession) AllPairsLength() (*big.Int, error) {
	return _NeoF.Contract.AllPairsLength(&_NeoF.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_NeoF *NeoFCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_NeoF *NeoFSession) FeeTo() (common.Address, error) {
	return _NeoF.Contract.FeeTo(&_NeoF.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_NeoF *NeoFCallerSession) FeeTo() (common.Address, error) {
	return _NeoF.Contract.FeeTo(&_NeoF.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_NeoF *NeoFCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_NeoF *NeoFSession) FeeToSetter() (common.Address, error) {
	return _NeoF.Contract.FeeToSetter(&_NeoF.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_NeoF *NeoFCallerSession) FeeToSetter() (common.Address, error) {
	return _NeoF.Contract.FeeToSetter(&_NeoF.CallOpts)
}

// FtsoData is a free data retrieval call binding the contract method 0xde5c4709.
//
// Solidity: function ftsoData() view returns(address ftso1, address ftso2, uint256 f1prc, address recip)
func (_NeoF *NeoFCaller) FtsoData(opts *bind.CallOpts) (struct {
	Ftso1 common.Address
	Ftso2 common.Address
	F1prc *big.Int
	Recip common.Address
}, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "ftsoData")

	outstruct := new(struct {
		Ftso1 common.Address
		Ftso2 common.Address
		F1prc *big.Int
		Recip common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ftso1 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Ftso2 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.F1prc = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Recip = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// FtsoData is a free data retrieval call binding the contract method 0xde5c4709.
//
// Solidity: function ftsoData() view returns(address ftso1, address ftso2, uint256 f1prc, address recip)
func (_NeoF *NeoFSession) FtsoData() (struct {
	Ftso1 common.Address
	Ftso2 common.Address
	F1prc *big.Int
	Recip common.Address
}, error) {
	return _NeoF.Contract.FtsoData(&_NeoF.CallOpts)
}

// FtsoData is a free data retrieval call binding the contract method 0xde5c4709.
//
// Solidity: function ftsoData() view returns(address ftso1, address ftso2, uint256 f1prc, address recip)
func (_NeoF *NeoFCallerSession) FtsoData() (struct {
	Ftso1 common.Address
	Ftso2 common.Address
	F1prc *big.Int
	Recip common.Address
}, error) {
	return _NeoF.Contract.FtsoData(&_NeoF.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_NeoF *NeoFCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_NeoF *NeoFSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _NeoF.Contract.GetPair(&_NeoF.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_NeoF *NeoFCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _NeoF.Contract.GetPair(&_NeoF.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NeoF *NeoFCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoF.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NeoF *NeoFSession) Owner() (common.Address, error) {
	return _NeoF.Contract.Owner(&_NeoF.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NeoF *NeoFCallerSession) Owner() (common.Address, error) {
	return _NeoF.Contract.Owner(&_NeoF.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_NeoF *NeoFTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _NeoF.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_NeoF *NeoFSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.CreatePair(&_NeoF.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_NeoF *NeoFTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.CreatePair(&_NeoF.TransactOpts, tokenA, tokenB)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _feeToSetter) returns()
func (_NeoF *NeoFTransactor) Initialize(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _NeoF.contract.Transact(opts, "initialize", _feeToSetter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _feeToSetter) returns()
func (_NeoF *NeoFSession) Initialize(_feeToSetter common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.Initialize(&_NeoF.TransactOpts, _feeToSetter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _feeToSetter) returns()
func (_NeoF *NeoFTransactorSession) Initialize(_feeToSetter common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.Initialize(&_NeoF.TransactOpts, _feeToSetter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeoF *NeoFTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoF.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeoF *NeoFSession) RenounceOwnership() (*types.Transaction, error) {
	return _NeoF.Contract.RenounceOwnership(&_NeoF.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeoF *NeoFTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NeoF.Contract.RenounceOwnership(&_NeoF.TransactOpts)
}

// SetDefProxyAdmin is a paid mutator transaction binding the contract method 0x0ba99698.
//
// Solidity: function setDefProxyAdmin(address _refPa) returns()
func (_NeoF *NeoFTransactor) SetDefProxyAdmin(opts *bind.TransactOpts, _refPa common.Address) (*types.Transaction, error) {
	return _NeoF.contract.Transact(opts, "setDefProxyAdmin", _refPa)
}

// SetDefProxyAdmin is a paid mutator transaction binding the contract method 0x0ba99698.
//
// Solidity: function setDefProxyAdmin(address _refPa) returns()
func (_NeoF *NeoFSession) SetDefProxyAdmin(_refPa common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.SetDefProxyAdmin(&_NeoF.TransactOpts, _refPa)
}

// SetDefProxyAdmin is a paid mutator transaction binding the contract method 0x0ba99698.
//
// Solidity: function setDefProxyAdmin(address _refPa) returns()
func (_NeoF *NeoFTransactorSession) SetDefProxyAdmin(_refPa common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.SetDefProxyAdmin(&_NeoF.TransactOpts, _refPa)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_NeoF *NeoFTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _NeoF.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_NeoF *NeoFSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.SetFeeTo(&_NeoF.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_NeoF *NeoFTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.SetFeeTo(&_NeoF.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_NeoF *NeoFTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _NeoF.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_NeoF *NeoFSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.SetFeeToSetter(&_NeoF.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_NeoF *NeoFTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.SetFeeToSetter(&_NeoF.TransactOpts, _feeToSetter)
}

// SetFtsoDetails is a paid mutator transaction binding the contract method 0x9c12cab3.
//
// Solidity: function setFtsoDetails(address _ftso1, address _ftso2, uint256 _ftso1percent, address _recip) returns()
func (_NeoF *NeoFTransactor) SetFtsoDetails(opts *bind.TransactOpts, _ftso1 common.Address, _ftso2 common.Address, _ftso1percent *big.Int, _recip common.Address) (*types.Transaction, error) {
	return _NeoF.contract.Transact(opts, "setFtsoDetails", _ftso1, _ftso2, _ftso1percent, _recip)
}

// SetFtsoDetails is a paid mutator transaction binding the contract method 0x9c12cab3.
//
// Solidity: function setFtsoDetails(address _ftso1, address _ftso2, uint256 _ftso1percent, address _recip) returns()
func (_NeoF *NeoFSession) SetFtsoDetails(_ftso1 common.Address, _ftso2 common.Address, _ftso1percent *big.Int, _recip common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.SetFtsoDetails(&_NeoF.TransactOpts, _ftso1, _ftso2, _ftso1percent, _recip)
}

// SetFtsoDetails is a paid mutator transaction binding the contract method 0x9c12cab3.
//
// Solidity: function setFtsoDetails(address _ftso1, address _ftso2, uint256 _ftso1percent, address _recip) returns()
func (_NeoF *NeoFTransactorSession) SetFtsoDetails(_ftso1 common.Address, _ftso2 common.Address, _ftso1percent *big.Int, _recip common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.SetFtsoDetails(&_NeoF.TransactOpts, _ftso1, _ftso2, _ftso1percent, _recip)
}

// SetLiqPoolRefImplementation is a paid mutator transaction binding the contract method 0x664ccee0.
//
// Solidity: function setLiqPoolRefImplementation(address _refImp) returns()
func (_NeoF *NeoFTransactor) SetLiqPoolRefImplementation(opts *bind.TransactOpts, _refImp common.Address) (*types.Transaction, error) {
	return _NeoF.contract.Transact(opts, "setLiqPoolRefImplementation", _refImp)
}

// SetLiqPoolRefImplementation is a paid mutator transaction binding the contract method 0x664ccee0.
//
// Solidity: function setLiqPoolRefImplementation(address _refImp) returns()
func (_NeoF *NeoFSession) SetLiqPoolRefImplementation(_refImp common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.SetLiqPoolRefImplementation(&_NeoF.TransactOpts, _refImp)
}

// SetLiqPoolRefImplementation is a paid mutator transaction binding the contract method 0x664ccee0.
//
// Solidity: function setLiqPoolRefImplementation(address _refImp) returns()
func (_NeoF *NeoFTransactorSession) SetLiqPoolRefImplementation(_refImp common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.SetLiqPoolRefImplementation(&_NeoF.TransactOpts, _refImp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeoF *NeoFTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NeoF.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeoF *NeoFSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.TransferOwnership(&_NeoF.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeoF *NeoFTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NeoF.Contract.TransferOwnership(&_NeoF.TransactOpts, newOwner)
}

// NeoFOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NeoF contract.
type NeoFOwnershipTransferredIterator struct {
	Event *NeoFOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NeoFOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoFOwnershipTransferred)
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
		it.Event = new(NeoFOwnershipTransferred)
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
func (it *NeoFOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoFOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoFOwnershipTransferred represents a OwnershipTransferred event raised by the NeoF contract.
type NeoFOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NeoF *NeoFFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NeoFOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NeoF.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NeoFOwnershipTransferredIterator{contract: _NeoF.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NeoF *NeoFFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NeoFOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NeoF.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoFOwnershipTransferred)
				if err := _NeoF.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NeoF *NeoFFilterer) ParseOwnershipTransferred(log types.Log) (*NeoFOwnershipTransferred, error) {
	event := new(NeoFOwnershipTransferred)
	if err := _NeoF.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoFPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the NeoF contract.
type NeoFPairCreatedIterator struct {
	Event *NeoFPairCreated // Event containing the contract specifics and raw log

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
func (it *NeoFPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoFPairCreated)
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
		it.Event = new(NeoFPairCreated)
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
func (it *NeoFPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoFPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoFPairCreated represents a PairCreated event raised by the NeoF contract.
type NeoFPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_NeoF *NeoFFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*NeoFPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _NeoF.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &NeoFPairCreatedIterator{contract: _NeoF.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_NeoF *NeoFFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *NeoFPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _NeoF.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoFPairCreated)
				if err := _NeoF.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_NeoF *NeoFFilterer) ParsePairCreated(log types.Log) (*NeoFPairCreated, error) {
	event := new(NeoFPairCreated)
	if err := _NeoF.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
