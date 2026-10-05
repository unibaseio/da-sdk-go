package simplefs

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestFs(t *testing.T) {
	rdir := "~/test/data"
	rdir, _ = homedir.Expand(rdir)
	fs, err := New(rdir)
	if err != nil {
		t.Fatal(err)
	}

	key := []byte("testa")
	val := []byte("abcdefg")
	err = fs.Put(key, val)
	if err != nil {
		t.Fatal(err)
	}

	nval, err := fs.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(val, nval) {
		t.Fatal("unequal val")
	}

	pval, err := fs.Get(key, 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(val[1:3], pval) {
		t.Fatal("unequal partial val")
	}

	wval, err := fs.Get(key, 2, 8)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(val[2:], wval) {
		t.Fatal("unequal partial out val")
	}
}

// Interrupted writes (.tmp) are removed from the store's own directory on
// start-up, never from the process's working directory.
func TestWalkRemovesTmpInPlace(t *testing.T) {
	base := t.TempDir()
	cwd := t.TempDir()
	t.Chdir(cwd)

	sf, err := New(base)
	if err != nil {
		t.Fatal(err)
	}
	if err := sf.Put([]byte("DsFile/abcd1234"), []byte("hello")); err != nil {
		t.Fatal(err)
	}
	fn, _ := sf.getPath([]byte("abcd1234"))
	stray := fn + ".tmp"
	if err := os.WriteFile(stray, []byte("partial"), 0644); err != nil {
		t.Fatal(err)
	}
	bystander := filepath.Join(cwd, "abcd1234.tmp")
	if err := os.WriteFile(bystander, []byte("not ours"), 0644); err != nil {
		t.Fatal(err)
	}
	os.Remove(filepath.Join(base, usage)) // force the walk

	sf2, err := New(base)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(stray); !os.IsNotExist(err) {
		t.Fatalf("interrupted write left in place: %v", err)
	}
	if _, err := os.Stat(bystander); err != nil {
		t.Fatalf("file in the working directory was removed: %v", err)
	}
	if sf2.size != 5 {
		t.Fatalf("size %d, want 5", sf2.size)
	}
}

func TestKeysCannotLeaveBaseDir(t *testing.T) {
	sf, err := New(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"", ".", "..", "DsFile/..", "/"} {
		if err := sf.Put([]byte(k), []byte("x")); err == nil {
			t.Errorf("Put(%q) accepted", k)
		}
		if _, err := sf.Get([]byte(k)); err == nil {
			t.Errorf("Get(%q) accepted", k)
		}
	}
	if _, err := sf.Get([]byte("abcd1234"), -1, 4); err == nil {
		t.Error("negative range accepted")
	}
}

// Overwriting a key keeps the data readable and the usage exact.
func TestPutOverwriteAccounting(t *testing.T) {
	sf, err := New(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	k := []byte("DsFile/ffee0011")
	if err := sf.Put(k, make([]byte, 10)); err != nil {
		t.Fatal(err)
	}
	if err := sf.Put(k, make([]byte, 3)); err != nil {
		t.Fatal(err)
	}
	if b, err := sf.Get(k); err != nil || len(b) != 3 {
		t.Fatalf("get after overwrite: %d %v", len(b), err)
	}
	if sf.size != 3 {
		t.Fatalf("size %d, want 3", sf.size)
	}
}
