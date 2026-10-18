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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_BOOST_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_COOLDOWN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REWARD_FUNDER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"boostBps\",\"inputs\":[{\"name\":\"dur\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pos\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCheckpoints.Checkpoint208\",\"components\":[{\"name\":\"_key\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"_value\",\"type\":\"uint208\",\"internalType\":\"uint208\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateBySig\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegates\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"earned\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"extendLock\",\"inputs\":[{\"name\":\"newDur\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"minted\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPastTotalSupply\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPastVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReward\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_ub\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastTimeRewardApplicable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastUpdate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"locks\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxBoostBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxLock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migratePrincipalAccounting\",\"inputs\":[{\"name\":\"_totalStaked\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"migratePrincipalAndUpgrader\",\"inputs\":[{\"name\":\"_totalStaked\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"upgrader\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"migrateUpgrader\",\"inputs\":[{\"name\":\"upgrader\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minLock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"notifyReward\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pending\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ready\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"periodFinish\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"principalMigrated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardPerToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardPerTokenStored\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setParams\",\"inputs\":[{\"name\":\"_minLock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_maxLock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_maxBoostBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_cooldown\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardToken\",\"inputs\":[{\"name\":\"_t\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dur\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopReward\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalStaked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ub\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"userRewardPerTokenPaid\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateChanged\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fromDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateVotesChanged\",\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"previousVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockExtended\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newEnd\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"vubMinted\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParamsUpdated\",\"inputs\":[{\"name\":\"minLock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"maxLock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"maxBoostBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"cooldown\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardAdded\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardPaid\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardStopped\",\"inputs\":[{\"name\":\"unstreamed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ubAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lockEnd\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"vubMinted\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unstaked\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ubAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ready\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ubAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20ExceededSafeSupply\",\"inputs\":[{\"name\":\"increasedSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC5805FutureLookup\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clock\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"ERC6372InconsistentClock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"VotesExpiredSignature\",\"inputs\":[{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a08060405234620000d157306080527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c16620000c257506001600160401b036002600160401b0319828216016200007c575b6040516151f39081620000d6823960805181818161141a0152612ddb0152f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80806200005c565b63f92ee8a960e01b8152600490fd5b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80628cc2621461048357806301ffc9a71461047e57806306fdde03146104795780630700037d1461047457806308617b321461046f578063095ea7b31461046a57806315456eba1461046557806318160ddd14610460578063182f27191461045b57806323b872dd14610456578063248a9ca3146104515780632def66201461044c5780632f2ff15d14610447578063313ce5671461044257806336568abe1461043d5780633a46b1a8146104385780633cbc963d146104335780633ccfd60b1461042e5780633d18b912146104295780633d9a581a146104245780634bf5d7e91461041f5780634f1ef2861461041a57806352d1902d14610415578063584fb91014610410578063587cde1e1461040b5780635c19a95c146104065780635de9a137146104015780635eebea20146103fc5780636c0b3e46146103f75780636fcfff45146103f257806370a08231146103ed5780637502e24e146103e8578063787a08a6146103e35780637b0a47ee146103de5780637ecebe00146103d957806380dc0672146103d457806380faa57d146103cf578063817b1cd2146103ca57806384b0196e146103c5578063859e6d6a146103c05780638aee8127146103bb5780638b41d35f146103b65780638b876347146103b15780638e539e8c146103ac57806391d14854146103a757806391ddadf4146103a2578063952e68cf1461039d57806395d89b41146103985780639ab24eb014610393578063a09f11431461038e578063a217fddf14610389578063a9059cbb14610384578063ad3cb1cc1461037f578063b49c71651461037a578063c046371114610375578063c0c53b8b14610370578063c3cda5201461036b578063ccc5749014610366578063cd3daf9d14610361578063d547741f1461035c578063dd62ed3e14610357578063de350feb14610352578063df136d651461034d578063e3d8900714610348578063ebe2b12b14610343578063f0090cf61461033e578063f037c63014610339578063f1127ed814610334578063f72c0d8b1461032f578063f7c618c11461032a5763ffa1ad7414610325575f80fd5b612830565b612808565b6127e1565b61274e565b612725565b6126eb565b6126c5565b6126a8565b61268b565b61254d565b612515565b6124cc565b6124b2565b612478565b6123aa565b612290565b612267565b6121d2565b61218d565b6120f3565b6120d9565b6120b2565b61207b565b611fc9565b611e6e565b611e43565b611de9565b611d14565b611cdc565b611cbf565b611c34565b611a90565b6119c0565b61191d565b6118f2565b6117fb565b6117a4565b611787565b611761565b611743565b6116ec565b611697565b611671565b61160e565b6115ab565b611589565b611544565b611471565b611408565b611389565b611279565b61125d565b61110d565b61101d565b610ffb565b610f29565b610ee2565b610ec7565b610e7c565b610d2d565b610cf4565b610c56565b610b7b565b610b52565b6109f9565b61093b565b6107b3565b610739565b6105a1565b6104f9565b6104ce565b600435906001600160a01b038216820361049e57565b5f80fd5b602435906001600160a01b038216820361049e57565b604435906001600160a01b038216820361049e57565b3461049e57602036600319011261049e5760206104f16104ec610488565b6128f0565b604051908152f35b3461049e57602036600319011261049e5760043563ffffffff60e01b811680910361049e57602090637965db0b60e01b811490811561053e575b506040519015158152f35b6301ffc9a760e01b1490505f610533565b91908251928382525f5b848110610579575050825f602080949584010152601f8019910116010190565b602081830181015184830182015201610559565b90602061059e92818152019061054f565b90565b3461049e575f36600319011261049e576040515f5f8051602061505e8339815191528054906105cf82612950565b80855291602091600191828116908115610664575060011461060c575b610608866105fc81880382611359565b6040519182918261058d565b0390f35b5f90815293507f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab05b838510610651575050505081016020016105fc826106085f6105ec565b8054868601840152938201938101610634565b9050869550610608969350602092506105fc94915060ff191682840152151560051b82010192935f6105ec565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace006020526040902090565b6001600160a01b03165f9081527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d016020526040902090565b3461049e57602036600319011261049e576001600160a01b0361075a610488565b165f52600b602052602060405f2054604051908152f35b600435906001600160401b038216820361049e57565b602435906001600160401b038216820361049e57565b606435906001600160401b038216820361049e57565b3461049e57608036600319011261049e576107cc610771565b6107d4610787565b604435916107e061079d565b6107e8613499565b6001600160401b0390818316801515908161092e575b50156108f8578461085662278d006108f39461083f6127107f591624d336d15110719b185ac9f817fab3324297374045a5d3fb1a0e88ab98899a1015612b07565b61084d61c350851115612b40565b84161115612b7a565b6001805467ffffffffffffffff60a01b191660a086901b67ffffffffffffffff60a01b1617905561089d856001600160401b03166001600160401b03196002541617600255565b6108a681600355565b6108c6826001600160401b03166001600160401b03196004541617600455565b604080516001600160401b0395861681529585166020870152850152909116606083015281906080820190565b0390a1005b60405162461bcd60e51b815260206004820152600e60248201526d626164206c6f636b2072616e676560901b6044820152606490fd5b905082851610155f6107fe565b3461049e57604036600319011261049e57610954610488565b60243533156109e1576001600160a01b0382169182156109c957610991829161097c33610691565b9060018060a01b03165f5260205260405f2090565b556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b604051634a1406b160e11b81525f6004820152602490fd5b60405163e602df0560e01b81525f6004820152602490fd5b3461049e57602036600319011261049e57335f908152600560205260409020600435907f0bcadf4a8215096a7cbb695c9385c84784ef6d32892221d088a4f13aa8d4919690610a49831515612bb7565b8054151580610b27575b610a5c90612bf1565b610a653361361a565b610b0060018201610aac610aa4610a9e610a99610a8985546001600160401b031690565b6001600160401b03421690612c2e565b612ee5565b876128a4565b612710900490565b92610ab88682546128e3565b9055610ace610ac986600c546128e3565b600c55565b610ad88333613669565b5f54610af39086906001600160a01b03165b30903390613723565b546001600160401b031690565b604080519485526001600160401b039091166020850152830152339180606081015b0390a2005b50610a5c610b3f60018301546001600160401b031690565b6001600160401b03429116119050610a53565b3461049e575f36600319011261049e5760205f805160206150be83398151915254604051908152f35b3461049e57602036600319011261049e57610b94613499565b5f8051602061517e833981519152805460ff8160401c16908115610c41575b50610c2f575f8051602061517e833981519152805467ffffffffffffffff19166002179055805460ff60401b1916600160401b179055610bf46004356137b4565b610bfc612c47565b604051600281527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29080602081016108f3565b60405163f92ee8a960e01b8152600490fd5b600291506001600160401b031610155f610bb3565b3461049e57606036600319011261049e57610c6f610488565b610c776104a2565b90604435610c883361097c84610691565b545f198110610c99575b5050613807565b818110610ccc5750506001600160a01b038116156109e15733156109c957610cc43361097c83610691565b505f80610c92565b604051637dc7a0d960e11b815233600482015260248101919091526044810191909152606490fd5b3461049e57602036600319011261049e576004355f525f8051602061515e8339815191526020526020600160405f200154604051908152f35b3461049e575f36600319011261049e57335f908152600560205260409020610d5781541515612c62565b610d85610d7d610d7160018401546001600160401b031690565b6001600160401b031690565b421015612c98565b610d8e3361361a565b54335f908152600560205260409020610dad905b60015f918281550155565b610db6336106c9565b5480610e6c575b50335f9081526006602052604090207f536c53e11db8105c787d8d5fce8b01f689aefd57771dad0d0c62c33af2ecc1f990610b2290610e4b90610e018582546128e3565b8155610af36001610e2c610e1d6004546001600160401b031690565b6001600160401b034216612ccd565b92019182906001600160401b03166001600160401b0319825416179055565b604080519485526001600160401b0390911660208501523393918291820190565b610e76903361385b565b5f610dbd565b3461049e57604036600319011261049e57610ec5600435610e9b6104a2565b90805f525f8051602061515e833981519152602052610ec0600160405f2001546135af565b613bd6565b005b3461049e575f36600319011261049e57602060405160128152f35b3461049e57604036600319011261049e57610efb6104a2565b336001600160a01b03821603610f1757610ec590600435613c0c565b60405163334bd91960e11b8152600490fd5b3461049e57604036600319011261049e57610f4a610f45610488565b610701565b610f55602435613c9e565b8154905f829160058411610fa8575b610f6f9350846143d1565b9081610f8d57505060205f5b6040516001600160d01b039091168152f35b610f98602092612889565b905f52815f20015460301c610f7b565b9192610fb381614234565b8103908111610ff657610f6f93855f5265ffffffffffff808360205f20015416908516105f14610fe4575091610f64565b929150610ff0906128d5565b90610f64565b612875565b3461049e575f36600319011261049e57602060ff600d54166040519015158152f35b3461049e575f36600319011261049e57335f52600660205260405f2080549081156110d65761105c610d7160016110649301546001600160401b031690565b421015612ce8565b335f90815260066020526040902061107b90610da2565b61108a610ac982600c54612897565b5f546110a290829033906001600160a01b0316613cdf565b60405190815233907f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5908060208101610b22565b60405162461bcd60e51b815260206004820152600f60248201526e6e6f7468696e672070656e64696e6760881b6044820152606490fd5b3461049e575f36600319011261049e576111263361361a565b335f908152600b60205260409020548061113c57005b335f908152600b60205260408120556001546001600160a01b03165f546001600160a01b039183916111769084166001600160a01b031690565b92808416908216146111c3575b61118f92503390613cdf565b60405190815233907fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486908060208101610b22565b90506111d96111d4600d5460ff1690565b612d23565b6040516370a0823160e01b815230600482015291602090839060249082905afa90811561125857611224849261118f945f91611229575b5061121d600c54856128e3565b1115612d82565b611183565b61124b915060203d602011611251575b6112438183611359565b810190612d68565b5f611210565b503d611239565b612d77565b3461049e575f36600319011261049e57602060405161c3508152f35b3461049e575f36600319011261049e5761129243614202565b65ffffffffffff806112a343614202565b169116036112f8576106086040516112ba8161131e565b601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015260405191829160208352602083019061054f565b6040516301bfc1c560e61b8152600490fd5b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b0382111761133957604052565b61130a565b60a081019081106001600160401b0382111761133957604052565b90601f801991011681019081106001600160401b0382111761133957604052565b604051906113878261131e565b565b604036600319011261049e5761139d610488565b602435906001600160401b039081831161049e573660238401121561049e57826004013591821161133957604051916113e0601f8201601f191660200184611359565b808352366024828601011161049e576020815f926024610ec597018387013784010152612dce565b3461049e575f36600319011261049e577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361145f5760206040515f8051602061511e8339815191528152f35b60405163703e46dd60e11b8152600490fd5b3461049e57604036600319011261049e5761148a6104a2565b611492613499565b5f8051602061517e83398151915290815460ff8160401c1690811561152f575b50610c2f575f8051602061517e833981519152805467ffffffffffffffff191660031790556114f491805460ff60401b1916600160401b179055600435612ed3565b6114fc612c47565b604051600381527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29080602081016108f3565b600391506001600160401b031610155f6114b2565b3461049e57602036600319011261049e5760206001600160a01b0380611568610488565b165f525f8051602061503e833981519152825260405f205416604051908152f35b3461049e57602036600319011261049e57610ec56115a5610488565b33613e0f565b3461049e57602036600319011261049e576001600160a01b036115cc610488565b165f52600560205260405f206001600160401b03600182549201541690610608604051928392839092916001600160401b036020916040840195845216910152565b3461049e57602036600319011261049e576001600160a01b0361162f610488565b165f52600660205260405f206001600160401b03600182549201541690610608604051928392839092916001600160401b036020916040840195845216910152565b3461049e575f36600319011261049e5760206001600160401b0360025416604051908152f35b3461049e57602036600319011261049e576116b3610f45610488565b5463ffffffff908181116116cd5760209160405191168152f35b604490604051906306dfcc6560e41b8252602060048301526024820152fd5b3461049e57602036600319011261049e5760206104f161170a610488565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b3461049e57602036600319011261049e5760206104f1610a99610771565b3461049e575f36600319011261049e5760206001600160401b0360045416604051908152f35b3461049e575f36600319011261049e576020600754604051908152f35b3461049e57602036600319011261049e576001600160a01b036117c5610488565b165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052602060405f2054604051908152f35b3461049e575f36600319011261049e57611813613514565b6001600160401b038042169080600854168210156118ba577f278b7aeb8d452471692dbcef95ab699085487ea0c4f3e5e2d13c84795a0f1ae2916118aa6118846108f39361185f6135db565b61187a846118756008546001600160401b031690565b612c2e565b60075491166128a4565b9161188e5f600755565b6001600160401b03166001600160401b03196008541617600855565b6040519081529081906020820190565b60405162461bcd60e51b815260206004820152601060248201526f6e6f206163746976652073747265616d60801b6044820152606490fd5b3461049e575f36600319011261049e57602061190c612f6c565b6001600160401b0360405191168152f35b3461049e575f36600319011261049e576020600c54604051908152f35b9161196e90949194600f60f81b845261196060209660e0602087015260e086019061054f565b90848203604086015261054f565b92606083015260018060a01b031660808201525f60a082015260c0818303910152602080845192838152019301915f5b8281106119ac575050505090565b83518552938101939281019260010161199e565b3461049e575f36600319011261049e577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100541580611a67575b15611a2a57611a06612988565b611a0e612a5b565b90610608611a1a612f8e565b604051938493309146918661193a565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10154156119f9565b3461049e57602036600319011261049e57610608611aac610771565b335f908152600560205260409020611ac681541515612bf1565b611b51610aa4611ae5610d716001546001600160401b039060a01c1690565b93611b066001600160401b0395868316908110159081611c15575b50612fb4565b611b4b611b1582874216612ccd565b94611b3b6001820197611b32610d718a546001600160401b031690565b90881611612feb565b611b443361361a565b5491612ee5565b906128a4565b81611b5b336106c9565b549182811115611bf257611b8e92611b7291612897565b93906001600160401b03166001600160401b0319825416179055565b81611be3575b604080516001600160401b039290921682526020820183905233917fc0d15b5a903ff969998bbe8c93fa9476a8374e008475560efd46f25b0b4301a29190a26040519081529081906020820190565b611bed8233613669565b611b94565b50611b8e91505f93906001600160401b03166001600160401b0319825416179055565b9050611c2c610d716002546001600160401b031690565b10155f611b00565b3461049e57602036600319011261049e57611c4d610488565b611c55613499565b6001600160401b0360085416421115611c8a57600180546001600160a01b0319166001600160a01b0392909216919091179055005b60405162461bcd60e51b815260206004820152600d60248201526c7265776172642061637469766560981b6044820152606490fd5b3461049e575f36600319011261049e57602060405162278d008152f35b3461049e57602036600319011261049e576001600160a01b03611cfd610488565b165f52600a602052602060405f2054604051908152f35b3461049e57602036600319011261049e57611d30600435613c9e565b5f8051602061513e833981519152908154905f829160058411611d91575b611d58935061435f565b9081611d6b5750506040515f8152602090f35b611d76602092612889565b905f525f8051602061519e833981519152015460301c610f7b565b9192611d9c81614234565b8103908111610ff657611d5893855f5265ffffffffffff80835f8051602061519e833981519152015416908516105f14611dd7575091611d4e565b929150611de3906128d5565b90611d4e565b3461049e57604036600319011261049e57602060ff611e37611e096104a2565b6004355f525f8051602061515e833981519152845260405f209060018060a01b03165f5260205260405f2090565b54166040519015158152f35b3461049e575f36600319011261049e576020611e5e43614202565b65ffffffffffff60405191168152f35b3461049e57604036600319011261049e576004357f0bcadf4a8215096a7cbb695c9385c84784ef6d32892221d088a4f13aa8d49196611eab610787565b611eb6831515612bb7565b611f30610aa4611f2a611f2484611edc610d716001546001600160401b039060a01c1690565b611efb6001600160401b0391828416908110159081611c155750612fb4565b335f908152600560205260409020611f14905415613024565b611f1d3361361a565b4216612ccd565b93612ee5565b856128a4565b90611f61611f3c61137a565b8581526001600160401b0383166020820152335f90815260056020526040902061305e565b611f70610ac985600c546128e3565b611f7a8233613669565b335f9081525f8051602061503e83398151915260205260409020546001600160a01b031615611fba575b5f54610b009085906001600160a01b0316610aea565b611fc43333613e0f565b611fa4565b3461049e575f36600319011261049e576040515f5f8051602061509e833981519152805490611ff782612950565b80855291602091600191828116908115610664575060011461202357610608866105fc81880382611359565b5f90815293507f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa5b838510612068575050505081016020016105fc826106085f6105ec565b805486860184015293820193810161204b565b3461049e57602036600319011261049e5760206001600160d01b036120a96120a4610f45610488565b613f0c565b16604051908152f35b3461049e575f36600319011261049e575f546040516001600160a01b039091168152602090f35b3461049e575f36600319011261049e5760206040515f8152f35b3461049e57604036600319011261049e5761210c610488565b3315612175576001600160a01b03161561215d5760405162461bcd60e51b81526020600482015260156024820152747655423a206e6f6e2d7472616e7366657261626c6560581b6044820152606490fd5b60405163ec442f0560e01b81525f6004820152602490fd5b604051634b637e8f60e11b81525f6004820152602490fd5b3461049e575f36600319011261049e576106086040516121ac8161131e565b60058152640352e302e360dc1b602082015260405191829160208352602083019061054f565b3461049e57602036600319011261049e576121eb610488565b6121f3613499565b5f8051602061517e83398151915290815460ff8160401c16908115612252575b50610c2f575f8051602061517e833981519152805467ffffffffffffffff191660031790556114f491805460ff60401b1916600160401b179055613092565b600391506001600160401b031610155f612213565b3461049e575f36600319011261049e5760206008546001600160401b036040519160401c168152f35b3461049e57606036600319011261049e576122a9610488565b6122b16104a2565b906122ba6104b8565b5f8051602061517e83398151915254926001600160401b0360ff8560401c16159416801590816123a2575b6001149081612398575b15908161238f575b50610c2f575f8051602061517e833981519152805467ffffffffffffffff1916600117905561232a928461236b57613127565b61233057005b612338612c47565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29080602081016108f3565b5f8051602061517e833981519152805460ff60401b1916600160401b179055613127565b9050155f6122f7565b303b1591506122ef565b8591506122e5565b3461049e5760c036600319011261049e576123c3610488565b6044359060243560643560ff8116810361049e5783421161245f57612453610ec59461245a926040519060208201927fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf845260018060a01b038816604084015286606084015260808301526080825261243b8261133e565b61244e60a43593608435935190206140f2565b614181565b9182614199565b613e0f565b604051632341d78760e11b815260048101859052602490fd5b3461049e575f36600319011261049e5760206040517f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f558152f35b3461049e575f36600319011261049e5760206104f1613391565b3461049e57604036600319011261049e57610ec56004356124eb6104a2565b90805f525f8051602061515e833981519152602052612510600160405f2001546135af565b613c0c565b3461049e57604036600319011261049e576020612544612533610488565b61097c61253e6104a2565b91610691565b54604051908152f35b3461049e57604036600319011261049e577fbbb707ba52ee8c7d03c6ff4ddd68fa3d2050fb9fff8f2616f6d5ed4eb5c2e32e60043561258a610787565b90612593613514565b61262b61188e836001600160401b038082166125b0811515613406565b6125b86135db565b81421691826125cf6008546001600160401b031690565b8083168210612663575050506125e86125ed91876128b7565b600755565b6125fa6007541515613442565b600880546fffffffffffffffff00000000000000001916604083901b67ffffffffffffffff60401b16179055612ccd565b6001546126429082906001600160a01b0316610aea565b604080519182526001600160401b03909216602082015290819081016108f3565b61267b6125e89361187a612686969461268194612c2e565b896128e3565b6128b7565b6125ed565b3461049e575f36600319011261049e576020600954604051908152f35b3461049e575f36600319011261049e576020600354604051908152f35b3461049e575f36600319011261049e5760206001600160401b0360085416604051908152f35b3461049e575f36600319011261049e5760206040517f61b754da92d8d8d7300489a35a466b9ed19cf4a61860a290f89bec3a75de2bcf8152f35b3461049e575f36600319011261049e5760206001600160401b0360015460a01c16604051908152f35b3461049e57604036600319011261049e57612767610488565b6024359063ffffffff8216820361049e576040916127996127a79261278a613481565b50612793613481565b50610701565b6127a1613481565b50614b1a565b508151906127b48261131e565b54602065ffffffffffff821692838152019060301c8152825191825260018060d01b039051166020820152f35b3461049e575f36600319011261049e5760206040515f805160206150fe8339815191528152f35b3461049e575f36600319011261049e576001546040516001600160a01b039091168152602090f35b3461049e575f36600319011261049e5761060860405161284f8161131e565b60058152640312e302e360dc1b602082015260405191829160208352602083019061054f565b634e487b7160e01b5f52601160045260245ffd5b5f19810191908211610ff657565b91908203918211610ff657565b81810292918115918404141715610ff657565b81156128c1570490565b634e487b7160e01b5f52601260045260245ffd5b9060018201809211610ff657565b91908201809211610ff657565b6128f9816106c9565b5490612903613391565b9060018060a01b031691825f52600a60205260405f20548203918211610ff657670de0b6b3a764000091612936916128a4565b04905f52600b60205260405f20548101809111610ff65790565b90600182811c9216801561297e575b602083101461296a57565b634e487b7160e01b5f52602260045260245ffd5b91607f169161295f565b604051905f825f8051602061507e833981519152918254926129a984612950565b80845293602091600191828116908115612a3557506001146129d5575b50505061138792500383611359565b5f9081527f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d9590935091905b828410612a1d57506113879450505081016020015f80806129c6565b85548885018301529485019487945092810192612a01565b925050506020925061138794915060ff191682840152151560051b8201015f80806129c6565b604051905f825f805160206150de83398151915291825492612a7c84612950565b80845293602091600191828116908115612a355750600114612aa75750505061138792500383611359565b5f9081527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b759590935091905b828410612aef57506113879450505081016020015f80806129c6565b85548885018301529485019487945092810192612ad3565b15612b0e57565b60405162461bcd60e51b815260206004820152600a6024820152690c4dedee6e840784062f60b31b6044820152606490fd5b15612b4757565b60405162461bcd60e51b815260206004820152600b60248201526a0626f6f7374203e206361760ac1b6044820152606490fd5b15612b8157565b60405162461bcd60e51b815260206004820152600e60248201526d0636f6f6c646f776e203e206361760941b6044820152606490fd5b15612bbe57565b60405162461bcd60e51b815260206004820152600b60248201526a1e995c9bc8185b5bdd5b9d60aa1b6044820152606490fd5b15612bf857565b60405162461bcd60e51b815260206004820152600e60248201526d6e6f20616374697665206c6f636b60901b6044820152606490fd5b6001600160401b039182169082160391908211610ff657565b5f8051602061517e833981519152805460ff60401b19169055565b15612c6957565b60405162461bcd60e51b81526020600482015260076024820152666e6f206c6f636b60c81b6044820152606490fd5b15612c9f57565b60405162461bcd60e51b81526020600482015260066024820152651b1bd8dad95960d21b6044820152606490fd5b9190916001600160401b0380809416911601918211610ff657565b15612cef57565b60405162461bcd60e51b815260206004820152600c60248201526b31b7b7b634b733903237bbb760a11b6044820152606490fd5b15612d2a57565b60405162461bcd60e51b81526020600482015260166024820152751c1c9a5b98da5c185b081b9bdd081b5a59dc985d195960521b6044820152606490fd5b9081602091031261049e575190565b6040513d5f823e3d90fd5b15612d8957565b60405162461bcd60e51b815260206004820152601c60248201527f72657761726420776f756c6420746f756368207072696e636970616c000000006044820152606490fd5b6001600160a01b039290917f00000000000000000000000000000000000000000000000000000000000000008416308114908115612eb8575b5061145f576020600494612e1961356b565b6040516352d1902d60e01b8152958691829087165afa5f9481612e97575b50612e5d57604051634c9c8ce360e01b81526001600160a01b0384166004820152602490fd5b90915f8051602061511e8339815191528403612e7e57611387929350614430565b604051632a87526960e21b815260048101859052602490fd5b612eb191955060203d602011611251576112438183611359565b935f612e37565b9050845f8051602061511e833981519152541614155f612e07565b90612ee0611387926137b4565b613d3e565b6001546001600160401b038281169160a01c811680831115612f6257816002541680931015612f585760035461270f19810191908211610ff657612f3b612f419284612f3484612f4899612c2e565b16906128a4565b93612c2e565b16906128b7565b612710908101809111610ff65790565b5050505060035490565b5050505061271090565b6001600160401b0380421690600854168082105f14612f89575090565b905090565b604051602081018181106001600160401b03821117611339576040525f8152905f368137565b15612fbb57565b60405162461bcd60e51b8152602060048201526008602482015267626164206c6f636b60c01b6044820152606490fd5b15612ff257565b60405162461bcd60e51b815260206004820152600a6024820152693737ba103637b733b2b960b11b6044820152606490fd5b1561302b57565b60405162461bcd60e51b815260206004820152600b60248201526a616374697665206c6f636b60a81b6044820152606490fd5b60016001600160401b03602061138794805185550151169101906001600160401b03166001600160401b0319825416179055565b60ff600d5416156130a65761138790613d3e565b60405162461bcd60e51b815260206004820152601760248201527f6d696772617465207072696e636970616c2066697273740000000000000000006044820152606490fd5b604051906130f88261131e565b60038252623b2aa160e91b6020830152565b604051906131178261131e565b60018252603160f81b6020830152565b906040516131348161131e565b601081526020906f2b37ba3296b2b9b1b937bbb2b2102aa160811b602082015261315c6130eb565b91613165614606565b61316d614606565b8151906001600160401b038211611339575f8051602061505e833981519152926131a08361319b8654612950565b614634565b602091601f84116001146132f35750926131df836131e6946132e59a99979461324d99975f926132e8575b50508160011b915f199060031b1c19161790565b9055614818565b6131ee614606565b6132076131f96130eb565b61320161310a565b90613f34565b61320f614606565b5f80546001600160a01b0319166001600160a01b039384161790551660018060a01b03166bffffffffffffffffffffffff60a01b6001541617600155565b6001805467ffffffffffffffff60a01b191661127560a71b1790556132836303c267006001600160401b03196002541617600255565b61328e6161a8600355565b6132a862093a806001600160401b03196004541617600455565b6132ba600160ff19600d541617600d55565b6132c381613984565b506132cd81613a35565b506132d781613b0b565b506132e061409d565b613b7a565b50565b015190505f806131cb565b5f8051602061505e8339815191525f529190601f1984167f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0935f905b828210613379575050936132e59998969361324d989693600193836131e69810613361575b505050811b019055614818565b01515f1960f88460031b161c191690555f8080613354565b8060018697829497870151815501960194019061332f565b5f805160206150be8339815191525480156133ff57600954906133cf6133b5612f6c565b61187a6001600160401b03918260085460401c1690612c2e565b90670de0b6b3a764000091828102928184041490151715610ff6576133f3916128b7565b8101809111610ff65790565b5060095490565b1561340d57565b60405162461bcd60e51b815260206004820152600d60248201526c3d32b93790323ab930ba34b7b760991b6044820152606490fd5b1561344957565b60405162461bcd60e51b815260206004820152601060248201526f1c995dd85c99081d1bdbc81cdb585b1b60821b6044820152606490fd5b6040519061348e8261131e565b5f6020838281520152565b335f9081527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee803602052604090207f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f559060ff905b5416156134f65750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b335f9081527fbe647c0ddb1a23068f378cf99f2f2f9a0cb078f402ce464e1c86ed30b305a70a602052604090207f61b754da92d8d8d7300489a35a466b9ed19cf4a61860a290f89bec3a75de2bcf9060ff906134ec565b335f9081527fab71e3f32666744d246edff3f96e4bdafee2e9867098cdd118a979a7464786a8602052604090205f805160206150fe8339815191529060ff906134ec565b5f8181525f8051602061515e83398151915260209081526040808320338452909152902060ff906134ec565b6135e3613391565b6009556113876135f1612f6c565b67ffffffffffffffff60401b6008549160401b169067ffffffffffffffff60401b191617600855565b613622613391565b6009556136306135f1612f6c565b6001600160a01b0381169081613644575050565b61364d906128f0565b905f52600b60205260405f2055600954600a60205260405f2055565b91906001600160a01b038316801561215d575f805160206150be833981519152908154838101809111610ff6575f805160206150be833981519152556136ae856106c9565b8054840190556040518381525f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602090a354926001600160d01b0384116136fd57611387929350614c25565b604051630e58ae9360e11b8152600481018590526001600160d01b036024820152604490fd5b6040516323b872dd60e01b5f9081526001600160a01b03938416600452938316602452604494909452909160209060648180855af160015f5114811615613795575b836040525f6060521561377757505050565b635274afe760e01b8352166001600160a01b03166004820152602490fd5b60018115166137ab57813b15153d151616613765565b833d5f823e3d90fd5b600d5460ff81166137cf57600191600c5560ff191617600d55565b60405162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481b5a59dc985d195960821b6044820152606490fd5b6001600160a01b03919082161561217557161561215d5760405162461bcd60e51b81526020600482015260156024820152747655423a206e6f6e2d7472616e7366657261626c6560581b6044820152606490fd5b6001600160a01b038082169291831561217557613877816106c9565b54838110613955578361388b9103916106c9565b555f805160206150be8339815191528281540390555f837fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020604051868152a36138d582614bbc565b926138df43614202565b6001600160d01b039485806138f2613ebb565b1691169003948511610ff6576113879461390b91614e85565b50505f9081525f8051602061503e83398151915260205260408120549080527fd4fb29e10204005f1a39963c6862b79a755e22f0177c53f05cdc3786c702f97454821691166144d2565b60405163391434e360e21b81526001600160a01b03929092166004830152602482015260448101839052606490fd5b6001600160a01b0381165f9081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f8051602061515e8339815191529060ff16613a2f575f808052602091825260408082206001600160a01b038516835290925220805460ff1916600117905533906001600160a01b03165f7f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50505f90565b6001600160a01b0381165f9081527f75e09417c5070057df3eafe5054d52fa0b2a87d64a235e197963b615aedee803602052604090207f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55905f8051602061515e8339815191529060ff905b5416613b04575f828152602091825260408082206001600160a01b038616835290925220805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b5050505f90565b6001600160a01b0381165f9081527fbe647c0ddb1a23068f378cf99f2f2f9a0cb078f402ce464e1c86ed30b305a70a602052604090207f61b754da92d8d8d7300489a35a466b9ed19cf4a61860a290f89bec3a75de2bcf905f8051602061515e8339815191529060ff90613aa0565b6001600160a01b0381165f9081527fab71e3f32666744d246edff3f96e4bdafee2e9867098cdd118a979a7464786a8602052604090205f805160206150fe833981519152905f8051602061515e8339815191529060ff90613aa0565b5f8181525f8051602061515e833981519152602081815260408084206001600160a01b038716855290915290912060ff90613aa0565b5f8181525f8051602061515e833981519152602081815260408084206001600160a01b03871685529091529091205460ff1615613b04575f828152602091825260408082206001600160a01b038616835290925220805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b65ffffffffffff613cae43614202565b1680821015613cc1575061059e90614202565b6044925060405191637669fc0f60e11b835260048301526024820152fd5b60405163a9059cbb60e01b5f9081526001600160a01b039384166004526024949094529260209060448180855af160015f5114811615613d28575b836040521561377757505050565b60018115166137ab57813b15153d151616613d1a565b6001600160a01b03811615613ddd575f805160206150fe8339815191525f8190525f8051602061515e8339815191526020527fab71e3f32666744d246edff3f96e4bdafee2e9867098cdd118a979a7464786a95414613da3576132e5906132e061409d565b60405162461bcd60e51b81526020600482015260126024820152711d5c19dc9859195c881a5b9cdd185b1b195960721b6044820152606490fd5b60405162461bcd60e51b815260206004820152600a602482015269075706772616465723d360b41b6044820152606490fd5b6001600160a01b038181165f8181525f8051602061503e8339815191526020526040812080548685166001600160a01b03198216811790925561138796941694613eb59390928691907f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9080a46001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b916144d2565b5f8051602061513e833981519152805480613ed65750505f90565b805f19810111610ff6577f88c46c62109817164d0ae1873830d4299a82e5daf552a3d8e989b27638fcf747915f52015460301c90565b805480613f195750505f90565b5f19918183810111610ff6575f5260205f2001015460301c90565b9190613f3e614606565b613f46614606565b82516001600160401b038111611339575f8051602061507e83398151915290613f7881613f738454612950565b6146ad565b602080601f831160011461400357509080613fac92613fb396975f926132e85750508160011b915f199060031b1c19161790565b9055614911565b613fdb5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10055565b6113875f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10155565b90601f198316966140415f8051602061507e8339815191525f527f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d90565b925f905b89821061408557505090839291600194613fb398991061406d575b505050811b019055614911565b01515f1960f88460031b161c191690555f8080614060565b80600185968294968601518155019501930190614045565b5f805160206150fe833981519152805f525f8051602061515e833981519152602052600160405f20018181549155817fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff5f80a4565b6140fa614d16565b614102614d80565b916040519260208401927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604085015260608401524660808401523060a084015260a0835260c08301918383106001600160401b038411176113395760429360e291846040528151902061190160f01b855260c282015201522090565b9161059e9391614190936149ed565b90929192614a8d565b6001600160a01b03165f8181527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590918190036141e4575050565b60449250604051916301d4b62360e61b835260048301526024820152fd5b65ffffffffffff90818111614215571690565b604490604051906306dfcc6560e41b8252603060048301526024820152fd5b600181111561059e57600181600160801b81101561434d575b6142f56142eb6142e16142d76142cd6142c361430197600488600160401b6142fc9a1015614340575b640100000000811015614333575b62010000811015614326575b61010081101561431a575b601081101561430e575b1015614306575b60030260011c6142bc818b6128b7565b0160011c90565b6142bc818a6128b7565b6142bc81896128b7565b6142bc81886128b7565b6142bc81876128b7565b6142bc81866128b7565b80936128b7565b821190565b900390565b60011b6142ac565b811c9160021b916142a5565b60081c91811b9161429b565b60101c9160081b91614290565b60201c9160101b91614284565b60401c9160201b91614276565b50600160401b9050608082901c61424d565b905b82811061436d57505090565b90918082169080831860011c8201809211610ff6575f8051602061513e8339815191525f5265ffffffffffff80835f8051602061519e833981519152015416908516105f146143bf5750915b90614361565b9291506143cb906128d5565b906143b9565b91905b8382106143e15750505090565b9091928083169080841860011c8201809211610ff657845f5265ffffffffffff808360205f20015416908416105f1461441e5750925b91906143d4565b93925061442a906128d5565b91614417565b90813b156144b1575f8051602061511e83398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2805115614496576132e591614b43565b50503461449f57565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b6001600160a01b038083169392919081169081851415806145fd575b6144fa575b5050505050565b8161456f575b50508261450f575b80806144f3565b7fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a7249161454661454061454c93610701565b91614bbc565b90614bef565b604080516001600160d01b039384168152919092166020820152a25f8080614508565b61457890610701565b61458184614bbc565b61458a43614202565b6001600160d01b0391828061459e86613f0c565b169116900392828411610ff6577fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724936145f3926145da92614f83565b6040805192851683529316602082015291829190820190565b0390a25f80614500565b508315156144ee565b60ff5f8051602061517e8339815191525460401c161561462257565b604051631afcd79f60e31b8152600490fd5b601f8111614640575050565b5f8051602061505e8339815191525f527f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0906020601f840160051c830193106146a3575b601f0160051c01905b818110614698575050565b5f815560010161468d565b9091508190614684565b601f81116146b9575050565b5f8051602061507e8339815191525f527f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d906020601f840160051c8301931061471c575b601f0160051c01905b818110614711575050565b5f8155600101614706565b90915081906146fd565b601f8111614732575050565b5f8051602061509e8339815191525f527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa906020601f840160051c83019310614795575b601f0160051c01905b81811061478a575050565b5f815560010161477f565b9091508190614776565b601f81116147ab575050565b5f805160206150de8339815191525f527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75906020601f840160051c8301931061480e575b601f0160051c01905b818110614803575050565b5f81556001016147f8565b90915081906147ef565b9081516001600160401b038111611339575f8051602061509e8339815191529061484b816148468454612950565b614726565b602080601f83116001146148805750819061487c9394955f926132e85750508160011b915f199060031b1c19161790565b9055565b90601f198316956148be5f8051602061509e8339815191525f527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa90565b925f905b8882106148f9575050836001959697106148e1575b505050811b019055565b01515f1960f88460031b161c191690555f80806148d7565b806001859682949686015181550195019301906148c2565b9081516001600160401b038111611339575f805160206150de833981519152906149448161493f8454612950565b61479f565b602080601f83116001146149755750819061487c9394955f926132e85750508160011b915f199060031b1c19161790565b90601f198316956149b35f805160206150de8339815191525f527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7590565b925f905b8882106149d5575050836001959697106148e157505050811b019055565b806001859682949686015181550195019301906149b7565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411614a64579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15611258575f516001600160a01b03811615614a5a57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115614a7957565b634e487b7160e01b5f52602160045260245ffd5b614a9681614a6f565b80614a9f575050565b614aa881614a6f565b60018103614ac25760405163f645eedf60e01b8152600490fd5b614acb81614a6f565b60028103614aec5760405163fce698f760e01b815260048101839052602490fd5b80614af8600392614a6f565b14614b005750565b6040516335e2f38360e21b81526004810191909152602490fd5b8054821015614b2f575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b905f8091602081519101845af48080614ba9575b15614b6657505061059e614cfd565b15614b8e57604051639996b31560e01b81526001600160a01b03919091166004820152602490fd5b3d15155f03612d775760405163d6bda27560e01b8152600490fd5b503d151580614b575750813b1515614b57565b6001600160d01b0390818111614bd0571690565b604490604051906306dfcc6560e41b825260d060048301526024820152fd5b90614bf943614202565b6001600160d01b03918280614c0d86613f0c565b16911601918211610ff657614c2192614f83565b9091565b90614c2f81614bbc565b91614c3943614202565b6001600160d01b03938480614c4c613ebb565b16911601848111610ff657614c6091614e85565b50506001600160a01b03908116908115614cc0575b5f8051602061503e8339815191526020527fd4fb29e10204005f1a39963c6862b79a755e22f0177c53f05cdc3786c702f974545f9283526040909220546113879450811691166144d2565b614cc983614bbc565b614cd243614202565b908580614cdd613ebb565b1691169003948511610ff65761138794614cf691614e85565b5050614c75565b604051903d82523d5f602084013e60203d830101604052565b614d1e612988565b8051908115614d2e576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100548015614d5b5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b614d88612a5b565b8051908115614d98576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101548015614d5b5790565b5f8051602061513e833981519152908154600160401b8110156113395760018101808455811015614b2f575f92909252805160209091015160301b65ffffffffffff191665ffffffffffff91909116175f8051602061519e8339815191529190910155565b8054600160401b81101561133957614e4791600182018155614b1a565b614e7257815160209092015160301b65ffffffffffff191665ffffffffffff92909216919091179055565b634e487b7160e01b5f525f60045260245ffd5b5f8051602061513e833981519152549192918015614f5a57614ea9614ecb91612889565b5f8051602061513e8339815191525f525f8051602061519e8339815191520190565b9081549165ffffffffffff90818416918316808311614f4857869203614f1057614f0992509065ffffffffffff82549181199060301b169116179055565b60301c9190565b5050614f4390614f2f614f2161137a565b65ffffffffffff9092168252565b6001600160d01b0385166020820152614dc5565b614f09565b604051632520601d60e01b8152600490fd5b50614f7e90614f6a614f2161137a565b6001600160d01b0384166020820152614dc5565b5f9190565b8054929392801561501857614f9a614fa591612889565b825f5260205f200190565b9182549265ffffffffffff91828516928116808411614f4857879303614fe45750614f0992509065ffffffffffff82549181199060301b169116179055565b915050614f4391615004614ff661137a565b65ffffffffffff9093168352565b6001600160d01b0386166020830152614e2a565b5090614f7e91615029614ff661137a565b6001600160d01b0385166020830152614e2a56fee8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d0052c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10252c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0452c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbce8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d0202dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0088c46c62109817164d0ae1873830d4299a82e5daf552a3d8e989b27638fcf748a2646970667358221220417a2fd1e134f97245793ba5b47215090bffc9f74040de2fd754a318c1994ee164736f6c63430008180033",
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

// MAXBOOSTBPS is a free data retrieval call binding the contract method 0x3d9a581a.
//
// Solidity: function MAX_BOOST_BPS() view returns(uint256)
func (_VUB *VUBCaller) MAXBOOSTBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "MAX_BOOST_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOOSTBPS is a free data retrieval call binding the contract method 0x3d9a581a.
//
// Solidity: function MAX_BOOST_BPS() view returns(uint256)
func (_VUB *VUBSession) MAXBOOSTBPS() (*big.Int, error) {
	return _VUB.Contract.MAXBOOSTBPS(&_VUB.CallOpts)
}

// MAXBOOSTBPS is a free data retrieval call binding the contract method 0x3d9a581a.
//
// Solidity: function MAX_BOOST_BPS() view returns(uint256)
func (_VUB *VUBCallerSession) MAXBOOSTBPS() (*big.Int, error) {
	return _VUB.Contract.MAXBOOSTBPS(&_VUB.CallOpts)
}

// MAXCOOLDOWN is a free data retrieval call binding the contract method 0x8b41d35f.
//
// Solidity: function MAX_COOLDOWN() view returns(uint64)
func (_VUB *VUBCaller) MAXCOOLDOWN(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "MAX_COOLDOWN")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXCOOLDOWN is a free data retrieval call binding the contract method 0x8b41d35f.
//
// Solidity: function MAX_COOLDOWN() view returns(uint64)
func (_VUB *VUBSession) MAXCOOLDOWN() (uint64, error) {
	return _VUB.Contract.MAXCOOLDOWN(&_VUB.CallOpts)
}

// MAXCOOLDOWN is a free data retrieval call binding the contract method 0x8b41d35f.
//
// Solidity: function MAX_COOLDOWN() view returns(uint64)
func (_VUB *VUBCallerSession) MAXCOOLDOWN() (uint64, error) {
	return _VUB.Contract.MAXCOOLDOWN(&_VUB.CallOpts)
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

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_VUB *VUBCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_VUB *VUBSession) UPGRADERROLE() ([32]byte, error) {
	return _VUB.Contract.UPGRADERROLE(&_VUB.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_VUB *VUBCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _VUB.Contract.UPGRADERROLE(&_VUB.CallOpts)
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

// PrincipalMigrated is a free data retrieval call binding the contract method 0x3cbc963d.
//
// Solidity: function principalMigrated() view returns(bool)
func (_VUB *VUBCaller) PrincipalMigrated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "principalMigrated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PrincipalMigrated is a free data retrieval call binding the contract method 0x3cbc963d.
//
// Solidity: function principalMigrated() view returns(bool)
func (_VUB *VUBSession) PrincipalMigrated() (bool, error) {
	return _VUB.Contract.PrincipalMigrated(&_VUB.CallOpts)
}

// PrincipalMigrated is a free data retrieval call binding the contract method 0x3cbc963d.
//
// Solidity: function principalMigrated() view returns(bool)
func (_VUB *VUBCallerSession) PrincipalMigrated() (bool, error) {
	return _VUB.Contract.PrincipalMigrated(&_VUB.CallOpts)
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

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_VUB *VUBCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VUB.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_VUB *VUBSession) TotalStaked() (*big.Int, error) {
	return _VUB.Contract.TotalStaked(&_VUB.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_VUB *VUBCallerSession) TotalStaked() (*big.Int, error) {
	return _VUB.Contract.TotalStaked(&_VUB.CallOpts)
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
// Solidity: function extendLock(uint64 newDur) returns(uint256 minted)
func (_VUB *VUBTransactor) ExtendLock(opts *bind.TransactOpts, newDur uint64) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "extendLock", newDur)
}

// ExtendLock is a paid mutator transaction binding the contract method 0x859e6d6a.
//
// Solidity: function extendLock(uint64 newDur) returns(uint256 minted)
func (_VUB *VUBSession) ExtendLock(newDur uint64) (*types.Transaction, error) {
	return _VUB.Contract.ExtendLock(&_VUB.TransactOpts, newDur)
}

// ExtendLock is a paid mutator transaction binding the contract method 0x859e6d6a.
//
// Solidity: function extendLock(uint64 newDur) returns(uint256 minted)
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

// MigratePrincipalAccounting is a paid mutator transaction binding the contract method 0x182f2719.
//
// Solidity: function migratePrincipalAccounting(uint256 _totalStaked) returns()
func (_VUB *VUBTransactor) MigratePrincipalAccounting(opts *bind.TransactOpts, _totalStaked *big.Int) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "migratePrincipalAccounting", _totalStaked)
}

// MigratePrincipalAccounting is a paid mutator transaction binding the contract method 0x182f2719.
//
// Solidity: function migratePrincipalAccounting(uint256 _totalStaked) returns()
func (_VUB *VUBSession) MigratePrincipalAccounting(_totalStaked *big.Int) (*types.Transaction, error) {
	return _VUB.Contract.MigratePrincipalAccounting(&_VUB.TransactOpts, _totalStaked)
}

// MigratePrincipalAccounting is a paid mutator transaction binding the contract method 0x182f2719.
//
// Solidity: function migratePrincipalAccounting(uint256 _totalStaked) returns()
func (_VUB *VUBTransactorSession) MigratePrincipalAccounting(_totalStaked *big.Int) (*types.Transaction, error) {
	return _VUB.Contract.MigratePrincipalAccounting(&_VUB.TransactOpts, _totalStaked)
}

// MigratePrincipalAndUpgrader is a paid mutator transaction binding the contract method 0x584fb910.
//
// Solidity: function migratePrincipalAndUpgrader(uint256 _totalStaked, address upgrader) returns()
func (_VUB *VUBTransactor) MigratePrincipalAndUpgrader(opts *bind.TransactOpts, _totalStaked *big.Int, upgrader common.Address) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "migratePrincipalAndUpgrader", _totalStaked, upgrader)
}

// MigratePrincipalAndUpgrader is a paid mutator transaction binding the contract method 0x584fb910.
//
// Solidity: function migratePrincipalAndUpgrader(uint256 _totalStaked, address upgrader) returns()
func (_VUB *VUBSession) MigratePrincipalAndUpgrader(_totalStaked *big.Int, upgrader common.Address) (*types.Transaction, error) {
	return _VUB.Contract.MigratePrincipalAndUpgrader(&_VUB.TransactOpts, _totalStaked, upgrader)
}

// MigratePrincipalAndUpgrader is a paid mutator transaction binding the contract method 0x584fb910.
//
// Solidity: function migratePrincipalAndUpgrader(uint256 _totalStaked, address upgrader) returns()
func (_VUB *VUBTransactorSession) MigratePrincipalAndUpgrader(_totalStaked *big.Int, upgrader common.Address) (*types.Transaction, error) {
	return _VUB.Contract.MigratePrincipalAndUpgrader(&_VUB.TransactOpts, _totalStaked, upgrader)
}

// MigrateUpgrader is a paid mutator transaction binding the contract method 0xb49c7165.
//
// Solidity: function migrateUpgrader(address upgrader) returns()
func (_VUB *VUBTransactor) MigrateUpgrader(opts *bind.TransactOpts, upgrader common.Address) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "migrateUpgrader", upgrader)
}

// MigrateUpgrader is a paid mutator transaction binding the contract method 0xb49c7165.
//
// Solidity: function migrateUpgrader(address upgrader) returns()
func (_VUB *VUBSession) MigrateUpgrader(upgrader common.Address) (*types.Transaction, error) {
	return _VUB.Contract.MigrateUpgrader(&_VUB.TransactOpts, upgrader)
}

// MigrateUpgrader is a paid mutator transaction binding the contract method 0xb49c7165.
//
// Solidity: function migrateUpgrader(address upgrader) returns()
func (_VUB *VUBTransactorSession) MigrateUpgrader(upgrader common.Address) (*types.Transaction, error) {
	return _VUB.Contract.MigrateUpgrader(&_VUB.TransactOpts, upgrader)
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

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_VUB *VUBTransactor) StopReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VUB.contract.Transact(opts, "stopReward")
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_VUB *VUBSession) StopReward() (*types.Transaction, error) {
	return _VUB.Contract.StopReward(&_VUB.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_VUB *VUBTransactorSession) StopReward() (*types.Transaction, error) {
	return _VUB.Contract.StopReward(&_VUB.TransactOpts)
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

// VUBRewardStoppedIterator is returned from FilterRewardStopped and is used to iterate over the raw logs and unpacked data for RewardStopped events raised by the VUB contract.
type VUBRewardStoppedIterator struct {
	Event *VUBRewardStopped // Event containing the contract specifics and raw log

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
func (it *VUBRewardStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VUBRewardStopped)
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
		it.Event = new(VUBRewardStopped)
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
func (it *VUBRewardStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VUBRewardStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VUBRewardStopped represents a RewardStopped event raised by the VUB contract.
type VUBRewardStopped struct {
	Unstreamed *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardStopped is a free log retrieval operation binding the contract event 0x278b7aeb8d452471692dbcef95ab699085487ea0c4f3e5e2d13c84795a0f1ae2.
//
// Solidity: event RewardStopped(uint256 unstreamed)
func (_VUB *VUBFilterer) FilterRewardStopped(opts *bind.FilterOpts) (*VUBRewardStoppedIterator, error) {

	logs, sub, err := _VUB.contract.FilterLogs(opts, "RewardStopped")
	if err != nil {
		return nil, err
	}
	return &VUBRewardStoppedIterator{contract: _VUB.contract, event: "RewardStopped", logs: logs, sub: sub}, nil
}

// WatchRewardStopped is a free log subscription operation binding the contract event 0x278b7aeb8d452471692dbcef95ab699085487ea0c4f3e5e2d13c84795a0f1ae2.
//
// Solidity: event RewardStopped(uint256 unstreamed)
func (_VUB *VUBFilterer) WatchRewardStopped(opts *bind.WatchOpts, sink chan<- *VUBRewardStopped) (event.Subscription, error) {

	logs, sub, err := _VUB.contract.WatchLogs(opts, "RewardStopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VUBRewardStopped)
				if err := _VUB.contract.UnpackLog(event, "RewardStopped", log); err != nil {
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

// ParseRewardStopped is a log parse operation binding the contract event 0x278b7aeb8d452471692dbcef95ab699085487ea0c4f3e5e2d13c84795a0f1ae2.
//
// Solidity: event RewardStopped(uint256 unstreamed)
func (_VUB *VUBFilterer) ParseRewardStopped(log types.Log) (*VUBRewardStopped, error) {
	event := new(VUBRewardStopped)
	if err := _VUB.contract.UnpackLog(event, "RewardStopped", log); err != nil {
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
