// example/upload — reference client for DIRECT DA upload (client → stream → chain).
//
// The lean, SDK-direct path (no hub): discover a stream via the gateway (or pin
// one with -stream), upload the file/dir straight to that stream, then the
// client's own wallet submits AddPiece on-chain. This is best for one-off
// files/dirs (fewest hops, no hub copy, user owns the piece). Continuous small
// writes (agent memory) should use the hub instead (batching + object store).
//
//	go run ./example/upload -sk <hexkey> -path ./weights.bin \
//	    -gateway http://<gateway> [-stream 0x<streamAddr>] [-name myfile]
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
	contract "github.com/unibaseio/da-sdk-go/contract/v2"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
)

func main() {
	gateway := flag.String("gateway", envOr("UNIBASE_GATEWAY", ""), "gateway URL (stream discovery + metadata)")
	stream := flag.String("stream", "", "pin a stream node address (else the gateway picks an online one)")
	skstr := flag.String("sk", "", "funded wallet private key (hex): holds ETH gas + UB bond")
	chain := flag.String("chain", envOr("CHAIN_TYPE", "base-sepolia"), "chain type")
	path := flag.String("path", "", "file or directory to upload")
	name := flag.String("name", "", "public object name for a single file (default: file basename)")
	n := flag.Int("n", 6, "erasure-code N")
	kk := flag.Int("k", 4, "erasure-code K")
	flag.Parse()

	if *gateway == "" || *skstr == "" || *path == "" {
		log.Fatal("need -gateway, -sk, -path")
	}
	sk, err := crypto.HexToECDSA(strings.TrimPrefix(*skstr, "0x"))
	if err != nil {
		log.Fatalf("bad -sk: %v", err)
	}
	// -stream pins a preferred stream (SDK reads STREAM_PRIORITY).
	if *stream != "" {
		os.Setenv("STREAM_PRIORITY", *stream)
	}

	au, err := key.BuildAuth(sk, []byte("upload"))
	if err != nil {
		log.Fatal(err)
	}
	sdk.Login(*gateway, au) // optional server-side charge/registration

	cm, err := contract.NewContractManage(sk, *chain)
	if err != nil {
		log.Fatal(err)
	}
	if err := cm.CheckBalance(au.Addr); err != nil {
		log.Fatalf("wallet %s: %v (needs ETH gas + UB)", au.Addr.Hex(), err)
	}
	policy := types.Policy{N: uint8(*n), K: uint8(*kk)}

	fp, _ := homedir.Expand(*path)
	fi, err := os.Stat(fp)
	if err != nil {
		log.Fatal(err)
	}
	if !fi.IsDir() {
		nm := *name
		if nm == "" {
			nm = filepath.Base(fp)
		}
		must(uploadOne(cm, au, policy, *gateway, fp, nm))
		return
	}
	// directory: upload each file under its basename
	err = filepath.Walk(fp, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		return uploadOne(cm, au, policy, *gateway, p, filepath.Base(p))
	})
	must(err)
}

func uploadOne(cm *contract.ContractManage, au types.Auth, policy types.Policy, gateway, fp, nm string) error {
	ff, streamer, err := sdk.Upload(gateway, au, policy, fp, nm) // → stream (direct) + UploadFileMeta(gateway)
	if err != nil {
		return err
	}
	pcs, err := sdk.CheckFileFull(ff, streamer, fp) // trustless: verify the streamer encoded the real bytes
	if err != nil {
		return err
	}
	for _, pc := range pcs {
		if _, err := cm.AddPiece(pc); err != nil { // client's own wallet stakes + submits on-chain
			return err
		}
	}
	fmt.Printf("uploaded %s → name=%s streamer=%s sha256=%s pieces=%d\n", fp, ff.Name, streamer.Hex(), ff.Hash, len(pcs))
	return nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
