package hub

import (
	"strconv"
	"testing"
)

// TestTotalCacheBounded: keys carry client-chosen strings, so a flood of
// distinct ones must not grow the map past its cap — and the cache must still
// serve hits afterwards.
func TestTotalCacheBounded(t *testing.T) {
	tc := newTotalCache()
	one := func() (int64, error) { return 1, nil }
	for i := 0; i < maxTotalCacheEntries+500; i++ {
		tc.get("needles|0xowner|bucket-"+strconv.Itoa(i), one)
	}
	if n := len(tc.m); n > maxTotalCacheEntries {
		t.Fatalf("cache grew to %d entries, cap is %d", n, maxTotalCacheEntries)
	}

	calls := 0
	count := func() (int64, error) { calls++; return 7, nil }
	tc.get("hot", count)
	if v, _ := tc.get("hot", count); v != 7 || calls != 1 {
		t.Fatalf("full cache stopped caching: v=%d calls=%d", v, calls)
	}
}
