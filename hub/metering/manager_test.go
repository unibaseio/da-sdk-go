package metering

import (
	"math/big"
	"path/filepath"
	"sync"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

// newTestDB opens a fresh on-disk sqlite DB (file, not :memory:, so concurrent
// connections share state) with the metering tables migrated.
func newTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.db")
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
	})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	db.Exec("PRAGMA busy_timeout = 60000;")
	if err := db.AutoMigrate(&types.MeterAccount{}, &types.MeterEvent{}, &types.MeterSettlement{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func newTestManager(t *testing.T, cfg Config) *Manager {
	t.Helper()
	if cfg.WriteBaseWei == nil {
		cfg.WriteBaseWei = big.NewInt(0)
	}
	if cfg.WritePerKBWei == nil {
		cfg.WritePerKBWei = big.NewInt(0)
	}
	if cfg.ReadPerRequestWei == nil {
		cfg.ReadPerRequestWei = big.NewInt(0)
	}
	if cfg.DefaultCreditLimitWei == nil {
		cfg.DefaultCreditLimitWei = big.NewInt(0)
	}
	return NewManager(newTestDB(t), cfg)
}

func bi(s string) *big.Int {
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("bad bigint: " + s)
	}
	return n
}

func TestPriceWriteRoundsUpToKB(t *testing.T) {
	m := newTestManager(t, Config{
		WriteBaseWei:  big.NewInt(100),
		WritePerKBWei: big.NewInt(10),
	})
	cases := []struct {
		bytes uint64
		want  string
	}{
		{0, "100"},     // empty -> base only
		{1, "110"},     // 1 byte -> 1 KB
		{1024, "110"},  // exactly 1 KB
		{1025, "120"},  // just over -> 2 KB
		{2500, "130"},  // ceil(2500/1024)=3
		{10240, "200"}, // exactly 10 KB
	}
	for _, c := range cases {
		got := m.PriceWrite(c.bytes).String()
		if got != c.want {
			t.Errorf("PriceWrite(%d) = %s, want %s", c.bytes, got, c.want)
		}
	}
}

func TestPriceWriteBigIntNoOverflow(t *testing.T) {
	// Values well beyond uint64 range (max ~1.8e19). 1e18 base, 2e18 per KB.
	m := newTestManager(t, Config{
		WriteBaseWei:  bi("1000000000000000000"),
		WritePerKBWei: bi("2000000000000000000"),
	})
	// 2500 bytes -> 3 KB -> 1e18 + 3*2e18 = 7e18
	got := m.PriceWrite(2500).String()
	if got != "7000000000000000000" {
		t.Fatalf("PriceWrite = %s, want 7000000000000000000", got)
	}
}

func TestRecordWriteIncrementsAccount(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:       true,
		WriteBaseWei:  big.NewInt(100),
		WritePerKBWei: big.NewInt(10),
	})
	owner := "0xABCdef0000000000000000000000000000000001"

	ev, err := m.RecordWrite(owner, "bucket1", "obj1", 2500)
	if err != nil {
		t.Fatalf("RecordWrite: %v", err)
	}
	if ev.FeeWei != "130" {
		t.Errorf("event fee = %s, want 130", ev.FeeWei)
	}
	if ev.Status != eventUnsettled {
		t.Errorf("event status = %s, want unsettled", ev.Status)
	}

	u, err := m.GetUsage(owner)
	if err != nil {
		t.Fatalf("GetUsage: %v", err)
	}
	if u.TotalWrites != 1 || u.TotalBytesWritten != 2500 {
		t.Errorf("writes=%d bytes=%d, want 1/2500", u.TotalWrites, u.TotalBytesWritten)
	}
	if u.TotalFeeWei != "130" || u.UnsettledFeeWei != "130" {
		t.Errorf("totalFee=%s unsettled=%s, want 130/130", u.TotalFeeWei, u.UnsettledFeeWei)
	}
	// owner must be canonicalized to lowercase
	if u.Owner != "0xabcdef0000000000000000000000000000000001" {
		t.Errorf("owner not lowercased: %s", u.Owner)
	}
}

func TestRecordWriteAccumulates(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:       true,
		WriteBaseWei:  bi("1000000000000000000"),
		WritePerKBWei: big.NewInt(0),
	})
	owner := "0x1111111111111111111111111111111111111111"
	for i := 0; i < 3; i++ {
		if _, err := m.RecordWrite(owner, "b", "o", 10); err != nil {
			t.Fatalf("RecordWrite: %v", err)
		}
	}
	u, _ := m.GetUsage(owner)
	if u.TotalWrites != 3 {
		t.Errorf("writes = %d, want 3", u.TotalWrites)
	}
	if u.UnsettledFeeWei != "3000000000000000000" {
		t.Errorf("unsettled = %s, want 3e18", u.UnsettledFeeWei)
	}
}

func TestConcurrentRecordWriteNoLostIncrements(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:       true,
		WriteBaseWei:  bi("1000000000000000000"), // 1e18 per write
		WritePerKBWei: big.NewInt(0),
	})
	owner := "0x2222222222222222222222222222222222222222"

	const n = 100
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			if _, err := m.RecordWrite(owner, "b", "o", 100); err != nil {
				t.Errorf("RecordWrite: %v", err)
			}
		}()
	}
	wg.Wait()

	u, _ := m.GetUsage(owner)
	if u.TotalWrites != n {
		t.Errorf("writes = %d, want %d", u.TotalWrites, n)
	}
	// 100 * 1e18 = 1e20 (exceeds uint64) — verifies big.Int accumulation.
	if u.UnsettledFeeWei != "100000000000000000000" {
		t.Errorf("unsettled = %s, want 1e20", u.UnsettledFeeWei)
	}
	// event count sanity
	var cnt int64
	m.db.Model(&types.MeterEvent{}).Where("owner = ?", owner).Count(&cnt)
	if cnt != n {
		t.Errorf("event count = %d, want %d", cnt, n)
	}
}

func TestGetUsageMissingAccount(t *testing.T) {
	m := newTestManager(t, Config{Enabled: true, DefaultCreditLimitWei: bi("500")})
	u, err := m.GetUsage("0x3333333333333333333333333333333333333333")
	if err != nil {
		t.Fatalf("GetUsage: %v", err)
	}
	if u.TotalWrites != 0 || u.UnsettledFeeWei != "0" || u.TotalFeeWei != "0" {
		t.Errorf("expected zeroed usage, got %+v", u)
	}
	if u.CreditLimitWei != "500" {
		t.Errorf("credit limit = %s, want 500 (default)", u.CreditLimitWei)
	}
}

func TestRecordReadIncrementsReads(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:           true,
		ReadPerRequestWei: big.NewInt(5),
	})
	owner := "0x4444444444444444444444444444444444444444"
	if _, err := m.RecordRead(owner, "b", "o", 42); err != nil {
		t.Fatalf("RecordRead: %v", err)
	}
	u, _ := m.GetUsage(owner)
	if u.TotalReads != 1 || u.TotalWrites != 0 {
		t.Errorf("reads=%d writes=%d, want 1/0", u.TotalReads, u.TotalWrites)
	}
	if u.UnsettledFeeWei != "5" {
		t.Errorf("unsettled = %s, want 5", u.UnsettledFeeWei)
	}
}
