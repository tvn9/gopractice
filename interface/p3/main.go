// Interface Example
package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func PrepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for _, d := range dishes {
		dish := d
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
}

func main() {
	dishes := []Preparer{
		Chicken("Chicken Wings"),
		Salad("Sidar Salad"),
	}
	PrepareDishes(dishes)
}
