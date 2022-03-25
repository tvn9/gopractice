// Finding the substring in a string
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"

func main() {
	lookFor := "lamb"

	// look for "lamb" in refString
	contain := strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	lookFor = "wolf"
	contain = strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	preFix := "Mary"
	begin := strings.HasPrefix(refString, preFix)
	fmt.Printf("The \"%s\" starts with \"%s\": %t \n", refString, preFix, begin)

	subFix := "lamb"
	ends := strings.HasSuffix(refString, subFix)
	fmt.Printf("The \"%s\" end with \"%s\": %t\n", refString, subFix, ends)

}
