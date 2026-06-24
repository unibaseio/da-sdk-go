package hub

import (
	"testing"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newStatTestServer(t *testing.T) *Server {
	t.Helper()
	// Unique, isolated in-memory DB per test (name = test name). cache=shared
	// keeps it alive across the pooled connection; MaxOpenConns(1) avoids the
	// pool dropping the only connection (which would discard the :memory: db).
	dsn := "file:" + t.Name() + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("db handle: %v", err)
	}
	sqlDB.SetMaxOpenConns(1)
	if err := db.AutoMigrate(&types.Needle{}, &types.Account{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return &Server{gdb: db, memStat: &memStatCache{}}
}

func seedNeedle(t *testing.T, s *Server, owner string, size uint64) {
	t.Helper()
	if err := s.gdb.Create(&types.Needle{Owner: owner, Size: size}).Error; err != nil {
		t.Fatalf("seed: %v", err)
	}
}

func seedAccount(t *testing.T, s *Server, name string) {
	t.Helper()
	if err := s.gdb.Create(&types.Account{Name: name}).Error; err != nil {
		t.Fatalf("seed account: %v", err)
	}
}

func TestComputeMemStats_AggregatesPerOwner(t *testing.T) {
	s := newStatTestServer(t)
	// owner A: 3 entries, 100+200+300 = 600 bytes (mixed + lower case → merge)
	seedNeedle(t, s, "0xAAAa000000000000000000000000000000000001", 100)
	seedNeedle(t, s, "0xAAAa000000000000000000000000000000000001", 200)
	seedNeedle(t, s, "0xaaaa000000000000000000000000000000000001", 300)
	// owner B: 1 entry, 50 bytes
	seedNeedle(t, s, "0xBBBB000000000000000000000000000000000002", 50)
	// 3 accounts on the hub (one has no memory)
	seedAccount(t, s, "0xaaaa000000000000000000000000000000000001")
	seedAccount(t, s, "0xbbbb000000000000000000000000000000000002")
	seedAccount(t, s, "0xcccc000000000000000000000000000000000003")

	ov, owners, err := s.computeMemStats()
	if err != nil {
		t.Fatalf("computeMemStats: %v", err)
	}

	if ov.TotalAddresses != 3 {
		t.Errorf("TotalAddresses: want 3, got %d", ov.TotalAddresses)
	}
	if ov.WalletsWithMemory != 2 {
		t.Errorf("WalletsWithMemory: want 2 (case-merged), got %d", ov.WalletsWithMemory)
	}
	if ov.MemoryCount != 4 {
		t.Errorf("MemoryCount: want 4, got %d", ov.MemoryCount)
	}
	if ov.MemoryBytes != 650 {
		t.Errorf("MemoryBytes: want 650, got %d", ov.MemoryBytes)
	}
	if len(owners) != 2 {
		t.Fatalf("want 2 owners, got %d", len(owners))
	}
	// sorted by bytes desc → A (600) first
	if owners[0].Owner != "0xaaaa000000000000000000000000000000000001" ||
		owners[0].Count != 3 || owners[0].Bytes != 600 {
		t.Errorf("owner A wrong: %+v", owners[0])
	}
	if owners[1].Count != 1 || owners[1].Bytes != 50 {
		t.Errorf("owner B wrong: %+v", owners[1])
	}
}

func TestComputeMemStats_Empty(t *testing.T) {
	s := newStatTestServer(t)
	ov, owners, err := s.computeMemStats()
	if err != nil {
		t.Fatal(err)
	}
	if ov.TotalAddresses != 0 || ov.WalletsWithMemory != 0 || ov.MemoryCount != 0 || ov.MemoryBytes != 0 {
		t.Fatalf("want all-zero overview, got %+v", ov)
	}
	if len(owners) != 0 {
		t.Fatalf("want 0 owners, got %d", len(owners))
	}
}

func TestMemStatCache_Pagination(t *testing.T) {
	s := newStatTestServer(t)
	// build a snapshot of 5 owners directly (bypass DB; test the in-memory page)
	owners := []types.MemoryStat{
		{Owner: "0x5", Count: 5, Bytes: 500},
		{Owner: "0x4", Count: 4, Bytes: 400},
		{Owner: "0x3", Count: 3, Bytes: 300},
		{Owner: "0x2", Count: 2, Bytes: 200},
		{Owner: "0x1", Count: 1, Bytes: 100},
	}
	s.memStat.set(&memStatSnapshot{
		overview:   types.MemoryOverview{WalletsWithMemory: 5},
		owners:     owners,
		computedAt: time.Unix(1700000000, 0),
	})

	p1 := s.memoryStatPage("", 0, 2)
	if p1.Total != 5 || len(p1.Items) != 2 || p1.Items[0].Bytes != 500 || p1.Items[1].Bytes != 400 {
		t.Fatalf("page1 wrong: %+v", p1)
	}
	if p1.ComputedAt != 1700000000 {
		t.Errorf("want ComputedAt set, got %d", p1.ComputedAt)
	}

	p3 := s.memoryStatPage("", 4, 2)
	if len(p3.Items) != 1 || p3.Items[0].Bytes != 100 {
		t.Fatalf("page3 wrong: %+v", p3)
	}

	// offset past the end → empty page, but total still correct
	pEnd := s.memoryStatPage("", 10, 2)
	if pEnd.Total != 5 || len(pEnd.Items) != 0 {
		t.Fatalf("pastEnd wrong: %+v", pEnd)
	}
}

func TestMemStatCache_OwnerFilter(t *testing.T) {
	s := newStatTestServer(t)
	s.memStat.set(&memStatSnapshot{
		owners: []types.MemoryStat{
			{Owner: "0xaaaa000000000000000000000000000000000001", Count: 5, Bytes: 500},
			{Owner: "0xbbbb000000000000000000000000000000000002", Count: 2, Bytes: 200},
		},
		computedAt: time.Unix(1700000000, 0),
	})

	// exact owner → 1-item list
	hit := s.memoryStatPage("0xaaaa000000000000000000000000000000000001", 0, 32)
	if hit.Total != 1 || len(hit.Items) != 1 || hit.Items[0].Count != 5 {
		t.Fatalf("owner hit wrong: %+v", hit)
	}

	// mixed-case query still matches (snapshot is lowercased)
	mixed := s.memoryStatPage("0xAAAA000000000000000000000000000000000001", 0, 32)
	if mixed.Total != 1 || len(mixed.Items) != 1 {
		t.Fatalf("mixed-case owner should match: %+v", mixed)
	}

	// unknown owner → empty list, total 0
	miss := s.memoryStatPage("0xdddd000000000000000000000000000000000009", 0, 32)
	if miss.Total != 0 || len(miss.Items) != 0 {
		t.Fatalf("unknown owner should be empty: %+v", miss)
	}
}

func TestMemStatCache_NotReady(t *testing.T) {
	s := newStatTestServer(t) // memStat set but no snapshot computed yet
	ov := s.memoryOverviewSnapshot()
	if ov.ComputedAt != 0 || ov.MemoryCount != 0 {
		t.Fatalf("want empty overview before first compute, got %+v", ov)
	}
	page := s.memoryStatPage("", 0, 10)
	if page.Total != 0 || len(page.Items) != 0 || page.ComputedAt != 0 {
		t.Fatalf("want empty page before first compute, got %+v", page)
	}
}

func TestRefreshMemStats_PopulatesCache(t *testing.T) {
	s := newStatTestServer(t)
	seedNeedle(t, s, "0xAbc0000000000000000000000000000000000001", 1000)
	seedAccount(t, s, "0xabc0000000000000000000000000000000000001")

	s.refreshMemStats()

	ov := s.memoryOverviewSnapshot()
	if ov.ComputedAt == 0 {
		t.Fatal("ComputedAt should be set after refresh")
	}
	if ov.WalletsWithMemory != 1 || ov.MemoryCount != 1 || ov.MemoryBytes != 1000 {
		t.Fatalf("overview wrong after refresh: %+v", ov)
	}
	page := s.memoryStatPage("", 0, 10)
	if page.Total != 1 || len(page.Items) != 1 || page.Items[0].Bytes != 1000 {
		t.Fatalf("page wrong after refresh: %+v", page)
	}
}
