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
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func (s *Server) loadGORM() {
	gpath := filepath.Join(s.rp.Path(), "gorm")

	os.MkdirAll(gpath, os.ModePerm)
	gpath = filepath.Join(gpath, "gorm.db")
	db, err := gorm.Open(sqlite.Open(gpath), &gorm.Config{
		Logger:                 glogger.Default.LogMode(glogger.Silent),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Enable WAL mode for better concurrency and data safety
	_ = db.Exec("PRAGMA journal_mode=WAL;")
	_ = db.Exec("PRAGMA synchronous = NORMAL;") // NORMAL provides good balance of safety and performance
	_ = db.Exec("PRAGMA cache_size = -64000;")  // 64MB cache (reduced for more frequent writes)
	_ = db.Exec("PRAGMA temp_store = MEMORY;")
	_ = db.Exec("PRAGMA mmap_size = 4000000000;")    // 4GB mmap
	_ = db.Exec("PRAGMA wal_autocheckpoint = 1000;") // Less frequent checkpoints to reduce lock contention
	_ = db.Exec("PRAGMA busy_timeout = 60000;")      // Increase timeout to 60 seconds for better handling of concurrent operations

	sqldb, err := db.DB()
	if err != nil {
		panic("failed to get database")
	}

	sqldb.SetMaxIdleConns(5)                   // Reduce idle connections to minimize lock contention
	sqldb.SetMaxOpenConns(20)                  // Reduce max connections to prevent overwhelming SQLite
	sqldb.SetConnMaxLifetime(15 * time.Minute) // Shorter connection lifetime for better resource management

	// Auto migrate tables
	db.AutoMigrate(&types.Account{})
	db.AutoMigrate(&types.Bucket{})
	db.AutoMigrate(&types.Needle{})
	db.AutoMigrate(&types.Volume{})
	db.AutoMigrate(&types.StatRecord{})
	db.AutoMigrate(&types.Conversation{})

	// Add indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_owner ON needles(owner);")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_bucket ON needles(bucket);")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_name ON needles(name);")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_owner_name ON needles(owner, name);")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_owner_bucket ON needles(owner, bucket);")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_owner_bucket_name ON needles(owner, bucket, name);")
	// listNeedle filters by LOWER(owner) and orders by id desc. Without a
	// (LOWER(owner), id) index that query SEARCHes by owner but then builds a
	// TEMP B-TREE to sort all matched rows by id — catastrophic for a wallet
	// with millions of needles (loads them all to return 32). This index serves
	// WHERE LOWER(owner)=? ORDER BY id via a reverse index walk (no temp sort).
	// NOTE: on a large existing DB the first startup builds this once (minutes,
	// write-locking) before serving — expected, do not restart mid-build.
	db.Exec("CREATE INDEX IF NOT EXISTS idx_needles_lower_owner_id ON needles(LOWER(owner), id);")

	s.gdb = db

	// Start periodic checkpoint routine for data safety
	go s.periodicCheckpoint()

	// iterate all needles to update bucket
	ni := os.Getenv("NEED_INIT")
	if ni != "" {
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
					s.addBucket(needle.Owner, bucketName)
					// update needle
					db.Model(&needle).Update("bucket", bucketName)
				} else {
					bucketName := needle.Owner
					s.addBucket(needle.Owner, bucketName)
					// update needle
					db.Model(&needle).Update("bucket", bucketName)
				}
			}
		}
	}

}

func (s *Server) addAccount(owner string) {
	var account types.Account
	result := s.gdb.First(&account, "name = ?", owner)
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
func (s *Server) addBucket(owner, bucket string) error {
	var gbucket types.Bucket
	result := s.gdb.First(&gbucket, "name = ? ", bucket)
	if result.RowsAffected > 0 {
		if !strings.EqualFold(gbucket.Owner, owner) {
			logger.Infof("bucket: %s is owned by %s", bucket, gbucket.Owner)
			return fmt.Errorf("bucket: %s is owned by %s", bucket, gbucket.Owner)
		}
		logger.Info("already has bucket: ", bucket)
		return nil
	}

	s.gdb.Create(&types.Bucket{
		Name:  bucket,
		Owner: owner,
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

func (s *Server) listBucket(owner string, offset, limit int) ([]types.BucketDisplay, error) {
	var buckets []types.Bucket
	q := s.gdb
	if owner != "" {
		q = q.Where("LOWER(owner) = ?", strings.ToLower(owner))
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

func (s *Server) addNeedle(owner, bucket, name string, findex uint64, start, length uint64) {
	s.gdb.Create(&types.Needle{
		Owner:  owner,
		Bucket: bucket,
		Name:   name,
		File:   findex,
		Start:  start,
		Size:   length,
	})

	if strings.HasSuffix(name, "_0") {
		connName := strings.TrimSuffix(name, "_0")
		// check if conversation already exists
		var conversation types.Conversation
		result := s.gdb.Where(&types.Conversation{Name: connName, Owner: owner, Bucket: bucket}).First(&conversation)
		if result.RowsAffected == 0 {
			s.gdb.Save(&types.Conversation{
				Name:   connName,
				Owner:  owner,
				Bucket: bucket,
			})
		}
	}
	logger.Info("create needle: ", owner)
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
		vol, err := s.getVolume(needles[i].Owner, needles[i].File)
		if err == nil && len(vol) > 0 {
			nd.Piece = vol[0].Piece
			nd.TxHash = vol[0].TxHash
			nd.ChainType = vol[0].ChainType
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
