package contract

import (
	"fmt"
	"math"
	"math/big"
)

func QueryERC20() {
	// 格式和query_contract.go中定义的一样，这里只写余额转换的方式

	bal := big.NewInt(100000002030000000)
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(18)))

	fmt.Printf("balance: %fETH\n", value)
}
