// example/download — reference client for the hub /v1 read API (S3-shaped).
// See da/HUB_API_V1_SPEC.md.
//
// Reads are public (content is client-encrypted), so no signature is needed:
//
//	GET /v1/buckets/{bucket}/objects/{key}          (metadata + verifiable receipt)
//	GET /v1/buckets/{bucket}/objects/{key}/content  (raw bytes)
//	GET /v1/buckets/{bucket}/objects/{key}/proof    (verification bundle)
//
//	go run ./example/download -bucket my-models -key weights.bin \
//	    -owner 0x... -out ./weights.bin -hub http://127.0.0.1:8086
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
)

func main() {
	hub := flag.String("hub", envOr("HUB_URL", "http://127.0.0.1:8086"), "hub base URL")
	bucket := flag.String("bucket", "", "bucket (namespace)")
	key := flag.String("key", "", "object key")
	owner := flag.String("owner", "", "owner address (scopes the lookup; optional)")
	out := flag.String("out", "", "output file (default: stdout summary only)")
	flag.Parse()

	if *bucket == "" || *key == "" {
		log.Fatal("need -bucket and -key")
	}

	client := &http.Client{Timeout: 120 * time.Second}
	base := *hub + "/v1/buckets/" + *bucket + "/objects/" + url.PathEscape(*key)
	q := ""
	if *owner != "" {
		q = "?owner=" + *owner
	}

	get := func(u string) ([]byte, int) {
		resp, err := client.Get(u)
		if err != nil {
			log.Fatalf("GET %s: %v", u, err)
		}
		defer resp.Body.Close()
		b, _ := io.ReadAll(resp.Body)
		return b, resp.StatusCode
	}

	// 1) receipt (metadata: size, sha256, commitment, chain, status)
	meta, code := get(base + q)
	fmt.Printf("GET object -> %d\n  %s\n", code, string(meta))

	// 2) verification bundle
	proof, pcode := get(base + "/proof" + q)
	fmt.Printf("GET proof -> %d\n  %s\n", pcode, string(proof))

	// 3) content
	body, ccode := get(base + "/content" + q)
	if ccode != http.StatusOK {
		fmt.Printf("GET content -> %d\n  %s\n", ccode, string(body))
		return
	}
	if *out == "" {
		fmt.Printf("GET content -> %d (%d bytes; pass -out to save)\n", ccode, len(body))
		return
	}
	fp, _ := homedir.Expand(*out)
	if err := os.WriteFile(fp, body, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GET content -> %d, wrote %d bytes to %s\n", ccode, len(body), fp)
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
