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
// client that sends the request and then does not read. The server has one
// global piece slot and the given per-client cap (0 = none).
func stalledPieceRead(t *testing.T, perClient int) (*Server, *httptest.Server, net.Conn, int) {
	t.Helper()
	t.Setenv("HUB_PIECE_WRITE_BASE_SEC", "1")
	t.Setenv("HUB_PIECE_WRITE_MIN_BPS", "1099511627776") // the size adds ~nothing
	s := newV1TestServer(t)
	big := bytes.Repeat([]byte{0x5a}, 64<<20) // far more than socket buffers hold
	s.ps = &recPieces{data: map[string][]byte{"c1d": big}}
	s.pieceSem = make(chan struct{}, 1)
	s.pieceClients = newClientSlots(perClient)
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
	return s, srv, conn, len(big)
}

// getPiece reads c1d in full from srv, reporting the status and body size.
func getPiece(t *testing.T, srv *httptest.Server) (int, int) {
	t.Helper()
	resp, err := http.Get(srv.URL + "/v1/pieces/c1d/content")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	n, _ := io.Copy(io.Discard, resp.Body)
	return resp.StatusCode, int(n)
}

// H-Q2 (rework): a client that stops reading must not block other readers.
// With the global slot held through the write, the second read waited
// maxSlotWait and got 503 while the first client sat on its write deadline.
func TestV1PieceContentStalledReaderDoesNotBlockOthers(t *testing.T) {
	old := maxSlotWait
	maxSlotWait = 500 * time.Millisecond
	t.Cleanup(func() { maxSlotWait = old })
	t.Setenv("HUB_PIECE_WRITE_BASE_SEC", "30") // the stalled write outlives the test's wait
	s, srv, _, size := stalledPieceRead(t, 0)
	time.Sleep(300 * time.Millisecond) // the stalled handler is now blocked writing
	if len(s.pieceSem) != 0 {
		t.Fatal("global piece slot still held while the response is being written")
	}
	code, n := getPiece(t, srv)
	if code != http.StatusOK || n != size {
		t.Fatalf("second reader got %d with %d bytes while another client stalled; want 200 with %d", code, n, size)
	}
}

// H-Q2 (rework): one client's concurrent piece requests are capped, write
// included, and its slot comes back once the bounded write ends.
func TestV1PieceContentPerClientCap(t *testing.T) {
	s, srv, _, size := stalledPieceRead(t, 1)
	time.Sleep(300 * time.Millisecond)
	if code, _ := getPiece(t, srv); code != http.StatusTooManyRequests {
		t.Fatalf("second concurrent request from the same client got %d, want 429", code)
	}
	deadline := time.Now().Add(10 * time.Second) // the stalled write ends after ~1 s
	for {
		s.pieceClients.mu.Lock()
		held := len(s.pieceClients.n)
		s.pieceClients.mu.Unlock()
		if held == 0 {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("client slot still held long after the write deadline")
		}
		time.Sleep(50 * time.Millisecond)
	}
	if code, n := getPiece(t, srv); code != http.StatusOK || n != size {
		t.Fatalf("after the stalled write ended: %d with %d bytes, want 200 with %d", code, n, size)
	}
}

func TestClientSlots(t *testing.T) {
	cs := newClientSlots(2)
	r1, ok1 := cs.acquire("a")
	_, ok2 := cs.acquire("a")
	_, ok3 := cs.acquire("a")
	_, okB := cs.acquire("b")
	if !ok1 || !ok2 || ok3 || !okB {
		t.Fatalf("cap 2: got %v %v %v (b %v)", ok1, ok2, ok3, okB)
	}
	r1()
	r1() // a second release is a no-op
	if _, ok := cs.acquire("a"); !ok {
		t.Fatal("slot not returned on release")
	}
	if _, ok := cs.acquire("a"); ok {
		t.Fatal("double release freed two slots")
	}
}

// H-Q2: the write is time-bounded: a client that stops reading loses the
// connection instead of holding the handler (and its buffer) forever.
func TestV1PieceContentWriteBounded(t *testing.T) {
	_, _, conn, size := stalledPieceRead(t, 0)
	time.Sleep(2500 * time.Millisecond) // past the 1 s write deadline
	conn.SetReadDeadline(time.Now().Add(20 * time.Second))
	got, _ := io.Copy(io.Discard, conn)
	if got >= int64(size) {
		t.Fatalf("a client that stopped reading still got the whole piece (%d bytes): the write was not bounded", got)
	}
}
