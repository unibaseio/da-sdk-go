// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package piece

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

// IPieceInitParams is an auto generated low-level Go binding around an user-defined struct.
type IPieceInitParams struct {
	Delay       uint64
	MinStore    uint64
	MaxStore    uint64
	MaxSize     uint64
	MinPrice    *big.Int
	StreamPrice *big.Int
}

// IPiecePieceInfo is an auto generated low-level Go binding around an user-defined struct.
type IPiecePieceInfo struct {
	Rsn      uint8
	Rsk      uint8
	Size     uint64
	Start    uint64
	Expire   uint64
	Owner    common.Address
	Streamer common.Address
	Price    *big.Int
	Balance  *big.Int
}

// IPieceReplicaInfo is an auto generated low-level Go binding around an user-defined struct.
type IPieceReplicaInfo struct {
	Piece    uint64
	StoredOn common.Address
	Root     [32]byte
}

// IPieceStoreInfo is an auto generated low-level Go binding around an user-defined struct.
type IPieceStoreInfo struct {
	Epoch   uint64
	Count   uint64
	Salary  *big.Int
	Revenue *big.Int
}

// IPieceStoreStat is an auto generated low-level Go binding around an user-defined struct.
type IPieceStoreStat struct {
	Count  uint64
	Salary *big.Int
}

// PieceMetaData contains all meta data concerning the Piece contract.
var PieceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAYER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addPiece\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_expire\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rsn\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rsk\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_s\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPieceFor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_expire\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rsn\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rsk\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_s\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addReplica\",\"inputs\":[{\"name\":\"_rn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkSReplica\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_com\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkStore\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEpoch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPIndex\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPRI\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPRInfo\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPiece\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"_pinfo\",\"type\":\"tuple\",\"internalType\":\"structIPiece.PieceInfo\",\"components\":[{\"name\":\"rsn\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rsk\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"start\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expire\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"streamer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRIndex\",\"inputs\":[{\"name\":\"_rn\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRS\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReplica\",\"inputs\":[{\"name\":\"_r\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"_ri\",\"type\":\"tuple\",\"internalType\":\"structIPiece.ReplicaInfo\",\"components\":[{\"name\":\"piece\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"storedOn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevenue\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSRAt\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSStat\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"_si\",\"type\":\"tuple\",\"internalType\":\"structIPiece.StoreStat\",\"components\":[{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"salary\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStore\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_si\",\"type\":\"tuple\",\"internalType\":\"structIPiece.StoreInfo\",\"components\":[{\"name\":\"epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"salary\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStoreCount\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStoreSalary\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_epoch\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIPiece.InitParams\",\"components\":[{\"name\":\"delay\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minStore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxStore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSize\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"streamPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxStore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"node\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"replicaCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reserveOf\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"retake\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settle\",\"inputs\":[{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"streamPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddPiece\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddPieceV2\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_pn\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"_price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_size\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_expire\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_rsn\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_rsk\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_s\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddReplica\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_sri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddReplicaV2\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_sri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_rn\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Retake\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Settle\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a08060405234620000d157306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c16620000c257506001600160401b036002600160401b0319828216016200007c575b6040516139389081620000d6823960805181818161136f015261146d0152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80806200005c565b63f92ee8a960e01b8152600490fd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a71461259557508063123f84591461233057806318198ad71461228f578063194cbaf514611a7f5780631ce85e7c14611a30578063248a9ca3146119f75780632565b159146119cd5780632620904f146118b85780632e1a7d4d146117c05780632f2ff15d1461177657806336568abe1461172e5780634b4ffccd146116845780634f1ef2861461142e57806350979d7e146113d957806352d1902d1461135c57806362b98032146112915780636a42b8f8146112675780636aadfac71461120157806375011aba146111e35780637a87f0c41461104b5780637cf1ee9114610c7457806385b6709514610c3b5780638879512514610c14578063894d35f814610b2d578063898e67f114610b015780638df82800146108775780638ef0e895146107ef578063900cf0cf146107c65780639108544c1461073257806391d14854146106dc578063926d7d7f146106a1578063a1e030d91461067a578063a217fddf1461065e578063ad3cb1cc14610618578063b650b50414610545578063ba8b2618146104e3578063bcbfd56b1461048e578063c3aa3804146103f5578063ccc57490146103ba578063d547741f1461036c578063d70754ec14610343578063e45be8eb14610325578063ee91365b146102b5578063fc0c546a1461028e578063fca5be7a146102645763ffa1ad7414610218575f80fd5b3461026157806003193601126102615761025d604051610237816126d1565b60058152640322e302e360dc1b60208201526040519182916020835260208301906127a1565b0390f35b80fd5b503461026157806003193601126102615760206001600160401b0360035460401c16604051908152f35b5034610261578060031936011261026157546040516001600160a01b039091168152602090f35b5034610261576040366003190112610261576102cf6125e8565b60243560ff811680910361032157604080936001600160401b038094168152600860205220905f52602052815f20541661030881612959565b50548251918252821c6001600160a01b03166020820152f35b5f80fd5b50346102615780600319360112610261576020600454604051908152f35b50346102615780600319360112610261576002546040516001600160a01b039091168152602090f35b5034610261576040366003190112610261576103b660043561038c612656565b908084525f805160206138c38339815191526020526103b16001604086200154612e05565b613622565b5080f35b503461026157806003193601126102615760206040517f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f558152f35b503461026157604036600319011261026157604090610412612640565b61041a6125fe565b83805192610427846126d1565b8484528460208095015260018060a01b031693848152600e8452818120926001600160401b0380911693845f528552825f205416948152600e845220905f5281526001835f20015481845161047b816126d1565b8481520190815283519283525190820152f35b50346102615760403660031901126102615760209060406104ad612640565b916104b66125fe565b9260018060a01b03168152600d8452206001600160401b038092165f52825260405f205416604051908152f35b503461026157602036600319011261026157600435906001600160401b03908183116102615750610531602061051e81943690600401612752565b8160405193828580945193849201612780565b8101600b8152030190205416604051908152f35b50346102615760403660031901126102615760a06105616125e8565b6105696125fe565b9061057381612831565b50916105b66105b16105926001600160401b0380965460101c166128ea565b6105ab60ff6105a087612831565b505460081c16612cf0565b90612d0b565b6128aa565b9260ff6105c284612831565b5054169360016105f76105e560ff6105d988612831565b505460081c1696612831565b50600283808b1b039101541694612959565b50015493604051958652602086015216604084015260608301526080820152f35b503461026157806003193601126102615761025d604051610638816126d1565b60058152640352e302e360dc1b60208201526040519182916020835260208301906127a1565b5034610261578060031936011261026157602090604051908152f35b503461026157806003193601126102615760206001600160401b03600a5416604051908152f35b503461026157806003193601126102615760206040517fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc48152f35b50346102615760403660031901126102615760406106f8612656565b9160043581525f805160206138c3833981519152602052209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b503461026157602036600319011261026157606061074e6125e8565b61075661293b565b506001600160401b0361076882612959565b5054169061077581612959565b505460401c6001600160a01b039081169160019061079290612959565b50015491604080516107a3816126b6565b858152602081019283520192835260405193845251166020830152516040820152f35b50346102615780600319360112610261576001546040516001600160a01b039091168152602090f35b50346102615760e0366003190112610261576001600160401b0360043581811161087357610821903690600401612752565b90604435908116810361032157610836612614565b60843560ff811681036103215761084b612770565b60c435939092906001600160a01b038516850361032157610870956024359033612e33565b80f35b8280fd5b503461026157602080600319360112610ac3576002546040516327c830a960e01b81526004803593926001600160a01b0392918491839190829086165afa8015610a68578590610ac7575b6108cd9150156127d3565b838160015416803b15610ac35781809160046040518094819363919840ad60e01b83525af18015610ab857610aa0575b5050600482826001541660405192838092639fa6a6e360e01b82525afa8015610a6857610932918691610a73575b50336136a1565b6044604082600254168151928380926361e728b160e01b8252336004830152600160248301525afa908115610a685785908692610a36575b5015610a0657338552600c8352600160408620015490816007029160078304036109f2577fa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d20493926109c16109e69387931015612a30565b338752600c8452600260408820016109da8382546128dd565b90553390875416612d37565b6040519283523392a280f35b634e487b7160e01b86526011600452602486fd5b60405162461bcd60e51b8152602060048201526008602482015267696e76206e6f646560c01b6044820152606490fd5b9050610a5a915060403d604011610a61575b610a5281836126ec565b810190612a13565b905f61096a565b503d610a48565b6040513d87823e3d90fd5b610a939150843d8611610a99575b610a8b81836126ec565b810190612812565b5f61092b565b503d610a81565b610aa99061266c565b610ab457835f6108fd565b8380fd5b6040513d84823e3d90fd5b5080fd5b508281813d8311610afa575b610add81836126ec565b81010312610af657610af16108cd916127c6565b6108c2565b8480fd5b503d610ad3565b5034610261576020366003190112610261576020610b25610b206125e8565b612ba5565b604051908152f35b503461026157606036600319011261026157610b47612640565b90610b506125fe565b90604435926001600160401b039283851161087357610b756040953690600401612752565b9160018060a01b03168352836020958694600d86522091165f52825282610bb28160405f2054169384938160405193828580945193849201612780565b8101600b815203019020541603610be65781610bd0610bd892612959565b505416612831565b505460901c16604051908152f35b60405162461bcd60e51b815260048101849052600660248201526534b73b1039b960d11b6044820152606490fd5b503461026157806003193601126102615760206001600160401b0360035416604051908152f35b5034610261576020366003190112610261576020906040906001600160a01b03610c63612640565b168152600f83522054604051908152f35b50346102615761014036600319011261026157610c8f612640565b610c97612656565b604435906001600160a01b03821682036103215760c0366063190112610ab45760405160c081018181106001600160401b0382111761103757604052610cdb612614565b8152610ce561262a565b602082015260a4356001600160401b038116810361032157604082015260c4356001600160401b038116810361032157606082015260e43560808201526101043560a0820152610124356001600160a01b0381169003610321575f805160206138e383398151915254926001600160401b03841680159081611027575b600114908161101d575b159081611014575b506110025760016001600160401b03198516175f805160206138e38339815191525560ff8460401c1615610fd6575b60ff5f805160206138e38339815191525460401c1615610fc4576001600160401b038251166001600160401b036020840151161115610f8b5760ff9460a0936bffffffffffffffffffffffff851b91600180871b031682895416178855600180861b03169060015416176001556002546001600160401b03841b8351851b1691600180861b03169063ffffffff60e01b1617176002556001600160401b036020820151166003549067ffffffffffffffff60401b604084015160401b16906001600160401b0360801b606085015160801b16926001600160401b0360c01b1617171760035560808101516004550151600555610ea1610124356133d5565b50610eae6101243561345f565b50610ebb6101243561350c565b50610f04604051610ecb8161269a565b8481528460208201528460408201528460608201528460808201528460a08201528460c08201528460e082015284610100820152612aaf565b610f27604051610f13816126b6565b848152846020820152846040820152612991565b60401c1615610f335780f35b68ff0000000000000000195f805160206138e383398151915254165f805160206138e3833981519152557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b60405162461bcd60e51b815260206004820152601160248201527064656c6179203e3d206d696e53746f726560781b6044820152606490fd5b604051631afcd79f60e31b8152600490fd5b68ffffffffffffffffff19841668010000000000000001175f805160206138e383398151915255610da3565b60405163f92ee8a960e01b8152600490fd5b9050155f610d74565b303b159150610d6c565b604086901c60ff16159150610d62565b634e487b7160e01b5f52604160045260245ffd5b5034610261576020366003190112610261576101206110686125e8565b611070612a6c565b5061010060ff61107f83612831565b5054169160ff61108e82612831565b505460081c16906001600160401b036110a682612831565b505460101c16906001600160401b036110be82612831565b505460501c166001600160401b036110d583612831565b505460901c166110e483612831565b50600101546001600160a01b0316906110fc84612831565b50600201546001600160a01b0316926004611124600361111b88612831565b50015496612831565b50015495604051976111358961269a565b8a8952602089015260408801526060870152608086015260a085015260c084015260e08301528282015260405192835260ff60208201511660208401526001600160401b0360408201511660408401526001600160401b0360608201511660608401526001600160401b03608082015116608084015260018060a01b0360a08201511660a084015260018060a01b0360c08201511660c084015260e081015160e08401520151610100820152f35b50346102615780600319360112610261576020600554604051908152f35b503461026157602036600319011261026157606061121d6125e8565b60ff61122882612831565b505416906001600160401b0361124e60ff61124284612831565b505460081c1692612831565b505460901c169060405192835260208301526040820152f35b503461026157806003193601126102615760206001600160401b0360025460a01c16604051908152f35b5034610261576020366003190112610261576112ab612640565b6001546001600160a01b03919083908316803b15610ac35781809160046040518094819363919840ad60e01b83525af18015610ab857611348575b505060206004926001541660405193848092639fa6a6e360e01b82525afa90811561133d5761087092849261131c575b506136a1565b61133691925060203d602011610a9957610a8b81836126ec565b905f611316565b6040513d85823e3d90fd5b6113519061266c565b61087357825f6112e6565b50346102615780600319360112610261577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036113c75760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b60405163703e46dd60e11b8152600490fd5b50346102615760403660031901126102615760209060406113f8612640565b916114016125fe565b9260018060a01b03168152600e8452206001600160401b038092165f52825260405f205416604051908152f35b50604036600319011261026157611443612640565b6024356001600160401b03811161087357611462903690600401612752565b6001600160a01b03917f00000000000000000000000000000000000000000000000000000000000000008316308114908115611656575b506113c757335f9081527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360209081526040909120549093907f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f559060ff161561163857506040516352d1902d60e01b8152908216938082600481885afa9182918793611604575b505061153f57604051634c9c8ce360e01b815260048101859052602490fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc908181036115ec5750813b156115d35780546001600160a01b03191684179055604051927fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8580a28151156115b9576103b69250613811565b5050346115c4575080f35b63b398979f60e01b8152600490fd5b604051634c9c8ce360e01b815260048101859052602490fd5b60249060405190632a87526960e21b82526004820152fd5b9080929350813d8311611631575b61161c81836126ec565b8101031261162d5751905f80611520565b8580fd5b503d611612565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b9050837f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc541614155f611499565b50346102615760203660031901126102615760406080916116a3612640565b81606084516116b18161267f565b8281528260208201528286820152015260018060a01b03168152600c6020522080546001600160401b039182821692600260018301549201549260606040516116f98161267f565b86815283602082019360401c168352604081019485520193845260405194855251166020840152516040830152516060820152f35b503461026157604036600319011261026157611748612656565b336001600160a01b03821603611764576103b690600435613622565b60405163334bd91960e11b8152600490fd5b5034610261576040366003190112610261576103b6600435611796612656565b908084525f805160206138c38339815191526020526117bb6001604086200154612e05565b6135b2565b503461026157602080600319360112610ac3576002546040516327c830a960e01b8152600480359391926001600160a01b03918391859190829085165afa8015610a685784938691611854575b50926109e69161183e7f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436495156127d3565b338752600f8452604087206109da8382546128dd565b809450838092503d83116118b1575b61186d81836126ec565b81010312610af657836109e69161183e6118a77f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364966127c6565b929550509161180d565b503d611863565b503461026157610100366003190112610261576118d3612640565b6024356001600160401b038111610873576118f2903690600401612752565b906118fb612614565b61190361262a565b61190b612770565b9060c4359260ff841684036103215760e435946001600160a01b03808716870361032157335f9081527fa9ed30483ca8e7e18edb58c7d68a44612fe881c88cebad0b52b41698e9844de260205260409020547fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc49060ff1615611638575081161561199c576108709660443591612e33565b60405162461bcd60e51b815260206004820152600960248201526834b73b1037bbb732b960b91b6044820152606490fd5b503461026157806003193601126102615760206001600160401b0360035460801c16604051908152f35b5034610261576020366003190112610261576001604060209260043581525f805160206138c38339815191528452200154604051908152f35b503461026157602036600319011261026157600435906001600160401b03908183116102615750611a6b602061051e81943690600401612752565b810160078152030190205416604051908152f35b5034610261576080366003190112610261576004356001600160401b038111610ac357611ab0903690600401612752565b611ab86125fe565b9060443560ff8116810361032157606435926001600160401b038411610af65736602385011215610af6576001600160401b03846004013511610af657366024856004013586010111610af65760015485906001600160a01b0316803b15610ac35781809160046040518094819363919840ad60e01b83525af18015610ab85761227b575b5050600154604051639fa6a6e360e01b81529290602090849060049082906001600160a01b03165afa92831561227057869361224f575b5083511561221c576004611b8783612831565b500154156121ec576001600160401b03611ba083612831565b505460901c166001600160401b03841610156121bd57611bbf82612831565b5060ff808316915416111561218e576001600160401b036040516020818751611beb8183858c01612780565b8101600b815203019020541661215f576001600160401b038216865260086020526040862060ff82165f526020526001600160401b0360405f20541661212e576001600160401b0382168652600960205260408620335f5260205260ff60405f2054166120fe57611c5b82612db5565b92611ca8611c986003611c6d86612831565b5001546001600160401b03611c918582611c868a612831565b505460901c16612902565b169061291b565b6001600160401b0386169061291b565b6004611cb385612831565b500154611cc26005548361292e565b116120cc57611cdf6004611cd586612831565b50019182546128dd565b9055600554611cf26004611cd586612831565b9055600554611d0084612831565b50600201546001600160a01b03168852600f6020526040882080549091611d269161292e565b9055611d3061293b565b936001600160401b0384168552336020860152611d5536886004013560248a0161270d565b602081519101206040860152611d766001600160401b03600a541695612991565b611e1f60405160208181611d908b83815193849201612780565b8101600b815203019020916001600160401b03199287848254161790556001600160401b0386168a52600860205260408a2060ff86165f5260205260405f2087848254161790556001600160401b0386168a52600960205260408a20335f5260205260405f20600160ff1982541617905560036001600160401b03611e1488612831565b50921691015461291b565b916001600160401b03611e3186612831565b505460901c1691611e4282336136a1565b338a52600d60205260408a20600c6020526001600160401b0360408c205460401c165f5260205260405f208782825416179055338a52600c60205260408a20805467ffffffffffffffff60401b611ea46001600160401b038360401c1661287e565b60401b169067ffffffffffffffff60401b1916179055338a52600c602052600160408b2001611ed485825461292e565b9055338a52600e60205260408a206001600160401b0383165f5260205260405f20908154906001600160401b03611f0c81841661287e565b169116179055338952600e6020526001600160401b0360408a2091165f52602052600160405f2001611f3f83825461292e565b9055338852600e60205260408820905f52602052611f65600160405f200191825461292e565b9055600254604080516361e728b160e01b81523360048201526001602482015291829060449082906001600160a01b03165afa9081156120c1578790889261209f575b5015610a0657338752600c6020526001604088200154908160070291600783040361208b577ff226a06b14819931288115f369b127b1a9986ea785f3a4cfcf16b0b445cf1ca59594926120026020959360ff931015612a30565b338952600c85526001600160401b036120466120268260408d205460401c166128ea565b9782604051998a988952168888015260c0604088015260c08701906127a1565b9316606085015216608083015281810360a08301526004860135808252602487018483013786838760040135830101523395601f80199160040135011601030190a280f35b634e487b7160e01b88526011600452602488fd5b90506120ba915060403d604011610a6157610a5281836126ec565b905f611fa8565b6040513d89823e3d90fd5b60405162461bcd60e51b815260206004820152600a6024820152696e6f2062616c616e636560b01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526008602482015267323ab81030b2323960c11b6044820152606490fd5b60405162461bcd60e51b815260206004820152600960248201526865786973742070726960b81b6044820152606490fd5b60405162461bcd60e51b815260206004820152600760248201526632bc34b9ba103960c91b6044820152606490fd5b60405162461bcd60e51b8152602060048201526007602482015266696e762070726960c81b6044820152606490fd5b60405162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b6044820152606490fd5b60405162461bcd60e51b81526020600482015260086024820152676e6f20706965636560c01b6044820152606490fd5b60405162461bcd60e51b815260206004820152600b60248201526a696e76207265706c69636160a81b6044820152606490fd5b61226991935060203d602011610a9957610a8b81836126ec565b915f611b74565b6040513d88823e3d90fd5b6122849061266c565b610af657845f611b3d565b5034610261576040366003190112610261576122a9612640565b906122b26125fe565b9160018060a01b031690818152600c6020526001600160401b038060408320541693168093106122ff576040918152600e60205220905f526020526020600160405f200154604051908152f35b60405162461bcd60e51b81526020600482015260096024820152680d2dcec40cae0dec6d60bb1b6044820152606490fd5b5034610321576020806003193601126103215761234b6125e8565b6002546040516327c830a960e01b815291926001600160a01b0392909182908290600490829087165afa8015612554575f9061255f575b61238d9150156127d3565b8160015416803b15610321575f809160046040518094819363919840ad60e01b83525af1801561255457612541575b50600481836001541660405192838092639fa6a6e360e01b82525afa908115610a68578591612524575b508260016123f386612831565b5001541633036124f65761242b9061240a85612831565b506001600160401b0380938180935460501c168260025460a01c16906128c2565b9216911610156124c957600461244085612831565b50015461244c85612ba5565b90818111156124bc5761248991612462916128dd565b80945b600461247088612831565b500161247d8382546128dd565b90553390885416612d37565b604051931683528201527f44647ea87405778ba90a8ca28d2b772b1350bcd5950e5111bbc8c506a5c9632a60403392a280f35b5050612489858094612465565b60405162461bcd60e51b81526004810183905260056024820152646561726c7960d81b6044820152606490fd5b60405162461bcd60e51b8152600481018390526006602482015265696e7620776360d01b6044820152606490fd5b61253b9150823d8411610a9957610a8b81836126ec565b5f6123e6565b61254c91945061266c565b5f925f6123bc565b6040513d5f823e3d90fd5b508181813d831161258e575b61257581836126ec565b810103126103215761258961238d916127c6565b612382565b503d61256b565b34610321576020366003190112610321576004359063ffffffff60e01b821680920361032157602091637965db0b60e01b81149081156125d7575b5015158152f35b6301ffc9a760e01b149050836125d0565b600435906001600160401b038216820361032157565b602435906001600160401b038216820361032157565b606435906001600160401b038216820361032157565b608435906001600160401b038216820361032157565b600435906001600160a01b038216820361032157565b602435906001600160a01b038216820361032157565b6001600160401b03811161103757604052565b608081019081106001600160401b0382111761103757604052565b61012081019081106001600160401b0382111761103757604052565b606081019081106001600160401b0382111761103757604052565b604081019081106001600160401b0382111761103757604052565b90601f801991011681019081106001600160401b0382111761103757604052565b9291926001600160401b0382116110375760405191612736601f8201601f1916602001846126ec565b829481845281830111610321578281602093845f960137010152565b9080601f830112156103215781602061276d9335910161270d565b90565b60a4359060ff8216820361032157565b5f5b8381106127915750505f910152565b8181015183820152602001612782565b906020916127ba81518092818552858086019101612780565b601f01601f1916010190565b5190811515820361032157565b156127da57565b60405162461bcd60e51b815260206004820152601060248201526f195b595c99d95b98de481c185d5cd95960821b6044820152606490fd5b9081602091031261032157516001600160401b03811681036103215790565b60065481101561286a5760059060065f52027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01905f90565b634e487b7160e01b5f52603260045260245ffd5b9060016001600160401b038093160191821161289657565b634e487b7160e01b5f52601160045260245ffd5b906001600160401b0380921660010191821161289657565b9190916001600160401b038080941691160191821161289657565b9190820391821161289657565b6001600160401b039081165f19019190821161289657565b6001600160401b03918216908216039190821161289657565b8181029291811591840414171561289657565b9190820180921161289657565b60405190612948826126b6565b5f6040838281528260208201520152565b600a5481101561286a57600a5f5260011b7fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a801905f90565b600a54600160401b811015611037578060016129b09201600a55612959565b919091612a00578051825460208301516001600160e01b03199091166001600160401b039290921691909117604091821b68010000000000000000600160e01b0316178355015160019190910155565b634e487b7160e01b5f525f60045260245ffd5b9190826040910312610321576020612a2a836127c6565b92015190565b15612a3757565b60405162461bcd60e51b815260206004820152600d60248201526c696e73756620706c656467653760981b6044820152606490fd5b60405190612a798261269a565b5f610100838281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e08201520152565b600654600160401b81101561103757806001612ace9201600655612831565b919091612a00578051825460208301516040840151606085015160808601516001600160d01b031990941660ff959095169490941760089290921b61ff00169190911760109190911b69ffffffffffffffff0000161760509290921b67ffffffffffffffff60501b169190911760909190911b67ffffffffffffffff60901b1617825560a08101516001830180546001600160a01b03199081166001600160a01b039384161790915560c0830151600285018054909216921691909117905560e08101516003830155610100015160049190910155565b60019060018060a01b036001541660409260405190639fa6a6e360e01b825260209460208360048181975afa928315612554575f93612ccf575b50612be985612831565b50936001600160401b0380955460901c168585161015612cc457905f96835f905b612c59575b50505050508282611c91612c44612c4a94612c5396611c916003612c3661276d9c9b612831565b5001549282611c8689612831565b92612db5565b6005549061292e565b9061291b565b612c6288612831565b505460ff8281169116811015612cbe578789165f5260088352835f20905f52825286835f20541615612c9a575b840160ff1684612c0a565b975f198114612cab57840197612c8f565b601184634e487b7160e01b5f525260245ffd5b50612c0f565b505050505050505f90565b612ce991935060203d602011610a9957610a8b81836126ec565b915f612bdf565b906001600160401b03809216601f0291821691820361289657565b906001600160401b03809116918215612d2357160490565b634e487b7160e01b5f52601260045260245ffd5b60405163a9059cbb60e01b5f9081526001600160a01b039384166004526024949094529260209060448180855af160015f5114811615612d96575b8360405215612d8057505050565b635274afe760e01b835216600482015260249150fd5b6001811516612dac57813b15153d151616612d72565b833d5f823e3d90fd5b6601ffffffffffff612dfc612df76105b161276d946105ab60ff6105a0612df16001600160401b03612de686612831565b505460101c166128ea565b93612831565b6128ea565b600f1c166128aa565b805f525f805160206138c383398151915260205260405f20335f5260205260ff60405f205416156116385750565b60015493975f97939692959094919392916001600160a01b0316803b15610321575f809160046040518094819363919840ad60e01b83525af18015612554576133bb575b50600154604051639fa6a6e360e01b815293929190602090859060049082906001600160a01b03165afa9384156133b057899461338f575b5086511561335e576001600160401b038a161561332e5760ff8316156132ff5760ff821660ff841610156132d1576003546001600160401b03612ef4818316876128c2565b166001600160401b038316106132a0576001600160401b03612f1b818360401c16876128c2565b166001600160401b0383161161326e57600454891061323d576001600160401b0360405160208181612f538d83815193849201612780565b81016007815203019020541661320a57612f7e6105b1612f728d6128ea565b6105ab60ff8816612cf0565b9060ff85166001600160401b03831602906001600160401b0382169182036128965760801c6001600160401b0316106131d757612fba906128ea565b600f1c6601ffffffffffff16612fcf906128aa565b612fd98583612902565b6001600160401b0316612fec908a61291b565b906001600160401b0316612fff9161291b565b60ff8316906005546130109161292e565b9061301a9161291b565b996006546001600160401b031697613030612a6c565b6001600160a01b0388811660a083015261010082018e905260e082018c90526001600160401b03848116604084015297881660608301819052978516608083015260ff8681168352871660208301529890981660c08901819052959761309590612aaf565b604051815181906130aa818360208701612780565b810160078152036020019020896001600160401b0319825416179055604051988998610120918a5260208a01528060408a015288016130e8916127a1565b9860608801526001600160401b031660808701526001600160401b031660a086015260ff1660c085015260ff1660e0840152610100830152600160a01b600190031692037fe37588ecbc243b6bdc3191c8d9743ce41ab2f677eea6aafbd2ce4cd59480bbdd91a280546040516323b872dd60e01b835233600452306024526044939093526001600160a01b031690808083815a93606492602095f19080516001148216156131b6575b83604052606052156131a1575050565b635274afe760e01b8252600482015260249150fd5b9060018115166131ce57823b15153d15161690613191565b833d83823e3d90fd5b60405162461bcd60e51b815260206004820152600b60248201526a6578636565642073697a6560a81b6044820152606490fd5b60405162461bcd60e51b815260206004820152600b60248201526a1c1a5958d948195e1a5cdd60aa1b6044820152606490fd5b60405162461bcd60e51b81526020600482015260096024820152686d696e20707269636560b81b6044820152606490fd5b60405162461bcd60e51b815260206004820152600a60248201526932bc31b2b2b210323ab960b11b6044820152606490fd5b60405162461bcd60e51b815260206004820152600960248201526839b437b93a10323ab960b91b6044820152606490fd5b60405162461bcd60e51b8152602060048201526006602482015265696e7620727360d01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526007602482015266696e762072736b60c81b6044820152606490fd5b60405162461bcd60e51b8152602060048201526008602482015267696e762073697a6560c01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526009602482015268696e7620706965636560b81b6044820152606490fd5b6133a991945060203d602011610a9957610a8b81836126ec565b925f612eaf565b6040513d8b823e3d90fd5b600493929198506133cb9061266c565b5f97909192612e77565b6001600160a01b03165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f805160206138c38339815191529060ff16613459575f805260205260405f20815f5260205260405f20600160ff1982541617905533905f5f805160206138a38339815191528180a4600190565b50505f90565b6001600160a01b03165f8181527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5591905f805160206138c38339815191529060ff1661350557825f5260205260405f20815f5260205260405f20600160ff1982541617905533915f805160206138a38339815191525f80a4600190565b5050505f90565b6001600160a01b03165f8181527fa9ed30483ca8e7e18edb58c7d68a44612fe881c88cebad0b52b41698e9844de260205260409020547fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc491905f805160206138c38339815191529060ff1661350557825f5260205260405f20815f5260205260405f20600160ff1982541617905533915f805160206138a38339815191525f80a4600190565b90815f525f805160206138c38339815191528060205260405f209160018060a01b031691825f5260205260ff60405f205416155f1461350557825f5260205260405f20815f5260205260405f20600160ff1982541617905533915f805160206138a38339815191525f80a4600190565b90815f525f805160206138c38339815191528060205260405f209160018060a01b031691825f5260205260ff60405f2054165f1461350557825f5260205260405f20815f5260205260405f2060ff19815416905533917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b91909160018060a01b0316805f52600c9260208481526040805f209586546001600160401b0380988282809416156137fa575b505416941693848110156137f0576136eb9061287e565b878116908582116137d257865f52600e808652845f20835f528652600180865f200154895f5285885261372482885f20019182546128dd565b9055885f5284875280865f200154828852865f20855f52885281875f200155885f528487528a865f2054871c16828852865f20855f528852865f20906001600160401b03198254161790556002918285101561378d575b505050613788915061287e565b6136eb565b895f528752855f209360011901908b821161289657613788948c6137c893165f528852865f20015491895f52858852865f200191825461292e565b90555f808061377b565b505092949550925f52525f20906001600160401b0319825416179055565b5050505050509050565b67ffffffffffffffff19168783161781555f6136d4565b905f8091602081519101845af4808061388f575b156138455750506040513d81523d5f602083013e60203d82010160405290565b1561386c57604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b3d1561387d576040513d5f823e3d90fd5b60405163d6bda27560e01b8152600490fd5b503d1515806138255750813b151561382556fe2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220f7220138e703be91d0d5e7eae93467a09ac099d009bc33c99bbf66a8fcdb64b964736f6c63430008180033",
}

// PieceABI is the input ABI used to generate the binding from.
// Deprecated: Use PieceMetaData.ABI instead.
var PieceABI = PieceMetaData.ABI

// PieceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PieceMetaData.Bin instead.
var PieceBin = PieceMetaData.Bin

// DeployPiece deploys a new Ethereum contract, binding an instance of Piece to it.
func DeployPiece(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Piece, error) {
	parsed, err := PieceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PieceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Piece{PieceCaller: PieceCaller{contract: contract}, PieceTransactor: PieceTransactor{contract: contract}, PieceFilterer: PieceFilterer{contract: contract}}, nil
}

// Piece is an auto generated Go binding around an Ethereum contract.
type Piece struct {
	PieceCaller     // Read-only binding to the contract
	PieceTransactor // Write-only binding to the contract
	PieceFilterer   // Log filterer for contract events
}

// PieceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PieceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PieceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PieceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PieceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PieceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PieceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PieceSession struct {
	Contract     *Piece            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PieceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PieceCallerSession struct {
	Contract *PieceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PieceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PieceTransactorSession struct {
	Contract     *PieceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PieceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PieceRaw struct {
	Contract *Piece // Generic contract binding to access the raw methods on
}

// PieceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PieceCallerRaw struct {
	Contract *PieceCaller // Generic read-only contract binding to access the raw methods on
}

// PieceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PieceTransactorRaw struct {
	Contract *PieceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPiece creates a new instance of Piece, bound to a specific deployed contract.
func NewPiece(address common.Address, backend bind.ContractBackend) (*Piece, error) {
	contract, err := bindPiece(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Piece{PieceCaller: PieceCaller{contract: contract}, PieceTransactor: PieceTransactor{contract: contract}, PieceFilterer: PieceFilterer{contract: contract}}, nil
}

// NewPieceCaller creates a new read-only instance of Piece, bound to a specific deployed contract.
func NewPieceCaller(address common.Address, caller bind.ContractCaller) (*PieceCaller, error) {
	contract, err := bindPiece(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PieceCaller{contract: contract}, nil
}

// NewPieceTransactor creates a new write-only instance of Piece, bound to a specific deployed contract.
func NewPieceTransactor(address common.Address, transactor bind.ContractTransactor) (*PieceTransactor, error) {
	contract, err := bindPiece(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PieceTransactor{contract: contract}, nil
}

// NewPieceFilterer creates a new log filterer instance of Piece, bound to a specific deployed contract.
func NewPieceFilterer(address common.Address, filterer bind.ContractFilterer) (*PieceFilterer, error) {
	contract, err := bindPiece(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PieceFilterer{contract: contract}, nil
}

// bindPiece binds a generic wrapper to an already deployed contract.
func bindPiece(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PieceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Piece *PieceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Piece.Contract.PieceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Piece *PieceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piece.Contract.PieceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Piece *PieceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Piece.Contract.PieceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Piece *PieceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Piece.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Piece *PieceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piece.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Piece *PieceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Piece.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Piece *PieceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Piece *PieceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Piece.Contract.DEFAULTADMINROLE(&_Piece.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Piece *PieceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Piece.Contract.DEFAULTADMINROLE(&_Piece.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Piece *PieceCaller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Piece *PieceSession) GOVERNORROLE() ([32]byte, error) {
	return _Piece.Contract.GOVERNORROLE(&_Piece.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Piece *PieceCallerSession) GOVERNORROLE() ([32]byte, error) {
	return _Piece.Contract.GOVERNORROLE(&_Piece.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Piece *PieceCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Piece *PieceSession) RELAYERROLE() ([32]byte, error) {
	return _Piece.Contract.RELAYERROLE(&_Piece.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Piece *PieceCallerSession) RELAYERROLE() ([32]byte, error) {
	return _Piece.Contract.RELAYERROLE(&_Piece.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Piece *PieceCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Piece *PieceSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Piece.Contract.UPGRADEINTERFACEVERSION(&_Piece.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Piece *PieceCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Piece.Contract.UPGRADEINTERFACEVERSION(&_Piece.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Piece *PieceCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Piece *PieceSession) VERSION() (string, error) {
	return _Piece.Contract.VERSION(&_Piece.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Piece *PieceCallerSession) VERSION() (string, error) {
	return _Piece.Contract.VERSION(&_Piece.CallOpts)
}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_Piece *PieceCaller) CheckSReplica(opts *bind.CallOpts, _a common.Address, _ri uint64, _com []byte) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "checkSReplica", _a, _ri, _com)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_Piece *PieceSession) CheckSReplica(_a common.Address, _ri uint64, _com []byte) (uint64, error) {
	return _Piece.Contract.CheckSReplica(&_Piece.CallOpts, _a, _ri, _com)
}

// CheckSReplica is a free data retrieval call binding the contract method 0x894d35f8.
//
// Solidity: function checkSReplica(address _a, uint64 _ri, bytes _com) view returns(uint64)
func (_Piece *PieceCallerSession) CheckSReplica(_a common.Address, _ri uint64, _com []byte) (uint64, error) {
	return _Piece.Contract.CheckSReplica(&_Piece.CallOpts, _a, _ri, _com)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Piece *PieceCaller) Delay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Piece *PieceSession) Delay() (uint64, error) {
	return _Piece.Contract.Delay(&_Piece.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Piece *PieceCallerSession) Delay() (uint64, error) {
	return _Piece.Contract.Delay(&_Piece.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Piece *PieceCaller) Epoch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Piece *PieceSession) Epoch() (common.Address, error) {
	return _Piece.Contract.Epoch(&_Piece.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Piece *PieceCallerSession) Epoch() (common.Address, error) {
	return _Piece.Contract.Epoch(&_Piece.CallOpts)
}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_Piece *PieceCaller) GetPIndex(opts *bind.CallOpts, _pn []byte) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPIndex", _pn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_Piece *PieceSession) GetPIndex(_pn []byte) (uint64, error) {
	return _Piece.Contract.GetPIndex(&_Piece.CallOpts, _pn)
}

// GetPIndex is a free data retrieval call binding the contract method 0x1ce85e7c.
//
// Solidity: function getPIndex(bytes _pn) view returns(uint64)
func (_Piece *PieceCallerSession) GetPIndex(_pn []byte) (uint64, error) {
	return _Piece.Contract.GetPIndex(&_Piece.CallOpts, _pn)
}

// GetPRI is a free data retrieval call binding the contract method 0xee91365b.
//
// Solidity: function getPRI(uint64 _pi, uint8 _pri) view returns(uint64, address)
func (_Piece *PieceCaller) GetPRI(opts *bind.CallOpts, _pi uint64, _pri uint8) (uint64, common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPRI", _pi, _pri)

	if err != nil {
		return *new(uint64), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetPRI is a free data retrieval call binding the contract method 0xee91365b.
//
// Solidity: function getPRI(uint64 _pi, uint8 _pri) view returns(uint64, address)
func (_Piece *PieceSession) GetPRI(_pi uint64, _pri uint8) (uint64, common.Address, error) {
	return _Piece.Contract.GetPRI(&_Piece.CallOpts, _pi, _pri)
}

// GetPRI is a free data retrieval call binding the contract method 0xee91365b.
//
// Solidity: function getPRI(uint64 _pi, uint8 _pri) view returns(uint64, address)
func (_Piece *PieceCallerSession) GetPRI(_pi uint64, _pri uint8) (uint64, common.Address, error) {
	return _Piece.Contract.GetPRI(&_Piece.CallOpts, _pi, _pri)
}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, address, bytes32)
func (_Piece *PieceCaller) GetPRInfo(opts *bind.CallOpts, _pi uint64, _ri uint64) (uint8, uint8, uint64, common.Address, [32]byte, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPRInfo", _pi, _ri)

	if err != nil {
		return *new(uint8), *new(uint8), *new(uint64), *new(common.Address), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return out0, out1, out2, out3, out4, err

}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, address, bytes32)
func (_Piece *PieceSession) GetPRInfo(_pi uint64, _ri uint64) (uint8, uint8, uint64, common.Address, [32]byte, error) {
	return _Piece.Contract.GetPRInfo(&_Piece.CallOpts, _pi, _ri)
}

// GetPRInfo is a free data retrieval call binding the contract method 0xb650b504.
//
// Solidity: function getPRInfo(uint64 _pi, uint64 _ri) view returns(uint8, uint8, uint64, address, bytes32)
func (_Piece *PieceCallerSession) GetPRInfo(_pi uint64, _ri uint64) (uint8, uint8, uint64, common.Address, [32]byte, error) {
	return _Piece.Contract.GetPRInfo(&_Piece.CallOpts, _pi, _ri)
}

// GetPiece is a free data retrieval call binding the contract method 0x7a87f0c4.
//
// Solidity: function getPiece(uint64 _pi) view returns((uint8,uint8,uint64,uint64,uint64,address,address,uint256,uint256) _pinfo)
func (_Piece *PieceCaller) GetPiece(opts *bind.CallOpts, _pi uint64) (IPiecePieceInfo, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getPiece", _pi)

	if err != nil {
		return *new(IPiecePieceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPiecePieceInfo)).(*IPiecePieceInfo)

	return out0, err

}

// GetPiece is a free data retrieval call binding the contract method 0x7a87f0c4.
//
// Solidity: function getPiece(uint64 _pi) view returns((uint8,uint8,uint64,uint64,uint64,address,address,uint256,uint256) _pinfo)
func (_Piece *PieceSession) GetPiece(_pi uint64) (IPiecePieceInfo, error) {
	return _Piece.Contract.GetPiece(&_Piece.CallOpts, _pi)
}

// GetPiece is a free data retrieval call binding the contract method 0x7a87f0c4.
//
// Solidity: function getPiece(uint64 _pi) view returns((uint8,uint8,uint64,uint64,uint64,address,address,uint256,uint256) _pinfo)
func (_Piece *PieceCallerSession) GetPiece(_pi uint64) (IPiecePieceInfo, error) {
	return _Piece.Contract.GetPiece(&_Piece.CallOpts, _pi)
}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_Piece *PieceCaller) GetRIndex(opts *bind.CallOpts, _rn []byte) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getRIndex", _rn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_Piece *PieceSession) GetRIndex(_rn []byte) (uint64, error) {
	return _Piece.Contract.GetRIndex(&_Piece.CallOpts, _rn)
}

// GetRIndex is a free data retrieval call binding the contract method 0xba8b2618.
//
// Solidity: function getRIndex(bytes _rn) view returns(uint64)
func (_Piece *PieceCallerSession) GetRIndex(_rn []byte) (uint64, error) {
	return _Piece.Contract.GetRIndex(&_Piece.CallOpts, _rn)
}

// GetRS is a free data retrieval call binding the contract method 0x6aadfac7.
//
// Solidity: function getRS(uint64 _pi) view returns(uint8, uint8, uint64)
func (_Piece *PieceCaller) GetRS(opts *bind.CallOpts, _pi uint64) (uint8, uint8, uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getRS", _pi)

	if err != nil {
		return *new(uint8), *new(uint8), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// GetRS is a free data retrieval call binding the contract method 0x6aadfac7.
//
// Solidity: function getRS(uint64 _pi) view returns(uint8, uint8, uint64)
func (_Piece *PieceSession) GetRS(_pi uint64) (uint8, uint8, uint64, error) {
	return _Piece.Contract.GetRS(&_Piece.CallOpts, _pi)
}

// GetRS is a free data retrieval call binding the contract method 0x6aadfac7.
//
// Solidity: function getRS(uint64 _pi) view returns(uint8, uint8, uint64)
func (_Piece *PieceCallerSession) GetRS(_pi uint64) (uint8, uint8, uint64, error) {
	return _Piece.Contract.GetRS(&_Piece.CallOpts, _pi)
}

// GetReplica is a free data retrieval call binding the contract method 0x9108544c.
//
// Solidity: function getReplica(uint64 _r) view returns((uint64,address,bytes32) _ri)
func (_Piece *PieceCaller) GetReplica(opts *bind.CallOpts, _r uint64) (IPieceReplicaInfo, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getReplica", _r)

	if err != nil {
		return *new(IPieceReplicaInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPieceReplicaInfo)).(*IPieceReplicaInfo)

	return out0, err

}

// GetReplica is a free data retrieval call binding the contract method 0x9108544c.
//
// Solidity: function getReplica(uint64 _r) view returns((uint64,address,bytes32) _ri)
func (_Piece *PieceSession) GetReplica(_r uint64) (IPieceReplicaInfo, error) {
	return _Piece.Contract.GetReplica(&_Piece.CallOpts, _r)
}

// GetReplica is a free data retrieval call binding the contract method 0x9108544c.
//
// Solidity: function getReplica(uint64 _r) view returns((uint64,address,bytes32) _ri)
func (_Piece *PieceCallerSession) GetReplica(_r uint64) (IPieceReplicaInfo, error) {
	return _Piece.Contract.GetReplica(&_Piece.CallOpts, _r)
}

// GetRevenue is a free data retrieval call binding the contract method 0x85b67095.
//
// Solidity: function getRevenue(address _a) view returns(uint256)
func (_Piece *PieceCaller) GetRevenue(opts *bind.CallOpts, _a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getRevenue", _a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevenue is a free data retrieval call binding the contract method 0x85b67095.
//
// Solidity: function getRevenue(address _a) view returns(uint256)
func (_Piece *PieceSession) GetRevenue(_a common.Address) (*big.Int, error) {
	return _Piece.Contract.GetRevenue(&_Piece.CallOpts, _a)
}

// GetRevenue is a free data retrieval call binding the contract method 0x85b67095.
//
// Solidity: function getRevenue(address _a) view returns(uint256)
func (_Piece *PieceCallerSession) GetRevenue(_a common.Address) (*big.Int, error) {
	return _Piece.Contract.GetRevenue(&_Piece.CallOpts, _a)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Piece *PieceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Piece *PieceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Piece.Contract.GetRoleAdmin(&_Piece.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Piece *PieceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Piece.Contract.GetRoleAdmin(&_Piece.CallOpts, role)
}

// GetSRAt is a free data retrieval call binding the contract method 0xbcbfd56b.
//
// Solidity: function getSRAt(address _a, uint64 _ri) view returns(uint64)
func (_Piece *PieceCaller) GetSRAt(opts *bind.CallOpts, _a common.Address, _ri uint64) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getSRAt", _a, _ri)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetSRAt is a free data retrieval call binding the contract method 0xbcbfd56b.
//
// Solidity: function getSRAt(address _a, uint64 _ri) view returns(uint64)
func (_Piece *PieceSession) GetSRAt(_a common.Address, _ri uint64) (uint64, error) {
	return _Piece.Contract.GetSRAt(&_Piece.CallOpts, _a, _ri)
}

// GetSRAt is a free data retrieval call binding the contract method 0xbcbfd56b.
//
// Solidity: function getSRAt(address _a, uint64 _ri) view returns(uint64)
func (_Piece *PieceCallerSession) GetSRAt(_a common.Address, _ri uint64) (uint64, error) {
	return _Piece.Contract.GetSRAt(&_Piece.CallOpts, _a, _ri)
}

// GetSStat is a free data retrieval call binding the contract method 0xc3aa3804.
//
// Solidity: function getSStat(address _a, uint64 _e) view returns((uint64,uint256) _si)
func (_Piece *PieceCaller) GetSStat(opts *bind.CallOpts, _a common.Address, _e uint64) (IPieceStoreStat, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getSStat", _a, _e)

	if err != nil {
		return *new(IPieceStoreStat), err
	}

	out0 := *abi.ConvertType(out[0], new(IPieceStoreStat)).(*IPieceStoreStat)

	return out0, err

}

// GetSStat is a free data retrieval call binding the contract method 0xc3aa3804.
//
// Solidity: function getSStat(address _a, uint64 _e) view returns((uint64,uint256) _si)
func (_Piece *PieceSession) GetSStat(_a common.Address, _e uint64) (IPieceStoreStat, error) {
	return _Piece.Contract.GetSStat(&_Piece.CallOpts, _a, _e)
}

// GetSStat is a free data retrieval call binding the contract method 0xc3aa3804.
//
// Solidity: function getSStat(address _a, uint64 _e) view returns((uint64,uint256) _si)
func (_Piece *PieceCallerSession) GetSStat(_a common.Address, _e uint64) (IPieceStoreStat, error) {
	return _Piece.Contract.GetSStat(&_Piece.CallOpts, _a, _e)
}

// GetStore is a free data retrieval call binding the contract method 0x4b4ffccd.
//
// Solidity: function getStore(address _a) view returns((uint64,uint64,uint256,uint256) _si)
func (_Piece *PieceCaller) GetStore(opts *bind.CallOpts, _a common.Address) (IPieceStoreInfo, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getStore", _a)

	if err != nil {
		return *new(IPieceStoreInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPieceStoreInfo)).(*IPieceStoreInfo)

	return out0, err

}

// GetStore is a free data retrieval call binding the contract method 0x4b4ffccd.
//
// Solidity: function getStore(address _a) view returns((uint64,uint64,uint256,uint256) _si)
func (_Piece *PieceSession) GetStore(_a common.Address) (IPieceStoreInfo, error) {
	return _Piece.Contract.GetStore(&_Piece.CallOpts, _a)
}

// GetStore is a free data retrieval call binding the contract method 0x4b4ffccd.
//
// Solidity: function getStore(address _a) view returns((uint64,uint64,uint256,uint256) _si)
func (_Piece *PieceCallerSession) GetStore(_a common.Address) (IPieceStoreInfo, error) {
	return _Piece.Contract.GetStore(&_Piece.CallOpts, _a)
}

// GetStoreCount is a free data retrieval call binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) view returns(uint64)
func (_Piece *PieceCaller) GetStoreCount(opts *bind.CallOpts, _a common.Address, _e uint64) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getStoreCount", _a, _e)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetStoreCount is a free data retrieval call binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) view returns(uint64)
func (_Piece *PieceSession) GetStoreCount(_a common.Address, _e uint64) (uint64, error) {
	return _Piece.Contract.GetStoreCount(&_Piece.CallOpts, _a, _e)
}

// GetStoreCount is a free data retrieval call binding the contract method 0x50979d7e.
//
// Solidity: function getStoreCount(address _a, uint64 _e) view returns(uint64)
func (_Piece *PieceCallerSession) GetStoreCount(_a common.Address, _e uint64) (uint64, error) {
	return _Piece.Contract.GetStoreCount(&_Piece.CallOpts, _a, _e)
}

// GetStoreSalary is a free data retrieval call binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) view returns(uint256)
func (_Piece *PieceCaller) GetStoreSalary(opts *bind.CallOpts, _a common.Address, _e uint64) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "getStoreSalary", _a, _e)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStoreSalary is a free data retrieval call binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) view returns(uint256)
func (_Piece *PieceSession) GetStoreSalary(_a common.Address, _e uint64) (*big.Int, error) {
	return _Piece.Contract.GetStoreSalary(&_Piece.CallOpts, _a, _e)
}

// GetStoreSalary is a free data retrieval call binding the contract method 0x18198ad7.
//
// Solidity: function getStoreSalary(address _a, uint64 _e) view returns(uint256)
func (_Piece *PieceCallerSession) GetStoreSalary(_a common.Address, _e uint64) (*big.Int, error) {
	return _Piece.Contract.GetStoreSalary(&_Piece.CallOpts, _a, _e)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Piece *PieceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Piece *PieceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Piece.Contract.HasRole(&_Piece.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Piece *PieceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Piece.Contract.HasRole(&_Piece.CallOpts, role, account)
}

// MaxSize is a free data retrieval call binding the contract method 0x2565b159.
//
// Solidity: function maxSize() view returns(uint64)
func (_Piece *PieceCaller) MaxSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "maxSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxSize is a free data retrieval call binding the contract method 0x2565b159.
//
// Solidity: function maxSize() view returns(uint64)
func (_Piece *PieceSession) MaxSize() (uint64, error) {
	return _Piece.Contract.MaxSize(&_Piece.CallOpts)
}

// MaxSize is a free data retrieval call binding the contract method 0x2565b159.
//
// Solidity: function maxSize() view returns(uint64)
func (_Piece *PieceCallerSession) MaxSize() (uint64, error) {
	return _Piece.Contract.MaxSize(&_Piece.CallOpts)
}

// MaxStore is a free data retrieval call binding the contract method 0xfca5be7a.
//
// Solidity: function maxStore() view returns(uint64)
func (_Piece *PieceCaller) MaxStore(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "maxStore")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxStore is a free data retrieval call binding the contract method 0xfca5be7a.
//
// Solidity: function maxStore() view returns(uint64)
func (_Piece *PieceSession) MaxStore() (uint64, error) {
	return _Piece.Contract.MaxStore(&_Piece.CallOpts)
}

// MaxStore is a free data retrieval call binding the contract method 0xfca5be7a.
//
// Solidity: function maxStore() view returns(uint64)
func (_Piece *PieceCallerSession) MaxStore() (uint64, error) {
	return _Piece.Contract.MaxStore(&_Piece.CallOpts)
}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Piece *PieceCaller) MinPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "minPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Piece *PieceSession) MinPrice() (*big.Int, error) {
	return _Piece.Contract.MinPrice(&_Piece.CallOpts)
}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Piece *PieceCallerSession) MinPrice() (*big.Int, error) {
	return _Piece.Contract.MinPrice(&_Piece.CallOpts)
}

// MinStore is a free data retrieval call binding the contract method 0x88795125.
//
// Solidity: function minStore() view returns(uint64)
func (_Piece *PieceCaller) MinStore(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "minStore")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinStore is a free data retrieval call binding the contract method 0x88795125.
//
// Solidity: function minStore() view returns(uint64)
func (_Piece *PieceSession) MinStore() (uint64, error) {
	return _Piece.Contract.MinStore(&_Piece.CallOpts)
}

// MinStore is a free data retrieval call binding the contract method 0x88795125.
//
// Solidity: function minStore() view returns(uint64)
func (_Piece *PieceCallerSession) MinStore() (uint64, error) {
	return _Piece.Contract.MinStore(&_Piece.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_Piece *PieceCaller) Node(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "node")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_Piece *PieceSession) Node() (common.Address, error) {
	return _Piece.Contract.Node(&_Piece.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xd70754ec.
//
// Solidity: function node() view returns(address)
func (_Piece *PieceCallerSession) Node() (common.Address, error) {
	return _Piece.Contract.Node(&_Piece.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Piece *PieceCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Piece *PieceSession) ProxiableUUID() ([32]byte, error) {
	return _Piece.Contract.ProxiableUUID(&_Piece.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Piece *PieceCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Piece.Contract.ProxiableUUID(&_Piece.CallOpts)
}

// ReplicaCount is a free data retrieval call binding the contract method 0xa1e030d9.
//
// Solidity: function replicaCount() view returns(uint64)
func (_Piece *PieceCaller) ReplicaCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "replicaCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ReplicaCount is a free data retrieval call binding the contract method 0xa1e030d9.
//
// Solidity: function replicaCount() view returns(uint64)
func (_Piece *PieceSession) ReplicaCount() (uint64, error) {
	return _Piece.Contract.ReplicaCount(&_Piece.CallOpts)
}

// ReplicaCount is a free data retrieval call binding the contract method 0xa1e030d9.
//
// Solidity: function replicaCount() view returns(uint64)
func (_Piece *PieceCallerSession) ReplicaCount() (uint64, error) {
	return _Piece.Contract.ReplicaCount(&_Piece.CallOpts)
}

// ReserveOf is a free data retrieval call binding the contract method 0x898e67f1.
//
// Solidity: function reserveOf(uint64 _pi) view returns(uint256)
func (_Piece *PieceCaller) ReserveOf(opts *bind.CallOpts, _pi uint64) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "reserveOf", _pi)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveOf is a free data retrieval call binding the contract method 0x898e67f1.
//
// Solidity: function reserveOf(uint64 _pi) view returns(uint256)
func (_Piece *PieceSession) ReserveOf(_pi uint64) (*big.Int, error) {
	return _Piece.Contract.ReserveOf(&_Piece.CallOpts, _pi)
}

// ReserveOf is a free data retrieval call binding the contract method 0x898e67f1.
//
// Solidity: function reserveOf(uint64 _pi) view returns(uint256)
func (_Piece *PieceCallerSession) ReserveOf(_pi uint64) (*big.Int, error) {
	return _Piece.Contract.ReserveOf(&_Piece.CallOpts, _pi)
}

// StreamPrice is a free data retrieval call binding the contract method 0x75011aba.
//
// Solidity: function streamPrice() view returns(uint256)
func (_Piece *PieceCaller) StreamPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "streamPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StreamPrice is a free data retrieval call binding the contract method 0x75011aba.
//
// Solidity: function streamPrice() view returns(uint256)
func (_Piece *PieceSession) StreamPrice() (*big.Int, error) {
	return _Piece.Contract.StreamPrice(&_Piece.CallOpts)
}

// StreamPrice is a free data retrieval call binding the contract method 0x75011aba.
//
// Solidity: function streamPrice() view returns(uint256)
func (_Piece *PieceCallerSession) StreamPrice() (*big.Int, error) {
	return _Piece.Contract.StreamPrice(&_Piece.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Piece *PieceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Piece *PieceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Piece.Contract.SupportsInterface(&_Piece.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Piece *PieceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Piece.Contract.SupportsInterface(&_Piece.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Piece *PieceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Piece *PieceSession) Token() (common.Address, error) {
	return _Piece.Contract.Token(&_Piece.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Piece *PieceCallerSession) Token() (common.Address, error) {
	return _Piece.Contract.Token(&_Piece.CallOpts)
}

// AddPiece is a paid mutator transaction binding the contract method 0x8ef0e895.
//
// Solidity: function addPiece(bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk, address _s) returns()
func (_Piece *PieceTransactor) AddPiece(opts *bind.TransactOpts, _pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8, _s common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "addPiece", _pn, _price, _size, _expire, rsn, rsk, _s)
}

// AddPiece is a paid mutator transaction binding the contract method 0x8ef0e895.
//
// Solidity: function addPiece(bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk, address _s) returns()
func (_Piece *PieceSession) AddPiece(_pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8, _s common.Address) (*types.Transaction, error) {
	return _Piece.Contract.AddPiece(&_Piece.TransactOpts, _pn, _price, _size, _expire, rsn, rsk, _s)
}

// AddPiece is a paid mutator transaction binding the contract method 0x8ef0e895.
//
// Solidity: function addPiece(bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk, address _s) returns()
func (_Piece *PieceTransactorSession) AddPiece(_pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8, _s common.Address) (*types.Transaction, error) {
	return _Piece.Contract.AddPiece(&_Piece.TransactOpts, _pn, _price, _size, _expire, rsn, rsk, _s)
}

// AddPieceFor is a paid mutator transaction binding the contract method 0x2620904f.
//
// Solidity: function addPieceFor(address owner, bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk, address _s) returns()
func (_Piece *PieceTransactor) AddPieceFor(opts *bind.TransactOpts, owner common.Address, _pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8, _s common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "addPieceFor", owner, _pn, _price, _size, _expire, rsn, rsk, _s)
}

// AddPieceFor is a paid mutator transaction binding the contract method 0x2620904f.
//
// Solidity: function addPieceFor(address owner, bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk, address _s) returns()
func (_Piece *PieceSession) AddPieceFor(owner common.Address, _pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8, _s common.Address) (*types.Transaction, error) {
	return _Piece.Contract.AddPieceFor(&_Piece.TransactOpts, owner, _pn, _price, _size, _expire, rsn, rsk, _s)
}

// AddPieceFor is a paid mutator transaction binding the contract method 0x2620904f.
//
// Solidity: function addPieceFor(address owner, bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 rsn, uint8 rsk, address _s) returns()
func (_Piece *PieceTransactorSession) AddPieceFor(owner common.Address, _pn []byte, _price *big.Int, _size uint64, _expire uint64, rsn uint8, rsk uint8, _s common.Address) (*types.Transaction, error) {
	return _Piece.Contract.AddPieceFor(&_Piece.TransactOpts, owner, _pn, _price, _size, _expire, rsn, rsk, _s)
}

// AddReplica is a paid mutator transaction binding the contract method 0x194cbaf5.
//
// Solidity: function addReplica(bytes _rn, uint64 _pi, uint8 _pri, bytes _proof) returns()
func (_Piece *PieceTransactor) AddReplica(opts *bind.TransactOpts, _rn []byte, _pi uint64, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "addReplica", _rn, _pi, _pri, _proof)
}

// AddReplica is a paid mutator transaction binding the contract method 0x194cbaf5.
//
// Solidity: function addReplica(bytes _rn, uint64 _pi, uint8 _pri, bytes _proof) returns()
func (_Piece *PieceSession) AddReplica(_rn []byte, _pi uint64, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _Piece.Contract.AddReplica(&_Piece.TransactOpts, _rn, _pi, _pri, _proof)
}

// AddReplica is a paid mutator transaction binding the contract method 0x194cbaf5.
//
// Solidity: function addReplica(bytes _rn, uint64 _pi, uint8 _pri, bytes _proof) returns()
func (_Piece *PieceTransactorSession) AddReplica(_rn []byte, _pi uint64, _pri uint8, _proof []byte) (*types.Transaction, error) {
	return _Piece.Contract.AddReplica(&_Piece.TransactOpts, _rn, _pi, _pri, _proof)
}

// CheckStore is a paid mutator transaction binding the contract method 0x62b98032.
//
// Solidity: function checkStore(address _a) returns()
func (_Piece *PieceTransactor) CheckStore(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "checkStore", _a)
}

// CheckStore is a paid mutator transaction binding the contract method 0x62b98032.
//
// Solidity: function checkStore(address _a) returns()
func (_Piece *PieceSession) CheckStore(_a common.Address) (*types.Transaction, error) {
	return _Piece.Contract.CheckStore(&_Piece.TransactOpts, _a)
}

// CheckStore is a paid mutator transaction binding the contract method 0x62b98032.
//
// Solidity: function checkStore(address _a) returns()
func (_Piece *PieceTransactorSession) CheckStore(_a common.Address) (*types.Transaction, error) {
	return _Piece.Contract.CheckStore(&_Piece.TransactOpts, _a)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Piece *PieceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Piece *PieceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Piece.Contract.GrantRole(&_Piece.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Piece *PieceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Piece.Contract.GrantRole(&_Piece.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x7cf1ee91.
//
// Solidity: function initialize(address _token, address _epoch, address _node, (uint64,uint64,uint64,uint64,uint256,uint256) params, address initialOwner) returns()
func (_Piece *PieceTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _epoch common.Address, _node common.Address, params IPieceInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "initialize", _token, _epoch, _node, params, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x7cf1ee91.
//
// Solidity: function initialize(address _token, address _epoch, address _node, (uint64,uint64,uint64,uint64,uint256,uint256) params, address initialOwner) returns()
func (_Piece *PieceSession) Initialize(_token common.Address, _epoch common.Address, _node common.Address, params IPieceInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _Piece.Contract.Initialize(&_Piece.TransactOpts, _token, _epoch, _node, params, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x7cf1ee91.
//
// Solidity: function initialize(address _token, address _epoch, address _node, (uint64,uint64,uint64,uint64,uint256,uint256) params, address initialOwner) returns()
func (_Piece *PieceTransactorSession) Initialize(_token common.Address, _epoch common.Address, _node common.Address, params IPieceInitParams, initialOwner common.Address) (*types.Transaction, error) {
	return _Piece.Contract.Initialize(&_Piece.TransactOpts, _token, _epoch, _node, params, initialOwner)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Piece *PieceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Piece *PieceSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Piece.Contract.RenounceRole(&_Piece.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Piece *PieceTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Piece.Contract.RenounceRole(&_Piece.TransactOpts, role, callerConfirmation)
}

// Retake is a paid mutator transaction binding the contract method 0x123f8459.
//
// Solidity: function retake(uint64 _pi) returns()
func (_Piece *PieceTransactor) Retake(opts *bind.TransactOpts, _pi uint64) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "retake", _pi)
}

// Retake is a paid mutator transaction binding the contract method 0x123f8459.
//
// Solidity: function retake(uint64 _pi) returns()
func (_Piece *PieceSession) Retake(_pi uint64) (*types.Transaction, error) {
	return _Piece.Contract.Retake(&_Piece.TransactOpts, _pi)
}

// Retake is a paid mutator transaction binding the contract method 0x123f8459.
//
// Solidity: function retake(uint64 _pi) returns()
func (_Piece *PieceTransactorSession) Retake(_pi uint64) (*types.Transaction, error) {
	return _Piece.Contract.Retake(&_Piece.TransactOpts, _pi)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Piece *PieceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Piece *PieceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Piece.Contract.RevokeRole(&_Piece.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Piece *PieceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Piece.Contract.RevokeRole(&_Piece.TransactOpts, role, account)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _m) returns()
func (_Piece *PieceTransactor) Settle(opts *bind.TransactOpts, _m *big.Int) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "settle", _m)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _m) returns()
func (_Piece *PieceSession) Settle(_m *big.Int) (*types.Transaction, error) {
	return _Piece.Contract.Settle(&_Piece.TransactOpts, _m)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _m) returns()
func (_Piece *PieceTransactorSession) Settle(_m *big.Int) (*types.Transaction, error) {
	return _Piece.Contract.Settle(&_Piece.TransactOpts, _m)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Piece *PieceTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Piece *PieceSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Piece.Contract.UpgradeToAndCall(&_Piece.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Piece *PieceTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Piece.Contract.UpgradeToAndCall(&_Piece.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _m) returns()
func (_Piece *PieceTransactor) Withdraw(opts *bind.TransactOpts, _m *big.Int) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "withdraw", _m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _m) returns()
func (_Piece *PieceSession) Withdraw(_m *big.Int) (*types.Transaction, error) {
	return _Piece.Contract.Withdraw(&_Piece.TransactOpts, _m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _m) returns()
func (_Piece *PieceTransactorSession) Withdraw(_m *big.Int) (*types.Transaction, error) {
	return _Piece.Contract.Withdraw(&_Piece.TransactOpts, _m)
}

// PieceAddPieceIterator is returned from FilterAddPiece and is used to iterate over the raw logs and unpacked data for AddPiece events raised by the Piece contract.
type PieceAddPieceIterator struct {
	Event *PieceAddPiece // Event containing the contract specifics and raw log

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
func (it *PieceAddPieceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceAddPiece)
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
		it.Event = new(PieceAddPiece)
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
func (it *PieceAddPieceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceAddPieceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceAddPiece represents a AddPiece event raised by the Piece contract.
type PieceAddPiece struct {
	A   common.Address
	Pi  uint64
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddPiece is a free log retrieval operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_Piece *PieceFilterer) FilterAddPiece(opts *bind.FilterOpts, _a []common.Address) (*PieceAddPieceIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "AddPiece", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceAddPieceIterator{contract: _Piece.contract, event: "AddPiece", logs: logs, sub: sub}, nil
}

// WatchAddPiece is a free log subscription operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_Piece *PieceFilterer) WatchAddPiece(opts *bind.WatchOpts, sink chan<- *PieceAddPiece, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "AddPiece", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceAddPiece)
				if err := _Piece.contract.UnpackLog(event, "AddPiece", log); err != nil {
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

// ParseAddPiece is a log parse operation binding the contract event 0x09ff0fb5f488978fdffbf3223e63e914f2a9b0c228844b8e2caedad8a85d220c.
//
// Solidity: event AddPiece(address indexed _a, uint64 _pi, uint64 _e)
func (_Piece *PieceFilterer) ParseAddPiece(log types.Log) (*PieceAddPiece, error) {
	event := new(PieceAddPiece)
	if err := _Piece.contract.UnpackLog(event, "AddPiece", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceAddPieceV2Iterator is returned from FilterAddPieceV2 and is used to iterate over the raw logs and unpacked data for AddPieceV2 events raised by the Piece contract.
type PieceAddPieceV2Iterator struct {
	Event *PieceAddPieceV2 // Event containing the contract specifics and raw log

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
func (it *PieceAddPieceV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceAddPieceV2)
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
		it.Event = new(PieceAddPieceV2)
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
func (it *PieceAddPieceV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceAddPieceV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceAddPieceV2 represents a AddPieceV2 event raised by the Piece contract.
type PieceAddPieceV2 struct {
	A      common.Address
	Pi     uint64
	E      uint64
	Pn     []byte
	Price  *big.Int
	Size   uint64
	Expire uint64
	Rsn    uint8
	Rsk    uint8
	S      common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddPieceV2 is a free log retrieval operation binding the contract event 0xe37588ecbc243b6bdc3191c8d9743ce41ab2f677eea6aafbd2ce4cd59480bbdd.
//
// Solidity: event AddPieceV2(address indexed _a, uint64 _pi, uint64 _e, bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 _rsn, uint8 _rsk, address _s)
func (_Piece *PieceFilterer) FilterAddPieceV2(opts *bind.FilterOpts, _a []common.Address) (*PieceAddPieceV2Iterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "AddPieceV2", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceAddPieceV2Iterator{contract: _Piece.contract, event: "AddPieceV2", logs: logs, sub: sub}, nil
}

// WatchAddPieceV2 is a free log subscription operation binding the contract event 0xe37588ecbc243b6bdc3191c8d9743ce41ab2f677eea6aafbd2ce4cd59480bbdd.
//
// Solidity: event AddPieceV2(address indexed _a, uint64 _pi, uint64 _e, bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 _rsn, uint8 _rsk, address _s)
func (_Piece *PieceFilterer) WatchAddPieceV2(opts *bind.WatchOpts, sink chan<- *PieceAddPieceV2, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "AddPieceV2", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceAddPieceV2)
				if err := _Piece.contract.UnpackLog(event, "AddPieceV2", log); err != nil {
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

// ParseAddPieceV2 is a log parse operation binding the contract event 0xe37588ecbc243b6bdc3191c8d9743ce41ab2f677eea6aafbd2ce4cd59480bbdd.
//
// Solidity: event AddPieceV2(address indexed _a, uint64 _pi, uint64 _e, bytes _pn, uint256 _price, uint64 _size, uint64 _expire, uint8 _rsn, uint8 _rsk, address _s)
func (_Piece *PieceFilterer) ParseAddPieceV2(log types.Log) (*PieceAddPieceV2, error) {
	event := new(PieceAddPieceV2)
	if err := _Piece.contract.UnpackLog(event, "AddPieceV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceAddReplicaIterator is returned from FilterAddReplica and is used to iterate over the raw logs and unpacked data for AddReplica events raised by the Piece contract.
type PieceAddReplicaIterator struct {
	Event *PieceAddReplica // Event containing the contract specifics and raw log

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
func (it *PieceAddReplicaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceAddReplica)
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
		it.Event = new(PieceAddReplica)
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
func (it *PieceAddReplicaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceAddReplicaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceAddReplica represents a AddReplica event raised by the Piece contract.
type PieceAddReplica struct {
	A   common.Address
	Ri  uint64
	Sri uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddReplica is a free log retrieval operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_Piece *PieceFilterer) FilterAddReplica(opts *bind.FilterOpts, _a []common.Address) (*PieceAddReplicaIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "AddReplica", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceAddReplicaIterator{contract: _Piece.contract, event: "AddReplica", logs: logs, sub: sub}, nil
}

// WatchAddReplica is a free log subscription operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_Piece *PieceFilterer) WatchAddReplica(opts *bind.WatchOpts, sink chan<- *PieceAddReplica, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "AddReplica", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceAddReplica)
				if err := _Piece.contract.UnpackLog(event, "AddReplica", log); err != nil {
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

// ParseAddReplica is a log parse operation binding the contract event 0x4b079e585085a3c94b2865036e598391fad4359764b36da39e21897be56349ef.
//
// Solidity: event AddReplica(address indexed _a, uint64 _ri, uint64 _sri)
func (_Piece *PieceFilterer) ParseAddReplica(log types.Log) (*PieceAddReplica, error) {
	event := new(PieceAddReplica)
	if err := _Piece.contract.UnpackLog(event, "AddReplica", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceAddReplicaV2Iterator is returned from FilterAddReplicaV2 and is used to iterate over the raw logs and unpacked data for AddReplicaV2 events raised by the Piece contract.
type PieceAddReplicaV2Iterator struct {
	Event *PieceAddReplicaV2 // Event containing the contract specifics and raw log

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
func (it *PieceAddReplicaV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceAddReplicaV2)
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
		it.Event = new(PieceAddReplicaV2)
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
func (it *PieceAddReplicaV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceAddReplicaV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceAddReplicaV2 represents a AddReplicaV2 event raised by the Piece contract.
type PieceAddReplicaV2 struct {
	A     common.Address
	Ri    uint64
	Sri   uint64
	Rn    []byte
	Pi    uint64
	Pri   uint8
	Proof []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddReplicaV2 is a free log retrieval operation binding the contract event 0xf226a06b14819931288115f369b127b1a9986ea785f3a4cfcf16b0b445cf1ca5.
//
// Solidity: event AddReplicaV2(address indexed _a, uint64 _ri, uint64 _sri, bytes _rn, uint64 _pi, uint8 _pri, bytes _proof)
func (_Piece *PieceFilterer) FilterAddReplicaV2(opts *bind.FilterOpts, _a []common.Address) (*PieceAddReplicaV2Iterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "AddReplicaV2", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceAddReplicaV2Iterator{contract: _Piece.contract, event: "AddReplicaV2", logs: logs, sub: sub}, nil
}

// WatchAddReplicaV2 is a free log subscription operation binding the contract event 0xf226a06b14819931288115f369b127b1a9986ea785f3a4cfcf16b0b445cf1ca5.
//
// Solidity: event AddReplicaV2(address indexed _a, uint64 _ri, uint64 _sri, bytes _rn, uint64 _pi, uint8 _pri, bytes _proof)
func (_Piece *PieceFilterer) WatchAddReplicaV2(opts *bind.WatchOpts, sink chan<- *PieceAddReplicaV2, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "AddReplicaV2", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceAddReplicaV2)
				if err := _Piece.contract.UnpackLog(event, "AddReplicaV2", log); err != nil {
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

// ParseAddReplicaV2 is a log parse operation binding the contract event 0xf226a06b14819931288115f369b127b1a9986ea785f3a4cfcf16b0b445cf1ca5.
//
// Solidity: event AddReplicaV2(address indexed _a, uint64 _ri, uint64 _sri, bytes _rn, uint64 _pi, uint8 _pri, bytes _proof)
func (_Piece *PieceFilterer) ParseAddReplicaV2(log types.Log) (*PieceAddReplicaV2, error) {
	event := new(PieceAddReplicaV2)
	if err := _Piece.contract.UnpackLog(event, "AddReplicaV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Piece contract.
type PieceInitializedIterator struct {
	Event *PieceInitialized // Event containing the contract specifics and raw log

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
func (it *PieceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceInitialized)
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
		it.Event = new(PieceInitialized)
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
func (it *PieceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceInitialized represents a Initialized event raised by the Piece contract.
type PieceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Piece *PieceFilterer) FilterInitialized(opts *bind.FilterOpts) (*PieceInitializedIterator, error) {

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PieceInitializedIterator{contract: _Piece.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Piece *PieceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PieceInitialized) (event.Subscription, error) {

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceInitialized)
				if err := _Piece.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Piece *PieceFilterer) ParseInitialized(log types.Log) (*PieceInitialized, error) {
	event := new(PieceInitialized)
	if err := _Piece.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceRetakeIterator is returned from FilterRetake and is used to iterate over the raw logs and unpacked data for Retake events raised by the Piece contract.
type PieceRetakeIterator struct {
	Event *PieceRetake // Event containing the contract specifics and raw log

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
func (it *PieceRetakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceRetake)
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
		it.Event = new(PieceRetake)
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
func (it *PieceRetakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceRetakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceRetake represents a Retake event raised by the Piece contract.
type PieceRetake struct {
	A   common.Address
	Pi  uint64
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRetake is a free log retrieval operation binding the contract event 0x44647ea87405778ba90a8ca28d2b772b1350bcd5950e5111bbc8c506a5c9632a.
//
// Solidity: event Retake(address indexed _a, uint64 _pi, uint256 _m)
func (_Piece *PieceFilterer) FilterRetake(opts *bind.FilterOpts, _a []common.Address) (*PieceRetakeIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Retake", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceRetakeIterator{contract: _Piece.contract, event: "Retake", logs: logs, sub: sub}, nil
}

// WatchRetake is a free log subscription operation binding the contract event 0x44647ea87405778ba90a8ca28d2b772b1350bcd5950e5111bbc8c506a5c9632a.
//
// Solidity: event Retake(address indexed _a, uint64 _pi, uint256 _m)
func (_Piece *PieceFilterer) WatchRetake(opts *bind.WatchOpts, sink chan<- *PieceRetake, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Retake", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceRetake)
				if err := _Piece.contract.UnpackLog(event, "Retake", log); err != nil {
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

// ParseRetake is a log parse operation binding the contract event 0x44647ea87405778ba90a8ca28d2b772b1350bcd5950e5111bbc8c506a5c9632a.
//
// Solidity: event Retake(address indexed _a, uint64 _pi, uint256 _m)
func (_Piece *PieceFilterer) ParseRetake(log types.Log) (*PieceRetake, error) {
	event := new(PieceRetake)
	if err := _Piece.contract.UnpackLog(event, "Retake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Piece contract.
type PieceRoleAdminChangedIterator struct {
	Event *PieceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PieceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceRoleAdminChanged)
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
		it.Event = new(PieceRoleAdminChanged)
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
func (it *PieceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceRoleAdminChanged represents a RoleAdminChanged event raised by the Piece contract.
type PieceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Piece *PieceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PieceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Piece.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PieceRoleAdminChangedIterator{contract: _Piece.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Piece *PieceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PieceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Piece.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceRoleAdminChanged)
				if err := _Piece.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Piece *PieceFilterer) ParseRoleAdminChanged(log types.Log) (*PieceRoleAdminChanged, error) {
	event := new(PieceRoleAdminChanged)
	if err := _Piece.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Piece contract.
type PieceRoleGrantedIterator struct {
	Event *PieceRoleGranted // Event containing the contract specifics and raw log

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
func (it *PieceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceRoleGranted)
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
		it.Event = new(PieceRoleGranted)
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
func (it *PieceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceRoleGranted represents a RoleGranted event raised by the Piece contract.
type PieceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Piece *PieceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PieceRoleGrantedIterator, error) {

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

	logs, sub, err := _Piece.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PieceRoleGrantedIterator{contract: _Piece.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Piece *PieceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PieceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Piece.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceRoleGranted)
				if err := _Piece.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Piece *PieceFilterer) ParseRoleGranted(log types.Log) (*PieceRoleGranted, error) {
	event := new(PieceRoleGranted)
	if err := _Piece.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Piece contract.
type PieceRoleRevokedIterator struct {
	Event *PieceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PieceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceRoleRevoked)
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
		it.Event = new(PieceRoleRevoked)
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
func (it *PieceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceRoleRevoked represents a RoleRevoked event raised by the Piece contract.
type PieceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Piece *PieceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PieceRoleRevokedIterator, error) {

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

	logs, sub, err := _Piece.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PieceRoleRevokedIterator{contract: _Piece.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Piece *PieceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PieceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Piece.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceRoleRevoked)
				if err := _Piece.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Piece *PieceFilterer) ParseRoleRevoked(log types.Log) (*PieceRoleRevoked, error) {
	event := new(PieceRoleRevoked)
	if err := _Piece.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the Piece contract.
type PieceSettleIterator struct {
	Event *PieceSettle // Event containing the contract specifics and raw log

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
func (it *PieceSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceSettle)
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
		it.Event = new(PieceSettle)
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
func (it *PieceSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceSettle represents a Settle event raised by the Piece contract.
type PieceSettle struct {
	A   common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) FilterSettle(opts *bind.FilterOpts, _a []common.Address) (*PieceSettleIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Settle", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceSettleIterator{contract: _Piece.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *PieceSettle, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Settle", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceSettle)
				if err := _Piece.contract.UnpackLog(event, "Settle", log); err != nil {
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

// ParseSettle is a log parse operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) ParseSettle(log types.Log) (*PieceSettle, error) {
	event := new(PieceSettle)
	if err := _Piece.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Piece contract.
type PieceUpgradedIterator struct {
	Event *PieceUpgraded // Event containing the contract specifics and raw log

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
func (it *PieceUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceUpgraded)
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
		it.Event = new(PieceUpgraded)
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
func (it *PieceUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceUpgraded represents a Upgraded event raised by the Piece contract.
type PieceUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Piece *PieceFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PieceUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PieceUpgradedIterator{contract: _Piece.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Piece *PieceFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PieceUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceUpgraded)
				if err := _Piece.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Piece *PieceFilterer) ParseUpgraded(log types.Log) (*PieceUpgraded, error) {
	event := new(PieceUpgraded)
	if err := _Piece.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PieceWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Piece contract.
type PieceWithdrawIterator struct {
	Event *PieceWithdraw // Event containing the contract specifics and raw log

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
func (it *PieceWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceWithdraw)
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
		it.Event = new(PieceWithdraw)
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
func (it *PieceWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceWithdraw represents a Withdraw event raised by the Piece contract.
type PieceWithdraw struct {
	A   common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) FilterWithdraw(opts *bind.FilterOpts, _a []common.Address) (*PieceWithdrawIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return &PieceWithdrawIterator{contract: _Piece.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PieceWithdraw, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceWithdraw)
				if err := _Piece.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_Piece *PieceFilterer) ParseWithdraw(log types.Log) (*PieceWithdraw, error) {
	event := new(PieceWithdraw)
	if err := _Piece.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
