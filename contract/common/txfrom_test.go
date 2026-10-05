package common

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// The sender of a failed tx is recovered for every tx type, so its replay
// runs as the real account (it was 0x0 for EIP-1559 txs).
func TestGetFromAllTxTypes(t *testing.T) {
	sk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	want := crypto.PubkeyToAddress(sk.PublicKey)
	chain := big.NewInt(84532)
	to := common.HexToAddress("0x1")

	txs := map[string]struct {
		tx     *types.Transaction
		signer types.Signer
	}{
		"homestead": {types.NewTx(&types.LegacyTx{Nonce: 1, To: &to, Gas: 21000, GasPrice: big.NewInt(1)}), types.HomesteadSigner{}},
		"eip155":    {types.NewTx(&types.LegacyTx{Nonce: 1, To: &to, Gas: 21000, GasPrice: big.NewInt(1)}), types.NewEIP155Signer(chain)},
		"accesslist": {types.NewTx(&types.AccessListTx{ChainID: chain, Nonce: 1, To: &to, Gas: 21000, GasPrice: big.NewInt(1)}),
			types.LatestSignerForChainID(chain)},
		"eip1559": {types.NewTx(&types.DynamicFeeTx{ChainID: chain, Nonce: 1, To: &to, Gas: 21000, GasTipCap: big.NewInt(1), GasFeeCap: big.NewInt(3)}),
			types.LatestSignerForChainID(chain)},
	}
	for name, c := range txs {
		stx, err := types.SignTx(c.tx, c.signer, sk)
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if got := getFrom(stx); got != want {
			t.Errorf("%s: sender %s, want %s", name, got, want)
		}
	}
}
