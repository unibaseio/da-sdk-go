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
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ChainURL = contract.BNBTestnetChainRPC
	ChainID  = contract.BNBTestnetChainID

	blockInterval  = uint64(3)
	bankAddr       = contract.BNBTestnetBankAddr
	tokenAddr      = contract.BNBTestnetTokenAddr
	deployDAOPhase = true
)

// V2 deployment configuration
var (
	// slots / minProveTime are in BLOCKS, so their wall-clock depends on the
	// chain's block interval. These defaults assume ~450ms blocks (16000*450ms
	// ≈ 2h epoch, 8000*450ms ≈ 1h prove window). ⚠️ base-sepolia is ~2s blocks,
	// so the SAME block counts become ~8.9h / ~4.4h — recalibrate per chain via
	// -slots / -min-prove-time (e.g. base-sepolia: -slots 3600 -min-prove-time
	// 1800 for ~2h / ~1h). Safety note: the prove window must comfortably exceed
	// the BN254 proof wall-clock (~4-5 min), which 1800*2s=1h does.
	slots    = uint64(16000)    // Epoch slots (blocks); see note above
	delay    = uint64(7)        // Piece delay in epochs
	minStore = uint64(1200)     // Minimum storage time, 1200 epochs = 100 days
	maxStore = uint64(12000)    // Maximum storage time
	maxSize  = uint64(33554432) // 32MB
	// Economic starting values (governance-adjustable post-launch). At UB=$0.1
	// minPrice 1e11 ≈ ~$3.6/TB/month (pre-redundancy) — between Filecoin and S3.
	// ⚠️ validate vs storer cost + market with the AddPiece val formula before
	// mainnet; was 1e8/1e9 ("just for test") — do NOT ship those.
	minPrice        = big.NewInt(1e11) // minimum storage price per unit, per epoch*MB
	streamPrice     = big.NewInt(1e12) // streaming price, per replica
	minProveTime    = big.NewInt(8000) // Min prove time (blocks); ~1h @450ms — see slots note
	challengeWindow = uint64(7)        // Challenge window in epochs for EProof
	minPledgeMap    = map[uint8]*big.Int{
		// type 1 (store) min pledge must cover the fraud-proof penalty
		// (challenge() locks basePenalty of the defender's stake — an
		// under-pledged store could not even be challenged)
		1: new(big.Int).Mul(contract.DefaultPenalty, big.NewInt(5)), // 5x penalty = 50000 UB (A2: concurrent-challenge capacity = 5)
		2: new(big.Int).Mul(big.NewInt(1e15), big.NewInt(10)),       // test
		3: new(big.Int).Mul(big.NewInt(1e15), big.NewInt(10)),       // test
	}

	baseAddr = common.HexToAddress("0xE0AD379735ba88B323298D091ff3b67Dd6C79852")
)

// DAO governance configuration
var (
	daoTokenName         = "UBDAO Governance Token"
	daoTokenSymbol       = "vUB"
	daoTokenSupply       = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1_000_000_000)) // 1B tokens
	daoVotingDelay       = big.NewInt(1)                                                  // 1 block
	daoVotingPeriod      = uint32(24_000)                                                 // ~3h at ~450ms block time on BSC
	daoProposalThreshold = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(2_500_000))      // 2.5M tokens (~0.25%)
	daoQuorumFraction    = big.NewInt(4)                                                  // 4%
	daoTimelockDelay     = big.NewInt(0)                                                  // no delay for testnet

	// governance token: "vub" (stake UB -> ve-weighted votes, default) or
	// "legacy" (standalone placeholder GovernanceToken). Override with -dao-gov-token.
	daoGovTokenKind = "vub"
	daoRewardToken  = common.Address{} // reward token for vUB APY; empty -> reuse UB token
)

func init() {
	dlog.SetLogLevel("DEBUG")
}

func main() {
	sk := flag.String("sk", "", "private key for sending transaction")
	rpc := flag.String("rpc", "", "chain rpc endpoint (default: bnb-testnet)")
	chainID := flag.Int64("chainid", 0, "chain id (default: bnb-testnet)")
	skipDAO := flag.Bool("skip-dao", false, "skip phase-2 DAO governance deployment")
	govToken := flag.String("dao-gov-token", "vub", "DAO governance token: vub (stake UB->ve votes) or legacy (placeholder GovernanceToken)")
	rewardTok := flag.String("dao-reward-token", "", "reward token address for vUB APY (default: reuse UB token)")
	slotsFlag := flag.Uint64("slots", 0, "epoch length in blocks (default 16000; use a small value on local anvil)")
	mptFlag := flag.Int64("min-prove-time", 0, "min prove time in blocks (default 8000)")
	tokenFlag := flag.String("token", "", "reuse an EXISTING ERC20 as the stake/penalty token (skip deploying a fresh test token); e.g. the canonical UB on this chain")
	flag.Parse()

	daoGovTokenKind = *govToken
	if *rewardTok != "" {
		daoRewardToken = common.HexToAddress(*rewardTok)
	}

	if *rpc != "" {
		ChainURL = *rpc
	}
	if *chainID != 0 {
		ChainID = *chainID
	}
	if *slotsFlag != 0 {
		slots = *slotsFlag
	}
	if *mptFlag != 0 {
		minProveTime = big.NewInt(*mptFlag)
	}
	deployDAOPhase = !*skipDAO

	fmt.Println("connect to: ", ChainURL)
	client, err := ethclient.DialContext(context.TODO(), ChainURL)
	if err != nil {
		return
	}
	defer client.Close()
	if *tokenFlag != "" {
		// reuse an existing ERC20 (e.g. the canonical UB) instead of deploying a
		// fresh test token — all 5 token-bearing impls get this address at init.
		tokenAddr = common.HexToAddress(*tokenFlag)
		log.Println("reusing existing token: ", tokenAddr.Hex())
	} else {
		DeployTokenTest(client, *sk)
	}
	deployall_v2(client, *sk)
}

func DeployTokenTest(client *ethclient.Client, sk string) error {
	au, err := makeAuth(sk)
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

	for _, p := range types.SupportedPolicies {
		err = SetRSVKRoot(client, sk, int(p.N), int(p.K))
		if err != nil {
			log.Println(err)
			return
		}
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
	// Get owner address from private key WITHOUT makeAuth — makeAuth advances the
	// local nonce counter, so calling it just to read the address would burn a
	// nonce that no tx ever uses, leaving an unminable gap (every later tx sits
	// queued behind the missing nonce). Derive the address directly instead.
	pk, err := crypto.HexToECDSA(sk)
	if err != nil {
		log.Println("Failed to parse key:", err)
		return
	}
	owner := crypto.PubkeyToAddress(pk.PublicKey)
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
	if err := waitForCode(client, rsproofProxy); err != nil { // guard RPC lag
		log.Println("RSProof proxy code not visible:", err)
		return
	}
	for _, p := range types.SupportedPolicies {
		if err := SetRSVKRootV2(client, sk, int(p.N), int(p.K), rsproofProxy); err != nil {
			log.Printf("Failed to set RS VK root for n=%d, k=%d: %v", p.N, p.K, err)
			return
		}
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
	if err := waitForCode(client, everifyProxy); err != nil { // guard RPC lag
		log.Println("EVerify proxy code not visible:", err)
		return
	}
	if err := waitForCode(client, eproofProxy); err != nil {
		log.Println("EProof proxy code not visible:", err)
		return
	}

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

	// Step 6: set the fraud-proof penalty (initialize defaults to 1e18;
	// dev-confirmed value is DefaultPenalty = 10000 UB)
	log.Println("=== Setting Base Penalty ===")
	if err := SetBasePenaltyV2(client, sk, rsproofProxy, eproofProxy, contract.DefaultPenalty); err != nil {
		log.Println("Failed to set base penalty:", err)
		return
	}

	// Step 7: ValidatorReward pool (FixB+A2) — deployed LAST so it does not
	// shift the deterministic addresses of the core proxies (the LocalAnvil
	// table + integration tests depend on the existing deploy order). Repoint
	// RSProof/EProof's penalty share to it; treasury funds it off-chain.
	log.Println("=== Deploying ValidatorReward (FixB+A2) ===")
	validatorRewardImpl, err := DeployValidatorRewardImpl(client, sk)
	if err != nil {
		log.Println("Failed to deploy ValidatorReward implementation:", err)
		return
	}
	validatorPool, err := DeployValidatorRewardProxy(client, sk, validatorRewardImpl, tokenAddr, owner)
	if err != nil {
		log.Println("Failed to deploy ValidatorReward proxy:", err)
		return
	}
	if err := waitForCode(client, validatorPool); err != nil { // guard RPC lag
		log.Println("ValidatorReward proxy code not visible:", err)
		return
	}
	if err := SetValidatorPool(client, sk, rsproofProxy, eproofProxy, validatorPool); err != nil {
		log.Println("Failed to set validator pool on RSProof/EProof:", err)
		return
	}

	log.Println("=== V2 Deployment Complete ===")
	log.Printf("Summary:\n")
	log.Printf("  EpochProxy: %s\n", epochProxy.Hex())
	log.Printf("  NodeProxy: %s\n", nodeProxy.Hex())
	log.Printf("  PieceProxy: %s\n", pieceProxy.Hex())
	log.Printf("  RSProofProxy: %s\n", rsproofProxy.Hex())
	log.Printf("  EVerifyProxy: %s\n", everifyProxy.Hex())
	log.Printf("  EProofProxy: %s\n", eproofProxy.Hex())
	log.Printf("  ValidatorRewardProxy: %s\n", validatorPool.Hex())

	// Phase 2: DAO governance
	if deployDAOPhase {
		log.Println("")
		log.Println("=== Phase 2: DAO Governance Deployment ===")
		deployDAO(client, sk, owner, tokenAddr, daoRewardToken, epochProxy, nodeProxy, pieceProxy, rsproofProxy, everifyProxy, eproofProxy)
	}
}
