package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

func SubmitSpace(baseUrl string, auth types.Auth, msm types.SpaceMeta) (types.SpaceMeta, error) {
	var re types.SpaceMeta
	mmb, err := json.Marshal(msm)
	if err != nil {
		return re, err
	}

	form := url.Values{}
	form.Set("txMsg", hex.EncodeToString(mmb))

	res, err := doRequest(context.TODO(), baseUrl, "/api/submitSpace", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return re, err
	}

	err = json.Unmarshal(res, &re)
	if err != nil {
		return re, err
	}

	return re, nil
}

func ConfirmSpace(baseUrl string, auth types.Auth, sn, sroot string, pf []byte) error {
	form := url.Values{}
	form.Set("name", sn)
	form.Set("root", sroot)
	form.Set("proof", hex.EncodeToString(pf))

	_, err := doRequest(context.TODO(), baseUrl, "/api/confirmSpace", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	return nil
}

func UpdateSpace(baseUrl string, auth types.Auth, sn string) error {
	form := url.Values{}
	form.Set("name", sn)

	_, err := doRequest(context.TODO(), baseUrl, "/api/updateSpace", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	return nil
}

func GetSpace(baseUrl string, auth types.Auth, name string) (types.SpaceResult, error) {
	res := types.SpaceResult{}
	form := url.Values{}
	form.Set("name", name)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getSpace", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ListSpace(baseUrl string, auth types.Auth, filter string) (types.ListSpaceResult, error) {
	res := types.ListSpaceResult{}
	opt := types.Options{
		UserDefined: make(map[string]string),
	}
	opt.UserDefined["filter"] = filter

	optyByte, err := json.Marshal(opt)
	if err != nil {
		return res, err
	}

	form := url.Values{}
	form.Set("option", hex.EncodeToString(optyByte))

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listSpace", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("space list: ", res)
	return res, nil
}
