package param

import (
	"errors"
	"io/fs"
	"testing"
)

func TestFetch(t *testing.T) {
	pn := "kzgsrs-bw6_761-lagrange-131072"

	err := Fetch(pn)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheck(t *testing.T) {
	err := CheckAll()
	if errors.Is(err, fs.ErrNotExist) {
		t.Skipf("SRS params not downloaded on this machine: %v", err)
	}
	if err != nil {
		t.Fatal(err)
	}
}
