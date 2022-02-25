package initial

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func InitClient() *ethclient.Client {
	client, err := ethclient.Dial("https://cloudflare-eth.com") // 连接到infura网关
	// or dial to local ipc
	//client, err = ethclient.Dial("/home/user/.ethereum/geth.ipc")
	// 连接到本地运行的ganache rpc主机
	client, err = ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
