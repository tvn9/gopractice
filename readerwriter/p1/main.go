// Reader and Writer Interface
package main

import (
	"fmt"
	"io"
	"strings"
)

/*
// Standard Lib interfaces
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
*/

// Example code to demonstrade Reader

func main() {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Original string:", str)

	reader := strings.NewReader(str)
	fmt.Printf("String from reader: %v\n%T\n", reader, reader)

	var newStr strings.Builder

	// Set the reading buffer to 5 bytes
	buffer := make([]byte, 5)

	for {
		numBytes, err := reader.Read(buffer)
		chunk := buffer[:numBytes]
		newStr.Write(chunk)

		fmt.Printf("Read %v bytes: %c\n", numBytes, chunk)
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("%v\n", newStr.String())
}
