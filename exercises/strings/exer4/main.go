// Count the Chars
package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Count the Chars
//
//  1. Change the following program to work with unicode
//     characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func main() {
	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "İNANÇ", it should return 5 not 7.

	args := os.Args
	byteLenght := len(args[1])
	runeLenght := utf8.RuneCountInString(os.Args[1])
	fmt.Println("bytes count =", byteLenght)
	fmt.Println("runes count =", runeLenght)
}
