package main

import (
	"fmt"
	"github.com/binkley/foo-go/funding"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())

	fund := funding.NewFund(3)
	fmt.Println("You started with", fund.Balance(), "zorkmids.")
	fund.Withdraw(1)
	fmt.Println("You now have", fund.Balance(), "zorkmids.")
}
