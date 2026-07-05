package hub

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// newV1TestServer stands up an in-memory sqlite-backed Server with the /v1
// routes mounted, so the HTTP layer (routing + auth + kind + cursor + receipt)
// can be exercised in-process. Upload/content paths need logfs and are covered
// by the on-chain smoke, not here.
func newV1TestServer(t *testing.T) *Server {
	t.Helper()
	gin.SetMode(gin.TestMode)
	dsn := "file:" + t.Name() + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1)
	if err := db.AutoMigrate(&types.Bucket{}, &types.Needle{}, &types.Account{}, &types.Volume{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	s := &Server{
		Router:        gin.New(),
		gdb:           db,
		memStat:       &memStatCache{},
		bucketDisplay: make(map[string]types.BucketDisplay),
	}
	s.registV1()
	return s
}

// testKey returns a fresh (address 0x…, privkey-hex-without-0x) pair.
func testKey(t *testing.T) (string, string) {
	t.Helper()
	k, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	return crypto.PubkeyToAddress(k.PublicKey).Hex(), hex.EncodeToString(crypto.FromECDSA(k))
}

func authHeader(addr, privk string) string {
	au := sdk.BuildAuth(addr, privk, []byte("hub"))
	b, _ := json.Marshal(au)
	return hex.EncodeToString(b)
}

func do(t *testing.T, s *Server, method, path, auth, body string) *httptest.ResponseRecorder {
	t.Helper()
	var r *http.Request
	if body != "" {
		r = httptest.NewRequest(method, path, strings.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
	} else {
		r = httptest.NewRequest(method, path, nil)
	}
	if auth != "" {
		r.Header.Set("Authorization", auth)
	}
	w := httptest.NewRecorder()
	s.Router.ServeHTTP(w, r)
	return w
}

func TestV1Buckets(t *testing.T) {
	s := newV1TestServer(t)
	addr, pk := testKey(t)
	auth := authHeader(addr, pk)

	// unauthenticated write → 401
	if w := do(t, s, "PUT", "/v1/buckets/kb1", "", `{"kind":"knowledgebase"}`); w.Code != http.StatusUnauthorized {
		t.Fatalf("unauth PUT: got %d want 401", w.Code)
	}

	// create bucket with kind → 201
	if w := do(t, s, "PUT", "/v1/buckets/kb1", auth, `{"kind":"knowledgebase"}`); w.Code != http.StatusCreated {
		t.Fatalf("create bucket: got %d body %s", w.Code, w.Body.String())
	}

	// GET bucket → kind reflected
	w := do(t, s, "GET", "/v1/buckets/kb1", "", "")
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), `"kind":"knowledgebase"`) {
		t.Fatalf("get bucket: got %d body %s", w.Code, w.Body.String())
	}

	// kind is immutable → 409
	if w := do(t, s, "PUT", "/v1/buckets/kb1", auth, `{"kind":"model"}`); w.Code != http.StatusConflict {
		t.Fatalf("kind change: got %d want 409 body %s", w.Code, w.Body.String())
	}

	// idempotent same kind → 200
	if w := do(t, s, "PUT", "/v1/buckets/kb1", auth, `{"kind":"knowledgebase"}`); w.Code != http.StatusOK {
		t.Fatalf("idempotent PUT: got %d want 200", w.Code)
	}

	// list buckets is enumeration → signed request scoped to the signer's own
	w = do(t, s, "GET", "/v1/buckets?kind=knowledgebase", auth, "")
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "kb1") {
		t.Fatalf("list buckets: got %d body %s", w.Code, w.Body.String())
	}

	// invalid kind → 400
	if w := do(t, s, "PUT", "/v1/buckets/bad", auth, `{"kind":"nope"}`); w.Code != http.StatusBadRequest {
		t.Fatalf("bad kind: got %d want 400", w.Code)
	}
}

func TestV1Objects(t *testing.T) {
	s := newV1TestServer(t)
	signer, spk := testKey(t)
	owner := strings.ToLower(signer)
	ownerAuth := authHeader(signer, spk)

	// seed a bucket + two needles; one with an on-chain volume (committed), one without (staged)
	s.gdb.Create(&types.Bucket{Name: "repo", Owner: owner, Kind: "model"})
	s.gdb.Create(&types.Needle{Owner: owner, Bucket: "repo", Name: "a.bin", File: 10, Size: 100})
	s.gdb.Create(&types.Needle{Owner: owner, Bucket: "repo", Name: "b.bin", File: 11, Size: 200})
	s.gdb.Create(&types.Volume{Owner: owner, File: 10, Piece: "a11abaaf", TxHash: "0xdead", ChainType: "base-sepolia"})

	// list objects is enumeration → anonymous is rejected, signed owner works
	if w := do(t, s, "GET", "/v1/buckets/repo/objects?owner="+owner, "", ""); w.Code != http.StatusUnauthorized {
		t.Fatalf("anon list objects: got %d want 401", w.Code)
	}
	w := do(t, s, "GET", "/v1/buckets/repo/objects?owner="+owner, ownerAuth, "")
	if w.Code != http.StatusOK {
		t.Fatalf("list objects: got %d body %s", w.Code, w.Body.String())
	}
	var lst struct {
		Objects []v1Receipt `json:"objects"`
	}
	json.Unmarshal(w.Body.Bytes(), &lst)
	if len(lst.Objects) != 2 {
		t.Fatalf("list objects: got %d want 2", len(lst.Objects))
	}

	// get the committed object → status committed + commitment + chain
	w = do(t, s, "GET", "/v1/buckets/repo/objects/a.bin?owner="+owner, "", "")
	var rc v1Receipt
	json.Unmarshal(w.Body.Bytes(), &rc)
	if rc.Status != "committed" || rc.Commitment != "a11abaaf" || rc.Chain == nil || rc.Chain.TxHash != "0xdead" {
		t.Fatalf("committed receipt wrong: %+v", rc)
	}

	// get the staged object → status staged, no commitment (fresh var: staged
	// response omits commitment/chain, so don't decode into a reused struct)
	w = do(t, s, "GET", "/v1/buckets/repo/objects/b.bin?owner="+owner, "", "")
	var rc2 v1Receipt
	json.Unmarshal(w.Body.Bytes(), &rc2)
	if rc2.Status != "staged" || rc2.Commitment != "" {
		t.Fatalf("staged receipt wrong: %+v", rc2)
	}

	// missing object → 404
	if w := do(t, s, "GET", "/v1/buckets/repo/objects/none?owner="+owner, "", ""); w.Code != http.StatusNotFound {
		t.Fatalf("missing object: got %d want 404", w.Code)
	}

	// write guards (return before logfs): PUT to a non-existent bucket → 404;
	// PUT to a bucket owned by someone else → 403.
	addr, pk := testKey(t)
	auth := authHeader(addr, pk)
	if w := do(t, s, "PUT", "/v1/buckets/nobucket/objects/x", auth, "hello"); w.Code != http.StatusNotFound {
		t.Fatalf("PUT no-bucket: got %d want 404 body %s", w.Code, w.Body.String())
	}
	// "repo" is owned by `owner` (0xabc…001), not by our fresh signer → 403
	if w := do(t, s, "PUT", "/v1/buckets/repo/objects/x", auth, "hello"); w.Code != http.StatusForbidden {
		t.Fatalf("PUT wrong-owner: got %d want 403 body %s", w.Code, w.Body.String())
	}

	// proof: committed object → 200 + commitment; staged → 425
	w = do(t, s, "GET", "/v1/buckets/repo/objects/a.bin/proof?owner="+owner, "", "")
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "a11abaaf") {
		t.Fatalf("proof committed: got %d body %s", w.Code, w.Body.String())
	}
	if w := do(t, s, "GET", "/v1/buckets/repo/objects/b.bin/proof?owner="+owner, "", ""); w.Code != http.StatusTooEarly {
		t.Fatalf("proof staged: got %d want 425", w.Code)
	}
}

func TestV1DeleteBucket(t *testing.T) {
	s := newV1TestServer(t)
	addr, pk := testKey(t)
	auth := authHeader(addr, pk)

	if w := do(t, s, "PUT", "/v1/buckets/tmp", auth, `{"kind":"file"}`); w.Code != http.StatusCreated {
		t.Fatalf("create: %d", w.Code)
	}
	// delete someone else's / non-existent
	if w := do(t, s, "DELETE", "/v1/buckets/nope", auth, ""); w.Code != http.StatusNotFound {
		t.Fatalf("delete missing: got %d want 404", w.Code)
	}
	// delete own → 200, then gone → 404
	if w := do(t, s, "DELETE", "/v1/buckets/tmp", auth, ""); w.Code != http.StatusOK {
		t.Fatalf("delete own: got %d body %s", w.Code, w.Body.String())
	}
	if w := do(t, s, "GET", "/v1/buckets/tmp", "", ""); w.Code != http.StatusNotFound {
		t.Fatalf("get after delete: got %d want 404", w.Code)
	}
}

func TestV1BucketsCursor(t *testing.T) {
	s := newV1TestServer(t)
	signer, spk := testKey(t)
	owner := strings.ToLower(signer)
	auth := authHeader(signer, spk)
	s.gdb.Create(&types.Bucket{Name: "b1", Owner: owner, Kind: "file"})
	s.gdb.Create(&types.Bucket{Name: "b2", Owner: owner, Kind: "file"})
	s.gdb.Create(&types.Bucket{Name: "b3", Owner: owner, Kind: "file"})

	// bucket listing is enumeration → requires a signed owner request
	// page 1: limit=2 → 2 buckets + a nextCursor
	w := do(t, s, "GET", "/v1/buckets?owner="+owner+"&limit=2", auth, "")
	var p struct {
		Buckets    []types.Bucket `json:"buckets"`
		NextCursor string         `json:"nextCursor"`
	}
	json.Unmarshal(w.Body.Bytes(), &p)
	if len(p.Buckets) != 2 || p.NextCursor == "" {
		t.Fatalf("page1: got %d buckets, cursor %q", len(p.Buckets), p.NextCursor)
	}
	// page 2: use cursor → remaining 1, no further cursor
	w = do(t, s, "GET", "/v1/buckets?owner="+owner+"&limit=2&cursor="+p.NextCursor, auth, "")
	var p2 struct {
		Buckets    []types.Bucket `json:"buckets"`
		NextCursor string         `json:"nextCursor"`
	}
	json.Unmarshal(w.Body.Bytes(), &p2)
	if len(p2.Buckets) != 1 || p2.NextCursor != "" {
		t.Fatalf("page2: got %d buckets, cursor %q", len(p2.Buckets), p2.NextCursor)
	}
}
