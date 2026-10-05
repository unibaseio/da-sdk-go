package hub

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

// A conversation read on a non-home shard fetches each staged record from the
// owner's home shard by (owner, bucket, key) — never by bare name — and caches
// it under the owner + row location; a peer answer of the wrong size is
// dropped, not cached (review 2026-10-04 M).
func TestConversationFallbackOwnerBound(t *testing.T) {
	s := newV1TestServer(t)
	if err := s.gdb.AutoMigrate(&types.Conversation{}); err != nil {
		t.Fatal(err)
	}
	s.readCache = newReadCache()

	var gotPaths []string
	peer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPaths = append(gotPaths, r.URL.EscapedPath()+"?"+r.URL.RawQuery)
		if r.Header.Get(shardFwdHeader) != testFwdSecret {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		switch {
		case strings.Contains(r.URL.Path, "/objects/chat_0/"):
			io.WriteString(w, "hello")
		case strings.Contains(r.URL.Path, "/objects/chat_1/"):
			io.WriteString(w, "some other, longer object")
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer peer.Close()
	s.shard = testShardRouter(t, 0, 2, "http://127.0.0.1:1,"+peer.URL)
	owner := ownerHomedAt(t, s.shard, 1)

	// no network in tests: the DA file-record fallback reports a miss
	orig := fetchOwnerFile
	fetchOwnerFile = func(*Server, string, common.Address, io.Writer) error { return errors.New("no record") }
	defer func() { fetchOwnerFile = orig }()

	s.gdb.Create(&types.Conversation{Name: "chat", Owner: owner, Bucket: "b/x"})
	s.gdb.Create(&types.Needle{Owner: owner, Bucket: "b/x", Name: "chat_0", File: 3, Start: 0, Size: 5})
	s.gdb.Create(&types.Needle{Owner: owner, Bucket: "b/x", Name: "chat_1", File: 3, Start: 8, Size: 4})

	msgs, err := s.getConversation(t.Context(), "chat", owner, "", 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	if b, _ := json.Marshal(msgs); string(b) != `["hello"]` {
		t.Fatalf("messages = %s", b)
	}
	want := "/v1/buckets/b%2Fx/objects/chat_0/content?owner=" + owner
	if len(gotPaths) == 0 || gotPaths[0] != want {
		t.Fatalf("peer request = %v, want %s", gotPaths, want)
	}
	if v, ok := s.readCache.get(owner, locKey(3, 0, 5)); !ok || string(v) != "hello" {
		t.Fatal("fetched record not cached under owner + location")
	}
	if _, ok := s.readCache.get(owner, locKey(3, 8, 4)); ok {
		t.Fatal("wrong-size answer was cached")
	}
	if _, ok := s.readCache.get(owner, "chat_0"); ok {
		t.Fatal("cached under the bare key")
	}
}
