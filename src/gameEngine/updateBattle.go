package gameEngine

import (
	"github.com/gen2brain/raylib-go/raylib"

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
		if engine.monster[engine.monsterBattle].hp <= 0 {
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