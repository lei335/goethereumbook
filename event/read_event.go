package event

import (
	"context"
	"fmt"
	"goethereum/contract/store"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func readEvent() {
	// 连接客户端
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2394201),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 解码经过ABI编码的日志，需要导入合约的ABI
	contractAbi, err := abi.JSON(strings.NewReader(store.StoreABI))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {

		itemlog, err := contractAbi.Unpack("ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		key := itemlog[0].([32]byte)
		value := itemlog[1].([32]byte)
		fmt.Println(string(key[:]))
		fmt.Println(string(value[:]))

		// 日志包含其他附加信息
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())

		// 读取事件日志的主题
		var topics [4]string
		for i, t := range vLog.Topics {
			topics[i] = t.Hex()
		}
		fmt.Println(topics[0]) // 事件的签名

		// topics[0] 等同于
		eventSignature := []byte("ItemSet(bytes32,bytes32)")
		hash := crypto.Keccak256Hash(eventSignature)
		fmt.Println(hash.Hex())
	}
}
