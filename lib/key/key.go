package key

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/lib/utils"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

type Key struct {
	ks    *keystore.KeyStore
	local *keystore.Key
}

func New(ksp string) *Key {
	ks := keystore.NewKeyStore(ksp, keystore.StandardScryptN, keystore.StandardScryptP)

	return &Key{
		ks: ks,
	}
}

func (k *Key) Has(addr common.Address) bool {
	return k.ks.HasAddress(addr)
}

func (k *Key) Create(pw string) (common.Address, error) {
	ac, err := k.ks.NewAccount(pw)
	if err != nil {
		return common.Address{}, err
	}
	return ac.Address, nil
}

func (k *Key) ImportECDSA(sk, pw string) (common.Address, error) {
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return common.Address{}, err
	}
	ac, err := k.ks.ImportECDSA(privateKey, pw)
	if err != nil {
		return common.Address{}, err
	}
	return ac.Address, nil
}

func (k *Key) Import(kjson, oldpw, pw string) (common.Address, error) {
	ac, err := k.ks.Import([]byte(kjson), oldpw, pw)
	if err != nil {
		return common.Address{}, err
	}
	return ac.Address, nil
}

func (k *Key) Load(addr common.Address, pw string) error {
	a := accounts.Account{
		Address: addr,
	}

	npw := utils.RandomBytes(32)
	if len(npw) != 32 {
		return fmt.Errorf("generate random fail")
	}

	kjson, err := k.ks.Export(a, pw, string(npw))
	if err != nil {
		return err
	}

	key, err := keystore.DecryptKey(kjson, string(npw))
	if err != nil {
		return err
	}
	k.local = key
	return nil
}

func (k *Key) Export() *keystore.Key {
	return k.local
}

func (k *Key) Sign(msg []byte) ([]byte, error) {
	if k.local == nil {
		return nil, fmt.Errorf("no local account")
	}
	return crypto.Sign(msg, k.local.PrivateKey)
}

func (k *Key) Decrypt(msg []byte) ([]byte, error) {
	if k.local == nil {
		return nil, fmt.Errorf("no local account")
	}

	decrypted, err := ecies.ImportECDSA(k.local.PrivateKey).Decrypt(msg, nil, nil)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

func (k *Key) Public() []byte {
	if k.local == nil {
		panic("no local account")
	}
	return crypto.FromECDSAPub(&k.local.PrivateKey.PublicKey)
}

func (k *Key) Address() common.Address {
	if k.local == nil {
		panic("no local account")
	}
	return k.local.Address
}

func (k *Key) BuildAuth(hash []byte) (types.Auth, error) {
	res := types.Auth{
		Addr: k.Address(),
		Hash: hash,
	}

	h := sha256.New()
	ts := time.Now().Unix()
	h.Write(hash)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(ts))
	h.Write(b)

	sum := h.Sum(nil)

	sign, err := k.Sign(sum)
	if err != nil {
		return res, err
	}

	res.Time = ts
	res.Sign = sign

	return res, nil
}

func BuildAuth(sk *ecdsa.PrivateKey, hash []byte) (types.Auth, error) {
	res := types.Auth{
		Addr: utils.ECDSAToAddr(sk),
		Hash: hash,
	}

	h := sha256.New()
	ts := time.Now().Unix()
	h.Write(hash)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(ts))
	h.Write(b)

	sum := h.Sum(nil)

	sign, err := crypto.Sign(sum, sk)
	if err != nil {
		return res, err
	}

	res.Time = ts
	res.Sign = sign

	return res, nil
}

func BuildAuthLocal(hash []byte) types.Auth {
	ks := keystore.NewKeyStore("/tmp/.dimo", keystore.StandardScryptN, keystore.StandardScryptP)
	laccount, _ := ks.NewAccount("test")

	h := sha256.New()
	ts := time.Now().Unix()
	h.Write(hash)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(ts))
	h.Write(b)

	sum := h.Sum(nil)

	ks.Unlock(laccount, "test")
	sign, err := ks.SignHash(laccount, sum[:])
	if err != nil {
		panic(err)
	}

	return types.Auth{
		Addr: laccount.Address,
		Time: ts,
		Hash: hash,
		Sign: sign,
	}
}
