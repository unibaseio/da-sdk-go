package logfs

import (
	"bytes"
	"math/rand"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alecthomas/units"
	"github.com/unibaseio/da-sdk-go/lib/kv"
	"github.com/unibaseio/da-sdk-go/lib/utils"
)

func TestBB(t *testing.T) {
	bb, err := units.ParseBase2Bytes(strings.TrimSpace("10MB"))
	if err != nil {
		t.Fatal(err)
	}
	if int(bb) <= 0 {
		t.Fatalf("unexpected size: %d", int(bb))
	}
}

func TestFs(t *testing.T) {
	basedir := t.TempDir()
	mdir := filepath.Join(basedir, "kv")
	ds, err := kv.NewBadgerStore(mdir, &kv.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}

	rdir := filepath.Join(basedir, "data")
	fs, err := New(ds, rdir, "0xaaa", "0xbcd")
	if err != nil {
		t.Fatal(err)
	}

	key := utils.RandomBytes(15)
	val := []byte("abcdefg")

	err = fs.Put(key, val)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		length := rand.Int31n(240) + 16
		nkey := utils.RandomBytes(int(length))
		length = rand.Int31n(1024 * 1024)
		nval := utils.RandomBytes(int(length))
		err = fs.Put(nkey, nval)
		if err != nil {
			t.Fatal(err)
		}
	}

	nval, err := fs.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(val, nval) {
		t.Fatal("unequal val")
	}

}
