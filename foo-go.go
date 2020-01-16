package main

import (
	"fmt"
	"github.com/binkley/foo-go/funding"
	"rsc.io/quote"
	"runtime"
)

func main() {
	fmt.Println(quote.Hello())

	fmt.Println()

	fund := funding.NewFund(3)
	fmt.Println("You started with", fund.Balance(), "zorkmids.")
	fund.Withdraw(1)
	fmt.Println("You now have", fund.Balance(), "zorkmids.")

	fmt.Println()

	fmt.Println("GOMAXPROCS is", runtime.GOMAXPROCS(0))
}
