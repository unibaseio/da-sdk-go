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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_SLOTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_SLOT_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_SLOTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_SLOT_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"check\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"current\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpochTime\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_slots\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlotSeconds\",\"inputs\":[{\"name\":\"_s\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlots\",\"inputs\":[{\"name\":\"_slots\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slotSeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slotSecondsChangedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slots\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slotsChangedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slotsWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetEpoch\",\"inputs\":[{\"name\":\"_e\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateSlotSeconds\",\"inputs\":[{\"name\":\"_s\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateSlots\",\"inputs\":[{\"name\":\"_s\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a080604052346100cc57306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100bd57506001600160401b036002600160401b031982821601610078575b60405161165e90816100d1823960805181818161091f0152610a0c0152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f8080610059565b63f92ee8a960e01b8152600490fd5b5f80fdfe6080604090808252600480361015610015575f80fd5b5f3560e01c91826301ffc9a714610f235750816309dabe3514610f0557816311ff884e14610ee957816312a02c8214610eaa578163248a9ca314610e745781632f2ff15d14610e4c57816336568abe14610e085781633a9e5c7b14610be257816348547d6914610bbb5781634f1ef2861461098357816352d1902d1461090c5781636738bc6a146108e5578163919840ad146108cd57816391d148541461087d5781639fa6a6e314610857578163a217fddf1461083d578163ad3cb1cc14610801578163ae888f9c146107e6578163c0f3f2e9146107c9578163cbccf0911461056a578163ccc5749014610530578163cd0de91414610514578163d547741f146104cc578163d7eecc7e1461020957508063e64d652a146101ec578063ef050b4f146101b5578063f9dad054146101975763ffa1ad7414610154575f80fd5b34610193575f36600319011261019357805161018f9161017382610fa1565b60058252640322e302e360dc1b60208301525191829182610ff1565b0390f35b5f80fd5b5034610193575f366003190112610193576020906003549051908152f35b5034610193576020366003190112610193576020906001600160401b036101da610f75565b165f5260078252805f20549051908152f35b5034610193575f366003190112610193576020905162278d008152f35b905034610193578160031936011261019357610223610f75565b9161022c610f8b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009384549160ff83851c1615926001600160401b0392838216801590816104c4575b60011490816104ba575b1590816104b1575b506104a15767ffffffffffffffff1982811660011789559185610482575b5060ff8854871c161561047257838116156104425767ffffffffffffffff60401b5f5491871b169067ffffffffffffffff60401b1916175f556102df6110a5565b5f8152855196602097888101914483525f8983015288825260608201908282108883111761042f57506103e5938a96937f9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd759896936103b5938c5251902087820190815261034b826110e5565b438252865f541661039d6103ab8d610370600161036786611038565b500154946110cd565b90518c810194855244602086015260c09190911b6001600160c01b03191660408501529182906048850190565b03601f198101835282610fd0565b51902090526110e5565b5f54846103c3818316611133565b1691829116175f555f526007845242875f20556103df81611318565b506113b5565b505f54168451908152a16103f557005b7fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29268ff00000000000000001981541690555160018152a1005b604190634e487b7160e01b5f525260245ffd5b855162461bcd60e51b8152602081890152600a6024820152697a65726f20736c6f747360b01b6044820152606490fd5b8551631afcd79f60e31b81528790fd5b68ffffffffffffffffff1916680100000000000000011788555f61029e565b855163f92ee8a960e01b81528790fd5b9050155f610280565b303b159150610278565b86915061026e565b8234610193578060031936011261019357610512913561050d60016104ef610f8b565b93835f525f805160206116098339815191526020525f2001546112ea565b6114f8565b005b8234610193575f36600319011261019357602091549051908152f35b8234610193575f36600319011261019357602090517f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f558152f35b823461019357602036600319011261019357610584610f75565b9161058d611271565b6001600160401b0380841693606480861015806107bc575b15610786575f548381871c16918288036105e6575b7fba1da09884224d94b905b0a3a828942e734c70ad3abf9cb4a66bda28fa7e64866020898951908152a1005b677fffffffffffffff8260411c16620f424080821161077e575b506801fffffffffffffffe83603f1c16908482046002148515171561076b57828210610763575b8910159081610758575b50156107205785548061071b5750825b600354908115918215610707575b5050156106c75750918593917fba1da09884224d94b905b0a3a828942e734c70ad3abf9cb4a66bda28fa7e648697959367ffffffffffffffff60401b836020991b169067ffffffffffffffff60401b1916175f55436003558086115f146106c057505b1690559181938580806105ba565b90506106b2565b865162461bcd60e51b8152602081880152601860248201527f736c6f7473206368616e67656420746869732065706f636800000000000000006044820152fd5b6107119250611084565b431015898061064f565b610641565b865162461bcd60e51b815260208188015260156024820152740e6d8dee8e640c6d0c2dcceca40c2c4deecca4064f605b1b6044820152fd5b905088111589610631565b829150610627565b601188634e487b7160e01b5f525260245ffd5b905089610600565b845162461bcd60e51b81526020818601526013602482015272736c6f7473206f7574206f6620626f756e647360681b6044820152fd5b50620f42408611156105a5565b8234610193575f3660031901126101935760209051620f42408152f35b8234610193575f366003190112610193576020905160648152f35b8234610193575f36600319011261019357805161018f9161082182610fa1565b60058252640352e302e360dc1b60208301525191829182610ff1565b8234610193575f36600319011261019357602090515f8152f35b8234610193575f366003190112610193576020906001600160401b035f54169051908152f35b8234610193578060031936011261019357602091610899610f8b565b90355f525f805160206116098339815191528352815f209060018060a01b03165f52825260ff815f20541690519015158152f35b34610193575f3660031901126101935761051261114b565b8234610193575f366003190112610193576020906001600160401b03600554169051908152f35b8234610193575f366003190112610193577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361097657602090517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b5163703e46dd60e11b8152fd5b905081600319360112610193578035916001600160a01b038084169190828503610193576024948535916001600160401b039081841161019357366023850112156101935783870135918211610ba9576020918551946109ec84601f19601f8501160187610fd0565b818652368a838301011161019357815f928b8693018389013786010152807f000000000000000000000000000000000000000000000000000000000000000016803014918215610b7b575b5050610b6b57610a45611271565b83516352d1902d60e01b81529080828881895afa9182915f93610b3b575b5050610a7d575050505191634c9c8ce360e01b8352820152fd5b85938591887f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc91828103610b265750843b15610b12575080546001600160a01b0319168317905551907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2825115610afc57506105129250611577565b91505034610b0657005b63b398979f60e01b8152fd5b8251634c9c8ce360e01b8152808801859052fd5b8351632a87526960e21b815280890191909152fd5b9080929350813d8311610b64575b610b538183610fd0565b810103126101935751905f80610a63565b503d610b49565b835163703e46dd60e11b81528690fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141590505f80610a37565b87604188634e487b7160e01b5f52525ffd5b8234610193575f366003190112610193575f548151911c6001600160401b03168152602090f35b90503461019357602036600319011261019357610bfd610f75565b610c05611271565b6001600160401b03918282169261025884101580610dfb575b15610db85760055492818416808603610c5e575b7fa3557cd174392277e79af681d511eef6b26d509069311d2b08d9a063e91485bb6020878951908152a1005b80158015610cab575b505050505067ffffffffffffffff1916811760055542600655807fa3557cd174392277e79af681d511eef6b26d509069311d2b08d9a063e91485bb60205f80610c32565b677fffffffffffffff8660011c168710159081610d7e575b5015610d3b57610ce292919080871115610d3457505b16600654611084565b4210610cf15780808080610c67565b606490602085519162461bcd60e51b8352820152601f60248201527f736c6f74207365636f6e6473206368616e67656420746869732065706f6368006044820152fd5b9050610cd9565b865162461bcd60e51b8152602081860152601c60248201527f736c6f74207365636f6e6473206368616e67652061626f7665203278000000006044820152606490fd5b90506801fffffffffffffffe8660011b16908282046002141715610da5578611155f610cc3565b601185634e487b7160e01b5f525260245ffd5b845162461bcd60e51b8152602081840152601a60248201527f736c6f74207365636f6e6473206f7574206f6620626f756e64730000000000006044820152606490fd5b5062278d00841115610c1e565b8234610193578060031936011261019357610e21610f8b565b90336001600160a01b03831603610e3d575061051291356114f8565b5163334bd91960e11b81529050fd5b82346101935780600319360112610193576105129135610e6f60016104ef610f8b565b611475565b823461019357602036600319011261019357602091355f525f8051602061160983398151915282526001815f2001549051908152f35b823461019357602036600319011261019357610ec4610f75565b6001610eda610ed283611038565b505492611038565b50015482519182526020820152f35b8234610193575f36600319011261019357602090516102588152f35b8234610193575f366003190112610193576020906006549051908152f35b903461019357602036600319011261019357359063ffffffff60e01b821680920361019357602091637965db0b60e01b8114908115610f64575b5015158152f35b6301ffc9a760e01b14905083610f5d565b600435906001600160401b038216820361019357565b602435906001600160a01b038216820361019357565b604081019081106001600160401b03821117610fbc57604052565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b03821117610fbc57604052565b602080825282518183018190529093925f5b82811061102457505060409293505f838284010152601f8019910116010190565b818101860151848201604001528501611003565b6002548110156110705760025f5260011b7f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01905f90565b634e487b7160e01b5f52603260045260245ffd5b9190820180921161109157565b634e487b7160e01b5f52601160045260245ffd5b60405190604082018281106001600160401b03821117610fbc576040525f6020838281520152565b9060016001600160401b038093160191821161109157565b60025468010000000000000000811015610fbc578060016111099201600255611038565b919091611120576020816001925184550151910155565b634e487b7160e01b5f525f60045260245ffd5b6001600160401b038091169081146110915760010190565b5f546001600160401b03908181169061116382611038565b50544303904382116110915760401c83161161124e5781600554168015159081611252575b5061124e577f9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75916112186020926111bd6110a5565b9043825261039d61120c6111de60016111d585611038565b500154936110cd565b6040805189810195865244602087015260c09290921b6001600160c01b031916908501529182906048850190565b519020848201526110e5565b5f5481611226818316611133565b1680916001600160401b031916175f555f52600782524260405f20555f5416604051908152a1565b5050565b6112699150825f52600760205260405f2054611084565b42105f611188565b335f9081527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f559060ff16156112cc5750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b805f525f8051602061160983398151915260205260405f20335f5260205260ff60405f205416156112cc5750565b6001600160a01b03165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f805160206116098339815191529060ff166113af575f805260205260405f20815f5260205260405f20600160ff1982541617905533905f7f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50505f90565b6001600160a01b03165f8181527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5591905f805160206116098339815191529060ff1661146e57825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b5050505f90565b90815f525f805160206116098339815191528060205260405f209160018060a01b031691825f5260205260ff60405f205416155f1461146e57825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b90815f525f805160206116098339815191528060205260405f209160018060a01b031691825f5260205260ff60405f2054165f1461146e57825f5260205260405f20815f5260205260405f2060ff19815416905533917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b905f8091602081519101845af480806115f5575b156115ab5750506040513d81523d5f602083013e60203d82010160405290565b156115d257604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b3d156115e3576040513d5f823e3d90fd5b60405163d6bda27560e01b8152600490fd5b503d15158061158b5750813b151561158b56fe02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a2646970667358221220a9ab437a9e805cbe697b0ef84b52bc47933a701758e47cc514ed492ee797d36a64736f6c63430008180033",
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

// MAXSLOTS is a free data retrieval call binding the contract method 0xc0f3f2e9.
//
// Solidity: function MAX_SLOTS() view returns(uint64)
func (_Epoch *EpochCaller) MAXSLOTS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "MAX_SLOTS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXSLOTS is a free data retrieval call binding the contract method 0xc0f3f2e9.
//
// Solidity: function MAX_SLOTS() view returns(uint64)
func (_Epoch *EpochSession) MAXSLOTS() (uint64, error) {
	return _Epoch.Contract.MAXSLOTS(&_Epoch.CallOpts)
}

// MAXSLOTS is a free data retrieval call binding the contract method 0xc0f3f2e9.
//
// Solidity: function MAX_SLOTS() view returns(uint64)
func (_Epoch *EpochCallerSession) MAXSLOTS() (uint64, error) {
	return _Epoch.Contract.MAXSLOTS(&_Epoch.CallOpts)
}

// MAXSLOTSECONDS is a free data retrieval call binding the contract method 0xe64d652a.
//
// Solidity: function MAX_SLOT_SECONDS() view returns(uint64)
func (_Epoch *EpochCaller) MAXSLOTSECONDS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "MAX_SLOT_SECONDS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXSLOTSECONDS is a free data retrieval call binding the contract method 0xe64d652a.
//
// Solidity: function MAX_SLOT_SECONDS() view returns(uint64)
func (_Epoch *EpochSession) MAXSLOTSECONDS() (uint64, error) {
	return _Epoch.Contract.MAXSLOTSECONDS(&_Epoch.CallOpts)
}

// MAXSLOTSECONDS is a free data retrieval call binding the contract method 0xe64d652a.
//
// Solidity: function MAX_SLOT_SECONDS() view returns(uint64)
func (_Epoch *EpochCallerSession) MAXSLOTSECONDS() (uint64, error) {
	return _Epoch.Contract.MAXSLOTSECONDS(&_Epoch.CallOpts)
}

// MINSLOTS is a free data retrieval call binding the contract method 0xae888f9c.
//
// Solidity: function MIN_SLOTS() view returns(uint64)
func (_Epoch *EpochCaller) MINSLOTS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "MIN_SLOTS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINSLOTS is a free data retrieval call binding the contract method 0xae888f9c.
//
// Solidity: function MIN_SLOTS() view returns(uint64)
func (_Epoch *EpochSession) MINSLOTS() (uint64, error) {
	return _Epoch.Contract.MINSLOTS(&_Epoch.CallOpts)
}

// MINSLOTS is a free data retrieval call binding the contract method 0xae888f9c.
//
// Solidity: function MIN_SLOTS() view returns(uint64)
func (_Epoch *EpochCallerSession) MINSLOTS() (uint64, error) {
	return _Epoch.Contract.MINSLOTS(&_Epoch.CallOpts)
}

// MINSLOTSECONDS is a free data retrieval call binding the contract method 0x11ff884e.
//
// Solidity: function MIN_SLOT_SECONDS() view returns(uint64)
func (_Epoch *EpochCaller) MINSLOTSECONDS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "MIN_SLOT_SECONDS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINSLOTSECONDS is a free data retrieval call binding the contract method 0x11ff884e.
//
// Solidity: function MIN_SLOT_SECONDS() view returns(uint64)
func (_Epoch *EpochSession) MINSLOTSECONDS() (uint64, error) {
	return _Epoch.Contract.MINSLOTSECONDS(&_Epoch.CallOpts)
}

// MINSLOTSECONDS is a free data retrieval call binding the contract method 0x11ff884e.
//
// Solidity: function MIN_SLOT_SECONDS() view returns(uint64)
func (_Epoch *EpochCallerSession) MINSLOTSECONDS() (uint64, error) {
	return _Epoch.Contract.MINSLOTSECONDS(&_Epoch.CallOpts)
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

// GetEpochTime is a free data retrieval call binding the contract method 0xef050b4f.
//
// Solidity: function getEpochTime(uint64 _epoch) view returns(uint256)
func (_Epoch *EpochCaller) GetEpochTime(opts *bind.CallOpts, _epoch uint64) (*big.Int, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "getEpochTime", _epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochTime is a free data retrieval call binding the contract method 0xef050b4f.
//
// Solidity: function getEpochTime(uint64 _epoch) view returns(uint256)
func (_Epoch *EpochSession) GetEpochTime(_epoch uint64) (*big.Int, error) {
	return _Epoch.Contract.GetEpochTime(&_Epoch.CallOpts, _epoch)
}

// GetEpochTime is a free data retrieval call binding the contract method 0xef050b4f.
//
// Solidity: function getEpochTime(uint64 _epoch) view returns(uint256)
func (_Epoch *EpochCallerSession) GetEpochTime(_epoch uint64) (*big.Int, error) {
	return _Epoch.Contract.GetEpochTime(&_Epoch.CallOpts, _epoch)
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

// SlotSeconds is a free data retrieval call binding the contract method 0x6738bc6a.
//
// Solidity: function slotSeconds() view returns(uint64)
func (_Epoch *EpochCaller) SlotSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "slotSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SlotSeconds is a free data retrieval call binding the contract method 0x6738bc6a.
//
// Solidity: function slotSeconds() view returns(uint64)
func (_Epoch *EpochSession) SlotSeconds() (uint64, error) {
	return _Epoch.Contract.SlotSeconds(&_Epoch.CallOpts)
}

// SlotSeconds is a free data retrieval call binding the contract method 0x6738bc6a.
//
// Solidity: function slotSeconds() view returns(uint64)
func (_Epoch *EpochCallerSession) SlotSeconds() (uint64, error) {
	return _Epoch.Contract.SlotSeconds(&_Epoch.CallOpts)
}

// SlotSecondsChangedAt is a free data retrieval call binding the contract method 0x09dabe35.
//
// Solidity: function slotSecondsChangedAt() view returns(uint256)
func (_Epoch *EpochCaller) SlotSecondsChangedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "slotSecondsChangedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotSecondsChangedAt is a free data retrieval call binding the contract method 0x09dabe35.
//
// Solidity: function slotSecondsChangedAt() view returns(uint256)
func (_Epoch *EpochSession) SlotSecondsChangedAt() (*big.Int, error) {
	return _Epoch.Contract.SlotSecondsChangedAt(&_Epoch.CallOpts)
}

// SlotSecondsChangedAt is a free data retrieval call binding the contract method 0x09dabe35.
//
// Solidity: function slotSecondsChangedAt() view returns(uint256)
func (_Epoch *EpochCallerSession) SlotSecondsChangedAt() (*big.Int, error) {
	return _Epoch.Contract.SlotSecondsChangedAt(&_Epoch.CallOpts)
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

// SlotsChangedAt is a free data retrieval call binding the contract method 0xf9dad054.
//
// Solidity: function slotsChangedAt() view returns(uint256)
func (_Epoch *EpochCaller) SlotsChangedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "slotsChangedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotsChangedAt is a free data retrieval call binding the contract method 0xf9dad054.
//
// Solidity: function slotsChangedAt() view returns(uint256)
func (_Epoch *EpochSession) SlotsChangedAt() (*big.Int, error) {
	return _Epoch.Contract.SlotsChangedAt(&_Epoch.CallOpts)
}

// SlotsChangedAt is a free data retrieval call binding the contract method 0xf9dad054.
//
// Solidity: function slotsChangedAt() view returns(uint256)
func (_Epoch *EpochCallerSession) SlotsChangedAt() (*big.Int, error) {
	return _Epoch.Contract.SlotsChangedAt(&_Epoch.CallOpts)
}

// SlotsWait is a free data retrieval call binding the contract method 0xcd0de914.
//
// Solidity: function slotsWait() view returns(uint256)
func (_Epoch *EpochCaller) SlotsWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "slotsWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotsWait is a free data retrieval call binding the contract method 0xcd0de914.
//
// Solidity: function slotsWait() view returns(uint256)
func (_Epoch *EpochSession) SlotsWait() (*big.Int, error) {
	return _Epoch.Contract.SlotsWait(&_Epoch.CallOpts)
}

// SlotsWait is a free data retrieval call binding the contract method 0xcd0de914.
//
// Solidity: function slotsWait() view returns(uint256)
func (_Epoch *EpochCallerSession) SlotsWait() (*big.Int, error) {
	return _Epoch.Contract.SlotsWait(&_Epoch.CallOpts)
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

// SetSlotSeconds is a paid mutator transaction binding the contract method 0x3a9e5c7b.
//
// Solidity: function setSlotSeconds(uint64 _s) returns()
func (_Epoch *EpochTransactor) SetSlotSeconds(opts *bind.TransactOpts, _s uint64) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "setSlotSeconds", _s)
}

// SetSlotSeconds is a paid mutator transaction binding the contract method 0x3a9e5c7b.
//
// Solidity: function setSlotSeconds(uint64 _s) returns()
func (_Epoch *EpochSession) SetSlotSeconds(_s uint64) (*types.Transaction, error) {
	return _Epoch.Contract.SetSlotSeconds(&_Epoch.TransactOpts, _s)
}

// SetSlotSeconds is a paid mutator transaction binding the contract method 0x3a9e5c7b.
//
// Solidity: function setSlotSeconds(uint64 _s) returns()
func (_Epoch *EpochTransactorSession) SetSlotSeconds(_s uint64) (*types.Transaction, error) {
	return _Epoch.Contract.SetSlotSeconds(&_Epoch.TransactOpts, _s)
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

// EpochUpdateSlotSecondsIterator is returned from FilterUpdateSlotSeconds and is used to iterate over the raw logs and unpacked data for UpdateSlotSeconds events raised by the Epoch contract.
type EpochUpdateSlotSecondsIterator struct {
	Event *EpochUpdateSlotSeconds // Event containing the contract specifics and raw log

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
func (it *EpochUpdateSlotSecondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochUpdateSlotSeconds)
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
		it.Event = new(EpochUpdateSlotSeconds)
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
func (it *EpochUpdateSlotSecondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochUpdateSlotSecondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochUpdateSlotSeconds represents a UpdateSlotSeconds event raised by the Epoch contract.
type EpochUpdateSlotSeconds struct {
	S   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateSlotSeconds is a free log retrieval operation binding the contract event 0xa3557cd174392277e79af681d511eef6b26d509069311d2b08d9a063e91485bb.
//
// Solidity: event UpdateSlotSeconds(uint64 _s)
func (_Epoch *EpochFilterer) FilterUpdateSlotSeconds(opts *bind.FilterOpts) (*EpochUpdateSlotSecondsIterator, error) {

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "UpdateSlotSeconds")
	if err != nil {
		return nil, err
	}
	return &EpochUpdateSlotSecondsIterator{contract: _Epoch.contract, event: "UpdateSlotSeconds", logs: logs, sub: sub}, nil
}

// WatchUpdateSlotSeconds is a free log subscription operation binding the contract event 0xa3557cd174392277e79af681d511eef6b26d509069311d2b08d9a063e91485bb.
//
// Solidity: event UpdateSlotSeconds(uint64 _s)
func (_Epoch *EpochFilterer) WatchUpdateSlotSeconds(opts *bind.WatchOpts, sink chan<- *EpochUpdateSlotSeconds) (event.Subscription, error) {

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "UpdateSlotSeconds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochUpdateSlotSeconds)
				if err := _Epoch.contract.UnpackLog(event, "UpdateSlotSeconds", log); err != nil {
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

// ParseUpdateSlotSeconds is a log parse operation binding the contract event 0xa3557cd174392277e79af681d511eef6b26d509069311d2b08d9a063e91485bb.
//
// Solidity: event UpdateSlotSeconds(uint64 _s)
func (_Epoch *EpochFilterer) ParseUpdateSlotSeconds(log types.Log) (*EpochUpdateSlotSeconds, error) {
	event := new(EpochUpdateSlotSeconds)
	if err := _Epoch.contract.UnpackLog(event, "UpdateSlotSeconds", log); err != nil {
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
