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

// INoAARecoveryRecoveryRequest is an auto generated low-level Go binding around an user-defined struct.
type INoAARecoveryRecoveryRequest struct {
	OldIdentityKey [32]byte
	NewIdentityKey [32]byte
	DgCommit       *big.Int
	ReadyAt        uint64
}

// MockNoAARecoveryMetaData contains all meta data concerning the MockNoAARecovery contract.
var MockNoAARecoveryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"recoveryDelay_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lockDuration_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeIdentity\",\"inputs\":[{\"name\":\"passportHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelRecovery\",\"inputs\":[{\"name\":\"passportHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeRecovery\",\"inputs\":[{\"name\":\"passportHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRecoveryRequest\",\"inputs\":[{\"name\":\"passportHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structINoAARecovery.RecoveryRequest\",\"components\":[{\"name\":\"oldIdentityKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newIdentityKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"dgCommit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"readyAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRecoveryLocked\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockRecovery\",\"inputs\":[{\"name\":\"passportHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockedUntil\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoveryDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRecovery\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"identityKey_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dgCommit_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"passportHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlockRecovery\",\"inputs\":[{\"name\":\"passportHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"RecoveryCancelled\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"pendingIdentityKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryContested\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"pendingIdentityKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"newIdentityKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryFinalized\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"oldIdentityKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"newIdentityKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryLocked\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryRequested\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"newIdentityKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"readyAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoveryUnlocked\",\"inputs\":[{\"name\":\"passportHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
	Bin: "0x60c060405234801561000f575f5ffd5b506040516116fa3803806116fa833981810160405281019061003191906100b1565b8167ffffffffffffffff1660808167ffffffffffffffff16815250508067ffffffffffffffff1660a08167ffffffffffffffff168152505050506100ef565b5f5ffd5b5f67ffffffffffffffff82169050919050565b61009081610074565b811461009a575f5ffd5b50565b5f815190506100ab81610087565b92915050565b5f5f604083850312156100c7576100c6610070565b5b5f6100d48582860161009d565b92505060206100e58582860161009d565b9150509250929050565b60805160a0516115d56101255f395f818161023501528181610319015261097601525f8181610259015261083001526115d55ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c806379204b681161006f57806379204b6814610163578063990f9bf71461017f5780639d754de8146101af578063d3e2a3df146101cb578063e99d9d48146101e7578063fe7b8de114610203576100a7565b806304554443146100ab5780633758ca7a146100c957806357164456146100e75780635ea3ec761461010357806376890e3214610133575b5f5ffd5b6100b3610233565b6040516100c09190610c9d565b60405180910390f35b6100d1610257565b6040516100de9190610c9d565b60405180910390f35b61010160048036038101906100fc9190610d52565b61027b565b005b61011d60048036038101906101189190610daf565b6103b8565b60405161012a9190610e63565b60405180910390f35b61014d60048036038101906101489190610daf565b610430565b60405161015a9190610e96565b60405180910390f35b61017d60048036038101906101789190610d52565b61044d565b005b61019960048036038101906101949190610daf565b61050e565b6040516101a69190610c9d565b60405180910390f35b6101c960048036038101906101c49190610d52565b610532565b005b6101e560048036038101906101e09190610ed9565b610691565b005b61020160048036038101906101fc9190610daf565b610a20565b005b61021d60048036038101906102189190610daf565b610c1b565b60405161022a9190610f92565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f5f8581526020019081526020015f206001015490505f5f1b81036102d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102cd9061102b565b60405180910390fd5b5f5f8581526020019081526020015f205f5f82015f9055600182015f9055600282015f9055600382015f6101000a81549067ffffffffffffffff021916905550507f0000000000000000000000000000000000000000000000000000000000000000426103439190611076565b60015f8681526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f9779648ab6a63d786fd726dd9c5f5873960dddca6252b628760bfd70e729a71d84826040516103aa9291906110b1565b60405180910390a150505050565b6103c0610c4d565b5f5f8381526020019081526020015f206040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250509050919050565b6002602052805f5260405f205f915054906101000a900460ff1681565b60025f8481526020019081526020015f205f9054906101000a900460ff166104aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a190611148565b60405180910390fd5b5f60025f8581526020019081526020015f205f6101000a81548160ff0219169083151502179055507fcf4323666218fdee66018eb7566c179ce553334f33de42e2ee711314977d87f9836040516105019190610f92565b60405180910390a1505050565b6001602052805f5260405f205f915054906101000a900467ffffffffffffffff1681565b60025f8481526020019081526020015f205f9054906101000a900460ff1615610590576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610587906111b0565b60405180910390fd5b600160025f8581526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f5f8581526020019081526020015f206001015490505f5f1b8114610654575f5f8581526020019081526020015f205f5f82015f9055600182015f9055600282015f9055600382015f6101000a81549067ffffffffffffffff021916905550507f9779648ab6a63d786fd726dd9c5f5873960dddca6252b628760bfd70e729a71d848260405161064b9291906110b1565b60405180910390a15b7ff1d7ea8fd15a531788ee7d98d5abf152e714bc1b1747fa7ec6dc5f44e3c29a8d846040516106839190610f92565b60405180910390a150505050565b5f86116106d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ca9061123e565b60405180910390fd5b60025f8581526020019081526020015f205f9054906101000a900460ff1615610731576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610728906111b0565b60405180910390fd5b4260015f8681526020019081526020015f205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1611156107a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610799906112a6565b60405180910390fd5b5f6107ac85610c1b565b9050865f1b81036107f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e990611334565b60405180910390fd5b5f5f5f8781526020019081526020015f2090505f816003015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16036108e7575f7f00000000000000000000000000000000000000000000000000000000000000004261085a9190611076565b905082825f0181905550885f1b826001018190555087826002018190555080826003015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f37449e8a7c9f62932eb6e88550d466720d853433bf2bb22c8ff42bc492b76739878a5f1b836040516108d79392919061138b565b60405180910390a1505050610a17565b5f81600101549050885f1b8103610933576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092a90611430565b60405180910390fd5b5f5f8881526020019081526020015f205f5f82015f9055600182015f9055600282015f9055600382015f6101000a81549067ffffffffffffffff021916905550507f0000000000000000000000000000000000000000000000000000000000000000426109a09190611076565b60015f8981526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fefb80809a6c7055a258032b2ed0b2fb8409d7060f11fa0b4181ecc1d956e9f2787828b5f1b604051610a0b9392919061144e565b60405180910390a15050505b50505050505050565b5f5f5f8381526020019081526020015f206040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505090505f816060015167ffffffffffffffff1603610adc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad39061102b565b60405180910390fd5b42816060015167ffffffffffffffff161115610b2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b24906114f3565b60405180910390fd5b805f0151610b3a83610c1b565b14610b7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7190611581565b60405180910390fd5b5f5f8381526020019081526020015f205f5f82015f9055600182015f9055600282015f9055600382015f6101000a81549067ffffffffffffffff02191690555050806020015160035f8481526020019081526020015f20819055507f5e578c8f35382a852852945932e1082f3805c53821c3a9ff95efa6cd1dc5223382825f01518360200151604051610c0f9392919061144e565b60405180910390a15050565b5f5f60035f8481526020019081526020015f205490505f5f1b8114610c405780610c45565b60015f1b5b915050919050565b60405180608001604052805f81526020015f81526020015f81526020015f67ffffffffffffffff1681525090565b5f67ffffffffffffffff82169050919050565b610c9781610c7b565b82525050565b5f602082019050610cb05f830184610c8e565b92915050565b5f5ffd5b5f5ffd5b5f819050919050565b610cd081610cbe565b8114610cda575f5ffd5b50565b5f81359050610ceb81610cc7565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610d1257610d11610cf1565b5b8235905067ffffffffffffffff811115610d2f57610d2e610cf5565b5b602083019150836001820283011115610d4b57610d4a610cf9565b5b9250929050565b5f5f5f60408486031215610d6957610d68610cb6565b5b5f610d7686828701610cdd565b935050602084013567ffffffffffffffff811115610d9757610d96610cba565b5b610da386828701610cfd565b92509250509250925092565b5f60208284031215610dc457610dc3610cb6565b5b5f610dd184828501610cdd565b91505092915050565b610de381610cbe565b82525050565b5f819050919050565b610dfb81610de9565b82525050565b610e0a81610c7b565b82525050565b608082015f820151610e245f850182610dda565b506020820151610e376020850182610dda565b506040820151610e4a6040850182610df2565b506060820151610e5d6060850182610e01565b50505050565b5f608082019050610e765f830184610e10565b92915050565b5f8115159050919050565b610e9081610e7c565b82525050565b5f602082019050610ea95f830184610e87565b92915050565b610eb881610de9565b8114610ec2575f5ffd5b50565b5f81359050610ed381610eaf565b92915050565b5f5f5f5f5f5f5f60c0888a031215610ef457610ef3610cb6565b5b5f610f018a828b01610cdd565b9750506020610f128a828b01610ec5565b9650506040610f238a828b01610ec5565b9550506060610f348a828b01610cdd565b9450506080610f458a828b01610cdd565b93505060a088013567ffffffffffffffff811115610f6657610f65610cba565b5b610f728a828b01610cfd565b925092505092959891949750929550565b610f8c81610cbe565b82525050565b5f602082019050610fa55f830184610f83565b92915050565b5f82825260208201905092915050565b7f4e6f41415265636f766572793a207265636f76657279206973206e6f742072655f8201527f7175657374656400000000000000000000000000000000000000000000000000602082015250565b5f611015602783610fab565b915061102082610fbb565b604082019050919050565b5f6020820190508181035f83015261104281611009565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61108082610c7b565b915061108b83610c7b565b9250828201905067ffffffffffffffff8111156110ab576110aa611049565b5b92915050565b5f6040820190506110c45f830185610f83565b6110d16020830184610f83565b9392505050565b7f4e6f41415265636f766572793a207265636f76657279206973206e6f74206c6f5f8201527f636b656400000000000000000000000000000000000000000000000000000000602082015250565b5f611132602483610fab565b915061113d826110d8565b604082019050919050565b5f6020820190508181035f83015261115f81611126565b9050919050565b7f4e6f41415265636f766572793a207265636f76657279206973206c6f636b65645f82015250565b5f61119a602083610fab565b91506111a582611166565b602082019050919050565b5f6020820190508181035f8301526111c78161118e565b9050919050565b7f4e6f41415265636f766572793a206964656e746974792063616e206e6f7420625f8201527f65207a65726f0000000000000000000000000000000000000000000000000000602082015250565b5f611228602683610fab565b9150611233826111ce565b604082019050919050565b5f6020820190508181035f8301526112558161121c565b9050919050565b7f4e6f41415265636f766572793a2070617373706f7274206973206c6f636b65645f82015250565b5f611290602083610fab565b915061129b8261125c565b602082019050919050565b5f6020820190508181035f8301526112bd81611284565b9050919050565b7f4e6f41415265636f766572793a206964656e7469747920616c726561647920615f8201527f6374697665000000000000000000000000000000000000000000000000000000602082015250565b5f61131e602583610fab565b9150611329826112c4565b604082019050919050565b5f6020820190508181035f83015261134b81611312565b9050919050565b5f819050919050565b5f61137561137061136b84610c7b565b611352565b610de9565b9050919050565b6113858161135b565b82525050565b5f60608201905061139e5f830186610f83565b6113ab6020830185610f83565b6113b8604083018461137c565b949350505050565b7f4e6f41415265636f766572793a207265636f7665727920616c726561647920725f8201527f6571756573746564000000000000000000000000000000000000000000000000602082015250565b5f61141a602883610fab565b9150611425826113c0565b604082019050919050565b5f6020820190508181035f8301526114478161140e565b9050919050565b5f6060820190506114615f830186610f83565b61146e6020830185610f83565b61147b6040830184610f83565b949350505050565b7f4e6f41415265636f766572793a207265636f76657279206973206e6f742072655f8201527f6164790000000000000000000000000000000000000000000000000000000000602082015250565b5f6114dd602383610fab565b91506114e882611483565b604082019050919050565b5f6020820190508181035f83015261150a816114d1565b9050919050565b7f4e6f41415265636f766572793a20616374697665206964656e746974792063685f8201527f616e676564000000000000000000000000000000000000000000000000000000602082015250565b5f61156b602583610fab565b915061157682611511565b604082019050919050565b5f6020820190508181035f8301526115988161155f565b905091905056fea26469706673582212209aa2f982724e4d103ce8367db9544745a792b04f5ecfdb55749379f5b40be17464736f6c634300081c0033",
}

// MockNoAARecoveryABI is the input ABI used to generate the binding from.
// Deprecated: Use MockNoAARecoveryMetaData.ABI instead.
var MockNoAARecoveryABI = MockNoAARecoveryMetaData.ABI

// MockNoAARecoveryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockNoAARecoveryMetaData.Bin instead.
var MockNoAARecoveryBin = MockNoAARecoveryMetaData.Bin

// DeployMockNoAARecovery deploys a new Ethereum contract, binding an instance of MockNoAARecovery to it.
func DeployMockNoAARecovery(auth *bind.TransactOpts, backend bind.ContractBackend, recoveryDelay_ uint64, lockDuration_ uint64) (common.Address, *types.Transaction, *MockNoAARecovery, error) {
	parsed, err := MockNoAARecoveryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockNoAARecoveryBin), backend, recoveryDelay_, lockDuration_)
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

// ActiveIdentity is a free data retrieval call binding the contract method 0xfe7b8de1.
//
// Solidity: function activeIdentity(bytes32 passportHash_) view returns(bytes32)
func (_MockNoAARecovery *MockNoAARecoveryCaller) ActiveIdentity(opts *bind.CallOpts, passportHash_ [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MockNoAARecovery.contract.Call(opts, &out, "activeIdentity", passportHash_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActiveIdentity is a free data retrieval call binding the contract method 0xfe7b8de1.
//
// Solidity: function activeIdentity(bytes32 passportHash_) view returns(bytes32)
func (_MockNoAARecovery *MockNoAARecoverySession) ActiveIdentity(passportHash_ [32]byte) ([32]byte, error) {
	return _MockNoAARecovery.Contract.ActiveIdentity(&_MockNoAARecovery.CallOpts, passportHash_)
}

// ActiveIdentity is a free data retrieval call binding the contract method 0xfe7b8de1.
//
// Solidity: function activeIdentity(bytes32 passportHash_) view returns(bytes32)
func (_MockNoAARecovery *MockNoAARecoveryCallerSession) ActiveIdentity(passportHash_ [32]byte) ([32]byte, error) {
	return _MockNoAARecovery.Contract.ActiveIdentity(&_MockNoAARecovery.CallOpts, passportHash_)
}

// GetRecoveryRequest is a free data retrieval call binding the contract method 0x5ea3ec76.
//
// Solidity: function getRecoveryRequest(bytes32 passportHash_) view returns((bytes32,bytes32,uint256,uint64))
func (_MockNoAARecovery *MockNoAARecoveryCaller) GetRecoveryRequest(opts *bind.CallOpts, passportHash_ [32]byte) (INoAARecoveryRecoveryRequest, error) {
	var out []interface{}
	err := _MockNoAARecovery.contract.Call(opts, &out, "getRecoveryRequest", passportHash_)

	if err != nil {
		return *new(INoAARecoveryRecoveryRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(INoAARecoveryRecoveryRequest)).(*INoAARecoveryRecoveryRequest)

	return out0, err

}

// GetRecoveryRequest is a free data retrieval call binding the contract method 0x5ea3ec76.
//
// Solidity: function getRecoveryRequest(bytes32 passportHash_) view returns((bytes32,bytes32,uint256,uint64))
func (_MockNoAARecovery *MockNoAARecoverySession) GetRecoveryRequest(passportHash_ [32]byte) (INoAARecoveryRecoveryRequest, error) {
	return _MockNoAARecovery.Contract.GetRecoveryRequest(&_MockNoAARecovery.CallOpts, passportHash_)
}

// GetRecoveryRequest is a free data retrieval call binding the contract method 0x5ea3ec76.
//
// Solidity: function getRecoveryRequest(bytes32 passportHash_) view returns((bytes32,bytes32,uint256,uint64))
func (_MockNoAARecovery *MockNoAARecoveryCallerSession) GetRecoveryRequest(passportHash_ [32]byte) (INoAARecoveryRecoveryRequest, error) {
	return _MockNoAARecovery.Contract.GetRecoveryRequest(&_MockNoAARecovery.CallOpts, passportHash_)
}

// IsRecoveryLocked is a free data retrieval call binding the contract method 0x76890e32.
//
// Solidity: function isRecoveryLocked(bytes32 ) view returns(bool)
func (_MockNoAARecovery *MockNoAARecoveryCaller) IsRecoveryLocked(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MockNoAARecovery.contract.Call(opts, &out, "isRecoveryLocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRecoveryLocked is a free data retrieval call binding the contract method 0x76890e32.
//
// Solidity: function isRecoveryLocked(bytes32 ) view returns(bool)
func (_MockNoAARecovery *MockNoAARecoverySession) IsRecoveryLocked(arg0 [32]byte) (bool, error) {
	return _MockNoAARecovery.Contract.IsRecoveryLocked(&_MockNoAARecovery.CallOpts, arg0)
}

// IsRecoveryLocked is a free data retrieval call binding the contract method 0x76890e32.
//
// Solidity: function isRecoveryLocked(bytes32 ) view returns(bool)
func (_MockNoAARecovery *MockNoAARecoveryCallerSession) IsRecoveryLocked(arg0 [32]byte) (bool, error) {
	return _MockNoAARecovery.Contract.IsRecoveryLocked(&_MockNoAARecovery.CallOpts, arg0)
}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoveryCaller) LockDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockNoAARecovery.contract.Call(opts, &out, "lockDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoverySession) LockDuration() (uint64, error) {
	return _MockNoAARecovery.Contract.LockDuration(&_MockNoAARecovery.CallOpts)
}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint64)
func (_MockNoAARecovery *MockNoAARecoveryCallerSession) LockDuration() (uint64, error) {
	return _MockNoAARecovery.Contract.LockDuration(&_MockNoAARecovery.CallOpts)
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

// CancelRecovery is a paid mutator transaction binding the contract method 0x57164456.
//
// Solidity: function cancelRecovery(bytes32 passportHash_, bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactor) CancelRecovery(opts *bind.TransactOpts, passportHash_ [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.contract.Transact(opts, "cancelRecovery", passportHash_, arg1)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x57164456.
//
// Solidity: function cancelRecovery(bytes32 passportHash_, bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoverySession) CancelRecovery(passportHash_ [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.CancelRecovery(&_MockNoAARecovery.TransactOpts, passportHash_, arg1)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x57164456.
//
// Solidity: function cancelRecovery(bytes32 passportHash_, bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactorSession) CancelRecovery(passportHash_ [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.CancelRecovery(&_MockNoAARecovery.TransactOpts, passportHash_, arg1)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash_) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactor) FinalizeRecovery(opts *bind.TransactOpts, passportHash_ [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.contract.Transact(opts, "finalizeRecovery", passportHash_)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash_) returns()
func (_MockNoAARecovery *MockNoAARecoverySession) FinalizeRecovery(passportHash_ [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.FinalizeRecovery(&_MockNoAARecovery.TransactOpts, passportHash_)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0xe99d9d48.
//
// Solidity: function finalizeRecovery(bytes32 passportHash_) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactorSession) FinalizeRecovery(passportHash_ [32]byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.FinalizeRecovery(&_MockNoAARecovery.TransactOpts, passportHash_)
}

// LockRecovery is a paid mutator transaction binding the contract method 0x9d754de8.
//
// Solidity: function lockRecovery(bytes32 passportHash_, bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactor) LockRecovery(opts *bind.TransactOpts, passportHash_ [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.contract.Transact(opts, "lockRecovery", passportHash_, arg1)
}

// LockRecovery is a paid mutator transaction binding the contract method 0x9d754de8.
//
// Solidity: function lockRecovery(bytes32 passportHash_, bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoverySession) LockRecovery(passportHash_ [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.LockRecovery(&_MockNoAARecovery.TransactOpts, passportHash_, arg1)
}

// LockRecovery is a paid mutator transaction binding the contract method 0x9d754de8.
//
// Solidity: function lockRecovery(bytes32 passportHash_, bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactorSession) LockRecovery(passportHash_ [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.LockRecovery(&_MockNoAARecovery.TransactOpts, passportHash_, arg1)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0xd3e2a3df.
//
// Solidity: function requestRecovery(bytes32 , uint256 identityKey_, uint256 dgCommit_, bytes32 passportHash_, bytes32 , bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactor) RequestRecovery(opts *bind.TransactOpts, arg0 [32]byte, identityKey_ *big.Int, dgCommit_ *big.Int, passportHash_ [32]byte, arg4 [32]byte, arg5 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.contract.Transact(opts, "requestRecovery", arg0, identityKey_, dgCommit_, passportHash_, arg4, arg5)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0xd3e2a3df.
//
// Solidity: function requestRecovery(bytes32 , uint256 identityKey_, uint256 dgCommit_, bytes32 passportHash_, bytes32 , bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoverySession) RequestRecovery(arg0 [32]byte, identityKey_ *big.Int, dgCommit_ *big.Int, passportHash_ [32]byte, arg4 [32]byte, arg5 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.RequestRecovery(&_MockNoAARecovery.TransactOpts, arg0, identityKey_, dgCommit_, passportHash_, arg4, arg5)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0xd3e2a3df.
//
// Solidity: function requestRecovery(bytes32 , uint256 identityKey_, uint256 dgCommit_, bytes32 passportHash_, bytes32 , bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactorSession) RequestRecovery(arg0 [32]byte, identityKey_ *big.Int, dgCommit_ *big.Int, passportHash_ [32]byte, arg4 [32]byte, arg5 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.RequestRecovery(&_MockNoAARecovery.TransactOpts, arg0, identityKey_, dgCommit_, passportHash_, arg4, arg5)
}

// UnlockRecovery is a paid mutator transaction binding the contract method 0x79204b68.
//
// Solidity: function unlockRecovery(bytes32 passportHash_, bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactor) UnlockRecovery(opts *bind.TransactOpts, passportHash_ [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.contract.Transact(opts, "unlockRecovery", passportHash_, arg1)
}

// UnlockRecovery is a paid mutator transaction binding the contract method 0x79204b68.
//
// Solidity: function unlockRecovery(bytes32 passportHash_, bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoverySession) UnlockRecovery(passportHash_ [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.UnlockRecovery(&_MockNoAARecovery.TransactOpts, passportHash_, arg1)
}

// UnlockRecovery is a paid mutator transaction binding the contract method 0x79204b68.
//
// Solidity: function unlockRecovery(bytes32 passportHash_, bytes ) returns()
func (_MockNoAARecovery *MockNoAARecoveryTransactorSession) UnlockRecovery(passportHash_ [32]byte, arg1 []byte) (*types.Transaction, error) {
	return _MockNoAARecovery.Contract.UnlockRecovery(&_MockNoAARecovery.TransactOpts, passportHash_, arg1)
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
	PassportHash       [32]byte
	PendingIdentityKey [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRecoveryCancelled is a free log retrieval operation binding the contract event 0x9779648ab6a63d786fd726dd9c5f5873960dddca6252b628760bfd70e729a71d.
//
// Solidity: event RecoveryCancelled(bytes32 passportHash, bytes32 pendingIdentityKey)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryCancelled(opts *bind.FilterOpts) (*MockNoAARecoveryRecoveryCancelledIterator, error) {

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryCancelled")
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryCancelledIterator{contract: _MockNoAARecovery.contract, event: "RecoveryCancelled", logs: logs, sub: sub}, nil
}

// WatchRecoveryCancelled is a free log subscription operation binding the contract event 0x9779648ab6a63d786fd726dd9c5f5873960dddca6252b628760bfd70e729a71d.
//
// Solidity: event RecoveryCancelled(bytes32 passportHash, bytes32 pendingIdentityKey)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryCancelled(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryCancelled) (event.Subscription, error) {

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryCancelled")
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

// ParseRecoveryCancelled is a log parse operation binding the contract event 0x9779648ab6a63d786fd726dd9c5f5873960dddca6252b628760bfd70e729a71d.
//
// Solidity: event RecoveryCancelled(bytes32 passportHash, bytes32 pendingIdentityKey)
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
	PassportHash       [32]byte
	PendingIdentityKey [32]byte
	NewIdentityKey     [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRecoveryContested is a free log retrieval operation binding the contract event 0xefb80809a6c7055a258032b2ed0b2fb8409d7060f11fa0b4181ecc1d956e9f27.
//
// Solidity: event RecoveryContested(bytes32 passportHash, bytes32 pendingIdentityKey, bytes32 newIdentityKey)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryContested(opts *bind.FilterOpts) (*MockNoAARecoveryRecoveryContestedIterator, error) {

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryContested")
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryContestedIterator{contract: _MockNoAARecovery.contract, event: "RecoveryContested", logs: logs, sub: sub}, nil
}

// WatchRecoveryContested is a free log subscription operation binding the contract event 0xefb80809a6c7055a258032b2ed0b2fb8409d7060f11fa0b4181ecc1d956e9f27.
//
// Solidity: event RecoveryContested(bytes32 passportHash, bytes32 pendingIdentityKey, bytes32 newIdentityKey)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryContested(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryContested) (event.Subscription, error) {

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryContested")
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

// ParseRecoveryContested is a log parse operation binding the contract event 0xefb80809a6c7055a258032b2ed0b2fb8409d7060f11fa0b4181ecc1d956e9f27.
//
// Solidity: event RecoveryContested(bytes32 passportHash, bytes32 pendingIdentityKey, bytes32 newIdentityKey)
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
	OldIdentityKey [32]byte
	NewIdentityKey [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRecoveryFinalized is a free log retrieval operation binding the contract event 0x5e578c8f35382a852852945932e1082f3805c53821c3a9ff95efa6cd1dc52233.
//
// Solidity: event RecoveryFinalized(bytes32 passportHash, bytes32 oldIdentityKey, bytes32 newIdentityKey)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryFinalized(opts *bind.FilterOpts) (*MockNoAARecoveryRecoveryFinalizedIterator, error) {

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryFinalized")
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryFinalizedIterator{contract: _MockNoAARecovery.contract, event: "RecoveryFinalized", logs: logs, sub: sub}, nil
}

// WatchRecoveryFinalized is a free log subscription operation binding the contract event 0x5e578c8f35382a852852945932e1082f3805c53821c3a9ff95efa6cd1dc52233.
//
// Solidity: event RecoveryFinalized(bytes32 passportHash, bytes32 oldIdentityKey, bytes32 newIdentityKey)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryFinalized(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryFinalized) (event.Subscription, error) {

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryFinalized")
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

// ParseRecoveryFinalized is a log parse operation binding the contract event 0x5e578c8f35382a852852945932e1082f3805c53821c3a9ff95efa6cd1dc52233.
//
// Solidity: event RecoveryFinalized(bytes32 passportHash, bytes32 oldIdentityKey, bytes32 newIdentityKey)
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
// Solidity: event RecoveryLocked(bytes32 passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryLocked(opts *bind.FilterOpts) (*MockNoAARecoveryRecoveryLockedIterator, error) {

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryLocked")
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryLockedIterator{contract: _MockNoAARecovery.contract, event: "RecoveryLocked", logs: logs, sub: sub}, nil
}

// WatchRecoveryLocked is a free log subscription operation binding the contract event 0xf1d7ea8fd15a531788ee7d98d5abf152e714bc1b1747fa7ec6dc5f44e3c29a8d.
//
// Solidity: event RecoveryLocked(bytes32 passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryLocked(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryLocked) (event.Subscription, error) {

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryLocked")
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
// Solidity: event RecoveryLocked(bytes32 passportHash)
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
	NewIdentityKey [32]byte
	ReadyAt        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRecoveryRequested is a free log retrieval operation binding the contract event 0x37449e8a7c9f62932eb6e88550d466720d853433bf2bb22c8ff42bc492b76739.
//
// Solidity: event RecoveryRequested(bytes32 passportHash, bytes32 newIdentityKey, uint256 readyAt)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryRequested(opts *bind.FilterOpts) (*MockNoAARecoveryRecoveryRequestedIterator, error) {

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryRequested")
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryRequestedIterator{contract: _MockNoAARecovery.contract, event: "RecoveryRequested", logs: logs, sub: sub}, nil
}

// WatchRecoveryRequested is a free log subscription operation binding the contract event 0x37449e8a7c9f62932eb6e88550d466720d853433bf2bb22c8ff42bc492b76739.
//
// Solidity: event RecoveryRequested(bytes32 passportHash, bytes32 newIdentityKey, uint256 readyAt)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryRequested(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryRequested) (event.Subscription, error) {

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryRequested")
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

// ParseRecoveryRequested is a log parse operation binding the contract event 0x37449e8a7c9f62932eb6e88550d466720d853433bf2bb22c8ff42bc492b76739.
//
// Solidity: event RecoveryRequested(bytes32 passportHash, bytes32 newIdentityKey, uint256 readyAt)
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
// Solidity: event RecoveryUnlocked(bytes32 passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) FilterRecoveryUnlocked(opts *bind.FilterOpts) (*MockNoAARecoveryRecoveryUnlockedIterator, error) {

	logs, sub, err := _MockNoAARecovery.contract.FilterLogs(opts, "RecoveryUnlocked")
	if err != nil {
		return nil, err
	}
	return &MockNoAARecoveryRecoveryUnlockedIterator{contract: _MockNoAARecovery.contract, event: "RecoveryUnlocked", logs: logs, sub: sub}, nil
}

// WatchRecoveryUnlocked is a free log subscription operation binding the contract event 0xcf4323666218fdee66018eb7566c179ce553334f33de42e2ee711314977d87f9.
//
// Solidity: event RecoveryUnlocked(bytes32 passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) WatchRecoveryUnlocked(opts *bind.WatchOpts, sink chan<- *MockNoAARecoveryRecoveryUnlocked) (event.Subscription, error) {

	logs, sub, err := _MockNoAARecovery.contract.WatchLogs(opts, "RecoveryUnlocked")
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
// Solidity: event RecoveryUnlocked(bytes32 passportHash)
func (_MockNoAARecovery *MockNoAARecoveryFilterer) ParseRecoveryUnlocked(log types.Log) (*MockNoAARecoveryRecoveryUnlocked, error) {
	event := new(MockNoAARecoveryRecoveryUnlocked)
	if err := _MockNoAARecovery.contract.UnpackLog(event, "RecoveryUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
