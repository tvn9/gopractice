// Function Literals: 01, Anonymous Function | 02, Closures | 03, Aliases
// As Function Parameter
package main

import (
	"fmt"
	"strings"
)

func customMsg(fn func(m string), msg string) {
	msg = strings.ToUpper(msg)
	fn(msg)
}

func surround() func(msg string) {
	return func(msg string) {
		fmt.Printf("%.*s\n", len(msg), "--------------")
		fmt.Println(msg)
		fmt.Printf("%.*s\n", len(msg), "--------------")
	}
}

func main() {
	msg := "Hello"
	fn := surround()
	customMsg(fn, msg)
}
