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
	list := [...]ShopList{
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
		{
			ItemNo:   1004,
			ItemName: "Butter",
			Price:    9.99,
		},
	}
	fmt.Println(list)
	fmt.Println("Total number of item:", len(list))

	totalCost := calcTotal(list)
	fmt.Printf("Total cost = %.2f\n", totalCost)
}

func calcTotal(l [4]ShopList) float32 {
	var total float32

	for i := 0; i < len(l); i++ {
		item := l[i]
		total += item.Price
	}

	return total
}
