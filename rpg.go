/* It's a simple rpg game, written on go language.
   Maded by Daniel. @just_perfection
*/

package main

import (
	"fmt"
	"strings"
)

func to_vld_str(str string) string {
	str = strings.ToLower(str)
	strings.Trim(str, str)
	return str
}

func main() {
	var is_continue bool
	fmt.Println("Hello, traveler! You are ready?")

	var input_r string
	fmt.Scan(&input_r)
	input_r = to_vld_str(input_r)

	if input_r == "yes" {
		is_continue = true
	}

	player_hp := 10
	player_dmg := 2
	player_kills := 0
	enemy_hp := 4
	enemy_damage := 1
	for is_continue {
		fmt.Println("They're goblin! 'A'ttack him or 'l'eave?")
		var atck_choose string
		fmt.Scan(&atck_choose)
		atck_choose = to_vld_str(atck_choose)
		switch atck_choose {
		case "a":
			enemy_hp -= player_dmg
			player_hp -= enemy_damage
			if enemy_hp == 0 {
				player_kills += 1
				enemy_hp = 4
				fmt.Printf("You killed goblin. Kills: %v\n", player_kills)
				continue
			}
			fmt.Printf("You attack goblin! He's hp [%v] -%v\n", enemy_hp, player_dmg)
			fmt.Printf("Goblin attack you! Your hp is [%v] -%v", player_hp, enemy_damage)
		case "l":
			if player_kills == 0 {
				fmt.Println("You leave from danger zone!")
				is_continue = false
			} else {
				fmt.Printf("Good job! You killed %v goblins\n", player_kills)
				is_continue = false
			}
		default:
			fmt.Println("Choose please!")
			is_continue = false
		}

	}
}
