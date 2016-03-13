package main

import (
	"fmt"
	"math"
)

type item struct {
	Name                string
	Cost, Damage, Armor int
}

var (
	weapons = []item{
		item{"Dagger", 8, 4, 0},
		item{"Shortsword", 10, 5, 0},
		item{"Warhammer", 25, 6, 0},
		item{"Longsword", 40, 7, 0},
		item{"Greataxe", 74, 8, 0},
	}

	armor = []item{
		item{"No armor", 0, 0, 0}, // no armor is a choice
		item{"Leather", 13, 0, 1},
		item{"Chainmail", 31, 0, 2},
		item{"Splintmail", 53, 0, 3},
		item{"Bandedmail", 75, 0, 4},
		item{"Platemail", 102, 0, 5},
	}

	rings = []item{
		item{"No ring", 0, 0, 0},
		item{"Damage+1", 25, 1, 0},
		item{"Damage+2", 50, 2, 0},
		item{"Damage+3", 100, 3, 0},
		item{"Defense+1", 20, 0, 1},
		item{"Defense+2", 40, 0, 2},
		item{"Defense+3", 80, 0, 3},
	}
)

func main() {
	boss_damage := 8
	boss_armor := 1

	min_cost := math.MaxInt32
	for _, w := range weapons {
		for _, a := range armor {
			for i, r1 := range rings {
				for j, r2 := range rings {
					if i == j && i != 0 {
						continue
					}
					cost := w.Cost + a.Cost + r1.Cost + r2.Cost
					my_damage := w.Damage + a.Damage + r1.Damage + r2.Damage
					my_armor := w.Armor + a.Armor + r1.Armor + r2.Armor
					offense := my_damage - boss_armor
					defense := boss_damage - my_armor
					fmt.Printf("Advantage: %d\tCost %5d\tDamage %5d\tArmor %5d: %s, %s, %s, %s\n",
						offense-defense, cost, my_damage, my_armor, w.Name, a.Name, r1.Name, r2.Name)

					my_points := 100
					boss_points := 104
					for {
						boss_points -= (my_damage - boss_armor)
						fmt.Printf("The player deals %d-%d = %d damage; the boss goes down to %d hit points.\n",
							my_damage, boss_armor, my_damage-boss_armor, boss_points)
						if boss_points <= 0 {
							fmt.Println("The player wins!\n")
							if cost < min_cost {
								min_cost = cost
							}
							break
						}

						my_points -= (boss_damage - my_armor)
						fmt.Printf("The boss deals %d-%d = %d damage; the player goes down to %d hit points.\n",
							boss_damage, my_armor, boss_damage-my_armor, my_points)
						if my_points <= 0 {
							fmt.Println("The boss wins!\n")
							break
						}
					}
					//fmt.Printf("Cost %5d\tDamage %5d\tArmor %5d: %s, %s, %s, %s\n", cost, my_damage, my_armor, w.Name, a.Name, r1.Name, r2.Name)
					//fmt.Printf("   offense %5d\tdefense%d\n", offense, defense)
				}
			}
		}
	}
	fmt.Println("Min cost: ", min_cost)
}
