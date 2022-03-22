package main

import "fmt"

func main() {
	toppings := []string{"black olives", "green peppers"}
	var pizza []string

	toppings = append(toppings, "onion", "extra cheese")
	pizza = append(pizza, toppings...)

	fmt.Printf("pizza       : %s\n", pizza)
}
