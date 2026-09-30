// Package evidence checks that an off-chain file referenced from an ERC-8004
// registry (agent registration file, feedback file, validation evidence) is
// intact and still available on Unibase DA.
//
// ERC-8004 stores a URI plus the keccak256 of the file; that proves the file
// was not changed but not that it still exists. Verify adds the DA side: the
// file's piece is registered on-chain, unexpired, and backed by at least K
// healthy replicas, so it can still be rebuilt from store nodes.
package evidence

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	contract "github.com/unibaseio/da-sdk-go/contract/v2"
)

// Chain is the read-only on-chain view Verify needs. NewChain adapts a
// *contract.ContractManage to it.
type Chain interface {
	GetPieceSerial(piece string) (uint64, error)
	GetPieceRS(pieceIndex uint64) (n, k uint8, expire uint64, err error)
	GetPieceReplica(pieceIndex uint64, slot uint8) (replica uint64, storedOn common.Address, err error)
	GetRSFake(pieceIndex uint64, slot uint8) (bool, error)
	NodeActive(addr common.Address) (bool, error)
	CurrentEpoch() (uint64, error)
}

type cmChain struct{ cm *contract.ContractManage }

// NewChain adapts a contract manager to Chain.
func NewChain(cm *contract.ContractManage) Chain { return cmChain{cm} }

func (c cmChain) GetPieceSerial(p string) (uint64, error) { return c.cm.GetPieceSerial(p) }
func (c cmChain) GetPieceRS(pi uint64) (uint8, uint8, uint64, error) {
	return c.cm.GetPieceRS(pi)
}
func (c cmChain) GetPieceReplica(pi uint64, slot uint8) (uint64, common.Address, error) {
	return c.cm.GetPieceReplica(pi, slot)
}
func (c cmChain) GetRSFake(pi uint64, slot uint8) (bool, error) { return c.cm.GetRSFake(pi, slot) }
func (c cmChain) NodeActive(a common.Address) (bool, error) {
	info, err := c.cm.GetPledgeInfo(a)
	return info.IsActive, err
}
func (c cmChain) CurrentEpoch() (uint64, error) { return c.cm.GetEpoch() }

// Range locates an object's bytes inside its DA piece (from the hub's /proof).
type Range struct {
	Volume uint64 `json:"volume"`
	Start  uint64 `json:"start"`
	Size   uint64 `json:"size"`
}

// Proof is the hub's verification bundle (GET …/objects/{key}/proof).
type Proof struct {
	Commitment string `json:"commitment"`
	Sha256     string `json:"sha256"`
	Keccak256  string `json:"keccak256"`
	Range      *Range `json:"range"`
	RangeNote  string `json:"rangeNote"`
	Chain      struct {
		TxHash    string `json:"txHash"`
		ChainType string `json:"chainType"`
	} `json:"chain"`
}

// TrustlessCheck is the result of hashing the object's range inside a piece
// rebuilt from store nodes, bypassing the hub's copy.
type TrustlessCheck struct {
	Hash string `json:"hash"`
	OK   bool   `json:"ok"`
}

// Result reports the three checks: content (hash matches), registered (piece on
// chain), available (unexpired with at least K healthy replicas).
type Result struct {
	URI          string          `json:"uri"`
	ExpectedHash string          `json:"expectedHash"`
	ContentHash  string          `json:"contentHash"`
	ContentOK    bool            `json:"contentOk"`
	Staged       bool            `json:"staged,omitempty"` // on the hub, not yet committed on-chain
	Committed    bool            `json:"committed"`
	Piece        string          `json:"piece,omitempty"`
	TxHash       string          `json:"txHash,omitempty"`
	Range        *Range          `json:"range,omitempty"`
	Registered   bool            `json:"registered"`
	PieceIndex   uint64          `json:"pieceIndex,omitempty"`
	N            uint8           `json:"n,omitempty"`
	K            uint8           `json:"k,omitempty"`
	ExpireEpoch  uint64          `json:"expireEpoch,omitempty"`
	CurrentEpoch uint64          `json:"currentEpoch,omitempty"`
	LiveReplicas int             `json:"liveReplicas"`
	Available    bool            `json:"available"`
	Trustless    *TrustlessCheck `json:"trustless,omitempty"`
	Notes        []string        `json:"notes,omitempty"`
}

// Options tunes Verify. The zero value is usable.
type Options struct {
	HTTPClient *http.Client
	// MaxBytes caps the evidence download (default 32 MiB).
	MaxBytes int64
	// FetchPiece, when set, rebuilds a piece from store nodes; Verify then
	// hashes the object's range inside it instead of trusting the hub's copy.
	FetchPiece func(piece string) ([]byte, error)
}

const defaultMaxBytes = 32 << 20

var hubObjectPath = regexp.MustCompile(`^/v1/buckets/[^/]+/objects/[^/]+$`)

// Verify checks the evidence at uri against the keccak256 recorded on-chain
// (expectedHash, 0x-hex). It returns an error only when the evidence cannot be
// fetched or the hub answers unexpectedly; failed checks are reported in the
// Result, not as errors.
func Verify(uri, expectedHash string, chain Chain, opt Options) (Result, error) {
	res := Result{URI: uri}
	want, err := normHash(expectedHash)
	if err != nil {
		return res, err
	}
	res.ExpectedHash = want

	hc := opt.HTTPClient
	if hc == nil {
		hc = &http.Client{Timeout: 60 * time.Second}
	}
	limit := opt.MaxBytes
	if limit <= 0 {
		limit = defaultMaxBytes
	}

	// 1. content
	body, status, err := get(hc, uri, limit)
	if err != nil {
		return res, fmt.Errorf("fetch evidence: %w", err)
	}
	if status != http.StatusOK {
		return res, fmt.Errorf("fetch evidence: HTTP %d", status)
	}
	res.ContentHash = keccakHex(body)
	res.ContentOK = res.ContentHash == want

	// 2. where it lives in DA (hub /proof)
	proofURL, ok := proofURLFor(uri)
	if !ok {
		res.Notes = append(res.Notes, "not a hub object URL; DA availability not checked")
		return res, nil
	}
	pb, status, err := get(hc, proofURL, 1<<20)
	if err != nil {
		return res, fmt.Errorf("fetch proof: %w", err)
	}
	switch status {
	case http.StatusOK:
	case http.StatusTooEarly:
		res.Staged = true
		res.Notes = append(res.Notes, "staged on the hub, not yet committed on-chain; retry later")
		return res, nil
	default:
		return res, fmt.Errorf("fetch proof: HTTP %d", status)
	}
	var p Proof
	if err := json.Unmarshal(pb, &p); err != nil {
		return res, fmt.Errorf("decode proof: %w", err)
	}
	if p.Commitment == "" {
		return res, fmt.Errorf("proof has no commitment")
	}
	res.Committed = true
	res.Piece = p.Commitment
	res.TxHash = p.Chain.TxHash
	res.Range = p.Range
	if p.Keccak256 != "" && !strings.EqualFold(p.Keccak256, res.ContentHash) {
		res.Notes = append(res.Notes, "hub-reported keccak256 differs from the fetched content")
	}
	if p.RangeNote != "" {
		res.Notes = append(res.Notes, p.RangeNote)
	}

	// 3. registered + available
	if chain == nil {
		res.Notes = append(res.Notes, "no chain client; on-chain checks skipped")
		return res, nil
	}
	pi, err := chain.GetPieceSerial(p.Commitment)
	if err != nil {
		return res, fmt.Errorf("piece index: %w", err)
	}
	res.PieceIndex = pi
	res.Registered = pi > 0
	if !res.Registered {
		res.Notes = append(res.Notes, "piece not registered on the Piece contract")
		return res, nil
	}
	n, k, expire, err := chain.GetPieceRS(pi)
	if err != nil {
		return res, fmt.Errorf("piece policy: %w", err)
	}
	cur, err := chain.CurrentEpoch()
	if err != nil {
		return res, fmt.Errorf("current epoch: %w", err)
	}
	res.N, res.K, res.ExpireEpoch, res.CurrentEpoch = n, k, expire, cur

	for slot := uint8(0); slot < n; slot++ {
		_, storedOn, err := chain.GetPieceReplica(pi, slot)
		if err != nil {
			return res, fmt.Errorf("replica %d: %w", slot, err)
		}
		if storedOn == (common.Address{}) {
			continue
		}
		fake, err := chain.GetRSFake(pi, slot)
		if err != nil {
			return res, fmt.Errorf("replica %d fraud status: %w", slot, err)
		}
		if fake {
			continue
		}
		active, err := chain.NodeActive(storedOn)
		if err != nil {
			return res, fmt.Errorf("replica %d node status: %w", slot, err)
		}
		if active {
			res.LiveReplicas++
		}
	}
	expired := cur >= expire
	if expired {
		res.Notes = append(res.Notes, fmt.Sprintf("piece expired at epoch %d (current %d)", expire, cur))
	}
	if res.LiveReplicas < int(k) {
		res.Notes = append(res.Notes, fmt.Sprintf("only %d healthy replicas, need %d to rebuild", res.LiveReplicas, k))
	}
	res.Available = !expired && res.LiveReplicas >= int(k)

	// 4. optional: rebuild the piece from store nodes and hash the range
	if opt.FetchPiece != nil {
		if p.Range == nil {
			res.Notes = append(res.Notes, "no byte range from the hub; trustless check skipped")
			return res, nil
		}
		data, err := opt.FetchPiece(p.Commitment)
		if err != nil {
			return res, fmt.Errorf("rebuild piece: %w", err)
		}
		end := p.Range.Start + p.Range.Size
		if end > uint64(len(data)) || end < p.Range.Start {
			res.Trustless = &TrustlessCheck{}
			res.Notes = append(res.Notes, fmt.Sprintf("range %d+%d is outside the %d-byte piece", p.Range.Start, p.Range.Size, len(data)))
			return res, nil
		}
		h := keccakHex(data[p.Range.Start:end])
		res.Trustless = &TrustlessCheck{Hash: h, OK: h == want}
	}
	return res, nil
}

// proofURLFor maps a hub object URL to its /proof URL, keeping the query
// (the owner parameter).
func proofURLFor(uri string) (string, bool) {
	u, err := url.Parse(uri)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		return "", false
	}
	if !hubObjectPath.MatchString(u.Path) {
		return "", false
	}
	u.Path += "/proof"
	return u.String(), true
}

func get(hc *http.Client, u string, limit int64) ([]byte, int, error) {
	resp, err := hc.Get(u)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(io.LimitReader(resp.Body, limit+1))
	if err != nil {
		return nil, resp.StatusCode, err
	}
	if int64(len(b)) > limit {
		return nil, resp.StatusCode, fmt.Errorf("response larger than %d bytes", limit)
	}
	return b, resp.StatusCode, nil
}

func normHash(h string) (string, error) {
	h = strings.ToLower(strings.TrimSpace(h))
	h = strings.TrimPrefix(h, "0x")
	if len(h) != 64 {
		return "", fmt.Errorf("expected a 32-byte keccak256 hash, got %d hex chars", len(h))
	}
	if _, err := hex.DecodeString(h); err != nil {
		return "", fmt.Errorf("bad hash: %w", err)
	}
	return "0x" + h, nil
}

// KeccakHex is the 0x-hex keccak256 of b — the hash to put in ERC-8004's
// feedbackHash / requestHash / responseHash fields.
func KeccakHex(b []byte) string { return keccakHex(b) }

func keccakHex(b []byte) string {
	return "0x" + hex.EncodeToString(crypto.Keccak256(b))
}
