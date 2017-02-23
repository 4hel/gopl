package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	zones := make(map[string]string)

	for _, arg := range os.Args[1:] {
		tokens := strings.Split(arg, "=")
		zones[tokens[0]] = tokens[1]
	}

	for k, v := range zones {
		time.Sleep(300 * time.Millisecond)
		go dial(k, v)
	}

	for {
	}
}

func dial(city string, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	copyTime(os.Stdout, conn, city)
}

func copyTime(dst io.Writer, src io.Reader, city string) {
	s := bufio.NewScanner(src)
	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(dst, "%10s: %s\n", city, s.Text())
	}
}
