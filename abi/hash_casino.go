// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// HashCasinoBetInfo is an auto generated low-level Go binding around an user-defined struct.
type HashCasinoBetInfo struct {
	GameType *big.Int
	BetType  *big.Int
	Position *big.Int
	Amount   *big.Int
	Owner    common.Address
}

// StructnameMetaData contains all meta data concerning the Structname contract.
var StructnameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BasicError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerCheckError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ErrorLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiple\",\"type\":\"uint256\"}],\"name\":\"LoserLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReceiveLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiple\",\"type\":\"uint256\"}],\"name\":\"WinnerLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBetInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gameType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structHashCasino.BetInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBetStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrevNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouletteNums\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"participate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pickWinners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nums\",\"type\":\"uint256[]\"}],\"name\":\"setRouletteNums\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StructnameABI is the input ABI used to generate the binding from.
// Deprecated: Use StructnameMetaData.ABI instead.
var StructnameABI = StructnameMetaData.ABI

// Structname is an auto generated Go binding around an Ethereum contract.
type Structname struct {
	StructnameCaller     // Read-only binding to the contract
	StructnameTransactor // Write-only binding to the contract
	StructnameFilterer   // Log filterer for contract events
}

// StructnameCaller is an auto generated read-only Go binding around an Ethereum contract.
type StructnameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructnameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StructnameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructnameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StructnameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructnameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StructnameSession struct {
	Contract     *Structname       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StructnameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StructnameCallerSession struct {
	Contract *StructnameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StructnameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StructnameTransactorSession struct {
	Contract     *StructnameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StructnameRaw is an auto generated low-level Go binding around an Ethereum contract.
type StructnameRaw struct {
	Contract *Structname // Generic contract binding to access the raw methods on
}

// StructnameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StructnameCallerRaw struct {
	Contract *StructnameCaller // Generic read-only contract binding to access the raw methods on
}

// StructnameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StructnameTransactorRaw struct {
	Contract *StructnameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStructname creates a new instance of Structname, bound to a specific deployed contract.
func NewStructname(address common.Address, backend bind.ContractBackend) (*Structname, error) {
	contract, err := bindStructname(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Structname{StructnameCaller: StructnameCaller{contract: contract}, StructnameTransactor: StructnameTransactor{contract: contract}, StructnameFilterer: StructnameFilterer{contract: contract}}, nil
}

// NewStructnameCaller creates a new read-only instance of Structname, bound to a specific deployed contract.
func NewStructnameCaller(address common.Address, caller bind.ContractCaller) (*StructnameCaller, error) {
	contract, err := bindStructname(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StructnameCaller{contract: contract}, nil
}

// NewStructnameTransactor creates a new write-only instance of Structname, bound to a specific deployed contract.
func NewStructnameTransactor(address common.Address, transactor bind.ContractTransactor) (*StructnameTransactor, error) {
	contract, err := bindStructname(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StructnameTransactor{contract: contract}, nil
}

// NewStructnameFilterer creates a new log filterer instance of Structname, bound to a specific deployed contract.
func NewStructnameFilterer(address common.Address, filterer bind.ContractFilterer) (*StructnameFilterer, error) {
	contract, err := bindStructname(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StructnameFilterer{contract: contract}, nil
}

// bindStructname binds a generic wrapper to an already deployed contract.
func bindStructname(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StructnameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Structname *StructnameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Structname.Contract.StructnameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Structname *StructnameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structname.Contract.StructnameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Structname *StructnameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Structname.Contract.StructnameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Structname *StructnameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Structname.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Structname *StructnameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structname.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Structname *StructnameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Structname.Contract.contract.Transact(opts, method, params...)
}

// GetBetInfos is a free data retrieval call binding the contract method 0x54cc5a0d.
//
// Solidity: function getBetInfos() view returns((uint256,uint256,uint256,uint256,address)[])
func (_Structname *StructnameCaller) GetBetInfos(opts *bind.CallOpts) ([]HashCasinoBetInfo, error) {
	var out []interface{}
	err := _Structname.contract.Call(opts, &out, "getBetInfos")

	if err != nil {
		return *new([]HashCasinoBetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]HashCasinoBetInfo)).(*[]HashCasinoBetInfo)

	return out0, err

}

// GetBetInfos is a free data retrieval call binding the contract method 0x54cc5a0d.
//
// Solidity: function getBetInfos() view returns((uint256,uint256,uint256,uint256,address)[])
func (_Structname *StructnameSession) GetBetInfos() ([]HashCasinoBetInfo, error) {
	return _Structname.Contract.GetBetInfos(&_Structname.CallOpts)
}

// GetBetInfos is a free data retrieval call binding the contract method 0x54cc5a0d.
//
// Solidity: function getBetInfos() view returns((uint256,uint256,uint256,uint256,address)[])
func (_Structname *StructnameCallerSession) GetBetInfos() ([]HashCasinoBetInfo, error) {
	return _Structname.Contract.GetBetInfos(&_Structname.CallOpts)
}

// GetBetStatus is a free data retrieval call binding the contract method 0x1fa859e8.
//
// Solidity: function getBetStatus() view returns(bool)
func (_Structname *StructnameCaller) GetBetStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Structname.contract.Call(opts, &out, "getBetStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBetStatus is a free data retrieval call binding the contract method 0x1fa859e8.
//
// Solidity: function getBetStatus() view returns(bool)
func (_Structname *StructnameSession) GetBetStatus() (bool, error) {
	return _Structname.Contract.GetBetStatus(&_Structname.CallOpts)
}

// GetBetStatus is a free data retrieval call binding the contract method 0x1fa859e8.
//
// Solidity: function getBetStatus() view returns(bool)
func (_Structname *StructnameCallerSession) GetBetStatus() (bool, error) {
	return _Structname.Contract.GetBetStatus(&_Structname.CallOpts)
}

// GetPoolAmount is a free data retrieval call binding the contract method 0x4ab4ba42.
//
// Solidity: function getPoolAmount() view returns(uint256)
func (_Structname *StructnameCaller) GetPoolAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Structname.contract.Call(opts, &out, "getPoolAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAmount is a free data retrieval call binding the contract method 0x4ab4ba42.
//
// Solidity: function getPoolAmount() view returns(uint256)
func (_Structname *StructnameSession) GetPoolAmount() (*big.Int, error) {
	return _Structname.Contract.GetPoolAmount(&_Structname.CallOpts)
}

// GetPoolAmount is a free data retrieval call binding the contract method 0x4ab4ba42.
//
// Solidity: function getPoolAmount() view returns(uint256)
func (_Structname *StructnameCallerSession) GetPoolAmount() (*big.Int, error) {
	return _Structname.Contract.GetPoolAmount(&_Structname.CallOpts)
}

// GetPrevNum is a free data retrieval call binding the contract method 0x27a6586b.
//
// Solidity: function getPrevNum() view returns(uint256)
func (_Structname *StructnameCaller) GetPrevNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Structname.contract.Call(opts, &out, "getPrevNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrevNum is a free data retrieval call binding the contract method 0x27a6586b.
//
// Solidity: function getPrevNum() view returns(uint256)
func (_Structname *StructnameSession) GetPrevNum() (*big.Int, error) {
	return _Structname.Contract.GetPrevNum(&_Structname.CallOpts)
}

// GetPrevNum is a free data retrieval call binding the contract method 0x27a6586b.
//
// Solidity: function getPrevNum() view returns(uint256)
func (_Structname *StructnameCallerSession) GetPrevNum() (*big.Int, error) {
	return _Structname.Contract.GetPrevNum(&_Structname.CallOpts)
}

// GetRouletteNums is a free data retrieval call binding the contract method 0x633bf639.
//
// Solidity: function getRouletteNums() view returns(uint256[])
func (_Structname *StructnameCaller) GetRouletteNums(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Structname.contract.Call(opts, &out, "getRouletteNums")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRouletteNums is a free data retrieval call binding the contract method 0x633bf639.
//
// Solidity: function getRouletteNums() view returns(uint256[])
func (_Structname *StructnameSession) GetRouletteNums() ([]*big.Int, error) {
	return _Structname.Contract.GetRouletteNums(&_Structname.CallOpts)
}

// GetRouletteNums is a free data retrieval call binding the contract method 0x633bf639.
//
// Solidity: function getRouletteNums() view returns(uint256[])
func (_Structname *StructnameCallerSession) GetRouletteNums() ([]*big.Int, error) {
	return _Structname.Contract.GetRouletteNums(&_Structname.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Structname *StructnameTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structname.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Structname *StructnameSession) Deposit() (*types.Transaction, error) {
	return _Structname.Contract.Deposit(&_Structname.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Structname *StructnameTransactorSession) Deposit() (*types.Transaction, error) {
	return _Structname.Contract.Deposit(&_Structname.TransactOpts)
}

// Participate is a paid mutator transaction binding the contract method 0xaf4e476f.
//
// Solidity: function participate(uint256 gameType, uint256 betType, uint256 position, uint256 amount) payable returns()
func (_Structname *StructnameTransactor) Participate(opts *bind.TransactOpts, gameType *big.Int, betType *big.Int, position *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Structname.contract.Transact(opts, "participate", gameType, betType, position, amount)
}

// Participate is a paid mutator transaction binding the contract method 0xaf4e476f.
//
// Solidity: function participate(uint256 gameType, uint256 betType, uint256 position, uint256 amount) payable returns()
func (_Structname *StructnameSession) Participate(gameType *big.Int, betType *big.Int, position *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Structname.Contract.Participate(&_Structname.TransactOpts, gameType, betType, position, amount)
}

// Participate is a paid mutator transaction binding the contract method 0xaf4e476f.
//
// Solidity: function participate(uint256 gameType, uint256 betType, uint256 position, uint256 amount) payable returns()
func (_Structname *StructnameTransactorSession) Participate(gameType *big.Int, betType *big.Int, position *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Structname.Contract.Participate(&_Structname.TransactOpts, gameType, betType, position, amount)
}

// PickWinners is a paid mutator transaction binding the contract method 0xdc07820f.
//
// Solidity: function pickWinners() returns()
func (_Structname *StructnameTransactor) PickWinners(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structname.contract.Transact(opts, "pickWinners")
}

// PickWinners is a paid mutator transaction binding the contract method 0xdc07820f.
//
// Solidity: function pickWinners() returns()
func (_Structname *StructnameSession) PickWinners() (*types.Transaction, error) {
	return _Structname.Contract.PickWinners(&_Structname.TransactOpts)
}

// PickWinners is a paid mutator transaction binding the contract method 0xdc07820f.
//
// Solidity: function pickWinners() returns()
func (_Structname *StructnameTransactorSession) PickWinners() (*types.Transaction, error) {
	return _Structname.Contract.PickWinners(&_Structname.TransactOpts)
}

// SetRouletteNums is a paid mutator transaction binding the contract method 0xe16b4603.
//
// Solidity: function setRouletteNums(uint256[] nums) returns()
func (_Structname *StructnameTransactor) SetRouletteNums(opts *bind.TransactOpts, nums []*big.Int) (*types.Transaction, error) {
	return _Structname.contract.Transact(opts, "setRouletteNums", nums)
}

// SetRouletteNums is a paid mutator transaction binding the contract method 0xe16b4603.
//
// Solidity: function setRouletteNums(uint256[] nums) returns()
func (_Structname *StructnameSession) SetRouletteNums(nums []*big.Int) (*types.Transaction, error) {
	return _Structname.Contract.SetRouletteNums(&_Structname.TransactOpts, nums)
}

// SetRouletteNums is a paid mutator transaction binding the contract method 0xe16b4603.
//
// Solidity: function setRouletteNums(uint256[] nums) returns()
func (_Structname *StructnameTransactorSession) SetRouletteNums(nums []*big.Int) (*types.Transaction, error) {
	return _Structname.Contract.SetRouletteNums(&_Structname.TransactOpts, nums)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Structname *StructnameTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structname.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Structname *StructnameSession) Withdraw() (*types.Transaction, error) {
	return _Structname.Contract.Withdraw(&_Structname.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Structname *StructnameTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Structname.Contract.Withdraw(&_Structname.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Structname *StructnameTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structname.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Structname *StructnameSession) Receive() (*types.Transaction, error) {
	return _Structname.Contract.Receive(&_Structname.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Structname *StructnameTransactorSession) Receive() (*types.Transaction, error) {
	return _Structname.Contract.Receive(&_Structname.TransactOpts)
}

// StructnameErrorLogIterator is returned from FilterErrorLog and is used to iterate over the raw logs and unpacked data for ErrorLog events raised by the Structname contract.
type StructnameErrorLogIterator struct {
	Event *StructnameErrorLog // Event containing the contract specifics and raw log

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
func (it *StructnameErrorLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StructnameErrorLog)
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
		it.Event = new(StructnameErrorLog)
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
func (it *StructnameErrorLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StructnameErrorLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StructnameErrorLog represents a ErrorLog event raised by the Structname contract.
type StructnameErrorLog struct {
	Sender  common.Address
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterErrorLog is a free log retrieval operation binding the contract event 0x3811264a3dc641fd8e586a867055d4058a1592a7d5d8c41027c42b03a762c7eb.
//
// Solidity: event ErrorLog(address indexed sender, string message)
func (_Structname *StructnameFilterer) FilterErrorLog(opts *bind.FilterOpts, sender []common.Address) (*StructnameErrorLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Structname.contract.FilterLogs(opts, "ErrorLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &StructnameErrorLogIterator{contract: _Structname.contract, event: "ErrorLog", logs: logs, sub: sub}, nil
}

// WatchErrorLog is a free log subscription operation binding the contract event 0x3811264a3dc641fd8e586a867055d4058a1592a7d5d8c41027c42b03a762c7eb.
//
// Solidity: event ErrorLog(address indexed sender, string message)
func (_Structname *StructnameFilterer) WatchErrorLog(opts *bind.WatchOpts, sink chan<- *StructnameErrorLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Structname.contract.WatchLogs(opts, "ErrorLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StructnameErrorLog)
				if err := _Structname.contract.UnpackLog(event, "ErrorLog", log); err != nil {
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

// ParseErrorLog is a log parse operation binding the contract event 0x3811264a3dc641fd8e586a867055d4058a1592a7d5d8c41027c42b03a762c7eb.
//
// Solidity: event ErrorLog(address indexed sender, string message)
func (_Structname *StructnameFilterer) ParseErrorLog(log types.Log) (*StructnameErrorLog, error) {
	event := new(StructnameErrorLog)
	if err := _Structname.contract.UnpackLog(event, "ErrorLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StructnameLoserLogIterator is returned from FilterLoserLog and is used to iterate over the raw logs and unpacked data for LoserLog events raised by the Structname contract.
type StructnameLoserLogIterator struct {
	Event *StructnameLoserLog // Event containing the contract specifics and raw log

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
func (it *StructnameLoserLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StructnameLoserLog)
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
		it.Event = new(StructnameLoserLog)
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
func (it *StructnameLoserLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StructnameLoserLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StructnameLoserLog represents a LoserLog event raised by the Structname contract.
type StructnameLoserLog struct {
	Sender   common.Address
	GameType *big.Int
	BetType  *big.Int
	Position *big.Int
	Amount   *big.Int
	Multiple *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLoserLog is a free log retrieval operation binding the contract event 0x10282305406f800e3c9d5acd4862eb9942ab1f0b661f0863669c4c2ec8e04bb7.
//
// Solidity: event LoserLog(address indexed sender, uint256 gameType, uint256 betType, uint256 position, uint256 amount, uint256 multiple)
func (_Structname *StructnameFilterer) FilterLoserLog(opts *bind.FilterOpts, sender []common.Address) (*StructnameLoserLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Structname.contract.FilterLogs(opts, "LoserLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &StructnameLoserLogIterator{contract: _Structname.contract, event: "LoserLog", logs: logs, sub: sub}, nil
}

// WatchLoserLog is a free log subscription operation binding the contract event 0x10282305406f800e3c9d5acd4862eb9942ab1f0b661f0863669c4c2ec8e04bb7.
//
// Solidity: event LoserLog(address indexed sender, uint256 gameType, uint256 betType, uint256 position, uint256 amount, uint256 multiple)
func (_Structname *StructnameFilterer) WatchLoserLog(opts *bind.WatchOpts, sink chan<- *StructnameLoserLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Structname.contract.WatchLogs(opts, "LoserLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StructnameLoserLog)
				if err := _Structname.contract.UnpackLog(event, "LoserLog", log); err != nil {
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

// ParseLoserLog is a log parse operation binding the contract event 0x10282305406f800e3c9d5acd4862eb9942ab1f0b661f0863669c4c2ec8e04bb7.
//
// Solidity: event LoserLog(address indexed sender, uint256 gameType, uint256 betType, uint256 position, uint256 amount, uint256 multiple)
func (_Structname *StructnameFilterer) ParseLoserLog(log types.Log) (*StructnameLoserLog, error) {
	event := new(StructnameLoserLog)
	if err := _Structname.contract.UnpackLog(event, "LoserLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StructnameReceiveLogIterator is returned from FilterReceiveLog and is used to iterate over the raw logs and unpacked data for ReceiveLog events raised by the Structname contract.
type StructnameReceiveLogIterator struct {
	Event *StructnameReceiveLog // Event containing the contract specifics and raw log

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
func (it *StructnameReceiveLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StructnameReceiveLog)
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
		it.Event = new(StructnameReceiveLog)
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
func (it *StructnameReceiveLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StructnameReceiveLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StructnameReceiveLog represents a ReceiveLog event raised by the Structname contract.
type StructnameReceiveLog struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceiveLog is a free log retrieval operation binding the contract event 0x3001d122178ee9da2e7224593b8383a606c707c305951086f284673d62757cf5.
//
// Solidity: event ReceiveLog(address indexed sender, uint256 value)
func (_Structname *StructnameFilterer) FilterReceiveLog(opts *bind.FilterOpts, sender []common.Address) (*StructnameReceiveLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Structname.contract.FilterLogs(opts, "ReceiveLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &StructnameReceiveLogIterator{contract: _Structname.contract, event: "ReceiveLog", logs: logs, sub: sub}, nil
}

// WatchReceiveLog is a free log subscription operation binding the contract event 0x3001d122178ee9da2e7224593b8383a606c707c305951086f284673d62757cf5.
//
// Solidity: event ReceiveLog(address indexed sender, uint256 value)
func (_Structname *StructnameFilterer) WatchReceiveLog(opts *bind.WatchOpts, sink chan<- *StructnameReceiveLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Structname.contract.WatchLogs(opts, "ReceiveLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StructnameReceiveLog)
				if err := _Structname.contract.UnpackLog(event, "ReceiveLog", log); err != nil {
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

// ParseReceiveLog is a log parse operation binding the contract event 0x3001d122178ee9da2e7224593b8383a606c707c305951086f284673d62757cf5.
//
// Solidity: event ReceiveLog(address indexed sender, uint256 value)
func (_Structname *StructnameFilterer) ParseReceiveLog(log types.Log) (*StructnameReceiveLog, error) {
	event := new(StructnameReceiveLog)
	if err := _Structname.contract.UnpackLog(event, "ReceiveLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StructnameWinnerLogIterator is returned from FilterWinnerLog and is used to iterate over the raw logs and unpacked data for WinnerLog events raised by the Structname contract.
type StructnameWinnerLogIterator struct {
	Event *StructnameWinnerLog // Event containing the contract specifics and raw log

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
func (it *StructnameWinnerLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StructnameWinnerLog)
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
		it.Event = new(StructnameWinnerLog)
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
func (it *StructnameWinnerLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StructnameWinnerLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StructnameWinnerLog represents a WinnerLog event raised by the Structname contract.
type StructnameWinnerLog struct {
	Sender   common.Address
	GameType *big.Int
	BetType  *big.Int
	Position *big.Int
	Amount   *big.Int
	Multiple *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWinnerLog is a free log retrieval operation binding the contract event 0xd6cf70a2d274f01d777b7778a37498f79396f5b7d07fcc905bc4b0ff58e321fb.
//
// Solidity: event WinnerLog(address indexed sender, uint256 gameType, uint256 betType, uint256 position, uint256 amount, uint256 multiple)
func (_Structname *StructnameFilterer) FilterWinnerLog(opts *bind.FilterOpts, sender []common.Address) (*StructnameWinnerLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Structname.contract.FilterLogs(opts, "WinnerLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &StructnameWinnerLogIterator{contract: _Structname.contract, event: "WinnerLog", logs: logs, sub: sub}, nil
}

// WatchWinnerLog is a free log subscription operation binding the contract event 0xd6cf70a2d274f01d777b7778a37498f79396f5b7d07fcc905bc4b0ff58e321fb.
//
// Solidity: event WinnerLog(address indexed sender, uint256 gameType, uint256 betType, uint256 position, uint256 amount, uint256 multiple)
func (_Structname *StructnameFilterer) WatchWinnerLog(opts *bind.WatchOpts, sink chan<- *StructnameWinnerLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Structname.contract.WatchLogs(opts, "WinnerLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StructnameWinnerLog)
				if err := _Structname.contract.UnpackLog(event, "WinnerLog", log); err != nil {
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

// ParseWinnerLog is a log parse operation binding the contract event 0xd6cf70a2d274f01d777b7778a37498f79396f5b7d07fcc905bc4b0ff58e321fb.
//
// Solidity: event WinnerLog(address indexed sender, uint256 gameType, uint256 betType, uint256 position, uint256 amount, uint256 multiple)
func (_Structname *StructnameFilterer) ParseWinnerLog(log types.Log) (*StructnameWinnerLog, error) {
	event := new(StructnameWinnerLog)
	if err := _Structname.contract.UnpackLog(event, "WinnerLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
