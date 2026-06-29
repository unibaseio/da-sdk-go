package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	contract "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/epoch"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/eproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/everify"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/node"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/piece"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/plonk/add"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/plonk/kzg"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/plonk/mul"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/plonk/rsone"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/proxy"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/rsproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/validatorreward"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// deploy-local monotonic nonce. The deploy tool fires txs back-to-back;
// contract.MakeAuth leaves auth.Nonce nil, so go-ethereum fetches PendingNonceAt
// per send — and under RPC read-after-write lag (Alchemy) two fast txs can reuse
// a nonce, surfacing as "already known" / "replacement underpriced". makeAuth
// assigns max(chain pending nonce, our last+1) so a lagging read never reuses one.
var (
	deployNonceNext uint64
	deployNonceInit bool
)

func makeAuth(sk string) (*bind.TransactOpts, error) {
	auth, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return nil, err
	}
	// Initialize the local counter ONCE from the chain's pending nonce, then
	// increment purely locally. Re-reading PendingNonceAt per call is unreliable
	// on load-balanced RPCs (Alchemy): a transient ghost/queued tx can bump the
	// reported nonce, and taking it would skip a nonce, leaving an unminable gap
	// (the tx sits queued forever behind the missing nonce). The deploy tool is
	// the sole, sequential sender for this key, so a local counter is gap-free.
	if !deployNonceInit {
		cl, err := ethclient.Dial(ChainURL)
		if err != nil {
			return nil, err
		}
		defer cl.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		chainNonce, err := cl.PendingNonceAt(ctx, auth.From)
		if err != nil {
			return nil, err
		}
		deployNonceNext = chainNonce
		deployNonceInit = true
	}
	auth.Nonce = new(big.Int).SetUint64(deployNonceNext)
	deployNonceNext++
	return auth, nil
}

// DeploymentRecord represents a single deployment entry
type DeploymentRecord struct {
	Timestamp   string            `json:"timestamp"`
	ChainURL    string            `json:"chain_url"`
	ChainID     int64             `json:"chain_id"`
	Deployments map[string]string `json:"deployments"`
}

// DeploymentHistory represents the complete deployment history
type DeploymentHistory struct {
	Records []DeploymentRecord `json:"records"`
}

// SaveDeployment saves a deployment record to the out file
func SaveDeployment(contractName string, contractAddr common.Address) error {
	const outFile = "out"

	// Read existing history
	history := DeploymentHistory{}
	data, err := os.ReadFile(outFile)
	if err == nil {
		// File exists, unmarshal existing data
		if err := json.Unmarshal(data, &history); err != nil {
			log.Printf("Warning: failed to parse existing out file: %v\n", err)
			history = DeploymentHistory{}
		}
	}

	// Get current timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Find or create today's record
	var currentRecord *DeploymentRecord
	for i := range history.Records {
		if history.Records[i].ChainURL == ChainURL {
			// Found existing record for this chain
			currentRecord = &history.Records[i]
			break
		}
	}

	if currentRecord == nil {
		// Create new record
		history.Records = append(history.Records, DeploymentRecord{
			Timestamp:   timestamp,
			ChainURL:    ChainURL,
			ChainID:     ChainID,
			Deployments: make(map[string]string),
		})
		currentRecord = &history.Records[len(history.Records)-1]
	}

	// Update deployment
	currentRecord.Timestamp = timestamp
	currentRecord.Deployments[contractName] = contractAddr.Hex()

	// Write back to file
	jsonData, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal deployment history: %w", err)
	}

	if err := os.WriteFile(outFile, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write out file: %w", err)
	}

	log.Printf("Saved deployment: %s -> %s\n", contractName, contractAddr.Hex())
	return nil
}

func DeployEpochImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	tAddr, tx, _, err := epoch.DeployEpoch(au, client)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Println("epoch: ", tAddr.Hex())
	SaveDeployment("Epoch", tAddr)
	return tAddr, nil
}

func DeployNodeImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	tAddr, tx, _, err := node.DeployNode(au, client)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	log.Println("node: ", tAddr.Hex())
	SaveDeployment("Node", tAddr)
	return tAddr, nil
}

func DeployPieceImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	tAddr, tx, _, err := piece.DeployPiece(au, client)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	log.Println("piece: ", tAddr.Hex())
	SaveDeployment("Piece", tAddr)
	return tAddr, nil
}

func DeployRSProofImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	tAddr, tx, _, err := rsproof.DeployRSProof(au, client)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	log.Println("rsproof: ", tAddr.Hex())
	SaveDeployment("RSProof", tAddr)
	return tAddr, nil
}

// waitForCode polls until the freshly-deployed contract's code is visible at
// addr. Guards against RPC read-after-write lag (e.g. Alchemy's load-balanced
// replicas returning empty code in the same second a proxy is created), which
// otherwise makes the next estimateGas fail with "no contract code at given
// address". CheckTx confirms the tx is mined; this confirms the read side caught up.
func waitForCode(client *ethclient.Client, addr common.Address) error {
	for i := 0; i < 30; i++ {
		code, err := client.CodeAt(context.Background(), addr, nil)
		if err == nil && len(code) > 0 {
			return nil
		}
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("timed out waiting for code at %s", addr.Hex())
}

func SetRSVKRootV2(client *ethclient.Client, sk string, n, k int, rsproofAddr common.Address) error {
	au, err := makeAuth(sk)
	if err != nil {
		return err
	}

	ni, err := rsproof.NewRSProof(rsproofAddr, client)
	if err != nil {
		return err
	}

	vt := new(big.Int)
	switch {
	case n == 6 && k == 4:
		vt.SetString(contract.RSn6k4VKRoot, 10)
	case n == 14 && k == 7:
		vt.SetString(contract.RSn14k7VKRoot, 10)
	case n == 32 && k == 16:
		vt.SetString(contract.RSn32k16VKRoot, 10)
	case n == 64 && k == 32:
		vt.SetString(contract.RSn64k32VKRoot, 10)
	default:
		return fmt.Errorf("unsupported rs policy: %d %d", n, k)
	}

	tx, err := ni.SetVKRoot(au, uint8(n), uint8(k), vt)
	if err != nil {
		return err
	}
	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Printf("set vk root: %d %d %d\n", n, k, vt)
	return nil
}

func DeployEproofImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	tAddr, tx, _, err := eproof.DeployEProof(au, client)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	log.Println("eproof: ", tAddr.Hex())
	SaveDeployment("EProof", tAddr)
	return tAddr, nil
}

func DeployEverifyImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	tAddr, tx, _, err := everify.DeployEVerify(au, client)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	log.Println("everify: ", tAddr.Hex())
	SaveDeployment("EVerify", tAddr)
	return tAddr, nil
}

func DeployRSPlonkImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	tAddr, tx, _, err := rsone.DeployPlonkVerifier(au, client)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Println("rsone: ", tAddr.Hex())
	SaveDeployment("RSOneVerifier", tAddr)
	return tAddr, nil
}

func DeployKZGPlonkImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	tAddr, tx, _, err := kzg.DeployPlonkVerifier(au, client)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Println("kzg: ", tAddr.Hex())
	SaveDeployment("KzgVerifier", tAddr)
	return tAddr, nil
}

func DeployMulPlonkImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	tAddr, tx, _, err := mul.DeployPlonkVerifier(au, client)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Println("mul: ", tAddr.Hex())
	SaveDeployment("MulVerifier", tAddr)
	return tAddr, nil
}

func DeployAddPlonkImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	tAddr, tx, _, err := add.DeployPlonkVerifier(au, client)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Println("add: ", tAddr.Hex())
	SaveDeployment("AddVerifier", tAddr)
	return tAddr, nil
}

// DeployEpochProxy deploys the ERC1967 proxy for Epoch with initialization
func DeployEpochProxy(client *ethclient.Client, sk string, implAddr common.Address, slots uint64, owner common.Address) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	// Get the Epoch ABI to encode the initialize call
	epochABI, err := epoch.EpochMetaData.GetAbi()
	if err != nil {
		return common.Address{}, err
	}

	// Encode the initialize call
	initData, err := epochABI.Pack("initialize", slots, owner)
	if err != nil {
		return common.Address{}, err
	}

	// Deploy the proxy
	proxyAddr, tx, _, err := proxy.DeployERC1967Proxy(au, client, implAddr, initData)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Printf("EpochProxy deployed at: %s\n", proxyAddr.Hex())
	SaveDeployment("EpochProxy", proxyAddr)
	return proxyAddr, nil
}

// DeployNodeProxy deploys the ERC1967 proxy for Node with initialization
func DeployNodeProxy(client *ethclient.Client, sk string, implAddr common.Address, token common.Address, epochProxy common.Address, owner common.Address) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	nodeABI, err := node.NodeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, err
	}

	initData, err := nodeABI.Pack("initialize", token, epochProxy, owner)
	if err != nil {
		return common.Address{}, err
	}

	proxyAddr, tx, _, err := proxy.DeployERC1967Proxy(au, client, implAddr, initData)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Printf("NodeProxy deployed at: %s\n", proxyAddr.Hex())
	SaveDeployment("NodeProxy", proxyAddr)
	return proxyAddr, nil
}

// DeployPieceProxy deploys the ERC1967 proxy for Piece with initialization
func DeployPieceProxy(client *ethclient.Client, sk string, implAddr common.Address, token common.Address, epochProxy common.Address, nodeProxy common.Address, initParams piece.IPieceInitParams, owner common.Address) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	pieceABI, err := piece.PieceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, err
	}

	initData, err := pieceABI.Pack("initialize", token, epochProxy, nodeProxy, initParams, owner)
	if err != nil {
		return common.Address{}, err
	}

	proxyAddr, tx, _, err := proxy.DeployERC1967Proxy(au, client, implAddr, initData)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Printf("PieceProxy deployed at: %s\n", proxyAddr.Hex())
	SaveDeployment("PieceProxy", proxyAddr)
	return proxyAddr, nil
}

// DeployRSProofProxy deploys the ERC1967 proxy for RSProof with initialization
func DeployRSProofProxy(client *ethclient.Client, sk string, implAddr common.Address, initParams rsproof.IRSProofInitParams, owner common.Address) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	rsproofABI, err := rsproof.RSProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, err
	}

	initData, err := rsproofABI.Pack("initialize", initParams, owner)
	if err != nil {
		return common.Address{}, err
	}

	proxyAddr, tx, _, err := proxy.DeployERC1967Proxy(au, client, implAddr, initData)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Printf("RSProofProxy deployed at: %s\n", proxyAddr.Hex())
	SaveDeployment("RSProofProxy", proxyAddr)
	return proxyAddr, nil
}

// DeployEVerifyProxy deploys the ERC1967 proxy for EVerify with initialization
func DeployEVerifyProxy(client *ethclient.Client, sk string, implAddr common.Address, epochProxy common.Address, pieceProxy common.Address, kzgVerifier common.Address, mulVerifier common.Address, addVerifier common.Address, owner common.Address) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	everifyABI, err := everify.EVerifyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, err
	}

	initData, err := everifyABI.Pack("initialize", epochProxy, pieceProxy, kzgVerifier, mulVerifier, addVerifier, owner)
	if err != nil {
		return common.Address{}, err
	}

	proxyAddr, tx, _, err := proxy.DeployERC1967Proxy(au, client, implAddr, initData)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Printf("EVerifyProxy deployed at: %s\n", proxyAddr.Hex())
	SaveDeployment("EVerifyProxy", proxyAddr)
	return proxyAddr, nil
}

// DeployEProofProxy deploys the ERC1967 proxy for EProof with initialization
func DeployEProofProxy(client *ethclient.Client, sk string, implAddr common.Address, initParams eproof.IEProofInitParams, owner common.Address) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}

	eproofABI, err := eproof.EProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, err
	}

	initData, err := eproofABI.Pack("initialize", initParams, owner)
	if err != nil {
		return common.Address{}, err
	}

	proxyAddr, tx, _, err := proxy.DeployERC1967Proxy(au, client, implAddr, initData)
	if err != nil {
		return common.Address{}, err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}

	log.Printf("EProofProxy deployed at: %s\n", proxyAddr.Hex())
	SaveDeployment("EProofProxy", proxyAddr)
	return proxyAddr, nil
}

// SetEProofAddress sets the EProof address in EVerify contract
func SetEProofAddress(client *ethclient.Client, sk string, everifyProxy common.Address, eproofProxy common.Address) error {
	au, err := makeAuth(sk)
	if err != nil {
		return err
	}

	everifyInstance, err := everify.NewEVerify(everifyProxy, client)
	if err != nil {
		return err
	}

	tx, err := everifyInstance.SetEProof(au, eproofProxy)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}

	log.Printf("EVerify at %s set EProof to %s\n", everifyProxy.Hex(), eproofProxy.Hex())
	return nil
}

// SetNodeAddresses sets EProof and RSProof addresses in Node contract
func SetNodeAddresses(client *ethclient.Client, sk string, nodeProxy common.Address, eproofProxy common.Address, rsproofProxy common.Address) error {
	au, err := makeAuth(sk)
	if err != nil {
		return err
	}

	nodeInstance, err := node.NewNode(nodeProxy, client)
	if err != nil {
		return err
	}

	tx, err := nodeInstance.SetAddress(au, eproofProxy, rsproofProxy)
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}

	log.Printf("Node at %s set EProof to %s and RSProof to %s\n", nodeProxy.Hex(), eproofProxy.Hex(), rsproofProxy.Hex())
	return nil
}

func SetMinPledgeV2(client *ethclient.Client, sk string, _typ uint8, val *big.Int, nodeAddr common.Address) error {
	au, err := makeAuth(sk)
	if err != nil {
		return err
	}

	ni, err := node.NewNode(nodeAddr, client)
	if err != nil {
		return err
	}

	tx, err := ni.Set(au, _typ, val)
	if err != nil {
		return err
	}
	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Printf("set nodeType:%d minpledge:%s success\n", _typ, val.String())
	return nil
}

// SetBasePenaltyV2 sets the fraud-proof penalty on both proof contracts.
// The proxies initialize basePenalty to 1e18; the dev-confirmed value is
// contract.DefaultPenalty (10000 UB), set post-deploy via the GOVERNOR_ROLE.
func SetBasePenaltyV2(client *ethclient.Client, sk string, rsproofAddr, eproofAddr common.Address, penalty *big.Int) error {
	au, err := makeAuth(sk)
	if err != nil {
		return err
	}

	ri, err := rsproof.NewRSProof(rsproofAddr, client)
	if err != nil {
		return err
	}
	tx, err := ri.SetBasePenalty(au, penalty)
	if err != nil {
		return err
	}
	if err := contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return err
	}
	log.Printf("rsproof basePenalty set to %s\n", penalty.String())

	// fresh auth for the second tx: makeAuth assigns an explicit, incrementing
	// nonce, so reusing the same `au` would send eproof's tx at rsproof's nonce
	// ("replacement underpriced"). One makeAuth per tx.
	au, err = makeAuth(sk)
	if err != nil {
		return err
	}
	ei, err := eproof.NewEProof(eproofAddr, client)
	if err != nil {
		return err
	}
	tx, err = ei.SetBasePenalty(au, penalty)
	if err != nil {
		return err
	}
	if err := contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return err
	}
	log.Printf("eproof basePenalty set to %s\n", penalty.String())
	return nil
}

// --- ValidatorReward (FixB+A2) ---------------------------------------------

func DeployValidatorRewardImpl(client *ethclient.Client, sk string) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}
	tAddr, tx, _, err := validatorreward.DeployValidatorReward(au, client)
	if err != nil {
		return common.Address{}, err
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return common.Address{}, err
	}
	log.Println("validatorReward: ", tAddr.Hex())
	SaveDeployment("ValidatorReward", tAddr)
	return tAddr, nil
}

func DeployValidatorRewardProxy(client *ethclient.Client, sk string, implAddr, token, owner common.Address) (common.Address, error) {
	au, err := makeAuth(sk)
	if err != nil {
		return common.Address{}, err
	}
	vrABI, err := validatorreward.ValidatorRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, err
	}
	initData, err := vrABI.Pack("initialize", token, owner)
	if err != nil {
		return common.Address{}, err
	}
	proxyAddr, tx, _, err := proxy.DeployERC1967Proxy(au, client, implAddr, initData)
	if err != nil {
		return common.Address{}, err
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return common.Address{}, err
	}
	log.Printf("ValidatorRewardProxy deployed at: %s\n", proxyAddr.Hex())
	SaveDeployment("ValidatorRewardProxy", proxyAddr)
	return proxyAddr, nil
}

// SetValidatorPool repoints RSProof + EProof's penalty-share recipient to the
// validator reward pool (FixB+A2). Until called, the protocol share defaults to
// `base` (backward-compatible).
func SetValidatorPool(client *ethclient.Client, sk string, rsproofProxy, eproofProxy, pool common.Address) error {
	au, err := makeAuth(sk)
	if err != nil {
		return err
	}
	ri, err := rsproof.NewRSProof(rsproofProxy, client)
	if err != nil {
		return err
	}
	tx, err := ri.SetValidatorPool(au, pool)
	if err != nil {
		return err
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return err
	}

	au, err = makeAuth(sk)
	if err != nil {
		return err
	}
	ei, err := eproof.NewEProof(eproofProxy, client)
	if err != nil {
		return err
	}
	tx, err = ei.SetValidatorPool(au, pool)
	if err != nil {
		return err
	}
	if err = contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return err
	}
	log.Printf("validatorPool set on RSProof+EProof: %s\n", pool.Hex())
	return nil
}
