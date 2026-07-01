package metering

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// The canonical JSON encoding and its hash are a wire contract: lock them so an
// accidental struct-field reorder or tag change is caught.
func TestSettlementReportCanonicalJSONAndHash(t *testing.T) {
	r := SettlementReport{
		Type:         ReportType,
		Owner:        "0xabc",
		AmountWei:    "1000",
		FromEventID:  1,
		ToEventID:    5,
		Writes:       3,
		Reads:        0,
		BytesWritten: 123,
		Timestamp:    1700000000,
	}
	const wantJSON = `{"amount_wei":"1000","bytes_written":123,"from_event_id":1,"owner":"0xabc","reads":0,"timestamp":1700000000,"to_event_id":5,"type":"da-hub-metering-settlement","writes":3}`
	const wantHash = "0xd66917ec882f5e24ae96921789926a21e5f56bb38ba4a9d5d65673abfdf41d30"

	b, err := r.CanonicalJSON()
	if err != nil {
		t.Fatalf("CanonicalJSON: %v", err)
	}
	if string(b) != wantJSON {
		t.Fatalf("canonical json mismatch:\n got %s\nwant %s", b, wantJSON)
	}

	sum, hexStr, err := r.Hash()
	if err != nil {
		t.Fatalf("Hash: %v", err)
	}
	if hexStr != wantHash {
		t.Fatalf("hash = %s, want %s", hexStr, wantHash)
	}
	// [32]byte and hex must agree
	if "0x"+common.Bytes2Hex(sum[:]) != hexStr {
		t.Fatalf("bytes/hex disagree")
	}
}

func TestSettlementReportHashDeterministic(t *testing.T) {
	mk := func() SettlementReport {
		return SettlementReport{Type: ReportType, Owner: "0xdeadbeef", AmountWei: "999", Timestamp: 42}
	}
	_, h1, _ := mk().Hash()
	_, h2, _ := mk().Hash()
	if h1 != h2 {
		t.Fatalf("hash not deterministic: %s vs %s", h1, h2)
	}
}

func TestNewERC8183ClientValidation(t *testing.T) {
	// Missing RPC.
	if _, err := newERC8183Client(Config{}); err == nil {
		t.Error("expected error for missing RPC URL")
	}
	// Bad contract addr.
	if _, err := newERC8183Client(Config{ChainRPCURL: "http://x", ERC8183ContractAddr: "nope"}); err == nil {
		t.Error("expected error for bad contract addr")
	}
	// Missing provider key.
	_, err := newERC8183Client(Config{
		ChainRPCURL:          "http://x",
		ERC8183ContractAddr:  "0x1111111111111111111111111111111111111111",
		ERC8183EvaluatorAddr: "0x2222222222222222222222222222222222222222",
		ERC20TokenAddr:       "0x3333333333333333333333333333333333333333",
	})
	if err == nil {
		t.Error("expected error for missing provider key")
	}
}

// Live integration test. Skipped unless RUN_LIVE_ERC8183_TESTS=1 and the chain
// env is configured. It never runs in CI by default.
func TestLiveERC8183CreateJob(t *testing.T) {
	if os.Getenv("RUN_LIVE_ERC8183_TESTS") != "1" {
		t.Skip("set RUN_LIVE_ERC8183_TESTS=1 and chain env to run the live ERC-8183 test")
	}
	cfg := LoadConfigFromEnv()
	c, err := newERC8183Client(cfg)
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	expiredAt := big.NewInt(time.Now().Add(time.Hour).Unix())
	jobID, tx, err := c.CreateJob(ctx, expiredAt, "live-test", big.NewInt(0))
	if err != nil {
		t.Fatalf("CreateJob: %v", err)
	}
	t.Logf("created job %s in tx %s", jobID, tx.Hex())
}
