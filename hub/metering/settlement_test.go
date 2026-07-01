package metering

import (
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

func TestSettleOffchainClearsDebt(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:      true,
		ChargeWrites: true,
		WriteBaseWei: bi("1000000000000000000"), // 1e18 per write
	})
	owner := "0xbbbb0000000000000000000000000000000000bb"

	for i := 0; i < 3; i++ {
		if _, err := m.RecordWrite(owner, "b", "o", 10); err != nil {
			t.Fatalf("RecordWrite: %v", err)
		}
	}

	resp, err := m.SettleOffchain(owner)
	if err != nil {
		t.Fatalf("SettleOffchain: %v", err)
	}
	if resp.Status != settlementConfirmed {
		t.Fatalf("status = %s, want confirmed", resp.Status)
	}
	if resp.SettledAmountWei != "3000000000000000000" {
		t.Errorf("settled = %s, want 3e18", resp.SettledAmountWei)
	}
	if resp.SettlementID == 0 {
		t.Error("settlement id should be non-zero")
	}

	// account debt cleared
	u, _ := m.GetUsage(owner)
	if u.UnsettledFeeWei != "0" {
		t.Errorf("unsettled after settle = %s, want 0", u.UnsettledFeeWei)
	}
	if u.TotalFeeWei != "3000000000000000000" {
		t.Errorf("total fee should be unchanged = %s, want 3e18", u.TotalFeeWei)
	}

	// all events marked settled, pointing at the settlement
	var unsettled int64
	m.db.Model(&types.MeterEvent{}).Where("owner = ? AND status = ?", owner, eventUnsettled).Count(&unsettled)
	if unsettled != 0 {
		t.Errorf("%d events still unsettled, want 0", unsettled)
	}
	var settled int64
	m.db.Model(&types.MeterEvent{}).
		Where("owner = ? AND status = ? AND settlement_id = ?", owner, eventSettled, resp.SettlementID).
		Count(&settled)
	if settled != 3 {
		t.Errorf("%d events settled under id %d, want 3", settled, resp.SettlementID)
	}

	// settlement record exists
	var srec types.MeterSettlement
	if err := m.db.First(&srec, resp.SettlementID).Error; err != nil {
		t.Fatalf("settlement record missing: %v", err)
	}
	if srec.AmountWei != "3000000000000000000" || srec.Mode != modeOffchain {
		t.Errorf("settlement record = %+v", srec)
	}
}

func TestSettleOffchainNoDebtNoop(t *testing.T) {
	m := newTestManager(t, Config{Enabled: true, ChargeWrites: true, WriteBaseWei: big.NewInt(100)})
	resp, err := m.SettleOffchain("0xcccc0000000000000000000000000000000000cc")
	if err != nil {
		t.Fatalf("SettleOffchain: %v", err)
	}
	if resp.Status != settleStatusNoop || resp.SettledAmountWei != "0" {
		t.Fatalf("want noop/0, got %+v", resp)
	}
}

func TestSettleOffchainIdempotent(t *testing.T) {
	m := newTestManager(t, Config{Enabled: true, ChargeWrites: true, WriteBaseWei: big.NewInt(100)})
	owner := "0xdddd0000000000000000000000000000000000dd"
	if _, err := m.RecordWrite(owner, "b", "o", 1); err != nil {
		t.Fatalf("RecordWrite: %v", err)
	}
	if _, err := m.SettleOffchain(owner); err != nil {
		t.Fatalf("first settle: %v", err)
	}
	// second settle: nothing left
	resp, err := m.SettleOffchain(owner)
	if err != nil {
		t.Fatalf("second settle: %v", err)
	}
	if resp.Status != settleStatusNoop {
		t.Fatalf("second settle status = %s, want noop", resp.Status)
	}
}

// --- route-level signer enforcement ---

func fixedSigner(addr string) SignerFunc {
	return func(c *gin.Context) string { return addr }
}

func newRouter(m *Manager, signer SignerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	pub := r.Group("/api")
	authed := r.Group("/api")
	m.RegisterRoutes(pub, authed, signer)
	return r
}

func TestSettleRouteRejectsUnsigned(t *testing.T) {
	m := newTestManager(t, Config{Enabled: true, ChargeWrites: true, WriteBaseWei: big.NewInt(100)})
	r := newRouter(m, fixedSigner("")) // no signer

	req := httptest.NewRequest(http.MethodPost, "/api/metering/settle", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", w.Code)
	}
}

func TestSettleRouteRejectsSignerMismatch(t *testing.T) {
	m := newTestManager(t, Config{Enabled: true, ChargeWrites: true, WriteBaseWei: big.NewInt(100)})
	signer := "0xeeee0000000000000000000000000000000000ee"
	r := newRouter(m, fixedSigner(signer))

	// owner in body differs from signer
	req := httptest.NewRequest(http.MethodPost,
		"/api/metering/settle?owner=0xffff0000000000000000000000000000000000ff", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", w.Code)
	}
}

func TestSettleRouteSucceedsForSigner(t *testing.T) {
	m := newTestManager(t, Config{Enabled: true, ChargeWrites: true, WriteBaseWei: big.NewInt(100)})
	signer := "0xeeee0000000000000000000000000000000000ee"
	if _, err := m.RecordWrite(signer, "b", "o", 1); err != nil {
		t.Fatalf("RecordWrite: %v", err)
	}
	r := newRouter(m, fixedSigner(signer))

	// owner omitted -> defaults to signer
	req := httptest.NewRequest(http.MethodPost, "/api/metering/settle", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body=%s", w.Code, w.Body.String())
	}
	u, _ := m.GetUsage(signer)
	if u.UnsettledFeeWei != "0" {
		t.Errorf("debt not cleared: unsettled = %s", u.UnsettledFeeWei)
	}
}
