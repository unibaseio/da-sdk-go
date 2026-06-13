package build

import "github.com/unibaseio/da-sdk-go/lib/env"

const (
	ServerURL = "http://54.251.11.180:8080"

	// mainnet chains (deploy order: BASE first, then BSC, then ETH)
	ETHMainnet  = "eth-mainnet"
	BSCMainnet  = "bsc-mainnet"
	BaseMainnet = "base"

	// testnets
	BaseSepolia   = "base-sepolia"
	LocalAnvil    = "local-anvil"
	BNBTestnetV2  = "bnb-testnet-v2"
	BNBTestnetDAO = "bnb-testnet-dao"
)

const chainTypeHint = "please set env 'CHAIN_TYPE' to one of 'base', 'base-sepolia', 'bsc-mainnet', 'eth-mainnet', 'bnb-testnet-v2' or 'bnb-testnet-dao'"

func CheckChain() string {
	ct := env.Str(env.ChainType, "")
	if ct == "" {
		panic(chainTypeHint)
	}
	switch ct {
	case BaseMainnet, BaseSepolia, BSCMainnet, ETHMainnet, BNBTestnetV2, BNBTestnetDAO:
		return ct
	default:
		panic(chainTypeHint)
	}
}
