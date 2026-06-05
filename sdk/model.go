package sdk

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/url"
	"os"
	"path"
	"strings"

	com "github.com/unibaseio/da-sdk-go/contract/common"
	contract "github.com/unibaseio/da-sdk-go/contract/v1"
	"github.com/unibaseio/da-sdk-go/lib/archive"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

func SubmitModelRepo(baseUrl string, auth types.Auth, mr types.RepoCore) error {
	mmb, err := json.Marshal(mr)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("repo", hex.EncodeToString(mmb))

	res, err := doRequest(context.TODO(), baseUrl, "/api/submitRepo", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	logger.Debug("response: ", res)

	return nil
}

func GetModelRepo(baseUrl string, auth types.Auth, rc types.RepoCore) (types.RepoMeta, error) {
	rm := new(types.RepoMeta)

	form := url.Values{}
	form.Set("name", rc.ID())

	res, err := doRequest(context.TODO(), baseUrl, "/api/getRepo", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return *rm, err
	}
	err = json.Unmarshal(res, rm)
	if err != nil {
		return *rm, err
	}
	logger.Debug("response: ", rm)

	return *rm, err
}

func SubmitModel(baseUrl string, auth types.Auth, mr types.ModelMeta) error {
	mmb, err := json.Marshal(mr)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("txMsg", hex.EncodeToString(mmb))

	res, err := doRequest(context.TODO(), baseUrl, "/api/submitModel", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	re := new(types.ModelMeta)
	err = json.Unmarshal(res, re)
	if err != nil {
		return err
	}

	return nil
}

func BuyModel(baseUrl string, auth types.Auth, modelName string) error {
	form := url.Values{}
	form.Set("name", modelName)

	_, err := doRequest(context.TODO(), baseUrl, "/api/buyModel", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	return nil
}

func ListModel(baseUrl string, auth types.Auth, filter string) (types.ListModelResult, error) {
	res := types.ListModelResult{}
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

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listModel", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("model list my: ", res)
	return res, nil
}

func GetModel(baseUrl string, auth types.Auth, modelName string) (types.ModelResult, error) {
	form := url.Values{}
	form.Set("name", modelName)

	re := types.ModelResult{}
	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getModel", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return re, err
	}

	err = json.Unmarshal(resByte, &re)
	if err != nil {
		return re, err
	}

	return re, nil
}

func ListLocalDir(d string) (types.ModelMeta, error) {
	mrm := types.ModelMeta{
		Name: path.Base(d),
	}

	dst := path.Join(d, archive.ShadowTar)
	_, err := os.Stat(dst)
	if err != nil {
		os.Remove(dst)
		err := archive.Tar(d, dst, "gz", bls.MaxSize)
		if err != nil {
			return mrm, err
		}
	}

	res, size, err := walk(d, "")
	if err != nil {
		return mrm, err
	}
	mrm.Files = res
	mrm.Size = size

	logger.Debug("files: ", mrm.Files)

	return mrm, nil
}

func UploadModelFiles(url string, sk *ecdsa.PrivateKey, au types.Auth, fp string) (types.ModelMeta, error) {
	mrm, err := ListLocalDir(fp)
	if err != nil {
		return mrm, err
	}

	cm, err := contract.NewContractManage(sk, com.OPSepolia)
	if err != nil {
		panic(err)
	}

	Login(url, au)

	for k, v := range mrm.Files {
		fr, err := GetFileReceipt(url, au, v)
		if err == nil && fr.Name == v {
			logger.Debug("local already has file: ", v)
			continue
		}

		sfp := path.Join(fp, k)
		ff, streamer, err := Upload(url, au, types.Policy{6, 4}, sfp, "")
		if err != nil {
			return mrm, err
		}

		pcs, err := CheckFileFull(ff, streamer, sfp)
		if err != nil {
			return mrm, err
		}

		for _, pc := range pcs {
			_, err = cm.AddPiece(pc)
			if err != nil {
				return mrm, err
			}
		}

		logger.Debugf("uploaded %s to %s, sha256: %s\n", sfp, streamer, ff.Hash)
		if k == archive.ShadowTar {
			os.Remove(sfp)
		}
	}
	return mrm, nil
}

func DownloadModel(url, rootDir string, au types.Auth, mrm types.ModelMeta, ks types.IPieceStore) error {
	has := Exists(rootDir)
	if !has {
		err := os.MkdirAll(rootDir, 0755)
		if err != nil {
			return err
		}
	}

	baseDir := path.Join(rootDir, mrm.Name)
	has = Exists(baseDir)
	if !has {
		err := os.MkdirAll(baseDir, 0755)
		if err != nil {
			return err
		}
	}
	h := sha256.New()
	for k, v := range mrm.Files {
		p := path.Join(baseDir, k)
		has := Exists(p)
		if has {
			logger.Debug("check local:", p)
			if len(v) != 64 {
				continue
			}
			osf, err := os.OpenFile(p, os.O_RDONLY, os.ModePerm)
			if err != nil {
				return err
			}
			h.Reset()
			_, err = io.Copy(h, osf)
			if err != nil {
				osf.Close()
				return err
			}
			osf.Close()
			hs := h.Sum(nil)
			if v == hex.EncodeToString(hs) {
				if k == archive.ShadowTar {
					err = archive.UnTar(p, baseDir, "gz")
					if err != nil {
						return err
					}
					os.Remove(p)
				}
				continue
			}

			logger.Warnf("remove local file %s due to unmatch hash", p)
			err = os.Remove(p)
			if err != nil {
				return err
			}
		}
		pdir := path.Dir(p)
		err := os.MkdirAll(pdir, 0755)
		if err != nil {
			return err
		}

		fi, err := os.Create(p)
		if err != nil {
			return err
		}

		// download it if sha256
		logger.Debug("parallel download: ", k, v)
		err = DownloadParallel(url, au, v, 8, ks, fi)
		if err != nil {
			return err
		}

		if k == archive.ShadowTar {
			err = archive.UnTar(p, baseDir, "gz")
			if err != nil {
				return err
			}
			os.Remove(p)
		}
	}

	return nil
}

func Exists(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}
