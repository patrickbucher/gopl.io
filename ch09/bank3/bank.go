// Package bank implements a concurrency-safe bank with only one account.
package bank

import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	defer mu.Unlock()
	mu.Lock()
	balance = balance + amount
}

func Balance() int {
	defer mu.Unlock()
	mu.Lock()
	return balance
}
