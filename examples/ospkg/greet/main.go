package main

import (
	"fmt"
	"os"
)

func main() {

	// Read cli input
	args := os.Args[1:]

	fmt.Println(args, len(args))

	if len(args) != 2 {
		fmt.Println("Usage: please enter your first and last name")
		fmt.Println("Example: command [first last]")
		return
	}

	fmt.Printf("Hello %s %s!\n", args[0], args[1])
}
