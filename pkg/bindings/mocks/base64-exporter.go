// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Base64ExporterABI is the input ABI used to generate the binding from.
const Base64ExporterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_encoded\",\"type\":\"bytes\"}],\"name\":\"base64decode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Base64ExporterBin is the compiled bytecode used for deploying new contracts.
const Base64ExporterBin = `608060405234801561001057600080fd5b50610bd7806100206000396000f3006080604052600436106100405763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633a718bfb8114610045575b600080fd5b34801561005157600080fd5b5061006560048035602481019101356100da565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561009f578181015183820152602001610087565b50505050905090810190601f1680156100cc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b606061011583838080601f0160208091040260200160405190810160405280939291908181526020018383808284375061011c945050505050565b9392505050565b60606000806000806000606060008060008a519450846040519080825280601f01601f19166020018201604052801561015f578160200160208202803883390190505b509350600085118015610173575060048506155b15156101e057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c69642062617365363420656e636f64696e67000000000000000000604482015290519081900360640190fd5b604080517f3d00000000000000000000000000000000000000000000000000000000000000815290519081900360010190208b518c90600119880190811061022457fe5b90602001015160f860020a900460f860020a026040516020018082600160f860020a031916600160f860020a03191681526001019150506040516020818303038152906040526040518082805190602001908083835b602083106102995780518252601f19909201916020918201910161027a565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614156102db576002850394506103d2565b604080517f3d00000000000000000000000000000000000000000000000000000000000000815290519081900360010190208b518c90600019880190811061031f57fe5b90602001015160f860020a900460f860020a026040516020018082600160f860020a031916600160f860020a03191681526001019150506040516020818303038152906040526040518082805190602001908083835b602083106103945780518252601f199092019160209182019101610375565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614156103d2576001850394505b505060048084040260005b81811015610734576040805160a081018252607b815260006020820152600080516020610b6c83398151915291810191909152600080516020610b8c8339815191526060820152600080516020610b4c83398151915260808201528b5160018301928d91811061044957fe5b90602001015160f860020a900460f860020a0260f860020a900481518110151561046f57fe5b90602001015160f860020a900460f860020a02985060a060405190810160405280607b815260200160008152602001600080516020610b6c8339815191528152602001600080516020610b8c8339815191528152602001600080516020610b4c8339815191528152508b828060010193508151811015156104ec57fe5b90602001015160f860020a900460f860020a0260f860020a900481518110151561051257fe5b90602001015160f860020a900460f860020a02975060a060405190810160405280607b815260200160008152602001600080516020610b6c8339815191528152602001600080516020610b8c8339815191528152602001600080516020610b4c8339815191528152508b8280600101935081518110151561058f57fe5b90602001015160f860020a900460f860020a0260f860020a90048151811015156105b557fe5b90602001015160f860020a900460f860020a02965060a060405190810160405280607b815260200160008152602001600080516020610b6c8339815191528152602001600080516020610b8c8339815191528152602001600080516020610b4c8339815191528152508b8280600101935081518110151561063257fe5b90602001015160f860020a900460f860020a0260f860020a900481518110151561065857fe5b0160200151845160f860020a9182900490910296506001840193600160f860020a03198b81166004026010828d1604171691869190811061069557fe5b906020010190600160f860020a031916908160001a905350835160018401936010600160f860020a03198b8116919091026004828c160417169186919081106106da57fe5b906020010190600160f860020a031916908160001a90535083516001840193600160f860020a031989811660400289171691869190811061071757fe5b906020010190600160f860020a031916908160001a9053506103dd565b818503600214156108c6576040805160a081018252607b815260006020820152600080516020610b6c83398151915291810191909152600080516020610b8c8339815191526060820152600080516020610b4c83398151915260808201528b5160018301928d9181106107a357fe5b90602001015160f860020a900460f860020a0260f860020a90048151811015156107c957fe5b90602001015160f860020a900460f860020a02985060a060405190810160405280607b815260200160008152602001600080516020610b6c8339815191528152602001600080516020610b8c8339815191528152602001600080516020610b4c8339815191528152508b8280600101935081518110151561084657fe5b90602001015160f860020a900460f860020a0260f860020a900481518110151561086c57fe5b0160200151845160f860020a9182900490910298506001840193600160f860020a03198b81166004026010828d160417169186919081106108a957fe5b906020010190600160f860020a031916908160001a905350610b3c565b81850360031415610b3c576040805160a081018252607b815260006020820152600080516020610b6c83398151915291810191909152600080516020610b8c8339815191526060820152600080516020610b4c83398151915260808201528b5160018301928d91811061093557fe5b90602001015160f860020a900460f860020a0260f860020a900481518110151561095b57fe5b90602001015160f860020a900460f860020a02985060a060405190810160405280607b815260200160008152602001600080516020610b6c8339815191528152602001600080516020610b8c8339815191528152602001600080516020610b4c8339815191528152508b828060010193508151811015156109d857fe5b90602001015160f860020a900460f860020a0260f860020a90048151811015156109fe57fe5b90602001015160f860020a900460f860020a02975060a060405190810160405280607b815260200160008152602001600080516020610b6c8339815191528152602001600080516020610b8c8339815191528152602001600080516020610b4c8339815191528152508b82806001019350815181101515610a7b57fe5b90602001015160f860020a900460f860020a0260f860020a9004815181101515610aa157fe5b0160200151845160f860020a9182900490910297506001840193600160f860020a03198b81166004026010828d16041716918691908110610ade57fe5b906020010190600160f860020a031916908160001a905350835160018401936010600160f860020a03198b8116919091026004828c16041716918691908110610b2357fe5b906020010190600160f860020a031916908160001a9053505b505081529796505050505050505600001a1b1c1d1e1f202122232425262728292a2b2c2d2e2f30313233000000000000000000000000000000003e003e003f3435363738393a3b3c3d00000000000000000102030405060708090a0b0c0d0e0f10111213141516171819000000003fa165627a7a72305820992c3bf1e90f632ae0ab04e104910e114f9b5a4112da332b40962387f57096e80029`

// DeployBase64Exporter deploys a new Ethereum contract, binding an instance of Base64Exporter to it.
func DeployBase64Exporter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Base64Exporter, error) {
	parsed, err := abi.JSON(strings.NewReader(Base64ExporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Base64ExporterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Base64Exporter{Base64ExporterCaller: Base64ExporterCaller{contract: contract}, Base64ExporterTransactor: Base64ExporterTransactor{contract: contract}, Base64ExporterFilterer: Base64ExporterFilterer{contract: contract}}, nil
}

// Base64Exporter is an auto generated Go binding around an Ethereum contract.
type Base64Exporter struct {
	Base64ExporterCaller     // Read-only binding to the contract
	Base64ExporterTransactor // Write-only binding to the contract
	Base64ExporterFilterer   // Log filterer for contract events
}

// Base64ExporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type Base64ExporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64ExporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Base64ExporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64ExporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Base64ExporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64ExporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Base64ExporterSession struct {
	Contract     *Base64Exporter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Base64ExporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Base64ExporterCallerSession struct {
	Contract *Base64ExporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Base64ExporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Base64ExporterTransactorSession struct {
	Contract     *Base64ExporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Base64ExporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type Base64ExporterRaw struct {
	Contract *Base64Exporter // Generic contract binding to access the raw methods on
}

// Base64ExporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Base64ExporterCallerRaw struct {
	Contract *Base64ExporterCaller // Generic read-only contract binding to access the raw methods on
}

// Base64ExporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Base64ExporterTransactorRaw struct {
	Contract *Base64ExporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBase64Exporter creates a new instance of Base64Exporter, bound to a specific deployed contract.
func NewBase64Exporter(address common.Address, backend bind.ContractBackend) (*Base64Exporter, error) {
	contract, err := bindBase64Exporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Base64Exporter{Base64ExporterCaller: Base64ExporterCaller{contract: contract}, Base64ExporterTransactor: Base64ExporterTransactor{contract: contract}, Base64ExporterFilterer: Base64ExporterFilterer{contract: contract}}, nil
}

// NewBase64ExporterCaller creates a new read-only instance of Base64Exporter, bound to a specific deployed contract.
func NewBase64ExporterCaller(address common.Address, caller bind.ContractCaller) (*Base64ExporterCaller, error) {
	contract, err := bindBase64Exporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Base64ExporterCaller{contract: contract}, nil
}

// NewBase64ExporterTransactor creates a new write-only instance of Base64Exporter, bound to a specific deployed contract.
func NewBase64ExporterTransactor(address common.Address, transactor bind.ContractTransactor) (*Base64ExporterTransactor, error) {
	contract, err := bindBase64Exporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Base64ExporterTransactor{contract: contract}, nil
}

// NewBase64ExporterFilterer creates a new log filterer instance of Base64Exporter, bound to a specific deployed contract.
func NewBase64ExporterFilterer(address common.Address, filterer bind.ContractFilterer) (*Base64ExporterFilterer, error) {
	contract, err := bindBase64Exporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Base64ExporterFilterer{contract: contract}, nil
}

// bindBase64Exporter binds a generic wrapper to an already deployed contract.
func bindBase64Exporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Base64ExporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base64Exporter *Base64ExporterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Base64Exporter.Contract.Base64ExporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base64Exporter *Base64ExporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base64Exporter.Contract.Base64ExporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base64Exporter *Base64ExporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base64Exporter.Contract.Base64ExporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base64Exporter *Base64ExporterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Base64Exporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base64Exporter *Base64ExporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base64Exporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base64Exporter *Base64ExporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base64Exporter.Contract.contract.Transact(opts, method, params...)
}

// Base64decode is a free data retrieval call binding the contract method 0x3a718bfb.
//
// Solidity: function base64decode(bytes _encoded) constant returns(bytes)
func (_Base64Exporter *Base64ExporterCaller) Base64decode(opts *bind.CallOpts, _encoded []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Base64Exporter.contract.Call(opts, out, "base64decode", _encoded)
	return *ret0, err
}

// Base64decode is a free data retrieval call binding the contract method 0x3a718bfb.
//
// Solidity: function base64decode(bytes _encoded) constant returns(bytes)
func (_Base64Exporter *Base64ExporterSession) Base64decode(_encoded []byte) ([]byte, error) {
	return _Base64Exporter.Contract.Base64decode(&_Base64Exporter.CallOpts, _encoded)
}

// Base64decode is a free data retrieval call binding the contract method 0x3a718bfb.
//
// Solidity: function base64decode(bytes _encoded) constant returns(bytes)
func (_Base64Exporter *Base64ExporterCallerSession) Base64decode(_encoded []byte) ([]byte, error) {
	return _Base64Exporter.Contract.Base64decode(&_Base64Exporter.CallOpts, _encoded)
}
