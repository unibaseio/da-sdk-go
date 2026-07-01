package metering

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unibaseio/da-sdk-go/contract/erc8183/go/erc8183"
)

// erc8183API is the subset of ERC-8183 operations settlement needs. It is an
// interface so the settlement orchestration can be unit-tested with a fake,
// without a live chain.
type erc8183API interface {
	CreateJob(ctx context.Context, expiredAt *big.Int, description string, providerAgentID *big.Int) (jobID *big.Int, tx common.Hash, err error)
	SetBudget(ctx context.Context, jobID, amount *big.Int) (common.Hash, error)
	Fund(ctx context.Context, jobID, expectedBudget *big.Int) (common.Hash, error)
	Submit(ctx context.Context, jobID *big.Int, deliverable [32]byte) (common.Hash, error)
}

// ERC8183Client drives the ERC-8183 escrow contract via the generated binding.
// It dials per call and uses the provider key to sign. All mutating calls wait
// for the transaction to be mined and return its hash.
type ERC8183Client struct {
	rpcURL        string
	contractAddr  common.Address
	evaluatorAddr common.Address
	tokenAddr     common.Address
	chainID       *big.Int
	sk            *ecdsa.PrivateKey
	provider      common.Address
}

var _ erc8183API = (*ERC8183Client)(nil)

// newERC8183Client validates config and builds a client. It requires the
// provider key (settlement submits signed transactions). It does NOT dial.
func newERC8183Client(cfg Config) (*ERC8183Client, error) {
	if cfg.ChainRPCURL == "" {
		return nil, fmt.Errorf("HUB_CHAIN_RPC_URL not set")
	}
	if !common.IsHexAddress(cfg.ERC8183ContractAddr) {
		return nil, fmt.Errorf("HUB_ERC8183_CONTRACT_ADDR invalid: %q", cfg.ERC8183ContractAddr)
	}
	if !common.IsHexAddress(cfg.ERC8183EvaluatorAddr) {
		return nil, fmt.Errorf("HUB_ERC8183_EVALUATOR_ADDR invalid: %q", cfg.ERC8183EvaluatorAddr)
	}
	if !common.IsHexAddress(cfg.ERC20TokenAddr) {
		return nil, fmt.Errorf("HUB_ERC20_TOKEN_ADDR invalid: %q", cfg.ERC20TokenAddr)
	}
	if cfg.ProviderPrivateKey == "" {
		return nil, fmt.Errorf("HUB_PROVIDER_PRIVATE_KEY not set")
	}
	sk, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.ProviderPrivateKey, "0x"))
	if err != nil {
		return nil, fmt.Errorf("provider private key: %w", err)
	}
	provider := crypto.PubkeyToAddress(sk.PublicKey)
	if common.IsHexAddress(cfg.ProviderAddress) {
		provider = common.HexToAddress(cfg.ProviderAddress)
	}
	return &ERC8183Client{
		rpcURL:        cfg.ChainRPCURL,
		contractAddr:  common.HexToAddress(cfg.ERC8183ContractAddr),
		evaluatorAddr: common.HexToAddress(cfg.ERC8183EvaluatorAddr),
		tokenAddr:     common.HexToAddress(cfg.ERC20TokenAddr),
		chainID:       big.NewInt(cfg.ChainID),
		sk:            sk,
		provider:      provider,
	}, nil
}

func (c *ERC8183Client) dial(ctx context.Context) (*erc8183.ERC8183, *ethclient.Client, error) {
	client, err := ethclient.DialContext(ctx, c.rpcURL)
	if err != nil {
		return nil, nil, err
	}
	inst, err := erc8183.NewERC8183(c.contractAddr, client)
	if err != nil {
		client.Close()
		return nil, nil, err
	}
	return inst, client, nil
}

func (c *ERC8183Client) transactor(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(c.sk, c.chainID)
	if err != nil {
		return nil, err
	}
	auth.Context = ctx
	return auth, nil
}

// CreateJob creates an escrow job with the provider as worker and the
// configured evaluator. It waits for the receipt and extracts the job id from
// the JobCreated event. The evaluator address is also used as the hook, mirroring
// the reference implementation.
func (c *ERC8183Client) CreateJob(ctx context.Context, expiredAt *big.Int, description string, providerAgentID *big.Int) (*big.Int, common.Hash, error) {
	inst, client, err := c.dial(ctx)
	if err != nil {
		return nil, common.Hash{}, err
	}
	defer client.Close()

	auth, err := c.transactor(ctx)
	if err != nil {
		return nil, common.Hash{}, err
	}

	tx, err := inst.CreateJob(auth, c.provider, c.evaluatorAddr, expiredAt, description, c.evaluatorAddr, providerAgentID)
	if err != nil {
		return nil, common.Hash{}, err
	}

	rcpt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return nil, tx.Hash(), err
	}
	if rcpt.Status != 1 {
		return nil, tx.Hash(), fmt.Errorf("createJob reverted (tx %s)", tx.Hash().Hex())
	}

	for _, lg := range rcpt.Logs {
		ev, perr := inst.ParseJobCreated(*lg)
		if perr == nil {
			return ev.JobId, tx.Hash(), nil
		}
	}
	return nil, tx.Hash(), fmt.Errorf("JobCreated event not found in receipt (tx %s)", tx.Hash().Hex())
}

// SetBudget sets the job budget in the configured payment token.
func (c *ERC8183Client) SetBudget(ctx context.Context, jobID, amount *big.Int) (common.Hash, error) {
	return c.send(ctx, func(inst *erc8183.ERC8183, auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return inst.SetBudget(auth, jobID, c.tokenAddr, amount, []byte{})
	}, "setBudget")
}

// Fund funds the job up to the expected budget.
func (c *ERC8183Client) Fund(ctx context.Context, jobID, expectedBudget *big.Int) (common.Hash, error) {
	return c.send(ctx, func(inst *erc8183.ERC8183, auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return inst.Fund(auth, jobID, expectedBudget, []byte{})
	}, "fund")
}

// Submit submits the deliverable (report hash) for the job.
func (c *ERC8183Client) Submit(ctx context.Context, jobID *big.Int, deliverable [32]byte) (common.Hash, error) {
	return c.send(ctx, func(inst *erc8183.ERC8183, auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return inst.Submit(auth, jobID, deliverable, []byte{})
	}, "submit")
}

// send dials, signs, submits, and waits for the receipt of a mutating call.
func (c *ERC8183Client) send(ctx context.Context, do func(*erc8183.ERC8183, *bind.TransactOpts) (*ethtypes.Transaction, error), name string) (common.Hash, error) {
	inst, client, err := c.dial(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	defer client.Close()

	auth, err := c.transactor(ctx)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := do(inst, auth)
	if err != nil {
		return common.Hash{}, err
	}

	rcpt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return tx.Hash(), err
	}
	if rcpt.Status != 1 {
		return tx.Hash(), fmt.Errorf("%s reverted (tx %s)", name, tx.Hash().Hex())
	}
	return tx.Hash(), nil
}

// ----------------------------------------------------------------------------
// Settlement report + hash
// ----------------------------------------------------------------------------

// SettlementReport is the canonical settlement payload whose sha256 is submitted
// on-chain as the ERC-8183 deliverable. Fields are ordered alphabetically so
// json.Marshal produces a stable, compact canonical form (locked by a test).
type SettlementReport struct {
	AmountWei    string `json:"amount_wei"`
	BytesWritten uint64 `json:"bytes_written"`
	FromEventID  uint   `json:"from_event_id"`
	Owner        string `json:"owner"`
	Reads        uint64 `json:"reads"`
	Timestamp    int64  `json:"timestamp"`
	ToEventID    uint   `json:"to_event_id"`
	Type         string `json:"type"`
	Writes       uint64 `json:"writes"`
}

// ReportType is the fixed discriminator for hub metering settlements.
const ReportType = "da-hub-metering-settlement"

// CanonicalJSON returns the deterministic byte encoding of the report used for
// hashing. json.Marshal emits struct fields in declaration order (alphabetical
// here) with no insignificant whitespace, so the output is stable.
func (r SettlementReport) CanonicalJSON() ([]byte, error) {
	return json.Marshal(r)
}

// Hash returns sha256(CanonicalJSON) as a [32]byte deliverable and its 0x-hex
// string form.
func (r SettlementReport) Hash() ([32]byte, string, error) {
	b, err := r.CanonicalJSON()
	if err != nil {
		return [32]byte{}, "", err
	}
	sum := sha256.Sum256(b)
	return sum, "0x" + common.Bytes2Hex(sum[:]), nil
}
