package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Testing bufio newscanner function - the input expected a file
	// command < file.txt
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanRunes)

	bc := 0
	for scanner.Scan() {
		bc++
	}

	fmt.Println("Number of byte:", bc)
}
