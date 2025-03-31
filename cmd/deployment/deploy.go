package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	contract "github.com/Emircaan/reverse-proxy/ethereum/contracts/auth"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatalf("Failed to connect ethereum client")
	}
	privateKey, err := crypto.HexToECDSA("edab3d2985f5a1ca915ee60c8f0cccbd61ced925996f287ab2903fbc9368a4de")
	if err != nil {
		log.Fatal("failed to convert private key")
	}
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		log.Fatal("failed to create authorized transactor")
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	address, tx, _, err := contract.DeployContracts(auth, client)

	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	fmt.Println("Contract deployed at address:", address.Hex())
	fmt.Println("Transaction Hash:", tx.Hash().Hex())

}
