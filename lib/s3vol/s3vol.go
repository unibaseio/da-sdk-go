// Package s3vol is the S3/MinIO VolumeBackend for logfs (P3-S). It stores each
// SEALED 31MiB volume as one object and serves a needle as a byte-range GET, so
// 34M needles cost O(volumes) objects rather than O(needles) — see
// da/HUB_CACHE_OPTIMIZATION.md §8.6.
//
// It is S3-compatible: point HUB_S3_ENDPOINT at AWS S3, a self-hosted MinIO, or
// any S3 gateway. Only the hub wires this in (behind HUB_BUFFER=s3); logfs itself
// stays dependency-free and local-only by default.
package s3vol

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/unibaseio/da-sdk-go/lib/env"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
)

// Store is a shared S3/MinIO client. Per-owner backends (Bind) namespace their
// objects by owner so one bucket serves all owners of one hub cluster.
type Store struct {
	cli    *minio.Client
	bucket string
}

// NewFromEnv builds a Store from HUB_S3_* env, or returns an error if the
// required endpoint/bucket are unset. Caller gates on HUB_BUFFER=s3.
//
//	HUB_S3_ENDPOINT    host[:port] (e.g. s3.ap-southeast-1.amazonaws.com, localhost:9000)
//	HUB_S3_BUCKET      bucket name (must already exist)
//	HUB_S3_ACCESS_KEY  / HUB_S3_SECRET_KEY  static credentials; if ACCESS_KEY is
//	                   empty, fall back to the EC2 instance IAM role (IMDS) —
//	                   the preferred, keyless path on AWS.
//	HUB_S3_REGION      default us-east-1
//	HUB_S3_USE_SSL     "false" to disable TLS (local MinIO); default on
func NewFromEnv() (*Store, error) {
	endpoint := env.Str("HUB_S3_ENDPOINT", "")
	bucket := env.Str("HUB_S3_BUCKET", "")
	if endpoint == "" || bucket == "" {
		return nil, fmt.Errorf("HUB_S3_ENDPOINT and HUB_S3_BUCKET are required for HUB_BUFFER=s3")
	}
	ak := env.Str("HUB_S3_ACCESS_KEY", "")
	sk := env.Str("HUB_S3_SECRET_KEY", "")
	region := env.Str("HUB_S3_REGION", "us-east-1")
	useSSL := env.Str("HUB_S3_USE_SSL", "true") != "false"

	// Static keys when provided (MinIO / non-AWS); otherwise the EC2 instance
	// IAM role via IMDS — no long-lived secrets on the box.
	var creds *credentials.Credentials
	if ak != "" {
		creds = credentials.NewStaticV4(ak, sk, "")
	} else {
		creds = credentials.NewIAM("")
	}

	cli, err := minio.New(endpoint, &minio.Options{
		Creds:  creds,
		Secure: useSSL,
		Region: region,
	})
	if err != nil {
		return nil, fmt.Errorf("s3vol: init client: %w", err)
	}
	return &Store{cli: cli, bucket: bucket}, nil
}

// Ping verifies the bucket is reachable; used for a startup log line so a
// misconfiguration surfaces early (reads/writes still degrade gracefully).
func (s *Store) Ping(ctx context.Context) error {
	ok, err := s.cli.BucketExists(ctx, s.bucket)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("s3vol: bucket %q not found", s.bucket)
	}
	return nil
}

// Bind returns a per-owner logfs.VolumeBackend. Objects are keyed
// vol/<local>/<addr>/<idx>.vol so a cluster's hubs share one bucket without
// collision and a reader replica can locate any writer's sealed volume.
func (s *Store) Bind(local, addr string) logfs.VolumeBackend {
	return &ownerBackend{s: s, prefix: fmt.Sprintf("vol/%s/%s/", local, addr)}
}

type ownerBackend struct {
	s      *Store
	prefix string
}

func (b *ownerBackend) key(idx uint64) string {
	return fmt.Sprintf("%s%d.vol", b.prefix, idx)
}

// PutVolume streams the sealed volume file to S3 as one object. Idempotent:
// volumes are append-only, so re-PUT of the same idx is the same bytes.
func (b *ownerBackend) PutVolume(idx uint64, localPath string) error {
	f, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	_, err = b.s.cli.PutObject(ctx, b.s.bucket, b.key(idx), f, fi.Size(), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return fmt.Errorf("s3vol: put %s: %w", b.key(idx), err)
	}
	return nil
}

// RangeAt reads n bytes at off from the volume object (a needle read).
func (b *ownerBackend) RangeAt(idx uint64, off, n int64) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := minio.GetObjectOptions{}
	if err := opts.SetRange(off, off+n-1); err != nil { // inclusive end
		return nil, err
	}
	obj, err := b.s.cli.GetObject(ctx, b.s.bucket, b.key(idx), opts)
	if err != nil {
		return nil, fmt.Errorf("s3vol: get %s: %w", b.key(idx), err)
	}
	defer obj.Close()

	buf := make([]byte, n)
	got, err := io.ReadFull(obj, buf)
	if err != nil && err != io.ErrUnexpectedEOF {
		return nil, fmt.Errorf("s3vol: read %s: %w", b.key(idx), err)
	}
	return buf[:got], nil
}
