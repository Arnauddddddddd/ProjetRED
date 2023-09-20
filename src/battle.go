package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func battle() {
			
	if monster[engine.monsterBattle].speed > character.speed {
		engine.playerTurn = false
	}

	if engine.playerTurn {
		if rl.IsMouseButtonPressed(rl.MouseRightButton) {
			monster[engine.monsterBattle].hp -= character.damage
			engine.playerTurn = false
		}
		if monster[engine.monsterBattle].hp <= 0 {
			monster[engine.monsterBattle].alive = false
			engine.battle = false
			engine.cam.Zoom = 3.5
		}
		if character.hp <= 0 {
			character.alive = false
			engine.battle = false
			engine.cam.Zoom = 3.5
		}
	} else {
		character.hp -= monster[engine.monsterBattle].damage
		engine.playerTurn = true
	}
}