package main

import (
	"fmt"
	"log"
	"math/big"

	contract "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/bank"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/control"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/epoch"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/eproof"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/everify"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/gpu"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/model"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/node"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/piece"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/plonk/add"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/plonk/kzg"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/plonk/mul"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/plonk/rsone"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/reward"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/rsproof"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/space"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Set(client *ethclient.Client, sk string, _typ string, ca common.Address) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	bi, err := bank.NewBank(bankAddr, client)
	if err != nil {
		return err
	}

	tx, err := bi.Set(au, _typ, ca)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println(_typ, " is set to: ", ca.String())
	return nil
}

func DeployBank(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := bank.DeployBank(au, client)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("bank: ", tAddr.Hex())
	bankAddr = tAddr

	Set(client, sk, "base", contract.Base)
	return nil
}

func DeployToken(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, ti, err := token.DeployToken(au, client, "Unibase", "UB")
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("token: ", tAddr.Hex())
	Set(client, sk, "token", tAddr)
	tokenAddr = tAddr

	owner, err := ti.Owner(&bind.CallOpts{From: contract.Base})
	if err != nil {
		return err
	}
	bal, err := ti.BalanceOf(&bind.CallOpts{From: contract.Base}, owner)
	if err != nil {
		return err
	}
	log.Println("owner has token: ", bal)
	return nil
}

func DeployEpoch(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := epoch.DeployEpoch(au, client, blockInterval)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}

	log.Println("epoch: ", tAddr.Hex())
	Set(client, sk, "epoch", tAddr)
	return nil
}

func DeployReward(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := reward.DeployReward(au, client, bankAddr)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("reward: ", tAddr.Hex())
	Set(client, sk, "reward", tAddr)
	return nil
}

func DeployControl(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := control.DeployControl(au, client, bankAddr)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("control: ", tAddr.Hex())
	Set(client, sk, "control", tAddr)
	return nil
}

func DeployNode(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := node.DeployNode(au, client, bankAddr)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("node: ", tAddr.Hex())
	Set(client, sk, "node", tAddr)
	return nil
}

func DeployPiece(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := piece.DeployPiece(au, client, bankAddr)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("piece: ", tAddr.Hex())
	Set(client, sk, "piece", tAddr)
	return nil
}

func DeployRSProof(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := rsproof.DeployRSProof(au, client, bankAddr, blockInterval)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("rsproof: ", tAddr.Hex())
	Set(client, sk, "rsproof", tAddr)

	return nil
}

func DeployRSPlonk(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := rsone.DeployPlonkVerifier(au, client)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}

	log.Println("rsone: ", tAddr.Hex())
	Set(client, sk, "rsone", tAddr)
	return nil
}

func SetRSVKRoot(client *ethclient.Client, sk string, n, k int) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	bi, err := bank.NewBank(bankAddr, client)
	if err != nil {
		return err
	}

	rsAddr, err := bi.Get(&bind.CallOpts{From: au.From}, "rsproof")
	if err != nil {
		return err
	}

	ni, err := rsproof.NewRSProof(rsAddr, client)
	if err != nil {
		return err
	}

	vt := new(big.Int)
	switch {
	case n == 6 && k == 4:
		vt.SetString(contract.RSn6k4VKRoot, 10)
	case n == 14 && k == 7:
		vt.SetString(contract.RSn14k7VKRoot, 10)
	case n == 32 && k == 16:
		vt.SetString(contract.RSn32k16VKRoot, 10)
	case n == 64 && k == 32:
		vt.SetString(contract.RSn64k32VKRoot, 10)
	default:
		return fmt.Errorf("unsupported rs policy: %d %d", n, k)
	}

	tx, err := ni.SetVKRoot(au, uint8(n), uint8(k), vt)
	if err != nil {
		return err
	}
	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Printf("set vk root: %d %d %d\n", n, k, vt)
	return nil
}

func DeployEproof(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := eproof.DeployEProof(au, client, bankAddr, blockInterval)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("eproof: ", tAddr.Hex())
	Set(client, sk, "eproof", tAddr)
	return nil
}

func DeployEverify(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := everify.DeployEVerify(au, client, bankAddr)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("everify: ", tAddr.Hex())
	Set(client, sk, "everify", tAddr)
	return nil
}

func DeployKZGPlonk(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := kzg.DeployPlonkVerifier(au, client)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}

	log.Println("kzg: ", tAddr.Hex())
	Set(client, sk, "kzg", tAddr)
	return nil
}

func DeployMulPlonk(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := mul.DeployPlonkVerifier(au, client)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}

	log.Println("mul: ", tAddr.Hex())
	Set(client, sk, "mul", tAddr)
	return nil
}

func DeployAddPlonk(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := add.DeployPlonkVerifier(au, client)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}

	log.Println("add: ", tAddr.Hex())
	Set(client, sk, "add", tAddr)
	return nil
}

func DeployModel(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := model.DeployModel(au, client, bankAddr)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("model: ", tAddr.Hex())
	Set(client, sk, "model", tAddr)
	return nil
}

func DeployGPU(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := gpu.DeployGPU(au, client, bankAddr)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("gpu: ", tAddr.Hex())
	Set(client, sk, "gpu", tAddr)
	return nil
}

func DeploySpace(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, _, err := space.DeploySpace(au, client, bankAddr)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("space: ", tAddr.Hex())
	Set(client, sk, "space", tAddr)
	return nil
}

func SetMinPledge(client *ethclient.Client, sk string, _typ uint8, val *big.Int) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	bi, err := bank.NewBank(bankAddr, client)
	if err != nil {
		return err
	}

	nodeAddr, err := bi.Get(&bind.CallOpts{From: au.From}, "node")
	if err != nil {
		return err
	}

	ni, err := node.NewNode(nodeAddr, client)
	if err != nil {
		return err
	}

	tx, err := ni.Set(au, _typ, val)
	if err != nil {
		return err
	}
	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("set minpledge success")
	return nil
}

func SetMiner(client *ethclient.Client, sk string, _a common.Address) error {
	ti, err := token.NewToken(tokenAddr, client)
	if err != nil {
		return err
	}
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tx, err := ti.TransferOwnership(au, _a)
	if err != nil {
		return err
	}
	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}
