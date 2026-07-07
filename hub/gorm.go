package hub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func (s *Server) loadGORM() {
	cfg := &gorm.Config{
		Logger:                 glogger.Default.LogMode(glogger.Silent),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	}

	// Backend selection (stage 2: shared index for multi-instance read scaling).
	//   HUB_DB_DRIVER=postgres + HUB_DB_DSN=...  → shared Postgres (RDS)
	//   default                                  → local SQLite (single node)
	driver := strings.ToLower(os.Getenv("HUB_DB_DRIVER"))
	dsn := os.Getenv("HUB_DB_DSN")
	if driver == "" && dsn != "" {
		driver = "postgres"
	}

	var db *gorm.DB
	var err error
	switch driver {
	case "postgres", "pg":
		if dsn == "" {
			panic("HUB_DB_DRIVER=postgres requires HUB_DB_DSN")
		}
		db, err = gorm.Open(postgres.Open(dsn), cfg)
		if err != nil {
			panic("failed to connect postgres: " + err.Error())
		}
		logger.Info("gorm backend: postgres (shared index)")
	default:
		gpath := filepath.Join(s.rp.Path(), "gorm")
		os.MkdirAll(gpath, os.ModePerm)
		gpath = filepath.Join(gpath, "gorm.db")
		db, err = gorm.Open(sqlite.Open(gpath), cfg)
		if err != nil {
			panic("failed to connect database")
		}
		// SQLite-only PRAGMAs (WAL, cache, mmap, busy timeout)
		_ = db.Exec("PRAGMA journal_mode=WAL;")
		_ = db.Exec("PRAGMA synchronous = NORMAL;") // NORMAL provides good balance of safety and performance
		_ = db.Exec("PRAGMA cache_size = -64000;")  // 64MB cache (reduced for more frequent writes)
		_ = db.Exec("PRAGMA temp_store = MEMORY;")
		_ = db.Exec("PRAGMA mmap_size = 4000000000;")    // 4GB mmap
		_ = db.Exec("PRAGMA wal_autocheckpoint = 1000;") // Less frequent checkpoints to reduce lock contention
		_ = db.Exec("PRAGMA busy_timeout = 60000;")      // Increase timeout to 60 seconds for better handling of concurrent operations
		logger.Infof("gorm backend: sqlite at %s", gpath)
	}

	sqldb, err := db.DB()
	if err != nil {
		panic("failed to get database")
	}

	sqldb.SetMaxIdleConns(5)                   // Reduce idle connections to minimize lock contention
	sqldb.SetMaxOpenConns(20)                  // Reduce max connections to prevent overwhelming SQLite
	sqldb.SetConnMaxLifetime(15 * time.Minute) // Shorter connection lifetime for better resource management

	s.gdb = db

	// Read/write split: when HUB_DB_DSN_READ is set (e.g. the Aurora reader
	// endpoint) register gorm's dbresolver so SELECTs go to the read replica and
	// writes + DDL stay on the writer (the main connection opened above). Unset →
	// single DB, behavior unchanged. Postgres only. Reads that GATE a write
	// (existence / uniqueness checks) must still hit the writer to avoid replica-lag
	// races — those call sites use .Clauses(dbresolver.Write).
	if (driver == "postgres" || driver == "pg") && os.Getenv("HUB_DB_DSN_READ") != "" {
		readDSN := os.Getenv("HUB_DB_DSN_READ")
		if err := db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{postgres.Open(readDSN)},
			Policy:   dbresolver.RandomPolicy{},
		}).
			// Bound + recycle the replica pool (the writer/source pool already gets
			// these via sqldb above). ConnMaxLifetime matters for Aurora read
			// auto-scaling: the reader CLUSTER endpoint load-balances only on NEW
			// connections, so without periodic recycling the pool stays pinned to the
			// current reader(s) and a freshly scaled-out replica never receives traffic.
			SetConnMaxLifetime(10 * time.Minute).
			SetMaxIdleConns(5).
			SetMaxOpenConns(20)); err != nil {
			panic("failed to register read replica (HUB_DB_DSN_READ): " + err.Error())
		}
		logger.Info("dbresolver: reads → HUB_DB_DSN_READ (replica), writes/DDL → writer")
	}

	// Schema + indexes: writer only. Reader replicas (HUB_READONLY) share the
	// same DB, so running AutoMigrate / CREATE INDEX from every replica would
	// race on DDL (Postgres) — and a reader has nothing to create. The covering
	// expression index (LOWER(owner), size) lets the memory-stats aggregation
	// and the WHERE LOWER(owner)=? filters run index-only; building it on a very
	// large table is a one-time, possibly multi-minute, write-locking op.
	if !s.readonly {
		db.AutoMigrate(&types.Account{})
		db.AutoMigrate(&types.Bucket{})
		db.AutoMigrate(&types.Needle{})
		db.AutoMigrate(&types.Volume{})
		db.AutoMigrate(&types.StatRecord{})
		db.AutoMigrate(&types.Conversation{})

		// Reads filter by LOWER(owner) (+ optional bucket/name) and order by id, so
		// only the LOWER(owner) expression indexes plus name/bucket are useful. The
		// raw (owner, …) composites can't serve LOWER(owner)=? — they were dead
		// weight that slowed every INSERT on this 34M-row table. Drop them (also on
		// existing DBs) and keep just the indexes the query planner actually uses.
		for _, dead := range []string{
			"idx_needles_owner", "idx_needles_owner_name",
			"idx_needles_owner_bucket", "idx_needles_owner_bucket_name",
		} {
			db.Exec("DROP INDEX IF EXISTS " + dead)
		}
		db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_bucket ON needles(bucket);")
		db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_name ON needles(name);")

		// The LOWER(owner) expression indexes are PARTIAL on `deleted_at IS NULL` to
		// match the predicate gorm injects into every soft-delete query. Without it
		// the planner can't go index-only (it must hit the heap to check deleted_at),
		// so memoryStat's GROUP BY fell back to a full 34M-row seq scan every refresh.
		// With it: memoryStat aggregates index-only, and listNeedle's offset-walk stays
		// in-index. Build the partial versions, THEN drop the old full ones — that
		// ordering keeps an owner index available throughout the one-time rebuild
		// (multi-minute on a large table, writer-startup only). Partial indexes work on
		// both Postgres and SQLite. NOTE: after the rebuild + the bulk load, run
		// `VACUUM ANALYZE needles;` so the visibility map is set and the scan is truly
		// index-only.
		db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_lower_owner_size_live ON needles(LOWER(owner), size) WHERE deleted_at IS NULL;")
		db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_lower_owner_id_live ON needles(LOWER(owner), id) WHERE deleted_at IS NULL;")
		// Old full (non-partial) versions, superseded by the *_live partials above.
		// DROP IF EXISTS is a no-op once gone, so this stays idempotent across restarts.
		db.Exec("DROP INDEX IF EXISTS idx_needles_lower_owner_size")
		db.Exec("DROP INDEX IF EXISTS idx_needles_lower_owner_id")

		// getVolume(owner,file) is called once per needle in the list paths (N+1);
		// index volumes so each lookup is a seek, not a full-table scan.
		db.Exec("CREATE INDEX IF NOT EXISTS idx_volumes_lower_owner_file ON volumes(LOWER(owner), file);")

		// Small reference tables today, but addBucket/addAccount do a name lookup on
		// the WRITE path (every upload), and the list endpoints filter by LOWER(owner);
		// index them so those stay seeks instead of seq scans as the tables grow.
		db.Exec("CREATE INDEX IF NOT EXISTS idx_buckets_name ON buckets(name);")
		db.Exec("CREATE INDEX IF NOT EXISTS idx_accounts_name ON accounts(name);")
		db.Exec("CREATE INDEX IF NOT EXISTS idx_conversations_name ON conversations(name);")
		db.Exec("CREATE INDEX IF NOT EXISTS idx_buckets_lower_owner_id ON buckets(LOWER(owner), id);")
		db.Exec("CREATE INDEX IF NOT EXISTS idx_conversations_lower_owner_id ON conversations(LOWER(owner), id);")

		// Ensure a PRIMARY KEY on id for every table (Postgres). A gorm/AutoMigrate-
		// created table always has it, but a table BULK-LOADED by an external tool
		// can be missing it: pgloader's "create no indexes" drops the PK too, and
		// AutoMigrate won't add a PK to a pre-existing table. Without it, every
		// ORDER BY id (the global list endpoints) degrades to a full seq scan + sort
		// — e.g. the unscoped /api/listNeedle was ~5s on 34M rows. Add it
		// idempotently. On a large table this is a one-time, blocking build, so the
		// migration should add needles' PK CONCURRENTLY before first start (see
		// docs/stage2-rds-deploy.md); then this just finds it present and skips.
		if driver == "postgres" || driver == "pg" {
			for _, t := range []string{"accounts", "buckets", "needles", "volumes", "stat_records", "conversations"} {
				db.Exec(fmt.Sprintf(`DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conrelid = '%s'::regclass AND contype = 'p') THEN
    ALTER TABLE %s ADD PRIMARY KEY (id);
  END IF;
END $$;`, t, t))
			}
		}
	}

	// Periodic WAL checkpoint is SQLite-only; Postgres self-manages durability.
	if s.isSQLite() {
		go s.periodicCheckpoint()
	}

	// One-time backfill rewrites rows → writer only (and only when requested).
	ni := os.Getenv("NEED_INIT")
	if ni != "" && !s.readonly {
		logger.Info("handle need init")
		var needles []types.Needle
		result := db.Model(&types.Needle{}).Where("name like ? and created_at >= ?",
			"%_0",
			time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC)).Find(&needles)
		if result.Error != nil {
			logger.Error("failed to get needles: ", result.Error)
			return
		}
		for _, needle := range needles {
			if len(needle.Name) <= 2 {
				continue
			}

			if strings.HasSuffix(needle.Name, "_0") {
				name := strings.TrimSuffix(needle.Name, "_0")
				db.Save(&types.Conversation{
					Name:   name,
					Owner:  needle.Owner,
					Bucket: needle.Bucket,
				})
				continue
			}
			if needle.Name != "" {
				continue
			}
			if needle.Bucket != "" {
				continue
			}
			fmt.Println("update bucket: ", needle.Name)
			var w bytes.Buffer
			s.logFSRead(needle.Owner, needle.Name, &w)
			if w.Len() > 0 {
				// decode w to json
				var meta map[string]interface{}
				err := json.Unmarshal(w.Bytes(), &meta)
				if err != nil {
					continue
				}
				bucketName, ok := meta["name"].(string)
				if ok {
					s.addBucket(needle.Owner, bucketName, "memory")
					// update needle
					db.Model(&needle).Update("bucket", bucketName)
				} else {
					bucketName := needle.Owner
					s.addBucket(needle.Owner, bucketName, "memory")
					// update needle
					db.Model(&needle).Update("bucket", bucketName)
				}
			}
		}
	}

}

func (s *Server) addAccount(owner string) {
	var account types.Account
	// read-before-write: must see the writer's latest, not a lagging replica.
	result := s.gdb.Clauses(dbresolver.Write).First(&account, "name = ?", owner)
	if result.RowsAffected > 0 {
		logger.Info("already has account: ", owner)
		return
	}

	s.gdb.Create(&types.Account{
		Name: owner,
	})
	logger.Info("create account: ", owner)
}

// TODO: bucket is global unique
// kind = object-store scenario ("memory"/"model"/"dataset"); empty defaults to memory.
func (s *Server) addBucket(owner, bucket, kind string) error {
	if kind == "" {
		kind = "memory"
	}
	var gbucket types.Bucket
	// read-before-write (uniqueness/ownership gate): force the writer so two quick
	// uploads of the same bucket name can't both miss it on a lagging replica.
	result := s.gdb.Clauses(dbresolver.Write).First(&gbucket, "name = ? ", bucket)
	if result.RowsAffected > 0 {
		if !strings.EqualFold(gbucket.Owner, owner) {
			logger.Infof("bucket: %s is owned by %s", bucket, gbucket.Owner)
			return fmt.Errorf("bucket: %s is owned by %s", bucket, gbucket.Owner)
		}
		// backfill kind on a legacy/empty-kind bucket (idempotent, memory unaffected).
		if gbucket.Kind == "" && kind != "memory" {
			s.gdb.Model(&gbucket).Update("kind", kind)
		}
		logger.Info("already has bucket: ", bucket)
		return nil
	}

	s.gdb.Create(&types.Bucket{
		Name:  bucket,
		Owner: owner,
		Kind:  kind,
	})
	logger.Info("create bucket: ", bucket)
	return nil
}

func (s *Server) getAccount(owner string) ([]types.Account, error) {
	var accounts []types.Account
	q := s.gdb
	if owner != "" {
		q = q.Where("LOWER(name) = ?", strings.ToLower(owner))
	}
	result := q.Find(&accounts)
	if result.Error != nil {
		return accounts, result.Error
	}
	return accounts, nil
}

func (s *Server) listAccount(offset, limit int) ([]types.Account, error) {
	var accounts []types.Account
	result := s.gdb.Order("id desc").Limit(limit).Offset(offset).Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}

	return accounts, nil
}

// computeMemStats runs the heavy aggregation ONCE: a single GROUP BY over
// needles yields the full per-owner list, from which the overview totals are
// derived (no separate COUNT/SUM scans). This is a full-table scan and must
// NOT be run inline per request (it 504'd behind the proxy on a large table) —
// it's driven by the background refresh loop and the result is cached.
//
// Grouping/counting use LOWER(owner) so mixed-case and lowercase variants of
// the same wallet merge (served index-only by idx_needles_lower_owner_size_live,
// the partial index whose `deleted_at IS NULL` predicate matches this query).
func (s *Server) computeMemStats() (types.MemoryOverview, []types.MemoryStat, error) {
	var ov types.MemoryOverview

	if err := s.gdb.Model(&types.Account{}).Count(&ov.TotalAddresses).Error; err != nil {
		return ov, nil, err
	}

	owners := []types.MemoryStat{}
	err := s.gdb.Model(&types.Needle{}).
		Select("LOWER(owner) as owner, count(*) as count, COALESCE(SUM(size),0) as bytes").
		Group("LOWER(owner)").
		Order("bytes desc").
		Scan(&owners).Error
	if err != nil {
		return ov, nil, err
	}

	var totalCount, totalBytes int64
	for i := range owners {
		owners[i].GB = float64(owners[i].Bytes) / 1e9
		totalCount += owners[i].Count
		totalBytes += owners[i].Bytes
	}
	ov.WalletsWithMemory = int64(len(owners))
	ov.MemoryCount = totalCount
	ov.MemoryBytes = totalBytes
	ov.MemoryGB = float64(totalBytes) / 1e9
	return ov, owners, nil
}

func (s *Server) getBucket(owner, bucket string) ([]types.BucketDisplay, error) {
	var buckets []types.Bucket
	q := s.gdb
	if owner != "" {
		q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
	}
	if bucket != "" {
		q = q.Where("name = ?", bucket)
	}
	result := q.Find(&buckets)
	if result.Error != nil {
		return nil, result.Error
	}
	res := make([]types.BucketDisplay, 0, len(buckets))
	for _, bucket := range buckets {
		s.bucketDisplayLock.RLock()
		bd, ok := s.bucketDisplay[bucket.Name]
		if ok && bucket.UpdatedAt.Equal(bd.Last) {
			res = append(res, bd)
			s.bucketDisplayLock.RUnlock()
			continue
		}
		s.bucketDisplayLock.RUnlock()
		if !ok {
			bd = types.BucketDisplay{
				Bucket: bucket,
				Last:   bucket.UpdatedAt,
			}
		}
		// read data
		needles, err := s.getNeedleByName(bucket.Name)
		if err == nil && len(needles) > 0 {
			if ok && needles[0].UpdatedAt.Equal(bd.Last) {
				res = append(res, bd)
				continue
			}
			bd.Last = needles[0].UpdatedAt

			var w bytes.Buffer
			s.logFSRead(needles[0].Owner, needles[0].Name, &w)
			if w.Len() > 0 {
				// decode w to json
				var mjson map[string]interface{}
				err := json.Unmarshal(w.Bytes(), &mjson)
				if err != nil {
					continue
				}
				description, ok := mjson["content"].(string)
				if ok {
					bd.Description = description
				}

				meta, ok := mjson["metadata"].(map[string]interface{})
				if ok {
					transport, ok := meta["transport"].(string)
					if ok {
						bd.Transport = transport
					}
					typ, ok := meta["type"].(string)
					if ok {
						bd.Type = typ
					}
					state, ok := meta["state"].(string)
					if ok {
						bd.State = state
					}
				}
			}
		}
		s.bucketDisplayLock.Lock()
		s.bucketDisplay[bucket.Name] = bd
		s.bucketDisplayLock.Unlock()
		res = append(res, bd)
	}
	return res, nil
}

func (s *Server) listBucket(owner, kind string, offset, limit int) ([]types.BucketDisplay, error) {
	var buckets []types.Bucket
	q := s.gdb
	if owner != "" {
		q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
	}
	// kind filter: "memory" also matches legacy empty-kind rows; others exact.
	if kind == "memory" {
		q = q.Where("kind = ? OR kind = '' OR kind IS NULL", kind)
	} else if kind != "" {
		q = q.Where("kind = ?", kind)
	}
	result := q.Order("id desc").Limit(limit).Offset(offset).Find(&buckets)
	if result.Error != nil {
		return nil, result.Error
	}

	res := make([]types.BucketDisplay, 0, len(buckets))
	for _, bucket := range buckets {
		s.bucketDisplayLock.RLock()
		bd, ok := s.bucketDisplay[bucket.Name]
		if ok && bucket.CreatedAt.Equal(bd.Last) {
			res = append(res, bd)
			s.bucketDisplayLock.RUnlock()
			continue
		}
		s.bucketDisplayLock.RUnlock()
		if !ok {
			bd = types.BucketDisplay{
				Bucket: bucket,
				Last:   bucket.CreatedAt,
			}
		}
		// read data
		needles, err := s.getNeedleByName(bucket.Name)
		if err == nil && len(needles) > 0 {
			if ok && needles[0].UpdatedAt.Equal(bd.Last) {
				res = append(res, bd)
				continue
			}
			bd.Last = needles[0].UpdatedAt

			var w bytes.Buffer
			s.logFSRead(needles[0].Owner, needles[0].Name, &w)
			if w.Len() > 0 {
				// decode w to json
				var mjson map[string]interface{}
				err := json.Unmarshal(w.Bytes(), &mjson)
				if err != nil {
					continue
				}
				description, ok := mjson["content"].(string)
				if ok {
					bd.Description = description
				}

				meta, ok := mjson["metadata"].(map[string]interface{})
				if ok {
					transport, ok := meta["transport"].(string)
					if ok {
						bd.Transport = transport
					}
					typ, ok := meta["type"].(string)
					if ok {
						bd.Type = typ
					}

					state, ok := meta["state"].(string)
					if ok {
						bd.State = state
					}
				}
			}
		}
		s.bucketDisplayLock.Lock()
		s.bucketDisplay[bucket.Name] = bd
		s.bucketDisplayLock.Unlock()
		res = append(res, bd)
	}

	return res, nil
}

func (s *Server) addNeedle(owner, bucket, name string, findex uint64, start, length uint64) error {
	if err := s.gdb.Create(&types.Needle{
		Owner:  owner,
		Bucket: bucket,
		Name:   name,
		File:   findex,
		Start:  start,
		Size:   length,
	}).Error; err != nil {
		// The blob is already in logfs but now unindexed. Surface the failure so
		// the caller reports the upload as failed (client retries → we re-index),
		// instead of silently returning 200 with a row missing from listNeedle.
		// This matters more on a networked Postgres than on local SQLite.
		logger.Errorf("create needle failed (owner=%s name=%s): %v", owner, name, err)
		return err
	}

	// Ensure the conversation row for this needle's group. The group id is the name
	// up to the LAST "_" — matching getConversation's `name like <conv>_%` query. This
	// works for any "_<index>" scheme (e.g. proto ids zero-pad the seq to 20 digits →
	// "..._00000000000000000000", and sync session ids are "..._<id>"), not just a
	// literal "_0" suffix (the old check missed every padded/suffixed name → grouping
	// stayed empty). Idempotent upsert; runs per needle so any entry heals the group.
	if idx := strings.LastIndex(name, "_"); idx > 0 {
		connName := name[:idx]
		var conversation types.Conversation
		// read-before-write: force the writer so we don't create a duplicate
		// conversation row off a lagging replica.
		result := s.gdb.Clauses(dbresolver.Write).Where(&types.Conversation{Name: connName, Owner: owner, Bucket: bucket}).First(&conversation)
		if result.RowsAffected == 0 {
			// non-fatal: the needle is indexed; conversation grouping can heal later.
			if err := s.gdb.Save(&types.Conversation{
				Name:   connName,
				Owner:  owner,
				Bucket: bucket,
			}).Error; err != nil {
				logger.Errorf("create conversation failed (name=%s): %v", connName, err)
			}
		}
	}
	logger.Info("create needle: ", owner)
	return nil
}

func (s *Server) getNeedleByName(name string) ([]types.Needle, error) {
	var needle []types.Needle
	result := s.gdb.Where(&types.Needle{Name: name}).Order("id desc").Limit(1).Find(&needle)
	if result.Error != nil {
		return needle, result.Error
	}
	return needle, nil
}

func (s *Server) getNeedleDisplay(owner, bucket, name string) ([]types.NeedleDisplay, error) {
	var needle []types.Needle
	q := s.gdb
	if name != "" {
		q = q.Where("name = ?", name)
	}
	if owner != "" {
		q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
	}
	if bucket != "" {
		q = q.Where("bucket = ?", bucket)
	}
	result := q.Order("id desc").Limit(1).Find(&needle)
	if result.Error != nil {
		return nil, result.Error
	}
	res := make([]types.NeedleDisplay, 0, len(needle))
	for i := 0; i < len(needle); i++ {
		nd := types.NeedleDisplay{
			CreatedAt: needle[i].CreatedAt,
			Name:      needle[i].Name,
			Owner:     needle[i].Owner,
			Bucket:    needle[i].Bucket,
			File:      needle[i].File,
			Start:     needle[i].Start,
			Size:      needle[i].Size,
		}
		vol, err := s.getVolume(needle[i].Owner, needle[i].File)
		if err == nil && len(vol) > 0 {
			nd.Piece = vol[0].Piece
			nd.TxHash = vol[0].TxHash
			nd.ChainType = vol[0].ChainType
		}
		res = append(res, nd)
	}

	return res, nil
}

func (s *Server) listNeedleDisplay(owner, bucket string, offset, limit int) ([]types.NeedleDisplay, error) {
	logger.Debug("list needle: ", owner, bucket, offset, limit)
	var needle []types.Needle
	q := s.gdb
	if owner != "" {
		q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
	}
	if bucket != "" {
		q = q.Where("bucket = ?", bucket)
	}
	result := q.Order("id desc").Limit(limit).Offset(offset).Find(&needle)
	if result.Error != nil {
		return nil, result.Error
	}

	vmap := s.volumesFor(needle) // one query instead of getVolume per needle (N+1)
	res := make([]types.NeedleDisplay, 0, len(needle))
	for i := 0; i < len(needle); i++ {
		nd := types.NeedleDisplay{
			CreatedAt: needle[i].CreatedAt,
			Name:      needle[i].Name,
			Owner:     needle[i].Owner,
			Bucket:    needle[i].Bucket,
			File:      needle[i].File,
			Start:     needle[i].Start,
			Size:      needle[i].Size,
		}
		if v, ok := vmap[volKey(needle[i].Owner, needle[i].File)]; ok {
			nd.Piece = v.Piece
			nd.TxHash = v.TxHash
			nd.ChainType = v.ChainType
		}
		res = append(res, nd)
	}

	return res, nil
}

func (s *Server) addVolume(owner string, findex uint64, piece, txn string) {
	s.gdb.Create(&types.Volume{
		ChainType: s.rp.Repo().Config().Chain.Type,
		Owner:     owner,
		File:      findex,
		Piece:     piece,
		TxHash:    txn,
	})
	logger.Info("create volume: ", piece)
}

func (s *Server) getVolume(owner string, fid uint64) ([]types.Volume, error) {
	var vol []types.Volume
	q := s.gdb
	if owner != "" {
		q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
	}
	if fid != 0 {
		q = q.Where("file = ?", fid)
	}
	result := q.Find(&vol)
	if result.Error != nil {
		return vol, result.Error
	}
	return vol, nil
}

// volKey keys the per-(owner,file) volume map; owner lowercased to match how
// volumes are queried.
func volKey(owner string, file uint64) string {
	return fmt.Sprintf("%s:%d", strings.ToLower(owner), file)
}

// volumesFor batch-loads the volumes for a page of needles in ONE query, keyed
// by (LOWER(owner), file) — avoids the N+1 of calling getVolume per needle in
// the list paths. Bounded by the page's distinct owners × files; the map lookup
// picks the exact (owner,file) pair, so any cross-product over-fetch is harmless.
func (s *Server) volumesFor(needles []types.Needle) map[string]types.Volume {
	m := make(map[string]types.Volume, len(needles))
	if len(needles) == 0 {
		return m
	}
	ownerSet := make(map[string]struct{})
	fileSet := make(map[uint64]struct{})
	for _, n := range needles {
		ownerSet[strings.ToLower(n.Owner)] = struct{}{}
		fileSet[n.File] = struct{}{}
	}
	owners := make([]string, 0, len(ownerSet))
	for o := range ownerSet {
		owners = append(owners, o)
	}
	files := make([]uint64, 0, len(fileSet))
	for f := range fileSet {
		files = append(files, f)
	}
	var vols []types.Volume
	if err := s.gdb.Where("LOWER(owner) IN ? AND file IN ?", owners, files).Find(&vols).Error; err != nil {
		logger.Warnf("batch volumesFor failed: %v", err)
		return m
	}
	for _, v := range vols {
		m[volKey(v.Owner, v.File)] = v
	}
	return m
}

func (s *Server) listVolume(owner string, offset, limit int) ([]types.Volume, error) {
	var vols []types.Volume
	q := s.gdb
	if owner != "" {
		q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
	}
	result := q.Order("id desc").Limit(limit).Offset(offset).Find(&vols)
	if result.Error != nil {
		return nil, result.Error
	}

	return vols, nil
}

func (s *Server) listConversationDisplay(addr, bucket string, offset, limit int) ([]types.Conversation, error) {
	var conversations []types.Conversation
	q := s.gdb
	if addr != "" {
		q = q.Where("LOWER(owner) = ?", strings.ToLower(addr))
	}
	if bucket != "" {
		q = q.Where("bucket = ?", bucket)
	}
	result := q.Order("id desc").Limit(limit).Offset(offset).Find(&conversations)
	if result.Error != nil {
		return nil, result.Error
	}
	return conversations, nil
}

func (s *Server) getConversationDisplay(addr, bucket, name string) ([]types.Conversation, error) {
	var conversations []types.Conversation
	q := s.gdb
	if addr != "" {
		q = q.Where("LOWER(owner) = ?", strings.ToLower(addr))
	}
	if bucket != "" {
		q = q.Where("bucket = ?", bucket)
	}
	if name != "" {
		q = q.Where("name = ?", name)
	}
	result := q.Find(&conversations)
	if result.Error != nil {
		return nil, result.Error
	}
	return conversations, nil
}

func (s *Server) listNeedleDisplayByConversation(addr, bucket, conversation string, offset, limit int) ([]types.NeedleDisplay, error) {
	var needles []types.Needle
	var result *gorm.DB

	// 首先获取conversation_0的id
	var firstNeedle types.Needle
	firstQuery := s.gdb.Model(&types.Needle{}).Where("name = ?", conversation+"_0")
	if addr != "" {
		firstQuery = firstQuery.Where("LOWER(owner) = ?", strings.ToLower(addr))
	}
	if bucket != "" {
		firstQuery = firstQuery.Where("bucket = ?", bucket)
	}

	firstResult := firstQuery.First(&firstNeedle)

	query := s.gdb.Model(&types.Needle{})
	if addr != "" {
		query = query.Where("LOWER(owner) = ?", strings.ToLower(addr))
	}
	if bucket != "" {
		query = query.Where("bucket = ?", bucket)
	}

	// 如果找到了conversation_0记录，使用id进行优化查询
	if firstResult.Error == nil {
		query = query.Where("name like ? AND id >= ?",
			conversation+"_%", firstNeedle.ID)
	} else {
		// 如果没有找到，使用原来的查询方式
		query = query.Where("name like ? and created_at >= ?",
			conversation+"_%",
			time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC))
	}

	result = query.Order("id desc").Limit(limit).Offset(offset).Find(&needles)
	if result.Error != nil {
		return nil, result.Error
	}
	vmap := s.volumesFor(needles) // batch instead of getVolume per needle (N+1)
	res := make([]types.NeedleDisplay, 0, len(needles))
	for i := 0; i < len(needles); i++ {
		nd := types.NeedleDisplay{
			CreatedAt: needles[i].CreatedAt,
			Name:      needles[i].Name,
			Owner:     needles[i].Owner,
			Bucket:    needles[i].Bucket,
			File:      needles[i].File,
			Start:     needles[i].Start,
			Size:      needles[i].Size,
		}
		if v, ok := vmap[volKey(needles[i].Owner, needles[i].File)]; ok {
			nd.Piece = v.Piece
			nd.TxHash = v.TxHash
			nd.ChainType = v.ChainType
		}
		res = append(res, nd)
	}
	return res, nil
}

func (s *Server) listConversation(addr, bucket string, offset, limit int) ([]string, error) {
	/*
		var needles []types.Needle
		// create time is time.Time >= 2025-03-07,
		// name end with "_0"
		// addr may be empty
		// bucket may be empty
		var result *gorm.DB
		if bucket == "" {
			result = s.gdb.Model(&types.Needle{}).Where("owner = ? and created_at >= ? and name like ?", addr, time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC), "%_0").Find(&needles)
		} else {
			result = s.gdb.Model(&types.Needle{}).Where("owner = ? and bucket = ? and created_at >= ? and name like ?", addr, bucket, time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC), "%_0").Find(&needles)
		}
		if result.Error != nil {
			return nil, result.Error
		}
		// name is conversation_index, so we need to split it to conversation_id
		// reomve duplicate
		conversations := make([]string, 0, len(needles))
		cmap := make(map[string]struct{})
		for _, needle := range needles {
			if len(needle.Name) <= 2 {
				continue
			}
			parts := strings.Split(needle.Name, "_")
			if len(parts) != 2 {
				continue
			}
			conversationID := parts[0]
			if _, ok := cmap[conversationID]; !ok {
				conversations = append(conversations, conversationID)
				cmap[conversationID] = struct{}{}
			}
		}
	*/
	conversations, err := s.listConversationDisplay(addr, bucket, offset, limit)
	if err != nil {
		return nil, err
	}

	var res []string
	cmap := make(map[string]struct{})
	for _, conversation := range conversations {
		if _, ok := cmap[conversation.Name]; !ok {
			res = append(res, conversation.Name)
			cmap[conversation.Name] = struct{}{}
		}
	}
	return res, nil
}

func (s *Server) getConversation(ctx context.Context, conversation, addr, bucket string, offset, limit int) ([]string, error) {
	var gconversation types.Conversation
	cq := s.gdb.Where("name = ?", conversation)
	if addr != "" {
		cq = cq.Where("LOWER(owner) = ?", strings.ToLower(addr))
	}
	if bucket != "" {
		cq = cq.Where("bucket = ?", bucket)
	}
	res := cq.Find(&gconversation)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return []string{}, nil
	}

	var needles []types.Needle
	query := s.gdb.Model(&types.Needle{}).Where("name like ? and created_at >= ?",
		conversation+"_%",
		time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC))
	if addr != "" {
		query = query.Where("LOWER(owner) = ?", strings.ToLower(addr))
	}
	if bucket != "" {
		query = query.Where("bucket = ?", bucket)
	}

	result := query.Order("id asc").Limit(limit).Offset(offset).Find(&needles)
	if result.Error != nil {
		return nil, result.Error
	}
	// name is conversation_index, so we need to split it to conversation_id
	// reomve duplicate
	conversations := make([]string, 0, len(needles))
	for _, needle := range needles {
		var w bytes.Buffer
		_, err := s.download(ctx, needle.Name, addr, &w)
		if err != nil {
			continue
		}
		conversations = append(conversations, w.String())
	}
	return conversations, nil
}

// periodicCheckpoint runs periodic WAL checkpoints to ensure data persistence
func (s *Server) periodicCheckpoint() {
	ticker := time.NewTicker(5 * time.Minute) // Checkpoint every 5 minutes to reduce lock contention
	defer ticker.Stop()

	for {
		select {
		case <-s.checkpointStop:
			// Perform final checkpoint before stopping
			if s.gdb != nil {
				logger.Info("performing final database checkpoint...")
				if err := s.gdb.Exec("PRAGMA wal_checkpoint(FULL);").Error; err != nil {
					logger.Errorf("final WAL checkpoint failed: %v", err)
				}
			}
			return
		case <-ticker.C:
			if s.gdb != nil {
				// Execute WAL checkpoint to persist data (use RESTART for better performance)
				if err := s.gdb.Exec("PRAGMA wal_checkpoint(RESTART);").Error; err != nil {
					logger.Debugf("WAL checkpoint failed: %v", err)
				}
			}
		}
	}
}
