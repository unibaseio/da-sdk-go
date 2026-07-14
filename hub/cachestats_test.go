package hub

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestV1CacheStats: the endpoint aggregates the Phase-Q counters into valid
// JSON and is nil-safe on a bare server (no readCache, no loaded owners, no
// concurrency cap).
func TestV1CacheStats(t *testing.T) {
	gin.SetMode(gin.TestMode)
	s := &Server{} // readCache nil, lfs empty, dlSem nil
	s.fsHit.Store(7)
	s.fsCreate.Store(2)
	s.dlTotal.Store(5)
	s.dlShared.Store(3)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/v1/cachestats", nil)

	s.v1CacheStats(c)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	var body map[string]map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("bad json: %v (%s)", err, w.Body.String())
	}
	if got := body["fs"]["hit"].(float64); got != 7 {
		t.Fatalf("fs.hit = %v, want 7", got)
	}
	if got := body["fs"]["create"].(float64); got != 2 {
		t.Fatalf("fs.create = %v, want 2", got)
	}
	if got := body["download"]["total"].(float64); got != 5 {
		t.Fatalf("download.total = %v, want 5", got)
	}
	if got := body["download"]["shared"].(float64); got != 3 {
		t.Fatalf("download.shared = %v, want 3", got)
	}
	if enabled := body["readcache"]["enabled"].(bool); enabled {
		t.Fatalf("readcache.enabled = true, want false on bare server")
	}
	if lim := body["download"]["concurrency_limit"].(float64); lim != 0 {
		t.Fatalf("download.concurrency_limit = %v, want 0 (unlimited)", lim)
	}
}
