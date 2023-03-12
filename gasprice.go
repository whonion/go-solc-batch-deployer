package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to the Ethereum node
	client, err := ethclient.Dial(os.Getenv("RPC_PROVIDER"))
	if err != nil {
		log.Fatal(err)
	}

	// Get the suggested gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Print the gas price
	fmt.Printf("Current gas price for chain with RPC: %v: %v wei\n", os.Getenv("RPC_PROVIDER"), gasPrice)
}
