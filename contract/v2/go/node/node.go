// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package node

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

// INodeNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type INodeNodeInfo struct {
	NodeType     uint8
	IsActive     bool
	ExitEpoch    uint64
	StakedAmount *big.Int
	LockedAmount *big.Int
}

// NodeMetaData contains all meta data concerning the Node contract.
var NodeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"check\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emergencyPause\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEmergencyPause\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emergencyPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEpoch\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eproof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_epoch\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lock\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minStake\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStakeOf\",\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeInfoOf\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structINode.NodeInfo\",\"components\":[{\"name\":\"nodeType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"exitEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nodeType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"exitEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"punish\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rsproof\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"set\",\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"money\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"_eproof\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rsproof\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDelay\",\"inputs\":[{\"name\":\"_delay\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmergencyPause\",\"inputs\":[{\"name\":\"p\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"terminate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlock\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"m\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DelayUpdated\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyPauseSet\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Punish\",\"inputs\":[{\"name\":\"_a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_typ\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_money\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Set\",\"inputs\":[{\"name\":\"_type\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_type\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Terminated\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"m\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a080604052346100cc57306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100bd57506001600160401b036002600160401b031982821601610078575b604051611d0d90816100d18239608051818181610b200152610c2f0152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f8080610059565b63f92ee8a960e01b8152600490fd5b5f80fdfe6080604081815260049182361015610015575f80fd5b5f925f3560e01c91826301ffc9a714611585575081630c08bf88146114245781630fdefd361461135e578163189a5a17146112eb578163248a9ca3146112b457816327c830a91461128e578163282d3fdf146111e75781632e1a7d4d14610f925781632f2ff15d14610f6857816334c2028614610f0057816336568abe14610eba5781633b58524d14610e6057816348ab5e6c14610de65781634f1ef28614610bad57816351858e2714610b8457816352d1902d14610b0b57816361e728b114610ac95781636a42b8f814610aa357816379ca7e0f14610a7a5781637eee288d14610a1357816381cc0c7a146109ea578163900cf0cf146109c157816391d14854146109715781639748dcdc14610774578163a217fddf14610759578163ad3cb1cc1461071c578163c0c53b8b1461056d578163c107330214610428578163ccc57490146103ed578163d3748dc314610371578163d547741f146103a1578163d8c5115514610371578163dd752e551461020d57508063fc0c546a146101e65763ffa1ad74146101a3575f80fd5b346101e257816003193601126101e25780516101de916101c282611655565b60058252640322e302e360dc1b60208301525191829182611691565b0390f35b5080fd5b50346101e257816003193601126101e25790548151911c6001600160a01b03168152602090f35b90503461036d578160031936011261036d57610227611603565b835483516323b872dd60e01b86523384523060249081523560448190529294929391831c6001600160a01b031660208760648180855af1600188511481161561034e575b828552876060521561033c575050916102ee9160017f3cf14181ae25669a913d72411736fc5c01f538fa503e963b0b2e56bcefb3edaf9433885260056020528288209081549060ff8260081c16159081610317575b506102f4575b50016102d38382546117c4565b90555160ff9094168452602084015233929081906040820190565b0390a280f35b61ffff191660ff8816176101001769ffffffffffffffff0000191681555f6102c6565b905061032686858501546117c4565b9060ff8a168b52602052848a205411155f6102c0565b602492635274afe760e01b8352820152fd5b600181151661036457813b15153d15161661026b565b823d89823e3d90fd5b8280fd5b90503461036d57602036600319011261036d57602092829160ff610393611603565b168252845220549051908152f35b9190503461036d578060031936011261036d576103e991356103e460016103c66115ed565b938387525f80516020611cb883398151915260205286200154611999565b611ba7565b5080f35b5050346101e257816003193601126101e257602090517f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f558152f35b839150346101e2576020908160031936011261036d578035906001600160401b038083168093036105695761045b611920565b6002546001600160a01b03168015158061055f575b6104b3575b5050507f7fd1e355d2fb7633f0bf87a494651b9fdd77c54aa4b8f425ff8c4889d34833f49293816001600160401b031986541617855551908152a180f35b84839188519283809263430d0a0960e11b82525afa908115610555578691610528575b501682106104e5578080610475565b845162461bcd60e51b8152908101839052601c60248201527f64656c61792062656c6f77206368616c6c656e67652077696e646f77000000006044820152606490fd5b6105489150853d871161054e575b6105408183611670565b8101906116d8565b876104d6565b503d610536565b87513d88823e3d90fd5b50803b1515610470565b8480fd5b90503461036d57606036600319011261036d576105886115d7565b906105916115ed565b6001600160a01b039260443591908483168303610718577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009485549460ff86891c1615956001600160401b03811680159081610710575b6001149081610706575b1590816106fd575b506106ee5767ffffffffffffffff1981166001178855866106cf575b5060ff8754891c16156106c157508754600180546001600160a01b031916949092169390931790556001600160e01b031990911690851b68010000000000000000600160e01b031617600717855561067790610671816119c7565b50611a64565b50610680578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f808280f35b8751631afcd79f60e31b8152fd5b68ffffffffffffffffff1916680100000000000000011787555f610616565b50875163f92ee8a960e01b8152fd5b9050155f6105fa565b303b1591506105f2565b8891506105e8565b5f80fd5b5050346101e257816003193601126101e25780516101de9161073d82611655565b60058252640352e302e360dc1b60208301525191829182611691565b5050346101e257816003193601126101e25751908152602090f35b90503461036d57606036600319011261036d5761078f6115d7565b906107986115ed565b91604435926107ae6107a86116f7565b156117e5565b6002546001600160a01b03928391821633148015610964575b6107d09061178b565b1693848752600560205285872093600185016107ed838254611824565b81556107ff8385878c548c1c166118a2565b5485549060ff82168a5282602052888a20541161085c575b50509254945160ff9095168552602085019290925216917fc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c329080604081015b0390a380f35b61ff0092919219168555878460015416803b156101e257818091858b518094819363919840ad60e01b83525af1801561095a57610942575b505060208460015416885193848092639fa6a6e360e01b82525afa8015610938577fc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c329561090f60ff92610856958c91610919575b50825469ffffffffffffffff0000191660109190911b69ffffffffffffffff000016178255565b9550819250610817565b610932915060203d60201161054e576105408183611670565b5f6108e8565b87513d8a823e3d90fd5b61094b90611613565b61095657875f610894565b8780fd5b89513d84823e3d90fd5b50600354821633146107c7565b90503461036d578160031936011261036d578160209361098f6115ed565b923581525f80516020611cb88339815191528552209060018060a01b03165f52825260ff815f20541690519015158152f35b5050346101e257816003193601126101e25760015490516001600160a01b039091168152602090f35b5050346101e257816003193601126101e25760025490516001600160a01b039091168152602090f35b5050346101e257806003193601126101e257600290610a306115d7565b82546001600160a01b039190821633148015610a6d575b610a509061178b565b1683526005602052822001610a686024358254611824565b905580f35b5060035482163314610a47565b5050346101e257816003193601126101e25760035490516001600160a01b039091168152602090f35b5050346101e257816003193601126101e2576001600160401b0360209254169051908152f35b828434610b085781600319360112610b085750610ae46115d7565b60243560ff8116810361071857610afa91611831565b825191151582526020820152f35b80fd5b828434610b085780600319360112610b0857507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610b7757602090517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b5163703e46dd60e11b8152fd5b5050346101e257816003193601126101e25760065490516001600160a01b039091168152602090f35b9180915060031936011261036d57610bc36115d7565b9060249384356001600160401b039384821161036d573660238301121561036d5781860135948511610dd457602094845192610c0887601f19601f8501160185611670565b81845236898383010111610569578185928a89930183870137840101526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116308114908115610da6575b50610d9657610c68611920565b84516352d1902d60e01b815290821695808289818a5afa9182918693610d66575b5050610ca457505050505191634c9c8ce360e01b8352820152fd5b9093919492967f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc91828103610d515750843b15610d3d575080546001600160a01b0319168317905551907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8680a2825115610d2457506103e99250611c26565b91505034610d3157505080f35b63b398979f60e01b8152fd5b8251634c9c8ce360e01b8152808801859052fd5b8351632a87526960e21b815280890191909152fd5b9080929350813d8311610d8f575b610d7e8183611670565b810103126105695751905f80610c89565b503d610d74565b845163703e46dd60e11b81528790fd5b9050817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc541614155f610c5b565b86604187634e487b7160e01b5f52525ffd5b90503461036d578160031936011261036d577fc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef6788191610e5a610e25611603565b9160243593610e32611920565b60ff84168752602052838187205551928392836020909392919360ff60408201951681520152565b0390a180f35b5050346101e257366003190112610b0857610e796115d7565b610e816115ed565b610e89611920565b60018060a01b0390816bffffffffffffffffffffffff60a01b93168360025416176002551690600354161760035580f35b8383346101e257806003193601126101e257610ed46115ed565b90336001600160a01b03831603610ef157506103e9919235611ba7565b5163334bd91960e11b81528390fd5b8334610b08576020366003190112610b0857610f1a6115d7565b610f22611920565b600680546001600160a01b0319166001600160a01b039290921691821790557fc32df8d34dee7f1c985371a6d6c56a54baf4e051bd1cc6f005d4eb28afb9d23f8280a280f35b9190503461036d578060031936011261036d576103e99135610f8d60016103c66115ed565b611b24565b9190503461036d57602090816003193601126111e357823592610fb66107a86116f7565b338552600583528185206001546001600160a01b03929087908416803b156101e2578180918488518094819363919840ad60e01b83525af180156111d9576111c1575b5050815460ff8160081c168015611139575b156110be5760ff60018401549116885281865261103961102e868a2054896117c4565b6002850154906117c4565b1161108f5750847f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d59493926001611085935b01611077838254611824565b905533908854851c166118a2565b519283523392a280f35b835162461bcd60e51b81529081018590526009602482015268696e73756620706c6560b81b6044820152606490fd5b5060018201546110d26002840154886117c4565b116111085750847f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d594939260016110859361106b565b835162461bcd60e51b8152908101859052600b60248201526a1a5b9cdd59881c1b19481d60aa1b6044820152606490fd5b5081868560015416875192838092639fa6a6e360e01b82525afa9081156111b757899161119a575b506001600160401b03808a5416818460101c16018181116111875781169116111561100b565b634e487b7160e01b8b526011855260248bfd5b6111b19150873d891161054e576105408183611670565b5f611161565b86513d8b823e3d90fd5b6111ca90611613565b6111d557865f610ff9565b8680fd5b86513d84823e3d90fd5b8380fd5b9190503461036d578060031936011261036d576112026115d7565b6002546001600160a01b039190821633148015611281575b6112239061178b565b16835260056020528083206001600282019161124260243584546117c4565b809355015410611250578280f35b906020606492519162461bcd60e51b8352820152600b60248201526a1b1bd8dac8195e18d9595960aa1b6044820152fd5b506003548216331461121a565b5050346101e257816003193601126101e2576020906112ab6116f7565b90519015158152f35b90503461036d57602036600319011261036d57816020936001923581525f80516020611cb883398151915285522001549051908152f35b5050346101e25760203660031901126101e25760a09181906001600160a01b036113136115d7565b1681526005602052209081549160026001820154910154916001600160401b0381519460ff8116865260ff8160081c161515602087015260101c169084015260608301526080820152f35b5050346101e25760203660031901126101e2578060a09261137d6115d7565b816080845161138b8161163a565b82815282602082015282868201528260608201520152600180861b031681526005602052208151906113bc8261163a565b80549260ff841693848452602084019060ff8160081c16151582526001600160401b0391828487019260101c16825260806002600187015496606089019788520154960195865283519687525115156020870152511690840152516060830152516080820152f35b905034610718575f36600319011261071857335f526005602052815f20805460ff8160081c16156115555761ff00191681556001546001600160a01b03908116803b15610718575f80918587518094819363919840ad60e01b83525af1801561154b57611537575b506001548451639fa6a6e360e01b81529360209285928391165afa91821561152a576114e193508492611509575b509069ffffffffffffffff000082549160101b169069ffffffffffffffff00001916179055565b337f98cd97fc1a1cc958cbd729b1bb531d4f3ea4925470bf23ea9af386640cbd07be8280a280f35b61152391925060203d60201161054e576105408183611670565b905f6114ba565b50505051903d90823e3d90fd5b611542919550611613565b5f93602061148c565b85513d5f823e3d90fd5b835162461bcd60e51b8152602081850152600a6024820152696e6f742061637469766560b01b6044820152606490fd5b903461071857602036600319011261071857359063ffffffff60e01b821680920361071857602091637965db0b60e01b81149081156115c6575b5015158152f35b6301ffc9a760e01b149050836115bf565b600435906001600160a01b038216820361071857565b602435906001600160a01b038216820361071857565b6004359060ff8216820361071857565b6001600160401b03811161162657604052565b634e487b7160e01b5f52604160045260245ffd5b60a081019081106001600160401b0382111761162657604052565b604081019081106001600160401b0382111761162657604052565b90601f801991011681019081106001600160401b0382111761162657604052565b602080825282518183018190529093925f5b8281106116c457505060409293505f838284010152601f8019910116010190565b8181018601518482016040015285016116a3565b9081602091031261071857516001600160401b03811681036107185790565b6006546001600160a01b031680151580611781575b61171557505f90565b6020600491604051928380926358c3de9360e11b82525afa908115611776575f9161173e575090565b90506020813d60201161176e575b8161175960209383611670565b81010312610718575180151581036107185790565b3d915061174c565b6040513d5f823e3d90fd5b50803b151561170c565b1561179257565b60405162461bcd60e51b815260206004820152600a60248201526934b73b1031b0b63632b960b11b6044820152606490fd5b919082018092116117d157565b634e487b7160e01b5f52601160045260245ffd5b156117ec57565b60405162461bcd60e51b815260206004820152601060248201526f195b595c99d95b98de481c185d5cd95960821b6044820152606490fd5b919082039182116117d157565b60018060a01b03165f52600560205260405f2090815460ff8160081c1615918215611890575b50506118895760018101546002909101548082106118805761187891611824565b905b60019190565b50505f9061187a565b505f905f90565b60ff9192508116911614155f80611857565b60405163a9059cbb60e01b5f9081526001600160a01b039384166004526024949094529260209060448180855af160015f5114811615611901575b83604052156118eb57505050565b635274afe760e01b835216600482015260249150fd5b600181151661191757813b15153d1516166118dd565b833d5f823e3d90fd5b335f9081527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f559060ff161561197b5750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b805f525f80516020611cb883398151915260205260405f20335f5260205260ff60405f2054161561197b5750565b6001600160a01b03165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f80516020611cb88339815191529060ff16611a5e575f805260205260405f20815f5260205260405f20600160ff1982541617905533905f7f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50505f90565b6001600160a01b03165f8181527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5591905f80516020611cb88339815191529060ff16611b1d57825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b5050505f90565b90815f525f80516020611cb88339815191528060205260405f209160018060a01b031691825f5260205260ff60405f205416155f14611b1d57825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b90815f525f80516020611cb88339815191528060205260405f209160018060a01b031691825f5260205260ff60405f2054165f14611b1d57825f5260205260405f20815f5260205260405f2060ff19815416905533917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b905f8091602081519101845af48080611ca4575b15611c5a5750506040513d81523d5f602083013e60203d82010160405290565b15611c8157604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b3d15611c92576040513d5f823e3d90fd5b60405163d6bda27560e01b8152600490fd5b503d151580611c3a5750813b1515611c3a56fe02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a26469706673582212209994043a1164b75c7ee0d220074282c59459c9ddf2af21d63081cf2ad346d22264736f6c63430008180033",
}

// NodeABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeMetaData.ABI instead.
var NodeABI = NodeMetaData.ABI

// NodeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NodeMetaData.Bin instead.
var NodeBin = NodeMetaData.Bin

// DeployNode deploys a new Ethereum contract, binding an instance of Node to it.
func DeployNode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Node, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// Node is an auto generated Go binding around an Ethereum contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Node *NodeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Node *NodeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Node.Contract.DEFAULTADMINROLE(&_Node.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Node *NodeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Node.Contract.DEFAULTADMINROLE(&_Node.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Node *NodeCaller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Node *NodeSession) GOVERNORROLE() ([32]byte, error) {
	return _Node.Contract.GOVERNORROLE(&_Node.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Node *NodeCallerSession) GOVERNORROLE() ([32]byte, error) {
	return _Node.Contract.GOVERNORROLE(&_Node.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Node *NodeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Node *NodeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Node.Contract.UPGRADEINTERFACEVERSION(&_Node.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Node *NodeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Node.Contract.UPGRADEINTERFACEVERSION(&_Node.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Node *NodeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Node *NodeSession) VERSION() (string, error) {
	return _Node.Contract.VERSION(&_Node.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Node *NodeCallerSession) VERSION() (string, error) {
	return _Node.Contract.VERSION(&_Node.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(bool, uint256)
func (_Node *NodeCaller) Check(opts *bind.CallOpts, _a common.Address, _type uint8) (bool, *big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "check", _a, _type)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(bool, uint256)
func (_Node *NodeSession) Check(_a common.Address, _type uint8) (bool, *big.Int, error) {
	return _Node.Contract.Check(&_Node.CallOpts, _a, _type)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(bool, uint256)
func (_Node *NodeCallerSession) Check(_a common.Address, _type uint8) (bool, *big.Int, error) {
	return _Node.Contract.Check(&_Node.CallOpts, _a, _type)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Node *NodeCaller) Delay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Node *NodeSession) Delay() (uint64, error) {
	return _Node.Contract.Delay(&_Node.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint64)
func (_Node *NodeCallerSession) Delay() (uint64, error) {
	return _Node.Contract.Delay(&_Node.CallOpts)
}

// EmergencyPause is a free data retrieval call binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() view returns(address)
func (_Node *NodeCaller) EmergencyPause(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "emergencyPause")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyPause is a free data retrieval call binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() view returns(address)
func (_Node *NodeSession) EmergencyPause() (common.Address, error) {
	return _Node.Contract.EmergencyPause(&_Node.CallOpts)
}

// EmergencyPause is a free data retrieval call binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() view returns(address)
func (_Node *NodeCallerSession) EmergencyPause() (common.Address, error) {
	return _Node.Contract.EmergencyPause(&_Node.CallOpts)
}

// EmergencyPaused is a free data retrieval call binding the contract method 0x27c830a9.
//
// Solidity: function emergencyPaused() view returns(bool)
func (_Node *NodeCaller) EmergencyPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "emergencyPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyPaused is a free data retrieval call binding the contract method 0x27c830a9.
//
// Solidity: function emergencyPaused() view returns(bool)
func (_Node *NodeSession) EmergencyPaused() (bool, error) {
	return _Node.Contract.EmergencyPaused(&_Node.CallOpts)
}

// EmergencyPaused is a free data retrieval call binding the contract method 0x27c830a9.
//
// Solidity: function emergencyPaused() view returns(bool)
func (_Node *NodeCallerSession) EmergencyPaused() (bool, error) {
	return _Node.Contract.EmergencyPaused(&_Node.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Node *NodeCaller) Epoch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Node *NodeSession) Epoch() (common.Address, error) {
	return _Node.Contract.Epoch(&_Node.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(address)
func (_Node *NodeCallerSession) Epoch() (common.Address, error) {
	return _Node.Contract.Epoch(&_Node.CallOpts)
}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Node *NodeCaller) Eproof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "eproof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Node *NodeSession) Eproof() (common.Address, error) {
	return _Node.Contract.Eproof(&_Node.CallOpts)
}

// Eproof is a free data retrieval call binding the contract method 0x81cc0c7a.
//
// Solidity: function eproof() view returns(address)
func (_Node *NodeCallerSession) Eproof() (common.Address, error) {
	return _Node.Contract.Eproof(&_Node.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Node *NodeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Node *NodeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Node.Contract.GetRoleAdmin(&_Node.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Node *NodeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Node.Contract.GetRoleAdmin(&_Node.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Node *NodeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Node *NodeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Node.Contract.HasRole(&_Node.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Node *NodeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Node.Contract.HasRole(&_Node.CallOpts, role, account)
}

// MinStake is a free data retrieval call binding the contract method 0xd8c51155.
//
// Solidity: function minStake(uint8 ) view returns(uint256)
func (_Node *NodeCaller) MinStake(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "minStake", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0xd8c51155.
//
// Solidity: function minStake(uint8 ) view returns(uint256)
func (_Node *NodeSession) MinStake(arg0 uint8) (*big.Int, error) {
	return _Node.Contract.MinStake(&_Node.CallOpts, arg0)
}

// MinStake is a free data retrieval call binding the contract method 0xd8c51155.
//
// Solidity: function minStake(uint8 ) view returns(uint256)
func (_Node *NodeCallerSession) MinStake(arg0 uint8) (*big.Int, error) {
	return _Node.Contract.MinStake(&_Node.CallOpts, arg0)
}

// MinStakeOf is a free data retrieval call binding the contract method 0xd3748dc3.
//
// Solidity: function minStakeOf(uint8 _type) view returns(uint256)
func (_Node *NodeCaller) MinStakeOf(opts *bind.CallOpts, _type uint8) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "minStakeOf", _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeOf is a free data retrieval call binding the contract method 0xd3748dc3.
//
// Solidity: function minStakeOf(uint8 _type) view returns(uint256)
func (_Node *NodeSession) MinStakeOf(_type uint8) (*big.Int, error) {
	return _Node.Contract.MinStakeOf(&_Node.CallOpts, _type)
}

// MinStakeOf is a free data retrieval call binding the contract method 0xd3748dc3.
//
// Solidity: function minStakeOf(uint8 _type) view returns(uint256)
func (_Node *NodeCallerSession) MinStakeOf(_type uint8) (*big.Int, error) {
	return _Node.Contract.MinStakeOf(&_Node.CallOpts, _type)
}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address a) view returns((uint8,bool,uint64,uint256,uint256))
func (_Node *NodeCaller) NodeInfoOf(opts *bind.CallOpts, a common.Address) (INodeNodeInfo, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "nodeInfoOf", a)

	if err != nil {
		return *new(INodeNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(INodeNodeInfo)).(*INodeNodeInfo)

	return out0, err

}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address a) view returns((uint8,bool,uint64,uint256,uint256))
func (_Node *NodeSession) NodeInfoOf(a common.Address) (INodeNodeInfo, error) {
	return _Node.Contract.NodeInfoOf(&_Node.CallOpts, a)
}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address a) view returns((uint8,bool,uint64,uint256,uint256))
func (_Node *NodeCallerSession) NodeInfoOf(a common.Address) (INodeNodeInfo, error) {
	return _Node.Contract.NodeInfoOf(&_Node.CallOpts, a)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint8 nodeType, bool isActive, uint64 exitEpoch, uint256 stakedAmount, uint256 lockedAmount)
func (_Node *NodeCaller) Nodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	NodeType     uint8
	IsActive     bool
	ExitEpoch    uint64
	StakedAmount *big.Int
	LockedAmount *big.Int
}, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "nodes", arg0)

	outstruct := new(struct {
		NodeType     uint8
		IsActive     bool
		ExitEpoch    uint64
		StakedAmount *big.Int
		LockedAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.IsActive = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ExitEpoch = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.StakedAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockedAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint8 nodeType, bool isActive, uint64 exitEpoch, uint256 stakedAmount, uint256 lockedAmount)
func (_Node *NodeSession) Nodes(arg0 common.Address) (struct {
	NodeType     uint8
	IsActive     bool
	ExitEpoch    uint64
	StakedAmount *big.Int
	LockedAmount *big.Int
}, error) {
	return _Node.Contract.Nodes(&_Node.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint8 nodeType, bool isActive, uint64 exitEpoch, uint256 stakedAmount, uint256 lockedAmount)
func (_Node *NodeCallerSession) Nodes(arg0 common.Address) (struct {
	NodeType     uint8
	IsActive     bool
	ExitEpoch    uint64
	StakedAmount *big.Int
	LockedAmount *big.Int
}, error) {
	return _Node.Contract.Nodes(&_Node.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Node *NodeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Node *NodeSession) ProxiableUUID() ([32]byte, error) {
	return _Node.Contract.ProxiableUUID(&_Node.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Node *NodeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Node.Contract.ProxiableUUID(&_Node.CallOpts)
}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Node *NodeCaller) Rsproof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "rsproof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Node *NodeSession) Rsproof() (common.Address, error) {
	return _Node.Contract.Rsproof(&_Node.CallOpts)
}

// Rsproof is a free data retrieval call binding the contract method 0x79ca7e0f.
//
// Solidity: function rsproof() view returns(address)
func (_Node *NodeCallerSession) Rsproof() (common.Address, error) {
	return _Node.Contract.Rsproof(&_Node.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Node *NodeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Node *NodeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Node.Contract.SupportsInterface(&_Node.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Node *NodeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Node.Contract.SupportsInterface(&_Node.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Node *NodeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Node *NodeSession) Token() (common.Address, error) {
	return _Node.Contract.Token(&_Node.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Node *NodeCallerSession) Token() (common.Address, error) {
	return _Node.Contract.Token(&_Node.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Node *NodeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Node *NodeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Node.Contract.GrantRole(&_Node.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Node *NodeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Node.Contract.GrantRole(&_Node.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _token, address _epoch, address initialOwner) returns()
func (_Node *NodeTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _epoch common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "initialize", _token, _epoch, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _token, address _epoch, address initialOwner) returns()
func (_Node *NodeSession) Initialize(_token common.Address, _epoch common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.Initialize(&_Node.TransactOpts, _token, _epoch, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _token, address _epoch, address initialOwner) returns()
func (_Node *NodeTransactorSession) Initialize(_token common.Address, _epoch common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.Initialize(&_Node.TransactOpts, _token, _epoch, initialOwner)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _from, uint256 _m) returns()
func (_Node *NodeTransactor) Lock(opts *bind.TransactOpts, _from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "lock", _from, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _from, uint256 _m) returns()
func (_Node *NodeSession) Lock(_from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Lock(&_Node.TransactOpts, _from, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _from, uint256 _m) returns()
func (_Node *NodeTransactorSession) Lock(_from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Lock(&_Node.TransactOpts, _from, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x9748dcdc.
//
// Solidity: function punish(address _from, address _to, uint256 _m) returns()
func (_Node *NodeTransactor) Punish(opts *bind.TransactOpts, _from common.Address, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "punish", _from, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x9748dcdc.
//
// Solidity: function punish(address _from, address _to, uint256 _m) returns()
func (_Node *NodeSession) Punish(_from common.Address, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Punish(&_Node.TransactOpts, _from, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x9748dcdc.
//
// Solidity: function punish(address _from, address _to, uint256 _m) returns()
func (_Node *NodeTransactorSession) Punish(_from common.Address, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Punish(&_Node.TransactOpts, _from, _to, _m)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Node *NodeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Node *NodeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Node.Contract.RenounceRole(&_Node.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Node *NodeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Node.Contract.RenounceRole(&_Node.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Node *NodeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Node *NodeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Node.Contract.RevokeRole(&_Node.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Node *NodeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Node.Contract.RevokeRole(&_Node.TransactOpts, role, account)
}

// Set is a paid mutator transaction binding the contract method 0x48ab5e6c.
//
// Solidity: function set(uint8 _type, uint256 money) returns()
func (_Node *NodeTransactor) Set(opts *bind.TransactOpts, _type uint8, money *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "set", _type, money)
}

// Set is a paid mutator transaction binding the contract method 0x48ab5e6c.
//
// Solidity: function set(uint8 _type, uint256 money) returns()
func (_Node *NodeSession) Set(_type uint8, money *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Set(&_Node.TransactOpts, _type, money)
}

// Set is a paid mutator transaction binding the contract method 0x48ab5e6c.
//
// Solidity: function set(uint8 _type, uint256 money) returns()
func (_Node *NodeTransactorSession) Set(_type uint8, money *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Set(&_Node.TransactOpts, _type, money)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address _eproof, address _rsproof) returns()
func (_Node *NodeTransactor) SetAddress(opts *bind.TransactOpts, _eproof common.Address, _rsproof common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "setAddress", _eproof, _rsproof)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address _eproof, address _rsproof) returns()
func (_Node *NodeSession) SetAddress(_eproof common.Address, _rsproof common.Address) (*types.Transaction, error) {
	return _Node.Contract.SetAddress(&_Node.TransactOpts, _eproof, _rsproof)
}

// SetAddress is a paid mutator transaction binding the contract method 0x3b58524d.
//
// Solidity: function setAddress(address _eproof, address _rsproof) returns()
func (_Node *NodeTransactorSession) SetAddress(_eproof common.Address, _rsproof common.Address) (*types.Transaction, error) {
	return _Node.Contract.SetAddress(&_Node.TransactOpts, _eproof, _rsproof)
}

// SetDelay is a paid mutator transaction binding the contract method 0xc1073302.
//
// Solidity: function setDelay(uint64 _delay) returns()
func (_Node *NodeTransactor) SetDelay(opts *bind.TransactOpts, _delay uint64) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "setDelay", _delay)
}

// SetDelay is a paid mutator transaction binding the contract method 0xc1073302.
//
// Solidity: function setDelay(uint64 _delay) returns()
func (_Node *NodeSession) SetDelay(_delay uint64) (*types.Transaction, error) {
	return _Node.Contract.SetDelay(&_Node.TransactOpts, _delay)
}

// SetDelay is a paid mutator transaction binding the contract method 0xc1073302.
//
// Solidity: function setDelay(uint64 _delay) returns()
func (_Node *NodeTransactorSession) SetDelay(_delay uint64) (*types.Transaction, error) {
	return _Node.Contract.SetDelay(&_Node.TransactOpts, _delay)
}

// SetEmergencyPause is a paid mutator transaction binding the contract method 0x34c20286.
//
// Solidity: function setEmergencyPause(address p) returns()
func (_Node *NodeTransactor) SetEmergencyPause(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "setEmergencyPause", p)
}

// SetEmergencyPause is a paid mutator transaction binding the contract method 0x34c20286.
//
// Solidity: function setEmergencyPause(address p) returns()
func (_Node *NodeSession) SetEmergencyPause(p common.Address) (*types.Transaction, error) {
	return _Node.Contract.SetEmergencyPause(&_Node.TransactOpts, p)
}

// SetEmergencyPause is a paid mutator transaction binding the contract method 0x34c20286.
//
// Solidity: function setEmergencyPause(address p) returns()
func (_Node *NodeTransactorSession) SetEmergencyPause(p common.Address) (*types.Transaction, error) {
	return _Node.Contract.SetEmergencyPause(&_Node.TransactOpts, p)
}

// Stake is a paid mutator transaction binding the contract method 0xdd752e55.
//
// Solidity: function stake(uint8 _type, uint256 m) returns()
func (_Node *NodeTransactor) Stake(opts *bind.TransactOpts, _type uint8, m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "stake", _type, m)
}

// Stake is a paid mutator transaction binding the contract method 0xdd752e55.
//
// Solidity: function stake(uint8 _type, uint256 m) returns()
func (_Node *NodeSession) Stake(_type uint8, m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Stake(&_Node.TransactOpts, _type, m)
}

// Stake is a paid mutator transaction binding the contract method 0xdd752e55.
//
// Solidity: function stake(uint8 _type, uint256 m) returns()
func (_Node *NodeTransactorSession) Stake(_type uint8, m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Stake(&_Node.TransactOpts, _type, m)
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Node *NodeTransactor) Terminate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "terminate")
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Node *NodeSession) Terminate() (*types.Transaction, error) {
	return _Node.Contract.Terminate(&_Node.TransactOpts)
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Node *NodeTransactorSession) Terminate() (*types.Transaction, error) {
	return _Node.Contract.Terminate(&_Node.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _from, uint256 _m) returns()
func (_Node *NodeTransactor) Unlock(opts *bind.TransactOpts, _from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "unlock", _from, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _from, uint256 _m) returns()
func (_Node *NodeSession) Unlock(_from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Unlock(&_Node.TransactOpts, _from, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _from, uint256 _m) returns()
func (_Node *NodeTransactorSession) Unlock(_from common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Unlock(&_Node.TransactOpts, _from, _m)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Node *NodeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Node *NodeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Node.Contract.UpgradeToAndCall(&_Node.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Node *NodeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Node.Contract.UpgradeToAndCall(&_Node.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 m) returns()
func (_Node *NodeTransactor) Withdraw(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "withdraw", m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 m) returns()
func (_Node *NodeSession) Withdraw(m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Withdraw(&_Node.TransactOpts, m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 m) returns()
func (_Node *NodeTransactorSession) Withdraw(m *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Withdraw(&_Node.TransactOpts, m)
}

// NodeDelayUpdatedIterator is returned from FilterDelayUpdated and is used to iterate over the raw logs and unpacked data for DelayUpdated events raised by the Node contract.
type NodeDelayUpdatedIterator struct {
	Event *NodeDelayUpdated // Event containing the contract specifics and raw log

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
func (it *NodeDelayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDelayUpdated)
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
		it.Event = new(NodeDelayUpdated)
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
func (it *NodeDelayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDelayUpdated represents a DelayUpdated event raised by the Node contract.
type NodeDelayUpdated struct {
	NewDelay uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayUpdated is a free log retrieval operation binding the contract event 0x7fd1e355d2fb7633f0bf87a494651b9fdd77c54aa4b8f425ff8c4889d34833f4.
//
// Solidity: event DelayUpdated(uint64 newDelay)
func (_Node *NodeFilterer) FilterDelayUpdated(opts *bind.FilterOpts) (*NodeDelayUpdatedIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "DelayUpdated")
	if err != nil {
		return nil, err
	}
	return &NodeDelayUpdatedIterator{contract: _Node.contract, event: "DelayUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayUpdated is a free log subscription operation binding the contract event 0x7fd1e355d2fb7633f0bf87a494651b9fdd77c54aa4b8f425ff8c4889d34833f4.
//
// Solidity: event DelayUpdated(uint64 newDelay)
func (_Node *NodeFilterer) WatchDelayUpdated(opts *bind.WatchOpts, sink chan<- *NodeDelayUpdated) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "DelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDelayUpdated)
				if err := _Node.contract.UnpackLog(event, "DelayUpdated", log); err != nil {
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

// ParseDelayUpdated is a log parse operation binding the contract event 0x7fd1e355d2fb7633f0bf87a494651b9fdd77c54aa4b8f425ff8c4889d34833f4.
//
// Solidity: event DelayUpdated(uint64 newDelay)
func (_Node *NodeFilterer) ParseDelayUpdated(log types.Log) (*NodeDelayUpdated, error) {
	event := new(NodeDelayUpdated)
	if err := _Node.contract.UnpackLog(event, "DelayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeEmergencyPauseSetIterator is returned from FilterEmergencyPauseSet and is used to iterate over the raw logs and unpacked data for EmergencyPauseSet events raised by the Node contract.
type NodeEmergencyPauseSetIterator struct {
	Event *NodeEmergencyPauseSet // Event containing the contract specifics and raw log

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
func (it *NodeEmergencyPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeEmergencyPauseSet)
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
		it.Event = new(NodeEmergencyPauseSet)
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
func (it *NodeEmergencyPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeEmergencyPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeEmergencyPauseSet represents a EmergencyPauseSet event raised by the Node contract.
type NodeEmergencyPauseSet struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyPauseSet is a free log retrieval operation binding the contract event 0xc32df8d34dee7f1c985371a6d6c56a54baf4e051bd1cc6f005d4eb28afb9d23f.
//
// Solidity: event EmergencyPauseSet(address indexed pauser)
func (_Node *NodeFilterer) FilterEmergencyPauseSet(opts *bind.FilterOpts, pauser []common.Address) (*NodeEmergencyPauseSetIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "EmergencyPauseSet", pauserRule)
	if err != nil {
		return nil, err
	}
	return &NodeEmergencyPauseSetIterator{contract: _Node.contract, event: "EmergencyPauseSet", logs: logs, sub: sub}, nil
}

// WatchEmergencyPauseSet is a free log subscription operation binding the contract event 0xc32df8d34dee7f1c985371a6d6c56a54baf4e051bd1cc6f005d4eb28afb9d23f.
//
// Solidity: event EmergencyPauseSet(address indexed pauser)
func (_Node *NodeFilterer) WatchEmergencyPauseSet(opts *bind.WatchOpts, sink chan<- *NodeEmergencyPauseSet, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "EmergencyPauseSet", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeEmergencyPauseSet)
				if err := _Node.contract.UnpackLog(event, "EmergencyPauseSet", log); err != nil {
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

// ParseEmergencyPauseSet is a log parse operation binding the contract event 0xc32df8d34dee7f1c985371a6d6c56a54baf4e051bd1cc6f005d4eb28afb9d23f.
//
// Solidity: event EmergencyPauseSet(address indexed pauser)
func (_Node *NodeFilterer) ParseEmergencyPauseSet(log types.Log) (*NodeEmergencyPauseSet, error) {
	event := new(NodeEmergencyPauseSet)
	if err := _Node.contract.UnpackLog(event, "EmergencyPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Node contract.
type NodeInitializedIterator struct {
	Event *NodeInitialized // Event containing the contract specifics and raw log

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
func (it *NodeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeInitialized)
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
		it.Event = new(NodeInitialized)
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
func (it *NodeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeInitialized represents a Initialized event raised by the Node contract.
type NodeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Node *NodeFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodeInitializedIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodeInitializedIterator{contract: _Node.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Node *NodeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodeInitialized) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeInitialized)
				if err := _Node.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Node *NodeFilterer) ParseInitialized(log types.Log) (*NodeInitialized, error) {
	event := new(NodeInitialized)
	if err := _Node.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodePunishIterator is returned from FilterPunish and is used to iterate over the raw logs and unpacked data for Punish events raised by the Node contract.
type NodePunishIterator struct {
	Event *NodePunish // Event containing the contract specifics and raw log

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
func (it *NodePunishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodePunish)
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
		it.Event = new(NodePunish)
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
func (it *NodePunishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodePunishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodePunish represents a Punish event raised by the Node contract.
type NodePunish struct {
	A     common.Address
	Typ   uint8
	To    common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPunish is a free log retrieval operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) FilterPunish(opts *bind.FilterOpts, _a []common.Address, _to []common.Address) (*NodePunishIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &NodePunishIterator{contract: _Node.contract, event: "Punish", logs: logs, sub: sub}, nil
}

// WatchPunish is a free log subscription operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) WatchPunish(opts *bind.WatchOpts, sink chan<- *NodePunish, _a []common.Address, _to []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodePunish)
				if err := _Node.contract.UnpackLog(event, "Punish", log); err != nil {
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

// ParsePunish is a log parse operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_Node *NodeFilterer) ParsePunish(log types.Log) (*NodePunish, error) {
	event := new(NodePunish)
	if err := _Node.contract.UnpackLog(event, "Punish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Node contract.
type NodeRoleAdminChangedIterator struct {
	Event *NodeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NodeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRoleAdminChanged)
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
		it.Event = new(NodeRoleAdminChanged)
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
func (it *NodeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRoleAdminChanged represents a RoleAdminChanged event raised by the Node contract.
type NodeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Node *NodeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NodeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Node.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NodeRoleAdminChangedIterator{contract: _Node.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Node *NodeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NodeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Node.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRoleAdminChanged)
				if err := _Node.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Node *NodeFilterer) ParseRoleAdminChanged(log types.Log) (*NodeRoleAdminChanged, error) {
	event := new(NodeRoleAdminChanged)
	if err := _Node.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Node contract.
type NodeRoleGrantedIterator struct {
	Event *NodeRoleGranted // Event containing the contract specifics and raw log

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
func (it *NodeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRoleGranted)
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
		it.Event = new(NodeRoleGranted)
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
func (it *NodeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRoleGranted represents a RoleGranted event raised by the Node contract.
type NodeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Node *NodeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NodeRoleGrantedIterator, error) {

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

	logs, sub, err := _Node.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NodeRoleGrantedIterator{contract: _Node.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Node *NodeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NodeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Node.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRoleGranted)
				if err := _Node.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Node *NodeFilterer) ParseRoleGranted(log types.Log) (*NodeRoleGranted, error) {
	event := new(NodeRoleGranted)
	if err := _Node.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Node contract.
type NodeRoleRevokedIterator struct {
	Event *NodeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NodeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRoleRevoked)
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
		it.Event = new(NodeRoleRevoked)
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
func (it *NodeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRoleRevoked represents a RoleRevoked event raised by the Node contract.
type NodeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Node *NodeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NodeRoleRevokedIterator, error) {

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

	logs, sub, err := _Node.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NodeRoleRevokedIterator{contract: _Node.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Node *NodeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NodeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Node.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRoleRevoked)
				if err := _Node.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Node *NodeFilterer) ParseRoleRevoked(log types.Log) (*NodeRoleRevoked, error) {
	event := new(NodeRoleRevoked)
	if err := _Node.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the Node contract.
type NodeSetIterator struct {
	Event *NodeSet // Event containing the contract specifics and raw log

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
func (it *NodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeSet)
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
		it.Event = new(NodeSet)
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
func (it *NodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeSet represents a Set event raised by the Node contract.
type NodeSet struct {
	Type uint8
	M    *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _m)
func (_Node *NodeFilterer) FilterSet(opts *bind.FilterOpts) (*NodeSetIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &NodeSetIterator{contract: _Node.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _m)
func (_Node *NodeFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *NodeSet) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeSet)
				if err := _Node.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _m)
func (_Node *NodeFilterer) ParseSet(log types.Log) (*NodeSet, error) {
	event := new(NodeSet)
	if err := _Node.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Node contract.
type NodeStakedIterator struct {
	Event *NodeStaked // Event containing the contract specifics and raw log

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
func (it *NodeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStaked)
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
		it.Event = new(NodeStaked)
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
func (it *NodeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStaked represents a Staked event raised by the Node contract.
type NodeStaked struct {
	A    common.Address
	Type uint8
	M    *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x3cf14181ae25669a913d72411736fc5c01f538fa503e963b0b2e56bcefb3edaf.
//
// Solidity: event Staked(address indexed a, uint8 _type, uint256 m)
func (_Node *NodeFilterer) FilterStaked(opts *bind.FilterOpts, a []common.Address) (*NodeStakedIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Staked", aRule)
	if err != nil {
		return nil, err
	}
	return &NodeStakedIterator{contract: _Node.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x3cf14181ae25669a913d72411736fc5c01f538fa503e963b0b2e56bcefb3edaf.
//
// Solidity: event Staked(address indexed a, uint8 _type, uint256 m)
func (_Node *NodeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *NodeStaked, a []common.Address) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Staked", aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStaked)
				if err := _Node.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x3cf14181ae25669a913d72411736fc5c01f538fa503e963b0b2e56bcefb3edaf.
//
// Solidity: event Staked(address indexed a, uint8 _type, uint256 m)
func (_Node *NodeFilterer) ParseStaked(log types.Log) (*NodeStaked, error) {
	event := new(NodeStaked)
	if err := _Node.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeTerminatedIterator is returned from FilterTerminated and is used to iterate over the raw logs and unpacked data for Terminated events raised by the Node contract.
type NodeTerminatedIterator struct {
	Event *NodeTerminated // Event containing the contract specifics and raw log

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
func (it *NodeTerminatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeTerminated)
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
		it.Event = new(NodeTerminated)
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
func (it *NodeTerminatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeTerminatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeTerminated represents a Terminated event raised by the Node contract.
type NodeTerminated struct {
	A   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTerminated is a free log retrieval operation binding the contract event 0x98cd97fc1a1cc958cbd729b1bb531d4f3ea4925470bf23ea9af386640cbd07be.
//
// Solidity: event Terminated(address indexed a)
func (_Node *NodeFilterer) FilterTerminated(opts *bind.FilterOpts, a []common.Address) (*NodeTerminatedIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Terminated", aRule)
	if err != nil {
		return nil, err
	}
	return &NodeTerminatedIterator{contract: _Node.contract, event: "Terminated", logs: logs, sub: sub}, nil
}

// WatchTerminated is a free log subscription operation binding the contract event 0x98cd97fc1a1cc958cbd729b1bb531d4f3ea4925470bf23ea9af386640cbd07be.
//
// Solidity: event Terminated(address indexed a)
func (_Node *NodeFilterer) WatchTerminated(opts *bind.WatchOpts, sink chan<- *NodeTerminated, a []common.Address) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Terminated", aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeTerminated)
				if err := _Node.contract.UnpackLog(event, "Terminated", log); err != nil {
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

// ParseTerminated is a log parse operation binding the contract event 0x98cd97fc1a1cc958cbd729b1bb531d4f3ea4925470bf23ea9af386640cbd07be.
//
// Solidity: event Terminated(address indexed a)
func (_Node *NodeFilterer) ParseTerminated(log types.Log) (*NodeTerminated, error) {
	event := new(NodeTerminated)
	if err := _Node.contract.UnpackLog(event, "Terminated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Node contract.
type NodeUpgradedIterator struct {
	Event *NodeUpgraded // Event containing the contract specifics and raw log

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
func (it *NodeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeUpgraded)
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
		it.Event = new(NodeUpgraded)
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
func (it *NodeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeUpgraded represents a Upgraded event raised by the Node contract.
type NodeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Node *NodeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NodeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NodeUpgradedIterator{contract: _Node.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Node *NodeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NodeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeUpgraded)
				if err := _Node.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Node *NodeFilterer) ParseUpgraded(log types.Log) (*NodeUpgraded, error) {
	event := new(NodeUpgraded)
	if err := _Node.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Node contract.
type NodeWithdrawnIterator struct {
	Event *NodeWithdrawn // Event containing the contract specifics and raw log

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
func (it *NodeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeWithdrawn)
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
		it.Event = new(NodeWithdrawn)
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
func (it *NodeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeWithdrawn represents a Withdrawn event raised by the Node contract.
type NodeWithdrawn struct {
	A   common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed a, uint256 m)
func (_Node *NodeFilterer) FilterWithdrawn(opts *bind.FilterOpts, a []common.Address) (*NodeWithdrawnIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "Withdrawn", aRule)
	if err != nil {
		return nil, err
	}
	return &NodeWithdrawnIterator{contract: _Node.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed a, uint256 m)
func (_Node *NodeFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *NodeWithdrawn, a []common.Address) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "Withdrawn", aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeWithdrawn)
				if err := _Node.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed a, uint256 m)
func (_Node *NodeFilterer) ParseWithdrawn(log types.Log) (*NodeWithdrawn, error) {
	event := new(NodeWithdrawn)
	if err := _Node.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
