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
	return str
}

// Return enemy hp after player attack
func attack_enemy(e_hp int, p_atck int) int {
	return e_hp - p_atck
}

// Return player hp after enemy attack
func attack_player(p_hp int, e_atck int) int {
	return p_hp - e_atck
}

func is_enemy_killed(e_hp int) bool {
	// It's a sugar in go lang.
	return e_hp == 0
}

func is_player_killed(p_hp int) bool {
	return p_hp == 0
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

	enemy_hp := 4
	enemy_dmg := 2

	player_hp := 10
	player_dmg := 2
	player_gold := 0
	_ = player_gold
	player_kills := 0

	for is_continue {
		fmt.Println("They're goblin! 'A'ttack him or 'l'eave? Also you can go to 't'avern.")
		var atck_choose string
		fmt.Scan(&atck_choose)
		atck_choose = to_vld_str(atck_choose)
		switch atck_choose {
		case "a":
			enemy_hp = attack_enemy(enemy_hp, player_dmg)
			if is_enemy_killed(enemy_hp) {
				player_kills += 1
				// Init new enemy.
				enemy_hp = 4
				fmt.Printf("You killed a goblin! Kills: %v\n", player_kills)
				continue
			}
			player_hp = attack_player(player_hp, enemy_dmg)
			if is_player_killed(player_hp) {
				fmt.Println("You are dead :(")
				is_continue = false
				continue
			}
			fmt.Printf("You attack goblin! He's hp [%v] -%v\n", enemy_hp, player_dmg)
			fmt.Printf("Goblin attack you! Your hp is [%v] -%v\n", player_hp, enemy_dmg)
		case "t":
			// TODO: Extract this variables to constants
			fmt.Printf("You in tavern. Do you want drink an el? [+%v], cost: [%v].", 3, 5)

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
