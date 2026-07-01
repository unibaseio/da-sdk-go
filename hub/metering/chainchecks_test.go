package metering

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// fakeERC20 is an injectable erc20API for tests. balance/allowance are fixed;
// err (if set) is returned from the read calls to simulate RPC failure.
type fakeERC20 struct {
	balance   *big.Int
	allowance *big.Int
	err       error

	// transfer* control TransferFrom for settlement tests (zero-valued by
	// default: returns the zero hash and nil error, preserving prior behavior).
	transferHash common.Hash
	transferErr  error

	balanceCalls   int
	allowanceCalls int
	transferCalls  int
}

func (f *fakeERC20) BalanceOf(ctx context.Context, owner common.Address) (*big.Int, error) {
	f.balanceCalls++
	if f.err != nil {
		return nil, f.err
	}
	return f.balance, nil
}

func (f *fakeERC20) Allowance(ctx context.Context, owner, spender common.Address) (*big.Int, error) {
	f.allowanceCalls++
	if f.err != nil {
		return nil, f.err
	}
	return f.allowance, nil
}

func (f *fakeERC20) TransferFrom(ctx context.Context, from, to common.Address, amount *big.Int) (common.Hash, error) {
	f.transferCalls++
	return f.transferHash, f.transferErr
}

const chainOwner = "0xaaaa0000000000000000000000000000000000aa"

// With CheckChain off, no RPC call is made even if a reader is present.
func TestCanWriteChainDisabledNoRPC(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:      true,
		ChargeWrites: true,
		WriteBaseWei: big.NewInt(100),
		CheckChain:   false,
	})
	fake := &fakeERC20{balance: big.NewInt(0), allowance: big.NewInt(0)}
	m.erc20 = fake

	res, err := m.CanWrite(chainOwner, 1)
	if err != nil {
		t.Fatalf("CanWrite: %v", err)
	}
	if !res.Allowed || res.Reason != reasonAllowed {
		t.Fatalf("want allowed, got %+v", res)
	}
	if fake.balanceCalls != 0 || fake.allowanceCalls != 0 {
		t.Fatalf("chain calls made while CheckChain=false: bal=%d allow=%d", fake.balanceCalls, fake.allowanceCalls)
	}
}

func TestCanWriteInsufficientBalance402(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:      true,
		ChargeWrites: true,
		WriteBaseWei: big.NewInt(1000),
		CheckChain:   true,
	})
	m.erc20 = &fakeERC20{balance: big.NewInt(500), allowance: bi("999999999")}

	res, _ := m.CanWrite(chainOwner, 1) // required = 1000, balance = 500
	if res.Allowed || res.Reason != reasonInsufficientBalance {
		t.Fatalf("want insufficient_balance, got %+v", res)
	}
	if res.BalanceWei != "500" {
		t.Errorf("balance = %s, want 500", res.BalanceWei)
	}
}

func TestCanWriteInsufficientAllowance402(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:      true,
		ChargeWrites: true,
		WriteBaseWei: big.NewInt(1000),
		CheckChain:   true,
	})
	// balance is plenty, allowance too small
	m.erc20 = &fakeERC20{balance: bi("999999999"), allowance: big.NewInt(500)}

	res, _ := m.CanWrite(chainOwner, 1) // required = 1000
	if res.Allowed || res.Reason != reasonInsufficientAllow {
		t.Fatalf("want insufficient_allowance, got %+v", res)
	}
	if res.AllowanceWei != "500" {
		t.Errorf("allowance = %s, want 500", res.AllowanceWei)
	}
}

func TestCanWriteChainSufficientAllows(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:      true,
		ChargeWrites: true,
		WriteBaseWei: big.NewInt(1000),
		CheckChain:   true,
	})
	m.erc20 = &fakeERC20{balance: bi("2000"), allowance: bi("2000")}

	res, _ := m.CanWrite(chainOwner, 1)
	if !res.Allowed || res.Reason != reasonAllowed {
		t.Fatalf("want allowed, got %+v", res)
	}
}

// An RPC error surfaces as chain_check_failed (a 402 refusal), not a Go error.
func TestCanWriteChainRPCErrorReported(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:      true,
		ChargeWrites: true,
		WriteBaseWei: big.NewInt(1000),
		CheckChain:   true,
	})
	m.erc20 = &fakeERC20{err: fmt.Errorf("dial tcp: connection refused")}

	res, err := m.CanWrite(chainOwner, 1)
	if err != nil {
		t.Fatalf("CanWrite must not return a Go error on RPC failure: %v", err)
	}
	if res.Allowed || res.Reason != reasonChainCheckFailed {
		t.Fatalf("want chain_check_failed, got %+v", res)
	}
}

// CheckChain on but no client configured -> chain_check_failed.
func TestCanWriteChainNoClient(t *testing.T) {
	m := newTestManager(t, Config{
		Enabled:      true,
		ChargeWrites: true,
		WriteBaseWei: big.NewInt(1000),
		CheckChain:   true,
	})
	// m.erc20 left nil (no valid chain config)
	res, _ := m.CanWrite(chainOwner, 1)
	if res.Allowed || res.Reason != reasonChainCheckFailed {
		t.Fatalf("want chain_check_failed, got %+v", res)
	}
}
