package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
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

	form := url.Values{}
	form.Set("txMsg", hex.EncodeToString(mmb))

	_, err = doRequest(context.TODO(), baseUrl, "/v1/edges", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	return nil
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

	resByte, err := Get(context.TODO(), u)
	if err != nil {
		return res, err
	}

	if err = unwrapItems(resByte, &res.Edges); err != nil {
		return res, err
	}

	logger.Debug("edge list: ", res)
	return res, nil
}
