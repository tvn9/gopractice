// Variatics function example
package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	total := sum(2, 3, 5, 7, 9)
	fmt.Println("Sum of 2, 3, 5, 7, and 9 is", total)

	// Let make some slices
	odd := []int{3, 5, 7, 9}
	even := []int{2, 4, 6, 8}

	oddEven := append(odd, even...)

	oddSum := sum(odd...)
	fmt.Printf("Sum of odd %v = %d\n", odd, oddSum)

	allSum := sum(oddEven...)
	fmt.Printf("Sum of oddEven %v = %d\n", oddEven, allSum)
}
