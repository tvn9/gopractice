package main

import (
	"bmiapp/cmdprompt"
	"fmt"
)

// Print user input data
func printInput(w, h float64) {
	fmt.Printf("your weight is %.2fkg\n", w)
	fmt.Printf("your height is %.2fm\n", h)
}

// Display info to user
func printBMI(b float64) {
	fmt.Printf("%s %.2f\n\n", cmdprompt.Result, b)
}
