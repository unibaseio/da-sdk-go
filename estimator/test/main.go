package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/unibaseio/da-sdk-go/estimator"
	"github.com/unibaseio/da-sdk-go/lib/utils"
	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	l2fee()
}

func l2fee() {
	l2client, err := ethclient.Dial(estimator.L2Endpoint)
	if err != nil {
		return
	}
	defer l2client.Close()
	gpo, err := bindings.NewGasPriceOracle(predeploys.GasPriceOracleAddr, l2client)
	if err != nil {
		return
	}

	data := utils.RandomBytes(2200)
	gasused := 370_000
	size := estimator.GetSize(estimator.FlzCompress(data))

	l1total := new(big.Int)
	l2total := new(big.Int)
	total := new(big.Int)
	l1average := new(big.Int)
	l2average := new(big.Int)
	average := new(big.Int)
	cnt := 0

	bn, err := l2client.BlockNumber(context.TODO())
	if err != nil {
		return
	}

	for {
		header, err := l2client.HeaderByNumber(context.TODO(), big.NewInt(int64(bn)))
		if err != nil {
			continue
		}

		opt := &bind.CallOpts{BlockNumber: big.NewInt(int64(bn))}
		p, err := estimator.GetPrice(gpo, opt)
		if err != nil {
			continue
		}
		cnt++
		bn -= 30

		l2fee := new(big.Int).Mul(header.BaseFee, big.NewInt(int64(gasused)))
		l2total.Add(l2total, l2fee)
		l2average.Div(l2total, big.NewInt(int64(cnt)))

		l1fee := estimator.L1Fee(p, size)
		l1total.Add(l1total, l1fee)
		l1average.Div(l1total, big.NewInt(int64(cnt)))

		fee := new(big.Int).Add(l1fee, l2fee)
		total.Add(total, fee)
		average.Div(total, big.NewInt(int64(cnt)))

		fmt.Printf("%d,head: %d, l1: %s, l2: %s, total: %s, l1 average: %s, l2 average: %s, average: %s\n", cnt, bn, utils.FormatEth(l1fee), utils.FormatEth(l2fee), utils.FormatEth(fee), utils.FormatEth(l1average), utils.FormatEth(l2average), utils.FormatEth(average))
	}
}

func l1fee() {
	l2client, err := ethclient.Dial(estimator.L1Endpoint)
	if err != nil {
		return
	}
	defer l2client.Close()

	gasused := 370_000

	l2total := new(big.Int)
	l2average := new(big.Int)
	cnt := 0

	bn, err := l2client.BlockNumber(context.TODO())
	if err != nil {
		return
	}

	for {
		header, err := l2client.HeaderByNumber(context.TODO(), big.NewInt(int64(bn)))
		if err != nil {
			continue
		}

		cnt++
		bn -= 30

		l2fee := new(big.Int).Mul(header.BaseFee, big.NewInt(int64(gasused)))
		l2total.Add(l2total, l2fee)
		l2average.Div(l2total, big.NewInt(int64(cnt)))

		fmt.Printf("%d,head: %d, total: %s, average: %s\n", cnt, bn, utils.FormatEth(l2fee), utils.FormatEth(l2average))
	}
}
