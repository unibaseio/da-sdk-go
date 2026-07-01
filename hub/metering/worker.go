package metering

import (
	"math/big"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

// Worker periodically settles accounts whose unsettled fee has reached the
// configured threshold. It runs only when HUB_METERING_AUTO_SETTLE=true.
//
// First version assumes a single hub instance. For multi-replica deployment,
// add DB locking or leader election before enabling in production.
type Worker struct {
	m         *Manager
	interval  time.Duration
	threshold *big.Int
	stop      chan struct{}
	done      chan struct{}
}

// AutoSettleEnabled reports whether the auto-settlement worker should run.
func (m *Manager) AutoSettleEnabled() bool {
	return m.Enabled() && m.cfg.AutoSettle
}

// NewWorker builds the worker from config. Interval defaults to 300s when unset.
func (m *Manager) NewWorker() *Worker {
	interval := time.Duration(m.cfg.SettleIntervalSec) * time.Second
	if interval <= 0 {
		interval = 300 * time.Second
	}
	return &Worker{
		m:         m,
		interval:  interval,
		threshold: new(big.Int).Set(m.cfg.SettleThresholdWei),
		stop:      make(chan struct{}),
		done:      make(chan struct{}),
	}
}

// Start launches the background loop. Non-blocking.
func (w *Worker) Start() {
	go func() {
		defer close(w.done)
		ticker := time.NewTicker(w.interval)
		defer ticker.Stop()
		logger.Infof("metering worker started: interval=%s threshold=%s", w.interval, w.threshold.String())
		for {
			select {
			case <-w.stop:
				return
			case <-ticker.C:
				w.runOnce()
			}
		}
	}()
}

// Stop signals the loop and waits for it to exit. Safe to call once.
func (w *Worker) Stop() {
	close(w.stop)
	<-w.done
	logger.Info("metering worker stopped")
}

// runOnce scans eligible accounts and settles those at/above the threshold. A
// failure or panic settling one account is logged and never aborts the scan.
func (w *Worker) runOnce() {
	var accts []types.MeterAccount
	// Non-numeric SQL comparison of decimal strings is unsafe, so only exclude
	// obviously-zero rows here and compare against the threshold with big.Int.
	err := w.m.db.
		Where("enabled = ? AND unsettled_fee_wei != '' AND unsettled_fee_wei != '0'", true).
		Find(&accts).Error
	if err != nil {
		logger.Warnf("metering worker: scan failed: %v", err)
		return
	}

	for _, a := range accts {
		unsettled := parseWei(a.UnsettledFeeWei)
		if unsettled.Cmp(w.threshold) < 0 {
			continue
		}
		w.settleOne(a.Owner)
	}
}

// settleOne settles a single owner, recovering from any panic so the scan
// continues and the hub never crashes on a settlement fault.
func (w *Worker) settleOne(owner string) {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("metering worker: settle panicked for %s: %v", owner, r)
		}
	}()
	resp, err := w.m.Settle(owner)
	if err != nil {
		logger.Warnf("metering worker: settle failed for %s: %v", owner, err)
		return
	}
	logger.Infof("metering worker: settled %s amount=%s status=%s", owner, resp.SettledAmountWei, resp.Status)
}
