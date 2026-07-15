package hub

import (
	"encoding/binary"
	"fmt"
	"path/filepath"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/env"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// getFS returns the per-owner LogFS instance, loading (and, on the write path,
// registering) it on demand.
//
// Concurrency: the common hit is a lock-free sync.Map load. On a miss, creation
// is deduplicated per-owner via a singleflight group, so logfs.New (a filesystem
// open) runs WITHOUT any server-wide lock — a brand-new owner no longer
// serializes unrelated owners (the old code held s.Lock() across logfs.New).
//
//   - create=true  (write path): loads or, for a never-seen owner, registers it
//     in the LOGINST registry + creates its account (idempotent).
//   - create=false (read path):  only loads owners already present in the
//     LOGINST registry; returns an error for unknown owners.
func (s *Server) getFS(addr string, create bool) (*logfs.LogFS, error) {
	if v, ok := s.lfs.Load(addr); ok {
		s.fsHit.Add(1)
		return v.(*logfs.LogFS), nil
	}

	v, err, _ := s.fsSF.Do(addr, func() (interface{}, error) {
		// Re-check: another goroutine may have created it between the fast-path
		// miss and acquiring the singleflight slot.
		if v, ok := s.lfs.Load(addr); ok {
			s.fsHit.Add(1)
			return v, nil
		}

		regKey := types.NewKey(types.DsLogFS, LOGINST, addr)
		has, herr := s.rp.MetaStore().Has(regKey)
		registered := herr == nil && has
		if !registered && !create {
			return nil, fmt.Errorf("no such owner: %s", addr)
		}

		fspath := filepath.Join(s.rp.Path(), LOGFS)
		fs, nerr := logfs.New(s.rp.MetaStore(), fspath, s.local.String(), addr, s.logFSOptions(addr)...)
		if nerr != nil {
			return nil, nerr
		}
		s.fsCreate.Add(1)

		if create {
			// account bookkeeping is idempotent; matches the old write path which
			// called addAccount on every in-memory (re)load of a writer.
			s.addAccount(addr)
			if !registered {
				s.registerOwner(addr)
			}
		}

		s.lfs.Store(addr, fs)
		return fs, nil
	})
	if err != nil {
		return nil, err
	}
	return v.(*logfs.LogFS), nil
}

// registerOwner assigns a never-seen owner a slot in the LOGINST registry
// (count, index->addr, addr marker) and bumps the owner count. Guarded by
// fscntMu; the per-owner singleflight in getFS ensures it runs at most once per
// new owner. Mirrors the original inline logic in logFSWrite.
func (s *Server) registerOwner(addr string) {
	s.fscntMu.Lock()
	defer s.fscntMu.Unlock()

	buf := make([]byte, 4)

	// addr marker: "this owner is registered"
	regKey := types.NewKey(types.DsLogFS, LOGINST, addr)
	binary.BigEndian.PutUint32(buf, 0)
	s.rp.MetaStore().Put(regKey, buf)

	// index -> addr (used by the drain fan-out to resolve owner from idx)
	idxKey := types.NewKey(types.DsLogFS, LOGINST, s.fscnt)
	s.rp.MetaStore().Put(idxKey, []byte(addr))

	// bump the owner count
	s.fscnt++
	cntKey := types.NewKey(types.DsLogFS, LOGINST)
	binary.BigEndian.PutUint32(buf, s.fscnt)
	s.rp.MetaStore().Put(cntKey, buf)

	logger.Infof("registered log inst: %s -> %d", addr, s.fscnt)
}

// logFSOptions returns the per-owner logfs construction options. When a durable
// volume backend is configured (HUB_BUFFER=s3), it binds an owner-namespaced
// backend so sealed volumes are uploaded and reads can fall back to it. Empty
// (local-only) by default.
func (s *Server) logFSOptions(addr string) []logfs.Option {
	if s.volStore == nil {
		return nil
	}
	opts := []logfs.Option{logfs.WithVolumeBackend(s.volStore.Bind(s.local.String(), addr))}
	// HUB_BUFFER_LOCAL_TTL (seconds): reclaim local disk for volumes confirmed
	// uploaded to S3, older than the TTL. 0 (default) = keep local forever.
	if ttl := env.Int("HUB_BUFFER_LOCAL_TTL", 0); ttl > 0 {
		opts = append(opts, logfs.WithLocalTTL(time.Duration(ttl)*time.Second))
	}
	return opts
}

// fscntGet returns the current registered-owner count (for the drain fan-out).
func (s *Server) fscntGet() uint32 {
	s.fscntMu.Lock()
	defer s.fscntMu.Unlock()
	return s.fscnt
}
