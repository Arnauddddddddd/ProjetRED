package main

import "github.com/gen2brain/raylib-go/raylib"

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)

	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	
	rl.CloseWindow()
}
