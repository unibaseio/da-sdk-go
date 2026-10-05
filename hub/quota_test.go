package hub

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/kv"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

func TestWriteQuotaCharge(t *testing.T) {
	t.Setenv("HUB_WRITE_QUOTA_BYTES", "1000")
	t.Setenv("HUB_WRITE_QUOTA_WINDOW_SEC", "86400")
	t.Setenv("HUB_WRITE_QUOTA_EXEMPT", " 0xEE ")
	q := newWriteQuota()
	if err := q.charge("0xA", 600); err != nil {
		t.Fatal(err)
	}
	if err := q.charge("0xa", 600); err == nil {
		t.Fatal("over quota accepted (case must not split a signer)")
	}
	if err := q.charge("0xa", 400); err != nil {
		t.Fatalf("remaining quota refused: %v", err)
	}
	if err := q.charge("0xb", 1001); err == nil {
		t.Fatal("single write over the whole quota accepted")
	}
	if err := q.charge("0xee", 1<<40); err != nil {
		t.Fatalf("exempt signer charged: %v", err)
	}

	t.Setenv("HUB_WRITE_QUOTA_BYTES", "0")
	if newWriteQuota() != nil {
		t.Fatal("0 must disable the quota")
	}
	var off *writeQuota
	if err := off.charge("0xa", 1<<40); err != nil {
		t.Fatal(err)
	}
}

// PUT and POST objects are charged to the signer; an over-quota PUT with a
// declared length is refused before its body is read.
func TestV1WriteQuota(t *testing.T) {
	t.Setenv("HUB_WRITE_QUOTA_BYTES", "10")
	s := newV1TestServer(t)
	s.quota = newWriteQuota()
	addr, pk := testKey(t)
	owner := strings.ToLower(addr)
	auth := authHeader(addr, pk)

	dir := t.TempDir()
	ds, err := kv.NewBadgerStore(filepath.Join(dir, "kv"), &kv.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ds.Close() })
	fs, err := logfs.New(ds, filepath.Join(dir, "logfs"), "0xhub", owner)
	if err != nil {
		t.Fatal(err)
	}
	s.lfs.Store(owner, fs)
	s.gdb.Create(&types.Bucket{Name: "b", Owner: owner, Kind: "memory"})

	put := func(key, body string) int {
		req := httptest.NewRequest("PUT", "/v1/buckets/b/objects/"+key, strings.NewReader(body))
		req.Header.Set("Authorization", auth)
		w := httptest.NewRecorder()
		s.Router.ServeHTTP(w, req)
		return w.Code
	}
	if c := put("k1", "123456"); c != http.StatusAccepted {
		t.Fatalf("first put: %d", c)
	}
	if c := put("k2", "123456"); c != http.StatusTooManyRequests {
		t.Fatalf("over-quota put: %d", c)
	}
	var n int64
	s.gdb.Model(&types.Needle{}).Where("name = ?", "k2").Count(&n)
	if n != 0 {
		t.Fatal("over-quota object was written")
	}

	// multipart batch: charged as a whole (4 left, 2x3 bytes requested)
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	for _, name := range []string{"a", "b"} {
		fw, _ := mw.CreateFormFile("files", name)
		fw.Write([]byte("xyz"))
	}
	mw.Close()
	req := httptest.NewRequest("POST", "/v1/buckets/b/objects", &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("Authorization", auth)
	w := httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("over-quota batch: %d %s", w.Code, w.Body.String())
	}
}

// Hub-paid seals are charged before any work; register=client is not.
func TestSealQuotaAndConcurrency(t *testing.T) {
	t.Setenv("HUB_WRITE_QUOTA_BYTES", "4")
	s := newV1TestServer(t)
	s.quota = newWriteQuota()
	addr, pk := testKey(t)
	seal := func(register string) *httptest.ResponseRecorder {
		var body bytes.Buffer
		mw := multipart.NewWriter(&body)
		mw.WriteField("owner", addr)
		mw.WriteField("register", register)
		mw.WriteField("rsn", "6")
		mw.WriteField("rsk", "4")
		fw, _ := mw.CreateFormFile("file", "blob")
		fw.Write([]byte("12345"))
		mw.Close()
		req := httptest.NewRequest("POST", "/v1/seal", &body)
		req.Header.Set("Content-Type", mw.FormDataContentType())
		req.Header.Set("Authorization", authHeader(addr, pk))
		w := httptest.NewRecorder()
		s.Router.ServeHTTP(w, req)
		return w
	}
	for _, mode := range []string{"hub", "hub_attributed"} {
		if w := seal(mode); w.Code != http.StatusTooManyRequests {
			t.Fatalf("%s seal over quota: %d %s", mode, w.Code, w.Body.String())
		}
	}

	// client mode skips the quota and queues for a seal slot: with the only
	// slot taken and no queueing allowed it is turned away with 503
	t.Setenv("HUB_SEAL_QUEUE_SEC", "0")
	s.sealSem = make(chan struct{}, 1)
	s.sealSem <- struct{}{}
	if w := seal("client"); w.Code != http.StatusServiceUnavailable || w.Header().Get("Retry-After") == "" {
		t.Fatalf("busy seal: %d %s", w.Code, w.Body.String())
	}
}

func TestStoreDurationCheck(t *testing.T) {
	d := storeDuration{min: 1200, max: 5000}
	if err := d.check(100, 1299); err == nil {
		t.Fatal("term below minStore accepted")
	}
	if err := d.check(100, 1300); err != nil {
		t.Fatal(err)
	}
	if err := d.check(100, 5101); err == nil {
		t.Fatal("term above maxStore accepted")
	}
	// the contract has no "0 = unbounded" case: maxStore 0 rejects every term
	if err := (storeDuration{min: 10}).check(100, 1<<40); err == nil {
		t.Fatal("term accepted with maxStore 0, which the contract rejects")
	}
}

// Public piece reads wait for a pieceSem slot and give up with 503.
func TestV1PieceContentBounded(t *testing.T) {
	s := newV1TestServer(t)
	s.pieceSem = make(chan struct{}, 1)
	s.pieceSem <- struct{}{}
	old := maxSlotWait
	maxSlotWait = 50 * time.Millisecond
	defer func() { maxSlotWait = old }()

	if w := do(t, s, "GET", "/v1/pieces/abcd/content", "", ""); w.Code != http.StatusServiceUnavailable {
		t.Fatalf("got %d %s", w.Code, w.Body.String())
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := acquireSem(ctx, s.pieceSem); err != errBusy {
		t.Fatalf("cancelled wait: %v", err)
	}
}

func TestReadAllSized(t *testing.T) {
	for _, hint := range []int64{-1, 0, 3, 5, 100} {
		b, err := readAllSized(strings.NewReader("hello"), hint)
		if err != nil || string(b) != "hello" {
			t.Fatalf("hint %d: %q %v", hint, b, err)
		}
	}
}

// A declared length is not trusted for the allocation: claiming the full cap
// with a tiny body must not reserve the cap.
func TestReadAllSizedBoundsPrealloc(t *testing.T) {
	b, err := readAllSized(strings.NewReader("hello"), 64<<20)
	if err != nil || string(b) != "hello" {
		t.Fatalf("%q %v", b, err)
	}
	if cap(b) > maxPrealloc+bytes.MinRead {
		t.Fatalf("preallocated %d bytes for a 5-byte body", cap(b))
	}
	big := strings.Repeat("x", 3<<20) // larger than the prealloc: still read whole
	if b, err := readAllSized(strings.NewReader(big), int64(len(big))); err != nil || len(b) != len(big) {
		t.Fatalf("large body: %d %v", len(b), err)
	}
}
