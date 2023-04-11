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
