package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

// Metering client helpers. These are convenience wrappers around the hub's
// /api/metering/* endpoints for SDK users who want to preflight cost or trigger
// settlement. The server remains authoritative — a client preflight is advisory
// only, and the hub re-checks admission on every upload.

// MeteringPricing mirrors GET /api/metering/pricing.
type MeteringPricing struct {
	Enabled           bool   `json:"enabled"`
	ChargeWrites      bool   `json:"charge_writes"`
	ChargeReads       bool   `json:"charge_reads"`
	WriteBaseWei      string `json:"write_base_wei"`
	WritePerKBWei     string `json:"write_per_kb_wei"`
	ReadPerRequestWei string `json:"read_per_request_wei"`
}

// MeteringUsage mirrors GET /api/metering/usage.
type MeteringUsage struct {
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

// MeteringCanWrite mirrors GET /api/metering/can-write.
type MeteringCanWrite struct {
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

// MeteringSettleTx holds the on-chain tx hashes from an ERC-8183 settlement.
type MeteringSettleTx struct {
	Transfer  string `json:"transfer,omitempty"`
	CreateJob string `json:"create_job,omitempty"`
	SetBudget string `json:"set_budget,omitempty"`
	Fund      string `json:"fund,omitempty"`
	Submit    string `json:"submit,omitempty"`
}

// MeteringSettle mirrors POST /api/metering/settle.
type MeteringSettle struct {
	Owner            string            `json:"owner"`
	SettledAmountWei string            `json:"settled_amount_wei"`
	Mode             string            `json:"mode"`
	SettlementID     uint              `json:"settlement_id,omitempty"`
	Status           string            `json:"status"`
	JobID            string            `json:"job_id,omitempty"`
	ReportHash       string            `json:"report_hash,omitempty"`
	Tx               *MeteringSettleTx `json:"tx,omitempty"`
}

// GetMeteringPricing fetches the hub's fee schedule (public, no auth).
func GetMeteringPricing(baseUrl string) (*MeteringPricing, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	resb, err := Get(ctx, baseUrl+"/api/metering/pricing")
	if err != nil {
		return nil, err
	}
	var res MeteringPricing
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetMeteringUsage fetches an owner's usage summary (public, no auth).
func GetMeteringUsage(baseUrl, owner string) (*MeteringUsage, error) {
	q := url.Values{}
	q.Set("owner", owner)
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	resb, err := Get(ctx, baseUrl+"/api/metering/usage?"+q.Encode())
	if err != nil {
		return nil, err
	}
	var res MeteringUsage
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CanWrite asks the hub whether a write of the given size would be admitted
// (public, no auth). This is advisory; the hub re-checks on upload.
func CanWrite(baseUrl, owner string, bytes uint64) (*MeteringCanWrite, error) {
	q := url.Values{}
	q.Set("owner", owner)
	q.Set("bytes", strconv.FormatUint(bytes, 10))
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	resb, err := Get(ctx, baseUrl+"/api/metering/can-write?"+q.Encode())
	if err != nil {
		return nil, err
	}
	var res MeteringCanWrite
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SettleMetering triggers settlement for the signer's account. It requires a
// valid Auth; the hub settles the account matching the recovered signer.
func SettleMetering(baseUrl string, auth types.Auth) (*MeteringSettle, error) {
	if len(auth.Sign) == 0 {
		return nil, fmt.Errorf("settle requires a signed Auth")
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancel()
	resb, err := doRequest(ctx, baseUrl, "/api/metering/settle", "", auth, nil)
	if err != nil {
		return nil, err
	}
	var res MeteringSettle
	if err := json.Unmarshal(resb, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
