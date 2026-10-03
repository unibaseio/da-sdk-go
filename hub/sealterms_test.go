package hub

import (
	"math/big"
	"testing"

	com "github.com/unibaseio/da-sdk-go/contract/common"
)

func TestSealTerms(t *testing.T) {
	const start, max = 100, 1201
	def := uint64(start + com.DefaultStoreEpoch)
	defPrice := big.NewInt(int64(com.DefaultReplicaPrice))
	huge := "1000000000000000000000000"

	for _, c := range []struct {
		name          string
		register      string
		expire, price string
		wantExpire    uint64
		wantPrice     *big.Int
	}{
		{"hub, defaults", "hub", "", "", def, defPrice},
		{"hub ignores a client price", "hub", "", huge, def, defPrice},
		{"hub caps a long term", "hub", "999999999", "", start + max, defPrice},
		{"hub keeps a shorter term", "hub", "500", "", 500, defPrice},
		{"bad expire falls back", "hub", "50", "", def, defPrice},
		{"client pays: its terms stand", "client", "999999999", huge, 999999999, func() *big.Int { p, _ := new(big.Int).SetString(huge, 10); return p }()},
		{"client: non-positive price ignored", "client", "", "0", def, defPrice},
		// hub_attributed only attributes ownership: the hub still pays
		{"hub_attributed ignores a client price", "hub_attributed", "", huge, def, defPrice},
		{"hub_attributed caps a long term", "hub_attributed", "999999999", "", start + max, defPrice},
	} {
		e, p := sealTerms(c.register, start, c.expire, c.price, max)
		if e != c.wantExpire || p.Cmp(c.wantPrice) != 0 {
			t.Errorf("%s: got (%d, %s), want (%d, %s)", c.name, e, p, c.wantExpire, c.wantPrice)
		}
	}
}

func TestSealFileName(t *testing.T) {
	if got := sealFileName("0xAbC", ""); got != "" {
		t.Fatalf("no name: got %q", got)
	}
	// a client cannot take a hub volume's name ("<owner>/<i>.vol")
	if got := sealFileName("0xAbC", "0xvictim/7.vol"); got != "seal/0xabc/0xvictim/7.vol" {
		t.Fatalf("got %q", got)
	}
}
