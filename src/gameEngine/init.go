package gameEngine

import "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth = 1600
	screenHeight = 900
)

type playerStruct struct {
	Sprite rl.Texture2D

	Src rl.Rectangle
	Dest rl.Rectangle

	Moving bool
	Dir int
	Up, Down, Right, Left bool
	Frame int

	Speed float32 
}

type mapStruct struct {
	tileDest rl.Rectangle // où sur l'écran
	tileSrc rl.Rectangle // où sur l'image
	tileMap []int
	srcMap []string
	mapW, mapH int
	colisionList [][]float32
}

type spriteStruct struct {
	grass rl.Texture2D
	hill rl.Texture2D
	fence rl.Texture2D
	house rl.Texture2D
	water rl.Texture2D
	tilled rl.Texture2D
	texture rl.Texture2D
	buttonPlay rl.Texture2D
	buttonPlayPressed rl.Texture2D
	buttonMenu rl.Texture2D
    buttonMenuPressed rl.Texture2D
	invBar rl.Texture2D
	heart rl.Texture2D
	heartContainer rl.Texture2D
	money rl.Texture2D
	layer rl.Texture2D
	bgForest rl.Texture2D
	buttonBattle rl.Texture2D
	gameName rl.Texture2D
	shopName rl.Texture2D
}

type monsterStruct struct {
	name string
	hp int
	hpMax int
	damage int
	speed int
	alive bool
	sprite rl.Texture2D
	Src rl.Rectangle
	Dest rl.Rectangle
	frameCount int
	frameNumber int
	deadTime int
	Xstart int
	speedFrame int
	goldLoot int
}

type itemStruct struct {
	name string
	gender string
	description string
	sprite rl.Texture2D
	damageUp int
	hpUp int
	speedUp int
	outBattle bool
	battle bool
	price int
}

type charcacterStruct struct {
	name string
	hp int
	hpMax int
	damage int
	damageBase int
	speed int
	class string
	gold int
	inventory []itemStruct
	showInventory bool
	alive bool
}

type shopStruct struct {
	name string
	items []itemStruct
	sprite rl.Texture2D
	Src rl.Rectangle
	Dest rl.Rectangle
	shopSprite rl.Texture2D
}

type menuStruct struct {
	sprite rl.Texture2D
	Dest rl.Rectangle
	Src rl.Rectangle
	frameCount int
}

func initt(engine *EngineStruct) {
	rl.InitWindow(screenWidth, screenHeight, "Le Jeu")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	engine.run = true
	engine.bgColor = rl.NewColor(147, 211, 196, 255)
	engine.menuSelector = true
	engine.battle = false
	engine.playerTurn = true

	engine.player.Sprite = rl.LoadTexture("../texture/GodotProject/World/Actor/Npc/Warrior/SpriteSheet.png")
	engine.player.Src = rl.NewRectangle(0, 0, 16, 16)
	engine.player.Dest = rl.NewRectangle(570, 300, 16, 16)

	engine.player.Speed = 1


	engine.monster = append(engine.monster, monsterStruct{"slime", 180, 100, 30, 0, true, rl.LoadTexture("../texture/res/Characters/slime.png"), rl.NewRectangle(0, 0, 32, 32), rl.NewRectangle(520, 700, 32, 32), 0, 3, 0, 0, 8, 12})
	engine.monster = append(engine.monster, monsterStruct{"knight", 200, 200, 80, 0, true, rl.LoadTexture("../texture/monster/DarkKnight.png"), rl.NewRectangle(0, 50, 64, 50), rl.NewRectangle(500, 650, 64, 50), 0, 1, 0, 64, 40, 12})
	engine.monster = append(engine.monster, monsterStruct{"el diablos", 200, 200, 80, 0, true, rl.LoadTexture("../texture/DungeonTilesetII_v1.6/donjon2.2.png"), rl.NewRectangle(0, 64, 16, 16), rl.NewRectangle(770, 600, 16, 16), 0, 7, 0, 0, 8, 12})
	engine.monster = append(engine.monster, monsterStruct{"l'escargos", 200, 200, 80, 0, true, rl.LoadTexture("../texture/DungeonTilesetII_v1.6/donjon2.2.png"), rl.NewRectangle(5, 370, 16, 30), rl.NewRectangle(790, 700, 16, 30), 0, 7, 0, 2, 8, 12})
	engine.monster = append(engine.monster, monsterStruct{"samurai", 200, 200, 80, 0, true, rl.LoadTexture("../texture/monster/Samurai.png"), rl.NewRectangle(0, 50, 64, 50), rl.NewRectangle(800, 650, 64, 50), 0, 9, 0, 0, 8, 12})
	engine.monster = append(engine.monster, monsterStruct{"Pigmen", 200, 200, 80, 0, true, rl.LoadTexture("../texture/monster/Pig_Big.png"), rl.NewRectangle(0, 106, 64, 38), rl.NewRectangle(650, 570, 64, 38), 0, 5, 0, 0, 8, 12})
	engine.monster = append(engine.monster, monsterStruct{"Pigmen", 200, 200, 80, 0, true, rl.LoadTexture("../texture/DungeonTilesetII_v1.6/donjon.png"), rl.NewRectangle(0, 100, 16, 32), rl.NewRectangle(650, 570, 16, 32), 0, 5, 0, 128, 8, 12})

	engine.sprite.grass = rl.LoadTexture("../texture/res/Tilesets/Grass.png")
	engine.sprite.hill = rl.LoadTexture("../texture/res/Tilesets/Hills.png")
	engine.sprite.fence = rl.LoadTexture("../texture/res/Tilesets/Fences.png")
	engine.sprite.house = rl.LoadTexture("../texture/res/Tilesets/Wooden House.png")
	engine.sprite.water = rl.LoadTexture("../texture/res/Tilesets/Water.png")
	engine.sprite.tilled = rl.LoadTexture("../texture/res/Tilesets/Tilled Dirt.png")
	engine.sprite.invBar = rl.LoadTexture("../texture/Retro Inventory/Original/Inventory_Example_04.png")
	engine.sprite.heart = rl.LoadTexture("../texture/PropsInPixels_16x/heart.png")
	engine.sprite.heartContainer = rl.LoadTexture("../texture/PropsInPixels_16x/heartContainer.png")
	engine.sprite.money = rl.LoadTexture("../texture/PropsInPixels_16x/money.png")
	engine.sprite.layer = rl.LoadTexture("../texture/calque.png")
	engine.sprite.bgForest = rl.LoadTexture("../texture/battle/PNG/game_background_4/game_background_4.png")
	engine.sprite.buttonBattle = rl.LoadTexture("../texture/2204_w017_n001_439a_p30_439-removebg-preview.png")
	engine.sprite.gameName = rl.LoadTexture("../texture/Capture_d_écran_du_2023-09-22_15-14-26-removebg-preview.png")
	engine.sprite.shopName = rl.LoadTexture("../texture/Capture_d_écran_du_2023-09-22_16-33-57-removebg-preview.png")

	engine.menu.sprite = rl.LoadTexture("../texture/output-onlinegiftools.png")
	engine.menu.Src = rl.NewRectangle(0, 0, 500, 267)
	engine.menu.Dest = rl.NewRectangle(float32(screenWidth), float32(screenHeight), 1600, 900)

	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
	
	
	engine.shopKeeper.Src = rl.NewRectangle(0, 0, 32, 32)
	engine.shopKeeper.Dest = rl.NewRectangle(500, 600, 32, 32)
	engine.shopKeeper.shopSprite = rl.LoadTexture("../texture/Retro Inventory/Original/Inventory_Example_03.png")
	engine.shopKeeper.sprite = rl.LoadTexture("../texture/AnimationSheet_Character.png")

	engine.shop = false


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

	loadMap(engine, "../map.txt")
}
