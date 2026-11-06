package sdk

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// M1/F9: the registration's signature covers the registration itself.
func TestRegisterEdgeSignedBindsTheRegistration(t *testing.T) {
	sk, _ := crypto.GenerateKey()
	var got types.Auth
	var txMsg []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		au, err := DecodeAuth(r.Header.Get("Authorization"))
		if err != nil {
			t.Error(err)
		}
		got = au
		b := new(bytes.Buffer)
		b.ReadFrom(r.Body)
		form, _ := url.ParseQuery(b.String())
		txMsg, _ = hex.DecodeString(form.Get("txMsg"))
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()
	em := types.EdgeMeta{Type: types.StoreType, Name: crypto.PubkeyToAddress(sk.PublicKey), ExposeURL: "http://10.0.0.1:8082", ChainType: "base-sepolia"}
	sign := func(h []byte) (types.Auth, error) { return key.BuildAuth(sk, h) }
	if err := RegisterEdgeSigned(srv.URL, sign, em); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got.Hash, EdgeAuthHash(txMsg)) {
		t.Fatal("the signature does not cover the registration")
	}
	if err := VerifyAuth(got); err != nil {
		t.Fatal(err)
	}
	// another registration does not match the captured signature
	em.ExposeURL = "http://203.0.113.9:8082"
	other, _ := json.Marshal(em)
	if bytes.Equal(got.Hash, EdgeAuthHash(other)) {
		t.Fatal("a captured signature covers another ExposeURL")
	}
}
