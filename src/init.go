package main

import "github.com/gen2brain/raylib-go/raylib"

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Le Jeu")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	player.playerSprite = rl.LoadTexture("../GodotProject/World/Actor/Player/GreenNinja/SpriteSheet.png")
	player.playerSrc = rl.NewRectangle(0, 0, 16, 16)
	player.playerDest = rl.NewRectangle(100, 100, 30, 30)

	player.playerSpeed = 2.5

	grassSprite = rl.LoadTexture("../res/Tilesets/Grass.png")

	bord.tileDest = rl.NewRectangle(0, 0, 16, 16)
	bord.tileSrc = rl.NewRectangle(0, 0, 16, 16)

    buttonMenu = rl.LoadTexture("../button_menu/png/Buttons/Rect-Text-Blue/Play-Idle.png")
    buttonMenuPressed = rl.LoadTexture("../button_menu/png/Buttons/Rect-Text-Blue/Play-Click.png")

	buttonPlay = rl.LoadTexture("../button_menu/png/Buttons/Rect-Icon-Blue/Play-Click.png")
	buttonPlayPressed = rl.LoadTexture("../button_menu/png/Buttons/Rect-Icon-Blue/Play-Idle.png")

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/music.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(player.playerDest.X - (player.playerDest.Width / 2)), float32(player.playerDest.Y - (player.playerDest.Height/2))), 0.0, 2.0)

	loadMap("../map.txt")
}
