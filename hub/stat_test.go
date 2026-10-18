package hub

import (
	"context"
	"testing"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestStatManagerStopReturns: Stop used to hold sm.mu while calling
// updateDailyStat, which takes sm.mu again — a self-deadlock that hung every
// graceful shutdown. Stop must return, and must do so after run has exited.
func TestStatManagerStopReturns(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:"+t.Name()+"?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1)
	if err := db.AutoMigrate(&types.Account{}, &types.Bucket{}, &types.Needle{}, &types.Volume{}, &types.StatRecord{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}

	sm := NewStatManager(db)
	if err := sm.Start(context.Background()); err != nil {
		t.Fatalf("start: %v", err)
	}

	stopped := make(chan struct{})
	go func() { sm.Stop(); close(stopped) }()
	select {
	case <-stopped:
	case <-time.After(10 * time.Second):
		t.Fatal("Stop did not return: StatManager deadlocked on shutdown")
	}
	select {
	case <-sm.done:
	default:
		t.Fatal("Stop returned before the run loop exited")
	}
}
