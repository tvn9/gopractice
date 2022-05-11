// Function Literals Example
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func compute(a, b int, op func(a, b int) int) int {
	fmt.Printf("Compute %d & %d\n", a, b)
	return op(a, b)
}

func main() {
	a, b := 10, 8
	r1 := compute(a, b, add)
	fmt.Printf("Result: %d\n", r1)

	fmt.Println("Compute", compute(a, b, func(a, b int) int {
		return a - b
	}))

	mul := func(a, b int) int {
		fmt.Printf("Multipling %d * %d = ", a, b)
		return a * b
	}
	fmt.Println(compute(a, b, mul))
}
