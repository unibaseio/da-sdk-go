package kv

import (
	"fmt"
	"testing"
)

// TestKVStatsCountsPuts: every Put is counted, slow-put fires when the threshold
// is ~0, and L0 table count is non-negative and callable on an open store.
func TestKVStatsCountsPuts(t *testing.T) {
	dir := t.TempDir()
	ds, err := NewBadgerStore(dir, &DefaultOptions)
	if err != nil {
		t.Fatalf("badger: %v", err)
	}
	defer ds.Close()

	// force every write to count as "slow" so the counter is deterministic
	ds.slowPutNanos = 1

	const n = 50
	for i := 0; i < n; i++ {
		if err := ds.Put([]byte(fmt.Sprintf("k-%03d", i)), []byte("v")); err != nil {
			t.Fatalf("put %d: %v", i, err)
		}
	}

	puts, slow, maxMs, l0 := ds.KVStats()
	if puts != n {
		t.Fatalf("puts = %d, want %d", puts, n)
	}
	if slow != n {
		t.Fatalf("slow_puts = %d, want %d (threshold=1ns)", slow, n)
	}
	if maxMs < 0 {
		t.Fatalf("max_put_ms negative: %d", maxMs)
	}
	if l0 < 0 {
		t.Fatalf("l0_tables negative: %d", l0)
	}
}

// TestKVStatsNoSlowByDefault: with the default 100ms threshold, small writes
// don't register as slow.
func TestKVStatsNoSlowByDefault(t *testing.T) {
	dir := t.TempDir()
	ds, err := NewBadgerStore(dir, &DefaultOptions)
	if err != nil {
		t.Fatalf("badger: %v", err)
	}
	defer ds.Close()

	for i := 0; i < 20; i++ {
		if err := ds.Put([]byte(fmt.Sprintf("q-%03d", i)), []byte("v")); err != nil {
			t.Fatal(err)
		}
	}
	puts, slow, _, _ := ds.KVStats()
	if puts != 20 {
		t.Fatalf("puts = %d, want 20", puts)
	}
	if slow != 0 {
		t.Fatalf("slow_puts = %d, want 0 for fast small writes", slow)
	}
}
