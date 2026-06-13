package common

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/unibaseio/da-sdk-go/build"
	"github.com/unibaseio/da-sdk-go/contract/v1/go/token"
	"github.com/unibaseio/da-sdk-go/lib/env"
	dlog "github.com/unibaseio/da-sdk-go/lib/log"
	"github.com/unibaseio/da-sdk-go/lib/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	KZGVKRoot = "3b8201b322c65a735690cf82c850b8624c29ec05400e06ba92a9aad12c37c1605812abbc9a1a11f500b3ab28b7751b52"

	RSn6k4VKRoot   = "6182580396057136035349698244013672118159269515303422592734042707265223434376"
	RSn14k7VKRoot  = "10373754662900994857464626600297287716762049534190727150581258556168809741757"
	RSn32k16VKRoot = "17745558306831082913906433185923529664704368402036214777413791674807342977788"
	RSn64k32VKRoot = "16234120671149928180196435146616388668330002147244144092564388600958703384116"

	INKZGVKRoot = "9235750480968672023981154929687398144904220020071947383737661744780443258323"
	INMulVKRoot = "20271222645096249312949571772787165775593093202630226028249238982739427154007"
	INAddVKRoot = "19807320097661814426496674775485679654967276569243947871856251314675094662138"
)

var (
	//DevChain   = "http://54.254.72.127:8501"
	//DevChainID = 222

	//http://unibasechain-scan-405529765.ap-southeast-1.elb.amazonaws.com/
	//L1Bridge   = common.HexToAddress("0xc072613dAaab3E9BcC8dDd23aE7c368DC5751984")
	//DevChain   = "https://chain.unibase.io"
	//DevChainID = 43134

	//DevChain   = "https://ethereum-holesky.publicnode.com"
	//DevChainID = 17000
	//BankAddr   = common.HexToAddress("0x6c579D5eF7846E2c6cE255Adc2E0BEF1411fEB5c")
	//TokenAddr  = common.HexToAddress("0x421BfaFCfa9370c64F65100246D02913Bc9079F4")
	//SyncHeight   = 2_391_000

	//https://sepolia-optimism.etherscan.io/
	//DevChain   = "https://11155420.rpc.thirdweb.com"

	DefaultGasLimit = 10_000_000
	DefaultGasPrice = 100_000_000 // 0.1gwei

	DefaultStreamPrice  = 1e12
	DefaultReplicaPrice = 1e11 // 1TB*100 epoch cost 10
	DefaultStoreEpoch   = 1201 // slight larger than minEpoch
	DelaySubmit         = 7

	DefaultSpacePrice = 1e10
	DefaultSpaceEpoch = 201

	// DefaultPenalty is the fraud-proof challenge stake / slash amount:
	// 10000 UB (decimals 18). Must match basePenalty in RSProof.sol/
	// EProof.sol, and Node minPledge(type 1) must be >= it. The value
	// exceeds int64, so it is a *big.Int — treat it as read-only.
	DefaultPenalty = new(big.Int).Mul(big.NewInt(10_000), big.NewInt(1e18))

	Base = common.HexToAddress("0x61Ea24745A3F7Bcbb67eD95B674fEcfbb331ABd0")
)

// local-anvil
var (
	LocalAnvil                     = build.LocalAnvil
	LocalAnvilChainRPC             = "http://127.0.0.1:8545"
	LocalAnvilChainRPCForFilterLog = "http://127.0.0.1:8545"
	LocalAnvilChainID              = 56
	LocalAnvilTokenAddr            = common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	LocalAnvilSyncHeight           = 0_000_000

	// use proxy contract address
	LocalAnvilEpochAddr   = common.HexToAddress("0xB7f8BC63BbcaD18155201308C8f3540b07f84F5e")
	LocalAnvilNodeAddr    = common.HexToAddress("0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0")
	LocalAnvilPieceAddr   = common.HexToAddress("0x0DCd1Bf9A1b36cE34237eEaFef220932846BCD82")
	LocalAnvilRSProofAddr = common.HexToAddress("0x9A676e781A523b5d0C0e43731313A708CB607508")
	LocalAnvilEProofAddr  = common.HexToAddress("0xc6e7DF5E7b4f2A278906862b61205850344D4e7d")
	LocalAnvilEVerifyAddr = common.HexToAddress("0x3Aa5ebB10DC797CAC828524e59A333d0A371443c")
	LocalAnvilStatAddr    = common.HexToAddress("")

	LocalAnvilRSOneAddr = common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
	LocalAnvilKZGAddr   = common.HexToAddress("0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9")
	LocalAnvilAddAddr   = common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0")
	LocalAnvilMulAddr   = common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9")
)

// base-sepolia (first deployment target; DA contract addresses TBD — fill in
// after deployment, together with the chain's KZGVKRoot)
var (
	BaseSepolia                     = build.BaseSepolia
	BaseSepoliaExplorer             = "https://sepolia.basescan.org/"
	BaseSepoliaChainRPC             = "https://base-sepolia-rpc.publicnode.com"
	BaseSepoliaChainRPCForFilterLog = "https://base-sepolia-rpc.publicnode.com"
	BaseSepoliaChainID              = int64(84532)
	BaseSepoliaTokenAddr            = common.HexToAddress("") // TODO: test UB/ERC20 on base-sepolia
	BaseSepoliaSyncHeight           = 0                       // TODO: set to deployment height

	BaseSepoliaEpochAddr   = common.HexToAddress("")
	BaseSepoliaNodeAddr    = common.HexToAddress("")
	BaseSepoliaPieceAddr   = common.HexToAddress("")
	BaseSepoliaRSProofAddr = common.HexToAddress("")
	BaseSepoliaEProofAddr  = common.HexToAddress("")
	BaseSepoliaEVerifyAddr = common.HexToAddress("")
	BaseSepoliaStatAddr    = common.HexToAddress("")

	BaseSepoliaRSOneAddr = common.HexToAddress("")
	BaseSepoliaKZGAddr   = common.HexToAddress("")
	BaseSepoliaAddAddr   = common.HexToAddress("")
	BaseSepoliaMulAddr   = common.HexToAddress("")
)

// base mainnet (deploy order #1; DA contracts + UB OFT bridge addr TBD)
var (
	BaseMainnet                     = build.BaseMainnet
	BaseMainnetExplorer             = "https://basescan.org/"
	BaseMainnetChainRPC             = "https://base-rpc.publicnode.com"
	BaseMainnetChainRPCForFilterLog = "https://base-rpc.publicnode.com"
	BaseMainnetChainID              = int64(8453)
	BaseMainnetTokenAddr            = common.HexToAddress("") // TODO: UB LayerZero OFT on BASE (bridge deployment pending)
	BaseMainnetSyncHeight           = 0                       // TODO: set to deployment height

	BaseMainnetEpochAddr   = common.HexToAddress("")
	BaseMainnetNodeAddr    = common.HexToAddress("")
	BaseMainnetPieceAddr   = common.HexToAddress("")
	BaseMainnetRSProofAddr = common.HexToAddress("")
	BaseMainnetEProofAddr  = common.HexToAddress("")
	BaseMainnetEVerifyAddr = common.HexToAddress("")
	BaseMainnetStatAddr    = common.HexToAddress("")

	BaseMainnetRSOneAddr = common.HexToAddress("")
	BaseMainnetKZGAddr   = common.HexToAddress("")
	BaseMainnetAddAddr   = common.HexToAddress("")
	BaseMainnetMulAddr   = common.HexToAddress("")
)

// bsc mainnet (deploy order #2)
// ⚠️ TokenAddr: UB LayerZero OFT bridge on BSC. Two conflicting deployment
// records exist for mainnet UB — verify the canonical address before any
// mainnet use of this table.
var (
	BSCMainnet                     = build.BSCMainnet
	BSCMainnetExplorer             = "https://bscscan.com/"
	BSCMainnetChainRPC             = "https://bsc-rpc.publicnode.com"
	BSCMainnetChainRPCForFilterLog = "https://bsc-rpc.publicnode.com"
	BSCMainnetChainID              = int64(56)
	BSCMainnetTokenAddr            = common.HexToAddress("0x40b8129B786D766267A7a118cF8C07E31CDB6Fde")
	BSCMainnetSyncHeight           = 0 // TODO: set to deployment height

	BSCMainnetEpochAddr   = common.HexToAddress("")
	BSCMainnetNodeAddr    = common.HexToAddress("")
	BSCMainnetPieceAddr   = common.HexToAddress("")
	BSCMainnetRSProofAddr = common.HexToAddress("")
	BSCMainnetEProofAddr  = common.HexToAddress("")
	BSCMainnetEVerifyAddr = common.HexToAddress("")
	BSCMainnetStatAddr    = common.HexToAddress("")

	BSCMainnetRSOneAddr = common.HexToAddress("")
	BSCMainnetKZGAddr   = common.HexToAddress("")
	BSCMainnetAddAddr   = common.HexToAddress("")
	BSCMainnetMulAddr   = common.HexToAddress("")
)

// eth mainnet (deploy order #3)
// ⚠️ TokenAddr: UB canonical ledger on ETH. Two conflicting deployment
// records exist for mainnet UB — verify the canonical address before any
// mainnet use of this table.
var (
	ETHMainnet                     = build.ETHMainnet
	ETHMainnetExplorer             = "https://etherscan.io/"
	ETHMainnetChainRPC             = "https://ethereum-rpc.publicnode.com"
	ETHMainnetChainRPCForFilterLog = "https://ethereum-rpc.publicnode.com"
	ETHMainnetChainID              = int64(1)
	ETHMainnetTokenAddr            = common.HexToAddress("0x6944E1DF6Bf5972305f9Ab25dF47ef10De01bcc8")
	ETHMainnetSyncHeight           = 0 // TODO: set to deployment height

	ETHMainnetEpochAddr   = common.HexToAddress("")
	ETHMainnetNodeAddr    = common.HexToAddress("")
	ETHMainnetPieceAddr   = common.HexToAddress("")
	ETHMainnetRSProofAddr = common.HexToAddress("")
	ETHMainnetEProofAddr  = common.HexToAddress("")
	ETHMainnetEVerifyAddr = common.HexToAddress("")
	ETHMainnetStatAddr    = common.HexToAddress("")

	ETHMainnetRSOneAddr = common.HexToAddress("")
	ETHMainnetKZGAddr   = common.HexToAddress("")
	ETHMainnetAddAddr   = common.HexToAddress("")
	ETHMainnetMulAddr   = common.HexToAddress("")
)

// bnb-testnet-v2
var (
	BNBTestnetV2                   = build.BNBTestnetV2
	BNBTestnetExplorer             = "https://testnet.bscscan.com/"
	BNBTestnetChainRPC             = "https://bsc-testnet-dataseed.bnbchain.org"
	BNBTestnetChainRPCForFilterLog = "https://bsc-prebsc-dataseed.bnbchain.org"
	BNBTestnetChainID              = int64(97)
	BNBTestnetBankAddr             = common.HexToAddress("0x5903805A3a50Fab318c8650bABC71F58900EE34e")
	BNBTestnetTokenAddr            = common.HexToAddress("0xC488F83A897E8AFF387D4124D46a63Dd33cb9c97")
	BNBTestnetSyncHeight           = 80_542_700

	BNBTestnetEpochAddr   = common.HexToAddress("0xf80Ff1FE31ac5872D0366aCAAF2BDa8a28AE2cA8")
	BNBTestnetNodeAddr    = common.HexToAddress("0x16c2A3634E71eC14e09cafbe67c6aBC06AE06Eb8")
	BNBTestnetPieceAddr   = common.HexToAddress("0x00CDaB61bc0bd8055D27E770A6Ee9149BCbd4fb7")
	BNBTestnetRSProofAddr = common.HexToAddress("0xdB8C0bf2B8510f729ce17b7048F98B0F8F757c7A")
	BNBTestnetEProofAddr  = common.HexToAddress("0xba0F80B3395c8e0722FF92A04aF315ab14Ee5C60")
	BNBTestnetEVerifyAddr = common.HexToAddress("0xBFFfD0708Ef5CE588622F2961B30D4BA8baD3072")
	BNBTestnetStatAddr    = common.HexToAddress("")

	BNBTestnetRSOneAddr = common.HexToAddress("0x8f8aA4BcC6f8A5eA36E8DFB8fCB14efEab5460d9")
	BNBTestnetKZGAddr   = common.HexToAddress("0xE78Eb0B875772E0C464DaAa2054AE8D7F4a7c06A")
	BNBTestnetAddAddr   = common.HexToAddress("0x37222C0a687079968a039874748218e0C43BFf65")
	BNBTestnetMulAddr   = common.HexToAddress("0x7581C27FC358208F3d64A0cd2E5733290D7C0CD9")
)

// bnb-testnet-dao
var (
	BNBTestnetDAO                     = build.BNBTestnetDAO
	BNBTestnetDAOExplorer             = "https://testnet.bscscan.com/"
	BNBTestnetDAOChainRPC             = "https://bsc-testnet-dataseed.bnbchain.org"
	BNBTestnetDAOChainRPCForFilterLog = "https://bsc-prebsc-dataseed.bnbchain.org"
	BNBTestnetDAOChainID              = int64(97)
	BNBTestnetDAOBankAddr             = common.HexToAddress("0x5903805A3a50Fab318c8650bABC71F58900EE34e")
	BNBTestnetDAOTokenAddr            = common.HexToAddress("0x7c5AC49563d7046906333AC2CA28E602327CFe1A")
	BNBTestnetDAOSyncHeight           = 108_347_000

	BNBTestnetDAOEpochAddr   = common.HexToAddress("0xfBF8CbDCFb200eD20e69e4F89b39CbC7A18855E6")
	BNBTestnetDAONodeAddr    = common.HexToAddress("0xAF83D927C968eDC0E50232536Aac2245b4E048d9")
	BNBTestnetDAOPieceAddr   = common.HexToAddress("0x0b03B50c5A65051F0b33b1155f6370b3EE8394c6")
	BNBTestnetDAORSProofAddr = common.HexToAddress("0x97Ce0b5006CD830a38635FE7001C0b988536f398")
	BNBTestnetDAOEProofAddr  = common.HexToAddress("0x8F4FfA5FbDA95762DF1B6aC19272DA6aa78A3860")
	BNBTestnetDAOEVerifyAddr = common.HexToAddress("0xe8a6b34880828329A2371d43184f6Cb6298C1d20")
	BNBTestnetDAOStatAddr    = common.HexToAddress("")

	BNBTestnetDAORSOneAddr = common.HexToAddress("0xFe01d13E43d9f551Acc4469c3FF4da3e2FF3e3C7")
	BNBTestnetDAOKZGAddr   = common.HexToAddress("0x804D9F4AEAa3A44B03A28c35c522c6E1cf7c371c")
	BNBTestnetDAOAddAddr   = common.HexToAddress("0x5E87B084d8A7207664aE408280D2BdEdA4F7f30D")
	BNBTestnetDAOMulAddr   = common.HexToAddress("0x9e0fb9415ED4306EB331532c130f896350765A0F")

	BNBTestnetDAOTimelockAddr = common.HexToAddress("0xfbEc1a00623d01AfFbBE68cca021c3f6fe08fD2a")
	BNBTestnetDAOGovernorAddr = common.HexToAddress("0x8c760975545030CEBF16C0f6E109d90Bc77B5Ed9")
	BNBTestnetDAOERC20Addr    = common.HexToAddress("0x53B5bC1cDc844B02F4BFeA71429657e5078c0349")
)

var Logger = dlog.Logger("contract")

func init() {
	DefaultGasPrice = env.Int(env.GasPrice, DefaultGasPrice)
	DefaultGasLimit = env.Int(env.GasLimit, DefaultGasLimit)

	Logger.Infof("gas price: %d", DefaultGasPrice)
	Logger.Infof("gas limit: %d", DefaultGasLimit)
}

func MakeAuth(ep string, chainID int64, hexSk string) (*bind.TransactOpts, error) {
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		return nil, err
	}

	return MakeAuthBySk(ep, big.NewInt(chainID), sk)
}

func CheckTx(ep string, txHash common.Hash) error {
	return checkTx(ep, txHash)
}

func MakeAuthBySk(ep string, chainID *big.Int, sk *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	auth := &bind.TransactOpts{}
	auth, err := bind.NewKeyedTransactorWithChainID(sk, chainID)
	if err != nil {
		return nil, fmt.Errorf("NewKeyedTransaction failed %s", err)
	}

	auth.Value = big.NewInt(0)
	// TODO(P1): switch proof-heavy txs to EstimateGas + safety margin once the
	// margin ratio is confirmed; a fixed limit either wastes or underprovisions.
	auth.GasLimit = uint64(DefaultGasLimit)
	client, err := ethclient.Dial(ep)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	Logger.Debugf("height: %d, basefee: %d, blob: %d", header.Number, header.BaseFee, header.BlobGasUsed)

	if header.BaseFee == nil || header.BaseFee.BitLen() == 0 {
		// pre-1559 chain, or zero base fee (e.g. BSC): keep legacy pricing
		auth.GasPrice = big.NewInt(int64(DefaultGasPrice))
		Logger.Debugf("no basefee, set legacy gas price to %d", DefaultGasPrice)
		return auth, nil
	}

	// EIP-1559 chain (ETH/BASE): dynamic-fee tx. The old baseFee*1.2 GasPrice
	// implied a tiny priority fee when baseFee is low and the tx could sit
	// unmined; price tip and fee cap separately instead.
	tip, err := client.SuggestGasTipCap(ctx)
	if err != nil || tip.Sign() == 0 {
		tip = big.NewInt(int64(DefaultGasPrice)) // fallback floor (0.1 gwei default)
		Logger.Debugf("suggest tip unavailable, fallback to %d", tip)
	}
	if v := env.Str(env.GasTip, ""); v != "" {
		if t, ok := new(big.Int).SetString(v, 10); ok {
			tip = t
		} else {
			Logger.Warn("invalid GAS_TIP: ", v)
		}
	}
	// feeCap = 2*baseFee + tip: survives a doubling of baseFee while pending;
	// only the actual baseFee+tip is paid, the cap is not.
	feeCap := new(big.Int).Mul(header.BaseFee, big.NewInt(2))
	feeCap.Add(feeCap, tip)

	auth.GasTipCap = tip
	auth.GasFeeCap = feeCap
	Logger.Debugf("set 1559 fees: tip %d, feeCap %d (basefee %d)", tip, feeCap, header.BaseFee)

	return auth, nil
}

func GetTransactionReceipt(endPoint string, hash common.Hash) (*types.Receipt, error) {
	client, err := ethclient.Dial(endPoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.TransactionReceipt(ctx, hash)
}

func GetTransactionRetry(endpoint string, h common.Hash) (*types.Transaction, error) {
	retry := 0
	for retry < 10 {
		tx, err := GetTransaction(endpoint, h)
		if err == nil {
			return tx, nil
		}
		retry++
		time.Sleep(time.Duration(retry) * time.Second)
	}
	return nil, fmt.Errorf("fail to get tx")
}

func GetTransaction(endpoint string, h common.Hash) (*types.Transaction, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, _, err := client.TransactionByHash(ctx, h)
	return res, err
}

func checkTx(endPoint string, txHash common.Hash) error {
	Logger.Debug("check tx: ", txHash.String())
	var receipt *types.Receipt
	var err error

	t := 0
	for i := 0; i < 10; i++ {
		t = 2*t + 1
		time.Sleep(time.Duration(t) * time.Second)
		receipt, err = GetTransactionReceipt(endPoint, txHash)
		if err == nil {
			break
		}
	}

	if receipt == nil {
		return fmt.Errorf("%s not packaged", txHash)
	}

	if receipt.Status == types.ReceiptStatusFailed { // 0 means fail
		for _, elog := range receipt.Logs {
			log.Printf("Log: %v\n", elog) // 打印日志信息
		}
		err = AnalyzeTransactionFailure(endPoint, txHash)
		if err != nil {
			Logger.Warn("tx revert: ", err)
			return err
		}

		if receipt.GasUsed != receipt.CumulativeGasUsed {
			return fmt.Errorf("%s transaction exceed gas limit", txHash)
		}
		return fmt.Errorf("%s transaction mined but execution failed, check your input", txHash)
	}
	Logger.Debugf("%s cost gas: %d, price: %d, blob gas: %d, price: %d", txHash.String(), receipt.GasUsed, receipt.EffectiveGasPrice, receipt.BlobGasUsed, receipt.BlobGasPrice)
	return nil
}

func AnalyzeTransactionFailure(endPoint string, txHash common.Hash) error {
	client, err := ethclient.Dial(endPoint)
	if err != nil {
		return err
	}
	defer client.Close()
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return fmt.Errorf("failed to get transaction by hash: %v", err)
	}

	if isPending {
		return fmt.Errorf("transaction is still pending")
	}

	// 获取交易回执
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return fmt.Errorf("failed to get transaction receipt: %v", err)
	}

	// 获取失败的合约调用信息
	callMsg := ethereum.CallMsg{
		From:     getFrom(tx),
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	_, err = client.CallContract(context.Background(), callMsg, receipt.BlockNumber)
	return err
}

func getFrom(tx *types.Transaction) common.Address {
	getSigner := func(trans *types.Transaction) types.Signer {
		v, _, _ := trans.RawSignatureValues()
		var isProtectedV bool
		for loop := true; loop; loop = false {
			if v.BitLen() <= 8 {
				vv := v.Uint64()
				isProtectedV = vv != 27 && vv != 28
				break
			}
			isProtectedV = true
		}
		if v.Sign() != 0 && isProtectedV {
			var chainId *big.Int
			for loop := true; loop; loop = false {
				if v.BitLen() <= 64 {
					vv := v.Uint64()
					if vv == 27 || vv == 28 {
						chainId = new(big.Int)
						break
					}
					chainId = new(big.Int).SetUint64((vv - 35) / 2)
					break
				}
				nv := new(big.Int).Sub(v, big.NewInt(35))
				chainId = nv.Div(nv, big.NewInt(2))
			}
			return types.NewEIP155Signer(chainId)
		} else {
			return types.HomesteadSigner{}
		}
	}
	signer := getSigner(tx)
	from, err := types.Sender(signer, tx)
	if err != nil {
		return common.Address{}
	}
	return from
}

func Transfer(ep string, sk *ecdsa.PrivateKey, toAddr common.Address, value *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, ep)
	if err != nil {
		return fmt.Errorf("dail %s fail %s", ep, err)
	}
	defer client.Close()

	fromAddr := utils.ECDSAToAddr(sk)
	Logger.Debugf("%s from has: %d", fromAddr, BalanceOf(ep, fromAddr))
	Logger.Debugf("%s to has: %d", toAddr, BalanceOf(ep, toAddr))

	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return err
	}

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	gasLimit := uint64(23000)
	gasPrice := header.BaseFee.Add(header.BaseFee, big.NewInt(int64(DefaultGasPrice)))

	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), sk)
	if err != nil {
		return err
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return err
	}

	err = checkTx(ep, signedTx.Hash())
	if err != nil {
		return err
	}
	Logger.Debugf("%s to has: %d", toAddr, BalanceOf(ep, toAddr))
	return nil
}

func BalanceOf(ep string, addr common.Address) *big.Int {
	client, err := rpc.Dial(ep)
	if err != nil {
		return big.NewInt(0)
	}
	defer client.Close()

	ctx, cancle := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancle()

	var result string
	err = client.CallContext(ctx, &result, "eth_getBalance", addr.String(), "latest")
	if err != nil {
		return big.NewInt(0)
	}

	val, _ := new(big.Int).SetString(result[2:], 16)
	return val
}

func TransferToken(ep string, chainID *big.Int, sk *ecdsa.PrivateKey, tokenAddr, toaddr common.Address, val *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	client, err := ethclient.DialContext(ctx, ep)
	if err != nil {
		return err
	}
	defer client.Close()
	ti, err := token.NewToken(tokenAddr, client)
	if err != nil {
		return err
	}
	au, err := MakeAuthBySk(ep, chainID, sk)
	if err != nil {
		return err
	}
	hasval, err := ti.BalanceOf(&bind.CallOpts{From: Base}, au.From)
	if err != nil {
		return err
	}
	Logger.Debugf("%s from has token: %d", au.From, hasval)

	hasval, err = ti.BalanceOf(&bind.CallOpts{From: Base}, toaddr)
	if err != nil {
		return err
	}
	Logger.Debugf("%s to has token: %d", toaddr, hasval)

	tx, err := ti.Transfer(au, toaddr, val)
	if err != nil {
		return err
	}
	err = checkTx(ep, tx.Hash())
	if err != nil {
		return err
	}

	hasval, err = ti.BalanceOf(&bind.CallOpts{From: Base}, toaddr)
	if err != nil {
		return err
	}
	Logger.Debugf("%s to has token: %d", toaddr, hasval)
	return nil
}

func BalanceOfToken(ep string, tokenaddr, addr common.Address) *big.Int {
	ctx, cancle := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancle()

	client, err := ethclient.DialContext(ctx, ep)
	if err != nil {
		return big.NewInt(0)
	}
	defer client.Close()

	ti, err := token.NewToken(tokenaddr, client)
	if err != nil {
		return big.NewInt(0)
	}

	hasval, err := ti.BalanceOf(&bind.CallOpts{From: addr}, addr)
	if err != nil {
		return big.NewInt(0)
	}
	Logger.Debugf("%s to has token: %d", addr, hasval)
	return hasval
}
