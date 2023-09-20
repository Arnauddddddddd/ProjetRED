package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func inventorySelector() {
	for i := 0; i < len(character.inventory); i++ {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) && (int32(497 + 67*i) < rl.GetMouseX() && rl.GetMouseX() < int32(497 + 67*i)+67) && 450 < rl.GetMouseY() && 
		rl.GetMouseY() < int32(520) {
			if len(character.inventory) >= i {
				character.hp += character.inventory[i].hpUp
				if character.hp >= character.hpMax {
					character.hp = character.hpMax
				}
				character.damage += character.inventory[i].damageUp
				player.Speed += float32(character.inventory[i].speedUp)
				character.inventory = append(character.inventory[:i], character.inventory[i+1:]...)
			}
		}
	}
}
