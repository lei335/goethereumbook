package transaction

import (
	"context"
	"encoding/hex"
	"fmt"
	"goethereum/initial"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func SendRawTx() {
	client := initial.InitClient()

	rawTx := "f8a9018504d448bfe382548c9428b149020d2152179873ec60bed6bf7cd705775d80b844a9059cbb0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d00000000000000000000000000000000000000000000003635c9adc5dea0000025a074cfd4992ba0ee40020fe1f3fde07eff9bec7edf34362e19af549d53b435760ca00f2cecb2f3ad2adfbcb3402c42a61a2fb9457812b4815442d1b266b034c0add8"

	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		log.Fatal(err)
	}

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
