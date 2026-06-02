package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	contract "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/epoch"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/eproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/node"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/rsproof"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ProxyAddresses holds all proxy contract addresses
type ProxyAddresses struct {
	EpochProxy   common.Address
	NodeProxy    common.Address
	PieceProxy   common.Address
	RSProofProxy common.Address
	EVerifyProxy common.Address
	EProofProxy  common.Address
	StatProxy    common.Address
}

// ReadProxyAddresses reads proxy addresses from the out JSON file
func ReadProxyAddresses(outFile string) (*ProxyAddresses, error) {
	// Read the out file
	data, err := os.ReadFile(outFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read out file: %w", err)
	}

	// Parse the deployment history
	var history DeploymentHistory
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, fmt.Errorf("failed to parse out file: %w", err)
	}

	// Find the record for the current chain
	var currentRecord *DeploymentRecord
	for i := range history.Records {
		if history.Records[i].ChainURL == ChainURL {
			currentRecord = &history.Records[i]
			break
		}
	}

	if currentRecord == nil {
		return nil, fmt.Errorf("no deployment record found for chain: %s", ChainURL)
	}

	// Extract proxy addresses
	addresses := &ProxyAddresses{}

	if addr, ok := currentRecord.Deployments["EpochProxy"]; ok {
		addresses.EpochProxy = common.HexToAddress(addr)
	}
	if addr, ok := currentRecord.Deployments["NodeProxy"]; ok {
		addresses.NodeProxy = common.HexToAddress(addr)
	}
	if addr, ok := currentRecord.Deployments["PieceProxy"]; ok {
		addresses.PieceProxy = common.HexToAddress(addr)
	}
	if addr, ok := currentRecord.Deployments["RSProofProxy"]; ok {
		addresses.RSProofProxy = common.HexToAddress(addr)
	}
	if addr, ok := currentRecord.Deployments["EVerifyProxy"]; ok {
		addresses.EVerifyProxy = common.HexToAddress(addr)
	}
	if addr, ok := currentRecord.Deployments["EProofProxy"]; ok {
		addresses.EProofProxy = common.HexToAddress(addr)
	}
	if addr, ok := currentRecord.Deployments["StatProxy"]; ok {
		addresses.StatProxy = common.HexToAddress(addr)
	}

	log.Println("=== Proxy Addresses from out file ===")
	log.Printf("EpochProxy: %s\n", addresses.EpochProxy.Hex())
	log.Printf("NodeProxy: %s\n", addresses.NodeProxy.Hex())
	log.Printf("PieceProxy: %s\n", addresses.PieceProxy.Hex())
	log.Printf("RSProofProxy: %s\n", addresses.RSProofProxy.Hex())
	log.Printf("EVerifyProxy: %s\n", addresses.EVerifyProxy.Hex())
	log.Printf("EProofProxy: %s\n", addresses.EProofProxy.Hex())
	log.Printf("StatProxy: %s\n", addresses.StatProxy.Hex())

	return addresses, nil
}

// UpdateEpochSlots updates the slots value in the Epoch contract
func UpdateEpochSlots(client *ethclient.Client, sk string, epochProxy common.Address, newSlots uint64) error {
	log.Println("\n=== Updating Epoch slots ===")

	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	epochInstance, err := epoch.NewEpoch(epochProxy, client)
	if err != nil {
		return fmt.Errorf("failed to create epoch instance: %w", err)
	}

	// Get current slots value
	currentSlots, err := epochInstance.Slots(nil)
	if err != nil {
		return fmt.Errorf("failed to get current slots: %w", err)
	}
	log.Printf("Current slots: %d\n", currentSlots)

	// Update to new slots value
	tx, err := epochInstance.SetSlots(au, newSlots)
	if err != nil {
		return fmt.Errorf("failed to set slots: %w", err)
	}

	if err := contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return fmt.Errorf("failed to check transaction: %w", err)
	}

	log.Printf("Updated slots to: %d\n", newSlots)

	// Verify the update
	updatedSlots, err := epochInstance.Slots(nil)
	if err != nil {
		return fmt.Errorf("failed to verify updated slots: %w", err)
	}

	if updatedSlots != newSlots {
		return fmt.Errorf("slots update verification failed: expected %d, got %d", newSlots, updatedSlots)
	}

	log.Printf("Slots successfully updated and verified: %d\n", updatedSlots)
	return nil
}

// UpdateNodeMinStake updates minimum stake for all node types from minPledgeMap
func UpdateNodeMinStake(client *ethclient.Client, sk string, nodeProxy common.Address) error {
	log.Println("\n=== Updating Node minimum stakes ===")

	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	nodeInstance, err := node.NewNode(nodeProxy, client)
	if err != nil {
		return fmt.Errorf("failed to create node instance: %w", err)
	}

	// Update min stake for each node type in the map
	for nodeType, newMinStake := range minPledgeMap {
		// Get current min stake value
		currentMinStake, err := nodeInstance.MinStakeOf(nil, nodeType)
		if err != nil {
			log.Printf("Warning: failed to get current min stake for type %d: %v\n", nodeType, err)
			currentMinStake = big.NewInt(0)
		}
		log.Printf("Node type %d - Current min stake: %s\n", nodeType, currentMinStake.String())

		// Update to new min stake value
		tx, err := nodeInstance.Set(au, nodeType, newMinStake)
		if err != nil {
			return fmt.Errorf("failed to set min stake for node type %d: %w", nodeType, err)
		}

		if err := contract.CheckTx(ChainURL, tx.Hash()); err != nil {
			return fmt.Errorf("failed to check transaction for node type %d: %w", nodeType, err)
		}

		log.Printf("Node type %d - Updated min stake to: %s\n", nodeType, newMinStake.String())

		// Verify the update
		updatedMinStake, err := nodeInstance.MinStakeOf(nil, nodeType)
		if err != nil {
			return fmt.Errorf("failed to verify updated min stake for node type %d: %w", nodeType, err)
		}

		if updatedMinStake.Cmp(newMinStake) != 0 {
			return fmt.Errorf("min stake update verification failed for node type %d: expected %s, got %s",
				nodeType, newMinStake.String(), updatedMinStake.String())
		}

		log.Printf("Node type %d - MinStake successfully updated and verified: %s\n", nodeType, updatedMinStake.String())
	}

	return nil
}

// UpdateRSProofVKRoot updates VKRoot for specific RS parameters
func UpdateRSProofVKRoot(client *ethclient.Client, sk string, rsproofProxy common.Address, n, k uint8) error {
	log.Printf("\n=== Updating RSProof VKRoot for RS(%d,%d) ===\n", n, k)

	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	rsproofInstance, err := rsproof.NewRSProof(rsproofProxy, client)
	if err != nil {
		return fmt.Errorf("failed to create rsproof instance: %w", err)
	}

	// Get the VKRoot value for this RS policy
	vkRoot := new(big.Int)
	switch {
	case n == 6 && k == 4:
		vkRoot.SetString(contract.RSn6k4VKRoot, 10)
	case n == 14 && k == 7:
		vkRoot.SetString(contract.RSn14k7VKRoot, 10)
	case n == 32 && k == 16:
		vkRoot.SetString(contract.RSn32k16VKRoot, 10)
	case n == 64 && k == 32:
		vkRoot.SetString(contract.RSn64k32VKRoot, 10)
	default:
		return fmt.Errorf("unsupported RS policy: n=%d, k=%d", n, k)
	}

	// Get current VKRoot value
	currentVKRoot, err := rsproofInstance.GetVKRoot(nil, n, k)
	if err != nil {
		log.Printf("Warning: failed to get current VKRoot: %v\n", err)
		currentVKRoot = big.NewInt(0)
	}
	log.Printf("Current VKRoot for RS(%d,%d): %s\n", n, k, currentVKRoot.String())

	// Only set if not already set
	if currentVKRoot.Cmp(big.NewInt(0)) == 0 {
		tx, err := rsproofInstance.SetVKRoot(au, n, k, vkRoot)
		if err != nil {
			return fmt.Errorf("failed to set VKRoot: %w", err)
		}

		if err := contract.CheckTx(ChainURL, tx.Hash()); err != nil {
			return fmt.Errorf("failed to check transaction: %w", err)
		}

		log.Printf("Set VKRoot for RS(%d,%d) to: %s\n", n, k, vkRoot.String())

		// Verify the update
		updatedVKRoot, err := rsproofInstance.GetVKRoot(nil, n, k)
		if err != nil {
			return fmt.Errorf("failed to verify updated VKRoot: %w", err)
		}

		if updatedVKRoot.Cmp(vkRoot) != 0 {
			return fmt.Errorf("VKRoot update verification failed: expected %s, got %s", vkRoot.String(), updatedVKRoot.String())
		}

		log.Printf("VKRoot successfully set and verified: %s\n", updatedVKRoot.String())
	} else {
		log.Printf("VKRoot already set for RS(%d,%d), skipping update\n", n, k)
	}

	return nil
}

// UpdateRSProofMinProveTime updates the minimum prove time for RSProof
func UpdateRSProofMinProveTime(client *ethclient.Client, sk string, rsproofProxy common.Address, newMinProveTime *big.Int) error {
	log.Println("\n=== Updating RSProof minimum prove time ===")

	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	rsproofInstance, err := rsproof.NewRSProof(rsproofProxy, client)
	if err != nil {
		return fmt.Errorf("failed to create rsproof instance: %w", err)
	}

	// Get current min prove time
	currentMinProveTime, err := rsproofInstance.MinProveTime(nil)
	if err != nil {
		return fmt.Errorf("failed to get current min prove time: %w", err)
	}
	log.Printf("Current RS min prove time: %s\n", currentMinProveTime.String())

	// Update to new min prove time
	tx, err := rsproofInstance.SetMinProveTime(au, newMinProveTime)
	if err != nil {
		return fmt.Errorf("failed to set min prove time: %w", err)
	}

	if err := contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return fmt.Errorf("failed to check transaction: %w", err)
	}

	log.Printf("Updated RS min prove time to: %s\n", newMinProveTime.String())

	// Verify the update
	updatedMinProveTime, err := rsproofInstance.MinProveTime(nil)
	if err != nil {
		return fmt.Errorf("failed to verify updated min prove time: %w", err)
	}

	if updatedMinProveTime.Cmp(newMinProveTime) != 0 {
		return fmt.Errorf("RS min prove time update verification failed: expected %s, got %s",
			newMinProveTime.String(), updatedMinProveTime.String())
	}

	log.Printf("RS MinProveTime successfully updated and verified: %s\n", updatedMinProveTime.String())
	return nil
}

// UpdateRSProofBasePenalty updates the basePenalty for RSProof
func UpdateRSProofBasePenalty(client *ethclient.Client, sk string, rsproofProxy common.Address, newBasePenalty *big.Int) error {
	log.Println("\n=== Updating RSProof basePenalty ===")

	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	rsproofInstance, err := rsproof.NewRSProof(rsproofProxy, client)
	if err != nil {
		return fmt.Errorf("failed to create rsproof instance: %w", err)
	}

	// Get current 
	current, err := rsproofInstance.BasePenalty(nil)
	if err != nil {
		return fmt.Errorf("failed to get current basePenalty: %w", err)
	}
	log.Printf("Current RS basePenalty: %s\n", current.String())

	// Update to new
	tx, err := rsproofInstance.SetBasePenalty(au, newBasePenalty)
	if err != nil {
		return fmt.Errorf("failed to set basePenalty: %w", err)
	}

	if err := contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return fmt.Errorf("failed to check transaction: %w", err)
	}

	log.Printf("Updated RS basePenalty to: %s\n", newBasePenalty.String())

	// Verify the update
	updated, err := rsproofInstance.BasePenalty(nil)
	if err != nil {
		return fmt.Errorf("failed to verify updated basePenalty: %w", err)
	}

	if updated.Cmp(newBasePenalty) != 0 {
		return fmt.Errorf("RS basePenalty update verification failed: expected %s, got %s",
			newBasePenalty.String(), updated.String())
	}

	log.Printf("RS BasePenalty successfully updated and verified: %s\n", updated.String())
	return nil
}

// UpdateEProofMinProveTime updates the minimum prove time for EProof
func UpdateEProofMinProveTime(client *ethclient.Client, sk string, eproofProxy common.Address, newMinProveTime *big.Int) error {
	log.Println("\n=== Updating EProof minimum prove time ===")

	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	eproofInstance, err := eproof.NewEProof(eproofProxy, client)
	if err != nil {
		return fmt.Errorf("failed to create eproof instance: %w", err)
	}

	// Get current min prove time
	currentMinProveTime, err := eproofInstance.MinProveTime(nil)
	if err != nil {
		return fmt.Errorf("failed to get current min prove time: %w", err)
	}
	log.Printf("Current E min prove time: %s\n", currentMinProveTime.String())

	// Update to new min prove time
	tx, err := eproofInstance.SetMinProveTime(au, newMinProveTime)
	if err != nil {
		return fmt.Errorf("failed to set min prove time: %w", err)
	}

	if err := contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return fmt.Errorf("failed to check transaction: %w", err)
	}

	log.Printf("Updated E min prove time to: %s\n", newMinProveTime.String())

	// Verify the update
	updatedMinProveTime, err := eproofInstance.MinProveTime(nil)
	if err != nil {
		return fmt.Errorf("failed to verify updated min prove time: %w", err)
	}

	if updatedMinProveTime.Cmp(newMinProveTime) != 0 {
		return fmt.Errorf("EProof min prove time update verification failed: expected %s, got %s",
			newMinProveTime.String(), updatedMinProveTime.String())
	}

	log.Printf("E MinProveTime successfully updated and verified: %s\n", updatedMinProveTime.String())
	return nil
}

// UpdateEProofBasePenalty updates the basePenalty for EProof
func UpdateEProofBasePenalty(client *ethclient.Client, sk string, eproofProxy common.Address, newBasePenalty *big.Int) error {
	log.Println("\n=== Updating EProof basePenalty ===")

	au, err := contract.MakeAuth(ChainURL, ChainID, sk)
	if err != nil {
		return err
	}

	eproofInstance, err := eproof.NewEProof(eproofProxy, client)
	if err != nil {
		return fmt.Errorf("failed to create eproof instance: %w", err)
	}

	// Get current
	current, err := eproofInstance.BasePenalty(nil)
	if err != nil {
		return fmt.Errorf("failed to get current basePenalty: %w", err)
	}
	log.Printf("Current E basePenalty: %s\n", current.String())

	// Update to new
	tx, err := eproofInstance.SetBasePenalty(au, newBasePenalty)
	if err != nil {
		return fmt.Errorf("failed to set basePenalty: %w", err)
	}

	if err := contract.CheckTx(ChainURL, tx.Hash()); err != nil {
		return fmt.Errorf("failed to check transaction: %w", err)
	}

	log.Printf("Updated E basePenalty to: %s\n", newBasePenalty.String())

	// Verify the update
	updated, err := eproofInstance.BasePenalty(nil)
	if err != nil {
		return fmt.Errorf("failed to verify updated basePenalty: %w", err)
	}

	if updated.Cmp(newBasePenalty) != 0 {
		return fmt.Errorf("EProof basePenalty update verification failed: expected %s, got %s",
			newBasePenalty.String(), updated.String())
	}

	log.Printf("E BasePenalty successfully updated and verified: %s\n", updated.String())
	return nil
}
