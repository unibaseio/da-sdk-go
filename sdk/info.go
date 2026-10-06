package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"syscall"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

// infoClient bounds /v1/info probes: the URL can come from the open edge
// registry, so a target that never answers, answers with an endless body, or
// redirects elsewhere must not hold or steer the caller.
var infoClient = &http.Client{
	Timeout: 10 * time.Second,
	CheckRedirect: func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

// maxInfoBytes caps a /v1/info response; a real one is well under 4 KiB.
const maxInfoBytes = 64 << 10

func Info(baseUrl string) (types.EdgeReceipt, error) {
	return info(infoClient, baseUrl)
}

func info(cl *http.Client, baseUrl string) (types.EdgeReceipt, error) {
	res := types.EdgeReceipt{}
	resp, err := cl.Get(baseUrl + "/v1/info")
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("response: %s", resp.Status)
	}

	resb, err := io.ReadAll(io.LimitReader(resp.Body, maxInfoBytes))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resb, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// ProbeEdge is Info for a URL taken from the edge registry, where anyone can
// register any URL. On top of Info's bounds it refuses to connect to loopback,
// link-local (incl. the cloud metadata addresses), unspecified, "this host"
// or multicast addresses, so the registry cannot aim the prober at the
// prober's own host or its cloud metadata. The check runs on the address
// actually dialed (probeTransport's dialer Control), after resolution: resolving
// first and dialing later let a DNS answer change in between (rebinding).
// Private VPC ranges stay allowed: nodes register their in-VPC addresses.
func ProbeEdge(baseUrl string) (types.EdgeReceipt, error) {
	if _, err := url.Parse(baseUrl); err != nil {
		return types.EdgeReceipt{}, err
	}
	cl := *infoClient // Info's bounds, with the vetting transport
	cl.Transport = probeTransport
	return info(&cl, baseUrl)
}

// probeTransport connects only to addresses probeAllowed accepts. No proxy:
// the dialed address must be the edge's own.
var probeTransport = &http.Transport{
	Proxy: nil,
	DialContext: (&net.Dialer{
		Timeout: 10 * time.Second,
		Control: probeControl,
	}).DialContext,
	TLSHandshakeTimeout: 10 * time.Second,
	DisableKeepAlives:   true,
}

// probeControl vets each address the probe dials (after DNS resolution).
func probeControl(network, address string, _ syscall.RawConn) error {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return fmt.Errorf("edge probe: bad address %q", address)
	}
	ip := net.ParseIP(host) // a zoned (%iface) address does not parse: refused
	if ip == nil || !probeAllowed(ip) {
		return fmt.Errorf("edge probe: refused address %s", address)
	}
	return nil
}

var (
	// AWS's IPv6 service range: instance metadata (fd00:ec2::254), DNS
	// (::253), time sync (::123). No node is reachable there.
	awsIPv6Services = mustCIDR("fd00:ec2::/32")
	// NAT64 (RFC 6052): an IPv4 address in IPv6 form, judged as that IPv4.
	nat64 = mustCIDR("64:ff9b::/96")
	// Local-use NAT64 (RFC 8215): the embedding is operator-chosen, so the
	// IPv4 behind an address cannot be read off; refused outright.
	nat64Local = mustCIDR("64:ff9b:1::/48")
	// 0.0.0.0/8 "this network": 0.x.x.x reaches the local host on Linux.
	thisNetwork = mustCIDR("0.0.0.0/8")
)

func mustCIDR(s string) *net.IPNet {
	_, n, err := net.ParseCIDR(s)
	if err != nil {
		panic(err)
	}
	return n
}

func probeAllowed(ip net.IP) bool {
	if nat64.Contains(ip) && ip.To4() == nil {
		return probeAllowed(net.IP(ip[12:16]))
	}
	return !(ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() ||
		ip.IsInterfaceLocalMulticast() || ip.IsMulticast() || ip.IsUnspecified() ||
		awsIPv6Services.Contains(ip) || thisNetwork.Contains(ip) || nat64Local.Contains(ip))
}

func Login(baseUrl string, auth types.Auth) error {
	form := url.Values{}
	form.Set("chain", chaintype)

	_, err := doRequest(context.TODO(), baseUrl, "/v1/login", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	return nil
}
