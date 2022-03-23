package main

import "fmt"

type names []string

func main() {
	// Declare a slice
	newName := names{"Mike", "Chris", "Mark"}

	// call changeDate function
	changeData(newName)

	fmt.Println(newName)

}

func changeData(n names) {
	n[1] = "Tim"
	fmt.Println(n)
}
