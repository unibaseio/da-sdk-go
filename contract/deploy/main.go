package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"

	contract "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/token"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/eproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/piece"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/rsproof"
	dlog "github.com/unibaseio/da-sdk-go/lib/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ChainURL = contract.BNBTestnetChainRPC
	ChainID  = contract.BNBTestnetChainID

	blockInterval = uint64(3)
	bankAddr      = contract.BNBTestnetBankAddr
	tokenAddr     = contract.BNBTestnetTokenAddr
)

// V2 deployment configuration
var (
	slots    = uint64(16000)    // Epoch slots, 16000 * blockInterval(450ms) = 2h
	delay    = uint64(7)        // Piece delay in epochs
	minStore = uint64(1200)     // Minimum storage time, 1200 epochs = 100 days
	maxStore = uint64(12000)    // Maximum storage time
	maxSize  = uint64(33554432) // 32MB
	// minPrice     = big.NewInt(1e11)    // Minimum price per unit, per epoch*MB
	// streamPrice  = big.NewInt(1e12)    // Streaming price, per replica
	minPrice        = big.NewInt(1e8)  // just for test
	streamPrice     = big.NewInt(1e9)  // just for test
	minProveTime    = big.NewInt(8000) // Minimum prove time for RS/E proofs, 1 hour
	challengeWindow = uint64(7)        // Challenge window in epochs for EProof
	minPledgeMap    = map[uint8]*big.Int{
		// 1: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10)), // 10 tokens for type 1
		// 2: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10)), // 10 tokens for type 2
		// 3: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10)), // 10 tokens for type 3
		1: new(big.Int).Mul(big.NewInt(1e15), big.NewInt(10)), // test
		2: new(big.Int).Mul(big.NewInt(1e15), big.NewInt(10)), // test
		3: new(big.Int).Mul(big.NewInt(1e15), big.NewInt(10)), // test
	}

	baseAddr = common.HexToAddress("0xE0AD379735ba88B323298D091ff3b67Dd6C79852")
)

// DAO governance configuration
var (
	daoTokenName         = "Unibase DAO Governance Token"
	daoTokenSymbol       = "UNIDAOT"
	daoTokenSupply       = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1_000_000_000)) // 1B tokens
	daoVotingDelay       = big.NewInt(1)                                                  // 1 block
	daoVotingPeriod      = uint32(24_000)                                                 // ~3h at ~450ms block time on BSC
	daoProposalThreshold = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(2_500_000))      // 2.5M tokens (~0.25%)
	daoQuorumFraction    = big.NewInt(4)                                                  // 4%
	daoTimelockDelay     = big.NewInt(0)                                                  // no delay for testnet
)

func init() {
	dlog.SetLogLevel("DEBUG")
}

func main() {
	fmt.Println("connect to: ", ChainURL)
	sk := flag.String("sk", "", "private key for sending transaction")
	flag.Parse()

	client, err := ethclient.DialContext(context.TODO(), ChainURL)
	if err != nil {
		return
	}
	defer client.Close()
	DeployTokenTest(client, *sk)
	deployall_v2(client, *sk)
}

func DeployTokenTest(client *ethclient.Client, sk string) error {
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	tAddr, tx, ti, err := token.DeployToken(au, client, "Unibase", "UB")
	if err != nil {
		return err
	}

	err = contract.CheckTx(ChainURL, tx.Hash())
	if err != nil {
		return err
	}
	log.Println("token: ", tAddr.Hex())
	tokenAddr = tAddr

	owner, err := ti.Owner(&bind.CallOpts{From: contract.Base})
	if err != nil {
		return err
	}
	bal, err := ti.BalanceOf(&bind.CallOpts{From: contract.Base}, owner)
	if err != nil {
		return err
	}
	log.Println("owner has token: ", bal)
	return nil
}

func deployall_v1(client *ethclient.Client, sk string) {
	minPledge := big.NewInt(1e18)
	minPledge.Mul(minPledge, big.NewInt(10))
	err := DeployBank(client, sk)
	if err != nil {
		log.Println(err)
		return
	}
	err = DeployToken(client, sk)
	if err != nil {
		log.Println(err)
		return
	}
	err = SetMiner(client, sk, bankAddr)
	if err != nil {
		log.Println(err)
		return
	}
	err = DeployEpoch(client, sk)
	if err != nil {
		log.Println(err)
		return
	}

	err = DeployNode(client, sk)
	if err != nil {
		log.Println(err)
		return
	}
	err = SetMinPledge(client, sk, 1, minPledge)
	if err != nil {
		log.Println(err)
		return
	}
	err = SetMinPledge(client, sk, 2, minPledge)
	if err != nil {
		log.Println(err)
		return
	}
	err = SetMinPledge(client, sk, 3, minPledge)
	if err != nil {
		log.Println(err)
		return
	}
	err = DeployReward(client, sk)
	if err != nil {
		log.Println(err)
		return
	}
	err = DeployControl(client, sk)
	if err != nil {
		log.Println(err)
		return
	}
	err = DeployPiece(client, sk)
	if err != nil {
		log.Println(err)
		return
	}

	err = DeployRSPlonk(client, sk)
	if err != nil {
		log.Println(err)
		return
	}

	err = DeployRSProof(client, sk)
	if err != nil {
		log.Println(err)
		return
	}

	err = SetRSVKRoot(client, sk, 6, 4)
	if err != nil {
		log.Println(err)
		return
	}

	err = SetRSVKRoot(client, sk, 14, 7)
	if err != nil {
		log.Println(err)
		return
	}

	err = SetRSVKRoot(client, sk, 32, 16)
	if err != nil {
		log.Println(err)
		return
	}

	err = SetRSVKRoot(client, sk, 64, 32)
	if err != nil {
		log.Println(err)
		return
	}

	err = DeployEproof(client, sk)
	if err != nil {
		log.Println(err)
		return
	}

	err = DeployEverify(client, sk)
	if err != nil {
		log.Println(err)
		return
	}

	err = DeployKZGPlonk(client, sk)
	if err != nil {
		log.Println(err)
		return
	}

	err = DeployMulPlonk(client, sk)
	if err != nil {
		log.Println(err)
		return
	}

	err = DeployAddPlonk(client, sk)
	if err != nil {
		log.Println(err)
		return
	}
	err = DeployGPU(client, sk)
	if err != nil {
		log.Println(err)
		return
	}
	err = DeployModel(client, sk)
	if err != nil {
		log.Println(err)
		return
	}
	err = DeploySpace(client, sk)
	if err != nil {
		log.Println(err)
		return
	}
}

// deployall_v2 deploys all V2 contracts:
// 1. Deploy verifier contracts (RSOne, Add, Mul, Kzg)
// 2. Deploy implementation contracts (Epoch, Node, Piece, RSProof, EProof, EVerify)
// 3. Deploy ERC1967 proxies with initialization
// 4. Set cross-contract references (EProof in EVerify, addresses in Node)
// 5. Configure RS VK roots and minimum pledges
func deployall_v2(client *ethclient.Client, sk string) {
	// Get owner address from private key
	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		log.Println("Failed to create auth:", err)
		return
	}
	owner := au.From
	log.Printf("Using owner address: %s\n", owner.Hex())

	// Step 1: Deploy verifier contracts (non-upgradeable)
	log.Println("=== Deploying Verifier Contracts ===")
	rsoneAddr, err := DeployRSPlonkImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy RSOneVerifier:", err)
		return
	}

	addAddr, err := DeployAddPlonkImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy AddVerifier:", err)
		return
	}

	mulAddr, err := DeployMulPlonkImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy MulVerifier:", err)
		return
	}

	kzgAddr, err := DeployKZGPlonkImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy KzgVerifier:", err)
		return
	}

	// Step 2: Deploy implementation contracts
	log.Println("=== Deploying Implementation Contracts ===")
	epochImpl, err := DeployEpochImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy Epoch implementation:", err)
		return
	}

	nodeImpl, err := DeployNodeImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy Node implementation:", err)
		return
	}

	pieceImpl, err := DeployPieceImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy Piece implementation:", err)
		return
	}

	rsproofImpl, err := DeployRSProofImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy RSProof implementation:", err)
		return
	}

	eproofImpl, err := DeployEproofImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy EProof implementation:", err)
		return
	}

	everifyImpl, err := DeployEverifyImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy EVerify implementation:", err)
		return
	}

	// Step 3: Deploy proxy contracts with initialization
	log.Println("=== Deploying Proxy Contracts ===")

	// Deploy Epoch proxy
	epochProxy, err := DeployEpochProxy(client, sk, epochImpl, slots, owner)
	if err != nil {
		log.Println("Failed to deploy Epoch proxy:", err)
		return
	}

	// Deploy Node proxy
	nodeProxy, err := DeployNodeProxy(client, sk, nodeImpl, tokenAddr, epochProxy, owner)
	if err != nil {
		log.Println("Failed to deploy Node proxy:", err)
		return
	}

	// Deploy Piece proxy with init params
	pieceInitParams := piece.IPieceInitParams{
		Delay:       delay,
		MinStore:    minStore,
		MaxStore:    maxStore,
		MaxSize:     maxSize,
		MinPrice:    minPrice,
		StreamPrice: streamPrice,
	}
	pieceProxy, err := DeployPieceProxy(client, sk, pieceImpl, tokenAddr, epochProxy, nodeProxy, pieceInitParams, owner)
	if err != nil {
		log.Println("Failed to deploy Piece proxy:", err)
		return
	}

	// Deploy RSProof proxy with init params
	rsproofInitParams := rsproof.IRSProofInitParams{
		Token:        tokenAddr,
		Piece:        pieceProxy,
		Node:         nodeProxy,
		Base:         baseAddr,
		Rsone:        rsoneAddr,
		MinProveTime: minProveTime,
	}
	rsproofProxy, err := DeployRSProofProxy(client, sk, rsproofImpl, rsproofInitParams, owner)
	if err != nil {
		log.Println("Failed to deploy RSProof proxy:", err)
		return
	}

	// Set RS VK roots for different policies
	log.Println("=== Setting RS VK Roots ===")
	if err := SetRSVKRootV2(client, sk, 6, 4, rsproofProxy); err != nil {
		log.Println("Failed to set RS VK root for n=6, k=4:", err)
		return
	}
	if err := SetRSVKRootV2(client, sk, 14, 7, rsproofProxy); err != nil {
		log.Println("Failed to set RS VK root for n=14, k=7:", err)
		return
	}
	if err := SetRSVKRootV2(client, sk, 32, 16, rsproofProxy); err != nil {
		log.Println("Failed to set RS VK root for n=32, k=16:", err)
		return
	}
	if err := SetRSVKRootV2(client, sk, 64, 32, rsproofProxy); err != nil {
		log.Println("Failed to set RS VK root for n=64, k=32:", err)
		return
	}

	// Deploy EVerify proxy
	everifyProxy, err := DeployEVerifyProxy(client, sk, everifyImpl, epochProxy, pieceProxy, kzgAddr, mulAddr, addAddr, owner)
	if err != nil {
		log.Println("Failed to deploy EVerify proxy:", err)
		return
	}

	// Deploy EProof proxy with init params
	eproofInitParams := eproof.IEProofInitParams{
		Epoch:           epochProxy,
		Node:            nodeProxy,
		Piece:           pieceProxy,
		Token:           tokenAddr,
		Everify:         everifyProxy,
		Base:            baseAddr,
		ChallengeWindow: challengeWindow,
		MinProveTime:    minProveTime,
	}
	eproofProxy, err := DeployEProofProxy(client, sk, eproofImpl, eproofInitParams, owner)
	if err != nil {
		log.Println("Failed to deploy EProof proxy:", err)
		return
	}

	// Step 4: Set cross-contract references
	log.Println("=== Setting Cross-Contract References ===")

	// Set EProof address in EVerify
	if err := SetEProofAddress(client, sk, everifyProxy, eproofProxy); err != nil {
		log.Println("Failed to set EProof address in EVerify:", err)
		return
	}

	// Set EProof and RSProof addresses in Node
	if err := SetNodeAddresses(client, sk, nodeProxy, eproofProxy, rsproofProxy); err != nil {
		log.Println("Failed to set addresses in Node:", err)
		return
	}

	// Step 5: Set minimum pledge for different node types
	log.Println("=== Setting Node Minimum Pledges ===")
	for nodeType, pledge := range minPledgeMap {
		if err := SetMinPledgeV2(client, sk, nodeType, pledge, nodeProxy); err != nil {
			log.Printf("Failed to set min pledge for node type %d: %v", nodeType, err)
			return
		}
		log.Printf("Set min pledge for node type %d to %s", nodeType, pledge.String())
	}

	log.Println("=== V2 Deployment Complete ===")
	log.Printf("Summary:\n")
	log.Printf("  EpochProxy: %s\n", epochProxy.Hex())
	log.Printf("  NodeProxy: %s\n", nodeProxy.Hex())
	log.Printf("  PieceProxy: %s\n", pieceProxy.Hex())
	log.Printf("  RSProofProxy: %s\n", rsproofProxy.Hex())
	log.Printf("  EVerifyProxy: %s\n", everifyProxy.Hex())
	log.Printf("  EProofProxy: %s\n", eproofProxy.Hex())

	// Phase 2: DAO governance
	log.Println("")
	log.Println("=== Phase 2: DAO Governance Deployment ===")
	deployDAO(client, sk, owner, epochProxy, nodeProxy, pieceProxy, rsproofProxy, everifyProxy, eproofProxy)
}
