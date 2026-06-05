package cmd

import (
	"fmt"

	"github.com/unibaseio/da-sdk-go/build"
	"github.com/unibaseio/da-sdk-go/lib/config"
	"github.com/unibaseio/da-sdk-go/lib/repo"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

var InitCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize a repo, and create default wallet",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		repoDir, err := homedir.Expand(cctx.String(RepoStr))
		if err != nil {
			return err
		}

		if repo.Exists(repoDir) {
			return fmt.Errorf("repo at '%s' is already initialized", repoDir)
		}

		rep, err := repo.NewFSRepo(repoDir, config.NewDefaultConfig())
		if err != nil {
			return err
		}
		defer rep.Close()

		pw := cctx.String(PasswordStr)
		if pw == "" {
			pw, err = InputPassWord()
			if err != nil {
				return err
			}
		}

		ac, err := rep.Key().Create(pw)
		if err != nil {
			return err
		}
		cfg := rep.Config()
		cfg.Wallet.Address = ac.String()
		cfg.Chain.Type = build.CheckChain()

		rep.ReplaceConfig(cfg)

		fmt.Printf("=====  connect chain: %s  =====\n", cfg.Chain.Type)
		fmt.Printf("=====  create account: %s  =====\n", ac)

		return nil
	},
}
