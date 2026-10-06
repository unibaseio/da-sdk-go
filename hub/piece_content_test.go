package hub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

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
