// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc8183

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

// AgenticCommerceJob is an auto generated low-level Go binding around an user-defined struct.
type AgenticCommerceJob struct {
	Id              *big.Int
	Client          common.Address
	Provider        common.Address
	Evaluator       common.Address
	Description     string
	Budget          *big.Int
	ExpiredAt       *big.Int
	Status          uint8
	Hook            common.Address
	PaymentToken    common.Address
	ProviderAgentId *big.Int
}

// ERC8183MetaData contains all meta data concerning the ERC8183 contract.
var ERC8183MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BudgetMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiryTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeesTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HookNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHook\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroBudget\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BudgetSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"evaluator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EvaluatorFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBP\",\"type\":\"uint256\"}],\"name\":\"EvaluatorFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"HookWhitelistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"evaluator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"name\":\"JobCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"evaluator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"JobCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"JobExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"JobFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rejector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"name\":\"JobRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"deliverable\",\"type\":\"bytes32\"}],\"name\":\"JobSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"platformTreasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PlatformFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBP\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"}],\"name\":\"PlatformFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"ProviderSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"claimRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"optParams\",\"type\":\"bytes\"}],\"name\":\"complete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"evaluator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"providerAgentId\",\"type\":\"uint256\"}],\"name\":\"createJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"evaluatorFeeBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedBudget\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"optParams\",\"type\":\"bytes\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"getJob\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"evaluator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"enumAgenticCommerce.JobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"providerAgentId\",\"type\":\"uint256\"}],\"internalType\":\"structAgenticCommerce.Job\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"evaluator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"enumAgenticCommerce.JobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"providerAgentId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeeBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"optParams\",\"type\":\"bytes\"}],\"name\":\"reject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"optParams\",\"type\":\"bytes\"}],\"name\":\"setBudget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBP_\",\"type\":\"uint256\"}],\"name\":\"setEvaluatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setHookWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBP_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"name\":\"setPlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"setProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"deliverable\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"optParams\",\"type\":\"bytes\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedHooks\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC8183ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC8183MetaData.ABI instead.
var ERC8183ABI = ERC8183MetaData.ABI

// ERC8183 is an auto generated Go binding around an Ethereum contract.
type ERC8183 struct {
	ERC8183Caller     // Read-only binding to the contract
	ERC8183Transactor // Write-only binding to the contract
	ERC8183Filterer   // Log filterer for contract events
}

// ERC8183Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC8183Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC8183Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC8183Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC8183Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC8183Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC8183Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC8183Session struct {
	Contract     *ERC8183          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC8183CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC8183CallerSession struct {
	Contract *ERC8183Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC8183TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC8183TransactorSession struct {
	Contract     *ERC8183Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC8183Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC8183Raw struct {
	Contract *ERC8183 // Generic contract binding to access the raw methods on
}

// ERC8183CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC8183CallerRaw struct {
	Contract *ERC8183Caller // Generic read-only contract binding to access the raw methods on
}

// ERC8183TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC8183TransactorRaw struct {
	Contract *ERC8183Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC8183 creates a new instance of ERC8183, bound to a specific deployed contract.
func NewERC8183(address common.Address, backend bind.ContractBackend) (*ERC8183, error) {
	contract, err := bindERC8183(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC8183{ERC8183Caller: ERC8183Caller{contract: contract}, ERC8183Transactor: ERC8183Transactor{contract: contract}, ERC8183Filterer: ERC8183Filterer{contract: contract}}, nil
}

// NewERC8183Caller creates a new read-only instance of ERC8183, bound to a specific deployed contract.
func NewERC8183Caller(address common.Address, caller bind.ContractCaller) (*ERC8183Caller, error) {
	contract, err := bindERC8183(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC8183Caller{contract: contract}, nil
}

// NewERC8183Transactor creates a new write-only instance of ERC8183, bound to a specific deployed contract.
func NewERC8183Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC8183Transactor, error) {
	contract, err := bindERC8183(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC8183Transactor{contract: contract}, nil
}

// NewERC8183Filterer creates a new log filterer instance of ERC8183, bound to a specific deployed contract.
func NewERC8183Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC8183Filterer, error) {
	contract, err := bindERC8183(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC8183Filterer{contract: contract}, nil
}

// bindERC8183 binds a generic wrapper to an already deployed contract.
func bindERC8183(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC8183MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC8183 *ERC8183Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC8183.Contract.ERC8183Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC8183 *ERC8183Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC8183.Contract.ERC8183Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC8183 *ERC8183Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC8183.Contract.ERC8183Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC8183 *ERC8183CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC8183.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC8183 *ERC8183TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC8183.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC8183 *ERC8183TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC8183.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ERC8183 *ERC8183Caller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ERC8183 *ERC8183Session) ADMINROLE() ([32]byte, error) {
	return _ERC8183.Contract.ADMINROLE(&_ERC8183.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ERC8183 *ERC8183CallerSession) ADMINROLE() ([32]byte, error) {
	return _ERC8183.Contract.ADMINROLE(&_ERC8183.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC8183 *ERC8183Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC8183 *ERC8183Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC8183.Contract.DEFAULTADMINROLE(&_ERC8183.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC8183 *ERC8183CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC8183.Contract.DEFAULTADMINROLE(&_ERC8183.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ERC8183 *ERC8183Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ERC8183 *ERC8183Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _ERC8183.Contract.UPGRADEINTERFACEVERSION(&_ERC8183.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ERC8183 *ERC8183CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ERC8183.Contract.UPGRADEINTERFACEVERSION(&_ERC8183.CallOpts)
}

// EvaluatorFeeBP is a free data retrieval call binding the contract method 0x2f0e31f4.
//
// Solidity: function evaluatorFeeBP() view returns(uint256)
func (_ERC8183 *ERC8183Caller) EvaluatorFeeBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "evaluatorFeeBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EvaluatorFeeBP is a free data retrieval call binding the contract method 0x2f0e31f4.
//
// Solidity: function evaluatorFeeBP() view returns(uint256)
func (_ERC8183 *ERC8183Session) EvaluatorFeeBP() (*big.Int, error) {
	return _ERC8183.Contract.EvaluatorFeeBP(&_ERC8183.CallOpts)
}

// EvaluatorFeeBP is a free data retrieval call binding the contract method 0x2f0e31f4.
//
// Solidity: function evaluatorFeeBP() view returns(uint256)
func (_ERC8183 *ERC8183CallerSession) EvaluatorFeeBP() (*big.Int, error) {
	return _ERC8183.Contract.EvaluatorFeeBP(&_ERC8183.CallOpts)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,address,address,string,uint256,uint256,uint8,address,address,uint256))
func (_ERC8183 *ERC8183Caller) GetJob(opts *bind.CallOpts, jobId *big.Int) (AgenticCommerceJob, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "getJob", jobId)

	if err != nil {
		return *new(AgenticCommerceJob), err
	}

	out0 := *abi.ConvertType(out[0], new(AgenticCommerceJob)).(*AgenticCommerceJob)

	return out0, err

}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,address,address,string,uint256,uint256,uint8,address,address,uint256))
func (_ERC8183 *ERC8183Session) GetJob(jobId *big.Int) (AgenticCommerceJob, error) {
	return _ERC8183.Contract.GetJob(&_ERC8183.CallOpts, jobId)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,address,address,string,uint256,uint256,uint8,address,address,uint256))
func (_ERC8183 *ERC8183CallerSession) GetJob(jobId *big.Int) (AgenticCommerceJob, error) {
	return _ERC8183.Contract.GetJob(&_ERC8183.CallOpts, jobId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC8183 *ERC8183Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC8183 *ERC8183Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC8183.Contract.GetRoleAdmin(&_ERC8183.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC8183 *ERC8183CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC8183.Contract.GetRoleAdmin(&_ERC8183.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC8183 *ERC8183Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC8183 *ERC8183Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC8183.Contract.HasRole(&_ERC8183.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC8183 *ERC8183CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC8183.Contract.HasRole(&_ERC8183.CallOpts, role, account)
}

// JobCounter is a free data retrieval call binding the contract method 0x50355d76.
//
// Solidity: function jobCounter() view returns(uint256)
func (_ERC8183 *ERC8183Caller) JobCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "jobCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JobCounter is a free data retrieval call binding the contract method 0x50355d76.
//
// Solidity: function jobCounter() view returns(uint256)
func (_ERC8183 *ERC8183Session) JobCounter() (*big.Int, error) {
	return _ERC8183.Contract.JobCounter(&_ERC8183.CallOpts)
}

// JobCounter is a free data retrieval call binding the contract method 0x50355d76.
//
// Solidity: function jobCounter() view returns(uint256)
func (_ERC8183 *ERC8183CallerSession) JobCounter() (*big.Int, error) {
	return _ERC8183.Contract.JobCounter(&_ERC8183.CallOpts)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, address client, address provider, address evaluator, string description, uint256 budget, uint256 expiredAt, uint8 status, address hook, address paymentToken, uint256 providerAgentId)
func (_ERC8183 *ERC8183Caller) Jobs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              *big.Int
	Client          common.Address
	Provider        common.Address
	Evaluator       common.Address
	Description     string
	Budget          *big.Int
	ExpiredAt       *big.Int
	Status          uint8
	Hook            common.Address
	PaymentToken    common.Address
	ProviderAgentId *big.Int
}, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "jobs", arg0)

	outstruct := new(struct {
		Id              *big.Int
		Client          common.Address
		Provider        common.Address
		Evaluator       common.Address
		Description     string
		Budget          *big.Int
		ExpiredAt       *big.Int
		Status          uint8
		Hook            common.Address
		PaymentToken    common.Address
		ProviderAgentId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Client = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Provider = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Evaluator = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Description = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Budget = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ExpiredAt = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.Hook = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.PaymentToken = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.ProviderAgentId = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, address client, address provider, address evaluator, string description, uint256 budget, uint256 expiredAt, uint8 status, address hook, address paymentToken, uint256 providerAgentId)
func (_ERC8183 *ERC8183Session) Jobs(arg0 *big.Int) (struct {
	Id              *big.Int
	Client          common.Address
	Provider        common.Address
	Evaluator       common.Address
	Description     string
	Budget          *big.Int
	ExpiredAt       *big.Int
	Status          uint8
	Hook            common.Address
	PaymentToken    common.Address
	ProviderAgentId *big.Int
}, error) {
	return _ERC8183.Contract.Jobs(&_ERC8183.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, address client, address provider, address evaluator, string description, uint256 budget, uint256 expiredAt, uint8 status, address hook, address paymentToken, uint256 providerAgentId)
func (_ERC8183 *ERC8183CallerSession) Jobs(arg0 *big.Int) (struct {
	Id              *big.Int
	Client          common.Address
	Provider        common.Address
	Evaluator       common.Address
	Description     string
	Budget          *big.Int
	ExpiredAt       *big.Int
	Status          uint8
	Hook            common.Address
	PaymentToken    common.Address
	ProviderAgentId *big.Int
}, error) {
	return _ERC8183.Contract.Jobs(&_ERC8183.CallOpts, arg0)
}

// PlatformFeeBP is a free data retrieval call binding the contract method 0xff96092a.
//
// Solidity: function platformFeeBP() view returns(uint256)
func (_ERC8183 *ERC8183Caller) PlatformFeeBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "platformFeeBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFeeBP is a free data retrieval call binding the contract method 0xff96092a.
//
// Solidity: function platformFeeBP() view returns(uint256)
func (_ERC8183 *ERC8183Session) PlatformFeeBP() (*big.Int, error) {
	return _ERC8183.Contract.PlatformFeeBP(&_ERC8183.CallOpts)
}

// PlatformFeeBP is a free data retrieval call binding the contract method 0xff96092a.
//
// Solidity: function platformFeeBP() view returns(uint256)
func (_ERC8183 *ERC8183CallerSession) PlatformFeeBP() (*big.Int, error) {
	return _ERC8183.Contract.PlatformFeeBP(&_ERC8183.CallOpts)
}

// PlatformTreasury is a free data retrieval call binding the contract method 0xe138818c.
//
// Solidity: function platformTreasury() view returns(address)
func (_ERC8183 *ERC8183Caller) PlatformTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "platformTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformTreasury is a free data retrieval call binding the contract method 0xe138818c.
//
// Solidity: function platformTreasury() view returns(address)
func (_ERC8183 *ERC8183Session) PlatformTreasury() (common.Address, error) {
	return _ERC8183.Contract.PlatformTreasury(&_ERC8183.CallOpts)
}

// PlatformTreasury is a free data retrieval call binding the contract method 0xe138818c.
//
// Solidity: function platformTreasury() view returns(address)
func (_ERC8183 *ERC8183CallerSession) PlatformTreasury() (common.Address, error) {
	return _ERC8183.Contract.PlatformTreasury(&_ERC8183.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC8183 *ERC8183Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC8183 *ERC8183Session) ProxiableUUID() ([32]byte, error) {
	return _ERC8183.Contract.ProxiableUUID(&_ERC8183.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC8183 *ERC8183CallerSession) ProxiableUUID() ([32]byte, error) {
	return _ERC8183.Contract.ProxiableUUID(&_ERC8183.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC8183 *ERC8183Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC8183 *ERC8183Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC8183.Contract.SupportsInterface(&_ERC8183.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC8183 *ERC8183CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC8183.Contract.SupportsInterface(&_ERC8183.CallOpts, interfaceId)
}

// WhitelistedHooks is a free data retrieval call binding the contract method 0x6d3b96c3.
//
// Solidity: function whitelistedHooks(address ) view returns(bool)
func (_ERC8183 *ERC8183Caller) WhitelistedHooks(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC8183.contract.Call(opts, &out, "whitelistedHooks", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedHooks is a free data retrieval call binding the contract method 0x6d3b96c3.
//
// Solidity: function whitelistedHooks(address ) view returns(bool)
func (_ERC8183 *ERC8183Session) WhitelistedHooks(arg0 common.Address) (bool, error) {
	return _ERC8183.Contract.WhitelistedHooks(&_ERC8183.CallOpts, arg0)
}

// WhitelistedHooks is a free data retrieval call binding the contract method 0x6d3b96c3.
//
// Solidity: function whitelistedHooks(address ) view returns(bool)
func (_ERC8183 *ERC8183CallerSession) WhitelistedHooks(arg0 common.Address) (bool, error) {
	return _ERC8183.Contract.WhitelistedHooks(&_ERC8183.CallOpts, arg0)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 jobId) returns()
func (_ERC8183 *ERC8183Transactor) ClaimRefund(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "claimRefund", jobId)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 jobId) returns()
func (_ERC8183 *ERC8183Session) ClaimRefund(jobId *big.Int) (*types.Transaction, error) {
	return _ERC8183.Contract.ClaimRefund(&_ERC8183.TransactOpts, jobId)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 jobId) returns()
func (_ERC8183 *ERC8183TransactorSession) ClaimRefund(jobId *big.Int) (*types.Transaction, error) {
	return _ERC8183.Contract.ClaimRefund(&_ERC8183.TransactOpts, jobId)
}

// Complete is a paid mutator transaction binding the contract method 0xd75bbdf3.
//
// Solidity: function complete(uint256 jobId, bytes32 reason, bytes optParams) returns()
func (_ERC8183 *ERC8183Transactor) Complete(opts *bind.TransactOpts, jobId *big.Int, reason [32]byte, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "complete", jobId, reason, optParams)
}

// Complete is a paid mutator transaction binding the contract method 0xd75bbdf3.
//
// Solidity: function complete(uint256 jobId, bytes32 reason, bytes optParams) returns()
func (_ERC8183 *ERC8183Session) Complete(jobId *big.Int, reason [32]byte, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.Complete(&_ERC8183.TransactOpts, jobId, reason, optParams)
}

// Complete is a paid mutator transaction binding the contract method 0xd75bbdf3.
//
// Solidity: function complete(uint256 jobId, bytes32 reason, bytes optParams) returns()
func (_ERC8183 *ERC8183TransactorSession) Complete(jobId *big.Int, reason [32]byte, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.Complete(&_ERC8183.TransactOpts, jobId, reason, optParams)
}

// CreateJob is a paid mutator transaction binding the contract method 0x1c226749.
//
// Solidity: function createJob(address provider, address evaluator, uint256 expiredAt, string description, address hook, uint256 providerAgentId) returns(uint256)
func (_ERC8183 *ERC8183Transactor) CreateJob(opts *bind.TransactOpts, provider common.Address, evaluator common.Address, expiredAt *big.Int, description string, hook common.Address, providerAgentId *big.Int) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "createJob", provider, evaluator, expiredAt, description, hook, providerAgentId)
}

// CreateJob is a paid mutator transaction binding the contract method 0x1c226749.
//
// Solidity: function createJob(address provider, address evaluator, uint256 expiredAt, string description, address hook, uint256 providerAgentId) returns(uint256)
func (_ERC8183 *ERC8183Session) CreateJob(provider common.Address, evaluator common.Address, expiredAt *big.Int, description string, hook common.Address, providerAgentId *big.Int) (*types.Transaction, error) {
	return _ERC8183.Contract.CreateJob(&_ERC8183.TransactOpts, provider, evaluator, expiredAt, description, hook, providerAgentId)
}

// CreateJob is a paid mutator transaction binding the contract method 0x1c226749.
//
// Solidity: function createJob(address provider, address evaluator, uint256 expiredAt, string description, address hook, uint256 providerAgentId) returns(uint256)
func (_ERC8183 *ERC8183TransactorSession) CreateJob(provider common.Address, evaluator common.Address, expiredAt *big.Int, description string, hook common.Address, providerAgentId *big.Int) (*types.Transaction, error) {
	return _ERC8183.Contract.CreateJob(&_ERC8183.TransactOpts, provider, evaluator, expiredAt, description, hook, providerAgentId)
}

// Fund is a paid mutator transaction binding the contract method 0xd2e13f50.
//
// Solidity: function fund(uint256 jobId, uint256 expectedBudget, bytes optParams) returns()
func (_ERC8183 *ERC8183Transactor) Fund(opts *bind.TransactOpts, jobId *big.Int, expectedBudget *big.Int, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "fund", jobId, expectedBudget, optParams)
}

// Fund is a paid mutator transaction binding the contract method 0xd2e13f50.
//
// Solidity: function fund(uint256 jobId, uint256 expectedBudget, bytes optParams) returns()
func (_ERC8183 *ERC8183Session) Fund(jobId *big.Int, expectedBudget *big.Int, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.Fund(&_ERC8183.TransactOpts, jobId, expectedBudget, optParams)
}

// Fund is a paid mutator transaction binding the contract method 0xd2e13f50.
//
// Solidity: function fund(uint256 jobId, uint256 expectedBudget, bytes optParams) returns()
func (_ERC8183 *ERC8183TransactorSession) Fund(jobId *big.Int, expectedBudget *big.Int, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.Fund(&_ERC8183.TransactOpts, jobId, expectedBudget, optParams)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC8183 *ERC8183Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC8183 *ERC8183Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC8183.Contract.GrantRole(&_ERC8183.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC8183 *ERC8183TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC8183.Contract.GrantRole(&_ERC8183.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address treasury_) returns()
func (_ERC8183 *ERC8183Transactor) Initialize(opts *bind.TransactOpts, treasury_ common.Address) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "initialize", treasury_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address treasury_) returns()
func (_ERC8183 *ERC8183Session) Initialize(treasury_ common.Address) (*types.Transaction, error) {
	return _ERC8183.Contract.Initialize(&_ERC8183.TransactOpts, treasury_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address treasury_) returns()
func (_ERC8183 *ERC8183TransactorSession) Initialize(treasury_ common.Address) (*types.Transaction, error) {
	return _ERC8183.Contract.Initialize(&_ERC8183.TransactOpts, treasury_)
}

// Reject is a paid mutator transaction binding the contract method 0x41dd26f5.
//
// Solidity: function reject(uint256 jobId, bytes32 reason, bytes optParams) returns()
func (_ERC8183 *ERC8183Transactor) Reject(opts *bind.TransactOpts, jobId *big.Int, reason [32]byte, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "reject", jobId, reason, optParams)
}

// Reject is a paid mutator transaction binding the contract method 0x41dd26f5.
//
// Solidity: function reject(uint256 jobId, bytes32 reason, bytes optParams) returns()
func (_ERC8183 *ERC8183Session) Reject(jobId *big.Int, reason [32]byte, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.Reject(&_ERC8183.TransactOpts, jobId, reason, optParams)
}

// Reject is a paid mutator transaction binding the contract method 0x41dd26f5.
//
// Solidity: function reject(uint256 jobId, bytes32 reason, bytes optParams) returns()
func (_ERC8183 *ERC8183TransactorSession) Reject(jobId *big.Int, reason [32]byte, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.Reject(&_ERC8183.TransactOpts, jobId, reason, optParams)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ERC8183 *ERC8183Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ERC8183 *ERC8183Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ERC8183.Contract.RenounceRole(&_ERC8183.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ERC8183 *ERC8183TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ERC8183.Contract.RenounceRole(&_ERC8183.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC8183 *ERC8183Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC8183 *ERC8183Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC8183.Contract.RevokeRole(&_ERC8183.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC8183 *ERC8183TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC8183.Contract.RevokeRole(&_ERC8183.TransactOpts, role, account)
}

// SetBudget is a paid mutator transaction binding the contract method 0xf3302b89.
//
// Solidity: function setBudget(uint256 jobId, address token, uint256 amount, bytes optParams) returns()
func (_ERC8183 *ERC8183Transactor) SetBudget(opts *bind.TransactOpts, jobId *big.Int, token common.Address, amount *big.Int, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "setBudget", jobId, token, amount, optParams)
}

// SetBudget is a paid mutator transaction binding the contract method 0xf3302b89.
//
// Solidity: function setBudget(uint256 jobId, address token, uint256 amount, bytes optParams) returns()
func (_ERC8183 *ERC8183Session) SetBudget(jobId *big.Int, token common.Address, amount *big.Int, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.SetBudget(&_ERC8183.TransactOpts, jobId, token, amount, optParams)
}

// SetBudget is a paid mutator transaction binding the contract method 0xf3302b89.
//
// Solidity: function setBudget(uint256 jobId, address token, uint256 amount, bytes optParams) returns()
func (_ERC8183 *ERC8183TransactorSession) SetBudget(jobId *big.Int, token common.Address, amount *big.Int, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.SetBudget(&_ERC8183.TransactOpts, jobId, token, amount, optParams)
}

// SetEvaluatorFee is a paid mutator transaction binding the contract method 0x84f15090.
//
// Solidity: function setEvaluatorFee(uint256 feeBP_) returns()
func (_ERC8183 *ERC8183Transactor) SetEvaluatorFee(opts *bind.TransactOpts, feeBP_ *big.Int) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "setEvaluatorFee", feeBP_)
}

// SetEvaluatorFee is a paid mutator transaction binding the contract method 0x84f15090.
//
// Solidity: function setEvaluatorFee(uint256 feeBP_) returns()
func (_ERC8183 *ERC8183Session) SetEvaluatorFee(feeBP_ *big.Int) (*types.Transaction, error) {
	return _ERC8183.Contract.SetEvaluatorFee(&_ERC8183.TransactOpts, feeBP_)
}

// SetEvaluatorFee is a paid mutator transaction binding the contract method 0x84f15090.
//
// Solidity: function setEvaluatorFee(uint256 feeBP_) returns()
func (_ERC8183 *ERC8183TransactorSession) SetEvaluatorFee(feeBP_ *big.Int) (*types.Transaction, error) {
	return _ERC8183.Contract.SetEvaluatorFee(&_ERC8183.TransactOpts, feeBP_)
}

// SetHookWhitelist is a paid mutator transaction binding the contract method 0xce79eb60.
//
// Solidity: function setHookWhitelist(address hook, bool status) returns()
func (_ERC8183 *ERC8183Transactor) SetHookWhitelist(opts *bind.TransactOpts, hook common.Address, status bool) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "setHookWhitelist", hook, status)
}

// SetHookWhitelist is a paid mutator transaction binding the contract method 0xce79eb60.
//
// Solidity: function setHookWhitelist(address hook, bool status) returns()
func (_ERC8183 *ERC8183Session) SetHookWhitelist(hook common.Address, status bool) (*types.Transaction, error) {
	return _ERC8183.Contract.SetHookWhitelist(&_ERC8183.TransactOpts, hook, status)
}

// SetHookWhitelist is a paid mutator transaction binding the contract method 0xce79eb60.
//
// Solidity: function setHookWhitelist(address hook, bool status) returns()
func (_ERC8183 *ERC8183TransactorSession) SetHookWhitelist(hook common.Address, status bool) (*types.Transaction, error) {
	return _ERC8183.Contract.SetHookWhitelist(&_ERC8183.TransactOpts, hook, status)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0xb4d884f6.
//
// Solidity: function setPlatformFee(uint256 feeBP_, address treasury_) returns()
func (_ERC8183 *ERC8183Transactor) SetPlatformFee(opts *bind.TransactOpts, feeBP_ *big.Int, treasury_ common.Address) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "setPlatformFee", feeBP_, treasury_)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0xb4d884f6.
//
// Solidity: function setPlatformFee(uint256 feeBP_, address treasury_) returns()
func (_ERC8183 *ERC8183Session) SetPlatformFee(feeBP_ *big.Int, treasury_ common.Address) (*types.Transaction, error) {
	return _ERC8183.Contract.SetPlatformFee(&_ERC8183.TransactOpts, feeBP_, treasury_)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0xb4d884f6.
//
// Solidity: function setPlatformFee(uint256 feeBP_, address treasury_) returns()
func (_ERC8183 *ERC8183TransactorSession) SetPlatformFee(feeBP_ *big.Int, treasury_ common.Address) (*types.Transaction, error) {
	return _ERC8183.Contract.SetPlatformFee(&_ERC8183.TransactOpts, feeBP_, treasury_)
}

// SetProvider is a paid mutator transaction binding the contract method 0xfac3246c.
//
// Solidity: function setProvider(uint256 jobId, address provider_, uint256 agentId) returns()
func (_ERC8183 *ERC8183Transactor) SetProvider(opts *bind.TransactOpts, jobId *big.Int, provider_ common.Address, agentId *big.Int) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "setProvider", jobId, provider_, agentId)
}

// SetProvider is a paid mutator transaction binding the contract method 0xfac3246c.
//
// Solidity: function setProvider(uint256 jobId, address provider_, uint256 agentId) returns()
func (_ERC8183 *ERC8183Session) SetProvider(jobId *big.Int, provider_ common.Address, agentId *big.Int) (*types.Transaction, error) {
	return _ERC8183.Contract.SetProvider(&_ERC8183.TransactOpts, jobId, provider_, agentId)
}

// SetProvider is a paid mutator transaction binding the contract method 0xfac3246c.
//
// Solidity: function setProvider(uint256 jobId, address provider_, uint256 agentId) returns()
func (_ERC8183 *ERC8183TransactorSession) SetProvider(jobId *big.Int, provider_ common.Address, agentId *big.Int) (*types.Transaction, error) {
	return _ERC8183.Contract.SetProvider(&_ERC8183.TransactOpts, jobId, provider_, agentId)
}

// Submit is a paid mutator transaction binding the contract method 0x9e63798d.
//
// Solidity: function submit(uint256 jobId, bytes32 deliverable, bytes optParams) returns()
func (_ERC8183 *ERC8183Transactor) Submit(opts *bind.TransactOpts, jobId *big.Int, deliverable [32]byte, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "submit", jobId, deliverable, optParams)
}

// Submit is a paid mutator transaction binding the contract method 0x9e63798d.
//
// Solidity: function submit(uint256 jobId, bytes32 deliverable, bytes optParams) returns()
func (_ERC8183 *ERC8183Session) Submit(jobId *big.Int, deliverable [32]byte, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.Submit(&_ERC8183.TransactOpts, jobId, deliverable, optParams)
}

// Submit is a paid mutator transaction binding the contract method 0x9e63798d.
//
// Solidity: function submit(uint256 jobId, bytes32 deliverable, bytes optParams) returns()
func (_ERC8183 *ERC8183TransactorSession) Submit(jobId *big.Int, deliverable [32]byte, optParams []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.Submit(&_ERC8183.TransactOpts, jobId, deliverable, optParams)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC8183 *ERC8183Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC8183.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC8183 *ERC8183Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.UpgradeToAndCall(&_ERC8183.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC8183 *ERC8183TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC8183.Contract.UpgradeToAndCall(&_ERC8183.TransactOpts, newImplementation, data)
}

// ERC8183BudgetSetIterator is returned from FilterBudgetSet and is used to iterate over the raw logs and unpacked data for BudgetSet events raised by the ERC8183 contract.
type ERC8183BudgetSetIterator struct {
	Event *ERC8183BudgetSet // Event containing the contract specifics and raw log

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
func (it *ERC8183BudgetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183BudgetSet)
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
		it.Event = new(ERC8183BudgetSet)
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
func (it *ERC8183BudgetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183BudgetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183BudgetSet represents a BudgetSet event raised by the ERC8183 contract.
type ERC8183BudgetSet struct {
	JobId  *big.Int
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBudgetSet is a free log retrieval operation binding the contract event 0x74f37e24047ef5ad02b6aac54caaf20e06c51585a911183b2c6f2db0d8d896cc.
//
// Solidity: event BudgetSet(uint256 indexed jobId, address indexed token, uint256 amount)
func (_ERC8183 *ERC8183Filterer) FilterBudgetSet(opts *bind.FilterOpts, jobId []*big.Int, token []common.Address) (*ERC8183BudgetSetIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "BudgetSet", jobIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183BudgetSetIterator{contract: _ERC8183.contract, event: "BudgetSet", logs: logs, sub: sub}, nil
}

// WatchBudgetSet is a free log subscription operation binding the contract event 0x74f37e24047ef5ad02b6aac54caaf20e06c51585a911183b2c6f2db0d8d896cc.
//
// Solidity: event BudgetSet(uint256 indexed jobId, address indexed token, uint256 amount)
func (_ERC8183 *ERC8183Filterer) WatchBudgetSet(opts *bind.WatchOpts, sink chan<- *ERC8183BudgetSet, jobId []*big.Int, token []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "BudgetSet", jobIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183BudgetSet)
				if err := _ERC8183.contract.UnpackLog(event, "BudgetSet", log); err != nil {
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

// ParseBudgetSet is a log parse operation binding the contract event 0x74f37e24047ef5ad02b6aac54caaf20e06c51585a911183b2c6f2db0d8d896cc.
//
// Solidity: event BudgetSet(uint256 indexed jobId, address indexed token, uint256 amount)
func (_ERC8183 *ERC8183Filterer) ParseBudgetSet(log types.Log) (*ERC8183BudgetSet, error) {
	event := new(ERC8183BudgetSet)
	if err := _ERC8183.contract.UnpackLog(event, "BudgetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183EvaluatorFeePaidIterator is returned from FilterEvaluatorFeePaid and is used to iterate over the raw logs and unpacked data for EvaluatorFeePaid events raised by the ERC8183 contract.
type ERC8183EvaluatorFeePaidIterator struct {
	Event *ERC8183EvaluatorFeePaid // Event containing the contract specifics and raw log

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
func (it *ERC8183EvaluatorFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183EvaluatorFeePaid)
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
		it.Event = new(ERC8183EvaluatorFeePaid)
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
func (it *ERC8183EvaluatorFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183EvaluatorFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183EvaluatorFeePaid represents a EvaluatorFeePaid event raised by the ERC8183 contract.
type ERC8183EvaluatorFeePaid struct {
	JobId     *big.Int
	Evaluator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvaluatorFeePaid is a free log retrieval operation binding the contract event 0x253dd534010ac976fa263caa123bae79b9c50292adf7ce67bdc5ec309f784e61.
//
// Solidity: event EvaluatorFeePaid(uint256 indexed jobId, address indexed evaluator, uint256 amount)
func (_ERC8183 *ERC8183Filterer) FilterEvaluatorFeePaid(opts *bind.FilterOpts, jobId []*big.Int, evaluator []common.Address) (*ERC8183EvaluatorFeePaidIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var evaluatorRule []interface{}
	for _, evaluatorItem := range evaluator {
		evaluatorRule = append(evaluatorRule, evaluatorItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "EvaluatorFeePaid", jobIdRule, evaluatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183EvaluatorFeePaidIterator{contract: _ERC8183.contract, event: "EvaluatorFeePaid", logs: logs, sub: sub}, nil
}

// WatchEvaluatorFeePaid is a free log subscription operation binding the contract event 0x253dd534010ac976fa263caa123bae79b9c50292adf7ce67bdc5ec309f784e61.
//
// Solidity: event EvaluatorFeePaid(uint256 indexed jobId, address indexed evaluator, uint256 amount)
func (_ERC8183 *ERC8183Filterer) WatchEvaluatorFeePaid(opts *bind.WatchOpts, sink chan<- *ERC8183EvaluatorFeePaid, jobId []*big.Int, evaluator []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var evaluatorRule []interface{}
	for _, evaluatorItem := range evaluator {
		evaluatorRule = append(evaluatorRule, evaluatorItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "EvaluatorFeePaid", jobIdRule, evaluatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183EvaluatorFeePaid)
				if err := _ERC8183.contract.UnpackLog(event, "EvaluatorFeePaid", log); err != nil {
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

// ParseEvaluatorFeePaid is a log parse operation binding the contract event 0x253dd534010ac976fa263caa123bae79b9c50292adf7ce67bdc5ec309f784e61.
//
// Solidity: event EvaluatorFeePaid(uint256 indexed jobId, address indexed evaluator, uint256 amount)
func (_ERC8183 *ERC8183Filterer) ParseEvaluatorFeePaid(log types.Log) (*ERC8183EvaluatorFeePaid, error) {
	event := new(ERC8183EvaluatorFeePaid)
	if err := _ERC8183.contract.UnpackLog(event, "EvaluatorFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183EvaluatorFeeSetIterator is returned from FilterEvaluatorFeeSet and is used to iterate over the raw logs and unpacked data for EvaluatorFeeSet events raised by the ERC8183 contract.
type ERC8183EvaluatorFeeSetIterator struct {
	Event *ERC8183EvaluatorFeeSet // Event containing the contract specifics and raw log

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
func (it *ERC8183EvaluatorFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183EvaluatorFeeSet)
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
		it.Event = new(ERC8183EvaluatorFeeSet)
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
func (it *ERC8183EvaluatorFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183EvaluatorFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183EvaluatorFeeSet represents a EvaluatorFeeSet event raised by the ERC8183 contract.
type ERC8183EvaluatorFeeSet struct {
	FeeBP *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEvaluatorFeeSet is a free log retrieval operation binding the contract event 0x8d3c1e4f1243a044dac6f85b2c1043fe0965504a5d1331bca64534c4b45d560b.
//
// Solidity: event EvaluatorFeeSet(uint256 feeBP)
func (_ERC8183 *ERC8183Filterer) FilterEvaluatorFeeSet(opts *bind.FilterOpts) (*ERC8183EvaluatorFeeSetIterator, error) {

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "EvaluatorFeeSet")
	if err != nil {
		return nil, err
	}
	return &ERC8183EvaluatorFeeSetIterator{contract: _ERC8183.contract, event: "EvaluatorFeeSet", logs: logs, sub: sub}, nil
}

// WatchEvaluatorFeeSet is a free log subscription operation binding the contract event 0x8d3c1e4f1243a044dac6f85b2c1043fe0965504a5d1331bca64534c4b45d560b.
//
// Solidity: event EvaluatorFeeSet(uint256 feeBP)
func (_ERC8183 *ERC8183Filterer) WatchEvaluatorFeeSet(opts *bind.WatchOpts, sink chan<- *ERC8183EvaluatorFeeSet) (event.Subscription, error) {

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "EvaluatorFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183EvaluatorFeeSet)
				if err := _ERC8183.contract.UnpackLog(event, "EvaluatorFeeSet", log); err != nil {
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

// ParseEvaluatorFeeSet is a log parse operation binding the contract event 0x8d3c1e4f1243a044dac6f85b2c1043fe0965504a5d1331bca64534c4b45d560b.
//
// Solidity: event EvaluatorFeeSet(uint256 feeBP)
func (_ERC8183 *ERC8183Filterer) ParseEvaluatorFeeSet(log types.Log) (*ERC8183EvaluatorFeeSet, error) {
	event := new(ERC8183EvaluatorFeeSet)
	if err := _ERC8183.contract.UnpackLog(event, "EvaluatorFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183HookWhitelistUpdatedIterator is returned from FilterHookWhitelistUpdated and is used to iterate over the raw logs and unpacked data for HookWhitelistUpdated events raised by the ERC8183 contract.
type ERC8183HookWhitelistUpdatedIterator struct {
	Event *ERC8183HookWhitelistUpdated // Event containing the contract specifics and raw log

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
func (it *ERC8183HookWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183HookWhitelistUpdated)
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
		it.Event = new(ERC8183HookWhitelistUpdated)
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
func (it *ERC8183HookWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183HookWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183HookWhitelistUpdated represents a HookWhitelistUpdated event raised by the ERC8183 contract.
type ERC8183HookWhitelistUpdated struct {
	Hook   common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHookWhitelistUpdated is a free log retrieval operation binding the contract event 0x7ee54953080e392a475a25b6acacb85417ca4e1953293c90934233ca13612510.
//
// Solidity: event HookWhitelistUpdated(address indexed hook, bool status)
func (_ERC8183 *ERC8183Filterer) FilterHookWhitelistUpdated(opts *bind.FilterOpts, hook []common.Address) (*ERC8183HookWhitelistUpdatedIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "HookWhitelistUpdated", hookRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183HookWhitelistUpdatedIterator{contract: _ERC8183.contract, event: "HookWhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchHookWhitelistUpdated is a free log subscription operation binding the contract event 0x7ee54953080e392a475a25b6acacb85417ca4e1953293c90934233ca13612510.
//
// Solidity: event HookWhitelistUpdated(address indexed hook, bool status)
func (_ERC8183 *ERC8183Filterer) WatchHookWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *ERC8183HookWhitelistUpdated, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "HookWhitelistUpdated", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183HookWhitelistUpdated)
				if err := _ERC8183.contract.UnpackLog(event, "HookWhitelistUpdated", log); err != nil {
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

// ParseHookWhitelistUpdated is a log parse operation binding the contract event 0x7ee54953080e392a475a25b6acacb85417ca4e1953293c90934233ca13612510.
//
// Solidity: event HookWhitelistUpdated(address indexed hook, bool status)
func (_ERC8183 *ERC8183Filterer) ParseHookWhitelistUpdated(log types.Log) (*ERC8183HookWhitelistUpdated, error) {
	event := new(ERC8183HookWhitelistUpdated)
	if err := _ERC8183.contract.UnpackLog(event, "HookWhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC8183 contract.
type ERC8183InitializedIterator struct {
	Event *ERC8183Initialized // Event containing the contract specifics and raw log

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
func (it *ERC8183InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183Initialized)
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
		it.Event = new(ERC8183Initialized)
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
func (it *ERC8183InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183Initialized represents a Initialized event raised by the ERC8183 contract.
type ERC8183Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC8183 *ERC8183Filterer) FilterInitialized(opts *bind.FilterOpts) (*ERC8183InitializedIterator, error) {

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC8183InitializedIterator{contract: _ERC8183.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC8183 *ERC8183Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC8183Initialized) (event.Subscription, error) {

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183Initialized)
				if err := _ERC8183.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC8183 *ERC8183Filterer) ParseInitialized(log types.Log) (*ERC8183Initialized, error) {
	event := new(ERC8183Initialized)
	if err := _ERC8183.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183JobCompletedIterator is returned from FilterJobCompleted and is used to iterate over the raw logs and unpacked data for JobCompleted events raised by the ERC8183 contract.
type ERC8183JobCompletedIterator struct {
	Event *ERC8183JobCompleted // Event containing the contract specifics and raw log

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
func (it *ERC8183JobCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183JobCompleted)
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
		it.Event = new(ERC8183JobCompleted)
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
func (it *ERC8183JobCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183JobCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183JobCompleted represents a JobCompleted event raised by the ERC8183 contract.
type ERC8183JobCompleted struct {
	JobId     *big.Int
	Evaluator common.Address
	Reason    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJobCompleted is a free log retrieval operation binding the contract event 0x0fd54bd364fa9e67f17b091aefe930932c09fe7651cf5ad02c71a418f3341444.
//
// Solidity: event JobCompleted(uint256 indexed jobId, address indexed evaluator, bytes32 reason)
func (_ERC8183 *ERC8183Filterer) FilterJobCompleted(opts *bind.FilterOpts, jobId []*big.Int, evaluator []common.Address) (*ERC8183JobCompletedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var evaluatorRule []interface{}
	for _, evaluatorItem := range evaluator {
		evaluatorRule = append(evaluatorRule, evaluatorItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "JobCompleted", jobIdRule, evaluatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183JobCompletedIterator{contract: _ERC8183.contract, event: "JobCompleted", logs: logs, sub: sub}, nil
}

// WatchJobCompleted is a free log subscription operation binding the contract event 0x0fd54bd364fa9e67f17b091aefe930932c09fe7651cf5ad02c71a418f3341444.
//
// Solidity: event JobCompleted(uint256 indexed jobId, address indexed evaluator, bytes32 reason)
func (_ERC8183 *ERC8183Filterer) WatchJobCompleted(opts *bind.WatchOpts, sink chan<- *ERC8183JobCompleted, jobId []*big.Int, evaluator []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var evaluatorRule []interface{}
	for _, evaluatorItem := range evaluator {
		evaluatorRule = append(evaluatorRule, evaluatorItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "JobCompleted", jobIdRule, evaluatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183JobCompleted)
				if err := _ERC8183.contract.UnpackLog(event, "JobCompleted", log); err != nil {
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

// ParseJobCompleted is a log parse operation binding the contract event 0x0fd54bd364fa9e67f17b091aefe930932c09fe7651cf5ad02c71a418f3341444.
//
// Solidity: event JobCompleted(uint256 indexed jobId, address indexed evaluator, bytes32 reason)
func (_ERC8183 *ERC8183Filterer) ParseJobCompleted(log types.Log) (*ERC8183JobCompleted, error) {
	event := new(ERC8183JobCompleted)
	if err := _ERC8183.contract.UnpackLog(event, "JobCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183JobCreatedIterator is returned from FilterJobCreated and is used to iterate over the raw logs and unpacked data for JobCreated events raised by the ERC8183 contract.
type ERC8183JobCreatedIterator struct {
	Event *ERC8183JobCreated // Event containing the contract specifics and raw log

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
func (it *ERC8183JobCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183JobCreated)
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
		it.Event = new(ERC8183JobCreated)
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
func (it *ERC8183JobCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183JobCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183JobCreated represents a JobCreated event raised by the ERC8183 contract.
type ERC8183JobCreated struct {
	JobId     *big.Int
	Client    common.Address
	Provider  common.Address
	Evaluator common.Address
	ExpiredAt *big.Int
	Hook      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJobCreated is a free log retrieval operation binding the contract event 0xb0f0239bfdd96453e24733e18bfc24b70d8fadf123dd977473518dd577ee79b9.
//
// Solidity: event JobCreated(uint256 indexed jobId, address indexed client, address indexed provider, address evaluator, uint256 expiredAt, address hook)
func (_ERC8183 *ERC8183Filterer) FilterJobCreated(opts *bind.FilterOpts, jobId []*big.Int, client []common.Address, provider []common.Address) (*ERC8183JobCreatedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "JobCreated", jobIdRule, clientRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183JobCreatedIterator{contract: _ERC8183.contract, event: "JobCreated", logs: logs, sub: sub}, nil
}

// WatchJobCreated is a free log subscription operation binding the contract event 0xb0f0239bfdd96453e24733e18bfc24b70d8fadf123dd977473518dd577ee79b9.
//
// Solidity: event JobCreated(uint256 indexed jobId, address indexed client, address indexed provider, address evaluator, uint256 expiredAt, address hook)
func (_ERC8183 *ERC8183Filterer) WatchJobCreated(opts *bind.WatchOpts, sink chan<- *ERC8183JobCreated, jobId []*big.Int, client []common.Address, provider []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "JobCreated", jobIdRule, clientRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183JobCreated)
				if err := _ERC8183.contract.UnpackLog(event, "JobCreated", log); err != nil {
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

// ParseJobCreated is a log parse operation binding the contract event 0xb0f0239bfdd96453e24733e18bfc24b70d8fadf123dd977473518dd577ee79b9.
//
// Solidity: event JobCreated(uint256 indexed jobId, address indexed client, address indexed provider, address evaluator, uint256 expiredAt, address hook)
func (_ERC8183 *ERC8183Filterer) ParseJobCreated(log types.Log) (*ERC8183JobCreated, error) {
	event := new(ERC8183JobCreated)
	if err := _ERC8183.contract.UnpackLog(event, "JobCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183JobExpiredIterator is returned from FilterJobExpired and is used to iterate over the raw logs and unpacked data for JobExpired events raised by the ERC8183 contract.
type ERC8183JobExpiredIterator struct {
	Event *ERC8183JobExpired // Event containing the contract specifics and raw log

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
func (it *ERC8183JobExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183JobExpired)
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
		it.Event = new(ERC8183JobExpired)
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
func (it *ERC8183JobExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183JobExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183JobExpired represents a JobExpired event raised by the ERC8183 contract.
type ERC8183JobExpired struct {
	JobId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterJobExpired is a free log retrieval operation binding the contract event 0x97237956f8810192811e2c3f273fd02c5d6295206fdd9c62e6fe2bfc19ba9232.
//
// Solidity: event JobExpired(uint256 indexed jobId)
func (_ERC8183 *ERC8183Filterer) FilterJobExpired(opts *bind.FilterOpts, jobId []*big.Int) (*ERC8183JobExpiredIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "JobExpired", jobIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183JobExpiredIterator{contract: _ERC8183.contract, event: "JobExpired", logs: logs, sub: sub}, nil
}

// WatchJobExpired is a free log subscription operation binding the contract event 0x97237956f8810192811e2c3f273fd02c5d6295206fdd9c62e6fe2bfc19ba9232.
//
// Solidity: event JobExpired(uint256 indexed jobId)
func (_ERC8183 *ERC8183Filterer) WatchJobExpired(opts *bind.WatchOpts, sink chan<- *ERC8183JobExpired, jobId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "JobExpired", jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183JobExpired)
				if err := _ERC8183.contract.UnpackLog(event, "JobExpired", log); err != nil {
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

// ParseJobExpired is a log parse operation binding the contract event 0x97237956f8810192811e2c3f273fd02c5d6295206fdd9c62e6fe2bfc19ba9232.
//
// Solidity: event JobExpired(uint256 indexed jobId)
func (_ERC8183 *ERC8183Filterer) ParseJobExpired(log types.Log) (*ERC8183JobExpired, error) {
	event := new(ERC8183JobExpired)
	if err := _ERC8183.contract.UnpackLog(event, "JobExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183JobFundedIterator is returned from FilterJobFunded and is used to iterate over the raw logs and unpacked data for JobFunded events raised by the ERC8183 contract.
type ERC8183JobFundedIterator struct {
	Event *ERC8183JobFunded // Event containing the contract specifics and raw log

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
func (it *ERC8183JobFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183JobFunded)
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
		it.Event = new(ERC8183JobFunded)
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
func (it *ERC8183JobFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183JobFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183JobFunded represents a JobFunded event raised by the ERC8183 contract.
type ERC8183JobFunded struct {
	JobId  *big.Int
	Client common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJobFunded is a free log retrieval operation binding the contract event 0xe3fbcc1ea1bdc559ec7f0347efde7655e58b5f45a30b0e4470a583c3ef5496b3.
//
// Solidity: event JobFunded(uint256 indexed jobId, address indexed client, uint256 amount)
func (_ERC8183 *ERC8183Filterer) FilterJobFunded(opts *bind.FilterOpts, jobId []*big.Int, client []common.Address) (*ERC8183JobFundedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "JobFunded", jobIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183JobFundedIterator{contract: _ERC8183.contract, event: "JobFunded", logs: logs, sub: sub}, nil
}

// WatchJobFunded is a free log subscription operation binding the contract event 0xe3fbcc1ea1bdc559ec7f0347efde7655e58b5f45a30b0e4470a583c3ef5496b3.
//
// Solidity: event JobFunded(uint256 indexed jobId, address indexed client, uint256 amount)
func (_ERC8183 *ERC8183Filterer) WatchJobFunded(opts *bind.WatchOpts, sink chan<- *ERC8183JobFunded, jobId []*big.Int, client []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "JobFunded", jobIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183JobFunded)
				if err := _ERC8183.contract.UnpackLog(event, "JobFunded", log); err != nil {
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

// ParseJobFunded is a log parse operation binding the contract event 0xe3fbcc1ea1bdc559ec7f0347efde7655e58b5f45a30b0e4470a583c3ef5496b3.
//
// Solidity: event JobFunded(uint256 indexed jobId, address indexed client, uint256 amount)
func (_ERC8183 *ERC8183Filterer) ParseJobFunded(log types.Log) (*ERC8183JobFunded, error) {
	event := new(ERC8183JobFunded)
	if err := _ERC8183.contract.UnpackLog(event, "JobFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183JobRejectedIterator is returned from FilterJobRejected and is used to iterate over the raw logs and unpacked data for JobRejected events raised by the ERC8183 contract.
type ERC8183JobRejectedIterator struct {
	Event *ERC8183JobRejected // Event containing the contract specifics and raw log

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
func (it *ERC8183JobRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183JobRejected)
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
		it.Event = new(ERC8183JobRejected)
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
func (it *ERC8183JobRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183JobRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183JobRejected represents a JobRejected event raised by the ERC8183 contract.
type ERC8183JobRejected struct {
	JobId    *big.Int
	Rejector common.Address
	Reason   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterJobRejected is a free log retrieval operation binding the contract event 0xae7362b1af91f4492868987b9c73990d780060811551b58728fbe96fd1bab275.
//
// Solidity: event JobRejected(uint256 indexed jobId, address indexed rejector, bytes32 reason)
func (_ERC8183 *ERC8183Filterer) FilterJobRejected(opts *bind.FilterOpts, jobId []*big.Int, rejector []common.Address) (*ERC8183JobRejectedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var rejectorRule []interface{}
	for _, rejectorItem := range rejector {
		rejectorRule = append(rejectorRule, rejectorItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "JobRejected", jobIdRule, rejectorRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183JobRejectedIterator{contract: _ERC8183.contract, event: "JobRejected", logs: logs, sub: sub}, nil
}

// WatchJobRejected is a free log subscription operation binding the contract event 0xae7362b1af91f4492868987b9c73990d780060811551b58728fbe96fd1bab275.
//
// Solidity: event JobRejected(uint256 indexed jobId, address indexed rejector, bytes32 reason)
func (_ERC8183 *ERC8183Filterer) WatchJobRejected(opts *bind.WatchOpts, sink chan<- *ERC8183JobRejected, jobId []*big.Int, rejector []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var rejectorRule []interface{}
	for _, rejectorItem := range rejector {
		rejectorRule = append(rejectorRule, rejectorItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "JobRejected", jobIdRule, rejectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183JobRejected)
				if err := _ERC8183.contract.UnpackLog(event, "JobRejected", log); err != nil {
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

// ParseJobRejected is a log parse operation binding the contract event 0xae7362b1af91f4492868987b9c73990d780060811551b58728fbe96fd1bab275.
//
// Solidity: event JobRejected(uint256 indexed jobId, address indexed rejector, bytes32 reason)
func (_ERC8183 *ERC8183Filterer) ParseJobRejected(log types.Log) (*ERC8183JobRejected, error) {
	event := new(ERC8183JobRejected)
	if err := _ERC8183.contract.UnpackLog(event, "JobRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183JobSubmittedIterator is returned from FilterJobSubmitted and is used to iterate over the raw logs and unpacked data for JobSubmitted events raised by the ERC8183 contract.
type ERC8183JobSubmittedIterator struct {
	Event *ERC8183JobSubmitted // Event containing the contract specifics and raw log

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
func (it *ERC8183JobSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183JobSubmitted)
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
		it.Event = new(ERC8183JobSubmitted)
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
func (it *ERC8183JobSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183JobSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183JobSubmitted represents a JobSubmitted event raised by the ERC8183 contract.
type ERC8183JobSubmitted struct {
	JobId       *big.Int
	Provider    common.Address
	Deliverable [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterJobSubmitted is a free log retrieval operation binding the contract event 0x80c17db79857f338a6a6df68a6883ecc0ce78e2202fe61ed979733573f40538e.
//
// Solidity: event JobSubmitted(uint256 indexed jobId, address indexed provider, bytes32 deliverable)
func (_ERC8183 *ERC8183Filterer) FilterJobSubmitted(opts *bind.FilterOpts, jobId []*big.Int, provider []common.Address) (*ERC8183JobSubmittedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "JobSubmitted", jobIdRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183JobSubmittedIterator{contract: _ERC8183.contract, event: "JobSubmitted", logs: logs, sub: sub}, nil
}

// WatchJobSubmitted is a free log subscription operation binding the contract event 0x80c17db79857f338a6a6df68a6883ecc0ce78e2202fe61ed979733573f40538e.
//
// Solidity: event JobSubmitted(uint256 indexed jobId, address indexed provider, bytes32 deliverable)
func (_ERC8183 *ERC8183Filterer) WatchJobSubmitted(opts *bind.WatchOpts, sink chan<- *ERC8183JobSubmitted, jobId []*big.Int, provider []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "JobSubmitted", jobIdRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183JobSubmitted)
				if err := _ERC8183.contract.UnpackLog(event, "JobSubmitted", log); err != nil {
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

// ParseJobSubmitted is a log parse operation binding the contract event 0x80c17db79857f338a6a6df68a6883ecc0ce78e2202fe61ed979733573f40538e.
//
// Solidity: event JobSubmitted(uint256 indexed jobId, address indexed provider, bytes32 deliverable)
func (_ERC8183 *ERC8183Filterer) ParseJobSubmitted(log types.Log) (*ERC8183JobSubmitted, error) {
	event := new(ERC8183JobSubmitted)
	if err := _ERC8183.contract.UnpackLog(event, "JobSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183PaymentReleasedIterator is returned from FilterPaymentReleased and is used to iterate over the raw logs and unpacked data for PaymentReleased events raised by the ERC8183 contract.
type ERC8183PaymentReleasedIterator struct {
	Event *ERC8183PaymentReleased // Event containing the contract specifics and raw log

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
func (it *ERC8183PaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183PaymentReleased)
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
		it.Event = new(ERC8183PaymentReleased)
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
func (it *ERC8183PaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183PaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183PaymentReleased represents a PaymentReleased event raised by the ERC8183 contract.
type ERC8183PaymentReleased struct {
	JobId    *big.Int
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaymentReleased is a free log retrieval operation binding the contract event 0x21d71db5be59bb9fa133895586b7404307dd33fb93b16db09dc6f1d9d7d231b0.
//
// Solidity: event PaymentReleased(uint256 indexed jobId, address indexed provider, uint256 amount)
func (_ERC8183 *ERC8183Filterer) FilterPaymentReleased(opts *bind.FilterOpts, jobId []*big.Int, provider []common.Address) (*ERC8183PaymentReleasedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "PaymentReleased", jobIdRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183PaymentReleasedIterator{contract: _ERC8183.contract, event: "PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchPaymentReleased is a free log subscription operation binding the contract event 0x21d71db5be59bb9fa133895586b7404307dd33fb93b16db09dc6f1d9d7d231b0.
//
// Solidity: event PaymentReleased(uint256 indexed jobId, address indexed provider, uint256 amount)
func (_ERC8183 *ERC8183Filterer) WatchPaymentReleased(opts *bind.WatchOpts, sink chan<- *ERC8183PaymentReleased, jobId []*big.Int, provider []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "PaymentReleased", jobIdRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183PaymentReleased)
				if err := _ERC8183.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
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

// ParsePaymentReleased is a log parse operation binding the contract event 0x21d71db5be59bb9fa133895586b7404307dd33fb93b16db09dc6f1d9d7d231b0.
//
// Solidity: event PaymentReleased(uint256 indexed jobId, address indexed provider, uint256 amount)
func (_ERC8183 *ERC8183Filterer) ParsePaymentReleased(log types.Log) (*ERC8183PaymentReleased, error) {
	event := new(ERC8183PaymentReleased)
	if err := _ERC8183.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183PlatformFeePaidIterator is returned from FilterPlatformFeePaid and is used to iterate over the raw logs and unpacked data for PlatformFeePaid events raised by the ERC8183 contract.
type ERC8183PlatformFeePaidIterator struct {
	Event *ERC8183PlatformFeePaid // Event containing the contract specifics and raw log

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
func (it *ERC8183PlatformFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183PlatformFeePaid)
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
		it.Event = new(ERC8183PlatformFeePaid)
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
func (it *ERC8183PlatformFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183PlatformFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183PlatformFeePaid represents a PlatformFeePaid event raised by the ERC8183 contract.
type ERC8183PlatformFeePaid struct {
	JobId            *big.Int
	PlatformTreasury common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeePaid is a free log retrieval operation binding the contract event 0x0e61cfd0ee655641b3ed0a1dd373ec3cacf4c090de1f33708e3a4091f4947662.
//
// Solidity: event PlatformFeePaid(uint256 indexed jobId, address indexed platformTreasury, uint256 amount)
func (_ERC8183 *ERC8183Filterer) FilterPlatformFeePaid(opts *bind.FilterOpts, jobId []*big.Int, platformTreasury []common.Address) (*ERC8183PlatformFeePaidIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var platformTreasuryRule []interface{}
	for _, platformTreasuryItem := range platformTreasury {
		platformTreasuryRule = append(platformTreasuryRule, platformTreasuryItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "PlatformFeePaid", jobIdRule, platformTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183PlatformFeePaidIterator{contract: _ERC8183.contract, event: "PlatformFeePaid", logs: logs, sub: sub}, nil
}

// WatchPlatformFeePaid is a free log subscription operation binding the contract event 0x0e61cfd0ee655641b3ed0a1dd373ec3cacf4c090de1f33708e3a4091f4947662.
//
// Solidity: event PlatformFeePaid(uint256 indexed jobId, address indexed platformTreasury, uint256 amount)
func (_ERC8183 *ERC8183Filterer) WatchPlatformFeePaid(opts *bind.WatchOpts, sink chan<- *ERC8183PlatformFeePaid, jobId []*big.Int, platformTreasury []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var platformTreasuryRule []interface{}
	for _, platformTreasuryItem := range platformTreasury {
		platformTreasuryRule = append(platformTreasuryRule, platformTreasuryItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "PlatformFeePaid", jobIdRule, platformTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183PlatformFeePaid)
				if err := _ERC8183.contract.UnpackLog(event, "PlatformFeePaid", log); err != nil {
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

// ParsePlatformFeePaid is a log parse operation binding the contract event 0x0e61cfd0ee655641b3ed0a1dd373ec3cacf4c090de1f33708e3a4091f4947662.
//
// Solidity: event PlatformFeePaid(uint256 indexed jobId, address indexed platformTreasury, uint256 amount)
func (_ERC8183 *ERC8183Filterer) ParsePlatformFeePaid(log types.Log) (*ERC8183PlatformFeePaid, error) {
	event := new(ERC8183PlatformFeePaid)
	if err := _ERC8183.contract.UnpackLog(event, "PlatformFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183PlatformFeeSetIterator is returned from FilterPlatformFeeSet and is used to iterate over the raw logs and unpacked data for PlatformFeeSet events raised by the ERC8183 contract.
type ERC8183PlatformFeeSetIterator struct {
	Event *ERC8183PlatformFeeSet // Event containing the contract specifics and raw log

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
func (it *ERC8183PlatformFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183PlatformFeeSet)
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
		it.Event = new(ERC8183PlatformFeeSet)
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
func (it *ERC8183PlatformFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183PlatformFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183PlatformFeeSet represents a PlatformFeeSet event raised by the ERC8183 contract.
type ERC8183PlatformFeeSet struct {
	FeeBP    *big.Int
	Treasury common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeSet is a free log retrieval operation binding the contract event 0x2c913b218a7fcc4905fb9f9dc3a7ecd1b0f9f8b086cdcd514f2fc4c885389930.
//
// Solidity: event PlatformFeeSet(uint256 feeBP, address indexed treasury)
func (_ERC8183 *ERC8183Filterer) FilterPlatformFeeSet(opts *bind.FilterOpts, treasury []common.Address) (*ERC8183PlatformFeeSetIterator, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "PlatformFeeSet", treasuryRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183PlatformFeeSetIterator{contract: _ERC8183.contract, event: "PlatformFeeSet", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeSet is a free log subscription operation binding the contract event 0x2c913b218a7fcc4905fb9f9dc3a7ecd1b0f9f8b086cdcd514f2fc4c885389930.
//
// Solidity: event PlatformFeeSet(uint256 feeBP, address indexed treasury)
func (_ERC8183 *ERC8183Filterer) WatchPlatformFeeSet(opts *bind.WatchOpts, sink chan<- *ERC8183PlatformFeeSet, treasury []common.Address) (event.Subscription, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "PlatformFeeSet", treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183PlatformFeeSet)
				if err := _ERC8183.contract.UnpackLog(event, "PlatformFeeSet", log); err != nil {
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

// ParsePlatformFeeSet is a log parse operation binding the contract event 0x2c913b218a7fcc4905fb9f9dc3a7ecd1b0f9f8b086cdcd514f2fc4c885389930.
//
// Solidity: event PlatformFeeSet(uint256 feeBP, address indexed treasury)
func (_ERC8183 *ERC8183Filterer) ParsePlatformFeeSet(log types.Log) (*ERC8183PlatformFeeSet, error) {
	event := new(ERC8183PlatformFeeSet)
	if err := _ERC8183.contract.UnpackLog(event, "PlatformFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183ProviderSetIterator is returned from FilterProviderSet and is used to iterate over the raw logs and unpacked data for ProviderSet events raised by the ERC8183 contract.
type ERC8183ProviderSetIterator struct {
	Event *ERC8183ProviderSet // Event containing the contract specifics and raw log

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
func (it *ERC8183ProviderSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183ProviderSet)
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
		it.Event = new(ERC8183ProviderSet)
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
func (it *ERC8183ProviderSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183ProviderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183ProviderSet represents a ProviderSet event raised by the ERC8183 contract.
type ERC8183ProviderSet struct {
	JobId    *big.Int
	Provider common.Address
	AgentId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderSet is a free log retrieval operation binding the contract event 0x49d0adc1d0f8a0d589f5eba9d9764b09b597ab6227e8ba63b0884599f1acd0e7.
//
// Solidity: event ProviderSet(uint256 indexed jobId, address indexed provider, uint256 agentId)
func (_ERC8183 *ERC8183Filterer) FilterProviderSet(opts *bind.FilterOpts, jobId []*big.Int, provider []common.Address) (*ERC8183ProviderSetIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "ProviderSet", jobIdRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183ProviderSetIterator{contract: _ERC8183.contract, event: "ProviderSet", logs: logs, sub: sub}, nil
}

// WatchProviderSet is a free log subscription operation binding the contract event 0x49d0adc1d0f8a0d589f5eba9d9764b09b597ab6227e8ba63b0884599f1acd0e7.
//
// Solidity: event ProviderSet(uint256 indexed jobId, address indexed provider, uint256 agentId)
func (_ERC8183 *ERC8183Filterer) WatchProviderSet(opts *bind.WatchOpts, sink chan<- *ERC8183ProviderSet, jobId []*big.Int, provider []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "ProviderSet", jobIdRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183ProviderSet)
				if err := _ERC8183.contract.UnpackLog(event, "ProviderSet", log); err != nil {
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

// ParseProviderSet is a log parse operation binding the contract event 0x49d0adc1d0f8a0d589f5eba9d9764b09b597ab6227e8ba63b0884599f1acd0e7.
//
// Solidity: event ProviderSet(uint256 indexed jobId, address indexed provider, uint256 agentId)
func (_ERC8183 *ERC8183Filterer) ParseProviderSet(log types.Log) (*ERC8183ProviderSet, error) {
	event := new(ERC8183ProviderSet)
	if err := _ERC8183.contract.UnpackLog(event, "ProviderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183RefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the ERC8183 contract.
type ERC8183RefundedIterator struct {
	Event *ERC8183Refunded // Event containing the contract specifics and raw log

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
func (it *ERC8183RefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183Refunded)
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
		it.Event = new(ERC8183Refunded)
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
func (it *ERC8183RefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183RefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183Refunded represents a Refunded event raised by the ERC8183 contract.
type ERC8183Refunded struct {
	JobId  *big.Int
	Client common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x7ca5472b7ea78c2c0141c5a12ee6d170cf4ce8ed06be3d22c8252ddfc7a6a2c4.
//
// Solidity: event Refunded(uint256 indexed jobId, address indexed client, uint256 amount)
func (_ERC8183 *ERC8183Filterer) FilterRefunded(opts *bind.FilterOpts, jobId []*big.Int, client []common.Address) (*ERC8183RefundedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "Refunded", jobIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183RefundedIterator{contract: _ERC8183.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x7ca5472b7ea78c2c0141c5a12ee6d170cf4ce8ed06be3d22c8252ddfc7a6a2c4.
//
// Solidity: event Refunded(uint256 indexed jobId, address indexed client, uint256 amount)
func (_ERC8183 *ERC8183Filterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *ERC8183Refunded, jobId []*big.Int, client []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "Refunded", jobIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183Refunded)
				if err := _ERC8183.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0x7ca5472b7ea78c2c0141c5a12ee6d170cf4ce8ed06be3d22c8252ddfc7a6a2c4.
//
// Solidity: event Refunded(uint256 indexed jobId, address indexed client, uint256 amount)
func (_ERC8183 *ERC8183Filterer) ParseRefunded(log types.Log) (*ERC8183Refunded, error) {
	event := new(ERC8183Refunded)
	if err := _ERC8183.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ERC8183 contract.
type ERC8183RoleAdminChangedIterator struct {
	Event *ERC8183RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC8183RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183RoleAdminChanged)
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
		it.Event = new(ERC8183RoleAdminChanged)
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
func (it *ERC8183RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183RoleAdminChanged represents a RoleAdminChanged event raised by the ERC8183 contract.
type ERC8183RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC8183 *ERC8183Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ERC8183RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183RoleAdminChangedIterator{contract: _ERC8183.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC8183 *ERC8183Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC8183RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183RoleAdminChanged)
				if err := _ERC8183.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC8183 *ERC8183Filterer) ParseRoleAdminChanged(log types.Log) (*ERC8183RoleAdminChanged, error) {
	event := new(ERC8183RoleAdminChanged)
	if err := _ERC8183.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC8183 contract.
type ERC8183RoleGrantedIterator struct {
	Event *ERC8183RoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC8183RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183RoleGranted)
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
		it.Event = new(ERC8183RoleGranted)
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
func (it *ERC8183RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183RoleGranted represents a RoleGranted event raised by the ERC8183 contract.
type ERC8183RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC8183 *ERC8183Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC8183RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183RoleGrantedIterator{contract: _ERC8183.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC8183 *ERC8183Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC8183RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183RoleGranted)
				if err := _ERC8183.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC8183 *ERC8183Filterer) ParseRoleGranted(log types.Log) (*ERC8183RoleGranted, error) {
	event := new(ERC8183RoleGranted)
	if err := _ERC8183.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC8183 contract.
type ERC8183RoleRevokedIterator struct {
	Event *ERC8183RoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC8183RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183RoleRevoked)
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
		it.Event = new(ERC8183RoleRevoked)
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
func (it *ERC8183RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183RoleRevoked represents a RoleRevoked event raised by the ERC8183 contract.
type ERC8183RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC8183 *ERC8183Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC8183RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183RoleRevokedIterator{contract: _ERC8183.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC8183 *ERC8183Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC8183RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183RoleRevoked)
				if err := _ERC8183.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC8183 *ERC8183Filterer) ParseRoleRevoked(log types.Log) (*ERC8183RoleRevoked, error) {
	event := new(ERC8183RoleRevoked)
	if err := _ERC8183.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC8183UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC8183 contract.
type ERC8183UpgradedIterator struct {
	Event *ERC8183Upgraded // Event containing the contract specifics and raw log

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
func (it *ERC8183UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC8183Upgraded)
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
		it.Event = new(ERC8183Upgraded)
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
func (it *ERC8183UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC8183UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC8183Upgraded represents a Upgraded event raised by the ERC8183 contract.
type ERC8183Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC8183 *ERC8183Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC8183UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC8183.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC8183UpgradedIterator{contract: _ERC8183.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC8183 *ERC8183Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC8183Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC8183.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC8183Upgraded)
				if err := _ERC8183.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ERC8183 *ERC8183Filterer) ParseUpgraded(log types.Log) (*ERC8183Upgraded, error) {
	event := new(ERC8183Upgraded)
	if err := _ERC8183.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
