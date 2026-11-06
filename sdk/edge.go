package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
	"net/url"
	"strings"

	"github.com/unibaseio/da-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/common"
)

func RegisterEdge(baseUrl string, auth types.Auth, em types.EdgeMeta) error {
	mmb, err := json.Marshal(em)
	if err != nil {
		return err
	}
	return registerEdgeRaw(baseUrl, auth, mmb)
}

// EdgeAuthHash is what a registration's signature covers: the registration
// itself. A header signed over a label (RegisterEdge with a "register" auth)
// could be replayed within its window to register the edge again with
// another ExposeURL (audit 2026-10-13 M1/F9).
func EdgeAuthHash(txMsg []byte) []byte {
	return crypto.Keccak256([]byte("da-edge-register/v1"), txMsg)
}

// EnsureEdgeSigned is EnsureEdge with the registration signed by sign over
// EdgeAuthHash of it, so the signature is good for this registration only.
func EnsureEdgeSigned(baseUrl string, sign func([]byte) (types.Auth, error), em types.EdgeMeta) (already bool, err error) {
	cur, gerr := GetEdge(baseUrl, types.Auth{}, em.Name)
	if gerr == nil && cur.ExposeURL != "" && cur.ExposeURL == em.ExposeURL {
		return true, nil
	}
	mmb, err := json.Marshal(em)
	if err != nil {
		return false, err
	}
	auth, err := sign(EdgeAuthHash(mmb))
	if err != nil {
		return false, err
	}
	return false, registerEdgeRaw(baseUrl, auth, mmb)
}

// RegisterEdgeSigned registers em with a signature bound to it (EdgeAuthHash).
func RegisterEdgeSigned(baseUrl string, sign func([]byte) (types.Auth, error), em types.EdgeMeta) error {
	mmb, err := json.Marshal(em)
	if err != nil {
		return err
	}
	auth, err := sign(EdgeAuthHash(mmb))
	if err != nil {
		return err
	}
	return registerEdgeRaw(baseUrl, auth, mmb)
}

func registerEdgeRaw(baseUrl string, auth types.Auth, mmb []byte) error {
	form := url.Values{}
	form.Set("txMsg", hex.EncodeToString(mmb))

	_, err := doRequest(context.TODO(), baseUrl, "/v1/edges", "", auth, strings.NewReader(form.Encode()))
	return err
}

// EnsureEdge registers the edge exactly once at startup. If this address is already
// registered with the same ExposeURL it is a no-op (returns already=true); otherwise
// it registers via em. It does NOT retry: a returned error is meant to abort node
// startup so operators immediately see the node failed to register, rather than the
// node silently retrying forever or running invisible in the registry. Re-checking
// ExposeURL (not just existence) also self-corrects a stale/blank entry left by an
// earlier partial registration.
func EnsureEdge(baseUrl string, auth types.Auth, em types.EdgeMeta) (already bool, err error) {
	cur, gerr := GetEdge(baseUrl, auth, auth.Addr)
	if gerr == nil && cur.ExposeURL != "" && cur.ExposeURL == em.ExposeURL {
		return true, nil
	}
	if rerr := RegisterEdge(baseUrl, auth, em); rerr != nil {
		return false, rerr
	}
	return false, nil
}

func GetEdge(baseUrl string, auth types.Auth, eaddr common.Address) (types.EdgeReceipt, error) {
	res := types.EdgeReceipt{}

	resByte, err := Get(context.TODO(), v1URL(baseUrl, "/v1/edges/"+url.PathEscape(eaddr.String())))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func ListEdge(baseUrl string, auth types.Auth, filter string) (types.ListEdgeResult, error) {
	res := types.ListEdgeResult{}

	u := v1URL(baseUrl, "/v1/edges")
	if filter != "" {
		u += andSep(u) + "type=" + url.QueryEscape(filter)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), MetaTimeout)
	defer cancel()
	resByte, err := Get(ctx, u)
	if err != nil {
		return res, err
	}

	if err = unwrapItems(resByte, &res.Edges); err != nil {
		return res, err
	}

	logger.Debug("edge list: ", res)
	return res, nil
}
