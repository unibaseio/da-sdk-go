package hub

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	contract "github.com/unibaseio/da-sdk-go/contract/v2"
	"github.com/unibaseio/da-sdk-go/lib/env"
	lerror "github.com/unibaseio/da-sdk-go/lib/error"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
)

func (s *Server) addUpload(g *gin.RouterGroup) {
	g.Group("/").POST("/uploadData", s.uploadData)
	g.Group("/").POST("/upload", s.upload)
	g.Group("/").POST("/uploadDir", s.uploadDir) // large-file folder (model/dataset)
}

// addUploadReadonly registers the write paths on a read-only replica so they
// return 503 instead of writing to the replica's local logfs/badger (which
// would fork the data). Writes must be routed to the primary by the LB.
func (s *Server) addUploadReadonly(g *gin.RouterGroup) {
	reject := func(c *gin.Context) {
		c.JSON(http.StatusServiceUnavailable,
			lerror.ToAPIError("hub", fmt.Errorf("this node is read-only; route writes to the primary")))
	}
	g.Group("/").POST("/uploadData", reject)
	g.Group("/").POST("/upload", reject)
	g.Group("/").POST("/uploadDir", reject)
}

// uploadDir ingests a folder of large files (model/dataset scenario) in ONE
// multipart request: each file becomes its own DA-backed object under bucket.
// kind is "model" or "dataset"; bucket = the repo name. Unlike memory's
// small-write coalescing, each file is written with the passthrough-large
// policy (its own volume) so it uploads to DA promptly and cleanly.
func (s *Server) uploadDir(c *gin.Context) {
	addr := c.PostForm("owner")
	if !RequireOwnerMatch(c, addr) {
		return
	}
	kind := c.PostForm("kind")
	if kind != "model" && kind != "dataset" {
		c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("kind must be 'model' or 'dataset'")))
		return
	}
	bucket := c.PostForm("bucket")
	if bucket == "" {
		c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("bucket (repo name) required")))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("no files (use form field 'files')")))
		return
	}

	metas := make([]types.MemeMeta, 0, len(files))
	for _, fh := range files {
		if fh.Size == 0 {
			c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("empty file: %s", fh.Filename)))
			return
		}
		fr, err := fh.Open()
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		mm, err := s.logFSWriteEx(addr, bucket, fh.Filename, kind, true, fr)
		fr.Close()
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		metas = append(metas, mm)
	}

	c.JSON(http.StatusOK, metas)
}

func (s *Server) uploadData(c *gin.Context) {
	addr := c.PostForm("owner")
	if !RequireOwnerMatch(c, addr) {
		return
	}
	bucket := c.PostForm("bucket")
	if bucket == "" {
		bucket = addr
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	if file == nil {
		c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("file is nil")))
		return
	}

	fr, err := file.Open()
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	if file.Size == 0 {
		c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("empty file")))
		return
	}
	// kind: memory (default) or knowledgebase — small-object/coalesce path either way.
	mm, err := s.logFSWriteEx(addr, bucket, file.Filename, memKind(c.PostForm("kind")), false, fr)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, mm)
}

func (s *Server) upload(c *gin.Context) {
	var mjson types.MemeStruct

	err := c.ShouldBindJSON(&mjson)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	if !RequireOwnerMatch(c, mjson.Owner) {
		return
	}

	if mjson.Bucket == "" {
		var meta map[string]interface{}
		err = json.Unmarshal([]byte(mjson.Message), &meta)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		bucketName, ok := meta["name"].(string)
		if ok {
			mjson.Bucket = bucketName
		} else {
			mjson.Bucket = mjson.Owner
		}
	}

	var buf bytes.Buffer
	buf.WriteString(mjson.Message)

	mm, err := s.logFSWriteEx(mjson.Owner, mjson.Bucket, mjson.ID, memKind(mjson.Kind), false, &buf)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, mm)
}

// memKind normalizes the kind for the small-object (coalesce) upload path:
// only "memory" (default) and "knowledgebase" (RAG chunk) belong here; empty or
// a large-file kind (model/dataset — those use /uploadDir) falls back to memory.
func memKind(k string) string {
	if k == "knowledgebase" {
		return k
	}
	return "memory"
}

// logFSWriteEx writes one object to the owner's logfs + indexes it.
//   - kind:  bucket scenario ("memory"/"model"/"dataset").
//   - large: passthrough-large policy — seal any pending small-write volume
//     first, then seal this object into its own volume (isolated, uploads
//     promptly). large=false keeps the memory coalescing behavior.
func (s *Server) logFSWriteEx(addr string, bucket string, key string, kind string, large bool, r io.Reader) (types.MemeMeta, error) {
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
	rbytes, err := io.ReadAll(r)
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

	// Drop any stale "missing" marker so this key is immediately downloadable.
	s.missCache.del(addr, key)

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

func (s *Server) logFSRead(addr string, key string, w io.Writer) (int64, error) {
	if addr == "" {
		ns, err := s.getNeedleByName(key)
		if err != nil || len(ns) == 0 {
			return 0, fmt.Errorf("no such needle %s", key)
		}
		// Found the exact owner the needle was stored under — read it directly.
		return s.logFSReadOne(ns[0].Owner, key, w)
	}

	// Client-supplied owner: try the canonical lowercase form first, then the
	// EIP-55 checksum form (legacy data was stored under the mixed-case
	// address). logFSReadOne only writes to w on success, so failed candidates
	// leave w untouched.
	var firstErr error
	for _, cand := range ownerCandidates(addr) {
		n, err := s.logFSReadOne(cand, key, w)
		if err == nil {
			return n, nil
		}
		if firstErr == nil {
			firstErr = err
		}
	}
	return 0, firstErr
}

func (s *Server) logFSReadOne(addr string, key string, w io.Writer) (int64, error) {
	fs, err := s.getFS(addr, false)
	if err != nil {
		return 0, err
	}

	lm, err := fs.GetMeta([]byte(key))
	if err != nil {
		return 0, err
	}

	wbytes, err := fs.GetData(lm)
	if err != nil {
		// todo: get from remote
		return 0, err
	}
	n, err := w.Write(wbytes)
	if err != nil {
		return 0, err
	}

	return int64(n), nil
}

func (s *Server) load() error {
	fspath := filepath.Join(s.rp.Path(), LOGFS)
	fs, err := logfs.New(s.rp.MetaStore(), fspath, s.local.String(), s.local.String())
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
func (s *Server) recoverStagedPiece(cm *contract.ContractManage, au types.Auth, key string, i uint64, pn string) bool {
	er, err := sdk.ListEdge(sdk.ServerURL, au, types.StreamType)
	if err != nil {
		return false
	}
	for _, st := range er.Edges {
		pr, err := sdk.GetPieceReceipt(st.ExposeURL, au, pn)
		if err != nil || pr.Name == "" {
			continue
		}
		if pr.Serial == 0 {
			txn, err := cm.AddPiece(pr.PieceCore)
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

	next := logfs.GetIndex(s.local.String(), key)
	dsKey = types.NewKey(types.DsLogFS, LOGINST, key)
	val, err = s.rp.MetaStore().Get(dsKey)
	if err == nil && len(val) == 8 {
		next = binary.BigEndian.Uint64(val)
	}

	logger.Debugf("check: %s %d %d", key, next, curIndex)
	if next >= curIndex {
		return
	}

	for i := next; i < curIndex; i++ {
		fname := fmt.Sprintf("%s/%d.vol", key, i)
		fp := filepath.Join(s.rp.Path(), LOGFS, key, fmt.Sprintf("%d.vol", i))

		fr, err := sdk.GetFileReceipt(sdk.ServerURL, au, fname)
		if err == nil {
			logger.Infof("%s/%d.vol is already uploaded, check its piece onchain", key, i)
			if fr.ChainType != s.rp.Repo().Config().Chain.Type {
				buf := make([]byte, 8)
				binary.BigEndian.PutUint64(buf, i+1)
				s.rp.MetaStore().Put(dsKey, buf)
				logger.Warnf("new chain type detected, ignore previous one")
				continue
			}
			er, err := sdk.ListEdge(sdk.ServerURL, au, types.StreamType)
			if err != nil {
				break
			}
			suc := 0
			for _, pn := range fr.Pieces {
				for _, st := range er.Edges {
					pr, err := sdk.GetPieceReceipt(st.ExposeURL, au, pn)
					if err == nil {
						if pr.Serial > 0 {
							suc++
						} else {
							txn, err := cm.AddPiece(pr.PieceCore)
							if err == nil {
								s.addVolume(key, i, pr.Name, txn)
								suc++
							}
						}
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
		res, streamer, err := sdk.Upload(sdk.ServerURL, au, policy, fp, fname)
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
			if pn != "" && s.recoverStagedPiece(cm, au, key, i, pn) {
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
