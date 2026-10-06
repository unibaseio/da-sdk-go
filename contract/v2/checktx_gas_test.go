package contract

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func failedTx(t *testing.T, f *fakeRPC, gas uint64) common.Hash {
	t.Helper()
	sk, _ := crypto.GenerateKey()
	to := common.HexToAddress("0x1234")
	tx, err := etypes.SignNewTx(sk, etypes.LatestSignerForChainID(big.NewInt(31337)), &etypes.LegacyTx{
		Nonce: 0, To: &to, Gas: gas, GasPrice: big.NewInt(1), Value: big.NewInt(0),
	})
	if err != nil {
		t.Fatal(err)
	}
	f.known[tx.Hash()] = tx
	f.receiptFailed = true
	return tx.Hash()
}

// T-OOG: a failed tx that was not first in its block (GasUsed !=
// CumulativeGasUsed) was reported as "exceed gas limit" although it used a
// fraction of its gas.
func TestCheckTxFailureNotFirstInBlock(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	h := failedTx(t, f, 500000)
	f.receiptGas, f.receiptCumul = 30000, 90000 // third tx of its block, used 6% of its gas

	err := c.CheckTx(h)
	if err == nil {
		t.Fatal("failed tx reported as success")
	}
	if strings.Contains(err.Error(), "gas") || !strings.Contains(err.Error(), "mined but execution failed") {
		t.Fatalf("got %v, want a plain execution failure (not a gas one)", err)
	}
}

// A failed tx that used its whole limit is reported as out of gas OR an
// INVALID opcode — both consume all gas, so neither can be claimed alone.
func TestCheckTxFailureUsedAllGas(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	h := failedTx(t, f, 50000)
	f.receiptGas, f.receiptCumul = 50000, 50000

	err := c.CheckTx(h)
	if err == nil || !strings.Contains(err.Error(), "out of gas, or an INVALID opcode") {
		t.Fatalf("got %v", err)
	}
}
