package hub

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/unibaseio/da-sdk-go/build"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
)

// shardFetchTimeout bounds one cross-shard object fetch (small conversation
// records; a slow home shard must not stall the whole listing).
const shardFetchTimeout = 15 * time.Second

// readNeedleFallback reads an index row's object when this hub's LogFS doesn't
// have it, bound to that row's owner throughout (review 2026-10-04 M). It
// replaced the by-name download() fallback (since removed), which on a non-home shard ran the
// owner-less piece/replica lookups and cached whatever came back under the
// owner's key — in the shared L2 too.
//
//  1. Staged objects live only on the owner's home shard: fetch this exact
//     (owner, bucket, key) from there and keep it only if its size is the row's.
//     It is cached under the row's location, as a local read would be.
//  2. Otherwise only a DA file record the owner registered under this name
//     (hash-checked by sdk.DownloadOf, size-checked here) — never a piece or
//     replica looked up by bare name. Not cached.
func (s *Server) readNeedleFallback(ctx context.Context, n types.Needle, w io.Writer) error {
	data, err := s.shardFetchObject(ctx, n)
	if err == nil {
		s.readCache.put(n.Owner, locKey(n.File, n.Start, n.Size), data)
		_, err = w.Write(data)
		return err
	}
	if !common.IsHexAddress(n.Owner) {
		return err
	}
	s.dlTotal.Add(1)
	v, ferr, shared := s.dlSF.Do("file:"+cacheKey(n.Owner, n.Name), func() (interface{}, error) {
		if s.dlSem != nil {
			select {
			case s.dlSem <- struct{}{}:
				defer func() { <-s.dlSem }()
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		var buf bytes.Buffer
		if err := fetchOwnerFile(s, n.Name, common.HexToAddress(n.Owner), &buf); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	})
	if shared {
		s.dlShared.Add(1)
	}
	if ferr != nil {
		return ferr
	}
	b := v.([]byte)
	if uint64(len(b)) != n.Size {
		return fmt.Errorf("file record %s of %s has %d bytes, row has %d", n.Name, n.Owner, len(b), n.Size)
	}
	_, err = w.Write(b)
	return err
}

// fetchOwnerFile downloads the DA file record that owner registered under name
// (a variable so tests can stub the network).
var fetchOwnerFile = func(s *Server, name string, owner common.Address, w io.Writer) error {
	if _, err := sdk.GetFileReceiptOf(build.ServerURL, name, owner); err != nil {
		return err
	}
	return sdk.DownloadOf(build.ServerURL, s.auth, name, owner, s.ps, w)
}

// shardFetchObject fetches the row's object from its owner's home shard (the
// /v1 content endpoint, marked as forwarded so the peer answers locally). An
// error when sharding is off, this hub is home, or the answer isn't the row's
// size.
func (s *Server) shardFetchObject(ctx context.Context, n types.Needle) ([]byte, error) {
	sr := s.shard
	if sr == nil || n.Owner == "" || n.Bucket == "" {
		return nil, fmt.Errorf("no home shard to ask")
	}
	home := sr.shardOf(n.Owner)
	if home == sr.index || sr.peers[home] == nil {
		return nil, fmt.Errorf("this hub is the home shard")
	}
	target := strings.TrimRight(sr.peers[home].String(), "/") + "/v1/buckets/" + url.PathEscape(n.Bucket) + "/objects/" + url.PathEscape(n.Name) +
		"/content?owner=" + url.QueryEscape(n.Owner)

	ctx, cancel := context.WithTimeout(ctx, shardFetchTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}
	sr.markForwarded(req)
	resp, err := sr.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("home shard %d: status %d", home, resp.StatusCode)
	}
	data, err := io.ReadAll(io.LimitReader(resp.Body, int64(n.Size)+1))
	if err != nil {
		return nil, err
	}
	if uint64(len(data)) != n.Size {
		return nil, fmt.Errorf("home shard %d returned %d bytes, row has %d", home, len(data), n.Size)
	}
	sr.readProxied.Add(1)
	return data, nil
}
