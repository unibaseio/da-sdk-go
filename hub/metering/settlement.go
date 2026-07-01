package metering

import (
	"math/big"
	"strings"
	"time"

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
// The erc8183 mode is wired in a later phase; until then all settlement is
// off-chain.
func (m *Manager) Settle(owner string) (*SettleResponse, error) {
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
