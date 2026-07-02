// ubcli — the Unibase CLI. First module: DA (direct decentralized storage).
//
// Direct path (no hub): the client discovers a stream via the gateway (or -stream
// pins one), uploads straight to it, and the client's own wallet submits AddPiece
// on-chain. Continuous small writes (agent memory) use the hub service instead.
//
// Agent-friendly: --json output, stdin/stdout piping, no interactive prompts
// (key via --key or UNIBASE_KEY), clear exit codes.
//
//	export UNIBASE_GATEWAY=http://<gateway>
//	export UNIBASE_KEY=<hex private key>          # upload only; reads are public
//	ubcli da upload   --path ./weights.bin [--stream 0x<addr>] [--name w.bin]
//	ubcli da download --name weights.bin --out ./weights.bin
//	ubcli da ls       --json
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ubcli",
		Usage: "Unibase CLI — verifiable decentralized storage & more",
		// Common flags (--hub/--key/--json) live on each leaf command (see
		// commonFlags) so they resolve from the subcommand and work in any
		// position, e.g. `ubcli da ls --bucket b --json`.
		Commands: []*cli.Command{daCommand()},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
