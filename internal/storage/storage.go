// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// VotingStorageABI is the input ABI used to generate the binding from.
const VotingStorageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"addToWhiteList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"indexes\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"encryptedOption\",\"type\":\"string\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vote\",\"type\":\"string\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"registerVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VotingStorageBin is the compiled bytecode used for deploying new contracts.
const VotingStorageBin = `0x60806040526003805467ffffffffffffffff1916905534801561002157600080fd5b5060008054600160a060020a03191633179055610913806100436000396000f3fe608060405234801561001057600080fd5b506004361061009a576000357c0100000000000000000000000000000000000000000000000000000000900480638778b27d116100785780638778b27d146101445780638da5cb5b14610169578063d8bff5a514610171578063f97a73901461028e5761009a565b8063372c12b11461009f57806347ee0394146100d957806382897ea714610101575b600080fd5b6100c5600480360360208110156100b557600080fd5b5035600160a060020a03166103bb565b604080519115158252519081900360200190f35b6100ff600480360360208110156100ef57600080fd5b5035600160a060020a03166103d0565b005b6101286004803603602081101561011757600080fd5b503567ffffffffffffffff1661040b565b60408051600160a060020a039092168252519081900360200190f35b61014c610426565b6040805167ffffffffffffffff9092168252519081900360200190f35b610128610436565b6101976004803603602081101561018757600080fd5b5035600160a060020a0316610445565b6040518084600160a060020a0316600160a060020a031681526020018060200180602001838103835285818151815260200191508051906020019080838360005b838110156101f05781810151838201526020016101d8565b50505050905090810190601f16801561021d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610250578181015183820152602001610238565b50505050905090810190601f16801561027d5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b6100ff600480360360408110156102a457600080fd5b8101906020810181356401000000008111156102bf57600080fd5b8201836020820111156102d157600080fd5b803590602001918460018302840111640100000000831117156102f357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561034657600080fd5b82018360208201111561035857600080fd5b8035906020019184600183028401116401000000008311171561037a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610588945050505050565b60046020526000908152604090205460ff1681565b600054600160a060020a031633146103e757600080fd5b600160a060020a03166000908152600460205260409020805460ff19166001179055565b600160205260009081526040902054600160a060020a031681565b60035467ffffffffffffffff1681565b600054600160a060020a031681565b6002602081815260009283526040928390208054600180830180548751601f60001994831615610100029490940190911696909604918201859004850286018501909652808552600160a060020a03909116949193928301828280156104ec5780601f106104c1576101008083540402835291602001916104ec565b820191906000526020600020905b8154815290600101906020018083116104cf57829003601f168201915b50505060028085018054604080516020601f600019610100600187161502019094169590950492830185900485028101850190915281815295969594509092509083018282801561057e5780601f106105535761010080835404028352916020019161057e565b820191906000526020600020905b81548152906001019060200180831161056157829003601f168201915b5050505050905083565b6003805467ffffffffffffffff8082166001011667ffffffffffffffff19909116179055604051825160009161064c91859133916020918201918291908501908083835b602083106105eb5780518252601f1990920191602091820191016105cc565b6001836020036101000a03801982511681845116808217855250505050505090500182600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401925050506040516020818303038152906040526106fd565b9050600061065a8284610717565b9050600160a060020a038116331461067157600080fd5b610679610821565b6020818101868152604080840187905233808552600090815260028452208351815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039091161781559051805184936106d792600185019291019061084c565b50604082015180516106f391600284019160209091019061084c565b5050505050505050565b805160208201206000906107108161079f565b9392505050565b600080600080610726856107f0565b60408051600080825260208083018085528d905260ff87168385015260608301869052608083018590529251959850939650919450919260019260a0808401939192601f1981019281900390910190855afa158015610789573d6000803e3d6000fd5b5050604051601f19015198975050505050505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b60008060008351604114151561080557600080fd5b5050506020810151604082015160609092015160001a92909190565b6060604051908101604052806000600160a060020a0316815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061088d57805160ff19168380011785556108ba565b828001600101855582156108ba579182015b828111156108ba57825182559160200191906001019061089f565b506108c69291506108ca565b5090565b6108e491905b808211156108c657600081556001016108d0565b9056fea165627a7a72305820dc1bc85694c6558b06dca20ac6a20df7892938a6c39021b79def0f6a793510640029`

// DeployVotingStorage deploys a new Ethereum contract, binding an instance of VotingStorage to it.
func DeployVotingStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VotingStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VotingStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VotingStorage{VotingStorageCaller: VotingStorageCaller{contract: contract}, VotingStorageTransactor: VotingStorageTransactor{contract: contract}, VotingStorageFilterer: VotingStorageFilterer{contract: contract}}, nil
}

// VotingStorage is an auto generated Go binding around an Ethereum contract.
type VotingStorage struct {
	VotingStorageCaller     // Read-only binding to the contract
	VotingStorageTransactor // Write-only binding to the contract
	VotingStorageFilterer   // Log filterer for contract events
}

// VotingStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingStorageSession struct {
	Contract     *VotingStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingStorageCallerSession struct {
	Contract *VotingStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VotingStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingStorageTransactorSession struct {
	Contract     *VotingStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VotingStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingStorageRaw struct {
	Contract *VotingStorage // Generic contract binding to access the raw methods on
}

// VotingStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingStorageCallerRaw struct {
	Contract *VotingStorageCaller // Generic read-only contract binding to access the raw methods on
}

// VotingStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingStorageTransactorRaw struct {
	Contract *VotingStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotingStorage creates a new instance of VotingStorage, bound to a specific deployed contract.
func NewVotingStorage(address common.Address, backend bind.ContractBackend) (*VotingStorage, error) {
	contract, err := bindVotingStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotingStorage{VotingStorageCaller: VotingStorageCaller{contract: contract}, VotingStorageTransactor: VotingStorageTransactor{contract: contract}, VotingStorageFilterer: VotingStorageFilterer{contract: contract}}, nil
}

// NewVotingStorageCaller creates a new read-only instance of VotingStorage, bound to a specific deployed contract.
func NewVotingStorageCaller(address common.Address, caller bind.ContractCaller) (*VotingStorageCaller, error) {
	contract, err := bindVotingStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingStorageCaller{contract: contract}, nil
}

// NewVotingStorageTransactor creates a new write-only instance of VotingStorage, bound to a specific deployed contract.
func NewVotingStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingStorageTransactor, error) {
	contract, err := bindVotingStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingStorageTransactor{contract: contract}, nil
}

// NewVotingStorageFilterer creates a new log filterer instance of VotingStorage, bound to a specific deployed contract.
func NewVotingStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingStorageFilterer, error) {
	contract, err := bindVotingStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingStorageFilterer{contract: contract}, nil
}

// bindVotingStorage binds a generic wrapper to an already deployed contract.
func bindVotingStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingStorage *VotingStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VotingStorage.Contract.VotingStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingStorage *VotingStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingStorage.Contract.VotingStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingStorage *VotingStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingStorage.Contract.VotingStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingStorage *VotingStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VotingStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingStorage *VotingStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingStorage *VotingStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingStorage.Contract.contract.Transact(opts, method, params...)
}

// Indexes is a free data retrieval call binding the contract method 0x82897ea7.
//
// Solidity: function indexes(uint64 ) constant returns(address)
func (_VotingStorage *VotingStorageCaller) Indexes(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VotingStorage.contract.Call(opts, out, "indexes", arg0)
	return *ret0, err
}

// Indexes is a free data retrieval call binding the contract method 0x82897ea7.
//
// Solidity: function indexes(uint64 ) constant returns(address)
func (_VotingStorage *VotingStorageSession) Indexes(arg0 uint64) (common.Address, error) {
	return _VotingStorage.Contract.Indexes(&_VotingStorage.CallOpts, arg0)
}

// Indexes is a free data retrieval call binding the contract method 0x82897ea7.
//
// Solidity: function indexes(uint64 ) constant returns(address)
func (_VotingStorage *VotingStorageCallerSession) Indexes(arg0 uint64) (common.Address, error) {
	return _VotingStorage.Contract.Indexes(&_VotingStorage.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VotingStorage *VotingStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VotingStorage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VotingStorage *VotingStorageSession) Owner() (common.Address, error) {
	return _VotingStorage.Contract.Owner(&_VotingStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VotingStorage *VotingStorageCallerSession) Owner() (common.Address, error) {
	return _VotingStorage.Contract.Owner(&_VotingStorage.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address voter, string encryptedOption, bytes signature)
func (_VotingStorage *VotingStorageCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Voter           common.Address
	EncryptedOption string
	Signature       []byte
}, error) {
	ret := new(struct {
		Voter           common.Address
		EncryptedOption string
		Signature       []byte
	})
	out := ret
	err := _VotingStorage.contract.Call(opts, out, "votes", arg0)
	return *ret, err
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address voter, string encryptedOption, bytes signature)
func (_VotingStorage *VotingStorageSession) Votes(arg0 common.Address) (struct {
	Voter           common.Address
	EncryptedOption string
	Signature       []byte
}, error) {
	return _VotingStorage.Contract.Votes(&_VotingStorage.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address voter, string encryptedOption, bytes signature)
func (_VotingStorage *VotingStorageCallerSession) Votes(arg0 common.Address) (struct {
	Voter           common.Address
	EncryptedOption string
	Signature       []byte
}, error) {
	return _VotingStorage.Contract.Votes(&_VotingStorage.CallOpts, arg0)
}

// VotesCount is a free data retrieval call binding the contract method 0x8778b27d.
//
// Solidity: function votesCount() constant returns(uint64)
func (_VotingStorage *VotingStorageCaller) VotesCount(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _VotingStorage.contract.Call(opts, out, "votesCount")
	return *ret0, err
}

// VotesCount is a free data retrieval call binding the contract method 0x8778b27d.
//
// Solidity: function votesCount() constant returns(uint64)
func (_VotingStorage *VotingStorageSession) VotesCount() (uint64, error) {
	return _VotingStorage.Contract.VotesCount(&_VotingStorage.CallOpts)
}

// VotesCount is a free data retrieval call binding the contract method 0x8778b27d.
//
// Solidity: function votesCount() constant returns(uint64)
func (_VotingStorage *VotingStorageCallerSession) VotesCount() (uint64, error) {
	return _VotingStorage.Contract.VotesCount(&_VotingStorage.CallOpts)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) constant returns(bool)
func (_VotingStorage *VotingStorageCaller) WhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VotingStorage.contract.Call(opts, out, "whiteList", arg0)
	return *ret0, err
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) constant returns(bool)
func (_VotingStorage *VotingStorageSession) WhiteList(arg0 common.Address) (bool, error) {
	return _VotingStorage.Contract.WhiteList(&_VotingStorage.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) constant returns(bool)
func (_VotingStorage *VotingStorageCallerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _VotingStorage.Contract.WhiteList(&_VotingStorage.CallOpts, arg0)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(address _voter) returns()
func (_VotingStorage *VotingStorageTransactor) AddToWhiteList(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _VotingStorage.contract.Transact(opts, "addToWhiteList", _voter)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(address _voter) returns()
func (_VotingStorage *VotingStorageSession) AddToWhiteList(_voter common.Address) (*types.Transaction, error) {
	return _VotingStorage.Contract.AddToWhiteList(&_VotingStorage.TransactOpts, _voter)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(address _voter) returns()
func (_VotingStorage *VotingStorageTransactorSession) AddToWhiteList(_voter common.Address) (*types.Transaction, error) {
	return _VotingStorage.Contract.AddToWhiteList(&_VotingStorage.TransactOpts, _voter)
}

// RegisterVote is a paid mutator transaction binding the contract method 0xf97a7390.
//
// Solidity: function registerVote(string _vote, bytes _sig) returns()
func (_VotingStorage *VotingStorageTransactor) RegisterVote(opts *bind.TransactOpts, _vote string, _sig []byte) (*types.Transaction, error) {
	return _VotingStorage.contract.Transact(opts, "registerVote", _vote, _sig)
}

// RegisterVote is a paid mutator transaction binding the contract method 0xf97a7390.
//
// Solidity: function registerVote(string _vote, bytes _sig) returns()
func (_VotingStorage *VotingStorageSession) RegisterVote(_vote string, _sig []byte) (*types.Transaction, error) {
	return _VotingStorage.Contract.RegisterVote(&_VotingStorage.TransactOpts, _vote, _sig)
}

// RegisterVote is a paid mutator transaction binding the contract method 0xf97a7390.
//
// Solidity: function registerVote(string _vote, bytes _sig) returns()
func (_VotingStorage *VotingStorageTransactorSession) RegisterVote(_vote string, _sig []byte) (*types.Transaction, error) {
	return _VotingStorage.Contract.RegisterVote(&_VotingStorage.TransactOpts, _vote, _sig)
}
