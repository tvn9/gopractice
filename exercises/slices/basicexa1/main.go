//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

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
	for i, p := range partList {
		fmt.Printf("%d %v\n", i, p)
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

	for i, p := range partList {
		fmt.Println(i, p)
	}

	// slice part 2, 3
	partList = partList[1:3]
	fmt.Println("\nThe list with only 2 part left")
	for i, p := range partList {
		fmt.Printf("%d %v\n", i, p)
	}

}
