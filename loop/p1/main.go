package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	// Get user input
	userInput, err := getUserChoice()
	if err != nil {
		fmt.Println(err)
		return
	}

	if userInput == "1" {
		fmt.Print("Please enter a number to calculate: ")
		num, err := getInputNumber()
		if err != nil {
			fmt.Println(err)
			return
		}
		calculateSum(num)
	} else if userInput == "2" {
		fmt.Print("Please enter a number to calculate: ")
		num, err := getInputNumber()
		if err != nil {
			fmt.Println(err)
			return
		}
		factorial(num)
	} else if userInput == "3" {
		addUserInputLoop()
	} else if userInput == "4" {
		fmt.Print("Please enter a list of numbers separated by commas [1,2,3, 4, 5]: ")
		nums, err := getNumList()
		if err != nil {
			fmt.Println(err)
			return
		}
		calcListOfNumbers(nums)

	} else {
		fmt.Println("Please try to enter option 1 to 4")
	}
}

func calculateSum(n int) {
	var sum int
	for i := 1; i <= n; i++ {
		if i == n {
			fmt.Print(i)
		} else {
			fmt.Printf("%d + ", i)
		}
		sum += i
	}
	fmt.Printf(" = %d\n", sum)
}

func factorial(n int) {
	product := 1
	for i := 1; i <= n; i++ {
		if i == n {
			fmt.Print(i)
		} else {
			fmt.Printf("%d * ", i)
		}
		product *= i
	}
	fmt.Printf(" = %d\n", product)
}

func addUserInputLoop() {
	userInput := true
	sum := 0

	for userInput {
		fmt.Print("Please enter a number to calculate: ")
		num, err := getInputNumber()
		if err != nil {
			fmt.Printf("something wrong with user input, Error: %s\n", err)
			userInput = false
			return
		}
		fmt.Printf("%d + %d = %d\n", sum, num, sum+num)
		sum += num
	}
}

func calcListOfNumbers(nums []int) {
	sum := 0
	for i, n := range nums {
		sum += n
		if i == len(nums)-1 {
			fmt.Print(n)
		} else {
			fmt.Printf("%d + ", n)
		}
	}
	fmt.Printf(" = %d\n", sum)
}

func getNumList() ([]int, error) {

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(-1)
	}

	userInput = strings.Replace(userInput, "\n", "", -1)
	userInput = strings.Replace(userInput, " ", "", -1)
	fmt.Println(userInput)
	numList := strings.Split(userInput, ",")

	var nums []int

	for _, n := range numList {
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s is not a number\n", n)
			return nil, errors.New("Please make sure input are number.")
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func getInputNumber() (int, error) {

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(-1)
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	inputNum, err := strconv.ParseInt(userInput, 10, 64)
	if err != nil {
		return 0, err
	}

	if inputNum > 0 {
		return int(inputNum), nil
	} else {
		return int(inputNum), errors.New("Error: input value must be larger than 0")
	}
}

func getUserChoice() (string, error) {

	fmt.Println("\nPlease Enter and option from 1 to 4")
	fmt.Println("1. Add up all the number of number X")
	fmt.Println("2. Build the factorial up to number X")
	fmt.Println("3. Sum up manually entered numbers")
	fmt.Println("4. Sum up a list of enter numbers")
	fmt.Print("Please select an option: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid Input")
		os.Exit(-1)
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	if userInput == "1" || userInput == "2" || userInput == "3" || userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("You have entered an invalid option")
	}
}
