package logfs

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/kv"
)

// fakeBackend is an in-memory VolumeBackend: PutVolume snapshots the sealed
// volume bytes, RangeAt serves a byte range from that snapshot. Mirrors what an
// S3/minio backend does (whole-volume object + range GET) with no infra.
type fakeBackend struct {
	mu   sync.Mutex
	vols map[uint64][]byte
	puts int
}

func newFakeBackend() *fakeBackend { return &fakeBackend{vols: map[uint64][]byte{}} }

func (b *fakeBackend) PutVolume(idx uint64, localPath string) error {
	data, err := os.ReadFile(localPath)
	if err != nil {
		return err
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	// copy: the local file may be reclaimed later
	cp := append([]byte(nil), data...)
	b.vols[idx] = cp
	b.puts++
	return nil
}

func (b *fakeBackend) RangeAt(idx uint64, off, n int64) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	v, ok := b.vols[idx]
	if !ok {
		return nil, fmt.Errorf("no such volume %d", idx)
	}
	if off < 0 || off+n > int64(len(v)) {
		return nil, fmt.Errorf("range out of bounds")
	}
	return append([]byte(nil), v[off:off+n]...), nil
}

func (b *fakeBackend) has(idx uint64) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	_, ok := b.vols[idx]
	return ok
}

func newTestFSWithBackend(t *testing.T, vb VolumeBackend) *LogFS {
	t.Helper()
	dir := t.TempDir()
	ds, err := kv.NewBadgerStore(filepath.Join(dir, "kv"), &kv.DefaultOptions)
	if err != nil {
		t.Fatalf("badger: %v", err)
	}
	fs, err := New(ds, filepath.Join(dir, "data"), "0xlocal", "0xowner", WithVolumeBackend(vb))
	if err != nil {
		t.Fatalf("logfs new: %v", err)
	}
	return fs
}

// waitFor polls until cond() or the deadline (seal-upload is async).
func waitFor(t *testing.T, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("condition not met before deadline")
}

// TestVolumeBackendUploadOnSeal: sealing a volume (via Roll) hands it to the
// backend exactly once, and the object bytes equal the local file.
func TestVolumeBackendUploadOnSeal(t *testing.T) {
	vb := newFakeBackend()
	fs := newTestFSWithBackend(t, vb)
	defer fs.Close()

	if err := fs.Put([]byte("k1"), bytes.Repeat([]byte{0xab}, 100)); err != nil {
		t.Fatal(err)
	}
	sealedIdx := fs.curIndex
	if err := fs.Roll(); err != nil { // seals the current volume
		t.Fatal(err)
	}
	waitFor(t, func() bool { return vb.has(sealedIdx) })

	// the uploaded object must byte-match the on-disk sealed volume
	localPath := filepath.Join(fs.basedir, fmt.Sprintf("%d.vol", sealedIdx))
	want, err := os.ReadFile(localPath)
	if err != nil {
		t.Fatal(err)
	}
	vb.mu.Lock()
	got := vb.vols[sealedIdx]
	vb.mu.Unlock()
	if !bytes.Equal(got, want) {
		t.Fatalf("uploaded bytes != local volume (%d vs %d)", len(got), len(want))
	}
}

// TestVolumeBackendReadFallback: when the local sealed volume is deleted (as TTL
// reclaim will do), reads transparently fall back to the backend RangeAt and
// still pass the content-hash check.
func TestVolumeBackendReadFallback(t *testing.T) {
	vb := newFakeBackend()
	fs := newTestFSWithBackend(t, vb)
	defer fs.Close()

	val := bytes.Repeat([]byte{0x5c}, 137)
	if err := fs.Put([]byte("kfall"), val); err != nil {
		t.Fatal(err)
	}
	sealedIdx := fs.curIndex
	if err := fs.Roll(); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return vb.has(sealedIdx) })

	// simulate TTL reclaim: drop the local sealed volume + any cached fd
	fs.fdc.closeAll()
	if err := os.Remove(filepath.Join(fs.basedir, fmt.Sprintf("%d.vol", sealedIdx))); err != nil {
		t.Fatal(err)
	}

	got, err := fs.Get([]byte("kfall"))
	if err != nil {
		t.Fatalf("get via backend fallback: %v", err)
	}
	if !bytes.Equal(got, val) {
		t.Fatalf("fallback read mismatch")
	}
}

// TestVolumeBackendCurrentVolumeStaysLocal: the open (unsealed) volume is never
// uploaded — reads of a just-written, not-yet-sealed needle resolve locally even
// with a backend attached (backend has nothing).
func TestVolumeBackendCurrentVolumeStaysLocal(t *testing.T) {
	vb := newFakeBackend()
	fs := newTestFSWithBackend(t, vb)
	defer fs.Close()

	val := bytes.Repeat([]byte{0x11}, 64)
	if err := fs.Put([]byte("konly"), val); err != nil {
		t.Fatal(err)
	}
	got, err := fs.Get([]byte("konly"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, val) {
		t.Fatal("local read mismatch")
	}
	vb.mu.Lock()
	puts := vb.puts
	vb.mu.Unlock()
	if puts != 0 {
		t.Fatalf("open volume must not be uploaded, saw %d puts", puts)
	}
}

// TestVolumeBackendLocalTTLReclaim: with a backend + short TTL, a sealed+
// uploaded volume's local file is reclaimed, yet the needle still reads back
// (served from the backend). The current open volume is never reclaimed.
func TestVolumeBackendLocalTTLReclaim(t *testing.T) {
	vb := newFakeBackend()
	dir := t.TempDir()
	ds, err := kv.NewBadgerStore(filepath.Join(dir, "kv"), &kv.DefaultOptions)
	if err != nil {
		t.Fatalf("badger: %v", err)
	}
	fs, err := New(ds, filepath.Join(dir, "data"), "0xlocal", "0xowner",
		WithVolumeBackend(vb), WithLocalTTL(50*time.Millisecond))
	if err != nil {
		t.Fatalf("logfs new: %v", err)
	}
	defer fs.Close()

	val := bytes.Repeat([]byte{0x7e}, 111)
	if err := fs.Put([]byte("kttl"), val); err != nil {
		t.Fatal(err)
	}
	sealedIdx := fs.curIndex
	if err := fs.Roll(); err != nil {
		t.Fatal(err)
	}
	sealedPath := filepath.Join(fs.basedir, fmt.Sprintf("%d.vol", sealedIdx))

	// wait for upload + reclaim to remove the local sealed volume
	waitFor(t, func() bool {
		_, statErr := os.Stat(sealedPath)
		return os.IsNotExist(statErr)
	})

	// the needle still reads back via the backend fallback
	got, err := fs.Get([]byte("kttl"))
	if err != nil {
		t.Fatalf("get after reclaim: %v", err)
	}
	if !bytes.Equal(got, val) {
		t.Fatal("post-reclaim read mismatch")
	}
}

// TestVolumeBackendNilDefault: no backend = historical behavior, reads local,
// no panic on the fallback path.
func TestVolumeBackendNilDefault(t *testing.T) {
	fs := newTestFS(t) // no backend
	defer fs.Close()
	if fs.volBackend != nil {
		t.Fatal("default backend should be nil")
	}
	val := bytes.Repeat([]byte{0x22}, 80)
	if err := fs.Put([]byte("kdef"), val); err != nil {
		t.Fatal(err)
	}
	got, err := fs.Get([]byte("kdef"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, val) {
		t.Fatal("default read mismatch")
	}
}
