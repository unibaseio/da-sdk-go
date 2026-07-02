package hub

// /v1 — resource-oriented, wallet-native object-store API (spec: da/HUB_API_V1_SPEC.md).
// S3 SHAPE (bucket/object/key) but web3/DA-native semantics: kind on the bucket,
// verifiable receipt (commitment + chain + status), async staged→committed.
//
// This is a THIN FAÇADE over the same Server (gdb + logfs) that /api/* uses —
// old /api/* is untouched; an object written via either surface is visible from
// the other (one storage engine, two API façades).

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/unibaseio/da-sdk-go/lib/env"
	lerror "github.com/unibaseio/da-sdk-go/lib/error"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

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
	Key          string     `json:"key"`
	Size         uint64     `json:"size"`
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
	r := v1Receipt{Bucket: bucket, Key: n.Name, Size: n.Size, Status: "staged", Availability: "pending"}
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

// registV1 mounts the /v1 surface. Public reads (no auth) + authed writes,
// mirroring the /api group split; same Server/gdb/logfs underneath.
func (s *Server) registV1() {
	pub := s.Router.Group("/v1")
	pub.Use(RateLimit())
	pub.GET("/buckets", s.v1ListBuckets)
	pub.GET("/buckets/:bucket", s.v1GetBucket)
	pub.GET("/buckets/:bucket/objects", s.v1ListObjects)
	pub.GET("/buckets/:bucket/objects/:key", s.v1GetObject)
	pub.GET("/buckets/:bucket/objects/:key/content", s.v1GetObjectContent)
	pub.GET("/stats", s.v1Stats)
	pub.GET("/overview", s.v1Overview)
	pub.GET("/owners/:owner", s.v1GetOwner)

	w := s.Router.Group("/v1")
	w.Use(maxBodyV1())
	w.Use(AuthMiddleware())
	w.Use(RateLimit())
	if s.readonly {
		reject := func(c *gin.Context) {
			c.JSON(http.StatusServiceUnavailable, lerror.ToAPIError("hub", fmt.Errorf("this node is read-only; route writes to the primary")))
		}
		w.PUT("/buckets/:bucket", reject)
		w.PUT("/buckets/:bucket/objects/:key", reject)
		w.POST("/buckets/:bucket/objects", reject)
	} else {
		w.PUT("/buckets/:bucket", s.v1PutBucket)
		w.PUT("/buckets/:bucket/objects/:key", s.v1PutObject)
		w.POST("/buckets/:bucket/objects", s.v1PostObjects)
	}
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

// GET /v1/buckets?owner=&kind=&cursor=&limit=
func (s *Server) v1ListBuckets(c *gin.Context) {
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	limit := v1Limit(c)
	res, err := s.listBucket(owner, c.Query("kind"), 0, limit) // reuse; offset 0 (buckets table small)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"buckets": res})
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

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	if len(data) == 0 {
		abortWithBadRequest(c, fmt.Errorf("empty body"))
		return
	}

	large := v1KindProfileLarge(b.Kind, int64(len(data)))
	if _, err := s.logFSWriteEx(owner, bucket, key, b.Kind, large, strings.NewReader(string(data))); err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	// staged: on logfs + indexed; committed later by drainInstance's AddPiece.
	c.JSON(http.StatusAccepted, v1Receipt{Bucket: bucket, Key: key, Size: uint64(len(data)), Status: "staged", Availability: "pending"})
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
	receipts := make([]v1Receipt, 0, len(files))
	for _, fh := range files {
		if fh.Size == 0 {
			abortWithBadRequest(c, fmt.Errorf("empty file: %s", fh.Filename))
			return
		}
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
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	bucket := c.Param("bucket")
	limit := v1Limit(c)

	// cursor pagination on the needles table by id (avoids offset walk on 34M rows):
	// WHERE bucket=? [AND LOWER(owner)=?] [AND id < cursor] ORDER BY id DESC LIMIT n
	q := s.gdb.Model(&types.Needle{}).Where("bucket = ?", bucket)
	if owner != "" {
		q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
	}
	if cur := c.Query("cursor"); cur != "" {
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
	if len(needles) == limit {
		next = strconv.FormatUint(uint64(needles[len(needles)-1].ID), 10)
	}
	c.JSON(http.StatusOK, gin.H{"objects": objs, "nextCursor": next})
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
	c.JSON(http.StatusOK, receiptFromNeedle(bucket, nds[0]))
}

// GET /v1/buckets/{bucket}/objects/{key}/content — download bytes.
func (s *Server) v1GetObjectContent(c *gin.Context) {
	owner, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}
	key := c.Param("key")
	var w strings.Builder
	// logFSRead resolves owner candidates (lowercase + legacy checksum); empty
	// owner falls back to a needle-name lookup.
	if _, err := s.logFSRead(owner, key, &w); err != nil {
		c.JSON(http.StatusNotFound, lerror.ToAPIError("hub", err))
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", []byte(w.String()))
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

// ---- helpers ---------------------------------------------------------------

func v1Limit(c *gin.Context) int {
	n, _ := strconv.Atoi(c.Query("limit"))
	if n <= 0 || n > 100 {
		n = 32
	}
	return n
}
