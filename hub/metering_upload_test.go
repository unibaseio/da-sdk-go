package hub

import (
	"math/big"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/unibaseio/da-sdk-go/hub/metering"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func newMeteringTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.db")
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
	})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&types.MeterAccount{}, &types.MeterEvent{}, &types.MeterSettlement{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func meterEventCount(t *testing.T, db *gorm.DB, owner string) int64 {
	t.Helper()
	var n int64
	db.Model(&types.MeterEvent{}).Where("owner = ?", owner).Count(&n)
	return n
}

// recordUploadWrite must be a no-op when metering is disabled.
func TestRecordUploadWriteDisabledNoEvent(t *testing.T) {
	db := newMeteringTestDB(t)
	s := &Server{metering: metering.NewManager(db, metering.Config{
		Enabled:       false, // disabled
		ChargeWrites:  true,
		WriteBaseWei:  big.NewInt(100),
		WritePerKBWei: big.NewInt(10),
	})}

	owner := "0xabc0000000000000000000000000000000000001"
	s.recordUploadWrite(owner, "b", "o", 2500)

	if n := meterEventCount(t, db, owner); n != 0 {
		t.Fatalf("disabled metering created %d events, want 0", n)
	}
}

// recordUploadWrite must be a no-op when write charging is off, even if enabled.
func TestRecordUploadWriteChargeWritesOffNoEvent(t *testing.T) {
	db := newMeteringTestDB(t)
	s := &Server{metering: metering.NewManager(db, metering.Config{
		Enabled:      true,
		ChargeWrites: false, // charging off
		WriteBaseWei: big.NewInt(100),
	})}

	owner := "0xabc0000000000000000000000000000000000002"
	s.recordUploadWrite(owner, "b", "o", 2500)

	if n := meterEventCount(t, db, owner); n != 0 {
		t.Fatalf("charge-writes-off created %d events, want 0", n)
	}
}

// A successful upload path records exactly one event with the correct fee.
func TestRecordUploadWriteEnabledCreatesEvent(t *testing.T) {
	db := newMeteringTestDB(t)
	s := &Server{metering: metering.NewManager(db, metering.Config{
		Enabled:       true,
		ChargeWrites:  true,
		WriteBaseWei:  big.NewInt(100),
		WritePerKBWei: big.NewInt(10),
	})}

	owner := "0xABC0000000000000000000000000000000000003"
	s.recordUploadWrite(owner, "bucket", "object", 2500)

	// owner is canonicalized to lowercase in the ledger
	lc := "0xabc0000000000000000000000000000000000003"
	if n := meterEventCount(t, db, lc); n != 1 {
		t.Fatalf("enabled metering created %d events, want 1", n)
	}

	var ev types.MeterEvent
	db.Where("owner = ?", lc).First(&ev)
	if ev.FeeWei != "130" { // 100 + ceil(2500/1024=3)*10
		t.Errorf("event fee = %s, want 130", ev.FeeWei)
	}
	if ev.Bytes != 2500 || ev.Operation != "write" {
		t.Errorf("event bytes=%d op=%s, want 2500/write", ev.Bytes, ev.Operation)
	}
}

// checkWriteAdmission allows when metering is disabled (no 402, no body).
func TestCheckWriteAdmissionDisabledAllows(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := newMeteringTestDB(t)
	s := &Server{metering: metering.NewManager(db, metering.Config{
		Enabled:               false,
		WriteBaseWei:          big.NewInt(100),
		DefaultCreditLimitWei: big.NewInt(1),
	})}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	if !s.checkWriteAdmission(c, "0xabc0000000000000000000000000000000000010", 100000) {
		t.Fatal("disabled metering should allow the write")
	}
}

// checkWriteAdmission returns 402 when the write exceeds the credit limit.
func TestCheckWriteAdmissionCreditLimit402(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := newMeteringTestDB(t)
	s := &Server{metering: metering.NewManager(db, metering.Config{
		Enabled:               true,
		ChargeWrites:          true,
		WriteBaseWei:          big.NewInt(100),
		DefaultCreditLimitWei: big.NewInt(150),
	})}
	owner := "0xabc0000000000000000000000000000000000011"

	// First write accrues 100 unsettled (allowed).
	w1 := httptest.NewRecorder()
	c1, _ := gin.CreateTestContext(w1)
	if !s.checkWriteAdmission(c1, owner, 1) {
		t.Fatal("first write should be allowed")
	}
	s.recordUploadWrite(owner, "b", "o", 1)

	// Second write: required 200 > 150 -> 402.
	w2 := httptest.NewRecorder()
	c2, _ := gin.CreateTestContext(w2)
	if s.checkWriteAdmission(c2, owner, 1) {
		t.Fatal("second write should be rejected")
	}
	if w2.Code != http.StatusPaymentRequired {
		t.Fatalf("status = %d, want 402", w2.Code)
	}
}
