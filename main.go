package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the Ethereum network
	client, err := ethclient.Dial(os.Getenv("RPC_PROVIDER"))
	if err != nil {
		log.Fatal("RPC_PROVIDER environment variable not set")
	}

	// Load your private key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal("PRIVATE_KEY environment variable not set")
	}
	// Load your chainID
	chainIdStr, ok := os.LookupEnv("CHAIN_ID")
	if !ok {
		log.Fatal("CHAIN_ID environment variable not set")
	}
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	// Load and deploy each contract in succession
	contractFiles, err := os.ReadDir("./contracts")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range contractFiles {
		if filepath.Ext(file.Name()) == ".sol" {
			// Compile contract
			contractPath := filepath.Join("contracts", file.Name())
			var cmd *exec.Cmd

			if runtime.GOOS == "windows" {
				cmd = exec.Command("cmd/solc.exe", "--bin", "--abi", "--optimize", "--output-dir", "compiled_contracts", "--evm-version", "byzantium", "--overwrite", contractPath)
			} else {
				cmd = exec.Command("solc", "--bin", "--abi", "--optimize", "--output-dir", "compiled_contracts", "--evm-version", "byzantium", "--overwrite", contractPath)
			}

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}

		// Get the absolute path to the bin folder
		binDir, err := filepath.Abs("./compiled_contracts")
		if err != nil {
			log.Fatal(err)
		}

		// Rename output files
		name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		binFilename := fmt.Sprintf("%s.bin", name)
		abiFilename := fmt.Sprintf("%s.abi", name)
		// binPath := filepath.Join(binDir, fmt.Sprintf("contracts_%s_sol_%s.bin", name, name))
		// abiPath := filepath.Join(binDir, fmt.Sprintf("contracts_%s_sol_%s.abi", name, name))

		newBinPath := filepath.Join(binDir, binFilename)
		newAbiPath := filepath.Join(binDir, abiFilename)

		// err = os.Rename(binPath, newBinPath)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// err = os.Rename(abiPath, newAbiPath)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// Get the bytecode and ABI from the compiled contract
		bytecodeBytes, err := os.ReadFile(newBinPath)
		if err != nil {
			log.Fatal(err)
		}
		bytecodeStr := string(bytecodeBytes)
		constructorBytes, err := hex.DecodeString(bytecodeStr[:len(bytecodeStr)-68])
		if err != nil {
			log.Fatal(err)
		}

		abiBytes, err := os.ReadFile(newAbiPath)
		if err != nil {
			log.Fatal(err)
		}
		// Set the gas price and gas limit
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// Calculate the gas required for deploying the contract
		estimateGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			From: crypto.PubkeyToAddress(privateKey.PublicKey),
			To:   nil,
			Data: constructorBytes,
		})
		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			fmt.Printf("Estimate gas overflow uint64\n")
			log.Fatal(err)
		}
		fmt.Printf("Estimated gas for deploy: %v\n", estimateGas)
		// Create a new instance of a transaction signer
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
		if err != nil {
			log.Fatal(err)
		}
		gasLimit := estimateGas
		auth.GasPrice = gasPrice
		auth.GasLimit = gasLimit + uint64(10000)

		// Load the contract's ABI
		contractABI, err := abi.JSON(bytes.NewReader(abiBytes))
		if err != nil {
			log.Fatal(err)
		}

		// Deploy the contract
		address, tx, _, err := bind.DeployContract(auth, contractABI, constructorBytes, client)
		if err != nil {
			log.Fatal(err)
		}

		// Wait for the transaction to be mined
		fmt.Printf("Contract %s waiting to be mined: 0x%x\n", file.Name(), tx.Hash())
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			log.Fatal(err)
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			log.Fatalf("contract %s deployment failed", file.Name())
		}

		// Print the contract address and the transaction hash
		fmt.Printf("Contract %s deployed to: %s\n", file.Name(), address.Hex())
	}
}
