package hub

import (
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// newKindTestServer stands up an in-memory sqlite with the bucket/needle/account
// schema so addBucket/listBucket kind logic can be exercised without chain/logfs.
func newKindTestServer(t *testing.T) *Server {
	t.Helper()
	dsn := "file:" + t.Name() + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1)
	if err := db.AutoMigrate(&types.Bucket{}, &types.Needle{}, &types.Account{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return &Server{
		gdb:           db,
		memStat:       &memStatCache{},
		bucketDisplay: make(map[string]types.BucketDisplay),
	}
}

// TestBucketKind covers the A-tier object-store kind model: addBucket records the
// scenario kind, listBucket filters by it, and the "memory" filter also matches
// legacy empty-kind rows (so pre-migration buckets keep showing under memory).
func TestBucketKind(t *testing.T) {
	s := newKindTestServer(t)
	const owner = "0xabc0000000000000000000000000000000000001"

	if err := s.addBucket(owner, "mem-a", "memory"); err != nil {
		t.Fatalf("addBucket memory: %v", err)
	}
	if err := s.addBucket(owner, "model-a", "model"); err != nil {
		t.Fatalf("addBucket model: %v", err)
	}
	if err := s.addBucket(owner, "ds-a", "dataset"); err != nil {
		t.Fatalf("addBucket dataset: %v", err)
	}
	// empty kind = defaults to memory (addBucket normalizes "" -> "memory")
	if err := s.addBucket(owner, "mem-default", ""); err != nil {
		t.Fatalf("addBucket default: %v", err)
	}
	// a legacy row written directly with no kind (simulates pre-migration data)
	if err := s.gdb.Create(&types.Bucket{Name: "legacy", Owner: owner}).Error; err != nil {
		t.Fatalf("seed legacy: %v", err)
	}

	count := func(kind string) int {
		res, err := s.listBucket(owner, kind, 0, 100)
		if err != nil {
			t.Fatalf("listBucket %q: %v", kind, err)
		}
		return len(res)
	}

	// memory filter: mem-a + mem-default + legacy(empty) = 3
	if got := count("memory"); got != 3 {
		t.Errorf("kind=memory: got %d buckets, want 3 (incl. default + legacy empty-kind)", got)
	}
	if got := count("model"); got != 1 {
		t.Errorf("kind=model: got %d, want 1", got)
	}
	if got := count("dataset"); got != 1 {
		t.Errorf("kind=dataset: got %d, want 1", got)
	}
	// no filter: all 5
	if got := count(""); got != 5 {
		t.Errorf("kind='': got %d, want 5 (all)", got)
	}
}

// TestBucketKindBackfill: uploading model content to a bucket that exists with an
// empty kind backfills its kind (idempotent; memory never gets overwritten).
func TestBucketKindBackfill(t *testing.T) {
	s := newKindTestServer(t)
	const owner = "0xabc0000000000000000000000000000000000002"

	// legacy empty-kind bucket
	if err := s.gdb.Create(&types.Bucket{Name: "repo1", Owner: owner}).Error; err != nil {
		t.Fatalf("seed: %v", err)
	}
	// second touch as a model → backfills kind=model on the empty-kind row
	if err := s.addBucket(owner, "repo1", "model"); err != nil {
		t.Fatalf("addBucket backfill: %v", err)
	}
	var b types.Bucket
	if err := s.gdb.First(&b, "name = ?", "repo1").Error; err != nil {
		t.Fatalf("reload: %v", err)
	}
	if b.Kind != "model" {
		t.Errorf("kind after backfill: got %q, want model", b.Kind)
	}
	// a memory touch must NOT overwrite an existing non-empty kind
	if err := s.addBucket(owner, "repo1", "memory"); err != nil {
		t.Fatalf("addBucket memory touch: %v", err)
	}
	s.gdb.First(&b, "name = ?", "repo1")
	if b.Kind != "model" {
		t.Errorf("kind clobbered by memory touch: got %q, want model", b.Kind)
	}
}
