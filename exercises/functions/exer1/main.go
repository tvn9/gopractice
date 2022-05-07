//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once
package main

import "fmt"

var (
	health, maxHealth int
	energy, maxEnergy int
	gameOver          bool
)

type Player struct {
	name      string
	health    int
	maxHealth int
	energy    int
	maxEnergy int
}

// Modify player health
func (p *Player) gainHealth(h int) {
	p.health += h
	if p.health > p.maxHealth {
		p.health = maxHealth
	}
}
func (p *Player) gainEnergy(e int) {
	p.energy += e
	if p.energy > p.maxEnergy {
		p.energy = maxEnergy
	}
}

func (p *Player) lossHealth(h int) bool {
	p.health -= h
	if p.health < 0 {
		p.health = health
		gameOver = true
	} else {
		gameOver = false
	}
	return gameOver
}
func (p *Player) lossEnergy(e int) bool {
	p.energy -= e
	if p.energy < 0 {
		p.energy = energy
		gameOver = true
	} else {
		gameOver = false
	}
	return gameOver
}

func (p *Player) printPlayerStat() {
	fmt.Printf("Player %q health: %v, energy: %v\n", p.name, p.health, p.energy)
}

func init() {
	maxHealth = 100
	maxEnergy = 1000
	health = 0
	energy = 0
}

func main() {
	// Create a players with initial health and energy to max level
	player1 := Player{
		name:   "Mike",
		health: maxHealth,
		energy: maxEnergy,
	}
	player2 := Player{
		name:   "Kevin",
		health: maxHealth,
		energy: maxEnergy,
	}

	// Print out player stats
	player1.printPlayerStat()
	player2.printPlayerStat()

	// Increase Players health
	fmt.Println("\nIncrease player1 health by 10, and player2 by 20")
	player1.gainHealth(10)
	player1.gainEnergy(500)
	player2.gainHealth(20)
	player2.gainEnergy(300)

	// print players status
	player1.printPlayerStat()
	player2.printPlayerStat()

	// decrease players health
	fmt.Println("\ndecrease players health and print status")
	p1h := player1.lossHealth(150)
	if p1h {
		fmt.Printf("Player %s loss, game over!\n", player1.name)
	} else {
		fmt.Printf("Player %s still alive!\n", player1.name)
	}
	p1e := player1.lossEnergy(500)
	if !p1e {
		fmt.Printf("Player %s is remain energy %d\n", player1.name, player1.energy)
	} else {
		fmt.Printf("Player %s has no power %d\n", player1.name, player1.energy)
	}
	p2h := player2.lossHealth(99)
	if p2h {
		fmt.Printf("Player %s loss, game over!\n", player2.name)
	} else {
		fmt.Printf("Player %s still alive!\n", player2.name)
	}
	p2e := player2.lossEnergy(1200)
	if !p2e {
		fmt.Printf("Player %s energy level %d\n", player2.name, player2.energy)
	} else {
		fmt.Printf("Player %s loss all energy %d\n", player2.name, player2.energy)
	}
	player1.printPlayerStat()
	player2.printPlayerStat()

}
