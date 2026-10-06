package hub

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	contract "github.com/unibaseio/da-sdk-go/contract/v2"
	"github.com/unibaseio/da-sdk-go/lib/env"
	"github.com/unibaseio/da-sdk-go/lib/log"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
	"github.com/unibaseio/da-sdk-go/lib/piece"
	"github.com/unibaseio/da-sdk-go/lib/repo"
	"github.com/unibaseio/da-sdk-go/lib/s3vol"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/lib/utils"
	"github.com/unibaseio/da-sdk-go/sdk"

	"github.com/ethereum/go-ethereum/common"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var logger = log.Logger("hub")

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

type Server struct {
	Router *gin.Engine

	typ string

	rp repo.Repo

	gdb *gorm.DB

	ps types.IPieceStore

	// per-owner LogFS instances. Lookups are lock-free (sync.Map); creation is
	// deduplicated per-owner by fsSF so logfs.New never runs under a global lock
	// (a new owner no longer serializes unrelated owners). See fsmanager.go.
	lfs     sync.Map // addr(string) -> *logfs.LogFS
	fsSF    singleflight.Group
	fscnt   uint32 // number of registered owners (LOGINST count)
	fscntMu sync.Mutex

	// getFS observability: lock-free map hits vs on-demand logfs.New creations.
	fsHit    atomic.Int64
	fsCreate atomic.Int64

	local common.Address

	auth types.Auth

	statManager *StatManager

	// read-through byte LRU of small hot objects (nil when HUB_READCACHE_MB=0)
	readCache *readCache

	// dedupes concurrent DA-download fallbacks for the same (owner,name) so a
	// cold-but-existing key is reconstructed once, not once per request.
	dlSF singleflight.Group
	// dlSem bounds concurrent DA reconstructs (HUB_DOWNLOAD_CONCURRENCY>0);
	// nil = unlimited = historical behavior. dlTotal/dlShared count fallbacks
	// and how many were served by an in-flight flight (singleflight coalescing).
	dlSem    chan struct{}
	dlTotal  atomic.Int64
	dlShared atomic.Int64

	// pieceSem bounds concurrent /v1/pieces/{cid}/content reads, rebuilds and
	// response writes (a piece may be ~1 GB); the write is time-bounded, see
	// pieceWriteTimeout. HUB_PIECE_DOWNLOAD_CONCURRENCY.
	pieceSem chan struct{}

	// quota is the per-signer budget for hub-paid writes (nil = off); sealSem
	// bounds concurrent seals (HUB_SEAL_CONCURRENCY). See quota.go / seal.go.
	quota   *writeQuota
	sealSem chan struct{}

	// cached Piece-contract store-duration bounds (minStore/maxStore) for seal
	storeDurMu sync.Mutex
	storeDur   storeDuration

	// P3-S: shared S3/MinIO backend for sealed volumes (nil = local-only buffer,
	// the default). Bound per-owner and passed to logfs.New via getFS/load.
	volStore *s3vol.Store

	// P4-Route: owner-sharded sticky write routing (nil = single-node, default).
	shard *shardRouter

	// lazily-built chain client for the /v1/seal path (hub-signed AddPiece)
	cmMu sync.Mutex
	cm   *contract.ContractManage

	// cached per-owner memory stats, recomputed in the background
	memStat *memStatCache

	// cached grand totals for /v1 list withTotal (COUNT over 34M+ rows)
	totals *totalCache

	// readonly = a reader replica (HUB_READONLY): shares the index DB but does not
	// own local writes — skips upload routes, chain submitter, writer loop, DDL.
	readonly bool

	httpServer *http.Server

	// Add channels for graceful shutdown
	shutdownChan   chan struct{}
	checkpointStop chan struct{}
	// shutdownOnce serializes Shutdown: the signal handler and the daemon both
	// call it on SIGINT/SIGTERM, and the second caller must wait, not re-close.
	shutdownOnce sync.Once

	// uploadNotify wakes the uploadTo drain loop when new data is written
	// (event-driven), so a write isn't stuck behind the periodic tick. Buffered
	// size 1 + non-blocking send = a coalescing signal (never blocks the writer).
	uploadNotify chan struct{}
}

func NewServer(rp repo.Repo) (*Server, error) {
	if len(siweDomains()) == 0 {
		logger.Warn("HUB_SIWE_DOMAINS is not set: SIWE sign-ins issued for any website are accepted")
	}
	log.SetLogLevel("DEBUG")

	gin.SetMode(gin.ReleaseMode)

	localAddr := rp.Key().Address()

	logger.Infof("hub %s starting...", localAddr)

	router := gin.Default()
	// Allow %2F-encoded slashes in /v1 path params (object keys / resource names
	// can contain "/"). Match on the raw path, decode the param value back.
	router.UseRawPath = true
	router.UnescapePathValues = true
	// client IP for the per-IP limiter: honor X-Forwarded-For only from trusted
	// proxies (HUB_TRUSTED_PROXIES, default private+loopback — see ratelimit.go)
	if err := configureClientIP(router); err != nil {
		return nil, err
	}

	auth, err := rp.Key().BuildAuth([]byte("hub"))
	if err != nil {
		return nil, err
	}

	s := &Server{
		Router: router,

		typ:   types.HubType,
		local: localAddr,
		rp:    rp,
		ps:    piece.New(rp.MetaStore(), rp.DataStore()),
		auth:  auth,

		readCache: newReadCache(),
		memStat:   &memStatCache{},
		totals:    newTotalCache(),

		readonly: os.Getenv("HUB_READONLY") != "",

		shutdownChan:   make(chan struct{}),
		checkpointStop: make(chan struct{}),
		uploadNotify:   make(chan struct{}, 1),
	}

	// Optional cap on concurrent DA reconstructs (expensive K-of-N fetches).
	// Default 0 = unlimited = historical behavior; singleflight already collapses
	// same-key floods, this bounds distinct-key fan-out under a read storm.
	if n := env.Int("HUB_DOWNLOAD_CONCURRENCY", 0); n > 0 {
		s.dlSem = make(chan struct{}, n)
	}
	// Public piece reads rebuild up to ~1 GB each, so unlike the object path
	// they are bounded by default.
	if n := env.Int("HUB_PIECE_DOWNLOAD_CONCURRENCY", defaultPieceDownloadConcurrency); n > 0 {
		s.pieceSem = make(chan struct{}, n)
	}
	s.quota = newWriteQuota()
	if n := env.Int("HUB_SEAL_CONCURRENCY", defaultSealConcurrency); n > 0 {
		s.sealSem = make(chan struct{}, n)
	}

	// P3-S: optional durable/shared sealed-volume backend (HUB_BUFFER=s3). Default
	// "local" keeps volumes on local disk only (unchanged). When s3, each owner's
	// LogFS uploads sealed volumes to S3/MinIO and reads fall back there when a
	// local volume is gone — enabling crash recovery + cross-replica reads.
	if strings.EqualFold(env.Str("HUB_BUFFER", "local"), "s3") {
		st, verr := s3vol.NewFromEnv()
		if verr != nil {
			return nil, fmt.Errorf("HUB_BUFFER=s3: %w", verr)
		}
		pctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		if perr := st.Ping(pctx); perr != nil {
			cancel()
			return nil, fmt.Errorf("HUB_BUFFER=s3 bucket unreachable: %w", perr)
		}
		cancel()
		s.volStore = st
		logger.Infof("HUB_BUFFER=s3: sealed volumes backed by S3 bucket")
	}

	// P4-Route: owner-sharded sticky write routing (nil unless HUB_SHARD_TOTAL>1).
	s.shard, err = newShardRouter()
	if err != nil {
		return nil, err
	}

	// P4: shared L2 read cache (Redis). Fail-open — an unreachable Redis degrades
	// to L1-only, so a bad ping is a warning, not a startup failure.
	if s.readCache != nil {
		if _, on := s.readCache.L2Stats(); on {
			if perr := s.readCache.l2.ping(); perr != nil {
				logger.Warnf("HUB_REDIS_ADDR set but Redis unreachable (L2 fails open to L1-only): %v", perr)
			} else {
				logger.Infof("read cache L2 (Redis) enabled")
			}
		}
	}

	if s.readonly {
		logger.Warn("HUB_READONLY set: running as a read-only replica (no writes, no chain submit, no schema DDL)")
	}

	err = s.register()
	if err != nil {
		return nil, err
	}

	s.load()

	s.loadGORM()

	// StatManager's background loop writes StatRecord; on a reader replica we
	// create it (so /v1/stats doesn't nil-panic) but don't start the writer.
	sm := NewStatManager(s.gdb)
	if !s.readonly {
		err = sm.Start(context.Background())
		if err != nil {
			return nil, err
		}
	}
	s.statManager = sm

	// chain submission is a write path — writer only.
	if !s.readonly {
		go s.uploadTo()
	}

	// memory-stats recompute is a full-index scan over the needles table. Run it
	// on the WRITER only: if every replica ran it independently against the shared
	// DB, the heavy scan would be multiplied N times. Stats are low-QPS
	// (dashboard), so the ALB routes /v1/overview to the writer (like uploads);
	// a replica that gets one serves an empty snapshot.
	if !s.readonly {
		s.startMemStats(context.Background())
	}

	// keep the hot global totals warm (needles/buckets grand counts)
	s.startTotalRefresh(context.Background())

	s.registRoute()

	s.httpServer = newHTTPServer(rp.Config().API.Endpoint, s.Router)

	// Setup signal handler for emergency shutdown
	s.SetupSignalHandler()

	return s, nil
}

func (s *Server) registRoute() {
	s.Router.Use(Cors())

	s.Router.Use(ginzap.Ginzap(log.Logger("gin").Desugar(), time.RFC3339, true))

	// Single clean, resource-oriented /v1 surface (no legacy /api) — same as the
	// gateway + nodes. Public reads + signed writes; content is client-encrypted,
	// so browsable reads expose only ciphertext + metadata. Object CRUD is S3-shaped
	// (buckets/objects); info/conversations are resources; seal is a /v1 operation.
	// See GATEWAY_API_V1_SPEC.md (Hub §) + v1.go.
	s.registV1()
}

// isSQLite reports whether the gorm backend is SQLite (vs Postgres).
func (s *Server) isSQLite() bool {
	return s.gdb != nil && s.gdb.Dialector.Name() == "sqlite"

}

// HTTP server timeouts (env-tunable, seconds; 0 disables one):
//   - HUB_HTTP_READ_HEADER_TIMEOUT_SEC (10): slowloris guard on the request line and
//     headers.
//   - HUB_HTTP_IDLE_TIMEOUT_SEC (120): keep-alive idle; above the ALB's default
//     60s idle timeout so the ALB, not the hub, closes idle connections (the
//     other way round yields sporadic 502s).
//   - HUB_HTTP_READ_TIMEOUT_SEC (600): whole request including the body. Uploads
//     are capped at HUB_MAX_MULTIPART_BYTES (64 MiB), so 10 min still admits a
//     ~110 KB/s client.
//   - HUB_HTTP_WRITE_TIMEOUT_SEC (0 = none): a write deadline would cut off
//     legitimate long responses — piece downloads up to ~1 GB, seal (encode +
//     on-chain wait) and ?wait=1 commits — so it stays off by default; those
//     handlers are bounded by their own contexts and semaphores instead.
const (
	defaultHTTPReadHeaderTimeoutSec = 10
	defaultHTTPIdleTimeoutSec       = 120
	defaultHTTPReadTimeoutSec       = 600
	defaultHTTPWriteTimeoutSec      = 0
)

func newHTTPServer(addr string, h http.Handler) *http.Server {
	sec := func(k string, def int) time.Duration {
		return time.Duration(env.Int(k, def)) * time.Second
	}
	return &http.Server{
		Addr:              addr,
		Handler:           h,
		ReadHeaderTimeout: sec("HUB_HTTP_READ_HEADER_TIMEOUT_SEC", defaultHTTPReadHeaderTimeoutSec),
		ReadTimeout:       sec("HUB_HTTP_READ_TIMEOUT_SEC", defaultHTTPReadTimeoutSec),
		WriteTimeout:      sec("HUB_HTTP_WRITE_TIMEOUT_SEC", defaultHTTPWriteTimeoutSec),
		IdleTimeout:       sec("HUB_HTTP_IDLE_TIMEOUT_SEC", defaultHTTPIdleTimeoutSec),
		MaxHeaderBytes:    1 << 20,
	}
}

// ListenAndServe starts the HTTP server
func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down both the HTTP server and persists data. It is
// safe to call more than once and concurrently: the work runs once, and later
// callers block until it has finished.
func (s *Server) Shutdown(ctx context.Context) error {
	s.shutdownOnce.Do(func() { s.shutdown(ctx) })
	return nil
}

func (s *Server) shutdown(ctx context.Context) {
	logger.Info("starting server shutdown...")

	// Signal checkpoint routine to stop and perform final checkpoint
	if s.checkpointStop != nil {
		close(s.checkpointStop)
	}
	// signal the drain loop to stop (an in-flight drain is not waited for)
	// before its LogFS instances are closed below
	if s.shutdownChan != nil {
		close(s.shutdownChan)
	}

	// First shutdown the HTTP server
	if s.httpServer != nil {
		logger.Info("shutting down HTTP server...")
		if err := s.httpServer.Shutdown(ctx); err != nil {
			logger.Errorf("failed to shutdown HTTP server: %v", err)
		}
	}

	// Stop the statistics manager (Stop() writes a final record — writer only)
	if s.statManager != nil && !s.readonly {
		logger.Info("stopping statistics manager...")
		s.statManager.Stop()
	}

	// Close all LogFS instances
	s.lfs.Range(func(k, v any) bool {
		addr := k.(string)
		lfs := v.(*logfs.LogFS)
		logger.Infof("closing LogFS for address: %s", addr)
		if err := lfs.Close(); err != nil {
			logger.Errorf("failed to close LogFS for %s: %v", addr, err)
		}
		return true
	})

	// Close the shared L2 read cache (Redis), if any.
	if s.readCache != nil && s.readCache.l2 != nil {
		s.readCache.l2.close()
	}

	// Persist database data
	if s.gdb != nil {
		logger.Info("persisting database data...")

		// Get the underlying SQL database
		sqlDB, err := s.gdb.DB()
		if err != nil {
			logger.Errorf("failed to get SQL database: %v", err)
		} else {
			// SQLite-only WAL persistence; Postgres manages its own durability.
			if s.isSQLite() {
				logger.Info("executing SQLite persistence commands...")

				// Force a final checkpoint to ensure WAL data is written to main database
				if err := s.gdb.Exec("PRAGMA wal_checkpoint(FULL);").Error; err != nil {
					logger.Errorf("failed to execute final WAL checkpoint: %v", err)
				}

				// Synchronize data to disk
				if err := s.gdb.Exec("PRAGMA synchronous = FULL;").Error; err != nil {
					logger.Errorf("failed to set synchronous mode: %v", err)
				}

				// Force fsync to ensure data is written to disk
				if err := s.gdb.Exec("PRAGMA wal_checkpoint(TRUNCATE);").Error; err != nil {
					logger.Errorf("failed to truncate WAL: %v", err)
				}
			}

			// Close the SQL database connection
			if err := sqlDB.Close(); err != nil {
				logger.Errorf("failed to close SQL database: %v", err)
			} else {
				logger.Info("database connection closed successfully")
			}
		}
	}

	// Close repository resources
	if s.rp != nil {
		logger.Info("closing repository...")
		if err := s.rp.Close(); err != nil {
			logger.Errorf("failed to close repository: %v", err)
		}
	}

	logger.Info("server shutdown completed")
}

// login re-announces this node to the gateway hourly, signing each time: the
// gateway rejects signatures older than ~10 minutes, and one reused header
// would be replayable by anyone who saw it.
func login(url string, sign sdk.Signer) {
	for {
		if auth, err := sign([]byte("login")); err == nil {
			sdk.Login(url, auth)
		}
		time.Sleep(time.Hour)
	}
}

func (s *Server) register() error {
	auth, err := s.rp.Key().BuildAuth([]byte("register"))
	if err != nil {
		return err
	}

	go login(s.rp.Config().Remote.URL, s.rp.Key().BuildAuth)

	mm := types.EdgeMeta{
		Type:      s.typ,
		Name:      auth.Addr,
		PublicKey: s.rp.Key().Public(),
		ExposeURL: s.rp.Config().API.Expose,
		Hardware:  utils.GetHardwareInfo(),
		ChainType: s.rp.Config().Chain.Type,
	}

	err = sdk.RegisterEdge(s.rp.Config().Remote.URL, auth, mm)
	if err != nil {
		logger.Debug("register hub fail:", err)
		return err
	}
	return nil
}
