package sdk

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

func TestGetMeteringPricing(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/metering/pricing" {
			t.Errorf("path = %s", r.URL.Path)
		}
		w.Write([]byte(`{"enabled":true,"charge_writes":true,"write_base_wei":"100","write_per_kb_wei":"10","read_per_request_wei":"0"}`))
	}))
	defer srv.Close()

	p, err := GetMeteringPricing(srv.URL)
	if err != nil {
		t.Fatalf("GetMeteringPricing: %v", err)
	}
	if !p.Enabled || p.WriteBaseWei != "100" || p.WritePerKBWei != "10" {
		t.Errorf("bad pricing: %+v", p)
	}
}

func TestGetMeteringUsage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("owner"); got != "0xabc" {
			t.Errorf("owner query = %s", got)
		}
		w.Write([]byte(`{"owner":"0xabc","total_writes":5,"unsettled_fee_wei":"500","total_fee_wei":"500","credit_limit_wei":"0"}`))
	}))
	defer srv.Close()

	u, err := GetMeteringUsage(srv.URL, "0xabc")
	if err != nil {
		t.Fatalf("GetMeteringUsage: %v", err)
	}
	if u.TotalWrites != 5 || u.UnsettledFeeWei != "500" {
		t.Errorf("bad usage: %+v", u)
	}
}

func TestCanWrite(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("owner") != "0xabc" || q.Get("bytes") != "2500" {
			t.Errorf("query = %v", q)
		}
		w.Write([]byte(`{"allowed":false,"reason":"credit_limit_exceeded","required_wei":"200","credit_limit_wei":"150"}`))
	}))
	defer srv.Close()

	res, err := CanWrite(srv.URL, "0xabc", 2500)
	if err != nil {
		t.Fatalf("CanWrite: %v", err)
	}
	if res.Allowed || res.Reason != "credit_limit_exceeded" || res.RequiredWei != "200" {
		t.Errorf("bad can-write: %+v", res)
	}
}

func TestSettleMeteringRequiresAuth(t *testing.T) {
	if _, err := SettleMetering("http://example.invalid", types.Auth{}); err == nil {
		t.Error("expected error for unsigned auth")
	}
}

func TestSettleMetering(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.Header.Get("Authorization") == "" {
			t.Error("missing Authorization header")
		}
		w.Write([]byte(`{"owner":"0xabc","settled_amount_wei":"500","mode":"offchain","status":"confirmed","settlement_id":1}`))
	}))
	defer srv.Close()

	auth := types.Auth{Sign: []byte("fake-signature-bytes")}
	res, err := SettleMetering(srv.URL, auth)
	if err != nil {
		t.Fatalf("SettleMetering: %v", err)
	}
	if res.Status != "confirmed" || res.SettledAmountWei != "500" || res.SettlementID != 1 {
		t.Errorf("bad settle: %+v", res)
	}
}
