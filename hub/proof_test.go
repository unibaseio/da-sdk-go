package hub

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/kv"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// /proof (and the object receipt) hash exactly the (owner, bucket, key) row's
// object: a later same-key write in another bucket of the same owner must not
// change the bundle (review 2026-10-04 M, like B4 for /content).
func TestV1ProofHashesTheRowsObject(t *testing.T) {
	s := newV1TestServer(t)
	s.readCache = newReadCache()

	dir := t.TempDir()
	ds, err := kv.NewBadgerStore(filepath.Join(dir, "kv"), &kv.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ds.Close() })

	alice := strings.ToLower("0x00000000000000000000000000000000000000a1")
	fs, err := logfs.New(ds, filepath.Join(dir, "logfs"), "0xhub", alice)
	if err != nil {
		t.Fatal(err)
	}
	s.lfs.Store(alice, fs)
	put := func(bucket, content string) {
		t.Helper()
		if err := fs.Put([]byte("k"), []byte(content)); err != nil {
			t.Fatal(err)
		}
		lm, _ := fs.GetMeta([]byte("k"))
		s.gdb.Create(&types.Needle{Owner: alice, Bucket: bucket, Name: "k", File: lm.Index, Start: lm.Start, Size: lm.Size})
		s.gdb.Create(&types.Volume{Owner: alice, File: lm.Index, Piece: "piece"})
	}
	s.gdb.Create(&types.Bucket{Name: "notes", Owner: alice, Kind: "memory"})
	s.gdb.Create(&types.Bucket{Name: "drafts", Owner: alice, Kind: "memory"})
	put("notes", "alice notes")
	put("drafts", "alice drafts") // same owner + key, other bucket, written last

	for bucket, content := range map[string]string{"notes": "alice notes", "drafts": "alice drafts"} {
		w := do(t, s, "GET", "/v1/buckets/"+bucket+"/objects/k/proof?owner="+alice, "", "")
		if w.Code != http.StatusOK {
			t.Fatalf("%s proof: %d %s", bucket, w.Code, w.Body.String())
		}
		var resp struct {
			Sha256    string `json:"sha256"`
			Keccak256 string `json:"keccak256"`
		}
		json.Unmarshal(w.Body.Bytes(), &resp)
		sum := sha256.Sum256([]byte(content))
		if resp.Sha256 != hex.EncodeToString(sum[:]) || resp.Keccak256 != keccakHex([]byte(content)) {
			t.Errorf("%s proof hashes another object: %+v", bucket, resp)
		}
	}

	// the receipt's sha256 comes from the LogFS meta, which now describes the
	// drafts object: notes gets none rather than the wrong one
	var rc v1Receipt
	w := do(t, s, "GET", "/v1/buckets/notes/objects/k?owner="+alice, "", "")
	json.Unmarshal(w.Body.Bytes(), &rc)
	if rc.Sha256 != "" {
		t.Errorf("notes receipt carries another object's sha256: %s", rc.Sha256)
	}
	w = do(t, s, "GET", "/v1/buckets/drafts/objects/k?owner="+alice, "", "")
	json.Unmarshal(w.Body.Bytes(), &rc)
	if sum := sha256.Sum256([]byte("alice drafts")); rc.Sha256 != hex.EncodeToString(sum[:]) {
		t.Errorf("drafts receipt sha256 = %q", rc.Sha256)
	}
}
