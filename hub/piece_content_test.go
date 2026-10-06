package hub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/unibaseio/da-sdk-go/build"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// pieceNet is a gateway that is also the (only) relay stream, serving one
// 6/4 piece of 248 bytes whose K data replicas are on store 0x5.
type pieceNet struct {
	srv  *httptest.Server
	cid  string
	data []byte
}

func newPieceNet(t *testing.T) *pieceNet {
	t.Helper()
	const k, slen = 4, 2
	n := &pieceNet{cid: fmt.Sprintf("%096x", 0xc1d)}
	n.data = make([]byte, 31*slen*k)
	for i := range n.data {
		n.data[i] = byte(i*5 + 1)
	}
	pr := types.PieceReceipt{
		PieceCore: types.PieceCore{Name: n.cid, Policy: types.Policy{N: 6, K: k}, Size: int64(len(n.data))},
		Replicas:  make([]string, k),
		StoredOn:  make([]common.Address, k),
	}
	shards := map[string][]byte{}
	for j := 0; j < k; j++ {
		name := fmt.Sprintf("%096x", j+1)
		pr.Replicas[j], pr.StoredOn[j] = name, common.HexToAddress("0x5")
		shards[name] = bls.Pad(n.data[j*31*slen : (j+1)*31*slen])
	}
	n.srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/v1/pieces/"+n.cid:
			json.NewEncoder(w).Encode(pr)
		case r.URL.Path == "/v1/edges":
			e := types.EdgeReceipt{EdgeMeta: types.EdgeMeta{
				Name: common.HexToAddress("0x9"), Type: types.StreamType, ExposeURL: n.srv.URL, ChainType: build.CheckChain(),
			}, OnChain: true}
			json.NewEncoder(w).Encode(map[string]any{"items": []types.EdgeReceipt{e}})
		case r.URL.Path == "/v1/download":
			r.ParseForm()
			b, ok := shards[r.Form.Get("name")]
			if !ok {
				http.NotFound(w, r)
				return
			}
			w.Write(b)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(n.srv.Close)
	old := build.ServerURL
	build.ServerURL = n.srv.URL
	t.Cleanup(func() { build.ServerURL = old })
	return n
}

// recPieces is an in-memory piece store recording what is put into it.
type recPieces struct {
	types.IPieceStore
	mu   sync.Mutex
	data map[string][]byte
	puts int
}

func (r *recPieces) GetPiece(_ context.Context, name string, w io.Writer, _ types.Options) (types.PieceReceipt, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	b, ok := r.data[name]
	if !ok {
		return types.PieceReceipt{}, fmt.Errorf("miss")
	}
	if w != nil {
		w.Write(b)
	}
	return types.PieceReceipt{}, nil
}

func (r *recPieces) PutPiece(_ context.Context, pc types.PieceCore, b []byte, _ bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.puts++
	r.data[pc.Name] = append([]byte(nil), b...)
	return nil
}

// S4(b): the public piece endpoint serves a rebuilt piece but never stores
// it: nothing checks a bare piece's bytes, and the store is shared with file
// reads.
func TestV1PieceContentNotCached(t *testing.T) {
	n := newPieceNet(t)
	s := newV1TestServer(t)
	ps := &recPieces{data: map[string][]byte{}}
	s.ps = ps

	for i := 0; i < 2; i++ {
		w := do(t, s, "GET", "/v1/pieces/"+n.cid+"/content", "", "")
		if w.Code != http.StatusOK || !bytes.Equal(w.Body.Bytes(), n.data) {
			t.Fatalf("read %d: %d %q", i, w.Code, strings.TrimSpace(w.Body.String()))
		}
	}
	if ps.puts != 0 {
		t.Fatalf("an unverified piece was put into the piece store (%d puts)", ps.puts)
	}

	// a piece the store does hold (checked as part of a file) is served from it
	ps.data[n.cid] = []byte("from-store")
	if w := do(t, s, "GET", "/v1/pieces/"+n.cid+"/content", "", ""); w.Body.String() != "from-store" {
		t.Fatalf("stored piece not served from the store: %q", w.Body.String())
	}
}

// stalledPieceRead starts a piece read of 64 MiB on a real server from a
// client that sends the request and then does not read.
func stalledPieceRead(t *testing.T) (*Server, net.Conn, int) {
	t.Helper()
	t.Setenv("HUB_PIECE_WRITE_BASE_SEC", "1")
	t.Setenv("HUB_PIECE_WRITE_MIN_BPS", "1099511627776") // the size adds ~nothing
	s := newV1TestServer(t)
	big := bytes.Repeat([]byte{0x5a}, 64<<20) // far more than socket buffers hold
	s.ps = &recPieces{data: map[string][]byte{"c1d": big}}
	s.pieceSem = make(chan struct{}, 1)
	srv := httptest.NewServer(s.Router)
	t.Cleanup(srv.Close)
	conn, err := net.Dial("tcp", srv.Listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { conn.Close() }) // runs before srv.Close
	if _, err := conn.Write([]byte("GET /v1/pieces/c1d/content HTTP/1.1\r\nHost: hub\r\n\r\n")); err != nil {
		t.Fatal(err)
	}
	return s, conn, len(big)
}

// H-Q2: the piece slot is held until the response has been written (it
// bounds the pieces held in memory), and released once the write ends.
func TestV1PieceContentSlotCoversWrite(t *testing.T) {
	s, _, _ := stalledPieceRead(t)
	time.Sleep(300 * time.Millisecond) // the handler is now blocked writing
	if len(s.pieceSem) != 1 {
		t.Fatal("piece slot released while the piece is still being written")
	}
	deadline := time.Now().Add(10 * time.Second)
	for len(s.pieceSem) != 0 && time.Now().Before(deadline) {
		time.Sleep(50 * time.Millisecond)
	}
	if len(s.pieceSem) != 0 {
		t.Fatal("slot still held long after the write deadline")
	}
}

// H-Q2: the write is time-bounded: a client that stops reading loses the
// connection instead of holding the handler (and its buffer) forever.
func TestV1PieceContentWriteBounded(t *testing.T) {
	_, conn, size := stalledPieceRead(t)
	time.Sleep(2500 * time.Millisecond) // past the 1 s write deadline
	conn.SetReadDeadline(time.Now().Add(20 * time.Second))
	got, _ := io.Copy(io.Discard, conn)
	if got >= int64(size) {
		t.Fatalf("a client that stopped reading still got the whole piece (%d bytes): the write was not bounded", got)
	}
}
