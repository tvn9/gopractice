// Pointers Example
package main

import "fmt"

func main() {

	var numPrt *int
	fmt.Printf("numPrt = *int; mem addrress: %p\n", &numPrt)

	num := 100
	num2 := 500
	fmt.Printf("num = 100; mem addrss: %p\n", &num)
	fmt.Printf("num2 = 100; mem addrss: %p\n", &num2)

	numPrt = &num
	fmt.Printf("numPrt = &num; mem address: %p, num mem address: %p\n", &numPrt, &num)

	*numPrt = 200
	fmt.Printf("*numPrt = 200: %d, num: %d\n", *numPrt, num)

	numPrt = &num2
	fmt.Printf("numPrt: %p; num2: %d\n", &numPrt, &num2)

	*numPrt = 200
	fmt.Printf("num %d, num2 %d\n", num, num2)
}
