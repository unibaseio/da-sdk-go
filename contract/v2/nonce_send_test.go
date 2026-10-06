package contract

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	com "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/piece"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// N5: a nonce reserved for a transaction that never left this process must be
// handed back, or every later transaction of the key queues behind the gap
// until the 5-minute stuck-nonce resync. These tests drive the real senders in
// set.go against an in-process fake node (fakerpc_test.go).

var someStore = common.HexToAddress("0x00000000000000000000000000000000000000aa")

// Gas estimation reverts (e.g. the game moved on): nothing was broadcast, so
// the next transaction must reuse the same nonce.
func TestNonceReleasedWhenGasEstimationReverts(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)

	f.rpcErr["eth_estimateGas"] = jsonRPCError{Code: 3, Message: "execution reverted: exceed pt"}
	if err := c.ChallengeOne(someStore, 7, 1); err == nil {
		t.Fatal("ChallengeOne succeeded although gas estimation reverted")
	}
	if n := f.count("eth_sendRawTransaction"); n != 0 {
		t.Fatalf("%d txs broadcast after a reverted estimation", n)
	}

	delete(f.rpcErr, "eth_estimateGas")
	if err := c.ChallengeOne(someStore, 7, 1); err != nil {
		t.Fatalf("second ChallengeOne: %v", err)
	}
	if len(f.sent) != 1 || f.sent[0].Nonce() != 0 {
		t.Fatalf("next tx nonce = %d, want 0 (the reverted send's nonce was leaked)", f.sent[0].Nonce())
	}
}

// The node answers the broadcast with an error (here: insufficient funds): the
// tx is in no pool, so its nonce is free again.
func TestNonceReleasedWhenNodeRejectsTx(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)

	f.rpcErr["eth_sendRawTransaction"] = jsonRPCError{Code: -32000, Message: "insufficient funds for gas * price + value"}
	if err := c.ChallengeOne(someStore, 7, 1); err == nil {
		t.Fatal("ChallengeOne succeeded although the node rejected the tx")
	}
	delete(f.rpcErr, "eth_sendRawTransaction")
	if err := c.ChallengeOne(someStore, 7, 1); err != nil {
		t.Fatalf("second ChallengeOne: %v", err)
	}
	if got := f.sent[0].Nonce(); got != 0 {
		t.Fatalf("next tx nonce = %d, want 0 (the rejected tx's nonce was leaked)", got)
	}
}

// An early return before anything is sent (AddReplica on a replica already on
// chain) must not reserve a nonce at all.
func TestNoNonceReservedOnEarlyReturn(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)

	pabi, err := abi.JSON(strings.NewReader(piece.PieceABI))
	if err != nil {
		t.Fatal(err)
	}
	word := func(v uint64) []byte {
		return common.LeftPadBytes(new(big.Int).SetUint64(v).Bytes(), 32)
	}
	f.callResult[common.Bytes2Hex(pabi.Methods["getPIndex"].ID)] = word(3)
	f.callResult[common.Bytes2Hex(pabi.Methods["getRIndex"].ID)] = word(9) // already registered

	var g bls.G1
	g.ScalarMultiplicationBase(big.NewInt(5))
	name := com.G1ToString(g)
	err = c.AddReplica(types.ReplicaCore{Name: name, Piece: name, Index: 0}, []byte{1})
	if err == nil || !strings.Contains(err.Error(), "already on chain") {
		t.Fatalf("AddReplica: %v, want 'already on chain'", err)
	}

	if err := c.ChallengeOne(someStore, 7, 1); err != nil {
		t.Fatalf("ChallengeOne: %v", err)
	}
	if got := f.sent[0].Nonce(); got != 0 {
		t.Fatalf("next tx nonce = %d, want 0 (AddReplica's early return leaked a nonce)", got)
	}
}

// "already known": the node holds this very tx (e.g. a retried broadcast), so
// it was sent; the nonce is spent and the call goes on to wait for it.
func TestAlreadyKnownCountsAsSent(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)

	f.rpcErr["eth_sendRawTransaction"] = jsonRPCError{Code: -32000, Message: "already known"}
	if err := c.ChallengeOne(someStore, 7, 1); err != nil {
		t.Fatalf("ChallengeOne with an already-pooled tx: %v", err)
	}
	if c.localNonce != 1 || !c.nonceReady {
		t.Fatalf("localNonce=%d ready=%v, want 1 true (nonce spent)", c.localNonce, c.nonceReady)
	}
}

// No answer to the broadcast: the node may have taken the tx, so the nonce is
// kept (reusing it could replace a pooled tx) — unless the node is then seen
// to hold the tx, in which case the call carries on as sent.
func TestNoAnswerKeepsNonce(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)

	f.drop["eth_sendRawTransaction"] = true
	if err := c.ChallengeOne(someStore, 7, 1); err == nil {
		t.Fatal("ChallengeOne succeeded although the broadcast got no answer")
	}
	if c.localNonce != 1 || !c.nonceReady {
		t.Fatalf("localNonce=%d ready=%v after an unanswered broadcast, want 1 true (kept)", c.localNonce, c.nonceReady)
	}
	if n := f.count("eth_getTransactionByHash"); n == 0 {
		t.Fatal("an unanswered broadcast was not looked up")
	}
}

// Another tx already holds the nonce: the next reservation re-reads the chain.
func TestNonceTooLowResyncs(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)

	f.rpcErr["eth_sendRawTransaction"] = jsonRPCError{Code: -32000, Message: "nonce too low: next nonce 5, tx nonce 0"}
	if err := c.ChallengeOne(someStore, 7, 1); err == nil {
		t.Fatal("ChallengeOne succeeded with nonce too low")
	}
	delete(f.rpcErr, "eth_sendRawTransaction")
	f.pendingNonce = 5
	if err := c.ChallengeOne(someStore, 7, 1); err != nil {
		t.Fatalf("ChallengeOne: %v", err)
	}
	if got := f.sent[0].Nonce(); got != 5 {
		t.Fatalf("next tx nonce = %d, want the chain's 5", got)
	}
}

// The broadcast's answer is lost but the node did take the tx: the lookup
// finds it and the call carries on as sent.
func TestNoAnswerButPooledCountsAsSent(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)

	f.dropAfter["eth_sendRawTransaction"] = true
	f.knowsSent = true
	if err := c.ChallengeOne(someStore, 7, 1); err != nil {
		t.Fatalf("ChallengeOne whose tx the node holds: %v", err)
	}
	if len(f.sent) != 1 || c.localNonce != 1 {
		t.Fatalf("sent=%d localNonce=%d, want 1 1", len(f.sent), c.localNonce)
	}
}
