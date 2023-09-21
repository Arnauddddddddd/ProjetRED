package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func inventorySelector() {
	for i := 0; i < len(engine.character.inventory); i++ {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) && (int32(497 + 67*i) < rl.GetMouseX() && rl.GetMouseX() < int32(497 + 67*i)+67) && 450 < rl.GetMouseY() && 
		rl.GetMouseY() < int32(520) {
			if len(engine.character.inventory) >= i {
				engine.character.hp += engine.character.inventory[i].hpUp
				if engine.character.hp >= engine.character.hpMax {
					engine.character.hp = engine.character.hpMax
				}
				engine.character.damage += engine.character.inventory[i].damageUp
				engine.player.Speed += float32(engine.character.inventory[i].speedUp)
				engine.character.inventory = append(engine.character.inventory[:i], engine.character.inventory[i+1:]...)
			}
		}
	}
}
