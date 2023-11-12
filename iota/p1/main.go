// Iota example 1
package main

import "fmt"

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

/*
func (d Direction) String() string {
	switch d {
	case North:
		return fmt.Sprintf("North")
	case South:
		return fmt.Sprintf("South")
	case East:
		return fmt.Sprintf("East")
	case West:
		return fmt.Sprintf("West")
	default:
		return "other direction"
	}
}
*/

// Alternative way to create a function with that return the index number
func (d Direction) String() string {
	return []string{"North", "East", "South", "West"}[d]
}

func main() {
	north := North
	fmt.Println(north)
}
