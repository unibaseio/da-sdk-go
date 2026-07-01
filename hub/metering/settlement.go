package metering

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"gorm.io/gorm"
)

const (
	modeOffchain = "offchain"
	modeERC8183  = "erc8183"

	settlementPending    = "pending"
	settlementSubmitting = "submitting"
	settlementConfirmed  = "confirmed"
	settlementFailed     = "failed"

	settleStatusNoop = "noop"
)

// SettleTx holds the on-chain transaction hashes produced by an ERC-8183
// settlement. Empty for off-chain settlements.
type SettleTx struct {
	Transfer  string `json:"transfer,omitempty"`
	CreateJob string `json:"create_job,omitempty"`
	Fund      string `json:"fund,omitempty"`
	Submit    string `json:"submit,omitempty"`
}

// SettleResponse is returned by POST /api/metering/settle.
type SettleResponse struct {
	Owner            string    `json:"owner"`
	SettledAmountWei string    `json:"settled_amount_wei"`
	Mode             string    `json:"mode"`
	SettlementID     uint      `json:"settlement_id,omitempty"`
	Status           string    `json:"status"`
	JobID            string    `json:"job_id,omitempty"`
	ReportHash       string    `json:"report_hash,omitempty"`
	Tx               *SettleTx `json:"tx,omitempty"`
}

// Settle clears an owner's unsettled fee using the configured settlement mode.
func (m *Manager) Settle(owner string) (*SettleResponse, error) {
	if m.cfg.SettlementMode == modeERC8183 {
		return m.SettleERC8183(owner)
	}
	return m.SettleOffchain(owner)
}

// unsettledForOwner loads unsettled events (oldest first), returning the events,
// their total fee, and the [from,to] event id range. Empty set => amount 0.
func unsettledForOwner(tx *gorm.DB, owner string) (events []types.MeterEvent, amount *big.Int, fromID, toID uint, err error) {
	if err = tx.Where("owner = ? AND status = ?", owner, eventUnsettled).
		Order("id asc").Find(&events).Error; err != nil {
		return nil, nil, 0, 0, err
	}
	amount = big.NewInt(0)
	for _, e := range events {
		amount.Add(amount, parseWei(e.FeeWei))
	}
	if len(events) > 0 {
		fromID = events[0].ID
		toID = events[len(events)-1].ID
	}
	return events, amount, fromID, toID, nil
}

// clearDebt marks the unsettled events (id in [fromID,toID]) as settled under
// settlementID and subtracts amount from the account's unsettled fee. It must
// run inside a transaction while holding the owner lock. LastSettledAt is set.
func clearDebt(tx *gorm.DB, owner string, amount *big.Int, fromID, toID, settlementID uint) error {
	if err := tx.Model(&types.MeterEvent{}).
		Where("owner = ? AND status = ? AND id <= ?", owner, eventUnsettled, toID).
		Updates(map[string]interface{}{"status": eventSettled, "settlement_id": settlementID}).Error; err != nil {
		return err
	}

	var acct types.MeterAccount
	if err := tx.Where("owner = ?", owner).First(&acct).Error; err != nil {
		return err
	}
	remaining := new(big.Int).Sub(parseWei(acct.UnsettledFeeWei), amount)
	if remaining.Sign() < 0 {
		remaining = big.NewInt(0)
	}
	now := time.Now()
	return tx.Model(&types.MeterAccount{}).Where("owner = ?", owner).Updates(map[string]interface{}{
		"unsettled_fee_wei": remaining.String(),
		"last_settled_at":   &now,
	}).Error
}

// SettleOffchain clears the owner's unsettled fee in a single DB transaction and
// marks the covered events settled. No chain interaction. Idempotent: a second
// call with nothing unsettled is a no-op.
func (m *Manager) SettleOffchain(owner string) (*SettleResponse, error) {
	owner = strings.ToLower(owner)

	unlock := m.accountLocks.lock(owner)
	defer unlock()

	resp := &SettleResponse{Owner: owner, Mode: modeOffchain}
	err := m.db.Transaction(func(tx *gorm.DB) error {
		events, amount, fromID, toID, err := unsettledForOwner(tx, owner)
		if err != nil {
			return err
		}
		if len(events) == 0 {
			resp.Status = settleStatusNoop
			resp.SettledAmountWei = "0"
			return nil
		}

		settlement := types.MeterSettlement{
			Owner:       owner,
			AmountWei:   amount.String(),
			FromEventID: fromID,
			ToEventID:   toID,
			Status:      settlementConfirmed,
			Mode:        modeOffchain,
		}
		if err := tx.Create(&settlement).Error; err != nil {
			return err
		}
		if err := clearDebt(tx, owner, amount, fromID, toID, settlement.ID); err != nil {
			return err
		}

		resp.SettledAmountWei = amount.String()
		resp.SettlementID = settlement.ID
		resp.Status = settlementConfirmed
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SettleERC8183 settles an owner's unsettled fee through the ERC-8183 escrow.
// Critical invariant: debt is cleared ONLY after all chain calls succeed.
//
//	Phase A (locked, DB tx): snapshot unsettled events, create a pending
//	  settlement, and mark those events "settling" so they are reserved.
//	Phase B (no lock): run the chain sequence
//	  transferFrom -> createJob -> setBudget -> fund -> submit.
//	Phase C (locked, DB tx):
//	  success -> mark reserved events settled, subtract amount from the account,
//	    settlement confirmed with tx hashes + job id + report hash.
//	  failure -> revert reserved events to unsettled (debt intact, retryable),
//	    settlement failed with the error and any tx hashes already produced.
func (m *Manager) SettleERC8183(owner string) (*SettleResponse, error) {
	owner = strings.ToLower(owner)
	resp := &SettleResponse{Owner: owner, Mode: modeERC8183, Tx: &SettleTx{}}

	var settlementID uint
	amount := big.NewInt(0)
	var fromID, toID uint
	var report SettlementReport
	empty := false

	// Phase A: reserve events under a pending settlement.
	unlock := m.accountLocks.lock(owner)
	err := m.db.Transaction(func(tx *gorm.DB) error {
		events, amt, f, tt, err := unsettledForOwner(tx, owner)
		if err != nil {
			return err
		}
		if len(events) == 0 {
			empty = true
			return nil
		}
		amount, fromID, toID = amt, f, tt

		var writes, reads, bytesWritten uint64
		for _, e := range events {
			if e.Operation == opWrite {
				writes++
				bytesWritten += e.Bytes
			} else {
				reads++
			}
		}

		s := types.MeterSettlement{
			Owner:       owner,
			AmountWei:   amt.String(),
			FromEventID: f,
			ToEventID:   tt,
			Status:      settlementPending,
			Mode:        modeERC8183,
		}
		if err := tx.Create(&s).Error; err != nil {
			return err
		}
		settlementID = s.ID
		report = SettlementReport{
			Type:         ReportType,
			Owner:        owner,
			AmountWei:    amt.String(),
			FromEventID:  f,
			ToEventID:    tt,
			Writes:       writes,
			Reads:        reads,
			BytesWritten: bytesWritten,
			Timestamp:    time.Now().Unix(),
		}
		return tx.Model(&types.MeterEvent{}).
			Where("owner = ? AND status = ? AND id <= ?", owner, eventUnsettled, tt).
			Update("status", eventSettling).Error
	})
	unlock()
	if err != nil {
		return nil, err
	}
	if empty {
		resp.Status = settleStatusNoop
		resp.SettledAmountWei = "0"
		return resp, nil
	}

	// Phase B: chain calls (no lock held; may take minutes).
	chainErr := m.runERC8183Settlement(owner, amount, report, resp)

	// Phase C: finalize.
	unlock = m.accountLocks.lock(owner)
	defer unlock()

	if chainErr != nil {
		// Revert reservation; debt stays intact and retryable.
		ferr := m.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&types.MeterEvent{}).
				Where("owner = ? AND status = ? AND id >= ? AND id <= ?", owner, eventSettling, fromID, toID).
				Update("status", eventUnsettled).Error; err != nil {
				return err
			}
			return tx.Model(&types.MeterSettlement{}).Where("id = ?", settlementID).
				Updates(settlementTxFields(resp, settlementFailed, chainErr.Error())).Error
		})
		if ferr != nil {
			return nil, fmt.Errorf("chain settle failed (%v); finalize also failed: %w", chainErr, ferr)
		}
		resp.Status = settlementFailed
		return resp, chainErr
	}

	serr := m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&types.MeterEvent{}).
			Where("owner = ? AND status = ? AND id <= ?", owner, eventSettling, toID).
			Updates(map[string]interface{}{"status": eventSettled, "settlement_id": settlementID}).Error; err != nil {
			return err
		}

		var acct types.MeterAccount
		if err := tx.Where("owner = ?", owner).First(&acct).Error; err != nil {
			return err
		}
		remaining := new(big.Int).Sub(parseWei(acct.UnsettledFeeWei), amount)
		if remaining.Sign() < 0 {
			remaining = big.NewInt(0)
		}
		now := time.Now()
		if err := tx.Model(&types.MeterAccount{}).Where("owner = ?", owner).Updates(map[string]interface{}{
			"unsettled_fee_wei": remaining.String(),
			"last_settled_at":   &now,
		}).Error; err != nil {
			return err
		}
		return tx.Model(&types.MeterSettlement{}).Where("id = ?", settlementID).
			Updates(settlementTxFields(resp, settlementConfirmed, "")).Error
	})
	if serr != nil {
		return nil, serr
	}

	resp.Status = settlementConfirmed
	resp.SettledAmountWei = amount.String()
	resp.SettlementID = settlementID
	return resp, nil
}

// settlementTxFields builds the settlement update map from the response's tx
// hashes, job id, report hash, plus the final status and error string.
func settlementTxFields(resp *SettleResponse, status, errStr string) map[string]interface{} {
	return map[string]interface{}{
		"status":        status,
		"error":         errStr,
		"transfer_tx":   resp.Tx.Transfer,
		"create_job_tx": resp.Tx.CreateJob,
		"fund_tx":       resp.Tx.Fund,
		"submit_tx":     resp.Tx.Submit,
		"job_id":        resp.JobID,
		"report_hash":   resp.ReportHash,
	}
}

// runERC8183Settlement executes the on-chain settlement sequence, populating
// resp.Tx / JobID / ReportHash as each step completes so partial progress is
// recorded even on failure.
func (m *Manager) runERC8183Settlement(owner string, amount *big.Int, report SettlementReport, resp *SettleResponse) error {
	if m.erc20 == nil || m.erc8183 == nil {
		return fmt.Errorf("chain clients not configured for erc8183 settlement")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// 1. Pull funds from the user to the provider.
	th, err := m.erc20.TransferFrom(ctx, common.HexToAddress(owner), m.provider, amount)
	if err != nil {
		return fmt.Errorf("transferFrom: %w", err)
	}
	resp.Tx.Transfer = th.Hex()

	// 2. Create the escrow job.
	expiredAt := big.NewInt(time.Now().Add(time.Hour).Unix())
	desc := fmt.Sprintf("da-hub-metering owner=%s amount=%s", owner, amount.String())
	jobID, cjTx, err := m.erc8183.CreateJob(ctx, expiredAt, desc, big.NewInt(0))
	if err != nil {
		return fmt.Errorf("createJob: %w", err)
	}
	resp.Tx.CreateJob = cjTx.Hex()
	resp.JobID = jobID.String()

	// 3. Set the budget.
	if _, err := m.erc8183.SetBudget(ctx, jobID, amount); err != nil {
		return fmt.Errorf("setBudget: %w", err)
	}

	// 4. Fund the job.
	fTx, err := m.erc8183.Fund(ctx, jobID, amount)
	if err != nil {
		return fmt.Errorf("fund: %w", err)
	}
	resp.Tx.Fund = fTx.Hex()

	// 5. Submit the report hash as the deliverable.
	deliverable, hexHash, err := report.Hash()
	if err != nil {
		return fmt.Errorf("report hash: %w", err)
	}
	resp.ReportHash = hexHash
	sTx, err := m.erc8183.Submit(ctx, jobID, deliverable)
	if err != nil {
		return fmt.Errorf("submit: %w", err)
	}
	resp.Tx.Submit = sTx.Hex()

	return nil
}
