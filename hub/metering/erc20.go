package metering

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/token"
)

// erc20API is the subset of ERC-20 operations metering needs. It is an
// interface so write-admission and settlement can be tested without a live RPC
// node by injecting a fake.
type erc20API interface {
	BalanceOf(ctx context.Context, owner common.Address) (*big.Int, error)
	Allowance(ctx context.Context, owner, spender common.Address) (*big.Int, error)
	TransferFrom(ctx context.Context, from, to common.Address, amount *big.Int) (common.Hash, error)
}

// ERC20Client talks to the configured ERC-20 token via the generated binding.
// It dials the RPC per call (like contract/v2 inst.go) so construction does not
// require the node to be reachable. Read calls (BalanceOf/Allowance) need no
// key; TransferFrom uses the provider private key.
type ERC20Client struct {
	rpcURL    string
	tokenAddr common.Address
	chainID   *big.Int
	sk        *ecdsa.PrivateKey // provider key; nil => read-only
}

var _ erc20API = (*ERC20Client)(nil)

// newERC20Client validates config and builds a client. It does NOT dial.
func newERC20Client(cfg Config) (*ERC20Client, error) {
	if cfg.ChainRPCURL == "" {
		return nil, fmt.Errorf("HUB_CHAIN_RPC_URL not set")
	}
	if !common.IsHexAddress(cfg.ERC20TokenAddr) {
		return nil, fmt.Errorf("HUB_ERC20_TOKEN_ADDR invalid: %q", cfg.ERC20TokenAddr)
	}
	ec := &ERC20Client{
		rpcURL:    cfg.ChainRPCURL,
		tokenAddr: common.HexToAddress(cfg.ERC20TokenAddr),
		chainID:   big.NewInt(cfg.ChainID),
	}
	if cfg.ProviderPrivateKey != "" {
		sk, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.ProviderPrivateKey, "0x"))
		if err != nil {
			return nil, fmt.Errorf("provider private key: %w", err)
		}
		ec.sk = sk
	}
	return ec, nil
}

func (e *ERC20Client) newToken(ctx context.Context) (*token.Token, *ethclient.Client, error) {
	client, err := ethclient.DialContext(ctx, e.rpcURL)
	if err != nil {
		return nil, nil, err
	}
	tk, err := token.NewToken(e.tokenAddr, client)
	if err != nil {
		client.Close()
		return nil, nil, err
	}
	return tk, client, nil
}

// BalanceOf returns the owner's token balance in base units.
func (e *ERC20Client) BalanceOf(ctx context.Context, owner common.Address) (*big.Int, error) {
	tk, client, err := e.newToken(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return tk.BalanceOf(&bind.CallOpts{Context: ctx}, owner)
}

// Allowance returns how much spender may pull from owner.
func (e *ERC20Client) Allowance(ctx context.Context, owner, spender common.Address) (*big.Int, error) {
	tk, client, err := e.newToken(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return tk.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
}

// TransferFrom pulls amount from `from` to `to` using the provider key. It
// waits for the receipt and fails on revert: settlement clears debt based on
// this call succeeding, so a submitted-but-reverted transfer (insufficient
// balance/allowance) must be reported as an error, not a success.
func (e *ERC20Client) TransferFrom(ctx context.Context, from, to common.Address, amount *big.Int) (common.Hash, error) {
	if e.sk == nil {
		return common.Hash{}, fmt.Errorf("provider private key not configured")
	}
	tk, client, err := e.newToken(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	defer client.Close()

	auth, err := bind.NewKeyedTransactorWithChainID(e.sk, e.chainID)
	if err != nil {
		return common.Hash{}, err
	}
	auth.Context = ctx

	tx, err := tk.TransferFrom(auth, from, to, amount)
	if err != nil {
		return common.Hash{}, err
	}
	rcpt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return tx.Hash(), err
	}
	if rcpt.Status != 1 {
		return tx.Hash(), fmt.Errorf("transferFrom reverted (tx %s)", tx.Hash().Hex())
	}
	return tx.Hash(), nil
}

// providerFromKey derives the provider address from the configured key.
func (e *ERC20Client) providerFromKey() (common.Address, bool) {
	if e.sk == nil {
		return common.Address{}, false
	}
	return crypto.PubkeyToAddress(e.sk.PublicKey), true
}
