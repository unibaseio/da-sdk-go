package metering

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// SignerFunc extracts the authenticated, lowercased signer address from the gin
// context (or "" if unauthenticated). It is injected by the hub package so the
// metering package does not import hub (which would be an import cycle).
type SignerFunc func(c *gin.Context) string

// PricingResponse is returned by GET /api/metering/pricing.
type PricingResponse struct {
	Enabled           bool   `json:"enabled"`
	ChargeWrites      bool   `json:"charge_writes"`
	ChargeReads       bool   `json:"charge_reads"`
	WriteBaseWei      string `json:"write_base_wei"`
	WritePerKBWei     string `json:"write_per_kb_wei"`
	ReadPerRequestWei string `json:"read_per_request_wei"`
}

// RegisterRoutes wires metering endpoints. Public read-only routes are added to
// pub; authenticated routes (added in later phases) go on authed and use signer
// to enforce signer == owner. Passing nil for signer is allowed while only
// public routes exist.
func (m *Manager) RegisterRoutes(pub, authed *gin.RouterGroup, signer SignerFunc) {
	pub.GET("/metering/pricing", m.handlePricing)
	pub.GET("/metering/usage", m.handleUsage)
}

func (m *Manager) handlePricing(c *gin.Context) {
	p := m.Pricing()
	c.JSON(http.StatusOK, PricingResponse{
		Enabled:           m.cfg.Enabled,
		ChargeWrites:      m.cfg.ChargeWrites,
		ChargeReads:       m.cfg.ChargeReads,
		WriteBaseWei:      p.WriteBaseWei.String(),
		WritePerKBWei:     p.WritePerKBWei.String(),
		ReadPerRequestWei: p.ReadPerRequestWei.String(),
	})
}

func (m *Manager) handleUsage(c *gin.Context) {
	owner := c.Query("owner")
	if owner == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner is required"})
		return
	}
	if !common.IsHexAddress(owner) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner must be a 0x-prefixed Ethereum address"})
		return
	}
	usage, err := m.GetUsage(owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usage)
}
