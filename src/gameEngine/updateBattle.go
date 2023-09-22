package gameEngine

import (
	"github.com/gen2brain/raylib-go/raylib"

	"fmt"
)

func updateBattle (engine *EngineStruct) {

	engine.run = !rl.WindowShouldClose()

	if engine.framCount % 8 == 1 { engine.monster[engine.monsterBattle].frameCount++ }
	if engine.monster[engine.monsterBattle].frameCount > 3 { engine.monster[engine.monsterBattle].frameCount = 0 }
	engine.monster[engine.monsterBattle].Src.X = engine.monster[engine.monsterBattle].Src.Width * float32(engine.monster[engine.monsterBattle].frameCount)

	if engine.monster[engine.monsterBattle].speed > engine.character.speed {
		engine.playerTurn = false
	}

	if engine.playerTurn {
		if rl.IsMouseButtonPressed(rl.MouseRightButton) {
			engine.monster[engine.monsterBattle].hp -= engine.character.damage
			engine.playerTurn = false
		}
		for i := 0; i < len(engine.character.inventory); i++ {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && (int32(458 + 69*i) < rl.GetMouseX() && rl.GetMouseX() < int32(458 + 69*i)+69) && 770 < rl.GetMouseY() && 
			rl.GetMouseY() < int32(840) {
				if engine.character.inventory[i].battle {
					if len(engine.character.inventory) >= i {
						engine.character.hp += engine.character.inventory[i].hpUp
						if engine.character.hp >= engine.character.hpMax {
							engine.character.hp = engine.character.hpMax
						}
						engine.character.damage += engine.character.inventory[i].damageUp
						engine.player.Speed += float32(engine.character.inventory[i].speedUp)
						engine.character.inventory = append(engine.character.inventory[:i], engine.character.inventory[i+1:]...)
					}
					engine.playerTurn = false
				} else {
					fmt.Println("NON")
				}
			}
		}
		if engine.monster[engine.monsterBattle].hp <= 0 {
			engine.monster[engine.monsterBattle].alive = false
			engine.battle = false
		}
		if engine.character.hp <= 0 {
			engine.character.alive = false
			engine.battle = false
		}
	} else {
		engine.character.hp -= engine.monster[engine.monsterBattle].damage
		engine.playerTurn = true
	}
}