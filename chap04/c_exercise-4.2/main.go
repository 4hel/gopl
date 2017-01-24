// calculate the sha256 sum of STDIN and print it to STDOUT

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var f384 = flag.Bool("sha384", false, "calculate 384 bit sum instead")
var f512 = flag.Bool("sha512", false, "calculate 512 bit sum instead")

func main() {
	checkFlags()
	// data of type []byte gets all from STDIN
	data, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Fprintf(os.Stderr, "sha256: %v\n", err)
	}
	if *f384 {
		fmt.Printf("%x\n", sha512.Sum384(data))
	} else if *f512 {
		fmt.Printf("%x\n", sha512.Sum512(data))
	} else {
		fmt.Printf("%x\n", sha256.Sum256(data))
	}
}

func checkFlags() {
	flag.Parse()
	if *f384 && *f512 {
		fmt.Fprintln(os.Stderr, "sha256: cannot do both 384 and 512, please decide")
		os.Exit(1)
	}
}
