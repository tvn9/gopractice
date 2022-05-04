// Pointer example
package main

import "fmt"

func main() {

	num := 999

	numPrt := &num

	// Print out the memory address
	fmt.Printf("num address = %p, num value = %d\n", &num, num)
	fmt.Printf("numPrt address %p, numPrt value = %d\n", numPrt, *numPrt)

	// Change the num value using Pointer
	*numPrt = 555
	fmt.Printf("num value is now = %d\n", num)

	// Change numPort value
	// numPrt = 888 // This won't work because numPrt type is a pointer to int *int

	// assign num to num2
	num2 := num
	fmt.Printf("num value = %d\n", num)
	fmt.Printf("num2 value = %d\n", num2)

	// Let change num2 veriable
	num2 = 100
	fmt.Printf("num value = %d\n", num)
	fmt.Printf("num2 value = %d\n", num2)

}
