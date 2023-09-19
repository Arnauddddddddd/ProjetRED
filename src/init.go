package main

import "github.com/gen2brain/raylib-go/raylib"

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Le Jeu")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	player.Sprite = rl.LoadTexture("../GodotProject/World/Actor/Player/GreenNinja/SpriteSheet.png")
	player.Src = rl.NewRectangle(0, 0, 16, 16)
	player.Dest = rl.NewRectangle(100, 100, 16, 16)

	player.Speed = 1

	sprite.grass = rl.LoadTexture("../res/Tilesets/Grass.png")
	sprite.hill = rl.LoadTexture("../res/Tilesets/Hills.png")
	sprite.fence = rl.LoadTexture("../res/Tilesets/Fences.png")
	sprite.house = rl.LoadTexture("../res/Tilesets/Wooden House.png")
	sprite.water = rl.LoadTexture("../res/Tilesets/Water.png")
	sprite.tilled = rl.LoadTexture("../res/Tilesets/Tilled Dirt.png")

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

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(player.Dest.X - (player.Dest.Width / 2)), float32(player.Dest.Y - (player.Dest.Height/2))), 0.0, 3.5)

	loadMap("../map.txt")
}
