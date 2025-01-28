/* It's a simple rpg game, written on go language.
   Maded by Daniel. @just_perfection
*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

// Player vars
var is_player_in_dialog bool
var player_kills int
var player_gold int
var player_dmg int
var player_hp int

// Enemy vars
var enemy_dmg int
var enemy_hp int

func init_player() {
	is_player_in_dialog = false
	player_kills = 0
	player_gold = 0
	player_dmg = 2
	player_hp = 10
}

func init_enemy() {
	enemy_dmg = 2
	enemy_hp = 4
}

// Return lowercase string without spaces, tabs
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

func go_tavern(p_gold int, p_kills int) {

}

func main() {
	fmt.Println("Hello, traveler! You are ready?")

	var input_r string
	fmt.Scan(&input_r)
	input_r = to_vld_str(input_r)

	if input_r != "yes" {
		return
	}

	init_player()
	init_enemy()
	// Main game loop
	for true {
		if !is_player_in_dialog {
			fmt.Println("They're goblin! 'A'ttack him or 'l'eave? Also you can go to 't'avern.")
		}

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
				break
			}

			player_hp = attack_player(player_hp, enemy_dmg)
			if is_player_killed(player_hp) {
				fmt.Println("You are dead :(")
				break
			}

			fmt.Printf("You attack goblin! He's hp [%v] -%v\n", enemy_hp, player_dmg)
			fmt.Printf("Goblin attack you! Your hp is [%v] -%v\n", player_hp, enemy_dmg)
		case "t":
			// TODO: Extract this variables to constants
			fmt.Printf("You in tavern. Do you want drink an el? [+%v], cost: [%v].", 3, 5)
			is_player_in_dialog = true

		case "l":
			if player_kills == 0 {
				fmt.Println("You leave from danger zone!")
				return
			} else {
				fmt.Printf("Good job! You killed %v goblin(s)\n", player_kills)
				return
			}

		default:
			fmt.Println("Choose please!")
			fmt.Printf("%v\n", errors.New("throw error: invalid choose."))
			return
		}

	}
}
