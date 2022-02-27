package transaction

import (
	"context"
	"fmt"
	"goethereum/initial"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func inquireTransaction() {
	client := initial.InitClient()
	// 完整区块
	blockNumber := big.NewInt(5671788)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
	}

	// 获取交易的发送方from
	tx := block.Transactions()[0]
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil); err == nil {
		fmt.Println(msg.From().Hex())
	}

	// 获取交易收据
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status) // 1:success; 0:fail
	fmt.Println(receipt.Logs)

	// 不获取块的情况下查询交易
	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for i := uint(0); i < count; i++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, i)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex())
	}

	// transactionByHash
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	fmt.Println(isPending)
}
