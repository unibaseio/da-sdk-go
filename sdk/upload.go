package sdk

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/archive"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/types"

	darchive "github.com/docker/docker/pkg/archive"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/go-homedir"
	"github.com/schollz/progressbar/v3"
)

func Upload(baseUrl string, auth types.Auth, policy types.Policy, filePath string, name string) (types.FileFull, common.Address, error) {
	var res types.FileFull
	er, err := ListEdge(baseUrl, auth, types.StreamType)
	if err != nil {
		return res, common.Address{}, err
	}

	logger.Debug("streams before: ", er.Edges)
	Disorder(er.Edges)
	logger.Debug("streams after: ", er.Edges)

	if os.Getenv("STREAM_PRIORITY") != "" {
		//"0x3a2e98eaaa6ba9e3102cb622945f102c38221268"
		priaddr := common.HexToAddress(os.Getenv("STREAM_PRIORITY"))
		at := 0
		for j, em := range er.Edges {
			if em.Name == priaddr {
				at = j
				break
			}
		}

		if at != 0 {
			tmp := er.Edges[at]
			er.Edges[at] = er.Edges[0]
			er.Edges[0] = tmp
		}
	}

	for _, em := range er.Edges {
		if em.Type != types.StreamType {
			continue
		}
		if !em.OnChain {
			continue
		}
		fr, err := UploadData(em.ExposeURL, auth, policy, filePath)
		if err != nil {
			logger.Debug("upload: ", filePath, " to: ", em.ExposeURL, " fail: ", err)
			if strings.Contains(err.Error(), "already has") {
				return res, em.Name, err
			}
			continue
		}

		if name != "" {
			fr.Name = name
		}

		if fr.ChainType == "" {
			fr.ChainType = chaintype
		}

		logger.Debug("upload meta: ", filePath, " to: ", baseUrl)
		err = UploadFileMeta(baseUrl, auth, fr.FileReceipt)
		return fr, em.Name, err
	}

	return res, common.Address{}, fmt.Errorf("no avail streamer")
}

func UploadData(baseUrl string, auth types.Auth, policy types.Policy, filePath string) (types.FileFull, error) {
	logger.Debug("upload: ", filePath, " to: ", baseUrl)
	var res types.FileFull
	p, err := homedir.Expand(filePath)
	if err != nil {
		return res, err
	}

	bar := progressbar.DefaultBytes(-1, "upload:")

	ipr, ipw := io.Pipe()
	mwriter := multipart.NewWriter(ipw)
	go func() {
		defer ipw.Close()
		defer mwriter.Close()

		err = mwriter.WriteField("rsn", strconv.Itoa(int(policy.N)))
		if err != nil {
			return
		}

		err = mwriter.WriteField("rsk", strconv.Itoa(int(policy.K)))
		if err != nil {
			return
		}

		part, err := mwriter.CreateFormFile("file", p)
		if err != nil {
			return
		}

		fi, err := os.Stat(p)
		if err != nil {
			return
		}

		if fi.IsDir() {
			pf, err := darchive.Tar(p, darchive.Gzip)
			if err != nil {
				return
			}
			defer pf.Close()

			pr := progressbar.NewReader(pf, bar)

			io.Copy(part, &pr)
		} else {
			pf, err := os.Open(p)
			if err != nil {
				return
			}
			defer pf.Close()

			pr := progressbar.NewReader(pf, bar)

			io.Copy(part, &pr)
		}
	}()

	haddr := baseUrl + "/api/upload"
	hreq, err := http.NewRequest("POST", haddr, ipr)
	if err != nil {
		return res, err
	}

	aub, err := json.Marshal(auth)
	if err != nil {
		return res, err
	}

	hreq.Header.Add("Authorization", hex.EncodeToString(aub))
	hreq.Header.Add("Content-Type", mwriter.FormDataContentType())
	defaultHTTPClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			WriteBufferSize:       16 << 10, // 16KiB moving up from 4KiB default
			ReadBufferSize:        16 << 10, // 16KiB moving up from 4KiB default
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableCompression:    true,
		},
	}

	resp, err := defaultHTTPClient.Do(hreq)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	resByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("response: %s, msg: %s", resp.Status, resByte)
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	bar.Finish()

	return res, nil
}

func walk(baseDir, curdir string) (map[string]string, uint64, error) {
	res := make(map[string]string)
	size := uint64(0)
	fulDir := path.Join(baseDir, curdir)
	rd, err := os.ReadDir(fulDir)
	if err != nil {
		return res, size, err
	}
	h := sha256.New()
	for _, fde := range rd {
		if fde.IsDir() {
			subCur := path.Join(curdir, fde.Name())
			subRes, subSize, err := walk(baseDir, subCur)
			for k, v := range subRes {
				res[k] = v
			}
			if err != nil {
				return res, size, err
			}
			size += subSize
			continue
		}

		fi, err := fde.Info()
		if err != nil {
			return res, size, err
		}
		if fi.Size() < bls.MaxSize && fi.Name() != archive.ShadowTar {
			continue
		}
		osf, err := os.OpenFile(path.Join(fulDir, fi.Name()), os.O_RDONLY, os.ModePerm)
		if err != nil {
			return res, size, err
		}
		h.Reset()
		_, err = io.Copy(h, osf)
		if err != nil {
			osf.Close()
			return res, size, err
		}
		osf.Close()
		hs := h.Sum(nil)
		res[path.Join(curdir, fi.Name())] = hex.EncodeToString(hs)
		size += uint64(fi.Size())
	}
	return res, size, nil
}
