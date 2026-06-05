// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governor

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

// DAOGovernorMetaData contains all meta data concerning the DAOGovernor contract.
var DAOGovernorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BALLOT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"COUNTING_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVote\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteBySig\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteWithReason\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteWithReasonAndParams\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteWithReasonAndParamsBySig\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getProposalId\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotesWithParams\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashProposal\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIVotes\"},{\"name\":\"timelock\",\"type\":\"address\",\"internalType\":\"contractTimelockControllerUpgradeable\"},{\"name\":\"votingDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"votingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_proposalThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumFraction\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proposalDeadline\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalEta\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalNeedsQueuing\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalProposer\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalSnapshot\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalVotes\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"againstVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"forVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"abstainVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"propose\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queue\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"quorum\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumDenominator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumNumerator\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumNumerator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relay\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setProposalThreshold\",\"inputs\":[{\"name\":\"newProposalThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVotingDelay\",\"inputs\":[{\"name\":\"newVotingDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVotingPeriod\",\"inputs\":[{\"name\":\"newVotingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"state\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIGovernor.ProposalState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timelock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC5805\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateQuorumNumerator\",\"inputs\":[{\"name\":\"newQuorumNumerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTimelock\",\"inputs\":[{\"name\":\"newTimelock\",\"type\":\"address\",\"internalType\":\"contractTimelockControllerUpgradeable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"votingDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"votingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalCanceled\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalCreated\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"targets\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"signatures\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"},{\"name\":\"voteStart\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"voteEnd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalExecuted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalQueued\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"etaSeconds\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalThresholdSet\",\"inputs\":[{\"name\":\"oldProposalThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newProposalThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumNumeratorUpdated\",\"inputs\":[{\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newQuorumNumerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimelockChange\",\"inputs\":[{\"name\":\"oldTimelock\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTimelock\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCast\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCastWithParams\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VotingDelaySet\",\"inputs\":[{\"name\":\"oldVotingDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotingDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VotingPeriodSet\",\"inputs\":[{\"name\":\"oldVotingPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotingPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GovernorAlreadyCastVote\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GovernorAlreadyQueuedProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorDisabledDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GovernorInsufficientProposerVotes\",\"inputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorInvalidProposalLength\",\"inputs\":[{\"name\":\"targets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"calldatas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"values\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorInvalidQuorumFraction\",\"inputs\":[{\"name\":\"quorumNumerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumDenominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorInvalidSignature\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GovernorInvalidVoteParams\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GovernorInvalidVoteType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GovernorInvalidVotingPeriod\",\"inputs\":[{\"name\":\"votingPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorNonexistentProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorNotQueuedProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorOnlyExecutor\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GovernorQueueNotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GovernorRestrictedProposer\",\"inputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GovernorUnableToCancel\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GovernorUnexpectedProposalState\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"uint8\",\"internalType\":\"enumIGovernor.ProposalState\"},{\"name\":\"expectedStates\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60808060405234610016576147ac908161001b8239f35b5f80fdfe60806040526004361015610022575b3615610018575f80fd5b6100206127f2565b005b5f3560e01c806301ffc9a71461033c57806302a251a31461033757806306f3f9e61461033257806306fdde031461032d578063143489d014610328578063150b7a0214610323578063160cbed71461031e57806322f120de146103195780632656227d146103145780632d63f6931461030f5780632fe3e2611461030a5780633932abb1146103055780633e4f49e61461030057806343859632146102fb578063452115d6146102f65780634bf5d7e9146102f1578063544ffc9c146102ec57806354fd4d50146102e757806356781388146102e25780635b8d0e0d146102dd5780635f398a14146102d857806360c4247f146102d357806379051887146102ce5780637b3c71d3146102c95780637d5e81e2146102c45780637ecebe00146102bf57806384b0196e146102ba5780638ff262e3146102b557806391ddadf4146102b057806397c3d334146102ab5780639a802a6d146102a6578063a7713a70146102a1578063a890c9101461029c578063a8f8a66814610279578063a9a9529414610297578063ab58fb8e14610292578063b58131b01461028d578063bc197c8114610288578063c01f9e3714610283578063c28bc2fa1461027e578063c59057e414610279578063d33219b414610274578063dd4e2ba51461026f578063deaaa7cc1461026a578063e540d01d14610265578063eb9019d414610260578063ece40cc11461025b578063f23a6e6114610256578063f8ce560a146102515763fc0c546a0361000e57611e1f565b611d88565b611d11565b611ced565b611c5f565b611c2d565b611bf3565b611b96565b611b62565b61195e565b611ad0565b611ab2565b611a12565b6119d7565b611996565b61197a565b6118d9565b6118ae565b6117ee565b6117d3565b6117a9565b611673565b6115a3565b6114b7565b61141e565b6113c9565b61139c565b61137e565b61130f565b611287565b61120c565b6111e1565b61118a565b61115b565b6110b3565b611052565b611025565b610fc4565b610f8a565b610f49565b610dbc565b610c61565b6109eb565b6107a0565b61061c565b61052c565b610402565b6103d0565b346103c25760203660031901126103c25760043563ffffffff60e01b81168091036103c2576020906366defe7760e11b81149081156103b1575b81156103a0575b811561038f575b506040519015158152f35b6301ffc9a760e01b1490505f610384565b630271189760e51b8114915061037d565b6332a2ad4360e11b81149150610376565b5f80fd5b5f9103126103c257565b346103c2575f3660031901126103c257602063ffffffff5f805160206146178339815191525460301c16604051908152f35b346103c25760203660031901126103c25760043561041e612812565b606481116104b3576001600160d01b03908161043861304f565b1661044161260b565b92808311610493577f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b46339979361047791841690614026565b505060408051918252602082019290925290819081015b0390a1005b6040516306dfcc6560e41b815260d0600482015260248101849052604490fd5b6044906040519063243e544560e01b8252600482015260646024820152fd5b5f5b8381106104e35750505f910152565b81810151838201526020016104d4565b9060209161050c815180928185528580860191016104d2565b601f01601f1916010190565b9060206105299281815201906104f3565b90565b346103c2575f3660031901126103c2576040515f5f8051602061473783398151915280549061055a82611e53565b808552916020916001918281169081156105ef5750600114610597575b6105938661058781880382610704565b60405191829182610518565b0390f35b5f90815293507fda13dda7583a39a3cd73e8830529c760837228fa4683752c823b17e10548aad55b8385106105dc57505050508101602001610587826105935f610577565b80548686018401529382019381016105bf565b90508695506105939693506020925061058794915060ff191682840152151560051b82010192935f610577565b346103c25760203660031901126103c2576004355f525f805160206145f7833981519152602052602060018060a01b0360405f205416604051908152f35b6001600160a01b038116036103c257565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b03811161069257604052565b61066b565b604081019081106001600160401b0382111761069257604052565b602081019081106001600160401b0382111761069257604052565b60c081019081106001600160401b0382111761069257604052565b61010081019081106001600160401b0382111761069257604052565b90601f801991011681019081106001600160401b0382111761069257604052565b6040519061073282610697565b565b6001600160401b03811161069257601f01601f191660200190565b92919261075b82610734565b916107696040519384610704565b8294818452818301116103c2578281602093845f960137010152565b9080601f830112156103c2578160206105299335910161074f565b346103c25760803660031901126103c2576107bc60043561065a565b6107c760243561065a565b6064356001600160401b0381116103c2576107e6903690600401610785565b505f80516020614717833981519152546001600160a01b0316300361081757604051630a85bd0160e11b8152602090f35b604051637485328f60e11b8152600490fd5b6001600160401b0381116106925760051b60200190565b9080601f830112156103c257602090823561085a81610829565b936108686040519586610704565b81855260208086019260051b8201019283116103c257602001905b828210610891575050505090565b838091833561089f8161065a565b815201910190610883565b9080601f830112156103c25760209082356108c481610829565b936108d26040519586610704565b81855260208086019260051b8201019283116103c257602001905b8282106108fb575050505090565b813581529083019083016108ed565b81601f820112156103c25780359160209161092484610829565b936109326040519586610704565b808552838086019160051b830101928084116103c257848301915b84831061095d5750505050505090565b82356001600160401b0381116103c257869161097e84848094890101610785565b81520192019161094d565b60806003198201126103c2576001600160401b03916004358381116103c257826109b591600401610840565b926024358181116103c257836109cd916004016108aa565b926044359182116103c2576109e49160040161090a565b9060643590565b346103c2576109f936610989565b909291610a088285838661278e565b92610a12846128af565b505f805160206147178339815191528054909590610a40906001600160a01b03165b6001600160a01b031690565b9260409687519063793d064960e11b825260209081836004818a5afa968715610be357610aa9988b945f99610c15575b50935163b1c5f42760e01b81523060601b6bffffffffffffffffffffffff19169094189883918591829081908d8b8b8e6004860161322e565b03915afa8015610be357610adb93610a34935f92610be8575b5050610acd8a61200a565b55546001600160a01b031690565b90813b156103c2575f8094610b06878b51998a97889687956308f2a0bb60e41b875260048701613273565b03925af1908115610be357610b2a92610b2592610bca575b50426132ce565b613b5a565b65ffffffffffff811615610bb957917f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892610ba884610b8a610593966001610b7088612037565b019065ffffffffffff1665ffffffffffff19825416179055565b835185815265ffffffffffff90911660208201529081906040820190565b0390a1519081529081906020820190565b8251634844252360e11b8152600490fd5b80610bd7610bdd9261067f565b806103c6565b5f610b1e565b61227c565b610c079250803d10610c0e575b610bff8183610704565b810190612fc7565b5f80610ac2565b503d610bf5565b87919950918689610c338795863d8811610c0e57610bff8183610704565b9b9350505091610a70565b65ffffffffffff8116036103c257565b6064359063ffffffff821682036103c257565b346103c25760c03660031901126103c257600435610c7e8161065a565b602435610c8a8161065a565b60443590610c9782610c3e565b610c9f610c4e565b925f8051602061475783398151915254936001600160401b0360ff8660401c1615951680159081610db4575b6001149081610daa575b159081610da1575b50610d8f575f80516020614757833981519152805467ffffffffffffffff19166001179055610d199385610d6b575b60a4359360843593612063565b610d1f57005b5f80516020614757833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290806020810161048e565b5f80516020614757833981519152805460ff60401b1916600160401b179055610d0c565b60405163f92ee8a960e01b8152600490fd5b9050155f610cdd565b303b159150610cd5565b869150610ccb565b610dc536610989565b610dd48183858795969761278e565b92610dde84612901565b50610dfe610deb85612037565b805460ff60f01b1916600160f01b179055565b5f8051602061471783398151915280546001600160a01b039691949087163003610edc575b926105939692610e3792610e48958861365d565b91543092166001600160a01b031690565b141580610eab575b610e91575b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f90602090a16040519081529081906020820190565b610ea65f5f805160206146f783398151915255565b610e55565b50610ed7610ed35f805160206146f7833981519152546001600160801b0381169060801c1490565b1590565b610e50565b9391959094925f5b8551811015610f395760019030610f0e610a34610f01848b612263565b516001600160a01b031690565b14610f1a575b01610ee4565b610f34610f27828b612263565b5160208151910120612ac8565b610f14565b5090959294919390929190610e23565b346103c25760203660031901126103c2576004355f525f805160206145f7833981519152602052602065ffffffffffff60405f205460a01c16604051908152f35b346103c2575f3660031901126103c25760206040517f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118152f35b346103c2575f3660031901126103c257602065ffffffffffff5f805160206146178339815191525416604051908152f35b634e487b7160e01b5f52602160045260245ffd5b6008111561101357565b610ff5565b9060088210156110135752565b346103c25760203660031901126103c2576020611043600435612b63565b6110506040518092611018565bf35b346103c25760403660031901126103c257602060ff6110a76024356110768161065a565b6004355f525f805160206146d78339815191528452600360405f20019060018060a01b03165f5260205260405f2090565b54166040519015158152f35b346103c2576110c136610989565b926110d18483858496959661278e565b6110da81612b63565b6008811015611013571580611132575b1561110e576105936110fe868686866137e4565b6040519081529081906020820190565b604051638fe5d8a960e01b81526004810191909152336024820152604490fd5b0390fd5b50805f525f805160206145f783398151915260205260018060a01b0360405f20541633146110ea565b346103c2575f3660031901126103c2576105936111766122c0565b6040519182916020835260208301906104f3565b346103c25760203660031901126103c2576004355f525f805160206146d783398151915260205260405f20805461059360026001840154930154604051938493846040919493926060820195825260208201520152565b346103c2575f3660031901126103c257610593611176612381565b6024359060ff821682036103c257565b346103c25760403660031901126103c257602061125261122a6111fc565b604051611236816106b2565b5f815260405191611246836106b2565b5f835233600435612cef565b604051908152f35b9181601f840112156103c2578235916001600160401b0383116103c257602083818601950101116103c257565b346103c25760c03660031901126103c2576112a06111fc565b604435906112ad8261065a565b6001600160401b03906064358281116103c2576112ce90369060040161125a565b6084358481116103c2576112e6903690600401610785565b9160a4359485116103c257610593956113066110fe963690600401610785565b9460043561239e565b346103c25760803660031901126103c2576113286111fc565b6001600160401b03906044358281116103c25761134990369060040161125a565b90916064359384116103c2576113746112529361136c6020963690600401610785565b93369161074f565b9033600435612cef565b346103c25760203660031901126103c2576020611252600435612e07565b346103c25760203660031901126103c2576100206004356113bc81610c3e565b6113c4612812565b612f05565b346103c25760603660031901126103c2576113e26111fc565b6044356001600160401b0381116103c25760209161141161140a61125293369060040161125a565b369161074f565b60405191611246836106b2565b346103c25760803660031901126103c2576001600160401b036004358181116103c25761144f903690600401610840565b906024358181116103c2576114689036906004016108aa565b916044358281116103c25761148190369060040161090a565b6064359283116103c257366023840112156103c257610593936114b16110fe94369060248160040135910161074f565b926124c3565b346103c25760203660031901126103c2576004356114d48161065a565b60018060a01b03165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052602060405f2054604051908152f35b9081518082526020808093019301915f5b828110611530575050505090565b835185529381019392810192600101611522565b916115799061156b61052997959693600f60f81b865260e0602087015260e08601906104f3565b9084820360408601526104f3565b60608301949094526001600160a01b031660808201525f60a082015280830360c090910152611511565b346103c2575f3660031901126103c2577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10054158061164a575b1561160d576115e9611e8b565b6115f1611f5e565b906105936115fd6125f5565b6040519384933091469186611544565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10154156115dc565b346103c25760803660031901126103c25760043561168f6111fc565b906044359161169d8361065a565b6064356001600160401b0381116103c257610ed36116c261176d923690600401610785565b6001600160a01b0386165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052604090208054600181019091556117679060405160208101917ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d7835288604083015260ff8816606083015260018060a01b038a16608083015260a082015260a0815261175f816106cd565b519020613916565b866139a0565b61178857906110fe9161059393611782612051565b92612c74565b6040516394ab6c0760e01b81526001600160a01b0384166004820152602490fd5b346103c2575f3660031901126103c25760206117c361260b565b65ffffffffffff60405191168152f35b346103c2575f3660031901126103c257602060405160648152f35b346103c25760603660031901126103c25760043561180b8161065a565b6044356001600160401b0381116103c25761182a903690600401610785565b505f805160206146b783398151915254604051630748d63560e31b81526001600160a01b03928316600482015260248035908201529160209183916044918391165afa8015610be357610593915f9161188f575b506040519081529081906020820190565b6118a8915060203d602011610c0e57610bff8183610704565b5f61187e565b346103c2575f3660031901126103c25760206001600160d01b036118d061304f565b16604051908152f35b346103c25760203660031901126103c2576004356118f68161065a565b6118fe612812565b5f805160206147178339815191528054604080516001600160a01b0380841682529094166020850181905292937f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b2264019190a16001600160a01b031916179055005b346103c257602061125261197136610989565b9291909161278e565b346103c25760203660031901126103c257602060405160018152f35b346103c25760203660031901126103c2576004355f525f805160206145f7833981519152602052602065ffffffffffff600160405f20015416604051908152f35b346103c2575f3660031901126103c25760207ed7616c8fe29c6c2fbe1d0c5bc8f2faa4c35b43746e70b24b4d532752affd0054604051908152f35b346103c25760a03660031901126103c257611a2e60043561065a565b611a3960243561065a565b6001600160401b036044358181116103c257611a599036906004016108aa565b506064358181116103c257611a729036906004016108aa565b506084359081116103c257611a8b903690600401610785565b50610593611a97612691565b6040516001600160e01b031990911681529081906020820190565b346103c25760203660031901126103c25760206112526004356126ba565b60603660031901126103c257600435611ae88161065a565b604435906001600160401b0382116103c2575f8091611b0e61002094369060040161125a565b90611b17612812565b81604051928392833781018481520391602435905af13d15611b5a573d90611b3e82610734565b91611b4c6040519384610704565b82523d5f602084013e6130a0565b6060906130a0565b346103c2575f3660031901126103c2575f80516020614717833981519152546040516001600160a01b039091168152602090f35b346103c2575f3660031901126103c257610593604051611bb581610697565b602081527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e60208201526040519182916020835260208301906104f3565b346103c2575f3660031901126103c25760206040517ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d78152f35b346103c25760203660031901126103c25760043563ffffffff811681036103c25761002090611c5a612812565b6130c9565b346103c25760403660031901126103c257600435611c7c8161065a565b5f604051611c89816106b2565b525f805160206146b783398151915254604051630748d63560e31b81526001600160a01b03928316600482015260248035908201529160209183916044918391165afa8015610be357610593915f9161188f57506040519081529081906020820190565b346103c25760203660031901126103c257611d06612812565b610020600435613159565b346103c25760a03660031901126103c257611d2d60043561065a565b611d3860243561065a565b6084356001600160401b0381116103c257611d57903690600401610785565b505f80516020614717833981519152546001600160a01b031630036108175760405163f23a6e6160e01b8152602090f35b346103c25760203660031901126103c2575f805160206146b783398151915254604051632394e7a360e21b81526004803590820181905291602090829060249082906001600160a01b03165afa8015610be357610593926110fe925f92611dfa575b50611df490612e07565b90613f25565b611df4919250611e189060203d602011610c0e57610bff8183610704565b9190611dea565b346103c2575f3660031901126103c2575f805160206146b7833981519152546040516001600160a01b039091168152602090f35b90600182811c92168015611e81575b6020831014611e6d57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691611e62565b604051905f825f8051602061467783398151915291825492611eac84611e53565b80845293602091600191828116908115611f385750600114611ed8575b50505061073292500383610704565b5f9081527f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d9590935091905b828410611f2057506107329450505081016020015f8080611ec9565b85548885018301529485019487945092810192611f04565b925050506020925061073294915060ff191682840152151560051b8201015f8080611ec9565b604051905f825f8051602061469783398151915291825492611f7f84611e53565b80845293602091600191828116908115611f385750600114611faa5750505061073292500383610704565b5f9081527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b759590935091905b828410611ff257506107329450505081016020015f8080611ec9565b85548885018301529485019487945092810192611fd6565b5f527f0d5829787b8befdbc6044ef7457d8a95c2a04bc99235349f1a212c063e59d40160205260405f2090565b5f525f805160206145f783398151915260205260405f2090565b6040519061205e826106b2565b5f8252565b929094939160405161207481610697565b600b815260206a2220a7a3b7bb32b93737b960a91b60208301526120966132db565b61209e612381565b906120a76132db565b82516001600160401b038111610692575f80516020614677833981519152916120d9826120d48554611e53565b613309565b602090601f83116001146121a657506121849461212a6107329c9b99956121919995612123866121969e9b9761217f975f9161219b575b508160011b915f199060031b1c19161790565b9055613474565b6121525f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10055565b61217a5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10155565b613579565b612953565b61218c6132db565b612974565b6129b1565b612a58565b90508601515f612110565b5f805160206146778339815191525f5290601f1983167f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d925f905b8282106122375750506107329c9b9995612191999560018661217f966121969f9c98966121849c9761212a971061221f575b5050811b019055613474565b8701515f1960f88460031b161c191690555f80612213565b80600185968294968c015181550195019301906121e1565b634e487b7160e01b5f52603260045260245ffd5b80518210156122775760209160051b010190565b61224f565b6040513d5f823e3d90fd5b6040519061229482610697565b601d82527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c740000006020830152565b5f805160206146b783398151915254604051634bf5d7e960e01b8152905f90829060049082906001600160a01b03165afa5f9181612307575b506105295750610529612287565b9091503d805f833e6123198183610704565b8101906020818303126103c2578051906001600160401b0382116103c2570181601f820112156103c257805161234e81610734565b9261235c6040519485610704565b818452602082840101116103c25761237a91602080850191016104d2565b905f6122f9565b6040519061238e82610697565b60018252603160f81b6020830152565b939092919695610ed36124739161246d8a6123ba36878961074f565b6001600160a01b0382165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590602081519101208b5160208d0120906040519260208401947f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81186528d604086015260ff8d16606086015260018060a01b0316608085015260a084015260c083015260e082015260e0815261175f816106e8565b8a6139a0565b61248e5761052995969161248891369161074f565b92612cef565b6040516394ab6c0760e01b81526001600160a01b0388166004820152602490fd5b634e487b7160e01b5f52601160045260245ffd5b91939290936124d28233612f5f565b156125dd577ed7616c8fe29c6c2fbe1d0c5bc8f2faa4c35b43746e70b24b4d532752affd0054948561250c575b6105299495503393613d5f565b65ffffffffffff5f198161251e61260b565b1601908082116125d8576020905f604051612538816106b2565b525f805160206146b783398151915254604051630748d63560e31b8152336004820152939091166024840152829060449082906001600160a01b03165afa908115610be3575f916125b9575b5086811061259257506124ff565b604051636121770b60e11b8152336004820152602481019190915260448101879052606490fd5b6125d2915060203d602011610c0e57610bff8183610704565b5f612584565b6124af565b60405163d9b3955760e01b8152336004820152602490fd5b604051612601816106b2565b5f8152905f368137565b5f805160206146b7833981519152546040516324776b7d60e21b815290602090829060049082906001600160a01b03165afa5f9181612654575b50610529575061052943613b5a565b9091506020813d602011612689575b8161267060209383610704565b810103126103c2575161268281610c3e565b905f612645565b3d9150612663565b5f80516020614717833981519152546001600160a01b031630036108175763bc197c8160e01b90565b5f525f805160206145f783398151915260205260405f205465ffffffffffff908163ffffffff8260d01c169160a01c16018181116125d8571690565b9081518082526020808093019301915f5b828110612715575050505090565b83516001600160a01b031685529381019392810192600101612707565b90808251908181526020809101926020808460051b8301019501935f915b8483106127605750505050505090565b909192939495848061277e600193601f198682030187528a516104f3565b9801930193019194939290612750565b92906127da926127ec926040519485926127ca6127b7602086019960808b5260a08701906126f6565b601f199687878303016040880152611511565b9085858303016060860152612732565b90608083015203908101835282610704565b51902090565b5f80516020614717833981519152546001600160a01b0316300361081757565b5f80516020614717833981519152546001600160a01b031633810361287857300361283957565b61284236610734565b61284f6040519182610704565b3681526020810190365f83375f602036830101525190205b806128706131af565b036128675750565b6040516347096e4760e01b8152336004820152602490fd5b6040906128ab5f939594606083019683526020830190611018565b0152565b6128b881612b63565b906008821015611013576010600160ff84161b16156128d5575090565b6128f8606492604051926331b75e4d60e01b845260048401526024830190611018565b60106044820152fd5b61290a81612b63565b906008821015611013576030600160ff84161b1615612927575090565b61294a606492604051926331b75e4d60e01b845260048401526024830190611018565b60306044820152fd5b6107329291611c5a61296f926129676132db565b6113c46132db565b613159565b61297c6132db565b6129846132db565b5f805160206146b783398151915280546001600160a01b0319166001600160a01b03909216919091179055565b6129b96132db565b6129c16132db565b606481116104b3576001600160d01b03806129da61304f565b16906129e461260b565b90808411612a3857839291612a1c917f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997951690614026565b505060408051918252602082019290925290819081015b0390a1565b6040516306dfcc6560e41b815260d0600482015260248101859052604490fd5b612a606132db565b612a686132db565b5f805160206147178339815191528054604080516001600160a01b0380841682529094166020850181905292937f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b2264019190a16001600160a01b031916179055565b5f805160206146f78339815191529081548060801c9160018301926001600160801b0380931683851614612b39575f527f7c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb0360205260405f20558254916001600160801b03199060801b169116179055565b634e487b715f5260416020526024601cfd5b908160209103126103c2575180151581036103c25790565b612b6c816136e7565b90612b7682611009565b60058203612c7057612b88915061200a565b545f8051602061471783398151915254612baa906001600160a01b0316610a34565b604051632c258a9f60e11b81526004810183905260209291908381602481855afa908115610be3575f91612c53575b5015612be757505050600590565b604051632ab0f52960e01b815260048101929092528290829060249082905afa918215610be3575f92612c26575b505015612c2157600790565b600290565b612c459250803d10612c4c575b612c3d8183610704565b810190612b4b565b5f80612c15565b503d612c33565b612c6a9150843d8611612c4c57612c3d8183610704565b5f612bd9565b5090565b91610529939160405193612c87856106b2565b5f8552612cef565b93909260ff612cbb9361052997958752166020860152604085015260a0606085015260a08401906104f3565b9160808184039101526104f3565b909260ff60809361052996958452166020830152604082015281606082015201906104f3565b929190612cfb84612b63565b6008811015611013576002600160ff83161b1615612dd95750612d51612d49612d43612d38612d2988612037565b5460a01c65ffffffffffff1690565b65ffffffffffff1690565b83612fd6565b838387613a07565b948051155f14612d9d5750612d977fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4938660405194859460018060a01b03169785612cc9565b0390a290565b612d97907fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712948760405195869560018060a01b03169886612c8f565b6040516331b75e4d60e01b815260048101869052606491612dfe906024830190611018565b60026044820152fd5b612e0f613af8565b9265ffffffffffff92509082168110612e315750505b6001600160d01b031690565b612e3c919250613b5a565b905f80516020614637833981519152918254905f91809360058211612eb1575b5050612e689350614307565b80612e7457505f612e25565b612ea5612e83612eac92613041565b5f805160206146378339815191525f525f805160206146578339815191520190565b5460301c90565b612e25565b612ebd829395926141dc565b83039283116125d857612e68955f5280835f80516020614657833981519152015416908516105f14612ef35750915b5f80612e5c565b929150612eff906132c0565b90612eec565b5f805160206146178339815191529081547fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93604065ffffffffffff81519481851686521693846020820152a165ffffffffffff1916179055565b90805160348110612fbf5760131981830101516001600160b01b03191669dc8f8d908f908c9a8dc360b01b01612fbf57612f9e91602919820190613bb8565b9015918215612fac57505090565b6001600160a01b03918216911614919050565b505050600190565b908160209103126103c2575190565b5f805160206146b783398151915254604051630748d63560e31b81526001600160a01b039283166004820152602481019390935260209183916044918391165afa908115610be3575f91613028575090565b610529915060203d602011610c0e57610bff8183610704565b5f198101919082116125d857565b5f8051602061463783398151915280548061306a5750505f90565b805f198101116125d8577f293b0181c8ec34cd3252e741689bdc21b70ee7a0ec76216439035a5c3718909a915f52015460301c90565b156130a85790565b8051156130b757602081519101fd5b60405163d6bda27560e01b8152600490fd5b63ffffffff908181169081156131415769ffffffff000000000000907f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e882860405f80516020614617833981519152958654958251918760301c1682526020820152a160301b169069ffffffff0000000000001916179055565b60405163f1cfbf0560e01b81525f6004820152602490fd5b7ed7616c8fe29c6c2fbe1d0c5bc8f2faa4c35b43746e70b24b4d532752affd00805460408051918252602082018490527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc0546191a155565b5f805160206146f7833981519152908154916001600160801b038084169360801c841461321c57835f527f7c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb03602052600160405f20945f8654965501166001600160801b0319825416179055565b634e487b715f5260316020526024601cfd5b94939261325a60809361324c6132689460a08a5260a08a01906126f6565b9088820360208a0152611511565b908682036040880152612732565b935f60608201520152565b91926132a260a0946132946132b0949998979960c0875260c08701906126f6565b908582036020870152611511565b908382036040850152612732565b945f606083015260808201520152565b90600182018092116125d857565b919082018092116125d857565b60ff5f805160206147578339815191525460401c16156132f757565b604051631afcd79f60e31b8152600490fd5b601f8111613315575050565b5f805160206146778339815191525f527f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d906020601f840160051c83019310613378575b601f0160051c01905b81811061336d575050565b5f8155600101613362565b9091508190613359565b601f811161338e575050565b5f805160206147378339815191525f527fda13dda7583a39a3cd73e8830529c760837228fa4683752c823b17e10548aad5906020601f840160051c830193106133f1575b601f0160051c01905b8181106133e6575050565b5f81556001016133db565b90915081906133d2565b601f8111613407575050565b5f805160206146978339815191525f527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75906020601f840160051c8301931061346a575b601f0160051c01905b81811061345f575050565b5f8155600101613454565b909150819061344b565b9081516001600160401b038111610692575f80516020614697833981519152906134a7816134a28454611e53565b6133fb565b602080601f83116001146134e8575081906134d99394955f926134dd575b50508160011b915f199060031b1c19161790565b9055565b015190505f806134c5565b90601f198316956135265f805160206146978339815191525f527f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7590565b925f905b88821061356157505083600195969710613549575b505050811b019055565b01515f1960f88460031b161c191690555f808061353f565b8060018596829496860151815501950193019061352a565b906135826132db565b81516001600160401b038111610692575f80516020614737833981519152906135b4816135af8454611e53565b613382565b602080601f83116001146135e5575081906134d99394955f926134dd5750508160011b915f199060031b1c19161790565b90601f198316956136235f805160206147378339815191525f527fda13dda7583a39a3cd73e8830529c760837228fa4683752c823b17e10548aad590565b925f905b8882106136455750508360019596971061354957505050811b019055565b80600185968294968601518155019501930190613627565b9290939160018060a01b035f80516020614717833981519152541690813b156103c2575f936136b76040519788958694859463e38335e560e01b86526bffffffffffffffffffffffff193060601b1618926004860161322e565b039134905af1908115610be3575f926136d5926136d8575b5061200a565b55565b6136e19061067f565b5f6136cf565b6136f081612037565b5460ff8160f01c166137dd5760f81c6137d757613712612d38612d2983612037565b80156137be57613723612d3861260b565b809110156137b857613734826126ba565b1061373f5750600190565b61374b610ed382614102565b8015613789575b1561375d5750600390565b612d38600161376e61377b93612037565b015465ffffffffffff1690565b61378457600490565b600590565b506137b3610ed3825f525f805160206146d783398151915260205260405f20600181015490541090565b613752565b50505f90565b604051636ad0607560e01b815260048101839052602490fd5b50600290565b5050600790565b906137f093929161278e565b6137f981612b63565b600881101561101357603b600160ff83161b16156138e9575061383361381e82612037565b80546001600160f81b0316600160f81b179055565b6040518181527f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c90602090a16138688161200a565b5480613872575090565b5f8051602061471783398151915254613893906001600160a01b0316610a34565b803b156103c25760405163c4d252f560e01b815260048101929092525f908290602490829084905af18015610be3576138d6575b505f6138d28261200a565b5590565b80610bd76138e39261067f565b5f6138c7565b9061390d606492604051926331b75e4d60e01b845260048401526024830190611018565b603b6044820152fd5b604290613921614547565b6139296145b1565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a08152613978816106cd565b519020906040519161190160f01b8352600283015260228201522090565b6004111561101357565b9091813b6139c8576139b291926141a2565b506139bc81613996565b159182612fac57505090565b6020918160645f935160405192630b135d3f60e11b9788855260048501526040602485015286820190604485015e01915afa905f5114601f3d11161690565b613a28909291925f525f805160206146d783398151915260205260405f2090565b9160038301613a51613a4a83839060018060a01b03165f5260205260405f2090565b5460ff1690565b613ad757613a7560ff9392613a82929060018060a01b03165f5260205260405f2090565b805460ff19166001179055565b1680613a995750613a948282546132ce565b905590565b60018103613ab05750600101613a948282546132ce565b600203613ac557600201613a948282546132ce565b6040516303599be160e11b8152600490fd5b6040516371c6af4960e01b81526001600160a01b0383166004820152602490fd5b5f80516020614637833981519152805480613b175750505f905f905f90565b805f198101116125d8577f293b0181c8ec34cd3252e741689bdc21b70ee7a0ec76216439035a5c3718909a915f52015460019165ffffffffffff82169160301c90565b65ffffffffffff90818111613b6d571690565b604490604051906306dfcc6560e41b8252603060048301526024820152fd5b908160011b91808304600214901517156125d857565b908160041b91808304601014901517156125d857565b9190825182118015613c3f575b613c1857613bd2816132c0565b821180613c21575b613be5901515613b8c565b602801806028116125d8578183038381116125d85703613c1857613c0892614379565b90916001600160a01b0390911690565b5050505f905f90565b50828101602001516001600160f01b03191661060f60f31b14613bda565b50818111613bc5565b90613c5282610829565b613c5f6040519182610704565b8281528092613c70601f1991610829565b01905f5b828110613c8057505050565b806060602080938501015201613c74565b9592613cc590613cd3939b9a989996959261012090895260209c60018060a01b03168d8a01528060408a01528801906126f6565b908682036060880152611511565b9784890360808601528251808a52818a019180808360051b8d01019501925f905b838210613d3157505050505061052996975090613d189184820360a0860152612732565b9360c083015260e08201526101008184039101526104f3565b90919293958380613d508f93600194601f199082030186528a516104f3565b98019201920190939291613cf4565b919493909294613d778651602088012082868661278e565b958351855190818114801590613efc575b8015613ef4575b613ecb57505065ffffffffffff9485613daa612d298a612037565b16613ea4577f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e095612a339363ffffffff613dff613de561260b565b5f80516020614617833981519152549480861691166132ce565b9260301c16613e82613e108c612037565b80546001600160a01b0319166001600160a01b038a16178155613e59613e3586613b5a565b825465ffffffffffff60a01b191660a09190911b65ffffffffffff60a01b16178255565b613e6283614423565b815463ffffffff60d01b191660d09190911b63ffffffff60d01b16179055565b613e96613e8f8951613c48565b91846132ce565b936040519889988d8a613c91565b87613eae81612b63565b6040516331b75e4d60e01b815291829161112e9160048401612890565b8351604051630447b05d60e41b8152600481019290925260248201526044810191909152606490fd5b508015613d8f565b508351811415613d88565b8115613f11570490565b634e487b7160e01b5f52601260045260245ffd5b5f1982820982820291828083109203918083039214613f92578160641115613f80577f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c29936064910990828211900360fe1b910360021c170290565b634e487b715f5260116020526024601cfd5b5050606491500490565b5f80516020614637833981519152908154600160401b811015610692576001810180845581101561227757610732925f525f80516020614657833981519152019061400465ffffffffffff825116839065ffffffffffff1665ffffffffffff19825416179055565b60200151815465ffffffffffff1660309190911b65ffffffffffff1916179055565b5f805160206146378339815191525491929180156140d957612e8361404a91613041565b9081549165ffffffffffff908184169183168083116140c75786920361408f5761408892509065ffffffffffff82549181199060301b169116179055565b60301c9190565b50506140c2906140ae6140a0610725565b65ffffffffffff9092168252565b6001600160d01b0385166020820152613f9c565b614088565b604051632520601d60e01b8152600490fd5b506140fd906140e96140a0610725565b6001600160d01b0384166020820152613f9c565b5f9190565b5f525f805160206146d783398151915260205260405f205f805160206145f783398151915260205265ffffffffffff60405f205460a01c166024602060018060a01b035f805160206146b7833981519152541660405192838092632394e7a360e21b82528660048301525afa8015610be35761419d9261418c925f92611dfa5750611df490612e07565b9160026001820154910154906132ce565b101590565b81519190604183036141d2576141cb9250602082015190606060408401519301515f1a90614453565b9192909190565b50505f9160029190565b600181111561052957600181600160801b8110156142f5575b61429d61429361428961427f61427561426b6142a997600488600160401b6142a49a10156142e8575b6401000000008110156142db575b620100008110156142ce575b6101008110156142c2575b60108110156142b6575b10156142ae575b60030260011c614264818b613f07565b0160011c90565b614264818a613f07565b6142648189613f07565b6142648188613f07565b6142648187613f07565b6142648186613f07565b8093613f07565b821190565b900390565b60011b614254565b811c9160021b9161424d565b60081c91811b91614243565b60101c9160081b91614238565b60201c9160101b9161422c565b60401c9160201b9161421e565b50600160401b9050608082901c6141f5565b905b82811061431557505090565b90918082169080831860011c82018092116125d8575f805160206146378339815191525f5265ffffffffffff80835f80516020614657833981519152015416908516105f146143675750915b90614309565b929150614373906132c0565b90614361565b929092614385846132c0565b831180614405575b61439f6143a691949293941515613b8c565b5f956132ce565b915b8183106143b85750505060019190565b9092919360ff6143d96143d46020888601015160ff60f81b1690565b6144d5565b1690600f82116143fa57906143ef600192613ba2565b0194019192906143a8565b505f94508493505050565b50602084820101516001600160f01b03191661060f60f31b1461438d565b63ffffffff90818111614434571690565b604490604051906306dfcc6560e41b8252602060048301526024820152fd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116144ca579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15610be3575f516001600160a01b038116156144c057905f905f90565b505f906001905f90565b5050505f9160039190565b60f81c602f81118061453d575b156144f157602f190160ff1690565b6060811180614533575b1561450a576056190160ff1690565b6040811180614529575b15614523576036190160ff1690565b5060ff90565b5060478110614514565b50606781106144fb565b50603a81106144e2565b61454f611e8b565b805190811561455f576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10054801561458c5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b6145b9611f5e565b80519081156145c9576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10154801561458c579056fe7c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb0100d7616c8fe29c6c2fbe1d0c5bc8f2faa4c35b43746e70b24b4d532752affd01e770710421fd2cad75ad828c61aa98f2d77d423a440b67872d0f65554148e000293b0181c8ec34cd3252e741689bdc21b70ee7a0ec76216439035a5c3718909ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1033ba4977254e415696610a40ebf2258dbfa0ec6a2ff64e84bfe715ff16977cc00a1cefa0f43667ef127a258e673c94202a79b656e62899531c4376d87a7f398007c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb020d5829787b8befdbc6044ef7457d8a95c2a04bc99235349f1a212c063e59d4007c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220fa85a6d5351176359ec1d86ffa7d4e2ee495c2f324a70f7c22ceba043cb37c2764736f6c63430008180033",
}

// DAOGovernorABI is the input ABI used to generate the binding from.
// Deprecated: Use DAOGovernorMetaData.ABI instead.
var DAOGovernorABI = DAOGovernorMetaData.ABI

// DAOGovernorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DAOGovernorMetaData.Bin instead.
var DAOGovernorBin = DAOGovernorMetaData.Bin

// DeployDAOGovernor deploys a new Ethereum contract, binding an instance of DAOGovernor to it.
func DeployDAOGovernor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DAOGovernor, error) {
	parsed, err := DAOGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DAOGovernorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DAOGovernor{DAOGovernorCaller: DAOGovernorCaller{contract: contract}, DAOGovernorTransactor: DAOGovernorTransactor{contract: contract}, DAOGovernorFilterer: DAOGovernorFilterer{contract: contract}}, nil
}

// DAOGovernor is an auto generated Go binding around an Ethereum contract.
type DAOGovernor struct {
	DAOGovernorCaller     // Read-only binding to the contract
	DAOGovernorTransactor // Write-only binding to the contract
	DAOGovernorFilterer   // Log filterer for contract events
}

// DAOGovernorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAOGovernorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOGovernorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAOGovernorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOGovernorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAOGovernorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOGovernorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAOGovernorSession struct {
	Contract     *DAOGovernor      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAOGovernorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAOGovernorCallerSession struct {
	Contract *DAOGovernorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DAOGovernorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAOGovernorTransactorSession struct {
	Contract     *DAOGovernorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DAOGovernorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAOGovernorRaw struct {
	Contract *DAOGovernor // Generic contract binding to access the raw methods on
}

// DAOGovernorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAOGovernorCallerRaw struct {
	Contract *DAOGovernorCaller // Generic read-only contract binding to access the raw methods on
}

// DAOGovernorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAOGovernorTransactorRaw struct {
	Contract *DAOGovernorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAOGovernor creates a new instance of DAOGovernor, bound to a specific deployed contract.
func NewDAOGovernor(address common.Address, backend bind.ContractBackend) (*DAOGovernor, error) {
	contract, err := bindDAOGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAOGovernor{DAOGovernorCaller: DAOGovernorCaller{contract: contract}, DAOGovernorTransactor: DAOGovernorTransactor{contract: contract}, DAOGovernorFilterer: DAOGovernorFilterer{contract: contract}}, nil
}

// NewDAOGovernorCaller creates a new read-only instance of DAOGovernor, bound to a specific deployed contract.
func NewDAOGovernorCaller(address common.Address, caller bind.ContractCaller) (*DAOGovernorCaller, error) {
	contract, err := bindDAOGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAOGovernorCaller{contract: contract}, nil
}

// NewDAOGovernorTransactor creates a new write-only instance of DAOGovernor, bound to a specific deployed contract.
func NewDAOGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*DAOGovernorTransactor, error) {
	contract, err := bindDAOGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAOGovernorTransactor{contract: contract}, nil
}

// NewDAOGovernorFilterer creates a new log filterer instance of DAOGovernor, bound to a specific deployed contract.
func NewDAOGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*DAOGovernorFilterer, error) {
	contract, err := bindDAOGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAOGovernorFilterer{contract: contract}, nil
}

// bindDAOGovernor binds a generic wrapper to an already deployed contract.
func bindDAOGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DAOGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAOGovernor *DAOGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAOGovernor.Contract.DAOGovernorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAOGovernor *DAOGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOGovernor.Contract.DAOGovernorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAOGovernor *DAOGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAOGovernor.Contract.DAOGovernorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAOGovernor *DAOGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAOGovernor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAOGovernor *DAOGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOGovernor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAOGovernor *DAOGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAOGovernor.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_DAOGovernor *DAOGovernorCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_DAOGovernor *DAOGovernorSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _DAOGovernor.Contract.BALLOTTYPEHASH(&_DAOGovernor.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_DAOGovernor *DAOGovernorCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _DAOGovernor.Contract.BALLOTTYPEHASH(&_DAOGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_DAOGovernor *DAOGovernorCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_DAOGovernor *DAOGovernorSession) CLOCKMODE() (string, error) {
	return _DAOGovernor.Contract.CLOCKMODE(&_DAOGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_DAOGovernor *DAOGovernorCallerSession) CLOCKMODE() (string, error) {
	return _DAOGovernor.Contract.CLOCKMODE(&_DAOGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_DAOGovernor *DAOGovernorCaller) COUNTINGMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "COUNTING_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_DAOGovernor *DAOGovernorSession) COUNTINGMODE() (string, error) {
	return _DAOGovernor.Contract.COUNTINGMODE(&_DAOGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_DAOGovernor *DAOGovernorCallerSession) COUNTINGMODE() (string, error) {
	return _DAOGovernor.Contract.COUNTINGMODE(&_DAOGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_DAOGovernor *DAOGovernorCaller) EXTENDEDBALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "EXTENDED_BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_DAOGovernor *DAOGovernorSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _DAOGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_DAOGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_DAOGovernor *DAOGovernorCallerSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _DAOGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_DAOGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_DAOGovernor *DAOGovernorCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_DAOGovernor *DAOGovernorSession) Clock() (*big.Int, error) {
	return _DAOGovernor.Contract.Clock(&_DAOGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_DAOGovernor *DAOGovernorCallerSession) Clock() (*big.Int, error) {
	return _DAOGovernor.Contract.Clock(&_DAOGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DAOGovernor *DAOGovernorCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "eip712Domain")

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
func (_DAOGovernor *DAOGovernorSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _DAOGovernor.Contract.Eip712Domain(&_DAOGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DAOGovernor *DAOGovernorCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _DAOGovernor.Contract.Eip712Domain(&_DAOGovernor.CallOpts)
}

// GetProposalId is a free data retrieval call binding the contract method 0xa8f8a668.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) GetProposalId(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "getProposalId", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalId is a free data retrieval call binding the contract method 0xa8f8a668.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) GetProposalId(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _DAOGovernor.Contract.GetProposalId(&_DAOGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// GetProposalId is a free data retrieval call binding the contract method 0xa8f8a668.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) GetProposalId(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _DAOGovernor.Contract.GetProposalId(&_DAOGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) GetVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "getVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.GetVotes(&_DAOGovernor.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.GetVotes(&_DAOGovernor.CallOpts, account, timepoint)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) GetVotesWithParams(opts *bind.CallOpts, account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "getVotesWithParams", account, timepoint, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _DAOGovernor.Contract.GetVotesWithParams(&_DAOGovernor.CallOpts, account, timepoint, params)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _DAOGovernor.Contract.GetVotesWithParams(&_DAOGovernor.CallOpts, account, timepoint, params)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_DAOGovernor *DAOGovernorCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_DAOGovernor *DAOGovernorSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _DAOGovernor.Contract.HasVoted(&_DAOGovernor.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_DAOGovernor *DAOGovernorCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _DAOGovernor.Contract.HasVoted(&_DAOGovernor.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) HashProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_DAOGovernor *DAOGovernorSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _DAOGovernor.Contract.HashProposal(&_DAOGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _DAOGovernor.Contract.HashProposal(&_DAOGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DAOGovernor *DAOGovernorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DAOGovernor *DAOGovernorSession) Name() (string, error) {
	return _DAOGovernor.Contract.Name(&_DAOGovernor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DAOGovernor *DAOGovernorCallerSession) Name() (string, error) {
	return _DAOGovernor.Contract.Name(&_DAOGovernor.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) Nonces(owner common.Address) (*big.Int, error) {
	return _DAOGovernor.Contract.Nonces(&_DAOGovernor.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _DAOGovernor.Contract.Nonces(&_DAOGovernor.CallOpts, owner)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) ProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.ProposalDeadline(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.ProposalDeadline(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) ProposalEta(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "proposalEta", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.ProposalEta(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.ProposalEta(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_DAOGovernor *DAOGovernorCaller) ProposalNeedsQueuing(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "proposalNeedsQueuing", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_DAOGovernor *DAOGovernorSession) ProposalNeedsQueuing(proposalId *big.Int) (bool, error) {
	return _DAOGovernor.Contract.ProposalNeedsQueuing(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_DAOGovernor *DAOGovernorCallerSession) ProposalNeedsQueuing(proposalId *big.Int) (bool, error) {
	return _DAOGovernor.Contract.ProposalNeedsQueuing(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_DAOGovernor *DAOGovernorCaller) ProposalProposer(opts *bind.CallOpts, proposalId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "proposalProposer", proposalId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_DAOGovernor *DAOGovernorSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _DAOGovernor.Contract.ProposalProposer(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_DAOGovernor *DAOGovernorCallerSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _DAOGovernor.Contract.ProposalProposer(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) ProposalSnapshot(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.ProposalSnapshot(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.ProposalSnapshot(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) ProposalThreshold() (*big.Int, error) {
	return _DAOGovernor.Contract.ProposalThreshold(&_DAOGovernor.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) ProposalThreshold() (*big.Int, error) {
	return _DAOGovernor.Contract.ProposalThreshold(&_DAOGovernor.CallOpts)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_DAOGovernor *DAOGovernorCaller) ProposalVotes(opts *bind.CallOpts, proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "proposalVotes", proposalId)

	outstruct := new(struct {
		AgainstVotes *big.Int
		ForVotes     *big.Int
		AbstainVotes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AgainstVotes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ForVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AbstainVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_DAOGovernor *DAOGovernorSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _DAOGovernor.Contract.ProposalVotes(&_DAOGovernor.CallOpts, proposalId)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_DAOGovernor *DAOGovernorCallerSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _DAOGovernor.Contract.ProposalVotes(&_DAOGovernor.CallOpts, proposalId)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) Quorum(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "quorum", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.Quorum(&_DAOGovernor.CallOpts, blockNumber)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.Quorum(&_DAOGovernor.CallOpts, blockNumber)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) QuorumDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "quorumDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) QuorumDenominator() (*big.Int, error) {
	return _DAOGovernor.Contract.QuorumDenominator(&_DAOGovernor.CallOpts)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) QuorumDenominator() (*big.Int, error) {
	return _DAOGovernor.Contract.QuorumDenominator(&_DAOGovernor.CallOpts)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) QuorumNumerator(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "quorumNumerator", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.QuorumNumerator(&_DAOGovernor.CallOpts, timepoint)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _DAOGovernor.Contract.QuorumNumerator(&_DAOGovernor.CallOpts, timepoint)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) QuorumNumerator0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "quorumNumerator0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) QuorumNumerator0() (*big.Int, error) {
	return _DAOGovernor.Contract.QuorumNumerator0(&_DAOGovernor.CallOpts)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) QuorumNumerator0() (*big.Int, error) {
	return _DAOGovernor.Contract.QuorumNumerator0(&_DAOGovernor.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_DAOGovernor *DAOGovernorCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_DAOGovernor *DAOGovernorSession) State(proposalId *big.Int) (uint8, error) {
	return _DAOGovernor.Contract.State(&_DAOGovernor.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_DAOGovernor *DAOGovernorCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _DAOGovernor.Contract.State(&_DAOGovernor.CallOpts, proposalId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAOGovernor *DAOGovernorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAOGovernor *DAOGovernorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DAOGovernor.Contract.SupportsInterface(&_DAOGovernor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAOGovernor *DAOGovernorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DAOGovernor.Contract.SupportsInterface(&_DAOGovernor.CallOpts, interfaceId)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_DAOGovernor *DAOGovernorCaller) Timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_DAOGovernor *DAOGovernorSession) Timelock() (common.Address, error) {
	return _DAOGovernor.Contract.Timelock(&_DAOGovernor.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_DAOGovernor *DAOGovernorCallerSession) Timelock() (common.Address, error) {
	return _DAOGovernor.Contract.Timelock(&_DAOGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DAOGovernor *DAOGovernorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DAOGovernor *DAOGovernorSession) Token() (common.Address, error) {
	return _DAOGovernor.Contract.Token(&_DAOGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DAOGovernor *DAOGovernorCallerSession) Token() (common.Address, error) {
	return _DAOGovernor.Contract.Token(&_DAOGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DAOGovernor *DAOGovernorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DAOGovernor *DAOGovernorSession) Version() (string, error) {
	return _DAOGovernor.Contract.Version(&_DAOGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DAOGovernor *DAOGovernorCallerSession) Version() (string, error) {
	return _DAOGovernor.Contract.Version(&_DAOGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) VotingDelay() (*big.Int, error) {
	return _DAOGovernor.Contract.VotingDelay(&_DAOGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) VotingDelay() (*big.Int, error) {
	return _DAOGovernor.Contract.VotingDelay(&_DAOGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_DAOGovernor *DAOGovernorCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOGovernor.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_DAOGovernor *DAOGovernorSession) VotingPeriod() (*big.Int, error) {
	return _DAOGovernor.Contract.VotingPeriod(&_DAOGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_DAOGovernor *DAOGovernorCallerSession) VotingPeriod() (*big.Int, error) {
	return _DAOGovernor.Contract.VotingPeriod(&_DAOGovernor.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactor) Cancel(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "cancel", targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_DAOGovernor *DAOGovernorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Cancel(&_DAOGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Cancel(&_DAOGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_DAOGovernor *DAOGovernorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _DAOGovernor.Contract.CastVote(&_DAOGovernor.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _DAOGovernor.Contract.CastVote(&_DAOGovernor.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "castVoteBySig", proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_DAOGovernor *DAOGovernorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.CastVoteBySig(&_DAOGovernor.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.CastVoteBySig(&_DAOGovernor.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_DAOGovernor *DAOGovernorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _DAOGovernor.Contract.CastVoteWithReason(&_DAOGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _DAOGovernor.Contract.CastVoteWithReason(&_DAOGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactor) CastVoteWithReasonAndParams(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_DAOGovernor *DAOGovernorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.CastVoteWithReasonAndParams(&_DAOGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.CastVoteWithReasonAndParams(&_DAOGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactor) CastVoteWithReasonAndParamsBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_DAOGovernor *DAOGovernorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_DAOGovernor.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_DAOGovernor.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_DAOGovernor *DAOGovernorTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "execute", targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_DAOGovernor *DAOGovernorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Execute(&_DAOGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_DAOGovernor *DAOGovernorTransactorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Execute(&_DAOGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x22f120de.
//
// Solidity: function initialize(address token, address timelock, uint48 votingDelay, uint32 votingPeriod, uint256 _proposalThreshold, uint256 quorumFraction) returns()
func (_DAOGovernor *DAOGovernorTransactor) Initialize(opts *bind.TransactOpts, token common.Address, timelock common.Address, votingDelay *big.Int, votingPeriod uint32, _proposalThreshold *big.Int, quorumFraction *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "initialize", token, timelock, votingDelay, votingPeriod, _proposalThreshold, quorumFraction)
}

// Initialize is a paid mutator transaction binding the contract method 0x22f120de.
//
// Solidity: function initialize(address token, address timelock, uint48 votingDelay, uint32 votingPeriod, uint256 _proposalThreshold, uint256 quorumFraction) returns()
func (_DAOGovernor *DAOGovernorSession) Initialize(token common.Address, timelock common.Address, votingDelay *big.Int, votingPeriod uint32, _proposalThreshold *big.Int, quorumFraction *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Initialize(&_DAOGovernor.TransactOpts, token, timelock, votingDelay, votingPeriod, _proposalThreshold, quorumFraction)
}

// Initialize is a paid mutator transaction binding the contract method 0x22f120de.
//
// Solidity: function initialize(address token, address timelock, uint48 votingDelay, uint32 votingPeriod, uint256 _proposalThreshold, uint256 quorumFraction) returns()
func (_DAOGovernor *DAOGovernorTransactorSession) Initialize(token common.Address, timelock common.Address, votingDelay *big.Int, votingPeriod uint32, _proposalThreshold *big.Int, quorumFraction *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Initialize(&_DAOGovernor.TransactOpts, token, timelock, votingDelay, votingPeriod, _proposalThreshold, quorumFraction)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_DAOGovernor *DAOGovernorTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_DAOGovernor *DAOGovernorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.OnERC1155BatchReceived(&_DAOGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_DAOGovernor *DAOGovernorTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.OnERC1155BatchReceived(&_DAOGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_DAOGovernor *DAOGovernorTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_DAOGovernor *DAOGovernorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.OnERC1155Received(&_DAOGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_DAOGovernor *DAOGovernorTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.OnERC1155Received(&_DAOGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DAOGovernor *DAOGovernorTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DAOGovernor *DAOGovernorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.OnERC721Received(&_DAOGovernor.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DAOGovernor *DAOGovernorTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.OnERC721Received(&_DAOGovernor.TransactOpts, arg0, arg1, arg2, arg3)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "propose", targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_DAOGovernor *DAOGovernorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Propose(&_DAOGovernor.TransactOpts, targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Propose(&_DAOGovernor.TransactOpts, targets, values, calldatas, description)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactor) Queue(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "queue", targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_DAOGovernor *DAOGovernorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Queue(&_DAOGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_DAOGovernor *DAOGovernorTransactorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Queue(&_DAOGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_DAOGovernor *DAOGovernorTransactor) Relay(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "relay", target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_DAOGovernor *DAOGovernorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Relay(&_DAOGovernor.TransactOpts, target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_DAOGovernor *DAOGovernorTransactorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _DAOGovernor.Contract.Relay(&_DAOGovernor.TransactOpts, target, value, data)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_DAOGovernor *DAOGovernorTransactor) SetProposalThreshold(opts *bind.TransactOpts, newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "setProposalThreshold", newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_DAOGovernor *DAOGovernorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.Contract.SetProposalThreshold(&_DAOGovernor.TransactOpts, newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_DAOGovernor *DAOGovernorTransactorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.Contract.SetProposalThreshold(&_DAOGovernor.TransactOpts, newProposalThreshold)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_DAOGovernor *DAOGovernorTransactor) SetVotingDelay(opts *bind.TransactOpts, newVotingDelay *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "setVotingDelay", newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_DAOGovernor *DAOGovernorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.Contract.SetVotingDelay(&_DAOGovernor.TransactOpts, newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_DAOGovernor *DAOGovernorTransactorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.Contract.SetVotingDelay(&_DAOGovernor.TransactOpts, newVotingDelay)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_DAOGovernor *DAOGovernorTransactor) SetVotingPeriod(opts *bind.TransactOpts, newVotingPeriod uint32) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "setVotingPeriod", newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_DAOGovernor *DAOGovernorSession) SetVotingPeriod(newVotingPeriod uint32) (*types.Transaction, error) {
	return _DAOGovernor.Contract.SetVotingPeriod(&_DAOGovernor.TransactOpts, newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_DAOGovernor *DAOGovernorTransactorSession) SetVotingPeriod(newVotingPeriod uint32) (*types.Transaction, error) {
	return _DAOGovernor.Contract.SetVotingPeriod(&_DAOGovernor.TransactOpts, newVotingPeriod)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_DAOGovernor *DAOGovernorTransactor) UpdateQuorumNumerator(opts *bind.TransactOpts, newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "updateQuorumNumerator", newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_DAOGovernor *DAOGovernorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.Contract.UpdateQuorumNumerator(&_DAOGovernor.TransactOpts, newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_DAOGovernor *DAOGovernorTransactorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _DAOGovernor.Contract.UpdateQuorumNumerator(&_DAOGovernor.TransactOpts, newQuorumNumerator)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_DAOGovernor *DAOGovernorTransactor) UpdateTimelock(opts *bind.TransactOpts, newTimelock common.Address) (*types.Transaction, error) {
	return _DAOGovernor.contract.Transact(opts, "updateTimelock", newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_DAOGovernor *DAOGovernorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _DAOGovernor.Contract.UpdateTimelock(&_DAOGovernor.TransactOpts, newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_DAOGovernor *DAOGovernorTransactorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _DAOGovernor.Contract.UpdateTimelock(&_DAOGovernor.TransactOpts, newTimelock)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAOGovernor *DAOGovernorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOGovernor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAOGovernor *DAOGovernorSession) Receive() (*types.Transaction, error) {
	return _DAOGovernor.Contract.Receive(&_DAOGovernor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAOGovernor *DAOGovernorTransactorSession) Receive() (*types.Transaction, error) {
	return _DAOGovernor.Contract.Receive(&_DAOGovernor.TransactOpts)
}

// DAOGovernorEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the DAOGovernor contract.
type DAOGovernorEIP712DomainChangedIterator struct {
	Event *DAOGovernorEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *DAOGovernorEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorEIP712DomainChanged)
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
		it.Event = new(DAOGovernorEIP712DomainChanged)
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
func (it *DAOGovernorEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorEIP712DomainChanged represents a EIP712DomainChanged event raised by the DAOGovernor contract.
type DAOGovernorEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DAOGovernor *DAOGovernorFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*DAOGovernorEIP712DomainChangedIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorEIP712DomainChangedIterator{contract: _DAOGovernor.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DAOGovernor *DAOGovernorFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *DAOGovernorEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorEIP712DomainChanged)
				if err := _DAOGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_DAOGovernor *DAOGovernorFilterer) ParseEIP712DomainChanged(log types.Log) (*DAOGovernorEIP712DomainChanged, error) {
	event := new(DAOGovernorEIP712DomainChanged)
	if err := _DAOGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DAOGovernor contract.
type DAOGovernorInitializedIterator struct {
	Event *DAOGovernorInitialized // Event containing the contract specifics and raw log

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
func (it *DAOGovernorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorInitialized)
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
		it.Event = new(DAOGovernorInitialized)
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
func (it *DAOGovernorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorInitialized represents a Initialized event raised by the DAOGovernor contract.
type DAOGovernorInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DAOGovernor *DAOGovernorFilterer) FilterInitialized(opts *bind.FilterOpts) (*DAOGovernorInitializedIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorInitializedIterator{contract: _DAOGovernor.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DAOGovernor *DAOGovernorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DAOGovernorInitialized) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorInitialized)
				if err := _DAOGovernor.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DAOGovernor *DAOGovernorFilterer) ParseInitialized(log types.Log) (*DAOGovernorInitialized, error) {
	event := new(DAOGovernorInitialized)
	if err := _DAOGovernor.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the DAOGovernor contract.
type DAOGovernorProposalCanceledIterator struct {
	Event *DAOGovernorProposalCanceled // Event containing the contract specifics and raw log

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
func (it *DAOGovernorProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorProposalCanceled)
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
		it.Event = new(DAOGovernorProposalCanceled)
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
func (it *DAOGovernorProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorProposalCanceled represents a ProposalCanceled event raised by the DAOGovernor contract.
type DAOGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_DAOGovernor *DAOGovernorFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*DAOGovernorProposalCanceledIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorProposalCanceledIterator{contract: _DAOGovernor.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_DAOGovernor *DAOGovernorFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *DAOGovernorProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorProposalCanceled)
				if err := _DAOGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_DAOGovernor *DAOGovernorFilterer) ParseProposalCanceled(log types.Log) (*DAOGovernorProposalCanceled, error) {
	event := new(DAOGovernorProposalCanceled)
	if err := _DAOGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the DAOGovernor contract.
type DAOGovernorProposalCreatedIterator struct {
	Event *DAOGovernorProposalCreated // Event containing the contract specifics and raw log

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
func (it *DAOGovernorProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorProposalCreated)
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
		it.Event = new(DAOGovernorProposalCreated)
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
func (it *DAOGovernorProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorProposalCreated represents a ProposalCreated event raised by the DAOGovernor contract.
type DAOGovernorProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_DAOGovernor *DAOGovernorFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*DAOGovernorProposalCreatedIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorProposalCreatedIterator{contract: _DAOGovernor.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_DAOGovernor *DAOGovernorFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *DAOGovernorProposalCreated) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorProposalCreated)
				if err := _DAOGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_DAOGovernor *DAOGovernorFilterer) ParseProposalCreated(log types.Log) (*DAOGovernorProposalCreated, error) {
	event := new(DAOGovernorProposalCreated)
	if err := _DAOGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the DAOGovernor contract.
type DAOGovernorProposalExecutedIterator struct {
	Event *DAOGovernorProposalExecuted // Event containing the contract specifics and raw log

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
func (it *DAOGovernorProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorProposalExecuted)
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
		it.Event = new(DAOGovernorProposalExecuted)
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
func (it *DAOGovernorProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorProposalExecuted represents a ProposalExecuted event raised by the DAOGovernor contract.
type DAOGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_DAOGovernor *DAOGovernorFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*DAOGovernorProposalExecutedIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorProposalExecutedIterator{contract: _DAOGovernor.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_DAOGovernor *DAOGovernorFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *DAOGovernorProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorProposalExecuted)
				if err := _DAOGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_DAOGovernor *DAOGovernorFilterer) ParseProposalExecuted(log types.Log) (*DAOGovernorProposalExecuted, error) {
	event := new(DAOGovernorProposalExecuted)
	if err := _DAOGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the DAOGovernor contract.
type DAOGovernorProposalQueuedIterator struct {
	Event *DAOGovernorProposalQueued // Event containing the contract specifics and raw log

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
func (it *DAOGovernorProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorProposalQueued)
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
		it.Event = new(DAOGovernorProposalQueued)
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
func (it *DAOGovernorProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorProposalQueued represents a ProposalQueued event raised by the DAOGovernor contract.
type DAOGovernorProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_DAOGovernor *DAOGovernorFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*DAOGovernorProposalQueuedIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorProposalQueuedIterator{contract: _DAOGovernor.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_DAOGovernor *DAOGovernorFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *DAOGovernorProposalQueued) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorProposalQueued)
				if err := _DAOGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
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

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_DAOGovernor *DAOGovernorFilterer) ParseProposalQueued(log types.Log) (*DAOGovernorProposalQueued, error) {
	event := new(DAOGovernorProposalQueued)
	if err := _DAOGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorProposalThresholdSetIterator is returned from FilterProposalThresholdSet and is used to iterate over the raw logs and unpacked data for ProposalThresholdSet events raised by the DAOGovernor contract.
type DAOGovernorProposalThresholdSetIterator struct {
	Event *DAOGovernorProposalThresholdSet // Event containing the contract specifics and raw log

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
func (it *DAOGovernorProposalThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorProposalThresholdSet)
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
		it.Event = new(DAOGovernorProposalThresholdSet)
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
func (it *DAOGovernorProposalThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorProposalThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorProposalThresholdSet represents a ProposalThresholdSet event raised by the DAOGovernor contract.
type DAOGovernorProposalThresholdSet struct {
	OldProposalThreshold *big.Int
	NewProposalThreshold *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProposalThresholdSet is a free log retrieval operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_DAOGovernor *DAOGovernorFilterer) FilterProposalThresholdSet(opts *bind.FilterOpts) (*DAOGovernorProposalThresholdSetIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorProposalThresholdSetIterator{contract: _DAOGovernor.contract, event: "ProposalThresholdSet", logs: logs, sub: sub}, nil
}

// WatchProposalThresholdSet is a free log subscription operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_DAOGovernor *DAOGovernorFilterer) WatchProposalThresholdSet(opts *bind.WatchOpts, sink chan<- *DAOGovernorProposalThresholdSet) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorProposalThresholdSet)
				if err := _DAOGovernor.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
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

// ParseProposalThresholdSet is a log parse operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_DAOGovernor *DAOGovernorFilterer) ParseProposalThresholdSet(log types.Log) (*DAOGovernorProposalThresholdSet, error) {
	event := new(DAOGovernorProposalThresholdSet)
	if err := _DAOGovernor.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorQuorumNumeratorUpdatedIterator is returned from FilterQuorumNumeratorUpdated and is used to iterate over the raw logs and unpacked data for QuorumNumeratorUpdated events raised by the DAOGovernor contract.
type DAOGovernorQuorumNumeratorUpdatedIterator struct {
	Event *DAOGovernorQuorumNumeratorUpdated // Event containing the contract specifics and raw log

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
func (it *DAOGovernorQuorumNumeratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorQuorumNumeratorUpdated)
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
		it.Event = new(DAOGovernorQuorumNumeratorUpdated)
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
func (it *DAOGovernorQuorumNumeratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorQuorumNumeratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the DAOGovernor contract.
type DAOGovernorQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterQuorumNumeratorUpdated is a free log retrieval operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_DAOGovernor *DAOGovernorFilterer) FilterQuorumNumeratorUpdated(opts *bind.FilterOpts) (*DAOGovernorQuorumNumeratorUpdatedIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorQuorumNumeratorUpdatedIterator{contract: _DAOGovernor.contract, event: "QuorumNumeratorUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumNumeratorUpdated is a free log subscription operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_DAOGovernor *DAOGovernorFilterer) WatchQuorumNumeratorUpdated(opts *bind.WatchOpts, sink chan<- *DAOGovernorQuorumNumeratorUpdated) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorQuorumNumeratorUpdated)
				if err := _DAOGovernor.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
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

// ParseQuorumNumeratorUpdated is a log parse operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_DAOGovernor *DAOGovernorFilterer) ParseQuorumNumeratorUpdated(log types.Log) (*DAOGovernorQuorumNumeratorUpdated, error) {
	event := new(DAOGovernorQuorumNumeratorUpdated)
	if err := _DAOGovernor.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorTimelockChangeIterator is returned from FilterTimelockChange and is used to iterate over the raw logs and unpacked data for TimelockChange events raised by the DAOGovernor contract.
type DAOGovernorTimelockChangeIterator struct {
	Event *DAOGovernorTimelockChange // Event containing the contract specifics and raw log

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
func (it *DAOGovernorTimelockChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorTimelockChange)
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
		it.Event = new(DAOGovernorTimelockChange)
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
func (it *DAOGovernorTimelockChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorTimelockChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorTimelockChange represents a TimelockChange event raised by the DAOGovernor contract.
type DAOGovernorTimelockChange struct {
	OldTimelock common.Address
	NewTimelock common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTimelockChange is a free log retrieval operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_DAOGovernor *DAOGovernorFilterer) FilterTimelockChange(opts *bind.FilterOpts) (*DAOGovernorTimelockChangeIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorTimelockChangeIterator{contract: _DAOGovernor.contract, event: "TimelockChange", logs: logs, sub: sub}, nil
}

// WatchTimelockChange is a free log subscription operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_DAOGovernor *DAOGovernorFilterer) WatchTimelockChange(opts *bind.WatchOpts, sink chan<- *DAOGovernorTimelockChange) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorTimelockChange)
				if err := _DAOGovernor.contract.UnpackLog(event, "TimelockChange", log); err != nil {
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

// ParseTimelockChange is a log parse operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_DAOGovernor *DAOGovernorFilterer) ParseTimelockChange(log types.Log) (*DAOGovernorTimelockChange, error) {
	event := new(DAOGovernorTimelockChange)
	if err := _DAOGovernor.contract.UnpackLog(event, "TimelockChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the DAOGovernor contract.
type DAOGovernorVoteCastIterator struct {
	Event *DAOGovernorVoteCast // Event containing the contract specifics and raw log

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
func (it *DAOGovernorVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorVoteCast)
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
		it.Event = new(DAOGovernorVoteCast)
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
func (it *DAOGovernorVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorVoteCast represents a VoteCast event raised by the DAOGovernor contract.
type DAOGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_DAOGovernor *DAOGovernorFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*DAOGovernorVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &DAOGovernorVoteCastIterator{contract: _DAOGovernor.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_DAOGovernor *DAOGovernorFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *DAOGovernorVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorVoteCast)
				if err := _DAOGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_DAOGovernor *DAOGovernorFilterer) ParseVoteCast(log types.Log) (*DAOGovernorVoteCast, error) {
	event := new(DAOGovernorVoteCast)
	if err := _DAOGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorVoteCastWithParamsIterator is returned from FilterVoteCastWithParams and is used to iterate over the raw logs and unpacked data for VoteCastWithParams events raised by the DAOGovernor contract.
type DAOGovernorVoteCastWithParamsIterator struct {
	Event *DAOGovernorVoteCastWithParams // Event containing the contract specifics and raw log

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
func (it *DAOGovernorVoteCastWithParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorVoteCastWithParams)
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
		it.Event = new(DAOGovernorVoteCastWithParams)
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
func (it *DAOGovernorVoteCastWithParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorVoteCastWithParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorVoteCastWithParams represents a VoteCastWithParams event raised by the DAOGovernor contract.
type DAOGovernorVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCastWithParams is a free log retrieval operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_DAOGovernor *DAOGovernorFilterer) FilterVoteCastWithParams(opts *bind.FilterOpts, voter []common.Address) (*DAOGovernorVoteCastWithParamsIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return &DAOGovernorVoteCastWithParamsIterator{contract: _DAOGovernor.contract, event: "VoteCastWithParams", logs: logs, sub: sub}, nil
}

// WatchVoteCastWithParams is a free log subscription operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_DAOGovernor *DAOGovernorFilterer) WatchVoteCastWithParams(opts *bind.WatchOpts, sink chan<- *DAOGovernorVoteCastWithParams, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorVoteCastWithParams)
				if err := _DAOGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
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

// ParseVoteCastWithParams is a log parse operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_DAOGovernor *DAOGovernorFilterer) ParseVoteCastWithParams(log types.Log) (*DAOGovernorVoteCastWithParams, error) {
	event := new(DAOGovernorVoteCastWithParams)
	if err := _DAOGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorVotingDelaySetIterator is returned from FilterVotingDelaySet and is used to iterate over the raw logs and unpacked data for VotingDelaySet events raised by the DAOGovernor contract.
type DAOGovernorVotingDelaySetIterator struct {
	Event *DAOGovernorVotingDelaySet // Event containing the contract specifics and raw log

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
func (it *DAOGovernorVotingDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorVotingDelaySet)
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
		it.Event = new(DAOGovernorVotingDelaySet)
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
func (it *DAOGovernorVotingDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorVotingDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorVotingDelaySet represents a VotingDelaySet event raised by the DAOGovernor contract.
type DAOGovernorVotingDelaySet struct {
	OldVotingDelay *big.Int
	NewVotingDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotingDelaySet is a free log retrieval operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_DAOGovernor *DAOGovernorFilterer) FilterVotingDelaySet(opts *bind.FilterOpts) (*DAOGovernorVotingDelaySetIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorVotingDelaySetIterator{contract: _DAOGovernor.contract, event: "VotingDelaySet", logs: logs, sub: sub}, nil
}

// WatchVotingDelaySet is a free log subscription operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_DAOGovernor *DAOGovernorFilterer) WatchVotingDelaySet(opts *bind.WatchOpts, sink chan<- *DAOGovernorVotingDelaySet) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorVotingDelaySet)
				if err := _DAOGovernor.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
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

// ParseVotingDelaySet is a log parse operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_DAOGovernor *DAOGovernorFilterer) ParseVotingDelaySet(log types.Log) (*DAOGovernorVotingDelaySet, error) {
	event := new(DAOGovernorVotingDelaySet)
	if err := _DAOGovernor.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOGovernorVotingPeriodSetIterator is returned from FilterVotingPeriodSet and is used to iterate over the raw logs and unpacked data for VotingPeriodSet events raised by the DAOGovernor contract.
type DAOGovernorVotingPeriodSetIterator struct {
	Event *DAOGovernorVotingPeriodSet // Event containing the contract specifics and raw log

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
func (it *DAOGovernorVotingPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOGovernorVotingPeriodSet)
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
		it.Event = new(DAOGovernorVotingPeriodSet)
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
func (it *DAOGovernorVotingPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOGovernorVotingPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOGovernorVotingPeriodSet represents a VotingPeriodSet event raised by the DAOGovernor contract.
type DAOGovernorVotingPeriodSet struct {
	OldVotingPeriod *big.Int
	NewVotingPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVotingPeriodSet is a free log retrieval operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_DAOGovernor *DAOGovernorFilterer) FilterVotingPeriodSet(opts *bind.FilterOpts) (*DAOGovernorVotingPeriodSetIterator, error) {

	logs, sub, err := _DAOGovernor.contract.FilterLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return &DAOGovernorVotingPeriodSetIterator{contract: _DAOGovernor.contract, event: "VotingPeriodSet", logs: logs, sub: sub}, nil
}

// WatchVotingPeriodSet is a free log subscription operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_DAOGovernor *DAOGovernorFilterer) WatchVotingPeriodSet(opts *bind.WatchOpts, sink chan<- *DAOGovernorVotingPeriodSet) (event.Subscription, error) {

	logs, sub, err := _DAOGovernor.contract.WatchLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOGovernorVotingPeriodSet)
				if err := _DAOGovernor.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
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

// ParseVotingPeriodSet is a log parse operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_DAOGovernor *DAOGovernorFilterer) ParseVotingPeriodSet(log types.Log) (*DAOGovernorVotingPeriodSet, error) {
	event := new(DAOGovernorVotingPeriodSet)
	if err := _DAOGovernor.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
