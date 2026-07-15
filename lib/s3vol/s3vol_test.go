package s3vol

import (
	"testing"
)

// TestNewFromEnvRequiresEndpointAndBucket: config guard — s3 buffer must not
// silently start without a target.
func TestNewFromEnvRequiresEndpointAndBucket(t *testing.T) {
	t.Setenv("HUB_S3_ENDPOINT", "")
	t.Setenv("HUB_S3_BUCKET", "")
	if _, err := NewFromEnv(); err == nil {
		t.Fatal("expected error when endpoint+bucket unset")
	}

	t.Setenv("HUB_S3_ENDPOINT", "localhost:9000")
	t.Setenv("HUB_S3_BUCKET", "")
	if _, err := NewFromEnv(); err == nil {
		t.Fatal("expected error when bucket unset")
	}
}

// TestNewFromEnvOKAndBindKey: a valid config builds a client and Bind produces
// owner-namespaced, per-volume object keys.
func TestNewFromEnvOKAndBindKey(t *testing.T) {
	t.Setenv("HUB_S3_ENDPOINT", "localhost:9000")
	t.Setenv("HUB_S3_BUCKET", "hub-vol")
	t.Setenv("HUB_S3_ACCESS_KEY", "ak")
	t.Setenv("HUB_S3_SECRET_KEY", "sk")
	t.Setenv("HUB_S3_USE_SSL", "false")

	st, err := NewFromEnv()
	if err != nil {
		t.Fatalf("NewFromEnv: %v", err)
	}

	b := st.Bind("0xlocalnode", "0xowner").(*ownerBackend)
	if got, want := b.key(7), "vol/0xlocalnode/0xowner/7.vol"; got != want {
		t.Fatalf("key = %q, want %q", got, want)
	}
	// distinct owners never collide
	b2 := st.Bind("0xlocalnode", "0xother").(*ownerBackend)
	if b.key(7) == b2.key(7) {
		t.Fatal("keys must differ per owner")
	}
}
