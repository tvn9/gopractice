// Range example
package main

import "fmt"

type Part struct {
	sn       string
	color    string
	material string
}

func main() {

	// partList variable holds list of parts use in an assembly line
	partList := []Part{
		{
			sn:       "P001",
			color:    "Red",
			material: "PVC",
		},
		{
			sn:       "P002",
			color:    "Blue",
			material: "Aluminum",
		},
		{
			sn:       "P003",
			color:    "Green",
			material: "PVC",
		},
	}

	// Print out the partList
	fmt.Println("\nThe first list with 3 parts")
	// For range loop
	for i, p := range partList {
		fmt.Printf("%d %v\n", i+1, p)
	}

	// Add two new parts to the partList
	p1 := Part{
		sn:       "P004",
		color:    "Black",
		material: "Fiberglass",
	}
	p2 := Part{
		sn:       "P005",
		color:    "Grey",
		material: "Fiberglass",
	}

	partList = append(partList, p1, p2)
	// Print out the partList
	fmt.Println("\nThe list after adding 2 more parts")
	// Print part list using for range loop
	for i, p := range partList {
		fmt.Printf("%d %s\n", i+1, p)
	}

	// slice part 2, 3, print out slice of parts with for range loop
	partList = partList[1:3]
	fmt.Println("\nThe list with only 2 part left")
	for i, p := range partList {
		fmt.Printf("%d %v\n", i+1, p)
	}
}
