package types

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name string
}

type Bucket struct {
	gorm.Model
	Name  string
	Owner string
}

type BucketDisplay struct {
	Bucket
	Description string
	State       string
	Transport   string
	Type        string
	Last        time.Time
}

type Needle struct {
	gorm.Model
	Name   string
	Owner  string
	Bucket string
	File   uint64
	Start  uint64
	Size   uint64
	Commit string
}

type Conversation struct {
	gorm.Model
	Name   string
	Owner  string
	Bucket string
	Count  uint64
}

type Volume struct {
	gorm.Model
	ChainType string
	Owner     string
	File      uint64
	Piece     string
	TxHash    string
}

type StatRecord struct {
	gorm.Model
	Day           time.Time `gorm:"primaryKey"`
	DailyAccounts int64
	DailyBuckets  int64
	DailyNeedles  int64
	DailyVolumes  int64
	TotalAccounts int64
	TotalBuckets  int64
	TotalNeedles  int64
	TotalVolumes  int64
	LastAccountID uint
	LastBucketID  uint
	LastNeedleID  uint
	LastVolumeID  uint
}

type NeedleDisplay struct {
	CreatedAt time.Time
	Name      string
	Owner     string
	Bucket    string
	File      uint64
	Start     uint64
	Size      uint64
	Piece     string
	ChainType string
	TxHash    string
}

// MemoryStat is one owner's aggregate memory usage.
type MemoryStat struct {
	Owner string  `json:"owner"` // wallet address (lowercased)
	Count int64   `json:"count"` // number of memory entries (needles)
	Bytes int64   `json:"bytes"` // total stored bytes
	GB    float64 `json:"gb"`    // bytes / 1e9, convenience
}

// MemoryStatResult is one page of per-owner memory stats.
type MemoryStatResult struct {
	Total  int64        `json:"total"` // total distinct owners (for pagination)
	Offset int          `json:"offset"`
	Length int          `json:"length"`
	Items  []MemoryStat `json:"items"`
}

// MemoryOverview is the dashboard summary (the four overview cards).
type MemoryOverview struct {
	TotalAddresses    int64   `json:"totalAddresses"`    // distinct addresses known to the hub (accounts)
	WalletsWithMemory int64   `json:"walletsWithMemory"` // owners with >=1 memory entry
	MemoryCount       int64   `json:"memoryCount"`       // total memory entries (needles)
	MemoryBytes       int64   `json:"memoryBytes"`       // total stored bytes
	MemoryGB          float64 `json:"memoryGB"`          // memoryBytes / 1e9
}

type Stat struct {
	Day           time.Time
	DailyAccounts int64 // new created accounts at this day
	DailyBuckets  int64 // new created buckets at this day
	DailyNeedles  int64 // new created needles at this day
	DailyVolumes  int64 // new created volumes at this day
	TotalAccounts int64 // total accounts until this day
	TotalBuckets  int64 // total buckets until this day
	TotalNeedles  int64 // total needles until this day
	TotalVolumes  int64 // total volumes until this day
}

type FileStat struct {
	Day        time.Time
	DailyEdges int64 // new created edges at this day
	DailyFiles int64 // new created files at this day
	DailySize  int64 // new created size at this day
	TotalEdges int64 // total edges until this day
	TotalFiles int64 // total files until this day
	TotalSize  int64 // total size until this day
}
