package main

import (
	"fmt"
	"sync"

	bank "gopl.io/ch09/bank3"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println(bank.Balance())

	// Alice:
	wg.Add(1)
	go func() {
		defer wg.Done()
		bank.Deposit(1000)
		bank.Deposit(100)
	}()

	// Bob:
	wg.Add(1)
	go func() {
		defer wg.Done()
		bank.Deposit(100)
		bank.Deposit(200)
	}()

	wg.Wait()
	fmt.Println(bank.Balance())
}
