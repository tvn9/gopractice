package main

import (
	"fmt"
	"os"
)

func main() {

	// Read user input
	args := os.Args

	fmt.Printf("First cli value: %s\n", args[0])
	fmt.Printf("Second cli value: %s\n", args[1])
	fmt.Printf("Third cli value: %s\n", args[2])
	fmt.Printf("Fourth cli value: %s\n", args[3])
}
