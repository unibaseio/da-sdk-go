// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stat

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

// StatMetaData contains all meta data concerning the Stat contract.
var StatMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkPiece\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEpoch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eproof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEProof\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_piece\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rsproof\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eproof\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"piece\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPiece\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rsproof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRSProof\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a080604052346100cc57306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100bd57506001600160401b036002600160401b031982821601610078575b6040516111c890816100d1823960805181818161035601526103f90152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f8080610059565b63f92ee8a960e01b8152600490fd5b5f80fdfe6080604090808252600480361015610015575f80fd5b5f3560e01c91826301ffc9a71461085d575081631459457a146106a3578163248a9ca31461066d5781632f2ff15d1461064557816336568abe146106015781634f1ef286146103ba57816352d1902d1461034357816379ca7e0f1461031b57816381cc0c7a146102f3578163900cf0cf146102cc57816391d148541461027c57816395ad21dc14610254578163a217fddf1461023a578163ad3cb1cc146101f6578163c2e8e717146101b3578163ccc5749014610179578163d547741f14610131575063ffa1ad74146100e6575f80fd5b3461012d575f36600319011261012d57805161012991610105826108db565b60058252640322e302e360dc1b602083015251918291602083526020830190610980565b0390f35b5f80fd5b823461012d578060031936011261012d57610177913561017260016101546108c5565b93835f525f805160206111738339815191526020525f200154611016565b611062565b005b823461012d575f36600319011261012d57602090517f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f558152f35b823461012d57602036600319011261012d578135916001600160401b03831161012d576101e86101ed9160209436910161092b565b6109f8565b90519015158152f35b823461012d575f36600319011261012d57805161012991610216826108db565b60058252640352e302e360dc1b602083015251918291602083526020830190610980565b823461012d575f36600319011261012d57602090515f8152f35b823461012d575f36600319011261012d5760015490516001600160a01b039091168152602090f35b823461012d578060031936011261012d576020916102986108c5565b90355f525f805160206111738339815191528352815f209060018060a01b03165f52825260ff815f20541690519015158152f35b823461012d575f36600319011261012d575f5490516001600160a01b039091168152602090f35b823461012d575f36600319011261012d5760035490516001600160a01b039091168152602090f35b823461012d575f36600319011261012d5760025490516001600160a01b039091168152602090f35b823461012d575f36600319011261012d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036103ad57602090517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b5163703e46dd60e11b8152fd5b90508160031936011261012d576103cf6108af565b6024356001600160401b03811161012d576103ed903690840161092b565b6001600160a01b0393907f000000000000000000000000000000000000000000000000000000000000000085163081149081156105d3575b506105c3577f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5594855f526020955f805160206111738339815191528752835f20335f52875260ff845f205416156105a557508316928251956352d1902d60e01b875280878781885afa9687915f98610575575b50506104b557505051634c9c8ce360e01b81529182015260249150fd5b838593877f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9081810361055f5750833b156105485780546001600160a01b0319168317905551907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2825115610532575061017792506110e1565b9150503461053c57005b63b398979f60e01b8152fd5b8151634c9c8ce360e01b8152808701849052602490fd5b86602491845191632a87526960e21b8352820152fd5b9080929850813d831161059e575b61058d818361090a565b8101031261012d5751955f80610498565b503d610583565b835163e2517d3f60e01b815233818801526024810191909152604490fd5b815163703e46dd60e11b81528490fd5b9050857f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc541614155f610425565b823461012d578060031936011261012d5761061a6108c5565b90336001600160a01b0383160361063657506101779135611062565b5163334bd91960e11b81529050fd5b823461012d578060031936011261012d57610177913561066860016101546108c5565b610f93565b823461012d57602036600319011261012d57602091355f525f8051602061117383398151915282526001815f2001549051908152f35b823461012d5760a036600319011261012d576106bd6108af565b6106c56108c5565b6001600160a01b0391604435838116919082900361012d576064359084821680920361012d5760843592858416840361012d577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009586549560ff878a1c1615966001600160401b03811680159081610855575b600114908161084b575b159081610842575b506108325767ffffffffffffffff198116600117895587610813575b5060ff88548a1c1615610803576107bf9697989950816bffffffffffffffffffffffff60a01b9416845f5416175f551682600154161760015581600254161760025560035416176003556107b981610e36565b50610ed3565b506107c657005b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b8851631afcd79f60e31b81528a90fd5b68ffffffffffffffffff1916680100000000000000011788558a610766565b895163f92ee8a960e01b81528b90fd5b9050158c61074a565b303b159150610742565b899150610738565b903461012d57602036600319011261012d57359063ffffffff60e01b821680920361012d57602091637965db0b60e01b811490811561089e575b5015158152f35b6301ffc9a760e01b14905083610897565b600435906001600160a01b038216820361012d57565b602435906001600160a01b038216820361012d57565b604081019081106001600160401b038211176108f657604052565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b038211176108f657604052565b81601f8201121561012d578035906001600160401b0382116108f6576040519261095f601f8401601f19166020018561090a565b8284526020838301011161012d57815f926020809301838601378301015290565b91908251928382525f5b8481106109aa575050825f602080949584010152601f8019910116010190565b60208183018101518483018201520161098a565b51906001600160401b038216820361012d57565b519060ff8216820361012d57565b9081602091031261012d5751801515810361012d5790565b5f54604051639fa6a6e360e01b815290602090829060049082906001600160a01b03165afa8015610c7e575f90610df6575b6001600160401b0391501660028110610dc757610a72906001602060018060a01b0360015416946040518095819263073a179f60e21b83528460048401526024830190610980565b0381875afa928315610c7e575f93610d8b575b506001600160401b038316928315610d825760405193636aadfac760e01b85526004850152606084602481885afa928315610c7e575f935f955f91610d1f575b506001600160401b038291161115610d1557600119016001600160401b038111610c2e576002546003545f978896959490926001600160a01b03928316921690855b610b29575b505050505050505060ff809116911610610b2557600190565b5f90565b60ff999596979995868916878c161015610d0e576040805163ee91365b60e01b81526001600160401b038416600482015260ff8d1660248201529081604481895afa908115610c7e575f905f92610cb6575b506001600160401b031615610c89576040516383f0362760e01b81526001600160401b038416600482015260ff8d166024820152602081604481885afa908115610c7e575f91610c97575b50610c8957604051636fb1c3fd60e11b81526001600160a01b0390911660048201526001600160401b0387166024820152602081604481885afa908115610c7e575f91610c4f575b50610c425786168701958611610c2e5760ff8780979b5b01169796610b07565b634e487b7160e01b5f52601160045260245ffd5b998780975060ff91610c25565b610c71915060203d602011610c77575b610c69818361090a565b8101906109e0565b5f610c0e565b503d610c5f565b6040513d5f823e3d90fd5b50998780975060ff91610c25565b610cb0915060203d602011610c7757610c69818361090a565b5f610bc6565b9150506040813d604011610d06575b81610cd26040938361090a565b8101031261012d576020610ce5826109be565b910151906001600160a01b038216820361012d576001600160401b03610b7b565b3d9150610cc5565b9950610b0c565b5050505050505f90565b94505093506060833d606011610d7a575b81610d3d6060938361090a565b8101031261012d57610d4e836109d2565b93806001600160401b03610d706040610d69602089016109d2565b97016109be565b9695969150610ac5565b3d9150610d30565b50505050505f90565b9092506020813d602011610dbf575b81610da76020938361090a565b8101031261012d57610db8906109be565b915f610a85565b3d9150610d9a565b60405162461bcd60e51b815260206004820152600760248201526634b73b1031bab960c91b6044820152606490fd5b506020813d602011610e2e575b81610e106020938361090a565b8101031261012d57610e296001600160401b03916109be565b610a2a565b3d9150610e03565b6001600160a01b03165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f805160206111738339815191529060ff16610ecd575f805260205260405f20815f5260205260405f20600160ff1982541617905533905f7f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50505f90565b6001600160a01b03165f8181527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5591905f805160206111738339815191529060ff16610f8c57825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b5050505f90565b90815f525f805160206111738339815191528060205260405f209160018060a01b031691825f5260205260ff60405f205416155f14610f8c57825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b805f525f8051602061117383398151915260205260405f20335f5260205260ff60405f205416156110445750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b90815f525f805160206111738339815191528060205260405f209160018060a01b031691825f5260205260ff60405f2054165f14610f8c57825f5260205260405f20815f5260205260405f2060ff19815416905533917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b905f8091602081519101845af4808061115f575b156111155750506040513d81523d5f602083013e60203d82010160405290565b1561113c57604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b3d1561114d576040513d5f823e3d90fd5b60405163d6bda27560e01b8152600490fd5b503d1515806110f55750813b15156110f556fe02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a2646970667358221220ef5f5666ef1c1e04206a749433d7e44f21fa3744c2c589f9fc355f2a208440a464736f6c63430008180033",
}

// StatABI is the input ABI used to generate the binding from.
// Deprecated: Use StatMetaData.ABI instead.
var StatABI = StatMetaData.ABI

// StatBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StatMetaData.Bin instead.
var StatBin = StatMetaData.Bin

// DeployStat deploys a new Ethereum contract, binding an instance of Stat to it.
func DeployStat(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stat, error) {
	parsed, err := StatMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StatBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stat{StatCaller: StatCaller{contract: contract}, StatTransactor: StatTransactor{contract: contract}, StatFilterer: StatFilterer{contract: contract}}, nil
}

// Stat is an auto generated Go binding around an Ethereum contract.
type Stat struct {
	StatCaller     // Read-only binding to the contract
	StatTransactor // Write-only binding to the contract
	StatFilterer   // Log filterer for contract events
}

// StatCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatSession struct {
	Contract     *Stat             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatCallerSession struct {
	Contract *StatCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatTransactorSession struct {
	Contract     *StatTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatRaw struct {
	Contract *Stat // Generic contract binding to access the raw methods on
}

// StatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatCallerRaw struct {
	Contract *StatCaller // Generic read-only contract binding to access the raw methods on
}

// StatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatTransactorRaw struct {
	Contract *StatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStat creates a new instance of Stat, bound to a specific deployed contract.
func NewStat(address common.Address, backend bind.ContractBackend) (*Stat, error) {
	contract, err := bindStat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stat{StatCaller: StatCaller{contract: contract}, StatTransactor: StatTransactor{contract: contract}, StatFilterer: StatFilterer{contract: contract}}, nil
}

// NewStatCaller creates a new read-only instance of Stat, bound to a specific deployed contract.
func NewStatCaller(address common.Address, caller bind.ContractCaller) (*StatCaller, error) {
	contract, err := bindStat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatCaller{contract: contract}, nil
}

// NewStatTransactor creates a new write-only instance of Stat, bound to a specific deployed contract.
func NewStatTransactor(address common.Address, transactor bind.ContractTransactor) (*StatTransactor, error) {
	contract, err := bindStat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatTransactor{contract: contract}, nil
}

// NewStatFilterer creates a new log filterer instance of Stat, bound to a specific deployed contract.
func NewStatFilterer(address common.Address, filterer bind.ContractFilterer) (*StatFilterer, error) {
	contract, err := bindStat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatFilterer{contract: contract}, nil
}

// bindStat binds a generic wrapper to an already deployed contract.
func bindStat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StatMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stat *StatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stat.Contract.StatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stat *StatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.Contract.StatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stat *StatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stat.Contract.StatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stat *StatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stat *StatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stat *StatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stat.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stat *StatCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stat *StatSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stat.Contract.DEFAULTADMINROLE(&_Stat.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stat *StatCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stat.Contract.DEFAULTADMINROLE(&_Stat.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Stat *StatCaller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Stat *StatSession) GOVERNORROLE() ([32]byte, error) {
	return _Stat.Contract.GOVERNORROLE(&_Stat.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Stat *StatCallerSession) GOVERNORROLE() ([32]byte, error) {
	return _Stat.Contract.GOVERNORROLE(&_Stat.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Stat *StatCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Stat *StatSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Stat.Contract.UPGRADEINTERFACEVERSION(&_Stat.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Stat *StatCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Stat.Contract.UPGRADEINTERFACEVERSION(&_Stat.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Stat *StatCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Stat *StatSession) VERSION() (string, error) {
	return _Stat.Contract.VERSION(&_Stat.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Stat *StatCallerSession) VERSION() (string, error) {
	return _Stat.Contract.VERSION(&_Stat.CallOpts)
}

// CheckPiece is a free data retrieval call binding the contract method 0xc2e8e717.
//
// Solidity: function checkPiece(bytes _pn) view returns(bool)
func (_Stat *StatCaller) CheckPiece(opts *bind.CallOpts, _pn []byte) (bool, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "checkPiece", _pn)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPiece is a free data retrieval call binding the contract method 0xc2e8e717.
//
// Solidity: function checkPiece(bytes _pn) view returns(bool)
func (_Stat *StatSession) CheckPiece(_pn []byte) (bool, error) {
	return _Stat.Contract.CheckPiece(&_Stat.CallOpts, _pn)
}

// CheckPiece is a free data retrieval call binding the contract method 0xc2e8e717.
//
// Solidity: function checkPiece(bytes _pn) view returns(bool)
func (_Stat *StatCallerSession) CheckPiece(_pn []byte) (bool, error) {
	return _Stat.Contract.CheckPiece(&_Stat.CallOpts, _pn)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Stat *StatCaller) Epoch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Stat *StatSession) Epoch() (common.Address, error) {
	return _Stat.Contract.Epoch(&_Stat.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Stat *StatCallerSession) Epoch() (common.Address, error) {
	return _Stat.Contract.Epoch(&_Stat.CallOpts)
}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Stat *StatCaller) Eproof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "eproof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Stat *StatSession) Eproof() (common.Address, error) {
	return _Stat.Contract.Eproof(&_Stat.CallOpts)
}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Stat *StatCallerSession) Eproof() (common.Address, error) {
	return _Stat.Contract.Eproof(&_Stat.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stat *StatCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stat *StatSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Stat.Contract.GetRoleAdmin(&_Stat.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stat *StatCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Stat.Contract.GetRoleAdmin(&_Stat.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stat *StatCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stat *StatSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Stat.Contract.HasRole(&_Stat.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stat *StatCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Stat.Contract.HasRole(&_Stat.CallOpts, role, account)
}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_Stat *StatCaller) Piece(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "piece")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_Stat *StatSession) Piece() (common.Address, error) {
	return _Stat.Contract.Piece(&_Stat.CallOpts)
}

// Piece is a free data retrieval call binding the contract method 0x95ad21dc.
//
// Solidity: function piece() view returns(address)
func (_Stat *StatCallerSession) Piece() (common.Address, error) {
	return _Stat.Contract.Piece(&_Stat.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Stat *StatCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Stat *StatSession) ProxiableUUID() ([32]byte, error) {
	return _Stat.Contract.ProxiableUUID(&_Stat.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Stat *StatCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Stat.Contract.ProxiableUUID(&_Stat.CallOpts)
}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Stat *StatCaller) Rsproof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "rsproof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Stat *StatSession) Rsproof() (common.Address, error) {
	return _Stat.Contract.Rsproof(&_Stat.CallOpts)
}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Stat *StatCallerSession) Rsproof() (common.Address, error) {
	return _Stat.Contract.Rsproof(&_Stat.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stat *StatCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Stat.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stat *StatSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stat.Contract.SupportsInterface(&_Stat.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stat *StatCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stat.Contract.SupportsInterface(&_Stat.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stat *StatTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stat *StatSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stat.Contract.GrantRole(&_Stat.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stat *StatTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stat.Contract.GrantRole(&_Stat.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _epoch, address _piece, address _rsproof, address _eproof, address initialOwner) returns()
func (_Stat *StatTransactor) Initialize(opts *bind.TransactOpts, _epoch common.Address, _piece common.Address, _rsproof common.Address, _eproof common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "initialize", _epoch, _piece, _rsproof, _eproof, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _epoch, address _piece, address _rsproof, address _eproof, address initialOwner) returns()
func (_Stat *StatSession) Initialize(_epoch common.Address, _piece common.Address, _rsproof common.Address, _eproof common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Stat.Contract.Initialize(&_Stat.TransactOpts, _epoch, _piece, _rsproof, _eproof, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _epoch, address _piece, address _rsproof, address _eproof, address initialOwner) returns()
func (_Stat *StatTransactorSession) Initialize(_epoch common.Address, _piece common.Address, _rsproof common.Address, _eproof common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Stat.Contract.Initialize(&_Stat.TransactOpts, _epoch, _piece, _rsproof, _eproof, initialOwner)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Stat *StatTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Stat *StatSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Stat.Contract.RenounceRole(&_Stat.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Stat *StatTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Stat.Contract.RenounceRole(&_Stat.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stat *StatTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stat *StatSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stat.Contract.RevokeRole(&_Stat.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stat *StatTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stat.Contract.RevokeRole(&_Stat.TransactOpts, role, account)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Stat *StatTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Stat.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Stat *StatSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Stat.Contract.UpgradeToAndCall(&_Stat.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Stat *StatTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Stat.Contract.UpgradeToAndCall(&_Stat.TransactOpts, newImplementation, data)
}

// StatInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Stat contract.
type StatInitializedIterator struct {
	Event *StatInitialized // Event containing the contract specifics and raw log

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
func (it *StatInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatInitialized)
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
		it.Event = new(StatInitialized)
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
func (it *StatInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatInitialized represents a Initialized event raised by the Stat contract.
type StatInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Stat *StatFilterer) FilterInitialized(opts *bind.FilterOpts) (*StatInitializedIterator, error) {

	logs, sub, err := _Stat.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StatInitializedIterator{contract: _Stat.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Stat *StatFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StatInitialized) (event.Subscription, error) {

	logs, sub, err := _Stat.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatInitialized)
				if err := _Stat.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Stat *StatFilterer) ParseInitialized(log types.Log) (*StatInitialized, error) {
	event := new(StatInitialized)
	if err := _Stat.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Stat contract.
type StatRoleAdminChangedIterator struct {
	Event *StatRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StatRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatRoleAdminChanged)
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
		it.Event = new(StatRoleAdminChanged)
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
func (it *StatRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatRoleAdminChanged represents a RoleAdminChanged event raised by the Stat contract.
type StatRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stat *StatFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StatRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Stat.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StatRoleAdminChangedIterator{contract: _Stat.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stat *StatFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StatRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Stat.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatRoleAdminChanged)
				if err := _Stat.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Stat *StatFilterer) ParseRoleAdminChanged(log types.Log) (*StatRoleAdminChanged, error) {
	event := new(StatRoleAdminChanged)
	if err := _Stat.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Stat contract.
type StatRoleGrantedIterator struct {
	Event *StatRoleGranted // Event containing the contract specifics and raw log

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
func (it *StatRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatRoleGranted)
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
		it.Event = new(StatRoleGranted)
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
func (it *StatRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatRoleGranted represents a RoleGranted event raised by the Stat contract.
type StatRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stat *StatFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StatRoleGrantedIterator, error) {

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

	logs, sub, err := _Stat.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StatRoleGrantedIterator{contract: _Stat.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stat *StatFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StatRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Stat.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatRoleGranted)
				if err := _Stat.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Stat *StatFilterer) ParseRoleGranted(log types.Log) (*StatRoleGranted, error) {
	event := new(StatRoleGranted)
	if err := _Stat.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Stat contract.
type StatRoleRevokedIterator struct {
	Event *StatRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StatRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatRoleRevoked)
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
		it.Event = new(StatRoleRevoked)
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
func (it *StatRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatRoleRevoked represents a RoleRevoked event raised by the Stat contract.
type StatRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stat *StatFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StatRoleRevokedIterator, error) {

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

	logs, sub, err := _Stat.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StatRoleRevokedIterator{contract: _Stat.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stat *StatFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StatRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Stat.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatRoleRevoked)
				if err := _Stat.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Stat *StatFilterer) ParseRoleRevoked(log types.Log) (*StatRoleRevoked, error) {
	event := new(StatRoleRevoked)
	if err := _Stat.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Stat contract.
type StatUpgradedIterator struct {
	Event *StatUpgraded // Event containing the contract specifics and raw log

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
func (it *StatUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatUpgraded)
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
		it.Event = new(StatUpgraded)
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
func (it *StatUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatUpgraded represents a Upgraded event raised by the Stat contract.
type StatUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Stat *StatFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StatUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Stat.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StatUpgradedIterator{contract: _Stat.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Stat *StatFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StatUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Stat.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatUpgraded)
				if err := _Stat.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Stat *StatFilterer) ParseUpgraded(log types.Log) (*StatUpgraded, error) {
	event := new(StatUpgraded)
	if err := _Stat.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
