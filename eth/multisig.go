// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// MultiSigWithRoleTransaction is an auto generated low-level Go binding around an user-defined struct.
type MultiSigWithRoleTransaction struct {
	Submitter     common.Address
	Target        common.Address
	Data          []byte
	Note          string
	Deadline      *big.Int
	Confirmations *big.Int
	Status        uint8
}

// MultisigMetaData contains all meta data concerning the Multisig contract.
var MultisigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"AddSubmitter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"CancelTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"ConfirmTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"ExecuteTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"RevokeOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"RevokeSubmitter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"RevokeTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_note\",\"type\":\"string\"}],\"name\":\"SubmitTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newWeight\",\"type\":\"uint256\"}],\"name\":\"UpdateWeight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"enumMultiSigWithRole.Role\",\"name\":\"_role\",\"type\":\"uint8\"}],\"name\":\"addRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"cancelTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTransactionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getCancelTxByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"},{\"internalType\":\"enumMultiSigWithRole.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structMultiSigWithRole.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getExecutedTxByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"},{\"internalType\":\"enumMultiSigWithRole.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structMultiSigWithRole.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOwnerByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getPendingTxByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"},{\"internalType\":\"enumMultiSigWithRole.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structMultiSigWithRole.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getSubmitterByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isSubmitter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"multiSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"multiSendTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"enumMultiSigWithRole.Role\",\"name\":\"_role\",\"type\":\"uint8\"}],\"name\":\"removeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"revokeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_note\",\"type\":\"string\"}],\"name\":\"submitTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCancelTxs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalExecutedTxs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOwner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPendingTxs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSubmitter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"},{\"internalType\":\"enumMultiSigWithRole.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"updateWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MultisigABI is the input ABI used to generate the binding from.
// Deprecated: Use MultisigMetaData.ABI instead.
var MultisigABI = MultisigMetaData.ABI

// Multisig is an auto generated Go binding around an Ethereum contract.
type Multisig struct {
	MultisigCaller     // Read-only binding to the contract
	MultisigTransactor // Write-only binding to the contract
	MultisigFilterer   // Log filterer for contract events
}

// MultisigCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisigSession struct {
	Contract     *Multisig         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisigCallerSession struct {
	Contract *MultisigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MultisigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisigTransactorSession struct {
	Contract     *MultisigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MultisigRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisigRaw struct {
	Contract *Multisig // Generic contract binding to access the raw methods on
}

// MultisigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisigCallerRaw struct {
	Contract *MultisigCaller // Generic read-only contract binding to access the raw methods on
}

// MultisigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisigTransactorRaw struct {
	Contract *MultisigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisig creates a new instance of Multisig, bound to a specific deployed contract.
func NewMultisig(address common.Address, backend bind.ContractBackend) (*Multisig, error) {
	contract, err := bindMultisig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multisig{MultisigCaller: MultisigCaller{contract: contract}, MultisigTransactor: MultisigTransactor{contract: contract}, MultisigFilterer: MultisigFilterer{contract: contract}}, nil
}

// NewMultisigCaller creates a new read-only instance of Multisig, bound to a specific deployed contract.
func NewMultisigCaller(address common.Address, caller bind.ContractCaller) (*MultisigCaller, error) {
	contract, err := bindMultisig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigCaller{contract: contract}, nil
}

// NewMultisigTransactor creates a new write-only instance of Multisig, bound to a specific deployed contract.
func NewMultisigTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisigTransactor, error) {
	contract, err := bindMultisig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigTransactor{contract: contract}, nil
}

// NewMultisigFilterer creates a new log filterer instance of Multisig, bound to a specific deployed contract.
func NewMultisigFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisigFilterer, error) {
	contract, err := bindMultisig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisigFilterer{contract: contract}, nil
}

// bindMultisig binds a generic wrapper to an already deployed contract.
func bindMultisig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisig *MultisigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisig.Contract.MultisigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisig *MultisigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisig.Contract.MultisigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisig *MultisigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisig.Contract.MultisigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisig *MultisigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisig *MultisigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisig *MultisigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisig.Contract.contract.Transact(opts, method, params...)
}

// CurrentTransactionId is a free data retrieval call binding the contract method 0x52b422b7.
//
// Solidity: function currentTransactionId() view returns(uint256)
func (_Multisig *MultisigCaller) CurrentTransactionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "currentTransactionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTransactionId is a free data retrieval call binding the contract method 0x52b422b7.
//
// Solidity: function currentTransactionId() view returns(uint256)
func (_Multisig *MultisigSession) CurrentTransactionId() (*big.Int, error) {
	return _Multisig.Contract.CurrentTransactionId(&_Multisig.CallOpts)
}

// CurrentTransactionId is a free data retrieval call binding the contract method 0x52b422b7.
//
// Solidity: function currentTransactionId() view returns(uint256)
func (_Multisig *MultisigCallerSession) CurrentTransactionId() (*big.Int, error) {
	return _Multisig.Contract.CurrentTransactionId(&_Multisig.CallOpts)
}

// GetCancelTxByIndex is a free data retrieval call binding the contract method 0xeab237bd.
//
// Solidity: function getCancelTxByIndex(uint256 _index) view returns(uint256 txId, (address,address,bytes,string,uint256,uint256,uint8))
func (_Multisig *MultisigCaller) GetCancelTxByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, MultiSigWithRoleTransaction, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getCancelTxByIndex", _index)

	if err != nil {
		return *new(*big.Int), *new(MultiSigWithRoleTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(MultiSigWithRoleTransaction)).(*MultiSigWithRoleTransaction)

	return out0, out1, err

}

// GetCancelTxByIndex is a free data retrieval call binding the contract method 0xeab237bd.
//
// Solidity: function getCancelTxByIndex(uint256 _index) view returns(uint256 txId, (address,address,bytes,string,uint256,uint256,uint8))
func (_Multisig *MultisigSession) GetCancelTxByIndex(_index *big.Int) (*big.Int, MultiSigWithRoleTransaction, error) {
	return _Multisig.Contract.GetCancelTxByIndex(&_Multisig.CallOpts, _index)
}

// GetCancelTxByIndex is a free data retrieval call binding the contract method 0xeab237bd.
//
// Solidity: function getCancelTxByIndex(uint256 _index) view returns(uint256 txId, (address,address,bytes,string,uint256,uint256,uint8))
func (_Multisig *MultisigCallerSession) GetCancelTxByIndex(_index *big.Int) (*big.Int, MultiSigWithRoleTransaction, error) {
	return _Multisig.Contract.GetCancelTxByIndex(&_Multisig.CallOpts, _index)
}

// GetExecutedTxByIndex is a free data retrieval call binding the contract method 0x7f00d91f.
//
// Solidity: function getExecutedTxByIndex(uint256 _index) view returns(uint256 txId, (address,address,bytes,string,uint256,uint256,uint8))
func (_Multisig *MultisigCaller) GetExecutedTxByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, MultiSigWithRoleTransaction, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getExecutedTxByIndex", _index)

	if err != nil {
		return *new(*big.Int), *new(MultiSigWithRoleTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(MultiSigWithRoleTransaction)).(*MultiSigWithRoleTransaction)

	return out0, out1, err

}

// GetExecutedTxByIndex is a free data retrieval call binding the contract method 0x7f00d91f.
//
// Solidity: function getExecutedTxByIndex(uint256 _index) view returns(uint256 txId, (address,address,bytes,string,uint256,uint256,uint8))
func (_Multisig *MultisigSession) GetExecutedTxByIndex(_index *big.Int) (*big.Int, MultiSigWithRoleTransaction, error) {
	return _Multisig.Contract.GetExecutedTxByIndex(&_Multisig.CallOpts, _index)
}

// GetExecutedTxByIndex is a free data retrieval call binding the contract method 0x7f00d91f.
//
// Solidity: function getExecutedTxByIndex(uint256 _index) view returns(uint256 txId, (address,address,bytes,string,uint256,uint256,uint8))
func (_Multisig *MultisigCallerSession) GetExecutedTxByIndex(_index *big.Int) (*big.Int, MultiSigWithRoleTransaction, error) {
	return _Multisig.Contract.GetExecutedTxByIndex(&_Multisig.CallOpts, _index)
}

// GetOwnerByIndex is a free data retrieval call binding the contract method 0x1b732d46.
//
// Solidity: function getOwnerByIndex(uint256 _index) view returns(address)
func (_Multisig *MultisigCaller) GetOwnerByIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getOwnerByIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerByIndex is a free data retrieval call binding the contract method 0x1b732d46.
//
// Solidity: function getOwnerByIndex(uint256 _index) view returns(address)
func (_Multisig *MultisigSession) GetOwnerByIndex(_index *big.Int) (common.Address, error) {
	return _Multisig.Contract.GetOwnerByIndex(&_Multisig.CallOpts, _index)
}

// GetOwnerByIndex is a free data retrieval call binding the contract method 0x1b732d46.
//
// Solidity: function getOwnerByIndex(uint256 _index) view returns(address)
func (_Multisig *MultisigCallerSession) GetOwnerByIndex(_index *big.Int) (common.Address, error) {
	return _Multisig.Contract.GetOwnerByIndex(&_Multisig.CallOpts, _index)
}

// GetPendingTxByIndex is a free data retrieval call binding the contract method 0xbfdc77cb.
//
// Solidity: function getPendingTxByIndex(uint256 _index) view returns(uint256 txId, (address,address,bytes,string,uint256,uint256,uint8))
func (_Multisig *MultisigCaller) GetPendingTxByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, MultiSigWithRoleTransaction, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getPendingTxByIndex", _index)

	if err != nil {
		return *new(*big.Int), *new(MultiSigWithRoleTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(MultiSigWithRoleTransaction)).(*MultiSigWithRoleTransaction)

	return out0, out1, err

}

// GetPendingTxByIndex is a free data retrieval call binding the contract method 0xbfdc77cb.
//
// Solidity: function getPendingTxByIndex(uint256 _index) view returns(uint256 txId, (address,address,bytes,string,uint256,uint256,uint8))
func (_Multisig *MultisigSession) GetPendingTxByIndex(_index *big.Int) (*big.Int, MultiSigWithRoleTransaction, error) {
	return _Multisig.Contract.GetPendingTxByIndex(&_Multisig.CallOpts, _index)
}

// GetPendingTxByIndex is a free data retrieval call binding the contract method 0xbfdc77cb.
//
// Solidity: function getPendingTxByIndex(uint256 _index) view returns(uint256 txId, (address,address,bytes,string,uint256,uint256,uint8))
func (_Multisig *MultisigCallerSession) GetPendingTxByIndex(_index *big.Int) (*big.Int, MultiSigWithRoleTransaction, error) {
	return _Multisig.Contract.GetPendingTxByIndex(&_Multisig.CallOpts, _index)
}

// GetSubmitterByIndex is a free data retrieval call binding the contract method 0x886beaef.
//
// Solidity: function getSubmitterByIndex(uint256 _index) view returns(address)
func (_Multisig *MultisigCaller) GetSubmitterByIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getSubmitterByIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubmitterByIndex is a free data retrieval call binding the contract method 0x886beaef.
//
// Solidity: function getSubmitterByIndex(uint256 _index) view returns(address)
func (_Multisig *MultisigSession) GetSubmitterByIndex(_index *big.Int) (common.Address, error) {
	return _Multisig.Contract.GetSubmitterByIndex(&_Multisig.CallOpts, _index)
}

// GetSubmitterByIndex is a free data retrieval call binding the contract method 0x886beaef.
//
// Solidity: function getSubmitterByIndex(uint256 _index) view returns(address)
func (_Multisig *MultisigCallerSession) GetSubmitterByIndex(_index *big.Int) (common.Address, error) {
	return _Multisig.Contract.GetSubmitterByIndex(&_Multisig.CallOpts, _index)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 _id, address _user) view returns(bool)
func (_Multisig *MultisigCaller) IsConfirmed(opts *bind.CallOpts, _id *big.Int, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "isConfirmed", _id, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 _id, address _user) view returns(bool)
func (_Multisig *MultisigSession) IsConfirmed(_id *big.Int, _user common.Address) (bool, error) {
	return _Multisig.Contract.IsConfirmed(&_Multisig.CallOpts, _id, _user)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 _id, address _user) view returns(bool)
func (_Multisig *MultisigCallerSession) IsConfirmed(_id *big.Int, _user common.Address) (bool, error) {
	return _Multisig.Contract.IsConfirmed(&_Multisig.CallOpts, _id, _user)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _user) view returns(bool)
func (_Multisig *MultisigCaller) IsOwner(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "isOwner", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _user) view returns(bool)
func (_Multisig *MultisigSession) IsOwner(_user common.Address) (bool, error) {
	return _Multisig.Contract.IsOwner(&_Multisig.CallOpts, _user)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _user) view returns(bool)
func (_Multisig *MultisigCallerSession) IsOwner(_user common.Address) (bool, error) {
	return _Multisig.Contract.IsOwner(&_Multisig.CallOpts, _user)
}

// IsSubmitter is a free data retrieval call binding the contract method 0xa926fdbc.
//
// Solidity: function isSubmitter(address _user) view returns(bool)
func (_Multisig *MultisigCaller) IsSubmitter(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "isSubmitter", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubmitter is a free data retrieval call binding the contract method 0xa926fdbc.
//
// Solidity: function isSubmitter(address _user) view returns(bool)
func (_Multisig *MultisigSession) IsSubmitter(_user common.Address) (bool, error) {
	return _Multisig.Contract.IsSubmitter(&_Multisig.CallOpts, _user)
}

// IsSubmitter is a free data retrieval call binding the contract method 0xa926fdbc.
//
// Solidity: function isSubmitter(address _user) view returns(bool)
func (_Multisig *MultisigCallerSession) IsSubmitter(_user common.Address) (bool, error) {
	return _Multisig.Contract.IsSubmitter(&_Multisig.CallOpts, _user)
}

// TotalCancelTxs is a free data retrieval call binding the contract method 0xc1af8660.
//
// Solidity: function totalCancelTxs() view returns(uint256)
func (_Multisig *MultisigCaller) TotalCancelTxs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "totalCancelTxs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCancelTxs is a free data retrieval call binding the contract method 0xc1af8660.
//
// Solidity: function totalCancelTxs() view returns(uint256)
func (_Multisig *MultisigSession) TotalCancelTxs() (*big.Int, error) {
	return _Multisig.Contract.TotalCancelTxs(&_Multisig.CallOpts)
}

// TotalCancelTxs is a free data retrieval call binding the contract method 0xc1af8660.
//
// Solidity: function totalCancelTxs() view returns(uint256)
func (_Multisig *MultisigCallerSession) TotalCancelTxs() (*big.Int, error) {
	return _Multisig.Contract.TotalCancelTxs(&_Multisig.CallOpts)
}

// TotalExecutedTxs is a free data retrieval call binding the contract method 0x5fe6d8f3.
//
// Solidity: function totalExecutedTxs() view returns(uint256)
func (_Multisig *MultisigCaller) TotalExecutedTxs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "totalExecutedTxs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalExecutedTxs is a free data retrieval call binding the contract method 0x5fe6d8f3.
//
// Solidity: function totalExecutedTxs() view returns(uint256)
func (_Multisig *MultisigSession) TotalExecutedTxs() (*big.Int, error) {
	return _Multisig.Contract.TotalExecutedTxs(&_Multisig.CallOpts)
}

// TotalExecutedTxs is a free data retrieval call binding the contract method 0x5fe6d8f3.
//
// Solidity: function totalExecutedTxs() view returns(uint256)
func (_Multisig *MultisigCallerSession) TotalExecutedTxs() (*big.Int, error) {
	return _Multisig.Contract.TotalExecutedTxs(&_Multisig.CallOpts)
}

// TotalOwner is a free data retrieval call binding the contract method 0x420b538c.
//
// Solidity: function totalOwner() view returns(uint256)
func (_Multisig *MultisigCaller) TotalOwner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "totalOwner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOwner is a free data retrieval call binding the contract method 0x420b538c.
//
// Solidity: function totalOwner() view returns(uint256)
func (_Multisig *MultisigSession) TotalOwner() (*big.Int, error) {
	return _Multisig.Contract.TotalOwner(&_Multisig.CallOpts)
}

// TotalOwner is a free data retrieval call binding the contract method 0x420b538c.
//
// Solidity: function totalOwner() view returns(uint256)
func (_Multisig *MultisigCallerSession) TotalOwner() (*big.Int, error) {
	return _Multisig.Contract.TotalOwner(&_Multisig.CallOpts)
}

// TotalPendingTxs is a free data retrieval call binding the contract method 0xfda58c11.
//
// Solidity: function totalPendingTxs() view returns(uint256)
func (_Multisig *MultisigCaller) TotalPendingTxs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "totalPendingTxs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPendingTxs is a free data retrieval call binding the contract method 0xfda58c11.
//
// Solidity: function totalPendingTxs() view returns(uint256)
func (_Multisig *MultisigSession) TotalPendingTxs() (*big.Int, error) {
	return _Multisig.Contract.TotalPendingTxs(&_Multisig.CallOpts)
}

// TotalPendingTxs is a free data retrieval call binding the contract method 0xfda58c11.
//
// Solidity: function totalPendingTxs() view returns(uint256)
func (_Multisig *MultisigCallerSession) TotalPendingTxs() (*big.Int, error) {
	return _Multisig.Contract.TotalPendingTxs(&_Multisig.CallOpts)
}

// TotalSubmitter is a free data retrieval call binding the contract method 0x3ec82cb7.
//
// Solidity: function totalSubmitter() view returns(uint256)
func (_Multisig *MultisigCaller) TotalSubmitter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "totalSubmitter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSubmitter is a free data retrieval call binding the contract method 0x3ec82cb7.
//
// Solidity: function totalSubmitter() view returns(uint256)
func (_Multisig *MultisigSession) TotalSubmitter() (*big.Int, error) {
	return _Multisig.Contract.TotalSubmitter(&_Multisig.CallOpts)
}

// TotalSubmitter is a free data retrieval call binding the contract method 0x3ec82cb7.
//
// Solidity: function totalSubmitter() view returns(uint256)
func (_Multisig *MultisigCallerSession) TotalSubmitter() (*big.Int, error) {
	return _Multisig.Contract.TotalSubmitter(&_Multisig.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address submitter, address target, bytes data, string note, uint256 deadline, uint256 confirmations, uint8 status)
func (_Multisig *MultisigCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Submitter     common.Address
	Target        common.Address
	Data          []byte
	Note          string
	Deadline      *big.Int
	Confirmations *big.Int
	Status        uint8
}, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		Submitter     common.Address
		Target        common.Address
		Data          []byte
		Note          string
		Deadline      *big.Int
		Confirmations *big.Int
		Status        uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Submitter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Target = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Note = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Deadline = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Confirmations = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address submitter, address target, bytes data, string note, uint256 deadline, uint256 confirmations, uint8 status)
func (_Multisig *MultisigSession) Transactions(arg0 *big.Int) (struct {
	Submitter     common.Address
	Target        common.Address
	Data          []byte
	Note          string
	Deadline      *big.Int
	Confirmations *big.Int
	Status        uint8
}, error) {
	return _Multisig.Contract.Transactions(&_Multisig.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address submitter, address target, bytes data, string note, uint256 deadline, uint256 confirmations, uint8 status)
func (_Multisig *MultisigCallerSession) Transactions(arg0 *big.Int) (struct {
	Submitter     common.Address
	Target        common.Address
	Data          []byte
	Note          string
	Deadline      *big.Int
	Confirmations *big.Int
	Status        uint8
}, error) {
	return _Multisig.Contract.Transactions(&_Multisig.CallOpts, arg0)
}

// Weight is a free data retrieval call binding the contract method 0xa1aab33f.
//
// Solidity: function weight() view returns(uint256)
func (_Multisig *MultisigCaller) Weight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "weight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Weight is a free data retrieval call binding the contract method 0xa1aab33f.
//
// Solidity: function weight() view returns(uint256)
func (_Multisig *MultisigSession) Weight() (*big.Int, error) {
	return _Multisig.Contract.Weight(&_Multisig.CallOpts)
}

// Weight is a free data retrieval call binding the contract method 0xa1aab33f.
//
// Solidity: function weight() view returns(uint256)
func (_Multisig *MultisigCallerSession) Weight() (*big.Int, error) {
	return _Multisig.Contract.Weight(&_Multisig.CallOpts)
}

// AddRole is a paid mutator transaction binding the contract method 0x44deb6f3.
//
// Solidity: function addRole(address _account, uint8 _role) returns()
func (_Multisig *MultisigTransactor) AddRole(opts *bind.TransactOpts, _account common.Address, _role uint8) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "addRole", _account, _role)
}

// AddRole is a paid mutator transaction binding the contract method 0x44deb6f3.
//
// Solidity: function addRole(address _account, uint8 _role) returns()
func (_Multisig *MultisigSession) AddRole(_account common.Address, _role uint8) (*types.Transaction, error) {
	return _Multisig.Contract.AddRole(&_Multisig.TransactOpts, _account, _role)
}

// AddRole is a paid mutator transaction binding the contract method 0x44deb6f3.
//
// Solidity: function addRole(address _account, uint8 _role) returns()
func (_Multisig *MultisigTransactorSession) AddRole(_account common.Address, _role uint8) (*types.Transaction, error) {
	return _Multisig.Contract.AddRole(&_Multisig.TransactOpts, _account, _role)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0x3380c0d8.
//
// Solidity: function cancelTransaction(uint256 _id) returns()
func (_Multisig *MultisigTransactor) CancelTransaction(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "cancelTransaction", _id)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0x3380c0d8.
//
// Solidity: function cancelTransaction(uint256 _id) returns()
func (_Multisig *MultisigSession) CancelTransaction(_id *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.CancelTransaction(&_Multisig.TransactOpts, _id)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0x3380c0d8.
//
// Solidity: function cancelTransaction(uint256 _id) returns()
func (_Multisig *MultisigTransactorSession) CancelTransaction(_id *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.CancelTransaction(&_Multisig.TransactOpts, _id)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _id) returns()
func (_Multisig *MultisigTransactor) ConfirmTransaction(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "confirmTransaction", _id)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _id) returns()
func (_Multisig *MultisigSession) ConfirmTransaction(_id *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ConfirmTransaction(&_Multisig.TransactOpts, _id)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _id) returns()
func (_Multisig *MultisigTransactorSession) ConfirmTransaction(_id *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ConfirmTransaction(&_Multisig.TransactOpts, _id)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _id) payable returns()
func (_Multisig *MultisigTransactor) ExecuteTransaction(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "executeTransaction", _id)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _id) payable returns()
func (_Multisig *MultisigSession) ExecuteTransaction(_id *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ExecuteTransaction(&_Multisig.TransactOpts, _id)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _id) payable returns()
func (_Multisig *MultisigTransactorSession) ExecuteTransaction(_id *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ExecuteTransaction(&_Multisig.TransactOpts, _id)
}

// MultiSend is a paid mutator transaction binding the contract method 0x9ec68f0f.
//
// Solidity: function multiSend(address _token, address[] _receivers, uint256[] _amounts) returns()
func (_Multisig *MultisigTransactor) MultiSend(opts *bind.TransactOpts, _token common.Address, _receivers []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "multiSend", _token, _receivers, _amounts)
}

// MultiSend is a paid mutator transaction binding the contract method 0x9ec68f0f.
//
// Solidity: function multiSend(address _token, address[] _receivers, uint256[] _amounts) returns()
func (_Multisig *MultisigSession) MultiSend(_token common.Address, _receivers []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.MultiSend(&_Multisig.TransactOpts, _token, _receivers, _amounts)
}

// MultiSend is a paid mutator transaction binding the contract method 0x9ec68f0f.
//
// Solidity: function multiSend(address _token, address[] _receivers, uint256[] _amounts) returns()
func (_Multisig *MultisigTransactorSession) MultiSend(_token common.Address, _receivers []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.MultiSend(&_Multisig.TransactOpts, _token, _receivers, _amounts)
}

// MultiSendTokens is a paid mutator transaction binding the contract method 0x76cf82eb.
//
// Solidity: function multiSendTokens(address[] _tokens, address[] _receivers, uint256[] _amounts) returns()
func (_Multisig *MultisigTransactor) MultiSendTokens(opts *bind.TransactOpts, _tokens []common.Address, _receivers []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "multiSendTokens", _tokens, _receivers, _amounts)
}

// MultiSendTokens is a paid mutator transaction binding the contract method 0x76cf82eb.
//
// Solidity: function multiSendTokens(address[] _tokens, address[] _receivers, uint256[] _amounts) returns()
func (_Multisig *MultisigSession) MultiSendTokens(_tokens []common.Address, _receivers []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.MultiSendTokens(&_Multisig.TransactOpts, _tokens, _receivers, _amounts)
}

// MultiSendTokens is a paid mutator transaction binding the contract method 0x76cf82eb.
//
// Solidity: function multiSendTokens(address[] _tokens, address[] _receivers, uint256[] _amounts) returns()
func (_Multisig *MultisigTransactorSession) MultiSendTokens(_tokens []common.Address, _receivers []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.MultiSendTokens(&_Multisig.TransactOpts, _tokens, _receivers, _amounts)
}

// Payment is a paid mutator transaction binding the contract method 0x99b23dc3.
//
// Solidity: function payment(address _token, address _receiver, uint256 _amount) returns()
func (_Multisig *MultisigTransactor) Payment(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "payment", _token, _receiver, _amount)
}

// Payment is a paid mutator transaction binding the contract method 0x99b23dc3.
//
// Solidity: function payment(address _token, address _receiver, uint256 _amount) returns()
func (_Multisig *MultisigSession) Payment(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.Payment(&_Multisig.TransactOpts, _token, _receiver, _amount)
}

// Payment is a paid mutator transaction binding the contract method 0x99b23dc3.
//
// Solidity: function payment(address _token, address _receiver, uint256 _amount) returns()
func (_Multisig *MultisigTransactorSession) Payment(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.Payment(&_Multisig.TransactOpts, _token, _receiver, _amount)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xb74866fb.
//
// Solidity: function removeRole(address _account, uint8 _role) returns()
func (_Multisig *MultisigTransactor) RemoveRole(opts *bind.TransactOpts, _account common.Address, _role uint8) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "removeRole", _account, _role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xb74866fb.
//
// Solidity: function removeRole(address _account, uint8 _role) returns()
func (_Multisig *MultisigSession) RemoveRole(_account common.Address, _role uint8) (*types.Transaction, error) {
	return _Multisig.Contract.RemoveRole(&_Multisig.TransactOpts, _account, _role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xb74866fb.
//
// Solidity: function removeRole(address _account, uint8 _role) returns()
func (_Multisig *MultisigTransactorSession) RemoveRole(_account common.Address, _role uint8) (*types.Transaction, error) {
	return _Multisig.Contract.RemoveRole(&_Multisig.TransactOpts, _account, _role)
}

// RevokeTransaction is a paid mutator transaction binding the contract method 0x735631ad.
//
// Solidity: function revokeTransaction(uint256 _id) returns()
func (_Multisig *MultisigTransactor) RevokeTransaction(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "revokeTransaction", _id)
}

// RevokeTransaction is a paid mutator transaction binding the contract method 0x735631ad.
//
// Solidity: function revokeTransaction(uint256 _id) returns()
func (_Multisig *MultisigSession) RevokeTransaction(_id *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.RevokeTransaction(&_Multisig.TransactOpts, _id)
}

// RevokeTransaction is a paid mutator transaction binding the contract method 0x735631ad.
//
// Solidity: function revokeTransaction(uint256 _id) returns()
func (_Multisig *MultisigTransactorSession) RevokeTransaction(_id *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.RevokeTransaction(&_Multisig.TransactOpts, _id)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xbe616434.
//
// Solidity: function submitTransaction(address _target, bytes _data, uint256 _deadline, string _note) returns()
func (_Multisig *MultisigTransactor) SubmitTransaction(opts *bind.TransactOpts, _target common.Address, _data []byte, _deadline *big.Int, _note string) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "submitTransaction", _target, _data, _deadline, _note)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xbe616434.
//
// Solidity: function submitTransaction(address _target, bytes _data, uint256 _deadline, string _note) returns()
func (_Multisig *MultisigSession) SubmitTransaction(_target common.Address, _data []byte, _deadline *big.Int, _note string) (*types.Transaction, error) {
	return _Multisig.Contract.SubmitTransaction(&_Multisig.TransactOpts, _target, _data, _deadline, _note)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xbe616434.
//
// Solidity: function submitTransaction(address _target, bytes _data, uint256 _deadline, string _note) returns()
func (_Multisig *MultisigTransactorSession) SubmitTransaction(_target common.Address, _data []byte, _deadline *big.Int, _note string) (*types.Transaction, error) {
	return _Multisig.Contract.SubmitTransaction(&_Multisig.TransactOpts, _target, _data, _deadline, _note)
}

// UpdateWeight is a paid mutator transaction binding the contract method 0xec3df6ad.
//
// Solidity: function updateWeight(uint256 _weight) returns()
func (_Multisig *MultisigTransactor) UpdateWeight(opts *bind.TransactOpts, _weight *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "updateWeight", _weight)
}

// UpdateWeight is a paid mutator transaction binding the contract method 0xec3df6ad.
//
// Solidity: function updateWeight(uint256 _weight) returns()
func (_Multisig *MultisigSession) UpdateWeight(_weight *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.UpdateWeight(&_Multisig.TransactOpts, _weight)
}

// UpdateWeight is a paid mutator transaction binding the contract method 0xec3df6ad.
//
// Solidity: function updateWeight(uint256 _weight) returns()
func (_Multisig *MultisigTransactorSession) UpdateWeight(_weight *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.UpdateWeight(&_Multisig.TransactOpts, _weight)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisig *MultisigTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisig.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisig *MultisigSession) Receive() (*types.Transaction, error) {
	return _Multisig.Contract.Receive(&_Multisig.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisig *MultisigTransactorSession) Receive() (*types.Transaction, error) {
	return _Multisig.Contract.Receive(&_Multisig.TransactOpts)
}

// MultisigAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Multisig contract.
type MultisigAddOwnerIterator struct {
	Event *MultisigAddOwner // Event containing the contract specifics and raw log

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
func (it *MultisigAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigAddOwner)
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
		it.Event = new(MultisigAddOwner)
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
func (it *MultisigAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigAddOwner represents a AddOwner event raised by the Multisig contract.
type MultisigAddOwner struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0xac1e9ef41b54c676ccf449d83ae6f2624bcdce8f5b93a6b48ce95874c332693d.
//
// Solidity: event AddOwner(address _account)
func (_Multisig *MultisigFilterer) FilterAddOwner(opts *bind.FilterOpts) (*MultisigAddOwnerIterator, error) {

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &MultisigAddOwnerIterator{contract: _Multisig.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0xac1e9ef41b54c676ccf449d83ae6f2624bcdce8f5b93a6b48ce95874c332693d.
//
// Solidity: event AddOwner(address _account)
func (_Multisig *MultisigFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *MultisigAddOwner) (event.Subscription, error) {

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigAddOwner)
				if err := _Multisig.contract.UnpackLog(event, "AddOwner", log); err != nil {
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

// ParseAddOwner is a log parse operation binding the contract event 0xac1e9ef41b54c676ccf449d83ae6f2624bcdce8f5b93a6b48ce95874c332693d.
//
// Solidity: event AddOwner(address _account)
func (_Multisig *MultisigFilterer) ParseAddOwner(log types.Log) (*MultisigAddOwner, error) {
	event := new(MultisigAddOwner)
	if err := _Multisig.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigAddSubmitterIterator is returned from FilterAddSubmitter and is used to iterate over the raw logs and unpacked data for AddSubmitter events raised by the Multisig contract.
type MultisigAddSubmitterIterator struct {
	Event *MultisigAddSubmitter // Event containing the contract specifics and raw log

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
func (it *MultisigAddSubmitterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigAddSubmitter)
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
		it.Event = new(MultisigAddSubmitter)
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
func (it *MultisigAddSubmitterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigAddSubmitterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigAddSubmitter represents a AddSubmitter event raised by the Multisig contract.
type MultisigAddSubmitter struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddSubmitter is a free log retrieval operation binding the contract event 0x12fed2eaa2dbcff8567935f8d8c3eb8917ddd2b0466d2c17623503b8e64df06e.
//
// Solidity: event AddSubmitter(address _account)
func (_Multisig *MultisigFilterer) FilterAddSubmitter(opts *bind.FilterOpts) (*MultisigAddSubmitterIterator, error) {

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "AddSubmitter")
	if err != nil {
		return nil, err
	}
	return &MultisigAddSubmitterIterator{contract: _Multisig.contract, event: "AddSubmitter", logs: logs, sub: sub}, nil
}

// WatchAddSubmitter is a free log subscription operation binding the contract event 0x12fed2eaa2dbcff8567935f8d8c3eb8917ddd2b0466d2c17623503b8e64df06e.
//
// Solidity: event AddSubmitter(address _account)
func (_Multisig *MultisigFilterer) WatchAddSubmitter(opts *bind.WatchOpts, sink chan<- *MultisigAddSubmitter) (event.Subscription, error) {

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "AddSubmitter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigAddSubmitter)
				if err := _Multisig.contract.UnpackLog(event, "AddSubmitter", log); err != nil {
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

// ParseAddSubmitter is a log parse operation binding the contract event 0x12fed2eaa2dbcff8567935f8d8c3eb8917ddd2b0466d2c17623503b8e64df06e.
//
// Solidity: event AddSubmitter(address _account)
func (_Multisig *MultisigFilterer) ParseAddSubmitter(log types.Log) (*MultisigAddSubmitter, error) {
	event := new(MultisigAddSubmitter)
	if err := _Multisig.contract.UnpackLog(event, "AddSubmitter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigCancelTransactionIterator is returned from FilterCancelTransaction and is used to iterate over the raw logs and unpacked data for CancelTransaction events raised by the Multisig contract.
type MultisigCancelTransactionIterator struct {
	Event *MultisigCancelTransaction // Event containing the contract specifics and raw log

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
func (it *MultisigCancelTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigCancelTransaction)
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
		it.Event = new(MultisigCancelTransaction)
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
func (it *MultisigCancelTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigCancelTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigCancelTransaction represents a CancelTransaction event raised by the Multisig contract.
type MultisigCancelTransaction struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelTransaction is a free log retrieval operation binding the contract event 0x4674c50741f01a47c994e5eb211efddf3e030f68cc89f2f026ab25860cda2c23.
//
// Solidity: event CancelTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) FilterCancelTransaction(opts *bind.FilterOpts, _id []*big.Int) (*MultisigCancelTransactionIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "CancelTransaction", _idRule)
	if err != nil {
		return nil, err
	}
	return &MultisigCancelTransactionIterator{contract: _Multisig.contract, event: "CancelTransaction", logs: logs, sub: sub}, nil
}

// WatchCancelTransaction is a free log subscription operation binding the contract event 0x4674c50741f01a47c994e5eb211efddf3e030f68cc89f2f026ab25860cda2c23.
//
// Solidity: event CancelTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) WatchCancelTransaction(opts *bind.WatchOpts, sink chan<- *MultisigCancelTransaction, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "CancelTransaction", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigCancelTransaction)
				if err := _Multisig.contract.UnpackLog(event, "CancelTransaction", log); err != nil {
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

// ParseCancelTransaction is a log parse operation binding the contract event 0x4674c50741f01a47c994e5eb211efddf3e030f68cc89f2f026ab25860cda2c23.
//
// Solidity: event CancelTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) ParseCancelTransaction(log types.Log) (*MultisigCancelTransaction, error) {
	event := new(MultisigCancelTransaction)
	if err := _Multisig.contract.UnpackLog(event, "CancelTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigConfirmTransactionIterator is returned from FilterConfirmTransaction and is used to iterate over the raw logs and unpacked data for ConfirmTransaction events raised by the Multisig contract.
type MultisigConfirmTransactionIterator struct {
	Event *MultisigConfirmTransaction // Event containing the contract specifics and raw log

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
func (it *MultisigConfirmTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigConfirmTransaction)
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
		it.Event = new(MultisigConfirmTransaction)
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
func (it *MultisigConfirmTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigConfirmTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigConfirmTransaction represents a ConfirmTransaction event raised by the Multisig contract.
type MultisigConfirmTransaction struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterConfirmTransaction is a free log retrieval operation binding the contract event 0x0c22acbd1a0497750234bd4ea85e2a367754152040cc2d8b241a8ad98e7e54af.
//
// Solidity: event ConfirmTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) FilterConfirmTransaction(opts *bind.FilterOpts, _id []*big.Int) (*MultisigConfirmTransactionIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "ConfirmTransaction", _idRule)
	if err != nil {
		return nil, err
	}
	return &MultisigConfirmTransactionIterator{contract: _Multisig.contract, event: "ConfirmTransaction", logs: logs, sub: sub}, nil
}

// WatchConfirmTransaction is a free log subscription operation binding the contract event 0x0c22acbd1a0497750234bd4ea85e2a367754152040cc2d8b241a8ad98e7e54af.
//
// Solidity: event ConfirmTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) WatchConfirmTransaction(opts *bind.WatchOpts, sink chan<- *MultisigConfirmTransaction, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "ConfirmTransaction", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigConfirmTransaction)
				if err := _Multisig.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
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

// ParseConfirmTransaction is a log parse operation binding the contract event 0x0c22acbd1a0497750234bd4ea85e2a367754152040cc2d8b241a8ad98e7e54af.
//
// Solidity: event ConfirmTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) ParseConfirmTransaction(log types.Log) (*MultisigConfirmTransaction, error) {
	event := new(MultisigConfirmTransaction)
	if err := _Multisig.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Multisig contract.
type MultisigDepositIterator struct {
	Event *MultisigDeposit // Event containing the contract specifics and raw log

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
func (it *MultisigDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigDeposit)
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
		it.Event = new(MultisigDeposit)
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
func (it *MultisigDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigDeposit represents a Deposit event raised by the Multisig contract.
type MultisigDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed _sender, uint256 _amount)
func (_Multisig *MultisigFilterer) FilterDeposit(opts *bind.FilterOpts, _sender []common.Address) (*MultisigDepositIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "Deposit", _senderRule)
	if err != nil {
		return nil, err
	}
	return &MultisigDepositIterator{contract: _Multisig.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed _sender, uint256 _amount)
func (_Multisig *MultisigFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultisigDeposit, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "Deposit", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigDeposit)
				if err := _Multisig.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed _sender, uint256 _amount)
func (_Multisig *MultisigFilterer) ParseDeposit(log types.Log) (*MultisigDeposit, error) {
	event := new(MultisigDeposit)
	if err := _Multisig.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigExecuteTransactionIterator is returned from FilterExecuteTransaction and is used to iterate over the raw logs and unpacked data for ExecuteTransaction events raised by the Multisig contract.
type MultisigExecuteTransactionIterator struct {
	Event *MultisigExecuteTransaction // Event containing the contract specifics and raw log

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
func (it *MultisigExecuteTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigExecuteTransaction)
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
		it.Event = new(MultisigExecuteTransaction)
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
func (it *MultisigExecuteTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigExecuteTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigExecuteTransaction represents a ExecuteTransaction event raised by the Multisig contract.
type MultisigExecuteTransaction struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterExecuteTransaction is a free log retrieval operation binding the contract event 0xae30dc3f11bb6b178aafe5e7fc568fb6d87200068a944a8015c0db1b4533dbb8.
//
// Solidity: event ExecuteTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) FilterExecuteTransaction(opts *bind.FilterOpts, _id []*big.Int) (*MultisigExecuteTransactionIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "ExecuteTransaction", _idRule)
	if err != nil {
		return nil, err
	}
	return &MultisigExecuteTransactionIterator{contract: _Multisig.contract, event: "ExecuteTransaction", logs: logs, sub: sub}, nil
}

// WatchExecuteTransaction is a free log subscription operation binding the contract event 0xae30dc3f11bb6b178aafe5e7fc568fb6d87200068a944a8015c0db1b4533dbb8.
//
// Solidity: event ExecuteTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) WatchExecuteTransaction(opts *bind.WatchOpts, sink chan<- *MultisigExecuteTransaction, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "ExecuteTransaction", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigExecuteTransaction)
				if err := _Multisig.contract.UnpackLog(event, "ExecuteTransaction", log); err != nil {
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

// ParseExecuteTransaction is a log parse operation binding the contract event 0xae30dc3f11bb6b178aafe5e7fc568fb6d87200068a944a8015c0db1b4533dbb8.
//
// Solidity: event ExecuteTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) ParseExecuteTransaction(log types.Log) (*MultisigExecuteTransaction, error) {
	event := new(MultisigExecuteTransaction)
	if err := _Multisig.contract.UnpackLog(event, "ExecuteTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigRevokeOwnerIterator is returned from FilterRevokeOwner and is used to iterate over the raw logs and unpacked data for RevokeOwner events raised by the Multisig contract.
type MultisigRevokeOwnerIterator struct {
	Event *MultisigRevokeOwner // Event containing the contract specifics and raw log

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
func (it *MultisigRevokeOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigRevokeOwner)
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
		it.Event = new(MultisigRevokeOwner)
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
func (it *MultisigRevokeOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigRevokeOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigRevokeOwner represents a RevokeOwner event raised by the Multisig contract.
type MultisigRevokeOwner struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevokeOwner is a free log retrieval operation binding the contract event 0xadfb8ca2578f0288f371f60ed3bf03edbbc7867f43e893c2da79cc5b21e93001.
//
// Solidity: event RevokeOwner(address _account)
func (_Multisig *MultisigFilterer) FilterRevokeOwner(opts *bind.FilterOpts) (*MultisigRevokeOwnerIterator, error) {

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "RevokeOwner")
	if err != nil {
		return nil, err
	}
	return &MultisigRevokeOwnerIterator{contract: _Multisig.contract, event: "RevokeOwner", logs: logs, sub: sub}, nil
}

// WatchRevokeOwner is a free log subscription operation binding the contract event 0xadfb8ca2578f0288f371f60ed3bf03edbbc7867f43e893c2da79cc5b21e93001.
//
// Solidity: event RevokeOwner(address _account)
func (_Multisig *MultisigFilterer) WatchRevokeOwner(opts *bind.WatchOpts, sink chan<- *MultisigRevokeOwner) (event.Subscription, error) {

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "RevokeOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigRevokeOwner)
				if err := _Multisig.contract.UnpackLog(event, "RevokeOwner", log); err != nil {
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

// ParseRevokeOwner is a log parse operation binding the contract event 0xadfb8ca2578f0288f371f60ed3bf03edbbc7867f43e893c2da79cc5b21e93001.
//
// Solidity: event RevokeOwner(address _account)
func (_Multisig *MultisigFilterer) ParseRevokeOwner(log types.Log) (*MultisigRevokeOwner, error) {
	event := new(MultisigRevokeOwner)
	if err := _Multisig.contract.UnpackLog(event, "RevokeOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigRevokeSubmitterIterator is returned from FilterRevokeSubmitter and is used to iterate over the raw logs and unpacked data for RevokeSubmitter events raised by the Multisig contract.
type MultisigRevokeSubmitterIterator struct {
	Event *MultisigRevokeSubmitter // Event containing the contract specifics and raw log

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
func (it *MultisigRevokeSubmitterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigRevokeSubmitter)
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
		it.Event = new(MultisigRevokeSubmitter)
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
func (it *MultisigRevokeSubmitterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigRevokeSubmitterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigRevokeSubmitter represents a RevokeSubmitter event raised by the Multisig contract.
type MultisigRevokeSubmitter struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevokeSubmitter is a free log retrieval operation binding the contract event 0x8b213f89a1b0b8f73887d531a2c340210b9f8ea3192d6d1d2dd46158374afc9b.
//
// Solidity: event RevokeSubmitter(address _account)
func (_Multisig *MultisigFilterer) FilterRevokeSubmitter(opts *bind.FilterOpts) (*MultisigRevokeSubmitterIterator, error) {

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "RevokeSubmitter")
	if err != nil {
		return nil, err
	}
	return &MultisigRevokeSubmitterIterator{contract: _Multisig.contract, event: "RevokeSubmitter", logs: logs, sub: sub}, nil
}

// WatchRevokeSubmitter is a free log subscription operation binding the contract event 0x8b213f89a1b0b8f73887d531a2c340210b9f8ea3192d6d1d2dd46158374afc9b.
//
// Solidity: event RevokeSubmitter(address _account)
func (_Multisig *MultisigFilterer) WatchRevokeSubmitter(opts *bind.WatchOpts, sink chan<- *MultisigRevokeSubmitter) (event.Subscription, error) {

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "RevokeSubmitter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigRevokeSubmitter)
				if err := _Multisig.contract.UnpackLog(event, "RevokeSubmitter", log); err != nil {
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

// ParseRevokeSubmitter is a log parse operation binding the contract event 0x8b213f89a1b0b8f73887d531a2c340210b9f8ea3192d6d1d2dd46158374afc9b.
//
// Solidity: event RevokeSubmitter(address _account)
func (_Multisig *MultisigFilterer) ParseRevokeSubmitter(log types.Log) (*MultisigRevokeSubmitter, error) {
	event := new(MultisigRevokeSubmitter)
	if err := _Multisig.contract.UnpackLog(event, "RevokeSubmitter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigRevokeTransactionIterator is returned from FilterRevokeTransaction and is used to iterate over the raw logs and unpacked data for RevokeTransaction events raised by the Multisig contract.
type MultisigRevokeTransactionIterator struct {
	Event *MultisigRevokeTransaction // Event containing the contract specifics and raw log

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
func (it *MultisigRevokeTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigRevokeTransaction)
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
		it.Event = new(MultisigRevokeTransaction)
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
func (it *MultisigRevokeTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigRevokeTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigRevokeTransaction represents a RevokeTransaction event raised by the Multisig contract.
type MultisigRevokeTransaction struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRevokeTransaction is a free log retrieval operation binding the contract event 0x92a917be073b79ab50f9d997e9514a5d101e052ce04228014947e7883a7c0346.
//
// Solidity: event RevokeTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) FilterRevokeTransaction(opts *bind.FilterOpts, _id []*big.Int) (*MultisigRevokeTransactionIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "RevokeTransaction", _idRule)
	if err != nil {
		return nil, err
	}
	return &MultisigRevokeTransactionIterator{contract: _Multisig.contract, event: "RevokeTransaction", logs: logs, sub: sub}, nil
}

// WatchRevokeTransaction is a free log subscription operation binding the contract event 0x92a917be073b79ab50f9d997e9514a5d101e052ce04228014947e7883a7c0346.
//
// Solidity: event RevokeTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) WatchRevokeTransaction(opts *bind.WatchOpts, sink chan<- *MultisigRevokeTransaction, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "RevokeTransaction", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigRevokeTransaction)
				if err := _Multisig.contract.UnpackLog(event, "RevokeTransaction", log); err != nil {
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

// ParseRevokeTransaction is a log parse operation binding the contract event 0x92a917be073b79ab50f9d997e9514a5d101e052ce04228014947e7883a7c0346.
//
// Solidity: event RevokeTransaction(uint256 indexed _id)
func (_Multisig *MultisigFilterer) ParseRevokeTransaction(log types.Log) (*MultisigRevokeTransaction, error) {
	event := new(MultisigRevokeTransaction)
	if err := _Multisig.contract.UnpackLog(event, "RevokeTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigSubmitTransactionIterator is returned from FilterSubmitTransaction and is used to iterate over the raw logs and unpacked data for SubmitTransaction events raised by the Multisig contract.
type MultisigSubmitTransactionIterator struct {
	Event *MultisigSubmitTransaction // Event containing the contract specifics and raw log

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
func (it *MultisigSubmitTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigSubmitTransaction)
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
		it.Event = new(MultisigSubmitTransaction)
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
func (it *MultisigSubmitTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigSubmitTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigSubmitTransaction represents a SubmitTransaction event raised by the Multisig contract.
type MultisigSubmitTransaction struct {
	Id       *big.Int
	Target   common.Address
	Data     []byte
	Deadline *big.Int
	Note     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitTransaction is a free log retrieval operation binding the contract event 0xb8dcc6ebcc290fc38fc7594290649a4e818da5c65971c2f86d232bcf2bba4e58.
//
// Solidity: event SubmitTransaction(uint256 indexed _id, address _target, bytes _data, uint256 _deadline, string _note)
func (_Multisig *MultisigFilterer) FilterSubmitTransaction(opts *bind.FilterOpts, _id []*big.Int) (*MultisigSubmitTransactionIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "SubmitTransaction", _idRule)
	if err != nil {
		return nil, err
	}
	return &MultisigSubmitTransactionIterator{contract: _Multisig.contract, event: "SubmitTransaction", logs: logs, sub: sub}, nil
}

// WatchSubmitTransaction is a free log subscription operation binding the contract event 0xb8dcc6ebcc290fc38fc7594290649a4e818da5c65971c2f86d232bcf2bba4e58.
//
// Solidity: event SubmitTransaction(uint256 indexed _id, address _target, bytes _data, uint256 _deadline, string _note)
func (_Multisig *MultisigFilterer) WatchSubmitTransaction(opts *bind.WatchOpts, sink chan<- *MultisigSubmitTransaction, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "SubmitTransaction", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigSubmitTransaction)
				if err := _Multisig.contract.UnpackLog(event, "SubmitTransaction", log); err != nil {
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

// InputDataDecoder is a log parse operation binding the contract event 0xb8dcc6ebcc290fc38fc7594290649a4e818da5c65971c2f86d232bcf2bba4e58.
//
// Solidity: event SubmitTransaction(uint256 indexed _id, address _target, bytes _data, uint256 _deadline, string _note)
func (_Multisig *MultisigFilterer) ParseSubmitTransaction(log types.Log) (*MultisigSubmitTransaction, error) {
	event := new(MultisigSubmitTransaction)
	if err := _Multisig.contract.UnpackLog(event, "SubmitTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigUpdateWeightIterator is returned from FilterUpdateWeight and is used to iterate over the raw logs and unpacked data for UpdateWeight events raised by the Multisig contract.
type MultisigUpdateWeightIterator struct {
	Event *MultisigUpdateWeight // Event containing the contract specifics and raw log

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
func (it *MultisigUpdateWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigUpdateWeight)
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
		it.Event = new(MultisigUpdateWeight)
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
func (it *MultisigUpdateWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigUpdateWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigUpdateWeight represents a UpdateWeight event raised by the Multisig contract.
type MultisigUpdateWeight struct {
	NewWeight *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateWeight is a free log retrieval operation binding the contract event 0x5fee8331c1b3641338a538b509e8ffd4f30fcbed95cadda5d1dcad3c3d54785b.
//
// Solidity: event UpdateWeight(uint256 _newWeight)
func (_Multisig *MultisigFilterer) FilterUpdateWeight(opts *bind.FilterOpts) (*MultisigUpdateWeightIterator, error) {

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "UpdateWeight")
	if err != nil {
		return nil, err
	}
	return &MultisigUpdateWeightIterator{contract: _Multisig.contract, event: "UpdateWeight", logs: logs, sub: sub}, nil
}

// WatchUpdateWeight is a free log subscription operation binding the contract event 0x5fee8331c1b3641338a538b509e8ffd4f30fcbed95cadda5d1dcad3c3d54785b.
//
// Solidity: event UpdateWeight(uint256 _newWeight)
func (_Multisig *MultisigFilterer) WatchUpdateWeight(opts *bind.WatchOpts, sink chan<- *MultisigUpdateWeight) (event.Subscription, error) {

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "UpdateWeight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigUpdateWeight)
				if err := _Multisig.contract.UnpackLog(event, "UpdateWeight", log); err != nil {
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

// ParseUpdateWeight is a log parse operation binding the contract event 0x5fee8331c1b3641338a538b509e8ffd4f30fcbed95cadda5d1dcad3c3d54785b.
//
// Solidity: event UpdateWeight(uint256 _newWeight)
func (_Multisig *MultisigFilterer) ParseUpdateWeight(log types.Log) (*MultisigUpdateWeight, error) {
	event := new(MultisigUpdateWeight)
	if err := _Multisig.contract.UnpackLog(event, "UpdateWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Multisig contract.
type MultisigWithdrawIterator struct {
	Event *MultisigWithdraw // Event containing the contract specifics and raw log

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
func (it *MultisigWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigWithdraw)
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
		it.Event = new(MultisigWithdraw)
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
func (it *MultisigWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigWithdraw represents a Withdraw event raised by the Multisig contract.
type MultisigWithdraw struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _receiver, uint256 _amount)
func (_Multisig *MultisigFilterer) FilterWithdraw(opts *bind.FilterOpts, _receiver []common.Address) (*MultisigWithdrawIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _Multisig.contract.FilterLogs(opts, "Withdraw", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &MultisigWithdrawIterator{contract: _Multisig.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _receiver, uint256 _amount)
func (_Multisig *MultisigFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MultisigWithdraw, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _Multisig.contract.WatchLogs(opts, "Withdraw", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigWithdraw)
				if err := _Multisig.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _receiver, uint256 _amount)
func (_Multisig *MultisigFilterer) ParseWithdraw(log types.Log) (*MultisigWithdraw, error) {
	event := new(MultisigWithdraw)
	if err := _Multisig.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
