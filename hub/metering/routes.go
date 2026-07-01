package metering

import (
	"net/http"
	"strconv"
	"strings"

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
	pub.GET("/metering/can-write", m.handleCanWrite)

	// Authenticated routes enforce signer == owner via the injected signer.
	if authed != nil && signer != nil {
		authed.POST("/metering/settle", m.handleSettle(signer))
	}
}

// resolveAuthedOwner returns the canonical lowercase owner for an authenticated
// metering action. The owner may be supplied in the form body or query; if
// omitted it defaults to the signer. A supplied owner must be a valid address
// and match the signer. On failure it writes the response and returns ok=false.
func resolveAuthedOwner(c *gin.Context, signer SignerFunc) (string, bool) {
	s := signer(c)
	if s == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no signer in context"})
		return "", false
	}
	owner := c.PostForm("owner")
	if owner == "" {
		owner = c.Query("owner")
	}
	if owner == "" {
		return strings.ToLower(s), true
	}
	if !common.IsHexAddress(owner) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner must be a 0x-prefixed Ethereum address"})
		return "", false
	}
	if !strings.EqualFold(owner, s) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "owner does not match signer"})
		return "", false
	}
	return strings.ToLower(owner), true
}

func (m *Manager) handleSettle(signer SignerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		owner, ok := resolveAuthedOwner(c, signer)
		if !ok {
			return
		}
		resp, err := m.Settle(owner)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	}
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

// handleCanWrite is a read-only preflight query. It always returns 200 with the
// decision body (the "allowed" flag conveys the result); the 402 status is
// reserved for real upload attempts.
func (m *Manager) handleCanWrite(c *gin.Context) {
	owner := c.Query("owner")
	if owner == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner is required"})
		return
	}
	if !common.IsHexAddress(owner) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner must be a 0x-prefixed Ethereum address"})
		return
	}
	var bytes uint64
	if s := c.Query("bytes"); s != "" {
		n, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bytes must be a non-negative integer"})
			return
		}
		bytes = n
	}
	res, err := m.CanWrite(owner, bytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
