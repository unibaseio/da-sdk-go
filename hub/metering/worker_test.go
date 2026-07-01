package metering

import (
	"math/big"
	"testing"
	"time"
)

// runOnce settles accounts at/above the threshold and ignores those below.
func TestWorkerRunOnceThreshold(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:            true,
		ChargeWrites:       true,
		WriteBaseWei:       bi("1000000000000000000"), // 1e18 per write
		SettleThresholdWei: bi("2000000000000000000"), // threshold 2e18
	})

	// rich: 3 writes -> 3e18 (>= threshold, should settle)
	rich := "0xaaaa1111111111111111111111111111111111aa"
	for i := 0; i < 3; i++ {
		if _, err := m.RecordWrite(rich, "b", "o", 10); err != nil {
			t.Fatalf("RecordWrite: %v", err)
		}
	}
	// poor: 1 write -> 1e18 (< threshold, should be ignored)
	poor := "0xbbbb2222222222222222222222222222222222bb"
	if _, err := m.RecordWrite(poor, "b", "o", 10); err != nil {
		t.Fatalf("RecordWrite: %v", err)
	}

	w := m.NewWorker()
	w.runOnce()

	ur, _ := m.GetUsage(rich)
	if ur.UnsettledFeeWei != "0" {
		t.Errorf("rich account not settled: unsettled = %s", ur.UnsettledFeeWei)
	}
	up, _ := m.GetUsage(poor)
	if up.UnsettledFeeWei != "1000000000000000000" {
		t.Errorf("poor account should be untouched: unsettled = %s, want 1e18", up.UnsettledFeeWei)
	}
}

// With threshold 0, any positive unsettled fee is settled.
func TestWorkerRunOnceZeroThresholdSettlesAll(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:            true,
		ChargeWrites:       true,
		WriteBaseWei:       big.NewInt(100),
		SettleThresholdWei: big.NewInt(0),
	})
	owner := "0xcccc3333333333333333333333333333333333cc"
	if _, err := m.RecordWrite(owner, "b", "o", 10); err != nil {
		t.Fatalf("RecordWrite: %v", err)
	}
	m.NewWorker().runOnce()

	u, _ := m.GetUsage(owner)
	if u.UnsettledFeeWei != "0" {
		t.Errorf("unsettled = %s, want 0", u.UnsettledFeeWei)
	}
}

// A settle failure for one account must not abort the scan or panic.
func TestWorkerRunOnceErrorDoesNotAbortScan(t *testing.T) {
	// erc8183 mode with a failing fund step -> Settle returns an error.
	fail8183 := &fakeERC8183{jobID: big.NewInt(1), failAt: "fund"}
	m := newTestManager(t, Config{
		Enabled:            true,
		ChargeWrites:       true,
		WriteBaseWei:       big.NewInt(100),
		SettlementMode:     modeERC8183,
		SettleThresholdWei: big.NewInt(0),
	})
	m.erc20 = &fakeERC20{}
	m.erc8183 = fail8183

	owner := "0xdddd4444444444444444444444444444444444dd"
	if _, err := m.RecordWrite(owner, "b", "o", 10); err != nil {
		t.Fatalf("RecordWrite: %v", err)
	}

	// must not panic; debt must remain
	m.NewWorker().runOnce()

	u, _ := m.GetUsage(owner)
	if u.UnsettledFeeWei == "0" {
		t.Error("debt cleared despite chain failure")
	}
}

// Start/Stop must be clean and not hang.
func TestWorkerStartStop(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:           true,
		ChargeWrites:      true,
		WriteBaseWei:      big.NewInt(100),
		SettleIntervalSec: 3600, // long; we only test lifecycle
	})
	w := m.NewWorker()
	w.Start()

	done := make(chan struct{})
	go func() { w.Stop(); close(done) }()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("Worker.Stop did not return in time")
	}
}
