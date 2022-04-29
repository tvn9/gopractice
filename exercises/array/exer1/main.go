//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again
package main

import (
	"fmt"
)

type ShopList struct {
	ItemNo   int
	ItemName string
	Price    float32
}

func main() {
	// Create a shopping list
	list := [4]ShopList{
		{
			ItemNo:   1001,
			ItemName: "Chicken",
			Price:    11.50,
		},
		{
			ItemNo:   1002,
			ItemName: "Eggs",
			Price:    8.99,
		},
		{
			ItemNo:   1003,
			ItemName: "Ice Screem",
			Price:    9.99,
		},
	}

	totalItem := countItems(list)

	// Print the list item on the list
	fmt.Println(list[totalItem-1])

	// Print total number of items
	fmt.Println("Total number of item:", totalItem)

	// Print the total cost
	totalCost := calcTotal(list)
	fmt.Printf("Total cost = %.2f\n", totalCost)

	// Add a new item on the ShopList
	list[3].ItemNo = 1004
	list[3].Price = 12.99

	// After adding a fourth item
	fmt.Println("After adding the fourth item")

	// Print list
	printList(list)
}

func calcTotal(l [4]ShopList) float32 {
	var total float32

	for i := 0; i < len(l); i++ {
		item := l[i]
		total += item.Price
	}

	return total
}

func countItems(list [4]ShopList) int {
	total := 0

	for i := 0; i < len(list); i++ {
		item := list[i]
		if item.ItemName != "" || item.ItemNo > 0 {
			total += 1
		}
	}
	return total
}

func printList(list [4]ShopList) {
	var (
		totalItem int
		totalCost float32
	)

	for i := 0; i < len(list); i++ {
		item := list[i]

		if item.ItemName != "" || item.ItemNo > 0 {
			totalItem += 1
			totalCost += item.Price
		}
	}

	fmt.Printf("Total number of items = %d\n", totalItem)
	fmt.Printf("Total cost for all tems = %.2f\n", totalCost)
}
