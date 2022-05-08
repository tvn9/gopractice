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

// moveCoordinate1 with receiver references Coordinate object
// without pointer (*Coordinate), this will modifiy X, Y value
// for the copy object, not the original Coordinate object
func (c1 Coordinate) moveCoordinate1(x, y int) Coordinate {
	c1.X += x
	c1.Y += y
	// return the copy of the Coordinate object
	return c1
}

// moveCoordinate with a Go specific syntax call Receiver function
// where receiver function take an object data type as input parameter
func (coord *Coordinate) moveCoordinate(x, y int) {
	coord.X += x
	coord.Y += y
}

func main() {
	// defind an object literal of type Coordinate
	coord := &Coordinate{5, 5} // This will works
	// coord := Coordinate{5, 5} // This will works too
	fmt.Printf("\nCoordinate initial value X=%d, Y=%d\n", coord.X, coord.Y)

	// Call moveCoordinate func to move coordinate X, Y by 2 points
	moveCoordinate(2, 2, coord)
	fmt.Printf("\nCoordinate X=5 + 2, Y=5 + 2 => X = %d, Y = %d\n", coord.X, coord.Y)

	// call function receiver without pointer to Coordinate object
	c1 := coord.moveCoordinate1(3, 3)
	fmt.Printf("\nc1 X=7, Y=7; add 3 to x, y => X=%d, Y=%d\n", c1.X, c1.Y)
	fmt.Printf("original coord.X, coord.Y => X=%d, Y=%d\n", coord.X, coord.Y)

	// Call function receiver by dot notation
	fmt.Printf("\nCoordinate X = %d, Y = %d ", coord.X, coord.Y)
	coord.moveCoordinate(-2, -2)
	fmt.Printf("after added -2 to x, y; => X = %d, Y = %d\n", coord.X, coord.Y)
}
