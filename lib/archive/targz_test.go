package archive

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestTarGzRoundTrip(t *testing.T) {
	src := t.TempDir()
	must := func(err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
	}
	must(os.MkdirAll(filepath.Join(src, "a", "b"), 0o755))
	must(os.WriteFile(filepath.Join(src, "top.txt"), []byte("top"), 0o644))
	must(os.WriteFile(filepath.Join(src, "a", "b", "deep.bin"), bytes.Repeat([]byte{7}, 70000), 0o644))
	must(os.Symlink("/etc/passwd", filepath.Join(src, "link")))

	rc, err := TarGz(src)
	must(err)
	defer rc.Close()
	dst := t.TempDir()
	must(UntarGz(rc, dst))

	if b, err := os.ReadFile(filepath.Join(dst, "top.txt")); err != nil || string(b) != "top" {
		t.Fatalf("top.txt: %q %v", b, err)
	}
	if b, err := os.ReadFile(filepath.Join(dst, "a", "b", "deep.bin")); err != nil || len(b) != 70000 {
		t.Fatalf("deep.bin: %d %v", len(b), err)
	}
	if _, err := os.Lstat(filepath.Join(dst, "link")); !os.IsNotExist(err) {
		t.Fatalf("symlink was archived: %v", err)
	}
}

func archiveOf(t *testing.T, hdrs ...*tar.Header) io.Reader {
	t.Helper()
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)
	for _, h := range hdrs {
		if h.Typeflag == tar.TypeReg {
			h.Size = 1
		}
		if err := tw.WriteHeader(h); err != nil {
			t.Fatal(err)
		}
		if h.Typeflag == tar.TypeReg {
			tw.Write([]byte("x"))
		}
	}
	tw.Close()
	gw.Close()
	return &buf
}

// The zip-slip family: nothing is written outside the target directory.
func TestUntarGzRefusesEscapes(t *testing.T) {
	for name, h := range map[string]*tar.Header{
		"parent":   {Name: "../evil", Typeflag: tar.TypeReg, Mode: 0o644},
		"nested":   {Name: "a/../../evil", Typeflag: tar.TypeReg, Mode: 0o644},
		"absolute": {Name: "/tmp/evil", Typeflag: tar.TypeReg, Mode: 0o644},
	} {
		parent := t.TempDir()
		dst := filepath.Join(parent, "out")
		if err := UntarGz(archiveOf(t, h), dst); err == nil {
			t.Errorf("%s: accepted", name)
		}
		if _, err := os.Stat(filepath.Join(parent, "evil")); err == nil {
			t.Errorf("%s: wrote outside the target", name)
		}
	}
}

// Links are skipped, never created (they could point outside the target).
func TestUntarGzSkipsLinks(t *testing.T) {
	dst := t.TempDir()
	err := UntarGz(archiveOf(t,
		&tar.Header{Name: "l", Typeflag: tar.TypeSymlink, Linkname: "/etc"},
		&tar.Header{Name: "h", Typeflag: tar.TypeLink, Linkname: "../x"},
		&tar.Header{Name: "ok", Typeflag: tar.TypeReg, Mode: 0o644},
	), dst)
	if err != nil {
		t.Fatal(err)
	}
	for _, n := range []string{"l", "h"} {
		if _, err := os.Lstat(filepath.Join(dst, n)); !os.IsNotExist(err) {
			t.Errorf("link %s created", n)
		}
	}
	if _, err := os.Stat(filepath.Join(dst, "ok")); err != nil {
		t.Errorf("regular file not extracted: %v", err)
	}
}
