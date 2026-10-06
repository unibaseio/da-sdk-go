package hub

import (
	"bytes"
	"context"
	"errors"
	"time"

	"github.com/unibaseio/da-sdk-go/build"
	"github.com/unibaseio/da-sdk-go/lib/env"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
)

// defaultPieceDownloadConcurrency bounds concurrent public piece reads,
// rebuilds and response writes (HUB_PIECE_DOWNLOAD_CONCURRENCY): a slot is
// held until the piece has been written out, so at most this many pieces are
// in memory for the endpoint.
const defaultPieceDownloadConcurrency = 4

// A piece response must be written within base + size/minRate
// (HUB_PIECE_WRITE_BASE_SEC, HUB_PIECE_WRITE_MIN_BPS): the hub has no server
// WriteTimeout (large object reads), so without this a client that stops
// reading would hold its slot and the piece's buffer forever. 1 GiB at the
// default 1 MiB/s gets about 17 minutes.
const (
	defaultPieceWriteBaseSec = 30
	defaultPieceWriteMinBPS  = 1 << 20
)

// pieceWriteTimeout is how long writing n bytes of a piece may take.
func pieceWriteTimeout(n int) time.Duration {
	base := time.Duration(env.Int64("HUB_PIECE_WRITE_BASE_SEC", defaultPieceWriteBaseSec)) * time.Second
	bps := env.Int64("HUB_PIECE_WRITE_MIN_BPS", defaultPieceWriteMinBPS)
	if bps <= 0 {
		bps = defaultPieceWriteMinBPS
	}
	return base + time.Duration(float64(n)/float64(bps)*float64(time.Second))
}

// errBusy is returned when a read gave up waiting for a download slot.
var errBusy = errors.New("too many downloads in progress; retry later")

// maxSlotWait bounds how long a read queues for a download slot (a var for tests).
var maxSlotWait = time.Minute

// acquireSem takes a slot of sem (nil = unbounded), giving up when ctx ends or
// after maxSlotWait.
func acquireSem(ctx context.Context, sem chan struct{}) (func(), error) {
	if sem == nil {
		return func() {}, nil
	}
	t := time.NewTimer(maxSlotWait)
	defer t.Stop()
	select {
	case sem <- struct{}{}:
		return func() { <-sem }, nil
	case <-ctx.Done():
		return nil, errBusy
	case <-t.C:
		return nil, errBusy
	}
}

// downloadPiece returns a committed piece by its DA commitment: from the local
// piece store when it holds it, else rebuilt from store nodes. The caller
// holds a pieceSem slot across the read and the response write (the endpoint
// is public and a piece may be ~1 GB); a rebuild also takes a dlSem slot like
// other DA reconstructs, and concurrent requests for one piece share a single
// rebuild.
//
// A rebuilt piece is NOT put into the piece store. A bare piece has nothing
// the hub can check it against (its name is a KZG commitment; checking bytes
// against it needs the full SRS), so caching it would let one lying store or
// stream plant bytes that every later reader of the store gets, including
// file reads (sdk.DownloadOf), which would then fail their hash. Cost: each
// cold read of a piece the hub has not checked rebuilds it (K replica fetches
// plus RS repair) instead of hitting local disk. That is bounded by pieceSem,
// dlSem and singleflight, and pieces the hub did check (fetched for a file
// whose hash matched) are still served from the store.
func (s *Server) downloadPiece(ctx context.Context, cid string) ([]byte, error) {
	var local bytes.Buffer
	if _, err := s.ps.GetPiece(ctx, cid, &local, types.Options{}); err == nil {
		return local.Bytes(), nil
	}
	v, err, _ := s.dlSF.Do("piece:"+cid, func() (interface{}, error) {
		release, err := acquireSem(ctx, s.dlSem)
		if err != nil {
			return nil, err
		}
		defer release()
		_, data, err := sdk.DownloadPiece(build.ServerURL, s.auth, cid)
		if err != nil {
			return nil, err
		}
		return data, nil
	})
	if err != nil {
		return nil, err
	}
	return v.([]byte), nil
}
