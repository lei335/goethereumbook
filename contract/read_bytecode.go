package contract

import (
	"context"
	"encoding/hex"
	"fmt"
	"goethereum/initial"
	"log"
)

func ReadContractByteCode() {
	client := initial.InitClient()

	contractAddress := DeployContract()

	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode))
}
