package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
	"github.com/unibaseio/da-sdk-go/sdk"
	"github.com/urfave/cli/v2"
)

// commonFlags are shared by every leaf command (defined per-command so they
// resolve from the subcommand context and can appear in any position).
func commonFlags(extra ...cli.Flag) []cli.Flag {
	base := []cli.Flag{
		&cli.StringFlag{Name: "hub", EnvVars: []string{"UNIBASE_HUB"}, Value: "http://127.0.0.1:8086", Usage: "hub base URL"},
		&cli.StringFlag{Name: "key", EnvVars: []string{"UNIBASE_KEY"}, Usage: "owner private key (hex); required for writes"},
		&cli.BoolFlag{Name: "json", Usage: "machine-readable JSON output"},
	}
	return append(base, extra...)
}

// daCommand is the `ubcli da` group — the DA hub /v1 object store.
func daCommand() *cli.Command {
	return &cli.Command{
		Name:  "da",
		Usage: "verifiable object storage (upload/download/list) on Unibase DA",
		Subcommands: []*cli.Command{
			daUploadCmd(),
			daDownloadCmd(),
			daLsCmd(),
		},
	}
}

// ---- v1 client -------------------------------------------------------------

type v1Client struct {
	hub, owner, privk string
	isJSON            bool
	hc                *http.Client
}

// newClient builds a client from the global flags. needKey=true (writes) errors
// if no key; reads leave owner empty.
func newClient(c *cli.Context, needKey bool) (*v1Client, error) {
	cl := &v1Client{
		hub:    strings.TrimRight(c.String("hub"), "/"),
		isJSON: c.Bool("json"),
		hc:     &http.Client{Timeout: 300 * time.Second},
	}
	if k := strings.TrimPrefix(c.String("key"), "0x"); k != "" {
		sk, err := crypto.HexToECDSA(k)
		if err != nil {
			return nil, fmt.Errorf("bad --key/UNIBASE_KEY: %w", err)
		}
		cl.owner = crypto.PubkeyToAddress(sk.PublicKey).Hex()
		cl.privk = k
	} else if needKey {
		return nil, fmt.Errorf("this command writes and needs a wallet key (--key or UNIBASE_KEY)")
	}
	return cl, nil
}

func (c *v1Client) authHeader() string {
	au := sdk.BuildAuth(c.owner, c.privk, []byte("hub"))
	b, _ := json.Marshal(au)
	return hex.EncodeToString(b)
}

func (c *v1Client) do(method, path, ctype string, body io.Reader, signed bool) (int, []byte, error) {
	r, err := http.NewRequest(method, c.hub+path, body)
	if err != nil {
		return 0, nil, err
	}
	if ctype != "" {
		r.Header.Set("Content-Type", ctype)
	}
	if signed {
		r.Header.Set("Authorization", c.authHeader())
	}
	resp, err := c.hc.Do(r)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	rb, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, rb, nil
}

// ---- commands --------------------------------------------------------------

func daUploadCmd() *cli.Command {
	return &cli.Command{
		Name:  "upload",
		Usage: "upload a file (or stdin) as an object; auto-creates the bucket",
		Flags: commonFlags(
			&cli.StringFlag{Name: "bucket", Required: true, Usage: "bucket (namespace)"},
			&cli.StringFlag{Name: "kind", Value: "file", Usage: "bucket kind: memory|knowledgebase|file|model|dataset"},
			&cli.StringFlag{Name: "name", Usage: "object key (default: file basename, or 'stdin')"},
			&cli.StringFlag{Name: "path", Usage: "file to upload; '-' or empty reads stdin"},
			&cli.BoolFlag{Name: "wait", Usage: "block until committed on-chain (returns commitment + tx)"},
		),
		Action: func(c *cli.Context) error {
			cl, err := newClient(c, true)
			if err != nil {
				return err
			}
			bucket := c.String("bucket")

			// read content (file or stdin)
			var data []byte
			key := c.String("name")
			if p := c.String("path"); p != "" && p != "-" {
				fp, _ := homedir.Expand(p)
				if data, err = os.ReadFile(fp); err != nil {
					return err
				}
				if key == "" {
					key = filepath.Base(fp)
				}
			} else {
				if data, err = io.ReadAll(os.Stdin); err != nil {
					return err
				}
				if key == "" {
					key = "stdin"
				}
			}
			if len(data) == 0 {
				return fmt.Errorf("nothing to upload (empty file/stdin)")
			}

			// 1) ensure bucket (idempotent; kind immutable → 409)
			code, body, err := cl.do("PUT", "/v1/buckets/"+url.PathEscape(bucket), "application/json",
				strings.NewReader(fmt.Sprintf(`{"kind":%q}`, c.String("kind"))), true)
			if err != nil {
				return err
			}
			if code != http.StatusOK && code != http.StatusCreated {
				return fmt.Errorf("create bucket: HTTP %d %s", code, body)
			}

			// 2) put object
			path := "/v1/buckets/" + url.PathEscape(bucket) + "/objects/" + url.PathEscape(key)
			if c.Bool("wait") {
				path += "?wait=1"
			}
			code, body, err = cl.do("PUT", path, "application/octet-stream", strings.NewReader(string(data)), true)
			if err != nil {
				return err
			}
			if code != http.StatusOK && code != http.StatusAccepted {
				return fmt.Errorf("upload: HTTP %d %s", code, body)
			}
			return cl.emitReceipt(body)
		},
	}
}

func daDownloadCmd() *cli.Command {
	return &cli.Command{
		Name:  "download",
		Usage: "download an object's bytes to a file or stdout",
		Flags: commonFlags(
			&cli.StringFlag{Name: "bucket", Required: true},
			&cli.StringFlag{Name: "name", Required: true, Usage: "object key"},
			&cli.StringFlag{Name: "owner", Usage: "owner address (scopes lookup; optional)"},
			&cli.StringFlag{Name: "out", Usage: "output file; '-' or empty writes stdout (pipe-friendly)"},
		),
		Action: func(c *cli.Context) error {
			cl, err := newClient(c, false)
			if err != nil {
				return err
			}
			base := "/v1/buckets/" + url.PathEscape(c.String("bucket")) + "/objects/" + url.PathEscape(c.String("name"))
			q := ""
			if o := c.String("owner"); o != "" {
				q = "?owner=" + o
			}
			code, body, err := cl.do("GET", base+"/content"+q, "", nil, false)
			if err != nil {
				return err
			}
			if code != http.StatusOK {
				return fmt.Errorf("download: HTTP %d %s", code, body)
			}
			out := c.String("out")
			if out == "" || out == "-" {
				os.Stdout.Write(body) // pure bytes → pipeable
				return nil
			}
			fp, _ := homedir.Expand(out)
			if err := os.WriteFile(fp, body, 0o644); err != nil {
				return err
			}
			if cl.isJSON {
				fmt.Printf(`{"out":%q,"bytes":%d}`+"\n", fp, len(body))
			} else {
				fmt.Printf("wrote %d bytes to %s\n", len(body), fp)
			}
			return nil
		},
	}
}

func daLsCmd() *cli.Command {
	return &cli.Command{
		Name:  "ls",
		Usage: "list objects in a bucket, or list buckets",
		Flags: commonFlags(
			&cli.StringFlag{Name: "bucket", Usage: "list objects in this bucket; omit to list buckets"},
			&cli.StringFlag{Name: "owner", Usage: "filter by owner"},
			&cli.StringFlag{Name: "kind", Usage: "filter buckets by kind"},
			&cli.IntFlag{Name: "limit", Value: 32},
		),
		Action: func(c *cli.Context) error {
			cl, err := newClient(c, false)
			if err != nil {
				return err
			}
			q := url.Values{}
			if o := c.String("owner"); o != "" {
				q.Set("owner", o)
			}
			q.Set("limit", fmt.Sprintf("%d", c.Int("limit")))

			var path, field string
			if b := c.String("bucket"); b != "" {
				path = "/v1/buckets/" + url.PathEscape(b) + "/objects?" + q.Encode()
				field = "objects"
			} else {
				if k := c.String("kind"); k != "" {
					q.Set("kind", k)
				}
				path = "/v1/buckets?" + q.Encode()
				field = "buckets"
			}
			code, body, err := cl.do("GET", path, "", nil, false)
			if err != nil {
				return err
			}
			if code != http.StatusOK {
				return fmt.Errorf("ls: HTTP %d %s", code, body)
			}
			if cl.isJSON {
				os.Stdout.Write(body)
				fmt.Println()
				return nil
			}
			return printList(body, field)
		},
	}
}

// ---- output helpers --------------------------------------------------------

// emitReceipt prints the object receipt: raw JSON when --json, else a summary.
func (c *v1Client) emitReceipt(body []byte) error {
	if c.isJSON {
		os.Stdout.Write(body)
		fmt.Println()
		return nil
	}
	var r struct {
		Bucket, Key, Status, Commitment, Availability string
		Size                                          uint64
		Chain                                         *struct{ TxHash string }
	}
	json.Unmarshal(body, &r)
	tx := ""
	if r.Chain != nil {
		tx = r.Chain.TxHash
	}
	fmt.Printf("%s/%s  %d bytes  status=%s\n", r.Bucket, r.Key, r.Size, r.Status)
	if r.Commitment != "" {
		fmt.Printf("  commitment=%s\n  tx=%s (%s)\n", r.Commitment, tx, r.Availability)
	}
	return nil
}

func printList(body []byte, field string) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(body, &m); err != nil {
		os.Stdout.Write(body)
		return nil
	}
	if field == "buckets" {
		var bs []struct{ Name, Owner, Kind string }
		json.Unmarshal(m["buckets"], &bs)
		for _, b := range bs {
			fmt.Printf("%-40s kind=%-14s owner=%s\n", b.Name, b.Kind, b.Owner)
		}
	} else {
		var os_ []struct {
			Key, Status, Commitment string
			Size                    uint64
		}
		json.Unmarshal(m["objects"], &os_)
		for _, o := range os_ {
			fmt.Printf("%-40s %10d  %-9s %s\n", o.Key, o.Size, o.Status, o.Commitment)
		}
	}
	if nc := strings.Trim(string(m["nextCursor"]), `"`); nc != "" {
		fmt.Printf("(more: --... cursor=%s)\n", nc)
	}
	return nil
}
