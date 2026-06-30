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

	StoreNodes     int64 // total store nodes until this day
	StreamNodes    int64 // total stream nodes until this day
	ValidatorNodes int64 // total validator nodes until this day
}
