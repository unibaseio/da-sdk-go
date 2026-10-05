package hub

// /v1 — resource-oriented, wallet-native object-store API (spec: da/HUB_API_V1_SPEC.md).
// S3 SHAPE (bucket/object/key) but web3/DA-native semantics: kind on the bucket,
// verifiable receipt (commitment + chain + status), async staged→committed.
//
// A thin façade over the Server's index (gdb) and per-owner LogFS; it is the
// hub's only HTTP surface.

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/unibaseio/da-sdk-go/lib/env"
	lerror "github.com/unibaseio/da-sdk-go/lib/error"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// v1WantTotal reports whether the caller asked for a grand total row count
// (?withTotal=1). Opt-in because COUNT(*) can be expensive on large needle
// tables — normal cursor/offset paging stays cheap; only clients that render a
// numbered pager (e.g. the block explorer) pay for it.
func v1WantTotal(c *gin.Context) bool {
	v := c.Query("withTotal")
	return v == "1" || v == "true"
}

// v1Offset reads an optional ?offset= (numbered pagination). Returns (n, true)
// when present and valid; the caller then uses OFFSET instead of cursor.
func v1Offset(c *gin.Context) (int, bool, error) {
	off := c.Query("offset")
	if off == "" {
		return 0, false, nil
	}
	n, err := strconv.Atoi(off)
	if err != nil || n < 0 {
		return 0, false, fmt.Errorf("invalid offset")
	}
	return n, true, nil
}

// v1KindProfileLarge maps a bucket kind (+ object size) to the storage profile:
// passthrough-large (own volume, prompt upload) vs coalesce-small (batch into a
// shared volume). memory/knowledgebase coalesce; model/dataset passthrough;
// file adapts by size.
const v1FileLargeThreshold = 1 << 20 // 1 MiB: a "file" bigger than this gets its own volume

func v1KindProfileLarge(kind string, size int64) bool {
	switch kind {
	case "model", "dataset":
		return true
	case "file":
		return size > v1FileLargeThreshold
	default: // memory, knowledgebase
		return false
	}
}

// v1KindValid reports whether kind is a recognized bucket scenario.
func v1KindValid(kind string) bool {
	switch kind {
	case "memory", "knowledgebase", "file", "model", "dataset":
		return true
	}
	return false
}

// v1Receipt is the verifiable object receipt (spec §4).
type v1Receipt struct {
	Bucket       string     `json:"bucket"`
	Owner        string     `json:"owner,omitempty"` // set on cross-bucket listings (GET /v1/objects)
	Key          string     `json:"key"`
	Size         uint64     `json:"size"`
	Sha256       string     `json:"sha256,omitempty"`     // content sha256 (best-effort, from logfs meta)
	Keccak256    string     `json:"keccak256,omitempty"`  // content keccak256, 0x-hex (the hash ERC-8004 fields carry)
	Commitment   string     `json:"commitment,omitempty"` // DA piece (content-addressed id)
	Status       string     `json:"status"`               // staged → committed
	Chain        *v1Chain   `json:"chain,omitempty"`
	Availability string     `json:"availability"` // verified once committed
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
}

type v1Chain struct {
	TxHash    string `json:"txHash"`
	ChainType string `json:"chainType"`
}

// receiptFromNeedle builds a receipt from a NeedleDisplay (which already joined
// the Volume for piece/txhash). status=committed once the volume's piece is on
// chain (Volume row exists with a piece), else staged (in logfs, awaiting the
// drainInstance AddPiece).
func receiptFromNeedle(bucket string, n types.NeedleDisplay) v1Receipt {
	r := v1Receipt{Bucket: bucket, Owner: n.Owner, Key: n.Name, Size: n.Size, Status: "staged", Availability: "pending"}
	if !n.CreatedAt.IsZero() {
		t := n.CreatedAt
		r.CreatedAt = &t
	}
	if n.Piece != "" {
		r.Commitment = n.Piece
		r.Status = "committed"
		r.Availability = "verified"
		r.Chain = &v1Chain{TxHash: n.TxHash, ChainType: n.ChainType}
	}
	return r
}

// maxBodyV1 caps request bodies on the v1 write group. Object writes stream a
// (possibly large) blob, so use the multipart cap for the whole group.
func maxBodyV1() gin.HandlerFunc {
	capBytes := env.Int64("HUB_MAX_MULTIPART_BYTES", defaultMaxMultipartBytes)
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, capBytes)
		c.Next()
	}
}

// registV1 mounts the /v1 surface: public reads (no auth) + authed writes.
func (s *Server) registV1() {
	pub := s.Router.Group("/v1")
	// public reads carry no signer at middleware time: per-IP tier only
	pub.Use(IPRateLimit())
	pub.GET("/info", s.v1Info)
	pub.GET("/buckets", s.v1ListBuckets)
	pub.GET("/buckets/:bucket", s.v1GetBucket)
	pub.GET("/objects", s.v1ListAllObjects)
	pub.GET("/buckets/:bucket/objects", s.v1ListObjects)
	pub.GET("/buckets/:bucket/objects/:key", s.v1GetObject)
	pub.GET("/buckets/:bucket/objects/:key/content", s.v1GetObjectContent)
	pub.GET("/buckets/:bucket/objects/:key/proof", s.v1GetObjectProof)
	pub.GET("/pieces/:name/content", s.v1GetPieceContent)
	pub.GET("/conversations", s.v1ListConversations)
	pub.GET("/conversations/:id", s.v1GetConversation)
	pub.GET("/stats", s.v1Stats)
	pub.GET("/cachestats", s.v1CacheStats)
	pub.GET("/overview", s.v1Overview)
	pub.GET("/owners/:owner", s.v1GetOwner)

	w := s.Router.Group("/v1")
	w.Use(maxBodyV1())
	// per-IP limit BEFORE signature verification (the expensive step), per-owner
	// limit after it (needs the recovered signer)
	w.Use(IPRateLimit())
	w.Use(AuthMiddleware())
	// P4-Route: after auth (owner = signer), forward a non-home owner's write to
	// its shard. No-op unless HUB_SHARD_TOTAL>1. Reads are never sharded.
	w.Use(s.shardWrite())
	w.Use(OwnerRateLimit())
	if s.readonly {
		reject := func(c *gin.Context) {
			c.JSON(http.StatusServiceUnavailable, lerror.ToAPIError("hub", fmt.Errorf("this node is read-only; route writes to the primary")))
		}
		w.PUT("/buckets/:bucket", reject)
		w.PUT("/buckets/:bucket/objects/:key", reject)
		w.POST("/buckets/:bucket/objects", reject)
		w.POST("/seal", reject)
	} else {
		w.PUT("/buckets/:bucket", s.v1PutBucket)
		w.DELETE("/buckets/:bucket", s.v1DeleteBucket)
		w.PUT("/buckets/:bucket/objects/:key", s.v1PutObject)
		w.POST("/buckets/:bucket/objects", s.v1PostObjects)
		// seal: commit one already-encrypted blob as a DA piece on-chain (register
		// modes hub / hub_attributed / client). An operation, not a resource verb —
		// the same /v1-operations carve-out as the nodes' upload/download. Reuses the
		// existing seal handler (form + DA_SEAL_ENDPOINT_SPEC contract preserved).
		w.POST("/seal", s.seal)
	}
}

// ---- Info ------------------------------------------------------------------

// GET /v1/info — hub node identity.
func (s *Server) v1Info(c *gin.Context) {
	c.JSON(http.StatusOK, types.EdgeReceipt{EdgeMeta: types.EdgeMeta{Type: s.typ, Name: s.local}})
}

// ---- Conversations ---------------------------------------------------------
// A conversation is a memory-domain view grouping an owner's objects. Read-only;
// enumeration requires owner==signer (RequireOwnerForList), point-get resolves the
// owner (ResolveOwnerForList) — same authz as objects.

// GET /v1/conversations?owner=&bucket=&offset=&limit= — list an owner's conversation
// ids (the raw prefixes).
func (s *Server) v1ListConversations(c *gin.Context) {
	owner, ok := RequireOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	offset, length, ok := v1Page(c, 1024)
	if !ok {
		return
	}
	res, err := s.listConversation(owner, c.Query("bucket"), offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"conversations": res})
}

// GET /v1/conversations/{id}?owner=&bucket= — the raw records under {id} in write
// order (each record = a stored payload).
func (s *Server) v1GetConversation(c *gin.Context) {
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	// a conversation is one owner's: without an owner the lookup would span
	// every account's conversations of that id
	if owner == "" {
		c.JSON(http.StatusBadRequest, lerror.ToAPIError("hub", fmt.Errorf("owner required")))
		return
	}
	offset, length, ok := v1Page(c, 1024)
	if !ok {
		return
	}
	res, err := s.getConversation(c.Request.Context(), c.Param("id"), owner, c.Query("bucket"), offset, length)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"conversation": c.Param("id"), "messages": res})
}

// ---- bucket lookup helper (direct gdb; kind normalized) --------------------

func (s *Server) v1LookupBucket(bucket string) (types.Bucket, bool) {
	var b types.Bucket
	res := s.gdb.Where("name = ?", bucket).First(&b)
	if res.RowsAffected == 0 {
		return b, false
	}
	if b.Kind == "" {
		b.Kind = "memory"
	}
	return b, true
}

// ---- Buckets ---------------------------------------------------------------

type v1PutBucketBody struct {
	Kind string `json:"kind"`
}

// PUT /v1/buckets/{bucket} — ensure a bucket with a kind (idempotent; kind immutable).
func (s *Server) v1PutBucket(c *gin.Context) {
	owner := CtxAuthAddr(c)
	bucket := c.Param("bucket")
	var body v1PutBucketBody
	_ = c.ShouldBindJSON(&body)
	if body.Kind == "" {
		body.Kind = "memory"
	}
	if !v1KindValid(body.Kind) {
		abortWithBadRequest(c, fmt.Errorf("invalid kind %q (want memory|knowledgebase|file|model|dataset)", body.Kind))
		return
	}
	if existing, ok := s.v1LookupBucket(bucket); ok {
		if !strings.EqualFold(existing.Owner, owner) {
			c.JSON(http.StatusForbidden, lerror.ToAPIError("hub", fmt.Errorf("bucket %s owned by another account", bucket)))
			return
		}
		if existing.Kind != body.Kind {
			c.JSON(http.StatusConflict, lerror.ToAPIError("hub", fmt.Errorf("bucket kind is immutable (%s); cannot change to %s", existing.Kind, body.Kind)))
			return
		}
		c.JSON(http.StatusOK, gin.H{"bucket": bucket, "owner": owner, "kind": body.Kind, "created": false})
		return
	}
	if err := s.addBucket(CanonOwner(owner), bucket, body.Kind); err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"bucket": bucket, "owner": owner, "kind": body.Kind, "created": true})
}

// GET /v1/buckets?owner=&kind=&cursor=&limit=  (cursor on bucket id)
func (s *Server) v1ListBuckets(c *gin.Context) {
	owner, ok := RequireOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	limit := v1Limit(c)
	kind := c.Query("kind")

	filter := func(q *gorm.DB) *gorm.DB {
		if owner != "" {
			q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
		}
		if kind == "memory" {
			q = q.Where("kind = ? OR kind = '' OR kind IS NULL", kind)
		} else if kind != "" {
			q = q.Where("kind = ?", kind)
		}
		return q
	}

	resp := gin.H{}
	if v1WantTotal(c) {
		total, err := s.cachedCount("buckets|"+strings.ToLower(owner)+"|"+kind, &types.Bucket{}, filter)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		resp["total"] = total
	}

	q := filter(s.gdb.Model(&types.Bucket{}))
	off, useOffset, err := v1Offset(c)
	if err != nil {
		abortWithBadRequest(c, err)
		return
	}
	if useOffset {
		q = q.Offset(off)
	} else if cur := c.Query("cursor"); cur != "" {
		id, err := strconv.ParseUint(cur, 10, 64)
		if err != nil {
			abortWithBadRequest(c, fmt.Errorf("invalid cursor"))
			return
		}
		q = q.Where("id < ?", id)
	}
	var buckets []types.Bucket
	if err := q.Order("id desc").Limit(limit).Find(&buckets).Error; err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	var next string
	if !useOffset && len(buckets) == limit {
		next = strconv.FormatUint(uint64(buckets[len(buckets)-1].ID), 10)
	}
	resp["buckets"] = buckets
	resp["nextCursor"] = next
	c.JSON(http.StatusOK, resp)
}

// GET /v1/buckets/{bucket}
func (s *Server) v1GetBucket(c *gin.Context) {
	b, ok := s.v1LookupBucket(c.Param("bucket"))
	if !ok {
		c.JSON(http.StatusNotFound, lerror.ToAPIError("hub", fmt.Errorf("no such bucket: %s", c.Param("bucket"))))
		return
	}
	c.JSON(http.StatusOK, gin.H{"bucket": b.Name, "owner": b.Owner, "kind": b.Kind, "createdAt": b.CreatedAt})
}

// ---- Objects ---------------------------------------------------------------

// PUT /v1/buckets/{bucket}/objects/{key} — put a single object (raw body).
func (s *Server) v1PutObject(c *gin.Context) {
	owner := CtxAuthAddr(c)
	bucket := c.Param("bucket")
	key := c.Param("key")

	b, ok := s.v1LookupBucket(bucket)
	if !ok {
		c.JSON(http.StatusNotFound, lerror.ToAPIError("hub", fmt.Errorf("no such bucket %s; PUT /v1/buckets/%s first", bucket, bucket)))
		return
	}
	if !strings.EqualFold(b.Owner, owner) {
		c.JSON(http.StatusForbidden, lerror.ToAPIError("hub", fmt.Errorf("bucket %s owned by another account", bucket)))
		return
	}

	// a declared length lets an over-quota write fail before its body is read
	if c.Request.ContentLength > 0 && !s.chargeWrite(c, c.Request.ContentLength) {
		return
	}
	// one copy of the body: read it (preallocating from Content-Length up to
	// maxPrealloc) and hand that slice to LogFS (it used to be copied twice more)
	data, err := readAllSized(c.Request.Body, c.Request.ContentLength)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	if len(data) == 0 {
		abortWithBadRequest(c, fmt.Errorf("empty body"))
		return
	}
	if c.Request.ContentLength <= 0 && !s.chargeWrite(c, int64(len(data))) {
		return
	}

	large := v1KindProfileLarge(b.Kind, int64(len(data)))
	if _, err := s.logFSWriteData(owner, bucket, key, b.Kind, large, data); err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	// ?wait=1: block until the object commits on-chain (drainInstance AddPiece),
	// then return 200 + committed receipt. Best for passthrough kinds (own volume,
	// uploads promptly); small coalesce kinds may stay staged until the batch flush.
	kh := keccakHex(data)
	if c.Query("wait") == "1" {
		if nd, done := s.v1WaitCommitted(c.Request.Context(), owner, bucket, key); done {
			r := receiptFromNeedle(bucket, nd)
			r.Keccak256 = kh
			c.JSON(http.StatusOK, r)
			return
		}
	}
	// staged: on logfs + indexed; committed later by drainInstance's AddPiece.
	c.JSON(http.StatusAccepted, v1Receipt{Bucket: bucket, Key: key, Size: uint64(len(data)), Keccak256: kh, Status: "staged", Availability: "pending"})
}

// POST /v1/buckets/{bucket}/objects — batch put (multipart files[]).
func (s *Server) v1PostObjects(c *gin.Context) {
	owner := CtxAuthAddr(c)
	bucket := c.Param("bucket")

	b, ok := s.v1LookupBucket(bucket)
	if !ok {
		c.JSON(http.StatusNotFound, lerror.ToAPIError("hub", fmt.Errorf("no such bucket %s; PUT /v1/buckets/%s first", bucket, bucket)))
		return
	}
	if !strings.EqualFold(b.Owner, owner) {
		c.JSON(http.StatusForbidden, lerror.ToAPIError("hub", fmt.Errorf("bucket %s owned by another account", bucket)))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		abortWithBadRequest(c, fmt.Errorf("no files (use form field 'files')"))
		return
	}
	var total int64
	for _, fh := range files {
		if fh.Size == 0 {
			abortWithBadRequest(c, fmt.Errorf("empty file: %s", fh.Filename))
			return
		}
		total += fh.Size
	}
	// the whole batch counts against the signer's quota, all or nothing
	if !s.chargeWrite(c, total) {
		return
	}
	receipts := make([]v1Receipt, 0, len(files))
	for _, fh := range files {
		fr, err := fh.Open()
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		large := v1KindProfileLarge(b.Kind, fh.Size)
		_, err = s.logFSWriteEx(owner, bucket, fh.Filename, b.Kind, large, fr)
		fr.Close()
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		receipts = append(receipts, v1Receipt{Bucket: bucket, Key: fh.Filename, Size: uint64(fh.Size), Status: "staged", Availability: "pending"})
	}
	c.JSON(http.StatusAccepted, gin.H{"objects": receipts})
}

// GET /v1/buckets/{bucket}/objects?owner=&cursor=&limit=
func (s *Server) v1ListObjects(c *gin.Context) {
	owner, ok := RequireOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Param("bucket")
	limit := v1Limit(c)

	// cursor pagination on the needles table by id (avoids offset walk on 34M rows):
	// WHERE bucket=? [AND LOWER(owner)=?] [AND id < cursor] ORDER BY id DESC LIMIT n.
	// Offset + total are opt-in (numbered pager) — see v1Offset/v1WantTotal.
	filter := func(q *gorm.DB) *gorm.DB {
		q = q.Where("bucket = ?", bucket)
		if owner != "" {
			q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
		}
		return q
	}

	resp := gin.H{}
	if v1WantTotal(c) {
		total, err := s.cachedCount("needles|"+strings.ToLower(owner)+"|"+bucket, &types.Needle{}, filter)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		resp["total"] = total
	}

	q := filter(s.gdb.Model(&types.Needle{}))
	off, useOffset, err := v1Offset(c)
	if err != nil {
		abortWithBadRequest(c, err)
		return
	}
	if useOffset {
		q = q.Offset(off)
	} else if cur := c.Query("cursor"); cur != "" {
		id, err := strconv.ParseUint(cur, 10, 64)
		if err != nil {
			abortWithBadRequest(c, fmt.Errorf("invalid cursor"))
			return
		}
		q = q.Where("id < ?", id)
	}
	var needles []types.Needle
	if err := q.Order("id desc").Limit(limit).Find(&needles).Error; err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	vmap := s.volumesFor(needles) // batch volume join (piece/txhash), avoids N+1
	objs := make([]v1Receipt, 0, len(needles))
	for i := range needles {
		nd := types.NeedleDisplay{CreatedAt: needles[i].CreatedAt, Name: needles[i].Name, Size: needles[i].Size}
		if v, ok := vmap[volKey(needles[i].Owner, needles[i].File)]; ok {
			nd.Piece = v.Piece
			nd.TxHash = v.TxHash
			nd.ChainType = v.ChainType
		}
		objs = append(objs, receiptFromNeedle(bucket, nd))
	}
	var next string
	if !useOffset && len(needles) == limit {
		next = strconv.FormatUint(uint64(needles[len(needles)-1].ID), 10)
	}
	resp["objects"] = objs
	resp["nextCursor"] = next
	c.JSON(http.StatusOK, resp)
}

// GET /v1/objects?owner=&cursor=&limit= — cross-bucket needle list. Global
// enumeration (no bucket path param), so it requires an owner-signed request
// like the other list endpoints: the owner scopes to their own needles, and a
// reader identity (HUB_READER_ADDRS) may enumerate any/all owners. Backs the
// explorer's top-level Memory page, which shows recent memory across all agents.
func (s *Server) v1ListAllObjects(c *gin.Context) {
	owner, ok := RequireOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	limit := v1Limit(c)

	filter := func(q *gorm.DB) *gorm.DB {
		if owner != "" {
			q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
		}
		return q
	}

	resp := gin.H{}
	if v1WantTotal(c) {
		total, err := s.cachedCount("needles|"+strings.ToLower(owner)+"|", &types.Needle{}, filter)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		resp["total"] = total
	}

	q := filter(s.gdb.Model(&types.Needle{}))
	off, useOffset, err := v1Offset(c)
	if err != nil {
		abortWithBadRequest(c, err)
		return
	}
	if useOffset {
		q = q.Offset(off)
	} else if cur := c.Query("cursor"); cur != "" {
		id, err := strconv.ParseUint(cur, 10, 64)
		if err != nil {
			abortWithBadRequest(c, fmt.Errorf("invalid cursor"))
			return
		}
		q = q.Where("id < ?", id)
	}
	var needles []types.Needle
	if err := q.Order("id desc").Limit(limit).Find(&needles).Error; err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	vmap := s.volumesFor(needles)
	objs := make([]v1Receipt, 0, len(needles))
	for i := range needles {
		nd := types.NeedleDisplay{
			CreatedAt: needles[i].CreatedAt,
			Name:      needles[i].Name,
			Owner:     needles[i].Owner,
			Bucket:    needles[i].Bucket,
			Size:      needles[i].Size,
		}
		if v, ok := vmap[volKey(needles[i].Owner, needles[i].File)]; ok {
			nd.Piece = v.Piece
			nd.TxHash = v.TxHash
			nd.ChainType = v.ChainType
		}
		objs = append(objs, receiptFromNeedle(needles[i].Bucket, nd))
	}
	var next string
	if !useOffset && len(needles) == limit {
		next = strconv.FormatUint(uint64(needles[len(needles)-1].ID), 10)
	}
	resp["objects"] = objs
	resp["nextCursor"] = next
	c.JSON(http.StatusOK, resp)
}

// GET /v1/buckets/{bucket}/objects/{key}
func (s *Server) v1GetObject(c *gin.Context) {
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Param("bucket")
	key := c.Param("key")
	nds, err := s.getNeedleDisplay(owner, bucket, key)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	if len(nds) == 0 {
		c.JSON(http.StatusNotFound, lerror.ToAPIError("hub", fmt.Errorf("no such object %s/%s", bucket, key)))
		return
	}
	rc := receiptFromNeedle(bucket, nds[0])
	rc.Sha256 = s.objectSha256At(nds[0]) // best-effort, this row's object only
	c.JSON(http.StatusOK, rc)
}

// GET /v1/buckets/{bucket}/objects/{key}/content — download bytes.
// The object is the newest index row for (owner, bucket, key) — the same row
// its receipt and /proof describe — and its bytes are read from the location
// that row records, never looked up by key alone.
func (s *Server) v1GetObjectContent(c *gin.Context) {
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Param("bucket")
	key := c.Param("key")
	nds, err := s.getNeedleDisplay(owner, bucket, key)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	if len(nds) == 0 {
		c.JSON(http.StatusNotFound, lerror.ToAPIError("hub", fmt.Errorf("no such object %s/%s", bucket, key)))
		return
	}
	n := nds[0]
	var w bytes.Buffer
	if _, err := s.logFSReadAt(n.Owner, key, n.File, n.Start, n.Size, &w); err != nil {
		// P4 shard fallback: a STAGED object's bytes live only on its owner's
		// home shard (index rows are shared, logfs blobs are not, and L2 fills
		// on read, not write). If we're not the home shard, forward the read
		// there instead of 404ing the read-after-write window.
		if s.shard != nil && s.shardReadProxy(c, n.Owner) {
			return
		}
		c.JSON(http.StatusNotFound, lerror.ToAPIError("hub", err))
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", w.Bytes())
}

// GET /v1/pieces/{name}/content — download a committed DA piece by its
// content-id (da_cid). Unlike object content (logfs by key), this resolves the
// piece off the DA network (GetPieceReceipt → DownloadPiece) via
// downloadPiece — the cold-tier read path (seal'd segments) needs this.
//
// The name is a DA commitment, so it is resolved on DA only: never against
// hub objects or file records, which anyone can create under that name.
func (s *Server) v1GetPieceContent(c *gin.Context) {
	if _, ok := ResolveOwnerForList(c, c.Query("owner")); !ok {
		return
	}
	data, err := s.downloadPiece(c.Request.Context(), c.Param("name"))
	if err != nil {
		code := http.StatusNotFound
		if errors.Is(err, errBusy) {
			code = http.StatusServiceUnavailable
		}
		c.JSON(code, lerror.ToAPIError("hub", err))
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", data)
}

// ---- Aggregates ------------------------------------------------------------

func (s *Server) v1Overview(c *gin.Context) {
	c.JSON(http.StatusOK, s.memoryOverviewSnapshot())
}

func (s *Server) v1Stats(c *gin.Context) {
	days, _ := strconv.Atoi(c.Query("days"))
	if days <= 0 {
		days = 7
	}
	if days > maxStatDays {
		days = maxStatDays // the result is sized by days: unbounded, one request could exhaust memory
	}
	if s.statManager == nil {
		c.JSON(http.StatusOK, []types.Stat{})
		return
	}
	c.JSON(http.StatusOK, s.statManager.GetStats(days))
}

func (s *Server) v1GetOwner(c *gin.Context) {
	owner := c.Param("owner")
	accs, err := s.getAccount(owner)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	if len(accs) == 0 {
		c.JSON(http.StatusNotFound, lerror.ToAPIError("hub", fmt.Errorf("no such owner: %s", owner)))
		return
	}
	c.JSON(http.StatusOK, gin.H{"owner": owner, "account": accs[0]})
}

// DELETE /v1/buckets/{bucket} — soft-delete the bucket + its index rows.
// NOTE: on-chain DA pieces are immutable and persist forever; this only removes
// the object from hub listings (a local index tombstone), not from DA.
func (s *Server) v1DeleteBucket(c *gin.Context) {
	owner := CtxAuthAddr(c)
	bucket := c.Param("bucket")
	b, ok := s.v1LookupBucket(bucket)
	if !ok {
		c.JSON(http.StatusNotFound, lerror.ToAPIError("hub", fmt.Errorf("no such bucket: %s", bucket)))
		return
	}
	if !strings.EqualFold(b.Owner, owner) {
		c.JSON(http.StatusForbidden, lerror.ToAPIError("hub", fmt.Errorf("bucket %s owned by another account", bucket)))
		return
	}
	// scoped to the owner, not the bare name: legacy data can hold another
	// owner's bucket or objects under the same name
	if err := s.gdb.Where("name = ? AND LOWER(owner) = ?", bucket, strings.ToLower(owner)).Delete(&types.Bucket{}).Error; err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	if err := s.gdb.Where("bucket = ? AND LOWER(owner) = ?", bucket, strings.ToLower(owner)).Delete(&types.Needle{}).Error; err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"bucket": bucket, "deleted": true,
		"note": "removed from index; on-chain DA data is immutable and persists"})
}

// v1Range locates an object's bytes inside its DA piece. The piece's original
// data is the whole logfs volume file, and logfs writes each object raw at
// Start (padding only after it), so [Start, Start+Size) of the reconstructed
// piece is the object.
type v1Range struct {
	Volume uint64 `json:"volume"`
	Start  uint64 `json:"start"`
	Size   uint64 `json:"size"`
}

// maxProofHashBytes caps how much content /proof reads to compute keccak256.
// Evidence and memory objects are small; large model/dataset objects skip it.
const maxProofHashBytes = 16 << 20

// GET /v1/buckets/{bucket}/objects/{key}/proof — verification bundle. Returns
// the DA commitment + on-chain pointer so anyone can independently verify:
// the content hashes (sha256, and keccak256 for small objects), and — when the
// volume is a single piece — the object's byte range inside that piece, so a
// verifier can rebuild the piece from store nodes and check it without the hub.
// (Full ZK availability-proof export is a future enhancement.)
func (s *Server) v1GetObjectProof(c *gin.Context) {
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Param("bucket")
	key := c.Param("key")
	nds, err := s.getNeedleDisplay(owner, bucket, key)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	if len(nds) == 0 {
		c.JSON(http.StatusNotFound, lerror.ToAPIError("hub", fmt.Errorf("no such object %s/%s", bucket, key)))
		return
	}
	n := nds[0]
	if n.Piece == "" {
		c.JSON(http.StatusTooEarly, lerror.ToAPIError("hub", fmt.Errorf("object not yet committed on-chain (status=staged)")))
		return
	}
	// hashes of exactly this row's object (owner, bucket, key → volume/start/
	// size), like /content — never of a same-key object in another bucket
	sh, kh := s.objectHashesAt(n)
	resp := gin.H{
		"bucket":     bucket,
		"key":        key,
		"commitment": n.Piece, // DA piece (content-addressed id)
		"sha256":     sh,
		"chain":      v1Chain{TxHash: n.TxHash, ChainType: n.ChainType},
		"verify":     "hash the content and compare; confirm the piece is registered and unexpired on the Piece contract; optionally rebuild the piece from store nodes and hash range.start..start+size",
	}
	if kh != "" {
		resp["keccak256"] = kh
	}
	vols, err := s.volumesOf(n.Owner, n.File)
	if err == nil {
		if cur := onChain(vols, s.currentChainType()); len(cur) > 0 {
			vols = cur
		}
	}
	switch {
	case err != nil:
		// range left out; the rest of the bundle still stands
	case len(vols) == 1:
		resp["range"] = v1Range{Volume: n.File, Start: n.Start, Size: n.Size}
	default:
		resp["rangeNote"] = fmt.Sprintf("volume %d spans %d pieces; the byte range inside a piece is not resolvable", n.File, len(vols))
	}
	c.JSON(http.StatusOK, resp)
}

// objectHashesAt returns the content sha256 (hex) and keccak256 (0x-hex) of
// the object an index row points at, best-effort (empty on failure). Objects up
// to maxProofHashBytes are read by the row's location and hashed; larger ones
// get only the sha256 recorded in the LogFS meta, and only while that meta
// still describes this row's location. The LogFS key index is per owner, not
// per bucket, so a lookup by key alone could describe a same-key object in
// another bucket.
func (s *Server) objectHashesAt(n types.NeedleDisplay) (sha, keccak string) {
	if n.Owner == "" || n.Size == 0 {
		return "", ""
	}
	if _, err := s.readFS(n.Owner); err != nil {
		return "", ""
	}
	if n.Size > maxProofHashBytes {
		return s.objectSha256At(n), ""
	}
	var buf bytes.Buffer
	if _, err := s.logFSReadAt(n.Owner, n.Name, n.File, n.Start, n.Size, &buf); err != nil || uint64(buf.Len()) != n.Size {
		return "", ""
	}
	sum := sha256.Sum256(buf.Bytes())
	return hex.EncodeToString(sum[:]), keccakHex(buf.Bytes())
}

// readFS resolves an owner's LogFS for a best-effort read: the loaded instance,
// else (with a repo) a registered owner's on-demand load. Never creates one.
func (s *Server) readFS(owner string) (*logfs.LogFS, error) {
	if s.rp == nil {
		if v, ok := s.lfs.Load(owner); ok {
			return v.(*logfs.LogFS), nil
		}
		return nil, fmt.Errorf("no repo")
	}
	return s.getFS(owner, false)
}

func keccakHex(b []byte) string {
	return "0x" + hex.EncodeToString(crypto.Keccak256(b))
}

// objectSha256At returns the sha256 (hex) the LogFS meta records for the row's
// key, only if that meta is this row's object (same volume, start and size) —
// cheap, no content read. Empty otherwise.
func (s *Server) objectSha256At(n types.NeedleDisplay) string {
	if n.Owner == "" {
		return ""
	}
	fs, err := s.readFS(n.Owner)
	if err != nil {
		return ""
	}
	m, err := fs.GetMeta([]byte(n.Name))
	if err != nil || len(m.Hash) == 0 || m.Index != n.File || m.Start != n.Start || m.Size != n.Size {
		return ""
	}
	return hex.EncodeToString(m.Hash)
}

// v1WaitCommitted polls until the object's volume lands on-chain (Piece set) or
// a bounded deadline (HUB_V1_WAIT_SEC, default 90s) / request cancellation.
func (s *Server) v1WaitCommitted(ctx context.Context, owner, bucket, key string) (types.NeedleDisplay, bool) {
	maxWait := time.Duration(env.Int("HUB_V1_WAIT_SEC", 90)) * time.Second
	timeout := time.After(maxWait)
	tick := time.NewTicker(2 * time.Second)
	defer tick.Stop()
	for {
		nds, _ := s.getNeedleDisplay(owner, bucket, key)
		if len(nds) > 0 && nds[0].Piece != "" {
			return nds[0], true
		}
		select {
		case <-ctx.Done():
			return types.NeedleDisplay{}, false
		case <-timeout:
			return types.NeedleDisplay{}, false
		case <-tick.C:
		}
	}
}

// ---- helpers ---------------------------------------------------------------

func v1Limit(c *gin.Context) int {
	n, _ := strconv.Atoi(c.Query("limit"))
	if n <= 0 || n > 100 {
		n = 32
	}
	return n
}

// v1Page reads ?offset= and ?length=: length defaults to and is capped at max,
// and negative values are refused (a negative length removes the SQL LIMIT).
// On a bad value it has already written a 400 and returns ok=false.
func v1Page(c *gin.Context, max int) (offset, length int, ok bool) {
	offset, _ = strconv.Atoi(c.Query("offset"))
	length, _ = strconv.Atoi(c.Query("length"))
	if offset < 0 || length < 0 {
		c.JSON(http.StatusBadRequest, lerror.ToAPIError("hub", fmt.Errorf("offset and length must not be negative")))
		return 0, 0, false
	}
	if length == 0 || length > max {
		length = max
	}
	return offset, length, true
}

// maxStatDays bounds GET /v1/stats?days=.
const maxStatDays = 366
