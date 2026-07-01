package hub

import (
	"math/big"
	"path/filepath"
	"testing"

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
