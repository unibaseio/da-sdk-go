package sdk

import (
	"crypto/ecdsa"
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
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/go-homedir"
	"github.com/schollz/progressbar/v3"
)

// Signer signs one request's Authorization for the given purpose label.
type Signer func(label []byte) (types.Auth, error)

// KeySigner signs with a private key, fresh on every call.
func KeySigner(sk *ecdsa.PrivateKey) Signer {
	return func(label []byte) (types.Auth, error) { return key.BuildAuth(sk, label) }
}

func Upload(baseUrl string, auth types.Auth, policy types.Policy, filePath string, name string) (types.FileFull, common.Address, error) {
	return UploadWith(baseUrl, func([]byte) (types.Auth, error) { return auth, nil }, policy, filePath, name)
}

// UploadWith is Upload with a signer instead of one fixed header: the stream
// upload and the gateway file record are each signed just before they are
// sent, so the record still carries a fresh signature after a long encode.
// Gateways and streams reject signatures older than a few minutes.
func UploadWith(baseUrl string, sign Signer, policy types.Policy, filePath string, name string) (types.FileFull, common.Address, error) {
	var res types.FileFull
	auth, err := sign([]byte("upload"))
	if err != nil {
		return res, common.Address{}, err
	}
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
		upAuth, err := sign([]byte("upload"))
		if err != nil {
			return res, common.Address{}, err
		}
		fr, err := UploadData(em.ExposeURL, upAuth, policy, filePath)
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
		metaAuth, err := sign([]byte("upload"))
		if err != nil {
			return fr, em.Name, err
		}
		err = UploadFileMeta(baseUrl, metaAuth, fr.FileReceipt)
		return fr, em.Name, err
	}

	return res, common.Address{}, fmt.Errorf("no avail streamer")
}

// uploadAnswerLimit bounds what UploadData reads of a stream's answer (it
// used to read without limit, 2026-10-06 audit N11): per piece, one witness —
// 13.5 KB raw for 64/32, ~18 KB as JSON — and a few names, so 64 KiB each, plus
// 64 KiB. A directory's tar size is not known up front: 64 MiB.
func uploadAnswerLimit(p string, policy types.Policy) int64 {
	const perPiece, base, unknown = 64 << 10, 64 << 10, 64 << 20
	fi, err := os.Stat(p)
	max := MaxPieceSize(policy)
	if err != nil || fi.IsDir() || max <= 0 {
		return unknown
	}
	pieces := int64(1)
	if fi.Size() > 0 {
		pieces = 1 + (fi.Size()-1)/max
	}
	return base + pieces*perPiece
}

func UploadData(baseUrl string, auth types.Auth, policy types.Policy, filePath string) (types.FileFull, error) {
	logger.Debug("upload: ", filePath, " to: ", baseUrl)
	var res types.FileFull
	p, err := homedir.Expand(filePath)
	if err != nil {
		return res, err
	}

	bar := progressbar.DefaultBytes(-1, "upload:")

	// The body is streamed from a goroutine. Any failure there (stat, open,
	// tar, read) aborts the request through CloseWithError and is returned —
	// it used to end the body early, so the stream encoded a truncated file
	// and the upload reported success. The bytes sent are hashed so the
	// stream's receipt can be checked against them.
	sent := sha256.New()
	var sentN int64
	ipr, ipw := io.Pipe()
	mwriter := multipart.NewWriter(ipw)
	done := make(chan error, 1)
	go func() {
		err := writeUploadBody(mwriter, p, policy, bar, sent, &sentN)
		if err == nil {
			err = mwriter.Close()
		}
		ipw.CloseWithError(err) // nil → EOF
		done <- err
	}()

	haddr := baseUrl + "/v1/upload"
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
		ipr.CloseWithError(err)
		if berr := <-done; berr != nil && berr != io.ErrClosedPipe {
			return res, fmt.Errorf("upload %s: %w", filePath, berr)
		}
		return res, err
	}
	defer resp.Body.Close()

	limit := uploadAnswerLimit(p, policy)
	resByte, err := io.ReadAll(io.LimitReader(resp.Body, limit+1))
	ipr.CloseWithError(io.ErrClosedPipe) // unblock the writer if the server answered early
	berr := <-done
	if err != nil {
		return res, err
	}
	if int64(len(resByte)) > limit {
		return res, fmt.Errorf("upload %s: the stream's answer exceeds %d bytes", filePath, limit)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return res, fmt.Errorf("response: %s, msg: %s", resp.Status, resByte)
	}
	if berr != nil {
		return res, fmt.Errorf("upload %s: %w", filePath, berr)
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	bar.Finish()

	// the stream's answer must describe exactly the bytes sent, encoded under
	// the requested policy
	if err := CheckFileFullShape(res, policy); err != nil {
		return res, fmt.Errorf("stream answer: %w", err)
	}
	if res.Size != sentN {
		return res, fmt.Errorf("stream answer: size %d, sent %d bytes", res.Size, sentN)
	}
	if got := hex.EncodeToString(sent.Sum(nil)); !strings.EqualFold(res.Hash, got) {
		return res, fmt.Errorf("stream answer: hash %s, sent data hashes to %s", res.Hash, got)
	}

	return res, nil
}

// writeUploadBody writes the multipart upload form for the file or directory
// (as tar.gz) at p, hashing and counting the file bytes into sent / n.
func writeUploadBody(mw *multipart.Writer, p string, policy types.Policy, bar *progressbar.ProgressBar, sent io.Writer, n *int64) error {
	if err := mw.WriteField("rsn", strconv.Itoa(int(policy.N))); err != nil {
		return err
	}
	if err := mw.WriteField("rsk", strconv.Itoa(int(policy.K))); err != nil {
		return err
	}
	part, err := mw.CreateFormFile("file", p)
	if err != nil {
		return err
	}
	fi, err := os.Stat(p)
	if err != nil {
		return err
	}
	var src io.ReadCloser
	if fi.IsDir() {
		src, err = archive.TarGz(p)
	} else {
		src, err = os.Open(p)
	}
	if err != nil {
		return err
	}
	defer src.Close()

	pr := progressbar.NewReader(src, bar)
	c, err := io.Copy(io.MultiWriter(part, sent), &pr)
	*n = c
	return err
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
