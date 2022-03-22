package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// reading input from cli
	args := os.Args[1:]
	// check for input errors
	if len(args) <= 0 {
		fmt.Println("Please enter a list of numbers, example: 2 3 4 5 6 54 332...")
		return
	}

	// declare a num slice
	var num []int

	// converts input string to numbers
	for _, n := range args {
		temp, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s is not a number\n", n)
			continue
		}

		// append input number from temp variable to num slice
		num = append(num, temp)
	}

	// sort the num slice from low to high
	sort.Ints(num)

	// print out the num slice to terminal
	fmt.Printf("Your inputs: %d\n", num)

}
