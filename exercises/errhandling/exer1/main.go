// Movie Ratings
package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

func main() {

	// Get age from command-line
	args := os.Args[1:]

	// check cli-input for error
	if len(args) != 1 {
		fmt.Println("Usage: command [13]")
		return
	}

	// Convert cli-args string to integer
	age, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%s is not a number.\n", os.Args[1])
		return
	}

	// Check for age and issue allow movie Ratings
	if age > 5 && age < 13 {
		fmt.Println("Rated-PG")
	} else if age > 12 && age < 18 {
		fmt.Println("Rated PG-13")
	} else if age > 17 {
		fmt.Println("Rated R")
	} else {
		fmt.Println("Wrong age.")
	}
}
