// Function Literals Example
package main

import "fmt"

// add func adds two numbers type int and return result
func add(a, b int) int {
	return a + b
}

// compute func receive closure func and return result from add func
func compute(a, b int, op func(a, b int) int) int {
	fmt.Printf("Compute %d & %d\n", a, b)
	return op(a, b)
}

func main() {
	// difine a, b variables of type in
	a, b := 10, 8

	// call/pass a, b, and "add(a, b, int) int" to compute func
	r1 := compute(a, b, add)
	fmt.Printf("Result: %d\n", r1)

	// just different way to pass result of compute func to other function
	// this all call anonymus inline funtion
	fmt.Printf("%d - %d = %d\n", a, b, compute(a, b, func(a, b int) int { return a - b }))

	// define a closure func to multiply two numbers
	mul := func(a, b int) int { fmt.Printf("Multipling %d * %d = ", a, b); return a * b }

	// call compute func and pass, a, b and mul func
	fmt.Println(compute(a, b, mul))
}
