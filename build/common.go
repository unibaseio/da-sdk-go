package build

import "os"

const (
	ServerURL    = "http://54.251.11.180:8080"
	OPSepolia    = "op-sepolia"
	OPBNBTestnet = "opbnb-testnet"
	//BNBTestnet   = "bnb-testnet-v2"
	LocalAnvil    = "local-anvil"
	BNBTestnetV2  = "bnb-testnet-v2"
	BNBTestnetDAO = "bnb-testnet-dao"
)

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
