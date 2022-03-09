package contract

import (
	"fmt"
	"log"
)

func queryContract() {
	instance := LoadContract()

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}
