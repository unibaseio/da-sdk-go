package hub

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

func TestLikeEscape(t *testing.T) {
	if got := likeEscape(`a%b_c\d`); got != `a\%b\_c\\d` {
		t.Fatalf("got %q", got)
	}
}

// GET /v1/conversations/{id}: one owner's conversation only, the id matched
// literally, and bounded paging.
func TestV1GetConversationBounds(t *testing.T) {
	s := newV1TestServer(t)
	if err := s.gdb.AutoMigrate(&types.Conversation{}); err != nil {
		t.Fatal(err)
	}
	alice := strings.ToLower("0x00000000000000000000000000000000000000a1")
	s.gdb.Create(&types.Conversation{Name: "%", Owner: alice, Bucket: "b"})

	// no owner: refused instead of spanning every account
	if w := do(t, s, "GET", "/v1/conversations/%25", "", ""); w.Code != http.StatusBadRequest {
		t.Fatalf("ownerless read: got %d %s", w.Code, w.Body.String())
	}
	// negative paging refused (a negative length removed the SQL LIMIT)
	if w := do(t, s, "GET", "/v1/conversations/x?owner="+alice+"&length=-1", "", ""); w.Code != http.StatusBadRequest {
		t.Fatalf("negative length: got %d", w.Code)
	}

	// an id of "%" matches only a needle literally named "%_<n>", not every row
	s.gdb.Create(&types.Needle{Owner: alice, Bucket: "b", Name: "secret_0", File: 0, Start: 0, Size: 1})
	w := do(t, s, "GET", "/v1/conversations/%25?owner="+alice, "", "")
	if w.Code != http.StatusOK {
		t.Fatalf("got %d %s", w.Code, w.Body.String())
	}
	var resp struct {
		Messages []string `json:"messages"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if len(resp.Messages) != 0 {
		t.Fatalf("id %%%% matched other conversations: %v", resp.Messages)
	}
}

func TestV1StatsDaysCapped(t *testing.T) {
	s := newV1TestServer(t)
	s.statManager = &StatManager{stats: map[string]*types.Stat{}}
	if w := do(t, s, "GET", "/v1/stats?days=1000000000000", "", ""); w.Code != http.StatusOK {
		t.Fatalf("got %d", w.Code)
	}
}
