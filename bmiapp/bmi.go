package main

import (
	"bmiapp/bmicalc"
	"bmiapp/cmdprompt"
)

func main() {

	// Print Welcome Info
	cmdprompt.PrintWelcome()

	// Get weight and height from user input
	w, h := getUserMetrics()

	// Print user input
	printInput(w, h)

	// Calculate BMI
	bmi := bmicalc.CalcBmi(w, h)

	// Print result to screen
	printBMI(bmi)

}
