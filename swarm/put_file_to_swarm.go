package swarmStorage

import (
	"fmt"
	"log"

	bzzclient "github.com/ethersphere/swarm/api/client"
)

func uploadToSwarm() {
	client := bzzclient.NewClient("http://127.0.0.1:8500")

	file, err := bzzclient.Open("./hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 上传文件，得到一个内容清单哈希值
	manifestHash, err := client.Upload(file, "", false, false, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(manifestHash) // 2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0

	// 之后就可以在这里查看上传的文件：
	// bzz://2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0
}
