package hub

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIsSQLite(t *testing.T) {
	s := newStatTestServer(t) // opens an in-memory sqlite
	if !s.isSQLite() {
		t.Fatal("expected isSQLite()=true for a sqlite-backed server")
	}

	var empty Server // nil gdb
	if empty.isSQLite() {
		t.Fatal("expected isSQLite()=false when gdb is nil")
	}
}

func TestReadonlyRejectsUpload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	g := r.Group("/api")
	(&Server{}).addUploadReadonly(g) // reject closure doesn't touch Server fields

	for _, path := range []string{"/api/upload", "/api/uploadData"} {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", path, nil)
		r.ServeHTTP(w, req)
		if w.Code != http.StatusServiceUnavailable {
			t.Fatalf("%s: want 503 on read-only replica, got %d body=%s", path, w.Code, w.Body.String())
		}
		if !strings.Contains(w.Body.String(), "read-only") {
			t.Errorf("%s: want 'read-only' in body, got %s", path, w.Body.String())
		}
	}
}
