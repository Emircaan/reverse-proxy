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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UserAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UserRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"authorizeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600180546001600160a01b031916331790556102a9806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806367c2a360146100515780638da5cb5b14610079578063ef097fc41461009d578063fe9fbb80146100c3575b600080fd5b6100776004803603602081101561006757600080fd5b50356001600160a01b03166100fd565b005b610081610192565b604080516001600160a01b039092168252519081900360200190f35b610077600480360360208110156100b357600080fd5b50356001600160a01b03166101a1565b6100e9600480360360208110156100d957600080fd5b50356001600160a01b0316610233565b604080519115158252519081900360200190f35b6001546001600160a01b031633146101465760405162461bcd60e51b81526004018080602001828103825260228152602001806102526022913960400191505060405180910390fd5b6001600160a01b038116600081815260208190526040808220805460ff19166001179055517fb0be505cf6e26533f3066ac7722c3f8a5e8a123f43187c7832d333c49603b1469190a250565b6001546001600160a01b031681565b6001546001600160a01b031633146101ea5760405162461bcd60e51b81526004018080602001828103825260228152602001806102526022913960400191505060405180910390fd5b6001600160a01b038116600081815260208190526040808220805460ff19169055517fb43ac7bf01d16ebcb8c868a4c03fecb1f6984ea4fa1a52c116ef282b8dbb17c59190a250565b6001600160a01b031660009081526020819052604090205460ff169056fe4f6e6c79206f776e65722063616e20706572666f726d207468697320616374696f6ea264697066735822122052d88049744aab69ae9d89885035880ddf1c94dac91581b671c029ec5786fde064736f6c63430007060033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address user) view returns(bool)
func (_Contracts *ContractsCaller) IsAuthorized(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isAuthorized", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address user) view returns(bool)
func (_Contracts *ContractsSession) IsAuthorized(user common.Address) (bool, error) {
	return _Contracts.Contract.IsAuthorized(&_Contracts.CallOpts, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address user) view returns(bool)
func (_Contracts *ContractsCallerSession) IsAuthorized(user common.Address) (bool, error) {
	return _Contracts.Contract.IsAuthorized(&_Contracts.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// AuthorizeUser is a paid mutator transaction binding the contract method 0x67c2a360.
//
// Solidity: function authorizeUser(address user) returns()
func (_Contracts *ContractsTransactor) AuthorizeUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "authorizeUser", user)
}

// AuthorizeUser is a paid mutator transaction binding the contract method 0x67c2a360.
//
// Solidity: function authorizeUser(address user) returns()
func (_Contracts *ContractsSession) AuthorizeUser(user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AuthorizeUser(&_Contracts.TransactOpts, user)
}

// AuthorizeUser is a paid mutator transaction binding the contract method 0x67c2a360.
//
// Solidity: function authorizeUser(address user) returns()
func (_Contracts *ContractsTransactorSession) AuthorizeUser(user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AuthorizeUser(&_Contracts.TransactOpts, user)
}

// RevokeUser is a paid mutator transaction binding the contract method 0xef097fc4.
//
// Solidity: function revokeUser(address user) returns()
func (_Contracts *ContractsTransactor) RevokeUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revokeUser", user)
}

// RevokeUser is a paid mutator transaction binding the contract method 0xef097fc4.
//
// Solidity: function revokeUser(address user) returns()
func (_Contracts *ContractsSession) RevokeUser(user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeUser(&_Contracts.TransactOpts, user)
}

// RevokeUser is a paid mutator transaction binding the contract method 0xef097fc4.
//
// Solidity: function revokeUser(address user) returns()
func (_Contracts *ContractsTransactorSession) RevokeUser(user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeUser(&_Contracts.TransactOpts, user)
}

// ContractsUserAuthorizedIterator is returned from FilterUserAuthorized and is used to iterate over the raw logs and unpacked data for UserAuthorized events raised by the Contracts contract.
type ContractsUserAuthorizedIterator struct {
	Event *ContractsUserAuthorized // Event containing the contract specifics and raw log

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
func (it *ContractsUserAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUserAuthorized)
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
		it.Event = new(ContractsUserAuthorized)
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
func (it *ContractsUserAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUserAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUserAuthorized represents a UserAuthorized event raised by the Contracts contract.
type ContractsUserAuthorized struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserAuthorized is a free log retrieval operation binding the contract event 0xb0be505cf6e26533f3066ac7722c3f8a5e8a123f43187c7832d333c49603b146.
//
// Solidity: event UserAuthorized(address indexed user)
func (_Contracts *ContractsFilterer) FilterUserAuthorized(opts *bind.FilterOpts, user []common.Address) (*ContractsUserAuthorizedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "UserAuthorized", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractsUserAuthorizedIterator{contract: _Contracts.contract, event: "UserAuthorized", logs: logs, sub: sub}, nil
}

// WatchUserAuthorized is a free log subscription operation binding the contract event 0xb0be505cf6e26533f3066ac7722c3f8a5e8a123f43187c7832d333c49603b146.
//
// Solidity: event UserAuthorized(address indexed user)
func (_Contracts *ContractsFilterer) WatchUserAuthorized(opts *bind.WatchOpts, sink chan<- *ContractsUserAuthorized, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "UserAuthorized", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUserAuthorized)
				if err := _Contracts.contract.UnpackLog(event, "UserAuthorized", log); err != nil {
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

// ParseUserAuthorized is a log parse operation binding the contract event 0xb0be505cf6e26533f3066ac7722c3f8a5e8a123f43187c7832d333c49603b146.
//
// Solidity: event UserAuthorized(address indexed user)
func (_Contracts *ContractsFilterer) ParseUserAuthorized(log types.Log) (*ContractsUserAuthorized, error) {
	event := new(ContractsUserAuthorized)
	if err := _Contracts.contract.UnpackLog(event, "UserAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsUserRevokedIterator is returned from FilterUserRevoked and is used to iterate over the raw logs and unpacked data for UserRevoked events raised by the Contracts contract.
type ContractsUserRevokedIterator struct {
	Event *ContractsUserRevoked // Event containing the contract specifics and raw log

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
func (it *ContractsUserRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUserRevoked)
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
		it.Event = new(ContractsUserRevoked)
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
func (it *ContractsUserRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUserRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUserRevoked represents a UserRevoked event raised by the Contracts contract.
type ContractsUserRevoked struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserRevoked is a free log retrieval operation binding the contract event 0xb43ac7bf01d16ebcb8c868a4c03fecb1f6984ea4fa1a52c116ef282b8dbb17c5.
//
// Solidity: event UserRevoked(address indexed user)
func (_Contracts *ContractsFilterer) FilterUserRevoked(opts *bind.FilterOpts, user []common.Address) (*ContractsUserRevokedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "UserRevoked", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractsUserRevokedIterator{contract: _Contracts.contract, event: "UserRevoked", logs: logs, sub: sub}, nil
}

// WatchUserRevoked is a free log subscription operation binding the contract event 0xb43ac7bf01d16ebcb8c868a4c03fecb1f6984ea4fa1a52c116ef282b8dbb17c5.
//
// Solidity: event UserRevoked(address indexed user)
func (_Contracts *ContractsFilterer) WatchUserRevoked(opts *bind.WatchOpts, sink chan<- *ContractsUserRevoked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "UserRevoked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUserRevoked)
				if err := _Contracts.contract.UnpackLog(event, "UserRevoked", log); err != nil {
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

// ParseUserRevoked is a log parse operation binding the contract event 0xb43ac7bf01d16ebcb8c868a4c03fecb1f6984ea4fa1a52c116ef282b8dbb17c5.
//
// Solidity: event UserRevoked(address indexed user)
func (_Contracts *ContractsFilterer) ParseUserRevoked(log types.Log) (*ContractsUserRevoked, error) {
	event := new(ContractsUserRevoked)
	if err := _Contracts.contract.UnpackLog(event, "UserRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
