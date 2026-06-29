package main

import (
	"fmt"
	"log"
	"math/big"

	contract "github.com/unibaseio/da-sdk-go/contract/common"
	daogovernor "github.com/unibaseio/da-sdk-go/contract/v2/go/dao/governor"
	daotimelock "github.com/unibaseio/da-sdk-go/contract/v2/go/dao/timelock"
	daotoken "github.com/unibaseio/da-sdk-go/contract/v2/go/dao/token"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/epoch"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/eproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/everify"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/node"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/piece"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/proxy"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/rsproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/vub"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployGovernanceTokenImpl(client *ethclient.Client, sk string, name, symbol string, supply *big.Int, owner common.Address) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}
	tAddr, tx, ti, err := daotoken.DeployGovernanceToken(au, client)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploy GovernanceToken: %w", err)
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return common.Address{}, err
	}
	log.Println("GovernanceToken deployed at:", tAddr.Hex())

	au, err = makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}
	tx, err = ti.Initialize(au, name, symbol, supply, owner)
	if err != nil {
		return common.Address{}, fmt.Errorf("initialize GovernanceToken: %w", err)
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return common.Address{}, err
	}
	log.Println("GovernanceToken initialized")
	SaveDeployment("GovernanceToken", tAddr)
	return tAddr, nil
}

// DeployVUBProxy deploys the vote-escrowed UB staking module (VUB) behind a
// UUPS proxy. VUB implements IVotes and replaces the legacy placeholder
// GovernanceToken as the token fed to DAOGovernor: stake UB -> ve-weighted vUB
// voting power.
// ubAddr is the staked token (UB on this chain); rewardAddr is the
// community-incentive token paying staking APY (may equal ubAddr on testnet).
func DeployVUBProxy(client *ethclient.Client, sk string, ubAddr, rewardAddr, owner common.Address) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}
	implAddr, tx, _, err := vub.DeployVUB(au, client)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploy VUB impl: %w", err)
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return common.Address{}, err
	}
	log.Println("VUB impl deployed at:", implAddr.Hex())

	vubABI, err := vub.VUBMetaData.GetAbi()
	if err != nil {
		return common.Address{}, err
	}
	initData, err := vubABI.Pack("initialize", ubAddr, rewardAddr, owner)
	if err != nil {
		return common.Address{}, err
	}

	au, err = makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}
	proxyAddr, tx, _, err := proxy.DeployERC1967Proxy(au, client, implAddr, initData)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploy VUB proxy: %w", err)
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return common.Address{}, err
	}
	log.Printf("VUBProxy deployed at: %s (ub=%s reward=%s)\n", proxyAddr.Hex(), ubAddr.Hex(), rewardAddr.Hex())
	SaveDeployment("VUB", proxyAddr)
	return proxyAddr, nil
}

func DeployDAOTimelockImpl(client *ethclient.Client, sk string, minDelay *big.Int, admin common.Address) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}
	tAddr, tx, ti, err := daotimelock.DeployDAOTimelock(au, client)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploy DAOTimelock: %w", err)
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return common.Address{}, err
	}
	log.Println("DAOTimelock deployed at:", tAddr.Hex())

	// proposers/executors are set to address(0) as placeholder;
	// PROPOSER_ROLE is granted to the governor after it's deployed.
	placeholder := []common.Address{{}}
	au, err = makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}
	tx, err = ti.Initialize(au, minDelay, placeholder, placeholder, admin)
	if err != nil {
		return common.Address{}, fmt.Errorf("initialize DAOTimelock: %w", err)
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return common.Address{}, err
	}
	log.Printf("DAOTimelock initialized (minDelay=%s)\n", minDelay)
	SaveDeployment("DAOTimelock", tAddr)
	return tAddr, nil
}

func DeployDAOGovernorImpl(client *ethclient.Client, sk string, govTokenAddr, timelockAddr common.Address, votingDelay *big.Int, votingPeriod uint32, proposalThreshold, quorumFraction *big.Int) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}
	gAddr, tx, gi, err := daogovernor.DeployDAOGovernor(au, client)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploy DAOGovernor: %w", err)
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return common.Address{}, err
	}
	log.Println("DAOGovernor deployed at:", gAddr.Hex())

	au, err = makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}
	tx, err = gi.Initialize(au, govTokenAddr, timelockAddr, votingDelay, votingPeriod, proposalThreshold, quorumFraction)
	if err != nil {
		return common.Address{}, fmt.Errorf("initialize DAOGovernor: %w", err)
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return common.Address{}, err
	}
	log.Println("DAOGovernor initialized")
	SaveDeployment("DAOGovernor", gAddr)
	return gAddr, nil
}

// SetupTimelockRoles grants PROPOSER_ROLE to the governor and
// EXECUTOR_ROLE to address(0) (anyone can execute) on the timelock.
func SetupTimelockRoles(client *ethclient.Client, sk string, timelockAddr, governorAddr common.Address) error {
	ti, err := daotimelock.NewDAOTimelock(timelockAddr, client)
	if err != nil {
		return fmt.Errorf("connect DAOTimelock: %w", err)
	}
	proposerRole, err := ti.PROPOSERROLE(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("get PROPOSER_ROLE: %w", err)
	}
	executorRole, err := ti.EXECUTORROLE(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("get EXECUTOR_ROLE: %w", err)
	}

	au, err := makeAuth(sk)
	if err != nil {
		return err
	}
	tx, err := ti.GrantRole(au, proposerRole, governorAddr)
	if err != nil {
		return fmt.Errorf("grant PROPOSER_ROLE: %w", err)
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return err
	}
	log.Printf("Granted PROPOSER_ROLE to Governor: %s\n", governorAddr.Hex())

	au, err = makeAuth(sk)
	if err != nil {
		return err
	}
	tx, err = ti.GrantRole(au, executorRole, common.Address{})
	if err != nil {
		return fmt.Errorf("grant EXECUTOR_ROLE: %w", err)
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return err
	}
	log.Println("Granted EXECUTOR_ROLE to anyone (address(0))")
	return nil
}

// GrantGovernorRoleToContracts grants GOVERNOR_ROLE on each core contract
// proxy to the DAOTimelock, enabling DAO-controlled parameter updates.
func GrantGovernorRoleToContracts(client *ethclient.Client, sk string, timelockAddr, epochProxy, nodeProxy, pieceProxy, rsproofProxy, everifyProxy, eproofProxy common.Address) error {
	epochInst, err := epoch.NewEpoch(epochProxy, client)
	if err != nil {
		return fmt.Errorf("connect Epoch: %w", err)
	}
	governorRole, err := epochInst.GOVERNORROLE(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("read GOVERNOR_ROLE: %w", err)
	}

	type grantCall struct {
		name string
		fn   func() (common.Hash, error)
	}

	nodeInst, err := node.NewNode(nodeProxy, client)
	if err != nil {
		return fmt.Errorf("connect Node: %w", err)
	}
	pieceInst, err := piece.NewPiece(pieceProxy, client)
	if err != nil {
		return fmt.Errorf("connect Piece: %w", err)
	}
	rsproofInst, err := rsproof.NewRSProof(rsproofProxy, client)
	if err != nil {
		return fmt.Errorf("connect RSProof: %w", err)
	}
	everifyInst, err := everify.NewEVerify(everifyProxy, client)
	if err != nil {
		return fmt.Errorf("connect EVerify: %w", err)
	}
	eproofInst, err := eproof.NewEProof(eproofProxy, client)
	if err != nil {
		return fmt.Errorf("connect EProof: %w", err)
	}

	grants := []grantCall{
		{"Epoch", func() (common.Hash, error) {
			au, e := makeAuth(sk)
			if e != nil {
				return common.Hash{}, e
			}
			tx, e := epochInst.GrantRole(au, governorRole, timelockAddr)
			if e != nil {
				return common.Hash{}, e
			}
			return tx.Hash(), nil
		}},
		{"Node", func() (common.Hash, error) {
			au, e := makeAuth(sk)
			if e != nil {
				return common.Hash{}, e
			}
			tx, e := nodeInst.GrantRole(au, governorRole, timelockAddr)
			if e != nil {
				return common.Hash{}, e
			}
			return tx.Hash(), nil
		}},
		{"Piece", func() (common.Hash, error) {
			au, e := makeAuth(sk)
			if e != nil {
				return common.Hash{}, e
			}
			tx, e := pieceInst.GrantRole(au, governorRole, timelockAddr)
			if e != nil {
				return common.Hash{}, e
			}
			return tx.Hash(), nil
		}},
		{"RSProof", func() (common.Hash, error) {
			au, e := makeAuth(sk)
			if e != nil {
				return common.Hash{}, e
			}
			tx, e := rsproofInst.GrantRole(au, governorRole, timelockAddr)
			if e != nil {
				return common.Hash{}, e
			}
			return tx.Hash(), nil
		}},
		{"EVerify", func() (common.Hash, error) {
			au, e := makeAuth(sk)
			if e != nil {
				return common.Hash{}, e
			}
			tx, e := everifyInst.GrantRole(au, governorRole, timelockAddr)
			if e != nil {
				return common.Hash{}, e
			}
			return tx.Hash(), nil
		}},
		{"EProof", func() (common.Hash, error) {
			au, e := makeAuth(sk)
			if e != nil {
				return common.Hash{}, e
			}
			tx, e := eproofInst.GrantRole(au, governorRole, timelockAddr)
			if e != nil {
				return common.Hash{}, e
			}
			return tx.Hash(), nil
		}},
	}

	for _, g := range grants {
		hash, err := g.fn()
		if err != nil {
			return fmt.Errorf("grant GOVERNOR_ROLE on %s: %w", g.name, err)
		}
		if err = contract.CheckTx(ChainURL, hash); err != nil {
			return fmt.Errorf("tx failed for %s: %w", g.name, err)
		}
		log.Printf("Granted GOVERNOR_ROLE to Timelock on %s\n", g.name)
	}
	return nil
}

func deployDAO(client *ethclient.Client, sk string, owner, ubAddr, rewardAddr, epochProxy, nodeProxy, pieceProxy, rsproofProxy, everifyProxy, eproofProxy common.Address) {
	// Step 1: governance token. Default = vUB (stake UB -> ve-weighted votes);
	// the legacy GovernanceToken remains available as a fallback via -dao-gov-token legacy.
	var govTokenAddr common.Address
	var err error
	if daoGovTokenKind == "legacy" {
		govTokenAddr, err = DeployGovernanceTokenImpl(client, sk, daoTokenName, daoTokenSymbol, daoTokenSupply, owner)
		if err != nil {
			log.Println("Failed to deploy GovernanceToken:", err)
			return
		}
	} else {
		if rewardAddr == (common.Address{}) {
			rewardAddr = ubAddr // testnet default: reward in the same token
		}
		govTokenAddr, err = DeployVUBProxy(client, sk, ubAddr, rewardAddr, owner)
		if err != nil {
			log.Println("Failed to deploy VUB:", err)
			return
		}
	}

	// Step 2: DAOTimelock
	timelockAddr, err := DeployDAOTimelockImpl(client, sk, daoTimelockDelay, owner)
	if err != nil {
		log.Println("Failed to deploy DAOTimelock:", err)
		return
	}

	// Step 3: DAOGovernor
	governorAddr, err := DeployDAOGovernorImpl(client, sk, govTokenAddr, timelockAddr, daoVotingDelay, daoVotingPeriod, daoProposalThreshold, daoQuorumFraction)
	if err != nil {
		log.Println("Failed to deploy DAOGovernor:", err)
		return
	}

	// Step 4: Setup timelock roles (PROPOSER → governor, EXECUTOR → anyone)
	if err = SetupTimelockRoles(client, sk, timelockAddr, governorAddr); err != nil {
		log.Println("Failed to setup timelock roles:", err)
		return
	}

	// Step 5: Grant GOVERNOR_ROLE to timelock on all core contract proxies
	if err = GrantGovernorRoleToContracts(client, sk, timelockAddr, epochProxy, nodeProxy, pieceProxy, rsproofProxy, everifyProxy, eproofProxy); err != nil {
		log.Println("Failed to grant GOVERNOR_ROLE to Timelock:", err)
		return
	}

	log.Println("=== DAO Governance Deployment Complete ===")
	log.Printf("  GovToken (%s): %s\n", daoGovTokenKind, govTokenAddr.Hex())
	log.Printf("  DAOTimelock:     %s\n", timelockAddr.Hex())
	log.Printf("  DAOGovernor:     %s\n", governorAddr.Hex())
}
