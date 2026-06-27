// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vub

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

// VUBMetaData contains all meta data concerning the VUB contract.
var VUBMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REWARD_FUNDER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"boostBps\",\"inputs\":[{\"name\":\"dur\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pos\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCheckpoints.Checkpoint208\",\"components\":[{\"name\":\"_key\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"_value\",\"type\":\"uint208\",\"internalType\":\"uint208\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateBySig\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegates\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"earned\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"extendLock\",\"inputs\":[{\"name\":\"newDur\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPastTotalSupply\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPastVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReward\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_ub\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastTimeRewardApplicable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastUpdate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"locks\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxBoostBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxLock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minLock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"notifyReward\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pending\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ready\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"periodFinish\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardPerToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardPerTokenStored\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setParams\",\"inputs\":[{\"name\":\"_minLock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_maxLock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_maxBoostBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_cooldown\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardToken\",\"inputs\":[{\"name\":\"_t\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dur\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ub\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"userRewardPerTokenPaid\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateChanged\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fromDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateVotesChanged\",\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"previousVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockExtended\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newEnd\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"vubMinted\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParamsUpdated\",\"inputs\":[{\"name\":\"minLock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"maxLock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"maxBoostBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"cooldown\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardAdded\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardPaid\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ubAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lockEnd\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"vubMinted\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unstaked\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ubAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ready\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ubAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20ExceededSafeSupply\",\"inputs\":[{\"name\":\"increasedSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC5805FutureLookup\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clock\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"ERC6372InconsistentClock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"VotesExpiredSignature\",\"inputs\":[{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a08060405234620000d157306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c16620000c257506001600160401b036002600160401b0319828216016200007c575b6040516148ff9081620000d6823960805181818161115a01526127490152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80806200005c565b63f92ee8a960e01b8152600490fd5b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80628cc262146103f357806301ffc9a7146103ee57806306fdde03146103e95780630700037d146103e457806308617b32146103df578063095ea7b3146103da57806315456eba146103d557806318160ddd146103d057806323b872dd146103cb578063248a9ca3146103c65780632def6620146103c15780632f2ff15d146103bc578063313ce567146103b757806336568abe146103b25780633a46b1a8146103ad5780633ccfd60b146103a85780633d18b912146103a35780634bf5d7e91461039e5780634f1ef2861461039957806352d1902d14610394578063587cde1e1461038f5780635c19a95c1461038a5780635de9a137146103855780635eebea20146103805780636c0b3e461461037b5780636fcfff451461037657806370a08231146103715780637502e24e1461036c578063787a08a6146103675780637b0a47ee146103625780637ecebe001461035d57806380faa57d1461035857806384b0196e14610353578063859e6d6a1461034e5780638aee8127146103495780638b876347146103445780638e539e8c1461033f57806391d148541461033a57806391ddadf414610335578063952e68cf1461033057806395d89b411461032b5780639ab24eb014610326578063a09f114314610321578063a217fddf1461031c578063a9059cbb14610317578063ad3cb1cc14610312578063c04637111461030d578063c0c53b8b14610308578063c3cda52014610303578063ccc57490146102fe578063cd3daf9d146102f9578063d547741f146102f4578063dd62ed3e146102ef578063de350feb146102ea578063df136d65146102e5578063e3d89007146102e0578063ebe2b12b146102db578063f0090cf6146102d6578063f037c630146102d1578063f1127ed8146102cc578063f7c618c1146102c75763ffa1ad74146102c2575f80fd5b6122db565b6122b3565b612220565b6121f7565b6121bd565b612197565b61217a565b61215d565b611ff9565b611fc1565b611f78565b611f5e565b611f24565b611e56565b611d19565b611cf0565b611cab565b611c11565b611bf7565b611bd0565b611b99565b611ae7565b61197b565b611950565b6118f6565b611821565b6117e9565b61175e565b6115e9565b611519565b611468565b611411565b6113f4565b6113ce565b6113b0565b611359565b611304565b6112de565b61127b565b611218565b6111f6565b6111b1565b611148565b6110c9565b610fb9565b610f40565b610e5f565b610d8d565b610d46565b610d2b565b610ce0565b610b91565b610b58565b610aba565b610a91565b61094c565b61088e565b610723565b6106a9565b610511565b610469565b61043e565b600435906001600160a01b038216820361040e57565b5f80fd5b602435906001600160a01b038216820361040e57565b604435906001600160a01b038216820361040e57565b3461040e57602036600319011261040e57602061046161045c6103f8565b61239b565b604051908152f35b3461040e57602036600319011261040e5760043563ffffffff60e01b811680910361040e57602090637965db0b60e01b81149081156104ae575b506040519015158152f35b6301ffc9a760e01b1490505f6104a3565b91908251928382525f5b8481106104e9575050825f602080949584010152601f8019910116010190565b6020818301810151848301820152016104c9565b90602061050e9281815201906104bf565b90565b3461040e575f36600319011261040e576040515f5f8051602061478a83398151915280549061053f826123fb565b808552916020916001918281169081156105d4575060011461057c575b6105788661056c81880382611099565b604051918291826104fd565b0390f35b5f90815293507f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab05b8385106105c15750505050810160200161056c826105785f61055c565b80548686018401529382019381016105a4565b90508695506105789693506020925061056c94915060ff191682840152151560051b82010192935f61055c565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace006020526040902090565b6001600160a01b03165f9081527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d016020526040902090565b3461040e57602036600319011261040e576001600160a01b036106ca6103f8565b165f52600b602052602060405f2054604051908152f35b600435906001600160401b038216820361040e57565b602435906001600160401b038216820361040e57565b606435906001600160401b038216820361040e57565b3461040e57608036600319011261040e5761073c6106e1565b6107446106f7565b6044359161075061070d565b92610759612dbf565b6001600160401b038083168015159182610881575b50501561084b577f591624d336d15110719b185ac9f817fab3324297374045a5d3fb1a0e88ab988993816107a96127106108469410156125b2565b6001805467ffffffffffffffff60a01b191660a086901b67ffffffffffffffff60a01b161790556107f0856001600160401b03166001600160401b03196002541617600255565b6107f981600355565b610819826001600160401b03166001600160401b03196004541617600455565b604080516001600160401b0395861681529585166020870152850152909116606083015281906080820190565b0390a1005b60405162461bcd60e51b815260206004820152600e60248201526d626164206c6f636b2072616e676560901b6044820152606490fd5b8516101590505f8061076e565b3461040e57604036600319011261040e576108a76103f8565b6024353315610934576001600160a01b03821691821561091c576108e482916108cf33610601565b9060018060a01b03165f5260205260405f2090565b556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b604051634a1406b160e11b81525f6004820152602490fd5b60405163e602df0560e01b81525f6004820152602490fd5b3461040e57602036600319011261040e57335f908152600560205260409020600435907f0bcadf4a8215096a7cbb695c9385c84784ef6d32892221d088a4f13aa8d491969061099c8315156125eb565b8054151580610a66575b6109af90612625565b6109b833612efc565b610a3f600182016109ff6109f76109f16109ec6109dc85546001600160401b031690565b6001600160401b03421690612662565b612851565b8761234f565b612710900490565b92610a0b86825461238e565b9055610a178333612f4b565b5f54610a329086906001600160a01b03165b30903390613005565b546001600160401b031690565b604080519485526001600160401b039091166020850152830152339180606081015b0390a2005b506109af610a7e60018301546001600160401b031690565b6001600160401b034291161190506109a6565b3461040e575f36600319011261040e5760205f805160206147ea83398151915254604051908152f35b3461040e57606036600319011261040e57610ad36103f8565b610adb610412565b90604435610aec336108cf84610601565b545f198110610afd575b5050613096565b818110610b305750506001600160a01b0381161561093457331561091c57610b28336108cf83610601565b505f80610af6565b604051637dc7a0d960e11b815233600482015260248101919091526044810191909152606490fd5b3461040e57602036600319011261040e576004355f525f8051602061486a8339815191526020526020600160405f200154604051908152f35b3461040e575f36600319011261040e57335f908152600560205260409020610bbb8154151561267b565b610be9610be1610bd560018401546001600160401b031690565b6001600160401b031690565b4210156126b1565b610bf233612efc565b54335f908152600560205260409020610c11905b60015f918281550155565b610c1a33610639565b5480610cd0575b50335f9081526006602052604090207f536c53e11db8105c787d8d5fce8b01f689aefd57771dad0d0c62c33af2ecc1f990610a6190610caf90610c6585825461238e565b8155610a326001610c90610c816004546001600160401b031690565b6001600160401b0342166126e6565b92019182906001600160401b03166001600160401b0319825416179055565b604080519485526001600160401b0390911660208501523393918291820190565b610cda90336130ea565b5f610c21565b3461040e57604036600319011261040e57610d29600435610cff610412565b90805f525f8051602061486a833981519152602052610d24600160405f200154612e91565b613409565b005b3461040e575f36600319011261040e57602060405160128152f35b3461040e57604036600319011261040e57610d5f610412565b336001600160a01b03821603610d7b57610d299060043561343f565b60405163334bd91960e11b8152600490fd5b3461040e57604036600319011261040e57610dae610da96103f8565b610671565b610db96024356134d1565b8154905f829160058411610e0c575b610dd3935084613af8565b9081610df157505060205f5b6040516001600160d01b039091168152f35b610dfc602092612334565b905f52815f20015460301c610ddf565b9192610e178161395b565b8103908111610e5a57610dd393855f5265ffffffffffff808360205f20015416908516105f14610e48575091610dc8565b929150610e5490612380565b90610dc8565b612320565b3461040e575f36600319011261040e57335f52600660205260405f208054908115610f0957610e9e610bd56001610ea69301546001600160401b031690565b421015612701565b335f908152600660205260409020610ebd90610c06565b5f54610ed590829033906001600160a01b0316613512565b60405190815233907f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5908060208101610a61565b60405162461bcd60e51b815260206004820152600f60248201526e6e6f7468696e672070656e64696e6760881b6044820152606490fd5b3461040e575f36600319011261040e57610f5933612efc565b335f52600b60205260405f2080549081610f6f57005b5f9055600154610f8b90829033906001600160a01b0316613512565b6040519081527fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e048660203392a2005b3461040e575f36600319011261040e57610fd243613929565b65ffffffffffff80610fe343613929565b1691160361103857610578604051610ffa8161105e565b601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c7400000060208201526040519182916020835260208301906104bf565b6040516301bfc1c560e61b8152600490fd5b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b0382111761107957604052565b61104a565b60a081019081106001600160401b0382111761107957604052565b90601f801991011681019081106001600160401b0382111761107957604052565b604051906110c78261105e565b565b604036600319011261040e576110dd6103f8565b602435906001600160401b039081831161040e573660238401121561040e5782600401359182116110795760405191611120601f8201601f191660200184611099565b808352366024828601011161040e576020815f926024610d299701838701378401015261273c565b3461040e575f36600319011261040e577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361119f5760206040515f8051602061482a8339815191528152f35b60405163703e46dd60e11b8152600490fd5b3461040e57602036600319011261040e5760206001600160a01b03806111d56103f8565b165f525f8051602061476a833981519152825260405f205416604051908152f35b3461040e57602036600319011261040e57610d296112126103f8565b3361358b565b3461040e57602036600319011261040e576001600160a01b036112396103f8565b165f52600560205260405f206001600160401b03600182549201541690610578604051928392839092916001600160401b036020916040840195845216910152565b3461040e57602036600319011261040e576001600160a01b0361129c6103f8565b165f52600660205260405f206001600160401b03600182549201541690610578604051928392839092916001600160401b036020916040840195845216910152565b3461040e575f36600319011261040e5760206001600160401b0360025416604051908152f35b3461040e57602036600319011261040e57611320610da96103f8565b5463ffffffff9081811161133a5760209160405191168152f35b604490604051906306dfcc6560e41b8252602060048301526024820152fd5b3461040e57602036600319011261040e5760206104616113776103f8565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b3461040e57602036600319011261040e5760206104616109ec6106e1565b3461040e575f36600319011261040e5760206001600160401b0360045416604051908152f35b3461040e575f36600319011261040e576020600754604051908152f35b3461040e57602036600319011261040e576001600160a01b036114326103f8565b165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052602060405f2054604051908152f35b3461040e575f36600319011261040e5760206114826128d8565b6001600160401b0360405191168152f35b916114c790949194600f60f81b84526114b960209660e0602087015260e08601906104bf565b9084820360408601526104bf565b92606083015260018060a01b031660808201525f60a082015260c0818303910152602080845192838152019301915f5b828110611505575050505090565b8351855293810193928101926001016114f7565b3461040e575f36600319011261040e577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1005415806115c0575b156115835761155f612433565b611567612506565b906105786115736128fa565b6040519384933091469186611493565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015415611552565b3461040e57602036600319011261040e576116026106e1565b335f9081526005602052604090207fc0d15b5a903ff969998bbe8c93fa9476a8374e008475560efd46f25b0b4301a290610a61909261164384541515612625565b61171661165b610bd56002546001600160401b031690565b6116716001600160401b03918285161115612920565b6116fa6109f7824216946116f461168882886126e6565b998a9560018201986116b06116a7610bd58c546001600160401b031690565b838a1611612957565b6116b933612efc565b89546001600160401b03169142908316111561174f576116df6116e8916116ee93612662565b925b5494612851565b91612851565b90612342565b9061234f565b92906001600160401b03166001600160401b0319825416179055565b80611740575b604051918291339583602090939291936001600160401b0360408201951681520152565b61174a8133612f4b565b61171c565b50506116ee6116e85f926116e1565b3461040e57602036600319011261040e576117776103f8565b61177f612dbf565b6001600160401b03600854164211156117b457600180546001600160a01b0319166001600160a01b0392909216919091179055005b60405162461bcd60e51b815260206004820152600d60248201526c7265776172642061637469766560981b6044820152606490fd5b3461040e57602036600319011261040e576001600160a01b0361180a6103f8565b165f52600a602052602060405f2054604051908152f35b3461040e57602036600319011261040e5761183d6004356134d1565b5f8051602061484a833981519152908154905f82916005841161189e575b6118659350613a86565b90816118785750506040515f8152602090f35b611883602092612334565b905f525f805160206148aa833981519152015460301c610ddf565b91926118a98161395b565b8103908111610e5a5761186593855f5265ffffffffffff80835f805160206148aa833981519152015416908516105f146118e457509161185b565b9291506118f090612380565b9061185b565b3461040e57604036600319011261040e57602060ff611944611916610412565b6004355f525f8051602061486a833981519152845260405f209060018060a01b03165f5260205260405f2090565b54166040519015158152f35b3461040e575f36600319011261040e57602061196b43613929565b65ffffffffffff60405191168152f35b3461040e57604036600319011261040e576004357f0bcadf4a8215096a7cbb695c9385c84784ef6d32892221d088a4f13aa8d491966119b86106f7565b6119c38315156125eb565b611a3e6109f7611a38611a32846119e9610bd56001546001600160401b039060a01c1690565b611a096001600160401b0391828416908110159081611ac8575b50612990565b335f908152600560205260409020611a229054156129c7565b611a2b33612efc565b42166126e6565b93612851565b8561234f565b90611a6f611a4a6110ba565b8581526001600160401b0383166020820152335f908152600560205260409020612a01565b611a798233612f4b565b335f9081525f8051602061476a83398151915260205260409020546001600160a01b031615611ab9575b5f54610a3f9085906001600160a01b0316610a29565b611ac3333361358b565b611aa3565b9050611adf610bd56002546001600160401b031690565b10155f611a03565b3461040e575f36600319011261040e576040515f5f805160206147ca833981519152805490611b15826123fb565b808552916020916001918281169081156105d45750600114611b41576105788661056c81880382611099565b5f90815293507f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa5b838510611b865750505050810160200161056c826105785f61055c565b8054868601840152938201938101611b69565b3461040e57602036600319011261040e5760206001600160d01b03611bc7611bc2610da96103f8565b613688565b16604051908152f35b3461040e575f36600319011261040e575f546040516001600160a01b039091168152602090f35b3461040e575f36600319011261040e5760206040515f8152f35b3461040e57604036600319011261040e57611c2a6103f8565b3315611c93576001600160a01b031615611c7b5760405162461bcd60e51b81526020600482015260156024820152747655423a206e6f6e2d7472616e7366657261626c6560581b6044820152606490fd5b60405163ec442f0560e01b81525f6004820152602490fd5b604051634b637e8f60e11b81525f6004820152602490fd5b3461040e575f36600319011261040e57610578604051611cca8161105e565b60058152640352e302e360dc1b60208201526040519182916020835260208301906104bf565b3461040e575f36600319011261040e5760206008546001600160401b036040519160401c168152f35b3461040e57606036600319011261040e57611d326103f8565b611d3a610412565b90611d43610428565b5f8051602061488a83398151915254926001600160401b0360ff8560401c1615941680159081611e4e575b6001149081611e44575b159081611e3b575b50611e29575f8051602061488a833981519152805467ffffffffffffffff19166001179055611db39284611e0557612a71565b611db957005b5f8051602061488a833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2908060208101610846565b5f8051602061488a833981519152805460ff60401b1916600160401b179055612a71565b60405163f92ee8a960e01b8152600490fd5b9050155f611d80565b303b159150611d78565b859150611d6e565b3461040e5760c036600319011261040e57611e6f6103f8565b6044359060243560643560ff8116810361040e57834211611f0b57611eff610d2994611f06926040519060208201927fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf845260018060a01b0388166040840152866060840152608083015260808252611ee78261107e565b611efa60a4359360843593519020613819565b6138a8565b91826138c0565b61358b565b604051632341d78760e11b815260048101859052602490fd5b3461040e575f36600319011261040e5760206040517f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f558152f35b3461040e575f36600319011261040e576020610461612cb7565b3461040e57604036600319011261040e57610d29600435611f97610412565b90805f525f8051602061486a833981519152602052611fbc600160405f200154612e91565b61343f565b3461040e57604036600319011261040e576020611ff0611fdf6103f8565b6108cf611fea610412565b91610601565b54604051908152f35b3461040e57604036600319011261040e577fbbb707ba52ee8c7d03c6ff4ddd68fa3d2050fb9fff8f2616f6d5ed4eb5c2e32e6004356120366106f7565b9061203f612e3a565b6120f36120d7836001600160401b0380821661205c811515612d2c565b612064612ebd565b814216918261207b6008546001600160401b031690565b808316821061212b575050506120946120999187612362565b600755565b6120a66007541515612d68565b600880546fffffffffffffffff00000000000000001916604083901b67ffffffffffffffff60401b161790556126e6565b6001600160401b03166001600160401b03196008541617600855565b60015461210a9082906001600160a01b0316610a29565b604080519182526001600160401b0390921660208201529081908101610846565b61214d61209493612143612158969461215394612662565b600754911661234f565b8961238e565b612362565b612099565b3461040e575f36600319011261040e576020600954604051908152f35b3461040e575f36600319011261040e576020600354604051908152f35b3461040e575f36600319011261040e5760206001600160401b0360085416604051908152f35b3461040e575f36600319011261040e5760206040517f61b754da92d8d8d7300489a35a466b9ed19cf4a61860a290f89bec3a75de2bcf8152f35b3461040e575f36600319011261040e5760206001600160401b0360015460a01c16604051908152f35b3461040e57604036600319011261040e576122396103f8565b6024359063ffffffff8216820361040e5760409161226b6122799261225c612da7565b50612265612da7565b50610671565b612273612da7565b50614246565b508151906122868261105e565b54602065ffffffffffff821692838152019060301c8152825191825260018060d01b039051166020820152f35b3461040e575f36600319011261040e576001546040516001600160a01b039091168152602090f35b3461040e575f36600319011261040e576105786040516122fa8161105e565b60058152640312e302e360dc1b60208201526040519182916020835260208301906104bf565b634e487b7160e01b5f52601160045260245ffd5b5f19810191908211610e5a57565b91908203918211610e5a57565b81810292918115918404141715610e5a57565b811561236c570490565b634e487b7160e01b5f52601260045260245ffd5b9060018201809211610e5a57565b91908201809211610e5a57565b6123a481610639565b54906123ae612cb7565b9060018060a01b031691825f52600a60205260405f20548203918211610e5a57670de0b6b3a7640000916123e19161234f565b04905f52600b60205260405f20548101809111610e5a5790565b90600182811c92168015612429575b602083101461241557565b634e487b7160e01b5f52602260045260245ffd5b91607f169161240a565b604051905f825f805160206147aa83398151915291825492612454846123fb565b808452936020916001918281169081156124e05750600114612480575b5050506110c792500383611099565b5f9081527f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d9590935091905b8284106124c857506110c79450505081016020015f8080612471565b855488850183015294850194879450928101926124ac565b92505050602092506110c794915060ff191682840152151560051b8201015f8080612471565b604051905f825f8051602061480a83398151915291825492612527846123fb565b808452936020916001918281169081156124e05750600114612552575050506110c792500383611099565b5f9081527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b759590935091905b82841061259a57506110c79450505081016020015f8080612471565b8554888501830152948501948794509281019261257e565b156125b957565b60405162461bcd60e51b815260206004820152600a6024820152690c4dedee6e840784062f60b31b6044820152606490fd5b156125f257565b60405162461bcd60e51b815260206004820152600b60248201526a1e995c9bc8185b5bdd5b9d60aa1b6044820152606490fd5b1561262c57565b60405162461bcd60e51b815260206004820152600e60248201526d6e6f20616374697665206c6f636b60901b6044820152606490fd5b6001600160401b039182169082160391908211610e5a57565b1561268257565b60405162461bcd60e51b81526020600482015260076024820152666e6f206c6f636b60c81b6044820152606490fd5b156126b857565b60405162461bcd60e51b81526020600482015260066024820152651b1bd8dad95960d21b6044820152606490fd5b9190916001600160401b0380809416911601918211610e5a57565b1561270857565b60405162461bcd60e51b815260206004820152600c60248201526b31b7b7b634b733903237bbb760a11b6044820152606490fd5b6001600160a01b039290917f00000000000000000000000000000000000000000000000000000000000000008416308114908115612836575b5061119f576020600494612787612dbf565b6040516352d1902d60e01b8152958691829087165afa5f9481612805575b506127cb57604051634c9c8ce360e01b81526001600160a01b0384166004820152602490fd5b90915f8051602061482a83398151915284036127ec576110c7929350613b57565b604051632a87526960e21b815260048101859052602490fd5b61282891955060203d60201161282f575b6128208183611099565b810190613571565b935f6127a5565b503d612816565b9050845f8051602061482a833981519152541614155f612775565b6001546001600160401b038281169160a01c8116808311156128ce578160025416809310156128c45760035461270f19810191908211610e5a576128a76128ad92846128a0846128b499612662565b169061234f565b93612662565b1690612362565b612710908101809111610e5a5790565b5050505060035490565b5050505061271090565b6001600160401b0380421690600854168082105f146128f5575090565b905090565b604051602081018181106001600160401b03821117611079576040525f8152905f368137565b1561292757565b60405162461bcd60e51b8152602060048201526008602482015267746f6f206c6f6e6760c01b6044820152606490fd5b1561295e57565b60405162461bcd60e51b815260206004820152600a6024820152693737ba103637b733b2b960b11b6044820152606490fd5b1561299757565b60405162461bcd60e51b8152602060048201526008602482015267626164206c6f636b60c01b6044820152606490fd5b156129ce57565b60405162461bcd60e51b815260206004820152600b60248201526a616374697665206c6f636b60a81b6044820152606490fd5b60016001600160401b0360206110c794805185550151169101906001600160401b03166001600160401b0319825416179055565b60405190612a428261105e565b60038252623b2aa160e91b6020830152565b60405190612a618261105e565b60018252603160f81b6020830152565b90604051612a7e8161105e565b601081526020906f2b37ba3296b2b9b1b937bbb2b2102aa160811b6020820152612aa6612a35565b91612aaf613d2d565b612ab7613d2d565b8151906001600160401b038211611079575f8051602061478a83398151915292612aea83612ae586546123fb565b613d5b565b602091601f8411600114612c19575092612b2983612b3094612c0b9a999794612b9799975f92612c0e575b50508160011b915f199060031b1c19161790565b9055613f3f565b612b38613d2d565b612b51612b43612a35565b612b4b612a54565b906136b0565b612b59613d2d565b5f80546001600160a01b0319166001600160a01b039384161790551660018060a01b03166bffffffffffffffffffffffff60a01b6001541617600155565b6001805467ffffffffffffffff60a01b191661127560a71b179055612bcd6303c267006001600160401b03196002541617600255565b612bd86161a8600355565b612bf262093a806001600160401b03196004541617600455565b612bfb81613213565b50612c05816132c4565b5061339a565b50565b015190505f80612b15565b5f8051602061478a8339815191525f529190601f1984167f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0935f905b828210612c9f57505093612c0b99989693612b9798969360019383612b309810612c87575b505050811b019055613f3f565b01515f1960f88460031b161c191690555f8080612c7a565b80600186978294978701518155019601940190612c55565b5f805160206147ea833981519152548015612d255760095490612cf5612cdb6128d8565b6121436001600160401b03918260085460401c1690612662565b90670de0b6b3a764000091828102928184041490151715610e5a57612d1991612362565b8101809111610e5a5790565b5060095490565b15612d3357565b60405162461bcd60e51b815260206004820152600d60248201526c3d32b93790323ab930ba34b7b760991b6044820152606490fd5b15612d6f57565b60405162461bcd60e51b815260206004820152601060248201526f1c995dd85c99081d1bdbc81cdb585b1b60821b6044820152606490fd5b60405190612db48261105e565b5f6020838281520152565b335f9081527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee803602052604090207f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f559060ff905b541615612e1c5750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b335f9081527fbe647c0ddb1a23068f378cf99f2f2f9a0cb078f402ce464e1c86ed30b305a70a602052604090207f61b754da92d8d8d7300489a35a466b9ed19cf4a61860a290f89bec3a75de2bcf9060ff90612e12565b5f8181525f8051602061486a83398151915260209081526040808320338452909152902060ff90612e12565b612ec5612cb7565b6009556110c7612ed36128d8565b67ffffffffffffffff60401b6008549160401b169067ffffffffffffffff60401b191617600855565b612f04612cb7565b600955612f12612ed36128d8565b6001600160a01b0381169081612f26575050565b612f2f9061239b565b905f52600b60205260405f2055600954600a60205260405f2055565b91906001600160a01b0383168015611c7b575f805160206147ea833981519152908154838101809111610e5a575f805160206147ea83398151915255612f9085610639565b8054840190556040518381525f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602090a354926001600160d01b038411612fdf576110c7929350614351565b604051630e58ae9360e11b8152600481018590526001600160d01b036024820152604490fd5b6040516323b872dd60e01b5f9081526001600160a01b03938416600452938316602452604494909452909160209060648180855af160015f5114811615613077575b836040525f6060521561305957505050565b635274afe760e01b8352166001600160a01b03166004820152602490fd5b600181151661308d57813b15153d151616613047565b833d5f823e3d90fd5b6001600160a01b039190821615611c93571615611c7b5760405162461bcd60e51b81526020600482015260156024820152747655423a206e6f6e2d7472616e7366657261626c6560581b6044820152606490fd5b6001600160a01b0380821692918315611c935761310681610639565b548381106131e4578361311a910391610639565b555f805160206147ea8339815191528281540390555f837fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020604051868152a3613164826142e8565b9261316e43613929565b6001600160d01b03948580613181613637565b1691169003948511610e5a576110c79461319a916145b1565b50505f9081525f8051602061476a83398151915260205260408120549080527fd4fb29e10204005f1a39963c6862b79a755e22f0177c53f05cdc3786c702f9745482169116613bf9565b60405163391434e360e21b81526001600160a01b03929092166004830152602482015260448101839052606490fd5b6001600160a01b0381165f9081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f8051602061486a8339815191529060ff166132be575f808052602091825260408082206001600160a01b038516835290925220805460ff1916600117905533906001600160a01b03165f7f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50505f90565b6001600160a01b0381165f9081527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee803602052604090207f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55905f8051602061486a8339815191529060ff905b5416613393575f828152602091825260408082206001600160a01b038616835290925220805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b5050505f90565b6001600160a01b0381165f9081527fbe647c0ddb1a23068f378cf99f2f2f9a0cb078f402ce464e1c86ed30b305a70a602052604090207f61b754da92d8d8d7300489a35a466b9ed19cf4a61860a290f89bec3a75de2bcf905f8051602061486a8339815191529060ff9061332f565b5f8181525f8051602061486a833981519152602081815260408084206001600160a01b038716855290915290912060ff9061332f565b5f8181525f8051602061486a833981519152602081815260408084206001600160a01b03871685529091529091205460ff1615613393575f828152602091825260408082206001600160a01b038616835290925220805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b65ffffffffffff6134e143613929565b16808210156134f4575061050e90613929565b6044925060405191637669fc0f60e11b835260048301526024820152fd5b60405163a9059cbb60e01b5f9081526001600160a01b039384166004526024949094529260209060448180855af160015f511481161561355b575b836040521561305957505050565b600181151661308d57813b15153d15161661354d565b9081602091031261040e575190565b6040513d5f823e3d90fd5b6001600160a01b038181165f8181525f8051602061476a8339815191526020526040812080548685166001600160a01b0319821681179092556110c7969416946136319390928691907f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9080a46001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b91613bf9565b5f8051602061484a8339815191528054806136525750505f90565b805f19810111610e5a577f88c46c62109817164d0ae1873830d4299a82e5daf552a3d8e989b27638fcf747915f52015460301c90565b8054806136955750505f90565b5f19918183810111610e5a575f5260205f2001015460301c90565b91906136ba613d2d565b6136c2613d2d565b82516001600160401b038111611079575f805160206147aa833981519152906136f4816136ef84546123fb565b613dd4565b602080601f831160011461377f575090806137289261372f96975f92612c0e5750508160011b915f199060031b1c19161790565b9055614038565b6137575f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10055565b6110c75f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10155565b90601f198316966137bd5f805160206147aa8339815191525f527f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d90565b925f905b8982106138015750509083929160019461372f9899106137e9575b505050811b019055614038565b01515f1960f88460031b161c191690555f80806137dc565b806001859682949686015181550195019301906137c1565b613821614442565b6138296144ac565b916040519260208401927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604085015260608401524660808401523060a084015260a0835260c08301918383106001600160401b038411176110795760429360e291846040528151902061190160f01b855260c282015201522090565b9161050e93916138b793614114565b909291926141b9565b6001600160a01b03165f8181527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0060205260409020805460018101909155909181900361390b575050565b60449250604051916301d4b62360e61b835260048301526024820152fd5b65ffffffffffff9081811161393c571690565b604490604051906306dfcc6560e41b8252603060048301526024820152fd5b600181111561050e57600181600160801b811015613a74575b613a1c613a12613a086139fe6139f46139ea613a2897600488600160401b613a239a1015613a67575b640100000000811015613a5a575b62010000811015613a4d575b610100811015613a41575b6010811015613a35575b1015613a2d575b60030260011c6139e3818b612362565b0160011c90565b6139e3818a612362565b6139e38189612362565b6139e38188612362565b6139e38187612362565b6139e38186612362565b8093612362565b821190565b900390565b60011b6139d3565b811c9160021b916139cc565b60081c91811b916139c2565b60101c9160081b916139b7565b60201c9160101b916139ab565b60401c9160201b9161399d565b50600160401b9050608082901c613974565b905b828110613a9457505090565b90918082169080831860011c8201809211610e5a575f8051602061484a8339815191525f5265ffffffffffff80835f805160206148aa833981519152015416908516105f14613ae65750915b90613a88565b929150613af290612380565b90613ae0565b91905b838210613b085750505090565b9091928083169080841860011c8201809211610e5a57845f5265ffffffffffff808360205f20015416908416105f14613b455750925b9190613afb565b939250613b5190612380565b91613b3e565b90813b15613bd8575f8051602061482a83398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2805115613bbd57612c0b9161426f565b505034613bc657565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b6001600160a01b03808316939291908116908185141580613d24575b613c21575b5050505050565b81613c96575b505082613c36575b8080613c1a565b7fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a72491613c6d613c67613c7393610671565b916142e8565b9061431b565b604080516001600160d01b039384168152919092166020820152a25f8080613c2f565b613c9f90610671565b613ca8846142e8565b613cb143613929565b6001600160d01b03918280613cc586613688565b169116900392828411610e5a577fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a72493613d1a92613d01926146af565b6040805192851683529316602082015291829190820190565b0390a25f80613c27565b50831515613c15565b60ff5f8051602061488a8339815191525460401c1615613d4957565b604051631afcd79f60e31b8152600490fd5b601f8111613d67575050565b5f8051602061478a8339815191525f527f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0906020601f840160051c83019310613dca575b601f0160051c01905b818110613dbf575050565b5f8155600101613db4565b9091508190613dab565b601f8111613de0575050565b5f805160206147aa8339815191525f527f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d906020601f840160051c83019310613e43575b601f0160051c01905b818110613e38575050565b5f8155600101613e2d565b9091508190613e24565b601f8111613e59575050565b5f805160206147ca8339815191525f527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa906020601f840160051c83019310613ebc575b601f0160051c01905b818110613eb1575050565b5f8155600101613ea6565b9091508190613e9d565b601f8111613ed2575050565b5f8051602061480a8339815191525f527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75906020601f840160051c83019310613f35575b601f0160051c01905b818110613f2a575050565b5f8155600101613f1f565b9091508190613f16565b9081516001600160401b038111611079575f805160206147ca83398151915290613f7281613f6d84546123fb565b613e4d565b602080601f8311600114613fa757508190613fa39394955f92612c0e5750508160011b915f199060031b1c19161790565b9055565b90601f19831695613fe55f805160206147ca8339815191525f527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa90565b925f905b88821061402057505083600195969710614008575b505050811b019055565b01515f1960f88460031b161c191690555f8080613ffe565b80600185968294968601518155019501930190613fe9565b9081516001600160401b038111611079575f8051602061480a8339815191529061406b8161406684546123fb565b613ec6565b602080601f831160011461409c57508190613fa39394955f92612c0e5750508160011b915f199060031b1c19161790565b90601f198316956140da5f8051602061480a8339815191525f527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7590565b925f905b8882106140fc5750508360019596971061400857505050811b019055565b806001859682949686015181550195019301906140de565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411614190579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa1561418b575f516001600160a01b0381161561418157905f905f90565b505f906001905f90565b613580565b5050505f9160039190565b600411156141a557565b634e487b7160e01b5f52602160045260245ffd5b6141c28161419b565b806141cb575050565b6141d48161419b565b600181036141ee5760405163f645eedf60e01b8152600490fd5b6141f78161419b565b600281036142185760405163fce698f760e01b815260048101839052602490fd5b8061422460039261419b565b1461422c5750565b6040516335e2f38360e21b81526004810191909152602490fd5b805482101561425b575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b905f8091602081519101845af480806142d5575b1561429257505061050e614429565b156142ba57604051639996b31560e01b81526001600160a01b03919091166004820152602490fd5b3d15155f036135805760405163d6bda27560e01b8152600490fd5b503d1515806142835750813b1515614283565b6001600160d01b03908181116142fc571690565b604490604051906306dfcc6560e41b825260d060048301526024820152fd5b9061432543613929565b6001600160d01b0391828061433986613688565b16911601918211610e5a5761434d926146af565b9091565b9061435b816142e8565b9161436543613929565b6001600160d01b03938480614378613637565b16911601848111610e5a5761438c916145b1565b50506001600160a01b039081169081156143ec575b5f8051602061476a8339815191526020527fd4fb29e10204005f1a39963c6862b79a755e22f0177c53f05cdc3786c702f974545f9283526040909220546110c7945081169116613bf9565b6143f5836142e8565b6143fe43613929565b908580614409613637565b1691169003948511610e5a576110c794614422916145b1565b50506143a1565b604051903d82523d5f602084013e60203d830101604052565b61444a612433565b805190811561445a576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1005480156144875790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b6144b4612506565b80519081156144c4576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015480156144875790565b5f8051602061484a833981519152908154600160401b811015611079576001810180845581101561425b575f92909252805160209091015160301b65ffffffffffff191665ffffffffffff91909116175f805160206148aa8339815191529190910155565b8054600160401b8110156110795761457391600182018155614246565b61459e57815160209092015160301b65ffffffffffff191665ffffffffffff92909216919091179055565b634e487b7160e01b5f525f60045260245ffd5b5f8051602061484a833981519152549192918015614686576145d56145f791612334565b5f8051602061484a8339815191525f525f805160206148aa8339815191520190565b9081549165ffffffffffff908184169183168083116146745786920361463c5761463592509065ffffffffffff82549181199060301b169116179055565b60301c9190565b505061466f9061465b61464d6110ba565b65ffffffffffff9092168252565b6001600160d01b03851660208201526144f1565b614635565b604051632520601d60e01b8152600490fd5b506146aa9061469661464d6110ba565b6001600160d01b03841660208201526144f1565b5f9190565b80549293928015614744576146c66146d191612334565b825f5260205f200190565b9182549265ffffffffffff9182851692811680841161467457879303614710575061463592509065ffffffffffff82549181199060301b169116179055565b91505061466f916147306147226110ba565b65ffffffffffff9093168352565b6001600160d01b0386166020830152614556565b50906146aa916147556147226110ba565b6001600160d01b038516602083015261455656fee8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d0052c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10252c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0452c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbce8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d0202dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0088c46c62109817164d0ae1873830d4299a82e5daf552a3d8e989b27638fcf748a2646970667358221220c942b2f385148fb7c55963446470b86ac34a6a367f0cb752efc7c5bd421ca29564736f6c63430008180033",
}

// VUBABI is the input ABI used to generate the binding from.
// Deprecated: Use VUBMetaData.ABI instead.
var VUBABI = VUBMetaData.ABI

// VUBBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VUBMetaData.Bin instead.
var VUBBin = VUBMetaData.Bin

// DeployVUB deploys a new Ethereum contract, binding an instance of VUB to it.
func DeployVUB(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VUB, error) {
	parsed, err := VUBMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VUBBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VUB{VUBCaller: VUBCaller{contract: contract}, VUBTransactor: VUBTransactor{contract: contract}, VUBFilterer: VUBFilterer{contract: contract}}, nil
}

// VUB is an auto generated Go binding around an Ethereum contract.
type VUB struct {
	VUBCaller     // Read-only binding to the contract
	VUBTransactor // Write-only binding to the contract
	VUBFilterer   // Log filterer for contract events
}

// VUBCaller is an auto generated read-only Go binding around an Ethereum contract.
type VUBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VUBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VUBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VUBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VUBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VUBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VUBSession struct {
	Contract     *VUB              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VUBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VUBCallerSession struct {
	Contract *VUBCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VUBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VUBTransactorSession struct {
	Contract     *VUBTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VUBRaw is an auto generated low-level Go binding around an Ethereum contract.
type VUBRaw struct {
	Contract *VUB // Generic contract binding to access the raw methods on
}

// VUBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VUBCallerRaw struct {
	Contract *VUBCaller // Generic read-only contract binding to access the raw methods on
}

// VUBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VUBTransactorRaw struct {
	Contract *VUBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVUB creates a new instance of VUB, bound to a specific deployed contract.
func NewVUB(address common.Address, backend bind.ContractBackend) (*VUB, error) {
	contract, err := bindVUB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VUB{VUBCaller: VUBCaller{contract: contract}, VUBTransactor: VUBTransactor{contract: contract}, VUBFilterer: VUBFilterer{contract: contract}}, nil
}

// NewVUBCaller creates a new read-only instance of VUB, bound to a specific deployed contract.
func NewVUBCaller(address common.Address, caller bind.ContractCaller) (*VUBCaller, error) {
	contract, err := bindVUB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VUBCaller{contract: contract}, nil
}

// NewVUBTransactor creates a new write-only instance of VUB, bound to a specific deployed contract.
func NewVUBTransactor(address common.Address, transactor bind.ContractTransactor) (*VUBTransactor, error) {
	contract, err := bindVUB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VUBTransactor{contract: contract}, nil
}

// NewVUBFilterer creates a new log filterer instance of VUB, bound to a specific deployed contract.
func NewVUBFilterer(address common.Address, filterer bind.ContractFilterer) (*VUBFilterer, error) {
	contract, err := bindVUB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VUBFilterer{contract: contract}, nil
}

// bindVUB binds a generic wrapper to an already deployed contract.
func bindVUB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VUBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VUB *VUBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VUB.Contract.VUBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VUB *VUBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VUB.Contract.VUBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VUB *VUBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VUB.Contract.VUBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VUB *VUBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VUB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VUB *VUBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VUB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VUB *VUBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VUB.Contract.contract.Transact(opts, method, params...)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_VUB *VUBCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_VUB *VUBSession) CLOCKMODE() (string, error) {
	return _VUB.Contract.CLOCKMODE(&_VUB.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_VUB *VUBCallerSession) CLOCKMODE() (string, error) {
	return _VUB.Contract.CLOCKMODE(&_VUB.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VUB *VUBCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VUB *VUBSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VUB.Contract.DEFAULTADMINROLE(&_VUB.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VUB *VUBCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VUB.Contract.DEFAULTADMINROLE(&_VUB.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_VUB *VUBCaller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_VUB *VUBSession) GOVERNORROLE() ([32]byte, error) {
	return _VUB.Contract.GOVERNORROLE(&_VUB.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_VUB *VUBCallerSession) GOVERNORROLE() ([32]byte, error) {
	return _VUB.Contract.GOVERNORROLE(&_VUB.CallOpts)
}

// REWARDFUNDERROLE is a free data retrieval call binding the contract method 0xf0090cf6.
//
// Solidity: function REWARD_FUNDER_ROLE() view returns(bytes32)
func (_VUB *VUBCaller) REWARDFUNDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "REWARD_FUNDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDFUNDERROLE is a free data retrieval call binding the contract method 0xf0090cf6.
//
// Solidity: function REWARD_FUNDER_ROLE() view returns(bytes32)
func (_VUB *VUBSession) REWARDFUNDERROLE() ([32]byte, error) {
	return _VUB.Contract.REWARDFUNDERROLE(&_VUB.CallOpts)
}

// REWARDFUNDERROLE is a free data retrieval call binding the contract method 0xf0090cf6.
//
// Solidity: function REWARD_FUNDER_ROLE() view returns(bytes32)
func (_VUB *VUBCallerSession) REWARDFUNDERROLE() ([32]byte, error) {
	return _VUB.Contract.REWARDFUNDERROLE(&_VUB.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VUB *VUBCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VUB *VUBSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _VUB.Contract.UPGRADEINTERFACEVERSION(&_VUB.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VUB *VUBCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _VUB.Contract.UPGRADEINTERFACEVERSION(&_VUB.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_VUB *VUBCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_VUB *VUBSession) VERSION() (string, error) {
	return _VUB.Contract.VERSION(&_VUB.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_VUB *VUBCallerSession) VERSION() (string, error) {
	return _VUB.Contract.VERSION(&_VUB.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VUB *VUBCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VUB *VUBSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VUB.Contract.Allowance(&_VUB.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VUB *VUBCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VUB.Contract.Allowance(&_VUB.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VUB *VUBCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VUB *VUBSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VUB.Contract.BalanceOf(&_VUB.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VUB *VUBCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VUB.Contract.BalanceOf(&_VUB.CallOpts, account)
}

// BoostBps is a free data retrieval call binding the contract method 0x7502e24e.
//
// Solidity: function boostBps(uint64 dur) view returns(uint256)
func (_VUB *VUBCaller) BoostBps(opts *bind.CallOpts, dur uint64) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "boostBps", dur)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoostBps is a free data retrieval call binding the contract method 0x7502e24e.
//
// Solidity: function boostBps(uint64 dur) view returns(uint256)
func (_VUB *VUBSession) BoostBps(dur uint64) (*big.Int, error) {
	return _VUB.Contract.BoostBps(&_VUB.CallOpts, dur)
}

// BoostBps is a free data retrieval call binding the contract method 0x7502e24e.
//
// Solidity: function boostBps(uint64 dur) view returns(uint256)
func (_VUB *VUBCallerSession) BoostBps(dur uint64) (*big.Int, error) {
	return _VUB.Contract.BoostBps(&_VUB.CallOpts, dur)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (_VUB *VUBCaller) Checkpoints(opts *bind.CallOpts, account common.Address, pos uint32) (CheckpointsCheckpoint208, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "checkpoints", account, pos)

	if err != nil {
		return *new(CheckpointsCheckpoint208), err
	}

	out0 := *abi.ConvertType(out[0], new(CheckpointsCheckpoint208)).(*CheckpointsCheckpoint208)

	return out0, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (_VUB *VUBSession) Checkpoints(account common.Address, pos uint32) (CheckpointsCheckpoint208, error) {
	return _VUB.Contract.Checkpoints(&_VUB.CallOpts, account, pos)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (_VUB *VUBCallerSession) Checkpoints(account common.Address, pos uint32) (CheckpointsCheckpoint208, error) {
	return _VUB.Contract.Checkpoints(&_VUB.CallOpts, account, pos)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_VUB *VUBCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_VUB *VUBSession) Clock() (*big.Int, error) {
	return _VUB.Contract.Clock(&_VUB.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_VUB *VUBCallerSession) Clock() (*big.Int, error) {
	return _VUB.Contract.Clock(&_VUB.CallOpts)
}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint64)
func (_VUB *VUBCaller) Cooldown(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "cooldown")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint64)
func (_VUB *VUBSession) Cooldown() (uint64, error) {
	return _VUB.Contract.Cooldown(&_VUB.CallOpts)
}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint64)
func (_VUB *VUBCallerSession) Cooldown() (uint64, error) {
	return _VUB.Contract.Cooldown(&_VUB.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VUB *VUBCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VUB *VUBSession) Decimals() (uint8, error) {
	return _VUB.Contract.Decimals(&_VUB.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VUB *VUBCallerSession) Decimals() (uint8, error) {
	return _VUB.Contract.Decimals(&_VUB.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_VUB *VUBCaller) Delegates(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "delegates", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_VUB *VUBSession) Delegates(account common.Address) (common.Address, error) {
	return _VUB.Contract.Delegates(&_VUB.CallOpts, account)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_VUB *VUBCallerSession) Delegates(account common.Address) (common.Address, error) {
	return _VUB.Contract.Delegates(&_VUB.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address a) view returns(uint256)
func (_VUB *VUBCaller) Earned(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "earned", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address a) view returns(uint256)
func (_VUB *VUBSession) Earned(a common.Address) (*big.Int, error) {
	return _VUB.Contract.Earned(&_VUB.CallOpts, a)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address a) view returns(uint256)
func (_VUB *VUBCallerSession) Earned(a common.Address) (*big.Int, error) {
	return _VUB.Contract.Earned(&_VUB.CallOpts, a)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VUB *VUBCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "eip712Domain")

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
func (_VUB *VUBSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _VUB.Contract.Eip712Domain(&_VUB.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VUB *VUBCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _VUB.Contract.Eip712Domain(&_VUB.CallOpts)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (_VUB *VUBCaller) GetPastTotalSupply(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "getPastTotalSupply", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (_VUB *VUBSession) GetPastTotalSupply(timepoint *big.Int) (*big.Int, error) {
	return _VUB.Contract.GetPastTotalSupply(&_VUB.CallOpts, timepoint)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (_VUB *VUBCallerSession) GetPastTotalSupply(timepoint *big.Int) (*big.Int, error) {
	return _VUB.Contract.GetPastTotalSupply(&_VUB.CallOpts, timepoint)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (_VUB *VUBCaller) GetPastVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "getPastVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (_VUB *VUBSession) GetPastVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _VUB.Contract.GetPastVotes(&_VUB.CallOpts, account, timepoint)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (_VUB *VUBCallerSession) GetPastVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _VUB.Contract.GetPastVotes(&_VUB.CallOpts, account, timepoint)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VUB *VUBCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VUB *VUBSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VUB.Contract.GetRoleAdmin(&_VUB.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VUB *VUBCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VUB.Contract.GetRoleAdmin(&_VUB.CallOpts, role)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_VUB *VUBCaller) GetVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_VUB *VUBSession) GetVotes(account common.Address) (*big.Int, error) {
	return _VUB.Contract.GetVotes(&_VUB.CallOpts, account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_VUB *VUBCallerSession) GetVotes(account common.Address) (*big.Int, error) {
	return _VUB.Contract.GetVotes(&_VUB.CallOpts, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VUB *VUBCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VUB *VUBSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VUB.Contract.HasRole(&_VUB.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VUB *VUBCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VUB.Contract.HasRole(&_VUB.CallOpts, role, account)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint64)
func (_VUB *VUBCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint64)
func (_VUB *VUBSession) LastTimeRewardApplicable() (uint64, error) {
	return _VUB.Contract.LastTimeRewardApplicable(&_VUB.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint64)
func (_VUB *VUBCallerSession) LastTimeRewardApplicable() (uint64, error) {
	return _VUB.Contract.LastTimeRewardApplicable(&_VUB.CallOpts)
}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint64)
func (_VUB *VUBCaller) LastUpdate(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "lastUpdate")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint64)
func (_VUB *VUBSession) LastUpdate() (uint64, error) {
	return _VUB.Contract.LastUpdate(&_VUB.CallOpts)
}

// LastUpdate is a free data retrieval call binding the contract method 0xc0463711.
//
// Solidity: function lastUpdate() view returns(uint64)
func (_VUB *VUBCallerSession) LastUpdate() (uint64, error) {
	return _VUB.Contract.LastUpdate(&_VUB.CallOpts)
}

// Locks is a free data retrieval call binding the contract method 0x5de9a137.
//
// Solidity: function locks(address ) view returns(uint256 amount, uint64 end)
func (_VUB *VUBCaller) Locks(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	End    uint64
}, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "locks", arg0)

	outstruct := new(struct {
		Amount *big.Int
		End    uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// Locks is a free data retrieval call binding the contract method 0x5de9a137.
//
// Solidity: function locks(address ) view returns(uint256 amount, uint64 end)
func (_VUB *VUBSession) Locks(arg0 common.Address) (struct {
	Amount *big.Int
	End    uint64
}, error) {
	return _VUB.Contract.Locks(&_VUB.CallOpts, arg0)
}

// Locks is a free data retrieval call binding the contract method 0x5de9a137.
//
// Solidity: function locks(address ) view returns(uint256 amount, uint64 end)
func (_VUB *VUBCallerSession) Locks(arg0 common.Address) (struct {
	Amount *big.Int
	End    uint64
}, error) {
	return _VUB.Contract.Locks(&_VUB.CallOpts, arg0)
}

// MaxBoostBps is a free data retrieval call binding the contract method 0xe3d89007.
//
// Solidity: function maxBoostBps() view returns(uint256)
func (_VUB *VUBCaller) MaxBoostBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "maxBoostBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBoostBps is a free data retrieval call binding the contract method 0xe3d89007.
//
// Solidity: function maxBoostBps() view returns(uint256)
func (_VUB *VUBSession) MaxBoostBps() (*big.Int, error) {
	return _VUB.Contract.MaxBoostBps(&_VUB.CallOpts)
}

// MaxBoostBps is a free data retrieval call binding the contract method 0xe3d89007.
//
// Solidity: function maxBoostBps() view returns(uint256)
func (_VUB *VUBCallerSession) MaxBoostBps() (*big.Int, error) {
	return _VUB.Contract.MaxBoostBps(&_VUB.CallOpts)
}

// MaxLock is a free data retrieval call binding the contract method 0x6c0b3e46.
//
// Solidity: function maxLock() view returns(uint64)
func (_VUB *VUBCaller) MaxLock(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "maxLock")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxLock is a free data retrieval call binding the contract method 0x6c0b3e46.
//
// Solidity: function maxLock() view returns(uint64)
func (_VUB *VUBSession) MaxLock() (uint64, error) {
	return _VUB.Contract.MaxLock(&_VUB.CallOpts)
}

// MaxLock is a free data retrieval call binding the contract method 0x6c0b3e46.
//
// Solidity: function maxLock() view returns(uint64)
func (_VUB *VUBCallerSession) MaxLock() (uint64, error) {
	return _VUB.Contract.MaxLock(&_VUB.CallOpts)
}

// MinLock is a free data retrieval call binding the contract method 0xf037c630.
//
// Solidity: function minLock() view returns(uint64)
func (_VUB *VUBCaller) MinLock(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "minLock")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinLock is a free data retrieval call binding the contract method 0xf037c630.
//
// Solidity: function minLock() view returns(uint64)
func (_VUB *VUBSession) MinLock() (uint64, error) {
	return _VUB.Contract.MinLock(&_VUB.CallOpts)
}

// MinLock is a free data retrieval call binding the contract method 0xf037c630.
//
// Solidity: function minLock() view returns(uint64)
func (_VUB *VUBCallerSession) MinLock() (uint64, error) {
	return _VUB.Contract.MinLock(&_VUB.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VUB *VUBCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VUB *VUBSession) Name() (string, error) {
	return _VUB.Contract.Name(&_VUB.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VUB *VUBCallerSession) Name() (string, error) {
	return _VUB.Contract.Name(&_VUB.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_VUB *VUBCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_VUB *VUBSession) Nonces(owner common.Address) (*big.Int, error) {
	return _VUB.Contract.Nonces(&_VUB.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_VUB *VUBCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _VUB.Contract.Nonces(&_VUB.CallOpts, owner)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_VUB *VUBCaller) NumCheckpoints(opts *bind.CallOpts, account common.Address) (uint32, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "numCheckpoints", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_VUB *VUBSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _VUB.Contract.NumCheckpoints(&_VUB.CallOpts, account)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_VUB *VUBCallerSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _VUB.Contract.NumCheckpoints(&_VUB.CallOpts, account)
}

// Pending is a free data retrieval call binding the contract method 0x5eebea20.
//
// Solidity: function pending(address ) view returns(uint256 amount, uint64 ready)
func (_VUB *VUBCaller) Pending(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	Ready  uint64
}, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "pending", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Ready  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ready = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// Pending is a free data retrieval call binding the contract method 0x5eebea20.
//
// Solidity: function pending(address ) view returns(uint256 amount, uint64 ready)
func (_VUB *VUBSession) Pending(arg0 common.Address) (struct {
	Amount *big.Int
	Ready  uint64
}, error) {
	return _VUB.Contract.Pending(&_VUB.CallOpts, arg0)
}

// Pending is a free data retrieval call binding the contract method 0x5eebea20.
//
// Solidity: function pending(address ) view returns(uint256 amount, uint64 ready)
func (_VUB *VUBCallerSession) Pending(arg0 common.Address) (struct {
	Amount *big.Int
	Ready  uint64
}, error) {
	return _VUB.Contract.Pending(&_VUB.CallOpts, arg0)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint64)
func (_VUB *VUBCaller) PeriodFinish(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint64)
func (_VUB *VUBSession) PeriodFinish() (uint64, error) {
	return _VUB.Contract.PeriodFinish(&_VUB.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint64)
func (_VUB *VUBCallerSession) PeriodFinish() (uint64, error) {
	return _VUB.Contract.PeriodFinish(&_VUB.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VUB *VUBCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VUB *VUBSession) ProxiableUUID() ([32]byte, error) {
	return _VUB.Contract.ProxiableUUID(&_VUB.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VUB *VUBCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VUB.Contract.ProxiableUUID(&_VUB.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_VUB *VUBCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_VUB *VUBSession) RewardPerToken() (*big.Int, error) {
	return _VUB.Contract.RewardPerToken(&_VUB.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_VUB *VUBCallerSession) RewardPerToken() (*big.Int, error) {
	return _VUB.Contract.RewardPerToken(&_VUB.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_VUB *VUBCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_VUB *VUBSession) RewardPerTokenStored() (*big.Int, error) {
	return _VUB.Contract.RewardPerTokenStored(&_VUB.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_VUB *VUBCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _VUB.Contract.RewardPerTokenStored(&_VUB.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_VUB *VUBCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_VUB *VUBSession) RewardRate() (*big.Int, error) {
	return _VUB.Contract.RewardRate(&_VUB.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_VUB *VUBCallerSession) RewardRate() (*big.Int, error) {
	return _VUB.Contract.RewardRate(&_VUB.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_VUB *VUBCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_VUB *VUBSession) RewardToken() (common.Address, error) {
	return _VUB.Contract.RewardToken(&_VUB.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_VUB *VUBCallerSession) RewardToken() (common.Address, error) {
	return _VUB.Contract.RewardToken(&_VUB.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_VUB *VUBCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_VUB *VUBSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _VUB.Contract.Rewards(&_VUB.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_VUB *VUBCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _VUB.Contract.Rewards(&_VUB.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VUB *VUBCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VUB *VUBSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VUB.Contract.SupportsInterface(&_VUB.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VUB *VUBCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VUB.Contract.SupportsInterface(&_VUB.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VUB *VUBCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VUB *VUBSession) Symbol() (string, error) {
	return _VUB.Contract.Symbol(&_VUB.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VUB *VUBCallerSession) Symbol() (string, error) {
	return _VUB.Contract.Symbol(&_VUB.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VUB *VUBCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VUB *VUBSession) TotalSupply() (*big.Int, error) {
	return _VUB.Contract.TotalSupply(&_VUB.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VUB *VUBCallerSession) TotalSupply() (*big.Int, error) {
	return _VUB.Contract.TotalSupply(&_VUB.CallOpts)
}

// Ub is a free data retrieval call binding the contract method 0xa09f1143.
//
// Solidity: function ub() view returns(address)
func (_VUB *VUBCaller) Ub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "ub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ub is a free data retrieval call binding the contract method 0xa09f1143.
//
// Solidity: function ub() view returns(address)
func (_VUB *VUBSession) Ub() (common.Address, error) {
	return _VUB.Contract.Ub(&_VUB.CallOpts)
}

// Ub is a free data retrieval call binding the contract method 0xa09f1143.
//
// Solidity: function ub() view returns(address)
func (_VUB *VUBCallerSession) Ub() (common.Address, error) {
	return _VUB.Contract.Ub(&_VUB.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_VUB *VUBCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_VUB *VUBSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _VUB.Contract.UserRewardPerTokenPaid(&_VUB.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_VUB *VUBCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _VUB.Contract.UserRewardPerTokenPaid(&_VUB.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_VUB *VUBTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_VUB *VUBSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VUB.Contract.Approve(&_VUB.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_VUB *VUBTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VUB.Contract.Approve(&_VUB.TransactOpts, spender, value)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_VUB *VUBTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_VUB *VUBSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _VUB.Contract.Delegate(&_VUB.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_VUB *VUBTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _VUB.Contract.Delegate(&_VUB.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_VUB *VUBTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_VUB *VUBSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VUB.Contract.DelegateBySig(&_VUB.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_VUB *VUBTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VUB.Contract.DelegateBySig(&_VUB.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// ExtendLock is a paid mutator transaction binding the contract method 0x859e6d6a.
//
// Solidity: function extendLock(uint64 newDur) returns()
func (_VUB *VUBTransactor) ExtendLock(opts *bind.TransactOpts, newDur uint64) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "extendLock", newDur)
}

// ExtendLock is a paid mutator transaction binding the contract method 0x859e6d6a.
//
// Solidity: function extendLock(uint64 newDur) returns()
func (_VUB *VUBSession) ExtendLock(newDur uint64) (*types.Transaction, error) {
	return _VUB.Contract.ExtendLock(&_VUB.TransactOpts, newDur)
}

// ExtendLock is a paid mutator transaction binding the contract method 0x859e6d6a.
//
// Solidity: function extendLock(uint64 newDur) returns()
func (_VUB *VUBTransactorSession) ExtendLock(newDur uint64) (*types.Transaction, error) {
	return _VUB.Contract.ExtendLock(&_VUB.TransactOpts, newDur)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_VUB *VUBTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_VUB *VUBSession) GetReward() (*types.Transaction, error) {
	return _VUB.Contract.GetReward(&_VUB.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_VUB *VUBTransactorSession) GetReward() (*types.Transaction, error) {
	return _VUB.Contract.GetReward(&_VUB.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VUB *VUBTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VUB *VUBSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VUB.Contract.GrantRole(&_VUB.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VUB *VUBTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VUB.Contract.GrantRole(&_VUB.TransactOpts, role, account)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x15456eba.
//
// Solidity: function increaseAmount(uint256 amount) returns()
func (_VUB *VUBTransactor) IncreaseAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "increaseAmount", amount)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x15456eba.
//
// Solidity: function increaseAmount(uint256 amount) returns()
func (_VUB *VUBSession) IncreaseAmount(amount *big.Int) (*types.Transaction, error) {
	return _VUB.Contract.IncreaseAmount(&_VUB.TransactOpts, amount)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x15456eba.
//
// Solidity: function increaseAmount(uint256 amount) returns()
func (_VUB *VUBTransactorSession) IncreaseAmount(amount *big.Int) (*types.Transaction, error) {
	return _VUB.Contract.IncreaseAmount(&_VUB.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _ub, address _rewardToken, address initialOwner) returns()
func (_VUB *VUBTransactor) Initialize(opts *bind.TransactOpts, _ub common.Address, _rewardToken common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "initialize", _ub, _rewardToken, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _ub, address _rewardToken, address initialOwner) returns()
func (_VUB *VUBSession) Initialize(_ub common.Address, _rewardToken common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _VUB.Contract.Initialize(&_VUB.TransactOpts, _ub, _rewardToken, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _ub, address _rewardToken, address initialOwner) returns()
func (_VUB *VUBTransactorSession) Initialize(_ub common.Address, _rewardToken common.Address, initialOwner common.Address) (*types.Transaction, error) {
	return _VUB.Contract.Initialize(&_VUB.TransactOpts, _ub, _rewardToken, initialOwner)
}

// NotifyReward is a paid mutator transaction binding the contract method 0xde350feb.
//
// Solidity: function notifyReward(uint256 amount, uint64 duration) returns()
func (_VUB *VUBTransactor) NotifyReward(opts *bind.TransactOpts, amount *big.Int, duration uint64) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "notifyReward", amount, duration)
}

// NotifyReward is a paid mutator transaction binding the contract method 0xde350feb.
//
// Solidity: function notifyReward(uint256 amount, uint64 duration) returns()
func (_VUB *VUBSession) NotifyReward(amount *big.Int, duration uint64) (*types.Transaction, error) {
	return _VUB.Contract.NotifyReward(&_VUB.TransactOpts, amount, duration)
}

// NotifyReward is a paid mutator transaction binding the contract method 0xde350feb.
//
// Solidity: function notifyReward(uint256 amount, uint64 duration) returns()
func (_VUB *VUBTransactorSession) NotifyReward(amount *big.Int, duration uint64) (*types.Transaction, error) {
	return _VUB.Contract.NotifyReward(&_VUB.TransactOpts, amount, duration)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VUB *VUBTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VUB *VUBSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VUB.Contract.RenounceRole(&_VUB.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VUB *VUBTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VUB.Contract.RenounceRole(&_VUB.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VUB *VUBTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VUB *VUBSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VUB.Contract.RevokeRole(&_VUB.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VUB *VUBTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VUB.Contract.RevokeRole(&_VUB.TransactOpts, role, account)
}

// SetParams is a paid mutator transaction binding the contract method 0x08617b32.
//
// Solidity: function setParams(uint64 _minLock, uint64 _maxLock, uint256 _maxBoostBps, uint64 _cooldown) returns()
func (_VUB *VUBTransactor) SetParams(opts *bind.TransactOpts, _minLock uint64, _maxLock uint64, _maxBoostBps *big.Int, _cooldown uint64) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "setParams", _minLock, _maxLock, _maxBoostBps, _cooldown)
}

// SetParams is a paid mutator transaction binding the contract method 0x08617b32.
//
// Solidity: function setParams(uint64 _minLock, uint64 _maxLock, uint256 _maxBoostBps, uint64 _cooldown) returns()
func (_VUB *VUBSession) SetParams(_minLock uint64, _maxLock uint64, _maxBoostBps *big.Int, _cooldown uint64) (*types.Transaction, error) {
	return _VUB.Contract.SetParams(&_VUB.TransactOpts, _minLock, _maxLock, _maxBoostBps, _cooldown)
}

// SetParams is a paid mutator transaction binding the contract method 0x08617b32.
//
// Solidity: function setParams(uint64 _minLock, uint64 _maxLock, uint256 _maxBoostBps, uint64 _cooldown) returns()
func (_VUB *VUBTransactorSession) SetParams(_minLock uint64, _maxLock uint64, _maxBoostBps *big.Int, _cooldown uint64) (*types.Transaction, error) {
	return _VUB.Contract.SetParams(&_VUB.TransactOpts, _minLock, _maxLock, _maxBoostBps, _cooldown)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _t) returns()
func (_VUB *VUBTransactor) SetRewardToken(opts *bind.TransactOpts, _t common.Address) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "setRewardToken", _t)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _t) returns()
func (_VUB *VUBSession) SetRewardToken(_t common.Address) (*types.Transaction, error) {
	return _VUB.Contract.SetRewardToken(&_VUB.TransactOpts, _t)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _t) returns()
func (_VUB *VUBTransactorSession) SetRewardToken(_t common.Address) (*types.Transaction, error) {
	return _VUB.Contract.SetRewardToken(&_VUB.TransactOpts, _t)
}

// Stake is a paid mutator transaction binding the contract method 0x952e68cf.
//
// Solidity: function stake(uint256 amount, uint64 dur) returns()
func (_VUB *VUBTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, dur uint64) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "stake", amount, dur)
}

// Stake is a paid mutator transaction binding the contract method 0x952e68cf.
//
// Solidity: function stake(uint256 amount, uint64 dur) returns()
func (_VUB *VUBSession) Stake(amount *big.Int, dur uint64) (*types.Transaction, error) {
	return _VUB.Contract.Stake(&_VUB.TransactOpts, amount, dur)
}

// Stake is a paid mutator transaction binding the contract method 0x952e68cf.
//
// Solidity: function stake(uint256 amount, uint64 dur) returns()
func (_VUB *VUBTransactorSession) Stake(amount *big.Int, dur uint64) (*types.Transaction, error) {
	return _VUB.Contract.Stake(&_VUB.TransactOpts, amount, dur)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_VUB *VUBTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_VUB *VUBSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VUB.Contract.Transfer(&_VUB.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_VUB *VUBTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VUB.Contract.Transfer(&_VUB.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_VUB *VUBTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_VUB *VUBSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VUB.Contract.TransferFrom(&_VUB.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_VUB *VUBTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VUB.Contract.TransferFrom(&_VUB.TransactOpts, from, to, value)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_VUB *VUBTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_VUB *VUBSession) Unstake() (*types.Transaction, error) {
	return _VUB.Contract.Unstake(&_VUB.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_VUB *VUBTransactorSession) Unstake() (*types.Transaction, error) {
	return _VUB.Contract.Unstake(&_VUB.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VUB *VUBTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VUB *VUBSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VUB.Contract.UpgradeToAndCall(&_VUB.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VUB *VUBTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VUB.Contract.UpgradeToAndCall(&_VUB.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_VUB *VUBTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_VUB *VUBSession) Withdraw() (*types.Transaction, error) {
	return _VUB.Contract.Withdraw(&_VUB.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_VUB *VUBTransactorSession) Withdraw() (*types.Transaction, error) {
	return _VUB.Contract.Withdraw(&_VUB.TransactOpts)
}

// VUBApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VUB contract.
type VUBApprovalIterator struct {
	Event *VUBApproval // Event containing the contract specifics and raw log

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
func (it *VUBApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBApproval)
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
		it.Event = new(VUBApproval)
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
func (it *VUBApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBApproval represents a Approval event raised by the VUB contract.
type VUBApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VUB *VUBFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VUBApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VUB.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VUBApprovalIterator{contract: _VUB.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VUB *VUBFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VUBApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VUB.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBApproval)
				if err := _VUB.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_VUB *VUBFilterer) ParseApproval(log types.Log) (*VUBApproval, error) {
	event := new(VUBApproval)
	if err := _VUB.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the VUB contract.
type VUBDelegateChangedIterator struct {
	Event *VUBDelegateChanged // Event containing the contract specifics and raw log

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
func (it *VUBDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBDelegateChanged)
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
		it.Event = new(VUBDelegateChanged)
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
func (it *VUBDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBDelegateChanged represents a DelegateChanged event raised by the VUB contract.
type VUBDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_VUB *VUBFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*VUBDelegateChangedIterator, error) {

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

	logs, sub, err := _VUB.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &VUBDelegateChangedIterator{contract: _VUB.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_VUB *VUBFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *VUBDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VUB.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBDelegateChanged)
				if err := _VUB.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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
func (_VUB *VUBFilterer) ParseDelegateChanged(log types.Log) (*VUBDelegateChanged, error) {
	event := new(VUBDelegateChanged)
	if err := _VUB.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the VUB contract.
type VUBDelegateVotesChangedIterator struct {
	Event *VUBDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *VUBDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBDelegateVotesChanged)
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
		it.Event = new(VUBDelegateVotesChanged)
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
func (it *VUBDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBDelegateVotesChanged represents a DelegateVotesChanged event raised by the VUB contract.
type VUBDelegateVotesChanged struct {
	Delegate      common.Address
	PreviousVotes *big.Int
	NewVotes      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_VUB *VUBFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*VUBDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _VUB.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &VUBDelegateVotesChangedIterator{contract: _VUB.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_VUB *VUBFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *VUBDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _VUB.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBDelegateVotesChanged)
				if err := _VUB.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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
func (_VUB *VUBFilterer) ParseDelegateVotesChanged(log types.Log) (*VUBDelegateVotesChanged, error) {
	event := new(VUBDelegateVotesChanged)
	if err := _VUB.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the VUB contract.
type VUBEIP712DomainChangedIterator struct {
	Event *VUBEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *VUBEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBEIP712DomainChanged)
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
		it.Event = new(VUBEIP712DomainChanged)
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
func (it *VUBEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBEIP712DomainChanged represents a EIP712DomainChanged event raised by the VUB contract.
type VUBEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VUB *VUBFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*VUBEIP712DomainChangedIterator, error) {

	logs, sub, err := _VUB.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &VUBEIP712DomainChangedIterator{contract: _VUB.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VUB *VUBFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *VUBEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _VUB.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBEIP712DomainChanged)
				if err := _VUB.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_VUB *VUBFilterer) ParseEIP712DomainChanged(log types.Log) (*VUBEIP712DomainChanged, error) {
	event := new(VUBEIP712DomainChanged)
	if err := _VUB.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VUB contract.
type VUBInitializedIterator struct {
	Event *VUBInitialized // Event containing the contract specifics and raw log

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
func (it *VUBInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBInitialized)
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
		it.Event = new(VUBInitialized)
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
func (it *VUBInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBInitialized represents a Initialized event raised by the VUB contract.
type VUBInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VUB *VUBFilterer) FilterInitialized(opts *bind.FilterOpts) (*VUBInitializedIterator, error) {

	logs, sub, err := _VUB.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VUBInitializedIterator{contract: _VUB.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VUB *VUBFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VUBInitialized) (event.Subscription, error) {

	logs, sub, err := _VUB.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBInitialized)
				if err := _VUB.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VUB *VUBFilterer) ParseInitialized(log types.Log) (*VUBInitialized, error) {
	event := new(VUBInitialized)
	if err := _VUB.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBLockExtendedIterator is returned from FilterLockExtended and is used to iterate over the raw logs and unpacked data for LockExtended events raised by the VUB contract.
type VUBLockExtendedIterator struct {
	Event *VUBLockExtended // Event containing the contract specifics and raw log

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
func (it *VUBLockExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBLockExtended)
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
		it.Event = new(VUBLockExtended)
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
func (it *VUBLockExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBLockExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBLockExtended represents a LockExtended event raised by the VUB contract.
type VUBLockExtended struct {
	User      common.Address
	NewEnd    uint64
	VubMinted *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLockExtended is a free log retrieval operation binding the contract event 0xc0d15b5a903ff969998bbe8c93fa9476a8374e008475560efd46f25b0b4301a2.
//
// Solidity: event LockExtended(address indexed user, uint64 newEnd, uint256 vubMinted)
func (_VUB *VUBFilterer) FilterLockExtended(opts *bind.FilterOpts, user []common.Address) (*VUBLockExtendedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VUB.contract.FilterLogs(opts, "LockExtended", userRule)
	if err != nil {
		return nil, err
	}
	return &VUBLockExtendedIterator{contract: _VUB.contract, event: "LockExtended", logs: logs, sub: sub}, nil
}

// WatchLockExtended is a free log subscription operation binding the contract event 0xc0d15b5a903ff969998bbe8c93fa9476a8374e008475560efd46f25b0b4301a2.
//
// Solidity: event LockExtended(address indexed user, uint64 newEnd, uint256 vubMinted)
func (_VUB *VUBFilterer) WatchLockExtended(opts *bind.WatchOpts, sink chan<- *VUBLockExtended, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VUB.contract.WatchLogs(opts, "LockExtended", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBLockExtended)
				if err := _VUB.contract.UnpackLog(event, "LockExtended", log); err != nil {
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

// ParseLockExtended is a log parse operation binding the contract event 0xc0d15b5a903ff969998bbe8c93fa9476a8374e008475560efd46f25b0b4301a2.
//
// Solidity: event LockExtended(address indexed user, uint64 newEnd, uint256 vubMinted)
func (_VUB *VUBFilterer) ParseLockExtended(log types.Log) (*VUBLockExtended, error) {
	event := new(VUBLockExtended)
	if err := _VUB.contract.UnpackLog(event, "LockExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBParamsUpdatedIterator is returned from FilterParamsUpdated and is used to iterate over the raw logs and unpacked data for ParamsUpdated events raised by the VUB contract.
type VUBParamsUpdatedIterator struct {
	Event *VUBParamsUpdated // Event containing the contract specifics and raw log

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
func (it *VUBParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBParamsUpdated)
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
		it.Event = new(VUBParamsUpdated)
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
func (it *VUBParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBParamsUpdated represents a ParamsUpdated event raised by the VUB contract.
type VUBParamsUpdated struct {
	MinLock     uint64
	MaxLock     uint64
	MaxBoostBps *big.Int
	Cooldown    uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParamsUpdated is a free log retrieval operation binding the contract event 0x591624d336d15110719b185ac9f817fab3324297374045a5d3fb1a0e88ab9889.
//
// Solidity: event ParamsUpdated(uint64 minLock, uint64 maxLock, uint256 maxBoostBps, uint64 cooldown)
func (_VUB *VUBFilterer) FilterParamsUpdated(opts *bind.FilterOpts) (*VUBParamsUpdatedIterator, error) {

	logs, sub, err := _VUB.contract.FilterLogs(opts, "ParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &VUBParamsUpdatedIterator{contract: _VUB.contract, event: "ParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchParamsUpdated is a free log subscription operation binding the contract event 0x591624d336d15110719b185ac9f817fab3324297374045a5d3fb1a0e88ab9889.
//
// Solidity: event ParamsUpdated(uint64 minLock, uint64 maxLock, uint256 maxBoostBps, uint64 cooldown)
func (_VUB *VUBFilterer) WatchParamsUpdated(opts *bind.WatchOpts, sink chan<- *VUBParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _VUB.contract.WatchLogs(opts, "ParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBParamsUpdated)
				if err := _VUB.contract.UnpackLog(event, "ParamsUpdated", log); err != nil {
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

// ParseParamsUpdated is a log parse operation binding the contract event 0x591624d336d15110719b185ac9f817fab3324297374045a5d3fb1a0e88ab9889.
//
// Solidity: event ParamsUpdated(uint64 minLock, uint64 maxLock, uint256 maxBoostBps, uint64 cooldown)
func (_VUB *VUBFilterer) ParseParamsUpdated(log types.Log) (*VUBParamsUpdated, error) {
	event := new(VUBParamsUpdated)
	if err := _VUB.contract.UnpackLog(event, "ParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the VUB contract.
type VUBRewardAddedIterator struct {
	Event *VUBRewardAdded // Event containing the contract specifics and raw log

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
func (it *VUBRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBRewardAdded)
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
		it.Event = new(VUBRewardAdded)
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
func (it *VUBRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBRewardAdded represents a RewardAdded event raised by the VUB contract.
type VUBRewardAdded struct {
	Amount   *big.Int
	Duration uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xbbb707ba52ee8c7d03c6ff4ddd68fa3d2050fb9fff8f2616f6d5ed4eb5c2e32e.
//
// Solidity: event RewardAdded(uint256 amount, uint64 duration)
func (_VUB *VUBFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*VUBRewardAddedIterator, error) {

	logs, sub, err := _VUB.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &VUBRewardAddedIterator{contract: _VUB.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xbbb707ba52ee8c7d03c6ff4ddd68fa3d2050fb9fff8f2616f6d5ed4eb5c2e32e.
//
// Solidity: event RewardAdded(uint256 amount, uint64 duration)
func (_VUB *VUBFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *VUBRewardAdded) (event.Subscription, error) {

	logs, sub, err := _VUB.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBRewardAdded)
				if err := _VUB.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xbbb707ba52ee8c7d03c6ff4ddd68fa3d2050fb9fff8f2616f6d5ed4eb5c2e32e.
//
// Solidity: event RewardAdded(uint256 amount, uint64 duration)
func (_VUB *VUBFilterer) ParseRewardAdded(log types.Log) (*VUBRewardAdded, error) {
	event := new(VUBRewardAdded)
	if err := _VUB.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the VUB contract.
type VUBRewardPaidIterator struct {
	Event *VUBRewardPaid // Event containing the contract specifics and raw log

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
func (it *VUBRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBRewardPaid)
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
		it.Event = new(VUBRewardPaid)
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
func (it *VUBRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBRewardPaid represents a RewardPaid event raised by the VUB contract.
type VUBRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_VUB *VUBFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*VUBRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VUB.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &VUBRewardPaidIterator{contract: _VUB.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_VUB *VUBFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *VUBRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VUB.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBRewardPaid)
				if err := _VUB.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_VUB *VUBFilterer) ParseRewardPaid(log types.Log) (*VUBRewardPaid, error) {
	event := new(VUBRewardPaid)
	if err := _VUB.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VUB contract.
type VUBRoleAdminChangedIterator struct {
	Event *VUBRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VUBRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBRoleAdminChanged)
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
		it.Event = new(VUBRoleAdminChanged)
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
func (it *VUBRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBRoleAdminChanged represents a RoleAdminChanged event raised by the VUB contract.
type VUBRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VUB *VUBFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VUBRoleAdminChangedIterator, error) {

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

	logs, sub, err := _VUB.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VUBRoleAdminChangedIterator{contract: _VUB.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VUB *VUBFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VUBRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _VUB.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBRoleAdminChanged)
				if err := _VUB.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_VUB *VUBFilterer) ParseRoleAdminChanged(log types.Log) (*VUBRoleAdminChanged, error) {
	event := new(VUBRoleAdminChanged)
	if err := _VUB.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VUB contract.
type VUBRoleGrantedIterator struct {
	Event *VUBRoleGranted // Event containing the contract specifics and raw log

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
func (it *VUBRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBRoleGranted)
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
		it.Event = new(VUBRoleGranted)
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
func (it *VUBRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBRoleGranted represents a RoleGranted event raised by the VUB contract.
type VUBRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VUB *VUBFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VUBRoleGrantedIterator, error) {

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

	logs, sub, err := _VUB.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VUBRoleGrantedIterator{contract: _VUB.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VUB *VUBFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VUBRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VUB.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBRoleGranted)
				if err := _VUB.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_VUB *VUBFilterer) ParseRoleGranted(log types.Log) (*VUBRoleGranted, error) {
	event := new(VUBRoleGranted)
	if err := _VUB.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VUB contract.
type VUBRoleRevokedIterator struct {
	Event *VUBRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VUBRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBRoleRevoked)
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
		it.Event = new(VUBRoleRevoked)
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
func (it *VUBRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBRoleRevoked represents a RoleRevoked event raised by the VUB contract.
type VUBRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VUB *VUBFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VUBRoleRevokedIterator, error) {

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

	logs, sub, err := _VUB.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VUBRoleRevokedIterator{contract: _VUB.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VUB *VUBFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VUBRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VUB.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBRoleRevoked)
				if err := _VUB.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_VUB *VUBFilterer) ParseRoleRevoked(log types.Log) (*VUBRoleRevoked, error) {
	event := new(VUBRoleRevoked)
	if err := _VUB.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the VUB contract.
type VUBStakedIterator struct {
	Event *VUBStaked // Event containing the contract specifics and raw log

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
func (it *VUBStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBStaked)
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
		it.Event = new(VUBStaked)
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
func (it *VUBStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBStaked represents a Staked event raised by the VUB contract.
type VUBStaked struct {
	User      common.Address
	UbAmount  *big.Int
	LockEnd   uint64
	VubMinted *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x0bcadf4a8215096a7cbb695c9385c84784ef6d32892221d088a4f13aa8d49196.
//
// Solidity: event Staked(address indexed user, uint256 ubAmount, uint64 lockEnd, uint256 vubMinted)
func (_VUB *VUBFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*VUBStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VUB.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &VUBStakedIterator{contract: _VUB.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x0bcadf4a8215096a7cbb695c9385c84784ef6d32892221d088a4f13aa8d49196.
//
// Solidity: event Staked(address indexed user, uint256 ubAmount, uint64 lockEnd, uint256 vubMinted)
func (_VUB *VUBFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *VUBStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VUB.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBStaked)
				if err := _VUB.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x0bcadf4a8215096a7cbb695c9385c84784ef6d32892221d088a4f13aa8d49196.
//
// Solidity: event Staked(address indexed user, uint256 ubAmount, uint64 lockEnd, uint256 vubMinted)
func (_VUB *VUBFilterer) ParseStaked(log types.Log) (*VUBStaked, error) {
	event := new(VUBStaked)
	if err := _VUB.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VUB contract.
type VUBTransferIterator struct {
	Event *VUBTransfer // Event containing the contract specifics and raw log

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
func (it *VUBTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBTransfer)
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
		it.Event = new(VUBTransfer)
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
func (it *VUBTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBTransfer represents a Transfer event raised by the VUB contract.
type VUBTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VUB *VUBFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VUBTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VUB.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VUBTransferIterator{contract: _VUB.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VUB *VUBFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VUBTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VUB.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBTransfer)
				if err := _VUB.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_VUB *VUBFilterer) ParseTransfer(log types.Log) (*VUBTransfer, error) {
	event := new(VUBTransfer)
	if err := _VUB.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the VUB contract.
type VUBUnstakedIterator struct {
	Event *VUBUnstaked // Event containing the contract specifics and raw log

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
func (it *VUBUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBUnstaked)
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
		it.Event = new(VUBUnstaked)
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
func (it *VUBUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBUnstaked represents a Unstaked event raised by the VUB contract.
type VUBUnstaked struct {
	User     common.Address
	UbAmount *big.Int
	Ready    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x536c53e11db8105c787d8d5fce8b01f689aefd57771dad0d0c62c33af2ecc1f9.
//
// Solidity: event Unstaked(address indexed user, uint256 ubAmount, uint64 ready)
func (_VUB *VUBFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*VUBUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VUB.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &VUBUnstakedIterator{contract: _VUB.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x536c53e11db8105c787d8d5fce8b01f689aefd57771dad0d0c62c33af2ecc1f9.
//
// Solidity: event Unstaked(address indexed user, uint256 ubAmount, uint64 ready)
func (_VUB *VUBFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *VUBUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VUB.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBUnstaked)
				if err := _VUB.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x536c53e11db8105c787d8d5fce8b01f689aefd57771dad0d0c62c33af2ecc1f9.
//
// Solidity: event Unstaked(address indexed user, uint256 ubAmount, uint64 ready)
func (_VUB *VUBFilterer) ParseUnstaked(log types.Log) (*VUBUnstaked, error) {
	event := new(VUBUnstaked)
	if err := _VUB.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VUB contract.
type VUBUpgradedIterator struct {
	Event *VUBUpgraded // Event containing the contract specifics and raw log

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
func (it *VUBUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBUpgraded)
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
		it.Event = new(VUBUpgraded)
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
func (it *VUBUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBUpgraded represents a Upgraded event raised by the VUB contract.
type VUBUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VUB *VUBFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VUBUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VUB.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VUBUpgradedIterator{contract: _VUB.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VUB *VUBFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VUBUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VUB.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBUpgraded)
				if err := _VUB.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_VUB *VUBFilterer) ParseUpgraded(log types.Log) (*VUBUpgraded, error) {
	event := new(VUBUpgraded)
	if err := _VUB.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VUBWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the VUB contract.
type VUBWithdrawnIterator struct {
	Event *VUBWithdrawn // Event containing the contract specifics and raw log

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
func (it *VUBWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBWithdrawn)
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
		it.Event = new(VUBWithdrawn)
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
func (it *VUBWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBWithdrawn represents a Withdrawn event raised by the VUB contract.
type VUBWithdrawn struct {
	User     common.Address
	UbAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 ubAmount)
func (_VUB *VUBFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*VUBWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VUB.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &VUBWithdrawnIterator{contract: _VUB.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 ubAmount)
func (_VUB *VUBFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *VUBWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VUB.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBWithdrawn)
				if err := _VUB.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
// Solidity: event Withdrawn(address indexed user, uint256 ubAmount)
func (_VUB *VUBFilterer) ParseWithdrawn(log types.Log) (*VUBWithdrawn, error) {
	event := new(VUBWithdrawn)
	if err := _VUB.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
