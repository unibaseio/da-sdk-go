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

// defaultPieceDownloadConcurrency bounds concurrent public piece reads
// (HUB_PIECE_DOWNLOAD_CONCURRENCY): each may hold a piece of up to ~1 GB.
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
// piece store, else rebuilt from store nodes (then kept locally). The whole read
// holds a pieceSem slot (the endpoint is public and a piece may be ~1 GB), a
// rebuild also a dlSem slot like other DA reconstructs, and concurrent requests
// for one piece share a single rebuild. The bytes are returned, not copied into
// a writer, so a piece is held once per flight.
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
		if _, err := sdk.GetPieceReceipt(build.ServerURL, s.auth, cid); err != nil {
			return nil, err
		}
		if err := sdk.DownloadPieceAndSave(build.ServerURL, s.auth, cid, s.ps); err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		if _, err := s.ps.GetPiece(ctx, cid, &buf, types.Options{}); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	})
	if err != nil {
		return nil, err
	}
	return v.([]byte), nil
}
