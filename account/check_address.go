package account

import (
	"context"
	"fmt"
	"goethereum/initial"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func checkAddress() {
	// 检查地址是否有效，使用正则表达式
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	// 检查地址是否为账户或智能合约
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := initial.InitClient().CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract)
}
