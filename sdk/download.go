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
		spr.PieceCore = pr.PieceCore
		if err := checkReplicaList(spr); err != nil {
			return pr.PieceCore, nil, err
		}
		pr.Replicas, pr.StoredOn, err = mergeStreamerReplicas(pr, spr)
		if err != nil {
			return pr.PieceCore, nil, err
		}
	}

	// every shard of the piece has the length the encoder gave it, fixed by the
	// gateway's Size and K: a store cannot set it by answering first
	shardLen, err := pieceShardLen(pr.PieceCore)
	if err != nil {
		return pr.PieceCore, nil, err
	}

	suc = 0
	for i, rep := range pr.Replicas {
		val, err := downloadReplica(baseUrl, streamURL, auth, rep, pr.StoredOn[i], int64(shardLen))
		if err == nil && len(val) != shardLen {
			err = fmt.Errorf("replica %s has %d bytes, want %d", rep, len(val), shardLen)
		}
		if err != nil {
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
		slen := shardLen / bls.PadSize
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

// pieceShardLen is the byte length of each of a piece's N shards, as the
// stream encodes it (stream/file.go Put: shardLen = 1+(n-1)/(31K) field
// elements of PadSize bytes). Size must be a piece size the policy allows.
func pieceShardLen(pc types.PieceCore) (int, error) {
	if err := pc.Policy.Check(); err != nil {
		return 0, err
	}
	if pc.Size <= 0 || pc.Size > MaxPieceSize(pc.Policy) {
		return 0, fmt.Errorf("piece %s: size %d outside 1..%d", pc.Name, pc.Size, MaxPieceSize(pc.Policy))
	}
	elems := 1 + (pc.Size-1)/(bls.UnPadSize*int64(pc.Policy.K))
	return int(elems * bls.PadSize), nil
}

// mergeStreamerReplicas fills the slots the gateway has no replica for from
// the streamer's receipt. A slot the gateway knows (the replica is on chain)
// keeps the gateway's name and store; the streamer must name the same replica
// there or the answer is refused. A name the streamer adds must be a
// commitment's hex (48 bytes) and appear in no other slot, so a streamer
// cannot serve one shard under two slots.
//
// Names of slots that are not on chain yet cannot be checked against anything
// a light client holds (binding a name to its bytes needs the SRS); the
// rebuilt piece is still only as good as the caller's own check (the file
// hash for file downloads).
func mergeStreamerReplicas(gw, st types.PieceReceipt) ([]string, []common.Address, error) {
	n := int(gw.Policy.N)
	names := make([]string, n)
	stores := make([]common.Address, n)
	seen := make(map[string]int, n)
	for i := 0; i < n && i < len(gw.Replicas); i++ {
		if gw.Replicas[i] == "" {
			continue
		}
		names[i], stores[i] = gw.Replicas[i], gw.StoredOn[i]
		seen[strings.ToLower(gw.Replicas[i])] = i
	}
	for i, name := range st.Replicas {
		if name == "" {
			continue
		}
		if names[i] != "" {
			if !strings.EqualFold(names[i], name) {
				return nil, nil, fmt.Errorf("piece %s slot %d: streamer names replica %s, the chain record has %s", gw.Name, i, name, names[i])
			}
			continue
		}
		if b, err := hex.DecodeString(name); err != nil || len(b) != bls.G1Size {
			return nil, nil, fmt.Errorf("piece %s slot %d: streamer names %q, not a commitment", gw.Name, i, name)
		}
		if j, dup := seen[strings.ToLower(name)]; dup {
			return nil, nil, fmt.Errorf("piece %s: streamer names replica %s for slots %d and %d", gw.Name, name, j, i)
		}
		seen[strings.ToLower(name)] = i
		names[i], stores[i] = name, st.StoredOn[i]
	}
	return names, stores, nil
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
	return downloadReplica(baseUrl, streamUrl, auth, name, addr, -1)
}

// downloadReplica is DownloadReplica reading at most max bytes of an answer
// (max < 0: unbounded); a longer one is an error, not a buffer to hold.
func downloadReplica(baseUrl, streamUrl string, auth types.Auth, name string, addr common.Address, max int64) ([]byte, error) {
	if addr != types.EmptyAddr {
		res, err := downloadReplicaFromStream(baseUrl, auth, name, addr, max)
		if err == nil {
			return res, nil
		}
	}

	return downloadReplicaOrigin(streamUrl, auth, name, max)
}

func DownloadReplicaOrigin(baseUrl string, auth types.Auth, name string) ([]byte, error) {
	return downloadReplicaOrigin(baseUrl, auth, name, -1)
}

func downloadReplicaOrigin(baseUrl string, auth types.Auth, name string, max int64) ([]byte, error) {
	form := url.Values{}
	form.Set("name", name)
	form.Set("type", "replica")

	logger.Debug("download replica: ", name, " at:", baseUrl)
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancle()
	resByte, err := doRequestLimit(ctx, baseUrl, "/v1/download", "", auth, strings.NewReader(form.Encode()), max)
	if err != nil {
		return nil, err
	}

	return resByte, nil
}

func DownloadReplicaFromStream(baseUrl string, auth types.Auth, name string, addr common.Address) ([]byte, error) {
	return downloadReplicaFromStream(baseUrl, auth, name, addr, -1)
}

func downloadReplicaFromStream(baseUrl string, auth types.Auth, name string, addr common.Address, max int64) ([]byte, error) {
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
		resByte, err := doRequestLimit(ctx, er.ExposeURL, "/v1/download", "", auth, strings.NewReader(form.Encode()), max)
		if err != nil {
			continue
		}

		return resByte, nil
	}
	return nil, fmt.Errorf("no avail stream")
}

// DownloadPieceAndSave rebuilds piece com from store nodes and keeps it in ks.
//
// Deprecated: a bare piece has nothing a light client can check its bytes
// against (the commitment needs the full SRS), so this caches whatever the
// stores returned, and every later reader of ks trusts it. Use
// DownloadPieceAndSaveVerified with a real check (a file hash, a volume
// digest), CheckFileParallelOf for files, or DownloadPiece without caching.
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
// pieces into ks. Nothing that fails the check is cached. When the check fails
// with pieces taken from ks (or a cached piece cannot be read), those pieces
// are evicted and the whole file is fetched once more without the cache.
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
	err = checkFileParallel(baseUrl, auth, fr, parallel, ks, true)
	if errors.Is(err, errBadCache) {
		logger.Warnf("%v; fetching the file again without the cache", err)
		err = checkFileParallel(baseUrl, auth, fr, parallel, ks, false)
	}
	return err
}

// errBadCache marks a download that failed because of what the piece cache
// held: the caller retries once without reading the cache.
var errBadCache = errors.New("cached pieces evicted")

// evictCached drops the cached copies of the file's pieces at idx (blobs
// only; DeleteData). It reports how many it dropped.
func evictCached(ks types.IPieceStore, fr types.FileReceipt, idx []int) int {
	n := 0
	for _, i := range idx {
		if err := ks.DeleteData(context.TODO(), fr.Pieces[i]); err != nil {
			logger.Warnf("evict cached piece %s: %v", fr.Pieces[i], err)
			continue
		}
		n++
	}
	return n
}

// readCached reads a whole cached piece, or fails without partial output.
func readCached(ks types.IPieceStore, com string) ([]byte, error) {
	var b bytes.Buffer
	if _, err := ks.GetPiece(context.TODO(), com, &b, types.Options{}); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// checkFileParallel is one pass of CheckFileParallelOf, reading pieces from ks
// only when useCache is set.
func checkFileParallel(baseUrl string, auth types.Auth, fr types.FileReceipt, parallel int, ks types.IPieceStore, useCache bool) error {
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
		if useCache && ks != nil {
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

	var cached []int
	for i, s := range slots {
		if s.cached {
			cached = append(cached, i)
		}
	}

	h := sha256.New()
	for i, s := range slots {
		if s.err != nil {
			return fmt.Errorf("piece %s: %w", fr.Pieces[i], s.err)
		}
		if s.cached {
			b, err := readCached(ks, fr.Pieces[i])
			if err != nil {
				// indexed but unreadable (e.g. its blob was evicted): refetch
				return fmt.Errorf("file %s: read cached piece %s: %v: %w", fr.Name, fr.Pieces[i], err, errBadCache)
			}
			h.Write(b)
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
		if evictCached(ks, fr, cached) > 0 {
			return fmt.Errorf("file %s: data hashes to %s, receipt says %s, %d pieces came from the cache: %w", fr.Name, got, fr.Hash, len(cached), errBadCache)
		}
		return fmt.Errorf("file %s: data hashes to %s, receipt says %s: %w", fr.Name, got, fr.Hash, ErrFileHashMismatch)
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
// and checks the result against the sha256 in its receipt.
//
// Without a cache (ks == nil) bytes are written to w as they arrive; on a
// mismatch the error comes at the end and w holds bad data, which the caller
// must discard.
//
// With a cache, the file is checked before anything is written to w: cached
// pieces are read from ks, the rest fetched, and only a file that hashes right
// is written out, after which the fetched pieces are cached. If the check
// fails and some pieces came from ks, those are evicted (a bad piece must not
// keep failing every later read of every file that holds it) and the whole
// file is fetched once more from the network.
func DownloadOf(baseUrl string, auth types.Auth, name string, owner common.Address, ks types.IPieceStore, w io.Writer) error {
	fr, err := GetFileReceiptOf(baseUrl, name, owner)
	if err != nil {
		return err
	}
	if err := checkReceiptHash(fr); err != nil {
		return err
	}
	if ks == nil {
		return streamFile(baseUrl, auth, fr, w)
	}
	err = downloadChecked(baseUrl, auth, fr, ks, w, true)
	if errors.Is(err, errBadCache) {
		logger.Warnf("%v; fetching the file again without the cache", err)
		err = downloadChecked(baseUrl, auth, fr, ks, w, false)
	}
	return err
}

// streamFile writes the file's pieces to w as they are fetched and checks the
// hash at the end.
func streamFile(baseUrl string, auth types.Auth, fr types.FileReceipt, w io.Writer) error {
	h := sha256.New()
	out := io.MultiWriter(w, h)
	for _, com := range fr.Pieces {
		_, resByte, err := DownloadPiece(baseUrl, auth, com)
		if err != nil {
			return err
		}
		if _, err := out.Write(resByte); err != nil {
			return err
		}
	}
	if got := hex.EncodeToString(h.Sum(nil)); !strings.EqualFold(got, fr.Hash) {
		return fmt.Errorf("file %s: downloaded data hashes to %s, receipt says %s: %w", fr.Name, got, fr.Hash, ErrFileHashMismatch)
	}
	return nil
}

// downloadChecked is one cache-backed pass of DownloadOf: hash every piece
// (from ks when useCache, else fetched), then write the file to w only if it
// matches. Fetched pieces are held in memory until then (as before); cached
// ones are read again for the write and hashed again, so a cache entry that
// changed in between is reported, not passed off as checked.
func downloadChecked(baseUrl string, auth types.Auth, fr types.FileReceipt, ks types.IPieceStore, w io.Writer, useCache bool) error {
	type part struct {
		cached bool
		pc     types.PieceCore
		data   []byte
	}
	parts := make([]part, len(fr.Pieces))
	var cached []int
	h := sha256.New()
	for i, com := range fr.Pieces {
		if useCache {
			if b, err := readCached(ks, com); err == nil {
				parts[i].cached = true
				cached = append(cached, i)
				h.Write(b)
				continue
			}
		}
		pc, resByte, err := DownloadPiece(baseUrl, auth, com)
		if err != nil {
			return err
		}
		parts[i] = part{pc: pc, data: resByte}
		h.Write(resByte)
	}
	want := fr.Hash
	if got := hex.EncodeToString(h.Sum(nil)); !strings.EqualFold(got, want) {
		if evictCached(ks, fr, cached) > 0 {
			return fmt.Errorf("file %s: data hashes to %s, receipt says %s, %d pieces came from the cache: %w", fr.Name, got, want, len(cached), errBadCache)
		}
		return fmt.Errorf("file %s: downloaded data hashes to %s, receipt says %s: %w", fr.Name, got, want, ErrFileHashMismatch)
	}

	h.Reset()
	out := io.MultiWriter(w, h)
	for i, p := range parts {
		data := p.data
		if p.cached {
			b, err := readCached(ks, fr.Pieces[i])
			if err != nil {
				return fmt.Errorf("file %s: re-read cached piece %s: %w", fr.Name, fr.Pieces[i], err)
			}
			data = b
		}
		if _, err := out.Write(data); err != nil {
			return err
		}
	}
	if got := hex.EncodeToString(h.Sum(nil)); !strings.EqualFold(got, want) {
		// a cached piece changed between the check and the write
		return fmt.Errorf("file %s: written data hashes to %s, receipt says %s: %w", fr.Name, got, want, ErrFileHashMismatch)
	}
	for _, p := range parts {
		if p.cached {
			continue
		}
		if err := ks.PutPiece(context.TODO(), p.pc, p.data, true); err != nil {
			logger.Warnf("cache piece %s: %v", p.pc.Name, err)
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

// DownloadWSize writes size bytes of file name from offset start, as far as
// they lie in one piece (todo: a range spanning pieces). A range cannot be
// checked against the file hash, so a piece fetched here is never put into ks:
// it would be served to later full-file reads as if checked. ks is only read.
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
			if b, err := readCached(ks, com); err == nil {
				_, err = w.Write(byteRange(b, pstart, size))
				return err
			}
		}

		_, resByte, err := DownloadPiece(baseUrl, auth, com)
		if err != nil {
			return err
		}
		_, err = w.Write(byteRange(resByte, pstart, size))
		return err
	}

	return nil
}

// byteRange is b[start:start+size], clamped to b.
func byteRange(b []byte, start, size int64) []byte {
	if start < 0 || start >= int64(len(b)) || size <= 0 {
		return nil
	}
	end := start + size
	if end > int64(len(b)) || end < start {
		end = int64(len(b))
	}
	return b[start:end]
}
