// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	timeout = 10 * time.Second
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer (*wg).Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func scan(c net.Conn, out chan<- string) {

	input := bufio.NewScanner(c)
	for input.Scan() {
		out <- input.Text()
	}
	close(out)
}

//!+
func handleConn(c net.Conn) {
	var wg sync.WaitGroup // number echo workers
	wordChan := make(chan string)
	timeChan := time.After(timeout)

	go scan(c, wordChan)

	for done := false; !done; {
		select {
		case word := <-wordChan:
			wg.Add(1)
			go echo(c, word, 1*time.Second, &wg)
			timeChan = time.After(timeout)
		case <-timeChan:
			done = true
		}
	}
	wg.Wait()
	c.(*net.TCPConn).CloseWrite()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
