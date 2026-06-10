package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"sync"

	com "github.com/unibaseio/da-sdk-go/contract/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractManage struct {
	Type            string
	ChainID         *big.Int
	RPC             string
	RPCForFilterLog string
	SyncHeight      int
	sk              *ecdsa.PrivateKey

	TokenAddr common.Address

	EpochAddr   common.Address
	NodeAddr    common.Address
	PieceAddr   common.Address
	RSProofAddr common.Address
	EProofAddr  common.Address
	EVerifyAddr common.Address
	StatAddr    common.Address

	RSOneAddr common.Address
	KZGAddr   common.Address
	AddAddr   common.Address
	MulAddr   common.Address

	// shared RPC client for c.RPC, created lazily and reused by all contract
	// bindings/calls (P2: one client instead of a dial per call)
	clientMu sync.Mutex
	client   *ethclient.Client
}

// Client returns the shared ethclient for c.RPC, dialing it on first use.
// The endpoints are HTTP, so the client needs no liveness management —
// each request rides Go's pooled HTTP transport.
func (c *ContractManage) Client(ctx context.Context) (*ethclient.Client, error) {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if c.client != nil {
		return c.client, nil
	}
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	c.client = client
	return client, nil
}

func NewContractManage(sk *ecdsa.PrivateKey, chainType string) (*ContractManage, error) {
	cm := &ContractManage{
		Type: chainType,
		sk:   sk,
	}

	switch chainType {
	case com.BaseSepolia:
		cm.RPC = com.BaseSepoliaChainRPC
		cm.RPCForFilterLog = com.BaseSepoliaChainRPCForFilterLog
		cm.ChainID = big.NewInt(com.BaseSepoliaChainID)
		cm.SyncHeight = com.BaseSepoliaSyncHeight

		cm.TokenAddr = com.BaseSepoliaTokenAddr

		cm.EpochAddr = com.BaseSepoliaEpochAddr
		cm.NodeAddr = com.BaseSepoliaNodeAddr
		cm.PieceAddr = com.BaseSepoliaPieceAddr
		cm.RSProofAddr = com.BaseSepoliaRSProofAddr
		cm.EProofAddr = com.BaseSepoliaEProofAddr
		cm.EVerifyAddr = com.BaseSepoliaEVerifyAddr
		cm.StatAddr = com.BaseSepoliaStatAddr

		cm.RSOneAddr = com.BaseSepoliaRSOneAddr
		cm.KZGAddr = com.BaseSepoliaKZGAddr
		cm.AddAddr = com.BaseSepoliaAddAddr
		cm.MulAddr = com.BaseSepoliaMulAddr
	case com.BaseMainnet:
		cm.RPC = com.BaseMainnetChainRPC
		cm.RPCForFilterLog = com.BaseMainnetChainRPCForFilterLog
		cm.ChainID = big.NewInt(com.BaseMainnetChainID)
		cm.SyncHeight = com.BaseMainnetSyncHeight

		cm.TokenAddr = com.BaseMainnetTokenAddr

		cm.EpochAddr = com.BaseMainnetEpochAddr
		cm.NodeAddr = com.BaseMainnetNodeAddr
		cm.PieceAddr = com.BaseMainnetPieceAddr
		cm.RSProofAddr = com.BaseMainnetRSProofAddr
		cm.EProofAddr = com.BaseMainnetEProofAddr
		cm.EVerifyAddr = com.BaseMainnetEVerifyAddr
		cm.StatAddr = com.BaseMainnetStatAddr

		cm.RSOneAddr = com.BaseMainnetRSOneAddr
		cm.KZGAddr = com.BaseMainnetKZGAddr
		cm.AddAddr = com.BaseMainnetAddAddr
		cm.MulAddr = com.BaseMainnetMulAddr
	case com.BSCMainnet:
		cm.RPC = com.BSCMainnetChainRPC
		cm.RPCForFilterLog = com.BSCMainnetChainRPCForFilterLog
		cm.ChainID = big.NewInt(com.BSCMainnetChainID)
		cm.SyncHeight = com.BSCMainnetSyncHeight

		cm.TokenAddr = com.BSCMainnetTokenAddr

		cm.EpochAddr = com.BSCMainnetEpochAddr
		cm.NodeAddr = com.BSCMainnetNodeAddr
		cm.PieceAddr = com.BSCMainnetPieceAddr
		cm.RSProofAddr = com.BSCMainnetRSProofAddr
		cm.EProofAddr = com.BSCMainnetEProofAddr
		cm.EVerifyAddr = com.BSCMainnetEVerifyAddr
		cm.StatAddr = com.BSCMainnetStatAddr

		cm.RSOneAddr = com.BSCMainnetRSOneAddr
		cm.KZGAddr = com.BSCMainnetKZGAddr
		cm.AddAddr = com.BSCMainnetAddAddr
		cm.MulAddr = com.BSCMainnetMulAddr
	case com.ETHMainnet:
		cm.RPC = com.ETHMainnetChainRPC
		cm.RPCForFilterLog = com.ETHMainnetChainRPCForFilterLog
		cm.ChainID = big.NewInt(com.ETHMainnetChainID)
		cm.SyncHeight = com.ETHMainnetSyncHeight

		cm.TokenAddr = com.ETHMainnetTokenAddr

		cm.EpochAddr = com.ETHMainnetEpochAddr
		cm.NodeAddr = com.ETHMainnetNodeAddr
		cm.PieceAddr = com.ETHMainnetPieceAddr
		cm.RSProofAddr = com.ETHMainnetRSProofAddr
		cm.EProofAddr = com.ETHMainnetEProofAddr
		cm.EVerifyAddr = com.ETHMainnetEVerifyAddr
		cm.StatAddr = com.ETHMainnetStatAddr

		cm.RSOneAddr = com.ETHMainnetRSOneAddr
		cm.KZGAddr = com.ETHMainnetKZGAddr
		cm.AddAddr = com.ETHMainnetAddAddr
		cm.MulAddr = com.ETHMainnetMulAddr
	case com.BNBTestnetV2:
		cm.RPC = com.BNBTestnetChainRPC
		cm.RPCForFilterLog = com.BNBTestnetChainRPCForFilterLog
		cm.ChainID = big.NewInt(int64(com.BNBTestnetChainID))
		cm.SyncHeight = com.BNBTestnetSyncHeight

		cm.TokenAddr = com.BNBTestnetTokenAddr

		cm.EpochAddr = com.BNBTestnetEpochAddr
		cm.NodeAddr = com.BNBTestnetNodeAddr
		cm.PieceAddr = com.BNBTestnetPieceAddr
		cm.RSProofAddr = com.BNBTestnetRSProofAddr
		cm.EProofAddr = com.BNBTestnetEProofAddr
		cm.EVerifyAddr = com.BNBTestnetEVerifyAddr
		cm.StatAddr = com.BNBTestnetStatAddr

		cm.RSOneAddr = com.BNBTestnetRSOneAddr
		cm.KZGAddr = com.BNBTestnetKZGAddr
		cm.AddAddr = com.BNBTestnetAddAddr
		cm.MulAddr = com.BNBTestnetMulAddr
	case com.BNBTestnetDAO:
		cm.RPC = com.BNBTestnetDAOChainRPC
		cm.RPCForFilterLog = com.BNBTestnetDAOChainRPCForFilterLog
		cm.ChainID = big.NewInt(int64(com.BNBTestnetDAOChainID))
		cm.SyncHeight = com.BNBTestnetDAOSyncHeight

		cm.TokenAddr = com.BNBTestnetDAOTokenAddr

		cm.EpochAddr = com.BNBTestnetDAOEpochAddr
		cm.NodeAddr = com.BNBTestnetDAONodeAddr
		cm.PieceAddr = com.BNBTestnetDAOPieceAddr
		cm.RSProofAddr = com.BNBTestnetDAORSProofAddr
		cm.EProofAddr = com.BNBTestnetDAOEProofAddr
		cm.EVerifyAddr = com.BNBTestnetDAOEVerifyAddr
		cm.StatAddr = com.BNBTestnetDAOStatAddr

		cm.RSOneAddr = com.BNBTestnetDAORSOneAddr
		cm.KZGAddr = com.BNBTestnetDAOKZGAddr
		cm.AddAddr = com.BNBTestnetDAOAddAddr
		cm.MulAddr = com.BNBTestnetDAOMulAddr
	case com.LocalAnvil:
		cm.RPC = com.LocalAnvilChainRPC
		cm.RPCForFilterLog = com.LocalAnvilChainRPCForFilterLog
		cm.ChainID = big.NewInt(int64(com.LocalAnvilChainID))
		cm.SyncHeight = com.LocalAnvilSyncHeight

		cm.TokenAddr = com.LocalAnvilTokenAddr

		cm.EpochAddr = com.LocalAnvilEpochAddr
		cm.NodeAddr = com.LocalAnvilNodeAddr
		cm.PieceAddr = com.LocalAnvilPieceAddr
		cm.RSProofAddr = com.LocalAnvilRSProofAddr
		cm.EProofAddr = com.LocalAnvilEProofAddr
		cm.EVerifyAddr = com.LocalAnvilEVerifyAddr
		cm.StatAddr = com.LocalAnvilStatAddr

		cm.RSOneAddr = com.LocalAnvilRSOneAddr
		cm.KZGAddr = com.LocalAnvilKZGAddr
		cm.AddAddr = com.LocalAnvilAddAddr
		cm.MulAddr = com.LocalAnvilMulAddr
	default:
		return nil, fmt.Errorf("unsupported chain type: %s, use 'base', 'base-sepolia', 'bsc-mainnet', 'eth-mainnet', 'bnb-testnet-v2', 'bnb-testnet-dao' or 'local-anvil'", chainType)
	}

	// RPC endpoint env overrides. Per-chain-id keys (CHAIN_RPC_<id> /
	// CHAIN_RPC_FILTER_<id>) avoid collisions when several chains run in one
	// process (e.g. the gateway); the generic CHAIN_RPC / CHAIN_RPC_FILTER are
	// kept as a fallback for single-chain deployments.
	idStr := cm.ChainID.String()
	if v := os.Getenv("CHAIN_RPC_" + idStr); v != "" {
		cm.RPC = v
	} else if v := os.Getenv("CHAIN_RPC"); v != "" {
		cm.RPC = v
	}
	if v := os.Getenv("CHAIN_RPC_FILTER_" + idStr); v != "" {
		cm.RPCForFilterLog = v
	} else if v := os.Getenv("CHAIN_RPC_FILTER"); v != "" {
		cm.RPCForFilterLog = v
	}

	// check chain RPC is connected
	// check chain id; the validated client is kept as the shared client
	client, err := cm.Client(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	if chainID.Cmp(cm.ChainID) != 0 {
		return nil, fmt.Errorf("chain id mismatch, expected %d, got %d", cm.ChainID, chainID)
	}

	com.Logger.Info("connected to chain: ", cm.RPC)

	return cm, nil
}

func (c *ContractManage) MakeAuth() (*bind.TransactOpts, error) {
	return com.MakeAuthBySk(c.RPC, c.ChainID, c.sk)
}

func (c *ContractManage) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	return com.GetTransactionReceipt(c.RPC, hash)
}

func (c *ContractManage) CheckTx(txHash common.Hash) error {
	return com.CheckTx(c.RPC, txHash)
}

func (c *ContractManage) Transfer(toAddr common.Address, value *big.Int) error {
	return com.Transfer(c.RPC, c.sk, toAddr, value)
}

func (c *ContractManage) TransferToken(toAddr common.Address, value *big.Int) error {
	return com.TransferToken(c.RPC, c.ChainID, c.sk, c.TokenAddr, toAddr, value)
}

func (c *ContractManage) BalanceOf(toAddr common.Address) *big.Int {
	return com.BalanceOf(c.RPC, toAddr)
}

func (c *ContractManage) BalanceOfToken(toAddr common.Address) *big.Int {
	return com.BalanceOfToken(c.RPC, c.TokenAddr, toAddr)
}
