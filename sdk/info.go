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
	res := types.EdgeReceipt{}
	resp, err := infoClient.Get(baseUrl + "/v1/info")
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
// register any URL. On top of Info's bounds it refuses hosts that resolve to
// loopback, link-local (incl. the cloud metadata address), unspecified or
// multicast addresses, so the registry cannot aim the prober at the prober's
// own host or its cloud metadata. Private VPC ranges stay allowed: nodes
// register their in-VPC addresses.
func ProbeEdge(baseUrl string) (types.EdgeReceipt, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return types.EdgeReceipt{}, err
	}
	ips, err := net.LookupIP(u.Hostname())
	if err != nil {
		return types.EdgeReceipt{}, err
	}
	for _, ip := range ips {
		if !probeAllowed(ip) {
			return types.EdgeReceipt{}, fmt.Errorf("edge url %s resolves to a refused address %s", baseUrl, ip)
		}
	}
	return Info(baseUrl)
}

func probeAllowed(ip net.IP) bool {
	return !(ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() ||
		ip.IsInterfaceLocalMulticast() || ip.IsMulticast() || ip.IsUnspecified())
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
