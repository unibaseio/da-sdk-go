package sdk

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

// fakeStream answers /v1/upload with fakeEncode of what it received, passed
// through lie so a test can make it dishonest.
func fakeStream(t *testing.T, lie func(ff *types.FileFull)) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, _, err := r.FormFile("file")
		var data []byte
		if err == nil {
			data, err = io.ReadAll(f)
		}
		if err != nil {
			// still answer: a client must not take this as a good upload
			data = nil
		}
		pol := types.Policy{N: 6, K: 4}
		var ff types.FileFull
		if len(data) == 0 {
			ff = types.FileFull{FileReceipt: types.FileReceipt{FileCore: types.FileCore{Policy: pol, Hash: strings.Repeat("00", 32)}}}
		} else {
			ff = fakeEncode(t, data, pol, testStream)
		}
		if lie != nil {
			lie(&ff)
		}
		json.NewEncoder(w).Encode(ff)
	}))
}

func TestUploadDataChecksStreamAnswer(t *testing.T) {
	pol := types.Policy{N: 6, K: 4}
	data := testData()
	fp := writeTemp(t, data)

	cases := []struct {
		name    string
		lie     func(ff *types.FileFull)
		wantErr bool
	}{
		{"honest", nil, false},
		{"blank hash", func(ff *types.FileFull) { ff.Hash = "" }, true},
		{"other hash", func(ff *types.FileFull) { ff.Hash = strings.Repeat("ab", 32) }, true},
		{"other policy", func(ff *types.FileFull) { ff.Policy = types.Policy{N: 14, K: 7} }, true},
		{"prefix only", func(ff *types.FileFull) { ff.Size, ff.PieceSizes[0] = 1000, 1000 }, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := fakeStream(t, c.lie)
			defer s.Close()
			ff, err := UploadData(s.URL, types.Auth{}, pol, fp)
			if (err != nil) != c.wantErr {
				t.Fatalf("err=%v, wantErr=%v", err, c.wantErr)
			}
			if err == nil {
				if _, err := CheckFileFullPolicy(ff, testStream, fp, pol); err != nil {
					t.Fatalf("honest answer fails the full check: %v", err)
				}
			}
		})
	}
}

func TestUploadDataReportsReadError(t *testing.T) {
	s := fakeStream(t, nil)
	defer s.Close()
	missing := filepath.Join(t.TempDir(), "nope")
	if _, err := UploadData(s.URL, types.Auth{}, types.Policy{N: 6, K: 4}, missing); err == nil {
		t.Fatal("upload of an unreadable file reported success")
	}
}
