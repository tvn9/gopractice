// Function Literals: 01, Anonymous Function | 02, Closures | 03, Aliases
// Anonymous Function
package main

import "fmt"

func hellWorld() {
	fmt.Printf("Hello, ")
	world := func() {
		fmt.Printf("World!\n")
	}
	world()
	world()
	world()
	world()
}

func main() {
	hellWorld()
}
