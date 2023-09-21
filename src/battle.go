package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func battle() {
	if engine.monster[engine.monsterBattle].speed > engine.character.speed {
		engine.playerTurn = false
	}

	if engine.playerTurn {
		if rl.IsMouseButtonPressed(rl.MouseRightButton) {
			engine.monster[engine.monsterBattle].hp -= engine.character.damage
			engine.playerTurn = false
		}
		if engine.monster[engine.monsterBattle].hp <= 0 {
			engine.monster[engine.monsterBattle].alive = false
			engine.battle = false
			engine.cam.Zoom = 3.5
		}
		if engine.character.hp <= 0 {
			engine.character.alive = false
			engine.battle = false
			engine.cam.Zoom = 3.5
		}
	} else {
		engine.character.hp -= engine.monster[engine.monsterBattle].damage
		engine.playerTurn = true
	}
}