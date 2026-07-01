package metering

import (
	"math/big"
	"strings"
	"sync"
	"time"

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
}

// NewManager builds a Manager. db must be a live *gorm.DB with the metering
// tables migrated.
func NewManager(db *gorm.DB, cfg Config) *Manager {
	return &Manager{db: db, cfg: cfg}
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
type keyedMutex struct {
	mu sync.Mutex
	m  map[string]*sync.Mutex
}

func (k *keyedMutex) lock(key string) func() {
	k.mu.Lock()
	if k.m == nil {
		k.m = make(map[string]*sync.Mutex)
	}
	mu, ok := k.m[key]
	if !ok {
		mu = &sync.Mutex{}
		k.m[key] = mu
	}
	k.mu.Unlock()

	mu.Lock()
	return mu.Unlock
}
