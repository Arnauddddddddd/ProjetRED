package main

import "github.com/gen2brain/raylib-go/raylib"

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Le Jeu")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	engine.run = true
	engine.bgColor = rl.NewColor(147, 211, 196, 255)
	engine.menuSelector = true

	player.Sprite = rl.LoadTexture("../texture/GodotProject/World/Actor/Player/GreenNinja/SpriteSheet.png")
	player.Src = rl.NewRectangle(0, 0, 16, 16)
	player.Dest = rl.NewRectangle(670, 600, 16, 16)

	player.Speed = 1

	monster = append(monster, monsterStruct{"slime", 180, 30, true, rl.LoadTexture("../texture/res/Characters/slime.png"), rl.NewRectangle(0, 0, 32, 32), rl.NewRectangle(720, 600, 32, 32), 0})

	sprite.grass = rl.LoadTexture("../texture/res/Tilesets/Grass.png")
	sprite.hill = rl.LoadTexture("../texture/res/Tilesets/Hills.png")
	sprite.fence = rl.LoadTexture("../texture/res/Tilesets/Fences.png")
	sprite.house = rl.LoadTexture("../texture/res/Tilesets/Wooden House.png")
	sprite.water = rl.LoadTexture("../texture/res/Tilesets/Water.png")
	sprite.tilled = rl.LoadTexture("../texture/res/Tilesets/Tilled Dirt.png")
	sprite.invBar = rl.LoadTexture("../texture/Free Inventory/Inventory_Bar.png")
	sprite.heart = rl.LoadTexture("../texture/PropsInPixels_16x/heart.png")
	sprite.heartContainer = rl.LoadTexture("../texture/PropsInPixels_16x/heartContainer.png")
	sprite.money = rl.LoadTexture("../texture/PropsInPixels_16x/money.png")
	sprite.layer = rl.LoadTexture("../texture/calque.png")

	bord.tileDest = rl.NewRectangle(0, 0, 16, 16)
	bord.tileSrc = rl.NewRectangle(0, 0, 16, 16)

    sprite.buttonMenu = rl.LoadTexture("../texture/button_menu/png/Buttons/Rect-Text-Blue/Play-Idle.png")
    sprite.buttonMenuPressed = rl.LoadTexture("../texture/button_menu/png/Buttons/Rect-Text-Blue/Play-Click.png")

	sprite.buttonPlay = rl.LoadTexture("../texture/button_menu/png/Buttons/Rect-Icon-Blue/Play-Click.png")
	sprite.buttonPlayPressed = rl.LoadTexture("../texture/button_menu/png/Buttons/Rect-Icon-Blue/Play-Idle.png")

	rl.InitAudioDevice()
	engine.music = rl.LoadMusicStream("texture/res/music.mp3")
	engine.musicPaused = false
	rl.PlayMusicStream(engine.music)

	engine.cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(player.Dest.X - (player.Dest.Width / 2)), float32(player.Dest.Y - (player.Dest.Height/2))), 0.0, 3.5)

	loadMap("../map.txt")
}
