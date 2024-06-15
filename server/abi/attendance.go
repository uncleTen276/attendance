// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package attendance_record

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

// AttendanceEmployeeRecord is an auto generated low-level Go binding around an user-defined struct.
type AttendanceEmployeeRecord struct {
	RecordId    *big.Int
	EmployeeId  *big.Int
	CheckInTime string
	Position    string
	Details     string
}

// AttendanceRecordMetaData contains all meta data concerning the AttendanceRecord contract.
var AttendanceRecordMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"attendanceRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"checkInTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"employeeRecordIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_employeeId\",\"type\":\"uint256\"}],\"name\":\"getAttendanceByEmployee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"recordId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"checkInTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structAttendance.EmployeeRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRecordId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"checkInTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"}],\"name\":\"recordAtttendance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_recordId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_checkInTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_details\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_position\",\"type\":\"string\"}],\"name\":\"updateAttendance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AttendanceRecordABI is the input ABI used to generate the binding from.
// Deprecated: Use AttendanceRecordMetaData.ABI instead.
var AttendanceRecordABI = AttendanceRecordMetaData.ABI

// AttendanceRecord is an auto generated Go binding around an Ethereum contract.
type AttendanceRecord struct {
	AttendanceRecordCaller     // Read-only binding to the contract
	AttendanceRecordTransactor // Write-only binding to the contract
	AttendanceRecordFilterer   // Log filterer for contract events
}

// AttendanceRecordCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttendanceRecordCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceRecordTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttendanceRecordTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceRecordFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttendanceRecordFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceRecordSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttendanceRecordSession struct {
	Contract     *AttendanceRecord // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttendanceRecordCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttendanceRecordCallerSession struct {
	Contract *AttendanceRecordCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AttendanceRecordTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttendanceRecordTransactorSession struct {
	Contract     *AttendanceRecordTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AttendanceRecordRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttendanceRecordRaw struct {
	Contract *AttendanceRecord // Generic contract binding to access the raw methods on
}

// AttendanceRecordCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttendanceRecordCallerRaw struct {
	Contract *AttendanceRecordCaller // Generic read-only contract binding to access the raw methods on
}

// AttendanceRecordTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttendanceRecordTransactorRaw struct {
	Contract *AttendanceRecordTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttendanceRecord creates a new instance of AttendanceRecord, bound to a specific deployed contract.
func NewAttendanceRecord(address common.Address, backend bind.ContractBackend) (*AttendanceRecord, error) {
	contract, err := bindAttendanceRecord(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttendanceRecord{AttendanceRecordCaller: AttendanceRecordCaller{contract: contract}, AttendanceRecordTransactor: AttendanceRecordTransactor{contract: contract}, AttendanceRecordFilterer: AttendanceRecordFilterer{contract: contract}}, nil
}

// NewAttendanceRecordCaller creates a new read-only instance of AttendanceRecord, bound to a specific deployed contract.
func NewAttendanceRecordCaller(address common.Address, caller bind.ContractCaller) (*AttendanceRecordCaller, error) {
	contract, err := bindAttendanceRecord(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttendanceRecordCaller{contract: contract}, nil
}

// NewAttendanceRecordTransactor creates a new write-only instance of AttendanceRecord, bound to a specific deployed contract.
func NewAttendanceRecordTransactor(address common.Address, transactor bind.ContractTransactor) (*AttendanceRecordTransactor, error) {
	contract, err := bindAttendanceRecord(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttendanceRecordTransactor{contract: contract}, nil
}

// NewAttendanceRecordFilterer creates a new log filterer instance of AttendanceRecord, bound to a specific deployed contract.
func NewAttendanceRecordFilterer(address common.Address, filterer bind.ContractFilterer) (*AttendanceRecordFilterer, error) {
	contract, err := bindAttendanceRecord(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttendanceRecordFilterer{contract: contract}, nil
}

// bindAttendanceRecord binds a generic wrapper to an already deployed contract.
func bindAttendanceRecord(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AttendanceRecordMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttendanceRecord *AttendanceRecordRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttendanceRecord.Contract.AttendanceRecordCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttendanceRecord *AttendanceRecordRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttendanceRecord.Contract.AttendanceRecordTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttendanceRecord *AttendanceRecordRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttendanceRecord.Contract.AttendanceRecordTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttendanceRecord *AttendanceRecordCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttendanceRecord.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttendanceRecord *AttendanceRecordTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttendanceRecord.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttendanceRecord *AttendanceRecordTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttendanceRecord.Contract.contract.Transact(opts, method, params...)
}

// AttendanceRecords is a free data retrieval call binding the contract method 0x5763c184.
//
// Solidity: function attendanceRecords(uint256 ) view returns(uint256 recordId, uint256 employeeId, string checkInTime, string position, string details)
func (_AttendanceRecord *AttendanceRecordCaller) AttendanceRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RecordId    *big.Int
	EmployeeId  *big.Int
	CheckInTime string
	Position    string
	Details     string
}, error) {
	var out []interface{}
	err := _AttendanceRecord.contract.Call(opts, &out, "attendanceRecords", arg0)

	outstruct := new(struct {
		RecordId    *big.Int
		EmployeeId  *big.Int
		CheckInTime string
		Position    string
		Details     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RecordId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EmployeeId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CheckInTime = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Position = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Details = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// AttendanceRecords is a free data retrieval call binding the contract method 0x5763c184.
//
// Solidity: function attendanceRecords(uint256 ) view returns(uint256 recordId, uint256 employeeId, string checkInTime, string position, string details)
func (_AttendanceRecord *AttendanceRecordSession) AttendanceRecords(arg0 *big.Int) (struct {
	RecordId    *big.Int
	EmployeeId  *big.Int
	CheckInTime string
	Position    string
	Details     string
}, error) {
	return _AttendanceRecord.Contract.AttendanceRecords(&_AttendanceRecord.CallOpts, arg0)
}

// AttendanceRecords is a free data retrieval call binding the contract method 0x5763c184.
//
// Solidity: function attendanceRecords(uint256 ) view returns(uint256 recordId, uint256 employeeId, string checkInTime, string position, string details)
func (_AttendanceRecord *AttendanceRecordCallerSession) AttendanceRecords(arg0 *big.Int) (struct {
	RecordId    *big.Int
	EmployeeId  *big.Int
	CheckInTime string
	Position    string
	Details     string
}, error) {
	return _AttendanceRecord.Contract.AttendanceRecords(&_AttendanceRecord.CallOpts, arg0)
}

// EmployeeRecordIds is a free data retrieval call binding the contract method 0x0618c89c.
//
// Solidity: function employeeRecordIds(uint256 , uint256 ) view returns(uint256)
func (_AttendanceRecord *AttendanceRecordCaller) EmployeeRecordIds(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AttendanceRecord.contract.Call(opts, &out, "employeeRecordIds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmployeeRecordIds is a free data retrieval call binding the contract method 0x0618c89c.
//
// Solidity: function employeeRecordIds(uint256 , uint256 ) view returns(uint256)
func (_AttendanceRecord *AttendanceRecordSession) EmployeeRecordIds(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _AttendanceRecord.Contract.EmployeeRecordIds(&_AttendanceRecord.CallOpts, arg0, arg1)
}

// EmployeeRecordIds is a free data retrieval call binding the contract method 0x0618c89c.
//
// Solidity: function employeeRecordIds(uint256 , uint256 ) view returns(uint256)
func (_AttendanceRecord *AttendanceRecordCallerSession) EmployeeRecordIds(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _AttendanceRecord.Contract.EmployeeRecordIds(&_AttendanceRecord.CallOpts, arg0, arg1)
}

// GetAttendanceByEmployee is a free data retrieval call binding the contract method 0x2e74c140.
//
// Solidity: function getAttendanceByEmployee(uint256 _employeeId) view returns((uint256,uint256,string,string,string)[])
func (_AttendanceRecord *AttendanceRecordCaller) GetAttendanceByEmployee(opts *bind.CallOpts, _employeeId *big.Int) ([]AttendanceEmployeeRecord, error) {
	var out []interface{}
	err := _AttendanceRecord.contract.Call(opts, &out, "getAttendanceByEmployee", _employeeId)

	if err != nil {
		return *new([]AttendanceEmployeeRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttendanceEmployeeRecord)).(*[]AttendanceEmployeeRecord)

	return out0, err

}

// GetAttendanceByEmployee is a free data retrieval call binding the contract method 0x2e74c140.
//
// Solidity: function getAttendanceByEmployee(uint256 _employeeId) view returns((uint256,uint256,string,string,string)[])
func (_AttendanceRecord *AttendanceRecordSession) GetAttendanceByEmployee(_employeeId *big.Int) ([]AttendanceEmployeeRecord, error) {
	return _AttendanceRecord.Contract.GetAttendanceByEmployee(&_AttendanceRecord.CallOpts, _employeeId)
}

// GetAttendanceByEmployee is a free data retrieval call binding the contract method 0x2e74c140.
//
// Solidity: function getAttendanceByEmployee(uint256 _employeeId) view returns((uint256,uint256,string,string,string)[])
func (_AttendanceRecord *AttendanceRecordCallerSession) GetAttendanceByEmployee(_employeeId *big.Int) ([]AttendanceEmployeeRecord, error) {
	return _AttendanceRecord.Contract.GetAttendanceByEmployee(&_AttendanceRecord.CallOpts, _employeeId)
}

// NextRecordId is a free data retrieval call binding the contract method 0xdbde207d.
//
// Solidity: function nextRecordId() view returns(uint256)
func (_AttendanceRecord *AttendanceRecordCaller) NextRecordId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttendanceRecord.contract.Call(opts, &out, "nextRecordId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRecordId is a free data retrieval call binding the contract method 0xdbde207d.
//
// Solidity: function nextRecordId() view returns(uint256)
func (_AttendanceRecord *AttendanceRecordSession) NextRecordId() (*big.Int, error) {
	return _AttendanceRecord.Contract.NextRecordId(&_AttendanceRecord.CallOpts)
}

// NextRecordId is a free data retrieval call binding the contract method 0xdbde207d.
//
// Solidity: function nextRecordId() view returns(uint256)
func (_AttendanceRecord *AttendanceRecordCallerSession) NextRecordId() (*big.Int, error) {
	return _AttendanceRecord.Contract.NextRecordId(&_AttendanceRecord.CallOpts)
}

// RecordAtttendance is a paid mutator transaction binding the contract method 0x29868f95.
//
// Solidity: function recordAtttendance(uint256 employeeId, string checkInTime, string details, string position) returns(uint256)
func (_AttendanceRecord *AttendanceRecordTransactor) RecordAtttendance(opts *bind.TransactOpts, employeeId *big.Int, checkInTime string, details string, position string) (*types.Transaction, error) {
	return _AttendanceRecord.contract.Transact(opts, "recordAtttendance", employeeId, checkInTime, details, position)
}

// RecordAtttendance is a paid mutator transaction binding the contract method 0x29868f95.
//
// Solidity: function recordAtttendance(uint256 employeeId, string checkInTime, string details, string position) returns(uint256)
func (_AttendanceRecord *AttendanceRecordSession) RecordAtttendance(employeeId *big.Int, checkInTime string, details string, position string) (*types.Transaction, error) {
	return _AttendanceRecord.Contract.RecordAtttendance(&_AttendanceRecord.TransactOpts, employeeId, checkInTime, details, position)
}

// RecordAtttendance is a paid mutator transaction binding the contract method 0x29868f95.
//
// Solidity: function recordAtttendance(uint256 employeeId, string checkInTime, string details, string position) returns(uint256)
func (_AttendanceRecord *AttendanceRecordTransactorSession) RecordAtttendance(employeeId *big.Int, checkInTime string, details string, position string) (*types.Transaction, error) {
	return _AttendanceRecord.Contract.RecordAtttendance(&_AttendanceRecord.TransactOpts, employeeId, checkInTime, details, position)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0x727d63be.
//
// Solidity: function updateAttendance(uint256 _recordId, string _checkInTime, string _details, string _position) returns()
func (_AttendanceRecord *AttendanceRecordTransactor) UpdateAttendance(opts *bind.TransactOpts, _recordId *big.Int, _checkInTime string, _details string, _position string) (*types.Transaction, error) {
	return _AttendanceRecord.contract.Transact(opts, "updateAttendance", _recordId, _checkInTime, _details, _position)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0x727d63be.
//
// Solidity: function updateAttendance(uint256 _recordId, string _checkInTime, string _details, string _position) returns()
func (_AttendanceRecord *AttendanceRecordSession) UpdateAttendance(_recordId *big.Int, _checkInTime string, _details string, _position string) (*types.Transaction, error) {
	return _AttendanceRecord.Contract.UpdateAttendance(&_AttendanceRecord.TransactOpts, _recordId, _checkInTime, _details, _position)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0x727d63be.
//
// Solidity: function updateAttendance(uint256 _recordId, string _checkInTime, string _details, string _position) returns()
func (_AttendanceRecord *AttendanceRecordTransactorSession) UpdateAttendance(_recordId *big.Int, _checkInTime string, _details string, _position string) (*types.Transaction, error) {
	return _AttendanceRecord.Contract.UpdateAttendance(&_AttendanceRecord.TransactOpts, _recordId, _checkInTime, _details, _position)
}
