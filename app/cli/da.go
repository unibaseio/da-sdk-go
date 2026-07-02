package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
	contract "github.com/unibaseio/da-sdk-go/contract/v2"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
	"github.com/urfave/cli/v2"
)

// commonFlags are shared by every leaf command (per-command so they resolve
// from the subcommand context and work in any position).
func commonFlags(extra ...cli.Flag) []cli.Flag {
	base := []cli.Flag{
		&cli.StringFlag{Name: "gateway", EnvVars: []string{"UNIBASE_GATEWAY"}, Usage: "gateway URL (stream discovery + metadata)"},
		&cli.StringFlag{Name: "key", EnvVars: []string{"UNIBASE_KEY"}, Usage: "funded wallet private key (hex); required for upload"},
		&cli.BoolFlag{Name: "json", Usage: "machine-readable JSON output"},
	}
	return append(base, extra...)
}

// daCommand — `ubcli da`: DIRECT DA storage (client → stream → chain). The lean
// SDK-direct path; no hub. (The hub is a separate service for continuous small
// writes / the object-store API.)
func daCommand() *cli.Command {
	return &cli.Command{
		Name:  "da",
		Usage: "verifiable decentralized storage — direct upload/download",
		Subcommands: []*cli.Command{
			daUploadCmd(),
			daDownloadCmd(),
			daLsCmd(),
		},
	}
}

func requireGateway(c *cli.Context) (string, error) {
	gw := strings.TrimRight(c.String("gateway"), "/")
	if gw == "" {
		return "", fmt.Errorf("need --gateway (or UNIBASE_GATEWAY)")
	}
	return gw, nil
}

func loadKey(c *cli.Context) (*ecdsa.PrivateKey, error) {
	k := strings.TrimPrefix(c.String("key"), "0x")
	if k == "" {
		return nil, fmt.Errorf("upload needs a funded wallet key (--key or UNIBASE_KEY): ETH gas + UB bond")
	}
	return crypto.HexToECDSA(k)
}

func daUploadCmd() *cli.Command {
	return &cli.Command{
		Name:  "upload",
		Usage: "upload a file or directory directly to a stream node (client submits AddPiece)",
		Flags: commonFlags(
			&cli.StringFlag{Name: "path", Required: true, Usage: "file or directory to upload"},
			&cli.StringFlag{Name: "name", Usage: "public name for a single file (default: basename)"},
			&cli.StringFlag{Name: "stream", Usage: "pin a stream node address (else gateway picks an online one)"},
			&cli.StringFlag{Name: "chain", EnvVars: []string{"CHAIN_TYPE"}, Value: "base-sepolia", Usage: "chain type"},
			&cli.IntFlag{Name: "n", Value: 6, Usage: "erasure-code N"},
			&cli.IntFlag{Name: "k", Value: 4, Usage: "erasure-code K"},
		),
		Action: func(c *cli.Context) error {
			gw, err := requireGateway(c)
			if err != nil {
				return err
			}
			sk, err := loadKey(c)
			if err != nil {
				return err
			}
			if s := c.String("stream"); s != "" {
				os.Setenv("STREAM_PRIORITY", s) // SDK prefers this stream in discovery
			}
			au, err := key.BuildAuth(sk, []byte("upload"))
			if err != nil {
				return err
			}
			sdk.Login(gw, au)
			cm, err := contract.NewContractManage(sk, c.String("chain"))
			if err != nil {
				return err
			}
			if err := cm.CheckBalance(au.Addr); err != nil {
				return fmt.Errorf("wallet %s: %w (needs ETH gas + UB)", au.Addr.Hex(), err)
			}
			policy := types.Policy{N: uint8(c.Int("n")), K: uint8(c.Int("k"))}

			fp, _ := homedir.Expand(c.String("path"))
			fi, err := os.Stat(fp)
			if err != nil {
				return err
			}
			isJSON := c.Bool("json")
			if !fi.IsDir() {
				nm := c.String("name")
				if nm == "" {
					nm = filepath.Base(fp)
				}
				return uploadOne(cm, au, policy, gw, fp, nm, isJSON)
			}
			return filepath.Walk(fp, func(p string, info os.FileInfo, e error) error {
				if e != nil || info.IsDir() {
					return e
				}
				return uploadOne(cm, au, policy, gw, p, filepath.Base(p), isJSON)
			})
		},
	}
}

func uploadOne(cm *contract.ContractManage, au types.Auth, policy types.Policy, gw, fp, nm string, isJSON bool) error {
	ff, streamer, err := sdk.Upload(gw, au, policy, fp, nm) // direct → stream + UploadFileMeta(gateway)
	if err != nil {
		return err
	}
	pcs, err := sdk.CheckFileFull(ff, streamer, fp) // trustless: verify streamer encoded the real bytes
	if err != nil {
		return err
	}
	var tx string
	for _, pc := range pcs {
		t, err := cm.AddPiece(pc) // client's own wallet stakes + submits on-chain
		if err != nil {
			return err
		}
		tx = t
	}
	if isJSON {
		b, _ := json.Marshal(map[string]any{
			"name": ff.Name, "sha256": ff.Hash, "streamer": streamer.Hex(),
			"pieces": len(pcs), "lastTx": tx, "chainType": ff.ChainType,
		})
		fmt.Println(string(b))
	} else {
		fmt.Printf("%s → name=%s sha256=%s streamer=%s pieces=%d tx=%s\n",
			fp, ff.Name, ff.Hash, streamer.Hex(), len(pcs), tx)
	}
	return nil
}

func daDownloadCmd() *cli.Command {
	return &cli.Command{
		Name:  "download",
		Usage: "download a file by name (reconstructed from DA pieces); public, no key needed",
		Flags: commonFlags(
			&cli.StringFlag{Name: "name", Required: true, Usage: "file name (as registered on upload)"},
			&cli.StringFlag{Name: "out", Usage: "output file; '-' or empty writes stdout"},
		),
		Action: func(c *cli.Context) error {
			gw, err := requireGateway(c)
			if err != nil {
				return err
			}
			// Reads still need a valid signed Authorization (gateway/stream reject
			// "nil authorization"), but no funds — use --key if given, else an
			// ephemeral key just to sign the request.
			sk, err := loadKey(c)
			if err != nil {
				if sk, err = crypto.GenerateKey(); err != nil {
					return err
				}
			}
			au, err := key.BuildAuth(sk, []byte("download"))
			if err != nil {
				return err
			}
			name := c.String("name")
			out := c.String("out")
			if out == "" || out == "-" {
				// stdout: single attempt (can't retry a consumed pipe)
				return sdk.Download(gw, au, name, nil, os.Stdout)
			}
			// file: retry to ride out the gateway's post-AddPiece piece-index sync
			// lag ("record not found" right after upload). Re-create (truncate)
			// each attempt so a partial write never lingers — memory-safe for large files.
			fp, _ := homedir.Expand(out)
			var lastErr error
			for attempt := 1; attempt <= 6; attempt++ {
				f, err := os.Create(fp)
				if err != nil {
					return err
				}
				lastErr = sdk.Download(gw, au, name, nil, f)
				f.Close()
				if lastErr == nil {
					if c.Bool("json") {
						fmt.Printf(`{"name":%q,"out":%q}`+"\n", name, fp)
					} else {
						fmt.Fprintf(os.Stderr, "downloaded %s → %s\n", name, fp)
					}
					return nil
				}
				if !strings.Contains(lastErr.Error(), "record not found") {
					break // a real error, not the sync-lag race
				}
				fmt.Fprintf(os.Stderr, "not indexed yet (gateway syncing AddPiece), retry %d/6…\n", attempt)
				time.Sleep(5 * time.Second)
			}
			os.Remove(fp)
			return lastErr
		},
	}
}

func daLsCmd() *cli.Command {
	return &cli.Command{
		Name:  "ls",
		Usage: "list files on the gateway",
		Flags: commonFlags(
			&cli.IntFlag{Name: "limit", Value: 32},
		),
		Action: func(c *cli.Context) error {
			gw, err := requireGateway(c)
			if err != nil {
				return err
			}
			u := fmt.Sprintf("%s/api/listFile?start=0&count=%d", gw, c.Int("limit"))
			resp, err := (&http.Client{Timeout: 30 * time.Second}).Get(u)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			if resp.StatusCode != 200 {
				return fmt.Errorf("ls: HTTP %d %s", resp.StatusCode, body)
			}
			if c.Bool("json") {
				os.Stdout.Write(body)
				fmt.Println()
				return nil
			}
			var r struct {
				Files []struct {
					Name, Hash, Owner string
					Size              int64
				}
			}
			if json.Unmarshal(body, &r) != nil {
				os.Stdout.Write(body)
				return nil
			}
			for _, f := range r.Files {
				h := f.Hash
				if len(h) > 16 {
					h = h[:16]
				}
				fmt.Printf("%-40s %10d  %s  owner=%s\n", f.Name, f.Size, h, f.Owner)
			}
			return nil
		},
	}
}
