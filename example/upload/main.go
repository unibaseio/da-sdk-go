// example/upload — reference client for the hub /v1 upload API (S3-shaped,
// wallet-native). See da/HUB_API_V1_SPEC.md.
//
// Flow:  PUT /v1/buckets/{bucket} {kind}                 (declare bucket + kind once)
//        PUT /v1/buckets/{bucket}/objects/{key}?wait=1    (upload one object)
//
// Auth: every write carries Authorization = hex(json(BuildAuth(...))); the
// recovered signer must equal the object owner (the bucket's owner).
//
//	go run ./example/upload -sk <hexkey> -bucket my-models -kind model \
//	    -path ./weights.bin -hub http://127.0.0.1:8086 -wait
package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
	"github.com/unibaseio/da-sdk-go/sdk"
)

func main() {
	hub := flag.String("hub", envOr("HUB_URL", "http://127.0.0.1:8086"), "hub base URL")
	skstr := flag.String("sk", "", "owner private key (hex); generated if empty")
	bucket := flag.String("bucket", "", "bucket (namespace) to upload into")
	kind := flag.String("kind", "file", "bucket kind: memory|knowledgebase|file|model|dataset")
	keyName := flag.String("key", "", "object key (default: file basename)")
	path := flag.String("path", "", "file to upload")
	wait := flag.Bool("wait", true, "block until the object is committed on-chain")
	flag.Parse()

	if *bucket == "" || *path == "" {
		log.Fatal("need -bucket and -path")
	}

	// owner key: the signer of every write; owner == this address.
	sk, err := crypto.HexToECDSA(strings.TrimPrefix(*skstr, "0x"))
	if err != nil {
		if sk, err = crypto.GenerateKey(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("=== generated key: %s ===\n", hex.EncodeToString(crypto.FromECDSA(sk)))
	}
	owner := crypto.PubkeyToAddress(sk.PublicKey).Hex()
	privk := hex.EncodeToString(crypto.FromECDSA(sk))

	// authHeader signs a fresh auth envelope (±10min window) per request.
	authHeader := func() string {
		au := sdk.BuildAuth(owner, privk, []byte("hub"))
		b, _ := json.Marshal(au)
		return hex.EncodeToString(b)
	}

	fp, _ := homedir.Expand(*path)
	data, err := os.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	key := *keyName
	if key == "" {
		key = filepath.Base(fp)
	}

	client := &http.Client{Timeout: 300 * time.Second}
	send := func(method, url, ctype, body string) {
		r, _ := http.NewRequest(method, url, strings.NewReader(body))
		r.Header.Set("Content-Type", ctype)
		r.Header.Set("Authorization", authHeader())
		resp, err := client.Do(r)
		if err != nil {
			log.Fatalf("%s %s: %v", method, url, err)
		}
		defer resp.Body.Close()
		rb, _ := io.ReadAll(resp.Body)
		fmt.Printf("%s %s -> %d\n  %s\n", method, url, resp.StatusCode, string(rb))
	}

	// 1) ensure the bucket exists with its kind (idempotent; kind immutable)
	send("PUT", *hub+"/v1/buckets/"+*bucket, "application/json", fmt.Sprintf(`{"kind":%q}`, *kind))

	// 2) upload the object; ?wait=1 blocks until the DA AddPiece lands on-chain
	//    (200 + committed receipt: commitment + chain), else 202 staged.
	url := *hub + "/v1/buckets/" + *bucket + "/objects/" + key
	if *wait {
		url += "?wait=1"
	}
	send("PUT", url, "application/octet-stream", string(data))
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
