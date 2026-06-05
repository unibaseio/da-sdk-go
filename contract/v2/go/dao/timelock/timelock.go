// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package timelock

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

// DAOTimelockMetaData contains all meta data concerning the DAOTimelock contract.
var DAOTimelockMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"CANCELLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXECUTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROPOSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getMinDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperationState\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumTimelockControllerUpgradeable.OperationState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimestamp\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOperation\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hashOperationBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"minDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"executors\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperation\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationDone\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationPending\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationReady\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"schedule\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"scheduleBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallExecuted\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSalt\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallScheduled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Cancelled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinDelayChange\",\"inputs\":[{\"name\":\"oldDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimelockInsufficientDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TimelockInvalidOperationLength\",\"inputs\":[{\"name\":\"targets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payloads\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"values\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TimelockUnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TimelockUnexecutedPredecessor\",\"inputs\":[{\"name\":\"predecessorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"TimelockUnexpectedOperationState\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedStates\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x6080806040523461001657611b8c908161001b8239f35b5f80fdfe60406080815260048036101561001e575b5050361561001c575f80fd5b005b5f3560e01c90816301d5062a14610d9357816301ffc9a714610d3f57816307bd026514610d18578163134008d314610c6257816313bc9f2014610c44578163150b7a0214610bf2578163248a9ca314610bbc5781632ab0f52914610b9e5781632f2ff15d14610b7657816331d5075014610b5857816336568abe14610b14578163584b153e14610aed57816364d6235314610a795781637958004c14610a375781638065657f14610a165781638f2a0bb01461088d5781638f61f4f51461085357816391d1485414610803578163a217fddf146107e9578163b08e51c0146107af578163b1c5f42714610784578163bc197c8114610701578163c4c4c7b31461048e578163c4d252f5146103a9578163d45c443514610376578163d547741f1461032e578163e38335e5146101e8578163f23a6e6114610196575063f27a0c92146101695780610010565b34610192575f366003190112610192576020905f80516020611b17833981519152549051908152f35b5f80fd5b82346101925760a0366003190112610192576101b0610e3e565b506101b9610e54565b506084356001600160401b038111610192576020926101da91369101610f49565b505163f23a6e6160e01b8152f35b90506101f336610fbf565b9098949591939296975f80516020611b378339815191525f525f80516020611af7833981519152602052855f205f805260205260ff865f20541615610320575b838314801590610316575b6102e8575061025661025d918a868a878b888f611314565b98896115f6565b5f5b81811061026f5761001c896116af565b80808a7fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b588a8a6102df6102c78f988c6102c0828e6102ba8f60019f6102b591859161129b565b6112bf565b9761129b565b35956112d3565b906102d48282878761165c565b8d5194859485611157565b0390a30161025f565b85516001624fcdef60e01b031981529081019283526020830185905260408301849052918291506060010390fd5b508483141561023e565b6103293361156c565b610233565b823461019257806003193601126101925761001c91356103716001610351610e54565b93835f525f80516020611af78339815191526020525f20015433906115bd565b6119b9565b823461019257602036600319011261019257602091355f525f80516020611ad78339815191528252805f20549051908152f35b8234610192576020366003190112610192578135917ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f783805f525f80516020611af7833981519152602052825f20335f5260205260ff835f205416156104735750610412836111d9565b1561045757505f908282525f80516020611ad78339815191526020528120557fbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb705f80a2005b826044925191635ead8eb560e01b835282015260066024820152fd5b604492519163e2517d3f60e01b835233908301526024820152fd5b905034610192576080366003190112610192578035906001600160401b03602435818111610192576104c39036908401611097565b92604435828111610192576104db9036908501611097565b6064356001600160a01b0381811696909291878103610192577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0097885460ff818c1c1615978116988915806106fa575b6001809b1490816106f0575b1590816106e7575b506106d9575067ffffffffffffffff19811689178a558893929190886106ba575b50610569611a61565b610571611a61565b61057a306116d9565b506106aa575b5088939192905f845b610656575b5050825f905b61061d575b50505050817f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5925f80516020611b17833981519152558151905f82526020820152a16105e157005b7fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29260209268ff000000000000000019815416905551908152a1005b90919293825182101561064f5750806106428461063b879486611aa2565b51166118b6565b5001908392918994610594565b9392610599565b90809294939550518110156106a0578061067c86610675869486611aa2565b5116611763565b506106928661068b8386611aa2565b5116611810565b5001908293918a9593610589565b899492939161058e565b6106b3906116d9565b505f610580565b68ffffffffffffffffff191668010000000000000001178a555f610560565b8b5163f92ee8a960e01b8152fd5b9050155f61053f565b303b159150610537565b508861052b565b82346101925760a03660031901126101925761071b610e3e565b50610724610e54565b506001600160401b03604435818111610192576107449036908501611037565b506064358181116101925761075c9036908501611037565b506084359081116101925760209261077691369101610f49565b505163bc197c8160e01b8152f35b8234610192576020906107a861079936610fbf565b96959095949194939293611314565b9051908152f35b8234610192575f36600319011261019257602090517ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f7838152f35b8234610192575f36600319011261019257602090515f8152f35b823461019257806003193601126101925760209161081f610e54565b90355f525f80516020611af78339815191528352815f209060018060a01b03165f52825260ff815f20541690519015158152f35b8234610192575f36600319011261019257602090517fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc18152f35b82346101925760c0366003190112610192576001600160401b03908235828111610192576108be9036908501610f8f565b93602435848111610192576108d69036908301610f8f565b94604435908111610192576108ee9036908401610f8f565b60649391933590608435978960a43594610907336114ea565b828214801590610a0c575b6109df575089848489858a610927968e611314565b99610932858c61144f565b8a5f5b8a838210610973578c80925061094757005b7f20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d03879160209151908152a2005b6001927f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca8b8b6109d48f8c88978f92898f8f8f6109c2916109bc6102b58680946109c99961129b565b9a61129b565b35986112d3565b91519687968761111f565b0390a3018b90610935565b89516001624fcdef60e01b0319815290810191825260208201849052604082018390529081906060010390fd5b5083821415610912565b8234610192576020906107a8610a2b36610eab565b94939093929192611246565b823461019257602036600319011261019257610a538235611202565b90519082811015610a6657602092508152f35b602183634e487b7160e01b5f525260245ffd5b823461019257602036600319011261019257813591303303610ad757507f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d55f80516020611b178339815191529182548151908152846020820152a155005b602491519063e2850c5960e01b82523390820152fd5b823461019257602036600319011261019257610b0b602092356111d9565b90519015158152f35b8234610192578060031936011261019257610b2d610e54565b90336001600160a01b03831603610b49575061001c91356119b9565b5163334bd91960e11b81529050fd5b823461019257602036600319011261019257610b0b602092356111c2565b823461019257806003193601126101925761001c9135610b996001610351610e54565b611949565b823461019257602036600319011261019257610b0b602092356111aa565b823461019257602036600319011261019257602091355f525f80516020611af783398151915282526001815f2001549051908152f35b823461019257608036600319011261019257610c0c610e3e565b50610c15610e54565b506064356001600160401b03811161019257602092610c3691369101610f49565b5051630a85bd0160e11b8152f35b823461019257602036600319011261019257610b0b6020923561117e565b61001c610cec5f610d027fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58610ce388610c9a36610eab565b5f80516020611b378339815191528b9a9697939598929a525f80516020611af7833981519152602052828b208b805260205260ff838c20541615610d0a575b8985858a8a611246565b998a98896115f6565b610cf88383888861165c565b5194859485611157565b0390a36116af565b610d133361156c565b610cd9565b8234610192575f36600319011261019257602090515f80516020611b378339815191528152f35b905034610192576020366003190112610192573563ffffffff60e01b811680910361019257602091630271189760e51b8214918215610d82575b50519015158152f35b6301ffc9a760e01b1491505f610d79565b82346101925760c036600319011261019257610dad610e3e565b9060243592604435936001600160401b03851161019257610df35f927f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca96369101610e7e565b95909160643595610e346084359760a43590610e0e336114ea565b610e1c8a828d8a8989611246565b9a8b97610e29848a61144f565b8a519687968761111f565b0390a38161094757005b600435906001600160a01b038216820361019257565b602435906001600160a01b038216820361019257565b35906001600160a01b038216820361019257565b9181601f84011215610192578235916001600160401b038311610192576020838186019501011161019257565b60a0600319820112610192576004356001600160a01b0381168103610192579160243591604435906001600160401b03821161019257610eed91600401610e7e565b90916064359060843590565b90601f801991011681019081106001600160401b03821117610f1a57604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b038111610f1a57601f01601f191660200190565b81601f8201121561019257803590610f6082610f2e565b92610f6e6040519485610ef9565b8284526020838301011161019257815f926020809301838601378301015290565b9181601f84011215610192578235916001600160401b038311610192576020808501948460051b01011161019257565b9060a0600319830112610192576001600160401b036004358181116101925783610feb91600401610f8f565b93909392602435838111610192578261100691600401610f8f565b9390939260443591821161019257610eed91600401610f8f565b6001600160401b038111610f1a5760051b60200190565b9080601f8301121561019257602090823561105181611020565b9361105f6040519586610ef9565b81855260208086019260051b82010192831161019257602001905b828210611088575050505090565b8135815290830190830161107a565b9080601f830112156101925760209082356110b181611020565b936110bf6040519586610ef9565b81855260208086019260051b82010192831161019257602001905b8282106110e8575050505090565b8380916110f484610e6a565b8152019101906110da565b908060209392818452848401375f828201840152601f01601f1916010190565b92909361114d926080959897969860018060a01b03168552602085015260a0604085015260a08401916110ff565b9460608201520152565b61117b949260609260018060a01b03168252602082015281604082015201916110ff565b90565b61118790611202565b60048110156111965760021490565b634e487b7160e01b5f52602160045260245ffd5b6111b390611202565b60048110156111965760031490565b6111cb90611202565b600481101561119657151590565b6111e290611202565b600481101561119657600181149081156111fa575090565b600291501490565b5f525f80516020611ad783398151915260205260405f205480155f1461122757505f90565b600181036112355750600390565b42101561124157600190565b600290565b9461127c61129594959293604051968795602087019960018060a01b03168a52604087015260a0606087015260c08601916110ff565b91608084015260a083015203601f198101835282610ef9565b51902090565b91908110156112ab5760051b0190565b634e487b7160e01b5f52603260045260245ffd5b356001600160a01b03811681036101925790565b91908110156112ab5760051b81013590601e19813603018212156101925701908135916001600160401b038311610192576020018236038113610192579190565b969294909695919560405196602091828901998060c08b0160a08d525260e08a0191905f5b8582821061142b5750505050888103601f1990810160408b0152888252976001600160fb1b038111610192579089969495939897929160051b80928a830137019380888601878703606089015252604085019460408260051b82010195835f925b8484106113c2575050505050506112959550608084015260a083015203908101835282610ef9565b9193969850919398999496603f198282030184528935601e19843603018112156101925783018681019190356001600160401b03811161019257803603831361019257611414889283926001956110ff565b9b0194019401918b98969394919a9997959a61139a565b80600192939495838060a01b0361144188610e6a565b168152019401929101611339565b90611459826111c2565b6114ca575f80516020611b17833981519152548082106114ac5750420190814211611498575f525f80516020611ad783398151915260205260405f2055565b634e487b7160e01b5f52601160045260245ffd5b6044925060405191635433660960e01b835260048301526024820152fd5b604051635ead8eb560e01b81526004810183905260016024820152604490fd5b6001600160a01b03165f8181527f5a8734c34b98d7c96eb2ea25f298989407e1f25da116ec139bcce0887bcb7cf760205260409020547fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc19060ff161561154e575050565b604492506040519163e2517d3f60e01b835260048301526024820152fd5b6001600160a01b03165f8181527f52fce5e8a5d0d9e8d1ea29f4525e512e9c27bf92cae50374d497f918ab48f38260205260409020545f80516020611b378339815191529060ff161561154e575050565b805f525f80516020611af783398151915260205260405f209160018060a01b031691825f5260205260ff60405f2054161561154e575050565b6115ff8161117e565b1561163d57508015158061162d575b6116155750565b6024906040519063121534c360e31b82526004820152fd5b50611637816111aa565b1561160e565b60449060405190635ead8eb560e01b8252600482015260046024820152fd5b6116a4935f93928493826040519384928337810185815203925af13d156116a7573d9061168882610f2e565b916116966040519384610ef9565b82523d5f602084013e611a38565b50565b606090611a38565b6116b88161117e565b1561163d575f525f80516020611ad7833981519152602052600160405f2055565b6001600160a01b03165f8181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f80516020611af78339815191529060ff1661175d575f805260205260405f20815f5260205260405f20600160ff1982541617905533905f5f80516020611ab78339815191528180a4600190565b50505f90565b6001600160a01b03165f8181527f5a8734c34b98d7c96eb2ea25f298989407e1f25da116ec139bcce0887bcb7cf760205260409020547fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc191905f80516020611af78339815191529060ff1661180957825f5260205260405f20815f5260205260405f20600160ff1982541617905533915f80516020611ab78339815191525f80a4600190565b5050505f90565b6001600160a01b03165f8181527ffa71e07f24c4701ef65a970775979de1292cfe909335cd18a32d2b7b7398791460205260409020547ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f78391905f80516020611af78339815191529060ff1661180957825f5260205260405f20815f5260205260405f20600160ff1982541617905533915f80516020611ab78339815191525f80a4600190565b6001600160a01b03165f8181527f52fce5e8a5d0d9e8d1ea29f4525e512e9c27bf92cae50374d497f918ab48f38260205260409020545f80516020611b3783398151915291905f80516020611af78339815191529060ff1661180957825f5260205260405f20815f5260205260405f20600160ff1982541617905533915f80516020611ab78339815191525f80a4600190565b90815f525f80516020611af78339815191528060205260405f209160018060a01b031691825f5260205260ff60405f205416155f1461180957825f5260205260405f20815f5260205260405f20600160ff1982541617905533915f80516020611ab78339815191525f80a4600190565b90815f525f80516020611af78339815191528060205260405f209160018060a01b031691825f5260205260ff60405f2054165f1461180957825f5260205260405f20815f5260205260405f2060ff19815416905533917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b15611a405790565b805115611a4f57602081519101fd5b60405163d6bda27560e01b8152600490fd5b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c1615611a9057565b604051631afcd79f60e31b8152600490fd5b80518210156112ab5760209160051b01019056fe2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9a37c2aa9d186a0969ff8a8267bf4e07e864c2f2768f5040949e28a624fb360002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268009a37c2aa9d186a0969ff8a8267bf4e07e864c2f2768f5040949e28a624fb3601d8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63a264697066735822122024bfd1c932b3f77bb7da7da1e085ee69b2a3d41632ccd74897cd54ce2e0059fb64736f6c63430008180033",
}

// DAOTimelockABI is the input ABI used to generate the binding from.
// Deprecated: Use DAOTimelockMetaData.ABI instead.
var DAOTimelockABI = DAOTimelockMetaData.ABI

// DAOTimelockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DAOTimelockMetaData.Bin instead.
var DAOTimelockBin = DAOTimelockMetaData.Bin

// DeployDAOTimelock deploys a new Ethereum contract, binding an instance of DAOTimelock to it.
func DeployDAOTimelock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DAOTimelock, error) {
	parsed, err := DAOTimelockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DAOTimelockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DAOTimelock{DAOTimelockCaller: DAOTimelockCaller{contract: contract}, DAOTimelockTransactor: DAOTimelockTransactor{contract: contract}, DAOTimelockFilterer: DAOTimelockFilterer{contract: contract}}, nil
}

// DAOTimelock is an auto generated Go binding around an Ethereum contract.
type DAOTimelock struct {
	DAOTimelockCaller     // Read-only binding to the contract
	DAOTimelockTransactor // Write-only binding to the contract
	DAOTimelockFilterer   // Log filterer for contract events
}

// DAOTimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAOTimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOTimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAOTimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOTimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAOTimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOTimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAOTimelockSession struct {
	Contract     *DAOTimelock      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAOTimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAOTimelockCallerSession struct {
	Contract *DAOTimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DAOTimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAOTimelockTransactorSession struct {
	Contract     *DAOTimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DAOTimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAOTimelockRaw struct {
	Contract *DAOTimelock // Generic contract binding to access the raw methods on
}

// DAOTimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAOTimelockCallerRaw struct {
	Contract *DAOTimelockCaller // Generic read-only contract binding to access the raw methods on
}

// DAOTimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAOTimelockTransactorRaw struct {
	Contract *DAOTimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAOTimelock creates a new instance of DAOTimelock, bound to a specific deployed contract.
func NewDAOTimelock(address common.Address, backend bind.ContractBackend) (*DAOTimelock, error) {
	contract, err := bindDAOTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAOTimelock{DAOTimelockCaller: DAOTimelockCaller{contract: contract}, DAOTimelockTransactor: DAOTimelockTransactor{contract: contract}, DAOTimelockFilterer: DAOTimelockFilterer{contract: contract}}, nil
}

// NewDAOTimelockCaller creates a new read-only instance of DAOTimelock, bound to a specific deployed contract.
func NewDAOTimelockCaller(address common.Address, caller bind.ContractCaller) (*DAOTimelockCaller, error) {
	contract, err := bindDAOTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAOTimelockCaller{contract: contract}, nil
}

// NewDAOTimelockTransactor creates a new write-only instance of DAOTimelock, bound to a specific deployed contract.
func NewDAOTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*DAOTimelockTransactor, error) {
	contract, err := bindDAOTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAOTimelockTransactor{contract: contract}, nil
}

// NewDAOTimelockFilterer creates a new log filterer instance of DAOTimelock, bound to a specific deployed contract.
func NewDAOTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*DAOTimelockFilterer, error) {
	contract, err := bindDAOTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAOTimelockFilterer{contract: contract}, nil
}

// bindDAOTimelock binds a generic wrapper to an already deployed contract.
func bindDAOTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DAOTimelockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAOTimelock *DAOTimelockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAOTimelock.Contract.DAOTimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAOTimelock *DAOTimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOTimelock.Contract.DAOTimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAOTimelock *DAOTimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAOTimelock.Contract.DAOTimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAOTimelock *DAOTimelockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAOTimelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAOTimelock *DAOTimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOTimelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAOTimelock *DAOTimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAOTimelock.Contract.contract.Transact(opts, method, params...)
}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockCaller) CANCELLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "CANCELLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockSession) CANCELLERROLE() ([32]byte, error) {
	return _DAOTimelock.Contract.CANCELLERROLE(&_DAOTimelock.CallOpts)
}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockCallerSession) CANCELLERROLE() ([32]byte, error) {
	return _DAOTimelock.Contract.CANCELLERROLE(&_DAOTimelock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DAOTimelock.Contract.DEFAULTADMINROLE(&_DAOTimelock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DAOTimelock.Contract.DEFAULTADMINROLE(&_DAOTimelock.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockSession) EXECUTORROLE() ([32]byte, error) {
	return _DAOTimelock.Contract.EXECUTORROLE(&_DAOTimelock.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _DAOTimelock.Contract.EXECUTORROLE(&_DAOTimelock.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockCaller) PROPOSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "PROPOSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockSession) PROPOSERROLE() ([32]byte, error) {
	return _DAOTimelock.Contract.PROPOSERROLE(&_DAOTimelock.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_DAOTimelock *DAOTimelockCallerSession) PROPOSERROLE() ([32]byte, error) {
	return _DAOTimelock.Contract.PROPOSERROLE(&_DAOTimelock.CallOpts)
}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_DAOTimelock *DAOTimelockCaller) GetMinDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "getMinDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_DAOTimelock *DAOTimelockSession) GetMinDelay() (*big.Int, error) {
	return _DAOTimelock.Contract.GetMinDelay(&_DAOTimelock.CallOpts)
}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_DAOTimelock *DAOTimelockCallerSession) GetMinDelay() (*big.Int, error) {
	return _DAOTimelock.Contract.GetMinDelay(&_DAOTimelock.CallOpts)
}

// GetOperationState is a free data retrieval call binding the contract method 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (_DAOTimelock *DAOTimelockCaller) GetOperationState(opts *bind.CallOpts, id [32]byte) (uint8, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "getOperationState", id)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperationState is a free data retrieval call binding the contract method 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (_DAOTimelock *DAOTimelockSession) GetOperationState(id [32]byte) (uint8, error) {
	return _DAOTimelock.Contract.GetOperationState(&_DAOTimelock.CallOpts, id)
}

// GetOperationState is a free data retrieval call binding the contract method 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (_DAOTimelock *DAOTimelockCallerSession) GetOperationState(id [32]byte) (uint8, error) {
	return _DAOTimelock.Contract.GetOperationState(&_DAOTimelock.CallOpts, id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DAOTimelock *DAOTimelockCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DAOTimelock *DAOTimelockSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DAOTimelock.Contract.GetRoleAdmin(&_DAOTimelock.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DAOTimelock *DAOTimelockCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DAOTimelock.Contract.GetRoleAdmin(&_DAOTimelock.CallOpts, role)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_DAOTimelock *DAOTimelockCaller) GetTimestamp(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "getTimestamp", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_DAOTimelock *DAOTimelockSession) GetTimestamp(id [32]byte) (*big.Int, error) {
	return _DAOTimelock.Contract.GetTimestamp(&_DAOTimelock.CallOpts, id)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_DAOTimelock *DAOTimelockCallerSession) GetTimestamp(id [32]byte) (*big.Int, error) {
	return _DAOTimelock.Contract.GetTimestamp(&_DAOTimelock.CallOpts, id)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DAOTimelock *DAOTimelockCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DAOTimelock *DAOTimelockSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DAOTimelock.Contract.HasRole(&_DAOTimelock.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DAOTimelock *DAOTimelockCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DAOTimelock.Contract.HasRole(&_DAOTimelock.CallOpts, role, account)
}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_DAOTimelock *DAOTimelockCaller) HashOperation(opts *bind.CallOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "hashOperation", target, value, data, predecessor, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_DAOTimelock *DAOTimelockSession) HashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _DAOTimelock.Contract.HashOperation(&_DAOTimelock.CallOpts, target, value, data, predecessor, salt)
}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_DAOTimelock *DAOTimelockCallerSession) HashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _DAOTimelock.Contract.HashOperation(&_DAOTimelock.CallOpts, target, value, data, predecessor, salt)
}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_DAOTimelock *DAOTimelockCaller) HashOperationBatch(opts *bind.CallOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "hashOperationBatch", targets, values, payloads, predecessor, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_DAOTimelock *DAOTimelockSession) HashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _DAOTimelock.Contract.HashOperationBatch(&_DAOTimelock.CallOpts, targets, values, payloads, predecessor, salt)
}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_DAOTimelock *DAOTimelockCallerSession) HashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _DAOTimelock.Contract.HashOperationBatch(&_DAOTimelock.CallOpts, targets, values, payloads, predecessor, salt)
}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockCaller) IsOperation(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "isOperation", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockSession) IsOperation(id [32]byte) (bool, error) {
	return _DAOTimelock.Contract.IsOperation(&_DAOTimelock.CallOpts, id)
}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockCallerSession) IsOperation(id [32]byte) (bool, error) {
	return _DAOTimelock.Contract.IsOperation(&_DAOTimelock.CallOpts, id)
}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockCaller) IsOperationDone(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "isOperationDone", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockSession) IsOperationDone(id [32]byte) (bool, error) {
	return _DAOTimelock.Contract.IsOperationDone(&_DAOTimelock.CallOpts, id)
}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockCallerSession) IsOperationDone(id [32]byte) (bool, error) {
	return _DAOTimelock.Contract.IsOperationDone(&_DAOTimelock.CallOpts, id)
}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockCaller) IsOperationPending(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "isOperationPending", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockSession) IsOperationPending(id [32]byte) (bool, error) {
	return _DAOTimelock.Contract.IsOperationPending(&_DAOTimelock.CallOpts, id)
}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockCallerSession) IsOperationPending(id [32]byte) (bool, error) {
	return _DAOTimelock.Contract.IsOperationPending(&_DAOTimelock.CallOpts, id)
}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockCaller) IsOperationReady(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "isOperationReady", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockSession) IsOperationReady(id [32]byte) (bool, error) {
	return _DAOTimelock.Contract.IsOperationReady(&_DAOTimelock.CallOpts, id)
}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_DAOTimelock *DAOTimelockCallerSession) IsOperationReady(id [32]byte) (bool, error) {
	return _DAOTimelock.Contract.IsOperationReady(&_DAOTimelock.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAOTimelock *DAOTimelockCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DAOTimelock.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAOTimelock *DAOTimelockSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DAOTimelock.Contract.SupportsInterface(&_DAOTimelock.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAOTimelock *DAOTimelockCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DAOTimelock.Contract.SupportsInterface(&_DAOTimelock.CallOpts, interfaceId)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_DAOTimelock *DAOTimelockTransactor) Cancel(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "cancel", id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_DAOTimelock *DAOTimelockSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.Cancel(&_DAOTimelock.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_DAOTimelock *DAOTimelockTransactorSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.Cancel(&_DAOTimelock.TransactOpts, id)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_DAOTimelock *DAOTimelockTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "execute", target, value, payload, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_DAOTimelock *DAOTimelockSession) Execute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.Execute(&_DAOTimelock.TransactOpts, target, value, payload, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_DAOTimelock *DAOTimelockTransactorSession) Execute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.Execute(&_DAOTimelock.TransactOpts, target, value, payload, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_DAOTimelock *DAOTimelockTransactor) ExecuteBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "executeBatch", targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_DAOTimelock *DAOTimelockSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.ExecuteBatch(&_DAOTimelock.TransactOpts, targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_DAOTimelock *DAOTimelockTransactorSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.ExecuteBatch(&_DAOTimelock.TransactOpts, targets, values, payloads, predecessor, salt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DAOTimelock *DAOTimelockTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DAOTimelock *DAOTimelockSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOTimelock.Contract.GrantRole(&_DAOTimelock.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DAOTimelock *DAOTimelockTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOTimelock.Contract.GrantRole(&_DAOTimelock.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4c4c7b3.
//
// Solidity: function initialize(uint256 minDelay, address[] proposers, address[] executors, address admin) returns()
func (_DAOTimelock *DAOTimelockTransactor) Initialize(opts *bind.TransactOpts, minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "initialize", minDelay, proposers, executors, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4c4c7b3.
//
// Solidity: function initialize(uint256 minDelay, address[] proposers, address[] executors, address admin) returns()
func (_DAOTimelock *DAOTimelockSession) Initialize(minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) (*types.Transaction, error) {
	return _DAOTimelock.Contract.Initialize(&_DAOTimelock.TransactOpts, minDelay, proposers, executors, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4c4c7b3.
//
// Solidity: function initialize(uint256 minDelay, address[] proposers, address[] executors, address admin) returns()
func (_DAOTimelock *DAOTimelockTransactorSession) Initialize(minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) (*types.Transaction, error) {
	return _DAOTimelock.Contract.Initialize(&_DAOTimelock.TransactOpts, minDelay, proposers, executors, admin)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_DAOTimelock *DAOTimelockTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_DAOTimelock *DAOTimelockSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.OnERC1155BatchReceived(&_DAOTimelock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_DAOTimelock *DAOTimelockTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.OnERC1155BatchReceived(&_DAOTimelock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_DAOTimelock *DAOTimelockTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_DAOTimelock *DAOTimelockSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.OnERC1155Received(&_DAOTimelock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_DAOTimelock *DAOTimelockTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.OnERC1155Received(&_DAOTimelock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DAOTimelock *DAOTimelockTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DAOTimelock *DAOTimelockSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.OnERC721Received(&_DAOTimelock.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DAOTimelock *DAOTimelockTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DAOTimelock.Contract.OnERC721Received(&_DAOTimelock.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DAOTimelock *DAOTimelockTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DAOTimelock *DAOTimelockSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DAOTimelock.Contract.RenounceRole(&_DAOTimelock.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DAOTimelock *DAOTimelockTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DAOTimelock.Contract.RenounceRole(&_DAOTimelock.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DAOTimelock *DAOTimelockTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DAOTimelock *DAOTimelockSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOTimelock.Contract.RevokeRole(&_DAOTimelock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DAOTimelock *DAOTimelockTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAOTimelock.Contract.RevokeRole(&_DAOTimelock.TransactOpts, role, account)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_DAOTimelock *DAOTimelockTransactor) Schedule(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "schedule", target, value, data, predecessor, salt, delay)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_DAOTimelock *DAOTimelockSession) Schedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _DAOTimelock.Contract.Schedule(&_DAOTimelock.TransactOpts, target, value, data, predecessor, salt, delay)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_DAOTimelock *DAOTimelockTransactorSession) Schedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _DAOTimelock.Contract.Schedule(&_DAOTimelock.TransactOpts, target, value, data, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_DAOTimelock *DAOTimelockTransactor) ScheduleBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "scheduleBatch", targets, values, payloads, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_DAOTimelock *DAOTimelockSession) ScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _DAOTimelock.Contract.ScheduleBatch(&_DAOTimelock.TransactOpts, targets, values, payloads, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_DAOTimelock *DAOTimelockTransactorSession) ScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _DAOTimelock.Contract.ScheduleBatch(&_DAOTimelock.TransactOpts, targets, values, payloads, predecessor, salt, delay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_DAOTimelock *DAOTimelockTransactor) UpdateDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _DAOTimelock.contract.Transact(opts, "updateDelay", newDelay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_DAOTimelock *DAOTimelockSession) UpdateDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _DAOTimelock.Contract.UpdateDelay(&_DAOTimelock.TransactOpts, newDelay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_DAOTimelock *DAOTimelockTransactorSession) UpdateDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _DAOTimelock.Contract.UpdateDelay(&_DAOTimelock.TransactOpts, newDelay)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAOTimelock *DAOTimelockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOTimelock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAOTimelock *DAOTimelockSession) Receive() (*types.Transaction, error) {
	return _DAOTimelock.Contract.Receive(&_DAOTimelock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAOTimelock *DAOTimelockTransactorSession) Receive() (*types.Transaction, error) {
	return _DAOTimelock.Contract.Receive(&_DAOTimelock.TransactOpts)
}

// DAOTimelockCallExecutedIterator is returned from FilterCallExecuted and is used to iterate over the raw logs and unpacked data for CallExecuted events raised by the DAOTimelock contract.
type DAOTimelockCallExecutedIterator struct {
	Event *DAOTimelockCallExecuted // Event containing the contract specifics and raw log

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
func (it *DAOTimelockCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTimelockCallExecuted)
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
		it.Event = new(DAOTimelockCallExecuted)
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
func (it *DAOTimelockCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTimelockCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTimelockCallExecuted represents a CallExecuted event raised by the DAOTimelock contract.
type DAOTimelockCallExecuted struct {
	Id     [32]byte
	Index  *big.Int
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallExecuted is a free log retrieval operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_DAOTimelock *DAOTimelockFilterer) FilterCallExecuted(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*DAOTimelockCallExecutedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _DAOTimelock.contract.FilterLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &DAOTimelockCallExecutedIterator{contract: _DAOTimelock.contract, event: "CallExecuted", logs: logs, sub: sub}, nil
}

// WatchCallExecuted is a free log subscription operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_DAOTimelock *DAOTimelockFilterer) WatchCallExecuted(opts *bind.WatchOpts, sink chan<- *DAOTimelockCallExecuted, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _DAOTimelock.contract.WatchLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTimelockCallExecuted)
				if err := _DAOTimelock.contract.UnpackLog(event, "CallExecuted", log); err != nil {
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

// ParseCallExecuted is a log parse operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_DAOTimelock *DAOTimelockFilterer) ParseCallExecuted(log types.Log) (*DAOTimelockCallExecuted, error) {
	event := new(DAOTimelockCallExecuted)
	if err := _DAOTimelock.contract.UnpackLog(event, "CallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOTimelockCallSaltIterator is returned from FilterCallSalt and is used to iterate over the raw logs and unpacked data for CallSalt events raised by the DAOTimelock contract.
type DAOTimelockCallSaltIterator struct {
	Event *DAOTimelockCallSalt // Event containing the contract specifics and raw log

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
func (it *DAOTimelockCallSaltIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTimelockCallSalt)
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
		it.Event = new(DAOTimelockCallSalt)
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
func (it *DAOTimelockCallSaltIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTimelockCallSaltIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTimelockCallSalt represents a CallSalt event raised by the DAOTimelock contract.
type DAOTimelockCallSalt struct {
	Id   [32]byte
	Salt [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCallSalt is a free log retrieval operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_DAOTimelock *DAOTimelockFilterer) FilterCallSalt(opts *bind.FilterOpts, id [][32]byte) (*DAOTimelockCallSaltIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DAOTimelock.contract.FilterLogs(opts, "CallSalt", idRule)
	if err != nil {
		return nil, err
	}
	return &DAOTimelockCallSaltIterator{contract: _DAOTimelock.contract, event: "CallSalt", logs: logs, sub: sub}, nil
}

// WatchCallSalt is a free log subscription operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_DAOTimelock *DAOTimelockFilterer) WatchCallSalt(opts *bind.WatchOpts, sink chan<- *DAOTimelockCallSalt, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DAOTimelock.contract.WatchLogs(opts, "CallSalt", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTimelockCallSalt)
				if err := _DAOTimelock.contract.UnpackLog(event, "CallSalt", log); err != nil {
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

// ParseCallSalt is a log parse operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_DAOTimelock *DAOTimelockFilterer) ParseCallSalt(log types.Log) (*DAOTimelockCallSalt, error) {
	event := new(DAOTimelockCallSalt)
	if err := _DAOTimelock.contract.UnpackLog(event, "CallSalt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOTimelockCallScheduledIterator is returned from FilterCallScheduled and is used to iterate over the raw logs and unpacked data for CallScheduled events raised by the DAOTimelock contract.
type DAOTimelockCallScheduledIterator struct {
	Event *DAOTimelockCallScheduled // Event containing the contract specifics and raw log

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
func (it *DAOTimelockCallScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTimelockCallScheduled)
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
		it.Event = new(DAOTimelockCallScheduled)
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
func (it *DAOTimelockCallScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTimelockCallScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTimelockCallScheduled represents a CallScheduled event raised by the DAOTimelock contract.
type DAOTimelockCallScheduled struct {
	Id          [32]byte
	Index       *big.Int
	Target      common.Address
	Value       *big.Int
	Data        []byte
	Predecessor [32]byte
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCallScheduled is a free log retrieval operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_DAOTimelock *DAOTimelockFilterer) FilterCallScheduled(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*DAOTimelockCallScheduledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _DAOTimelock.contract.FilterLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &DAOTimelockCallScheduledIterator{contract: _DAOTimelock.contract, event: "CallScheduled", logs: logs, sub: sub}, nil
}

// WatchCallScheduled is a free log subscription operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_DAOTimelock *DAOTimelockFilterer) WatchCallScheduled(opts *bind.WatchOpts, sink chan<- *DAOTimelockCallScheduled, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _DAOTimelock.contract.WatchLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTimelockCallScheduled)
				if err := _DAOTimelock.contract.UnpackLog(event, "CallScheduled", log); err != nil {
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

// ParseCallScheduled is a log parse operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_DAOTimelock *DAOTimelockFilterer) ParseCallScheduled(log types.Log) (*DAOTimelockCallScheduled, error) {
	event := new(DAOTimelockCallScheduled)
	if err := _DAOTimelock.contract.UnpackLog(event, "CallScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOTimelockCancelledIterator is returned from FilterCancelled and is used to iterate over the raw logs and unpacked data for Cancelled events raised by the DAOTimelock contract.
type DAOTimelockCancelledIterator struct {
	Event *DAOTimelockCancelled // Event containing the contract specifics and raw log

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
func (it *DAOTimelockCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTimelockCancelled)
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
		it.Event = new(DAOTimelockCancelled)
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
func (it *DAOTimelockCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTimelockCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTimelockCancelled represents a Cancelled event raised by the DAOTimelock contract.
type DAOTimelockCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelled is a free log retrieval operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_DAOTimelock *DAOTimelockFilterer) FilterCancelled(opts *bind.FilterOpts, id [][32]byte) (*DAOTimelockCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DAOTimelock.contract.FilterLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &DAOTimelockCancelledIterator{contract: _DAOTimelock.contract, event: "Cancelled", logs: logs, sub: sub}, nil
}

// WatchCancelled is a free log subscription operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_DAOTimelock *DAOTimelockFilterer) WatchCancelled(opts *bind.WatchOpts, sink chan<- *DAOTimelockCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DAOTimelock.contract.WatchLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTimelockCancelled)
				if err := _DAOTimelock.contract.UnpackLog(event, "Cancelled", log); err != nil {
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

// ParseCancelled is a log parse operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_DAOTimelock *DAOTimelockFilterer) ParseCancelled(log types.Log) (*DAOTimelockCancelled, error) {
	event := new(DAOTimelockCancelled)
	if err := _DAOTimelock.contract.UnpackLog(event, "Cancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOTimelockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DAOTimelock contract.
type DAOTimelockInitializedIterator struct {
	Event *DAOTimelockInitialized // Event containing the contract specifics and raw log

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
func (it *DAOTimelockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTimelockInitialized)
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
		it.Event = new(DAOTimelockInitialized)
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
func (it *DAOTimelockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTimelockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTimelockInitialized represents a Initialized event raised by the DAOTimelock contract.
type DAOTimelockInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DAOTimelock *DAOTimelockFilterer) FilterInitialized(opts *bind.FilterOpts) (*DAOTimelockInitializedIterator, error) {

	logs, sub, err := _DAOTimelock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DAOTimelockInitializedIterator{contract: _DAOTimelock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DAOTimelock *DAOTimelockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DAOTimelockInitialized) (event.Subscription, error) {

	logs, sub, err := _DAOTimelock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTimelockInitialized)
				if err := _DAOTimelock.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DAOTimelock *DAOTimelockFilterer) ParseInitialized(log types.Log) (*DAOTimelockInitialized, error) {
	event := new(DAOTimelockInitialized)
	if err := _DAOTimelock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOTimelockMinDelayChangeIterator is returned from FilterMinDelayChange and is used to iterate over the raw logs and unpacked data for MinDelayChange events raised by the DAOTimelock contract.
type DAOTimelockMinDelayChangeIterator struct {
	Event *DAOTimelockMinDelayChange // Event containing the contract specifics and raw log

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
func (it *DAOTimelockMinDelayChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTimelockMinDelayChange)
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
		it.Event = new(DAOTimelockMinDelayChange)
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
func (it *DAOTimelockMinDelayChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTimelockMinDelayChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTimelockMinDelayChange represents a MinDelayChange event raised by the DAOTimelock contract.
type DAOTimelockMinDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinDelayChange is a free log retrieval operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_DAOTimelock *DAOTimelockFilterer) FilterMinDelayChange(opts *bind.FilterOpts) (*DAOTimelockMinDelayChangeIterator, error) {

	logs, sub, err := _DAOTimelock.contract.FilterLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return &DAOTimelockMinDelayChangeIterator{contract: _DAOTimelock.contract, event: "MinDelayChange", logs: logs, sub: sub}, nil
}

// WatchMinDelayChange is a free log subscription operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_DAOTimelock *DAOTimelockFilterer) WatchMinDelayChange(opts *bind.WatchOpts, sink chan<- *DAOTimelockMinDelayChange) (event.Subscription, error) {

	logs, sub, err := _DAOTimelock.contract.WatchLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTimelockMinDelayChange)
				if err := _DAOTimelock.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
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

// ParseMinDelayChange is a log parse operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_DAOTimelock *DAOTimelockFilterer) ParseMinDelayChange(log types.Log) (*DAOTimelockMinDelayChange, error) {
	event := new(DAOTimelockMinDelayChange)
	if err := _DAOTimelock.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOTimelockRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DAOTimelock contract.
type DAOTimelockRoleAdminChangedIterator struct {
	Event *DAOTimelockRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DAOTimelockRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTimelockRoleAdminChanged)
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
		it.Event = new(DAOTimelockRoleAdminChanged)
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
func (it *DAOTimelockRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTimelockRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTimelockRoleAdminChanged represents a RoleAdminChanged event raised by the DAOTimelock contract.
type DAOTimelockRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DAOTimelock *DAOTimelockFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DAOTimelockRoleAdminChangedIterator, error) {

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

	logs, sub, err := _DAOTimelock.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DAOTimelockRoleAdminChangedIterator{contract: _DAOTimelock.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DAOTimelock *DAOTimelockFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DAOTimelockRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _DAOTimelock.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTimelockRoleAdminChanged)
				if err := _DAOTimelock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_DAOTimelock *DAOTimelockFilterer) ParseRoleAdminChanged(log types.Log) (*DAOTimelockRoleAdminChanged, error) {
	event := new(DAOTimelockRoleAdminChanged)
	if err := _DAOTimelock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOTimelockRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DAOTimelock contract.
type DAOTimelockRoleGrantedIterator struct {
	Event *DAOTimelockRoleGranted // Event containing the contract specifics and raw log

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
func (it *DAOTimelockRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTimelockRoleGranted)
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
		it.Event = new(DAOTimelockRoleGranted)
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
func (it *DAOTimelockRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTimelockRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTimelockRoleGranted represents a RoleGranted event raised by the DAOTimelock contract.
type DAOTimelockRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAOTimelock *DAOTimelockFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DAOTimelockRoleGrantedIterator, error) {

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

	logs, sub, err := _DAOTimelock.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DAOTimelockRoleGrantedIterator{contract: _DAOTimelock.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAOTimelock *DAOTimelockFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DAOTimelockRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DAOTimelock.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTimelockRoleGranted)
				if err := _DAOTimelock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_DAOTimelock *DAOTimelockFilterer) ParseRoleGranted(log types.Log) (*DAOTimelockRoleGranted, error) {
	event := new(DAOTimelockRoleGranted)
	if err := _DAOTimelock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOTimelockRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DAOTimelock contract.
type DAOTimelockRoleRevokedIterator struct {
	Event *DAOTimelockRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DAOTimelockRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTimelockRoleRevoked)
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
		it.Event = new(DAOTimelockRoleRevoked)
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
func (it *DAOTimelockRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTimelockRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTimelockRoleRevoked represents a RoleRevoked event raised by the DAOTimelock contract.
type DAOTimelockRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAOTimelock *DAOTimelockFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DAOTimelockRoleRevokedIterator, error) {

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

	logs, sub, err := _DAOTimelock.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DAOTimelockRoleRevokedIterator{contract: _DAOTimelock.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAOTimelock *DAOTimelockFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DAOTimelockRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DAOTimelock.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTimelockRoleRevoked)
				if err := _DAOTimelock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_DAOTimelock *DAOTimelockFilterer) ParseRoleRevoked(log types.Log) (*DAOTimelockRoleRevoked, error) {
	event := new(DAOTimelockRoleRevoked)
	if err := _DAOTimelock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
