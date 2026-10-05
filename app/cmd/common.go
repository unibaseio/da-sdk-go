package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/howeyc/gopass"
)

const (
	RepoStr      = "repo"
	PasswordStr  = "password"
	EndpointStr  = "bind"
	ExposeStr    = "expose"
	RemoteURLStr = "remote"
)

// InputPassWord prompts for the keystore password on the terminal. There is
// no built-in default: a node without a TTY (systemd) must get it from
// --password.
func InputPassWord() (string, error) {
	type result struct {
		pw  string
		err error
	}
	ch := make(chan result, 1) // buffered: a late answer after the timeout is dropped, not raced
	go func() {
		fmt.Fprint(os.Stderr, "Please input your password (at least 8): ")
		pd, err := gopass.GetPasswdMasked()
		ch <- result{string(pd), err}
	}()

	select {
	case r := <-ch:
		if r.err != nil {
			return "", fmt.Errorf("read password: %w (pass --password)", r.err)
		}
		if len(r.pw) < 8 {
			return "", fmt.Errorf("password length should be at least 8")
		}
		return r.pw, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("no keystore password: pass --password")
	}
}
