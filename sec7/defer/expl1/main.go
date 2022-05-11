// Defer example code
package main

import "fmt"

func one() {
	fmt.Println("1")
}

func two() {
	fmt.Println("2")
}

func three() {
	fmt.Println("3")
}

func sample() {
	fmt.Println("Begin")

	defer three()
	defer two()
	defer one()

	fmt.Println("End")
}

func main() {
	sample()
}
