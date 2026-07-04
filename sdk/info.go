package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

func Info(baseUrl string) (types.EdgeReceipt, error) {
	res := types.EdgeReceipt{}
	resp, err := http.Get(baseUrl + "/v1/info")
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("response: %s", resp.Status)
	}

	resb, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resb, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func Login(baseUrl string, auth types.Auth) error {
	form := url.Values{}
	form.Set("chaintype", chaintype)

	_, err := doRequest(context.TODO(), baseUrl, "/v1/login", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	return nil
}
