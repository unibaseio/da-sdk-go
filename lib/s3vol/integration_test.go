package s3vol

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/minio/minio-go/v7"
)

// TestS3RoundTripIntegration exercises the real PutVolume/RangeAt against a live
// S3-compatible endpoint (a local MinIO container). It is skipped unless
// HUB_S3_TEST_ENDPOINT is set, so `go test` stays hermetic by default.
//
// Run locally with:
//
//	docker run -d -p 9000:9000 -e MINIO_ROOT_USER=minioadmin \
//	  -e MINIO_ROOT_PASSWORD=minioadmin minio/minio server /data
//	HUB_S3_TEST_ENDPOINT=localhost:9000 HUB_S3_TEST_BUCKET=hub-vol \
//	  HUB_S3_TEST_AK=minioadmin HUB_S3_TEST_SK=minioadmin \
//	  go test ./lib/s3vol/ -run Integration -v
func TestS3RoundTripIntegration(t *testing.T) {
	endpoint := os.Getenv("HUB_S3_TEST_ENDPOINT")
	if endpoint == "" {
		t.Skip("set HUB_S3_TEST_ENDPOINT (+ _BUCKET/_AK/_SK) to run the S3 integration test")
	}
	bucket := envOr("HUB_S3_TEST_BUCKET", "hub-vol")

	t.Setenv("HUB_S3_ENDPOINT", endpoint)
	t.Setenv("HUB_S3_BUCKET", bucket)
	t.Setenv("HUB_S3_ACCESS_KEY", os.Getenv("HUB_S3_TEST_AK"))
	t.Setenv("HUB_S3_SECRET_KEY", os.Getenv("HUB_S3_TEST_SK"))
	t.Setenv("HUB_S3_USE_SSL", envOr("HUB_S3_TEST_SSL", "false"))

	st, err := NewFromEnv()
	if err != nil {
		t.Fatalf("NewFromEnv: %v", err)
	}
	ctx := context.Background()

	// ensure the bucket exists
	if ok, berr := st.cli.BucketExists(ctx, bucket); berr != nil {
		t.Fatalf("BucketExists: %v", berr)
	} else if !ok {
		if merr := st.cli.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); merr != nil {
			t.Fatalf("MakeBucket: %v", merr)
		}
	}
	if perr := st.Ping(ctx); perr != nil {
		t.Fatalf("Ping: %v", perr)
	}

	// write a fake "volume" file with recognizable bytes
	vol := make([]byte, 4096)
	for i := range vol {
		vol[i] = byte(i)
	}
	dir := t.TempDir()
	path := filepath.Join(dir, "42.vol")
	if err := os.WriteFile(path, vol, 0644); err != nil {
		t.Fatal(err)
	}

	be := st.Bind("0xlocaltest", "0xownertest")
	const idx = 42
	if err := be.PutVolume(idx, path); err != nil {
		t.Fatalf("PutVolume: %v", err)
	}

	// full-object range
	got, err := be.RangeAt(idx, 0, int64(len(vol)))
	if err != nil {
		t.Fatalf("RangeAt full: %v", err)
	}
	if !bytes.Equal(got, vol) {
		t.Fatalf("full range mismatch (%d vs %d)", len(got), len(vol))
	}

	// sub-range (a needle inside the volume)
	off, n := int64(1000), int64(256)
	sub, err := be.RangeAt(idx, off, n)
	if err != nil {
		t.Fatalf("RangeAt sub: %v", err)
	}
	if !bytes.Equal(sub, vol[off:off+n]) {
		t.Fatalf("sub range mismatch")
	}

	// cleanup
	_ = st.cli.RemoveObject(ctx, bucket, be.(*ownerBackend).key(idx), minio.RemoveObjectOptions{})
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
