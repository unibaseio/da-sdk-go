package evidence

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

type slot struct {
	node   common.Address
	fake   bool
	active bool
}

type fakeChain struct {
	pieces map[string]uint64
	n, k   uint8
	expire uint64
	epoch  uint64
	slots  []slot
}

func (f *fakeChain) GetPieceSerial(p string) (uint64, error) { return f.pieces[p], nil }
func (f *fakeChain) GetPieceRS(uint64) (uint8, uint8, uint64, error) {
	return f.n, f.k, f.expire, nil
}
func (f *fakeChain) GetPieceReplica(_ uint64, s uint8) (uint64, common.Address, error) {
	if int(s) >= len(f.slots) {
		return 0, common.Address{}, nil
	}
	return uint64(s) + 1, f.slots[s].node, nil
}
func (f *fakeChain) GetRSFake(_ uint64, s uint8) (bool, error) {
	if int(s) >= len(f.slots) {
		return false, nil
	}
	return f.slots[s].fake, nil
}
func (f *fakeChain) NodeActive(a common.Address) (bool, error) {
	for _, s := range f.slots {
		if s.node == a {
			return s.active, nil
		}
	}
	return false, nil
}
func (f *fakeChain) CurrentEpoch() (uint64, error) { return f.epoch, nil }

func healthyChain() *fakeChain {
	c := &fakeChain{pieces: map[string]uint64{"piece-a": 7}, n: 6, k: 4, expire: 100, epoch: 50}
	for i := 0; i < 6; i++ {
		c.slots = append(c.slots, slot{node: common.BigToAddress(bigInt(i + 1)), active: true})
	}
	return c
}

// hub serves one evidence object and its /proof bundle.
type hub struct {
	content     []byte
	proofStatus int
	proof       Proof
}

func (h *hub) server(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/proof"):
			if h.proofStatus != 0 && h.proofStatus != http.StatusOK {
				w.WriteHeader(h.proofStatus)
				return
			}
			json.NewEncoder(w).Encode(h.proof)
		case strings.HasPrefix(r.URL.Path, "/v1/buckets/"):
			w.Write(h.content)
		default:
			http.NotFound(w, r)
		}
	}))
}

func newHub(content []byte) *hub {
	h := &hub{content: content}
	h.proof.Commitment = "piece-a"
	h.proof.Keccak256 = keccakHex(content)
	h.proof.Range = &Range{Volume: 0, Start: 4096, Size: uint64(len(content))}
	return h
}

func objURL(base string) string {
	return base + "/v1/buckets/8004-evidence/objects/0xabc?owner=0x01"
}

func TestVerifyHealthy(t *testing.T) {
	content := []byte(`{"agentId":1,"value":95}`)
	srv := newHub(content).server(t)
	defer srv.Close()

	r, err := Verify(objURL(srv.URL), keccakHex(content), healthyChain(), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if !r.ContentOK || !r.Committed || !r.Registered || !r.Available || r.LiveReplicas != 6 {
		t.Fatalf("healthy evidence should pass every check: %+v", r)
	}
	if r.PieceIndex != 7 || r.N != 6 || r.K != 4 {
		t.Fatalf("piece details wrong: %+v", r)
	}
	// without --trustless nothing ties the piece to the content: say so
	if r.PieceBound || !hasNote(r, NotPieceBoundNote) {
		t.Fatalf("non-trustless result must not claim the piece holds the content: %+v", r)
	}
}

func hasNote(r Result, note string) bool {
	for _, n := range r.Notes {
		if n == note {
			return true
		}
	}
	return false
}

func TestVerifyTamperedContent(t *testing.T) {
	srv := newHub([]byte(`{"value":1}`)).server(t)
	defer srv.Close()

	r, err := Verify(objURL(srv.URL), keccakHex([]byte(`{"value":95}`)), healthyChain(), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if r.ContentOK {
		t.Fatal("changed content must fail the content check")
	}
}

func TestVerifyExpired(t *testing.T) {
	content := []byte("x")
	srv := newHub(content).server(t)
	defer srv.Close()
	c := healthyChain()
	c.epoch = c.expire

	r, _ := Verify(objURL(srv.URL), keccakHex(content), c, Options{})
	if r.Available {
		t.Fatalf("expired piece must not be available: %+v", r)
	}
}

func TestVerifyReplicaHealth(t *testing.T) {
	content := []byte("x")
	srv := newHub(content).server(t)
	defer srv.Close()

	// 6/4: forge 2, one empty slot, one inactive node → 2 healthy < 4
	c := healthyChain()
	c.slots[0].fake = true
	c.slots[1].fake = true
	c.slots[2].node = common.Address{}
	c.slots[3].active = false

	r, _ := Verify(objURL(srv.URL), keccakHex(content), c, Options{})
	if r.LiveReplicas != 2 || r.Available {
		t.Fatalf("want 2 healthy replicas and unavailable, got %+v", r)
	}

	// exactly k healthy (N-K forged) is still enough to rebuild
	c = healthyChain()
	c.slots[0].fake = true
	c.slots[1].fake = true
	r, _ = Verify(objURL(srv.URL), keccakHex(content), c, Options{})
	if r.LiveReplicas != 4 || !r.Available {
		t.Fatalf("k healthy replicas should be available, got %+v", r)
	}
}

func TestVerifyStaged(t *testing.T) {
	content := []byte("x")
	h := newHub(content)
	h.proofStatus = http.StatusTooEarly
	srv := h.server(t)
	defer srv.Close()

	r, err := Verify(objURL(srv.URL), keccakHex(content), healthyChain(), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if !r.ContentOK || !r.Staged || r.Committed || r.Available {
		t.Fatalf("staged object: content ok, not yet committed: %+v", r)
	}
}

func TestVerifyUnregisteredPiece(t *testing.T) {
	content := []byte("x")
	srv := newHub(content).server(t)
	defer srv.Close()
	c := healthyChain()
	c.pieces = map[string]uint64{}

	r, _ := Verify(objURL(srv.URL), keccakHex(content), c, Options{})
	if r.Registered || r.Available {
		t.Fatalf("unregistered piece must fail: %+v", r)
	}
}

func TestVerifyTrustless(t *testing.T) {
	content := []byte(`{"agentId":1}`)
	srv := newHub(content).server(t)
	defer srv.Close()

	piece := make([]byte, 8192)
	copy(piece[4096:], content)
	good := func(string) ([]byte, error) { return piece, nil }
	r, err := Verify(objURL(srv.URL), keccakHex(content), healthyChain(), Options{FetchPiece: good})
	if err != nil {
		t.Fatal(err)
	}
	if r.Trustless == nil || !r.Trustless.OK {
		t.Fatalf("rebuilt piece holds the evidence at range: %+v", r.Trustless)
	}
	if !r.PieceBound || hasNote(r, NotPieceBoundNote) {
		t.Fatalf("trustless pass binds the piece: %+v", r)
	}

	// the hub served the right bytes, but the piece on store nodes differs
	bad := make([]byte, 8192)
	r, _ = Verify(objURL(srv.URL), keccakHex(content), healthyChain(), Options{FetchPiece: func(string) ([]byte, error) { return bad, nil }})
	if r.Trustless == nil || r.Trustless.OK {
		t.Fatalf("mismatching piece must fail the trustless check: %+v", r.Trustless)
	}
	if r.PieceBound {
		t.Fatal("failed trustless check reported as piece-bound")
	}

	// range past the end of the piece is reported, not a panic
	short := make([]byte, 100)
	r, _ = Verify(objURL(srv.URL), keccakHex(content), healthyChain(), Options{FetchPiece: func(string) ([]byte, error) { return short, nil }})
	if r.Trustless == nil || r.Trustless.OK {
		t.Fatalf("out-of-range must fail: %+v", r.Trustless)
	}
}

func TestVerifyNonHubURL(t *testing.T) {
	content := []byte("x")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write(content) }))
	defer srv.Close()

	r, err := Verify(srv.URL+"/agent.json", keccakHex(content), healthyChain(), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if !r.ContentOK || r.Committed || len(r.Notes) == 0 {
		t.Fatalf("plain https evidence: content checked, DA skipped with a note: %+v", r)
	}
}

func TestNormHash(t *testing.T) {
	h := keccakHex([]byte("x"))
	for _, in := range []string{h, strings.ToUpper(h[2:]), " " + h + " "} {
		if got, err := normHash(in); err != nil || got != h {
			t.Fatalf("normHash(%q) = %q, %v", in, got, err)
		}
	}
	if _, err := normHash("0x1234"); err == nil {
		t.Fatal("short hash must be rejected")
	}
}

func TestProofURLFor(t *testing.T) {
	got, ok := proofURLFor("https://hub.example/v1/buckets/b/objects/0xabc?owner=0x01")
	if !ok || got != "https://hub.example/v1/buckets/b/objects/0xabc/proof?owner=0x01" {
		t.Fatalf("proofURLFor = %q %v", got, ok)
	}
	for _, bad := range []string{"ipfs://cid", "https://x/agent.json", "https://x/v1/buckets/b/objects/k/proof"} {
		if _, ok := proofURLFor(bad); ok {
			t.Fatalf("%s is not a hub object URL", bad)
		}
	}
}

func bigInt(i int) *big.Int { return big.NewInt(int64(i)) }
