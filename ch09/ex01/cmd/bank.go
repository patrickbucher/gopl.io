package main

import (
	"fmt"
	"sync"

	bank "gopl.io/ch09/ex01"
)

func main() {
	var wg sync.WaitGroup

	// Alice:
	go func() {
		defer wg.Done()
		wg.Add(1)
		bank.Withdraw(500)
		bank.Deposit(1000)
		bank.Deposit(100)
	}()

	// Bob:
	go func() {
		defer wg.Done()
		wg.Add(1)
		bank.Deposit(100)
		bank.Deposit(200)
		bank.Withdraw(1000)
		bank.Withdraw(1000)
	}()
	wg.Wait()
	fmt.Println(bank.Balance())
}
