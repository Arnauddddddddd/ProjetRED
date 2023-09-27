package gameEngine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func inventorySelector(engine *EngineStruct, coordx int, coordy int, cordy2 int, cordy3 int, cordy4 int, slotLen int) {
	for i := 0; i < len(engine.character.inventory); i++ {
		if i < 4 {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && int32((coordx) + slotLen*i) < rl.GetMouseX() && rl.GetMouseX() < int32((coordx) + slotLen*i+slotLen) && int32(365) < rl.GetMouseY() && rl.GetMouseY() < int32(cordy2) {
				inventoryEffect(engine, i)
			}
		}
	}
	if len(engine.character.inventory) > 4 {
		for j := 4; j < len(engine.character.inventory); j++ {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && int32((coordx) + slotLen*(j-4)) < rl.GetMouseX() && rl.GetMouseX() < int32(coordx + slotLen*(j-4)+ slotLen) && int32(cordy3) < rl.GetMouseY() && rl.GetMouseY() < int32(cordy4) {
				inventoryEffect(engine, j)
			}
		}
	}
}

func inventoryEffect(engine *EngineStruct, u int) {
	if engine.character.inventory[u].battle && engine.battle.inBattle { 
		engine.character.damage += engine.character.inventory[u].damageUp
		engine.character.hp += engine.character.inventory[u].hpUp
		if engine.character.hp >= engine.character.hpMax {
			engine.character.hp = engine.character.hpMax
		}
		if engine.character.inventory[u].gender == "Axe" {
			engine.monster[engine.battle.monsterBattle].hp -= 250
		}
		if engine.character.inventory[u].gender == "Fork" {
			engine.monster[engine.battle.monsterBattle].hp -= 100
		}
		engine.playerTurn = false
		engine.character.inventory = append(engine.character.inventory[:u], engine.character.inventory[u+1:]...)

	} else if engine.character.inventory[u].outBattle && !engine.battle.inBattle {
		engine.character.hp += engine.character.inventory[u].hpUp
		engine.player.Speed += float32(engine.character.inventory[u].speedUp)
		if engine.character.inventory[u].gender == "Improvement" {
			engine.character.hpMax *= 1.3
			engine.character.damageBase *= 1.3
			engine.character.damage *= 1.3
			engine.character.hp += 5000
		}
		if engine.character.hp >= engine.character.hpMax {
			engine.character.hp = engine.character.hpMax
		}
		engine.character.inventory = append(engine.character.inventory[:u], engine.character.inventory[u+1:]...)
	}
}
