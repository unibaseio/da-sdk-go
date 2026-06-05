package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/unibaseio/da-sdk-go/lib/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

var StopCmd = &cli.Command{
	Name:  "stop",
	Usage: "stop daemon instance",
	Action: func(cctx *cli.Context) error {
		pidpath, err := homedir.Expand(cctx.String(RepoStr))
		if err != nil {
			return err
		}

		pd, err := os.ReadFile(path.Join(pidpath, "pid"))
		if err != nil {
			return err
		}

		err = utils.KillProcess(string(pd))
		if err != nil {
			return err
		}

		fmt.Println("daemon instance gracefully exited...")

		return nil
	},
}
