package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return int(*c), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		*c++
	}
	return int(*c), nil
}

func main() {
	var c LineCounter
	fmt.Fprint(&c, "asdf asdf\n asdf asdf")
	fmt.Println(c)
}
