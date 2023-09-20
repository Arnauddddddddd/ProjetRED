package main

import (
	"github.com/gen2brain/raylib-go/raylib"
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
}

type monsterStruct struct {
	name string
	hp int
	damage int
	speed int
	alive bool
	sprite rl.Texture2D
	Src rl.Rectangle
	Dest rl.Rectangle
	frameCount int
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
}

type charcacterStruct struct {
	name string
	hp int
	hpMax int
	damage int
	speed int
	class string
	gold int
	inventory []itemStruct
	showInventory bool
	alive bool
}

type engineStruct struct {
	run bool
	bgColor rl.Color

	framCount int

	musicPaused bool
	music rl.Music

	cam rl.Camera2D
	menuSelector bool

	battle bool
	monsterBattle int
}


var (

	//healPotion = itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../PropsInPixels_16x/potion.png"), 0, 40, 0}
	
	player playerStruct
	bord mapStruct
	sprite spriteStruct
	engine engineStruct
	monster []monsterStruct
	character charcacterStruct

)
