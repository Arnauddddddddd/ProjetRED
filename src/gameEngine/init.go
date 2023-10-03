package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

	showHud bool
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
	donjon rl.Texture2D
	donjon2 rl.Texture2D
	house rl.Texture2D
	water rl.Texture2D
	tilled rl.Texture2D
	texture rl.Texture2D
	buttonPlay rl.Texture2D
	buttonPlayPressed rl.Texture2D
	buttonMenu rl.Texture2D
    buttonMenuPressed rl.Texture2D
	invBar rl.Texture2D
	redHeart rl.Texture2D
	yellowHeart rl.Texture2D
	money rl.Texture2D
	layer rl.Texture2D
	bgForest rl.Texture2D
	bgDesert rl.Texture2D
	bgDungeon rl.Texture2D
}

type monsterStruct struct {
	name string
	hp float32
	hpMax float32
	damage float32
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
	goldLoot float32
	increase float32
	where string
	turn bool
}

type itemStruct struct {
	name string
	gender string
	description string
	sprite rl.Texture2D
	damageUp float32
	hpUp float32
	speedUp float32
	outBattle bool
	battle bool
	infinitySale bool
	price float32
}

type charcacterStruct struct {
	name string
	hp float32
	hpMax float32
	damage float32
	damageBase float32
	speed int
	gold float32
	inventory []itemStruct
	showInventory bool
	showText bool
	alive bool
}

type shopStruct struct {
	items []itemStruct
	sprite rl.Texture2D
	Src rl.Rectangle
	Dest rl.Rectangle
	shopSprite rl.Texture2D
	showPrice []int
}

type menuStruct struct {
	sprite rl.Texture2D
	Dest rl.Rectangle
	Src rl.Rectangle
	frameCount int
}

type textBoxStruct struct {
	sprite rl.Texture2D
	space rl.Texture2D
	frameCountSpace int
	frameCountText int
	textWriting bool
	textToPrint string
	textPrint string
}

type battleStruct struct {
	inBattle bool
	monsterBattle int
	slash bool
	buttonBattleAttack rl.Texture2D
	buttonBattleFattality rl.Texture2D
	buttonBattlePressed []rl.Texture2D
	slashSprite rl.Texture2D
	slashFrameCount int
	slashSrc rl.Rectangle
	fatality string
}

func (engine *EngineStruct) init() { // initialisation général
	rl.InitWindow(screenWidth, screenHeight, "Le Jeu")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	engine.run = true
	engine.bgColor = rl.NewColor(147, 211, 196, 255)
	engine.menuSelector = true
	engine.menuPlay = true
	engine.battle.inBattle = false
	engine.playerTurn = true

	engine.fontText = rl.LoadFont("../texture/txt/Dico.ttf")
	engine.fontNum = rl.LoadFont("../texture/txt/prstart.ttf")

	engine.player.Src = rl.NewRectangle(0, 0, 16, 16)
	engine.player.Dest = rl.NewRectangle(1057, 1633, 16, 16)

	engine.player.Speed = 1.0


	engine.initMonster()
	engine.initSprite()
	engine.initShopKeeper()


	engine.menu.Src = rl.NewRectangle(0, 0, 500, 267)
	engine.menu.Dest = rl.NewRectangle(float32(screenWidth), float32(screenHeight), 1600, 900)


	engine.bord.tileDest = rl.NewRectangle(0, 0, 16, 16)
	engine.bord.tileSrc = rl.NewRectangle(0, 0, 16, 16)

	engine.cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(engine.player.Dest.X - (engine.player.Dest.Width / 2)), float32(engine.player.Dest.Y - (engine.player.Dest.Height/2))), 0.0, 4.0)

	engine.loadMap("../map.txt")
}

func (engine *EngineStruct) initMonster() { // On initialise tout les monstres

	engine.monster = append(engine.monster, monsterStruct{"slime", 60, 60, 30, 0, true, rl.LoadTexture("../texture/monster/slime.png"), rl.NewRectangle(0, 0, 32, 32), rl.NewRectangle(1577, 1596, 32, 32), 0, 3, 0, 0, 8, 50, 6, "jungle", false})
	engine.monster = append(engine.monster, monsterStruct{"slime", 60, 60, 30, 0, true, rl.LoadTexture("../texture/monster/slime.png"), rl.NewRectangle(0, 0, 32, 32), rl.NewRectangle(1437, 1448, 32, 32), 0, 3, 0, 0, 8, 50, 6,"jungle", false})
	engine.monster = append(engine.monster, monsterStruct{"l'escargos", 100, 100, 40, 0, true, rl.LoadTexture("../texture/monster/donjon2.2.png"), rl.NewRectangle(5, 370, 16, 30), rl.NewRectangle(1701, 1073, 16, 30), 0, 7, 0, 2, 8, 70, 6,"jungle", true})
	engine.monster = append(engine.monster, monsterStruct{"l'escargos", 100, 100, 40, 0, true, rl.LoadTexture("../texture/monster/donjon2.2.png"), rl.NewRectangle(5, 370, 16, 30), rl.NewRectangle(1512, 800, 16, 30), 0, 7, 0, 2, 8, 70, 6,"jungle", true})
	engine.monster = append(engine.monster, monsterStruct{"croco", 130, 130, 25, 0, true, rl.LoadTexture("../texture/monster/donjon2.2.png"), rl.NewRectangle(0, 230, 16, 32), rl.NewRectangle(1631, 924, 16, 32), 0, 2, 0, 802, 16, 80, 6,"jungle", true})
	engine.monster = append(engine.monster, monsterStruct{"wizard", 80, 80, 60, 0, true, rl.LoadTexture("../texture/monster/donjon.png"), rl.NewRectangle(5, 138, 16, 25), rl.NewRectangle(1386, 910, 16, 25), 0, 7, 0, 128, 8, 100, 10,"jungle", false})
	engine.monster = append(engine.monster, monsterStruct{"blue knight", 150, 150, 60, 0, true, rl.LoadTexture("../texture/monster/donjon.png"), rl.NewRectangle(0, 100, 16, 32), rl.NewRectangle(1250, 1100, 16, 32), 0, 5, 0, 128, 8, 200, 6,"jungle", false})
	
	engine.monster = append(engine.monster, monsterStruct{"Pigmen", 200, 200, 60, 0, true, rl.LoadTexture("../texture/monster/Pig_Big.png"), rl.NewRectangle(0, 60, 64, 38), rl.NewRectangle(443, 1029, 64, 38), 0, 4, 0, 0, 8, 300, 6,"desert", false})
	engine.monster = append(engine.monster, monsterStruct{"Pigmen", 200, 200, 60, 0, true, rl.LoadTexture("../texture/monster/Pig_Big.png"), rl.NewRectangle(0, 60, 64, 38), rl.NewRectangle(399, 1185, 64, 38), 0, 4, 0, 0, 8, 300, 6,"desert", false})
	engine.monster = append(engine.monster, monsterStruct{"Pigmen", 200, 200, 60, 0, true, rl.LoadTexture("../texture/monster/Pig_Big.png"), rl.NewRectangle(0, 60, 64, 38), rl.NewRectangle(287, 1367, 64, 38), 0, 4, 0, 0, 8, 300, 6,"desert", false})
	engine.monster = append(engine.monster, monsterStruct{"Pigmen", 200, 200, 60, 0, true, rl.LoadTexture("../texture/monster/Pig_Big.png"), rl.NewRectangle(0, 60, 64, 38), rl.NewRectangle(700, 1370, 64, 38), 0, 4, 0, 0, 8, 300, 6,"desert", false})
	engine.monster = append(engine.monster, monsterStruct{"el diablos2", 60, 60, 100, 0, true, rl.LoadTexture("../texture/monster/donjon.png"), rl.NewRectangle(0, 272, 16, 24), rl.NewRectangle(923, 1310, 16, 24), 0, 7, 0, 368, 8, 150,  6,"desert", false})
	engine.monster = append(engine.monster, monsterStruct{"el diablos2", 60, 60, 100, 0, true, rl.LoadTexture("../texture/monster/donjon.png"), rl.NewRectangle(0, 272, 16, 24), rl.NewRectangle(409, 1425, 16, 24), 0, 7, 0, 368, 8, 150,  6,"desert", false})
	engine.monster = append(engine.monster, monsterStruct{"knight", 500, 500, 90, 0, true, rl.LoadTexture("../texture/monster/DarkKnight.png"), rl.NewRectangle(0, 50, 64, 50), rl.NewRectangle(741, 929, 64, 50), 0, 2, 4, 3, 8, 500, 6,"desert", false})
	engine.monster = append(engine.monster, monsterStruct{"knight", 500, 500, 90, 0, true, rl.LoadTexture("../texture/monster/DarkKnight.png"), rl.NewRectangle(0, 50, 64, 50), rl.NewRectangle(791, 929, 64, 50), 0, 2, 4, 3, 8, 500, 6,"desert", false})
	
	engine.monster = append(engine.monster, monsterStruct{"demon", 600, 600, 98, 0, true, rl.LoadTexture("../texture/monster/demon.png"), rl.NewRectangle(0, 0, 32, 42), rl.NewRectangle(328, 768, 32, 42), 0, 2, 8, 0, 9, 1000,  6,"dungeon", false})
	engine.monster = append(engine.monster, monsterStruct{"goblin", 700, 700, 120, 0, true, rl.LoadTexture("../texture/monster/donjon2.2.png"), rl.NewRectangle(0, 380, 32, 40), rl.NewRectangle(1113, 680, 32, 40), 0, 2, 0, 848, 16, 1500,  6,"dungeon", true})
	engine.monster = append(engine.monster, monsterStruct{"goblin", 700, 700, 120, 0, true, rl.LoadTexture("../texture/monster/donjon2.2.png"), rl.NewRectangle(0, 380, 32, 40), rl.NewRectangle(1113, 820, 32, 40), 0, 2, 0, 848, 16, 1500,  6,"dungeon", true})
	engine.monster = append(engine.monster, monsterStruct{"golem", 750, 750, 170, 0, true, rl.LoadTexture("../texture/monster/donjon2.2.png"), rl.NewRectangle(0, 326, 32, 40), rl.NewRectangle(931, 670, 32, 40), 0, 2, 0, 848, 12, 1750,  6,"dungeon", true})
	engine.monster = append(engine.monster, monsterStruct{"samurai", 800, 800, 200, 0, true, rl.LoadTexture("../texture/monster/Samurai.png"), rl.NewRectangle(0, 50, 64, 50), rl.NewRectangle(507, 607, 64, 50), 0, 7, 0, 50, 12, 2000,  6,"dungeon", false})
	engine.monster = append(engine.monster, monsterStruct{"knight", 900, 800, 250, 0, true, rl.LoadTexture("../texture/monster/DarkKnight.png"), rl.NewRectangle(0, 50, 64, 50), rl.NewRectangle(335, 510, 64, 50), 0, 9, 9, 3, 10, 2000,  6,"dungeon", false})
	engine.monster = append(engine.monster, monsterStruct{"slime", 999, 999, 400, 0, true, rl.LoadTexture("../texture/monster/slime2.png"), rl.NewRectangle(0, 0, 32, 32), rl.NewRectangle(1577, 410, 32*8, 32*8), 0, 3, 0, 320, 8, 12,  40,"dungeon", false})
	}

func (engine *EngineStruct) initSprite() { // On initialise toutes les texture / sprites
	engine.sprite.grass = rl.LoadTexture("../texture/map/forest_.png")
	engine.sprite.donjon = rl.LoadTexture("../texture/map/donjon.png")
	engine.sprite.donjon2 = rl.LoadTexture("../texture/map/Dungeon_Tileset.png")
	engine.sprite.house = rl.LoadTexture("../texture/map/Wooden House.png")
	engine.sprite.water = rl.LoadTexture("../texture/map/Water.png")
	engine.sprite.tilled = rl.LoadTexture("../texture/map/desert_.png")
	engine.sprite.invBar = rl.LoadTexture("../texture/inventory/Inventory_Example_04.png")
	engine.sprite.redHeart = rl.LoadTexture("../texture/shop/HeartsRed.png")
	engine.sprite.yellowHeart = rl.LoadTexture("../texture/shop/HeartsYellow.png")
	engine.sprite.money = rl.LoadTexture("../texture/shop/money.png")
	engine.sprite.layer = rl.LoadTexture("../texture/calque.png")
	engine.sprite.bgForest = rl.LoadTexture("../texture/battle/game_background_4.png")
	engine.sprite.bgDesert = rl.LoadTexture("../texture/battle/game_background_1.png")
	engine.sprite.bgDungeon = rl.LoadTexture("../texture/battle/game_background_3.png")
	engine.battle.buttonBattleAttack = rl.LoadTexture("../texture/buttons/2204_w017_n001_439a_p30_439-removebg-preview.png")
	engine.battle.buttonBattleFattality = rl.LoadTexture("../texture/buttons/2204_w017_n001_439a_p30_439-removebg-preview.png")
	engine.battle.buttonBattlePressed = []rl.Texture2D{rl.LoadTexture("../texture/buttons/2204_w017_n001_439a_p30_439-removebg-preview.png"), rl.LoadTexture("../texture/buttons/button-removebg-preview.png")}
	engine.battle.slashSprite = rl.LoadTexture("../texture/slash2-PhotoRoom.png-PhotoRoom.png")
	engine.sprite.buttonMenu = rl.LoadTexture("../texture/buttons/Play-Idle2.png")
    engine.sprite.buttonMenuPressed = rl.LoadTexture("../texture/buttons/Play-Click2.png")
	engine.sprite.buttonPlay = rl.LoadTexture("../texture/buttons/Play-Click.png")
	engine.sprite.buttonPlayPressed = rl.LoadTexture("../texture/buttons/Play-Idle.png")

	engine.shopKeeper.shopSprite = rl.LoadTexture("../texture/inventory/Inventory_Example_03.png")
	engine.shopKeeper.sprite = rl.LoadTexture("../texture/shop/AnimationSheet_Character.png")
	
	engine.textBox.sprite = rl.LoadTexture("../texture//buttons/buttons_2x.png")
	engine.textBox.space = rl.LoadTexture("../texture/buttons/ps_style1.png")

	engine.menu.sprite = rl.LoadTexture("../texture/output-onlinegiftools.png")

	engine.player.Sprite = rl.LoadTexture("../texture/player.png")
}

func (engine *EngineStruct) initShopKeeper() { // On initialise tout les items du shop
	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Heal Potion", "Potion", "A simple magic potion that restores you 50 hp", rl.LoadTexture("../texture/shop/potion.png"), 0, 50, 0, true, true, true, 20})
	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Strengt item", "Fork", "Use this fork in combat to inflict \n100 hp on your opponent.", rl.LoadTexture("../texture/shop/fork.png"), 0, 0, 0, false, true, true, 100})
	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Improvement Item", "Improvement", "EVOLVE : to permanently\n increase your damage and health by 20%.", rl.LoadTexture("../texture/shop/star.png"), 0, 0, 0, true, false, true, 100})
	engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Key", "special", "The key to access the devastated lands", rl.LoadTexture("../texture/shop/key.png"), 50, 50, 0, false, false, false, 200})
	
	//engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Strengt item", "Flask", "By using this item, your damage will be \ndoubled during your fight, if you \ntook a fork or an ax before, this additional \ndamage will also be doubled", rl.LoadTexture("../texture/shop/flask.png"), engine.character.damageBase, 0, 0, false, true, true, 1})
	//engine.shopKeeper.items = append(engine.shopKeeper.items, itemStruct{"Dead item", "dead", "By using this item, your damage will be \ndoubled during your fight, if you \ntook a fork or an ax before, this additional \ndamage will also be doubled", rl.LoadTexture("../texture/shop/dead2.png"), 0, 0, 0, false, false, true, 1})
	
	engine.battle.slashSrc = rl.NewRectangle(0, 0, 70, 39)

	engine.shopKeeper.Src = rl.NewRectangle(0, 0, 32, 32)
	engine.shopKeeper.Dest = rl.NewRectangle(1624, 1269, 32, 32)
	//engine.shopKeeper.Dest = rl.NewRectangle(686, 1153, 32, 32)

	engine.shopKeeper.showPrice = append(engine.shopKeeper.showPrice, 0)
	engine.shopKeeper.showPrice = append(engine.shopKeeper.showPrice, 0)

	engine.shop = false
}