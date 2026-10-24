// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock

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

// MockNoAARecoveryMetaData contains all meta data concerning the MockNoAARecovery contract.
var MockNoAARecoveryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"recoveryDelay_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contestLock_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelRecovery\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contestLock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeRecovery\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"identityKey\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockRecovery\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"locked\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockedUntil\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoveryDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRecovery\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newIdentityKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requests\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"newIdentityKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"readyAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockRecovery\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"RecoveryCancelled\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryContested\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"lockedUntil\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryFinalized\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newIdentityKey\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryLocked\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryRequested\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newIdentityKey\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"readyAt\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryUnlocked\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
	Bin: "0x60c060405234801561000f575f5ffd5b506040516110bd3803806110bd833981810160405281019061003191906100b1565b8167ffffffffffffffff1660808167ffffffffffffffff16815250508067ffffffffffffffff1660a08167ffffffffffffffff168152505050506100ef565b5f5ffd5b5f67ffffffffffffffff82169050919050565b61009081610074565b811461009a575f5ffd5b50565b5f815190506100ab81610087565b92915050565b5f5f604083850312156100c7576100c6610070565b5b5f6100d48582860161009d565b92505060206100e58582860161009d565b9150509250929050565b60805160a051610f9f61011e5f395f818161023601526105a601525f818161025a015261064b0152610f9f5ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c806372cc02841161006f57806372cc02841461014f578063990f9bf71461016b5780639d8669851461019b578063cbe9e764146101cc578063e99d9d48146101fc578063f7bd7e5314610218576100a7565b80631bed85ec146100ab5780633758ca7a146100c95780636286d8c8146100e7578063631fec7f146101175780636a6669b914610133575b5f5ffd5b6100b3610234565b6040516100c09190610a43565b60405180910390f35b6100d1610258565b6040516100de9190610a43565b60405180910390f35b61010160048036038101906100fc9190610a93565b61027c565b60405161010e9190610ad6565b60405180910390f35b610131600480360381019061012c9190610a93565b610291565b005b61014d60048036038101906101489190610a93565b610348565b005b61016960048036038101906101649190610b19565b6103fd565b005b61018560048036038101906101809190610a93565b610720565b6040516101929190610a43565b60405180910390f35b6101b560048036038101906101b09190610a93565b610744565b6040516101c3929190610b57565b60405180910390f35b6101e660048036038101906101e19190610a93565b610776565b6040516101f39190610b98565b60405180910390f35b61021660048036038101906102119190610a93565b610793565b005b610232600480360381019061022d9190610a93565b61096a565b005b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b6003602052805f5260405f205f915090505481565b5f5f5f8381526020019081526020015f205f0154036102e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102dc90610c0b565b60405180910390fd5b5f5f8281526020019081526020015f205f5f82015f9055600182015f6101000a81549067ffffffffffffffff02191690555050807f7d6dceb0ad9e1c7d903468b0d49a3151dcb4d9f5fb0c5d0940b3ee2a7a8fd5d160405160405180910390a250565b60025f8281526020019081526020015f205f9054906101000a900460ff166103a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039c90610c73565b60405180910390fd5b5f60025f8381526020019081526020015f205f6101000a81548160ff021916908315150217905550807fcf4323666218fdee66018eb7566c179ce553334f33de42e2ee711314977d87f960405160405180910390a250565b5f810361043f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043690610cdb565b60405180910390fd5b60025f8381526020019081526020015f205f9054906101000a900460ff161561049d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049490610d43565b60405180910390fd5b4260015f8481526020019081526020015f205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16111561050e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050590610dab565b60405180910390fd5b5f5f5f8481526020019081526020015f2090505f815f0154146106485781815f015403610570576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056790610e13565b60405180910390fd5b5f5f8481526020019081526020015f205f5f82015f9055600182015f6101000a81549067ffffffffffffffff021916905550505f7f0000000000000000000000000000000000000000000000000000000000000000426105d09190610e5e565b90508060015f8681526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550837f88aba7e1192eea8406b20d949e0e266837957afac0d62ff1693e0a7d5df7c835826040516106399190610a43565b60405180910390a2505061071c565b5f7f0000000000000000000000000000000000000000000000000000000000000000426106759190610e5e565b905060405180604001604052808481526020018267ffffffffffffffff168152505f5f8681526020019081526020015f205f820151815f01556020820151816001015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550905050837fffb8f640e7ca536ef5d2baf5a2ea5ed9d5ece60da943b95f984ee2e8058d8ddf8483604051610711929190610b57565b60405180910390a250505b5050565b6001602052805f5260405f205f915054906101000a900467ffffffffffffffff1681565b5f602052805f5260405f205f91509050805f015490806001015f9054906101000a900467ffffffffffffffff16905082565b6002602052805f5260405f205f915054906101000a900460ff1681565b5f5f5f8381526020019081526020015f206040518060400160405290815f8201548152602001600182015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505090505f815f015103610830576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082790610c0b565b60405180910390fd5b60025f8381526020019081526020015f205f9054906101000a900460ff161561088e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088590610d43565b60405180910390fd5b806020015167ffffffffffffffff164210156108df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d690610ee3565b60405180910390fd5b805f015160035f8481526020019081526020015f20819055505f5f8381526020019081526020015f205f5f82015f9055600182015f6101000a81549067ffffffffffffffff02191690555050817f0c64074ccd537c0bb34b4f79afb1c468f9e33d66f8601162fb063830fdfddb3b825f015160405161095e9190610ad6565b60405180910390a25050565b60025f8281526020019081526020015f205f9054906101000a900460ff16156109c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109bf90610f4b565b60405180910390fd5b600160025f8381526020019081526020015f205f6101000a81548160ff021916908315150217905550807ff1d7ea8fd15a531788ee7d98d5abf152e714bc1b1747fa7ec6dc5f44e3c29a8d60405160405180910390a250565b5f67ffffffffffffffff82169050919050565b610a3d81610a21565b82525050565b5f602082019050610a565f830184610a34565b92915050565b5f5ffd5b5f819050919050565b610a7281610a60565b8114610a7c575f5ffd5b50565b5f81359050610a8d81610a69565b92915050565b5f60208284031215610aa857610aa7610a5c565b5b5f610ab584828501610a7f565b91505092915050565b5f819050919050565b610ad081610abe565b82525050565b5f602082019050610ae95f830184610ac7565b92915050565b610af881610abe565b8114610b02575f5ffd5b50565b5f81359050610b1381610aef565b92915050565b5f5f60408385031215610b2f57610b2e610a5c565b5b5f610b3c85828601610a7f565b9250506020610b4d85828601610b05565b9150509250929050565b5f604082019050610b6a5f830185610ac7565b610b776020830184610a34565b9392505050565b5f8115159050919050565b610b9281610b7e565b82525050565b5f602082019050610bab5f830184610b89565b92915050565b5f82825260208201905092915050565b7f6e6f2072657175657374000000000000000000000000000000000000000000005f82015250565b5f610bf5600a83610bb1565b9150610c0082610bc1565b602082019050919050565b5f6020820190508181035f830152610c2281610be9565b9050919050565b7f6e6f74206c6f636b6564000000000000000000000000000000000000000000005f82015250565b5f610c5d600a83610bb1565b9150610c6882610c29565b602082019050919050565b5f6020820190508181035f830152610c8a81610c51565b9050919050565b7f7a65726f206b65790000000000000000000000000000000000000000000000005f82015250565b5f610cc5600883610bb1565b9150610cd082610c91565b602082019050919050565b5f6020820190508181035f830152610cf281610cb9565b9050919050565b7f6c6f636b656400000000000000000000000000000000000000000000000000005f82015250565b5f610d2d600683610bb1565b9150610d3882610cf9565b602082019050919050565b5f6020820190508181035f830152610d5a81610d21565b9050919050565b7f636f6e74657374206c6f636b00000000000000000000000000000000000000005f82015250565b5f610d95600c83610bb1565b9150610da082610d61565b602082019050919050565b5f6020820190508181035f830152610dc281610d89565b9050919050565b7f616c7265616479207265717565737465640000000000000000000000000000005f82015250565b5f610dfd601183610bb1565b9150610e0882610dc9565b602082019050919050565b5f6020820190508181035f830152610e2a81610df1565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610e6882610a21565b9150610e7383610a21565b9250828201905067ffffffffffffffff811115610e9357610e92610e31565b5b92915050565b7f6e6f7420726561647900000000000000000000000000000000000000000000005f82015250565b5f610ecd600983610bb1565b9150610ed882610e99565b602082019050919050565b5f6020820190508181035f830152610efa81610ec1565b9050919050565b7f616c7265616479206c6f636b65640000000000000000000000000000000000005f82015250565b5f610f35600e83610bb1565b9150610f4082610f01565b602082019050919050565b5f6020820190508181035f830152610f6281610f29565b905091905056fea2646970667358221220ee96bee1a29ff10c52281eab335406d475504c9e4b55bcf0d581402c44f9cbf464736f6c634300081c0033",
}

// MockNoAARecoveryABI is the input ABI used to generate the binding from.
// Deprecated: Use MockNoAARecoveryMetaData.ABI instead.
var MockNoAARecoveryABI = MockNoAARecoveryMetaData.ABI

// MockNoAARecoveryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockNoAARecoveryMetaData.Bin instead.
var MockNoAARecoveryBin = MockNoAARecoveryMetaData.Bin

// DeployMockNoAARecovery deploys a new Ethereum contract, binding an instance of MockNoAARecovery to it.
func DeployMockNoAARecovery(auth *bind.TransactOpts, backend bind.ContractBackend, recoveryDelay_ uint64, contestLock_ uint64) (common.Address, *types.Transaction, *MockNoAARecovery, error) {
	parsed, err := MockNoAARecoveryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockNoAARecoveryBin), backend, recoveryDelay_, contestLock_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockNoAARecovery{MockNoAARecoveryCaller: MockNoAARecoveryCaller{contract: contract}, MockNoAARecoveryTransactor: MockNoAARecoveryTransactor{contract: contract}, MockNoAARecoveryFilterer: MockNoAARecoveryFilterer{contract: contract}}, nil
}

// MockNoAARecovery is an auto generated Go binding around an Ethereum contract.
type MockNoAARecovery struct {
	MockNoAARecoveryCaller     // Read-only binding to the contract
	MockNoAARecoveryTransactor // Write-only binding to the contract
	MockNoAARecoveryFilterer   // Log filterer for contract events
}

// MockNoAARecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockNoAARecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNoAARecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockNoAARecoveryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNoAARecoveryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockNoAARecoveryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNoAARecoverySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockNoAARecoverySession struct {
	Contract     *MockNoAARecovery // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockNoAARecoveryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockNoAARecoveryCallerSession struct {
	Contract *MockNoAARecoveryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MockNoAARecoveryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockNoAARecoveryTransactorSession struct {
	Contract     *MockNoAARecoveryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MockNoAARecoveryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockNoAARecoveryRaw struct {
	Contract *MockNoAARecovery // Generic contract binding to access the raw methods on
}

// MockNoAARecoveryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockNoAARecoveryCallerRaw struct {
	Contract *MockNoAARecoveryCaller // Generic read-only contract binding to access the raw methods on
}

// MockNoAARecoveryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockNoAARecoveryTransactorRaw struct {
	Contract *MockNoAARecoveryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockNoAARecovery creates a new instance of MockNoAARecovery, bound to a specific deployed contract.
func NewMockNoAARecovery(address common.Address, backend bind.ContractBackend) (*MockNoAARecovery, error) {
	contract, err := bindMockNoAARecovery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockNoAARecovery{MockNoAARecoveryCaller: MockNoAARecoveryCaller{contract: contract}, MockNoAARecoveryTransactor: MockNoAARecoveryTransactor{contract: contract}, MockNoAARecoveryFilterer: MockNoAARecoveryFilterer{contract: contract}}, nil
}

// NewMockNoAARecoveryCaller creates a new read-only instance of MockNoAARecovery, bound to a specific deployed contract.
func NewMockNoAARecoveryCaller(address common.Address, caller bind.ContractCaller) (*MockNoAARecoveryCaller, error) {
	contract, err := bindMockNoAARecovery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryCaller{contract: contract}, nil
}

// NewMockNoAARecoveryTransactor creates a new write-only instance of MockNoAARecovery, bound to a specific deployed contract.
func NewMockNoAARecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*MockNoAARecoveryTransactor, error) {
	contract, err := bindMockNoAARecovery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryTransactor{contract: contract}, nil
}

// NewMockNoAARecoveryFilterer creates a new log filterer instance of MockNoAARecovery, bound to a specific deployed contract.
func NewMockNoAARecoveryFilterer(address common.Address, filterer bind.ContractFilterer) (*MockNoAARecoveryFilterer, error) {
	contract, err := bindMockNoAARecovery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryFilterer{contract: contract}, nil
}

// bindMockNoAARecovery binds a generic wrapper to an already deployed contract.
func bindMockNoAARecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockNoAARecoveryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockNoAARecovery *MockNoAARecoveryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockNoAARecovery.Contract.MockNoAARecoveryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockNoAARecovery *MockNoAARecoveryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.MockNoAARecoveryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockNoAARecovery *MockNoAARecoveryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.MockNoAARecoveryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockNoAARecovery *MockNoAARecoveryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockNoAARecovery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockNoAARecovery *MockNoAARecoveryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockNoAARecovery *MockNoAARecoveryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.contract.Transact(opts, method, params...)
}

// ContestLock is a free data retrieval call binding the contract method 0x1bed85ec.
//
// Solidity: function contestLock() view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoveryCaller) ContestLock(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockNoAARecovery.contract.Call(opts, &out, "contestLock")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ContestLock is a free data retrieval call binding the contract method 0x1bed85ec.
//
// Solidity: function contestLock() view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoverySession) ContestLock() (uint64, error) {
	return _MockNoAARecovery.Contract.ContestLock(&_MockNoAARecovery.CallOpts)
}

// ContestLock is a free data retrieval call binding the contract method 0x1bed85ec.
//
// Solidity: function contestLock() view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoveryCallerSession) ContestLock() (uint64, error) {
	return _MockNoAARecovery.Contract.ContestLock(&_MockNoAARecovery.CallOpts)
}

// IdentityKey is a free data retrieval call binding the contract method 0x6286d8c8.
//
// Solidity: function identityKey(bytes32 ) view returns(uint256)
func (_MockNoAARecovery *MockNoAARecoveryCaller) IdentityKey(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MockNoAARecovery.contract.Call(opts, &out, "identityKey", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdentityKey is a free data retrieval call binding the contract method 0x6286d8c8.
//
// Solidity: function identityKey(bytes32 ) view returns(uint256)
func (_MockNoAARecovery *MockNoAARecoverySession) IdentityKey(arg0 [32]byte) (*big.Int, error) {
	return _MockNoAARecovery.Contract.IdentityKey(&_MockNoAARecovery.CallOpts, arg0)
}

// IdentityKey is a free data retrieval call binding the contract method 0x6286d8c8.
//
// Solidity: function identityKey(bytes32 ) view returns(uint256)
func (_MockNoAARecovery *MockNoAARecoveryCallerSession) IdentityKey(arg0 [32]byte) (*big.Int, error) {
	return _MockNoAARecovery.Contract.IdentityKey(&_MockNoAARecovery.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbe9e764.
//
// Solidity: function locked(bytes32 ) view returns(bool)
func (_MockNoAARecovery *MockNoAARecoveryCaller) Locked(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MockNoAARecovery.contract.Call(opts, &out, "locked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcbe9e764.
//
// Solidity: function locked(bytes32 ) view returns(bool)
func (_MockNoAARecovery *MockNoAARecoverySession) Locked(arg0 [32]byte) (bool, error) {
	return _MockNoAARecovery.Contract.Locked(&_MockNoAARecovery.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbe9e764.
//
// Solidity: function locked(bytes32 ) view returns(bool)
func (_MockNoAARecovery *MockNoAARecoveryCallerSession) Locked(arg0 [32]byte) (bool, error) {
	return _MockNoAARecovery.Contract.Locked(&_MockNoAARecovery.CallOpts, arg0)
}

// LockedUntil is a free data retrieval call binding the contract method 0x990f9bf7.
//
// Solidity: function lockedUntil(bytes32 ) view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoveryCaller) LockedUntil(opts *bind.CallOpts, arg0 [32]byte) (uint64, error) {
	var out []interface{}
	err := _MockNoAARecovery.contract.Call(opts, &out, "lockedUntil", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LockedUntil is a free data retrieval call binding the contract method 0x990f9bf7.
//
// Solidity: function lockedUntil(bytes32 ) view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoverySession) LockedUntil(arg0 [32]byte) (uint64, error) {
	return _MockNoAARecovery.Contract.LockedUntil(&_MockNoAARecovery.CallOpts, arg0)
}

// LockedUntil is a free data retrieval call binding the contract method 0x990f9bf7.
//
// Solidity: function lockedUntil(bytes32 ) view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoveryCallerSession) LockedUntil(arg0 [32]byte) (uint64, error) {
	return _MockNoAARecovery.Contract.LockedUntil(&_MockNoAARecovery.CallOpts, arg0)
}

// RecoveryDelay is a free data retrieval call binding the contract method 0x3758ca7a.
//
// Solidity: function recoveryDelay() view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoveryCaller) RecoveryDelay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockNoAARecovery.contract.Call(opts, &out, "recoveryDelay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RecoveryDelay is a free data retrieval call binding the contract method 0x3758ca7a.
//
// Solidity: function recoveryDelay() view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoverySession) RecoveryDelay() (uint64, error) {
	return _MockNoAARecovery.Contract.RecoveryDelay(&_MockNoAARecovery.CallOpts)
}

// RecoveryDelay is a free data retrieval call binding the contract method 0x3758ca7a.
//
// Solidity: function recoveryDelay() view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoveryCallerSession) RecoveryDelay() (uint64, error) {
	return _MockNoAARecovery.Contract.RecoveryDelay(&_MockNoAARecovery.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256 newIdentityKey, uint64 readyAt)
func (_MockNoAARecovery *MockNoAARecoveryCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	NewIdentityKey *big.Int
	ReadyAt        uint64
}, error) {
	var out []interface{}
	err := _MockNoAARecovery.contract.Call(opts, &out, "requests", arg0)

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
// Solidity: function requests(bytes32 ) view returns(uint256 newIdentityKey, uint64 readyAt)
func (_MockNoAARecovery *MockNoAARecoverySession) Requests(arg0 [32]byte) (struct {
	NewIdentityKey *big.Int
	ReadyAt        uint64
}, error) {
	return _MockNoAARecovery.Contract.Requests(&_MockNoAARecovery.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256 newIdentityKey, uint64 readyAt)
func (_MockNoAARecovery *MockNoAARecoveryCallerSession) Requests(arg0 [32]byte) (struct {
	NewIdentityKey *big.Int
	ReadyAt        uint64
}, error) {
	return _MockNoAARecovery.Contract.Requests(&_MockNoAARecovery.CallOpts, arg0)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x631fec7f.
//
// Solidity: function cancelRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactor) CancelRecovery(opts *bind.TransactOpts, passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.contract.Transact(opts, "cancelRecovery", passportHash)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x631fec7f.
//
// Solidity: function cancelRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoverySession) CancelRecovery(passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.CancelRecovery(&_MockNoAARecovery.TransactOpts, passportHash)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x631fec7f.
//
// Solidity: function cancelRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactorSession) CancelRecovery(passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.CancelRecovery(&_MockNoAARecovery.TransactOpts, passportHash)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactor) FinalizeRecovery(opts *bind.TransactOpts, passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.contract.Transact(opts, "finalizeRecovery", passportHash)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoverySession) FinalizeRecovery(passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.FinalizeRecovery(&_MockNoAARecovery.TransactOpts, passportHash)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactorSession) FinalizeRecovery(passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.FinalizeRecovery(&_MockNoAARecovery.TransactOpts, passportHash)
}

// LockRecovery is a paid mutator transaction binding the contract method 0xf7bd7e53.
//
// Solidity: function lockRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactor) LockRecovery(opts *bind.TransactOpts, passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.contract.Transact(opts, "lockRecovery", passportHash)
}

// LockRecovery is a paid mutator transaction binding the contract method 0xf7bd7e53.
//
// Solidity: function lockRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoverySession) LockRecovery(passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.LockRecovery(&_MockNoAARecovery.TransactOpts, passportHash)
}

// LockRecovery is a paid mutator transaction binding the contract method 0xf7bd7e53.
//
// Solidity: function lockRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactorSession) LockRecovery(passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.LockRecovery(&_MockNoAARecovery.TransactOpts, passportHash)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0x72cc0284.
//
// Solidity: function requestRecovery(bytes32 passportHash, uint256 newIdentityKey) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactor) RequestRecovery(opts *bind.TransactOpts, passportHash [32]byte, newIdentityKey *big.Int) (*types.Transaction, error) {
	return _MockNoAARecovery.contract.Transact(opts, "requestRecovery", passportHash, newIdentityKey)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0x72cc0284.
//
// Solidity: function requestRecovery(bytes32 passportHash, uint256 newIdentityKey) returns()
func (_MockNoAARecovery *MockNoAARecoverySession) RequestRecovery(passportHash [32]byte, newIdentityKey *big.Int) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.RequestRecovery(&_MockNoAARecovery.TransactOpts, passportHash, newIdentityKey)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0x72cc0284.
//
// Solidity: function requestRecovery(bytes32 passportHash, uint256 newIdentityKey) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactorSession) RequestRecovery(passportHash [32]byte, newIdentityKey *big.Int) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.RequestRecovery(&_MockNoAARecovery.TransactOpts, passportHash, newIdentityKey)
}

// UnlockRecovery is a paid mutator transaction binding the contract method 0x6a6669b9.
//
// Solidity: function unlockRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactor) UnlockRecovery(opts *bind.TransactOpts, passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.contract.Transact(opts, "unlockRecovery", passportHash)
}

// UnlockRecovery is a paid mutator transaction binding the contract method 0x6a6669b9.
//
// Solidity: function unlockRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoverySession) UnlockRecovery(passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.UnlockRecovery(&_MockNoAARecovery.TransactOpts, passportHash)
}

// UnlockRecovery is a paid mutator transaction binding the contract method 0x6a6669b9.
//
// Solidity: function unlockRecovery(bytes32 passportHash) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactorSession) UnlockRecovery(passportHash [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.UnlockRecovery(&_MockNoAARecovery.TransactOpts, passportHash)
}

// MockNoAARecoveryRecoveryCancelledIterator is returned from FilterRecoveryCancelled and is used to iterate over the raw logs and unpacked data for RecoveryCancelled events raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryCancelledIterator struct {
	Event *MockNoAARecoveryRecoveryCancelled // Event containing the contract specifics and raw log

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
func (it *MockNoAARecoveryRecoveryCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockNoAARecoveryRecoveryCancelled)
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
		it.Event = new(MockNoAARecoveryRecoveryCancelled)
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
func (it *MockNoAARecoveryRecoveryCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockNoAARecoveryRecoveryCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockNoAARecoveryRecoveryCancelled represents a RecoveryCancelled event raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryCancelled struct {
	PassportHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecoveryCancelled is a free log retrieval operation binding the contract event 0x7d6dceb0ad9e1c7d903468b0d49a3151dcb4d9f5fb0c5d0940b3ee2a7a8fd5d1.
//
// Solidity: event RecoveryCancelled(bytes32 indexed passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryCancelled(opts *bind.FilterOpts, passportHash [][32]byte) (*MockNoAARecoveryRecoveryCancelledIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryCancelled", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryCancelledIterator{contract: _MockNoAARecovery.contract, event: "RecoveryCancelled", logs: logs, sub: sub}, nil
}

// WatchRecoveryCancelled is a free log subscription operation binding the contract event 0x7d6dceb0ad9e1c7d903468b0d49a3151dcb4d9f5fb0c5d0940b3ee2a7a8fd5d1.
//
// Solidity: event RecoveryCancelled(bytes32 indexed passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryCancelled(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryCancelled, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryCancelled", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockNoAARecoveryRecoveryCancelled)
				if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryCancelled", log); err != nil {
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
func (_MockNoAARecovery *MockNoAARecoveryFilterer) ParseRecoveryCancelled(log types.Log) (*MockNoAARecoveryRecoveryCancelled, error) {
	event := new(MockNoAARecoveryRecoveryCancelled)
	if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockNoAARecoveryRecoveryContestedIterator is returned from FilterRecoveryContested and is used to iterate over the raw logs and unpacked data for RecoveryContested events raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryContestedIterator struct {
	Event *MockNoAARecoveryRecoveryContested // Event containing the contract specifics and raw log

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
func (it *MockNoAARecoveryRecoveryContestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockNoAARecoveryRecoveryContested)
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
		it.Event = new(MockNoAARecoveryRecoveryContested)
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
func (it *MockNoAARecoveryRecoveryContestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockNoAARecoveryRecoveryContestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockNoAARecoveryRecoveryContested represents a RecoveryContested event raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryContested struct {
	PassportHash [32]byte
	LockedUntil  uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecoveryContested is a free log retrieval operation binding the contract event 0x88aba7e1192eea8406b20d949e0e266837957afac0d62ff1693e0a7d5df7c835.
//
// Solidity: event RecoveryContested(bytes32 indexed passportHash, uint64 lockedUntil)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryContested(opts *bind.FilterOpts, passportHash [][32]byte) (*MockNoAARecoveryRecoveryContestedIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryContested", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryContestedIterator{contract: _MockNoAARecovery.contract, event: "RecoveryContested", logs: logs, sub: sub}, nil
}

// WatchRecoveryContested is a free log subscription operation binding the contract event 0x88aba7e1192eea8406b20d949e0e266837957afac0d62ff1693e0a7d5df7c835.
//
// Solidity: event RecoveryContested(bytes32 indexed passportHash, uint64 lockedUntil)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryContested(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryContested, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryContested", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockNoAARecoveryRecoveryContested)
				if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryContested", log); err != nil {
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
func (_MockNoAARecovery *MockNoAARecoveryFilterer) ParseRecoveryContested(log types.Log) (*MockNoAARecoveryRecoveryContested, error) {
	event := new(MockNoAARecoveryRecoveryContested)
	if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryContested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockNoAARecoveryRecoveryFinalizedIterator is returned from FilterRecoveryFinalized and is used to iterate over the raw logs and unpacked data for RecoveryFinalized events raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryFinalizedIterator struct {
	Event *MockNoAARecoveryRecoveryFinalized // Event containing the contract specifics and raw log

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
func (it *MockNoAARecoveryRecoveryFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockNoAARecoveryRecoveryFinalized)
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
		it.Event = new(MockNoAARecoveryRecoveryFinalized)
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
func (it *MockNoAARecoveryRecoveryFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockNoAARecoveryRecoveryFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockNoAARecoveryRecoveryFinalized represents a RecoveryFinalized event raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryFinalized struct {
	PassportHash   [32]byte
	NewIdentityKey *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRecoveryFinalized is a free log retrieval operation binding the contract event 0x0c64074ccd537c0bb34b4f79afb1c468f9e33d66f8601162fb063830fdfddb3b.
//
// Solidity: event RecoveryFinalized(bytes32 indexed passportHash, uint256 newIdentityKey)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryFinalized(opts *bind.FilterOpts, passportHash [][32]byte) (*MockNoAARecoveryRecoveryFinalizedIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryFinalized", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryFinalizedIterator{contract: _MockNoAARecovery.contract, event: "RecoveryFinalized", logs: logs, sub: sub}, nil
}

// WatchRecoveryFinalized is a free log subscription operation binding the contract event 0x0c64074ccd537c0bb34b4f79afb1c468f9e33d66f8601162fb063830fdfddb3b.
//
// Solidity: event RecoveryFinalized(bytes32 indexed passportHash, uint256 newIdentityKey)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryFinalized(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryFinalized, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryFinalized", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockNoAARecoveryRecoveryFinalized)
				if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryFinalized", log); err != nil {
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
func (_MockNoAARecovery *MockNoAARecoveryFilterer) ParseRecoveryFinalized(log types.Log) (*MockNoAARecoveryRecoveryFinalized, error) {
	event := new(MockNoAARecoveryRecoveryFinalized)
	if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockNoAARecoveryRecoveryLockedIterator is returned from FilterRecoveryLocked and is used to iterate over the raw logs and unpacked data for RecoveryLocked events raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryLockedIterator struct {
	Event *MockNoAARecoveryRecoveryLocked // Event containing the contract specifics and raw log

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
func (it *MockNoAARecoveryRecoveryLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockNoAARecoveryRecoveryLocked)
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
		it.Event = new(MockNoAARecoveryRecoveryLocked)
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
func (it *MockNoAARecoveryRecoveryLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockNoAARecoveryRecoveryLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockNoAARecoveryRecoveryLocked represents a RecoveryLocked event raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryLocked struct {
	PassportHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecoveryLocked is a free log retrieval operation binding the contract event 0xf1d7ea8fd15a531788ee7d98d5abf152e714bc1b1747fa7ec6dc5f44e3c29a8d.
//
// Solidity: event RecoveryLocked(bytes32 indexed passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryLocked(opts *bind.FilterOpts, passportHash [][32]byte) (*MockNoAARecoveryRecoveryLockedIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryLocked", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryLockedIterator{contract: _MockNoAARecovery.contract, event: "RecoveryLocked", logs: logs, sub: sub}, nil
}

// WatchRecoveryLocked is a free log subscription operation binding the contract event 0xf1d7ea8fd15a531788ee7d98d5abf152e714bc1b1747fa7ec6dc5f44e3c29a8d.
//
// Solidity: event RecoveryLocked(bytes32 indexed passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryLocked(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryLocked, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryLocked", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockNoAARecoveryRecoveryLocked)
				if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryLocked", log); err != nil {
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
func (_MockNoAARecovery *MockNoAARecoveryFilterer) ParseRecoveryLocked(log types.Log) (*MockNoAARecoveryRecoveryLocked, error) {
	event := new(MockNoAARecoveryRecoveryLocked)
	if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockNoAARecoveryRecoveryRequestedIterator is returned from FilterRecoveryRequested and is used to iterate over the raw logs and unpacked data for RecoveryRequested events raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryRequestedIterator struct {
	Event *MockNoAARecoveryRecoveryRequested // Event containing the contract specifics and raw log

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
func (it *MockNoAARecoveryRecoveryRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockNoAARecoveryRecoveryRequested)
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
		it.Event = new(MockNoAARecoveryRecoveryRequested)
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
func (it *MockNoAARecoveryRecoveryRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockNoAARecoveryRecoveryRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockNoAARecoveryRecoveryRequested represents a RecoveryRequested event raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryRequested struct {
	PassportHash   [32]byte
	NewIdentityKey *big.Int
	ReadyAt        uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRecoveryRequested is a free log retrieval operation binding the contract event 0xffb8f640e7ca536ef5d2baf5a2ea5ed9d5ece60da943b95f984ee2e8058d8ddf.
//
// Solidity: event RecoveryRequested(bytes32 indexed passportHash, uint256 newIdentityKey, uint64 readyAt)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryRequested(opts *bind.FilterOpts, passportHash [][32]byte) (*MockNoAARecoveryRecoveryRequestedIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryRequested", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryRequestedIterator{contract: _MockNoAARecovery.contract, event: "RecoveryRequested", logs: logs, sub: sub}, nil
}

// WatchRecoveryRequested is a free log subscription operation binding the contract event 0xffb8f640e7ca536ef5d2baf5a2ea5ed9d5ece60da943b95f984ee2e8058d8ddf.
//
// Solidity: event RecoveryRequested(bytes32 indexed passportHash, uint256 newIdentityKey, uint64 readyAt)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryRequested(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryRequested, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryRequested", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockNoAARecoveryRecoveryRequested)
				if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryRequested", log); err != nil {
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
func (_MockNoAARecovery *MockNoAARecoveryFilterer) ParseRecoveryRequested(log types.Log) (*MockNoAARecoveryRecoveryRequested, error) {
	event := new(MockNoAARecoveryRecoveryRequested)
	if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockNoAARecoveryRecoveryUnlockedIterator is returned from FilterRecoveryUnlocked and is used to iterate over the raw logs and unpacked data for RecoveryUnlocked events raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryUnlockedIterator struct {
	Event *MockNoAARecoveryRecoveryUnlocked // Event containing the contract specifics and raw log

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
func (it *MockNoAARecoveryRecoveryUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockNoAARecoveryRecoveryUnlocked)
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
		it.Event = new(MockNoAARecoveryRecoveryUnlocked)
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
func (it *MockNoAARecoveryRecoveryUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockNoAARecoveryRecoveryUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockNoAARecoveryRecoveryUnlocked represents a RecoveryUnlocked event raised by the MockNoAARecovery contract.
type MockNoAARecoveryRecoveryUnlocked struct {
	PassportHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecoveryUnlocked is a free log retrieval operation binding the contract event 0xcf4323666218fdee66018eb7566c179ce553334f33de42e2ee711314977d87f9.
//
// Solidity: event RecoveryUnlocked(bytes32 indexed passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryUnlocked(opts *bind.FilterOpts, passportHash [][32]byte) (*MockNoAARecoveryRecoveryUnlockedIterator, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryUnlocked", passportHashRule)
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryUnlockedIterator{contract: _MockNoAARecovery.contract, event: "RecoveryUnlocked", logs: logs, sub: sub}, nil
}

// WatchRecoveryUnlocked is a free log subscription operation binding the contract event 0xcf4323666218fdee66018eb7566c179ce553334f33de42e2ee711314977d87f9.
//
// Solidity: event RecoveryUnlocked(bytes32 indexed passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryUnlocked(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryUnlocked, passportHash [][32]byte) (event.Subscription, error) {

	var passportHashRule []interface{}
	for _, passportHashItem := range passportHash {
		passportHashRule = append(passportHashRule, passportHashItem)
	}

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryUnlocked", passportHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockNoAARecoveryRecoveryUnlocked)
				if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryUnlocked", log); err != nil {
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
func (_MockNoAARecovery *MockNoAARecoveryFilterer) ParseRecoveryUnlocked(log types.Log) (*MockNoAARecoveryRecoveryUnlocked, error) {
	event := new(MockNoAARecoveryRecoveryUnlocked)
	if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
