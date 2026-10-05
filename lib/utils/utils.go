package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"os/exec"
	"runtime"
	"strconv"

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

func HexToAddress(addr string) common.Address {
	return common.HexToAddress(addr)
}

func RandomBytes(length int) []byte {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)

	return randomBytes
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

func KillProcess(pid string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("kill", "-15", pid).Run()
	case "windows":
		return exec.Command("taskkill", "/F", "/T", "/PID", pid).Run()
	default:
		return fmt.Errorf("unsupported platform %s", runtime.GOOS)
	}
}
