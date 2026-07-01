package metering

import (
	"math/big"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

func TestCanWriteDisabledAlwaysAllows(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:               false,
		WriteBaseWei:          big.NewInt(1000000),
		DefaultCreditLimitWei: big.NewInt(1), // tiny, but ignored when disabled
	})
	res, err := m.CanWrite("0x5555555555555555555555555555555555555555", 999999)
	if err != nil {
		t.Fatalf("CanWrite: %v", err)
	}
	if !res.Allowed || res.Reason != reasonMeteringDisabled {
		t.Fatalf("want allowed/metering_disabled, got %+v", res)
	}
}

func TestCanWriteWithinCreditLimit(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:               true,
		ChargeWrites:          true,
		WriteBaseWei:          big.NewInt(100),
		DefaultCreditLimitWei: big.NewInt(150),
	})
	// fresh owner, no unsettled fee: required = est = 100 <= 150
	res, err := m.CanWrite("0x6666666666666666666666666666666666666666", 1)
	if err != nil {
		t.Fatalf("CanWrite: %v", err)
	}
	if !res.Allowed || res.Reason != reasonAllowed {
		t.Fatalf("want allowed, got %+v", res)
	}
	if res.RequiredWei != "100" || res.CreditLimitWei != "150" {
		t.Errorf("required=%s limit=%s, want 100/150", res.RequiredWei, res.CreditLimitWei)
	}
}

func TestCanWriteExceedsCreditLimit(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:               true,
		ChargeWrites:          true,
		WriteBaseWei:          big.NewInt(100),
		DefaultCreditLimitWei: big.NewInt(150),
	})
	owner := "0x7777777777777777777777777777777777777777"
	// First write accrues 100 unsettled and creates the account with limit 150.
	if _, err := m.RecordWrite(owner, "b", "o", 1); err != nil {
		t.Fatalf("RecordWrite: %v", err)
	}
	// Next write: required = unsettled(100) + est(100) = 200 > 150 -> reject.
	res, err := m.CanWrite(owner, 1)
	if err != nil {
		t.Fatalf("CanWrite: %v", err)
	}
	if res.Allowed || res.Reason != reasonCreditLimitExceeded {
		t.Fatalf("want rejected/credit_limit_exceeded, got %+v", res)
	}
	if res.RequiredWei != "200" || res.UnsettledFeeWei != "100" {
		t.Errorf("required=%s unsettled=%s, want 200/100", res.RequiredWei, res.UnsettledFeeWei)
	}
}

func TestCanWriteZeroCreditLimitUnlimited(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:               true,
		ChargeWrites:          true,
		WriteBaseWei:          bi("1000000000000000000000"), // huge
		DefaultCreditLimitWei: big.NewInt(0),                // 0 == unlimited
	})
	res, err := m.CanWrite("0x8888888888888888888888888888888888888888", 100000)
	if err != nil {
		t.Fatalf("CanWrite: %v", err)
	}
	if !res.Allowed {
		t.Fatalf("zero credit limit should be unlimited, got %+v", res)
	}
}

func TestCanWriteDisabledAccountRejected(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:      true,
		ChargeWrites: true,
		WriteBaseWei: big.NewInt(100),
	})
	owner := "0x9999999999999999999999999999999999999999"
	if _, err := m.RecordWrite(owner, "b", "o", 1); err != nil {
		t.Fatalf("RecordWrite: %v", err)
	}
	// Disable the account.
	if err := m.db.Model(&types.MeterAccount{}).Where("owner = ?", owner).
		Updates(map[string]interface{}{"enabled": false, "status": statusDisabled}).Error; err != nil {
		t.Fatalf("disable: %v", err)
	}
	res, err := m.CanWrite(owner, 1)
	if err != nil {
		t.Fatalf("CanWrite: %v", err)
	}
	if res.Allowed || res.Reason != reasonAccountDisabled {
		t.Fatalf("want rejected/account_disabled, got %+v", res)
	}
}
