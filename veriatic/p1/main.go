// Veriatic function example
package main

import "fmt"

func sum(sums ...int) int {
	sum := 0
	if len(sums) == 0 {
		return sum
	}
	for _, n := range sums {
		sum += n
	}
	return sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := []int{10, 11, 12}
	sum1 := sum(a...)
	sum2 := sum(b...)
	all := append(a, b...)

	sum := sum(all...)
	fmt.Printf("Sum of %v = %d\n", a, sum1)
	fmt.Printf("Sum of %v = %d\n", b, sum2)
	fmt.Printf("All %d\n", all)
	fmt.Printf("Sum of all = %d\n", sum)

}
