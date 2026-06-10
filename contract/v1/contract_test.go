package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	com "github.com/unibaseio/da-sdk-go/contract/common"

	"github.com/unibaseio/da-sdk-go/build"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestTransfer(t *testing.T) {
	admin := common.HexToAddress("0xE251825869b2262e82AA843feea29B7E308764CC")
	cm, err := NewContractManage(nil, build.BNBTestnetV2)
	if err != nil {
		t.Fatal(err)
	}

	has := cm.BalanceOf(admin)
	fmt.Println("admin has: ", has)

	sk, addr := makeAccount()
	valt := big.NewInt(1e18)
	valt.Mul(valt, big.NewInt(10000))

	val := big.NewInt(2e17)
	addr = common.HexToAddress("0xcf2bf532adbed038b849416b2346633c57bcc3fe")
	err = testtransfer(addr, val, valt)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(sk)
}

func TestFakeReplica(t *testing.T) {
	valt := big.NewInt(1e18)
	valt.Mul(valt, big.NewInt(10))

	pkey, paddr := makeAccount()

	cm, err := NewContractManage(pkey, build.BNBTestnetV2)
	if err != nil {
		t.Fatal(err)
	}

	val := big.NewInt(1e15)
	err = testtransfer(paddr, val, valt)
	if err != nil {
		t.Fatal(err)
	}

	err = cm.RegisterNode(1, nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	fi, err := cm.NewPiece(ctx)
	if err != nil {
		t.Fatal(err)
	}

	au, err := cm.MakeAuth()
	if err != nil {
		t.Fatal(err)
	}

	rb := []byte("test-" + time.Now().String())
	_pi := uint64(76)
	_pri := uint8(1)
	pf := make([]byte, 32)
	fmt.Println("submitreplica0: ", cm.BalanceOf(au.From))
	tx, err := fi.AddReplica(au, rb, _pi, _pri, pf)
	if err != nil {
		t.Fatal(err)
	}
	err = cm.CheckTx(tx.Hash())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("submitreplica1: ", cm.BalanceOf(au.From))
}

func TestReward(t *testing.T) {
	sk, addr := makeAccount()
	valt := big.NewInt(1e18)
	val := big.NewInt(1e14)
	err := testtransfer(addr, val, valt)
	if err != nil {
		t.Fatal(err)
	}

	cm, err := NewContractManage(sk, build.BNBTestnetV2)
	if err != nil {
		t.Fatal(err)
	}
	ri, err := cm.NewReward(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	au, err := cm.MakeAuth()
	if err != nil {
		t.Fatal(err)
	}
	for i := uint64(64); i < 3000; i += 64 {
		tx, err := ri.Mint(au, i)
		if err != nil {
			t.Fatal(i, err)
		}

		err = cm.CheckTx(tx.Hash())
		if err != nil {
			t.Fatal(i, err)
		}
	}
}

func TestTotalReward(t *testing.T) {
	cm, err := NewContractManage(nil, build.BNBTestnetV2)
	if err != nil {
		t.Fatal(err)
	}
	ri, err := cm.NewReward(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	ts, err := ri.TotalShare(&bind.CallOpts{From: com.Base})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ts)
	cur, err := ri.Current(&bind.CallOpts{From: com.Base})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cur)
	val, err := ri.Rest(&bind.CallOpts{From: com.Base})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(val)
	rsi, err := ri.GetEReward(&bind.CallOpts{From: com.Base}, cur-1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsi)
	t.Fatal()
}

func TestNodeCheck(t *testing.T) {
	psk, paddr := makeAccount()
	cm, err := NewContractManage(psk, build.BNBTestnetV2)
	if err != nil {
		t.Fatal(err)
	}
	err = testtransfer(paddr, big.NewInt(1e14), big.NewInt(2e18))
	if err != nil {
		t.Fatal(err)
	}

	err = cm.RegisterNode(1, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = cm.CheckNode(paddr, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal("")
}

func testtransfer(to common.Address, val, valt *big.Int) error {
	skbyte, err := os.ReadFile("/tmp/sk")
	if err != nil {
		return err
	}
	hexSk := string(skbyte[:64])
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		return err
	}
	cm, err := NewContractManage(sk, build.BNBTestnetV2)
	if err != nil {
		return err
	}

	err = cm.Transfer(to, val)
	if err != nil {
		return err
	}
	return cm.TransferToken(to, valt)
}

func makeAccount() (*ecdsa.PrivateKey, common.Address) {
	privk, err := crypto.GenerateKey()
	if err != nil {
		panic("generate")
	}
	publicKey := privk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("public")
	}

	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("generate: ", addr)
	//skbyte := crypto.FromECDSA(privk)
	//common.Bytes2Hex(skbyte)

	return privk, addr
}

func TestBlock(t *testing.T) {
	sk, _ := makeAccount()
	cm, err := NewContractManage(sk, build.BNBTestnetV2)
	if err != nil {
		t.Fatal(err)
	}

	latest, err := cm.GetBlockNumber()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(latest)

	latest, err = cm.GetEpoch()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(latest)
	eb, _, err := cm.GetEpochInfo(latest)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(eb.Uint64())

	pi, err := cm.NewPiece(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	pcu, err := pi.Current(&bind.CallOpts{From: com.Base})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(pcu)
}

func TestChoose(t *testing.T) {
	sk, addr := makeAccount()
	cm, err := NewContractManage(sk, build.BNBTestnetV2)
	if err != nil {
		t.Fatal(err)
	}

	seed := [32]byte{100}
	count := 10
	pcnt := 0
	for i := 0; i < 20; i++ {
		ci, err := cm.choose(addr, seed, uint64(count), uint64(pcnt), uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		nci := Choose(addr, seed, uint64(count), uint64(pcnt), uint64(i))
		if ci != nci {
			t.Fatal("unequal at:", i, ci, nci)
		}
	}
}

func TestOrder(t *testing.T) {
	sk, _ := makeAccount()
	cm, err := NewContractManage(sk, build.BNBTestnetV2)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		count := rand.Uint64() / 2
		pcnt := rand.Int63n(int64(count))
		ci, err := cm.getOrder(uint64(count), uint64(pcnt))
		if err != nil {
			t.Fatal(err)
		}
		nci, _ := GetOrder(uint64(count), uint64(pcnt))
		if ci != nci {
			t.Fatal("unequal at:", i, ci, nci)
		}
	}
}
