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

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minutes, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	strSlice := strings.Split(input, ":")
	if len(strSlice) != 3 {
		return Time{}, &TimeParseError{"Invalid number of elements", input}
	} else {
		hour, err := strconv.Atoi(strSlice[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour %v", err), input}
		}
		minute, err := strconv.Atoi(strSlice[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minute %v", err), input}
		}
		second, err := strconv.Atoi(strSlice[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing second %v", err), input}
		}

		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{"hour out or range: 0 <= hour <= 23", fmt.Sprintf("%v", hour)}
		}
		if minute > 59 || minute < 0 {
			return Time{}, &TimeParseError{"minute out or range: 0 <= minute <= 60", fmt.Sprintf("%v", minute)}
		}
		if second > 59 || second < 0 {
			return Time{}, &TimeParseError{"second out or range: 0 <= second <= 60", fmt.Sprintf("%v", second)}
		}
		return Time{hour, minute, second}, nil
	}
}
