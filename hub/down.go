package hub

import (
	"bytes"
	"context"
	"errors"
	"time"

	"github.com/unibaseio/da-sdk-go/build"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
)

// defaultPieceDownloadConcurrency bounds concurrent public piece reads and
// rebuilds (HUB_PIECE_DOWNLOAD_CONCURRENCY). It does not bound memory held
// while responses are written: the slot is released before the bytes go out.
const defaultPieceDownloadConcurrency = 4

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
// piece store when it holds it, else rebuilt from store nodes. The read (not
// the response write) holds a pieceSem slot (the endpoint is public and a
// piece may be ~1 GB), a rebuild also a dlSem slot
// like other DA reconstructs, and concurrent requests for one piece share a
// single rebuild.
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
	release, err := acquireSem(ctx, s.pieceSem)
	if err != nil {
		return nil, err
	}
	defer release()

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
