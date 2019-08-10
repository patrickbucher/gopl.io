// Package bank implements a concurrency-safe bank with only one account.
package bank

import "log"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

type withdrawal struct {
	amount  int
	success bool
}

var withdrawals = make(chan withdrawal) // withdraw amount from deposit

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func Withdraw(amount int) bool {
	success := Balance() >= amount
	withdrawals <- withdrawal{amount, success}
	return success
}

func teller() {
	var balance int // calance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
			log.Println("deposit =", amount, "balance =", balance)
		case balances <- balance:
		case withdrawal := <-withdrawals:
			if withdrawal.success {
				balance -= withdrawal.amount
				log.Println("withdraw =", withdrawal.amount, "balance =", balance)
			} else {
				log.Println("withdraw =", withdrawal.amount, "from balance =", balance,
					"computer says 'No'")
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
