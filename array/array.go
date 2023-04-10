package main

import "fmt"

func main() {
	// Declare an array
	var productNames [4]string
	// Define an array
	productNames = [4]string{"A Book", "B Book", "C Book"}

	// Adding element to existing array
	productNames[3] = "D Bock"

	// define an array
	prices := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Print the array data
	fmt.Println(prices)
	fmt.Println(productNames)

	// How to access array indexes
	fmt.Println(prices[0], prices[3], prices[len(prices)-1], prices[1:6], prices[2:], prices[:5])
	fmt.Println(productNames[:], productNames[1:4], productNames[len(productNames)-1])

	// Can I change arrange value in any array index?
	prices[0] = 10
	fmt.Println(prices)

	// Can I remove a value from and array?
	prices1 := prices[3:]
	fmt.Println(prices1, prices)

	priceList1 := prices[2:]

	fmt.Println(priceList1)

	fmt.Println(prices)
}
