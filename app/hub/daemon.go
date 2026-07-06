package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"
	"time"

	"github.com/unibaseio/da-sdk-go/app/cmd"
	"github.com/unibaseio/da-sdk-go/build"
	"github.com/unibaseio/da-sdk-go/hub"
	"github.com/unibaseio/da-sdk-go/lib/repo"
	"github.com/unibaseio/da-sdk-go/lib/utils"
	"github.com/unibaseio/da-sdk-go/sdk"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

var serverCmd = &cli.Command{
	Name:  "daemon",
	Usage: "dimo hub daemon",
	Subcommands: []*cli.Command{
		runCmd,
		cmd.StopCmd,
	},
}

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "run hub node for upload and download",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    cmd.EndpointStr,
			Aliases: []string{"b"},
			Usage:   "input your endpoint",
			Value:   "127.0.0.1:8086",
		},
		&cli.StringFlag{
			Name:    cmd.RemoteURLStr,
			Aliases: []string{"r"},
			Usage:   "input remote server url",
			Value:   build.ServerURL,
		},
		&cli.StringFlag{
			Name:    cmd.ExposeStr,
			Aliases: []string{"e"},
			Usage:   "input expose url for external access",
			Value:   "http://127.0.0.1:8086",
		},
	},
	Action: func(cctx *cli.Context) error {
		rpath, err := homedir.Expand(cctx.String(cmd.RepoStr))
		if err != nil {
			return err
		}

		if !repo.Exists(rpath) {
			return fmt.Errorf("please init first")
		}

		rp, err := repo.NewFSRepo(rpath, nil)
		if err != nil {
			return err
		}
		cfg := rp.Config()

		ct := build.CheckChain()
		if ct != rp.Config().Chain.Type {
			return fmt.Errorf("env 'CHAIN_TYPE' should be same with config %s", rp.Config().Chain.Type)
		}

		cfg.API.Endpoint = cctx.String(cmd.EndpointStr)
		cfg.Remote.URL = cctx.String(cmd.RemoteURLStr)
		cfg.API.Expose = cctx.String(cmd.ExposeStr)

		he := os.Getenv("EXPOSE_URL")
		if he != "" {
			cfg.API.Expose = he
		}

		// Reachability probe only — the gateway is needed for DA sealing and
		// discovery, not for serving (local logfs + index DB). A down or
		// mid-upgrade gateway must not keep the hub offline; registration
		// retries in the background (hub.Server.registerLoop).
		if _, err := sdk.Info(cfg.Remote.URL); err != nil {
			log.Printf("WARN: gateway %s not reachable at boot (%v) — starting anyway; background registration will retry", cfg.Remote.URL, err)
		}
		rp.ReplaceConfig(cfg)

		pw := cctx.String(cmd.PasswordStr)
		if pw == "" {
			pw, err = cmd.InputPassWord()
			if err != nil {
				return err
			}
		}

		err = rp.Key().Load(utils.HexToAddress(cfg.Wallet.Address), pw)
		if err != nil {
			return err
		}

		srv, err := hub.NewServer(rp)
		if err != nil {
			return err
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		log.Println("listen on port: ", cfg.API.Endpoint)

		pid := os.Getpid()
		pids := []byte(strconv.Itoa(pid))
		err = os.WriteFile(path.Join(rpath, "pid"), pids, 0644)
		if err != nil {
			return err
		}

		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("forced to shutdown: ", err)
		}

		log.Println("hub daemon exited")
		return nil
	},
}
