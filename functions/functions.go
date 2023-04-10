package main

import (
	"fmt"
)

func main() {

	a := 30
	b := 25
	fmt.Printf("%d + %d = %d\n", a, b, addTwoNumbers(a, b))

	fmt.Printf("%d - %d = %d\n", a, b, subtractTwoNumbers(a, b))

	fmt.Printf("%d * %d = %d\n", a, b, multiplyTwoNumbers(a, b))

	fmt.Printf("%.2f / %.2f = %.2f\n", 10.0, 3.0, devideTwoNumbers(10.0, 3.0))

	fmt.Printf("%d + %d + %d + %d + %d = %d\n", a, b, 4, 5, 6, addNumbers(a, b, 4, 5, 6))

}

func addTwoNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func subtractTwoNumbers(num1 int, num2 int) (difference int) {
	difference = num1 - num2
	return
}

func multiplyTwoNumbers(num1 int, num2 int) (product int) {
	product = num1 * num2
	return
}

func devideTwoNumbers(num1 float64, num2 float64) (quotient float64) {
	quotient = num1 / num2
	return
}

func addNumbers(num ...int) int {
	var sum int
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}
