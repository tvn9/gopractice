// Goroutine example
package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capLetters []rune

	capIt := func(r rune) {
		capLetters = append(capLetters, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)
	}

	for _, l := range data {
		go capIt(l)
	}
	time.Sleep(110 * time.Millisecond)
	fmt.Printf("Capitalized: %c\n", capLetters)
}
