package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

func CreateBucket(baseUrl string, auth types.Auth, bn string) (minio.BucketInfo, error) {
	res := minio.BucketInfo{}
	form := url.Values{}
	form.Set("bucket", bn)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/createBucket", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("create bucket: ", res)

	return res, nil
}

func GetBucket(baseUrl string, auth types.Auth, bn string) (minio.BucketInfo, error) {
	res := minio.BucketInfo{}
	form := url.Values{}
	form.Set("bucket", bn)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getBucket", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}
	logger.Debug("get bucket: ", res)
	return res, nil
}

func DeleteBucket(baseUrl string, auth types.Auth, bn string) error {
	form := url.Values{}
	form.Set("bucket", bn)

	_, err := doRequest(context.TODO(), baseUrl, "/api/deleteBucket", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	return nil
}

func ListBucket(baseUrl string, auth types.Auth) ([]minio.BucketInfo, error) {
	res := []minio.BucketInfo{}
	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listBucket", "", auth, nil)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("bucket list: ", res)
	return res, nil
}

func PutObject(baseUrl string, auth types.Auth, bn, on string, nm types.NeedleMeta) error {
	form := url.Values{}
	form.Set("bucket", bn)
	form.Set("object", on)
	nmb, err := json.Marshal(nm)
	if err != nil {
		return err
	}
	form.Set("needle", hex.EncodeToString(nmb))
	_, err = doRequest(context.TODO(), baseUrl, "/api/putObject", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	return nil
}

func GetObject(baseUrl string, auth types.Auth, bn, on string) (minio.ObjectInfo, error) {
	res := minio.ObjectInfo{}
	form := url.Values{}
	form.Set("bucket", bn)
	form.Set("object", on)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getObject", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func DeleteObject(baseUrl string, auth types.Auth, bn, on string) error {
	form := url.Values{}
	form.Set("bucket", bn)
	form.Set("object", on)

	_, err := doRequest(context.TODO(), baseUrl, "/api/deleteObject", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	return nil
}

func ListObject(baseUrl string, auth types.Auth, bn string, opt types.Options) (minio.ListBucketResult, error) {
	res := minio.ListBucketResult{}

	form := url.Values{}
	form.Set("bucket", bn)
	optyByte, err := json.Marshal(opt)
	if err != nil {
		return res, err
	}
	form.Set("option", hex.EncodeToString(optyByte))

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listObject", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("object list: ", res)
	return res, nil
}
