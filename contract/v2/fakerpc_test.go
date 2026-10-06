package contract

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// fakeRPC is a minimal in-process JSON-RPC node: enough of the eth_ API for a
// contract binding to build, sign and send a transaction, and for CheckTx to
// see it mined. Each method can be overridden to return an error, and the
// calls made are recorded. No chain, no Anvil.
type fakeRPC struct {
	t   *testing.T
	srv *httptest.Server

	mu    sync.Mutex
	calls map[string]int
	sent  []*types.Transaction

	pendingNonce uint64
	// per-method overrides: return this JSON-RPC error instead of a result
	rpcErr map[string]jsonRPCError
	// methods whose HTTP connection is dropped without an answer
	drop map[string]bool
	// methods that are processed (a sent tx is recorded) but whose answer is lost
	dropAfter map[string]bool
	// eth_call result by 4-byte selector (hex, no 0x); default: 32 zero bytes
	callResult map[string][]byte
	// eth_getTransactionByHash answers "found" for sent txs when true
	knowsSent bool
	// txs eth_getTransactionByHash knows regardless of knowsSent
	known map[common.Hash]*types.Transaction
	// receipt shape: failed status and gas figures (zero = defaults)
	receiptFailed            bool
	receiptGas, receiptCumul uint64
}

type jsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newFakeRPC(t *testing.T) *fakeRPC {
	f := &fakeRPC{
		t:          t,
		calls:      map[string]int{},
		rpcErr:     map[string]jsonRPCError{},
		drop:       map[string]bool{},
		dropAfter:  map[string]bool{},
		known:      map[common.Hash]*types.Transaction{},
		callResult: map[string][]byte{},
	}
	f.srv = httptest.NewServer(http.HandlerFunc(f.serve))
	t.Cleanup(f.srv.Close)
	return f
}

func (f *fakeRPC) count(m string) int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.calls[m]
}

func (f *fakeRPC) serve(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var req struct {
		ID     json.RawMessage   `json:"id"`
		Method string            `json:"method"`
		Params []json.RawMessage `json:"params"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "batch not supported", http.StatusBadRequest)
		return
	}
	f.mu.Lock()
	f.calls[req.Method]++
	e, hasErr := f.rpcErr[req.Method]
	drop := f.drop[req.Method]
	dropAfter := f.dropAfter[req.Method]
	f.mu.Unlock()

	if dropAfter {
		f.result(req.Method, req.Params)
	}
	if drop || dropAfter {
		hj, ok := w.(http.Hijacker)
		if !ok {
			f.t.Fatal("cannot hijack")
		}
		conn, _, _ := hj.Hijack()
		conn.Close()
		return
	}

	resp := map[string]interface{}{"jsonrpc": "2.0", "id": req.ID}
	if hasErr {
		resp["error"] = e
	} else {
		resp["result"] = f.result(req.Method, req.Params)
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func (f *fakeRPC) result(method string, params []json.RawMessage) interface{} {
	switch method {
	case "eth_chainId":
		return "0x7a69"
	case "eth_blockNumber":
		return "0x10"
	case "eth_getBlockByNumber":
		h := &types.Header{
			ParentHash: common.Hash{1},
			Difficulty: big.NewInt(1),
			Number:     big.NewInt(16),
			GasLimit:   30_000_000,
			Time:       1,
		}
		return h
	case "eth_getTransactionCount":
		f.mu.Lock()
		defer f.mu.Unlock()
		return hexutil.Uint64(f.pendingNonce)
	case "eth_getCode":
		return "0x6080"
	case "eth_estimateGas":
		return "0x186a0"
	case "eth_gasPrice", "eth_maxPriorityFeePerGas":
		return "0x3b9aca00"
	case "eth_getBalance":
		return "0x0"
	case "eth_call":
		var call struct {
			Data  hexutil.Bytes `json:"data"`
			Input hexutil.Bytes `json:"input"`
		}
		_ = json.Unmarshal(params[0], &call)
		data := call.Input
		if len(data) == 0 {
			data = call.Data
		}
		if len(data) >= 4 {
			f.mu.Lock()
			res, ok := f.callResult[common.Bytes2Hex(data[:4])]
			f.mu.Unlock()
			if ok {
				return hexutil.Bytes(res)
			}
		}
		return hexutil.Bytes(make([]byte, 32))
	case "eth_sendRawTransaction":
		var raw hexutil.Bytes
		_ = json.Unmarshal(params[0], &raw)
		tx := new(types.Transaction)
		if err := tx.UnmarshalBinary(raw); err != nil {
			f.t.Errorf("fake rpc: bad raw tx: %v", err)
			return nil
		}
		f.mu.Lock()
		f.sent = append(f.sent, tx)
		f.mu.Unlock()
		return tx.Hash()
	case "eth_getTransactionByHash":
		var h common.Hash
		_ = json.Unmarshal(params[0], &h)
		f.mu.Lock()
		defer f.mu.Unlock()
		if tx, ok := f.known[h]; ok {
			return minedTxJSON(f.t, tx)
		}
		if f.knowsSent {
			for _, tx := range f.sent {
				if tx.Hash() == h {
					return tx
				}
			}
		}
		return nil
	case "eth_getTransactionReceipt":
		var h common.Hash
		_ = json.Unmarshal(params[0], &h)
		f.mu.Lock()
		status, gas, cumul := types.ReceiptStatusSuccessful, uint64(21000), uint64(21000)
		if f.receiptFailed {
			status = types.ReceiptStatusFailed
		}
		if f.receiptGas != 0 {
			gas, cumul = f.receiptGas, f.receiptCumul
		}
		f.mu.Unlock()
		return &types.Receipt{
			Status:            status,
			CumulativeGasUsed: cumul,
			GasUsed:           gas,
			Logs:              []*types.Log{},
			TxHash:            h,
			BlockHash:         common.Hash{2},
			BlockNumber:       big.NewInt(16),
		}
	}
	f.t.Errorf("fake rpc: unexpected method %s", method)
	return nil
}

// managerOn returns a manager whose tx/call endpoint is the fake node.
func managerOn(t *testing.T, f *fakeRPC) *ContractManage {
	t.Helper()
	sk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	return &ContractManage{
		RPC:        f.srv.URL,
		rpcs:       []string{f.srv.URL},
		ChainID:    big.NewInt(31337),
		sk:         sk,
		PieceAddr:  common.HexToAddress("0x1001"),
		EProofAddr: common.HexToAddress("0x1002"),
		TokenAddr:  common.HexToAddress("0x1003"),
	}
}

func hasSubstr(err error, s string) bool {
	return err != nil && strings.Contains(err.Error(), s)
}

// minedTxJSON is tx as eth_getTransactionByHash returns a mined tx.
func minedTxJSON(t *testing.T, tx *types.Transaction) map[string]interface{} {
	b, err := tx.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	m := map[string]interface{}{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	m["blockNumber"] = "0x10"
	m["blockHash"] = common.Hash{2}
	m["transactionIndex"] = "0x0"
	return m
}
