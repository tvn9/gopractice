//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package main

import (
	"fmt"

	"github.com/tvn9/learngo/exercises/errors/exer1/timeparse"
)

func main() {
	test, err := timeparse.ParseTime("19:29:100")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test)
}
