package metering

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// fakeERC8183 records call order and can be made to fail at a chosen step.
type fakeERC8183 struct {
	jobID   *big.Int
	failAt  string // "", "createJob", "setBudget", "fund", "submit"
	txHash  common.Hash
	callLog []string
}

func (f *fakeERC8183) CreateJob(ctx context.Context, expiredAt *big.Int, description string, providerAgentID *big.Int) (*big.Int, common.Hash, error) {
	f.callLog = append(f.callLog, "createJob")
	if f.failAt == "createJob" {
		return nil, common.Hash{}, fmt.Errorf("createJob boom")
	}
	return f.jobID, f.txHash, nil
}

func (f *fakeERC8183) SetBudget(ctx context.Context, jobID, amount *big.Int) (common.Hash, error) {
	f.callLog = append(f.callLog, "setBudget")
	if f.failAt == "setBudget" {
		return common.Hash{}, fmt.Errorf("setBudget boom")
	}
	return f.txHash, nil
}

func (f *fakeERC8183) Fund(ctx context.Context, jobID, expectedBudget *big.Int) (common.Hash, error) {
	f.callLog = append(f.callLog, "fund")
	if f.failAt == "fund" {
		return common.Hash{}, fmt.Errorf("fund boom")
	}
	return f.txHash, nil
}

func (f *fakeERC8183) Submit(ctx context.Context, jobID *big.Int, deliverable [32]byte) (common.Hash, error) {
	f.callLog = append(f.callLog, "submit")
	if f.failAt == "submit" {
		return common.Hash{}, fmt.Errorf("submit boom")
	}
	return f.txHash, nil
}

func newERC8183TestManager(t *testing.T, erc20 erc20API, erc8183 erc8183API) *Manager {
	m := newTestManager(t, Config{
		Enabled:        true,
		ChargeWrites:   true,
		WriteBaseWei:   bi("1000000000000000000"), // 1e18/write
		SettlementMode: modeERC8183,
	})
	m.erc20 = erc20
	m.erc8183 = erc8183
	m.provider = common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	return m
}

func TestSettleERC8183SuccessClearsDebt(t *testing.T) {
	txh := common.HexToHash("0xabc123")
	m := newERC8183TestManager(t,
		&fakeERC20{transferHash: txh},
		&fakeERC8183{jobID: big.NewInt(77), txHash: txh},
	)
	owner := "0x1010101010101010101010101010101010101010"
	for i := 0; i < 2; i++ {
		if _, err := m.RecordWrite(owner, "b", "o", 10); err != nil {
			t.Fatalf("RecordWrite: %v", err)
		}
	}

	resp, err := m.Settle(owner) // mode=erc8183 -> SettleERC8183
	if err != nil {
		t.Fatalf("Settle: %v", err)
	}
	if resp.Status != settlementConfirmed {
		t.Fatalf("status = %s, want confirmed", resp.Status)
	}
	if resp.SettledAmountWei != "2000000000000000000" {
		t.Errorf("settled = %s, want 2e18", resp.SettledAmountWei)
	}
	if resp.JobID != "77" {
		t.Errorf("job id = %s, want 77", resp.JobID)
	}
	if resp.Tx == nil || resp.Tx.Transfer == "" || resp.Tx.CreateJob == "" || resp.Tx.Fund == "" || resp.Tx.Submit == "" {
		t.Errorf("tx hashes not all populated: %+v", resp.Tx)
	}
	if resp.ReportHash == "" {
		t.Error("report hash not set")
	}

	// debt cleared
	u, _ := m.GetUsage(owner)
	if u.UnsettledFeeWei != "0" {
		t.Errorf("unsettled = %s, want 0", u.UnsettledFeeWei)
	}

	// events settled; settlement record persisted with tx hashes + job id
	var unsettled int64
	m.db.Model(&types.MeterEvent{}).Where("owner = ? AND status != ?", owner, eventSettled).Count(&unsettled)
	if unsettled != 0 {
		t.Errorf("%d events not settled", unsettled)
	}
	var srec types.MeterSettlement
	m.db.First(&srec, resp.SettlementID)
	if srec.Status != settlementConfirmed || srec.JobID != "77" || srec.TransferTx == "" || srec.SubmitTx == "" {
		t.Errorf("settlement record = %+v", srec)
	}
}

func TestSettleERC8183ChainFailureKeepsDebt(t *testing.T) {
	txh := common.HexToHash("0xabc123")
	fake8183 := &fakeERC8183{jobID: big.NewInt(88), txHash: txh, failAt: "fund"}
	m := newERC8183TestManager(t, &fakeERC20{transferHash: txh}, fake8183)

	owner := "0x2020202020202020202020202020202020202020"
	for i := 0; i < 3; i++ {
		if _, err := m.RecordWrite(owner, "b", "o", 10); err != nil {
			t.Fatalf("RecordWrite: %v", err)
		}
	}

	resp, err := m.Settle(owner)
	if err == nil {
		t.Fatal("expected settlement error on chain failure")
	}
	if resp == nil || resp.Status != settlementFailed {
		t.Fatalf("resp status = %v, want failed", resp)
	}

	// CRITICAL: debt must be intact.
	u, _ := m.GetUsage(owner)
	if u.UnsettledFeeWei != "3000000000000000000" {
		t.Fatalf("debt was cleared on failure! unsettled = %s, want 3e18", u.UnsettledFeeWei)
	}

	// events must be back to unsettled (retryable), none settled
	var settled int64
	m.db.Model(&types.MeterEvent{}).Where("owner = ? AND status = ?", owner, eventSettled).Count(&settled)
	if settled != 0 {
		t.Errorf("%d events marked settled despite chain failure", settled)
	}
	var unsettled int64
	m.db.Model(&types.MeterEvent{}).Where("owner = ? AND status = ?", owner, eventUnsettled).Count(&unsettled)
	if unsettled != 3 {
		t.Errorf("%d events unsettled, want 3 (retryable)", unsettled)
	}

	// settlement record marked failed, with the tx hashes produced before the failure
	var srec types.MeterSettlement
	m.db.First(&srec, "owner = ?", owner)
	if srec.Status != settlementFailed || srec.Error == "" {
		t.Errorf("settlement record = %+v, want failed with error", srec)
	}
	if srec.TransferTx == "" || srec.CreateJobTx == "" {
		t.Errorf("pre-failure tx hashes not saved: %+v", srec)
	}
	if srec.FundTx != "" || srec.SubmitTx != "" {
		t.Errorf("post-failure tx hashes should be empty: %+v", srec)
	}

	// a retry can now succeed and clear the debt
	fake8183.failAt = ""
	resp2, err := m.Settle(owner)
	if err != nil {
		t.Fatalf("retry Settle: %v", err)
	}
	if resp2.Status != settlementConfirmed {
		t.Fatalf("retry status = %s, want confirmed", resp2.Status)
	}
	u2, _ := m.GetUsage(owner)
	if u2.UnsettledFeeWei != "0" {
		t.Errorf("debt not cleared on retry: %s", u2.UnsettledFeeWei)
	}
}
