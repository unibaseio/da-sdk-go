package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

func UploadFileMeta(baseUrl string, auth types.Auth, fcws types.FileReceipt) error {
	form := url.Values{}
	fcwsb, err := json.Marshal(fcws)
	if err != nil {
		return err
	}
	form.Set("meta", hex.EncodeToString(fcwsb))

	_, err = doRequest(context.TODO(), baseUrl, "/v1/files", "", auth, strings.NewReader(form.Encode()))
	return err
}

func GetReplicaReceipt(baseUrl string, auth types.Auth, name string) (types.ReplicaReceipt, error) {
	var res types.ReplicaReceipt

	resByte, err := Get(context.TODO(), v1URL(baseUrl, "/v1/replicas/"+url.PathEscape(name)))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ListReplicaByEdge(baseUrl string, addr string, start, count int) (types.ListReplicaResult, error) {
	res := types.ListReplicaResult{}

	u := v1URL(baseUrl, "/v1/replicas")
	u += andSep(u) + "edge=" + url.QueryEscape(addr) +
		"&offset=" + strconv.Itoa(start) + "&limit=" + strconv.Itoa(count)

	resByte, err := Get(context.TODO(), u)
	if err != nil {
		return res, err
	}

	if err = unwrapItems(resByte, &res.Replicas); err != nil {
		return res, err
	}

	logger.Debug("replica list: ", res)
	return res, nil
}

func GetPieceOfEdge(baseUrl string, name string) (types.PieceReceipt, error) {
	var res types.PieceReceipt

	resByte, err := Get(context.TODO(), v1URL(baseUrl, "/v1/pieces/"+url.PathEscape(name)))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ListReplica(baseUrl string, auth types.Auth, filter string) (types.ListReplicaResult, error) {
	res := types.ListReplicaResult{}

	u := v1URL(baseUrl, "/v1/replicas")
	if filter != "" {
		u += andSep(u) + "filter=" + url.QueryEscape(filter)
	}

	resByte, err := Get(context.TODO(), u)
	if err != nil {
		return res, err
	}

	if err = unwrapItems(resByte, &res.Replicas); err != nil {
		return res, err
	}

	logger.Debug("replica list: ", res)
	return res, nil
}

func GetPieceReceipt(baseUrl string, auth types.Auth, name string) (types.PieceReceipt, error) {
	var res types.PieceReceipt

	resByte, err := Get(context.TODO(), v1URL(baseUrl, "/v1/pieces/"+url.PathEscape(name)))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ListPiece(baseUrl string, auth types.Auth, filter string) (types.ListPieceResult, error) {
	return listPiece(baseUrl, filter, "")
}

// ListDispatch lists the pieces a stream can still hand to store: the stream
// leaves out pieces store already holds or cannot be given a slot of, and puts
// the most under-replicated first.
func ListDispatch(baseUrl string, store common.Address) (types.ListPieceResult, error) {
	return listPiece(baseUrl, "dispatch", store.Hex())
}

func listPiece(baseUrl, filter, store string) (types.ListPieceResult, error) {
	res := types.ListPieceResult{}

	u := v1URL(baseUrl, "/v1/pieces")
	if filter != "" {
		u += andSep(u) + "filter=" + url.QueryEscape(filter)
	}
	if store != "" {
		u += andSep(u) + "store=" + url.QueryEscape(store)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), MetaTimeout)
	defer cancel()
	resByte, err := Get(ctx, u)
	if err != nil {
		return res, err
	}

	if err = unwrapItems(resByte, &res.Pieces); err != nil {
		return res, err
	}

	logger.Debug("piece list: ", res)
	return res, nil
}

func GetFileReceipt(baseUrl string, auth types.Auth, name string) (types.FileReceipt, error) {
	var res types.FileReceipt

	resByte, err := Get(context.TODO(), v1URL(baseUrl, "/v1/files/"+url.PathEscape(name)))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ListFile(baseUrl string, auth types.Auth, filter string) (types.ListFileResult, error) {
	res := types.ListFileResult{}

	u := v1URL(baseUrl, "/v1/files")
	if filter != "" {
		u += andSep(u) + "filter=" + url.QueryEscape(filter)
	}

	resByte, err := Get(context.TODO(), u)
	if err != nil {
		return res, err
	}

	if err = unwrapItems(resByte, &res.Files); err != nil {
		return res, err
	}

	logger.Debug("file list: ", res)
	return res, nil
}

func RequestPiece(baseUrl string, auth types.Auth, name string) (types.PieceWitness, error) {
	form := url.Values{}
	form.Set("name", name)
	form.Set("chaintype", chaintype)

	var res types.PieceWitness
	ctx, cancel := context.WithTimeout(context.TODO(), MetaTimeout)
	defer cancel()
	resByte, err := doRequest(ctx, baseUrl, "/v1/requestPiece", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ConfirmPiece(baseUrl string, auth types.Auth, name, proof string) ([]byte, error) {
	form := url.Values{}
	form.Set("name", name)
	form.Set("chaintype", chaintype)

	resByte, err := doRequest(context.TODO(), baseUrl, "/v1/confirmPiece", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	return resByte, nil
}
