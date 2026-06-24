package hub

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

// defaultMemStatRefreshSec is how often (seconds) the per-owner memory
// aggregates are recomputed in the background. Conservative by default because
// each recompute is an index scan over the whole needles table (tens of
// millions of rows). Override with HUB_MEMSTAT_REFRESH_SEC.
const defaultMemStatRefreshSec int64 = 300

func memStatRefreshInterval() time.Duration {
	return time.Duration(envInt64("HUB_MEMSTAT_REFRESH_SEC", defaultMemStatRefreshSec)) * time.Second
}

// memStatSnapshot is a point-in-time result of computeMemStats, served to the
// /api/memoryOverview and /api/memoryStat endpoints without touching the DB.
type memStatSnapshot struct {
	overview   types.MemoryOverview
	owners     []types.MemoryStat // sorted by bytes desc
	computedAt time.Time
}

type memStatCache struct {
	mu   sync.RWMutex
	snap *memStatSnapshot
}

func (m *memStatCache) get() *memStatSnapshot {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.snap
}

func (m *memStatCache) set(s *memStatSnapshot) {
	m.mu.Lock()
	m.snap = s
	m.mu.Unlock()
}

// startMemStats launches the background recompute loop. The aggregates are
// full-table scans — too heavy to run inline per request (they timed out
// behind the reverse proxy) — so we compute them periodically and serve the
// cached snapshot instantly. The first compute runs async so server startup
// is never blocked; until it lands the endpoints return an empty (ComputedAt=0)
// result.
func (s *Server) startMemStats(ctx context.Context) {
	go func() {
		s.refreshMemStats()
		t := time.NewTicker(memStatRefreshInterval())
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-s.shutdownChan:
				return
			case <-t.C:
				s.refreshMemStats()
			}
		}
	}()
}

func (s *Server) refreshMemStats() {
	ov, owners, err := s.computeMemStats()
	if err != nil {
		logger.Warnf("memstat recompute failed: %v", err)
		return
	}
	s.memStat.set(&memStatSnapshot{
		overview:   ov,
		owners:     owners,
		computedAt: time.Now(),
	})
	logger.Infof("memstat refreshed: %d owners, %d entries, %d bytes",
		ov.WalletsWithMemory, ov.MemoryCount, ov.MemoryBytes)
}

// memoryOverviewSnapshot returns the cached overview (zero value with
// ComputedAt=0 if the first compute hasn't completed yet).
func (s *Server) memoryOverviewSnapshot() types.MemoryOverview {
	snap := s.memStat.get()
	if snap == nil {
		return types.MemoryOverview{}
	}
	ov := snap.overview
	ov.ComputedAt = snap.computedAt.Unix()
	return ov
}

// memoryStatPage paginates the cached per-owner list in memory. When owner is
// non-empty it filters to that single wallet (case-insensitive) — the result
// is still a list (0 or 1 item).
func (s *Server) memoryStatPage(owner string, offset, length int) types.MemoryStatResult {
	res := types.MemoryStatResult{Offset: offset, Length: length, Items: []types.MemoryStat{}}
	snap := s.memStat.get()
	if snap == nil {
		return res
	}
	res.ComputedAt = snap.computedAt.Unix()

	items := snap.owners
	if owner != "" {
		// snapshot owners are lowercased (GROUP BY LOWER(owner)); match the same.
		want := strings.ToLower(strings.TrimSpace(owner))
		items = []types.MemoryStat{}
		for _, o := range snap.owners {
			if o.Owner == want {
				items = []types.MemoryStat{o}
				break
			}
		}
	}

	res.Total = int64(len(items))
	if offset < len(items) {
		end := offset + length
		if end > len(items) {
			end = len(items)
		}
		res.Items = items[offset:end]
	}
	return res
}
