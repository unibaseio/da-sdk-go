package sdk

import (
	"context"
	"encoding/binary"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
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

// S8: AWS's IPv6 metadata (and its other service addresses), NAT64 forms of
// refused IPv4 addresses and 0.0.0.0/8 are refused too.
func TestProbeAllowedMore(t *testing.T) {
	for _, c := range []struct {
		ip   string
		want bool
	}{
		{"fd00:ec2::254", false}, // AWS IMDS over IPv6
		{"fd00:ec2::253", false}, // AWS DNS
		{"64:ff9b::a9fe:a9fe", false},
		{"64:ff9b::7f00:1", false},
		{"0.1.2.3", false},
		{"::ffff:169.254.169.254", false},
		{"fd12:3456::1", true}, // other unique-local (private) ranges stay allowed
		{"2600:1f18::1", true},
		{"64:ff9b::808:808", true},
		{"64:ff9b:1::a00:1", false}, // local-use NAT64 (RFC 8215): embedding unknown
		{"64:ff9b:1:ffff::1", false},
	} {
		if got := probeAllowed(net.ParseIP(c.ip)); got != c.want {
			t.Errorf("%s: got %v, want %v", c.ip, got, c.want)
		}
	}
}

// fakeDNS answers A queries from answer(n) for the n-th A query (1-based) and
// AAAA queries with no records. It serves on 127.0.0.1 over UDP.
func fakeDNS(t *testing.T, answer func(n int32) net.IP) string {
	t.Helper()
	pc, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pc.Close() })
	var aQueries atomic.Int32
	go func() {
		buf := make([]byte, 1500)
		for {
			n, from, err := pc.ReadFrom(buf)
			if err != nil {
				return
			}
			q := buf[:n]
			if len(q) < 12 {
				continue
			}
			i := 12
			for i < len(q) && q[i] != 0 {
				i += int(q[i]) + 1
			}
			if i+5 > len(q) {
				continue
			}
			qtype := binary.BigEndian.Uint16(q[i+1:])
			question := q[12 : i+5]
			resp := []byte{q[0], q[1], 0x81, 0x80, 0, 1, 0, 0, 0, 0, 0, 0}
			resp = append(resp, question...)
			if qtype == 1 { // A
				ip := answer(aQueries.Add(1)).To4()
				resp[7] = 1
				resp = append(resp, 0xc0, 0x0c, 0, 1, 0, 1, 0, 0, 0, 0, 0, 4)
				resp = append(resp, ip...)
			}
			pc.WriteTo(resp, from)
		}
	}()
	return pc.LocalAddr().String()
}

// useResolver points the process's default resolver at a fake DNS server.
func useResolver(t *testing.T, dnsAddr string) {
	t.Helper()
	old := net.DefaultResolver
	net.DefaultResolver = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, _, _ string) (net.Conn, error) {
			return (&net.Dialer{}).DialContext(ctx, "udp", dnsAddr)
		},
	}
	t.Cleanup(func() { net.DefaultResolver = old })
}

// localAllowedIPv4 is an address of this host that probeAllowed accepts (a
// LAN address), so a test can serve on it without leaving the host.
func localAllowedIPv4(t *testing.T) net.IP {
	t.Helper()
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		t.Skip(err)
	}
	for _, a := range addrs {
		if n, ok := a.(*net.IPNet); ok && n.IP.To4() != nil && probeAllowed(n.IP) {
			return n.IP.To4()
		}
	}
	t.Skip("no non-loopback IPv4 address on this host")
	return nil
}

// countingInfo serves /v1/info on l, counting requests.
func countingInfo(t *testing.T, l net.Listener) *atomic.Int32 {
	t.Helper()
	var hits atomic.Int32
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.Write([]byte(`{}`))
	}))
	srv.Listener.Close()
	srv.Listener = l
	srv.Start()
	t.Cleanup(srv.Close)
	return &hits
}

// S8: DNS rebinding. The name first resolves to an allowed address (a decoy
// edge on this host's LAN address), then to loopback, both on the same port.
// A prober that resolves to check and resolves again to connect ends up on
// the loopback service; ProbeEdge must connect only where its check looked.
func TestProbeEdgeDNSRebinding(t *testing.T) {
	lan := localAllowedIPv4(t)
	var decoyL, loopL net.Listener
	for try := 0; try < 20 && loopL == nil; try++ {
		l, err := net.Listen("tcp", net.JoinHostPort(lan.String(), "0"))
		if err != nil {
			t.Skip(err)
		}
		_, port, _ := net.SplitHostPort(l.Addr().String())
		if ll, err := net.Listen("tcp", "127.0.0.1:"+port); err == nil {
			decoyL, loopL = l, ll
		} else {
			l.Close()
		}
	}
	if loopL == nil {
		t.Skip("no port free on both addresses")
	}
	_, port, _ := net.SplitHostPort(decoyL.Addr().String())
	decoy := countingInfo(t, decoyL)
	loop := countingInfo(t, loopL)
	useResolver(t, fakeDNS(t, func(n int32) net.IP {
		if n == 1 {
			return lan
		}
		return net.ParseIP("127.0.0.1")
	}))

	// (a sandbox may intercept non-loopback connections, so the decoy can
	// be unreachable: bound the wait)
	oldTimeout := infoClient.Timeout
	infoClient.Timeout = 2 * time.Second
	defer func() { infoClient.Timeout = oldTimeout }()

	_, err := ProbeEdge("http://rebind.example:" + port)
	if loop.Load() != 0 {
		t.Fatalf("probe reached the loopback service after a DNS rebind (err=%v)", err)
	}
	if err == nil && decoy.Load() != 1 {
		t.Fatalf("probe succeeded without reaching the address it resolved (decoy hits %d)", decoy.Load())
	}
	t.Logf("probe went to the first answer only: decoy hits %d, err %v", decoy.Load(), err)
}

// S8: a name that resolves straight to loopback is refused at dial time.
func TestProbeEdgeRefusesResolvedLoopback(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()
	_, port, _ := net.SplitHostPort(srv.Listener.Addr().String())
	useResolver(t, fakeDNS(t, func(int32) net.IP { return net.ParseIP("127.0.0.1") }))
	if _, err := ProbeEdge("http://loop.example:" + port); err == nil || !strings.Contains(err.Error(), "refused address") || hits.Load() != 0 {
		t.Fatalf("err=%v hits=%d", err, hits.Load())
	}
}
