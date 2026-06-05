package estimator

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/utils"
	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestL2Fee(t *testing.T) {
	l2client, err := ethclient.Dial(L2Endpoint)
	if err != nil {
		t.Fatal(err)
	}
	defer l2client.Close()
	gpo, err := bindings.NewGasPriceOracle(predeploys.GasPriceOracleAddr, l2client)
	if err != nil {
		t.Fatal(err)
	}

	p, err := GetPrice(gpo, nil)
	if err != nil {
		t.Fatal(err)
	}

	data := utils.RandomBytes(2200)

	l1fee := L1Fee(p, GetSize(FlzCompress(data)))
	fmt.Println("l1fee: ", l1fee)
	l1fee, err = gpo.GetL1Fee(nil, data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("l1fee: ", l1fee)
}

func TestL2(t *testing.T) {
	l2client, err := ethclient.Dial(L2Endpoint)
	if err != nil {
		t.Fatal(err)
	}
	defer l2client.Close()
	gpo, err := bindings.NewGasPriceOracle(predeploys.GasPriceOracleAddr, l2client)
	if err != nil {
		t.Fatal(err)
	}
	bn, err := l2client.BlockNumber(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(bn)
	opt := &bind.CallOpts{BlockNumber: big.NewInt(int64(bn - 1))}
	ver, err := gpo.Version(opt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("version: ", ver)

	isf, err := gpo.IsFjord(opt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("fjord: ", isf)

	bf, err := gpo.L1BaseFee(opt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("basefee: ", bf)

	bfs, err := gpo.BaseFeeScalar(opt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("basefeescalar: ", bfs)

	bbf, err := gpo.BlobBaseFee(opt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("blobbasefee: ", bbf)

	bbfs, err := gpo.BlobBaseFeeScalar(opt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("blobbasefeescalar: ", bbfs)

	data := utils.RandomBytes(10)

	l1fee, err := gpo.GetL1Fee(opt, data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("l1fee: ", l1fee)
}
