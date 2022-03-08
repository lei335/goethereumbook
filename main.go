package main

/*
go项目目录结构通常如下：
- my-go-project
	- cmd
	- pkg or lib
	- internal
	- go.mod && go.sum
	- Makefile
*/

import (
	"fmt"
	"goethereum/transaction"
)

func main() {
	fmt.Println("Hello, go ethereum!")

	transaction.SendRawTx()
}
