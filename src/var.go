package main

import "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth = 1000
	screenHeight = 480
)

type playerStruct struct {
	playerSprite rl.Texture2D

	playerSrc rl.Rectangle
	playerDest rl.Rectangle

	playerMoving bool
	playerDir int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrame int

	playerSpeed float32 
}

var (
	run = true
	bgColor = rl.NewColor(147, 211, 196, 255)

	player playerStruct

	grassSprite rl.Texture2D

	framCount int

	musicPaused bool
	music rl.Music

	cam rl.Camera2D

    buttonMenu rl.Texture2D
    buttonMenuPressed rl.Texture2D

	buttonPlay rl.Texture2D
	buttonPlayPressed rl.Texture2D

)

