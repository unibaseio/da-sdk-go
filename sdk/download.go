package sdk

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/bls/erasure"
	"github.com/unibaseio/da-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/semaphore"
)

func DownloadPiece(baseUrl string, auth types.Auth, name string) (types.PieceCore, []byte, error) {
	logger.Debug("download piece: ", name, " from: ", baseUrl)
	pr, err := GetPieceReceipt(baseUrl, auth, name)
	if err != nil {
		return pr.PieceCore, nil, err
	}

	res := make([][]byte, pr.Policy.N)
	suc := 0
	need := make([]int, 0, pr.Policy.K)
	survived := make([]int, 0, pr.Policy.K)
	for i, rep := range pr.Replicas {
		if rep == "" {
			logger.Warnf("piece %s %d is not stored", pr.Name, i)
			continue
		}
		suc++
	}
	streamURL := ""
	if suc < int(pr.Policy.K) {
		er, err := GetEdge(baseUrl, auth, pr.Streamer)
		if err != nil {
			return pr.PieceCore, nil, err
		}
		streamURL = er.ExposeURL
		pr, err = GetPieceReceipt(streamURL, auth, name)
		if err != nil {
			return pr.PieceCore, nil, err
		}
	}

	suc = 0
	for i, rep := range pr.Replicas {
		val, err := DownloadReplica(baseUrl, streamURL, auth, rep, pr.StoredOn[i])
		if err != nil || len(val) == 0 {
			if i < int(pr.Policy.K) {
				need = append(need, i)
			}
			logger.Debugf("%s download fail: %v", rep, err)
			continue
		}

		res[i] = val
		suc++
		survived = append(survived, i)
		if suc >= int(pr.Policy.K) {
			break
		}
	}
	if suc < int(pr.Policy.K) {
		return pr.PieceCore, nil, fmt.Errorf("no enough replica")
	}
	// repair
	if len(need) > 0 {
		slen := len(res[survived[0]]) / bls.PadSize
		for _, v := range need {
			res[v] = make([]byte, 0, slen*bls.PadSize)
		}

		rs, err := erasure.NewRS(int(pr.Policy.N), int(pr.Policy.K))
		if err != nil {
			return pr.PieceCore, nil, err
		}
		re, err := rs.NewReconst(survived)
		if err != nil {
			return pr.PieceCore, nil, err
		}
		encoded := make([][]byte, int(pr.Policy.K))
		for i := 0; i < slen; i++ {
			for j := 0; j < int(pr.Policy.K); j++ {
				encoded[j] = res[survived[j]][i*bls.PadSize : (i+1)*bls.PadSize]
			}
			par, err := re.Encode(encoded, need)
			if err != nil {
				return pr.PieceCore, nil, err
			}
			for i, v := range need {
				res[v] = append(res[v], par[i]...)
			}
		}
	}

	pbyte := make([]byte, 0, pr.Size)
	for i := 0; i < int(pr.Policy.K); i++ {
		data, err := bls.Unpad(res[i])
		if err != nil {
			return pr.PieceCore, nil, err
		}
		pbyte = append(pbyte, data...)
	}

	return pr.PieceCore, pbyte[:pr.Size], nil
}

func DownloadReplica(baseUrl, streamUrl string, auth types.Auth, name string, addr common.Address) ([]byte, error) {
	if addr != types.EmptyAddr {
		res, err := DownloadReplicaFromStream(baseUrl, auth, name, addr)
		if err == nil {
			return res, nil
		}
	}

	return DownloadReplicaOrigin(streamUrl, auth, name)
}

func DownloadReplicaOrigin(baseUrl string, auth types.Auth, name string) ([]byte, error) {
	form := url.Values{}
	form.Set("name", name)
	form.Set("type", "replica")

	logger.Debug("download replica: ", name, " at:", baseUrl)
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancle()
	resByte, err := doRequest(ctx, baseUrl, "/v1/download", "", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	return resByte, nil
}

func DownloadReplicaFromStream(baseUrl string, auth types.Auth, name string, addr common.Address) ([]byte, error) {
	logger.Debug("download replica: ", name, " via stream")
	el, err := ListEdge(baseUrl, auth, types.StreamType)
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Set("name", name)
	form.Set("storedOn", addr.String())

	for _, er := range relayStreams(el.Edges, chaintype) {
		logger.Debug("download replica: ", name, " via stream: ", er.Name, " at: ", er.ExposeURL)
		ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
		defer cancle()
		resByte, err := doRequest(ctx, er.ExposeURL, "/v1/download", "", auth, strings.NewReader(form.Encode()))
		if err != nil {
			continue
		}

		return resByte, nil
	}
	return nil, fmt.Errorf("no avail stream")
}

func DownloadPieceAndSave(baseUrl string, auth types.Auth, com string, ks types.IPieceStore) error {
	if ks != nil {
		_, err := ks.GetPiece(context.TODO(), com, nil, types.Options{})
		if err == nil {
			return nil
		}
	}

	pc, resByte, err := DownloadPiece(baseUrl, auth, com)
	if err != nil {
		return err
	}

	if ks != nil {
		ks.PutPiece(context.TODO(), pc, resByte, true)
	}
	return nil
}

func CheckFile(baseUrl string, auth types.Auth, name string, ks types.IPieceStore) error {
	fr, err := GetFileReceipt(baseUrl, auth, name)
	if err != nil {
		return err
	}

	for _, com := range fr.Pieces {
		err = DownloadPieceAndSave(baseUrl, auth, com, ks)
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckFileParallel(baseUrl string, auth types.Auth, name string, parallel int, ks types.IPieceStore) error {
	logger.Debug("download piece in parallel: ", parallel)
	fr, err := GetFileReceipt(baseUrl, auth, name)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	sm := semaphore.NewWeighted(int64(parallel))
	for i, com := range fr.Pieces {
		err := sm.Acquire(context.TODO(), 1)
		if err != nil {
			return err
		}
		wg.Add(1)
		go func(ni int, com string, ks types.IPieceStore) {
			defer sm.Release(1)
			defer wg.Done()

			DownloadPieceAndSave(baseUrl, auth, com, ks)
		}(i, com, ks)
	}
	wg.Wait()

	return nil
}

// relayStreams keeps the registered streams a replica may be fetched through:
// active on-chain and, when the chain is known, on it. The registry is open,
// so this narrows who can answer; the bytes are still checked by the caller.
func relayStreams(edges []types.EdgeReceipt, chain string) []types.EdgeReceipt {
	out := make([]types.EdgeReceipt, 0, len(edges))
	for _, e := range edges {
		if !e.OnChain {
			continue
		}
		if chain != "" && e.ChainType != chain {
			continue
		}
		out = append(out, e)
	}
	return out
}

func Download(baseUrl string, auth types.Auth, name string, ks types.IPieceStore, w io.Writer) error {
	return DownloadOf(baseUrl, auth, name, types.EmptyAddr, ks, w)
}

// DownloadOf downloads owner's file (EmptyAddr: any owner's file of that name)
// and checks the result against the sha256 in its receipt. Pieces fetched
// from the network go into ks only after the whole file checks out, so data
// from a lying store or stream is never cached. Bytes are written to w as
// they arrive; on a mismatch the error comes at the end and w holds bad data,
// which the caller must discard.
func DownloadOf(baseUrl string, auth types.Auth, name string, owner common.Address, ks types.IPieceStore, w io.Writer) error {
	fr, err := GetFileReceiptOf(baseUrl, name, owner)
	if err != nil {
		return err
	}

	h := sha256.New()
	out := io.MultiWriter(w, h)
	type fetched struct {
		pc   types.PieceCore
		data []byte
	}
	var toCache []fetched

	for _, com := range fr.Pieces {
		if ks != nil {
			var b bytes.Buffer
			_, err := ks.GetPiece(context.TODO(), com, &b, types.Options{})
			if err == nil {
				out.Write(b.Bytes())
				continue
			}
		}

		pc, resByte, err := DownloadPiece(baseUrl, auth, com)
		if err != nil {
			return err
		}
		if ks != nil {
			toCache = append(toCache, fetched{pc, resByte})
		}
		out.Write(resByte)
	}

	if got := hex.EncodeToString(h.Sum(nil)); fr.Hash != "" && !strings.EqualFold(got, fr.Hash) {
		return fmt.Errorf("file %s: downloaded data hashes to %s, receipt says %s", name, got, fr.Hash)
	}
	for _, f := range toCache {
		ks.PutPiece(context.TODO(), f.pc, f.data, true)
	}
	return nil
}

func DownloadParallel(baseUrl string, auth types.Auth, name string, parallel int, ks types.IPieceStore, w io.Writer) error {
	if ks != nil {
		err := CheckFileParallel(baseUrl, auth, name, parallel, ks)
		if err != nil {
			return err
		}
	}

	return Download(baseUrl, auth, name, ks, w)
}

// todo: handle size > piece size
func DownloadWSize(baseUrl string, auth types.Auth, name string, ks types.IPieceStore, w io.Writer, start, size int64) error {
	logger.Debugf("download file %s %d %d ", name, start, size)
	fr, err := GetFileReceipt(baseUrl, auth, name)
	if err != nil {
		return err
	}

	sstart := int64(0)
	pstart := int64(0)
	for _, com := range fr.Pieces {
		pr, err := GetPieceReceipt(baseUrl, auth, com)
		if err != nil {
			return err
		}

		tmp := sstart + pr.Size
		if start >= tmp {
			sstart = tmp
			continue
		}
		if sstart >= start+size {
			return nil
		}
		pstart = start - sstart
		sstart = tmp

		if ks != nil {
			var b bytes.Buffer
			_, err := ks.GetPiece(context.TODO(), com, &b, types.Options{})
			if err == nil {
				_, err = w.Write(b.Bytes())
				return err
			}
		}

		pc, resByte, err := DownloadPiece(baseUrl, auth, com)
		if err != nil {
			return err
		}

		if ks != nil {
			err = ks.PutPiece(context.TODO(), pc, resByte, true)
			if err != nil {
				return err
			}
			var b bytes.Buffer
			_, err = ks.GetPiece(context.TODO(), com, &b, types.Options{
				UserDefined: map[string]string{
					"start": strconv.FormatInt(pstart, 10),
					"size":  strconv.FormatInt(size, 10),
				},
			})
			if err != nil {
				return err
			}
			_, err = w.Write(b.Bytes())
			return err
		}
	}

	return nil
}
