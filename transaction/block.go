package transaction

import (
	"context"
	"fmt"
	"goethereum/initial"
	"log"
	"math/big"
)

func inquireBlock() {
	client := initial.InitClient()

	// 区块头
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String()) // block number

	// 完整区块
	blockNumber := big.NewInt(5671788)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())
	fmt.Println(block.GasLimit())
	fmt.Println(block.Time())
	fmt.Println(len(block.Transactions()))

	// 区块交易数目
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
