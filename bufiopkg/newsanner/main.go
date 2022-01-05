package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Testing bufio newscanner function
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanRunes)

	wc := 0
	for scanner.Scan() {
		wc++
	}

	fmt.Println(wc)

}
