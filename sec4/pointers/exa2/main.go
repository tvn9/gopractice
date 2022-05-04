// Pointer example
package main

import "fmt"

type GameScore struct {
	score int
}

func changeScore(g *GameScore) {
	g.score += 1
	fmt.Println("Game score", g)
}

func replace(old *string, new string, g *GameScore) {

}
func main() {

	// Change game score
	game1 := GameScore{}

	game1.score = 8

}
