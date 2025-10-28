package main

import (
	//"context"
	"fmt"
	"log"
	"math/big"
	"sepolia_block/contract/counter" // 生成的绑定代码

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, _ := ethclient.Dial("https://sepolia.infura.io/v3/你的key")

	// 选择其中一种方式：

	// 方式A: 部署新合约
	//deployAndUse(client)

	// 方式B: 使用已部署的合约
	useExistingContract(client, "0xc7c4566FF8508FA96881a6b6f3B6f919a0846319")
}

// 方式A: 部署新合约
func deployAndUse(client *ethclient.Client) {
	// 1. 准备部署参数
	privateKey, _ := crypto.HexToECDSA("your-private-key")
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))

	// 2. 部署合约
	address, tx, instance, err := counter.DeployCounter(auth, client, big.NewInt(0))
	if err != nil {
		log.Fatal("部署失败:", err)
	}

	fmt.Printf("合约部署地址: %s\n", address.Hex())
	fmt.Printf("部署交易: %s\n", tx.Hash().Hex())

	// 3. 使用新部署的合约
	useContract(instance, address.Hex())
}

// 方式B: 使用已部署的合约
func useExistingContract(client *ethclient.Client, contractAddress string) {
	address := common.HexToAddress(contractAddress)
	instance, err := counter.NewCounter(address, client)
	if err != nil {
		log.Fatal("连接合约失败:", err)
	}

	fmt.Printf("连接到现有合约: %s\n", contractAddress)
	useContract(instance, contractAddress)
}

// 使用合约的通用函数
func useContract(instance *counter.Counter, address string) {
	// 调用合约方法
	count, err := instance.GetCount(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("合约 %s 的当前计数: %d\n", address, count)

	// 可以继续调用其他方法...
	// instance.Increment(auth)
	// instance.Reset(auth)
}
