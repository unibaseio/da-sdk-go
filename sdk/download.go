package sdk

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
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

	if pr.Name != name {
		return pr.PieceCore, nil, fmt.Errorf("receipt for piece %s names %s", name, pr.Name)
	}
	if err := pr.Policy.Check(); err != nil {
		return pr.PieceCore, nil, err
	}
	if err := checkReplicaList(pr); err != nil {
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
		// Too few stored replicas: the piece's streamer may still hold its
		// staged shards. Only its replica list is taken from it — name, policy
		// and size stay as the gateway recorded them — and only from an
		// on-chain stream of this chain (anyone can name themselves streamer).
		er, err := GetEdge(baseUrl, auth, pr.Streamer)
		if err != nil {
			return pr.PieceCore, nil, err
		}
		if len(relayStreams([]types.EdgeReceipt{er}, chaintype)) == 0 {
			return pr.PieceCore, nil, fmt.Errorf("streamer %s of piece %s is not an on-chain stream of this chain", pr.Streamer, name)
		}
		streamURL = er.ExposeURL
		spr, err := GetPieceReceipt(streamURL, auth, name)
		if err != nil {
			return pr.PieceCore, nil, err
		}
		if spr.Name != name {
			return pr.PieceCore, nil, fmt.Errorf("streamer answered for %s, asked for %s", spr.Name, name)
		}
		pr.Replicas, pr.StoredOn = spr.Replicas, spr.StoredOn
		if err := checkReplicaList(pr); err != nil {
			return pr.PieceCore, nil, err
		}
	}

	suc = 0
	shardLen := -1
	for i, rep := range pr.Replicas {
		val, err := DownloadReplica(baseUrl, streamURL, auth, rep, pr.StoredOn[i])
		if err == nil && len(val) > 0 {
			// shards of one piece are equal-length multiples of PadSize
			if len(val)%bls.PadSize != 0 || (shardLen >= 0 && len(val) != shardLen) {
				err = fmt.Errorf("replica %s has a malformed length %d", rep, len(val))
			} else {
				shardLen = len(val)
			}
		}
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
	if pr.Size < 0 || int64(len(pbyte)) < pr.Size {
		return pr.PieceCore, nil, fmt.Errorf("piece %s rebuilt to %d bytes, receipt says %d", name, len(pbyte), pr.Size)
	}

	return pr.PieceCore, pbyte[:pr.Size], nil
}

// checkReplicaList makes a receipt's replica list safe to index: one name and
// one storing address per slot, at most N slots.
func checkReplicaList(pr types.PieceReceipt) error {
	if len(pr.Replicas) > int(pr.Policy.N) || len(pr.StoredOn) < len(pr.Replicas) {
		return fmt.Errorf("piece %s: %d replicas / %d stores for N=%d", pr.Name, len(pr.Replicas), len(pr.StoredOn), pr.Policy.N)
	}
	return nil
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

// DownloadPieceAndSave rebuilds piece com from store nodes and keeps it in ks.
// ⚠️ A bare piece has nothing a light client can check its bytes against (the
// commitment needs the full SRS), so this caches what the stores returned.
// Callers that can check the data (a file hash, a volume digest) should use
// DownloadPieceAndSaveVerified; file downloads use CheckFileParallelOf.
func DownloadPieceAndSave(baseUrl string, auth types.Auth, com string, ks types.IPieceStore) error {
	return DownloadPieceAndSaveVerified(baseUrl, auth, com, ks, nil)
}

// DownloadPieceAndSaveVerified is DownloadPieceAndSave that runs verify (when
// non-nil) on the rebuilt piece and caches it only if verify accepts it.
func DownloadPieceAndSaveVerified(baseUrl string, auth types.Auth, com string, ks types.IPieceStore, verify func(types.PieceCore, []byte) error) error {
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
	if verify != nil {
		if err := verify(pc, resByte); err != nil {
			return fmt.Errorf("piece %s: %w", com, err)
		}
	}

	if ks != nil {
		return ks.PutPiece(context.TODO(), pc, resByte, true)
	}
	return nil
}

// ErrFileHashMismatch is wrapped by the errors a download returns when the
// rebuilt file does not hash to its receipt's sha256.
var ErrFileHashMismatch = errors.New("downloaded data does not match the file hash")

// checkReceiptHash rejects a file receipt whose Hash cannot be checked: a
// blank hash would otherwise let any data through.
func checkReceiptHash(fr types.FileReceipt) error {
	if hb, err := hex.DecodeString(fr.Hash); err != nil || len(hb) != sha256.Size {
		return fmt.Errorf("file %s: receipt hash %q is not a sha256; the download cannot be verified", fr.Name, fr.Hash)
	}
	return nil
}

func CheckFile(baseUrl string, auth types.Auth, name string, ks types.IPieceStore) error {
	return CheckFileParallelOf(baseUrl, auth, name, types.EmptyAddr, 1, ks)
}

func CheckFileParallel(baseUrl string, auth types.Auth, name string, parallel int, ks types.IPieceStore) error {
	return CheckFileParallelOf(baseUrl, auth, name, types.EmptyAddr, parallel, ks)
}

// CheckFileParallelOf fetches owner's file (EmptyAddr: any owner's file of
// that name) piece by piece, parallel at a time, into temporary files, checks
// the whole file against the receipt's sha256, and only then puts the fetched
// pieces into ks. Nothing that fails the check is cached.
func CheckFileParallelOf(baseUrl string, auth types.Auth, name string, owner common.Address, parallel int, ks types.IPieceStore) error {
	logger.Debug("download piece in parallel: ", parallel)
	if parallel < 1 {
		parallel = 1
	}
	fr, err := GetFileReceiptOf(baseUrl, name, owner)
	if err != nil {
		return err
	}
	if err := checkReceiptHash(fr); err != nil {
		return err
	}

	dir, err := os.MkdirTemp("", "da-check-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	type slot struct {
		cached bool
		pc     types.PieceCore
		path   string
		err    error
	}
	slots := make([]slot, len(fr.Pieces))

	var wg sync.WaitGroup
	sm := semaphore.NewWeighted(int64(parallel))
	for i, com := range fr.Pieces {
		if ks != nil {
			if _, err := ks.GetPiece(context.TODO(), com, nil, types.Options{}); err == nil {
				slots[i].cached = true
				continue
			}
		}
		if err := sm.Acquire(context.TODO(), 1); err != nil {
			return err
		}
		wg.Add(1)
		go func(i int, com string) {
			defer sm.Release(1)
			defer wg.Done()
			pc, data, err := DownloadPiece(baseUrl, auth, com)
			if err == nil {
				p := filepath.Join(dir, strconv.Itoa(i))
				err = os.WriteFile(p, data, 0o600)
				slots[i].path = p
			}
			slots[i].pc, slots[i].err = pc, err
		}(i, com)
	}
	wg.Wait()

	h := sha256.New()
	for i, s := range slots {
		if s.err != nil {
			return fmt.Errorf("piece %s: %w", fr.Pieces[i], s.err)
		}
		if s.cached {
			if _, err := ks.GetPiece(context.TODO(), fr.Pieces[i], h, types.Options{}); err != nil {
				return err
			}
			continue
		}
		f, err := os.Open(s.path)
		if err != nil {
			return err
		}
		_, err = io.Copy(h, f)
		f.Close()
		if err != nil {
			return err
		}
	}
	if got := hex.EncodeToString(h.Sum(nil)); !strings.EqualFold(got, fr.Hash) {
		return fmt.Errorf("file %s: data hashes to %s, receipt says %s: %w", name, got, fr.Hash, ErrFileHashMismatch)
	}

	if ks == nil {
		return nil
	}
	for _, s := range slots {
		if s.cached {
			continue
		}
		data, err := os.ReadFile(s.path)
		if err != nil {
			return err
		}
		if err := ks.PutPiece(context.TODO(), s.pc, data, true); err != nil {
			return err
		}
	}
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
	if err := checkReceiptHash(fr); err != nil {
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

	if got := hex.EncodeToString(h.Sum(nil)); !strings.EqualFold(got, fr.Hash) {
		return fmt.Errorf("file %s: downloaded data hashes to %s, receipt says %s: %w", name, got, fr.Hash, ErrFileHashMismatch)
	}
	for _, f := range toCache {
		if err := ks.PutPiece(context.TODO(), f.pc, f.data, true); err != nil {
			logger.Warnf("cache piece %s: %v", f.pc.Name, err)
		}
	}
	return nil
}

func DownloadParallel(baseUrl string, auth types.Auth, name string, parallel int, ks types.IPieceStore, w io.Writer) error {
	return DownloadParallelOf(baseUrl, auth, name, types.EmptyAddr, parallel, ks, w)
}

// DownloadParallelOf is DownloadOf that, with a cache, first fetches the
// pieces parallel at a time and caches them once the whole file checks out.
// A failed prefetch falls back to the serial download, except on a hash
// mismatch, which is returned.
func DownloadParallelOf(baseUrl string, auth types.Auth, name string, owner common.Address, parallel int, ks types.IPieceStore, w io.Writer) error {
	if ks != nil {
		err := CheckFileParallelOf(baseUrl, auth, name, owner, parallel, ks)
		if errors.Is(err, ErrFileHashMismatch) {
			return err
		}
		if err != nil {
			logger.Warnf("parallel fetch of %s: %v; downloading serially", name, err)
		}
	}

	return DownloadOf(baseUrl, auth, name, owner, ks, w)
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
