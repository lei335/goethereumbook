package swarmStorage

import (
	"fmt"
	"io/ioutil"
	"log"

	bzzclient "github.com/ethersphere/swarm/api/client"
)

func downloadFromSwarm() {
	manifestHash := "f9192507e2e8e118bfedac428c3aa1dec4ae156e954128ec5fb27f63ee67bcac"

	client := bzzclient.NewClient("http://127.0.0.1:8500")

	// 获取清单
	manifest, isEncrypted, err := client.DownloadManifest(manifestHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(isEncrypted)

	// 遍历清单条目
	for _, entry := range manifest.Entries {
		fmt.Println(entry.Hash)
		fmt.Println(entry.ContentType)
		fmt.Println(entry.Path)
	}

	// 下载文件
	file, err := client.Download(manifestHash, "")
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content)) // hello world
}
