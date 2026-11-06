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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAYER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addPiece\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_expire\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rsn\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rsk\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_s\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPieceFor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_expire\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rsn\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rsk\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_s\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addReplica\",\"inputs\":[{\"name\":\"_rn\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkSReplica\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_com\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkStore\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEpoch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freeForgedSlot\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPIndex\",\"inputs\":[{\"name\":\"_pn\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPRI\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPRInfo\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPiece\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"_pinfo\",\"type\":\"tuple\",\"internalType\":\"structIPiece.PieceInfo\",\"components\":[{\"name\":\"rsn\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rsk\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"start\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expire\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"streamer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRIndex\",\"inputs\":[{\"name\":\"_rn\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRS\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReplica\",\"inputs\":[{\"name\":\"_r\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"_ri\",\"type\":\"tuple\",\"internalType\":\"structIPiece.ReplicaInfo\",\"components\":[{\"name\":\"piece\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"storedOn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevenue\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSRAt\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSStat\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"_si\",\"type\":\"tuple\",\"internalType\":\"structIPiece.StoreStat\",\"components\":[{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"salary\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStore\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_si\",\"type\":\"tuple\",\"internalType\":\"structIPiece.StoreInfo\",\"components\":[{\"name\":\"epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"salary\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStoreCount\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStoreSalary\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_e\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_epoch\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIPiece.InitParams\",\"components\":[{\"name\":\"delay\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minStore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxStore\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSize\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"streamPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxStore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"node\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"replicaCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reserveOf\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"retake\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settle\",\"inputs\":[{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"streamPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddPiece\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddPieceV2\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_pn\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"_price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_size\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_expire\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_rsn\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_rsk\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_s\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddReplica\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_sri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddReplicaV2\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_sri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_rn\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Retake\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Settle\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotFreed\",\"inputs\":[{\"name\":\"_pi\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"_pri\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_ri\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"_store\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_refund\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a08060405234620000d157306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c16620000c257506001600160401b036002600160401b0319828216016200007c575b604051613de39081620000d682396080518181816113ec01526114ea0152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80806200005c565b63f92ee8a960e01b8152600490fd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a7146126bd57508063123f84591461245857806318198ad7146123b7578063194cbaf514611afc5780631ce85e7c14611aad578063248a9ca314611a745780632565b15914611a4a5780632620904f146119355780632e1a7d4d1461183d5780632f2ff15d146117f357806336568abe146117ab5780634b4ffccd146117015780634f1ef286146114ab57806350979d7e1461145657806352d1902d146113d957806362b980321461130e5780636a42b8f8146112e45780636aadfac71461127e57806375011aba146112605780637a87f0c4146110c85780637cf1ee9114610cf157806385b6709514610cb85780638879512514610c91578063894d35f814610baa578063898e67f114610b7e5780638df82800146108f45780638ef0e8951461086b578063900cf0cf146108425780639108544c146107ae57806391d1485414610758578063926d7d7f1461071d578063a1e030d9146106f6578063a217fddf146106da578063ad3cb1cc14610694578063b650b504146105c1578063ba8b26181461055f578063baa3cb0214610531578063bcbfd56b146104dc578063c3aa380414610443578063ccc5749014610408578063d547741f146103ba578063d70754ec14610391578063e45be8eb14610373578063ee91365b14610308578063fad73d52146102cb578063fc0c546a146102a4578063fca5be7a1461027a5763ffa1ad741461022e575f80fd5b3461027757806003193601126102775761027360405161024d816127f9565b60058152640322e302e360dc1b60208201526040519182916020835260208301906128d9565b0390f35b80fd5b503461027757806003193601126102775760206001600160401b0360035460401c16604051908152f35b5034610277578060031936011261027757546040516001600160a01b039091168152602090f35b50346102775760203660031901126102775760ff60406020926001600160401b036102f4612710565b168152601084522054166040519015158152f35b503461027757604036600319011261027757604090610325612710565b9060ff836103316128a8565b926001600160401b03809516815260086020522091165f52602052815f20541661035a81612a91565b50548251918252821c6001600160a01b03166020820152f35b50346102775780600319360112610277576020600454604051908152f35b50346102775780600319360112610277576002546040516001600160a01b039091168152602090f35b5034610277576040366003190112610277576104046004356103da61277e565b908084525f80516020613d6e8339815191526020526103ff60016040862001546132b0565b613acd565b5080f35b503461027757806003193601126102775760206040517f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f558152f35b503461027757604036600319011261027757604090610460612768565b610468612726565b83805192610475846127f9565b8484528460208095015260018060a01b031693848152600e8452818120926001600160401b0380911693845f528552825f205416948152600e845220905f5281526001835f2001548184516104c9816127f9565b8481520190815283519283525190820152f35b50346102775760403660031901126102775760209060406104fb612768565b91610504612726565b9260018060a01b03168152600d8452206001600160401b038092165f52825260405f205416604051908152f35b50346102775760403660031901126102775761055c61054e612710565b6105566128a8565b90612ebd565b80f35b503461027757602036600319011261027757600435906001600160401b039081831161027757506105ad602061059a8194369060040161287a565b81604051938285809451938492016128b8565b8101600b8152030190205416604051908152f35b50346102775760403660031901126102775760a06105dd612710565b6105e5612726565b906105ef81612969565b509161063261062d61060e6001600160401b0380965460101c16612a22565b61062760ff61061c87612969565b505460081c16612e76565b90612e91565b6129e2565b9260ff61063e84612969565b50541693600161067361066160ff61065588612969565b505460081c1696612969565b50600283808b1b039101541694612a91565b50015493604051958652602086015216604084015260608301526080820152f35b50346102775780600319360112610277576102736040516106b4816127f9565b60058152640352e302e360dc1b60208201526040519182916020835260208301906128d9565b5034610277578060031936011261027757602090604051908152f35b503461027757806003193601126102775760206001600160401b03600a5416604051908152f35b503461027757806003193601126102775760206040517fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc48152f35b503461027757604036600319011261027757604061077461277e565b9160043581525f80516020613d6e833981519152602052209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b50346102775760203660031901126102775760606107ca612710565b6107d2612a73565b506001600160401b036107e482612a91565b505416906107f181612a91565b505460401c6001600160a01b039081169160019061080e90612a91565b500154916040805161081f816127de565b858152602081019283520192835260405193845251166020830152516040820152f35b50346102775780600319360112610277576001546040516001600160a01b039091168152602090f35b50346102775760e0366003190112610277576001600160401b036004358181116108f05761089d90369060040161287a565b9060443590811681036108ec576108b261273c565b60843560ff811681036108ec576108c7612898565b60c435939092906001600160a01b03851685036108ec5761055c9560243590336132de565b5f80fd5b8280fd5b503461027757602080600319360112610b40576002546040516327c830a960e01b81526004803593926001600160a01b0392918491839190829086165afa8015610ae5578590610b44575b61094a91501561290b565b838160015416803b15610b405781809160046040518094819363919840ad60e01b83525af18015610b3557610b1d575b5050600482826001541660405192838092639fa6a6e360e01b82525afa8015610ae5576109af918691610af0575b5033613b4c565b6044604082600254168151928380926361e728b160e01b8252336004830152600160248301525afa908115610ae55785908692610ab3575b5015610a8357338552600c835260016040862001549081600702916007830403610a6f577fa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d2049392610a3e610a639387931015612b68565b338752600c845260026040882001610a57838254612a15565b905533908754166131e2565b6040519283523392a280f35b634e487b7160e01b86526011600452602486fd5b60405162461bcd60e51b8152602060048201526008602482015267696e76206e6f646560c01b6044820152606490fd5b9050610ad7915060403d604011610ade575b610acf8183612814565b810190612b4b565b905f6109e7565b503d610ac5565b6040513d87823e3d90fd5b610b109150843d8611610b16575b610b088183612814565b81019061294a565b5f6109a8565b503d610afe565b610b2690612794565b610b3157835f61097a565b8380fd5b6040513d84823e3d90fd5b5080fd5b508281813d8311610b77575b610b5a8183612814565b81010312610b7357610b6e61094a916128fe565b61093f565b8480fd5b503d610b50565b5034610277576020366003190112610277576020610ba2610b9d612710565b612cdd565b604051908152f35b503461027757606036600319011261027757610bc4612768565b90610bcd612726565b90604435926001600160401b03928385116108f057610bf2604095369060040161287a565b9160018060a01b03168352836020958694600d86522091165f52825282610c2f8160405f20541693849381604051938285809451938492016128b8565b8101600b815203019020541603610c635781610c4d610c5592612a91565b505416612969565b505460901c16604051908152f35b60405162461bcd60e51b815260048101849052600660248201526534b73b1039b960d11b6044820152606490fd5b503461027757806003193601126102775760206001600160401b0360035416604051908152f35b5034610277576020366003190112610277576020906040906001600160a01b03610ce0612768565b168152600f83522054604051908152f35b50346102775761014036600319011261027757610d0c612768565b610d1461277e565b604435906001600160a01b03821682036108ec5760c0366063190112610b315760405160c081018181106001600160401b038211176110b457604052610d5861273c565b8152610d62612752565b602082015260a4356001600160401b03811681036108ec57604082015260c4356001600160401b03811681036108ec57606082015260e43560808201526101043560a0820152610124356001600160a01b03811690036108ec575f80516020613d8e83398151915254926001600160401b038416801590816110a4575b600114908161109a575b159081611091575b5061107f5760016001600160401b03198516175f80516020613d8e8339815191525560ff8460401c1615611053575b60ff5f80516020613d8e8339815191525460401c1615611041576001600160401b038251166001600160401b0360208401511611156110085760ff9460a0936bffffffffffffffffffffffff851b91600180871b031682895416178855600180861b03169060015416176001556002546001600160401b03841b8351851b1691600180861b03169063ffffffff60e01b1617176002556001600160401b036020820151166003549067ffffffffffffffff60401b604084015160401b16906001600160401b0360801b606085015160801b16926001600160401b0360c01b1617171760035560808101516004550151600555610f1e61012435613880565b50610f2b6101243561390a565b50610f38610124356139b7565b50610f81604051610f48816127c2565b8481528460208201528460408201528460608201528460808201528460a08201528460c08201528460e082015284610100820152612be7565b610fa4604051610f90816127de565b848152846020820152846040820152612ac9565b60401c1615610fb05780f35b68ff0000000000000000195f80516020613d8e83398151915254165f80516020613d8e833981519152557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b60405162461bcd60e51b815260206004820152601160248201527064656c6179203e3d206d696e53746f726560781b6044820152606490fd5b604051631afcd79f60e31b8152600490fd5b68ffffffffffffffffff19841668010000000000000001175f80516020613d8e83398151915255610e20565b60405163f92ee8a960e01b8152600490fd5b9050155f610df1565b303b159150610de9565b604086901c60ff16159150610ddf565b634e487b7160e01b5f52604160045260245ffd5b5034610277576020366003190112610277576101206110e5612710565b6110ed612ba4565b5061010060ff6110fc83612969565b5054169160ff61110b82612969565b505460081c16906001600160401b0361112382612969565b505460101c16906001600160401b0361113b82612969565b505460501c166001600160401b0361115283612969565b505460901c1661116183612969565b50600101546001600160a01b03169061117984612969565b50600201546001600160a01b03169260046111a1600361119888612969565b50015496612969565b50015495604051976111b2896127c2565b8a8952602089015260408801526060870152608086015260a085015260c084015260e08301528282015260405192835260ff60208201511660208401526001600160401b0360408201511660408401526001600160401b0360608201511660608401526001600160401b03608082015116608084015260018060a01b0360a08201511660a084015260018060a01b0360c08201511660c084015260e081015160e08401520151610100820152f35b50346102775780600319360112610277576020600554604051908152f35b503461027757602036600319011261027757606061129a612710565b60ff6112a582612969565b505416906001600160401b036112cb60ff6112bf84612969565b505460081c1692612969565b505460901c169060405192835260208301526040820152f35b503461027757806003193601126102775760206001600160401b0360025460a01c16604051908152f35b503461027757602036600319011261027757611328612768565b6001546001600160a01b03919083908316803b15610b405781809160046040518094819363919840ad60e01b83525af18015610b35576113c5575b505060206004926001541660405193848092639fa6a6e360e01b82525afa9081156113ba5761055c928492611399575b50613b4c565b6113b391925060203d602011610b1657610b088183612814565b905f611393565b6040513d85823e3d90fd5b6113ce90612794565b6108f057825f611363565b50346102775780600319360112610277577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036114445760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b60405163703e46dd60e11b8152600490fd5b5034610277576040366003190112610277576020906040611475612768565b9161147e612726565b9260018060a01b03168152600e8452206001600160401b038092165f52825260405f205416604051908152f35b506040366003190112610277576114c0612768565b6024356001600160401b0381116108f0576114df90369060040161287a565b6001600160a01b03917f000000000000000000000000000000000000000000000000000000000000000083163081149081156116d3575b5061144457335f9081527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360209081526040909120549093907f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f559060ff16156116b557506040516352d1902d60e01b8152908216938082600481885afa9182918793611681575b50506115bc57604051634c9c8ce360e01b815260048101859052602490fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc908181036116695750813b156116505780546001600160a01b03191684179055604051927fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8580a2815115611636576104049250613cbc565b505034611641575080f35b63b398979f60e01b8152600490fd5b604051634c9c8ce360e01b815260048101859052602490fd5b60249060405190632a87526960e21b82526004820152fd5b9080929350813d83116116ae575b6116998183612814565b810103126116aa5751905f8061159d565b8580fd5b503d61168f565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b9050837f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc541614155f611516565b5034610277576020366003190112610277576040608091611720612768565b816060845161172e816127a7565b8281528260208201528286820152015260018060a01b03168152600c6020522080546001600160401b03918282169260026001830154920154926060604051611776816127a7565b86815283602082019360401c168352604081019485520193845260405194855251166020840152516040830152516060820152f35b5034610277576040366003190112610277576117c561277e565b336001600160a01b038216036117e15761040490600435613acd565b60405163334bd91960e11b8152600490fd5b50346102775760403660031901126102775761040460043561181361277e565b908084525f80516020613d6e83398151915260205261183860016040862001546132b0565b613a5d565b503461027757602080600319360112610b40576002546040516327c830a960e01b8152600480359391926001600160a01b03918391859190829085165afa8015610ae557849386916118d1575b5092610a63916118bb7f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364951561290b565b338752600f845260408720610a57838254612a15565b809450838092503d831161192e575b6118ea8183612814565b81010312610b735783610a63916118bb6119247f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364966128fe565b929550509161188a565b503d6118e0565b50346102775761010036600319011261027757611950612768565b6024356001600160401b0381116108f05761196f90369060040161287a565b9061197861273c565b611980612752565b611988612898565b9060c4359260ff841684036108ec5760e435946001600160a01b0380871687036108ec57335f9081527fa9ed30483ca8e7e18edb58c7d68a44612fe881c88cebad0b52b41698e9844de260205260409020547fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc49060ff16156116b55750811615611a195761055c96604435916132de565b60405162461bcd60e51b815260206004820152600960248201526834b73b1037bbb732b960b91b6044820152606490fd5b503461027757806003193601126102775760206001600160401b0360035460801c16604051908152f35b5034610277576020366003190112610277576001604060209260043581525f80516020613d6e8339815191528452200154604051908152f35b503461027757602036600319011261027757600435906001600160401b03908183116102775750611ae8602061059a8194369060040161287a565b810160078152030190205416604051908152f35b5034610277576080366003190112610277576004356001600160401b038111610b4057611b2d90369060040161287a565b611b35612726565b9060443560ff811681036108ec57606435926001600160401b038411610b735736602385011215610b73576001600160401b03846004013511610b7357366024856004013586010111610b735760015485906001600160a01b0316803b15610b405781809160046040518094819363919840ad60e01b83525af18015610b35576123a3575b5050600154604051639fa6a6e360e01b81529290602090849060049082906001600160a01b03165afa928315612398578693612377575b50835115612344576004611c0483612969565b50015415612314576001600160401b03611c1d83612969565b505460901c166001600160401b03841610156122e557611c3c82612969565b5060ff80831691541611156122b6576001600160401b036040516020818751611c688183858c016128b8565b8101600b8152030190205416158015612278575b15612249576001600160401b038216865260086020526040862060ff82165f526020526001600160401b0360405f205416612218576001600160401b0382168652600960205260408620335f5260205260ff60405f2054166121e857611ce182613260565b92611d2e611d1e6003611cf386612969565b5001546001600160401b03611d178582611d0c8a612969565b505460901c16612a3a565b1690612a53565b6001600160401b03861690612a53565b6001600160401b038416885260116020526040882060ff84165f5260205260ff60405f2054165f146121df5787905b6004611d6886612969565b500154611d758383612a66565b116121ad57611d926004611d8887612969565b5001918254612a15565b90556004611d9f85612969565b5001611dac828254612a15565b9055611db784612969565b50600201546001600160a01b03168852600f6020526040882080549091611ddd91612a66565b90556001600160401b038316875260116020526040872060ff83165f5260205260405f2060ff198154169055611e11612a73565b936001600160401b0384168552336020860152611e3636886004013560248a01612835565b602081519101206040860152611e576001600160401b03600a541695612ac9565b611f0060405160208181611e718b838151938492016128b8565b8101600b815203019020916001600160401b03199287848254161790556001600160401b0386168a52600860205260408a2060ff86165f5260205260405f2087848254161790556001600160401b0386168a52600960205260408a20335f5260205260405f20600160ff1982541617905560036001600160401b03611ef588612969565b509216910154612a53565b916001600160401b03611f1286612969565b505460901c1691611f238233613b4c565b338a52600d60205260408a20600c6020526001600160401b0360408c205460401c165f5260205260405f208782825416179055338a52600c60205260408a20805467ffffffffffffffff60401b611f856001600160401b038360401c166129b6565b60401b169067ffffffffffffffff60401b1916179055338a52600c602052600160408b2001611fb5858254612a66565b9055338a52600e60205260408a206001600160401b0383165f5260205260405f20908154906001600160401b03611fed8184166129b6565b169116179055338952600e6020526001600160401b0360408a2091165f52602052600160405f2001612020838254612a66565b9055338852600e60205260408820905f52602052612046600160405f2001918254612a66565b9055600254604080516361e728b160e01b81523360048201526001602482015291829060449082906001600160a01b03165afa9081156121a25787908892612180575b5015610a8357338752600c6020526001604088200154908160070291600783040361216c577ff226a06b14819931288115f369b127b1a9986ea785f3a4cfcf16b0b445cf1ca59594926120e36020959360ff931015612b68565b338952600c85526001600160401b036121276121078260408d205460401c16612a22565b9782604051998a988952168888015260c0604088015260c08701906128d9565b9316606085015216608083015281810360a08301526004860135808252602487018483013786838760040135830101523395601f80199160040135011601030190a280f35b634e487b7160e01b88526011600452602488fd5b905061219b915060403d604011610ade57610acf8183612814565b905f612089565b6040513d89823e3d90fd5b60405162461bcd60e51b815260206004820152600a6024820152696e6f2062616c616e636560b01b6044820152606490fd5b60055490611d5d565b60405162461bcd60e51b8152602060048201526008602482015267323ab81030b2323960c11b6044820152606490fd5b60405162461bcd60e51b815260206004820152600960248201526865786973742070726960b81b6044820152606490fd5b60405162461bcd60e51b815260206004820152600760248201526632bc34b9ba103960c91b6044820152606490fd5b506001600160401b0360405160208187516122968183858c016128b8565b8101600b81520301902054168652601060205260ff604087205416611c7c565b60405162461bcd60e51b8152602060048201526007602482015266696e762070726960c81b6044820152606490fd5b60405162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b6044820152606490fd5b60405162461bcd60e51b81526020600482015260086024820152676e6f20706965636560c01b6044820152606490fd5b60405162461bcd60e51b815260206004820152600b60248201526a696e76207265706c69636160a81b6044820152606490fd5b61239191935060203d602011610b1657610b088183612814565b915f611bf1565b6040513d88823e3d90fd5b6123ac90612794565b610b7357845f611bba565b5034610277576040366003190112610277576123d1612768565b906123da612726565b9160018060a01b031690818152600c6020526001600160401b03806040832054169316809310612427576040918152600e60205220905f526020526020600160405f200154604051908152f35b60405162461bcd60e51b81526020600482015260096024820152680d2dcec40cae0dec6d60bb1b6044820152606490fd5b50346108ec576020806003193601126108ec57612473612710565b6002546040516327c830a960e01b815291926001600160a01b0392909182908290600490829087165afa801561267c575f90612687575b6124b591501561290b565b8160015416803b156108ec575f809160046040518094819363919840ad60e01b83525af1801561267c57612669575b50600481836001541660405192838092639fa6a6e360e01b82525afa908115610ae557859161264c575b5082600161251b86612969565b50015416330361261e576125539061253285612969565b506001600160401b0380938180935460501c168260025460a01c16906129fa565b9216911610156125f157600461256885612969565b50015461257485612cdd565b90818111156125e4576125b19161258a91612a15565b80945b600461259888612969565b50016125a5838254612a15565b905533908854166131e2565b604051931683528201527f44647ea87405778ba90a8ca28d2b772b1350bcd5950e5111bbc8c506a5c9632a60403392a280f35b50506125b185809461258d565b60405162461bcd60e51b81526004810183905260056024820152646561726c7960d81b6044820152606490fd5b60405162461bcd60e51b8152600481018390526006602482015265696e7620776360d01b6044820152606490fd5b6126639150823d8411610b1657610b088183612814565b5f61250e565b612674919450612794565b5f925f6124e4565b6040513d5f823e3d90fd5b508181813d83116126b6575b61269d8183612814565b810103126108ec576126b16124b5916128fe565b6124aa565b503d612693565b346108ec5760203660031901126108ec576004359063ffffffff60e01b82168092036108ec57602091637965db0b60e01b81149081156126ff575b5015158152f35b6301ffc9a760e01b149050836126f8565b600435906001600160401b03821682036108ec57565b602435906001600160401b03821682036108ec57565b606435906001600160401b03821682036108ec57565b608435906001600160401b03821682036108ec57565b600435906001600160a01b03821682036108ec57565b602435906001600160a01b03821682036108ec57565b6001600160401b0381116110b457604052565b608081019081106001600160401b038211176110b457604052565b61012081019081106001600160401b038211176110b457604052565b606081019081106001600160401b038211176110b457604052565b604081019081106001600160401b038211176110b457604052565b90601f801991011681019081106001600160401b038211176110b457604052565b9291926001600160401b0382116110b4576040519161285e601f8201601f191660200184612814565b8294818452818301116108ec578281602093845f960137010152565b9080601f830112156108ec5781602061289593359101612835565b90565b60a4359060ff821682036108ec57565b6024359060ff821682036108ec57565b5f5b8381106128c95750505f910152565b81810151838201526020016128ba565b906020916128f2815180928185528580860191016128b8565b601f01601f1916010190565b519081151582036108ec57565b1561291257565b60405162461bcd60e51b815260206004820152601060248201526f195b595c99d95b98de481c185d5cd95960821b6044820152606490fd5b908160209103126108ec57516001600160401b03811681036108ec5790565b6006548110156129a25760059060065f52027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01905f90565b634e487b7160e01b5f52603260045260245ffd5b9060016001600160401b03809316019182116129ce57565b634e487b7160e01b5f52601160045260245ffd5b906001600160401b038092166001019182116129ce57565b9190916001600160401b03808094169116019182116129ce57565b919082039182116129ce57565b6001600160401b039081165f1901919082116129ce57565b6001600160401b0391821690821603919082116129ce57565b818102929181159184041417156129ce57565b919082018092116129ce57565b60405190612a80826127de565b5f6040838281528260208201520152565b600a548110156129a257600a5f5260011b7fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a801905f90565b600a54600160401b8110156110b457806001612ae89201600a55612a91565b919091612b38578051825460208301516001600160e01b03199091166001600160401b039290921691909117604091821b68010000000000000000600160e01b0316178355015160019190910155565b634e487b7160e01b5f525f60045260245ffd5b91908260409103126108ec576020612b62836128fe565b92015190565b15612b6f57565b60405162461bcd60e51b815260206004820152600d60248201526c696e73756620706c656467653760981b6044820152606490fd5b60405190612bb1826127c2565b5f610100838281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e08201520152565b600654600160401b8110156110b457806001612c069201600655612969565b919091612b38578051825460208301516040840151606085015160808601516001600160d01b031990941660ff959095169490941760089290921b61ff00169190911760109190911b69ffffffffffffffff0000161760509290921b67ffffffffffffffff60501b169190911760909190911b67ffffffffffffffff60901b1617825560a08101516001830180546001600160a01b03199081166001600160a01b039384161790915560c0830151600285018054909216921691909117905560e08101516003830155610100015160049190910155565b6001805460408051639fa6a6e360e01b815260209594600494909387908390879082906001600160a01b03165afa918215612e6c575f92612e4d575b50612d2384612969565b50926001600160401b0380945460901c168484161015612e42575f97889687845b612d90575b505050505090611d17612d85849383611d17612d7a6128959a9b6003612d71612d8b9b612969565b50015490612a53565b9282611d0c89612969565b92613260565b612a66565b612d9988612969565b505460ff828116918116821015612e3b57888a1691825f5260088652865f20815f52865289875f20541615612dd7575b505050840160ff1684612d44565b90919c5f198114612e285787019c5f5260118552855f20905f528452845f20541615612e06575b5f8080612dc9565b979060ff85612e1b8196959460055490612a66565b9a92939495505050612dfe565b601185634e487b7160e01b5f525260245ffd5b5050612d49565b505f96505050505050565b612e65919250873d8911610b1657610b088183612814565b905f612d19565b83513d5f823e3d90fd5b906001600160401b03809216601f029182169182036129ce57565b906001600160401b03809116918215612ea957160490565b634e487b7160e01b5f52601260045260245ffd5b600254604080516379ca7e0f60e01b81529293926001600160a01b03925f929160209182908290600490829089165afa908115612e6c579085915f916131a8575b50163303613179576001600160401b039081871696875f526008825260ff845f20971696875f52825282845f20541692831561316e5760019680885416803b156108ec575f8091600489518094819363919840ad60e01b83525af1801561316457613150575b5080846004928a5416885193848092639fa6a6e360e01b82525afa918215613146579160809997959391899795937fd9a99dcb8bb098fb6535dccbe640188d0a0f125b786fff52c803be82df5168bb9c9a92613127575b50612fc586612a91565b5054871c16968581526010855286812060ff19908a828254161790558d8252600886528782208b5f528652875f206001600160401b031981541690558d825260098652878220895f528652875f208181541690558d8252601186528782208b5f52865289885f209182541617905580988361303f86612969565b505460901c169184841683811061306b575b5050505050505082519586528501528301526060820152a2565b869b5093879386938b611d17948e61310699600e6130a28e60049f9e61310c9f6003613099611d1792612969565b50015492613260565b9a6130ad8985613b4c565b838552600c815286868620016130c48d8254612a15565b905583855252838320905f528d5283835f20016130e28a8254612a15565b90558152600e8c5220825f528a528b5f20016130ff868254612a15565b9055612a3a565b96612969565b5001613119868254612a66565b90555f808080808080613051565b61313f919250853d8711610b1657610b088183612814565b905f612fbb565b87513d8a823e3d90fd5b61315b919750612794565b5f956004612f64565b87513d5f823e3d90fd5b505050505050505050565b606491519062461bcd60e51b82526004820152600a60248201526934b73b1031b0b63632b960b11b6044820152fd5b809250838092503d83116131db575b6131c18183612814565b810103126108ec575184811681036108ec5784905f612efe565b503d6131b7565b60405163a9059cbb60e01b5f9081526001600160a01b039384166004526024949094529260209060448180855af160015f5114811615613241575b836040521561322b57505050565b635274afe760e01b835216600482015260249150fd5b600181151661325757813b15153d15161661321d565b833d5f823e3d90fd5b6601ffffffffffff6132a76132a261062d6128959461062760ff61061c61329c6001600160401b0361329186612969565b505460101c16612a22565b93612969565b612a22565b600f1c166129e2565b805f525f80516020613d6e83398151915260205260405f20335f5260205260ff60405f205416156116b55750565b60015493975f97939692959094919392916001600160a01b0316803b156108ec575f809160046040518094819363919840ad60e01b83525af1801561267c57613866575b50600154604051639fa6a6e360e01b815293929190602090859060049082906001600160a01b03165afa93841561385b57899461383a575b50865115613809576001600160401b038a16156137d95760ff8316156137aa5760ff821660ff8416101561377c576003546001600160401b0361339f818316876129fa565b166001600160401b0383161061374b576001600160401b036133c6818360401c16876129fa565b166001600160401b038316116137195760045489106136e8576001600160401b03604051602081816133fe8d838151938492016128b8565b8101600781520301902054166136b55761342961062d61341d8d612a22565b61062760ff8816612e76565b9060ff85166001600160401b03831602906001600160401b0382169182036129ce5760801c6001600160401b0316106136825761346590612a22565b600f1c6601ffffffffffff1661347a906129e2565b6134848583612a3a565b6001600160401b0316613497908a612a53565b906001600160401b03166134aa91612a53565b60ff8316906005546134bb91612a66565b906134c591612a53565b996006546001600160401b0316976134db612ba4565b6001600160a01b0388811660a083015261010082018e905260e082018c90526001600160401b03848116604084015297881660608301819052978516608083015260ff8681168352871660208301529890981660c08901819052959761354090612be7565b604051815181906135558183602087016128b8565b810160078152036020019020896001600160401b0319825416179055604051988998610120918a5260208a01528060408a01528801613593916128d9565b9860608801526001600160401b031660808701526001600160401b031660a086015260ff1660c085015260ff1660e0840152610100830152600160a01b600190031692037fe37588ecbc243b6bdc3191c8d9743ce41ab2f677eea6aafbd2ce4cd59480bbdd91a280546040516323b872dd60e01b835233600452306024526044939093526001600160a01b031690808083815a93606492602095f1908051600114821615613661575b836040526060521561364c575050565b635274afe760e01b8252600482015260249150fd5b90600181151661367957823b15153d1516169061363c565b833d83823e3d90fd5b60405162461bcd60e51b815260206004820152600b60248201526a6578636565642073697a6560a81b6044820152606490fd5b60405162461bcd60e51b815260206004820152600b60248201526a1c1a5958d948195e1a5cdd60aa1b6044820152606490fd5b60405162461bcd60e51b81526020600482015260096024820152686d696e20707269636560b81b6044820152606490fd5b60405162461bcd60e51b815260206004820152600a60248201526932bc31b2b2b210323ab960b11b6044820152606490fd5b60405162461bcd60e51b815260206004820152600960248201526839b437b93a10323ab960b91b6044820152606490fd5b60405162461bcd60e51b8152602060048201526006602482015265696e7620727360d01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526007602482015266696e762072736b60c81b6044820152606490fd5b60405162461bcd60e51b8152602060048201526008602482015267696e762073697a6560c01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526009602482015268696e7620706965636560b81b6044820152606490fd5b61385491945060203d602011610b1657610b088183612814565b925f61335a565b6040513d8b823e3d90fd5b6004939291985061387690612794565b5f97909192613322565b6001600160a01b03165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f80516020613d6e8339815191529060ff16613904575f805260205260405f20815f5260205260405f20600160ff1982541617905533905f5f80516020613d4e8339815191528180a4600190565b50505f90565b6001600160a01b03165f8181527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5591905f80516020613d6e8339815191529060ff166139b057825f5260205260405f20815f5260205260405f20600160ff1982541617905533915f80516020613d4e8339815191525f80a4600190565b5050505f90565b6001600160a01b03165f8181527fa9ed30483ca8e7e18edb58c7d68a44612fe881c88cebad0b52b41698e9844de260205260409020547fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc491905f80516020613d6e8339815191529060ff166139b057825f5260205260405f20815f5260205260405f20600160ff1982541617905533915f80516020613d4e8339815191525f80a4600190565b90815f525f80516020613d6e8339815191528060205260405f209160018060a01b031691825f5260205260ff60405f205416155f146139b057825f5260205260405f20815f5260205260405f20600160ff1982541617905533915f80516020613d4e8339815191525f80a4600190565b90815f525f80516020613d6e8339815191528060205260405f209160018060a01b031691825f5260205260ff60405f2054165f146139b057825f5260205260405f20815f5260205260405f2060ff19815416905533917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b91909160018060a01b0316805f52600c9260208481526040805f209586546001600160401b038098828280941615613ca5575b50541694169384811015613c9b57613b96906129b6565b87811690858211613c7d57865f52600e808652845f20835f528652600180865f200154895f52858852613bcf82885f2001918254612a15565b9055885f5284875280865f200154828852865f20855f52885281875f200155885f528487528a865f2054871c16828852865f20855f528852865f20906001600160401b031982541617905560029182851015613c38575b505050613c3391506129b6565b613b96565b895f528752855f209360011901908b82116129ce57613c33948c613c7393165f528852865f20015491895f52858852865f2001918254612a66565b90555f8080613c26565b505092949550925f52525f20906001600160401b0319825416179055565b5050505050509050565b67ffffffffffffffff19168783161781555f613b7f565b905f8091602081519101845af48080613d3a575b15613cf05750506040513d81523d5f602083013e60203d82010160405290565b15613d1757604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b3d15613d28576040513d5f823e3d90fd5b60405163d6bda27560e01b8152600490fd5b503d151580613cd05750813b1515613cd056fe2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220a6e9b8571cf5bd80296124c998348233a116ba0b4bbd89d82fa426251fee9e4764736f6c63430008180033",
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

// Forged is a free data retrieval call binding the contract method 0xfad73d52.
//
// Solidity: function forged(uint64 ) view returns(bool)
func (_Piece *PieceCaller) Forged(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _Piece.contract.Call(opts, &out, "forged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Forged is a free data retrieval call binding the contract method 0xfad73d52.
//
// Solidity: function forged(uint64 ) view returns(bool)
func (_Piece *PieceSession) Forged(arg0 uint64) (bool, error) {
	return _Piece.Contract.Forged(&_Piece.CallOpts, arg0)
}

// Forged is a free data retrieval call binding the contract method 0xfad73d52.
//
// Solidity: function forged(uint64 ) view returns(bool)
func (_Piece *PieceCallerSession) Forged(arg0 uint64) (bool, error) {
	return _Piece.Contract.Forged(&_Piece.CallOpts, arg0)
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

// FreeForgedSlot is a paid mutator transaction binding the contract method 0xbaa3cb02.
//
// Solidity: function freeForgedSlot(uint64 _pi, uint8 _pri) returns()
func (_Piece *PieceTransactor) FreeForgedSlot(opts *bind.TransactOpts, _pi uint64, _pri uint8) (*types.Transaction, error) {
	return _Piece.contract.Transact(opts, "freeForgedSlot", _pi, _pri)
}

// FreeForgedSlot is a paid mutator transaction binding the contract method 0xbaa3cb02.
//
// Solidity: function freeForgedSlot(uint64 _pi, uint8 _pri) returns()
func (_Piece *PieceSession) FreeForgedSlot(_pi uint64, _pri uint8) (*types.Transaction, error) {
	return _Piece.Contract.FreeForgedSlot(&_Piece.TransactOpts, _pi, _pri)
}

// FreeForgedSlot is a paid mutator transaction binding the contract method 0xbaa3cb02.
//
// Solidity: function freeForgedSlot(uint64 _pi, uint8 _pri) returns()
func (_Piece *PieceTransactorSession) FreeForgedSlot(_pi uint64, _pri uint8) (*types.Transaction, error) {
	return _Piece.Contract.FreeForgedSlot(&_Piece.TransactOpts, _pi, _pri)
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

// PieceSlotFreedIterator is returned from FilterSlotFreed and is used to iterate over the raw logs and unpacked data for SlotFreed events raised by the Piece contract.
type PieceSlotFreedIterator struct {
	Event *PieceSlotFreed // Event containing the contract specifics and raw log

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
func (it *PieceSlotFreedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PieceSlotFreed)
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
		it.Event = new(PieceSlotFreed)
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
func (it *PieceSlotFreedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PieceSlotFreedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PieceSlotFreed represents a SlotFreed event raised by the Piece contract.
type PieceSlotFreed struct {
	Pi     uint64
	Pri    uint8
	Ri     uint64
	Store  common.Address
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlotFreed is a free log retrieval operation binding the contract event 0xd9a99dcb8bb098fb6535dccbe640188d0a0f125b786fff52c803be82df5168bb.
//
// Solidity: event SlotFreed(uint64 indexed _pi, uint8 _pri, uint64 _ri, address _store, uint256 _refund)
func (_Piece *PieceFilterer) FilterSlotFreed(opts *bind.FilterOpts, _pi []uint64) (*PieceSlotFreedIterator, error) {

	var _piRule []interface{}
	for _, _piItem := range _pi {
		_piRule = append(_piRule, _piItem)
	}

	logs, sub, err := _Piece.contract.FilterLogs(opts, "SlotFreed", _piRule)
	if err != nil {
		return nil, err
	}
	return &PieceSlotFreedIterator{contract: _Piece.contract, event: "SlotFreed", logs: logs, sub: sub}, nil
}

// WatchSlotFreed is a free log subscription operation binding the contract event 0xd9a99dcb8bb098fb6535dccbe640188d0a0f125b786fff52c803be82df5168bb.
//
// Solidity: event SlotFreed(uint64 indexed _pi, uint8 _pri, uint64 _ri, address _store, uint256 _refund)
func (_Piece *PieceFilterer) WatchSlotFreed(opts *bind.WatchOpts, sink chan<- *PieceSlotFreed, _pi []uint64) (event.Subscription, error) {

	var _piRule []interface{}
	for _, _piItem := range _pi {
		_piRule = append(_piRule, _piItem)
	}

	logs, sub, err := _Piece.contract.WatchLogs(opts, "SlotFreed", _piRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PieceSlotFreed)
				if err := _Piece.contract.UnpackLog(event, "SlotFreed", log); err != nil {
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

// ParseSlotFreed is a log parse operation binding the contract event 0xd9a99dcb8bb098fb6535dccbe640188d0a0f125b786fff52c803be82df5168bb.
//
// Solidity: event SlotFreed(uint64 indexed _pi, uint8 _pri, uint64 _ri, address _store, uint256 _refund)
func (_Piece *PieceFilterer) ParseSlotFreed(log types.Log) (*PieceSlotFreed, error) {
	event := new(PieceSlotFreed)
	if err := _Piece.contract.UnpackLog(event, "SlotFreed", log); err != nil {
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
