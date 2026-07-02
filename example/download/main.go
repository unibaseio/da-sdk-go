// example/download — reference client for DIRECT DA download (gateway → pieces).
//
// Lean, SDK-direct path (no hub): fetch the file receipt from the gateway and
// reconstruct the bytes from its DA pieces (store replicas via the stream WS
// relay, or the streamer's staging copy). No funds needed — but the request
// still carries a signed Authorization (an ephemeral key is fine).
//
//	go run ./example/download -gateway http://<gateway> -name <fileName> -out ./out.bin
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/sdk"
)

func main() {
	gateway := flag.String("gateway", envOr("UNIBASE_GATEWAY", ""), "gateway URL")
	name := flag.String("name", "", "file name (as registered on upload)")
	out := flag.String("out", "", "output file; '-' or empty writes stdout")
	skstr := flag.String("sk", "", "signing key (hex); optional — ephemeral if empty (no funds needed)")
	flag.Parse()

	if *gateway == "" || *name == "" {
		log.Fatal("need -gateway and -name")
	}

	// reads need a valid signed Authorization but no funds; ephemeral key is fine.
	sk, err := crypto.HexToECDSA(strings.TrimPrefix(*skstr, "0x"))
	if err != nil {
		if sk, err = crypto.GenerateKey(); err != nil {
			log.Fatal(err)
		}
	}
	au, err := key.BuildAuth(sk, []byte("download"))
	if err != nil {
		log.Fatal(err)
	}

	w := os.Stdout
	if *out != "" && *out != "-" {
		fp, _ := homedir.Expand(*out)
		f, err := os.Create(fp)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	}
	if err := sdk.Download(*gateway, au, *name, nil, w); err != nil {
		log.Fatal(err)
	}
	if w != os.Stdout {
		fmt.Fprintf(os.Stderr, "downloaded %s → %s\n", *name, w.Name())
	}
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
