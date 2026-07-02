package metering

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// interruptSettlement simulates a crash mid-settlement: events reserved as
// "settling" under a settlement row left in the given status, with transferTx
// recorded (or not) depending on how far the chain sequence got.
func interruptSettlement(t *testing.T, m *Manager, owner, transferTx string) types.MeterSettlement {
	t.Helper()
	var events []types.MeterEvent
	if err := m.db.Where("owner = ? AND status = ?", owner, eventUnsettled).
		Order("id asc").Find(&events).Error; err != nil || len(events) == 0 {
		t.Fatalf("no unsettled events to reserve: %v", err)
	}
	amount := big.NewInt(0)
	for _, e := range events {
		amount.Add(amount, parseWei(e.FeeWei))
	}
	s := types.MeterSettlement{
		Owner:       owner,
		AmountWei:   amount.String(),
		FromEventID: events[0].ID,
		ToEventID:   events[len(events)-1].ID,
		Status:      settlementPending,
		Mode:        modeERC8183,
		TransferTx:  transferTx,
	}
	if transferTx != "" {
		s.Status = settlementSubmitting
	}
	if err := m.db.Create(&s).Error; err != nil {
		t.Fatalf("create settlement: %v", err)
	}
	if err := m.db.Model(&types.MeterEvent{}).
		Where("owner = ? AND status = ?", owner, eventUnsettled).
		Update("status", eventSettling).Error; err != nil {
		t.Fatalf("reserve events: %v", err)
	}
	return s
}

// A settlement interrupted before the transfer confirmed must be reverted:
// events back to unsettled, settlement failed, debt intact and retryable.
func TestRecoverInterruptedBeforeTransfer(t *testing.T) {
	cfg := Config{Enabled: true, ChargeWrites: true, WriteBaseWei: bi("1000000000000000000"), SettlementMode: modeERC8183}
	m := newTestManager(t, cfg)
	owner := "0x3030303030303030303030303030303030303030"
	for i := 0; i < 2; i++ {
		if _, err := m.RecordWrite(owner, "b", "o", 10); err != nil {
			t.Fatalf("RecordWrite: %v", err)
		}
	}
	s := interruptSettlement(t, m, owner, "")

	// Simulate restart: a fresh Manager over the same DB runs recovery.
	m2 := NewManager(m.db, m.cfg)

	var unsettled int64
	m2.db.Model(&types.MeterEvent{}).Where("owner = ? AND status = ?", owner, eventUnsettled).Count(&unsettled)
	if unsettled != 2 {
		t.Errorf("%d events unsettled after recovery, want 2", unsettled)
	}
	var srec types.MeterSettlement
	m2.db.First(&srec, s.ID)
	if srec.Status != settlementFailed || srec.Error == "" {
		t.Errorf("settlement = %+v, want failed with error", srec)
	}
	u, _ := m2.GetUsage(owner)
	if u.UnsettledFeeWei != "2000000000000000000" {
		t.Errorf("unsettled fee = %s, want 2e18 (debt intact)", u.UnsettledFeeWei)
	}

	// The recovered debt must be settleable again.
	m2.erc20 = &fakeERC20{}
	m2.erc8183 = &fakeERC8183{jobID: big.NewInt(1)}
	m2.provider = common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	resp, err := m2.Settle(owner)
	if err != nil || resp.Status != settlementConfirmed {
		t.Fatalf("retry settle: %v (%+v)", err, resp)
	}
	u2, _ := m2.GetUsage(owner)
	if u2.UnsettledFeeWei != "0" {
		t.Errorf("unsettled fee = %s after retry, want 0", u2.UnsettledFeeWei)
	}
}

// A settlement interrupted after the transfer confirmed means the user paid:
// recovery finalizes it, clearing the covered debt exactly once.
func TestRecoverInterruptedAfterTransfer(t *testing.T) {
	cfg := Config{Enabled: true, ChargeWrites: true, WriteBaseWei: bi("1000000000000000000"), SettlementMode: modeERC8183}
	m := newTestManager(t, cfg)
	owner := "0x4040404040404040404040404040404040404040"
	for i := 0; i < 3; i++ {
		if _, err := m.RecordWrite(owner, "b", "o", 10); err != nil {
			t.Fatalf("RecordWrite: %v", err)
		}
	}
	s := interruptSettlement(t, m, owner, "0xdeadbeef")

	m2 := NewManager(m.db, m.cfg)

	var settled int64
	m2.db.Model(&types.MeterEvent{}).Where("owner = ? AND status = ? AND settlement_id = ?", owner, eventSettled, s.ID).Count(&settled)
	if settled != 3 {
		t.Errorf("%d events settled under recovered settlement, want 3", settled)
	}
	var srec types.MeterSettlement
	m2.db.First(&srec, s.ID)
	if srec.Status != settlementConfirmed {
		t.Errorf("settlement status = %s, want confirmed", srec.Status)
	}
	u, _ := m2.GetUsage(owner)
	if u.UnsettledFeeWei != "0" {
		t.Errorf("unsettled fee = %s, want 0 (user already paid)", u.UnsettledFeeWei)
	}
}

// Events reserved by another (interrupted) settlement with lower ids must not
// be swept into a later settlement's confirmed range.
func TestSettleERC8183DoesNotSweepForeignSettlingEvents(t *testing.T) {
	txh := common.HexToHash("0xabc123")
	m := newERC8183TestManager(t,
		&fakeERC20{transferHash: txh},
		&fakeERC8183{jobID: big.NewInt(9), txHash: txh},
	)
	owner := "0x5050505050505050505050505050505050505050"

	// Orphan: one event stuck in "settling" from a prior interrupted attempt.
	if _, err := m.RecordWrite(owner, "b", "orphan", 10); err != nil {
		t.Fatalf("RecordWrite: %v", err)
	}
	if err := m.db.Model(&types.MeterEvent{}).
		Where("owner = ?", owner).Update("status", eventSettling).Error; err != nil {
		t.Fatalf("orphan setup: %v", err)
	}

	// Two fresh unsettled events, then settle.
	for i := 0; i < 2; i++ {
		if _, err := m.RecordWrite(owner, "b", "o", 10); err != nil {
			t.Fatalf("RecordWrite: %v", err)
		}
	}
	resp, err := m.Settle(owner)
	if err != nil {
		t.Fatalf("Settle: %v", err)
	}
	if resp.SettledAmountWei != "2000000000000000000" {
		t.Errorf("settled = %s, want 2e18 (orphan excluded)", resp.SettledAmountWei)
	}

	// The orphan must still be "settling", untouched by this settlement.
	var orphan types.MeterEvent
	m.db.Where("owner = ? AND object_name = ?", owner, "orphan").First(&orphan)
	if orphan.Status != eventSettling || orphan.SettlementID != 0 {
		t.Errorf("orphan = status=%s settlement_id=%d, want settling/0", orphan.Status, orphan.SettlementID)
	}
	// Its fee stays in the account's unsettled total (recovery owns it).
	u, _ := m.GetUsage(owner)
	if u.UnsettledFeeWei != "1000000000000000000" {
		t.Errorf("unsettled fee = %s, want 1e18 (orphan's fee)", u.UnsettledFeeWei)
	}
}

// erc8183 auto-settle requires a positive threshold; offchain does not.
func TestAutoSettleEnabledThresholdGate(t *testing.T) {
	cases := []struct {
		name      string
		mode      string
		threshold *big.Int
		want      bool
	}{
		{"erc8183 zero threshold", modeERC8183, big.NewInt(0), false},
		{"erc8183 positive threshold", modeERC8183, big.NewInt(1), true},
		{"offchain zero threshold", modeOffchain, big.NewInt(0), true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := newTestManager(t, Config{
				Enabled:            true,
				AutoSettle:         true,
				SettlementMode:     c.mode,
				SettleThresholdWei: c.threshold,
			})
			if got := m.AutoSettleEnabled(); got != c.want {
				t.Errorf("AutoSettleEnabled = %v, want %v", got, c.want)
			}
		})
	}
}
