package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	com "github.com/unibaseio/da-sdk-go/contract/common"
	contract "github.com/unibaseio/da-sdk-go/contract/v2"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
)

func main() {
	skstr := flag.String("sk", "", "private key for sending transaction")
	pathstr := flag.String("path", "", "dir or file path to upload")
	//mf := flag.Bool("model", false, "upload type: model or regular file/dir")
	fname := flag.Bool("name", false, "file name in public, default is sha256")
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

	err = UploadFile(sk, fp, *fname)
	if err != nil {
		log.Println(err)
	}
}

func UploadFile(sk *ecdsa.PrivateKey, fp string, fname bool) error {
	au, err := key.BuildAuth(sk, []byte("upload"))
	if err != nil {
		return err
	}

	// charge from server
	sdk.Login(sdk.ServerURL, au)

	cm, err := contract.NewContractManage(sk, com.BaseSepolia)
	if err != nil {
		return err
	}

	err = cm.CheckBalance(au.Addr)
	if err != nil {
		return err
	}

	fi, err := os.Stat(fp)
	if err != nil {
		return err
	}

	policy := types.Policy{
		N: 6,
		K: 4,
	}

	if !fi.IsDir() {
		nm := ""
		if fname {
			nm = filepath.Base(fp)
		}

		// upload to stream and submit to gateway
		res, streamer, err := sdk.Upload(sdk.ServerURL, au, policy, fp, nm)
		if err != nil {
			return err
		}
		pcs, err := sdk.CheckFileFull(res, streamer, fp)
		if err != nil {
			return err
		}
		log.Printf("upload %s to %s, sha256: %s\n", fp, streamer, res.Hash)
		log.Printf("submit %s to chain\n", res.Name)

		// submit meta to chain
		for _, pc := range pcs {
			_, err = cm.AddPiece(pc)
			if err != nil {
				return err
			}
		}

		return nil
	}

	// recursive upload files in directory
	return filepath.Walk(fp, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		nm := ""
		if fname {
			nm = filepath.Base(fileName)
		}
		res, streamer, err := sdk.Upload(sdk.ServerURL, au, policy, fileName, nm)
		if err != nil {
			return nil
		}
		pcs, err := sdk.CheckFileFull(res, streamer, fp)
		if err != nil {
			return err
		}

		log.Printf("upload %s to %s, sha256: %s\n", fp, streamer, res.Hash)
		log.Printf("submit %s to chain\n", res.Name)

		// submit meta to chain
		for _, pc := range pcs {
			_, err = cm.AddPiece(pc)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
