package cmdprompt

import "fmt"

const mainTitle = "\nBMI Calculator"
const separator = "================================="
const WeightPrompt = "Please enter your weight (kg): "
const HeightPrompt = "Please enter your height (m): "
const Result = "Your BMI is"

func PrintWelcome() {
	// Display program usage info to user
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
