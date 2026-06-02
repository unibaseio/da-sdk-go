package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

func ListAccountHub(baseUrl string) error {
	resb, err := doRequest(context.TODO(), baseUrl, "/api/listAccount", "application/json", types.Auth{}, nil)
	if err != nil {
		return err
	}

	var res []types.Account
	err = json.Unmarshal(resb, &res)
	if err != nil {
		return err
	}
	logger.Info("got account: ", res)

	return nil
}

func ListNeedleHub(baseUrl string, owner string) error {
	form := url.Values{}
	form.Set("owner", owner)

	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancle()
	resb, err := doRequest(ctx, baseUrl, "/api/listNeedle", "", types.Auth{}, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	var res []types.Needle
	err = json.Unmarshal(resb, &res)
	if err != nil {
		return err
	}
	logger.Info("got needle: ", res)

	return nil
}

func UploadHub(baseUrl, owner, filename, msg string) error {
	ms := types.MemeStruct{
		Owner:   owner,
		ID:      filename,
		Message: msg,
	}
	msb, err := json.Marshal(ms)
	if err != nil {
		return err
	}

	resb, err := doRequest(context.TODO(), baseUrl, "/api/upload", "application/json", types.Auth{}, bytes.NewReader(msb))
	if err != nil {
		return err
	}

	var res types.MemeMeta
	err = json.Unmarshal(resb, &res)
	if err != nil {
		return err
	}
	logger.Info("upload done: ", res)

	return nil
}

func UploadHubData(baseUrl, owner, filename string, data []byte) error {
	ipr, ipw := io.Pipe()
	mwriter := multipart.NewWriter(ipw)
	go func() {
		defer ipw.Close()
		defer mwriter.Close()

		err := mwriter.WriteField("owner", owner)
		if err != nil {
			return
		}

		part, err := mwriter.CreateFormFile("file", filename)
		if err != nil {
			return
		}

		part.Write(data)
	}()

	resb, err := doRequest(context.TODO(), baseUrl, "/api/uploadData", mwriter.FormDataContentType(), types.Auth{}, ipr)
	if err != nil {
		return err
	}

	var res types.MemeMeta
	err = json.Unmarshal(resb, &res)
	if err != nil {
		return err
	}
	logger.Info("upload done: ", res)

	return nil
}

func DownloadHubData(baseUrl string, owner, filename string) ([]byte, error) {
	form := url.Values{}
	form.Set("id", filename)
	form.Set("owner", owner)

	logger.Debugf("download %s %s from hub %s", filename, owner, baseUrl)
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancle()
	resByte, err := doRequest(ctx, baseUrl, "/api/download", "", types.Auth{}, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	return resByte, nil
}
