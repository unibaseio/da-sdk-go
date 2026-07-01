package types

import (
	"time"

	"gorm.io/gorm"
)

// Metering ledger models. All token amounts are stored as base-unit (wei)
// decimal strings, never floats. Arithmetic is done with math/big in the
// metering manager. Empty amount strings are treated as "0".

// MeterAccount is the per-owner accounting summary. Owner is the canonical
// lowercase 0x address.
type MeterAccount struct {
	gorm.Model
	Owner   string `gorm:"uniqueIndex;size:64"`
	Enabled bool
	Status  string // active, disabled

	TotalReads        uint64
	TotalWrites       uint64
	TotalBytesWritten uint64

	TotalFeeWei     string // lifetime fee accrued (wei decimal string)
	UnsettledFeeWei string // fee not yet settled on-chain (wei decimal string)
	CreditLimitWei  string // 0 == unlimited local credit
	LastSettledAt   *time.Time
}

// MeterEvent is one billable (or zero-fee observational) usage event.
type MeterEvent struct {
	gorm.Model
	Owner        string `gorm:"index:idx_meter_events_owner_status,priority:1;size:64"`
	Operation    string // write, read
	ObjectName   string
	Bucket       string
	Bytes        uint64
	FeeWei       string
	Status       string `gorm:"index:idx_meter_events_owner_status,priority:2"` // unsettled, settling, settled
	SettlementID uint   `gorm:"index"`
}

// MeterSettlement records one settlement attempt over a range of events.
type MeterSettlement struct {
	gorm.Model
	Owner       string `gorm:"index:idx_meter_settlements_owner_status,priority:1;size:64"`
	AmountWei   string
	FromEventID uint
	ToEventID   uint
	Status      string `gorm:"index:idx_meter_settlements_owner_status,priority:2"` // pending, submitting, confirmed, failed
	Mode        string // offchain, erc8183

	TransferTx  string
	CreateJobTx string
	FundTx      string
	SubmitTx    string
	JobID       string
	ReportHash  string
	Error       string
}
