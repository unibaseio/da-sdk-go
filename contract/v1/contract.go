package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	com "github.com/unibaseio/da-sdk-go/contract/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractManage struct {
	Type       string
	ChainID    *big.Int
	RPC        string
	BankAddr   common.Address
	TokenAddr  common.Address
	SyncHeight int
	sk         *ecdsa.PrivateKey
}

func NewContractManage(sk *ecdsa.PrivateKey, chainType string) (*ContractManage, error) {
	cm := &ContractManage{
		Type: chainType,
		sk:   sk,
	}

	switch chainType {
	case com.OPSepolia:
		cm.RPC = com.OPSepoliaChainRPC
		cm.ChainID = big.NewInt(int64(com.OPSepoliaChainID))
		cm.BankAddr = com.OPSepoliaBankAddr
		cm.TokenAddr = com.OPSepoliaTokenAddr
		cm.SyncHeight = com.OPSepoliaSyncHeight
	case com.OPBNBTestnet:
		cm.RPC = com.OPBNBTestnetChainRPC
		cm.ChainID = big.NewInt(int64(com.OPBNBTestnetChainID))
		cm.BankAddr = com.OPBNBTestnetBankAddr
		cm.TokenAddr = com.OPBNBTestnetTokenAddr
		cm.SyncHeight = com.OPBNBTestnetSyncHeight
	case com.BNBTestnetV2:
		cm.RPC = com.BNBTestnetChainRPC
		cm.ChainID = big.NewInt(int64(com.BNBTestnetChainID))
		cm.BankAddr = com.BNBTestnetBankAddr
		cm.TokenAddr = com.BNBTestnetTokenAddr
		cm.SyncHeight = com.BNBTestnetSyncHeight
		chainRPC := os.Getenv("CHAIN_RPC_97")
		if chainRPC != "" {
			cm.RPC = chainRPC
		}
	default:
		return nil, fmt.Errorf("unsupportted chain type: %s, use 'bnb-testnet', 'op-sepolia' or 'opbnb-testnet'", chainType)
	}

	chainRPC := os.Getenv("CHAIN_RPC")
	if chainRPC != "" {
		cm.RPC = chainRPC
	}

	// check chain RPC is connected
	// check chain id
	client, err := ethclient.Dial(cm.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

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
