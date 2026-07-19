// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package epoch

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

// EpochMetaData contains all meta data concerning the Epoch contract.
var EpochMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"check\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"current\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_slots\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlots\",\"inputs\":[{\"name\":\"_slots\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slots\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetEpoch\",\"inputs\":[{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateSlots\",\"inputs\":[{\"name\":\"_s\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a080604052346100cc57306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100bd57506001600160401b036002600160401b031982821601610078575b60405161107f90816100d182396080518181816105b601526106a30152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f8080610059565b63f92ee8a960e01b8152600490fd5b5f80fdfe6080604090808252600480361015610015575f80fd5b5f3560e01c91826301ffc9a71461095a5750816312a02c821461091b578163248a9ca3146108e55781632f2ff15d146108bd57816336568abe1461087957816348547d69146108525781634f1ef2861461061a57816352d1902d146105a3578163919840ad1461058b57816391d148541461053b5781639fa6a6e314610515578163a217fddf146104fb578163ad3cb1cc146104bf578163cbccf09114610437578163ccc57490146103fd578163d547741f146103b5578163d7eecc7e14610129575063ffa1ad74146100e6575f80fd5b34610125575f36600319011261012557805161012191610105826109d8565b60058252640322e302e360dc1b60208301525191829182610a28565b0390f35b5f80fd5b9050346101255781600319360112610125576101436109ac565b9061014c6109c2565b917ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0092835460ff81871c1615936001600160401b0393848316801590816103ad575b60011490816103a3575b15908161039a575b5061038b5767ffffffffffffffff198381166001178855928661036c575b5060ff8754891c161561035d576101d88582161515610c59565b67ffffffffffffffff60401b5f5491891b169067ffffffffffffffff60401b1916175f55610204610abb565b905f8252875160208101914483525f8a83015289825260608201908282108883111761034a57506102fd93602096937f9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd759896936102d9938d5251902087820190815261026f82610b0f565b4382525f5487166102c16102cf8e610294600161028b86610a6f565b50015494610ae3565b90518c810194855244602086015260c09190911b6001600160c01b03191660408501529182906048850190565b03601f198101835282610a07565b5190209052610b0f565b5f5490846102e8818416610b5d565b169116175f556102f781610d39565b50610dd6565b505f54168551908152a161030d57005b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b604190634e487b7160e01b5f525260245ffd5b508651631afcd79f60e31b8152fd5b68ffffffffffffffffff1916680100000000000000011787555f6101be565b50865163f92ee8a960e01b8152fd5b9050155f6101a0565b303b159150610198565b87915061018e565b82346101255780600319360112610125576103fb91356103f660016103d86109c2565b93835f525f8051602061102a8339815191526020525f200154610d0b565b610f19565b005b8234610125575f36600319011261012557602090517f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f558152f35b82346101255760203660031901126101255760207fba1da09884224d94b905b0a3a828942e734c70ad3abf9cb4a66bda28fa7e6486916104756109ac565b9061047e610c92565b6001600160401b03821691610494831515610c59565b67ffffffffffffffff60401b5f5491831b169067ffffffffffffffff60401b1916175f5551908152a1005b8234610125575f366003190112610125578051610121916104df826109d8565b60058252640352e302e360dc1b60208301525191829182610a28565b8234610125575f36600319011261012557602090515f8152f35b8234610125575f366003190112610125576020906001600160401b035f54169051908152f35b82346101255780600319360112610125576020916105576109c2565b90355f525f8051602061102a8339815191528352815f209060018060a01b03165f52825260ff815f20541690519015158152f35b34610125575f366003190112610125576103fb610b75565b8234610125575f366003190112610125577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361060d57602090517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b5163703e46dd60e11b8152fd5b905081600319360112610125578035916001600160a01b038084169190828503610125576024948535916001600160401b0390818411610125573660238501121561012557838701359182116108405760209185519461068384601f19601f8501160187610a07565b818652368a838301011161012557815f928b8693018389013786010152807f000000000000000000000000000000000000000000000000000000000000000016803014918215610812575b5050610802576106dc610c92565b83516352d1902d60e01b81529080828881895afa9182915f936107d2575b5050610714575050505191634c9c8ce360e01b8352820152fd5b85938591887f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc918281036107bd5750843b156107a9575080546001600160a01b0319168317905551907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a282511561079357506103fb9250610f98565b9150503461079d57005b63b398979f60e01b8152fd5b8251634c9c8ce360e01b8152808801859052fd5b8351632a87526960e21b815280890191909152fd5b9080929350813d83116107fb575b6107ea8183610a07565b810103126101255751905f806106fa565b503d6107e0565b835163703e46dd60e11b81528690fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141590505f806106ce565b87604188634e487b7160e01b5f52525ffd5b8234610125575f366003190112610125575f548151911c6001600160401b03168152602090f35b82346101255780600319360112610125576108926109c2565b90336001600160a01b038316036108ae57506103fb9135610f19565b5163334bd91960e11b81529050fd5b82346101255780600319360112610125576103fb91356108e060016103d86109c2565b610e96565b823461012557602036600319011261012557602091355f525f8051602061102a83398151915282526001815f2001549051908152f35b8234610125576020366003190112610125576109356109ac565b600161094b61094383610a6f565b505492610a6f565b50015482519182526020820152f35b903461012557602036600319011261012557359063ffffffff60e01b821680920361012557602091637965db0b60e01b811490811561099b575b5015158152f35b6301ffc9a760e01b14905083610994565b600435906001600160401b038216820361012557565b602435906001600160a01b038216820361012557565b604081019081106001600160401b038211176109f357604052565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b038211176109f357604052565b602080825282518183018190529093925f5b828110610a5b57505060409293505f838284010152601f8019910116010190565b818101860151848201604001528501610a3a565b600254811015610aa75760025f5260011b7f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01905f90565b634e487b7160e01b5f52603260045260245ffd5b60405190604082018281106001600160401b038211176109f3576040525f6020838281520152565b9060016001600160401b0380931601918211610afb57565b634e487b7160e01b5f52601160045260245ffd5b600254680100000000000000008110156109f357806001610b339201600255610a6f565b919091610b4a576020816001925184550151910155565b634e487b7160e01b5f525f60045260245ffd5b6001600160401b03809116908114610afb5760010190565b5f546001600160401b039081811690610b8d82610a6f565b5054430390438211610afb5760401c831611610c55577f9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd7591610c2e602092610bd3610abb565b904382526102c1610c22610bf46001610beb85610a6f565b50015493610ae3565b6040805189810195865244602087015260c09290921b6001600160c01b031916908501529182906048850190565b51902084820152610b0f565b5f5490610c3c818316610b5d565b1680916001600160401b031916175f55604051908152a1565b5050565b15610c6057565b60405162461bcd60e51b815260206004820152600a6024820152697a65726f20736c6f747360b01b6044820152606490fd5b335f9081527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f559060ff1615610ced5750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b805f525f8051602061102a83398151915260205260405f20335f5260205260ff60405f20541615610ced5750565b6001600160a01b03165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f8051602061102a8339815191529060ff16610dd0575f805260205260405f20815f5260205260405f20600160ff1982541617905533905f7f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50505f90565b6001600160a01b03165f8181527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5591905f8051602061102a8339815191529060ff16610e8f57825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b5050505f90565b90815f525f8051602061102a8339815191528060205260405f209160018060a01b031691825f5260205260ff60405f205416155f14610e8f57825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b90815f525f8051602061102a8339815191528060205260405f209160018060a01b031691825f5260205260ff60405f2054165f14610e8f57825f5260205260405f20815f5260205260405f2060ff19815416905533917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b905f8091602081519101845af48080611016575b15610fcc5750506040513d81523d5f602083013e60203d82010160405290565b15610ff357604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b3d15611004576040513d5f823e3d90fd5b60405163d6bda27560e01b8152600490fd5b503d151580610fac5750813b1515610fac56fe02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a2646970667358221220ebfceeee0ed50fa07a9d24c04aa8b10a8a1743059765ee5b8df88ba0cdf40c5d64736f6c63430008180033",
}

// EpochABI is the input ABI used to generate the binding from.
// Deprecated: Use EpochMetaData.ABI instead.
var EpochABI = EpochMetaData.ABI

// EpochBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EpochMetaData.Bin instead.
var EpochBin = EpochMetaData.Bin

// DeployEpoch deploys a new Ethereum contract, binding an instance of Epoch to it.
func DeployEpoch(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Epoch, error) {
	parsed, err := EpochMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EpochBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Epoch{EpochCaller: EpochCaller{contract: contract}, EpochTransactor: EpochTransactor{contract: contract}, EpochFilterer: EpochFilterer{contract: contract}}, nil
}

// Epoch is an auto generated Go binding around an Ethereum contract.
type Epoch struct {
	EpochCaller     // Read-only binding to the contract
	EpochTransactor // Write-only binding to the contract
	EpochFilterer   // Log filterer for contract events
}

// EpochCaller is an auto generated read-only Go binding around an Ethereum contract.
type EpochCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EpochTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EpochFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EpochSession struct {
	Contract     *Epoch            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EpochCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EpochCallerSession struct {
	Contract *EpochCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EpochTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EpochTransactorSession struct {
	Contract     *EpochTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EpochRaw is an auto generated low-level Go binding around an Ethereum contract.
type EpochRaw struct {
	Contract *Epoch // Generic contract binding to access the raw methods on
}

// EpochCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EpochCallerRaw struct {
	Contract *EpochCaller // Generic read-only contract binding to access the raw methods on
}

// EpochTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EpochTransactorRaw struct {
	Contract *EpochTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEpoch creates a new instance of Epoch, bound to a specific deployed contract.
func NewEpoch(address common.Address, backend bind.ContractBackend) (*Epoch, error) {
	contract, err := bindEpoch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Epoch{EpochCaller: EpochCaller{contract: contract}, EpochTransactor: EpochTransactor{contract: contract}, EpochFilterer: EpochFilterer{contract: contract}}, nil
}

// NewEpochCaller creates a new read-only instance of Epoch, bound to a specific deployed contract.
func NewEpochCaller(address common.Address, caller bind.ContractCaller) (*EpochCaller, error) {
	contract, err := bindEpoch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EpochCaller{contract: contract}, nil
}

// NewEpochTransactor creates a new write-only instance of Epoch, bound to a specific deployed contract.
func NewEpochTransactor(address common.Address, transactor bind.ContractTransactor) (*EpochTransactor, error) {
	contract, err := bindEpoch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EpochTransactor{contract: contract}, nil
}

// NewEpochFilterer creates a new log filterer instance of Epoch, bound to a specific deployed contract.
func NewEpochFilterer(address common.Address, filterer bind.ContractFilterer) (*EpochFilterer, error) {
	contract, err := bindEpoch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EpochFilterer{contract: contract}, nil
}

// bindEpoch binds a generic wrapper to an already deployed contract.
func bindEpoch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EpochMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Epoch *EpochRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Epoch.Contract.EpochCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Epoch *EpochRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Epoch.Contract.EpochTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Epoch *EpochRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Epoch.Contract.EpochTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Epoch *EpochCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Epoch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Epoch *EpochTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Epoch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Epoch *EpochTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Epoch.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Epoch *EpochCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Epoch *EpochSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Epoch.Contract.DEFAULTADMINROLE(&_Epoch.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Epoch *EpochCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Epoch.Contract.DEFAULTADMINROLE(&_Epoch.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Epoch *EpochCaller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Epoch *EpochSession) GOVERNORROLE() ([32]byte, error) {
	return _Epoch.Contract.GOVERNORROLE(&_Epoch.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Epoch *EpochCallerSession) GOVERNORROLE() ([32]byte, error) {
	return _Epoch.Contract.GOVERNORROLE(&_Epoch.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Epoch *EpochCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Epoch *EpochSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Epoch.Contract.UPGRADEINTERFACEVERSION(&_Epoch.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Epoch *EpochCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Epoch.Contract.UPGRADEINTERFACEVERSION(&_Epoch.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Epoch *EpochCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Epoch *EpochSession) VERSION() (string, error) {
	return _Epoch.Contract.VERSION(&_Epoch.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Epoch *EpochCallerSession) VERSION() (string, error) {
	return _Epoch.Contract.VERSION(&_Epoch.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Epoch *EpochCaller) Current(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "current")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Epoch *EpochSession) Current() (uint64, error) {
	return _Epoch.Contract.Current(&_Epoch.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Epoch *EpochCallerSession) Current() (uint64, error) {
	return _Epoch.Contract.Current(&_Epoch.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns(uint256, bytes32)
func (_Epoch *EpochCaller) GetEpoch(opts *bind.CallOpts, _epoch uint64) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "getEpoch", _epoch)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns(uint256, bytes32)
func (_Epoch *EpochSession) GetEpoch(_epoch uint64) (*big.Int, [32]byte, error) {
	return _Epoch.Contract.GetEpoch(&_Epoch.CallOpts, _epoch)
}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns(uint256, bytes32)
func (_Epoch *EpochCallerSession) GetEpoch(_epoch uint64) (*big.Int, [32]byte, error) {
	return _Epoch.Contract.GetEpoch(&_Epoch.CallOpts, _epoch)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Epoch *EpochCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Epoch *EpochSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Epoch.Contract.GetRoleAdmin(&_Epoch.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Epoch *EpochCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Epoch.Contract.GetRoleAdmin(&_Epoch.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Epoch *EpochCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Epoch *EpochSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Epoch.Contract.HasRole(&_Epoch.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Epoch *EpochCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Epoch.Contract.HasRole(&_Epoch.CallOpts, role, account)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Epoch *EpochCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Epoch *EpochSession) ProxiableUUID() ([32]byte, error) {
	return _Epoch.Contract.ProxiableUUID(&_Epoch.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Epoch *EpochCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Epoch.Contract.ProxiableUUID(&_Epoch.CallOpts)
}

// Slots is a free data retrieval call binding the contract method 0x48547d69.
//
// Solidity: function slots() view returns(uint64)
func (_Epoch *EpochCaller) Slots(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "slots")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Slots is a free data retrieval call binding the contract method 0x48547d69.
//
// Solidity: function slots() view returns(uint64)
func (_Epoch *EpochSession) Slots() (uint64, error) {
	return _Epoch.Contract.Slots(&_Epoch.CallOpts)
}

// Slots is a free data retrieval call binding the contract method 0x48547d69.
//
// Solidity: function slots() view returns(uint64)
func (_Epoch *EpochCallerSession) Slots() (uint64, error) {
	return _Epoch.Contract.Slots(&_Epoch.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Epoch *EpochCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Epoch *EpochSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Epoch.Contract.SupportsInterface(&_Epoch.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Epoch *EpochCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Epoch.Contract.SupportsInterface(&_Epoch.CallOpts, interfaceId)
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_Epoch *EpochTransactor) Check(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "check")
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_Epoch *EpochSession) Check() (*types.Transaction, error) {
	return _Epoch.Contract.Check(&_Epoch.TransactOpts)
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_Epoch *EpochTransactorSession) Check() (*types.Transaction, error) {
	return _Epoch.Contract.Check(&_Epoch.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Epoch *EpochTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Epoch *EpochSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.GrantRole(&_Epoch.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Epoch *EpochTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.GrantRole(&_Epoch.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 _slots, address initialOwner) returns()
func (_Epoch *EpochTransactor) Initialize(opts *bind.TransactOpts, _slots uint64, initialOwner common.Address) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "initialize", _slots, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 _slots, address initialOwner) returns()
func (_Epoch *EpochSession) Initialize(_slots uint64, initialOwner common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.Initialize(&_Epoch.TransactOpts, _slots, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7eecc7e.
//
// Solidity: function initialize(uint64 _slots, address initialOwner) returns()
func (_Epoch *EpochTransactorSession) Initialize(_slots uint64, initialOwner common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.Initialize(&_Epoch.TransactOpts, _slots, initialOwner)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Epoch *EpochTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Epoch *EpochSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.RenounceRole(&_Epoch.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Epoch *EpochTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.RenounceRole(&_Epoch.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Epoch *EpochTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Epoch *EpochSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.RevokeRole(&_Epoch.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Epoch *EpochTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Epoch.Contract.RevokeRole(&_Epoch.TransactOpts, role, account)
}

// SetSlots is a paid mutator transaction binding the contract method 0xcbccf091.
//
// Solidity: function setSlots(uint64 _slots) returns()
func (_Epoch *EpochTransactor) SetSlots(opts *bind.TransactOpts, _slots uint64) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "setSlots", _slots)
}

// SetSlots is a paid mutator transaction binding the contract method 0xcbccf091.
//
// Solidity: function setSlots(uint64 _slots) returns()
func (_Epoch *EpochSession) SetSlots(_slots uint64) (*types.Transaction, error) {
	return _Epoch.Contract.SetSlots(&_Epoch.TransactOpts, _slots)
}

// SetSlots is a paid mutator transaction binding the contract method 0xcbccf091.
//
// Solidity: function setSlots(uint64 _slots) returns()
func (_Epoch *EpochTransactorSession) SetSlots(_slots uint64) (*types.Transaction, error) {
	return _Epoch.Contract.SetSlots(&_Epoch.TransactOpts, _slots)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Epoch *EpochTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Epoch *EpochSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Epoch.Contract.UpgradeToAndCall(&_Epoch.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Epoch *EpochTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Epoch.Contract.UpgradeToAndCall(&_Epoch.TransactOpts, newImplementation, data)
}

// EpochInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Epoch contract.
type EpochInitializedIterator struct {
	Event *EpochInitialized // Event containing the contract specifics and raw log

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
func (it *EpochInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochInitialized)
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
		it.Event = new(EpochInitialized)
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
func (it *EpochInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochInitialized represents a Initialized event raised by the Epoch contract.
type EpochInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Epoch *EpochFilterer) FilterInitialized(opts *bind.FilterOpts) (*EpochInitializedIterator, error) {

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EpochInitializedIterator{contract: _Epoch.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Epoch *EpochFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EpochInitialized) (event.Subscription, error) {

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochInitialized)
				if err := _Epoch.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Epoch *EpochFilterer) ParseInitialized(log types.Log) (*EpochInitialized, error) {
	event := new(EpochInitialized)
	if err := _Epoch.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Epoch contract.
type EpochRoleAdminChangedIterator struct {
	Event *EpochRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EpochRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRoleAdminChanged)
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
		it.Event = new(EpochRoleAdminChanged)
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
func (it *EpochRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRoleAdminChanged represents a RoleAdminChanged event raised by the Epoch contract.
type EpochRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Epoch *EpochFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EpochRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EpochRoleAdminChangedIterator{contract: _Epoch.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Epoch *EpochFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EpochRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRoleAdminChanged)
				if err := _Epoch.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Epoch *EpochFilterer) ParseRoleAdminChanged(log types.Log) (*EpochRoleAdminChanged, error) {
	event := new(EpochRoleAdminChanged)
	if err := _Epoch.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Epoch contract.
type EpochRoleGrantedIterator struct {
	Event *EpochRoleGranted // Event containing the contract specifics and raw log

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
func (it *EpochRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRoleGranted)
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
		it.Event = new(EpochRoleGranted)
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
func (it *EpochRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRoleGranted represents a RoleGranted event raised by the Epoch contract.
type EpochRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Epoch *EpochFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EpochRoleGrantedIterator, error) {

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

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EpochRoleGrantedIterator{contract: _Epoch.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Epoch *EpochFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EpochRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRoleGranted)
				if err := _Epoch.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Epoch *EpochFilterer) ParseRoleGranted(log types.Log) (*EpochRoleGranted, error) {
	event := new(EpochRoleGranted)
	if err := _Epoch.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Epoch contract.
type EpochRoleRevokedIterator struct {
	Event *EpochRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EpochRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRoleRevoked)
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
		it.Event = new(EpochRoleRevoked)
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
func (it *EpochRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRoleRevoked represents a RoleRevoked event raised by the Epoch contract.
type EpochRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Epoch *EpochFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EpochRoleRevokedIterator, error) {

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

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EpochRoleRevokedIterator{contract: _Epoch.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Epoch *EpochFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EpochRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRoleRevoked)
				if err := _Epoch.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Epoch *EpochFilterer) ParseRoleRevoked(log types.Log) (*EpochRoleRevoked, error) {
	event := new(EpochRoleRevoked)
	if err := _Epoch.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochSetEpochIterator is returned from FilterSetEpoch and is used to iterate over the raw logs and unpacked data for SetEpoch events raised by the Epoch contract.
type EpochSetEpochIterator struct {
	Event *EpochSetEpoch // Event containing the contract specifics and raw log

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
func (it *EpochSetEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochSetEpoch)
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
		it.Event = new(EpochSetEpoch)
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
func (it *EpochSetEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochSetEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochSetEpoch represents a SetEpoch event raised by the Epoch contract.
type EpochSetEpoch struct {
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetEpoch is a free log retrieval operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_Epoch *EpochFilterer) FilterSetEpoch(opts *bind.FilterOpts) (*EpochSetEpochIterator, error) {

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "SetEpoch")
	if err != nil {
		return nil, err
	}
	return &EpochSetEpochIterator{contract: _Epoch.contract, event: "SetEpoch", logs: logs, sub: sub}, nil
}

// WatchSetEpoch is a free log subscription operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_Epoch *EpochFilterer) WatchSetEpoch(opts *bind.WatchOpts, sink chan<- *EpochSetEpoch) (event.Subscription, error) {

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "SetEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochSetEpoch)
				if err := _Epoch.contract.UnpackLog(event, "SetEpoch", log); err != nil {
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

// ParseSetEpoch is a log parse operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_Epoch *EpochFilterer) ParseSetEpoch(log types.Log) (*EpochSetEpoch, error) {
	event := new(EpochSetEpoch)
	if err := _Epoch.contract.UnpackLog(event, "SetEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochUpdateSlotsIterator is returned from FilterUpdateSlots and is used to iterate over the raw logs and unpacked data for UpdateSlots events raised by the Epoch contract.
type EpochUpdateSlotsIterator struct {
	Event *EpochUpdateSlots // Event containing the contract specifics and raw log

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
func (it *EpochUpdateSlotsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochUpdateSlots)
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
		it.Event = new(EpochUpdateSlots)
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
func (it *EpochUpdateSlotsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochUpdateSlotsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochUpdateSlots represents a UpdateSlots event raised by the Epoch contract.
type EpochUpdateSlots struct {
	S   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateSlots is a free log retrieval operation binding the contract event 0xba1da09884224d94b905b0a3a828942e734c70ad3abf9cb4a66bda28fa7e6486.
//
// Solidity: event UpdateSlots(uint64 _s)
func (_Epoch *EpochFilterer) FilterUpdateSlots(opts *bind.FilterOpts) (*EpochUpdateSlotsIterator, error) {

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "UpdateSlots")
	if err != nil {
		return nil, err
	}
	return &EpochUpdateSlotsIterator{contract: _Epoch.contract, event: "UpdateSlots", logs: logs, sub: sub}, nil
}

// WatchUpdateSlots is a free log subscription operation binding the contract event 0xba1da09884224d94b905b0a3a828942e734c70ad3abf9cb4a66bda28fa7e6486.
//
// Solidity: event UpdateSlots(uint64 _s)
func (_Epoch *EpochFilterer) WatchUpdateSlots(opts *bind.WatchOpts, sink chan<- *EpochUpdateSlots) (event.Subscription, error) {

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "UpdateSlots")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochUpdateSlots)
				if err := _Epoch.contract.UnpackLog(event, "UpdateSlots", log); err != nil {
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

// ParseUpdateSlots is a log parse operation binding the contract event 0xba1da09884224d94b905b0a3a828942e734c70ad3abf9cb4a66bda28fa7e6486.
//
// Solidity: event UpdateSlots(uint64 _s)
func (_Epoch *EpochFilterer) ParseUpdateSlots(log types.Log) (*EpochUpdateSlots, error) {
	event := new(EpochUpdateSlots)
	if err := _Epoch.contract.UnpackLog(event, "UpdateSlots", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Epoch contract.
type EpochUpgradedIterator struct {
	Event *EpochUpgraded // Event containing the contract specifics and raw log

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
func (it *EpochUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochUpgraded)
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
		it.Event = new(EpochUpgraded)
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
func (it *EpochUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochUpgraded represents a Upgraded event raised by the Epoch contract.
type EpochUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Epoch *EpochFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EpochUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EpochUpgradedIterator{contract: _Epoch.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Epoch *EpochFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EpochUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochUpgraded)
				if err := _Epoch.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Epoch *EpochFilterer) ParseUpgraded(log types.Log) (*EpochUpgraded, error) {
	event := new(EpochUpgraded)
	if err := _Epoch.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
