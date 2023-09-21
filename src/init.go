package main

import "github.com/gen2brain/raylib-go/raylib"

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Le Jeu")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	engine.run = true
	engine.bgColor = rl.NewColor(147, 211, 196, 255)
	engine.menuSelector = true
	engine.battle = false
	engine.playerTurn = true

	engine.player.Sprite = rl.LoadTexture("../texture/GodotProject/World/Actor/Player/GreenNinja/SpriteSheet.png")
	engine.player.Src = rl.NewRectangle(0, 0, 16, 16)
	engine.player.Dest = rl.NewRectangle(570, 700, 16, 16)

	engine.player.Speed = 1

	engine.monster = append(engine.monster, monsterStruct{"slime", 180, 200, 30, 0, true, rl.LoadTexture("../texture/res/Characters/slime.png"), rl.NewRectangle(0, 0, 32, 32), rl.NewRectangle(520, 700, 32, 32), 0, 0})

	engine.sprite.grass = rl.LoadTexture("../texture/res/Tilesets/Grass.png")
	engine.sprite.hill = rl.LoadTexture("../texture/res/Tilesets/Hills.png")
	engine.sprite.fence = rl.LoadTexture("../texture/res/Tilesets/Fences.png")
	engine.sprite.house = rl.LoadTexture("../texture/res/Tilesets/Wooden House.png")
	engine.sprite.water = rl.LoadTexture("../texture/res/Tilesets/Water.png")
	engine.sprite.tilled = rl.LoadTexture("../texture/res/Tilesets/Tilled Dirt.png")
	engine.sprite.invBar = rl.LoadTexture("../texture/Free Inventory/Inventory_Bar.png")
	engine.sprite.heart = rl.LoadTexture("../texture/PropsInPixels_16x/heart.png")
	engine.sprite.heartContainer = rl.LoadTexture("../texture/PropsInPixels_16x/heartContainer.png")
	engine.sprite.money = rl.LoadTexture("../texture/PropsInPixels_16x/money.png")
	engine.sprite.layer = rl.LoadTexture("../texture/calque.png")
	engine.sprite.bgForest = rl.LoadTexture("../texture/battle/PNG/game_background_4/game_background_4.png")
	engine.sprite.buttonAttack = rl.LoadTexture("../texture/Pixel Buttons pack 4.png")

	engine.bord.tileDest = rl.NewRectangle(0, 0, 16, 16)
	engine.bord.tileSrc = rl.NewRectangle(0, 0, 16, 16)

    engine.sprite.buttonMenu = rl.LoadTexture("../texture/button_menu/png/Buttons/Rect-Text-Blue/Play-Idle.png")
    engine.sprite.buttonMenuPressed = rl.LoadTexture("../texture/button_menu/png/Buttons/Rect-Text-Blue/Play-Click.png")

	engine.sprite.buttonPlay = rl.LoadTexture("../texture/button_menu/png/Buttons/Rect-Icon-Blue/Play-Click.png")
	engine.sprite.buttonPlayPressed = rl.LoadTexture("../texture/button_menu/png/Buttons/Rect-Icon-Blue/Play-Idle.png")

	rl.InitAudioDevice()
	engine.music = rl.LoadMusicStream("texture/res/music.mp3")
	engine.musicPaused = false
	rl.PlayMusicStream(engine.music)

	engine.cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(engine.player.Dest.X - (engine.player.Dest.Width / 2)), float32(engine.player.Dest.Y - (engine.player.Dest.Height/2))), 0.0, 3.5)

	loadMap("../map.txt")
}
