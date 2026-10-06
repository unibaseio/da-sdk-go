package hub

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	contract "github.com/unibaseio/da-sdk-go/contract/v2"
	"github.com/unibaseio/da-sdk-go/lib/env"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
)

// logFSWriteEx writes one object to the owner's logfs + indexes it.
//   - kind:  bucket scenario ("memory"/"model"/"dataset").
//   - large: passthrough-large policy — seal any pending small-write volume
//     first, then seal this object into its own volume (isolated, uploads
//     promptly). large=false keeps the memory coalescing behavior.
func (s *Server) logFSWriteEx(addr string, bucket string, key string, kind string, large bool, r io.Reader) (types.MemeMeta, error) {
	size := int64(-1)
	if sz, ok := r.(interface{ Size() int64 }); ok { // multipart file parts
		size = sz.Size()
	}
	rbytes, err := readAllSized(r, size)
	if err != nil {
		return types.MemeMeta{}, err
	}
	return s.logFSWriteData(addr, bucket, key, kind, large, rbytes)
}

// maxPrealloc bounds readAllSized's up-front allocation.
const maxPrealloc = 1 << 20

// readAllSized reads r to EOF into a buffer preallocated from a size hint
// (Content-Length, multipart part size; <=0 = unknown), so a body is not
// re-copied while the buffer grows. size is only a hint: a short or long read
// is still read exactly. At most maxPrealloc is allocated up front; a larger
// body grows as its bytes arrive.
func readAllSized(r io.Reader, size int64) ([]byte, error) {
	if size <= 0 {
		return io.ReadAll(r)
	}
	if lim := env.Int64("HUB_MAX_MULTIPART_BYTES", defaultMaxMultipartBytes); size > lim {
		size = lim
	}
	// The size is the client's claim: preallocate only a bounded part of it,
	// so a declared length with no body behind it cannot reserve the whole
	// cap per connection. Larger bodies grow as their bytes actually arrive.
	if size > maxPrealloc {
		size = maxPrealloc
	}
	// +MinRead: bytes.Buffer.ReadFrom wants that much free room per read, so an
	// exact-size buffer would still regrow (copy) at the end
	buf := bytes.NewBuffer(make([]byte, 0, size+bytes.MinRead))
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// logFSWriteData is logFSWriteEx for bytes already in memory (no extra copy).
func (s *Server) logFSWriteData(addr string, bucket string, key string, kind string, large bool, rbytes []byte) (types.MemeMeta, error) {
	var err error
	if addr == "" {
		addr = s.local.String()
	} else {
		// Canonicalize the client-provided owner to lowercase so the same
		// wallet always lands in one namespace regardless of the case it was
		// sent in (EIP-55 checksum vs lowercase).
		addr = CanonOwner(addr)
	}

	if bucket == "" {
		bucket = addr
	}

	err = s.addBucket(addr, bucket, kind)
	if err != nil {
		return types.MemeMeta{}, err
	}

	fs, err := s.getFS(addr, true)
	if err != nil {
		return types.MemeMeta{}, err
	}

	// passthrough-large: isolate this object — seal any pending small-write
	// volume first so the big file doesn't share a volume with tiny needles.
	if large {
		if err := fs.Roll(); err != nil {
			return types.MemeMeta{}, err
		}
	}

	err = fs.Put([]byte(key), rbytes)
	if err != nil {
		return types.MemeMeta{}, err
	}

	// seal the large object into its own completed volume so it uploads promptly.
	if large {
		if err := fs.Roll(); err != nil {
			return types.MemeMeta{}, err
		}
	}

	// Drop any stale "missing" marker + cached value so this key reflects the
	// new write immediately.
	s.readCache.del(addr, key)

	lm, err := fs.GetMeta([]byte(key))
	if err != nil {
		return types.MemeMeta{}, err
	}

	if err := s.addNeedle(addr, bucket, key, lm.Index, lm.Start, lm.Size); err != nil {
		return types.MemeMeta{}, err
	}

	// wake the drain loop (coalescing, never blocks the writer)
	select {
	case s.uploadNotify <- struct{}{}:
	default:
	}

	mm := types.MemeMeta{
		File:  fmt.Sprintf("%s/%d.log", addr, lm.Index),
		Start: lm.Start,
		Size:  lm.Size,
	}

	return mm, nil
}

// logFSReadAt reads one object by the location its index row records
// (volume, start, size). Reading by key alone conflates objects sharing a key:
// the logfs key index is per owner, not per bucket, so the newest write of a
// key in any bucket wins. When the key's meta still points at this row (the
// usual case) the read is hash-checked through it.
func (s *Server) logFSReadAt(owner, key string, file, start, size uint64, w io.Writer) (int64, error) {
	ck := locKey(file, start, size) // location, not key
	if wbytes, ok := s.readCache.get(owner, ck); ok {
		n, err := w.Write(wbytes)
		return int64(n), err
	}
	fs, err := s.readFS(owner)
	if err != nil {
		return 0, err
	}
	var wbytes []byte
	if lm, merr := fs.GetMeta([]byte(key)); merr == nil && lm.Index == file && lm.Start == start && lm.Size == size {
		wbytes, err = fs.GetData(lm)
	} else {
		wbytes, err = fs.GetDataAt(file, start, size)
	}
	if err != nil {
		return 0, err
	}
	s.readCache.put(owner, ck, wbytes)
	n, err := w.Write(wbytes)
	return int64(n), err
}

// locKey is the read-cache name of an object by its location in the owner's
// LogFS (volume, start, size) — unique per owner, unlike the key.
func locKey(file, start, size uint64) string {
	return fmt.Sprintf("@%d/%d/%d", file, start, size)
}

func (s *Server) load() error {
	fspath := filepath.Join(s.rp.Path(), LOGFS)
	fs, err := logfs.New(s.rp.MetaStore(), fspath, s.local.String(), s.local.String(), s.logFSOptions(s.local.String())...)
	if err != nil {
		return err
	}
	s.lfs.Store(s.local.String(), fs)

	dsKey := types.NewKey(types.DsLogFS, LOGINST)
	val, err := s.rp.MetaStore().Get(dsKey)
	if err == nil && len(val) == 4 {
		s.fscnt = binary.BigEndian.Uint32(val)
	}

	if s.fscnt == 0 {
		s.fscnt = 1
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, s.fscnt)
		s.rp.MetaStore().Put(dsKey, buf)

		dsKey := types.NewKey(types.DsLogFS, LOGINST, 0)
		s.rp.MetaStore().Put(dsKey, []byte(s.local.String()))
	}
	/*
		for i := uint32(0); i < s.fscnt; i++ {
			dsKey := types.NewKey(types.DsLogFS, LOGINST, i)
			val, err := s.rp.MetaStore().Get(dsKey)
			if err != nil {
				break
			}

			s.addAccount(string(val))
		}
	*/
	logger.Infof("load log inst: %d", s.fscnt)
	return nil
}

func (s *Server) uploadTo() {
	sk := s.rp.Key().Export().PrivateKey
	au, err := key.BuildAuth(sk, []byte("upload"))
	if err != nil {
		panic(err)
	}

	policy := types.Policy{
		N: uint8(env.Int("HUB_RS_N", 6)),
		K: uint8(env.Int("HUB_RS_K", 4)),
	}

	cm, err := contract.NewContractManage(sk, s.rp.Repo().Config().Chain.Type)
	if err != nil {
		panic(err)
	}

	tick := time.Duration(env.Int("HUB_UPLOAD_TICK_SEC", 30)) * time.Second
	ticker := time.NewTicker(tick)
	defer ticker.Stop()
	for {
		// event-driven: drain on a new write, with a periodic fallback tick.
		select {
		case <-s.shutdownChan:
			return
		case <-s.uploadNotify:
		case <-ticker.C:
		}
		if err := cm.CheckBalance(au.Addr); err != nil {
			logger.Warnf("upload: balance check failed: %v", err)
			continue
		}
		// a fresh header per pass: streams bound how old a signature may be
		if fresh, err := key.BuildAuth(sk, []byte("upload")); err == nil {
			au = fresh
		}
		s.drainAll(cm, au, policy)
	}
}

// drainAll fans the per-owner drains out across a bounded worker pool. Each
// owner is an independent log instance, so they run concurrently; concurrent
// AddPiece is safe via the serialized-nonce manager (one cm shared). Within a
// single owner, drainInstance keeps volumes ordered — the per-owner offset
// advances sequentially. Worker count: HUB_UPLOAD_WORKERS (default NumCPU).
func (s *Server) drainAll(cm *contract.ContractManage, au types.Auth, policy types.Policy) {
	n := s.fscntGet()

	workers := env.Int("HUB_UPLOAD_WORKERS", runtime.NumCPU())
	if workers < 1 {
		workers = 1
	}
	if uint32(workers) > n {
		workers = int(n)
	}
	logger.Infof("upload drain: %d owners, %d workers", n, workers)

	sem := make(chan struct{}, workers)
	var wg sync.WaitGroup
	for i := uint32(0); i < n; i++ {
		sem <- struct{}{}
		wg.Add(1)
		go func(idx uint32) {
			defer wg.Done()
			defer func() { <-sem }()
			s.drainInstance(cm, au, policy, idx)
		}(i)
	}
	wg.Wait()
}

// getLFS returns the in-memory log instance for an owner (nil if not loaded).
func (s *Server) getLFS(addr string) *logfs.LogFS {
	if v, ok := s.lfs.Load(addr); ok {
		return v.(*logfs.LogFS)
	}
	return nil
}

var (
	alreadyHasPieceRe   = regexp.MustCompile(`already has piece: ([0-9a-fA-F]+)`)
	alreadyHasReplicaRe = regexp.MustCompile(`already has replica: ([0-9a-fA-F]+)`)
)

// parseAlreadyHasPiece pulls the piece name out of a stream "already has piece: <hex>"
// error (returns "" for anything else).
func parseAlreadyHasPiece(msg string) string {
	if m := alreadyHasPieceRe.FindStringSubmatch(msg); m != nil {
		return m[1]
	}
	return ""
}

// parseAlreadyHasReplica pulls the replica name out of "already has replica: <hex>".
func parseAlreadyHasReplica(msg string) string {
	if m := alreadyHasReplicaRe.FindStringSubmatch(msg); m != nil {
		return m[1]
	}
	return ""
}

// pieceOfReplica resolves a staged replica name to the piece it belongs to, by
// asking a stream for the replica receipt (which carries .Piece). "" if unknown.
func (s *Server) pieceOfReplica(au types.Auth, rn string) string {
	er, err := sdk.ListEdge(sdk.ServerURL, au, types.StreamType)
	if err != nil {
		return ""
	}
	for _, st := range er.Edges {
		if rr, err := sdk.GetReplicaReceipt(st.ExposeURL, au, rn); err == nil && rr.Piece != "" {
			return rr.Piece
		}
	}
	return ""
}

// recoverStagedPiece registers an already-staged piece on-chain (if not yet) and
// records its volume, so a vol whose earlier AddPiece failed still commits instead
// of being stuck "staged" forever. Mirrors the path-1 (GetFileReceipt) recovery,
// but keyed by the piece name from the "already has piece" error. Returns true when
// the vol is committed (or already on-chain) so the caller advances the offset.
// stagedPieceCore is what the hub registers on chain for a piece it staged
// earlier, from a stream's receipt for it. Only what the hub can check is
// kept: the drain's own policy, the piece name it asked for, the stream that
// holds the piece, and the size of the local volume (a volume, at most 31 MiB,
// is always a single piece). Price and expiry are left to AddPiece's defaults,
// exactly as on the normal path: the receipt's would decide the bond the hub
// locks, and a stream could inflate them (2026-10-06 audit, N10).
func stagedPieceCore(pr types.PieceReceipt, stream common.Address, pn string, policy types.Policy, volSize int64) (types.PieceCore, error) {
	switch {
	case !strings.EqualFold(pr.Name, pn):
		return types.PieceCore{}, fmt.Errorf("receipt is for piece %s, not %s", pr.Name, pn)
	case pr.Streamer != stream:
		return types.PieceCore{}, fmt.Errorf("piece %s is streamed by %s, receipt came from %s", pn, pr.Streamer, stream)
	case pr.Policy != policy:
		return types.PieceCore{}, fmt.Errorf("piece %s has policy %d/%d, the drain uses %d/%d", pn, pr.Policy.N, pr.Policy.K, policy.N, policy.K)
	case volSize <= 0 || pr.Size != volSize:
		return types.PieceCore{}, fmt.Errorf("piece %s has size %d, the volume is %d bytes", pn, pr.Size, volSize)
	}
	return types.PieceCore{Policy: policy, Name: pr.Name, Size: volSize, Streamer: stream}, nil
}

// trustedStreams lists the streams the hub takes staged-piece receipts from:
// active on chain and on the hub's chain (anyone can register an edge).
func (s *Server) trustedStreams(au types.Auth) ([]types.EdgeReceipt, error) {
	er, err := sdk.ListEdge(sdk.ServerURL, au, types.StreamType)
	if err != nil {
		return nil, err
	}
	chain := s.rp.Repo().Config().Chain.Type
	out := make([]types.EdgeReceipt, 0, len(er.Edges))
	for _, st := range er.Edges {
		if st.OnChain && st.ChainType == chain {
			out = append(out, st)
		}
	}
	return out, nil
}

func (s *Server) recoverStagedPiece(cm *contract.ContractManage, au types.Auth, key string, i uint64, pn string, policy types.Policy, volSize int64) bool {
	streams, err := s.trustedStreams(au)
	if err != nil {
		return false
	}
	for _, st := range streams {
		pr, err := sdk.GetPieceReceipt(st.ExposeURL, au, pn)
		if err != nil || pr.Name == "" {
			continue
		}
		if pr.Serial == 0 {
			pc, err := stagedPieceCore(pr, st.Name, pn, policy, volSize)
			if err != nil {
				logger.Warnf("recover staged piece %s: %v", pn, err)
				continue
			}
			txn, err := cm.AddPiece(pc)
			if err != nil {
				logger.Warnf("recover staged piece %s: AddPiece failed: %v", pn, err)
				return false
			}
			s.addVolume(key, i, pr.Name, txn)
			logger.Infof("recovered staged piece %s -> on-chain (tx %s)", pn, txn)
		} else {
			s.addVolume(key, i, pr.Name, "")
			logger.Infof("recovered staged piece %s (already on-chain, serial %d)", pn, pr.Serial)
		}
		return true
	}
	return false
}

// drainInstance uploads + commits one owner's (log instance idx) pending
// volumes in order, advancing that owner's offset as each volume lands. Encode
// (CPU) and AddPiece (chain) of different owners overlap because drainInstance
// runs concurrently per owner.
func (s *Server) drainInstance(cm *contract.ContractManage, au types.Auth, policy types.Policy, idx uint32) {
	dsKey := types.NewKey(types.DsLogFS, LOGINST, idx)
	val, err := s.rp.MetaStore().Get(dsKey)
	if err != nil {
		return
	}

	key := string(val)

	// P2 time-flush: commit a partial (open) volume once it ages past the
	// threshold, so slow/low-volume owners' small writes reach DA within bounded
	// latency instead of waiting for the 31MiB size trigger. Still batched —
	// everything accumulated in the window goes into one piece. Roll() advances
	// the persisted offset, so the freshly-completed volume is uploaded below.
	if maxAge := time.Duration(env.Int("HUB_FLUSH_MAX_AGE_SEC", 300)) * time.Second; maxAge > 0 {
		if fs := s.getLFS(key); fs != nil {
			if sz, age := fs.Pending(); sz > 0 && age >= maxAge {
				if err := fs.Roll(); err != nil {
					logger.Warnf("time-flush %s failed: %v", key, err)
				} else {
					logger.Infof("time-flush %s: rolled %d bytes (age %s)", key, sz, age.Truncate(time.Second))
				}
			}
		}
	}

	logger.Debugf("check: %s %d", key, idx)
	dsKey = types.NewKey(types.DsLogFS, key)
	val, err = s.rp.MetaStore().Get(dsKey)
	if err != nil || len(val) != 8 {
		return
	}
	curIndex := binary.BigEndian.Uint64(val)

	next := drainNext(s.rp.MetaStore(), s.local.String(), key)
	dsKey = drainOffsetKey(key)

	logger.Debugf("check: %s %d %d", key, next, curIndex)
	if next >= curIndex {
		return
	}

	for i := next; i < curIndex; i++ {
		fname := fmt.Sprintf("%s/%d.vol", key, i)
		fp := filepath.Join(s.rp.Path(), LOGFS, key, fmt.Sprintf("%d.vol", i))

		// only the hub's own record counts: anyone can register a file under
		// this (predictable) volume name, and trusting it would skip the volume
		vi, serr := os.Stat(fp)
		if serr != nil {
			// volumes the drain has not committed are kept on disk (the
			// reclaim gate); one missing here predates that gate or was
			// removed by hand, and nothing below can work without it
			logger.Errorf("drain %s vol %d: local volume unavailable (%v); restore %s to continue", key, i, serr, fp)
			break
		}
		volSize := vi.Size()
		fr, err := sdk.GetFileReceiptOf(sdk.ServerURL, fname, au.Addr)
		if err == nil && !volumeMatches(fp, fr.Hash) {
			// a record of ours under this name that is not this volume's bytes
			// (e.g. planted before seal names were namespaced): don't trust it
			logger.Warnf("%s/%d.vol: file record %s does not match the local volume; uploading", key, i, fr.Hash)
			err = fmt.Errorf("record does not match volume")
		}
		if err == nil {
			logger.Infof("%s/%d.vol is already uploaded, check its piece onchain", key, i)
			if fr.ChainType != s.rp.Repo().Config().Chain.Type {
				buf := make([]byte, 8)
				binary.BigEndian.PutUint64(buf, i+1)
				s.rp.MetaStore().Put(dsKey, buf)
				logger.Warnf("new chain type detected, ignore previous one")
				continue
			}
			streams, err := s.trustedStreams(au)
			if err != nil {
				break
			}
			suc := 0
			for _, pn := range fr.Pieces {
				// one registration per piece, from the first stream that holds it
				for _, st := range streams {
					pr, err := sdk.GetPieceReceipt(st.ExposeURL, au, pn)
					if err != nil {
						continue
					}
					if pr.Serial > 0 {
						suc++
						break
					}
					pc, err := stagedPieceCore(pr, st.Name, pn, policy, volSize)
					if err != nil {
						logger.Warnf("%s/%d.vol: %v", key, i, err)
						continue
					}
					txn, err := cm.AddPiece(pc)
					if err == nil {
						s.addVolume(key, i, pc.Name, txn)
						suc++
						break
					}
				}
			}
			if suc == len(fr.Pieces) {
				buf := make([]byte, 8)
				binary.BigEndian.PutUint64(buf, i+1)
				s.rp.MetaStore().Put(dsKey, buf)
				continue
			}
			continue
		}
		// upload to stream and submit to gateway
		// each request signed when sent: encoding a volume takes minutes, and
		// the gateway and streams reject signatures older than ~10 minutes
		res, streamer, err := sdk.UploadWith(sdk.ServerURL, s.rp.Key().BuildAuth, policy, fp, fname)
		if err != nil {
			// The piece is already staged on a stream from a prior attempt whose
			// on-chain AddPiece never completed (hub briefly out of gas, or a
			// concurrent drain pass double-uploaded this slow-encoding vol). Recover:
			// register that staged piece on-chain if it isn't yet, record the volume,
			// and advance. The old code just skipped (piece) or broke (replica) → the
			// vol stayed staged-but-uncommitted forever. "already has replica" carries
			// a replica name, so resolve it to its piece first.
			pn := parseAlreadyHasPiece(err.Error())
			if pn == "" {
				if rn := parseAlreadyHasReplica(err.Error()); rn != "" {
					pn = s.pieceOfReplica(au, rn)
				}
			}
			if pn != "" && s.recoverStagedPiece(cm, au, key, i, pn, policy, volSize) {
				buf := make([]byte, 8)
				binary.BigEndian.PutUint64(buf, i+1)
				s.rp.MetaStore().Put(dsKey, buf)
				continue
			}
			logger.Warnf("drain %s vol %d: %v", key, i, err)
			break
		}
		pcs, err := sdk.CheckFileFull(res, streamer, fp)
		if err != nil {
			break
		}
		log.Printf("upload %s to %s, sha256: %s\n", fp, streamer, res.Hash)
		log.Printf("submit %s to chain\n", res.Name)
		// submit meta to chain
		var terr error
		for _, pc := range pcs {
			txn, err := cm.AddPiece(pc)
			if err != nil {
				terr = err
				break
			}
			s.addVolume(key, i, pc.Name, txn)
		}
		if terr != nil {
			break
		}
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, i+1)
		s.rp.MetaStore().Put(dsKey, buf)
	}
}

// drainOffsetKey is where drainInstance keeps an owner's next volume to
// commit to DA: every volume below it is committed.
func drainOffsetKey(owner string) []byte {
	return types.NewKey(types.DsLogFS, LOGINST, owner)
}

// drainNext reads an owner's drain offset (its first volume index when the
// drain has not committed any).
func drainNext(ds types.IKVStore, local, owner string) uint64 {
	if val, err := ds.Get(drainOffsetKey(owner)); err == nil && len(val) == 8 {
		return binary.BigEndian.Uint64(val)
	}
	return logfs.GetIndex(local, owner)
}

// volumeMatches reports whether the sealed volume file at fp has sha256 hash
// (hex), i.e. a file record with that hash really is this volume.
func volumeMatches(fp, hash string) bool {
	f, err := os.Open(fp)
	if err != nil {
		return false
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return false
	}
	return strings.EqualFold(hex.EncodeToString(h.Sum(nil)), hash)
}
