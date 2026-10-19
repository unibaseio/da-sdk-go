package contract

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	etypes "github.com/ethereum/go-ethereum/core/types"
	com "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/node"
)

// NodeExitInfo is a node's standing on the Node contract at c.NodeAddr and
// what its exit allows.
type NodeExitInfo struct {
	Info          node.INodeNodeInfo
	Delay         uint64 // epochs between terminate and a full withdrawal
	Current       uint64 // the Node contract's epoch now
	Withdrawable  *big.Int
	WithdrawAfter uint64 // epoch from which the whole unlocked stake can go (0 while active)
}

// NodeExit reads this node's stake and exit state on c.NodeAddr. Retiring a
// node from an old deployment points c.NodeAddr at that deployment's Node.
func (c *ContractManage) NodeExit() (NodeExitInfo, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), readCallTimeout())
	defer cancel()
	ni, err := c.NewNode(ctx)
	if err != nil {
		return NodeExitInfo{}, err
	}
	opts := &bind.CallOpts{From: com.Base, Context: ctx}
	var r NodeExitInfo
	if r.Info, err = ni.NodeInfoOf(opts, c.From()); err != nil {
		return r, err
	}
	if r.Delay, err = ni.Delay(opts); err != nil {
		return r, err
	}
	ea, err := ni.Epoch(opts)
	if err != nil {
		return r, err
	}
	ep := c.EpochAddr
	c.EpochAddr = ea
	ei, err := c.NewEpoch(ctx)
	c.EpochAddr = ep
	if err != nil {
		return r, err
	}
	if r.Current, err = ei.Current(opts); err != nil {
		return r, err
	}
	r.Withdrawable = new(big.Int)
	if !r.Info.IsActive {
		r.WithdrawAfter = r.Info.ExitEpoch + r.Delay + 1
		if r.Current >= r.WithdrawAfter && r.Info.StakedAmount.Cmp(r.Info.LockedAmount) > 0 {
			r.Withdrawable.Sub(r.Info.StakedAmount, r.Info.LockedAmount)
		}
	}
	return r, nil
}

// TerminateNode ends this node's activity on c.NodeAddr; Node.withdraw then
// releases the stake once Node.delay epochs have passed.
func (c *ContractManage) TerminateNode() error {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancel()
	ni, err := c.NewNode(ctx)
	if err != nil {
		return err
	}
	return c.sendAndWait(ctx, func(au *bind.TransactOpts) (*etypes.Transaction, error) {
		return ni.Terminate(au)
	})
}

// WithdrawStake withdraws m of this node's stake from c.NodeAddr.
func (c *ContractManage) WithdrawStake(m *big.Int) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancel()
	ni, err := c.NewNode(ctx)
	if err != nil {
		return err
	}
	return c.sendAndWait(ctx, func(au *bind.TransactOpts) (*etypes.Transaction, error) {
		return ni.Withdraw(au, m)
	})
}
