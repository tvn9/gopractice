// interface example
package main

import "fmt"

type Resetter interface {
	Reset()
}

type Coordinate struct {
	X, Y int
}
type Player struct {
	health   int
	position Coordinate
}

func (p *Player) Reset() {
	p.health = 100
	p.position = Coordinate{0, 0}
}

func Reset(r Resetter) {
	r.Reset()
}

func main() {
	// Create a Player
	p1 := Player{
		health: 50,
		position: Coordinate{
			X: 10,
			Y: 30,
		},
	}

	p2 := Player{50, Coordinate{20, 10}}

	fmt.Println("Player 1", p1)
	Reset(&p1)
	fmt.Println("Player 1", p1)

	fmt.Println("Player 2", p2)
	p2.Reset()
	fmt.Println("Player 2", p2)
}
