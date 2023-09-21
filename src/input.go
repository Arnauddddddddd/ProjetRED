package main

import "github.com/gen2brain/raylib-go/raylib"

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		engine.player.Moving = true
		engine.player.Dir = 1
		engine.player.Up = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		engine.player.Moving = true
		engine.player.Dir = 0
		engine.player.Down = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		engine.player.Moving = true
		engine.player.Dir = 2
		engine.player.Left = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		engine.player.Moving = true
		engine.player.Dir = 3
		engine.player.Right = true
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		engine.musicPaused = !engine.musicPaused
	}
	if !engine.battle {
		if rl.IsKeyPressed(rl.KeyE) {
			engine.character.showInventory = !engine.character.showInventory
		}
	}
}
