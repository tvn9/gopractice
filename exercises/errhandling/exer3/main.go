// Leap Year
package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {
	// Get the year number from cli
	args := os.Args[1:]

	// verify for valid cli-arguments
	if len(args) != 1 {
		fmt.Println("Usage: command [2020]")
		return
	}

	// convert cli-input string to number
	year, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%q is a not a valid year.\n", args[0])
		return
	}

	/*
		// validate leaf year rules
		if year%100 == 0 && year%400 != 0 {
			fmt.Printf("Year %d is not a leaf year!\n", year)
		} else if year%100 == 0 && year%400 == 0 {
			fmt.Printf("Year %d is a leaf year!\n", year)
		} else if year%4 == 0 {
			fmt.Printf("Year %d is a leaf year!\n", year)
		} else {
			fmt.Printf("Year %d is not a leaf year!\n", year)
		}
	*/

	// Better validation
	if year%400 == 0 {
		fmt.Printf("Year %d is a leaf year!\n", year)
	} else if year%100 == 0 {
		fmt.Printf("Year %d is not a leaf year!\n", year)
	} else if year%4 == 0 {
		fmt.Printf("Year %d is a leaf year!\n", year)
	} else {
		fmt.Printf("Year %d is a not leaf year!\n", year)
	}
}
