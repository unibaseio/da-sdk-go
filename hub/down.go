package hub

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/unibaseio/da-sdk-go/build"
	lerror "github.com/unibaseio/da-sdk-go/lib/error"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
)

func (s *Server) addDownload(g *gin.RouterGroup) {
	g.Group("/").POST("/download", s.downloadByPOST)
	g.Group("/").GET("/download", s.downloadByGET)
}

func (s *Server) downloadByGET(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.Query("name")
	if mn == "" {
		mn = c.Query("id")
	}
	addr, ok := ResolveOwnerForList(c, c.Query("owner"))
	if !ok {
		return
	}

	head := fmt.Sprintf("attachment; filename=\"%s\"", mn)
	extraHeaders := map[string]string{
		"Content-Disposition": head,
	}

	var w bytes.Buffer
	size, err := s.download(ctx, mn, addr, &w)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.DataFromReader(http.StatusOK, size, "text/plain", &w, extraHeaders)
}

func (s *Server) downloadByPOST(c *gin.Context) {
	ctx := c.Request.Context()
	mn := c.PostForm("name")
	if mn == "" {
		mn = c.PostForm("id")
	}
	addr, ok := ResolveOwnerForList(c, c.PostForm("owner"))
	if !ok {
		return
	}

	head := fmt.Sprintf("attachment; filename=\"%s\"", mn)
	extraHeaders := map[string]string{
		"Content-Disposition": head,
	}

	var w bytes.Buffer
	size, err := s.download(ctx, mn, addr, &w)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}
	c.DataFromReader(http.StatusOK, size, "text/plain", &w, extraHeaders)
}

func (s *Server) download(ctx context.Context, name, owner string, w io.Writer) (int64, error) {
	size, _, err := s.logFSRead(owner, name, w)
	if err == nil {
		return size, nil
	}

	// Local miss. If this key was very recently confirmed absent, skip the
	// expensive remote fallback chain below — this absorbs download floods for
	// keys that don't exist (each full miss otherwise costs remote round-trips).
	if s.missCache.has(owner, name) {
		return 0, fmt.Errorf("no such file: %s at %s", name, owner)
	}

	// Cold-but-existing keys: dedupe the expensive DA reconstruct so N concurrent
	// requests for the same object trigger one fetch, shared to all. The result
	// is buffered (as downloadByGET already does) and populated into the hot read
	// cache; large objects are excluded by the cache's own size guard.
	s.dlTotal.Add(1)
	v, err, shared := s.dlSF.Do(missKey(owner, name), func() (interface{}, error) {
		// Optionally cap concurrent distinct-key reconstructs (HUB_DOWNLOAD_CONCURRENCY);
		// nil semaphore = unlimited. Only the flight leader acquires — sharers ride
		// its result — so the bound counts real DA fetches, not waiters.
		if s.dlSem != nil {
			select {
			case s.dlSem <- struct{}{}:
				defer func() { <-s.dlSem }()
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		var buf bytes.Buffer
		if derr := s.downloadRemote(ctx, name, owner, &buf); derr != nil {
			return nil, derr
		}
		b := buf.Bytes()
		s.readCache.put(owner, name, b)
		return b, nil
	})
	if shared {
		s.dlShared.Add(1)
	}
	if err != nil {
		return 0, err
	}
	data := v.([]byte)
	n, werr := w.Write(data)
	if werr != nil {
		return 0, werr
	}
	return int64(n), nil
}

// downloadRemote runs the DA fallback chain (file → piece → replica) into w.
// On a total miss it records a negative-cache marker and returns an error.
func (s *Server) downloadRemote(ctx context.Context, name, owner string, w io.Writer) error {
	// a file record is looked up as the caller's own when the caller is known
	// (names are per owner), and the rebuilt file is checked against its hash
	fowner := types.EmptyAddr
	if common.IsHexAddress(owner) {
		fowner = common.HexToAddress(owner)
	}
	if _, err := sdk.GetFileReceiptOf(build.ServerURL, name, fowner); err == nil {
		// own buffer: a failed attempt (e.g. hash mismatch) has written partial
		// bytes, which must not prefix the next fallback's output
		var fb bytes.Buffer
		if err = sdk.DownloadOf(build.ServerURL, s.auth, name, fowner, s.ps, &fb); err == nil {
			_, err = w.Write(fb.Bytes())
			return err
		}
		logger.Warnf("download file %s (owner %s) failed: %s", name, owner, err)
	}

	if _, err := sdk.GetPieceReceipt(build.ServerURL, s.auth, name); err == nil {
		if err = sdk.DownloadPieceAndSave(build.ServerURL, s.auth, name, s.ps); err != nil {
			return err
		}
		if _, err := s.ps.GetPiece(ctx, name, w, types.Options{}); err == nil {
			return nil
		}
	}

	if _, err := s.ps.GetReplica(ctx, name, w, types.Options{}); err == nil {
		return nil
	}

	// Fully missing everywhere — remember so repeat requests stay cheap.
	s.missCache.add(owner, name)
	return fmt.Errorf("no such file: %s at %s", name, owner)
}

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
