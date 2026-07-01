package build

import "os"

const (
	OPSepolia    = "op-sepolia"
	OPBNBTestnet = "opbnb-testnet"
	//BNBTestnet   = "bnb-testnet-v2"
	LocalAnvil    = "local-anvil"
	BNBTestnetV2  = "bnb-testnet-v2"
	BNBTestnetDAO = "bnb-testnet-dao"
)

// ServerURL is the default gateway endpoint that hubs/clients use as their
// remote. Override at runtime with the SERVER_URL env var (e.g. via docker
// compose) so the address isn't baked into the binary.
var ServerURL = func() string {
	if v := os.Getenv("SERVER_URL"); v != "" {
		return v
	}
	return "http://54.251.11.180:8080"
}()

func CheckChain() string {
	ct := os.Getenv("CHAIN_TYPE")
	if ct == "" {
		panic("please set env 'CHAIN_TYPE' to 'op-sepolia', 'opbnb-testnet', 'bnb-testnet-v2' or 'bnb-testnet-dao'")
	}
	switch ct {
	case OPSepolia, OPBNBTestnet, BNBTestnetV2, BNBTestnetDAO:
		return ct
	default:
		panic("please set env 'CHAIN_TYPE' to 'op-sepolia', 'opbnb-testnet', 'bnb-testnet-v2' or 'bnb-testnet-dao'")
	}
}
