// Breaking the string into words
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"

func main() {

	// Convert refString to words slice using strings.Fields (using whitespace as seperator)
	words := strings.Fields(refString)
	for i, w := range words {
		fmt.Printf("Word %d is: %s\n", i, w)
	}

	// Call the SplitUnderScore from anyother.go
	splitUnderScore()
}
