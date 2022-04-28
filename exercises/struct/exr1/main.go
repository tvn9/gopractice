//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

// Rectangle struct hold the Rectangle coordinates variable
type Rectangle struct {
	X1, X2 uint32
	Y1, Y2 uint32
}

func main() {

	// Create a rectangle
	rec1 := Rectangle{
		X1: 0,
		X2: 20,
		Y1: 0,
		Y2: 10,
	}

	// variable x and y compute the width and high of the rectangle using the X1, X2, Y1, Y2 coordinates
	w := rec1.X1 + rec1.X2
	h := rec1.Y1 + rec1.Y2

	// Print the rec1 width and high result to terminal
	fmt.Printf("Rec1 width = %d\n", w)
	fmt.Printf("Rec1 high = %d\n", h)

	// get rec1 area by calling calcArea function
	rec1Area := calcArea(w, h)
	fmt.Printf("Area of rec1 = %d\n", rec1Area)

	rec1Perimeter := calcPerimeter(w, h)
	fmt.Printf("rec1 perimeter = %d\n", rec1Perimeter)

	// increase the size of the w and h by 2
	w *= 2
	h *= 2

	rec2Area := calcArea(w, h)
	rec2Perimeter := calcPerimeter(w, h)
	fmt.Printf("rec2 area = %d\n", rec2Area)
	fmt.Printf("rec2 perimeter = %d\n", rec2Perimeter)
}

// calcArea function calculate the are of a rectangle
func calcArea(w, y uint32) uint32 {
	area := w * y
	return area
}

func calcPerimeter(w, h uint32) uint32 {
	perimeter := (w + h) * 2
	return perimeter
}
