package account

import (
	"context"
	"fmt"
	initial "goethereum/initial"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func address() {
	addr := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f") // 20Byte, length is 40 without prefix '0x'

	fmt.Println(addr.Hex())        //  0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(addr.Hash().Hex()) // 左边填充0到32Byte长度，0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
}

func balance() {
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	client := initial.InitClient()
	balance, err := client.BalanceAt(context.Background(), account, nil) // 区块指定为nil表示最新的余额
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// 精度转换
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	// 待处理账户余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)
}
