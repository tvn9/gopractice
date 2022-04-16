// Simplify the Lear Year from previous exercise
package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Simplify the Leap Year
//
//  1. Look at the solution of "the previous exercise".
//
//  2. And simplify the code (especially the if statements!).
//
// EXPECTED OUTPUT
//  It's the same as the previous exercise.
// ---------------------------------------------------------

func main() {

	// Get the year number from CLI
	args := os.Args[1:]

	// Validate input
	if len(args) != 1 {
		fmt.Println("Usage: command [2019]")
		return
	}

	// convert year number string to number
	year, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%q is not a valid year\n", args[0])
		return
	}

	// Validate lear Year
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("Year %d is a leaf year.\n", year)
	} else {
		fmt.Printf("Year %d is not a leaf year.\n", year)
	}

}
