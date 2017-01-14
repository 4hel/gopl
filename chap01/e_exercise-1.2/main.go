// Echo2 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i+1, " ", arg)
	}
}
