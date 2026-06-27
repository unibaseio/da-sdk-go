package contract

import (
	"bytes"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	piece "github.com/unibaseio/da-sdk-go/contract/v2/go/piece"
)

// TestDecodeAddPieceFields verifies the selector dispatch in decodeAddPieceFields:
// addPiece and addPieceFor must yield identical piece fields, with addPieceFor's
// leading `owner` arg correctly dropped (the bug we guard against is decoding
// addPieceFor calldata with addPiece's layout, which shifts every field).
func TestDecodeAddPieceFields(t *testing.T) {
	cabi, err := abi.JSON(strings.NewReader(piece.PieceMetaData.ABI))
	if err != nil {
		t.Fatalf("load piece abi: %v", err)
	}

	pn := []byte("piece-commitment-bytes")
	price := big.NewInt(1_000_000)
	size := uint64(64 * 1024 * 1024)
	expire := uint64(4242)
	rsn := uint8(6)
	rsk := uint8(4)
	streamer := common.HexToAddress("0x1111111111111111111111111111111111111111")
	owner := common.HexToAddress("0x2222222222222222222222222222222222222222")

	assertFields := func(t *testing.T, fields []interface{}) {
		t.Helper()
		if !bytes.Equal(fields[0].([]byte), pn) {
			t.Errorf("pn mismatch: %x", fields[0].([]byte))
		}
		if fields[1].(*big.Int).Cmp(price) != 0 {
			t.Errorf("price mismatch: %s", fields[1].(*big.Int))
		}
		if fields[2].(uint64) != size {
			t.Errorf("size mismatch: %d", fields[2].(uint64))
		}
		if fields[3].(uint64) != expire {
			t.Errorf("expire mismatch: %d", fields[3].(uint64))
		}
		if fields[4].(uint8) != rsn {
			t.Errorf("rsn mismatch: %d", fields[4].(uint8))
		}
		if fields[5].(uint8) != rsk {
			t.Errorf("rsk mismatch: %d", fields[5].(uint8))
		}
		if fields[6].(common.Address) != streamer {
			t.Errorf("streamer mismatch: %s", fields[6].(common.Address))
		}
	}

	t.Run("addPiece", func(t *testing.T) {
		data, err := cabi.Pack("addPiece", pn, price, size, expire, rsn, rsk, streamer)
		if err != nil {
			t.Fatalf("pack addPiece: %v", err)
		}
		fields, err := decodeAddPieceFields(cabi, data)
		if err != nil {
			t.Fatalf("decode addPiece: %v", err)
		}
		assertFields(t, fields)
	})

	t.Run("addPieceFor drops leading owner", func(t *testing.T) {
		data, err := cabi.Pack("addPieceFor", owner, pn, price, size, expire, rsn, rsk, streamer)
		if err != nil {
			t.Fatalf("pack addPieceFor: %v", err)
		}
		fields, err := decodeAddPieceFields(cabi, data)
		if err != nil {
			t.Fatalf("decode addPieceFor: %v", err)
		}
		// must be the SAME 7 fields as addPiece — owner shifted off
		assertFields(t, fields)
	})

	t.Run("unknown selector errors", func(t *testing.T) {
		if _, err := decodeAddPieceFields(cabi, []byte{0xde, 0xad, 0xbe, 0xef, 0x00}); err == nil {
			t.Fatal("expected error for unknown selector")
		}
		if _, err := decodeAddPieceFields(cabi, []byte{0x01, 0x02}); err == nil {
			t.Fatal("expected error for short calldata")
		}
	})
}
