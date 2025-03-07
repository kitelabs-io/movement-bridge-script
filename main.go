package main

import (
	"bridge/contracts"
	"context"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var (
	adapterABI, _ = contracts.AdapterMetaData.GetAbi()
	erc20ABI, _   = contracts.Erc20MetaData.GetAbi()
)

var (
	adapterAddress common.Address

	client       *ethclient.Client
	transactor   *bind.TransactOpts
	bridgeAmount *big.Int
	moveAddress  [32]byte
)

func init() {
	godotenv.Load()
	adapterAddressEnv := os.Getenv("ADAPTER_ADDRESS")
	if adapterAddressEnv == "" {
		log.Panicln("ADAPTER_ADDRESS is required")
	}

	moveAddressEnv := os.Getenv("MOVE_ADDRESS")
	if moveAddressEnv == "" {
		log.Panicln("MOVE_ADDRESS is required")
	}
	moveAddressBytes := common.FromHex(moveAddressEnv)
	if len(moveAddressBytes) != 32 {
		log.Panicln("MOVE_ADDRESS should be 32 bytes, current length: ", len(moveAddressBytes))
	}

	copy(moveAddress[:], moveAddressBytes)
	log.Printf("Move address: %v\n", common.Bytes2Hex(moveAddress[:]))

	adapterAddress = common.HexToAddress(adapterAddressEnv)
	client = getClient()
	transactor = getTransactor()
	bridgeAmount = getBridgeAmount()

}

func main() {

	adapterContract, err := contracts.NewAdapter(adapterAddress, client)
	if err != nil {
		log.Panicf("Failed to create adapter contract: %v\n", err)
	}

	token, err := adapterContract.Token(nil)
	if err != nil {
		log.Panicf("Failed to get token address: %v\n", err)
	}

	log.Printf("Token address: %s\n", token.Hex())

	tokenContract, err := contracts.NewErc20(token, client)
	if err != nil {
		log.Panicf("Failed to create token contract: %v\n", err)
	}

	allowance, err := tokenContract.Allowance(nil, transactor.From, adapterAddress)
	if err != nil {
		log.Panicf("Failed to get allowance: %v\n", err)
	}
	if allowance.Cmp(bridgeAmount) < 0 {
		log.Printf("Approving %v to adapter\n", bridgeAmount)

		if tx, err := tokenContract.Approve(transactor, adapterAddress, bridgeAmount); err != nil {
			log.Panicf("failed to approve for adapter: %v\n", err)
		} else {
			log.Printf("approve tx submitted: %s, waiting for confirm\n", tx.Hash().Hex())
			waitForTransactionSuccess(tx.Hash())
			log.Printf("approve success: %s", tx.Hash().Hex())
		}
	} else {
		log.Printf("Already approved %v to adapter\n", allowance)
	}

	fee, err := adapterContract.QuoteSend(nil, contracts.SendParam{
		DstEid:      30325, // move eid
		To:          moveAddress,
		AmountLD:    bridgeAmount,
		MinAmountLD: bridgeAmount,
	}, false)

	if err != nil {
		log.Panicf("failed to estimate lz send fee: %v", err)
	}

	log.Printf("lz fee: %v\n", fee.NativeFee.String())

	transactor.Value = fee.NativeFee
	if tx, err := adapterContract.Send(transactor, contracts.SendParam{
		DstEid:      30325, // move eid
		To:          moveAddress,
		AmountLD:    bridgeAmount,
		MinAmountLD: bridgeAmount,
	}, fee, transactor.From); err != nil {
		log.Panicf("failed to bridge: %v", err)
	} else {
		log.Printf("bridge tx submitted: %s, waiting for confirm\n", tx.Hash().Hex())
		waitForTransactionSuccess(tx.Hash())
		log.Printf("bridge success: %s", tx.Hash().Hex())
	}
}

func waitForTransactionSuccess(txHash common.Hash) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			log.Printf("getting transaction receipt - error: %v, sleep for 1s...\n", err)
			time.Sleep(time.Second)
		}
		if receipt != nil {
			if receipt.Status == 1 {
				log.Printf("Transaction success: %s", txHash.Hex())
				break
			} else {
				log.Panicf("Transaction failed: %s", txHash.Hex())
			}
		}
	}
}

func getClient() *ethclient.Client {
	rpcUrl := os.Getenv("RPC_URL")
	if rpcUrl == "" {
		log.Panicln("RPC_URL env is required")
	}
	var err error
	client, err = ethclient.Dial(rpcUrl)
	if err != nil {
		log.Panicf("Failed to connect to the Ethereum client: %v", err)
	}
	return client
}

func getTransactor() *bind.TransactOpts {
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		log.Panicln("PRIVATE_KEY env is required")
	}

	pk, err := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
	if err != nil {
		log.Panicf("Failed to parse private key: %v", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Panicf("Failed to get chain id: %v", err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		log.Panicf("Failed to create transactor: %v", err)
	}

	return transactor
}

func getBridgeAmount() *big.Int {
	amountEnv := os.Getenv("BRIDGE_AMOUNT")
	if amountEnv == "" {
		log.Panicln("BRIDGE_AMOUNT env is required")
	}

	amount, _ := new(big.Int).SetString(amountEnv, 10)

	return amount
}
