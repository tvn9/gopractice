// Odd or Even
package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

func main() {
	// Get a number from cli
	args := os.Args[1:]

	// Handle cli input error
	if len(args) != 1 {
		fmt.Println("Pick a number")
		return
	}

	// Verify if number is even and divisible by 8
	num, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%q is not a number.\n", args[0])
		return
	}

	if num <= 0 {
		fmt.Println("Please enter a positive number larger than 0")
	} else if num%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8\n", num)
	} else if num%2 == 0 {
		fmt.Printf("%d is an even number.\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}
}
