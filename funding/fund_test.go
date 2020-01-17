package funding

import (
	"sync"
	"testing"
)

type Transactor func(fund *Fund)

type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

type FundServer struct {
	commands chan TransactionCommand
	fund     *Fund
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan TransactionCommand),
		fund:     NewFund(initialBalance),
	}

	go server.loop()

	return server
}

func (s *FundServer) Transact(transactor Transactor) {
	command := TransactionCommand{
		Transactor: transactor,
		Done:       make(chan bool),
	}
	s.commands <- command
	<-command.Done
}

func (s *FundServer) Balance() int {
	var balance int

	s.Transact(func(f *Fund) {
		balance = f.Balance()
	})

	return balance
}

func (s *FundServer) Withdraw(amount int) {
	s.Transact(func(f *Fund) {
		f.Withdraw(amount)
	})
}

func (s *FundServer) loop() {
	for transaction := range s.commands {
		transaction.Transactor(s.fund)
		transaction.Done <- true
	}
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
