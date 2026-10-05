package hub

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
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

// A reader replica (HUB_READONLY) must not write to its local LogFS/index,
// which would fork the data: every /v1 write answers 503, still behind auth.
func TestReadonlyRejectsWrites(t *testing.T) {
	s := newV1TestServerMode(t, true)
	hdr := buildHeader(t, testSK, "hub", time.Now().Unix())

	var form bytes.Buffer
	mw := multipart.NewWriter(&form)
	mw.WriteField("owner", testAddr)
	mw.Close()

	reqs := []*http.Request{
		signedWrite(hdr, "b"),
		httptest.NewRequest("PUT", "/v1/buckets/b/objects/k", strings.NewReader("data")),
		httptest.NewRequest("POST", "/v1/buckets/b/objects", bytes.NewReader(form.Bytes())),
		httptest.NewRequest("POST", "/v1/seal", bytes.NewReader(form.Bytes())),
	}
	reqs[2].Header.Set("Content-Type", mw.FormDataContentType())
	reqs[3].Header.Set("Content-Type", mw.FormDataContentType())
	for _, req := range reqs {
		req.Header.Set("Authorization", hdr)
		w := serve(s, req)
		if w.Code != http.StatusServiceUnavailable {
			t.Fatalf("%s %s: want 503 on read-only replica, got %d body=%s", req.Method, req.URL.Path, w.Code, w.Body.String())
		}
		if !strings.Contains(w.Body.String(), "read-only") {
			t.Errorf("%s %s: want 'read-only' in body, got %s", req.Method, req.URL.Path, w.Body.String())
		}
	}

	// DELETE isn't mounted on a reader at all
	req := httptest.NewRequest("DELETE", "/v1/buckets/b", nil)
	req.Header.Set("Authorization", hdr)
	if w := serve(s, req); w.Code >= 200 && w.Code < 300 {
		t.Fatalf("DELETE on read-only replica: got %d", w.Code)
	}

	// the reject runs after auth: an unsigned write is still a 401
	if w := serve(s, signedWrite("", "b")); w.Code != http.StatusUnauthorized {
		t.Fatalf("unsigned write on read-only replica: want 401, got %d", w.Code)
	}

	var n int64
	s.gdb.Model(&types.Bucket{}).Count(&n)
	if n != 0 {
		t.Fatalf("read-only replica wrote %d bucket rows", n)
	}
}
