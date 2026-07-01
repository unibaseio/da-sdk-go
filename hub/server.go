package hub

import (
	"context"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	contract "github.com/unibaseio/da-sdk-go/contract/v2"
	"github.com/unibaseio/da-sdk-go/docs"
	"github.com/unibaseio/da-sdk-go/lib/log"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
	"github.com/unibaseio/da-sdk-go/lib/piece"
	"github.com/unibaseio/da-sdk-go/lib/repo"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/lib/utils"
	"github.com/unibaseio/da-sdk-go/sdk"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	sync.RWMutex
	fscnt uint32
	lfs   map[string]*logfs.LogFS

	local common.Address

	auth types.Auth

	statManager *StatManager

	bucketDisplayLock sync.RWMutex
	bucketDisplay     map[string]types.BucketDisplay

	// negative cache of download keys confirmed missing (download-flood guard)
	missCache *missCache

	// lazily-built chain client for the /api/seal path (hub-signed AddPiece)
	cmMu sync.Mutex
	cm   *contract.ContractManage

	// cached per-owner memory stats, recomputed in the background
	memStat *memStatCache

	// readonly = a reader replica (HUB_READONLY): shares the index DB but does not
	// own local writes — skips upload routes, chain submitter, writer loop, DDL.
	readonly bool

	httpServer *http.Server

	// Add channels for graceful shutdown
	shutdownChan   chan struct{}
	checkpointStop chan struct{}

	// uploadNotify wakes the uploadTo drain loop when new data is written
	// (event-driven), so a write isn't stuck behind the periodic tick. Buffered
	// size 1 + non-blocking send = a coalescing signal (never blocks the writer).
	uploadNotify chan struct{}
}

func NewServer(rp repo.Repo) (*Server, error) {
	log.SetLogLevel("DEBUG")

	gin.SetMode(gin.ReleaseMode)

	localAddr := rp.Key().Address()

	logger.Infof("hub %s starting...", localAddr)

	router := gin.Default()

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

		bucketDisplay: make(map[string]types.BucketDisplay),
		lfs:           make(map[string]*logfs.LogFS),

		missCache: newMissCache(),
		memStat:   &memStatCache{},

		readonly: os.Getenv("HUB_READONLY") != "",

		shutdownChan:   make(chan struct{}),
		checkpointStop: make(chan struct{}),
		uploadNotify:   make(chan struct{}, 1),
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
	// create it (so /api/stat doesn't nil-panic) but don't start the writer.
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
	// (dashboard), so the ALB routes /api/memoryStat + /api/memoryOverview to the
	// writer (like uploads); a replica that gets one serves an empty snapshot.
	if !s.readonly {
		s.startMemStats(context.Background())
	}

	s.registRoute()

	s.httpServer = &http.Server{
		Addr:    rp.Config().API.Endpoint,
		Handler: s.Router,
	}

	// Setup signal handler for emergency shutdown
	s.SetupSignalHandler()

	return s, nil
}

func (s *Server) registRoute() {
	swaghost := s.rp.Config().API.Expose
	if swaghost != "" {
		swaghost = strings.TrimPrefix(swaghost, "http://")
		docs.SwaggerInfo.Host = swaghost
	}

	s.Router.Use(Cors())

	s.Router.Use(ginzap.Ginzap(log.Logger("gin").Desugar(), time.RFC3339, true))

	s.Router.Use(static.Serve("/", static.LocalFile("assets", true)))

	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public, read-only /api group: body-size cap + per-IP rate limit, but NO
	// Authorization required — these endpoints are browsable like a block
	// explorer. Stored content is client-encrypted, so listing/downloading
	// exposes only ciphertext plus metadata. The per-IP limit (generous by
	// default) throttles a single abusive host without blocking the shared
	// explorer reverse-proxy IP; the download negative cache is the primary
	// flood absorber.
	pub := s.Router.Group("/api")
	pub.Use(MaxBodySize())
	pub.Use(RateLimit())

	s.addInfo(pub)
	s.addDownload(pub)
	s.addList(pub)
	s.addConversation(pub)
	s.addStat(pub)

	// Mutating /api group: body-size cap → auth → rate limit. AuthMiddleware
	// requires a valid signed Authorization header, and each handler further
	// enforces owner == signer (RequireOwnerMatch).
	authed := s.Router.Group("/api")
	authed.Use(MaxBodySize())
	authed.Use(AuthMiddleware())
	authed.Use(RateLimit())

	if s.readonly {
		// reader replicas must not write locally — reject writes (ALB routes writes to the primary).
		s.addUploadReadonly(authed)
	} else {
		s.addUpload(authed)
		s.addSeal(authed) // seal is a write path (Upload + AddPiece)
	}
}

// isSQLite reports whether the gorm backend is SQLite (vs Postgres).
func (s *Server) isSQLite() bool {
	return s.gdb != nil && s.gdb.Dialector.Name() == "sqlite"

}

// ListenAndServe starts the HTTP server
func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down both the HTTP server and persists data
func (s *Server) Shutdown(ctx context.Context) error {
	logger.Info("starting server shutdown...")

	// Signal checkpoint routine to stop and perform final checkpoint
	close(s.checkpointStop)

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
	s.Lock()
	for addr, lfs := range s.lfs {
		logger.Infof("closing LogFS for address: %s", addr)
		if err := lfs.Close(); err != nil {
			logger.Errorf("failed to close LogFS for %s: %v", addr, err)
		}
	}
	s.Unlock()

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
	return nil
}

func login(url string, auth types.Auth) {
	for {
		sdk.Login(url, auth)
		time.Sleep(time.Hour)
	}
}

func (s *Server) register() error {
	auth, err := s.rp.Key().BuildAuth([]byte("register"))
	if err != nil {
		return err
	}

	go login(s.rp.Config().Remote.URL, auth)

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

func (s *Server) addInfo(g *gin.RouterGroup) {
	g.Group("/").GET("/info", func(c *gin.Context) {
		res := types.EdgeReceipt{
			EdgeMeta: types.EdgeMeta{
				Type: s.typ,
				Name: s.local,
			},
		}

		c.JSON(http.StatusOK, res)
	})
}
