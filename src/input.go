package main

import "github.com/gen2brain/raylib-go/raylib"

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		player.Moving = true
		player.Dir = 1
		player.Up = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		player.Moving = true
		player.Dir = 0
		player.Down = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		player.Moving = true
		player.Dir = 2
		player.Left = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		player.Moving = true
		player.Dir = 3
		player.Right = true
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		engine.musicPaused = !engine.musicPaused
	}
	if !engine.battle {
		if rl.IsKeyPressed(rl.KeyE) {
			character.showInventory = !character.showInventory
		}
	}
}
