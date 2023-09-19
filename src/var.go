package main

import "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth = 1000
	screenHeight = 480
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
}

type monsterStruct struct {
	name string
	hp int
	damage int
	alive bool
	coord []int
}

type engineStruct struct {
	run bool
	bgColor rl.Color

	framCount int

	musicPaused bool
	music rl.Music

	cam rl.Camera2D
}


var (
	
	player playerStruct
	bord mapStruct
	sprite spriteStruct
	engine engineStruct
	monster []monsterStruct

)

