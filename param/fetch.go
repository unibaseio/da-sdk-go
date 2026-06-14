package param

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/env"
	"github.com/unibaseio/da-sdk-go/lib/log"

	"github.com/mitchellh/go-homedir"
	"github.com/schollz/progressbar/v3"
)

var logger = log.Logger("param")

const (
	chunkSize = 8 * 1024 * 1024
	paramdir  = "~/.plonk"
)

// gateway is the base URL for the shared SRS/param download; override with
// PARAM_SOURCE (e.g. an internal mirror/cache).
var gateway = env.Str(env.ParamSource, "https://unibase-params.s3.ap-southeast-1.amazonaws.com/")

func CheckAll() error {
	for pn, pi := range paramMap {
		err := checkFile(pn, pi)
		if err != nil {
			return err
		}
	}
	return nil
}

func FetchAll() error {
	for pn := range paramMap {
		err := Fetch(pn)
		if err != nil {
			return err
		}
	}
	return nil
}

func FetchKZG() error {
	pn := "kzgsrs-bls12_377-kzg"

	err := Fetch(pn)
	if err != nil {
		return err
	}
	return nil
}

func FetchKZGVK() error {
	pn := "kzgsrs-bls12_377-kzg-vk"

	err := Fetch(pn)
	if err != nil {
		return err
	}
	return nil
}

func Fetch(pn string) error {
	pi, ok := paramMap[pn]
	if !ok {
		return fmt.Errorf("no such param in config")
	}
	err := checkFile(pn, pi)
	if err == nil {
		return nil
	}
	logger.Warn(err)
	return downloadFile(pn, pi)
}

func checkFile(fn string, pi paramInfo) error {
	fp, err := homedir.Expand(paramdir)
	if err != nil {
		return err
	}
	fp = filepath.Join(fp, fn)
	fi, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer fi.Close()
	fs, err := fi.Stat()
	if err != nil {
		return err
	}
	if fs.Size() != pi.size {
		return fmt.Errorf("mismatch size at %s", fn)
	}
	h := sha256.New()
	_, err = io.Copy(h, fi)
	if err != nil {
		return err
	}

	if hex.EncodeToString(h.Sum(nil)) != pi.hash {
		return fmt.Errorf("mismatch hash at %s", fn)
	}
	logger.Debug("local has right param: ", fn)
	return nil
}

func downloadFile(pn string, pi paramInfo) error {
	logger.Debug("start fetch param: ", pn)
	fp, err := homedir.Expand(paramdir)
	if err != nil {
		return err
	}
	os.MkdirAll(fp, 0755)

	fp = filepath.Join(fp, pn)
	os.Remove(fp)
	fi, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer fi.Close()

	url := gateway + pn
	bar := progressbar.DefaultBytes(pi.size, pn+" download:")

	h := sha256.New()

	mw := io.MultiWriter(h, fi, bar)

	for start := int64(0); start < pi.size; start += chunkSize {
		end := start + chunkSize - 1
		if end >= pi.size {
			end = pi.size - 1
		}

		val, err := downloadChunk(url, start, end)
		if err != nil {
			time.Sleep(3 * time.Second)
			val, err = downloadChunk(url, start, end)
			if err != nil {
				return err
			}
		}
		mw.Write(val)
	}
	bar.Finish()
	if hex.EncodeToString(h.Sum(nil)) != pi.hash {
		return fmt.Errorf("mismatch hash at downloaded %s", pn)
	}
	logger.Debug("finish fetch param: ", pn)
	return nil
}

func downloadChunk(url string, start, end int64) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent {
		return nil, fmt.Errorf("server returned status code %d", resp.StatusCode)
	}

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
