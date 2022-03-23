package main

import (
	"fmt"
	"os"
)

func main() {

	// Using os.Args package for getting CLI arguments
	args := os.Args

	// Print all command arguments
	fmt.Println(args)

	// Print first cli args[0]
	programName := args[0]
	fmt.Printf("The binary name is: %s\n", programName)

	// The arguments after the binary
	otherArgs := args[1:]
	fmt.Printf("other arguments: %s\n", otherArgs)

	for i, a := range otherArgs {
		fmt.Printf("#%d %s\n", i, a)
	}
}
