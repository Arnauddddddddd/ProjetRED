package gameEngine

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func inventorySelector(engine *EngineStruct) {
	for i := 0; i < len(engine.character.inventory); i++ {
		if i < 4 {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && (int32(640 + 79*i) < rl.GetMouseX() && rl.GetMouseX() < int32(640 + 79*i)+79) && int32(365) < rl.GetMouseY() && rl.GetMouseY() < int32(445) {
				inventoryEffect(engine, i)
			}
		}
	}
	if len(engine.character.inventory) > 4 {
		for j := 4; j < len(engine.character.inventory); j++ {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && (int32(640 + 79*(j-4)) < rl.GetMouseX() && rl.GetMouseX() < int32(640 + 79*(j-4))+79) && int32(455) < rl.GetMouseY() && rl.GetMouseY() < int32(530) {
				inventoryEffect(engine, j)
			}
		}
	}
}

func inventoryEffect(engine *EngineStruct, i int) {
	if engine.character.inventory[i].outBattle {
		engine.character.hp += engine.character.inventory[i].hpUp
		if engine.character.hp >= engine.character.hpMax {
			engine.character.hp = engine.character.hpMax
		}
		engine.character.damage += engine.character.inventory[i].damageUp
		engine.player.Speed += float32(engine.character.inventory[i].speedUp)
		engine.character.inventory = append(engine.character.inventory[:i], engine.character.inventory[i+1:]...)
	} else {
		fmt.Println("NON NON")
	}
}
