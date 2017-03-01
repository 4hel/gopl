//found limit of ~ 2.5 mio goroutines with 8GB Memory
package main

import "fmt"
import "time"

const stages = 2500000

func main() {

	start := make(chan int)
	var end chan int
	lastOutput := start

	for i := 0; i < stages; i++ {
		newOutput := make(chan int)
		go pass(lastOutput, newOutput)
		lastOutput = newOutput
		if i == stages-1 {
			end = newOutput
		}
	}
	begin := time.Now()
	start <- 43
	fmt.Println("end:", <-end)
	fmt.Printf("took: %s\n", time.Since(begin))
}

func pass(input chan int, output chan int) {
	val := <-input
	//fmt.Print(">>> " + strconv.Itoa(val) + " >>> ")
	output <- val
}
