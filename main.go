package main

import (
	"fmt"

	"github.com/seungsu3579/learngo/banking"
)

func main() {
	account := banking.Account{Owner: "luca.ss", Balance: 1000}
	fmt.Println(account)
}
