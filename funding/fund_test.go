package funding

import (
	"sync"
	"testing"
)

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
