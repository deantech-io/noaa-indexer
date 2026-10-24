// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// NoAARecoveryMetaData contains all meta data concerning the NoAARecovery contract.
var NoAARecoveryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"finalizeRecovery\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockedUntil\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requests\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"newIdentityKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"readyAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RecoveryCancelled\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryContested\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"lockedUntil\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryFinalized\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newIdentityKey\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryLocked\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryRequested\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newIdentityKey\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"readyAt\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryUnlocked\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
}

// NoAARecoveryABI is the input ABI used to generate the binding from.
// Deprecated: Use NoAARecoveryMetaData.ABI instead.
var NoAARecoveryABI = NoAARecoveryMetaData.ABI

// NoAARecovery is an auto generated Go binding around an Ethereum contract.
type NoAARecovery struct {
	NoAARecoveryCaller     // Read-only binding to the contract
	NoAARecoveryTransactor // Write-only binding to the contract
	NoAARecoveryFilterer   // Log filterer for contract events
}

// NoAARecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoAARecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoAARecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoAARecoveryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoAARecoveryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NoAARecoveryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoAARecoverySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoAARecoverySession struct {
	Contract     *NoAARecovery     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoAARecoveryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoAARecoveryCallerSession struct {
	Contract *NoAARecoveryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NoAARecoveryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoAARecoveryTransactorSession struct {
	Contract     *NoAARecoveryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NoAARecoveryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoAARecoveryRaw struct {
	Contract *NoAARecovery // Generic contract binding to access the raw methods on
}

// NoAARecoveryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoAARecoveryCallerRaw struct {
	Contract *NoAARecoveryCaller // Generic read-only contract binding to access the raw methods on
}

// NoAARecoveryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoAARecoveryTransactorRaw struct {
	Contract *NoAARecoveryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNoAARecovery creates a new instance of NoAARecovery, bound to a specific deployed contract.
func NewNoAARecovery(address common.Address, backend bind.ContractBackend) (*NoAARecovery, error) {
	contract, err := bindNoAARecovery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoAARecovery{NoAARecoveryCaller: NoAARecoveryCaller{contract: contract}, NoAARecoveryTransactor: NoAARecoveryTransactor{contract: contract}, NoAARecoveryFilterer: NoAARecoveryFilterer{contract: contract}}, nil
}

// NewNoAARecoveryCaller creates a new read-only instance of NoAARecovery, bound to a specific deployed contract.
func NewNoAARecoveryCaller(address common.Address, caller bind.ContractCaller) (*NoAARecoveryCaller, error) {
	contract, err := bindNoAARecovery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryCaller{contract: contract}, nil
}

// NewNoAARecoveryTransactor creates a new write-only instance of NoAARecovery, bound to a specific deployed contract.
func NewNoAARecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*NoAARecoveryTransactor, error) {
	contract, err := bindNoAARecovery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryTransactor{contract: contract}, nil
}

// NewNoAARecoveryFilterer creates a new log filterer instance of NoAARecovery, bound to a specific deployed contract.
func NewNoAARecoveryFilterer(address common.Address, filterer bind.ContractFilterer) (*NoAARecoveryFilterer, error) {
	contract, err := bindNoAARecovery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryFilterer{contract: contract}, nil
}

// bindNoAARecovery binds a generic wrapper to an already deployed contract.
func bindNoAARecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NoAARecoveryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoAARecovery *NoAARecoveryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoAARecovery.Contract.NoAARecoveryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoAARecovery *NoAARecoveryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoAARecovery.Contract.NoAARecoveryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoAARecovery *NoAARecoveryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoAARecovery.Contract.NoAARecoveryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoAARecovery *NoAARecoveryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoAARecovery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoAARecovery *NoAARecoveryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoAARecovery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoAARecovery *NoAARecoveryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoAARecovery.Contract.contract.Transact(opts, method, params...)
}

// LockedUntil is a free data retrieval call binding the contract method 0x990f9bf7.
//
// Solidity: function lockedUntil(bytes32 passportHash) view returns(uint64)
func (_NoAARecovery *NoAARecoveryCaller) LockedUntil(opts *bind.CallOpts, passportHash [32]byte) (uint64, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "lockedUntil", passportHash)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LockedUntil is a free data retrieval call binding the contract method 0x990f9bf7.
//
// Solidity: function lockedUntil(bytes32 passportHash) view returns(uint64)
func (_NoAARecovery *NoAARecoverySession) LockedUntil(passportHash [32]byte) (uint64, error) {
	return _NoAARecovery.Contract.LockedUntil(&_NoAARecovery.CallOpts, passportHash)
}

// LockedUntil is a free data retrieval call binding the contract method 0x990f9bf7.
//
// Solidity: function lockedUntil(bytes32 passportHash) view returns(uint64)
func (_NoAARecovery *NoAARecoveryCallerSession) LockedUntil(passportHash [32]byte) (uint64, error) {
	return _NoAARecovery.Contract.LockedUntil(&_NoAARecovery.CallOpts, passportHash)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 passportHash) view returns(uint256 newIdentityKey, uint64 readyAt)
func (_NoAARecovery *NoAARecoveryCaller) Requests(opts *bind.CallOpts, passportHash [32]byte) (struct {
	NewIdentityKey *big.Int
	ReadyAt        uint64
}, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "requests", passportHash)

	outstruct := new(struct {
		NewIdentityKey *big.Int
		ReadyAt        uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewIdentityKey = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReadyAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 passportHash) view returns(uint256 newIdentityKey, uint64 readyAt)
func (_NoAARecovery *NoAARecoverySession) Requests(passportHash [32]byte) (struct {
	NewIdentityKey *big.Int
	ReadyAt        uint64
}, error) {
	return _NoAARecovery.Contract.Requests(&_NoAARecovery.CallOpts, passportHash)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 passportHash) view returns(uint256 newIdentityKey, uint64 readyAt)
func (_NoAARecovery *NoAARecoveryCallerSession) Requests(passportHash [32]byte) (struct {
	NewIdentityKey *big.Int
	ReadyAt        uint64
}, error) {
	return _NoAARecovery.Contract.Requests(&_NoAARecovery.CallOpts, passportHash)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash) returns()
func (_NoAARecovery *NoAARecoveryTransactor) FinalizeRecovery(opts *bind.TransactOpts, passportHash [32]byte) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "finalizeRecovery", passportHash)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash) returns()
func (_NoAARecovery *NoAARecoverySession) FinalizeRecovery(passportHash [32]byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.FinalizeRecovery(&_NoAARecovery.TransactOpts, passportHash)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash) returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) FinalizeRecovery(passportHash [32]byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.FinalizeRecovery(&_NoAARecovery.TransactOpts, passportHash)
}

// NoAARecoveryRecoveryCancelledIterator is returned from FilterRecoveryCancelled and is used to iterate over the raw logs and unpacked data for RecoveryCancelled events raised by the NoAARecovery contract.
type NoAARecoveryRecoveryCancelledIterator struct {
	Event *NoAARecoveryRecoveryCancelled // Event containing the contract specifics and raw log

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
func (it *NoAARecoveryRecoveryCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoAARecoveryRecoveryCancelled)
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
		it.Event = new(NoAARecoveryRecoveryCancelled)
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
func (it *NoAARecoveryRecoveryCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoAARecoveryRecoveryCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoAARecoveryRecoveryCancelled represents a RecoveryCancelled event raised by the NoAARecovery contract.
type NoAARecoveryRecoveryCancelled struct {
	PassportHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecoveryCancelled is a free log retrieval operation binding the contract event 0x7d6dceb0ad9e1c7d903468b0d49a3151dcb4d9f5fb0c5d0940b3ee2a7a8fd5d1.
//
// Solidity: event RecoveryCancelled(bytes32 indexed passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryCancelled(opts *bind.FilterOpts, passportHash [][32]byte) (*NoAARecoveryRecoveryCancelledIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryCancelled", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryCancelledIterator{contract: _NoAARecovery.contract, event: "RecoveryCancelled", logs: logs, sub: sub}, nil
}

// WatchRecoveryCancelled is a free log subscription operation binding the contract event 0x7d6dceb0ad9e1c7d903468b0d49a3151dcb4d9f5fb0c5d0940b3ee2a7a8fd5d1.
//
// Solidity: event RecoveryCancelled(bytes32 indexed passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryCancelled(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryCancelled, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryCancelled", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoAARecoveryRecoveryCancelled)
				if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryCancelled", log); err != nil {
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

// ParseRecoveryCancelled is a log parse operation binding the contract event 0x7d6dceb0ad9e1c7d903468b0d49a3151dcb4d9f5fb0c5d0940b3ee2a7a8fd5d1.
//
// Solidity: event RecoveryCancelled(bytes32 indexed passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) ParseRecoveryCancelled(log types.Log) (*NoAARecoveryRecoveryCancelled, error) {
	event := new(NoAARecoveryRecoveryCancelled)
	if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoAARecoveryRecoveryContestedIterator is returned from FilterRecoveryContested and is used to iterate over the raw logs and unpacked data for RecoveryContested events raised by the NoAARecovery contract.
type NoAARecoveryRecoveryContestedIterator struct {
	Event *NoAARecoveryRecoveryContested // Event containing the contract specifics and raw log

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
func (it *NoAARecoveryRecoveryContestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoAARecoveryRecoveryContested)
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
		it.Event = new(NoAARecoveryRecoveryContested)
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
func (it *NoAARecoveryRecoveryContestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoAARecoveryRecoveryContestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoAARecoveryRecoveryContested represents a RecoveryContested event raised by the NoAARecovery contract.
type NoAARecoveryRecoveryContested struct {
	PassportHash [32]byte
	LockedUntil  uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecoveryContested is a free log retrieval operation binding the contract event 0x88aba7e1192eea8406b20d949e0e266837957afac0d62ff1693e0a7d5df7c835.
//
// Solidity: event RecoveryContested(bytes32 indexed passportHash, uint64 lockedUntil)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryContested(opts *bind.FilterOpts, passportHash [][32]byte) (*NoAARecoveryRecoveryContestedIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryContested", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryContestedIterator{contract: _NoAARecovery.contract, event: "RecoveryContested", logs: logs, sub: sub}, nil
}

// WatchRecoveryContested is a free log subscription operation binding the contract event 0x88aba7e1192eea8406b20d949e0e266837957afac0d62ff1693e0a7d5df7c835.
//
// Solidity: event RecoveryContested(bytes32 indexed passportHash, uint64 lockedUntil)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryContested(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryContested, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryContested", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoAARecoveryRecoveryContested)
				if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryContested", log); err != nil {
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

// ParseRecoveryContested is a log parse operation binding the contract event 0x88aba7e1192eea8406b20d949e0e266837957afac0d62ff1693e0a7d5df7c835.
//
// Solidity: event RecoveryContested(bytes32 indexed passportHash, uint64 lockedUntil)
func (_NoAARecovery *NoAARecoveryFilterer) ParseRecoveryContested(log types.Log) (*NoAARecoveryRecoveryContested, error) {
	event := new(NoAARecoveryRecoveryContested)
	if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryContested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoAARecoveryRecoveryFinalizedIterator is returned from FilterRecoveryFinalized and is used to iterate over the raw logs and unpacked data for RecoveryFinalized events raised by the NoAARecovery contract.
type NoAARecoveryRecoveryFinalizedIterator struct {
	Event *NoAARecoveryRecoveryFinalized // Event containing the contract specifics and raw log

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
func (it *NoAARecoveryRecoveryFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoAARecoveryRecoveryFinalized)
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
		it.Event = new(NoAARecoveryRecoveryFinalized)
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
func (it *NoAARecoveryRecoveryFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoAARecoveryRecoveryFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoAARecoveryRecoveryFinalized represents a RecoveryFinalized event raised by the NoAARecovery contract.
type NoAARecoveryRecoveryFinalized struct {
	PassportHash   [32]byte
	NewIdentityKey *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRecoveryFinalized is a free log retrieval operation binding the contract event 0x0c64074ccd537c0bb34b4f79afb1c468f9e33d66f8601162fb063830fdfddb3b.
//
// Solidity: event RecoveryFinalized(bytes32 indexed passportHash, uint256 newIdentityKey)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryFinalized(opts *bind.FilterOpts, passportHash [][32]byte) (*NoAARecoveryRecoveryFinalizedIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryFinalized", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryFinalizedIterator{contract: _NoAARecovery.contract, event: "RecoveryFinalized", logs: logs, sub: sub}, nil
}

// WatchRecoveryFinalized is a free log subscription operation binding the contract event 0x0c64074ccd537c0bb34b4f79afb1c468f9e33d66f8601162fb063830fdfddb3b.
//
// Solidity: event RecoveryFinalized(bytes32 indexed passportHash, uint256 newIdentityKey)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryFinalized(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryFinalized, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryFinalized", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoAARecoveryRecoveryFinalized)
				if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryFinalized", log); err != nil {
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

// ParseRecoveryFinalized is a log parse operation binding the contract event 0x0c64074ccd537c0bb34b4f79afb1c468f9e33d66f8601162fb063830fdfddb3b.
//
// Solidity: event RecoveryFinalized(bytes32 indexed passportHash, uint256 newIdentityKey)
func (_NoAARecovery *NoAARecoveryFilterer) ParseRecoveryFinalized(log types.Log) (*NoAARecoveryRecoveryFinalized, error) {
	event := new(NoAARecoveryRecoveryFinalized)
	if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoAARecoveryRecoveryLockedIterator is returned from FilterRecoveryLocked and is used to iterate over the raw logs and unpacked data for RecoveryLocked events raised by the NoAARecovery contract.
type NoAARecoveryRecoveryLockedIterator struct {
	Event *NoAARecoveryRecoveryLocked // Event containing the contract specifics and raw log

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
func (it *NoAARecoveryRecoveryLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoAARecoveryRecoveryLocked)
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
		it.Event = new(NoAARecoveryRecoveryLocked)
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
func (it *NoAARecoveryRecoveryLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoAARecoveryRecoveryLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoAARecoveryRecoveryLocked represents a RecoveryLocked event raised by the NoAARecovery contract.
type NoAARecoveryRecoveryLocked struct {
	PassportHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecoveryLocked is a free log retrieval operation binding the contract event 0xf1d7ea8fd15a531788ee7d98d5abf152e714bc1b1747fa7ec6dc5f44e3c29a8d.
//
// Solidity: event RecoveryLocked(bytes32 indexed passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryLocked(opts *bind.FilterOpts, passportHash [][32]byte) (*NoAARecoveryRecoveryLockedIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryLocked", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryLockedIterator{contract: _NoAARecovery.contract, event: "RecoveryLocked", logs: logs, sub: sub}, nil
}

// WatchRecoveryLocked is a free log subscription operation binding the contract event 0xf1d7ea8fd15a531788ee7d98d5abf152e714bc1b1747fa7ec6dc5f44e3c29a8d.
//
// Solidity: event RecoveryLocked(bytes32 indexed passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryLocked(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryLocked, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryLocked", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoAARecoveryRecoveryLocked)
				if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryLocked", log); err != nil {
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

// ParseRecoveryLocked is a log parse operation binding the contract event 0xf1d7ea8fd15a531788ee7d98d5abf152e714bc1b1747fa7ec6dc5f44e3c29a8d.
//
// Solidity: event RecoveryLocked(bytes32 indexed passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) ParseRecoveryLocked(log types.Log) (*NoAARecoveryRecoveryLocked, error) {
	event := new(NoAARecoveryRecoveryLocked)
	if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoAARecoveryRecoveryRequestedIterator is returned from FilterRecoveryRequested and is used to iterate over the raw logs and unpacked data for RecoveryRequested events raised by the NoAARecovery contract.
type NoAARecoveryRecoveryRequestedIterator struct {
	Event *NoAARecoveryRecoveryRequested // Event containing the contract specifics and raw log

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
func (it *NoAARecoveryRecoveryRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoAARecoveryRecoveryRequested)
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
		it.Event = new(NoAARecoveryRecoveryRequested)
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
func (it *NoAARecoveryRecoveryRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoAARecoveryRecoveryRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoAARecoveryRecoveryRequested represents a RecoveryRequested event raised by the NoAARecovery contract.
type NoAARecoveryRecoveryRequested struct {
	PassportHash   [32]byte
	NewIdentityKey *big.Int
	ReadyAt        uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRecoveryRequested is a free log retrieval operation binding the contract event 0xffb8f640e7ca536ef5d2baf5a2ea5ed9d5ece60da943b95f984ee2e8058d8ddf.
//
// Solidity: event RecoveryRequested(bytes32 indexed passportHash, uint256 newIdentityKey, uint64 readyAt)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryRequested(opts *bind.FilterOpts, passportHash [][32]byte) (*NoAARecoveryRecoveryRequestedIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryRequested", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryRequestedIterator{contract: _NoAARecovery.contract, event: "RecoveryRequested", logs: logs, sub: sub}, nil
}

// WatchRecoveryRequested is a free log subscription operation binding the contract event 0xffb8f640e7ca536ef5d2baf5a2ea5ed9d5ece60da943b95f984ee2e8058d8ddf.
//
// Solidity: event RecoveryRequested(bytes32 indexed passportHash, uint256 newIdentityKey, uint64 readyAt)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryRequested(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryRequested, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryRequested", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoAARecoveryRecoveryRequested)
				if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryRequested", log); err != nil {
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

// ParseRecoveryRequested is a log parse operation binding the contract event 0xffb8f640e7ca536ef5d2baf5a2ea5ed9d5ece60da943b95f984ee2e8058d8ddf.
//
// Solidity: event RecoveryRequested(bytes32 indexed passportHash, uint256 newIdentityKey, uint64 readyAt)
func (_NoAARecovery *NoAARecoveryFilterer) ParseRecoveryRequested(log types.Log) (*NoAARecoveryRecoveryRequested, error) {
	event := new(NoAARecoveryRecoveryRequested)
	if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoAARecoveryRecoveryUnlockedIterator is returned from FilterRecoveryUnlocked and is used to iterate over the raw logs and unpacked data for RecoveryUnlocked events raised by the NoAARecovery contract.
type NoAARecoveryRecoveryUnlockedIterator struct {
	Event *NoAARecoveryRecoveryUnlocked // Event containing the contract specifics and raw log

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
func (it *NoAARecoveryRecoveryUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoAARecoveryRecoveryUnlocked)
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
		it.Event = new(NoAARecoveryRecoveryUnlocked)
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
func (it *NoAARecoveryRecoveryUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoAARecoveryRecoveryUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoAARecoveryRecoveryUnlocked represents a RecoveryUnlocked event raised by the NoAARecovery contract.
type NoAARecoveryRecoveryUnlocked struct {
	PassportHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecoveryUnlocked is a free log retrieval operation binding the contract event 0xcf4323666218fdee66018eb7566c179ce553334f33de42e2ee711314977d87f9.
//
// Solidity: event RecoveryUnlocked(bytes32 indexed passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryUnlocked(opts *bind.FilterOpts, passportHash [][32]byte) (*NoAARecoveryRecoveryUnlockedIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryUnlocked", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryUnlockedIterator{contract: _NoAARecovery.contract, event: "RecoveryUnlocked", logs: logs, sub: sub}, nil
}

// WatchRecoveryUnlocked is a free log subscription operation binding the contract event 0xcf4323666218fdee66018eb7566c179ce553334f33de42e2ee711314977d87f9.
//
// Solidity: event RecoveryUnlocked(bytes32 indexed passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryUnlocked(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryUnlocked, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryUnlocked", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoAARecoveryRecoveryUnlocked)
				if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryUnlocked", log); err != nil {
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

// ParseRecoveryUnlocked is a log parse operation binding the contract event 0xcf4323666218fdee66018eb7566c179ce553334f33de42e2ee711314977d87f9.
//
// Solidity: event RecoveryUnlocked(bytes32 indexed passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) ParseRecoveryUnlocked(log types.Log) (*NoAARecoveryRecoveryUnlocked, error) {
	event := new(NoAARecoveryRecoveryUnlocked)
	if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
