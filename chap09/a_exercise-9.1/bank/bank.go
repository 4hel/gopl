// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type withdraw struct {
	amount     int
	resultChan chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdraw)

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func Withdraw(amount int) bool {
	result := make(chan bool)
	w := withdraw{amount, result}
	withdraws <- w
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case w := <-withdraws:
			if w.amount <= balance {
				balance -= w.amount
				w.resultChan <- true
			} else {
				w.resultChan <- false
			}

		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
