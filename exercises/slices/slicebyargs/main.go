package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// uncomment the slice below
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw", "Queen", "Titanic"}

	// 1. Print the []string above
	fmt.Printf("Ships name: %q\n", ships)

	// 2. Get the start and stop position from the cli args
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf("Please enter start and stop value 1 - %d, example: 2 4\n", len(ships))
		return
	}

	// convert cli arguments to number
	start, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%s is not a number\n", args[0])
		return
	}

	stop, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("%s is not a number\n", args[1])
		return
	}

	fmt.Println(start, stop)

	// check start and stop numbers syntax,
	// start number must be smaller than stop number for the slice to work
	if start > stop {
		fmt.Println("first number must be lower than second number, please try again!")
		return
	}

	if start < 1 || stop > len(ships) {
		fmt.Println("Invalid start or stop value.")
		return
	}

	start -= 1

	// Perform the slice and print result
	newSlice := ships[start:stop]
	fmt.Printf("You selection: %q\n", newSlice)

}
