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

	_, err = doRequest(context.TODO(), baseUrl, "/api/registerEdge", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	return nil
}

func GetEdge(baseUrl string, auth types.Auth, eaddr common.Address) (types.EdgeReceipt, error) {
	res := types.EdgeReceipt{}

	form := url.Values{}
	form.Set("name", eaddr.String())
	form.Set("chaintype", chaintype)
	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getEdge", "", auth, strings.NewReader(form.Encode()))
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
	opt := types.Options{
		UserDefined: make(map[string]string),
	}
	opt.UserDefined["filter"] = filter
	opt.UserDefined["chaintype"] = chaintype

	optyByte, err := json.Marshal(opt)
	if err != nil {
		return res, err
	}

	form := url.Values{}
	form.Set("option", hex.EncodeToString(optyByte))

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listEdge", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("edge list: ", res)
	return res, nil
}
