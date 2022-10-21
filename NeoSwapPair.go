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

// NeoSwapPair_ftsoData is an auto generated low-level Go binding around an user-defined struct.
type NeoSwapPair_ftsoData struct {
	Ftso1 common.Address
	Ftso2 common.Address
	F1prc *big.Int
	Recip common.Address
}

// NeoLPMetaData contains all meta data concerning the NeoLP contract.
var NeoLPMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"ftso1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ftso2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"f1prc\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recip\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structNeoSwapPair._ftsoData\",\"name\":\"fd\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"info\",\"type\":\"bytes32\"}],\"name\":\"dInfo\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WNat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewardsFromFTSO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facVals\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"ftso1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ftso2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"f1prc\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recip\",\"type\":\"address\"}],\"internalType\":\"structNeoSwapPair._ftsoData\",\"name\":\"_fd\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_intRA\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fixOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDelegating\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"message\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"ftso1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ftso2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"f1prc\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recip\",\"type\":\"address\"}],\"internalType\":\"structNeoSwapPair._ftsoData\",\"name\":\"_fd\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_change\",\"type\":\"bool\"}],\"name\":\"setDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"showSettings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_ra\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_id\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_raddr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"showshit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tra\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"claimeps\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lca\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// NeoLPABI is the input ABI used to generate the binding from.
// Deprecated: Use NeoLPMetaData.ABI instead.
var NeoLPABI = NeoLPMetaData.ABI

// NeoLP is an auto generated Go binding around an Ethereum contract.
type NeoLP struct {
	NeoLPCaller     // Read-only binding to the contract
	NeoLPTransactor // Write-only binding to the contract
	NeoLPFilterer   // Log filterer for contract events
}

// NeoLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type NeoLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeoLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NeoLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeoLPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NeoLPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeoLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NeoLPSession struct {
	Contract     *NeoLP            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NeoLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NeoLPCallerSession struct {
	Contract *NeoLPCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NeoLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NeoLPTransactorSession struct {
	Contract     *NeoLPTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NeoLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type NeoLPRaw struct {
	Contract *NeoLP // Generic contract binding to access the raw methods on
}

// NeoLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NeoLPCallerRaw struct {
	Contract *NeoLPCaller // Generic read-only contract binding to access the raw methods on
}

// NeoLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NeoLPTransactorRaw struct {
	Contract *NeoLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNeoLP creates a new instance of NeoLP, bound to a specific deployed contract.
func NewNeoLP(address common.Address, backend bind.ContractBackend) (*NeoLP, error) {
	contract, err := bindNeoLP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NeoLP{NeoLPCaller: NeoLPCaller{contract: contract}, NeoLPTransactor: NeoLPTransactor{contract: contract}, NeoLPFilterer: NeoLPFilterer{contract: contract}}, nil
}

// NewNeoLPCaller creates a new read-only instance of NeoLP, bound to a specific deployed contract.
func NewNeoLPCaller(address common.Address, caller bind.ContractCaller) (*NeoLPCaller, error) {
	contract, err := bindNeoLP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NeoLPCaller{contract: contract}, nil
}

// NewNeoLPTransactor creates a new write-only instance of NeoLP, bound to a specific deployed contract.
func NewNeoLPTransactor(address common.Address, transactor bind.ContractTransactor) (*NeoLPTransactor, error) {
	contract, err := bindNeoLP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NeoLPTransactor{contract: contract}, nil
}

// NewNeoLPFilterer creates a new log filterer instance of NeoLP, bound to a specific deployed contract.
func NewNeoLPFilterer(address common.Address, filterer bind.ContractFilterer) (*NeoLPFilterer, error) {
	contract, err := bindNeoLP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NeoLPFilterer{contract: contract}, nil
}

// bindNeoLP binds a generic wrapper to an already deployed contract.
func bindNeoLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NeoLPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NeoLP *NeoLPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NeoLP.Contract.NeoLPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NeoLP *NeoLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoLP.Contract.NeoLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NeoLP *NeoLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NeoLP.Contract.NeoLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NeoLP *NeoLPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NeoLP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NeoLP *NeoLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NeoLP *NeoLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NeoLP.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_NeoLP *NeoLPCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_NeoLP *NeoLPSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _NeoLP.Contract.DOMAINSEPARATOR(&_NeoLP.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_NeoLP *NeoLPCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _NeoLP.Contract.DOMAINSEPARATOR(&_NeoLP.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_NeoLP *NeoLPCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_NeoLP *NeoLPSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _NeoLP.Contract.MINIMUMLIQUIDITY(&_NeoLP.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_NeoLP *NeoLPCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _NeoLP.Contract.MINIMUMLIQUIDITY(&_NeoLP.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_NeoLP *NeoLPCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_NeoLP *NeoLPSession) PERMITTYPEHASH() ([32]byte, error) {
	return _NeoLP.Contract.PERMITTYPEHASH(&_NeoLP.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_NeoLP *NeoLPCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _NeoLP.Contract.PERMITTYPEHASH(&_NeoLP.CallOpts)
}

// WNat is a free data retrieval call binding the contract method 0x3c0c2618.
//
// Solidity: function WNat() view returns(address)
func (_NeoLP *NeoLPCaller) WNat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "WNat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WNat is a free data retrieval call binding the contract method 0x3c0c2618.
//
// Solidity: function WNat() view returns(address)
func (_NeoLP *NeoLPSession) WNat() (common.Address, error) {
	return _NeoLP.Contract.WNat(&_NeoLP.CallOpts)
}

// WNat is a free data retrieval call binding the contract method 0x3c0c2618.
//
// Solidity: function WNat() view returns(address)
func (_NeoLP *NeoLPCallerSession) WNat() (common.Address, error) {
	return _NeoLP.Contract.WNat(&_NeoLP.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_NeoLP *NeoLPCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_NeoLP *NeoLPSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NeoLP.Contract.Allowance(&_NeoLP.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_NeoLP *NeoLPCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NeoLP.Contract.Allowance(&_NeoLP.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_NeoLP *NeoLPCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_NeoLP *NeoLPSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _NeoLP.Contract.BalanceOf(&_NeoLP.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_NeoLP *NeoLPCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _NeoLP.Contract.BalanceOf(&_NeoLP.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NeoLP *NeoLPCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NeoLP *NeoLPSession) Decimals() (uint8, error) {
	return _NeoLP.Contract.Decimals(&_NeoLP.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NeoLP *NeoLPCallerSession) Decimals() (uint8, error) {
	return _NeoLP.Contract.Decimals(&_NeoLP.CallOpts)
}

// FacVals is a free data retrieval call binding the contract method 0x2b726882.
//
// Solidity: function facVals() view returns((address,address,uint256,address) _fd, address _intRA)
func (_NeoLP *NeoLPCaller) FacVals(opts *bind.CallOpts) (struct {
	Fd    NeoSwapPair_ftsoData
	IntRA common.Address
}, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "facVals")

	outstruct := new(struct {
		Fd    NeoSwapPair_ftsoData
		IntRA common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fd = *abi.ConvertType(out[0], new(NeoSwapPair_ftsoData)).(*NeoSwapPair_ftsoData)
	outstruct.IntRA = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// FacVals is a free data retrieval call binding the contract method 0x2b726882.
//
// Solidity: function facVals() view returns((address,address,uint256,address) _fd, address _intRA)
func (_NeoLP *NeoLPSession) FacVals() (struct {
	Fd    NeoSwapPair_ftsoData
	IntRA common.Address
}, error) {
	return _NeoLP.Contract.FacVals(&_NeoLP.CallOpts)
}

// FacVals is a free data retrieval call binding the contract method 0x2b726882.
//
// Solidity: function facVals() view returns((address,address,uint256,address) _fd, address _intRA)
func (_NeoLP *NeoLPCallerSession) FacVals() (struct {
	Fd    NeoSwapPair_ftsoData
	IntRA common.Address
}, error) {
	return _NeoLP.Contract.FacVals(&_NeoLP.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_NeoLP *NeoLPCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_NeoLP *NeoLPSession) Factory() (common.Address, error) {
	return _NeoLP.Contract.Factory(&_NeoLP.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_NeoLP *NeoLPCallerSession) Factory() (common.Address, error) {
	return _NeoLP.Contract.Factory(&_NeoLP.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_NeoLP *NeoLPCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_NeoLP *NeoLPSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _NeoLP.Contract.GetReserves(&_NeoLP.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_NeoLP *NeoLPCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _NeoLP.Contract.GetReserves(&_NeoLP.CallOpts)
}

// IsDelegating is a free data retrieval call binding the contract method 0xcd3c5b4d.
//
// Solidity: function isDelegating() view returns(bool)
func (_NeoLP *NeoLPCaller) IsDelegating(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "isDelegating")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegating is a free data retrieval call binding the contract method 0xcd3c5b4d.
//
// Solidity: function isDelegating() view returns(bool)
func (_NeoLP *NeoLPSession) IsDelegating() (bool, error) {
	return _NeoLP.Contract.IsDelegating(&_NeoLP.CallOpts)
}

// IsDelegating is a free data retrieval call binding the contract method 0xcd3c5b4d.
//
// Solidity: function isDelegating() view returns(bool)
func (_NeoLP *NeoLPCallerSession) IsDelegating() (bool, error) {
	return _NeoLP.Contract.IsDelegating(&_NeoLP.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_NeoLP *NeoLPCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_NeoLP *NeoLPSession) KLast() (*big.Int, error) {
	return _NeoLP.Contract.KLast(&_NeoLP.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_NeoLP *NeoLPCallerSession) KLast() (*big.Int, error) {
	return _NeoLP.Contract.KLast(&_NeoLP.CallOpts)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() view returns(string)
func (_NeoLP *NeoLPCaller) Message(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "message")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() view returns(string)
func (_NeoLP *NeoLPSession) Message() (string, error) {
	return _NeoLP.Contract.Message(&_NeoLP.CallOpts)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() view returns(string)
func (_NeoLP *NeoLPCallerSession) Message() (string, error) {
	return _NeoLP.Contract.Message(&_NeoLP.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NeoLP *NeoLPCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NeoLP *NeoLPSession) Name() (string, error) {
	return _NeoLP.Contract.Name(&_NeoLP.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NeoLP *NeoLPCallerSession) Name() (string, error) {
	return _NeoLP.Contract.Name(&_NeoLP.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NeoLP *NeoLPCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NeoLP *NeoLPSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _NeoLP.Contract.Nonces(&_NeoLP.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NeoLP *NeoLPCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _NeoLP.Contract.Nonces(&_NeoLP.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NeoLP *NeoLPCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NeoLP *NeoLPSession) Owner() (common.Address, error) {
	return _NeoLP.Contract.Owner(&_NeoLP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NeoLP *NeoLPCallerSession) Owner() (common.Address, error) {
	return _NeoLP.Contract.Owner(&_NeoLP.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_NeoLP *NeoLPCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_NeoLP *NeoLPSession) Price0CumulativeLast() (*big.Int, error) {
	return _NeoLP.Contract.Price0CumulativeLast(&_NeoLP.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_NeoLP *NeoLPCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _NeoLP.Contract.Price0CumulativeLast(&_NeoLP.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_NeoLP *NeoLPCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_NeoLP *NeoLPSession) Price1CumulativeLast() (*big.Int, error) {
	return _NeoLP.Contract.Price1CumulativeLast(&_NeoLP.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_NeoLP *NeoLPCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _NeoLP.Contract.Price1CumulativeLast(&_NeoLP.CallOpts)
}

// ShowSettings is a free data retrieval call binding the contract method 0xdc12042f.
//
// Solidity: function showSettings() view returns(uint256 _ra, bool _id, address _raddr)
func (_NeoLP *NeoLPCaller) ShowSettings(opts *bind.CallOpts) (struct {
	Ra    *big.Int
	Id    bool
	Raddr common.Address
}, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "showSettings")

	outstruct := new(struct {
		Ra    *big.Int
		Id    bool
		Raddr common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ra = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Id = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Raddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ShowSettings is a free data retrieval call binding the contract method 0xdc12042f.
//
// Solidity: function showSettings() view returns(uint256 _ra, bool _id, address _raddr)
func (_NeoLP *NeoLPSession) ShowSettings() (struct {
	Ra    *big.Int
	Id    bool
	Raddr common.Address
}, error) {
	return _NeoLP.Contract.ShowSettings(&_NeoLP.CallOpts)
}

// ShowSettings is a free data retrieval call binding the contract method 0xdc12042f.
//
// Solidity: function showSettings() view returns(uint256 _ra, bool _id, address _raddr)
func (_NeoLP *NeoLPCallerSession) ShowSettings() (struct {
	Ra    *big.Int
	Id    bool
	Raddr common.Address
}, error) {
	return _NeoLP.Contract.ShowSettings(&_NeoLP.CallOpts)
}

// Showshit is a free data retrieval call binding the contract method 0xfd00e9f9.
//
// Solidity: function showshit() view returns(uint256 tra, uint256[] claimeps, uint256 lca)
func (_NeoLP *NeoLPCaller) Showshit(opts *bind.CallOpts) (struct {
	Tra      *big.Int
	Claimeps []*big.Int
	Lca      *big.Int
}, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "showshit")

	outstruct := new(struct {
		Tra      *big.Int
		Claimeps []*big.Int
		Lca      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tra = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Claimeps = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Lca = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Showshit is a free data retrieval call binding the contract method 0xfd00e9f9.
//
// Solidity: function showshit() view returns(uint256 tra, uint256[] claimeps, uint256 lca)
func (_NeoLP *NeoLPSession) Showshit() (struct {
	Tra      *big.Int
	Claimeps []*big.Int
	Lca      *big.Int
}, error) {
	return _NeoLP.Contract.Showshit(&_NeoLP.CallOpts)
}

// Showshit is a free data retrieval call binding the contract method 0xfd00e9f9.
//
// Solidity: function showshit() view returns(uint256 tra, uint256[] claimeps, uint256 lca)
func (_NeoLP *NeoLPCallerSession) Showshit() (struct {
	Tra      *big.Int
	Claimeps []*big.Int
	Lca      *big.Int
}, error) {
	return _NeoLP.Contract.Showshit(&_NeoLP.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NeoLP *NeoLPCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NeoLP *NeoLPSession) Symbol() (string, error) {
	return _NeoLP.Contract.Symbol(&_NeoLP.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NeoLP *NeoLPCallerSession) Symbol() (string, error) {
	return _NeoLP.Contract.Symbol(&_NeoLP.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_NeoLP *NeoLPCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_NeoLP *NeoLPSession) Token0() (common.Address, error) {
	return _NeoLP.Contract.Token0(&_NeoLP.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_NeoLP *NeoLPCallerSession) Token0() (common.Address, error) {
	return _NeoLP.Contract.Token0(&_NeoLP.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_NeoLP *NeoLPCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_NeoLP *NeoLPSession) Token1() (common.Address, error) {
	return _NeoLP.Contract.Token1(&_NeoLP.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_NeoLP *NeoLPCallerSession) Token1() (common.Address, error) {
	return _NeoLP.Contract.Token1(&_NeoLP.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NeoLP *NeoLPCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NeoLP.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NeoLP *NeoLPSession) TotalSupply() (*big.Int, error) {
	return _NeoLP.Contract.TotalSupply(&_NeoLP.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NeoLP *NeoLPCallerSession) TotalSupply() (*big.Int, error) {
	return _NeoLP.Contract.TotalSupply(&_NeoLP.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NeoLP *NeoLPTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NeoLP *NeoLPSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeoLP.Contract.Approve(&_NeoLP.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NeoLP *NeoLPTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeoLP.Contract.Approve(&_NeoLP.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_NeoLP *NeoLPTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_NeoLP *NeoLPSession) Burn(to common.Address) (*types.Transaction, error) {
	return _NeoLP.Contract.Burn(&_NeoLP.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_NeoLP *NeoLPTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _NeoLP.Contract.Burn(&_NeoLP.TransactOpts, to)
}

// ClaimRewardsFromFTSO is a paid mutator transaction binding the contract method 0xb35c66d2.
//
// Solidity: function claimRewardsFromFTSO() returns()
func (_NeoLP *NeoLPTransactor) ClaimRewardsFromFTSO(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "claimRewardsFromFTSO")
}

// ClaimRewardsFromFTSO is a paid mutator transaction binding the contract method 0xb35c66d2.
//
// Solidity: function claimRewardsFromFTSO() returns()
func (_NeoLP *NeoLPSession) ClaimRewardsFromFTSO() (*types.Transaction, error) {
	return _NeoLP.Contract.ClaimRewardsFromFTSO(&_NeoLP.TransactOpts)
}

// ClaimRewardsFromFTSO is a paid mutator transaction binding the contract method 0xb35c66d2.
//
// Solidity: function claimRewardsFromFTSO() returns()
func (_NeoLP *NeoLPTransactorSession) ClaimRewardsFromFTSO() (*types.Transaction, error) {
	return _NeoLP.Contract.ClaimRewardsFromFTSO(&_NeoLP.TransactOpts)
}

// FixOwner is a paid mutator transaction binding the contract method 0x08856caf.
//
// Solidity: function fixOwner() returns()
func (_NeoLP *NeoLPTransactor) FixOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "fixOwner")
}

// FixOwner is a paid mutator transaction binding the contract method 0x08856caf.
//
// Solidity: function fixOwner() returns()
func (_NeoLP *NeoLPSession) FixOwner() (*types.Transaction, error) {
	return _NeoLP.Contract.FixOwner(&_NeoLP.TransactOpts)
}

// FixOwner is a paid mutator transaction binding the contract method 0x08856caf.
//
// Solidity: function fixOwner() returns()
func (_NeoLP *NeoLPTransactorSession) FixOwner() (*types.Transaction, error) {
	return _NeoLP.Contract.FixOwner(&_NeoLP.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_NeoLP *NeoLPTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_NeoLP *NeoLPSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _NeoLP.Contract.Initialize(&_NeoLP.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_NeoLP *NeoLPTransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _NeoLP.Contract.Initialize(&_NeoLP.TransactOpts, _token0, _token1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_NeoLP *NeoLPTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_NeoLP *NeoLPSession) Mint(to common.Address) (*types.Transaction, error) {
	return _NeoLP.Contract.Mint(&_NeoLP.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_NeoLP *NeoLPTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _NeoLP.Contract.Mint(&_NeoLP.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_NeoLP *NeoLPTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_NeoLP *NeoLPSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NeoLP.Contract.Permit(&_NeoLP.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_NeoLP *NeoLPTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NeoLP.Contract.Permit(&_NeoLP.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeoLP *NeoLPTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeoLP *NeoLPSession) RenounceOwnership() (*types.Transaction, error) {
	return _NeoLP.Contract.RenounceOwnership(&_NeoLP.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeoLP *NeoLPTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NeoLP.Contract.RenounceOwnership(&_NeoLP.TransactOpts)
}

// SetDelegation is a paid mutator transaction binding the contract method 0x1225a796.
//
// Solidity: function setDelegation((address,address,uint256,address) _fd, bool _change) returns()
func (_NeoLP *NeoLPTransactor) SetDelegation(opts *bind.TransactOpts, _fd NeoSwapPair_ftsoData, _change bool) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "setDelegation", _fd, _change)
}

// SetDelegation is a paid mutator transaction binding the contract method 0x1225a796.
//
// Solidity: function setDelegation((address,address,uint256,address) _fd, bool _change) returns()
func (_NeoLP *NeoLPSession) SetDelegation(_fd NeoSwapPair_ftsoData, _change bool) (*types.Transaction, error) {
	return _NeoLP.Contract.SetDelegation(&_NeoLP.TransactOpts, _fd, _change)
}

// SetDelegation is a paid mutator transaction binding the contract method 0x1225a796.
//
// Solidity: function setDelegation((address,address,uint256,address) _fd, bool _change) returns()
func (_NeoLP *NeoLPTransactorSession) SetDelegation(_fd NeoSwapPair_ftsoData, _change bool) (*types.Transaction, error) {
	return _NeoLP.Contract.SetDelegation(&_NeoLP.TransactOpts, _fd, _change)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_NeoLP *NeoLPTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_NeoLP *NeoLPSession) Skim(to common.Address) (*types.Transaction, error) {
	return _NeoLP.Contract.Skim(&_NeoLP.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_NeoLP *NeoLPTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _NeoLP.Contract.Skim(&_NeoLP.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_NeoLP *NeoLPTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_NeoLP *NeoLPSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _NeoLP.Contract.Swap(&_NeoLP.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_NeoLP *NeoLPTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _NeoLP.Contract.Swap(&_NeoLP.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_NeoLP *NeoLPTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_NeoLP *NeoLPSession) Sync() (*types.Transaction, error) {
	return _NeoLP.Contract.Sync(&_NeoLP.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_NeoLP *NeoLPTransactorSession) Sync() (*types.Transaction, error) {
	return _NeoLP.Contract.Sync(&_NeoLP.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NeoLP *NeoLPTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NeoLP *NeoLPSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeoLP.Contract.Transfer(&_NeoLP.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NeoLP *NeoLPTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeoLP.Contract.Transfer(&_NeoLP.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NeoLP *NeoLPTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NeoLP *NeoLPSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeoLP.Contract.TransferFrom(&_NeoLP.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NeoLP *NeoLPTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeoLP.Contract.TransferFrom(&_NeoLP.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeoLP *NeoLPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeoLP *NeoLPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NeoLP.Contract.TransferOwnership(&_NeoLP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeoLP *NeoLPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NeoLP.Contract.TransferOwnership(&_NeoLP.TransactOpts, newOwner)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_NeoLP *NeoLPTransactor) WithdrawRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoLP.contract.Transact(opts, "withdrawRewards")
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_NeoLP *NeoLPSession) WithdrawRewards() (*types.Transaction, error) {
	return _NeoLP.Contract.WithdrawRewards(&_NeoLP.TransactOpts)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_NeoLP *NeoLPTransactorSession) WithdrawRewards() (*types.Transaction, error) {
	return _NeoLP.Contract.WithdrawRewards(&_NeoLP.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NeoLP *NeoLPTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NeoLP.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NeoLP *NeoLPSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NeoLP.Contract.Fallback(&_NeoLP.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NeoLP *NeoLPTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NeoLP.Contract.Fallback(&_NeoLP.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NeoLP *NeoLPTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeoLP.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NeoLP *NeoLPSession) Receive() (*types.Transaction, error) {
	return _NeoLP.Contract.Receive(&_NeoLP.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NeoLP *NeoLPTransactorSession) Receive() (*types.Transaction, error) {
	return _NeoLP.Contract.Receive(&_NeoLP.TransactOpts)
}

// NeoLPApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NeoLP contract.
type NeoLPApprovalIterator struct {
	Event *NeoLPApproval // Event containing the contract specifics and raw log

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
func (it *NeoLPApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoLPApproval)
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
		it.Event = new(NeoLPApproval)
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
func (it *NeoLPApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoLPApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoLPApproval represents a Approval event raised by the NeoLP contract.
type NeoLPApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NeoLP *NeoLPFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NeoLPApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NeoLP.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NeoLPApprovalIterator{contract: _NeoLP.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NeoLP *NeoLPFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NeoLPApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NeoLP.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoLPApproval)
				if err := _NeoLP.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NeoLP *NeoLPFilterer) ParseApproval(log types.Log) (*NeoLPApproval, error) {
	event := new(NeoLPApproval)
	if err := _NeoLP.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoLPBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the NeoLP contract.
type NeoLPBurnIterator struct {
	Event *NeoLPBurn // Event containing the contract specifics and raw log

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
func (it *NeoLPBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoLPBurn)
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
		it.Event = new(NeoLPBurn)
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
func (it *NeoLPBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoLPBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoLPBurn represents a Burn event raised by the NeoLP contract.
type NeoLPBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_NeoLP *NeoLPFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*NeoLPBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NeoLP.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NeoLPBurnIterator{contract: _NeoLP.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_NeoLP *NeoLPFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *NeoLPBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NeoLP.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoLPBurn)
				if err := _NeoLP.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_NeoLP *NeoLPFilterer) ParseBurn(log types.Log) (*NeoLPBurn, error) {
	event := new(NeoLPBurn)
	if err := _NeoLP.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoLPMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the NeoLP contract.
type NeoLPMintIterator struct {
	Event *NeoLPMint // Event containing the contract specifics and raw log

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
func (it *NeoLPMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoLPMint)
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
		it.Event = new(NeoLPMint)
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
func (it *NeoLPMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoLPMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoLPMint represents a Mint event raised by the NeoLP contract.
type NeoLPMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_NeoLP *NeoLPFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*NeoLPMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NeoLP.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &NeoLPMintIterator{contract: _NeoLP.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_NeoLP *NeoLPFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *NeoLPMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NeoLP.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoLPMint)
				if err := _NeoLP.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_NeoLP *NeoLPFilterer) ParseMint(log types.Log) (*NeoLPMint, error) {
	event := new(NeoLPMint)
	if err := _NeoLP.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoLPOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NeoLP contract.
type NeoLPOwnershipTransferredIterator struct {
	Event *NeoLPOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NeoLPOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoLPOwnershipTransferred)
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
		it.Event = new(NeoLPOwnershipTransferred)
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
func (it *NeoLPOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoLPOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoLPOwnershipTransferred represents a OwnershipTransferred event raised by the NeoLP contract.
type NeoLPOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NeoLP *NeoLPFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NeoLPOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NeoLP.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NeoLPOwnershipTransferredIterator{contract: _NeoLP.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NeoLP *NeoLPFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NeoLPOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NeoLP.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoLPOwnershipTransferred)
				if err := _NeoLP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NeoLP *NeoLPFilterer) ParseOwnershipTransferred(log types.Log) (*NeoLPOwnershipTransferred, error) {
	event := new(NeoLPOwnershipTransferred)
	if err := _NeoLP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoLPSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the NeoLP contract.
type NeoLPSwapIterator struct {
	Event *NeoLPSwap // Event containing the contract specifics and raw log

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
func (it *NeoLPSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoLPSwap)
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
		it.Event = new(NeoLPSwap)
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
func (it *NeoLPSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoLPSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoLPSwap represents a Swap event raised by the NeoLP contract.
type NeoLPSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_NeoLP *NeoLPFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*NeoLPSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NeoLP.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NeoLPSwapIterator{contract: _NeoLP.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_NeoLP *NeoLPFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *NeoLPSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NeoLP.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoLPSwap)
				if err := _NeoLP.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_NeoLP *NeoLPFilterer) ParseSwap(log types.Log) (*NeoLPSwap, error) {
	event := new(NeoLPSwap)
	if err := _NeoLP.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoLPSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the NeoLP contract.
type NeoLPSyncIterator struct {
	Event *NeoLPSync // Event containing the contract specifics and raw log

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
func (it *NeoLPSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoLPSync)
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
		it.Event = new(NeoLPSync)
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
func (it *NeoLPSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoLPSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoLPSync represents a Sync event raised by the NeoLP contract.
type NeoLPSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_NeoLP *NeoLPFilterer) FilterSync(opts *bind.FilterOpts) (*NeoLPSyncIterator, error) {

	logs, sub, err := _NeoLP.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &NeoLPSyncIterator{contract: _NeoLP.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_NeoLP *NeoLPFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *NeoLPSync) (event.Subscription, error) {

	logs, sub, err := _NeoLP.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoLPSync)
				if err := _NeoLP.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_NeoLP *NeoLPFilterer) ParseSync(log types.Log) (*NeoLPSync, error) {
	event := new(NeoLPSync)
	if err := _NeoLP.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoLPTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NeoLP contract.
type NeoLPTransferIterator struct {
	Event *NeoLPTransfer // Event containing the contract specifics and raw log

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
func (it *NeoLPTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoLPTransfer)
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
		it.Event = new(NeoLPTransfer)
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
func (it *NeoLPTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoLPTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoLPTransfer represents a Transfer event raised by the NeoLP contract.
type NeoLPTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NeoLP *NeoLPFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NeoLPTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NeoLP.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NeoLPTransferIterator{contract: _NeoLP.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NeoLP *NeoLPFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NeoLPTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NeoLP.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoLPTransfer)
				if err := _NeoLP.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NeoLP *NeoLPFilterer) ParseTransfer(log types.Log) (*NeoLPTransfer, error) {
	event := new(NeoLPTransfer)
	if err := _NeoLP.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NeoLPDInfoIterator is returned from FilterDInfo and is used to iterate over the raw logs and unpacked data for DInfo events raised by the NeoLP contract.
type NeoLPDInfoIterator struct {
	Event *NeoLPDInfo // Event containing the contract specifics and raw log

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
func (it *NeoLPDInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeoLPDInfo)
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
		it.Event = new(NeoLPDInfo)
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
func (it *NeoLPDInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeoLPDInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeoLPDInfo represents a DInfo event raised by the NeoLP contract.
type NeoLPDInfo struct {
	Fd   NeoSwapPair_ftsoData
	Info [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDInfo is a free log retrieval operation binding the contract event 0x6b85700bafc133797158cf9265f14de360654ae02a0f644c329bff0334784e99.
//
// Solidity: event dInfo((address,address,uint256,address) fd, bytes32 info)
func (_NeoLP *NeoLPFilterer) FilterDInfo(opts *bind.FilterOpts) (*NeoLPDInfoIterator, error) {

	logs, sub, err := _NeoLP.contract.FilterLogs(opts, "dInfo")
	if err != nil {
		return nil, err
	}
	return &NeoLPDInfoIterator{contract: _NeoLP.contract, event: "dInfo", logs: logs, sub: sub}, nil
}

// WatchDInfo is a free log subscription operation binding the contract event 0x6b85700bafc133797158cf9265f14de360654ae02a0f644c329bff0334784e99.
//
// Solidity: event dInfo((address,address,uint256,address) fd, bytes32 info)
func (_NeoLP *NeoLPFilterer) WatchDInfo(opts *bind.WatchOpts, sink chan<- *NeoLPDInfo) (event.Subscription, error) {

	logs, sub, err := _NeoLP.contract.WatchLogs(opts, "dInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeoLPDInfo)
				if err := _NeoLP.contract.UnpackLog(event, "dInfo", log); err != nil {
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

// ParseDInfo is a log parse operation binding the contract event 0x6b85700bafc133797158cf9265f14de360654ae02a0f644c329bff0334784e99.
//
// Solidity: event dInfo((address,address,uint256,address) fd, bytes32 info)
func (_NeoLP *NeoLPFilterer) ParseDInfo(log types.Log) (*NeoLPDInfo, error) {
	event := new(NeoLPDInfo)
	if err := _NeoLP.contract.UnpackLog(event, "dInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
