// Reader and Writer Interface
package main

import (
	"bufio"
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

// Example code to demonstrade Reader and Writer

func main() {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Original string:", str)

	source := strings.NewReader(str)
	buffered := bufio.NewReader(source)

	newStr, err := buffered.ReadString('\n')
	if err == io.EOF {
		fmt.Println(newStr)
	} else {
		fmt.Println("something went wrong...")
	}
}
