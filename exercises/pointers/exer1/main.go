//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change
package main

import "fmt"

// Security struct holds security states
type Security struct {
	name   string
	secTag bool
}

// active, inactive hold true, false boolean value
const (
	active, inactive = true, false
)

func activateSec(item *Security, state bool) {
	item.secTag = state
}

func main() {

	// Define a list of security items
	frontDoor := Security{name: "Front Door", secTag: active}
	windows := Security{name: "All Windows", secTag: active}
	garageDoor := Security{name: "Garage Door", secTag: active}
	secCams := Security{name: "All Security Cams", secTag: active}

	// homeSecurity slice holds all home security items
	homeSecurity := []Security{frontDoor, windows, garageDoor, secCams}

	// Print out the homeSecurity items and it states
	fmt.Println("\nHome Security Items set to Active States")
	for i, v := range homeSecurity {
		fmt.Printf("%d %s security state active %t\n", i+1, v.name, v.secTag)
	}

	// Deactivate individual security item
	fmt.Println("\nDeactivate some security items")
	activateSec(&frontDoor, inactive)
	activateSec(&windows, inactive)
	fmt.Printf("%s security state %t\n", frontDoor.name, frontDoor.secTag)
	fmt.Printf("%s security state %t\n", windows.name, windows.secTag)

	// Deactivate all securities
	fmt.Println("\nDeactivate all security items")
	for i, v := range homeSecurity {
		activateSec(&v, inactive)
		fmt.Printf("%d %s security state: %t\n", i+1, v.name, v.secTag)
	}
}
