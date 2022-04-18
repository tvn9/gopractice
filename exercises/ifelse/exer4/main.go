// Arg count
package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	// Get arguments from CLI
	args := os.Args[1:]

	// Varlidate at least 1 arguments enterred from CLI
	if len(args) < 1 {
		fmt.Println("Give me args")
		return
	}

	// Count arguements
	argCount := len(args)

	// variable for is or are
	var isorare, argsorarg string

	if argCount == 1 {
		isorare = "is"
		argsorarg = "argument"
	} else {
		isorare = "are"
		argsorarg = "arguments"
	}

	fmt.Printf("There %s %d %s ", isorare, argCount, argsorarg)
	for _, a := range args {
		fmt.Printf("%q ", a)
	}
	fmt.Println()
}
