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

type Player struct {
	Name      string
	Health    int
	MaxHealth int
	Energy    int
	MaxEnergy int
}

func (p *Player) modifyHealth(health int) {
	p.Health += health
	fmt.Println("Health modified by", health)
}

func (p *Player) modifyEnergy(energy int) {
	p.Energy += energy
	fmt.Println("Energy modified by", energy)
}

func main() {
	p := Player{"Player1", 100, 100, 100, 100}
	fmt.Println("Initial Player stats:", p)
	p.modifyHealth(10)
	fmt.Println("Player stats after health modification:", p)
	p.modifyEnergy(-10)
	fmt.Println("Player stats after energy modification:", p)
}
