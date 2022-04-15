// Right Trim It
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Right Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     only the right-most part of the given string
//  3. Trim it from the right part only
//  4. Print the number of characters it contains.
//
// RESTRICTION
//  Your program should work with unicode string values.
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func main() {
	// currently it prints 17
	// it should print 5

	name := "inanç           "
	name = strings.TrimRight(name, " ")
	fmt.Println(name, utf8.RuneCountInString(name))
}
