// Days in the Month
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func main() {

	// Get the year from CLI
	args := os.Args[1:]

	// Validate cli-argument
	if len(args) != 1 {
		fmt.Println("usage: command [january]")
		return
	}

	// Convert input to lower case
	month := strings.ToLower(args[0])
	year := time.Now().Year()

	var feb string

	// validates year
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		feb = "29"
		fmt.Println(year, "is a leaf year")
	} else {
		feb = "28"
		fmt.Println(year, "is not a leaf year")
	}

	if month == "january" || month == "march" ||
		month == "may" || month == "july" ||
		month == "august" || month == "october" ||
		month == "december" {
		fmt.Printf("%s has 31 days.\n", strings.Title(month))
	} else if month == "april" || month == "june" ||
		month == "september" || month == "november" {
		fmt.Printf("%s has 30 days.\n", strings.Title(month))
	} else if month == "february" {
		fmt.Printf("%s has %s days.\n", strings.Title(month), feb)
	} else {
		fmt.Printf("%s is not a month.\n", month)
	}
}
