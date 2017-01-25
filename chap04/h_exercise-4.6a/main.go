// in place squash a run of unicode spaces into a single ascii space on []byte slice

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "\t\tstärt   and tab:\t\tend\nof line"
	fmt.Println(string(squash([]byte(s))))
}

func squash(bytes []byte) []byte {
	lastSpace := false
	readPointer := 0
	writePointer := 0

	for readPointer < (len(bytes)) {
		rune, size := utf8.DecodeRune(bytes[readPointer:])

		if !unicode.IsSpace(rune) {
			utf8.EncodeRune(bytes[writePointer:], rune)
			writePointer += size
			lastSpace = false
		} else {
			if !lastSpace {
				size = utf8.EncodeRune(bytes[writePointer:], int32(' '))
				writePointer += size
			}
			lastSpace = true
		}
		readPointer += size

	}

	return bytes[:writePointer]
}
