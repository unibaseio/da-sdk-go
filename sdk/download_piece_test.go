package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/bls/erasure"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// rsNet is a gateway plus one stream serving a single 6/4 piece with real
// Reed-Solomon parity, so DownloadPiece can rebuild missing data shards.
// The piece is 248 bytes: K=4 shards of 2 field elements (64 bytes each).
type rsNet struct {
	gw, st *httptest.Server
	data   []byte
	names  []string // honest replica name per slot
	mu     sync.Mutex
	shards map[string][]byte // name -> bytes the stores/stream serve
	gwRec  types.PieceReceipt
	stRec  types.PieceReceipt
}

var rsStore = common.HexToAddress("0x5")
var rsStreamer = common.HexToAddress("0x9")

func rsName(i int) string { return fmt.Sprintf("%096x", i+1) }

func newRSNet(t *testing.T) *rsNet {
	t.Helper()
	const n, k, slen = 6, 4, 2
	pol := types.Policy{N: n, K: k}
	n6 := &rsNet{shards: map[string][]byte{}}
	n6.data = make([]byte, 31*slen*k)
	for i := range n6.data {
		n6.data[i] = byte(i*7 + 3)
	}
	shards := make([][]byte, n)
	for j := 0; j < k; j++ {
		shards[j] = bls.Pad(n6.data[j*31*slen : (j+1)*31*slen])
	}
	rs, err := erasure.NewRS(n, k)
	if err != nil {
		t.Fatal(err)
	}
	need := []int{4, 5}
	enc := make([][]byte, k)
	for e := 0; e < slen; e++ {
		for j := 0; j < k; j++ {
			enc[j] = shards[j][e*bls.PadSize : (e+1)*bls.PadSize]
		}
		par, err := rs.Encode(enc, need)
		if err != nil {
			t.Fatal(err)
		}
		for j, v := range need {
			shards[v] = append(shards[v], par[j]...)
		}
	}
	core := types.PieceCore{Name: "p", Policy: pol, Size: int64(len(n6.data)), Streamer: rsStreamer}
	n6.gwRec = types.PieceReceipt{PieceCore: core, Replicas: make([]string, n), StoredOn: make([]common.Address, n)}
	n6.stRec = types.PieceReceipt{PieceCore: core, Replicas: make([]string, n), StoredOn: make([]common.Address, n)}
	for i := 0; i < n; i++ {
		name := rsName(i)
		n6.names = append(n6.names, name)
		n6.shards[name] = shards[i]
		n6.gwRec.Replicas[i], n6.gwRec.StoredOn[i] = name, rsStore
		n6.stRec.Replicas[i], n6.stRec.StoredOn[i] = name, rsStore
	}

	n6.st = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n6.mu.Lock()
		defer n6.mu.Unlock()
		switch {
		case r.URL.Path == "/v1/pieces/p":
			json.NewEncoder(w).Encode(n6.stRec)
		case r.URL.Path == "/v1/download":
			r.ParseForm()
			b, ok := n6.shards[r.Form.Get("name")]
			if !ok {
				http.NotFound(w, r)
				return
			}
			w.Write(b)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(n6.st.Close)
	edge := types.EdgeReceipt{EdgeMeta: types.EdgeMeta{
		Name: rsStreamer, Type: types.StreamType, ExposeURL: n6.st.URL, ChainType: chaintype,
	}, OnChain: true}
	n6.gw = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n6.mu.Lock()
		defer n6.mu.Unlock()
		switch {
		case r.URL.Path == "/v1/pieces/p":
			json.NewEncoder(w).Encode(n6.gwRec)
		case r.URL.Path == "/v1/edges":
			json.NewEncoder(w).Encode(map[string]any{"items": []types.EdgeReceipt{edge}})
		case strings.HasPrefix(r.URL.Path, "/v1/edges/"):
			json.NewEncoder(w).Encode(edge)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(n6.gw.Close)
	return n6
}

func (n *rsNet) set(f func(n *rsNet)) {
	n.mu.Lock()
	defer n.mu.Unlock()
	f(n)
}

func TestDownloadPieceHonest(t *testing.T) {
	n := newRSNet(t)
	_, data, err := DownloadPiece(n.gw.URL, types.Auth{}, "p")
	if err != nil || !bytes.Equal(data, n.data) {
		t.Fatalf("honest piece: err=%v equal=%v", err, bytes.Equal(data, n.data))
	}
}

// S5: the shard length comes from the gateway's Size and K, not from whichever
// replica answers first. A store answering slot 0 with a longer, still
// PadSize-aligned blob must be skipped (and the slot rebuilt), not make every
// honest shard look malformed.
func TestDownloadPieceShardLengthFromSize(t *testing.T) {
	n := newRSNet(t)
	n.set(func(n *rsNet) { n.shards[n.names[0]] = bytes.Repeat([]byte{0x01}, 3*bls.PadSize) })
	_, data, err := DownloadPiece(n.gw.URL, types.Auth{}, "p")
	if err != nil {
		t.Fatalf("one wrong-length store failed the download: %v", err)
	}
	if !bytes.Equal(data, n.data) {
		t.Fatal("rebuilt piece differs")
	}
}

func TestDownloadPieceRejectsBadSize(t *testing.T) {
	for _, size := range []int64{0, -1, MaxPieceSize(types.Policy{N: 6, K: 4}) + 1} {
		n := newRSNet(t)
		n.set(func(n *rsNet) { n.gwRec.Size = size })
		if _, _, err := DownloadPiece(n.gw.URL, types.Auth{}, "p"); err == nil {
			t.Fatalf("size %d accepted", size)
		}
	}
}
