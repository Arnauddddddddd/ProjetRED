package main

import "github.com/gen2brain/raylib-go/raylib"

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Le Jeu")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("../res/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("../res/Characters/Basic Charakter Spritesheet.png")
    buttonMenu = rl.LoadTexture("../button_menu/png/Buttons/Rect-Text-Blue/Play-Idle.png")
    buttonMenuPressed = rl.LoadTexture("../button_menu/png/Buttons/Rect-Text-Blue/Play-Click.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)



	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/music.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(playerDest.X - (playerDest.Width / 2)), float32(playerDest.Y - (playerDest.Height/2))), 0.0, 1.0)

}
