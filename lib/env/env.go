// Package env centralizes environment-variable access for the DA stack:
// typed getters with defaults (no more scattered os.Getenv + ad-hoc parsing)
// and one place that names every managed variable (no string typos).
//
// Parsing failures fall back to the default and warn via the standard log
// (NOT lib/log — that package reads LOG_LEVEL through here, so importing it
// would be a cycle).
package env

import (
	"log"
	"os"
	"strconv"
	"time"
)

// Managed environment variable names — the single source of truth. Values are
// unchanged from the historical strings, so existing deployments keep working.
const (
	ChainType      = "CHAIN_TYPE"       // chain selection (build.CheckChain / ContractManage)
	ChainRPC       = "CHAIN_RPC"        // generic RPC override (tx/call)
	ChainRPCFilter = "CHAIN_RPC_FILTER" // generic FilterLogs RPC override
	// per-chain overrides are CHAIN_RPC_<id> / CHAIN_RPC_FILTER_<id>; build them
	// with ChainRPC+"_"+id etc.

	GasLimit = "GAS_LIMIT"
	GasPrice = "GAS_PRICE"
	GasTip   = "GAS_TIP"

	LogLevel = "LOG_LEVEL"
	LogFile  = "LOG_FILE"

	NodePassword = "NODE_PASSWORD"
	ExposeURL    = "EXPOSE_URL"
	SyncHeight   = "SYNC_HEIGHT"

	StreamUploadConcurrency = "STREAM_UPLOAD_CONCURRENCY"

	// gateway backend split
	GatewayMode      = "GATEWAY_MODE"
	GatewayDBDriver  = "GATEWAY_DB_DRIVER"
	GatewayDBDSN     = "GATEWAY_DB_DSN"
	GatewayRedisAddr = "GATEWAY_REDIS_ADDR"

	// service endpoints (hardcoded defaults, env-overridable for deployments)
	ServerURL   = "SERVER_URL"   // default DA server the SDK talks to
	ParamSource = "PARAM_SOURCE" // base URL for the shared SRS/param download
	HubURL      = "HUB_URL"      // default hub server
)

// Str returns the env value or def if unset/empty.
func Str(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// Int returns the env value parsed as int, or def on unset/parse error.
func Int(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		log.Printf("env: invalid int %s=%q, using default %d", key, v, def)
		return def
	}
	return n
}

// Int64 returns the env value parsed as int64, or def on unset/parse error.
func Int64(key string, def int64) int64 {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Printf("env: invalid int64 %s=%q, using default %d", key, v, def)
		return def
	}
	return n
}

// Float returns the env value parsed as float64, or def on unset/parse error.
func Float(key string, def float64) float64 {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Printf("env: invalid float %s=%q, using default %v", key, v, def)
		return def
	}
	return f
}

// Bool returns the env value parsed as bool (1/true/yes/on), or def if unset.
func Bool(key string, def bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Printf("env: invalid bool %s=%q, using default %v", key, v, def)
		return def
	}
	return b
}

// Duration returns the env value parsed as a Go duration, or def on error.
func Duration(key string, def time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		log.Printf("env: invalid duration %s=%q, using default %v", key, v, def)
		return def
	}
	return d
}
