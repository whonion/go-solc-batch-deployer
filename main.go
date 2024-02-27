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

// Struct to store EVM chain information
type evmChain struct {
	ChainID     int64
	ChainName   string
	ExplorerURL string
	RPCURL      string
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load your private keys
	privateKeysStr, ok := os.LookupEnv("PRIVATE_KEY")
	if !ok {
		log.Fatal("PRIVATE_KEY environment variable not set")
	}
	privateKeys := strings.Split(privateKeysStr, "\n")

	// Print the number of private keys for testing
	fmt.Printf("Loaded wallets: %d\n", len(privateKeys))

	// Load your chainIDs as a comma-separated string
	chainIDsStr, ok := os.LookupEnv("CHAIN_ID")
	if !ok {
		log.Fatal("CHAIN_ID environment variable not set")
	}
	chainIDs := strings.Split(chainIDsStr, ",")

	// Convert chain IDs to integers
	var chains []int64
	for _, id := range chainIDs {
		chainID, err := strconv.ParseInt(strings.TrimSpace(id), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		chains = append(chains, chainID)
	}

	// Define EVM chains with corresponding explorer and RPC URLs
	evmChains := map[int64]evmChain{
		1: {
			ChainID:     1,
			ChainName:   "Ethereum",
			ExplorerURL: "https://etherscan.io",
			RPCURL:      "https://eth.llamarpc.com",
		},
		10: {
			ChainID:     10,
			ChainName:   "Optimism",
			ExplorerURL: "https://optimistic.etherscan.io",
			RPCURL:      "https://rpc.ankr.com/optimism",
		},
		56: {
			ChainID:     56,
			ChainName:   "BSC",
			ExplorerURL: "https://bscscan.com",
			RPCURL:      "https://1rpc.io/bnb",
		},
		42161: {
			ChainID:     42161,
			ChainName:   "Arbitrum",
			ExplorerURL: "https://arbiscan.io",
			RPCURL:      "https://rpc.ankr.com/arbitrum",
		},
		42170: {
			ChainID:     42170,
			ChainName:   "Arbitrum Nova",
			ExplorerURL: "https://nova.arbitrum.io",
			RPCURL:      "https://nova.arbitrum.io/rpc",
		},
		137: {
			ChainID:     137,
			ChainName:   "Polygon",
			ExplorerURL: "https://polygonscan.com",
			RPCURL:      "https://polygon.llamarpc.com",
		},
		100: {
			ChainID:     100,
			ChainName:   "Gnosis",
			ExplorerURL: "https://gnosisscan.io",
			RPCURL:      "https://gnosis-rpc.com",
		},
		324: {
			ChainID:     324,
			ChainName:   "zkSync Era",
			ExplorerURL: "https://explorer.zksync.io",
			RPCURL:      "https://mainnet.era.zksync.io",
		},
		1101: {
			ChainID:     1101,
			ChainName:   "zkEVM",
			ExplorerURL: "https://zkevm.polygonscan.com",
			RPCURL:      "https://1rpc.io/zkevm",
		},
		43114: {
			ChainID:     43114,
			ChainName:   "Avalanche",
			ExplorerURL: "https://snowtrace.io",
			RPCURL:      "https://rpc.ankr.com/avalanche",
		},
		1313161554: {
			ChainID:     1313161554,
			ChainName:   "Aurora",
			ExplorerURL: "https://mainnet.aurora.dev",
			RPCURL:      "https://1rpc.io/aurora",
		},
		250: {
			ChainID:     250,
			ChainName:   "Fantom",
			ExplorerURL: "https://ftmscan.com",
			RPCURL:      "https://1rpc.io/ftm",
		},
		5: {
			ChainID:     5,
			ChainName:   "Goerly Testnet",
			ExplorerURL: "https://goerli.etherscan.io",
			RPCURL:      "https://rpc.ankr.com/eth_goerli",
		},
		11155111: {
			ChainID:     11155111,
			ChainName:   "Sepolia",
			ExplorerURL: "https://sepolia.etherscan.io",
			RPCURL:      "https://endpoints.omniatech.io/v1/eth/sepolia/public",
		},
		534353: {
			ChainID:     534353,
			ChainName:   "Scroll Alpha Testnet",
			ExplorerURL: "https://scrollscan.com",
			RPCURL:      "https://scroll-alpha-public.unifra.io",
		},
		534351: {
			ChainID:     534351,
			ChainName:   "Scroll Alpha Testnet",
			ExplorerURL: "https://sepolia-blockscout.scroll.io",
			RPCURL:      "https://1rpc.io/scroll/sepolia",
		},
		534352: {
			ChainID:     534352,
			ChainName:   "Scroll Mainnet",
			ExplorerURL: "https://scrollscan.com",
			RPCURL:      "https://rpc.scroll.io",
		},
		59140: {
			ChainID:     59140,
			ChainName:   "Linea(Testnet)",
			ExplorerURL: "https://goerli.lineascan.build",
			RPCURL:      "https://rpc.goerli.linea.build",
		},
		59144: {
			ChainID:     59144,
			ChainName:   "Linea Mainnet",
			ExplorerURL: "https://lineascan.build",
			RPCURL:      "https://rpc.linea.build",
		},
		6102: {
			ChainID:     6102,
			ChainName:   "Cascadia Testnet",
			ExplorerURL: "https://explorer.cascadia.foundation",
			RPCURL:      "https://testnet.cascadia.foundation",
		},
		34443: {
			ChainID:     34443,
			ChainName:   "Mode network",
			ExplorerURL: "https://explorer.mode.network",
			RPCURL:      "https://1rpc.io/mode",
		},
		7000: {
			ChainID:     7000,
			ChainName:   "ZetaChain Mainnet",
			ExplorerURL: "https://explorer.zetachain.com",
			RPCURL:      "https://zeta.rpcgrid.com",
		},
		8453: {
			ChainID:     8453,
			ChainName:   "Base Mainnet",
			ExplorerURL: "https://basescan.org",
			RPCURL:      "https://base-rpc.publicnode.com",
		},
		204: {
			ChainID:     204,
			ChainName:   "opBNB",
			ExplorerURL: "https://mainnet.opbnbscan.com",
			RPCURL:      "https://opbnb-rpc.publicnode.com",
		},
	}

	// Iterate over each private key
	for _, privateKeyStr := range privateKeys {
		privateKey, err := crypto.HexToECDSA(strings.TrimSpace(privateKeyStr))
		if err != nil {
			log.Fatalf("Error parsing private key: %s", err)
		}

		// Iterate over each chain
		for _, chainID := range chains {
			// Check if chainID exists in the evmChains map
			chain, ok := evmChains[chainID]
			if !ok {
				log.Fatalf("⛓️ Unsupported CHAIN_ID: %d", chainID)
			}

			// Connect to RPC Provider
			rpcProvider, ok := os.LookupEnv("RPC_PROVIDER")
			if !ok {
				fmt.Printf("🛑 RPC_PROVIDER environment variable not set\n")
				fmt.Printf("RPC will use the public node from 🌎 https://chainlist.org/\n")
				rpcProvider = chain.RPCURL
			}

			client, err := ethclient.Dial(rpcProvider)
			if err != nil {
				log.Fatal(err)
			}

			// Load and deploy each contract in succession
			contractFiles, err := os.ReadDir("./contracts")
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range contractFiles {
				{
					if filepath.Ext(file.Name()) == ".sol" {
						// Compile contract
						contractPath := filepath.Join("contracts", file.Name())
						var cmd *exec.Cmd

						if runtime.GOOS == "windows" {
							cmd = exec.Command("cmd/solc.exe", "--allow-paths", "./node_modules/openzeppelin-solidity/", "--bin", "--abi", "--optimize", "--output-dir", "compiled_contracts", "--evm-version", "byzantium", "--overwrite", contractPath)
						} else {
							cmd = exec.Command("solc", "--allow-paths", "./node_modules/openzeppelin-solidity/", "--bin", "--abi", "--optimize", "--output-dir", "compiled_contracts", "--evm-version", "byzantium", "--overwrite", contractPath)
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
					binPath := filepath.Join(binDir, binFilename)
					abiPath := filepath.Join(binDir, abiFilename)

					// Get the bytecode and ABI from the compiled contract
					bytecodeBytes, err := os.ReadFile(binPath)
					if err != nil {
						log.Fatal(err)
					}
					bytecodeStr := string(bytecodeBytes)
					constructorBytes, err := hex.DecodeString(bytecodeStr[:len(bytecodeStr)-68])
					if err != nil {
						log.Fatal(err)
					}

					abiBytes, err := os.ReadFile(abiPath)
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
					fmt.Printf("💸  Estimated gas for deploy: %v\n", estimateGas)
					fmt.Printf("🛜   Contract %s will be deploy to: %s chain\n", file.Name(), chain.ChainName)
					// Create a new instance of a transaction signer
					auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
					if err != nil {
						log.Fatal(err)
					}
					gasLimit := estimateGas
					auth.GasPrice = gasPrice
					auth.GasLimit = gasLimit + uint64(50000)

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
					fmt.Printf("📃  Contract %s waiting to be mined: %s\n", file.Name(), chain.ExplorerURL+"/tx/"+tx.Hash().Hex())
					receipt, err := bind.WaitMined(context.Background(), client, tx)
					if err != nil {
						log.Fatal(err)
					}
					if receipt.Status != types.ReceiptStatusSuccessful {
						log.Fatalf("❌  contract %s deployment failed", file.Name())
					}

					// Print the contract address and the transaction hash
					fmt.Printf("🚀  Contract %s deployed to: %s\n", file.Name(), chain.ExplorerURL+"/address/"+address.Hex())
				}
			}
		}

	}
}
