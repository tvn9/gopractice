// Receiver Functions Example
package main

import "fmt"

// Coordinate struct hold X and Y example Coordinate point
type Coordinate struct {
	X, Y int
}

// moveCoordinate function read in x, y type int data and Coordinate struct type
func moveCoordinate(x, y int, coord *Coordinate) {
	coord.X += y
	coord.Y += y
}

// moveCoordinate with a Go specific syntax call Receiver function
// where receiver function take an object data type as input parameter
func (coord *Coordinate) moveCoordinate(x, y int) {
	coord.X += x
	coord.Y += y
}

func main() {
	// defind an object literal of type Coordinate
	coord := Coordinate{5, 5}

	// Call moveCoordinate func to move coordinate X, Y by 2 points
	moveCoordinate(2, 2, &coord)
	fmt.Printf("Coordinate X=5 + 2, Y=5 + 2 => X = %d, Y = %d\n", coord.X, coord.Y)

	// Call function receiver by dot notation
	coord.moveCoordinate(-2, -2)
	fmt.Printf("Coordinate X=7 - 2, Y=7 - 2 => X = %d, Y = %d\n", coord.X, coord.Y)

}
