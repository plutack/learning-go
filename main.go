package main

import (
	"fmt"
	"time"

	"github.com/plutack/intro-to-go/account"
)

func main() {
	var (
		talutAccount, bayoAccount, sholaAccount *account.Account
		err                                     error
	)
	talutAccount, err = account.New("Talut", 1000, "savings")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		err = nil
	}
	fmt.Printf("talut account should be undefined: %v\n", talutAccount)
	time.Sleep(1 * time.Second)

	bayoAccount, err = account.New("Bayo", 2000, "savings")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		err = nil
	}
	fmt.Printf("valid account: %v\n", *bayoAccount)
	time.Sleep(1 * time.Second)
	account.New("Talut", 5999, "nil")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	time.Sleep(1 * time.Second)
	sholaAccount, err = account.New("Shola", 4000, "fixedDeposit")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		err = nil
	}
	fmt.Printf("valid account: %v\n", sholaAccount)
}
