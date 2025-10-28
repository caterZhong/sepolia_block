// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter

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

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialCount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"CountIncreased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"incrementBy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161092d38038061092d83398181016040528101906100329190610112565b8060008190555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35061015c565b60008151905061010c81610145565b92915050565b60006020828403121561012457600080fd5b6000610132848285016100fd565b91505092915050565b6000819050919050565b61014e8161013b565b811461015957600080fd5b50565b6107c28061016b6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806303df179c146100675780638da5cb5b14610097578063a87d942c146100b5578063d09de08a146100d3578063d826f88f146100f1578063f2fde38b146100fb575b600080fd5b610081600480360381019061007c9190610529565b610117565b60405161008e9190610671565b60405180910390f35b61009f61018b565b6040516100ac9190610616565b60405180910390f35b6100bd6101b1565b6040516100ca9190610671565b60405180910390f35b6100db6101ba565b6040516100e89190610671565b60405180910390f35b6100f961022d565b005b61011560048036038101906101109190610500565b610316565b005b60008160008082825461012a919061069d565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167ff460188d4f685d0830992f7cc91050b3d1929947e6bd0bd69bf748efd1c82af96000546040516101799190610671565b60405180910390a26000549050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054905090565b600060016000808282546101ce919061069d565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167ff460188d4f685d0830992f7cc91050b3d1929947e6bd0bd69bf748efd1c82af960005460405161021d9190610671565b60405180910390a2600054905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b490610631565b60405180910390fd5b600080819055503373ffffffffffffffffffffffffffffffffffffffff167ff460188d4f685d0830992f7cc91050b3d1929947e6bd0bd69bf748efd1c82af960005460405161030c9190610671565b60405180910390a2565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039d90610631565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610416576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040d90610651565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000813590506104e58161075e565b92915050565b6000813590506104fa81610775565b92915050565b60006020828403121561051257600080fd5b6000610520848285016104d6565b91505092915050565b60006020828403121561053b57600080fd5b6000610549848285016104eb565b91505092915050565b61055b816106f3565b82525050565b600061056e60218361068c565b91507f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f60008301527f6e000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006105d460208361068c565b91507f4e6577206f776e65722063616e6e6f74206265207a65726f20616464726573736000830152602082019050919050565b61061081610725565b82525050565b600060208201905061062b6000830184610552565b92915050565b6000602082019050818103600083015261064a81610561565b9050919050565b6000602082019050818103600083015261066a816105c7565b9050919050565b60006020820190506106866000830184610607565b92915050565b600082825260208201905092915050565b60006106a882610725565b91506106b383610725565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156106e8576106e761072f565b5b828201905092915050565b60006106fe82610705565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b610767816106f3565b811461077257600080fd5b50565b61077e81610725565b811461078957600080fd5b5056fea264697066735822122024d0091bf363d991bb01641ebd5897e22380b4997618d749322f734412422a5464736f6c63430008000033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend, _initialCount *big.Int) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend, _initialCount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterCallerSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Counter *CounterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Counter *CounterSession) Owner() (common.Address, error) {
	return _Counter.Contract.Owner(&_Counter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Counter *CounterCallerSession) Owner() (common.Address, error) {
	return _Counter.Contract.Owner(&_Counter.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(uint256)
func (_Counter *CounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(uint256)
func (_Counter *CounterSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(uint256)
func (_Counter *CounterTransactorSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// IncrementBy is a paid mutator transaction binding the contract method 0x03df179c.
//
// Solidity: function incrementBy(uint256 _value) returns(uint256)
func (_Counter *CounterTransactor) IncrementBy(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "incrementBy", _value)
}

// IncrementBy is a paid mutator transaction binding the contract method 0x03df179c.
//
// Solidity: function incrementBy(uint256 _value) returns(uint256)
func (_Counter *CounterSession) IncrementBy(_value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.IncrementBy(&_Counter.TransactOpts, _value)
}

// IncrementBy is a paid mutator transaction binding the contract method 0x03df179c.
//
// Solidity: function incrementBy(uint256 _value) returns(uint256)
func (_Counter *CounterTransactorSession) IncrementBy(_value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.IncrementBy(&_Counter.TransactOpts, _value)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterTransactorSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Counter *CounterTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Counter *CounterSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Counter.Contract.TransferOwnership(&_Counter.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Counter *CounterTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Counter.Contract.TransferOwnership(&_Counter.TransactOpts, _newOwner)
}

// CounterCountIncreasedIterator is returned from FilterCountIncreased and is used to iterate over the raw logs and unpacked data for CountIncreased events raised by the Counter contract.
type CounterCountIncreasedIterator struct {
	Event *CounterCountIncreased // Event containing the contract specifics and raw log

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
func (it *CounterCountIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterCountIncreased)
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
		it.Event = new(CounterCountIncreased)
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
func (it *CounterCountIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterCountIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterCountIncreased represents a CountIncreased event raised by the Counter contract.
type CounterCountIncreased struct {
	NewCount *big.Int
	By       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountIncreased is a free log retrieval operation binding the contract event 0xf460188d4f685d0830992f7cc91050b3d1929947e6bd0bd69bf748efd1c82af9.
//
// Solidity: event CountIncreased(uint256 newCount, address indexed by)
func (_Counter *CounterFilterer) FilterCountIncreased(opts *bind.FilterOpts, by []common.Address) (*CounterCountIncreasedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Counter.contract.FilterLogs(opts, "CountIncreased", byRule)
	if err != nil {
		return nil, err
	}
	return &CounterCountIncreasedIterator{contract: _Counter.contract, event: "CountIncreased", logs: logs, sub: sub}, nil
}

// WatchCountIncreased is a free log subscription operation binding the contract event 0xf460188d4f685d0830992f7cc91050b3d1929947e6bd0bd69bf748efd1c82af9.
//
// Solidity: event CountIncreased(uint256 newCount, address indexed by)
func (_Counter *CounterFilterer) WatchCountIncreased(opts *bind.WatchOpts, sink chan<- *CounterCountIncreased, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Counter.contract.WatchLogs(opts, "CountIncreased", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterCountIncreased)
				if err := _Counter.contract.UnpackLog(event, "CountIncreased", log); err != nil {
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

// ParseCountIncreased is a log parse operation binding the contract event 0xf460188d4f685d0830992f7cc91050b3d1929947e6bd0bd69bf748efd1c82af9.
//
// Solidity: event CountIncreased(uint256 newCount, address indexed by)
func (_Counter *CounterFilterer) ParseCountIncreased(log types.Log) (*CounterCountIncreased, error) {
	event := new(CounterCountIncreased)
	if err := _Counter.contract.UnpackLog(event, "CountIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Counter contract.
type CounterOwnershipTransferredIterator struct {
	Event *CounterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CounterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterOwnershipTransferred)
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
		it.Event = new(CounterOwnershipTransferred)
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
func (it *CounterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterOwnershipTransferred represents a OwnershipTransferred event raised by the Counter contract.
type CounterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Counter *CounterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CounterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Counter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CounterOwnershipTransferredIterator{contract: _Counter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Counter *CounterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CounterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Counter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterOwnershipTransferred)
				if err := _Counter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Counter *CounterFilterer) ParseOwnershipTransferred(log types.Log) (*CounterOwnershipTransferred, error) {
	event := new(CounterOwnershipTransferred)
	if err := _Counter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
