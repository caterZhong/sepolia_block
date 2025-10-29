package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Infura Sepolia 测试网络 URL
	// 替换 YOUR-INFURA-API-KEY 为你的实际 API Key
	infuraURL := "https://sepolia.infura.io/v3/你的key"

	// 连接到 Sepolia 测试网络
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	fmt.Println("Successfully connected to Sepolia test network!")

	// 创建上下文 with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 查询最新区块号
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to get latest block header: %v", err)
	}

	fmt.Printf("Latest block number: %d\n", header.Number)

	// 查询指定区块（这里查询最新区块和前一个区块）
	queryBlocks(ctx, client, header.Number)
	//if header.Number.Cmp(big.NewInt(0)) > 0 {
	//	queryBlocks(ctx, client, new(big.Int).Sub(header.Number, big.NewInt(1)))
	//}

}

func queryBlocks(ctx context.Context, client *ethclient.Client, blockNumber *big.Int) {
	separator := strings.Repeat("=", 50)
	fmt.Printf("\n%s\n", separator)
	fmt.Printf("Querying block: %d\n", blockNumber)
	fmt.Println(separator)

	// 获取区块详细信息
	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		log.Printf("Failed to get block %d: %v", blockNumber, err)
		return
	}

	// 输出区块基本信息
	fmt.Printf("Block Hash: %s\n", block.Hash().Hex())
	fmt.Printf("Block Number: %d\n", block.Number())
	fmt.Printf("Block Timestamp: %s (%d)\n",
		time.Unix(int64(block.Time()), 0).Format("2006-01-02 15:04:05"),
		block.Time())
	fmt.Printf("Transaction Count: %d\n", len(block.Transactions()))
	fmt.Printf("Parent Hash: %s\n", block.ParentHash().Hex())
	fmt.Printf("Difficulty: %d\n", block.Difficulty())
	fmt.Printf("Gas Used: %d\n", block.GasUsed())
	fmt.Printf("Gas Limit: %d\n", block.GasLimit())
	fmt.Printf("Nonce: %d\n", block.Nonce())

	// 矿工地址 - 直接从区块中获取
	if block.Coinbase() != (common.Address{}) {
		fmt.Printf("Miner: %s\n", block.Coinbase().Hex())
	} else {
		fmt.Printf("Miner: Unknown\n")
	}

	// 如果有交易，显示前几笔交易的哈希
	if len(block.Transactions()) > 0 {
		fmt.Printf("\nFirst %d transactions:\n", min(3, len(block.Transactions())))
		for i := 0; i < min(3, len(block.Transactions())); i++ {
			fmt.Printf("  TX %d: %s\n", i+1, block.Transactions()[i].Hash().Hex())
		}
	}

	// 区块大小估算
	fmt.Printf("Block Size (approx): %d bytes\n", block.Size())

	// 添加额外的区块信息
	fmt.Printf("Uncle Count: %d\n", len(block.Uncles()))
	if len(block.Uncles()) > 0 {
		fmt.Printf("Uncle Hashes:\n")
		for i, uncle := range block.Uncles() {
			fmt.Printf("  Uncle %d: %s\n", i+1, uncle.Hash().Hex())
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
