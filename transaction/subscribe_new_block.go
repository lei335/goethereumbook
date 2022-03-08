package transaction

import (
	"context"
	"fmt"
	"goethereum/initial"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

func subscribeNewBlock() {
	client := initial.InitClient()

	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())

			// 获得区块完整信息
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(block.Coinbase().Hex())
		}
	}
}
