// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	"./bank"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(100)
		if bank.Withdraw(400) {
			fmt.Println("alice can withdraw 400")
		} else {
			fmt.Println("alice cannot withdraw 400")
		}
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		if bank.Withdraw(100) {
			fmt.Println("bob can withdraw 100")
		} else {
			fmt.Println("bob cannot withdraw 100")
		}
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
