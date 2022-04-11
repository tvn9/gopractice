package main

import (
	"fmt"
	"strings"
)

const str = "Mary_had a little_lamb"

func splitUnderScore() {
	words := strings.Split(str, "_")
	for i, w := range words {
		fmt.Printf("Word %d is: %s\n", i, w)
	}
}
