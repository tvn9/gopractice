package main

import (
	"bmiapp/cmdprompt"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (float64, float64) {
	// Prompt user to input thier information
	weight := getUserInput(cmdprompt.WeightPrompt)
	height := getUserInput(cmdprompt.HeightPrompt)
	return weight, height
}

func getUserInput(promptText string) (value float64) {
	// remove the "\n" from the cli input
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	value, _ = strconv.ParseFloat(userInput, 64)
	return value
}
