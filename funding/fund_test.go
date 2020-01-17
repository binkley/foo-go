package funding

import (
	"fmt"
	"sync"
	"testing"
)

type FundServer struct {
	commands chan interface{}
	fund     *Fund
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}

	go server.loop()

	return server
}

func (s *FundServer) Balance() int {
	responseChan := make(chan int)
	s.commands <- BalanceCommand{Response: responseChan}
	return <-responseChan
}

func (s *FundServer) Withdraw(amount int) {
	s.commands <- WithdrawCommand{Amount: amount}
}

func (s *FundServer) loop() {
	for command := range s.commands {
		switch command.(type) {
		case WithdrawCommand:
			withdrawal := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawal.Amount)

		case BalanceCommand:
			getBalance := command.(BalanceCommand)
			balance := s.fund.Balance()
			getBalance.Response <- balance

		default:
			panic(fmt.Sprintf("Unknown command: %v", command))
		}
	}
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}

const WORKERS = 10

func BenchmarkFund(b *testing.B) {
	// Skip N = 1
	if b.N < WORKERS {
		return
	}

	// Add as many dollars as we have iterations this run
	server := NewFundServer(b.N)

	dollarsPerFounder := b.N / WORKERS

	var wg sync.WaitGroup

	// Burn through them one at a time until they are all gone
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := 0; i < dollarsPerFounder; i++ {
				server.Withdraw(1)
			}
		}()
	}

	wg.Wait()

	balance := server.Balance()

	if balance != 0 {
		b.Error("Balance wasn't zero:", balance)
	}
}
