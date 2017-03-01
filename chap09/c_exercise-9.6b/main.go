// found limit of ~ 4.5 mio messages per second
package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go pass(chan1, chan2)
	go pass(chan2, chan1)
	fmt.Println("ping pong for 1 second...")
	chan1 <- 1
	time.Sleep(1 * time.Second)
	select {
	case n := <-chan1:
		fmt.Println("msg sent:", n)
	case n := <-chan2:
		fmt.Println("msg sent:", n)
	}
	close(chan1)
	close(chan2)
}

func pass(input chan int, output chan int) {
	for val := range input {
		val++
		output <- val
	}
}
