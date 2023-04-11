package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter you age: ")

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	userAge, _ := strconv.ParseInt(userInput, 0, 64)

	// one wait to setup a bool veriable
	// isOldEnough := userAge >= 18

	if userAge >= 18 {
		fmt.Println("Welcome to the club!")
	} else {
		fmt.Println("Goodbye!")
	}
}
