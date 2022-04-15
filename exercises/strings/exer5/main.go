// Improved Banger
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Improved Banger
//
//  Change the Banger program the work with Unicode
//  characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  İNANÇ!!!!!
// ---------------------------------------------------------

func main() {
	msg := os.Args[1]

	// s := msg + strings.Repeat("!", len(msg))
	s := msg + strings.Repeat("!", utf8.RuneCountInString(msg))

	fmt.Println(s)
}
