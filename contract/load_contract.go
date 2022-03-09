package contract

import (
	"goethereum/contract/store"
	"goethereum/initial"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func LoadContract() *store.Store {
	client := initial.InitClient()

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err := store.NewStore(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}
