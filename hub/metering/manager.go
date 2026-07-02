package metering

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/log"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"gorm.io/gorm"
)

var logger = log.Logger("metering")

const (
	opWrite = "write"
	opRead  = "read"

	statusActive   = "active"
	statusDisabled = "disabled"

	eventUnsettled = "unsettled"
	eventSettling  = "settling"
	eventSettled   = "settled"

	// CanWrite reasons
	reasonMeteringDisabled    = "metering_disabled"
	reasonAccountDisabled     = "account_disabled"
	reasonCreditLimitExceeded = "credit_limit_exceeded"
	reasonInsufficientAllow   = "insufficient_allowance"
	reasonInsufficientBalance = "insufficient_balance"
	reasonChainCheckFailed    = "chain_check_failed"
	reasonAllowed             = "allowed"
)

// Manager owns pricing, the usage ledger, and (in later phases) write-admission
// checks and settlement. It is safe for concurrent use.
type Manager struct {
	db  *gorm.DB
	cfg Config

	// accountLocks serializes read-modify-write updates of a single owner's
	// MeterAccount totals. Token amounts can exceed uint64, so totals are kept
	// as decimal strings and updated with math/big in Go rather than via SQL
	// integer arithmetic. This assumes a single hub instance (see plan).
	accountLocks keyedMutex

	// erc20 is the token client used for chain balance/allowance checks and
	// settlement transfers. It is nil when chain features are off or config is
	// invalid; callers must handle nil. Held as an interface so tests can inject
	// a fake without a live RPC node.
	erc20 erc20API

	// erc8183 is the escrow client used by erc8183 settlement mode. Nil unless
	// SettlementMode=erc8183 and config is valid. Interface for testability.
	erc8183 erc8183API

	// provider is the spender/recipient address for allowance checks and
	// transferFrom. Derived from HUB_PROVIDER_ADDRESS or the provider key.
	provider common.Address
}

// NewManager builds a Manager. db must be a live *gorm.DB with the metering
// tables migrated. Nil wei amounts are normalized to zero so fee arithmetic
// never dereferences a nil *big.Int.
func NewManager(db *gorm.DB, cfg Config) *Manager {
	cfg.WriteBaseWei = orZero(cfg.WriteBaseWei)
	cfg.WritePerKBWei = orZero(cfg.WritePerKBWei)
	cfg.ReadPerRequestWei = orZero(cfg.ReadPerRequestWei)
	cfg.DefaultCreditLimitWei = orZero(cfg.DefaultCreditLimitWei)
	cfg.SettleThresholdWei = orZero(cfg.SettleThresholdWei)

	m := &Manager{db: db, cfg: cfg}

	// Build the token client when chain features are requested. A build error
	// is logged, not fatal: CanWrite reports chain_check_failed and settlement
	// surfaces the misconfiguration, but the hub still starts.
	if cfg.CheckChain || cfg.SettlementMode == modeERC8183 {
		ec, err := newERC20Client(cfg)
		if err != nil {
			logger.Warnf("metering: erc20 client not available: %v", err)
		} else {
			m.erc20 = ec
			m.provider = resolveProvider(cfg, ec)
			if (m.provider == common.Address{}) {
				logger.Warnf("metering: provider address unset; allowance checks will use zero address")
			}
		}
	}

	if cfg.SettlementMode == modeERC8183 {
		ec8, err := newERC8183Client(cfg)
		if err != nil {
			logger.Warnf("metering: erc8183 client not available: %v", err)
		} else {
			m.erc8183 = ec8
			if (m.provider == common.Address{}) {
				m.provider = ec8.provider
			}
		}
	}

	// Clean up settlements a previous process left mid-flight, before any
	// traffic can observe (or settle around) their reserved events.
	if cfg.Enabled {
		m.recoverInterruptedSettlements()
	}
	return m
}

// resolveProvider picks the address used as the ERC-20 spender (allowance
// checks) and transferFrom recipient. When a signing key is configured it wins:
// transferFrom is signed by the key, so allowances must be granted to the
// key-derived address or the pull reverts. HUB_PROVIDER_ADDRESS applies only to
// key-less (read-only chain check) deployments, where it declares the address
// that will eventually spend.
func resolveProvider(cfg Config, ec *ERC20Client) common.Address {
	keyAddr, hasKey := ec.providerFromKey()
	if common.IsHexAddress(cfg.ProviderAddress) {
		cfgAddr := common.HexToAddress(cfg.ProviderAddress)
		if hasKey && cfgAddr != keyAddr {
			logger.Warnf("metering: HUB_PROVIDER_ADDRESS %s differs from the key-derived address %s; using the key address (it is the actual transferFrom spender)", cfg.ProviderAddress, keyAddr.Hex())
			return keyAddr
		}
		return cfgAddr
	}
	if hasKey {
		return keyAddr
	}
	return common.Address{}
}

// orZero returns n, or a fresh zero *big.Int when n is nil.
func orZero(n *big.Int) *big.Int {
	if n == nil {
		return big.NewInt(0)
	}
	return n
}

// Enabled reports whether metering is on. It is nil-safe so callers can guard
// with s.metering.Enabled() even before initialization.
func (m *Manager) Enabled() bool {
	return m != nil && m.cfg.Enabled
}

// Config returns a copy of the loaded configuration.
func (m *Manager) Config() Config {
	return m.cfg
}

// ShouldChargeWrites reports whether write events should be recorded. Nil-safe.
func (m *Manager) ShouldChargeWrites() bool {
	return m.Enabled() && m.cfg.ChargeWrites
}

// ShouldChargeReads reports whether read events should be recorded. Nil-safe.
func (m *Manager) ShouldChargeReads() bool {
	return m.Enabled() && m.cfg.ChargeReads
}

// ----------------------------------------------------------------------------
// Pricing
// ----------------------------------------------------------------------------

// Pricing is the immutable fee schedule.
type Pricing struct {
	WriteBaseWei      *big.Int
	WritePerKBWei     *big.Int
	ReadPerRequestWei *big.Int
}

// Pricing returns copies of the configured fee parameters.
func (m *Manager) Pricing() Pricing {
	return Pricing{
		WriteBaseWei:      new(big.Int).Set(m.cfg.WriteBaseWei),
		WritePerKBWei:     new(big.Int).Set(m.cfg.WritePerKBWei),
		ReadPerRequestWei: new(big.Int).Set(m.cfg.ReadPerRequestWei),
	}
}

// PriceWrite computes write_base + ceil(payloadBytes/1024) * write_per_kb.
// Returns a fresh *big.Int (never nil).
func (m *Manager) PriceWrite(payloadBytes uint64) *big.Int {
	fee := new(big.Int).Set(m.cfg.WriteBaseWei)
	if payloadBytes > 0 {
		billableKB := (payloadBytes + 1023) / 1024 // ceil
		kbFee := new(big.Int).Mul(m.cfg.WritePerKBWei, new(big.Int).SetUint64(billableKB))
		fee.Add(fee, kbFee)
	}
	return fee
}

// PriceRead returns the flat per-request read fee.
func (m *Manager) PriceRead() *big.Int {
	return new(big.Int).Set(m.cfg.ReadPerRequestWei)
}

// ----------------------------------------------------------------------------
// Ledger recording
// ----------------------------------------------------------------------------

// RecordWrite creates an unsettled write event and increments the owner's
// account totals atomically. Callers MUST only invoke this after a successful
// storage write. Returns the created event.
func (m *Manager) RecordWrite(owner, bucket, object string, bytes uint64) (*types.MeterEvent, error) {
	fee := m.PriceWrite(bytes)
	return m.record(opWrite, owner, bucket, object, bytes, fee)
}

// RecordRead creates an unsettled read event and increments read totals.
func (m *Manager) RecordRead(owner, bucket, object string, bytes uint64) (*types.MeterEvent, error) {
	fee := m.PriceRead()
	return m.record(opRead, owner, bucket, object, bytes, fee)
}

func (m *Manager) record(op, owner, bucket, object string, bytes uint64, fee *big.Int) (*types.MeterEvent, error) {
	owner = strings.ToLower(owner)
	feeStr := fee.String()

	// Serialize per-owner so concurrent writes don't lose increments.
	unlock := m.accountLocks.lock(owner)
	defer unlock()

	var event types.MeterEvent
	err := m.db.Transaction(func(tx *gorm.DB) error {
		event = types.MeterEvent{
			Owner:      owner,
			Operation:  op,
			ObjectName: object,
			Bucket:     bucket,
			Bytes:      bytes,
			FeeWei:     feeStr,
			Status:     eventUnsettled,
		}
		if err := tx.Create(&event).Error; err != nil {
			return err
		}

		var acct types.MeterAccount
		res := tx.Where("owner = ?", owner).First(&acct)
		if res.Error != nil {
			if res.Error != gorm.ErrRecordNotFound {
				return res.Error
			}
			acct = types.MeterAccount{
				Owner:           owner,
				Enabled:         true,
				Status:          statusActive,
				TotalFeeWei:     "0",
				UnsettledFeeWei: "0",
				CreditLimitWei:  m.cfg.DefaultCreditLimitWei.String(),
			}
			if err := tx.Create(&acct).Error; err != nil {
				return err
			}
		}

		if op == opWrite {
			acct.TotalWrites++
			acct.TotalBytesWritten += bytes
		} else {
			acct.TotalReads++
		}
		acct.TotalFeeWei = addWei(acct.TotalFeeWei, fee).String()
		acct.UnsettledFeeWei = addWei(acct.UnsettledFeeWei, fee).String()

		return tx.Model(&acct).Where("owner = ?", owner).Updates(map[string]interface{}{
			"total_reads":         acct.TotalReads,
			"total_writes":        acct.TotalWrites,
			"total_bytes_written": acct.TotalBytesWritten,
			"total_fee_wei":       acct.TotalFeeWei,
			"unsettled_fee_wei":   acct.UnsettledFeeWei,
		}).Error
	})
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// ----------------------------------------------------------------------------
// Usage
// ----------------------------------------------------------------------------

// UsageResponse is the read model returned by GET /api/metering/usage.
type UsageResponse struct {
	Owner             string     `json:"owner"`
	Enabled           bool       `json:"enabled"`
	Status            string     `json:"status"`
	TotalReads        uint64     `json:"total_reads"`
	TotalWrites       uint64     `json:"total_writes"`
	TotalBytesWritten uint64     `json:"total_bytes_written"`
	TotalFeeWei       string     `json:"total_fee_wei"`
	UnsettledFeeWei   string     `json:"unsettled_fee_wei"`
	CreditLimitWei    string     `json:"credit_limit_wei"`
	LastSettledAt     *time.Time `json:"last_settled_at,omitempty"`
}

// GetUsage returns the current ledger summary for owner. A missing account is
// reported as an empty (zeroed) usage rather than an error.
func (m *Manager) GetUsage(owner string) (*UsageResponse, error) {
	owner = strings.ToLower(owner)
	var acct types.MeterAccount
	res := m.db.Where("owner = ?", owner).First(&acct)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return &UsageResponse{
				Owner:           owner,
				Enabled:         m.cfg.Enabled,
				Status:          statusActive,
				TotalFeeWei:     "0",
				UnsettledFeeWei: "0",
				CreditLimitWei:  m.cfg.DefaultCreditLimitWei.String(),
			}, nil
		}
		return nil, res.Error
	}
	return &UsageResponse{
		Owner:             acct.Owner,
		Enabled:           acct.Enabled,
		Status:            acct.Status,
		TotalReads:        acct.TotalReads,
		TotalWrites:       acct.TotalWrites,
		TotalBytesWritten: acct.TotalBytesWritten,
		TotalFeeWei:       normWei(acct.TotalFeeWei),
		UnsettledFeeWei:   normWei(acct.UnsettledFeeWei),
		CreditLimitWei:    normWei(acct.CreditLimitWei),
		LastSettledAt:     acct.LastSettledAt,
	}, nil
}

// ----------------------------------------------------------------------------
// Write admission
// ----------------------------------------------------------------------------

// CanWriteResult is the admission decision for a prospective write.
type CanWriteResult struct {
	Allowed         bool   `json:"allowed"`
	Reason          string `json:"reason"`
	RequiredWei     string `json:"required_wei"`
	EstimatedFeeWei string `json:"estimated_fee_wei"`
	UnsettledFeeWei string `json:"unsettled_fee_wei"`
	CreditLimitWei  string `json:"credit_limit_wei"`
	BalanceWei      string `json:"balance_wei,omitempty"`
	AllowanceWei    string `json:"allowance_wei,omitempty"`
	Action          string `json:"action,omitempty"`
}

// CanWrite decides whether owner may perform a write of estimatedBytes.
//
//	estimated_fee = PriceWrite(estimatedBytes)
//	required      = unsettled_fee + estimated_fee
//
// Rejects when the account is disabled, or (credit_limit > 0 and required >
// credit_limit). Chain balance/allowance checks are applied only when
// HUB_METERING_CHECK_CHAIN=true (added in a later phase). When metering is
// disabled it always allows. A non-nil error means a local lookup failed and
// the caller should decide its own fail-open/closed policy.
func (m *Manager) CanWrite(owner string, estimatedBytes uint64) (*CanWriteResult, error) {
	owner = strings.ToLower(owner)
	est := m.PriceWrite(estimatedBytes)

	res := &CanWriteResult{
		EstimatedFeeWei: est.String(),
	}

	if !m.Enabled() {
		res.Allowed = true
		res.Reason = reasonMeteringDisabled
		res.RequiredWei = est.String()
		res.UnsettledFeeWei = "0"
		res.CreditLimitWei = "0"
		return res, nil
	}

	unsettled := big.NewInt(0)
	creditLimit := new(big.Int).Set(m.cfg.DefaultCreditLimitWei)
	disabled := false

	var acct types.MeterAccount
	r := m.db.Where("owner = ?", owner).First(&acct)
	if r.Error == nil {
		unsettled = parseWei(acct.UnsettledFeeWei)
		creditLimit = parseWei(acct.CreditLimitWei)
		disabled = !acct.Enabled || acct.Status == statusDisabled
	} else if r.Error != gorm.ErrRecordNotFound {
		return nil, r.Error
	}

	required := new(big.Int).Add(unsettled, est)
	res.RequiredWei = required.String()
	res.UnsettledFeeWei = unsettled.String()
	res.CreditLimitWei = creditLimit.String()

	if disabled {
		res.Allowed = false
		res.Reason = reasonAccountDisabled
		res.Action = "contact_provider"
		return res, nil
	}

	if creditLimit.Sign() > 0 && required.Cmp(creditLimit) > 0 {
		res.Allowed = false
		res.Reason = reasonCreditLimitExceeded
		res.Action = "settle"
		return res, nil
	}

	// Chain balance/allowance checks run only when HUB_METERING_CHECK_CHAIN is
	// on. Otherwise local credit is authoritative and no RPC call is made.
	if m.cfg.CheckChain {
		m.applyChainChecks(owner, required, res)
		return res, nil
	}

	res.Allowed = true
	res.Reason = reasonAllowed
	return res, nil
}

// applyChainChecks reads on-chain balance and allowance and sets the decision
// on res. Any RPC/config problem is reported as chain_check_failed (a refusal
// surfaced as 402), never as a Go error, so upload handlers do not return 500.
func (m *Manager) applyChainChecks(owner string, required *big.Int, res *CanWriteResult) {
	if m.erc20 == nil {
		res.Allowed = false
		res.Reason = reasonChainCheckFailed
		res.Action = "check_provider_config"
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	ownerAddr := common.HexToAddress(owner)

	bal, err := m.erc20.BalanceOf(ctx, ownerAddr)
	if err != nil {
		logger.Warnf("metering: balanceOf(%s) failed: %v", owner, err)
		res.Allowed = false
		res.Reason = reasonChainCheckFailed
		res.Action = "retry"
		return
	}
	res.BalanceWei = bal.String()

	allow, err := m.erc20.Allowance(ctx, ownerAddr, m.provider)
	if err != nil {
		logger.Warnf("metering: allowance(%s) failed: %v", owner, err)
		res.Allowed = false
		res.Reason = reasonChainCheckFailed
		res.Action = "retry"
		return
	}
	res.AllowanceWei = allow.String()

	if bal.Cmp(required) < 0 {
		res.Allowed = false
		res.Reason = reasonInsufficientBalance
		res.Action = "deposit"
		return
	}
	if allow.Cmp(required) < 0 {
		res.Allowed = false
		res.Reason = reasonInsufficientAllow
		res.Action = "approve"
		return
	}

	res.Allowed = true
	res.Reason = reasonAllowed
}

// ----------------------------------------------------------------------------
// wei string helpers
// ----------------------------------------------------------------------------

// parseWei converts a stored decimal string to *big.Int; empty/invalid -> 0.
func parseWei(s string) *big.Int {
	if s == "" {
		return big.NewInt(0)
	}
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		logger.Warnf("invalid wei string in ledger: %q, treating as 0", s)
		return big.NewInt(0)
	}
	return n
}

// normWei normalizes a stored wei string for output ("" -> "0").
func normWei(s string) string {
	return parseWei(s).String()
}

// addWei returns base + delta as a *big.Int.
func addWei(base string, delta *big.Int) *big.Int {
	return new(big.Int).Add(parseWei(base), delta)
}

// ----------------------------------------------------------------------------
// keyedMutex
// ----------------------------------------------------------------------------

// keyedMutex hands out one mutex per string key so callers can serialize work
// on a single key while allowing different keys to proceed concurrently.
// Entries are reference-counted and removed once unused, so the map does not
// grow without bound as new owners appear.
type keyedMutex struct {
	mu sync.Mutex
	m  map[string]*keyedMutexEntry
}

type keyedMutexEntry struct {
	mu   sync.Mutex
	refs int
}

func (k *keyedMutex) lock(key string) func() {
	k.mu.Lock()
	if k.m == nil {
		k.m = make(map[string]*keyedMutexEntry)
	}
	e, ok := k.m[key]
	if !ok {
		e = &keyedMutexEntry{}
		k.m[key] = e
	}
	e.refs++
	k.mu.Unlock()

	e.mu.Lock()
	return func() {
		e.mu.Unlock()
		k.mu.Lock()
		e.refs--
		if e.refs == 0 {
			delete(k.m, key)
		}
		k.mu.Unlock()
	}
}
