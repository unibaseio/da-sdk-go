// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorreward

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

// ValidatorRewardMetaData contains all meta data concerning the ValidatorReward contract.
var ValidatorRewardMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addValidator\",\"inputs\":[{\"name\":\"_v\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"attest\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"attestCount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"available\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"distribute\",\"inputs\":[{\"name\":\"_vals\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fund\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastAttest\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pending\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeValidator\",\"inputs\":[{\"name\":\"_v\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalPending\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Attested\",\"inputs\":[{\"name\":\"v\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"epoch\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"v\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Distributed\",\"inputs\":[{\"name\":\"v\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Funded\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorAdded\",\"inputs\":[{\"name\":\"v\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemoved\",\"inputs\":[{\"name\":\"v\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a080604052346100cc57306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100bd57506001600160401b036002600160401b031982821601610078575b6040516116cd90816100d182396080518181816104cb01526105b00152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f8080610059565b63f92ee8a960e01b8152600490fd5b5f80fdfe6080604090808252600480361015610015575f80fd5b5f3560e01c91826301ffc9a714611033575081630f43a67714611015578163248a9ca314610fdf5781632929abe614610e035781632aeaa0a014610d165781632f2ff15d14610cee57816335aa2e4414610cae57816336568abe14610c6a5781633f90916a14610c4c57816340a141ff14610b18578163485cc955146109ac57816348a0d754146109895781634d238c8e146108505781634e71d92d146107625781634f1ef2861461052f57816352d1902d146104b85781635eebea201461048157816391d1485414610431578163a217fddf14610417578163ad3cb1cc146103db578163ca1d209d146102f2578163ccc57490146102b8578163d547741f14610270578163eeea7ca81461023b57508063facd743b146101ff578063fc0c546a146101d8578063fd122be6146101975763ffa1ad7414610154575f80fd5b34610193575f36600319011261019357805161018f916101738261112b565b60058252640312e302e360dc1b6020830152519182918261117d565b0390f35b5f80fd5b5034610193576020366003190112610193576020906001600160a01b036101bc6110cc565b165f526003825267ffffffffffffffff815f2054169051908152f35b5034610193575f366003190112610193575f5490516001600160a01b039091168152602090f35b5034610193576020366003190112610193576020906001600160a01b036102246110cc565b165f526001825260ff815f20541690519015158152f35b8234610193576020366003190112610193576020916001600160a01b036102606110cc565b165f528252805f20549051908152f35b82346101935780600319360112610193576102b691356102b160016102936110b6565b93835f525f805160206116788339815191526020525f20015461135f565b611567565b005b8234610193575f36600319011261019357602090517f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f558152f35b82346101935760203660031901126101935781359182156103b35760018060a01b035f54168251916323b872dd60e01b5f52338152306024528460445260205f60648180865af160015f5114811615610394575b8385525f6060521561038157505050519081527f5af8184bef8e4b45eb9f6ed7734d04da38ced226495548f46e0c8ff8d7d9a52460203392a2005b635274afe760e01b835282015260249150fd5b60018115166103aa57823b15153d151616610346565b833d5f823e3d90fd5b606491519062461bcd60e51b82526020818301526024820152637a65726f60e01b6044820152fd5b8234610193575f36600319011261019357805161018f916103fb8261112b565b60058252640352e302e360dc1b6020830152519182918261117d565b8234610193575f36600319011261019357602090515f8152f35b823461019357806003193601126101935760209161044d6110b6565b90355f525f805160206116788339815191528352815f209060018060a01b03165f52825260ff815f20541690519015158152f35b8234610193576020366003190112610193576020906001600160a01b036104a66110cc565b165f5260058252805f20549051908152f35b8234610193575f366003190112610193577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361052257602090517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b5163703e46dd60e11b8152fd5b905081600319360112610193576105446110cc565b602492833567ffffffffffffffff92838211610193573660238301121561019357818501359384116107505760209383519261058986601f19601f850116018561115b565b8184523688838301011161019357815f928988930183870137840101526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116308114908115610722575b50610712576105e96112e6565b83516352d1902d60e01b81529082169480828881895afa9182915f936106e2575b5050610624575050505191634c9c8ce360e01b8352820152fd5b85938591887f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc918281036106cd5750843b156106b9575080546001600160a01b0319168317905551907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28251156106a357506102b692506115e6565b915050346106ad57005b63b398979f60e01b8152fd5b8251634c9c8ce360e01b8152808801859052fd5b8351632a87526960e21b815280890191909152fd5b9080929350813d831161070b575b6106fa818361115b565b810103126101935751905f8061060a565b503d6106f0565b835163703e46dd60e11b81528690fd5b9050817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc541614155f6105dc565b85604186634e487b7160e01b5f52525ffd5b8234610193575f36600319011261019357335f526005602052805f205491821561082457335f5260056020525f8281205561079f83600654611245565b60065560018060a01b035f541682519163a9059cbb60e01b5f523381528460245260205f60448180865af160015f511481161561080e575b8385521561038157505050519081527fd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a60203392a2005b60018115166103aa57823b15153d1516166107d7565b6020606492519162461bcd60e51b835282015260076024820152666e6f7468696e6760c81b6044820152fd5b9050346101935760203660031901126101935761086b6110cc565b906108746112e6565b6001600160a01b03821692831561095a57835f52600160205260ff815f20541661092e57835f5260016020525f20600160ff19825416179055600254906801000000000000000082101561091b5750906108d78260016108f594016002556110e2565b90919060018060a01b038084549260031b9316831b921b1916179055565b7fe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec38849875f80a2005b604190634e487b7160e01b5f525260245ffd5b906020606492519162461bcd60e51b8352820152600660248201526565786973747360d01b6044820152fd5b906020606492519162461bcd60e51b835282015260096024820152683d32b9379030b2323960b91b6044820152fd5b8234610193575f366003190112610193576020906109a5611252565b9051908152f35b9050346101935781600319360112610193576109c66110cc565b906109cf6110b6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009283549260ff84871c16159367ffffffffffffffff811680159081610b10575b6001149081610b06575b159081610afd575b50610aee5767ffffffffffffffff198116600117865584610acf575b5060ff8554871c1615610ac157505f80546001600160a01b0319166001600160a01b03909216919091179055610a7d90610a778161138d565b50611424565b50610a8457005b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b8551631afcd79f60e31b8152fd5b68ffffffffffffffffff1916680100000000000000011785555f610a3e565b50855163f92ee8a960e01b8152fd5b9050155f610a22565b303b159150610a1a565b869150610a10565b823461019357602036600319011261019357610b326110cc565b91610b3b6112e6565b60018060a01b0380931692835f526001926001602052610b6060ff825f205416611209565b845f5260016020525f2060ff198154169055600280545f5b818110610ba7575b867fe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f15f80a2005b8387610bb2836110e2565b929054600393841b1c1614610bc957508501610b78565b955091935f19929091838101908111610c3957906108d785610bed610bfa946110e2565b9054908a1b1c16916110e2565b8354908115610c2657500192610c0f846110e2565b81939154921b1b1916905555818080808080610b80565b603190634e487b7160e01b5f525260245ffd5b601183634e487b7160e01b5f525260245ffd5b8234610193575f366003190112610193576020906006549051908152f35b8234610193578060031936011261019357610c836110b6565b90336001600160a01b03831603610c9f57506102b69135611567565b5163334bd91960e11b81529050fd5b90503461019357602036600319011261019357359060025482101561019357610cd86020926110e2565b905491519160018060a01b039160031b1c168152f35b82346101935780600319360112610193576102b69135610d1160016102936110b6565b6114e4565b90503461019357602091826003193601126101935781359267ffffffffffffffff80851680950361019357335f5260018252610d5760ff845f205416611209565b335f5260038252825f205416841115610dd4578290335f5260038152825f208567ffffffffffffffff19825416179055525f2080549160018301809311610dc1575055337f8ad601fbe08199f30553631f9519230e5fe1a3c6da337f967c91d2838f9986e15f80a3005b601190634e487b7160e01b5f525260245ffd5b905162461bcd60e51b815291820152600b60248201526a0e6e8c2d8ca40cae0dec6d60ab1b6044820152606490fd5b823461019357806003193601126101935767ffffffffffffffff90823582811161019357610e349036908501611085565b909260243590811161019357610e4d9036908601611085565b9094610e576112e6565b818303610fb6575f805b838110610f965750610e71611252565b10610f6457505f5b828110610e8257005b6001600160a01b0380610e968386896111c4565b610e9f906111f5565b165f5260209060018252855f205460ff16610eb990611209565b610ec483858a6111c4565b3581610ed185888b6111c4565b610eda906111f5565b165f5260058352865f2090815490610ef1916111d4565b9055610efe83858a6111c4565b35600690815490610f0e916111d4565b9055610f1b8386896111c4565b610f24906111f5565b90610f3084868b6111c4565b3591875192835216917fb649c98f58055c520df0dcb5709eff2e931217ff2fb1e21376130d31bbb1c0af91a2600101610e79565b606490602085519162461bcd60e51b8352820152600c60248201526b195e18d959591cc81c1bdbdb60a21b6044820152fd5b90610faf600191610fa884878c6111c4565b35906111d4565b9101610e61565b606490602085519162461bcd60e51b835282015260036024820152623632b760e91b6044820152fd5b823461019357602036600319011261019357602091355f525f8051602061167883398151915282526001815f2001549051908152f35b8234610193575f366003190112610193576020906002549051908152f35b903461019357602036600319011261019357359063ffffffff60e01b821680920361019357602091637965db0b60e01b8114908115611074575b5015158152f35b6301ffc9a760e01b1490508361106d565b9181601f840112156101935782359167ffffffffffffffff8311610193576020808501948460051b01011161019357565b602435906001600160a01b038216820361019357565b600435906001600160a01b038216820361019357565b6002548110156111175760025f527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01905f90565b634e487b7160e01b5f52603260045260245ffd5b6040810190811067ffffffffffffffff82111761114757604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761114757604052565b602080825282518183018190529093925f5b8281106111b057505060409293505f838284010152601f8019910116010190565b81810186015184820160400152850161118f565b91908110156111175760051b0190565b919082018092116111e157565b634e487b7160e01b5f52601160045260245ffd5b356001600160a01b03811681036101935790565b1561121057565b60405162461bcd60e51b815260206004820152600d60248201526c3737ba103b30b634b230ba37b960991b6044820152606490fd5b919082039182116111e157565b5f546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa9081156112db575f916112a9575b50600654808211156112a3576112a091611245565b90565b50505f90565b90506020813d6020116112d3575b816112c46020938361115b565b8101031261019357515f61128b565b3d91506112b7565b6040513d5f823e3d90fd5b335f9081527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f559060ff16156113415750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b805f525f8051602061167883398151915260205260405f20335f5260205260ff60405f205416156113415750565b6001600160a01b03165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f805160206116788339815191529060ff166112a3575f805260205260405f20815f5260205260405f20600160ff1982541617905533905f7f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b6001600160a01b03165f8181527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee80360205260409020547f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5591905f805160206116788339815191529060ff166114dd57825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b5050505f90565b90815f525f805160206116788339815191528060205260405f209160018060a01b031691825f5260205260ff60405f205416155f146114dd57825f5260205260405f20815f5260205260405f20600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b90815f525f805160206116788339815191528060205260405f209160018060a01b031691825f5260205260ff60405f2054165f146114dd57825f5260205260405f20815f5260205260405f2060ff19815416905533917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b905f8091602081519101845af48080611664575b1561161a5750506040513d81523d5f602083013e60203d82010160405290565b1561164157604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b3d15611652576040513d5f823e3d90fd5b60405163d6bda27560e01b8152600490fd5b503d1515806115fa5750813b15156115fa56fe02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a2646970667358221220a5074ad5e5696b2e03726ae4a4f0029a1f2ce3d6c1ecf90a61a4cde7ca0484c064736f6c63430008180033",
}

// ValidatorRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorRewardMetaData.ABI instead.
var ValidatorRewardABI = ValidatorRewardMetaData.ABI

// ValidatorRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorRewardMetaData.Bin instead.
var ValidatorRewardBin = ValidatorRewardMetaData.Bin

// DeployValidatorReward deploys a new Ethereum contract, binding an instance of ValidatorReward to it.
func DeployValidatorReward(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorReward, error) {
	parsed, err := ValidatorRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorRewardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorReward{ValidatorRewardCaller: ValidatorRewardCaller{contract: contract}, ValidatorRewardTransactor: ValidatorRewardTransactor{contract: contract}, ValidatorRewardFilterer: ValidatorRewardFilterer{contract: contract}}, nil
}

// ValidatorReward is an auto generated Go binding around an Ethereum contract.
type ValidatorReward struct {
	ValidatorRewardCaller     // Read-only binding to the contract
	ValidatorRewardTransactor // Write-only binding to the contract
	ValidatorRewardFilterer   // Log filterer for contract events
}

// ValidatorRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorRewardSession struct {
	Contract     *ValidatorReward  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorRewardCallerSession struct {
	Contract *ValidatorRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ValidatorRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorRewardTransactorSession struct {
	Contract     *ValidatorRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ValidatorRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRewardRaw struct {
	Contract *ValidatorReward // Generic contract binding to access the raw methods on
}

// ValidatorRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorRewardCallerRaw struct {
	Contract *ValidatorRewardCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorRewardTransactorRaw struct {
	Contract *ValidatorRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorReward creates a new instance of ValidatorReward, bound to a specific deployed contract.
func NewValidatorReward(address common.Address, backend bind.ContractBackend) (*ValidatorReward, error) {
	contract, err := bindValidatorReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorReward{ValidatorRewardCaller: ValidatorRewardCaller{contract: contract}, ValidatorRewardTransactor: ValidatorRewardTransactor{contract: contract}, ValidatorRewardFilterer: ValidatorRewardFilterer{contract: contract}}, nil
}

// NewValidatorRewardCaller creates a new read-only instance of ValidatorReward, bound to a specific deployed contract.
func NewValidatorRewardCaller(address common.Address, caller bind.ContractCaller) (*ValidatorRewardCaller, error) {
	contract, err := bindValidatorReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardCaller{contract: contract}, nil
}

// NewValidatorRewardTransactor creates a new write-only instance of ValidatorReward, bound to a specific deployed contract.
func NewValidatorRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorRewardTransactor, error) {
	contract, err := bindValidatorReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardTransactor{contract: contract}, nil
}

// NewValidatorRewardFilterer creates a new log filterer instance of ValidatorReward, bound to a specific deployed contract.
func NewValidatorRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorRewardFilterer, error) {
	contract, err := bindValidatorReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardFilterer{contract: contract}, nil
}

// bindValidatorReward binds a generic wrapper to an already deployed contract.
func bindValidatorReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorRewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorReward *ValidatorRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorReward.Contract.ValidatorRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorReward *ValidatorRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorReward.Contract.ValidatorRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorReward *ValidatorRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorReward.Contract.ValidatorRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorReward *ValidatorRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorReward *ValidatorRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorReward *ValidatorRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorReward.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ValidatorReward *ValidatorRewardCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ValidatorReward *ValidatorRewardSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ValidatorReward.Contract.DEFAULTADMINROLE(&_ValidatorReward.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ValidatorReward *ValidatorRewardCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ValidatorReward.Contract.DEFAULTADMINROLE(&_ValidatorReward.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_ValidatorReward *ValidatorRewardCaller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_ValidatorReward *ValidatorRewardSession) GOVERNORROLE() ([32]byte, error) {
	return _ValidatorReward.Contract.GOVERNORROLE(&_ValidatorReward.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_ValidatorReward *ValidatorRewardCallerSession) GOVERNORROLE() ([32]byte, error) {
	return _ValidatorReward.Contract.GOVERNORROLE(&_ValidatorReward.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ValidatorReward *ValidatorRewardCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ValidatorReward *ValidatorRewardSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ValidatorReward.Contract.UPGRADEINTERFACEVERSION(&_ValidatorReward.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ValidatorReward *ValidatorRewardCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ValidatorReward.Contract.UPGRADEINTERFACEVERSION(&_ValidatorReward.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ValidatorReward *ValidatorRewardCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ValidatorReward *ValidatorRewardSession) VERSION() (string, error) {
	return _ValidatorReward.Contract.VERSION(&_ValidatorReward.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ValidatorReward *ValidatorRewardCallerSession) VERSION() (string, error) {
	return _ValidatorReward.Contract.VERSION(&_ValidatorReward.CallOpts)
}

// AttestCount is a free data retrieval call binding the contract method 0xeeea7ca8.
//
// Solidity: function attestCount(address ) view returns(uint256)
func (_ValidatorReward *ValidatorRewardCaller) AttestCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "attestCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestCount is a free data retrieval call binding the contract method 0xeeea7ca8.
//
// Solidity: function attestCount(address ) view returns(uint256)
func (_ValidatorReward *ValidatorRewardSession) AttestCount(arg0 common.Address) (*big.Int, error) {
	return _ValidatorReward.Contract.AttestCount(&_ValidatorReward.CallOpts, arg0)
}

// AttestCount is a free data retrieval call binding the contract method 0xeeea7ca8.
//
// Solidity: function attestCount(address ) view returns(uint256)
func (_ValidatorReward *ValidatorRewardCallerSession) AttestCount(arg0 common.Address) (*big.Int, error) {
	return _ValidatorReward.Contract.AttestCount(&_ValidatorReward.CallOpts, arg0)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_ValidatorReward *ValidatorRewardCaller) Available(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "available")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_ValidatorReward *ValidatorRewardSession) Available() (*big.Int, error) {
	return _ValidatorReward.Contract.Available(&_ValidatorReward.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_ValidatorReward *ValidatorRewardCallerSession) Available() (*big.Int, error) {
	return _ValidatorReward.Contract.Available(&_ValidatorReward.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ValidatorReward *ValidatorRewardCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ValidatorReward *ValidatorRewardSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ValidatorReward.Contract.GetRoleAdmin(&_ValidatorReward.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ValidatorReward *ValidatorRewardCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ValidatorReward.Contract.GetRoleAdmin(&_ValidatorReward.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ValidatorReward *ValidatorRewardCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ValidatorReward *ValidatorRewardSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ValidatorReward.Contract.HasRole(&_ValidatorReward.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ValidatorReward *ValidatorRewardCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ValidatorReward.Contract.HasRole(&_ValidatorReward.CallOpts, role, account)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_ValidatorReward *ValidatorRewardCaller) IsValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "isValidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_ValidatorReward *ValidatorRewardSession) IsValidator(arg0 common.Address) (bool, error) {
	return _ValidatorReward.Contract.IsValidator(&_ValidatorReward.CallOpts, arg0)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_ValidatorReward *ValidatorRewardCallerSession) IsValidator(arg0 common.Address) (bool, error) {
	return _ValidatorReward.Contract.IsValidator(&_ValidatorReward.CallOpts, arg0)
}

// LastAttest is a free data retrieval call binding the contract method 0xfd122be6.
//
// Solidity: function lastAttest(address ) view returns(uint64)
func (_ValidatorReward *ValidatorRewardCaller) LastAttest(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "lastAttest", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAttest is a free data retrieval call binding the contract method 0xfd122be6.
//
// Solidity: function lastAttest(address ) view returns(uint64)
func (_ValidatorReward *ValidatorRewardSession) LastAttest(arg0 common.Address) (uint64, error) {
	return _ValidatorReward.Contract.LastAttest(&_ValidatorReward.CallOpts, arg0)
}

// LastAttest is a free data retrieval call binding the contract method 0xfd122be6.
//
// Solidity: function lastAttest(address ) view returns(uint64)
func (_ValidatorReward *ValidatorRewardCallerSession) LastAttest(arg0 common.Address) (uint64, error) {
	return _ValidatorReward.Contract.LastAttest(&_ValidatorReward.CallOpts, arg0)
}

// Pending is a free data retrieval call binding the contract method 0x5eebea20.
//
// Solidity: function pending(address ) view returns(uint256)
func (_ValidatorReward *ValidatorRewardCaller) Pending(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "pending", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pending is a free data retrieval call binding the contract method 0x5eebea20.
//
// Solidity: function pending(address ) view returns(uint256)
func (_ValidatorReward *ValidatorRewardSession) Pending(arg0 common.Address) (*big.Int, error) {
	return _ValidatorReward.Contract.Pending(&_ValidatorReward.CallOpts, arg0)
}

// Pending is a free data retrieval call binding the contract method 0x5eebea20.
//
// Solidity: function pending(address ) view returns(uint256)
func (_ValidatorReward *ValidatorRewardCallerSession) Pending(arg0 common.Address) (*big.Int, error) {
	return _ValidatorReward.Contract.Pending(&_ValidatorReward.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ValidatorReward *ValidatorRewardCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ValidatorReward *ValidatorRewardSession) ProxiableUUID() ([32]byte, error) {
	return _ValidatorReward.Contract.ProxiableUUID(&_ValidatorReward.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ValidatorReward *ValidatorRewardCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ValidatorReward.Contract.ProxiableUUID(&_ValidatorReward.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ValidatorReward *ValidatorRewardCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ValidatorReward *ValidatorRewardSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ValidatorReward.Contract.SupportsInterface(&_ValidatorReward.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ValidatorReward *ValidatorRewardCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ValidatorReward.Contract.SupportsInterface(&_ValidatorReward.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ValidatorReward *ValidatorRewardCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ValidatorReward *ValidatorRewardSession) Token() (common.Address, error) {
	return _ValidatorReward.Contract.Token(&_ValidatorReward.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ValidatorReward *ValidatorRewardCallerSession) Token() (common.Address, error) {
	return _ValidatorReward.Contract.Token(&_ValidatorReward.CallOpts)
}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_ValidatorReward *ValidatorRewardCaller) TotalPending(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "totalPending")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_ValidatorReward *ValidatorRewardSession) TotalPending() (*big.Int, error) {
	return _ValidatorReward.Contract.TotalPending(&_ValidatorReward.CallOpts)
}

// TotalPending is a free data retrieval call binding the contract method 0x3f90916a.
//
// Solidity: function totalPending() view returns(uint256)
func (_ValidatorReward *ValidatorRewardCallerSession) TotalPending() (*big.Int, error) {
	return _ValidatorReward.Contract.TotalPending(&_ValidatorReward.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorReward *ValidatorRewardCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorReward *ValidatorRewardSession) ValidatorCount() (*big.Int, error) {
	return _ValidatorReward.Contract.ValidatorCount(&_ValidatorReward.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorReward *ValidatorRewardCallerSession) ValidatorCount() (*big.Int, error) {
	return _ValidatorReward.Contract.ValidatorCount(&_ValidatorReward.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_ValidatorReward *ValidatorRewardCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ValidatorReward.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_ValidatorReward *ValidatorRewardSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorReward.Contract.Validators(&_ValidatorReward.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_ValidatorReward *ValidatorRewardCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorReward.Contract.Validators(&_ValidatorReward.CallOpts, arg0)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _v) returns()
func (_ValidatorReward *ValidatorRewardTransactor) AddValidator(opts *bind.TransactOpts, _v common.Address) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "addValidator", _v)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _v) returns()
func (_ValidatorReward *ValidatorRewardSession) AddValidator(_v common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.AddValidator(&_ValidatorReward.TransactOpts, _v)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _v) returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) AddValidator(_v common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.AddValidator(&_ValidatorReward.TransactOpts, _v)
}

// Attest is a paid mutator transaction binding the contract method 0x2aeaa0a0.
//
// Solidity: function attest(uint64 _epoch) returns()
func (_ValidatorReward *ValidatorRewardTransactor) Attest(opts *bind.TransactOpts, _epoch uint64) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "attest", _epoch)
}

// Attest is a paid mutator transaction binding the contract method 0x2aeaa0a0.
//
// Solidity: function attest(uint64 _epoch) returns()
func (_ValidatorReward *ValidatorRewardSession) Attest(_epoch uint64) (*types.Transaction, error) {
	return _ValidatorReward.Contract.Attest(&_ValidatorReward.TransactOpts, _epoch)
}

// Attest is a paid mutator transaction binding the contract method 0x2aeaa0a0.
//
// Solidity: function attest(uint64 _epoch) returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) Attest(_epoch uint64) (*types.Transaction, error) {
	return _ValidatorReward.Contract.Attest(&_ValidatorReward.TransactOpts, _epoch)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ValidatorReward *ValidatorRewardTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ValidatorReward *ValidatorRewardSession) Claim() (*types.Transaction, error) {
	return _ValidatorReward.Contract.Claim(&_ValidatorReward.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) Claim() (*types.Transaction, error) {
	return _ValidatorReward.Contract.Claim(&_ValidatorReward.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0x2929abe6.
//
// Solidity: function distribute(address[] _vals, uint256[] _amounts) returns()
func (_ValidatorReward *ValidatorRewardTransactor) Distribute(opts *bind.TransactOpts, _vals []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "distribute", _vals, _amounts)
}

// Distribute is a paid mutator transaction binding the contract method 0x2929abe6.
//
// Solidity: function distribute(address[] _vals, uint256[] _amounts) returns()
func (_ValidatorReward *ValidatorRewardSession) Distribute(_vals []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _ValidatorReward.Contract.Distribute(&_ValidatorReward.TransactOpts, _vals, _amounts)
}

// Distribute is a paid mutator transaction binding the contract method 0x2929abe6.
//
// Solidity: function distribute(address[] _vals, uint256[] _amounts) returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) Distribute(_vals []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _ValidatorReward.Contract.Distribute(&_ValidatorReward.TransactOpts, _vals, _amounts)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) returns()
func (_ValidatorReward *ValidatorRewardTransactor) Fund(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "fund", _amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) returns()
func (_ValidatorReward *ValidatorRewardSession) Fund(_amount *big.Int) (*types.Transaction, error) {
	return _ValidatorReward.Contract.Fund(&_ValidatorReward.TransactOpts, _amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) Fund(_amount *big.Int) (*types.Transaction, error) {
	return _ValidatorReward.Contract.Fund(&_ValidatorReward.TransactOpts, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ValidatorReward *ValidatorRewardTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ValidatorReward *ValidatorRewardSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.GrantRole(&_ValidatorReward.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.GrantRole(&_ValidatorReward.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token, address initialOwner) returns()
func (_ValidatorReward *ValidatorRewardTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "initialize", _token, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token, address initialOwner) returns()
func (_ValidatorReward *ValidatorRewardSession) Initialize(_token common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.Initialize(&_ValidatorReward.TransactOpts, _token, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token, address initialOwner) returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) Initialize(_token common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.Initialize(&_ValidatorReward.TransactOpts, _token, initialOwner)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _v) returns()
func (_ValidatorReward *ValidatorRewardTransactor) RemoveValidator(opts *bind.TransactOpts, _v common.Address) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "removeValidator", _v)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _v) returns()
func (_ValidatorReward *ValidatorRewardSession) RemoveValidator(_v common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.RemoveValidator(&_ValidatorReward.TransactOpts, _v)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _v) returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) RemoveValidator(_v common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.RemoveValidator(&_ValidatorReward.TransactOpts, _v)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ValidatorReward *ValidatorRewardTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ValidatorReward *ValidatorRewardSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.RenounceRole(&_ValidatorReward.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.RenounceRole(&_ValidatorReward.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ValidatorReward *ValidatorRewardTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ValidatorReward *ValidatorRewardSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.RevokeRole(&_ValidatorReward.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorReward.Contract.RevokeRole(&_ValidatorReward.TransactOpts, role, account)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ValidatorReward *ValidatorRewardTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ValidatorReward.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ValidatorReward *ValidatorRewardSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ValidatorReward.Contract.UpgradeToAndCall(&_ValidatorReward.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ValidatorReward *ValidatorRewardTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ValidatorReward.Contract.UpgradeToAndCall(&_ValidatorReward.TransactOpts, newImplementation, data)
}

// ValidatorRewardAttestedIterator is returned from FilterAttested and is used to iterate over the raw logs and unpacked data for Attested events raised by the ValidatorReward contract.
type ValidatorRewardAttestedIterator struct {
	Event *ValidatorRewardAttested // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardAttestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardAttested)
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
		it.Event = new(ValidatorRewardAttested)
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
func (it *ValidatorRewardAttestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardAttestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardAttested represents a Attested event raised by the ValidatorReward contract.
type ValidatorRewardAttested struct {
	V     common.Address
	Epoch uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAttested is a free log retrieval operation binding the contract event 0x8ad601fbe08199f30553631f9519230e5fe1a3c6da337f967c91d2838f9986e1.
//
// Solidity: event Attested(address indexed v, uint64 indexed epoch)
func (_ValidatorReward *ValidatorRewardFilterer) FilterAttested(opts *bind.FilterOpts, v []common.Address, epoch []uint64) (*ValidatorRewardAttestedIterator, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "Attested", vRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardAttestedIterator{contract: _ValidatorReward.contract, event: "Attested", logs: logs, sub: sub}, nil
}

// WatchAttested is a free log subscription operation binding the contract event 0x8ad601fbe08199f30553631f9519230e5fe1a3c6da337f967c91d2838f9986e1.
//
// Solidity: event Attested(address indexed v, uint64 indexed epoch)
func (_ValidatorReward *ValidatorRewardFilterer) WatchAttested(opts *bind.WatchOpts, sink chan<- *ValidatorRewardAttested, v []common.Address, epoch []uint64) (event.Subscription, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "Attested", vRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardAttested)
				if err := _ValidatorReward.contract.UnpackLog(event, "Attested", log); err != nil {
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

// ParseAttested is a log parse operation binding the contract event 0x8ad601fbe08199f30553631f9519230e5fe1a3c6da337f967c91d2838f9986e1.
//
// Solidity: event Attested(address indexed v, uint64 indexed epoch)
func (_ValidatorReward *ValidatorRewardFilterer) ParseAttested(log types.Log) (*ValidatorRewardAttested, error) {
	event := new(ValidatorRewardAttested)
	if err := _ValidatorReward.contract.UnpackLog(event, "Attested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the ValidatorReward contract.
type ValidatorRewardClaimedIterator struct {
	Event *ValidatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardClaimed)
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
		it.Event = new(ValidatorRewardClaimed)
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
func (it *ValidatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardClaimed represents a Claimed event raised by the ValidatorReward contract.
type ValidatorRewardClaimed struct {
	V      common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed v, uint256 amount)
func (_ValidatorReward *ValidatorRewardFilterer) FilterClaimed(opts *bind.FilterOpts, v []common.Address) (*ValidatorRewardClaimedIterator, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "Claimed", vRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardClaimedIterator{contract: _ValidatorReward.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed v, uint256 amount)
func (_ValidatorReward *ValidatorRewardFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ValidatorRewardClaimed, v []common.Address) (event.Subscription, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "Claimed", vRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardClaimed)
				if err := _ValidatorReward.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed v, uint256 amount)
func (_ValidatorReward *ValidatorRewardFilterer) ParseClaimed(log types.Log) (*ValidatorRewardClaimed, error) {
	event := new(ValidatorRewardClaimed)
	if err := _ValidatorReward.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardDistributedIterator is returned from FilterDistributed and is used to iterate over the raw logs and unpacked data for Distributed events raised by the ValidatorReward contract.
type ValidatorRewardDistributedIterator struct {
	Event *ValidatorRewardDistributed // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardDistributed)
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
		it.Event = new(ValidatorRewardDistributed)
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
func (it *ValidatorRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardDistributed represents a Distributed event raised by the ValidatorReward contract.
type ValidatorRewardDistributed struct {
	V      common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributed is a free log retrieval operation binding the contract event 0xb649c98f58055c520df0dcb5709eff2e931217ff2fb1e21376130d31bbb1c0af.
//
// Solidity: event Distributed(address indexed v, uint256 amount)
func (_ValidatorReward *ValidatorRewardFilterer) FilterDistributed(opts *bind.FilterOpts, v []common.Address) (*ValidatorRewardDistributedIterator, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "Distributed", vRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardDistributedIterator{contract: _ValidatorReward.contract, event: "Distributed", logs: logs, sub: sub}, nil
}

// WatchDistributed is a free log subscription operation binding the contract event 0xb649c98f58055c520df0dcb5709eff2e931217ff2fb1e21376130d31bbb1c0af.
//
// Solidity: event Distributed(address indexed v, uint256 amount)
func (_ValidatorReward *ValidatorRewardFilterer) WatchDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorRewardDistributed, v []common.Address) (event.Subscription, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "Distributed", vRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardDistributed)
				if err := _ValidatorReward.contract.UnpackLog(event, "Distributed", log); err != nil {
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

// ParseDistributed is a log parse operation binding the contract event 0xb649c98f58055c520df0dcb5709eff2e931217ff2fb1e21376130d31bbb1c0af.
//
// Solidity: event Distributed(address indexed v, uint256 amount)
func (_ValidatorReward *ValidatorRewardFilterer) ParseDistributed(log types.Log) (*ValidatorRewardDistributed, error) {
	event := new(ValidatorRewardDistributed)
	if err := _ValidatorReward.contract.UnpackLog(event, "Distributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardFundedIterator is returned from FilterFunded and is used to iterate over the raw logs and unpacked data for Funded events raised by the ValidatorReward contract.
type ValidatorRewardFundedIterator struct {
	Event *ValidatorRewardFunded // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardFunded)
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
		it.Event = new(ValidatorRewardFunded)
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
func (it *ValidatorRewardFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardFunded represents a Funded event raised by the ValidatorReward contract.
type ValidatorRewardFunded struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFunded is a free log retrieval operation binding the contract event 0x5af8184bef8e4b45eb9f6ed7734d04da38ced226495548f46e0c8ff8d7d9a524.
//
// Solidity: event Funded(address indexed from, uint256 amount)
func (_ValidatorReward *ValidatorRewardFilterer) FilterFunded(opts *bind.FilterOpts, from []common.Address) (*ValidatorRewardFundedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "Funded", fromRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardFundedIterator{contract: _ValidatorReward.contract, event: "Funded", logs: logs, sub: sub}, nil
}

// WatchFunded is a free log subscription operation binding the contract event 0x5af8184bef8e4b45eb9f6ed7734d04da38ced226495548f46e0c8ff8d7d9a524.
//
// Solidity: event Funded(address indexed from, uint256 amount)
func (_ValidatorReward *ValidatorRewardFilterer) WatchFunded(opts *bind.WatchOpts, sink chan<- *ValidatorRewardFunded, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "Funded", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardFunded)
				if err := _ValidatorReward.contract.UnpackLog(event, "Funded", log); err != nil {
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

// ParseFunded is a log parse operation binding the contract event 0x5af8184bef8e4b45eb9f6ed7734d04da38ced226495548f46e0c8ff8d7d9a524.
//
// Solidity: event Funded(address indexed from, uint256 amount)
func (_ValidatorReward *ValidatorRewardFilterer) ParseFunded(log types.Log) (*ValidatorRewardFunded, error) {
	event := new(ValidatorRewardFunded)
	if err := _ValidatorReward.contract.UnpackLog(event, "Funded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ValidatorReward contract.
type ValidatorRewardInitializedIterator struct {
	Event *ValidatorRewardInitialized // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardInitialized)
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
		it.Event = new(ValidatorRewardInitialized)
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
func (it *ValidatorRewardInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardInitialized represents a Initialized event raised by the ValidatorReward contract.
type ValidatorRewardInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ValidatorReward *ValidatorRewardFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValidatorRewardInitializedIterator, error) {

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardInitializedIterator{contract: _ValidatorReward.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ValidatorReward *ValidatorRewardFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorRewardInitialized) (event.Subscription, error) {

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardInitialized)
				if err := _ValidatorReward.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ValidatorReward *ValidatorRewardFilterer) ParseInitialized(log types.Log) (*ValidatorRewardInitialized, error) {
	event := new(ValidatorRewardInitialized)
	if err := _ValidatorReward.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ValidatorReward contract.
type ValidatorRewardRoleAdminChangedIterator struct {
	Event *ValidatorRewardRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardRoleAdminChanged)
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
		it.Event = new(ValidatorRewardRoleAdminChanged)
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
func (it *ValidatorRewardRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardRoleAdminChanged represents a RoleAdminChanged event raised by the ValidatorReward contract.
type ValidatorRewardRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ValidatorReward *ValidatorRewardFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ValidatorRewardRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardRoleAdminChangedIterator{contract: _ValidatorReward.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ValidatorReward *ValidatorRewardFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ValidatorRewardRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardRoleAdminChanged)
				if err := _ValidatorReward.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ValidatorReward *ValidatorRewardFilterer) ParseRoleAdminChanged(log types.Log) (*ValidatorRewardRoleAdminChanged, error) {
	event := new(ValidatorRewardRoleAdminChanged)
	if err := _ValidatorReward.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ValidatorReward contract.
type ValidatorRewardRoleGrantedIterator struct {
	Event *ValidatorRewardRoleGranted // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardRoleGranted)
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
		it.Event = new(ValidatorRewardRoleGranted)
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
func (it *ValidatorRewardRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardRoleGranted represents a RoleGranted event raised by the ValidatorReward contract.
type ValidatorRewardRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ValidatorReward *ValidatorRewardFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ValidatorRewardRoleGrantedIterator, error) {

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

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardRoleGrantedIterator{contract: _ValidatorReward.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ValidatorReward *ValidatorRewardFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ValidatorRewardRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardRoleGranted)
				if err := _ValidatorReward.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ValidatorReward *ValidatorRewardFilterer) ParseRoleGranted(log types.Log) (*ValidatorRewardRoleGranted, error) {
	event := new(ValidatorRewardRoleGranted)
	if err := _ValidatorReward.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ValidatorReward contract.
type ValidatorRewardRoleRevokedIterator struct {
	Event *ValidatorRewardRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardRoleRevoked)
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
		it.Event = new(ValidatorRewardRoleRevoked)
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
func (it *ValidatorRewardRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardRoleRevoked represents a RoleRevoked event raised by the ValidatorReward contract.
type ValidatorRewardRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ValidatorReward *ValidatorRewardFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ValidatorRewardRoleRevokedIterator, error) {

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

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardRoleRevokedIterator{contract: _ValidatorReward.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ValidatorReward *ValidatorRewardFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ValidatorRewardRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardRoleRevoked)
				if err := _ValidatorReward.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ValidatorReward *ValidatorRewardFilterer) ParseRoleRevoked(log types.Log) (*ValidatorRewardRoleRevoked, error) {
	event := new(ValidatorRewardRoleRevoked)
	if err := _ValidatorReward.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ValidatorReward contract.
type ValidatorRewardUpgradedIterator struct {
	Event *ValidatorRewardUpgraded // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardUpgraded)
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
		it.Event = new(ValidatorRewardUpgraded)
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
func (it *ValidatorRewardUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardUpgraded represents a Upgraded event raised by the ValidatorReward contract.
type ValidatorRewardUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ValidatorReward *ValidatorRewardFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ValidatorRewardUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardUpgradedIterator{contract: _ValidatorReward.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ValidatorReward *ValidatorRewardFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ValidatorRewardUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardUpgraded)
				if err := _ValidatorReward.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ValidatorReward *ValidatorRewardFilterer) ParseUpgraded(log types.Log) (*ValidatorRewardUpgraded, error) {
	event := new(ValidatorRewardUpgraded)
	if err := _ValidatorReward.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the ValidatorReward contract.
type ValidatorRewardValidatorAddedIterator struct {
	Event *ValidatorRewardValidatorAdded // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardValidatorAdded)
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
		it.Event = new(ValidatorRewardValidatorAdded)
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
func (it *ValidatorRewardValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardValidatorAdded represents a ValidatorAdded event raised by the ValidatorReward contract.
type ValidatorRewardValidatorAdded struct {
	V   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed v)
func (_ValidatorReward *ValidatorRewardFilterer) FilterValidatorAdded(opts *bind.FilterOpts, v []common.Address) (*ValidatorRewardValidatorAddedIterator, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "ValidatorAdded", vRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardValidatorAddedIterator{contract: _ValidatorReward.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed v)
func (_ValidatorReward *ValidatorRewardFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ValidatorRewardValidatorAdded, v []common.Address) (event.Subscription, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "ValidatorAdded", vRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardValidatorAdded)
				if err := _ValidatorReward.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed v)
func (_ValidatorReward *ValidatorRewardFilterer) ParseValidatorAdded(log types.Log) (*ValidatorRewardValidatorAdded, error) {
	event := new(ValidatorRewardValidatorAdded)
	if err := _ValidatorReward.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the ValidatorReward contract.
type ValidatorRewardValidatorRemovedIterator struct {
	Event *ValidatorRewardValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardValidatorRemoved)
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
		it.Event = new(ValidatorRewardValidatorRemoved)
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
func (it *ValidatorRewardValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardValidatorRemoved represents a ValidatorRemoved event raised by the ValidatorReward contract.
type ValidatorRewardValidatorRemoved struct {
	V   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed v)
func (_ValidatorReward *ValidatorRewardFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, v []common.Address) (*ValidatorRewardValidatorRemovedIterator, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}

	logs, sub, err := _ValidatorReward.contract.FilterLogs(opts, "ValidatorRemoved", vRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardValidatorRemovedIterator{contract: _ValidatorReward.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed v)
func (_ValidatorReward *ValidatorRewardFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ValidatorRewardValidatorRemoved, v []common.Address) (event.Subscription, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}

	logs, sub, err := _ValidatorReward.contract.WatchLogs(opts, "ValidatorRemoved", vRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardValidatorRemoved)
				if err := _ValidatorReward.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed v)
func (_ValidatorReward *ValidatorRewardFilterer) ParseValidatorRemoved(log types.Log) (*ValidatorRewardValidatorRemoved, error) {
	event := new(ValidatorRewardValidatorRemoved)
	if err := _ValidatorReward.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
