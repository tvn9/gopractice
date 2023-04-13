package main

import "fmt"

func main() {

	// define a new slice
	friends := []string{"Mike", "Marko"}
	fmt.Println(friends, len(friends), cap(friends)) // len = 2, cap = 2

	// adding more elements to the slice
	friends = append(friends, "Tim")
	fmt.Println(friends, len(friends), cap(friends)) // len = 3, cap 4

	// adding some more elements
	friends = append(friends, "Ken", "Tuan")
	fmt.Println(friends, len(friends), cap(friends)) // len = 4, cap 8

}
