package gameEngine

import (
	"github.com/gen2brain/raylib-go/raylib"

	"fmt"

)

func updateBattle (engine *EngineStruct) {

	engine.run = !rl.WindowShouldClose()

	engine.framCount++

	if engine.framCount % 8 == 1 { engine.monster[engine.monsterBattle].frameCount++ }
	if engine.monster[engine.monsterBattle].frameCount > 3 { engine.monster[engine.monsterBattle].frameCount = 0 }
	engine.monster[engine.monsterBattle].Src.X = engine.monster[engine.monsterBattle].Src.Width * float32(engine.monster[engine.monsterBattle].frameCount)

	if engine.monster[engine.monsterBattle].speed > engine.character.speed {
		engine.playerTurn = false
	}

	if engine.playerTurn {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) &&  (126 < rl.GetMouseX() && rl.GetMouseX() <  468 && 758 < rl.GetMouseY() && rl.GetMouseY() < 840) {
			engine.monster[engine.monsterBattle].hp -= engine.character.damage
			engine.playerTurn = false
		}
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) &&  (1140 < rl.GetMouseX() && rl.GetMouseX() <  1492 && 758 < rl.GetMouseY() && rl.GetMouseY() < 840) {
			fmt.Println(rl.GetMouseX())
		}
		if engine.monster[engine.monsterBattle].hp <= 0 {
			engine.character.gold += engine.monster[engine.monsterBattle].goldLoot
			engine.character.damage = engine.character.damageBase
			engine.monster[engine.monsterBattle].alive = false
			engine.battle = false
		}
		if engine.character.hp <= 0 {
			engine.character.damage = engine.character.damageBase
			engine.character.alive = false
			engine.battle = false
		}
	} else {
		engine.character.hp -= engine.monster[engine.monsterBattle].damage
		engine.playerTurn = true
	}
}