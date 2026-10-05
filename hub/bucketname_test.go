package hub

import (
	"net/http"
	"strings"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

// Deleting a bucket removes only the signer's rows: legacy data can hold
// another owner's bucket row or objects under the same name.
func TestV1DeleteBucketScopedToOwner(t *testing.T) {
	s := newV1TestServer(t)
	alice, ak := testKey(t)
	bob := "0x00000000000000000000000000000000000000b2"
	la := strings.ToLower(alice)

	s.gdb.Create(&types.Bucket{Name: "shared", Owner: la, Kind: "memory"})  // oldest: alice owns it
	s.gdb.Create(&types.Bucket{Name: "shared", Owner: bob, Kind: "memory"}) // legacy duplicate
	s.gdb.Create(&types.Needle{Owner: la, Bucket: "shared", Name: "a"})
	s.gdb.Create(&types.Needle{Owner: bob, Bucket: "shared", Name: "b"})

	if w := do(t, s, "DELETE", "/v1/buckets/shared", authHeader(alice, ak), ""); w.Code != http.StatusOK {
		t.Fatalf("delete: %d %s", w.Code, w.Body.String())
	}
	var nb, nn int64
	s.gdb.Model(&types.Bucket{}).Where("owner = ?", bob).Count(&nb)
	s.gdb.Model(&types.Needle{}).Where("owner = ?", bob).Count(&nn)
	if nb != 1 || nn != 1 {
		t.Fatalf("bob's rows deleted: buckets=%d needles=%d", nb, nn)
	}
	s.gdb.Model(&types.Needle{}).Where("owner = ?", la).Count(&nn)
	if nn != 0 {
		t.Fatalf("alice's needles left: %d", nn)
	}
}

func TestBucketNameUniqueIndex(t *testing.T) {
	s := newV1TestServer(t)
	ensureBucketNameUnique(s.gdb)
	if err := s.addBucket("0xa1", "n1", "memory"); err != nil {
		t.Fatal(err)
	}
	// a racing second owner's insert is refused by the DB...
	if err := s.gdb.Create(&types.Bucket{Name: "n1", Owner: "0xb2"}).Error; err == nil {
		t.Fatal("duplicate live bucket name accepted")
	}
	// ...and addBucket reports the name as taken
	if err := s.addBucket("0xb2", "n1", "memory"); err == nil {
		t.Fatal("second owner got the bucket")
	}
	if err := s.addBucket("0xA1", "n1", "memory"); err != nil {
		t.Fatalf("owner re-add: %v", err)
	}
	// a deleted bucket's name is free again
	s.gdb.Where("name = ?", "n1").Delete(&types.Bucket{})
	if err := s.addBucket("0xb2", "n1", "memory"); err != nil {
		t.Fatalf("reuse after delete: %v", err)
	}
}

// Existing duplicates must not break startup: the index is skipped, not forced.
func TestBucketNameUniqueSkippedOnDuplicates(t *testing.T) {
	s := newV1TestServer(t)
	s.gdb.Create(&types.Bucket{Name: "dup", Owner: "0xa1"})
	s.gdb.Create(&types.Bucket{Name: "dup", Owner: "0xb2"})
	ensureBucketNameUnique(s.gdb)
	var n int64
	s.gdb.Raw("SELECT COUNT(*) FROM sqlite_master WHERE type='index' AND name='idx_buckets_name_live_unique'").Scan(&n)
	if n != 0 {
		t.Fatal("unique index created over duplicates")
	}
	if err := s.addBucket("0xb2", "dup", "memory"); err == nil {
		t.Fatal("non-first owner allowed")
	}
}
