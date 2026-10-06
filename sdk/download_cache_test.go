package sdk

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// memPieces is an in-memory piece cache that records what was put into it.
type memPieces struct {
	types.IPieceStore
	mu   sync.Mutex
	data map[string][]byte
}

func newMemPieces() *memPieces { return &memPieces{data: map[string][]byte{}} }

func (m *memPieces) GetPiece(_ context.Context, name string, w io.Writer, _ types.Options) (types.PieceReceipt, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	b, ok := m.data[name]
	if !ok {
		return types.PieceReceipt{}, fmt.Errorf("miss")
	}
	if w != nil {
		w.Write(b)
	}
	return types.PieceReceipt{}, nil
}

func (m *memPieces) PutPiece(_ context.Context, pc types.PieceCore, b []byte, _ bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[pc.Name] = append([]byte(nil), b...)
	return nil
}

func (m *memPieces) DeleteData(_ context.Context, name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, name)
	return nil
}

func (m *memPieces) len() int { m.mu.Lock(); defer m.mu.Unlock(); return len(m.data) }

// fakeNet serves a gateway + stream for one file of two pieces. Each piece is
// 248 bytes (K=4 shards of 2 field elements, no zero padding). corrupt, when
// set, flips a byte in every replica the stream relays.
type fakeNet struct {
	srv     *httptest.Server
	file    []byte
	fr      types.FileReceipt
	shards  map[string][]byte
	pieces  map[string]types.PieceReceipt
	corrupt bool
}

func newFakeNet(t *testing.T, hash func(sum string) string) *fakeNet {
	t.Helper()
	const k, slen = 4, 2
	pol := types.Policy{N: 6, K: k}
	per := 31 * slen
	n := &fakeNet{shards: map[string][]byte{}, pieces: map[string]types.PieceReceipt{}}
	var names []string
	for p := 0; p < 2; p++ {
		data := make([]byte, per*k)
		for i := range data {
			data[i] = byte(p*101 + i*13 + 1)
		}
		n.file = append(n.file, data...)
		name := fmt.Sprintf("piece%d", p)
		pr := types.PieceReceipt{
			PieceCore: types.PieceCore{Name: name, Policy: pol, Size: int64(len(data))},
			Replicas:  make([]string, k),
			StoredOn:  make([]common.Address, k),
		}
		for j := 0; j < k; j++ {
			rn := fmt.Sprintf("%s-r%d", name, j)
			n.shards[rn] = bls.Pad(data[j*per : (j+1)*per])
			pr.Replicas[j] = rn
			pr.StoredOn[j] = common.HexToAddress("0x5")
		}
		n.pieces[name] = pr
		names = append(names, name)
	}
	sum := sha256.Sum256(n.file)
	n.fr = types.FileReceipt{
		FileCore: types.FileCore{Name: "f", Policy: pol, Hash: hash(hex.EncodeToString(sum[:])), Size: int64(len(n.file))},
		Pieces:   names,
	}
	n.srv = httptest.NewServer(http.HandlerFunc(n.serve))
	t.Cleanup(n.srv.Close)
	return n
}

func (n *fakeNet) serve(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/v1/files/"):
		json.NewEncoder(w).Encode(n.fr)
	case strings.HasPrefix(r.URL.Path, "/v1/pieces/"):
		pr, ok := n.pieces[strings.TrimPrefix(r.URL.Path, "/v1/pieces/")]
		if !ok {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(pr)
	case r.URL.Path == "/v1/edges":
		e := types.EdgeReceipt{EdgeMeta: types.EdgeMeta{
			Name: common.HexToAddress("0x9"), Type: types.StreamType, ExposeURL: n.srv.URL, ChainType: chaintype,
		}, OnChain: true}
		json.NewEncoder(w).Encode(map[string]any{"items": []types.EdgeReceipt{e}})
	case r.URL.Path == "/v1/download":
		r.ParseForm()
		b, ok := n.shards[r.Form.Get("name")]
		if !ok {
			http.NotFound(w, r)
			return
		}
		b = append([]byte(nil), b...)
		if n.corrupt {
			b[5] ^= 0xff
		}
		w.Write(b)
	default:
		http.NotFound(w, r)
	}
}

func same(s string) string { return s }

func TestCheckFileParallelCachesOnlyVerifiedPieces(t *testing.T) {
	n := newFakeNet(t, same)
	ks := newMemPieces()
	if err := CheckFileParallelOf(n.srv.URL, types.Auth{}, "f", types.EmptyAddr, 2, ks); err != nil {
		t.Fatalf("honest file: %v", err)
	}
	if ks.len() != 2 {
		t.Fatalf("cached %d pieces, want 2", ks.len())
	}
	var out bytes.Buffer
	if err := DownloadParallelOf(n.srv.URL, types.Auth{}, "f", types.EmptyAddr, 2, ks, &out); err != nil {
		t.Fatalf("download: %v", err)
	}
	if !bytes.Equal(out.Bytes(), n.file) {
		t.Fatal("downloaded data differs")
	}
}

func TestCheckFileParallelRejectsLyingStores(t *testing.T) {
	n := newFakeNet(t, same)
	n.corrupt = true
	ks := newMemPieces()
	err := CheckFileParallelOf(n.srv.URL, types.Auth{}, "f", types.EmptyAddr, 2, ks)
	if !errors.Is(err, ErrFileHashMismatch) {
		t.Fatalf("err=%v, want a hash mismatch", err)
	}
	if ks.len() != 0 {
		t.Fatalf("cached %d unverified pieces", ks.len())
	}
	var out bytes.Buffer
	if err := DownloadParallelOf(n.srv.URL, types.Auth{}, "f", types.EmptyAddr, 2, ks, &out); !errors.Is(err, ErrFileHashMismatch) {
		t.Fatalf("download err=%v, want a hash mismatch", err)
	}
	if ks.len() != 0 {
		t.Fatalf("download cached %d unverified pieces", ks.len())
	}
}

func TestDownloadRejectsBlankHash(t *testing.T) {
	n := newFakeNet(t, func(string) string { return "" })
	ks := newMemPieces()
	if err := CheckFileParallelOf(n.srv.URL, types.Auth{}, "f", types.EmptyAddr, 2, ks); err == nil {
		t.Fatal("check accepted a receipt without a hash")
	}
	var out bytes.Buffer
	if err := DownloadOf(n.srv.URL, types.Auth{}, "f", types.EmptyAddr, ks, &out); err == nil {
		t.Fatal("download accepted a receipt without a hash")
	}
	if ks.len() != 0 {
		t.Fatalf("cached %d pieces", ks.len())
	}
}

func TestDownloadPieceAndSaveVerified(t *testing.T) {
	n := newFakeNet(t, same)
	ks := newMemPieces()
	reject := func(types.PieceCore, []byte) error { return fmt.Errorf("no") }
	if err := DownloadPieceAndSaveVerified(n.srv.URL, types.Auth{}, "piece0", ks, reject); err == nil || ks.len() != 0 {
		t.Fatalf("rejected piece: err=%v cached=%d", err, ks.len())
	}
	if err := DownloadPieceAndSaveVerified(n.srv.URL, types.Auth{}, "piece0", ks, nil); err != nil || ks.len() != 1 {
		t.Fatalf("accepted piece: err=%v cached=%d", err, ks.len())
	}
}

// S4: a bad piece in the cache must not fail every later read of the file:
// on a hash mismatch the cached pieces are evicted and the file fetched again.
func TestDownloadEvictsPoisonedCache(t *testing.T) {
	for _, mode := range []string{"serial", "parallel"} {
		n := newFakeNet(t, same)
		ks := newMemPieces()
		poisoned := bytes.Repeat([]byte{0xee}, 248)
		ks.PutPiece(context.Background(), types.PieceCore{Name: "piece1"}, poisoned, true)

		var out bytes.Buffer
		var err error
		if mode == "serial" {
			err = DownloadOf(n.srv.URL, types.Auth{}, "f", types.EmptyAddr, ks, &out)
		} else {
			err = DownloadParallelOf(n.srv.URL, types.Auth{}, "f", types.EmptyAddr, 2, ks, &out)
		}
		if err != nil {
			t.Fatalf("%s: poisoned cache entry failed the download: %v", mode, err)
		}
		if !bytes.Equal(out.Bytes(), n.file) {
			t.Fatalf("%s: wrote %d bytes, not the file", mode, out.Len())
		}
		var b bytes.Buffer
		if _, err := ks.GetPiece(context.Background(), "piece1", &b, types.Options{}); err != nil || bytes.Equal(b.Bytes(), poisoned) {
			t.Fatalf("%s: poisoned piece still cached (err=%v)", mode, err)
		}
	}
}

// With a cache, a file that fails its hash is not written to w at all.
func TestDownloadWithCacheWritesOnlyCheckedData(t *testing.T) {
	n := newFakeNet(t, same)
	n.corrupt = true
	ks := newMemPieces()
	var out bytes.Buffer
	if err := DownloadOf(n.srv.URL, types.Auth{}, "f", types.EmptyAddr, ks, &out); !errors.Is(err, ErrFileHashMismatch) {
		t.Fatalf("err=%v, want a hash mismatch", err)
	}
	if out.Len() != 0 || ks.len() != 0 {
		t.Fatalf("wrote %d bytes / cached %d pieces of a file that failed its hash", out.Len(), ks.len())
	}
}

// A range read cannot be checked, so it must not cache what it fetched.
func TestDownloadWSizeDoesNotCache(t *testing.T) {
	n := newFakeNet(t, same)
	ks := newMemPieces()
	var out bytes.Buffer
	if err := DownloadWSize(n.srv.URL, types.Auth{}, "f", ks, &out, 250, 10); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(out.Bytes(), n.file[250:260]) {
		t.Fatalf("range read %x, want %x", out.Bytes(), n.file[250:260])
	}
	if ks.len() != 0 {
		t.Fatalf("range read cached %d unchecked pieces", ks.len())
	}
}
