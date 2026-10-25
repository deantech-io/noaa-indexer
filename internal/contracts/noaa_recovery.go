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

// NoAARecoveryRecoveryRequest is an auto generated low-level Go binding around an user-defined struct.
type NoAARecoveryRecoveryRequest struct {
	OldIdentityKey [32]byte
	NewIdentityKey [32]byte
	DgCommit       *big.Int
	ReadyAt        uint64
}

// NoAARecoveryMetaData contains all meta data concerning the NoAARecovery contract.
var NoAARecoveryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubSignals\",\"type\":\"bytes32[]\"}],\"name\":\"InvalidNoirProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pendingIdentityKey\",\"type\":\"bytes32\"}],\"name\":\"RecoveryCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pendingIdentityKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newIdentityKey\",\"type\":\"bytes32\"}],\"name\":\"RecoveryContested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldIdentityKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newIdentityKey\",\"type\":\"bytes32\"}],\"name\":\"RecoveryFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"}],\"name\":\"RecoveryLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newIdentityKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"readyAt\",\"type\":\"uint256\"}],\"name\":\"RecoveryRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"passportHash\",\"type\":\"bytes32\"}],\"name\":\"RecoveryUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REVOKED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stateKeeper_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"registration_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keyPossessionVerifier_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"recoveryDelay_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lockDuration_\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"zkTypes_\",\"type\":\"bytes32[]\"}],\"name\":\"__NoAARecovery_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"allowedZkTypes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"passportHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints_\",\"type\":\"bytes\"}],\"name\":\"cancelRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"passportHash_\",\"type\":\"bytes32\"}],\"name\":\"finalizeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"passportHash_\",\"type\":\"bytes32\"},{\"internalType\":\"enumNoAARecovery.Action\",\"name\":\"action_\",\"type\":\"uint8\"}],\"name\":\"getKeyPossessionMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"passportHash_\",\"type\":\"bytes32\"}],\"name\":\"getRecoveryRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"oldIdentityKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newIdentityKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"readyAt\",\"type\":\"uint64\"}],\"internalType\":\"structNoAARecovery.RecoveryRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isRecoveryLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyPossessionVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"passportHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints_\",\"type\":\"bytes\"}],\"name\":\"lockRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"lockedUntil\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoveryDelay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registration\",\"outputs\":[{\"internalType\":\"contractRegistration2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"certificatesRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"identityKey_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dgCommit_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"passportHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"zkType_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints_\",\"type\":\"bytes\"}],\"name\":\"requestRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keyPossessionVerifier_\",\"type\":\"address\"}],\"name\":\"setKeyPossessionVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"recoveryDelay_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lockDuration_\",\"type\":\"uint64\"}],\"name\":\"setRecoveryParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"zkTypes_\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"allowed_\",\"type\":\"bool\"}],\"name\":\"setZkTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateKeeper\",\"outputs\":[{\"internalType\":\"contractStateKeeper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"passportHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints_\",\"type\":\"bytes\"}],\"name\":\"unlockRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// REVOKED is a free data retrieval call binding the contract method 0xe6c292e8.
//
// Solidity: function REVOKED() view returns(bytes32)
func (_NoAARecovery *NoAARecoveryCaller) REVOKED(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "REVOKED")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REVOKED is a free data retrieval call binding the contract method 0xe6c292e8.
//
// Solidity: function REVOKED() view returns(bytes32)
func (_NoAARecovery *NoAARecoverySession) REVOKED() ([32]byte, error) {
	return _NoAARecovery.Contract.REVOKED(&_NoAARecovery.CallOpts)
}

// REVOKED is a free data retrieval call binding the contract method 0xe6c292e8.
//
// Solidity: function REVOKED() view returns(bytes32)
func (_NoAARecovery *NoAARecoveryCallerSession) REVOKED() ([32]byte, error) {
	return _NoAARecovery.Contract.REVOKED(&_NoAARecovery.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NoAARecovery *NoAARecoveryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NoAARecovery *NoAARecoverySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NoAARecovery.Contract.UPGRADEINTERFACEVERSION(&_NoAARecovery.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NoAARecovery *NoAARecoveryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NoAARecovery.Contract.UPGRADEINTERFACEVERSION(&_NoAARecovery.CallOpts)
}

// USED is a free data retrieval call binding the contract method 0xc152a336.
//
// Solidity: function USED() view returns(bytes32)
func (_NoAARecovery *NoAARecoveryCaller) USED(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "USED")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// USED is a free data retrieval call binding the contract method 0xc152a336.
//
// Solidity: function USED() view returns(bytes32)
func (_NoAARecovery *NoAARecoverySession) USED() ([32]byte, error) {
	return _NoAARecovery.Contract.USED(&_NoAARecovery.CallOpts)
}

// USED is a free data retrieval call binding the contract method 0xc152a336.
//
// Solidity: function USED() view returns(bytes32)
func (_NoAARecovery *NoAARecoveryCallerSession) USED() ([32]byte, error) {
	return _NoAARecovery.Contract.USED(&_NoAARecovery.CallOpts)
}

// AllowedZkTypes is a free data retrieval call binding the contract method 0x486636c7.
//
// Solidity: function allowedZkTypes(bytes32 ) view returns(bool)
func (_NoAARecovery *NoAARecoveryCaller) AllowedZkTypes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "allowedZkTypes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedZkTypes is a free data retrieval call binding the contract method 0x486636c7.
//
// Solidity: function allowedZkTypes(bytes32 ) view returns(bool)
func (_NoAARecovery *NoAARecoverySession) AllowedZkTypes(arg0 [32]byte) (bool, error) {
	return _NoAARecovery.Contract.AllowedZkTypes(&_NoAARecovery.CallOpts, arg0)
}

// AllowedZkTypes is a free data retrieval call binding the contract method 0x486636c7.
//
// Solidity: function allowedZkTypes(bytes32 ) view returns(bool)
func (_NoAARecovery *NoAARecoveryCallerSession) AllowedZkTypes(arg0 [32]byte) (bool, error) {
	return _NoAARecovery.Contract.AllowedZkTypes(&_NoAARecovery.CallOpts, arg0)
}

// GetKeyPossessionMessage is a free data retrieval call binding the contract method 0x15766b52.
//
// Solidity: function getKeyPossessionMessage(bytes32 passportHash_, uint8 action_) view returns(bytes32)
func (_NoAARecovery *NoAARecoveryCaller) GetKeyPossessionMessage(opts *bind.CallOpts, passportHash_ [32]byte, action_ uint8) ([32]byte, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "getKeyPossessionMessage", passportHash_, action_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeyPossessionMessage is a free data retrieval call binding the contract method 0x15766b52.
//
// Solidity: function getKeyPossessionMessage(bytes32 passportHash_, uint8 action_) view returns(bytes32)
func (_NoAARecovery *NoAARecoverySession) GetKeyPossessionMessage(passportHash_ [32]byte, action_ uint8) ([32]byte, error) {
	return _NoAARecovery.Contract.GetKeyPossessionMessage(&_NoAARecovery.CallOpts, passportHash_, action_)
}

// GetKeyPossessionMessage is a free data retrieval call binding the contract method 0x15766b52.
//
// Solidity: function getKeyPossessionMessage(bytes32 passportHash_, uint8 action_) view returns(bytes32)
func (_NoAARecovery *NoAARecoveryCallerSession) GetKeyPossessionMessage(passportHash_ [32]byte, action_ uint8) ([32]byte, error) {
	return _NoAARecovery.Contract.GetKeyPossessionMessage(&_NoAARecovery.CallOpts, passportHash_, action_)
}

// GetRecoveryRequest is a free data retrieval call binding the contract method 0x5ea3ec76.
//
// Solidity: function getRecoveryRequest(bytes32 passportHash_) view returns((bytes32,bytes32,uint256,uint64))
func (_NoAARecovery *NoAARecoveryCaller) GetRecoveryRequest(opts *bind.CallOpts, passportHash_ [32]byte) (NoAARecoveryRecoveryRequest, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "getRecoveryRequest", passportHash_)

	if err != nil {
		return *new(NoAARecoveryRecoveryRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(NoAARecoveryRecoveryRequest)).(*NoAARecoveryRecoveryRequest)

	return out0, err

}

// GetRecoveryRequest is a free data retrieval call binding the contract method 0x5ea3ec76.
//
// Solidity: function getRecoveryRequest(bytes32 passportHash_) view returns((bytes32,bytes32,uint256,uint64))
func (_NoAARecovery *NoAARecoverySession) GetRecoveryRequest(passportHash_ [32]byte) (NoAARecoveryRecoveryRequest, error) {
	return _NoAARecovery.Contract.GetRecoveryRequest(&_NoAARecovery.CallOpts, passportHash_)
}

// GetRecoveryRequest is a free data retrieval call binding the contract method 0x5ea3ec76.
//
// Solidity: function getRecoveryRequest(bytes32 passportHash_) view returns((bytes32,bytes32,uint256,uint64))
func (_NoAARecovery *NoAARecoveryCallerSession) GetRecoveryRequest(passportHash_ [32]byte) (NoAARecoveryRecoveryRequest, error) {
	return _NoAARecovery.Contract.GetRecoveryRequest(&_NoAARecovery.CallOpts, passportHash_)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NoAARecovery *NoAARecoveryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NoAARecovery *NoAARecoverySession) Implementation() (common.Address, error) {
	return _NoAARecovery.Contract.Implementation(&_NoAARecovery.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NoAARecovery *NoAARecoveryCallerSession) Implementation() (common.Address, error) {
	return _NoAARecovery.Contract.Implementation(&_NoAARecovery.CallOpts)
}

// IsRecoveryLocked is a free data retrieval call binding the contract method 0x76890e32.
//
// Solidity: function isRecoveryLocked(bytes32 ) view returns(bool)
func (_NoAARecovery *NoAARecoveryCaller) IsRecoveryLocked(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "isRecoveryLocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRecoveryLocked is a free data retrieval call binding the contract method 0x76890e32.
//
// Solidity: function isRecoveryLocked(bytes32 ) view returns(bool)
func (_NoAARecovery *NoAARecoverySession) IsRecoveryLocked(arg0 [32]byte) (bool, error) {
	return _NoAARecovery.Contract.IsRecoveryLocked(&_NoAARecovery.CallOpts, arg0)
}

// IsRecoveryLocked is a free data retrieval call binding the contract method 0x76890e32.
//
// Solidity: function isRecoveryLocked(bytes32 ) view returns(bool)
func (_NoAARecovery *NoAARecoveryCallerSession) IsRecoveryLocked(arg0 [32]byte) (bool, error) {
	return _NoAARecovery.Contract.IsRecoveryLocked(&_NoAARecovery.CallOpts, arg0)
}

// KeyPossessionVerifier is a free data retrieval call binding the contract method 0xfc3c2843.
//
// Solidity: function keyPossessionVerifier() view returns(address)
func (_NoAARecovery *NoAARecoveryCaller) KeyPossessionVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "keyPossessionVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyPossessionVerifier is a free data retrieval call binding the contract method 0xfc3c2843.
//
// Solidity: function keyPossessionVerifier() view returns(address)
func (_NoAARecovery *NoAARecoverySession) KeyPossessionVerifier() (common.Address, error) {
	return _NoAARecovery.Contract.KeyPossessionVerifier(&_NoAARecovery.CallOpts)
}

// KeyPossessionVerifier is a free data retrieval call binding the contract method 0xfc3c2843.
//
// Solidity: function keyPossessionVerifier() view returns(address)
func (_NoAARecovery *NoAARecoveryCallerSession) KeyPossessionVerifier() (common.Address, error) {
	return _NoAARecovery.Contract.KeyPossessionVerifier(&_NoAARecovery.CallOpts)
}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint64)
func (_NoAARecovery *NoAARecoveryCaller) LockDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "lockDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint64)
func (_NoAARecovery *NoAARecoverySession) LockDuration() (uint64, error) {
	return _NoAARecovery.Contract.LockDuration(&_NoAARecovery.CallOpts)
}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint64)
func (_NoAARecovery *NoAARecoveryCallerSession) LockDuration() (uint64, error) {
	return _NoAARecovery.Contract.LockDuration(&_NoAARecovery.CallOpts)
}

// LockedUntil is a free data retrieval call binding the contract method 0x990f9bf7.
//
// Solidity: function lockedUntil(bytes32 ) view returns(uint64)
func (_NoAARecovery *NoAARecoveryCaller) LockedUntil(opts *bind.CallOpts, arg0 [32]byte) (uint64, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "lockedUntil", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LockedUntil is a free data retrieval call binding the contract method 0x990f9bf7.
//
// Solidity: function lockedUntil(bytes32 ) view returns(uint64)
func (_NoAARecovery *NoAARecoverySession) LockedUntil(arg0 [32]byte) (uint64, error) {
	return _NoAARecovery.Contract.LockedUntil(&_NoAARecovery.CallOpts, arg0)
}

// LockedUntil is a free data retrieval call binding the contract method 0x990f9bf7.
//
// Solidity: function lockedUntil(bytes32 ) view returns(uint64)
func (_NoAARecovery *NoAARecoveryCallerSession) LockedUntil(arg0 [32]byte) (uint64, error) {
	return _NoAARecovery.Contract.LockedUntil(&_NoAARecovery.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x9e317f12.
//
// Solidity: function nonces(bytes32 ) view returns(uint256)
func (_NoAARecovery *NoAARecoveryCaller) Nonces(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x9e317f12.
//
// Solidity: function nonces(bytes32 ) view returns(uint256)
func (_NoAARecovery *NoAARecoverySession) Nonces(arg0 [32]byte) (*big.Int, error) {
	return _NoAARecovery.Contract.Nonces(&_NoAARecovery.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x9e317f12.
//
// Solidity: function nonces(bytes32 ) view returns(uint256)
func (_NoAARecovery *NoAARecoveryCallerSession) Nonces(arg0 [32]byte) (*big.Int, error) {
	return _NoAARecovery.Contract.Nonces(&_NoAARecovery.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NoAARecovery *NoAARecoveryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NoAARecovery *NoAARecoverySession) ProxiableUUID() ([32]byte, error) {
	return _NoAARecovery.Contract.ProxiableUUID(&_NoAARecovery.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NoAARecovery *NoAARecoveryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NoAARecovery.Contract.ProxiableUUID(&_NoAARecovery.CallOpts)
}

// RecoveryDelay is a free data retrieval call binding the contract method 0x3758ca7a.
//
// Solidity: function recoveryDelay() view returns(uint64)
func (_NoAARecovery *NoAARecoveryCaller) RecoveryDelay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "recoveryDelay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RecoveryDelay is a free data retrieval call binding the contract method 0x3758ca7a.
//
// Solidity: function recoveryDelay() view returns(uint64)
func (_NoAARecovery *NoAARecoverySession) RecoveryDelay() (uint64, error) {
	return _NoAARecovery.Contract.RecoveryDelay(&_NoAARecovery.CallOpts)
}

// RecoveryDelay is a free data retrieval call binding the contract method 0x3758ca7a.
//
// Solidity: function recoveryDelay() view returns(uint64)
func (_NoAARecovery *NoAARecoveryCallerSession) RecoveryDelay() (uint64, error) {
	return _NoAARecovery.Contract.RecoveryDelay(&_NoAARecovery.CallOpts)
}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(address)
func (_NoAARecovery *NoAARecoveryCaller) Registration(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "registration")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(address)
func (_NoAARecovery *NoAARecoverySession) Registration() (common.Address, error) {
	return _NoAARecovery.Contract.Registration(&_NoAARecovery.CallOpts)
}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(address)
func (_NoAARecovery *NoAARecoveryCallerSession) Registration() (common.Address, error) {
	return _NoAARecovery.Contract.Registration(&_NoAARecovery.CallOpts)
}

// StateKeeper is a free data retrieval call binding the contract method 0xc3e6f8f9.
//
// Solidity: function stateKeeper() view returns(address)
func (_NoAARecovery *NoAARecoveryCaller) StateKeeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoAARecovery.contract.Call(opts, &out, "stateKeeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateKeeper is a free data retrieval call binding the contract method 0xc3e6f8f9.
//
// Solidity: function stateKeeper() view returns(address)
func (_NoAARecovery *NoAARecoverySession) StateKeeper() (common.Address, error) {
	return _NoAARecovery.Contract.StateKeeper(&_NoAARecovery.CallOpts)
}

// StateKeeper is a free data retrieval call binding the contract method 0xc3e6f8f9.
//
// Solidity: function stateKeeper() view returns(address)
func (_NoAARecovery *NoAARecoveryCallerSession) StateKeeper() (common.Address, error) {
	return _NoAARecovery.Contract.StateKeeper(&_NoAARecovery.CallOpts)
}

// NoAARecoveryInit is a paid mutator transaction binding the contract method 0x657ef2ef.
//
// Solidity: function __NoAARecovery_init(address stateKeeper_, address registration_, address keyPossessionVerifier_, uint64 recoveryDelay_, uint64 lockDuration_, bytes32[] zkTypes_) returns()
func (_NoAARecovery *NoAARecoveryTransactor) NoAARecoveryInit(opts *bind.TransactOpts, stateKeeper_ common.Address, registration_ common.Address, keyPossessionVerifier_ common.Address, recoveryDelay_ uint64, lockDuration_ uint64, zkTypes_ [][32]byte) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "__NoAARecovery_init", stateKeeper_, registration_, keyPossessionVerifier_, recoveryDelay_, lockDuration_, zkTypes_)
}

// NoAARecoveryInit is a paid mutator transaction binding the contract method 0x657ef2ef.
//
// Solidity: function __NoAARecovery_init(address stateKeeper_, address registration_, address keyPossessionVerifier_, uint64 recoveryDelay_, uint64 lockDuration_, bytes32[] zkTypes_) returns()
func (_NoAARecovery *NoAARecoverySession) NoAARecoveryInit(stateKeeper_ common.Address, registration_ common.Address, keyPossessionVerifier_ common.Address, recoveryDelay_ uint64, lockDuration_ uint64, zkTypes_ [][32]byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.NoAARecoveryInit(&_NoAARecovery.TransactOpts, stateKeeper_, registration_, keyPossessionVerifier_, recoveryDelay_, lockDuration_, zkTypes_)
}

// NoAARecoveryInit is a paid mutator transaction binding the contract method 0x657ef2ef.
//
// Solidity: function __NoAARecovery_init(address stateKeeper_, address registration_, address keyPossessionVerifier_, uint64 recoveryDelay_, uint64 lockDuration_, bytes32[] zkTypes_) returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) NoAARecoveryInit(stateKeeper_ common.Address, registration_ common.Address, keyPossessionVerifier_ common.Address, recoveryDelay_ uint64, lockDuration_ uint64, zkTypes_ [][32]byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.NoAARecoveryInit(&_NoAARecovery.TransactOpts, stateKeeper_, registration_, keyPossessionVerifier_, recoveryDelay_, lockDuration_, zkTypes_)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x57164456.
//
// Solidity: function cancelRecovery(bytes32 passportHash_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoveryTransactor) CancelRecovery(opts *bind.TransactOpts, passportHash_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "cancelRecovery", passportHash_, zkPoints_)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x57164456.
//
// Solidity: function cancelRecovery(bytes32 passportHash_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoverySession) CancelRecovery(passportHash_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.CancelRecovery(&_NoAARecovery.TransactOpts, passportHash_, zkPoints_)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x57164456.
//
// Solidity: function cancelRecovery(bytes32 passportHash_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) CancelRecovery(passportHash_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.CancelRecovery(&_NoAARecovery.TransactOpts, passportHash_, zkPoints_)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash_) returns()
func (_NoAARecovery *NoAARecoveryTransactor) FinalizeRecovery(opts *bind.TransactOpts, passportHash_ [32]byte) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "finalizeRecovery", passportHash_)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash_) returns()
func (_NoAARecovery *NoAARecoverySession) FinalizeRecovery(passportHash_ [32]byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.FinalizeRecovery(&_NoAARecovery.TransactOpts, passportHash_)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash_) returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) FinalizeRecovery(passportHash_ [32]byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.FinalizeRecovery(&_NoAARecovery.TransactOpts, passportHash_)
}

// LockRecovery is a paid mutator transaction binding the contract method 0x9d754de8.
//
// Solidity: function lockRecovery(bytes32 passportHash_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoveryTransactor) LockRecovery(opts *bind.TransactOpts, passportHash_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "lockRecovery", passportHash_, zkPoints_)
}

// LockRecovery is a paid mutator transaction binding the contract method 0x9d754de8.
//
// Solidity: function lockRecovery(bytes32 passportHash_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoverySession) LockRecovery(passportHash_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.LockRecovery(&_NoAARecovery.TransactOpts, passportHash_, zkPoints_)
}

// LockRecovery is a paid mutator transaction binding the contract method 0x9d754de8.
//
// Solidity: function lockRecovery(bytes32 passportHash_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) LockRecovery(passportHash_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.LockRecovery(&_NoAARecovery.TransactOpts, passportHash_, zkPoints_)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0xd3e2a3df.
//
// Solidity: function requestRecovery(bytes32 certificatesRoot_, uint256 identityKey_, uint256 dgCommit_, bytes32 passportHash_, bytes32 zkType_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoveryTransactor) RequestRecovery(opts *bind.TransactOpts, certificatesRoot_ [32]byte, identityKey_ *big.Int, dgCommit_ *big.Int, passportHash_ [32]byte, zkType_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "requestRecovery", certificatesRoot_, identityKey_, dgCommit_, passportHash_, zkType_, zkPoints_)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0xd3e2a3df.
//
// Solidity: function requestRecovery(bytes32 certificatesRoot_, uint256 identityKey_, uint256 dgCommit_, bytes32 passportHash_, bytes32 zkType_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoverySession) RequestRecovery(certificatesRoot_ [32]byte, identityKey_ *big.Int, dgCommit_ *big.Int, passportHash_ [32]byte, zkType_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.RequestRecovery(&_NoAARecovery.TransactOpts, certificatesRoot_, identityKey_, dgCommit_, passportHash_, zkType_, zkPoints_)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0xd3e2a3df.
//
// Solidity: function requestRecovery(bytes32 certificatesRoot_, uint256 identityKey_, uint256 dgCommit_, bytes32 passportHash_, bytes32 zkType_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) RequestRecovery(certificatesRoot_ [32]byte, identityKey_ *big.Int, dgCommit_ *big.Int, passportHash_ [32]byte, zkType_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.RequestRecovery(&_NoAARecovery.TransactOpts, certificatesRoot_, identityKey_, dgCommit_, passportHash_, zkType_, zkPoints_)
}

// SetKeyPossessionVerifier is a paid mutator transaction binding the contract method 0x0932a10b.
//
// Solidity: function setKeyPossessionVerifier(address keyPossessionVerifier_) returns()
func (_NoAARecovery *NoAARecoveryTransactor) SetKeyPossessionVerifier(opts *bind.TransactOpts, keyPossessionVerifier_ common.Address) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "setKeyPossessionVerifier", keyPossessionVerifier_)
}

// SetKeyPossessionVerifier is a paid mutator transaction binding the contract method 0x0932a10b.
//
// Solidity: function setKeyPossessionVerifier(address keyPossessionVerifier_) returns()
func (_NoAARecovery *NoAARecoverySession) SetKeyPossessionVerifier(keyPossessionVerifier_ common.Address) (*types.Transaction, error) {
	return _NoAARecovery.Contract.SetKeyPossessionVerifier(&_NoAARecovery.TransactOpts, keyPossessionVerifier_)
}

// SetKeyPossessionVerifier is a paid mutator transaction binding the contract method 0x0932a10b.
//
// Solidity: function setKeyPossessionVerifier(address keyPossessionVerifier_) returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) SetKeyPossessionVerifier(keyPossessionVerifier_ common.Address) (*types.Transaction, error) {
	return _NoAARecovery.Contract.SetKeyPossessionVerifier(&_NoAARecovery.TransactOpts, keyPossessionVerifier_)
}

// SetRecoveryParams is a paid mutator transaction binding the contract method 0xcf67d005.
//
// Solidity: function setRecoveryParams(uint64 recoveryDelay_, uint64 lockDuration_) returns()
func (_NoAARecovery *NoAARecoveryTransactor) SetRecoveryParams(opts *bind.TransactOpts, recoveryDelay_ uint64, lockDuration_ uint64) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "setRecoveryParams", recoveryDelay_, lockDuration_)
}

// SetRecoveryParams is a paid mutator transaction binding the contract method 0xcf67d005.
//
// Solidity: function setRecoveryParams(uint64 recoveryDelay_, uint64 lockDuration_) returns()
func (_NoAARecovery *NoAARecoverySession) SetRecoveryParams(recoveryDelay_ uint64, lockDuration_ uint64) (*types.Transaction, error) {
	return _NoAARecovery.Contract.SetRecoveryParams(&_NoAARecovery.TransactOpts, recoveryDelay_, lockDuration_)
}

// SetRecoveryParams is a paid mutator transaction binding the contract method 0xcf67d005.
//
// Solidity: function setRecoveryParams(uint64 recoveryDelay_, uint64 lockDuration_) returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) SetRecoveryParams(recoveryDelay_ uint64, lockDuration_ uint64) (*types.Transaction, error) {
	return _NoAARecovery.Contract.SetRecoveryParams(&_NoAARecovery.TransactOpts, recoveryDelay_, lockDuration_)
}

// SetZkTypes is a paid mutator transaction binding the contract method 0x21926987.
//
// Solidity: function setZkTypes(bytes32[] zkTypes_, bool allowed_) returns()
func (_NoAARecovery *NoAARecoveryTransactor) SetZkTypes(opts *bind.TransactOpts, zkTypes_ [][32]byte, allowed_ bool) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "setZkTypes", zkTypes_, allowed_)
}

// SetZkTypes is a paid mutator transaction binding the contract method 0x21926987.
//
// Solidity: function setZkTypes(bytes32[] zkTypes_, bool allowed_) returns()
func (_NoAARecovery *NoAARecoverySession) SetZkTypes(zkTypes_ [][32]byte, allowed_ bool) (*types.Transaction, error) {
	return _NoAARecovery.Contract.SetZkTypes(&_NoAARecovery.TransactOpts, zkTypes_, allowed_)
}

// SetZkTypes is a paid mutator transaction binding the contract method 0x21926987.
//
// Solidity: function setZkTypes(bytes32[] zkTypes_, bool allowed_) returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) SetZkTypes(zkTypes_ [][32]byte, allowed_ bool) (*types.Transaction, error) {
	return _NoAARecovery.Contract.SetZkTypes(&_NoAARecovery.TransactOpts, zkTypes_, allowed_)
}

// UnlockRecovery is a paid mutator transaction binding the contract method 0x79204b68.
//
// Solidity: function unlockRecovery(bytes32 passportHash_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoveryTransactor) UnlockRecovery(opts *bind.TransactOpts, passportHash_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "unlockRecovery", passportHash_, zkPoints_)
}

// UnlockRecovery is a paid mutator transaction binding the contract method 0x79204b68.
//
// Solidity: function unlockRecovery(bytes32 passportHash_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoverySession) UnlockRecovery(passportHash_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.UnlockRecovery(&_NoAARecovery.TransactOpts, passportHash_, zkPoints_)
}

// UnlockRecovery is a paid mutator transaction binding the contract method 0x79204b68.
//
// Solidity: function unlockRecovery(bytes32 passportHash_, bytes zkPoints_) returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) UnlockRecovery(passportHash_ [32]byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.UnlockRecovery(&_NoAARecovery.TransactOpts, passportHash_, zkPoints_)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NoAARecovery *NoAARecoveryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NoAARecovery.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NoAARecovery *NoAARecoverySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.UpgradeToAndCall(&_NoAARecovery.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NoAARecovery *NoAARecoveryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NoAARecovery.Contract.UpgradeToAndCall(&_NoAARecovery.TransactOpts, newImplementation, data)
}

// NoAARecoveryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NoAARecovery contract.
type NoAARecoveryInitializedIterator struct {
	Event *NoAARecoveryInitialized // Event containing the contract specifics and raw log

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
func (it *NoAARecoveryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoAARecoveryInitialized)
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
		it.Event = new(NoAARecoveryInitialized)
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
func (it *NoAARecoveryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoAARecoveryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoAARecoveryInitialized represents a Initialized event raised by the NoAARecovery contract.
type NoAARecoveryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NoAARecovery *NoAARecoveryFilterer) FilterInitialized(opts *bind.FilterOpts) (*NoAARecoveryInitializedIterator, error) {

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryInitializedIterator{contract: _NoAARecovery.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NoAARecovery *NoAARecoveryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NoAARecoveryInitialized) (event.Subscription, error) {

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoAARecoveryInitialized)
				if err := _NoAARecovery.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NoAARecovery *NoAARecoveryFilterer) ParseInitialized(log types.Log) (*NoAARecoveryInitialized, error) {
	event := new(NoAARecoveryInitialized)
	if err := _NoAARecovery.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	PassportHash       [32]byte
	PendingIdentityKey [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRecoveryCancelled is a free log retrieval operation binding the contract event 0x9779648ab6a63d786fd726dd9c5f5873960dddca6252b628760bfd70e729a71d.
//
// Solidity: event RecoveryCancelled(bytes32 passportHash, bytes32 pendingIdentityKey)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryCancelled(opts *bind.FilterOpts) (*NoAARecoveryRecoveryCancelledIterator, error) {

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryCancelled")
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryCancelledIterator{contract: _NoAARecovery.contract, event: "RecoveryCancelled", logs: logs, sub: sub}, nil
}

// WatchRecoveryCancelled is a free log subscription operation binding the contract event 0x9779648ab6a63d786fd726dd9c5f5873960dddca6252b628760bfd70e729a71d.
//
// Solidity: event RecoveryCancelled(bytes32 passportHash, bytes32 pendingIdentityKey)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryCancelled(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryCancelled) (event.Subscription, error) {

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryCancelled")
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

// ParseRecoveryCancelled is a log parse operation binding the contract event 0x9779648ab6a63d786fd726dd9c5f5873960dddca6252b628760bfd70e729a71d.
//
// Solidity: event RecoveryCancelled(bytes32 passportHash, bytes32 pendingIdentityKey)
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
	PassportHash       [32]byte
	PendingIdentityKey [32]byte
	NewIdentityKey     [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRecoveryContested is a free log retrieval operation binding the contract event 0xefb80809a6c7055a258032b2ed0b2fb8409d7060f11fa0b4181ecc1d956e9f27.
//
// Solidity: event RecoveryContested(bytes32 passportHash, bytes32 pendingIdentityKey, bytes32 newIdentityKey)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryContested(opts *bind.FilterOpts) (*NoAARecoveryRecoveryContestedIterator, error) {

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryContested")
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryContestedIterator{contract: _NoAARecovery.contract, event: "RecoveryContested", logs: logs, sub: sub}, nil
}

// WatchRecoveryContested is a free log subscription operation binding the contract event 0xefb80809a6c7055a258032b2ed0b2fb8409d7060f11fa0b4181ecc1d956e9f27.
//
// Solidity: event RecoveryContested(bytes32 passportHash, bytes32 pendingIdentityKey, bytes32 newIdentityKey)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryContested(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryContested) (event.Subscription, error) {

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryContested")
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

// ParseRecoveryContested is a log parse operation binding the contract event 0xefb80809a6c7055a258032b2ed0b2fb8409d7060f11fa0b4181ecc1d956e9f27.
//
// Solidity: event RecoveryContested(bytes32 passportHash, bytes32 pendingIdentityKey, bytes32 newIdentityKey)
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
	OldIdentityKey [32]byte
	NewIdentityKey [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRecoveryFinalized is a free log retrieval operation binding the contract event 0x5e578c8f35382a852852945932e1082f3805c53821c3a9ff95efa6cd1dc52233.
//
// Solidity: event RecoveryFinalized(bytes32 passportHash, bytes32 oldIdentityKey, bytes32 newIdentityKey)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryFinalized(opts *bind.FilterOpts) (*NoAARecoveryRecoveryFinalizedIterator, error) {

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryFinalized")
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryFinalizedIterator{contract: _NoAARecovery.contract, event: "RecoveryFinalized", logs: logs, sub: sub}, nil
}

// WatchRecoveryFinalized is a free log subscription operation binding the contract event 0x5e578c8f35382a852852945932e1082f3805c53821c3a9ff95efa6cd1dc52233.
//
// Solidity: event RecoveryFinalized(bytes32 passportHash, bytes32 oldIdentityKey, bytes32 newIdentityKey)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryFinalized(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryFinalized) (event.Subscription, error) {

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryFinalized")
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

// ParseRecoveryFinalized is a log parse operation binding the contract event 0x5e578c8f35382a852852945932e1082f3805c53821c3a9ff95efa6cd1dc52233.
//
// Solidity: event RecoveryFinalized(bytes32 passportHash, bytes32 oldIdentityKey, bytes32 newIdentityKey)
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
// Solidity: event RecoveryLocked(bytes32 passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryLocked(opts *bind.FilterOpts) (*NoAARecoveryRecoveryLockedIterator, error) {

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryLocked")
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryLockedIterator{contract: _NoAARecovery.contract, event: "RecoveryLocked", logs: logs, sub: sub}, nil
}

// WatchRecoveryLocked is a free log subscription operation binding the contract event 0xf1d7ea8fd15a531788ee7d98d5abf152e714bc1b1747fa7ec6dc5f44e3c29a8d.
//
// Solidity: event RecoveryLocked(bytes32 passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryLocked(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryLocked) (event.Subscription, error) {

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryLocked")
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
// Solidity: event RecoveryLocked(bytes32 passportHash)
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
	NewIdentityKey [32]byte
	ReadyAt        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRecoveryRequested is a free log retrieval operation binding the contract event 0x37449e8a7c9f62932eb6e88550d466720d853433bf2bb22c8ff42bc492b76739.
//
// Solidity: event RecoveryRequested(bytes32 passportHash, bytes32 newIdentityKey, uint256 readyAt)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryRequested(opts *bind.FilterOpts) (*NoAARecoveryRecoveryRequestedIterator, error) {

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryRequested")
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryRequestedIterator{contract: _NoAARecovery.contract, event: "RecoveryRequested", logs: logs, sub: sub}, nil
}

// WatchRecoveryRequested is a free log subscription operation binding the contract event 0x37449e8a7c9f62932eb6e88550d466720d853433bf2bb22c8ff42bc492b76739.
//
// Solidity: event RecoveryRequested(bytes32 passportHash, bytes32 newIdentityKey, uint256 readyAt)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryRequested(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryRequested) (event.Subscription, error) {

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryRequested")
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

// ParseRecoveryRequested is a log parse operation binding the contract event 0x37449e8a7c9f62932eb6e88550d466720d853433bf2bb22c8ff42bc492b76739.
//
// Solidity: event RecoveryRequested(bytes32 passportHash, bytes32 newIdentityKey, uint256 readyAt)
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
// Solidity: event RecoveryUnlocked(bytes32 passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) FilterRecoveryUnlocked(opts *bind.FilterOpts) (*NoAARecoveryRecoveryUnlockedIterator, error) {

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "RecoveryUnlocked")
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryRecoveryUnlockedIterator{contract: _NoAARecovery.contract, event: "RecoveryUnlocked", logs: logs, sub: sub}, nil
}

// WatchRecoveryUnlocked is a free log subscription operation binding the contract event 0xcf4323666218fdee66018eb7566c179ce553334f33de42e2ee711314977d87f9.
//
// Solidity: event RecoveryUnlocked(bytes32 passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) WatchRecoveryUnlocked(opts *bind.WatchOpts, sink chan<- *NoAARecoveryRecoveryUnlocked) (event.Subscription, error) {

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "RecoveryUnlocked")
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
// Solidity: event RecoveryUnlocked(bytes32 passportHash)
func (_NoAARecovery *NoAARecoveryFilterer) ParseRecoveryUnlocked(log types.Log) (*NoAARecoveryRecoveryUnlocked, error) {
	event := new(NoAARecoveryRecoveryUnlocked)
	if err := _NoAARecovery.contract.UnpackLog(event, "RecoveryUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoAARecoveryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NoAARecovery contract.
type NoAARecoveryUpgradedIterator struct {
	Event *NoAARecoveryUpgraded // Event containing the contract specifics and raw log

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
func (it *NoAARecoveryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoAARecoveryUpgraded)
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
		it.Event = new(NoAARecoveryUpgraded)
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
func (it *NoAARecoveryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoAARecoveryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoAARecoveryUpgraded represents a Upgraded event raised by the NoAARecovery contract.
type NoAARecoveryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NoAARecovery *NoAARecoveryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NoAARecoveryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NoAARecovery.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NoAARecoveryUpgradedIterator{contract: _NoAARecovery.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NoAARecovery *NoAARecoveryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NoAARecoveryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NoAARecovery.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoAARecoveryUpgraded)
				if err := _NoAARecovery.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NoAARecovery *NoAARecoveryFilterer) ParseUpgraded(log types.Log) (*NoAARecoveryUpgraded, error) {
	event := new(NoAARecoveryUpgraded)
	if err := _NoAARecovery.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
