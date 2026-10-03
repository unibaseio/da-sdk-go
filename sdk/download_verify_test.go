package sdk

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// cachedPieces serves pieces from memory (the hub's piece store).
type cachedPieces struct {
	types.IPieceStore
	data map[string][]byte
}

func (c cachedPieces) GetPiece(_ context.Context, name string, w io.Writer, _ types.Options) (types.PieceReceipt, error) {
	b, ok := c.data[name]
	if !ok {
		return types.PieceReceipt{}, fmt.Errorf("miss")
	}
	w.Write(b)
	return types.PieceReceipt{}, nil
}

// gatewayWith serves one file receipt and, like a gateway predating owner
// filtering, ignores ?owner=.
func gatewayWith(t *testing.T, fr types.FileReceipt) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/v1/files/") {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(fr)
	}))
}

func TestGetFileReceiptOfChecksOwner(t *testing.T) {
	hub := common.HexToAddress("0x01")
	attacker := common.HexToAddress("0x02")
	gw := gatewayWith(t, types.FileReceipt{FileCore: types.FileCore{Name: "v", Owner: attacker}})
	defer gw.Close()

	if _, err := GetFileReceiptOf(gw.URL, "v", hub); err == nil {
		t.Fatal("accepted another owner's record for the hub's file")
	}
	if fr, err := GetFileReceiptOf(gw.URL, "v", attacker); err != nil || fr.Owner != attacker {
		t.Fatalf("owner's own record: %v %v", fr.Owner, err)
	}
	if _, err := GetFileReceiptOf(gw.URL, "v", types.EmptyAddr); err != nil {
		t.Fatalf("any-owner lookup: %v", err)
	}
}

func TestDownloadChecksFileHash(t *testing.T) {
	owner := common.HexToAddress("0x01")
	parts := map[string][]byte{"p1": []byte("hello "), "p2": []byte("world")}
	good := sha256.Sum256([]byte("hello world"))

	for _, c := range []struct {
		name    string
		hash    string
		wantErr bool
	}{
		{"matching hash", hex.EncodeToString(good[:]), false},
		{"forged data", strings.Repeat("ab", 32), true},
	} {
		gw := gatewayWith(t, types.FileReceipt{
			FileCore: types.FileCore{Name: "f", Hash: c.hash, Owner: owner},
			Pieces:   []string{"p1", "p2"},
		})
		var out bytes.Buffer
		err := DownloadOf(gw.URL, types.Auth{}, "f", owner, cachedPieces{data: parts}, &out)
		gw.Close()
		if (err != nil) != c.wantErr {
			t.Fatalf("%s: err=%v, wantErr=%v", c.name, err, c.wantErr)
		}
		if out.String() != "hello world" {
			t.Fatalf("%s: wrote %q", c.name, out.String())
		}
	}
}

func TestRelayStreams(t *testing.T) {
	e := func(name string, chain string, on bool) types.EdgeReceipt {
		return types.EdgeReceipt{EdgeMeta: types.EdgeMeta{Name: common.HexToAddress(name), ChainType: chain}, OnChain: on}
	}
	edges := []types.EdgeReceipt{e("0x1", "base-sepolia", true), e("0x2", "base-sepolia", false), e("0x3", "bnb-testnet-dao", true)}
	if got := relayStreams(edges, "base-sepolia"); len(got) != 1 || got[0].Name != common.HexToAddress("0x1") {
		t.Fatalf("got %v", got)
	}
	if got := relayStreams(edges, ""); len(got) != 2 {
		t.Fatalf("unknown chain: got %d streams, want the 2 on-chain ones", len(got))
	}
}
