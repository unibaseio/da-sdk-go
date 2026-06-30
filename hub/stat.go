package hub

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) addStat(g *gin.RouterGroup) {
	g.Group("/").POST("/stat", s.getStatByPost)
	g.Group("/").GET("/stat", s.getStatByGet)

	g.Group("/").POST("/memoryStat", s.listMemoryStatByPost)
	g.Group("/").GET("/memoryStat", s.listMemoryStatByGet)

	g.Group("/").GET("/memoryOverview", s.getMemoryOverview)
	g.Group("/").POST("/memoryOverview", s.getMemoryOverview)
}

// getMemoryOverview godoc
//
//	@Summary		Memory dashboard overview
//	@Description	Hub-wide totals: distinct addresses, wallets with memory, total memory entries, total size (GB).
//	@Tags			statistics
//	@Produce		json
//	@Success		200	{object}	types.MemoryOverview
//	@Failure		599	{object}	lerror.APIError
//	@Router			/api/memoryOverview [get]
func (s *Server) getMemoryOverview(c *gin.Context) {
	// served from the background-computed snapshot (no inline DB scan)
	c.JSON(http.StatusOK, s.memoryOverviewSnapshot())
}

// listMemoryStatByGet godoc
//
//	@Summary		Per-owner memory stats (paginated)
//	@Description	List each wallet's memory entry count and total size (GB), grouped by owner, ordered by size desc. Pass owner to filter to a single wallet (case-insensitive); result stays a list (0 or 1 item).
//	@Tags			statistics
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	false	"filter to a single wallet address (case-insensitive)"
//	@Param			offset	query		int		false	"pagination offset" default(0)
//	@Param			length	query		int		false	"page size (max 100)" default(32)
//	@Success		200		{object}	types.MemoryStatResult
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/memoryStat [get]
func (s *Server) listMemoryStatByGet(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	length, _ := strconv.Atoi(c.Query("length"))
	s.serveMemoryStat(c, c.Query("owner"), offset, length)
}

// listMemoryStatByPost godoc
//
//	@Summary		Per-owner memory stats (paginated, POST)
//	@Description	List each wallet's memory entry count and total size (GB), grouped by owner, ordered by size desc. Pass owner to filter to a single wallet (case-insensitive); result stays a list (0 or 1 item).
//	@Tags			statistics
//	@Accept			application/x-www-form-urlencoded
//	@Produce		json
//	@Param			owner	formData	string	false	"filter to a single wallet address (case-insensitive)"
//	@Param			offset	formData	int		false	"pagination offset" default(0)
//	@Param			length	formData	int		false	"page size (max 100)" default(32)
//	@Success		200		{object}	types.MemoryStatResult
//	@Failure		599		{object}	lerror.APIError
//	@Router			/api/memoryStat [post]
func (s *Server) listMemoryStatByPost(c *gin.Context) {
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	length, _ := strconv.Atoi(c.PostForm("length"))
	s.serveMemoryStat(c, c.PostForm("owner"), offset, length)
}

func (s *Server) serveMemoryStat(c *gin.Context, owner string, offset, length int) {
	if length <= 0 {
		length = 32
	}
	if length > 100 {
		length = 100
	}
	if offset < 0 {
		offset = 0
	}
	// served (and filtered) from the background-computed snapshot (no inline DB scan)
	c.JSON(http.StatusOK, s.memoryStatPage(owner, offset, length))
}

// @Summary Get statistics by POST
// @Description Get daily statistics for accounts, needles, and volumes using POST method
// @Tags statistics
// @Accept x-www-form-urlencoded
// @Produce json
// @Param count formData int false "Number of days to return (default: 30)"
// @Success 200 {array} types.Stat "Array of daily statistics"
// @Failure 400 {object} map[string]string "Invalid count parameter"
// @Router /api/stat [post]
func (s *Server) getStatByPost(c *gin.Context) {
	count := 30
	countStr := c.PostForm("count")
	countInt, err := strconv.Atoi(countStr)
	if err == nil && countInt > 0 {
		count = countInt
	}
	stats := s.statManager.GetStats(count)
	c.JSON(200, stats)
}

// @Summary Get statistics by GET
// @Description Get daily statistics for accounts, needles, and volumes using GET method
// @Tags statistics
// @Accept json
// @Produce json
// @Param count query int false "Number of days to return (default: 30)"
// @Success 200 {array} types.Stat "Array of daily statistics"
// @Failure 400 {object} map[string]string "Invalid count parameter"
// @Router /api/stat [get]
func (s *Server) getStatByGet(c *gin.Context) {
	count := 30
	countStr := c.Query("count")
	countInt, err := strconv.Atoi(countStr)
	if err == nil && countInt > 0 {
		count = countInt
	}
	stats := s.statManager.GetStats(count)
	c.JSON(200, stats)
}

// StatManager manages daily statistics for accounts, needles, and volumes
type StatManager struct {
	mu       sync.RWMutex
	db       *gorm.DB
	stats    map[string]*types.Stat // key: YYYY-MM-DD
	lastDay  time.Time
	stopChan chan struct{}

	// Track last processed IDs for incremental updates
	lastAccountID uint
	lastBucketID  uint
	lastNeedleID  uint
	lastVolumeID  uint
}

// NewStatManager creates a new StatManager instance
func NewStatManager(db *gorm.DB) *StatManager {

	return &StatManager{
		db:       db,
		stats:    make(map[string]*types.Stat),
		stopChan: make(chan struct{}),
	}
}

// loadStats loads statistics from database
func (sm *StatManager) loadStats() (int, error) {
	// find the latest 30 day records
	now := time.Now().UTC()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	recordCount := 0
	for i := 30; i >= 0; i-- {
		var record types.StatRecord
		daytime := now.AddDate(0, 0, -i)
		err := sm.db.Order("id desc").Where("day = ?", daytime).First(&record).Error
		if err != nil {
			fmt.Println("no record for day: ", daytime, err)
			continue
		}
		fmt.Println("found record at day: ", daytime, record)
		day := record.Day.Format("2006-01-02")
		sm.stats[day] = &types.Stat{
			Day:           record.Day,
			DailyAccounts: record.DailyAccounts,
			DailyBuckets:  record.DailyBuckets,
			DailyNeedles:  record.DailyNeedles,
			DailyVolumes:  record.DailyVolumes,
			TotalAccounts: record.TotalAccounts,
			TotalBuckets:  record.TotalBuckets,
			TotalNeedles:  record.TotalNeedles,
			TotalVolumes:  record.TotalVolumes,
		}
		recordCount++
		sm.lastDay = daytime
	}

	logger.Infof("found %d records", recordCount)

	if recordCount == 0 {
		return 0, nil
	}

	var record types.StatRecord
	err := sm.db.Order("id desc").Where("day = ?", sm.lastDay).First(&record).Error
	if err != nil {
		return 0, nil
	}

	sm.lastAccountID = record.LastAccountID
	sm.lastBucketID = record.LastBucketID
	sm.lastNeedleID = record.LastNeedleID
	sm.lastVolumeID = record.LastVolumeID

	for sm.lastDay.AddDate(0, 0, 1).Before(now) {
		sm.lastDay = sm.lastDay.AddDate(0, 0, 1)
		day := sm.lastDay.Format("2006-01-02")
		st := &types.Stat{
			Day:           sm.lastDay,
			TotalAccounts: record.TotalAccounts,
			TotalBuckets:  record.TotalBuckets,
			TotalNeedles:  record.TotalNeedles,
			TotalVolumes:  record.TotalVolumes,
		}
		sm.stats[day] = st
		recordCount++
		sm.saveStat(st)
	}

	return recordCount, nil
}

// saveStat saves a stat record to database
func (sm *StatManager) saveStat(stat *types.Stat) error {
	record := types.StatRecord{
		Day:           stat.Day,
		DailyAccounts: stat.DailyAccounts,
		DailyBuckets:  stat.DailyBuckets,
		DailyNeedles:  stat.DailyNeedles,
		DailyVolumes:  stat.DailyVolumes,
		TotalAccounts: stat.TotalAccounts,
		TotalBuckets:  stat.TotalBuckets,
		TotalNeedles:  stat.TotalNeedles,
		TotalVolumes:  stat.TotalVolumes,
		LastAccountID: sm.lastAccountID,
		LastBucketID:  sm.lastBucketID,
		LastNeedleID:  sm.lastNeedleID,
		LastVolumeID:  sm.lastVolumeID,
	}

	// if record for the day exists, update it
	sm.db.Order("id desc").Where("day = ?", record.Day).Assign(record).FirstOrCreate(&record)

	return nil
}

// Start initializes the StatManager with historical data and starts the background update routine
func (sm *StatManager) Start(ctx context.Context) error {
	// Load existing stats from database
	statcount, err := sm.loadStats()
	if err != nil {
		return err
	}

	initStat := os.Getenv("INIT_STAT")
	// If no stats exist, initialize with historical data
	if statcount == 0 || initStat != "" {
		logger.Warnf("no enough stats found, initializing with historical data")
		sm.stats = make(map[string]*types.Stat)
		now := time.Now().UTC()
		now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

		startDay := now.AddDate(0, 0, -31)
		sm.updateStat(startDay)

		for i := 30; i >= 0; i-- {
			day := now.AddDate(0, 0, -i)
			sm.updateDailyStat(day)
		}
		sm.lastDay = now
	}

	// Start background update routine
	go sm.run(ctx)
	return nil
}

// Stop stops the background update routine
func (sm *StatManager) Stop() {
	logger.Info("stopping statistics manager...")

	// Signal the background routine to stop
	close(sm.stopChan)

	// Save current statistics before shutdown
	sm.mu.Lock()
	defer sm.mu.Unlock()

	now := time.Now().UTC()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	// Update and save current day's statistics
	if now.Day() != sm.lastDay.Day() {
		sm.updateDailyStat(sm.lastDay)
		sm.lastDay = now
	}
	// Update current day's data
	sm.updateDailyStat(now)

	logger.Info("statistics manager stopped and data saved")
}

// run executes the background update routine
func (sm *StatManager) run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-sm.stopChan:
			return
		case <-ticker.C:
			now := time.Now().UTC()
			// Align to day boundary
			now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

			// Update previous day's data if date changed
			if now.Day() != sm.lastDay.Day() {
				sm.updateDailyStat(sm.lastDay)
				sm.lastDay = now
			}
			// Update current day's data
			sm.updateDailyStat(now)
		}
	}
}

// updateStat updates statistics for before a day
func (sm *StatManager) updateStat(t time.Time) {
	// Ensure time is aligned to day boundary
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	day := t.Format("2006-01-02")
	nextDay := t.Add(24 * time.Hour)

	// Get total counts up to the specified day and update last IDs
	var totalAccounts []types.Account
	sm.db.Model(&types.Account{}).Where("created_at < ?", nextDay).Find(&totalAccounts)
	totalAccountsCount := int64(len(totalAccounts))
	if len(totalAccounts) > 0 {
		sm.lastAccountID = totalAccounts[len(totalAccounts)-1].ID
	}

	var totalBuckets []types.Bucket
	sm.db.Model(&types.Bucket{}).Where("created_at < ?", nextDay).Find(&totalBuckets)
	totalBucketsCount := int64(len(totalBuckets))
	if len(totalBuckets) > 0 {
		sm.lastBucketID = totalBuckets[len(totalBuckets)-1].ID
	}

	var totalNeedles []types.Needle
	sm.db.Model(&types.Needle{}).Where("created_at < ?", nextDay).Find(&totalNeedles)
	totalNeedlesCount := int64(len(totalNeedles))
	if len(totalNeedles) > 0 {
		sm.lastNeedleID = totalNeedles[len(totalNeedles)-1].ID
	}

	var totalVolumes []types.Volume
	sm.db.Model(&types.Volume{}).Where("created_at < ?", nextDay).Find(&totalVolumes)
	totalVolumesCount := int64(len(totalVolumes))
	if len(totalVolumes) > 0 {
		sm.lastVolumeID = totalVolumes[len(totalVolumes)-1].ID
	}

	stat := &types.Stat{
		Day:           t,
		TotalAccounts: totalAccountsCount,
		TotalBuckets:  totalBucketsCount,
		TotalNeedles:  totalNeedlesCount,
		TotalVolumes:  totalVolumesCount,
	}

	sm.mu.Lock()
	sm.stats[day] = stat
	sm.mu.Unlock()
	logger.Info("stat: ", stat)
}

// updateDayStat updates statistics for a specific day
func (sm *StatManager) updateDailyStat(t time.Time) {
	// Ensure time is aligned to day boundary
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	day := t.Format("2006-01-02")
	nextDay := t.Add(24 * time.Hour)

	// Get daily counts using incremental updates with optimized queries
	var dailyAccounts []types.Account
	sm.db.Model(&types.Account{}).
		Where("id > ? AND created_at < ? AND created_at >= ?", sm.lastAccountID, nextDay, t).
		Order("id ASC").
		Find(&dailyAccounts)
	if len(dailyAccounts) > 0 {
		sm.lastAccountID = dailyAccounts[len(dailyAccounts)-1].ID
	}
	dailyAccountsCount := int64(len(dailyAccounts))

	var dailyBuckets []types.Bucket
	sm.db.Model(&types.Bucket{}).
		Where("id > ? AND created_at < ? AND created_at >= ?", sm.lastBucketID, nextDay, t).
		Order("id ASC").
		Find(&dailyBuckets)
	if len(dailyBuckets) > 0 {
		sm.lastBucketID = dailyBuckets[len(dailyBuckets)-1].ID
	}
	dailyBucketsCount := int64(len(dailyBuckets))

	var dailyNeedles []types.Needle
	sm.db.Model(&types.Needle{}).
		Where("id > ? AND created_at < ? AND created_at >= ?", sm.lastNeedleID, nextDay, t).
		Order("id ASC").
		Find(&dailyNeedles)
	if len(dailyNeedles) > 0 {
		sm.lastNeedleID = dailyNeedles[len(dailyNeedles)-1].ID
	}
	dailyNeedlesCount := int64(len(dailyNeedles))

	var dailyVolumes []types.Volume
	sm.db.Model(&types.Volume{}).
		Where("id > ? AND created_at < ? AND created_at >= ?", sm.lastVolumeID, nextDay, t).
		Order("id ASC").
		Find(&dailyVolumes)
	if len(dailyVolumes) > 0 {
		sm.lastVolumeID = dailyVolumes[len(dailyVolumes)-1].ID
	}
	dailyVolumesCount := int64(len(dailyVolumes))

	sm.mu.Lock()
	stat, ok := sm.stats[day]
	if !ok {
		stat = &types.Stat{
			Day: t,
		}
		sm.stats[day] = stat
	}
	// Accumulate daily incremental data
	stat.DailyAccounts = stat.DailyAccounts + dailyAccountsCount
	stat.DailyBuckets = stat.DailyBuckets + dailyBucketsCount
	stat.DailyNeedles = stat.DailyNeedles + dailyNeedlesCount
	stat.DailyVolumes = stat.DailyVolumes + dailyVolumesCount

	prevDayTime := t.AddDate(0, 0, -1)
	prevDay := prevDayTime.Format("2006-01-02")
	prevDayStat, ok := sm.stats[prevDay]
	if !ok {
		prevDayStat = &types.Stat{
			Day: prevDayTime,
		}
		sm.stats[prevDay] = prevDayStat
	}
	// Calculate totals using accumulated daily data
	stat.TotalAccounts = prevDayStat.TotalAccounts + stat.DailyAccounts
	stat.TotalBuckets = prevDayStat.TotalBuckets + stat.DailyBuckets
	stat.TotalNeedles = prevDayStat.TotalNeedles + stat.DailyNeedles
	stat.TotalVolumes = prevDayStat.TotalVolumes + stat.DailyVolumes

	sm.mu.Unlock()

	// Save the updated stat to database
	if err := sm.saveStat(stat); err != nil {
		logger.Error("Failed to save stat: ", err)
	}

	logger.Info("stat: ", stat)
}

// GetStats returns a copy of all statistics
func (sm *StatManager) GetStats(count int) []types.Stat {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// Return a copy to prevent external modification
	// sort the result by day desc
	result := make([]types.Stat, 0, count)
	now := time.Now().UTC()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	for i := 0; i < count; i++ {
		day := now.AddDate(0, 0, -i)
		stat, ok := sm.stats[day.Format("2006-01-02")]
		if !ok {
			continue
		}
		result = append(result, *stat)
	}

	return result
}
