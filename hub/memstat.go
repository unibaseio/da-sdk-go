package hub

import (
	"context"
	"sync"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

// memStatRefreshInterval is how often the per-owner memory aggregates are
// recomputed in the background.
const memStatRefreshInterval = 2 * time.Minute

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
		t := time.NewTicker(memStatRefreshInterval)
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

// memoryStatPage paginates the cached per-owner list in memory.
func (s *Server) memoryStatPage(offset, length int) types.MemoryStatResult {
	res := types.MemoryStatResult{Offset: offset, Length: length, Items: []types.MemoryStat{}}
	snap := s.memStat.get()
	if snap == nil {
		return res
	}
	res.Total = int64(len(snap.owners))
	res.ComputedAt = snap.computedAt.Unix()
	if offset < len(snap.owners) {
		end := offset + length
		if end > len(snap.owners) {
			end = len(snap.owners)
		}
		res.Items = snap.owners[offset:end]
	}
	return res
}
