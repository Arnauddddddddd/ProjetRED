package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func quit() {
	rl.UnloadTexture(engine.sprite.grass)
	rl.UnloadTexture(engine.sprite.hill)
	rl.UnloadTexture(engine.sprite.fence)
	rl.UnloadTexture(engine.sprite.house)
	rl.UnloadTexture(engine.sprite.water)
	rl.UnloadTexture(engine.sprite.tilled)
	rl.UnloadTexture(engine.sprite.invBar)

	rl.UnloadTexture(engine.player.Sprite)

	rl.UnloadMusicStream(engine.music)
	rl.CloseAudioDevice()
	
	rl.CloseWindow()
}
