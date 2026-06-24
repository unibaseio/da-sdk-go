package hub

import (
	"testing"

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
	if err := db.AutoMigrate(&types.Needle{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return &Server{gdb: db}
}

func seedNeedle(t *testing.T, s *Server, owner string, size uint64) {
	t.Helper()
	if err := s.gdb.Create(&types.Needle{Owner: owner, Size: size}).Error; err != nil {
		t.Fatalf("seed: %v", err)
	}
}

func TestListMemoryStat_AggregatesPerOwner(t *testing.T) {
	s := newStatTestServer(t)
	// owner A: 3 entries, 100+200+300 = 600 bytes
	seedNeedle(t, s, "0xAAAa000000000000000000000000000000000001", 100)
	seedNeedle(t, s, "0xAAAa000000000000000000000000000000000001", 200)
	seedNeedle(t, s, "0xaaaa000000000000000000000000000000000001", 300) // same wallet, lowercase
	// owner B: 1 entry, 50 bytes
	seedNeedle(t, s, "0xBBBB000000000000000000000000000000000002", 50)

	res, err := s.listMemoryStat(0, 10)
	if err != nil {
		t.Fatalf("listMemoryStat: %v", err)
	}

	// 2 distinct owners (case-insensitive merge of A's two cases)
	if res.Total != 2 {
		t.Fatalf("want total=2 distinct owners, got %d", res.Total)
	}
	if len(res.Items) != 2 {
		t.Fatalf("want 2 items, got %d", len(res.Items))
	}

	// ordered by bytes desc → A (600) first
	a := res.Items[0]
	if a.Owner != "0xaaaa000000000000000000000000000000000001" {
		t.Errorf("want owner A (lowercased) first, got %s", a.Owner)
	}
	if a.Count != 3 || a.Bytes != 600 {
		t.Errorf("owner A: want count=3 bytes=600, got count=%d bytes=%d", a.Count, a.Bytes)
	}
	if a.GB != 600.0/1e9 {
		t.Errorf("owner A: want gb=%v, got %v", 600.0/1e9, a.GB)
	}

	b := res.Items[1]
	if b.Count != 1 || b.Bytes != 50 {
		t.Errorf("owner B: want count=1 bytes=50, got count=%d bytes=%d", b.Count, b.Bytes)
	}
}

func TestListMemoryStat_Pagination(t *testing.T) {
	s := newStatTestServer(t)
	// 5 owners, sizes 500,400,300,200,100 → deterministic desc order
	owners := []struct {
		addr string
		size uint64
	}{
		{"0x0000000000000000000000000000000000000005", 500},
		{"0x0000000000000000000000000000000000000004", 400},
		{"0x0000000000000000000000000000000000000003", 300},
		{"0x0000000000000000000000000000000000000002", 200},
		{"0x0000000000000000000000000000000000000001", 100},
	}
	for _, o := range owners {
		seedNeedle(t, s, o.addr, o.size)
	}

	page1, err := s.listMemoryStat(0, 2)
	if err != nil {
		t.Fatal(err)
	}
	if page1.Total != 5 {
		t.Fatalf("want total=5, got %d", page1.Total)
	}
	if len(page1.Items) != 2 || page1.Items[0].Bytes != 500 || page1.Items[1].Bytes != 400 {
		t.Fatalf("page1 unexpected: %+v", page1.Items)
	}

	page2, err := s.listMemoryStat(2, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(page2.Items) != 2 || page2.Items[0].Bytes != 300 || page2.Items[1].Bytes != 200 {
		t.Fatalf("page2 unexpected: %+v", page2.Items)
	}

	page3, err := s.listMemoryStat(4, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(page3.Items) != 1 || page3.Items[0].Bytes != 100 {
		t.Fatalf("page3 unexpected: %+v", page3.Items)
	}
}

func TestListMemoryStat_Empty(t *testing.T) {
	s := newStatTestServer(t)
	res, err := s.listMemoryStat(0, 10)
	if err != nil {
		t.Fatal(err)
	}
	if res.Total != 0 || len(res.Items) != 0 {
		t.Fatalf("want empty, got total=%d items=%d", res.Total, len(res.Items))
	}
}
