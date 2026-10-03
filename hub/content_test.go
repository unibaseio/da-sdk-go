package hub

import (
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/kv"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// Object content is the bytes of the exact (owner, bucket, key) row — never
// another bucket's or another owner's object that shares the key (B4).
func TestV1ContentReadsTheRowsOwnBytes(t *testing.T) {
	s := newV1TestServer(t)
	s.readCache = newReadCache()

	dir := t.TempDir()
	ds, err := kv.NewBadgerStore(filepath.Join(dir, "kv"), &kv.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ds.Close() })

	alice := strings.ToLower("0x00000000000000000000000000000000000000a1")
	bob := strings.ToLower("0x00000000000000000000000000000000000000b2")

	put := func(owner, bucket, key, content string) {
		t.Helper()
		v, ok := s.lfs.Load(owner)
		if !ok {
			fs, err := logfs.New(ds, filepath.Join(dir, "logfs"), "0xhub", owner)
			if err != nil {
				t.Fatal(err)
			}
			s.lfs.Store(owner, fs)
			v = fs
		}
		fs := v.(*logfs.LogFS)
		if err := fs.Put([]byte(key), []byte(content)); err != nil {
			t.Fatal(err)
		}
		lm, err := fs.GetMeta([]byte(key))
		if err != nil {
			t.Fatal(err)
		}
		s.gdb.Create(&types.Needle{Owner: owner, Bucket: bucket, Name: key, File: lm.Index, Start: lm.Start, Size: lm.Size})
	}

	s.gdb.Create(&types.Bucket{Name: "notes", Owner: alice, Kind: "file"})
	s.gdb.Create(&types.Bucket{Name: "drafts", Owner: alice, Kind: "file"})
	s.gdb.Create(&types.Bucket{Name: "bobs", Owner: bob, Kind: "file"})
	put(alice, "notes", "todo.txt", "alice notes")
	put(alice, "drafts", "todo.txt", "alice drafts") // same owner, same key, other bucket
	put(bob, "bobs", "todo.txt", "bob's text")      // other owner, same key, written last

	for _, c := range []struct{ path, want string }{
		{"/v1/buckets/notes/objects/todo.txt/content?owner=" + alice, "alice notes"},
		{"/v1/buckets/drafts/objects/todo.txt/content?owner=" + alice, "alice drafts"},
		{"/v1/buckets/notes/objects/todo.txt/content", "alice notes"}, // no owner: the bucket decides
		{"/v1/buckets/bobs/objects/todo.txt/content", "bob's text"},
	} {
		w := do(t, s, "GET", c.path, "", "")
		if w.Code != http.StatusOK || w.Body.String() != c.want {
			t.Errorf("%s: got %d %q, want %q", c.path, w.Code, w.Body.String(), c.want)
		}
	}

	// a key that exists only in another bucket is not found here
	if w := do(t, s, "GET", "/v1/buckets/bobs/objects/missing/content", "", ""); w.Code != http.StatusNotFound {
		t.Errorf("missing object: got %d", w.Code)
	}
}
