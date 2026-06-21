package hub

import (
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	com "github.com/unibaseio/da-sdk-go/contract/common"
	contract "github.com/unibaseio/da-sdk-go/contract/v2"
	lerror "github.com/unibaseio/da-sdk-go/lib/error"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
)

func (s *Server) addSeal(g *gin.RouterGroup) {
	g.Group("/").POST("/seal", s.seal)
}

// chainManager lazily builds (and caches) the chain client used by /api/seal to
// sign AddPiece with the hub's own key (v1, hub-paid registration).
func (s *Server) chainManager() (*contract.ContractManage, error) {
	s.cmMu.Lock()
	defer s.cmMu.Unlock()
	if s.cm != nil {
		return s.cm, nil
	}
	sk := s.rp.Key().Export().PrivateKey
	cm, err := contract.NewContractManage(sk, s.rp.Config().Chain.Type)
	if err != nil {
		return nil, err
	}
	s.cm = cm
	return cm, nil
}

// seal implements the membase × DA integration endpoint (DA_SEAL_ENDPOINT_SPEC):
// a client POSTs one already-encrypted segment blob; the hub erasure-encodes +
// stages it (sdk.Upload) and either registers it on-chain itself (v1,
// register=hub, default — client zero-gas/zero-key) or returns the fields the
// client needs to self-sign AddPiece (v2, register=client).
func (s *Server) seal(c *gin.Context) {
	owner := c.PostForm("owner")
	if !RequireOwnerMatch(c, owner) {
		return
	}
	register := c.DefaultPostForm("register", "hub")

	rsn, err1 := strconv.Atoi(c.PostForm("rsn"))
	rsk, err2 := strconv.Atoi(c.PostForm("rsk"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, lerror.ToAPIError("hub", fmt.Errorf("invalid rsn/rsk")))
		return
	}
	policy := types.Policy{N: uint8(rsn), K: uint8(rsk)}
	if err := policy.Check(); err != nil {
		c.JSON(http.StatusBadRequest, lerror.ToAPIError("hub", err))
		return
	}

	fh, err := c.FormFile("file")
	if err != nil || fh == nil || fh.Size == 0 {
		c.JSON(http.StatusBadRequest, lerror.ToAPIError("hub", fmt.Errorf("missing or empty file")))
		return
	}

	// persist the ciphertext to a temp file — sdk.Upload is file-based, and the
	// hub never inspects the bytes (client-side encrypted).
	src, err := fh.Open()
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	defer src.Close()
	tmp, err := os.CreateTemp("", "seal-*.bin")
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	if _, err := io.Copy(tmp, src); err != nil {
		tmp.Close()
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	tmp.Close()

	name := c.PostForm("name") // optional; "" => hub-generated (commitment hex)

	sk := s.rp.Key().Export().PrivateKey
	au, err := key.BuildAuth(sk, []byte("seal"))
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	cm, err := s.chainManager()
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	// 1. erasure-encode + stream-stage + KZG commit.
	res, streamer, err := sdk.Upload(sdk.ServerURL, au, policy, tmpPath, name)
	if err != nil {
		c.JSON(http.StatusBadGateway, lerror.ToAPIError("hub", err))
		return
	}
	// 2. trustless check: the streamer encoded the actual bytes.
	pcs, err := sdk.CheckFileFull(res, streamer, tmpPath)
	if err != nil {
		c.JSON(http.StatusBadGateway, lerror.ToAPIError("hub", err))
		return
	}
	if len(pcs) == 0 {
		c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("no piece produced")))
		return
	}
	pc := pcs[0] // a seal segment (<= ~1GB/piece) is a single piece

	// resolve start/expire/price so the response (and v2 cost) are exact and
	// match what AddPiece locks.
	start, err := cm.GetEpoch()
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	expire := start + uint64(com.DefaultStoreEpoch)
	if v := c.PostForm("expire"); v != "" {
		if e, err := strconv.ParseUint(v, 10, 64); err == nil {
			expire = e
		}
	}
	price := big.NewInt(int64(com.DefaultReplicaPrice))
	if v := c.PostForm("price"); v != "" {
		if p, ok := new(big.Int).SetString(v, 10); ok {
			price = p
		}
	}
	pc.Start = start
	pc.Expire = expire
	pc.Price = price

	// idempotency: da_cid is content-addressed, so a retried seal of the same
	// blob yields the same piece. If it is already on chain, don't re-register.
	serial, _ := cm.GetPieceSerial(pc.Name)

	if register == "hub" {
		// v1: hub signs AddPiece + pays gas; client needs no chain interaction.
		txn := ""
		if serial == 0 {
			txn, err = cm.AddPiece(pc)
			if err != nil {
				c.JSON(599, lerror.ToAPIError("hub", err))
				return
			}
			serial, _ = cm.GetPieceSerial(pc.Name) // best-effort
		}
		c.JSON(http.StatusOK, gin.H{
			"register":     "hub",
			"da_cid":       pc.Name,
			"size":         pc.Size,
			"policy":       gin.H{"n": policy.N, "k": policy.K},
			"expire":       expire,
			"add_piece_tx": txn, // "" if already registered (idempotent retry)
			"piece_serial": serial,
		})
		return
	}

	// v2: hub does NOT register; return the fields the client self-signs with.
	pn, err := com.G1StringInSolidity(pc.Name)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"register":       "client",
		"da_cid":         pc.Name,
		"pn_solidity":    hex.EncodeToString(pn),
		"size":           pc.Size,
		"policy":         gin.H{"n": policy.N, "k": policy.K},
		"price":          price.String(),
		"start":          start,
		"expire":         expire,
		"streamer":       streamer.Hex(),
		"cost":           contract.AddPieceCost(pc).String(),
		"piece_contract": cm.PieceAddr.Hex(),
		"token_contract": cm.TokenAddr.Hex(),
		"piece_serial":   serial, // >0 if already registered (idempotent)
		// client must AddPiece before this epoch or the piece falls out of
		// stream staging (~15 min) and is under-replicated (spec Q2).
		"staging_deadline_epoch": pc.Start + uint64(com.DelaySubmit),
	})
}
