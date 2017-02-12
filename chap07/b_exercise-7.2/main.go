package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	w     io.Writer
	count int64
}

func (bc *ByteCounter) Write(p []byte) (int, error) {
	n, err := bc.w.Write(p)
	bc.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := ByteCounter{w, 0}
	return &bc, &bc.count
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "Hello World\n")
	fmt.Println(*count)
}
