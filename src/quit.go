package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func quit() {
	fmt.Println(bord.colisionList)
	rl.UnloadTexture(sprite.grass)
	rl.UnloadTexture(sprite.hill)
	rl.UnloadTexture(sprite.fence)
	rl.UnloadTexture(sprite.house)
	rl.UnloadTexture(sprite.water)
	rl.UnloadTexture(sprite.tilled)

	rl.UnloadTexture(player.Sprite)

	rl.UnloadMusicStream(engine.music)
	rl.CloseAudioDevice()
	
	rl.CloseWindow()
}
