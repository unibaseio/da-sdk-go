package bls

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/lib/utils"

	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	"github.com/consensys/gnark-crypto/ecc/bw6-761/fr/mimc"
)

func TestKZG(t *testing.T) {
	pk := GenKZGKey(MaxShard/8, big.NewInt(12345678))

	pkb := pk.Serialize()
	vkb := pk.VerifyKey().Serialize()

	npk, err := NewKZGPublicKey(pkb)
	if err != nil {
		t.Fatal(err)
	}
	nvk := new(VerifyKey)
	err = nvk.Deserialize(vkb)
	if err != nil {
		t.Fatal(err)
	}

	rnd := utils.RandomBytes(32)
	data := utils.RandomBytes(30)

	r := bytes.NewReader(data)

	fi, err := os.Create("/tmp/testfile2")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()

	h := sha256.New()
	mw := io.MultiWriter(fi, h)

	tee := io.TeeReader(r, mw)

	c, err := pk.GenCommitments(UnPadSize, tee)
	if err != nil {
		t.Fatal(err)
	}

	if len(c) == 0 {
		t.Fatal("empty commitment")
	}

	rp := bytes.NewReader(data)
	ic := NewChallenge(rnd)

	pr, err := npk.GenProofs(ic, UnPadSize, rp)
	if err != nil {
		t.Fatal(err)
	}

	if len(pr) == 0 {
		t.Fatal("empty proof")
	}

	for _, com := range c {
		err := ic.Add(com)
		if err != nil {
			t.Fatal(err)
		}
	}

	if len(pr) > 1 {
		for i := 1; i < len(pr); i++ {
			pr[0].Add(pr[i])
		}
	}

	err = nvk.VerifyProof(ic, pr[0])
	if err != nil {
		t.Fatal(err)
	}

	t.Log("data sha256: ", hex.EncodeToString(h.Sum(nil)))
}

func TestProof4(t *testing.T) {
	pk := GenKZGKey(MaxShard, big.NewInt(123456789))
	vk := &VerifyKey{
		&pk.Vk,
	}
	ic, ip := genMoveProof(pk)
	nt := time.Now()
	err := vk.VerifyProof(ic, ip)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("move proof:", time.Since(nt))
}

func genMoveProof(pk *PublicKey) (types.IChallenge, types.IProof) {
	fileSize := MaxSize / 4

	rnd := GenRandom(32)
	data := GenRandom(fileSize)

	offset1 := MaxShard / 2
	offset2 := MaxShard * 3 / 4
	ic := NewChallenge(rnd, offset1, offset2)

	pf, err := pk.GenProof(ic, UnPadSize, data)
	if err != nil {
		panic(err)
	}

	c1, err := pk.GenCommitment(UnPadSize, data, 0)
	if err != nil {
		panic(err)
	}

	c2, err := pk.GenCommitment(UnPadSize, data, offset1)
	if err != nil {
		panic(err)
	}

	c3, err := pk.GenCommitment(UnPadSize, data, offset2)
	if err != nil {
		panic(err)
	}

	err = ic.Add(c1)
	if err != nil {
		panic(err)
	}

	err = ic.Add(c2)
	if err != nil {
		panic(err)
	}

	err = ic.Add(c3)
	if err != nil {
		panic(err)
	}

	return ic, pf
}

func TestKZGDec(t *testing.T) {
	pk := GenKZGKey(MaxShard, big.NewInt(12345678))

	rnd := utils.RandomBytes(32)
	data := utils.RandomBytes(1 * MaxSize)

	nt := time.Now()
	cval, err := pk.GenCommitment(MaxShard, rnd, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("cost:", time.Since(nt))

	c, err := pk.GenCommitment(UnPadSize, data, 0)
	if err != nil {
		t.Fatal(err)
	}

	c.Add(cval)
	ic := NewChallenge(rnd)
	ic.Add(c)

	pdata := Pad(data)
	err = Encrypt(pdata, rnd)
	if err != nil {
		t.Fatal(err)
	}
	pr, err := pk.GenProof(ic, 32, pdata)
	if err != nil {
		t.Fatal(err)
	}

	err = pk.VerifyKey().VerifyProof(ic, pr)
	if err != nil {
		t.Fatal(err)
	}

	err = Decrypt(pdata, rnd)
	if err != nil {
		t.Fatal(err)
	}
	updata, err := Unpad(pdata)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, updata) {
		t.Fatal("unequal dec")
	}
}

func TestKZGSloth(t *testing.T) {
	opk := GenKZGKey(MaxShard, big.NewInt(12345678))

	pkb := opk.Serialize()

	pk, err := NewKZGPublicKey(pkb)
	if err != nil {
		t.Fatal(err)
	}

	rnd := utils.RandomBytes(32)
	data := utils.RandomBytes(1 * MaxSize)

	pdata := make([]byte, MaxSize)
	copy(pdata, data)
	pdata = Pad(pdata)

	err = SlothEncode(pdata, rnd)
	if err != nil {
		t.Fatal(err)
	}

	c, err := pk.GenCommitment(PadSize, pdata, 0)
	if err != nil {
		t.Fatal(err)
	}

	cser := c.Serialize()

	cnew, err := NewCommitment(cser)
	if err != nil {
		t.Fatal(err)
	}
	ic := NewChallenge(rnd)
	err = ic.Add(cnew)
	if err != nil {
		t.Fatal(err)
	}
	pr, err := pk.GenProof(ic, PadSize, pdata)
	if err != nil {
		t.Fatal(err)
	}

	err = pk.VerifyKey().VerifyProof(ic, pr)
	if err != nil {
		t.Fatal(err)
	}

	err = SlothDecode(pdata, rnd)
	if err != nil {
		t.Fatal(err)
	}
	updata, err := Unpad(pdata)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, updata) {
		t.Fatal("unequal dec")
	}
}

func TestPad(t *testing.T) {
	fsize := 3*MaxSize + 28
	data := utils.RandomBytes(fsize)
	nt := time.Now()
	end := Pad(data)
	t.Log("pad cost: ", time.Since(nt), len(end))
	dec, err := Unpad(end)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("unpad cost: ", time.Since(nt), len(dec))

	if !bytes.Equal(data, dec[:fsize]) {
		t.Fatal("unequal data")
	}
	t.Fatal()
}

func TestEncrypt(t *testing.T) {
	fsize := 3*MaxSize + 28
	data := utils.RandomBytes(fsize)
	rnd := utils.RandomBytes(32)
	nt := time.Now()
	var fr_r Fr
	fr_r.SetBytes(rnd)
	fr_big := new(big.Int)
	fr_r.BigInt(fr_big)
	fr_r.Exp(fr_r, fr_big)
	t.Log("exp cost: ", time.Since(nt))

	pdata := Pad(data)

	enc := make([]byte, len(pdata))
	copy(enc, pdata)
	err := Encrypt(enc, rnd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("enc cost: ", time.Since(nt))

	err = Decrypt(enc, rnd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dec cost: ", time.Since(nt))
	if !bytes.Equal(pdata, enc) {
		t.Fatal("unequal data")
	}

	updata, err := Unpad(enc)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, updata) {
		t.Fatal("unequal data")
	}
	t.Fatal()
}

func TestSloth(t *testing.T) {
	mod := fr.Modulus()
	mod1 := new(big.Int).Sub(mod, big.NewInt(1))
	t.Log(mod1.Text(16))
	big_v := big.NewInt(17)
	big_v.ModInverse(big_v, mod1)
	t.Log(big_v.Text(16))
	big_v.SetInt64(11)
	big_v.ModInverse(big_v, mod1)
	t.Log(big_v.Text(16))

	rnd := utils.RandomBytes(32)
	data := make([]byte, 32*1024*1024)
	var fr_r Fr
	for i := 0; i < 1024*1024; i++ {
		fr_r.SetRandom()
		copy(data[i*32:(i+1)*32], fr_r.Marshal())
	}

	enc := make([]byte, len(data))
	copy(enc, data)
	nt := time.Now()
	err := SlothEncode(enc, rnd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("enc cost: ", time.Since(nt))
	nt = time.Now()
	err = SlothDecode(enc, rnd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dec cost: ", time.Since(nt))
	if !bytes.Equal(data, enc) {
		t.Fatal("unequal")
	}
	t.Fatal()
}

func TestSlothV2(t *testing.T) {
	rnd := utils.RandomBytes(32)
	data := make([]byte, 32*1024*1024)
	var fr_r Fr
	for i := 0; i < 1024*1024; i++ {
		fr_r.SetRandom()
		copy(data[i*32:(i+1)*32], fr_r.Marshal())
	}

	enc := make([]byte, len(data))
	copy(enc, data)
	nt := time.Now()
	err := SlothEncodeV2(enc, rnd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("enc v2 cost: ", time.Since(nt))
	nt = time.Now()
	err = SlothDecodeV2(enc, rnd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dec v2 cost: ", time.Since(nt))
	if !bytes.Equal(data, enc) {
		t.Fatal("unequal")
	}
	t.Fatal()
}

func TestPoint(t *testing.T) {
	fmt.Println("test")
	nt := time.Now()
	pk := GenPointKey(MaxShard/4, big.NewInt(12345678))
	fmt.Println("gen cost: ", time.Since(nt))
	nt = time.Now()
	pkb := pk.Serialize()
	t.Log("ser cost: ", time.Since(nt))
	nt = time.Now()
	npk := new(PointPublicKey)
	err := npk.Deserialize(pkb)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("der cost: ", time.Since(nt))

	vkb := pk.VerifyKey().Serialize()
	nvk := new(PointVerifyKey)
	err = nvk.Deserialize(vkb)
	if err != nil {
		t.Fatal(err)
	}

	data := utils.RandomBytes(MaxSize / 256)

	com, err := pk.GenCommitment(UnPadSize, data)
	if err != nil {
		t.Fatal(err)
	}

	for rnd := 0; rnd < MaxShard/256; rnd++ {
		ic := NewPointChallenge(rnd)
		pf, err := npk.GenProof(ic, UnPadSize, data)
		if err != nil {
			t.Fatal(rnd, err)
		}
		ic.Add(com)

		err = nvk.VerifyProof(ic, pf)
		if err != nil {
			t.Fatal(rnd, err)
		}
	}
}

func TestMultiPoint(t *testing.T) {
	fmt.Println("test")
	nt := time.Now()
	pk := GenPointKey(MaxShard/128, big.NewInt(12345678))
	fmt.Println("gen cost: ", time.Since(nt))

	data := utils.RandomBytes(MaxSize / 256)

	com, err := pk.GenCommitment(UnPadSize, data)
	if err != nil {
		t.Fatal(err)
	}

	rnds := make([]int, 10)
	for i := 0; i < len(rnds); i++ {
		rnds[i] = rand.Intn(MaxShard / 256)
	}
	mp, err := pk.GenMultiProof(rnds, UnPadSize, data)
	if err != nil {
		t.Fatal(err)
	}

	var commit G1
	commit.SetBytes(com.Raw())
	err = pk.ToVerifyKey().VerifyMultiProof(commit, rnds, mp)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMarshalSpeed(t *testing.T) {
	ppk := GenPointKey(20240, big.NewInt(895576))
	rawres := make([]byte, 0, (ppk.N+1)*G2Size*2)
	res := make([]byte, 0, (ppk.N+1)*G2Size)

	nt := time.Now()
	for i := 0; i < ppk.N+1; i++ {
		tmp := ppk.G2[i].RawBytes()
		rawres = append(rawres, tmp[:]...)
	}
	t.Log("raw marshal cost: ", time.Since(nt))

	nt = time.Now()
	for i := 0; i < ppk.N+1; i++ {
		tmp := ppk.G2[i].Bytes()
		res = append(res, tmp[:]...)
	}
	t.Log("compress marshal cost: ", time.Since(nt))

	nt = time.Now()
	for i := 0; i < ppk.N+1; i++ {
		ppk.G2[i].SetBytes(rawres[i*G2Size*2 : (i+1)*G2Size*2])
	}
	t.Log("raw unmarsal cost: ", time.Since(nt))

	nt = time.Now()
	for i := 0; i < ppk.N+1; i++ {
		ppk.G2[i].SetBytes(res[i*G2Size : (i+1)*G2Size])
	}
	t.Log("compress unmarsal cost: ", time.Since(nt))
	t.Fatal()
}

func TestHash(t *testing.T) {
	nt := time.Now()
	//h, err := recursion.NewShort(curveID.ScalarField(), fr.Modulus())
	//if err != nil {
	//	t.Fatal(err)
	//}
	h := mimc.NewMiMC()

	fieldSize := h.BlockSize()
	t.Log(fieldSize, h.Size())
	var value fr.Element
	bigval := new(big.Int)
	for i := 0; i < 1024*256; i++ {
		h.Reset()
		value.SetRandom()
		value.BigInt(bigval)

		bval := bigval.Bytes()
		bval = append(make([]byte, fieldSize-len(bval)), bval...)
		h.Write(bval)
		h.Write(bval)
		h.Sum(nil)
	}

	t.Fatal("cost: ", time.Since(nt))
}

func GenRandom(len int) []byte {
	res := make([]byte, len)
	for i := 0; i < len; i += 7 {
		val := rand.Int63()
		for j := 0; i+j < len && j < 7; j++ {
			res[i+j] = byte(val)
			val >>= 8
		}
	}
	return res
}

// 4M: 10s vs 100 ms
func TestReverse(t *testing.T) {
	var fr, v Fr
	fr.SetRandom()
	v.SetRandom()
	nt := time.Now()
	for i := 0; i < 4*1024*1024; i++ {
		v.Mul(&v, &fr)
		//v.Div(&v, &fr)
		fr.Add(&fr, &v)
	}

	t.Fatal("cost: ", time.Since(nt))
}

// 10s, 320ms
func TestSlothV3(t *testing.T) {
	rnd := utils.RandomBytes(32)
	data := make([]byte, 32*4*1024*1024)
	var fr_r Fr
	for i := 0; i < 4*1024*1024; i++ {
		fr_r.SetRandom()
		copy(data[i*32:(i+1)*32], fr_r.Marshal())
	}

	enc := make([]byte, len(data))
	copy(enc, data)
	nt := time.Now()
	err := SlothEncodeV3(enc, rnd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("enc v3 cost: ", time.Since(nt))
	nt = time.Now()
	err = SlothDecodeV3(enc, rnd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dec v3 cost: ", time.Since(nt))
	if !bytes.Equal(data, enc) {
		t.Fatal("unequal")
	}
	t.Fatal()
}

func TestSlothV4(t *testing.T) {
	rnd := utils.RandomBytes(32)
	var fr_r Fr
	fr_r.SetBytes(rnd)
	total := 32 * 1024 * 1024
	shards := make([]Fr, total)
	shards[0].SetOne()
	nt := time.Now()
	for i := 1; i < total; i++ {
		shards[i].Mul(&shards[i-1], &fr_r)
	}
	t.Log("v4 cost: ", time.Since(nt))

	var point Fr
	point.SetRandom()
	nt = time.Now()
	val := Eval(shards, point)
	t.Log("v4 eval cost: ", time.Since(nt))

	nt = time.Now()
	fr_r.Mul(&fr_r, &point)
	var nval, tmp Fr
	nval.Exp(fr_r, big.NewInt(int64(total)))
	tmp.SetOne()
	nval.Sub(&nval, &tmp)
	tmp.Sub(&fr_r, &tmp)
	nval.Div(&nval, &tmp)
	t.Log("v4 eval1 cost: ", time.Since(nt))
	if !nval.Equal(&val) {
		t.Log("unequal")
	}

	nt = time.Now()
	h := Divide(shards, point)
	t.Log("v4 divide cost: ", time.Since(nt))

	nt = time.Now()
	re := Mul(h, point)
	t.Log("v4 mul cost: ", time.Since(nt))
	re[0].Add(&re[0], &val)

	for i := 1; i < total; i++ {
		if !shards[i].Equal(&re[i]) {
			t.Fatal("not equal at: ", i)
		}
	}

	t.Fatal()
}

func TestMarshal(t *testing.T) {
	ew := NewEncodeWitness(6, 4)
	ew.H.ScalarMultiplicationBase(big.NewInt(10))
	ew.ClaimedValues[0].SetRandom()
	ewb := ew.Serialize()
	new := new(EncodeWitness)
	err := new.Deserialize(ewb)
	if err != nil {
		t.Fatal(err)
	}
}
