// --Summary:
//
//	Implement receiver functions to create stat modifications
//	for a video game character.
//
// --Requirements:
// * Implement a player having the following statistics:
//   - Health, Max Health
//   - Energy, Max Energy
//   - Name
//   - Implement receiver functions to modify the `Health` and `Energy`
//     statistics of the player.
//   - Print out the statistic change within each function
//   - Execute each function at least once
package main

import "fmt"

type Player struct {
	name      string
	health    uint
	maxHealth uint
	energy    uint
	maxEnergy uint
}

func (p *Player) addMaxHealth(h uint) {
	if h > 0 {
		p.maxHealth += h
	}
}

func (p *Player) setPlayerName(n string) {
	if n != "" {
		p.name = n
	}
}

func (p *Player) addMaxEnergy(e uint) {
	if e > 0 {
		p.maxEnergy += e
	}
}

func (p *Player) addHealth(h uint) {
	if h > 0 {
		p.health += h
	}
	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}
}

func (p *Player) applyDamage(d uint) {
	if p.health-d > p.health {
		p.health = 0
	} else {
		p.health -= d
	}
}

func (p *Player) addEnergy(e uint) {
	if e > 0 {
		p.energy += e
	}
	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	}
}

func (p *Player) consumeEnergy(e uint) {
	if p.energy-e > p.maxEnergy {
		p.energy = 0
	} else {
		p.energy -= e
	}
}

func (p *Player) printPlayerStat() {
	fmt.Printf("Player %s, health: %d, energy: %d, maxHealth: %d, maxEnergy: %d\n",
		p.name, p.health, p.energy, p.maxHealth, p.maxEnergy)
}

func main() {

	// Create a players with initial health and energy to max level
	player1 := &Player{}
	player1.setPlayerName("Mike")
	player1.addMaxHealth(500)
	player1.addMaxEnergy(500)
	player1.addHealth(100)
	player1.addEnergy(500)
	player1.applyDamage(99)
	player1.addHealth(80)
	player1.consumeEnergy(20)

	player1.consumeEnergy(9999)

	// print players status
	player1.printPlayerStat()

}
