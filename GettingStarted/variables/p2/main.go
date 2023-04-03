package main

import (
	"fmt"
)

func main() {
	// full declaration of variable
	var hello string = "Hello"

	// full declaration with empty value
	var world string

	// now assign value the variable
	world = "World"

	// multiple declaration of the same type
	var firstName, lastName string = "Thanh", "Nguyen"
	fullname := firstName + " " + lastName

	// Now let print the above variable to the screen
	fmt.Println(hello, world)
	fmt.Println(firstName, lastName)

	// Now let try some shorthand declaration
	job := ""
	job = "Programmer"

	workAt := "Startup Software Company"

	yearBorn := 1975

	// exmaple of string concatination
	fmt.Println(firstName+" "+lastName+" is a", job, "he works at", workAt)
	fmt.Println(fullname, "was born in", yearBorn)
}
