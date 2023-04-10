package main

import "fmt"

type Product struct {
	Title       string
	Description string
	Price       float64
}

func main() {
	// Array with 3 hobbies
	hobbies := [3]string{"Tennis", "Home Contruction", "Road Travel"}
	// Print array to screen
	fmt.Println(hobbies)

	// The first element
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// new array with first element
	newHobbies := hobbies[0:2]
	fmt.Println(newHobbies)

	// The second and third element in a new list
	newHobbies = hobbies[1:3]
	fmt.Println(newHobbies)

	// Create a slice
	hobbiesslice := []string{"Tennis", "Home Construction"}

	// Adding new element to slice
	hobbiesslice = append(hobbiesslice, "Road Travel")
	fmt.Println(hobbiesslice)

	// Creat a new struct
	vihicle := []Product{
		{
			Title:       "F150",
			Description: "F150 Puckup Truck",
			Price:       39000.00,
		},
		{
			Title:       "Silverado",
			Description: "Silverado 1500 Pickup Truck",
			Price:       38000.00,
		},
	}

	car := Product{
		Title:       "Toyota",
		Description: "Prius Hybrid with 60Miles MPG",
		Price:       28000.00,
	}

	// Adding a new product
	vihicle = append(vihicle, car)
	fmt.Println(vihicle)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
