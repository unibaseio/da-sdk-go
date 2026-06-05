package main

import (
	"fmt"
	"os"

	"github.com/unibaseio/da-sdk-go/app/cmd"
	"github.com/unibaseio/da-sdk-go/build"
	"github.com/urfave/cli/v2"
)

func main() {
	local := make([]*cli.Command, 0, 1)
	local = append(local, cmd.InitCmd)
	local = append(local, serverCmd)

	app := cli.App{
		Name:    "hub",
		Version: build.UserVersion(),
		Usage:   "hub node",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  cmd.RepoStr,
				Value: "~/.hub",
				Usage: "app home dir",
			},
			&cli.StringFlag{
				Name:    cmd.PasswordStr,
				Aliases: []string{"pwd"},
				Value:   "aidemo123",
			},
		},
		Commands: local,
	}
	app.Setup()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err) // nolint:errcheck
		os.Exit(1)
	}
}
