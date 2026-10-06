package common

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// T-OOG: a failed tx that consumed its whole gas limit was reported as "ran
// out of gas", but an INVALID opcode consumes all gas too. The message must
// not claim either cause alone.
func TestCheckTxUsedAllGasWording(t *testing.T) {
	sk, _ := crypto.GenerateKey()
	to := common.HexToAddress("0x1234")
	tx, err := types.SignNewTx(sk, types.LatestSignerForChainID(big.NewInt(31337)), &types.LegacyTx{
		To: &to, Gas: 50000, GasPrice: big.NewInt(1), Value: big.NewInt(0),
	})
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var req struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
		}
		_ = json.Unmarshal(body, &req)
		var res interface{}
		switch req.Method {
		case "eth_getTransactionReceipt":
			res = &types.Receipt{Status: types.ReceiptStatusFailed, GasUsed: 50000, CumulativeGasUsed: 50000,
				Logs: []*types.Log{}, TxHash: tx.Hash(), BlockHash: common.Hash{2}, BlockNumber: big.NewInt(16)}
		case "eth_getTransactionByHash":
			b, _ := tx.MarshalJSON()
			m := map[string]interface{}{}
			_ = json.Unmarshal(b, &m)
			m["blockNumber"], m["blockHash"], m["transactionIndex"] = "0x10", common.Hash{2}, "0x0"
			res = m
		case "eth_call":
			res = "0x"
		default:
			t.Errorf("unexpected %s", req.Method)
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "id": req.ID, "result": res})
	}))
	defer srv.Close()

	err = CheckTx(srv.URL, tx.Hash())
	if err == nil {
		t.Fatal("failed tx reported as success")
	}
	if !strings.Contains(err.Error(), "out of gas, or an INVALID opcode") {
		t.Fatalf("got %q, want both possible causes named", err)
	}
}
