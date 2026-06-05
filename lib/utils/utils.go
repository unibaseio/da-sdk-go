package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	mrand "math/rand"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/disk"
	"golang.org/x/crypto/sha3"
)

func ECDSAToAddr(sk *ecdsa.PrivateKey) common.Address {
	publicKey := sk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func HexToAddr(sk string) common.Address {
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return common.Address{}
	}

	return ECDSAToAddr(privateKey)
}

func ToEthAddress(pubkey []byte) []byte {
	if len(pubkey) == 65 {
		d := sha3.NewLegacyKeccak256()
		d.Write(pubkey[1:])
		payload := d.Sum(nil)
		return payload[12:]
	}

	return pubkey
}

func GetDiskStatus(path string) (types.DiskStats, error) {
	m := types.DiskStats{
		Path: path,
	}
	dus, err := disk.Usage(path)
	if err != nil {
		return m, err
	}

	m.Total = dus.Total
	m.Free = dus.Free

	m.Used = m.Total - m.Free
	return m, nil
}

func DataToName(d []byte) string {
	dsum := sha256.Sum256(d)
	return hex.EncodeToString(dsum[:])
}

func HexToAddress(addr string) common.Address {
	return common.HexToAddress(addr)
}

func RandomBytes(length int) []byte {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)

	return randomBytes
}

func ShuffleString(array []string) {
	var temp string
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := len(array) - 1; i >= 0; i-- {
		num := r.Intn(i + 1)
		temp = array[i]
		array[i] = array[num]
		array[num] = temp
	}
}

func LocalIp() string {
	address, _ := net.InterfaceAddrs()
	var ip = "localhost"
	for _, address := range address {
		if ipAddress, ok := address.(*net.IPNet); ok && !ipAddress.IP.IsLoopback() {
			if ipAddress.IP.To4() != nil {
				ip = ipAddress.IP.String()
			}
		}
	}
	return ip
}

func Disorder(array []interface{}) {
	var temp interface{}
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := len(array) - 1; i >= 0; i-- {
		num := r.Intn(i + 1)
		temp = array[i]
		array[i] = array[num]
		array[num] = temp
	}
}

func GetGPUInfo(id string) (types.GPUCore, error) {
	res := types.GPUCore{}
	cmd := exec.Command("nvidia-smi", "--query-gpu=name,uuid,memory.total", "--format=csv,noheader", "--id="+id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return res, err
	}

	outs := strings.Trim(string(out), "\n")
	slice := strings.Split(outs, ", ")
	if len(slice) != 3 {
		return res, fmt.Errorf("wrong gpu detail")
	}
	res.Type = slice[0]
	res.Name = slice[1]
	res.Memory = slice[2]
	return res, nil
}

const (
	KiB = 1024
	MiB = 1048576
	GiB = 1073741824
	TiB = 1099511627776

	KB = 1e3
	MB = 1e6
	GB = 1e9
	TB = 1e12

	Wei  = 1
	GWei = 1e9
	Eth  = 1e18
)

func FormatBytes(i int64) (result string) {
	switch {
	case i >= TiB:
		result = fmt.Sprintf("%.02f TiB", float64(i)/TiB)
	case i >= GiB:
		result = fmt.Sprintf("%.02f GiB", float64(i)/GiB)
	case i >= MiB:
		result = fmt.Sprintf("%.02f MiB", float64(i)/MiB)
	case i >= KiB:
		result = fmt.Sprintf("%.02f KiB", float64(i)/KiB)
	default:
		result = fmt.Sprintf("%d B", i)
	}
	return
}

func FormatEth(i *big.Int) string {
	f := new(big.Float).SetInt(i)
	res, _ := f.Float64()
	switch {
	case res >= Eth:
		return fmt.Sprintf("%.02f Eth", res/Eth)
	case res >= GWei:
		return fmt.Sprintf("%.02f Gwei", res/GWei)
	default:
		return fmt.Sprintf("%d Wei", i.Int64())
	}
}

func GetHardwareInfo() types.HardwareInfo {
	res := types.HardwareInfo{}
	ci, err := cpu.Info()
	if err == nil && len(ci) > 0 {
		res.CPU = ci[0].ModelName + ", " + strconv.Itoa(len(ci)) + " Cores"
	}
	vms, err := mem.VirtualMemory()
	if err == nil {
		res.Memory = FormatBytes(int64(vms.Total))
	}
	return res
}
