// unibase — the Unibase CLI. First module: DA (hub /v1 object store).
//
// Designed to be agent-friendly: --json structured output, stdin/stdout piping,
// no interactive prompts (key via --key or UNIBASE_KEY), clear exit codes.
//
//	export UNIBASE_HUB=http://127.0.0.1:8086
//	export UNIBASE_KEY=<hex private key>          # writes only; reads are public
//	unibase da upload   --bucket my-models --kind model --path ./weights.bin --wait
//	unibase da download --bucket my-models --key weights.bin --out ./weights.bin
//	unibase da ls       --bucket my-models --json
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "unibase",
		Usage: "Unibase CLI — verifiable decentralized storage & more",
		// Common flags (--hub/--key/--json) live on each leaf command (see
		// commonFlags) so they resolve from the subcommand and work in any
		// position, e.g. `unibase da ls --bucket b --json`.
		Commands: []*cli.Command{daCommand()},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
