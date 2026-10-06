package hub

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"testing/iotest"
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

// H-Q1: while the signer table is full of signers still using their quota, a
// new signer is refused instead of sharing (and possibly draining) one
// overflow bucket with every other newcomer.
func TestWriteQuotaFullTableRefusesNewSigners(t *testing.T) {
	t.Setenv("HUB_WRITE_QUOTA_BYTES", "1000")
	t.Setenv("HUB_WRITE_QUOTA_WINDOW_SEC", "86400")
	q := newWriteQuota()
	for i := 0; i < 100_000; i++ {
		if err := q.charge(fmt.Sprintf("0x%040x", i), 1000); err != nil {
			t.Fatalf("signer %d: %v", i, err)
		}
	}
	if err := q.charge("0xnew1", 1); err == nil {
		t.Fatal("new signer admitted while every slot is a signer in debt")
	}
	// a signer that has a slot keeps its own budget
	if err := q.charge(fmt.Sprintf("0x%040x", 7), 1); err == nil {
		t.Fatal("a spent signer was given more")
	}
}

// H-Q1: buckets that have refilled completely are dropped (losslessly), so a
// table filled by tiny writes does not push new signers into a shared bucket:
// each new signer still gets its own full quota.
func TestWriteQuotaRefilledBucketsMakeRoom(t *testing.T) {
	t.Setenv("HUB_WRITE_QUOTA_BYTES", "1000")
	t.Setenv("HUB_WRITE_QUOTA_WINDOW_SEC", "1") // refills 1000 B/s
	q := newWriteQuota()
	for i := 0; i < 100_000; i++ {
		if err := q.charge(fmt.Sprintf("0x%040x", i), 1); err != nil {
			t.Fatalf("signer %d: %v", i, err)
		}
	}
	time.Sleep(1100 * time.Millisecond) // every bucket full again
	for _, s := range []string{"0xnewa", "0xnewb"} {
		if err := q.charge(s, 1000); err != nil {
			t.Fatalf("%s: %v (new signers share one bucket)", s, err)
		}
	}
}

// H-Q1: a write that fails before anything is stored gives its charge back.
func TestV1PutRefundsFailedWrite(t *testing.T) {
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

	// declares 6 bytes, the connection breaks after 2
	req := httptest.NewRequest("PUT", "/v1/buckets/b/objects/k1", nil)
	req.Body = io.NopCloser(io.MultiReader(strings.NewReader("12"), iotest.ErrReader(errors.New("connection reset"))))
	req.ContentLength = 6
	req.Header.Set("Authorization", auth)
	w := httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	if w.Code < 400 {
		t.Fatalf("broken upload: %d", w.Code)
	}

	req = httptest.NewRequest("PUT", "/v1/buckets/b/objects/k2", strings.NewReader("0123456789"))
	req.Header.Set("Authorization", auth)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	if w.Code != http.StatusAccepted {
		t.Fatalf("full-quota write after a failed one: %d %s (the failed write kept its charge)", w.Code, w.Body.String())
	}
}

// H-Q1: a hub-paid seal turned away before any work gives its charge back.
func TestSealRefundsWhenBusy(t *testing.T) {
	t.Setenv("HUB_WRITE_QUOTA_BYTES", "10")
	t.Setenv("HUB_SEAL_QUEUE_SEC", "0")
	s := newV1TestServer(t)
	s.quota = newWriteQuota()
	s.sealSem = make(chan struct{}, 1)
	s.sealSem <- struct{}{}
	addr, pk := testKey(t)
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	mw.WriteField("owner", addr)
	mw.WriteField("register", "hub")
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
	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("busy seal: %d %s", w.Code, w.Body.String())
	}
	if err := s.quota.charge(addr, 10); err != nil {
		t.Fatalf("seal refused for being busy kept its charge: %v", err)
	}
}

func TestWriteQuotaRefund(t *testing.T) {
	t.Setenv("HUB_WRITE_QUOTA_BYTES", "1000")
	t.Setenv("HUB_WRITE_QUOTA_WINDOW_SEC", "86400")
	q := newWriteQuota()
	if err := q.charge("0xa", 1000); err != nil {
		t.Fatal(err)
	}
	q.refund("0xA", 400)
	if err := q.charge("0xa", 400); err != nil {
		t.Fatalf("refunded bytes not available: %v", err)
	}
	if err := q.charge("0xa", 1); err == nil {
		t.Fatal("refund gave back more than it was given")
	}
	q.refund("0xa", 1<<40) // capped at the full quota
	if err := q.charge("0xa", 1000); err != nil {
		t.Fatal(err)
	}
	if err := q.charge("0xa", 1); err == nil {
		t.Fatal("refund overfilled the bucket")
	}
	var off *writeQuota
	off.refund("0xa", 1) // disabled quota: no-op
}
