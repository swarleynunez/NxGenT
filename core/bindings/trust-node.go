// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// TypesEvidenceMetric is an auto generated low-level Go binding around an user-defined struct.
type TypesEvidenceMetric struct {
	Id    string
	Value *big.Int
}

// TypesSLAMetric is an auto generated low-level Go binding around an user-defined struct.
type TypesSLAMetric struct {
	Id        string
	Threshold *big.Int
	Weighting *big.Int
}

// TrustNodeMetaData contains all meta data concerning the TrustNode contract.
var TrustNodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AcceptedSLA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"NewSLA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"TerminatedSLA\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SLAs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"setAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"terminatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"acceptSLA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"canAcceptSLA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"mAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"nextSLAId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"registeredAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEvidencesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"slaId\",\"type\":\"uint64\"}],\"name\":\"getLastEvidenceTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slaId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metricId\",\"type\":\"string\"}],\"name\":\"getSLAMetric\",\"outputs\":[{\"internalType\":\"SD59x18\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"SD59x18\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slaId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metricId\",\"type\":\"string\"}],\"name\":\"hasSLAMetric\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"isSLAActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSLACustomer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"lastEvidenceTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_customer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"SD59x18\",\"name\":\"threshold\",\"type\":\"int256\"},{\"internalType\":\"SD59x18\",\"name\":\"weighting\",\"type\":\"int256\"}],\"internalType\":\"structTypes.SLAMetric[]\",\"name\":\"_metrics\",\"type\":\"tuple[]\"}],\"name\":\"setSLA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_slaId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"SD59x18\",\"name\":\"value\",\"type\":\"int256\"}],\"internalType\":\"structTypes.EvidenceMetric[]\",\"name\":\"_metrics\",\"type\":\"tuple[]\"}],\"name\":\"storeEvidence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"slaId\",\"type\":\"uint64\"}],\"name\":\"updateLastEvidenceTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TrustNodeABI is the input ABI used to generate the binding from.
// Deprecated: Use TrustNodeMetaData.ABI instead.
var TrustNodeABI = TrustNodeMetaData.ABI

// TrustNode is an auto generated Go binding around an Ethereum contract.
type TrustNode struct {
	TrustNodeCaller     // Read-only binding to the contract
	TrustNodeTransactor // Write-only binding to the contract
	TrustNodeFilterer   // Log filterer for contract events
}

// TrustNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrustNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrustNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrustNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrustNodeSession struct {
	Contract     *TrustNode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrustNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrustNodeCallerSession struct {
	Contract *TrustNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TrustNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrustNodeTransactorSession struct {
	Contract     *TrustNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TrustNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrustNodeRaw struct {
	Contract *TrustNode // Generic contract binding to access the raw methods on
}

// TrustNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrustNodeCallerRaw struct {
	Contract *TrustNodeCaller // Generic read-only contract binding to access the raw methods on
}

// TrustNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrustNodeTransactorRaw struct {
	Contract *TrustNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrustNode creates a new instance of TrustNode, bound to a specific deployed contract.
func NewTrustNode(address common.Address, backend bind.ContractBackend) (*TrustNode, error) {
	contract, err := bindTrustNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrustNode{TrustNodeCaller: TrustNodeCaller{contract: contract}, TrustNodeTransactor: TrustNodeTransactor{contract: contract}, TrustNodeFilterer: TrustNodeFilterer{contract: contract}}, nil
}

// NewTrustNodeCaller creates a new read-only instance of TrustNode, bound to a specific deployed contract.
func NewTrustNodeCaller(address common.Address, caller bind.ContractCaller) (*TrustNodeCaller, error) {
	contract, err := bindTrustNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrustNodeCaller{contract: contract}, nil
}

// NewTrustNodeTransactor creates a new write-only instance of TrustNode, bound to a specific deployed contract.
func NewTrustNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*TrustNodeTransactor, error) {
	contract, err := bindTrustNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrustNodeTransactor{contract: contract}, nil
}

// NewTrustNodeFilterer creates a new log filterer instance of TrustNode, bound to a specific deployed contract.
func NewTrustNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*TrustNodeFilterer, error) {
	contract, err := bindTrustNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrustNodeFilterer{contract: contract}, nil
}

// bindTrustNode binds a generic wrapper to an already deployed contract.
func bindTrustNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrustNodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustNode *TrustNodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustNode.Contract.TrustNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustNode *TrustNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustNode.Contract.TrustNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustNode *TrustNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustNode.Contract.TrustNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustNode *TrustNodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustNode *TrustNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustNode *TrustNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustNode.Contract.contract.Transact(opts, method, params...)
}

// SLAs is a free data retrieval call binding the contract method 0xe08aff0e.
//
// Solidity: function SLAs(uint64 ) view returns(address customer, uint256 setAt, uint256 acceptedAt, uint256 terminatedAt)
func (_TrustNode *TrustNodeCaller) SLAs(opts *bind.CallOpts, arg0 uint64) (struct {
	Customer     common.Address
	SetAt        *big.Int
	AcceptedAt   *big.Int
	TerminatedAt *big.Int
}, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "SLAs", arg0)

	outstruct := new(struct {
		Customer     common.Address
		SetAt        *big.Int
		AcceptedAt   *big.Int
		TerminatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Customer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SetAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AcceptedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TerminatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SLAs is a free data retrieval call binding the contract method 0xe08aff0e.
//
// Solidity: function SLAs(uint64 ) view returns(address customer, uint256 setAt, uint256 acceptedAt, uint256 terminatedAt)
func (_TrustNode *TrustNodeSession) SLAs(arg0 uint64) (struct {
	Customer     common.Address
	SetAt        *big.Int
	AcceptedAt   *big.Int
	TerminatedAt *big.Int
}, error) {
	return _TrustNode.Contract.SLAs(&_TrustNode.CallOpts, arg0)
}

// SLAs is a free data retrieval call binding the contract method 0xe08aff0e.
//
// Solidity: function SLAs(uint64 ) view returns(address customer, uint256 setAt, uint256 acceptedAt, uint256 terminatedAt)
func (_TrustNode *TrustNodeCallerSession) SLAs(arg0 uint64) (struct {
	Customer     common.Address
	SetAt        *big.Int
	AcceptedAt   *big.Int
	TerminatedAt *big.Int
}, error) {
	return _TrustNode.Contract.SLAs(&_TrustNode.CallOpts, arg0)
}

// CanAcceptSLA is a free data retrieval call binding the contract method 0x43858181.
//
// Solidity: function canAcceptSLA(uint64 id, address addr) view returns(bool)
func (_TrustNode *TrustNodeCaller) CanAcceptSLA(opts *bind.CallOpts, id uint64, addr common.Address) (bool, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "canAcceptSLA", id, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanAcceptSLA is a free data retrieval call binding the contract method 0x43858181.
//
// Solidity: function canAcceptSLA(uint64 id, address addr) view returns(bool)
func (_TrustNode *TrustNodeSession) CanAcceptSLA(id uint64, addr common.Address) (bool, error) {
	return _TrustNode.Contract.CanAcceptSLA(&_TrustNode.CallOpts, id, addr)
}

// CanAcceptSLA is a free data retrieval call binding the contract method 0x43858181.
//
// Solidity: function canAcceptSLA(uint64 id, address addr) view returns(bool)
func (_TrustNode *TrustNodeCallerSession) CanAcceptSLA(id uint64, addr common.Address) (bool, error) {
	return _TrustNode.Contract.CanAcceptSLA(&_TrustNode.CallOpts, id, addr)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address mAddr, address nAddr, string ip, uint64 nextSLAId, uint256 registeredAt)
func (_TrustNode *TrustNodeCaller) Data(opts *bind.CallOpts) (struct {
	MAddr        common.Address
	NAddr        common.Address
	Ip           string
	NextSLAId    uint64
	RegisteredAt *big.Int
}, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "data")

	outstruct := new(struct {
		MAddr        common.Address
		NAddr        common.Address
		Ip           string
		NextSLAId    uint64
		RegisteredAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NAddr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Ip = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.NextSLAId = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.RegisteredAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address mAddr, address nAddr, string ip, uint64 nextSLAId, uint256 registeredAt)
func (_TrustNode *TrustNodeSession) Data() (struct {
	MAddr        common.Address
	NAddr        common.Address
	Ip           string
	NextSLAId    uint64
	RegisteredAt *big.Int
}, error) {
	return _TrustNode.Contract.Data(&_TrustNode.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address mAddr, address nAddr, string ip, uint64 nextSLAId, uint256 registeredAt)
func (_TrustNode *TrustNodeCallerSession) Data() (struct {
	MAddr        common.Address
	NAddr        common.Address
	Ip           string
	NextSLAId    uint64
	RegisteredAt *big.Int
}, error) {
	return _TrustNode.Contract.Data(&_TrustNode.CallOpts)
}

// GetEvidencesCount is a free data retrieval call binding the contract method 0x530745a0.
//
// Solidity: function getEvidencesCount() view returns(uint256)
func (_TrustNode *TrustNodeCaller) GetEvidencesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "getEvidencesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEvidencesCount is a free data retrieval call binding the contract method 0x530745a0.
//
// Solidity: function getEvidencesCount() view returns(uint256)
func (_TrustNode *TrustNodeSession) GetEvidencesCount() (*big.Int, error) {
	return _TrustNode.Contract.GetEvidencesCount(&_TrustNode.CallOpts)
}

// GetEvidencesCount is a free data retrieval call binding the contract method 0x530745a0.
//
// Solidity: function getEvidencesCount() view returns(uint256)
func (_TrustNode *TrustNodeCallerSession) GetEvidencesCount() (*big.Int, error) {
	return _TrustNode.Contract.GetEvidencesCount(&_TrustNode.CallOpts)
}

// GetIP is a free data retrieval call binding the contract method 0x96355b2e.
//
// Solidity: function getIP() view returns(string)
func (_TrustNode *TrustNodeCaller) GetIP(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "getIP")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetIP is a free data retrieval call binding the contract method 0x96355b2e.
//
// Solidity: function getIP() view returns(string)
func (_TrustNode *TrustNodeSession) GetIP() (string, error) {
	return _TrustNode.Contract.GetIP(&_TrustNode.CallOpts)
}

// GetIP is a free data retrieval call binding the contract method 0x96355b2e.
//
// Solidity: function getIP() view returns(string)
func (_TrustNode *TrustNodeCallerSession) GetIP() (string, error) {
	return _TrustNode.Contract.GetIP(&_TrustNode.CallOpts)
}

// GetLastEvidenceTime is a free data retrieval call binding the contract method 0x844ac36c.
//
// Solidity: function getLastEvidenceTime(address target, uint64 slaId) view returns(uint256)
func (_TrustNode *TrustNodeCaller) GetLastEvidenceTime(opts *bind.CallOpts, target common.Address, slaId uint64) (*big.Int, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "getLastEvidenceTime", target, slaId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastEvidenceTime is a free data retrieval call binding the contract method 0x844ac36c.
//
// Solidity: function getLastEvidenceTime(address target, uint64 slaId) view returns(uint256)
func (_TrustNode *TrustNodeSession) GetLastEvidenceTime(target common.Address, slaId uint64) (*big.Int, error) {
	return _TrustNode.Contract.GetLastEvidenceTime(&_TrustNode.CallOpts, target, slaId)
}

// GetLastEvidenceTime is a free data retrieval call binding the contract method 0x844ac36c.
//
// Solidity: function getLastEvidenceTime(address target, uint64 slaId) view returns(uint256)
func (_TrustNode *TrustNodeCallerSession) GetLastEvidenceTime(target common.Address, slaId uint64) (*big.Int, error) {
	return _TrustNode.Contract.GetLastEvidenceTime(&_TrustNode.CallOpts, target, slaId)
}

// GetSLAMetric is a free data retrieval call binding the contract method 0xda9793c0.
//
// Solidity: function getSLAMetric(uint64 slaId, string metricId) view returns(int256, int256)
func (_TrustNode *TrustNodeCaller) GetSLAMetric(opts *bind.CallOpts, slaId uint64, metricId string) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "getSLAMetric", slaId, metricId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetSLAMetric is a free data retrieval call binding the contract method 0xda9793c0.
//
// Solidity: function getSLAMetric(uint64 slaId, string metricId) view returns(int256, int256)
func (_TrustNode *TrustNodeSession) GetSLAMetric(slaId uint64, metricId string) (*big.Int, *big.Int, error) {
	return _TrustNode.Contract.GetSLAMetric(&_TrustNode.CallOpts, slaId, metricId)
}

// GetSLAMetric is a free data retrieval call binding the contract method 0xda9793c0.
//
// Solidity: function getSLAMetric(uint64 slaId, string metricId) view returns(int256, int256)
func (_TrustNode *TrustNodeCallerSession) GetSLAMetric(slaId uint64, metricId string) (*big.Int, *big.Int, error) {
	return _TrustNode.Contract.GetSLAMetric(&_TrustNode.CallOpts, slaId, metricId)
}

// HasSLAMetric is a free data retrieval call binding the contract method 0xf2ad26b4.
//
// Solidity: function hasSLAMetric(uint64 slaId, string metricId) view returns(bool)
func (_TrustNode *TrustNodeCaller) HasSLAMetric(opts *bind.CallOpts, slaId uint64, metricId string) (bool, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "hasSLAMetric", slaId, metricId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSLAMetric is a free data retrieval call binding the contract method 0xf2ad26b4.
//
// Solidity: function hasSLAMetric(uint64 slaId, string metricId) view returns(bool)
func (_TrustNode *TrustNodeSession) HasSLAMetric(slaId uint64, metricId string) (bool, error) {
	return _TrustNode.Contract.HasSLAMetric(&_TrustNode.CallOpts, slaId, metricId)
}

// HasSLAMetric is a free data retrieval call binding the contract method 0xf2ad26b4.
//
// Solidity: function hasSLAMetric(uint64 slaId, string metricId) view returns(bool)
func (_TrustNode *TrustNodeCallerSession) HasSLAMetric(slaId uint64, metricId string) (bool, error) {
	return _TrustNode.Contract.HasSLAMetric(&_TrustNode.CallOpts, slaId, metricId)
}

// IsSLAActive is a free data retrieval call binding the contract method 0x71d2bdaa.
//
// Solidity: function isSLAActive(uint64 id) view returns(bool)
func (_TrustNode *TrustNodeCaller) IsSLAActive(opts *bind.CallOpts, id uint64) (bool, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "isSLAActive", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSLAActive is a free data retrieval call binding the contract method 0x71d2bdaa.
//
// Solidity: function isSLAActive(uint64 id) view returns(bool)
func (_TrustNode *TrustNodeSession) IsSLAActive(id uint64) (bool, error) {
	return _TrustNode.Contract.IsSLAActive(&_TrustNode.CallOpts, id)
}

// IsSLAActive is a free data retrieval call binding the contract method 0x71d2bdaa.
//
// Solidity: function isSLAActive(uint64 id) view returns(bool)
func (_TrustNode *TrustNodeCallerSession) IsSLAActive(id uint64) (bool, error) {
	return _TrustNode.Contract.IsSLAActive(&_TrustNode.CallOpts, id)
}

// IsSLACustomer is a free data retrieval call binding the contract method 0x7fd2eb2d.
//
// Solidity: function isSLACustomer(uint64 id, address addr) view returns(bool)
func (_TrustNode *TrustNodeCaller) IsSLACustomer(opts *bind.CallOpts, id uint64, addr common.Address) (bool, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "isSLACustomer", id, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSLACustomer is a free data retrieval call binding the contract method 0x7fd2eb2d.
//
// Solidity: function isSLACustomer(uint64 id, address addr) view returns(bool)
func (_TrustNode *TrustNodeSession) IsSLACustomer(id uint64, addr common.Address) (bool, error) {
	return _TrustNode.Contract.IsSLACustomer(&_TrustNode.CallOpts, id, addr)
}

// IsSLACustomer is a free data retrieval call binding the contract method 0x7fd2eb2d.
//
// Solidity: function isSLACustomer(uint64 id, address addr) view returns(bool)
func (_TrustNode *TrustNodeCallerSession) IsSLACustomer(id uint64, addr common.Address) (bool, error) {
	return _TrustNode.Contract.IsSLACustomer(&_TrustNode.CallOpts, id, addr)
}

// LastEvidenceTimes is a free data retrieval call binding the contract method 0x0cc9718a.
//
// Solidity: function lastEvidenceTimes(address , uint64 ) view returns(uint256)
func (_TrustNode *TrustNodeCaller) LastEvidenceTimes(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) (*big.Int, error) {
	var out []interface{}
	err := _TrustNode.contract.Call(opts, &out, "lastEvidenceTimes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEvidenceTimes is a free data retrieval call binding the contract method 0x0cc9718a.
//
// Solidity: function lastEvidenceTimes(address , uint64 ) view returns(uint256)
func (_TrustNode *TrustNodeSession) LastEvidenceTimes(arg0 common.Address, arg1 uint64) (*big.Int, error) {
	return _TrustNode.Contract.LastEvidenceTimes(&_TrustNode.CallOpts, arg0, arg1)
}

// LastEvidenceTimes is a free data retrieval call binding the contract method 0x0cc9718a.
//
// Solidity: function lastEvidenceTimes(address , uint64 ) view returns(uint256)
func (_TrustNode *TrustNodeCallerSession) LastEvidenceTimes(arg0 common.Address, arg1 uint64) (*big.Int, error) {
	return _TrustNode.Contract.LastEvidenceTimes(&_TrustNode.CallOpts, arg0, arg1)
}

// AcceptSLA is a paid mutator transaction binding the contract method 0x13012c68.
//
// Solidity: function acceptSLA(uint64 id) returns()
func (_TrustNode *TrustNodeTransactor) AcceptSLA(opts *bind.TransactOpts, id uint64) (*types.Transaction, error) {
	return _TrustNode.contract.Transact(opts, "acceptSLA", id)
}

// AcceptSLA is a paid mutator transaction binding the contract method 0x13012c68.
//
// Solidity: function acceptSLA(uint64 id) returns()
func (_TrustNode *TrustNodeSession) AcceptSLA(id uint64) (*types.Transaction, error) {
	return _TrustNode.Contract.AcceptSLA(&_TrustNode.TransactOpts, id)
}

// AcceptSLA is a paid mutator transaction binding the contract method 0x13012c68.
//
// Solidity: function acceptSLA(uint64 id) returns()
func (_TrustNode *TrustNodeTransactorSession) AcceptSLA(id uint64) (*types.Transaction, error) {
	return _TrustNode.Contract.AcceptSLA(&_TrustNode.TransactOpts, id)
}

// SetSLA is a paid mutator transaction binding the contract method 0x3a0808d2.
//
// Solidity: function setSLA(address _customer, (string,int256,int256)[] _metrics) returns()
func (_TrustNode *TrustNodeTransactor) SetSLA(opts *bind.TransactOpts, _customer common.Address, _metrics []TypesSLAMetric) (*types.Transaction, error) {
	return _TrustNode.contract.Transact(opts, "setSLA", _customer, _metrics)
}

// SetSLA is a paid mutator transaction binding the contract method 0x3a0808d2.
//
// Solidity: function setSLA(address _customer, (string,int256,int256)[] _metrics) returns()
func (_TrustNode *TrustNodeSession) SetSLA(_customer common.Address, _metrics []TypesSLAMetric) (*types.Transaction, error) {
	return _TrustNode.Contract.SetSLA(&_TrustNode.TransactOpts, _customer, _metrics)
}

// SetSLA is a paid mutator transaction binding the contract method 0x3a0808d2.
//
// Solidity: function setSLA(address _customer, (string,int256,int256)[] _metrics) returns()
func (_TrustNode *TrustNodeTransactorSession) SetSLA(_customer common.Address, _metrics []TypesSLAMetric) (*types.Transaction, error) {
	return _TrustNode.Contract.SetSLA(&_TrustNode.TransactOpts, _customer, _metrics)
}

// StoreEvidence is a paid mutator transaction binding the contract method 0x466d8542.
//
// Solidity: function storeEvidence(uint64 _slaId, (string,int256)[] _metrics) returns()
func (_TrustNode *TrustNodeTransactor) StoreEvidence(opts *bind.TransactOpts, _slaId uint64, _metrics []TypesEvidenceMetric) (*types.Transaction, error) {
	return _TrustNode.contract.Transact(opts, "storeEvidence", _slaId, _metrics)
}

// StoreEvidence is a paid mutator transaction binding the contract method 0x466d8542.
//
// Solidity: function storeEvidence(uint64 _slaId, (string,int256)[] _metrics) returns()
func (_TrustNode *TrustNodeSession) StoreEvidence(_slaId uint64, _metrics []TypesEvidenceMetric) (*types.Transaction, error) {
	return _TrustNode.Contract.StoreEvidence(&_TrustNode.TransactOpts, _slaId, _metrics)
}

// StoreEvidence is a paid mutator transaction binding the contract method 0x466d8542.
//
// Solidity: function storeEvidence(uint64 _slaId, (string,int256)[] _metrics) returns()
func (_TrustNode *TrustNodeTransactorSession) StoreEvidence(_slaId uint64, _metrics []TypesEvidenceMetric) (*types.Transaction, error) {
	return _TrustNode.Contract.StoreEvidence(&_TrustNode.TransactOpts, _slaId, _metrics)
}

// UpdateLastEvidenceTime is a paid mutator transaction binding the contract method 0x7bccd032.
//
// Solidity: function updateLastEvidenceTime(address target, uint64 slaId) returns()
func (_TrustNode *TrustNodeTransactor) UpdateLastEvidenceTime(opts *bind.TransactOpts, target common.Address, slaId uint64) (*types.Transaction, error) {
	return _TrustNode.contract.Transact(opts, "updateLastEvidenceTime", target, slaId)
}

// UpdateLastEvidenceTime is a paid mutator transaction binding the contract method 0x7bccd032.
//
// Solidity: function updateLastEvidenceTime(address target, uint64 slaId) returns()
func (_TrustNode *TrustNodeSession) UpdateLastEvidenceTime(target common.Address, slaId uint64) (*types.Transaction, error) {
	return _TrustNode.Contract.UpdateLastEvidenceTime(&_TrustNode.TransactOpts, target, slaId)
}

// UpdateLastEvidenceTime is a paid mutator transaction binding the contract method 0x7bccd032.
//
// Solidity: function updateLastEvidenceTime(address target, uint64 slaId) returns()
func (_TrustNode *TrustNodeTransactorSession) UpdateLastEvidenceTime(target common.Address, slaId uint64) (*types.Transaction, error) {
	return _TrustNode.Contract.UpdateLastEvidenceTime(&_TrustNode.TransactOpts, target, slaId)
}

// TrustNodeAcceptedSLAIterator is returned from FilterAcceptedSLA and is used to iterate over the raw logs and unpacked data for AcceptedSLA events raised by the TrustNode contract.
type TrustNodeAcceptedSLAIterator struct {
	Event *TrustNodeAcceptedSLA // Event containing the contract specifics and raw log

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
func (it *TrustNodeAcceptedSLAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustNodeAcceptedSLA)
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
		it.Event = new(TrustNodeAcceptedSLA)
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
func (it *TrustNodeAcceptedSLAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustNodeAcceptedSLAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustNodeAcceptedSLA represents a AcceptedSLA event raised by the TrustNode contract.
type TrustNodeAcceptedSLA struct {
	From common.Address
	Id   uint64
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAcceptedSLA is a free log retrieval operation binding the contract event 0x00b29cc80a099b3a17f375c363566d3f1feb0e144bc2aac7f88181fc42866c8e.
//
// Solidity: event AcceptedSLA(address from, uint64 id, address to)
func (_TrustNode *TrustNodeFilterer) FilterAcceptedSLA(opts *bind.FilterOpts) (*TrustNodeAcceptedSLAIterator, error) {

	logs, sub, err := _TrustNode.contract.FilterLogs(opts, "AcceptedSLA")
	if err != nil {
		return nil, err
	}
	return &TrustNodeAcceptedSLAIterator{contract: _TrustNode.contract, event: "AcceptedSLA", logs: logs, sub: sub}, nil
}

// WatchAcceptedSLA is a free log subscription operation binding the contract event 0x00b29cc80a099b3a17f375c363566d3f1feb0e144bc2aac7f88181fc42866c8e.
//
// Solidity: event AcceptedSLA(address from, uint64 id, address to)
func (_TrustNode *TrustNodeFilterer) WatchAcceptedSLA(opts *bind.WatchOpts, sink chan<- *TrustNodeAcceptedSLA) (event.Subscription, error) {

	logs, sub, err := _TrustNode.contract.WatchLogs(opts, "AcceptedSLA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustNodeAcceptedSLA)
				if err := _TrustNode.contract.UnpackLog(event, "AcceptedSLA", log); err != nil {
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

// ParseAcceptedSLA is a log parse operation binding the contract event 0x00b29cc80a099b3a17f375c363566d3f1feb0e144bc2aac7f88181fc42866c8e.
//
// Solidity: event AcceptedSLA(address from, uint64 id, address to)
func (_TrustNode *TrustNodeFilterer) ParseAcceptedSLA(log types.Log) (*TrustNodeAcceptedSLA, error) {
	event := new(TrustNodeAcceptedSLA)
	if err := _TrustNode.contract.UnpackLog(event, "AcceptedSLA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustNodeNewSLAIterator is returned from FilterNewSLA and is used to iterate over the raw logs and unpacked data for NewSLA events raised by the TrustNode contract.
type TrustNodeNewSLAIterator struct {
	Event *TrustNodeNewSLA // Event containing the contract specifics and raw log

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
func (it *TrustNodeNewSLAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustNodeNewSLA)
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
		it.Event = new(TrustNodeNewSLA)
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
func (it *TrustNodeNewSLAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustNodeNewSLAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustNodeNewSLA represents a NewSLA event raised by the TrustNode contract.
type TrustNodeNewSLA struct {
	From common.Address
	Id   uint64
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewSLA is a free log retrieval operation binding the contract event 0x24ffc7a120388b6bc1612ba533982e1946b399be13371a0d3e5222e65eb74185.
//
// Solidity: event NewSLA(address from, uint64 id, address to)
func (_TrustNode *TrustNodeFilterer) FilterNewSLA(opts *bind.FilterOpts) (*TrustNodeNewSLAIterator, error) {

	logs, sub, err := _TrustNode.contract.FilterLogs(opts, "NewSLA")
	if err != nil {
		return nil, err
	}
	return &TrustNodeNewSLAIterator{contract: _TrustNode.contract, event: "NewSLA", logs: logs, sub: sub}, nil
}

// WatchNewSLA is a free log subscription operation binding the contract event 0x24ffc7a120388b6bc1612ba533982e1946b399be13371a0d3e5222e65eb74185.
//
// Solidity: event NewSLA(address from, uint64 id, address to)
func (_TrustNode *TrustNodeFilterer) WatchNewSLA(opts *bind.WatchOpts, sink chan<- *TrustNodeNewSLA) (event.Subscription, error) {

	logs, sub, err := _TrustNode.contract.WatchLogs(opts, "NewSLA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustNodeNewSLA)
				if err := _TrustNode.contract.UnpackLog(event, "NewSLA", log); err != nil {
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

// ParseNewSLA is a log parse operation binding the contract event 0x24ffc7a120388b6bc1612ba533982e1946b399be13371a0d3e5222e65eb74185.
//
// Solidity: event NewSLA(address from, uint64 id, address to)
func (_TrustNode *TrustNodeFilterer) ParseNewSLA(log types.Log) (*TrustNodeNewSLA, error) {
	event := new(TrustNodeNewSLA)
	if err := _TrustNode.contract.UnpackLog(event, "NewSLA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustNodeTerminatedSLAIterator is returned from FilterTerminatedSLA and is used to iterate over the raw logs and unpacked data for TerminatedSLA events raised by the TrustNode contract.
type TrustNodeTerminatedSLAIterator struct {
	Event *TrustNodeTerminatedSLA // Event containing the contract specifics and raw log

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
func (it *TrustNodeTerminatedSLAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustNodeTerminatedSLA)
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
		it.Event = new(TrustNodeTerminatedSLA)
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
func (it *TrustNodeTerminatedSLAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustNodeTerminatedSLAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustNodeTerminatedSLA represents a TerminatedSLA event raised by the TrustNode contract.
type TrustNodeTerminatedSLA struct {
	From common.Address
	Id   uint64
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTerminatedSLA is a free log retrieval operation binding the contract event 0x8ba6d1b44b72c8f9f774c453f3d90dbc250adddbd7d91c6ca0a1aa3b47c250fd.
//
// Solidity: event TerminatedSLA(address from, uint64 id, address to)
func (_TrustNode *TrustNodeFilterer) FilterTerminatedSLA(opts *bind.FilterOpts) (*TrustNodeTerminatedSLAIterator, error) {

	logs, sub, err := _TrustNode.contract.FilterLogs(opts, "TerminatedSLA")
	if err != nil {
		return nil, err
	}
	return &TrustNodeTerminatedSLAIterator{contract: _TrustNode.contract, event: "TerminatedSLA", logs: logs, sub: sub}, nil
}

// WatchTerminatedSLA is a free log subscription operation binding the contract event 0x8ba6d1b44b72c8f9f774c453f3d90dbc250adddbd7d91c6ca0a1aa3b47c250fd.
//
// Solidity: event TerminatedSLA(address from, uint64 id, address to)
func (_TrustNode *TrustNodeFilterer) WatchTerminatedSLA(opts *bind.WatchOpts, sink chan<- *TrustNodeTerminatedSLA) (event.Subscription, error) {

	logs, sub, err := _TrustNode.contract.WatchLogs(opts, "TerminatedSLA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustNodeTerminatedSLA)
				if err := _TrustNode.contract.UnpackLog(event, "TerminatedSLA", log); err != nil {
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

// ParseTerminatedSLA is a log parse operation binding the contract event 0x8ba6d1b44b72c8f9f774c453f3d90dbc250adddbd7d91c6ca0a1aa3b47c250fd.
//
// Solidity: event TerminatedSLA(address from, uint64 id, address to)
func (_TrustNode *TrustNodeFilterer) ParseTerminatedSLA(log types.Log) (*TrustNodeTerminatedSLA, error) {
	event := new(TrustNodeTerminatedSLA)
	if err := _TrustNode.contract.UnpackLog(event, "TerminatedSLA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
