// Iota example
package main

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) string() string {
	return []string{"North", "East", "South", "West"}[d]
}

func main() {
	north := North
	fmt.Printf("North: %d %T\n", north, north)
	fmt.Printf("East: %d %T\n", East, East)
	fmt.Printf("South: %d %T\n", South, South)
	fmt.Printf("West: %d %T\n", West, West)

	str := north.string()
	for i, v := range str {
		fmt.Printf("Testing iota %d %v\n", i, v)
	}

}
