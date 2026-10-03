package sdk

import (
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProbeAllowed(t *testing.T) {
	for _, c := range []struct {
		ip   string
		want bool
	}{
		{"10.0.21.247", true}, // in-VPC node addresses are legitimate
		{"52.77.86.192", true},
		{"127.0.0.1", false},
		{"169.254.169.254", false}, // cloud metadata
		{"0.0.0.0", false},
		{"::1", false},
		{"fe80::1", false},
		{"224.0.0.1", false},
	} {
		if got := probeAllowed(net.ParseIP(c.ip)); got != c.want {
			t.Errorf("%s: got %v, want %v", c.ip, got, c.want)
		}
	}
}

func TestProbeEdgeRefusesLoopback(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(`{}`)) }))
	defer srv.Close()
	if _, err := ProbeEdge(srv.URL); err == nil || !strings.Contains(err.Error(), "refused") {
		t.Fatalf("loopback edge probed: %v", err)
	}
	// the plain Info (gateway health checks, local dev) still reaches it
	if _, err := Info(srv.URL); err != nil {
		t.Fatalf("Info on loopback: %v", err)
	}
}

func TestInfoDoesNotFollowRedirects(t *testing.T) {
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(`{}`)) }))
	defer target.Close()
	redir := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL+"/v1/info", http.StatusFound)
	}))
	defer redir.Close()
	if _, err := Info(redir.URL); err == nil {
		t.Fatal("followed a redirect")
	}
}
