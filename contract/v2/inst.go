package contract

import (
	"context"

	"github.com/unibaseio/da-sdk-go/contract/v2/go/epoch"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/eproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/everify"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/node"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/piece"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/plonk/add"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/plonk/kzg"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/plonk/mul"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/plonk/rsone"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/rsproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/token"

	"github.com/ethereum/go-ethereum/ethclient"
)

func (c *ContractManage) NewEpoch(ctx context.Context) (*epoch.Epoch, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return epoch.NewEpoch(c.EpochAddr, client)
}

func (c *ContractManage) NewNode(ctx context.Context) (*node.Node, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return node.NewNode(c.NodeAddr, client)
}

func (c *ContractManage) NewPiece(ctx context.Context) (*piece.Piece, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return piece.NewPiece(c.PieceAddr, client)
}

func (c *ContractManage) NewRSProof(ctx context.Context) (*rsproof.RSProof, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return rsproof.NewRSProof(c.RSProofAddr, client)
}

func (c *ContractManage) NewRSOne(ctx context.Context) (*rsone.PlonkVerifier, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return rsone.NewPlonkVerifier(c.RSOneAddr, client)
}

func (c *ContractManage) NewKZGPlonk(ctx context.Context) (*kzg.PlonkVerifier, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return kzg.NewPlonkVerifier(c.KZGAddr, client)
}

func (c *ContractManage) NewMulPlonk(ctx context.Context) (*mul.PlonkVerifier, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return mul.NewPlonkVerifier(c.MulAddr, client)
}

func (c *ContractManage) NewAddPlonk(ctx context.Context) (*add.PlonkVerifier, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return add.NewPlonkVerifier(c.AddAddr, client)
}

func (c *ContractManage) NewEProof(ctx context.Context) (*eproof.EProof, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return eproof.NewEProof(c.EProofAddr, client)
}

func (c *ContractManage) NewEVerify(ctx context.Context) (*everify.EVerify, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return everify.NewEVerify(c.EVerifyAddr, client)
}

func (c *ContractManage) NewToken(ctx context.Context) (*token.Token, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return token.NewToken(c.TokenAddr, client)
}
