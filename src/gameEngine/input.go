package gameEngine

import "github.com/gen2brain/raylib-go/raylib"

func (engine *EngineStruct) input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) { // si on appuie sur Z ou flèche du bas on change la direction (voir le spriteSheet) et on définit qu'on bouge
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
	if rl.IsKeyPressed(rl.KeyE) { // on affiche l'inventaire ou désaffiche l'inventaire en appuyant sur E
		engine.character.showInventory = !engine.character.showInventory
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		engine.player.Speed = 3
		engine.bord.colisionList = [][]float32{}
	} else {
		engine.player.Speed = 1
	}
}
