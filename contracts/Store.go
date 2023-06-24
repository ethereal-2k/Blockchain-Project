// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"compareAccessPolicy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"compareMeta1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"compareMeta2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPKC\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPKF\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"input1\",\"type\":\"string\"}],\"name\":\"setAccessP\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"input1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"input2\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"input3\",\"type\":\"bytes\"}],\"name\":\"setMeta1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"input1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"input2\",\"type\":\"string\"}],\"name\":\"setMeta2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setPKC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setPKF\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input2\",\"type\":\"bytes\"}],\"name\":\"setParams\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506115db806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806392e0bdfd1161008c578063d087d28811610066578063d087d28814610a46578063d0ec2e3a14610ac9578063d41c4b6014610b4c578063ffd58c6d14610c9e576100ea565b806392e0bdfd146106cf578063be6a2a94146107a2578063c78e0f711461085d576100ea565b80635e615a6b116100c85780635e615a6b146103cf5780637bebc55b146104be57806386e08293146105795780638ee3d37a1461064c576100ea565b80630944a67f146100ef57806334e784a1146101c257806339a40fe91461027d575b600080fd5b6101a86004803603602081101561010557600080fd5b810190808035906020019064010000000081111561012257600080fd5b82018360208201111561013457600080fd5b8035906020019184600183028401116401000000008311171561015657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610d59565b604051808215151515815260200191505060405180910390f35b61027b600480360360208110156101d857600080fd5b81019080803590602001906401000000008111156101f557600080fd5b82018360208201111561020757600080fd5b8035906020019184600183028401116401000000008311171561022957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610e50565b005b6103cd6004803603604081101561029357600080fd5b81019080803590602001906401000000008111156102b057600080fd5b8201836020820111156102c257600080fd5b803590602001918460018302840111640100000000831117156102e457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561034757600080fd5b82018360208201111561035957600080fd5b8035906020019184600183028401116401000000008311171561037b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610e6a565b005b6103d7610e9c565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561041b578082015181840152602081019050610400565b50505050905090810190601f1680156104485780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610481578082015181840152602081019050610466565b50505050905090810190601f1680156104ae5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b610577600480360360208110156104d457600080fd5b81019080803590602001906401000000008111156104f157600080fd5b82018360208201111561050357600080fd5b8035906020019184600183028401116401000000008311171561052557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610fe3565b005b6106326004803603602081101561058f57600080fd5b81019080803590602001906401000000008111156105ac57600080fd5b8201836020820111156105be57600080fd5b803590602001918460018302840111640100000000831117156105e057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610ffd565b604051808215151515815260200191505060405180910390f35b6106546110f4565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610694578082015181840152602081019050610679565b50505050905090810190601f1680156106c15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610788600480360360208110156106e557600080fd5b810190808035906020019064010000000081111561070257600080fd5b82018360208201111561071457600080fd5b8035906020019184600183028401116401000000008311171561073657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611196565b604051808215151515815260200191505060405180910390f35b61085b600480360360208110156107b857600080fd5b81019080803590602001906401000000008111156107d557600080fd5b8201836020820111156107e757600080fd5b8035906020019184600183028401116401000000008311171561080957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061128d565b005b610a446004803603606081101561087357600080fd5b810190808035906020019064010000000081111561089057600080fd5b8201836020820111156108a257600080fd5b803590602001918460018302840111640100000000831117156108c457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561092757600080fd5b82018360208201111561093957600080fd5b8035906020019184600183028401116401000000008311171561095b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156109be57600080fd5b8201836020820111156109d057600080fd5b803590602001918460018302840111640100000000831117156109f257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506112a7565b005b610a4e6112f1565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a8e578082015181840152602081019050610a73565b50505050905090810190601f168015610abb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610ad1611393565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610b11578082015181840152602081019050610af6565b50505050905090810190601f168015610b3e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610c9c60048036036040811015610b6257600080fd5b8101908080359060200190640100000000811115610b7f57600080fd5b820183602082011115610b9157600080fd5b80359060200191846001830284011164010000000083111715610bb357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190640100000000811115610c1657600080fd5b820183602082011115610c2857600080fd5b80359060200191846001830284011164010000000083111715610c4a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611435565b005b610d5760048036036020811015610cb457600080fd5b8101908080359060200190640100000000811115610cd157600080fd5b820183602082011115610ce357600080fd5b80359060200191846001830284011164010000000083111715610d0557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611467565b005b6000816040516020018082805190602001908083835b60208310610d925780518252602082019150602081019050602083039250610d6f565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060056040516020018082805460018160011615610100020316600290048015610e2d5780601f10610e0b576101008083540402835291820191610e2d565b820191906000526020600020905b815481529060010190602001808311610e19575b505091505060405160208183030381529060405280519060200120149050919050565b8060039080519060200190610e66929190611481565b5050565b8160009080519060200190610e80929190611481565b508060019080519060200190610e97929190611481565b505050565b60608060006001818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f385780601f10610f0d57610100808354040283529160200191610f38565b820191906000526020600020905b815481529060010190602001808311610f1b57829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fd45780601f10610fa957610100808354040283529160200191610fd4565b820191906000526020600020905b815481529060010190602001808311610fb757829003601f168201915b50505050509050915091509091565b80600a9080519060200190610ff9929190611481565b5050565b6000816040516020018082805190602001908083835b602083106110365780518252602082019150602081019050602083039250611013565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120600860405160200180828054600181600116156101000203166002900480156110d15780601f106110af5761010080835404028352918201916110d1565b820191906000526020600020905b8154815290600101906020018083116110bd575b505091505060405160208183030381529060405280519060200120149050919050565b606060038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561118c5780601f106111615761010080835404028352916020019161118c565b820191906000526020600020905b81548152906001019060200180831161116f57829003601f168201915b5050505050905090565b6000816040516020018082805190602001908083835b602083106111cf57805182526020820191506020810190506020830392506111ac565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001206009604051602001808280546001816001161561010002031660029004801561126a5780601f1061124857610100808354040283529182019161126a565b820191906000526020600020905b815481529060010190602001808311611256575b505091505060405160208183030381529060405280519060200120149050919050565b80600290805190602001906112a3929190611481565b5050565b82600490805190602001906112bd929190611501565b5081600590805190602001906112d4929190611501565b5080600690805190602001906112eb929190611481565b50505050565b6060600a8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113895780601f1061135e57610100808354040283529160200191611389565b820191906000526020600020905b81548152906001019060200180831161136c57829003601f168201915b5050505050905090565b606060028054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561142b5780601f106114005761010080835404028352916020019161142b565b820191906000526020600020905b81548152906001019060200180831161140e57829003601f168201915b5050505050905090565b816007908051906020019061144b929190611501565b508060089080519060200190611462929190611501565b505050565b806009908051906020019061147d929190611501565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106114c257805160ff19168380011785556114f0565b828001600101855582156114f0579182015b828111156114ef5782518255916020019190600101906114d4565b5b5090506114fd9190611581565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061154257805160ff1916838001178555611570565b82800160010185558215611570579182015b8281111561156f578251825591602001919060010190611554565b5b50905061157d9190611581565b5090565b6115a391905b8082111561159f576000816000905550600101611587565b5090565b9056fea265627a7a72315820ff28cbe5080c21d252e7a2e5cacb21f18193b6ce742173768c6dccf51811a3bc64736f6c63430005100032",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// CompareAccessPolicy is a free data retrieval call binding the contract method 0x92e0bdfd.
//
// Solidity: function compareAccessPolicy(string req) view returns(bool)
func (_Store *StoreCaller) CompareAccessPolicy(opts *bind.CallOpts, req string) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "compareAccessPolicy", req)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareAccessPolicy is a free data retrieval call binding the contract method 0x92e0bdfd.
//
// Solidity: function compareAccessPolicy(string req) view returns(bool)
func (_Store *StoreSession) CompareAccessPolicy(req string) (bool, error) {
	return _Store.Contract.CompareAccessPolicy(&_Store.CallOpts, req)
}

// CompareAccessPolicy is a free data retrieval call binding the contract method 0x92e0bdfd.
//
// Solidity: function compareAccessPolicy(string req) view returns(bool)
func (_Store *StoreCallerSession) CompareAccessPolicy(req string) (bool, error) {
	return _Store.Contract.CompareAccessPolicy(&_Store.CallOpts, req)
}

// CompareMeta1 is a free data retrieval call binding the contract method 0x0944a67f.
//
// Solidity: function compareMeta1(string req) view returns(bool)
func (_Store *StoreCaller) CompareMeta1(opts *bind.CallOpts, req string) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "compareMeta1", req)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareMeta1 is a free data retrieval call binding the contract method 0x0944a67f.
//
// Solidity: function compareMeta1(string req) view returns(bool)
func (_Store *StoreSession) CompareMeta1(req string) (bool, error) {
	return _Store.Contract.CompareMeta1(&_Store.CallOpts, req)
}

// CompareMeta1 is a free data retrieval call binding the contract method 0x0944a67f.
//
// Solidity: function compareMeta1(string req) view returns(bool)
func (_Store *StoreCallerSession) CompareMeta1(req string) (bool, error) {
	return _Store.Contract.CompareMeta1(&_Store.CallOpts, req)
}

// CompareMeta2 is a free data retrieval call binding the contract method 0x86e08293.
//
// Solidity: function compareMeta2(string req) view returns(bool)
func (_Store *StoreCaller) CompareMeta2(opts *bind.CallOpts, req string) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "compareMeta2", req)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareMeta2 is a free data retrieval call binding the contract method 0x86e08293.
//
// Solidity: function compareMeta2(string req) view returns(bool)
func (_Store *StoreSession) CompareMeta2(req string) (bool, error) {
	return _Store.Contract.CompareMeta2(&_Store.CallOpts, req)
}

// CompareMeta2 is a free data retrieval call binding the contract method 0x86e08293.
//
// Solidity: function compareMeta2(string req) view returns(bool)
func (_Store *StoreCallerSession) CompareMeta2(req string) (bool, error) {
	return _Store.Contract.CompareMeta2(&_Store.CallOpts, req)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(bytes)
func (_Store *StoreCaller) GetNonce(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(bytes)
func (_Store *StoreSession) GetNonce() ([]byte, error) {
	return _Store.Contract.GetNonce(&_Store.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(bytes)
func (_Store *StoreCallerSession) GetNonce() ([]byte, error) {
	return _Store.Contract.GetNonce(&_Store.CallOpts)
}

// GetPKC is a free data retrieval call binding the contract method 0xd0ec2e3a.
//
// Solidity: function getPKC() view returns(bytes)
func (_Store *StoreCaller) GetPKC(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getPKC")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPKC is a free data retrieval call binding the contract method 0xd0ec2e3a.
//
// Solidity: function getPKC() view returns(bytes)
func (_Store *StoreSession) GetPKC() ([]byte, error) {
	return _Store.Contract.GetPKC(&_Store.CallOpts)
}

// GetPKC is a free data retrieval call binding the contract method 0xd0ec2e3a.
//
// Solidity: function getPKC() view returns(bytes)
func (_Store *StoreCallerSession) GetPKC() ([]byte, error) {
	return _Store.Contract.GetPKC(&_Store.CallOpts)
}

// GetPKF is a free data retrieval call binding the contract method 0x8ee3d37a.
//
// Solidity: function getPKF() view returns(bytes)
func (_Store *StoreCaller) GetPKF(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getPKF")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPKF is a free data retrieval call binding the contract method 0x8ee3d37a.
//
// Solidity: function getPKF() view returns(bytes)
func (_Store *StoreSession) GetPKF() ([]byte, error) {
	return _Store.Contract.GetPKF(&_Store.CallOpts)
}

// GetPKF is a free data retrieval call binding the contract method 0x8ee3d37a.
//
// Solidity: function getPKF() view returns(bytes)
func (_Store *StoreCallerSession) GetPKF() ([]byte, error) {
	return _Store.Contract.GetPKF(&_Store.CallOpts)
}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns(bytes, bytes)
func (_Store *StoreCaller) GetParams(opts *bind.CallOpts) ([]byte, []byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getParams")

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns(bytes, bytes)
func (_Store *StoreSession) GetParams() ([]byte, []byte, error) {
	return _Store.Contract.GetParams(&_Store.CallOpts)
}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns(bytes, bytes)
func (_Store *StoreCallerSession) GetParams() ([]byte, []byte, error) {
	return _Store.Contract.GetParams(&_Store.CallOpts)
}

// SetAccessP is a paid mutator transaction binding the contract method 0xffd58c6d.
//
// Solidity: function setAccessP(string input1) returns()
func (_Store *StoreTransactor) SetAccessP(opts *bind.TransactOpts, input1 string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setAccessP", input1)
}

// SetAccessP is a paid mutator transaction binding the contract method 0xffd58c6d.
//
// Solidity: function setAccessP(string input1) returns()
func (_Store *StoreSession) SetAccessP(input1 string) (*types.Transaction, error) {
	return _Store.Contract.SetAccessP(&_Store.TransactOpts, input1)
}

// SetAccessP is a paid mutator transaction binding the contract method 0xffd58c6d.
//
// Solidity: function setAccessP(string input1) returns()
func (_Store *StoreTransactorSession) SetAccessP(input1 string) (*types.Transaction, error) {
	return _Store.Contract.SetAccessP(&_Store.TransactOpts, input1)
}

// SetMeta1 is a paid mutator transaction binding the contract method 0xc78e0f71.
//
// Solidity: function setMeta1(string input1, string input2, bytes input3) returns()
func (_Store *StoreTransactor) SetMeta1(opts *bind.TransactOpts, input1 string, input2 string, input3 []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setMeta1", input1, input2, input3)
}

// SetMeta1 is a paid mutator transaction binding the contract method 0xc78e0f71.
//
// Solidity: function setMeta1(string input1, string input2, bytes input3) returns()
func (_Store *StoreSession) SetMeta1(input1 string, input2 string, input3 []byte) (*types.Transaction, error) {
	return _Store.Contract.SetMeta1(&_Store.TransactOpts, input1, input2, input3)
}

// SetMeta1 is a paid mutator transaction binding the contract method 0xc78e0f71.
//
// Solidity: function setMeta1(string input1, string input2, bytes input3) returns()
func (_Store *StoreTransactorSession) SetMeta1(input1 string, input2 string, input3 []byte) (*types.Transaction, error) {
	return _Store.Contract.SetMeta1(&_Store.TransactOpts, input1, input2, input3)
}

// SetMeta2 is a paid mutator transaction binding the contract method 0xd41c4b60.
//
// Solidity: function setMeta2(string input1, string input2) returns()
func (_Store *StoreTransactor) SetMeta2(opts *bind.TransactOpts, input1 string, input2 string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setMeta2", input1, input2)
}

// SetMeta2 is a paid mutator transaction binding the contract method 0xd41c4b60.
//
// Solidity: function setMeta2(string input1, string input2) returns()
func (_Store *StoreSession) SetMeta2(input1 string, input2 string) (*types.Transaction, error) {
	return _Store.Contract.SetMeta2(&_Store.TransactOpts, input1, input2)
}

// SetMeta2 is a paid mutator transaction binding the contract method 0xd41c4b60.
//
// Solidity: function setMeta2(string input1, string input2) returns()
func (_Store *StoreTransactorSession) SetMeta2(input1 string, input2 string) (*types.Transaction, error) {
	return _Store.Contract.SetMeta2(&_Store.TransactOpts, input1, input2)
}

// SetNonce is a paid mutator transaction binding the contract method 0x7bebc55b.
//
// Solidity: function setNonce(bytes input) returns()
func (_Store *StoreTransactor) SetNonce(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setNonce", input)
}

// SetNonce is a paid mutator transaction binding the contract method 0x7bebc55b.
//
// Solidity: function setNonce(bytes input) returns()
func (_Store *StoreSession) SetNonce(input []byte) (*types.Transaction, error) {
	return _Store.Contract.SetNonce(&_Store.TransactOpts, input)
}

// SetNonce is a paid mutator transaction binding the contract method 0x7bebc55b.
//
// Solidity: function setNonce(bytes input) returns()
func (_Store *StoreTransactorSession) SetNonce(input []byte) (*types.Transaction, error) {
	return _Store.Contract.SetNonce(&_Store.TransactOpts, input)
}

// SetPKC is a paid mutator transaction binding the contract method 0xbe6a2a94.
//
// Solidity: function setPKC(bytes input) returns()
func (_Store *StoreTransactor) SetPKC(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setPKC", input)
}

// SetPKC is a paid mutator transaction binding the contract method 0xbe6a2a94.
//
// Solidity: function setPKC(bytes input) returns()
func (_Store *StoreSession) SetPKC(input []byte) (*types.Transaction, error) {
	return _Store.Contract.SetPKC(&_Store.TransactOpts, input)
}

// SetPKC is a paid mutator transaction binding the contract method 0xbe6a2a94.
//
// Solidity: function setPKC(bytes input) returns()
func (_Store *StoreTransactorSession) SetPKC(input []byte) (*types.Transaction, error) {
	return _Store.Contract.SetPKC(&_Store.TransactOpts, input)
}

// SetPKF is a paid mutator transaction binding the contract method 0x34e784a1.
//
// Solidity: function setPKF(bytes input) returns()
func (_Store *StoreTransactor) SetPKF(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setPKF", input)
}

// SetPKF is a paid mutator transaction binding the contract method 0x34e784a1.
//
// Solidity: function setPKF(bytes input) returns()
func (_Store *StoreSession) SetPKF(input []byte) (*types.Transaction, error) {
	return _Store.Contract.SetPKF(&_Store.TransactOpts, input)
}

// SetPKF is a paid mutator transaction binding the contract method 0x34e784a1.
//
// Solidity: function setPKF(bytes input) returns()
func (_Store *StoreTransactorSession) SetPKF(input []byte) (*types.Transaction, error) {
	return _Store.Contract.SetPKF(&_Store.TransactOpts, input)
}

// SetParams is a paid mutator transaction binding the contract method 0x39a40fe9.
//
// Solidity: function setParams(bytes input1, bytes input2) returns()
func (_Store *StoreTransactor) SetParams(opts *bind.TransactOpts, input1 []byte, input2 []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setParams", input1, input2)
}

// SetParams is a paid mutator transaction binding the contract method 0x39a40fe9.
//
// Solidity: function setParams(bytes input1, bytes input2) returns()
func (_Store *StoreSession) SetParams(input1 []byte, input2 []byte) (*types.Transaction, error) {
	return _Store.Contract.SetParams(&_Store.TransactOpts, input1, input2)
}

// SetParams is a paid mutator transaction binding the contract method 0x39a40fe9.
//
// Solidity: function setParams(bytes input1, bytes input2) returns()
func (_Store *StoreTransactorSession) SetParams(input1 []byte, input2 []byte) (*types.Transaction, error) {
	return _Store.Contract.SetParams(&_Store.TransactOpts, input1, input2)
}
