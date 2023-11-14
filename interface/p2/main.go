// Interface Example
package main

import "fmt"

type Coordinate struct {
	X int
	Y int
}

type Resetter interface {
	Reset()
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
	p1 := Player{}
	p1.health = 100
	p1.position.X = 5
	p1.position.Y = 5
	fmt.Println(p1)

	Reset(&p1)
	fmt.Println(p1)
}
