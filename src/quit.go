package main

import "github.com/gen2brain/raylib-go/raylib"

func quit() {
	rl.UnloadTexture(sprite.grass)
	rl.UnloadTexture(sprite.hill)
	rl.UnloadTexture(sprite.fence)
	rl.UnloadTexture(sprite.house)
	rl.UnloadTexture(sprite.water)
	rl.UnloadTexture(sprite.tilled)

	rl.UnloadTexture(player.Sprite)

	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	
	rl.CloseWindow()
}
