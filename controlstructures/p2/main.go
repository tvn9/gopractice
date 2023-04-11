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

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("something wrong with the input, error:", err)
		os.Exit(-1)
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	userAge, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	} else if userAge <= 0 {
		fmt.Println("Please enter age greater than zero")
		return
	}

	if userAge >= 60 {
		fmt.Println("Wecome to the senior club!")
	} else if userAge < 60 && userAge >= 30 {
		fmt.Println("Best age to be anywhere and do anything!")
	} else if userAge < 30 && userAge >= 18 {
		fmt.Println("Best time to learn everything you can")
	} else {
		fmt.Println("Play, learn, and have alot of friends")
	}
}
