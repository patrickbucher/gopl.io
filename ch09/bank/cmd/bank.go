package main

import (
	"fmt"

	"gopl.io/ch09/bank"
)

func main() {
	// Alice:
	go func() {
		bank.Deposit(200)                // A1
		fmt.Println("=", bank.Balance()) // A2
	}()
	// Bob:
	go bank.Deposit(100) // B
}
