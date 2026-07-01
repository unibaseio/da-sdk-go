package metering

import (
	"math/big"
	"os"
	"strconv"
)

// Config holds all metering settings. It is loaded from the environment and is
// designed so that the zero value / default (Enabled=false) leaves all existing
// hub behavior unchanged.
type Config struct {
	Enabled      bool
	ChargeWrites bool
	ChargeReads  bool

	WriteBaseWei      *big.Int
	WritePerKBWei     *big.Int
	ReadPerRequestWei *big.Int

	DefaultCreditLimitWei *big.Int
	CheckChain            bool
	SettlementMode        string // offchain, erc8183

	AutoSettle         bool
	SettleIntervalSec  int
	SettleThresholdWei *big.Int

	ProviderAddress      string
	ProviderPrivateKey   string
	ERC20TokenAddr       string
	ERC8183ContractAddr  string
	ERC8183EvaluatorAddr string
	ChainRPCURL          string
	ChainID              int64
}

// LoadConfigFromEnv reads HUB_METERING_* and related env vars. Metering is
// disabled by default; only HUB_METERING_ENABLED=true turns it on.
func LoadConfigFromEnv() Config {
	return Config{
		Enabled:      envBool("HUB_METERING_ENABLED", false),
		ChargeWrites: envBool("HUB_METERING_CHARGE_WRITES", true),
		ChargeReads:  envBool("HUB_METERING_CHARGE_READS", false),

		WriteBaseWei:      envBigInt("HUB_METERING_WRITE_BASE_WEI", big.NewInt(0)),
		WritePerKBWei:     envBigInt("HUB_METERING_WRITE_PER_KB_WEI", big.NewInt(0)),
		ReadPerRequestWei: envBigInt("HUB_METERING_READ_PER_REQUEST_WEI", big.NewInt(0)),

		DefaultCreditLimitWei: envBigInt("HUB_METERING_DEFAULT_CREDIT_LIMIT_WEI", big.NewInt(0)),
		CheckChain:            envBool("HUB_METERING_CHECK_CHAIN", false),
		SettlementMode:        envString("HUB_METERING_SETTLEMENT_MODE", "offchain"),

		AutoSettle:         envBool("HUB_METERING_AUTO_SETTLE", false),
		SettleIntervalSec:  int(envInt64("HUB_METERING_SETTLE_INTERVAL_SEC", 300)),
		SettleThresholdWei: envBigInt("HUB_METERING_SETTLE_THRESHOLD_WEI", big.NewInt(0)),

		ProviderAddress:      envString("HUB_PROVIDER_ADDRESS", ""),
		ProviderPrivateKey:   envString("HUB_PROVIDER_PRIVATE_KEY", ""),
		ERC20TokenAddr:       envString("HUB_ERC20_TOKEN_ADDR", ""),
		ERC8183ContractAddr:  envString("HUB_ERC8183_CONTRACT_ADDR", ""),
		ERC8183EvaluatorAddr: envString("HUB_ERC8183_EVALUATOR_ADDR", ""),
		ChainRPCURL:          envString("HUB_CHAIN_RPC_URL", ""),
		ChainID:              envInt64("HUB_CHAIN_ID", 97),
	}
}

func envBool(key string, fallback bool) bool {
	if v := os.Getenv(key); v != "" {
		b, err := strconv.ParseBool(v)
		if err == nil {
			return b
		}
	}
	return fallback
}

func envString(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envInt64(key string, fallback int64) int64 {
	if v := os.Getenv(key); v != "" {
		n, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return n
		}
	}
	return fallback
}

// envBigInt parses a base-10 wei string. On empty or malformed input it returns
// a copy of fallback so money is never silently corrupted by a float.
func envBigInt(key string, fallback *big.Int) *big.Int {
	if v := os.Getenv(key); v != "" {
		n, ok := new(big.Int).SetString(v, 10)
		if ok {
			return n
		}
	}
	return new(big.Int).Set(fallback)
}
