// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// CheckpointsCheckpoint208 is an auto generated low-level Go binding around an user-defined struct.
type CheckpointsCheckpoint208 struct {
	Key   *big.Int
	Value *big.Int
}

// GovernanceTokenMetaData contains all meta data concerning the GovernanceToken contract.
var GovernanceTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pos\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCheckpoints.Checkpoint208\",\"components\":[{\"name\":\"_key\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"_value\",\"type\":\"uint208\",\"internalType\":\"uint208\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateBySig\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegates\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPastTotalSupply\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPastVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateChanged\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fromDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateVotesChanged\",\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"previousVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20ExceededSafeSupply\",\"inputs\":[{\"name\":\"increasedSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2612ExpiredSignature\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC2612InvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC5805FutureLookup\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clock\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"ERC6372InconsistentClock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"VotesExpiredSignature\",\"inputs\":[{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x608080604052346100165761295d908161001b8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c806306fdde0314610194578063095ea7b31461018f57806318160ddd1461018a57806323b872dd14610185578063313ce567146101805780633644e5151461017b5780633a46b1a8146101765780634bf5d7e914610171578063587cde1e1461016c5780635c19a95c146101675780636fcfff451461016257806370a082311461015d5780637ecebe001461015857806384b0196e146101535780638e539e8c1461014e57806391ddadf41461014957806395d89b41146101445780639ab24eb01461013f578063a9059cbb1461013a578063bd3a13f614610135578063c3cda52014610130578063d505accf1461012b578063dd62ed3e146101265763f1127ed814610121575f80fd5b610f9d565b610f65565b610e13565b610d4a565b610bca565b610acc565b610a95565b6109e3565b6109b8565b6108e3565b610813565b610736565b6106df565b61068a565b610666565b610621565b610590565b6104be565b61049c565b610481565b61038e565b610352565b610321565b6101eb565b91908251928382525f5b8481106101c3575050825f602080949584010152601f8019910116010190565b6020818301810151848301820152016101a3565b9060206101e8928181520190610199565b90565b346102db575f3660031901126102db576040515f5f8051602061284883398151915280549061021982611030565b808552916020916001918281169081156102ae5750600114610256575b6102528661024681880382610b43565b604051918291826101d7565b0390f35b5f90815293507f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab05b83851061029b57505050508101602001610246826102525f610236565b805486860184015293820193810161027e565b90508695506102529693506020925061024694915060ff191682840152151560051b82010192935f610236565b5f80fd5b600435906001600160a01b03821682036102db57565b602435906001600160a01b03821682036102db57565b606435906001600160a01b03821682036102db57565b346102db5760403660031901126102db5761034761033d6102df565b60243590336119a0565b602060405160018152f35b346102db575f3660031901126102db5760207f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0254604051908152f35b346102db5760603660031901126102db576103a76102df565b6103af6102f5565b604435906103d5336103c0856111e7565b9060018060a01b03165f5260205260405f2090565b54925f1984106103f6575b6103ea9350611407565b60405160018152602090f35b82841061045b576001600160a01b0381161561044357331561042b57826103ea9403610425336103c0846111e7565b556103e0565b604051634a1406b160e11b81525f6004820152602490fd5b60405163e602df0560e01b81525f6004820152602490fd5b604051637dc7a0d960e11b81523360048201526024810185905260448101849052606490fd5b346102db575f3660031901126102db57602060405160128152f35b346102db575f3660031901126102db5760206104b66119f8565b604051908152f35b346102db5760403660031901126102db576104df6104da6102df565b61121f565b6104ea6024356114e7565b8154905f82916005841161053d575b610504935084611c5b565b908161052257505060205f5b6040516001600160d01b039091168152f35b61052d60209261153c565b905f52815f20015460301c610510565b919261054881611abe565b810390811161058b5761050493855f5265ffffffffffff808360205f20015416908516105f146105795750916104f9565b9291506105859061154a565b906104f9565b611528565b346102db575f3660031901126102db576105a943611a6e565b65ffffffffffff806105ba43611a6e565b1691160361060f576102526040516105d181610b06565b601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c740000006020820152604051918291602083526020830190610199565b6040516301bfc1c560e61b8152600490fd5b346102db5760203660031901126102db5760206001600160a01b03806106456102df565b165f525f80516020612828833981519152825260405f205416604051908152f35b346102db5760203660031901126102db576106886106826102df565b33611558565b005b346102db5760203660031901126102db576106a66104da6102df565b5463ffffffff908181116106c05760209160405191168152f35b604490604051906306dfcc6560e41b8252602060048301526024820152fd5b346102db5760203660031901126102db5760206104b66106fd6102df565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b346102db5760203660031901126102db576001600160a01b036107576102df565b165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052602060405f2054604051908152f35b916107c190949194600f60f81b84526107b360209660e0602087015260e0860190610199565b908482036040860152610199565b92606083015260018060a01b031660808201525f60a082015260c0818303910152602080845192838152019301915f5b8281106107ff575050505090565b8351855293810193928101926001016107f1565b346102db575f3660031901126102db577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1005415806108ba575b1561087d57610859611068565b61086161113b565b9061025261086d61128f565b604051938493309146918661078d565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101541561084c565b346102db5760203660031901126102db576108ff6004356114e7565b5f805160206128c8833981519152908154905f829160058411610960575b6109279350611be9565b908161093a5750506040515f8152602090f35b61094560209261153c565b905f525f80516020612908833981519152015460301c610510565b919261096b81611abe565b810390811161058b5761092793855f5265ffffffffffff80835f80516020612908833981519152015416908516105f146109a657509161091d565b9291506109b29061154a565b9061091d565b346102db575f3660031901126102db5760206109d343611a6e565b65ffffffffffff60405191168152f35b346102db575f3660031901126102db576040515f5f80516020612888833981519152805490610a1182611030565b808552916020916001918281169081156102ae5750600114610a3d576102528661024681880382610b43565b5f90815293507f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa5b838510610a8257505050508101602001610246826102525f610236565b8054868601840152938201938101610a65565b346102db5760203660031901126102db5760206001600160d01b03610ac3610abe6104da6102df565b611655565b16604051908152f35b346102db5760403660031901126102db57610347610ae86102df565b6024359033611407565b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff821117610b2257604052565b610af2565b60a0810190811067ffffffffffffffff821117610b2257604052565b90601f8019910116810190811067ffffffffffffffff821117610b2257604052565b60405190610b7282610b06565b565b81601f820112156102db5780359067ffffffffffffffff8211610b225760405192610ba9601f8401601f191660200185610b43565b828452602083830101116102db57815f926020809301838601378301015290565b346102db5760803660031901126102db5767ffffffffffffffff6004358181116102db57610bfc903690600401610b74565b6024358281116102db57610c14903690600401610b74565b90610c1d61030b565b905f805160206128e8833981519152549360ff8560401c1615941680159081610d22575b6001149081610d18575b159081610d0f575b50610cfd575f805160206128e8833981519152805467ffffffffffffffff19166001179055610c8b9284610cd9575b604435916112b6565b610c9157005b5f805160206128e8833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b5f805160206128e8833981519152805460ff60401b1916600160401b179055610c82565b60405163f92ee8a960e01b8152600490fd5b9050155f610c53565b303b159150610c4b565b859150610c41565b6064359060ff821682036102db57565b6084359060ff821682036102db57565b346102db5760c03660031901126102db57610d636102df565b60443590602435610d72610d2a565b834211610dfa57610dee61068894610df5926040519060208201927fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf845260018060a01b0388166040840152866060840152608083015260808252610dd682610b27565b610de960a43593608435935190206118ee565b611914565b918261192c565b611558565b604051632341d78760e11b815260048101859052602490fd5b346102db5760e03660031901126102db57610e2c6102df565b610e346102f5565b60443590606435610e43610d3a565b814211610f4c576001600160a01b038581165f8181527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602090815260409182902080546001810190915582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c99281019283529283019390935292861660608201526080810187905260a081019190915260c080820194909452928352610f05929091610ef260e083610b43565b610de960c4359360a435935190206118ee565b6001600160a01b0384811690821603610f225750610688926119a0565b6040516325c0072360e11b81526001600160a01b0391821660048201529084166024820152604490fd5b60405163313c898160e11b815260048101839052602490fd5b346102db5760403660031901126102db576020610f94610f836102df565b6103c0610f8e6102f5565b916111e7565b54604051908152f35b346102db5760403660031901126102db57610fb66102df565b6024359063ffffffff821682036102db57604091610fe8610ff692610fd96113ef565b50610fe26113ef565b5061121f565b610ff06113ef565b5061230f565b5081519061100382610b06565b54602065ffffffffffff821692838152019060301c8152825191825260018060d01b039051166020820152f35b90600182811c9216801561105e575b602083101461104a57565b634e487b7160e01b5f52602260045260245ffd5b91607f169161103f565b604051905f825f805160206128688339815191529182549261108984611030565b8084529360209160019182811690811561111557506001146110b5575b505050610b7292500383610b43565b5f9081527f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d9590935091905b8284106110fd5750610b729450505081016020015f80806110a6565b855488850183015294850194879450928101926110e1565b9250505060209250610b7294915060ff191682840152151560051b8201015f80806110a6565b604051905f825f805160206128a88339815191529182549261115c84611030565b80845293602091600191828116908115611115575060011461118757505050610b7292500383610b43565b5f9081527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b759590935091905b8284106111cf5750610b729450505081016020015f80806110a6565b855488850183015294850194879450928101926111b3565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b6001600160a01b03165f9081527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d016020526040902090565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace006020526040902090565b6040516020810181811067ffffffffffffffff821117610b22576040525f8152905f368137565b939290936112c2611dee565b6112ca611dee565b805167ffffffffffffffff8111610b22575f80516020612848833981519152906112fd816112f88454611030565b611e1c565b602080601f831160011461135557509661134593926113398361134094610b729a9b5f9161134a575b508160011b915f199060031b1c19161790565b9055612000565b61167d565b61180e565b90508601515f611326565b5f805160206128488339815191525f5290601f1983167f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0925f905b8282106113d757505083610b72999a9361134597969361134096600194106113bf575b5050811b019055612000565b8701515f1960f88460031b161c191690555f806113b3565b80600185968294968b01518155019501930190611390565b604051906113fc82610b06565b5f6020838281520152565b6001600160a01b03808216949392919085156114cf57821680156114b75761142e82611257565b54958487106114885784610b7296970361144784611257565b5561145184611257565b8054860190556040518581527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602090a361253c565b60405163391434e360e21b81526001600160a01b03841660048201526024810188905260448101869052606490fd5b60405163ec442f0560e01b81525f6004820152602490fd5b604051634b637e8f60e11b81525f6004820152602490fd5b65ffffffffffff6114f743611a6e565b168082101561150a57506101e890611a6e565b6044925060405191637669fc0f60e11b835260048301526024820152fd5b634e487b7160e01b5f52601160045260245ffd5b5f1981019190821161058b57565b906001820180921161058b57565b6001600160a01b038181165f8181525f805160206128288339815191526020526040812080548685166001600160a01b031982168117909255610b72969416946115fe9390928691907f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9080a46001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b91611cba565b5f805160206128c883398151915280548061161f5750505f90565b805f1981011161058b577f88c46c62109817164d0ae1873830d4299a82e5daf552a3d8e989b27638fcf747915f52015460301c90565b8054806116625750505f90565b5f1991818381011161058b575f5260205f2001015460301c90565b90611686611dee565b60405161169281610b06565b6001808252602090603160f81b60208401526116ac611dee565b84519067ffffffffffffffff8211610b22575f80516020612868833981519152926116e0836116db8654611030565b611e95565b602091601f841160011461177857505090806117169261171d96975f9261176d575b50508160011b915f199060031b1c19161790565b90556120fa565b6117455f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10055565b610b725f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10155565b015190505f80611702565b5f805160206128688339815191525f908152601f198516987f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d949390925b8a83106117f7575050509083929160019461171d9899106117df575b505050811b0190556120fa565b01515f1960f88460031b161c191690555f80806117d2565b8385015186559485019493810193918101916117b6565b91906001600160a01b03831680156114b7577f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0290815483810180911161058b577f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025561187985611257565b8054840190556040518381525f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602090a354926001600160d01b0384116118c857610b729293506124b3565b604051630e58ae9360e11b8152600481018590526001600160d01b036024820152604490fd5b6042906118f96119f8565b906040519161190160f01b8352600283015260228201522090565b916101e89391611923936121d7565b90929192612282565b6001600160a01b0381165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052604090208054600181019091559091819003611978575050565b6040516301d4b62360e61b81526001600160a01b039092166004830152602482015260449150fd5b6001600160a01b03808216929190831561044357821693841561042b57806119ee7f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925946103c06020956111e7565b55604051908152a3565b611a00612338565b611a086123a2565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815260c0810181811067ffffffffffffffff821117610b225760405251902090565b65ffffffffffff90818111611a81571690565b604490604051906306dfcc6560e41b8252603060048301526024820152fd5b8115611aaa570490565b634e487b7160e01b5f52601260045260245ffd5b60018111156101e857600181600160801b811015611bd7575b611b7f611b75611b6b611b61611b57611b4d611b8b97600488600160401b611b869a1015611bca575b640100000000811015611bbd575b62010000811015611bb0575b610100811015611ba4575b6010811015611b98575b1015611b90575b60030260011c611b46818b611aa0565b0160011c90565b611b46818a611aa0565b611b468189611aa0565b611b468188611aa0565b611b468187611aa0565b611b468186611aa0565b8093611aa0565b821190565b900390565b60011b611b36565b811c9160021b91611b2f565b60081c91811b91611b25565b60101c9160081b91611b1a565b60201c9160101b91611b0e565b60401c9160201b91611b00565b50600160401b9050608082901c611ad7565b905b828110611bf757505090565b90918082169080831860011c820180921161058b575f805160206128c88339815191525f5265ffffffffffff80835f80516020612908833981519152015416908516105f14611c495750915b90611beb565b929150611c559061154a565b90611c43565b91905b838210611c6b5750505090565b9091928083169080841860011c820180921161058b57845f5265ffffffffffff808360205f20015416908416105f14611ca85750925b9190611c5e565b939250611cb49061154a565b91611ca1565b6001600160a01b03808316939291908116908185141580611de5575b611ce2575b5050505050565b81611d57575b505082611cf7575b8080611cdb565b7fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a72491611d2e611d28611d349361121f565b916123e7565b9061241a565b604080516001600160d01b039384168152919092166020820152a25f8080611cf0565b611d609061121f565b611d69846123e7565b611d7243611a6e565b6001600160d01b03918280611d8686611655565b16911690039282841161058b577fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a72493611ddb92611dc29261276d565b6040805192851683529316602082015291829190820190565b0390a25f80611ce8565b50831515611cd6565b60ff5f805160206128e88339815191525460401c1615611e0a57565b604051631afcd79f60e31b8152600490fd5b601f8111611e28575050565b5f805160206128488339815191525f527f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0906020601f840160051c83019310611e8b575b601f0160051c01905b818110611e80575050565b5f8155600101611e75565b9091508190611e6c565b601f8111611ea1575050565b5f805160206128688339815191525f527f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d906020601f840160051c83019310611f04575b601f0160051c01905b818110611ef9575050565b5f8155600101611eee565b9091508190611ee5565b601f8111611f1a575050565b5f805160206128888339815191525f527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa906020601f840160051c83019310611f7d575b601f0160051c01905b818110611f72575050565b5f8155600101611f67565b9091508190611f5e565b601f8111611f93575050565b5f805160206128a88339815191525f527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75906020601f840160051c83019310611ff6575b601f0160051c01905b818110611feb575050565b5f8155600101611fe0565b9091508190611fd7565b90815167ffffffffffffffff8111610b22575f80516020612888833981519152906120348161202f8454611030565b611f0e565b602080601f8311600114612069575081906120659394955f9261176d5750508160011b915f199060031b1c19161790565b9055565b90601f198316956120a75f805160206128888339815191525f527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa90565b925f905b8882106120e2575050836001959697106120ca575b505050811b019055565b01515f1960f88460031b161c191690555f80806120c0565b806001859682949686015181550195019301906120ab565b90815167ffffffffffffffff8111610b22575f805160206128a88339815191529061212e816121298454611030565b611f87565b602080601f831160011461215f575081906120659394955f9261176d5750508160011b915f199060031b1c19161790565b90601f1983169561219d5f805160206128a88339815191525f527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7590565b925f905b8882106121bf575050836001959697106120ca57505050811b019055565b806001859682949686015181550195019301906121a1565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411612259579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa1561224e575f516001600160a01b0381161561224457905f905f90565b505f906001905f90565b6040513d5f823e3d90fd5b5050505f9160039190565b6004111561226e57565b634e487b7160e01b5f52602160045260245ffd5b61228b81612264565b80612294575050565b61229d81612264565b600181036122b75760405163f645eedf60e01b8152600490fd5b6122c081612264565b600281036122e15760405163fce698f760e01b815260048101839052602490fd5b806122ed600392612264565b146122f55750565b6040516335e2f38360e21b81526004810191909152602490fd5b8054821015612324575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b612340611068565b8051908115612350576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10054801561237d5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b6123aa61113b565b80519081156123ba576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10154801561237d5790565b6001600160d01b03908181116123fb571690565b604490604051906306dfcc6560e41b825260d060048301526024820152fd5b9061242443611a6e565b6001600160d01b0391828061243886611655565b1691160191821161058b5761244c9261276d565b9091565b61245943611a6e565b906001600160d01b0390818061246d611604565b1691160190811161058b5761244c9161266f565b61248a43611a6e565b906001600160d01b0390818061249e611604565b169116900390811161058b5761244c9161266f565b90610b72916124c96124c4836123e7565b612450565b50506001600160a01b03908116908115612524575b5f805160206128288339815191526020527fd4fb29e10204005f1a39963c6862b79a755e22f0177c53f05cdc3786c702f974545f92835260409092205481169116611cba565b612535612530846123e7565b612481565b50506124de565b610b7292916001600160a01b0391821691908190831561259c575b16918215612589575b5f525f805160206128288339815191526020528060405f205416915f5260405f20541690611cba565b612595612530856123e7565b5050612560565b6125a86124c4866123e7565b5050612557565b5f805160206128c8833981519152908154600160401b811015610b225760018101808455811015612324575f92909252805160209091015160301b65ffffffffffff191665ffffffffffff91909116175f805160206129088339815191529190910155565b8054600160401b811015610b22576126319160018201815561230f565b61265c57815160209092015160301b65ffffffffffff191665ffffffffffff92909216919091179055565b634e487b7160e01b5f525f60045260245ffd5b5f805160206128c8833981519152549192918015612744576126936126b59161153c565b5f805160206128c88339815191525f525f805160206129088339815191520190565b9081549165ffffffffffff90818416918316808311612732578692036126fa576126f392509065ffffffffffff82549181199060301b169116179055565b60301c9190565b505061272d9061271961270b610b65565b65ffffffffffff9092168252565b6001600160d01b03851660208201526125af565b6126f3565b604051632520601d60e01b8152600490fd5b506127689061275461270b610b65565b6001600160d01b03841660208201526125af565b5f9190565b805492939280156128025761278461278f9161153c565b825f5260205f200190565b9182549265ffffffffffff91828516928116808411612732578793036127ce57506126f392509065ffffffffffff82549181199060301b169116179055565b91505061272d916127ee6127e0610b65565b65ffffffffffff9093168352565b6001600160d01b0386166020830152612614565b5090612768916128136127e0610b65565b6001600160d01b038516602083015261261456fee8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d0052c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10252c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace04a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103e8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d02f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0088c46c62109817164d0ae1873830d4299a82e5daf552a3d8e989b27638fcf748a26469706673582212206e702fad43107abeb03a32849644a7d13d7778944b07db9841e7803de736b7c364736f6c63430008180033",
}

// GovernanceTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceTokenMetaData.ABI instead.
var GovernanceTokenABI = GovernanceTokenMetaData.ABI

// GovernanceTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernanceTokenMetaData.Bin instead.
var GovernanceTokenBin = GovernanceTokenMetaData.Bin

// DeployGovernanceToken deploys a new Ethereum contract, binding an instance of GovernanceToken to it.
func DeployGovernanceToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GovernanceToken, error) {
	parsed, err := GovernanceTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernanceToken{GovernanceTokenCaller: GovernanceTokenCaller{contract: contract}, GovernanceTokenTransactor: GovernanceTokenTransactor{contract: contract}, GovernanceTokenFilterer: GovernanceTokenFilterer{contract: contract}}, nil
}

// GovernanceToken is an auto generated Go binding around an Ethereum contract.
type GovernanceToken struct {
	GovernanceTokenCaller     // Read-only binding to the contract
	GovernanceTokenTransactor // Write-only binding to the contract
	GovernanceTokenFilterer   // Log filterer for contract events
}

// GovernanceTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceTokenSession struct {
	Contract     *GovernanceToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceTokenCallerSession struct {
	Contract *GovernanceTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GovernanceTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTokenTransactorSession struct {
	Contract     *GovernanceTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GovernanceTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceTokenRaw struct {
	Contract *GovernanceToken // Generic contract binding to access the raw methods on
}

// GovernanceTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceTokenCallerRaw struct {
	Contract *GovernanceTokenCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTokenTransactorRaw struct {
	Contract *GovernanceTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceToken creates a new instance of GovernanceToken, bound to a specific deployed contract.
func NewGovernanceToken(address common.Address, backend bind.ContractBackend) (*GovernanceToken, error) {
	contract, err := bindGovernanceToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceToken{GovernanceTokenCaller: GovernanceTokenCaller{contract: contract}, GovernanceTokenTransactor: GovernanceTokenTransactor{contract: contract}, GovernanceTokenFilterer: GovernanceTokenFilterer{contract: contract}}, nil
}

// NewGovernanceTokenCaller creates a new read-only instance of GovernanceToken, bound to a specific deployed contract.
func NewGovernanceTokenCaller(address common.Address, caller bind.ContractCaller) (*GovernanceTokenCaller, error) {
	contract, err := bindGovernanceToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenCaller{contract: contract}, nil
}

// NewGovernanceTokenTransactor creates a new write-only instance of GovernanceToken, bound to a specific deployed contract.
func NewGovernanceTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTokenTransactor, error) {
	contract, err := bindGovernanceToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenTransactor{contract: contract}, nil
}

// NewGovernanceTokenFilterer creates a new log filterer instance of GovernanceToken, bound to a specific deployed contract.
func NewGovernanceTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceTokenFilterer, error) {
	contract, err := bindGovernanceToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenFilterer{contract: contract}, nil
}

// bindGovernanceToken binds a generic wrapper to an already deployed contract.
func bindGovernanceToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GovernanceTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceToken *GovernanceTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceToken.Contract.GovernanceTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceToken *GovernanceTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceToken.Contract.GovernanceTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceToken *GovernanceTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceToken.Contract.GovernanceTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceToken *GovernanceTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceToken *GovernanceTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceToken *GovernanceTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceToken.Contract.contract.Transact(opts, method, params...)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_GovernanceToken *GovernanceTokenCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_GovernanceToken *GovernanceTokenSession) CLOCKMODE() (string, error) {
	return _GovernanceToken.Contract.CLOCKMODE(&_GovernanceToken.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_GovernanceToken *GovernanceTokenCallerSession) CLOCKMODE() (string, error) {
	return _GovernanceToken.Contract.CLOCKMODE(&_GovernanceToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_GovernanceToken *GovernanceTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_GovernanceToken *GovernanceTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _GovernanceToken.Contract.DOMAINSEPARATOR(&_GovernanceToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_GovernanceToken *GovernanceTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _GovernanceToken.Contract.DOMAINSEPARATOR(&_GovernanceToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GovernanceToken *GovernanceTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.Allowance(&_GovernanceToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.Allowance(&_GovernanceToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GovernanceToken *GovernanceTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.BalanceOf(&_GovernanceToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.BalanceOf(&_GovernanceToken.CallOpts, account)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (_GovernanceToken *GovernanceTokenCaller) Checkpoints(opts *bind.CallOpts, account common.Address, pos uint32) (CheckpointsCheckpoint208, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "checkpoints", account, pos)

	if err != nil {
		return *new(CheckpointsCheckpoint208), err
	}

	out0 := *abi.ConvertType(out[0], new(CheckpointsCheckpoint208)).(*CheckpointsCheckpoint208)

	return out0, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (_GovernanceToken *GovernanceTokenSession) Checkpoints(account common.Address, pos uint32) (CheckpointsCheckpoint208, error) {
	return _GovernanceToken.Contract.Checkpoints(&_GovernanceToken.CallOpts, account, pos)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (_GovernanceToken *GovernanceTokenCallerSession) Checkpoints(account common.Address, pos uint32) (CheckpointsCheckpoint208, error) {
	return _GovernanceToken.Contract.Checkpoints(&_GovernanceToken.CallOpts, account, pos)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_GovernanceToken *GovernanceTokenCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_GovernanceToken *GovernanceTokenSession) Clock() (*big.Int, error) {
	return _GovernanceToken.Contract.Clock(&_GovernanceToken.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_GovernanceToken *GovernanceTokenCallerSession) Clock() (*big.Int, error) {
	return _GovernanceToken.Contract.Clock(&_GovernanceToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GovernanceToken *GovernanceTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GovernanceToken *GovernanceTokenSession) Decimals() (uint8, error) {
	return _GovernanceToken.Contract.Decimals(&_GovernanceToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GovernanceToken *GovernanceTokenCallerSession) Decimals() (uint8, error) {
	return _GovernanceToken.Contract.Decimals(&_GovernanceToken.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_GovernanceToken *GovernanceTokenCaller) Delegates(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "delegates", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_GovernanceToken *GovernanceTokenSession) Delegates(account common.Address) (common.Address, error) {
	return _GovernanceToken.Contract.Delegates(&_GovernanceToken.CallOpts, account)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_GovernanceToken *GovernanceTokenCallerSession) Delegates(account common.Address) (common.Address, error) {
	return _GovernanceToken.Contract.Delegates(&_GovernanceToken.CallOpts, account)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GovernanceToken *GovernanceTokenCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GovernanceToken *GovernanceTokenSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GovernanceToken.Contract.Eip712Domain(&_GovernanceToken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GovernanceToken *GovernanceTokenCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GovernanceToken.Contract.Eip712Domain(&_GovernanceToken.CallOpts)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCaller) GetPastTotalSupply(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "getPastTotalSupply", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (_GovernanceToken *GovernanceTokenSession) GetPastTotalSupply(timepoint *big.Int) (*big.Int, error) {
	return _GovernanceToken.Contract.GetPastTotalSupply(&_GovernanceToken.CallOpts, timepoint)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCallerSession) GetPastTotalSupply(timepoint *big.Int) (*big.Int, error) {
	return _GovernanceToken.Contract.GetPastTotalSupply(&_GovernanceToken.CallOpts, timepoint)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCaller) GetPastVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "getPastVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (_GovernanceToken *GovernanceTokenSession) GetPastVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _GovernanceToken.Contract.GetPastVotes(&_GovernanceToken.CallOpts, account, timepoint)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCallerSession) GetPastVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _GovernanceToken.Contract.GetPastVotes(&_GovernanceToken.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCaller) GetVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_GovernanceToken *GovernanceTokenSession) GetVotes(account common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.GetVotes(&_GovernanceToken.CallOpts, account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCallerSession) GetVotes(account common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.GetVotes(&_GovernanceToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernanceToken *GovernanceTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernanceToken *GovernanceTokenSession) Name() (string, error) {
	return _GovernanceToken.Contract.Name(&_GovernanceToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernanceToken *GovernanceTokenCallerSession) Name() (string, error) {
	return _GovernanceToken.Contract.Name(&_GovernanceToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_GovernanceToken *GovernanceTokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.Nonces(&_GovernanceToken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.Nonces(&_GovernanceToken.CallOpts, owner)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_GovernanceToken *GovernanceTokenCaller) NumCheckpoints(opts *bind.CallOpts, account common.Address) (uint32, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "numCheckpoints", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_GovernanceToken *GovernanceTokenSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _GovernanceToken.Contract.NumCheckpoints(&_GovernanceToken.CallOpts, account)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_GovernanceToken *GovernanceTokenCallerSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _GovernanceToken.Contract.NumCheckpoints(&_GovernanceToken.CallOpts, account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GovernanceToken *GovernanceTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GovernanceToken *GovernanceTokenSession) Symbol() (string, error) {
	return _GovernanceToken.Contract.Symbol(&_GovernanceToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GovernanceToken *GovernanceTokenCallerSession) Symbol() (string, error) {
	return _GovernanceToken.Contract.Symbol(&_GovernanceToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GovernanceToken *GovernanceTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GovernanceToken *GovernanceTokenSession) TotalSupply() (*big.Int, error) {
	return _GovernanceToken.Contract.TotalSupply(&_GovernanceToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GovernanceToken *GovernanceTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _GovernanceToken.Contract.TotalSupply(&_GovernanceToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_GovernanceToken *GovernanceTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GovernanceToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_GovernanceToken *GovernanceTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GovernanceToken.Contract.Approve(&_GovernanceToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_GovernanceToken *GovernanceTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GovernanceToken.Contract.Approve(&_GovernanceToken.TransactOpts, spender, value)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_GovernanceToken *GovernanceTokenTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _GovernanceToken.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_GovernanceToken *GovernanceTokenSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _GovernanceToken.Contract.Delegate(&_GovernanceToken.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_GovernanceToken *GovernanceTokenTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _GovernanceToken.Contract.Delegate(&_GovernanceToken.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_GovernanceToken *GovernanceTokenTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GovernanceToken.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_GovernanceToken *GovernanceTokenSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GovernanceToken.Contract.DelegateBySig(&_GovernanceToken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_GovernanceToken *GovernanceTokenTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GovernanceToken.Contract.DelegateBySig(&_GovernanceToken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0xbd3a13f6.
//
// Solidity: function initialize(string name, string symbol, uint256 initialSupply, address recipient) returns()
func (_GovernanceToken *GovernanceTokenTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, initialSupply *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GovernanceToken.contract.Transact(opts, "initialize", name, symbol, initialSupply, recipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xbd3a13f6.
//
// Solidity: function initialize(string name, string symbol, uint256 initialSupply, address recipient) returns()
func (_GovernanceToken *GovernanceTokenSession) Initialize(name string, symbol string, initialSupply *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GovernanceToken.Contract.Initialize(&_GovernanceToken.TransactOpts, name, symbol, initialSupply, recipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xbd3a13f6.
//
// Solidity: function initialize(string name, string symbol, uint256 initialSupply, address recipient) returns()
func (_GovernanceToken *GovernanceTokenTransactorSession) Initialize(name string, symbol string, initialSupply *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GovernanceToken.Contract.Initialize(&_GovernanceToken.TransactOpts, name, symbol, initialSupply, recipient)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_GovernanceToken *GovernanceTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GovernanceToken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_GovernanceToken *GovernanceTokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GovernanceToken.Contract.Permit(&_GovernanceToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_GovernanceToken *GovernanceTokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GovernanceToken.Contract.Permit(&_GovernanceToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_GovernanceToken *GovernanceTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GovernanceToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_GovernanceToken *GovernanceTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GovernanceToken.Contract.Transfer(&_GovernanceToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_GovernanceToken *GovernanceTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GovernanceToken.Contract.Transfer(&_GovernanceToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_GovernanceToken *GovernanceTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GovernanceToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_GovernanceToken *GovernanceTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GovernanceToken.Contract.TransferFrom(&_GovernanceToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_GovernanceToken *GovernanceTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GovernanceToken.Contract.TransferFrom(&_GovernanceToken.TransactOpts, from, to, value)
}

// GovernanceTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GovernanceToken contract.
type GovernanceTokenApprovalIterator struct {
	Event *GovernanceTokenApproval // Event containing the contract specifics and raw log

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
func (it *GovernanceTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTokenApproval)
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
		it.Event = new(GovernanceTokenApproval)
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
func (it *GovernanceTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTokenApproval represents a Approval event raised by the GovernanceToken contract.
type GovernanceTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GovernanceTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GovernanceToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenApprovalIterator{contract: _GovernanceToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GovernanceTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GovernanceToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTokenApproval)
				if err := _GovernanceToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) ParseApproval(log types.Log) (*GovernanceTokenApproval, error) {
	event := new(GovernanceTokenApproval)
	if err := _GovernanceToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTokenDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the GovernanceToken contract.
type GovernanceTokenDelegateChangedIterator struct {
	Event *GovernanceTokenDelegateChanged // Event containing the contract specifics and raw log

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
func (it *GovernanceTokenDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTokenDelegateChanged)
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
		it.Event = new(GovernanceTokenDelegateChanged)
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
func (it *GovernanceTokenDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTokenDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTokenDelegateChanged represents a DelegateChanged event raised by the GovernanceToken contract.
type GovernanceTokenDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_GovernanceToken *GovernanceTokenFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*GovernanceTokenDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _GovernanceToken.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenDelegateChangedIterator{contract: _GovernanceToken.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_GovernanceToken *GovernanceTokenFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *GovernanceTokenDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _GovernanceToken.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTokenDelegateChanged)
				if err := _GovernanceToken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_GovernanceToken *GovernanceTokenFilterer) ParseDelegateChanged(log types.Log) (*GovernanceTokenDelegateChanged, error) {
	event := new(GovernanceTokenDelegateChanged)
	if err := _GovernanceToken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTokenDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the GovernanceToken contract.
type GovernanceTokenDelegateVotesChangedIterator struct {
	Event *GovernanceTokenDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *GovernanceTokenDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTokenDelegateVotesChanged)
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
		it.Event = new(GovernanceTokenDelegateVotesChanged)
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
func (it *GovernanceTokenDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTokenDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTokenDelegateVotesChanged represents a DelegateVotesChanged event raised by the GovernanceToken contract.
type GovernanceTokenDelegateVotesChanged struct {
	Delegate      common.Address
	PreviousVotes *big.Int
	NewVotes      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_GovernanceToken *GovernanceTokenFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*GovernanceTokenDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _GovernanceToken.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenDelegateVotesChangedIterator{contract: _GovernanceToken.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_GovernanceToken *GovernanceTokenFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *GovernanceTokenDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _GovernanceToken.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTokenDelegateVotesChanged)
				if err := _GovernanceToken.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_GovernanceToken *GovernanceTokenFilterer) ParseDelegateVotesChanged(log types.Log) (*GovernanceTokenDelegateVotesChanged, error) {
	event := new(GovernanceTokenDelegateVotesChanged)
	if err := _GovernanceToken.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTokenEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the GovernanceToken contract.
type GovernanceTokenEIP712DomainChangedIterator struct {
	Event *GovernanceTokenEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *GovernanceTokenEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTokenEIP712DomainChanged)
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
		it.Event = new(GovernanceTokenEIP712DomainChanged)
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
func (it *GovernanceTokenEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTokenEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTokenEIP712DomainChanged represents a EIP712DomainChanged event raised by the GovernanceToken contract.
type GovernanceTokenEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GovernanceToken *GovernanceTokenFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*GovernanceTokenEIP712DomainChangedIterator, error) {

	logs, sub, err := _GovernanceToken.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenEIP712DomainChangedIterator{contract: _GovernanceToken.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GovernanceToken *GovernanceTokenFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *GovernanceTokenEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _GovernanceToken.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTokenEIP712DomainChanged)
				if err := _GovernanceToken.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GovernanceToken *GovernanceTokenFilterer) ParseEIP712DomainChanged(log types.Log) (*GovernanceTokenEIP712DomainChanged, error) {
	event := new(GovernanceTokenEIP712DomainChanged)
	if err := _GovernanceToken.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GovernanceToken contract.
type GovernanceTokenInitializedIterator struct {
	Event *GovernanceTokenInitialized // Event containing the contract specifics and raw log

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
func (it *GovernanceTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTokenInitialized)
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
		it.Event = new(GovernanceTokenInitialized)
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
func (it *GovernanceTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTokenInitialized represents a Initialized event raised by the GovernanceToken contract.
type GovernanceTokenInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GovernanceToken *GovernanceTokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*GovernanceTokenInitializedIterator, error) {

	logs, sub, err := _GovernanceToken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenInitializedIterator{contract: _GovernanceToken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GovernanceToken *GovernanceTokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GovernanceTokenInitialized) (event.Subscription, error) {

	logs, sub, err := _GovernanceToken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTokenInitialized)
				if err := _GovernanceToken.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GovernanceToken *GovernanceTokenFilterer) ParseInitialized(log types.Log) (*GovernanceTokenInitialized, error) {
	event := new(GovernanceTokenInitialized)
	if err := _GovernanceToken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GovernanceToken contract.
type GovernanceTokenTransferIterator struct {
	Event *GovernanceTokenTransfer // Event containing the contract specifics and raw log

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
func (it *GovernanceTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTokenTransfer)
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
		it.Event = new(GovernanceTokenTransfer)
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
func (it *GovernanceTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTokenTransfer represents a Transfer event raised by the GovernanceToken contract.
type GovernanceTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GovernanceTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GovernanceToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenTransferIterator{contract: _GovernanceToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GovernanceTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GovernanceToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTokenTransfer)
				if err := _GovernanceToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) ParseTransfer(log types.Log) (*GovernanceTokenTransfer, error) {
	event := new(GovernanceTokenTransfer)
	if err := _GovernanceToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
