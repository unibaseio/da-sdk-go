package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/kv"
	"github.com/unibaseio/da-sdk-go/lib/piece"
	"github.com/unibaseio/da-sdk-go/lib/simplefs"
	"github.com/unibaseio/da-sdk-go/sdk"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
)

func main() {
	skstr := flag.String("sk", "", "private key for sending transaction")
	namestr := flag.String("name", "", "file or model name to download")
	pathstr := flag.String("path", "", "dir or file path to save")
	mf := flag.Bool("model", false, "download type: model or regular file")
	flag.Parse()

	sk, err := crypto.HexToECDSA(*skstr)
	if err != nil {
		sk, err = crypto.GenerateKey()
		if err != nil {
			return
		}
		skbyte := crypto.FromECDSA(sk)
		fmt.Printf("=== generate privatekey: %s ===\n", hex.EncodeToString(skbyte))
	}

	fp, err := homedir.Expand(*pathstr)
	if err != nil {
		return
	}

	if *mf {
		err = DownloadModel(sk, *namestr, fp)
		if err != nil {
			log.Println(err)
		}
	} else {
		err = DownloadFile(sk, *namestr, fp)
		if err != nil {
			log.Println(err)
		}
	}
}

func DownloadFile(sk *ecdsa.PrivateKey, fname, fp string) error {
	au, err := key.BuildAuth(sk, []byte("download"))
	if err != nil {
		return err
	}

	finfo, err := os.Stat(fp)
	if err == nil {
		if finfo.IsDir() {
			fp = filepath.Join(fp, fname)
			log.Println("will save file to: ", fp)
		} else {
			fmt.Printf("overwrite %s? \n", fp)
			fmt.Printf("please input 'yes' to continue: ")
			var input string
			fmt.Scan(&input)
			if input != "yes" {
				return fmt.Errorf("confirm your path again")
			}
		}
	}

	fi, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer fi.Close()

	fr, err := sdk.GetFileReceipt(sdk.ServerURL, au, fname)
	if err != nil {
		return err
	}
	if fr.Size < 1024*1024*256 {
		err = sdk.Download(sdk.ServerURL, au, fname, nil, fi)
		if err != nil {
			return err
		}
		log.Printf("download %s to: %s\n", fname, fp)
		return nil
	}

	cachedir, err := homedir.Expand("~/.dimocache")
	if err != nil {
		return err
	}

	log.Printf("use local dir '%s' as cache\n", cachedir)
	defer os.RemoveAll(cachedir)

	ds, err := kv.NewBadgerStore(path.Join(cachedir, "meta"), nil)
	if err != nil {
		return err
	}
	fs, err := simplefs.New(path.Join(cachedir, "data"))
	if err != nil {
		return err
	}

	ks := piece.New(ds, fs)

	// download 8 in parallel
	err = sdk.DownloadParallel(sdk.ServerURL, au, fp, 8, ks, fi)
	if err != nil {
		return err
	}

	log.Printf("download %s to: %s\n", fname, fp)
	return nil
}

func DownloadModel(sk *ecdsa.PrivateKey, name, fp string) error {
	au, err := key.BuildAuth(sk, []byte("upload"))
	if err != nil {
		return err
	}

	finfo, err := os.Stat(fp)
	if err != nil {
		return err
	}

	if !finfo.IsDir() {
		return fmt.Errorf("please input dir for saving model files")
	}

	cachedir, err := homedir.Expand("~/.dimocache")
	if err != nil {
		return err
	}

	log.Printf("use local dir '%s' as cache\n", cachedir)
	defer os.RemoveAll(cachedir)

	ds, err := kv.NewBadgerStore(path.Join(cachedir, "meta"), nil)
	if err != nil {
		return err
	}
	fs, err := simplefs.New(path.Join(cachedir, "data"))
	if err != nil {
		return err
	}

	ks := piece.New(ds, fs)

	mrm, err := sdk.GetModel(sdk.ServerURL, au, name)
	if err != nil {
		return err
	}

	err = sdk.DownloadModel(sdk.ServerURL, fp, au, mrm.ModelMeta, ks)
	if err != nil {
		return err
	}
	log.Printf("download model %s to dir: %s\n", name, fp)
	return nil
}
