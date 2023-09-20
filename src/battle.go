package main

import "github.com/gen2brain/raylib-go/raylib"

func battle(monster monsterStruct) {
	play := "player"
	rl.ClearBackground(engine.bgColor)
	rl.DrawTexture(sprite.bgForest, 0, 0, rl.White)
	if monster.speed > character.speed {
		play = "monster"
	}
	if play == "player" {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			monster.hp -= character.damage
			play = "monster"
		}
	}
	if play == "monster" {
		character.hp -= monster.damage
		play = "player"
	}
}