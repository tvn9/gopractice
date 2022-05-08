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

func (d Direction) NewString() string {
	return []string{"North", "East", "South", "West"}[d]
}

func main() {
	north := North
	fmt.Printf("North: %d %T\n", north, north)
	fmt.Printf("East: %d %T\n", East, East)
	fmt.Printf("South: %d %T\n", South, South)
	fmt.Printf("West: %d %T\n", West, West)

	n, s, e, w := North, South, East, West

	fmt.Println()
	fmt.Println(n.String())
	fmt.Println(s.String())
	fmt.Println(e.String())
	fmt.Println(w.String())

	fmt.Println()
	fmt.Println(n.NewString())
	fmt.Println(s.NewString())
	fmt.Println(e.NewString())
	fmt.Println(w.NewString())

}
