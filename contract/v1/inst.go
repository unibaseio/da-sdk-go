package contract

import (
	"context"

	com "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/bank"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/control"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/epoch"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/eproof"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/everify"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/gpu"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/model"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/node"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/piece"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/plonk/kzg"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/plonk/rsone"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/reward"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/rsproof"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/space"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (c *ContractManage) NewBank(ctx context.Context) (*bank.Bank, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return bank.NewBank(c.BankAddr, client)
}

func (c *ContractManage) NewToken(ctx context.Context) (*token.Token, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	taddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "token")
	if err != nil {
		return nil, err
	}
	return token.NewToken(taddr, client)
}

func (c *ContractManage) NewEpoch(ctx context.Context) (*epoch.Epoch, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	eaddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "epoch")
	if err != nil {
		return nil, err
	}
	return epoch.NewEpoch(eaddr, client)
}

func (c *ContractManage) NewControl(ctx context.Context) (*control.Control, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	caddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "control")
	if err != nil {
		return nil, err
	}
	return control.NewControl(caddr, client)
}

func (c *ContractManage) NewNode(ctx context.Context) (*node.Node, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	naddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "node")
	if err != nil {
		return nil, err
	}
	return node.NewNode(naddr, client)
}

func (c *ContractManage) NewPiece(ctx context.Context) (*piece.Piece, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	faddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "piece")
	if err != nil {
		return nil, err
	}
	return piece.NewPiece(faddr, client)
}

func (c *ContractManage) NewReward(ctx context.Context) (*reward.Reward, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	faddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "reward")
	if err != nil {
		return nil, err
	}
	return reward.NewReward(faddr, client)
}

func (c *ContractManage) NewRSProof(ctx context.Context) (*rsproof.RSProof, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	faddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "rsproof")
	if err != nil {
		return nil, err
	}
	return rsproof.NewRSProof(faddr, client)
}

func (c *ContractManage) NewRSOne(ctx context.Context) (*rsone.PlonkVerifier, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	faddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "rsone")
	if err != nil {
		return nil, err
	}
	return rsone.NewPlonkVerifier(faddr, client)
}

func (c *ContractManage) NewKZGPlonk(ctx context.Context) (*kzg.PlonkVerifier, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	faddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "kzg")
	if err != nil {
		return nil, err
	}
	return kzg.NewPlonkVerifier(faddr, client)
}

func (c *ContractManage) NewMulPlonk(ctx context.Context) (*kzg.PlonkVerifier, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	faddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "mul")
	if err != nil {
		return nil, err
	}
	return kzg.NewPlonkVerifier(faddr, client)
}

func (c *ContractManage) NewAddPlonk(ctx context.Context) (*kzg.PlonkVerifier, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	faddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "add")
	if err != nil {
		return nil, err
	}
	return kzg.NewPlonkVerifier(faddr, client)
}

func (c *ContractManage) NewEProof(ctx context.Context) (*eproof.EProof, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	faddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "eproof")
	if err != nil {
		return nil, err
	}
	return eproof.NewEProof(faddr, client)
}

func (c *ContractManage) NewEVerify(ctx context.Context) (*everify.EVerify, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	faddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "everify")
	if err != nil {
		return nil, err
	}
	return everify.NewEVerify(faddr, client)
}

func (c *ContractManage) NewGPU(ctx context.Context) (*gpu.GPU, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	gaddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "gpu")
	if err != nil {
		return nil, err
	}
	return gpu.NewGPU(gaddr, client)
}

func (c *ContractManage) NewModel(ctx context.Context) (*model.Model, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	maddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "model")
	if err != nil {
		return nil, err
	}
	return model.NewModel(maddr, client)
}

func (c *ContractManage) NewSpace(ctx context.Context) (*space.Space, error) {
	client, err := ethclient.DialContext(ctx, c.RPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bi, err := bank.NewBank(c.BankAddr, client)
	if err != nil {
		return nil, err
	}
	aaddr, err := bi.Get(&bind.CallOpts{From: com.Base}, "space")
	if err != nil {
		return nil, err
	}
	return space.NewSpace(aaddr, client)
}
